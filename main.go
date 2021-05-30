package main

import (
	"github.com/toannguyen3105/port-scanner/port"
)

func main() {
	// Scan with domain
	port.GetOpenPorts("www.value24h.vn", port.PortRange{Start: 75, End: 85})

	// Scan with IP
	port.GetOpenPorts("103.221.220.245", port.PortRange{Start: 400, End: 450})
}
