package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	port = flag.String("port", ":8080", "port to listen on")
)

func main() {
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			hostname := os.Getenv("HOSTNAME")
			fmt.Fprintf(w, "Hostname: %s\n", hostname)
			return
		}
	})
	log.Fatal(http.ListenAndServe(*port, mux))
}
