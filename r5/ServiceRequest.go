package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	QuantityQuantity        *Quantity                          `json:"quantityQuantity"`
	QuantityRatio           *Ratio                             `json:"quantityRatio"`
	QuantityRange           *Range                             `json:"quantityRange"`
	Subject                 Reference                          `json:"subject"`
	Focus                   []Reference                        `json:"focus,omitempty"`
	Encounter               *Reference                         `json:"encounter,omitempty"`
	OccurrenceDateTime      *string                            `json:"occurrenceDateTime"`
	OccurrencePeriod        *Period                            `json:"occurrencePeriod"`
	OccurrenceTiming        *Timing                            `json:"occurrenceTiming"`
	AsNeededBoolean         *bool                              `json:"asNeededBoolean"`
	AsNeededCodeableConcept *CodeableConcept                   `json:"asNeededCodeableConcept"`
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
	InstructionMarkdown  *string     `json:"instructionMarkdown"`
	InstructionReference *Reference  `json:"instructionReference"`
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

func (resource *ServiceRequest) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ServiceRequest) T_Status() templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ServiceRequest) T_Intent() templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("intent", nil, optionsValueSet)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet)
}
func (resource *ServiceRequest) T_Category(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *ServiceRequest) T_Priority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet)
}
func (resource *ServiceRequest) T_PerformerType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("performerType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("performerType", resource.PerformerType, optionsValueSet)
}
func (resource *ServiceRequest) T_BodySite(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("bodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("bodySite", &resource.BodySite[0], optionsValueSet)
}
func (resource *ServiceRequest) T_OrderDetailParameterCode(numOrderDetail int, numParameter int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.OrderDetail[numOrderDetail].Parameter) >= numParameter {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.OrderDetail[numOrderDetail].Parameter[numParameter].Code, optionsValueSet)
}
