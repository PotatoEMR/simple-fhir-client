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

// http://hl7.org/fhir/r5/StructureDefinition/VisionPrescription
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
	Created           time.Time                             `json:"created,format:'2006-01-02T15:04:05Z07:00'"`
	Patient           Reference                             `json:"patient"`
	Encounter         *Reference                            `json:"encounter,omitempty"`
	DateWritten       time.Time                             `json:"dateWritten,format:'2006-01-02T15:04:05Z07:00'"`
	Prescriber        Reference                             `json:"prescriber"`
	LensSpecification []VisionPrescriptionLensSpecification `json:"lensSpecification"`
}

// http://hl7.org/fhir/r5/StructureDefinition/VisionPrescription
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

// http://hl7.org/fhir/r5/StructureDefinition/VisionPrescription
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
func (r VisionPrescription) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "VisionPrescription/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "VisionPrescription"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *VisionPrescription) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("VisionPrescription.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("VisionPrescription.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *VisionPrescription) T_Created(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("VisionPrescription.Created", nil, htmlAttrs)
	}
	return DateTimeInput("VisionPrescription.Created", &resource.Created, htmlAttrs)
}
func (resource *VisionPrescription) T_DateWritten(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("VisionPrescription.DateWritten", nil, htmlAttrs)
	}
	return DateTimeInput("VisionPrescription.DateWritten", &resource.DateWritten, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationProduct(numLensSpecification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return CodeableConceptSelect("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Product", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Product", &resource.LensSpecification[numLensSpecification].Product, optionsValueSet, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationEye(numLensSpecification int, htmlAttrs string) templ.Component {
	optionsValueSet := VSVision_eye_codes

	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return CodeSelect("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Eye", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Eye", &resource.LensSpecification[numLensSpecification].Eye, optionsValueSet, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationSphere(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Sphere", nil, htmlAttrs)
	}
	return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Sphere", resource.LensSpecification[numLensSpecification].Sphere, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationCylinder(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Cylinder", nil, htmlAttrs)
	}
	return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Cylinder", resource.LensSpecification[numLensSpecification].Cylinder, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationAxis(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return IntInput("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Axis", nil, htmlAttrs)
	}
	return IntInput("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Axis", resource.LensSpecification[numLensSpecification].Axis, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationAdd(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Add", nil, htmlAttrs)
	}
	return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Add", resource.LensSpecification[numLensSpecification].Add, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationPower(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Power", nil, htmlAttrs)
	}
	return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Power", resource.LensSpecification[numLensSpecification].Power, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationBackCurve(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].BackCurve", nil, htmlAttrs)
	}
	return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].BackCurve", resource.LensSpecification[numLensSpecification].BackCurve, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationDiameter(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Diameter", nil, htmlAttrs)
	}
	return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Diameter", resource.LensSpecification[numLensSpecification].Diameter, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationColor(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return StringInput("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Color", nil, htmlAttrs)
	}
	return StringInput("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Color", resource.LensSpecification[numLensSpecification].Color, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationBrand(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return StringInput("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Brand", nil, htmlAttrs)
	}
	return StringInput("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Brand", resource.LensSpecification[numLensSpecification].Brand, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationNote(numLensSpecification int, numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) || numNote >= len(resource.LensSpecification[numLensSpecification].Note) {
		return AnnotationTextArea("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Note["+strconv.Itoa(numNote)+"]", &resource.LensSpecification[numLensSpecification].Note[numNote], htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationPrismAmount(numLensSpecification int, numPrism int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) || numPrism >= len(resource.LensSpecification[numLensSpecification].Prism) {
		return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Prism["+strconv.Itoa(numPrism)+"].Amount", nil, htmlAttrs)
	}
	return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Prism["+strconv.Itoa(numPrism)+"].Amount", &resource.LensSpecification[numLensSpecification].Prism[numPrism].Amount, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationPrismBase(numLensSpecification int, numPrism int, htmlAttrs string) templ.Component {
	optionsValueSet := VSVision_base_codes

	if resource == nil || numLensSpecification >= len(resource.LensSpecification) || numPrism >= len(resource.LensSpecification[numLensSpecification].Prism) {
		return CodeSelect("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Prism["+strconv.Itoa(numPrism)+"].Base", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Prism["+strconv.Itoa(numPrism)+"].Base", &resource.LensSpecification[numLensSpecification].Prism[numPrism].Base, optionsValueSet, htmlAttrs)
}
