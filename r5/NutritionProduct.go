//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/NutritionProduct
type NutritionProduct struct {
	Id                *string                          `json:"id,omitempty"`
	Meta              *Meta                            `json:"meta,omitempty"`
	ImplicitRules     *string                          `json:"implicitRules,omitempty"`
	Language          *string                          `json:"language,omitempty"`
	Text              *Narrative                       `json:"text,omitempty"`
	Contained         []Resource                       `json:"contained,omitempty"`
	Extension         []Extension                      `json:"extension,omitempty"`
	ModifierExtension []Extension                      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept                 `json:"code,omitempty"`
	Status            string                           `json:"status"`
	Category          []CodeableConcept                `json:"category,omitempty"`
	Manufacturer      []Reference                      `json:"manufacturer,omitempty"`
	Nutrient          []NutritionProductNutrient       `json:"nutrient,omitempty"`
	Ingredient        []NutritionProductIngredient     `json:"ingredient,omitempty"`
	KnownAllergen     []CodeableReference              `json:"knownAllergen,omitempty"`
	Characteristic    []NutritionProductCharacteristic `json:"characteristic,omitempty"`
	Instance          []NutritionProductInstance       `json:"instance,omitempty"`
	Note              []Annotation                     `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionProduct
type NutritionProductNutrient struct {
	Id                *string            `json:"id,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Item              *CodeableReference `json:"item,omitempty"`
	Amount            []Ratio            `json:"amount,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionProduct
type NutritionProductIngredient struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Item              CodeableReference `json:"item"`
	Amount            []Ratio           `json:"amount,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionProduct
type NutritionProductCharacteristic struct {
	Id                   *string         `json:"id,omitempty"`
	Extension            []Extension     `json:"extension,omitempty"`
	ModifierExtension    []Extension     `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept `json:"type"`
	ValueCodeableConcept CodeableConcept `json:"valueCodeableConcept"`
	ValueString          string          `json:"valueString"`
	ValueQuantity        Quantity        `json:"valueQuantity"`
	ValueBase64Binary    string          `json:"valueBase64Binary"`
	ValueAttachment      Attachment      `json:"valueAttachment"`
	ValueBoolean         bool            `json:"valueBoolean"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionProduct
type NutritionProductInstance struct {
	Id                    *string      `json:"id,omitempty"`
	Extension             []Extension  `json:"extension,omitempty"`
	ModifierExtension     []Extension  `json:"modifierExtension,omitempty"`
	Quantity              *Quantity    `json:"quantity,omitempty"`
	Identifier            []Identifier `json:"identifier,omitempty"`
	Name                  *string      `json:"name,omitempty"`
	LotNumber             *string      `json:"lotNumber,omitempty"`
	Expiry                *string      `json:"expiry,omitempty"`
	UseBy                 *string      `json:"useBy,omitempty"`
	BiologicalSourceEvent *Identifier  `json:"biologicalSourceEvent,omitempty"`
}

type OtherNutritionProduct NutritionProduct

// on convert struct to json, automatically add resourceType=NutritionProduct
func (r NutritionProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNutritionProduct
		ResourceType string `json:"resourceType"`
	}{
		OtherNutritionProduct: OtherNutritionProduct(r),
		ResourceType:          "NutritionProduct",
	})
}
