package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/NutritionOrder
type NutritionOrder struct {
	Id                     *string                       `json:"id,omitempty"`
	Meta                   *Meta                         `json:"meta,omitempty"`
	ImplicitRules          *string                       `json:"implicitRules,omitempty"`
	Language               *string                       `json:"language,omitempty"`
	Text                   *Narrative                    `json:"text,omitempty"`
	Contained              []Resource                    `json:"contained,omitempty"`
	Extension              []Extension                   `json:"extension,omitempty"`
	ModifierExtension      []Extension                   `json:"modifierExtension,omitempty"`
	Identifier             []Identifier                  `json:"identifier,omitempty"`
	InstantiatesCanonical  []string                      `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri        []string                      `json:"instantiatesUri,omitempty"`
	Instantiates           []string                      `json:"instantiates,omitempty"`
	Status                 string                        `json:"status"`
	Intent                 string                        `json:"intent"`
	Patient                Reference                     `json:"patient"`
	Encounter              *Reference                    `json:"encounter,omitempty"`
	DateTime               time.Time                     `json:"dateTime,format:'2006-01-02T15:04:05Z07:00'"`
	Orderer                *Reference                    `json:"orderer,omitempty"`
	AllergyIntolerance     []Reference                   `json:"allergyIntolerance,omitempty"`
	FoodPreferenceModifier []CodeableConcept             `json:"foodPreferenceModifier,omitempty"`
	ExcludeFoodModifier    []CodeableConcept             `json:"excludeFoodModifier,omitempty"`
	OralDiet               *NutritionOrderOralDiet       `json:"oralDiet,omitempty"`
	Supplement             []NutritionOrderSupplement    `json:"supplement,omitempty"`
	EnteralFormula         *NutritionOrderEnteralFormula `json:"enteralFormula,omitempty"`
	Note                   []Annotation                  `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/NutritionOrder
type NutritionOrderOralDiet struct {
	Id                   *string                          `json:"id,omitempty"`
	Extension            []Extension                      `json:"extension,omitempty"`
	ModifierExtension    []Extension                      `json:"modifierExtension,omitempty"`
	Type                 []CodeableConcept                `json:"type,omitempty"`
	Schedule             []Timing                         `json:"schedule,omitempty"`
	Nutrient             []NutritionOrderOralDietNutrient `json:"nutrient,omitempty"`
	Texture              []NutritionOrderOralDietTexture  `json:"texture,omitempty"`
	FluidConsistencyType []CodeableConcept                `json:"fluidConsistencyType,omitempty"`
	Instruction          *string                          `json:"instruction,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/NutritionOrder
type NutritionOrderOralDietNutrient struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Modifier          *CodeableConcept `json:"modifier,omitempty"`
	Amount            *Quantity        `json:"amount,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/NutritionOrder
type NutritionOrderOralDietTexture struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Modifier          *CodeableConcept `json:"modifier,omitempty"`
	FoodType          *CodeableConcept `json:"foodType,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/NutritionOrder
type NutritionOrderSupplement struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	ProductName       *string          `json:"productName,omitempty"`
	Schedule          []Timing         `json:"schedule,omitempty"`
	Quantity          *Quantity        `json:"quantity,omitempty"`
	Instruction       *string          `json:"instruction,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/NutritionOrder
type NutritionOrderEnteralFormula struct {
	Id                        *string                                      `json:"id,omitempty"`
	Extension                 []Extension                                  `json:"extension,omitempty"`
	ModifierExtension         []Extension                                  `json:"modifierExtension,omitempty"`
	BaseFormulaType           *CodeableConcept                             `json:"baseFormulaType,omitempty"`
	BaseFormulaProductName    *string                                      `json:"baseFormulaProductName,omitempty"`
	AdditiveType              *CodeableConcept                             `json:"additiveType,omitempty"`
	AdditiveProductName       *string                                      `json:"additiveProductName,omitempty"`
	CaloricDensity            *Quantity                                    `json:"caloricDensity,omitempty"`
	RouteofAdministration     *CodeableConcept                             `json:"routeofAdministration,omitempty"`
	Administration            []NutritionOrderEnteralFormulaAdministration `json:"administration,omitempty"`
	MaxVolumeToDeliver        *Quantity                                    `json:"maxVolumeToDeliver,omitempty"`
	AdministrationInstruction *string                                      `json:"administrationInstruction,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/NutritionOrder
type NutritionOrderEnteralFormulaAdministration struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Schedule          *Timing     `json:"schedule,omitempty"`
	Quantity          *Quantity   `json:"quantity,omitempty"`
	RateQuantity      *Quantity   `json:"rateQuantity,omitempty"`
	RateRatio         *Ratio      `json:"rateRatio,omitempty"`
}

type OtherNutritionOrder NutritionOrder

// on convert struct to json, automatically add resourceType=NutritionOrder
func (r NutritionOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNutritionOrder
		ResourceType string `json:"resourceType"`
	}{
		OtherNutritionOrder: OtherNutritionOrder(r),
		ResourceType:        "NutritionOrder",
	})
}
func (r NutritionOrder) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "NutritionOrder/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "NutritionOrder"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *NutritionOrder) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("NutritionOrder.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("NutritionOrder.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *NutritionOrder) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("NutritionOrder.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("NutritionOrder.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *NutritionOrder) T_Instantiates(numInstantiates int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiates >= len(resource.Instantiates) {
		return StringInput("NutritionOrder.Instantiates["+strconv.Itoa(numInstantiates)+"]", nil, htmlAttrs)
	}
	return StringInput("NutritionOrder.Instantiates["+strconv.Itoa(numInstantiates)+"]", &resource.Instantiates[numInstantiates], htmlAttrs)
}
func (resource *NutritionOrder) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("NutritionOrder.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("NutritionOrder.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_Intent(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("NutritionOrder.Intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("NutritionOrder.Intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_DateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("NutritionOrder.DateTime", nil, htmlAttrs)
	}
	return DateTimeInput("NutritionOrder.DateTime", &resource.DateTime, htmlAttrs)
}
func (resource *NutritionOrder) T_FoodPreferenceModifier(numFoodPreferenceModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numFoodPreferenceModifier >= len(resource.FoodPreferenceModifier) {
		return CodeableConceptSelect("NutritionOrder.FoodPreferenceModifier["+strconv.Itoa(numFoodPreferenceModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NutritionOrder.FoodPreferenceModifier["+strconv.Itoa(numFoodPreferenceModifier)+"]", &resource.FoodPreferenceModifier[numFoodPreferenceModifier], optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_ExcludeFoodModifier(numExcludeFoodModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numExcludeFoodModifier >= len(resource.ExcludeFoodModifier) {
		return CodeableConceptSelect("NutritionOrder.ExcludeFoodModifier["+strconv.Itoa(numExcludeFoodModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NutritionOrder.ExcludeFoodModifier["+strconv.Itoa(numExcludeFoodModifier)+"]", &resource.ExcludeFoodModifier[numExcludeFoodModifier], optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("NutritionOrder.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("NutritionOrder.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietType(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numType >= len(resource.OralDiet.Type) {
		return CodeableConceptSelect("NutritionOrder.OralDiet.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NutritionOrder.OralDiet.Type["+strconv.Itoa(numType)+"]", &resource.OralDiet.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietFluidConsistencyType(numFluidConsistencyType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numFluidConsistencyType >= len(resource.OralDiet.FluidConsistencyType) {
		return CodeableConceptSelect("NutritionOrder.OralDiet.FluidConsistencyType["+strconv.Itoa(numFluidConsistencyType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NutritionOrder.OralDiet.FluidConsistencyType["+strconv.Itoa(numFluidConsistencyType)+"]", &resource.OralDiet.FluidConsistencyType[numFluidConsistencyType], optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietInstruction(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("NutritionOrder.OralDiet.Instruction", nil, htmlAttrs)
	}
	return StringInput("NutritionOrder.OralDiet.Instruction", resource.OralDiet.Instruction, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietNutrientModifier(numNutrient int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numNutrient >= len(resource.OralDiet.Nutrient) {
		return CodeableConceptSelect("NutritionOrder.OralDiet.Nutrient["+strconv.Itoa(numNutrient)+"].Modifier", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NutritionOrder.OralDiet.Nutrient["+strconv.Itoa(numNutrient)+"].Modifier", resource.OralDiet.Nutrient[numNutrient].Modifier, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietTextureModifier(numTexture int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTexture >= len(resource.OralDiet.Texture) {
		return CodeableConceptSelect("NutritionOrder.OralDiet.Texture["+strconv.Itoa(numTexture)+"].Modifier", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NutritionOrder.OralDiet.Texture["+strconv.Itoa(numTexture)+"].Modifier", resource.OralDiet.Texture[numTexture].Modifier, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietTextureFoodType(numTexture int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTexture >= len(resource.OralDiet.Texture) {
		return CodeableConceptSelect("NutritionOrder.OralDiet.Texture["+strconv.Itoa(numTexture)+"].FoodType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NutritionOrder.OralDiet.Texture["+strconv.Itoa(numTexture)+"].FoodType", resource.OralDiet.Texture[numTexture].FoodType, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementType(numSupplement int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return CodeableConceptSelect("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].Type", resource.Supplement[numSupplement].Type, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementProductName(numSupplement int, htmlAttrs string) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return StringInput("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].ProductName", nil, htmlAttrs)
	}
	return StringInput("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].ProductName", resource.Supplement[numSupplement].ProductName, htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementInstruction(numSupplement int, htmlAttrs string) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return StringInput("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].Instruction", nil, htmlAttrs)
	}
	return StringInput("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].Instruction", resource.Supplement[numSupplement].Instruction, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaBaseFormulaType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("NutritionOrder.EnteralFormula.BaseFormulaType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NutritionOrder.EnteralFormula.BaseFormulaType", resource.EnteralFormula.BaseFormulaType, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaBaseFormulaProductName(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("NutritionOrder.EnteralFormula.BaseFormulaProductName", nil, htmlAttrs)
	}
	return StringInput("NutritionOrder.EnteralFormula.BaseFormulaProductName", resource.EnteralFormula.BaseFormulaProductName, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdditiveType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("NutritionOrder.EnteralFormula.AdditiveType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NutritionOrder.EnteralFormula.AdditiveType", resource.EnteralFormula.AdditiveType, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdditiveProductName(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("NutritionOrder.EnteralFormula.AdditiveProductName", nil, htmlAttrs)
	}
	return StringInput("NutritionOrder.EnteralFormula.AdditiveProductName", resource.EnteralFormula.AdditiveProductName, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaRouteofAdministration(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("NutritionOrder.EnteralFormula.RouteofAdministration", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NutritionOrder.EnteralFormula.RouteofAdministration", resource.EnteralFormula.RouteofAdministration, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationInstruction(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("NutritionOrder.EnteralFormula.AdministrationInstruction", nil, htmlAttrs)
	}
	return StringInput("NutritionOrder.EnteralFormula.AdministrationInstruction", resource.EnteralFormula.AdministrationInstruction, htmlAttrs)
}
