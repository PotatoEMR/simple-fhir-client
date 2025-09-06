package r4b

//generated with command go run ./bultaoreune -nodownload
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

func (resource *DetectedIssue) T_Id() templ.Component {

	if resource == nil {
		return StringInput("DetectedIssue.Id", nil)
	}
	return StringInput("DetectedIssue.Id", resource.Id)
}
func (resource *DetectedIssue) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("DetectedIssue.ImplicitRules", nil)
	}
	return StringInput("DetectedIssue.ImplicitRules", resource.ImplicitRules)
}
func (resource *DetectedIssue) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("DetectedIssue.Language", nil, optionsValueSet)
	}
	return CodeSelect("DetectedIssue.Language", resource.Language, optionsValueSet)
}
func (resource *DetectedIssue) T_Status() templ.Component {
	optionsValueSet := VSObservation_status

	if resource == nil {
		return CodeSelect("DetectedIssue.Status", nil, optionsValueSet)
	}
	return CodeSelect("DetectedIssue.Status", &resource.Status, optionsValueSet)
}
func (resource *DetectedIssue) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DetectedIssue.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DetectedIssue.Code", resource.Code, optionsValueSet)
}
func (resource *DetectedIssue) T_Severity() templ.Component {
	optionsValueSet := VSDetectedissue_severity

	if resource == nil {
		return CodeSelect("DetectedIssue.Severity", nil, optionsValueSet)
	}
	return CodeSelect("DetectedIssue.Severity", resource.Severity, optionsValueSet)
}
func (resource *DetectedIssue) T_Detail() templ.Component {

	if resource == nil {
		return StringInput("DetectedIssue.Detail", nil)
	}
	return StringInput("DetectedIssue.Detail", resource.Detail)
}
func (resource *DetectedIssue) T_Reference() templ.Component {

	if resource == nil {
		return StringInput("DetectedIssue.Reference", nil)
	}
	return StringInput("DetectedIssue.Reference", resource.Reference)
}
func (resource *DetectedIssue) T_EvidenceId(numEvidence int) templ.Component {

	if resource == nil || len(resource.Evidence) >= numEvidence {
		return StringInput("DetectedIssue.Evidence["+strconv.Itoa(numEvidence)+"].Id", nil)
	}
	return StringInput("DetectedIssue.Evidence["+strconv.Itoa(numEvidence)+"].Id", resource.Evidence[numEvidence].Id)
}
func (resource *DetectedIssue) T_EvidenceCode(numEvidence int, numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Evidence) >= numEvidence || len(resource.Evidence[numEvidence].Code) >= numCode {
		return CodeableConceptSelect("DetectedIssue.Evidence["+strconv.Itoa(numEvidence)+"].Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DetectedIssue.Evidence["+strconv.Itoa(numEvidence)+"].Code["+strconv.Itoa(numCode)+"]", &resource.Evidence[numEvidence].Code[numCode], optionsValueSet)
}
func (resource *DetectedIssue) T_MitigationId(numMitigation int) templ.Component {

	if resource == nil || len(resource.Mitigation) >= numMitigation {
		return StringInput("DetectedIssue.Mitigation["+strconv.Itoa(numMitigation)+"].Id", nil)
	}
	return StringInput("DetectedIssue.Mitigation["+strconv.Itoa(numMitigation)+"].Id", resource.Mitigation[numMitigation].Id)
}
func (resource *DetectedIssue) T_MitigationAction(numMitigation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Mitigation) >= numMitigation {
		return CodeableConceptSelect("DetectedIssue.Mitigation["+strconv.Itoa(numMitigation)+"].Action", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DetectedIssue.Mitigation["+strconv.Itoa(numMitigation)+"].Action", &resource.Mitigation[numMitigation].Action, optionsValueSet)
}
func (resource *DetectedIssue) T_MitigationDate(numMitigation int) templ.Component {

	if resource == nil || len(resource.Mitigation) >= numMitigation {
		return StringInput("DetectedIssue.Mitigation["+strconv.Itoa(numMitigation)+"].Date", nil)
	}
	return StringInput("DetectedIssue.Mitigation["+strconv.Itoa(numMitigation)+"].Date", resource.Mitigation[numMitigation].Date)
}
