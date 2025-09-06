package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Condition
type Condition struct {
	Id                 *string             `json:"id,omitempty"`
	Meta               *Meta               `json:"meta,omitempty"`
	ImplicitRules      *string             `json:"implicitRules,omitempty"`
	Language           *string             `json:"language,omitempty"`
	Text               *Narrative          `json:"text,omitempty"`
	Contained          []Resource          `json:"contained,omitempty"`
	Extension          []Extension         `json:"extension,omitempty"`
	ModifierExtension  []Extension         `json:"modifierExtension,omitempty"`
	Identifier         []Identifier        `json:"identifier,omitempty"`
	ClinicalStatus     *CodeableConcept    `json:"clinicalStatus,omitempty"`
	VerificationStatus *CodeableConcept    `json:"verificationStatus,omitempty"`
	Category           []CodeableConcept   `json:"category,omitempty"`
	Severity           *CodeableConcept    `json:"severity,omitempty"`
	Code               *CodeableConcept    `json:"code,omitempty"`
	BodySite           []CodeableConcept   `json:"bodySite,omitempty"`
	Subject            Reference           `json:"subject"`
	Encounter          *Reference          `json:"encounter,omitempty"`
	OnsetDateTime      *string             `json:"onsetDateTime,omitempty"`
	OnsetAge           *Age                `json:"onsetAge,omitempty"`
	OnsetPeriod        *Period             `json:"onsetPeriod,omitempty"`
	OnsetRange         *Range              `json:"onsetRange,omitempty"`
	OnsetString        *string             `json:"onsetString,omitempty"`
	AbatementDateTime  *string             `json:"abatementDateTime,omitempty"`
	AbatementAge       *Age                `json:"abatementAge,omitempty"`
	AbatementPeriod    *Period             `json:"abatementPeriod,omitempty"`
	AbatementRange     *Range              `json:"abatementRange,omitempty"`
	AbatementString    *string             `json:"abatementString,omitempty"`
	RecordedDate       *string             `json:"recordedDate,omitempty"`
	Recorder           *Reference          `json:"recorder,omitempty"`
	Asserter           *Reference          `json:"asserter,omitempty"`
	Stage              []ConditionStage    `json:"stage,omitempty"`
	Evidence           []ConditionEvidence `json:"evidence,omitempty"`
	Note               []Annotation        `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Condition
type ConditionStage struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Summary           *CodeableConcept `json:"summary,omitempty"`
	Assessment        []Reference      `json:"assessment,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Condition
type ConditionEvidence struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `json:"code,omitempty"`
	Detail            []Reference       `json:"detail,omitempty"`
}

type OtherCondition Condition

// on convert struct to json, automatically add resourceType=Condition
func (r Condition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCondition
		ResourceType string `json:"resourceType"`
	}{
		OtherCondition: OtherCondition(r),
		ResourceType:   "Condition",
	})
}

func (resource *Condition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Condition.Id", nil)
	}
	return StringInput("Condition.Id", resource.Id)
}
func (resource *Condition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Condition.ImplicitRules", nil)
	}
	return StringInput("Condition.ImplicitRules", resource.ImplicitRules)
}
func (resource *Condition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Condition.Language", nil, optionsValueSet)
	}
	return CodeSelect("Condition.Language", resource.Language, optionsValueSet)
}
func (resource *Condition) T_ClinicalStatus() templ.Component {
	optionsValueSet := VSCondition_clinical

	if resource == nil {
		return CodeableConceptSelect("Condition.ClinicalStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Condition.ClinicalStatus", resource.ClinicalStatus, optionsValueSet)
}
func (resource *Condition) T_VerificationStatus() templ.Component {
	optionsValueSet := VSCondition_ver_status

	if resource == nil {
		return CodeableConceptSelect("Condition.VerificationStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Condition.VerificationStatus", resource.VerificationStatus, optionsValueSet)
}
func (resource *Condition) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("Condition.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Condition.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *Condition) T_Severity(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Condition.Severity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Condition.Severity", resource.Severity, optionsValueSet)
}
func (resource *Condition) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Condition.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Condition.Code", resource.Code, optionsValueSet)
}
func (resource *Condition) T_BodySite(numBodySite int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.BodySite) >= numBodySite {
		return CodeableConceptSelect("Condition.BodySite["+strconv.Itoa(numBodySite)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Condition.BodySite["+strconv.Itoa(numBodySite)+"]", &resource.BodySite[numBodySite], optionsValueSet)
}
func (resource *Condition) T_RecordedDate() templ.Component {

	if resource == nil {
		return StringInput("Condition.RecordedDate", nil)
	}
	return StringInput("Condition.RecordedDate", resource.RecordedDate)
}
func (resource *Condition) T_StageId(numStage int) templ.Component {

	if resource == nil || len(resource.Stage) >= numStage {
		return StringInput("Condition.Stage["+strconv.Itoa(numStage)+"].Id", nil)
	}
	return StringInput("Condition.Stage["+strconv.Itoa(numStage)+"].Id", resource.Stage[numStage].Id)
}
func (resource *Condition) T_StageSummary(numStage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Stage) >= numStage {
		return CodeableConceptSelect("Condition.Stage["+strconv.Itoa(numStage)+"].Summary", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Condition.Stage["+strconv.Itoa(numStage)+"].Summary", resource.Stage[numStage].Summary, optionsValueSet)
}
func (resource *Condition) T_StageType(numStage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Stage) >= numStage {
		return CodeableConceptSelect("Condition.Stage["+strconv.Itoa(numStage)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Condition.Stage["+strconv.Itoa(numStage)+"].Type", resource.Stage[numStage].Type, optionsValueSet)
}
func (resource *Condition) T_EvidenceId(numEvidence int) templ.Component {

	if resource == nil || len(resource.Evidence) >= numEvidence {
		return StringInput("Condition.Evidence["+strconv.Itoa(numEvidence)+"].Id", nil)
	}
	return StringInput("Condition.Evidence["+strconv.Itoa(numEvidence)+"].Id", resource.Evidence[numEvidence].Id)
}
func (resource *Condition) T_EvidenceCode(numEvidence int, numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Evidence) >= numEvidence || len(resource.Evidence[numEvidence].Code) >= numCode {
		return CodeableConceptSelect("Condition.Evidence["+strconv.Itoa(numEvidence)+"].Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Condition.Evidence["+strconv.Itoa(numEvidence)+"].Code["+strconv.Itoa(numCode)+"]", &resource.Evidence[numEvidence].Code[numCode], optionsValueSet)
}
