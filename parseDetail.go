package main

import (
	"fmt"
	"strconv"
)

var fuck = make(map[string]bool)

func (s *parser) parseDetail(stateNum, icsid, subj, code, career, units string, add, drop bool) (*courseDetailS, error) {
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

<link rel='stylesheet' type='text/css' href='/cs/public/cache/SSS_STYLESHEET_1.css' />
<title>Enquire Timetable</title>
<script language='JavaScript'>
var baseKey_win0 = "";
var altKey_win0 = "05678\xbc\xbe\xbf\xde";
var ctrlKey_win0 = "JK";
var saveWarningKeys_win0 = "";
var bTabOverTB_win0 = false;
var bTabOverPg_win0 = false;
var bTabOverNonPS_win0 = false;
var strCurrUrl='https://cusis.cuhk.edu.hk/psc/public/EMPLOYEE/HRMS/c/CU_SCR_MENU.CU_TMSR801.GBL?&PAGE=SSS_CRSE_OFFER_DTL';
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
	const afterStateNum = `' />
<input type='hidden' name='ICAction' value='None' />
<input type='hidden' name='ICXPos' value='0' />
<input type='hidden' name='ICYPos' value='0' />
<input type='hidden' name='ICFocus' value='' />
<input type='hidden' name='ICSaveWarningFilter' value='0' />
<input type='hidden' name='ICChanged' value='-1' />
<input type='hidden' name='ICResubmit' value='0' />
<input type='hidden' name='ICSID' value='`
	const afterICSID = `' />
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
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' class='PSPAGECONTAINER' cols='10' width='705'>
<tr>
<td width='0' height='3'></td>
<td width='4'></td>
<td width='4'></td>
<td width='4'></td>
<td width='1'></td>
<td width='183'></td>
<td width='4'></td>
<td width='450'></td>
<td width='20'></td>
<td width='35'></td>
</tr>
<tr>
<td height='35'></td>
<td colspan='8'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='635'>
<tr><td width='633'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='3' width='633' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='12' height='0'></td>
<td width='1'></td>
<td width='620'></td>
</tr>
<tr>
<td height='14' colspan='2'></td>
<td  valign='top' align='left'>
<span  class='PATRANSACTIONTITLE' >Teaching Timetable</span>
</td>
</tr>
<tr>
<td height='20'></td>
<td colspan='2'  valign='top' align='left'>
<span  class='PAPAGETITLE' >Course Detail</span>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='16' colspan='10'></td>
</tr>
<tr>
<td height='8' colspan='3'></td>
<td colspan='4'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='192'>
<tr><td width='190' height='6'>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='10' colspan='10'></td>
</tr>
<tr>
<td height='14' colspan='4'></td>
<td colspan='6' nowrap='nowrap'  valign='top' align='left'>
<span  class='PSHYPERLINK' >
<a name='DERIVED_SAA_CRS_RETURN_PB' id='DERIVED_SAA_CRS_RETURN_PB' tabindex='57' href="javascript:submitAction_win0(document.win0,'DERIVED_SAA_CRS_RETURN_PB');"  class='PSHYPERLINK'  title="Return" >Return to Teaching Timetable</a></span>
</td>
</tr>
<tr>
<td height='8' colspan='3'></td>
<td colspan='4'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='192'>
<tr><td width='190' height='6'>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='15' colspan='10'></td>
</tr>
<tr>
<td height='23' colspan='5'></td>
<td colspan='5'  valign='top' align='left'>
<span  class='PALEVEL0SECONDARY' >`
	/*ACCT 1111 - Foundations in Financial Accounting*/
	const afterTitle = `</span>
</td>
</tr>
<tr>
<td height='8' colspan='3'></td>
<td colspan='4'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='192'>
<tr><td width='190' height='6'>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='4' colspan='10'></td>
</tr>
<tr>
<td height='`
	const afterHeight = `' colspan='2'></td>
<td colspan='6'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='SSSGROUPBOXLTBLUEWBO'  width='646'>
<tr><td width='644'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='7' width='644' class='SSSGROUPBOXLTBLUE' style='border-style:none' >
<tr>
<td width='15' height='0'></td>
<td width='12'></td>
<td width='56'></td>
<td width='371'></td>
<td width='4'></td>
<td width='184'></td>
<td width='2'></td>
</tr>
<tr>
<td height='19'></td>
<td colspan='3' rowspan='2'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1' class='PSGROUPBOXNBO'  width='439'>
<tr><td class='SSSGROUPBOXLTBLUE' align='left'>Course Detail</td></tr>
<tr><td width='439'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='5' width='439' class='PSGROUPBOX' style='border-style:none' >
<tr>
<td width='12' height='13'></td>
<td width='168'></td>
<td width='9'></td>
<td width='227'></td>
<td width='23'></td>
</tr>
<tr>
<td height='2'></td>
<td colspan='2' rowspan='2'  valign='top' align='left'>
<label for='SSR_CRSE_OFF_VW_ACAD_CAREER$0' class='PSDROPDOWNLABEL' >Career</label>
</td>
</tr>
<tr>
<td height='20'></td>
<td colspan='2' rowspan='2'  valign='top' align='left'>
<span  class='PSDROPDOWNLIST_DISPONLY' >`
	/*Undergraduate*/
	const afterCareer = `</span>
</td>
</tr>
<tr>
<td height='1'></td>
<td colspan='2' rowspan='2'  valign='top' align='left'>
<label for='DERIVED_CRSECAT_UNITS_RANGE$0' class='PSEDITBOXLABEL' >Units</label>
</td>
</tr>
<tr>
<td height='21'></td>
<td colspan='2'  valign='top' align='left'>
<span  class='PSEDITBOX_DISPONLY' >`
	/*3.00*/
	const afterUnits = `</span>
</td>
</tr>
<tr>
<td height='17'></td>
<td colspan='2'  valign='top' align='left'>
<label for='SSR_CRSE_OFF_VW_GRADING_BASIS$0' class='PSDROPDOWNLABEL' >Grading Basis</label>
</td>
<td colspan='2'  valign='top' align='left'>
<span  class='PSDROPDOWNLIST_DISPONLY' >`
	const afterGrading = `</span>
</td>
</tr>
<tr>
<td height='5' colspan='2'></td>
<td colspan='2' rowspan='2'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  id='SSR_DUMMY_RECVW$scroll$0' width='236'>
<tr><td width='234'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='4' width='234' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='8' height='3'></td>
<td width='104'></td>
<td width='86'></td>
<td width='36'></td>
</tr>
<tr>
<td height='2' colspan='4'></td>
<td rowspan='2'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='128'>
<tr><td width='126' height='14'>
</td></tr>
</table>
</td>
</tr>
`
	const beforeComponent = `<tr>
<td height='14'></td>
<td rowspan='2'  valign='top' align='left'>
<span  class='PSEDITBOX_DISPONLY' >`
	const beforeRequired = `</span>
</td>
<td colspan='2' rowspan='2'  valign='top' align='left'>
<span  class='PSEDITBOX_DISPONLY' >`
	const afterRequired = `</span>
</td>
</tr>
<tr>
<td height='2'></td>
<td rowspan='2'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='128'>
<tr><td width='126' height='14'>
</td></tr>
</table>
</td>
</tr>
`
	const beforeLastComponent = `<tr>
<td height='13'></td>
<td  valign='top' align='left'>
<span  class='PSEDITBOX_DISPONLY' >`
	const beforeLastRequired = `</span>
</td>
<td colspan='2'  valign='top' align='left'>
<span  class='PSEDITBOX_DISPONLY' >`
	const afterLastRequired = `</span>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
`
	const courseComponents1a = `<td height='`
	const courseComponents1b = `'></td>
<td rowspan='2'  valign='top' align='left'>
<label for='SR_LBL_WRK_CRSE_COMPONENT_LBL$0' class='PSEDITBOXLABEL' >Course Components</label>
</td>
</tr>
<tr>
<td height='9'></td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='`
	const courseComponents1c = `'></td>
<td rowspan='6'></td>
<td rowspan='2'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='183'>
`
	const viewClassSection = `<tr><td width='181'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='6' width='181' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='3' height='0'></td>
<td width='4'></td>
<td width='4'></td>
<td width='124'></td>
<td width='44'></td>
<td width='2'></td>
</tr>
<tr>
<td height='16' colspan='2'></td>
<td colspan='2'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='127'>
<tr><td width='125'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='2' width='125' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='3' height='5'></td>
<td width='122'></td>
</tr>
<tr>
<td height='9'></td>
<td nowrap='nowrap'  valign='top' align='left'>
<span  class='SSSBUTTON_ACTIONLINK' >
<a name='DERIVED_SAA_CRS_SSR_PB_GO' id='DERIVED_SAA_CRS_SSR_PB_GO' tabindex='84' href="javascript:submitAction_win0(document.win0,'DERIVED_SAA_CRS_SSR_PB_GO');"  class='SSSBUTTON_ACTIONLINK' >View Class Sections</a></span>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='4' colspan='6'></td>
</tr>
<tr>
<td height='18' colspan='3'></td>
<td colspan='3' nowrap='nowrap'  valign='top' align='left'>
<span  class='SSSBUTTON_ACTIONLINK' >
<a name='CU_DERIVED_CUR_CU_CRSE_OUT_BTN' id='CU_DERIVED_CUR_CU_CRSE_OUT_BTN' tabindex='85' href="javascript:submitAction_win0(document.win0,'CU_DERIVED_CUR_CU_CRSE_OUT_BTN');"  class='SSSBUTTON_ACTIONLINK' >View Course Outcome</a></span>
</td>
</tr>
<tr>
<td height='21'></td>
<td colspan='4'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='175'>
<tr><td width='173'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='2' width='173' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='8' height='0'></td>
<td width='165'></td>
</tr>
<tr>
<td height='11'></td>
<td  valign='top' align='left'>
<span  class='PSLONGEDITBOX' >&nbsp;</span>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='3' colspan='6'></td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
`
	const notScheduled = `<tr><td width='181'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='5' width='181' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='3' height='0'></td>
<td width='8'></td>
<td width='164'></td>
<td width='4'></td>
<td width='2'></td>
</tr>
<tr>
<td height='25'></td>
<td colspan='2'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='171'>
<tr><td width='169'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='2' width='169' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='7' height='0'></td>
<td width='162'></td>
</tr>
<tr>
<td height='15'></td>
<td  valign='top' align='left'>
<span  class='PSPSMALLTEXT' >***&nbsp; This course has not been scheduled.&nbsp; ***</span>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='4' colspan='5'></td>
</tr>
<tr>
<td height='18' colspan='2'></td>
<td colspan='3' nowrap='nowrap'  valign='top' align='left'>
<span  class='SSSBUTTON_ACTIONLINK' >
<a name='CU_DERIVED_CUR_CU_CRSE_OUT_BTN' id='CU_DERIVED_CUR_CU_CRSE_OUT_BTN' tabindex='85' href="javascript:submitAction_win0(document.win0,'CU_DERIVED_CUR_CU_CRSE_OUT_BTN');"  class='SSSBUTTON_ACTIONLINK' >View Course Outcome</a></span>
</td>
</tr>
<tr>
<td height='21'></td>
<td colspan='3'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='175'>
<tr><td width='173'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='2' width='173' class='PABACKGROUNDINVISIBLE' style='border-style:none' >
<tr>
<td width='8' height='0'></td>
<td width='165'></td>
</tr>
<tr>
<td height='11'></td>
<td  valign='top' align='left'>
<span  class='PSLONGEDITBOX' >&nbsp;</span>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='3' colspan='5'></td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
`
	const enrollmentInformation = `<tr>
<td height='27'></td>
<td colspan='3' rowspan='2'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1' class='PSGROUPBOXNBO'  width='403'>
<tr><td class='SSSGROUPBOXLTBLUE' align='left'>Enrollment Information</td></tr>
<tr><td width='403'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='3' width='403' class='PSGROUPBOX' style='border-style:none' >
<tr>
<td width='12' height='11'></td>
<td width='177'></td>
<td width='214'></td>
</tr>
`
	const noEnrollmentInformation = `<tr>
<td height='0' colspan='4'></td>
</tr>
<tr>
<td height='0' colspan='4'></td>
</tr>
`
	const addConsent = `<tr>
<td height='20'></td>
<td  valign='top' align='left'>
<label for='SSR_CRSE_OFF_VW_CONSENT$0' class='PSDROPDOWNLABEL' >Add Consent</label>
</td>
`
	const dropConsent = `<tr>
<td height='20'></td>
<td  valign='top' align='left'>
<label for='SSR_CRSE_OFF_VW_SSR_DROP_CONSENT$0' class='PSDROPDOWNLABEL' >Drop Consent</label>
</td>
`
	const beforeConsent1 = `<td  valign='top' align='left'>
<span  class='PSDROPDOWNLIST_DISPONLY' >`
	const beforeConsent2 = `<td rowspan='2'  valign='top' align='left'>
<span  class='PSDROPDOWNLIST_DISPONLY' >`
	const afterConsent = `Consent Required</span>
</td>
</tr>
`
	const beforeAttribute = `<tr>
<td height='32'></td>
<td  valign='top' align='left'>
<label for='DERIVED_CRSECAT_SSR_CRSE_ATTR_LONG$0' class='PSEDITBOXLABEL' >Course Attribute</label>
</td>
<td  valign='top' align='left'>
<span  class='PSLONGEDITBOX' >`
	const afterAttribute = `</span>
</td>
</tr>
`
	const beforeRequirement = `<tr>
<td height='1'></td>
<td rowspan='2'  valign='top' align='left'>
<label for='SSR_CRSE_OFF_VW_RQRMNT_GROUP$0' class='PSEDITBOXLABEL' >Enrollment Requirement</label>
</td>
</tr>
<tr>
<td height='19'></td>
<td  valign='top' align='left'>
<span  class='PSEDITBOX_DISPONLY' >`
	const afterRequirement = `</span>
</td>
</tr>
`
	const afterEnrollmentInformationa = `</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='`
	const afterEnrollmentInformationb = `'></td>
</tr>
`
	const noDescription = `<tr>
<td height='0' colspan='4'></td>
</tr>
<tr>
<td height='7' colspan='2'></td>
<td  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='56'>
<tr><td width='54' height='5'>
</td></tr>
</table>
</td>
</tr>
`
	const beforeDescription = `<tr>
<td height='88'></td>
<td colspan='3'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1' class='PSGROUPBOXNBO'  width='403'>
<tr><td class='SSSGROUPBOXLTBLUE' align='left'>Description</td></tr>
<tr><td width='403'>
<table  id='ACE_width' border='0' cellpadding='0' cellspacing='0' cols='2' width='403' class='PSGROUPBOX' style='border-style:none' >
<tr>
<td width='15' height='11'></td>
<td width='388'></td>
</tr>
<tr>
<td height='64'></td>
<td  valign='top' align='left'>
<span  class='PSLONGEDITBOX' >`
	const afterDescription = `</span>
</td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='7' colspan='2'></td>
<td  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='56'>
<tr><td width='54' height='5'>
</td></tr>
</table>
</td>
</tr>
`
	const beforeLast = `<tr>
<td height='6' colspan='4'></td>
</tr>
</table>
</td></tr>
</table>
</td>
</tr>
`
	const last = `<tr>
<td height='16' colspan='10'></td>
</tr>
<tr>
<td height='8' colspan='2'></td>
<td colspan='4'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='192'>
<tr><td width='190' height='6'>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='6' colspan='10'></td>
</tr>
<tr>
<td height='14' colspan='3'></td>
<td colspan='7' nowrap='nowrap'  valign='top' align='left'>
<span  class='PSHYPERLINK' >
<a name='DERIVED_SAA_CRS_RETURN_PB$160$' id='DERIVED_SAA_CRS_RETURN_PB$160$' tabindex='478' href="javascript:submitAction_win0(document.win0,'DERIVED_SAA_CRS_RETURN_PB$160$');"  class='PSHYPERLINK'  title="Return" >Return to Teaching Timetable</a></span>
</td>
</tr>
<tr>
<td height='8' colspan='2'></td>
<td colspan='4'  valign='top' align='left'>
<table cellpadding='0' cellspacing='0' cols='1'  class='PABACKGROUNDINVISIBLEWBO'  width='192'>
<tr><td width='190' height='6'>
</td></tr>
</table>
</td>
</tr>
<tr>
<td height='12' colspan='10'></td>
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
	err := s.spanErr(beforeStateNum)
	if err != nil {
		return nil, fmt.Errorf("beforeStateNum\n%v", err)
	}
	if err = s.spanErr(stateNum); err != nil {
		return nil, fmt.Errorf("stateNum\n%v", err)
	}
	if err = s.spanErr(afterStateNum); err != nil {
		return nil, fmt.Errorf("afterStateNum\n%v", err)
	}
	if err = s.spanErr(icsid); err != nil {
		return nil, fmt.Errorf("icsid\n%v", err)
	}
	if err = s.spanErr(afterICSID); err != nil {
		return nil, fmt.Errorf("afterICSID\n%v", err)
	}
	if err = s.spanErr(subj + ` ` + code + ` - `); err != nil {
		return nil, fmt.Errorf("title\n%v", err)
	}
	var ret courseDetailS
	if ret.title, err = s.splitErr(afterTitle); err != nil {
		return nil, fmt.Errorf("afterTitle\n%v", err)
	}
	if _, err = s.splitErr(afterHeight); err != nil {
		return nil, fmt.Errorf("afterHeight\n%v", err)
	}
	if err = s.spanLastErr(last); err != nil {
		return nil, fmt.Errorf("last\n%v", err)
	}
	detailString := s.s
	if err = s.spanErr(career); err != nil {
		return nil, fmt.Errorf("career\n%v", err)
	}
	if err = s.spanErr(afterCareer); err != nil {
		return nil, fmt.Errorf("afterCareer\n%v", err)
	}
	if ret.units, err = s.splitErr(afterUnits); err != nil {
		return nil, fmt.Errorf("afterUnits\n%v", err)
	}
	if ret.grading, err = s.splitErr(afterGrading); err != nil {
		return nil, fmt.Errorf("afterGrading\n%v", err)
	}
	for {
		if s.span(beforeLastComponent) {
			t, err := s.splitErr(beforeLastRequired)
			if err != nil {
				return nil, fmt.Errorf(
					"beforeLastRequired %v\n%v",
					len(ret.components),
					err,
				)
			}
			u, err := s.splitErr(afterLastRequired)
			if err != nil {
				return nil, fmt.Errorf(
					"afterLastRequired %v\n%v",
					len(ret.components),
					err,
				)
			}
			ret.components = append(ret.components, [2]string{t, u})
			break
		}
		if err = s.spanErr(beforeComponent); err != nil {
			return nil, fmt.Errorf(
				"beforeComponent %v\n%v",
				len(ret.components),
				err,
			)
		}
		t, err := s.splitErr(beforeRequired)
		if err != nil {
			return nil, fmt.Errorf(
				"beforeRequired %v\n%v",
				len(ret.components),
				err,
			)
		}
		u, err := s.splitErr(afterRequired)
		if err != nil {
			return nil, fmt.Errorf(
				"afterRequired %v\n%v",
				len(ret.components),
				err,
			)
		}
		ret.components = append(ret.components, [2]string{t, u})
	}
	if err = s.spanErr(courseComponents1a); err != nil {
		return nil, fmt.Errorf("courseComponents1a\n%v", err)
	}
	if err = s.spanErr(strconv.Itoa(16*len(ret.components) - 1)); err != nil {
		return nil, fmt.Errorf("courseComponents1a_\n%v", err)
	}
	if err = s.spanErr(courseComponents1b); err != nil {
		return nil, fmt.Errorf("courseComponents1b\n%v", err)
	}
	if err = s.spanErr(strconv.Itoa(16*len(ret.components) + 81)); err != nil {
		return nil, fmt.Errorf("courseComponents1b_\n%v", err)
	}
	if err = s.spanErr(courseComponents1c); err != nil {
		return nil, fmt.Errorf("courseComponents1c\n%v", err)
	}
	if s.span(viewClassSection) {
		ret.scheduled = true
	} else if err = s.spanErr(notScheduled); err != nil {
		return nil, fmt.Errorf("notScheduled\n%v", err)
	}
	if s.span(enrollmentInformation) {
		if s.span(addConsent) {
			if s.span(beforeConsent1) {
			} else if err = s.spanErr(beforeConsent2); err != nil {
				return nil, fmt.Errorf("beforeConsent2\n%v", err)
			}
			if ret.addConsent, err = s.splitErr(afterConsent); err != nil {
				return nil, fmt.Errorf("afterConsent\n%v", err)
			}
		}
		if s.span(dropConsent) {
			if s.span(beforeConsent1) {
			} else if err = s.spanErr(beforeConsent2); err != nil {
				return nil, fmt.Errorf("beforeConsent2\n%v", err)
			}
			if ret.dropConsent, err = s.splitErr(afterConsent); err != nil {
				return nil, fmt.Errorf("afterConsent\n%v", err)
			}
		}
		if s.span(beforeRequirement) {
			if ret.requirement, err = s.splitErr(afterRequirement); err != nil {
				return nil, fmt.Errorf("afterRequirementa\n%v", err)
			}
		}
		if s.span(beforeAttribute) {
			if ret.attribute, err = s.splitErr(afterAttribute); err != nil {
				return nil, fmt.Errorf("afterAttribute\n%v", err)
			}
		}
		if err = s.spanErr(afterEnrollmentInformationa); err != nil {
			return nil, fmt.Errorf("afterEnrollmentInformationa\n%v", err)
		}
		if t, err := s.splitErr(afterEnrollmentInformationb); err == nil {
			if !fuck[t] {
				fuck[t] = true
				fmt.Println(t, ret)
			}
		} else {
			return nil, fmt.Errorf("afterEnrollmentInformationb\n%v", err)
		}
	} else if err = s.spanErr(noEnrollmentInformation); err != nil {
		return nil, fmt.Errorf("noEnrollmentInformation\n%v", err)
	}
	if s.span(beforeDescription) {
		if ret.description, err = s.splitErr(afterDescription); err != nil {
			return nil, fmt.Errorf("afterDescription\n%v", err)
		}
	} else if err = s.spanErr(noDescription); err != nil {
		return nil, fmt.Errorf("noDescription\n%v", err)
	}
	if err = s.equalErr(beforeLast); err != nil {
		return nil, fmt.Errorf("beforeLast\n%v", err)
	}
	s.s = detailString
	return &ret, nil
}
