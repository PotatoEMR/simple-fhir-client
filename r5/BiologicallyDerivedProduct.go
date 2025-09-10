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
	ExpirationDate          *string                               `json:"expirationDate,omitempty"`
	Collection              *BiologicallyDerivedProductCollection `json:"collection,omitempty"`
	StorageTempRequirements *Range                                `json:"storageTempRequirements,omitempty"`
	Property                []BiologicallyDerivedProductProperty  `json:"property,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProductCollection struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Collector         *Reference  `json:"collector,omitempty"`
	Source            *Reference  `json:"source,omitempty"`
	CollectedDateTime *string     `json:"collectedDateTime,omitempty"`
	CollectedPeriod   *Period     `json:"collectedPeriod,omitempty"`
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
		return DateTimeInput("expirationDate", nil, htmlAttrs)
	}
	return DateTimeInput("expirationDate", resource.ExpirationDate, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_CollectionCollectedDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("collection.collectedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("collection.collectedDateTime", resource.Collection.CollectedDateTime, htmlAttrs)
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
func (resource *BiologicallyDerivedProduct) T_PropertyValueString(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("property["+strconv.Itoa(numProperty)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("property["+strconv.Itoa(numProperty)+"].valueString", &resource.Property[numProperty].ValueString, htmlAttrs)
}
