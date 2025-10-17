package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Date                  *FhirDateTime                    `json:"date,omitempty"`
	Participant           []FamilyMemberHistoryParticipant `json:"participant,omitempty"`
	Name                  *string                          `json:"name,omitempty"`
	Relationship          CodeableConcept                  `json:"relationship"`
	Sex                   *CodeableConcept                 `json:"sex,omitempty"`
	BornPeriod            *Period                          `json:"bornPeriod,omitempty"`
	BornDate              *FhirDate                        `json:"bornDate,omitempty"`
	BornString            *string                          `json:"bornString,omitempty"`
	AgeAge                *Age                             `json:"ageAge,omitempty"`
	AgeRange              *Range                           `json:"ageRange,omitempty"`
	AgeString             *string                          `json:"ageString,omitempty"`
	EstimatedAge          *bool                            `json:"estimatedAge,omitempty"`
	DeceasedBoolean       *bool                            `json:"deceasedBoolean,omitempty"`
	DeceasedAge           *Age                             `json:"deceasedAge,omitempty"`
	DeceasedRange         *Range                           `json:"deceasedRange,omitempty"`
	DeceasedDate          *FhirDate                        `json:"deceasedDate,omitempty"`
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
	PerformedDateTime  *FhirDateTime    `json:"performedDateTime,omitempty"`
	Note               []Annotation     `json:"note,omitempty"`
}

type OtherFamilyMemberHistory FamilyMemberHistory

// struct -> json, automatically add resourceType=Patient
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
func (r FamilyMemberHistory) ResourceType() string {
	return "FamilyMemberHistory"
}

func (resource *FamilyMemberHistory) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *FamilyMemberHistory) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSHistory_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DataAbsentReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("dataAbsentReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("dataAbsentReason", resource.DataAbsentReason, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Patient(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "patient", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "patient", &resource.Patient, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Relationship(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relationship", &resource.Relationship, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Sex(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("sex", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("sex", resource.Sex, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_BornPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("bornPeriod", nil, htmlAttrs)
	}
	return PeriodInput("bornPeriod", resource.BornPeriod, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_BornDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("bornDate", nil, htmlAttrs)
	}
	return FhirDateInput("bornDate", resource.BornDate, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_BornString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("bornString", nil, htmlAttrs)
	}
	return StringInput("bornString", resource.BornString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_AgeAge(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return AgeInput("ageAge", nil, htmlAttrs)
	}
	return AgeInput("ageAge", resource.AgeAge, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_AgeRange(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return RangeInput("ageRange", nil, htmlAttrs)
	}
	return RangeInput("ageRange", resource.AgeRange, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_AgeString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("ageString", nil, htmlAttrs)
	}
	return StringInput("ageString", resource.AgeString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_EstimatedAge(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("estimatedAge", nil, htmlAttrs)
	}
	return BoolInput("estimatedAge", resource.EstimatedAge, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DeceasedBoolean(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("deceasedBoolean", nil, htmlAttrs)
	}
	return BoolInput("deceasedBoolean", resource.DeceasedBoolean, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DeceasedAge(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return AgeInput("deceasedAge", nil, htmlAttrs)
	}
	return AgeInput("deceasedAge", resource.DeceasedAge, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DeceasedRange(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return RangeInput("deceasedRange", nil, htmlAttrs)
	}
	return RangeInput("deceasedRange", resource.DeceasedRange, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DeceasedDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("deceasedDate", nil, htmlAttrs)
	}
	return FhirDateInput("deceasedDate", resource.DeceasedDate, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_DeceasedString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("deceasedString", nil, htmlAttrs)
	}
	return StringInput("deceasedString", resource.DeceasedString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Reason(numReason int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], htmlAttrs)
}
func (resource *FamilyMemberHistory) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ParticipantFunction(numParticipant int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].function", resource.Participant[numParticipant].Function, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ParticipantActor(frs []FhirResource, numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return ReferenceInput(frs, "participant["+strconv.Itoa(numParticipant)+"].actor", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "participant["+strconv.Itoa(numParticipant)+"].actor", &resource.Participant[numParticipant].Actor, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionCode(numCondition int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return CodeableConceptSelect("condition["+strconv.Itoa(numCondition)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("condition["+strconv.Itoa(numCondition)+"].code", &resource.Condition[numCondition].Code, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionOutcome(numCondition int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return CodeableConceptSelect("condition["+strconv.Itoa(numCondition)+"].outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("condition["+strconv.Itoa(numCondition)+"].outcome", resource.Condition[numCondition].Outcome, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionContributedToDeath(numCondition int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return BoolInput("condition["+strconv.Itoa(numCondition)+"].contributedToDeath", nil, htmlAttrs)
	}
	return BoolInput("condition["+strconv.Itoa(numCondition)+"].contributedToDeath", resource.Condition[numCondition].ContributedToDeath, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionOnsetAge(numCondition int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return AgeInput("condition["+strconv.Itoa(numCondition)+"].onsetAge", nil, htmlAttrs)
	}
	return AgeInput("condition["+strconv.Itoa(numCondition)+"].onsetAge", resource.Condition[numCondition].OnsetAge, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionOnsetRange(numCondition int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return RangeInput("condition["+strconv.Itoa(numCondition)+"].onsetRange", nil, htmlAttrs)
	}
	return RangeInput("condition["+strconv.Itoa(numCondition)+"].onsetRange", resource.Condition[numCondition].OnsetRange, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionOnsetPeriod(numCondition int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return PeriodInput("condition["+strconv.Itoa(numCondition)+"].onsetPeriod", nil, htmlAttrs)
	}
	return PeriodInput("condition["+strconv.Itoa(numCondition)+"].onsetPeriod", resource.Condition[numCondition].OnsetPeriod, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionOnsetString(numCondition int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return StringInput("condition["+strconv.Itoa(numCondition)+"].onsetString", nil, htmlAttrs)
	}
	return StringInput("condition["+strconv.Itoa(numCondition)+"].onsetString", resource.Condition[numCondition].OnsetString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ConditionNote(numCondition int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) || numNote >= len(resource.Condition[numCondition].Note) {
		return AnnotationTextArea("condition["+strconv.Itoa(numCondition)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("condition["+strconv.Itoa(numCondition)+"].note["+strconv.Itoa(numNote)+"]", &resource.Condition[numCondition].Note[numNote], htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ProcedureCode(numProcedure int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].code", &resource.Procedure[numProcedure].Code, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ProcedureOutcome(numProcedure int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].outcome", resource.Procedure[numProcedure].Outcome, optionsValueSet, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ProcedureContributedToDeath(numProcedure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return BoolInput("procedure["+strconv.Itoa(numProcedure)+"].contributedToDeath", nil, htmlAttrs)
	}
	return BoolInput("procedure["+strconv.Itoa(numProcedure)+"].contributedToDeath", resource.Procedure[numProcedure].ContributedToDeath, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ProcedurePerformedAge(numProcedure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return AgeInput("procedure["+strconv.Itoa(numProcedure)+"].performedAge", nil, htmlAttrs)
	}
	return AgeInput("procedure["+strconv.Itoa(numProcedure)+"].performedAge", resource.Procedure[numProcedure].PerformedAge, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ProcedurePerformedRange(numProcedure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return RangeInput("procedure["+strconv.Itoa(numProcedure)+"].performedRange", nil, htmlAttrs)
	}
	return RangeInput("procedure["+strconv.Itoa(numProcedure)+"].performedRange", resource.Procedure[numProcedure].PerformedRange, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ProcedurePerformedPeriod(numProcedure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return PeriodInput("procedure["+strconv.Itoa(numProcedure)+"].performedPeriod", nil, htmlAttrs)
	}
	return PeriodInput("procedure["+strconv.Itoa(numProcedure)+"].performedPeriod", resource.Procedure[numProcedure].PerformedPeriod, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ProcedurePerformedString(numProcedure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return StringInput("procedure["+strconv.Itoa(numProcedure)+"].performedString", nil, htmlAttrs)
	}
	return StringInput("procedure["+strconv.Itoa(numProcedure)+"].performedString", resource.Procedure[numProcedure].PerformedString, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ProcedurePerformedDateTime(numProcedure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return FhirDateTimeInput("procedure["+strconv.Itoa(numProcedure)+"].performedDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("procedure["+strconv.Itoa(numProcedure)+"].performedDateTime", resource.Procedure[numProcedure].PerformedDateTime, htmlAttrs)
}
func (resource *FamilyMemberHistory) T_ProcedureNote(numProcedure int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) || numNote >= len(resource.Procedure[numProcedure].Note) {
		return AnnotationTextArea("procedure["+strconv.Itoa(numProcedure)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("procedure["+strconv.Itoa(numProcedure)+"].note["+strconv.Itoa(numNote)+"]", &resource.Procedure[numProcedure].Note[numNote], htmlAttrs)
}
