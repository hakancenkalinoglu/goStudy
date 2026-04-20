package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID       int
	UserName string
	Role     string
}

// 1. Handler Home route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	//gelen isteklerin methodunu kontrol edebilirsin
	if r.Method != http.MethodGet {
		http.Error(w, "Sadece GET methodu destekleniyor", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("Home sayfasına bir istek geldi")
	w.Write([]byte("Web Sunucumuza Hoş Geldiniz!"))
}

// 2. api handler
func apiHandler(w http.ResponseWriter, r *http.Request) {
	u1 := User{
		ID:       1,
		UserName: "Ahmet",
		Role:     "Admin",
	}

	//client'a veri tipinin json olduğuınu belirtiyoruz
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//json halene Marshal edip doğrudan ResponseWriter'a yazıyoruz
	json.NewEncoder(w).Encode(u1)
}
func main() {
	//---- ROUTER ----
	//"Catch-All" (Her Şeyi Yakalayan / Joker)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("api/user", apiHandler)

	//START SERVER
	port := ":8080"
	fmt.Printf("Server %s postunda başlatılıyor..\n", port)

	//nil çünkü DefaultServeMux kullanılıyor
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Server başlatılamadı", err)
	}
}
