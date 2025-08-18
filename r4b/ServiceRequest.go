//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

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
	QuantityQuantity        *Quantity         `json:"quantityQuantity"`
	QuantityRatio           *Ratio            `json:"quantityRatio"`
	QuantityRange           *Range            `json:"quantityRange"`
	Subject                 Reference         `json:"subject"`
	Encounter               *Reference        `json:"encounter,omitempty"`
	OccurrenceDateTime      *string           `json:"occurrenceDateTime"`
	OccurrencePeriod        *Period           `json:"occurrencePeriod"`
	OccurrenceTiming        *Timing           `json:"occurrenceTiming"`
	AsNeededBoolean         *bool             `json:"asNeededBoolean"`
	AsNeededCodeableConcept *CodeableConcept  `json:"asNeededCodeableConcept"`
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
