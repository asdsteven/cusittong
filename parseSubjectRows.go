package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (s *parser) parseSubjectRows(rows, tabindex int) ([][2]string, error) {
	/*<tr align='left'>
	  <td scope='row' class='PSSRCHRESULTSEVENROW' nowrap='nowrap'><a href="javascript: submitAction_win0(document.win0,'#ICRow3');" tabindex='35' class='PSSRCHRESULTSEVENROW' >ARCH</a></td>
	  <td class='PSSRCHRESULTSEVENROW' nowrap='nowrap'>Architectural Studies</td>
	  </tr>
	  <tr align='left'>
	  <td scope='row' class='PSSRCHRESULTSEVENROW' nowrap='nowrap'><a name='SEARCH_RESULTLAST' href="javascript: submitAction_win0(document.win0,'#ICRow99');" tabindex='131' class='PSSRCHRESULTSEVENROW' >SPAN</a></td>
	  <td class='PSSRCHRESULTSEVENROW' nowrap='nowrap'>Spanish</td>
	  </tr>*/
	const (
		beforeNameClass = `<tr align='left'>
<td scope='row' class='PSSRCHRESULTS`
		beforeName      = `ROW' nowrap='nowrap'><a `
		nameFirst       = `name='SEARCH_RESULT1' `
		nameLast        = `name='SEARCH_RESULTLAST' `
		beforeRow       = `href="javascript: submitAction_win0(document.win0,'#ICRow`
		beforeTabindex  = `');" tabindex='`
		beforeSlugClass = `' class='PSSRCHRESULTS`
		beforeSlug      = `ROW' >`
		beforeFullClass = `</a></td>
<td class='PSSRCHRESULTS`
		beforeFull = `ROW' nowrap='nowrap'>`
		afterFull  = `</td>
</tr>
`
	)
	parseRow := func(row, tabindex int, name string) ([2]string, error) {
		class := [2]string{"ODD", "EVEN"}[row%2]
		a := strings.Join([]string{beforeNameClass, class, beforeName, name, beforeRow, strconv.Itoa(row), beforeTabindex, strconv.Itoa(tabindex), beforeSlugClass, class, beforeSlug}, "")
		b := strings.Join([]string{beforeFullClass, class, beforeFull}, "")
		if err := s.spanErr(a); err != nil {
			return [2]string{}, fmt.Errorf("beforeSlug\n%v", err)
		}
		slug, err := s.takeErr(4)
		if err != nil {
			return [2]string{}, fmt.Errorf("slug\n%v", err)
		}
		if err := s.spanErr(b); err != nil {
			return [2]string{}, fmt.Errorf("afterSlug\n%v", err)
		}
		full, err := s.splitErr(afterFull)
		if err != nil {
			return [2]string{}, fmt.Errorf("full\n%v", err)
		}
		return [2]string{slug, full}, nil
	}
	subject, err := parseRow(0, tabindex, nameFirst)
	if err != nil {
		return nil, fmt.Errorf("first subject: %v", err)
	}
	ret := [][2]string{subject}
	for row := 1; row < rows-1; row++ {
		subject, err = parseRow(row, tabindex+row, "")
		if err != nil {
			return nil, fmt.Errorf("%vth subject: %v", row, err)
		}
		ret = append(ret, subject)
	}
	subject, err = parseRow(rows-1, tabindex+rows-1, nameLast)
	if err != nil {
		return nil, fmt.Errorf("last subject: %v", err)
	}
	ret = append(ret, subject)
	if err := s.equalErr(""); err != nil {
		return nil, fmt.Errorf("last\n%v", err)
	}
	return ret, nil
}
