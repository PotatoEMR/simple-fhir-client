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
	ExpirationDate          *time.Time                            `json:"expirationDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	CollectedDateTime *time.Time  `json:"collectedDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *BiologicallyDerivedProduct) T_ProductCategory(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("BiologicallyDerivedProduct.ProductCategory", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("BiologicallyDerivedProduct.ProductCategory", resource.ProductCategory, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProductCode(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("BiologicallyDerivedProduct.ProductCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("BiologicallyDerivedProduct.ProductCode", resource.ProductCode, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_Division(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("BiologicallyDerivedProduct.Division", nil, htmlAttrs)
	}
	return StringInput("BiologicallyDerivedProduct.Division", resource.Division, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ProductStatus(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("BiologicallyDerivedProduct.ProductStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("BiologicallyDerivedProduct.ProductStatus", resource.ProductStatus, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_ExpirationDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("BiologicallyDerivedProduct.ExpirationDate", nil, htmlAttrs)
	}
	return DateTimeInput("BiologicallyDerivedProduct.ExpirationDate", resource.ExpirationDate, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_CollectionCollectedDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("BiologicallyDerivedProduct.Collection.CollectedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("BiologicallyDerivedProduct.Collection.CollectedDateTime", resource.Collection.CollectedDateTime, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("BiologicallyDerivedProduct.Property["+strconv.Itoa(numProperty)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("BiologicallyDerivedProduct.Property["+strconv.Itoa(numProperty)+"].Type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_PropertyValueBoolean(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return BoolInput("BiologicallyDerivedProduct.Property["+strconv.Itoa(numProperty)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("BiologicallyDerivedProduct.Property["+strconv.Itoa(numProperty)+"].ValueBoolean", &resource.Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_PropertyValueInteger(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return IntInput("BiologicallyDerivedProduct.Property["+strconv.Itoa(numProperty)+"].ValueInteger", nil, htmlAttrs)
	}
	return IntInput("BiologicallyDerivedProduct.Property["+strconv.Itoa(numProperty)+"].ValueInteger", &resource.Property[numProperty].ValueInteger, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_PropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("BiologicallyDerivedProduct.Property["+strconv.Itoa(numProperty)+"].ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("BiologicallyDerivedProduct.Property["+strconv.Itoa(numProperty)+"].ValueCodeableConcept", &resource.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *BiologicallyDerivedProduct) T_PropertyValueString(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("BiologicallyDerivedProduct.Property["+strconv.Itoa(numProperty)+"].ValueString", nil, htmlAttrs)
	}
	return StringInput("BiologicallyDerivedProduct.Property["+strconv.Itoa(numProperty)+"].ValueString", &resource.Property[numProperty].ValueString, htmlAttrs)
}
