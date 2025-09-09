package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Expiry                *time.Time   `json:"expiry,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	UseBy                 *time.Time   `json:"useBy,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *NutritionProduct) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("NutritionProduct.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NutritionProduct.Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSNutritionproduct_status

	if resource == nil {
		return CodeSelect("NutritionProduct.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("NutritionProduct.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("NutritionProduct.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NutritionProduct.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("NutritionProduct.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("NutritionProduct.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *NutritionProduct) T_CharacteristicType(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("NutritionProduct.Characteristic["+strconv.Itoa(numCharacteristic)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NutritionProduct.Characteristic["+strconv.Itoa(numCharacteristic)+"].Type", &resource.Characteristic[numCharacteristic].Type, optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_CharacteristicValueCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("NutritionProduct.Characteristic["+strconv.Itoa(numCharacteristic)+"].ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NutritionProduct.Characteristic["+strconv.Itoa(numCharacteristic)+"].ValueCodeableConcept", &resource.Characteristic[numCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *NutritionProduct) T_CharacteristicValueString(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("NutritionProduct.Characteristic["+strconv.Itoa(numCharacteristic)+"].ValueString", nil, htmlAttrs)
	}
	return StringInput("NutritionProduct.Characteristic["+strconv.Itoa(numCharacteristic)+"].ValueString", &resource.Characteristic[numCharacteristic].ValueString, htmlAttrs)
}
func (resource *NutritionProduct) T_CharacteristicValueBase64Binary(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("NutritionProduct.Characteristic["+strconv.Itoa(numCharacteristic)+"].ValueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("NutritionProduct.Characteristic["+strconv.Itoa(numCharacteristic)+"].ValueBase64Binary", &resource.Characteristic[numCharacteristic].ValueBase64Binary, htmlAttrs)
}
func (resource *NutritionProduct) T_CharacteristicValueBoolean(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("NutritionProduct.Characteristic["+strconv.Itoa(numCharacteristic)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("NutritionProduct.Characteristic["+strconv.Itoa(numCharacteristic)+"].ValueBoolean", &resource.Characteristic[numCharacteristic].ValueBoolean, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceName(numInstance int, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].Name", nil, htmlAttrs)
	}
	return StringInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].Name", resource.Instance[numInstance].Name, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceLotNumber(numInstance int, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].LotNumber", nil, htmlAttrs)
	}
	return StringInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].LotNumber", resource.Instance[numInstance].LotNumber, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceExpiry(numInstance int, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return DateTimeInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].Expiry", nil, htmlAttrs)
	}
	return DateTimeInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].Expiry", resource.Instance[numInstance].Expiry, htmlAttrs)
}
func (resource *NutritionProduct) T_InstanceUseBy(numInstance int, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return DateTimeInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].UseBy", nil, htmlAttrs)
	}
	return DateTimeInput("NutritionProduct.Instance["+strconv.Itoa(numInstance)+"].UseBy", resource.Instance[numInstance].UseBy, htmlAttrs)
}
