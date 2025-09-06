package r4

//generated with command go run ./bultaoreune -nodownload
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

func (resource *MedicinalProductIngredient) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductIngredient.Id", nil)
	}
	return StringInput("MedicinalProductIngredient.Id", resource.Id)
}
func (resource *MedicinalProductIngredient) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductIngredient.ImplicitRules", nil)
	}
	return StringInput("MedicinalProductIngredient.ImplicitRules", resource.ImplicitRules)
}
func (resource *MedicinalProductIngredient) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MedicinalProductIngredient.Language", nil, optionsValueSet)
	}
	return CodeSelect("MedicinalProductIngredient.Language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductIngredient) T_Role(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductIngredient.Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductIngredient.Role", &resource.Role, optionsValueSet)
}
func (resource *MedicinalProductIngredient) T_AllergenicIndicator() templ.Component {

	if resource == nil {
		return BoolInput("MedicinalProductIngredient.AllergenicIndicator", nil)
	}
	return BoolInput("MedicinalProductIngredient.AllergenicIndicator", resource.AllergenicIndicator)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceId(numSpecifiedSubstance int) templ.Component {

	if resource == nil || len(resource.SpecifiedSubstance) >= numSpecifiedSubstance {
		return StringInput("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Id", nil)
	}
	return StringInput("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Id", resource.SpecifiedSubstance[numSpecifiedSubstance].Id)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceCode(numSpecifiedSubstance int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SpecifiedSubstance) >= numSpecifiedSubstance {
		return CodeableConceptSelect("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Code", &resource.SpecifiedSubstance[numSpecifiedSubstance].Code, optionsValueSet)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceGroup(numSpecifiedSubstance int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SpecifiedSubstance) >= numSpecifiedSubstance {
		return CodeableConceptSelect("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Group", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Group", &resource.SpecifiedSubstance[numSpecifiedSubstance].Group, optionsValueSet)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceConfidentiality(numSpecifiedSubstance int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SpecifiedSubstance) >= numSpecifiedSubstance {
		return CodeableConceptSelect("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Confidentiality", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Confidentiality", resource.SpecifiedSubstance[numSpecifiedSubstance].Confidentiality, optionsValueSet)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceStrengthId(numSpecifiedSubstance int, numStrength int) templ.Component {

	if resource == nil || len(resource.SpecifiedSubstance) >= numSpecifiedSubstance || len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength) >= numStrength {
		return StringInput("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Strength["+strconv.Itoa(numStrength)+"].Id", nil)
	}
	return StringInput("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Strength["+strconv.Itoa(numStrength)+"].Id", resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].Id)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceStrengthMeasurementPoint(numSpecifiedSubstance int, numStrength int) templ.Component {

	if resource == nil || len(resource.SpecifiedSubstance) >= numSpecifiedSubstance || len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength) >= numStrength {
		return StringInput("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Strength["+strconv.Itoa(numStrength)+"].MeasurementPoint", nil)
	}
	return StringInput("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Strength["+strconv.Itoa(numStrength)+"].MeasurementPoint", resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].MeasurementPoint)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceStrengthCountry(numSpecifiedSubstance int, numStrength int, numCountry int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SpecifiedSubstance) >= numSpecifiedSubstance || len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength) >= numStrength || len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].Country) >= numCountry {
		return CodeableConceptSelect("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Strength["+strconv.Itoa(numStrength)+"].Country["+strconv.Itoa(numCountry)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Strength["+strconv.Itoa(numStrength)+"].Country["+strconv.Itoa(numCountry)+"]", &resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].Country[numCountry], optionsValueSet)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceStrengthReferenceStrengthId(numSpecifiedSubstance int, numStrength int, numReferenceStrength int) templ.Component {

	if resource == nil || len(resource.SpecifiedSubstance) >= numSpecifiedSubstance || len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength) >= numStrength || len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].ReferenceStrength) >= numReferenceStrength {
		return StringInput("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Id", nil)
	}
	return StringInput("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Id", resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].ReferenceStrength[numReferenceStrength].Id)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceStrengthReferenceStrengthSubstance(numSpecifiedSubstance int, numStrength int, numReferenceStrength int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SpecifiedSubstance) >= numSpecifiedSubstance || len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength) >= numStrength || len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].ReferenceStrength) >= numReferenceStrength {
		return CodeableConceptSelect("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Substance", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Substance", resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].ReferenceStrength[numReferenceStrength].Substance, optionsValueSet)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceStrengthReferenceStrengthMeasurementPoint(numSpecifiedSubstance int, numStrength int, numReferenceStrength int) templ.Component {

	if resource == nil || len(resource.SpecifiedSubstance) >= numSpecifiedSubstance || len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength) >= numStrength || len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].ReferenceStrength) >= numReferenceStrength {
		return StringInput("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].MeasurementPoint", nil)
	}
	return StringInput("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].MeasurementPoint", resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].ReferenceStrength[numReferenceStrength].MeasurementPoint)
}
func (resource *MedicinalProductIngredient) T_SpecifiedSubstanceStrengthReferenceStrengthCountry(numSpecifiedSubstance int, numStrength int, numReferenceStrength int, numCountry int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SpecifiedSubstance) >= numSpecifiedSubstance || len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength) >= numStrength || len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].ReferenceStrength) >= numReferenceStrength || len(resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].ReferenceStrength[numReferenceStrength].Country) >= numCountry {
		return CodeableConceptSelect("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Country["+strconv.Itoa(numCountry)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductIngredient.SpecifiedSubstance["+strconv.Itoa(numSpecifiedSubstance)+"].Strength["+strconv.Itoa(numStrength)+"].ReferenceStrength["+strconv.Itoa(numReferenceStrength)+"].Country["+strconv.Itoa(numCountry)+"]", &resource.SpecifiedSubstance[numSpecifiedSubstance].Strength[numStrength].ReferenceStrength[numReferenceStrength].Country[numCountry], optionsValueSet)
}
func (resource *MedicinalProductIngredient) T_SubstanceId() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductIngredient.Substance.Id", nil)
	}
	return StringInput("MedicinalProductIngredient.Substance.Id", resource.Substance.Id)
}
func (resource *MedicinalProductIngredient) T_SubstanceCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductIngredient.Substance.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductIngredient.Substance.Code", &resource.Substance.Code, optionsValueSet)
}
