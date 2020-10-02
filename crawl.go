package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const cusis = `https://cusis.cuhk.edu.hk/psc/public/EMPLOYEE/HRMS/c/CU_SCR_MENU.CU_TMSR801.GBL`

func responseToParser(res *http.Response, err error) (*parser, error) {
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err := res.Body.Close(); err != nil {
		return nil, err
	}
	s := &parser{body}
	if s.equal(cusisAuthorizeError) {
		return nil, errCusisAuthorize
	}
	if s.span(cusisError) {
		return nil, errCusis
	}
	return s, nil
}

var reboot func() (*http.Client, string, error)

func crawlCourse(subject [2]string) ([]rowHeadS, error) {
	c, icsid, err := reboot()
	if err != nil {
		return nil, fmt.Errorf("reboot: %v", err)
	}
	s, err := responseToParser(c.PostForm(cusis, url.Values{
		"ICSID":                 {icsid},
		"ICAction":              {"CU_RC_TMSR801_SSR_PB_CLASS_SRCH"},
		"CU_RC_TMSR801_SUBJECT": {subject[0]},
	}))
	if err != nil {
		return nil, fmt.Errorf("post search: %v", err)
	}
	rowHeads, err := s.parseCourses(
		icsid,
		"Undergraduate",
		thisTerm[1],
		strings.Replace(subject[1], "&", "&amp;", -1),
		subject[0],
	)
	if err != nil {
		return nil, fmt.Errorf("parseCourses: %v", err)
	}
	stateNum := 3
	for i, rowHead := range rowHeads {
		if rowHead.reserves == nil {
			continue
		}
		s, err := responseToParser(c.PostForm(cusis, url.Values{
			"ICSID":    {icsid},
			"ICAction": {`ENRL_CAP$` + strconv.Itoa(rowHead.row)},
		}))
		if err != nil {
			return nil, fmt.Errorf("post cap: %v", err)
		}
		stateNum++
		rowHeads[i].reserves, err = s.parseReserves(
			strconv.Itoa(stateNum),
			icsid,
			subject[0]+rowHead.code+rowHead.group+rowHead.rowBodys[0].section,
			rowHead.nbr,
			rowHead.rowBodys[0].component,
			rowHead.title,
		)
		if err != nil {
			return nil, fmt.Errorf("parseReserves %v\n%v", i, err)
		}
		s, err = responseToParser(c.PostForm(cusis, url.Values{
			"ICSID":    {icsid},
			"ICAction": {`CU_RC_TMSR801_SSR_PB_CLOSE$0`},
		}))
		if err != nil {
			return nil, fmt.Errorf("post return: %v", err)
		}
		stateNum++
	}
	return rowHeads, nil
}

func crawlCourses(subjects [][2]string, workers int) error {
	type resultS struct {
		subjects [2]string
		rowHeads []rowHeadS
	}
	results := make(chan resultS)
	jobs := make(chan [2]string)
	errs := make(chan error, workers)
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for job := range jobs {
				for retry := 0; ; retry++ {
					r, err := crawlCourse(job)
					if err == nil {
						results <- resultS{job, r}
						break
					}
					err = fmt.Errorf("crawlCourses(%v): %v", job, err)
					if retry == 3 {
						errs <- err
						return
					}
					writeLogs <- fmt.Sprintf("retry %v: %v\n", retry, err)
				}
			}
		}()
	}
	go func() {
		defer func() {
			close(jobs)
			wg.Wait()
			close(results)
			close(errs)
		}()
		for _, subject := range subjects {
			select {
			case err := <-errs:
				errs <- err
				return
			case jobs <- subject:
			}
		}
	}()
	changed := false
	for res := range results {
		for _, rowHead := range res.rowHeads {
			reserves := rowHead.reserves
			for _, rowBody := range rowHead.rowBodys {
				s := res.subjects[0] + rowHead.code + rowHead.group + rowBody.section
				history := db[s]
				if history == nil {
					changed = true
					db[s] = &historyS{
						rowHead.group,
						rowHead.title,
						rowBody.quota,
						rowBody.quota - rowBody.vacancy,
						reserves,
						nil,
					}
				} else {
					now := time.Now().Unix()
					changeLen := len(history.changes)
					if history.title != rowHead.title {
						history.changes = append(history.changes, changeS{
							now,
							4,
							history.title,
						})
						history.title = rowHead.title
					}
					if history.quota != rowBody.quota {
						history.changes = append(history.changes, changeS{
							now,
							5,
							history.quota,
						})
						history.quota = rowBody.quota
					}
					if history.enroll != rowBody.quota-rowBody.vacancy {
						history.changes = append(history.changes, changeS{
							now,
							6,
							history.enroll,
						})
						history.enroll = rowBody.quota - rowBody.vacancy
					}
					for i := 0; i < len(history.reserves) || i < len(reserves); i++ {
						if i >= len(history.reserves) {
							history.changes = append(history.changes, changeS{
								now,
								7 + i*3,
								"",
							})
							history.changes = append(history.changes, changeS{
								now,
								8 + i*3,
								0,
							})
							history.changes = append(history.changes, changeS{
								now,
								9 + i*3,
								0,
							})
							history.reserves = append(history.reserves, reserves[i])
							continue
						}
						if i >= len(reserves) {
							if history.reserves[i] != (reserveS{}) {
								history.changes = append(history.changes, changeS{
									now,
									7 + i*3,
									history.reserves[i].major,
								})
								history.changes = append(history.changes, changeS{
									now,
									8 + i*3,
									history.reserves[i].quota,
								})
								history.changes = append(history.changes, changeS{
									now,
									9 + i*3,
									history.reserves[i].enroll,
								})
							}
							history.reserves[i] = reserveS{}
							continue
						}
						if history.reserves[i].major != reserves[i].major {
							history.changes = append(history.changes, changeS{
								now,
								7 + i*3,
								history.reserves[i].major,
							})
							history.reserves[i].major = reserves[i].major
						}
						if history.reserves[i].quota != reserves[i].quota {
							history.changes = append(history.changes, changeS{
								now,
								8 + i*3,
								history.reserves[i].quota,
							})
							history.reserves[i].quota = reserves[i].quota
						}
						if history.reserves[i].enroll != reserves[i].enroll {
							history.changes = append(history.changes, changeS{
								now,
								9 + i*3,
								history.reserves[i].enroll,
							})
							history.reserves[i].enroll = reserves[i].enroll
						}
					}
					if changeLen != len(history.changes) {
						changed = true
					}
				}
				reserves = nil
			}
		}
	}
	if err := <-errs; err != nil {
		return err
	}
	if changed {
		select {
		case githubPush <- struct{}{}:
		default:
		}
	}
	return nil
}

func boot() (*http.Client, *parser, error) {
	j, err := cookiejar.New(nil)
	if err != nil {
		return nil, nil, err
	}
	var c = &http.Client{Jar: j}
	s, err := responseToParser(c.Get(cusis))
	if err != nil {
		return nil, nil, err
	}
	return c, s, nil
}

func crawlSubjects() ([][2]string, error) {
	c, s, err := boot()
	if err != nil {
		return nil, fmt.Errorf("boot: %v", err)
	}
	icsid, err := s.parseBoot()
	if err != nil {
		return nil, fmt.Errorf("parseBoot\n%v", err)
	}
	s, err = responseToParser(c.PostForm(cusis, url.Values{
		"ICSID":                       {icsid},
		"ICAction":                    {`CLASS_SRCH_WRK2_STRM$50$`},
		`CLASS_SRCH_WRK2_ACAD_CAREER`: {"UG"},
		`CLASS_SRCH_WRK2_STRM$50$`:    {thisTerm[0]},
	}))
	if err != nil {
		return nil, fmt.Errorf("postTerm\n%v", err)
	}
	if err := s.parseBoot2(icsid); err != nil {
		return nil, fmt.Errorf("parseBoot2\n%v", err)
	}
	termString := string(s.s)
	term, err := s.parseTerms()
	if err != nil {
		return nil, fmt.Errorf("parseTerms\n%v", err)
	}
	term[1] = strings.TrimSpace(term[1])
	if term[0] != thisTerm[0] {
		return nil, fmt.Errorf("no term\n%v %v", term, thisTerm)
	}
	if thisTerm != *term {
		thisTerm = *term
		if err = githubFetchDB(); err != nil{ 
			return nil, fmt.Errorf("githubFetchDB: %v", err)
		}
	}
	reboot = func() (*http.Client, string, error) {
		c, s, err := boot()
		if err != nil {
			return nil, "", err
		}
		icsid, err := s.parseBoot()
		if err != nil {
			return nil, "", fmt.Errorf("parseBoot\n%v", err)
		}
		s, err = responseToParser(c.PostForm(cusis, url.Values{
			"ICSID":                       {icsid},
			"ICAction":                    {`CLASS_SRCH_WRK2_STRM$50$`},
			`CLASS_SRCH_WRK2_ACAD_CAREER`: {"UG"},
			`CLASS_SRCH_WRK2_STRM$50$`:    {thisTerm[0]},
		}))
		if err != nil {
			return nil, "", fmt.Errorf("postTerm\n%v", err)
		}
		if err := s.parseBoot2(icsid); err != nil {
			return nil, "", fmt.Errorf("parseBoot2\n%v", err)
		}
		if err = s.equalErr(termString); err != nil {
			return nil, "", fmt.Errorf("termString\n%v", err)
		}
		return c, icsid, nil
	}
	s, err = responseToParser(c.PostForm(cusis, url.Values{
		"ICSID":    {icsid},
		"ICAction": {"CU_RC_TMSR801_SUBJECT$prompt"},
	}))
	if err != nil {
		return nil, fmt.Errorf("post prompt err: %v", err)
	}
	subjects, err := s.parseSubjects(icsid)
	if err != nil {
		return nil, fmt.Errorf("parseSubjects\n%v", err)
	}
	if subjects == nil {
		s, err = responseToParser(c.PostForm(cusis, url.Values{
			"ICSID":    {icsid},
			"ICAction": {"#ICViewAll"},
		}))
		subjects, err = s.parseSubjectsAll(icsid)
		if err != nil {
			return nil, fmt.Errorf("parseSubjectsAll\n%v", err)
		}
	}
	return subjects, nil
}

func crawl() {
	var prevWorkers int
	var prevDuration time.Duration
	workers := 64
	for {
		subjects, err := crawlSubjects()
		if err != nil {
			writeLogs <- fmt.Sprintf("crawlSubjects err: %v", err)
			continue
		}
		for {
			t := time.Now()
			if err = crawlCourses(subjects, workers); err != nil {
				writeLogs <- fmt.Sprintf("crawlCourses err: %v", err)
				continue
			}
			duration := time.Since(t)
			writeLogs <- fmt.Sprintf("%v workers, %v seconds\n", workers, duration.Seconds())
			if prevWorkers == 0 {
				workers, prevWorkers = workers+1, workers
			} else if prevDuration < duration {
				workers, prevWorkers = prevWorkers, workers
			} else if prevWorkers < workers {
				workers, prevWorkers = workers+1, workers
			} else {
				workers, prevWorkers = workers-1, workers
			}
			prevDuration = duration
		}
	}
}
