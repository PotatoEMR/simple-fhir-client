package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/NutritionOrder
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
	DateTime               FhirDateTime                  `json:"dateTime"`
	Orderer                *Reference                    `json:"orderer,omitempty"`
	AllergyIntolerance     []Reference                   `json:"allergyIntolerance,omitempty"`
	FoodPreferenceModifier []CodeableConcept             `json:"foodPreferenceModifier,omitempty"`
	ExcludeFoodModifier    []CodeableConcept             `json:"excludeFoodModifier,omitempty"`
	OralDiet               *NutritionOrderOralDiet       `json:"oralDiet,omitempty"`
	Supplement             []NutritionOrderSupplement    `json:"supplement,omitempty"`
	EnteralFormula         *NutritionOrderEnteralFormula `json:"enteralFormula,omitempty"`
	Note                   []Annotation                  `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/NutritionOrder
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

// http://hl7.org/fhir/r4b/StructureDefinition/NutritionOrder
type NutritionOrderOralDietNutrient struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Modifier          *CodeableConcept `json:"modifier,omitempty"`
	Amount            *Quantity        `json:"amount,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/NutritionOrder
type NutritionOrderOralDietTexture struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Modifier          *CodeableConcept `json:"modifier,omitempty"`
	FoodType          *CodeableConcept `json:"foodType,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/NutritionOrder
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

// http://hl7.org/fhir/r4b/StructureDefinition/NutritionOrder
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

// http://hl7.org/fhir/r4b/StructureDefinition/NutritionOrder
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
func (resource *NutritionOrder) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *NutritionOrder) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *NutritionOrder) T_Instantiates(numInstantiates int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiates >= len(resource.Instantiates) {
		return StringInput("instantiates["+strconv.Itoa(numInstantiates)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiates["+strconv.Itoa(numInstantiates)+"]", &resource.Instantiates[numInstantiates], htmlAttrs)
}
func (resource *NutritionOrder) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_Intent(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_Patient(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("patient", nil, htmlAttrs)
	}
	return ReferenceInput("patient", &resource.Patient, htmlAttrs)
}
func (resource *NutritionOrder) T_Encounter(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("encounter", nil, htmlAttrs)
	}
	return ReferenceInput("encounter", resource.Encounter, htmlAttrs)
}
func (resource *NutritionOrder) T_DateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("dateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("dateTime", &resource.DateTime, htmlAttrs)
}
func (resource *NutritionOrder) T_Orderer(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("orderer", nil, htmlAttrs)
	}
	return ReferenceInput("orderer", resource.Orderer, htmlAttrs)
}
func (resource *NutritionOrder) T_AllergyIntolerance(numAllergyIntolerance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllergyIntolerance >= len(resource.AllergyIntolerance) {
		return ReferenceInput("allergyIntolerance["+strconv.Itoa(numAllergyIntolerance)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("allergyIntolerance["+strconv.Itoa(numAllergyIntolerance)+"]", &resource.AllergyIntolerance[numAllergyIntolerance], htmlAttrs)
}
func (resource *NutritionOrder) T_FoodPreferenceModifier(numFoodPreferenceModifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFoodPreferenceModifier >= len(resource.FoodPreferenceModifier) {
		return CodeableConceptSelect("foodPreferenceModifier["+strconv.Itoa(numFoodPreferenceModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("foodPreferenceModifier["+strconv.Itoa(numFoodPreferenceModifier)+"]", &resource.FoodPreferenceModifier[numFoodPreferenceModifier], optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_ExcludeFoodModifier(numExcludeFoodModifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numExcludeFoodModifier >= len(resource.ExcludeFoodModifier) {
		return CodeableConceptSelect("excludeFoodModifier["+strconv.Itoa(numExcludeFoodModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("excludeFoodModifier["+strconv.Itoa(numExcludeFoodModifier)+"]", &resource.ExcludeFoodModifier[numExcludeFoodModifier], optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietType(numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numType >= len(resource.OralDiet.Type) {
		return CodeableConceptSelect("oralDiet.type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("oralDiet.type["+strconv.Itoa(numType)+"]", &resource.OralDiet.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietSchedule(numSchedule int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSchedule >= len(resource.OralDiet.Schedule) {
		return TimingInput("oralDiet.schedule["+strconv.Itoa(numSchedule)+"]", nil, htmlAttrs)
	}
	return TimingInput("oralDiet.schedule["+strconv.Itoa(numSchedule)+"]", &resource.OralDiet.Schedule[numSchedule], htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietFluidConsistencyType(numFluidConsistencyType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFluidConsistencyType >= len(resource.OralDiet.FluidConsistencyType) {
		return CodeableConceptSelect("oralDiet.fluidConsistencyType["+strconv.Itoa(numFluidConsistencyType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("oralDiet.fluidConsistencyType["+strconv.Itoa(numFluidConsistencyType)+"]", &resource.OralDiet.FluidConsistencyType[numFluidConsistencyType], optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietInstruction(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("oralDiet.instruction", nil, htmlAttrs)
	}
	return StringInput("oralDiet.instruction", resource.OralDiet.Instruction, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietNutrientModifier(numNutrient int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNutrient >= len(resource.OralDiet.Nutrient) {
		return CodeableConceptSelect("oralDiet.nutrient["+strconv.Itoa(numNutrient)+"].modifier", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("oralDiet.nutrient["+strconv.Itoa(numNutrient)+"].modifier", resource.OralDiet.Nutrient[numNutrient].Modifier, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietNutrientAmount(numNutrient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNutrient >= len(resource.OralDiet.Nutrient) {
		return QuantityInput("oralDiet.nutrient["+strconv.Itoa(numNutrient)+"].amount", nil, htmlAttrs)
	}
	return QuantityInput("oralDiet.nutrient["+strconv.Itoa(numNutrient)+"].amount", resource.OralDiet.Nutrient[numNutrient].Amount, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietTextureModifier(numTexture int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTexture >= len(resource.OralDiet.Texture) {
		return CodeableConceptSelect("oralDiet.texture["+strconv.Itoa(numTexture)+"].modifier", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("oralDiet.texture["+strconv.Itoa(numTexture)+"].modifier", resource.OralDiet.Texture[numTexture].Modifier, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietTextureFoodType(numTexture int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTexture >= len(resource.OralDiet.Texture) {
		return CodeableConceptSelect("oralDiet.texture["+strconv.Itoa(numTexture)+"].foodType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("oralDiet.texture["+strconv.Itoa(numTexture)+"].foodType", resource.OralDiet.Texture[numTexture].FoodType, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementType(numSupplement int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return CodeableConceptSelect("supplement["+strconv.Itoa(numSupplement)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("supplement["+strconv.Itoa(numSupplement)+"].type", resource.Supplement[numSupplement].Type, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementProductName(numSupplement int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return StringInput("supplement["+strconv.Itoa(numSupplement)+"].productName", nil, htmlAttrs)
	}
	return StringInput("supplement["+strconv.Itoa(numSupplement)+"].productName", resource.Supplement[numSupplement].ProductName, htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementSchedule(numSupplement int, numSchedule int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) || numSchedule >= len(resource.Supplement[numSupplement].Schedule) {
		return TimingInput("supplement["+strconv.Itoa(numSupplement)+"].schedule["+strconv.Itoa(numSchedule)+"]", nil, htmlAttrs)
	}
	return TimingInput("supplement["+strconv.Itoa(numSupplement)+"].schedule["+strconv.Itoa(numSchedule)+"]", &resource.Supplement[numSupplement].Schedule[numSchedule], htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementQuantity(numSupplement int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return QuantityInput("supplement["+strconv.Itoa(numSupplement)+"].quantity", nil, htmlAttrs)
	}
	return QuantityInput("supplement["+strconv.Itoa(numSupplement)+"].quantity", resource.Supplement[numSupplement].Quantity, htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementInstruction(numSupplement int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return StringInput("supplement["+strconv.Itoa(numSupplement)+"].instruction", nil, htmlAttrs)
	}
	return StringInput("supplement["+strconv.Itoa(numSupplement)+"].instruction", resource.Supplement[numSupplement].Instruction, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaBaseFormulaType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("enteralFormula.baseFormulaType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("enteralFormula.baseFormulaType", resource.EnteralFormula.BaseFormulaType, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaBaseFormulaProductName(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("enteralFormula.baseFormulaProductName", nil, htmlAttrs)
	}
	return StringInput("enteralFormula.baseFormulaProductName", resource.EnteralFormula.BaseFormulaProductName, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdditiveType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("enteralFormula.additiveType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("enteralFormula.additiveType", resource.EnteralFormula.AdditiveType, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdditiveProductName(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("enteralFormula.additiveProductName", nil, htmlAttrs)
	}
	return StringInput("enteralFormula.additiveProductName", resource.EnteralFormula.AdditiveProductName, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaCaloricDensity(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return QuantityInput("enteralFormula.caloricDensity", nil, htmlAttrs)
	}
	return QuantityInput("enteralFormula.caloricDensity", resource.EnteralFormula.CaloricDensity, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaRouteofAdministration(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("enteralFormula.routeofAdministration", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("enteralFormula.routeofAdministration", resource.EnteralFormula.RouteofAdministration, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaMaxVolumeToDeliver(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return QuantityInput("enteralFormula.maxVolumeToDeliver", nil, htmlAttrs)
	}
	return QuantityInput("enteralFormula.maxVolumeToDeliver", resource.EnteralFormula.MaxVolumeToDeliver, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationInstruction(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("enteralFormula.administrationInstruction", nil, htmlAttrs)
	}
	return StringInput("enteralFormula.administrationInstruction", resource.EnteralFormula.AdministrationInstruction, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationSchedule(numAdministration int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdministration >= len(resource.EnteralFormula.Administration) {
		return TimingInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].schedule", nil, htmlAttrs)
	}
	return TimingInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].schedule", resource.EnteralFormula.Administration[numAdministration].Schedule, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationQuantity(numAdministration int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdministration >= len(resource.EnteralFormula.Administration) {
		return QuantityInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].quantity", nil, htmlAttrs)
	}
	return QuantityInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].quantity", resource.EnteralFormula.Administration[numAdministration].Quantity, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationRateQuantity(numAdministration int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdministration >= len(resource.EnteralFormula.Administration) {
		return QuantityInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].rateQuantity", nil, htmlAttrs)
	}
	return QuantityInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].rateQuantity", resource.EnteralFormula.Administration[numAdministration].RateQuantity, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationRateRatio(numAdministration int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdministration >= len(resource.EnteralFormula.Administration) {
		return RatioInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].rateRatio", nil, htmlAttrs)
	}
	return RatioInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].rateRatio", resource.EnteralFormula.Administration[numAdministration].RateRatio, htmlAttrs)
}
