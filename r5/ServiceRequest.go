package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ServiceRequest
type ServiceRequest struct {
	Id                      *string                            `json:"id,omitempty"`
	Meta                    *Meta                              `json:"meta,omitempty"`
	ImplicitRules           *string                            `json:"implicitRules,omitempty"`
	Language                *string                            `json:"language,omitempty"`
	Text                    *Narrative                         `json:"text,omitempty"`
	Contained               []Resource                         `json:"contained,omitempty"`
	Extension               []Extension                        `json:"extension,omitempty"`
	ModifierExtension       []Extension                        `json:"modifierExtension,omitempty"`
	Identifier              []Identifier                       `json:"identifier,omitempty"`
	InstantiatesCanonical   []string                           `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri         []string                           `json:"instantiatesUri,omitempty"`
	BasedOn                 []Reference                        `json:"basedOn,omitempty"`
	Replaces                []Reference                        `json:"replaces,omitempty"`
	Requisition             *Identifier                        `json:"requisition,omitempty"`
	Status                  string                             `json:"status"`
	Intent                  string                             `json:"intent"`
	Category                []CodeableConcept                  `json:"category,omitempty"`
	Priority                *string                            `json:"priority,omitempty"`
	DoNotPerform            *bool                              `json:"doNotPerform,omitempty"`
	Code                    *CodeableReference                 `json:"code,omitempty"`
	OrderDetail             []ServiceRequestOrderDetail        `json:"orderDetail,omitempty"`
	QuantityQuantity        *Quantity                          `json:"quantityQuantity,omitempty"`
	QuantityRatio           *Ratio                             `json:"quantityRatio,omitempty"`
	QuantityRange           *Range                             `json:"quantityRange,omitempty"`
	Subject                 Reference                          `json:"subject"`
	Focus                   []Reference                        `json:"focus,omitempty"`
	Encounter               *Reference                         `json:"encounter,omitempty"`
	OccurrenceDateTime      *string                            `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod        *Period                            `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming        *Timing                            `json:"occurrenceTiming,omitempty"`
	AsNeededBoolean         *bool                              `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept                   `json:"asNeededCodeableConcept,omitempty"`
	AuthoredOn              *string                            `json:"authoredOn,omitempty"`
	Requester               *Reference                         `json:"requester,omitempty"`
	PerformerType           *CodeableConcept                   `json:"performerType,omitempty"`
	Performer               []Reference                        `json:"performer,omitempty"`
	Location                []CodeableReference                `json:"location,omitempty"`
	Reason                  []CodeableReference                `json:"reason,omitempty"`
	Insurance               []Reference                        `json:"insurance,omitempty"`
	SupportingInfo          []CodeableReference                `json:"supportingInfo,omitempty"`
	Specimen                []Reference                        `json:"specimen,omitempty"`
	BodySite                []CodeableConcept                  `json:"bodySite,omitempty"`
	BodyStructure           *Reference                         `json:"bodyStructure,omitempty"`
	Note                    []Annotation                       `json:"note,omitempty"`
	PatientInstruction      []ServiceRequestPatientInstruction `json:"patientInstruction,omitempty"`
	RelevantHistory         []Reference                        `json:"relevantHistory,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ServiceRequest
type ServiceRequestOrderDetail struct {
	Id                *string                              `json:"id,omitempty"`
	Extension         []Extension                          `json:"extension,omitempty"`
	ModifierExtension []Extension                          `json:"modifierExtension,omitempty"`
	ParameterFocus    *CodeableReference                   `json:"parameterFocus,omitempty"`
	Parameter         []ServiceRequestOrderDetailParameter `json:"parameter"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ServiceRequest
type ServiceRequestOrderDetailParameter struct {
	Id                   *string         `json:"id,omitempty"`
	Extension            []Extension     `json:"extension,omitempty"`
	ModifierExtension    []Extension     `json:"modifierExtension,omitempty"`
	Code                 CodeableConcept `json:"code"`
	ValueQuantity        Quantity        `json:"valueQuantity"`
	ValueRatio           Ratio           `json:"valueRatio"`
	ValueRange           Range           `json:"valueRange"`
	ValueBoolean         bool            `json:"valueBoolean"`
	ValueCodeableConcept CodeableConcept `json:"valueCodeableConcept"`
	ValueString          string          `json:"valueString"`
	ValuePeriod          Period          `json:"valuePeriod"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ServiceRequest
type ServiceRequestPatientInstruction struct {
	Id                   *string     `json:"id,omitempty"`
	Extension            []Extension `json:"extension,omitempty"`
	ModifierExtension    []Extension `json:"modifierExtension,omitempty"`
	InstructionMarkdown  *string     `json:"instructionMarkdown,omitempty"`
	InstructionReference *Reference  `json:"instructionReference,omitempty"`
}

type OtherServiceRequest ServiceRequest

// on convert struct to json, automatically add resourceType=ServiceRequest
func (r ServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherServiceRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherServiceRequest: OtherServiceRequest(r),
		ResourceType:        "ServiceRequest",
	})
}
func (r ServiceRequest) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ServiceRequest/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ServiceRequest"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ServiceRequest) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *ServiceRequest) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *ServiceRequest) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_Intent(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_Priority(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_DoNotPerform(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("doNotPerform", nil, htmlAttrs)
	}
	return BoolInput("doNotPerform", resource.DoNotPerform, htmlAttrs)
}
func (resource *ServiceRequest) T_OccurrenceDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("occurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("occurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *ServiceRequest) T_AsNeededBoolean(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("asNeededBoolean", nil, htmlAttrs)
	}
	return BoolInput("asNeededBoolean", resource.AsNeededBoolean, htmlAttrs)
}
func (resource *ServiceRequest) T_AsNeededCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("asNeededCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("asNeededCodeableConcept", resource.AsNeededCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_AuthoredOn(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("authoredOn", nil, htmlAttrs)
	}
	return DateTimeInput("authoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *ServiceRequest) T_PerformerType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("performerType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performerType", resource.PerformerType, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_BodySite(numBodySite int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBodySite >= len(resource.BodySite) {
		return CodeableConceptSelect("bodySite["+strconv.Itoa(numBodySite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("bodySite["+strconv.Itoa(numBodySite)+"]", &resource.BodySite[numBodySite], optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *ServiceRequest) T_OrderDetailParameterCode(numOrderDetail int, numParameter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOrderDetail >= len(resource.OrderDetail) || numParameter >= len(resource.OrderDetail[numOrderDetail].Parameter) {
		return CodeableConceptSelect("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].code", &resource.OrderDetail[numOrderDetail].Parameter[numParameter].Code, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_OrderDetailParameterValueBoolean(numOrderDetail int, numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOrderDetail >= len(resource.OrderDetail) || numParameter >= len(resource.OrderDetail[numOrderDetail].Parameter) {
		return BoolInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueBoolean", &resource.OrderDetail[numOrderDetail].Parameter[numParameter].ValueBoolean, htmlAttrs)
}
func (resource *ServiceRequest) T_OrderDetailParameterValueCodeableConcept(numOrderDetail int, numParameter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOrderDetail >= len(resource.OrderDetail) || numParameter >= len(resource.OrderDetail[numOrderDetail].Parameter) {
		return CodeableConceptSelect("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueCodeableConcept", &resource.OrderDetail[numOrderDetail].Parameter[numParameter].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_OrderDetailParameterValueString(numOrderDetail int, numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOrderDetail >= len(resource.OrderDetail) || numParameter >= len(resource.OrderDetail[numOrderDetail].Parameter) {
		return StringInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("orderDetail["+strconv.Itoa(numOrderDetail)+"].parameter["+strconv.Itoa(numParameter)+"].valueString", &resource.OrderDetail[numOrderDetail].Parameter[numParameter].ValueString, htmlAttrs)
}
func (resource *ServiceRequest) T_PatientInstructionInstructionMarkdown(numPatientInstruction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPatientInstruction >= len(resource.PatientInstruction) {
		return StringInput("patientInstruction["+strconv.Itoa(numPatientInstruction)+"].instructionMarkdown", nil, htmlAttrs)
	}
	return StringInput("patientInstruction["+strconv.Itoa(numPatientInstruction)+"].instructionMarkdown", resource.PatientInstruction[numPatientInstruction].InstructionMarkdown, htmlAttrs)
}
