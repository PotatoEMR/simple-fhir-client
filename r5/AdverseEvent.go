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
	OccurrenceDateTime      *time.Time                       `json:"occurrenceDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	OccurrencePeriod        *Period                          `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming        *Timing                          `json:"occurrenceTiming,omitempty"`
	Detected                *time.Time                       `json:"detected,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	RecordedDate            *time.Time                       `json:"recordedDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *AdverseEvent) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSAdverse_event_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Actuality(htmlAttrs string) templ.Component {
	optionsValueSet := VSAdverse_event_actuality

	if resource == nil {
		return CodeSelect("Actuality", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Actuality", &resource.Actuality, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_OccurrenceDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("OccurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("OccurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *AdverseEvent) T_Detected(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Detected", nil, htmlAttrs)
	}
	return DateTimeInput("Detected", resource.Detected, htmlAttrs)
}
func (resource *AdverseEvent) T_RecordedDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("RecordedDate", nil, htmlAttrs)
	}
	return DateTimeInput("RecordedDate", resource.RecordedDate, htmlAttrs)
}
func (resource *AdverseEvent) T_Seriousness(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Seriousness", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Seriousness", resource.Seriousness, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Outcome(numOutcome int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOutcome >= len(resource.Outcome) {
		return CodeableConceptSelect("Outcome["+strconv.Itoa(numOutcome)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Outcome["+strconv.Itoa(numOutcome)+"]", &resource.Outcome[numOutcome], optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_ExpectedInResearchStudy(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("ExpectedInResearchStudy", nil, htmlAttrs)
	}
	return BoolInput("ExpectedInResearchStudy", resource.ExpectedInResearchStudy, htmlAttrs)
}
func (resource *AdverseEvent) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *AdverseEvent) T_ParticipantFunction(numParticipant int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeableConceptSelect("Participant["+strconv.Itoa(numParticipant)+"]Function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Participant["+strconv.Itoa(numParticipant)+"]Function", resource.Participant[numParticipant].Function, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_SuspectEntityInstanceCodeableConcept(numSuspectEntity int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSuspectEntity >= len(resource.SuspectEntity) {
		return CodeableConceptSelect("SuspectEntity["+strconv.Itoa(numSuspectEntity)+"]InstanceCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SuspectEntity["+strconv.Itoa(numSuspectEntity)+"]InstanceCodeableConcept", &resource.SuspectEntity[numSuspectEntity].InstanceCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityAssessmentMethod(numSuspectEntity int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSuspectEntity >= len(resource.SuspectEntity) {
		return CodeableConceptSelect("SuspectEntity["+strconv.Itoa(numSuspectEntity)+"]Causality.AssessmentMethod", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SuspectEntity["+strconv.Itoa(numSuspectEntity)+"]Causality.AssessmentMethod", resource.SuspectEntity[numSuspectEntity].Causality.AssessmentMethod, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityEntityRelatedness(numSuspectEntity int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSuspectEntity >= len(resource.SuspectEntity) {
		return CodeableConceptSelect("SuspectEntity["+strconv.Itoa(numSuspectEntity)+"]Causality.EntityRelatedness", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SuspectEntity["+strconv.Itoa(numSuspectEntity)+"]Causality.EntityRelatedness", resource.SuspectEntity[numSuspectEntity].Causality.EntityRelatedness, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_ContributingFactorItemCodeableConcept(numContributingFactor int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numContributingFactor >= len(resource.ContributingFactor) {
		return CodeableConceptSelect("ContributingFactor["+strconv.Itoa(numContributingFactor)+"]ItemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ContributingFactor["+strconv.Itoa(numContributingFactor)+"]ItemCodeableConcept", &resource.ContributingFactor[numContributingFactor].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_PreventiveActionItemCodeableConcept(numPreventiveAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPreventiveAction >= len(resource.PreventiveAction) {
		return CodeableConceptSelect("PreventiveAction["+strconv.Itoa(numPreventiveAction)+"]ItemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PreventiveAction["+strconv.Itoa(numPreventiveAction)+"]ItemCodeableConcept", &resource.PreventiveAction[numPreventiveAction].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_MitigatingActionItemCodeableConcept(numMitigatingAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMitigatingAction >= len(resource.MitigatingAction) {
		return CodeableConceptSelect("MitigatingAction["+strconv.Itoa(numMitigatingAction)+"]ItemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MitigatingAction["+strconv.Itoa(numMitigatingAction)+"]ItemCodeableConcept", &resource.MitigatingAction[numMitigatingAction].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_SupportingInfoItemCodeableConcept(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodeableConceptSelect("SupportingInfo["+strconv.Itoa(numSupportingInfo)+"]ItemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SupportingInfo["+strconv.Itoa(numSupportingInfo)+"]ItemCodeableConcept", &resource.SupportingInfo[numSupportingInfo].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
