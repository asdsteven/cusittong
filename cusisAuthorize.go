package main

const cusisError = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN" "http://www.w3.org/TR/REC-html40/loose.dtd"`

const cusisAuthorizeError = `<html dir=ltr lang=en>
<HEAD>
<!--
* ******************************************************************
* Added by: Kenneth Tsang IBM
* To prevent users from logging in using the public site
* by redirecting to the CUSIS login page.
* ******************************************************************
-->
<meta HTTP-EQUIV='Refresh' CONTENT='0; URL=../../psp/csprd/?cmd=login'>
<!--
* ******************************************************************
* ORACLE CONFIDENTIAL.  For authorized use only.  Except for as
* expressly authorized by Oracle, do not disclose, copy, reproduce,
* distribute, or modify.
* ******************************************************************
*
-->
<TITLE>Oracle | PeopleSoft Enterprise 8 Sign-in</TITLE>

<link rel="stylesheet" href="/public/styles.css" rel="stylesheet" type="text/css">
<style type="text/css">
<!--
a:link {
        color: #48598c;
}
a:visited {
        color: #48598C;
}
a:hover {
        color: #48598C;
}
a:active {
        color: #48598C;
}
body {
        margin-left: 0px;
        margin-top: 0px;
        margin-right: 0px;
        margin-bottom: 0px;
}
-->
</style></HEAD>
<script LANGUAGE="JavaScript">
    function signin(form)
	  {
	  var docLoc=new String(document.location);
	  var iLast= docLoc.lastIndexOf("?&");
	  if (docLoc.length == (iLast+2))
	      {
	      docLoc = docLoc.substring(0, iLast);
	      }

	  if (docLoc.indexOf("?cmd=")==-1 && docLoc.indexOf("?")!=-1)
	     {
		 if (docLoc.indexOf("&cmd=login")==-1)
	        {
	        var i=docLoc.length - 1;
	        var j= docLoc.lastIndexOf("&");
	        if (j!=-1 && i==j)
		       form.action=docLoc+form.action.substring(form.action.indexOf("?")+1,form.action.length);
	        else
		       form.action=docLoc+"&"+form.action.substring(form.action.indexOf("?")+1,form.action.length);
	        }
	  else
	        form.action=docLoc.substring(0,docLoc.indexOf("&cmd=login"))+"&cmd=login"+docLoc.substring(docLoc.indexOf("&languageCd="),docLoc.length);
	     }

	     var now=new Date();
	     form.timezoneOffset.value=now.getTimezoneOffset();
	     return ;
         }

    function setFocus()
      {
      try
         {document.login.userid.focus()}
      catch (e)
         {};
      return;
      }

function setErrorImg()
    {
    var login_error = document.getElementById('login_error').innerHTML;
    var discovery_error = document.getElementById('discovery_error').innerHTML;
    login_error = login_error.replace(/^\s+/,"");       // delete leading spaces
    discovery_error = discovery_error.replace(/^\s+/,"");

    if (login_error.length != 0 || discovery_error.length != 0)
        {
        document.getElementById('error_img').style.display = 'block';
        document.getElementById('error_link').focus();
        }
    else
        setFocus();
    }

function submitAction(form)
{
signin(form);
form.Submit.disabled=true;
form.submit();
}
</script>
<BODY onLoad="setErrorImg(); ">
<table width="100%" height="99%"  border="0" cellpadding="0" cellspacing="0">
  <tr>
    <td height="96" align="center" valign="middle"><img
      src="/public/images/OPSE_logo.gif" alt="Oracle PeopleSoft logo" width="322" height="96"></td>
  </tr>
  <tr>
    <td height="250" align="center" valign="middle"><TABLE width="590" border=0 cellPadding=0 cellSpacing=0>
        <TR>
          <TD width="600" height="138" align=center valign="top">
                <table width="100%" height="273" border="0" cellpadding="0" cellspacing="0">
                  <tr>
                    <td><table width="100%" border="0" cellspacing="0" cellpadding="0">
                        <tr>
                          <td width="50%" height="120" valign="middle" class="psloginframe"><table width="100%" border="0" cellspacing="0" cellpadding="0">
                              <form action="?cmd=login&languageCd=ENG" method="post" id="login" name="login" autocomplete=off onSubmit="signin(document.login)">
                                <input type="hidden" name="timezoneOffset" value="0">
                              <table width="100%" border="0" cellspacing="0" cellpadding="0">
                              <tr>
                                <td colspan="3"><span class="psloginlabel"><img src="/public/images/shim.gif" width="1" height="8" alt=""></span></td>
                              </tr>
                              <tr>
                                <div style="text-align:center">
                                <h1 id="error_img" style="display:none"><a id ="error_link" href="javascript:setFocus();" tabindex="1"><img src="/public/images/PT_LOGIN_ERROR.gif" border="0" alt="Error"/></a></h1>
                                <h2 class="psloginerror" id="login_error"> java.lang.RuntimeException: Failover string csas5:9000,csas6:9000,csas7:9000,csas8:9000,csas9:9000,csas10:9000,csas11:9000:csas4:9000 has invalid format. </h2>
                                <h2 class="psloginerror" id="discovery_error">  </h2>
                                </div>
                              </tr>
                              <tr>
                                <td width="45%" height="25" align="right"><label for='userid' class='psloginlabel'>User ID:</label></td>
                                <td width="1%" height="25">&nbsp;</td>
                                <td width="54%" height="25"><input id="userid" name="userid" type="text" class="pslogineditbox" value="" size="15"></td>
                              </tr>
                              <tr>
                                <td height="35" align="right"><label for='pwd' class='psloginlabel'>Password:</label></td>
                                <td height="35">&nbsp;</td>
                                <td height="35"><input TYPE="password" id="pwd" name="pwd" class="pslogineditbox" size="15"></td>
                              </tr>
                              <tr>
                                <td height="35">&nbsp;</td>
                                <td height="35">&nbsp;</td>
                                <td height="35"><input name="Submit" type="submit" class="psloginbutton" value="Sign In" onclick="submitAction(document.login)"></td>
                              </tr>
                          </table></td>
                          <td width="50%" rowspan="2" align="center" class="pslanguageframe"><table width="90%" border="0" cellpadding="0" cellspacing="0">
                            <tr>
                              <td valign="top" class="psloginlabel"><img src="/public/images/shim.gif" width="1" height="8" alt=""></td>
                            </tr>
                            <tr>
                              <td height="25" valign="middle" class="psloginlabel">Select a Language:</td>
                            </tr>
                            <tr>
                              <td valign="top" class="psloginlabel"><table width="100%" border="0" cellspacing="0" cellpadding="0">
                                  <tr>
                                    <td width="50%" height="18" class="pslogintext"><a href="?cmd=login&amp;languageCd=ENG">English</a></td>
                                    <td width="50%" class="pslogintext"><a href="?cmd=login&amp;languageCd=ESP">Espa&ntilde;ol</a></td>
                                  </tr>
                                  <tr>
                                    <td width="50%" class="pslogintext"><a href="?cmd=login&amp;languageCd=DAN">Dansk</a><a href="?cmd=login&amp;languageCd=ESP"></a></td>
                                    <td width="50%" height="18" class="pslogintext"><a href="?cmd=login&amp;languageCd=GER">Deutsch</a><a href="?cmd=login&amp;languageCd=DAN"></a></td>
                                  </tr>
                                  <tr>
                                    <td width="50%" height="18" class="pslogintext"><a href="?cmd=login&amp;languageCd=FRA">Fran&ccedil;ais</a></td>
                                    <td width="50%" class="pslogintext"><a href="?cmd=login&amp;languageCd=CFR">Fran&ccedil;ais&nbsp;du&nbsp;Canada</a><a href="?cmd=login&amp;languageCd=GER"></a></td>
                                  </tr>
                                  <tr>
                                    <td width="50%" height="18" class="pslogintext"><a href="?cmd=login&amp;languageCd=ITA">Italiano</a><a href="?cmd=login&amp;languageCd=FRA"></a></td>
                                    <td width="50%" class="pslogintext"><a href="?cmd=login&amp;languageCd=HUN">Magyar</a></td>
                                  </tr>
                                  <tr>
                                    <td width="50%" height="18" class="pslogintext"><a href="?cmd=login&amp;languageCd=DUT">Nederlands</a><a href="?cmd=login&amp;languageCd=CFR"></a></td>
                                    <td width="50%" class="pslogintext"><a href="?cmd=login&amp;languageCd=NOR">Norsk</a><a href="?cmd=login&amp;languageCd=ITA"></a></td>
                                  </tr>
                                  <tr>
                                    <td width="50%" class="pslogintext"><a href="?cmd=login&amp;languageCd=POL">Polski</a></td>
                                    <td width="50%" height="18" class="pslogintext"><a href="?cmd=login&amp;languageCd=POR">Portugu&ecirc;s</a><a href="?cmd=login&amp;languageCd=HUN"></a></td>
                                  </tr>
                                  <tr>
                                    <td width="50%" height="18" class="pslogintext"><a href="?cmd=login&amp;languageCd=FIN">Suomi</a><a href="?cmd=login&amp;languageCd=DUT"></a></td>
                                    <td width="50%" height="18" class="pslogintext"><a href="?cmd=login&amp;languageCd=SVE">Svenska</a></td>
                                  </tr>
                                  <tr>
                                    <td width="50%" height="18" class="textnormal"><a href="?cmd=login&languageCd=CZE""><img src="/public/images/language_cze.gif" width="44" height="16" border="0" title="Czech" alt="Czech" /></a></td>
                                    <td width="50%" height="18" class="textnormal"><a href="?cmd=login&languageCd=JPN"><img src="/public/images/language_japanese.gif" width="37" height="16" border="0" title="Japanese" alt="Japanese" /></a></td>
                                  </tr>
                                  <tr>
                                    <td width="50%" height="18" class="textnormal"><a href="?cmd=login&languageCd=KOR"><img src="/public/images/korean.gif" width="34" height="16" border="0" title="Korean" alt="Korean" /></a></td>
                                    <td width="50%" height="18" class="textnormal"><a href="?cmd=login&languageCd=RUS"><img src="/public/images/language_rus.gif" width="50" height="16" border="0" title="Russian" alt="Russian" /></a></td>
                                  </tr>
                                  <tr>
                                    <td width="50%" height="18" class="textnormal"><a href="?cmd=login&languageCd=THA"><img src="/public/images/language_tha.gif" width="24" height="16" border="0" title="Thai" alt="Thai" /></a></td>
                                    <td width="50%" height="18" class="textnormal"><a href="?cmd=login&languageCd=ZHS"><img src="/public/images/schinese.gif" width="53" height="16" border="0" title="Simplified Chinese" alt="Simplified Chinese" /></a></td>
                                  </tr>
                                  <tr>
                                    <td width="50%" height="18" class="textnormal"><a href="?cmd=login&languageCd=ZHT"><img src="/public/images/tchinese.gif" width="53" height="16" border="0" title="Traditional Chinese" alt="Traditional Chinese" /></a></td>
                                    <td width="50%" height="18" class="textnormal"><a href="?cmd=login&languageCd=ARA"><img src="/public/images/arabic.gif" width="32" height="16" border="0" title="Arabic" alt="Arabic" /></a></td>
                                  </tr>
                                  <tr>
                                    <td colspan="2"><img src="/public/images/shim.gif" width="1" height="10" alt=""></td>
                                  </tr>
                              </table></td>
                            </tr>
                          </table></td>
                        </tr>
                        <tr>
                          <td height="119" valign="middle" class="psmessageframe"><table width="100%" border="0" cellspacing="0" cellpadding="5">
                            <tr>
                              <td><div align="center">
                                <p class="pslogintext">  </p>
                              </div></td>
                            </tr>
                          </table></td>
                        </tr>

                    </table></td>
                  </tr>
          </table></TD>
        </TR>
        <TR>
          <TD align=center>&nbsp;</TD>
        </TR>
      </TBODY>
    </TABLE></td>
  </tr>
  <tr>
    <td valign="bottom"><table width="100%"  border="0" cellspacing="0" cellpadding="5">
      <tr>
        <td height="30" colspan="2" valign="bottom" class="pslogincopyright">&nbsp;</td>
        </tr>
      <tr>
        <td width="50%" valign="bottom" class="pslogincopyright">Copyright &copy; 2000, 2007, Oracle. All rights reserved. PeopleSoft is a registered trademark of Oracle Corporation and/or its affiliates. Other names may be trademarks of their respective owners.</td>
        <td width="569">&nbsp;</td>
      </tr>
    </table></td>
  </tr>
</table>
</BODY></HTML>
`
