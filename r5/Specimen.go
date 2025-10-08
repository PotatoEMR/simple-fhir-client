package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Specimen
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
	ReceivedTime        *FhirDateTime        `json:"receivedTime,omitempty"`
	Parent              []Reference          `json:"parent,omitempty"`
	Request             []Reference          `json:"request,omitempty"`
	Combined            *string              `json:"combined,omitempty"`
	Role                []CodeableConcept    `json:"role,omitempty"`
	Feature             []SpecimenFeature    `json:"feature,omitempty"`
	Collection          *SpecimenCollection  `json:"collection,omitempty"`
	Processing          []SpecimenProcessing `json:"processing,omitempty"`
	Container           []SpecimenContainer  `json:"container,omitempty"`
	Condition           []CodeableConcept    `json:"condition,omitempty"`
	Note                []Annotation         `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Specimen
type SpecimenFeature struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Description       string          `json:"description"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Specimen
type SpecimenCollection struct {
	Id                           *string            `json:"id,omitempty"`
	Extension                    []Extension        `json:"extension,omitempty"`
	ModifierExtension            []Extension        `json:"modifierExtension,omitempty"`
	Collector                    *Reference         `json:"collector,omitempty"`
	CollectedDateTime            *FhirDateTime      `json:"collectedDateTime,omitempty"`
	CollectedPeriod              *Period            `json:"collectedPeriod,omitempty"`
	Duration                     *Duration          `json:"duration,omitempty"`
	Quantity                     *Quantity          `json:"quantity,omitempty"`
	Method                       *CodeableConcept   `json:"method,omitempty"`
	Device                       *CodeableReference `json:"device,omitempty"`
	Procedure                    *Reference         `json:"procedure,omitempty"`
	BodySite                     *CodeableReference `json:"bodySite,omitempty"`
	FastingStatusCodeableConcept *CodeableConcept   `json:"fastingStatusCodeableConcept,omitempty"`
	FastingStatusDuration        *Duration          `json:"fastingStatusDuration,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Specimen
type SpecimenProcessing struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Method            *CodeableConcept `json:"method,omitempty"`
	Additive          []Reference      `json:"additive,omitempty"`
	TimeDateTime      *FhirDateTime    `json:"timeDateTime,omitempty"`
	TimePeriod        *Period          `json:"timePeriod,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Specimen
type SpecimenContainer struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Device            Reference   `json:"device"`
	Location          *Reference  `json:"location,omitempty"`
	SpecimenQuantity  *Quantity   `json:"specimenQuantity,omitempty"`
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
func (resource *Specimen) T_AccessionIdentifier(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IdentifierInput("accessionIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("accessionIdentifier", resource.AccessionIdentifier, htmlAttrs)
}
func (resource *Specimen) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSpecimen_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *Specimen) T_ReceivedTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("receivedTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("receivedTime", resource.ReceivedTime, htmlAttrs)
}
func (resource *Specimen) T_Parent(frs []FhirResource, numParent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParent >= len(resource.Parent) {
		return ReferenceInput(frs, "parent["+strconv.Itoa(numParent)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "parent["+strconv.Itoa(numParent)+"]", &resource.Parent[numParent], htmlAttrs)
}
func (resource *Specimen) T_Request(frs []FhirResource, numRequest int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRequest >= len(resource.Request) {
		return ReferenceInput(frs, "request["+strconv.Itoa(numRequest)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "request["+strconv.Itoa(numRequest)+"]", &resource.Request[numRequest], htmlAttrs)
}
func (resource *Specimen) T_Combined(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSpecimen_combined

	if resource == nil {
		return CodeSelect("combined", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("combined", resource.Combined, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_Role(numRole int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRole >= len(resource.Role) {
		return CodeableConceptSelect("role["+strconv.Itoa(numRole)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("role["+strconv.Itoa(numRole)+"]", &resource.Role[numRole], optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_Condition(numCondition int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return CodeableConceptSelect("condition["+strconv.Itoa(numCondition)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("condition["+strconv.Itoa(numCondition)+"]", &resource.Condition[numCondition], optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Specimen) T_FeatureType(numFeature int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFeature >= len(resource.Feature) {
		return CodeableConceptSelect("feature["+strconv.Itoa(numFeature)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("feature["+strconv.Itoa(numFeature)+"].type", &resource.Feature[numFeature].Type, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_FeatureDescription(numFeature int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFeature >= len(resource.Feature) {
		return StringInput("feature["+strconv.Itoa(numFeature)+"].description", nil, htmlAttrs)
	}
	return StringInput("feature["+strconv.Itoa(numFeature)+"].description", &resource.Feature[numFeature].Description, htmlAttrs)
}
func (resource *Specimen) T_CollectionCollector(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "collection.collector", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "collection.collector", resource.Collection.Collector, htmlAttrs)
}
func (resource *Specimen) T_CollectionCollectedDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("collection.collectedDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("collection.collectedDateTime", resource.Collection.CollectedDateTime, htmlAttrs)
}
func (resource *Specimen) T_CollectionCollectedPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("collection.collectedPeriod", nil, htmlAttrs)
	}
	return PeriodInput("collection.collectedPeriod", resource.Collection.CollectedPeriod, htmlAttrs)
}
func (resource *Specimen) T_CollectionDuration(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DurationInput("collection.duration", nil, htmlAttrs)
	}
	return DurationInput("collection.duration", resource.Collection.Duration, htmlAttrs)
}
func (resource *Specimen) T_CollectionQuantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("collection.quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("collection.quantity", resource.Collection.Quantity, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_CollectionMethod(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("collection.method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("collection.method", resource.Collection.Method, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_CollectionDevice(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("collection.device", nil, htmlAttrs)
	}
	return CodeableReferenceInput("collection.device", resource.Collection.Device, htmlAttrs)
}
func (resource *Specimen) T_CollectionProcedure(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "collection.procedure", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "collection.procedure", resource.Collection.Procedure, htmlAttrs)
}
func (resource *Specimen) T_CollectionBodySite(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("collection.bodySite", nil, htmlAttrs)
	}
	return CodeableReferenceInput("collection.bodySite", resource.Collection.BodySite, htmlAttrs)
}
func (resource *Specimen) T_CollectionFastingStatusCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("collection.fastingStatusCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("collection.fastingStatusCodeableConcept", resource.Collection.FastingStatusCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_CollectionFastingStatusDuration(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DurationInput("collection.fastingStatusDuration", nil, htmlAttrs)
	}
	return DurationInput("collection.fastingStatusDuration", resource.Collection.FastingStatusDuration, htmlAttrs)
}
func (resource *Specimen) T_ProcessingDescription(numProcessing int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return StringInput("processing["+strconv.Itoa(numProcessing)+"].description", nil, htmlAttrs)
	}
	return StringInput("processing["+strconv.Itoa(numProcessing)+"].description", resource.Processing[numProcessing].Description, htmlAttrs)
}
func (resource *Specimen) T_ProcessingMethod(numProcessing int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return CodeableConceptSelect("processing["+strconv.Itoa(numProcessing)+"].method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("processing["+strconv.Itoa(numProcessing)+"].method", resource.Processing[numProcessing].Method, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_ProcessingAdditive(frs []FhirResource, numProcessing int, numAdditive int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) || numAdditive >= len(resource.Processing[numProcessing].Additive) {
		return ReferenceInput(frs, "processing["+strconv.Itoa(numProcessing)+"].additive["+strconv.Itoa(numAdditive)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "processing["+strconv.Itoa(numProcessing)+"].additive["+strconv.Itoa(numAdditive)+"]", &resource.Processing[numProcessing].Additive[numAdditive], htmlAttrs)
}
func (resource *Specimen) T_ProcessingTimeDateTime(numProcessing int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return FhirDateTimeInput("processing["+strconv.Itoa(numProcessing)+"].timeDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("processing["+strconv.Itoa(numProcessing)+"].timeDateTime", resource.Processing[numProcessing].TimeDateTime, htmlAttrs)
}
func (resource *Specimen) T_ProcessingTimePeriod(numProcessing int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return PeriodInput("processing["+strconv.Itoa(numProcessing)+"].timePeriod", nil, htmlAttrs)
	}
	return PeriodInput("processing["+strconv.Itoa(numProcessing)+"].timePeriod", resource.Processing[numProcessing].TimePeriod, htmlAttrs)
}
func (resource *Specimen) T_ContainerDevice(frs []FhirResource, numContainer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContainer >= len(resource.Container) {
		return ReferenceInput(frs, "container["+strconv.Itoa(numContainer)+"].device", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "container["+strconv.Itoa(numContainer)+"].device", &resource.Container[numContainer].Device, htmlAttrs)
}
func (resource *Specimen) T_ContainerLocation(frs []FhirResource, numContainer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContainer >= len(resource.Container) {
		return ReferenceInput(frs, "container["+strconv.Itoa(numContainer)+"].location", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "container["+strconv.Itoa(numContainer)+"].location", resource.Container[numContainer].Location, htmlAttrs)
}
func (resource *Specimen) T_ContainerSpecimenQuantity(numContainer int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numContainer >= len(resource.Container) {
		return QuantityInput("container["+strconv.Itoa(numContainer)+"].specimenQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("container["+strconv.Itoa(numContainer)+"].specimenQuantity", resource.Container[numContainer].SpecimenQuantity, optionsValueSet, htmlAttrs)
}
