package main

func PatientEverything(fhirVersion string) string {
	return `// Search for everything about a patient, returned as bundle
//
// PatientEverythingGroup does exact same search, but returns list of each resource type
func (fc *FhirClient) PatientEverythingBundled(patId string) (*` + fhirVersion + `.Bundle, error) {
	if patId == "" {
		return nil, errors.New("PatientEverything called with empty pat id")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Patient", patId, "$everything")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating req, your fault not fhir server's %s", err)
	}

	var parsed ` + fhirVersion + `.Bundle
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
	    if _, ok := err.(*r4.OperationOutcome); !ok {
	        err = fmt.Errorf("PatientEverythingBundled: %s", err)
	    }
	}
	return &parsed, err
}

// Search for everything about a patient, returns a list of each resource type
//
// PatientEverythingBundled does exact same search, but returns bundle of all resources
func (fc *FhirClient) PatientEverythingGrouped(patId string) (*ResourceGroup, error) {
	bundle, err := fc.PatientEverythingBundled(patId)
	if err != nil {
		return nil, err
	}
	return BundleToGroup(bundle)
}`
}
