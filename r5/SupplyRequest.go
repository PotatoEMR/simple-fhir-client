package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/SupplyRequest
type SupplyRequest struct {
	Id                 *string                  `json:"id,omitempty"`
	Meta               *Meta                    `json:"meta,omitempty"`
	ImplicitRules      *string                  `json:"implicitRules,omitempty"`
	Language           *string                  `json:"language,omitempty"`
	Text               *Narrative               `json:"text,omitempty"`
	Contained          []Resource               `json:"contained,omitempty"`
	Extension          []Extension              `json:"extension,omitempty"`
	ModifierExtension  []Extension              `json:"modifierExtension,omitempty"`
	Identifier         []Identifier             `json:"identifier,omitempty"`
	Status             *string                  `json:"status,omitempty"`
	BasedOn            []Reference              `json:"basedOn,omitempty"`
	Category           *CodeableConcept         `json:"category,omitempty"`
	Priority           *string                  `json:"priority,omitempty"`
	DeliverFor         *Reference               `json:"deliverFor,omitempty"`
	Item               CodeableReference        `json:"item"`
	Quantity           Quantity                 `json:"quantity"`
	Parameter          []SupplyRequestParameter `json:"parameter,omitempty"`
	OccurrenceDateTime *time.Time               `json:"occurrenceDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	OccurrencePeriod   *Period                  `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing                  `json:"occurrenceTiming,omitempty"`
	AuthoredOn         *time.Time               `json:"authoredOn,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Requester          *Reference               `json:"requester,omitempty"`
	Supplier           []Reference              `json:"supplier,omitempty"`
	Reason             []CodeableReference      `json:"reason,omitempty"`
	DeliverFrom        *Reference               `json:"deliverFrom,omitempty"`
	DeliverTo          *Reference               `json:"deliverTo,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SupplyRequest
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

// on convert struct to json, automatically add resourceType=SupplyRequest
func (r SupplyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSupplyRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherSupplyRequest: OtherSupplyRequest(r),
		ResourceType:       "SupplyRequest",
	})
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
func (resource *SupplyRequest) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSSupplyrequest_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *SupplyRequest) T_Category(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category", resource.Category, optionsValueSet, htmlAttrs)
}
func (resource *SupplyRequest) T_Priority(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *SupplyRequest) T_OccurrenceDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("OccurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("OccurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *SupplyRequest) T_AuthoredOn(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("AuthoredOn", nil, htmlAttrs)
	}
	return DateTimeInput("AuthoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *SupplyRequest) T_ParameterCode(numParameter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeableConceptSelect("Parameter["+strconv.Itoa(numParameter)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Parameter["+strconv.Itoa(numParameter)+"]Code", resource.Parameter[numParameter].Code, optionsValueSet, htmlAttrs)
}
func (resource *SupplyRequest) T_ParameterValueCodeableConcept(numParameter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeableConceptSelect("Parameter["+strconv.Itoa(numParameter)+"]ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Parameter["+strconv.Itoa(numParameter)+"]ValueCodeableConcept", resource.Parameter[numParameter].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *SupplyRequest) T_ParameterValueBoolean(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return BoolInput("Parameter["+strconv.Itoa(numParameter)+"]ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Parameter["+strconv.Itoa(numParameter)+"]ValueBoolean", resource.Parameter[numParameter].ValueBoolean, htmlAttrs)
}
