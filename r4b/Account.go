package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Account
type Account struct {
	Id                *string            `json:"id,omitempty"`
	Meta              *Meta              `json:"meta,omitempty"`
	ImplicitRules     *string            `json:"implicitRules,omitempty"`
	Language          *string            `json:"language,omitempty"`
	Text              *Narrative         `json:"text,omitempty"`
	Contained         []Resource         `json:"contained,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Identifier        []Identifier       `json:"identifier,omitempty"`
	Status            string             `json:"status"`
	Type              *CodeableConcept   `json:"type,omitempty"`
	Name              *string            `json:"name,omitempty"`
	Subject           []Reference        `json:"subject,omitempty"`
	ServicePeriod     *Period            `json:"servicePeriod,omitempty"`
	Coverage          []AccountCoverage  `json:"coverage,omitempty"`
	Owner             *Reference         `json:"owner,omitempty"`
	Description       *string            `json:"description,omitempty"`
	Guarantor         []AccountGuarantor `json:"guarantor,omitempty"`
	PartOf            *Reference         `json:"partOf,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Account
type AccountCoverage struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Coverage          Reference   `json:"coverage"`
	Priority          *int        `json:"priority,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Account
type AccountGuarantor struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Party             Reference   `json:"party"`
	OnHold            *bool       `json:"onHold,omitempty"`
	Period            *Period     `json:"period,omitempty"`
}

type OtherAccount Account

// on convert struct to json, automatically add resourceType=Account
func (r Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAccount
		ResourceType string `json:"resourceType"`
	}{
		OtherAccount: OtherAccount(r),
		ResourceType: "Account",
	})
}
func (r Account) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Account/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Account"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Account) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAccount_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Account) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Account) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *Account) T_Subject(frs []FhirResource, numSubject int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubject >= len(resource.Subject) {
		return ReferenceInput(frs, "subject["+strconv.Itoa(numSubject)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject["+strconv.Itoa(numSubject)+"]", &resource.Subject[numSubject], htmlAttrs)
}
func (resource *Account) T_ServicePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("servicePeriod", nil, htmlAttrs)
	}
	return PeriodInput("servicePeriod", resource.ServicePeriod, htmlAttrs)
}
func (resource *Account) T_Owner(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "owner", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "owner", resource.Owner, htmlAttrs)
}
func (resource *Account) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *Account) T_PartOf(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "partOf", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf", resource.PartOf, htmlAttrs)
}
func (resource *Account) T_CoverageCoverage(frs []FhirResource, numCoverage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCoverage >= len(resource.Coverage) {
		return ReferenceInput(frs, "coverage["+strconv.Itoa(numCoverage)+"].coverage", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "coverage["+strconv.Itoa(numCoverage)+"].coverage", &resource.Coverage[numCoverage].Coverage, htmlAttrs)
}
func (resource *Account) T_CoveragePriority(numCoverage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCoverage >= len(resource.Coverage) {
		return IntInput("coverage["+strconv.Itoa(numCoverage)+"].priority", nil, htmlAttrs)
	}
	return IntInput("coverage["+strconv.Itoa(numCoverage)+"].priority", resource.Coverage[numCoverage].Priority, htmlAttrs)
}
func (resource *Account) T_GuarantorParty(frs []FhirResource, numGuarantor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGuarantor >= len(resource.Guarantor) {
		return ReferenceInput(frs, "guarantor["+strconv.Itoa(numGuarantor)+"].party", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "guarantor["+strconv.Itoa(numGuarantor)+"].party", &resource.Guarantor[numGuarantor].Party, htmlAttrs)
}
func (resource *Account) T_GuarantorOnHold(numGuarantor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGuarantor >= len(resource.Guarantor) {
		return BoolInput("guarantor["+strconv.Itoa(numGuarantor)+"].onHold", nil, htmlAttrs)
	}
	return BoolInput("guarantor["+strconv.Itoa(numGuarantor)+"].onHold", resource.Guarantor[numGuarantor].OnHold, htmlAttrs)
}
func (resource *Account) T_GuarantorPeriod(numGuarantor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGuarantor >= len(resource.Guarantor) {
		return PeriodInput("guarantor["+strconv.Itoa(numGuarantor)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("guarantor["+strconv.Itoa(numGuarantor)+"].period", resource.Guarantor[numGuarantor].Period, htmlAttrs)
}
