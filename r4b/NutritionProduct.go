package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *NutritionProduct) T_Id() templ.Component {

	if resource == nil {
		return StringInput("NutritionProduct.Id", nil)
	}
	return StringInput("NutritionProduct.Id", resource.Id)
}
func (resource *NutritionProduct) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("NutritionProduct.ImplicitRules", nil)
	}
	return StringInput("NutritionProduct.ImplicitRules", resource.ImplicitRules)
}
func (resource *NutritionProduct) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("NutritionProduct.Language", nil, optionsValueSet)
	}
	return CodeSelect("NutritionProduct.Language", resource.Language, optionsValueSet)
}
func (resource *NutritionProduct) T_Status() templ.Component {
	optionsValueSet := VSNutritionproduct_status

	if resource == nil {
		return CodeSelect("NutritionProduct.Status", nil, optionsValueSet)
	}
	return CodeSelect("NutritionProduct.Status", &resource.Status, optionsValueSet)
}
func (resource *NutritionProduct) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("NutritionProduct.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionProduct.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *NutritionProduct) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("NutritionProduct.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionProduct.Code", resource.Code, optionsValueSet)
}
func (resource *NutritionProduct) T_NutrientId(numNutrient int) templ.Component {

	if resource == nil || len(resource.Nutrient) >= numNutrient {
		return StringInput("NutritionProduct.Nutrient["+strconv.Itoa(numNutrient)+"].Id", nil)
	}
	return StringInput("NutritionProduct.Nutrient["+strconv.Itoa(numNutrient)+"].Id", resource.Nutrient[numNutrient].Id)
}
func (resource *NutritionProduct) T_IngredientId(numIngredient int) templ.Component {

	if resource == nil || len(resource.Ingredient) >= numIngredient {
		return StringInput("NutritionProduct.Ingredient["+strconv.Itoa(numIngredient)+"].Id", nil)
	}
	return StringInput("NutritionProduct.Ingredient["+strconv.Itoa(numIngredient)+"].Id", resource.Ingredient[numIngredient].Id)
}
func (resource *NutritionProduct) T_ProductCharacteristicId(numProductCharacteristic int) templ.Component {

	if resource == nil || len(resource.ProductCharacteristic) >= numProductCharacteristic {
		return StringInput("NutritionProduct.ProductCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].Id", nil)
	}
	return StringInput("NutritionProduct.ProductCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].Id", resource.ProductCharacteristic[numProductCharacteristic].Id)
}
func (resource *NutritionProduct) T_ProductCharacteristicType(numProductCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ProductCharacteristic) >= numProductCharacteristic {
		return CodeableConceptSelect("NutritionProduct.ProductCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionProduct.ProductCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].Type", &resource.ProductCharacteristic[numProductCharacteristic].Type, optionsValueSet)
}
func (resource *NutritionProduct) T_InstanceId() templ.Component {

	if resource == nil {
		return StringInput("NutritionProduct.Instance.Id", nil)
	}
	return StringInput("NutritionProduct.Instance.Id", resource.Instance.Id)
}
func (resource *NutritionProduct) T_InstanceLotNumber() templ.Component {

	if resource == nil {
		return StringInput("NutritionProduct.Instance.LotNumber", nil)
	}
	return StringInput("NutritionProduct.Instance.LotNumber", resource.Instance.LotNumber)
}
func (resource *NutritionProduct) T_InstanceExpiry() templ.Component {

	if resource == nil {
		return StringInput("NutritionProduct.Instance.Expiry", nil)
	}
	return StringInput("NutritionProduct.Instance.Expiry", resource.Instance.Expiry)
}
func (resource *NutritionProduct) T_InstanceUseBy() templ.Component {

	if resource == nil {
		return StringInput("NutritionProduct.Instance.UseBy", nil)
	}
	return StringInput("NutritionProduct.Instance.UseBy", resource.Instance.UseBy)
}
