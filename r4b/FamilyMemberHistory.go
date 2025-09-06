package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/FamilyMemberHistory
type FamilyMemberHistory struct {
	Id                    *string                        `json:"id,omitempty"`
	Meta                  *Meta                          `json:"meta,omitempty"`
	ImplicitRules         *string                        `json:"implicitRules,omitempty"`
	Language              *string                        `json:"language,omitempty"`
	Text                  *Narrative                     `json:"text,omitempty"`
	Contained             []Resource                     `json:"contained,omitempty"`
	Extension             []Extension                    `json:"extension,omitempty"`
	ModifierExtension     []Extension                    `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                   `json:"identifier,omitempty"`
	InstantiatesCanonical []string                       `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                       `json:"instantiatesUri,omitempty"`
	Status                string                         `json:"status"`
	DataAbsentReason      *CodeableConcept               `json:"dataAbsentReason,omitempty"`
	Patient               Reference                      `json:"patient"`
	Date                  *string                        `json:"date,omitempty"`
	Name                  *string                        `json:"name,omitempty"`
	Relationship          CodeableConcept                `json:"relationship"`
	Sex                   *CodeableConcept               `json:"sex,omitempty"`
	BornPeriod            *Period                        `json:"bornPeriod,omitempty"`
	BornDate              *string                        `json:"bornDate,omitempty"`
	BornString            *string                        `json:"bornString,omitempty"`
	AgeAge                *Age                           `json:"ageAge,omitempty"`
	AgeRange              *Range                         `json:"ageRange,omitempty"`
	AgeString             *string                        `json:"ageString,omitempty"`
	EstimatedAge          *bool                          `json:"estimatedAge,omitempty"`
	DeceasedBoolean       *bool                          `json:"deceasedBoolean,omitempty"`
	DeceasedAge           *Age                           `json:"deceasedAge,omitempty"`
	DeceasedRange         *Range                         `json:"deceasedRange,omitempty"`
	DeceasedDate          *string                        `json:"deceasedDate,omitempty"`
	DeceasedString        *string                        `json:"deceasedString,omitempty"`
	ReasonCode            []CodeableConcept              `json:"reasonCode,omitempty"`
	ReasonReference       []Reference                    `json:"reasonReference,omitempty"`
	Note                  []Annotation                   `json:"note,omitempty"`
	Condition             []FamilyMemberHistoryCondition `json:"condition,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/FamilyMemberHistory
type FamilyMemberHistoryCondition struct {
	Id                 *string          `json:"id,omitempty"`
	Extension          []Extension      `json:"extension,omitempty"`
	ModifierExtension  []Extension      `json:"modifierExtension,omitempty"`
	Code               CodeableConcept  `json:"code"`
	Outcome            *CodeableConcept `json:"outcome,omitempty"`
	ContributedToDeath *bool            `json:"contributedToDeath,omitempty"`
	OnsetAge           *Age             `json:"onsetAge,omitempty"`
	OnsetRange         *Range           `json:"onsetRange,omitempty"`
	OnsetPeriod        *Period          `json:"onsetPeriod,omitempty"`
	OnsetString        *string          `json:"onsetString,omitempty"`
	Note               []Annotation     `json:"note,omitempty"`
}

type OtherFamilyMemberHistory FamilyMemberHistory

// on convert struct to json, automatically add resourceType=FamilyMemberHistory
func (r FamilyMemberHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFamilyMemberHistory
		ResourceType string `json:"resourceType"`
	}{
		OtherFamilyMemberHistory: OtherFamilyMemberHistory(r),
		ResourceType:             "FamilyMemberHistory",
	})
}

func (resource *FamilyMemberHistory) T_Id() templ.Component {

	if resource == nil {
		return StringInput("FamilyMemberHistory.Id", nil)
	}
	return StringInput("FamilyMemberHistory.Id", resource.Id)
}
func (resource *FamilyMemberHistory) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("FamilyMemberHistory.ImplicitRules", nil)
	}
	return StringInput("FamilyMemberHistory.ImplicitRules", resource.ImplicitRules)
}
func (resource *FamilyMemberHistory) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("FamilyMemberHistory.Language", nil, optionsValueSet)
	}
	return CodeSelect("FamilyMemberHistory.Language", resource.Language, optionsValueSet)
}
func (resource *FamilyMemberHistory) T_InstantiatesCanonical(numInstantiatesCanonical int) templ.Component {

	if resource == nil || len(resource.InstantiatesCanonical) >= numInstantiatesCanonical {
		return StringInput("FamilyMemberHistory.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil)
	}
	return StringInput("FamilyMemberHistory.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical])
}
func (resource *FamilyMemberHistory) T_InstantiatesUri(numInstantiatesUri int) templ.Component {

	if resource == nil || len(resource.InstantiatesUri) >= numInstantiatesUri {
		return StringInput("FamilyMemberHistory.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil)
	}
	return StringInput("FamilyMemberHistory.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri])
}
func (resource *FamilyMemberHistory) T_Status() templ.Component {
	optionsValueSet := VSHistory_status

	if resource == nil {
		return CodeSelect("FamilyMemberHistory.Status", nil, optionsValueSet)
	}
	return CodeSelect("FamilyMemberHistory.Status", &resource.Status, optionsValueSet)
}
func (resource *FamilyMemberHistory) T_DataAbsentReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("FamilyMemberHistory.DataAbsentReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("FamilyMemberHistory.DataAbsentReason", resource.DataAbsentReason, optionsValueSet)
}
func (resource *FamilyMemberHistory) T_Date() templ.Component {

	if resource == nil {
		return StringInput("FamilyMemberHistory.Date", nil)
	}
	return StringInput("FamilyMemberHistory.Date", resource.Date)
}
func (resource *FamilyMemberHistory) T_Name() templ.Component {

	if resource == nil {
		return StringInput("FamilyMemberHistory.Name", nil)
	}
	return StringInput("FamilyMemberHistory.Name", resource.Name)
}
func (resource *FamilyMemberHistory) T_Relationship(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("FamilyMemberHistory.Relationship", nil, optionsValueSet)
	}
	return CodeableConceptSelect("FamilyMemberHistory.Relationship", &resource.Relationship, optionsValueSet)
}
func (resource *FamilyMemberHistory) T_Sex(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("FamilyMemberHistory.Sex", nil, optionsValueSet)
	}
	return CodeableConceptSelect("FamilyMemberHistory.Sex", resource.Sex, optionsValueSet)
}
func (resource *FamilyMemberHistory) T_EstimatedAge() templ.Component {

	if resource == nil {
		return BoolInput("FamilyMemberHistory.EstimatedAge", nil)
	}
	return BoolInput("FamilyMemberHistory.EstimatedAge", resource.EstimatedAge)
}
func (resource *FamilyMemberHistory) T_ReasonCode(numReasonCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ReasonCode) >= numReasonCode {
		return CodeableConceptSelect("FamilyMemberHistory.ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("FamilyMemberHistory.ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet)
}
func (resource *FamilyMemberHistory) T_ConditionId(numCondition int) templ.Component {

	if resource == nil || len(resource.Condition) >= numCondition {
		return StringInput("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].Id", nil)
	}
	return StringInput("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].Id", resource.Condition[numCondition].Id)
}
func (resource *FamilyMemberHistory) T_ConditionCode(numCondition int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Condition) >= numCondition {
		return CodeableConceptSelect("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].Code", &resource.Condition[numCondition].Code, optionsValueSet)
}
func (resource *FamilyMemberHistory) T_ConditionOutcome(numCondition int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Condition) >= numCondition {
		return CodeableConceptSelect("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].Outcome", nil, optionsValueSet)
	}
	return CodeableConceptSelect("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].Outcome", resource.Condition[numCondition].Outcome, optionsValueSet)
}
func (resource *FamilyMemberHistory) T_ConditionContributedToDeath(numCondition int) templ.Component {

	if resource == nil || len(resource.Condition) >= numCondition {
		return BoolInput("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].ContributedToDeath", nil)
	}
	return BoolInput("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].ContributedToDeath", resource.Condition[numCondition].ContributedToDeath)
}
