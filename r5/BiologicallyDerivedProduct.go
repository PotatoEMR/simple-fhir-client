package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProduct struct {
	Id                      *string                               `json:"id,omitempty"`
	Meta                    *Meta                                 `json:"meta,omitempty"`
	ImplicitRules           *string                               `json:"implicitRules,omitempty"`
	Language                *string                               `json:"language,omitempty"`
	Text                    *Narrative                            `json:"text,omitempty"`
	Contained               []Resource                            `json:"contained,omitempty"`
	Extension               []Extension                           `json:"extension,omitempty"`
	ModifierExtension       []Extension                           `json:"modifierExtension,omitempty"`
	ProductCategory         *Coding                               `json:"productCategory,omitempty"`
	ProductCode             *CodeableConcept                      `json:"productCode,omitempty"`
	Parent                  []Reference                           `json:"parent,omitempty"`
	Request                 []Reference                           `json:"request,omitempty"`
	Identifier              []Identifier                          `json:"identifier,omitempty"`
	BiologicalSourceEvent   *Identifier                           `json:"biologicalSourceEvent,omitempty"`
	ProcessingFacility      []Reference                           `json:"processingFacility,omitempty"`
	Division                *string                               `json:"division,omitempty"`
	ProductStatus           *Coding                               `json:"productStatus,omitempty"`
	ExpirationDate          *FhirDateTime                         `json:"expirationDate,omitempty"`
	Collection              *BiologicallyDerivedProductCollection `json:"collection,omitempty"`
	StorageTempRequirements *Range                                `json:"storageTempRequirements,omitempty"`
	Property                []BiologicallyDerivedProductProperty  `json:"property,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProductCollection struct {
	Id                *string       `json:"id,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	Collector         *Reference    `json:"collector,omitempty"`
	Source            *Reference    `json:"source,omitempty"`
	CollectedDateTime *FhirDateTime `json:"collectedDateTime,omitempty"`
	CollectedPeriod   *Period       `json:"collectedPeriod,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProductProperty struct {
	Id                   *string         `json:"id,omitempty"`
	Extension            []Extension     `json:"extension,omitempty"`
	ModifierExtension    []Extension     `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept `json:"type"`
	ValueBoolean         bool            `json:"valueBoolean"`
	ValueInteger         int             `json:"valueInteger"`
	ValueCodeableConcept CodeableConcept `json:"valueCodeableConcept"`
	ValuePeriod          Period          `json:"valuePeriod"`
	ValueQuantity        Quantity        `json:"valueQuantity"`
	ValueRange           Range           `json:"valueRange"`
	ValueRatio           Ratio           `json:"valueRatio"`
	ValueString          string          `json:"valueString"`
	ValueAttachment      Attachment      `json:"valueAttachment"`
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
func (resource *BiologicallyDerivedProduct) T_ProductCategory(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("productCategory", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("productCategory", resource.ProductCategory, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProductCode(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("productCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("productCode", resource.ProductCode, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_Parent(frs []FhirResource, numParent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParent >= len(resource.Parent) {
		return ReferenceInput(frs, "parent["+strconv.Itoa(numParent)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "parent["+strconv.Itoa(numParent)+"]", &resource.Parent[numParent], htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_Request(frs []FhirResource, numRequest int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRequest >= len(resource.Request) {
		return ReferenceInput(frs, "request["+strconv.Itoa(numRequest)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "request["+strconv.Itoa(numRequest)+"]", &resource.Request[numRequest], htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_BiologicalSourceEvent(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IdentifierInput("biologicalSourceEvent", nil, htmlAttrs)
	}
	return IdentifierInput("biologicalSourceEvent", resource.BiologicalSourceEvent, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProcessingFacility(frs []FhirResource, numProcessingFacility int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcessingFacility >= len(resource.ProcessingFacility) {
		return ReferenceInput(frs, "processingFacility["+strconv.Itoa(numProcessingFacility)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "processingFacility["+strconv.Itoa(numProcessingFacility)+"]", &resource.ProcessingFacility[numProcessingFacility], htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_Division(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("division", nil, htmlAttrs)
	}
	return StringInput("division", resource.Division, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProductStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("productStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("productStatus", resource.ProductStatus, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ExpirationDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("expirationDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("expirationDate", resource.ExpirationDate, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_StorageTempRequirements(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return RangeInput("storageTempRequirements", nil, htmlAttrs)
	}
	return RangeInput("storageTempRequirements", resource.StorageTempRequirements, htmlAttrs)
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
func (resource *BiologicallyDerivedProduct) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_PropertyValueBoolean(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", &resource.Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_PropertyValueInteger(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return IntInput("property["+strconv.Itoa(numProperty)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("property["+strconv.Itoa(numProperty)+"].valueInteger", &resource.Property[numProperty].ValueInteger, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_PropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", &resource.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_PropertyValuePeriod(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return PeriodInput("property["+strconv.Itoa(numProperty)+"].valuePeriod", nil, htmlAttrs)
	}
	return PeriodInput("property["+strconv.Itoa(numProperty)+"].valuePeriod", &resource.Property[numProperty].ValuePeriod, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_PropertyValueQuantity(numProperty int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return QuantityInput("property["+strconv.Itoa(numProperty)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("property["+strconv.Itoa(numProperty)+"].valueQuantity", &resource.Property[numProperty].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_PropertyValueRange(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return RangeInput("property["+strconv.Itoa(numProperty)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("property["+strconv.Itoa(numProperty)+"].valueRange", &resource.Property[numProperty].ValueRange, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_PropertyValueRatio(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return RatioInput("property["+strconv.Itoa(numProperty)+"].valueRatio", nil, htmlAttrs)
	}
	return RatioInput("property["+strconv.Itoa(numProperty)+"].valueRatio", &resource.Property[numProperty].ValueRatio, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_PropertyValueString(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("property["+strconv.Itoa(numProperty)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("property["+strconv.Itoa(numProperty)+"].valueString", &resource.Property[numProperty].ValueString, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_PropertyValueAttachment(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return AttachmentInput("property["+strconv.Itoa(numProperty)+"].valueAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("property["+strconv.Itoa(numProperty)+"].valueAttachment", &resource.Property[numProperty].ValueAttachment, htmlAttrs)
}
