package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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

func (resource *Procedure) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Procedure.Id", nil)
	}
	return StringInput("Procedure.Id", resource.Id)
}
func (resource *Procedure) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Procedure.ImplicitRules", nil)
	}
	return StringInput("Procedure.ImplicitRules", resource.ImplicitRules)
}
func (resource *Procedure) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Procedure.Language", nil, optionsValueSet)
	}
	return CodeSelect("Procedure.Language", resource.Language, optionsValueSet)
}
func (resource *Procedure) T_InstantiatesCanonical(numInstantiatesCanonical int) templ.Component {

	if resource == nil || len(resource.InstantiatesCanonical) >= numInstantiatesCanonical {
		return StringInput("Procedure.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil)
	}
	return StringInput("Procedure.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical])
}
func (resource *Procedure) T_InstantiatesUri(numInstantiatesUri int) templ.Component {

	if resource == nil || len(resource.InstantiatesUri) >= numInstantiatesUri {
		return StringInput("Procedure.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil)
	}
	return StringInput("Procedure.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri])
}
func (resource *Procedure) T_Status() templ.Component {
	optionsValueSet := VSEvent_status

	if resource == nil {
		return CodeSelect("Procedure.Status", nil, optionsValueSet)
	}
	return CodeSelect("Procedure.Status", &resource.Status, optionsValueSet)
}
func (resource *Procedure) T_StatusReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Procedure.StatusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Procedure.StatusReason", resource.StatusReason, optionsValueSet)
}
func (resource *Procedure) T_Category(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Procedure.Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Procedure.Category", resource.Category, optionsValueSet)
}
func (resource *Procedure) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Procedure.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Procedure.Code", resource.Code, optionsValueSet)
}
func (resource *Procedure) T_ReasonCode(numReasonCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ReasonCode) >= numReasonCode {
		return CodeableConceptSelect("Procedure.ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Procedure.ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet)
}
func (resource *Procedure) T_BodySite(numBodySite int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.BodySite) >= numBodySite {
		return CodeableConceptSelect("Procedure.BodySite["+strconv.Itoa(numBodySite)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Procedure.BodySite["+strconv.Itoa(numBodySite)+"]", &resource.BodySite[numBodySite], optionsValueSet)
}
func (resource *Procedure) T_Outcome(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Procedure.Outcome", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Procedure.Outcome", resource.Outcome, optionsValueSet)
}
func (resource *Procedure) T_Complication(numComplication int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Complication) >= numComplication {
		return CodeableConceptSelect("Procedure.Complication["+strconv.Itoa(numComplication)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Procedure.Complication["+strconv.Itoa(numComplication)+"]", &resource.Complication[numComplication], optionsValueSet)
}
func (resource *Procedure) T_FollowUp(numFollowUp int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.FollowUp) >= numFollowUp {
		return CodeableConceptSelect("Procedure.FollowUp["+strconv.Itoa(numFollowUp)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Procedure.FollowUp["+strconv.Itoa(numFollowUp)+"]", &resource.FollowUp[numFollowUp], optionsValueSet)
}
func (resource *Procedure) T_UsedCode(numUsedCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.UsedCode) >= numUsedCode {
		return CodeableConceptSelect("Procedure.UsedCode["+strconv.Itoa(numUsedCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Procedure.UsedCode["+strconv.Itoa(numUsedCode)+"]", &resource.UsedCode[numUsedCode], optionsValueSet)
}
func (resource *Procedure) T_PerformerId(numPerformer int) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return StringInput("Procedure.Performer["+strconv.Itoa(numPerformer)+"].Id", nil)
	}
	return StringInput("Procedure.Performer["+strconv.Itoa(numPerformer)+"].Id", resource.Performer[numPerformer].Id)
}
func (resource *Procedure) T_PerformerFunction(numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return CodeableConceptSelect("Procedure.Performer["+strconv.Itoa(numPerformer)+"].Function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Procedure.Performer["+strconv.Itoa(numPerformer)+"].Function", resource.Performer[numPerformer].Function, optionsValueSet)
}
func (resource *Procedure) T_FocalDeviceId(numFocalDevice int) templ.Component {

	if resource == nil || len(resource.FocalDevice) >= numFocalDevice {
		return StringInput("Procedure.FocalDevice["+strconv.Itoa(numFocalDevice)+"].Id", nil)
	}
	return StringInput("Procedure.FocalDevice["+strconv.Itoa(numFocalDevice)+"].Id", resource.FocalDevice[numFocalDevice].Id)
}
func (resource *Procedure) T_FocalDeviceAction(numFocalDevice int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.FocalDevice) >= numFocalDevice {
		return CodeableConceptSelect("Procedure.FocalDevice["+strconv.Itoa(numFocalDevice)+"].Action", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Procedure.FocalDevice["+strconv.Itoa(numFocalDevice)+"].Action", resource.FocalDevice[numFocalDevice].Action, optionsValueSet)
}
