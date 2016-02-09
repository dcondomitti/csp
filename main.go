package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func ReportHandler(rw http.ResponseWriter, req *http.Request) {
	var r Report

	if req.Method != "POST" {
		log.WithFields(log.Fields{
			"event":  "request_rejected",
			"error":  "method_not_allowed",
			"method": req.Method,
		}).Info("Method not allowed")

		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.WithFields(log.Fields{
			"event": "request_rejected",
			"error": err,
		}).Error("Unable to read request body")

		http.Error(rw, "Unable to parse JSON", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(bytes.NewBuffer(body)).Decode(&r); err != nil {
		log.WithFields(log.Fields{
			"event": "request_rejected",
			"error": err,
			"body":  body,
		}).Error("Unable to parse JSON")

		http.Error(rw, "Unable to parse JSON", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusAccepted)

	log.WithFields(log.Fields{
		"event":  "request_accepted",
		"report": r,
	}).Info("Request accepted")

	return

}

func HealthHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "OK")
}
