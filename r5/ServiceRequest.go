package r5

//generated with command go run ./bultaoreune -nodownload
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

func (resource *ServiceRequest) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ServiceRequest.Id", nil)
	}
	return StringInput("ServiceRequest.Id", resource.Id)
}
func (resource *ServiceRequest) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ServiceRequest.ImplicitRules", nil)
	}
	return StringInput("ServiceRequest.ImplicitRules", resource.ImplicitRules)
}
func (resource *ServiceRequest) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ServiceRequest.Language", nil, optionsValueSet)
	}
	return CodeSelect("ServiceRequest.Language", resource.Language, optionsValueSet)
}
func (resource *ServiceRequest) T_InstantiatesCanonical(numInstantiatesCanonical int) templ.Component {

	if resource == nil || len(resource.InstantiatesCanonical) >= numInstantiatesCanonical {
		return StringInput("ServiceRequest.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil)
	}
	return StringInput("ServiceRequest.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical])
}
func (resource *ServiceRequest) T_InstantiatesUri(numInstantiatesUri int) templ.Component {

	if resource == nil || len(resource.InstantiatesUri) >= numInstantiatesUri {
		return StringInput("ServiceRequest.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil)
	}
	return StringInput("ServiceRequest.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri])
}
func (resource *ServiceRequest) T_Status() templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("ServiceRequest.Status", nil, optionsValueSet)
	}
	return CodeSelect("ServiceRequest.Status", &resource.Status, optionsValueSet)
}
func (resource *ServiceRequest) T_Intent() templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("ServiceRequest.Intent", nil, optionsValueSet)
	}
	return CodeSelect("ServiceRequest.Intent", &resource.Intent, optionsValueSet)
}
func (resource *ServiceRequest) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("ServiceRequest.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ServiceRequest.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *ServiceRequest) T_Priority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("ServiceRequest.Priority", nil, optionsValueSet)
	}
	return CodeSelect("ServiceRequest.Priority", resource.Priority, optionsValueSet)
}
func (resource *ServiceRequest) T_DoNotPerform() templ.Component {

	if resource == nil {
		return BoolInput("ServiceRequest.DoNotPerform", nil)
	}
	return BoolInput("ServiceRequest.DoNotPerform", resource.DoNotPerform)
}
func (resource *ServiceRequest) T_AuthoredOn() templ.Component {

	if resource == nil {
		return StringInput("ServiceRequest.AuthoredOn", nil)
	}
	return StringInput("ServiceRequest.AuthoredOn", resource.AuthoredOn)
}
func (resource *ServiceRequest) T_PerformerType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ServiceRequest.PerformerType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ServiceRequest.PerformerType", resource.PerformerType, optionsValueSet)
}
func (resource *ServiceRequest) T_BodySite(numBodySite int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.BodySite) >= numBodySite {
		return CodeableConceptSelect("ServiceRequest.BodySite["+strconv.Itoa(numBodySite)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ServiceRequest.BodySite["+strconv.Itoa(numBodySite)+"]", &resource.BodySite[numBodySite], optionsValueSet)
}
func (resource *ServiceRequest) T_OrderDetailId(numOrderDetail int) templ.Component {

	if resource == nil || len(resource.OrderDetail) >= numOrderDetail {
		return StringInput("ServiceRequest.OrderDetail["+strconv.Itoa(numOrderDetail)+"].Id", nil)
	}
	return StringInput("ServiceRequest.OrderDetail["+strconv.Itoa(numOrderDetail)+"].Id", resource.OrderDetail[numOrderDetail].Id)
}
func (resource *ServiceRequest) T_OrderDetailParameterId(numOrderDetail int, numParameter int) templ.Component {

	if resource == nil || len(resource.OrderDetail) >= numOrderDetail || len(resource.OrderDetail[numOrderDetail].Parameter) >= numParameter {
		return StringInput("ServiceRequest.OrderDetail["+strconv.Itoa(numOrderDetail)+"].Parameter["+strconv.Itoa(numParameter)+"].Id", nil)
	}
	return StringInput("ServiceRequest.OrderDetail["+strconv.Itoa(numOrderDetail)+"].Parameter["+strconv.Itoa(numParameter)+"].Id", resource.OrderDetail[numOrderDetail].Parameter[numParameter].Id)
}
func (resource *ServiceRequest) T_OrderDetailParameterCode(numOrderDetail int, numParameter int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.OrderDetail) >= numOrderDetail || len(resource.OrderDetail[numOrderDetail].Parameter) >= numParameter {
		return CodeableConceptSelect("ServiceRequest.OrderDetail["+strconv.Itoa(numOrderDetail)+"].Parameter["+strconv.Itoa(numParameter)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ServiceRequest.OrderDetail["+strconv.Itoa(numOrderDetail)+"].Parameter["+strconv.Itoa(numParameter)+"].Code", &resource.OrderDetail[numOrderDetail].Parameter[numParameter].Code, optionsValueSet)
}
func (resource *ServiceRequest) T_PatientInstructionId(numPatientInstruction int) templ.Component {

	if resource == nil || len(resource.PatientInstruction) >= numPatientInstruction {
		return StringInput("ServiceRequest.PatientInstruction["+strconv.Itoa(numPatientInstruction)+"].Id", nil)
	}
	return StringInput("ServiceRequest.PatientInstruction["+strconv.Itoa(numPatientInstruction)+"].Id", resource.PatientInstruction[numPatientInstruction].Id)
}
