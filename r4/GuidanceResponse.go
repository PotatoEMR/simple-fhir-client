package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/a-h/templ"
)

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
	OccurrenceDateTime    *FhirDateTime     `json:"occurrenceDateTime,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
func (r GuidanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGuidanceResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherGuidanceResponse: OtherGuidanceResponse(r),
		ResourceType:          "GuidanceResponse",
	})
}

// json -> struct, first reject if resourceType != GuidanceResponse
func (r *GuidanceResponse) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "GuidanceResponse" {
		return errors.New("resourceType not GuidanceResponse")
	}
	return json.Unmarshal(data, (*OtherGuidanceResponse)(r))
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
func (resource *GuidanceResponse) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *GuidanceResponse) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *GuidanceResponse) T_OccurrenceDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("occurrenceDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("occurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *GuidanceResponse) T_Performer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "performer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "performer", resource.Performer, htmlAttrs)
}
func (resource *GuidanceResponse) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *GuidanceResponse) T_ReasonReference(frs []FhirResource, numReasonReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonReference >= len(resource.ReasonReference) {
		return ReferenceInput(frs, "reasonReference["+strconv.Itoa(numReasonReference)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "reasonReference["+strconv.Itoa(numReasonReference)+"]", &resource.ReasonReference[numReasonReference], htmlAttrs)
}
func (resource *GuidanceResponse) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *GuidanceResponse) T_EvaluationMessage(frs []FhirResource, numEvaluationMessage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvaluationMessage >= len(resource.EvaluationMessage) {
		return ReferenceInput(frs, "evaluationMessage["+strconv.Itoa(numEvaluationMessage)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "evaluationMessage["+strconv.Itoa(numEvaluationMessage)+"]", &resource.EvaluationMessage[numEvaluationMessage], htmlAttrs)
}
func (resource *GuidanceResponse) T_OutputParameters(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "outputParameters", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "outputParameters", resource.OutputParameters, htmlAttrs)
}
func (resource *GuidanceResponse) T_Result(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "result", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "result", resource.Result, htmlAttrs)
}
func (resource *GuidanceResponse) T_DataRequirement(numDataRequirement int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDataRequirement >= len(resource.DataRequirement) {
		return DataRequirementInput("dataRequirement["+strconv.Itoa(numDataRequirement)+"]", nil, htmlAttrs)
	}
	return DataRequirementInput("dataRequirement["+strconv.Itoa(numDataRequirement)+"]", &resource.DataRequirement[numDataRequirement], htmlAttrs)
}
