:: --------------------------------------------------------------
:: Project: CloudflareSpeedTest Auto Update DNS Record
:: Version: 1.0.6
:: Author: XIU2
:: Project: https://github.com/XIU2/CloudflareSpeedTest
:: --------------------------------------------------------------
@echo off
Setlocal Enabledelayedexpansion

:: Here you can add or modify CFST run parameters; echo.| is used to automatically press Enter to exit the program (no need to add -p 0 parameter anymore)
echo.|cfst.exe -o "result_ddns.txt"

:: Check if the result file exists; if it does not exist, it means the result count is 0
if not exist result_ddns.txt (
    echo.
    echo CFST speed test returned 0 IPs, skipping the following steps...
    goto :END
)

for /f "skip=1 tokens=1 delims=," %%i in (result_ddns.txt) do (
    Echo %%i
    if "%%i"=="" (
        echo.
        echo CFST speed test returned 0 IPs, skipping the following steps...
        goto :END
    )
:: API Key method (global permissions)
    curl -X PUT "https://api.cloudflare.com/client/v4/zones/DomainID/dns_records/DNSRecordID" ^
            -H "X-Auth-Email: AccountEmail" ^
            -H "X-Auth-Key: APIKeyObtainedEarlier" ^
            -H "Content-Type: application/json" ^
            --data "{\"type\":\"A\",\"name\":\"FullDomainName\",\"content\":\"%%i\",\"ttl\":1,\"proxied\":true}"
:: API Token method (custom permissions); if you want to use this method, delete or comment out the above lines and remove the leading "::" comment marker from the lines below.
::    curl -X PUT "https://api.cloudflare.com/client/v4/zones/DomainID/dns_records/DNSRecordID" ^
::            -H "Authorization: Bearer APIKeyObtainedEarlier" ^
::            -H "Content-Type: application/json" ^
::            --data "{\"type\":\"A\",\"name\":\"FullDomainName\",\"content\":\"%%i\",\"ttl\":1,\"proxied\":true}"

        goto :END
)
:END
pause
