package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type Report struct {
	CspReport `json:"csp-report"`
}

type CspReport struct {
	DocumentUri       string `json:"document-uri"`
	Referrer          string `json:"referrer"`
	ViolatedDirective string `json:"violated-directive"`
	OriginalPolicy    string `json:"original-policy"`
	BlockedUri        string `json:"blocked-uri"`
	SourceFile        string `json:"source-file"`
	LineNumber        int    `json:"line-number"`
	ColumnNumber      int    `json:"column-number"`
	StatusCode        int    `json:"status-code"`
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
