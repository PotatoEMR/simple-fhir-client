package r4b

//generated with command go run ./bultaoreune
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
	Id                *string       `json:"id,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	Quantity          *Quantity     `json:"quantity,omitempty"`
	Identifier        []Identifier  `json:"identifier,omitempty"`
	LotNumber         *string       `json:"lotNumber,omitempty"`
	Expiry            *FhirDateTime `json:"expiry,omitempty"`
	UseBy             *FhirDateTime `json:"useBy,omitempty"`
}

type OtherNutritionProduct NutritionProduct

// struct -> json, automatically add resourceType=Patient
func (r NutritionProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNutritionProduct
		ResourceType string `json:"resourceType"`
	}{
		OtherNutritionProduct: OtherNutritionProduct(r),
		ResourceType:          "NutritionProduct",
	})
}

func (r NutritionProduct) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "NutritionProduct/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "NutritionProduct"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r NutritionProduct) ResourceType() string {
	return "NutritionProduct"
}

func (resource *NutritionProduct) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSNutritionproduct_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_Manufacturer(frs []FhirResource, numManufacturer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numManufacturer >= len(resource.Manufacturer) {
		return ReferenceInput(frs, "manufacturer["+strconv.Itoa(numManufacturer)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "manufacturer["+strconv.Itoa(numManufacturer)+"]", &resource.Manufacturer[numManufacturer], htmlAttrs)
}
func (resource *NutritionProduct) T_KnownAllergen(numKnownAllergen int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numKnownAllergen >= len(resource.KnownAllergen) {
		return CodeableReferenceInput("knownAllergen["+strconv.Itoa(numKnownAllergen)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("knownAllergen["+strconv.Itoa(numKnownAllergen)+"]", &resource.KnownAllergen[numKnownAllergen], htmlAttrs)
}
func (resource *NutritionProduct) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *NutritionProduct) T_NutrientItem(numNutrient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNutrient >= len(resource.Nutrient) {
		return CodeableReferenceInput("nutrient["+strconv.Itoa(numNutrient)+"].item", nil, htmlAttrs)
	}
	return CodeableReferenceInput("nutrient["+strconv.Itoa(numNutrient)+"].item", resource.Nutrient[numNutrient].Item, htmlAttrs)
}
func (resource *NutritionProduct) T_NutrientAmount(numNutrient int, numAmount int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNutrient >= len(resource.Nutrient) || numAmount >= len(resource.Nutrient[numNutrient].Amount) {
		return RatioInput("nutrient["+strconv.Itoa(numNutrient)+"].amount["+strconv.Itoa(numAmount)+"]", nil, htmlAttrs)
	}
	return RatioInput("nutrient["+strconv.Itoa(numNutrient)+"].amount["+strconv.Itoa(numAmount)+"]", &resource.Nutrient[numNutrient].Amount[numAmount], htmlAttrs)
}
func (resource *NutritionProduct) T_IngredientItem(numIngredient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableReferenceInput("ingredient["+strconv.Itoa(numIngredient)+"].item", nil, htmlAttrs)
	}
	return CodeableReferenceInput("ingredient["+strconv.Itoa(numIngredient)+"].item", &resource.Ingredient[numIngredient].Item, htmlAttrs)
}
func (resource *NutritionProduct) T_IngredientAmount(numIngredient int, numAmount int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) || numAmount >= len(resource.Ingredient[numIngredient].Amount) {
		return RatioInput("ingredient["+strconv.Itoa(numIngredient)+"].amount["+strconv.Itoa(numAmount)+"]", nil, htmlAttrs)
	}
	return RatioInput("ingredient["+strconv.Itoa(numIngredient)+"].amount["+strconv.Itoa(numAmount)+"]", &resource.Ingredient[numIngredient].Amount[numAmount], htmlAttrs)
}
func (resource *NutritionProduct) T_ProductCharacteristicType(numProductCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProductCharacteristic >= len(resource.ProductCharacteristic) {
		return CodeableConceptSelect("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].type", &resource.ProductCharacteristic[numProductCharacteristic].Type, optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_ProductCharacteristicValueCodeableConcept(numProductCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProductCharacteristic >= len(resource.ProductCharacteristic) {
		return CodeableConceptSelect("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].valueCodeableConcept", &resource.ProductCharacteristic[numProductCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_ProductCharacteristicValueString(numProductCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProductCharacteristic >= len(resource.ProductCharacteristic) {
		return StringInput("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].valueString", &resource.ProductCharacteristic[numProductCharacteristic].ValueString, htmlAttrs)
}
func (resource *NutritionProduct) T_ProductCharacteristicValueQuantity(numProductCharacteristic int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numProductCharacteristic >= len(resource.ProductCharacteristic) {
		return QuantityInput("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].valueQuantity", &resource.ProductCharacteristic[numProductCharacteristic].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_ProductCharacteristicValueBase64Binary(numProductCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProductCharacteristic >= len(resource.ProductCharacteristic) {
		return StringInput("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].valueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].valueBase64Binary", &resource.ProductCharacteristic[numProductCharacteristic].ValueBase64Binary, htmlAttrs)
}
func (resource *NutritionProduct) T_ProductCharacteristicValueAttachment(numProductCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProductCharacteristic >= len(resource.ProductCharacteristic) {
		return AttachmentInput("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].valueAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].valueAttachment", &resource.ProductCharacteristic[numProductCharacteristic].ValueAttachment, htmlAttrs)
}
func (resource *NutritionProduct) T_ProductCharacteristicValueBoolean(numProductCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProductCharacteristic >= len(resource.ProductCharacteristic) {
		return BoolInput("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].valueBoolean", &resource.ProductCharacteristic[numProductCharacteristic].ValueBoolean, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceQuantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("instance.quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("instance.quantity", resource.Instance.Quantity, optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceLotNumber(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("instance.lotNumber", nil, htmlAttrs)
	}
	return StringInput("instance.lotNumber", resource.Instance.LotNumber, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceExpiry(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("instance.expiry", nil, htmlAttrs)
	}
	return FhirDateTimeInput("instance.expiry", resource.Instance.Expiry, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceUseBy(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("instance.useBy", nil, htmlAttrs)
	}
	return FhirDateTimeInput("instance.useBy", resource.Instance.UseBy, htmlAttrs)
}
