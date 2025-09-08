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

// http://hl7.org/fhir/r5/StructureDefinition/NutritionOrder
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
	BasedOn                []Reference                   `json:"basedOn,omitempty"`
	GroupIdentifier        *Identifier                   `json:"groupIdentifier,omitempty"`
	Status                 string                        `json:"status"`
	Intent                 string                        `json:"intent"`
	Priority               *string                       `json:"priority,omitempty"`
	Subject                Reference                     `json:"subject"`
	Encounter              *Reference                    `json:"encounter,omitempty"`
	SupportingInformation  []Reference                   `json:"supportingInformation,omitempty"`
	DateTime               time.Time                     `json:"dateTime,format:'2006-01-02T15:04:05Z07:00'"`
	Orderer                *Reference                    `json:"orderer,omitempty"`
	Performer              []CodeableReference           `json:"performer,omitempty"`
	AllergyIntolerance     []Reference                   `json:"allergyIntolerance,omitempty"`
	FoodPreferenceModifier []CodeableConcept             `json:"foodPreferenceModifier,omitempty"`
	ExcludeFoodModifier    []CodeableConcept             `json:"excludeFoodModifier,omitempty"`
	OutsideFoodAllowed     *bool                         `json:"outsideFoodAllowed,omitempty"`
	OralDiet               *NutritionOrderOralDiet       `json:"oralDiet,omitempty"`
	Supplement             []NutritionOrderSupplement    `json:"supplement,omitempty"`
	EnteralFormula         *NutritionOrderEnteralFormula `json:"enteralFormula,omitempty"`
	Note                   []Annotation                  `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionOrder
type NutritionOrderOralDiet struct {
	Id                   *string                          `json:"id,omitempty"`
	Extension            []Extension                      `json:"extension,omitempty"`
	ModifierExtension    []Extension                      `json:"modifierExtension,omitempty"`
	Type                 []CodeableConcept                `json:"type,omitempty"`
	Schedule             *NutritionOrderOralDietSchedule  `json:"schedule,omitempty"`
	Nutrient             []NutritionOrderOralDietNutrient `json:"nutrient,omitempty"`
	Texture              []NutritionOrderOralDietTexture  `json:"texture,omitempty"`
	FluidConsistencyType []CodeableConcept                `json:"fluidConsistencyType,omitempty"`
	Instruction          *string                          `json:"instruction,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionOrder
type NutritionOrderOralDietSchedule struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Timing            []Timing         `json:"timing,omitempty"`
	AsNeeded          *bool            `json:"asNeeded,omitempty"`
	AsNeededFor       *CodeableConcept `json:"asNeededFor,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionOrder
type NutritionOrderOralDietNutrient struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Modifier          *CodeableConcept `json:"modifier,omitempty"`
	Amount            *Quantity        `json:"amount,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionOrder
type NutritionOrderOralDietTexture struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Modifier          *CodeableConcept `json:"modifier,omitempty"`
	FoodType          *CodeableConcept `json:"foodType,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionOrder
type NutritionOrderSupplement struct {
	Id                *string                           `json:"id,omitempty"`
	Extension         []Extension                       `json:"extension,omitempty"`
	ModifierExtension []Extension                       `json:"modifierExtension,omitempty"`
	Type              *CodeableReference                `json:"type,omitempty"`
	ProductName       *string                           `json:"productName,omitempty"`
	Schedule          *NutritionOrderSupplementSchedule `json:"schedule,omitempty"`
	Quantity          *Quantity                         `json:"quantity,omitempty"`
	Instruction       *string                           `json:"instruction,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionOrder
type NutritionOrderSupplementSchedule struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Timing            []Timing         `json:"timing,omitempty"`
	AsNeeded          *bool            `json:"asNeeded,omitempty"`
	AsNeededFor       *CodeableConcept `json:"asNeededFor,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionOrder
type NutritionOrderEnteralFormula struct {
	Id                        *string                                      `json:"id,omitempty"`
	Extension                 []Extension                                  `json:"extension,omitempty"`
	ModifierExtension         []Extension                                  `json:"modifierExtension,omitempty"`
	BaseFormulaType           *CodeableReference                           `json:"baseFormulaType,omitempty"`
	BaseFormulaProductName    *string                                      `json:"baseFormulaProductName,omitempty"`
	DeliveryDevice            []CodeableReference                          `json:"deliveryDevice,omitempty"`
	Additive                  []NutritionOrderEnteralFormulaAdditive       `json:"additive,omitempty"`
	CaloricDensity            *Quantity                                    `json:"caloricDensity,omitempty"`
	RouteOfAdministration     *CodeableConcept                             `json:"routeOfAdministration,omitempty"`
	Administration            []NutritionOrderEnteralFormulaAdministration `json:"administration,omitempty"`
	MaxVolumeToDeliver        *Quantity                                    `json:"maxVolumeToDeliver,omitempty"`
	AdministrationInstruction *string                                      `json:"administrationInstruction,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionOrder
type NutritionOrderEnteralFormulaAdditive struct {
	Id                *string            `json:"id,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Type              *CodeableReference `json:"type,omitempty"`
	ProductName       *string            `json:"productName,omitempty"`
	Quantity          *Quantity          `json:"quantity,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionOrder
type NutritionOrderEnteralFormulaAdministration struct {
	Id                *string                                             `json:"id,omitempty"`
	Extension         []Extension                                         `json:"extension,omitempty"`
	ModifierExtension []Extension                                         `json:"modifierExtension,omitempty"`
	Schedule          *NutritionOrderEnteralFormulaAdministrationSchedule `json:"schedule,omitempty"`
	Quantity          *Quantity                                           `json:"quantity,omitempty"`
	RateQuantity      *Quantity                                           `json:"rateQuantity,omitempty"`
	RateRatio         *Ratio                                              `json:"rateRatio,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionOrder
type NutritionOrderEnteralFormulaAdministrationSchedule struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Timing            []Timing         `json:"timing,omitempty"`
	AsNeeded          *bool            `json:"asNeeded,omitempty"`
	AsNeededFor       *CodeableConcept `json:"asNeededFor,omitempty"`
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
		return StringInput("InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *NutritionOrder) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *NutritionOrder) T_Instantiates(numInstantiates int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiates >= len(resource.Instantiates) {
		return StringInput("Instantiates["+strconv.Itoa(numInstantiates)+"]", nil, htmlAttrs)
	}
	return StringInput("Instantiates["+strconv.Itoa(numInstantiates)+"]", &resource.Instantiates[numInstantiates], htmlAttrs)
}
func (resource *NutritionOrder) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_Intent(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("Intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_Priority(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_DateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("DateTime", nil, htmlAttrs)
	}
	return DateTimeInput("DateTime", &resource.DateTime, htmlAttrs)
}
func (resource *NutritionOrder) T_FoodPreferenceModifier(numFoodPreferenceModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numFoodPreferenceModifier >= len(resource.FoodPreferenceModifier) {
		return CodeableConceptSelect("FoodPreferenceModifier["+strconv.Itoa(numFoodPreferenceModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("FoodPreferenceModifier["+strconv.Itoa(numFoodPreferenceModifier)+"]", &resource.FoodPreferenceModifier[numFoodPreferenceModifier], optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_ExcludeFoodModifier(numExcludeFoodModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numExcludeFoodModifier >= len(resource.ExcludeFoodModifier) {
		return CodeableConceptSelect("ExcludeFoodModifier["+strconv.Itoa(numExcludeFoodModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ExcludeFoodModifier["+strconv.Itoa(numExcludeFoodModifier)+"]", &resource.ExcludeFoodModifier[numExcludeFoodModifier], optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_OutsideFoodAllowed(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("OutsideFoodAllowed", nil, htmlAttrs)
	}
	return BoolInput("OutsideFoodAllowed", resource.OutsideFoodAllowed, htmlAttrs)
}
func (resource *NutritionOrder) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietType(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numType >= len(resource.OralDiet.Type) {
		return CodeableConceptSelect("OralDietType["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OralDietType["+strconv.Itoa(numType)+"]", &resource.OralDiet.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietFluidConsistencyType(numFluidConsistencyType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numFluidConsistencyType >= len(resource.OralDiet.FluidConsistencyType) {
		return CodeableConceptSelect("OralDietFluidConsistencyType["+strconv.Itoa(numFluidConsistencyType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OralDietFluidConsistencyType["+strconv.Itoa(numFluidConsistencyType)+"]", &resource.OralDiet.FluidConsistencyType[numFluidConsistencyType], optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietInstruction(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("OralDietInstruction", nil, htmlAttrs)
	}
	return StringInput("OralDietInstruction", resource.OralDiet.Instruction, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietScheduleAsNeeded(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("OralDietSchedule.AsNeeded", nil, htmlAttrs)
	}
	return BoolInput("OralDietSchedule.AsNeeded", resource.OralDiet.Schedule.AsNeeded, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietScheduleAsNeededFor(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("OralDietSchedule.AsNeededFor", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OralDietSchedule.AsNeededFor", resource.OralDiet.Schedule.AsNeededFor, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietNutrientModifier(numNutrient int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numNutrient >= len(resource.OralDiet.Nutrient) {
		return CodeableConceptSelect("OralDietNutrient["+strconv.Itoa(numNutrient)+"].Modifier", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OralDietNutrient["+strconv.Itoa(numNutrient)+"].Modifier", resource.OralDiet.Nutrient[numNutrient].Modifier, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietTextureModifier(numTexture int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTexture >= len(resource.OralDiet.Texture) {
		return CodeableConceptSelect("OralDietTexture["+strconv.Itoa(numTexture)+"].Modifier", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OralDietTexture["+strconv.Itoa(numTexture)+"].Modifier", resource.OralDiet.Texture[numTexture].Modifier, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietTextureFoodType(numTexture int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTexture >= len(resource.OralDiet.Texture) {
		return CodeableConceptSelect("OralDietTexture["+strconv.Itoa(numTexture)+"].FoodType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OralDietTexture["+strconv.Itoa(numTexture)+"].FoodType", resource.OralDiet.Texture[numTexture].FoodType, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementProductName(numSupplement int, htmlAttrs string) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return StringInput("Supplement["+strconv.Itoa(numSupplement)+"]ProductName", nil, htmlAttrs)
	}
	return StringInput("Supplement["+strconv.Itoa(numSupplement)+"]ProductName", resource.Supplement[numSupplement].ProductName, htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementInstruction(numSupplement int, htmlAttrs string) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return StringInput("Supplement["+strconv.Itoa(numSupplement)+"]Instruction", nil, htmlAttrs)
	}
	return StringInput("Supplement["+strconv.Itoa(numSupplement)+"]Instruction", resource.Supplement[numSupplement].Instruction, htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementScheduleAsNeeded(numSupplement int, htmlAttrs string) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return BoolInput("Supplement["+strconv.Itoa(numSupplement)+"]Schedule.AsNeeded", nil, htmlAttrs)
	}
	return BoolInput("Supplement["+strconv.Itoa(numSupplement)+"]Schedule.AsNeeded", resource.Supplement[numSupplement].Schedule.AsNeeded, htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementScheduleAsNeededFor(numSupplement int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return CodeableConceptSelect("Supplement["+strconv.Itoa(numSupplement)+"]Schedule.AsNeededFor", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Supplement["+strconv.Itoa(numSupplement)+"]Schedule.AsNeededFor", resource.Supplement[numSupplement].Schedule.AsNeededFor, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaBaseFormulaProductName(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EnteralFormulaBaseFormulaProductName", nil, htmlAttrs)
	}
	return StringInput("EnteralFormulaBaseFormulaProductName", resource.EnteralFormula.BaseFormulaProductName, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaRouteOfAdministration(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("EnteralFormulaRouteOfAdministration", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EnteralFormulaRouteOfAdministration", resource.EnteralFormula.RouteOfAdministration, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationInstruction(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EnteralFormulaAdministrationInstruction", nil, htmlAttrs)
	}
	return StringInput("EnteralFormulaAdministrationInstruction", resource.EnteralFormula.AdministrationInstruction, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdditiveProductName(numAdditive int, htmlAttrs string) templ.Component {
	if resource == nil || numAdditive >= len(resource.EnteralFormula.Additive) {
		return StringInput("EnteralFormulaAdditive["+strconv.Itoa(numAdditive)+"].ProductName", nil, htmlAttrs)
	}
	return StringInput("EnteralFormulaAdditive["+strconv.Itoa(numAdditive)+"].ProductName", resource.EnteralFormula.Additive[numAdditive].ProductName, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationScheduleAsNeeded(numAdministration int, htmlAttrs string) templ.Component {
	if resource == nil || numAdministration >= len(resource.EnteralFormula.Administration) {
		return BoolInput("EnteralFormulaAdministration["+strconv.Itoa(numAdministration)+"].Schedule.AsNeeded", nil, htmlAttrs)
	}
	return BoolInput("EnteralFormulaAdministration["+strconv.Itoa(numAdministration)+"].Schedule.AsNeeded", resource.EnteralFormula.Administration[numAdministration].Schedule.AsNeeded, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationScheduleAsNeededFor(numAdministration int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAdministration >= len(resource.EnteralFormula.Administration) {
		return CodeableConceptSelect("EnteralFormulaAdministration["+strconv.Itoa(numAdministration)+"].Schedule.AsNeededFor", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EnteralFormulaAdministration["+strconv.Itoa(numAdministration)+"].Schedule.AsNeededFor", resource.EnteralFormula.Administration[numAdministration].Schedule.AsNeededFor, optionsValueSet, htmlAttrs)
}
