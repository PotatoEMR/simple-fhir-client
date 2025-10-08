package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ChargeItem
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
	Encounter              *Reference            `json:"encounter,omitempty"`
	OccurrenceDateTime     *FhirDateTime         `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod       *Period               `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming       *Timing               `json:"occurrenceTiming,omitempty"`
	Performer              []ChargeItemPerformer `json:"performer,omitempty"`
	PerformingOrganization *Reference            `json:"performingOrganization,omitempty"`
	RequestingOrganization *Reference            `json:"requestingOrganization,omitempty"`
	CostCenter             *Reference            `json:"costCenter,omitempty"`
	Quantity               *Quantity             `json:"quantity,omitempty"`
	Bodysite               []CodeableConcept     `json:"bodysite,omitempty"`
	UnitPriceComponent     *MonetaryComponent    `json:"unitPriceComponent,omitempty"`
	TotalPriceComponent    *MonetaryComponent    `json:"totalPriceComponent,omitempty"`
	OverrideReason         *CodeableConcept      `json:"overrideReason,omitempty"`
	Enterer                *Reference            `json:"enterer,omitempty"`
	EnteredDate            *FhirDateTime         `json:"enteredDate,omitempty"`
	Reason                 []CodeableConcept     `json:"reason,omitempty"`
	Service                []CodeableReference   `json:"service,omitempty"`
	Product                []CodeableReference   `json:"product,omitempty"`
	Account                []Reference           `json:"account,omitempty"`
	Note                   []Annotation          `json:"note,omitempty"`
	SupportingInformation  []Reference           `json:"supportingInformation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ChargeItem
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
func (resource *ChargeItem) T_DefinitionUri(numDefinitionUri int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDefinitionUri >= len(resource.DefinitionUri) {
		return StringInput("definitionUri["+strconv.Itoa(numDefinitionUri)+"]", nil, htmlAttrs)
	}
	return StringInput("definitionUri["+strconv.Itoa(numDefinitionUri)+"]", &resource.DefinitionUri[numDefinitionUri], htmlAttrs)
}
func (resource *ChargeItem) T_DefinitionCanonical(numDefinitionCanonical int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDefinitionCanonical >= len(resource.DefinitionCanonical) {
		return StringInput("definitionCanonical["+strconv.Itoa(numDefinitionCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("definitionCanonical["+strconv.Itoa(numDefinitionCanonical)+"]", &resource.DefinitionCanonical[numDefinitionCanonical], htmlAttrs)
}
func (resource *ChargeItem) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSChargeitem_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItem) T_PartOf(frs []FhirResource, numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *ChargeItem) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItem) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *ChargeItem) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *ChargeItem) T_OccurrenceDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("occurrenceDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("occurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *ChargeItem) T_OccurrencePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("occurrencePeriod", nil, htmlAttrs)
	}
	return PeriodInput("occurrencePeriod", resource.OccurrencePeriod, htmlAttrs)
}
func (resource *ChargeItem) T_OccurrenceTiming(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return TimingInput("occurrenceTiming", nil, htmlAttrs)
	}
	return TimingInput("occurrenceTiming", resource.OccurrenceTiming, htmlAttrs)
}
func (resource *ChargeItem) T_PerformingOrganization(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "performingOrganization", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "performingOrganization", resource.PerformingOrganization, htmlAttrs)
}
func (resource *ChargeItem) T_RequestingOrganization(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "requestingOrganization", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "requestingOrganization", resource.RequestingOrganization, htmlAttrs)
}
func (resource *ChargeItem) T_CostCenter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "costCenter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "costCenter", resource.CostCenter, htmlAttrs)
}
func (resource *ChargeItem) T_Quantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("quantity", resource.Quantity, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItem) T_Bodysite(numBodysite int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBodysite >= len(resource.Bodysite) {
		return CodeableConceptSelect("bodysite["+strconv.Itoa(numBodysite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("bodysite["+strconv.Itoa(numBodysite)+"]", &resource.Bodysite[numBodysite], optionsValueSet, htmlAttrs)
}
func (resource *ChargeItem) T_UnitPriceComponent(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return MonetaryComponentInput("unitPriceComponent", nil, htmlAttrs)
	}
	return MonetaryComponentInput("unitPriceComponent", resource.UnitPriceComponent, htmlAttrs)
}
func (resource *ChargeItem) T_TotalPriceComponent(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return MonetaryComponentInput("totalPriceComponent", nil, htmlAttrs)
	}
	return MonetaryComponentInput("totalPriceComponent", resource.TotalPriceComponent, htmlAttrs)
}
func (resource *ChargeItem) T_OverrideReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("overrideReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("overrideReason", resource.OverrideReason, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItem) T_Enterer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "enterer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "enterer", resource.Enterer, htmlAttrs)
}
func (resource *ChargeItem) T_EnteredDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("enteredDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("enteredDate", resource.EnteredDate, htmlAttrs)
}
func (resource *ChargeItem) T_Reason(numReason int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableConceptSelect("reason["+strconv.Itoa(numReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], optionsValueSet, htmlAttrs)
}
func (resource *ChargeItem) T_Service(numService int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numService >= len(resource.Service) {
		return CodeableReferenceInput("service["+strconv.Itoa(numService)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("service["+strconv.Itoa(numService)+"]", &resource.Service[numService], htmlAttrs)
}
func (resource *ChargeItem) T_Product(numProduct int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProduct >= len(resource.Product) {
		return CodeableReferenceInput("product["+strconv.Itoa(numProduct)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("product["+strconv.Itoa(numProduct)+"]", &resource.Product[numProduct], htmlAttrs)
}
func (resource *ChargeItem) T_Account(frs []FhirResource, numAccount int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAccount >= len(resource.Account) {
		return ReferenceInput(frs, "account["+strconv.Itoa(numAccount)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "account["+strconv.Itoa(numAccount)+"]", &resource.Account[numAccount], htmlAttrs)
}
func (resource *ChargeItem) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *ChargeItem) T_SupportingInformation(frs []FhirResource, numSupportingInformation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInformation >= len(resource.SupportingInformation) {
		return ReferenceInput(frs, "supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", &resource.SupportingInformation[numSupportingInformation], htmlAttrs)
}
func (resource *ChargeItem) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
func (resource *ChargeItem) T_PerformerActor(frs []FhirResource, numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"].actor", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"].actor", &resource.Performer[numPerformer].Actor, htmlAttrs)
}
