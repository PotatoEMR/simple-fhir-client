package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ManufacturedItemDefinition
type ManufacturedItemDefinition struct {
	Id                   *string                               `json:"id,omitempty"`
	Meta                 *Meta                                 `json:"meta,omitempty"`
	ImplicitRules        *string                               `json:"implicitRules,omitempty"`
	Language             *string                               `json:"language,omitempty"`
	Text                 *Narrative                            `json:"text,omitempty"`
	Contained            []Resource                            `json:"contained,omitempty"`
	Extension            []Extension                           `json:"extension,omitempty"`
	ModifierExtension    []Extension                           `json:"modifierExtension,omitempty"`
	Identifier           []Identifier                          `json:"identifier,omitempty"`
	Status               string                                `json:"status"`
	Name                 *string                               `json:"name,omitempty"`
	ManufacturedDoseForm CodeableConcept                       `json:"manufacturedDoseForm"`
	UnitOfPresentation   *CodeableConcept                      `json:"unitOfPresentation,omitempty"`
	Manufacturer         []Reference                           `json:"manufacturer,omitempty"`
	MarketingStatus      []MarketingStatus                     `json:"marketingStatus,omitempty"`
	Ingredient           []CodeableConcept                     `json:"ingredient,omitempty"`
	Property             []ManufacturedItemDefinitionProperty  `json:"property,omitempty"`
	Component            []ManufacturedItemDefinitionComponent `json:"component,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ManufacturedItemDefinition
type ManufacturedItemDefinitionProperty struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept  `json:"type"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueDate            *string          `json:"valueDate,omitempty"`
	ValueBoolean         *bool            `json:"valueBoolean,omitempty"`
	ValueMarkdown        *string          `json:"valueMarkdown,omitempty"`
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty"`
	ValueReference       *Reference       `json:"valueReference,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ManufacturedItemDefinition
type ManufacturedItemDefinitionComponent struct {
	Id                *string                                          `json:"id,omitempty"`
	Extension         []Extension                                      `json:"extension,omitempty"`
	ModifierExtension []Extension                                      `json:"modifierExtension,omitempty"`
	Type              CodeableConcept                                  `json:"type"`
	Function          []CodeableConcept                                `json:"function,omitempty"`
	Amount            []Quantity                                       `json:"amount,omitempty"`
	Constituent       []ManufacturedItemDefinitionComponentConstituent `json:"constituent,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ManufacturedItemDefinition
type ManufacturedItemDefinitionComponentConstituent struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Amount            []Quantity          `json:"amount,omitempty"`
	Location          []CodeableConcept   `json:"location,omitempty"`
	Function          []CodeableConcept   `json:"function,omitempty"`
	HasIngredient     []CodeableReference `json:"hasIngredient,omitempty"`
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
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_ManufacturedDoseForm(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("manufacturedDoseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("manufacturedDoseForm", &resource.ManufacturedDoseForm, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_UnitOfPresentation(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("unitOfPresentation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("unitOfPresentation", resource.UnitOfPresentation, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_Ingredient(numIngredient int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableConceptSelect("ingredient["+strconv.Itoa(numIngredient)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ingredient["+strconv.Itoa(numIngredient)+"]", &resource.Ingredient[numIngredient], optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_PropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", resource.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_PropertyValueDate(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return DateInput("property["+strconv.Itoa(numProperty)+"].valueDate", nil, htmlAttrs)
	}
	return DateInput("property["+strconv.Itoa(numProperty)+"].valueDate", resource.Property[numProperty].ValueDate, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_PropertyValueBoolean(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", resource.Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_PropertyValueMarkdown(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("property["+strconv.Itoa(numProperty)+"].valueMarkdown", nil, htmlAttrs)
	}
	return StringInput("property["+strconv.Itoa(numProperty)+"].valueMarkdown", resource.Property[numProperty].ValueMarkdown, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_ComponentType(numComponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].type", &resource.Component[numComponent].Type, optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_ComponentFunction(numComponent int, numFunction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) || numFunction >= len(resource.Component[numComponent].Function) {
		return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].function["+strconv.Itoa(numFunction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].function["+strconv.Itoa(numFunction)+"]", &resource.Component[numComponent].Function[numFunction], optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_ComponentConstituentLocation(numComponent int, numConstituent int, numLocation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) || numConstituent >= len(resource.Component[numComponent].Constituent) || numLocation >= len(resource.Component[numComponent].Constituent[numConstituent].Location) {
		return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].constituent["+strconv.Itoa(numConstituent)+"].location["+strconv.Itoa(numLocation)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].constituent["+strconv.Itoa(numConstituent)+"].location["+strconv.Itoa(numLocation)+"]", &resource.Component[numComponent].Constituent[numConstituent].Location[numLocation], optionsValueSet, htmlAttrs)
}
func (resource *ManufacturedItemDefinition) T_ComponentConstituentFunction(numComponent int, numConstituent int, numFunction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) || numConstituent >= len(resource.Component[numComponent].Constituent) || numFunction >= len(resource.Component[numComponent].Constituent[numConstituent].Function) {
		return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].constituent["+strconv.Itoa(numConstituent)+"].function["+strconv.Itoa(numFunction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("component["+strconv.Itoa(numComponent)+"].constituent["+strconv.Itoa(numConstituent)+"].function["+strconv.Itoa(numFunction)+"]", &resource.Component[numComponent].Constituent[numConstituent].Function[numFunction], optionsValueSet, htmlAttrs)
}
