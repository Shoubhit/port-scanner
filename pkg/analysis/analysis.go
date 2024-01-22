package analysis

// AnalysisResponse represents the network security analysis report response.
type AnalysisResponse struct {
	IPAddress   string `json:"ipAddress"`
	PortScan    string `json:"portScan"`
	Protocol    string `json:"protocol"`
	DNSAnalysis string `json:"dnsAnalysis"`
}
