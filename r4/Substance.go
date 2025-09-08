package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Substance
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
	Status            *string               `json:"status,omitempty"`
	Category          []CodeableConcept     `json:"category,omitempty"`
	Code              CodeableConcept       `json:"code"`
	Description       *string               `json:"description,omitempty"`
	Instance          []SubstanceInstance   `json:"instance,omitempty"`
	Ingredient        []SubstanceIngredient `json:"ingredient,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Substance
type SubstanceInstance struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Identifier        *Identifier `json:"identifier,omitempty"`
	Expiry            *time.Time  `json:"expiry,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Quantity          *Quantity   `json:"quantity,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Substance
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
func (r Substance) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Substance/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Substance"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Substance) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSSubstance_status

	if resource == nil {
		return CodeSelect("Substance.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Substance.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Substance) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Substance.Category."+strconv.Itoa(numCategory)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Substance.Category."+strconv.Itoa(numCategory)+".", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Substance) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Substance.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Substance.Code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Substance) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Substance.Description", nil, htmlAttrs)
	}
	return StringInput("Substance.Description", resource.Description, htmlAttrs)
}
func (resource *Substance) T_InstanceExpiry(numInstance int, htmlAttrs string) templ.Component {

	if resource == nil || numInstance >= len(resource.Instance) {
		return DateTimeInput("Substance.Instance."+strconv.Itoa(numInstance)+"..Expiry", nil, htmlAttrs)
	}
	return DateTimeInput("Substance.Instance."+strconv.Itoa(numInstance)+"..Expiry", resource.Instance[numInstance].Expiry, htmlAttrs)
}
func (resource *Substance) T_IngredientSubstanceCodeableConcept(numIngredient int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableConceptSelect("Substance.Ingredient."+strconv.Itoa(numIngredient)+"..SubstanceCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Substance.Ingredient."+strconv.Itoa(numIngredient)+"..SubstanceCodeableConcept", &resource.Ingredient[numIngredient].SubstanceCodeableConcept, optionsValueSet, htmlAttrs)
}
