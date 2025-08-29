package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	Expiry            *string     `json:"expiry,omitempty"`
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

func (resource *Substance) SubstanceLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Substance) SubstanceStatus() templ.Component {
	optionsValueSet := VSSubstance_status
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
