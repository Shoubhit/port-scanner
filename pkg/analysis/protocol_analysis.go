package analysis

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Function to detect potential attack indicators in TCP packets
func detectTCPAttack(tcp *gopacket.TCP) {
	// Add your logic here to analyze TCP packets for potential attack indicators
	// Example: Check for unusual flags
	if tcp.FIN || tcp.URG {
		fmt.Println("Potential TCP attack detected: Unusual FIN or URG flags")
	}
}

// Function to detect potential attack indicators in HTTP packets
func detectHTTPAttack(applicationLayerPayload []byte) {
	// Add your logic here to analyze HTTP packets for potential attack indicators
	// Example: Check for SQL injection attempts
	if bytesContain(applicationLayerPayload, []byte("DROP TABLE")) {
		fmt.Println("Potential SQL injection attempt detected")
	}
}

// Helper function to check if a byte slice contains a substring
func bytesContain(source, target []byte) bool {
	return len(source) >= len(target) && bytesIndex(source, target) != -1
}

// Helper function to find the index of a substring in a byte slice
func bytesIndex(source, target []byte) int {
	for i := 0; i <= len(source)-len(target); i++ {
		if bytesEqual(source[i:i+len(target)], target) {
			return i
		}
	}
	return -1
}

// Helper function to check if two byte slices are equal
func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	handle, err := pcap.OpenLive("en0", 1600, true, 1*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		// Handle TCP packets
		if tcpLayer := packet.Layer(gopacket.LayerTypeTCP); tcpLayer != nil {
			tcp, _ := tcpLayer.(*gopacket.TCP)
			fmt.Printf("TCP Packet: %s:%d -> %s:%d\n", packet.NetworkLayer().NetworkFlow().Src(), tcp.SrcPort, packet.NetworkLayer().NetworkFlow().Dst(), tcp.DstPort)
			detectTCPAttack(tcp)
		}

		// Handle HTTP packets (assuming port 80 for simplicity)
		if applicationLayer := packet.ApplicationLayer(); applicationLayer != nil && packet.TransportLayer().TransportFlow().Dst().Port == 80 {
			fmt.Printf("HTTP Packet:\n%s\n", applicationLayer.Payload())
			detectHTTPAttack(applicationLayer.Payload())
		}
	}
}
