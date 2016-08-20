package main

import (
	"fmt"
	"html"
	"strconv"
	"strings"
)

func (s *parser) parseCourseRow(cols16 bool, row int) (interface{}, error) {
	class := [2]string{
		`class='PSLEVEL1GRIDODDROW'`,
		`class='PSLEVEL1GRIDEVENROW'`,
	}[row%2]
	rowString := strconv.Itoa(row)
	tabindex1 := strconv.Itoa(row*6 + 38)
	tabindex2 := strconv.Itoa(row*6 + 39)
	tabindex3 := strconv.Itoa(row*6 + 40)
	emptyRowHead := `<tr valign='center'>
<td align='left'  ` + class + `  height='15'>
&nbsp;
</td>
<td align='left'  ` + class + ` >
&nbsp;
</td>
<td align='left'  ` + class + ` >
&nbsp;
</td>
<td align='right'  ` + class + ` >
&nbsp;
</td>
<td align='left'  ` + class + ` >
&nbsp;
</td>
`
	head := `<tr valign='center'>
<td align='left'  ` + class + `  height='15'>
<span  class='PSEDITBOX_DISPONLY' >`
	code := `</span>
</td>
<td align='left'  ` + class + ` >
<span  class='PSHYPERLINK' >
<a name='CLASS_NBR$` + rowString + `' id='CLASS_NBR$` + rowString + `' tabindex='` + tabindex1 + `' href="javascript:submitAction_win0(document.win0,'CLASS_NBR$` + rowString + `');"  class='PSHYPERLINK'  title="Class Nbr" >`
	nbr := `</a></span>
</td>
<td align='left'  ` + class + ` >
<span  class='PSHYPERLINK' >
<a name='CRSE_TITLE$` + rowString + `' id='CRSE_TITLE$` + rowString + `' tabindex='` + tabindex2 + `' href="javascript:submitAction_win0(document.win0,'CRSE_TITLE$` + rowString + `');"  class='PSHYPERLINK'  title="Fixed length 78" `
	beforeTitle1 := `>`
	beforeTitle2 := ` accesskey='`
	beforeTitle3 := `'>`
	title := `</a></span>
</td>
<td align='right'  ` + class + ` >
<span  class='PSEDITBOX_DISPONLY' >`
	units := `</span>
</td>
<td align='left'  ` + class + ` >
<span  class='PSLONGEDITBOX' >`
	teachers := `</span>
</td>
`
	emptyRowBody := `<td align='left'  ` + class + ` >
&nbsp;
</td>
<td align='right'  ` + class + ` >
&nbsp;
</td>
<td align='left'  ` + class + ` >
&nbsp;
</td>
<td align='left'  ` + class + ` >
&nbsp;
</td>
<td align='left'  ` + class + ` >
&nbsp;
</td>
`
	if !cols16 {
		emptyRowBody = `<td align='left'  ` + class + ` >
&nbsp;
</td>
<td align='left'  ` + class + ` >
&nbsp;
</td>
<td align='left'  ` + class + ` >
&nbsp;
</td>
<td align='left'  ` + class + ` >
&nbsp;
</td>
`
	}
	beforeQuota1 := `<td align='left'  ` + class + ` >
<span  class='PSHYPERLINKDISABLED' >
` //' style="color:dimgray;" >
	quota1 := `</span>
</td>
`
	beforeQuota2 := `<td align='left'  ` + class + ` >
<span  class='PSHYPERLINK' >
<a name='ENRL_CAP$` + rowString + `' id='ENRL_CAP$` + rowString + `' tabindex='` + tabindex3 + `' href="javascript:submitAction_win0(document.win0,'ENRL_CAP$` + rowString + `');"  class='PSHYPERLINK'  title="Enrollment Capacity" >`
	quota2 := `</a></span>
</td>
`
	beforeVacancy := `<td align='right'  ` + class + ` >
<span  class='PSEDITBOX_DISPONLY' >`
	vacancy := `</span>
</td>
`
	beforeComponent := `<td align='left'  ` + class + ` >
<span  class='PSEDITBOX_DISPONLY' >`
	component := `</span>
</td>
<td align='left'  ` + class + ` >
<span  class='PSEDITBOX_DISPONLY' >`
	section := `</span>
</td>
<td align='left'  ` + class + ` >
<span  class='PSEDITBOX_DISPONLY' >`
	language := `</span>
</td>
`
	beforePeriod := `<td align='left'  ` + class + ` >
<span  class='PSEDITBOX_DISPONLY' >`
	period := `</span>
</td>
<td align='left'  ` + class + ` >
<span  class='PSEDITBOX_DISPONLY' >`
	room := `</span>
</td>
<td align='left'  ` + class + ` >
<span  class='PSEDITBOX_DISPONLY' >`
	date := `</span>
</td>
<td align='left'  ` + class + ` >
<span  class='PSEDITBOX_DISPONLY' >`
	add := `</span>
</td>
<td align='left'  ` + class + ` >
<span  class='PSEDITBOX_DISPONLY' >`
	drop := `</span>
</td>
`
	emptyDept := `<td align='left'  ` + class + ` >
&nbsp;
</td>
</tr>
`
	beforeDept := `<td align='left'  ` + class + ` >
<span  class='PSEDITBOX_DISPONLY' >`
	dept := `</span>
</td>
</tr>
`
	var rowHead rowHeadS
	var rowBody rowBodyS
	var rowFoot rowFootS
	var err error
	parseRowFoot := func() error {
		if err = s.spanErr(beforePeriod); err != nil {
			return err
		}
		if cols16 {
			if !s.span("TBA") {
				if t, err := s.takeErr(20); err == nil {
					rowFoot.period = &t
				} else {
					return err
				}
			}
			if err = s.spanErr(period); err != nil {
				return err
			}
			if !s.span("TBA") {
				if t, err := s.splitErr(room); err == nil {
					rowFoot.room = &t
				} else {
					return err
				}
			} else if err := s.spanErr(room); err != nil {
				return err
			}
			if !s.span("TBA") {
				if t, err := s.takeErr(23); err == nil {
					rowFoot.date = &t
				} else {
					return err
				}
			}
			if err = s.spanErr(date); err != nil {
				return err
			}
		} else {
			if !s.span("TBA") {
				if t, err := s.takeErr(23); err == nil {
					rowFoot.date = &t
				} else {
					return err
				}
			}
			if err = s.spanErr(date); err != nil {
				return err
			}
			if !s.span("TBA") {
				if t, err := s.takeErr(20); err == nil {
					rowFoot.period = &t
				} else {
					return err
				}
			}
			if err = s.spanErr(period); err != nil {
				return err
			}
			if !s.span("TBA") {
				if t, err := s.splitErr(room); err == nil {
					rowFoot.room = &t
				} else {
					return err
				}
			} else if err := s.spanErr(room); err != nil {
				return err
			}
		}
		if s.span("Yes") {
			rowFoot.add = true
		} else if err = s.spanErr("&nbsp;"); err != nil {
			return err
		}
		if err = s.spanErr(add); err != nil {
			return err
		}
		if s.span("Yes") {
			rowFoot.drop = true
		} else if err = s.spanErr("&nbsp;"); err != nil {
			return err
		}
		if err = s.spanErr(drop); err != nil {
			return err
		}
		return nil
	}
	parseRowBody := func() error {
		if cols16 {
			if err = s.spanErr(beforeVacancy); err != nil {
				return err
			}
			if t, err := s.splitErr(vacancy); err == nil {
				if rowBody.vacancy, err = strconv.Atoi(t); err != nil {
					return err
				}
			} else {
				return err
			}
		}
		if err = s.spanErr(beforeComponent); err != nil {
			return err
		}
		if rowBody.component, err = s.splitErr(component); err != nil {
			return err
		}
		if rowBody.section, err = s.splitErr(section); err != nil {
			return err
		}
		if rowBody.language, err = s.splitErr(language); err != nil {
			return err
		}
		if err = parseRowFoot(); err != nil {
			return fmt.Errorf("parseRowFoot\n%v", err)
		}
		rowBody.rowFoot = []rowFootS{rowFoot}
		if err = s.spanErr(beforeDept); err != nil {
			return err
		}
		if rowBody.dept, err = s.splitErr(dept); err != nil {
			return err
		}
		return nil
	}
	parseRowHead := func() error {
		if err = s.spanErr(head); err != nil {
			return err
		}
		if rowHead.code, err = s.splitErr(code); err != nil {
			return err
		}
		if rowHead.nbr, err = s.takeErr(4); err != nil {
			return err
		}
		if err = s.spanErr(nbr); err != nil {
			return err
		}
		if s.span(beforeTitle1) {
			if t, err := s.splitErr(title); err == nil {
				rowHead.title = html.UnescapeString(t)
			} else {
				return err
			}
		} else {
			if err = s.spanErr(beforeTitle2); err != nil {
				return err
			}
			key := s.takeRune()
			if err = s.spanErr(beforeTitle3); err != nil {
				return err
			}
			t, err := s.splitErr(title)
			if err != nil {
				return err
			}
			rowHead.title = html.UnescapeString(strings.Replace(t, `<u>`+key+`</u>`, key, 1))
		}
		if rowHead.units, err = s.splitErr(units); err != nil {
			return err
		}
		if rowHead.teachers, err = s.splitErr(teachers); err != nil {
			return err
		}
		if s.span(beforeQuota1) {
			if t, err := s.splitErr(quota1); err == nil {
				if rowBody.quota, err = strconv.Atoi(t); err != nil {
					return err
				}
			} else {
				return err

			}
		} else if err = s.spanErr(beforeQuota2); err == nil {
			if t, err := s.splitErr(quota2); err == nil {
				if rowBody.quota, err = strconv.Atoi(t); err != nil {
					return err
				}
				rowHead.reserves = []reserveS{}
			} else {
				return err

			}
		} else {
			return err
		}
		if err = parseRowBody(); err != nil {
			return fmt.Errorf("parseRowBody\n%v", err)
		}
		rowHead.rowBody = []rowBodyS{rowBody}
		return nil
	}
	if s.span(emptyRowHead) {
		if s.span(emptyRowBody) {
			if err = parseRowFoot(); err != nil {
				return nil, fmt.Errorf("parseRowFoot\n%v", err)
			}
			if err = s.spanErr(emptyDept); err != nil {
				return nil, fmt.Errorf("emptyRowToe\n%v", err)
			}
			return rowFoot, nil
		}
		if err = s.spanErr(beforeQuota1); err != nil {
			return nil, fmt.Errorf("beforeQuota1\n%v", err)
		}
		if t, err := s.splitErr(quota1); err == nil {
			if rowBody.quota, err = strconv.Atoi(t); err != nil {
				return nil, fmt.Errorf("quota1: %v", err)
			}
		} else {
			return nil, fmt.Errorf("quota1\n%v", err)
		}
		if err = parseRowBody(); err != nil {
			return nil, fmt.Errorf("parseRowBody\n%v", err)
		}
		return rowBody, nil
	}
	if err = parseRowHead(); err != nil {
		return nil, fmt.Errorf("parseRowHead\n%v", err)
	}
	return rowHead, nil
}
