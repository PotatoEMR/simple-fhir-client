package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	PresentationRatio            *Ratio                                         `json:"presentationRatio,omitempty"`
	PresentationRatioRange       *RatioRange                                    `json:"presentationRatioRange,omitempty"`
	PresentationCodeableConcept  *CodeableConcept                               `json:"presentationCodeableConcept,omitempty"`
	PresentationQuantity         *Quantity                                      `json:"presentationQuantity,omitempty"`
	TextPresentation             *string                                        `json:"textPresentation,omitempty"`
	ConcentrationRatio           *Ratio                                         `json:"concentrationRatio,omitempty"`
	ConcentrationRatioRange      *RatioRange                                    `json:"concentrationRatioRange,omitempty"`
	ConcentrationCodeableConcept *CodeableConcept                               `json:"concentrationCodeableConcept,omitempty"`
	ConcentrationQuantity        *Quantity                                      `json:"concentrationQuantity,omitempty"`
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

func (resource *Ingredient) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Ingredient.Id", nil)
	}
	return StringInput("Ingredient.Id", resource.Id)
}
func (resource *Ingredient) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Ingredient.ImplicitRules", nil)
	}
	return StringInput("Ingredient.ImplicitRules", resource.ImplicitRules)
}
func (resource *Ingredient) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Ingredient.Language", nil, optionsValueSet)
	}
	return CodeSelect("Ingredient.Language", resource.Language, optionsValueSet)
}
func (resource *Ingredient) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Ingredient.Status", nil, optionsValueSet)
	}
	return CodeSelect("Ingredient.Status", &resource.Status, optionsValueSet)
}
func (resource *Ingredient) T_Role(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Ingredient.Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Ingredient.Role", &resource.Role, optionsValueSet)
}
func (resource *Ingredient) T_Function(numFunction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Function) >= numFunction {
		return CodeableConceptSelect("Ingredient.Function["+strconv.Itoa(numFunction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Ingredient.Function["+strconv.Itoa(numFunction)+"]", &resource.Function[numFunction], optionsValueSet)
}
func (resource *Ingredient) T_Group(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Ingredient.Group", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Ingredient.Group", resource.Group, optionsValueSet)
}
func (resource *Ingredient) T_AllergenicIndicator() templ.Component {

	if resource == nil {
		return BoolInput("Ingredient.AllergenicIndicator", nil)
	}
	return BoolInput("Ingredient.AllergenicIndicator", resource.AllergenicIndicator)
}
func (resource *Ingredient) T_Comment() templ.Component {

	if resource == nil {
		return StringInput("Ingredient.Comment", nil)
	}
	return StringInput("Ingredient.Comment", resource.Comment)
}
func (resource *Ingredient) T_ManufacturerId(numManufacturer int) templ.Component {

	if resource == nil || len(resource.Manufacturer) >= numManufacturer {
		return StringInput("Ingredient.Manufacturer["+strconv.Itoa(numManufacturer)+"].Id", nil)
	}
	return StringInput("Ingredient.Manufacturer["+strconv.Itoa(numManufacturer)+"].Id", resource.Manufacturer[numManufacturer].Id)
}
func (resource *Ingredient) T_ManufacturerRole(numManufacturer int) templ.Component {
	optionsValueSet := VSIngredient_manufacturer_role

	if resource == nil || len(resource.Manufacturer) >= numManufacturer {
		return CodeSelect("Ingredient.Manufacturer["+strconv.Itoa(numManufacturer)+"].Role", nil, optionsValueSet)
	}
	return CodeSelect("Ingredient.Manufacturer["+strconv.Itoa(numManufacturer)+"].Role", resource.Manufacturer[numManufacturer].Role, optionsValueSet)
}
func (resource *Ingredient) T_SubstanceId() templ.Component {

	if resource == nil {
		return StringInput("Ingredient.Substance.Id", nil)
	}
	return StringInput("Ingredient.Substance.Id", resource.Substance.Id)
}
func (resource *Ingredient) T_SubstanceStrengthId(numStrength int) templ.Component {

	if resource == nil || len(resource.Substance.Strength) >= numStrength {
		return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].Id", nil)
	}
	return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].Id", resource.Substance.Strength[numStrength].Id)
}
func (resource *Ingredient) T_SubstanceStrengthTextPresentation(numStrength int) templ.Component {

	if resource == nil || len(resource.Substance.Strength) >= numStrength {
		return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].TextPresentation", nil)
	}
	return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].TextPresentation", resource.Substance.Strength[numStrength].TextPresentation)
}
func (resource *Ingredient) T_SubstanceStrengthTextConcentration(numStrength int) templ.Component {

	if resource == nil || len(resource.Substance.Strength) >= numStrength {
		return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].TextConcentration", nil)
	}
	return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].TextConcentration", resource.Substance.Strength[numStrength].TextConcentration)
}
func (resource *Ingredient) T_SubstanceStrengthBasis(numStrength int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Substance.Strength) >= numStrength {
		return CodeableConceptSelect("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].Basis", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].Basis", resource.Substance.Strength[numStrength].Basis, optionsValueSet)
}
func (resource *Ingredient) T_SubstanceStrengthMeasurementPoint(numStrength int) templ.Component {

	if resource == nil || len(resource.Substance.Strength) >= numStrength {
		return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].MeasurementPoint", nil)
	}
	return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].MeasurementPoint", resource.Substance.Strength[numStrength].MeasurementPoint)
}
func (resource *Ingredient) T_SubstanceStrengthCountry(numStrength int, numCountry int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Substance.Strength) >= numStrength || len(resource.Substance.Strength[numStrength].Country) >= numCountry {
		return CodeableConceptSelect("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].Country["+strconv.Itoa(numCountry)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].Country["+strconv.Itoa(numCountry)+"]", &resource.Substance.Strength[numStrength].Country[numCountry], optionsValueSet)
}
func (resource *Ingredient) T_SubstanceStrengthReferenceStrengthId(numStrength int, numReferenceStrength int) templ.Component {

	if resource == nil || len(resource.Substance.Strength) >= numStrength || len(resource.Substance.Strength[numStrength].ReferenceStrength) >= numReferenceStrength {
		return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Id", nil)
	}
	return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Id", resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].Id)
}
func (resource *Ingredient) T_SubstanceStrengthReferenceStrengthMeasurementPoint(numStrength int, numReferenceStrength int) templ.Component {

	if resource == nil || len(resource.Substance.Strength) >= numStrength || len(resource.Substance.Strength[numStrength].ReferenceStrength) >= numReferenceStrength {
		return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].MeasurementPoint", nil)
	}
	return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].MeasurementPoint", resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].MeasurementPoint)
}
func (resource *Ingredient) T_SubstanceStrengthReferenceStrengthCountry(numStrength int, numReferenceStrength int, numCountry int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Substance.Strength) >= numStrength || len(resource.Substance.Strength[numStrength].ReferenceStrength) >= numReferenceStrength || len(resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].Country) >= numCountry {
		return CodeableConceptSelect("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Country["+strconv.Itoa(numCountry)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Country["+strconv.Itoa(numCountry)+"]", &resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].Country[numCountry], optionsValueSet)
}
