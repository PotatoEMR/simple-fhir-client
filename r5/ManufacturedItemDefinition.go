package r5

//generated with command go run ./bultaoreune -nodownload
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
func (resource *ManufacturedItemDefinition) T_Name() templ.Component {

	if resource == nil {
		return StringInput("ManufacturedItemDefinition.Name", nil)
	}
	return StringInput("ManufacturedItemDefinition.Name", resource.Name)
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
func (resource *ManufacturedItemDefinition) T_ComponentId(numComponent int) templ.Component {

	if resource == nil || len(resource.Component) >= numComponent {
		return StringInput("ManufacturedItemDefinition.Component["+strconv.Itoa(numComponent)+"].Id", nil)
	}
	return StringInput("ManufacturedItemDefinition.Component["+strconv.Itoa(numComponent)+"].Id", resource.Component[numComponent].Id)
}
func (resource *ManufacturedItemDefinition) T_ComponentType(numComponent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Component) >= numComponent {
		return CodeableConceptSelect("ManufacturedItemDefinition.Component["+strconv.Itoa(numComponent)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ManufacturedItemDefinition.Component["+strconv.Itoa(numComponent)+"].Type", &resource.Component[numComponent].Type, optionsValueSet)
}
func (resource *ManufacturedItemDefinition) T_ComponentFunction(numComponent int, numFunction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Component) >= numComponent || len(resource.Component[numComponent].Function) >= numFunction {
		return CodeableConceptSelect("ManufacturedItemDefinition.Component["+strconv.Itoa(numComponent)+"].Function["+strconv.Itoa(numFunction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ManufacturedItemDefinition.Component["+strconv.Itoa(numComponent)+"].Function["+strconv.Itoa(numFunction)+"]", &resource.Component[numComponent].Function[numFunction], optionsValueSet)
}
func (resource *ManufacturedItemDefinition) T_ComponentConstituentId(numComponent int, numConstituent int) templ.Component {

	if resource == nil || len(resource.Component) >= numComponent || len(resource.Component[numComponent].Constituent) >= numConstituent {
		return StringInput("ManufacturedItemDefinition.Component["+strconv.Itoa(numComponent)+"].Constituent["+strconv.Itoa(numConstituent)+"].Id", nil)
	}
	return StringInput("ManufacturedItemDefinition.Component["+strconv.Itoa(numComponent)+"].Constituent["+strconv.Itoa(numConstituent)+"].Id", resource.Component[numComponent].Constituent[numConstituent].Id)
}
func (resource *ManufacturedItemDefinition) T_ComponentConstituentLocation(numComponent int, numConstituent int, numLocation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Component) >= numComponent || len(resource.Component[numComponent].Constituent) >= numConstituent || len(resource.Component[numComponent].Constituent[numConstituent].Location) >= numLocation {
		return CodeableConceptSelect("ManufacturedItemDefinition.Component["+strconv.Itoa(numComponent)+"].Constituent["+strconv.Itoa(numConstituent)+"].Location["+strconv.Itoa(numLocation)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ManufacturedItemDefinition.Component["+strconv.Itoa(numComponent)+"].Constituent["+strconv.Itoa(numConstituent)+"].Location["+strconv.Itoa(numLocation)+"]", &resource.Component[numComponent].Constituent[numConstituent].Location[numLocation], optionsValueSet)
}
func (resource *ManufacturedItemDefinition) T_ComponentConstituentFunction(numComponent int, numConstituent int, numFunction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Component) >= numComponent || len(resource.Component[numComponent].Constituent) >= numConstituent || len(resource.Component[numComponent].Constituent[numConstituent].Function) >= numFunction {
		return CodeableConceptSelect("ManufacturedItemDefinition.Component["+strconv.Itoa(numComponent)+"].Constituent["+strconv.Itoa(numConstituent)+"].Function["+strconv.Itoa(numFunction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ManufacturedItemDefinition.Component["+strconv.Itoa(numComponent)+"].Constituent["+strconv.Itoa(numConstituent)+"].Function["+strconv.Itoa(numFunction)+"]", &resource.Component[numComponent].Constituent[numConstituent].Function[numFunction], optionsValueSet)
}
