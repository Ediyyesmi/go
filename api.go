package main

import (
	"encoding/json"
	"fmt"

	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Merhaba, dünya!")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Content-Type kontrolü
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Geçersiz içerik türü", http.StatusUnsupportedMediaType)
		return
	}

	var data map[string]interface{}
	// JSON verisini çözümle
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Geçersiz JSON", http.StatusBadRequest)
		return
	}

	// Yanıt olarak gelen veriyi döndür
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/post", postHandler) // POST isteği için yeni endpoint

	fmt.Println("Sunucu 8080 portunda çalışıyor...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Sunucu başlatılamadı:", err)
	}

}
