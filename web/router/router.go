package router

import (
	"net/http"

	"github.com/Shoubhit/secure-api/web/handler"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Handlers
	analyzeHandler := http.HandlerFunc(handler.AnalyzeHandler)
	reportHandler := http.HandlerFunc(handler.GetReports)

	// Routes
	r.Handle("/analyze", analyzeHandler).Methods("POST")
	r.HandleFunc("/reports", reportHandler).Methods("GET")

	return r
}
