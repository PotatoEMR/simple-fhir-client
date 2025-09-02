package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *AdverseEvent) AdverseEventLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
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
func (resource *AdverseEvent) AdverseEventEvent(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("event", nil, optionsValueSet)
	}
	return CodeableConceptSelect("event", resource.Event, optionsValueSet)
}
func (resource *AdverseEvent) AdverseEventSeriousness(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("seriousness", nil, optionsValueSet)
	}
	return CodeableConceptSelect("seriousness", resource.Seriousness, optionsValueSet)
}
func (resource *AdverseEvent) AdverseEventSeverity() templ.Component {
	optionsValueSet := VSAdverse_event_severity

	if resource != nil {
		return CodeableConceptSelect("severity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("severity", resource.Severity, optionsValueSet)
}
func (resource *AdverseEvent) AdverseEventOutcome() templ.Component {
	optionsValueSet := VSAdverse_event_outcome

	if resource != nil {
		return CodeableConceptSelect("outcome", nil, optionsValueSet)
	}
	return CodeableConceptSelect("outcome", resource.Outcome, optionsValueSet)
}
func (resource *AdverseEvent) AdverseEventSuspectEntityCausalityAssessment(numSuspectEntity int, numCausality int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.SuspectEntity[numSuspectEntity].Causality) >= numCausality {
		return CodeableConceptSelect("assessment", nil, optionsValueSet)
	}
	return CodeableConceptSelect("assessment", resource.SuspectEntity[numSuspectEntity].Causality[numCausality].Assessment, optionsValueSet)
}
func (resource *AdverseEvent) AdverseEventSuspectEntityCausalityMethod(numSuspectEntity int, numCausality int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.SuspectEntity[numSuspectEntity].Causality) >= numCausality {
		return CodeableConceptSelect("method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("method", resource.SuspectEntity[numSuspectEntity].Causality[numCausality].Method, optionsValueSet)
}
