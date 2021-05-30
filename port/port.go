package port

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type PortResult struct {
	Port    int
	State   bool
	Service string
}

type PortRange struct {
	Start, End int
}

type ScanResult struct {
	hostname string
	ip       []net.IP
	results  []PortResult
}

var common = map[int]string{
	7:    "echo",
	20:   "ftp",
	21:   "ftp",
	22:   "ssh",
	23:   "telnet",
	25:   "smtp",
	43:   "whois",
	53:   "dns",
	67:   "dhcp",
	68:   "dhcp",
	80:   "http",
	110:  "pop3",
	123:  "ntp",
	137:  "netbios",
	139:  "netbios",
	143:  "imap4",
	443:  "https",
	513:  "rlogin",
	540:  "rtsp",
	587:  "smtp",
	873:  "rsync",
	902:  "vmware",
	989:  "ftps",
	990:  "ftps",
	1194: "openvpn",
	3306: "mysql",
	5000: "unpn",
	8080: "https-proxy",
	8443: "https-alt",
	2345: "test",
}

func ScanPort(hostname string, port int) PortResult {
	result := PortResult{Port: port}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		result.State = false
		return result
	}

	defer conn.Close()
	result.State = true
	return result
}

func ScanPorts(hostname string, ports PortRange) (ScanResult, bool) {
	var results []PortResult
	var scanned ScanResult
	addr, err := net.LookupIP(hostname)

	if err != nil {
		return scanned, false
	}

	for i := ports.Start; i <= ports.End; i++ {
		if v, ok := common[i]; ok {
			result := ScanPort(hostname, i)
			result.Service = v
			results = append(results, result)
		}
	}
	scanned = ScanResult{
		hostname: hostname,
		ip:       addr,
		results:  results,
	}
	return scanned, true
}

func DisplayScanResult(result ScanResult) {
	ip := result.ip[len(result.ip)-1]
	fmt.Printf("Open ports for %s (%s)\n", result.hostname, ip.String())
	for _, v := range result.results {
		if v.State {
			fmt.Printf("%d %s\n", v.Port, v.Service)
		}
	}
}

func GetOpenPorts(hostname string, ports PortRange) {
	scanned, ok := ScanPorts(hostname, ports)
	if ok {
		DisplayScanResult(scanned)
	} else {
		fmt.Printf("Error: Invalid IP address")
	}
}
