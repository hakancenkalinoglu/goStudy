package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type WebSite struct {
	URL         string    `json:"url"`
	IsActive    bool      `json:"isActive"`
	LastChecked time.Time `json:"lastChecked"`
}

func handleSites(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(sites)
		return
	}

	if r.Method == http.MethodPost {
		var newSite WebSite

		err := json.NewDecoder(r.Body).Decode(&newSite)
		if err != nil {
			http.Error(w, "Unsupported JSON Format", http.StatusBadRequest)
			return
		}
		mu.Lock()
		sites = append(sites, newSite)
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newSite)

		return
	}

	http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)

}

var sites []WebSite
var mu sync.Mutex

func main() {
	http.HandleFunc("/sites", handleSites)
	go checkSites()

	http.ListenAndServe(":8080", nil)

}
