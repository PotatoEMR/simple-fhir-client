package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	CollectedDateTime *string     `json:"collectedDateTime"`
	CollectedPeriod   *Period     `json:"collectedPeriod"`
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

func (resource *BiologicallyDerivedProduct) BiologicallyDerivedProductLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) BiologicallyDerivedProductProductCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodingSelect("productCategory", nil, optionsValueSet)
	}
	return CodingSelect("productCategory", resource.ProductCategory, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) BiologicallyDerivedProductProductCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("productCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productCode", resource.ProductCode, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) BiologicallyDerivedProductProductStatus(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodingSelect("productStatus", nil, optionsValueSet)
	}
	return CodingSelect("productStatus", resource.ProductStatus, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) BiologicallyDerivedProductPropertyType(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Property) >= numProperty {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Property[numProperty].Type, optionsValueSet)
}
