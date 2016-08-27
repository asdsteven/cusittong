package main

import (
	"fmt"
	"strconv"
)

func (s *parser) parseCourses(icsid, career, term, subject, subj string) ([]rowHeadS, error) {
	cols, width1, width2, width3 := "16", "1453", "803", "1315"
	if career == "Postgraduate - Research" || career == "Postgraduate - Taught" {
		cols, width1, width2, width3 = "15", "1408", "758", "1270"
	}
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
var strCurrUrl='https://cusis.cuhk.edu.hk/psc/public/EMPLOYEE/HRMS/c/CU_SCR_MENU.CU_TMSR801.GBL?&PAGE=CU_TMSR801_RST';
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
	const afterICSID = `' />
<input type='hidden' name='ICFind' value='' />
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
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' class='PSPAGECONTAINER' cols='9' width='`
	const afterICSIDWidth = `'>
<tr>
<td width='4' height='4'></td>
<td width='8'></td>
<td width='8'></td>
<td width='4'></td>
<td width='4'></td>
<td width='360'></td>
<td width='148'></td>
<td width='`
	const beforeCareer = `'></td>
<td width='114'></td>
</tr>
<tr>
<td height='27' colspan='2'></td>
<td colspan='7'  valign='top' align='left'>
<label for='DERIVED_AA2_DERIVED_TITLE' class='PAPAGETITLE' >Teaching Timetable</label>
</td>
</tr>
<tr>
<td height='24'></td>
<td colspan='5'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='383'>
<tr><td width='381'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='3' width='381' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='7' height='0'></td>
<td width='133'></td>
<td width='241'></td>
</tr>
<tr>
<td height='1'></td>
<td rowspan='2'  valign='top' align='left'>
<label for='ACAD_CAR_TBL_DESCR' class='PSEDITBOXLABEL' >Course Career:</label>
</td>
</tr>
<tr>
<td height='13'></td>
<td  valign='top' align='left'>
<span  class='PSEDITBOX_DISPONLY' >`
	const afterCareer = `</span>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='4' colspan='9'></td>
</tr>
<tr>
<td height='21'></td>
<td colspan='5'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='383'>
<tr><td width='381'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='3' width='381' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='7' height='0'></td>
<td width='133'></td>
<td width='241'></td>
</tr>
<tr>
<td height='1'></td>
<td rowspan='2'  valign='top' align='left'>
<label for='TERM_TBL_DESCR' class='PSEDITBOXLABEL' >Term:</label>
</td>
</tr>
<tr>
<td height='10'></td>
<td  valign='top' align='left'>
<span  class='PSEDITBOX_DISPONLY' >`
	const afterTerm = `</span>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='8' colspan='9'></td>
</tr>
<tr>
<td height='21'></td>
<td colspan='5'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='383'>
<tr><td width='381'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='3' width='381' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='7' height='0'></td>
<td width='133'></td>
<td width='241'></td>
</tr>
<tr>
<td height='1'></td>
<td rowspan='2'  valign='top' align='left'>
<label for='SUBJECT_DESCR' class='PSEDITBOXLABEL' >Subject:</label>
</td>
</tr>
<tr>
<td height='10'></td>
<td  valign='top' align='left'>
<span  class='PSEDITBOX_DISPONLY' >`
	const afterSubject = `</span>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='8' colspan='9'></td>
</tr>
<tr>
<td height='26' colspan='5'></td>
<td colspan='4'  valign='top' align='left'>
<span  class='PSHYPERLINK' >
<a name='CU_RC_TMSR801_SSR_PB_NEW_SEARCH' id='CU_RC_TMSR801_SSR_PB_NEW_SEARCH' tabindex='23' href="javascript:submitAction_win0(document.win0,'CU_RC_TMSR801_SSR_PB_NEW_SEARCH');"  class='PSHYPERLINK'  title="Search" >New Search</a></span>
</td>
</tr>
<tr>
<td height='`
	const afterHeight = `' colspan='4'></td>
<td colspan='4'  valign='top' align='left'>
<table border='1' cellspacing='0' class='PSLEVEL1GRIDWBO'  id='CLASS_LIST$scroll$0' dir='ltr' cellpadding='2' cols='`
	const afterCols = `' width='`
	const afterWidth = `'>
<tr><td class='PSLEVEL1GRIDLABEL'  colspan='`
	const beforeOf = `' align='right'><a name='CLASS_LIST$hfind$0' id='CLASS_LIST$hfind$0' tabindex='26' onclick="return FindString_win0(document.win0.ICFind);" href="javascript:submitAction_win0(document.win0,'CLASS_LIST$hfind$0');"  class='PSLEVEL1GRIDLABEL' >Find</a>&nbsp;|&nbsp;<a name='CLASS_LIST$hexcel$0' id='CLASS_LIST$hexcel$0' tabindex='27' href="javascript:submitAction_win0(document.win0,'CLASS_LIST$hexcel$0');"><img src=/cs/public/cache/PT_DOWNLOAD_1.gif name='CLASS_LIST$hexcel$img$0' alt='Download' title='Download' border='0' /></a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span  class='PSHEADERHYPERLINKD' >First</span>&nbsp;<img src=/cs/public/cache/PT_PREVIOUSROW_D_1.gif name='CLASS_LIST$hup$img$0' alt='Show previous row (inactive button) (Alt+,)' title='Show previous row (inactive button) (Alt+,)' border='0' />&nbsp;<span class='PSGRIDCOUNTER' >`
	const beforeRows16 = `</span>&nbsp;<img src=/cs/public/cache/PT_NEXTROW_D_1.gif name='CLASS_LIST$hdown$img$0' alt='Show next row (inactive button) (Alt+.)' title='Show next row (inactive button) (Alt+.)' border='0' />&nbsp;<span  class='PSHEADERHYPERLINKD' >Last</span>&nbsp;</td></tr>
<tr valign='center'>
<th scope='col' width='82' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Class Code</th>
<th scope='col' width='43' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Class Nbr</th>
<th scope='col' width='165' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Course Title</th>
<th scope='col' width='38' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Units</th>
<th scope='col' width='157' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Teaching Staff</th>
<th scope='col' width='47' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Quota(s)</th>
<th scope='col' width='46' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Vacancy</th>
<th scope='col' width='61' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Course Component</th>
<th scope='col' width='40' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Section Code</th>
<th scope='col' width='50' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Language</th>
<th scope='col' width='105' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Period</th>
<th scope='col' width='61' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Room</th>
<th scope='col' width='75' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Meeting Date</th>
<th scope='col' width='42' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Add Consent</th>
<th scope='col' width='44' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Drop Consent</th>
<th scope='col' width='148' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Course Offering Dept</th>
</tr>
`
	const beforeRows15 = `</span>&nbsp;<img src=/cs/public/cache/PT_NEXTROW_D_1.gif name='CLASS_LIST$hdown$img$0' alt='Show next row (inactive button) (Alt+.)' title='Show next row (inactive button) (Alt+.)' border='0' />&nbsp;<span  class='PSHEADERHYPERLINKD' >Last</span>&nbsp;</td></tr>
<tr valign='center'>
<th scope='col' width='82' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Class Code</th>
<th scope='col' width='43' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Class Nbr</th>
<th scope='col' width='165' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Course Title</th>
<th scope='col' width='38' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Units</th>
<th scope='col' width='157' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Teaching Staff</th>
<th scope='col' width='47' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Quota(s)</th>
<th scope='col' width='61' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Course Component</th>
<th scope='col' width='40' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Section Code</th>
<th scope='col' width='50' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Language</th>
<th scope='col' width='83' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Meeting Date</th>
<th scope='col' width='105' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Period</th>
<th scope='col' width='61' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Room</th>
<th scope='col' width='42' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Add Consent</th>
<th scope='col' width='44' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Drop Consent</th>
<th scope='col' width='148' align='left' class='PSLEVEL1GRIDCOLUMNHDR' >Course Offering Dept</th>
</tr>
`
	const afterRows = `</table>
</td>
</tr>
<tr>
<td height='8' colspan='9'></td>
</tr>
<tr>
<td height='28' colspan='3'></td>
<td colspan='4'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='515'>
<tr><td width='513'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='2' width='513' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='7' height='9'></td>
<td width='506'></td>
</tr>
<tr>
<td height='17'></td>
<td  valign='top' align='left'>
<span  class='PSHYPERLINK' >
<a name='CU_RC_TMSR801_SSR_PB_NEW_SEARCH$41$$0' id='CU_RC_TMSR801_SSR_PB_NEW_SEARCH$41$$0' tabindex='`
	const last = `' href="javascript:submitAction_win0(document.win0,'CU_RC_TMSR801_SSR_PB_NEW_SEARCH$41$$0');"  class='PSHYPERLINK'  title="Search" >New Search</a></span>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='15' colspan='9'></td>
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
	if err := s.spanErr("3"); err != nil {
		return nil, fmt.Errorf("stateNum\n%v", err)
	}
	if err := s.spanErr(beforeICSID); err != nil {
		return nil, fmt.Errorf("beforeICSID\n%v", err)
	}
	if err := s.spanErr(icsid); err != nil {
		return nil, fmt.Errorf("icsid\n%v", err)
	}
	if err := s.spanErr(afterICSID); err != nil {
		return nil, fmt.Errorf("afterICSID\n%v", err)
	}
	if err := s.spanErr(width1); err != nil {
		return nil, fmt.Errorf("width1\n%v", err)
	}
	if err := s.spanErr(afterICSIDWidth); err != nil {
		return nil, fmt.Errorf("afterICSIDWidth\n%v", err)
	}
	if err := s.spanErr(width2); err != nil {
		return nil, fmt.Errorf("width2\n%v", err)
	}
	if err := s.spanErr(beforeCareer); err != nil {
		return nil, fmt.Errorf("beforeCareer\n%v", err)
	}
	if err := s.spanErr(career); err != nil {
		return nil, fmt.Errorf("career\n%v", err)
	}
	if err := s.spanErr(afterCareer); err != nil {
		return nil, fmt.Errorf("afterCareer\n%v", err)
	}
	if err := s.spanErr(term); err != nil {
		return nil, fmt.Errorf("term\n%v", err)
	}
	if err := s.spanErr(afterTerm); err != nil {
		return nil, fmt.Errorf("afterTerm\n%v", err)
	}
	if err := s.spanErr(subject); err != nil {
		return nil, fmt.Errorf("subject\n%v", err)
	}
	if err := s.spanErr(afterSubject); err != nil {
		return nil, fmt.Errorf("afterSubject\n%v", err)
	}
	t, err := s.splitErr(afterHeight)
	if err != nil {
		return nil, fmt.Errorf("height\n%v", err)
	}
	height, err := strconv.Atoi(t)
	if err != nil {
		return nil, fmt.Errorf("height: %v", err)
	}
	if err := s.spanErr(cols); err != nil {
		return nil, fmt.Errorf("cols\n%v", err)
	}
	if err := s.spanErr(afterCols); err != nil {
		return nil, fmt.Errorf("afterCols\n%v", err)
	}
	if err := s.spanErr(width3); err != nil {
		return nil, fmt.Errorf("width3\n%v", err)
	}
	if err := s.spanErr(afterWidth); err != nil {
		return nil, fmt.Errorf("afterwidth\n%v", err)
	}
	if err := s.spanErr(cols); err != nil {
		return nil, fmt.Errorf("cols2\n%v", err)
	}
	if err := s.spanErr(beforeOf); err != nil {
		return nil, fmt.Errorf("beforeOf\n%v", err)
	}
	rows := (height - 35) / 22
	if rows == 1 {
		if err := s.spanErr("1 of 1"); err != nil {
			return nil, fmt.Errorf("rows %v\n%v", rows, err)
		}
	} else {
		if err := s.spanErr(fmt.Sprintf("1-%v of %v", rows, rows)); err != nil {
			return nil, fmt.Errorf("rows %v\n%v", rows, err)
		}
	}
	cols16 := cols == "16"
	if cols16 {
		if err := s.spanErr(beforeRows16); err != nil {
			return nil, fmt.Errorf("beforeRows16\n%v", err)
		}
	} else {
		if err := s.spanErr(beforeRows15); err != nil {
			return nil, fmt.Errorf("beforeRows15\n%v", err)
		}
	}
	var ret []rowHeadS
	group := ""
	for row := 0; row < rows; row++ {
		r, err := s.parseCourseRow(cols16, row, subj, group)
		if err != nil {
			return nil, fmt.Errorf("%vth row\n%v", row, err)
		}
		switch r := r.(type) {
		case rowHeadS:
			ret = append(ret, r)
			group = r.group
		case rowBodyS:
			if len(ret) == 0 {
				return nil, fmt.Errorf("first rowBody")
			}
			ret[len(ret)-1].rowBody = append(ret[len(ret)-1].rowBody, r)
		case rowFootS:
			if len(ret) == 0 {
				return nil, fmt.Errorf("first rowFoot")
			}
			o := ret[len(ret)-1].rowBody
			o[len(o)-1].rowFoot = append(o[len(o)-1].rowFoot, r)
		}
	}
	if err := s.spanErr(afterRows); err != nil {
		return nil, fmt.Errorf("afterRows\n%v", err)
	}
	if err := s.spanErr(strconv.Itoa(38 + rows*6 + 13)); err != nil {
		return nil, fmt.Errorf("last tabindex\n%v", err)
	}
	if err := s.equalErr(last); err != nil {
		return nil, fmt.Errorf("last\n%v", err)
	}
	return ret, nil
}
