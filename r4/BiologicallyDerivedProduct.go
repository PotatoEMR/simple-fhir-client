package r4

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	CollectedDateTime *string     `json:"collectedDateTime"`
	CollectedPeriod   *Period     `json:"collectedPeriod"`
}

// http://hl7.org/fhir/r4/StructureDefinition/BiologicallyDerivedProduct
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

// http://hl7.org/fhir/r4/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProductManipulation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       *string     `json:"description,omitempty"`
	TimeDateTime      *string     `json:"timeDateTime"`
	TimePeriod        *Period     `json:"timePeriod"`
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

func (resource *BiologicallyDerivedProduct) BiologicallyDerivedProductLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) BiologicallyDerivedProductProductCategory() templ.Component {
	optionsValueSet := VSProduct_category
	currentVal := ""
	if resource != nil {
		currentVal = *resource.ProductCategory
	}
	return CodeSelect("productCategory", currentVal, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) BiologicallyDerivedProductStatus() templ.Component {
	optionsValueSet := VSProduct_status
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *BiologicallyDerivedProduct) BiologicallyDerivedProductStorageScale(numStorage int) templ.Component {
	optionsValueSet := VSProduct_storage_scale
	currentVal := ""
	if resource != nil && len(resource.Storage) >= numStorage {
		currentVal = *resource.Storage[numStorage].Scale
	}
	return CodeSelect("scale", currentVal, optionsValueSet)
}
