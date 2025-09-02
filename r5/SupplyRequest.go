package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	OccurrenceDateTime *string                  `json:"occurrenceDateTime"`
	OccurrencePeriod   *Period                  `json:"occurrencePeriod"`
	OccurrenceTiming   *Timing                  `json:"occurrenceTiming"`
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

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *SupplyRequest) SupplyRequestStatus() templ.Component {
	optionsValueSet := VSSupplyrequest_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", resource.Status, optionsValueSet)
}
func (resource *SupplyRequest) SupplyRequestCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.Category, optionsValueSet)
}
func (resource *SupplyRequest) SupplyRequestPriority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource != nil {
		return CodeSelect("priority", nil, optionsValueSet)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet)
}
func (resource *SupplyRequest) SupplyRequestParameterCode(numParameter int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Parameter) >= numParameter {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Parameter[numParameter].Code, optionsValueSet)
}
