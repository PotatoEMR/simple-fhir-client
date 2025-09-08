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

// http://hl7.org/fhir/r4b/StructureDefinition/AllergyIntolerance
type AllergyIntolerance struct {
	Id                 *string                      `json:"id,omitempty"`
	Meta               *Meta                        `json:"meta,omitempty"`
	ImplicitRules      *string                      `json:"implicitRules,omitempty"`
	Language           *string                      `json:"language,omitempty"`
	Text               *Narrative                   `json:"text,omitempty"`
	Contained          []Resource                   `json:"contained,omitempty"`
	Extension          []Extension                  `json:"extension,omitempty"`
	ModifierExtension  []Extension                  `json:"modifierExtension,omitempty"`
	Identifier         []Identifier                 `json:"identifier,omitempty"`
	ClinicalStatus     *CodeableConcept             `json:"clinicalStatus,omitempty"`
	VerificationStatus *CodeableConcept             `json:"verificationStatus,omitempty"`
	Type               *string                      `json:"type,omitempty"`
	Category           []string                     `json:"category,omitempty"`
	Criticality        *string                      `json:"criticality,omitempty"`
	Code               *CodeableConcept             `json:"code,omitempty"`
	Patient            Reference                    `json:"patient"`
	Encounter          *Reference                   `json:"encounter,omitempty"`
	OnsetDateTime      *time.Time                   `json:"onsetDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	OnsetAge           *Age                         `json:"onsetAge,omitempty"`
	OnsetPeriod        *Period                      `json:"onsetPeriod,omitempty"`
	OnsetRange         *Range                       `json:"onsetRange,omitempty"`
	OnsetString        *string                      `json:"onsetString,omitempty"`
	RecordedDate       *time.Time                   `json:"recordedDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Recorder           *Reference                   `json:"recorder,omitempty"`
	Asserter           *Reference                   `json:"asserter,omitempty"`
	LastOccurrence     *time.Time                   `json:"lastOccurrence,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Note               []Annotation                 `json:"note,omitempty"`
	Reaction           []AllergyIntoleranceReaction `json:"reaction,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/AllergyIntolerance
type AllergyIntoleranceReaction struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Substance         *CodeableConcept  `json:"substance,omitempty"`
	Manifestation     []CodeableConcept `json:"manifestation"`
	Description       *string           `json:"description,omitempty"`
	Onset             *time.Time        `json:"onset,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Severity          *string           `json:"severity,omitempty"`
	ExposureRoute     *CodeableConcept  `json:"exposureRoute,omitempty"`
	Note              []Annotation      `json:"note,omitempty"`
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
func (r AllergyIntolerance) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "AllergyIntolerance/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "AllergyIntolerance"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *AllergyIntolerance) T_ClinicalStatus(htmlAttrs string) templ.Component {
	optionsValueSet := VSAllergyintolerance_clinical

	if resource == nil {
		return CodeableConceptSelect("ClinicalStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ClinicalStatus", resource.ClinicalStatus, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_VerificationStatus(htmlAttrs string) templ.Component {
	optionsValueSet := VSAllergyintolerance_verification

	if resource == nil {
		return CodeableConceptSelect("VerificationStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("VerificationStatus", resource.VerificationStatus, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_Type(htmlAttrs string) templ.Component {
	optionsValueSet := VSAllergy_intolerance_type

	if resource == nil {
		return CodeSelect("Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_Category(numCategory int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAllergy_intolerance_category

	if resource == nil || numCategory >= len(resource.Category) {
		return CodeSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_Criticality(htmlAttrs string) templ.Component {
	optionsValueSet := VSAllergy_intolerance_criticality

	if resource == nil {
		return CodeSelect("Criticality", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Criticality", resource.Criticality, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_OnsetDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("OnsetDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("OnsetDateTime", resource.OnsetDateTime, htmlAttrs)
}
func (resource *AllergyIntolerance) T_OnsetString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("OnsetString", nil, htmlAttrs)
	}
	return StringInput("OnsetString", resource.OnsetString, htmlAttrs)
}
func (resource *AllergyIntolerance) T_RecordedDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("RecordedDate", nil, htmlAttrs)
	}
	return DateTimeInput("RecordedDate", resource.RecordedDate, htmlAttrs)
}
func (resource *AllergyIntolerance) T_LastOccurrence(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("LastOccurrence", nil, htmlAttrs)
	}
	return DateTimeInput("LastOccurrence", resource.LastOccurrence, htmlAttrs)
}
func (resource *AllergyIntolerance) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *AllergyIntolerance) T_ReactionSubstance(numReaction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) {
		return CodeableConceptSelect("Reaction["+strconv.Itoa(numReaction)+"]Substance", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Reaction["+strconv.Itoa(numReaction)+"]Substance", resource.Reaction[numReaction].Substance, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_ReactionManifestation(numReaction int, numManifestation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) || numManifestation >= len(resource.Reaction[numReaction].Manifestation) {
		return CodeableConceptSelect("Reaction["+strconv.Itoa(numReaction)+"]Manifestation["+strconv.Itoa(numManifestation)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Reaction["+strconv.Itoa(numReaction)+"]Manifestation["+strconv.Itoa(numManifestation)+"]", &resource.Reaction[numReaction].Manifestation[numManifestation], optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_ReactionDescription(numReaction int, htmlAttrs string) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) {
		return StringInput("Reaction["+strconv.Itoa(numReaction)+"]Description", nil, htmlAttrs)
	}
	return StringInput("Reaction["+strconv.Itoa(numReaction)+"]Description", resource.Reaction[numReaction].Description, htmlAttrs)
}
func (resource *AllergyIntolerance) T_ReactionOnset(numReaction int, htmlAttrs string) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) {
		return DateTimeInput("Reaction["+strconv.Itoa(numReaction)+"]Onset", nil, htmlAttrs)
	}
	return DateTimeInput("Reaction["+strconv.Itoa(numReaction)+"]Onset", resource.Reaction[numReaction].Onset, htmlAttrs)
}
func (resource *AllergyIntolerance) T_ReactionSeverity(numReaction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSReaction_event_severity

	if resource == nil || numReaction >= len(resource.Reaction) {
		return CodeSelect("Reaction["+strconv.Itoa(numReaction)+"]Severity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Reaction["+strconv.Itoa(numReaction)+"]Severity", resource.Reaction[numReaction].Severity, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_ReactionExposureRoute(numReaction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) {
		return CodeableConceptSelect("Reaction["+strconv.Itoa(numReaction)+"]ExposureRoute", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Reaction["+strconv.Itoa(numReaction)+"]ExposureRoute", resource.Reaction[numReaction].ExposureRoute, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_ReactionNote(numReaction int, numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) || numNote >= len(resource.Reaction[numReaction].Note) {
		return AnnotationTextArea("Reaction["+strconv.Itoa(numReaction)+"]Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Reaction["+strconv.Itoa(numReaction)+"]Note["+strconv.Itoa(numNote)+"]", &resource.Reaction[numReaction].Note[numNote], htmlAttrs)
}
