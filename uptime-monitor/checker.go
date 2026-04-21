package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkSites() {
	for {
		for i := range sites {
			response, err := http.Get(sites[i].URL)
			fmt.Printf("Checking %s \n", sites[i].URL)
			mu.Lock()
			if err != nil || response.StatusCode != http.StatusOK {
				sites[i].IsActive = false
				sites[i].LastChecked = time.Now()
			} else {
				sites[i].IsActive = true
				sites[i].LastChecked = time.Now()
			}
			mu.Unlock()

		}
		time.Sleep(10 * time.Second)
	}
}
