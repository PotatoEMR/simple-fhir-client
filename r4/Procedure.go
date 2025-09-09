package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Procedure
type Procedure struct {
	Id                    *string                `json:"id,omitempty"`
	Meta                  *Meta                  `json:"meta,omitempty"`
	ImplicitRules         *string                `json:"implicitRules,omitempty"`
	Language              *string                `json:"language,omitempty"`
	Text                  *Narrative             `json:"text,omitempty"`
	Contained             []Resource             `json:"contained,omitempty"`
	Extension             []Extension            `json:"extension,omitempty"`
	ModifierExtension     []Extension            `json:"modifierExtension,omitempty"`
	Identifier            []Identifier           `json:"identifier,omitempty"`
	InstantiatesCanonical []string               `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string               `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference            `json:"basedOn,omitempty"`
	PartOf                []Reference            `json:"partOf,omitempty"`
	Status                string                 `json:"status"`
	StatusReason          *CodeableConcept       `json:"statusReason,omitempty"`
	Category              *CodeableConcept       `json:"category,omitempty"`
	Code                  *CodeableConcept       `json:"code,omitempty"`
	Subject               Reference              `json:"subject"`
	Encounter             *Reference             `json:"encounter,omitempty"`
	PerformedDateTime     *string                `json:"performedDateTime,omitempty"`
	PerformedPeriod       *Period                `json:"performedPeriod,omitempty"`
	PerformedString       *string                `json:"performedString,omitempty"`
	PerformedAge          *Age                   `json:"performedAge,omitempty"`
	PerformedRange        *Range                 `json:"performedRange,omitempty"`
	Recorder              *Reference             `json:"recorder,omitempty"`
	Asserter              *Reference             `json:"asserter,omitempty"`
	Performer             []ProcedurePerformer   `json:"performer,omitempty"`
	Location              *Reference             `json:"location,omitempty"`
	ReasonCode            []CodeableConcept      `json:"reasonCode,omitempty"`
	ReasonReference       []Reference            `json:"reasonReference,omitempty"`
	BodySite              []CodeableConcept      `json:"bodySite,omitempty"`
	Outcome               *CodeableConcept       `json:"outcome,omitempty"`
	Report                []Reference            `json:"report,omitempty"`
	Complication          []CodeableConcept      `json:"complication,omitempty"`
	ComplicationDetail    []Reference            `json:"complicationDetail,omitempty"`
	FollowUp              []CodeableConcept      `json:"followUp,omitempty"`
	Note                  []Annotation           `json:"note,omitempty"`
	FocalDevice           []ProcedureFocalDevice `json:"focalDevice,omitempty"`
	UsedReference         []Reference            `json:"usedReference,omitempty"`
	UsedCode              []CodeableConcept      `json:"usedCode,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Procedure
type ProcedurePerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
	OnBehalfOf        *Reference       `json:"onBehalfOf,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Procedure
type ProcedureFocalDevice struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Action            *CodeableConcept `json:"action,omitempty"`
	Manipulated       Reference        `json:"manipulated"`
}

type OtherProcedure Procedure

// on convert struct to json, automatically add resourceType=Procedure
func (r Procedure) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProcedure
		ResourceType string `json:"resourceType"`
	}{
		OtherProcedure: OtherProcedure(r),
		ResourceType:   "Procedure",
	})
}
func (r Procedure) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Procedure/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Procedure"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Procedure) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *Procedure) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *Procedure) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSEvent_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_StatusReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Category(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category", resource.Category, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_PerformedDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("performedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("performedDateTime", resource.PerformedDateTime, htmlAttrs)
}
func (resource *Procedure) T_PerformedString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("performedString", nil, htmlAttrs)
	}
	return StringInput("performedString", resource.PerformedString, htmlAttrs)
}
func (resource *Procedure) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_BodySite(numBodySite int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBodySite >= len(resource.BodySite) {
		return CodeableConceptSelect("bodySite["+strconv.Itoa(numBodySite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("bodySite["+strconv.Itoa(numBodySite)+"]", &resource.BodySite[numBodySite], optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Outcome(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("outcome", resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Complication(numComplication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numComplication >= len(resource.Complication) {
		return CodeableConceptSelect("complication["+strconv.Itoa(numComplication)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("complication["+strconv.Itoa(numComplication)+"]", &resource.Complication[numComplication], optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_FollowUp(numFollowUp int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numFollowUp >= len(resource.FollowUp) {
		return CodeableConceptSelect("followUp["+strconv.Itoa(numFollowUp)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("followUp["+strconv.Itoa(numFollowUp)+"]", &resource.FollowUp[numFollowUp], optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Procedure) T_UsedCode(numUsedCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numUsedCode >= len(resource.UsedCode) {
		return CodeableConceptSelect("usedCode["+strconv.Itoa(numUsedCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("usedCode["+strconv.Itoa(numUsedCode)+"]", &resource.UsedCode[numUsedCode], optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_FocalDeviceAction(numFocalDevice int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numFocalDevice >= len(resource.FocalDevice) {
		return CodeableConceptSelect("focalDevice["+strconv.Itoa(numFocalDevice)+"].action", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("focalDevice["+strconv.Itoa(numFocalDevice)+"].action", resource.FocalDevice[numFocalDevice].Action, optionsValueSet, htmlAttrs)
}
