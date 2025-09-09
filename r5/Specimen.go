package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	ReceivedTime        *time.Time           `json:"receivedTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	CollectedDateTime            *time.Time         `json:"collectedDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	TimeDateTime      *time.Time       `json:"timeDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *Specimen) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSSpecimen_status

	if resource == nil {
		return CodeSelect("Specimen.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Specimen.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Specimen.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Specimen.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_ReceivedTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Specimen.ReceivedTime", nil, htmlAttrs)
	}
	return DateTimeInput("Specimen.ReceivedTime", resource.ReceivedTime, htmlAttrs)
}
func (resource *Specimen) T_Combined(htmlAttrs string) templ.Component {
	optionsValueSet := VSSpecimen_combined

	if resource == nil {
		return CodeSelect("Specimen.Combined", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Specimen.Combined", resource.Combined, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_Role(numRole int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRole >= len(resource.Role) {
		return CodeableConceptSelect("Specimen.Role["+strconv.Itoa(numRole)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Specimen.Role["+strconv.Itoa(numRole)+"]", &resource.Role[numRole], optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_Condition(numCondition int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return CodeableConceptSelect("Specimen.Condition["+strconv.Itoa(numCondition)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Specimen.Condition["+strconv.Itoa(numCondition)+"]", &resource.Condition[numCondition], optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Specimen.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Specimen.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Specimen) T_FeatureType(numFeature int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numFeature >= len(resource.Feature) {
		return CodeableConceptSelect("Specimen.Feature["+strconv.Itoa(numFeature)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Specimen.Feature["+strconv.Itoa(numFeature)+"].Type", &resource.Feature[numFeature].Type, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_FeatureDescription(numFeature int, htmlAttrs string) templ.Component {
	if resource == nil || numFeature >= len(resource.Feature) {
		return StringInput("Specimen.Feature["+strconv.Itoa(numFeature)+"].Description", nil, htmlAttrs)
	}
	return StringInput("Specimen.Feature["+strconv.Itoa(numFeature)+"].Description", &resource.Feature[numFeature].Description, htmlAttrs)
}
func (resource *Specimen) T_CollectionCollectedDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Specimen.Collection.CollectedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Specimen.Collection.CollectedDateTime", resource.Collection.CollectedDateTime, htmlAttrs)
}
func (resource *Specimen) T_CollectionMethod(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Specimen.Collection.Method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Specimen.Collection.Method", resource.Collection.Method, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_CollectionFastingStatusCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Specimen.Collection.FastingStatusCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Specimen.Collection.FastingStatusCodeableConcept", resource.Collection.FastingStatusCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_ProcessingDescription(numProcessing int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return StringInput("Specimen.Processing["+strconv.Itoa(numProcessing)+"].Description", nil, htmlAttrs)
	}
	return StringInput("Specimen.Processing["+strconv.Itoa(numProcessing)+"].Description", resource.Processing[numProcessing].Description, htmlAttrs)
}
func (resource *Specimen) T_ProcessingMethod(numProcessing int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return CodeableConceptSelect("Specimen.Processing["+strconv.Itoa(numProcessing)+"].Method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Specimen.Processing["+strconv.Itoa(numProcessing)+"].Method", resource.Processing[numProcessing].Method, optionsValueSet, htmlAttrs)
}
func (resource *Specimen) T_ProcessingTimeDateTime(numProcessing int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return DateTimeInput("Specimen.Processing["+strconv.Itoa(numProcessing)+"].TimeDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Specimen.Processing["+strconv.Itoa(numProcessing)+"].TimeDateTime", resource.Processing[numProcessing].TimeDateTime, htmlAttrs)
}
