package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	CollectedDateTime *time.Time  `json:"collectedDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	TimeDateTime      *time.Time       `json:"timeDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	TimePeriod        *Period          `json:"timePeriod,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProductManipulation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       *string     `json:"description,omitempty"`
	TimeDateTime      *time.Time  `json:"timeDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
		return CodeSelect("ProductCategory", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ProductCategory", resource.ProductCategory, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProductCode(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ProductCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ProductCode", resource.ProductCode, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSProduct_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_Quantity(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("Quantity", nil, htmlAttrs)
	}
	return IntInput("Quantity", resource.Quantity, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_CollectionCollectedDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("CollectionCollectedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("CollectionCollectedDateTime", resource.Collection.CollectedDateTime, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProcessingDescription(numProcessing int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return StringInput("Processing["+strconv.Itoa(numProcessing)+"]Description", nil, htmlAttrs)
	}
	return StringInput("Processing["+strconv.Itoa(numProcessing)+"]Description", resource.Processing[numProcessing].Description, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProcessingProcedure(numProcessing int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return CodeableConceptSelect("Processing["+strconv.Itoa(numProcessing)+"]Procedure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Processing["+strconv.Itoa(numProcessing)+"]Procedure", resource.Processing[numProcessing].Procedure, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProcessingTimeDateTime(numProcessing int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return DateTimeInput("Processing["+strconv.Itoa(numProcessing)+"]TimeDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Processing["+strconv.Itoa(numProcessing)+"]TimeDateTime", resource.Processing[numProcessing].TimeDateTime, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ManipulationDescription(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ManipulationDescription", nil, htmlAttrs)
	}
	return StringInput("ManipulationDescription", resource.Manipulation.Description, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ManipulationTimeDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("ManipulationTimeDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("ManipulationTimeDateTime", resource.Manipulation.TimeDateTime, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_StorageDescription(numStorage int, htmlAttrs string) templ.Component {
	if resource == nil || numStorage >= len(resource.Storage) {
		return StringInput("Storage["+strconv.Itoa(numStorage)+"]Description", nil, htmlAttrs)
	}
	return StringInput("Storage["+strconv.Itoa(numStorage)+"]Description", resource.Storage[numStorage].Description, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_StorageTemperature(numStorage int, htmlAttrs string) templ.Component {
	if resource == nil || numStorage >= len(resource.Storage) {
		return Float64Input("Storage["+strconv.Itoa(numStorage)+"]Temperature", nil, htmlAttrs)
	}
	return Float64Input("Storage["+strconv.Itoa(numStorage)+"]Temperature", resource.Storage[numStorage].Temperature, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_StorageScale(numStorage int, htmlAttrs string) templ.Component {
	optionsValueSet := VSProduct_storage_scale

	if resource == nil || numStorage >= len(resource.Storage) {
		return CodeSelect("Storage["+strconv.Itoa(numStorage)+"]Scale", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Storage["+strconv.Itoa(numStorage)+"]Scale", resource.Storage[numStorage].Scale, optionsValueSet, htmlAttrs)
}
