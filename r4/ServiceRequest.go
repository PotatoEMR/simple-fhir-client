package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/ServiceRequest
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
	OccurrenceDateTime      *time.Time        `json:"occurrenceDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	OccurrencePeriod        *Period           `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming        *Timing           `json:"occurrenceTiming,omitempty"`
	AsNeededBoolean         *bool             `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept  `json:"asNeededCodeableConcept,omitempty"`
	AuthoredOn              *time.Time        `json:"authoredOn,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *ServiceRequest) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *ServiceRequest) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *ServiceRequest) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_Intent(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("Intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_Priority(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_DoNotPerform(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("DoNotPerform", nil, htmlAttrs)
	}
	return BoolInput("DoNotPerform", resource.DoNotPerform, htmlAttrs)
}
func (resource *ServiceRequest) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_OrderDetail(numOrderDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOrderDetail >= len(resource.OrderDetail) {
		return CodeableConceptSelect("OrderDetail["+strconv.Itoa(numOrderDetail)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OrderDetail["+strconv.Itoa(numOrderDetail)+"]", &resource.OrderDetail[numOrderDetail], optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_OccurrenceDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("OccurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("OccurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *ServiceRequest) T_AsNeededBoolean(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("AsNeededBoolean", nil, htmlAttrs)
	}
	return BoolInput("AsNeededBoolean", resource.AsNeededBoolean, htmlAttrs)
}
func (resource *ServiceRequest) T_AsNeededCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("AsNeededCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AsNeededCodeableConcept", resource.AsNeededCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_AuthoredOn(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("AuthoredOn", nil, htmlAttrs)
	}
	return DateTimeInput("AuthoredOn", resource.AuthoredOn, htmlAttrs)
}
func (resource *ServiceRequest) T_PerformerType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PerformerType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PerformerType", resource.PerformerType, optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_LocationCode(numLocationCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLocationCode >= len(resource.LocationCode) {
		return CodeableConceptSelect("LocationCode["+strconv.Itoa(numLocationCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("LocationCode["+strconv.Itoa(numLocationCode)+"]", &resource.LocationCode[numLocationCode], optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_BodySite(numBodySite int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBodySite >= len(resource.BodySite) {
		return CodeableConceptSelect("BodySite["+strconv.Itoa(numBodySite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("BodySite["+strconv.Itoa(numBodySite)+"]", &resource.BodySite[numBodySite], optionsValueSet, htmlAttrs)
}
func (resource *ServiceRequest) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *ServiceRequest) T_PatientInstruction(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PatientInstruction", nil, htmlAttrs)
	}
	return StringInput("PatientInstruction", resource.PatientInstruction, htmlAttrs)
}
