package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	CollectedDateTime *string     `json:"collectedDateTime"`
	CollectedPeriod   *Period     `json:"collectedPeriod"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProductProcessing struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Procedure         *CodeableConcept `json:"procedure,omitempty"`
	Additive          *Reference       `json:"additive,omitempty"`
	TimeDateTime      *string          `json:"timeDateTime"`
	TimePeriod        *Period          `json:"timePeriod"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProductManipulation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       *string     `json:"description,omitempty"`
	TimeDateTime      *string     `json:"timeDateTime"`
	TimePeriod        *Period     `json:"timePeriod"`
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

func (resource *BiologicallyDerivedProduct) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) T_ProductCategory() templ.Component {
	optionsValueSet := VSProduct_category

	if resource == nil {
		return CodeSelect("productCategory", nil, optionsValueSet)
	}
	return CodeSelect("productCategory", resource.ProductCategory, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) T_ProductCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("productCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productCode", resource.ProductCode, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) T_Status() templ.Component {
	optionsValueSet := VSProduct_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", resource.Status, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) T_ProcessingProcedure(numProcessing int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Processing) >= numProcessing {
		return CodeableConceptSelect("procedure", nil, optionsValueSet)
	}
	return CodeableConceptSelect("procedure", resource.Processing[numProcessing].Procedure, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) T_StorageScale(numStorage int) templ.Component {
	optionsValueSet := VSProduct_storage_scale

	if resource == nil && len(resource.Storage) >= numStorage {
		return CodeSelect("scale", nil, optionsValueSet)
	}
	return CodeSelect("scale", resource.Storage[numStorage].Scale, optionsValueSet)
}
