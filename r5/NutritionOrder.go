package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	DateTime               string                        `json:"dateTime"`
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

func (resource *NutritionOrder) T_Id() templ.Component {

	if resource == nil {
		return StringInput("NutritionOrder.Id", nil)
	}
	return StringInput("NutritionOrder.Id", resource.Id)
}
func (resource *NutritionOrder) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("NutritionOrder.ImplicitRules", nil)
	}
	return StringInput("NutritionOrder.ImplicitRules", resource.ImplicitRules)
}
func (resource *NutritionOrder) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("NutritionOrder.Language", nil, optionsValueSet)
	}
	return CodeSelect("NutritionOrder.Language", resource.Language, optionsValueSet)
}
func (resource *NutritionOrder) T_InstantiatesCanonical(numInstantiatesCanonical int) templ.Component {

	if resource == nil || len(resource.InstantiatesCanonical) >= numInstantiatesCanonical {
		return StringInput("NutritionOrder.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil)
	}
	return StringInput("NutritionOrder.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical])
}
func (resource *NutritionOrder) T_InstantiatesUri(numInstantiatesUri int) templ.Component {

	if resource == nil || len(resource.InstantiatesUri) >= numInstantiatesUri {
		return StringInput("NutritionOrder.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil)
	}
	return StringInput("NutritionOrder.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri])
}
func (resource *NutritionOrder) T_Instantiates(numInstantiates int) templ.Component {

	if resource == nil || len(resource.Instantiates) >= numInstantiates {
		return StringInput("NutritionOrder.Instantiates["+strconv.Itoa(numInstantiates)+"]", nil)
	}
	return StringInput("NutritionOrder.Instantiates["+strconv.Itoa(numInstantiates)+"]", &resource.Instantiates[numInstantiates])
}
func (resource *NutritionOrder) T_Status() templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("NutritionOrder.Status", nil, optionsValueSet)
	}
	return CodeSelect("NutritionOrder.Status", &resource.Status, optionsValueSet)
}
func (resource *NutritionOrder) T_Intent() templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("NutritionOrder.Intent", nil, optionsValueSet)
	}
	return CodeSelect("NutritionOrder.Intent", &resource.Intent, optionsValueSet)
}
func (resource *NutritionOrder) T_Priority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("NutritionOrder.Priority", nil, optionsValueSet)
	}
	return CodeSelect("NutritionOrder.Priority", resource.Priority, optionsValueSet)
}
func (resource *NutritionOrder) T_DateTime() templ.Component {

	if resource == nil {
		return StringInput("NutritionOrder.DateTime", nil)
	}
	return StringInput("NutritionOrder.DateTime", &resource.DateTime)
}
func (resource *NutritionOrder) T_FoodPreferenceModifier(numFoodPreferenceModifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.FoodPreferenceModifier) >= numFoodPreferenceModifier {
		return CodeableConceptSelect("NutritionOrder.FoodPreferenceModifier["+strconv.Itoa(numFoodPreferenceModifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionOrder.FoodPreferenceModifier["+strconv.Itoa(numFoodPreferenceModifier)+"]", &resource.FoodPreferenceModifier[numFoodPreferenceModifier], optionsValueSet)
}
func (resource *NutritionOrder) T_ExcludeFoodModifier(numExcludeFoodModifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ExcludeFoodModifier) >= numExcludeFoodModifier {
		return CodeableConceptSelect("NutritionOrder.ExcludeFoodModifier["+strconv.Itoa(numExcludeFoodModifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionOrder.ExcludeFoodModifier["+strconv.Itoa(numExcludeFoodModifier)+"]", &resource.ExcludeFoodModifier[numExcludeFoodModifier], optionsValueSet)
}
func (resource *NutritionOrder) T_OutsideFoodAllowed() templ.Component {

	if resource == nil {
		return BoolInput("NutritionOrder.OutsideFoodAllowed", nil)
	}
	return BoolInput("NutritionOrder.OutsideFoodAllowed", resource.OutsideFoodAllowed)
}
func (resource *NutritionOrder) T_OralDietId() templ.Component {

	if resource == nil {
		return StringInput("NutritionOrder.OralDiet.Id", nil)
	}
	return StringInput("NutritionOrder.OralDiet.Id", resource.OralDiet.Id)
}
func (resource *NutritionOrder) T_OralDietType(numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.OralDiet.Type) >= numType {
		return CodeableConceptSelect("NutritionOrder.OralDiet.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionOrder.OralDiet.Type["+strconv.Itoa(numType)+"]", &resource.OralDiet.Type[numType], optionsValueSet)
}
func (resource *NutritionOrder) T_OralDietFluidConsistencyType(numFluidConsistencyType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.OralDiet.FluidConsistencyType) >= numFluidConsistencyType {
		return CodeableConceptSelect("NutritionOrder.OralDiet.FluidConsistencyType["+strconv.Itoa(numFluidConsistencyType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionOrder.OralDiet.FluidConsistencyType["+strconv.Itoa(numFluidConsistencyType)+"]", &resource.OralDiet.FluidConsistencyType[numFluidConsistencyType], optionsValueSet)
}
func (resource *NutritionOrder) T_OralDietInstruction() templ.Component {

	if resource == nil {
		return StringInput("NutritionOrder.OralDiet.Instruction", nil)
	}
	return StringInput("NutritionOrder.OralDiet.Instruction", resource.OralDiet.Instruction)
}
func (resource *NutritionOrder) T_OralDietScheduleId() templ.Component {

	if resource == nil {
		return StringInput("NutritionOrder.OralDiet.Schedule.Id", nil)
	}
	return StringInput("NutritionOrder.OralDiet.Schedule.Id", resource.OralDiet.Schedule.Id)
}
func (resource *NutritionOrder) T_OralDietScheduleAsNeeded() templ.Component {

	if resource == nil {
		return BoolInput("NutritionOrder.OralDiet.Schedule.AsNeeded", nil)
	}
	return BoolInput("NutritionOrder.OralDiet.Schedule.AsNeeded", resource.OralDiet.Schedule.AsNeeded)
}
func (resource *NutritionOrder) T_OralDietScheduleAsNeededFor(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("NutritionOrder.OralDiet.Schedule.AsNeededFor", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionOrder.OralDiet.Schedule.AsNeededFor", resource.OralDiet.Schedule.AsNeededFor, optionsValueSet)
}
func (resource *NutritionOrder) T_OralDietNutrientId(numNutrient int) templ.Component {

	if resource == nil || len(resource.OralDiet.Nutrient) >= numNutrient {
		return StringInput("NutritionOrder.OralDiet.Nutrient["+strconv.Itoa(numNutrient)+"].Id", nil)
	}
	return StringInput("NutritionOrder.OralDiet.Nutrient["+strconv.Itoa(numNutrient)+"].Id", resource.OralDiet.Nutrient[numNutrient].Id)
}
func (resource *NutritionOrder) T_OralDietNutrientModifier(numNutrient int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.OralDiet.Nutrient) >= numNutrient {
		return CodeableConceptSelect("NutritionOrder.OralDiet.Nutrient["+strconv.Itoa(numNutrient)+"].Modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionOrder.OralDiet.Nutrient["+strconv.Itoa(numNutrient)+"].Modifier", resource.OralDiet.Nutrient[numNutrient].Modifier, optionsValueSet)
}
func (resource *NutritionOrder) T_OralDietTextureId(numTexture int) templ.Component {

	if resource == nil || len(resource.OralDiet.Texture) >= numTexture {
		return StringInput("NutritionOrder.OralDiet.Texture["+strconv.Itoa(numTexture)+"].Id", nil)
	}
	return StringInput("NutritionOrder.OralDiet.Texture["+strconv.Itoa(numTexture)+"].Id", resource.OralDiet.Texture[numTexture].Id)
}
func (resource *NutritionOrder) T_OralDietTextureModifier(numTexture int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.OralDiet.Texture) >= numTexture {
		return CodeableConceptSelect("NutritionOrder.OralDiet.Texture["+strconv.Itoa(numTexture)+"].Modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionOrder.OralDiet.Texture["+strconv.Itoa(numTexture)+"].Modifier", resource.OralDiet.Texture[numTexture].Modifier, optionsValueSet)
}
func (resource *NutritionOrder) T_OralDietTextureFoodType(numTexture int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.OralDiet.Texture) >= numTexture {
		return CodeableConceptSelect("NutritionOrder.OralDiet.Texture["+strconv.Itoa(numTexture)+"].FoodType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionOrder.OralDiet.Texture["+strconv.Itoa(numTexture)+"].FoodType", resource.OralDiet.Texture[numTexture].FoodType, optionsValueSet)
}
func (resource *NutritionOrder) T_SupplementId(numSupplement int) templ.Component {

	if resource == nil || len(resource.Supplement) >= numSupplement {
		return StringInput("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].Id", nil)
	}
	return StringInput("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].Id", resource.Supplement[numSupplement].Id)
}
func (resource *NutritionOrder) T_SupplementProductName(numSupplement int) templ.Component {

	if resource == nil || len(resource.Supplement) >= numSupplement {
		return StringInput("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].ProductName", nil)
	}
	return StringInput("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].ProductName", resource.Supplement[numSupplement].ProductName)
}
func (resource *NutritionOrder) T_SupplementInstruction(numSupplement int) templ.Component {

	if resource == nil || len(resource.Supplement) >= numSupplement {
		return StringInput("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].Instruction", nil)
	}
	return StringInput("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].Instruction", resource.Supplement[numSupplement].Instruction)
}
func (resource *NutritionOrder) T_SupplementScheduleId(numSupplement int) templ.Component {

	if resource == nil || len(resource.Supplement) >= numSupplement {
		return StringInput("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].Schedule.Id", nil)
	}
	return StringInput("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].Schedule.Id", resource.Supplement[numSupplement].Schedule.Id)
}
func (resource *NutritionOrder) T_SupplementScheduleAsNeeded(numSupplement int) templ.Component {

	if resource == nil || len(resource.Supplement) >= numSupplement {
		return BoolInput("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].Schedule.AsNeeded", nil)
	}
	return BoolInput("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].Schedule.AsNeeded", resource.Supplement[numSupplement].Schedule.AsNeeded)
}
func (resource *NutritionOrder) T_SupplementScheduleAsNeededFor(numSupplement int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Supplement) >= numSupplement {
		return CodeableConceptSelect("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].Schedule.AsNeededFor", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionOrder.Supplement["+strconv.Itoa(numSupplement)+"].Schedule.AsNeededFor", resource.Supplement[numSupplement].Schedule.AsNeededFor, optionsValueSet)
}
func (resource *NutritionOrder) T_EnteralFormulaId() templ.Component {

	if resource == nil {
		return StringInput("NutritionOrder.EnteralFormula.Id", nil)
	}
	return StringInput("NutritionOrder.EnteralFormula.Id", resource.EnteralFormula.Id)
}
func (resource *NutritionOrder) T_EnteralFormulaBaseFormulaProductName() templ.Component {

	if resource == nil {
		return StringInput("NutritionOrder.EnteralFormula.BaseFormulaProductName", nil)
	}
	return StringInput("NutritionOrder.EnteralFormula.BaseFormulaProductName", resource.EnteralFormula.BaseFormulaProductName)
}
func (resource *NutritionOrder) T_EnteralFormulaRouteOfAdministration(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("NutritionOrder.EnteralFormula.RouteOfAdministration", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionOrder.EnteralFormula.RouteOfAdministration", resource.EnteralFormula.RouteOfAdministration, optionsValueSet)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationInstruction() templ.Component {

	if resource == nil {
		return StringInput("NutritionOrder.EnteralFormula.AdministrationInstruction", nil)
	}
	return StringInput("NutritionOrder.EnteralFormula.AdministrationInstruction", resource.EnteralFormula.AdministrationInstruction)
}
func (resource *NutritionOrder) T_EnteralFormulaAdditiveId(numAdditive int) templ.Component {

	if resource == nil || len(resource.EnteralFormula.Additive) >= numAdditive {
		return StringInput("NutritionOrder.EnteralFormula.Additive["+strconv.Itoa(numAdditive)+"].Id", nil)
	}
	return StringInput("NutritionOrder.EnteralFormula.Additive["+strconv.Itoa(numAdditive)+"].Id", resource.EnteralFormula.Additive[numAdditive].Id)
}
func (resource *NutritionOrder) T_EnteralFormulaAdditiveProductName(numAdditive int) templ.Component {

	if resource == nil || len(resource.EnteralFormula.Additive) >= numAdditive {
		return StringInput("NutritionOrder.EnteralFormula.Additive["+strconv.Itoa(numAdditive)+"].ProductName", nil)
	}
	return StringInput("NutritionOrder.EnteralFormula.Additive["+strconv.Itoa(numAdditive)+"].ProductName", resource.EnteralFormula.Additive[numAdditive].ProductName)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationId(numAdministration int) templ.Component {

	if resource == nil || len(resource.EnteralFormula.Administration) >= numAdministration {
		return StringInput("NutritionOrder.EnteralFormula.Administration["+strconv.Itoa(numAdministration)+"].Id", nil)
	}
	return StringInput("NutritionOrder.EnteralFormula.Administration["+strconv.Itoa(numAdministration)+"].Id", resource.EnteralFormula.Administration[numAdministration].Id)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationScheduleId(numAdministration int) templ.Component {

	if resource == nil || len(resource.EnteralFormula.Administration) >= numAdministration {
		return StringInput("NutritionOrder.EnteralFormula.Administration["+strconv.Itoa(numAdministration)+"].Schedule.Id", nil)
	}
	return StringInput("NutritionOrder.EnteralFormula.Administration["+strconv.Itoa(numAdministration)+"].Schedule.Id", resource.EnteralFormula.Administration[numAdministration].Schedule.Id)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationScheduleAsNeeded(numAdministration int) templ.Component {

	if resource == nil || len(resource.EnteralFormula.Administration) >= numAdministration {
		return BoolInput("NutritionOrder.EnteralFormula.Administration["+strconv.Itoa(numAdministration)+"].Schedule.AsNeeded", nil)
	}
	return BoolInput("NutritionOrder.EnteralFormula.Administration["+strconv.Itoa(numAdministration)+"].Schedule.AsNeeded", resource.EnteralFormula.Administration[numAdministration].Schedule.AsNeeded)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationScheduleAsNeededFor(numAdministration int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.EnteralFormula.Administration) >= numAdministration {
		return CodeableConceptSelect("NutritionOrder.EnteralFormula.Administration["+strconv.Itoa(numAdministration)+"].Schedule.AsNeededFor", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionOrder.EnteralFormula.Administration["+strconv.Itoa(numAdministration)+"].Schedule.AsNeededFor", resource.EnteralFormula.Administration[numAdministration].Schedule.AsNeededFor, optionsValueSet)
}
