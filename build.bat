set CVSROOT=:pserver:builduser@cvsserver:2401/cvs
set ANT_OPTS=-Xms512M -Xmx512M

SET GIT_PATH=C:\Program Files\Git\cmd
if exist "%GIT_PATH%" goto path_git_found
SET GIT_PATH=D:\Program Files\Git\cmd
:path_git_found

SET SIGNTOOL_PATH=C:\Program Files\Microsoft SDKs\Windows\v6.1\Bin
if exist "%SIGNTOOL_PATH%" goto path_signtool_found
SET SIGNTOOL_PATH=C:\Program Files\Microsoft SDKs\Windows\v7.0A\bin
:path_signtool_found

SET PATH=%SIGNTOOL_PATH%;%PATH%;%GIT_PATH%

set PATH_INNO=c:\Program Files\Inno Setup 5
if exist "%PATH_INNO%" goto path_inno_found
set PATH_INNO=E:\Program Files\Inno Setup 5
:path_inno_found

cvs co thirdparty_packages\code_signing_certificate

gradlew goBuild

"%SIGNTOOL_PATH%/signtool.exe" sign /f thirdparty_packages\code_signing_certificate\comodo_code_signing.pfx /p raymedi123 /t http://timestamp.comodoca.com/authenticode Output\*.exe

"%PATH_INNO%/compil32.exe" /cc peripheral.iss

"%SIGNTOOL_PATH%/signtool.exe" sign /f thirdparty_packages\code_signing_certificate\comodo_code_signing.pfx /p raymedi123 /t http://timestamp.comodoca.com/authenticode Output\ServQuick_Peripherals.exe

gradlew uploadArtifacts -PuploadType=Daily