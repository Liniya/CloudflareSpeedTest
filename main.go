package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/XIU2/CloudflareSpeedTest/task"
	"github.com/XIU2/CloudflareSpeedTest/utils"
)

var (
	version, versionNew string
)

func init() {
	var printVersion bool
	var help = `
CloudflareSpeedTest ` + version + `
Test the latency and speed of all CDN or website IPs to find the fastest IP (IPv4+IPv6)!
https://github.com/XIU2/CloudflareSpeedTest

Parameters:
    -n 200
        Number of threads for latency testing; more threads speed up latency testing, but avoid setting too high on low-performance devices (e.g., routers); (default: 200, max: 1000)
    -t 4
        Number of ping attempts per IP; number of times to test latency for each IP; (default: 4)
    -dn 10
        Number of download speed tests; after latency testing and sorting, the number of IPs to test for download speed from the lowest latency; (default: 10)
    -dt 10
        Download test duration; maximum time for each IP's download speed test; cannot be too short; (default: 10 seconds)
    -tp 443
        Specify test port; port used for latency and download tests; (default: 443)
    -url https://cf.xiu2.xyz/url
        Specify test URL; URL used for latency (HTTPing) and download tests; default URL is not guaranteed to be available; recommend using a self-hosted one;

    -httping
        Switch test mode; use HTTP protocol for latency testing instead of TCP; test URL is specified by the [-url] parameter; (default: TCPing)
    -httping-code 200
        Valid status code; HTTP status code considered valid during HTTPing latency tests; only one code allowed; (default: 200 301 302)
    -cfcolo HKG,KHH,NRT,LAX,SEA,SJC,FRA,MAD
        Match specific regions; IATA airport codes or country/city codes, separated by commas; only available in HTTPing mode; (default: all regions)

    -tl 200
        Maximum average latency; only output IPs with average latency below the specified value; conditions can be combined; (default: 9999 ms)
    -tll 40
        Minimum average latency; only output IPs with average latency above the specified value; (default: 0 ms)
    -tlr 0.2
        Maximum packet loss rate; only output IPs with packet loss rate less than or equal to the specified value; range: 0.00~1.00; 0 filters out any IP with packet loss; (default: 1.00)
    -sl 5
        Minimum download speed; only output IPs with download speed above the specified value; stops testing once the specified number [-dn] is reached; (default: 0.00 MB/s)

    -p 10
        Number of results to display; directly show the specified number of results after testing; if 0, no results are displayed and the program exits; (default: 10)
    -f ip.txt
        IP range data file; if path contains spaces, enclose in quotes; supports IP ranges from other CDNs; (default: ip.txt)
    -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
        Specify IP ranges directly via parameter; comma-separated list of IP ranges; (default: empty)
    -o result.csv
        Write results to file; if path contains spaces, enclose in quotes; if empty, no file is written [-o ""]; (default: result.csv)

    -dd
        Disable download speed test; after disabling, results are sorted by latency instead of download speed; (default: enabled)
    -allip
        Test all IPs; test every IP in the IP range (IPv4 only); (default: randomly test one IP per /24 segment)

    -debug
        Debug mode; outputs additional logs during unexpected situations to help diagnose issues; (default: disabled)

    -v
        Print program version and check for updates
    -h
        Print help information
`
	var minDelay, maxDelay, downloadTime int
	var maxLossRate float64
	flag.IntVar(&task.Routines, "n", 200, "Latency test threads")
	flag.IntVar(&task.PingTimes, "t", 4, "Latency test attempts")
	flag.IntVar(&task.TestCount, "dn", 10, "Download test count")
	flag.IntVar(&downloadTime, "dt", 10, "Download test duration")
	flag.IntVar(&task.TCPPort, "tp", 443, "Specify test port")
	flag.StringVar(&task.URL, "url", "https://cf.xiu2.xyz/url", "Specify test URL")

	flag.BoolVar(&task.Httping, "httping", false, "Switch test mode")
	flag.IntVar(&task.HttpingStatusCode, "httping-code", 0, "Valid status code")
	flag.StringVar(&task.HttpingCFColo, "cfcolo", "", "Match specific regions")

	flag.IntVar(&maxDelay, "tl", 9999, "Maximum average latency")
	flag.IntVar(&minDelay, "tll", 0, "Minimum average latency")
	flag.Float64Var(&maxLossRate, "tlr", 1, "Maximum packet loss rate")
	flag.Float64Var(&task.MinSpeed, "sl", 0, "Minimum download speed")

	flag.IntVar(&utils.PrintNum, "p", 10, "Number of results to display")
	flag.StringVar(&task.IPFile, "f", "ip.txt", "IP range data file")
	flag.StringVar(&task.IPText, "ip", "", "Specify IP ranges")
	flag.StringVar(&utils.Output, "o", "result.csv", "Output results file")

	flag.BoolVar(&task.Disable, "dd", false, "Disable download test")
	flag.BoolVar(&task.TestAll, "allip", false, "Test all IPs")

	flag.BoolVar(&utils.Debug, "debug", false, "Debug mode")

	flag.BoolVar(&printVersion, "v", false, "Print version")
	flag.Usage = func() { fmt.Print(help) }
	flag.Parse()

	if task.MinSpeed > 0 && time.Duration(maxDelay)*time.Millisecond == utils.InputMaxDelay {
		utils.Yellow.Println("[Tip] When using [-sl], it is recommended to combine with [-tl] to avoid endlessly testing due to insufficient number of IPs meeting [-dn] requirement...")
	}
	utils.InputMaxDelay = time.Duration(maxDelay) * time.Millisecond
	utils.InputMinDelay = time.Duration(minDelay) * time.Millisecond
	utils.InputMaxLossRate = float32(maxLossRate)
	task.Timeout = time.Duration(downloadTime) * time.Second
	task.HttpingCFColomap = task.MapColoMap()

	if printVersion {
		println(version)
		fmt.Println("Checking for updates...")
		checkUpdate()
		if versionNew != "" {
			utils.Yellow.Printf("*** New version [%s] detected! Please update at [https://github.com/XIU2/CloudflareSpeedTest] ***", versionNew)
		} else {
			utils.Green.Println("You are on the latest version [" + version + "]!")
		}
		os.Exit(0)
	}
}

func main() {
	task.InitRandSeed() // Initialize random seed

	fmt.Printf("# XIU2/CloudflareSpeedTest %s \n\n", version)

	// Start latency test + filter by latency/packet loss
	pingData := task.NewPing().Run().FilterDelay().FilterLossRate()
	// Start download speed test
	speedData := task.TestDownloadSpeed(pingData)
	utils.ExportCsv(speedData) // Export to file
	speedData.Print()          // Print results
	endPrint()                 // Exit appropriately (for Windows)
}

// Exit appropriately (for Windows)
func endPrint() {
	if utils.NoPrintResult() { // If no results need to be printed, exit directly
		return
	}
	if runtime.GOOS == "windows" { // If on Windows, require Enter key press or Ctrl+C to exit (prevents auto-closing when double-clicking to run)
		fmt.Printf("Press Enter or Ctrl+C to exit.")
		fmt.Scanln()
	}
}

// Check for updates
func checkUpdate() {
	timeout := 10 * time.Second
	client := http.Client{Timeout: timeout}
	res, err := client.Get("https://api.xiu2.xyz/ver/cloudflarespeedtest.txt")
	if err != nil {
		return
	}
	// Read resource data body: []byte
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	// Close resource stream
	defer res.Body.Close()
	if string(body) != version {
		versionNew = string(body)
	}
}
