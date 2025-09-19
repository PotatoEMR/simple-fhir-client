package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	OccurrenceDateTime    *FhirDateTime       `json:"occurrenceDateTime,omitempty"`
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
func (resource *GuidanceResponse) T_RequestIdentifier(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IdentifierInput("requestIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("requestIdentifier", resource.RequestIdentifier, htmlAttrs)
}
func (resource *GuidanceResponse) T_ModuleUri(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("moduleUri", nil, htmlAttrs)
	}
	return StringInput("moduleUri", &resource.ModuleUri, htmlAttrs)
}
func (resource *GuidanceResponse) T_ModuleCanonical(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("moduleCanonical", nil, htmlAttrs)
	}
	return StringInput("moduleCanonical", &resource.ModuleCanonical, htmlAttrs)
}
func (resource *GuidanceResponse) T_ModuleCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("moduleCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("moduleCodeableConcept", &resource.ModuleCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *GuidanceResponse) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSGuidance_response_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *GuidanceResponse) T_Subject(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("subject", nil, htmlAttrs)
	}
	return ReferenceInput("subject", resource.Subject, htmlAttrs)
}
func (resource *GuidanceResponse) T_Encounter(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("encounter", nil, htmlAttrs)
	}
	return ReferenceInput("encounter", resource.Encounter, htmlAttrs)
}
func (resource *GuidanceResponse) T_OccurrenceDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("occurrenceDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("occurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *GuidanceResponse) T_Performer(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("performer", nil, htmlAttrs)
	}
	return ReferenceInput("performer", resource.Performer, htmlAttrs)
}
func (resource *GuidanceResponse) T_Reason(numReason int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], htmlAttrs)
}
func (resource *GuidanceResponse) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *GuidanceResponse) T_EvaluationMessage(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("evaluationMessage", nil, htmlAttrs)
	}
	return ReferenceInput("evaluationMessage", resource.EvaluationMessage, htmlAttrs)
}
func (resource *GuidanceResponse) T_OutputParameters(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("outputParameters", nil, htmlAttrs)
	}
	return ReferenceInput("outputParameters", resource.OutputParameters, htmlAttrs)
}
func (resource *GuidanceResponse) T_Result(numResult int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResult >= len(resource.Result) {
		return ReferenceInput("result["+strconv.Itoa(numResult)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("result["+strconv.Itoa(numResult)+"]", &resource.Result[numResult], htmlAttrs)
}
func (resource *GuidanceResponse) T_DataRequirement(numDataRequirement int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDataRequirement >= len(resource.DataRequirement) {
		return DataRequirementInput("dataRequirement["+strconv.Itoa(numDataRequirement)+"]", nil, htmlAttrs)
	}
	return DataRequirementInput("dataRequirement["+strconv.Itoa(numDataRequirement)+"]", &resource.DataRequirement[numDataRequirement], htmlAttrs)
}
