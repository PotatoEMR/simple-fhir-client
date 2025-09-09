package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *VisionPrescription) T_Created(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("created", nil, htmlAttrs)
	}
	return DateTimeInput("created", &resource.Created, htmlAttrs)
}
func (resource *VisionPrescription) T_DateWritten(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("dateWritten", nil, htmlAttrs)
	}
	return DateTimeInput("dateWritten", &resource.DateWritten, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationProduct(numLensSpecification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return CodeableConceptSelect("lensSpecification["+strconv.Itoa(numLensSpecification)+"].product", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("lensSpecification["+strconv.Itoa(numLensSpecification)+"].product", &resource.LensSpecification[numLensSpecification].Product, optionsValueSet, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationEye(numLensSpecification int, htmlAttrs string) templ.Component {
	optionsValueSet := VSVision_eye_codes

	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return CodeSelect("lensSpecification["+strconv.Itoa(numLensSpecification)+"].eye", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("lensSpecification["+strconv.Itoa(numLensSpecification)+"].eye", &resource.LensSpecification[numLensSpecification].Eye, optionsValueSet, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationSphere(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return Float64Input("lensSpecification["+strconv.Itoa(numLensSpecification)+"].sphere", nil, htmlAttrs)
	}
	return Float64Input("lensSpecification["+strconv.Itoa(numLensSpecification)+"].sphere", resource.LensSpecification[numLensSpecification].Sphere, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationCylinder(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return Float64Input("lensSpecification["+strconv.Itoa(numLensSpecification)+"].cylinder", nil, htmlAttrs)
	}
	return Float64Input("lensSpecification["+strconv.Itoa(numLensSpecification)+"].cylinder", resource.LensSpecification[numLensSpecification].Cylinder, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationAxis(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return IntInput("lensSpecification["+strconv.Itoa(numLensSpecification)+"].axis", nil, htmlAttrs)
	}
	return IntInput("lensSpecification["+strconv.Itoa(numLensSpecification)+"].axis", resource.LensSpecification[numLensSpecification].Axis, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationAdd(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return Float64Input("lensSpecification["+strconv.Itoa(numLensSpecification)+"].add", nil, htmlAttrs)
	}
	return Float64Input("lensSpecification["+strconv.Itoa(numLensSpecification)+"].add", resource.LensSpecification[numLensSpecification].Add, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationPower(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return Float64Input("lensSpecification["+strconv.Itoa(numLensSpecification)+"].power", nil, htmlAttrs)
	}
	return Float64Input("lensSpecification["+strconv.Itoa(numLensSpecification)+"].power", resource.LensSpecification[numLensSpecification].Power, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationBackCurve(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return Float64Input("lensSpecification["+strconv.Itoa(numLensSpecification)+"].backCurve", nil, htmlAttrs)
	}
	return Float64Input("lensSpecification["+strconv.Itoa(numLensSpecification)+"].backCurve", resource.LensSpecification[numLensSpecification].BackCurve, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationDiameter(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return Float64Input("lensSpecification["+strconv.Itoa(numLensSpecification)+"].diameter", nil, htmlAttrs)
	}
	return Float64Input("lensSpecification["+strconv.Itoa(numLensSpecification)+"].diameter", resource.LensSpecification[numLensSpecification].Diameter, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationColor(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return StringInput("lensSpecification["+strconv.Itoa(numLensSpecification)+"].color", nil, htmlAttrs)
	}
	return StringInput("lensSpecification["+strconv.Itoa(numLensSpecification)+"].color", resource.LensSpecification[numLensSpecification].Color, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationBrand(numLensSpecification int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) {
		return StringInput("lensSpecification["+strconv.Itoa(numLensSpecification)+"].brand", nil, htmlAttrs)
	}
	return StringInput("lensSpecification["+strconv.Itoa(numLensSpecification)+"].brand", resource.LensSpecification[numLensSpecification].Brand, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationNote(numLensSpecification int, numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) || numNote >= len(resource.LensSpecification[numLensSpecification].Note) {
		return AnnotationTextArea("lensSpecification["+strconv.Itoa(numLensSpecification)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("lensSpecification["+strconv.Itoa(numLensSpecification)+"].note["+strconv.Itoa(numNote)+"]", &resource.LensSpecification[numLensSpecification].Note[numNote], htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationPrismAmount(numLensSpecification int, numPrism int, htmlAttrs string) templ.Component {
	if resource == nil || numLensSpecification >= len(resource.LensSpecification) || numPrism >= len(resource.LensSpecification[numLensSpecification].Prism) {
		return Float64Input("lensSpecification["+strconv.Itoa(numLensSpecification)+"].prism["+strconv.Itoa(numPrism)+"].amount", nil, htmlAttrs)
	}
	return Float64Input("lensSpecification["+strconv.Itoa(numLensSpecification)+"].prism["+strconv.Itoa(numPrism)+"].amount", &resource.LensSpecification[numLensSpecification].Prism[numPrism].Amount, htmlAttrs)
}
func (resource *VisionPrescription) T_LensSpecificationPrismBase(numLensSpecification int, numPrism int, htmlAttrs string) templ.Component {
	optionsValueSet := VSVision_base_codes

	if resource == nil || numLensSpecification >= len(resource.LensSpecification) || numPrism >= len(resource.LensSpecification[numLensSpecification].Prism) {
		return CodeSelect("lensSpecification["+strconv.Itoa(numLensSpecification)+"].prism["+strconv.Itoa(numPrism)+"].base", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("lensSpecification["+strconv.Itoa(numLensSpecification)+"].prism["+strconv.Itoa(numPrism)+"].base", &resource.LensSpecification[numLensSpecification].Prism[numPrism].Base, optionsValueSet, htmlAttrs)
}
