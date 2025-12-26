:: --------------------------------------------------------------
:: Project: CloudflareSpeedTest Auto Update Hosts
:: Version: 1.0.5
:: Author: XIU2
:: Project: https://github.com/XIU2/CloudflareSpeedTest
:: --------------------------------------------------------------
@echo off
Setlocal Enabledelayedexpansion

:: Check if administrator privileges have been obtained

>nul 2>&1 "%SYSTEMROOT%\system32\cacls.exe" "%SYSTEMROOT%\system32\config\system" 

if '%errorlevel%' NEQ '0' (  
    goto UACPrompt  
) else ( goto gotAdmin )  

:: Write a vbs script to run this batch file with administrator privileges

:UACPrompt  
    echo Set UAC = CreateObject^("Shell.Application"^) > "%temp%\getadmin.vbs" 
    echo UAC.ShellExecute "%~s0", "", "", "runas", 1 >> "%temp%\getadmin.vbs" 
    "%temp%\getadmin.vbs" 
    exit /B  

:: If the temporary vbs script exists, delete it
  
:gotAdmin  
    if exist "%temp%\getadmin.vbs" ( del "%temp%\getadmin.vbs" )  
    pushd "%CD%" 
    CD /D "%~dp0" 


:: The above section checks if administrator privileges have been obtained; if not, obtain them. Below is the main code of this script.


:: If the nowip_hosts.txt file does not exist, it means this is the first time running this script
if not exist "nowip_hosts.txt" (
    echo This script's purpose is to obtain the fastest IP after CFST speed test and replace the Cloudflare CDN IP in the Hosts file.
    echo Before use, please read: https://github.com/XIU2/CloudflareSpeedTest/issues/42#issuecomment-768273768
    echo.
    echo For first-time use, please change all Cloudflare CDN IPs in the Hosts file to a single IP.
    set /p nowip="Enter that Cloudflare CDN IP and press Enter (this step will not be needed again):"
    echo !nowip!>nowip_hosts.txt
    echo.
)  

:: Get the currently used Cloudflare CDN IP from the nowip_hosts.txt file
set /p nowip=<nowip_hosts.txt
echo Starting speed test...


:: This RESET is for users who need the function "keep cycling speed tests until a suitable IP is found"
:: If you need this function, change the following 3 goto :STOP to goto :RESET
:RESET


:: Here you can add or modify CFST run parameters; echo.| is used to automatically press Enter to exit the program (no need to add -p 0 parameter anymore)
echo.|cfst.exe -o "result_hosts.txt"


:: Check if the result file exists; if it does not exist, it means the result count is 0
if not exist result_hosts.txt (
    echo.
    echo CFST speed test returned 0 IPs, skipping the following steps...
    goto :STOP
)

:: Get the fastest IP from the first line
for /f "skip=1 tokens=1 delims=," %%i in ('more result_hosts.txt') do (
    SET bestip=%%i
    goto :END
)

:END

:: Check if the obtained fastest IP is empty, and whether it is the same as the old IP
if "%bestip%"=="" (
    echo.
    echo CFST speed test returned 0 IPs, skipping the following steps...
    goto :STOP
)
if "%bestip%"=="%nowip%" (
    echo.
    echo CFST speed test returned 0 IPs, skipping the following steps...
    goto :STOP
)


:: The following code is only needed for the function "keep cycling speed tests until a suitable IP is found"
:: Considering that when a download speed lower limit is specified but no IP meets all conditions, CFST will output all IP results
:: Therefore, when using the -sl parameter, remove the leading :: comment marker below to perform line count checking (e.g., if 10 download test IPs are specified, set the value below to 11)
::set /a v=0
::for /f %%a in ('type result_hosts.txt') do set /a v+=1
::if %v% GTR 11 (
::    echo.
::    echo CFST speed test did not find any IP fully meeting the conditions, restarting test...
::    goto :RESET
::)


echo %bestip%>nowip_hosts.txt
echo.
echo Old IP: %nowip%
echo New IP: %bestip%

CD /d "C:\Windows\System32\drivers\etc"
echo.
echo Backing up Hosts file (hosts_backup)...
copy hosts hosts_backup
echo.
echo Starting replacement...
(
    for /f "tokens=*" %%i in (hosts_backup) do (
        set s=%%i
        set s=!s:%nowip%=%bestip%!
        echo !s!
        )
)>hosts

echo Done...
echo.
:STOP
pause 
