package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
)

func browse() error {
	if err := crawlTerms(); err != nil {
		return fmt.Errorf("crawlTerms err: %v", err)
	}
	if err := crawlSubjects(); err != nil {
		return fmt.Errorf("crawlSubjects err: %v", err)
	}
	c, icsid, err := reboot(3, 1)
	if err != nil {
		return fmt.Errorf("reboot: %v", err)
	}
	s, err := responseToParser(c.PostForm(cusis, url.Values{
		"ICSID":                 {icsid},
		"ICAction":              {"CU_RC_TMSR801_SSR_PB_CLASS_SRCH"},
		"CU_RC_TMSR801_SUBJECT": {"ACCT"},
	}))
	if err != nil {
		return fmt.Errorf("post search: %v", err)
	}
	s, err = responseToParser(c.PostForm(cusis, url.Values{
		"ICSID":    {icsid},
		"ICAction": {"CRSE_TITLE$0"},
	}))
	if err != nil {
		return fmt.Errorf("post: %v", err)
	}
	s, err = responseToParser(c.PostForm(cusis, url.Values{
		"ICSID":    {icsid},
		"ICAction": {"DERIVED_SAA_CRS_SSR_PB_GO"},
	}))
	if err != nil {
		return fmt.Errorf("post: %v", err)
	}
	s, err = responseToParser(c.PostForm(cusis, url.Values{
		"ICSID":    {icsid},
		"ICAction": {"CLASS_SECTION$0"},
	}))
	if err != nil {
		return fmt.Errorf("post: %v", err)
	}
	ioutil.WriteFile("parseClass.go", s.s, 0666)
	return nil
}
