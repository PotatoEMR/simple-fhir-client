package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/ChargeItemDefinition
type ChargeItemDefinition struct {
	Id                *string                             `json:"id,omitempty"`
	Meta              *Meta                               `json:"meta,omitempty"`
	ImplicitRules     *string                             `json:"implicitRules,omitempty"`
	Language          *string                             `json:"language,omitempty"`
	Text              *Narrative                          `json:"text,omitempty"`
	Contained         []Resource                          `json:"contained,omitempty"`
	Extension         []Extension                         `json:"extension,omitempty"`
	ModifierExtension []Extension                         `json:"modifierExtension,omitempty"`
	Url               string                              `json:"url"`
	Identifier        []Identifier                        `json:"identifier,omitempty"`
	Version           *string                             `json:"version,omitempty"`
	Title             *string                             `json:"title,omitempty"`
	DerivedFromUri    []string                            `json:"derivedFromUri,omitempty"`
	PartOf            []string                            `json:"partOf,omitempty"`
	Replaces          []string                            `json:"replaces,omitempty"`
	Status            string                              `json:"status"`
	Experimental      *bool                               `json:"experimental,omitempty"`
	Date              *string                             `json:"date,omitempty"`
	Publisher         *string                             `json:"publisher,omitempty"`
	Contact           []ContactDetail                     `json:"contact,omitempty"`
	Description       *string                             `json:"description,omitempty"`
	UseContext        []UsageContext                      `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                   `json:"jurisdiction,omitempty"`
	Copyright         *string                             `json:"copyright,omitempty"`
	ApprovalDate      *string                             `json:"approvalDate,omitempty"`
	LastReviewDate    *string                             `json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period                             `json:"effectivePeriod,omitempty"`
	Code              *CodeableConcept                    `json:"code,omitempty"`
	Instance          []Reference                         `json:"instance,omitempty"`
	Applicability     []ChargeItemDefinitionApplicability `json:"applicability,omitempty"`
	PropertyGroup     []ChargeItemDefinitionPropertyGroup `json:"propertyGroup,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ChargeItemDefinition
type ChargeItemDefinitionApplicability struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Expression        *string     `json:"expression,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ChargeItemDefinition
type ChargeItemDefinitionPropertyGroup struct {
	Id                *string                                           `json:"id,omitempty"`
	Extension         []Extension                                       `json:"extension,omitempty"`
	ModifierExtension []Extension                                       `json:"modifierExtension,omitempty"`
	PriceComponent    []ChargeItemDefinitionPropertyGroupPriceComponent `json:"priceComponent,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ChargeItemDefinition
type ChargeItemDefinitionPropertyGroupPriceComponent struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              string           `json:"type"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Factor            *float64         `json:"factor,omitempty"`
	Amount            *Money           `json:"amount,omitempty"`
}

type OtherChargeItemDefinition ChargeItemDefinition

// on convert struct to json, automatically add resourceType=ChargeItemDefinition
func (r ChargeItemDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherChargeItemDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherChargeItemDefinition: OtherChargeItemDefinition(r),
		ResourceType:              "ChargeItemDefinition",
	})
}

func (resource *ChargeItemDefinition) ChargeItemDefinitionLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *ChargeItemDefinition) ChargeItemDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *ChargeItemDefinition) ChargeItemDefinitionPropertyGroupPriceComponentType(numPropertyGroup int, numPriceComponent int) templ.Component {
	optionsValueSet := VSInvoice_priceComponentType
	currentVal := ""
	if resource != nil && len(resource.PropertyGroup[numPropertyGroup].PriceComponent) >= numPriceComponent {
		currentVal = resource.PropertyGroup[numPropertyGroup].PriceComponent[numPriceComponent].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
