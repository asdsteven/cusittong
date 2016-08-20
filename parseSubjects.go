package main

import (
	"fmt"
)

func (s *parser) parseSubjects(icsid string) ([][2]string, error) {
	const beforeFocus = `<html dir='ltr' lang='en'>
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
var baseKey_win0 = "\x1b\r\n";
var altKey_win0 = "05678\xbc\xbe\xbf\xde\xbc\xbe\xde1";
var ctrlKey_win0 = "JK";
var saveWarningKeys_win0 = "";
var bTabOverTB_win0 = false;
var bTabOverPg_win0 = false;
var bTabOverNonPS_win0 = false;
var strCurrUrl='https://cusis.cuhk.edu.hk/psc/public/EMPLOYEE/HRMS/c/CU_SCR_MENU.CU_TMSR801.GBL?&PAGE=CU_TMSR801_ENTRY';
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
<body class='PSSRCHPAGE'  onload="
setFocus_win0('`
	const focus1 = `SEARCH_RESULT1`
	const focus2 = `CU_SUBJECT_DVW_SUBJECT`
	const beforeICSID = `',-1);
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
<input type='hidden' name='ICStateNum' value='3' />
<input type='hidden' name='ICAction' value='None' />
<input type='hidden' name='ICXPos' value='0' />
<input type='hidden' name='ICYPos' value='0' />
<input type='hidden' name='ICFocus' value='' />
<input type='hidden' name='ICSaveWarningFilter' value='0' />
<input type='hidden' name='ICChanged' value='-1' />
<input type='hidden' name='ICResubmit' value='0' />
<input type='hidden' name='ICSID' value='`
	const afterICSID = `' />
<span class='PSSRCHACTION' >Look Up Course Subject</span><br /><br />
<table cellpadding='0' cellspacing='2' class='PSPAGECONTAINER' >
<tr valign='top'>
<td>
<label for='#ICKeySelect'><span class='PSSRCHSUBTITLE' >Search by:&nbsp;&nbsp;&nbsp;</span></label><select name='#ICKeySelect' id='#ICKeySelect' tabindex='12' class='PSDROPDOWNLIST' ' onchange="submitAction_win0(this.form,this.name);">
<option value='2'>Description</option>
<option value='0' selected='selected'>Subject Area</option>
</select>
</td>
<td>
<span class='PSDROPDOWNLIST' >begins with</span>
</td>
<td>
<input type='text' name='CU_SUBJECT_DVW_SUBJECT' id='CU_SUBJECT_DVW_SUBJECT' tabindex='14' value=""  class='PSEDITBOX' style="width:140px; " maxlength='8'  />
</td>
</tr>
</table>
`
	const noResult = `<br /><input type=button id='#ICSearch' name='#ICSearch'  class='PSPUSHBUTTONTBLOOKUP'  value='Look Up' onclick="javascript:submitAction_win0(document.win0, '#ICSearch');" tabindex='24' PSaccesskey='\' alt='Look up (Alt+1)' title='Look up (Alt+1)'>
<input type=button id='#ICCancel' name='#ICCancel'  class='PSPUSHBUTTONTBCANCEL'  value='Cancel' onclick="javascript:submitAction_win0(document.win0, '#ICCancel');" tabindex='25' alt='Cancel (Esc)' title='Cancel (Esc)'>
<a name='#ICAdvSearch' class='PSHYPERLINK' href="javascript: submitAction_win0(document.win0,'#ICAdvSearch');" tabindex='26'>Advanced Lookup</a>
<br /><br />
<span class='PSSRCHINSTRUCTIONS' >No matching values were found.</span>
<div id='DetachDiv' height=0 width=0 frameborder=0></div>
</form>
<a name='ICLastAnchor_win0'></a>
</body>
</html>
`
	const beforeViewAll1 = `<br /><input type=button id='#ICSearch' name='#ICSearch'  class='PSPUSHBUTTONTBLOOKUP'  value='Look Up' onclick="javascript:submitAction_win0(document.win0, '#ICSearch');" tabindex='24' PSaccesskey='\' alt='Look up (Alt+1)' title='Look up (Alt+1)'>
<input type=button id='#ICCancel' name='#ICCancel'  class='PSPUSHBUTTONTBCANCEL'  value='Cancel' onclick="javascript:submitAction_win0(document.win0, '#ICCancel');" tabindex='25' alt='Cancel (Esc)' title='Cancel (Esc)'>
<a name='#ICAdvSearch' class='PSHYPERLINK' href="javascript: submitAction_win0(document.win0,'#ICAdvSearch');" tabindex='26'>Advanced Lookup</a>
<br /><br />
<span class='PSSRCHSUBTITLE' >Search Results</span>
<br />
<table class='PSSRCHRESULTSTITLE' dir='>ltr' cellpadding='1' cellspacing='0'>
<tr valign='baseline'>
<td align='left' nowrap='nowrap'><span class='PSSRCHRESULTSHYPERLINKD' >View All</span>
&nbsp;&nbsp;&nbsp;</td>
<td align='right' nowrap='nowrap'>
<span class='PSSRCHRESULTSHYPERLINKD' >First</span>
&nbsp;<img src='/cs/public/cache/PT_PREVIOUSROW_D_1.gif' name='PrevPage' alt='Show previous rows (inactive button) (Alt+,)' title='Show previous rows (inactive button) (Alt+,)' align='middle' />
&nbsp;<span style="font-size:80%;">1-`
	const afterViewAll1 = `</span>
&nbsp;<img src='/cs/public/cache/PT_NEXTROW_D_1.gif' name='NextPage' alt='Show next rows (inactive button) (Alt+.)' title='Show next rows (inactive button) (Alt+.)' align='middle' />
&nbsp;<span class='PSSRCHRESULTSHYPERLINKD' >Last</span>
</td></tr>
<tr align='left'><th scope='col' class='PSSRCHRESULTSHDR' ><a tabindex='27' class='PSSRCHRESULTSHDR' name='#ICSortCol0' href="javascript:submitAction_win0(document.win0,'#ICSortCol0');" title='Click column heading to sort ascending'>Subject Area</a></th><th scope='col' class='PSSRCHRESULTSHDR' ><a tabindex='28' class='PSSRCHRESULTSHDR' name='#ICSortCol1' href="javascript:submitAction_win0(document.win0,'#ICSortCol1');" title='Click column heading to sort ascending'>Description</a></th></tr>
`
	const beforeViewSome = `<br /><input type=button id='#ICSearch' name='#ICSearch'  class='PSPUSHBUTTONTBLOOKUP'  value='Look Up' onclick="javascript:submitAction_win0(document.win0, '#ICSearch');" tabindex='24' PSaccesskey='\' alt='Look up (Alt+1)' title='Look up (Alt+1)'>
<input type=button id='#ICCancel' name='#ICCancel'  class='PSPUSHBUTTONTBCANCEL'  value='Cancel' onclick="javascript:submitAction_win0(document.win0, '#ICCancel');" tabindex='25' alt='Cancel (Esc)' title='Cancel (Esc)'>
<a name='#ICAdvSearch' class='PSHYPERLINK' href="javascript: submitAction_win0(document.win0,'#ICAdvSearch');" tabindex='26'>Advanced Lookup</a>
<br /><br />
<span class='PSSRCHSUBTITLE' >Search Results</span>
<br />
<table class='PSSRCHRESULTSTITLE' dir='>ltr' cellpadding='1' cellspacing='0'>
<tr valign='baseline'>
<td align='left' nowrap='nowrap'><span class='PSSRCHRESULTSHYPERLINK' ><a name='#ICViewAll' class='PSSRCHRESULTSHYPERLINK' href="javascript: submitAction_win0(document.win0,'#ICViewAll');" tabindex='27'>View All</a></span>
&nbsp;&nbsp;&nbsp;</td>
<td align='right' nowrap='nowrap'>
<span class='PSSRCHRESULTSHYPERLINKD' >First</span>
&nbsp;<img src='/cs/public/cache/PT_PREVIOUSROW_D_1.gif' name='PrevPage' alt='Show previous rows (inactive button) (Alt+,)' title='Show previous rows (inactive button) (Alt+,)' align='middle' />
&nbsp;<span style="font-size:80%;">1-100 of `
	const last = `</table>
<div id='DetachDiv' height=0 width=0 frameborder=0></div>
</form>
<a name='ICLastAnchor_win0'></a>
</body>
</html>
`
	if err := s.spanErr(beforeFocus); err != nil {
		return nil, fmt.Errorf("beforeFocus\n%v", err)
	}
	var focus int
	if s.span(focus1) {
		focus = 1
	} else if err := s.spanErr(focus2); err == nil {
		focus = 2
	} else {
		return nil, fmt.Errorf("focus2\n%v", err)
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
	if focus == 2 {
		if err := s.equalErr(noResult); err != nil {
			return nil, fmt.Errorf("noResult\n%v", err)
		}
		return [][2]string{}, nil
	}
	if err := s.spanLastErr(last); err != nil {
		return nil, fmt.Errorf("last\n%v", err)
	}
	if s.span(beforeViewAll1) {
		rows, err := s.parseRows()
		if err != nil {
			return nil, fmt.Errorf("parseRows\n%v", err)
		}
		if err := s.spanErr(afterViewAll1); err != nil {
			return nil, fmt.Errorf("afterViewAll1\n%v", err)
		}
		return s.parseSubjectRows(rows, 29)
	}
	err := s.spanErr(beforeViewSome)
	if err != nil {
		return nil, fmt.Errorf("beforeViewSome\n%v", err)
	}
	return nil, nil
}
