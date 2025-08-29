package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	ValueQuantity        *Quantity        `json:"valueQuantity"`
	ValueDate            *string          `json:"valueDate"`
	ValueBoolean         *bool            `json:"valueBoolean"`
	ValueMarkdown        *string          `json:"valueMarkdown"`
	ValueAttachment      *Attachment      `json:"valueAttachment"`
	ValueReference       *Reference       `json:"valueReference"`
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

func (resource *ManufacturedItemDefinition) ManufacturedItemDefinitionLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *ManufacturedItemDefinition) ManufacturedItemDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
