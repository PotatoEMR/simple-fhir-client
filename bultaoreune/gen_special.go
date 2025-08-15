package main

func PatientEverything(fhirVersion string) string {
	return `func (fc *FhirClient) PatientEverything(patId string) (*` + fhirVersion + `.Bundle, error) {
		reqUrl, err := url.JoinPath(fc.BaseUrl, "Patient", patId, "$everything")
		if err != nil {
			return nil, err
		}
		req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
		if err != nil {
			return nil, fmt.Errorf("error creating req, your fault not fhir server's %s", err)
		}

		resp, err := makeRequestCheckError(fc, req)
		if err != nil {
			return nil, fmt.Errorf("makeRequestCheckError %s", err)
		}

		var parsed r4.Bundle
		err = json.NewDecoder(resp.Body).Decode(&parsed)
		if err != nil {
			return nil, fmt.Errorf("error decoding json %s", err)
		}
		resp.Body.Close()
		return &parsed, nil
	}

	func (fc *FhirClient)PatientEverythingGrouped(patId string)(*ResourceGroup, error){
		bundle, err := fc.PatientEverything(patId)
		if err != nil {
			return nil, fmt.Errorf("PatientEverything error %s", err)
		}
		return BundleToGroup(bundle)
	}

`
}
