package r4

//generated with command go run ./bultaoreune
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
	Id                *string       `json:"id,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	Collector         *Reference    `json:"collector,omitempty"`
	Source            *Reference    `json:"source,omitempty"`
	CollectedDateTime *FhirDateTime `json:"collectedDateTime,omitempty"`
	CollectedPeriod   *Period       `json:"collectedPeriod,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProductProcessing struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Procedure         *CodeableConcept `json:"procedure,omitempty"`
	Additive          *Reference       `json:"additive,omitempty"`
	TimeDateTime      *FhirDateTime    `json:"timeDateTime,omitempty"`
	TimePeriod        *Period          `json:"timePeriod,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProductManipulation struct {
	Id                *string       `json:"id,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	Description       *string       `json:"description,omitempty"`
	TimeDateTime      *FhirDateTime `json:"timeDateTime,omitempty"`
	TimePeriod        *Period       `json:"timePeriod,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
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
func (r BiologicallyDerivedProduct) ResourceType() string {
	return "BiologicallyDerivedProduct"
}

func (resource *BiologicallyDerivedProduct) T_ProductCategory(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSProduct_category

	if resource == nil {
		return CodeSelect("productCategory", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("productCategory", resource.ProductCategory, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProductCode(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("productCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("productCode", resource.ProductCode, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSProduct_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_Request(frs []FhirResource, numRequest int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRequest >= len(resource.Request) {
		return ReferenceInput(frs, "request["+strconv.Itoa(numRequest)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "request["+strconv.Itoa(numRequest)+"]", &resource.Request[numRequest], htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_Quantity(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("quantity", nil, htmlAttrs)
	}
	return IntInput("quantity", resource.Quantity, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_Parent(frs []FhirResource, numParent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParent >= len(resource.Parent) {
		return ReferenceInput(frs, "parent["+strconv.Itoa(numParent)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "parent["+strconv.Itoa(numParent)+"]", &resource.Parent[numParent], htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_CollectionCollector(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "collection.collector", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "collection.collector", resource.Collection.Collector, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_CollectionSource(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "collection.source", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "collection.source", resource.Collection.Source, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_CollectionCollectedDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("collection.collectedDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("collection.collectedDateTime", resource.Collection.CollectedDateTime, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_CollectionCollectedPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("collection.collectedPeriod", nil, htmlAttrs)
	}
	return PeriodInput("collection.collectedPeriod", resource.Collection.CollectedPeriod, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProcessingDescription(numProcessing int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return StringInput("processing["+strconv.Itoa(numProcessing)+"].description", nil, htmlAttrs)
	}
	return StringInput("processing["+strconv.Itoa(numProcessing)+"].description", resource.Processing[numProcessing].Description, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProcessingProcedure(numProcessing int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return CodeableConceptSelect("processing["+strconv.Itoa(numProcessing)+"].procedure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("processing["+strconv.Itoa(numProcessing)+"].procedure", resource.Processing[numProcessing].Procedure, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProcessingAdditive(frs []FhirResource, numProcessing int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return ReferenceInput(frs, "processing["+strconv.Itoa(numProcessing)+"].additive", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "processing["+strconv.Itoa(numProcessing)+"].additive", resource.Processing[numProcessing].Additive, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProcessingTimeDateTime(numProcessing int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return FhirDateTimeInput("processing["+strconv.Itoa(numProcessing)+"].timeDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("processing["+strconv.Itoa(numProcessing)+"].timeDateTime", resource.Processing[numProcessing].TimeDateTime, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProcessingTimePeriod(numProcessing int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcessing >= len(resource.Processing) {
		return PeriodInput("processing["+strconv.Itoa(numProcessing)+"].timePeriod", nil, htmlAttrs)
	}
	return PeriodInput("processing["+strconv.Itoa(numProcessing)+"].timePeriod", resource.Processing[numProcessing].TimePeriod, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ManipulationDescription(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("manipulation.description", nil, htmlAttrs)
	}
	return StringInput("manipulation.description", resource.Manipulation.Description, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ManipulationTimeDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("manipulation.timeDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("manipulation.timeDateTime", resource.Manipulation.TimeDateTime, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ManipulationTimePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("manipulation.timePeriod", nil, htmlAttrs)
	}
	return PeriodInput("manipulation.timePeriod", resource.Manipulation.TimePeriod, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_StorageDescription(numStorage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStorage >= len(resource.Storage) {
		return StringInput("storage["+strconv.Itoa(numStorage)+"].description", nil, htmlAttrs)
	}
	return StringInput("storage["+strconv.Itoa(numStorage)+"].description", resource.Storage[numStorage].Description, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_StorageTemperature(numStorage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStorage >= len(resource.Storage) {
		return Float64Input("storage["+strconv.Itoa(numStorage)+"].temperature", nil, htmlAttrs)
	}
	return Float64Input("storage["+strconv.Itoa(numStorage)+"].temperature", resource.Storage[numStorage].Temperature, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_StorageScale(numStorage int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSProduct_storage_scale

	if resource == nil || numStorage >= len(resource.Storage) {
		return CodeSelect("storage["+strconv.Itoa(numStorage)+"].scale", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("storage["+strconv.Itoa(numStorage)+"].scale", resource.Storage[numStorage].Scale, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_StorageDuration(numStorage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStorage >= len(resource.Storage) {
		return PeriodInput("storage["+strconv.Itoa(numStorage)+"].duration", nil, htmlAttrs)
	}
	return PeriodInput("storage["+strconv.Itoa(numStorage)+"].duration", resource.Storage[numStorage].Duration, htmlAttrs)
}
