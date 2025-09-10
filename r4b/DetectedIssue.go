package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/DetectedIssue
type DetectedIssue struct {
	Id                 *string                   `json:"id,omitempty"`
	Meta               *Meta                     `json:"meta,omitempty"`
	ImplicitRules      *string                   `json:"implicitRules,omitempty"`
	Language           *string                   `json:"language,omitempty"`
	Text               *Narrative                `json:"text,omitempty"`
	Contained          []Resource                `json:"contained,omitempty"`
	Extension          []Extension               `json:"extension,omitempty"`
	ModifierExtension  []Extension               `json:"modifierExtension,omitempty"`
	Identifier         []Identifier              `json:"identifier,omitempty"`
	Status             string                    `json:"status"`
	Code               *CodeableConcept          `json:"code,omitempty"`
	Severity           *string                   `json:"severity,omitempty"`
	Patient            *Reference                `json:"patient,omitempty"`
	IdentifiedDateTime *string                   `json:"identifiedDateTime,omitempty"`
	IdentifiedPeriod   *Period                   `json:"identifiedPeriod,omitempty"`
	Author             *Reference                `json:"author,omitempty"`
	Implicated         []Reference               `json:"implicated,omitempty"`
	Evidence           []DetectedIssueEvidence   `json:"evidence,omitempty"`
	Detail             *string                   `json:"detail,omitempty"`
	Reference          *string                   `json:"reference,omitempty"`
	Mitigation         []DetectedIssueMitigation `json:"mitigation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/DetectedIssue
type DetectedIssueEvidence struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `json:"code,omitempty"`
	Detail            []Reference       `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/DetectedIssue
type DetectedIssueMitigation struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Action            CodeableConcept `json:"action"`
	Date              *string         `json:"date,omitempty"`
	Author            *Reference      `json:"author,omitempty"`
}

type OtherDetectedIssue DetectedIssue

// on convert struct to json, automatically add resourceType=DetectedIssue
func (r DetectedIssue) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDetectedIssue
		ResourceType string `json:"resourceType"`
	}{
		OtherDetectedIssue: OtherDetectedIssue(r),
		ResourceType:       "DetectedIssue",
	})
}
func (r DetectedIssue) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "DetectedIssue/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "DetectedIssue"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *DetectedIssue) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSObservation_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DetectedIssue) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *DetectedIssue) T_Severity(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDetectedissue_severity

	if resource == nil {
		return CodeSelect("severity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("severity", resource.Severity, optionsValueSet, htmlAttrs)
}
func (resource *DetectedIssue) T_IdentifiedDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("identifiedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("identifiedDateTime", resource.IdentifiedDateTime, htmlAttrs)
}
func (resource *DetectedIssue) T_Detail(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("detail", nil, htmlAttrs)
	}
	return StringInput("detail", resource.Detail, htmlAttrs)
}
func (resource *DetectedIssue) T_Reference(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("reference", nil, htmlAttrs)
	}
	return StringInput("reference", resource.Reference, htmlAttrs)
}
func (resource *DetectedIssue) T_EvidenceCode(numEvidence int, numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvidence >= len(resource.Evidence) || numCode >= len(resource.Evidence[numEvidence].Code) {
		return CodeableConceptSelect("evidence["+strconv.Itoa(numEvidence)+"].code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("evidence["+strconv.Itoa(numEvidence)+"].code["+strconv.Itoa(numCode)+"]", &resource.Evidence[numEvidence].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *DetectedIssue) T_MitigationAction(numMitigation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMitigation >= len(resource.Mitigation) {
		return CodeableConceptSelect("mitigation["+strconv.Itoa(numMitigation)+"].action", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("mitigation["+strconv.Itoa(numMitigation)+"].action", &resource.Mitigation[numMitigation].Action, optionsValueSet, htmlAttrs)
}
func (resource *DetectedIssue) T_MitigationDate(numMitigation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMitigation >= len(resource.Mitigation) {
		return DateTimeInput("mitigation["+strconv.Itoa(numMitigation)+"].date", nil, htmlAttrs)
	}
	return DateTimeInput("mitigation["+strconv.Itoa(numMitigation)+"].date", resource.Mitigation[numMitigation].Date, htmlAttrs)
}
