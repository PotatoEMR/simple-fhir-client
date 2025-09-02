package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/ManufacturedItemDefinition
type ManufacturedItemDefinition struct {
	Id                   *string                              `json:"id,omitempty"`
	Meta                 *Meta                                `json:"meta,omitempty"`
	ImplicitRules        *string                              `json:"implicitRules,omitempty"`
	Language             *string                              `json:"language,omitempty"`
	Text                 *Narrative                           `json:"text,omitempty"`
	Contained            []Resource                           `json:"contained,omitempty"`
	Extension            []Extension                          `json:"extension,omitempty"`
	ModifierExtension    []Extension                          `json:"modifierExtension,omitempty"`
	Identifier           []Identifier                         `json:"identifier,omitempty"`
	Status               string                               `json:"status"`
	ManufacturedDoseForm CodeableConcept                      `json:"manufacturedDoseForm"`
	UnitOfPresentation   *CodeableConcept                     `json:"unitOfPresentation,omitempty"`
	Manufacturer         []Reference                          `json:"manufacturer,omitempty"`
	Ingredient           []CodeableConcept                    `json:"ingredient,omitempty"`
	Property             []ManufacturedItemDefinitionProperty `json:"property,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ManufacturedItemDefinition
type ManufacturedItemDefinitionProperty struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept  `json:"type"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	ValueQuantity        *Quantity        `json:"valueQuantity"`
	ValueDate            *string          `json:"valueDate"`
	ValueBoolean         *bool            `json:"valueBoolean"`
	ValueAttachment      *Attachment      `json:"valueAttachment"`
}

type OtherManufacturedItemDefinition ManufacturedItemDefinition

// on convert struct to json, automatically add resourceType=ManufacturedItemDefinition
func (r ManufacturedItemDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherManufacturedItemDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherManufacturedItemDefinition: OtherManufacturedItemDefinition(r),
		ResourceType:                    "ManufacturedItemDefinition",
	})
}

func (resource *ManufacturedItemDefinition) ManufacturedItemDefinitionLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ManufacturedItemDefinition) ManufacturedItemDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ManufacturedItemDefinition) ManufacturedItemDefinitionManufacturedDoseForm(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("manufacturedDoseForm", nil, optionsValueSet)
	}
	return CodeableConceptSelect("manufacturedDoseForm", &resource.ManufacturedDoseForm, optionsValueSet)
}
func (resource *ManufacturedItemDefinition) ManufacturedItemDefinitionUnitOfPresentation(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("unitOfPresentation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("unitOfPresentation", resource.UnitOfPresentation, optionsValueSet)
}
func (resource *ManufacturedItemDefinition) ManufacturedItemDefinitionIngredient(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("ingredient", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ingredient", &resource.Ingredient[0], optionsValueSet)
}
func (resource *ManufacturedItemDefinition) ManufacturedItemDefinitionPropertyType(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Property) >= numProperty {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Property[numProperty].Type, optionsValueSet)
}
