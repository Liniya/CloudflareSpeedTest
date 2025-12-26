package task

import (
	"bufio"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const defaultInputFile = "ip.txt"

var (
	// TestAll test all ip
	TestAll = false
	// IPFile is the filename of IP Rangs
	IPFile = defaultInputFile
	IPText string
)

func InitRandSeed() {
	rand.Seed(time.Now().UnixNano())
}

func isIPv4(ip string) bool {
	return strings.Contains(ip, ".")
}

func randIPEndWith(num byte) byte {
	if num == 0 { // For /32 single IP
		return byte(0)
	}
	return byte(rand.Intn(int(num)))
}

type IPRanges struct {
	ips     []*net.IPAddr
	mask    string
	firstIP net.IP
	ipNet   *net.IPNet
}

func newIPRanges() *IPRanges {
	return &IPRanges{
		ips: make([]*net.IPAddr, 0),
	}
}

// If it's a single IP, add subnet mask; otherwise get subnet mask (r.mask)
func (r *IPRanges) fixIP(ip string) string {
	// If there's no '/' then it's a single IP, so add /32 or /128 subnet mask
	if i := strings.IndexByte(ip, '/'); i < 0 {
		if isIPv4(ip) {
			r.mask = "/32"
		} else {
			r.mask = "/128"
		}
		ip += r.mask
	} else {
		r.mask = ip[i:]
	}
	return ip
}

// Parse IP range to get IP, IP range, and subnet mask
func (r *IPRanges) parseCIDR(ip string) {
	var err error
	if r.firstIP, r.ipNet, err = net.ParseCIDR(r.fixIP(ip)); err != nil {
		log.Fatalln("ParseCIDR err", err)
	}
}

func (r *IPRanges) appendIPv4(d byte) {
	r.appendIP(net.IPv4(r.firstIP[12], r.firstIP[13], r.firstIP[14], d))
}

func (r *IPRanges) appendIP(ip net.IP) {
	r.ips = append(r.ips, &net.IPAddr{IP: ip})
}

// Return the minimum and available count of the fourth octet
func (r *IPRanges) getIPRange() (minIP, hosts byte) {
	minIP = r.firstIP[15] & r.ipNet.Mask[3] // Minimum value of the fourth octet

	// Get number of available hosts based on subnet mask
	m := net.IPv4Mask(255, 255, 255, 255)
	for i, v := range r.ipNet.Mask {
		m[i] ^= v
	}
	total, _ := strconv.ParseInt(m.String(), 16, 32) // Total available IPs
	if total > 255 {                                 // Correct available IP count for fourth octet
		hosts = 255
		return
	}
	hosts = byte(total)
	return
}

func (r *IPRanges) chooseIPv4() {
	if r.mask == "/32" { // For single IP, no randomization needed; just add it directly
		r.appendIP(r.firstIP)
	} else {
		minIP, hosts := r.getIPRange()    // Get minimum and available count of fourth octet
		for r.ipNet.Contains(r.firstIP) { // Continue looping while IP is within range
			if TestAll { // If testing all IPs
				for i := 0; i <= int(hosts); i++ { // Iterate through fourth octet from min to max
					r.appendIPv4(byte(i) + minIP)
				}
			} else { // Randomize last octet of IP
				r.appendIPv4(minIP + randIPEndWith(hosts))
			}
			r.firstIP[14]++ // 0.0.(X+1).X
			if r.firstIP[14] == 0 {
				r.firstIP[13]++ // 0.(X+1).X.X
				if r.firstIP[13] == 0 {
					r.firstIP[12]++ // (X+1).X.X.X
				}
			}
		}
	}
}

func (r *IPRanges) chooseIPv6() {
	if r.mask == "/128" { // For single IP, no randomization needed; just add it directly
		r.appendIP(r.firstIP)
	} else {
		var tempIP uint8                  // Temporary variable to store previous octet value
		for r.ipNet.Contains(r.firstIP) { // Continue looping while IP is within range
			r.firstIP[15] = randIPEndWith(255) // Randomize last octet
			r.firstIP[14] = randIPEndWith(255) // Randomize second-to-last octet

			targetIP := make([]byte, len(r.firstIP))
			copy(targetIP, r.firstIP)
			r.appendIP(targetIP) // Add IP to pool

			for i := 13; i >= 0; i-- { // Randomize from third-to-last octet backwards
				tempIP = r.firstIP[i]              // Save previous octet value
				r.firstIP[i] += randIPEndWith(255) // Add random value 0~255 to current octet
				if r.firstIP[i] >= tempIP {        // If current octet >= previous, randomization succeeded; exit loop
					break
				}
			}
		}
	}
}

func loadIPRanges() []*net.IPAddr {
	ranges := newIPRanges()
	if IPText != "" { // Get IP ranges from parameter
		IPs := strings.Split(IPText, ",") // Split by comma into array and iterate
		for _, IP := range IPs {
			IP = strings.TrimSpace(IP) // Trim leading/trailing whitespace (spaces, tabs, newlines)
			if IP == "" {              // Skip empty entries (e.g., leading, trailing, or consecutive ,, cases)
				continue
			}
			ranges.parseCIDR(IP) // Parse IP range to get IP, range, subnet mask
			if isIPv4(IP) {      // Generate IPs to test (single/random/all)
				ranges.chooseIPv4()
			} else {
				ranges.chooseIPv6()
			}
		}
	} else { // Get IP ranges from file
		if IPFile == "" {
			IPFile = defaultInputFile
		}
		file, err := os.Open(IPFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() { // Iterate through each line in file
			line := strings.TrimSpace(scanner.Text()) // Trim leading/trailing whitespace
			if line == "" {                           // Skip empty lines
				continue
			}
			ranges.parseCIDR(line) // Parse IP range to get IP, range, subnet mask
			if isIPv4(line) {      // Generate IPs to test (single/random/all)
				ranges.chooseIPv4()
			} else {
				ranges.chooseIPv6()
			}
		}
	}
	return ranges.ips
}
