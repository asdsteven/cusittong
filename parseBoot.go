package main

func (s *parser) parseBoot() (icsid string, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(parseError); ok {
				err = e.error
			} else {
				panic(r)
			}
		}
	}()
	s.spanPanic(`before ICSID`, `<html dir='ltr' lang='en'>
<!-- Copyright (c) 2000, 2007, Oracle. All rights reserved. -->
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<script language='JavaScript'>
var  totalTimeoutMilliseconds = 2073600000; 
var  warningTimeoutMilliseconds = 2073600000; 
var timeOutURL = 'https://cusis.cuhk.edu.hk/psc/public/EMPLOYEE/HRMS/?cmd=expire';
var timeoutWarningPageURL = 'https://cusis.cuhk.edu.hk/psc/public/EMPLOYEE/HRMS/s/WEBLIB_TIMEOUT.PT_TIMEOUTWARNING.FieldFormula.IScript_TIMEOUTWARNING';
</script> 

<link rel='stylesheet' type='text/css' href='/cs/public/cache/SSS_STYLESHEET_1.css' />
<title>Enquire Timetable</title>
<script language='JavaScript'>
var baseKey_win0 = "\x1b\r\n";
var altKey_win0 = "05678\xbc\xbe\xbf\xde";
var ctrlKey_win0 = "JK";
var saveWarningKeys_win0 = "";
var bTabOverTB_win0 = false;
var bTabOverPg_win0 = false;
var bTabOverNonPS_win0 = false;
var strCurrUrl='https://cusis.cuhk.edu.hk/psc/public/EMPLOYEE/HRMS/c/CU_SCR_MENU.CU_TMSR801.GBL&PAGE=CU_TMSR801_ENTRY';
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
setFocus_win0('CU_RC_TMSR801_SUBJECT',-1);
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
<input type='hidden' name='ICStateNum' value='1' />
<input type='hidden' name='ICAction' value='None' />
<input type='hidden' name='ICXPos' value='0' />
<input type='hidden' name='ICYPos' value='0' />
<input type='hidden' name='ICFocus' value='' />
<input type='hidden' name='ICSaveWarningFilter' value='0' />
<input type='hidden' name='ICChanged' value='-1' />
<input type='hidden' name='ICResubmit' value='0' />
<input type='hidden' name='ICSID' value='`)
	icsid = s.takePanic(`ICSID`, 12)
	s.spanPanic(`after ICSID`, `' />
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
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' class='PSPAGECONTAINER' cols='8' width='731'>
<tr>
<td width='4' height='0'></td>
<td width='8'></td>
<td width='4'></td>
<td width='424'></td>
<td width='4'></td>
<td width='76'></td>
<td width='28'></td>
<td width='183'></td>
</tr>
<tr>
<td height='52'></td>
<td colspan='6'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='543'>
<tr><td class='PAGROUPBOXLABELINVISIBLE' align='left'>Groupbox</td></tr>
<tr><td width='541'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='3' width='541' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='7' height='0'></td>
<td width='2'></td>
<td width='532'></td>
</tr>
<tr>
<td height='22' colspan='2'></td>
<td  valign='top' align='left'>
<span  class='PATRANSACTIONTITLE' >Teaching Timetable</span>
</td>
</tr>
<tr>
<td height='18'></td>
<td colspan='2'  valign='top' align='left'>
<label for='DERIVED_CLSRCH_SSR_CLASS_LBL' class='PAPAGETITLE' >Enter Search Criteria</label>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='8' colspan='8'></td>
</tr>
<tr>
<td height='21' colspan='3'></td>
<td  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='423'>
<tr><td width='421'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='3' width='421' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='3' height='0'></td>
<td width='144'></td>
<td width='274'></td>
</tr>
<tr>
<td height='15'></td>
<td  valign='top' align='left'>
<label for='CLASS_SRCH_WRK2_ACAD_CAREER' class='PSDROPDOWNLABEL' >Course Career</label>
</td>
<td  valign='top' align='left'>
<select name='CLASS_SRCH_WRK2_ACAD_CAREER' id='CLASS_SRCH_WRK2_ACAD_CAREER' tabindex='48' size='1'  class='PSDROPDOWNLIST' style="width:180px; " onchange="submitAction_win0(this.form,this.name);" >
<option value=""></option>
<option value="PGDE">Postgraduate - PGDE</option>
<option value="RPG">Postgraduate - Research</option>
<option value="TPG">Postgraduate - Taught</option>
<option value="UG" selected='selected'>Undergraduate</option>
</select>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='4' colspan='8'></td>
</tr>
<tr>
<td height='25' colspan='3'></td>
<td colspan='2'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='427'>
<tr><td width='425'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='3' width='425' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='3' height='0'></td>
<td width='144'></td>
<td width='278'></td>
</tr>
<tr>
<td height='15'></td>
<td  valign='top' align='left'>
<label for='CLASS_SRCH_WRK2_STRM$50$' class='PSDROPDOWNLABEL' >Term</label>
</td>
<td  valign='top' align='left'>
<select name='CLASS_SRCH_WRK2_STRM$50$' id='CLASS_SRCH_WRK2_STRM$50$' tabindex='25' size='1'  class='PSDROPDOWNLIST' style="width:220px; " onchange="submitAction_win0(this.form,this.name);" >
<option value=""></option>
`)
	/*<option value="1940">						2016-17 Acad Year (Medicine)</option>
	  <option value="1945" selected='selected'>					2016-17 Term 1</option>
	  <option value="1955">				2016-17 Term 2</option>
	  <option value="1965">			2016-17 Term 3</option>
	  <option value="1970">		2016-17 Term 4</option>
	  <option value="1990">	2016-17 Summer Session</option>*/
	s.spanLastPanic(`last`, `</select>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='0' colspan='8'></td>
</tr>
<tr>
<td height='93' colspan='2'></td>
<td colspan='4'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PSGROUPBOXWBO'  width='507'>
<tr><td class='PSGROUPBOXLABEL' align='left'>Class Search Criteria</td></tr>
<tr><td width='505'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='4' width='505' class='PSGROUPBOX' style='border-style:none' >
<tr>
<td width='3' height='0'></td>
<td width='120'></td>
<td width='376'></td>
<td width='6'></td>
</tr>
<tr>
<td height='28'></td>
<td colspan='2'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='495'>
<tr><td width='493'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='4' width='493' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='3' height='0'></td>
<td width='144'></td>
<td width='125'></td>
<td width='221'></td>
</tr>
<tr>
<td height='4' colspan='2'></td>
<td rowspan='3' nowrap='nowrap'  valign='top' align='left'>
<input type='text' name='CU_RC_TMSR801_SUBJECT' id='CU_RC_TMSR801_SUBJECT' tabindex='38' value=""  class='PSEDITBOX' style="width:58px; " maxlength='8' onchange="oChange_win0=this;"  /><a name='CU_RC_TMSR801_SUBJECT$prompt' id='CU_RC_TMSR801_SUBJECT$prompt' tabindex='39' href="javascript:submitAction_win0(document.win0,'CU_RC_TMSR801_SUBJECT$prompt');"><img src="/cs/public/cache/PT_PROMPT_LOOKUP_1.gif" alt='Look up Course Subject (Alt+5)' title='Look up Course Subject (Alt+5)' border='0' align='absmiddle' /></a>
</td>
</tr>
<tr>
<td height='1'></td>
<td rowspan='2'  valign='top' align='left'>
<label for='CU_RC_TMSR801_SUBJECT' class='PSEDITBOXLABEL' >Course Subject</label>
</td>
</tr>
<tr>
<td height='15'></td>
<td  valign='top' align='left'>
<span  class='PSEDITBOX_DISPONLY' >&nbsp;</span>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='2' colspan='4'></td>
</tr>
<tr>
<td height='23'></td>
<td  valign='top' align='left'>
<span  class='PSTEXT' >Or</span>
</td>
</tr>
<tr>
<td height='26'></td>
<td colspan='2'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='495'>
<tr><td width='493'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='4' width='493' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='3' height='0'></td>
<td width='144'></td>
<td width='125'></td>
<td width='221'></td>
</tr>
<tr>
<td height='5'></td>
<td rowspan='2'  valign='top' align='left'>
<label for='CU_RC_TMSR801_ACAD_ORG' class='PSEDITBOXLABEL' >Course Offering Dept</label>
</td>
<td rowspan='2' nowrap='nowrap'  valign='top' align='left'>
<input type='text' name='CU_RC_TMSR801_ACAD_ORG' id='CU_RC_TMSR801_ACAD_ORG' tabindex='42' value=""  class='PSEDITBOX' style="width:60px; " maxlength='10' onchange="oChange_win0=this;"  /><a name='CU_RC_TMSR801_ACAD_ORG$prompt' id='CU_RC_TMSR801_ACAD_ORG$prompt' tabindex='43' href="javascript:submitAction_win0(document.win0,'CU_RC_TMSR801_ACAD_ORG$prompt');"><img src="/cs/public/cache/PT_PROMPT_LOOKUP_1.gif" alt='Look up Course Offering Dept (Alt+5)' title='Look up Course Offering Dept (Alt+5)' border='0' align='absmiddle' /></a>
</td>
</tr>
<tr>
<td height='15'></td>
<td  valign='top' align='left'>
<span  class='PSEDITBOX_DISPONLY' >&nbsp;</span>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='6' colspan='4'></td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='20' colspan='8'></td>
</tr>
<tr>
<td height='28'></td>
<td colspan='5'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='515'>
<tr><td width='513'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='2' width='513' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='439' height='13'></td>
<td width='74'></td>
</tr>
<tr>
<td height='13'></td>
<td  valign='top' align='left'>
<span  class='SSSBUTTON_CONFIRMLINK' >
<a name='CU_RC_TMSR801_SSR_PB_CLASS_SRCH' id='CU_RC_TMSR801_SSR_PB_CLASS_SRCH' tabindex='52' href="javascript:submitAction_win0(document.win0,'CU_RC_TMSR801_SSR_PB_CLASS_SRCH');"  class='SSSBUTTON_CONFIRMLINK' >Search</a></span>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='16' colspan='8'></td>
</tr>
</table>
</td></tr>
</table>
<div id='DetachDiv' height=0 width=0 frameborder=0></div>
</form>
<a name='ICLastAnchor_win0'></a>
</body>
</html>
`)
	return
}
