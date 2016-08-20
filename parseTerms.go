package main

import (
	"fmt"
)

func (s *parser) parseTerms() ([][2]string, error) {
	const (
		option0 = `<option value="`
		option1 = `">`
		option2 = `" selected='selected'>`
		option3 = "</option>\n"
	)
	var ret [][2]string
	for len(s.s) > 0 {
		if err := s.spanErr(option0); err != nil {
			return nil, fmt.Errorf("option0\n%v", err)
		}
		slug, err := s.takeErr(4)
		if err != nil {
			return nil, fmt.Errorf("slug\n%v", err)
		}
		if s.span(option1) {
		} else if err := s.spanErr(option2); err == nil {
		} else {
			return nil, fmt.Errorf("option2\n%v", err)
		}
		full, err := s.splitErr(option3)
		if err != nil {
			return nil, fmt.Errorf("full\n%v", err)
		}
		ret = append(ret, [2]string{slug, full})
	}
	return ret, nil
}
