package r4b

//generated with command go run ./bultaoreune
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
	ValueDate            *FhirDate        `json:"valueDate,omitempty"`
	ValueBoolean         *bool            `json:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty"`
}

type OtherManufacturedItemDefinition ManufacturedItemDefinition

// struct -> json, automatically add resourceType=Patient
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
func (r ManufacturedItemDefinition) ResourceType() string {
	return "ManufacturedItemDefinition"
}

func (resource *ManufacturedItemDefinition) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_ManufacturedDoseForm(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("manufacturedDoseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("manufacturedDoseForm", &resource.ManufacturedDoseForm, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_UnitOfPresentation(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("unitOfPresentation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("unitOfPresentation", resource.UnitOfPresentation, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_Manufacturer(frs []FhirResource, numManufacturer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numManufacturer >= len(resource.Manufacturer) {
		return ReferenceInput(frs, "manufacturer["+strconv.Itoa(numManufacturer)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "manufacturer["+strconv.Itoa(numManufacturer)+"]", &resource.Manufacturer[numManufacturer], htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_Ingredient(numIngredient int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableConceptSelect("ingredient["+strconv.Itoa(numIngredient)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ingredient["+strconv.Itoa(numIngredient)+"]", &resource.Ingredient[numIngredient], optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_PropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", resource.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_PropertyValueQuantity(numProperty int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return QuantityInput("property["+strconv.Itoa(numProperty)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("property["+strconv.Itoa(numProperty)+"].valueQuantity", resource.Property[numProperty].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_PropertyValueDate(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return FhirDateInput("property["+strconv.Itoa(numProperty)+"].valueDate", nil, htmlAttrs)
	}
	return FhirDateInput("property["+strconv.Itoa(numProperty)+"].valueDate", resource.Property[numProperty].ValueDate, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_PropertyValueBoolean(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", resource.Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_PropertyValueAttachment(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return AttachmentInput("property["+strconv.Itoa(numProperty)+"].valueAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("property["+strconv.Itoa(numProperty)+"].valueAttachment", resource.Property[numProperty].ValueAttachment, htmlAttrs)
}
