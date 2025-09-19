package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/DeviceRequest
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
	OccurrenceDateTime    *FhirDateTime            `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *Period                  `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming      *Timing                  `json:"occurrenceTiming,omitempty"`
	AuthoredOn            *FhirDateTime            `json:"authoredOn,omitempty"`
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

// http://hl7.org/fhir/r4/StructureDefinition/DeviceRequest
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
func (resource *DeviceRequest) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *DeviceRequest) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *DeviceRequest) T_BasedOn(numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput("basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *DeviceRequest) T_PriorRequest(numPriorRequest int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPriorRequest >= len(resource.PriorRequest) {
		return ReferenceInput("priorRequest["+strconv.Itoa(numPriorRequest)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("priorRequest["+strconv.Itoa(numPriorRequest)+"]", &resource.PriorRequest[numPriorRequest], htmlAttrs)
}
func (resource *DeviceRequest) T_GroupIdentifier(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IdentifierInput("groupIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("groupIdentifier", resource.GroupIdentifier, htmlAttrs)
}
func (resource *DeviceRequest) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DeviceRequest) T_Intent(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *DeviceRequest) T_Priority(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *DeviceRequest) T_CodeReference(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("codeReference", nil, htmlAttrs)
	}
	return ReferenceInput("codeReference", &resource.CodeReference, htmlAttrs)
}
func (resource *DeviceRequest) T_CodeCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("codeCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("codeCodeableConcept", &resource.CodeCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *DeviceRequest) T_Subject(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("subject", nil, htmlAttrs)
	}
	return ReferenceInput("subject", &resource.Subject, htmlAttrs)
}
func (resource *DeviceRequest) T_Encounter(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("encounter", nil, htmlAttrs)
	}
	return ReferenceInput("encounter", resource.Encounter, htmlAttrs)
}
func (resource *DeviceRequest) T_OccurrenceDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("occurrenceDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("occurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *DeviceRequest) T_OccurrencePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("occurrencePeriod", nil, htmlAttrs)
	}
	return PeriodInput("occurrencePeriod", resource.OccurrencePeriod, htmlAttrs)
}
func (resource *DeviceRequest) T_OccurrenceTiming(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return TimingInput("occurrenceTiming", nil, htmlAttrs)
	}
	return TimingInput("occurrenceTiming", resource.OccurrenceTiming, htmlAttrs)
}
func (resource *DeviceRequest) T_AuthoredOn(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("authoredOn", nil, htmlAttrs)
	}
	return FhirDateTimeInput("authoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *DeviceRequest) T_Requester(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("requester", nil, htmlAttrs)
	}
	return ReferenceInput("requester", resource.Requester, htmlAttrs)
}
func (resource *DeviceRequest) T_PerformerType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("performerType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performerType", resource.PerformerType, optionsValueSet, htmlAttrs)
}
func (resource *DeviceRequest) T_Performer(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("performer", nil, htmlAttrs)
	}
	return ReferenceInput("performer", resource.Performer, htmlAttrs)
}
func (resource *DeviceRequest) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *DeviceRequest) T_ReasonReference(numReasonReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonReference >= len(resource.ReasonReference) {
		return ReferenceInput("reasonReference["+strconv.Itoa(numReasonReference)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("reasonReference["+strconv.Itoa(numReasonReference)+"]", &resource.ReasonReference[numReasonReference], htmlAttrs)
}
func (resource *DeviceRequest) T_Insurance(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return ReferenceInput("insurance["+strconv.Itoa(numInsurance)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("insurance["+strconv.Itoa(numInsurance)+"]", &resource.Insurance[numInsurance], htmlAttrs)
}
func (resource *DeviceRequest) T_SupportingInfo(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return ReferenceInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"]", &resource.SupportingInfo[numSupportingInfo], htmlAttrs)
}
func (resource *DeviceRequest) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *DeviceRequest) T_RelevantHistory(numRelevantHistory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelevantHistory >= len(resource.RelevantHistory) {
		return ReferenceInput("relevantHistory["+strconv.Itoa(numRelevantHistory)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("relevantHistory["+strconv.Itoa(numRelevantHistory)+"]", &resource.RelevantHistory[numRelevantHistory], htmlAttrs)
}
func (resource *DeviceRequest) T_ParameterCode(numParameter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeableConceptSelect("parameter["+strconv.Itoa(numParameter)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("parameter["+strconv.Itoa(numParameter)+"].code", resource.Parameter[numParameter].Code, optionsValueSet, htmlAttrs)
}
func (resource *DeviceRequest) T_ParameterValueCodeableConcept(numParameter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeableConceptSelect("parameter["+strconv.Itoa(numParameter)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("parameter["+strconv.Itoa(numParameter)+"].valueCodeableConcept", resource.Parameter[numParameter].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *DeviceRequest) T_ParameterValueQuantity(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return QuantityInput("parameter["+strconv.Itoa(numParameter)+"].valueQuantity", nil, htmlAttrs)
	}
	return QuantityInput("parameter["+strconv.Itoa(numParameter)+"].valueQuantity", resource.Parameter[numParameter].ValueQuantity, htmlAttrs)
}
func (resource *DeviceRequest) T_ParameterValueRange(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return RangeInput("parameter["+strconv.Itoa(numParameter)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("parameter["+strconv.Itoa(numParameter)+"].valueRange", resource.Parameter[numParameter].ValueRange, htmlAttrs)
}
func (resource *DeviceRequest) T_ParameterValueBoolean(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return BoolInput("parameter["+strconv.Itoa(numParameter)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("parameter["+strconv.Itoa(numParameter)+"].valueBoolean", resource.Parameter[numParameter].ValueBoolean, htmlAttrs)
}
