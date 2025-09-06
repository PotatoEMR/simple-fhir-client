package r5

//generated with command go run ./bultaoreune -nodownload
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
	ReceivedTime        *string              `json:"receivedTime,omitempty"`
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
	CollectedDateTime            *string            `json:"collectedDateTime,omitempty"`
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
	TimeDateTime      *string          `json:"timeDateTime,omitempty"`
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

func (resource *Specimen) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Specimen.Id", nil)
	}
	return StringInput("Specimen.Id", resource.Id)
}
func (resource *Specimen) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Specimen.ImplicitRules", nil)
	}
	return StringInput("Specimen.ImplicitRules", resource.ImplicitRules)
}
func (resource *Specimen) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Specimen.Language", nil, optionsValueSet)
	}
	return CodeSelect("Specimen.Language", resource.Language, optionsValueSet)
}
func (resource *Specimen) T_Status() templ.Component {
	optionsValueSet := VSSpecimen_status

	if resource == nil {
		return CodeSelect("Specimen.Status", nil, optionsValueSet)
	}
	return CodeSelect("Specimen.Status", resource.Status, optionsValueSet)
}
func (resource *Specimen) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Specimen.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Specimen.Type", resource.Type, optionsValueSet)
}
func (resource *Specimen) T_ReceivedTime() templ.Component {

	if resource == nil {
		return StringInput("Specimen.ReceivedTime", nil)
	}
	return StringInput("Specimen.ReceivedTime", resource.ReceivedTime)
}
func (resource *Specimen) T_Combined() templ.Component {
	optionsValueSet := VSSpecimen_combined

	if resource == nil {
		return CodeSelect("Specimen.Combined", nil, optionsValueSet)
	}
	return CodeSelect("Specimen.Combined", resource.Combined, optionsValueSet)
}
func (resource *Specimen) T_Role(numRole int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Role) >= numRole {
		return CodeableConceptSelect("Specimen.Role["+strconv.Itoa(numRole)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Specimen.Role["+strconv.Itoa(numRole)+"]", &resource.Role[numRole], optionsValueSet)
}
func (resource *Specimen) T_Condition(numCondition int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Condition) >= numCondition {
		return CodeableConceptSelect("Specimen.Condition["+strconv.Itoa(numCondition)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Specimen.Condition["+strconv.Itoa(numCondition)+"]", &resource.Condition[numCondition], optionsValueSet)
}
func (resource *Specimen) T_FeatureId(numFeature int) templ.Component {

	if resource == nil || len(resource.Feature) >= numFeature {
		return StringInput("Specimen.Feature["+strconv.Itoa(numFeature)+"].Id", nil)
	}
	return StringInput("Specimen.Feature["+strconv.Itoa(numFeature)+"].Id", resource.Feature[numFeature].Id)
}
func (resource *Specimen) T_FeatureType(numFeature int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Feature) >= numFeature {
		return CodeableConceptSelect("Specimen.Feature["+strconv.Itoa(numFeature)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Specimen.Feature["+strconv.Itoa(numFeature)+"].Type", &resource.Feature[numFeature].Type, optionsValueSet)
}
func (resource *Specimen) T_FeatureDescription(numFeature int) templ.Component {

	if resource == nil || len(resource.Feature) >= numFeature {
		return StringInput("Specimen.Feature["+strconv.Itoa(numFeature)+"].Description", nil)
	}
	return StringInput("Specimen.Feature["+strconv.Itoa(numFeature)+"].Description", &resource.Feature[numFeature].Description)
}
func (resource *Specimen) T_CollectionId() templ.Component {

	if resource == nil {
		return StringInput("Specimen.Collection.Id", nil)
	}
	return StringInput("Specimen.Collection.Id", resource.Collection.Id)
}
func (resource *Specimen) T_CollectionMethod(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Specimen.Collection.Method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Specimen.Collection.Method", resource.Collection.Method, optionsValueSet)
}
func (resource *Specimen) T_ProcessingId(numProcessing int) templ.Component {

	if resource == nil || len(resource.Processing) >= numProcessing {
		return StringInput("Specimen.Processing["+strconv.Itoa(numProcessing)+"].Id", nil)
	}
	return StringInput("Specimen.Processing["+strconv.Itoa(numProcessing)+"].Id", resource.Processing[numProcessing].Id)
}
func (resource *Specimen) T_ProcessingDescription(numProcessing int) templ.Component {

	if resource == nil || len(resource.Processing) >= numProcessing {
		return StringInput("Specimen.Processing["+strconv.Itoa(numProcessing)+"].Description", nil)
	}
	return StringInput("Specimen.Processing["+strconv.Itoa(numProcessing)+"].Description", resource.Processing[numProcessing].Description)
}
func (resource *Specimen) T_ProcessingMethod(numProcessing int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Processing) >= numProcessing {
		return CodeableConceptSelect("Specimen.Processing["+strconv.Itoa(numProcessing)+"].Method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Specimen.Processing["+strconv.Itoa(numProcessing)+"].Method", resource.Processing[numProcessing].Method, optionsValueSet)
}
func (resource *Specimen) T_ContainerId(numContainer int) templ.Component {

	if resource == nil || len(resource.Container) >= numContainer {
		return StringInput("Specimen.Container["+strconv.Itoa(numContainer)+"].Id", nil)
	}
	return StringInput("Specimen.Container["+strconv.Itoa(numContainer)+"].Id", resource.Container[numContainer].Id)
}
