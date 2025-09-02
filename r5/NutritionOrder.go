package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	RateQuantity      *Quantity                                           `json:"rateQuantity"`
	RateRatio         *Ratio                                              `json:"rateRatio"`
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

func (resource *NutritionOrder) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *NutritionOrder) T_Status() templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *NutritionOrder) T_Intent() templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("intent", nil, optionsValueSet)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet)
}
func (resource *NutritionOrder) T_Priority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet)
}
func (resource *NutritionOrder) T_FoodPreferenceModifier(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("foodPreferenceModifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("foodPreferenceModifier", &resource.FoodPreferenceModifier[0], optionsValueSet)
}
func (resource *NutritionOrder) T_ExcludeFoodModifier(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("excludeFoodModifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("excludeFoodModifier", &resource.ExcludeFoodModifier[0], optionsValueSet)
}
func (resource *NutritionOrder) T_OralDietType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.OralDiet.Type[0], optionsValueSet)
}
func (resource *NutritionOrder) T_OralDietFluidConsistencyType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("fluidConsistencyType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("fluidConsistencyType", &resource.OralDiet.FluidConsistencyType[0], optionsValueSet)
}
func (resource *NutritionOrder) T_OralDietScheduleAsNeededFor(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("asNeededFor", nil, optionsValueSet)
	}
	return CodeableConceptSelect("asNeededFor", resource.OralDiet.Schedule.AsNeededFor, optionsValueSet)
}
func (resource *NutritionOrder) T_OralDietNutrientModifier(numNutrient int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.OralDiet.Nutrient) >= numNutrient {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", resource.OralDiet.Nutrient[numNutrient].Modifier, optionsValueSet)
}
func (resource *NutritionOrder) T_OralDietTextureModifier(numTexture int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.OralDiet.Texture) >= numTexture {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", resource.OralDiet.Texture[numTexture].Modifier, optionsValueSet)
}
func (resource *NutritionOrder) T_OralDietTextureFoodType(numTexture int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.OralDiet.Texture) >= numTexture {
		return CodeableConceptSelect("foodType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("foodType", resource.OralDiet.Texture[numTexture].FoodType, optionsValueSet)
}
func (resource *NutritionOrder) T_SupplementScheduleAsNeededFor(numSupplement int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Supplement) >= numSupplement {
		return CodeableConceptSelect("asNeededFor", nil, optionsValueSet)
	}
	return CodeableConceptSelect("asNeededFor", resource.Supplement[numSupplement].Schedule.AsNeededFor, optionsValueSet)
}
func (resource *NutritionOrder) T_EnteralFormulaRouteOfAdministration(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("routeOfAdministration", nil, optionsValueSet)
	}
	return CodeableConceptSelect("routeOfAdministration", resource.EnteralFormula.RouteOfAdministration, optionsValueSet)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationScheduleAsNeededFor(numAdministration int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.EnteralFormula.Administration) >= numAdministration {
		return CodeableConceptSelect("asNeededFor", nil, optionsValueSet)
	}
	return CodeableConceptSelect("asNeededFor", resource.EnteralFormula.Administration[numAdministration].Schedule.AsNeededFor, optionsValueSet)
}
