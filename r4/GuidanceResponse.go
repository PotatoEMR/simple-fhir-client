//generated August 14 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/GuidanceResponse
type GuidanceResponse struct {
	Id                    *string           `json:"id,omitempty"`
	Meta                  *Meta             `json:"meta,omitempty"`
	ImplicitRules         *string           `json:"implicitRules,omitempty"`
	Language              *string           `json:"language,omitempty"`
	Text                  *Narrative        `json:"text,omitempty"`
	Contained             []Resource        `json:"contained,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	RequestIdentifier     *Identifier       `json:"requestIdentifier,omitempty"`
	Identifier            []Identifier      `json:"identifier,omitempty"`
	ModuleUri             string            `json:"moduleUri"`
	ModuleCanonical       string            `json:"moduleCanonical"`
	ModuleCodeableConcept CodeableConcept   `json:"moduleCodeableConcept"`
	Status                string            `json:"status"`
	Subject               *Reference        `json:"subject,omitempty"`
	Encounter             *Reference        `json:"encounter,omitempty"`
	OccurrenceDateTime    *string           `json:"occurrenceDateTime,omitempty"`
	Performer             *Reference        `json:"performer,omitempty"`
	ReasonCode            []CodeableConcept `json:"reasonCode,omitempty"`
	ReasonReference       []Reference       `json:"reasonReference,omitempty"`
	Note                  []Annotation      `json:"note,omitempty"`
	EvaluationMessage     []Reference       `json:"evaluationMessage,omitempty"`
	OutputParameters      *Reference        `json:"outputParameters,omitempty"`
	Result                *Reference        `json:"result,omitempty"`
	DataRequirement       []DataRequirement `json:"dataRequirement,omitempty"`
}

type OtherGuidanceResponse GuidanceResponse

// on convert struct to json, automatically add resourceType=GuidanceResponse
func (r GuidanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGuidanceResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherGuidanceResponse: OtherGuidanceResponse(r),
		ResourceType:          "GuidanceResponse",
	})
}
