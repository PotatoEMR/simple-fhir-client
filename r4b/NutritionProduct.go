package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Expiry            *time.Time   `json:"expiry,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	UseBy             *time.Time   `json:"useBy,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *NutritionProduct) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSNutritionproduct_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *NutritionProduct) T_ProductCharacteristicType(numProductCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProductCharacteristic >= len(resource.ProductCharacteristic) {
		return CodeableConceptSelect("ProductCharacteristic["+strconv.Itoa(numProductCharacteristic)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ProductCharacteristic["+strconv.Itoa(numProductCharacteristic)+"]Type", &resource.ProductCharacteristic[numProductCharacteristic].Type, optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_ProductCharacteristicValueCodeableConcept(numProductCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProductCharacteristic >= len(resource.ProductCharacteristic) {
		return CodeableConceptSelect("ProductCharacteristic["+strconv.Itoa(numProductCharacteristic)+"]ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ProductCharacteristic["+strconv.Itoa(numProductCharacteristic)+"]ValueCodeableConcept", &resource.ProductCharacteristic[numProductCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_ProductCharacteristicValueString(numProductCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numProductCharacteristic >= len(resource.ProductCharacteristic) {
		return StringInput("ProductCharacteristic["+strconv.Itoa(numProductCharacteristic)+"]ValueString", nil, htmlAttrs)
	}
	return StringInput("ProductCharacteristic["+strconv.Itoa(numProductCharacteristic)+"]ValueString", &resource.ProductCharacteristic[numProductCharacteristic].ValueString, htmlAttrs)
}
func (resource *NutritionProduct) T_ProductCharacteristicValueBase64Binary(numProductCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numProductCharacteristic >= len(resource.ProductCharacteristic) {
		return StringInput("ProductCharacteristic["+strconv.Itoa(numProductCharacteristic)+"]ValueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("ProductCharacteristic["+strconv.Itoa(numProductCharacteristic)+"]ValueBase64Binary", &resource.ProductCharacteristic[numProductCharacteristic].ValueBase64Binary, htmlAttrs)
}
func (resource *NutritionProduct) T_ProductCharacteristicValueBoolean(numProductCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numProductCharacteristic >= len(resource.ProductCharacteristic) {
		return BoolInput("ProductCharacteristic["+strconv.Itoa(numProductCharacteristic)+"]ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("ProductCharacteristic["+strconv.Itoa(numProductCharacteristic)+"]ValueBoolean", &resource.ProductCharacteristic[numProductCharacteristic].ValueBoolean, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceLotNumber(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("InstanceLotNumber", nil, htmlAttrs)
	}
	return StringInput("InstanceLotNumber", resource.Instance.LotNumber, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceExpiry(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("InstanceExpiry", nil, htmlAttrs)
	}
	return DateTimeInput("InstanceExpiry", resource.Instance.Expiry, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceUseBy(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("InstanceUseBy", nil, htmlAttrs)
	}
	return DateTimeInput("InstanceUseBy", resource.Instance.UseBy, htmlAttrs)
}
