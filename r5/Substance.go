package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Expiry            *time.Time            `json:"expiry,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *Substance) T_Instance(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Instance", nil, htmlAttrs)
	}
	return BoolInput("Instance", &resource.Instance, htmlAttrs)
}
func (resource *Substance) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSSubstance_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Substance) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Substance) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *Substance) T_Expiry(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Expiry", nil, htmlAttrs)
	}
	return DateTimeInput("Expiry", resource.Expiry, htmlAttrs)
}
func (resource *Substance) T_IngredientSubstanceCodeableConcept(numIngredient int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableConceptSelect("Ingredient["+strconv.Itoa(numIngredient)+"]SubstanceCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Ingredient["+strconv.Itoa(numIngredient)+"]SubstanceCodeableConcept", &resource.Ingredient[numIngredient].SubstanceCodeableConcept, optionsValueSet, htmlAttrs)
}
