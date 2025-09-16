package r5

//generated with command go run ./bultaoreune
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
	OnsetDateTime      *FhirDateTime                   `json:"onsetDateTime,omitempty"`
	OnsetAge           *Age                            `json:"onsetAge,omitempty"`
	OnsetPeriod        *Period                         `json:"onsetPeriod,omitempty"`
	OnsetRange         *Range                          `json:"onsetRange,omitempty"`
	OnsetString        *string                         `json:"onsetString,omitempty"`
	RecordedDate       *FhirDateTime                   `json:"recordedDate,omitempty"`
	Participant        []AllergyIntoleranceParticipant `json:"participant,omitempty"`
	LastOccurrence     *FhirDateTime                   `json:"lastOccurrence,omitempty"`
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
	Onset             *FhirDateTime       `json:"onset,omitempty"`
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
func (resource *AllergyIntolerance) T_ClinicalStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("clinicalStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("clinicalStatus", resource.ClinicalStatus, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_VerificationStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("verificationStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("verificationStatus", resource.VerificationStatus, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_Category(numCategory int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAllergy_intolerance_category

	if resource == nil || numCategory >= len(resource.Category) {
		return CodeSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_Criticality(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAllergy_intolerance_criticality

	if resource == nil {
		return CodeSelect("criticality", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("criticality", resource.Criticality, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_OnsetDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("onsetDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("onsetDateTime", resource.OnsetDateTime, htmlAttrs)
}
func (resource *AllergyIntolerance) T_OnsetString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("onsetString", nil, htmlAttrs)
	}
	return StringInput("onsetString", resource.OnsetString, htmlAttrs)
}
func (resource *AllergyIntolerance) T_RecordedDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("recordedDate", nil, htmlAttrs)
	}
	return DateTimeInput("recordedDate", resource.RecordedDate, htmlAttrs)
}
func (resource *AllergyIntolerance) T_LastOccurrence(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("lastOccurrence", nil, htmlAttrs)
	}
	return DateTimeInput("lastOccurrence", resource.LastOccurrence, htmlAttrs)
}
func (resource *AllergyIntolerance) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *AllergyIntolerance) T_ParticipantFunction(numParticipant int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].function", resource.Participant[numParticipant].Function, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_ReactionSubstance(numReaction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) {
		return CodeableConceptSelect("reaction["+strconv.Itoa(numReaction)+"].substance", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reaction["+strconv.Itoa(numReaction)+"].substance", resource.Reaction[numReaction].Substance, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_ReactionDescription(numReaction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) {
		return StringInput("reaction["+strconv.Itoa(numReaction)+"].description", nil, htmlAttrs)
	}
	return StringInput("reaction["+strconv.Itoa(numReaction)+"].description", resource.Reaction[numReaction].Description, htmlAttrs)
}
func (resource *AllergyIntolerance) T_ReactionOnset(numReaction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) {
		return DateTimeInput("reaction["+strconv.Itoa(numReaction)+"].onset", nil, htmlAttrs)
	}
	return DateTimeInput("reaction["+strconv.Itoa(numReaction)+"].onset", resource.Reaction[numReaction].Onset, htmlAttrs)
}
func (resource *AllergyIntolerance) T_ReactionSeverity(numReaction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSReaction_event_severity

	if resource == nil || numReaction >= len(resource.Reaction) {
		return CodeSelect("reaction["+strconv.Itoa(numReaction)+"].severity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("reaction["+strconv.Itoa(numReaction)+"].severity", resource.Reaction[numReaction].Severity, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_ReactionExposureRoute(numReaction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) {
		return CodeableConceptSelect("reaction["+strconv.Itoa(numReaction)+"].exposureRoute", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reaction["+strconv.Itoa(numReaction)+"].exposureRoute", resource.Reaction[numReaction].ExposureRoute, optionsValueSet, htmlAttrs)
}
func (resource *AllergyIntolerance) T_ReactionNote(numReaction int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) || numNote >= len(resource.Reaction[numReaction].Note) {
		return AnnotationTextArea("reaction["+strconv.Itoa(numReaction)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("reaction["+strconv.Itoa(numReaction)+"].note["+strconv.Itoa(numNote)+"]", &resource.Reaction[numReaction].Note[numNote], htmlAttrs)
}
