package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueDate            *string          `json:"valueDate,omitempty"`
	ValueBoolean         *bool            `json:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty"`
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

func (resource *ManufacturedItemDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ManufacturedItemDefinition.Id", nil)
	}
	return StringInput("ManufacturedItemDefinition.Id", resource.Id)
}
func (resource *ManufacturedItemDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ManufacturedItemDefinition.ImplicitRules", nil)
	}
	return StringInput("ManufacturedItemDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *ManufacturedItemDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ManufacturedItemDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("ManufacturedItemDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *ManufacturedItemDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ManufacturedItemDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("ManufacturedItemDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *ManufacturedItemDefinition) T_ManufacturedDoseForm(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ManufacturedItemDefinition.ManufacturedDoseForm", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ManufacturedItemDefinition.ManufacturedDoseForm", &resource.ManufacturedDoseForm, optionsValueSet)
}
func (resource *ManufacturedItemDefinition) T_UnitOfPresentation(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ManufacturedItemDefinition.UnitOfPresentation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ManufacturedItemDefinition.UnitOfPresentation", resource.UnitOfPresentation, optionsValueSet)
}
func (resource *ManufacturedItemDefinition) T_Ingredient(numIngredient int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Ingredient) >= numIngredient {
		return CodeableConceptSelect("ManufacturedItemDefinition.Ingredient["+strconv.Itoa(numIngredient)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ManufacturedItemDefinition.Ingredient["+strconv.Itoa(numIngredient)+"]", &resource.Ingredient[numIngredient], optionsValueSet)
}
func (resource *ManufacturedItemDefinition) T_PropertyId(numProperty int) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return StringInput("ManufacturedItemDefinition.Property["+strconv.Itoa(numProperty)+"].Id", nil)
	}
	return StringInput("ManufacturedItemDefinition.Property["+strconv.Itoa(numProperty)+"].Id", resource.Property[numProperty].Id)
}
func (resource *ManufacturedItemDefinition) T_PropertyType(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return CodeableConceptSelect("ManufacturedItemDefinition.Property["+strconv.Itoa(numProperty)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ManufacturedItemDefinition.Property["+strconv.Itoa(numProperty)+"].Type", &resource.Property[numProperty].Type, optionsValueSet)
}
