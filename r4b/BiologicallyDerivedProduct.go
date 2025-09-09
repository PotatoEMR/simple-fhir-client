package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProduct struct {
	Id                *string                                 `json:"id,omitempty"`
	Meta              *Meta                                   `json:"meta,omitempty"`
	ImplicitRules     *string                                 `json:"implicitRules,omitempty"`
	Language          *string                                 `json:"language,omitempty"`
	Text              *Narrative                              `json:"text,omitempty"`
	Contained         []Resource                              `json:"contained,omitempty"`
	Extension         []Extension                             `json:"extension,omitempty"`
	ModifierExtension []Extension                             `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                            `json:"identifier,omitempty"`
	ProductCategory   *string                                 `json:"productCategory,omitempty"`
	ProductCode       *CodeableConcept                        `json:"productCode,omitempty"`
	Status            *string                                 `json:"status,omitempty"`
	Request           []Reference                             `json:"request,omitempty"`
	Quantity          *int                                    `json:"quantity,omitempty"`
	Parent            []Reference                             `json:"parent,omitempty"`
	Collection        *BiologicallyDerivedProductCollection   `json:"collection,omitempty"`
	Processing        []BiologicallyDerivedProductProcessing  `json:"processing,omitempty"`
	Manipulation      *BiologicallyDerivedProductManipulation `json:"manipulation,omitempty"`
	Storage           []BiologicallyDerivedProductStorage     `json:"storage,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProductCollection struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Collector         *Reference  `json:"collector,omitempty"`
	Source            *Reference  `json:"source,omitempty"`
	CollectedDateTime *string     `json:"collectedDateTime,omitempty"`
	CollectedPeriod   *Period     `json:"collectedPeriod,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProductProcessing struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Procedure         *CodeableConcept `json:"procedure,omitempty"`
	Additive          *Reference       `json:"additive,omitempty"`
	TimeDateTime      *string          `json:"timeDateTime,omitempty"`
	TimePeriod        *Period          `json:"timePeriod,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProductManipulation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       *string     `json:"description,omitempty"`
	TimeDateTime      *string     `json:"timeDateTime,omitempty"`
	TimePeriod        *Period     `json:"timePeriod,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProductStorage struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Temperature       *float64    `json:"temperature,omitempty"`
	Scale             *string     `json:"scale,omitempty"`
	Duration          *Period     `json:"duration,omitempty"`
}

type OtherBiologicallyDerivedProduct BiologicallyDerivedProduct

// on convert struct to json, automatically add resourceType=BiologicallyDerivedProduct
func (r BiologicallyDerivedProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBiologicallyDerivedProduct
		ResourceType string `json:"resourceType"`
	}{
		OtherBiologicallyDerivedProduct: OtherBiologicallyDerivedProduct(r),
		ResourceType:                    "BiologicallyDerivedProduct",
	})
}
func (r BiologicallyDerivedProduct) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "BiologicallyDerivedProduct/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "BiologicallyDerivedProduct"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *BiologicallyDerivedProduct) T_ProductCategory(htmlAttrs string) templ.Component {
	optionsValueSet := VSProduct_category

	if resource == nil {
		return CodeSelect("productCategory", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("productCategory", resource.ProductCategory, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProductCode(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("productCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("productCode", resource.ProductCode, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSProduct_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_Quantity(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("quantity", nil, htmlAttrs)
	}
	return IntInput("quantity", resource.Quantity, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_CollectionCollectedDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("collection.collectedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("collection.collectedDateTime", resource.Collection.CollectedDateTime, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProcessingDescription(numProcessing int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return StringInput("processing["+strconv.Itoa(numProcessing)+"].description", nil, htmlAttrs)
	}
	return StringInput("processing["+strconv.Itoa(numProcessing)+"].description", resource.Processing[numProcessing].Description, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProcessingProcedure(numProcessing int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return CodeableConceptSelect("processing["+strconv.Itoa(numProcessing)+"].procedure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("processing["+strconv.Itoa(numProcessing)+"].procedure", resource.Processing[numProcessing].Procedure, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProcessingTimeDateTime(numProcessing int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return DateTimeInput("processing["+strconv.Itoa(numProcessing)+"].timeDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("processing["+strconv.Itoa(numProcessing)+"].timeDateTime", resource.Processing[numProcessing].TimeDateTime, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ManipulationDescription(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("manipulation.description", nil, htmlAttrs)
	}
	return StringInput("manipulation.description", resource.Manipulation.Description, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ManipulationTimeDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("manipulation.timeDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("manipulation.timeDateTime", resource.Manipulation.TimeDateTime, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_StorageDescription(numStorage int, htmlAttrs string) templ.Component {
	if resource == nil || numStorage >= len(resource.Storage) {
		return StringInput("storage["+strconv.Itoa(numStorage)+"].description", nil, htmlAttrs)
	}
	return StringInput("storage["+strconv.Itoa(numStorage)+"].description", resource.Storage[numStorage].Description, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_StorageTemperature(numStorage int, htmlAttrs string) templ.Component {
	if resource == nil || numStorage >= len(resource.Storage) {
		return Float64Input("storage["+strconv.Itoa(numStorage)+"].temperature", nil, htmlAttrs)
	}
	return Float64Input("storage["+strconv.Itoa(numStorage)+"].temperature", resource.Storage[numStorage].Temperature, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_StorageScale(numStorage int, htmlAttrs string) templ.Component {
	optionsValueSet := VSProduct_storage_scale

	if resource == nil || numStorage >= len(resource.Storage) {
		return CodeSelect("storage["+strconv.Itoa(numStorage)+"].scale", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("storage["+strconv.Itoa(numStorage)+"].scale", resource.Storage[numStorage].Scale, optionsValueSet, htmlAttrs)
}
