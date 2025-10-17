package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Procedure
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
	Category              []CodeableConcept      `json:"category,omitempty"`
	Code                  *CodeableConcept       `json:"code,omitempty"`
	Subject               Reference              `json:"subject"`
	Focus                 *Reference             `json:"focus,omitempty"`
	Encounter             *Reference             `json:"encounter,omitempty"`
	OccurrenceDateTime    *FhirDateTime          `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *Period                `json:"occurrencePeriod,omitempty"`
	OccurrenceString      *string                `json:"occurrenceString,omitempty"`
	OccurrenceAge         *Age                   `json:"occurrenceAge,omitempty"`
	OccurrenceRange       *Range                 `json:"occurrenceRange,omitempty"`
	OccurrenceTiming      *Timing                `json:"occurrenceTiming,omitempty"`
	Recorded              *FhirDateTime          `json:"recorded,omitempty"`
	Recorder              *Reference             `json:"recorder,omitempty"`
	ReportedBoolean       *bool                  `json:"reportedBoolean,omitempty"`
	ReportedReference     *Reference             `json:"reportedReference,omitempty"`
	Performer             []ProcedurePerformer   `json:"performer,omitempty"`
	Location              *Reference             `json:"location,omitempty"`
	Reason                []CodeableReference    `json:"reason,omitempty"`
	BodySite              []CodeableConcept      `json:"bodySite,omitempty"`
	Outcome               *CodeableConcept       `json:"outcome,omitempty"`
	Report                []Reference            `json:"report,omitempty"`
	Complication          []CodeableReference    `json:"complication,omitempty"`
	FollowUp              []CodeableConcept      `json:"followUp,omitempty"`
	Note                  []Annotation           `json:"note,omitempty"`
	FocalDevice           []ProcedureFocalDevice `json:"focalDevice,omitempty"`
	Used                  []CodeableReference    `json:"used,omitempty"`
	SupportingInfo        []Reference            `json:"supportingInfo,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Procedure
type ProcedurePerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
	OnBehalfOf        *Reference       `json:"onBehalfOf,omitempty"`
	Period            *Period          `json:"period,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Procedure
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
func (r Procedure) ResourceType() string {
	return "Procedure"
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
func (resource *Procedure) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
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
func (resource *Procedure) T_Focus(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "focus", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "focus", resource.Focus, htmlAttrs)
}
func (resource *Procedure) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *Procedure) T_OccurrenceDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("occurrenceDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("occurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *Procedure) T_OccurrencePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("occurrencePeriod", nil, htmlAttrs)
	}
	return PeriodInput("occurrencePeriod", resource.OccurrencePeriod, htmlAttrs)
}
func (resource *Procedure) T_OccurrenceString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("occurrenceString", nil, htmlAttrs)
	}
	return StringInput("occurrenceString", resource.OccurrenceString, htmlAttrs)
}
func (resource *Procedure) T_OccurrenceAge(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return AgeInput("occurrenceAge", nil, htmlAttrs)
	}
	return AgeInput("occurrenceAge", resource.OccurrenceAge, htmlAttrs)
}
func (resource *Procedure) T_OccurrenceRange(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return RangeInput("occurrenceRange", nil, htmlAttrs)
	}
	return RangeInput("occurrenceRange", resource.OccurrenceRange, htmlAttrs)
}
func (resource *Procedure) T_OccurrenceTiming(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return TimingInput("occurrenceTiming", nil, htmlAttrs)
	}
	return TimingInput("occurrenceTiming", resource.OccurrenceTiming, htmlAttrs)
}
func (resource *Procedure) T_Recorded(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("recorded", nil, htmlAttrs)
	}
	return FhirDateTimeInput("recorded", resource.Recorded, htmlAttrs)
}
func (resource *Procedure) T_Recorder(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "recorder", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "recorder", resource.Recorder, htmlAttrs)
}
func (resource *Procedure) T_ReportedBoolean(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("reportedBoolean", nil, htmlAttrs)
	}
	return BoolInput("reportedBoolean", resource.ReportedBoolean, htmlAttrs)
}
func (resource *Procedure) T_ReportedReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "reportedReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "reportedReference", resource.ReportedReference, htmlAttrs)
}
func (resource *Procedure) T_Location(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "location", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "location", resource.Location, htmlAttrs)
}
func (resource *Procedure) T_Reason(numReason int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], htmlAttrs)
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
func (resource *Procedure) T_Complication(numComplication int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComplication >= len(resource.Complication) {
		return CodeableReferenceInput("complication["+strconv.Itoa(numComplication)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("complication["+strconv.Itoa(numComplication)+"]", &resource.Complication[numComplication], htmlAttrs)
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
func (resource *Procedure) T_Used(numUsed int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUsed >= len(resource.Used) {
		return CodeableReferenceInput("used["+strconv.Itoa(numUsed)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("used["+strconv.Itoa(numUsed)+"]", &resource.Used[numUsed], htmlAttrs)
}
func (resource *Procedure) T_SupportingInfo(frs []FhirResource, numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return ReferenceInput(frs, "supportingInfo["+strconv.Itoa(numSupportingInfo)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "supportingInfo["+strconv.Itoa(numSupportingInfo)+"]", &resource.SupportingInfo[numSupportingInfo], htmlAttrs)
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
func (resource *Procedure) T_PerformerPeriod(numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return PeriodInput("performer["+strconv.Itoa(numPerformer)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("performer["+strconv.Itoa(numPerformer)+"].period", resource.Performer[numPerformer].Period, htmlAttrs)
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
