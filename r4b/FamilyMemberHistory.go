package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	BornPeriod            *Period                        `json:"bornPeriod"`
	BornDate              *string                        `json:"bornDate"`
	BornString            *string                        `json:"bornString"`
	AgeAge                *Age                           `json:"ageAge"`
	AgeRange              *Range                         `json:"ageRange"`
	AgeString             *string                        `json:"ageString"`
	EstimatedAge          *bool                          `json:"estimatedAge,omitempty"`
	DeceasedBoolean       *bool                          `json:"deceasedBoolean"`
	DeceasedAge           *Age                           `json:"deceasedAge"`
	DeceasedRange         *Range                         `json:"deceasedRange"`
	DeceasedDate          *string                        `json:"deceasedDate"`
	DeceasedString        *string                        `json:"deceasedString"`
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
	OnsetAge           *Age             `json:"onsetAge"`
	OnsetRange         *Range           `json:"onsetRange"`
	OnsetPeriod        *Period          `json:"onsetPeriod"`
	OnsetString        *string          `json:"onsetString"`
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

func (resource *FamilyMemberHistory) FamilyMemberHistoryLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *FamilyMemberHistory) FamilyMemberHistoryStatus() templ.Component {
	optionsValueSet := VSHistory_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *FamilyMemberHistory) FamilyMemberHistoryDataAbsentReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("dataAbsentReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("dataAbsentReason", resource.DataAbsentReason, optionsValueSet)
}
func (resource *FamilyMemberHistory) FamilyMemberHistoryRelationship(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("relationship", nil, optionsValueSet)
	}
	return CodeableConceptSelect("relationship", &resource.Relationship, optionsValueSet)
}
func (resource *FamilyMemberHistory) FamilyMemberHistorySex(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("sex", nil, optionsValueSet)
	}
	return CodeableConceptSelect("sex", resource.Sex, optionsValueSet)
}
func (resource *FamilyMemberHistory) FamilyMemberHistoryReasonCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("reasonCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reasonCode", &resource.ReasonCode[0], optionsValueSet)
}
func (resource *FamilyMemberHistory) FamilyMemberHistoryConditionCode(numCondition int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Condition) >= numCondition {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Condition[numCondition].Code, optionsValueSet)
}
func (resource *FamilyMemberHistory) FamilyMemberHistoryConditionOutcome(numCondition int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Condition) >= numCondition {
		return CodeableConceptSelect("outcome", nil, optionsValueSet)
	}
	return CodeableConceptSelect("outcome", resource.Condition[numCondition].Outcome, optionsValueSet)
}
