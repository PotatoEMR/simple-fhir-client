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
	Date                  *time.Time                  `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Detected              *time.Time                  `json:"detected,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	RecordedDate          *time.Time                  `json:"recordedDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r AdverseEvent) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "AdverseEvent/" + *r.Id
		ref.Reference = &refStr
	}
	ref.Identifier = r.Identifier
	rtype := "AdverseEvent"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *AdverseEvent) T_Actuality(htmlAttrs string) templ.Component {
	optionsValueSet := VSAdverse_event_actuality

	if resource == nil {
		return CodeSelect("AdverseEvent.Actuality", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("AdverseEvent.Actuality", &resource.Actuality, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("AdverseEvent.Category."+strconv.Itoa(numCategory)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AdverseEvent.Category."+strconv.Itoa(numCategory)+".", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Event(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("AdverseEvent.Event", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AdverseEvent.Event", resource.Event, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Date(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("AdverseEvent.Date", nil, htmlAttrs)
	}
	return DateTimeInput("AdverseEvent.Date", resource.Date, htmlAttrs)
}
func (resource *AdverseEvent) T_Detected(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("AdverseEvent.Detected", nil, htmlAttrs)
	}
	return DateTimeInput("AdverseEvent.Detected", resource.Detected, htmlAttrs)
}
func (resource *AdverseEvent) T_RecordedDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("AdverseEvent.RecordedDate", nil, htmlAttrs)
	}
	return DateTimeInput("AdverseEvent.RecordedDate", resource.RecordedDate, htmlAttrs)
}
func (resource *AdverseEvent) T_Seriousness(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("AdverseEvent.Seriousness", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AdverseEvent.Seriousness", resource.Seriousness, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Severity(htmlAttrs string) templ.Component {
	optionsValueSet := VSAdverse_event_severity

	if resource == nil {
		return CodeableConceptSelect("AdverseEvent.Severity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AdverseEvent.Severity", resource.Severity, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Outcome(htmlAttrs string) templ.Component {
	optionsValueSet := VSAdverse_event_outcome

	if resource == nil {
		return CodeableConceptSelect("AdverseEvent.Outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AdverseEvent.Outcome", resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityAssessment(numSuspectEntity int, numCausality int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSuspectEntity >= len(resource.SuspectEntity) || numCausality >= len(resource.SuspectEntity[numSuspectEntity].Causality) {
		return CodeableConceptSelect("AdverseEvent.SuspectEntity."+strconv.Itoa(numSuspectEntity)+"..Causality."+strconv.Itoa(numCausality)+"..Assessment", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AdverseEvent.SuspectEntity."+strconv.Itoa(numSuspectEntity)+"..Causality."+strconv.Itoa(numCausality)+"..Assessment", resource.SuspectEntity[numSuspectEntity].Causality[numCausality].Assessment, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityProductRelatedness(numSuspectEntity int, numCausality int, htmlAttrs string) templ.Component {

	if resource == nil || numSuspectEntity >= len(resource.SuspectEntity) || numCausality >= len(resource.SuspectEntity[numSuspectEntity].Causality) {
		return StringInput("AdverseEvent.SuspectEntity."+strconv.Itoa(numSuspectEntity)+"..Causality."+strconv.Itoa(numCausality)+"..ProductRelatedness", nil, htmlAttrs)
	}
	return StringInput("AdverseEvent.SuspectEntity."+strconv.Itoa(numSuspectEntity)+"..Causality."+strconv.Itoa(numCausality)+"..ProductRelatedness", resource.SuspectEntity[numSuspectEntity].Causality[numCausality].ProductRelatedness, htmlAttrs)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityMethod(numSuspectEntity int, numCausality int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSuspectEntity >= len(resource.SuspectEntity) || numCausality >= len(resource.SuspectEntity[numSuspectEntity].Causality) {
		return CodeableConceptSelect("AdverseEvent.SuspectEntity."+strconv.Itoa(numSuspectEntity)+"..Causality."+strconv.Itoa(numCausality)+"..Method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AdverseEvent.SuspectEntity."+strconv.Itoa(numSuspectEntity)+"..Causality."+strconv.Itoa(numCausality)+"..Method", resource.SuspectEntity[numSuspectEntity].Causality[numCausality].Method, optionsValueSet, htmlAttrs)
}
