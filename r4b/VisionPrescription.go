//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

// http://hl7.org/fhir/r4b/StructureDefinition/VisionPrescription
type VisionPrescription struct {
	Id                *string                               `json:"id,omitempty"`
	Meta              *Meta                                 `json:"meta,omitempty"`
	ImplicitRules     *string                               `json:"implicitRules,omitempty"`
	Language          *string                               `json:"language,omitempty"`
	Text              *Narrative                            `json:"text,omitempty"`
	Contained         []Resource                            `json:"contained,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                          `json:"identifier,omitempty"`
	Status            string                                `json:"status"`
	Created           string                                `json:"created"`
	Patient           Reference                             `json:"patient"`
	Encounter         *Reference                            `json:"encounter,omitempty"`
	DateWritten       string                                `json:"dateWritten"`
	Prescriber        Reference                             `json:"prescriber"`
	LensSpecification []VisionPrescriptionLensSpecification `json:"lensSpecification"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/VisionPrescription
type VisionPrescriptionLensSpecification struct {
	Id                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Product           CodeableConcept                            `json:"product"`
	Eye               string                                     `json:"eye"`
	Sphere            *float64                                   `json:"sphere,omitempty"`
	Cylinder          *float64                                   `json:"cylinder,omitempty"`
	Axis              *int                                       `json:"axis,omitempty"`
	Prism             []VisionPrescriptionLensSpecificationPrism `json:"prism,omitempty"`
	Add               *float64                                   `json:"add,omitempty"`
	Power             *float64                                   `json:"power,omitempty"`
	BackCurve         *float64                                   `json:"backCurve,omitempty"`
	Diameter          *float64                                   `json:"diameter,omitempty"`
	Duration          *Quantity                                  `json:"duration,omitempty"`
	Color             *string                                    `json:"color,omitempty"`
	Brand             *string                                    `json:"brand,omitempty"`
	Note              []Annotation                               `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/VisionPrescription
type VisionPrescriptionLensSpecificationPrism struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Amount            float64     `json:"amount"`
	Base              string      `json:"base"`
}

type OtherVisionPrescription VisionPrescription

// on convert struct to json, automatically add resourceType=VisionPrescription
func (r VisionPrescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherVisionPrescription
		ResourceType string `json:"resourceType"`
	}{
		OtherVisionPrescription: OtherVisionPrescription(r),
		ResourceType:            "VisionPrescription",
	})
}
