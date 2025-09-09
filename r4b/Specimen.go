package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Specimen
type Specimen struct {
	Id                  *string              `json:"id,omitempty"`
	Meta                *Meta                `json:"meta,omitempty"`
	ImplicitRules       *string              `json:"implicitRules,omitempty"`
	Language            *string              `json:"language,omitempty"`
	Text                *Narrative           `json:"text,omitempty"`
	Contained           []Resource           `json:"contained,omitempty"`
	Extension           []Extension          `json:"extension,omitempty"`
	ModifierExtension   []Extension          `json:"modifierExtension,omitempty"`
	Identifier          []Identifier         `json:"identifier,omitempty"`
	AccessionIdentifier *Identifier          `json:"accessionIdentifier,omitempty"`
	Status              *string              `json:"status,omitempty"`
	Type                *CodeableConcept     `json:"type,omitempty"`
	Subject             *Reference           `json:"subject,omitempty"`
	ReceivedTime        *string              `json:"receivedTime,omitempty"`
	Parent              []Reference          `json:"parent,omitempty"`
	Request             []Reference          `json:"request,omitempty"`
	Collection          *SpecimenCollection  `json:"collection,omitempty"`
	Processing          []SpecimenProcessing `json:"processing,omitempty"`
	Container           []SpecimenContainer  `json:"container,omitempty"`
	Condition           []CodeableConcept    `json:"condition,omitempty"`
	Note                []Annotation         `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Specimen
type SpecimenCollection struct {
	Id                           *string          `json:"id,omitempty"`
	Extension                    []Extension      `json:"extension,omitempty"`
	ModifierExtension            []Extension      `json:"modifierExtension,omitempty"`
	Collector                    *Reference       `json:"collector,omitempty"`
	CollectedDateTime            *string          `json:"collectedDateTime,omitempty"`
	CollectedPeriod              *Period          `json:"collectedPeriod,omitempty"`
	Duration                     *Duration        `json:"duration,omitempty"`
	Quantity                     *Quantity        `json:"quantity,omitempty"`
	Method                       *CodeableConcept `json:"method,omitempty"`
	BodySite                     *CodeableConcept `json:"bodySite,omitempty"`
	FastingStatusCodeableConcept *CodeableConcept `json:"fastingStatusCodeableConcept,omitempty"`
	FastingStatusDuration        *Duration        `json:"fastingStatusDuration,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Specimen
type SpecimenProcessing struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Procedure         *CodeableConcept `json:"procedure,omitempty"`
	Additive          []Reference      `json:"additive,omitempty"`
	TimeDateTime      *string          `json:"timeDateTime,omitempty"`
	TimePeriod        *Period          `json:"timePeriod,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Specimen
type SpecimenContainer struct {
	Id                      *string          `json:"id,omitempty"`
	Extension               []Extension      `json:"extension,omitempty"`
	ModifierExtension       []Extension      `json:"modifierExtension,omitempty"`
	Identifier              []Identifier     `json:"identifier,omitempty"`
	Description             *string          `json:"description,omitempty"`
	Type                    *CodeableConcept `json:"type,omitempty"`
	Capacity                *Quantity        `json:"capacity,omitempty"`
	SpecimenQuantity        *Quantity        `json:"specimenQuantity,omitempty"`
	AdditiveCodeableConcept *CodeableConcept `json:"additiveCodeableConcept,omitempty"`
	AdditiveReference       *Reference       `json:"additiveReference,omitempty"`
}

type OtherSpecimen Specimen

// on convert struct to json, automatically add resourceType=Specimen
func (r Specimen) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSpecimen
		ResourceType string `json:"resourceType"`
	}{
		OtherSpecimen: OtherSpecimen(r),
		ResourceType:  "Specimen",
	})
}
func (r Specimen) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Specimen/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Specimen"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Specimen) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSSpecimen_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_ReceivedTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("receivedTime", nil, htmlAttrs)
	}
	return DateTimeInput("receivedTime", resource.ReceivedTime, htmlAttrs)
}
func (resource *Specimen) T_Condition(numCondition int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return CodeableConceptSelect("condition["+strconv.Itoa(numCondition)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("condition["+strconv.Itoa(numCondition)+"]", &resource.Condition[numCondition], optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Specimen) T_CollectionCollectedDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("collection.collectedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("collection.collectedDateTime", resource.Collection.CollectedDateTime, htmlAttrs)
}
func (resource *Specimen) T_CollectionMethod(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("collection.method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("collection.method", resource.Collection.Method, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_CollectionBodySite(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("collection.bodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("collection.bodySite", resource.Collection.BodySite, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_CollectionFastingStatusCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("collection.fastingStatusCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("collection.fastingStatusCodeableConcept", resource.Collection.FastingStatusCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_ProcessingDescription(numProcessing int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return StringInput("processing["+strconv.Itoa(numProcessing)+"].description", nil, htmlAttrs)
	}
	return StringInput("processing["+strconv.Itoa(numProcessing)+"].description", resource.Processing[numProcessing].Description, htmlAttrs)
}
func (resource *Specimen) T_ProcessingProcedure(numProcessing int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return CodeableConceptSelect("processing["+strconv.Itoa(numProcessing)+"].procedure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("processing["+strconv.Itoa(numProcessing)+"].procedure", resource.Processing[numProcessing].Procedure, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_ProcessingTimeDateTime(numProcessing int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return DateTimeInput("processing["+strconv.Itoa(numProcessing)+"].timeDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("processing["+strconv.Itoa(numProcessing)+"].timeDateTime", resource.Processing[numProcessing].TimeDateTime, htmlAttrs)
}
func (resource *Specimen) T_ContainerDescription(numContainer int, htmlAttrs string) templ.Component {
	if resource == nil || numContainer >= len(resource.Container) {
		return StringInput("container["+strconv.Itoa(numContainer)+"].description", nil, htmlAttrs)
	}
	return StringInput("container["+strconv.Itoa(numContainer)+"].description", resource.Container[numContainer].Description, htmlAttrs)
}
func (resource *Specimen) T_ContainerType(numContainer int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numContainer >= len(resource.Container) {
		return CodeableConceptSelect("container["+strconv.Itoa(numContainer)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("container["+strconv.Itoa(numContainer)+"].type", resource.Container[numContainer].Type, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_ContainerAdditiveCodeableConcept(numContainer int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numContainer >= len(resource.Container) {
		return CodeableConceptSelect("container["+strconv.Itoa(numContainer)+"].additiveCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("container["+strconv.Itoa(numContainer)+"].additiveCodeableConcept", resource.Container[numContainer].AdditiveCodeableConcept, optionsValueSet, htmlAttrs)
}
