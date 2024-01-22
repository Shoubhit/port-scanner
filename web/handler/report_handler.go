package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shoubhit/secure-api/pkg/database"
)

func GetReports(w http.ResponseWriter, r *http.Request) {
	reports, err := database.GetAllReports()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}
