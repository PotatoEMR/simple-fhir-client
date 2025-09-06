package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/NutritionIntake
type NutritionIntake struct {
	Id                    *string                          `json:"id,omitempty"`
	Meta                  *Meta                            `json:"meta,omitempty"`
	ImplicitRules         *string                          `json:"implicitRules,omitempty"`
	Language              *string                          `json:"language,omitempty"`
	Text                  *Narrative                       `json:"text,omitempty"`
	Contained             []Resource                       `json:"contained,omitempty"`
	Extension             []Extension                      `json:"extension,omitempty"`
	ModifierExtension     []Extension                      `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                     `json:"identifier,omitempty"`
	InstantiatesCanonical []string                         `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                         `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference                      `json:"basedOn,omitempty"`
	PartOf                []Reference                      `json:"partOf,omitempty"`
	Status                string                           `json:"status"`
	StatusReason          []CodeableConcept                `json:"statusReason,omitempty"`
	Code                  *CodeableConcept                 `json:"code,omitempty"`
	Subject               Reference                        `json:"subject"`
	Encounter             *Reference                       `json:"encounter,omitempty"`
	OccurrenceDateTime    *string                          `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *Period                          `json:"occurrencePeriod,omitempty"`
	Recorded              *string                          `json:"recorded,omitempty"`
	ReportedBoolean       *bool                            `json:"reportedBoolean,omitempty"`
	ReportedReference     *Reference                       `json:"reportedReference,omitempty"`
	ConsumedItem          []NutritionIntakeConsumedItem    `json:"consumedItem"`
	IngredientLabel       []NutritionIntakeIngredientLabel `json:"ingredientLabel,omitempty"`
	Performer             []NutritionIntakePerformer       `json:"performer,omitempty"`
	Location              *Reference                       `json:"location,omitempty"`
	DerivedFrom           []Reference                      `json:"derivedFrom,omitempty"`
	Reason                []CodeableReference              `json:"reason,omitempty"`
	Note                  []Annotation                     `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionIntake
type NutritionIntakeConsumedItem struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	NutritionProduct  CodeableReference `json:"nutritionProduct"`
	Schedule          *Timing           `json:"schedule,omitempty"`
	Amount            *Quantity         `json:"amount,omitempty"`
	Rate              *Quantity         `json:"rate,omitempty"`
	NotConsumed       *bool             `json:"notConsumed,omitempty"`
	NotConsumedReason *CodeableConcept  `json:"notConsumedReason,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionIntake
type NutritionIntakeIngredientLabel struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Nutrient          CodeableReference `json:"nutrient"`
	Amount            Quantity          `json:"amount"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NutritionIntake
type NutritionIntakePerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

type OtherNutritionIntake NutritionIntake

// on convert struct to json, automatically add resourceType=NutritionIntake
func (r NutritionIntake) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNutritionIntake
		ResourceType string `json:"resourceType"`
	}{
		OtherNutritionIntake: OtherNutritionIntake(r),
		ResourceType:         "NutritionIntake",
	})
}

func (resource *NutritionIntake) T_Id() templ.Component {

	if resource == nil {
		return StringInput("NutritionIntake.Id", nil)
	}
	return StringInput("NutritionIntake.Id", resource.Id)
}
func (resource *NutritionIntake) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("NutritionIntake.ImplicitRules", nil)
	}
	return StringInput("NutritionIntake.ImplicitRules", resource.ImplicitRules)
}
func (resource *NutritionIntake) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("NutritionIntake.Language", nil, optionsValueSet)
	}
	return CodeSelect("NutritionIntake.Language", resource.Language, optionsValueSet)
}
func (resource *NutritionIntake) T_InstantiatesCanonical(numInstantiatesCanonical int) templ.Component {

	if resource == nil || len(resource.InstantiatesCanonical) >= numInstantiatesCanonical {
		return StringInput("NutritionIntake.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil)
	}
	return StringInput("NutritionIntake.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical])
}
func (resource *NutritionIntake) T_InstantiatesUri(numInstantiatesUri int) templ.Component {

	if resource == nil || len(resource.InstantiatesUri) >= numInstantiatesUri {
		return StringInput("NutritionIntake.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil)
	}
	return StringInput("NutritionIntake.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri])
}
func (resource *NutritionIntake) T_Status() templ.Component {
	optionsValueSet := VSEvent_status

	if resource == nil {
		return CodeSelect("NutritionIntake.Status", nil, optionsValueSet)
	}
	return CodeSelect("NutritionIntake.Status", &resource.Status, optionsValueSet)
}
func (resource *NutritionIntake) T_StatusReason(numStatusReason int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.StatusReason) >= numStatusReason {
		return CodeableConceptSelect("NutritionIntake.StatusReason["+strconv.Itoa(numStatusReason)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionIntake.StatusReason["+strconv.Itoa(numStatusReason)+"]", &resource.StatusReason[numStatusReason], optionsValueSet)
}
func (resource *NutritionIntake) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("NutritionIntake.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionIntake.Code", resource.Code, optionsValueSet)
}
func (resource *NutritionIntake) T_Recorded() templ.Component {

	if resource == nil {
		return StringInput("NutritionIntake.Recorded", nil)
	}
	return StringInput("NutritionIntake.Recorded", resource.Recorded)
}
func (resource *NutritionIntake) T_ConsumedItemId(numConsumedItem int) templ.Component {

	if resource == nil || len(resource.ConsumedItem) >= numConsumedItem {
		return StringInput("NutritionIntake.ConsumedItem["+strconv.Itoa(numConsumedItem)+"].Id", nil)
	}
	return StringInput("NutritionIntake.ConsumedItem["+strconv.Itoa(numConsumedItem)+"].Id", resource.ConsumedItem[numConsumedItem].Id)
}
func (resource *NutritionIntake) T_ConsumedItemType(numConsumedItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ConsumedItem) >= numConsumedItem {
		return CodeableConceptSelect("NutritionIntake.ConsumedItem["+strconv.Itoa(numConsumedItem)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionIntake.ConsumedItem["+strconv.Itoa(numConsumedItem)+"].Type", &resource.ConsumedItem[numConsumedItem].Type, optionsValueSet)
}
func (resource *NutritionIntake) T_ConsumedItemNotConsumed(numConsumedItem int) templ.Component {

	if resource == nil || len(resource.ConsumedItem) >= numConsumedItem {
		return BoolInput("NutritionIntake.ConsumedItem["+strconv.Itoa(numConsumedItem)+"].NotConsumed", nil)
	}
	return BoolInput("NutritionIntake.ConsumedItem["+strconv.Itoa(numConsumedItem)+"].NotConsumed", resource.ConsumedItem[numConsumedItem].NotConsumed)
}
func (resource *NutritionIntake) T_ConsumedItemNotConsumedReason(numConsumedItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ConsumedItem) >= numConsumedItem {
		return CodeableConceptSelect("NutritionIntake.ConsumedItem["+strconv.Itoa(numConsumedItem)+"].NotConsumedReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionIntake.ConsumedItem["+strconv.Itoa(numConsumedItem)+"].NotConsumedReason", resource.ConsumedItem[numConsumedItem].NotConsumedReason, optionsValueSet)
}
func (resource *NutritionIntake) T_IngredientLabelId(numIngredientLabel int) templ.Component {

	if resource == nil || len(resource.IngredientLabel) >= numIngredientLabel {
		return StringInput("NutritionIntake.IngredientLabel["+strconv.Itoa(numIngredientLabel)+"].Id", nil)
	}
	return StringInput("NutritionIntake.IngredientLabel["+strconv.Itoa(numIngredientLabel)+"].Id", resource.IngredientLabel[numIngredientLabel].Id)
}
func (resource *NutritionIntake) T_PerformerId(numPerformer int) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return StringInput("NutritionIntake.Performer["+strconv.Itoa(numPerformer)+"].Id", nil)
	}
	return StringInput("NutritionIntake.Performer["+strconv.Itoa(numPerformer)+"].Id", resource.Performer[numPerformer].Id)
}
func (resource *NutritionIntake) T_PerformerFunction(numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return CodeableConceptSelect("NutritionIntake.Performer["+strconv.Itoa(numPerformer)+"].Function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NutritionIntake.Performer["+strconv.Itoa(numPerformer)+"].Function", resource.Performer[numPerformer].Function, optionsValueSet)
}
