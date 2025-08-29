package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	CollectedDateTime            *string          `json:"collectedDateTime"`
	CollectedPeriod              *Period          `json:"collectedPeriod"`
	Duration                     *Duration        `json:"duration,omitempty"`
	Quantity                     *Quantity        `json:"quantity,omitempty"`
	Method                       *CodeableConcept `json:"method,omitempty"`
	BodySite                     *CodeableConcept `json:"bodySite,omitempty"`
	FastingStatusCodeableConcept *CodeableConcept `json:"fastingStatusCodeableConcept"`
	FastingStatusDuration        *Duration        `json:"fastingStatusDuration"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Specimen
type SpecimenProcessing struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Procedure         *CodeableConcept `json:"procedure,omitempty"`
	Additive          []Reference      `json:"additive,omitempty"`
	TimeDateTime      *string          `json:"timeDateTime"`
	TimePeriod        *Period          `json:"timePeriod"`
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
	AdditiveCodeableConcept *CodeableConcept `json:"additiveCodeableConcept"`
	AdditiveReference       *Reference       `json:"additiveReference"`
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

func (resource *Specimen) SpecimenLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Specimen) SpecimenStatus() templ.Component {
	optionsValueSet := VSSpecimen_status
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
