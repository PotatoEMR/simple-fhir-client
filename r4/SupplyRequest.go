package r4

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	OccurrenceDateTime  *string                  `json:"occurrenceDateTime"`
	OccurrencePeriod    *Period                  `json:"occurrencePeriod"`
	OccurrenceTiming    *Timing                  `json:"occurrenceTiming"`
	AuthoredOn          *string                  `json:"authoredOn,omitempty"`
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
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	ValueQuantity        *Quantity        `json:"valueQuantity"`
	ValueRange           *Range           `json:"valueRange"`
	ValueBoolean         *bool            `json:"valueBoolean"`
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
func (resource *SupplyRequest) SupplyRequestLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *SupplyRequest) SupplyRequestStatus() templ.Component {
	optionsValueSet := VSSupplyrequest_status
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *SupplyRequest) SupplyRequestPriority() templ.Component {
	optionsValueSet := VSRequest_priority
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Priority
	}
	return CodeSelect("priority", currentVal, optionsValueSet)
}
