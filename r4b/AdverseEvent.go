package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/AdverseEvent
type AdverseEvent struct {
	Id                    *string                     `json:"id,omitempty"`
	Meta                  *Meta                       `json:"meta,omitempty"`
	ImplicitRules         *string                     `json:"implicitRules,omitempty"`
	Language              *string                     `json:"language,omitempty"`
	Text                  *Narrative                  `json:"text,omitempty"`
	Contained             []Resource                  `json:"contained,omitempty"`
	Extension             []Extension                 `json:"extension,omitempty"`
	ModifierExtension     []Extension                 `json:"modifierExtension,omitempty"`
	Identifier            *Identifier                 `json:"identifier,omitempty"`
	Actuality             string                      `json:"actuality"`
	Category              []CodeableConcept           `json:"category,omitempty"`
	Event                 *CodeableConcept            `json:"event,omitempty"`
	Subject               Reference                   `json:"subject"`
	Encounter             *Reference                  `json:"encounter,omitempty"`
	Date                  *string                     `json:"date,omitempty"`
	Detected              *string                     `json:"detected,omitempty"`
	RecordedDate          *string                     `json:"recordedDate,omitempty"`
	ResultingCondition    []Reference                 `json:"resultingCondition,omitempty"`
	Location              *Reference                  `json:"location,omitempty"`
	Seriousness           *CodeableConcept            `json:"seriousness,omitempty"`
	Severity              *CodeableConcept            `json:"severity,omitempty"`
	Outcome               *CodeableConcept            `json:"outcome,omitempty"`
	Recorder              *Reference                  `json:"recorder,omitempty"`
	Contributor           []Reference                 `json:"contributor,omitempty"`
	SuspectEntity         []AdverseEventSuspectEntity `json:"suspectEntity,omitempty"`
	SubjectMedicalHistory []Reference                 `json:"subjectMedicalHistory,omitempty"`
	ReferenceDocument     []Reference                 `json:"referenceDocument,omitempty"`
	Study                 []Reference                 `json:"study,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/AdverseEvent
type AdverseEventSuspectEntity struct {
	Id                *string                              `json:"id,omitempty"`
	Extension         []Extension                          `json:"extension,omitempty"`
	ModifierExtension []Extension                          `json:"modifierExtension,omitempty"`
	Instance          Reference                            `json:"instance"`
	Causality         []AdverseEventSuspectEntityCausality `json:"causality,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/AdverseEvent
type AdverseEventSuspectEntityCausality struct {
	Id                 *string          `json:"id,omitempty"`
	Extension          []Extension      `json:"extension,omitempty"`
	ModifierExtension  []Extension      `json:"modifierExtension,omitempty"`
	Assessment         *CodeableConcept `json:"assessment,omitempty"`
	ProductRelatedness *string          `json:"productRelatedness,omitempty"`
	Author             *Reference       `json:"author,omitempty"`
	Method             *CodeableConcept `json:"method,omitempty"`
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
func (resource *AdverseEvent) T_Event(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("AdverseEvent.Event", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdverseEvent.Event", resource.Event, optionsValueSet)
}
func (resource *AdverseEvent) T_Date() templ.Component {

	if resource == nil {
		return StringInput("AdverseEvent.Date", nil)
	}
	return StringInput("AdverseEvent.Date", resource.Date)
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
func (resource *AdverseEvent) T_Severity() templ.Component {
	optionsValueSet := VSAdverse_event_severity

	if resource == nil {
		return CodeableConceptSelect("AdverseEvent.Severity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdverseEvent.Severity", resource.Severity, optionsValueSet)
}
func (resource *AdverseEvent) T_Outcome() templ.Component {
	optionsValueSet := VSAdverse_event_outcome

	if resource == nil {
		return CodeableConceptSelect("AdverseEvent.Outcome", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdverseEvent.Outcome", resource.Outcome, optionsValueSet)
}
func (resource *AdverseEvent) T_SuspectEntityId(numSuspectEntity int) templ.Component {

	if resource == nil || len(resource.SuspectEntity) >= numSuspectEntity {
		return StringInput("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Id", nil)
	}
	return StringInput("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Id", resource.SuspectEntity[numSuspectEntity].Id)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityId(numSuspectEntity int, numCausality int) templ.Component {

	if resource == nil || len(resource.SuspectEntity) >= numSuspectEntity || len(resource.SuspectEntity[numSuspectEntity].Causality) >= numCausality {
		return StringInput("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Causality["+strconv.Itoa(numCausality)+"].Id", nil)
	}
	return StringInput("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Causality["+strconv.Itoa(numCausality)+"].Id", resource.SuspectEntity[numSuspectEntity].Causality[numCausality].Id)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityAssessment(numSuspectEntity int, numCausality int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SuspectEntity) >= numSuspectEntity || len(resource.SuspectEntity[numSuspectEntity].Causality) >= numCausality {
		return CodeableConceptSelect("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Causality["+strconv.Itoa(numCausality)+"].Assessment", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Causality["+strconv.Itoa(numCausality)+"].Assessment", resource.SuspectEntity[numSuspectEntity].Causality[numCausality].Assessment, optionsValueSet)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityProductRelatedness(numSuspectEntity int, numCausality int) templ.Component {

	if resource == nil || len(resource.SuspectEntity) >= numSuspectEntity || len(resource.SuspectEntity[numSuspectEntity].Causality) >= numCausality {
		return StringInput("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Causality["+strconv.Itoa(numCausality)+"].ProductRelatedness", nil)
	}
	return StringInput("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Causality["+strconv.Itoa(numCausality)+"].ProductRelatedness", resource.SuspectEntity[numSuspectEntity].Causality[numCausality].ProductRelatedness)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityMethod(numSuspectEntity int, numCausality int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SuspectEntity) >= numSuspectEntity || len(resource.SuspectEntity[numSuspectEntity].Causality) >= numCausality {
		return CodeableConceptSelect("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Causality["+strconv.Itoa(numCausality)+"].Method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AdverseEvent.SuspectEntity["+strconv.Itoa(numSuspectEntity)+"].Causality["+strconv.Itoa(numCausality)+"].Method", resource.SuspectEntity[numSuspectEntity].Causality[numCausality].Method, optionsValueSet)
}
