package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductIngredient
type MedicinalProductIngredient struct {
	Id                  *string                                        `json:"id,omitempty"`
	Meta                *Meta                                          `json:"meta,omitempty"`
	ImplicitRules       *string                                        `json:"implicitRules,omitempty"`
	Language            *string                                        `json:"language,omitempty"`
	Text                *Narrative                                     `json:"text,omitempty"`
	Contained           []Resource                                     `json:"contained,omitempty"`
	Extension           []Extension                                    `json:"extension,omitempty"`
	ModifierExtension   []Extension                                    `json:"modifierExtension,omitempty"`
	Identifier          *Identifier                                    `json:"identifier,omitempty"`
	Role                CodeableConcept                                `json:"role"`
	AllergenicIndicator *bool                                          `json:"allergenicIndicator,omitempty"`
	Manufacturer        []Reference                                    `json:"manufacturer,omitempty"`
	SpecifiedSubstance  []MedicinalProductIngredientSpecifiedSubstance `json:"specifiedSubstance,omitempty"`
	Substance           *MedicinalProductIngredientSubstance           `json:"substance,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductIngredient
type MedicinalProductIngredientSpecifiedSubstance struct {
	Id                *string                                                `json:"id,omitempty"`
	Extension         []Extension                                            `json:"extension,omitempty"`
	ModifierExtension []Extension                                            `json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                        `json:"code"`
	Group             CodeableConcept                                        `json:"group"`
	Confidentiality   *CodeableConcept                                       `json:"confidentiality,omitempty"`
	Strength          []MedicinalProductIngredientSpecifiedSubstanceStrength `json:"strength,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductIngredient
type MedicinalProductIngredientSpecifiedSubstanceStrength struct {
	Id                    *string                                                                 `json:"id,omitempty"`
	Extension             []Extension                                                             `json:"extension,omitempty"`
	ModifierExtension     []Extension                                                             `json:"modifierExtension,omitempty"`
	Presentation          Ratio                                                                   `json:"presentation"`
	PresentationLowLimit  *Ratio                                                                  `json:"presentationLowLimit,omitempty"`
	Concentration         *Ratio                                                                  `json:"concentration,omitempty"`
	ConcentrationLowLimit *Ratio                                                                  `json:"concentrationLowLimit,omitempty"`
	MeasurementPoint      *string                                                                 `json:"measurementPoint,omitempty"`
	Country               []CodeableConcept                                                       `json:"country,omitempty"`
	ReferenceStrength     []MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength `json:"referenceStrength,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductIngredient
type MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Substance         *CodeableConcept  `json:"substance,omitempty"`
	Strength          Ratio             `json:"strength"`
	StrengthLowLimit  *Ratio            `json:"strengthLowLimit,omitempty"`
	MeasurementPoint  *string           `json:"measurementPoint,omitempty"`
	Country           []CodeableConcept `json:"country,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductIngredient
type MedicinalProductIngredientSubstance struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code"`
}

type OtherMedicinalProductIngredient MedicinalProductIngredient

// on convert struct to json, automatically add resourceType=MedicinalProductIngredient
func (r MedicinalProductIngredient) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductIngredient
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductIngredient: OtherMedicinalProductIngredient(r),
		ResourceType:                    "MedicinalProductIngredient",
	})
}

func (resource *MedicinalProductIngredient) MedicinalProductIngredientLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
