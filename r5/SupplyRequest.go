package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	OccurrenceDateTime *string                  `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                  `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing                  `json:"occurrenceTiming,omitempty"`
	AuthoredOn         *string                  `json:"authoredOn,omitempty"`
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

func (resource *SupplyRequest) T_Id() templ.Component {

	if resource == nil {
		return StringInput("SupplyRequest.Id", nil)
	}
	return StringInput("SupplyRequest.Id", resource.Id)
}
func (resource *SupplyRequest) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("SupplyRequest.ImplicitRules", nil)
	}
	return StringInput("SupplyRequest.ImplicitRules", resource.ImplicitRules)
}
func (resource *SupplyRequest) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("SupplyRequest.Language", nil, optionsValueSet)
	}
	return CodeSelect("SupplyRequest.Language", resource.Language, optionsValueSet)
}
func (resource *SupplyRequest) T_Status() templ.Component {
	optionsValueSet := VSSupplyrequest_status

	if resource == nil {
		return CodeSelect("SupplyRequest.Status", nil, optionsValueSet)
	}
	return CodeSelect("SupplyRequest.Status", resource.Status, optionsValueSet)
}
func (resource *SupplyRequest) T_Category(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SupplyRequest.Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SupplyRequest.Category", resource.Category, optionsValueSet)
}
func (resource *SupplyRequest) T_Priority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("SupplyRequest.Priority", nil, optionsValueSet)
	}
	return CodeSelect("SupplyRequest.Priority", resource.Priority, optionsValueSet)
}
func (resource *SupplyRequest) T_AuthoredOn() templ.Component {

	if resource == nil {
		return StringInput("SupplyRequest.AuthoredOn", nil)
	}
	return StringInput("SupplyRequest.AuthoredOn", resource.AuthoredOn)
}
func (resource *SupplyRequest) T_ParameterId(numParameter int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return StringInput("SupplyRequest.Parameter["+strconv.Itoa(numParameter)+"].Id", nil)
	}
	return StringInput("SupplyRequest.Parameter["+strconv.Itoa(numParameter)+"].Id", resource.Parameter[numParameter].Id)
}
func (resource *SupplyRequest) T_ParameterCode(numParameter int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return CodeableConceptSelect("SupplyRequest.Parameter["+strconv.Itoa(numParameter)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SupplyRequest.Parameter["+strconv.Itoa(numParameter)+"].Code", resource.Parameter[numParameter].Code, optionsValueSet)
}
