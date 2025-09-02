package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/DeviceRequest
type DeviceRequest struct {
	Id                    *string                  `json:"id,omitempty"`
	Meta                  *Meta                    `json:"meta,omitempty"`
	ImplicitRules         *string                  `json:"implicitRules,omitempty"`
	Language              *string                  `json:"language,omitempty"`
	Text                  *Narrative               `json:"text,omitempty"`
	Contained             []Resource               `json:"contained,omitempty"`
	Extension             []Extension              `json:"extension,omitempty"`
	ModifierExtension     []Extension              `json:"modifierExtension,omitempty"`
	Identifier            []Identifier             `json:"identifier,omitempty"`
	InstantiatesCanonical []string                 `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                 `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference              `json:"basedOn,omitempty"`
	PriorRequest          []Reference              `json:"priorRequest,omitempty"`
	GroupIdentifier       *Identifier              `json:"groupIdentifier,omitempty"`
	Status                *string                  `json:"status,omitempty"`
	Intent                string                   `json:"intent"`
	Priority              *string                  `json:"priority,omitempty"`
	CodeReference         Reference                `json:"codeReference"`
	CodeCodeableConcept   CodeableConcept          `json:"codeCodeableConcept"`
	Parameter             []DeviceRequestParameter `json:"parameter,omitempty"`
	Subject               Reference                `json:"subject"`
	Encounter             *Reference               `json:"encounter,omitempty"`
	OccurrenceDateTime    *string                  `json:"occurrenceDateTime"`
	OccurrencePeriod      *Period                  `json:"occurrencePeriod"`
	OccurrenceTiming      *Timing                  `json:"occurrenceTiming"`
	AuthoredOn            *string                  `json:"authoredOn,omitempty"`
	Requester             *Reference               `json:"requester,omitempty"`
	PerformerType         *CodeableConcept         `json:"performerType,omitempty"`
	Performer             *Reference               `json:"performer,omitempty"`
	ReasonCode            []CodeableConcept        `json:"reasonCode,omitempty"`
	ReasonReference       []Reference              `json:"reasonReference,omitempty"`
	Insurance             []Reference              `json:"insurance,omitempty"`
	SupportingInfo        []Reference              `json:"supportingInfo,omitempty"`
	Note                  []Annotation             `json:"note,omitempty"`
	RelevantHistory       []Reference              `json:"relevantHistory,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/DeviceRequest
type DeviceRequestParameter struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Code                 *CodeableConcept `json:"code,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	ValueQuantity        *Quantity        `json:"valueQuantity"`
	ValueRange           *Range           `json:"valueRange"`
	ValueBoolean         *bool            `json:"valueBoolean"`
}

type OtherDeviceRequest DeviceRequest

// on convert struct to json, automatically add resourceType=DeviceRequest
func (r DeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceRequest: OtherDeviceRequest(r),
		ResourceType:       "DeviceRequest",
	})
}

func (resource *DeviceRequest) DeviceRequestLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *DeviceRequest) DeviceRequestStatus() templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", resource.Status, optionsValueSet)
}
func (resource *DeviceRequest) DeviceRequestIntent() templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("intent", nil, optionsValueSet)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet)
}
func (resource *DeviceRequest) DeviceRequestPriority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet)
}
func (resource *DeviceRequest) DeviceRequestPerformerType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("performerType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("performerType", resource.PerformerType, optionsValueSet)
}
func (resource *DeviceRequest) DeviceRequestReasonCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("reasonCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reasonCode", &resource.ReasonCode[0], optionsValueSet)
}
func (resource *DeviceRequest) DeviceRequestParameterCode(numParameter int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Parameter) >= numParameter {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Parameter[numParameter].Code, optionsValueSet)
}
