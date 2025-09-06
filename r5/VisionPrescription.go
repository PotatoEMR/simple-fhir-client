package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Created           string                                `json:"created"`
	Patient           Reference                             `json:"patient"`
	Encounter         *Reference                            `json:"encounter,omitempty"`
	DateWritten       string                                `json:"dateWritten"`
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

func (resource *VisionPrescription) T_Id() templ.Component {

	if resource == nil {
		return StringInput("VisionPrescription.Id", nil)
	}
	return StringInput("VisionPrescription.Id", resource.Id)
}
func (resource *VisionPrescription) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("VisionPrescription.ImplicitRules", nil)
	}
	return StringInput("VisionPrescription.ImplicitRules", resource.ImplicitRules)
}
func (resource *VisionPrescription) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("VisionPrescription.Language", nil, optionsValueSet)
	}
	return CodeSelect("VisionPrescription.Language", resource.Language, optionsValueSet)
}
func (resource *VisionPrescription) T_Status() templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("VisionPrescription.Status", nil, optionsValueSet)
	}
	return CodeSelect("VisionPrescription.Status", &resource.Status, optionsValueSet)
}
func (resource *VisionPrescription) T_Created() templ.Component {

	if resource == nil {
		return StringInput("VisionPrescription.Created", nil)
	}
	return StringInput("VisionPrescription.Created", &resource.Created)
}
func (resource *VisionPrescription) T_DateWritten() templ.Component {

	if resource == nil {
		return StringInput("VisionPrescription.DateWritten", nil)
	}
	return StringInput("VisionPrescription.DateWritten", &resource.DateWritten)
}
func (resource *VisionPrescription) T_LensSpecificationId(numLensSpecification int) templ.Component {

	if resource == nil || len(resource.LensSpecification) >= numLensSpecification {
		return StringInput("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Id", nil)
	}
	return StringInput("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Id", resource.LensSpecification[numLensSpecification].Id)
}
func (resource *VisionPrescription) T_LensSpecificationProduct(numLensSpecification int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.LensSpecification) >= numLensSpecification {
		return CodeableConceptSelect("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Product", nil, optionsValueSet)
	}
	return CodeableConceptSelect("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Product", &resource.LensSpecification[numLensSpecification].Product, optionsValueSet)
}
func (resource *VisionPrescription) T_LensSpecificationEye(numLensSpecification int) templ.Component {
	optionsValueSet := VSVision_eye_codes

	if resource == nil || len(resource.LensSpecification) >= numLensSpecification {
		return CodeSelect("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Eye", nil, optionsValueSet)
	}
	return CodeSelect("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Eye", &resource.LensSpecification[numLensSpecification].Eye, optionsValueSet)
}
func (resource *VisionPrescription) T_LensSpecificationSphere(numLensSpecification int) templ.Component {

	if resource == nil || len(resource.LensSpecification) >= numLensSpecification {
		return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Sphere", nil)
	}
	return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Sphere", resource.LensSpecification[numLensSpecification].Sphere)
}
func (resource *VisionPrescription) T_LensSpecificationCylinder(numLensSpecification int) templ.Component {

	if resource == nil || len(resource.LensSpecification) >= numLensSpecification {
		return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Cylinder", nil)
	}
	return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Cylinder", resource.LensSpecification[numLensSpecification].Cylinder)
}
func (resource *VisionPrescription) T_LensSpecificationAxis(numLensSpecification int) templ.Component {

	if resource == nil || len(resource.LensSpecification) >= numLensSpecification {
		return IntInput("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Axis", nil)
	}
	return IntInput("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Axis", resource.LensSpecification[numLensSpecification].Axis)
}
func (resource *VisionPrescription) T_LensSpecificationAdd(numLensSpecification int) templ.Component {

	if resource == nil || len(resource.LensSpecification) >= numLensSpecification {
		return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Add", nil)
	}
	return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Add", resource.LensSpecification[numLensSpecification].Add)
}
func (resource *VisionPrescription) T_LensSpecificationPower(numLensSpecification int) templ.Component {

	if resource == nil || len(resource.LensSpecification) >= numLensSpecification {
		return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Power", nil)
	}
	return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Power", resource.LensSpecification[numLensSpecification].Power)
}
func (resource *VisionPrescription) T_LensSpecificationBackCurve(numLensSpecification int) templ.Component {

	if resource == nil || len(resource.LensSpecification) >= numLensSpecification {
		return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].BackCurve", nil)
	}
	return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].BackCurve", resource.LensSpecification[numLensSpecification].BackCurve)
}
func (resource *VisionPrescription) T_LensSpecificationDiameter(numLensSpecification int) templ.Component {

	if resource == nil || len(resource.LensSpecification) >= numLensSpecification {
		return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Diameter", nil)
	}
	return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Diameter", resource.LensSpecification[numLensSpecification].Diameter)
}
func (resource *VisionPrescription) T_LensSpecificationColor(numLensSpecification int) templ.Component {

	if resource == nil || len(resource.LensSpecification) >= numLensSpecification {
		return StringInput("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Color", nil)
	}
	return StringInput("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Color", resource.LensSpecification[numLensSpecification].Color)
}
func (resource *VisionPrescription) T_LensSpecificationBrand(numLensSpecification int) templ.Component {

	if resource == nil || len(resource.LensSpecification) >= numLensSpecification {
		return StringInput("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Brand", nil)
	}
	return StringInput("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Brand", resource.LensSpecification[numLensSpecification].Brand)
}
func (resource *VisionPrescription) T_LensSpecificationPrismId(numLensSpecification int, numPrism int) templ.Component {

	if resource == nil || len(resource.LensSpecification) >= numLensSpecification || len(resource.LensSpecification[numLensSpecification].Prism) >= numPrism {
		return StringInput("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Prism["+strconv.Itoa(numPrism)+"].Id", nil)
	}
	return StringInput("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Prism["+strconv.Itoa(numPrism)+"].Id", resource.LensSpecification[numLensSpecification].Prism[numPrism].Id)
}
func (resource *VisionPrescription) T_LensSpecificationPrismAmount(numLensSpecification int, numPrism int) templ.Component {

	if resource == nil || len(resource.LensSpecification) >= numLensSpecification || len(resource.LensSpecification[numLensSpecification].Prism) >= numPrism {
		return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Prism["+strconv.Itoa(numPrism)+"].Amount", nil)
	}
	return Float64Input("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Prism["+strconv.Itoa(numPrism)+"].Amount", &resource.LensSpecification[numLensSpecification].Prism[numPrism].Amount)
}
func (resource *VisionPrescription) T_LensSpecificationPrismBase(numLensSpecification int, numPrism int) templ.Component {
	optionsValueSet := VSVision_base_codes

	if resource == nil || len(resource.LensSpecification) >= numLensSpecification || len(resource.LensSpecification[numLensSpecification].Prism) >= numPrism {
		return CodeSelect("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Prism["+strconv.Itoa(numPrism)+"].Base", nil, optionsValueSet)
	}
	return CodeSelect("VisionPrescription.LensSpecification["+strconv.Itoa(numLensSpecification)+"].Prism["+strconv.Itoa(numPrism)+"].Base", &resource.LensSpecification[numLensSpecification].Prism[numPrism].Base, optionsValueSet)
}
