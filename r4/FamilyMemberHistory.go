package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/FamilyMemberHistory
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

// http://hl7.org/fhir/r4/StructureDefinition/FamilyMemberHistory
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
func (r FamilyMemberHistory) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "FamilyMemberHistory/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "FamilyMemberHistory"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *FamilyMemberHistory) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *FamilyMemberHistory) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSHistory_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DataAbsentReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("dataAbsentReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("dataAbsentReason", resource.DataAbsentReason, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Relationship(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relationship", &resource.Relationship, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Sex(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("sex", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("sex", resource.Sex, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_BornDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("bornDate", nil, htmlAttrs)
	}
	return DateInput("bornDate", resource.BornDate, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_BornString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("bornString", nil, htmlAttrs)
	}
	return StringInput("bornString", resource.BornString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_AgeString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ageString", nil, htmlAttrs)
	}
	return StringInput("ageString", resource.AgeString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_EstimatedAge(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("estimatedAge", nil, htmlAttrs)
	}
	return BoolInput("estimatedAge", resource.EstimatedAge, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DeceasedBoolean(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("deceasedBoolean", nil, htmlAttrs)
	}
	return BoolInput("deceasedBoolean", resource.DeceasedBoolean, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DeceasedDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("deceasedDate", nil, htmlAttrs)
	}
	return DateInput("deceasedDate", resource.DeceasedDate, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DeceasedString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("deceasedString", nil, htmlAttrs)
	}
	return StringInput("deceasedString", resource.DeceasedString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionCode(numCondition int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return CodeableConceptSelect("condition["+strconv.Itoa(numCondition)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("condition["+strconv.Itoa(numCondition)+"].code", &resource.Condition[numCondition].Code, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionOutcome(numCondition int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return CodeableConceptSelect("condition["+strconv.Itoa(numCondition)+"].outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("condition["+strconv.Itoa(numCondition)+"].outcome", resource.Condition[numCondition].Outcome, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionContributedToDeath(numCondition int, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return BoolInput("condition["+strconv.Itoa(numCondition)+"].contributedToDeath", nil, htmlAttrs)
	}
	return BoolInput("condition["+strconv.Itoa(numCondition)+"].contributedToDeath", resource.Condition[numCondition].ContributedToDeath, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionOnsetString(numCondition int, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return StringInput("condition["+strconv.Itoa(numCondition)+"].onsetString", nil, htmlAttrs)
	}
	return StringInput("condition["+strconv.Itoa(numCondition)+"].onsetString", resource.Condition[numCondition].OnsetString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionNote(numCondition int, numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) || numNote >= len(resource.Condition[numCondition].Note) {
		return AnnotationTextArea("condition["+strconv.Itoa(numCondition)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("condition["+strconv.Itoa(numCondition)+"].note["+strconv.Itoa(numNote)+"]", &resource.Condition[numCondition].Note[numNote], htmlAttrs)
}
