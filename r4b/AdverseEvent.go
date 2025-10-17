package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
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
	Date                  *FhirDateTime               `json:"date,omitempty"`
	Detected              *FhirDateTime               `json:"detected,omitempty"`
	RecordedDate          *FhirDateTime               `json:"recordedDate,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
func (r AdverseEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAdverseEvent
		ResourceType string `json:"resourceType"`
	}{
		OtherAdverseEvent: OtherAdverseEvent(r),
		ResourceType:      "AdverseEvent",
	})
}

// json -> struct, first reject if resourceType != AdverseEvent
func (r *AdverseEvent) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "AdverseEvent" {
		return errors.New("resourceType not AdverseEvent")
	}
	return json.Unmarshal(data, (*OtherAdverseEvent)(r))
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
func (resource *AdverseEvent) T_Event(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("event", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("event", resource.Event, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *AdverseEvent) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *AdverseEvent) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *AdverseEvent) T_Detected(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("detected", nil, htmlAttrs)
	}
	return FhirDateTimeInput("detected", resource.Detected, htmlAttrs)
}
func (resource *AdverseEvent) T_RecordedDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("recordedDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("recordedDate", resource.RecordedDate, htmlAttrs)
}
func (resource *AdverseEvent) T_ResultingCondition(frs []FhirResource, numResultingCondition int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResultingCondition >= len(resource.ResultingCondition) {
		return ReferenceInput(frs, "resultingCondition["+strconv.Itoa(numResultingCondition)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "resultingCondition["+strconv.Itoa(numResultingCondition)+"]", &resource.ResultingCondition[numResultingCondition], htmlAttrs)
}
func (resource *AdverseEvent) T_Location(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "location", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "location", resource.Location, htmlAttrs)
}
func (resource *AdverseEvent) T_Seriousness(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("seriousness", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("seriousness", resource.Seriousness, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Severity(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAdverse_event_severity

	if resource == nil {
		return CodeableConceptSelect("severity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("severity", resource.Severity, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Outcome(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAdverse_event_outcome

	if resource == nil {
		return CodeableConceptSelect("outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("outcome", resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_Recorder(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "recorder", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "recorder", resource.Recorder, htmlAttrs)
}
func (resource *AdverseEvent) T_Contributor(frs []FhirResource, numContributor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContributor >= len(resource.Contributor) {
		return ReferenceInput(frs, "contributor["+strconv.Itoa(numContributor)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "contributor["+strconv.Itoa(numContributor)+"]", &resource.Contributor[numContributor], htmlAttrs)
}
func (resource *AdverseEvent) T_SubjectMedicalHistory(frs []FhirResource, numSubjectMedicalHistory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubjectMedicalHistory >= len(resource.SubjectMedicalHistory) {
		return ReferenceInput(frs, "subjectMedicalHistory["+strconv.Itoa(numSubjectMedicalHistory)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subjectMedicalHistory["+strconv.Itoa(numSubjectMedicalHistory)+"]", &resource.SubjectMedicalHistory[numSubjectMedicalHistory], htmlAttrs)
}
func (resource *AdverseEvent) T_ReferenceDocument(frs []FhirResource, numReferenceDocument int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReferenceDocument >= len(resource.ReferenceDocument) {
		return ReferenceInput(frs, "referenceDocument["+strconv.Itoa(numReferenceDocument)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "referenceDocument["+strconv.Itoa(numReferenceDocument)+"]", &resource.ReferenceDocument[numReferenceDocument], htmlAttrs)
}
func (resource *AdverseEvent) T_Study(frs []FhirResource, numStudy int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStudy >= len(resource.Study) {
		return ReferenceInput(frs, "study["+strconv.Itoa(numStudy)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "study["+strconv.Itoa(numStudy)+"]", &resource.Study[numStudy], htmlAttrs)
}
func (resource *AdverseEvent) T_SuspectEntityInstance(frs []FhirResource, numSuspectEntity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSuspectEntity >= len(resource.SuspectEntity) {
		return ReferenceInput(frs, "suspectEntity["+strconv.Itoa(numSuspectEntity)+"].instance", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "suspectEntity["+strconv.Itoa(numSuspectEntity)+"].instance", &resource.SuspectEntity[numSuspectEntity].Instance, htmlAttrs)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityAssessment(numSuspectEntity int, numCausality int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSuspectEntity >= len(resource.SuspectEntity) || numCausality >= len(resource.SuspectEntity[numSuspectEntity].Causality) {
		return CodeableConceptSelect("suspectEntity["+strconv.Itoa(numSuspectEntity)+"].causality["+strconv.Itoa(numCausality)+"].assessment", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("suspectEntity["+strconv.Itoa(numSuspectEntity)+"].causality["+strconv.Itoa(numCausality)+"].assessment", resource.SuspectEntity[numSuspectEntity].Causality[numCausality].Assessment, optionsValueSet, htmlAttrs)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityProductRelatedness(numSuspectEntity int, numCausality int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSuspectEntity >= len(resource.SuspectEntity) || numCausality >= len(resource.SuspectEntity[numSuspectEntity].Causality) {
		return StringInput("suspectEntity["+strconv.Itoa(numSuspectEntity)+"].causality["+strconv.Itoa(numCausality)+"].productRelatedness", nil, htmlAttrs)
	}
	return StringInput("suspectEntity["+strconv.Itoa(numSuspectEntity)+"].causality["+strconv.Itoa(numCausality)+"].productRelatedness", resource.SuspectEntity[numSuspectEntity].Causality[numCausality].ProductRelatedness, htmlAttrs)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityAuthor(frs []FhirResource, numSuspectEntity int, numCausality int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSuspectEntity >= len(resource.SuspectEntity) || numCausality >= len(resource.SuspectEntity[numSuspectEntity].Causality) {
		return ReferenceInput(frs, "suspectEntity["+strconv.Itoa(numSuspectEntity)+"].causality["+strconv.Itoa(numCausality)+"].author", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "suspectEntity["+strconv.Itoa(numSuspectEntity)+"].causality["+strconv.Itoa(numCausality)+"].author", resource.SuspectEntity[numSuspectEntity].Causality[numCausality].Author, htmlAttrs)
}
func (resource *AdverseEvent) T_SuspectEntityCausalityMethod(numSuspectEntity int, numCausality int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSuspectEntity >= len(resource.SuspectEntity) || numCausality >= len(resource.SuspectEntity[numSuspectEntity].Causality) {
		return CodeableConceptSelect("suspectEntity["+strconv.Itoa(numSuspectEntity)+"].causality["+strconv.Itoa(numCausality)+"].method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("suspectEntity["+strconv.Itoa(numSuspectEntity)+"].causality["+strconv.Itoa(numCausality)+"].method", resource.SuspectEntity[numSuspectEntity].Causality[numCausality].Method, optionsValueSet, htmlAttrs)
}
