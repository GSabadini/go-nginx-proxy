package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", info)

	http.HandleFunc("/exp", func(w http.ResponseWriter, req *http.Request){
		w.Write([]byte("Hello world"))
	})

	log.Println("Start HTTP server :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func info(w http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()

	data := struct {
		Hostname string      `json:"hostname,omitempty"`
		IP       string      `json:"ip,omitempty"`
		Headers  http.Header `json:"headers,omitempty"`
		URL      string      `json:"url,omitempty"`
		Host     string      `json:"host,omitempty"`
		Method   string      `json:"method,omitempty"`
	}{
		Hostname: hostname,
		IP:       getIP(req),
		Headers:  req.Header,
		URL:      req.URL.RequestURI(),
		Host:     req.Host,
		Method:   req.Method,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getIP(req *http.Request) string {
	forwarded := req.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}

	return req.RemoteAddr
}