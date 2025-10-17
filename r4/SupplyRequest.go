package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/SupplyRequest
type SupplyRequest struct {
	Id                  *string                  `json:"id,omitempty"`
	Meta                *Meta                    `json:"meta,omitempty"`
	ImplicitRules       *string                  `json:"implicitRules,omitempty"`
	Language            *string                  `json:"language,omitempty"`
	Text                *Narrative               `json:"text,omitempty"`
	Contained           []Resource               `json:"contained,omitempty"`
	Extension           []Extension              `json:"extension,omitempty"`
	ModifierExtension   []Extension              `json:"modifierExtension,omitempty"`
	Identifier          []Identifier             `json:"identifier,omitempty"`
	Status              *string                  `json:"status,omitempty"`
	Category            *CodeableConcept         `json:"category,omitempty"`
	Priority            *string                  `json:"priority,omitempty"`
	ItemCodeableConcept CodeableConcept          `json:"itemCodeableConcept"`
	ItemReference       Reference                `json:"itemReference"`
	Quantity            Quantity                 `json:"quantity"`
	Parameter           []SupplyRequestParameter `json:"parameter,omitempty"`
	OccurrenceDateTime  *FhirDateTime            `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod    *Period                  `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming    *Timing                  `json:"occurrenceTiming,omitempty"`
	AuthoredOn          *FhirDateTime            `json:"authoredOn,omitempty"`
	Requester           *Reference               `json:"requester,omitempty"`
	Supplier            []Reference              `json:"supplier,omitempty"`
	ReasonCode          []CodeableConcept        `json:"reasonCode,omitempty"`
	ReasonReference     []Reference              `json:"reasonReference,omitempty"`
	DeliverFrom         *Reference               `json:"deliverFrom,omitempty"`
	DeliverTo           *Reference               `json:"deliverTo,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SupplyRequest
type SupplyRequestParameter struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Code                 *CodeableConcept `json:"code,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range           `json:"valueRange,omitempty"`
	ValueBoolean         *bool            `json:"valueBoolean,omitempty"`
}

type OtherSupplyRequest SupplyRequest

// struct -> json, automatically add resourceType=Patient
func (r SupplyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSupplyRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherSupplyRequest: OtherSupplyRequest(r),
		ResourceType:       "SupplyRequest",
	})
}

// json -> struct, first reject if resourceType != SupplyRequest
func (r *SupplyRequest) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "SupplyRequest" {
		return errors.New("resourceType not SupplyRequest")
	}
	return json.Unmarshal(data, (*OtherSupplyRequest)(r))
}

func (r SupplyRequest) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "SupplyRequest/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "SupplyRequest"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *SupplyRequest) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSupplyrequest_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *SupplyRequest) T_Category(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category", resource.Category, optionsValueSet, htmlAttrs)
}
func (resource *SupplyRequest) T_Priority(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *SupplyRequest) T_ItemCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("itemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("itemCodeableConcept", &resource.ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *SupplyRequest) T_ItemReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "itemReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "itemReference", &resource.ItemReference, htmlAttrs)
}
func (resource *SupplyRequest) T_Quantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("quantity", &resource.Quantity, optionsValueSet, htmlAttrs)
}
func (resource *SupplyRequest) T_OccurrenceDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("occurrenceDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("occurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *SupplyRequest) T_OccurrencePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("occurrencePeriod", nil, htmlAttrs)
	}
	return PeriodInput("occurrencePeriod", resource.OccurrencePeriod, htmlAttrs)
}
func (resource *SupplyRequest) T_OccurrenceTiming(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return TimingInput("occurrenceTiming", nil, htmlAttrs)
	}
	return TimingInput("occurrenceTiming", resource.OccurrenceTiming, htmlAttrs)
}
func (resource *SupplyRequest) T_AuthoredOn(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("authoredOn", nil, htmlAttrs)
	}
	return FhirDateTimeInput("authoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *SupplyRequest) T_Requester(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "requester", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "requester", resource.Requester, htmlAttrs)
}
func (resource *SupplyRequest) T_Supplier(frs []FhirResource, numSupplier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupplier >= len(resource.Supplier) {
		return ReferenceInput(frs, "supplier["+strconv.Itoa(numSupplier)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "supplier["+strconv.Itoa(numSupplier)+"]", &resource.Supplier[numSupplier], htmlAttrs)
}
func (resource *SupplyRequest) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *SupplyRequest) T_ReasonReference(frs []FhirResource, numReasonReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonReference >= len(resource.ReasonReference) {
		return ReferenceInput(frs, "reasonReference["+strconv.Itoa(numReasonReference)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "reasonReference["+strconv.Itoa(numReasonReference)+"]", &resource.ReasonReference[numReasonReference], htmlAttrs)
}
func (resource *SupplyRequest) T_DeliverFrom(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "deliverFrom", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "deliverFrom", resource.DeliverFrom, htmlAttrs)
}
func (resource *SupplyRequest) T_DeliverTo(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "deliverTo", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "deliverTo", resource.DeliverTo, htmlAttrs)
}
func (resource *SupplyRequest) T_ParameterCode(numParameter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeableConceptSelect("parameter["+strconv.Itoa(numParameter)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("parameter["+strconv.Itoa(numParameter)+"].code", resource.Parameter[numParameter].Code, optionsValueSet, htmlAttrs)
}
func (resource *SupplyRequest) T_ParameterValueCodeableConcept(numParameter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeableConceptSelect("parameter["+strconv.Itoa(numParameter)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("parameter["+strconv.Itoa(numParameter)+"].valueCodeableConcept", resource.Parameter[numParameter].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *SupplyRequest) T_ParameterValueQuantity(numParameter int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return QuantityInput("parameter["+strconv.Itoa(numParameter)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("parameter["+strconv.Itoa(numParameter)+"].valueQuantity", resource.Parameter[numParameter].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *SupplyRequest) T_ParameterValueRange(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return RangeInput("parameter["+strconv.Itoa(numParameter)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("parameter["+strconv.Itoa(numParameter)+"].valueRange", resource.Parameter[numParameter].ValueRange, htmlAttrs)
}
func (resource *SupplyRequest) T_ParameterValueBoolean(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return BoolInput("parameter["+strconv.Itoa(numParameter)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("parameter["+strconv.Itoa(numParameter)+"].valueBoolean", resource.Parameter[numParameter].ValueBoolean, htmlAttrs)
}
