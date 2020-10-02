package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)

type parser struct {
	s []byte
}

type parseError struct {
	error
}

func (s *parser) takeErr(i int) (string, error) {
	if len(s.s) < i {
		return "", errors.New(string(s.s))
	}
	ret := string(s.s[:i])
	s.s = s.s[i:]
	return ret, nil
}

func (s *parser) takePanic(e string, i int) string {
	if len(s.s) < i {
		panic(parseError{fmt.Errorf("%v\n%v", e, string(s.s))})
	}
	ret := string(s.s[:i])
	s.s = s.s[i:]
	return ret
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

func (s *parser) spanPanic(e, t string) {
	if len(s.s) < len(t) {
		panic(parseError{fmt.Errorf("%v\n%v", e, s.diff(t))})
	}
	for i, v := range s.s[:len(t)] {
		if t[i] != v {
			panic(parseError{fmt.Errorf("%v\n%v", e, s.diff(t))})
		}
	}
	s.s = s.s[len(t):]
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

func (s *parser) spanLastPanic(e, t string) {
	if len(s.s) < len(t) {
		panic(parseError{fmt.Errorf("%v\n%v", e, s.diff(t))})
	}
	for i, v := range s.s[len(s.s)-len(t):] {
		if t[i] != v {
			panic(parseError{fmt.Errorf("%v\n%v", e, (&parser{s.s[len(s.s)-len(t):]}).diff(t))})
		}
	}
	s.s = s.s[:len(s.s)-len(t)]
}

func (s *parser) index(t string) int {
	if len(t) > len(s.s) {
		return -1
	}
	for i := range s.s[:len(s.s)-len(t)+1] {
		found := true
		for j, v := range s.s[i : i+len(t)] {
			if t[j] != v {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}

func (s *parser) splitErr(t string) (string, error) {
	i := s.index(t)
	if i != -1 {
		ret := string(s.s[:i])
		s.s = s.s[i+len(t):]
		return ret, nil
	}
	l, r := 0, len(t)
	for l+1 < r {
		m := (l + r) / 2
		i := s.index(t[:m])
		if i == -1 {
			r = m
		} else {
			l = m
		}
	}
	i = s.index(t[:l])
	return "", errors.New((&parser{s.s[i:]}).diff(t))
}

func (s *parser) splitPanic(e, t string) string {
	i := s.index(t)
	if i != -1 {
		ret := string(s.s[:i])
		s.s = s.s[i+len(t):]
		return ret
	}
	l, r := 0, len(t)
	for l+1 < r {
		m := (l + r) / 2
		i := s.index(t[:m])
		if i == -1 {
			r = m
		} else {
			l = m
		}
	}
	i = s.index(t[:l])
	panic(parseError{fmt.Errorf("%v\n%v", e, (&parser{s.s[i:]}).diff(t))})
}

func (s *parser) splitIntPanic(e, t string) int {
	i, err := strconv.Atoi(s.splitPanic(e, t))
	if err != nil {
		panic(parseError{fmt.Errorf("%v: %v", e, err)})
	}
	return i
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

func (s *parser) equalPanic(e, t string) {
	if len(s.s) != len(t) {
		panic(parseError{fmt.Errorf("%v\n%v", e, s.diff(t))})
	}
	for i, v := range s.s {
		if t[i] != v {
			panic(parseError{fmt.Errorf("%v\n%v", e, s.diff(t))})
		}
	}
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
	err := ioutil.WriteFile("s.html", s.s, 0666)
	if err != nil {
		log.Printf("err writeFile s.html: %v\n", err)
	}
	err = ioutil.WriteFile("t.html", []byte(t), 0666)
	if err != nil {
		log.Printf("err writeFile t.html: %v\n", err)
	}
	ss := append(strings.Split(string(s.s), "\n"), "", "", "", "", "")
	tt := append(strings.Split(t, "\n"), "", "", "", "", "")
	for i, tv := range tt {
		if tv != ss[i] {
			return `<<<---diff---------
` + strings.Join(ss[i:(i+6)], "\n") + `
===================
` + strings.Join(tt[i:(i+6)], "\n") + `
------------>>>>>>>

`
		}
	}
	return ""
}
