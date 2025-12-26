# XIU2/CloudflareSpeedTest - Script

This directory contains scripts that call **CFST** and extend it to implement additional personalized features.

****
> [!TIP]
> I designed CFST as a command-line program specifically for its **generality**, because it's impossible to include every possible requirement directly in the software (especially niche or highly personalized needs). Doing so would increase maintenance complexity and burden, and lead to an overly bloated application (`"becoming the kind of thing I hate"`). One advantage of command-line programs is that they can be easily combined with other software and scripts.

For example, the scripts I've written here use CFST to perform speed tests and obtain results, then **freely process** those results according to your own needs (such as modifying Hosts files).

Overall, the scripts I've written are simple and serve single purposes. Besides meeting some users' needs, they are primarily meant as **reference examples** of how to combine CFST with scripts. Users familiar with scripting or programming can easily create their own personalized solutions.

Of course, if you have useful custom scripts, feel free to share them via [**Issues**](https://github.com/XIU2/CloudflareSpeedTest/issues), [**Discussions**](https://github.com/XIU2/CloudflareSpeedTest/discussions), or **Pull requests** so others can benefit too!

> Tip: Click the three horizontal lines icon in the top-right corner to view the table of contents.

****

## 📑 cfst_hosts.sh / cfst_hosts.bat (Included in the package)

This script runs CFST to find the fastest IP and replaces the old CDN IP in the Hosts file.

> **Author:** [@XIU2](https://github.com/xiu2)  
> **Usage instructions / Feedback: https://github.com/XIU2/CloudflareSpeedTest/discussions/312**

<details>
<summary><code><strong>「 Changelog」</strong></code></summary>

****

#### December 15, 2025, Version v1.0.5 (cfst_hosts.bat)
 - **1. Fixed** issue where the first IP line could not be retrieved under CFST's new version 

#### December 17, 2021, Version v1.0.4
 - **1. Optimized** functionality for "keep looping tests if no suitable IP is found" — fixed issue where re-testing didn't occur when a download speed limit was specified (commented out by default)   

#### December 17, 2021, Version v1.0.3
 - **1. Added** option to keep looping tests if no suitable IP is found (commented out by default)  
 - **2. Optimized** code  

#### September 29, 2021, Version v1.0.2
 - **1. Fixed** issue where the script did not exit when the number of speed test results was 0  

#### April 29, 2021, Version v1.0.1
 - **1. Optimized** no longer requires the -p 0 parameter to avoid exiting on Enter key press (now displays results without risk of accidental exit)  

#### January 28, 2021, Version v1.0.0
 - **1. Released** first version  

</details>

****

## 📑 cfst_3proxy.bat (Included in the package)

This script runs CFST to find the fastest IP and replaces the old Cloudflare CDN IP in the 3Proxy configuration file.  
It can redirect all Cloudflare CDN IPs to the fastest one, achieving permanent acceleration for all websites using Cloudflare CDN (no need to manually add each domain to Hosts).

> **Author:** [@XIU2](https://github.com/xiu2)  
> **Usage instructions / Feedback: https://github.com/XIU2/CloudflareSpeedTest/discussions/71**

<details>
<summary><code><strong>「 Changelog」</strong></code></summary>

****

#### December 15, 2025, Version v1.0.6
 - **1. Fixed** issue where the first IP line could not be retrieved under CFST's new version 

#### December 17, 2021, Version v1.0.5
 - **1. Optimized** functionality for "keep looping tests if no suitable IP is found" — fixed issue where re-testing didn't occur when a download speed limit was specified (commented out by default)   

#### December 17, 2021, Version v1.0.4
 - **1. Added** option to keep looping tests if no suitable IP is found (commented out by default)  
 - **2. Optimized** code  

#### September 29, 2021, Version v1.0.3
 - **1. Fixed** issue where the script did not exit when the number of speed test results was 0  

#### April 29, 2021, Version v1.0.2
 - **1. Optimized** no longer requires the -p 0 parameter to avoid exiting on Enter key press (now displays results without risk of accidental exit)  

#### March 16, 2021, Version v1.0.1
 - **1. Optimized** code and comments  

#### March 13, 2021, Version v1.0.0
 - **1. Released** first version  

</details>

****

## 📑 cfst_dnspod.sh

If your domain is hosted on **Dnspod**, you can use Dnspod's official API to automatically update DNS records!  
This script runs CFST to find the fastest IP and updates the domain's DNS record to that IP via the Dnspod API.

> **Author:** [@imashen](https://github.com/imashen)  
> **Usage instructions / Feedback: https://github.com/XIU2/CloudflareSpeedTest/pull/533**

<details>
<summary><code><strong>「 Changelog」</strong></code></summary>

****

#### August 6, 2024, Version v1.0.0
 - **1. Released** first version  

</details>

****

## 📑 cfst_ddns.sh / cfst_ddns.bat

If your domain is hosted on **Cloudflare**, you can use Cloudflare's official API to automatically update DNS records!  
This script runs CFST to find the fastest IP and updates the domain's DNS record to that IP via the Cloudflare API.

> **Author:** [@XIU2](https://github.com/xiu2)  
> **Usage instructions / Feedback: https://github.com/XIU2/CloudflareSpeedTest/discussions/481**

<details>
<summary><code><strong>「 Changelog」</strong></code></summary>

****

#### December 15, 2025, Version v1.0.6 (cfst_ddns.bat)
 - **1. Fixed** issue where the first IP line could not be retrieved under CFST's new version 

#### October 6, 2024, Version v1.0.5
 - **1. Added** support for API tokens (compared to global API keys, API tokens allow fine-grained permission control)   

#### December 17, 2021, Version v1.0.4
 - **1. Added** option to keep looping tests if no suitable IP is found (commented out by default)  
 - **2. Optimized** code  

#### September 29, 2021, Version v1.0.3
 - **1. Fixed** issue where the script did not exit when the number of speed test results was 0  

#### April 29, 2021, Version v1.0.2
 - **1. Optimized** no longer requires the -p 0 parameter to avoid exiting on Enter key press (now displays results without risk of accidental exit)  

#### January 27, 2021, Version v1.0.1
 - **1. Optimized** configuration to be read from a file  

#### January 26, 2021, Version v1.0.0
 - **1. Released** first version  

</details>

****

## 📑 cfst_dnsmasq.sh

This script runs CFST to find the fastest IP and replaces the old Cloudflare CDN IP in the dnsmasq configuration file.  

> **Author:** [@Sving1024](https://github.com/Sving1024)  
> **Usage instructions / Feedback: https://github.com/XIU2/CloudflareSpeedTest/discussions/566**

<details>
<summary><code><strong>「 Changelog」</strong></code></summary>

****

#### January 22, 2025, Version v1.0.1
 - **1. Fixed** IPv6 issue  

#### December 28, 2024, Version v1.0.0
 - **1. Released** first version  

</details>

****

## Feature Suggestions / Feedback

If you encounter any issues using these scripts, first check the corresponding **"Usage Instructions"** thread to see if others have asked similar questions.  
If you don't find a similar issue, please comment directly in the corresponding **"Usage Instructions"** thread to ask the author.
