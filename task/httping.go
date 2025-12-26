package task

import (
	//"crypto/tls"

	"io"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/XIU2/CloudflareSpeedTest/utils"
)

var (
	Httping               bool
	HttpingStatusCode     int
	HttpingCFColo         string
	HttpingCFColomap      *sync.Map
	RegexpColoIATACode    = regexp.MustCompile(`[A-Z]{3}`)  // Match IATA airport codes (commonly known as three-letter airport codes)
	RegexpColoCountryCode = regexp.MustCompile(`[A-Z]{2}`)  // Match country codes (e.g., US, CN, UK)
	RegexpColoGcore       = regexp.MustCompile(`^[a-z]{2}`) // Match city codes (lowercase, e.g., us, cn, uk)
)

// pingReceived pingTotalTime
func (p *Ping) httping(ip *net.IPAddr) (int, time.Duration, string) {
	hc := http.Client{
		Timeout: time.Second * 2,
		Transport: &http.Transport{
			DialContext: getDialContext(ip),
			//TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Skip certificate verification
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // Prevent redirection
		},
	}

	// First request to get HTTP status code and region code
	var colo string
	{
		request, err := http.NewRequest(http.MethodHead, URL, nil)
		if err != nil {
			if utils.Debug { // Debug mode: output more info
				utils.Red.Printf("[Debug] IP: %s, Failed to create latency test request, error: %v, test URL: %s\n", ip.String(), err, URL)
			}
			return 0, 0, ""
		}
		request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.80 Safari/537.36")
		response, err := hc.Do(request)
		if err != nil {
			if utils.Debug { // Debug mode: output more info
				utils.Red.Printf("[Debug] IP: %s, Latency test failed, error: %v, test URL: %s\n", ip.String(), err, URL)
			}
			return 0, 0, ""
		}
		defer response.Body.Close()

		//fmt.Println("IP:", ip, "StatusCode:", response.StatusCode, response.Request.URL)
		// If no HTTP status code is specified or the specified code is invalid, default to considering only 200, 301, 302 as successful
		if HttpingStatusCode == 0 || HttpingStatusCode < 100 && HttpingStatusCode > 599 {
			if response.StatusCode != 200 && response.StatusCode != 301 && response.StatusCode != 302 {
				if utils.Debug { // Debug mode: output more info
					utils.Red.Printf("[Debug] IP: %s, Latency test terminated, HTTP status code: %d, test URL: %s\n", ip.String(), response.StatusCode, URL)
				}
				return 0, 0, ""
			}
		} else {
			if response.StatusCode != HttpingStatusCode {
				if utils.Debug { // Debug mode: output more info
					utils.Red.Printf("[Debug] IP: %s, Latency test terminated, HTTP status code: %d, specified HTTP status code: %d, test URL: %s\n", ip.String(), response.StatusCode, HttpingStatusCode, URL)
				}
				return 0, 0, ""
			}
		}

		io.Copy(io.Discard, response.Body)

		// Get region code from header
		colo = getHeaderColo(response.Header)

		// Only match region codes if specified
		if HttpingCFColo != "" {
			// Check if region code matches specified regions
			colo = p.filterColo(colo)
			if colo == "" { // No match or doesn't meet specified region; terminate test for this IP
				if utils.Debug { // Debug mode: output more info
					utils.Red.Printf("[Debug] IP: %s, Region code mismatch: %s\n", ip.String(), colo)
				}
				return 0, 0, ""
			}
		}
	}

	// Loop to measure latency
	success := 0
	var delay time.Duration
	for i := 0; i < PingTimes; i++ {
		request, err := http.NewRequest(http.MethodHead, URL, nil)
		if err != nil {
			log.Fatal("Unexpected error, please report:", err)
			return 0, 0, ""
		}
		request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.80 Safari/537.36")
		if i == PingTimes-1 {
			request.Header.Set("Connection", "close")
		}
		startTime := time.Now()
		response, err := hc.Do(request)
		if err != nil {
			continue
		}
		success++
		io.Copy(io.Discard, response.Body)
		_ = response.Body.Close()
		duration := time.Since(startTime)
		delay += duration
	}

	return success, delay, colo
}

func MapColoMap() *sync.Map {
	if HttpingCFColo == "" {
		return nil
	}
	// Convert -cfcolo parameter regions to uppercase and format
	colos := strings.Split(strings.ToUpper(HttpingCFColo), ",")
	colomap := &sync.Map{}
	for _, colo := range colos {
		colomap.Store(colo, colo)
	}
	return colomap
}

// Extract region code from response header
func getHeaderColo(header http.Header) (colo string) {
	if header.Get("server") != "" {
		// If it's Cloudflare CDN
		// server: cloudflare
		// cf-ray: 7bd32409eda7b020-SJC
		if header.Get("server") == "cloudflare" {
			if colo = header.Get("cf-ray"); colo != "" {
				return RegexpColoIATACode.FindString(colo)
			}
		}
		// If it's CDN77 CDN (test URL: https://www.cdn77.com)
		// server: CDN77-Turbo
		// x-77-pop: losangelesUSCA // US shows as USCA, not compatible yet; extract only US
		// x-77-pop: frankfurtDE
		// x-77-pop: amsterdamNL
		// x-77-pop: singaporeSG
		if header.Get("server") == "CDN77-Turbo" {
			if colo = header.Get("x-77-pop"); colo != "" {
				return RegexpColoCountryCode.FindString(colo)
			}
		}
		// If it's Bunny CDN (test URL: https://bunny.net)
		// server: BunnyCDN-TW1-1121
		if colo = header.Get("server"); strings.Contains(colo, "BunnyCDN-") {
			return RegexpColoCountryCode.FindString(strings.TrimPrefix(colo, "BunnyCDN-")) // Remove BunnyCDN- prefix then match
		}
	}
	// If it's AWS CloudFront CDN (test URL: https://d7uri8nf7uskq.cloudfront.net/tools/list-cloudfront-ips)
	// x-amz-cf-pop: SIN52-P1
	if colo = header.Get("x-amz-cf-pop"); colo != "" {
		return RegexpColoIATACode.FindString(colo)
	}
	// If it's Fastly CDN (test URL: https://fastly.jsdelivr.net/gh/XIU2/CloudflareSpeedTest@master/go.mod)
	// x-served-by: cache-qpg1275-QPG
	// x-served-by: cache-fra-etou8220141-FRA, cache-hhr-khhr2060043-HHR (last one is actual location)
	if colo = header.Get("x-served-by"); colo != "" {
		if matches := RegexpColoIATACode.FindAllString(colo, -1); len(matches) > 0 {
			return matches[len(matches)-1] // Fastly's x-served-by may contain multiple region codes; take the last one
		}
	}
	// Gcore CDN headers (city codes only, not country codes), test URL: https://assets.gcore.pro/assets/icons/shield-lock.svg
	// x-id-fe: fr5-hw-edge-gc17
	// x-shard: fr5-shard0-default
	// x-id: fr5-hw-edge-gc28
	if colo = header.Get("x-id-fe"); colo != "" {
		if colo = RegexpColoGcore.FindString(colo); colo != "" {
			return strings.ToUpper(colo) // Convert lowercase region code to uppercase
		}
	}

	// If no header info is found, it's not a supported CDN; return empty string
	return ""
}

// Process region code
func (p *Ping) filterColo(colo string) string {
	if colo == "" {
		return ""
	}
	// If -cfcolo parameter is not specified, return directly
	if HttpingCFColomap == nil {
		return colo
	}
	// Match if airport code matches specified region
	_, ok := HttpingCFColomap.Load(colo)
	if ok {
		return colo
	}
	return ""
}
