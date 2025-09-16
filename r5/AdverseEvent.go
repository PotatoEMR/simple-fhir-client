package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/AdverseEvent
type AdverseEvent struct {
	Id                      *string                          `json:"id,omitempty"`
	Meta                    *Meta                            `json:"meta,omitempty"`
	ImplicitRules           *string                          `json:"implicitRules,omitempty"`
	Language                *string                          `json:"language,omitempty"`
	Text                    *Narrative                       `json:"text,omitempty"`
	Contained               []Resource                       `json:"contained,omitempty"`
	Extension               []Extension                      `json:"extension,omitempty"`
	ModifierExtension       []Extension                      `json:"modifierExtension,omitempty"`
	Identifier              []Identifier                     `json:"identifier,omitempty"`
	Status                  string                           `json:"status"`
	Actuality               string                           `json:"actuality"`
	Category                []CodeableConcept                `json:"category,omitempty"`
	Code                    *CodeableConcept                 `json:"code,omitempty"`
	Subject                 Reference                        `json:"subject"`
	Encounter               *Reference                       `json:"encounter,omitempty"`
	OccurrenceDateTime      *FhirDateTime                    `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod        *Period                          `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming        *Timing                          `json:"occurrenceTiming,omitempty"`
	Detected                *FhirDateTime                    `json:"detected,omitempty"`
	RecordedDate            *FhirDateTime                    `json:"recordedDate,omitempty"`
	ResultingEffect         []Reference                      `json:"resultingEffect,omitempty"`
	Location                *Reference                       `json:"location,omitempty"`
	Seriousness             *CodeableConcept                 `json:"seriousness,omitempty"`
	Outcome                 []CodeableConcept                `json:"outcome,omitempty"`
	Recorder                *Reference                       `json:"recorder,omitempty"`
	Participant             []AdverseEventParticipant        `json:"participant,omitempty"`
	Study                   []Reference                      `json:"study,omitempty"`
	ExpectedInResearchStudy *bool                            `json:"expectedInResearchStudy,omitempty"`
	SuspectEntity           []AdverseEventSuspectEntity      `json:"suspectEntity,omitempty"`
	ContributingFactor      []AdverseEventContributingFactor `json:"contributingFactor,omitempty"`
	PreventiveAction        []AdverseEventPreventiveAction   `json:"preventiveAction,omitempty"`
	MitigatingAction        []AdverseEventMitigatingAction   `json:"mitigatingAction,omitempty"`
	SupportingInfo          []AdverseEventSupportingInfo     `json:"supportingInfo,omitempty"`
	Note                    []Annotation                     `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AdverseEvent
type AdverseEventParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AdverseEvent
type AdverseEventSuspectEntity struct {
	Id                      *string                             `json:"id,omitempty"`
	Extension               []Extension                         `json:"extension,omitempty"`
	ModifierExtension       []Extension                         `json:"modifierExtension,omitempty"`
	InstanceCodeableConcept CodeableConcept                     `json:"instanceCodeableConcept"`
	InstanceReference       Reference                           `json:"instanceReference"`
	Causality               *AdverseEventSuspectEntityCausality `json:"causality,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AdverseEvent
type AdverseEventSuspectEntityCausality struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	AssessmentMethod  *CodeableConcept `json:"assessmentMethod,omitempty"`
	EntityRelatedness *CodeableConcept `json:"entityRelatedness,omitempty"`
	Author            *Reference       `json:"author,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AdverseEvent
type AdverseEventContributingFactor struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	ItemReference       Reference       `json:"itemReference"`
	ItemCodeableConcept CodeableConcept `json:"itemCodeableConcept"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AdverseEvent
type AdverseEventPreventiveAction struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	ItemReference       Reference       `json:"itemReference"`
	ItemCodeableConcept CodeableConcept `json:"itemCodeableConcept"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AdverseEvent
type AdverseEventMitigatingAction struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	ItemReference       Reference       `json:"itemReference"`
	ItemCodeableConcept CodeableConcept `json:"itemCodeableConcept"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AdverseEvent
type AdverseEventSupportingInfo struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	ItemReference       Reference       `json:"itemReference"`
	ItemCodeableConcept CodeableConcept `json:"itemCodeableConcept"`
}

type OtherAdverseEvent AdverseEvent

// on convert struct to json, automatically add resourceType=AdverseEvent
func (r AdverseEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAdverseEvent
		ResourceType string `json:"resourceType"`
	}{
		OtherAdverseEvent: OtherAdverseEvent(r),
		ResourceType:      "AdverseEvent",
	})
}
func (r AdverseEvent) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "AdverseEvent/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "AdverseEvent"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *AdverseEvent) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAdverse_event_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Actuality(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAdverse_event_actuality

	if resource == nil {
		return CodeSelect("actuality", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("actuality", &resource.Actuality, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_OccurrenceDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("occurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("occurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *AdverseEvent) T_Detected(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("detected", nil, htmlAttrs)
	}
	return DateTimeInput("detected", resource.Detected, htmlAttrs)
}
func (resource *AdverseEvent) T_RecordedDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("recordedDate", nil, htmlAttrs)
	}
	return DateTimeInput("recordedDate", resource.RecordedDate, htmlAttrs)
}
func (resource *AdverseEvent) T_Seriousness(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("seriousness", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("seriousness", resource.Seriousness, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Outcome(numOutcome int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOutcome >= len(resource.Outcome) {
		return CodeableConceptSelect("outcome["+strconv.Itoa(numOutcome)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("outcome["+strconv.Itoa(numOutcome)+"]", &resource.Outcome[numOutcome], optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_ExpectedInResearchStudy(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("expectedInResearchStudy", nil, htmlAttrs)
	}
	return BoolInput("expectedInResearchStudy", resource.ExpectedInResearchStudy, htmlAttrs)
}
func (resource *AdverseEvent) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *AdverseEvent) T_ParticipantFunction(numParticipant int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].function", resource.Participant[numParticipant].Function, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_SuspectEntityInstanceCodeableConcept(numSuspectEntity int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSuspectEntity >= len(resource.SuspectEntity) {
		return CodeableConceptSelect("suspectEntity["+strconv.Itoa(numSuspectEntity)+"].instanceCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("suspectEntity["+strconv.Itoa(numSuspectEntity)+"].instanceCodeableConcept", &resource.SuspectEntity[numSuspectEntity].InstanceCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityAssessmentMethod(numSuspectEntity int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSuspectEntity >= len(resource.SuspectEntity) {
		return CodeableConceptSelect("suspectEntity["+strconv.Itoa(numSuspectEntity)+"].causality.assessmentMethod", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("suspectEntity["+strconv.Itoa(numSuspectEntity)+"].causality.assessmentMethod", resource.SuspectEntity[numSuspectEntity].Causality.AssessmentMethod, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityEntityRelatedness(numSuspectEntity int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSuspectEntity >= len(resource.SuspectEntity) {
		return CodeableConceptSelect("suspectEntity["+strconv.Itoa(numSuspectEntity)+"].causality.entityRelatedness", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("suspectEntity["+strconv.Itoa(numSuspectEntity)+"].causality.entityRelatedness", resource.SuspectEntity[numSuspectEntity].Causality.EntityRelatedness, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_ContributingFactorItemCodeableConcept(numContributingFactor int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContributingFactor >= len(resource.ContributingFactor) {
		return CodeableConceptSelect("contributingFactor["+strconv.Itoa(numContributingFactor)+"].itemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("contributingFactor["+strconv.Itoa(numContributingFactor)+"].itemCodeableConcept", &resource.ContributingFactor[numContributingFactor].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_PreventiveActionItemCodeableConcept(numPreventiveAction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPreventiveAction >= len(resource.PreventiveAction) {
		return CodeableConceptSelect("preventiveAction["+strconv.Itoa(numPreventiveAction)+"].itemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("preventiveAction["+strconv.Itoa(numPreventiveAction)+"].itemCodeableConcept", &resource.PreventiveAction[numPreventiveAction].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_MitigatingActionItemCodeableConcept(numMitigatingAction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMitigatingAction >= len(resource.MitigatingAction) {
		return CodeableConceptSelect("mitigatingAction["+strconv.Itoa(numMitigatingAction)+"].itemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("mitigatingAction["+strconv.Itoa(numMitigatingAction)+"].itemCodeableConcept", &resource.MitigatingAction[numMitigatingAction].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_SupportingInfoItemCodeableConcept(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodeableConceptSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].itemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].itemCodeableConcept", &resource.SupportingInfo[numSupportingInfo].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
