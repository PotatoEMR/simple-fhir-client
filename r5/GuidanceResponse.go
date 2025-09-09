package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

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
	OccurrenceDateTime    *time.Time          `json:"occurrenceDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r GuidanceResponse) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "GuidanceResponse/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "GuidanceResponse"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *GuidanceResponse) T_ModuleUri(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("GuidanceResponse.ModuleUri", nil, htmlAttrs)
	}
	return StringInput("GuidanceResponse.ModuleUri", &resource.ModuleUri, htmlAttrs)
}
func (resource *GuidanceResponse) T_ModuleCanonical(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("GuidanceResponse.ModuleCanonical", nil, htmlAttrs)
	}
	return StringInput("GuidanceResponse.ModuleCanonical", &resource.ModuleCanonical, htmlAttrs)
}
func (resource *GuidanceResponse) T_ModuleCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("GuidanceResponse.ModuleCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("GuidanceResponse.ModuleCodeableConcept", &resource.ModuleCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *GuidanceResponse) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSGuidance_response_status

	if resource == nil {
		return CodeSelect("GuidanceResponse.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("GuidanceResponse.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *GuidanceResponse) T_OccurrenceDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("GuidanceResponse.OccurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("GuidanceResponse.OccurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *GuidanceResponse) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("GuidanceResponse.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("GuidanceResponse.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
