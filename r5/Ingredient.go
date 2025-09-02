package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *Ingredient) IngredientLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Ingredient) IngredientStatus() templ.Component {
	optionsValueSet := VSPublication_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Ingredient) IngredientRole(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", &resource.Role, optionsValueSet)
}
func (resource *Ingredient) IngredientFunction(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("function", &resource.Function[0], optionsValueSet)
}
func (resource *Ingredient) IngredientGroup(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("group", nil, optionsValueSet)
	}
	return CodeableConceptSelect("group", resource.Group, optionsValueSet)
}
func (resource *Ingredient) IngredientManufacturerRole(numManufacturer int) templ.Component {
	optionsValueSet := VSIngredient_manufacturer_role

	if resource != nil && len(resource.Manufacturer) >= numManufacturer {
		return CodeSelect("role", nil, optionsValueSet)
	}
	return CodeSelect("role", resource.Manufacturer[numManufacturer].Role, optionsValueSet)
}
func (resource *Ingredient) IngredientSubstanceStrengthBasis(numStrength int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Substance.Strength) >= numStrength {
		return CodeableConceptSelect("basis", nil, optionsValueSet)
	}
	return CodeableConceptSelect("basis", resource.Substance.Strength[numStrength].Basis, optionsValueSet)
}
func (resource *Ingredient) IngredientSubstanceStrengthCountry(numStrength int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Substance.Strength) >= numStrength {
		return CodeableConceptSelect("country", nil, optionsValueSet)
	}
	return CodeableConceptSelect("country", &resource.Substance.Strength[numStrength].Country[0], optionsValueSet)
}
func (resource *Ingredient) IngredientSubstanceStrengthReferenceStrengthCountry(numStrength int, numReferenceStrength int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Substance.Strength[numStrength].ReferenceStrength) >= numReferenceStrength {
		return CodeableConceptSelect("country", nil, optionsValueSet)
	}
	return CodeableConceptSelect("country", &resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].Country[0], optionsValueSet)
}
