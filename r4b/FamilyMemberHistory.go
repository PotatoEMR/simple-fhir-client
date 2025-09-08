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
	Date                  *time.Time                     `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Name                  *string                        `json:"name,omitempty"`
	Relationship          CodeableConcept                `json:"relationship"`
	Sex                   *CodeableConcept               `json:"sex,omitempty"`
	BornPeriod            *Period                        `json:"bornPeriod,omitempty"`
	BornDate              *time.Time                     `json:"bornDate,omitempty,format:'2006-01-02'"`
	BornString            *string                        `json:"bornString,omitempty"`
	AgeAge                *Age                           `json:"ageAge,omitempty"`
	AgeRange              *Range                         `json:"ageRange,omitempty"`
	AgeString             *string                        `json:"ageString,omitempty"`
	EstimatedAge          *bool                          `json:"estimatedAge,omitempty"`
	DeceasedBoolean       *bool                          `json:"deceasedBoolean,omitempty"`
	DeceasedAge           *Age                           `json:"deceasedAge,omitempty"`
	DeceasedRange         *Range                         `json:"deceasedRange,omitempty"`
	DeceasedDate          *time.Time                     `json:"deceasedDate,omitempty,format:'2006-01-02'"`
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
		return StringInput("InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *FamilyMemberHistory) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSHistory_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DataAbsentReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("DataAbsentReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DataAbsentReason", resource.DataAbsentReason, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Relationship(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Relationship", &resource.Relationship, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Sex(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Sex", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Sex", resource.Sex, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_BornDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("BornDate", nil, htmlAttrs)
	}
	return DateInput("BornDate", resource.BornDate, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_BornString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("BornString", nil, htmlAttrs)
	}
	return StringInput("BornString", resource.BornString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_AgeString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("AgeString", nil, htmlAttrs)
	}
	return StringInput("AgeString", resource.AgeString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_EstimatedAge(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("EstimatedAge", nil, htmlAttrs)
	}
	return BoolInput("EstimatedAge", resource.EstimatedAge, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DeceasedBoolean(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("DeceasedBoolean", nil, htmlAttrs)
	}
	return BoolInput("DeceasedBoolean", resource.DeceasedBoolean, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DeceasedDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("DeceasedDate", nil, htmlAttrs)
	}
	return DateInput("DeceasedDate", resource.DeceasedDate, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DeceasedString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DeceasedString", nil, htmlAttrs)
	}
	return StringInput("DeceasedString", resource.DeceasedString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionCode(numCondition int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return CodeableConceptSelect("Condition["+strconv.Itoa(numCondition)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Condition["+strconv.Itoa(numCondition)+"]Code", &resource.Condition[numCondition].Code, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionOutcome(numCondition int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return CodeableConceptSelect("Condition["+strconv.Itoa(numCondition)+"]Outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Condition["+strconv.Itoa(numCondition)+"]Outcome", resource.Condition[numCondition].Outcome, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionContributedToDeath(numCondition int, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return BoolInput("Condition["+strconv.Itoa(numCondition)+"]ContributedToDeath", nil, htmlAttrs)
	}
	return BoolInput("Condition["+strconv.Itoa(numCondition)+"]ContributedToDeath", resource.Condition[numCondition].ContributedToDeath, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionOnsetString(numCondition int, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return StringInput("Condition["+strconv.Itoa(numCondition)+"]OnsetString", nil, htmlAttrs)
	}
	return StringInput("Condition["+strconv.Itoa(numCondition)+"]OnsetString", resource.Condition[numCondition].OnsetString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionNote(numCondition int, numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) || numNote >= len(resource.Condition[numCondition].Note) {
		return AnnotationTextArea("Condition["+strconv.Itoa(numCondition)+"]Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Condition["+strconv.Itoa(numCondition)+"]Note["+strconv.Itoa(numNote)+"]", &resource.Condition[numCondition].Note[numNote], htmlAttrs)
}
