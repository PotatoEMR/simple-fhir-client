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
func (resource *Ingredient) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_For(numFor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFor >= len(resource.For) {
		return ReferenceInput("for["+strconv.Itoa(numFor)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("for["+strconv.Itoa(numFor)+"]", &resource.For[numFor], htmlAttrs)
}
func (resource *Ingredient) T_Role(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("role", &resource.Role, optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_Function(numFunction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFunction >= len(resource.Function) {
		return CodeableConceptSelect("function["+strconv.Itoa(numFunction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("function["+strconv.Itoa(numFunction)+"]", &resource.Function[numFunction], optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_AllergenicIndicator(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("allergenicIndicator", nil, htmlAttrs)
	}
	return BoolInput("allergenicIndicator", resource.AllergenicIndicator, htmlAttrs)
}
func (resource *Ingredient) T_ManufacturerRole(numManufacturer int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSIngredient_manufacturer_role

	if resource == nil || numManufacturer >= len(resource.Manufacturer) {
		return CodeSelect("manufacturer["+strconv.Itoa(numManufacturer)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("manufacturer["+strconv.Itoa(numManufacturer)+"].role", resource.Manufacturer[numManufacturer].Role, optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_ManufacturerManufacturer(numManufacturer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numManufacturer >= len(resource.Manufacturer) {
		return ReferenceInput("manufacturer["+strconv.Itoa(numManufacturer)+"].manufacturer", nil, htmlAttrs)
	}
	return ReferenceInput("manufacturer["+strconv.Itoa(numManufacturer)+"].manufacturer", &resource.Manufacturer[numManufacturer].Manufacturer, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceCode(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("substance.code", nil, htmlAttrs)
	}
	return CodeableReferenceInput("substance.code", &resource.Substance.Code, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthPresentationRatio(numStrength int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) {
		return RatioInput("substance.strength["+strconv.Itoa(numStrength)+"].presentationRatio", nil, htmlAttrs)
	}
	return RatioInput("substance.strength["+strconv.Itoa(numStrength)+"].presentationRatio", resource.Substance.Strength[numStrength].PresentationRatio, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthPresentationRatioRange(numStrength int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) {
		return RatioRangeInput("substance.strength["+strconv.Itoa(numStrength)+"].presentationRatioRange", nil, htmlAttrs)
	}
	return RatioRangeInput("substance.strength["+strconv.Itoa(numStrength)+"].presentationRatioRange", resource.Substance.Strength[numStrength].PresentationRatioRange, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthTextPresentation(numStrength int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) {
		return StringInput("substance.strength["+strconv.Itoa(numStrength)+"].textPresentation", nil, htmlAttrs)
	}
	return StringInput("substance.strength["+strconv.Itoa(numStrength)+"].textPresentation", resource.Substance.Strength[numStrength].TextPresentation, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthConcentrationRatio(numStrength int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) {
		return RatioInput("substance.strength["+strconv.Itoa(numStrength)+"].concentrationRatio", nil, htmlAttrs)
	}
	return RatioInput("substance.strength["+strconv.Itoa(numStrength)+"].concentrationRatio", resource.Substance.Strength[numStrength].ConcentrationRatio, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthConcentrationRatioRange(numStrength int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) {
		return RatioRangeInput("substance.strength["+strconv.Itoa(numStrength)+"].concentrationRatioRange", nil, htmlAttrs)
	}
	return RatioRangeInput("substance.strength["+strconv.Itoa(numStrength)+"].concentrationRatioRange", resource.Substance.Strength[numStrength].ConcentrationRatioRange, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthTextConcentration(numStrength int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) {
		return StringInput("substance.strength["+strconv.Itoa(numStrength)+"].textConcentration", nil, htmlAttrs)
	}
	return StringInput("substance.strength["+strconv.Itoa(numStrength)+"].textConcentration", resource.Substance.Strength[numStrength].TextConcentration, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthMeasurementPoint(numStrength int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) {
		return StringInput("substance.strength["+strconv.Itoa(numStrength)+"].measurementPoint", nil, htmlAttrs)
	}
	return StringInput("substance.strength["+strconv.Itoa(numStrength)+"].measurementPoint", resource.Substance.Strength[numStrength].MeasurementPoint, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthCountry(numStrength int, numCountry int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) || numCountry >= len(resource.Substance.Strength[numStrength].Country) {
		return CodeableConceptSelect("substance.strength["+strconv.Itoa(numStrength)+"].country["+strconv.Itoa(numCountry)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("substance.strength["+strconv.Itoa(numStrength)+"].country["+strconv.Itoa(numCountry)+"]", &resource.Substance.Strength[numStrength].Country[numCountry], optionsValueSet, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthReferenceStrengthSubstance(numStrength int, numReferenceStrength int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) || numReferenceStrength >= len(resource.Substance.Strength[numStrength].ReferenceStrength) {
		return CodeableReferenceInput("substance.strength["+strconv.Itoa(numStrength)+"].referenceStrength["+strconv.Itoa(numReferenceStrength)+"].substance", nil, htmlAttrs)
	}
	return CodeableReferenceInput("substance.strength["+strconv.Itoa(numStrength)+"].referenceStrength["+strconv.Itoa(numReferenceStrength)+"].substance", resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].Substance, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthReferenceStrengthStrengthRatio(numStrength int, numReferenceStrength int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) || numReferenceStrength >= len(resource.Substance.Strength[numStrength].ReferenceStrength) {
		return RatioInput("substance.strength["+strconv.Itoa(numStrength)+"].referenceStrength["+strconv.Itoa(numReferenceStrength)+"].strengthRatio", nil, htmlAttrs)
	}
	return RatioInput("substance.strength["+strconv.Itoa(numStrength)+"].referenceStrength["+strconv.Itoa(numReferenceStrength)+"].strengthRatio", &resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].StrengthRatio, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthReferenceStrengthStrengthRatioRange(numStrength int, numReferenceStrength int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) || numReferenceStrength >= len(resource.Substance.Strength[numStrength].ReferenceStrength) {
		return RatioRangeInput("substance.strength["+strconv.Itoa(numStrength)+"].referenceStrength["+strconv.Itoa(numReferenceStrength)+"].strengthRatioRange", nil, htmlAttrs)
	}
	return RatioRangeInput("substance.strength["+strconv.Itoa(numStrength)+"].referenceStrength["+strconv.Itoa(numReferenceStrength)+"].strengthRatioRange", &resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].StrengthRatioRange, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthReferenceStrengthMeasurementPoint(numStrength int, numReferenceStrength int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) || numReferenceStrength >= len(resource.Substance.Strength[numStrength].ReferenceStrength) {
		return StringInput("substance.strength["+strconv.Itoa(numStrength)+"].referenceStrength["+strconv.Itoa(numReferenceStrength)+"].measurementPoint", nil, htmlAttrs)
	}
	return StringInput("substance.strength["+strconv.Itoa(numStrength)+"].referenceStrength["+strconv.Itoa(numReferenceStrength)+"].measurementPoint", resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].MeasurementPoint, htmlAttrs)
}
func (resource *Ingredient) T_SubstanceStrengthReferenceStrengthCountry(numStrength int, numReferenceStrength int, numCountry int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStrength >= len(resource.Substance.Strength) || numReferenceStrength >= len(resource.Substance.Strength[numStrength].ReferenceStrength) || numCountry >= len(resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].Country) {
		return CodeableConceptSelect("substance.strength["+strconv.Itoa(numStrength)+"].referenceStrength["+strconv.Itoa(numReferenceStrength)+"].country["+strconv.Itoa(numCountry)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("substance.strength["+strconv.Itoa(numStrength)+"].referenceStrength["+strconv.Itoa(numReferenceStrength)+"].country["+strconv.Itoa(numCountry)+"]", &resource.Substance.Strength[numStrength].ReferenceStrength[numReferenceStrength].Country[numCountry], optionsValueSet, htmlAttrs)
}
