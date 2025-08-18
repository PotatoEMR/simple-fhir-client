//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

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
	CollectedDateTime            *string            `json:"collectedDateTime"`
	CollectedPeriod              *Period            `json:"collectedPeriod"`
	Duration                     *Duration          `json:"duration,omitempty"`
	Quantity                     *Quantity          `json:"quantity,omitempty"`
	Method                       *CodeableConcept   `json:"method,omitempty"`
	Device                       *CodeableReference `json:"device,omitempty"`
	Procedure                    *Reference         `json:"procedure,omitempty"`
	BodySite                     *CodeableReference `json:"bodySite,omitempty"`
	FastingStatusCodeableConcept *CodeableConcept   `json:"fastingStatusCodeableConcept"`
	FastingStatusDuration        *Duration          `json:"fastingStatusDuration"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Specimen
type SpecimenProcessing struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Method            *CodeableConcept `json:"method,omitempty"`
	Additive          []Reference      `json:"additive,omitempty"`
	TimeDateTime      *string          `json:"timeDateTime"`
	TimePeriod        *Period          `json:"timePeriod"`
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
