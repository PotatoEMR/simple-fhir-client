package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/DeviceRequest
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
	Replaces              []Reference              `json:"replaces,omitempty"`
	GroupIdentifier       *Identifier              `json:"groupIdentifier,omitempty"`
	Status                *string                  `json:"status,omitempty"`
	Intent                string                   `json:"intent"`
	Priority              *string                  `json:"priority,omitempty"`
	DoNotPerform          *bool                    `json:"doNotPerform,omitempty"`
	Code                  CodeableReference        `json:"code"`
	Quantity              *int                     `json:"quantity,omitempty"`
	Parameter             []DeviceRequestParameter `json:"parameter,omitempty"`
	Subject               Reference                `json:"subject"`
	Encounter             *Reference               `json:"encounter,omitempty"`
	OccurrenceDateTime    *string                  `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *Period                  `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming      *Timing                  `json:"occurrenceTiming,omitempty"`
	AuthoredOn            *string                  `json:"authoredOn,omitempty"`
	Requester             *Reference               `json:"requester,omitempty"`
	Performer             *CodeableReference       `json:"performer,omitempty"`
	Reason                []CodeableReference      `json:"reason,omitempty"`
	AsNeeded              *bool                    `json:"asNeeded,omitempty"`
	AsNeededFor           *CodeableConcept         `json:"asNeededFor,omitempty"`
	Insurance             []Reference              `json:"insurance,omitempty"`
	SupportingInfo        []Reference              `json:"supportingInfo,omitempty"`
	Note                  []Annotation             `json:"note,omitempty"`
	RelevantHistory       []Reference              `json:"relevantHistory,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceRequest
type DeviceRequestParameter struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Code                 *CodeableConcept `json:"code,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range           `json:"valueRange,omitempty"`
	ValueBoolean         *bool            `json:"valueBoolean,omitempty"`
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

func (resource *DeviceRequest) T_Id() templ.Component {

	if resource == nil {
		return StringInput("DeviceRequest.Id", nil)
	}
	return StringInput("DeviceRequest.Id", resource.Id)
}
func (resource *DeviceRequest) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("DeviceRequest.ImplicitRules", nil)
	}
	return StringInput("DeviceRequest.ImplicitRules", resource.ImplicitRules)
}
func (resource *DeviceRequest) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("DeviceRequest.Language", nil, optionsValueSet)
	}
	return CodeSelect("DeviceRequest.Language", resource.Language, optionsValueSet)
}
func (resource *DeviceRequest) T_InstantiatesCanonical(numInstantiatesCanonical int) templ.Component {

	if resource == nil || len(resource.InstantiatesCanonical) >= numInstantiatesCanonical {
		return StringInput("DeviceRequest.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil)
	}
	return StringInput("DeviceRequest.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical])
}
func (resource *DeviceRequest) T_InstantiatesUri(numInstantiatesUri int) templ.Component {

	if resource == nil || len(resource.InstantiatesUri) >= numInstantiatesUri {
		return StringInput("DeviceRequest.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil)
	}
	return StringInput("DeviceRequest.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri])
}
func (resource *DeviceRequest) T_Status() templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("DeviceRequest.Status", nil, optionsValueSet)
	}
	return CodeSelect("DeviceRequest.Status", resource.Status, optionsValueSet)
}
func (resource *DeviceRequest) T_Intent() templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("DeviceRequest.Intent", nil, optionsValueSet)
	}
	return CodeSelect("DeviceRequest.Intent", &resource.Intent, optionsValueSet)
}
func (resource *DeviceRequest) T_Priority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("DeviceRequest.Priority", nil, optionsValueSet)
	}
	return CodeSelect("DeviceRequest.Priority", resource.Priority, optionsValueSet)
}
func (resource *DeviceRequest) T_DoNotPerform() templ.Component {

	if resource == nil {
		return BoolInput("DeviceRequest.DoNotPerform", nil)
	}
	return BoolInput("DeviceRequest.DoNotPerform", resource.DoNotPerform)
}
func (resource *DeviceRequest) T_Quantity() templ.Component {

	if resource == nil {
		return IntInput("DeviceRequest.Quantity", nil)
	}
	return IntInput("DeviceRequest.Quantity", resource.Quantity)
}
func (resource *DeviceRequest) T_AuthoredOn() templ.Component {

	if resource == nil {
		return StringInput("DeviceRequest.AuthoredOn", nil)
	}
	return StringInput("DeviceRequest.AuthoredOn", resource.AuthoredOn)
}
func (resource *DeviceRequest) T_AsNeeded() templ.Component {

	if resource == nil {
		return BoolInput("DeviceRequest.AsNeeded", nil)
	}
	return BoolInput("DeviceRequest.AsNeeded", resource.AsNeeded)
}
func (resource *DeviceRequest) T_AsNeededFor(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DeviceRequest.AsNeededFor", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceRequest.AsNeededFor", resource.AsNeededFor, optionsValueSet)
}
func (resource *DeviceRequest) T_ParameterId(numParameter int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return StringInput("DeviceRequest.Parameter["+strconv.Itoa(numParameter)+"].Id", nil)
	}
	return StringInput("DeviceRequest.Parameter["+strconv.Itoa(numParameter)+"].Id", resource.Parameter[numParameter].Id)
}
func (resource *DeviceRequest) T_ParameterCode(numParameter int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return CodeableConceptSelect("DeviceRequest.Parameter["+strconv.Itoa(numParameter)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceRequest.Parameter["+strconv.Itoa(numParameter)+"].Code", resource.Parameter[numParameter].Code, optionsValueSet)
}
