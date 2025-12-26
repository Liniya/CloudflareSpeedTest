:: --------------------------------------------------------------
:: Project: CloudflareSpeedTest Auto Update 3Proxy
:: Version: 1.0.6
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


:: If the nowip_3proxy.txt file does not exist, it means this is the first time running this script
if not exist "nowip_3proxy.txt" (
    echo This script's purpose is to obtain the fastest IP after CFST speed test and replace the Cloudflare CDN IP in the 3Proxy configuration file.
    echo It can redirect all Cloudflare CDN IPs to the fastest IP, achieving permanent acceleration for all websites using Cloudflare CDN (no need to add domains one by one to Hosts).
    echo Before use, please read: https://github.com/XIU2/CloudflareSpeedTest/discussions/71
    echo.
    set /p nowip="Enter the current Cloudflare CDN IP used by 3Proxy and press Enter (this step will not be needed again):"
    echo !nowip!>nowip_3proxy.txt
    echo.
)  

:: Get the currently used Cloudflare CDN IP from the nowip_3proxy.txt file
set /p nowip=<nowip_3proxy.txt
echo Starting speed test...


:: This RESET is for users who need the function "keep cycling speed tests until a suitable IP is found"
:: If you need this function, change the following 3 goto :STOP to goto :RESET
:RESET


:: Here you can add or modify CFST run parameters; echo.| is used to automatically press Enter to exit the program (no need to add -p 0 parameter anymore)
echo.|cfst.exe -o "result_3proxy.txt"


:: Check if the result file exists; if it does not exist, it means the result count is 0
if not exist result_3proxy.txt (
    echo.
    echo CFST speed test returned 0 IPs, skipping the following steps...
    goto :STOP
)

:: Get the fastest IP from the first line
for /f "skip=1 tokens=1 delims=," %%i in ('more result_3proxy.txt') do (
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
::for /f %%a in ('type result_3proxy.txt') do set /a v+=1
::if %v% GTR 11 (
::    echo.
::    echo CFST speed test did not find any IP fully meeting the conditions, restarting test...
::    goto :RESET
::)


echo %bestip%>nowip_3proxy.txt
echo.
echo Old IP: %nowip%
echo New IP: %bestip%



:: Please change "D:\Program Files\3Proxy" in quotes to your 3Proxy program directory
CD /d "D:\Program Files\3Proxy"
:: Make sure that before running this script, you have tested that 3Proxy can run normally and is functional!



echo.
echo Backing up 3proxy.cfg file (3proxy.cfg_backup)...
copy 3proxy.cfg 3proxy.cfg_backup
echo.
echo Starting replacement...
(
    for /f "tokens=*" %%i in (3proxy.cfg_backup) do (
        set s=%%i
        set s=!s:%nowip%=%bestip%!
        echo !s!
        )
)>3proxy.cfg

net stop 3proxy
net start 3proxy

echo Done...
echo.
:STOP
pause 
