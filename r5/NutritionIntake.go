//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

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
	OccurrenceDateTime    *string                          `json:"occurrenceDateTime"`
	OccurrencePeriod      *Period                          `json:"occurrencePeriod"`
	Recorded              *string                          `json:"recorded,omitempty"`
	ReportedBoolean       *bool                            `json:"reportedBoolean"`
	ReportedReference     *Reference                       `json:"reportedReference"`
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
