package r5Client

import (
	"net/http"
	"strings"
	"time"
)

// http client for fhir server, supports Create, Read, Update, Patch, Delete, Search.
// create with base url eg client.New("https://r4.smarthealthit.org")
type FhirClient struct {
	BaseUrl string
	*http.Client
}

// http client for fhir server, supports Create, Read, Update, Patch, Delete, Search.
// create with base url eg client.New("https://r4.smarthealthit.org")
func New(FhirServer_BaseUrl string) *FhirClient {
	if !strings.HasPrefix(FhirServer_BaseUrl, "https://") {
		FhirServer_BaseUrl = "https://" + FhirServer_BaseUrl
	}
	return &FhirClient{
		BaseUrl: FhirServer_BaseUrl,
		Client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}
