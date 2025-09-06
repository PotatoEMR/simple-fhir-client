package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/ServiceRequest
type ServiceRequest struct {
	Id                      *string           `json:"id,omitempty"`
	Meta                    *Meta             `json:"meta,omitempty"`
	ImplicitRules           *string           `json:"implicitRules,omitempty"`
	Language                *string           `json:"language,omitempty"`
	Text                    *Narrative        `json:"text,omitempty"`
	Contained               []Resource        `json:"contained,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Identifier              []Identifier      `json:"identifier,omitempty"`
	InstantiatesCanonical   []string          `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri         []string          `json:"instantiatesUri,omitempty"`
	BasedOn                 []Reference       `json:"basedOn,omitempty"`
	Replaces                []Reference       `json:"replaces,omitempty"`
	Requisition             *Identifier       `json:"requisition,omitempty"`
	Status                  string            `json:"status"`
	Intent                  string            `json:"intent"`
	Category                []CodeableConcept `json:"category,omitempty"`
	Priority                *string           `json:"priority,omitempty"`
	DoNotPerform            *bool             `json:"doNotPerform,omitempty"`
	Code                    *CodeableConcept  `json:"code,omitempty"`
	OrderDetail             []CodeableConcept `json:"orderDetail,omitempty"`
	QuantityQuantity        *Quantity         `json:"quantityQuantity,omitempty"`
	QuantityRatio           *Ratio            `json:"quantityRatio,omitempty"`
	QuantityRange           *Range            `json:"quantityRange,omitempty"`
	Subject                 Reference         `json:"subject"`
	Encounter               *Reference        `json:"encounter,omitempty"`
	OccurrenceDateTime      *string           `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod        *Period           `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming        *Timing           `json:"occurrenceTiming,omitempty"`
	AsNeededBoolean         *bool             `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept  `json:"asNeededCodeableConcept,omitempty"`
	AuthoredOn              *string           `json:"authoredOn,omitempty"`
	Requester               *Reference        `json:"requester,omitempty"`
	PerformerType           *CodeableConcept  `json:"performerType,omitempty"`
	Performer               []Reference       `json:"performer,omitempty"`
	LocationCode            []CodeableConcept `json:"locationCode,omitempty"`
	LocationReference       []Reference       `json:"locationReference,omitempty"`
	ReasonCode              []CodeableConcept `json:"reasonCode,omitempty"`
	ReasonReference         []Reference       `json:"reasonReference,omitempty"`
	Insurance               []Reference       `json:"insurance,omitempty"`
	SupportingInfo          []Reference       `json:"supportingInfo,omitempty"`
	Specimen                []Reference       `json:"specimen,omitempty"`
	BodySite                []CodeableConcept `json:"bodySite,omitempty"`
	Note                    []Annotation      `json:"note,omitempty"`
	PatientInstruction      *string           `json:"patientInstruction,omitempty"`
	RelevantHistory         []Reference       `json:"relevantHistory,omitempty"`
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
func (resource *ServiceRequest) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ServiceRequest.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ServiceRequest.Code", resource.Code, optionsValueSet)
}
func (resource *ServiceRequest) T_OrderDetail(numOrderDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.OrderDetail) >= numOrderDetail {
		return CodeableConceptSelect("ServiceRequest.OrderDetail["+strconv.Itoa(numOrderDetail)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ServiceRequest.OrderDetail["+strconv.Itoa(numOrderDetail)+"]", &resource.OrderDetail[numOrderDetail], optionsValueSet)
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
func (resource *ServiceRequest) T_LocationCode(numLocationCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.LocationCode) >= numLocationCode {
		return CodeableConceptSelect("ServiceRequest.LocationCode["+strconv.Itoa(numLocationCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ServiceRequest.LocationCode["+strconv.Itoa(numLocationCode)+"]", &resource.LocationCode[numLocationCode], optionsValueSet)
}
func (resource *ServiceRequest) T_ReasonCode(numReasonCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ReasonCode) >= numReasonCode {
		return CodeableConceptSelect("ServiceRequest.ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ServiceRequest.ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet)
}
func (resource *ServiceRequest) T_BodySite(numBodySite int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.BodySite) >= numBodySite {
		return CodeableConceptSelect("ServiceRequest.BodySite["+strconv.Itoa(numBodySite)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ServiceRequest.BodySite["+strconv.Itoa(numBodySite)+"]", &resource.BodySite[numBodySite], optionsValueSet)
}
func (resource *ServiceRequest) T_PatientInstruction() templ.Component {

	if resource == nil {
		return StringInput("ServiceRequest.PatientInstruction", nil)
	}
	return StringInput("ServiceRequest.PatientInstruction", resource.PatientInstruction)
}
