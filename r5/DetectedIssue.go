package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/DetectedIssue
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
	Category           []CodeableConcept         `json:"category,omitempty"`
	Code               *CodeableConcept          `json:"code,omitempty"`
	Severity           *string                   `json:"severity,omitempty"`
	Subject            *Reference                `json:"subject,omitempty"`
	Encounter          *Reference                `json:"encounter,omitempty"`
	IdentifiedDateTime *string                   `json:"identifiedDateTime"`
	IdentifiedPeriod   *Period                   `json:"identifiedPeriod"`
	Author             *Reference                `json:"author,omitempty"`
	Implicated         []Reference               `json:"implicated,omitempty"`
	Evidence           []DetectedIssueEvidence   `json:"evidence,omitempty"`
	Detail             *string                   `json:"detail,omitempty"`
	Reference          *string                   `json:"reference,omitempty"`
	Mitigation         []DetectedIssueMitigation `json:"mitigation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DetectedIssue
type DetectedIssueEvidence struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `json:"code,omitempty"`
	Detail            []Reference       `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DetectedIssue
type DetectedIssueMitigation struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Action            CodeableConcept `json:"action"`
	Date              *string         `json:"date,omitempty"`
	Author            *Reference      `json:"author,omitempty"`
	Note              []Annotation    `json:"note,omitempty"`
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

func (resource *DetectedIssue) DetectedIssueLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *DetectedIssue) DetectedIssueStatus() templ.Component {
	optionsValueSet := VSDetectedissue_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *DetectedIssue) DetectedIssueSeverity() templ.Component {
	optionsValueSet := VSDetectedissue_severity
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Severity
	}
	return CodeSelect("severity", currentVal, optionsValueSet)
}
