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
	OccurrenceDateTime     *time.Time            `json:"occurrenceDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	EnteredDate            *time.Time            `json:"enteredDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r ChargeItem) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ChargeItem/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ChargeItem"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ChargeItem) T_DefinitionUri(numDefinitionUri int, htmlAttrs string) templ.Component {
	if resource == nil || numDefinitionUri >= len(resource.DefinitionUri) {
		return StringInput("DefinitionUri["+strconv.Itoa(numDefinitionUri)+"]", nil, htmlAttrs)
	}
	return StringInput("DefinitionUri["+strconv.Itoa(numDefinitionUri)+"]", &resource.DefinitionUri[numDefinitionUri], htmlAttrs)
}
func (resource *ChargeItem) T_DefinitionCanonical(numDefinitionCanonical int, htmlAttrs string) templ.Component {
	if resource == nil || numDefinitionCanonical >= len(resource.DefinitionCanonical) {
		return StringInput("DefinitionCanonical["+strconv.Itoa(numDefinitionCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("DefinitionCanonical["+strconv.Itoa(numDefinitionCanonical)+"]", &resource.DefinitionCanonical[numDefinitionCanonical], htmlAttrs)
}
func (resource *ChargeItem) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSChargeitem_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItem) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItem) T_OccurrenceDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("OccurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("OccurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *ChargeItem) T_Bodysite(numBodysite int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBodysite >= len(resource.Bodysite) {
		return CodeableConceptSelect("Bodysite["+strconv.Itoa(numBodysite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Bodysite["+strconv.Itoa(numBodysite)+"]", &resource.Bodysite[numBodysite], optionsValueSet, htmlAttrs)
}
func (resource *ChargeItem) T_FactorOverride(htmlAttrs string) templ.Component {
	if resource == nil {
		return Float64Input("FactorOverride", nil, htmlAttrs)
	}
	return Float64Input("FactorOverride", resource.FactorOverride, htmlAttrs)
}
func (resource *ChargeItem) T_OverrideReason(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("OverrideReason", nil, htmlAttrs)
	}
	return StringInput("OverrideReason", resource.OverrideReason, htmlAttrs)
}
func (resource *ChargeItem) T_EnteredDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("EnteredDate", nil, htmlAttrs)
	}
	return DateTimeInput("EnteredDate", resource.EnteredDate, htmlAttrs)
}
func (resource *ChargeItem) T_Reason(numReason int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableConceptSelect("Reason["+strconv.Itoa(numReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], optionsValueSet, htmlAttrs)
}
func (resource *ChargeItem) T_ProductCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ProductCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ProductCodeableConcept", resource.ProductCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItem) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *ChargeItem) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("Performer["+strconv.Itoa(numPerformer)+"]Function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Performer["+strconv.Itoa(numPerformer)+"]Function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
