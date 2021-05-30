package main

import (
	"github.com/toannguyen3105/port-scanner/port"
)

func main() {
	port.GetOpenPorts("www.freecodecamp.com", port.PortRange{Start: 75, End: 85})

	// called with ip address
	port.GetOpenPorts("104.26.10.78", port.PortRange{Start: 8079, End: 8090})

	// verbose called with ip address and no host name returnd -- single open port
	port.GetOpenPorts("137.74.187.104", port.PortRange{Start: 400, End: 450})

	// verbose called with host name -- multiple ports returned
	port.GetOpenPorts("scanme.nmap.,org", port.PortRange{Start: 20, End: 80})
}
