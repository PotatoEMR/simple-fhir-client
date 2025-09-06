package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Condition
type Condition struct {
	Id                 *string                `json:"id,omitempty"`
	Meta               *Meta                  `json:"meta,omitempty"`
	ImplicitRules      *string                `json:"implicitRules,omitempty"`
	Language           *string                `json:"language,omitempty"`
	Text               *Narrative             `json:"text,omitempty"`
	Contained          []Resource             `json:"contained,omitempty"`
	Extension          []Extension            `json:"extension,omitempty"`
	ModifierExtension  []Extension            `json:"modifierExtension,omitempty"`
	Identifier         []Identifier           `json:"identifier,omitempty"`
	ClinicalStatus     CodeableConcept        `json:"clinicalStatus"`
	VerificationStatus *CodeableConcept       `json:"verificationStatus,omitempty"`
	Category           []CodeableConcept      `json:"category,omitempty"`
	Severity           *CodeableConcept       `json:"severity,omitempty"`
	Code               *CodeableConcept       `json:"code,omitempty"`
	BodySite           []CodeableConcept      `json:"bodySite,omitempty"`
	Subject            Reference              `json:"subject"`
	Encounter          *Reference             `json:"encounter,omitempty"`
	OnsetDateTime      *string                `json:"onsetDateTime,omitempty"`
	OnsetAge           *Age                   `json:"onsetAge,omitempty"`
	OnsetPeriod        *Period                `json:"onsetPeriod,omitempty"`
	OnsetRange         *Range                 `json:"onsetRange,omitempty"`
	OnsetString        *string                `json:"onsetString,omitempty"`
	AbatementDateTime  *string                `json:"abatementDateTime,omitempty"`
	AbatementAge       *Age                   `json:"abatementAge,omitempty"`
	AbatementPeriod    *Period                `json:"abatementPeriod,omitempty"`
	AbatementRange     *Range                 `json:"abatementRange,omitempty"`
	AbatementString    *string                `json:"abatementString,omitempty"`
	RecordedDate       *string                `json:"recordedDate,omitempty"`
	Participant        []ConditionParticipant `json:"participant,omitempty"`
	Stage              []ConditionStage       `json:"stage,omitempty"`
	Evidence           []CodeableReference    `json:"evidence,omitempty"`
	Note               []Annotation           `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Condition
type ConditionParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Condition
type ConditionStage struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Summary           *CodeableConcept `json:"summary,omitempty"`
	Assessment        []Reference      `json:"assessment,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
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
func (resource *Condition) T_ClinicalStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Condition.ClinicalStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Condition.ClinicalStatus", &resource.ClinicalStatus, optionsValueSet)
}
func (resource *Condition) T_VerificationStatus(optionsValueSet []Coding) templ.Component {

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
func (resource *Condition) T_ParticipantId(numParticipant int) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return StringInput("Condition.Participant["+strconv.Itoa(numParticipant)+"].Id", nil)
	}
	return StringInput("Condition.Participant["+strconv.Itoa(numParticipant)+"].Id", resource.Participant[numParticipant].Id)
}
func (resource *Condition) T_ParticipantFunction(numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return CodeableConceptSelect("Condition.Participant["+strconv.Itoa(numParticipant)+"].Function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Condition.Participant["+strconv.Itoa(numParticipant)+"].Function", resource.Participant[numParticipant].Function, optionsValueSet)
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
