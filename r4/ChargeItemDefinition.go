package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *ChargeItemDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ChargeItemDefinition.Id", nil)
	}
	return StringInput("ChargeItemDefinition.Id", resource.Id)
}
func (resource *ChargeItemDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ChargeItemDefinition.ImplicitRules", nil)
	}
	return StringInput("ChargeItemDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *ChargeItemDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ChargeItemDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("ChargeItemDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *ChargeItemDefinition) T_Url() templ.Component {

	if resource == nil {
		return StringInput("ChargeItemDefinition.Url", nil)
	}
	return StringInput("ChargeItemDefinition.Url", &resource.Url)
}
func (resource *ChargeItemDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("ChargeItemDefinition.Version", nil)
	}
	return StringInput("ChargeItemDefinition.Version", resource.Version)
}
func (resource *ChargeItemDefinition) T_Title() templ.Component {

	if resource == nil {
		return StringInput("ChargeItemDefinition.Title", nil)
	}
	return StringInput("ChargeItemDefinition.Title", resource.Title)
}
func (resource *ChargeItemDefinition) T_DerivedFromUri(numDerivedFromUri int) templ.Component {

	if resource == nil || len(resource.DerivedFromUri) >= numDerivedFromUri {
		return StringInput("ChargeItemDefinition.DerivedFromUri["+strconv.Itoa(numDerivedFromUri)+"]", nil)
	}
	return StringInput("ChargeItemDefinition.DerivedFromUri["+strconv.Itoa(numDerivedFromUri)+"]", &resource.DerivedFromUri[numDerivedFromUri])
}
func (resource *ChargeItemDefinition) T_PartOf(numPartOf int) templ.Component {

	if resource == nil || len(resource.PartOf) >= numPartOf {
		return StringInput("ChargeItemDefinition.PartOf["+strconv.Itoa(numPartOf)+"]", nil)
	}
	return StringInput("ChargeItemDefinition.PartOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf])
}
func (resource *ChargeItemDefinition) T_Replaces(numReplaces int) templ.Component {

	if resource == nil || len(resource.Replaces) >= numReplaces {
		return StringInput("ChargeItemDefinition.Replaces["+strconv.Itoa(numReplaces)+"]", nil)
	}
	return StringInput("ChargeItemDefinition.Replaces["+strconv.Itoa(numReplaces)+"]", &resource.Replaces[numReplaces])
}
func (resource *ChargeItemDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ChargeItemDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("ChargeItemDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *ChargeItemDefinition) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("ChargeItemDefinition.Experimental", nil)
	}
	return BoolInput("ChargeItemDefinition.Experimental", resource.Experimental)
}
func (resource *ChargeItemDefinition) T_Date() templ.Component {

	if resource == nil {
		return StringInput("ChargeItemDefinition.Date", nil)
	}
	return StringInput("ChargeItemDefinition.Date", resource.Date)
}
func (resource *ChargeItemDefinition) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("ChargeItemDefinition.Publisher", nil)
	}
	return StringInput("ChargeItemDefinition.Publisher", resource.Publisher)
}
func (resource *ChargeItemDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("ChargeItemDefinition.Description", nil)
	}
	return StringInput("ChargeItemDefinition.Description", resource.Description)
}
func (resource *ChargeItemDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("ChargeItemDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ChargeItemDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *ChargeItemDefinition) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("ChargeItemDefinition.Copyright", nil)
	}
	return StringInput("ChargeItemDefinition.Copyright", resource.Copyright)
}
func (resource *ChargeItemDefinition) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("ChargeItemDefinition.ApprovalDate", nil)
	}
	return StringInput("ChargeItemDefinition.ApprovalDate", resource.ApprovalDate)
}
func (resource *ChargeItemDefinition) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("ChargeItemDefinition.LastReviewDate", nil)
	}
	return StringInput("ChargeItemDefinition.LastReviewDate", resource.LastReviewDate)
}
func (resource *ChargeItemDefinition) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ChargeItemDefinition.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ChargeItemDefinition.Code", resource.Code, optionsValueSet)
}
func (resource *ChargeItemDefinition) T_ApplicabilityId(numApplicability int) templ.Component {

	if resource == nil || len(resource.Applicability) >= numApplicability {
		return StringInput("ChargeItemDefinition.Applicability["+strconv.Itoa(numApplicability)+"].Id", nil)
	}
	return StringInput("ChargeItemDefinition.Applicability["+strconv.Itoa(numApplicability)+"].Id", resource.Applicability[numApplicability].Id)
}
func (resource *ChargeItemDefinition) T_ApplicabilityDescription(numApplicability int) templ.Component {

	if resource == nil || len(resource.Applicability) >= numApplicability {
		return StringInput("ChargeItemDefinition.Applicability["+strconv.Itoa(numApplicability)+"].Description", nil)
	}
	return StringInput("ChargeItemDefinition.Applicability["+strconv.Itoa(numApplicability)+"].Description", resource.Applicability[numApplicability].Description)
}
func (resource *ChargeItemDefinition) T_ApplicabilityLanguage(numApplicability int) templ.Component {

	if resource == nil || len(resource.Applicability) >= numApplicability {
		return StringInput("ChargeItemDefinition.Applicability["+strconv.Itoa(numApplicability)+"].Language", nil)
	}
	return StringInput("ChargeItemDefinition.Applicability["+strconv.Itoa(numApplicability)+"].Language", resource.Applicability[numApplicability].Language)
}
func (resource *ChargeItemDefinition) T_ApplicabilityExpression(numApplicability int) templ.Component {

	if resource == nil || len(resource.Applicability) >= numApplicability {
		return StringInput("ChargeItemDefinition.Applicability["+strconv.Itoa(numApplicability)+"].Expression", nil)
	}
	return StringInput("ChargeItemDefinition.Applicability["+strconv.Itoa(numApplicability)+"].Expression", resource.Applicability[numApplicability].Expression)
}
func (resource *ChargeItemDefinition) T_PropertyGroupId(numPropertyGroup int) templ.Component {

	if resource == nil || len(resource.PropertyGroup) >= numPropertyGroup {
		return StringInput("ChargeItemDefinition.PropertyGroup["+strconv.Itoa(numPropertyGroup)+"].Id", nil)
	}
	return StringInput("ChargeItemDefinition.PropertyGroup["+strconv.Itoa(numPropertyGroup)+"].Id", resource.PropertyGroup[numPropertyGroup].Id)
}
func (resource *ChargeItemDefinition) T_PropertyGroupPriceComponentId(numPropertyGroup int, numPriceComponent int) templ.Component {

	if resource == nil || len(resource.PropertyGroup) >= numPropertyGroup || len(resource.PropertyGroup[numPropertyGroup].PriceComponent) >= numPriceComponent {
		return StringInput("ChargeItemDefinition.PropertyGroup["+strconv.Itoa(numPropertyGroup)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Id", nil)
	}
	return StringInput("ChargeItemDefinition.PropertyGroup["+strconv.Itoa(numPropertyGroup)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Id", resource.PropertyGroup[numPropertyGroup].PriceComponent[numPriceComponent].Id)
}
func (resource *ChargeItemDefinition) T_PropertyGroupPriceComponentType(numPropertyGroup int, numPriceComponent int) templ.Component {
	optionsValueSet := VSInvoice_priceComponentType

	if resource == nil || len(resource.PropertyGroup) >= numPropertyGroup || len(resource.PropertyGroup[numPropertyGroup].PriceComponent) >= numPriceComponent {
		return CodeSelect("ChargeItemDefinition.PropertyGroup["+strconv.Itoa(numPropertyGroup)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("ChargeItemDefinition.PropertyGroup["+strconv.Itoa(numPropertyGroup)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Type", &resource.PropertyGroup[numPropertyGroup].PriceComponent[numPriceComponent].Type, optionsValueSet)
}
func (resource *ChargeItemDefinition) T_PropertyGroupPriceComponentCode(numPropertyGroup int, numPriceComponent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PropertyGroup) >= numPropertyGroup || len(resource.PropertyGroup[numPropertyGroup].PriceComponent) >= numPriceComponent {
		return CodeableConceptSelect("ChargeItemDefinition.PropertyGroup["+strconv.Itoa(numPropertyGroup)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ChargeItemDefinition.PropertyGroup["+strconv.Itoa(numPropertyGroup)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Code", resource.PropertyGroup[numPropertyGroup].PriceComponent[numPriceComponent].Code, optionsValueSet)
}
func (resource *ChargeItemDefinition) T_PropertyGroupPriceComponentFactor(numPropertyGroup int, numPriceComponent int) templ.Component {

	if resource == nil || len(resource.PropertyGroup) >= numPropertyGroup || len(resource.PropertyGroup[numPropertyGroup].PriceComponent) >= numPriceComponent {
		return Float64Input("ChargeItemDefinition.PropertyGroup["+strconv.Itoa(numPropertyGroup)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Factor", nil)
	}
	return Float64Input("ChargeItemDefinition.PropertyGroup["+strconv.Itoa(numPropertyGroup)+"].PriceComponent["+strconv.Itoa(numPriceComponent)+"].Factor", resource.PropertyGroup[numPropertyGroup].PriceComponent[numPriceComponent].Factor)
}
