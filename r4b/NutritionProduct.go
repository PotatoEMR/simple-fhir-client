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
func (resource *NutritionProduct) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
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
func (resource *NutritionProduct) T_ProductCharacteristicValueBase64Binary(numProductCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProductCharacteristic >= len(resource.ProductCharacteristic) {
		return StringInput("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].valueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].valueBase64Binary", &resource.ProductCharacteristic[numProductCharacteristic].ValueBase64Binary, htmlAttrs)
}
func (resource *NutritionProduct) T_ProductCharacteristicValueBoolean(numProductCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProductCharacteristic >= len(resource.ProductCharacteristic) {
		return BoolInput("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("productCharacteristic["+strconv.Itoa(numProductCharacteristic)+"].valueBoolean", &resource.ProductCharacteristic[numProductCharacteristic].ValueBoolean, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceLotNumber(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("instance.lotNumber", nil, htmlAttrs)
	}
	return StringInput("instance.lotNumber", resource.Instance.LotNumber, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceExpiry(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("instance.expiry", nil, htmlAttrs)
	}
	return DateTimeInput("instance.expiry", resource.Instance.Expiry, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceUseBy(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("instance.useBy", nil, htmlAttrs)
	}
	return DateTimeInput("instance.useBy", resource.Instance.UseBy, htmlAttrs)
}
