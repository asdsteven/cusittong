package main

import (
	"fmt"
	"html"
	"strconv"
	"strings"
	"time"
)

func (s *parser) parseCourseRow(row int, subj, group string) (ret interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(parseError); ok {
				err = e.error
			} else {
				panic(r)
			}
		}
	}()
	class := [2]string{
		`class='PSLEVEL1GRIDODDROW'`,
		`class='PSLEVEL1GRIDEVENROW'`,
	}[row%2]
	rowString := strconv.Itoa(row)
	tabindex1 := strconv.Itoa(row*6 + 36)
	tabindex2 := strconv.Itoa(row*6 + 37)
	tabindex3 := strconv.Itoa(row*6 + 38)
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
<span  class='PSEDITBOX_DISPONLY' >` + subj
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
	beforeQuota1 := `<td align='left'  ` + class + ` >
<span  class='PSHYPERLINKDISABLED' >
`
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
	parseRowFoot := func() {
		s.spanPanic(`before period`, beforePeriod)
		var p, r, d string
		p = s.splitPanic(`period`, period)
		r = s.splitPanic(`room`, room)
		d = s.splitPanic(`date`, date)
		if p != "TBA" {
			if len(p) != 20 {
				panic(parseError{fmt.Errorf("p %v", p)})
			}
			if p[3:5] == "00" || p[13:15] == "00" {
				panic(parseError{fmt.Errorf("p 12:00 %v", p)})
			}
			start, err := time.ParseInLocation("03:04PM", p[3:10], hongKong)
			if err != nil {
				panic(parseError{fmt.Errorf("p start %v: %v", p[3:10], err)})
			}
			end, err := time.ParseInLocation("03:04PM", p[13:], hongKong)
			if err != nil {
				panic(parseError{fmt.Errorf("p end %v: %v", p[3:10], err)})
			}
			if start.Unix() >= end.Unix() {
				panic(parseError{fmt.Errorf("p %v >= %v: %v", start, end, p)})
			}
			rowFoot.weekday = p[:2]
			rowFoot.period = []time.Time{start, end}
		}
		rowFoot.room = r
		if d != "TBA" {
			if len(d) != 23 {
				panic(parseError{fmt.Errorf("d %v", d)})
			}
			start, err := time.ParseInLocation("02/01/2006", d[:10], hongKong)
			if err != nil {
				panic(parseError{fmt.Errorf("d start %v: %v", d[:10], err)})
			}
			end, err := time.ParseInLocation("02/01/2006", d[13:], hongKong)
			if err != nil {
				panic(parseError{fmt.Errorf("d end %v: %v", d[13:], err)})
			}
			rowFoot.date = []time.Time{start, end}
		}
		if s.span("Yes") {
			rowFoot.add = true
		} else {
			s.spanPanic(`before add`, "&nbsp;")
		}
		s.spanPanic(`add`, add)
		if s.span("Yes") {
			rowFoot.drop = true
		} else {
			s.spanPanic(`before drop`, "&nbsp;")
		}
		s.spanPanic(`drop`, drop)
	}
	parseRowBody := func(group string) {
		s.spanPanic(`before vacancy`, beforeVacancy)
		rowBody.vacancy = s.splitIntPanic(`vacancy`, vacancy)
		s.spanPanic(`before component`, beforeComponent)
		rowBody.component = s.splitPanic(`component`, component)
		s.spanPanic(`group`, group)
		rowBody.section = s.splitPanic(`section`, section)
		rowBody.language = s.splitPanic(`language`, language)
		parseRowFoot()
		rowBody.rowFoots = []rowFootS{rowFoot}
		s.spanPanic(`before dept`, beforeDept)
		rowBody.dept = s.splitPanic(`dept`, dept)
	}
	parseRowHead := func() {
		s.spanPanic(`before code`, head)
		rowHead.code = s.takePanic(`code`, 4)
		rowHead.group = s.splitPanic(`group`, code)
		rowHead.nbr = s.splitPanic(`nbr`, nbr)
		if s.span(beforeTitle1) {
			t := s.splitPanic(`title`, title)
			rowHead.title = html.UnescapeString(t)
		} else {
			s.spanPanic(`before title2`, beforeTitle2)
			key := s.takeRune()
			s.spanPanic(`before title3`, beforeTitle3)
			t := s.splitPanic(`title`, title)
			rowHead.title = html.UnescapeString(strings.Replace(t, `<u>`+key+`</u>`, key, 1))
		}
		rowHead.units = s.splitPanic(`units`, units)
		for _, t := range strings.Split(s.splitPanic(`teachers`, teachers), `<br />`) {
			t = strings.TrimSpace(t)
			if t == `-` {
				continue
			}
			if !strings.HasPrefix(t, `- `) {
				panic(parseError{fmt.Errorf(
					"mismatch teacher\n%v",
					rowHead.teachers,
				)})
			}
			t = t[2:]
			rowHead.teachers = append(rowHead.teachers, t)
		}
		if s.span(beforeQuota1) {
			rowBody.quota = s.splitIntPanic(`quota1`, quota1)
		} else {
			s.spanPanic(`before quota2`, beforeQuota2)
			rowBody.quota = s.splitIntPanic(`quota2`, quota2)
			rowHead.reserves = []reserveS{}
		}
		parseRowBody(rowHead.group)
		rowHead.rowBodys = []rowBodyS{rowBody}
		rowHead.row = row
	}
	if s.span(emptyRowHead) {
		if s.span(emptyRowBody) {
			parseRowFoot()
			s.spanPanic(`empty dept`, emptyDept)
			return rowFoot, nil
		}
		s.spanPanic(`before quota1`, beforeQuota1)
		rowBody.quota = s.splitIntPanic(`quota1`, quota1)
		parseRowBody(group)
		return rowBody, nil
	}
	parseRowHead()
	return rowHead, nil
}
