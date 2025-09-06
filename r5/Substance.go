package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Substance
type Substance struct {
	Id                *string               `json:"id,omitempty"`
	Meta              *Meta                 `json:"meta,omitempty"`
	ImplicitRules     *string               `json:"implicitRules,omitempty"`
	Language          *string               `json:"language,omitempty"`
	Text              *Narrative            `json:"text,omitempty"`
	Contained         []Resource            `json:"contained,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `json:"identifier,omitempty"`
	Instance          bool                  `json:"instance"`
	Status            *string               `json:"status,omitempty"`
	Category          []CodeableConcept     `json:"category,omitempty"`
	Code              CodeableReference     `json:"code"`
	Description       *string               `json:"description,omitempty"`
	Expiry            *string               `json:"expiry,omitempty"`
	Quantity          *Quantity             `json:"quantity,omitempty"`
	Ingredient        []SubstanceIngredient `json:"ingredient,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Substance
type SubstanceIngredient struct {
	Id                       *string         `json:"id,omitempty"`
	Extension                []Extension     `json:"extension,omitempty"`
	ModifierExtension        []Extension     `json:"modifierExtension,omitempty"`
	Quantity                 *Ratio          `json:"quantity,omitempty"`
	SubstanceCodeableConcept CodeableConcept `json:"substanceCodeableConcept"`
	SubstanceReference       Reference       `json:"substanceReference"`
}

type OtherSubstance Substance

// on convert struct to json, automatically add resourceType=Substance
func (r Substance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstance
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstance: OtherSubstance(r),
		ResourceType:   "Substance",
	})
}

func (resource *Substance) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Substance.Id", nil)
	}
	return StringInput("Substance.Id", resource.Id)
}
func (resource *Substance) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Substance.ImplicitRules", nil)
	}
	return StringInput("Substance.ImplicitRules", resource.ImplicitRules)
}
func (resource *Substance) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Substance.Language", nil, optionsValueSet)
	}
	return CodeSelect("Substance.Language", resource.Language, optionsValueSet)
}
func (resource *Substance) T_Instance() templ.Component {

	if resource == nil {
		return BoolInput("Substance.Instance", nil)
	}
	return BoolInput("Substance.Instance", &resource.Instance)
}
func (resource *Substance) T_Status() templ.Component {
	optionsValueSet := VSSubstance_status

	if resource == nil {
		return CodeSelect("Substance.Status", nil, optionsValueSet)
	}
	return CodeSelect("Substance.Status", resource.Status, optionsValueSet)
}
func (resource *Substance) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("Substance.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Substance.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *Substance) T_Description() templ.Component {

	if resource == nil {
		return StringInput("Substance.Description", nil)
	}
	return StringInput("Substance.Description", resource.Description)
}
func (resource *Substance) T_Expiry() templ.Component {

	if resource == nil {
		return StringInput("Substance.Expiry", nil)
	}
	return StringInput("Substance.Expiry", resource.Expiry)
}
func (resource *Substance) T_IngredientId(numIngredient int) templ.Component {

	if resource == nil || len(resource.Ingredient) >= numIngredient {
		return StringInput("Substance.Ingredient["+strconv.Itoa(numIngredient)+"].Id", nil)
	}
	return StringInput("Substance.Ingredient["+strconv.Itoa(numIngredient)+"].Id", resource.Ingredient[numIngredient].Id)
}
