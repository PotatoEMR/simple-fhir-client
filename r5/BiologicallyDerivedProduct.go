package r5

//generated with command go run ./bultaoreune -nodownload
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
func (resource *BiologicallyDerivedProduct) T_ProductCategory(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("BiologicallyDerivedProduct.ProductCategory", nil, optionsValueSet)
	}
	return CodingSelect("BiologicallyDerivedProduct.ProductCategory", resource.ProductCategory, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) T_ProductCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("BiologicallyDerivedProduct.ProductCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("BiologicallyDerivedProduct.ProductCode", resource.ProductCode, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) T_Division() templ.Component {

	if resource == nil {
		return StringInput("BiologicallyDerivedProduct.Division", nil)
	}
	return StringInput("BiologicallyDerivedProduct.Division", resource.Division)
}
func (resource *BiologicallyDerivedProduct) T_ProductStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("BiologicallyDerivedProduct.ProductStatus", nil, optionsValueSet)
	}
	return CodingSelect("BiologicallyDerivedProduct.ProductStatus", resource.ProductStatus, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) T_ExpirationDate() templ.Component {

	if resource == nil {
		return StringInput("BiologicallyDerivedProduct.ExpirationDate", nil)
	}
	return StringInput("BiologicallyDerivedProduct.ExpirationDate", resource.ExpirationDate)
}
func (resource *BiologicallyDerivedProduct) T_CollectionId() templ.Component {

	if resource == nil {
		return StringInput("BiologicallyDerivedProduct.Collection.Id", nil)
	}
	return StringInput("BiologicallyDerivedProduct.Collection.Id", resource.Collection.Id)
}
func (resource *BiologicallyDerivedProduct) T_PropertyId(numProperty int) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return StringInput("BiologicallyDerivedProduct.Property["+strconv.Itoa(numProperty)+"].Id", nil)
	}
	return StringInput("BiologicallyDerivedProduct.Property["+strconv.Itoa(numProperty)+"].Id", resource.Property[numProperty].Id)
}
func (resource *BiologicallyDerivedProduct) T_PropertyType(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return CodeableConceptSelect("BiologicallyDerivedProduct.Property["+strconv.Itoa(numProperty)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("BiologicallyDerivedProduct.Property["+strconv.Itoa(numProperty)+"].Type", &resource.Property[numProperty].Type, optionsValueSet)
}
