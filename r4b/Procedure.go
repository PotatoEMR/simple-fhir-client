package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
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
	PerformedDateTime     *FhirDateTime          `json:"performedDateTime,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
func (r Procedure) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProcedure
		ResourceType string `json:"resourceType"`
	}{
		OtherProcedure: OtherProcedure(r),
		ResourceType:   "Procedure",
	})
}

// json -> struct, first reject if resourceType != Procedure
func (r *Procedure) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "Procedure" {
		return errors.New("resourceType not Procedure")
	}
	return json.Unmarshal(data, (*OtherProcedure)(r))
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
func (resource *Procedure) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *Procedure) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *Procedure) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *Procedure) T_PartOf(frs []FhirResource, numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *Procedure) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSEvent_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_StatusReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Category(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category", resource.Category, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *Procedure) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *Procedure) T_PerformedDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("performedDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("performedDateTime", resource.PerformedDateTime, htmlAttrs)
}
func (resource *Procedure) T_PerformedPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("performedPeriod", nil, htmlAttrs)
	}
	return PeriodInput("performedPeriod", resource.PerformedPeriod, htmlAttrs)
}
func (resource *Procedure) T_PerformedString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("performedString", nil, htmlAttrs)
	}
	return StringInput("performedString", resource.PerformedString, htmlAttrs)
}
func (resource *Procedure) T_PerformedAge(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return AgeInput("performedAge", nil, htmlAttrs)
	}
	return AgeInput("performedAge", resource.PerformedAge, htmlAttrs)
}
func (resource *Procedure) T_PerformedRange(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return RangeInput("performedRange", nil, htmlAttrs)
	}
	return RangeInput("performedRange", resource.PerformedRange, htmlAttrs)
}
func (resource *Procedure) T_Recorder(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "recorder", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "recorder", resource.Recorder, htmlAttrs)
}
func (resource *Procedure) T_Asserter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "asserter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "asserter", resource.Asserter, htmlAttrs)
}
func (resource *Procedure) T_Location(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "location", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "location", resource.Location, htmlAttrs)
}
func (resource *Procedure) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_ReasonReference(frs []FhirResource, numReasonReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonReference >= len(resource.ReasonReference) {
		return ReferenceInput(frs, "reasonReference["+strconv.Itoa(numReasonReference)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "reasonReference["+strconv.Itoa(numReasonReference)+"]", &resource.ReasonReference[numReasonReference], htmlAttrs)
}
func (resource *Procedure) T_BodySite(numBodySite int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBodySite >= len(resource.BodySite) {
		return CodeableConceptSelect("bodySite["+strconv.Itoa(numBodySite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("bodySite["+strconv.Itoa(numBodySite)+"]", &resource.BodySite[numBodySite], optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Outcome(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("outcome", resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Report(frs []FhirResource, numReport int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReport >= len(resource.Report) {
		return ReferenceInput(frs, "report["+strconv.Itoa(numReport)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "report["+strconv.Itoa(numReport)+"]", &resource.Report[numReport], htmlAttrs)
}
func (resource *Procedure) T_Complication(numComplication int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComplication >= len(resource.Complication) {
		return CodeableConceptSelect("complication["+strconv.Itoa(numComplication)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("complication["+strconv.Itoa(numComplication)+"]", &resource.Complication[numComplication], optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_ComplicationDetail(frs []FhirResource, numComplicationDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComplicationDetail >= len(resource.ComplicationDetail) {
		return ReferenceInput(frs, "complicationDetail["+strconv.Itoa(numComplicationDetail)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "complicationDetail["+strconv.Itoa(numComplicationDetail)+"]", &resource.ComplicationDetail[numComplicationDetail], htmlAttrs)
}
func (resource *Procedure) T_FollowUp(numFollowUp int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFollowUp >= len(resource.FollowUp) {
		return CodeableConceptSelect("followUp["+strconv.Itoa(numFollowUp)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("followUp["+strconv.Itoa(numFollowUp)+"]", &resource.FollowUp[numFollowUp], optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Procedure) T_UsedReference(frs []FhirResource, numUsedReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUsedReference >= len(resource.UsedReference) {
		return ReferenceInput(frs, "usedReference["+strconv.Itoa(numUsedReference)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "usedReference["+strconv.Itoa(numUsedReference)+"]", &resource.UsedReference[numUsedReference], htmlAttrs)
}
func (resource *Procedure) T_UsedCode(numUsedCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUsedCode >= len(resource.UsedCode) {
		return CodeableConceptSelect("usedCode["+strconv.Itoa(numUsedCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("usedCode["+strconv.Itoa(numUsedCode)+"]", &resource.UsedCode[numUsedCode], optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_PerformerActor(frs []FhirResource, numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"].actor", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"].actor", &resource.Performer[numPerformer].Actor, htmlAttrs)
}
func (resource *Procedure) T_PerformerOnBehalfOf(frs []FhirResource, numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"].onBehalfOf", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"].onBehalfOf", resource.Performer[numPerformer].OnBehalfOf, htmlAttrs)
}
func (resource *Procedure) T_FocalDeviceAction(numFocalDevice int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFocalDevice >= len(resource.FocalDevice) {
		return CodeableConceptSelect("focalDevice["+strconv.Itoa(numFocalDevice)+"].action", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("focalDevice["+strconv.Itoa(numFocalDevice)+"].action", resource.FocalDevice[numFocalDevice].Action, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_FocalDeviceManipulated(frs []FhirResource, numFocalDevice int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFocalDevice >= len(resource.FocalDevice) {
		return ReferenceInput(frs, "focalDevice["+strconv.Itoa(numFocalDevice)+"].manipulated", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "focalDevice["+strconv.Itoa(numFocalDevice)+"].manipulated", &resource.FocalDevice[numFocalDevice].Manipulated, htmlAttrs)
}
