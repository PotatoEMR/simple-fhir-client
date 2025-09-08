package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
func (r MedicinalProductIngredient) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicinalProductIngredient/" + *r.Id
		ref.Reference = &refStr
	}
	ref.Identifier = r.Identifier
	rtype := "MedicinalProductIngredient"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MedicinalProductIngredient) T_Role(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Role", &resource.Role, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductIngredient) T_AllergenicIndicator(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("AllergenicIndicator", nil, htmlAttrs)
	}
	return BoolInput("AllergenicIndicator", resource.AllergenicIndicator, htmlAttrs)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceCode(numSpecifiedSubstance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecifiedSubstance >= len(resource.SpecifiedSubstance) {
		return CodeableConceptSelect("SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"]Code", &resource.SpecifiedSubstance[numSpecifiedSubstance].Code, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceGroup(numSpecifiedSubstance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecifiedSubstance >= len(resource.SpecifiedSubstance) {
		return CodeableConceptSelect("SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"]Group", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"]Group", &resource.SpecifiedSubstance[numSpecifiedSubstance].Group, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceConfidentiality(numSpecifiedSubstance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecifiedSubstance >= len(resource.SpecifiedSubstance) {
		return CodeableConceptSelect("SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"]Confidentiality", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"]Confidentiality", resource.SpecifiedSubstance[numSpecifiedSubstance].Confidentiality, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceStrengthMeasurementPoint(numSpecifiedSubstance int, numStrength int, htmlAttrs string) templ.Component {
	if resource == nil || numSpecifiedSubstance >= len(resource.SpecifiedSubstance) || numStrength >= len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength) {
		return StringInput("SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"]Strength["+strconv.Itoa(numStrength)+"].MeasurementPoint", nil, htmlAttrs)
	}
	return StringInput("SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"]Strength["+strconv.Itoa(numStrength)+"].MeasurementPoint", resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].MeasurementPoint, htmlAttrs)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceStrengthCountry(numSpecifiedSubstance int, numStrength int, numCountry int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecifiedSubstance >= len(resource.SpecifiedSubstance) || numStrength >= len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength) || numCountry >= len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].Country) {
		return CodeableConceptSelect("SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"]Strength["+strconv.Itoa(numStrength)+"].Country["+strconv.Itoa(numCountry)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"]Strength["+strconv.Itoa(numStrength)+"].Country["+strconv.Itoa(numCountry)+"]", &resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].Country[numCountry], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceStrengthReferenceStrengthSubstance(numSpecifiedSubstance int, numStrength int, numReferenceStrength int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecifiedSubstance >= len(resource.SpecifiedSubstance) || numStrength >= len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength) || numReferenceStrength >= len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].ReferenceStrength) {
		return CodeableConceptSelect("SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"]Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Substance", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"]Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Substance", resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].ReferenceStrength[numReferenceStrength].Substance, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceStrengthReferenceStrengthMeasurementPoint(numSpecifiedSubstance int, numStrength int, numReferenceStrength int, htmlAttrs string) templ.Component {
	if resource == nil || numSpecifiedSubstance >= len(resource.SpecifiedSubstance) || numStrength >= len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength) || numReferenceStrength >= len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].ReferenceStrength) {
		return StringInput("SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"]Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].MeasurementPoint", nil, htmlAttrs)
	}
	return StringInput("SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"]Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].MeasurementPoint", resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].ReferenceStrength[numReferenceStrength].MeasurementPoint, htmlAttrs)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceStrengthReferenceStrengthCountry(numSpecifiedSubstance int, numStrength int, numReferenceStrength int, numCountry int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecifiedSubstance >= len(resource.SpecifiedSubstance) || numStrength >= len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength) || numReferenceStrength >= len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].ReferenceStrength) || numCountry >= len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].ReferenceStrength[numReferenceStrength].Country) {
		return CodeableConceptSelect("SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"]Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Country["+strconv.Itoa(numCountry)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"]Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Country["+strconv.Itoa(numCountry)+"]", &resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].ReferenceStrength[numReferenceStrength].Country[numCountry], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductIngredient) T_SubstanceCode(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SubstanceCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceCode", &resource.Substance.Code, optionsValueSet, htmlAttrs)
}
