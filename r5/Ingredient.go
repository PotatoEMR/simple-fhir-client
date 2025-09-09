package r5

//generated with command go run ./bultaoreune
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
func (r Ingredient) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Ingredient/" + *r.Id
		ref.Reference = &refStr
	}
	ref.Identifier = r.Identifier
	rtype := "Ingredient"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Ingredient) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Ingredient.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Ingredient.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_Role(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Ingredient.Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Ingredient.Role", &resource.Role, optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_Function(numFunction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numFunction >= len(resource.Function) {
		return CodeableConceptSelect("Ingredient.Function["+strconv.Itoa(numFunction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Ingredient.Function["+strconv.Itoa(numFunction)+"]", &resource.Function[numFunction], optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_Group(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Ingredient.Group", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Ingredient.Group", resource.Group, optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_AllergenicIndicator(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Ingredient.AllergenicIndicator", nil, htmlAttrs)
	}
	return BoolInput("Ingredient.AllergenicIndicator", resource.AllergenicIndicator, htmlAttrs)
}
func (resource *Ingredient) T_Comment(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Ingredient.Comment", nil, htmlAttrs)
	}
	return StringInput("Ingredient.Comment", resource.Comment, htmlAttrs)
}
func (resource *Ingredient) T_ManufacturerRole(numManufacturer int, htmlAttrs string) templ.Component {
	optionsValueSet := VSIngredient_manufacturer_role

	if resource == nil || numManufacturer >= len(resource.Manufacturer) {
		return CodeSelect("Ingredient.Manufacturer["+strconv.Itoa(numManufacturer)+"].Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Ingredient.Manufacturer["+strconv.Itoa(numManufacturer)+"].Role", resource.Manufacturer[numManufacturer].Role, optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthPresentationCodeableConcept(numStrength int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) {
		return CodeableConceptSelect("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].PresentationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].PresentationCodeableConcept", resource.Substance.Strength[numStrength].PresentationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthTextPresentation(numStrength int, htmlAttrs string) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) {
		return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].TextPresentation", nil, htmlAttrs)
	}
	return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].TextPresentation", resource.Substance.Strength[numStrength].TextPresentation, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthConcentrationCodeableConcept(numStrength int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) {
		return CodeableConceptSelect("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].ConcentrationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].ConcentrationCodeableConcept", resource.Substance.Strength[numStrength].ConcentrationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthTextConcentration(numStrength int, htmlAttrs string) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) {
		return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].TextConcentration", nil, htmlAttrs)
	}
	return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].TextConcentration", resource.Substance.Strength[numStrength].TextConcentration, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthBasis(numStrength int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) {
		return CodeableConceptSelect("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].Basis", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].Basis", resource.Substance.Strength[numStrength].Basis, optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthMeasurementPoint(numStrength int, htmlAttrs string) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) {
		return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].MeasurementPoint", nil, htmlAttrs)
	}
	return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].MeasurementPoint", resource.Substance.Strength[numStrength].MeasurementPoint, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthCountry(numStrength int, numCountry int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) || numCountry >= len(resource.Substance.Strength[numStrength].Country) {
		return CodeableConceptSelect("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].Country["+strconv.Itoa(numCountry)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].Country["+strconv.Itoa(numCountry)+"]", &resource.Substance.Strength[numStrength].Country[numCountry], optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthReferenceStrengthMeasurementPoint(numStrength int, numReferenceStrength int, htmlAttrs string) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) || numReferenceStrength >= len(resource.Substance.Strength[numStrength].ReferenceStrength) {
		return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].MeasurementPoint", nil, htmlAttrs)
	}
	return StringInput("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].MeasurementPoint", resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].MeasurementPoint, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthReferenceStrengthCountry(numStrength int, numReferenceStrength int, numCountry int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) || numReferenceStrength >= len(resource.Substance.Strength[numStrength].ReferenceStrength) || numCountry >= len(resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].Country) {
		return CodeableConceptSelect("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Country["+strconv.Itoa(numCountry)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Ingredient.Substance.Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Country["+strconv.Itoa(numCountry)+"]", &resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].Country[numCountry], optionsValueSet, htmlAttrs)
}
