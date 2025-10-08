package r5

//generated with command go run ./bultaoreune
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
	DateTime               FhirDateTime                  `json:"dateTime"`
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
func (resource *NutritionOrder) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *NutritionOrder) T_GroupIdentifier(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IdentifierInput("groupIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("groupIdentifier", resource.GroupIdentifier, htmlAttrs)
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
func (resource *NutritionOrder) T_Priority(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *NutritionOrder) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *NutritionOrder) T_SupportingInformation(frs []FhirResource, numSupportingInformation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInformation >= len(resource.SupportingInformation) {
		return ReferenceInput(frs, "supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", &resource.SupportingInformation[numSupportingInformation], htmlAttrs)
}
func (resource *NutritionOrder) T_DateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("dateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("dateTime", &resource.DateTime, htmlAttrs)
}
func (resource *NutritionOrder) T_Orderer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "orderer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "orderer", resource.Orderer, htmlAttrs)
}
func (resource *NutritionOrder) T_Performer(numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableReferenceInput("performer["+strconv.Itoa(numPerformer)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("performer["+strconv.Itoa(numPerformer)+"]", &resource.Performer[numPerformer], htmlAttrs)
}
func (resource *NutritionOrder) T_AllergyIntolerance(frs []FhirResource, numAllergyIntolerance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllergyIntolerance >= len(resource.AllergyIntolerance) {
		return ReferenceInput(frs, "allergyIntolerance["+strconv.Itoa(numAllergyIntolerance)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "allergyIntolerance["+strconv.Itoa(numAllergyIntolerance)+"]", &resource.AllergyIntolerance[numAllergyIntolerance], htmlAttrs)
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
func (resource *NutritionOrder) T_OutsideFoodAllowed(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("outsideFoodAllowed", nil, htmlAttrs)
	}
	return BoolInput("outsideFoodAllowed", resource.OutsideFoodAllowed, htmlAttrs)
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
func (resource *NutritionOrder) T_OralDietScheduleTiming(numTiming int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTiming >= len(resource.OralDiet.Schedule.Timing) {
		return TimingInput("oralDiet.schedule.timing["+strconv.Itoa(numTiming)+"]", nil, htmlAttrs)
	}
	return TimingInput("oralDiet.schedule.timing["+strconv.Itoa(numTiming)+"]", &resource.OralDiet.Schedule.Timing[numTiming], htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietScheduleAsNeeded(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("oralDiet.schedule.asNeeded", nil, htmlAttrs)
	}
	return BoolInput("oralDiet.schedule.asNeeded", resource.OralDiet.Schedule.AsNeeded, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietScheduleAsNeededFor(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("oralDiet.schedule.asNeededFor", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("oralDiet.schedule.asNeededFor", resource.OralDiet.Schedule.AsNeededFor, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietNutrientModifier(numNutrient int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNutrient >= len(resource.OralDiet.Nutrient) {
		return CodeableConceptSelect("oralDiet.nutrient["+strconv.Itoa(numNutrient)+"].modifier", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("oralDiet.nutrient["+strconv.Itoa(numNutrient)+"].modifier", resource.OralDiet.Nutrient[numNutrient].Modifier, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_OralDietNutrientAmount(numNutrient int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numNutrient >= len(resource.OralDiet.Nutrient) {
		return QuantityInput("oralDiet.nutrient["+strconv.Itoa(numNutrient)+"].amount", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("oralDiet.nutrient["+strconv.Itoa(numNutrient)+"].amount", resource.OralDiet.Nutrient[numNutrient].Amount, optionsValueSet, htmlAttrs)
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
func (resource *NutritionOrder) T_SupplementType(numSupplement int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return CodeableReferenceInput("supplement["+strconv.Itoa(numSupplement)+"].type", nil, htmlAttrs)
	}
	return CodeableReferenceInput("supplement["+strconv.Itoa(numSupplement)+"].type", resource.Supplement[numSupplement].Type, htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementProductName(numSupplement int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return StringInput("supplement["+strconv.Itoa(numSupplement)+"].productName", nil, htmlAttrs)
	}
	return StringInput("supplement["+strconv.Itoa(numSupplement)+"].productName", resource.Supplement[numSupplement].ProductName, htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementQuantity(numSupplement int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return QuantityInput("supplement["+strconv.Itoa(numSupplement)+"].quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("supplement["+strconv.Itoa(numSupplement)+"].quantity", resource.Supplement[numSupplement].Quantity, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementInstruction(numSupplement int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return StringInput("supplement["+strconv.Itoa(numSupplement)+"].instruction", nil, htmlAttrs)
	}
	return StringInput("supplement["+strconv.Itoa(numSupplement)+"].instruction", resource.Supplement[numSupplement].Instruction, htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementScheduleTiming(numSupplement int, numTiming int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) || numTiming >= len(resource.Supplement[numSupplement].Schedule.Timing) {
		return TimingInput("supplement["+strconv.Itoa(numSupplement)+"].schedule.timing["+strconv.Itoa(numTiming)+"]", nil, htmlAttrs)
	}
	return TimingInput("supplement["+strconv.Itoa(numSupplement)+"].schedule.timing["+strconv.Itoa(numTiming)+"]", &resource.Supplement[numSupplement].Schedule.Timing[numTiming], htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementScheduleAsNeeded(numSupplement int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return BoolInput("supplement["+strconv.Itoa(numSupplement)+"].schedule.asNeeded", nil, htmlAttrs)
	}
	return BoolInput("supplement["+strconv.Itoa(numSupplement)+"].schedule.asNeeded", resource.Supplement[numSupplement].Schedule.AsNeeded, htmlAttrs)
}
func (resource *NutritionOrder) T_SupplementScheduleAsNeededFor(numSupplement int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupplement >= len(resource.Supplement) {
		return CodeableConceptSelect("supplement["+strconv.Itoa(numSupplement)+"].schedule.asNeededFor", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("supplement["+strconv.Itoa(numSupplement)+"].schedule.asNeededFor", resource.Supplement[numSupplement].Schedule.AsNeededFor, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaBaseFormulaType(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("enteralFormula.baseFormulaType", nil, htmlAttrs)
	}
	return CodeableReferenceInput("enteralFormula.baseFormulaType", resource.EnteralFormula.BaseFormulaType, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaBaseFormulaProductName(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("enteralFormula.baseFormulaProductName", nil, htmlAttrs)
	}
	return StringInput("enteralFormula.baseFormulaProductName", resource.EnteralFormula.BaseFormulaProductName, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaDeliveryDevice(numDeliveryDevice int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDeliveryDevice >= len(resource.EnteralFormula.DeliveryDevice) {
		return CodeableReferenceInput("enteralFormula.deliveryDevice["+strconv.Itoa(numDeliveryDevice)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("enteralFormula.deliveryDevice["+strconv.Itoa(numDeliveryDevice)+"]", &resource.EnteralFormula.DeliveryDevice[numDeliveryDevice], htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaCaloricDensity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("enteralFormula.caloricDensity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("enteralFormula.caloricDensity", resource.EnteralFormula.CaloricDensity, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaRouteOfAdministration(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("enteralFormula.routeOfAdministration", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("enteralFormula.routeOfAdministration", resource.EnteralFormula.RouteOfAdministration, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaMaxVolumeToDeliver(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("enteralFormula.maxVolumeToDeliver", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("enteralFormula.maxVolumeToDeliver", resource.EnteralFormula.MaxVolumeToDeliver, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationInstruction(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("enteralFormula.administrationInstruction", nil, htmlAttrs)
	}
	return StringInput("enteralFormula.administrationInstruction", resource.EnteralFormula.AdministrationInstruction, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdditiveType(numAdditive int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdditive >= len(resource.EnteralFormula.Additive) {
		return CodeableReferenceInput("enteralFormula.additive["+strconv.Itoa(numAdditive)+"].type", nil, htmlAttrs)
	}
	return CodeableReferenceInput("enteralFormula.additive["+strconv.Itoa(numAdditive)+"].type", resource.EnteralFormula.Additive[numAdditive].Type, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdditiveProductName(numAdditive int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdditive >= len(resource.EnteralFormula.Additive) {
		return StringInput("enteralFormula.additive["+strconv.Itoa(numAdditive)+"].productName", nil, htmlAttrs)
	}
	return StringInput("enteralFormula.additive["+strconv.Itoa(numAdditive)+"].productName", resource.EnteralFormula.Additive[numAdditive].ProductName, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdditiveQuantity(numAdditive int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numAdditive >= len(resource.EnteralFormula.Additive) {
		return QuantityInput("enteralFormula.additive["+strconv.Itoa(numAdditive)+"].quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("enteralFormula.additive["+strconv.Itoa(numAdditive)+"].quantity", resource.EnteralFormula.Additive[numAdditive].Quantity, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationQuantity(numAdministration int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numAdministration >= len(resource.EnteralFormula.Administration) {
		return QuantityInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].quantity", resource.EnteralFormula.Administration[numAdministration].Quantity, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationRateQuantity(numAdministration int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numAdministration >= len(resource.EnteralFormula.Administration) {
		return QuantityInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].rateQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].rateQuantity", resource.EnteralFormula.Administration[numAdministration].RateQuantity, optionsValueSet, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationRateRatio(numAdministration int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdministration >= len(resource.EnteralFormula.Administration) {
		return RatioInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].rateRatio", nil, htmlAttrs)
	}
	return RatioInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].rateRatio", resource.EnteralFormula.Administration[numAdministration].RateRatio, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationScheduleTiming(numAdministration int, numTiming int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdministration >= len(resource.EnteralFormula.Administration) || numTiming >= len(resource.EnteralFormula.Administration[numAdministration].Schedule.Timing) {
		return TimingInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].schedule.timing["+strconv.Itoa(numTiming)+"]", nil, htmlAttrs)
	}
	return TimingInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].schedule.timing["+strconv.Itoa(numTiming)+"]", &resource.EnteralFormula.Administration[numAdministration].Schedule.Timing[numTiming], htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationScheduleAsNeeded(numAdministration int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdministration >= len(resource.EnteralFormula.Administration) {
		return BoolInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].schedule.asNeeded", nil, htmlAttrs)
	}
	return BoolInput("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].schedule.asNeeded", resource.EnteralFormula.Administration[numAdministration].Schedule.AsNeeded, htmlAttrs)
}
func (resource *NutritionOrder) T_EnteralFormulaAdministrationScheduleAsNeededFor(numAdministration int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdministration >= len(resource.EnteralFormula.Administration) {
		return CodeableConceptSelect("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].schedule.asNeededFor", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("enteralFormula.administration["+strconv.Itoa(numAdministration)+"].schedule.asNeededFor", resource.EnteralFormula.Administration[numAdministration].Schedule.AsNeededFor, optionsValueSet, htmlAttrs)
}
