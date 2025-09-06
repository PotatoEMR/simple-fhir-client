package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
func (resource *NutritionProduct) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("NutritionProduct.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionProduct.Code", resource.Code, optionsValueSet)
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
func (resource *NutritionProduct) T_CharacteristicId(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return StringInput("NutritionProduct.Characteristic["+strconv.Itoa(numCharacteristic)+"].Id", nil)
	}
	return StringInput("NutritionProduct.Characteristic["+strconv.Itoa(numCharacteristic)+"].Id", resource.Characteristic[numCharacteristic].Id)
}
func (resource *NutritionProduct) T_CharacteristicType(numCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return CodeableConceptSelect("NutritionProduct.Characteristic["+strconv.Itoa(numCharacteristic)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionProduct.Characteristic["+strconv.Itoa(numCharacteristic)+"].Type", &resource.Characteristic[numCharacteristic].Type, optionsValueSet)
}
func (resource *NutritionProduct) T_InstanceId(numInstance int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance {
		return StringInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].Id", nil)
	}
	return StringInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].Id", resource.Instance[numInstance].Id)
}
func (resource *NutritionProduct) T_InstanceName(numInstance int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance {
		return StringInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].Name", nil)
	}
	return StringInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].Name", resource.Instance[numInstance].Name)
}
func (resource *NutritionProduct) T_InstanceLotNumber(numInstance int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance {
		return StringInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].LotNumber", nil)
	}
	return StringInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].LotNumber", resource.Instance[numInstance].LotNumber)
}
func (resource *NutritionProduct) T_InstanceExpiry(numInstance int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance {
		return StringInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].Expiry", nil)
	}
	return StringInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].Expiry", resource.Instance[numInstance].Expiry)
}
func (resource *NutritionProduct) T_InstanceUseBy(numInstance int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance {
		return StringInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].UseBy", nil)
	}
	return StringInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].UseBy", resource.Instance[numInstance].UseBy)
}
