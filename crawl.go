package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
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

var reboot func() (*http.Client, string, error)

func crawlCourses() error {
	f := func(career, term, subject int) ([]rowHeadS, error) {
		c, icsid, err := reboot()
		if err != nil {
			return nil, fmt.Errorf("reboot: %v", err)
		}
		s, err := responseToParser(c.PostForm(cusis, url.Values{
			"ICSID":                       {icsid},
			"ICAction":                    {"CLASS_SRCH_WRK2_STRM$50$"},
			"CLASS_SRCH_WRK2_ACAD_CAREER": {db.careers[career].slug},
			"CLASS_SRCH_WRK2_STRM$50$":    {db.terms[term].slug},
		}))
		if err != nil {
			return nil, fmt.Errorf("post career term: %v", err)
		}
		careerOption := toCareerOption(db.careers[career].slug, db.careers)
		termOption := toTermOption(db.terms[term].slug, db.terms)
		if err := s.parseCareerTerm(icsid, careerOption, termOption); err != nil {
			return nil, fmt.Errorf("parseCareerTerm: %v", err)
		}
		s, err = responseToParser(c.PostForm(cusis, url.Values{
			"ICSID":                 {icsid},
			"ICAction":              {"CU_RC_TMSR801_SSR_PB_CLASS_SRCH"},
			"CU_RC_TMSR801_SUBJECT": {db.subjects[subject].slug},
		}))
		if err != nil {
			return nil, fmt.Errorf("post search: %v", err)
		}
		rowHeads, err := s.parseCourses(c, icsid, db.careers[career].en, strings.TrimSpace(db.terms[term].en), db.subjects[subject].en)
		if err != nil {
			return nil, fmt.Errorf("parseCourses: %v", err)
		}
		return rowHeads, nil
	}
	type resultS struct {
		c, t, s int
		r       []rowHeadS
	}
	results := make(chan interface{})
	go func() {
		jobs := make(chan struct{}, 10)
		g := func(c, t, s int) {
			retry := 0
			for {
				r, err := f(c, t, s)
				if err == nil {
					results <- resultS{c, t, s, r}
					<-jobs
					return
				}
				retry++
				if retry == 3 {
					results <- fmt.Errorf("crawlCourses(%v, %v, %v): %v", db.careers[c].slug, strings.TrimSpace(db.terms[t].en), db.subjects[s].slug, err)
					<-jobs
					return
				}
				log.Printf("retry crawlCourses(%v, %v, %v): %v", db.careers[c].slug, strings.TrimSpace(db.terms[t].en), db.subjects[s].slug, err)
			}
		}
		for _, ct := range db.careerTerms {
			for _, s := range ct.subjects {
				jobs <- struct{}{}
				go g(ct.career, ct.term, s)
			}
		}
	}()
	var errs []string
	for _, ct := range db.careerTerms {
		for range ct.subjects {
			r := <-results
			switch r := r.(type) {
			case error:
				errs = append(errs, r.Error())
			case resultS:
				log.Println(db.careers[r.c].slug, strings.TrimSpace(db.terms[r.t].en), db.subjects[r.s].slug)
			}
		}
	}
	if errs != nil {
		return fmt.Errorf("%v", strings.Join(errs, "\n"))
	}
	return nil
}

func crawlSubjects() error {
	f := func(career, term string) ([][2]string, error) {
		c, icsid, err := reboot()
		if err != nil {
			return nil, fmt.Errorf("reboot err: %v", err)
		}
		s, err := responseToParser(c.PostForm(cusis, url.Values{
			"ICSID":                       {icsid},
			"ICAction":                    {"CLASS_SRCH_WRK2_STRM$50$"},
			"CLASS_SRCH_WRK2_ACAD_CAREER": {career},
			"CLASS_SRCH_WRK2_STRM$50$":    {term},
		}))
		if err != nil {
			return nil, fmt.Errorf("post career term err: %v", err)
		}
		careerOption := toCareerOption(career, db.careers)
		termOption := toTermOption(term, db.terms)
		if err := s.parseCareerTerm(icsid, careerOption, termOption); err != nil {
			return nil, fmt.Errorf("parseCareerTerm\n%v", err)
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
	type resultS struct {
		c, t int
		s    [][2]string
	}
	results := make(chan interface{})
	go func() {
		jobs := make(chan struct{}, 10)
		g := func(c, t int) {
			retry := 0
			for {
				s, err := f(db.careers[c].slug, db.terms[t].slug)
				if err == nil {
					results <- resultS{c, t, s}
					<-jobs
					return
				}
				retry++
				if retry == 3 {
					results <- fmt.Errorf("crawlSubjects(%v, %v): %v", db.careers[c].slug, strings.TrimSpace(db.terms[t].en), err)
					<-jobs
					return
				}
				log.Printf("retry crawlSubjects(%v, %v): %v\n", db.careers[c].slug, strings.TrimSpace(db.terms[t].en), err)
			}
		}
		for c := range db.careers {
			for t := range db.terms {
				jobs <- struct{}{}
				go g(c, t)
			}
		}
	}()
	var errs []string
	var careerTerms []careerTermS
	var subjects []subjectS
	m := make(map[[2]string]int)
	for range db.careers {
		for range db.terms {
			r := <-results
			switch r := r.(type) {
			case error:
				errs = append(errs, r.Error())
			case resultS:
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
		}
	}
	if errs != nil {
		return fmt.Errorf("%v", strings.Join(errs, "\n"))
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
	db.terms = terms
	reboot = func() (*http.Client, string, error) {
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
	for _, ct := range db.careerTerms {
		fmt.Printf("%v %v %v\n", db.careers[ct.career].slug, strings.TrimSpace(db.terms[ct.term].en), len(ct.subjects))
	}
	if err := crawlCourses(); err != nil {
		return fmt.Errorf("crawlCourses err: %v", err)
	}
	return nil
}
