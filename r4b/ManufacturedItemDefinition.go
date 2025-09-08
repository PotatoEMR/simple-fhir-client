package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	ValueDate            *time.Time       `json:"valueDate,omitempty,format:'2006-01-02'"`
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
func (r ManufacturedItemDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ManufacturedItemDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ManufacturedItemDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ManufacturedItemDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ManufacturedItemDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ManufacturedItemDefinition.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_ManufacturedDoseForm(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ManufacturedItemDefinition.ManufacturedDoseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ManufacturedItemDefinition.ManufacturedDoseForm", &resource.ManufacturedDoseForm, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_UnitOfPresentation(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ManufacturedItemDefinition.UnitOfPresentation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ManufacturedItemDefinition.UnitOfPresentation", resource.UnitOfPresentation, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_Ingredient(numIngredient int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableConceptSelect("ManufacturedItemDefinition.Ingredient."+strconv.Itoa(numIngredient)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ManufacturedItemDefinition.Ingredient."+strconv.Itoa(numIngredient)+".", &resource.Ingredient[numIngredient], optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("ManufacturedItemDefinition.Property."+strconv.Itoa(numProperty)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ManufacturedItemDefinition.Property."+strconv.Itoa(numProperty)+"..Type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_PropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("ManufacturedItemDefinition.Property."+strconv.Itoa(numProperty)+"..ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ManufacturedItemDefinition.Property."+strconv.Itoa(numProperty)+"..ValueCodeableConcept", resource.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_PropertyValueDate(numProperty int, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return DateInput("ManufacturedItemDefinition.Property."+strconv.Itoa(numProperty)+"..ValueDate", nil, htmlAttrs)
	}
	return DateInput("ManufacturedItemDefinition.Property."+strconv.Itoa(numProperty)+"..ValueDate", resource.Property[numProperty].ValueDate, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_PropertyValueBoolean(numProperty int, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return BoolInput("ManufacturedItemDefinition.Property."+strconv.Itoa(numProperty)+"..ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("ManufacturedItemDefinition.Property."+strconv.Itoa(numProperty)+"..ValueBoolean", resource.Property[numProperty].ValueBoolean, htmlAttrs)
}
