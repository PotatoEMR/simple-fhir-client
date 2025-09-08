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
	OccurrenceDateTime    *time.Time               `json:"occurrenceDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	OccurrencePeriod      *Period                  `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming      *Timing                  `json:"occurrenceTiming,omitempty"`
	AuthoredOn            *time.Time               `json:"authoredOn,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r DeviceRequest) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "DeviceRequest/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "DeviceRequest"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *DeviceRequest) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs string) templ.Component {

	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("DeviceRequest.InstantiatesCanonical."+strconv.Itoa(numInstantiatesCanonical)+".", nil, htmlAttrs)
	}
	return StringInput("DeviceRequest.InstantiatesCanonical."+strconv.Itoa(numInstantiatesCanonical)+".", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *DeviceRequest) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs string) templ.Component {

	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("DeviceRequest.InstantiatesUri."+strconv.Itoa(numInstantiatesUri)+".", nil, htmlAttrs)
	}
	return StringInput("DeviceRequest.InstantiatesUri."+strconv.Itoa(numInstantiatesUri)+".", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *DeviceRequest) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("DeviceRequest.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DeviceRequest.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DeviceRequest) T_Intent(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("DeviceRequest.Intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DeviceRequest.Intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *DeviceRequest) T_Priority(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("DeviceRequest.Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DeviceRequest.Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *DeviceRequest) T_DoNotPerform(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("DeviceRequest.DoNotPerform", nil, htmlAttrs)
	}
	return BoolInput("DeviceRequest.DoNotPerform", resource.DoNotPerform, htmlAttrs)
}
func (resource *DeviceRequest) T_Quantity(htmlAttrs string) templ.Component {

	if resource == nil {
		return IntInput("DeviceRequest.Quantity", nil, htmlAttrs)
	}
	return IntInput("DeviceRequest.Quantity", resource.Quantity, htmlAttrs)
}
func (resource *DeviceRequest) T_OccurrenceDateTime(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("DeviceRequest.OccurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("DeviceRequest.OccurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *DeviceRequest) T_AuthoredOn(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("DeviceRequest.AuthoredOn", nil, htmlAttrs)
	}
	return DateTimeInput("DeviceRequest.AuthoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *DeviceRequest) T_AsNeeded(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("DeviceRequest.AsNeeded", nil, htmlAttrs)
	}
	return BoolInput("DeviceRequest.AsNeeded", resource.AsNeeded, htmlAttrs)
}
func (resource *DeviceRequest) T_AsNeededFor(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DeviceRequest.AsNeededFor", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceRequest.AsNeededFor", resource.AsNeededFor, optionsValueSet, htmlAttrs)
}
func (resource *DeviceRequest) T_Note(numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("DeviceRequest.Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("DeviceRequest.Note."+strconv.Itoa(numNote)+".", &resource.Note[numNote], htmlAttrs)
}
func (resource *DeviceRequest) T_ParameterCode(numParameter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeableConceptSelect("DeviceRequest.Parameter."+strconv.Itoa(numParameter)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceRequest.Parameter."+strconv.Itoa(numParameter)+"..Code", resource.Parameter[numParameter].Code, optionsValueSet, htmlAttrs)
}
func (resource *DeviceRequest) T_ParameterValueCodeableConcept(numParameter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeableConceptSelect("DeviceRequest.Parameter."+strconv.Itoa(numParameter)+"..ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceRequest.Parameter."+strconv.Itoa(numParameter)+"..ValueCodeableConcept", resource.Parameter[numParameter].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *DeviceRequest) T_ParameterValueBoolean(numParameter int, htmlAttrs string) templ.Component {

	if resource == nil || numParameter >= len(resource.Parameter) {
		return BoolInput("DeviceRequest.Parameter."+strconv.Itoa(numParameter)+"..ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("DeviceRequest.Parameter."+strconv.Itoa(numParameter)+"..ValueBoolean", resource.Parameter[numParameter].ValueBoolean, htmlAttrs)
}
