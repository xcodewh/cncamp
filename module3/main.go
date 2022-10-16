package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	_ "net/http/pprof"
)

func main() {
	fmt.Println("Starting http server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/readHeaders", readHeaders)
	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering healthz handler")
	//Log IP, http status code
	log.Printf("Request IP: %s, HTTP status code: %d", r.RemoteAddr, http.StatusOK)
	io.WriteString(w, "200\n")
}

func readHeaders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering readHeaders handler")
	//Log IP, http status code
	log.Printf("Request IP: %s, HTTP status code: %d", r.RemoteAddr, http.StatusOK)
	resp := make(map[string]string)
	for k, v := range r.Header {
		joinedVal := strings.Join(v, ",")
		w.Header().Set(k, joinedVal)
		resp[k] = joinedVal
	}

	//Read enviroment variable VERSION and write into response header
	w.Header().Set("VERSION", os.Getenv("VERSION"))
	jsonResp, _ := json.Marshal(resp)
	io.WriteString(w, string(jsonResp))
}
