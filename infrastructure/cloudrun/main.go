package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	reports "serverless-functions-go/infrastructure/cloudrun/reports"
)

func main() {
	http.HandleFunc("/google/reports/ad", reports.GetAdPerformanceReport)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
