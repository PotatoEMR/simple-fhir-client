package r5

//generated with command go run ./bultaoreune -nodownload
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
	OccurrenceDateTime      *string                          `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod        *Period                          `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming        *Timing                          `json:"occurrenceTiming,omitempty"`
	Detected                *string                          `json:"detected,omitempty"`
	RecordedDate            *string                          `json:"recordedDate,omitempty"`
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

func (resource *AdverseEvent) T_Id() templ.Component {

	if resource == nil {
		return StringInput("AdverseEvent.Id", nil)
	}
	return StringInput("AdverseEvent.Id", resource.Id)
}
func (resource *AdverseEvent) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("AdverseEvent.ImplicitRules", nil)
	}
	return StringInput("AdverseEvent.ImplicitRules", resource.ImplicitRules)
}
func (resource *AdverseEvent) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("AdverseEvent.Language", nil, optionsValueSet)
	}
	return CodeSelect("AdverseEvent.Language", resource.Language, optionsValueSet)
}
func (resource *AdverseEvent) T_Status() templ.Component {
	optionsValueSet := VSAdverse_event_status

	if resource == nil {
		return CodeSelect("AdverseEvent.Status", nil, optionsValueSet)
	}
	return CodeSelect("AdverseEvent.Status", &resource.Status, optionsValueSet)
}
func (resource *AdverseEvent) T_Actuality() templ.Component {
	optionsValueSet := VSAdverse_event_actuality

	if resource == nil {
		return CodeSelect("AdverseEvent.Actuality", nil, optionsValueSet)
	}
	return CodeSelect("AdverseEvent.Actuality", &resource.Actuality, optionsValueSet)
}
func (resource *AdverseEvent) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("AdverseEvent.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdverseEvent.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *AdverseEvent) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("AdverseEvent.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdverseEvent.Code", resource.Code, optionsValueSet)
}
func (resource *AdverseEvent) T_Detected() templ.Component {

	if resource == nil {
		return StringInput("AdverseEvent.Detected", nil)
	}
	return StringInput("AdverseEvent.Detected", resource.Detected)
}
func (resource *AdverseEvent) T_RecordedDate() templ.Component {

	if resource == nil {
		return StringInput("AdverseEvent.RecordedDate", nil)
	}
	return StringInput("AdverseEvent.RecordedDate", resource.RecordedDate)
}
func (resource *AdverseEvent) T_Seriousness(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("AdverseEvent.Seriousness", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdverseEvent.Seriousness", resource.Seriousness, optionsValueSet)
}
func (resource *AdverseEvent) T_Outcome(numOutcome int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Outcome) >= numOutcome {
		return CodeableConceptSelect("AdverseEvent.Outcome["+strconv.Itoa(numOutcome)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdverseEvent.Outcome["+strconv.Itoa(numOutcome)+"]", &resource.Outcome[numOutcome], optionsValueSet)
}
func (resource *AdverseEvent) T_ExpectedInResearchStudy() templ.Component {

	if resource == nil {
		return BoolInput("AdverseEvent.ExpectedInResearchStudy", nil)
	}
	return BoolInput("AdverseEvent.ExpectedInResearchStudy", resource.ExpectedInResearchStudy)
}
func (resource *AdverseEvent) T_ParticipantId(numParticipant int) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return StringInput("AdverseEvent.Participant["+strconv.Itoa(numParticipant)+"].Id", nil)
	}
	return StringInput("AdverseEvent.Participant["+strconv.Itoa(numParticipant)+"].Id", resource.Participant[numParticipant].Id)
}
func (resource *AdverseEvent) T_ParticipantFunction(numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return CodeableConceptSelect("AdverseEvent.Participant["+strconv.Itoa(numParticipant)+"].Function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdverseEvent.Participant["+strconv.Itoa(numParticipant)+"].Function", resource.Participant[numParticipant].Function, optionsValueSet)
}
func (resource *AdverseEvent) T_SuspectEntityId(numSuspectEntity int) templ.Component {

	if resource == nil || len(resource.SuspectEntity) >= numSuspectEntity {
		return StringInput("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Id", nil)
	}
	return StringInput("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Id", resource.SuspectEntity[numSuspectEntity].Id)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityId(numSuspectEntity int) templ.Component {

	if resource == nil || len(resource.SuspectEntity) >= numSuspectEntity {
		return StringInput("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Causality.Id", nil)
	}
	return StringInput("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Causality.Id", resource.SuspectEntity[numSuspectEntity].Causality.Id)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityAssessmentMethod(numSuspectEntity int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SuspectEntity) >= numSuspectEntity {
		return CodeableConceptSelect("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Causality.AssessmentMethod", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Causality.AssessmentMethod", resource.SuspectEntity[numSuspectEntity].Causality.AssessmentMethod, optionsValueSet)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityEntityRelatedness(numSuspectEntity int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SuspectEntity) >= numSuspectEntity {
		return CodeableConceptSelect("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Causality.EntityRelatedness", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Causality.EntityRelatedness", resource.SuspectEntity[numSuspectEntity].Causality.EntityRelatedness, optionsValueSet)
}
func (resource *AdverseEvent) T_ContributingFactorId(numContributingFactor int) templ.Component {

	if resource == nil || len(resource.ContributingFactor) >= numContributingFactor {
		return StringInput("AdverseEvent.ContributingFactor["+strconv.Itoa(numContributingFactor)+"].Id", nil)
	}
	return StringInput("AdverseEvent.ContributingFactor["+strconv.Itoa(numContributingFactor)+"].Id", resource.ContributingFactor[numContributingFactor].Id)
}
func (resource *AdverseEvent) T_PreventiveActionId(numPreventiveAction int) templ.Component {

	if resource == nil || len(resource.PreventiveAction) >= numPreventiveAction {
		return StringInput("AdverseEvent.PreventiveAction["+strconv.Itoa(numPreventiveAction)+"].Id", nil)
	}
	return StringInput("AdverseEvent.PreventiveAction["+strconv.Itoa(numPreventiveAction)+"].Id", resource.PreventiveAction[numPreventiveAction].Id)
}
func (resource *AdverseEvent) T_MitigatingActionId(numMitigatingAction int) templ.Component {

	if resource == nil || len(resource.MitigatingAction) >= numMitigatingAction {
		return StringInput("AdverseEvent.MitigatingAction["+strconv.Itoa(numMitigatingAction)+"].Id", nil)
	}
	return StringInput("AdverseEvent.MitigatingAction["+strconv.Itoa(numMitigatingAction)+"].Id", resource.MitigatingAction[numMitigatingAction].Id)
}
func (resource *AdverseEvent) T_SupportingInfoId(numSupportingInfo int) templ.Component {

	if resource == nil || len(resource.SupportingInfo) >= numSupportingInfo {
		return StringInput("AdverseEvent.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Id", nil)
	}
	return StringInput("AdverseEvent.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Id", resource.SupportingInfo[numSupportingInfo].Id)
}
