package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/AllergyIntolerance
type AllergyIntolerance struct {
	Id                 *string                         `json:"id,omitempty"`
	Meta               *Meta                           `json:"meta,omitempty"`
	ImplicitRules      *string                         `json:"implicitRules,omitempty"`
	Language           *string                         `json:"language,omitempty"`
	Text               *Narrative                      `json:"text,omitempty"`
	Contained          []Resource                      `json:"contained,omitempty"`
	Extension          []Extension                     `json:"extension,omitempty"`
	ModifierExtension  []Extension                     `json:"modifierExtension,omitempty"`
	Identifier         []Identifier                    `json:"identifier,omitempty"`
	ClinicalStatus     *CodeableConcept                `json:"clinicalStatus,omitempty"`
	VerificationStatus *CodeableConcept                `json:"verificationStatus,omitempty"`
	Type               *CodeableConcept                `json:"type,omitempty"`
	Category           []string                        `json:"category,omitempty"`
	Criticality        *string                         `json:"criticality,omitempty"`
	Code               *CodeableConcept                `json:"code,omitempty"`
	Patient            Reference                       `json:"patient"`
	Encounter          *Reference                      `json:"encounter,omitempty"`
	OnsetDateTime      *string                         `json:"onsetDateTime,omitempty"`
	OnsetAge           *Age                            `json:"onsetAge,omitempty"`
	OnsetPeriod        *Period                         `json:"onsetPeriod,omitempty"`
	OnsetRange         *Range                          `json:"onsetRange,omitempty"`
	OnsetString        *string                         `json:"onsetString,omitempty"`
	RecordedDate       *string                         `json:"recordedDate,omitempty"`
	Participant        []AllergyIntoleranceParticipant `json:"participant,omitempty"`
	LastOccurrence     *string                         `json:"lastOccurrence,omitempty"`
	Note               []Annotation                    `json:"note,omitempty"`
	Reaction           []AllergyIntoleranceReaction    `json:"reaction,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AllergyIntolerance
type AllergyIntoleranceParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AllergyIntolerance
type AllergyIntoleranceReaction struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Substance         *CodeableConcept    `json:"substance,omitempty"`
	Manifestation     []CodeableReference `json:"manifestation"`
	Description       *string             `json:"description,omitempty"`
	Onset             *string             `json:"onset,omitempty"`
	Severity          *string             `json:"severity,omitempty"`
	ExposureRoute     *CodeableConcept    `json:"exposureRoute,omitempty"`
	Note              []Annotation        `json:"note,omitempty"`
}

type OtherAllergyIntolerance AllergyIntolerance

// on convert struct to json, automatically add resourceType=AllergyIntolerance
func (r AllergyIntolerance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAllergyIntolerance
		ResourceType string `json:"resourceType"`
	}{
		OtherAllergyIntolerance: OtherAllergyIntolerance(r),
		ResourceType:            "AllergyIntolerance",
	})
}

func (resource *AllergyIntolerance) T_Id() templ.Component {

	if resource == nil {
		return StringInput("AllergyIntolerance.Id", nil)
	}
	return StringInput("AllergyIntolerance.Id", resource.Id)
}
func (resource *AllergyIntolerance) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("AllergyIntolerance.ImplicitRules", nil)
	}
	return StringInput("AllergyIntolerance.ImplicitRules", resource.ImplicitRules)
}
func (resource *AllergyIntolerance) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("AllergyIntolerance.Language", nil, optionsValueSet)
	}
	return CodeSelect("AllergyIntolerance.Language", resource.Language, optionsValueSet)
}
func (resource *AllergyIntolerance) T_ClinicalStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("AllergyIntolerance.ClinicalStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AllergyIntolerance.ClinicalStatus", resource.ClinicalStatus, optionsValueSet)
}
func (resource *AllergyIntolerance) T_VerificationStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("AllergyIntolerance.VerificationStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AllergyIntolerance.VerificationStatus", resource.VerificationStatus, optionsValueSet)
}
func (resource *AllergyIntolerance) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("AllergyIntolerance.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AllergyIntolerance.Type", resource.Type, optionsValueSet)
}
func (resource *AllergyIntolerance) T_Category(numCategory int) templ.Component {
	optionsValueSet := VSAllergy_intolerance_category

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeSelect("AllergyIntolerance.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeSelect("AllergyIntolerance.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *AllergyIntolerance) T_Criticality() templ.Component {
	optionsValueSet := VSAllergy_intolerance_criticality

	if resource == nil {
		return CodeSelect("AllergyIntolerance.Criticality", nil, optionsValueSet)
	}
	return CodeSelect("AllergyIntolerance.Criticality", resource.Criticality, optionsValueSet)
}
func (resource *AllergyIntolerance) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("AllergyIntolerance.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AllergyIntolerance.Code", resource.Code, optionsValueSet)
}
func (resource *AllergyIntolerance) T_RecordedDate() templ.Component {

	if resource == nil {
		return StringInput("AllergyIntolerance.RecordedDate", nil)
	}
	return StringInput("AllergyIntolerance.RecordedDate", resource.RecordedDate)
}
func (resource *AllergyIntolerance) T_LastOccurrence() templ.Component {

	if resource == nil {
		return StringInput("AllergyIntolerance.LastOccurrence", nil)
	}
	return StringInput("AllergyIntolerance.LastOccurrence", resource.LastOccurrence)
}
func (resource *AllergyIntolerance) T_ParticipantId(numParticipant int) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return StringInput("AllergyIntolerance.Participant["+strconv.Itoa(numParticipant)+"].Id", nil)
	}
	return StringInput("AllergyIntolerance.Participant["+strconv.Itoa(numParticipant)+"].Id", resource.Participant[numParticipant].Id)
}
func (resource *AllergyIntolerance) T_ParticipantFunction(numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return CodeableConceptSelect("AllergyIntolerance.Participant["+strconv.Itoa(numParticipant)+"].Function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AllergyIntolerance.Participant["+strconv.Itoa(numParticipant)+"].Function", resource.Participant[numParticipant].Function, optionsValueSet)
}
func (resource *AllergyIntolerance) T_ReactionId(numReaction int) templ.Component {

	if resource == nil || len(resource.Reaction) >= numReaction {
		return StringInput("AllergyIntolerance.Reaction["+strconv.Itoa(numReaction)+"].Id", nil)
	}
	return StringInput("AllergyIntolerance.Reaction["+strconv.Itoa(numReaction)+"].Id", resource.Reaction[numReaction].Id)
}
func (resource *AllergyIntolerance) T_ReactionSubstance(numReaction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Reaction) >= numReaction {
		return CodeableConceptSelect("AllergyIntolerance.Reaction["+strconv.Itoa(numReaction)+"].Substance", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AllergyIntolerance.Reaction["+strconv.Itoa(numReaction)+"].Substance", resource.Reaction[numReaction].Substance, optionsValueSet)
}
func (resource *AllergyIntolerance) T_ReactionDescription(numReaction int) templ.Component {

	if resource == nil || len(resource.Reaction) >= numReaction {
		return StringInput("AllergyIntolerance.Reaction["+strconv.Itoa(numReaction)+"].Description", nil)
	}
	return StringInput("AllergyIntolerance.Reaction["+strconv.Itoa(numReaction)+"].Description", resource.Reaction[numReaction].Description)
}
func (resource *AllergyIntolerance) T_ReactionOnset(numReaction int) templ.Component {

	if resource == nil || len(resource.Reaction) >= numReaction {
		return StringInput("AllergyIntolerance.Reaction["+strconv.Itoa(numReaction)+"].Onset", nil)
	}
	return StringInput("AllergyIntolerance.Reaction["+strconv.Itoa(numReaction)+"].Onset", resource.Reaction[numReaction].Onset)
}
func (resource *AllergyIntolerance) T_ReactionSeverity(numReaction int) templ.Component {
	optionsValueSet := VSReaction_event_severity

	if resource == nil || len(resource.Reaction) >= numReaction {
		return CodeSelect("AllergyIntolerance.Reaction["+strconv.Itoa(numReaction)+"].Severity", nil, optionsValueSet)
	}
	return CodeSelect("AllergyIntolerance.Reaction["+strconv.Itoa(numReaction)+"].Severity", resource.Reaction[numReaction].Severity, optionsValueSet)
}
func (resource *AllergyIntolerance) T_ReactionExposureRoute(numReaction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Reaction) >= numReaction {
		return CodeableConceptSelect("AllergyIntolerance.Reaction["+strconv.Itoa(numReaction)+"].ExposureRoute", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AllergyIntolerance.Reaction["+strconv.Itoa(numReaction)+"].ExposureRoute", resource.Reaction[numReaction].ExposureRoute, optionsValueSet)
}
