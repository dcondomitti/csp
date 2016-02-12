package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
)

var revision string

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.WithFields(log.Fields{
			"error": "no_hostname",
		}).Fatal("Unable to determine hostname")
	}

	http.HandleFunc("/csp-report", ReportHandler)

	http.HandleFunc("/health", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "OK")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "csp report handler rev %s on %s", revision, hostname)
	})

	log.WithFields(log.Fields{
		"event":    "startup",
		"port":     port,
		"hostname": hostname,
		"revision": revision,
	}).Info("Starting up")

	http.ListenAndServe(port, nil)
}
