package main

import (
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)

type parser struct {
	s []byte
}

func (s *parser) takeErr(i int) (string, error) {
	if len(s.s) < i {
		return "", errors.New(string(s.s))
	}
	ret := string(s.s[:i])
	s.s = s.s[i:]
	return ret, nil
}

func (s *parser) takeRune() string {
	_, i := utf8.DecodeRune(s.s)
	ret := string(s.s[:i])
	s.s = s.s[i:]
	return ret
}

func (s *parser) span(t string) bool {
	if len(s.s) < len(t) {
		return false
	}
	for i, v := range s.s[:len(t)] {
		if t[i] != v {
			return false
		}
	}
	s.s = s.s[len(t):]
	return true
}

func (s *parser) spanErr(t string) error {
	if len(s.s) < len(t) {
		return errors.New(s.diff(t))
	}
	for i, v := range s.s[:len(t)] {
		if t[i] != v {
			return errors.New(s.diff(t))
		}
	}
	s.s = s.s[len(t):]
	return nil
}

func (s *parser) spanLastErr(t string) error {
	if len(s.s) < len(t) {
		return errors.New(s.diff(t))
	}
	for i, v := range s.s[len(s.s)-len(t):] {
		if t[i] != v {
			return errors.New((&parser{s.s[len(s.s)-len(t):]}).diff(t))
		}
	}
	s.s = s.s[:len(s.s)-len(t)]
	return nil
}

func (s *parser) splitErr(t string) (string, error) {
	for i := range s.s[:len(s.s)-len(t)+1] {
		found := true
		for j, v := range s.s[i : i+len(t)] {
			if t[j] != v {
				found = false
				break
			}
		}
		if found {
			ret := string(s.s[:i])
			s.s = s.s[i+len(t):]
			return ret, nil
		}
	}
	return "", errors.New(s.diff(t))
}

func (s *parser) equal(t string) bool {
	if len(s.s) != len(t) {
		return false
	}
	for i, v := range s.s {
		if t[i] != v {
			return false
		}
	}
	return true
}

func (s *parser) equalErr(t string) error {
	if len(s.s) != len(t) {
		return errors.New(s.diff(t))
	}
	for i, v := range s.s {
		if t[i] != v {
			return errors.New(s.diff(t))
		}
	}
	return nil
}

func (s *parser) parseRows() (int, error) {
	t, err := s.splitErr(` of `)
	if err != nil {
		return 0, err
	}
	if err := s.spanErr(t); err != nil {
		return 0, err
	}
	rows, err := strconv.Atoi(t)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s *parser) diff(t string) string {
	f := func(s string) string {
		t := strings.SplitN(s, "\n", 6)
		if len(t) == 6 {
			t = t[:5]
		}
		return strings.Join(t, "\n")
	}
	err := ioutil.WriteFile("a.html", s.s, 0666)
	if err != nil {
		log.Printf("err writeFile: %v\n", err)
	}
	for i, tv := range t {
		sv, _ := utf8.DecodeRune(s.s[i:])
		if tv != sv {
			/*			x := `<<<---byte diff----
						` + fmt.Sprintf("%v", []byte(f(string(s.s[i:])))) + `
						===================
						` + fmt.Sprintf("%v", []byte(f(t[i:]))) + `
						------------>>>>>>>
						`*/
			return `<<<---diff---------
` + f(string(s.s[i:])) + `
===================
` + f(t[i:]) + `
------------>>>>>>>

`
		}
	}
	return ""
}
