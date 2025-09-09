package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/FamilyMemberHistory
type FamilyMemberHistory struct {
	Id                    *string                          `json:"id,omitempty"`
	Meta                  *Meta                            `json:"meta,omitempty"`
	ImplicitRules         *string                          `json:"implicitRules,omitempty"`
	Language              *string                          `json:"language,omitempty"`
	Text                  *Narrative                       `json:"text,omitempty"`
	Contained             []Resource                       `json:"contained,omitempty"`
	Extension             []Extension                      `json:"extension,omitempty"`
	ModifierExtension     []Extension                      `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                     `json:"identifier,omitempty"`
	InstantiatesCanonical []string                         `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                         `json:"instantiatesUri,omitempty"`
	Status                string                           `json:"status"`
	DataAbsentReason      *CodeableConcept                 `json:"dataAbsentReason,omitempty"`
	Patient               Reference                        `json:"patient"`
	Date                  *time.Time                       `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Participant           []FamilyMemberHistoryParticipant `json:"participant,omitempty"`
	Name                  *string                          `json:"name,omitempty"`
	Relationship          CodeableConcept                  `json:"relationship"`
	Sex                   *CodeableConcept                 `json:"sex,omitempty"`
	BornPeriod            *Period                          `json:"bornPeriod,omitempty"`
	BornDate              *time.Time                       `json:"bornDate,omitempty,format:'2006-01-02'"`
	BornString            *string                          `json:"bornString,omitempty"`
	AgeAge                *Age                             `json:"ageAge,omitempty"`
	AgeRange              *Range                           `json:"ageRange,omitempty"`
	AgeString             *string                          `json:"ageString,omitempty"`
	EstimatedAge          *bool                            `json:"estimatedAge,omitempty"`
	DeceasedBoolean       *bool                            `json:"deceasedBoolean,omitempty"`
	DeceasedAge           *Age                             `json:"deceasedAge,omitempty"`
	DeceasedRange         *Range                           `json:"deceasedRange,omitempty"`
	DeceasedDate          *time.Time                       `json:"deceasedDate,omitempty,format:'2006-01-02'"`
	DeceasedString        *string                          `json:"deceasedString,omitempty"`
	Reason                []CodeableReference              `json:"reason,omitempty"`
	Note                  []Annotation                     `json:"note,omitempty"`
	Condition             []FamilyMemberHistoryCondition   `json:"condition,omitempty"`
	Procedure             []FamilyMemberHistoryProcedure   `json:"procedure,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/FamilyMemberHistory
type FamilyMemberHistoryParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r5/StructureDefinition/FamilyMemberHistory
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

// http://hl7.org/fhir/r5/StructureDefinition/FamilyMemberHistory
type FamilyMemberHistoryProcedure struct {
	Id                 *string          `json:"id,omitempty"`
	Extension          []Extension      `json:"extension,omitempty"`
	ModifierExtension  []Extension      `json:"modifierExtension,omitempty"`
	Code               CodeableConcept  `json:"code"`
	Outcome            *CodeableConcept `json:"outcome,omitempty"`
	ContributedToDeath *bool            `json:"contributedToDeath,omitempty"`
	PerformedAge       *Age             `json:"performedAge,omitempty"`
	PerformedRange     *Range           `json:"performedRange,omitempty"`
	PerformedPeriod    *Period          `json:"performedPeriod,omitempty"`
	PerformedString    *string          `json:"performedString,omitempty"`
	PerformedDateTime  *time.Time       `json:"performedDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
		return StringInput("FamilyMemberHistory.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("FamilyMemberHistory.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *FamilyMemberHistory) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("FamilyMemberHistory.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("FamilyMemberHistory.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSHistory_status

	if resource == nil {
		return CodeSelect("FamilyMemberHistory.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("FamilyMemberHistory.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DataAbsentReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("FamilyMemberHistory.DataAbsentReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("FamilyMemberHistory.DataAbsentReason", resource.DataAbsentReason, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("FamilyMemberHistory.Date", nil, htmlAttrs)
	}
	return DateTimeInput("FamilyMemberHistory.Date", resource.Date, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("FamilyMemberHistory.Name", nil, htmlAttrs)
	}
	return StringInput("FamilyMemberHistory.Name", resource.Name, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Relationship(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("FamilyMemberHistory.Relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("FamilyMemberHistory.Relationship", &resource.Relationship, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Sex(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("FamilyMemberHistory.Sex", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("FamilyMemberHistory.Sex", resource.Sex, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_BornDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("FamilyMemberHistory.BornDate", nil, htmlAttrs)
	}
	return DateInput("FamilyMemberHistory.BornDate", resource.BornDate, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_BornString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("FamilyMemberHistory.BornString", nil, htmlAttrs)
	}
	return StringInput("FamilyMemberHistory.BornString", resource.BornString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_AgeString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("FamilyMemberHistory.AgeString", nil, htmlAttrs)
	}
	return StringInput("FamilyMemberHistory.AgeString", resource.AgeString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_EstimatedAge(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("FamilyMemberHistory.EstimatedAge", nil, htmlAttrs)
	}
	return BoolInput("FamilyMemberHistory.EstimatedAge", resource.EstimatedAge, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DeceasedBoolean(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("FamilyMemberHistory.DeceasedBoolean", nil, htmlAttrs)
	}
	return BoolInput("FamilyMemberHistory.DeceasedBoolean", resource.DeceasedBoolean, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DeceasedDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("FamilyMemberHistory.DeceasedDate", nil, htmlAttrs)
	}
	return DateInput("FamilyMemberHistory.DeceasedDate", resource.DeceasedDate, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DeceasedString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("FamilyMemberHistory.DeceasedString", nil, htmlAttrs)
	}
	return StringInput("FamilyMemberHistory.DeceasedString", resource.DeceasedString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("FamilyMemberHistory.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("FamilyMemberHistory.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ParticipantFunction(numParticipant int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeableConceptSelect("FamilyMemberHistory.Participant["+strconv.Itoa(numParticipant)+"].Function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("FamilyMemberHistory.Participant["+strconv.Itoa(numParticipant)+"].Function", resource.Participant[numParticipant].Function, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionCode(numCondition int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return CodeableConceptSelect("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].Code", &resource.Condition[numCondition].Code, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionOutcome(numCondition int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return CodeableConceptSelect("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].Outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].Outcome", resource.Condition[numCondition].Outcome, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionContributedToDeath(numCondition int, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return BoolInput("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].ContributedToDeath", nil, htmlAttrs)
	}
	return BoolInput("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].ContributedToDeath", resource.Condition[numCondition].ContributedToDeath, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionOnsetString(numCondition int, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return StringInput("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].OnsetString", nil, htmlAttrs)
	}
	return StringInput("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].OnsetString", resource.Condition[numCondition].OnsetString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionNote(numCondition int, numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) || numNote >= len(resource.Condition[numCondition].Note) {
		return AnnotationTextArea("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("FamilyMemberHistory.Condition["+strconv.Itoa(numCondition)+"].Note["+strconv.Itoa(numNote)+"]", &resource.Condition[numCondition].Note[numNote], htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ProcedureCode(numProcedure int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return CodeableConceptSelect("FamilyMemberHistory.Procedure["+strconv.Itoa(numProcedure)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("FamilyMemberHistory.Procedure["+strconv.Itoa(numProcedure)+"].Code", &resource.Procedure[numProcedure].Code, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ProcedureOutcome(numProcedure int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return CodeableConceptSelect("FamilyMemberHistory.Procedure["+strconv.Itoa(numProcedure)+"].Outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("FamilyMemberHistory.Procedure["+strconv.Itoa(numProcedure)+"].Outcome", resource.Procedure[numProcedure].Outcome, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ProcedureContributedToDeath(numProcedure int, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return BoolInput("FamilyMemberHistory.Procedure["+strconv.Itoa(numProcedure)+"].ContributedToDeath", nil, htmlAttrs)
	}
	return BoolInput("FamilyMemberHistory.Procedure["+strconv.Itoa(numProcedure)+"].ContributedToDeath", resource.Procedure[numProcedure].ContributedToDeath, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ProcedurePerformedString(numProcedure int, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return StringInput("FamilyMemberHistory.Procedure["+strconv.Itoa(numProcedure)+"].PerformedString", nil, htmlAttrs)
	}
	return StringInput("FamilyMemberHistory.Procedure["+strconv.Itoa(numProcedure)+"].PerformedString", resource.Procedure[numProcedure].PerformedString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ProcedurePerformedDateTime(numProcedure int, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return DateTimeInput("FamilyMemberHistory.Procedure["+strconv.Itoa(numProcedure)+"].PerformedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("FamilyMemberHistory.Procedure["+strconv.Itoa(numProcedure)+"].PerformedDateTime", resource.Procedure[numProcedure].PerformedDateTime, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ProcedureNote(numProcedure int, numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) || numNote >= len(resource.Procedure[numProcedure].Note) {
		return AnnotationTextArea("FamilyMemberHistory.Procedure["+strconv.Itoa(numProcedure)+"].Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("FamilyMemberHistory.Procedure["+strconv.Itoa(numProcedure)+"].Note["+strconv.Itoa(numNote)+"]", &resource.Procedure[numProcedure].Note[numNote], htmlAttrs)
}
