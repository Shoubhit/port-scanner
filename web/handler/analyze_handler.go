package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shoubhit/secure-api/pkg/analysis"
	"github.com/Shoubhit/secure-api/pkg/database"
)

type RequestParams struct {
	IPAddress string `json:"ipAddress"`
}

func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	// Extract parameters from the request
	var params RequestParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: Invalid request parameters")
		return
	}

	// Simulate network security analysis
	rawScanResult := analysis.InitialScan(params.IPAddress)
	portScanResult := ""
	for _, sr := range rawScanResult {
		portScanResult += fmt.Sprintf("{Port: %s, State: %s, Service: %s}\n", sr.Port, sr.State, sr.Service)
	}
	protocolAnalysisResult := "Protocol analysis result: Insecure protocol detected - HTTP/1.0"
	dnsAnalysisResult := "DNS analysis result: No signs of malicious activities"

	// Save the analysis report in the database with a timestamp
	report := database.Report{
		IPAddress:   params.IPAddress,
		PortScan:    portScanResult,
		Protocol:    protocolAnalysisResult,
		DNSAnalysis: dnsAnalysisResult,
	}
	if err := database.AddReport(report); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error saving report: %v", err)
		return
	}

	response := analysis.AnalysisResponse{
		IPAddress:   params.IPAddress,
		PortScan:    portScanResult,
		Protocol:    protocolAnalysisResult,
		DNSAnalysis: dnsAnalysisResult,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
