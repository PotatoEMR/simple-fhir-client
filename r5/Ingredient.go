//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/Ingredient
type Ingredient struct {
	Id                  *string                  `json:"id,omitempty"`
	Meta                *Meta                    `json:"meta,omitempty"`
	ImplicitRules       *string                  `json:"implicitRules,omitempty"`
	Language            *string                  `json:"language,omitempty"`
	Text                *Narrative               `json:"text,omitempty"`
	Contained           []Resource               `json:"contained,omitempty"`
	Extension           []Extension              `json:"extension,omitempty"`
	ModifierExtension   []Extension              `json:"modifierExtension,omitempty"`
	Identifier          *Identifier              `json:"identifier,omitempty"`
	Status              string                   `json:"status"`
	For                 []Reference              `json:"for,omitempty"`
	Role                CodeableConcept          `json:"role"`
	Function            []CodeableConcept        `json:"function,omitempty"`
	Group               *CodeableConcept         `json:"group,omitempty"`
	AllergenicIndicator *bool                    `json:"allergenicIndicator,omitempty"`
	Comment             *string                  `json:"comment,omitempty"`
	Manufacturer        []IngredientManufacturer `json:"manufacturer,omitempty"`
	Substance           IngredientSubstance      `json:"substance"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Ingredient
type IngredientManufacturer struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Role              *string     `json:"role,omitempty"`
	Manufacturer      Reference   `json:"manufacturer"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Ingredient
type IngredientSubstance struct {
	Id                *string                       `json:"id,omitempty"`
	Extension         []Extension                   `json:"extension,omitempty"`
	ModifierExtension []Extension                   `json:"modifierExtension,omitempty"`
	Code              CodeableReference             `json:"code"`
	Strength          []IngredientSubstanceStrength `json:"strength,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Ingredient
type IngredientSubstanceStrength struct {
	Id                           *string                                        `json:"id,omitempty"`
	Extension                    []Extension                                    `json:"extension,omitempty"`
	ModifierExtension            []Extension                                    `json:"modifierExtension,omitempty"`
	PresentationRatio            *Ratio                                         `json:"presentationRatio"`
	PresentationRatioRange       *RatioRange                                    `json:"presentationRatioRange"`
	PresentationCodeableConcept  *CodeableConcept                               `json:"presentationCodeableConcept"`
	PresentationQuantity         *Quantity                                      `json:"presentationQuantity"`
	TextPresentation             *string                                        `json:"textPresentation,omitempty"`
	ConcentrationRatio           *Ratio                                         `json:"concentrationRatio"`
	ConcentrationRatioRange      *RatioRange                                    `json:"concentrationRatioRange"`
	ConcentrationCodeableConcept *CodeableConcept                               `json:"concentrationCodeableConcept"`
	ConcentrationQuantity        *Quantity                                      `json:"concentrationQuantity"`
	TextConcentration            *string                                        `json:"textConcentration,omitempty"`
	Basis                        *CodeableConcept                               `json:"basis,omitempty"`
	MeasurementPoint             *string                                        `json:"measurementPoint,omitempty"`
	Country                      []CodeableConcept                              `json:"country,omitempty"`
	ReferenceStrength            []IngredientSubstanceStrengthReferenceStrength `json:"referenceStrength,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Ingredient
type IngredientSubstanceStrengthReferenceStrength struct {
	Id                 *string           `json:"id,omitempty"`
	Extension          []Extension       `json:"extension,omitempty"`
	ModifierExtension  []Extension       `json:"modifierExtension,omitempty"`
	Substance          CodeableReference `json:"substance"`
	StrengthRatio      Ratio             `json:"strengthRatio"`
	StrengthRatioRange RatioRange        `json:"strengthRatioRange"`
	StrengthQuantity   Quantity          `json:"strengthQuantity"`
	MeasurementPoint   *string           `json:"measurementPoint,omitempty"`
	Country            []CodeableConcept `json:"country,omitempty"`
}

type OtherIngredient Ingredient

// on convert struct to json, automatically add resourceType=Ingredient
func (r Ingredient) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherIngredient
		ResourceType string `json:"resourceType"`
	}{
		OtherIngredient: OtherIngredient(r),
		ResourceType:    "Ingredient",
	})
}
