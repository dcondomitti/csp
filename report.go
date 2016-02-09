package main

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
