package task

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/XIU2/CloudflareSpeedTest/utils"

	"github.com/VividCortex/ewma"
)

const (
	bufferSize                     = 1024
	defaultURL                     = "https://cf.xiu2.xyz/url"
	defaultTimeout                 = 10 * time.Second
	defaultDisableDownload         = false
	defaultTestNum                 = 10
	defaultMinSpeed        float64 = 0.0
)

var (
	URL     = defaultURL
	Timeout = defaultTimeout
	Disable = defaultDisableDownload

	TestCount = defaultTestNum
	MinSpeed  = defaultMinSpeed
)

func checkDownloadDefault() {
	if URL == "" {
		URL = defaultURL
	}
	if Timeout <= 0 {
		Timeout = defaultTimeout
	}
	if TestCount <= 0 {
		TestCount = defaultTestNum
	}
	if MinSpeed <= 0.0 {
		MinSpeed = defaultMinSpeed
	}
}

func TestDownloadSpeed(ipSet utils.PingDelaySet) (speedSet utils.DownloadSpeedSet) {
	checkDownloadDefault()
	if Disable {
		return utils.DownloadSpeedSet(ipSet)
	}
	if len(ipSet) <= 0 { // Continue download test only if IP array length > 0
		utils.Yellow.Println("[Info] Latency test result has 0 IPs, skipping download test.")
		return
	}
	testNum := TestCount                        // Number of IPs waiting for download test; default equals download test count (-dn)
	if len(ipSet) < TestCount || MinSpeed > 0 { // If filtered IP array length is less than download test count (-dn), or if minimum download speed (-sl) is specified (may require testing all until target number is reached), adjust waiting queue to IP array length
		testNum = len(ipSet)
	}
	if testNum < TestCount { // If waiting queue is less than download test count (-dn), adjust download test count to waiting queue size
		TestCount = testNum
	}

	utils.Cyan.Printf("Starting download speed test (Min: %.2f MB/s, Count: %d, Queue: %d)\n", MinSpeed, TestCount, testNum)
	// Control download test progress bar length to match latency test progress bar length (perfectionist)
	bar_a := len(strconv.Itoa(len(ipSet)))
	bar_b := "     "
	for i := 0; i < bar_a; i++ {
		bar_b += " "
	}
	bar := utils.NewBar(TestCount, bar_b, "")
	for i := 0; i < testNum; i++ {
		speed, colo := downloadHandler(ipSet[i].IP)
		ipSet[i].DownloadSpeed = speed
		if ipSet[i].Colo == "" { // Only write if Colo is empty (meaning it wasn't retrieved during HTTPing)
			ipSet[i].Colo = colo
		}
		// Filter results after each IP's download test using [minimum download speed] condition
		if speed >= MinSpeed*1024*1024 {
			bar.Grow(1, "")
			speedSet = append(speedSet, ipSet[i]) // Add to new array if above minimum download speed
			if len(speedSet) == TestCount {       // If target number of IPs is reached, break loop
				break
			}
		}
	}
	bar.Done()
	if MinSpeed == 0.00 { // If no minimum download speed specified, return all test data
		speedSet = utils.DownloadSpeedSet(ipSet)
	} else if utils.Debug && len(speedSet) == 0 { // If minimum download speed is specified, in debug mode, and no IPs meet condition, return all data for user to adjust conditions
		utils.Yellow.Println("[Debug] No IPs met the minimum download speed condition; returning all test data (to help adjust conditions next time).")
		speedSet = utils.DownloadSpeedSet(ipSet)
	}
	// Sort by speed
	sort.Sort(speedSet)
	return
}

func getDialContext(ip *net.IPAddr) func(ctx context.Context, network, address string) (net.Conn, error) {
	var fakeSourceAddr string
	if isIPv4(ip.String()) {
		fakeSourceAddr = fmt.Sprintf("%s:%d", ip.String(), TCPPort)
	} else {
		fakeSourceAddr = fmt.Sprintf("[%s]:%d", ip.String(), TCPPort)
	}
	return func(ctx context.Context, network, address string) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, network, fakeSourceAddr)
	}
}

// Unified request error debug output
func printDownloadDebugInfo(ip *net.IPAddr, err error, statusCode int, url, lastRedirectURL string, response *http.Response) {
	finalURL := url // Default final URL (so it outputs even if response is nil)
	if lastRedirectURL != "" {
		finalURL = lastRedirectURL // If lastRedirectURL is not empty, redirection occurred; prioritize outputting the final redirected URL
	} else if response != nil && response.Request != nil && response.Request.URL != nil {
		finalURL = response.Request.URL.String() // If response is not nil and Request and URL are not nil, get last successful response address
	}
	if url != finalURL { // If URL and final address differ, redirection occurred; error caused by redirected address
		if statusCode > 0 { // If status code > 0, error caused by HTTP status code
			utils.Red.Printf("[Debug] IP: %s, Download test terminated, HTTP status code: %d, download URL: %s, erroneous redirected URL: %s\n", ip.String(), statusCode, url, finalURL)
		} else {
			utils.Red.Printf("[Debug] IP: %s, Download test failed, error: %v, download URL: %s, erroneous redirected URL: %s\n", ip.String(), err, url, finalURL)
		}
	} else { // If URL and final address match, no redirection occurred
		if statusCode > 0 { // If status code > 0, error caused by HTTP status code
			utils.Red.Printf("[Debug] IP: %s, Download test terminated, HTTP status code: %d, download URL: %s\n", ip.String(), statusCode, url)
		} else {
			utils.Red.Printf("[Debug] IP: %s, Download test failed, error: %v, download URL: %s\n", ip.String(), err, url)
		}
	}
}

// return download Speed
func downloadHandler(ip *net.IPAddr) (float64, string) {
	var lastRedirectURL string // Record last redirect target for output during errors
	client := &http.Client{
		Transport: &http.Transport{DialContext: getDialContext(ip)},
		Timeout:   Timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			lastRedirectURL = req.URL.String() // Record each redirect target for output during errors
			if len(via) > 10 {                 // Limit to max 10 redirects
				if utils.Debug { // Debug mode: output more info
					utils.Red.Printf("[Debug] IP: %s, Too many redirects in download test, terminating, download URL: %s\n", ip.String(), req.URL.String())
				}
				return http.ErrUseLastResponse
			}
			if req.Header.Get("Referer") == defaultURL { // When using default download URL, don't carry Referer during redirect
				req.Header.Del("Referer")
			}
			return nil
		},
	}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		if utils.Debug { // Debug mode: output more info
			utils.Red.Printf("[Debug] IP: %s, Failed to create download test request, error: %v, download URL: %s\n", ip.String(), err, URL)
		}
		return 0.0, ""
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.80 Safari/537.36")

	response, err := client.Do(req)
	if err != nil {
		if utils.Debug { // Debug mode: output more info
			printDownloadDebugInfo(ip, err, 0, URL, lastRedirectURL, response)
		}
		return 0.0, ""
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		if utils.Debug { // Debug mode: output more info
			printDownloadDebugInfo(ip, nil, response.StatusCode, URL, lastRedirectURL, response)
		}
		return 0.0, ""
	}

	// Get region code from header
	colo := getHeaderColo(response.Header)

	timeStart := time.Now()           // Start time (current)
	timeEnd := timeStart.Add(Timeout) // End time = start + download test duration

	contentLength := response.ContentLength // File size
	buffer := make([]byte, bufferSize)

	var (
		contentRead     int64 = 0
		timeSlice             = Timeout / 100
		timeCounter           = 1
		lastContentRead int64 = 0
	)

	var nextTime = timeStart.Add(timeSlice * time.Duration(timeCounter))
	e := ewma.NewMovingAverage()

	// Loop to calculate; exit when file is fully downloaded (contentRead == contentLength)
	for contentLength != contentRead {
		currentTime := time.Now()
		if currentTime.After(nextTime) {
			timeCounter++
			nextTime = timeStart.Add(timeSlice * time.Duration(timeCounter))
			e.Add(float64(contentRead - lastContentRead))
			lastContentRead = contentRead
		}
		// If exceeded download test duration, exit loop (terminate test)
		if currentTime.After(timeEnd) {
			break
		}
		bufferRead, err := response.Body.Read(buffer)
		if err != nil {
			if err != io.EOF { // If error occurs during download and it's not EOF, exit loop (terminate test)
				break
			} else if contentLength == -1 { // File downloaded completely and size is unknown; exit loop (terminate test), e.g., https://speed.cloudflare.com/__down?bytes=200000000; if downloaded within 10s, speed result may be significantly lower or show as 0.00 (too fast)
				break
			}
			// Get previous time slice
			last_time_slice := timeStart.Add(timeSlice * time.Duration(timeCounter-1))
			// Downloaded data / (current time - previous time slice / time slice)
			e.Add(float64(contentRead-lastContentRead) / (float64(currentTime.Sub(last_time_slice)) / float64(timeSlice)))
		}
		contentRead += int64(bufferRead)
	}
	return e.Value() / (Timeout.Seconds() / 120), colo
}
