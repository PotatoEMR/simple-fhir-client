package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Ingredient
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
	AllergenicIndicator *bool                    `json:"allergenicIndicator,omitempty"`
	Manufacturer        []IngredientManufacturer `json:"manufacturer,omitempty"`
	Substance           IngredientSubstance      `json:"substance"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Ingredient
type IngredientManufacturer struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Role              *string     `json:"role,omitempty"`
	Manufacturer      Reference   `json:"manufacturer"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Ingredient
type IngredientSubstance struct {
	Id                *string                       `json:"id,omitempty"`
	Extension         []Extension                   `json:"extension,omitempty"`
	ModifierExtension []Extension                   `json:"modifierExtension,omitempty"`
	Code              CodeableReference             `json:"code"`
	Strength          []IngredientSubstanceStrength `json:"strength,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Ingredient
type IngredientSubstanceStrength struct {
	Id                      *string                                        `json:"id,omitempty"`
	Extension               []Extension                                    `json:"extension,omitempty"`
	ModifierExtension       []Extension                                    `json:"modifierExtension,omitempty"`
	PresentationRatio       *Ratio                                         `json:"presentationRatio,omitempty"`
	PresentationRatioRange  *RatioRange                                    `json:"presentationRatioRange,omitempty"`
	TextPresentation        *string                                        `json:"textPresentation,omitempty"`
	ConcentrationRatio      *Ratio                                         `json:"concentrationRatio,omitempty"`
	ConcentrationRatioRange *RatioRange                                    `json:"concentrationRatioRange,omitempty"`
	TextConcentration       *string                                        `json:"textConcentration,omitempty"`
	MeasurementPoint        *string                                        `json:"measurementPoint,omitempty"`
	Country                 []CodeableConcept                              `json:"country,omitempty"`
	ReferenceStrength       []IngredientSubstanceStrengthReferenceStrength `json:"referenceStrength,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Ingredient
type IngredientSubstanceStrengthReferenceStrength struct {
	Id                 *string            `json:"id,omitempty"`
	Extension          []Extension        `json:"extension,omitempty"`
	ModifierExtension  []Extension        `json:"modifierExtension,omitempty"`
	Substance          *CodeableReference `json:"substance,omitempty"`
	StrengthRatio      Ratio              `json:"strengthRatio"`
	StrengthRatioRange RatioRange         `json:"strengthRatioRange"`
	MeasurementPoint   *string            `json:"measurementPoint,omitempty"`
	Country            []CodeableConcept  `json:"country,omitempty"`
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
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_Role(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Role", &resource.Role, optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_Function(numFunction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numFunction >= len(resource.Function) {
		return CodeableConceptSelect("Function["+strconv.Itoa(numFunction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Function["+strconv.Itoa(numFunction)+"]", &resource.Function[numFunction], optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_AllergenicIndicator(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("AllergenicIndicator", nil, htmlAttrs)
	}
	return BoolInput("AllergenicIndicator", resource.AllergenicIndicator, htmlAttrs)
}
func (resource *Ingredient) T_ManufacturerRole(numManufacturer int, htmlAttrs string) templ.Component {
	optionsValueSet := VSIngredient_manufacturer_role

	if resource == nil || numManufacturer >= len(resource.Manufacturer) {
		return CodeSelect("Manufacturer["+strconv.Itoa(numManufacturer)+"]Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Manufacturer["+strconv.Itoa(numManufacturer)+"]Role", resource.Manufacturer[numManufacturer].Role, optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthTextPresentation(numStrength int, htmlAttrs string) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) {
		return StringInput("SubstanceStrength["+strconv.Itoa(numStrength)+"].TextPresentation", nil, htmlAttrs)
	}
	return StringInput("SubstanceStrength["+strconv.Itoa(numStrength)+"].TextPresentation", resource.Substance.Strength[numStrength].TextPresentation, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthTextConcentration(numStrength int, htmlAttrs string) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) {
		return StringInput("SubstanceStrength["+strconv.Itoa(numStrength)+"].TextConcentration", nil, htmlAttrs)
	}
	return StringInput("SubstanceStrength["+strconv.Itoa(numStrength)+"].TextConcentration", resource.Substance.Strength[numStrength].TextConcentration, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthMeasurementPoint(numStrength int, htmlAttrs string) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) {
		return StringInput("SubstanceStrength["+strconv.Itoa(numStrength)+"].MeasurementPoint", nil, htmlAttrs)
	}
	return StringInput("SubstanceStrength["+strconv.Itoa(numStrength)+"].MeasurementPoint", resource.Substance.Strength[numStrength].MeasurementPoint, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthCountry(numStrength int, numCountry int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) || numCountry >= len(resource.Substance.Strength[numStrength].Country) {
		return CodeableConceptSelect("SubstanceStrength["+strconv.Itoa(numStrength)+"].Country["+strconv.Itoa(numCountry)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceStrength["+strconv.Itoa(numStrength)+"].Country["+strconv.Itoa(numCountry)+"]", &resource.Substance.Strength[numStrength].Country[numCountry], optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthReferenceStrengthMeasurementPoint(numStrength int, numReferenceStrength int, htmlAttrs string) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) || numReferenceStrength >= len(resource.Substance.Strength[numStrength].ReferenceStrength) {
		return StringInput("SubstanceStrength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].MeasurementPoint", nil, htmlAttrs)
	}
	return StringInput("SubstanceStrength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].MeasurementPoint", resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].MeasurementPoint, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthReferenceStrengthCountry(numStrength int, numReferenceStrength int, numCountry int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) || numReferenceStrength >= len(resource.Substance.Strength[numStrength].ReferenceStrength) || numCountry >= len(resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].Country) {
		return CodeableConceptSelect("SubstanceStrength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Country["+strconv.Itoa(numCountry)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceStrength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Country["+strconv.Itoa(numCountry)+"]", &resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].Country[numCountry], optionsValueSet, htmlAttrs)
}
