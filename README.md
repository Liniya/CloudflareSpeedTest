# XIU2/CloudflareSpeedTest

> [!NOTE]
> This is an up-to-date fork (for 28.12.2025) of [**XIU2/CloudflareSpeedTest**](https://github.com/XIU2/CloudflareSpeedTest) that attempts to fix timeout support in the download test code for a specific blocking method called "[TCP 16-20](https://github.com/net4people/bbs/issues/490)" used by censors in certain countries. Additionally, all the Chinese text is machine-translated to English. Please support the original author(s) [here](https://github.com/XIU2/CloudflareSpeedTest).

[![Go Version](https://img.shields.io/github/go-mod/go-version/XIU2/CloudflareSpeedTest.svg?style=flat-square&label=Go&color=00ADD8&logo=go)](https://github.com/XIU2/CloudflareSpeedTest/)
[![Release Version](https://img.shields.io/github/v/release/XIU2/CloudflareSpeedTest.svg?style=flat-square&label=Release&color=00ADD8&logo=github)](https://github.com/XIU2/CloudflareSpeedTest/releases/latest)
[![GitHub license](https://img.shields.io/github/license/XIU2/CloudflareSpeedTest.svg?style=flat-square&label=License&color=00ADD8&logo=github)](https://github.com/XIU2/CloudflareSpeedTest/)
[![GitHub Star](https://img.shields.io/github/stars/XIU2/CloudflareSpeedTest.svg?style=flat-square&label=Star&color=00ADD8&logo=github)](https://github.com/XIU2/CloudflareSpeedTest/)
[![GitHub Fork](https://img.shields.io/github/forks/XIU2/CloudflareSpeedTest.svg?style=flat-square&label=Fork&color=00ADD8&logo=github)](https://github.com/XIU2/CloudflareSpeedTest/)

> [!NOTE]
> This is an up-to-date fork (for 28.12.2025) of [**XIU2/CloudflareSpeedTest**](https://github.com/XIU2/CloudflareSpeedTest) that attempts to fix timeout support in the download test code for a specific blocking method called "[TCP 16-20](https://github.com/net4people/bbs/issues/490)" used by censors in certain countries. Additionally, all the Chinese text is machine-translated to English. Please support the original author(s) [here](https://github.com/XIU2/CloudflareSpeedTest).

Many websites abroad use Cloudflare CDN, but the IP addresses allocated to visitors from mainland China are often unfriendly (high latency, high packet loss, slow speed).  
Although Cloudflare has published all its [IP ranges](https://www.cloudflare.com/zh-cn/ips/), finding the best ones among so many IPs could be exhausting—hence this tool.

**"Select Optimal IPs"** to test Cloudflare CDN latency and speed, and find the fastest IP (IPv4+IPv6)! If you find it useful, please give it a ⭐ to encourage me!

> _Check out my other open-source projects: [**TrackersList.com** - The most popular BT Tracker list! Improve your BT download speed~](https://github.com/XIU2/TrackersListCollection) <img src="https://img.shields.io/github/stars/XIU2/TrackersListCollection.svg?style=flat-square&label=Star&color=4285dd&logo=github" height="16px" />_  
> _[**UserScript** - 🐵 GitHub fast download, enhanced Zhihu, auto seamless pagination, eye-protection mode, and more than a dozen **Tampermonkey scripts**~](https://github.com/XIU2/UserScript) <img src="https://img.shields.io/github/stars/XIU2/UserScript.svg?style=flat-square&label=Star&color=4285dd&logo=github" height="16px" />_  
> _[**SNIProxy** - 🧷 My simple SNI Proxy (supports all platforms, full system, front proxy, easy configuration, etc.~)](https://github.com/XIU2/SNIProxy) <img src="https://img.shields.io/github/stars/XIU2/SNIProxy.svg?style=flat-square&label=Star&color=4285dd&logo=github" height="16px" />_  

Of course, this project also supports latency testing for **other CDNs or websites with multiple resolved IPs**, but you'll need to find your own download test URLs accordingly.

> [!IMPORTANT]
> Cloudflare CDN explicitly prohibits proxy usage. For any use involving proxying through CDN, you assume all risks. Do not rely excessively on it. See [#382](https://github.com/XIU2/CloudflareSpeedTest/discussions/382) [#383](https://github.com/XIU2/CloudflareSpeedTest/discussions/383)

****
## \# Quick Start

### Download and Run

1. Download the compiled executable file ( [Github Releases](https://github.com/XIU2/CloudflareSpeedTest/releases) / [Lanzou Cloud](https://xiu.lanzoub.com/b0742hkxe) ) and extract it.  
2. Double-click to run the `cfst.exe` file (Windows system), and wait for the speed test to complete...

<details>
<summary><code><strong>「 Click to view other installation methods for Windows 」</strong></code></summary>

****

If you have scoop (a command-line package manager for Windows), you can install it like this:

```sh
# Add the most popular Chinese software repository: dorado
scoop bucket add dorado https://github.com/chawyehsu/dorado
# Install cloudflare-speedtest
scoop install dorado/cloudflare-speedtest
```

</details>

<details>
<summary><code><strong>「 Click to view usage examples for Linux 」</strong></code></summary>

****

The following commands are only examples; please check the [**Releases**](https://github.com/XIU2/CloudflareSpeedTest/releases) page for the correct version number and filename.

``` yaml
# If using for the first time, it's recommended to create a new folder (skip this step for future updates)
mkdir cfst

# Enter the folder (for future updates, just repeat the download and extract steps below)
cd cfst

# Download the CFST archive (replace [version] and [filename] in the URL as needed)
wget -N https://github.com/XIU2/CloudflareSpeedTest/releases/download/v2.3.4/cfst_linux_amd64.tar.gz
# If you're downloading in a Chinese network environment, use one of these mirror accelerators:
# wget -N https://ghfast.top/https://github.com/XIU2/CloudflareSpeedTest/releases/download/v2.3.4/cfst_linux_arm64.tar.gz
# wget -N https://wget.la/https://github.com/XIU2/CloudflareSpeedTest/releases/download/v2.3.4/cfst_linux_arm64.tar.gz
# wget -N https://ghproxy.net/https://github.com/XIU2/CloudflareSpeedTest/releases/download/v2.3.4/cfst_linux_arm64.tar.gz
# wget -N https://gh-proxy.com/https://github.com/XIU2/CloudflareSpeedTest/releases/download/v2.3.4/cfst_linux_arm64.tar.gz
# wget -N https://hk.gh-proxy.com/https://github.com/XIU2/CloudflareSpeedTest/releases/download/v2.3.4/cfst_linux_arm64.tar.gz
# If download fails, try removing the -N parameter (if updating, remember to delete the old archive first: rm cfst_linux_amd64.tar.gz )

# Extract (no need to delete old files; they will be overwritten directly; replace filename as needed)
tar -zxf cfst_linux_amd64.tar.gz

# Grant execute permission
chmod +x cfst

# Run (without parameters)
./cfst

# Run (with example parameters)
./cfst -tl 200 -dn 20
```

> If the average latency is extremely low (e.g., 0.xx), it means CFST is using a proxy during speed testing. Please disable your proxy software before testing.  
> If running on a router, it's recommended to disable the router's proxy (or exclude CFST from it); otherwise, test results may be inaccurate/unusable.

</details>

****

> _Simple tutorial for running CFST independently on mobile phones: **[Android](https://github.com/XIU2/CloudflareSpeedTest/discussions/61)、[Android APP](https://github.com/xianshenglu/cloudflare-ip-tester-app)、[IOS](https://github.com/XIU2/CloudflareSpeedTest/discussions/321)**_

> [!NOTE]
> Note! This software is only suitable for websites. **It does NOT support selecting IPs for Cloudflare WARP, which uses UDP protocol**. See: [#392](https://github.com/XIU2/CloudflareSpeedTest/discussions/392)

****
## \# Quick Usage

After testing, the default output shows the **top 10 fastest IPs**. Example (output content only):

``` bash
IP Address        Sent   Received  Loss Rate  Avg Latency  Download Speed(MB/s)  Region Code
104.27.200.69     4      4         0.00       146.23       28.64                 LAX
172.67.60.78      4      4         0.00       139.82       15.02                 SEA
104.25.140.153    4      4         0.00       146.49       14.90                 SJC
104.27.192.65     4      4         0.00       140.28       14.07                 LAX
172.67.62.214     4      4         0.00       139.29       12.71                 LAX
104.27.207.5      4      4         0.00       145.92       11.95                 LAX
172.67.54.193     4      4         0.00       146.71       11.55                 LAX
104.22.66.8       4      4         0.00       147.42       11.11                 SEA
104.27.197.63     4      4         0.00       131.29       10.26                 FRA
172.67.58.91      4      4         0.00       140.19       9.14                  SJC
...

# If average latency is extremely low (e.g., 0.xx), it means CFST is using a proxy during speed testing. Disable your proxy software before testing.
# If running on a router, disable the router's proxy (or exclude CFST from it); otherwise, test results may be inaccurate/unusable.

# Since each test randomly selects an IP within each IP range, results will always differ between runs—this is normal!

# Note: I found that the first latency test after booting the computer shows significantly higher latency (same with manual TCPing), but subsequent tests are normal.
# Therefore, it's recommended that after booting, before your first official test, run a quick test on a few IPs (no need to wait for completion; just wait until the progress bar moves, then quit).

# The default process of this software:
# 1. Latency test (default TCPing mode; use HTTPing mode by adding parameter)
# 2. Sort by latency (ascending order, filtered by conditions; IPs with different packet loss rates are sorted separately, so some low-latency but high-loss IPs may appear later)
# 3. Download speed test (starts from the lowest-latency IP and proceeds sequentially; stops after testing 10 IPs by default)
# 4. Sort by download speed (descending order)
# 5. Output results (controlled by parameters whether to output to console (-p 0) or file (-o ""))

# Note: The output file result.csv may show Chinese garbled when opened with Microsoft Excel—this is normal; other spreadsheet software or Notepad display correctly.
```

The first row of the results is the **fastest IP**, combining the lowest latency and highest download speed!

The full results are saved in `result.csv` in the current directory. Open it with **Notepad/spreadsheet software**; format:

```
IP Address,Sent,Received,Loss Rate,Avg Latency,Download Speed(MB/s),Region Code
104.27.200.69,4,4,0.00,146.23,28.64,LAX
```

> [!NOTE]
> _If you find **download speed is 0.00**, use **debug mode `-debug`** to troubleshoot; see: [**# Download Speed is 0.00?**](https://github.com/Liniya/CloudflareSpeedTest#-download-speed-is-000)_

> _You can further filter and process the full results based on your needs, or check out advanced usage with **custom filtering conditions**!_

****
## \# Advanced Usage

Running without parameters uses default settings. To get more comprehensive or customized results, you can customize parameters.

```Dart
C:\>cfst.exe -h

CloudflareSpeedTest vX.X.X
Test latency and speed of all IPs for CDN or websites, find the fastest IP (IPv4+IPv6)!
https://github.com/XIU2/CloudflareSpeedTest

Parameters:
    -n 200
        Latency test threads; more threads mean faster latency testing, but don't set too high on low-performance devices (e.g., routers); (default 200, max 1000)
    -t 4
        Latency test count per IP; number of times to ping each IP; (default 4)
    -dn 10
        Number of IPs to download test; after latency testing and sorting, how many lowest-latency IPs to test for download speed; (default 10)
    -dt 10
        Download test duration per IP; minimum time allowed for each download test; cannot be too short; (default 10 seconds)
    -tp 443
        Specify test port; port used for latency and download tests; (default port 443)
    -url https://cf.xiu2.xyz/url
        Specify test URL; address used for latency (HTTPing) and download tests; default URL may not be reliable, it's recommended to build your own;
        During download testing, the software retrieves the region code of the IP from the HTTP response header (supports Cloudflare, AWS CloudFront, Fastly, Gcore, CDN77, Bunny, etc.)

    -httping
        Switch test mode; change latency test to HTTP protocol; test address is specified by [-url] parameter; (default TCPing)
        When using HTTP testing mode, the software retrieves the region code of the IP from the HTTP response header (supports Cloudflare, AWS CloudFront, Fastly, Gcore, CDN77, Bunny, etc.) and displays it.
        Note: HTTPing is essentially a form of network scanning. If you're running this on a server, reduce concurrency (-n), or you may be suspended by strict providers.
        If you notice that the first HTTPing test shows normal available IPs, but subsequent tests show fewer or even zero, then recover after a while, it may be due to your ISP or Cloudflare CDN detecting network scanning and triggering temporary restrictions. Lowering concurrency (-n) can reduce this.
    -httping-code 200
        Valid status code; HTTP latency test considers only this HTTP status code as valid; only one code allowed; (default 200 301 302)
    -cfcolo HKG,KHH,NRT,LAX,SEA,SJC,FRA,MAD
        Match specific regions; IATA airport region codes or country/city codes, separated by commas, case-insensitive; only available in HTTPing mode; (default all regions)
        Supports Cloudflare, AWS CloudFront, Fastly, Gcore, CDN77, Bunny, etc.
        Cloudflare, AWS CloudFront, Fastly use IATA 3-letter airport codes, e.g.: HKG,LAX
        CDN77, Bunny use 2-letter country/region codes, e.g.: US,CN
        Gcore uses 2-letter city codes, e.g.: FR,AM
        So when using -cfcolo to specify region codes, choose the correct type based on the CDN.

    -tl 200
        Max average latency; only output IPs with average latency below this value; can be combined with other limits; (default 9999 ms)
    -tll 40
        Min average latency; only output IPs with average latency above this value; (default 0 ms)
    -tlr 0.2
        Max packet loss rate; only output IPs with packet loss rate <= this value; range 0.00~1.00; 0 filters out any IP with packet loss; (default 1.00)
    -sl 5
        Min download speed; only output IPs with download speed above this value; stops testing once -dn number is reached; (default 0.00 MB/s)

    -p 10
        Number of results to display; show specified number of results after test; 0 means no output and exit immediately; (default 10)
    -f ip.txt
        IP range data file; if path contains spaces, use quotes; supports other CDN IP ranges; (default ip.txt)
    -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
        Specify IP ranges directly via parameter; comma-separated list of IPs; (default empty)
    -o result.csv
        Write results to file; if path contains spaces, use quotes; empty value means no file output [-o ""]; (default result.csv)
        Note: In some environments, using -o "" may be ignored as an empty parameter causing errors; use -o " " (space) to fix.

    -dd
        Disable download test; after latency testing, sort results by latency instead of download speed; (default enabled)
    -allip
        Test all IPs; test every IP in the IP range (IPv4 only); (default: randomly test one IP per /24 segment)

    -debug
        Debug output mode; outputs more logs during unexpected events to help diagnose issues; (default off)
        Currently, this feature only logs errors during HTTPing latency testing and download testing. If any test fails due to various reasons, the error reason will be logged.
        E.g., HTTPing latency test fails due to invalid HTTP status code, invalid test URL, or timeout.
        E.g., Download test fails due to invalid URL (blocked, 403 status code, timeout), resulting in 0.00 display.

    -v
        Print program version + check for updates
    -h
        Print help
```

### Interface Explanation

To avoid confusion about output during testing (e.g., "available", "queue" numbers, download test "interrupted" halfway, or "stuck"), here's an explanation.

<details>
<summary><code><strong>「 Click to expand content 」</strong></code></summary>

****

> This example uses common parameters: `-tll 40 -tl 150 -sl 1 -dn 5`, outputting:

```python
# XIU2/CloudflareSpeedTest vX.X.X

Starting latency test (mode: TCP, port: 443, range: 40 ~ 150 ms, packet loss: 1.00)
321 / 321 [-----------------------------------------------------------] Available: 30
Starting download test (min speed: 1.00 MB/s, count: 5, queue: 10)
3 / 5 [-----------------------------------------↗--------------------]
IP Address        Sent   Received  Loss Rate  Avg Latency  Download Speed(MB/s)  Region Code
XXX.XXX.XXX.XXX   4      4         0.00       83.32        3.66                  LAX
XXX.XXX.XXX.XXX   4      4         0.00       107.81       2.49                  LAX
XXX.XXX.XXX.XXX   4      3         0.25       149.59       1.04                  N/A

Full test results written to result.csv. View with Notepad/spreadsheet software.
Press Enter or Ctrl+C to exit.
```

****

> New users may wonder: "Why are there 30 available IPs after latency testing, but only 3 in the final output?"  
> What does "queue" mean in download test? Do I need to queue?

CFST first performs latency testing. During this phase, the progress bar shows the number of available IPs (e.g., `Available: 30`). Note: "available" means IPs that did not time out during latency testing, unrelated to latency/loss limits. After latency testing completes, due to specified latency limits and packet loss conditions, only `10` IPs remain eligible for download testing (i.e., the `queue: 10`).

In the above example: 321 IPs were tested; 30 passed without timeout. Then, filtered by latency range (40~150 ms) and packet loss limit, only 10 remained. If you used `-dd` to disable download testing, these 10 would be output directly. But since download testing is enabled, CFST proceeds to test these 10 IPs for download speed (`queue: 10`).

> Since download testing is single-threaded and tests one IP at a time, the number of waiting IPs is called the "queue".

****

> You may notice: "I specified `-dn 5` to find 5 IPs with good download speed, but why did it stop at 3?"

The download test progress bar `3 / 5` means: found `3` IPs meeting the download speed minimum condition (i.e., >1 MB/s), and `5` is your target number (`-dn 5`).

> Also note: if the available IP count after latency testing is less than `-dn`, then the progress bar's total will match the available count, not `-dn`.

After testing all 10 IPs, only 3 met the download speed condition (>1 MB/s); the other 7 did not.

So this isn't "interrupted before reaching 5"—it means all IPs were tested, but only 3 met your criteria.

****

Another scenario: when many IPs are available (hundreds or thousands), and you set a download speed limit, you may see: "Why is the download progress bar stuck at `X / 5`?"

This isn't stuck—it's because the progress bar only increments when an IP meets your download speed condition. So if none meet it, CFST keeps testing, making the progress bar appear frozen. This is a hint: your download speed expectation is too high for your network. Lower `-sl`.

****

If you find no IPs meet your conditions after testing all, lower your download speed limit (`-sl`) or remove it.  
Specify `-dn 20` without `-sl` to test only the top 20 lowest-latency IPs and stop—saving time.

****

If all queue IPs are tested but none meet the download speed condition, you may need to lower your expectation. But first, know the actual speed range. Use `-debug` mode to ignore conditions and output all results—you'll see actual speeds and can adjust `-sl` accordingly.

> Note: If you **don't specify** a download speed limit (`-sl`), CFST will always output all test results.

Similarly, for latency testing, `Available: 30` and `Queue: 10` tell you whether your latency/loss conditions are too strict. If many IPs are available but few remain after filtering, adjust your limits.

These mechanisms help you know if your latency/loss conditions are appropriate, and if your download speed condition is realistic.

</details>

****

### Usage Examples

On Windows, use CMD to pass parameters, or add them to a shortcut target.

> [!TIP]
> - All parameters have **default values**; omit them when using defaults (**choose as needed**); parameters are **order-independent**.  
> - In **PowerShell**, replace `cfst.exe` with `.\cfst.exe`.  
> - On Linux/macOS, replace `cfst.exe` with `./cfst`.

****

#### \# Run with Parameters in CMD

If unfamiliar with command-line programs, here's how to run with parameters.

<details>
<summary><code><strong>「 Click to expand content 」</strong></code></summary>

****

Many users open CMD and run CFST with absolute path, getting errors because `-f ip.txt` uses a relative path. You'd need to specify the absolute path to ip.txt, which is inconvenient. Instead, enter the CFST directory and use relative paths:

**Method One**:
1. Open the CFST program directory  
2. Hold <kbd>Shift + Right-click</kbd> to show context menu  
3. Select **[Open Command Window Here]** to open CMD in current directory  
4. Enter command with parameters, e.g.: `cfst.exe -tl 200 -dn 20`

**Method Two**:
1. Open the CFST program directory  
2. Click the address bar, select all, type `cmd`, and press Enter to open CMD in current directory  
4. Enter command with parameters, e.g.: `cfst.exe -tl 200 -dn 20`

> You can also open any CMD window and type `cd /d "D:\Program Files\cfst"` to enter the program directory.

> **Tip**: In **PowerShell**, replace `cfst.exe` with `.\cfst.exe`.  
> **Note**: In **PowerShell**, `-o ""` may be ignored as an empty parameter causing errors; use `-o " "` (space) to fix.

</details>

****

#### \# Run with Parameters via Windows Shortcut

If you rarely change parameters (e.g., always double-click), use a shortcut for convenience.

<details>
<summary><code><strong>「 Click to expand content 」</strong></code></summary>

****

Right-click `cfst.exe` → **[Create Shortcut]**, then right-click the shortcut → **[Properties]**, modify the **Target**:

``` bash
# To avoid outputting result file, add -o " " (space inside quotes; omitting space causes empty parameter to be ignored and error)
D:\ABC\cfst\cfst.exe -tl 200 -dn 20 -o " "

# If path contains spaces, put parameters outside quotes; ensure space between quote and dash.
"D:\Program Files\cfst\cfst.exe" -tl 200 -dn 20 -o " "

# Note: Shortcut "Start in" cannot be empty; otherwise absolute path will fail to find ip.txt
```

</details>

****

#### \# IPv4/IPv6

<details>
<summary><code><strong>「 Click to expand content 」</strong></code></summary>
****
``` bash
# Specify built-in IPv4 data file to test these IPv4 addresses (-f default is ip.txt, so parameter can be omitted)
cfst.exe -f ip.txt

# Specify built-in IPv6 data file to test these IPv6 addresses
# Since v2.1.0, IPv4+IPv6 mixed testing is supported and -ipv6 parameter removed; one file can contain both IPv4 and IPv6 addresses
cfst.exe -f ipv6.txt

# Or specify IPs directly via parameter
cfst.exe -ip 1.1.1.1,2606:4700::/32
```

> When testing IPv6, you may notice varying test counts each time; see reason: [#120](https://github.com/XIU2/CloudflareSpeedTest/issues/120)  
> Since IPv6 is vast (billions), and most IP ranges are unused, I only included some usable IPv6 ranges in `ipv6.txt`. You can scan and add/delete as desired; ASN data source: [bgp.he.net](https://bgp.he.net/AS13335#_prefixes6)

</details>

****

#### \# HTTPing

<details>
<summary><code><strong>「 Click to expand content 」</strong></code></summary>

****

Two latency test modes: **TCP protocol** and **HTTP protocol**.  
TCP is faster and less resource-intensive, timeout: 1 second (default).  
HTTP is suitable for quickly testing if a domain resolves to an IP and is accessible, timeout: 2 seconds.  
For the same IP, latencies typically follow: **ICMP < TCP < HTTP**; the latter is more sensitive to packet loss/network fluctuations.

> Note: HTTPing is essentially a form of **network scanning**. If running on a server, **reduce concurrency** (`-n`), or you may be suspended by strict providers. If HTTPing shows normal available IPs initially but fewer or zero later, then recovers after a while, your ISP or Cloudflare CDN may have detected network scanning and triggered temporary restrictions—lowering concurrency (`-n`) reduces this.

> Also, CFST only retrieves **response headers** during HTTPing (not page content), so file size doesn't affect HTTPing test (but download testing still needs a large file). Similar to `curl -i`.

> During HTTPing, the software retrieves the region code from HTTP response headers (supports Cloudflare, AWS CloudFront, Fastly, Gcore, CDN77, Bunny, etc.) and displays it. TCPing cannot do this (but download testing can, since it's also an HTTP request).

``` bash
# Just add -httping to switch to HTTP protocol latency test mode
cfst.exe -httping

# Software determines availability based on HTTP status code returned (timeout also counts); default considers 200, 301, 302 as valid. You can manually specify one valid HTTP status code (you must know what status code the test URL normally returns)
cfst.exe -httping -httping-code 200

# Use -url to specify HTTPing test address (can be any webpage URL, not limited to specific file)
cfst.exe -httping -url https://cf.xiu2.xyz/url
# If testing other websites/CDNs, specify an address using that website/CDN (default URL is for Cloudflare only)

# Note: If test URL uses HTTP protocol, add -tp 80 (this parameter affects port used in latency/download tests)
# Similarly, to test port 80, you must specify a http:// URL with -url (and it won't force redirect to HTTPS); for non-80/443 ports, ensure the download test URL supports access via that port.
cfst.exe -httping -tp 80 -url http://cdn.cloudflare.steamstatic.com/steam/apps/5952/movie_max.webm
```

</details>

****

#### \# Match Specific Regions

<details>
<summary><code><strong>「 Click to expand content 」</strong></code></summary>

****

Cloudflare CDN nodes use Anycast IPs—each IP's server location isn't fixed but dynamically changes based on **region, ISP, and time**. The same IP may route to different servers for different users (e.g., a US user gets a US node, a Japanese user gets a Japan node; mainland China is special—it gets routed to other countries). Different IP ranges have different routing/logic.

> Note! Although Cloudflare has many Asian nodes, **you may not get them**. Singaporeans might find many Singapore nodes easily, but you might find none—even after scanning all—because CDN controls routing. Anycast IP routing changes frequently; today an IP might be US, tomorrow it could be Europe (this is just an example; changes aren't that frequent; affected by congestion, cost, etc.). So **don't expect too much** from this feature!

Or pick any Cloudflare CDN IP (e.g., `104.16.123.96`), then use global ping test sites like [ping.sx/ping?t=104.16.123.96]—you'll find this IP has single-digit latency globally, often 0.X ms; only mainland China shows high latency (hundreds of ms).

This is Anycast technology—only mainland China's special network conditions require selecting optimal CDN IPs.

Besides retrieving region codes via **HTTP response headers**, you can manually visit `http://CloudflareIP/cdn-cgi/trace` to see the actual node region code assigned to you.

> This feature supports **Cloudflare, AWS CloudFront, Fastly, Gcore, CDN77, Bunny**.  
> But not all CDNs support Anycast; many restrict IP ranges per website.

> **Cloudflare, AWS CloudFront, Fastly** use **IATA 3-letter airport codes**, e.g.: HKG,LAX  
> **CDN77, Bunny** use **2-letter country/region codes**, e.g.: US,CN  
> **Gcore** uses **2-letter city codes**, e.g.: FR,AM  
> So when using `-cfcolo`, specify the correct code type for your CDN.

> Note: To filter AWS CloudFront regions, use `-url` to specify a download test URL using AWS CloudFront CDN (since default is Cloudflare). Sometimes HTTPing AWS CloudFront URLs returns 403; add `-httping-code 403` to correctly retrieve region code.

``` bash
# After specifying region names, latency test results will only include IPs from those regions (if -dd not used, download testing continues)
# If latency test returns 0, no available IPs from specified regions were found.
# Node region names are IATA airport codes or country/city codes; specify multiple with commas, lowercase supported since v2.2.3

cfst.exe -httping -cfcolo HKG,KHH,NRT,LAX,SEA,SJC,FRA,MAD

# Note: This parameter only works in HTTPing latency test mode (since region code is obtained from HTTP response headers)

# Also, during HTTPing, the software retrieves the region code from HTTP response headers (supports Cloudflare, AWS CloudFront, Fastly, Gcore, CDN77, Bunny, etc.) and displays it. TCPing cannot do this (but download testing can, since it's also an HTTP request)
```

> **IATA 3-letter airport codes**: https://www.cloudflarestatus.com/  
> **2-letter country codes**: [https://zh.wikipedia.org/wiki/ISO_3166-1二位字母代码#正式分配代码](https://zh.wikipedia.org/wiki/ISO_3166-1%E4%BA%8C%E4%BD%8D%E5%AD%97%E6%AF%8D%E4%BB%A3%E7%A0%81#%E6%AD%A3%E5%BC%8F%E5%88%86%E9%85%8D%E4%BB%A3%E7%A0%81)

</details>

****

#### \# Relative/Absolute Paths

<details>
<summary><code><strong>「 Click to expand content 」</strong></code></summary>

****

``` bash
# Specify IPv4 data file, no output to console, output to file (-p 0)
cfst.exe -f 1.txt -p 0 -dd

# Specify IPv4 data file, no output to file, show results directly (-p 10, -o empty but quotes required)
cfst.exe -f 2.txt -o "" -p 10 -dd

# Specify IPv4 data file and output to file (relative path, current directory; use quotes if path contains spaces)
cfst.exe -f 3.txt -o result.txt -dd


# Specify IPv4 data file and output to file (relative path, abc folder under current directory; use quotes if path contains spaces)
# Linux (abc folder under CFST program directory)
./cfst -f abc/3.txt -o abc/result.txt -dd

# Windows (note backslash)
cfst.exe -f abc\3.txt -o abc\result.txt -dd


# Specify IPv4 data file and output to file (absolute path, C:\abc\ directory; use quotes if path contains spaces)
# Linux (/abc/ directory)
./cfst -f /abc/4.txt -o /abc/result.csv -dd

# Windows (note backslash)
cfst.exe -f C:\abc\4.txt -o C:\abc\result.csv -dd


# If running CFST with absolute path, -f/-o file paths must also be absolute; otherwise "file not found" error!
# Linux (/abc/ directory)
/abc/cfst -f /abc/4.txt -o /abc/result.csv -dd

# Windows (note backslash)
C:\abc\cfst.exe -f C:\abc\4.txt -o C:\abc\result.csv -dd
```
</details>

****

#### \# Test Other Ports

<details>
<summary><code><strong>「 Click to expand content 」</strong></code></summary>

****

``` bash
# To test non-default port 443, use -tp parameter (affects both latency and download tests)

# To test port 80 for latency + download (if -dd disables download test, no need), specify http:// protocol download URL (won't force redirect to HTTPS)
cfst.exe -tp 80 -url http://cdn.cloudflare.steamstatic.com/steam/apps/5952/movie_max.webm

# For non-80/443 ports, ensure your download test URL supports access via that port.
```

</details>

****

#### \# Custom Test URL

<details>
<summary><code><strong>「 Click to expand content 」</strong></code></summary>

****

``` bash
# This parameter applies to download test and HTTP latency test. For the latter, the URL can be any webpage (not limited to specific files).

# Requirements: Directly downloadable, file size >200MB, uses Cloudflare CDN
cfst.exe -url https://cf.xiu2.xyz/url

# Note: If test URL uses HTTP protocol (won't force redirect to HTTPS), add -tp 80 (affects latency/download test port); for non-80/443 ports, ensure download URL supports that port.
cfst.exe -tp 80 -url http://cdn.cloudflare.steamstatic.com/steam/apps/5952/movie_max.webm
```

</details>

****

#### \# Custom Test Conditions (Specify Latency/Packet Loss/Download Speed Ranges)

<details>
<summary><code><strong>「 Click to expand content 」</strong></code></summary>

****

> Note: The **available count** on the latency test progress bar refers only to IPs that did not time out during latency testing, unrelated to latency limits.

- Only specify **[Max Average Latency]** condition

``` bash
# Max average latency: 200 ms, min download speed: 0 MB/s
# Find IPs with average latency < 200 ms, then perform 10 download tests sorted by ascending latency
cfst.exe -tl 200
```

> If no IP meets the latency condition, nothing is output.

****

- Only specify **[Max Average Latency]** condition, and only latency test (no download test)

``` bash
# Max average latency: 200 ms, min download speed: 0 MB/s, count: unknown
# Output only IPs with latency < 200ms; no download test (so -dn parameter is ignored)
cfst.exe -tl 200 -dd
```

- Only specify **[Max Packet Loss Rate]** condition

``` bash
# Max packet loss rate: 0.25
# Find IPs with packet loss rate <= 0.25; range 0.00~1.00; if -tlr 0, filter out any IP with packet loss
cfst.exe -tlr 0.25
```

****

- Only specify **[Min Download Speed]** condition

``` bash
# Max average latency: 9999 ms, min download speed: 5 MB/s, count: 10 (optional)
# Find 10 IPs with latency < 9999 ms and download speed > 5 MB/s before stopping
cfst.exe -sl 5 -dn 10
```

> If no IP meets the speed condition, nothing is output. You may need to lower your download speed expectation. Use `-debug` mode to ignore conditions and output all results—you'll see actual speeds and can adjust `-sl` accordingly.  
> Note: If you **don't specify** a download speed limit (`-sl`), CFST always outputs all results.

> If no average latency limit is specified, and you don't find enough IPs meeting the speed condition, testing continues indefinitely.  
> Suggest specifying both **[Min Download Speed]** and **[Max Average Latency]**; if target count isn't reached by latency limit, testing stops.

****

- Simultaneously specify **[Max Average Latency]** + **[Min Download Speed]** conditions

``` bash
# Both support decimals (e.g., -sl 0.5)
# Max average latency: 200 ms, min download speed: 5.6 MB/s, count: 10 (optional)
# Find 10 IPs with latency < 200 ms and download speed > 5.6 MB/s before stopping
cfst.exe -tl 200 -sl 5.6 -dn 10
```

> If no IP meets the latency condition, nothing is output.  
> If no IP meets the speed condition, nothing is output—but use `-debug` mode to ignore conditions and output all results (helps adjust conditions next time).  
> So first test without conditions to see typical latency/speed ranges, avoiding overly strict/lenient conditions!

> Since Cloudflare's published IP ranges include both **origin IPs** and **anycast IPs**, origin IPs are unusable—download speed is 0.00.  
> Run with `-sl 0.01` (min download speed) to filter out **origin IPs** (download speed < 0.01 MB/s).

****

To avoid confusion, here's what output you can expect under various condition combinations.

**No latency/speed conditions specified (all defaults):**
- Always output **all test results**

****

**Any latency condition specified (`-tl`, `-tll`, regardless of `-debug`):**
- If at least 1 IP meets condition, output **only those IPs** (if download test not disabled, continue download testing)  
- If no IP meets condition, output **nothing** (if download test not disabled, skip download test due to 0 count)

****

**Any download speed condition (`-sl`) specified:**

When **debug mode is off** (i.e., no `-debug` parameter; same logic as latency):

- If at least 1 IP meets condition, output **only those IPs**  
- If no IP meets condition, output **nothing**

When **debug mode is on** (i.e., `-debug` added; latency testing doesn't apply second rule below):

- If at least 1 IP meets condition, output **only those IPs**  
- If no IP meets condition, output **all test results**

</details>

****

#### \# Test a Single or Multiple IPs

<details>
<summary><code><strong>「 Click to expand content 」</strong></code></summary>

****

**Method One**:
Specify IPs directly via parameter.
``` bash
# Enter CFST directory, then run:
# Windows (in CMD)
cfst.exe -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32

# Linux
./cfst -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
```

****

**Method Two**:
Write these IPs into a text file, e.g., `1.txt`

```
1.1.1.1
1.1.1.200
1.0.0.1/24
2606:4700::/32
```

> For single IPs, `/32` subnet mask can be omitted (i.e., `1.1.1.1` equals `1.1.1.1/32`).  
> Subnet mask `/24` means last segment: `1.0.0.1~1.0.0.255`.

Then run CFST with `-f 1.txt` to specify IP range file.

``` bash
# Enter CFST directory, then run:
# Windows (in CMD)
cfst.exe -f 1.txt

# Linux
./cfst -f 1.txt

# For IP ranges like 1.0.0.1/24, only one random IP per range is tested; to test all IPs in the range, add -allip parameter.
```

</details>

****

#### \# Download Speed is 0.00?

<details>
<summary><code><strong>「 Click to expand content 」</strong></code></summary>
****
**\#\# Simple Principle Explanation:**

First, understand that CFST's download test is essentially the same as adding the `IP download-test-url-domain` to your hosts file and accessing the URL in a browser—it just automates this (similar to `curl -I --resolve download-test-url-domain:443:IP https://download-test-url`).

So if all download speeds are 0.00 MB/s, it means **an error occurred during download testing**, causing immediate termination (displayed as 0.00). Only these possibilities exist:

1. **Download test URL is invalid**
2. **Tested IP address is invalid**
3. **Your network has issues**

****

**\#\# Debug Mode Troubleshooting:**

First, **add `-debug`** to your original CFST command to enable debug mode, then re-run the test. Any errors during download testing will be shown for diagnosis.

Common download test failure reasons (Go language native error messages, mostly English):

1. `... read: connection reset by peer ...  `  
**Connection reset**, possibly download URL blocked or target IP HTTPS blocked; could be firewall or ISP (e.g., China Mobile or some regions' whitelist); or server simply reset your invalid request.
2. `... HTTP status code: 403 ...`  
Directly showing **HTTP status code** is easy to diagnose: 403 means access forbidden, 404 means file not found (search HTTP status codes for others)
3. `... context deadline exceeded (Client.Timeout exceeded while awaiting headers) ...`  
This is usually **request timeout**, possibly due to IP or network issues, or `-dt` download test time set too short (default 10s is not short)
4. `... tls: handshake failure ...` or `... tls: failed to verify certificate ...`  
These **TLS handshake failures/SSL certificate errors** mean the download URL and test IP server don't match—either the URL or IP is wrong (e.g., download URL is on Fastly CDN but test IP is Cloudflare, or vice versa)
5. `... tls: failed to verify certificate: x509: certificate is valid for XXX.XX, not YYY.YY ...`  
This means **SSL certificate doesn't include your download URL domain**—either the certificate is misconfigured, or the server IP doesn't have a certificate for that domain (e.g., using Google's server IP to test Baidu's domain)
6. `... tls: failed to verify certificate: x509: certificate has expired or is not yet valid: current time ...`  
This means **SSL certificate expired or not yet valid**; could also be same as 4/5 above (these 3 errors may occur together on same server IP)
7. `... tls: failed to verify certificate: x509: certificate signed by unknown authority.`  
This means **system certificate configuration is broken**, causing TLS handshake to fail. Only encountered in Termux (solution: see end of https://github.com/XIU2/CloudflareSpeedTest/discussions/61)

> If you encounter other errors and still don't understand after translation, open an Issue or Discussion—I'll update this section.  
> But when asking, **include full CFST output from debug mode** (or screenshot).

After checking above reasons, if still unresolved, try these further steps:

****

**1. Download test URL issue**:

Go to [#490](https://github.com/XIU2/CloudflareSpeedTest/discussions/490) and try other download URLs.

If one works, your original URL was invalid (note: default URL is a load-balanced redirect link that auto-redirects to community-shared公益 download URLs; availability varies by region; for stability, **build your own**—see methods in the post).

If many URLs still show 0.00, consider other possibilities.

****

**2. Tested IP address issue**:

The IPs you're testing may pass TCP tests but fail HTTP connections due to various reasons (e.g., origin IPs, enterprise-only IPs). Try other IPs.

****

**3. Your network issue**:

This is harder. If using a computer + broadband, try turning off Wi-Fi and enabling mobile data, then connect phone via USB to enable USB tethering (varies by phone OS; search for instructions), unplug Ethernet cable—now your computer uses mobile data. Run CFST again to see if results change (also try above methods to cross-validate).

If results improve, your broadband is the issue. If still 0.00, it's problematic...

****

</details>

****

#### \# Permanently Accelerate All Websites Using Cloudflare CDN (No Need to Add Domains to Hosts One by One)

I previously mentioned that the goal of this tool is to **accelerate access to Cloudflare CDN websites via hosts file modification**.

But as [**#8**](https://github.com/XIU2/CloudflareSpeedTest/issues/8) says, manually adding domains to hosts is **too tedious**. So I found a **permanent solution**! See this [**"Still Adding Hosts One by One? The Perfect Local Acceleration Method for All Cloudflare CDN Websites!"**](https://github.com/XIU2/CloudflareSpeedTest/discussions/71) and another [tutorial using local DNS service to modify domain resolution IP to selected IP](https://github.com/XIU2/CloudflareSpeedTest/discussions/317).

****

#### \# Auto-update Hosts

Considering many users need to replace IPs in hosts file after getting the fastest Cloudflare CDN IP.

See this [**Issues**](https://github.com/XIU2/CloudflareSpeedTest/discussions/312) for **Windows/Linux auto-update Hosts scripts**!

****

## Feedback

If you encounter issues, first check [**Issues**](https://github.com/XIU2/CloudflareSpeedTest/issues), [Discussions](https://github.com/XIU2/CloudflareSpeedTest/discussions) to see if others asked the same (check [**Closed**](https://github.com/XIU2/CloudflareSpeedTest/issues?q=is%3Aissue+is%3Aclosed)).  
If no similar issue, open a new [**Issues**](https://github.com/XIU2/CloudflareSpeedTest/issues/new).

> [!NOTE]
> **Note**! For anything **not related to CFST's issues or feature suggestions**, please use the project's forum (top `💬 Discussions`)  

****

## If this helped you, consider a donation~🎉✨

![WeChat Donation](https://github.com/XIU2/XIU2/blob/master/img/zs-01.png)![Alipay Donation](https://github.com/XIU2/XIU2/blob/master/img/zs-02.png)

****

## Derivative Projects

- _https://github.com/xianshenglu/cloudflare-ip-tester-app_  
_**CFST Android App [#202](https://github.com/XIU2/CloudflareSpeedTest/discussions/320)**_

- _https://github.com/mingxiaoyu/luci-app-cloudflarespeedtest_  
_**CFST OpenWrt Router Plugin [#174](https://github.com/XIU2/CloudflareSpeedTest/discussions/319)**_

- _https://github.com/immortalwrt-collections/openwrt-cdnspeedtest_  
_**CFST Native OpenWrt Build Version [#64](https://github.com/XIU2/CloudflareSpeedTest/discussions/64)**_

- _https://github.com/GuangYu-yu/CloudflareST-Rust_  
_**CFST Rust Version**_

- _https://github.com/hoseinnikkhah/CloudflareSpeedTest-English_  
_**English language version of CFST (Text language differences only) [#64](https://github.com/XIU2/CloudflareSpeedTest/issues/68)**_

> _Only lists some derivative projects promoted in this project; if any are missing, let me know~_

****

## Acknowledgments

- _https://github.com/Spedoske/CloudflareScanner_

> _Since that project hasn't been updated for a long time and I had many feature requests, I learned Go and built this (amateur)..._  
> _This software is based on that project but has been **completely refactored with many features added/bugs fixed**; features are actively added/optimized based on user feedback (free time)..._

****

## Manual Compilation

<details>
<summary><code><strong>「 Click to expand content 」</strong></code></summary>

****

For convenience, I embed the version number into the code's version variable during compilation. So when compiling manually, use `-ldflags` with `go build` to specify version:

```bash
go build -ldflags "-s -w -X main.version=v1.0.0"
# In CloudflareSpeedTest directory, run this command via CLI (e.g., CMD, Bat script) to compile a binary for your current system's OS, bitness, architecture with version v1.0.0 (Go auto-detects your system)
```

To compile for **other systems/architectures** on Windows 64-bit, specify **GOOS** and **GOARCH** variables.

E.g., compile Linux AMD64 binary on Windows:

```bat
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-s -w -X main.version=v1.0.0"
```

E.g., compile Windows 32-bit binary on Linux:

```bash
GOOS=windows
GOARCH=386
go build -ldflags "-s -w -X main.version=v1.0.0"
```

> Run `go tool dist list` to see supported combinations.

****

For batch compilation, I define a version variable; subsequent builds just use it.  
Also, batch compilation requires separate folders (or different filenames); use `-o` to specify.

```bat
:: Windows:
SET version=v1.0.0
SET GOOS=linux
SET GOARCH=amd64
go build -o Releases\cfst_linux_amd64\cfst -ldflags "-s -w -X main.version=%version%"
```

```bash
# Linux:
version=v1.0.0
GOOS=windows
GOARCH=386
go build -o Releases/cfst_windows_386/cfst.exe -ldflags "-s -w -X main.version=${version}"
```

</details>

****

## License

The GPL-3.0 License.
