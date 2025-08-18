//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

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
	DateTime               string                        `json:"dateTime"`
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
	RateQuantity      *Quantity   `json:"rateQuantity"`
	RateRatio         *Ratio      `json:"rateRatio"`
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
