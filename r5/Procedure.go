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
	OccurrenceDateTime    *time.Time             `json:"occurrenceDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	OccurrencePeriod      *Period                `json:"occurrencePeriod,omitempty"`
	OccurrenceString      *string                `json:"occurrenceString,omitempty"`
	OccurrenceAge         *Age                   `json:"occurrenceAge,omitempty"`
	OccurrenceRange       *Range                 `json:"occurrenceRange,omitempty"`
	OccurrenceTiming      *Timing                `json:"occurrenceTiming,omitempty"`
	Recorded              *time.Time             `json:"recorded,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *Procedure) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Procedure.Category."+strconv.Itoa(numCategory)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Procedure.Category."+strconv.Itoa(numCategory)+".", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Procedure.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Procedure.Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Procedure) T_OccurrenceDateTime(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Procedure.OccurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Procedure.OccurrenceDateTime", resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *Procedure) T_OccurrenceString(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Procedure.OccurrenceString", nil, htmlAttrs)
	}
	return StringInput("Procedure.OccurrenceString", resource.OccurrenceString, htmlAttrs)
}
func (resource *Procedure) T_Recorded(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Procedure.Recorded", nil, htmlAttrs)
	}
	return DateTimeInput("Procedure.Recorded", resource.Recorded, htmlAttrs)
}
func (resource *Procedure) T_ReportedBoolean(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("Procedure.ReportedBoolean", nil, htmlAttrs)
	}
	return BoolInput("Procedure.ReportedBoolean", resource.ReportedBoolean, htmlAttrs)
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
