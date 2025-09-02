package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/DetectedIssue
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
	IdentifiedDateTime *string                   `json:"identifiedDateTime"`
	IdentifiedPeriod   *Period                   `json:"identifiedPeriod"`
	Author             *Reference                `json:"author,omitempty"`
	Implicated         []Reference               `json:"implicated,omitempty"`
	Evidence           []DetectedIssueEvidence   `json:"evidence,omitempty"`
	Detail             *string                   `json:"detail,omitempty"`
	Reference          *string                   `json:"reference,omitempty"`
	Mitigation         []DetectedIssueMitigation `json:"mitigation,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/DetectedIssue
type DetectedIssueEvidence struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `json:"code,omitempty"`
	Detail            []Reference       `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/DetectedIssue
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

func (resource *DetectedIssue) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *DetectedIssue) T_Status() templ.Component {
	optionsValueSet := VSObservation_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *DetectedIssue) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet)
}
func (resource *DetectedIssue) T_Severity() templ.Component {
	optionsValueSet := VSDetectedissue_severity

	if resource == nil {
		return CodeSelect("severity", nil, optionsValueSet)
	}
	return CodeSelect("severity", resource.Severity, optionsValueSet)
}
func (resource *DetectedIssue) T_EvidenceCode(numEvidence int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Evidence) >= numEvidence {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Evidence[numEvidence].Code[0], optionsValueSet)
}
func (resource *DetectedIssue) T_MitigationAction(numMitigation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Mitigation) >= numMitigation {
		return CodeableConceptSelect("action", nil, optionsValueSet)
	}
	return CodeableConceptSelect("action", &resource.Mitigation[numMitigation].Action, optionsValueSet)
}
