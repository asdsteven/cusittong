package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (s *parser) parseReserves(stateNum, icsid, code, nbr, component, title string) ([]reserveS, error) {
	const beforeStateNum = `<html dir='ltr' lang='en'>
<!-- Copyright (c) 2000, 2007, Oracle. All rights reserved. -->
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<script language='JavaScript'>
var  totalTimeoutMilliseconds = 2073600000; 
var  warningTimeoutMilliseconds = 2073600000; 
var timeOutURL = 'https://cusis.cuhk.edu.hk/psc/public/EMPLOYEE/HRMS/?cmd=expire';
var timeoutWarningPageURL = 'https://cusis.cuhk.edu.hk/psc/public/EMPLOYEE/HRMS/s/WEBLIB_TIMEOUT.PT_TIMEOUTWARNING.FieldFormula.IScript_TIMEOUTWARNING';
</script> 

<link rel='stylesheet' type='text/css' href='/cs/public/cache/PSSTYLEDEF_ALT_1.css' />
<title>Enquire Timetable</title>
<script language='JavaScript'>
var baseKey_win0 = "\x1b";
var altKey_win0 = "05678\xbc\xbe\xbf\xde";
var ctrlKey_win0 = "JK";
var saveWarningKeys_win0 = "";
var bTabOverTB_win0 = false;
var bTabOverPg_win0 = false;
var bTabOverNonPS_win0 = false;
var strCurrUrl='https://cusis.cuhk.edu.hk/psc/public/EMPLOYEE/HRMS/c/CU_SCR_MENU.CU_TMSR801.GBL?&PAGE=CU_TMSR801_RSV_CAP';
</script>
<script language='JavaScript' type='text/javascript' src='/cs/public/cache/PT_SCRIPTGENERIC_1.js'>
</script>
<SCRIPT LANGUAGE='JavaScript'>
try {
document.domain = "cuhk.edu.hk";
}
catch(err) {;}
</SCRIPT>
<script language='JavaScript' type='text/javascript' src='/cs/public/cache/PT_PAGESCRIPT_win0_1.js'>
</script>
<script language='JavaScript' type='text/javascript' src='/cs/public/cache/PT_COPYURL_1.js'>
</script>
<script language='JavaScript' type='text/javascript' src='/cs/public/cache/PT_ISCROSSDOMAIN_1.js'>
</script>
<script language='JavaScript' type='text/javascript' src='/cs/public/cache/PT_SAVEWARNINGSCRIPT_1.js'>
</script>
<script language='JavaScript'>
var nResubmit=0;
setupTimeout();
function submitAction_win0(form, name)
{
form.ICAction.value=name;
form.ICXPos.value=getScrollX();
form.ICYPos.value=getScrollY();
processing_win0(1,3000);
form.ICResubmit.value=nResubmit;
form.submit();
nResubmit++;
}
</script>
<script language='JavaScript' type='text/javascript' src='/cs/public/cache/PT_EDITSCRIPT_win0_1.js'>
</script>
<script language='JavaScript' type='text/javascript' src='/cs/public/cache/PT_EDITSCRIPT2_win0_1.js'>
</script>
</head>
<body class='PSPAGE'  onload="
setEventHandlers_win0('ICFirstAnchor_win0', 'ICLastAnchor_win0', false);
processing_win0(0,3000);
if (!isCrossDomainTop()) top.document.title=document.title;
setKeyEventHandler_win0();
nResubmit=0;
">
<div id="WAIT_win0" style="position:absolute;right:0;display:block;">
<img align='right' src='/cs/public/cache/PT_PROCESSING_1.gif' class='PSPROCESSING' alt='Processing... please wait' title='Processing... please wait' />
</div>
<a name='ICFirstAnchor_win0'></a>
<form name='win0' method='post' action="../../../EMPLOYEE/HRMS/c/CU_SCR_MENU.CU_TMSR801.GBL"  autocomplete='off'>
<input type='hidden' name='ICType' value='Panel' />
<input type='hidden' name='ICElementNum' value='0' />
<input type='hidden' name='ICStateNum' value='`
	const beforeICSID = `' />
<input type='hidden' name='ICAction' value='None' />
<input type='hidden' name='ICXPos' value='0' />
<input type='hidden' name='ICYPos' value='0' />
<input type='hidden' name='ICFocus' value='' />
<input type='hidden' name='ICSaveWarningFilter' value='0' />
<input type='hidden' name='ICChanged' value='-1' />
<input type='hidden' name='ICResubmit' value='0' />
<input type='hidden' name='ICSID' value='`
	const beforeCode = `' />
<div ID='PAGEBAR'><table cols='3' width='100%' cellpadding='0' cellspacing='0' hspace='0' vspace='0'>
<tr>
<td width='80%'></td><td width='10%' nowrap='nowrap' align='right'><label for='#ICDataLang'><span class='PSDROPDOWNLABEL' >Data Language:&nbsp;&nbsp;</span><select name='#ICDataLang' id='#ICDataLang' tabindex='1' class='PSMULTILANG'  PSaccesskey='9' onchange="submitAction_win0(this.form,this.name);">
<option value="ENG" selected='selected'>English</option>
<option value="ZHS">Simplified Chinese</option>
<option value="ZHT">Traditional Chinese</option>
</select></label></td>
<td width='10%' nowrap='nowrap' align='right'></td></tr>
</table>
<br />
</div><table class='PSPAGECONTAINER' ><tr><td>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' class='PSPAGECONTAINER' cols='10' width='814'>
<tr>
<td width='16' height='8'></td>
<td width='8'></td>
<td width='4'></td>
<td width='76'></td>
<td width='9'></td>
<td width='32'></td>
<td width='63'></td>
<td width='300'></td>
<td width='61'></td>
<td width='245'></td>
</tr>
<tr>
<td height='48'></td>
<td colspan='6'  valign='top' align='left'>
<span  class='PAPAGETITLE' >Reserve Capacity</span>
</td>
</tr>
<tr>
<td height='1' colspan='2'></td>
<td colspan='3' rowspan='2'  valign='top' align='left'>
<label for='DERIVED_CLSRCH_DESCR1' class='PSEDITBOXLABEL' >Course:</label>
</td>
</tr>
<tr>
<td height='19' colspan='2'></td>
<td colspan='4' rowspan='2'  valign='top' align='left'>
<span  class='PSEDITBOX_DISPONLY' >`
	const beforeNbr = `</span>
</td>
</tr>
<tr>
<td height='1' colspan='2'></td>
<td colspan='2' rowspan='2'  valign='top' align='left'>
<label for='CLASS_SRCH_WRK2_CLASS_NBR' class='PSEDITBOXLABEL' >Class Nbr:</label>
</td>
</tr>
<tr>
<td height='19' colspan='2'></td>
<td colspan='2'  valign='top' align='right'>
<span  class='PSEDITBOX_DISPONLY' >`
	const beforeComponent = `</span>
</td>
</tr>
<tr>
<td height='1' colspan='2'></td>
<td colspan='3' rowspan='2'  valign='top' align='left'>
<label for='CLASS_SRCH_WRK2_SSR_COMPONENT' class='PSEDITBOXLABEL' >Component:</label>
</td>
</tr>
<tr>
<td height='19' colspan='2'></td>
<td colspan='4' rowspan='2'  valign='top' align='left'>
<span  class='PSEDITBOX_DISPONLY' >`
	const beforeTitle = `</span>
</td>
</tr>
<tr>
<td height='1' colspan='2'></td>
<td colspan='3' rowspan='2'  valign='top' align='left'>
<label for='DERIVED_CLSRCH_COURSE_TITLE_LONG' class='PSEDITBOXLABEL' >Title:</label>
</td>
</tr>
<tr>
<td height='34' colspan='2'></td>
<td colspan='5'  valign='top' align='left'>
<span  class='PSEDITBOX_DISPONLY' >`
	const beforeHeight = `</span>
</td>
</tr>
<tr>
<td height='`
	const beforeRows = `' colspan='3'></td>
<td colspan='5'  valign='top' align='left'>
<table border='1' cellspacing='0' class='PSLEVEL1GRIDWBO'  id='CU_CLS_RSV_VW$scroll$0' dir='ltr' cellpadding='2' cols='3' width='480'>
<tr valign='center'>
<th scope='col' width='249' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Reserved for</th>
<th scope='col' width='85' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Enrolment Quota</th>
<th scope='col' width='82' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Enrolment Total</th>
</tr>
`
	const beforeMajorClass = `<tr valign='center'>
<td align='left'  `
	const beforeMajor = `  height='14'>
<span  class='PSEDITBOX_DISPONLY' >`
	const beforeQuotaClass = `</span>
</td>
<td align='right'  `
	const beforeQuota = ` >
<span  class='PSEDITBOX_DISPONLY' >`
	const beforeEnrollClass = `</span>
</td>
<td align='right'  `
	const beforeEnroll = ` >
<span  class='PSEDITBOX_DISPONLY' >`
	const afterEnroll = `</span>
</td>
</tr>
`
	/*<tr valign='center'>
	  <td align='left'  class='PSLEVEL1GRIDEVENROW'  height='14'>
	  <span  class='PSEDITBOX_DISPONLY' >For UG ECON Major</span>
	  </td>
	  <td align='right'  class='PSLEVEL1GRIDEVENROW' >
	  <span  class='PSEDITBOX_DISPONLY' >5</span>
	  </td>
	  <td align='right'  class='PSLEVEL1GRIDEVENROW' >
	  <span  class='PSEDITBOX_DISPONLY' >1</span>
	  </td>
	  </tr>*/
	const beforeLast = `</table>
</td>
</tr>
<tr>
<td height='14' colspan='10'></td>
</tr>
<tr>
<td height='105' colspan='3'></td>
<td colspan='7'  valign='top' align='left'>
<span  class='PSHYPERLINK' >
<a name='CU_RC_TMSR801_SSR_PB_CLOSE$0' id='CU_RC_TMSR801_SSR_PB_CLOSE$0' tabindex='`
	const last = `' href="javascript:submitAction_win0(document.win0,'CU_RC_TMSR801_SSR_PB_CLOSE$0');"  class='PSHYPERLINK'  title="Close Push Button" >Return</a></span>
</td>
</tr>
</table>
</td></tr>
</table>
<div id='DetachDiv' height=0 width=0 frameborder=0></div>
</form>
<a name='ICLastAnchor_win0'></a>
</body>
</html>
`
	if err := s.spanErr(beforeStateNum); err != nil {
		return nil, fmt.Errorf("beforeStateNum\n%v", err)
	}
	if err := s.spanErr(stateNum); err != nil {
		return nil, fmt.Errorf("stateNum\n%v", err)
	}
	if err := s.spanErr(beforeICSID); err != nil {
		return nil, fmt.Errorf("beforeICSID\n%v", err)
	}
	if err := s.spanErr(icsid); err != nil {
		return nil, fmt.Errorf("icsid\n%v", err)
	}
	if err := s.spanErr(beforeCode); err != nil {
		return nil, fmt.Errorf("beforeCode\n%v", err)
	}
	if err := s.spanErr(code); err != nil {
		return nil, fmt.Errorf("code\n%v", err)
	}
	if err := s.spanErr(beforeNbr); err != nil {
		return nil, fmt.Errorf("beforeNbr\n%v", err)
	}
	if err := s.spanErr(nbr); err != nil {
		return nil, fmt.Errorf("nbr\n%v", err)
	}
	if err := s.spanErr(beforeComponent); err != nil {
		return nil, fmt.Errorf("beforeComponent\n%v", err)
	}
	if err := s.spanErr(component); err != nil {
		return nil, fmt.Errorf("component\n%v", err)
	}
	if err := s.spanErr(beforeTitle); err != nil {
		return nil, fmt.Errorf("beforeTitle\n%v", err)
	}
	if t, err := s.splitErr(beforeHeight); err == nil {
		entities := [...][2]string{
			{`&amp;Staff`, `Staff`},
			{`&amp;Ancient`, `Ancient`},
			{`&nbsp;`, ` `},
			{`&amp;`, `&`},
			{`&#039;`, `'`},
			{`R&amp;W`, `RW`},
			{`E&amp;H`, `EH`},
			{`Cul&amp;Soc`, `CulSoc`},
			{`Health&amp;Healthy`, `HealthHealthy`},
			{`Sexuality&amp;Race`, `SexualityRace`},
			{`Matrices&amp;Lin`, `MatricesLin`},
			{` &amp;T `, ` T `},
			{` T&amp;P `, ` TP `},
			{` Sch&#039;n&amp;EDU `, ` Sch'nEDU `},
			{` Women&amp;Gender `, ` WomenGender `},
		}
		var u []byte
		for i := 0; i < len(t); i++ {
			x := t[i : i+1]
			for _, v := range entities {
				if strings.HasPrefix(t[i:], v[0]) {
					x = v[1]
					i += len(v[0]) - 1
					break
				}
			}
			u = append(u, x...)
		}
		t = string(u)
		if title != t {
			return nil, fmt.Errorf("title\n%v != \n%v", t, title)
		}
	} else {
		return nil, fmt.Errorf("beforeHeight\n%v", err)
	}
	t, err := s.splitErr(beforeRows)
	if err != nil {
		return nil, fmt.Errorf("beforeRows\n%v", err)
	}
	height, err := strconv.Atoi(t)
	if err != nil {
		return nil, fmt.Errorf("height: %v", err)
	}
	rows := (height - 20) / 21
	var reserves []reserveS
	for row := 0; row < rows; row++ {
		class := [2]string{
			`class='PSLEVEL1GRIDODDROW'`,
			`class='PSLEVEL1GRIDEVENROW'`,
		}[row%2]
		if err := s.spanErr(beforeMajorClass + class + beforeMajor); err != nil {
			return nil, fmt.Errorf("beforeMajor\n%v", err)
		}
		major, err := s.splitErr(beforeQuotaClass + class + beforeQuota)
		if err != nil {
			return nil, fmt.Errorf("beforeQuota\n%v", err)
		}
		major = strings.Replace(major, "&amp;", "&", -1)
		t, err := s.splitErr(beforeEnrollClass + class + beforeEnroll)
		if err != nil {
			return nil, fmt.Errorf("beforeEnroll\n%v", err)
		}
		quota, err := strconv.Atoi(t)
		if err != nil {
			return nil, fmt.Errorf("quota: %v", err)
		}
		t, err = s.splitErr(afterEnroll)
		if err != nil {
			return nil, fmt.Errorf("afterEnroll\n%v", err)
		}
		enroll := 0
		if t != `&nbsp;` {
			enroll, err = strconv.Atoi(t)
			if err != nil {
				return nil, fmt.Errorf("enroll: %v", err)
			}
		}
		reserves = append(reserves, reserveS{major, quota, enroll})
	}
	if err := s.spanErr(beforeLast); err != nil {
		return nil, fmt.Errorf("beforeLast\n%v", err)
	}
	if err := s.spanErr(strconv.Itoa(42 + rows*2)); err != nil {
		return nil, fmt.Errorf("lastTabindex\n%v", err)
	}
	if err := s.equalErr(last); err != nil {
		return nil, fmt.Errorf("last\n%v", err)
	}
	return reserves, nil
}
