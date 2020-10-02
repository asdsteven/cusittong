package main

import (
	"fmt"
)

func (s *parser) parseTerms() (*[2]string, error) {
	const (
		option0 = `<option value="`
		option1 = `">`
		option2 = `" selected='selected'>`
		option3 = "</option>\n"
	)
	var ret *[2]string
	for len(s.s) > 0 {
		selected := false
		if err := s.spanErr(option0); err != nil {
			return nil, fmt.Errorf("option0\n%v", err)
		}
		slug, err := s.takeErr(4)
		if err != nil {
			return nil, fmt.Errorf("slug\n%v", err)
		}
		if s.span(option1) {
		} else if err := s.spanErr(option2); err == nil {
			selected = true
		} else {
			return nil, fmt.Errorf("option2\n%v", err)
		}
		full, err := s.splitErr(option3)
		if err != nil {
			return nil, fmt.Errorf("full\n%v", err)
		}
		if selected {
			if ret != nil {
				return nil, fmt.Errorf("multiple selected")
			}
			ret = &[2]string{slug, full}
		}
	}
	if ret == nil {
		return nil, fmt.Errorf("no selected")
	}
	return ret, nil
}
