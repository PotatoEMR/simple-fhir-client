package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	IdentifiedDateTime *time.Time                `json:"identifiedDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	Date              *time.Time      `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *DetectedIssue) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSObservation_status

	if resource == nil {
		return CodeSelect("DetectedIssue.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DetectedIssue.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DetectedIssue) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("DetectedIssue.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DetectedIssue.Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *DetectedIssue) T_Severity(htmlAttrs string) templ.Component {
	optionsValueSet := VSDetectedissue_severity

	if resource == nil {
		return CodeSelect("DetectedIssue.Severity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DetectedIssue.Severity", resource.Severity, optionsValueSet, htmlAttrs)
}
func (resource *DetectedIssue) T_IdentifiedDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("DetectedIssue.IdentifiedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("DetectedIssue.IdentifiedDateTime", resource.IdentifiedDateTime, htmlAttrs)
}
func (resource *DetectedIssue) T_Detail(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DetectedIssue.Detail", nil, htmlAttrs)
	}
	return StringInput("DetectedIssue.Detail", resource.Detail, htmlAttrs)
}
func (resource *DetectedIssue) T_Reference(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DetectedIssue.Reference", nil, htmlAttrs)
	}
	return StringInput("DetectedIssue.Reference", resource.Reference, htmlAttrs)
}
func (resource *DetectedIssue) T_EvidenceCode(numEvidence int, numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEvidence >= len(resource.Evidence) || numCode >= len(resource.Evidence[numEvidence].Code) {
		return CodeableConceptSelect("DetectedIssue.Evidence["+strconv.Itoa(numEvidence)+"].Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DetectedIssue.Evidence["+strconv.Itoa(numEvidence)+"].Code["+strconv.Itoa(numCode)+"]", &resource.Evidence[numEvidence].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *DetectedIssue) T_MitigationAction(numMitigation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMitigation >= len(resource.Mitigation) {
		return CodeableConceptSelect("DetectedIssue.Mitigation["+strconv.Itoa(numMitigation)+"].Action", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DetectedIssue.Mitigation["+strconv.Itoa(numMitigation)+"].Action", &resource.Mitigation[numMitigation].Action, optionsValueSet, htmlAttrs)
}
func (resource *DetectedIssue) T_MitigationDate(numMitigation int, htmlAttrs string) templ.Component {
	if resource == nil || numMitigation >= len(resource.Mitigation) {
		return DateTimeInput("DetectedIssue.Mitigation["+strconv.Itoa(numMitigation)+"].Date", nil, htmlAttrs)
	}
	return DateTimeInput("DetectedIssue.Mitigation["+strconv.Itoa(numMitigation)+"].Date", resource.Mitigation[numMitigation].Date, htmlAttrs)
}
