package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"sync"
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
		return nil, fmt.Errorf("cusisAuthorizeError")
	}
	if s.span(cusisError) {
		return nil, fmt.Errorf("cusisError")
	}
	return s, nil
}

var reboot func(career, term int) (*http.Client, string, error)

func rowHeadsToFragment(career, term, subject int, rowHeads []rowHeadS) (*fragmentS, error) {
	var courses []courseS
	var groups []groupS
	var teachers []teacherS
	var classes []classS
	var meetings []meetingS
	coursesMap := make(map[string]int)
	teachersMap := make(map[string]int)
	for _, rowHead := range rowHeads {
		if _, ok := coursesMap[rowHead.code]; !ok {
			coursesMap[rowHead.code] = len(courses)
			courses = append(courses, courseS{
				career,
				subject,
				rowHead.code,
				rowHead.title,
				rowHead.units,
				nil,
			})
		}
		course := coursesMap[rowHead.code]
		if courses[course].title != rowHead.title || courses[course].units != rowHead.units {
			return nil, fmt.Errorf(
				"mismatch course\n%v %v\n%v %v",
				courses[course].title,
				courses[course].units,
				rowHead.title,
				rowHead.units,
			)
		}
		group := groupS{course, term, rowHead.group, nil, nil}
		for _, rowBody := range rowHead.rowBody {
			class := classS{
				len(groups),
				"",
				rowBody.quota,
				rowBody.vacancy,
				rowBody.component,
				rowBody.section,
				rowBody.language,
				nil,
				nil,
				rowBody.rowFoot[0].add,
				rowBody.rowFoot[0].drop,
				rowBody.dept,
			}
			for _, rowFoot := range rowBody.rowFoot {
				if rowFoot.add != class.add || rowFoot.drop != class.drop {
					return nil, fmt.Errorf("differ add drop\n%v", rowFoot)
				}
				class.meetings = append(class.meetings, len(meetings))
				meetings = append(meetings, meetingS{
					len(classes),
					rowFoot.period,
					rowFoot.room,
					rowFoot.date,
				})
			}
			group.classes = append(group.classes, len(classes))
			classes = append(classes, class)
		}
		for _, t := range rowHead.teachers {
			if _, ok := teachersMap[t]; !ok {
				teachersMap[t] = len(teachers)
				teachers = append(teachers, teacherS{t, nil})
			}
			teachers[teachersMap[t]].classes = append(teachers[teachersMap[t]].classes, group.classes[0])
			classes[group.classes[0]].teachers = append(classes[group.classes[0]].teachers, teachersMap[t])
		}
		courses[course].groups = append(courses[course].groups, len(groups))
		groups = append(groups, group)
	}
	return &fragmentS{
		courses,
		groups,
		teachers,
		classes,
		meetings,
	}, nil
}

func crawlCourse(career, term, subject int) (*fragmentS, error) {
	c, icsid, err := reboot(career, term)
	if err != nil {
		return nil, fmt.Errorf("reboot: %v", err)
	}
	s, err := responseToParser(c.PostForm(cusis, url.Values{
		"ICSID":                 {icsid},
		"ICAction":              {"CU_RC_TMSR801_SSR_PB_CLASS_SRCH"},
		"CU_RC_TMSR801_SUBJECT": {db.subjects[subject].slug},
	}))
	if err != nil {
		return nil, fmt.Errorf("post search: %v", err)
	}
	rowHeads, err := s.parseCourses(
		icsid,
		db.careers[career].en,
		db.terms[term].en,
		db.subjects[subject].en,
		db.subjects[subject].slug,
	)
	stateNum := 3
	for i, rowHead := range rowHeads {
		if !rowHead.reserves {
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
		_, err = s.parseReserves(
			strconv.Itoa(stateNum),
			icsid,
			db.subjects[subject].slug+rowHead.code+rowHead.group+rowHead.rowBody[0].section,
			rowHead.nbr,
			rowHead.rowBody[0].component,
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
	_, err = rowHeadsToFragment(career, term, subject, rowHeads)
	if err != nil {
		return nil, fmt.Errorf("fragment\n%v", err)
	}
	prevCode := ""
	for i, rowHead := range rowHeads {
		if rowHead.code == prevCode {
			continue
		}
		prevCode = rowHead.code
		s, err := responseToParser(c.PostForm(cusis, url.Values{
			"ICSID":    {icsid},
			"ICAction": {`CRSE_TITLE$` + strconv.Itoa(rowHead.row)},
		}))
		if err != nil {
			return nil, fmt.Errorf("post title: %v", err)
		}
		stateNum++
		detail, err := s.parseDetail(
			strconv.Itoa(stateNum),
			icsid,
			db.subjects[subject].slug,
			rowHead.code,
			db.careers[career].en,
			rowHead.units,
			rowHead.rowBody[0].rowFoot[0].add,
			rowHead.rowBody[0].rowFoot[0].drop,
		)
		detailString := string(s.s)
		if err != nil {
			return nil, fmt.Errorf("parseDetail %v\n%v", i, err)
		}
		if detail.scheduled {
			s, err = responseToParser(c.PostForm(cusis, url.Values{
				"ICSID":    {icsid},
				"ICAction": {"DERIVED_SAA_CRS_SSR_PB_GO"},
			}))
			if err != nil {
				return nil, fmt.Errorf("post sections: %v", err)
			}
			stateNum++
			sections, err := s.parseSections(
				strconv.Itoa(stateNum),
				icsid,
				db.subjects[subject].slug,
				rowHead.code,
				detail.title,
				detailString,
			)
			if err != nil {
				return nil, fmt.Errorf("parseSections %v\n%v", i, err)
			}
			if sections.more {
				s, err = responseToParser(c.PostForm(cusis, url.Values{
					"ICSID":                    {icsid},
					"ICAction":                 {"CLASS_TBL_VW5$fviewall$0"},
					"DERIVED_SAA_CRS_TERM_ALT": {sections.terms[sections.term][0]},
				}))
				if err != nil {
					return nil, fmt.Errorf("post more: %v", err)
				}
				stateNum++
				sections, err = s.parseSections(
					strconv.Itoa(stateNum),
					icsid,
					db.subjects[subject].slug,
					rowHead.code,
					detail.title,
					detailString,
				)
				if err != nil {
					return nil, fmt.Errorf("parseSectionsAll %v\n%v", i, err)
				}
			}
		}
		s, err = responseToParser(c.PostForm(cusis, url.Values{
			"ICSID":    {icsid},
			"ICAction": {`DERIVED_SAA_CRS_RETURN_PB`},
		}))
		if err != nil {
			return nil, fmt.Errorf("post return: %v", err)
		}
		stateNum++
	}
	if err != nil {
		return nil, fmt.Errorf("parseCourses: %v", err)
	}
	return nil, nil
}

func crawlCourses() error {
	type resultS struct {
		c, t, s int
		r       *fragmentS
	}
	results := make(chan resultS)
	jobs := make(chan [3]int)
	errs := make(chan error, 10)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for job := range jobs {
				for retry := 0; ; retry++ {
					r, err := crawlCourse(job[0], job[1], job[2])
					if err == nil {
						results <- resultS{job[0], job[1], job[2], r}
						break
					}
					err = fmt.Errorf(
						"crawlCourses(%v, %v, %v): %v",
						db.careers[job[0]].slug,
						db.terms[job[1]].en,
						db.subjects[job[2]].slug,
						err,
					)
					if retry == 3 {
						errs <- err
						return
					}
					log.Printf("retry %v: %v", retry, err)
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
		for _, ct := range db.careerTerms {
			for _, s := range ct.subjects {
				select {
				case err := <-errs:
					errs <- err
					return
				case jobs <- [3]int{ct.career, ct.term, s}:
				}
			}
		}
	}()
	for range results {
	}
	if err := <-errs; err != nil {
		return err
	}
	return nil
}

func crawlSubject(career, term int) ([][2]string, error) {
	c, icsid, err := reboot(career, term)
	if err != nil {
		return nil, fmt.Errorf("reboot err: %v", err)
	}
	s, err := responseToParser(c.PostForm(cusis, url.Values{
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

func crawlSubjects() error {
	type resultS struct {
		c, t int
		s    [][2]string
	}
	results := make(chan resultS)
	errs := make(chan error, 10)
	jobs := make(chan [2]int)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for job := range jobs {
				for retry := 0; ; retry++ {
					s, err := crawlSubject(job[0], job[1])
					if err == nil {
						results <- resultS{job[0], job[1], s}
						break
					}
					err = fmt.Errorf(
						"crawlSubjects(%v, %v): %v",
						db.careers[job[0]].slug,
						db.terms[job[1]].en,
						err,
					)
					if retry == 3 {
						errs <- err
						return
					}
					log.Printf("retry %v: %v", retry, err)
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
		for c := range db.careers {
			for t := range db.terms {
				select {
				case err := <-errs:
					errs <- err
					return
				case jobs <- [2]int{c, t}:
				}
			}
		}
	}()
	var careerTerms []careerTermS
	var subjects []subjectS
	m := make(map[[2]string]int)
	for r := range results {
		var x []int
		for _, v := range r.s {
			if _, ok := m[v]; !ok {
				m[v] = len(subjects)
				subjects = append(subjects, subjectS{v[0], v[1], nil})
			}
			x = append(x, m[v])
		}
		careerTerms = append(careerTerms, careerTermS{r.c, r.t, x})
	}
	if err := <-errs; err != nil {
		return err
	}
	db.subjects = subjects
	db.careerTerms = careerTerms
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

func crawlTerms() error {
	c, s, err := boot()
	if err != nil {
		return fmt.Errorf("boot: %v", err)
	}
	icsid, err := s.parseBoot()
	if err != nil {
		return fmt.Errorf("parseBoot\n%v", err)
	}
	termString := string(s.s)
	en, err := s.parseTerms()
	if err != nil {
		return fmt.Errorf("parseTerms\n%v", err)
	}
	s, err = responseToParser(c.PostForm(cusis, url.Values{
		"ICSID":       {icsid},
		"#ICDataLang": {"ZHT"},
		"ICAction":    {"#ICDataLang"},
	}))
	if err != nil {
		return fmt.Errorf("post ZHT err: %v", err)
	}
	ch, err := s.parseTermsCH(icsid)
	if err != nil {
		return fmt.Errorf("parseTermsCH err: %v", err)
	}
	if len(en) != len(ch) {
		return fmt.Errorf("len(en) != len(ch)\n%v\n%v", en, ch)
	}
	var terms []termS
	for _, u := range en {
		for _, v := range ch {
			if u[0] == v[0] {
				terms = append(terms, termS{u[0], u[1], v[1], nil})
				break
			}
		}
	}
	if len(terms) != len(en) {
		return fmt.Errorf("differ en ch\n%v\n%v", en, ch)
	}
	db.terms = make([]termS, len(terms))
	copy(db.terms, terms)
	for i, v := range db.terms {
		db.terms[i].en = strings.TrimSpace(v.en)
	}
	reboot = func(career, term int) (*http.Client, string, error) {
		c, s, err := boot()
		if err != nil {
			return nil, "", err
		}
		icsid, err := s.parseBoot()
		if err != nil {
			return nil, "", fmt.Errorf("parseBoot\n%v", err)
		}
		if err = s.equalErr(termString); err != nil {
			return nil, "", fmt.Errorf("termString\n%v", err)
		}
		s, err = responseToParser(c.PostForm(cusis, url.Values{
			"ICSID":                       {icsid},
			"ICAction":                    {"CLASS_SRCH_WRK2_STRM$50$"},
			"CLASS_SRCH_WRK2_ACAD_CAREER": {db.careers[career].slug},
			"CLASS_SRCH_WRK2_STRM$50$":    {terms[term].slug},
		}))
		if err != nil {
			return nil, "", fmt.Errorf("post career term err: %v", err)
		}
		careerOption := toCareerOption(career, db.careers)
		termOption := toTermOption(term, terms)
		if err := s.parseCareerTerm(icsid, careerOption, termOption); err != nil {
			return nil, "", fmt.Errorf("parseCareerTerm\n%v", err)
		}
		return c, icsid, nil
	}
	return nil
}

func crawl() error {
	if err := crawlTerms(); err != nil {
		return fmt.Errorf("crawlTerms err: %v", err)
	}
	if err := crawlSubjects(); err != nil {
		return fmt.Errorf("crawlSubjects err: %v", err)
	}
	if err := crawlCourses(); err != nil {
		return fmt.Errorf("crawlCourses err: %v", err)
	}
	return nil
}
