package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/GuidanceResponse
type GuidanceResponse struct {
	Id                    *string             `json:"id,omitempty"`
	Meta                  *Meta               `json:"meta,omitempty"`
	ImplicitRules         *string             `json:"implicitRules,omitempty"`
	Language              *string             `json:"language,omitempty"`
	Text                  *Narrative          `json:"text,omitempty"`
	Contained             []Resource          `json:"contained,omitempty"`
	Extension             []Extension         `json:"extension,omitempty"`
	ModifierExtension     []Extension         `json:"modifierExtension,omitempty"`
	RequestIdentifier     *Identifier         `json:"requestIdentifier,omitempty"`
	Identifier            []Identifier        `json:"identifier,omitempty"`
	ModuleUri             string              `json:"moduleUri"`
	ModuleCanonical       string              `json:"moduleCanonical"`
	ModuleCodeableConcept CodeableConcept     `json:"moduleCodeableConcept"`
	Status                string              `json:"status"`
	Subject               *Reference          `json:"subject,omitempty"`
	Encounter             *Reference          `json:"encounter,omitempty"`
	OccurrenceDateTime    *string             `json:"occurrenceDateTime,omitempty"`
	Performer             *Reference          `json:"performer,omitempty"`
	Reason                []CodeableReference `json:"reason,omitempty"`
	Note                  []Annotation        `json:"note,omitempty"`
	EvaluationMessage     *Reference          `json:"evaluationMessage,omitempty"`
	OutputParameters      *Reference          `json:"outputParameters,omitempty"`
	Result                []Reference         `json:"result,omitempty"`
	DataRequirement       []DataRequirement   `json:"dataRequirement,omitempty"`
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

func (resource *GuidanceResponse) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *GuidanceResponse) T_Status() templ.Component {
	optionsValueSet := VSGuidance_response_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
