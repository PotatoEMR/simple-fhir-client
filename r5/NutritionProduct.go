package r5

//generated with command go run ./bultaoreune
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
	Id                    *string       `json:"id,omitempty"`
	Extension             []Extension   `json:"extension,omitempty"`
	ModifierExtension     []Extension   `json:"modifierExtension,omitempty"`
	Quantity              *Quantity     `json:"quantity,omitempty"`
	Identifier            []Identifier  `json:"identifier,omitempty"`
	Name                  *string       `json:"name,omitempty"`
	LotNumber             *string       `json:"lotNumber,omitempty"`
	Expiry                *FhirDateTime `json:"expiry,omitempty"`
	UseBy                 *FhirDateTime `json:"useBy,omitempty"`
	BiologicalSourceEvent *Identifier   `json:"biologicalSourceEvent,omitempty"`
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
func (resource *NutritionProduct) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
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
func (resource *NutritionProduct) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *NutritionProduct) T_CharacteristicType(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].type", &resource.Characteristic[numCharacteristic].Type, optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_CharacteristicValueCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].valueCodeableConcept", &resource.Characteristic[numCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_CharacteristicValueString(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueString", &resource.Characteristic[numCharacteristic].ValueString, htmlAttrs)
}
func (resource *NutritionProduct) T_CharacteristicValueBase64Binary(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueBase64Binary", &resource.Characteristic[numCharacteristic].ValueBase64Binary, htmlAttrs)
}
func (resource *NutritionProduct) T_CharacteristicValueBoolean(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueBoolean", &resource.Characteristic[numCharacteristic].ValueBoolean, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceName(numInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].name", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].name", resource.Instance[numInstance].Name, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceLotNumber(numInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].lotNumber", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].lotNumber", resource.Instance[numInstance].LotNumber, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceExpiry(numInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return DateTimeInput("instance["+strconv.Itoa(numInstance)+"].expiry", nil, htmlAttrs)
	}
	return DateTimeInput("instance["+strconv.Itoa(numInstance)+"].expiry", resource.Instance[numInstance].Expiry, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceUseBy(numInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return DateTimeInput("instance["+strconv.Itoa(numInstance)+"].useBy", nil, htmlAttrs)
	}
	return DateTimeInput("instance["+strconv.Itoa(numInstance)+"].useBy", resource.Instance[numInstance].UseBy, htmlAttrs)
}
