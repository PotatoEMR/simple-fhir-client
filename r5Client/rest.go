package r5Client

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

	r5 "github.com/PotatoEMR/simple-fhir-client/r5"
)

// search for resource by its Sp<resource> search params, return bundle of matching resources
//
// SearchGrouped does exact same search, but returns list of each resource type
func (fc *FhirClient) SearchBundled(sp SearchParam) (*r5.Bundle, error) {
	reqUrl, err := sp.SpEncode(&fc.BaseUrl)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, *reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("searchBundle error creating req, your fault not fhir server's %s", err)
	}

	resp, err := makeRequestCheckError(fc, req)
	if err != nil {
		return nil, fmt.Errorf("searchBundle makeRequestCheckError %s", err)
	}

	var parsed r5.Bundle
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, fmt.Errorf("searchBundle error decoding json %s", err)
	}
	resp.Body.Close()
	return &parsed, nil
}

// search for resource by its Sp<resource> search params, return bundle of matching resources
//
// SearchBundle does exact same search, but returns bundle of all resources
func (fc *FhirClient) SearchGrouped(sp SearchParam) (*ResourceGroup, error) {
	bundle, err := fc.SearchBundled(sp)
	if err != nil {
		return nil, fmt.Errorf("search error %s", err)
	}
	return BundleToGroup(bundle)
}

// Search for everything about a patient, returned as bundle
//
// PatientEverythingGroup does exact same search, but returns list of each resource type
func (fc *FhirClient) PatientEverythingBundled(patId string) (*r5.Bundle, error) {
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

	var parsed r5.Bundle
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, fmt.Errorf("error decoding json %s", err)
	}
	resp.Body.Close()
	return &parsed, nil
}

// Search for everything about a patient, returns a list of each resource type
//
// PatientEverythingBundled does exact same search, but returns bundle of all resources
func (fc *FhirClient) PatientEverythingGrouped(patId string) (*ResourceGroup, error) {
	bundle, err := fc.PatientEverythingBundled(patId)
	if err != nil {
		return nil, fmt.Errorf("patientEverything error %s", err)
	}
	return BundleToGroup(bundle)
}

// create Account, return id of created Account or OperationOutcome error
func (fc *FhirClient) CreateAccount(createResource *r5.Account) (*r5.Account, error) {
	if createResource == nil {
		return nil, errors.New("CreateAccount called with nil Account")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Account")
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

	var parsed r5.Account
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Account from fhir server by id, return Account or OperationOutcome error
func (fc *FhirClient) ReadAccount(id string) (*r5.Account, error) {
	if id == "" {
		return nil, errors.New("ReadAccount given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Account", id)
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

	var parsed r5.Account
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Account if exists on server, else create new Account with given id, return Account or OperationOutcome error
func (fc *FhirClient) UpdateAccount(updateResource *r5.Account) (*r5.Account, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateAccount called with nil Account")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Account", *updateResource.Id)
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

	var parsed r5.Account
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Account or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchAccount(patchResource *r5.Account) (*r5.Account, error) {
	if patchResource == nil {
		return nil, errors.New("PatchAccount given nil Account")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchAccount given Account without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Account", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchAccount error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchAccount error using patch representation as json")
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

	var parsed r5.Account
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Account and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAccount(deleteResource *r5.Account) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteAccount given nil Account")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteAccount given Account with nil Id (Id required to delete)")
	}
	return fc.DeleteAccountById(*deleteResource.Id)
}

// delete Account with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAccountById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Account", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ActivityDefinition, return id of created ActivityDefinition or OperationOutcome error
func (fc *FhirClient) CreateActivityDefinition(createResource *r5.ActivityDefinition) (*r5.ActivityDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateActivityDefinition called with nil ActivityDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ActivityDefinition")
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

	var parsed r5.ActivityDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ActivityDefinition from fhir server by id, return ActivityDefinition or OperationOutcome error
func (fc *FhirClient) ReadActivityDefinition(id string) (*r5.ActivityDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadActivityDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ActivityDefinition", id)
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

	var parsed r5.ActivityDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ActivityDefinition if exists on server, else create new ActivityDefinition with given id, return ActivityDefinition or OperationOutcome error
func (fc *FhirClient) UpdateActivityDefinition(updateResource *r5.ActivityDefinition) (*r5.ActivityDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateActivityDefinition called with nil ActivityDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ActivityDefinition", *updateResource.Id)
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

	var parsed r5.ActivityDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ActivityDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchActivityDefinition(patchResource *r5.ActivityDefinition) (*r5.ActivityDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchActivityDefinition given nil ActivityDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchActivityDefinition given ActivityDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ActivityDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchActivityDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchActivityDefinition error using patch representation as json")
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

	var parsed r5.ActivityDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ActivityDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteActivityDefinition(deleteResource *r5.ActivityDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteActivityDefinition given nil ActivityDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteActivityDefinition given ActivityDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteActivityDefinitionById(*deleteResource.Id)
}

// delete ActivityDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteActivityDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ActivityDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ActorDefinition, return id of created ActorDefinition or OperationOutcome error
func (fc *FhirClient) CreateActorDefinition(createResource *r5.ActorDefinition) (*r5.ActorDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateActorDefinition called with nil ActorDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ActorDefinition")
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

	var parsed r5.ActorDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ActorDefinition from fhir server by id, return ActorDefinition or OperationOutcome error
func (fc *FhirClient) ReadActorDefinition(id string) (*r5.ActorDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadActorDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ActorDefinition", id)
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

	var parsed r5.ActorDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ActorDefinition if exists on server, else create new ActorDefinition with given id, return ActorDefinition or OperationOutcome error
func (fc *FhirClient) UpdateActorDefinition(updateResource *r5.ActorDefinition) (*r5.ActorDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateActorDefinition called with nil ActorDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ActorDefinition", *updateResource.Id)
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

	var parsed r5.ActorDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ActorDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchActorDefinition(patchResource *r5.ActorDefinition) (*r5.ActorDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchActorDefinition given nil ActorDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchActorDefinition given ActorDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ActorDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchActorDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchActorDefinition error using patch representation as json")
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

	var parsed r5.ActorDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ActorDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteActorDefinition(deleteResource *r5.ActorDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteActorDefinition given nil ActorDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteActorDefinition given ActorDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteActorDefinitionById(*deleteResource.Id)
}

// delete ActorDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteActorDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ActorDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create AdministrableProductDefinition, return id of created AdministrableProductDefinition or OperationOutcome error
func (fc *FhirClient) CreateAdministrableProductDefinition(createResource *r5.AdministrableProductDefinition) (*r5.AdministrableProductDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateAdministrableProductDefinition called with nil AdministrableProductDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AdministrableProductDefinition")
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

	var parsed r5.AdministrableProductDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read AdministrableProductDefinition from fhir server by id, return AdministrableProductDefinition or OperationOutcome error
func (fc *FhirClient) ReadAdministrableProductDefinition(id string) (*r5.AdministrableProductDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadAdministrableProductDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AdministrableProductDefinition", id)
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

	var parsed r5.AdministrableProductDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace AdministrableProductDefinition if exists on server, else create new AdministrableProductDefinition with given id, return AdministrableProductDefinition or OperationOutcome error
func (fc *FhirClient) UpdateAdministrableProductDefinition(updateResource *r5.AdministrableProductDefinition) (*r5.AdministrableProductDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateAdministrableProductDefinition called with nil AdministrableProductDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AdministrableProductDefinition", *updateResource.Id)
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

	var parsed r5.AdministrableProductDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return AdministrableProductDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchAdministrableProductDefinition(patchResource *r5.AdministrableProductDefinition) (*r5.AdministrableProductDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchAdministrableProductDefinition given nil AdministrableProductDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchAdministrableProductDefinition given AdministrableProductDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AdministrableProductDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchAdministrableProductDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchAdministrableProductDefinition error using patch representation as json")
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

	var parsed r5.AdministrableProductDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete AdministrableProductDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAdministrableProductDefinition(deleteResource *r5.AdministrableProductDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteAdministrableProductDefinition given nil AdministrableProductDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteAdministrableProductDefinition given AdministrableProductDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteAdministrableProductDefinitionById(*deleteResource.Id)
}

// delete AdministrableProductDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAdministrableProductDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AdministrableProductDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create AdverseEvent, return id of created AdverseEvent or OperationOutcome error
func (fc *FhirClient) CreateAdverseEvent(createResource *r5.AdverseEvent) (*r5.AdverseEvent, error) {
	if createResource == nil {
		return nil, errors.New("CreateAdverseEvent called with nil AdverseEvent")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AdverseEvent")
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

	var parsed r5.AdverseEvent
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read AdverseEvent from fhir server by id, return AdverseEvent or OperationOutcome error
func (fc *FhirClient) ReadAdverseEvent(id string) (*r5.AdverseEvent, error) {
	if id == "" {
		return nil, errors.New("ReadAdverseEvent given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AdverseEvent", id)
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

	var parsed r5.AdverseEvent
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace AdverseEvent if exists on server, else create new AdverseEvent with given id, return AdverseEvent or OperationOutcome error
func (fc *FhirClient) UpdateAdverseEvent(updateResource *r5.AdverseEvent) (*r5.AdverseEvent, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateAdverseEvent called with nil AdverseEvent")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AdverseEvent", *updateResource.Id)
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

	var parsed r5.AdverseEvent
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return AdverseEvent or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchAdverseEvent(patchResource *r5.AdverseEvent) (*r5.AdverseEvent, error) {
	if patchResource == nil {
		return nil, errors.New("PatchAdverseEvent given nil AdverseEvent")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchAdverseEvent given AdverseEvent without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AdverseEvent", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchAdverseEvent error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchAdverseEvent error using patch representation as json")
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

	var parsed r5.AdverseEvent
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete AdverseEvent and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAdverseEvent(deleteResource *r5.AdverseEvent) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteAdverseEvent given nil AdverseEvent")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteAdverseEvent given AdverseEvent with nil Id (Id required to delete)")
	}
	return fc.DeleteAdverseEventById(*deleteResource.Id)
}

// delete AdverseEvent with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAdverseEventById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AdverseEvent", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create AllergyIntolerance, return id of created AllergyIntolerance or OperationOutcome error
func (fc *FhirClient) CreateAllergyIntolerance(createResource *r5.AllergyIntolerance) (*r5.AllergyIntolerance, error) {
	if createResource == nil {
		return nil, errors.New("CreateAllergyIntolerance called with nil AllergyIntolerance")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AllergyIntolerance")
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

	var parsed r5.AllergyIntolerance
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read AllergyIntolerance from fhir server by id, return AllergyIntolerance or OperationOutcome error
func (fc *FhirClient) ReadAllergyIntolerance(id string) (*r5.AllergyIntolerance, error) {
	if id == "" {
		return nil, errors.New("ReadAllergyIntolerance given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AllergyIntolerance", id)
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

	var parsed r5.AllergyIntolerance
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace AllergyIntolerance if exists on server, else create new AllergyIntolerance with given id, return AllergyIntolerance or OperationOutcome error
func (fc *FhirClient) UpdateAllergyIntolerance(updateResource *r5.AllergyIntolerance) (*r5.AllergyIntolerance, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateAllergyIntolerance called with nil AllergyIntolerance")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AllergyIntolerance", *updateResource.Id)
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

	var parsed r5.AllergyIntolerance
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return AllergyIntolerance or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchAllergyIntolerance(patchResource *r5.AllergyIntolerance) (*r5.AllergyIntolerance, error) {
	if patchResource == nil {
		return nil, errors.New("PatchAllergyIntolerance given nil AllergyIntolerance")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchAllergyIntolerance given AllergyIntolerance without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AllergyIntolerance", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchAllergyIntolerance error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchAllergyIntolerance error using patch representation as json")
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

	var parsed r5.AllergyIntolerance
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete AllergyIntolerance and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAllergyIntolerance(deleteResource *r5.AllergyIntolerance) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteAllergyIntolerance given nil AllergyIntolerance")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteAllergyIntolerance given AllergyIntolerance with nil Id (Id required to delete)")
	}
	return fc.DeleteAllergyIntoleranceById(*deleteResource.Id)
}

// delete AllergyIntolerance with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAllergyIntoleranceById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AllergyIntolerance", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Appointment, return id of created Appointment or OperationOutcome error
func (fc *FhirClient) CreateAppointment(createResource *r5.Appointment) (*r5.Appointment, error) {
	if createResource == nil {
		return nil, errors.New("CreateAppointment called with nil Appointment")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Appointment")
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

	var parsed r5.Appointment
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Appointment from fhir server by id, return Appointment or OperationOutcome error
func (fc *FhirClient) ReadAppointment(id string) (*r5.Appointment, error) {
	if id == "" {
		return nil, errors.New("ReadAppointment given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Appointment", id)
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

	var parsed r5.Appointment
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Appointment if exists on server, else create new Appointment with given id, return Appointment or OperationOutcome error
func (fc *FhirClient) UpdateAppointment(updateResource *r5.Appointment) (*r5.Appointment, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateAppointment called with nil Appointment")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Appointment", *updateResource.Id)
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

	var parsed r5.Appointment
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Appointment or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchAppointment(patchResource *r5.Appointment) (*r5.Appointment, error) {
	if patchResource == nil {
		return nil, errors.New("PatchAppointment given nil Appointment")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchAppointment given Appointment without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Appointment", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchAppointment error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchAppointment error using patch representation as json")
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

	var parsed r5.Appointment
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Appointment and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAppointment(deleteResource *r5.Appointment) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteAppointment given nil Appointment")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteAppointment given Appointment with nil Id (Id required to delete)")
	}
	return fc.DeleteAppointmentById(*deleteResource.Id)
}

// delete Appointment with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAppointmentById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Appointment", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create AppointmentResponse, return id of created AppointmentResponse or OperationOutcome error
func (fc *FhirClient) CreateAppointmentResponse(createResource *r5.AppointmentResponse) (*r5.AppointmentResponse, error) {
	if createResource == nil {
		return nil, errors.New("CreateAppointmentResponse called with nil AppointmentResponse")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AppointmentResponse")
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

	var parsed r5.AppointmentResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read AppointmentResponse from fhir server by id, return AppointmentResponse or OperationOutcome error
func (fc *FhirClient) ReadAppointmentResponse(id string) (*r5.AppointmentResponse, error) {
	if id == "" {
		return nil, errors.New("ReadAppointmentResponse given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AppointmentResponse", id)
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

	var parsed r5.AppointmentResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace AppointmentResponse if exists on server, else create new AppointmentResponse with given id, return AppointmentResponse or OperationOutcome error
func (fc *FhirClient) UpdateAppointmentResponse(updateResource *r5.AppointmentResponse) (*r5.AppointmentResponse, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateAppointmentResponse called with nil AppointmentResponse")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AppointmentResponse", *updateResource.Id)
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

	var parsed r5.AppointmentResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return AppointmentResponse or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchAppointmentResponse(patchResource *r5.AppointmentResponse) (*r5.AppointmentResponse, error) {
	if patchResource == nil {
		return nil, errors.New("PatchAppointmentResponse given nil AppointmentResponse")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchAppointmentResponse given AppointmentResponse without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AppointmentResponse", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchAppointmentResponse error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchAppointmentResponse error using patch representation as json")
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

	var parsed r5.AppointmentResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete AppointmentResponse and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAppointmentResponse(deleteResource *r5.AppointmentResponse) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteAppointmentResponse given nil AppointmentResponse")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteAppointmentResponse given AppointmentResponse with nil Id (Id required to delete)")
	}
	return fc.DeleteAppointmentResponseById(*deleteResource.Id)
}

// delete AppointmentResponse with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAppointmentResponseById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AppointmentResponse", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ArtifactAssessment, return id of created ArtifactAssessment or OperationOutcome error
func (fc *FhirClient) CreateArtifactAssessment(createResource *r5.ArtifactAssessment) (*r5.ArtifactAssessment, error) {
	if createResource == nil {
		return nil, errors.New("CreateArtifactAssessment called with nil ArtifactAssessment")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ArtifactAssessment")
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

	var parsed r5.ArtifactAssessment
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ArtifactAssessment from fhir server by id, return ArtifactAssessment or OperationOutcome error
func (fc *FhirClient) ReadArtifactAssessment(id string) (*r5.ArtifactAssessment, error) {
	if id == "" {
		return nil, errors.New("ReadArtifactAssessment given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ArtifactAssessment", id)
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

	var parsed r5.ArtifactAssessment
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ArtifactAssessment if exists on server, else create new ArtifactAssessment with given id, return ArtifactAssessment or OperationOutcome error
func (fc *FhirClient) UpdateArtifactAssessment(updateResource *r5.ArtifactAssessment) (*r5.ArtifactAssessment, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateArtifactAssessment called with nil ArtifactAssessment")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ArtifactAssessment", *updateResource.Id)
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

	var parsed r5.ArtifactAssessment
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ArtifactAssessment or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchArtifactAssessment(patchResource *r5.ArtifactAssessment) (*r5.ArtifactAssessment, error) {
	if patchResource == nil {
		return nil, errors.New("PatchArtifactAssessment given nil ArtifactAssessment")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchArtifactAssessment given ArtifactAssessment without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ArtifactAssessment", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchArtifactAssessment error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchArtifactAssessment error using patch representation as json")
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

	var parsed r5.ArtifactAssessment
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ArtifactAssessment and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteArtifactAssessment(deleteResource *r5.ArtifactAssessment) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteArtifactAssessment given nil ArtifactAssessment")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteArtifactAssessment given ArtifactAssessment with nil Id (Id required to delete)")
	}
	return fc.DeleteArtifactAssessmentById(*deleteResource.Id)
}

// delete ArtifactAssessment with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteArtifactAssessmentById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ArtifactAssessment", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create AuditEvent, return id of created AuditEvent or OperationOutcome error
func (fc *FhirClient) CreateAuditEvent(createResource *r5.AuditEvent) (*r5.AuditEvent, error) {
	if createResource == nil {
		return nil, errors.New("CreateAuditEvent called with nil AuditEvent")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AuditEvent")
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

	var parsed r5.AuditEvent
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read AuditEvent from fhir server by id, return AuditEvent or OperationOutcome error
func (fc *FhirClient) ReadAuditEvent(id string) (*r5.AuditEvent, error) {
	if id == "" {
		return nil, errors.New("ReadAuditEvent given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AuditEvent", id)
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

	var parsed r5.AuditEvent
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace AuditEvent if exists on server, else create new AuditEvent with given id, return AuditEvent or OperationOutcome error
func (fc *FhirClient) UpdateAuditEvent(updateResource *r5.AuditEvent) (*r5.AuditEvent, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateAuditEvent called with nil AuditEvent")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AuditEvent", *updateResource.Id)
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

	var parsed r5.AuditEvent
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return AuditEvent or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchAuditEvent(patchResource *r5.AuditEvent) (*r5.AuditEvent, error) {
	if patchResource == nil {
		return nil, errors.New("PatchAuditEvent given nil AuditEvent")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchAuditEvent given AuditEvent without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AuditEvent", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchAuditEvent error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchAuditEvent error using patch representation as json")
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

	var parsed r5.AuditEvent
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete AuditEvent and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAuditEvent(deleteResource *r5.AuditEvent) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteAuditEvent given nil AuditEvent")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteAuditEvent given AuditEvent with nil Id (Id required to delete)")
	}
	return fc.DeleteAuditEventById(*deleteResource.Id)
}

// delete AuditEvent with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAuditEventById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AuditEvent", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Basic, return id of created Basic or OperationOutcome error
func (fc *FhirClient) CreateBasic(createResource *r5.Basic) (*r5.Basic, error) {
	if createResource == nil {
		return nil, errors.New("CreateBasic called with nil Basic")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Basic")
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

	var parsed r5.Basic
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Basic from fhir server by id, return Basic or OperationOutcome error
func (fc *FhirClient) ReadBasic(id string) (*r5.Basic, error) {
	if id == "" {
		return nil, errors.New("ReadBasic given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Basic", id)
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

	var parsed r5.Basic
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Basic if exists on server, else create new Basic with given id, return Basic or OperationOutcome error
func (fc *FhirClient) UpdateBasic(updateResource *r5.Basic) (*r5.Basic, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateBasic called with nil Basic")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Basic", *updateResource.Id)
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

	var parsed r5.Basic
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Basic or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchBasic(patchResource *r5.Basic) (*r5.Basic, error) {
	if patchResource == nil {
		return nil, errors.New("PatchBasic given nil Basic")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchBasic given Basic without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Basic", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchBasic error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchBasic error using patch representation as json")
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

	var parsed r5.Basic
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Basic and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBasic(deleteResource *r5.Basic) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteBasic given nil Basic")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteBasic given Basic with nil Id (Id required to delete)")
	}
	return fc.DeleteBasicById(*deleteResource.Id)
}

// delete Basic with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBasicById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Basic", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Binary, return id of created Binary or OperationOutcome error
func (fc *FhirClient) CreateBinary(createResource *r5.Binary) (*r5.Binary, error) {
	if createResource == nil {
		return nil, errors.New("CreateBinary called with nil Binary")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Binary")
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

	var parsed r5.Binary
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Binary from fhir server by id, return Binary or OperationOutcome error
func (fc *FhirClient) ReadBinary(id string) (*r5.Binary, error) {
	if id == "" {
		return nil, errors.New("ReadBinary given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Binary", id)
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

	var parsed r5.Binary
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Binary if exists on server, else create new Binary with given id, return Binary or OperationOutcome error
func (fc *FhirClient) UpdateBinary(updateResource *r5.Binary) (*r5.Binary, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateBinary called with nil Binary")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Binary", *updateResource.Id)
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

	var parsed r5.Binary
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Binary or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchBinary(patchResource *r5.Binary) (*r5.Binary, error) {
	if patchResource == nil {
		return nil, errors.New("PatchBinary given nil Binary")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchBinary given Binary without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Binary", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchBinary error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchBinary error using patch representation as json")
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

	var parsed r5.Binary
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Binary and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBinary(deleteResource *r5.Binary) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteBinary given nil Binary")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteBinary given Binary with nil Id (Id required to delete)")
	}
	return fc.DeleteBinaryById(*deleteResource.Id)
}

// delete Binary with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBinaryById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Binary", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create BiologicallyDerivedProduct, return id of created BiologicallyDerivedProduct or OperationOutcome error
func (fc *FhirClient) CreateBiologicallyDerivedProduct(createResource *r5.BiologicallyDerivedProduct) (*r5.BiologicallyDerivedProduct, error) {
	if createResource == nil {
		return nil, errors.New("CreateBiologicallyDerivedProduct called with nil BiologicallyDerivedProduct")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BiologicallyDerivedProduct")
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

	var parsed r5.BiologicallyDerivedProduct
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read BiologicallyDerivedProduct from fhir server by id, return BiologicallyDerivedProduct or OperationOutcome error
func (fc *FhirClient) ReadBiologicallyDerivedProduct(id string) (*r5.BiologicallyDerivedProduct, error) {
	if id == "" {
		return nil, errors.New("ReadBiologicallyDerivedProduct given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BiologicallyDerivedProduct", id)
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

	var parsed r5.BiologicallyDerivedProduct
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace BiologicallyDerivedProduct if exists on server, else create new BiologicallyDerivedProduct with given id, return BiologicallyDerivedProduct or OperationOutcome error
func (fc *FhirClient) UpdateBiologicallyDerivedProduct(updateResource *r5.BiologicallyDerivedProduct) (*r5.BiologicallyDerivedProduct, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateBiologicallyDerivedProduct called with nil BiologicallyDerivedProduct")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BiologicallyDerivedProduct", *updateResource.Id)
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

	var parsed r5.BiologicallyDerivedProduct
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return BiologicallyDerivedProduct or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchBiologicallyDerivedProduct(patchResource *r5.BiologicallyDerivedProduct) (*r5.BiologicallyDerivedProduct, error) {
	if patchResource == nil {
		return nil, errors.New("PatchBiologicallyDerivedProduct given nil BiologicallyDerivedProduct")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchBiologicallyDerivedProduct given BiologicallyDerivedProduct without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BiologicallyDerivedProduct", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchBiologicallyDerivedProduct error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchBiologicallyDerivedProduct error using patch representation as json")
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

	var parsed r5.BiologicallyDerivedProduct
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete BiologicallyDerivedProduct and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBiologicallyDerivedProduct(deleteResource *r5.BiologicallyDerivedProduct) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteBiologicallyDerivedProduct given nil BiologicallyDerivedProduct")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteBiologicallyDerivedProduct given BiologicallyDerivedProduct with nil Id (Id required to delete)")
	}
	return fc.DeleteBiologicallyDerivedProductById(*deleteResource.Id)
}

// delete BiologicallyDerivedProduct with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBiologicallyDerivedProductById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BiologicallyDerivedProduct", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create BiologicallyDerivedProductDispense, return id of created BiologicallyDerivedProductDispense or OperationOutcome error
func (fc *FhirClient) CreateBiologicallyDerivedProductDispense(createResource *r5.BiologicallyDerivedProductDispense) (*r5.BiologicallyDerivedProductDispense, error) {
	if createResource == nil {
		return nil, errors.New("CreateBiologicallyDerivedProductDispense called with nil BiologicallyDerivedProductDispense")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BiologicallyDerivedProductDispense")
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

	var parsed r5.BiologicallyDerivedProductDispense
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read BiologicallyDerivedProductDispense from fhir server by id, return BiologicallyDerivedProductDispense or OperationOutcome error
func (fc *FhirClient) ReadBiologicallyDerivedProductDispense(id string) (*r5.BiologicallyDerivedProductDispense, error) {
	if id == "" {
		return nil, errors.New("ReadBiologicallyDerivedProductDispense given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BiologicallyDerivedProductDispense", id)
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

	var parsed r5.BiologicallyDerivedProductDispense
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace BiologicallyDerivedProductDispense if exists on server, else create new BiologicallyDerivedProductDispense with given id, return BiologicallyDerivedProductDispense or OperationOutcome error
func (fc *FhirClient) UpdateBiologicallyDerivedProductDispense(updateResource *r5.BiologicallyDerivedProductDispense) (*r5.BiologicallyDerivedProductDispense, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateBiologicallyDerivedProductDispense called with nil BiologicallyDerivedProductDispense")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BiologicallyDerivedProductDispense", *updateResource.Id)
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

	var parsed r5.BiologicallyDerivedProductDispense
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return BiologicallyDerivedProductDispense or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchBiologicallyDerivedProductDispense(patchResource *r5.BiologicallyDerivedProductDispense) (*r5.BiologicallyDerivedProductDispense, error) {
	if patchResource == nil {
		return nil, errors.New("PatchBiologicallyDerivedProductDispense given nil BiologicallyDerivedProductDispense")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchBiologicallyDerivedProductDispense given BiologicallyDerivedProductDispense without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BiologicallyDerivedProductDispense", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchBiologicallyDerivedProductDispense error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchBiologicallyDerivedProductDispense error using patch representation as json")
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

	var parsed r5.BiologicallyDerivedProductDispense
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete BiologicallyDerivedProductDispense and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBiologicallyDerivedProductDispense(deleteResource *r5.BiologicallyDerivedProductDispense) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteBiologicallyDerivedProductDispense given nil BiologicallyDerivedProductDispense")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteBiologicallyDerivedProductDispense given BiologicallyDerivedProductDispense with nil Id (Id required to delete)")
	}
	return fc.DeleteBiologicallyDerivedProductDispenseById(*deleteResource.Id)
}

// delete BiologicallyDerivedProductDispense with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBiologicallyDerivedProductDispenseById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BiologicallyDerivedProductDispense", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create BodyStructure, return id of created BodyStructure or OperationOutcome error
func (fc *FhirClient) CreateBodyStructure(createResource *r5.BodyStructure) (*r5.BodyStructure, error) {
	if createResource == nil {
		return nil, errors.New("CreateBodyStructure called with nil BodyStructure")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BodyStructure")
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

	var parsed r5.BodyStructure
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read BodyStructure from fhir server by id, return BodyStructure or OperationOutcome error
func (fc *FhirClient) ReadBodyStructure(id string) (*r5.BodyStructure, error) {
	if id == "" {
		return nil, errors.New("ReadBodyStructure given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BodyStructure", id)
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

	var parsed r5.BodyStructure
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace BodyStructure if exists on server, else create new BodyStructure with given id, return BodyStructure or OperationOutcome error
func (fc *FhirClient) UpdateBodyStructure(updateResource *r5.BodyStructure) (*r5.BodyStructure, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateBodyStructure called with nil BodyStructure")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BodyStructure", *updateResource.Id)
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

	var parsed r5.BodyStructure
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return BodyStructure or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchBodyStructure(patchResource *r5.BodyStructure) (*r5.BodyStructure, error) {
	if patchResource == nil {
		return nil, errors.New("PatchBodyStructure given nil BodyStructure")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchBodyStructure given BodyStructure without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BodyStructure", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchBodyStructure error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchBodyStructure error using patch representation as json")
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

	var parsed r5.BodyStructure
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete BodyStructure and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBodyStructure(deleteResource *r5.BodyStructure) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteBodyStructure given nil BodyStructure")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteBodyStructure given BodyStructure with nil Id (Id required to delete)")
	}
	return fc.DeleteBodyStructureById(*deleteResource.Id)
}

// delete BodyStructure with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBodyStructureById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BodyStructure", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create CanonicalResource, return id of created CanonicalResource or OperationOutcome error
func (fc *FhirClient) CreateCanonicalResource(createResource *r5.CanonicalResource) (*r5.CanonicalResource, error) {
	if createResource == nil {
		return nil, errors.New("CreateCanonicalResource called with nil CanonicalResource")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CanonicalResource")
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

	var parsed r5.CanonicalResource
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read CanonicalResource from fhir server by id, return CanonicalResource or OperationOutcome error
func (fc *FhirClient) ReadCanonicalResource(id string) (*r5.CanonicalResource, error) {
	if id == "" {
		return nil, errors.New("ReadCanonicalResource given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CanonicalResource", id)
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

	var parsed r5.CanonicalResource
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace CanonicalResource if exists on server, else create new CanonicalResource with given id, return CanonicalResource or OperationOutcome error
func (fc *FhirClient) UpdateCanonicalResource(updateResource *r5.CanonicalResource) (*r5.CanonicalResource, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateCanonicalResource called with nil CanonicalResource")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CanonicalResource", *updateResource.Id)
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

	var parsed r5.CanonicalResource
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return CanonicalResource or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCanonicalResource(patchResource *r5.CanonicalResource) (*r5.CanonicalResource, error) {
	if patchResource == nil {
		return nil, errors.New("PatchCanonicalResource given nil CanonicalResource")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchCanonicalResource given CanonicalResource without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CanonicalResource", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchCanonicalResource error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchCanonicalResource error using patch representation as json")
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

	var parsed r5.CanonicalResource
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete CanonicalResource and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCanonicalResource(deleteResource *r5.CanonicalResource) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCanonicalResource given nil CanonicalResource")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCanonicalResource given CanonicalResource with nil Id (Id required to delete)")
	}
	return fc.DeleteCanonicalResourceById(*deleteResource.Id)
}

// delete CanonicalResource with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCanonicalResourceById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CanonicalResource", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create CapabilityStatement, return id of created CapabilityStatement or OperationOutcome error
func (fc *FhirClient) CreateCapabilityStatement(createResource *r5.CapabilityStatement) (*r5.CapabilityStatement, error) {
	if createResource == nil {
		return nil, errors.New("CreateCapabilityStatement called with nil CapabilityStatement")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CapabilityStatement")
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

	var parsed r5.CapabilityStatement
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read CapabilityStatement from fhir server by id, return CapabilityStatement or OperationOutcome error
func (fc *FhirClient) ReadCapabilityStatement(id string) (*r5.CapabilityStatement, error) {
	if id == "" {
		return nil, errors.New("ReadCapabilityStatement given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CapabilityStatement", id)
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

	var parsed r5.CapabilityStatement
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace CapabilityStatement if exists on server, else create new CapabilityStatement with given id, return CapabilityStatement or OperationOutcome error
func (fc *FhirClient) UpdateCapabilityStatement(updateResource *r5.CapabilityStatement) (*r5.CapabilityStatement, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateCapabilityStatement called with nil CapabilityStatement")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CapabilityStatement", *updateResource.Id)
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

	var parsed r5.CapabilityStatement
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return CapabilityStatement or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCapabilityStatement(patchResource *r5.CapabilityStatement) (*r5.CapabilityStatement, error) {
	if patchResource == nil {
		return nil, errors.New("PatchCapabilityStatement given nil CapabilityStatement")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchCapabilityStatement given CapabilityStatement without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CapabilityStatement", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchCapabilityStatement error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchCapabilityStatement error using patch representation as json")
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

	var parsed r5.CapabilityStatement
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete CapabilityStatement and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCapabilityStatement(deleteResource *r5.CapabilityStatement) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCapabilityStatement given nil CapabilityStatement")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCapabilityStatement given CapabilityStatement with nil Id (Id required to delete)")
	}
	return fc.DeleteCapabilityStatementById(*deleteResource.Id)
}

// delete CapabilityStatement with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCapabilityStatementById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CapabilityStatement", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create CarePlan, return id of created CarePlan or OperationOutcome error
func (fc *FhirClient) CreateCarePlan(createResource *r5.CarePlan) (*r5.CarePlan, error) {
	if createResource == nil {
		return nil, errors.New("CreateCarePlan called with nil CarePlan")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CarePlan")
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

	var parsed r5.CarePlan
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read CarePlan from fhir server by id, return CarePlan or OperationOutcome error
func (fc *FhirClient) ReadCarePlan(id string) (*r5.CarePlan, error) {
	if id == "" {
		return nil, errors.New("ReadCarePlan given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CarePlan", id)
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

	var parsed r5.CarePlan
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace CarePlan if exists on server, else create new CarePlan with given id, return CarePlan or OperationOutcome error
func (fc *FhirClient) UpdateCarePlan(updateResource *r5.CarePlan) (*r5.CarePlan, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateCarePlan called with nil CarePlan")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CarePlan", *updateResource.Id)
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

	var parsed r5.CarePlan
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return CarePlan or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCarePlan(patchResource *r5.CarePlan) (*r5.CarePlan, error) {
	if patchResource == nil {
		return nil, errors.New("PatchCarePlan given nil CarePlan")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchCarePlan given CarePlan without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CarePlan", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchCarePlan error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchCarePlan error using patch representation as json")
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

	var parsed r5.CarePlan
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete CarePlan and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCarePlan(deleteResource *r5.CarePlan) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCarePlan given nil CarePlan")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCarePlan given CarePlan with nil Id (Id required to delete)")
	}
	return fc.DeleteCarePlanById(*deleteResource.Id)
}

// delete CarePlan with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCarePlanById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CarePlan", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create CareTeam, return id of created CareTeam or OperationOutcome error
func (fc *FhirClient) CreateCareTeam(createResource *r5.CareTeam) (*r5.CareTeam, error) {
	if createResource == nil {
		return nil, errors.New("CreateCareTeam called with nil CareTeam")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CareTeam")
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

	var parsed r5.CareTeam
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read CareTeam from fhir server by id, return CareTeam or OperationOutcome error
func (fc *FhirClient) ReadCareTeam(id string) (*r5.CareTeam, error) {
	if id == "" {
		return nil, errors.New("ReadCareTeam given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CareTeam", id)
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

	var parsed r5.CareTeam
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace CareTeam if exists on server, else create new CareTeam with given id, return CareTeam or OperationOutcome error
func (fc *FhirClient) UpdateCareTeam(updateResource *r5.CareTeam) (*r5.CareTeam, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateCareTeam called with nil CareTeam")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CareTeam", *updateResource.Id)
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

	var parsed r5.CareTeam
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return CareTeam or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCareTeam(patchResource *r5.CareTeam) (*r5.CareTeam, error) {
	if patchResource == nil {
		return nil, errors.New("PatchCareTeam given nil CareTeam")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchCareTeam given CareTeam without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CareTeam", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchCareTeam error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchCareTeam error using patch representation as json")
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

	var parsed r5.CareTeam
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete CareTeam and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCareTeam(deleteResource *r5.CareTeam) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCareTeam given nil CareTeam")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCareTeam given CareTeam with nil Id (Id required to delete)")
	}
	return fc.DeleteCareTeamById(*deleteResource.Id)
}

// delete CareTeam with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCareTeamById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CareTeam", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ChargeItem, return id of created ChargeItem or OperationOutcome error
func (fc *FhirClient) CreateChargeItem(createResource *r5.ChargeItem) (*r5.ChargeItem, error) {
	if createResource == nil {
		return nil, errors.New("CreateChargeItem called with nil ChargeItem")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ChargeItem")
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

	var parsed r5.ChargeItem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ChargeItem from fhir server by id, return ChargeItem or OperationOutcome error
func (fc *FhirClient) ReadChargeItem(id string) (*r5.ChargeItem, error) {
	if id == "" {
		return nil, errors.New("ReadChargeItem given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ChargeItem", id)
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

	var parsed r5.ChargeItem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ChargeItem if exists on server, else create new ChargeItem with given id, return ChargeItem or OperationOutcome error
func (fc *FhirClient) UpdateChargeItem(updateResource *r5.ChargeItem) (*r5.ChargeItem, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateChargeItem called with nil ChargeItem")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ChargeItem", *updateResource.Id)
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

	var parsed r5.ChargeItem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ChargeItem or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchChargeItem(patchResource *r5.ChargeItem) (*r5.ChargeItem, error) {
	if patchResource == nil {
		return nil, errors.New("PatchChargeItem given nil ChargeItem")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchChargeItem given ChargeItem without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ChargeItem", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchChargeItem error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchChargeItem error using patch representation as json")
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

	var parsed r5.ChargeItem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ChargeItem and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteChargeItem(deleteResource *r5.ChargeItem) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteChargeItem given nil ChargeItem")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteChargeItem given ChargeItem with nil Id (Id required to delete)")
	}
	return fc.DeleteChargeItemById(*deleteResource.Id)
}

// delete ChargeItem with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteChargeItemById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ChargeItem", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ChargeItemDefinition, return id of created ChargeItemDefinition or OperationOutcome error
func (fc *FhirClient) CreateChargeItemDefinition(createResource *r5.ChargeItemDefinition) (*r5.ChargeItemDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateChargeItemDefinition called with nil ChargeItemDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ChargeItemDefinition")
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

	var parsed r5.ChargeItemDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ChargeItemDefinition from fhir server by id, return ChargeItemDefinition or OperationOutcome error
func (fc *FhirClient) ReadChargeItemDefinition(id string) (*r5.ChargeItemDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadChargeItemDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ChargeItemDefinition", id)
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

	var parsed r5.ChargeItemDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ChargeItemDefinition if exists on server, else create new ChargeItemDefinition with given id, return ChargeItemDefinition or OperationOutcome error
func (fc *FhirClient) UpdateChargeItemDefinition(updateResource *r5.ChargeItemDefinition) (*r5.ChargeItemDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateChargeItemDefinition called with nil ChargeItemDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ChargeItemDefinition", *updateResource.Id)
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

	var parsed r5.ChargeItemDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ChargeItemDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchChargeItemDefinition(patchResource *r5.ChargeItemDefinition) (*r5.ChargeItemDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchChargeItemDefinition given nil ChargeItemDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchChargeItemDefinition given ChargeItemDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ChargeItemDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchChargeItemDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchChargeItemDefinition error using patch representation as json")
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

	var parsed r5.ChargeItemDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ChargeItemDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteChargeItemDefinition(deleteResource *r5.ChargeItemDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteChargeItemDefinition given nil ChargeItemDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteChargeItemDefinition given ChargeItemDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteChargeItemDefinitionById(*deleteResource.Id)
}

// delete ChargeItemDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteChargeItemDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ChargeItemDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Citation, return id of created Citation or OperationOutcome error
func (fc *FhirClient) CreateCitation(createResource *r5.Citation) (*r5.Citation, error) {
	if createResource == nil {
		return nil, errors.New("CreateCitation called with nil Citation")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Citation")
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

	var parsed r5.Citation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Citation from fhir server by id, return Citation or OperationOutcome error
func (fc *FhirClient) ReadCitation(id string) (*r5.Citation, error) {
	if id == "" {
		return nil, errors.New("ReadCitation given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Citation", id)
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

	var parsed r5.Citation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Citation if exists on server, else create new Citation with given id, return Citation or OperationOutcome error
func (fc *FhirClient) UpdateCitation(updateResource *r5.Citation) (*r5.Citation, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateCitation called with nil Citation")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Citation", *updateResource.Id)
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

	var parsed r5.Citation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Citation or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCitation(patchResource *r5.Citation) (*r5.Citation, error) {
	if patchResource == nil {
		return nil, errors.New("PatchCitation given nil Citation")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchCitation given Citation without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Citation", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchCitation error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchCitation error using patch representation as json")
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

	var parsed r5.Citation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Citation and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCitation(deleteResource *r5.Citation) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCitation given nil Citation")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCitation given Citation with nil Id (Id required to delete)")
	}
	return fc.DeleteCitationById(*deleteResource.Id)
}

// delete Citation with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCitationById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Citation", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Claim, return id of created Claim or OperationOutcome error
func (fc *FhirClient) CreateClaim(createResource *r5.Claim) (*r5.Claim, error) {
	if createResource == nil {
		return nil, errors.New("CreateClaim called with nil Claim")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Claim")
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

	var parsed r5.Claim
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Claim from fhir server by id, return Claim or OperationOutcome error
func (fc *FhirClient) ReadClaim(id string) (*r5.Claim, error) {
	if id == "" {
		return nil, errors.New("ReadClaim given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Claim", id)
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

	var parsed r5.Claim
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Claim if exists on server, else create new Claim with given id, return Claim or OperationOutcome error
func (fc *FhirClient) UpdateClaim(updateResource *r5.Claim) (*r5.Claim, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateClaim called with nil Claim")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Claim", *updateResource.Id)
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

	var parsed r5.Claim
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Claim or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchClaim(patchResource *r5.Claim) (*r5.Claim, error) {
	if patchResource == nil {
		return nil, errors.New("PatchClaim given nil Claim")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchClaim given Claim without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Claim", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchClaim error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchClaim error using patch representation as json")
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

	var parsed r5.Claim
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Claim and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteClaim(deleteResource *r5.Claim) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteClaim given nil Claim")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteClaim given Claim with nil Id (Id required to delete)")
	}
	return fc.DeleteClaimById(*deleteResource.Id)
}

// delete Claim with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteClaimById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Claim", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ClaimResponse, return id of created ClaimResponse or OperationOutcome error
func (fc *FhirClient) CreateClaimResponse(createResource *r5.ClaimResponse) (*r5.ClaimResponse, error) {
	if createResource == nil {
		return nil, errors.New("CreateClaimResponse called with nil ClaimResponse")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClaimResponse")
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

	var parsed r5.ClaimResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ClaimResponse from fhir server by id, return ClaimResponse or OperationOutcome error
func (fc *FhirClient) ReadClaimResponse(id string) (*r5.ClaimResponse, error) {
	if id == "" {
		return nil, errors.New("ReadClaimResponse given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClaimResponse", id)
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

	var parsed r5.ClaimResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ClaimResponse if exists on server, else create new ClaimResponse with given id, return ClaimResponse or OperationOutcome error
func (fc *FhirClient) UpdateClaimResponse(updateResource *r5.ClaimResponse) (*r5.ClaimResponse, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateClaimResponse called with nil ClaimResponse")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClaimResponse", *updateResource.Id)
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

	var parsed r5.ClaimResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ClaimResponse or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchClaimResponse(patchResource *r5.ClaimResponse) (*r5.ClaimResponse, error) {
	if patchResource == nil {
		return nil, errors.New("PatchClaimResponse given nil ClaimResponse")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchClaimResponse given ClaimResponse without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClaimResponse", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchClaimResponse error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchClaimResponse error using patch representation as json")
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

	var parsed r5.ClaimResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ClaimResponse and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteClaimResponse(deleteResource *r5.ClaimResponse) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteClaimResponse given nil ClaimResponse")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteClaimResponse given ClaimResponse with nil Id (Id required to delete)")
	}
	return fc.DeleteClaimResponseById(*deleteResource.Id)
}

// delete ClaimResponse with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteClaimResponseById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClaimResponse", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ClinicalImpression, return id of created ClinicalImpression or OperationOutcome error
func (fc *FhirClient) CreateClinicalImpression(createResource *r5.ClinicalImpression) (*r5.ClinicalImpression, error) {
	if createResource == nil {
		return nil, errors.New("CreateClinicalImpression called with nil ClinicalImpression")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClinicalImpression")
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

	var parsed r5.ClinicalImpression
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ClinicalImpression from fhir server by id, return ClinicalImpression or OperationOutcome error
func (fc *FhirClient) ReadClinicalImpression(id string) (*r5.ClinicalImpression, error) {
	if id == "" {
		return nil, errors.New("ReadClinicalImpression given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClinicalImpression", id)
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

	var parsed r5.ClinicalImpression
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ClinicalImpression if exists on server, else create new ClinicalImpression with given id, return ClinicalImpression or OperationOutcome error
func (fc *FhirClient) UpdateClinicalImpression(updateResource *r5.ClinicalImpression) (*r5.ClinicalImpression, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateClinicalImpression called with nil ClinicalImpression")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClinicalImpression", *updateResource.Id)
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

	var parsed r5.ClinicalImpression
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ClinicalImpression or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchClinicalImpression(patchResource *r5.ClinicalImpression) (*r5.ClinicalImpression, error) {
	if patchResource == nil {
		return nil, errors.New("PatchClinicalImpression given nil ClinicalImpression")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchClinicalImpression given ClinicalImpression without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClinicalImpression", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchClinicalImpression error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchClinicalImpression error using patch representation as json")
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

	var parsed r5.ClinicalImpression
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ClinicalImpression and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteClinicalImpression(deleteResource *r5.ClinicalImpression) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteClinicalImpression given nil ClinicalImpression")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteClinicalImpression given ClinicalImpression with nil Id (Id required to delete)")
	}
	return fc.DeleteClinicalImpressionById(*deleteResource.Id)
}

// delete ClinicalImpression with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteClinicalImpressionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClinicalImpression", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ClinicalUseDefinition, return id of created ClinicalUseDefinition or OperationOutcome error
func (fc *FhirClient) CreateClinicalUseDefinition(createResource *r5.ClinicalUseDefinition) (*r5.ClinicalUseDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateClinicalUseDefinition called with nil ClinicalUseDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClinicalUseDefinition")
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

	var parsed r5.ClinicalUseDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ClinicalUseDefinition from fhir server by id, return ClinicalUseDefinition or OperationOutcome error
func (fc *FhirClient) ReadClinicalUseDefinition(id string) (*r5.ClinicalUseDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadClinicalUseDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClinicalUseDefinition", id)
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

	var parsed r5.ClinicalUseDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ClinicalUseDefinition if exists on server, else create new ClinicalUseDefinition with given id, return ClinicalUseDefinition or OperationOutcome error
func (fc *FhirClient) UpdateClinicalUseDefinition(updateResource *r5.ClinicalUseDefinition) (*r5.ClinicalUseDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateClinicalUseDefinition called with nil ClinicalUseDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClinicalUseDefinition", *updateResource.Id)
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

	var parsed r5.ClinicalUseDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ClinicalUseDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchClinicalUseDefinition(patchResource *r5.ClinicalUseDefinition) (*r5.ClinicalUseDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchClinicalUseDefinition given nil ClinicalUseDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchClinicalUseDefinition given ClinicalUseDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClinicalUseDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchClinicalUseDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchClinicalUseDefinition error using patch representation as json")
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

	var parsed r5.ClinicalUseDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ClinicalUseDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteClinicalUseDefinition(deleteResource *r5.ClinicalUseDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteClinicalUseDefinition given nil ClinicalUseDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteClinicalUseDefinition given ClinicalUseDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteClinicalUseDefinitionById(*deleteResource.Id)
}

// delete ClinicalUseDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteClinicalUseDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClinicalUseDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create CodeSystem, return id of created CodeSystem or OperationOutcome error
func (fc *FhirClient) CreateCodeSystem(createResource *r5.CodeSystem) (*r5.CodeSystem, error) {
	if createResource == nil {
		return nil, errors.New("CreateCodeSystem called with nil CodeSystem")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CodeSystem")
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

	var parsed r5.CodeSystem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read CodeSystem from fhir server by id, return CodeSystem or OperationOutcome error
func (fc *FhirClient) ReadCodeSystem(id string) (*r5.CodeSystem, error) {
	if id == "" {
		return nil, errors.New("ReadCodeSystem given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CodeSystem", id)
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

	var parsed r5.CodeSystem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace CodeSystem if exists on server, else create new CodeSystem with given id, return CodeSystem or OperationOutcome error
func (fc *FhirClient) UpdateCodeSystem(updateResource *r5.CodeSystem) (*r5.CodeSystem, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateCodeSystem called with nil CodeSystem")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CodeSystem", *updateResource.Id)
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

	var parsed r5.CodeSystem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return CodeSystem or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCodeSystem(patchResource *r5.CodeSystem) (*r5.CodeSystem, error) {
	if patchResource == nil {
		return nil, errors.New("PatchCodeSystem given nil CodeSystem")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchCodeSystem given CodeSystem without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CodeSystem", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchCodeSystem error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchCodeSystem error using patch representation as json")
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

	var parsed r5.CodeSystem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete CodeSystem and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCodeSystem(deleteResource *r5.CodeSystem) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCodeSystem given nil CodeSystem")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCodeSystem given CodeSystem with nil Id (Id required to delete)")
	}
	return fc.DeleteCodeSystemById(*deleteResource.Id)
}

// delete CodeSystem with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCodeSystemById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CodeSystem", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Communication, return id of created Communication or OperationOutcome error
func (fc *FhirClient) CreateCommunication(createResource *r5.Communication) (*r5.Communication, error) {
	if createResource == nil {
		return nil, errors.New("CreateCommunication called with nil Communication")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Communication")
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

	var parsed r5.Communication
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Communication from fhir server by id, return Communication or OperationOutcome error
func (fc *FhirClient) ReadCommunication(id string) (*r5.Communication, error) {
	if id == "" {
		return nil, errors.New("ReadCommunication given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Communication", id)
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

	var parsed r5.Communication
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Communication if exists on server, else create new Communication with given id, return Communication or OperationOutcome error
func (fc *FhirClient) UpdateCommunication(updateResource *r5.Communication) (*r5.Communication, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateCommunication called with nil Communication")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Communication", *updateResource.Id)
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

	var parsed r5.Communication
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Communication or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCommunication(patchResource *r5.Communication) (*r5.Communication, error) {
	if patchResource == nil {
		return nil, errors.New("PatchCommunication given nil Communication")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchCommunication given Communication without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Communication", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchCommunication error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchCommunication error using patch representation as json")
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

	var parsed r5.Communication
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Communication and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCommunication(deleteResource *r5.Communication) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCommunication given nil Communication")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCommunication given Communication with nil Id (Id required to delete)")
	}
	return fc.DeleteCommunicationById(*deleteResource.Id)
}

// delete Communication with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCommunicationById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Communication", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create CommunicationRequest, return id of created CommunicationRequest or OperationOutcome error
func (fc *FhirClient) CreateCommunicationRequest(createResource *r5.CommunicationRequest) (*r5.CommunicationRequest, error) {
	if createResource == nil {
		return nil, errors.New("CreateCommunicationRequest called with nil CommunicationRequest")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CommunicationRequest")
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

	var parsed r5.CommunicationRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read CommunicationRequest from fhir server by id, return CommunicationRequest or OperationOutcome error
func (fc *FhirClient) ReadCommunicationRequest(id string) (*r5.CommunicationRequest, error) {
	if id == "" {
		return nil, errors.New("ReadCommunicationRequest given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CommunicationRequest", id)
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

	var parsed r5.CommunicationRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace CommunicationRequest if exists on server, else create new CommunicationRequest with given id, return CommunicationRequest or OperationOutcome error
func (fc *FhirClient) UpdateCommunicationRequest(updateResource *r5.CommunicationRequest) (*r5.CommunicationRequest, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateCommunicationRequest called with nil CommunicationRequest")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CommunicationRequest", *updateResource.Id)
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

	var parsed r5.CommunicationRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return CommunicationRequest or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCommunicationRequest(patchResource *r5.CommunicationRequest) (*r5.CommunicationRequest, error) {
	if patchResource == nil {
		return nil, errors.New("PatchCommunicationRequest given nil CommunicationRequest")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchCommunicationRequest given CommunicationRequest without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CommunicationRequest", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchCommunicationRequest error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchCommunicationRequest error using patch representation as json")
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

	var parsed r5.CommunicationRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete CommunicationRequest and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCommunicationRequest(deleteResource *r5.CommunicationRequest) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCommunicationRequest given nil CommunicationRequest")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCommunicationRequest given CommunicationRequest with nil Id (Id required to delete)")
	}
	return fc.DeleteCommunicationRequestById(*deleteResource.Id)
}

// delete CommunicationRequest with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCommunicationRequestById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CommunicationRequest", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create CompartmentDefinition, return id of created CompartmentDefinition or OperationOutcome error
func (fc *FhirClient) CreateCompartmentDefinition(createResource *r5.CompartmentDefinition) (*r5.CompartmentDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateCompartmentDefinition called with nil CompartmentDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CompartmentDefinition")
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

	var parsed r5.CompartmentDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read CompartmentDefinition from fhir server by id, return CompartmentDefinition or OperationOutcome error
func (fc *FhirClient) ReadCompartmentDefinition(id string) (*r5.CompartmentDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadCompartmentDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CompartmentDefinition", id)
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

	var parsed r5.CompartmentDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace CompartmentDefinition if exists on server, else create new CompartmentDefinition with given id, return CompartmentDefinition or OperationOutcome error
func (fc *FhirClient) UpdateCompartmentDefinition(updateResource *r5.CompartmentDefinition) (*r5.CompartmentDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateCompartmentDefinition called with nil CompartmentDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CompartmentDefinition", *updateResource.Id)
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

	var parsed r5.CompartmentDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return CompartmentDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCompartmentDefinition(patchResource *r5.CompartmentDefinition) (*r5.CompartmentDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchCompartmentDefinition given nil CompartmentDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchCompartmentDefinition given CompartmentDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CompartmentDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchCompartmentDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchCompartmentDefinition error using patch representation as json")
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

	var parsed r5.CompartmentDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete CompartmentDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCompartmentDefinition(deleteResource *r5.CompartmentDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCompartmentDefinition given nil CompartmentDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCompartmentDefinition given CompartmentDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteCompartmentDefinitionById(*deleteResource.Id)
}

// delete CompartmentDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCompartmentDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CompartmentDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Composition, return id of created Composition or OperationOutcome error
func (fc *FhirClient) CreateComposition(createResource *r5.Composition) (*r5.Composition, error) {
	if createResource == nil {
		return nil, errors.New("CreateComposition called with nil Composition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Composition")
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

	var parsed r5.Composition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Composition from fhir server by id, return Composition or OperationOutcome error
func (fc *FhirClient) ReadComposition(id string) (*r5.Composition, error) {
	if id == "" {
		return nil, errors.New("ReadComposition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Composition", id)
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

	var parsed r5.Composition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Composition if exists on server, else create new Composition with given id, return Composition or OperationOutcome error
func (fc *FhirClient) UpdateComposition(updateResource *r5.Composition) (*r5.Composition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateComposition called with nil Composition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Composition", *updateResource.Id)
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

	var parsed r5.Composition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Composition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchComposition(patchResource *r5.Composition) (*r5.Composition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchComposition given nil Composition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchComposition given Composition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Composition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchComposition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchComposition error using patch representation as json")
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

	var parsed r5.Composition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Composition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteComposition(deleteResource *r5.Composition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteComposition given nil Composition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteComposition given Composition with nil Id (Id required to delete)")
	}
	return fc.DeleteCompositionById(*deleteResource.Id)
}

// delete Composition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCompositionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Composition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ConceptMap, return id of created ConceptMap or OperationOutcome error
func (fc *FhirClient) CreateConceptMap(createResource *r5.ConceptMap) (*r5.ConceptMap, error) {
	if createResource == nil {
		return nil, errors.New("CreateConceptMap called with nil ConceptMap")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ConceptMap")
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

	var parsed r5.ConceptMap
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ConceptMap from fhir server by id, return ConceptMap or OperationOutcome error
func (fc *FhirClient) ReadConceptMap(id string) (*r5.ConceptMap, error) {
	if id == "" {
		return nil, errors.New("ReadConceptMap given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ConceptMap", id)
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

	var parsed r5.ConceptMap
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ConceptMap if exists on server, else create new ConceptMap with given id, return ConceptMap or OperationOutcome error
func (fc *FhirClient) UpdateConceptMap(updateResource *r5.ConceptMap) (*r5.ConceptMap, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateConceptMap called with nil ConceptMap")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ConceptMap", *updateResource.Id)
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

	var parsed r5.ConceptMap
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ConceptMap or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchConceptMap(patchResource *r5.ConceptMap) (*r5.ConceptMap, error) {
	if patchResource == nil {
		return nil, errors.New("PatchConceptMap given nil ConceptMap")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchConceptMap given ConceptMap without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ConceptMap", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchConceptMap error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchConceptMap error using patch representation as json")
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

	var parsed r5.ConceptMap
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ConceptMap and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteConceptMap(deleteResource *r5.ConceptMap) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteConceptMap given nil ConceptMap")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteConceptMap given ConceptMap with nil Id (Id required to delete)")
	}
	return fc.DeleteConceptMapById(*deleteResource.Id)
}

// delete ConceptMap with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteConceptMapById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ConceptMap", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Condition, return id of created Condition or OperationOutcome error
func (fc *FhirClient) CreateCondition(createResource *r5.Condition) (*r5.Condition, error) {
	if createResource == nil {
		return nil, errors.New("CreateCondition called with nil Condition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Condition")
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

	var parsed r5.Condition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Condition from fhir server by id, return Condition or OperationOutcome error
func (fc *FhirClient) ReadCondition(id string) (*r5.Condition, error) {
	if id == "" {
		return nil, errors.New("ReadCondition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Condition", id)
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

	var parsed r5.Condition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Condition if exists on server, else create new Condition with given id, return Condition or OperationOutcome error
func (fc *FhirClient) UpdateCondition(updateResource *r5.Condition) (*r5.Condition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateCondition called with nil Condition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Condition", *updateResource.Id)
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

	var parsed r5.Condition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Condition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCondition(patchResource *r5.Condition) (*r5.Condition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchCondition given nil Condition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchCondition given Condition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Condition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchCondition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchCondition error using patch representation as json")
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

	var parsed r5.Condition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Condition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCondition(deleteResource *r5.Condition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCondition given nil Condition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCondition given Condition with nil Id (Id required to delete)")
	}
	return fc.DeleteConditionById(*deleteResource.Id)
}

// delete Condition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteConditionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Condition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ConditionDefinition, return id of created ConditionDefinition or OperationOutcome error
func (fc *FhirClient) CreateConditionDefinition(createResource *r5.ConditionDefinition) (*r5.ConditionDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateConditionDefinition called with nil ConditionDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ConditionDefinition")
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

	var parsed r5.ConditionDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ConditionDefinition from fhir server by id, return ConditionDefinition or OperationOutcome error
func (fc *FhirClient) ReadConditionDefinition(id string) (*r5.ConditionDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadConditionDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ConditionDefinition", id)
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

	var parsed r5.ConditionDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ConditionDefinition if exists on server, else create new ConditionDefinition with given id, return ConditionDefinition or OperationOutcome error
func (fc *FhirClient) UpdateConditionDefinition(updateResource *r5.ConditionDefinition) (*r5.ConditionDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateConditionDefinition called with nil ConditionDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ConditionDefinition", *updateResource.Id)
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

	var parsed r5.ConditionDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ConditionDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchConditionDefinition(patchResource *r5.ConditionDefinition) (*r5.ConditionDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchConditionDefinition given nil ConditionDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchConditionDefinition given ConditionDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ConditionDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchConditionDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchConditionDefinition error using patch representation as json")
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

	var parsed r5.ConditionDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ConditionDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteConditionDefinition(deleteResource *r5.ConditionDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteConditionDefinition given nil ConditionDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteConditionDefinition given ConditionDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteConditionDefinitionById(*deleteResource.Id)
}

// delete ConditionDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteConditionDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ConditionDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Consent, return id of created Consent or OperationOutcome error
func (fc *FhirClient) CreateConsent(createResource *r5.Consent) (*r5.Consent, error) {
	if createResource == nil {
		return nil, errors.New("CreateConsent called with nil Consent")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Consent")
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

	var parsed r5.Consent
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Consent from fhir server by id, return Consent or OperationOutcome error
func (fc *FhirClient) ReadConsent(id string) (*r5.Consent, error) {
	if id == "" {
		return nil, errors.New("ReadConsent given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Consent", id)
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

	var parsed r5.Consent
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Consent if exists on server, else create new Consent with given id, return Consent or OperationOutcome error
func (fc *FhirClient) UpdateConsent(updateResource *r5.Consent) (*r5.Consent, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateConsent called with nil Consent")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Consent", *updateResource.Id)
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

	var parsed r5.Consent
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Consent or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchConsent(patchResource *r5.Consent) (*r5.Consent, error) {
	if patchResource == nil {
		return nil, errors.New("PatchConsent given nil Consent")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchConsent given Consent without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Consent", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchConsent error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchConsent error using patch representation as json")
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

	var parsed r5.Consent
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Consent and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteConsent(deleteResource *r5.Consent) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteConsent given nil Consent")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteConsent given Consent with nil Id (Id required to delete)")
	}
	return fc.DeleteConsentById(*deleteResource.Id)
}

// delete Consent with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteConsentById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Consent", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Contract, return id of created Contract or OperationOutcome error
func (fc *FhirClient) CreateContract(createResource *r5.Contract) (*r5.Contract, error) {
	if createResource == nil {
		return nil, errors.New("CreateContract called with nil Contract")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Contract")
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

	var parsed r5.Contract
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Contract from fhir server by id, return Contract or OperationOutcome error
func (fc *FhirClient) ReadContract(id string) (*r5.Contract, error) {
	if id == "" {
		return nil, errors.New("ReadContract given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Contract", id)
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

	var parsed r5.Contract
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Contract if exists on server, else create new Contract with given id, return Contract or OperationOutcome error
func (fc *FhirClient) UpdateContract(updateResource *r5.Contract) (*r5.Contract, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateContract called with nil Contract")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Contract", *updateResource.Id)
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

	var parsed r5.Contract
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Contract or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchContract(patchResource *r5.Contract) (*r5.Contract, error) {
	if patchResource == nil {
		return nil, errors.New("PatchContract given nil Contract")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchContract given Contract without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Contract", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchContract error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchContract error using patch representation as json")
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

	var parsed r5.Contract
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Contract and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteContract(deleteResource *r5.Contract) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteContract given nil Contract")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteContract given Contract with nil Id (Id required to delete)")
	}
	return fc.DeleteContractById(*deleteResource.Id)
}

// delete Contract with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteContractById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Contract", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Coverage, return id of created Coverage or OperationOutcome error
func (fc *FhirClient) CreateCoverage(createResource *r5.Coverage) (*r5.Coverage, error) {
	if createResource == nil {
		return nil, errors.New("CreateCoverage called with nil Coverage")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Coverage")
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

	var parsed r5.Coverage
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Coverage from fhir server by id, return Coverage or OperationOutcome error
func (fc *FhirClient) ReadCoverage(id string) (*r5.Coverage, error) {
	if id == "" {
		return nil, errors.New("ReadCoverage given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Coverage", id)
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

	var parsed r5.Coverage
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Coverage if exists on server, else create new Coverage with given id, return Coverage or OperationOutcome error
func (fc *FhirClient) UpdateCoverage(updateResource *r5.Coverage) (*r5.Coverage, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateCoverage called with nil Coverage")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Coverage", *updateResource.Id)
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

	var parsed r5.Coverage
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Coverage or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCoverage(patchResource *r5.Coverage) (*r5.Coverage, error) {
	if patchResource == nil {
		return nil, errors.New("PatchCoverage given nil Coverage")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchCoverage given Coverage without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Coverage", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchCoverage error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchCoverage error using patch representation as json")
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

	var parsed r5.Coverage
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Coverage and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCoverage(deleteResource *r5.Coverage) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCoverage given nil Coverage")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCoverage given Coverage with nil Id (Id required to delete)")
	}
	return fc.DeleteCoverageById(*deleteResource.Id)
}

// delete Coverage with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCoverageById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Coverage", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create CoverageEligibilityRequest, return id of created CoverageEligibilityRequest or OperationOutcome error
func (fc *FhirClient) CreateCoverageEligibilityRequest(createResource *r5.CoverageEligibilityRequest) (*r5.CoverageEligibilityRequest, error) {
	if createResource == nil {
		return nil, errors.New("CreateCoverageEligibilityRequest called with nil CoverageEligibilityRequest")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CoverageEligibilityRequest")
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

	var parsed r5.CoverageEligibilityRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read CoverageEligibilityRequest from fhir server by id, return CoverageEligibilityRequest or OperationOutcome error
func (fc *FhirClient) ReadCoverageEligibilityRequest(id string) (*r5.CoverageEligibilityRequest, error) {
	if id == "" {
		return nil, errors.New("ReadCoverageEligibilityRequest given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CoverageEligibilityRequest", id)
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

	var parsed r5.CoverageEligibilityRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace CoverageEligibilityRequest if exists on server, else create new CoverageEligibilityRequest with given id, return CoverageEligibilityRequest or OperationOutcome error
func (fc *FhirClient) UpdateCoverageEligibilityRequest(updateResource *r5.CoverageEligibilityRequest) (*r5.CoverageEligibilityRequest, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateCoverageEligibilityRequest called with nil CoverageEligibilityRequest")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CoverageEligibilityRequest", *updateResource.Id)
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

	var parsed r5.CoverageEligibilityRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return CoverageEligibilityRequest or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCoverageEligibilityRequest(patchResource *r5.CoverageEligibilityRequest) (*r5.CoverageEligibilityRequest, error) {
	if patchResource == nil {
		return nil, errors.New("PatchCoverageEligibilityRequest given nil CoverageEligibilityRequest")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchCoverageEligibilityRequest given CoverageEligibilityRequest without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CoverageEligibilityRequest", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchCoverageEligibilityRequest error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchCoverageEligibilityRequest error using patch representation as json")
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

	var parsed r5.CoverageEligibilityRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete CoverageEligibilityRequest and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCoverageEligibilityRequest(deleteResource *r5.CoverageEligibilityRequest) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCoverageEligibilityRequest given nil CoverageEligibilityRequest")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCoverageEligibilityRequest given CoverageEligibilityRequest with nil Id (Id required to delete)")
	}
	return fc.DeleteCoverageEligibilityRequestById(*deleteResource.Id)
}

// delete CoverageEligibilityRequest with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCoverageEligibilityRequestById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CoverageEligibilityRequest", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create CoverageEligibilityResponse, return id of created CoverageEligibilityResponse or OperationOutcome error
func (fc *FhirClient) CreateCoverageEligibilityResponse(createResource *r5.CoverageEligibilityResponse) (*r5.CoverageEligibilityResponse, error) {
	if createResource == nil {
		return nil, errors.New("CreateCoverageEligibilityResponse called with nil CoverageEligibilityResponse")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CoverageEligibilityResponse")
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

	var parsed r5.CoverageEligibilityResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read CoverageEligibilityResponse from fhir server by id, return CoverageEligibilityResponse or OperationOutcome error
func (fc *FhirClient) ReadCoverageEligibilityResponse(id string) (*r5.CoverageEligibilityResponse, error) {
	if id == "" {
		return nil, errors.New("ReadCoverageEligibilityResponse given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CoverageEligibilityResponse", id)
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

	var parsed r5.CoverageEligibilityResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace CoverageEligibilityResponse if exists on server, else create new CoverageEligibilityResponse with given id, return CoverageEligibilityResponse or OperationOutcome error
func (fc *FhirClient) UpdateCoverageEligibilityResponse(updateResource *r5.CoverageEligibilityResponse) (*r5.CoverageEligibilityResponse, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateCoverageEligibilityResponse called with nil CoverageEligibilityResponse")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CoverageEligibilityResponse", *updateResource.Id)
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

	var parsed r5.CoverageEligibilityResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return CoverageEligibilityResponse or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCoverageEligibilityResponse(patchResource *r5.CoverageEligibilityResponse) (*r5.CoverageEligibilityResponse, error) {
	if patchResource == nil {
		return nil, errors.New("PatchCoverageEligibilityResponse given nil CoverageEligibilityResponse")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchCoverageEligibilityResponse given CoverageEligibilityResponse without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CoverageEligibilityResponse", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchCoverageEligibilityResponse error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchCoverageEligibilityResponse error using patch representation as json")
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

	var parsed r5.CoverageEligibilityResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete CoverageEligibilityResponse and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCoverageEligibilityResponse(deleteResource *r5.CoverageEligibilityResponse) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCoverageEligibilityResponse given nil CoverageEligibilityResponse")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCoverageEligibilityResponse given CoverageEligibilityResponse with nil Id (Id required to delete)")
	}
	return fc.DeleteCoverageEligibilityResponseById(*deleteResource.Id)
}

// delete CoverageEligibilityResponse with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCoverageEligibilityResponseById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CoverageEligibilityResponse", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create DetectedIssue, return id of created DetectedIssue or OperationOutcome error
func (fc *FhirClient) CreateDetectedIssue(createResource *r5.DetectedIssue) (*r5.DetectedIssue, error) {
	if createResource == nil {
		return nil, errors.New("CreateDetectedIssue called with nil DetectedIssue")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DetectedIssue")
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

	var parsed r5.DetectedIssue
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read DetectedIssue from fhir server by id, return DetectedIssue or OperationOutcome error
func (fc *FhirClient) ReadDetectedIssue(id string) (*r5.DetectedIssue, error) {
	if id == "" {
		return nil, errors.New("ReadDetectedIssue given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DetectedIssue", id)
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

	var parsed r5.DetectedIssue
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace DetectedIssue if exists on server, else create new DetectedIssue with given id, return DetectedIssue or OperationOutcome error
func (fc *FhirClient) UpdateDetectedIssue(updateResource *r5.DetectedIssue) (*r5.DetectedIssue, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateDetectedIssue called with nil DetectedIssue")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DetectedIssue", *updateResource.Id)
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

	var parsed r5.DetectedIssue
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return DetectedIssue or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDetectedIssue(patchResource *r5.DetectedIssue) (*r5.DetectedIssue, error) {
	if patchResource == nil {
		return nil, errors.New("PatchDetectedIssue given nil DetectedIssue")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchDetectedIssue given DetectedIssue without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DetectedIssue", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchDetectedIssue error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchDetectedIssue error using patch representation as json")
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

	var parsed r5.DetectedIssue
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete DetectedIssue and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDetectedIssue(deleteResource *r5.DetectedIssue) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDetectedIssue given nil DetectedIssue")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDetectedIssue given DetectedIssue with nil Id (Id required to delete)")
	}
	return fc.DeleteDetectedIssueById(*deleteResource.Id)
}

// delete DetectedIssue with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDetectedIssueById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DetectedIssue", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Device, return id of created Device or OperationOutcome error
func (fc *FhirClient) CreateDevice(createResource *r5.Device) (*r5.Device, error) {
	if createResource == nil {
		return nil, errors.New("CreateDevice called with nil Device")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Device")
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

	var parsed r5.Device
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Device from fhir server by id, return Device or OperationOutcome error
func (fc *FhirClient) ReadDevice(id string) (*r5.Device, error) {
	if id == "" {
		return nil, errors.New("ReadDevice given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Device", id)
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

	var parsed r5.Device
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Device if exists on server, else create new Device with given id, return Device or OperationOutcome error
func (fc *FhirClient) UpdateDevice(updateResource *r5.Device) (*r5.Device, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateDevice called with nil Device")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Device", *updateResource.Id)
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

	var parsed r5.Device
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Device or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDevice(patchResource *r5.Device) (*r5.Device, error) {
	if patchResource == nil {
		return nil, errors.New("PatchDevice given nil Device")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchDevice given Device without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Device", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchDevice error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchDevice error using patch representation as json")
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

	var parsed r5.Device
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Device and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDevice(deleteResource *r5.Device) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDevice given nil Device")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDevice given Device with nil Id (Id required to delete)")
	}
	return fc.DeleteDeviceById(*deleteResource.Id)
}

// delete Device with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Device", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create DeviceAssociation, return id of created DeviceAssociation or OperationOutcome error
func (fc *FhirClient) CreateDeviceAssociation(createResource *r5.DeviceAssociation) (*r5.DeviceAssociation, error) {
	if createResource == nil {
		return nil, errors.New("CreateDeviceAssociation called with nil DeviceAssociation")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceAssociation")
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

	var parsed r5.DeviceAssociation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read DeviceAssociation from fhir server by id, return DeviceAssociation or OperationOutcome error
func (fc *FhirClient) ReadDeviceAssociation(id string) (*r5.DeviceAssociation, error) {
	if id == "" {
		return nil, errors.New("ReadDeviceAssociation given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceAssociation", id)
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

	var parsed r5.DeviceAssociation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace DeviceAssociation if exists on server, else create new DeviceAssociation with given id, return DeviceAssociation or OperationOutcome error
func (fc *FhirClient) UpdateDeviceAssociation(updateResource *r5.DeviceAssociation) (*r5.DeviceAssociation, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateDeviceAssociation called with nil DeviceAssociation")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceAssociation", *updateResource.Id)
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

	var parsed r5.DeviceAssociation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return DeviceAssociation or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDeviceAssociation(patchResource *r5.DeviceAssociation) (*r5.DeviceAssociation, error) {
	if patchResource == nil {
		return nil, errors.New("PatchDeviceAssociation given nil DeviceAssociation")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchDeviceAssociation given DeviceAssociation without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceAssociation", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchDeviceAssociation error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchDeviceAssociation error using patch representation as json")
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

	var parsed r5.DeviceAssociation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete DeviceAssociation and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceAssociation(deleteResource *r5.DeviceAssociation) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDeviceAssociation given nil DeviceAssociation")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDeviceAssociation given DeviceAssociation with nil Id (Id required to delete)")
	}
	return fc.DeleteDeviceAssociationById(*deleteResource.Id)
}

// delete DeviceAssociation with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceAssociationById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceAssociation", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create DeviceDefinition, return id of created DeviceDefinition or OperationOutcome error
func (fc *FhirClient) CreateDeviceDefinition(createResource *r5.DeviceDefinition) (*r5.DeviceDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateDeviceDefinition called with nil DeviceDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceDefinition")
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

	var parsed r5.DeviceDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read DeviceDefinition from fhir server by id, return DeviceDefinition or OperationOutcome error
func (fc *FhirClient) ReadDeviceDefinition(id string) (*r5.DeviceDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadDeviceDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceDefinition", id)
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

	var parsed r5.DeviceDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace DeviceDefinition if exists on server, else create new DeviceDefinition with given id, return DeviceDefinition or OperationOutcome error
func (fc *FhirClient) UpdateDeviceDefinition(updateResource *r5.DeviceDefinition) (*r5.DeviceDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateDeviceDefinition called with nil DeviceDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceDefinition", *updateResource.Id)
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

	var parsed r5.DeviceDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return DeviceDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDeviceDefinition(patchResource *r5.DeviceDefinition) (*r5.DeviceDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchDeviceDefinition given nil DeviceDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchDeviceDefinition given DeviceDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchDeviceDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchDeviceDefinition error using patch representation as json")
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

	var parsed r5.DeviceDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete DeviceDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceDefinition(deleteResource *r5.DeviceDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDeviceDefinition given nil DeviceDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDeviceDefinition given DeviceDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteDeviceDefinitionById(*deleteResource.Id)
}

// delete DeviceDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create DeviceDispense, return id of created DeviceDispense or OperationOutcome error
func (fc *FhirClient) CreateDeviceDispense(createResource *r5.DeviceDispense) (*r5.DeviceDispense, error) {
	if createResource == nil {
		return nil, errors.New("CreateDeviceDispense called with nil DeviceDispense")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceDispense")
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

	var parsed r5.DeviceDispense
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read DeviceDispense from fhir server by id, return DeviceDispense or OperationOutcome error
func (fc *FhirClient) ReadDeviceDispense(id string) (*r5.DeviceDispense, error) {
	if id == "" {
		return nil, errors.New("ReadDeviceDispense given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceDispense", id)
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

	var parsed r5.DeviceDispense
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace DeviceDispense if exists on server, else create new DeviceDispense with given id, return DeviceDispense or OperationOutcome error
func (fc *FhirClient) UpdateDeviceDispense(updateResource *r5.DeviceDispense) (*r5.DeviceDispense, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateDeviceDispense called with nil DeviceDispense")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceDispense", *updateResource.Id)
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

	var parsed r5.DeviceDispense
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return DeviceDispense or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDeviceDispense(patchResource *r5.DeviceDispense) (*r5.DeviceDispense, error) {
	if patchResource == nil {
		return nil, errors.New("PatchDeviceDispense given nil DeviceDispense")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchDeviceDispense given DeviceDispense without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceDispense", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchDeviceDispense error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchDeviceDispense error using patch representation as json")
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

	var parsed r5.DeviceDispense
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete DeviceDispense and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceDispense(deleteResource *r5.DeviceDispense) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDeviceDispense given nil DeviceDispense")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDeviceDispense given DeviceDispense with nil Id (Id required to delete)")
	}
	return fc.DeleteDeviceDispenseById(*deleteResource.Id)
}

// delete DeviceDispense with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceDispenseById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceDispense", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create DeviceMetric, return id of created DeviceMetric or OperationOutcome error
func (fc *FhirClient) CreateDeviceMetric(createResource *r5.DeviceMetric) (*r5.DeviceMetric, error) {
	if createResource == nil {
		return nil, errors.New("CreateDeviceMetric called with nil DeviceMetric")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceMetric")
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

	var parsed r5.DeviceMetric
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read DeviceMetric from fhir server by id, return DeviceMetric or OperationOutcome error
func (fc *FhirClient) ReadDeviceMetric(id string) (*r5.DeviceMetric, error) {
	if id == "" {
		return nil, errors.New("ReadDeviceMetric given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceMetric", id)
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

	var parsed r5.DeviceMetric
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace DeviceMetric if exists on server, else create new DeviceMetric with given id, return DeviceMetric or OperationOutcome error
func (fc *FhirClient) UpdateDeviceMetric(updateResource *r5.DeviceMetric) (*r5.DeviceMetric, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateDeviceMetric called with nil DeviceMetric")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceMetric", *updateResource.Id)
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

	var parsed r5.DeviceMetric
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return DeviceMetric or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDeviceMetric(patchResource *r5.DeviceMetric) (*r5.DeviceMetric, error) {
	if patchResource == nil {
		return nil, errors.New("PatchDeviceMetric given nil DeviceMetric")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchDeviceMetric given DeviceMetric without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceMetric", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchDeviceMetric error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchDeviceMetric error using patch representation as json")
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

	var parsed r5.DeviceMetric
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete DeviceMetric and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceMetric(deleteResource *r5.DeviceMetric) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDeviceMetric given nil DeviceMetric")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDeviceMetric given DeviceMetric with nil Id (Id required to delete)")
	}
	return fc.DeleteDeviceMetricById(*deleteResource.Id)
}

// delete DeviceMetric with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceMetricById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceMetric", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create DeviceRequest, return id of created DeviceRequest or OperationOutcome error
func (fc *FhirClient) CreateDeviceRequest(createResource *r5.DeviceRequest) (*r5.DeviceRequest, error) {
	if createResource == nil {
		return nil, errors.New("CreateDeviceRequest called with nil DeviceRequest")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceRequest")
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

	var parsed r5.DeviceRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read DeviceRequest from fhir server by id, return DeviceRequest or OperationOutcome error
func (fc *FhirClient) ReadDeviceRequest(id string) (*r5.DeviceRequest, error) {
	if id == "" {
		return nil, errors.New("ReadDeviceRequest given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceRequest", id)
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

	var parsed r5.DeviceRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace DeviceRequest if exists on server, else create new DeviceRequest with given id, return DeviceRequest or OperationOutcome error
func (fc *FhirClient) UpdateDeviceRequest(updateResource *r5.DeviceRequest) (*r5.DeviceRequest, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateDeviceRequest called with nil DeviceRequest")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceRequest", *updateResource.Id)
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

	var parsed r5.DeviceRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return DeviceRequest or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDeviceRequest(patchResource *r5.DeviceRequest) (*r5.DeviceRequest, error) {
	if patchResource == nil {
		return nil, errors.New("PatchDeviceRequest given nil DeviceRequest")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchDeviceRequest given DeviceRequest without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceRequest", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchDeviceRequest error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchDeviceRequest error using patch representation as json")
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

	var parsed r5.DeviceRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete DeviceRequest and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceRequest(deleteResource *r5.DeviceRequest) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDeviceRequest given nil DeviceRequest")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDeviceRequest given DeviceRequest with nil Id (Id required to delete)")
	}
	return fc.DeleteDeviceRequestById(*deleteResource.Id)
}

// delete DeviceRequest with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceRequestById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceRequest", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create DeviceUsage, return id of created DeviceUsage or OperationOutcome error
func (fc *FhirClient) CreateDeviceUsage(createResource *r5.DeviceUsage) (*r5.DeviceUsage, error) {
	if createResource == nil {
		return nil, errors.New("CreateDeviceUsage called with nil DeviceUsage")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceUsage")
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

	var parsed r5.DeviceUsage
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read DeviceUsage from fhir server by id, return DeviceUsage or OperationOutcome error
func (fc *FhirClient) ReadDeviceUsage(id string) (*r5.DeviceUsage, error) {
	if id == "" {
		return nil, errors.New("ReadDeviceUsage given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceUsage", id)
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

	var parsed r5.DeviceUsage
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace DeviceUsage if exists on server, else create new DeviceUsage with given id, return DeviceUsage or OperationOutcome error
func (fc *FhirClient) UpdateDeviceUsage(updateResource *r5.DeviceUsage) (*r5.DeviceUsage, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateDeviceUsage called with nil DeviceUsage")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceUsage", *updateResource.Id)
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

	var parsed r5.DeviceUsage
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return DeviceUsage or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDeviceUsage(patchResource *r5.DeviceUsage) (*r5.DeviceUsage, error) {
	if patchResource == nil {
		return nil, errors.New("PatchDeviceUsage given nil DeviceUsage")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchDeviceUsage given DeviceUsage without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceUsage", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchDeviceUsage error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchDeviceUsage error using patch representation as json")
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

	var parsed r5.DeviceUsage
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete DeviceUsage and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceUsage(deleteResource *r5.DeviceUsage) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDeviceUsage given nil DeviceUsage")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDeviceUsage given DeviceUsage with nil Id (Id required to delete)")
	}
	return fc.DeleteDeviceUsageById(*deleteResource.Id)
}

// delete DeviceUsage with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceUsageById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceUsage", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create DiagnosticReport, return id of created DiagnosticReport or OperationOutcome error
func (fc *FhirClient) CreateDiagnosticReport(createResource *r5.DiagnosticReport) (*r5.DiagnosticReport, error) {
	if createResource == nil {
		return nil, errors.New("CreateDiagnosticReport called with nil DiagnosticReport")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DiagnosticReport")
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

	var parsed r5.DiagnosticReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read DiagnosticReport from fhir server by id, return DiagnosticReport or OperationOutcome error
func (fc *FhirClient) ReadDiagnosticReport(id string) (*r5.DiagnosticReport, error) {
	if id == "" {
		return nil, errors.New("ReadDiagnosticReport given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DiagnosticReport", id)
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

	var parsed r5.DiagnosticReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace DiagnosticReport if exists on server, else create new DiagnosticReport with given id, return DiagnosticReport or OperationOutcome error
func (fc *FhirClient) UpdateDiagnosticReport(updateResource *r5.DiagnosticReport) (*r5.DiagnosticReport, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateDiagnosticReport called with nil DiagnosticReport")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DiagnosticReport", *updateResource.Id)
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

	var parsed r5.DiagnosticReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return DiagnosticReport or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDiagnosticReport(patchResource *r5.DiagnosticReport) (*r5.DiagnosticReport, error) {
	if patchResource == nil {
		return nil, errors.New("PatchDiagnosticReport given nil DiagnosticReport")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchDiagnosticReport given DiagnosticReport without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DiagnosticReport", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchDiagnosticReport error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchDiagnosticReport error using patch representation as json")
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

	var parsed r5.DiagnosticReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete DiagnosticReport and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDiagnosticReport(deleteResource *r5.DiagnosticReport) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDiagnosticReport given nil DiagnosticReport")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDiagnosticReport given DiagnosticReport with nil Id (Id required to delete)")
	}
	return fc.DeleteDiagnosticReportById(*deleteResource.Id)
}

// delete DiagnosticReport with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDiagnosticReportById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DiagnosticReport", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create DocumentReference, return id of created DocumentReference or OperationOutcome error
func (fc *FhirClient) CreateDocumentReference(createResource *r5.DocumentReference) (*r5.DocumentReference, error) {
	if createResource == nil {
		return nil, errors.New("CreateDocumentReference called with nil DocumentReference")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DocumentReference")
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

	var parsed r5.DocumentReference
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read DocumentReference from fhir server by id, return DocumentReference or OperationOutcome error
func (fc *FhirClient) ReadDocumentReference(id string) (*r5.DocumentReference, error) {
	if id == "" {
		return nil, errors.New("ReadDocumentReference given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DocumentReference", id)
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

	var parsed r5.DocumentReference
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace DocumentReference if exists on server, else create new DocumentReference with given id, return DocumentReference or OperationOutcome error
func (fc *FhirClient) UpdateDocumentReference(updateResource *r5.DocumentReference) (*r5.DocumentReference, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateDocumentReference called with nil DocumentReference")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DocumentReference", *updateResource.Id)
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

	var parsed r5.DocumentReference
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return DocumentReference or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDocumentReference(patchResource *r5.DocumentReference) (*r5.DocumentReference, error) {
	if patchResource == nil {
		return nil, errors.New("PatchDocumentReference given nil DocumentReference")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchDocumentReference given DocumentReference without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DocumentReference", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchDocumentReference error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchDocumentReference error using patch representation as json")
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

	var parsed r5.DocumentReference
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete DocumentReference and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDocumentReference(deleteResource *r5.DocumentReference) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDocumentReference given nil DocumentReference")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDocumentReference given DocumentReference with nil Id (Id required to delete)")
	}
	return fc.DeleteDocumentReferenceById(*deleteResource.Id)
}

// delete DocumentReference with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDocumentReferenceById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DocumentReference", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Encounter, return id of created Encounter or OperationOutcome error
func (fc *FhirClient) CreateEncounter(createResource *r5.Encounter) (*r5.Encounter, error) {
	if createResource == nil {
		return nil, errors.New("CreateEncounter called with nil Encounter")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Encounter")
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

	var parsed r5.Encounter
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Encounter from fhir server by id, return Encounter or OperationOutcome error
func (fc *FhirClient) ReadEncounter(id string) (*r5.Encounter, error) {
	if id == "" {
		return nil, errors.New("ReadEncounter given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Encounter", id)
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

	var parsed r5.Encounter
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Encounter if exists on server, else create new Encounter with given id, return Encounter or OperationOutcome error
func (fc *FhirClient) UpdateEncounter(updateResource *r5.Encounter) (*r5.Encounter, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateEncounter called with nil Encounter")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Encounter", *updateResource.Id)
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

	var parsed r5.Encounter
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Encounter or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEncounter(patchResource *r5.Encounter) (*r5.Encounter, error) {
	if patchResource == nil {
		return nil, errors.New("PatchEncounter given nil Encounter")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchEncounter given Encounter without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Encounter", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchEncounter error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchEncounter error using patch representation as json")
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

	var parsed r5.Encounter
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Encounter and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEncounter(deleteResource *r5.Encounter) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEncounter given nil Encounter")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEncounter given Encounter with nil Id (Id required to delete)")
	}
	return fc.DeleteEncounterById(*deleteResource.Id)
}

// delete Encounter with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEncounterById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Encounter", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create EncounterHistory, return id of created EncounterHistory or OperationOutcome error
func (fc *FhirClient) CreateEncounterHistory(createResource *r5.EncounterHistory) (*r5.EncounterHistory, error) {
	if createResource == nil {
		return nil, errors.New("CreateEncounterHistory called with nil EncounterHistory")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EncounterHistory")
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

	var parsed r5.EncounterHistory
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read EncounterHistory from fhir server by id, return EncounterHistory or OperationOutcome error
func (fc *FhirClient) ReadEncounterHistory(id string) (*r5.EncounterHistory, error) {
	if id == "" {
		return nil, errors.New("ReadEncounterHistory given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EncounterHistory", id)
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

	var parsed r5.EncounterHistory
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace EncounterHistory if exists on server, else create new EncounterHistory with given id, return EncounterHistory or OperationOutcome error
func (fc *FhirClient) UpdateEncounterHistory(updateResource *r5.EncounterHistory) (*r5.EncounterHistory, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateEncounterHistory called with nil EncounterHistory")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EncounterHistory", *updateResource.Id)
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

	var parsed r5.EncounterHistory
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return EncounterHistory or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEncounterHistory(patchResource *r5.EncounterHistory) (*r5.EncounterHistory, error) {
	if patchResource == nil {
		return nil, errors.New("PatchEncounterHistory given nil EncounterHistory")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchEncounterHistory given EncounterHistory without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EncounterHistory", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchEncounterHistory error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchEncounterHistory error using patch representation as json")
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

	var parsed r5.EncounterHistory
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete EncounterHistory and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEncounterHistory(deleteResource *r5.EncounterHistory) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEncounterHistory given nil EncounterHistory")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEncounterHistory given EncounterHistory with nil Id (Id required to delete)")
	}
	return fc.DeleteEncounterHistoryById(*deleteResource.Id)
}

// delete EncounterHistory with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEncounterHistoryById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EncounterHistory", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Endpoint, return id of created Endpoint or OperationOutcome error
func (fc *FhirClient) CreateEndpoint(createResource *r5.Endpoint) (*r5.Endpoint, error) {
	if createResource == nil {
		return nil, errors.New("CreateEndpoint called with nil Endpoint")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Endpoint")
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

	var parsed r5.Endpoint
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Endpoint from fhir server by id, return Endpoint or OperationOutcome error
func (fc *FhirClient) ReadEndpoint(id string) (*r5.Endpoint, error) {
	if id == "" {
		return nil, errors.New("ReadEndpoint given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Endpoint", id)
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

	var parsed r5.Endpoint
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Endpoint if exists on server, else create new Endpoint with given id, return Endpoint or OperationOutcome error
func (fc *FhirClient) UpdateEndpoint(updateResource *r5.Endpoint) (*r5.Endpoint, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateEndpoint called with nil Endpoint")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Endpoint", *updateResource.Id)
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

	var parsed r5.Endpoint
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Endpoint or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEndpoint(patchResource *r5.Endpoint) (*r5.Endpoint, error) {
	if patchResource == nil {
		return nil, errors.New("PatchEndpoint given nil Endpoint")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchEndpoint given Endpoint without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Endpoint", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchEndpoint error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchEndpoint error using patch representation as json")
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

	var parsed r5.Endpoint
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Endpoint and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEndpoint(deleteResource *r5.Endpoint) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEndpoint given nil Endpoint")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEndpoint given Endpoint with nil Id (Id required to delete)")
	}
	return fc.DeleteEndpointById(*deleteResource.Id)
}

// delete Endpoint with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEndpointById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Endpoint", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create EnrollmentRequest, return id of created EnrollmentRequest or OperationOutcome error
func (fc *FhirClient) CreateEnrollmentRequest(createResource *r5.EnrollmentRequest) (*r5.EnrollmentRequest, error) {
	if createResource == nil {
		return nil, errors.New("CreateEnrollmentRequest called with nil EnrollmentRequest")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EnrollmentRequest")
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

	var parsed r5.EnrollmentRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read EnrollmentRequest from fhir server by id, return EnrollmentRequest or OperationOutcome error
func (fc *FhirClient) ReadEnrollmentRequest(id string) (*r5.EnrollmentRequest, error) {
	if id == "" {
		return nil, errors.New("ReadEnrollmentRequest given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EnrollmentRequest", id)
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

	var parsed r5.EnrollmentRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace EnrollmentRequest if exists on server, else create new EnrollmentRequest with given id, return EnrollmentRequest or OperationOutcome error
func (fc *FhirClient) UpdateEnrollmentRequest(updateResource *r5.EnrollmentRequest) (*r5.EnrollmentRequest, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateEnrollmentRequest called with nil EnrollmentRequest")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EnrollmentRequest", *updateResource.Id)
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

	var parsed r5.EnrollmentRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return EnrollmentRequest or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEnrollmentRequest(patchResource *r5.EnrollmentRequest) (*r5.EnrollmentRequest, error) {
	if patchResource == nil {
		return nil, errors.New("PatchEnrollmentRequest given nil EnrollmentRequest")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchEnrollmentRequest given EnrollmentRequest without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EnrollmentRequest", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchEnrollmentRequest error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchEnrollmentRequest error using patch representation as json")
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

	var parsed r5.EnrollmentRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete EnrollmentRequest and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEnrollmentRequest(deleteResource *r5.EnrollmentRequest) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEnrollmentRequest given nil EnrollmentRequest")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEnrollmentRequest given EnrollmentRequest with nil Id (Id required to delete)")
	}
	return fc.DeleteEnrollmentRequestById(*deleteResource.Id)
}

// delete EnrollmentRequest with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEnrollmentRequestById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EnrollmentRequest", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create EnrollmentResponse, return id of created EnrollmentResponse or OperationOutcome error
func (fc *FhirClient) CreateEnrollmentResponse(createResource *r5.EnrollmentResponse) (*r5.EnrollmentResponse, error) {
	if createResource == nil {
		return nil, errors.New("CreateEnrollmentResponse called with nil EnrollmentResponse")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EnrollmentResponse")
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

	var parsed r5.EnrollmentResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read EnrollmentResponse from fhir server by id, return EnrollmentResponse or OperationOutcome error
func (fc *FhirClient) ReadEnrollmentResponse(id string) (*r5.EnrollmentResponse, error) {
	if id == "" {
		return nil, errors.New("ReadEnrollmentResponse given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EnrollmentResponse", id)
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

	var parsed r5.EnrollmentResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace EnrollmentResponse if exists on server, else create new EnrollmentResponse with given id, return EnrollmentResponse or OperationOutcome error
func (fc *FhirClient) UpdateEnrollmentResponse(updateResource *r5.EnrollmentResponse) (*r5.EnrollmentResponse, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateEnrollmentResponse called with nil EnrollmentResponse")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EnrollmentResponse", *updateResource.Id)
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

	var parsed r5.EnrollmentResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return EnrollmentResponse or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEnrollmentResponse(patchResource *r5.EnrollmentResponse) (*r5.EnrollmentResponse, error) {
	if patchResource == nil {
		return nil, errors.New("PatchEnrollmentResponse given nil EnrollmentResponse")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchEnrollmentResponse given EnrollmentResponse without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EnrollmentResponse", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchEnrollmentResponse error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchEnrollmentResponse error using patch representation as json")
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

	var parsed r5.EnrollmentResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete EnrollmentResponse and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEnrollmentResponse(deleteResource *r5.EnrollmentResponse) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEnrollmentResponse given nil EnrollmentResponse")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEnrollmentResponse given EnrollmentResponse with nil Id (Id required to delete)")
	}
	return fc.DeleteEnrollmentResponseById(*deleteResource.Id)
}

// delete EnrollmentResponse with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEnrollmentResponseById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EnrollmentResponse", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create EpisodeOfCare, return id of created EpisodeOfCare or OperationOutcome error
func (fc *FhirClient) CreateEpisodeOfCare(createResource *r5.EpisodeOfCare) (*r5.EpisodeOfCare, error) {
	if createResource == nil {
		return nil, errors.New("CreateEpisodeOfCare called with nil EpisodeOfCare")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EpisodeOfCare")
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

	var parsed r5.EpisodeOfCare
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read EpisodeOfCare from fhir server by id, return EpisodeOfCare or OperationOutcome error
func (fc *FhirClient) ReadEpisodeOfCare(id string) (*r5.EpisodeOfCare, error) {
	if id == "" {
		return nil, errors.New("ReadEpisodeOfCare given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EpisodeOfCare", id)
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

	var parsed r5.EpisodeOfCare
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace EpisodeOfCare if exists on server, else create new EpisodeOfCare with given id, return EpisodeOfCare or OperationOutcome error
func (fc *FhirClient) UpdateEpisodeOfCare(updateResource *r5.EpisodeOfCare) (*r5.EpisodeOfCare, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateEpisodeOfCare called with nil EpisodeOfCare")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EpisodeOfCare", *updateResource.Id)
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

	var parsed r5.EpisodeOfCare
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return EpisodeOfCare or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEpisodeOfCare(patchResource *r5.EpisodeOfCare) (*r5.EpisodeOfCare, error) {
	if patchResource == nil {
		return nil, errors.New("PatchEpisodeOfCare given nil EpisodeOfCare")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchEpisodeOfCare given EpisodeOfCare without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EpisodeOfCare", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchEpisodeOfCare error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchEpisodeOfCare error using patch representation as json")
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

	var parsed r5.EpisodeOfCare
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete EpisodeOfCare and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEpisodeOfCare(deleteResource *r5.EpisodeOfCare) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEpisodeOfCare given nil EpisodeOfCare")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEpisodeOfCare given EpisodeOfCare with nil Id (Id required to delete)")
	}
	return fc.DeleteEpisodeOfCareById(*deleteResource.Id)
}

// delete EpisodeOfCare with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEpisodeOfCareById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EpisodeOfCare", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create EventDefinition, return id of created EventDefinition or OperationOutcome error
func (fc *FhirClient) CreateEventDefinition(createResource *r5.EventDefinition) (*r5.EventDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateEventDefinition called with nil EventDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EventDefinition")
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

	var parsed r5.EventDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read EventDefinition from fhir server by id, return EventDefinition or OperationOutcome error
func (fc *FhirClient) ReadEventDefinition(id string) (*r5.EventDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadEventDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EventDefinition", id)
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

	var parsed r5.EventDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace EventDefinition if exists on server, else create new EventDefinition with given id, return EventDefinition or OperationOutcome error
func (fc *FhirClient) UpdateEventDefinition(updateResource *r5.EventDefinition) (*r5.EventDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateEventDefinition called with nil EventDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EventDefinition", *updateResource.Id)
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

	var parsed r5.EventDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return EventDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEventDefinition(patchResource *r5.EventDefinition) (*r5.EventDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchEventDefinition given nil EventDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchEventDefinition given EventDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EventDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchEventDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchEventDefinition error using patch representation as json")
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

	var parsed r5.EventDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete EventDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEventDefinition(deleteResource *r5.EventDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEventDefinition given nil EventDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEventDefinition given EventDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteEventDefinitionById(*deleteResource.Id)
}

// delete EventDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEventDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EventDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Evidence, return id of created Evidence or OperationOutcome error
func (fc *FhirClient) CreateEvidence(createResource *r5.Evidence) (*r5.Evidence, error) {
	if createResource == nil {
		return nil, errors.New("CreateEvidence called with nil Evidence")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Evidence")
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

	var parsed r5.Evidence
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Evidence from fhir server by id, return Evidence or OperationOutcome error
func (fc *FhirClient) ReadEvidence(id string) (*r5.Evidence, error) {
	if id == "" {
		return nil, errors.New("ReadEvidence given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Evidence", id)
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

	var parsed r5.Evidence
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Evidence if exists on server, else create new Evidence with given id, return Evidence or OperationOutcome error
func (fc *FhirClient) UpdateEvidence(updateResource *r5.Evidence) (*r5.Evidence, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateEvidence called with nil Evidence")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Evidence", *updateResource.Id)
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

	var parsed r5.Evidence
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Evidence or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEvidence(patchResource *r5.Evidence) (*r5.Evidence, error) {
	if patchResource == nil {
		return nil, errors.New("PatchEvidence given nil Evidence")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchEvidence given Evidence without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Evidence", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchEvidence error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchEvidence error using patch representation as json")
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

	var parsed r5.Evidence
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Evidence and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEvidence(deleteResource *r5.Evidence) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEvidence given nil Evidence")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEvidence given Evidence with nil Id (Id required to delete)")
	}
	return fc.DeleteEvidenceById(*deleteResource.Id)
}

// delete Evidence with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEvidenceById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Evidence", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create EvidenceReport, return id of created EvidenceReport or OperationOutcome error
func (fc *FhirClient) CreateEvidenceReport(createResource *r5.EvidenceReport) (*r5.EvidenceReport, error) {
	if createResource == nil {
		return nil, errors.New("CreateEvidenceReport called with nil EvidenceReport")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EvidenceReport")
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

	var parsed r5.EvidenceReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read EvidenceReport from fhir server by id, return EvidenceReport or OperationOutcome error
func (fc *FhirClient) ReadEvidenceReport(id string) (*r5.EvidenceReport, error) {
	if id == "" {
		return nil, errors.New("ReadEvidenceReport given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EvidenceReport", id)
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

	var parsed r5.EvidenceReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace EvidenceReport if exists on server, else create new EvidenceReport with given id, return EvidenceReport or OperationOutcome error
func (fc *FhirClient) UpdateEvidenceReport(updateResource *r5.EvidenceReport) (*r5.EvidenceReport, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateEvidenceReport called with nil EvidenceReport")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EvidenceReport", *updateResource.Id)
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

	var parsed r5.EvidenceReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return EvidenceReport or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEvidenceReport(patchResource *r5.EvidenceReport) (*r5.EvidenceReport, error) {
	if patchResource == nil {
		return nil, errors.New("PatchEvidenceReport given nil EvidenceReport")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchEvidenceReport given EvidenceReport without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EvidenceReport", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchEvidenceReport error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchEvidenceReport error using patch representation as json")
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

	var parsed r5.EvidenceReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete EvidenceReport and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEvidenceReport(deleteResource *r5.EvidenceReport) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEvidenceReport given nil EvidenceReport")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEvidenceReport given EvidenceReport with nil Id (Id required to delete)")
	}
	return fc.DeleteEvidenceReportById(*deleteResource.Id)
}

// delete EvidenceReport with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEvidenceReportById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EvidenceReport", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create EvidenceVariable, return id of created EvidenceVariable or OperationOutcome error
func (fc *FhirClient) CreateEvidenceVariable(createResource *r5.EvidenceVariable) (*r5.EvidenceVariable, error) {
	if createResource == nil {
		return nil, errors.New("CreateEvidenceVariable called with nil EvidenceVariable")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EvidenceVariable")
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

	var parsed r5.EvidenceVariable
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read EvidenceVariable from fhir server by id, return EvidenceVariable or OperationOutcome error
func (fc *FhirClient) ReadEvidenceVariable(id string) (*r5.EvidenceVariable, error) {
	if id == "" {
		return nil, errors.New("ReadEvidenceVariable given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EvidenceVariable", id)
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

	var parsed r5.EvidenceVariable
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace EvidenceVariable if exists on server, else create new EvidenceVariable with given id, return EvidenceVariable or OperationOutcome error
func (fc *FhirClient) UpdateEvidenceVariable(updateResource *r5.EvidenceVariable) (*r5.EvidenceVariable, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateEvidenceVariable called with nil EvidenceVariable")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EvidenceVariable", *updateResource.Id)
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

	var parsed r5.EvidenceVariable
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return EvidenceVariable or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEvidenceVariable(patchResource *r5.EvidenceVariable) (*r5.EvidenceVariable, error) {
	if patchResource == nil {
		return nil, errors.New("PatchEvidenceVariable given nil EvidenceVariable")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchEvidenceVariable given EvidenceVariable without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EvidenceVariable", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchEvidenceVariable error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchEvidenceVariable error using patch representation as json")
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

	var parsed r5.EvidenceVariable
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete EvidenceVariable and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEvidenceVariable(deleteResource *r5.EvidenceVariable) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEvidenceVariable given nil EvidenceVariable")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEvidenceVariable given EvidenceVariable with nil Id (Id required to delete)")
	}
	return fc.DeleteEvidenceVariableById(*deleteResource.Id)
}

// delete EvidenceVariable with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEvidenceVariableById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EvidenceVariable", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ExampleScenario, return id of created ExampleScenario or OperationOutcome error
func (fc *FhirClient) CreateExampleScenario(createResource *r5.ExampleScenario) (*r5.ExampleScenario, error) {
	if createResource == nil {
		return nil, errors.New("CreateExampleScenario called with nil ExampleScenario")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ExampleScenario")
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

	var parsed r5.ExampleScenario
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ExampleScenario from fhir server by id, return ExampleScenario or OperationOutcome error
func (fc *FhirClient) ReadExampleScenario(id string) (*r5.ExampleScenario, error) {
	if id == "" {
		return nil, errors.New("ReadExampleScenario given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ExampleScenario", id)
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

	var parsed r5.ExampleScenario
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ExampleScenario if exists on server, else create new ExampleScenario with given id, return ExampleScenario or OperationOutcome error
func (fc *FhirClient) UpdateExampleScenario(updateResource *r5.ExampleScenario) (*r5.ExampleScenario, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateExampleScenario called with nil ExampleScenario")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ExampleScenario", *updateResource.Id)
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

	var parsed r5.ExampleScenario
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ExampleScenario or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchExampleScenario(patchResource *r5.ExampleScenario) (*r5.ExampleScenario, error) {
	if patchResource == nil {
		return nil, errors.New("PatchExampleScenario given nil ExampleScenario")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchExampleScenario given ExampleScenario without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ExampleScenario", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchExampleScenario error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchExampleScenario error using patch representation as json")
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

	var parsed r5.ExampleScenario
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ExampleScenario and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteExampleScenario(deleteResource *r5.ExampleScenario) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteExampleScenario given nil ExampleScenario")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteExampleScenario given ExampleScenario with nil Id (Id required to delete)")
	}
	return fc.DeleteExampleScenarioById(*deleteResource.Id)
}

// delete ExampleScenario with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteExampleScenarioById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ExampleScenario", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ExplanationOfBenefit, return id of created ExplanationOfBenefit or OperationOutcome error
func (fc *FhirClient) CreateExplanationOfBenefit(createResource *r5.ExplanationOfBenefit) (*r5.ExplanationOfBenefit, error) {
	if createResource == nil {
		return nil, errors.New("CreateExplanationOfBenefit called with nil ExplanationOfBenefit")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ExplanationOfBenefit")
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

	var parsed r5.ExplanationOfBenefit
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ExplanationOfBenefit from fhir server by id, return ExplanationOfBenefit or OperationOutcome error
func (fc *FhirClient) ReadExplanationOfBenefit(id string) (*r5.ExplanationOfBenefit, error) {
	if id == "" {
		return nil, errors.New("ReadExplanationOfBenefit given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ExplanationOfBenefit", id)
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

	var parsed r5.ExplanationOfBenefit
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ExplanationOfBenefit if exists on server, else create new ExplanationOfBenefit with given id, return ExplanationOfBenefit or OperationOutcome error
func (fc *FhirClient) UpdateExplanationOfBenefit(updateResource *r5.ExplanationOfBenefit) (*r5.ExplanationOfBenefit, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateExplanationOfBenefit called with nil ExplanationOfBenefit")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ExplanationOfBenefit", *updateResource.Id)
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

	var parsed r5.ExplanationOfBenefit
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ExplanationOfBenefit or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchExplanationOfBenefit(patchResource *r5.ExplanationOfBenefit) (*r5.ExplanationOfBenefit, error) {
	if patchResource == nil {
		return nil, errors.New("PatchExplanationOfBenefit given nil ExplanationOfBenefit")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchExplanationOfBenefit given ExplanationOfBenefit without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ExplanationOfBenefit", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchExplanationOfBenefit error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchExplanationOfBenefit error using patch representation as json")
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

	var parsed r5.ExplanationOfBenefit
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ExplanationOfBenefit and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteExplanationOfBenefit(deleteResource *r5.ExplanationOfBenefit) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteExplanationOfBenefit given nil ExplanationOfBenefit")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteExplanationOfBenefit given ExplanationOfBenefit with nil Id (Id required to delete)")
	}
	return fc.DeleteExplanationOfBenefitById(*deleteResource.Id)
}

// delete ExplanationOfBenefit with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteExplanationOfBenefitById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ExplanationOfBenefit", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create FamilyMemberHistory, return id of created FamilyMemberHistory or OperationOutcome error
func (fc *FhirClient) CreateFamilyMemberHistory(createResource *r5.FamilyMemberHistory) (*r5.FamilyMemberHistory, error) {
	if createResource == nil {
		return nil, errors.New("CreateFamilyMemberHistory called with nil FamilyMemberHistory")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "FamilyMemberHistory")
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

	var parsed r5.FamilyMemberHistory
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read FamilyMemberHistory from fhir server by id, return FamilyMemberHistory or OperationOutcome error
func (fc *FhirClient) ReadFamilyMemberHistory(id string) (*r5.FamilyMemberHistory, error) {
	if id == "" {
		return nil, errors.New("ReadFamilyMemberHistory given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "FamilyMemberHistory", id)
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

	var parsed r5.FamilyMemberHistory
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace FamilyMemberHistory if exists on server, else create new FamilyMemberHistory with given id, return FamilyMemberHistory or OperationOutcome error
func (fc *FhirClient) UpdateFamilyMemberHistory(updateResource *r5.FamilyMemberHistory) (*r5.FamilyMemberHistory, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateFamilyMemberHistory called with nil FamilyMemberHistory")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "FamilyMemberHistory", *updateResource.Id)
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

	var parsed r5.FamilyMemberHistory
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return FamilyMemberHistory or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchFamilyMemberHistory(patchResource *r5.FamilyMemberHistory) (*r5.FamilyMemberHistory, error) {
	if patchResource == nil {
		return nil, errors.New("PatchFamilyMemberHistory given nil FamilyMemberHistory")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchFamilyMemberHistory given FamilyMemberHistory without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "FamilyMemberHistory", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchFamilyMemberHistory error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchFamilyMemberHistory error using patch representation as json")
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

	var parsed r5.FamilyMemberHistory
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete FamilyMemberHistory and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteFamilyMemberHistory(deleteResource *r5.FamilyMemberHistory) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteFamilyMemberHistory given nil FamilyMemberHistory")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteFamilyMemberHistory given FamilyMemberHistory with nil Id (Id required to delete)")
	}
	return fc.DeleteFamilyMemberHistoryById(*deleteResource.Id)
}

// delete FamilyMemberHistory with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteFamilyMemberHistoryById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "FamilyMemberHistory", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Flag, return id of created Flag or OperationOutcome error
func (fc *FhirClient) CreateFlag(createResource *r5.Flag) (*r5.Flag, error) {
	if createResource == nil {
		return nil, errors.New("CreateFlag called with nil Flag")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Flag")
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

	var parsed r5.Flag
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Flag from fhir server by id, return Flag or OperationOutcome error
func (fc *FhirClient) ReadFlag(id string) (*r5.Flag, error) {
	if id == "" {
		return nil, errors.New("ReadFlag given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Flag", id)
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

	var parsed r5.Flag
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Flag if exists on server, else create new Flag with given id, return Flag or OperationOutcome error
func (fc *FhirClient) UpdateFlag(updateResource *r5.Flag) (*r5.Flag, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateFlag called with nil Flag")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Flag", *updateResource.Id)
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

	var parsed r5.Flag
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Flag or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchFlag(patchResource *r5.Flag) (*r5.Flag, error) {
	if patchResource == nil {
		return nil, errors.New("PatchFlag given nil Flag")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchFlag given Flag without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Flag", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchFlag error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchFlag error using patch representation as json")
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

	var parsed r5.Flag
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Flag and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteFlag(deleteResource *r5.Flag) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteFlag given nil Flag")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteFlag given Flag with nil Id (Id required to delete)")
	}
	return fc.DeleteFlagById(*deleteResource.Id)
}

// delete Flag with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteFlagById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Flag", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create FormularyItem, return id of created FormularyItem or OperationOutcome error
func (fc *FhirClient) CreateFormularyItem(createResource *r5.FormularyItem) (*r5.FormularyItem, error) {
	if createResource == nil {
		return nil, errors.New("CreateFormularyItem called with nil FormularyItem")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "FormularyItem")
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

	var parsed r5.FormularyItem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read FormularyItem from fhir server by id, return FormularyItem or OperationOutcome error
func (fc *FhirClient) ReadFormularyItem(id string) (*r5.FormularyItem, error) {
	if id == "" {
		return nil, errors.New("ReadFormularyItem given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "FormularyItem", id)
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

	var parsed r5.FormularyItem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace FormularyItem if exists on server, else create new FormularyItem with given id, return FormularyItem or OperationOutcome error
func (fc *FhirClient) UpdateFormularyItem(updateResource *r5.FormularyItem) (*r5.FormularyItem, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateFormularyItem called with nil FormularyItem")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "FormularyItem", *updateResource.Id)
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

	var parsed r5.FormularyItem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return FormularyItem or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchFormularyItem(patchResource *r5.FormularyItem) (*r5.FormularyItem, error) {
	if patchResource == nil {
		return nil, errors.New("PatchFormularyItem given nil FormularyItem")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchFormularyItem given FormularyItem without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "FormularyItem", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchFormularyItem error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchFormularyItem error using patch representation as json")
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

	var parsed r5.FormularyItem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete FormularyItem and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteFormularyItem(deleteResource *r5.FormularyItem) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteFormularyItem given nil FormularyItem")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteFormularyItem given FormularyItem with nil Id (Id required to delete)")
	}
	return fc.DeleteFormularyItemById(*deleteResource.Id)
}

// delete FormularyItem with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteFormularyItemById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "FormularyItem", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create GenomicStudy, return id of created GenomicStudy or OperationOutcome error
func (fc *FhirClient) CreateGenomicStudy(createResource *r5.GenomicStudy) (*r5.GenomicStudy, error) {
	if createResource == nil {
		return nil, errors.New("CreateGenomicStudy called with nil GenomicStudy")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GenomicStudy")
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

	var parsed r5.GenomicStudy
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read GenomicStudy from fhir server by id, return GenomicStudy or OperationOutcome error
func (fc *FhirClient) ReadGenomicStudy(id string) (*r5.GenomicStudy, error) {
	if id == "" {
		return nil, errors.New("ReadGenomicStudy given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GenomicStudy", id)
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

	var parsed r5.GenomicStudy
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace GenomicStudy if exists on server, else create new GenomicStudy with given id, return GenomicStudy or OperationOutcome error
func (fc *FhirClient) UpdateGenomicStudy(updateResource *r5.GenomicStudy) (*r5.GenomicStudy, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateGenomicStudy called with nil GenomicStudy")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GenomicStudy", *updateResource.Id)
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

	var parsed r5.GenomicStudy
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return GenomicStudy or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchGenomicStudy(patchResource *r5.GenomicStudy) (*r5.GenomicStudy, error) {
	if patchResource == nil {
		return nil, errors.New("PatchGenomicStudy given nil GenomicStudy")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchGenomicStudy given GenomicStudy without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GenomicStudy", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchGenomicStudy error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchGenomicStudy error using patch representation as json")
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

	var parsed r5.GenomicStudy
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete GenomicStudy and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGenomicStudy(deleteResource *r5.GenomicStudy) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteGenomicStudy given nil GenomicStudy")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteGenomicStudy given GenomicStudy with nil Id (Id required to delete)")
	}
	return fc.DeleteGenomicStudyById(*deleteResource.Id)
}

// delete GenomicStudy with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGenomicStudyById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GenomicStudy", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Goal, return id of created Goal or OperationOutcome error
func (fc *FhirClient) CreateGoal(createResource *r5.Goal) (*r5.Goal, error) {
	if createResource == nil {
		return nil, errors.New("CreateGoal called with nil Goal")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Goal")
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

	var parsed r5.Goal
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Goal from fhir server by id, return Goal or OperationOutcome error
func (fc *FhirClient) ReadGoal(id string) (*r5.Goal, error) {
	if id == "" {
		return nil, errors.New("ReadGoal given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Goal", id)
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

	var parsed r5.Goal
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Goal if exists on server, else create new Goal with given id, return Goal or OperationOutcome error
func (fc *FhirClient) UpdateGoal(updateResource *r5.Goal) (*r5.Goal, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateGoal called with nil Goal")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Goal", *updateResource.Id)
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

	var parsed r5.Goal
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Goal or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchGoal(patchResource *r5.Goal) (*r5.Goal, error) {
	if patchResource == nil {
		return nil, errors.New("PatchGoal given nil Goal")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchGoal given Goal without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Goal", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchGoal error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchGoal error using patch representation as json")
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

	var parsed r5.Goal
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Goal and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGoal(deleteResource *r5.Goal) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteGoal given nil Goal")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteGoal given Goal with nil Id (Id required to delete)")
	}
	return fc.DeleteGoalById(*deleteResource.Id)
}

// delete Goal with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGoalById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Goal", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create GraphDefinition, return id of created GraphDefinition or OperationOutcome error
func (fc *FhirClient) CreateGraphDefinition(createResource *r5.GraphDefinition) (*r5.GraphDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateGraphDefinition called with nil GraphDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GraphDefinition")
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

	var parsed r5.GraphDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read GraphDefinition from fhir server by id, return GraphDefinition or OperationOutcome error
func (fc *FhirClient) ReadGraphDefinition(id string) (*r5.GraphDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadGraphDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GraphDefinition", id)
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

	var parsed r5.GraphDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace GraphDefinition if exists on server, else create new GraphDefinition with given id, return GraphDefinition or OperationOutcome error
func (fc *FhirClient) UpdateGraphDefinition(updateResource *r5.GraphDefinition) (*r5.GraphDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateGraphDefinition called with nil GraphDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GraphDefinition", *updateResource.Id)
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

	var parsed r5.GraphDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return GraphDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchGraphDefinition(patchResource *r5.GraphDefinition) (*r5.GraphDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchGraphDefinition given nil GraphDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchGraphDefinition given GraphDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GraphDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchGraphDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchGraphDefinition error using patch representation as json")
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

	var parsed r5.GraphDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete GraphDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGraphDefinition(deleteResource *r5.GraphDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteGraphDefinition given nil GraphDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteGraphDefinition given GraphDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteGraphDefinitionById(*deleteResource.Id)
}

// delete GraphDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGraphDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GraphDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Group, return id of created Group or OperationOutcome error
func (fc *FhirClient) CreateGroup(createResource *r5.Group) (*r5.Group, error) {
	if createResource == nil {
		return nil, errors.New("CreateGroup called with nil Group")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Group")
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

	var parsed r5.Group
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Group from fhir server by id, return Group or OperationOutcome error
func (fc *FhirClient) ReadGroup(id string) (*r5.Group, error) {
	if id == "" {
		return nil, errors.New("ReadGroup given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Group", id)
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

	var parsed r5.Group
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Group if exists on server, else create new Group with given id, return Group or OperationOutcome error
func (fc *FhirClient) UpdateGroup(updateResource *r5.Group) (*r5.Group, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateGroup called with nil Group")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Group", *updateResource.Id)
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

	var parsed r5.Group
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Group or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchGroup(patchResource *r5.Group) (*r5.Group, error) {
	if patchResource == nil {
		return nil, errors.New("PatchGroup given nil Group")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchGroup given Group without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Group", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchGroup error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchGroup error using patch representation as json")
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

	var parsed r5.Group
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Group and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGroup(deleteResource *r5.Group) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteGroup given nil Group")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteGroup given Group with nil Id (Id required to delete)")
	}
	return fc.DeleteGroupById(*deleteResource.Id)
}

// delete Group with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGroupById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Group", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create GuidanceResponse, return id of created GuidanceResponse or OperationOutcome error
func (fc *FhirClient) CreateGuidanceResponse(createResource *r5.GuidanceResponse) (*r5.GuidanceResponse, error) {
	if createResource == nil {
		return nil, errors.New("CreateGuidanceResponse called with nil GuidanceResponse")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GuidanceResponse")
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

	var parsed r5.GuidanceResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read GuidanceResponse from fhir server by id, return GuidanceResponse or OperationOutcome error
func (fc *FhirClient) ReadGuidanceResponse(id string) (*r5.GuidanceResponse, error) {
	if id == "" {
		return nil, errors.New("ReadGuidanceResponse given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GuidanceResponse", id)
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

	var parsed r5.GuidanceResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace GuidanceResponse if exists on server, else create new GuidanceResponse with given id, return GuidanceResponse or OperationOutcome error
func (fc *FhirClient) UpdateGuidanceResponse(updateResource *r5.GuidanceResponse) (*r5.GuidanceResponse, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateGuidanceResponse called with nil GuidanceResponse")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GuidanceResponse", *updateResource.Id)
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

	var parsed r5.GuidanceResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return GuidanceResponse or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchGuidanceResponse(patchResource *r5.GuidanceResponse) (*r5.GuidanceResponse, error) {
	if patchResource == nil {
		return nil, errors.New("PatchGuidanceResponse given nil GuidanceResponse")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchGuidanceResponse given GuidanceResponse without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GuidanceResponse", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchGuidanceResponse error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchGuidanceResponse error using patch representation as json")
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

	var parsed r5.GuidanceResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete GuidanceResponse and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGuidanceResponse(deleteResource *r5.GuidanceResponse) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteGuidanceResponse given nil GuidanceResponse")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteGuidanceResponse given GuidanceResponse with nil Id (Id required to delete)")
	}
	return fc.DeleteGuidanceResponseById(*deleteResource.Id)
}

// delete GuidanceResponse with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGuidanceResponseById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GuidanceResponse", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create HealthcareService, return id of created HealthcareService or OperationOutcome error
func (fc *FhirClient) CreateHealthcareService(createResource *r5.HealthcareService) (*r5.HealthcareService, error) {
	if createResource == nil {
		return nil, errors.New("CreateHealthcareService called with nil HealthcareService")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "HealthcareService")
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

	var parsed r5.HealthcareService
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read HealthcareService from fhir server by id, return HealthcareService or OperationOutcome error
func (fc *FhirClient) ReadHealthcareService(id string) (*r5.HealthcareService, error) {
	if id == "" {
		return nil, errors.New("ReadHealthcareService given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "HealthcareService", id)
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

	var parsed r5.HealthcareService
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace HealthcareService if exists on server, else create new HealthcareService with given id, return HealthcareService or OperationOutcome error
func (fc *FhirClient) UpdateHealthcareService(updateResource *r5.HealthcareService) (*r5.HealthcareService, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateHealthcareService called with nil HealthcareService")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "HealthcareService", *updateResource.Id)
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

	var parsed r5.HealthcareService
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return HealthcareService or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchHealthcareService(patchResource *r5.HealthcareService) (*r5.HealthcareService, error) {
	if patchResource == nil {
		return nil, errors.New("PatchHealthcareService given nil HealthcareService")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchHealthcareService given HealthcareService without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "HealthcareService", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchHealthcareService error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchHealthcareService error using patch representation as json")
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

	var parsed r5.HealthcareService
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete HealthcareService and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteHealthcareService(deleteResource *r5.HealthcareService) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteHealthcareService given nil HealthcareService")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteHealthcareService given HealthcareService with nil Id (Id required to delete)")
	}
	return fc.DeleteHealthcareServiceById(*deleteResource.Id)
}

// delete HealthcareService with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteHealthcareServiceById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "HealthcareService", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ImagingSelection, return id of created ImagingSelection or OperationOutcome error
func (fc *FhirClient) CreateImagingSelection(createResource *r5.ImagingSelection) (*r5.ImagingSelection, error) {
	if createResource == nil {
		return nil, errors.New("CreateImagingSelection called with nil ImagingSelection")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImagingSelection")
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

	var parsed r5.ImagingSelection
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ImagingSelection from fhir server by id, return ImagingSelection or OperationOutcome error
func (fc *FhirClient) ReadImagingSelection(id string) (*r5.ImagingSelection, error) {
	if id == "" {
		return nil, errors.New("ReadImagingSelection given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImagingSelection", id)
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

	var parsed r5.ImagingSelection
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ImagingSelection if exists on server, else create new ImagingSelection with given id, return ImagingSelection or OperationOutcome error
func (fc *FhirClient) UpdateImagingSelection(updateResource *r5.ImagingSelection) (*r5.ImagingSelection, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateImagingSelection called with nil ImagingSelection")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImagingSelection", *updateResource.Id)
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

	var parsed r5.ImagingSelection
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ImagingSelection or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchImagingSelection(patchResource *r5.ImagingSelection) (*r5.ImagingSelection, error) {
	if patchResource == nil {
		return nil, errors.New("PatchImagingSelection given nil ImagingSelection")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchImagingSelection given ImagingSelection without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImagingSelection", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchImagingSelection error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchImagingSelection error using patch representation as json")
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

	var parsed r5.ImagingSelection
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ImagingSelection and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImagingSelection(deleteResource *r5.ImagingSelection) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteImagingSelection given nil ImagingSelection")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteImagingSelection given ImagingSelection with nil Id (Id required to delete)")
	}
	return fc.DeleteImagingSelectionById(*deleteResource.Id)
}

// delete ImagingSelection with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImagingSelectionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImagingSelection", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ImagingStudy, return id of created ImagingStudy or OperationOutcome error
func (fc *FhirClient) CreateImagingStudy(createResource *r5.ImagingStudy) (*r5.ImagingStudy, error) {
	if createResource == nil {
		return nil, errors.New("CreateImagingStudy called with nil ImagingStudy")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImagingStudy")
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

	var parsed r5.ImagingStudy
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ImagingStudy from fhir server by id, return ImagingStudy or OperationOutcome error
func (fc *FhirClient) ReadImagingStudy(id string) (*r5.ImagingStudy, error) {
	if id == "" {
		return nil, errors.New("ReadImagingStudy given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImagingStudy", id)
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

	var parsed r5.ImagingStudy
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ImagingStudy if exists on server, else create new ImagingStudy with given id, return ImagingStudy or OperationOutcome error
func (fc *FhirClient) UpdateImagingStudy(updateResource *r5.ImagingStudy) (*r5.ImagingStudy, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateImagingStudy called with nil ImagingStudy")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImagingStudy", *updateResource.Id)
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

	var parsed r5.ImagingStudy
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ImagingStudy or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchImagingStudy(patchResource *r5.ImagingStudy) (*r5.ImagingStudy, error) {
	if patchResource == nil {
		return nil, errors.New("PatchImagingStudy given nil ImagingStudy")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchImagingStudy given ImagingStudy without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImagingStudy", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchImagingStudy error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchImagingStudy error using patch representation as json")
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

	var parsed r5.ImagingStudy
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ImagingStudy and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImagingStudy(deleteResource *r5.ImagingStudy) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteImagingStudy given nil ImagingStudy")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteImagingStudy given ImagingStudy with nil Id (Id required to delete)")
	}
	return fc.DeleteImagingStudyById(*deleteResource.Id)
}

// delete ImagingStudy with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImagingStudyById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImagingStudy", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Immunization, return id of created Immunization or OperationOutcome error
func (fc *FhirClient) CreateImmunization(createResource *r5.Immunization) (*r5.Immunization, error) {
	if createResource == nil {
		return nil, errors.New("CreateImmunization called with nil Immunization")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Immunization")
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

	var parsed r5.Immunization
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Immunization from fhir server by id, return Immunization or OperationOutcome error
func (fc *FhirClient) ReadImmunization(id string) (*r5.Immunization, error) {
	if id == "" {
		return nil, errors.New("ReadImmunization given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Immunization", id)
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

	var parsed r5.Immunization
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Immunization if exists on server, else create new Immunization with given id, return Immunization or OperationOutcome error
func (fc *FhirClient) UpdateImmunization(updateResource *r5.Immunization) (*r5.Immunization, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateImmunization called with nil Immunization")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Immunization", *updateResource.Id)
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

	var parsed r5.Immunization
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Immunization or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchImmunization(patchResource *r5.Immunization) (*r5.Immunization, error) {
	if patchResource == nil {
		return nil, errors.New("PatchImmunization given nil Immunization")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchImmunization given Immunization without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Immunization", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchImmunization error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchImmunization error using patch representation as json")
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

	var parsed r5.Immunization
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Immunization and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImmunization(deleteResource *r5.Immunization) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteImmunization given nil Immunization")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteImmunization given Immunization with nil Id (Id required to delete)")
	}
	return fc.DeleteImmunizationById(*deleteResource.Id)
}

// delete Immunization with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImmunizationById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Immunization", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ImmunizationEvaluation, return id of created ImmunizationEvaluation or OperationOutcome error
func (fc *FhirClient) CreateImmunizationEvaluation(createResource *r5.ImmunizationEvaluation) (*r5.ImmunizationEvaluation, error) {
	if createResource == nil {
		return nil, errors.New("CreateImmunizationEvaluation called with nil ImmunizationEvaluation")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImmunizationEvaluation")
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

	var parsed r5.ImmunizationEvaluation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ImmunizationEvaluation from fhir server by id, return ImmunizationEvaluation or OperationOutcome error
func (fc *FhirClient) ReadImmunizationEvaluation(id string) (*r5.ImmunizationEvaluation, error) {
	if id == "" {
		return nil, errors.New("ReadImmunizationEvaluation given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImmunizationEvaluation", id)
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

	var parsed r5.ImmunizationEvaluation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ImmunizationEvaluation if exists on server, else create new ImmunizationEvaluation with given id, return ImmunizationEvaluation or OperationOutcome error
func (fc *FhirClient) UpdateImmunizationEvaluation(updateResource *r5.ImmunizationEvaluation) (*r5.ImmunizationEvaluation, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateImmunizationEvaluation called with nil ImmunizationEvaluation")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImmunizationEvaluation", *updateResource.Id)
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

	var parsed r5.ImmunizationEvaluation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ImmunizationEvaluation or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchImmunizationEvaluation(patchResource *r5.ImmunizationEvaluation) (*r5.ImmunizationEvaluation, error) {
	if patchResource == nil {
		return nil, errors.New("PatchImmunizationEvaluation given nil ImmunizationEvaluation")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchImmunizationEvaluation given ImmunizationEvaluation without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImmunizationEvaluation", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchImmunizationEvaluation error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchImmunizationEvaluation error using patch representation as json")
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

	var parsed r5.ImmunizationEvaluation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ImmunizationEvaluation and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImmunizationEvaluation(deleteResource *r5.ImmunizationEvaluation) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteImmunizationEvaluation given nil ImmunizationEvaluation")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteImmunizationEvaluation given ImmunizationEvaluation with nil Id (Id required to delete)")
	}
	return fc.DeleteImmunizationEvaluationById(*deleteResource.Id)
}

// delete ImmunizationEvaluation with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImmunizationEvaluationById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImmunizationEvaluation", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ImmunizationRecommendation, return id of created ImmunizationRecommendation or OperationOutcome error
func (fc *FhirClient) CreateImmunizationRecommendation(createResource *r5.ImmunizationRecommendation) (*r5.ImmunizationRecommendation, error) {
	if createResource == nil {
		return nil, errors.New("CreateImmunizationRecommendation called with nil ImmunizationRecommendation")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImmunizationRecommendation")
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

	var parsed r5.ImmunizationRecommendation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ImmunizationRecommendation from fhir server by id, return ImmunizationRecommendation or OperationOutcome error
func (fc *FhirClient) ReadImmunizationRecommendation(id string) (*r5.ImmunizationRecommendation, error) {
	if id == "" {
		return nil, errors.New("ReadImmunizationRecommendation given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImmunizationRecommendation", id)
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

	var parsed r5.ImmunizationRecommendation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ImmunizationRecommendation if exists on server, else create new ImmunizationRecommendation with given id, return ImmunizationRecommendation or OperationOutcome error
func (fc *FhirClient) UpdateImmunizationRecommendation(updateResource *r5.ImmunizationRecommendation) (*r5.ImmunizationRecommendation, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateImmunizationRecommendation called with nil ImmunizationRecommendation")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImmunizationRecommendation", *updateResource.Id)
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

	var parsed r5.ImmunizationRecommendation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ImmunizationRecommendation or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchImmunizationRecommendation(patchResource *r5.ImmunizationRecommendation) (*r5.ImmunizationRecommendation, error) {
	if patchResource == nil {
		return nil, errors.New("PatchImmunizationRecommendation given nil ImmunizationRecommendation")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchImmunizationRecommendation given ImmunizationRecommendation without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImmunizationRecommendation", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchImmunizationRecommendation error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchImmunizationRecommendation error using patch representation as json")
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

	var parsed r5.ImmunizationRecommendation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ImmunizationRecommendation and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImmunizationRecommendation(deleteResource *r5.ImmunizationRecommendation) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteImmunizationRecommendation given nil ImmunizationRecommendation")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteImmunizationRecommendation given ImmunizationRecommendation with nil Id (Id required to delete)")
	}
	return fc.DeleteImmunizationRecommendationById(*deleteResource.Id)
}

// delete ImmunizationRecommendation with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImmunizationRecommendationById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImmunizationRecommendation", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ImplementationGuide, return id of created ImplementationGuide or OperationOutcome error
func (fc *FhirClient) CreateImplementationGuide(createResource *r5.ImplementationGuide) (*r5.ImplementationGuide, error) {
	if createResource == nil {
		return nil, errors.New("CreateImplementationGuide called with nil ImplementationGuide")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImplementationGuide")
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

	var parsed r5.ImplementationGuide
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ImplementationGuide from fhir server by id, return ImplementationGuide or OperationOutcome error
func (fc *FhirClient) ReadImplementationGuide(id string) (*r5.ImplementationGuide, error) {
	if id == "" {
		return nil, errors.New("ReadImplementationGuide given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImplementationGuide", id)
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

	var parsed r5.ImplementationGuide
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ImplementationGuide if exists on server, else create new ImplementationGuide with given id, return ImplementationGuide or OperationOutcome error
func (fc *FhirClient) UpdateImplementationGuide(updateResource *r5.ImplementationGuide) (*r5.ImplementationGuide, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateImplementationGuide called with nil ImplementationGuide")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImplementationGuide", *updateResource.Id)
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

	var parsed r5.ImplementationGuide
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ImplementationGuide or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchImplementationGuide(patchResource *r5.ImplementationGuide) (*r5.ImplementationGuide, error) {
	if patchResource == nil {
		return nil, errors.New("PatchImplementationGuide given nil ImplementationGuide")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchImplementationGuide given ImplementationGuide without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImplementationGuide", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchImplementationGuide error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchImplementationGuide error using patch representation as json")
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

	var parsed r5.ImplementationGuide
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ImplementationGuide and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImplementationGuide(deleteResource *r5.ImplementationGuide) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteImplementationGuide given nil ImplementationGuide")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteImplementationGuide given ImplementationGuide with nil Id (Id required to delete)")
	}
	return fc.DeleteImplementationGuideById(*deleteResource.Id)
}

// delete ImplementationGuide with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImplementationGuideById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImplementationGuide", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Ingredient, return id of created Ingredient or OperationOutcome error
func (fc *FhirClient) CreateIngredient(createResource *r5.Ingredient) (*r5.Ingredient, error) {
	if createResource == nil {
		return nil, errors.New("CreateIngredient called with nil Ingredient")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Ingredient")
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

	var parsed r5.Ingredient
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Ingredient from fhir server by id, return Ingredient or OperationOutcome error
func (fc *FhirClient) ReadIngredient(id string) (*r5.Ingredient, error) {
	if id == "" {
		return nil, errors.New("ReadIngredient given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Ingredient", id)
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

	var parsed r5.Ingredient
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Ingredient if exists on server, else create new Ingredient with given id, return Ingredient or OperationOutcome error
func (fc *FhirClient) UpdateIngredient(updateResource *r5.Ingredient) (*r5.Ingredient, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateIngredient called with nil Ingredient")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Ingredient", *updateResource.Id)
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

	var parsed r5.Ingredient
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Ingredient or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchIngredient(patchResource *r5.Ingredient) (*r5.Ingredient, error) {
	if patchResource == nil {
		return nil, errors.New("PatchIngredient given nil Ingredient")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchIngredient given Ingredient without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Ingredient", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchIngredient error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchIngredient error using patch representation as json")
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

	var parsed r5.Ingredient
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Ingredient and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteIngredient(deleteResource *r5.Ingredient) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteIngredient given nil Ingredient")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteIngredient given Ingredient with nil Id (Id required to delete)")
	}
	return fc.DeleteIngredientById(*deleteResource.Id)
}

// delete Ingredient with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteIngredientById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Ingredient", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create InsurancePlan, return id of created InsurancePlan or OperationOutcome error
func (fc *FhirClient) CreateInsurancePlan(createResource *r5.InsurancePlan) (*r5.InsurancePlan, error) {
	if createResource == nil {
		return nil, errors.New("CreateInsurancePlan called with nil InsurancePlan")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "InsurancePlan")
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

	var parsed r5.InsurancePlan
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read InsurancePlan from fhir server by id, return InsurancePlan or OperationOutcome error
func (fc *FhirClient) ReadInsurancePlan(id string) (*r5.InsurancePlan, error) {
	if id == "" {
		return nil, errors.New("ReadInsurancePlan given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "InsurancePlan", id)
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

	var parsed r5.InsurancePlan
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace InsurancePlan if exists on server, else create new InsurancePlan with given id, return InsurancePlan or OperationOutcome error
func (fc *FhirClient) UpdateInsurancePlan(updateResource *r5.InsurancePlan) (*r5.InsurancePlan, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateInsurancePlan called with nil InsurancePlan")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "InsurancePlan", *updateResource.Id)
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

	var parsed r5.InsurancePlan
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return InsurancePlan or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchInsurancePlan(patchResource *r5.InsurancePlan) (*r5.InsurancePlan, error) {
	if patchResource == nil {
		return nil, errors.New("PatchInsurancePlan given nil InsurancePlan")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchInsurancePlan given InsurancePlan without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "InsurancePlan", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchInsurancePlan error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchInsurancePlan error using patch representation as json")
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

	var parsed r5.InsurancePlan
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete InsurancePlan and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteInsurancePlan(deleteResource *r5.InsurancePlan) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteInsurancePlan given nil InsurancePlan")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteInsurancePlan given InsurancePlan with nil Id (Id required to delete)")
	}
	return fc.DeleteInsurancePlanById(*deleteResource.Id)
}

// delete InsurancePlan with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteInsurancePlanById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "InsurancePlan", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create InventoryItem, return id of created InventoryItem or OperationOutcome error
func (fc *FhirClient) CreateInventoryItem(createResource *r5.InventoryItem) (*r5.InventoryItem, error) {
	if createResource == nil {
		return nil, errors.New("CreateInventoryItem called with nil InventoryItem")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "InventoryItem")
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

	var parsed r5.InventoryItem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read InventoryItem from fhir server by id, return InventoryItem or OperationOutcome error
func (fc *FhirClient) ReadInventoryItem(id string) (*r5.InventoryItem, error) {
	if id == "" {
		return nil, errors.New("ReadInventoryItem given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "InventoryItem", id)
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

	var parsed r5.InventoryItem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace InventoryItem if exists on server, else create new InventoryItem with given id, return InventoryItem or OperationOutcome error
func (fc *FhirClient) UpdateInventoryItem(updateResource *r5.InventoryItem) (*r5.InventoryItem, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateInventoryItem called with nil InventoryItem")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "InventoryItem", *updateResource.Id)
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

	var parsed r5.InventoryItem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return InventoryItem or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchInventoryItem(patchResource *r5.InventoryItem) (*r5.InventoryItem, error) {
	if patchResource == nil {
		return nil, errors.New("PatchInventoryItem given nil InventoryItem")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchInventoryItem given InventoryItem without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "InventoryItem", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchInventoryItem error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchInventoryItem error using patch representation as json")
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

	var parsed r5.InventoryItem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete InventoryItem and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteInventoryItem(deleteResource *r5.InventoryItem) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteInventoryItem given nil InventoryItem")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteInventoryItem given InventoryItem with nil Id (Id required to delete)")
	}
	return fc.DeleteInventoryItemById(*deleteResource.Id)
}

// delete InventoryItem with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteInventoryItemById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "InventoryItem", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create InventoryReport, return id of created InventoryReport or OperationOutcome error
func (fc *FhirClient) CreateInventoryReport(createResource *r5.InventoryReport) (*r5.InventoryReport, error) {
	if createResource == nil {
		return nil, errors.New("CreateInventoryReport called with nil InventoryReport")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "InventoryReport")
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

	var parsed r5.InventoryReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read InventoryReport from fhir server by id, return InventoryReport or OperationOutcome error
func (fc *FhirClient) ReadInventoryReport(id string) (*r5.InventoryReport, error) {
	if id == "" {
		return nil, errors.New("ReadInventoryReport given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "InventoryReport", id)
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

	var parsed r5.InventoryReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace InventoryReport if exists on server, else create new InventoryReport with given id, return InventoryReport or OperationOutcome error
func (fc *FhirClient) UpdateInventoryReport(updateResource *r5.InventoryReport) (*r5.InventoryReport, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateInventoryReport called with nil InventoryReport")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "InventoryReport", *updateResource.Id)
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

	var parsed r5.InventoryReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return InventoryReport or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchInventoryReport(patchResource *r5.InventoryReport) (*r5.InventoryReport, error) {
	if patchResource == nil {
		return nil, errors.New("PatchInventoryReport given nil InventoryReport")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchInventoryReport given InventoryReport without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "InventoryReport", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchInventoryReport error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchInventoryReport error using patch representation as json")
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

	var parsed r5.InventoryReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete InventoryReport and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteInventoryReport(deleteResource *r5.InventoryReport) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteInventoryReport given nil InventoryReport")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteInventoryReport given InventoryReport with nil Id (Id required to delete)")
	}
	return fc.DeleteInventoryReportById(*deleteResource.Id)
}

// delete InventoryReport with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteInventoryReportById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "InventoryReport", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Invoice, return id of created Invoice or OperationOutcome error
func (fc *FhirClient) CreateInvoice(createResource *r5.Invoice) (*r5.Invoice, error) {
	if createResource == nil {
		return nil, errors.New("CreateInvoice called with nil Invoice")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Invoice")
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

	var parsed r5.Invoice
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Invoice from fhir server by id, return Invoice or OperationOutcome error
func (fc *FhirClient) ReadInvoice(id string) (*r5.Invoice, error) {
	if id == "" {
		return nil, errors.New("ReadInvoice given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Invoice", id)
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

	var parsed r5.Invoice
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Invoice if exists on server, else create new Invoice with given id, return Invoice or OperationOutcome error
func (fc *FhirClient) UpdateInvoice(updateResource *r5.Invoice) (*r5.Invoice, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateInvoice called with nil Invoice")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Invoice", *updateResource.Id)
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

	var parsed r5.Invoice
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Invoice or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchInvoice(patchResource *r5.Invoice) (*r5.Invoice, error) {
	if patchResource == nil {
		return nil, errors.New("PatchInvoice given nil Invoice")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchInvoice given Invoice without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Invoice", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchInvoice error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchInvoice error using patch representation as json")
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

	var parsed r5.Invoice
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Invoice and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteInvoice(deleteResource *r5.Invoice) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteInvoice given nil Invoice")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteInvoice given Invoice with nil Id (Id required to delete)")
	}
	return fc.DeleteInvoiceById(*deleteResource.Id)
}

// delete Invoice with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteInvoiceById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Invoice", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Library, return id of created Library or OperationOutcome error
func (fc *FhirClient) CreateLibrary(createResource *r5.Library) (*r5.Library, error) {
	if createResource == nil {
		return nil, errors.New("CreateLibrary called with nil Library")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Library")
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

	var parsed r5.Library
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Library from fhir server by id, return Library or OperationOutcome error
func (fc *FhirClient) ReadLibrary(id string) (*r5.Library, error) {
	if id == "" {
		return nil, errors.New("ReadLibrary given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Library", id)
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

	var parsed r5.Library
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Library if exists on server, else create new Library with given id, return Library or OperationOutcome error
func (fc *FhirClient) UpdateLibrary(updateResource *r5.Library) (*r5.Library, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateLibrary called with nil Library")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Library", *updateResource.Id)
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

	var parsed r5.Library
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Library or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchLibrary(patchResource *r5.Library) (*r5.Library, error) {
	if patchResource == nil {
		return nil, errors.New("PatchLibrary given nil Library")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchLibrary given Library without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Library", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchLibrary error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchLibrary error using patch representation as json")
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

	var parsed r5.Library
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Library and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteLibrary(deleteResource *r5.Library) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteLibrary given nil Library")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteLibrary given Library with nil Id (Id required to delete)")
	}
	return fc.DeleteLibraryById(*deleteResource.Id)
}

// delete Library with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteLibraryById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Library", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Linkage, return id of created Linkage or OperationOutcome error
func (fc *FhirClient) CreateLinkage(createResource *r5.Linkage) (*r5.Linkage, error) {
	if createResource == nil {
		return nil, errors.New("CreateLinkage called with nil Linkage")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Linkage")
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

	var parsed r5.Linkage
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Linkage from fhir server by id, return Linkage or OperationOutcome error
func (fc *FhirClient) ReadLinkage(id string) (*r5.Linkage, error) {
	if id == "" {
		return nil, errors.New("ReadLinkage given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Linkage", id)
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

	var parsed r5.Linkage
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Linkage if exists on server, else create new Linkage with given id, return Linkage or OperationOutcome error
func (fc *FhirClient) UpdateLinkage(updateResource *r5.Linkage) (*r5.Linkage, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateLinkage called with nil Linkage")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Linkage", *updateResource.Id)
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

	var parsed r5.Linkage
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Linkage or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchLinkage(patchResource *r5.Linkage) (*r5.Linkage, error) {
	if patchResource == nil {
		return nil, errors.New("PatchLinkage given nil Linkage")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchLinkage given Linkage without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Linkage", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchLinkage error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchLinkage error using patch representation as json")
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

	var parsed r5.Linkage
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Linkage and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteLinkage(deleteResource *r5.Linkage) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteLinkage given nil Linkage")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteLinkage given Linkage with nil Id (Id required to delete)")
	}
	return fc.DeleteLinkageById(*deleteResource.Id)
}

// delete Linkage with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteLinkageById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Linkage", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create List, return id of created List or OperationOutcome error
func (fc *FhirClient) CreateList(createResource *r5.List) (*r5.List, error) {
	if createResource == nil {
		return nil, errors.New("CreateList called with nil List")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "List")
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

	var parsed r5.List
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read List from fhir server by id, return List or OperationOutcome error
func (fc *FhirClient) ReadList(id string) (*r5.List, error) {
	if id == "" {
		return nil, errors.New("ReadList given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "List", id)
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

	var parsed r5.List
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace List if exists on server, else create new List with given id, return List or OperationOutcome error
func (fc *FhirClient) UpdateList(updateResource *r5.List) (*r5.List, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateList called with nil List")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "List", *updateResource.Id)
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

	var parsed r5.List
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return List or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchList(patchResource *r5.List) (*r5.List, error) {
	if patchResource == nil {
		return nil, errors.New("PatchList given nil List")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchList given List without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "List", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchList error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchList error using patch representation as json")
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

	var parsed r5.List
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete List and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteList(deleteResource *r5.List) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteList given nil List")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteList given List with nil Id (Id required to delete)")
	}
	return fc.DeleteListById(*deleteResource.Id)
}

// delete List with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteListById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "List", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Location, return id of created Location or OperationOutcome error
func (fc *FhirClient) CreateLocation(createResource *r5.Location) (*r5.Location, error) {
	if createResource == nil {
		return nil, errors.New("CreateLocation called with nil Location")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Location")
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

	var parsed r5.Location
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Location from fhir server by id, return Location or OperationOutcome error
func (fc *FhirClient) ReadLocation(id string) (*r5.Location, error) {
	if id == "" {
		return nil, errors.New("ReadLocation given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Location", id)
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

	var parsed r5.Location
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Location if exists on server, else create new Location with given id, return Location or OperationOutcome error
func (fc *FhirClient) UpdateLocation(updateResource *r5.Location) (*r5.Location, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateLocation called with nil Location")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Location", *updateResource.Id)
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

	var parsed r5.Location
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Location or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchLocation(patchResource *r5.Location) (*r5.Location, error) {
	if patchResource == nil {
		return nil, errors.New("PatchLocation given nil Location")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchLocation given Location without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Location", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchLocation error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchLocation error using patch representation as json")
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

	var parsed r5.Location
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Location and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteLocation(deleteResource *r5.Location) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteLocation given nil Location")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteLocation given Location with nil Id (Id required to delete)")
	}
	return fc.DeleteLocationById(*deleteResource.Id)
}

// delete Location with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteLocationById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Location", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ManufacturedItemDefinition, return id of created ManufacturedItemDefinition or OperationOutcome error
func (fc *FhirClient) CreateManufacturedItemDefinition(createResource *r5.ManufacturedItemDefinition) (*r5.ManufacturedItemDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateManufacturedItemDefinition called with nil ManufacturedItemDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ManufacturedItemDefinition")
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

	var parsed r5.ManufacturedItemDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ManufacturedItemDefinition from fhir server by id, return ManufacturedItemDefinition or OperationOutcome error
func (fc *FhirClient) ReadManufacturedItemDefinition(id string) (*r5.ManufacturedItemDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadManufacturedItemDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ManufacturedItemDefinition", id)
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

	var parsed r5.ManufacturedItemDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ManufacturedItemDefinition if exists on server, else create new ManufacturedItemDefinition with given id, return ManufacturedItemDefinition or OperationOutcome error
func (fc *FhirClient) UpdateManufacturedItemDefinition(updateResource *r5.ManufacturedItemDefinition) (*r5.ManufacturedItemDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateManufacturedItemDefinition called with nil ManufacturedItemDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ManufacturedItemDefinition", *updateResource.Id)
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

	var parsed r5.ManufacturedItemDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ManufacturedItemDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchManufacturedItemDefinition(patchResource *r5.ManufacturedItemDefinition) (*r5.ManufacturedItemDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchManufacturedItemDefinition given nil ManufacturedItemDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchManufacturedItemDefinition given ManufacturedItemDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ManufacturedItemDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchManufacturedItemDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchManufacturedItemDefinition error using patch representation as json")
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

	var parsed r5.ManufacturedItemDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ManufacturedItemDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteManufacturedItemDefinition(deleteResource *r5.ManufacturedItemDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteManufacturedItemDefinition given nil ManufacturedItemDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteManufacturedItemDefinition given ManufacturedItemDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteManufacturedItemDefinitionById(*deleteResource.Id)
}

// delete ManufacturedItemDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteManufacturedItemDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ManufacturedItemDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Measure, return id of created Measure or OperationOutcome error
func (fc *FhirClient) CreateMeasure(createResource *r5.Measure) (*r5.Measure, error) {
	if createResource == nil {
		return nil, errors.New("CreateMeasure called with nil Measure")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Measure")
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

	var parsed r5.Measure
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Measure from fhir server by id, return Measure or OperationOutcome error
func (fc *FhirClient) ReadMeasure(id string) (*r5.Measure, error) {
	if id == "" {
		return nil, errors.New("ReadMeasure given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Measure", id)
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

	var parsed r5.Measure
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Measure if exists on server, else create new Measure with given id, return Measure or OperationOutcome error
func (fc *FhirClient) UpdateMeasure(updateResource *r5.Measure) (*r5.Measure, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateMeasure called with nil Measure")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Measure", *updateResource.Id)
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

	var parsed r5.Measure
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Measure or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMeasure(patchResource *r5.Measure) (*r5.Measure, error) {
	if patchResource == nil {
		return nil, errors.New("PatchMeasure given nil Measure")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchMeasure given Measure without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Measure", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchMeasure error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchMeasure error using patch representation as json")
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

	var parsed r5.Measure
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Measure and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMeasure(deleteResource *r5.Measure) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMeasure given nil Measure")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMeasure given Measure with nil Id (Id required to delete)")
	}
	return fc.DeleteMeasureById(*deleteResource.Id)
}

// delete Measure with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMeasureById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Measure", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create MeasureReport, return id of created MeasureReport or OperationOutcome error
func (fc *FhirClient) CreateMeasureReport(createResource *r5.MeasureReport) (*r5.MeasureReport, error) {
	if createResource == nil {
		return nil, errors.New("CreateMeasureReport called with nil MeasureReport")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MeasureReport")
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

	var parsed r5.MeasureReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read MeasureReport from fhir server by id, return MeasureReport or OperationOutcome error
func (fc *FhirClient) ReadMeasureReport(id string) (*r5.MeasureReport, error) {
	if id == "" {
		return nil, errors.New("ReadMeasureReport given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MeasureReport", id)
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

	var parsed r5.MeasureReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace MeasureReport if exists on server, else create new MeasureReport with given id, return MeasureReport or OperationOutcome error
func (fc *FhirClient) UpdateMeasureReport(updateResource *r5.MeasureReport) (*r5.MeasureReport, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateMeasureReport called with nil MeasureReport")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MeasureReport", *updateResource.Id)
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

	var parsed r5.MeasureReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return MeasureReport or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMeasureReport(patchResource *r5.MeasureReport) (*r5.MeasureReport, error) {
	if patchResource == nil {
		return nil, errors.New("PatchMeasureReport given nil MeasureReport")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchMeasureReport given MeasureReport without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MeasureReport", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchMeasureReport error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchMeasureReport error using patch representation as json")
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

	var parsed r5.MeasureReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete MeasureReport and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMeasureReport(deleteResource *r5.MeasureReport) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMeasureReport given nil MeasureReport")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMeasureReport given MeasureReport with nil Id (Id required to delete)")
	}
	return fc.DeleteMeasureReportById(*deleteResource.Id)
}

// delete MeasureReport with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMeasureReportById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MeasureReport", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Medication, return id of created Medication or OperationOutcome error
func (fc *FhirClient) CreateMedication(createResource *r5.Medication) (*r5.Medication, error) {
	if createResource == nil {
		return nil, errors.New("CreateMedication called with nil Medication")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Medication")
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

	var parsed r5.Medication
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Medication from fhir server by id, return Medication or OperationOutcome error
func (fc *FhirClient) ReadMedication(id string) (*r5.Medication, error) {
	if id == "" {
		return nil, errors.New("ReadMedication given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Medication", id)
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

	var parsed r5.Medication
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Medication if exists on server, else create new Medication with given id, return Medication or OperationOutcome error
func (fc *FhirClient) UpdateMedication(updateResource *r5.Medication) (*r5.Medication, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateMedication called with nil Medication")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Medication", *updateResource.Id)
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

	var parsed r5.Medication
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Medication or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMedication(patchResource *r5.Medication) (*r5.Medication, error) {
	if patchResource == nil {
		return nil, errors.New("PatchMedication given nil Medication")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchMedication given Medication without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Medication", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchMedication error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchMedication error using patch representation as json")
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

	var parsed r5.Medication
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Medication and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedication(deleteResource *r5.Medication) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMedication given nil Medication")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMedication given Medication with nil Id (Id required to delete)")
	}
	return fc.DeleteMedicationById(*deleteResource.Id)
}

// delete Medication with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Medication", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create MedicationAdministration, return id of created MedicationAdministration or OperationOutcome error
func (fc *FhirClient) CreateMedicationAdministration(createResource *r5.MedicationAdministration) (*r5.MedicationAdministration, error) {
	if createResource == nil {
		return nil, errors.New("CreateMedicationAdministration called with nil MedicationAdministration")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationAdministration")
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

	var parsed r5.MedicationAdministration
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read MedicationAdministration from fhir server by id, return MedicationAdministration or OperationOutcome error
func (fc *FhirClient) ReadMedicationAdministration(id string) (*r5.MedicationAdministration, error) {
	if id == "" {
		return nil, errors.New("ReadMedicationAdministration given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationAdministration", id)
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

	var parsed r5.MedicationAdministration
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace MedicationAdministration if exists on server, else create new MedicationAdministration with given id, return MedicationAdministration or OperationOutcome error
func (fc *FhirClient) UpdateMedicationAdministration(updateResource *r5.MedicationAdministration) (*r5.MedicationAdministration, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateMedicationAdministration called with nil MedicationAdministration")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationAdministration", *updateResource.Id)
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

	var parsed r5.MedicationAdministration
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return MedicationAdministration or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMedicationAdministration(patchResource *r5.MedicationAdministration) (*r5.MedicationAdministration, error) {
	if patchResource == nil {
		return nil, errors.New("PatchMedicationAdministration given nil MedicationAdministration")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchMedicationAdministration given MedicationAdministration without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationAdministration", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchMedicationAdministration error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchMedicationAdministration error using patch representation as json")
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

	var parsed r5.MedicationAdministration
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete MedicationAdministration and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationAdministration(deleteResource *r5.MedicationAdministration) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMedicationAdministration given nil MedicationAdministration")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMedicationAdministration given MedicationAdministration with nil Id (Id required to delete)")
	}
	return fc.DeleteMedicationAdministrationById(*deleteResource.Id)
}

// delete MedicationAdministration with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationAdministrationById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationAdministration", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create MedicationDispense, return id of created MedicationDispense or OperationOutcome error
func (fc *FhirClient) CreateMedicationDispense(createResource *r5.MedicationDispense) (*r5.MedicationDispense, error) {
	if createResource == nil {
		return nil, errors.New("CreateMedicationDispense called with nil MedicationDispense")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationDispense")
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

	var parsed r5.MedicationDispense
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read MedicationDispense from fhir server by id, return MedicationDispense or OperationOutcome error
func (fc *FhirClient) ReadMedicationDispense(id string) (*r5.MedicationDispense, error) {
	if id == "" {
		return nil, errors.New("ReadMedicationDispense given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationDispense", id)
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

	var parsed r5.MedicationDispense
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace MedicationDispense if exists on server, else create new MedicationDispense with given id, return MedicationDispense or OperationOutcome error
func (fc *FhirClient) UpdateMedicationDispense(updateResource *r5.MedicationDispense) (*r5.MedicationDispense, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateMedicationDispense called with nil MedicationDispense")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationDispense", *updateResource.Id)
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

	var parsed r5.MedicationDispense
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return MedicationDispense or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMedicationDispense(patchResource *r5.MedicationDispense) (*r5.MedicationDispense, error) {
	if patchResource == nil {
		return nil, errors.New("PatchMedicationDispense given nil MedicationDispense")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchMedicationDispense given MedicationDispense without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationDispense", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchMedicationDispense error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchMedicationDispense error using patch representation as json")
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

	var parsed r5.MedicationDispense
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete MedicationDispense and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationDispense(deleteResource *r5.MedicationDispense) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMedicationDispense given nil MedicationDispense")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMedicationDispense given MedicationDispense with nil Id (Id required to delete)")
	}
	return fc.DeleteMedicationDispenseById(*deleteResource.Id)
}

// delete MedicationDispense with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationDispenseById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationDispense", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create MedicationKnowledge, return id of created MedicationKnowledge or OperationOutcome error
func (fc *FhirClient) CreateMedicationKnowledge(createResource *r5.MedicationKnowledge) (*r5.MedicationKnowledge, error) {
	if createResource == nil {
		return nil, errors.New("CreateMedicationKnowledge called with nil MedicationKnowledge")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationKnowledge")
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

	var parsed r5.MedicationKnowledge
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read MedicationKnowledge from fhir server by id, return MedicationKnowledge or OperationOutcome error
func (fc *FhirClient) ReadMedicationKnowledge(id string) (*r5.MedicationKnowledge, error) {
	if id == "" {
		return nil, errors.New("ReadMedicationKnowledge given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationKnowledge", id)
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

	var parsed r5.MedicationKnowledge
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace MedicationKnowledge if exists on server, else create new MedicationKnowledge with given id, return MedicationKnowledge or OperationOutcome error
func (fc *FhirClient) UpdateMedicationKnowledge(updateResource *r5.MedicationKnowledge) (*r5.MedicationKnowledge, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateMedicationKnowledge called with nil MedicationKnowledge")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationKnowledge", *updateResource.Id)
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

	var parsed r5.MedicationKnowledge
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return MedicationKnowledge or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMedicationKnowledge(patchResource *r5.MedicationKnowledge) (*r5.MedicationKnowledge, error) {
	if patchResource == nil {
		return nil, errors.New("PatchMedicationKnowledge given nil MedicationKnowledge")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchMedicationKnowledge given MedicationKnowledge without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationKnowledge", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchMedicationKnowledge error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchMedicationKnowledge error using patch representation as json")
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

	var parsed r5.MedicationKnowledge
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete MedicationKnowledge and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationKnowledge(deleteResource *r5.MedicationKnowledge) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMedicationKnowledge given nil MedicationKnowledge")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMedicationKnowledge given MedicationKnowledge with nil Id (Id required to delete)")
	}
	return fc.DeleteMedicationKnowledgeById(*deleteResource.Id)
}

// delete MedicationKnowledge with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationKnowledgeById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationKnowledge", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create MedicationRequest, return id of created MedicationRequest or OperationOutcome error
func (fc *FhirClient) CreateMedicationRequest(createResource *r5.MedicationRequest) (*r5.MedicationRequest, error) {
	if createResource == nil {
		return nil, errors.New("CreateMedicationRequest called with nil MedicationRequest")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationRequest")
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

	var parsed r5.MedicationRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read MedicationRequest from fhir server by id, return MedicationRequest or OperationOutcome error
func (fc *FhirClient) ReadMedicationRequest(id string) (*r5.MedicationRequest, error) {
	if id == "" {
		return nil, errors.New("ReadMedicationRequest given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationRequest", id)
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

	var parsed r5.MedicationRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace MedicationRequest if exists on server, else create new MedicationRequest with given id, return MedicationRequest or OperationOutcome error
func (fc *FhirClient) UpdateMedicationRequest(updateResource *r5.MedicationRequest) (*r5.MedicationRequest, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateMedicationRequest called with nil MedicationRequest")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationRequest", *updateResource.Id)
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

	var parsed r5.MedicationRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return MedicationRequest or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMedicationRequest(patchResource *r5.MedicationRequest) (*r5.MedicationRequest, error) {
	if patchResource == nil {
		return nil, errors.New("PatchMedicationRequest given nil MedicationRequest")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchMedicationRequest given MedicationRequest without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationRequest", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchMedicationRequest error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchMedicationRequest error using patch representation as json")
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

	var parsed r5.MedicationRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete MedicationRequest and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationRequest(deleteResource *r5.MedicationRequest) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMedicationRequest given nil MedicationRequest")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMedicationRequest given MedicationRequest with nil Id (Id required to delete)")
	}
	return fc.DeleteMedicationRequestById(*deleteResource.Id)
}

// delete MedicationRequest with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationRequestById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationRequest", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create MedicationStatement, return id of created MedicationStatement or OperationOutcome error
func (fc *FhirClient) CreateMedicationStatement(createResource *r5.MedicationStatement) (*r5.MedicationStatement, error) {
	if createResource == nil {
		return nil, errors.New("CreateMedicationStatement called with nil MedicationStatement")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationStatement")
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

	var parsed r5.MedicationStatement
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read MedicationStatement from fhir server by id, return MedicationStatement or OperationOutcome error
func (fc *FhirClient) ReadMedicationStatement(id string) (*r5.MedicationStatement, error) {
	if id == "" {
		return nil, errors.New("ReadMedicationStatement given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationStatement", id)
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

	var parsed r5.MedicationStatement
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace MedicationStatement if exists on server, else create new MedicationStatement with given id, return MedicationStatement or OperationOutcome error
func (fc *FhirClient) UpdateMedicationStatement(updateResource *r5.MedicationStatement) (*r5.MedicationStatement, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateMedicationStatement called with nil MedicationStatement")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationStatement", *updateResource.Id)
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

	var parsed r5.MedicationStatement
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return MedicationStatement or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMedicationStatement(patchResource *r5.MedicationStatement) (*r5.MedicationStatement, error) {
	if patchResource == nil {
		return nil, errors.New("PatchMedicationStatement given nil MedicationStatement")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchMedicationStatement given MedicationStatement without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationStatement", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchMedicationStatement error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchMedicationStatement error using patch representation as json")
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

	var parsed r5.MedicationStatement
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete MedicationStatement and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationStatement(deleteResource *r5.MedicationStatement) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMedicationStatement given nil MedicationStatement")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMedicationStatement given MedicationStatement with nil Id (Id required to delete)")
	}
	return fc.DeleteMedicationStatementById(*deleteResource.Id)
}

// delete MedicationStatement with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationStatementById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationStatement", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create MedicinalProductDefinition, return id of created MedicinalProductDefinition or OperationOutcome error
func (fc *FhirClient) CreateMedicinalProductDefinition(createResource *r5.MedicinalProductDefinition) (*r5.MedicinalProductDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateMedicinalProductDefinition called with nil MedicinalProductDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicinalProductDefinition")
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

	var parsed r5.MedicinalProductDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read MedicinalProductDefinition from fhir server by id, return MedicinalProductDefinition or OperationOutcome error
func (fc *FhirClient) ReadMedicinalProductDefinition(id string) (*r5.MedicinalProductDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadMedicinalProductDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicinalProductDefinition", id)
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

	var parsed r5.MedicinalProductDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace MedicinalProductDefinition if exists on server, else create new MedicinalProductDefinition with given id, return MedicinalProductDefinition or OperationOutcome error
func (fc *FhirClient) UpdateMedicinalProductDefinition(updateResource *r5.MedicinalProductDefinition) (*r5.MedicinalProductDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateMedicinalProductDefinition called with nil MedicinalProductDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicinalProductDefinition", *updateResource.Id)
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

	var parsed r5.MedicinalProductDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return MedicinalProductDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMedicinalProductDefinition(patchResource *r5.MedicinalProductDefinition) (*r5.MedicinalProductDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchMedicinalProductDefinition given nil MedicinalProductDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchMedicinalProductDefinition given MedicinalProductDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicinalProductDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchMedicinalProductDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchMedicinalProductDefinition error using patch representation as json")
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

	var parsed r5.MedicinalProductDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete MedicinalProductDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicinalProductDefinition(deleteResource *r5.MedicinalProductDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMedicinalProductDefinition given nil MedicinalProductDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMedicinalProductDefinition given MedicinalProductDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteMedicinalProductDefinitionById(*deleteResource.Id)
}

// delete MedicinalProductDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicinalProductDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicinalProductDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create MessageDefinition, return id of created MessageDefinition or OperationOutcome error
func (fc *FhirClient) CreateMessageDefinition(createResource *r5.MessageDefinition) (*r5.MessageDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateMessageDefinition called with nil MessageDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MessageDefinition")
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

	var parsed r5.MessageDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read MessageDefinition from fhir server by id, return MessageDefinition or OperationOutcome error
func (fc *FhirClient) ReadMessageDefinition(id string) (*r5.MessageDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadMessageDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MessageDefinition", id)
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

	var parsed r5.MessageDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace MessageDefinition if exists on server, else create new MessageDefinition with given id, return MessageDefinition or OperationOutcome error
func (fc *FhirClient) UpdateMessageDefinition(updateResource *r5.MessageDefinition) (*r5.MessageDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateMessageDefinition called with nil MessageDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MessageDefinition", *updateResource.Id)
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

	var parsed r5.MessageDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return MessageDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMessageDefinition(patchResource *r5.MessageDefinition) (*r5.MessageDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchMessageDefinition given nil MessageDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchMessageDefinition given MessageDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MessageDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchMessageDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchMessageDefinition error using patch representation as json")
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

	var parsed r5.MessageDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete MessageDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMessageDefinition(deleteResource *r5.MessageDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMessageDefinition given nil MessageDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMessageDefinition given MessageDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteMessageDefinitionById(*deleteResource.Id)
}

// delete MessageDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMessageDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MessageDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create MessageHeader, return id of created MessageHeader or OperationOutcome error
func (fc *FhirClient) CreateMessageHeader(createResource *r5.MessageHeader) (*r5.MessageHeader, error) {
	if createResource == nil {
		return nil, errors.New("CreateMessageHeader called with nil MessageHeader")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MessageHeader")
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

	var parsed r5.MessageHeader
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read MessageHeader from fhir server by id, return MessageHeader or OperationOutcome error
func (fc *FhirClient) ReadMessageHeader(id string) (*r5.MessageHeader, error) {
	if id == "" {
		return nil, errors.New("ReadMessageHeader given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MessageHeader", id)
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

	var parsed r5.MessageHeader
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace MessageHeader if exists on server, else create new MessageHeader with given id, return MessageHeader or OperationOutcome error
func (fc *FhirClient) UpdateMessageHeader(updateResource *r5.MessageHeader) (*r5.MessageHeader, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateMessageHeader called with nil MessageHeader")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MessageHeader", *updateResource.Id)
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

	var parsed r5.MessageHeader
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return MessageHeader or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMessageHeader(patchResource *r5.MessageHeader) (*r5.MessageHeader, error) {
	if patchResource == nil {
		return nil, errors.New("PatchMessageHeader given nil MessageHeader")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchMessageHeader given MessageHeader without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MessageHeader", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchMessageHeader error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchMessageHeader error using patch representation as json")
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

	var parsed r5.MessageHeader
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete MessageHeader and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMessageHeader(deleteResource *r5.MessageHeader) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMessageHeader given nil MessageHeader")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMessageHeader given MessageHeader with nil Id (Id required to delete)")
	}
	return fc.DeleteMessageHeaderById(*deleteResource.Id)
}

// delete MessageHeader with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMessageHeaderById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MessageHeader", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create MolecularSequence, return id of created MolecularSequence or OperationOutcome error
func (fc *FhirClient) CreateMolecularSequence(createResource *r5.MolecularSequence) (*r5.MolecularSequence, error) {
	if createResource == nil {
		return nil, errors.New("CreateMolecularSequence called with nil MolecularSequence")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MolecularSequence")
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

	var parsed r5.MolecularSequence
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read MolecularSequence from fhir server by id, return MolecularSequence or OperationOutcome error
func (fc *FhirClient) ReadMolecularSequence(id string) (*r5.MolecularSequence, error) {
	if id == "" {
		return nil, errors.New("ReadMolecularSequence given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MolecularSequence", id)
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

	var parsed r5.MolecularSequence
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace MolecularSequence if exists on server, else create new MolecularSequence with given id, return MolecularSequence or OperationOutcome error
func (fc *FhirClient) UpdateMolecularSequence(updateResource *r5.MolecularSequence) (*r5.MolecularSequence, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateMolecularSequence called with nil MolecularSequence")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MolecularSequence", *updateResource.Id)
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

	var parsed r5.MolecularSequence
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return MolecularSequence or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMolecularSequence(patchResource *r5.MolecularSequence) (*r5.MolecularSequence, error) {
	if patchResource == nil {
		return nil, errors.New("PatchMolecularSequence given nil MolecularSequence")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchMolecularSequence given MolecularSequence without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MolecularSequence", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchMolecularSequence error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchMolecularSequence error using patch representation as json")
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

	var parsed r5.MolecularSequence
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete MolecularSequence and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMolecularSequence(deleteResource *r5.MolecularSequence) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMolecularSequence given nil MolecularSequence")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMolecularSequence given MolecularSequence with nil Id (Id required to delete)")
	}
	return fc.DeleteMolecularSequenceById(*deleteResource.Id)
}

// delete MolecularSequence with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMolecularSequenceById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MolecularSequence", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create NamingSystem, return id of created NamingSystem or OperationOutcome error
func (fc *FhirClient) CreateNamingSystem(createResource *r5.NamingSystem) (*r5.NamingSystem, error) {
	if createResource == nil {
		return nil, errors.New("CreateNamingSystem called with nil NamingSystem")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NamingSystem")
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

	var parsed r5.NamingSystem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read NamingSystem from fhir server by id, return NamingSystem or OperationOutcome error
func (fc *FhirClient) ReadNamingSystem(id string) (*r5.NamingSystem, error) {
	if id == "" {
		return nil, errors.New("ReadNamingSystem given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NamingSystem", id)
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

	var parsed r5.NamingSystem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace NamingSystem if exists on server, else create new NamingSystem with given id, return NamingSystem or OperationOutcome error
func (fc *FhirClient) UpdateNamingSystem(updateResource *r5.NamingSystem) (*r5.NamingSystem, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateNamingSystem called with nil NamingSystem")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NamingSystem", *updateResource.Id)
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

	var parsed r5.NamingSystem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return NamingSystem or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchNamingSystem(patchResource *r5.NamingSystem) (*r5.NamingSystem, error) {
	if patchResource == nil {
		return nil, errors.New("PatchNamingSystem given nil NamingSystem")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchNamingSystem given NamingSystem without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NamingSystem", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchNamingSystem error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchNamingSystem error using patch representation as json")
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

	var parsed r5.NamingSystem
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete NamingSystem and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteNamingSystem(deleteResource *r5.NamingSystem) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteNamingSystem given nil NamingSystem")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteNamingSystem given NamingSystem with nil Id (Id required to delete)")
	}
	return fc.DeleteNamingSystemById(*deleteResource.Id)
}

// delete NamingSystem with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteNamingSystemById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NamingSystem", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create NutritionIntake, return id of created NutritionIntake or OperationOutcome error
func (fc *FhirClient) CreateNutritionIntake(createResource *r5.NutritionIntake) (*r5.NutritionIntake, error) {
	if createResource == nil {
		return nil, errors.New("CreateNutritionIntake called with nil NutritionIntake")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionIntake")
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

	var parsed r5.NutritionIntake
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read NutritionIntake from fhir server by id, return NutritionIntake or OperationOutcome error
func (fc *FhirClient) ReadNutritionIntake(id string) (*r5.NutritionIntake, error) {
	if id == "" {
		return nil, errors.New("ReadNutritionIntake given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionIntake", id)
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

	var parsed r5.NutritionIntake
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace NutritionIntake if exists on server, else create new NutritionIntake with given id, return NutritionIntake or OperationOutcome error
func (fc *FhirClient) UpdateNutritionIntake(updateResource *r5.NutritionIntake) (*r5.NutritionIntake, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateNutritionIntake called with nil NutritionIntake")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionIntake", *updateResource.Id)
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

	var parsed r5.NutritionIntake
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return NutritionIntake or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchNutritionIntake(patchResource *r5.NutritionIntake) (*r5.NutritionIntake, error) {
	if patchResource == nil {
		return nil, errors.New("PatchNutritionIntake given nil NutritionIntake")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchNutritionIntake given NutritionIntake without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionIntake", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchNutritionIntake error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchNutritionIntake error using patch representation as json")
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

	var parsed r5.NutritionIntake
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete NutritionIntake and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteNutritionIntake(deleteResource *r5.NutritionIntake) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteNutritionIntake given nil NutritionIntake")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteNutritionIntake given NutritionIntake with nil Id (Id required to delete)")
	}
	return fc.DeleteNutritionIntakeById(*deleteResource.Id)
}

// delete NutritionIntake with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteNutritionIntakeById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionIntake", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create NutritionOrder, return id of created NutritionOrder or OperationOutcome error
func (fc *FhirClient) CreateNutritionOrder(createResource *r5.NutritionOrder) (*r5.NutritionOrder, error) {
	if createResource == nil {
		return nil, errors.New("CreateNutritionOrder called with nil NutritionOrder")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionOrder")
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

	var parsed r5.NutritionOrder
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read NutritionOrder from fhir server by id, return NutritionOrder or OperationOutcome error
func (fc *FhirClient) ReadNutritionOrder(id string) (*r5.NutritionOrder, error) {
	if id == "" {
		return nil, errors.New("ReadNutritionOrder given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionOrder", id)
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

	var parsed r5.NutritionOrder
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace NutritionOrder if exists on server, else create new NutritionOrder with given id, return NutritionOrder or OperationOutcome error
func (fc *FhirClient) UpdateNutritionOrder(updateResource *r5.NutritionOrder) (*r5.NutritionOrder, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateNutritionOrder called with nil NutritionOrder")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionOrder", *updateResource.Id)
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

	var parsed r5.NutritionOrder
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return NutritionOrder or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchNutritionOrder(patchResource *r5.NutritionOrder) (*r5.NutritionOrder, error) {
	if patchResource == nil {
		return nil, errors.New("PatchNutritionOrder given nil NutritionOrder")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchNutritionOrder given NutritionOrder without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionOrder", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchNutritionOrder error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchNutritionOrder error using patch representation as json")
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

	var parsed r5.NutritionOrder
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete NutritionOrder and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteNutritionOrder(deleteResource *r5.NutritionOrder) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteNutritionOrder given nil NutritionOrder")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteNutritionOrder given NutritionOrder with nil Id (Id required to delete)")
	}
	return fc.DeleteNutritionOrderById(*deleteResource.Id)
}

// delete NutritionOrder with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteNutritionOrderById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionOrder", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create NutritionProduct, return id of created NutritionProduct or OperationOutcome error
func (fc *FhirClient) CreateNutritionProduct(createResource *r5.NutritionProduct) (*r5.NutritionProduct, error) {
	if createResource == nil {
		return nil, errors.New("CreateNutritionProduct called with nil NutritionProduct")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionProduct")
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

	var parsed r5.NutritionProduct
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read NutritionProduct from fhir server by id, return NutritionProduct or OperationOutcome error
func (fc *FhirClient) ReadNutritionProduct(id string) (*r5.NutritionProduct, error) {
	if id == "" {
		return nil, errors.New("ReadNutritionProduct given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionProduct", id)
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

	var parsed r5.NutritionProduct
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace NutritionProduct if exists on server, else create new NutritionProduct with given id, return NutritionProduct or OperationOutcome error
func (fc *FhirClient) UpdateNutritionProduct(updateResource *r5.NutritionProduct) (*r5.NutritionProduct, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateNutritionProduct called with nil NutritionProduct")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionProduct", *updateResource.Id)
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

	var parsed r5.NutritionProduct
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return NutritionProduct or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchNutritionProduct(patchResource *r5.NutritionProduct) (*r5.NutritionProduct, error) {
	if patchResource == nil {
		return nil, errors.New("PatchNutritionProduct given nil NutritionProduct")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchNutritionProduct given NutritionProduct without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionProduct", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchNutritionProduct error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchNutritionProduct error using patch representation as json")
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

	var parsed r5.NutritionProduct
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete NutritionProduct and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteNutritionProduct(deleteResource *r5.NutritionProduct) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteNutritionProduct given nil NutritionProduct")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteNutritionProduct given NutritionProduct with nil Id (Id required to delete)")
	}
	return fc.DeleteNutritionProductById(*deleteResource.Id)
}

// delete NutritionProduct with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteNutritionProductById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionProduct", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Observation, return id of created Observation or OperationOutcome error
func (fc *FhirClient) CreateObservation(createResource *r5.Observation) (*r5.Observation, error) {
	if createResource == nil {
		return nil, errors.New("CreateObservation called with nil Observation")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Observation")
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

	var parsed r5.Observation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Observation from fhir server by id, return Observation or OperationOutcome error
func (fc *FhirClient) ReadObservation(id string) (*r5.Observation, error) {
	if id == "" {
		return nil, errors.New("ReadObservation given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Observation", id)
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

	var parsed r5.Observation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Observation if exists on server, else create new Observation with given id, return Observation or OperationOutcome error
func (fc *FhirClient) UpdateObservation(updateResource *r5.Observation) (*r5.Observation, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateObservation called with nil Observation")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Observation", *updateResource.Id)
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

	var parsed r5.Observation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Observation or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchObservation(patchResource *r5.Observation) (*r5.Observation, error) {
	if patchResource == nil {
		return nil, errors.New("PatchObservation given nil Observation")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchObservation given Observation without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Observation", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchObservation error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchObservation error using patch representation as json")
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

	var parsed r5.Observation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Observation and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteObservation(deleteResource *r5.Observation) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteObservation given nil Observation")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteObservation given Observation with nil Id (Id required to delete)")
	}
	return fc.DeleteObservationById(*deleteResource.Id)
}

// delete Observation with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteObservationById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Observation", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ObservationDefinition, return id of created ObservationDefinition or OperationOutcome error
func (fc *FhirClient) CreateObservationDefinition(createResource *r5.ObservationDefinition) (*r5.ObservationDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateObservationDefinition called with nil ObservationDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ObservationDefinition")
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

	var parsed r5.ObservationDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ObservationDefinition from fhir server by id, return ObservationDefinition or OperationOutcome error
func (fc *FhirClient) ReadObservationDefinition(id string) (*r5.ObservationDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadObservationDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ObservationDefinition", id)
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

	var parsed r5.ObservationDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ObservationDefinition if exists on server, else create new ObservationDefinition with given id, return ObservationDefinition or OperationOutcome error
func (fc *FhirClient) UpdateObservationDefinition(updateResource *r5.ObservationDefinition) (*r5.ObservationDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateObservationDefinition called with nil ObservationDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ObservationDefinition", *updateResource.Id)
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

	var parsed r5.ObservationDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ObservationDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchObservationDefinition(patchResource *r5.ObservationDefinition) (*r5.ObservationDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchObservationDefinition given nil ObservationDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchObservationDefinition given ObservationDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ObservationDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchObservationDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchObservationDefinition error using patch representation as json")
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

	var parsed r5.ObservationDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ObservationDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteObservationDefinition(deleteResource *r5.ObservationDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteObservationDefinition given nil ObservationDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteObservationDefinition given ObservationDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteObservationDefinitionById(*deleteResource.Id)
}

// delete ObservationDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteObservationDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ObservationDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create OperationDefinition, return id of created OperationDefinition or OperationOutcome error
func (fc *FhirClient) CreateOperationDefinition(createResource *r5.OperationDefinition) (*r5.OperationDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateOperationDefinition called with nil OperationDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OperationDefinition")
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

	var parsed r5.OperationDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read OperationDefinition from fhir server by id, return OperationDefinition or OperationOutcome error
func (fc *FhirClient) ReadOperationDefinition(id string) (*r5.OperationDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadOperationDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OperationDefinition", id)
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

	var parsed r5.OperationDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace OperationDefinition if exists on server, else create new OperationDefinition with given id, return OperationDefinition or OperationOutcome error
func (fc *FhirClient) UpdateOperationDefinition(updateResource *r5.OperationDefinition) (*r5.OperationDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateOperationDefinition called with nil OperationDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OperationDefinition", *updateResource.Id)
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

	var parsed r5.OperationDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return OperationDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchOperationDefinition(patchResource *r5.OperationDefinition) (*r5.OperationDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchOperationDefinition given nil OperationDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchOperationDefinition given OperationDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OperationDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchOperationDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchOperationDefinition error using patch representation as json")
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

	var parsed r5.OperationDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete OperationDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteOperationDefinition(deleteResource *r5.OperationDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteOperationDefinition given nil OperationDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteOperationDefinition given OperationDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteOperationDefinitionById(*deleteResource.Id)
}

// delete OperationDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteOperationDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OperationDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create OperationOutcome, return id of created OperationOutcome or OperationOutcome error
func (fc *FhirClient) CreateOperationOutcome(createResource *r5.OperationOutcome) (*r5.OperationOutcome, error) {
	if createResource == nil {
		return nil, errors.New("CreateOperationOutcome called with nil OperationOutcome")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OperationOutcome")
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read OperationOutcome from fhir server by id, return OperationOutcome or OperationOutcome error
func (fc *FhirClient) ReadOperationOutcome(id string) (*r5.OperationOutcome, error) {
	if id == "" {
		return nil, errors.New("ReadOperationOutcome given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OperationOutcome", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace OperationOutcome if exists on server, else create new OperationOutcome with given id, return OperationOutcome or OperationOutcome error
func (fc *FhirClient) UpdateOperationOutcome(updateResource *r5.OperationOutcome) (*r5.OperationOutcome, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateOperationOutcome called with nil OperationOutcome")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OperationOutcome", *updateResource.Id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return OperationOutcome or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchOperationOutcome(patchResource *r5.OperationOutcome) (*r5.OperationOutcome, error) {
	if patchResource == nil {
		return nil, errors.New("PatchOperationOutcome given nil OperationOutcome")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchOperationOutcome given OperationOutcome without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OperationOutcome", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchOperationOutcome error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchOperationOutcome error using patch representation as json")
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete OperationOutcome and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteOperationOutcome(deleteResource *r5.OperationOutcome) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteOperationOutcome given nil OperationOutcome")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteOperationOutcome given OperationOutcome with nil Id (Id required to delete)")
	}
	return fc.DeleteOperationOutcomeById(*deleteResource.Id)
}

// delete OperationOutcome with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteOperationOutcomeById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OperationOutcome", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Organization, return id of created Organization or OperationOutcome error
func (fc *FhirClient) CreateOrganization(createResource *r5.Organization) (*r5.Organization, error) {
	if createResource == nil {
		return nil, errors.New("CreateOrganization called with nil Organization")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Organization")
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

	var parsed r5.Organization
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Organization from fhir server by id, return Organization or OperationOutcome error
func (fc *FhirClient) ReadOrganization(id string) (*r5.Organization, error) {
	if id == "" {
		return nil, errors.New("ReadOrganization given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Organization", id)
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

	var parsed r5.Organization
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Organization if exists on server, else create new Organization with given id, return Organization or OperationOutcome error
func (fc *FhirClient) UpdateOrganization(updateResource *r5.Organization) (*r5.Organization, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateOrganization called with nil Organization")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Organization", *updateResource.Id)
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

	var parsed r5.Organization
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Organization or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchOrganization(patchResource *r5.Organization) (*r5.Organization, error) {
	if patchResource == nil {
		return nil, errors.New("PatchOrganization given nil Organization")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchOrganization given Organization without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Organization", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchOrganization error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchOrganization error using patch representation as json")
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

	var parsed r5.Organization
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Organization and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteOrganization(deleteResource *r5.Organization) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteOrganization given nil Organization")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteOrganization given Organization with nil Id (Id required to delete)")
	}
	return fc.DeleteOrganizationById(*deleteResource.Id)
}

// delete Organization with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteOrganizationById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Organization", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create OrganizationAffiliation, return id of created OrganizationAffiliation or OperationOutcome error
func (fc *FhirClient) CreateOrganizationAffiliation(createResource *r5.OrganizationAffiliation) (*r5.OrganizationAffiliation, error) {
	if createResource == nil {
		return nil, errors.New("CreateOrganizationAffiliation called with nil OrganizationAffiliation")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OrganizationAffiliation")
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

	var parsed r5.OrganizationAffiliation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read OrganizationAffiliation from fhir server by id, return OrganizationAffiliation or OperationOutcome error
func (fc *FhirClient) ReadOrganizationAffiliation(id string) (*r5.OrganizationAffiliation, error) {
	if id == "" {
		return nil, errors.New("ReadOrganizationAffiliation given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OrganizationAffiliation", id)
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

	var parsed r5.OrganizationAffiliation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace OrganizationAffiliation if exists on server, else create new OrganizationAffiliation with given id, return OrganizationAffiliation or OperationOutcome error
func (fc *FhirClient) UpdateOrganizationAffiliation(updateResource *r5.OrganizationAffiliation) (*r5.OrganizationAffiliation, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateOrganizationAffiliation called with nil OrganizationAffiliation")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OrganizationAffiliation", *updateResource.Id)
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

	var parsed r5.OrganizationAffiliation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return OrganizationAffiliation or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchOrganizationAffiliation(patchResource *r5.OrganizationAffiliation) (*r5.OrganizationAffiliation, error) {
	if patchResource == nil {
		return nil, errors.New("PatchOrganizationAffiliation given nil OrganizationAffiliation")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchOrganizationAffiliation given OrganizationAffiliation without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OrganizationAffiliation", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchOrganizationAffiliation error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchOrganizationAffiliation error using patch representation as json")
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

	var parsed r5.OrganizationAffiliation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete OrganizationAffiliation and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteOrganizationAffiliation(deleteResource *r5.OrganizationAffiliation) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteOrganizationAffiliation given nil OrganizationAffiliation")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteOrganizationAffiliation given OrganizationAffiliation with nil Id (Id required to delete)")
	}
	return fc.DeleteOrganizationAffiliationById(*deleteResource.Id)
}

// delete OrganizationAffiliation with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteOrganizationAffiliationById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OrganizationAffiliation", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create PackagedProductDefinition, return id of created PackagedProductDefinition or OperationOutcome error
func (fc *FhirClient) CreatePackagedProductDefinition(createResource *r5.PackagedProductDefinition) (*r5.PackagedProductDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreatePackagedProductDefinition called with nil PackagedProductDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PackagedProductDefinition")
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

	var parsed r5.PackagedProductDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read PackagedProductDefinition from fhir server by id, return PackagedProductDefinition or OperationOutcome error
func (fc *FhirClient) ReadPackagedProductDefinition(id string) (*r5.PackagedProductDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadPackagedProductDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PackagedProductDefinition", id)
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

	var parsed r5.PackagedProductDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace PackagedProductDefinition if exists on server, else create new PackagedProductDefinition with given id, return PackagedProductDefinition or OperationOutcome error
func (fc *FhirClient) UpdatePackagedProductDefinition(updateResource *r5.PackagedProductDefinition) (*r5.PackagedProductDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdatePackagedProductDefinition called with nil PackagedProductDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PackagedProductDefinition", *updateResource.Id)
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

	var parsed r5.PackagedProductDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return PackagedProductDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPackagedProductDefinition(patchResource *r5.PackagedProductDefinition) (*r5.PackagedProductDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchPackagedProductDefinition given nil PackagedProductDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchPackagedProductDefinition given PackagedProductDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PackagedProductDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchPackagedProductDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchPackagedProductDefinition error using patch representation as json")
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

	var parsed r5.PackagedProductDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete PackagedProductDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePackagedProductDefinition(deleteResource *r5.PackagedProductDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePackagedProductDefinition given nil PackagedProductDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePackagedProductDefinition given PackagedProductDefinition with nil Id (Id required to delete)")
	}
	return fc.DeletePackagedProductDefinitionById(*deleteResource.Id)
}

// delete PackagedProductDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePackagedProductDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PackagedProductDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Patient, return id of created Patient or OperationOutcome error
func (fc *FhirClient) CreatePatient(createResource *r5.Patient) (*r5.Patient, error) {
	if createResource == nil {
		return nil, errors.New("CreatePatient called with nil Patient")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Patient")
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

	var parsed r5.Patient
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Patient from fhir server by id, return Patient or OperationOutcome error
func (fc *FhirClient) ReadPatient(id string) (*r5.Patient, error) {
	if id == "" {
		return nil, errors.New("ReadPatient given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Patient", id)
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

	var parsed r5.Patient
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Patient if exists on server, else create new Patient with given id, return Patient or OperationOutcome error
func (fc *FhirClient) UpdatePatient(updateResource *r5.Patient) (*r5.Patient, error) {
	if updateResource == nil {
		return nil, errors.New("UpdatePatient called with nil Patient")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Patient", *updateResource.Id)
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

	var parsed r5.Patient
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Patient or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPatient(patchResource *r5.Patient) (*r5.Patient, error) {
	if patchResource == nil {
		return nil, errors.New("PatchPatient given nil Patient")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchPatient given Patient without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Patient", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchPatient error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchPatient error using patch representation as json")
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

	var parsed r5.Patient
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Patient and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePatient(deleteResource *r5.Patient) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePatient given nil Patient")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePatient given Patient with nil Id (Id required to delete)")
	}
	return fc.DeletePatientById(*deleteResource.Id)
}

// delete Patient with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePatientById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Patient", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create PaymentNotice, return id of created PaymentNotice or OperationOutcome error
func (fc *FhirClient) CreatePaymentNotice(createResource *r5.PaymentNotice) (*r5.PaymentNotice, error) {
	if createResource == nil {
		return nil, errors.New("CreatePaymentNotice called with nil PaymentNotice")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PaymentNotice")
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

	var parsed r5.PaymentNotice
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read PaymentNotice from fhir server by id, return PaymentNotice or OperationOutcome error
func (fc *FhirClient) ReadPaymentNotice(id string) (*r5.PaymentNotice, error) {
	if id == "" {
		return nil, errors.New("ReadPaymentNotice given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PaymentNotice", id)
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

	var parsed r5.PaymentNotice
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace PaymentNotice if exists on server, else create new PaymentNotice with given id, return PaymentNotice or OperationOutcome error
func (fc *FhirClient) UpdatePaymentNotice(updateResource *r5.PaymentNotice) (*r5.PaymentNotice, error) {
	if updateResource == nil {
		return nil, errors.New("UpdatePaymentNotice called with nil PaymentNotice")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PaymentNotice", *updateResource.Id)
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

	var parsed r5.PaymentNotice
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return PaymentNotice or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPaymentNotice(patchResource *r5.PaymentNotice) (*r5.PaymentNotice, error) {
	if patchResource == nil {
		return nil, errors.New("PatchPaymentNotice given nil PaymentNotice")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchPaymentNotice given PaymentNotice without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PaymentNotice", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchPaymentNotice error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchPaymentNotice error using patch representation as json")
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

	var parsed r5.PaymentNotice
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete PaymentNotice and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePaymentNotice(deleteResource *r5.PaymentNotice) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePaymentNotice given nil PaymentNotice")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePaymentNotice given PaymentNotice with nil Id (Id required to delete)")
	}
	return fc.DeletePaymentNoticeById(*deleteResource.Id)
}

// delete PaymentNotice with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePaymentNoticeById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PaymentNotice", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create PaymentReconciliation, return id of created PaymentReconciliation or OperationOutcome error
func (fc *FhirClient) CreatePaymentReconciliation(createResource *r5.PaymentReconciliation) (*r5.PaymentReconciliation, error) {
	if createResource == nil {
		return nil, errors.New("CreatePaymentReconciliation called with nil PaymentReconciliation")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PaymentReconciliation")
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

	var parsed r5.PaymentReconciliation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read PaymentReconciliation from fhir server by id, return PaymentReconciliation or OperationOutcome error
func (fc *FhirClient) ReadPaymentReconciliation(id string) (*r5.PaymentReconciliation, error) {
	if id == "" {
		return nil, errors.New("ReadPaymentReconciliation given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PaymentReconciliation", id)
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

	var parsed r5.PaymentReconciliation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace PaymentReconciliation if exists on server, else create new PaymentReconciliation with given id, return PaymentReconciliation or OperationOutcome error
func (fc *FhirClient) UpdatePaymentReconciliation(updateResource *r5.PaymentReconciliation) (*r5.PaymentReconciliation, error) {
	if updateResource == nil {
		return nil, errors.New("UpdatePaymentReconciliation called with nil PaymentReconciliation")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PaymentReconciliation", *updateResource.Id)
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

	var parsed r5.PaymentReconciliation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return PaymentReconciliation or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPaymentReconciliation(patchResource *r5.PaymentReconciliation) (*r5.PaymentReconciliation, error) {
	if patchResource == nil {
		return nil, errors.New("PatchPaymentReconciliation given nil PaymentReconciliation")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchPaymentReconciliation given PaymentReconciliation without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PaymentReconciliation", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchPaymentReconciliation error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchPaymentReconciliation error using patch representation as json")
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

	var parsed r5.PaymentReconciliation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete PaymentReconciliation and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePaymentReconciliation(deleteResource *r5.PaymentReconciliation) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePaymentReconciliation given nil PaymentReconciliation")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePaymentReconciliation given PaymentReconciliation with nil Id (Id required to delete)")
	}
	return fc.DeletePaymentReconciliationById(*deleteResource.Id)
}

// delete PaymentReconciliation with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePaymentReconciliationById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PaymentReconciliation", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Permission, return id of created Permission or OperationOutcome error
func (fc *FhirClient) CreatePermission(createResource *r5.Permission) (*r5.Permission, error) {
	if createResource == nil {
		return nil, errors.New("CreatePermission called with nil Permission")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Permission")
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

	var parsed r5.Permission
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Permission from fhir server by id, return Permission or OperationOutcome error
func (fc *FhirClient) ReadPermission(id string) (*r5.Permission, error) {
	if id == "" {
		return nil, errors.New("ReadPermission given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Permission", id)
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

	var parsed r5.Permission
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Permission if exists on server, else create new Permission with given id, return Permission or OperationOutcome error
func (fc *FhirClient) UpdatePermission(updateResource *r5.Permission) (*r5.Permission, error) {
	if updateResource == nil {
		return nil, errors.New("UpdatePermission called with nil Permission")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Permission", *updateResource.Id)
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

	var parsed r5.Permission
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Permission or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPermission(patchResource *r5.Permission) (*r5.Permission, error) {
	if patchResource == nil {
		return nil, errors.New("PatchPermission given nil Permission")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchPermission given Permission without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Permission", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchPermission error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchPermission error using patch representation as json")
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

	var parsed r5.Permission
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Permission and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePermission(deleteResource *r5.Permission) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePermission given nil Permission")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePermission given Permission with nil Id (Id required to delete)")
	}
	return fc.DeletePermissionById(*deleteResource.Id)
}

// delete Permission with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePermissionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Permission", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Person, return id of created Person or OperationOutcome error
func (fc *FhirClient) CreatePerson(createResource *r5.Person) (*r5.Person, error) {
	if createResource == nil {
		return nil, errors.New("CreatePerson called with nil Person")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Person")
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

	var parsed r5.Person
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Person from fhir server by id, return Person or OperationOutcome error
func (fc *FhirClient) ReadPerson(id string) (*r5.Person, error) {
	if id == "" {
		return nil, errors.New("ReadPerson given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Person", id)
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

	var parsed r5.Person
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Person if exists on server, else create new Person with given id, return Person or OperationOutcome error
func (fc *FhirClient) UpdatePerson(updateResource *r5.Person) (*r5.Person, error) {
	if updateResource == nil {
		return nil, errors.New("UpdatePerson called with nil Person")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Person", *updateResource.Id)
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

	var parsed r5.Person
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Person or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPerson(patchResource *r5.Person) (*r5.Person, error) {
	if patchResource == nil {
		return nil, errors.New("PatchPerson given nil Person")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchPerson given Person without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Person", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchPerson error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchPerson error using patch representation as json")
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

	var parsed r5.Person
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Person and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePerson(deleteResource *r5.Person) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePerson given nil Person")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePerson given Person with nil Id (Id required to delete)")
	}
	return fc.DeletePersonById(*deleteResource.Id)
}

// delete Person with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePersonById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Person", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create PlanDefinition, return id of created PlanDefinition or OperationOutcome error
func (fc *FhirClient) CreatePlanDefinition(createResource *r5.PlanDefinition) (*r5.PlanDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreatePlanDefinition called with nil PlanDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PlanDefinition")
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

	var parsed r5.PlanDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read PlanDefinition from fhir server by id, return PlanDefinition or OperationOutcome error
func (fc *FhirClient) ReadPlanDefinition(id string) (*r5.PlanDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadPlanDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PlanDefinition", id)
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

	var parsed r5.PlanDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace PlanDefinition if exists on server, else create new PlanDefinition with given id, return PlanDefinition or OperationOutcome error
func (fc *FhirClient) UpdatePlanDefinition(updateResource *r5.PlanDefinition) (*r5.PlanDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdatePlanDefinition called with nil PlanDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PlanDefinition", *updateResource.Id)
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

	var parsed r5.PlanDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return PlanDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPlanDefinition(patchResource *r5.PlanDefinition) (*r5.PlanDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchPlanDefinition given nil PlanDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchPlanDefinition given PlanDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PlanDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchPlanDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchPlanDefinition error using patch representation as json")
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

	var parsed r5.PlanDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete PlanDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePlanDefinition(deleteResource *r5.PlanDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePlanDefinition given nil PlanDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePlanDefinition given PlanDefinition with nil Id (Id required to delete)")
	}
	return fc.DeletePlanDefinitionById(*deleteResource.Id)
}

// delete PlanDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePlanDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PlanDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Practitioner, return id of created Practitioner or OperationOutcome error
func (fc *FhirClient) CreatePractitioner(createResource *r5.Practitioner) (*r5.Practitioner, error) {
	if createResource == nil {
		return nil, errors.New("CreatePractitioner called with nil Practitioner")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Practitioner")
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

	var parsed r5.Practitioner
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Practitioner from fhir server by id, return Practitioner or OperationOutcome error
func (fc *FhirClient) ReadPractitioner(id string) (*r5.Practitioner, error) {
	if id == "" {
		return nil, errors.New("ReadPractitioner given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Practitioner", id)
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

	var parsed r5.Practitioner
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Practitioner if exists on server, else create new Practitioner with given id, return Practitioner or OperationOutcome error
func (fc *FhirClient) UpdatePractitioner(updateResource *r5.Practitioner) (*r5.Practitioner, error) {
	if updateResource == nil {
		return nil, errors.New("UpdatePractitioner called with nil Practitioner")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Practitioner", *updateResource.Id)
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

	var parsed r5.Practitioner
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Practitioner or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPractitioner(patchResource *r5.Practitioner) (*r5.Practitioner, error) {
	if patchResource == nil {
		return nil, errors.New("PatchPractitioner given nil Practitioner")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchPractitioner given Practitioner without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Practitioner", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchPractitioner error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchPractitioner error using patch representation as json")
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

	var parsed r5.Practitioner
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Practitioner and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePractitioner(deleteResource *r5.Practitioner) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePractitioner given nil Practitioner")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePractitioner given Practitioner with nil Id (Id required to delete)")
	}
	return fc.DeletePractitionerById(*deleteResource.Id)
}

// delete Practitioner with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePractitionerById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Practitioner", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create PractitionerRole, return id of created PractitionerRole or OperationOutcome error
func (fc *FhirClient) CreatePractitionerRole(createResource *r5.PractitionerRole) (*r5.PractitionerRole, error) {
	if createResource == nil {
		return nil, errors.New("CreatePractitionerRole called with nil PractitionerRole")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PractitionerRole")
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

	var parsed r5.PractitionerRole
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read PractitionerRole from fhir server by id, return PractitionerRole or OperationOutcome error
func (fc *FhirClient) ReadPractitionerRole(id string) (*r5.PractitionerRole, error) {
	if id == "" {
		return nil, errors.New("ReadPractitionerRole given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PractitionerRole", id)
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

	var parsed r5.PractitionerRole
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace PractitionerRole if exists on server, else create new PractitionerRole with given id, return PractitionerRole or OperationOutcome error
func (fc *FhirClient) UpdatePractitionerRole(updateResource *r5.PractitionerRole) (*r5.PractitionerRole, error) {
	if updateResource == nil {
		return nil, errors.New("UpdatePractitionerRole called with nil PractitionerRole")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PractitionerRole", *updateResource.Id)
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

	var parsed r5.PractitionerRole
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return PractitionerRole or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPractitionerRole(patchResource *r5.PractitionerRole) (*r5.PractitionerRole, error) {
	if patchResource == nil {
		return nil, errors.New("PatchPractitionerRole given nil PractitionerRole")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchPractitionerRole given PractitionerRole without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PractitionerRole", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchPractitionerRole error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchPractitionerRole error using patch representation as json")
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

	var parsed r5.PractitionerRole
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete PractitionerRole and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePractitionerRole(deleteResource *r5.PractitionerRole) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePractitionerRole given nil PractitionerRole")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePractitionerRole given PractitionerRole with nil Id (Id required to delete)")
	}
	return fc.DeletePractitionerRoleById(*deleteResource.Id)
}

// delete PractitionerRole with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePractitionerRoleById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PractitionerRole", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Procedure, return id of created Procedure or OperationOutcome error
func (fc *FhirClient) CreateProcedure(createResource *r5.Procedure) (*r5.Procedure, error) {
	if createResource == nil {
		return nil, errors.New("CreateProcedure called with nil Procedure")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Procedure")
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

	var parsed r5.Procedure
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Procedure from fhir server by id, return Procedure or OperationOutcome error
func (fc *FhirClient) ReadProcedure(id string) (*r5.Procedure, error) {
	if id == "" {
		return nil, errors.New("ReadProcedure given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Procedure", id)
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

	var parsed r5.Procedure
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Procedure if exists on server, else create new Procedure with given id, return Procedure or OperationOutcome error
func (fc *FhirClient) UpdateProcedure(updateResource *r5.Procedure) (*r5.Procedure, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateProcedure called with nil Procedure")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Procedure", *updateResource.Id)
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

	var parsed r5.Procedure
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Procedure or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchProcedure(patchResource *r5.Procedure) (*r5.Procedure, error) {
	if patchResource == nil {
		return nil, errors.New("PatchProcedure given nil Procedure")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchProcedure given Procedure without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Procedure", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchProcedure error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchProcedure error using patch representation as json")
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

	var parsed r5.Procedure
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Procedure and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteProcedure(deleteResource *r5.Procedure) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteProcedure given nil Procedure")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteProcedure given Procedure with nil Id (Id required to delete)")
	}
	return fc.DeleteProcedureById(*deleteResource.Id)
}

// delete Procedure with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteProcedureById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Procedure", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Provenance, return id of created Provenance or OperationOutcome error
func (fc *FhirClient) CreateProvenance(createResource *r5.Provenance) (*r5.Provenance, error) {
	if createResource == nil {
		return nil, errors.New("CreateProvenance called with nil Provenance")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Provenance")
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

	var parsed r5.Provenance
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Provenance from fhir server by id, return Provenance or OperationOutcome error
func (fc *FhirClient) ReadProvenance(id string) (*r5.Provenance, error) {
	if id == "" {
		return nil, errors.New("ReadProvenance given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Provenance", id)
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

	var parsed r5.Provenance
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Provenance if exists on server, else create new Provenance with given id, return Provenance or OperationOutcome error
func (fc *FhirClient) UpdateProvenance(updateResource *r5.Provenance) (*r5.Provenance, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateProvenance called with nil Provenance")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Provenance", *updateResource.Id)
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

	var parsed r5.Provenance
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Provenance or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchProvenance(patchResource *r5.Provenance) (*r5.Provenance, error) {
	if patchResource == nil {
		return nil, errors.New("PatchProvenance given nil Provenance")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchProvenance given Provenance without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Provenance", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchProvenance error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchProvenance error using patch representation as json")
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

	var parsed r5.Provenance
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Provenance and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteProvenance(deleteResource *r5.Provenance) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteProvenance given nil Provenance")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteProvenance given Provenance with nil Id (Id required to delete)")
	}
	return fc.DeleteProvenanceById(*deleteResource.Id)
}

// delete Provenance with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteProvenanceById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Provenance", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Questionnaire, return id of created Questionnaire or OperationOutcome error
func (fc *FhirClient) CreateQuestionnaire(createResource *r5.Questionnaire) (*r5.Questionnaire, error) {
	if createResource == nil {
		return nil, errors.New("CreateQuestionnaire called with nil Questionnaire")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Questionnaire")
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

	var parsed r5.Questionnaire
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Questionnaire from fhir server by id, return Questionnaire or OperationOutcome error
func (fc *FhirClient) ReadQuestionnaire(id string) (*r5.Questionnaire, error) {
	if id == "" {
		return nil, errors.New("ReadQuestionnaire given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Questionnaire", id)
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

	var parsed r5.Questionnaire
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Questionnaire if exists on server, else create new Questionnaire with given id, return Questionnaire or OperationOutcome error
func (fc *FhirClient) UpdateQuestionnaire(updateResource *r5.Questionnaire) (*r5.Questionnaire, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateQuestionnaire called with nil Questionnaire")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Questionnaire", *updateResource.Id)
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

	var parsed r5.Questionnaire
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Questionnaire or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchQuestionnaire(patchResource *r5.Questionnaire) (*r5.Questionnaire, error) {
	if patchResource == nil {
		return nil, errors.New("PatchQuestionnaire given nil Questionnaire")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchQuestionnaire given Questionnaire without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Questionnaire", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchQuestionnaire error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchQuestionnaire error using patch representation as json")
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

	var parsed r5.Questionnaire
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Questionnaire and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteQuestionnaire(deleteResource *r5.Questionnaire) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteQuestionnaire given nil Questionnaire")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteQuestionnaire given Questionnaire with nil Id (Id required to delete)")
	}
	return fc.DeleteQuestionnaireById(*deleteResource.Id)
}

// delete Questionnaire with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteQuestionnaireById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Questionnaire", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create QuestionnaireResponse, return id of created QuestionnaireResponse or OperationOutcome error
func (fc *FhirClient) CreateQuestionnaireResponse(createResource *r5.QuestionnaireResponse) (*r5.QuestionnaireResponse, error) {
	if createResource == nil {
		return nil, errors.New("CreateQuestionnaireResponse called with nil QuestionnaireResponse")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "QuestionnaireResponse")
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

	var parsed r5.QuestionnaireResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read QuestionnaireResponse from fhir server by id, return QuestionnaireResponse or OperationOutcome error
func (fc *FhirClient) ReadQuestionnaireResponse(id string) (*r5.QuestionnaireResponse, error) {
	if id == "" {
		return nil, errors.New("ReadQuestionnaireResponse given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "QuestionnaireResponse", id)
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

	var parsed r5.QuestionnaireResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace QuestionnaireResponse if exists on server, else create new QuestionnaireResponse with given id, return QuestionnaireResponse or OperationOutcome error
func (fc *FhirClient) UpdateQuestionnaireResponse(updateResource *r5.QuestionnaireResponse) (*r5.QuestionnaireResponse, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateQuestionnaireResponse called with nil QuestionnaireResponse")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "QuestionnaireResponse", *updateResource.Id)
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

	var parsed r5.QuestionnaireResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return QuestionnaireResponse or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchQuestionnaireResponse(patchResource *r5.QuestionnaireResponse) (*r5.QuestionnaireResponse, error) {
	if patchResource == nil {
		return nil, errors.New("PatchQuestionnaireResponse given nil QuestionnaireResponse")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchQuestionnaireResponse given QuestionnaireResponse without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "QuestionnaireResponse", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchQuestionnaireResponse error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchQuestionnaireResponse error using patch representation as json")
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

	var parsed r5.QuestionnaireResponse
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete QuestionnaireResponse and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteQuestionnaireResponse(deleteResource *r5.QuestionnaireResponse) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteQuestionnaireResponse given nil QuestionnaireResponse")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteQuestionnaireResponse given QuestionnaireResponse with nil Id (Id required to delete)")
	}
	return fc.DeleteQuestionnaireResponseById(*deleteResource.Id)
}

// delete QuestionnaireResponse with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteQuestionnaireResponseById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "QuestionnaireResponse", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create RegulatedAuthorization, return id of created RegulatedAuthorization or OperationOutcome error
func (fc *FhirClient) CreateRegulatedAuthorization(createResource *r5.RegulatedAuthorization) (*r5.RegulatedAuthorization, error) {
	if createResource == nil {
		return nil, errors.New("CreateRegulatedAuthorization called with nil RegulatedAuthorization")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RegulatedAuthorization")
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

	var parsed r5.RegulatedAuthorization
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read RegulatedAuthorization from fhir server by id, return RegulatedAuthorization or OperationOutcome error
func (fc *FhirClient) ReadRegulatedAuthorization(id string) (*r5.RegulatedAuthorization, error) {
	if id == "" {
		return nil, errors.New("ReadRegulatedAuthorization given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RegulatedAuthorization", id)
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

	var parsed r5.RegulatedAuthorization
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace RegulatedAuthorization if exists on server, else create new RegulatedAuthorization with given id, return RegulatedAuthorization or OperationOutcome error
func (fc *FhirClient) UpdateRegulatedAuthorization(updateResource *r5.RegulatedAuthorization) (*r5.RegulatedAuthorization, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateRegulatedAuthorization called with nil RegulatedAuthorization")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RegulatedAuthorization", *updateResource.Id)
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

	var parsed r5.RegulatedAuthorization
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return RegulatedAuthorization or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchRegulatedAuthorization(patchResource *r5.RegulatedAuthorization) (*r5.RegulatedAuthorization, error) {
	if patchResource == nil {
		return nil, errors.New("PatchRegulatedAuthorization given nil RegulatedAuthorization")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchRegulatedAuthorization given RegulatedAuthorization without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RegulatedAuthorization", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchRegulatedAuthorization error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchRegulatedAuthorization error using patch representation as json")
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

	var parsed r5.RegulatedAuthorization
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete RegulatedAuthorization and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRegulatedAuthorization(deleteResource *r5.RegulatedAuthorization) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteRegulatedAuthorization given nil RegulatedAuthorization")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteRegulatedAuthorization given RegulatedAuthorization with nil Id (Id required to delete)")
	}
	return fc.DeleteRegulatedAuthorizationById(*deleteResource.Id)
}

// delete RegulatedAuthorization with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRegulatedAuthorizationById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RegulatedAuthorization", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create RelatedPerson, return id of created RelatedPerson or OperationOutcome error
func (fc *FhirClient) CreateRelatedPerson(createResource *r5.RelatedPerson) (*r5.RelatedPerson, error) {
	if createResource == nil {
		return nil, errors.New("CreateRelatedPerson called with nil RelatedPerson")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RelatedPerson")
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

	var parsed r5.RelatedPerson
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read RelatedPerson from fhir server by id, return RelatedPerson or OperationOutcome error
func (fc *FhirClient) ReadRelatedPerson(id string) (*r5.RelatedPerson, error) {
	if id == "" {
		return nil, errors.New("ReadRelatedPerson given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RelatedPerson", id)
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

	var parsed r5.RelatedPerson
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace RelatedPerson if exists on server, else create new RelatedPerson with given id, return RelatedPerson or OperationOutcome error
func (fc *FhirClient) UpdateRelatedPerson(updateResource *r5.RelatedPerson) (*r5.RelatedPerson, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateRelatedPerson called with nil RelatedPerson")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RelatedPerson", *updateResource.Id)
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

	var parsed r5.RelatedPerson
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return RelatedPerson or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchRelatedPerson(patchResource *r5.RelatedPerson) (*r5.RelatedPerson, error) {
	if patchResource == nil {
		return nil, errors.New("PatchRelatedPerson given nil RelatedPerson")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchRelatedPerson given RelatedPerson without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RelatedPerson", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchRelatedPerson error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchRelatedPerson error using patch representation as json")
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

	var parsed r5.RelatedPerson
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete RelatedPerson and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRelatedPerson(deleteResource *r5.RelatedPerson) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteRelatedPerson given nil RelatedPerson")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteRelatedPerson given RelatedPerson with nil Id (Id required to delete)")
	}
	return fc.DeleteRelatedPersonById(*deleteResource.Id)
}

// delete RelatedPerson with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRelatedPersonById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RelatedPerson", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create RequestOrchestration, return id of created RequestOrchestration or OperationOutcome error
func (fc *FhirClient) CreateRequestOrchestration(createResource *r5.RequestOrchestration) (*r5.RequestOrchestration, error) {
	if createResource == nil {
		return nil, errors.New("CreateRequestOrchestration called with nil RequestOrchestration")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RequestOrchestration")
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

	var parsed r5.RequestOrchestration
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read RequestOrchestration from fhir server by id, return RequestOrchestration or OperationOutcome error
func (fc *FhirClient) ReadRequestOrchestration(id string) (*r5.RequestOrchestration, error) {
	if id == "" {
		return nil, errors.New("ReadRequestOrchestration given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RequestOrchestration", id)
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

	var parsed r5.RequestOrchestration
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace RequestOrchestration if exists on server, else create new RequestOrchestration with given id, return RequestOrchestration or OperationOutcome error
func (fc *FhirClient) UpdateRequestOrchestration(updateResource *r5.RequestOrchestration) (*r5.RequestOrchestration, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateRequestOrchestration called with nil RequestOrchestration")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RequestOrchestration", *updateResource.Id)
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

	var parsed r5.RequestOrchestration
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return RequestOrchestration or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchRequestOrchestration(patchResource *r5.RequestOrchestration) (*r5.RequestOrchestration, error) {
	if patchResource == nil {
		return nil, errors.New("PatchRequestOrchestration given nil RequestOrchestration")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchRequestOrchestration given RequestOrchestration without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RequestOrchestration", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchRequestOrchestration error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchRequestOrchestration error using patch representation as json")
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

	var parsed r5.RequestOrchestration
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete RequestOrchestration and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRequestOrchestration(deleteResource *r5.RequestOrchestration) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteRequestOrchestration given nil RequestOrchestration")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteRequestOrchestration given RequestOrchestration with nil Id (Id required to delete)")
	}
	return fc.DeleteRequestOrchestrationById(*deleteResource.Id)
}

// delete RequestOrchestration with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRequestOrchestrationById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RequestOrchestration", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Requirements, return id of created Requirements or OperationOutcome error
func (fc *FhirClient) CreateRequirements(createResource *r5.Requirements) (*r5.Requirements, error) {
	if createResource == nil {
		return nil, errors.New("CreateRequirements called with nil Requirements")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Requirements")
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

	var parsed r5.Requirements
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Requirements from fhir server by id, return Requirements or OperationOutcome error
func (fc *FhirClient) ReadRequirements(id string) (*r5.Requirements, error) {
	if id == "" {
		return nil, errors.New("ReadRequirements given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Requirements", id)
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

	var parsed r5.Requirements
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Requirements if exists on server, else create new Requirements with given id, return Requirements or OperationOutcome error
func (fc *FhirClient) UpdateRequirements(updateResource *r5.Requirements) (*r5.Requirements, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateRequirements called with nil Requirements")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Requirements", *updateResource.Id)
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

	var parsed r5.Requirements
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Requirements or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchRequirements(patchResource *r5.Requirements) (*r5.Requirements, error) {
	if patchResource == nil {
		return nil, errors.New("PatchRequirements given nil Requirements")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchRequirements given Requirements without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Requirements", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchRequirements error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchRequirements error using patch representation as json")
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

	var parsed r5.Requirements
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Requirements and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRequirements(deleteResource *r5.Requirements) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteRequirements given nil Requirements")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteRequirements given Requirements with nil Id (Id required to delete)")
	}
	return fc.DeleteRequirementsById(*deleteResource.Id)
}

// delete Requirements with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRequirementsById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Requirements", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ResearchStudy, return id of created ResearchStudy or OperationOutcome error
func (fc *FhirClient) CreateResearchStudy(createResource *r5.ResearchStudy) (*r5.ResearchStudy, error) {
	if createResource == nil {
		return nil, errors.New("CreateResearchStudy called with nil ResearchStudy")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchStudy")
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

	var parsed r5.ResearchStudy
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ResearchStudy from fhir server by id, return ResearchStudy or OperationOutcome error
func (fc *FhirClient) ReadResearchStudy(id string) (*r5.ResearchStudy, error) {
	if id == "" {
		return nil, errors.New("ReadResearchStudy given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchStudy", id)
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

	var parsed r5.ResearchStudy
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ResearchStudy if exists on server, else create new ResearchStudy with given id, return ResearchStudy or OperationOutcome error
func (fc *FhirClient) UpdateResearchStudy(updateResource *r5.ResearchStudy) (*r5.ResearchStudy, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateResearchStudy called with nil ResearchStudy")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchStudy", *updateResource.Id)
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

	var parsed r5.ResearchStudy
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ResearchStudy or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchResearchStudy(patchResource *r5.ResearchStudy) (*r5.ResearchStudy, error) {
	if patchResource == nil {
		return nil, errors.New("PatchResearchStudy given nil ResearchStudy")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchResearchStudy given ResearchStudy without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchStudy", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchResearchStudy error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchResearchStudy error using patch representation as json")
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

	var parsed r5.ResearchStudy
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ResearchStudy and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteResearchStudy(deleteResource *r5.ResearchStudy) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteResearchStudy given nil ResearchStudy")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteResearchStudy given ResearchStudy with nil Id (Id required to delete)")
	}
	return fc.DeleteResearchStudyById(*deleteResource.Id)
}

// delete ResearchStudy with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteResearchStudyById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchStudy", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ResearchSubject, return id of created ResearchSubject or OperationOutcome error
func (fc *FhirClient) CreateResearchSubject(createResource *r5.ResearchSubject) (*r5.ResearchSubject, error) {
	if createResource == nil {
		return nil, errors.New("CreateResearchSubject called with nil ResearchSubject")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchSubject")
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

	var parsed r5.ResearchSubject
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ResearchSubject from fhir server by id, return ResearchSubject or OperationOutcome error
func (fc *FhirClient) ReadResearchSubject(id string) (*r5.ResearchSubject, error) {
	if id == "" {
		return nil, errors.New("ReadResearchSubject given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchSubject", id)
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

	var parsed r5.ResearchSubject
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ResearchSubject if exists on server, else create new ResearchSubject with given id, return ResearchSubject or OperationOutcome error
func (fc *FhirClient) UpdateResearchSubject(updateResource *r5.ResearchSubject) (*r5.ResearchSubject, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateResearchSubject called with nil ResearchSubject")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchSubject", *updateResource.Id)
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

	var parsed r5.ResearchSubject
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ResearchSubject or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchResearchSubject(patchResource *r5.ResearchSubject) (*r5.ResearchSubject, error) {
	if patchResource == nil {
		return nil, errors.New("PatchResearchSubject given nil ResearchSubject")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchResearchSubject given ResearchSubject without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchSubject", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchResearchSubject error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchResearchSubject error using patch representation as json")
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

	var parsed r5.ResearchSubject
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ResearchSubject and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteResearchSubject(deleteResource *r5.ResearchSubject) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteResearchSubject given nil ResearchSubject")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteResearchSubject given ResearchSubject with nil Id (Id required to delete)")
	}
	return fc.DeleteResearchSubjectById(*deleteResource.Id)
}

// delete ResearchSubject with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteResearchSubjectById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchSubject", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create RiskAssessment, return id of created RiskAssessment or OperationOutcome error
func (fc *FhirClient) CreateRiskAssessment(createResource *r5.RiskAssessment) (*r5.RiskAssessment, error) {
	if createResource == nil {
		return nil, errors.New("CreateRiskAssessment called with nil RiskAssessment")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RiskAssessment")
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

	var parsed r5.RiskAssessment
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read RiskAssessment from fhir server by id, return RiskAssessment or OperationOutcome error
func (fc *FhirClient) ReadRiskAssessment(id string) (*r5.RiskAssessment, error) {
	if id == "" {
		return nil, errors.New("ReadRiskAssessment given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RiskAssessment", id)
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

	var parsed r5.RiskAssessment
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace RiskAssessment if exists on server, else create new RiskAssessment with given id, return RiskAssessment or OperationOutcome error
func (fc *FhirClient) UpdateRiskAssessment(updateResource *r5.RiskAssessment) (*r5.RiskAssessment, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateRiskAssessment called with nil RiskAssessment")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RiskAssessment", *updateResource.Id)
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

	var parsed r5.RiskAssessment
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return RiskAssessment or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchRiskAssessment(patchResource *r5.RiskAssessment) (*r5.RiskAssessment, error) {
	if patchResource == nil {
		return nil, errors.New("PatchRiskAssessment given nil RiskAssessment")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchRiskAssessment given RiskAssessment without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RiskAssessment", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchRiskAssessment error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchRiskAssessment error using patch representation as json")
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

	var parsed r5.RiskAssessment
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete RiskAssessment and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRiskAssessment(deleteResource *r5.RiskAssessment) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteRiskAssessment given nil RiskAssessment")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteRiskAssessment given RiskAssessment with nil Id (Id required to delete)")
	}
	return fc.DeleteRiskAssessmentById(*deleteResource.Id)
}

// delete RiskAssessment with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRiskAssessmentById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RiskAssessment", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Schedule, return id of created Schedule or OperationOutcome error
func (fc *FhirClient) CreateSchedule(createResource *r5.Schedule) (*r5.Schedule, error) {
	if createResource == nil {
		return nil, errors.New("CreateSchedule called with nil Schedule")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Schedule")
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

	var parsed r5.Schedule
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Schedule from fhir server by id, return Schedule or OperationOutcome error
func (fc *FhirClient) ReadSchedule(id string) (*r5.Schedule, error) {
	if id == "" {
		return nil, errors.New("ReadSchedule given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Schedule", id)
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

	var parsed r5.Schedule
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Schedule if exists on server, else create new Schedule with given id, return Schedule or OperationOutcome error
func (fc *FhirClient) UpdateSchedule(updateResource *r5.Schedule) (*r5.Schedule, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSchedule called with nil Schedule")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Schedule", *updateResource.Id)
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

	var parsed r5.Schedule
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Schedule or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSchedule(patchResource *r5.Schedule) (*r5.Schedule, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSchedule given nil Schedule")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSchedule given Schedule without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Schedule", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSchedule error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSchedule error using patch representation as json")
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

	var parsed r5.Schedule
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Schedule and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSchedule(deleteResource *r5.Schedule) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSchedule given nil Schedule")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSchedule given Schedule with nil Id (Id required to delete)")
	}
	return fc.DeleteScheduleById(*deleteResource.Id)
}

// delete Schedule with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteScheduleById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Schedule", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create SearchParameter, return id of created SearchParameter or OperationOutcome error
func (fc *FhirClient) CreateSearchParameter(createResource *r5.SearchParameter) (*r5.SearchParameter, error) {
	if createResource == nil {
		return nil, errors.New("CreateSearchParameter called with nil SearchParameter")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SearchParameter")
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

	var parsed r5.SearchParameter
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read SearchParameter from fhir server by id, return SearchParameter or OperationOutcome error
func (fc *FhirClient) ReadSearchParameter(id string) (*r5.SearchParameter, error) {
	if id == "" {
		return nil, errors.New("ReadSearchParameter given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SearchParameter", id)
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

	var parsed r5.SearchParameter
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace SearchParameter if exists on server, else create new SearchParameter with given id, return SearchParameter or OperationOutcome error
func (fc *FhirClient) UpdateSearchParameter(updateResource *r5.SearchParameter) (*r5.SearchParameter, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSearchParameter called with nil SearchParameter")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SearchParameter", *updateResource.Id)
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

	var parsed r5.SearchParameter
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return SearchParameter or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSearchParameter(patchResource *r5.SearchParameter) (*r5.SearchParameter, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSearchParameter given nil SearchParameter")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSearchParameter given SearchParameter without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SearchParameter", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSearchParameter error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSearchParameter error using patch representation as json")
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

	var parsed r5.SearchParameter
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete SearchParameter and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSearchParameter(deleteResource *r5.SearchParameter) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSearchParameter given nil SearchParameter")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSearchParameter given SearchParameter with nil Id (Id required to delete)")
	}
	return fc.DeleteSearchParameterById(*deleteResource.Id)
}

// delete SearchParameter with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSearchParameterById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SearchParameter", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ServiceRequest, return id of created ServiceRequest or OperationOutcome error
func (fc *FhirClient) CreateServiceRequest(createResource *r5.ServiceRequest) (*r5.ServiceRequest, error) {
	if createResource == nil {
		return nil, errors.New("CreateServiceRequest called with nil ServiceRequest")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ServiceRequest")
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

	var parsed r5.ServiceRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ServiceRequest from fhir server by id, return ServiceRequest or OperationOutcome error
func (fc *FhirClient) ReadServiceRequest(id string) (*r5.ServiceRequest, error) {
	if id == "" {
		return nil, errors.New("ReadServiceRequest given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ServiceRequest", id)
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

	var parsed r5.ServiceRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ServiceRequest if exists on server, else create new ServiceRequest with given id, return ServiceRequest or OperationOutcome error
func (fc *FhirClient) UpdateServiceRequest(updateResource *r5.ServiceRequest) (*r5.ServiceRequest, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateServiceRequest called with nil ServiceRequest")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ServiceRequest", *updateResource.Id)
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

	var parsed r5.ServiceRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ServiceRequest or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchServiceRequest(patchResource *r5.ServiceRequest) (*r5.ServiceRequest, error) {
	if patchResource == nil {
		return nil, errors.New("PatchServiceRequest given nil ServiceRequest")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchServiceRequest given ServiceRequest without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ServiceRequest", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchServiceRequest error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchServiceRequest error using patch representation as json")
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

	var parsed r5.ServiceRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ServiceRequest and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteServiceRequest(deleteResource *r5.ServiceRequest) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteServiceRequest given nil ServiceRequest")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteServiceRequest given ServiceRequest with nil Id (Id required to delete)")
	}
	return fc.DeleteServiceRequestById(*deleteResource.Id)
}

// delete ServiceRequest with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteServiceRequestById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ServiceRequest", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Slot, return id of created Slot or OperationOutcome error
func (fc *FhirClient) CreateSlot(createResource *r5.Slot) (*r5.Slot, error) {
	if createResource == nil {
		return nil, errors.New("CreateSlot called with nil Slot")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Slot")
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

	var parsed r5.Slot
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Slot from fhir server by id, return Slot or OperationOutcome error
func (fc *FhirClient) ReadSlot(id string) (*r5.Slot, error) {
	if id == "" {
		return nil, errors.New("ReadSlot given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Slot", id)
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

	var parsed r5.Slot
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Slot if exists on server, else create new Slot with given id, return Slot or OperationOutcome error
func (fc *FhirClient) UpdateSlot(updateResource *r5.Slot) (*r5.Slot, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSlot called with nil Slot")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Slot", *updateResource.Id)
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

	var parsed r5.Slot
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Slot or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSlot(patchResource *r5.Slot) (*r5.Slot, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSlot given nil Slot")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSlot given Slot without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Slot", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSlot error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSlot error using patch representation as json")
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

	var parsed r5.Slot
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Slot and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSlot(deleteResource *r5.Slot) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSlot given nil Slot")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSlot given Slot with nil Id (Id required to delete)")
	}
	return fc.DeleteSlotById(*deleteResource.Id)
}

// delete Slot with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSlotById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Slot", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Specimen, return id of created Specimen or OperationOutcome error
func (fc *FhirClient) CreateSpecimen(createResource *r5.Specimen) (*r5.Specimen, error) {
	if createResource == nil {
		return nil, errors.New("CreateSpecimen called with nil Specimen")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Specimen")
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

	var parsed r5.Specimen
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Specimen from fhir server by id, return Specimen or OperationOutcome error
func (fc *FhirClient) ReadSpecimen(id string) (*r5.Specimen, error) {
	if id == "" {
		return nil, errors.New("ReadSpecimen given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Specimen", id)
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

	var parsed r5.Specimen
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Specimen if exists on server, else create new Specimen with given id, return Specimen or OperationOutcome error
func (fc *FhirClient) UpdateSpecimen(updateResource *r5.Specimen) (*r5.Specimen, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSpecimen called with nil Specimen")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Specimen", *updateResource.Id)
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

	var parsed r5.Specimen
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Specimen or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSpecimen(patchResource *r5.Specimen) (*r5.Specimen, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSpecimen given nil Specimen")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSpecimen given Specimen without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Specimen", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSpecimen error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSpecimen error using patch representation as json")
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

	var parsed r5.Specimen
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Specimen and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSpecimen(deleteResource *r5.Specimen) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSpecimen given nil Specimen")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSpecimen given Specimen with nil Id (Id required to delete)")
	}
	return fc.DeleteSpecimenById(*deleteResource.Id)
}

// delete Specimen with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSpecimenById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Specimen", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create SpecimenDefinition, return id of created SpecimenDefinition or OperationOutcome error
func (fc *FhirClient) CreateSpecimenDefinition(createResource *r5.SpecimenDefinition) (*r5.SpecimenDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateSpecimenDefinition called with nil SpecimenDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SpecimenDefinition")
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

	var parsed r5.SpecimenDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read SpecimenDefinition from fhir server by id, return SpecimenDefinition or OperationOutcome error
func (fc *FhirClient) ReadSpecimenDefinition(id string) (*r5.SpecimenDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadSpecimenDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SpecimenDefinition", id)
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

	var parsed r5.SpecimenDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace SpecimenDefinition if exists on server, else create new SpecimenDefinition with given id, return SpecimenDefinition or OperationOutcome error
func (fc *FhirClient) UpdateSpecimenDefinition(updateResource *r5.SpecimenDefinition) (*r5.SpecimenDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSpecimenDefinition called with nil SpecimenDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SpecimenDefinition", *updateResource.Id)
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

	var parsed r5.SpecimenDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return SpecimenDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSpecimenDefinition(patchResource *r5.SpecimenDefinition) (*r5.SpecimenDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSpecimenDefinition given nil SpecimenDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSpecimenDefinition given SpecimenDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SpecimenDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSpecimenDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSpecimenDefinition error using patch representation as json")
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

	var parsed r5.SpecimenDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete SpecimenDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSpecimenDefinition(deleteResource *r5.SpecimenDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSpecimenDefinition given nil SpecimenDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSpecimenDefinition given SpecimenDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteSpecimenDefinitionById(*deleteResource.Id)
}

// delete SpecimenDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSpecimenDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SpecimenDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create StructureDefinition, return id of created StructureDefinition or OperationOutcome error
func (fc *FhirClient) CreateStructureDefinition(createResource *r5.StructureDefinition) (*r5.StructureDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateStructureDefinition called with nil StructureDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "StructureDefinition")
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

	var parsed r5.StructureDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read StructureDefinition from fhir server by id, return StructureDefinition or OperationOutcome error
func (fc *FhirClient) ReadStructureDefinition(id string) (*r5.StructureDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadStructureDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "StructureDefinition", id)
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

	var parsed r5.StructureDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace StructureDefinition if exists on server, else create new StructureDefinition with given id, return StructureDefinition or OperationOutcome error
func (fc *FhirClient) UpdateStructureDefinition(updateResource *r5.StructureDefinition) (*r5.StructureDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateStructureDefinition called with nil StructureDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "StructureDefinition", *updateResource.Id)
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

	var parsed r5.StructureDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return StructureDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchStructureDefinition(patchResource *r5.StructureDefinition) (*r5.StructureDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchStructureDefinition given nil StructureDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchStructureDefinition given StructureDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "StructureDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchStructureDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchStructureDefinition error using patch representation as json")
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

	var parsed r5.StructureDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete StructureDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteStructureDefinition(deleteResource *r5.StructureDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteStructureDefinition given nil StructureDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteStructureDefinition given StructureDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteStructureDefinitionById(*deleteResource.Id)
}

// delete StructureDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteStructureDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "StructureDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create StructureMap, return id of created StructureMap or OperationOutcome error
func (fc *FhirClient) CreateStructureMap(createResource *r5.StructureMap) (*r5.StructureMap, error) {
	if createResource == nil {
		return nil, errors.New("CreateStructureMap called with nil StructureMap")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "StructureMap")
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

	var parsed r5.StructureMap
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read StructureMap from fhir server by id, return StructureMap or OperationOutcome error
func (fc *FhirClient) ReadStructureMap(id string) (*r5.StructureMap, error) {
	if id == "" {
		return nil, errors.New("ReadStructureMap given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "StructureMap", id)
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

	var parsed r5.StructureMap
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace StructureMap if exists on server, else create new StructureMap with given id, return StructureMap or OperationOutcome error
func (fc *FhirClient) UpdateStructureMap(updateResource *r5.StructureMap) (*r5.StructureMap, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateStructureMap called with nil StructureMap")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "StructureMap", *updateResource.Id)
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

	var parsed r5.StructureMap
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return StructureMap or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchStructureMap(patchResource *r5.StructureMap) (*r5.StructureMap, error) {
	if patchResource == nil {
		return nil, errors.New("PatchStructureMap given nil StructureMap")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchStructureMap given StructureMap without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "StructureMap", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchStructureMap error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchStructureMap error using patch representation as json")
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

	var parsed r5.StructureMap
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete StructureMap and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteStructureMap(deleteResource *r5.StructureMap) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteStructureMap given nil StructureMap")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteStructureMap given StructureMap with nil Id (Id required to delete)")
	}
	return fc.DeleteStructureMapById(*deleteResource.Id)
}

// delete StructureMap with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteStructureMapById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "StructureMap", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Subscription, return id of created Subscription or OperationOutcome error
func (fc *FhirClient) CreateSubscription(createResource *r5.Subscription) (*r5.Subscription, error) {
	if createResource == nil {
		return nil, errors.New("CreateSubscription called with nil Subscription")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Subscription")
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

	var parsed r5.Subscription
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Subscription from fhir server by id, return Subscription or OperationOutcome error
func (fc *FhirClient) ReadSubscription(id string) (*r5.Subscription, error) {
	if id == "" {
		return nil, errors.New("ReadSubscription given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Subscription", id)
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

	var parsed r5.Subscription
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Subscription if exists on server, else create new Subscription with given id, return Subscription or OperationOutcome error
func (fc *FhirClient) UpdateSubscription(updateResource *r5.Subscription) (*r5.Subscription, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSubscription called with nil Subscription")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Subscription", *updateResource.Id)
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

	var parsed r5.Subscription
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Subscription or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSubscription(patchResource *r5.Subscription) (*r5.Subscription, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSubscription given nil Subscription")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSubscription given Subscription without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Subscription", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSubscription error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSubscription error using patch representation as json")
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

	var parsed r5.Subscription
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Subscription and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubscription(deleteResource *r5.Subscription) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSubscription given nil Subscription")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSubscription given Subscription with nil Id (Id required to delete)")
	}
	return fc.DeleteSubscriptionById(*deleteResource.Id)
}

// delete Subscription with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubscriptionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Subscription", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create SubscriptionStatus, return id of created SubscriptionStatus or OperationOutcome error
func (fc *FhirClient) CreateSubscriptionStatus(createResource *r5.SubscriptionStatus) (*r5.SubscriptionStatus, error) {
	if createResource == nil {
		return nil, errors.New("CreateSubscriptionStatus called with nil SubscriptionStatus")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubscriptionStatus")
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

	var parsed r5.SubscriptionStatus
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read SubscriptionStatus from fhir server by id, return SubscriptionStatus or OperationOutcome error
func (fc *FhirClient) ReadSubscriptionStatus(id string) (*r5.SubscriptionStatus, error) {
	if id == "" {
		return nil, errors.New("ReadSubscriptionStatus given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubscriptionStatus", id)
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

	var parsed r5.SubscriptionStatus
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace SubscriptionStatus if exists on server, else create new SubscriptionStatus with given id, return SubscriptionStatus or OperationOutcome error
func (fc *FhirClient) UpdateSubscriptionStatus(updateResource *r5.SubscriptionStatus) (*r5.SubscriptionStatus, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSubscriptionStatus called with nil SubscriptionStatus")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubscriptionStatus", *updateResource.Id)
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

	var parsed r5.SubscriptionStatus
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return SubscriptionStatus or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSubscriptionStatus(patchResource *r5.SubscriptionStatus) (*r5.SubscriptionStatus, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSubscriptionStatus given nil SubscriptionStatus")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSubscriptionStatus given SubscriptionStatus without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubscriptionStatus", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSubscriptionStatus error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSubscriptionStatus error using patch representation as json")
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

	var parsed r5.SubscriptionStatus
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete SubscriptionStatus and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubscriptionStatus(deleteResource *r5.SubscriptionStatus) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSubscriptionStatus given nil SubscriptionStatus")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSubscriptionStatus given SubscriptionStatus with nil Id (Id required to delete)")
	}
	return fc.DeleteSubscriptionStatusById(*deleteResource.Id)
}

// delete SubscriptionStatus with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubscriptionStatusById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubscriptionStatus", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create SubscriptionTopic, return id of created SubscriptionTopic or OperationOutcome error
func (fc *FhirClient) CreateSubscriptionTopic(createResource *r5.SubscriptionTopic) (*r5.SubscriptionTopic, error) {
	if createResource == nil {
		return nil, errors.New("CreateSubscriptionTopic called with nil SubscriptionTopic")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubscriptionTopic")
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

	var parsed r5.SubscriptionTopic
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read SubscriptionTopic from fhir server by id, return SubscriptionTopic or OperationOutcome error
func (fc *FhirClient) ReadSubscriptionTopic(id string) (*r5.SubscriptionTopic, error) {
	if id == "" {
		return nil, errors.New("ReadSubscriptionTopic given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubscriptionTopic", id)
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

	var parsed r5.SubscriptionTopic
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace SubscriptionTopic if exists on server, else create new SubscriptionTopic with given id, return SubscriptionTopic or OperationOutcome error
func (fc *FhirClient) UpdateSubscriptionTopic(updateResource *r5.SubscriptionTopic) (*r5.SubscriptionTopic, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSubscriptionTopic called with nil SubscriptionTopic")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubscriptionTopic", *updateResource.Id)
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

	var parsed r5.SubscriptionTopic
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return SubscriptionTopic or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSubscriptionTopic(patchResource *r5.SubscriptionTopic) (*r5.SubscriptionTopic, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSubscriptionTopic given nil SubscriptionTopic")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSubscriptionTopic given SubscriptionTopic without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubscriptionTopic", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSubscriptionTopic error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSubscriptionTopic error using patch representation as json")
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

	var parsed r5.SubscriptionTopic
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete SubscriptionTopic and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubscriptionTopic(deleteResource *r5.SubscriptionTopic) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSubscriptionTopic given nil SubscriptionTopic")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSubscriptionTopic given SubscriptionTopic with nil Id (Id required to delete)")
	}
	return fc.DeleteSubscriptionTopicById(*deleteResource.Id)
}

// delete SubscriptionTopic with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubscriptionTopicById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubscriptionTopic", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Substance, return id of created Substance or OperationOutcome error
func (fc *FhirClient) CreateSubstance(createResource *r5.Substance) (*r5.Substance, error) {
	if createResource == nil {
		return nil, errors.New("CreateSubstance called with nil Substance")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Substance")
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

	var parsed r5.Substance
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Substance from fhir server by id, return Substance or OperationOutcome error
func (fc *FhirClient) ReadSubstance(id string) (*r5.Substance, error) {
	if id == "" {
		return nil, errors.New("ReadSubstance given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Substance", id)
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

	var parsed r5.Substance
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Substance if exists on server, else create new Substance with given id, return Substance or OperationOutcome error
func (fc *FhirClient) UpdateSubstance(updateResource *r5.Substance) (*r5.Substance, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSubstance called with nil Substance")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Substance", *updateResource.Id)
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

	var parsed r5.Substance
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Substance or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSubstance(patchResource *r5.Substance) (*r5.Substance, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSubstance given nil Substance")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSubstance given Substance without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Substance", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSubstance error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSubstance error using patch representation as json")
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

	var parsed r5.Substance
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Substance and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstance(deleteResource *r5.Substance) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSubstance given nil Substance")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSubstance given Substance with nil Id (Id required to delete)")
	}
	return fc.DeleteSubstanceById(*deleteResource.Id)
}

// delete Substance with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstanceById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Substance", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create SubstanceDefinition, return id of created SubstanceDefinition or OperationOutcome error
func (fc *FhirClient) CreateSubstanceDefinition(createResource *r5.SubstanceDefinition) (*r5.SubstanceDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateSubstanceDefinition called with nil SubstanceDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceDefinition")
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

	var parsed r5.SubstanceDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read SubstanceDefinition from fhir server by id, return SubstanceDefinition or OperationOutcome error
func (fc *FhirClient) ReadSubstanceDefinition(id string) (*r5.SubstanceDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadSubstanceDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceDefinition", id)
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

	var parsed r5.SubstanceDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace SubstanceDefinition if exists on server, else create new SubstanceDefinition with given id, return SubstanceDefinition or OperationOutcome error
func (fc *FhirClient) UpdateSubstanceDefinition(updateResource *r5.SubstanceDefinition) (*r5.SubstanceDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSubstanceDefinition called with nil SubstanceDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceDefinition", *updateResource.Id)
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

	var parsed r5.SubstanceDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return SubstanceDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSubstanceDefinition(patchResource *r5.SubstanceDefinition) (*r5.SubstanceDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSubstanceDefinition given nil SubstanceDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSubstanceDefinition given SubstanceDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSubstanceDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSubstanceDefinition error using patch representation as json")
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

	var parsed r5.SubstanceDefinition
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete SubstanceDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstanceDefinition(deleteResource *r5.SubstanceDefinition) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSubstanceDefinition given nil SubstanceDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSubstanceDefinition given SubstanceDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteSubstanceDefinitionById(*deleteResource.Id)
}

// delete SubstanceDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstanceDefinitionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceDefinition", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create SubstanceNucleicAcid, return id of created SubstanceNucleicAcid or OperationOutcome error
func (fc *FhirClient) CreateSubstanceNucleicAcid(createResource *r5.SubstanceNucleicAcid) (*r5.SubstanceNucleicAcid, error) {
	if createResource == nil {
		return nil, errors.New("CreateSubstanceNucleicAcid called with nil SubstanceNucleicAcid")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceNucleicAcid")
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

	var parsed r5.SubstanceNucleicAcid
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read SubstanceNucleicAcid from fhir server by id, return SubstanceNucleicAcid or OperationOutcome error
func (fc *FhirClient) ReadSubstanceNucleicAcid(id string) (*r5.SubstanceNucleicAcid, error) {
	if id == "" {
		return nil, errors.New("ReadSubstanceNucleicAcid given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceNucleicAcid", id)
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

	var parsed r5.SubstanceNucleicAcid
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace SubstanceNucleicAcid if exists on server, else create new SubstanceNucleicAcid with given id, return SubstanceNucleicAcid or OperationOutcome error
func (fc *FhirClient) UpdateSubstanceNucleicAcid(updateResource *r5.SubstanceNucleicAcid) (*r5.SubstanceNucleicAcid, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSubstanceNucleicAcid called with nil SubstanceNucleicAcid")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceNucleicAcid", *updateResource.Id)
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

	var parsed r5.SubstanceNucleicAcid
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return SubstanceNucleicAcid or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSubstanceNucleicAcid(patchResource *r5.SubstanceNucleicAcid) (*r5.SubstanceNucleicAcid, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSubstanceNucleicAcid given nil SubstanceNucleicAcid")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSubstanceNucleicAcid given SubstanceNucleicAcid without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceNucleicAcid", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSubstanceNucleicAcid error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSubstanceNucleicAcid error using patch representation as json")
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

	var parsed r5.SubstanceNucleicAcid
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete SubstanceNucleicAcid and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstanceNucleicAcid(deleteResource *r5.SubstanceNucleicAcid) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSubstanceNucleicAcid given nil SubstanceNucleicAcid")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSubstanceNucleicAcid given SubstanceNucleicAcid with nil Id (Id required to delete)")
	}
	return fc.DeleteSubstanceNucleicAcidById(*deleteResource.Id)
}

// delete SubstanceNucleicAcid with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstanceNucleicAcidById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceNucleicAcid", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create SubstancePolymer, return id of created SubstancePolymer or OperationOutcome error
func (fc *FhirClient) CreateSubstancePolymer(createResource *r5.SubstancePolymer) (*r5.SubstancePolymer, error) {
	if createResource == nil {
		return nil, errors.New("CreateSubstancePolymer called with nil SubstancePolymer")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstancePolymer")
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

	var parsed r5.SubstancePolymer
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read SubstancePolymer from fhir server by id, return SubstancePolymer or OperationOutcome error
func (fc *FhirClient) ReadSubstancePolymer(id string) (*r5.SubstancePolymer, error) {
	if id == "" {
		return nil, errors.New("ReadSubstancePolymer given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstancePolymer", id)
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

	var parsed r5.SubstancePolymer
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace SubstancePolymer if exists on server, else create new SubstancePolymer with given id, return SubstancePolymer or OperationOutcome error
func (fc *FhirClient) UpdateSubstancePolymer(updateResource *r5.SubstancePolymer) (*r5.SubstancePolymer, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSubstancePolymer called with nil SubstancePolymer")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstancePolymer", *updateResource.Id)
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

	var parsed r5.SubstancePolymer
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return SubstancePolymer or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSubstancePolymer(patchResource *r5.SubstancePolymer) (*r5.SubstancePolymer, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSubstancePolymer given nil SubstancePolymer")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSubstancePolymer given SubstancePolymer without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstancePolymer", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSubstancePolymer error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSubstancePolymer error using patch representation as json")
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

	var parsed r5.SubstancePolymer
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete SubstancePolymer and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstancePolymer(deleteResource *r5.SubstancePolymer) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSubstancePolymer given nil SubstancePolymer")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSubstancePolymer given SubstancePolymer with nil Id (Id required to delete)")
	}
	return fc.DeleteSubstancePolymerById(*deleteResource.Id)
}

// delete SubstancePolymer with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstancePolymerById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstancePolymer", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create SubstanceProtein, return id of created SubstanceProtein or OperationOutcome error
func (fc *FhirClient) CreateSubstanceProtein(createResource *r5.SubstanceProtein) (*r5.SubstanceProtein, error) {
	if createResource == nil {
		return nil, errors.New("CreateSubstanceProtein called with nil SubstanceProtein")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceProtein")
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

	var parsed r5.SubstanceProtein
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read SubstanceProtein from fhir server by id, return SubstanceProtein or OperationOutcome error
func (fc *FhirClient) ReadSubstanceProtein(id string) (*r5.SubstanceProtein, error) {
	if id == "" {
		return nil, errors.New("ReadSubstanceProtein given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceProtein", id)
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

	var parsed r5.SubstanceProtein
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace SubstanceProtein if exists on server, else create new SubstanceProtein with given id, return SubstanceProtein or OperationOutcome error
func (fc *FhirClient) UpdateSubstanceProtein(updateResource *r5.SubstanceProtein) (*r5.SubstanceProtein, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSubstanceProtein called with nil SubstanceProtein")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceProtein", *updateResource.Id)
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

	var parsed r5.SubstanceProtein
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return SubstanceProtein or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSubstanceProtein(patchResource *r5.SubstanceProtein) (*r5.SubstanceProtein, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSubstanceProtein given nil SubstanceProtein")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSubstanceProtein given SubstanceProtein without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceProtein", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSubstanceProtein error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSubstanceProtein error using patch representation as json")
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

	var parsed r5.SubstanceProtein
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete SubstanceProtein and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstanceProtein(deleteResource *r5.SubstanceProtein) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSubstanceProtein given nil SubstanceProtein")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSubstanceProtein given SubstanceProtein with nil Id (Id required to delete)")
	}
	return fc.DeleteSubstanceProteinById(*deleteResource.Id)
}

// delete SubstanceProtein with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstanceProteinById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceProtein", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create SubstanceReferenceInformation, return id of created SubstanceReferenceInformation or OperationOutcome error
func (fc *FhirClient) CreateSubstanceReferenceInformation(createResource *r5.SubstanceReferenceInformation) (*r5.SubstanceReferenceInformation, error) {
	if createResource == nil {
		return nil, errors.New("CreateSubstanceReferenceInformation called with nil SubstanceReferenceInformation")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceReferenceInformation")
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

	var parsed r5.SubstanceReferenceInformation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read SubstanceReferenceInformation from fhir server by id, return SubstanceReferenceInformation or OperationOutcome error
func (fc *FhirClient) ReadSubstanceReferenceInformation(id string) (*r5.SubstanceReferenceInformation, error) {
	if id == "" {
		return nil, errors.New("ReadSubstanceReferenceInformation given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceReferenceInformation", id)
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

	var parsed r5.SubstanceReferenceInformation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace SubstanceReferenceInformation if exists on server, else create new SubstanceReferenceInformation with given id, return SubstanceReferenceInformation or OperationOutcome error
func (fc *FhirClient) UpdateSubstanceReferenceInformation(updateResource *r5.SubstanceReferenceInformation) (*r5.SubstanceReferenceInformation, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSubstanceReferenceInformation called with nil SubstanceReferenceInformation")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceReferenceInformation", *updateResource.Id)
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

	var parsed r5.SubstanceReferenceInformation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return SubstanceReferenceInformation or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSubstanceReferenceInformation(patchResource *r5.SubstanceReferenceInformation) (*r5.SubstanceReferenceInformation, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSubstanceReferenceInformation given nil SubstanceReferenceInformation")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSubstanceReferenceInformation given SubstanceReferenceInformation without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceReferenceInformation", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSubstanceReferenceInformation error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSubstanceReferenceInformation error using patch representation as json")
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

	var parsed r5.SubstanceReferenceInformation
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete SubstanceReferenceInformation and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstanceReferenceInformation(deleteResource *r5.SubstanceReferenceInformation) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSubstanceReferenceInformation given nil SubstanceReferenceInformation")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSubstanceReferenceInformation given SubstanceReferenceInformation with nil Id (Id required to delete)")
	}
	return fc.DeleteSubstanceReferenceInformationById(*deleteResource.Id)
}

// delete SubstanceReferenceInformation with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstanceReferenceInformationById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceReferenceInformation", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create SubstanceSourceMaterial, return id of created SubstanceSourceMaterial or OperationOutcome error
func (fc *FhirClient) CreateSubstanceSourceMaterial(createResource *r5.SubstanceSourceMaterial) (*r5.SubstanceSourceMaterial, error) {
	if createResource == nil {
		return nil, errors.New("CreateSubstanceSourceMaterial called with nil SubstanceSourceMaterial")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceSourceMaterial")
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

	var parsed r5.SubstanceSourceMaterial
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read SubstanceSourceMaterial from fhir server by id, return SubstanceSourceMaterial or OperationOutcome error
func (fc *FhirClient) ReadSubstanceSourceMaterial(id string) (*r5.SubstanceSourceMaterial, error) {
	if id == "" {
		return nil, errors.New("ReadSubstanceSourceMaterial given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceSourceMaterial", id)
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

	var parsed r5.SubstanceSourceMaterial
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace SubstanceSourceMaterial if exists on server, else create new SubstanceSourceMaterial with given id, return SubstanceSourceMaterial or OperationOutcome error
func (fc *FhirClient) UpdateSubstanceSourceMaterial(updateResource *r5.SubstanceSourceMaterial) (*r5.SubstanceSourceMaterial, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSubstanceSourceMaterial called with nil SubstanceSourceMaterial")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceSourceMaterial", *updateResource.Id)
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

	var parsed r5.SubstanceSourceMaterial
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return SubstanceSourceMaterial or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSubstanceSourceMaterial(patchResource *r5.SubstanceSourceMaterial) (*r5.SubstanceSourceMaterial, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSubstanceSourceMaterial given nil SubstanceSourceMaterial")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSubstanceSourceMaterial given SubstanceSourceMaterial without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceSourceMaterial", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSubstanceSourceMaterial error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSubstanceSourceMaterial error using patch representation as json")
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

	var parsed r5.SubstanceSourceMaterial
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete SubstanceSourceMaterial and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstanceSourceMaterial(deleteResource *r5.SubstanceSourceMaterial) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSubstanceSourceMaterial given nil SubstanceSourceMaterial")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSubstanceSourceMaterial given SubstanceSourceMaterial with nil Id (Id required to delete)")
	}
	return fc.DeleteSubstanceSourceMaterialById(*deleteResource.Id)
}

// delete SubstanceSourceMaterial with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstanceSourceMaterialById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceSourceMaterial", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create SupplyDelivery, return id of created SupplyDelivery or OperationOutcome error
func (fc *FhirClient) CreateSupplyDelivery(createResource *r5.SupplyDelivery) (*r5.SupplyDelivery, error) {
	if createResource == nil {
		return nil, errors.New("CreateSupplyDelivery called with nil SupplyDelivery")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SupplyDelivery")
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

	var parsed r5.SupplyDelivery
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read SupplyDelivery from fhir server by id, return SupplyDelivery or OperationOutcome error
func (fc *FhirClient) ReadSupplyDelivery(id string) (*r5.SupplyDelivery, error) {
	if id == "" {
		return nil, errors.New("ReadSupplyDelivery given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SupplyDelivery", id)
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

	var parsed r5.SupplyDelivery
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace SupplyDelivery if exists on server, else create new SupplyDelivery with given id, return SupplyDelivery or OperationOutcome error
func (fc *FhirClient) UpdateSupplyDelivery(updateResource *r5.SupplyDelivery) (*r5.SupplyDelivery, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSupplyDelivery called with nil SupplyDelivery")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SupplyDelivery", *updateResource.Id)
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

	var parsed r5.SupplyDelivery
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return SupplyDelivery or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSupplyDelivery(patchResource *r5.SupplyDelivery) (*r5.SupplyDelivery, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSupplyDelivery given nil SupplyDelivery")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSupplyDelivery given SupplyDelivery without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SupplyDelivery", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSupplyDelivery error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSupplyDelivery error using patch representation as json")
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

	var parsed r5.SupplyDelivery
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete SupplyDelivery and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSupplyDelivery(deleteResource *r5.SupplyDelivery) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSupplyDelivery given nil SupplyDelivery")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSupplyDelivery given SupplyDelivery with nil Id (Id required to delete)")
	}
	return fc.DeleteSupplyDeliveryById(*deleteResource.Id)
}

// delete SupplyDelivery with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSupplyDeliveryById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SupplyDelivery", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create SupplyRequest, return id of created SupplyRequest or OperationOutcome error
func (fc *FhirClient) CreateSupplyRequest(createResource *r5.SupplyRequest) (*r5.SupplyRequest, error) {
	if createResource == nil {
		return nil, errors.New("CreateSupplyRequest called with nil SupplyRequest")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SupplyRequest")
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

	var parsed r5.SupplyRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read SupplyRequest from fhir server by id, return SupplyRequest or OperationOutcome error
func (fc *FhirClient) ReadSupplyRequest(id string) (*r5.SupplyRequest, error) {
	if id == "" {
		return nil, errors.New("ReadSupplyRequest given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SupplyRequest", id)
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

	var parsed r5.SupplyRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace SupplyRequest if exists on server, else create new SupplyRequest with given id, return SupplyRequest or OperationOutcome error
func (fc *FhirClient) UpdateSupplyRequest(updateResource *r5.SupplyRequest) (*r5.SupplyRequest, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateSupplyRequest called with nil SupplyRequest")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SupplyRequest", *updateResource.Id)
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

	var parsed r5.SupplyRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return SupplyRequest or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSupplyRequest(patchResource *r5.SupplyRequest) (*r5.SupplyRequest, error) {
	if patchResource == nil {
		return nil, errors.New("PatchSupplyRequest given nil SupplyRequest")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchSupplyRequest given SupplyRequest without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SupplyRequest", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchSupplyRequest error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchSupplyRequest error using patch representation as json")
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

	var parsed r5.SupplyRequest
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete SupplyRequest and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSupplyRequest(deleteResource *r5.SupplyRequest) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSupplyRequest given nil SupplyRequest")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSupplyRequest given SupplyRequest with nil Id (Id required to delete)")
	}
	return fc.DeleteSupplyRequestById(*deleteResource.Id)
}

// delete SupplyRequest with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSupplyRequestById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SupplyRequest", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Task, return id of created Task or OperationOutcome error
func (fc *FhirClient) CreateTask(createResource *r5.Task) (*r5.Task, error) {
	if createResource == nil {
		return nil, errors.New("CreateTask called with nil Task")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Task")
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

	var parsed r5.Task
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Task from fhir server by id, return Task or OperationOutcome error
func (fc *FhirClient) ReadTask(id string) (*r5.Task, error) {
	if id == "" {
		return nil, errors.New("ReadTask given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Task", id)
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

	var parsed r5.Task
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Task if exists on server, else create new Task with given id, return Task or OperationOutcome error
func (fc *FhirClient) UpdateTask(updateResource *r5.Task) (*r5.Task, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateTask called with nil Task")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Task", *updateResource.Id)
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

	var parsed r5.Task
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Task or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchTask(patchResource *r5.Task) (*r5.Task, error) {
	if patchResource == nil {
		return nil, errors.New("PatchTask given nil Task")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchTask given Task without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Task", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchTask error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchTask error using patch representation as json")
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

	var parsed r5.Task
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Task and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTask(deleteResource *r5.Task) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteTask given nil Task")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteTask given Task with nil Id (Id required to delete)")
	}
	return fc.DeleteTaskById(*deleteResource.Id)
}

// delete Task with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTaskById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Task", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create TerminologyCapabilities, return id of created TerminologyCapabilities or OperationOutcome error
func (fc *FhirClient) CreateTerminologyCapabilities(createResource *r5.TerminologyCapabilities) (*r5.TerminologyCapabilities, error) {
	if createResource == nil {
		return nil, errors.New("CreateTerminologyCapabilities called with nil TerminologyCapabilities")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TerminologyCapabilities")
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

	var parsed r5.TerminologyCapabilities
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read TerminologyCapabilities from fhir server by id, return TerminologyCapabilities or OperationOutcome error
func (fc *FhirClient) ReadTerminologyCapabilities(id string) (*r5.TerminologyCapabilities, error) {
	if id == "" {
		return nil, errors.New("ReadTerminologyCapabilities given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TerminologyCapabilities", id)
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

	var parsed r5.TerminologyCapabilities
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace TerminologyCapabilities if exists on server, else create new TerminologyCapabilities with given id, return TerminologyCapabilities or OperationOutcome error
func (fc *FhirClient) UpdateTerminologyCapabilities(updateResource *r5.TerminologyCapabilities) (*r5.TerminologyCapabilities, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateTerminologyCapabilities called with nil TerminologyCapabilities")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TerminologyCapabilities", *updateResource.Id)
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

	var parsed r5.TerminologyCapabilities
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return TerminologyCapabilities or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchTerminologyCapabilities(patchResource *r5.TerminologyCapabilities) (*r5.TerminologyCapabilities, error) {
	if patchResource == nil {
		return nil, errors.New("PatchTerminologyCapabilities given nil TerminologyCapabilities")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchTerminologyCapabilities given TerminologyCapabilities without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TerminologyCapabilities", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchTerminologyCapabilities error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchTerminologyCapabilities error using patch representation as json")
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

	var parsed r5.TerminologyCapabilities
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete TerminologyCapabilities and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTerminologyCapabilities(deleteResource *r5.TerminologyCapabilities) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteTerminologyCapabilities given nil TerminologyCapabilities")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteTerminologyCapabilities given TerminologyCapabilities with nil Id (Id required to delete)")
	}
	return fc.DeleteTerminologyCapabilitiesById(*deleteResource.Id)
}

// delete TerminologyCapabilities with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTerminologyCapabilitiesById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TerminologyCapabilities", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create TestPlan, return id of created TestPlan or OperationOutcome error
func (fc *FhirClient) CreateTestPlan(createResource *r5.TestPlan) (*r5.TestPlan, error) {
	if createResource == nil {
		return nil, errors.New("CreateTestPlan called with nil TestPlan")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestPlan")
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

	var parsed r5.TestPlan
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read TestPlan from fhir server by id, return TestPlan or OperationOutcome error
func (fc *FhirClient) ReadTestPlan(id string) (*r5.TestPlan, error) {
	if id == "" {
		return nil, errors.New("ReadTestPlan given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestPlan", id)
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

	var parsed r5.TestPlan
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace TestPlan if exists on server, else create new TestPlan with given id, return TestPlan or OperationOutcome error
func (fc *FhirClient) UpdateTestPlan(updateResource *r5.TestPlan) (*r5.TestPlan, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateTestPlan called with nil TestPlan")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestPlan", *updateResource.Id)
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

	var parsed r5.TestPlan
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return TestPlan or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchTestPlan(patchResource *r5.TestPlan) (*r5.TestPlan, error) {
	if patchResource == nil {
		return nil, errors.New("PatchTestPlan given nil TestPlan")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchTestPlan given TestPlan without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestPlan", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchTestPlan error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchTestPlan error using patch representation as json")
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

	var parsed r5.TestPlan
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete TestPlan and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTestPlan(deleteResource *r5.TestPlan) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteTestPlan given nil TestPlan")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteTestPlan given TestPlan with nil Id (Id required to delete)")
	}
	return fc.DeleteTestPlanById(*deleteResource.Id)
}

// delete TestPlan with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTestPlanById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestPlan", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create TestReport, return id of created TestReport or OperationOutcome error
func (fc *FhirClient) CreateTestReport(createResource *r5.TestReport) (*r5.TestReport, error) {
	if createResource == nil {
		return nil, errors.New("CreateTestReport called with nil TestReport")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestReport")
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

	var parsed r5.TestReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read TestReport from fhir server by id, return TestReport or OperationOutcome error
func (fc *FhirClient) ReadTestReport(id string) (*r5.TestReport, error) {
	if id == "" {
		return nil, errors.New("ReadTestReport given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestReport", id)
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

	var parsed r5.TestReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace TestReport if exists on server, else create new TestReport with given id, return TestReport or OperationOutcome error
func (fc *FhirClient) UpdateTestReport(updateResource *r5.TestReport) (*r5.TestReport, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateTestReport called with nil TestReport")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestReport", *updateResource.Id)
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

	var parsed r5.TestReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return TestReport or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchTestReport(patchResource *r5.TestReport) (*r5.TestReport, error) {
	if patchResource == nil {
		return nil, errors.New("PatchTestReport given nil TestReport")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchTestReport given TestReport without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestReport", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchTestReport error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchTestReport error using patch representation as json")
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

	var parsed r5.TestReport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete TestReport and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTestReport(deleteResource *r5.TestReport) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteTestReport given nil TestReport")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteTestReport given TestReport with nil Id (Id required to delete)")
	}
	return fc.DeleteTestReportById(*deleteResource.Id)
}

// delete TestReport with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTestReportById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestReport", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create TestScript, return id of created TestScript or OperationOutcome error
func (fc *FhirClient) CreateTestScript(createResource *r5.TestScript) (*r5.TestScript, error) {
	if createResource == nil {
		return nil, errors.New("CreateTestScript called with nil TestScript")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestScript")
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

	var parsed r5.TestScript
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read TestScript from fhir server by id, return TestScript or OperationOutcome error
func (fc *FhirClient) ReadTestScript(id string) (*r5.TestScript, error) {
	if id == "" {
		return nil, errors.New("ReadTestScript given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestScript", id)
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

	var parsed r5.TestScript
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace TestScript if exists on server, else create new TestScript with given id, return TestScript or OperationOutcome error
func (fc *FhirClient) UpdateTestScript(updateResource *r5.TestScript) (*r5.TestScript, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateTestScript called with nil TestScript")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestScript", *updateResource.Id)
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

	var parsed r5.TestScript
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return TestScript or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchTestScript(patchResource *r5.TestScript) (*r5.TestScript, error) {
	if patchResource == nil {
		return nil, errors.New("PatchTestScript given nil TestScript")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchTestScript given TestScript without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestScript", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchTestScript error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchTestScript error using patch representation as json")
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

	var parsed r5.TestScript
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete TestScript and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTestScript(deleteResource *r5.TestScript) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteTestScript given nil TestScript")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteTestScript given TestScript with nil Id (Id required to delete)")
	}
	return fc.DeleteTestScriptById(*deleteResource.Id)
}

// delete TestScript with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTestScriptById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestScript", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create Transport, return id of created Transport or OperationOutcome error
func (fc *FhirClient) CreateTransport(createResource *r5.Transport) (*r5.Transport, error) {
	if createResource == nil {
		return nil, errors.New("CreateTransport called with nil Transport")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Transport")
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

	var parsed r5.Transport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read Transport from fhir server by id, return Transport or OperationOutcome error
func (fc *FhirClient) ReadTransport(id string) (*r5.Transport, error) {
	if id == "" {
		return nil, errors.New("ReadTransport given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Transport", id)
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

	var parsed r5.Transport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace Transport if exists on server, else create new Transport with given id, return Transport or OperationOutcome error
func (fc *FhirClient) UpdateTransport(updateResource *r5.Transport) (*r5.Transport, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateTransport called with nil Transport")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Transport", *updateResource.Id)
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

	var parsed r5.Transport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return Transport or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchTransport(patchResource *r5.Transport) (*r5.Transport, error) {
	if patchResource == nil {
		return nil, errors.New("PatchTransport given nil Transport")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchTransport given Transport without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Transport", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchTransport error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchTransport error using patch representation as json")
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

	var parsed r5.Transport
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete Transport and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTransport(deleteResource *r5.Transport) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteTransport given nil Transport")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteTransport given Transport with nil Id (Id required to delete)")
	}
	return fc.DeleteTransportById(*deleteResource.Id)
}

// delete Transport with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTransportById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Transport", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create ValueSet, return id of created ValueSet or OperationOutcome error
func (fc *FhirClient) CreateValueSet(createResource *r5.ValueSet) (*r5.ValueSet, error) {
	if createResource == nil {
		return nil, errors.New("CreateValueSet called with nil ValueSet")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ValueSet")
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

	var parsed r5.ValueSet
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read ValueSet from fhir server by id, return ValueSet or OperationOutcome error
func (fc *FhirClient) ReadValueSet(id string) (*r5.ValueSet, error) {
	if id == "" {
		return nil, errors.New("ReadValueSet given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ValueSet", id)
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

	var parsed r5.ValueSet
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace ValueSet if exists on server, else create new ValueSet with given id, return ValueSet or OperationOutcome error
func (fc *FhirClient) UpdateValueSet(updateResource *r5.ValueSet) (*r5.ValueSet, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateValueSet called with nil ValueSet")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ValueSet", *updateResource.Id)
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

	var parsed r5.ValueSet
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return ValueSet or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchValueSet(patchResource *r5.ValueSet) (*r5.ValueSet, error) {
	if patchResource == nil {
		return nil, errors.New("PatchValueSet given nil ValueSet")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchValueSet given ValueSet without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ValueSet", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchValueSet error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchValueSet error using patch representation as json")
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

	var parsed r5.ValueSet
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete ValueSet and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteValueSet(deleteResource *r5.ValueSet) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteValueSet given nil ValueSet")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteValueSet given ValueSet with nil Id (Id required to delete)")
	}
	return fc.DeleteValueSetById(*deleteResource.Id)
}

// delete ValueSet with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteValueSetById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ValueSet", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create VerificationResult, return id of created VerificationResult or OperationOutcome error
func (fc *FhirClient) CreateVerificationResult(createResource *r5.VerificationResult) (*r5.VerificationResult, error) {
	if createResource == nil {
		return nil, errors.New("CreateVerificationResult called with nil VerificationResult")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "VerificationResult")
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

	var parsed r5.VerificationResult
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read VerificationResult from fhir server by id, return VerificationResult or OperationOutcome error
func (fc *FhirClient) ReadVerificationResult(id string) (*r5.VerificationResult, error) {
	if id == "" {
		return nil, errors.New("ReadVerificationResult given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "VerificationResult", id)
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

	var parsed r5.VerificationResult
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace VerificationResult if exists on server, else create new VerificationResult with given id, return VerificationResult or OperationOutcome error
func (fc *FhirClient) UpdateVerificationResult(updateResource *r5.VerificationResult) (*r5.VerificationResult, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateVerificationResult called with nil VerificationResult")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "VerificationResult", *updateResource.Id)
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

	var parsed r5.VerificationResult
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return VerificationResult or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchVerificationResult(patchResource *r5.VerificationResult) (*r5.VerificationResult, error) {
	if patchResource == nil {
		return nil, errors.New("PatchVerificationResult given nil VerificationResult")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchVerificationResult given VerificationResult without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "VerificationResult", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchVerificationResult error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchVerificationResult error using patch representation as json")
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

	var parsed r5.VerificationResult
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete VerificationResult and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteVerificationResult(deleteResource *r5.VerificationResult) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteVerificationResult given nil VerificationResult")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteVerificationResult given VerificationResult with nil Id (Id required to delete)")
	}
	return fc.DeleteVerificationResultById(*deleteResource.Id)
}

// delete VerificationResult with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteVerificationResultById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "VerificationResult", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// create VisionPrescription, return id of created VisionPrescription or OperationOutcome error
func (fc *FhirClient) CreateVisionPrescription(createResource *r5.VisionPrescription) (*r5.VisionPrescription, error) {
	if createResource == nil {
		return nil, errors.New("CreateVisionPrescription called with nil VisionPrescription")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "VisionPrescription")
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

	var parsed r5.VisionPrescription
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// read VisionPrescription from fhir server by id, return VisionPrescription or OperationOutcome error
func (fc *FhirClient) ReadVisionPrescription(id string) (*r5.VisionPrescription, error) {
	if id == "" {
		return nil, errors.New("ReadVisionPrescription given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "VisionPrescription", id)
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

	var parsed r5.VisionPrescription
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// replace VisionPrescription if exists on server, else create new VisionPrescription with given id, return VisionPrescription or OperationOutcome error
func (fc *FhirClient) UpdateVisionPrescription(updateResource *r5.VisionPrescription) (*r5.VisionPrescription, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateVisionPrescription called with nil VisionPrescription")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "VisionPrescription", *updateResource.Id)
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

	var parsed r5.VisionPrescription
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// add or replace fields in existing resource on server, return VisionPrescription or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchVisionPrescription(patchResource *r5.VisionPrescription) (*r5.VisionPrescription, error) {
	if patchResource == nil {
		return nil, errors.New("PatchVisionPrescription given nil VisionPrescription")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchVisionPrescription given VisionPrescription without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "VisionPrescription", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchVisionPrescription error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchVisionPrescription error using patch representation as json")
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

	var parsed r5.VisionPrescription
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

// delete VisionPrescription and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteVisionPrescription(deleteResource *r5.VisionPrescription) (*r5.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteVisionPrescription given nil VisionPrescription")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteVisionPrescription given VisionPrescription with nil Id (Id required to delete)")
	}
	return fc.DeleteVisionPrescriptionById(*deleteResource.Id)
}

// delete VisionPrescription with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteVisionPrescriptionById(id string) (*r5.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "VisionPrescription", id)
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

	var parsed r5.OperationOutcome
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &parsed, nil
}

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
		var opOutcome r5.OperationOutcome
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
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value any    `json:"value"`
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
}
