//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

// http://hl7.org/fhir/r4b/StructureDefinition/NutritionProduct
type NutritionProduct struct {
	Id                    *string                                 `json:"id,omitempty"`
	Meta                  *Meta                                   `json:"meta,omitempty"`
	ImplicitRules         *string                                 `json:"implicitRules,omitempty"`
	Language              *string                                 `json:"language,omitempty"`
	Text                  *Narrative                              `json:"text,omitempty"`
	Contained             []Resource                              `json:"contained,omitempty"`
	Extension             []Extension                             `json:"extension,omitempty"`
	ModifierExtension     []Extension                             `json:"modifierExtension,omitempty"`
	Status                string                                  `json:"status"`
	Category              []CodeableConcept                       `json:"category,omitempty"`
	Code                  *CodeableConcept                        `json:"code,omitempty"`
	Manufacturer          []Reference                             `json:"manufacturer,omitempty"`
	Nutrient              []NutritionProductNutrient              `json:"nutrient,omitempty"`
	Ingredient            []NutritionProductIngredient            `json:"ingredient,omitempty"`
	KnownAllergen         []CodeableReference                     `json:"knownAllergen,omitempty"`
	ProductCharacteristic []NutritionProductProductCharacteristic `json:"productCharacteristic,omitempty"`
	Instance              *NutritionProductInstance               `json:"instance,omitempty"`
	Note                  []Annotation                            `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/NutritionProduct
type NutritionProductNutrient struct {
	Id                *string            `json:"id,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Item              *CodeableReference `json:"item,omitempty"`
	Amount            []Ratio            `json:"amount,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/NutritionProduct
type NutritionProductIngredient struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Item              CodeableReference `json:"item"`
	Amount            []Ratio           `json:"amount,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/NutritionProduct
type NutritionProductProductCharacteristic struct {
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

// http://hl7.org/fhir/r4b/StructureDefinition/NutritionProduct
type NutritionProductInstance struct {
	Id                *string      `json:"id,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Quantity          *Quantity    `json:"quantity,omitempty"`
	Identifier        []Identifier `json:"identifier,omitempty"`
	LotNumber         *string      `json:"lotNumber,omitempty"`
	Expiry            *string      `json:"expiry,omitempty"`
	UseBy             *string      `json:"useBy,omitempty"`
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
