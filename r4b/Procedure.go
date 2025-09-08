package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Procedure
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
	PerformedDateTime     *time.Time             `json:"performedDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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

// http://hl7.org/fhir/r4b/StructureDefinition/Procedure
type ProcedurePerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
	OnBehalfOf        *Reference       `json:"onBehalfOf,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Procedure
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
		return StringInput("Procedure.InstantiatesCanonical."+strconv.Itoa(numInstantiatesCanonical)+".", nil, htmlAttrs)
	}
	return StringInput("Procedure.InstantiatesCanonical."+strconv.Itoa(numInstantiatesCanonical)+".", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *Procedure) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs string) templ.Component {

	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("Procedure.InstantiatesUri."+strconv.Itoa(numInstantiatesUri)+".", nil, htmlAttrs)
	}
	return StringInput("Procedure.InstantiatesUri."+strconv.Itoa(numInstantiatesUri)+".", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *Procedure) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSEvent_status

	if resource == nil {
		return CodeSelect("Procedure.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Procedure.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_StatusReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Procedure.StatusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Procedure.StatusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Category(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Procedure.Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Procedure.Category", resource.Category, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Procedure.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Procedure.Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_PerformedDateTime(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Procedure.PerformedDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Procedure.PerformedDateTime", resource.PerformedDateTime, htmlAttrs)
}
func (resource *Procedure) T_PerformedString(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Procedure.PerformedString", nil, htmlAttrs)
	}
	return StringInput("Procedure.PerformedString", resource.PerformedString, htmlAttrs)
}
func (resource *Procedure) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("Procedure.ReasonCode."+strconv.Itoa(numReasonCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Procedure.ReasonCode."+strconv.Itoa(numReasonCode)+".", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_BodySite(numBodySite int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numBodySite >= len(resource.BodySite) {
		return CodeableConceptSelect("Procedure.BodySite."+strconv.Itoa(numBodySite)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Procedure.BodySite."+strconv.Itoa(numBodySite)+".", &resource.BodySite[numBodySite], optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Outcome(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Procedure.Outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Procedure.Outcome", resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Complication(numComplication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numComplication >= len(resource.Complication) {
		return CodeableConceptSelect("Procedure.Complication."+strconv.Itoa(numComplication)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Procedure.Complication."+strconv.Itoa(numComplication)+".", &resource.Complication[numComplication], optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_FollowUp(numFollowUp int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numFollowUp >= len(resource.FollowUp) {
		return CodeableConceptSelect("Procedure.FollowUp."+strconv.Itoa(numFollowUp)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Procedure.FollowUp."+strconv.Itoa(numFollowUp)+".", &resource.FollowUp[numFollowUp], optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Note(numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Procedure.Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("Procedure.Note."+strconv.Itoa(numNote)+".", &resource.Note[numNote], htmlAttrs)
}
func (resource *Procedure) T_UsedCode(numUsedCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numUsedCode >= len(resource.UsedCode) {
		return CodeableConceptSelect("Procedure.UsedCode."+strconv.Itoa(numUsedCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Procedure.UsedCode."+strconv.Itoa(numUsedCode)+".", &resource.UsedCode[numUsedCode], optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("Procedure.Performer."+strconv.Itoa(numPerformer)+"..Function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Procedure.Performer."+strconv.Itoa(numPerformer)+"..Function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_FocalDeviceAction(numFocalDevice int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numFocalDevice >= len(resource.FocalDevice) {
		return CodeableConceptSelect("Procedure.FocalDevice."+strconv.Itoa(numFocalDevice)+"..Action", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Procedure.FocalDevice."+strconv.Itoa(numFocalDevice)+"..Action", resource.FocalDevice[numFocalDevice].Action, optionsValueSet, htmlAttrs)
}
