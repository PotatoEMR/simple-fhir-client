package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *AdverseEvent) AdverseEventLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *AdverseEvent) AdverseEventStatus() templ.Component {
	optionsValueSet := VSAdverse_event_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *AdverseEvent) AdverseEventActuality() templ.Component {
	optionsValueSet := VSAdverse_event_actuality

	if resource != nil {
		return CodeSelect("actuality", nil, optionsValueSet)
	}
	return CodeSelect("actuality", &resource.Actuality, optionsValueSet)
}
func (resource *AdverseEvent) AdverseEventCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *AdverseEvent) AdverseEventCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet)
}
func (resource *AdverseEvent) AdverseEventSeriousness(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("seriousness", nil, optionsValueSet)
	}
	return CodeableConceptSelect("seriousness", resource.Seriousness, optionsValueSet)
}
func (resource *AdverseEvent) AdverseEventOutcome(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("outcome", nil, optionsValueSet)
	}
	return CodeableConceptSelect("outcome", &resource.Outcome[0], optionsValueSet)
}
func (resource *AdverseEvent) AdverseEventParticipantFunction(numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Participant) >= numParticipant {
		return CodeableConceptSelect("function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("function", resource.Participant[numParticipant].Function, optionsValueSet)
}
func (resource *AdverseEvent) AdverseEventSuspectEntityCausalityAssessmentMethod(numSuspectEntity int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.SuspectEntity) >= numSuspectEntity {
		return CodeableConceptSelect("assessmentMethod", nil, optionsValueSet)
	}
	return CodeableConceptSelect("assessmentMethod", resource.SuspectEntity[numSuspectEntity].Causality.AssessmentMethod, optionsValueSet)
}
func (resource *AdverseEvent) AdverseEventSuspectEntityCausalityEntityRelatedness(numSuspectEntity int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.SuspectEntity) >= numSuspectEntity {
		return CodeableConceptSelect("entityRelatedness", nil, optionsValueSet)
	}
	return CodeableConceptSelect("entityRelatedness", resource.SuspectEntity[numSuspectEntity].Causality.EntityRelatedness, optionsValueSet)
}
