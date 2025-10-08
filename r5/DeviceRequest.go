package r5

//generated with command go run ./bultaoreune
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
	OccurrenceDateTime    *FhirDateTime            `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *Period                  `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming      *Timing                  `json:"occurrenceTiming,omitempty"`
	AuthoredOn            *FhirDateTime            `json:"authoredOn,omitempty"`
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
func (resource *DeviceRequest) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *DeviceRequest) T_Replaces(frs []FhirResource, numReplaces int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReplaces >= len(resource.Replaces) {
		return ReferenceInput(frs, "replaces["+strconv.Itoa(numReplaces)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "replaces["+strconv.Itoa(numReplaces)+"]", &resource.Replaces[numReplaces], htmlAttrs)
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
func (resource *DeviceRequest) T_DoNotPerform(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("doNotPerform", nil, htmlAttrs)
	}
	return BoolInput("doNotPerform", resource.DoNotPerform, htmlAttrs)
}
func (resource *DeviceRequest) T_Code(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("code", nil, htmlAttrs)
	}
	return CodeableReferenceInput("code", &resource.Code, htmlAttrs)
}
func (resource *DeviceRequest) T_Quantity(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("quantity", nil, htmlAttrs)
	}
	return IntInput("quantity", resource.Quantity, htmlAttrs)
}
func (resource *DeviceRequest) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *DeviceRequest) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
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
func (resource *DeviceRequest) T_Requester(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "requester", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "requester", resource.Requester, htmlAttrs)
}
func (resource *DeviceRequest) T_Performer(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("performer", nil, htmlAttrs)
	}
	return CodeableReferenceInput("performer", resource.Performer, htmlAttrs)
}
func (resource *DeviceRequest) T_Reason(numReason int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], htmlAttrs)
}
func (resource *DeviceRequest) T_AsNeeded(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("asNeeded", nil, htmlAttrs)
	}
	return BoolInput("asNeeded", resource.AsNeeded, htmlAttrs)
}
func (resource *DeviceRequest) T_AsNeededFor(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("asNeededFor", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("asNeededFor", resource.AsNeededFor, optionsValueSet, htmlAttrs)
}
func (resource *DeviceRequest) T_Insurance(frs []FhirResource, numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return ReferenceInput(frs, "insurance["+strconv.Itoa(numInsurance)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "insurance["+strconv.Itoa(numInsurance)+"]", &resource.Insurance[numInsurance], htmlAttrs)
}
func (resource *DeviceRequest) T_SupportingInfo(frs []FhirResource, numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return ReferenceInput(frs, "supportingInfo["+strconv.Itoa(numSupportingInfo)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "supportingInfo["+strconv.Itoa(numSupportingInfo)+"]", &resource.SupportingInfo[numSupportingInfo], htmlAttrs)
}
func (resource *DeviceRequest) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *DeviceRequest) T_RelevantHistory(frs []FhirResource, numRelevantHistory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelevantHistory >= len(resource.RelevantHistory) {
		return ReferenceInput(frs, "relevantHistory["+strconv.Itoa(numRelevantHistory)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "relevantHistory["+strconv.Itoa(numRelevantHistory)+"]", &resource.RelevantHistory[numRelevantHistory], htmlAttrs)
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
func (resource *DeviceRequest) T_ParameterValueQuantity(numParameter int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return QuantityInput("parameter["+strconv.Itoa(numParameter)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("parameter["+strconv.Itoa(numParameter)+"].valueQuantity", resource.Parameter[numParameter].ValueQuantity, optionsValueSet, htmlAttrs)
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
