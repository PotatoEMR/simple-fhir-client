//generated August 14 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

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
	OccurrenceDateTime      *string                          `json:"occurrenceDateTime"`
	OccurrencePeriod        *Period                          `json:"occurrencePeriod"`
	OccurrenceTiming        *Timing                          `json:"occurrenceTiming"`
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
