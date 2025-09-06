package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/ChargeItem
type ChargeItem struct {
	Id                     *string               `json:"id,omitempty"`
	Meta                   *Meta                 `json:"meta,omitempty"`
	ImplicitRules          *string               `json:"implicitRules,omitempty"`
	Language               *string               `json:"language,omitempty"`
	Text                   *Narrative            `json:"text,omitempty"`
	Contained              []Resource            `json:"contained,omitempty"`
	Extension              []Extension           `json:"extension,omitempty"`
	ModifierExtension      []Extension           `json:"modifierExtension,omitempty"`
	Identifier             []Identifier          `json:"identifier,omitempty"`
	DefinitionUri          []string              `json:"definitionUri,omitempty"`
	DefinitionCanonical    []string              `json:"definitionCanonical,omitempty"`
	Status                 string                `json:"status"`
	PartOf                 []Reference           `json:"partOf,omitempty"`
	Code                   CodeableConcept       `json:"code"`
	Subject                Reference             `json:"subject"`
	Context                *Reference            `json:"context,omitempty"`
	OccurrenceDateTime     *string               `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod       *Period               `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming       *Timing               `json:"occurrenceTiming,omitempty"`
	Performer              []ChargeItemPerformer `json:"performer,omitempty"`
	PerformingOrganization *Reference            `json:"performingOrganization,omitempty"`
	RequestingOrganization *Reference            `json:"requestingOrganization,omitempty"`
	CostCenter             *Reference            `json:"costCenter,omitempty"`
	Quantity               *Quantity             `json:"quantity,omitempty"`
	Bodysite               []CodeableConcept     `json:"bodysite,omitempty"`
	FactorOverride         *float64              `json:"factorOverride,omitempty"`
	PriceOverride          *Money                `json:"priceOverride,omitempty"`
	OverrideReason         *string               `json:"overrideReason,omitempty"`
	Enterer                *Reference            `json:"enterer,omitempty"`
	EnteredDate            *string               `json:"enteredDate,omitempty"`
	Reason                 []CodeableConcept     `json:"reason,omitempty"`
	Service                []Reference           `json:"service,omitempty"`
	ProductReference       *Reference            `json:"productReference,omitempty"`
	ProductCodeableConcept *CodeableConcept      `json:"productCodeableConcept,omitempty"`
	Account                []Reference           `json:"account,omitempty"`
	Note                   []Annotation          `json:"note,omitempty"`
	SupportingInformation  []Reference           `json:"supportingInformation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ChargeItem
type ChargeItemPerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

type OtherChargeItem ChargeItem

// on convert struct to json, automatically add resourceType=ChargeItem
func (r ChargeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherChargeItem
		ResourceType string `json:"resourceType"`
	}{
		OtherChargeItem: OtherChargeItem(r),
		ResourceType:    "ChargeItem",
	})
}

func (resource *ChargeItem) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ChargeItem.Id", nil)
	}
	return StringInput("ChargeItem.Id", resource.Id)
}
func (resource *ChargeItem) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ChargeItem.ImplicitRules", nil)
	}
	return StringInput("ChargeItem.ImplicitRules", resource.ImplicitRules)
}
func (resource *ChargeItem) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ChargeItem.Language", nil, optionsValueSet)
	}
	return CodeSelect("ChargeItem.Language", resource.Language, optionsValueSet)
}
func (resource *ChargeItem) T_DefinitionUri(numDefinitionUri int) templ.Component {

	if resource == nil || len(resource.DefinitionUri) >= numDefinitionUri {
		return StringInput("ChargeItem.DefinitionUri["+strconv.Itoa(numDefinitionUri)+"]", nil)
	}
	return StringInput("ChargeItem.DefinitionUri["+strconv.Itoa(numDefinitionUri)+"]", &resource.DefinitionUri[numDefinitionUri])
}
func (resource *ChargeItem) T_DefinitionCanonical(numDefinitionCanonical int) templ.Component {

	if resource == nil || len(resource.DefinitionCanonical) >= numDefinitionCanonical {
		return StringInput("ChargeItem.DefinitionCanonical["+strconv.Itoa(numDefinitionCanonical)+"]", nil)
	}
	return StringInput("ChargeItem.DefinitionCanonical["+strconv.Itoa(numDefinitionCanonical)+"]", &resource.DefinitionCanonical[numDefinitionCanonical])
}
func (resource *ChargeItem) T_Status() templ.Component {
	optionsValueSet := VSChargeitem_status

	if resource == nil {
		return CodeSelect("ChargeItem.Status", nil, optionsValueSet)
	}
	return CodeSelect("ChargeItem.Status", &resource.Status, optionsValueSet)
}
func (resource *ChargeItem) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ChargeItem.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ChargeItem.Code", &resource.Code, optionsValueSet)
}
func (resource *ChargeItem) T_Bodysite(numBodysite int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Bodysite) >= numBodysite {
		return CodeableConceptSelect("ChargeItem.Bodysite["+strconv.Itoa(numBodysite)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ChargeItem.Bodysite["+strconv.Itoa(numBodysite)+"]", &resource.Bodysite[numBodysite], optionsValueSet)
}
func (resource *ChargeItem) T_FactorOverride() templ.Component {

	if resource == nil {
		return Float64Input("ChargeItem.FactorOverride", nil)
	}
	return Float64Input("ChargeItem.FactorOverride", resource.FactorOverride)
}
func (resource *ChargeItem) T_OverrideReason() templ.Component {

	if resource == nil {
		return StringInput("ChargeItem.OverrideReason", nil)
	}
	return StringInput("ChargeItem.OverrideReason", resource.OverrideReason)
}
func (resource *ChargeItem) T_EnteredDate() templ.Component {

	if resource == nil {
		return StringInput("ChargeItem.EnteredDate", nil)
	}
	return StringInput("ChargeItem.EnteredDate", resource.EnteredDate)
}
func (resource *ChargeItem) T_Reason(numReason int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Reason) >= numReason {
		return CodeableConceptSelect("ChargeItem.Reason["+strconv.Itoa(numReason)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ChargeItem.Reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], optionsValueSet)
}
func (resource *ChargeItem) T_PerformerId(numPerformer int) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return StringInput("ChargeItem.Performer["+strconv.Itoa(numPerformer)+"].Id", nil)
	}
	return StringInput("ChargeItem.Performer["+strconv.Itoa(numPerformer)+"].Id", resource.Performer[numPerformer].Id)
}
func (resource *ChargeItem) T_PerformerFunction(numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return CodeableConceptSelect("ChargeItem.Performer["+strconv.Itoa(numPerformer)+"].Function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ChargeItem.Performer["+strconv.Itoa(numPerformer)+"].Function", resource.Performer[numPerformer].Function, optionsValueSet)
}
