package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/BiologicallyDerivedProduct
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

// http://hl7.org/fhir/r4/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProductCollection struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Collector         *Reference  `json:"collector,omitempty"`
	Source            *Reference  `json:"source,omitempty"`
	CollectedDateTime *string     `json:"collectedDateTime,omitempty"`
	CollectedPeriod   *Period     `json:"collectedPeriod,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/BiologicallyDerivedProduct
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

// http://hl7.org/fhir/r4/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProductManipulation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       *string     `json:"description,omitempty"`
	TimeDateTime      *string     `json:"timeDateTime,omitempty"`
	TimePeriod        *Period     `json:"timePeriod,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/BiologicallyDerivedProduct
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

func (resource *BiologicallyDerivedProduct) T_Id() templ.Component {

	if resource == nil {
		return StringInput("BiologicallyDerivedProduct.Id", nil)
	}
	return StringInput("BiologicallyDerivedProduct.Id", resource.Id)
}
func (resource *BiologicallyDerivedProduct) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("BiologicallyDerivedProduct.ImplicitRules", nil)
	}
	return StringInput("BiologicallyDerivedProduct.ImplicitRules", resource.ImplicitRules)
}
func (resource *BiologicallyDerivedProduct) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("BiologicallyDerivedProduct.Language", nil, optionsValueSet)
	}
	return CodeSelect("BiologicallyDerivedProduct.Language", resource.Language, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) T_ProductCategory() templ.Component {
	optionsValueSet := VSProduct_category

	if resource == nil {
		return CodeSelect("BiologicallyDerivedProduct.ProductCategory", nil, optionsValueSet)
	}
	return CodeSelect("BiologicallyDerivedProduct.ProductCategory", resource.ProductCategory, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) T_ProductCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("BiologicallyDerivedProduct.ProductCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("BiologicallyDerivedProduct.ProductCode", resource.ProductCode, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) T_Status() templ.Component {
	optionsValueSet := VSProduct_status

	if resource == nil {
		return CodeSelect("BiologicallyDerivedProduct.Status", nil, optionsValueSet)
	}
	return CodeSelect("BiologicallyDerivedProduct.Status", resource.Status, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) T_Quantity() templ.Component {

	if resource == nil {
		return IntInput("BiologicallyDerivedProduct.Quantity", nil)
	}
	return IntInput("BiologicallyDerivedProduct.Quantity", resource.Quantity)
}
func (resource *BiologicallyDerivedProduct) T_CollectionId() templ.Component {

	if resource == nil {
		return StringInput("BiologicallyDerivedProduct.Collection.Id", nil)
	}
	return StringInput("BiologicallyDerivedProduct.Collection.Id", resource.Collection.Id)
}
func (resource *BiologicallyDerivedProduct) T_ProcessingId(numProcessing int) templ.Component {

	if resource == nil || len(resource.Processing) >= numProcessing {
		return StringInput("BiologicallyDerivedProduct.Processing["+strconv.Itoa(numProcessing)+"].Id", nil)
	}
	return StringInput("BiologicallyDerivedProduct.Processing["+strconv.Itoa(numProcessing)+"].Id", resource.Processing[numProcessing].Id)
}
func (resource *BiologicallyDerivedProduct) T_ProcessingDescription(numProcessing int) templ.Component {

	if resource == nil || len(resource.Processing) >= numProcessing {
		return StringInput("BiologicallyDerivedProduct.Processing["+strconv.Itoa(numProcessing)+"].Description", nil)
	}
	return StringInput("BiologicallyDerivedProduct.Processing["+strconv.Itoa(numProcessing)+"].Description", resource.Processing[numProcessing].Description)
}
func (resource *BiologicallyDerivedProduct) T_ProcessingProcedure(numProcessing int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Processing) >= numProcessing {
		return CodeableConceptSelect("BiologicallyDerivedProduct.Processing["+strconv.Itoa(numProcessing)+"].Procedure", nil, optionsValueSet)
	}
	return CodeableConceptSelect("BiologicallyDerivedProduct.Processing["+strconv.Itoa(numProcessing)+"].Procedure", resource.Processing[numProcessing].Procedure, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) T_ManipulationId() templ.Component {

	if resource == nil {
		return StringInput("BiologicallyDerivedProduct.Manipulation.Id", nil)
	}
	return StringInput("BiologicallyDerivedProduct.Manipulation.Id", resource.Manipulation.Id)
}
func (resource *BiologicallyDerivedProduct) T_ManipulationDescription() templ.Component {

	if resource == nil {
		return StringInput("BiologicallyDerivedProduct.Manipulation.Description", nil)
	}
	return StringInput("BiologicallyDerivedProduct.Manipulation.Description", resource.Manipulation.Description)
}
func (resource *BiologicallyDerivedProduct) T_StorageId(numStorage int) templ.Component {

	if resource == nil || len(resource.Storage) >= numStorage {
		return StringInput("BiologicallyDerivedProduct.Storage["+strconv.Itoa(numStorage)+"].Id", nil)
	}
	return StringInput("BiologicallyDerivedProduct.Storage["+strconv.Itoa(numStorage)+"].Id", resource.Storage[numStorage].Id)
}
func (resource *BiologicallyDerivedProduct) T_StorageDescription(numStorage int) templ.Component {

	if resource == nil || len(resource.Storage) >= numStorage {
		return StringInput("BiologicallyDerivedProduct.Storage["+strconv.Itoa(numStorage)+"].Description", nil)
	}
	return StringInput("BiologicallyDerivedProduct.Storage["+strconv.Itoa(numStorage)+"].Description", resource.Storage[numStorage].Description)
}
func (resource *BiologicallyDerivedProduct) T_StorageTemperature(numStorage int) templ.Component {

	if resource == nil || len(resource.Storage) >= numStorage {
		return Float64Input("BiologicallyDerivedProduct.Storage["+strconv.Itoa(numStorage)+"].Temperature", nil)
	}
	return Float64Input("BiologicallyDerivedProduct.Storage["+strconv.Itoa(numStorage)+"].Temperature", resource.Storage[numStorage].Temperature)
}
func (resource *BiologicallyDerivedProduct) T_StorageScale(numStorage int) templ.Component {
	optionsValueSet := VSProduct_storage_scale

	if resource == nil || len(resource.Storage) >= numStorage {
		return CodeSelect("BiologicallyDerivedProduct.Storage["+strconv.Itoa(numStorage)+"].Scale", nil, optionsValueSet)
	}
	return CodeSelect("BiologicallyDerivedProduct.Storage["+strconv.Itoa(numStorage)+"].Scale", resource.Storage[numStorage].Scale, optionsValueSet)
}
