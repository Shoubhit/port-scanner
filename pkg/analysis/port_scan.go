package analysis

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port    string
	State   string
	Service string
}

func getServiceName(port int) string {
	// Define a mapping between port numbers and corresponding services
	serviceMap := map[int]string{
		22:   "SSH",
		53:   "DNS",
		80:   "HTTP",
		443:  "HTTPS",
		8080: "HTTP Alt",
		8443: "HTTPS Alt",
		3389: "Remote Desktop",
		20:   "FTP Data",
		21:   "FTP Control",
		23:   "Telnet",
		1433: "Microsoft SQL",
		1434: "Microsoft SQL Monitor",
		3306: "MySQL",
		137:  "NetBIOS Name Service",
		138:  "NetBIOS Datagram Service",
		139:  "NetBIOS Session Service",
		445:  "Microsoft-DS",
	}

	// Check if the port is in the mapping
	if serviceName, ok := serviceMap[port]; ok {
		return serviceName
	}

	// Default service name for unknown ports
	return "Unknown"
}

func scanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: strconv.Itoa(port) + string("/") + protocol}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 5*time.Second)

	if err != nil {
		result.Service = getServiceName(port)
		result.State = "Closed"
		return result
	}
	defer conn.Close()

	// Set the service name based on the port number
	result.Service = getServiceName(port)
	result.State = "Open"
	return result
}

func InitialScan(hostname string) []ScanResult {
	// Define the list of ports to scan
	udpPorts := []int{137, 138, 139, 445}
	tcpPorts := []int{22, 53, 80, 443, 8080, 8443, 3389, 20, 21, 23, 1433, 1434, 3306}

	var results []ScanResult

	// Scan UDP ports
	for _, port := range udpPorts {
		results = append(results, scanPort("udp", hostname, port))
	}

	// Scan TCP ports
	for _, port := range tcpPorts {
		results = append(results, scanPort("tcp", hostname, port))
	}

	return results
}

func WideScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 0; i <= 49152; i++ {
		results = append(results, scanPort("udp", hostname, i))
	}

	for i := 0; i <= 49152; i++ {
		results = append(results, scanPort("tcp", hostname, i))
	}

	return results
}
