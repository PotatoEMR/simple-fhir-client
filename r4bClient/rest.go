package r4bClient

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

	"github.com/PotatoEMR/simple-fhir-client/r4"
	r4b "github.com/PotatoEMR/simple-fhir-client/r4b"
)

// search for resource by its Sp<resource> search params, return bundle of matching resources
//
// SearchGrouped does exact same search, but returns list of each resource type
func (fc *FhirClient) SearchBundled(sp SearchParam) (*r4b.Bundle, error) {
	reqUrl, err := sp.SpEncode(&fc.BaseUrl)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, *reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("searchBundle error creating req, your fault not fhir server's %s", err)
	}

	var parsed r4b.Bundle
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("SearchBundled: %s", err)
		}
	}
	return &parsed, err
}

// search for resource by its Sp<resource> search params, return bundle of matching resources
//
// SearchBundle does exact same search, but returns bundle of all resources
func (fc *FhirClient) SearchGrouped(sp SearchParam) (*ResourceGroup, error) {
	bundle, err := fc.SearchBundled(sp)
	if err != nil {
		return nil, err
	}
	return BundleToGroup(bundle)
}

// Search for everything about a patient, returned as bundle
//
// PatientEverythingGroup does exact same search, but returns list of each resource type
func (fc *FhirClient) PatientEverythingBundled(patId string) (*r4b.Bundle, error) {
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

	var parsed r4b.Bundle
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
}

// create Account, return id of created Account or OperationOutcome error
func (fc *FhirClient) CreateAccount(createResource *r4b.Account) (*r4b.Account, error) {
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

	var parsed r4b.Account
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateAccount: %s", err)
		}
	}
	return &parsed, err
}

// read Account from fhir server by id, return Account or OperationOutcome error
func (fc *FhirClient) ReadAccount(id string) (*r4b.Account, error) {
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

	var parsed r4b.Account
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadAccount: %s", err)
		}
	}
	return &parsed, err
}

// replace Account if exists on server, else create new Account with given id, return Account or OperationOutcome error
func (fc *FhirClient) UpdateAccount(updateResource *r4b.Account) (*r4b.Account, error) {
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

	var parsed r4b.Account
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateAccount: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Account or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchAccount(patchResource *r4b.Account) (*r4b.Account, error) {
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

	var parsed r4b.Account
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchAccount: %s", err)
		}
	}
	return &parsed, err
}

// delete Account and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAccount(deleteResource *r4b.Account) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteAccount given nil Account")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteAccount given Account with nil Id (Id required to delete)")
	}
	return fc.DeleteAccountById(*deleteResource.Id)
}

// delete Account with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAccountById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Account", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteAccountById: %s", err)
		}
	}
	return &parsed, err
}

// create ActivityDefinition, return id of created ActivityDefinition or OperationOutcome error
func (fc *FhirClient) CreateActivityDefinition(createResource *r4b.ActivityDefinition) (*r4b.ActivityDefinition, error) {
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

	var parsed r4b.ActivityDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateActivityDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read ActivityDefinition from fhir server by id, return ActivityDefinition or OperationOutcome error
func (fc *FhirClient) ReadActivityDefinition(id string) (*r4b.ActivityDefinition, error) {
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

	var parsed r4b.ActivityDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadActivityDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace ActivityDefinition if exists on server, else create new ActivityDefinition with given id, return ActivityDefinition or OperationOutcome error
func (fc *FhirClient) UpdateActivityDefinition(updateResource *r4b.ActivityDefinition) (*r4b.ActivityDefinition, error) {
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

	var parsed r4b.ActivityDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateActivityDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ActivityDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchActivityDefinition(patchResource *r4b.ActivityDefinition) (*r4b.ActivityDefinition, error) {
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

	var parsed r4b.ActivityDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchActivityDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete ActivityDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteActivityDefinition(deleteResource *r4b.ActivityDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteActivityDefinition given nil ActivityDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteActivityDefinition given ActivityDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteActivityDefinitionById(*deleteResource.Id)
}

// delete ActivityDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteActivityDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ActivityDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteActivityDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create AdministrableProductDefinition, return id of created AdministrableProductDefinition or OperationOutcome error
func (fc *FhirClient) CreateAdministrableProductDefinition(createResource *r4b.AdministrableProductDefinition) (*r4b.AdministrableProductDefinition, error) {
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

	var parsed r4b.AdministrableProductDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateAdministrableProductDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read AdministrableProductDefinition from fhir server by id, return AdministrableProductDefinition or OperationOutcome error
func (fc *FhirClient) ReadAdministrableProductDefinition(id string) (*r4b.AdministrableProductDefinition, error) {
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

	var parsed r4b.AdministrableProductDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadAdministrableProductDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace AdministrableProductDefinition if exists on server, else create new AdministrableProductDefinition with given id, return AdministrableProductDefinition or OperationOutcome error
func (fc *FhirClient) UpdateAdministrableProductDefinition(updateResource *r4b.AdministrableProductDefinition) (*r4b.AdministrableProductDefinition, error) {
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

	var parsed r4b.AdministrableProductDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateAdministrableProductDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return AdministrableProductDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchAdministrableProductDefinition(patchResource *r4b.AdministrableProductDefinition) (*r4b.AdministrableProductDefinition, error) {
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

	var parsed r4b.AdministrableProductDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchAdministrableProductDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete AdministrableProductDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAdministrableProductDefinition(deleteResource *r4b.AdministrableProductDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteAdministrableProductDefinition given nil AdministrableProductDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteAdministrableProductDefinition given AdministrableProductDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteAdministrableProductDefinitionById(*deleteResource.Id)
}

// delete AdministrableProductDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAdministrableProductDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AdministrableProductDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteAdministrableProductDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create AdverseEvent, return id of created AdverseEvent or OperationOutcome error
func (fc *FhirClient) CreateAdverseEvent(createResource *r4b.AdverseEvent) (*r4b.AdverseEvent, error) {
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

	var parsed r4b.AdverseEvent
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateAdverseEvent: %s", err)
		}
	}
	return &parsed, err
}

// read AdverseEvent from fhir server by id, return AdverseEvent or OperationOutcome error
func (fc *FhirClient) ReadAdverseEvent(id string) (*r4b.AdverseEvent, error) {
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

	var parsed r4b.AdverseEvent
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadAdverseEvent: %s", err)
		}
	}
	return &parsed, err
}

// replace AdverseEvent if exists on server, else create new AdverseEvent with given id, return AdverseEvent or OperationOutcome error
func (fc *FhirClient) UpdateAdverseEvent(updateResource *r4b.AdverseEvent) (*r4b.AdverseEvent, error) {
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

	var parsed r4b.AdverseEvent
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateAdverseEvent: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return AdverseEvent or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchAdverseEvent(patchResource *r4b.AdverseEvent) (*r4b.AdverseEvent, error) {
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

	var parsed r4b.AdverseEvent
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchAdverseEvent: %s", err)
		}
	}
	return &parsed, err
}

// delete AdverseEvent and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAdverseEvent(deleteResource *r4b.AdverseEvent) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteAdverseEvent given nil AdverseEvent")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteAdverseEvent given AdverseEvent with nil Id (Id required to delete)")
	}
	return fc.DeleteAdverseEventById(*deleteResource.Id)
}

// delete AdverseEvent with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAdverseEventById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AdverseEvent", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteAdverseEventById: %s", err)
		}
	}
	return &parsed, err
}

// create AllergyIntolerance, return id of created AllergyIntolerance or OperationOutcome error
func (fc *FhirClient) CreateAllergyIntolerance(createResource *r4b.AllergyIntolerance) (*r4b.AllergyIntolerance, error) {
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

	var parsed r4b.AllergyIntolerance
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateAllergyIntolerance: %s", err)
		}
	}
	return &parsed, err
}

// read AllergyIntolerance from fhir server by id, return AllergyIntolerance or OperationOutcome error
func (fc *FhirClient) ReadAllergyIntolerance(id string) (*r4b.AllergyIntolerance, error) {
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

	var parsed r4b.AllergyIntolerance
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadAllergyIntolerance: %s", err)
		}
	}
	return &parsed, err
}

// replace AllergyIntolerance if exists on server, else create new AllergyIntolerance with given id, return AllergyIntolerance or OperationOutcome error
func (fc *FhirClient) UpdateAllergyIntolerance(updateResource *r4b.AllergyIntolerance) (*r4b.AllergyIntolerance, error) {
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

	var parsed r4b.AllergyIntolerance
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateAllergyIntolerance: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return AllergyIntolerance or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchAllergyIntolerance(patchResource *r4b.AllergyIntolerance) (*r4b.AllergyIntolerance, error) {
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

	var parsed r4b.AllergyIntolerance
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchAllergyIntolerance: %s", err)
		}
	}
	return &parsed, err
}

// delete AllergyIntolerance and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAllergyIntolerance(deleteResource *r4b.AllergyIntolerance) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteAllergyIntolerance given nil AllergyIntolerance")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteAllergyIntolerance given AllergyIntolerance with nil Id (Id required to delete)")
	}
	return fc.DeleteAllergyIntoleranceById(*deleteResource.Id)
}

// delete AllergyIntolerance with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAllergyIntoleranceById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AllergyIntolerance", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteAllergyIntoleranceById: %s", err)
		}
	}
	return &parsed, err
}

// create Appointment, return id of created Appointment or OperationOutcome error
func (fc *FhirClient) CreateAppointment(createResource *r4b.Appointment) (*r4b.Appointment, error) {
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

	var parsed r4b.Appointment
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateAppointment: %s", err)
		}
	}
	return &parsed, err
}

// read Appointment from fhir server by id, return Appointment or OperationOutcome error
func (fc *FhirClient) ReadAppointment(id string) (*r4b.Appointment, error) {
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

	var parsed r4b.Appointment
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadAppointment: %s", err)
		}
	}
	return &parsed, err
}

// replace Appointment if exists on server, else create new Appointment with given id, return Appointment or OperationOutcome error
func (fc *FhirClient) UpdateAppointment(updateResource *r4b.Appointment) (*r4b.Appointment, error) {
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

	var parsed r4b.Appointment
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateAppointment: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Appointment or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchAppointment(patchResource *r4b.Appointment) (*r4b.Appointment, error) {
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

	var parsed r4b.Appointment
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchAppointment: %s", err)
		}
	}
	return &parsed, err
}

// delete Appointment and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAppointment(deleteResource *r4b.Appointment) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteAppointment given nil Appointment")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteAppointment given Appointment with nil Id (Id required to delete)")
	}
	return fc.DeleteAppointmentById(*deleteResource.Id)
}

// delete Appointment with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAppointmentById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Appointment", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteAppointmentById: %s", err)
		}
	}
	return &parsed, err
}

// create AppointmentResponse, return id of created AppointmentResponse or OperationOutcome error
func (fc *FhirClient) CreateAppointmentResponse(createResource *r4b.AppointmentResponse) (*r4b.AppointmentResponse, error) {
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

	var parsed r4b.AppointmentResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateAppointmentResponse: %s", err)
		}
	}
	return &parsed, err
}

// read AppointmentResponse from fhir server by id, return AppointmentResponse or OperationOutcome error
func (fc *FhirClient) ReadAppointmentResponse(id string) (*r4b.AppointmentResponse, error) {
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

	var parsed r4b.AppointmentResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadAppointmentResponse: %s", err)
		}
	}
	return &parsed, err
}

// replace AppointmentResponse if exists on server, else create new AppointmentResponse with given id, return AppointmentResponse or OperationOutcome error
func (fc *FhirClient) UpdateAppointmentResponse(updateResource *r4b.AppointmentResponse) (*r4b.AppointmentResponse, error) {
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

	var parsed r4b.AppointmentResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateAppointmentResponse: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return AppointmentResponse or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchAppointmentResponse(patchResource *r4b.AppointmentResponse) (*r4b.AppointmentResponse, error) {
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

	var parsed r4b.AppointmentResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchAppointmentResponse: %s", err)
		}
	}
	return &parsed, err
}

// delete AppointmentResponse and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAppointmentResponse(deleteResource *r4b.AppointmentResponse) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteAppointmentResponse given nil AppointmentResponse")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteAppointmentResponse given AppointmentResponse with nil Id (Id required to delete)")
	}
	return fc.DeleteAppointmentResponseById(*deleteResource.Id)
}

// delete AppointmentResponse with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAppointmentResponseById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AppointmentResponse", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteAppointmentResponseById: %s", err)
		}
	}
	return &parsed, err
}

// create AuditEvent, return id of created AuditEvent or OperationOutcome error
func (fc *FhirClient) CreateAuditEvent(createResource *r4b.AuditEvent) (*r4b.AuditEvent, error) {
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

	var parsed r4b.AuditEvent
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateAuditEvent: %s", err)
		}
	}
	return &parsed, err
}

// read AuditEvent from fhir server by id, return AuditEvent or OperationOutcome error
func (fc *FhirClient) ReadAuditEvent(id string) (*r4b.AuditEvent, error) {
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

	var parsed r4b.AuditEvent
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadAuditEvent: %s", err)
		}
	}
	return &parsed, err
}

// replace AuditEvent if exists on server, else create new AuditEvent with given id, return AuditEvent or OperationOutcome error
func (fc *FhirClient) UpdateAuditEvent(updateResource *r4b.AuditEvent) (*r4b.AuditEvent, error) {
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

	var parsed r4b.AuditEvent
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateAuditEvent: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return AuditEvent or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchAuditEvent(patchResource *r4b.AuditEvent) (*r4b.AuditEvent, error) {
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

	var parsed r4b.AuditEvent
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchAuditEvent: %s", err)
		}
	}
	return &parsed, err
}

// delete AuditEvent and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAuditEvent(deleteResource *r4b.AuditEvent) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteAuditEvent given nil AuditEvent")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteAuditEvent given AuditEvent with nil Id (Id required to delete)")
	}
	return fc.DeleteAuditEventById(*deleteResource.Id)
}

// delete AuditEvent with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteAuditEventById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "AuditEvent", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteAuditEventById: %s", err)
		}
	}
	return &parsed, err
}

// create Basic, return id of created Basic or OperationOutcome error
func (fc *FhirClient) CreateBasic(createResource *r4b.Basic) (*r4b.Basic, error) {
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

	var parsed r4b.Basic
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateBasic: %s", err)
		}
	}
	return &parsed, err
}

// read Basic from fhir server by id, return Basic or OperationOutcome error
func (fc *FhirClient) ReadBasic(id string) (*r4b.Basic, error) {
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

	var parsed r4b.Basic
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadBasic: %s", err)
		}
	}
	return &parsed, err
}

// replace Basic if exists on server, else create new Basic with given id, return Basic or OperationOutcome error
func (fc *FhirClient) UpdateBasic(updateResource *r4b.Basic) (*r4b.Basic, error) {
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

	var parsed r4b.Basic
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateBasic: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Basic or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchBasic(patchResource *r4b.Basic) (*r4b.Basic, error) {
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

	var parsed r4b.Basic
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchBasic: %s", err)
		}
	}
	return &parsed, err
}

// delete Basic and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBasic(deleteResource *r4b.Basic) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteBasic given nil Basic")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteBasic given Basic with nil Id (Id required to delete)")
	}
	return fc.DeleteBasicById(*deleteResource.Id)
}

// delete Basic with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBasicById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Basic", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteBasicById: %s", err)
		}
	}
	return &parsed, err
}

// create Binary, return id of created Binary or OperationOutcome error
func (fc *FhirClient) CreateBinary(createResource *r4b.Binary) (*r4b.Binary, error) {
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

	var parsed r4b.Binary
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateBinary: %s", err)
		}
	}
	return &parsed, err
}

// read Binary from fhir server by id, return Binary or OperationOutcome error
func (fc *FhirClient) ReadBinary(id string) (*r4b.Binary, error) {
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

	var parsed r4b.Binary
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadBinary: %s", err)
		}
	}
	return &parsed, err
}

// replace Binary if exists on server, else create new Binary with given id, return Binary or OperationOutcome error
func (fc *FhirClient) UpdateBinary(updateResource *r4b.Binary) (*r4b.Binary, error) {
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

	var parsed r4b.Binary
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateBinary: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Binary or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchBinary(patchResource *r4b.Binary) (*r4b.Binary, error) {
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

	var parsed r4b.Binary
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchBinary: %s", err)
		}
	}
	return &parsed, err
}

// delete Binary and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBinary(deleteResource *r4b.Binary) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteBinary given nil Binary")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteBinary given Binary with nil Id (Id required to delete)")
	}
	return fc.DeleteBinaryById(*deleteResource.Id)
}

// delete Binary with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBinaryById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Binary", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteBinaryById: %s", err)
		}
	}
	return &parsed, err
}

// create BiologicallyDerivedProduct, return id of created BiologicallyDerivedProduct or OperationOutcome error
func (fc *FhirClient) CreateBiologicallyDerivedProduct(createResource *r4b.BiologicallyDerivedProduct) (*r4b.BiologicallyDerivedProduct, error) {
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

	var parsed r4b.BiologicallyDerivedProduct
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateBiologicallyDerivedProduct: %s", err)
		}
	}
	return &parsed, err
}

// read BiologicallyDerivedProduct from fhir server by id, return BiologicallyDerivedProduct or OperationOutcome error
func (fc *FhirClient) ReadBiologicallyDerivedProduct(id string) (*r4b.BiologicallyDerivedProduct, error) {
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

	var parsed r4b.BiologicallyDerivedProduct
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadBiologicallyDerivedProduct: %s", err)
		}
	}
	return &parsed, err
}

// replace BiologicallyDerivedProduct if exists on server, else create new BiologicallyDerivedProduct with given id, return BiologicallyDerivedProduct or OperationOutcome error
func (fc *FhirClient) UpdateBiologicallyDerivedProduct(updateResource *r4b.BiologicallyDerivedProduct) (*r4b.BiologicallyDerivedProduct, error) {
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

	var parsed r4b.BiologicallyDerivedProduct
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateBiologicallyDerivedProduct: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return BiologicallyDerivedProduct or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchBiologicallyDerivedProduct(patchResource *r4b.BiologicallyDerivedProduct) (*r4b.BiologicallyDerivedProduct, error) {
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

	var parsed r4b.BiologicallyDerivedProduct
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchBiologicallyDerivedProduct: %s", err)
		}
	}
	return &parsed, err
}

// delete BiologicallyDerivedProduct and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBiologicallyDerivedProduct(deleteResource *r4b.BiologicallyDerivedProduct) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteBiologicallyDerivedProduct given nil BiologicallyDerivedProduct")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteBiologicallyDerivedProduct given BiologicallyDerivedProduct with nil Id (Id required to delete)")
	}
	return fc.DeleteBiologicallyDerivedProductById(*deleteResource.Id)
}

// delete BiologicallyDerivedProduct with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBiologicallyDerivedProductById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BiologicallyDerivedProduct", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteBiologicallyDerivedProductById: %s", err)
		}
	}
	return &parsed, err
}

// create BodyStructure, return id of created BodyStructure or OperationOutcome error
func (fc *FhirClient) CreateBodyStructure(createResource *r4b.BodyStructure) (*r4b.BodyStructure, error) {
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

	var parsed r4b.BodyStructure
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateBodyStructure: %s", err)
		}
	}
	return &parsed, err
}

// read BodyStructure from fhir server by id, return BodyStructure or OperationOutcome error
func (fc *FhirClient) ReadBodyStructure(id string) (*r4b.BodyStructure, error) {
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

	var parsed r4b.BodyStructure
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadBodyStructure: %s", err)
		}
	}
	return &parsed, err
}

// replace BodyStructure if exists on server, else create new BodyStructure with given id, return BodyStructure or OperationOutcome error
func (fc *FhirClient) UpdateBodyStructure(updateResource *r4b.BodyStructure) (*r4b.BodyStructure, error) {
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

	var parsed r4b.BodyStructure
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateBodyStructure: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return BodyStructure or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchBodyStructure(patchResource *r4b.BodyStructure) (*r4b.BodyStructure, error) {
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

	var parsed r4b.BodyStructure
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchBodyStructure: %s", err)
		}
	}
	return &parsed, err
}

// delete BodyStructure and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBodyStructure(deleteResource *r4b.BodyStructure) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteBodyStructure given nil BodyStructure")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteBodyStructure given BodyStructure with nil Id (Id required to delete)")
	}
	return fc.DeleteBodyStructureById(*deleteResource.Id)
}

// delete BodyStructure with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteBodyStructureById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "BodyStructure", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteBodyStructureById: %s", err)
		}
	}
	return &parsed, err
}

// create CapabilityStatement, return id of created CapabilityStatement or OperationOutcome error
func (fc *FhirClient) CreateCapabilityStatement(createResource *r4b.CapabilityStatement) (*r4b.CapabilityStatement, error) {
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

	var parsed r4b.CapabilityStatement
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateCapabilityStatement: %s", err)
		}
	}
	return &parsed, err
}

// read CapabilityStatement from fhir server by id, return CapabilityStatement or OperationOutcome error
func (fc *FhirClient) ReadCapabilityStatement(id string) (*r4b.CapabilityStatement, error) {
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

	var parsed r4b.CapabilityStatement
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadCapabilityStatement: %s", err)
		}
	}
	return &parsed, err
}

// replace CapabilityStatement if exists on server, else create new CapabilityStatement with given id, return CapabilityStatement or OperationOutcome error
func (fc *FhirClient) UpdateCapabilityStatement(updateResource *r4b.CapabilityStatement) (*r4b.CapabilityStatement, error) {
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

	var parsed r4b.CapabilityStatement
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateCapabilityStatement: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return CapabilityStatement or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCapabilityStatement(patchResource *r4b.CapabilityStatement) (*r4b.CapabilityStatement, error) {
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

	var parsed r4b.CapabilityStatement
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchCapabilityStatement: %s", err)
		}
	}
	return &parsed, err
}

// delete CapabilityStatement and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCapabilityStatement(deleteResource *r4b.CapabilityStatement) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCapabilityStatement given nil CapabilityStatement")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCapabilityStatement given CapabilityStatement with nil Id (Id required to delete)")
	}
	return fc.DeleteCapabilityStatementById(*deleteResource.Id)
}

// delete CapabilityStatement with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCapabilityStatementById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CapabilityStatement", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteCapabilityStatementById: %s", err)
		}
	}
	return &parsed, err
}

// create CarePlan, return id of created CarePlan or OperationOutcome error
func (fc *FhirClient) CreateCarePlan(createResource *r4b.CarePlan) (*r4b.CarePlan, error) {
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

	var parsed r4b.CarePlan
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateCarePlan: %s", err)
		}
	}
	return &parsed, err
}

// read CarePlan from fhir server by id, return CarePlan or OperationOutcome error
func (fc *FhirClient) ReadCarePlan(id string) (*r4b.CarePlan, error) {
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

	var parsed r4b.CarePlan
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadCarePlan: %s", err)
		}
	}
	return &parsed, err
}

// replace CarePlan if exists on server, else create new CarePlan with given id, return CarePlan or OperationOutcome error
func (fc *FhirClient) UpdateCarePlan(updateResource *r4b.CarePlan) (*r4b.CarePlan, error) {
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

	var parsed r4b.CarePlan
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateCarePlan: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return CarePlan or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCarePlan(patchResource *r4b.CarePlan) (*r4b.CarePlan, error) {
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

	var parsed r4b.CarePlan
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchCarePlan: %s", err)
		}
	}
	return &parsed, err
}

// delete CarePlan and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCarePlan(deleteResource *r4b.CarePlan) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCarePlan given nil CarePlan")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCarePlan given CarePlan with nil Id (Id required to delete)")
	}
	return fc.DeleteCarePlanById(*deleteResource.Id)
}

// delete CarePlan with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCarePlanById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CarePlan", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteCarePlanById: %s", err)
		}
	}
	return &parsed, err
}

// create CareTeam, return id of created CareTeam or OperationOutcome error
func (fc *FhirClient) CreateCareTeam(createResource *r4b.CareTeam) (*r4b.CareTeam, error) {
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

	var parsed r4b.CareTeam
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateCareTeam: %s", err)
		}
	}
	return &parsed, err
}

// read CareTeam from fhir server by id, return CareTeam or OperationOutcome error
func (fc *FhirClient) ReadCareTeam(id string) (*r4b.CareTeam, error) {
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

	var parsed r4b.CareTeam
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadCareTeam: %s", err)
		}
	}
	return &parsed, err
}

// replace CareTeam if exists on server, else create new CareTeam with given id, return CareTeam or OperationOutcome error
func (fc *FhirClient) UpdateCareTeam(updateResource *r4b.CareTeam) (*r4b.CareTeam, error) {
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

	var parsed r4b.CareTeam
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateCareTeam: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return CareTeam or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCareTeam(patchResource *r4b.CareTeam) (*r4b.CareTeam, error) {
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

	var parsed r4b.CareTeam
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchCareTeam: %s", err)
		}
	}
	return &parsed, err
}

// delete CareTeam and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCareTeam(deleteResource *r4b.CareTeam) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCareTeam given nil CareTeam")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCareTeam given CareTeam with nil Id (Id required to delete)")
	}
	return fc.DeleteCareTeamById(*deleteResource.Id)
}

// delete CareTeam with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCareTeamById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CareTeam", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteCareTeamById: %s", err)
		}
	}
	return &parsed, err
}

// create CatalogEntry, return id of created CatalogEntry or OperationOutcome error
func (fc *FhirClient) CreateCatalogEntry(createResource *r4b.CatalogEntry) (*r4b.CatalogEntry, error) {
	if createResource == nil {
		return nil, errors.New("CreateCatalogEntry called with nil CatalogEntry")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CatalogEntry")
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

	var parsed r4b.CatalogEntry
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateCatalogEntry: %s", err)
		}
	}
	return &parsed, err
}

// read CatalogEntry from fhir server by id, return CatalogEntry or OperationOutcome error
func (fc *FhirClient) ReadCatalogEntry(id string) (*r4b.CatalogEntry, error) {
	if id == "" {
		return nil, errors.New("ReadCatalogEntry given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CatalogEntry", id)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.CatalogEntry
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadCatalogEntry: %s", err)
		}
	}
	return &parsed, err
}

// replace CatalogEntry if exists on server, else create new CatalogEntry with given id, return CatalogEntry or OperationOutcome error
func (fc *FhirClient) UpdateCatalogEntry(updateResource *r4b.CatalogEntry) (*r4b.CatalogEntry, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateCatalogEntry called with nil CatalogEntry")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CatalogEntry", *updateResource.Id)
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

	var parsed r4b.CatalogEntry
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateCatalogEntry: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return CatalogEntry or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCatalogEntry(patchResource *r4b.CatalogEntry) (*r4b.CatalogEntry, error) {
	if patchResource == nil {
		return nil, errors.New("PatchCatalogEntry given nil CatalogEntry")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchCatalogEntry given CatalogEntry without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CatalogEntry", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchCatalogEntry error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchCatalogEntry error using patch representation as json")
	}
	req, err := http.NewRequest(http.MethodPatch, reqUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json-patch+json")

	var parsed r4b.CatalogEntry
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchCatalogEntry: %s", err)
		}
	}
	return &parsed, err
}

// delete CatalogEntry and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCatalogEntry(deleteResource *r4b.CatalogEntry) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCatalogEntry given nil CatalogEntry")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCatalogEntry given CatalogEntry with nil Id (Id required to delete)")
	}
	return fc.DeleteCatalogEntryById(*deleteResource.Id)
}

// delete CatalogEntry with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCatalogEntryById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CatalogEntry", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteCatalogEntryById: %s", err)
		}
	}
	return &parsed, err
}

// create ChargeItem, return id of created ChargeItem or OperationOutcome error
func (fc *FhirClient) CreateChargeItem(createResource *r4b.ChargeItem) (*r4b.ChargeItem, error) {
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

	var parsed r4b.ChargeItem
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateChargeItem: %s", err)
		}
	}
	return &parsed, err
}

// read ChargeItem from fhir server by id, return ChargeItem or OperationOutcome error
func (fc *FhirClient) ReadChargeItem(id string) (*r4b.ChargeItem, error) {
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

	var parsed r4b.ChargeItem
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadChargeItem: %s", err)
		}
	}
	return &parsed, err
}

// replace ChargeItem if exists on server, else create new ChargeItem with given id, return ChargeItem or OperationOutcome error
func (fc *FhirClient) UpdateChargeItem(updateResource *r4b.ChargeItem) (*r4b.ChargeItem, error) {
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

	var parsed r4b.ChargeItem
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateChargeItem: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ChargeItem or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchChargeItem(patchResource *r4b.ChargeItem) (*r4b.ChargeItem, error) {
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

	var parsed r4b.ChargeItem
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchChargeItem: %s", err)
		}
	}
	return &parsed, err
}

// delete ChargeItem and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteChargeItem(deleteResource *r4b.ChargeItem) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteChargeItem given nil ChargeItem")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteChargeItem given ChargeItem with nil Id (Id required to delete)")
	}
	return fc.DeleteChargeItemById(*deleteResource.Id)
}

// delete ChargeItem with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteChargeItemById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ChargeItem", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteChargeItemById: %s", err)
		}
	}
	return &parsed, err
}

// create ChargeItemDefinition, return id of created ChargeItemDefinition or OperationOutcome error
func (fc *FhirClient) CreateChargeItemDefinition(createResource *r4b.ChargeItemDefinition) (*r4b.ChargeItemDefinition, error) {
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

	var parsed r4b.ChargeItemDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateChargeItemDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read ChargeItemDefinition from fhir server by id, return ChargeItemDefinition or OperationOutcome error
func (fc *FhirClient) ReadChargeItemDefinition(id string) (*r4b.ChargeItemDefinition, error) {
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

	var parsed r4b.ChargeItemDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadChargeItemDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace ChargeItemDefinition if exists on server, else create new ChargeItemDefinition with given id, return ChargeItemDefinition or OperationOutcome error
func (fc *FhirClient) UpdateChargeItemDefinition(updateResource *r4b.ChargeItemDefinition) (*r4b.ChargeItemDefinition, error) {
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

	var parsed r4b.ChargeItemDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateChargeItemDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ChargeItemDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchChargeItemDefinition(patchResource *r4b.ChargeItemDefinition) (*r4b.ChargeItemDefinition, error) {
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

	var parsed r4b.ChargeItemDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchChargeItemDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete ChargeItemDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteChargeItemDefinition(deleteResource *r4b.ChargeItemDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteChargeItemDefinition given nil ChargeItemDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteChargeItemDefinition given ChargeItemDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteChargeItemDefinitionById(*deleteResource.Id)
}

// delete ChargeItemDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteChargeItemDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ChargeItemDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteChargeItemDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create Citation, return id of created Citation or OperationOutcome error
func (fc *FhirClient) CreateCitation(createResource *r4b.Citation) (*r4b.Citation, error) {
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

	var parsed r4b.Citation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateCitation: %s", err)
		}
	}
	return &parsed, err
}

// read Citation from fhir server by id, return Citation or OperationOutcome error
func (fc *FhirClient) ReadCitation(id string) (*r4b.Citation, error) {
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

	var parsed r4b.Citation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadCitation: %s", err)
		}
	}
	return &parsed, err
}

// replace Citation if exists on server, else create new Citation with given id, return Citation or OperationOutcome error
func (fc *FhirClient) UpdateCitation(updateResource *r4b.Citation) (*r4b.Citation, error) {
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

	var parsed r4b.Citation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateCitation: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Citation or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCitation(patchResource *r4b.Citation) (*r4b.Citation, error) {
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

	var parsed r4b.Citation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchCitation: %s", err)
		}
	}
	return &parsed, err
}

// delete Citation and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCitation(deleteResource *r4b.Citation) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCitation given nil Citation")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCitation given Citation with nil Id (Id required to delete)")
	}
	return fc.DeleteCitationById(*deleteResource.Id)
}

// delete Citation with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCitationById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Citation", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteCitationById: %s", err)
		}
	}
	return &parsed, err
}

// create Claim, return id of created Claim or OperationOutcome error
func (fc *FhirClient) CreateClaim(createResource *r4b.Claim) (*r4b.Claim, error) {
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

	var parsed r4b.Claim
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateClaim: %s", err)
		}
	}
	return &parsed, err
}

// read Claim from fhir server by id, return Claim or OperationOutcome error
func (fc *FhirClient) ReadClaim(id string) (*r4b.Claim, error) {
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

	var parsed r4b.Claim
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadClaim: %s", err)
		}
	}
	return &parsed, err
}

// replace Claim if exists on server, else create new Claim with given id, return Claim or OperationOutcome error
func (fc *FhirClient) UpdateClaim(updateResource *r4b.Claim) (*r4b.Claim, error) {
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

	var parsed r4b.Claim
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateClaim: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Claim or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchClaim(patchResource *r4b.Claim) (*r4b.Claim, error) {
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

	var parsed r4b.Claim
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchClaim: %s", err)
		}
	}
	return &parsed, err
}

// delete Claim and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteClaim(deleteResource *r4b.Claim) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteClaim given nil Claim")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteClaim given Claim with nil Id (Id required to delete)")
	}
	return fc.DeleteClaimById(*deleteResource.Id)
}

// delete Claim with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteClaimById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Claim", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteClaimById: %s", err)
		}
	}
	return &parsed, err
}

// create ClaimResponse, return id of created ClaimResponse or OperationOutcome error
func (fc *FhirClient) CreateClaimResponse(createResource *r4b.ClaimResponse) (*r4b.ClaimResponse, error) {
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

	var parsed r4b.ClaimResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateClaimResponse: %s", err)
		}
	}
	return &parsed, err
}

// read ClaimResponse from fhir server by id, return ClaimResponse or OperationOutcome error
func (fc *FhirClient) ReadClaimResponse(id string) (*r4b.ClaimResponse, error) {
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

	var parsed r4b.ClaimResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadClaimResponse: %s", err)
		}
	}
	return &parsed, err
}

// replace ClaimResponse if exists on server, else create new ClaimResponse with given id, return ClaimResponse or OperationOutcome error
func (fc *FhirClient) UpdateClaimResponse(updateResource *r4b.ClaimResponse) (*r4b.ClaimResponse, error) {
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

	var parsed r4b.ClaimResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateClaimResponse: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ClaimResponse or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchClaimResponse(patchResource *r4b.ClaimResponse) (*r4b.ClaimResponse, error) {
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

	var parsed r4b.ClaimResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchClaimResponse: %s", err)
		}
	}
	return &parsed, err
}

// delete ClaimResponse and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteClaimResponse(deleteResource *r4b.ClaimResponse) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteClaimResponse given nil ClaimResponse")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteClaimResponse given ClaimResponse with nil Id (Id required to delete)")
	}
	return fc.DeleteClaimResponseById(*deleteResource.Id)
}

// delete ClaimResponse with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteClaimResponseById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClaimResponse", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteClaimResponseById: %s", err)
		}
	}
	return &parsed, err
}

// create ClinicalImpression, return id of created ClinicalImpression or OperationOutcome error
func (fc *FhirClient) CreateClinicalImpression(createResource *r4b.ClinicalImpression) (*r4b.ClinicalImpression, error) {
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

	var parsed r4b.ClinicalImpression
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateClinicalImpression: %s", err)
		}
	}
	return &parsed, err
}

// read ClinicalImpression from fhir server by id, return ClinicalImpression or OperationOutcome error
func (fc *FhirClient) ReadClinicalImpression(id string) (*r4b.ClinicalImpression, error) {
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

	var parsed r4b.ClinicalImpression
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadClinicalImpression: %s", err)
		}
	}
	return &parsed, err
}

// replace ClinicalImpression if exists on server, else create new ClinicalImpression with given id, return ClinicalImpression or OperationOutcome error
func (fc *FhirClient) UpdateClinicalImpression(updateResource *r4b.ClinicalImpression) (*r4b.ClinicalImpression, error) {
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

	var parsed r4b.ClinicalImpression
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateClinicalImpression: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ClinicalImpression or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchClinicalImpression(patchResource *r4b.ClinicalImpression) (*r4b.ClinicalImpression, error) {
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

	var parsed r4b.ClinicalImpression
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchClinicalImpression: %s", err)
		}
	}
	return &parsed, err
}

// delete ClinicalImpression and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteClinicalImpression(deleteResource *r4b.ClinicalImpression) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteClinicalImpression given nil ClinicalImpression")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteClinicalImpression given ClinicalImpression with nil Id (Id required to delete)")
	}
	return fc.DeleteClinicalImpressionById(*deleteResource.Id)
}

// delete ClinicalImpression with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteClinicalImpressionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClinicalImpression", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteClinicalImpressionById: %s", err)
		}
	}
	return &parsed, err
}

// create ClinicalUseDefinition, return id of created ClinicalUseDefinition or OperationOutcome error
func (fc *FhirClient) CreateClinicalUseDefinition(createResource *r4b.ClinicalUseDefinition) (*r4b.ClinicalUseDefinition, error) {
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

	var parsed r4b.ClinicalUseDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateClinicalUseDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read ClinicalUseDefinition from fhir server by id, return ClinicalUseDefinition or OperationOutcome error
func (fc *FhirClient) ReadClinicalUseDefinition(id string) (*r4b.ClinicalUseDefinition, error) {
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

	var parsed r4b.ClinicalUseDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadClinicalUseDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace ClinicalUseDefinition if exists on server, else create new ClinicalUseDefinition with given id, return ClinicalUseDefinition or OperationOutcome error
func (fc *FhirClient) UpdateClinicalUseDefinition(updateResource *r4b.ClinicalUseDefinition) (*r4b.ClinicalUseDefinition, error) {
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

	var parsed r4b.ClinicalUseDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateClinicalUseDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ClinicalUseDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchClinicalUseDefinition(patchResource *r4b.ClinicalUseDefinition) (*r4b.ClinicalUseDefinition, error) {
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

	var parsed r4b.ClinicalUseDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchClinicalUseDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete ClinicalUseDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteClinicalUseDefinition(deleteResource *r4b.ClinicalUseDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteClinicalUseDefinition given nil ClinicalUseDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteClinicalUseDefinition given ClinicalUseDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteClinicalUseDefinitionById(*deleteResource.Id)
}

// delete ClinicalUseDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteClinicalUseDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ClinicalUseDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteClinicalUseDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create CodeSystem, return id of created CodeSystem or OperationOutcome error
func (fc *FhirClient) CreateCodeSystem(createResource *r4b.CodeSystem) (*r4b.CodeSystem, error) {
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

	var parsed r4b.CodeSystem
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateCodeSystem: %s", err)
		}
	}
	return &parsed, err
}

// read CodeSystem from fhir server by id, return CodeSystem or OperationOutcome error
func (fc *FhirClient) ReadCodeSystem(id string) (*r4b.CodeSystem, error) {
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

	var parsed r4b.CodeSystem
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadCodeSystem: %s", err)
		}
	}
	return &parsed, err
}

// replace CodeSystem if exists on server, else create new CodeSystem with given id, return CodeSystem or OperationOutcome error
func (fc *FhirClient) UpdateCodeSystem(updateResource *r4b.CodeSystem) (*r4b.CodeSystem, error) {
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

	var parsed r4b.CodeSystem
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateCodeSystem: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return CodeSystem or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCodeSystem(patchResource *r4b.CodeSystem) (*r4b.CodeSystem, error) {
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

	var parsed r4b.CodeSystem
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchCodeSystem: %s", err)
		}
	}
	return &parsed, err
}

// delete CodeSystem and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCodeSystem(deleteResource *r4b.CodeSystem) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCodeSystem given nil CodeSystem")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCodeSystem given CodeSystem with nil Id (Id required to delete)")
	}
	return fc.DeleteCodeSystemById(*deleteResource.Id)
}

// delete CodeSystem with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCodeSystemById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CodeSystem", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteCodeSystemById: %s", err)
		}
	}
	return &parsed, err
}

// create Communication, return id of created Communication or OperationOutcome error
func (fc *FhirClient) CreateCommunication(createResource *r4b.Communication) (*r4b.Communication, error) {
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

	var parsed r4b.Communication
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateCommunication: %s", err)
		}
	}
	return &parsed, err
}

// read Communication from fhir server by id, return Communication or OperationOutcome error
func (fc *FhirClient) ReadCommunication(id string) (*r4b.Communication, error) {
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

	var parsed r4b.Communication
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadCommunication: %s", err)
		}
	}
	return &parsed, err
}

// replace Communication if exists on server, else create new Communication with given id, return Communication or OperationOutcome error
func (fc *FhirClient) UpdateCommunication(updateResource *r4b.Communication) (*r4b.Communication, error) {
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

	var parsed r4b.Communication
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateCommunication: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Communication or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCommunication(patchResource *r4b.Communication) (*r4b.Communication, error) {
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

	var parsed r4b.Communication
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchCommunication: %s", err)
		}
	}
	return &parsed, err
}

// delete Communication and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCommunication(deleteResource *r4b.Communication) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCommunication given nil Communication")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCommunication given Communication with nil Id (Id required to delete)")
	}
	return fc.DeleteCommunicationById(*deleteResource.Id)
}

// delete Communication with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCommunicationById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Communication", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteCommunicationById: %s", err)
		}
	}
	return &parsed, err
}

// create CommunicationRequest, return id of created CommunicationRequest or OperationOutcome error
func (fc *FhirClient) CreateCommunicationRequest(createResource *r4b.CommunicationRequest) (*r4b.CommunicationRequest, error) {
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

	var parsed r4b.CommunicationRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateCommunicationRequest: %s", err)
		}
	}
	return &parsed, err
}

// read CommunicationRequest from fhir server by id, return CommunicationRequest or OperationOutcome error
func (fc *FhirClient) ReadCommunicationRequest(id string) (*r4b.CommunicationRequest, error) {
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

	var parsed r4b.CommunicationRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadCommunicationRequest: %s", err)
		}
	}
	return &parsed, err
}

// replace CommunicationRequest if exists on server, else create new CommunicationRequest with given id, return CommunicationRequest or OperationOutcome error
func (fc *FhirClient) UpdateCommunicationRequest(updateResource *r4b.CommunicationRequest) (*r4b.CommunicationRequest, error) {
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

	var parsed r4b.CommunicationRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateCommunicationRequest: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return CommunicationRequest or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCommunicationRequest(patchResource *r4b.CommunicationRequest) (*r4b.CommunicationRequest, error) {
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

	var parsed r4b.CommunicationRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchCommunicationRequest: %s", err)
		}
	}
	return &parsed, err
}

// delete CommunicationRequest and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCommunicationRequest(deleteResource *r4b.CommunicationRequest) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCommunicationRequest given nil CommunicationRequest")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCommunicationRequest given CommunicationRequest with nil Id (Id required to delete)")
	}
	return fc.DeleteCommunicationRequestById(*deleteResource.Id)
}

// delete CommunicationRequest with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCommunicationRequestById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CommunicationRequest", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteCommunicationRequestById: %s", err)
		}
	}
	return &parsed, err
}

// create CompartmentDefinition, return id of created CompartmentDefinition or OperationOutcome error
func (fc *FhirClient) CreateCompartmentDefinition(createResource *r4b.CompartmentDefinition) (*r4b.CompartmentDefinition, error) {
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

	var parsed r4b.CompartmentDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateCompartmentDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read CompartmentDefinition from fhir server by id, return CompartmentDefinition or OperationOutcome error
func (fc *FhirClient) ReadCompartmentDefinition(id string) (*r4b.CompartmentDefinition, error) {
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

	var parsed r4b.CompartmentDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadCompartmentDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace CompartmentDefinition if exists on server, else create new CompartmentDefinition with given id, return CompartmentDefinition or OperationOutcome error
func (fc *FhirClient) UpdateCompartmentDefinition(updateResource *r4b.CompartmentDefinition) (*r4b.CompartmentDefinition, error) {
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

	var parsed r4b.CompartmentDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateCompartmentDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return CompartmentDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCompartmentDefinition(patchResource *r4b.CompartmentDefinition) (*r4b.CompartmentDefinition, error) {
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

	var parsed r4b.CompartmentDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchCompartmentDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete CompartmentDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCompartmentDefinition(deleteResource *r4b.CompartmentDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCompartmentDefinition given nil CompartmentDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCompartmentDefinition given CompartmentDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteCompartmentDefinitionById(*deleteResource.Id)
}

// delete CompartmentDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCompartmentDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CompartmentDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteCompartmentDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create Composition, return id of created Composition or OperationOutcome error
func (fc *FhirClient) CreateComposition(createResource *r4b.Composition) (*r4b.Composition, error) {
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

	var parsed r4b.Composition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateComposition: %s", err)
		}
	}
	return &parsed, err
}

// read Composition from fhir server by id, return Composition or OperationOutcome error
func (fc *FhirClient) ReadComposition(id string) (*r4b.Composition, error) {
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

	var parsed r4b.Composition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadComposition: %s", err)
		}
	}
	return &parsed, err
}

// replace Composition if exists on server, else create new Composition with given id, return Composition or OperationOutcome error
func (fc *FhirClient) UpdateComposition(updateResource *r4b.Composition) (*r4b.Composition, error) {
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

	var parsed r4b.Composition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateComposition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Composition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchComposition(patchResource *r4b.Composition) (*r4b.Composition, error) {
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

	var parsed r4b.Composition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchComposition: %s", err)
		}
	}
	return &parsed, err
}

// delete Composition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteComposition(deleteResource *r4b.Composition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteComposition given nil Composition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteComposition given Composition with nil Id (Id required to delete)")
	}
	return fc.DeleteCompositionById(*deleteResource.Id)
}

// delete Composition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCompositionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Composition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteCompositionById: %s", err)
		}
	}
	return &parsed, err
}

// create ConceptMap, return id of created ConceptMap or OperationOutcome error
func (fc *FhirClient) CreateConceptMap(createResource *r4b.ConceptMap) (*r4b.ConceptMap, error) {
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

	var parsed r4b.ConceptMap
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateConceptMap: %s", err)
		}
	}
	return &parsed, err
}

// read ConceptMap from fhir server by id, return ConceptMap or OperationOutcome error
func (fc *FhirClient) ReadConceptMap(id string) (*r4b.ConceptMap, error) {
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

	var parsed r4b.ConceptMap
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadConceptMap: %s", err)
		}
	}
	return &parsed, err
}

// replace ConceptMap if exists on server, else create new ConceptMap with given id, return ConceptMap or OperationOutcome error
func (fc *FhirClient) UpdateConceptMap(updateResource *r4b.ConceptMap) (*r4b.ConceptMap, error) {
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

	var parsed r4b.ConceptMap
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateConceptMap: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ConceptMap or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchConceptMap(patchResource *r4b.ConceptMap) (*r4b.ConceptMap, error) {
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

	var parsed r4b.ConceptMap
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchConceptMap: %s", err)
		}
	}
	return &parsed, err
}

// delete ConceptMap and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteConceptMap(deleteResource *r4b.ConceptMap) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteConceptMap given nil ConceptMap")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteConceptMap given ConceptMap with nil Id (Id required to delete)")
	}
	return fc.DeleteConceptMapById(*deleteResource.Id)
}

// delete ConceptMap with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteConceptMapById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ConceptMap", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteConceptMapById: %s", err)
		}
	}
	return &parsed, err
}

// create Condition, return id of created Condition or OperationOutcome error
func (fc *FhirClient) CreateCondition(createResource *r4b.Condition) (*r4b.Condition, error) {
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

	var parsed r4b.Condition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateCondition: %s", err)
		}
	}
	return &parsed, err
}

// read Condition from fhir server by id, return Condition or OperationOutcome error
func (fc *FhirClient) ReadCondition(id string) (*r4b.Condition, error) {
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

	var parsed r4b.Condition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadCondition: %s", err)
		}
	}
	return &parsed, err
}

// replace Condition if exists on server, else create new Condition with given id, return Condition or OperationOutcome error
func (fc *FhirClient) UpdateCondition(updateResource *r4b.Condition) (*r4b.Condition, error) {
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

	var parsed r4b.Condition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateCondition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Condition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCondition(patchResource *r4b.Condition) (*r4b.Condition, error) {
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

	var parsed r4b.Condition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchCondition: %s", err)
		}
	}
	return &parsed, err
}

// delete Condition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCondition(deleteResource *r4b.Condition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCondition given nil Condition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCondition given Condition with nil Id (Id required to delete)")
	}
	return fc.DeleteConditionById(*deleteResource.Id)
}

// delete Condition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteConditionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Condition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteConditionById: %s", err)
		}
	}
	return &parsed, err
}

// create Consent, return id of created Consent or OperationOutcome error
func (fc *FhirClient) CreateConsent(createResource *r4b.Consent) (*r4b.Consent, error) {
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

	var parsed r4b.Consent
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateConsent: %s", err)
		}
	}
	return &parsed, err
}

// read Consent from fhir server by id, return Consent or OperationOutcome error
func (fc *FhirClient) ReadConsent(id string) (*r4b.Consent, error) {
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

	var parsed r4b.Consent
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadConsent: %s", err)
		}
	}
	return &parsed, err
}

// replace Consent if exists on server, else create new Consent with given id, return Consent or OperationOutcome error
func (fc *FhirClient) UpdateConsent(updateResource *r4b.Consent) (*r4b.Consent, error) {
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

	var parsed r4b.Consent
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateConsent: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Consent or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchConsent(patchResource *r4b.Consent) (*r4b.Consent, error) {
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

	var parsed r4b.Consent
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchConsent: %s", err)
		}
	}
	return &parsed, err
}

// delete Consent and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteConsent(deleteResource *r4b.Consent) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteConsent given nil Consent")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteConsent given Consent with nil Id (Id required to delete)")
	}
	return fc.DeleteConsentById(*deleteResource.Id)
}

// delete Consent with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteConsentById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Consent", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteConsentById: %s", err)
		}
	}
	return &parsed, err
}

// create Contract, return id of created Contract or OperationOutcome error
func (fc *FhirClient) CreateContract(createResource *r4b.Contract) (*r4b.Contract, error) {
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

	var parsed r4b.Contract
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateContract: %s", err)
		}
	}
	return &parsed, err
}

// read Contract from fhir server by id, return Contract or OperationOutcome error
func (fc *FhirClient) ReadContract(id string) (*r4b.Contract, error) {
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

	var parsed r4b.Contract
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadContract: %s", err)
		}
	}
	return &parsed, err
}

// replace Contract if exists on server, else create new Contract with given id, return Contract or OperationOutcome error
func (fc *FhirClient) UpdateContract(updateResource *r4b.Contract) (*r4b.Contract, error) {
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

	var parsed r4b.Contract
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateContract: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Contract or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchContract(patchResource *r4b.Contract) (*r4b.Contract, error) {
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

	var parsed r4b.Contract
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchContract: %s", err)
		}
	}
	return &parsed, err
}

// delete Contract and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteContract(deleteResource *r4b.Contract) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteContract given nil Contract")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteContract given Contract with nil Id (Id required to delete)")
	}
	return fc.DeleteContractById(*deleteResource.Id)
}

// delete Contract with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteContractById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Contract", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteContractById: %s", err)
		}
	}
	return &parsed, err
}

// create Coverage, return id of created Coverage or OperationOutcome error
func (fc *FhirClient) CreateCoverage(createResource *r4b.Coverage) (*r4b.Coverage, error) {
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

	var parsed r4b.Coverage
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateCoverage: %s", err)
		}
	}
	return &parsed, err
}

// read Coverage from fhir server by id, return Coverage or OperationOutcome error
func (fc *FhirClient) ReadCoverage(id string) (*r4b.Coverage, error) {
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

	var parsed r4b.Coverage
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadCoverage: %s", err)
		}
	}
	return &parsed, err
}

// replace Coverage if exists on server, else create new Coverage with given id, return Coverage or OperationOutcome error
func (fc *FhirClient) UpdateCoverage(updateResource *r4b.Coverage) (*r4b.Coverage, error) {
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

	var parsed r4b.Coverage
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateCoverage: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Coverage or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCoverage(patchResource *r4b.Coverage) (*r4b.Coverage, error) {
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

	var parsed r4b.Coverage
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchCoverage: %s", err)
		}
	}
	return &parsed, err
}

// delete Coverage and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCoverage(deleteResource *r4b.Coverage) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCoverage given nil Coverage")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCoverage given Coverage with nil Id (Id required to delete)")
	}
	return fc.DeleteCoverageById(*deleteResource.Id)
}

// delete Coverage with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCoverageById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Coverage", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteCoverageById: %s", err)
		}
	}
	return &parsed, err
}

// create CoverageEligibilityRequest, return id of created CoverageEligibilityRequest or OperationOutcome error
func (fc *FhirClient) CreateCoverageEligibilityRequest(createResource *r4b.CoverageEligibilityRequest) (*r4b.CoverageEligibilityRequest, error) {
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

	var parsed r4b.CoverageEligibilityRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateCoverageEligibilityRequest: %s", err)
		}
	}
	return &parsed, err
}

// read CoverageEligibilityRequest from fhir server by id, return CoverageEligibilityRequest or OperationOutcome error
func (fc *FhirClient) ReadCoverageEligibilityRequest(id string) (*r4b.CoverageEligibilityRequest, error) {
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

	var parsed r4b.CoverageEligibilityRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadCoverageEligibilityRequest: %s", err)
		}
	}
	return &parsed, err
}

// replace CoverageEligibilityRequest if exists on server, else create new CoverageEligibilityRequest with given id, return CoverageEligibilityRequest or OperationOutcome error
func (fc *FhirClient) UpdateCoverageEligibilityRequest(updateResource *r4b.CoverageEligibilityRequest) (*r4b.CoverageEligibilityRequest, error) {
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

	var parsed r4b.CoverageEligibilityRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateCoverageEligibilityRequest: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return CoverageEligibilityRequest or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCoverageEligibilityRequest(patchResource *r4b.CoverageEligibilityRequest) (*r4b.CoverageEligibilityRequest, error) {
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

	var parsed r4b.CoverageEligibilityRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchCoverageEligibilityRequest: %s", err)
		}
	}
	return &parsed, err
}

// delete CoverageEligibilityRequest and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCoverageEligibilityRequest(deleteResource *r4b.CoverageEligibilityRequest) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCoverageEligibilityRequest given nil CoverageEligibilityRequest")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCoverageEligibilityRequest given CoverageEligibilityRequest with nil Id (Id required to delete)")
	}
	return fc.DeleteCoverageEligibilityRequestById(*deleteResource.Id)
}

// delete CoverageEligibilityRequest with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCoverageEligibilityRequestById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CoverageEligibilityRequest", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteCoverageEligibilityRequestById: %s", err)
		}
	}
	return &parsed, err
}

// create CoverageEligibilityResponse, return id of created CoverageEligibilityResponse or OperationOutcome error
func (fc *FhirClient) CreateCoverageEligibilityResponse(createResource *r4b.CoverageEligibilityResponse) (*r4b.CoverageEligibilityResponse, error) {
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

	var parsed r4b.CoverageEligibilityResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateCoverageEligibilityResponse: %s", err)
		}
	}
	return &parsed, err
}

// read CoverageEligibilityResponse from fhir server by id, return CoverageEligibilityResponse or OperationOutcome error
func (fc *FhirClient) ReadCoverageEligibilityResponse(id string) (*r4b.CoverageEligibilityResponse, error) {
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

	var parsed r4b.CoverageEligibilityResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadCoverageEligibilityResponse: %s", err)
		}
	}
	return &parsed, err
}

// replace CoverageEligibilityResponse if exists on server, else create new CoverageEligibilityResponse with given id, return CoverageEligibilityResponse or OperationOutcome error
func (fc *FhirClient) UpdateCoverageEligibilityResponse(updateResource *r4b.CoverageEligibilityResponse) (*r4b.CoverageEligibilityResponse, error) {
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

	var parsed r4b.CoverageEligibilityResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateCoverageEligibilityResponse: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return CoverageEligibilityResponse or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchCoverageEligibilityResponse(patchResource *r4b.CoverageEligibilityResponse) (*r4b.CoverageEligibilityResponse, error) {
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

	var parsed r4b.CoverageEligibilityResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchCoverageEligibilityResponse: %s", err)
		}
	}
	return &parsed, err
}

// delete CoverageEligibilityResponse and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCoverageEligibilityResponse(deleteResource *r4b.CoverageEligibilityResponse) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteCoverageEligibilityResponse given nil CoverageEligibilityResponse")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteCoverageEligibilityResponse given CoverageEligibilityResponse with nil Id (Id required to delete)")
	}
	return fc.DeleteCoverageEligibilityResponseById(*deleteResource.Id)
}

// delete CoverageEligibilityResponse with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteCoverageEligibilityResponseById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "CoverageEligibilityResponse", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteCoverageEligibilityResponseById: %s", err)
		}
	}
	return &parsed, err
}

// create DetectedIssue, return id of created DetectedIssue or OperationOutcome error
func (fc *FhirClient) CreateDetectedIssue(createResource *r4b.DetectedIssue) (*r4b.DetectedIssue, error) {
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

	var parsed r4b.DetectedIssue
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateDetectedIssue: %s", err)
		}
	}
	return &parsed, err
}

// read DetectedIssue from fhir server by id, return DetectedIssue or OperationOutcome error
func (fc *FhirClient) ReadDetectedIssue(id string) (*r4b.DetectedIssue, error) {
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

	var parsed r4b.DetectedIssue
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadDetectedIssue: %s", err)
		}
	}
	return &parsed, err
}

// replace DetectedIssue if exists on server, else create new DetectedIssue with given id, return DetectedIssue or OperationOutcome error
func (fc *FhirClient) UpdateDetectedIssue(updateResource *r4b.DetectedIssue) (*r4b.DetectedIssue, error) {
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

	var parsed r4b.DetectedIssue
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateDetectedIssue: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return DetectedIssue or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDetectedIssue(patchResource *r4b.DetectedIssue) (*r4b.DetectedIssue, error) {
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

	var parsed r4b.DetectedIssue
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchDetectedIssue: %s", err)
		}
	}
	return &parsed, err
}

// delete DetectedIssue and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDetectedIssue(deleteResource *r4b.DetectedIssue) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDetectedIssue given nil DetectedIssue")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDetectedIssue given DetectedIssue with nil Id (Id required to delete)")
	}
	return fc.DeleteDetectedIssueById(*deleteResource.Id)
}

// delete DetectedIssue with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDetectedIssueById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DetectedIssue", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteDetectedIssueById: %s", err)
		}
	}
	return &parsed, err
}

// create Device, return id of created Device or OperationOutcome error
func (fc *FhirClient) CreateDevice(createResource *r4b.Device) (*r4b.Device, error) {
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

	var parsed r4b.Device
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateDevice: %s", err)
		}
	}
	return &parsed, err
}

// read Device from fhir server by id, return Device or OperationOutcome error
func (fc *FhirClient) ReadDevice(id string) (*r4b.Device, error) {
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

	var parsed r4b.Device
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadDevice: %s", err)
		}
	}
	return &parsed, err
}

// replace Device if exists on server, else create new Device with given id, return Device or OperationOutcome error
func (fc *FhirClient) UpdateDevice(updateResource *r4b.Device) (*r4b.Device, error) {
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

	var parsed r4b.Device
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateDevice: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Device or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDevice(patchResource *r4b.Device) (*r4b.Device, error) {
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

	var parsed r4b.Device
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchDevice: %s", err)
		}
	}
	return &parsed, err
}

// delete Device and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDevice(deleteResource *r4b.Device) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDevice given nil Device")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDevice given Device with nil Id (Id required to delete)")
	}
	return fc.DeleteDeviceById(*deleteResource.Id)
}

// delete Device with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Device", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteDeviceById: %s", err)
		}
	}
	return &parsed, err
}

// create DeviceDefinition, return id of created DeviceDefinition or OperationOutcome error
func (fc *FhirClient) CreateDeviceDefinition(createResource *r4b.DeviceDefinition) (*r4b.DeviceDefinition, error) {
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

	var parsed r4b.DeviceDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateDeviceDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read DeviceDefinition from fhir server by id, return DeviceDefinition or OperationOutcome error
func (fc *FhirClient) ReadDeviceDefinition(id string) (*r4b.DeviceDefinition, error) {
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

	var parsed r4b.DeviceDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadDeviceDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace DeviceDefinition if exists on server, else create new DeviceDefinition with given id, return DeviceDefinition or OperationOutcome error
func (fc *FhirClient) UpdateDeviceDefinition(updateResource *r4b.DeviceDefinition) (*r4b.DeviceDefinition, error) {
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

	var parsed r4b.DeviceDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateDeviceDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return DeviceDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDeviceDefinition(patchResource *r4b.DeviceDefinition) (*r4b.DeviceDefinition, error) {
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

	var parsed r4b.DeviceDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchDeviceDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete DeviceDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceDefinition(deleteResource *r4b.DeviceDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDeviceDefinition given nil DeviceDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDeviceDefinition given DeviceDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteDeviceDefinitionById(*deleteResource.Id)
}

// delete DeviceDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteDeviceDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create DeviceMetric, return id of created DeviceMetric or OperationOutcome error
func (fc *FhirClient) CreateDeviceMetric(createResource *r4b.DeviceMetric) (*r4b.DeviceMetric, error) {
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

	var parsed r4b.DeviceMetric
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateDeviceMetric: %s", err)
		}
	}
	return &parsed, err
}

// read DeviceMetric from fhir server by id, return DeviceMetric or OperationOutcome error
func (fc *FhirClient) ReadDeviceMetric(id string) (*r4b.DeviceMetric, error) {
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

	var parsed r4b.DeviceMetric
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadDeviceMetric: %s", err)
		}
	}
	return &parsed, err
}

// replace DeviceMetric if exists on server, else create new DeviceMetric with given id, return DeviceMetric or OperationOutcome error
func (fc *FhirClient) UpdateDeviceMetric(updateResource *r4b.DeviceMetric) (*r4b.DeviceMetric, error) {
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

	var parsed r4b.DeviceMetric
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateDeviceMetric: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return DeviceMetric or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDeviceMetric(patchResource *r4b.DeviceMetric) (*r4b.DeviceMetric, error) {
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

	var parsed r4b.DeviceMetric
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchDeviceMetric: %s", err)
		}
	}
	return &parsed, err
}

// delete DeviceMetric and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceMetric(deleteResource *r4b.DeviceMetric) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDeviceMetric given nil DeviceMetric")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDeviceMetric given DeviceMetric with nil Id (Id required to delete)")
	}
	return fc.DeleteDeviceMetricById(*deleteResource.Id)
}

// delete DeviceMetric with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceMetricById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceMetric", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteDeviceMetricById: %s", err)
		}
	}
	return &parsed, err
}

// create DeviceRequest, return id of created DeviceRequest or OperationOutcome error
func (fc *FhirClient) CreateDeviceRequest(createResource *r4b.DeviceRequest) (*r4b.DeviceRequest, error) {
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

	var parsed r4b.DeviceRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateDeviceRequest: %s", err)
		}
	}
	return &parsed, err
}

// read DeviceRequest from fhir server by id, return DeviceRequest or OperationOutcome error
func (fc *FhirClient) ReadDeviceRequest(id string) (*r4b.DeviceRequest, error) {
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

	var parsed r4b.DeviceRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadDeviceRequest: %s", err)
		}
	}
	return &parsed, err
}

// replace DeviceRequest if exists on server, else create new DeviceRequest with given id, return DeviceRequest or OperationOutcome error
func (fc *FhirClient) UpdateDeviceRequest(updateResource *r4b.DeviceRequest) (*r4b.DeviceRequest, error) {
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

	var parsed r4b.DeviceRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateDeviceRequest: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return DeviceRequest or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDeviceRequest(patchResource *r4b.DeviceRequest) (*r4b.DeviceRequest, error) {
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

	var parsed r4b.DeviceRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchDeviceRequest: %s", err)
		}
	}
	return &parsed, err
}

// delete DeviceRequest and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceRequest(deleteResource *r4b.DeviceRequest) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDeviceRequest given nil DeviceRequest")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDeviceRequest given DeviceRequest with nil Id (Id required to delete)")
	}
	return fc.DeleteDeviceRequestById(*deleteResource.Id)
}

// delete DeviceRequest with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceRequestById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceRequest", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteDeviceRequestById: %s", err)
		}
	}
	return &parsed, err
}

// create DeviceUseStatement, return id of created DeviceUseStatement or OperationOutcome error
func (fc *FhirClient) CreateDeviceUseStatement(createResource *r4b.DeviceUseStatement) (*r4b.DeviceUseStatement, error) {
	if createResource == nil {
		return nil, errors.New("CreateDeviceUseStatement called with nil DeviceUseStatement")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceUseStatement")
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

	var parsed r4b.DeviceUseStatement
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateDeviceUseStatement: %s", err)
		}
	}
	return &parsed, err
}

// read DeviceUseStatement from fhir server by id, return DeviceUseStatement or OperationOutcome error
func (fc *FhirClient) ReadDeviceUseStatement(id string) (*r4b.DeviceUseStatement, error) {
	if id == "" {
		return nil, errors.New("ReadDeviceUseStatement given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceUseStatement", id)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.DeviceUseStatement
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadDeviceUseStatement: %s", err)
		}
	}
	return &parsed, err
}

// replace DeviceUseStatement if exists on server, else create new DeviceUseStatement with given id, return DeviceUseStatement or OperationOutcome error
func (fc *FhirClient) UpdateDeviceUseStatement(updateResource *r4b.DeviceUseStatement) (*r4b.DeviceUseStatement, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateDeviceUseStatement called with nil DeviceUseStatement")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceUseStatement", *updateResource.Id)
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

	var parsed r4b.DeviceUseStatement
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateDeviceUseStatement: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return DeviceUseStatement or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDeviceUseStatement(patchResource *r4b.DeviceUseStatement) (*r4b.DeviceUseStatement, error) {
	if patchResource == nil {
		return nil, errors.New("PatchDeviceUseStatement given nil DeviceUseStatement")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchDeviceUseStatement given DeviceUseStatement without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceUseStatement", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchDeviceUseStatement error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchDeviceUseStatement error using patch representation as json")
	}
	req, err := http.NewRequest(http.MethodPatch, reqUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json-patch+json")

	var parsed r4b.DeviceUseStatement
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchDeviceUseStatement: %s", err)
		}
	}
	return &parsed, err
}

// delete DeviceUseStatement and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceUseStatement(deleteResource *r4b.DeviceUseStatement) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDeviceUseStatement given nil DeviceUseStatement")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDeviceUseStatement given DeviceUseStatement with nil Id (Id required to delete)")
	}
	return fc.DeleteDeviceUseStatementById(*deleteResource.Id)
}

// delete DeviceUseStatement with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDeviceUseStatementById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DeviceUseStatement", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteDeviceUseStatementById: %s", err)
		}
	}
	return &parsed, err
}

// create DiagnosticReport, return id of created DiagnosticReport or OperationOutcome error
func (fc *FhirClient) CreateDiagnosticReport(createResource *r4b.DiagnosticReport) (*r4b.DiagnosticReport, error) {
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

	var parsed r4b.DiagnosticReport
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateDiagnosticReport: %s", err)
		}
	}
	return &parsed, err
}

// read DiagnosticReport from fhir server by id, return DiagnosticReport or OperationOutcome error
func (fc *FhirClient) ReadDiagnosticReport(id string) (*r4b.DiagnosticReport, error) {
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

	var parsed r4b.DiagnosticReport
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadDiagnosticReport: %s", err)
		}
	}
	return &parsed, err
}

// replace DiagnosticReport if exists on server, else create new DiagnosticReport with given id, return DiagnosticReport or OperationOutcome error
func (fc *FhirClient) UpdateDiagnosticReport(updateResource *r4b.DiagnosticReport) (*r4b.DiagnosticReport, error) {
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

	var parsed r4b.DiagnosticReport
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateDiagnosticReport: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return DiagnosticReport or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDiagnosticReport(patchResource *r4b.DiagnosticReport) (*r4b.DiagnosticReport, error) {
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

	var parsed r4b.DiagnosticReport
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchDiagnosticReport: %s", err)
		}
	}
	return &parsed, err
}

// delete DiagnosticReport and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDiagnosticReport(deleteResource *r4b.DiagnosticReport) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDiagnosticReport given nil DiagnosticReport")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDiagnosticReport given DiagnosticReport with nil Id (Id required to delete)")
	}
	return fc.DeleteDiagnosticReportById(*deleteResource.Id)
}

// delete DiagnosticReport with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDiagnosticReportById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DiagnosticReport", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteDiagnosticReportById: %s", err)
		}
	}
	return &parsed, err
}

// create DocumentManifest, return id of created DocumentManifest or OperationOutcome error
func (fc *FhirClient) CreateDocumentManifest(createResource *r4b.DocumentManifest) (*r4b.DocumentManifest, error) {
	if createResource == nil {
		return nil, errors.New("CreateDocumentManifest called with nil DocumentManifest")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DocumentManifest")
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

	var parsed r4b.DocumentManifest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateDocumentManifest: %s", err)
		}
	}
	return &parsed, err
}

// read DocumentManifest from fhir server by id, return DocumentManifest or OperationOutcome error
func (fc *FhirClient) ReadDocumentManifest(id string) (*r4b.DocumentManifest, error) {
	if id == "" {
		return nil, errors.New("ReadDocumentManifest given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DocumentManifest", id)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.DocumentManifest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadDocumentManifest: %s", err)
		}
	}
	return &parsed, err
}

// replace DocumentManifest if exists on server, else create new DocumentManifest with given id, return DocumentManifest or OperationOutcome error
func (fc *FhirClient) UpdateDocumentManifest(updateResource *r4b.DocumentManifest) (*r4b.DocumentManifest, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateDocumentManifest called with nil DocumentManifest")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DocumentManifest", *updateResource.Id)
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

	var parsed r4b.DocumentManifest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateDocumentManifest: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return DocumentManifest or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDocumentManifest(patchResource *r4b.DocumentManifest) (*r4b.DocumentManifest, error) {
	if patchResource == nil {
		return nil, errors.New("PatchDocumentManifest given nil DocumentManifest")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchDocumentManifest given DocumentManifest without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DocumentManifest", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchDocumentManifest error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchDocumentManifest error using patch representation as json")
	}
	req, err := http.NewRequest(http.MethodPatch, reqUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json-patch+json")

	var parsed r4b.DocumentManifest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchDocumentManifest: %s", err)
		}
	}
	return &parsed, err
}

// delete DocumentManifest and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDocumentManifest(deleteResource *r4b.DocumentManifest) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDocumentManifest given nil DocumentManifest")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDocumentManifest given DocumentManifest with nil Id (Id required to delete)")
	}
	return fc.DeleteDocumentManifestById(*deleteResource.Id)
}

// delete DocumentManifest with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDocumentManifestById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DocumentManifest", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteDocumentManifestById: %s", err)
		}
	}
	return &parsed, err
}

// create DocumentReference, return id of created DocumentReference or OperationOutcome error
func (fc *FhirClient) CreateDocumentReference(createResource *r4b.DocumentReference) (*r4b.DocumentReference, error) {
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

	var parsed r4b.DocumentReference
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateDocumentReference: %s", err)
		}
	}
	return &parsed, err
}

// read DocumentReference from fhir server by id, return DocumentReference or OperationOutcome error
func (fc *FhirClient) ReadDocumentReference(id string) (*r4b.DocumentReference, error) {
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

	var parsed r4b.DocumentReference
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadDocumentReference: %s", err)
		}
	}
	return &parsed, err
}

// replace DocumentReference if exists on server, else create new DocumentReference with given id, return DocumentReference or OperationOutcome error
func (fc *FhirClient) UpdateDocumentReference(updateResource *r4b.DocumentReference) (*r4b.DocumentReference, error) {
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

	var parsed r4b.DocumentReference
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateDocumentReference: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return DocumentReference or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchDocumentReference(patchResource *r4b.DocumentReference) (*r4b.DocumentReference, error) {
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

	var parsed r4b.DocumentReference
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchDocumentReference: %s", err)
		}
	}
	return &parsed, err
}

// delete DocumentReference and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDocumentReference(deleteResource *r4b.DocumentReference) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteDocumentReference given nil DocumentReference")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteDocumentReference given DocumentReference with nil Id (Id required to delete)")
	}
	return fc.DeleteDocumentReferenceById(*deleteResource.Id)
}

// delete DocumentReference with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteDocumentReferenceById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "DocumentReference", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteDocumentReferenceById: %s", err)
		}
	}
	return &parsed, err
}

// create Encounter, return id of created Encounter or OperationOutcome error
func (fc *FhirClient) CreateEncounter(createResource *r4b.Encounter) (*r4b.Encounter, error) {
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

	var parsed r4b.Encounter
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateEncounter: %s", err)
		}
	}
	return &parsed, err
}

// read Encounter from fhir server by id, return Encounter or OperationOutcome error
func (fc *FhirClient) ReadEncounter(id string) (*r4b.Encounter, error) {
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

	var parsed r4b.Encounter
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadEncounter: %s", err)
		}
	}
	return &parsed, err
}

// replace Encounter if exists on server, else create new Encounter with given id, return Encounter or OperationOutcome error
func (fc *FhirClient) UpdateEncounter(updateResource *r4b.Encounter) (*r4b.Encounter, error) {
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

	var parsed r4b.Encounter
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateEncounter: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Encounter or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEncounter(patchResource *r4b.Encounter) (*r4b.Encounter, error) {
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

	var parsed r4b.Encounter
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchEncounter: %s", err)
		}
	}
	return &parsed, err
}

// delete Encounter and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEncounter(deleteResource *r4b.Encounter) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEncounter given nil Encounter")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEncounter given Encounter with nil Id (Id required to delete)")
	}
	return fc.DeleteEncounterById(*deleteResource.Id)
}

// delete Encounter with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEncounterById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Encounter", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteEncounterById: %s", err)
		}
	}
	return &parsed, err
}

// create Endpoint, return id of created Endpoint or OperationOutcome error
func (fc *FhirClient) CreateEndpoint(createResource *r4b.Endpoint) (*r4b.Endpoint, error) {
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

	var parsed r4b.Endpoint
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateEndpoint: %s", err)
		}
	}
	return &parsed, err
}

// read Endpoint from fhir server by id, return Endpoint or OperationOutcome error
func (fc *FhirClient) ReadEndpoint(id string) (*r4b.Endpoint, error) {
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

	var parsed r4b.Endpoint
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadEndpoint: %s", err)
		}
	}
	return &parsed, err
}

// replace Endpoint if exists on server, else create new Endpoint with given id, return Endpoint or OperationOutcome error
func (fc *FhirClient) UpdateEndpoint(updateResource *r4b.Endpoint) (*r4b.Endpoint, error) {
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

	var parsed r4b.Endpoint
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateEndpoint: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Endpoint or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEndpoint(patchResource *r4b.Endpoint) (*r4b.Endpoint, error) {
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

	var parsed r4b.Endpoint
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchEndpoint: %s", err)
		}
	}
	return &parsed, err
}

// delete Endpoint and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEndpoint(deleteResource *r4b.Endpoint) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEndpoint given nil Endpoint")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEndpoint given Endpoint with nil Id (Id required to delete)")
	}
	return fc.DeleteEndpointById(*deleteResource.Id)
}

// delete Endpoint with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEndpointById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Endpoint", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteEndpointById: %s", err)
		}
	}
	return &parsed, err
}

// create EnrollmentRequest, return id of created EnrollmentRequest or OperationOutcome error
func (fc *FhirClient) CreateEnrollmentRequest(createResource *r4b.EnrollmentRequest) (*r4b.EnrollmentRequest, error) {
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

	var parsed r4b.EnrollmentRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateEnrollmentRequest: %s", err)
		}
	}
	return &parsed, err
}

// read EnrollmentRequest from fhir server by id, return EnrollmentRequest or OperationOutcome error
func (fc *FhirClient) ReadEnrollmentRequest(id string) (*r4b.EnrollmentRequest, error) {
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

	var parsed r4b.EnrollmentRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadEnrollmentRequest: %s", err)
		}
	}
	return &parsed, err
}

// replace EnrollmentRequest if exists on server, else create new EnrollmentRequest with given id, return EnrollmentRequest or OperationOutcome error
func (fc *FhirClient) UpdateEnrollmentRequest(updateResource *r4b.EnrollmentRequest) (*r4b.EnrollmentRequest, error) {
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

	var parsed r4b.EnrollmentRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateEnrollmentRequest: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return EnrollmentRequest or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEnrollmentRequest(patchResource *r4b.EnrollmentRequest) (*r4b.EnrollmentRequest, error) {
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

	var parsed r4b.EnrollmentRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchEnrollmentRequest: %s", err)
		}
	}
	return &parsed, err
}

// delete EnrollmentRequest and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEnrollmentRequest(deleteResource *r4b.EnrollmentRequest) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEnrollmentRequest given nil EnrollmentRequest")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEnrollmentRequest given EnrollmentRequest with nil Id (Id required to delete)")
	}
	return fc.DeleteEnrollmentRequestById(*deleteResource.Id)
}

// delete EnrollmentRequest with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEnrollmentRequestById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EnrollmentRequest", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteEnrollmentRequestById: %s", err)
		}
	}
	return &parsed, err
}

// create EnrollmentResponse, return id of created EnrollmentResponse or OperationOutcome error
func (fc *FhirClient) CreateEnrollmentResponse(createResource *r4b.EnrollmentResponse) (*r4b.EnrollmentResponse, error) {
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

	var parsed r4b.EnrollmentResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateEnrollmentResponse: %s", err)
		}
	}
	return &parsed, err
}

// read EnrollmentResponse from fhir server by id, return EnrollmentResponse or OperationOutcome error
func (fc *FhirClient) ReadEnrollmentResponse(id string) (*r4b.EnrollmentResponse, error) {
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

	var parsed r4b.EnrollmentResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadEnrollmentResponse: %s", err)
		}
	}
	return &parsed, err
}

// replace EnrollmentResponse if exists on server, else create new EnrollmentResponse with given id, return EnrollmentResponse or OperationOutcome error
func (fc *FhirClient) UpdateEnrollmentResponse(updateResource *r4b.EnrollmentResponse) (*r4b.EnrollmentResponse, error) {
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

	var parsed r4b.EnrollmentResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateEnrollmentResponse: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return EnrollmentResponse or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEnrollmentResponse(patchResource *r4b.EnrollmentResponse) (*r4b.EnrollmentResponse, error) {
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

	var parsed r4b.EnrollmentResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchEnrollmentResponse: %s", err)
		}
	}
	return &parsed, err
}

// delete EnrollmentResponse and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEnrollmentResponse(deleteResource *r4b.EnrollmentResponse) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEnrollmentResponse given nil EnrollmentResponse")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEnrollmentResponse given EnrollmentResponse with nil Id (Id required to delete)")
	}
	return fc.DeleteEnrollmentResponseById(*deleteResource.Id)
}

// delete EnrollmentResponse with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEnrollmentResponseById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EnrollmentResponse", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteEnrollmentResponseById: %s", err)
		}
	}
	return &parsed, err
}

// create EpisodeOfCare, return id of created EpisodeOfCare or OperationOutcome error
func (fc *FhirClient) CreateEpisodeOfCare(createResource *r4b.EpisodeOfCare) (*r4b.EpisodeOfCare, error) {
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

	var parsed r4b.EpisodeOfCare
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateEpisodeOfCare: %s", err)
		}
	}
	return &parsed, err
}

// read EpisodeOfCare from fhir server by id, return EpisodeOfCare or OperationOutcome error
func (fc *FhirClient) ReadEpisodeOfCare(id string) (*r4b.EpisodeOfCare, error) {
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

	var parsed r4b.EpisodeOfCare
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadEpisodeOfCare: %s", err)
		}
	}
	return &parsed, err
}

// replace EpisodeOfCare if exists on server, else create new EpisodeOfCare with given id, return EpisodeOfCare or OperationOutcome error
func (fc *FhirClient) UpdateEpisodeOfCare(updateResource *r4b.EpisodeOfCare) (*r4b.EpisodeOfCare, error) {
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

	var parsed r4b.EpisodeOfCare
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateEpisodeOfCare: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return EpisodeOfCare or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEpisodeOfCare(patchResource *r4b.EpisodeOfCare) (*r4b.EpisodeOfCare, error) {
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

	var parsed r4b.EpisodeOfCare
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchEpisodeOfCare: %s", err)
		}
	}
	return &parsed, err
}

// delete EpisodeOfCare and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEpisodeOfCare(deleteResource *r4b.EpisodeOfCare) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEpisodeOfCare given nil EpisodeOfCare")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEpisodeOfCare given EpisodeOfCare with nil Id (Id required to delete)")
	}
	return fc.DeleteEpisodeOfCareById(*deleteResource.Id)
}

// delete EpisodeOfCare with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEpisodeOfCareById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EpisodeOfCare", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteEpisodeOfCareById: %s", err)
		}
	}
	return &parsed, err
}

// create EventDefinition, return id of created EventDefinition or OperationOutcome error
func (fc *FhirClient) CreateEventDefinition(createResource *r4b.EventDefinition) (*r4b.EventDefinition, error) {
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

	var parsed r4b.EventDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateEventDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read EventDefinition from fhir server by id, return EventDefinition or OperationOutcome error
func (fc *FhirClient) ReadEventDefinition(id string) (*r4b.EventDefinition, error) {
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

	var parsed r4b.EventDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadEventDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace EventDefinition if exists on server, else create new EventDefinition with given id, return EventDefinition or OperationOutcome error
func (fc *FhirClient) UpdateEventDefinition(updateResource *r4b.EventDefinition) (*r4b.EventDefinition, error) {
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

	var parsed r4b.EventDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateEventDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return EventDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEventDefinition(patchResource *r4b.EventDefinition) (*r4b.EventDefinition, error) {
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

	var parsed r4b.EventDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchEventDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete EventDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEventDefinition(deleteResource *r4b.EventDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEventDefinition given nil EventDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEventDefinition given EventDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteEventDefinitionById(*deleteResource.Id)
}

// delete EventDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEventDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EventDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteEventDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create Evidence, return id of created Evidence or OperationOutcome error
func (fc *FhirClient) CreateEvidence(createResource *r4b.Evidence) (*r4b.Evidence, error) {
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

	var parsed r4b.Evidence
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateEvidence: %s", err)
		}
	}
	return &parsed, err
}

// read Evidence from fhir server by id, return Evidence or OperationOutcome error
func (fc *FhirClient) ReadEvidence(id string) (*r4b.Evidence, error) {
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

	var parsed r4b.Evidence
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadEvidence: %s", err)
		}
	}
	return &parsed, err
}

// replace Evidence if exists on server, else create new Evidence with given id, return Evidence or OperationOutcome error
func (fc *FhirClient) UpdateEvidence(updateResource *r4b.Evidence) (*r4b.Evidence, error) {
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

	var parsed r4b.Evidence
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateEvidence: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Evidence or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEvidence(patchResource *r4b.Evidence) (*r4b.Evidence, error) {
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

	var parsed r4b.Evidence
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchEvidence: %s", err)
		}
	}
	return &parsed, err
}

// delete Evidence and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEvidence(deleteResource *r4b.Evidence) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEvidence given nil Evidence")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEvidence given Evidence with nil Id (Id required to delete)")
	}
	return fc.DeleteEvidenceById(*deleteResource.Id)
}

// delete Evidence with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEvidenceById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Evidence", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteEvidenceById: %s", err)
		}
	}
	return &parsed, err
}

// create EvidenceReport, return id of created EvidenceReport or OperationOutcome error
func (fc *FhirClient) CreateEvidenceReport(createResource *r4b.EvidenceReport) (*r4b.EvidenceReport, error) {
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

	var parsed r4b.EvidenceReport
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateEvidenceReport: %s", err)
		}
	}
	return &parsed, err
}

// read EvidenceReport from fhir server by id, return EvidenceReport or OperationOutcome error
func (fc *FhirClient) ReadEvidenceReport(id string) (*r4b.EvidenceReport, error) {
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

	var parsed r4b.EvidenceReport
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadEvidenceReport: %s", err)
		}
	}
	return &parsed, err
}

// replace EvidenceReport if exists on server, else create new EvidenceReport with given id, return EvidenceReport or OperationOutcome error
func (fc *FhirClient) UpdateEvidenceReport(updateResource *r4b.EvidenceReport) (*r4b.EvidenceReport, error) {
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

	var parsed r4b.EvidenceReport
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateEvidenceReport: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return EvidenceReport or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEvidenceReport(patchResource *r4b.EvidenceReport) (*r4b.EvidenceReport, error) {
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

	var parsed r4b.EvidenceReport
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchEvidenceReport: %s", err)
		}
	}
	return &parsed, err
}

// delete EvidenceReport and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEvidenceReport(deleteResource *r4b.EvidenceReport) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEvidenceReport given nil EvidenceReport")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEvidenceReport given EvidenceReport with nil Id (Id required to delete)")
	}
	return fc.DeleteEvidenceReportById(*deleteResource.Id)
}

// delete EvidenceReport with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEvidenceReportById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EvidenceReport", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteEvidenceReportById: %s", err)
		}
	}
	return &parsed, err
}

// create EvidenceVariable, return id of created EvidenceVariable or OperationOutcome error
func (fc *FhirClient) CreateEvidenceVariable(createResource *r4b.EvidenceVariable) (*r4b.EvidenceVariable, error) {
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

	var parsed r4b.EvidenceVariable
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateEvidenceVariable: %s", err)
		}
	}
	return &parsed, err
}

// read EvidenceVariable from fhir server by id, return EvidenceVariable or OperationOutcome error
func (fc *FhirClient) ReadEvidenceVariable(id string) (*r4b.EvidenceVariable, error) {
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

	var parsed r4b.EvidenceVariable
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadEvidenceVariable: %s", err)
		}
	}
	return &parsed, err
}

// replace EvidenceVariable if exists on server, else create new EvidenceVariable with given id, return EvidenceVariable or OperationOutcome error
func (fc *FhirClient) UpdateEvidenceVariable(updateResource *r4b.EvidenceVariable) (*r4b.EvidenceVariable, error) {
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

	var parsed r4b.EvidenceVariable
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateEvidenceVariable: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return EvidenceVariable or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchEvidenceVariable(patchResource *r4b.EvidenceVariable) (*r4b.EvidenceVariable, error) {
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

	var parsed r4b.EvidenceVariable
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchEvidenceVariable: %s", err)
		}
	}
	return &parsed, err
}

// delete EvidenceVariable and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEvidenceVariable(deleteResource *r4b.EvidenceVariable) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteEvidenceVariable given nil EvidenceVariable")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteEvidenceVariable given EvidenceVariable with nil Id (Id required to delete)")
	}
	return fc.DeleteEvidenceVariableById(*deleteResource.Id)
}

// delete EvidenceVariable with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteEvidenceVariableById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "EvidenceVariable", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteEvidenceVariableById: %s", err)
		}
	}
	return &parsed, err
}

// create ExampleScenario, return id of created ExampleScenario or OperationOutcome error
func (fc *FhirClient) CreateExampleScenario(createResource *r4b.ExampleScenario) (*r4b.ExampleScenario, error) {
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

	var parsed r4b.ExampleScenario
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateExampleScenario: %s", err)
		}
	}
	return &parsed, err
}

// read ExampleScenario from fhir server by id, return ExampleScenario or OperationOutcome error
func (fc *FhirClient) ReadExampleScenario(id string) (*r4b.ExampleScenario, error) {
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

	var parsed r4b.ExampleScenario
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadExampleScenario: %s", err)
		}
	}
	return &parsed, err
}

// replace ExampleScenario if exists on server, else create new ExampleScenario with given id, return ExampleScenario or OperationOutcome error
func (fc *FhirClient) UpdateExampleScenario(updateResource *r4b.ExampleScenario) (*r4b.ExampleScenario, error) {
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

	var parsed r4b.ExampleScenario
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateExampleScenario: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ExampleScenario or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchExampleScenario(patchResource *r4b.ExampleScenario) (*r4b.ExampleScenario, error) {
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

	var parsed r4b.ExampleScenario
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchExampleScenario: %s", err)
		}
	}
	return &parsed, err
}

// delete ExampleScenario and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteExampleScenario(deleteResource *r4b.ExampleScenario) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteExampleScenario given nil ExampleScenario")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteExampleScenario given ExampleScenario with nil Id (Id required to delete)")
	}
	return fc.DeleteExampleScenarioById(*deleteResource.Id)
}

// delete ExampleScenario with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteExampleScenarioById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ExampleScenario", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteExampleScenarioById: %s", err)
		}
	}
	return &parsed, err
}

// create ExplanationOfBenefit, return id of created ExplanationOfBenefit or OperationOutcome error
func (fc *FhirClient) CreateExplanationOfBenefit(createResource *r4b.ExplanationOfBenefit) (*r4b.ExplanationOfBenefit, error) {
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

	var parsed r4b.ExplanationOfBenefit
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateExplanationOfBenefit: %s", err)
		}
	}
	return &parsed, err
}

// read ExplanationOfBenefit from fhir server by id, return ExplanationOfBenefit or OperationOutcome error
func (fc *FhirClient) ReadExplanationOfBenefit(id string) (*r4b.ExplanationOfBenefit, error) {
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

	var parsed r4b.ExplanationOfBenefit
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadExplanationOfBenefit: %s", err)
		}
	}
	return &parsed, err
}

// replace ExplanationOfBenefit if exists on server, else create new ExplanationOfBenefit with given id, return ExplanationOfBenefit or OperationOutcome error
func (fc *FhirClient) UpdateExplanationOfBenefit(updateResource *r4b.ExplanationOfBenefit) (*r4b.ExplanationOfBenefit, error) {
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

	var parsed r4b.ExplanationOfBenefit
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateExplanationOfBenefit: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ExplanationOfBenefit or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchExplanationOfBenefit(patchResource *r4b.ExplanationOfBenefit) (*r4b.ExplanationOfBenefit, error) {
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

	var parsed r4b.ExplanationOfBenefit
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchExplanationOfBenefit: %s", err)
		}
	}
	return &parsed, err
}

// delete ExplanationOfBenefit and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteExplanationOfBenefit(deleteResource *r4b.ExplanationOfBenefit) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteExplanationOfBenefit given nil ExplanationOfBenefit")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteExplanationOfBenefit given ExplanationOfBenefit with nil Id (Id required to delete)")
	}
	return fc.DeleteExplanationOfBenefitById(*deleteResource.Id)
}

// delete ExplanationOfBenefit with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteExplanationOfBenefitById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ExplanationOfBenefit", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteExplanationOfBenefitById: %s", err)
		}
	}
	return &parsed, err
}

// create FamilyMemberHistory, return id of created FamilyMemberHistory or OperationOutcome error
func (fc *FhirClient) CreateFamilyMemberHistory(createResource *r4b.FamilyMemberHistory) (*r4b.FamilyMemberHistory, error) {
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

	var parsed r4b.FamilyMemberHistory
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateFamilyMemberHistory: %s", err)
		}
	}
	return &parsed, err
}

// read FamilyMemberHistory from fhir server by id, return FamilyMemberHistory or OperationOutcome error
func (fc *FhirClient) ReadFamilyMemberHistory(id string) (*r4b.FamilyMemberHistory, error) {
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

	var parsed r4b.FamilyMemberHistory
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadFamilyMemberHistory: %s", err)
		}
	}
	return &parsed, err
}

// replace FamilyMemberHistory if exists on server, else create new FamilyMemberHistory with given id, return FamilyMemberHistory or OperationOutcome error
func (fc *FhirClient) UpdateFamilyMemberHistory(updateResource *r4b.FamilyMemberHistory) (*r4b.FamilyMemberHistory, error) {
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

	var parsed r4b.FamilyMemberHistory
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateFamilyMemberHistory: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return FamilyMemberHistory or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchFamilyMemberHistory(patchResource *r4b.FamilyMemberHistory) (*r4b.FamilyMemberHistory, error) {
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

	var parsed r4b.FamilyMemberHistory
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchFamilyMemberHistory: %s", err)
		}
	}
	return &parsed, err
}

// delete FamilyMemberHistory and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteFamilyMemberHistory(deleteResource *r4b.FamilyMemberHistory) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteFamilyMemberHistory given nil FamilyMemberHistory")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteFamilyMemberHistory given FamilyMemberHistory with nil Id (Id required to delete)")
	}
	return fc.DeleteFamilyMemberHistoryById(*deleteResource.Id)
}

// delete FamilyMemberHistory with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteFamilyMemberHistoryById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "FamilyMemberHistory", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteFamilyMemberHistoryById: %s", err)
		}
	}
	return &parsed, err
}

// create Flag, return id of created Flag or OperationOutcome error
func (fc *FhirClient) CreateFlag(createResource *r4b.Flag) (*r4b.Flag, error) {
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

	var parsed r4b.Flag
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateFlag: %s", err)
		}
	}
	return &parsed, err
}

// read Flag from fhir server by id, return Flag or OperationOutcome error
func (fc *FhirClient) ReadFlag(id string) (*r4b.Flag, error) {
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

	var parsed r4b.Flag
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadFlag: %s", err)
		}
	}
	return &parsed, err
}

// replace Flag if exists on server, else create new Flag with given id, return Flag or OperationOutcome error
func (fc *FhirClient) UpdateFlag(updateResource *r4b.Flag) (*r4b.Flag, error) {
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

	var parsed r4b.Flag
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateFlag: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Flag or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchFlag(patchResource *r4b.Flag) (*r4b.Flag, error) {
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

	var parsed r4b.Flag
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchFlag: %s", err)
		}
	}
	return &parsed, err
}

// delete Flag and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteFlag(deleteResource *r4b.Flag) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteFlag given nil Flag")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteFlag given Flag with nil Id (Id required to delete)")
	}
	return fc.DeleteFlagById(*deleteResource.Id)
}

// delete Flag with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteFlagById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Flag", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteFlagById: %s", err)
		}
	}
	return &parsed, err
}

// create Goal, return id of created Goal or OperationOutcome error
func (fc *FhirClient) CreateGoal(createResource *r4b.Goal) (*r4b.Goal, error) {
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

	var parsed r4b.Goal
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateGoal: %s", err)
		}
	}
	return &parsed, err
}

// read Goal from fhir server by id, return Goal or OperationOutcome error
func (fc *FhirClient) ReadGoal(id string) (*r4b.Goal, error) {
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

	var parsed r4b.Goal
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadGoal: %s", err)
		}
	}
	return &parsed, err
}

// replace Goal if exists on server, else create new Goal with given id, return Goal or OperationOutcome error
func (fc *FhirClient) UpdateGoal(updateResource *r4b.Goal) (*r4b.Goal, error) {
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

	var parsed r4b.Goal
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateGoal: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Goal or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchGoal(patchResource *r4b.Goal) (*r4b.Goal, error) {
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

	var parsed r4b.Goal
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchGoal: %s", err)
		}
	}
	return &parsed, err
}

// delete Goal and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGoal(deleteResource *r4b.Goal) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteGoal given nil Goal")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteGoal given Goal with nil Id (Id required to delete)")
	}
	return fc.DeleteGoalById(*deleteResource.Id)
}

// delete Goal with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGoalById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Goal", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteGoalById: %s", err)
		}
	}
	return &parsed, err
}

// create GraphDefinition, return id of created GraphDefinition or OperationOutcome error
func (fc *FhirClient) CreateGraphDefinition(createResource *r4b.GraphDefinition) (*r4b.GraphDefinition, error) {
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

	var parsed r4b.GraphDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateGraphDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read GraphDefinition from fhir server by id, return GraphDefinition or OperationOutcome error
func (fc *FhirClient) ReadGraphDefinition(id string) (*r4b.GraphDefinition, error) {
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

	var parsed r4b.GraphDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadGraphDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace GraphDefinition if exists on server, else create new GraphDefinition with given id, return GraphDefinition or OperationOutcome error
func (fc *FhirClient) UpdateGraphDefinition(updateResource *r4b.GraphDefinition) (*r4b.GraphDefinition, error) {
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

	var parsed r4b.GraphDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateGraphDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return GraphDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchGraphDefinition(patchResource *r4b.GraphDefinition) (*r4b.GraphDefinition, error) {
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

	var parsed r4b.GraphDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchGraphDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete GraphDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGraphDefinition(deleteResource *r4b.GraphDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteGraphDefinition given nil GraphDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteGraphDefinition given GraphDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteGraphDefinitionById(*deleteResource.Id)
}

// delete GraphDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGraphDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GraphDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteGraphDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create Group, return id of created Group or OperationOutcome error
func (fc *FhirClient) CreateGroup(createResource *r4b.Group) (*r4b.Group, error) {
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

	var parsed r4b.Group
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateGroup: %s", err)
		}
	}
	return &parsed, err
}

// read Group from fhir server by id, return Group or OperationOutcome error
func (fc *FhirClient) ReadGroup(id string) (*r4b.Group, error) {
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

	var parsed r4b.Group
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadGroup: %s", err)
		}
	}
	return &parsed, err
}

// replace Group if exists on server, else create new Group with given id, return Group or OperationOutcome error
func (fc *FhirClient) UpdateGroup(updateResource *r4b.Group) (*r4b.Group, error) {
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

	var parsed r4b.Group
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateGroup: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Group or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchGroup(patchResource *r4b.Group) (*r4b.Group, error) {
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

	var parsed r4b.Group
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchGroup: %s", err)
		}
	}
	return &parsed, err
}

// delete Group and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGroup(deleteResource *r4b.Group) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteGroup given nil Group")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteGroup given Group with nil Id (Id required to delete)")
	}
	return fc.DeleteGroupById(*deleteResource.Id)
}

// delete Group with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGroupById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Group", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteGroupById: %s", err)
		}
	}
	return &parsed, err
}

// create GuidanceResponse, return id of created GuidanceResponse or OperationOutcome error
func (fc *FhirClient) CreateGuidanceResponse(createResource *r4b.GuidanceResponse) (*r4b.GuidanceResponse, error) {
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

	var parsed r4b.GuidanceResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateGuidanceResponse: %s", err)
		}
	}
	return &parsed, err
}

// read GuidanceResponse from fhir server by id, return GuidanceResponse or OperationOutcome error
func (fc *FhirClient) ReadGuidanceResponse(id string) (*r4b.GuidanceResponse, error) {
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

	var parsed r4b.GuidanceResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadGuidanceResponse: %s", err)
		}
	}
	return &parsed, err
}

// replace GuidanceResponse if exists on server, else create new GuidanceResponse with given id, return GuidanceResponse or OperationOutcome error
func (fc *FhirClient) UpdateGuidanceResponse(updateResource *r4b.GuidanceResponse) (*r4b.GuidanceResponse, error) {
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

	var parsed r4b.GuidanceResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateGuidanceResponse: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return GuidanceResponse or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchGuidanceResponse(patchResource *r4b.GuidanceResponse) (*r4b.GuidanceResponse, error) {
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

	var parsed r4b.GuidanceResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchGuidanceResponse: %s", err)
		}
	}
	return &parsed, err
}

// delete GuidanceResponse and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGuidanceResponse(deleteResource *r4b.GuidanceResponse) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteGuidanceResponse given nil GuidanceResponse")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteGuidanceResponse given GuidanceResponse with nil Id (Id required to delete)")
	}
	return fc.DeleteGuidanceResponseById(*deleteResource.Id)
}

// delete GuidanceResponse with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteGuidanceResponseById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "GuidanceResponse", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteGuidanceResponseById: %s", err)
		}
	}
	return &parsed, err
}

// create HealthcareService, return id of created HealthcareService or OperationOutcome error
func (fc *FhirClient) CreateHealthcareService(createResource *r4b.HealthcareService) (*r4b.HealthcareService, error) {
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

	var parsed r4b.HealthcareService
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateHealthcareService: %s", err)
		}
	}
	return &parsed, err
}

// read HealthcareService from fhir server by id, return HealthcareService or OperationOutcome error
func (fc *FhirClient) ReadHealthcareService(id string) (*r4b.HealthcareService, error) {
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

	var parsed r4b.HealthcareService
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadHealthcareService: %s", err)
		}
	}
	return &parsed, err
}

// replace HealthcareService if exists on server, else create new HealthcareService with given id, return HealthcareService or OperationOutcome error
func (fc *FhirClient) UpdateHealthcareService(updateResource *r4b.HealthcareService) (*r4b.HealthcareService, error) {
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

	var parsed r4b.HealthcareService
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateHealthcareService: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return HealthcareService or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchHealthcareService(patchResource *r4b.HealthcareService) (*r4b.HealthcareService, error) {
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

	var parsed r4b.HealthcareService
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchHealthcareService: %s", err)
		}
	}
	return &parsed, err
}

// delete HealthcareService and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteHealthcareService(deleteResource *r4b.HealthcareService) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteHealthcareService given nil HealthcareService")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteHealthcareService given HealthcareService with nil Id (Id required to delete)")
	}
	return fc.DeleteHealthcareServiceById(*deleteResource.Id)
}

// delete HealthcareService with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteHealthcareServiceById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "HealthcareService", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteHealthcareServiceById: %s", err)
		}
	}
	return &parsed, err
}

// create ImagingStudy, return id of created ImagingStudy or OperationOutcome error
func (fc *FhirClient) CreateImagingStudy(createResource *r4b.ImagingStudy) (*r4b.ImagingStudy, error) {
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

	var parsed r4b.ImagingStudy
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateImagingStudy: %s", err)
		}
	}
	return &parsed, err
}

// read ImagingStudy from fhir server by id, return ImagingStudy or OperationOutcome error
func (fc *FhirClient) ReadImagingStudy(id string) (*r4b.ImagingStudy, error) {
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

	var parsed r4b.ImagingStudy
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadImagingStudy: %s", err)
		}
	}
	return &parsed, err
}

// replace ImagingStudy if exists on server, else create new ImagingStudy with given id, return ImagingStudy or OperationOutcome error
func (fc *FhirClient) UpdateImagingStudy(updateResource *r4b.ImagingStudy) (*r4b.ImagingStudy, error) {
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

	var parsed r4b.ImagingStudy
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateImagingStudy: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ImagingStudy or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchImagingStudy(patchResource *r4b.ImagingStudy) (*r4b.ImagingStudy, error) {
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

	var parsed r4b.ImagingStudy
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchImagingStudy: %s", err)
		}
	}
	return &parsed, err
}

// delete ImagingStudy and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImagingStudy(deleteResource *r4b.ImagingStudy) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteImagingStudy given nil ImagingStudy")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteImagingStudy given ImagingStudy with nil Id (Id required to delete)")
	}
	return fc.DeleteImagingStudyById(*deleteResource.Id)
}

// delete ImagingStudy with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImagingStudyById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImagingStudy", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteImagingStudyById: %s", err)
		}
	}
	return &parsed, err
}

// create Immunization, return id of created Immunization or OperationOutcome error
func (fc *FhirClient) CreateImmunization(createResource *r4b.Immunization) (*r4b.Immunization, error) {
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

	var parsed r4b.Immunization
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateImmunization: %s", err)
		}
	}
	return &parsed, err
}

// read Immunization from fhir server by id, return Immunization or OperationOutcome error
func (fc *FhirClient) ReadImmunization(id string) (*r4b.Immunization, error) {
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

	var parsed r4b.Immunization
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadImmunization: %s", err)
		}
	}
	return &parsed, err
}

// replace Immunization if exists on server, else create new Immunization with given id, return Immunization or OperationOutcome error
func (fc *FhirClient) UpdateImmunization(updateResource *r4b.Immunization) (*r4b.Immunization, error) {
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

	var parsed r4b.Immunization
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateImmunization: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Immunization or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchImmunization(patchResource *r4b.Immunization) (*r4b.Immunization, error) {
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

	var parsed r4b.Immunization
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchImmunization: %s", err)
		}
	}
	return &parsed, err
}

// delete Immunization and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImmunization(deleteResource *r4b.Immunization) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteImmunization given nil Immunization")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteImmunization given Immunization with nil Id (Id required to delete)")
	}
	return fc.DeleteImmunizationById(*deleteResource.Id)
}

// delete Immunization with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImmunizationById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Immunization", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteImmunizationById: %s", err)
		}
	}
	return &parsed, err
}

// create ImmunizationEvaluation, return id of created ImmunizationEvaluation or OperationOutcome error
func (fc *FhirClient) CreateImmunizationEvaluation(createResource *r4b.ImmunizationEvaluation) (*r4b.ImmunizationEvaluation, error) {
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

	var parsed r4b.ImmunizationEvaluation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateImmunizationEvaluation: %s", err)
		}
	}
	return &parsed, err
}

// read ImmunizationEvaluation from fhir server by id, return ImmunizationEvaluation or OperationOutcome error
func (fc *FhirClient) ReadImmunizationEvaluation(id string) (*r4b.ImmunizationEvaluation, error) {
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

	var parsed r4b.ImmunizationEvaluation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadImmunizationEvaluation: %s", err)
		}
	}
	return &parsed, err
}

// replace ImmunizationEvaluation if exists on server, else create new ImmunizationEvaluation with given id, return ImmunizationEvaluation or OperationOutcome error
func (fc *FhirClient) UpdateImmunizationEvaluation(updateResource *r4b.ImmunizationEvaluation) (*r4b.ImmunizationEvaluation, error) {
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

	var parsed r4b.ImmunizationEvaluation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateImmunizationEvaluation: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ImmunizationEvaluation or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchImmunizationEvaluation(patchResource *r4b.ImmunizationEvaluation) (*r4b.ImmunizationEvaluation, error) {
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

	var parsed r4b.ImmunizationEvaluation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchImmunizationEvaluation: %s", err)
		}
	}
	return &parsed, err
}

// delete ImmunizationEvaluation and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImmunizationEvaluation(deleteResource *r4b.ImmunizationEvaluation) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteImmunizationEvaluation given nil ImmunizationEvaluation")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteImmunizationEvaluation given ImmunizationEvaluation with nil Id (Id required to delete)")
	}
	return fc.DeleteImmunizationEvaluationById(*deleteResource.Id)
}

// delete ImmunizationEvaluation with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImmunizationEvaluationById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImmunizationEvaluation", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteImmunizationEvaluationById: %s", err)
		}
	}
	return &parsed, err
}

// create ImmunizationRecommendation, return id of created ImmunizationRecommendation or OperationOutcome error
func (fc *FhirClient) CreateImmunizationRecommendation(createResource *r4b.ImmunizationRecommendation) (*r4b.ImmunizationRecommendation, error) {
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

	var parsed r4b.ImmunizationRecommendation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateImmunizationRecommendation: %s", err)
		}
	}
	return &parsed, err
}

// read ImmunizationRecommendation from fhir server by id, return ImmunizationRecommendation or OperationOutcome error
func (fc *FhirClient) ReadImmunizationRecommendation(id string) (*r4b.ImmunizationRecommendation, error) {
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

	var parsed r4b.ImmunizationRecommendation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadImmunizationRecommendation: %s", err)
		}
	}
	return &parsed, err
}

// replace ImmunizationRecommendation if exists on server, else create new ImmunizationRecommendation with given id, return ImmunizationRecommendation or OperationOutcome error
func (fc *FhirClient) UpdateImmunizationRecommendation(updateResource *r4b.ImmunizationRecommendation) (*r4b.ImmunizationRecommendation, error) {
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

	var parsed r4b.ImmunizationRecommendation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateImmunizationRecommendation: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ImmunizationRecommendation or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchImmunizationRecommendation(patchResource *r4b.ImmunizationRecommendation) (*r4b.ImmunizationRecommendation, error) {
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

	var parsed r4b.ImmunizationRecommendation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchImmunizationRecommendation: %s", err)
		}
	}
	return &parsed, err
}

// delete ImmunizationRecommendation and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImmunizationRecommendation(deleteResource *r4b.ImmunizationRecommendation) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteImmunizationRecommendation given nil ImmunizationRecommendation")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteImmunizationRecommendation given ImmunizationRecommendation with nil Id (Id required to delete)")
	}
	return fc.DeleteImmunizationRecommendationById(*deleteResource.Id)
}

// delete ImmunizationRecommendation with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImmunizationRecommendationById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImmunizationRecommendation", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteImmunizationRecommendationById: %s", err)
		}
	}
	return &parsed, err
}

// create ImplementationGuide, return id of created ImplementationGuide or OperationOutcome error
func (fc *FhirClient) CreateImplementationGuide(createResource *r4b.ImplementationGuide) (*r4b.ImplementationGuide, error) {
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

	var parsed r4b.ImplementationGuide
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateImplementationGuide: %s", err)
		}
	}
	return &parsed, err
}

// read ImplementationGuide from fhir server by id, return ImplementationGuide or OperationOutcome error
func (fc *FhirClient) ReadImplementationGuide(id string) (*r4b.ImplementationGuide, error) {
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

	var parsed r4b.ImplementationGuide
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadImplementationGuide: %s", err)
		}
	}
	return &parsed, err
}

// replace ImplementationGuide if exists on server, else create new ImplementationGuide with given id, return ImplementationGuide or OperationOutcome error
func (fc *FhirClient) UpdateImplementationGuide(updateResource *r4b.ImplementationGuide) (*r4b.ImplementationGuide, error) {
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

	var parsed r4b.ImplementationGuide
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateImplementationGuide: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ImplementationGuide or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchImplementationGuide(patchResource *r4b.ImplementationGuide) (*r4b.ImplementationGuide, error) {
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

	var parsed r4b.ImplementationGuide
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchImplementationGuide: %s", err)
		}
	}
	return &parsed, err
}

// delete ImplementationGuide and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImplementationGuide(deleteResource *r4b.ImplementationGuide) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteImplementationGuide given nil ImplementationGuide")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteImplementationGuide given ImplementationGuide with nil Id (Id required to delete)")
	}
	return fc.DeleteImplementationGuideById(*deleteResource.Id)
}

// delete ImplementationGuide with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteImplementationGuideById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ImplementationGuide", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteImplementationGuideById: %s", err)
		}
	}
	return &parsed, err
}

// create Ingredient, return id of created Ingredient or OperationOutcome error
func (fc *FhirClient) CreateIngredient(createResource *r4b.Ingredient) (*r4b.Ingredient, error) {
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

	var parsed r4b.Ingredient
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateIngredient: %s", err)
		}
	}
	return &parsed, err
}

// read Ingredient from fhir server by id, return Ingredient or OperationOutcome error
func (fc *FhirClient) ReadIngredient(id string) (*r4b.Ingredient, error) {
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

	var parsed r4b.Ingredient
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadIngredient: %s", err)
		}
	}
	return &parsed, err
}

// replace Ingredient if exists on server, else create new Ingredient with given id, return Ingredient or OperationOutcome error
func (fc *FhirClient) UpdateIngredient(updateResource *r4b.Ingredient) (*r4b.Ingredient, error) {
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

	var parsed r4b.Ingredient
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateIngredient: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Ingredient or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchIngredient(patchResource *r4b.Ingredient) (*r4b.Ingredient, error) {
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

	var parsed r4b.Ingredient
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchIngredient: %s", err)
		}
	}
	return &parsed, err
}

// delete Ingredient and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteIngredient(deleteResource *r4b.Ingredient) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteIngredient given nil Ingredient")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteIngredient given Ingredient with nil Id (Id required to delete)")
	}
	return fc.DeleteIngredientById(*deleteResource.Id)
}

// delete Ingredient with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteIngredientById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Ingredient", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteIngredientById: %s", err)
		}
	}
	return &parsed, err
}

// create InsurancePlan, return id of created InsurancePlan or OperationOutcome error
func (fc *FhirClient) CreateInsurancePlan(createResource *r4b.InsurancePlan) (*r4b.InsurancePlan, error) {
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

	var parsed r4b.InsurancePlan
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateInsurancePlan: %s", err)
		}
	}
	return &parsed, err
}

// read InsurancePlan from fhir server by id, return InsurancePlan or OperationOutcome error
func (fc *FhirClient) ReadInsurancePlan(id string) (*r4b.InsurancePlan, error) {
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

	var parsed r4b.InsurancePlan
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadInsurancePlan: %s", err)
		}
	}
	return &parsed, err
}

// replace InsurancePlan if exists on server, else create new InsurancePlan with given id, return InsurancePlan or OperationOutcome error
func (fc *FhirClient) UpdateInsurancePlan(updateResource *r4b.InsurancePlan) (*r4b.InsurancePlan, error) {
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

	var parsed r4b.InsurancePlan
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateInsurancePlan: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return InsurancePlan or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchInsurancePlan(patchResource *r4b.InsurancePlan) (*r4b.InsurancePlan, error) {
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

	var parsed r4b.InsurancePlan
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchInsurancePlan: %s", err)
		}
	}
	return &parsed, err
}

// delete InsurancePlan and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteInsurancePlan(deleteResource *r4b.InsurancePlan) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteInsurancePlan given nil InsurancePlan")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteInsurancePlan given InsurancePlan with nil Id (Id required to delete)")
	}
	return fc.DeleteInsurancePlanById(*deleteResource.Id)
}

// delete InsurancePlan with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteInsurancePlanById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "InsurancePlan", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteInsurancePlanById: %s", err)
		}
	}
	return &parsed, err
}

// create Invoice, return id of created Invoice or OperationOutcome error
func (fc *FhirClient) CreateInvoice(createResource *r4b.Invoice) (*r4b.Invoice, error) {
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

	var parsed r4b.Invoice
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateInvoice: %s", err)
		}
	}
	return &parsed, err
}

// read Invoice from fhir server by id, return Invoice or OperationOutcome error
func (fc *FhirClient) ReadInvoice(id string) (*r4b.Invoice, error) {
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

	var parsed r4b.Invoice
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadInvoice: %s", err)
		}
	}
	return &parsed, err
}

// replace Invoice if exists on server, else create new Invoice with given id, return Invoice or OperationOutcome error
func (fc *FhirClient) UpdateInvoice(updateResource *r4b.Invoice) (*r4b.Invoice, error) {
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

	var parsed r4b.Invoice
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateInvoice: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Invoice or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchInvoice(patchResource *r4b.Invoice) (*r4b.Invoice, error) {
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

	var parsed r4b.Invoice
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchInvoice: %s", err)
		}
	}
	return &parsed, err
}

// delete Invoice and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteInvoice(deleteResource *r4b.Invoice) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteInvoice given nil Invoice")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteInvoice given Invoice with nil Id (Id required to delete)")
	}
	return fc.DeleteInvoiceById(*deleteResource.Id)
}

// delete Invoice with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteInvoiceById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Invoice", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteInvoiceById: %s", err)
		}
	}
	return &parsed, err
}

// create Library, return id of created Library or OperationOutcome error
func (fc *FhirClient) CreateLibrary(createResource *r4b.Library) (*r4b.Library, error) {
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

	var parsed r4b.Library
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateLibrary: %s", err)
		}
	}
	return &parsed, err
}

// read Library from fhir server by id, return Library or OperationOutcome error
func (fc *FhirClient) ReadLibrary(id string) (*r4b.Library, error) {
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

	var parsed r4b.Library
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadLibrary: %s", err)
		}
	}
	return &parsed, err
}

// replace Library if exists on server, else create new Library with given id, return Library or OperationOutcome error
func (fc *FhirClient) UpdateLibrary(updateResource *r4b.Library) (*r4b.Library, error) {
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

	var parsed r4b.Library
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateLibrary: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Library or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchLibrary(patchResource *r4b.Library) (*r4b.Library, error) {
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

	var parsed r4b.Library
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchLibrary: %s", err)
		}
	}
	return &parsed, err
}

// delete Library and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteLibrary(deleteResource *r4b.Library) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteLibrary given nil Library")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteLibrary given Library with nil Id (Id required to delete)")
	}
	return fc.DeleteLibraryById(*deleteResource.Id)
}

// delete Library with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteLibraryById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Library", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteLibraryById: %s", err)
		}
	}
	return &parsed, err
}

// create Linkage, return id of created Linkage or OperationOutcome error
func (fc *FhirClient) CreateLinkage(createResource *r4b.Linkage) (*r4b.Linkage, error) {
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

	var parsed r4b.Linkage
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateLinkage: %s", err)
		}
	}
	return &parsed, err
}

// read Linkage from fhir server by id, return Linkage or OperationOutcome error
func (fc *FhirClient) ReadLinkage(id string) (*r4b.Linkage, error) {
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

	var parsed r4b.Linkage
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadLinkage: %s", err)
		}
	}
	return &parsed, err
}

// replace Linkage if exists on server, else create new Linkage with given id, return Linkage or OperationOutcome error
func (fc *FhirClient) UpdateLinkage(updateResource *r4b.Linkage) (*r4b.Linkage, error) {
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

	var parsed r4b.Linkage
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateLinkage: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Linkage or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchLinkage(patchResource *r4b.Linkage) (*r4b.Linkage, error) {
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

	var parsed r4b.Linkage
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchLinkage: %s", err)
		}
	}
	return &parsed, err
}

// delete Linkage and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteLinkage(deleteResource *r4b.Linkage) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteLinkage given nil Linkage")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteLinkage given Linkage with nil Id (Id required to delete)")
	}
	return fc.DeleteLinkageById(*deleteResource.Id)
}

// delete Linkage with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteLinkageById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Linkage", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteLinkageById: %s", err)
		}
	}
	return &parsed, err
}

// create List, return id of created List or OperationOutcome error
func (fc *FhirClient) CreateList(createResource *r4b.List) (*r4b.List, error) {
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

	var parsed r4b.List
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateList: %s", err)
		}
	}
	return &parsed, err
}

// read List from fhir server by id, return List or OperationOutcome error
func (fc *FhirClient) ReadList(id string) (*r4b.List, error) {
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

	var parsed r4b.List
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadList: %s", err)
		}
	}
	return &parsed, err
}

// replace List if exists on server, else create new List with given id, return List or OperationOutcome error
func (fc *FhirClient) UpdateList(updateResource *r4b.List) (*r4b.List, error) {
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

	var parsed r4b.List
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateList: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return List or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchList(patchResource *r4b.List) (*r4b.List, error) {
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

	var parsed r4b.List
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchList: %s", err)
		}
	}
	return &parsed, err
}

// delete List and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteList(deleteResource *r4b.List) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteList given nil List")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteList given List with nil Id (Id required to delete)")
	}
	return fc.DeleteListById(*deleteResource.Id)
}

// delete List with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteListById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "List", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteListById: %s", err)
		}
	}
	return &parsed, err
}

// create Location, return id of created Location or OperationOutcome error
func (fc *FhirClient) CreateLocation(createResource *r4b.Location) (*r4b.Location, error) {
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

	var parsed r4b.Location
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateLocation: %s", err)
		}
	}
	return &parsed, err
}

// read Location from fhir server by id, return Location or OperationOutcome error
func (fc *FhirClient) ReadLocation(id string) (*r4b.Location, error) {
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

	var parsed r4b.Location
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadLocation: %s", err)
		}
	}
	return &parsed, err
}

// replace Location if exists on server, else create new Location with given id, return Location or OperationOutcome error
func (fc *FhirClient) UpdateLocation(updateResource *r4b.Location) (*r4b.Location, error) {
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

	var parsed r4b.Location
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateLocation: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Location or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchLocation(patchResource *r4b.Location) (*r4b.Location, error) {
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

	var parsed r4b.Location
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchLocation: %s", err)
		}
	}
	return &parsed, err
}

// delete Location and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteLocation(deleteResource *r4b.Location) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteLocation given nil Location")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteLocation given Location with nil Id (Id required to delete)")
	}
	return fc.DeleteLocationById(*deleteResource.Id)
}

// delete Location with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteLocationById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Location", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteLocationById: %s", err)
		}
	}
	return &parsed, err
}

// create ManufacturedItemDefinition, return id of created ManufacturedItemDefinition or OperationOutcome error
func (fc *FhirClient) CreateManufacturedItemDefinition(createResource *r4b.ManufacturedItemDefinition) (*r4b.ManufacturedItemDefinition, error) {
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

	var parsed r4b.ManufacturedItemDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateManufacturedItemDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read ManufacturedItemDefinition from fhir server by id, return ManufacturedItemDefinition or OperationOutcome error
func (fc *FhirClient) ReadManufacturedItemDefinition(id string) (*r4b.ManufacturedItemDefinition, error) {
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

	var parsed r4b.ManufacturedItemDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadManufacturedItemDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace ManufacturedItemDefinition if exists on server, else create new ManufacturedItemDefinition with given id, return ManufacturedItemDefinition or OperationOutcome error
func (fc *FhirClient) UpdateManufacturedItemDefinition(updateResource *r4b.ManufacturedItemDefinition) (*r4b.ManufacturedItemDefinition, error) {
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

	var parsed r4b.ManufacturedItemDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateManufacturedItemDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ManufacturedItemDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchManufacturedItemDefinition(patchResource *r4b.ManufacturedItemDefinition) (*r4b.ManufacturedItemDefinition, error) {
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

	var parsed r4b.ManufacturedItemDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchManufacturedItemDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete ManufacturedItemDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteManufacturedItemDefinition(deleteResource *r4b.ManufacturedItemDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteManufacturedItemDefinition given nil ManufacturedItemDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteManufacturedItemDefinition given ManufacturedItemDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteManufacturedItemDefinitionById(*deleteResource.Id)
}

// delete ManufacturedItemDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteManufacturedItemDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ManufacturedItemDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteManufacturedItemDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create Measure, return id of created Measure or OperationOutcome error
func (fc *FhirClient) CreateMeasure(createResource *r4b.Measure) (*r4b.Measure, error) {
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

	var parsed r4b.Measure
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateMeasure: %s", err)
		}
	}
	return &parsed, err
}

// read Measure from fhir server by id, return Measure or OperationOutcome error
func (fc *FhirClient) ReadMeasure(id string) (*r4b.Measure, error) {
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

	var parsed r4b.Measure
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadMeasure: %s", err)
		}
	}
	return &parsed, err
}

// replace Measure if exists on server, else create new Measure with given id, return Measure or OperationOutcome error
func (fc *FhirClient) UpdateMeasure(updateResource *r4b.Measure) (*r4b.Measure, error) {
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

	var parsed r4b.Measure
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateMeasure: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Measure or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMeasure(patchResource *r4b.Measure) (*r4b.Measure, error) {
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

	var parsed r4b.Measure
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchMeasure: %s", err)
		}
	}
	return &parsed, err
}

// delete Measure and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMeasure(deleteResource *r4b.Measure) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMeasure given nil Measure")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMeasure given Measure with nil Id (Id required to delete)")
	}
	return fc.DeleteMeasureById(*deleteResource.Id)
}

// delete Measure with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMeasureById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Measure", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteMeasureById: %s", err)
		}
	}
	return &parsed, err
}

// create MeasureReport, return id of created MeasureReport or OperationOutcome error
func (fc *FhirClient) CreateMeasureReport(createResource *r4b.MeasureReport) (*r4b.MeasureReport, error) {
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

	var parsed r4b.MeasureReport
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateMeasureReport: %s", err)
		}
	}
	return &parsed, err
}

// read MeasureReport from fhir server by id, return MeasureReport or OperationOutcome error
func (fc *FhirClient) ReadMeasureReport(id string) (*r4b.MeasureReport, error) {
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

	var parsed r4b.MeasureReport
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadMeasureReport: %s", err)
		}
	}
	return &parsed, err
}

// replace MeasureReport if exists on server, else create new MeasureReport with given id, return MeasureReport or OperationOutcome error
func (fc *FhirClient) UpdateMeasureReport(updateResource *r4b.MeasureReport) (*r4b.MeasureReport, error) {
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

	var parsed r4b.MeasureReport
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateMeasureReport: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return MeasureReport or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMeasureReport(patchResource *r4b.MeasureReport) (*r4b.MeasureReport, error) {
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

	var parsed r4b.MeasureReport
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchMeasureReport: %s", err)
		}
	}
	return &parsed, err
}

// delete MeasureReport and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMeasureReport(deleteResource *r4b.MeasureReport) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMeasureReport given nil MeasureReport")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMeasureReport given MeasureReport with nil Id (Id required to delete)")
	}
	return fc.DeleteMeasureReportById(*deleteResource.Id)
}

// delete MeasureReport with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMeasureReportById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MeasureReport", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteMeasureReportById: %s", err)
		}
	}
	return &parsed, err
}

// create Media, return id of created Media or OperationOutcome error
func (fc *FhirClient) CreateMedia(createResource *r4b.Media) (*r4b.Media, error) {
	if createResource == nil {
		return nil, errors.New("CreateMedia called with nil Media")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Media")
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

	var parsed r4b.Media
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateMedia: %s", err)
		}
	}
	return &parsed, err
}

// read Media from fhir server by id, return Media or OperationOutcome error
func (fc *FhirClient) ReadMedia(id string) (*r4b.Media, error) {
	if id == "" {
		return nil, errors.New("ReadMedia given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Media", id)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.Media
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadMedia: %s", err)
		}
	}
	return &parsed, err
}

// replace Media if exists on server, else create new Media with given id, return Media or OperationOutcome error
func (fc *FhirClient) UpdateMedia(updateResource *r4b.Media) (*r4b.Media, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateMedia called with nil Media")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Media", *updateResource.Id)
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

	var parsed r4b.Media
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateMedia: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Media or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMedia(patchResource *r4b.Media) (*r4b.Media, error) {
	if patchResource == nil {
		return nil, errors.New("PatchMedia given nil Media")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchMedia given Media without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Media", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchMedia error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchMedia error using patch representation as json")
	}
	req, err := http.NewRequest(http.MethodPatch, reqUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json-patch+json")

	var parsed r4b.Media
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchMedia: %s", err)
		}
	}
	return &parsed, err
}

// delete Media and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedia(deleteResource *r4b.Media) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMedia given nil Media")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMedia given Media with nil Id (Id required to delete)")
	}
	return fc.DeleteMediaById(*deleteResource.Id)
}

// delete Media with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMediaById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Media", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteMediaById: %s", err)
		}
	}
	return &parsed, err
}

// create Medication, return id of created Medication or OperationOutcome error
func (fc *FhirClient) CreateMedication(createResource *r4b.Medication) (*r4b.Medication, error) {
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

	var parsed r4b.Medication
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateMedication: %s", err)
		}
	}
	return &parsed, err
}

// read Medication from fhir server by id, return Medication or OperationOutcome error
func (fc *FhirClient) ReadMedication(id string) (*r4b.Medication, error) {
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

	var parsed r4b.Medication
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadMedication: %s", err)
		}
	}
	return &parsed, err
}

// replace Medication if exists on server, else create new Medication with given id, return Medication or OperationOutcome error
func (fc *FhirClient) UpdateMedication(updateResource *r4b.Medication) (*r4b.Medication, error) {
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

	var parsed r4b.Medication
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateMedication: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Medication or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMedication(patchResource *r4b.Medication) (*r4b.Medication, error) {
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

	var parsed r4b.Medication
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchMedication: %s", err)
		}
	}
	return &parsed, err
}

// delete Medication and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedication(deleteResource *r4b.Medication) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMedication given nil Medication")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMedication given Medication with nil Id (Id required to delete)")
	}
	return fc.DeleteMedicationById(*deleteResource.Id)
}

// delete Medication with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Medication", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteMedicationById: %s", err)
		}
	}
	return &parsed, err
}

// create MedicationAdministration, return id of created MedicationAdministration or OperationOutcome error
func (fc *FhirClient) CreateMedicationAdministration(createResource *r4b.MedicationAdministration) (*r4b.MedicationAdministration, error) {
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

	var parsed r4b.MedicationAdministration
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateMedicationAdministration: %s", err)
		}
	}
	return &parsed, err
}

// read MedicationAdministration from fhir server by id, return MedicationAdministration or OperationOutcome error
func (fc *FhirClient) ReadMedicationAdministration(id string) (*r4b.MedicationAdministration, error) {
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

	var parsed r4b.MedicationAdministration
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadMedicationAdministration: %s", err)
		}
	}
	return &parsed, err
}

// replace MedicationAdministration if exists on server, else create new MedicationAdministration with given id, return MedicationAdministration or OperationOutcome error
func (fc *FhirClient) UpdateMedicationAdministration(updateResource *r4b.MedicationAdministration) (*r4b.MedicationAdministration, error) {
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

	var parsed r4b.MedicationAdministration
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateMedicationAdministration: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return MedicationAdministration or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMedicationAdministration(patchResource *r4b.MedicationAdministration) (*r4b.MedicationAdministration, error) {
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

	var parsed r4b.MedicationAdministration
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchMedicationAdministration: %s", err)
		}
	}
	return &parsed, err
}

// delete MedicationAdministration and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationAdministration(deleteResource *r4b.MedicationAdministration) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMedicationAdministration given nil MedicationAdministration")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMedicationAdministration given MedicationAdministration with nil Id (Id required to delete)")
	}
	return fc.DeleteMedicationAdministrationById(*deleteResource.Id)
}

// delete MedicationAdministration with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationAdministrationById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationAdministration", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteMedicationAdministrationById: %s", err)
		}
	}
	return &parsed, err
}

// create MedicationDispense, return id of created MedicationDispense or OperationOutcome error
func (fc *FhirClient) CreateMedicationDispense(createResource *r4b.MedicationDispense) (*r4b.MedicationDispense, error) {
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

	var parsed r4b.MedicationDispense
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateMedicationDispense: %s", err)
		}
	}
	return &parsed, err
}

// read MedicationDispense from fhir server by id, return MedicationDispense or OperationOutcome error
func (fc *FhirClient) ReadMedicationDispense(id string) (*r4b.MedicationDispense, error) {
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

	var parsed r4b.MedicationDispense
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadMedicationDispense: %s", err)
		}
	}
	return &parsed, err
}

// replace MedicationDispense if exists on server, else create new MedicationDispense with given id, return MedicationDispense or OperationOutcome error
func (fc *FhirClient) UpdateMedicationDispense(updateResource *r4b.MedicationDispense) (*r4b.MedicationDispense, error) {
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

	var parsed r4b.MedicationDispense
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateMedicationDispense: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return MedicationDispense or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMedicationDispense(patchResource *r4b.MedicationDispense) (*r4b.MedicationDispense, error) {
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

	var parsed r4b.MedicationDispense
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchMedicationDispense: %s", err)
		}
	}
	return &parsed, err
}

// delete MedicationDispense and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationDispense(deleteResource *r4b.MedicationDispense) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMedicationDispense given nil MedicationDispense")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMedicationDispense given MedicationDispense with nil Id (Id required to delete)")
	}
	return fc.DeleteMedicationDispenseById(*deleteResource.Id)
}

// delete MedicationDispense with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationDispenseById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationDispense", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteMedicationDispenseById: %s", err)
		}
	}
	return &parsed, err
}

// create MedicationKnowledge, return id of created MedicationKnowledge or OperationOutcome error
func (fc *FhirClient) CreateMedicationKnowledge(createResource *r4b.MedicationKnowledge) (*r4b.MedicationKnowledge, error) {
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

	var parsed r4b.MedicationKnowledge
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateMedicationKnowledge: %s", err)
		}
	}
	return &parsed, err
}

// read MedicationKnowledge from fhir server by id, return MedicationKnowledge or OperationOutcome error
func (fc *FhirClient) ReadMedicationKnowledge(id string) (*r4b.MedicationKnowledge, error) {
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

	var parsed r4b.MedicationKnowledge
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadMedicationKnowledge: %s", err)
		}
	}
	return &parsed, err
}

// replace MedicationKnowledge if exists on server, else create new MedicationKnowledge with given id, return MedicationKnowledge or OperationOutcome error
func (fc *FhirClient) UpdateMedicationKnowledge(updateResource *r4b.MedicationKnowledge) (*r4b.MedicationKnowledge, error) {
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

	var parsed r4b.MedicationKnowledge
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateMedicationKnowledge: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return MedicationKnowledge or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMedicationKnowledge(patchResource *r4b.MedicationKnowledge) (*r4b.MedicationKnowledge, error) {
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

	var parsed r4b.MedicationKnowledge
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchMedicationKnowledge: %s", err)
		}
	}
	return &parsed, err
}

// delete MedicationKnowledge and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationKnowledge(deleteResource *r4b.MedicationKnowledge) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMedicationKnowledge given nil MedicationKnowledge")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMedicationKnowledge given MedicationKnowledge with nil Id (Id required to delete)")
	}
	return fc.DeleteMedicationKnowledgeById(*deleteResource.Id)
}

// delete MedicationKnowledge with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationKnowledgeById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationKnowledge", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteMedicationKnowledgeById: %s", err)
		}
	}
	return &parsed, err
}

// create MedicationRequest, return id of created MedicationRequest or OperationOutcome error
func (fc *FhirClient) CreateMedicationRequest(createResource *r4b.MedicationRequest) (*r4b.MedicationRequest, error) {
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

	var parsed r4b.MedicationRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateMedicationRequest: %s", err)
		}
	}
	return &parsed, err
}

// read MedicationRequest from fhir server by id, return MedicationRequest or OperationOutcome error
func (fc *FhirClient) ReadMedicationRequest(id string) (*r4b.MedicationRequest, error) {
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

	var parsed r4b.MedicationRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadMedicationRequest: %s", err)
		}
	}
	return &parsed, err
}

// replace MedicationRequest if exists on server, else create new MedicationRequest with given id, return MedicationRequest or OperationOutcome error
func (fc *FhirClient) UpdateMedicationRequest(updateResource *r4b.MedicationRequest) (*r4b.MedicationRequest, error) {
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

	var parsed r4b.MedicationRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateMedicationRequest: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return MedicationRequest or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMedicationRequest(patchResource *r4b.MedicationRequest) (*r4b.MedicationRequest, error) {
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

	var parsed r4b.MedicationRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchMedicationRequest: %s", err)
		}
	}
	return &parsed, err
}

// delete MedicationRequest and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationRequest(deleteResource *r4b.MedicationRequest) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMedicationRequest given nil MedicationRequest")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMedicationRequest given MedicationRequest with nil Id (Id required to delete)")
	}
	return fc.DeleteMedicationRequestById(*deleteResource.Id)
}

// delete MedicationRequest with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationRequestById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationRequest", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteMedicationRequestById: %s", err)
		}
	}
	return &parsed, err
}

// create MedicationStatement, return id of created MedicationStatement or OperationOutcome error
func (fc *FhirClient) CreateMedicationStatement(createResource *r4b.MedicationStatement) (*r4b.MedicationStatement, error) {
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

	var parsed r4b.MedicationStatement
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateMedicationStatement: %s", err)
		}
	}
	return &parsed, err
}

// read MedicationStatement from fhir server by id, return MedicationStatement or OperationOutcome error
func (fc *FhirClient) ReadMedicationStatement(id string) (*r4b.MedicationStatement, error) {
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

	var parsed r4b.MedicationStatement
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadMedicationStatement: %s", err)
		}
	}
	return &parsed, err
}

// replace MedicationStatement if exists on server, else create new MedicationStatement with given id, return MedicationStatement or OperationOutcome error
func (fc *FhirClient) UpdateMedicationStatement(updateResource *r4b.MedicationStatement) (*r4b.MedicationStatement, error) {
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

	var parsed r4b.MedicationStatement
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateMedicationStatement: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return MedicationStatement or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMedicationStatement(patchResource *r4b.MedicationStatement) (*r4b.MedicationStatement, error) {
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

	var parsed r4b.MedicationStatement
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchMedicationStatement: %s", err)
		}
	}
	return &parsed, err
}

// delete MedicationStatement and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationStatement(deleteResource *r4b.MedicationStatement) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMedicationStatement given nil MedicationStatement")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMedicationStatement given MedicationStatement with nil Id (Id required to delete)")
	}
	return fc.DeleteMedicationStatementById(*deleteResource.Id)
}

// delete MedicationStatement with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicationStatementById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicationStatement", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteMedicationStatementById: %s", err)
		}
	}
	return &parsed, err
}

// create MedicinalProductDefinition, return id of created MedicinalProductDefinition or OperationOutcome error
func (fc *FhirClient) CreateMedicinalProductDefinition(createResource *r4b.MedicinalProductDefinition) (*r4b.MedicinalProductDefinition, error) {
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

	var parsed r4b.MedicinalProductDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateMedicinalProductDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read MedicinalProductDefinition from fhir server by id, return MedicinalProductDefinition or OperationOutcome error
func (fc *FhirClient) ReadMedicinalProductDefinition(id string) (*r4b.MedicinalProductDefinition, error) {
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

	var parsed r4b.MedicinalProductDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadMedicinalProductDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace MedicinalProductDefinition if exists on server, else create new MedicinalProductDefinition with given id, return MedicinalProductDefinition or OperationOutcome error
func (fc *FhirClient) UpdateMedicinalProductDefinition(updateResource *r4b.MedicinalProductDefinition) (*r4b.MedicinalProductDefinition, error) {
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

	var parsed r4b.MedicinalProductDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateMedicinalProductDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return MedicinalProductDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMedicinalProductDefinition(patchResource *r4b.MedicinalProductDefinition) (*r4b.MedicinalProductDefinition, error) {
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

	var parsed r4b.MedicinalProductDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchMedicinalProductDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete MedicinalProductDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicinalProductDefinition(deleteResource *r4b.MedicinalProductDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMedicinalProductDefinition given nil MedicinalProductDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMedicinalProductDefinition given MedicinalProductDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteMedicinalProductDefinitionById(*deleteResource.Id)
}

// delete MedicinalProductDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMedicinalProductDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MedicinalProductDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteMedicinalProductDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create MessageDefinition, return id of created MessageDefinition or OperationOutcome error
func (fc *FhirClient) CreateMessageDefinition(createResource *r4b.MessageDefinition) (*r4b.MessageDefinition, error) {
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

	var parsed r4b.MessageDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateMessageDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read MessageDefinition from fhir server by id, return MessageDefinition or OperationOutcome error
func (fc *FhirClient) ReadMessageDefinition(id string) (*r4b.MessageDefinition, error) {
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

	var parsed r4b.MessageDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadMessageDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace MessageDefinition if exists on server, else create new MessageDefinition with given id, return MessageDefinition or OperationOutcome error
func (fc *FhirClient) UpdateMessageDefinition(updateResource *r4b.MessageDefinition) (*r4b.MessageDefinition, error) {
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

	var parsed r4b.MessageDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateMessageDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return MessageDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMessageDefinition(patchResource *r4b.MessageDefinition) (*r4b.MessageDefinition, error) {
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

	var parsed r4b.MessageDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchMessageDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete MessageDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMessageDefinition(deleteResource *r4b.MessageDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMessageDefinition given nil MessageDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMessageDefinition given MessageDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteMessageDefinitionById(*deleteResource.Id)
}

// delete MessageDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMessageDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MessageDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteMessageDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create MessageHeader, return id of created MessageHeader or OperationOutcome error
func (fc *FhirClient) CreateMessageHeader(createResource *r4b.MessageHeader) (*r4b.MessageHeader, error) {
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

	var parsed r4b.MessageHeader
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateMessageHeader: %s", err)
		}
	}
	return &parsed, err
}

// read MessageHeader from fhir server by id, return MessageHeader or OperationOutcome error
func (fc *FhirClient) ReadMessageHeader(id string) (*r4b.MessageHeader, error) {
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

	var parsed r4b.MessageHeader
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadMessageHeader: %s", err)
		}
	}
	return &parsed, err
}

// replace MessageHeader if exists on server, else create new MessageHeader with given id, return MessageHeader or OperationOutcome error
func (fc *FhirClient) UpdateMessageHeader(updateResource *r4b.MessageHeader) (*r4b.MessageHeader, error) {
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

	var parsed r4b.MessageHeader
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateMessageHeader: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return MessageHeader or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMessageHeader(patchResource *r4b.MessageHeader) (*r4b.MessageHeader, error) {
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

	var parsed r4b.MessageHeader
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchMessageHeader: %s", err)
		}
	}
	return &parsed, err
}

// delete MessageHeader and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMessageHeader(deleteResource *r4b.MessageHeader) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMessageHeader given nil MessageHeader")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMessageHeader given MessageHeader with nil Id (Id required to delete)")
	}
	return fc.DeleteMessageHeaderById(*deleteResource.Id)
}

// delete MessageHeader with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMessageHeaderById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MessageHeader", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteMessageHeaderById: %s", err)
		}
	}
	return &parsed, err
}

// create MolecularSequence, return id of created MolecularSequence or OperationOutcome error
func (fc *FhirClient) CreateMolecularSequence(createResource *r4b.MolecularSequence) (*r4b.MolecularSequence, error) {
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

	var parsed r4b.MolecularSequence
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateMolecularSequence: %s", err)
		}
	}
	return &parsed, err
}

// read MolecularSequence from fhir server by id, return MolecularSequence or OperationOutcome error
func (fc *FhirClient) ReadMolecularSequence(id string) (*r4b.MolecularSequence, error) {
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

	var parsed r4b.MolecularSequence
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadMolecularSequence: %s", err)
		}
	}
	return &parsed, err
}

// replace MolecularSequence if exists on server, else create new MolecularSequence with given id, return MolecularSequence or OperationOutcome error
func (fc *FhirClient) UpdateMolecularSequence(updateResource *r4b.MolecularSequence) (*r4b.MolecularSequence, error) {
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

	var parsed r4b.MolecularSequence
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateMolecularSequence: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return MolecularSequence or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchMolecularSequence(patchResource *r4b.MolecularSequence) (*r4b.MolecularSequence, error) {
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

	var parsed r4b.MolecularSequence
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchMolecularSequence: %s", err)
		}
	}
	return &parsed, err
}

// delete MolecularSequence and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMolecularSequence(deleteResource *r4b.MolecularSequence) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteMolecularSequence given nil MolecularSequence")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteMolecularSequence given MolecularSequence with nil Id (Id required to delete)")
	}
	return fc.DeleteMolecularSequenceById(*deleteResource.Id)
}

// delete MolecularSequence with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteMolecularSequenceById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "MolecularSequence", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteMolecularSequenceById: %s", err)
		}
	}
	return &parsed, err
}

// create NamingSystem, return id of created NamingSystem or OperationOutcome error
func (fc *FhirClient) CreateNamingSystem(createResource *r4b.NamingSystem) (*r4b.NamingSystem, error) {
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

	var parsed r4b.NamingSystem
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateNamingSystem: %s", err)
		}
	}
	return &parsed, err
}

// read NamingSystem from fhir server by id, return NamingSystem or OperationOutcome error
func (fc *FhirClient) ReadNamingSystem(id string) (*r4b.NamingSystem, error) {
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

	var parsed r4b.NamingSystem
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadNamingSystem: %s", err)
		}
	}
	return &parsed, err
}

// replace NamingSystem if exists on server, else create new NamingSystem with given id, return NamingSystem or OperationOutcome error
func (fc *FhirClient) UpdateNamingSystem(updateResource *r4b.NamingSystem) (*r4b.NamingSystem, error) {
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

	var parsed r4b.NamingSystem
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateNamingSystem: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return NamingSystem or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchNamingSystem(patchResource *r4b.NamingSystem) (*r4b.NamingSystem, error) {
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

	var parsed r4b.NamingSystem
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchNamingSystem: %s", err)
		}
	}
	return &parsed, err
}

// delete NamingSystem and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteNamingSystem(deleteResource *r4b.NamingSystem) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteNamingSystem given nil NamingSystem")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteNamingSystem given NamingSystem with nil Id (Id required to delete)")
	}
	return fc.DeleteNamingSystemById(*deleteResource.Id)
}

// delete NamingSystem with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteNamingSystemById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NamingSystem", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteNamingSystemById: %s", err)
		}
	}
	return &parsed, err
}

// create NutritionOrder, return id of created NutritionOrder or OperationOutcome error
func (fc *FhirClient) CreateNutritionOrder(createResource *r4b.NutritionOrder) (*r4b.NutritionOrder, error) {
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

	var parsed r4b.NutritionOrder
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateNutritionOrder: %s", err)
		}
	}
	return &parsed, err
}

// read NutritionOrder from fhir server by id, return NutritionOrder or OperationOutcome error
func (fc *FhirClient) ReadNutritionOrder(id string) (*r4b.NutritionOrder, error) {
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

	var parsed r4b.NutritionOrder
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadNutritionOrder: %s", err)
		}
	}
	return &parsed, err
}

// replace NutritionOrder if exists on server, else create new NutritionOrder with given id, return NutritionOrder or OperationOutcome error
func (fc *FhirClient) UpdateNutritionOrder(updateResource *r4b.NutritionOrder) (*r4b.NutritionOrder, error) {
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

	var parsed r4b.NutritionOrder
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateNutritionOrder: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return NutritionOrder or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchNutritionOrder(patchResource *r4b.NutritionOrder) (*r4b.NutritionOrder, error) {
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

	var parsed r4b.NutritionOrder
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchNutritionOrder: %s", err)
		}
	}
	return &parsed, err
}

// delete NutritionOrder and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteNutritionOrder(deleteResource *r4b.NutritionOrder) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteNutritionOrder given nil NutritionOrder")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteNutritionOrder given NutritionOrder with nil Id (Id required to delete)")
	}
	return fc.DeleteNutritionOrderById(*deleteResource.Id)
}

// delete NutritionOrder with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteNutritionOrderById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionOrder", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteNutritionOrderById: %s", err)
		}
	}
	return &parsed, err
}

// create NutritionProduct, return id of created NutritionProduct or OperationOutcome error
func (fc *FhirClient) CreateNutritionProduct(createResource *r4b.NutritionProduct) (*r4b.NutritionProduct, error) {
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

	var parsed r4b.NutritionProduct
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateNutritionProduct: %s", err)
		}
	}
	return &parsed, err
}

// read NutritionProduct from fhir server by id, return NutritionProduct or OperationOutcome error
func (fc *FhirClient) ReadNutritionProduct(id string) (*r4b.NutritionProduct, error) {
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

	var parsed r4b.NutritionProduct
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadNutritionProduct: %s", err)
		}
	}
	return &parsed, err
}

// replace NutritionProduct if exists on server, else create new NutritionProduct with given id, return NutritionProduct or OperationOutcome error
func (fc *FhirClient) UpdateNutritionProduct(updateResource *r4b.NutritionProduct) (*r4b.NutritionProduct, error) {
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

	var parsed r4b.NutritionProduct
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateNutritionProduct: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return NutritionProduct or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchNutritionProduct(patchResource *r4b.NutritionProduct) (*r4b.NutritionProduct, error) {
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

	var parsed r4b.NutritionProduct
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchNutritionProduct: %s", err)
		}
	}
	return &parsed, err
}

// delete NutritionProduct and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteNutritionProduct(deleteResource *r4b.NutritionProduct) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteNutritionProduct given nil NutritionProduct")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteNutritionProduct given NutritionProduct with nil Id (Id required to delete)")
	}
	return fc.DeleteNutritionProductById(*deleteResource.Id)
}

// delete NutritionProduct with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteNutritionProductById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "NutritionProduct", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteNutritionProductById: %s", err)
		}
	}
	return &parsed, err
}

// create Observation, return id of created Observation or OperationOutcome error
func (fc *FhirClient) CreateObservation(createResource *r4b.Observation) (*r4b.Observation, error) {
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

	var parsed r4b.Observation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateObservation: %s", err)
		}
	}
	return &parsed, err
}

// read Observation from fhir server by id, return Observation or OperationOutcome error
func (fc *FhirClient) ReadObservation(id string) (*r4b.Observation, error) {
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

	var parsed r4b.Observation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadObservation: %s", err)
		}
	}
	return &parsed, err
}

// replace Observation if exists on server, else create new Observation with given id, return Observation or OperationOutcome error
func (fc *FhirClient) UpdateObservation(updateResource *r4b.Observation) (*r4b.Observation, error) {
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

	var parsed r4b.Observation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateObservation: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Observation or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchObservation(patchResource *r4b.Observation) (*r4b.Observation, error) {
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

	var parsed r4b.Observation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchObservation: %s", err)
		}
	}
	return &parsed, err
}

// delete Observation and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteObservation(deleteResource *r4b.Observation) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteObservation given nil Observation")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteObservation given Observation with nil Id (Id required to delete)")
	}
	return fc.DeleteObservationById(*deleteResource.Id)
}

// delete Observation with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteObservationById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Observation", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteObservationById: %s", err)
		}
	}
	return &parsed, err
}

// create ObservationDefinition, return id of created ObservationDefinition or OperationOutcome error
func (fc *FhirClient) CreateObservationDefinition(createResource *r4b.ObservationDefinition) (*r4b.ObservationDefinition, error) {
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

	var parsed r4b.ObservationDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateObservationDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read ObservationDefinition from fhir server by id, return ObservationDefinition or OperationOutcome error
func (fc *FhirClient) ReadObservationDefinition(id string) (*r4b.ObservationDefinition, error) {
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

	var parsed r4b.ObservationDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadObservationDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace ObservationDefinition if exists on server, else create new ObservationDefinition with given id, return ObservationDefinition or OperationOutcome error
func (fc *FhirClient) UpdateObservationDefinition(updateResource *r4b.ObservationDefinition) (*r4b.ObservationDefinition, error) {
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

	var parsed r4b.ObservationDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateObservationDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ObservationDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchObservationDefinition(patchResource *r4b.ObservationDefinition) (*r4b.ObservationDefinition, error) {
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

	var parsed r4b.ObservationDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchObservationDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete ObservationDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteObservationDefinition(deleteResource *r4b.ObservationDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteObservationDefinition given nil ObservationDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteObservationDefinition given ObservationDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteObservationDefinitionById(*deleteResource.Id)
}

// delete ObservationDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteObservationDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ObservationDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteObservationDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create OperationDefinition, return id of created OperationDefinition or OperationOutcome error
func (fc *FhirClient) CreateOperationDefinition(createResource *r4b.OperationDefinition) (*r4b.OperationDefinition, error) {
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

	var parsed r4b.OperationDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateOperationDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read OperationDefinition from fhir server by id, return OperationDefinition or OperationOutcome error
func (fc *FhirClient) ReadOperationDefinition(id string) (*r4b.OperationDefinition, error) {
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

	var parsed r4b.OperationDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadOperationDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace OperationDefinition if exists on server, else create new OperationDefinition with given id, return OperationDefinition or OperationOutcome error
func (fc *FhirClient) UpdateOperationDefinition(updateResource *r4b.OperationDefinition) (*r4b.OperationDefinition, error) {
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

	var parsed r4b.OperationDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateOperationDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return OperationDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchOperationDefinition(patchResource *r4b.OperationDefinition) (*r4b.OperationDefinition, error) {
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

	var parsed r4b.OperationDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchOperationDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete OperationDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteOperationDefinition(deleteResource *r4b.OperationDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteOperationDefinition given nil OperationDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteOperationDefinition given OperationDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteOperationDefinitionById(*deleteResource.Id)
}

// delete OperationDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteOperationDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OperationDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteOperationDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create OperationOutcome, return id of created OperationOutcome or OperationOutcome error
func (fc *FhirClient) CreateOperationOutcome(createResource *r4b.OperationOutcome) (*r4b.OperationOutcome, error) {
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

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateOperationOutcome: %s", err)
		}
	}
	return &parsed, err
}

// read OperationOutcome from fhir server by id, return OperationOutcome or OperationOutcome error
func (fc *FhirClient) ReadOperationOutcome(id string) (*r4b.OperationOutcome, error) {
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

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadOperationOutcome: %s", err)
		}
	}
	return &parsed, err
}

// replace OperationOutcome if exists on server, else create new OperationOutcome with given id, return OperationOutcome or OperationOutcome error
func (fc *FhirClient) UpdateOperationOutcome(updateResource *r4b.OperationOutcome) (*r4b.OperationOutcome, error) {
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

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateOperationOutcome: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return OperationOutcome or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchOperationOutcome(patchResource *r4b.OperationOutcome) (*r4b.OperationOutcome, error) {
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

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchOperationOutcome: %s", err)
		}
	}
	return &parsed, err
}

// delete OperationOutcome and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteOperationOutcome(deleteResource *r4b.OperationOutcome) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteOperationOutcome given nil OperationOutcome")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteOperationOutcome given OperationOutcome with nil Id (Id required to delete)")
	}
	return fc.DeleteOperationOutcomeById(*deleteResource.Id)
}

// delete OperationOutcome with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteOperationOutcomeById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OperationOutcome", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteOperationOutcomeById: %s", err)
		}
	}
	return &parsed, err
}

// create Organization, return id of created Organization or OperationOutcome error
func (fc *FhirClient) CreateOrganization(createResource *r4b.Organization) (*r4b.Organization, error) {
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

	var parsed r4b.Organization
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateOrganization: %s", err)
		}
	}
	return &parsed, err
}

// read Organization from fhir server by id, return Organization or OperationOutcome error
func (fc *FhirClient) ReadOrganization(id string) (*r4b.Organization, error) {
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

	var parsed r4b.Organization
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadOrganization: %s", err)
		}
	}
	return &parsed, err
}

// replace Organization if exists on server, else create new Organization with given id, return Organization or OperationOutcome error
func (fc *FhirClient) UpdateOrganization(updateResource *r4b.Organization) (*r4b.Organization, error) {
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

	var parsed r4b.Organization
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateOrganization: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Organization or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchOrganization(patchResource *r4b.Organization) (*r4b.Organization, error) {
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

	var parsed r4b.Organization
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchOrganization: %s", err)
		}
	}
	return &parsed, err
}

// delete Organization and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteOrganization(deleteResource *r4b.Organization) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteOrganization given nil Organization")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteOrganization given Organization with nil Id (Id required to delete)")
	}
	return fc.DeleteOrganizationById(*deleteResource.Id)
}

// delete Organization with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteOrganizationById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Organization", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteOrganizationById: %s", err)
		}
	}
	return &parsed, err
}

// create OrganizationAffiliation, return id of created OrganizationAffiliation or OperationOutcome error
func (fc *FhirClient) CreateOrganizationAffiliation(createResource *r4b.OrganizationAffiliation) (*r4b.OrganizationAffiliation, error) {
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

	var parsed r4b.OrganizationAffiliation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateOrganizationAffiliation: %s", err)
		}
	}
	return &parsed, err
}

// read OrganizationAffiliation from fhir server by id, return OrganizationAffiliation or OperationOutcome error
func (fc *FhirClient) ReadOrganizationAffiliation(id string) (*r4b.OrganizationAffiliation, error) {
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

	var parsed r4b.OrganizationAffiliation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadOrganizationAffiliation: %s", err)
		}
	}
	return &parsed, err
}

// replace OrganizationAffiliation if exists on server, else create new OrganizationAffiliation with given id, return OrganizationAffiliation or OperationOutcome error
func (fc *FhirClient) UpdateOrganizationAffiliation(updateResource *r4b.OrganizationAffiliation) (*r4b.OrganizationAffiliation, error) {
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

	var parsed r4b.OrganizationAffiliation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateOrganizationAffiliation: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return OrganizationAffiliation or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchOrganizationAffiliation(patchResource *r4b.OrganizationAffiliation) (*r4b.OrganizationAffiliation, error) {
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

	var parsed r4b.OrganizationAffiliation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchOrganizationAffiliation: %s", err)
		}
	}
	return &parsed, err
}

// delete OrganizationAffiliation and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteOrganizationAffiliation(deleteResource *r4b.OrganizationAffiliation) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteOrganizationAffiliation given nil OrganizationAffiliation")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteOrganizationAffiliation given OrganizationAffiliation with nil Id (Id required to delete)")
	}
	return fc.DeleteOrganizationAffiliationById(*deleteResource.Id)
}

// delete OrganizationAffiliation with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteOrganizationAffiliationById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "OrganizationAffiliation", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteOrganizationAffiliationById: %s", err)
		}
	}
	return &parsed, err
}

// create PackagedProductDefinition, return id of created PackagedProductDefinition or OperationOutcome error
func (fc *FhirClient) CreatePackagedProductDefinition(createResource *r4b.PackagedProductDefinition) (*r4b.PackagedProductDefinition, error) {
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

	var parsed r4b.PackagedProductDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreatePackagedProductDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read PackagedProductDefinition from fhir server by id, return PackagedProductDefinition or OperationOutcome error
func (fc *FhirClient) ReadPackagedProductDefinition(id string) (*r4b.PackagedProductDefinition, error) {
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

	var parsed r4b.PackagedProductDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadPackagedProductDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace PackagedProductDefinition if exists on server, else create new PackagedProductDefinition with given id, return PackagedProductDefinition or OperationOutcome error
func (fc *FhirClient) UpdatePackagedProductDefinition(updateResource *r4b.PackagedProductDefinition) (*r4b.PackagedProductDefinition, error) {
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

	var parsed r4b.PackagedProductDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdatePackagedProductDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return PackagedProductDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPackagedProductDefinition(patchResource *r4b.PackagedProductDefinition) (*r4b.PackagedProductDefinition, error) {
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

	var parsed r4b.PackagedProductDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchPackagedProductDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete PackagedProductDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePackagedProductDefinition(deleteResource *r4b.PackagedProductDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePackagedProductDefinition given nil PackagedProductDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePackagedProductDefinition given PackagedProductDefinition with nil Id (Id required to delete)")
	}
	return fc.DeletePackagedProductDefinitionById(*deleteResource.Id)
}

// delete PackagedProductDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePackagedProductDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PackagedProductDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeletePackagedProductDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create Patient, return id of created Patient or OperationOutcome error
func (fc *FhirClient) CreatePatient(createResource *r4b.Patient) (*r4b.Patient, error) {
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

	var parsed r4b.Patient
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreatePatient: %s", err)
		}
	}
	return &parsed, err
}

// read Patient from fhir server by id, return Patient or OperationOutcome error
func (fc *FhirClient) ReadPatient(id string) (*r4b.Patient, error) {
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

	var parsed r4b.Patient
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadPatient: %s", err)
		}
	}
	return &parsed, err
}

// replace Patient if exists on server, else create new Patient with given id, return Patient or OperationOutcome error
func (fc *FhirClient) UpdatePatient(updateResource *r4b.Patient) (*r4b.Patient, error) {
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

	var parsed r4b.Patient
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdatePatient: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Patient or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPatient(patchResource *r4b.Patient) (*r4b.Patient, error) {
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

	var parsed r4b.Patient
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchPatient: %s", err)
		}
	}
	return &parsed, err
}

// delete Patient and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePatient(deleteResource *r4b.Patient) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePatient given nil Patient")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePatient given Patient with nil Id (Id required to delete)")
	}
	return fc.DeletePatientById(*deleteResource.Id)
}

// delete Patient with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePatientById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Patient", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeletePatientById: %s", err)
		}
	}
	return &parsed, err
}

// create PaymentNotice, return id of created PaymentNotice or OperationOutcome error
func (fc *FhirClient) CreatePaymentNotice(createResource *r4b.PaymentNotice) (*r4b.PaymentNotice, error) {
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

	var parsed r4b.PaymentNotice
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreatePaymentNotice: %s", err)
		}
	}
	return &parsed, err
}

// read PaymentNotice from fhir server by id, return PaymentNotice or OperationOutcome error
func (fc *FhirClient) ReadPaymentNotice(id string) (*r4b.PaymentNotice, error) {
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

	var parsed r4b.PaymentNotice
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadPaymentNotice: %s", err)
		}
	}
	return &parsed, err
}

// replace PaymentNotice if exists on server, else create new PaymentNotice with given id, return PaymentNotice or OperationOutcome error
func (fc *FhirClient) UpdatePaymentNotice(updateResource *r4b.PaymentNotice) (*r4b.PaymentNotice, error) {
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

	var parsed r4b.PaymentNotice
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdatePaymentNotice: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return PaymentNotice or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPaymentNotice(patchResource *r4b.PaymentNotice) (*r4b.PaymentNotice, error) {
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

	var parsed r4b.PaymentNotice
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchPaymentNotice: %s", err)
		}
	}
	return &parsed, err
}

// delete PaymentNotice and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePaymentNotice(deleteResource *r4b.PaymentNotice) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePaymentNotice given nil PaymentNotice")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePaymentNotice given PaymentNotice with nil Id (Id required to delete)")
	}
	return fc.DeletePaymentNoticeById(*deleteResource.Id)
}

// delete PaymentNotice with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePaymentNoticeById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PaymentNotice", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeletePaymentNoticeById: %s", err)
		}
	}
	return &parsed, err
}

// create PaymentReconciliation, return id of created PaymentReconciliation or OperationOutcome error
func (fc *FhirClient) CreatePaymentReconciliation(createResource *r4b.PaymentReconciliation) (*r4b.PaymentReconciliation, error) {
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

	var parsed r4b.PaymentReconciliation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreatePaymentReconciliation: %s", err)
		}
	}
	return &parsed, err
}

// read PaymentReconciliation from fhir server by id, return PaymentReconciliation or OperationOutcome error
func (fc *FhirClient) ReadPaymentReconciliation(id string) (*r4b.PaymentReconciliation, error) {
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

	var parsed r4b.PaymentReconciliation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadPaymentReconciliation: %s", err)
		}
	}
	return &parsed, err
}

// replace PaymentReconciliation if exists on server, else create new PaymentReconciliation with given id, return PaymentReconciliation or OperationOutcome error
func (fc *FhirClient) UpdatePaymentReconciliation(updateResource *r4b.PaymentReconciliation) (*r4b.PaymentReconciliation, error) {
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

	var parsed r4b.PaymentReconciliation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdatePaymentReconciliation: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return PaymentReconciliation or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPaymentReconciliation(patchResource *r4b.PaymentReconciliation) (*r4b.PaymentReconciliation, error) {
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

	var parsed r4b.PaymentReconciliation
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchPaymentReconciliation: %s", err)
		}
	}
	return &parsed, err
}

// delete PaymentReconciliation and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePaymentReconciliation(deleteResource *r4b.PaymentReconciliation) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePaymentReconciliation given nil PaymentReconciliation")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePaymentReconciliation given PaymentReconciliation with nil Id (Id required to delete)")
	}
	return fc.DeletePaymentReconciliationById(*deleteResource.Id)
}

// delete PaymentReconciliation with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePaymentReconciliationById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PaymentReconciliation", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeletePaymentReconciliationById: %s", err)
		}
	}
	return &parsed, err
}

// create Person, return id of created Person or OperationOutcome error
func (fc *FhirClient) CreatePerson(createResource *r4b.Person) (*r4b.Person, error) {
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

	var parsed r4b.Person
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreatePerson: %s", err)
		}
	}
	return &parsed, err
}

// read Person from fhir server by id, return Person or OperationOutcome error
func (fc *FhirClient) ReadPerson(id string) (*r4b.Person, error) {
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

	var parsed r4b.Person
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadPerson: %s", err)
		}
	}
	return &parsed, err
}

// replace Person if exists on server, else create new Person with given id, return Person or OperationOutcome error
func (fc *FhirClient) UpdatePerson(updateResource *r4b.Person) (*r4b.Person, error) {
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

	var parsed r4b.Person
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdatePerson: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Person or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPerson(patchResource *r4b.Person) (*r4b.Person, error) {
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

	var parsed r4b.Person
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchPerson: %s", err)
		}
	}
	return &parsed, err
}

// delete Person and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePerson(deleteResource *r4b.Person) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePerson given nil Person")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePerson given Person with nil Id (Id required to delete)")
	}
	return fc.DeletePersonById(*deleteResource.Id)
}

// delete Person with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePersonById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Person", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeletePersonById: %s", err)
		}
	}
	return &parsed, err
}

// create PlanDefinition, return id of created PlanDefinition or OperationOutcome error
func (fc *FhirClient) CreatePlanDefinition(createResource *r4b.PlanDefinition) (*r4b.PlanDefinition, error) {
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

	var parsed r4b.PlanDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreatePlanDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read PlanDefinition from fhir server by id, return PlanDefinition or OperationOutcome error
func (fc *FhirClient) ReadPlanDefinition(id string) (*r4b.PlanDefinition, error) {
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

	var parsed r4b.PlanDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadPlanDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace PlanDefinition if exists on server, else create new PlanDefinition with given id, return PlanDefinition or OperationOutcome error
func (fc *FhirClient) UpdatePlanDefinition(updateResource *r4b.PlanDefinition) (*r4b.PlanDefinition, error) {
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

	var parsed r4b.PlanDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdatePlanDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return PlanDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPlanDefinition(patchResource *r4b.PlanDefinition) (*r4b.PlanDefinition, error) {
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

	var parsed r4b.PlanDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchPlanDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete PlanDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePlanDefinition(deleteResource *r4b.PlanDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePlanDefinition given nil PlanDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePlanDefinition given PlanDefinition with nil Id (Id required to delete)")
	}
	return fc.DeletePlanDefinitionById(*deleteResource.Id)
}

// delete PlanDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePlanDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PlanDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeletePlanDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create Practitioner, return id of created Practitioner or OperationOutcome error
func (fc *FhirClient) CreatePractitioner(createResource *r4b.Practitioner) (*r4b.Practitioner, error) {
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

	var parsed r4b.Practitioner
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreatePractitioner: %s", err)
		}
	}
	return &parsed, err
}

// read Practitioner from fhir server by id, return Practitioner or OperationOutcome error
func (fc *FhirClient) ReadPractitioner(id string) (*r4b.Practitioner, error) {
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

	var parsed r4b.Practitioner
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadPractitioner: %s", err)
		}
	}
	return &parsed, err
}

// replace Practitioner if exists on server, else create new Practitioner with given id, return Practitioner or OperationOutcome error
func (fc *FhirClient) UpdatePractitioner(updateResource *r4b.Practitioner) (*r4b.Practitioner, error) {
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

	var parsed r4b.Practitioner
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdatePractitioner: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Practitioner or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPractitioner(patchResource *r4b.Practitioner) (*r4b.Practitioner, error) {
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

	var parsed r4b.Practitioner
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchPractitioner: %s", err)
		}
	}
	return &parsed, err
}

// delete Practitioner and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePractitioner(deleteResource *r4b.Practitioner) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePractitioner given nil Practitioner")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePractitioner given Practitioner with nil Id (Id required to delete)")
	}
	return fc.DeletePractitionerById(*deleteResource.Id)
}

// delete Practitioner with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePractitionerById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Practitioner", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeletePractitionerById: %s", err)
		}
	}
	return &parsed, err
}

// create PractitionerRole, return id of created PractitionerRole or OperationOutcome error
func (fc *FhirClient) CreatePractitionerRole(createResource *r4b.PractitionerRole) (*r4b.PractitionerRole, error) {
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

	var parsed r4b.PractitionerRole
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreatePractitionerRole: %s", err)
		}
	}
	return &parsed, err
}

// read PractitionerRole from fhir server by id, return PractitionerRole or OperationOutcome error
func (fc *FhirClient) ReadPractitionerRole(id string) (*r4b.PractitionerRole, error) {
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

	var parsed r4b.PractitionerRole
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadPractitionerRole: %s", err)
		}
	}
	return &parsed, err
}

// replace PractitionerRole if exists on server, else create new PractitionerRole with given id, return PractitionerRole or OperationOutcome error
func (fc *FhirClient) UpdatePractitionerRole(updateResource *r4b.PractitionerRole) (*r4b.PractitionerRole, error) {
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

	var parsed r4b.PractitionerRole
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdatePractitionerRole: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return PractitionerRole or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchPractitionerRole(patchResource *r4b.PractitionerRole) (*r4b.PractitionerRole, error) {
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

	var parsed r4b.PractitionerRole
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchPractitionerRole: %s", err)
		}
	}
	return &parsed, err
}

// delete PractitionerRole and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePractitionerRole(deleteResource *r4b.PractitionerRole) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeletePractitionerRole given nil PractitionerRole")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeletePractitionerRole given PractitionerRole with nil Id (Id required to delete)")
	}
	return fc.DeletePractitionerRoleById(*deleteResource.Id)
}

// delete PractitionerRole with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeletePractitionerRoleById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "PractitionerRole", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeletePractitionerRoleById: %s", err)
		}
	}
	return &parsed, err
}

// create Procedure, return id of created Procedure or OperationOutcome error
func (fc *FhirClient) CreateProcedure(createResource *r4b.Procedure) (*r4b.Procedure, error) {
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

	var parsed r4b.Procedure
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateProcedure: %s", err)
		}
	}
	return &parsed, err
}

// read Procedure from fhir server by id, return Procedure or OperationOutcome error
func (fc *FhirClient) ReadProcedure(id string) (*r4b.Procedure, error) {
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

	var parsed r4b.Procedure
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadProcedure: %s", err)
		}
	}
	return &parsed, err
}

// replace Procedure if exists on server, else create new Procedure with given id, return Procedure or OperationOutcome error
func (fc *FhirClient) UpdateProcedure(updateResource *r4b.Procedure) (*r4b.Procedure, error) {
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

	var parsed r4b.Procedure
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateProcedure: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Procedure or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchProcedure(patchResource *r4b.Procedure) (*r4b.Procedure, error) {
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

	var parsed r4b.Procedure
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchProcedure: %s", err)
		}
	}
	return &parsed, err
}

// delete Procedure and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteProcedure(deleteResource *r4b.Procedure) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteProcedure given nil Procedure")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteProcedure given Procedure with nil Id (Id required to delete)")
	}
	return fc.DeleteProcedureById(*deleteResource.Id)
}

// delete Procedure with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteProcedureById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Procedure", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteProcedureById: %s", err)
		}
	}
	return &parsed, err
}

// create Provenance, return id of created Provenance or OperationOutcome error
func (fc *FhirClient) CreateProvenance(createResource *r4b.Provenance) (*r4b.Provenance, error) {
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

	var parsed r4b.Provenance
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateProvenance: %s", err)
		}
	}
	return &parsed, err
}

// read Provenance from fhir server by id, return Provenance or OperationOutcome error
func (fc *FhirClient) ReadProvenance(id string) (*r4b.Provenance, error) {
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

	var parsed r4b.Provenance
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadProvenance: %s", err)
		}
	}
	return &parsed, err
}

// replace Provenance if exists on server, else create new Provenance with given id, return Provenance or OperationOutcome error
func (fc *FhirClient) UpdateProvenance(updateResource *r4b.Provenance) (*r4b.Provenance, error) {
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

	var parsed r4b.Provenance
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateProvenance: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Provenance or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchProvenance(patchResource *r4b.Provenance) (*r4b.Provenance, error) {
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

	var parsed r4b.Provenance
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchProvenance: %s", err)
		}
	}
	return &parsed, err
}

// delete Provenance and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteProvenance(deleteResource *r4b.Provenance) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteProvenance given nil Provenance")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteProvenance given Provenance with nil Id (Id required to delete)")
	}
	return fc.DeleteProvenanceById(*deleteResource.Id)
}

// delete Provenance with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteProvenanceById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Provenance", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteProvenanceById: %s", err)
		}
	}
	return &parsed, err
}

// create Questionnaire, return id of created Questionnaire or OperationOutcome error
func (fc *FhirClient) CreateQuestionnaire(createResource *r4b.Questionnaire) (*r4b.Questionnaire, error) {
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

	var parsed r4b.Questionnaire
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateQuestionnaire: %s", err)
		}
	}
	return &parsed, err
}

// read Questionnaire from fhir server by id, return Questionnaire or OperationOutcome error
func (fc *FhirClient) ReadQuestionnaire(id string) (*r4b.Questionnaire, error) {
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

	var parsed r4b.Questionnaire
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadQuestionnaire: %s", err)
		}
	}
	return &parsed, err
}

// replace Questionnaire if exists on server, else create new Questionnaire with given id, return Questionnaire or OperationOutcome error
func (fc *FhirClient) UpdateQuestionnaire(updateResource *r4b.Questionnaire) (*r4b.Questionnaire, error) {
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

	var parsed r4b.Questionnaire
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateQuestionnaire: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Questionnaire or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchQuestionnaire(patchResource *r4b.Questionnaire) (*r4b.Questionnaire, error) {
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

	var parsed r4b.Questionnaire
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchQuestionnaire: %s", err)
		}
	}
	return &parsed, err
}

// delete Questionnaire and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteQuestionnaire(deleteResource *r4b.Questionnaire) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteQuestionnaire given nil Questionnaire")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteQuestionnaire given Questionnaire with nil Id (Id required to delete)")
	}
	return fc.DeleteQuestionnaireById(*deleteResource.Id)
}

// delete Questionnaire with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteQuestionnaireById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Questionnaire", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteQuestionnaireById: %s", err)
		}
	}
	return &parsed, err
}

// create QuestionnaireResponse, return id of created QuestionnaireResponse or OperationOutcome error
func (fc *FhirClient) CreateQuestionnaireResponse(createResource *r4b.QuestionnaireResponse) (*r4b.QuestionnaireResponse, error) {
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

	var parsed r4b.QuestionnaireResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateQuestionnaireResponse: %s", err)
		}
	}
	return &parsed, err
}

// read QuestionnaireResponse from fhir server by id, return QuestionnaireResponse or OperationOutcome error
func (fc *FhirClient) ReadQuestionnaireResponse(id string) (*r4b.QuestionnaireResponse, error) {
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

	var parsed r4b.QuestionnaireResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadQuestionnaireResponse: %s", err)
		}
	}
	return &parsed, err
}

// replace QuestionnaireResponse if exists on server, else create new QuestionnaireResponse with given id, return QuestionnaireResponse or OperationOutcome error
func (fc *FhirClient) UpdateQuestionnaireResponse(updateResource *r4b.QuestionnaireResponse) (*r4b.QuestionnaireResponse, error) {
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

	var parsed r4b.QuestionnaireResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateQuestionnaireResponse: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return QuestionnaireResponse or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchQuestionnaireResponse(patchResource *r4b.QuestionnaireResponse) (*r4b.QuestionnaireResponse, error) {
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

	var parsed r4b.QuestionnaireResponse
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchQuestionnaireResponse: %s", err)
		}
	}
	return &parsed, err
}

// delete QuestionnaireResponse and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteQuestionnaireResponse(deleteResource *r4b.QuestionnaireResponse) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteQuestionnaireResponse given nil QuestionnaireResponse")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteQuestionnaireResponse given QuestionnaireResponse with nil Id (Id required to delete)")
	}
	return fc.DeleteQuestionnaireResponseById(*deleteResource.Id)
}

// delete QuestionnaireResponse with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteQuestionnaireResponseById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "QuestionnaireResponse", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteQuestionnaireResponseById: %s", err)
		}
	}
	return &parsed, err
}

// create RegulatedAuthorization, return id of created RegulatedAuthorization or OperationOutcome error
func (fc *FhirClient) CreateRegulatedAuthorization(createResource *r4b.RegulatedAuthorization) (*r4b.RegulatedAuthorization, error) {
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

	var parsed r4b.RegulatedAuthorization
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateRegulatedAuthorization: %s", err)
		}
	}
	return &parsed, err
}

// read RegulatedAuthorization from fhir server by id, return RegulatedAuthorization or OperationOutcome error
func (fc *FhirClient) ReadRegulatedAuthorization(id string) (*r4b.RegulatedAuthorization, error) {
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

	var parsed r4b.RegulatedAuthorization
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadRegulatedAuthorization: %s", err)
		}
	}
	return &parsed, err
}

// replace RegulatedAuthorization if exists on server, else create new RegulatedAuthorization with given id, return RegulatedAuthorization or OperationOutcome error
func (fc *FhirClient) UpdateRegulatedAuthorization(updateResource *r4b.RegulatedAuthorization) (*r4b.RegulatedAuthorization, error) {
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

	var parsed r4b.RegulatedAuthorization
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateRegulatedAuthorization: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return RegulatedAuthorization or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchRegulatedAuthorization(patchResource *r4b.RegulatedAuthorization) (*r4b.RegulatedAuthorization, error) {
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

	var parsed r4b.RegulatedAuthorization
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchRegulatedAuthorization: %s", err)
		}
	}
	return &parsed, err
}

// delete RegulatedAuthorization and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRegulatedAuthorization(deleteResource *r4b.RegulatedAuthorization) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteRegulatedAuthorization given nil RegulatedAuthorization")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteRegulatedAuthorization given RegulatedAuthorization with nil Id (Id required to delete)")
	}
	return fc.DeleteRegulatedAuthorizationById(*deleteResource.Id)
}

// delete RegulatedAuthorization with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRegulatedAuthorizationById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RegulatedAuthorization", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteRegulatedAuthorizationById: %s", err)
		}
	}
	return &parsed, err
}

// create RelatedPerson, return id of created RelatedPerson or OperationOutcome error
func (fc *FhirClient) CreateRelatedPerson(createResource *r4b.RelatedPerson) (*r4b.RelatedPerson, error) {
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

	var parsed r4b.RelatedPerson
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateRelatedPerson: %s", err)
		}
	}
	return &parsed, err
}

// read RelatedPerson from fhir server by id, return RelatedPerson or OperationOutcome error
func (fc *FhirClient) ReadRelatedPerson(id string) (*r4b.RelatedPerson, error) {
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

	var parsed r4b.RelatedPerson
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadRelatedPerson: %s", err)
		}
	}
	return &parsed, err
}

// replace RelatedPerson if exists on server, else create new RelatedPerson with given id, return RelatedPerson or OperationOutcome error
func (fc *FhirClient) UpdateRelatedPerson(updateResource *r4b.RelatedPerson) (*r4b.RelatedPerson, error) {
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

	var parsed r4b.RelatedPerson
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateRelatedPerson: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return RelatedPerson or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchRelatedPerson(patchResource *r4b.RelatedPerson) (*r4b.RelatedPerson, error) {
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

	var parsed r4b.RelatedPerson
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchRelatedPerson: %s", err)
		}
	}
	return &parsed, err
}

// delete RelatedPerson and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRelatedPerson(deleteResource *r4b.RelatedPerson) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteRelatedPerson given nil RelatedPerson")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteRelatedPerson given RelatedPerson with nil Id (Id required to delete)")
	}
	return fc.DeleteRelatedPersonById(*deleteResource.Id)
}

// delete RelatedPerson with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRelatedPersonById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RelatedPerson", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteRelatedPersonById: %s", err)
		}
	}
	return &parsed, err
}

// create RequestGroup, return id of created RequestGroup or OperationOutcome error
func (fc *FhirClient) CreateRequestGroup(createResource *r4b.RequestGroup) (*r4b.RequestGroup, error) {
	if createResource == nil {
		return nil, errors.New("CreateRequestGroup called with nil RequestGroup")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RequestGroup")
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

	var parsed r4b.RequestGroup
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateRequestGroup: %s", err)
		}
	}
	return &parsed, err
}

// read RequestGroup from fhir server by id, return RequestGroup or OperationOutcome error
func (fc *FhirClient) ReadRequestGroup(id string) (*r4b.RequestGroup, error) {
	if id == "" {
		return nil, errors.New("ReadRequestGroup given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RequestGroup", id)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.RequestGroup
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadRequestGroup: %s", err)
		}
	}
	return &parsed, err
}

// replace RequestGroup if exists on server, else create new RequestGroup with given id, return RequestGroup or OperationOutcome error
func (fc *FhirClient) UpdateRequestGroup(updateResource *r4b.RequestGroup) (*r4b.RequestGroup, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateRequestGroup called with nil RequestGroup")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RequestGroup", *updateResource.Id)
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

	var parsed r4b.RequestGroup
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateRequestGroup: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return RequestGroup or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchRequestGroup(patchResource *r4b.RequestGroup) (*r4b.RequestGroup, error) {
	if patchResource == nil {
		return nil, errors.New("PatchRequestGroup given nil RequestGroup")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchRequestGroup given RequestGroup without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RequestGroup", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchRequestGroup error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchRequestGroup error using patch representation as json")
	}
	req, err := http.NewRequest(http.MethodPatch, reqUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json-patch+json")

	var parsed r4b.RequestGroup
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchRequestGroup: %s", err)
		}
	}
	return &parsed, err
}

// delete RequestGroup and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRequestGroup(deleteResource *r4b.RequestGroup) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteRequestGroup given nil RequestGroup")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteRequestGroup given RequestGroup with nil Id (Id required to delete)")
	}
	return fc.DeleteRequestGroupById(*deleteResource.Id)
}

// delete RequestGroup with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRequestGroupById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RequestGroup", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteRequestGroupById: %s", err)
		}
	}
	return &parsed, err
}

// create ResearchDefinition, return id of created ResearchDefinition or OperationOutcome error
func (fc *FhirClient) CreateResearchDefinition(createResource *r4b.ResearchDefinition) (*r4b.ResearchDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateResearchDefinition called with nil ResearchDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchDefinition")
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

	var parsed r4b.ResearchDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateResearchDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read ResearchDefinition from fhir server by id, return ResearchDefinition or OperationOutcome error
func (fc *FhirClient) ReadResearchDefinition(id string) (*r4b.ResearchDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadResearchDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchDefinition", id)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.ResearchDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadResearchDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace ResearchDefinition if exists on server, else create new ResearchDefinition with given id, return ResearchDefinition or OperationOutcome error
func (fc *FhirClient) UpdateResearchDefinition(updateResource *r4b.ResearchDefinition) (*r4b.ResearchDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateResearchDefinition called with nil ResearchDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchDefinition", *updateResource.Id)
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

	var parsed r4b.ResearchDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateResearchDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ResearchDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchResearchDefinition(patchResource *r4b.ResearchDefinition) (*r4b.ResearchDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchResearchDefinition given nil ResearchDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchResearchDefinition given ResearchDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchResearchDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchResearchDefinition error using patch representation as json")
	}
	req, err := http.NewRequest(http.MethodPatch, reqUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json-patch+json")

	var parsed r4b.ResearchDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchResearchDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete ResearchDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteResearchDefinition(deleteResource *r4b.ResearchDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteResearchDefinition given nil ResearchDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteResearchDefinition given ResearchDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteResearchDefinitionById(*deleteResource.Id)
}

// delete ResearchDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteResearchDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteResearchDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create ResearchElementDefinition, return id of created ResearchElementDefinition or OperationOutcome error
func (fc *FhirClient) CreateResearchElementDefinition(createResource *r4b.ResearchElementDefinition) (*r4b.ResearchElementDefinition, error) {
	if createResource == nil {
		return nil, errors.New("CreateResearchElementDefinition called with nil ResearchElementDefinition")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchElementDefinition")
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

	var parsed r4b.ResearchElementDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateResearchElementDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read ResearchElementDefinition from fhir server by id, return ResearchElementDefinition or OperationOutcome error
func (fc *FhirClient) ReadResearchElementDefinition(id string) (*r4b.ResearchElementDefinition, error) {
	if id == "" {
		return nil, errors.New("ReadResearchElementDefinition given empty id")
	}
	parts := strings.Split(id, "/") //param is id but can pass in url too
	id = parts[len(parts)-1]
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchElementDefinition", id)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.ResearchElementDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadResearchElementDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace ResearchElementDefinition if exists on server, else create new ResearchElementDefinition with given id, return ResearchElementDefinition or OperationOutcome error
func (fc *FhirClient) UpdateResearchElementDefinition(updateResource *r4b.ResearchElementDefinition) (*r4b.ResearchElementDefinition, error) {
	if updateResource == nil {
		return nil, errors.New("UpdateResearchElementDefinition called with nil ResearchElementDefinition")
	}
	if updateResource.Id == nil {
		//update will PUT resource at its Id, so it doesn't make sense without one
		//going to generate an id for you, modifying your original struct (!!)
		genId := "bultaoreune-PUT-" + strings.ReplaceAll(time.Now().Format(time.RFC3339Nano), ":", "-")
		updateResource.Id = &genId
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchElementDefinition", *updateResource.Id)
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

	var parsed r4b.ResearchElementDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateResearchElementDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ResearchElementDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchResearchElementDefinition(patchResource *r4b.ResearchElementDefinition) (*r4b.ResearchElementDefinition, error) {
	if patchResource == nil {
		return nil, errors.New("PatchResearchElementDefinition given nil ResearchElementDefinition")
	}
	if patchResource.Id == nil {
		return nil, errors.New("PatchResearchElementDefinition given ResearchElementDefinition without id, can't modify resource on server")
	}
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchElementDefinition", *patchResource.Id)
	if err != nil {
		return nil, err
	}

	structAsPatch, err := convertToPatch(patchResource)
	if err != nil {
		return nil, errors.New("PatchResearchElementDefinition error converting struct to patch representation")
	}
	reqBody, err := json.Marshal(structAsPatch)
	if err != nil {
		return nil, errors.New("PatchResearchElementDefinition error using patch representation as json")
	}
	req, err := http.NewRequest(http.MethodPatch, reqUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json-patch+json")

	var parsed r4b.ResearchElementDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchResearchElementDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete ResearchElementDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteResearchElementDefinition(deleteResource *r4b.ResearchElementDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteResearchElementDefinition given nil ResearchElementDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteResearchElementDefinition given ResearchElementDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteResearchElementDefinitionById(*deleteResource.Id)
}

// delete ResearchElementDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteResearchElementDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchElementDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteResearchElementDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create ResearchStudy, return id of created ResearchStudy or OperationOutcome error
func (fc *FhirClient) CreateResearchStudy(createResource *r4b.ResearchStudy) (*r4b.ResearchStudy, error) {
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

	var parsed r4b.ResearchStudy
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateResearchStudy: %s", err)
		}
	}
	return &parsed, err
}

// read ResearchStudy from fhir server by id, return ResearchStudy or OperationOutcome error
func (fc *FhirClient) ReadResearchStudy(id string) (*r4b.ResearchStudy, error) {
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

	var parsed r4b.ResearchStudy
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadResearchStudy: %s", err)
		}
	}
	return &parsed, err
}

// replace ResearchStudy if exists on server, else create new ResearchStudy with given id, return ResearchStudy or OperationOutcome error
func (fc *FhirClient) UpdateResearchStudy(updateResource *r4b.ResearchStudy) (*r4b.ResearchStudy, error) {
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

	var parsed r4b.ResearchStudy
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateResearchStudy: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ResearchStudy or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchResearchStudy(patchResource *r4b.ResearchStudy) (*r4b.ResearchStudy, error) {
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

	var parsed r4b.ResearchStudy
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchResearchStudy: %s", err)
		}
	}
	return &parsed, err
}

// delete ResearchStudy and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteResearchStudy(deleteResource *r4b.ResearchStudy) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteResearchStudy given nil ResearchStudy")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteResearchStudy given ResearchStudy with nil Id (Id required to delete)")
	}
	return fc.DeleteResearchStudyById(*deleteResource.Id)
}

// delete ResearchStudy with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteResearchStudyById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchStudy", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteResearchStudyById: %s", err)
		}
	}
	return &parsed, err
}

// create ResearchSubject, return id of created ResearchSubject or OperationOutcome error
func (fc *FhirClient) CreateResearchSubject(createResource *r4b.ResearchSubject) (*r4b.ResearchSubject, error) {
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

	var parsed r4b.ResearchSubject
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateResearchSubject: %s", err)
		}
	}
	return &parsed, err
}

// read ResearchSubject from fhir server by id, return ResearchSubject or OperationOutcome error
func (fc *FhirClient) ReadResearchSubject(id string) (*r4b.ResearchSubject, error) {
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

	var parsed r4b.ResearchSubject
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadResearchSubject: %s", err)
		}
	}
	return &parsed, err
}

// replace ResearchSubject if exists on server, else create new ResearchSubject with given id, return ResearchSubject or OperationOutcome error
func (fc *FhirClient) UpdateResearchSubject(updateResource *r4b.ResearchSubject) (*r4b.ResearchSubject, error) {
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

	var parsed r4b.ResearchSubject
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateResearchSubject: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ResearchSubject or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchResearchSubject(patchResource *r4b.ResearchSubject) (*r4b.ResearchSubject, error) {
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

	var parsed r4b.ResearchSubject
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchResearchSubject: %s", err)
		}
	}
	return &parsed, err
}

// delete ResearchSubject and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteResearchSubject(deleteResource *r4b.ResearchSubject) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteResearchSubject given nil ResearchSubject")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteResearchSubject given ResearchSubject with nil Id (Id required to delete)")
	}
	return fc.DeleteResearchSubjectById(*deleteResource.Id)
}

// delete ResearchSubject with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteResearchSubjectById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ResearchSubject", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteResearchSubjectById: %s", err)
		}
	}
	return &parsed, err
}

// create RiskAssessment, return id of created RiskAssessment or OperationOutcome error
func (fc *FhirClient) CreateRiskAssessment(createResource *r4b.RiskAssessment) (*r4b.RiskAssessment, error) {
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

	var parsed r4b.RiskAssessment
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateRiskAssessment: %s", err)
		}
	}
	return &parsed, err
}

// read RiskAssessment from fhir server by id, return RiskAssessment or OperationOutcome error
func (fc *FhirClient) ReadRiskAssessment(id string) (*r4b.RiskAssessment, error) {
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

	var parsed r4b.RiskAssessment
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadRiskAssessment: %s", err)
		}
	}
	return &parsed, err
}

// replace RiskAssessment if exists on server, else create new RiskAssessment with given id, return RiskAssessment or OperationOutcome error
func (fc *FhirClient) UpdateRiskAssessment(updateResource *r4b.RiskAssessment) (*r4b.RiskAssessment, error) {
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

	var parsed r4b.RiskAssessment
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateRiskAssessment: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return RiskAssessment or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchRiskAssessment(patchResource *r4b.RiskAssessment) (*r4b.RiskAssessment, error) {
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

	var parsed r4b.RiskAssessment
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchRiskAssessment: %s", err)
		}
	}
	return &parsed, err
}

// delete RiskAssessment and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRiskAssessment(deleteResource *r4b.RiskAssessment) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteRiskAssessment given nil RiskAssessment")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteRiskAssessment given RiskAssessment with nil Id (Id required to delete)")
	}
	return fc.DeleteRiskAssessmentById(*deleteResource.Id)
}

// delete RiskAssessment with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteRiskAssessmentById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "RiskAssessment", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteRiskAssessmentById: %s", err)
		}
	}
	return &parsed, err
}

// create Schedule, return id of created Schedule or OperationOutcome error
func (fc *FhirClient) CreateSchedule(createResource *r4b.Schedule) (*r4b.Schedule, error) {
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

	var parsed r4b.Schedule
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateSchedule: %s", err)
		}
	}
	return &parsed, err
}

// read Schedule from fhir server by id, return Schedule or OperationOutcome error
func (fc *FhirClient) ReadSchedule(id string) (*r4b.Schedule, error) {
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

	var parsed r4b.Schedule
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadSchedule: %s", err)
		}
	}
	return &parsed, err
}

// replace Schedule if exists on server, else create new Schedule with given id, return Schedule or OperationOutcome error
func (fc *FhirClient) UpdateSchedule(updateResource *r4b.Schedule) (*r4b.Schedule, error) {
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

	var parsed r4b.Schedule
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateSchedule: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Schedule or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSchedule(patchResource *r4b.Schedule) (*r4b.Schedule, error) {
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

	var parsed r4b.Schedule
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchSchedule: %s", err)
		}
	}
	return &parsed, err
}

// delete Schedule and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSchedule(deleteResource *r4b.Schedule) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSchedule given nil Schedule")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSchedule given Schedule with nil Id (Id required to delete)")
	}
	return fc.DeleteScheduleById(*deleteResource.Id)
}

// delete Schedule with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteScheduleById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Schedule", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteScheduleById: %s", err)
		}
	}
	return &parsed, err
}

// create SearchParameter, return id of created SearchParameter or OperationOutcome error
func (fc *FhirClient) CreateSearchParameter(createResource *r4b.SearchParameter) (*r4b.SearchParameter, error) {
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

	var parsed r4b.SearchParameter
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateSearchParameter: %s", err)
		}
	}
	return &parsed, err
}

// read SearchParameter from fhir server by id, return SearchParameter or OperationOutcome error
func (fc *FhirClient) ReadSearchParameter(id string) (*r4b.SearchParameter, error) {
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

	var parsed r4b.SearchParameter
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadSearchParameter: %s", err)
		}
	}
	return &parsed, err
}

// replace SearchParameter if exists on server, else create new SearchParameter with given id, return SearchParameter or OperationOutcome error
func (fc *FhirClient) UpdateSearchParameter(updateResource *r4b.SearchParameter) (*r4b.SearchParameter, error) {
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

	var parsed r4b.SearchParameter
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateSearchParameter: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return SearchParameter or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSearchParameter(patchResource *r4b.SearchParameter) (*r4b.SearchParameter, error) {
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

	var parsed r4b.SearchParameter
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchSearchParameter: %s", err)
		}
	}
	return &parsed, err
}

// delete SearchParameter and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSearchParameter(deleteResource *r4b.SearchParameter) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSearchParameter given nil SearchParameter")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSearchParameter given SearchParameter with nil Id (Id required to delete)")
	}
	return fc.DeleteSearchParameterById(*deleteResource.Id)
}

// delete SearchParameter with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSearchParameterById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SearchParameter", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteSearchParameterById: %s", err)
		}
	}
	return &parsed, err
}

// create ServiceRequest, return id of created ServiceRequest or OperationOutcome error
func (fc *FhirClient) CreateServiceRequest(createResource *r4b.ServiceRequest) (*r4b.ServiceRequest, error) {
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

	var parsed r4b.ServiceRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateServiceRequest: %s", err)
		}
	}
	return &parsed, err
}

// read ServiceRequest from fhir server by id, return ServiceRequest or OperationOutcome error
func (fc *FhirClient) ReadServiceRequest(id string) (*r4b.ServiceRequest, error) {
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

	var parsed r4b.ServiceRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadServiceRequest: %s", err)
		}
	}
	return &parsed, err
}

// replace ServiceRequest if exists on server, else create new ServiceRequest with given id, return ServiceRequest or OperationOutcome error
func (fc *FhirClient) UpdateServiceRequest(updateResource *r4b.ServiceRequest) (*r4b.ServiceRequest, error) {
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

	var parsed r4b.ServiceRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateServiceRequest: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ServiceRequest or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchServiceRequest(patchResource *r4b.ServiceRequest) (*r4b.ServiceRequest, error) {
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

	var parsed r4b.ServiceRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchServiceRequest: %s", err)
		}
	}
	return &parsed, err
}

// delete ServiceRequest and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteServiceRequest(deleteResource *r4b.ServiceRequest) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteServiceRequest given nil ServiceRequest")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteServiceRequest given ServiceRequest with nil Id (Id required to delete)")
	}
	return fc.DeleteServiceRequestById(*deleteResource.Id)
}

// delete ServiceRequest with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteServiceRequestById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ServiceRequest", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteServiceRequestById: %s", err)
		}
	}
	return &parsed, err
}

// create Slot, return id of created Slot or OperationOutcome error
func (fc *FhirClient) CreateSlot(createResource *r4b.Slot) (*r4b.Slot, error) {
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

	var parsed r4b.Slot
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateSlot: %s", err)
		}
	}
	return &parsed, err
}

// read Slot from fhir server by id, return Slot or OperationOutcome error
func (fc *FhirClient) ReadSlot(id string) (*r4b.Slot, error) {
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

	var parsed r4b.Slot
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadSlot: %s", err)
		}
	}
	return &parsed, err
}

// replace Slot if exists on server, else create new Slot with given id, return Slot or OperationOutcome error
func (fc *FhirClient) UpdateSlot(updateResource *r4b.Slot) (*r4b.Slot, error) {
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

	var parsed r4b.Slot
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateSlot: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Slot or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSlot(patchResource *r4b.Slot) (*r4b.Slot, error) {
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

	var parsed r4b.Slot
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchSlot: %s", err)
		}
	}
	return &parsed, err
}

// delete Slot and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSlot(deleteResource *r4b.Slot) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSlot given nil Slot")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSlot given Slot with nil Id (Id required to delete)")
	}
	return fc.DeleteSlotById(*deleteResource.Id)
}

// delete Slot with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSlotById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Slot", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteSlotById: %s", err)
		}
	}
	return &parsed, err
}

// create Specimen, return id of created Specimen or OperationOutcome error
func (fc *FhirClient) CreateSpecimen(createResource *r4b.Specimen) (*r4b.Specimen, error) {
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

	var parsed r4b.Specimen
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateSpecimen: %s", err)
		}
	}
	return &parsed, err
}

// read Specimen from fhir server by id, return Specimen or OperationOutcome error
func (fc *FhirClient) ReadSpecimen(id string) (*r4b.Specimen, error) {
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

	var parsed r4b.Specimen
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadSpecimen: %s", err)
		}
	}
	return &parsed, err
}

// replace Specimen if exists on server, else create new Specimen with given id, return Specimen or OperationOutcome error
func (fc *FhirClient) UpdateSpecimen(updateResource *r4b.Specimen) (*r4b.Specimen, error) {
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

	var parsed r4b.Specimen
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateSpecimen: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Specimen or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSpecimen(patchResource *r4b.Specimen) (*r4b.Specimen, error) {
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

	var parsed r4b.Specimen
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchSpecimen: %s", err)
		}
	}
	return &parsed, err
}

// delete Specimen and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSpecimen(deleteResource *r4b.Specimen) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSpecimen given nil Specimen")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSpecimen given Specimen with nil Id (Id required to delete)")
	}
	return fc.DeleteSpecimenById(*deleteResource.Id)
}

// delete Specimen with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSpecimenById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Specimen", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteSpecimenById: %s", err)
		}
	}
	return &parsed, err
}

// create SpecimenDefinition, return id of created SpecimenDefinition or OperationOutcome error
func (fc *FhirClient) CreateSpecimenDefinition(createResource *r4b.SpecimenDefinition) (*r4b.SpecimenDefinition, error) {
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

	var parsed r4b.SpecimenDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateSpecimenDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read SpecimenDefinition from fhir server by id, return SpecimenDefinition or OperationOutcome error
func (fc *FhirClient) ReadSpecimenDefinition(id string) (*r4b.SpecimenDefinition, error) {
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

	var parsed r4b.SpecimenDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadSpecimenDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace SpecimenDefinition if exists on server, else create new SpecimenDefinition with given id, return SpecimenDefinition or OperationOutcome error
func (fc *FhirClient) UpdateSpecimenDefinition(updateResource *r4b.SpecimenDefinition) (*r4b.SpecimenDefinition, error) {
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

	var parsed r4b.SpecimenDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateSpecimenDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return SpecimenDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSpecimenDefinition(patchResource *r4b.SpecimenDefinition) (*r4b.SpecimenDefinition, error) {
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

	var parsed r4b.SpecimenDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchSpecimenDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete SpecimenDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSpecimenDefinition(deleteResource *r4b.SpecimenDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSpecimenDefinition given nil SpecimenDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSpecimenDefinition given SpecimenDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteSpecimenDefinitionById(*deleteResource.Id)
}

// delete SpecimenDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSpecimenDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SpecimenDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteSpecimenDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create StructureDefinition, return id of created StructureDefinition or OperationOutcome error
func (fc *FhirClient) CreateStructureDefinition(createResource *r4b.StructureDefinition) (*r4b.StructureDefinition, error) {
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

	var parsed r4b.StructureDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateStructureDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read StructureDefinition from fhir server by id, return StructureDefinition or OperationOutcome error
func (fc *FhirClient) ReadStructureDefinition(id string) (*r4b.StructureDefinition, error) {
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

	var parsed r4b.StructureDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadStructureDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace StructureDefinition if exists on server, else create new StructureDefinition with given id, return StructureDefinition or OperationOutcome error
func (fc *FhirClient) UpdateStructureDefinition(updateResource *r4b.StructureDefinition) (*r4b.StructureDefinition, error) {
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

	var parsed r4b.StructureDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateStructureDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return StructureDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchStructureDefinition(patchResource *r4b.StructureDefinition) (*r4b.StructureDefinition, error) {
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

	var parsed r4b.StructureDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchStructureDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete StructureDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteStructureDefinition(deleteResource *r4b.StructureDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteStructureDefinition given nil StructureDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteStructureDefinition given StructureDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteStructureDefinitionById(*deleteResource.Id)
}

// delete StructureDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteStructureDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "StructureDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteStructureDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create StructureMap, return id of created StructureMap or OperationOutcome error
func (fc *FhirClient) CreateStructureMap(createResource *r4b.StructureMap) (*r4b.StructureMap, error) {
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

	var parsed r4b.StructureMap
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateStructureMap: %s", err)
		}
	}
	return &parsed, err
}

// read StructureMap from fhir server by id, return StructureMap or OperationOutcome error
func (fc *FhirClient) ReadStructureMap(id string) (*r4b.StructureMap, error) {
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

	var parsed r4b.StructureMap
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadStructureMap: %s", err)
		}
	}
	return &parsed, err
}

// replace StructureMap if exists on server, else create new StructureMap with given id, return StructureMap or OperationOutcome error
func (fc *FhirClient) UpdateStructureMap(updateResource *r4b.StructureMap) (*r4b.StructureMap, error) {
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

	var parsed r4b.StructureMap
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateStructureMap: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return StructureMap or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchStructureMap(patchResource *r4b.StructureMap) (*r4b.StructureMap, error) {
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

	var parsed r4b.StructureMap
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchStructureMap: %s", err)
		}
	}
	return &parsed, err
}

// delete StructureMap and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteStructureMap(deleteResource *r4b.StructureMap) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteStructureMap given nil StructureMap")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteStructureMap given StructureMap with nil Id (Id required to delete)")
	}
	return fc.DeleteStructureMapById(*deleteResource.Id)
}

// delete StructureMap with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteStructureMapById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "StructureMap", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteStructureMapById: %s", err)
		}
	}
	return &parsed, err
}

// create Subscription, return id of created Subscription or OperationOutcome error
func (fc *FhirClient) CreateSubscription(createResource *r4b.Subscription) (*r4b.Subscription, error) {
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

	var parsed r4b.Subscription
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateSubscription: %s", err)
		}
	}
	return &parsed, err
}

// read Subscription from fhir server by id, return Subscription or OperationOutcome error
func (fc *FhirClient) ReadSubscription(id string) (*r4b.Subscription, error) {
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

	var parsed r4b.Subscription
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadSubscription: %s", err)
		}
	}
	return &parsed, err
}

// replace Subscription if exists on server, else create new Subscription with given id, return Subscription or OperationOutcome error
func (fc *FhirClient) UpdateSubscription(updateResource *r4b.Subscription) (*r4b.Subscription, error) {
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

	var parsed r4b.Subscription
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateSubscription: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Subscription or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSubscription(patchResource *r4b.Subscription) (*r4b.Subscription, error) {
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

	var parsed r4b.Subscription
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchSubscription: %s", err)
		}
	}
	return &parsed, err
}

// delete Subscription and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubscription(deleteResource *r4b.Subscription) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSubscription given nil Subscription")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSubscription given Subscription with nil Id (Id required to delete)")
	}
	return fc.DeleteSubscriptionById(*deleteResource.Id)
}

// delete Subscription with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubscriptionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Subscription", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteSubscriptionById: %s", err)
		}
	}
	return &parsed, err
}

// create SubscriptionStatus, return id of created SubscriptionStatus or OperationOutcome error
func (fc *FhirClient) CreateSubscriptionStatus(createResource *r4b.SubscriptionStatus) (*r4b.SubscriptionStatus, error) {
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

	var parsed r4b.SubscriptionStatus
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateSubscriptionStatus: %s", err)
		}
	}
	return &parsed, err
}

// read SubscriptionStatus from fhir server by id, return SubscriptionStatus or OperationOutcome error
func (fc *FhirClient) ReadSubscriptionStatus(id string) (*r4b.SubscriptionStatus, error) {
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

	var parsed r4b.SubscriptionStatus
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadSubscriptionStatus: %s", err)
		}
	}
	return &parsed, err
}

// replace SubscriptionStatus if exists on server, else create new SubscriptionStatus with given id, return SubscriptionStatus or OperationOutcome error
func (fc *FhirClient) UpdateSubscriptionStatus(updateResource *r4b.SubscriptionStatus) (*r4b.SubscriptionStatus, error) {
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

	var parsed r4b.SubscriptionStatus
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateSubscriptionStatus: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return SubscriptionStatus or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSubscriptionStatus(patchResource *r4b.SubscriptionStatus) (*r4b.SubscriptionStatus, error) {
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

	var parsed r4b.SubscriptionStatus
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchSubscriptionStatus: %s", err)
		}
	}
	return &parsed, err
}

// delete SubscriptionStatus and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubscriptionStatus(deleteResource *r4b.SubscriptionStatus) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSubscriptionStatus given nil SubscriptionStatus")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSubscriptionStatus given SubscriptionStatus with nil Id (Id required to delete)")
	}
	return fc.DeleteSubscriptionStatusById(*deleteResource.Id)
}

// delete SubscriptionStatus with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubscriptionStatusById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubscriptionStatus", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteSubscriptionStatusById: %s", err)
		}
	}
	return &parsed, err
}

// create SubscriptionTopic, return id of created SubscriptionTopic or OperationOutcome error
func (fc *FhirClient) CreateSubscriptionTopic(createResource *r4b.SubscriptionTopic) (*r4b.SubscriptionTopic, error) {
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

	var parsed r4b.SubscriptionTopic
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateSubscriptionTopic: %s", err)
		}
	}
	return &parsed, err
}

// read SubscriptionTopic from fhir server by id, return SubscriptionTopic or OperationOutcome error
func (fc *FhirClient) ReadSubscriptionTopic(id string) (*r4b.SubscriptionTopic, error) {
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

	var parsed r4b.SubscriptionTopic
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadSubscriptionTopic: %s", err)
		}
	}
	return &parsed, err
}

// replace SubscriptionTopic if exists on server, else create new SubscriptionTopic with given id, return SubscriptionTopic or OperationOutcome error
func (fc *FhirClient) UpdateSubscriptionTopic(updateResource *r4b.SubscriptionTopic) (*r4b.SubscriptionTopic, error) {
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

	var parsed r4b.SubscriptionTopic
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateSubscriptionTopic: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return SubscriptionTopic or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSubscriptionTopic(patchResource *r4b.SubscriptionTopic) (*r4b.SubscriptionTopic, error) {
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

	var parsed r4b.SubscriptionTopic
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchSubscriptionTopic: %s", err)
		}
	}
	return &parsed, err
}

// delete SubscriptionTopic and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubscriptionTopic(deleteResource *r4b.SubscriptionTopic) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSubscriptionTopic given nil SubscriptionTopic")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSubscriptionTopic given SubscriptionTopic with nil Id (Id required to delete)")
	}
	return fc.DeleteSubscriptionTopicById(*deleteResource.Id)
}

// delete SubscriptionTopic with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubscriptionTopicById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubscriptionTopic", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteSubscriptionTopicById: %s", err)
		}
	}
	return &parsed, err
}

// create Substance, return id of created Substance or OperationOutcome error
func (fc *FhirClient) CreateSubstance(createResource *r4b.Substance) (*r4b.Substance, error) {
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

	var parsed r4b.Substance
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateSubstance: %s", err)
		}
	}
	return &parsed, err
}

// read Substance from fhir server by id, return Substance or OperationOutcome error
func (fc *FhirClient) ReadSubstance(id string) (*r4b.Substance, error) {
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

	var parsed r4b.Substance
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadSubstance: %s", err)
		}
	}
	return &parsed, err
}

// replace Substance if exists on server, else create new Substance with given id, return Substance or OperationOutcome error
func (fc *FhirClient) UpdateSubstance(updateResource *r4b.Substance) (*r4b.Substance, error) {
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

	var parsed r4b.Substance
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateSubstance: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Substance or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSubstance(patchResource *r4b.Substance) (*r4b.Substance, error) {
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

	var parsed r4b.Substance
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchSubstance: %s", err)
		}
	}
	return &parsed, err
}

// delete Substance and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstance(deleteResource *r4b.Substance) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSubstance given nil Substance")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSubstance given Substance with nil Id (Id required to delete)")
	}
	return fc.DeleteSubstanceById(*deleteResource.Id)
}

// delete Substance with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstanceById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Substance", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteSubstanceById: %s", err)
		}
	}
	return &parsed, err
}

// create SubstanceDefinition, return id of created SubstanceDefinition or OperationOutcome error
func (fc *FhirClient) CreateSubstanceDefinition(createResource *r4b.SubstanceDefinition) (*r4b.SubstanceDefinition, error) {
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

	var parsed r4b.SubstanceDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateSubstanceDefinition: %s", err)
		}
	}
	return &parsed, err
}

// read SubstanceDefinition from fhir server by id, return SubstanceDefinition or OperationOutcome error
func (fc *FhirClient) ReadSubstanceDefinition(id string) (*r4b.SubstanceDefinition, error) {
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

	var parsed r4b.SubstanceDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadSubstanceDefinition: %s", err)
		}
	}
	return &parsed, err
}

// replace SubstanceDefinition if exists on server, else create new SubstanceDefinition with given id, return SubstanceDefinition or OperationOutcome error
func (fc *FhirClient) UpdateSubstanceDefinition(updateResource *r4b.SubstanceDefinition) (*r4b.SubstanceDefinition, error) {
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

	var parsed r4b.SubstanceDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateSubstanceDefinition: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return SubstanceDefinition or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSubstanceDefinition(patchResource *r4b.SubstanceDefinition) (*r4b.SubstanceDefinition, error) {
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

	var parsed r4b.SubstanceDefinition
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchSubstanceDefinition: %s", err)
		}
	}
	return &parsed, err
}

// delete SubstanceDefinition and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstanceDefinition(deleteResource *r4b.SubstanceDefinition) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSubstanceDefinition given nil SubstanceDefinition")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSubstanceDefinition given SubstanceDefinition with nil Id (Id required to delete)")
	}
	return fc.DeleteSubstanceDefinitionById(*deleteResource.Id)
}

// delete SubstanceDefinition with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSubstanceDefinitionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SubstanceDefinition", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteSubstanceDefinitionById: %s", err)
		}
	}
	return &parsed, err
}

// create SupplyDelivery, return id of created SupplyDelivery or OperationOutcome error
func (fc *FhirClient) CreateSupplyDelivery(createResource *r4b.SupplyDelivery) (*r4b.SupplyDelivery, error) {
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

	var parsed r4b.SupplyDelivery
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateSupplyDelivery: %s", err)
		}
	}
	return &parsed, err
}

// read SupplyDelivery from fhir server by id, return SupplyDelivery or OperationOutcome error
func (fc *FhirClient) ReadSupplyDelivery(id string) (*r4b.SupplyDelivery, error) {
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

	var parsed r4b.SupplyDelivery
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadSupplyDelivery: %s", err)
		}
	}
	return &parsed, err
}

// replace SupplyDelivery if exists on server, else create new SupplyDelivery with given id, return SupplyDelivery or OperationOutcome error
func (fc *FhirClient) UpdateSupplyDelivery(updateResource *r4b.SupplyDelivery) (*r4b.SupplyDelivery, error) {
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

	var parsed r4b.SupplyDelivery
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateSupplyDelivery: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return SupplyDelivery or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSupplyDelivery(patchResource *r4b.SupplyDelivery) (*r4b.SupplyDelivery, error) {
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

	var parsed r4b.SupplyDelivery
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchSupplyDelivery: %s", err)
		}
	}
	return &parsed, err
}

// delete SupplyDelivery and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSupplyDelivery(deleteResource *r4b.SupplyDelivery) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSupplyDelivery given nil SupplyDelivery")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSupplyDelivery given SupplyDelivery with nil Id (Id required to delete)")
	}
	return fc.DeleteSupplyDeliveryById(*deleteResource.Id)
}

// delete SupplyDelivery with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSupplyDeliveryById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SupplyDelivery", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteSupplyDeliveryById: %s", err)
		}
	}
	return &parsed, err
}

// create SupplyRequest, return id of created SupplyRequest or OperationOutcome error
func (fc *FhirClient) CreateSupplyRequest(createResource *r4b.SupplyRequest) (*r4b.SupplyRequest, error) {
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

	var parsed r4b.SupplyRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateSupplyRequest: %s", err)
		}
	}
	return &parsed, err
}

// read SupplyRequest from fhir server by id, return SupplyRequest or OperationOutcome error
func (fc *FhirClient) ReadSupplyRequest(id string) (*r4b.SupplyRequest, error) {
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

	var parsed r4b.SupplyRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadSupplyRequest: %s", err)
		}
	}
	return &parsed, err
}

// replace SupplyRequest if exists on server, else create new SupplyRequest with given id, return SupplyRequest or OperationOutcome error
func (fc *FhirClient) UpdateSupplyRequest(updateResource *r4b.SupplyRequest) (*r4b.SupplyRequest, error) {
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

	var parsed r4b.SupplyRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateSupplyRequest: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return SupplyRequest or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchSupplyRequest(patchResource *r4b.SupplyRequest) (*r4b.SupplyRequest, error) {
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

	var parsed r4b.SupplyRequest
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchSupplyRequest: %s", err)
		}
	}
	return &parsed, err
}

// delete SupplyRequest and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSupplyRequest(deleteResource *r4b.SupplyRequest) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteSupplyRequest given nil SupplyRequest")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteSupplyRequest given SupplyRequest with nil Id (Id required to delete)")
	}
	return fc.DeleteSupplyRequestById(*deleteResource.Id)
}

// delete SupplyRequest with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteSupplyRequestById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "SupplyRequest", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteSupplyRequestById: %s", err)
		}
	}
	return &parsed, err
}

// create Task, return id of created Task or OperationOutcome error
func (fc *FhirClient) CreateTask(createResource *r4b.Task) (*r4b.Task, error) {
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

	var parsed r4b.Task
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateTask: %s", err)
		}
	}
	return &parsed, err
}

// read Task from fhir server by id, return Task or OperationOutcome error
func (fc *FhirClient) ReadTask(id string) (*r4b.Task, error) {
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

	var parsed r4b.Task
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadTask: %s", err)
		}
	}
	return &parsed, err
}

// replace Task if exists on server, else create new Task with given id, return Task or OperationOutcome error
func (fc *FhirClient) UpdateTask(updateResource *r4b.Task) (*r4b.Task, error) {
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

	var parsed r4b.Task
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateTask: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return Task or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchTask(patchResource *r4b.Task) (*r4b.Task, error) {
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

	var parsed r4b.Task
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchTask: %s", err)
		}
	}
	return &parsed, err
}

// delete Task and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTask(deleteResource *r4b.Task) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteTask given nil Task")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteTask given Task with nil Id (Id required to delete)")
	}
	return fc.DeleteTaskById(*deleteResource.Id)
}

// delete Task with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTaskById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "Task", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteTaskById: %s", err)
		}
	}
	return &parsed, err
}

// create TerminologyCapabilities, return id of created TerminologyCapabilities or OperationOutcome error
func (fc *FhirClient) CreateTerminologyCapabilities(createResource *r4b.TerminologyCapabilities) (*r4b.TerminologyCapabilities, error) {
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

	var parsed r4b.TerminologyCapabilities
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateTerminologyCapabilities: %s", err)
		}
	}
	return &parsed, err
}

// read TerminologyCapabilities from fhir server by id, return TerminologyCapabilities or OperationOutcome error
func (fc *FhirClient) ReadTerminologyCapabilities(id string) (*r4b.TerminologyCapabilities, error) {
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

	var parsed r4b.TerminologyCapabilities
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadTerminologyCapabilities: %s", err)
		}
	}
	return &parsed, err
}

// replace TerminologyCapabilities if exists on server, else create new TerminologyCapabilities with given id, return TerminologyCapabilities or OperationOutcome error
func (fc *FhirClient) UpdateTerminologyCapabilities(updateResource *r4b.TerminologyCapabilities) (*r4b.TerminologyCapabilities, error) {
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

	var parsed r4b.TerminologyCapabilities
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateTerminologyCapabilities: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return TerminologyCapabilities or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchTerminologyCapabilities(patchResource *r4b.TerminologyCapabilities) (*r4b.TerminologyCapabilities, error) {
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

	var parsed r4b.TerminologyCapabilities
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchTerminologyCapabilities: %s", err)
		}
	}
	return &parsed, err
}

// delete TerminologyCapabilities and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTerminologyCapabilities(deleteResource *r4b.TerminologyCapabilities) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteTerminologyCapabilities given nil TerminologyCapabilities")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteTerminologyCapabilities given TerminologyCapabilities with nil Id (Id required to delete)")
	}
	return fc.DeleteTerminologyCapabilitiesById(*deleteResource.Id)
}

// delete TerminologyCapabilities with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTerminologyCapabilitiesById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TerminologyCapabilities", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteTerminologyCapabilitiesById: %s", err)
		}
	}
	return &parsed, err
}

// create TestReport, return id of created TestReport or OperationOutcome error
func (fc *FhirClient) CreateTestReport(createResource *r4b.TestReport) (*r4b.TestReport, error) {
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

	var parsed r4b.TestReport
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateTestReport: %s", err)
		}
	}
	return &parsed, err
}

// read TestReport from fhir server by id, return TestReport or OperationOutcome error
func (fc *FhirClient) ReadTestReport(id string) (*r4b.TestReport, error) {
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

	var parsed r4b.TestReport
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadTestReport: %s", err)
		}
	}
	return &parsed, err
}

// replace TestReport if exists on server, else create new TestReport with given id, return TestReport or OperationOutcome error
func (fc *FhirClient) UpdateTestReport(updateResource *r4b.TestReport) (*r4b.TestReport, error) {
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

	var parsed r4b.TestReport
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateTestReport: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return TestReport or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchTestReport(patchResource *r4b.TestReport) (*r4b.TestReport, error) {
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

	var parsed r4b.TestReport
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchTestReport: %s", err)
		}
	}
	return &parsed, err
}

// delete TestReport and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTestReport(deleteResource *r4b.TestReport) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteTestReport given nil TestReport")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteTestReport given TestReport with nil Id (Id required to delete)")
	}
	return fc.DeleteTestReportById(*deleteResource.Id)
}

// delete TestReport with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTestReportById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestReport", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteTestReportById: %s", err)
		}
	}
	return &parsed, err
}

// create TestScript, return id of created TestScript or OperationOutcome error
func (fc *FhirClient) CreateTestScript(createResource *r4b.TestScript) (*r4b.TestScript, error) {
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

	var parsed r4b.TestScript
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateTestScript: %s", err)
		}
	}
	return &parsed, err
}

// read TestScript from fhir server by id, return TestScript or OperationOutcome error
func (fc *FhirClient) ReadTestScript(id string) (*r4b.TestScript, error) {
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

	var parsed r4b.TestScript
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadTestScript: %s", err)
		}
	}
	return &parsed, err
}

// replace TestScript if exists on server, else create new TestScript with given id, return TestScript or OperationOutcome error
func (fc *FhirClient) UpdateTestScript(updateResource *r4b.TestScript) (*r4b.TestScript, error) {
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

	var parsed r4b.TestScript
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateTestScript: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return TestScript or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchTestScript(patchResource *r4b.TestScript) (*r4b.TestScript, error) {
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

	var parsed r4b.TestScript
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchTestScript: %s", err)
		}
	}
	return &parsed, err
}

// delete TestScript and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTestScript(deleteResource *r4b.TestScript) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteTestScript given nil TestScript")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteTestScript given TestScript with nil Id (Id required to delete)")
	}
	return fc.DeleteTestScriptById(*deleteResource.Id)
}

// delete TestScript with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteTestScriptById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "TestScript", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteTestScriptById: %s", err)
		}
	}
	return &parsed, err
}

// create ValueSet, return id of created ValueSet or OperationOutcome error
func (fc *FhirClient) CreateValueSet(createResource *r4b.ValueSet) (*r4b.ValueSet, error) {
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

	var parsed r4b.ValueSet
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateValueSet: %s", err)
		}
	}
	return &parsed, err
}

// read ValueSet from fhir server by id, return ValueSet or OperationOutcome error
func (fc *FhirClient) ReadValueSet(id string) (*r4b.ValueSet, error) {
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

	var parsed r4b.ValueSet
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadValueSet: %s", err)
		}
	}
	return &parsed, err
}

// replace ValueSet if exists on server, else create new ValueSet with given id, return ValueSet or OperationOutcome error
func (fc *FhirClient) UpdateValueSet(updateResource *r4b.ValueSet) (*r4b.ValueSet, error) {
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

	var parsed r4b.ValueSet
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateValueSet: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return ValueSet or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchValueSet(patchResource *r4b.ValueSet) (*r4b.ValueSet, error) {
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

	var parsed r4b.ValueSet
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchValueSet: %s", err)
		}
	}
	return &parsed, err
}

// delete ValueSet and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteValueSet(deleteResource *r4b.ValueSet) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteValueSet given nil ValueSet")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteValueSet given ValueSet with nil Id (Id required to delete)")
	}
	return fc.DeleteValueSetById(*deleteResource.Id)
}

// delete ValueSet with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteValueSetById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "ValueSet", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteValueSetById: %s", err)
		}
	}
	return &parsed, err
}

// create VerificationResult, return id of created VerificationResult or OperationOutcome error
func (fc *FhirClient) CreateVerificationResult(createResource *r4b.VerificationResult) (*r4b.VerificationResult, error) {
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

	var parsed r4b.VerificationResult
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateVerificationResult: %s", err)
		}
	}
	return &parsed, err
}

// read VerificationResult from fhir server by id, return VerificationResult or OperationOutcome error
func (fc *FhirClient) ReadVerificationResult(id string) (*r4b.VerificationResult, error) {
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

	var parsed r4b.VerificationResult
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadVerificationResult: %s", err)
		}
	}
	return &parsed, err
}

// replace VerificationResult if exists on server, else create new VerificationResult with given id, return VerificationResult or OperationOutcome error
func (fc *FhirClient) UpdateVerificationResult(updateResource *r4b.VerificationResult) (*r4b.VerificationResult, error) {
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

	var parsed r4b.VerificationResult
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateVerificationResult: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return VerificationResult or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchVerificationResult(patchResource *r4b.VerificationResult) (*r4b.VerificationResult, error) {
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

	var parsed r4b.VerificationResult
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchVerificationResult: %s", err)
		}
	}
	return &parsed, err
}

// delete VerificationResult and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteVerificationResult(deleteResource *r4b.VerificationResult) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteVerificationResult given nil VerificationResult")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteVerificationResult given VerificationResult with nil Id (Id required to delete)")
	}
	return fc.DeleteVerificationResultById(*deleteResource.Id)
}

// delete VerificationResult with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteVerificationResultById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "VerificationResult", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteVerificationResultById: %s", err)
		}
	}
	return &parsed, err
}

// create VisionPrescription, return id of created VisionPrescription or OperationOutcome error
func (fc *FhirClient) CreateVisionPrescription(createResource *r4b.VisionPrescription) (*r4b.VisionPrescription, error) {
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

	var parsed r4b.VisionPrescription
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("CreateVisionPrescription: %s", err)
		}
	}
	return &parsed, err
}

// read VisionPrescription from fhir server by id, return VisionPrescription or OperationOutcome error
func (fc *FhirClient) ReadVisionPrescription(id string) (*r4b.VisionPrescription, error) {
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

	var parsed r4b.VisionPrescription
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("ReadVisionPrescription: %s", err)
		}
	}
	return &parsed, err
}

// replace VisionPrescription if exists on server, else create new VisionPrescription with given id, return VisionPrescription or OperationOutcome error
func (fc *FhirClient) UpdateVisionPrescription(updateResource *r4b.VisionPrescription) (*r4b.VisionPrescription, error) {
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

	var parsed r4b.VisionPrescription
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("UpdateVisionPrescription: %s", err)
		}
	}
	return &parsed, err
}

// add or replace fields in existing resource on server, return VisionPrescription or error OperationOutcome
//
// uses normal json patch, not fhirpath
func (fc *FhirClient) PatchVisionPrescription(patchResource *r4b.VisionPrescription) (*r4b.VisionPrescription, error) {
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

	var parsed r4b.VisionPrescription
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("PatchVisionPrescription: %s", err)
		}
	}
	return &parsed, err
}

// delete VisionPrescription and all resources that reference it, return OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteVisionPrescription(deleteResource *r4b.VisionPrescription) (*r4b.OperationOutcome, error) {
	if deleteResource == nil {
		return nil, errors.New("DeleteVisionPrescription given nil VisionPrescription")
	}
	if deleteResource.Id == nil {
		return nil, errors.New("DeleteVisionPrescription given VisionPrescription with nil Id (Id required to delete)")
	}
	return fc.DeleteVisionPrescriptionById(*deleteResource.Id)
}

// delete VisionPrescription with given id and all resources that reference it, return normal OperationOutcome or error OperationOutcome
func (fc *FhirClient) DeleteVisionPrescriptionById(id string) (*r4b.OperationOutcome, error) {
	reqUrl, err := url.JoinPath(fc.BaseUrl, "VisionPrescription", id)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?_cascade=delete" //might want to use url or something idk
	req, err := http.NewRequest(http.MethodDelete, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	var parsed r4b.OperationOutcome
	err = getFhirResourceOrError(fc, req, &parsed)
	if err != nil {
		if _, ok := err.(*r4.OperationOutcome); !ok {
			err = fmt.Errorf("DeleteVisionPrescriptionById: %s", err)
		}
	}
	return &parsed, err
}

// when parsing, check which fhir resource type (AllergyIntolerance, Patient, OperationOutcome, etc)
type checkType struct {
	ResourceType string `json:"resourceType"`
}

// get a fhir resource of type you asked for,
// or an operationoutcome error from server response,
// or some other string error if server misbehaving eg down
func getFhirResourceOrError(fc *FhirClient, req *http.Request, retResource r4b.FhirResource) error {
	req.Header.Set("Accept", "application/fhir+json")
	resp, err := fc.Do(req)
	if err != nil {
		return fmt.Errorf("getFhirResourceOrError, making req got err %s", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("getFhirResourceOrError: reading body got err %s", err)
	}
	defer resp.Body.Close()
	var ct checkType
	err = json.Unmarshal(body, &ct)
	if err != nil {
		return fmt.Errorf("getFhirResourceOrError: response body does not have resourceType: %s", string(body))
	}
	if ct.ResourceType == retResource.ResourceType() {
		err = json.Unmarshal(body, &retResource)
		if err == nil {
			//managed to unmarshal server response into desired resource type
			return nil
		} else {
			return fmt.Errorf("getFhirResourceOrError: %s resourceType could not unmarshal from body: %s", ct.ResourceType, string(body))
		}
	} else if ct.ResourceType == "OperationOutcome" {
		var oo r4b.OperationOutcome
		err = json.Unmarshal(body, &oo)
		if err == nil {
			return oo
		} else {
			return fmt.Errorf("getFhirResourceOrError: could not unmarshal OperationOutcome from body: %s", string(body))
		}
	}
	return fmt.Errorf("getFhirResourceOrError: response body has no fhir resource type: %s", string(body))
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
