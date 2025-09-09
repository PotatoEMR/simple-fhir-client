package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/ChargeItemDefinition
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

// http://hl7.org/fhir/r4b/StructureDefinition/ChargeItemDefinition
type ChargeItemDefinitionApplicability struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Expression        *string     `json:"expression,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ChargeItemDefinition
type ChargeItemDefinitionPropertyGroup struct {
	Id                *string                                           `json:"id,omitempty"`
	Extension         []Extension                                       `json:"extension,omitempty"`
	ModifierExtension []Extension                                       `json:"modifierExtension,omitempty"`
	PriceComponent    []ChargeItemDefinitionPropertyGroupPriceComponent `json:"priceComponent,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ChargeItemDefinition
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
func (r ChargeItemDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ChargeItemDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ChargeItemDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ChargeItemDefinition) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", &resource.Url, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_DerivedFromUri(numDerivedFromUri int, htmlAttrs string) templ.Component {
	if resource == nil || numDerivedFromUri >= len(resource.DerivedFromUri) {
		return StringInput("derivedFromUri["+strconv.Itoa(numDerivedFromUri)+"]", nil, htmlAttrs)
	}
	return StringInput("derivedFromUri["+strconv.Itoa(numDerivedFromUri)+"]", &resource.DerivedFromUri[numDerivedFromUri], htmlAttrs)
}
func (resource *ChargeItemDefinition) T_PartOf(numPartOf int, htmlAttrs string) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return StringInput("partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return StringInput("partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Replaces(numReplaces int, htmlAttrs string) templ.Component {
	if resource == nil || numReplaces >= len(resource.Replaces) {
		return StringInput("replaces["+strconv.Itoa(numReplaces)+"]", nil, htmlAttrs)
	}
	return StringInput("replaces["+strconv.Itoa(numReplaces)+"]", &resource.Replaces[numReplaces], htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("approvalDate", nil, htmlAttrs)
	}
	return DateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("lastReviewDate", nil, htmlAttrs)
	}
	return DateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_ApplicabilityDescription(numApplicability int, htmlAttrs string) templ.Component {
	if resource == nil || numApplicability >= len(resource.Applicability) {
		return StringInput("applicability["+strconv.Itoa(numApplicability)+"].description", nil, htmlAttrs)
	}
	return StringInput("applicability["+strconv.Itoa(numApplicability)+"].description", resource.Applicability[numApplicability].Description, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_ApplicabilityExpression(numApplicability int, htmlAttrs string) templ.Component {
	if resource == nil || numApplicability >= len(resource.Applicability) {
		return StringInput("applicability["+strconv.Itoa(numApplicability)+"].expression", nil, htmlAttrs)
	}
	return StringInput("applicability["+strconv.Itoa(numApplicability)+"].expression", resource.Applicability[numApplicability].Expression, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_PropertyGroupPriceComponentType(numPropertyGroup int, numPriceComponent int, htmlAttrs string) templ.Component {
	optionsValueSet := VSInvoice_priceComponentType

	if resource == nil || numPropertyGroup >= len(resource.PropertyGroup) || numPriceComponent >= len(resource.PropertyGroup[numPropertyGroup].PriceComponent) {
		return CodeSelect("propertyGroup["+strconv.Itoa(numPropertyGroup)+"].priceComponent["+strconv.Itoa(numPriceComponent)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("propertyGroup["+strconv.Itoa(numPropertyGroup)+"].priceComponent["+strconv.Itoa(numPriceComponent)+"].type", &resource.PropertyGroup[numPropertyGroup].PriceComponent[numPriceComponent].Type, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_PropertyGroupPriceComponentCode(numPropertyGroup int, numPriceComponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPropertyGroup >= len(resource.PropertyGroup) || numPriceComponent >= len(resource.PropertyGroup[numPropertyGroup].PriceComponent) {
		return CodeableConceptSelect("propertyGroup["+strconv.Itoa(numPropertyGroup)+"].priceComponent["+strconv.Itoa(numPriceComponent)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("propertyGroup["+strconv.Itoa(numPropertyGroup)+"].priceComponent["+strconv.Itoa(numPriceComponent)+"].code", resource.PropertyGroup[numPropertyGroup].PriceComponent[numPriceComponent].Code, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItemDefinition) T_PropertyGroupPriceComponentFactor(numPropertyGroup int, numPriceComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numPropertyGroup >= len(resource.PropertyGroup) || numPriceComponent >= len(resource.PropertyGroup[numPropertyGroup].PriceComponent) {
		return Float64Input("propertyGroup["+strconv.Itoa(numPropertyGroup)+"].priceComponent["+strconv.Itoa(numPriceComponent)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("propertyGroup["+strconv.Itoa(numPropertyGroup)+"].priceComponent["+strconv.Itoa(numPriceComponent)+"].factor", resource.PropertyGroup[numPropertyGroup].PriceComponent[numPriceComponent].Factor, htmlAttrs)
}
