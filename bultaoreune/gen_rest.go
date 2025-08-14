package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func writeRest(domainResources []string, generateClientDirVer, fhirVersion string) {
	var sb strings.Builder
	//search is generic for now
	search := `// search for resource by its Sp<resource> search params, return bundle of matching resources
func Search[T SearchParam](sp T, fc *FhirClient) (*FHIR_VERSION.Bundle, error) {
	reqUrl, err := sp.SpEncode(&fc.BaseUrl)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, *reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating req, your fault not fhir server's %s", err)
	}

	resp, err := makeRequestCheckError(fc, req)
	if err != nil {
		return nil, fmt.Errorf("makeRequestCheckError %s", err)
	}

	var parsed FHIR_VERSION.Bundle
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, fmt.Errorf("error decoding json %s", err)
	}
	resp.Body.Close()
	return &parsed, nil
	}
	
	`
	searchFV := strings.ReplaceAll(search, "FHIR_VERSION", fhirVersion)
	sb.WriteString(searchFV)

	//searchGrouped is one generic func, but switch statement for each resource
	sb.WriteString(`func SearchGrouped[T SearchParam](sp T, fc *FhirClient) (*ResourceGroup, error) {
	bundle, err := Search(sp, fc)
	if err != nil {
		return nil, fmt.Errorf("Search error %s", err)
	}
	grp := ResourceGroup{}
	for _, e := range bundle.Entry {
		switch res := e.Resource.(type) {`)
	for _, dr := range domainResources {
		sb.WriteString("case *" + fhirVersion + "." + dr + ":\n")
		sb.WriteString("grp." + dr + "_list = append(grp." + dr + "_list, res)\n")
	}
	sb.WriteString(`default:
			return nil, errors.New("bundle entry not a domain resource, could not put in resourcegroup")
		}
	}
	return &grp, nil}`)

	//rest of crud one for each resource
	for _, dr := range domainResources {
		const templateStr = `
		// create DOMAIN_RESOURCE, return id of created DOMAIN_RESOURCE or OperationOutcome error
func (fc *FhirClient) CreateDOMAIN_RESOURCE(createResource *FHIR_VERSION.DOMAIN_RESOURCE) (*FHIR_VERSION.DOMAIN_RESOURCE, error) {
	if createResource == nil {
		return nil, errors.New("CreateDOMAIN_RESOURCE called with nil DOMAIN_RESOURCE")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DOMAIN_RESOURCE")
	if err != nil {
		return nil, err
	}
	reqBody, err := json.Marshal(createResource)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/fhir+json")

	resp, err := makeRequestCheckError(fc, req)
	if err != nil {
		return nil, fmt.Errorf("makeRequestCheckError %s", err)
	}

	var parsed FHIR_VERSION.DOMAIN_RESOURCE
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}
	
// read DOMAIN_RESOURCE from fhir server by id, return DOMAIN_RESOURCE or OperationOutcome error
func (fc *FhirClient) ReadDOMAIN_RESOURCE(id string) (*FHIR_VERSION.DOMAIN_RESOURCE, error) {
	if id == "" {
		return nil, errors.New("ReadDOMAIN_RESOURCE given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DOMAIN_RESOURCE", id)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := makeRequestCheckError(fc, req)
	if err != nil {
		return nil, fmt.Errorf("makeRequestCheckError %s", err)
	}

	var parsed FHIR_VERSION.DOMAIN_RESOURCE
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace DOMAIN_RESOURCE if exists on server, else create new DOMAIN_RESOURCE with given id, return DOMAIN_RESOURCE or OperationOutcome error
func (fc *FhirClient) UpdateDOMAIN_RESOURCE(updateResource *FHIR_VERSION.DOMAIN_RESOURCE) (*FHIR_VERSION.DOMAIN_RESOURCE, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateDOMAIN_RESOURCE called with nil DOMAIN_RESOURCE")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DOMAIN_RESOURCE", *updateResource.Id)
	if err != nil {
		return nil, err
	}
	reqBody, err := json.Marshal(updateResource)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, reqUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/fhir+json")

	resp, err := makeRequestCheckError(fc, req)
	if err != nil {
		return nil, fmt.Errorf("makeRequestCheckError %s", err)
	}

	var parsed FHIR_VERSION.DOMAIN_RESOURCE
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}


// add or replace fields in existing resource on server, return DOMAIN_RESOURCE or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDOMAIN_RESOURCE(patchResource *FHIR_VERSION.DOMAIN_RESOURCE) (*FHIR_VERSION.DOMAIN_RESOURCE, error) {
	if patchResource == nil {
		return nil, errors.New("PatchDOMAIN_RESOURCE given nil DOMAIN_RESOURCE")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchDOMAIN_RESOURCE given DOMAIN_RESOURCE without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DOMAIN_RESOURCE", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchDOMAIN_RESOURCE error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchDOMAIN_RESOURCE error using patch representation as json")
	}
	req, err := http.NewRequest(http.MethodPatch, reqUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json-patch+json")

	resp, err := makeRequestCheckError(fc, req)
	if err != nil {
		return nil, fmt.Errorf("makeRequestCheckError %s", err)
	}

	var parsed FHIR_VERSION.DOMAIN_RESOURCE
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}



// delete DOMAIN_RESOURCE and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDOMAIN_RESOURCE(deleteResource *FHIR_VERSION.DOMAIN_RESOURCE) (*FHIR_VERSION.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDOMAIN_RESOURCE given nil DOMAIN_RESOURCE")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDOMAIN_RESOURCE given DOMAIN_RESOURCE with nil Id (Id required to delete)")
	}
	return fc.DeleteDOMAIN_RESOURCEById(*deleteResource.Id)
}

// delete DOMAIN_RESOURCE with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDOMAIN_RESOURCEById(id string) (*FHIR_VERSION.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DOMAIN_RESOURCE", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := makeRequestCheckError(fc, req)
	if err != nil {
		return nil, fmt.Errorf("makeRequestCheckError %s", err)
	}

	var parsed FHIR_VERSION.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}`
		restDomainResorce := strings.ReplaceAll(templateStr, "DOMAIN_RESOURCE", dr)
		sb.WriteString(strings.ReplaceAll(restDomainResorce, "FHIR_VERSION", fhirVersion))
	}
	cf, _ := os.Create(path.Join(generateClientDirVer, "rest.go"))
	fmt.Fprintf(cf, "package %sClient\n", fhirVersion)
	fmt.Fprintln(cf, `
		import (
		"bytes"
		"encoding/json"
		"errors"
		"fmt"
		"io"
		"net/http"
		"net/url"
		"strings"
		"time"

		`+fhirVersion+" \"github.com/PotatoEMR/simple-fhir-client/"+fhirVersion+"\")\n")
	fmt.Fprintln(cf, sb.String())
	helpers := `
	// common checks for malformed response or operationoutcome response, rather than expected resource
func makeRequestCheckError(fc *FhirClient, req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", "application/fhir+json")
	resp, err := fc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request %s", err)
	}
	ct := resp.Header.Get("Content-Type")
	if ct == "text/html" {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %s received HTML response (expects json)\n%s", resp.Status, string(body))
	}
	if resp.StatusCode >= 300 {
		var opOutcome ` + fhirVersion + `.OperationOutcome
		err = json.NewDecoder(resp.Body).Decode(&opOutcome)
		if err != nil {
			if ct == "" {
				ct = "not provided in response header!"
			}
			return nil, fmt.Errorf("status %s, content type %s, could not parse OperationOutcome, error %s", resp.Status, ct, err.Error())
		}
		return nil, fmt.Errorf("status %s, %s", resp.Status, opOutcome.String())
	}
	return resp, nil
}
	
type PatchOperation struct {
	Op    string BACKTICKjson:"op"BACKTICK
	Path  string BACKTICKjson:"path"BACKTICK
	Value any    BACKTICKjson:"value"BACKTICK
}

// AI generated function (might be bad idk) for converting struct to a patch
// each field gets op/path/value:
//
//	{
//	  "op": "add",
//	  "path": "/name/0/text",
//	  "value": "Ash"
//	}
//
// used to PATCH with content type application/json-patch+json (not fhirpath)
func convertToPatch(data any) ([]PatchOperation, error) {
	var patches []PatchOperation

	// First convert to map[string]any to handle the JSON data
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var dataMap map[string]any
	err = json.Unmarshal(jsonData, &dataMap)
	if err != nil {
		return nil, err
	}

	// Recursively process the map
	processPatch("", dataMap, &patches)

	return patches, nil
}

func processPatch(basePath string, data any, patches *[]PatchOperation) {
	switch v := data.(type) {
	case map[string]any:
		for key, value := range v {
			path := basePath + "/" + key
			if value != nil {
				processPatch(path, value, patches)
			}
		}
	case []any:
		for i, value := range v {
			path := fmt.Sprintf("%s/%d", basePath, i)
			if value != nil {
				processPatch(path, value, patches)
			}
		}
	default:
		// This is a leaf value, create a patch operation
		*patches = append(*patches, PatchOperation{
			Op:    "merge",
			Path:  basePath,
			Value: v,
		})
	}
}`
	helperJsonBackticks := strings.ReplaceAll(helpers, "BACKTICK", "`")
	fmt.Fprintln(cf, helperJsonBackticks)
}
