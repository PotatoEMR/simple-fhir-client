package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/CarePlan
type CarePlan struct {
	Id                    *string            `json:"id,omitempty"`
	Meta                  *Meta              `json:"meta,omitempty"`
	ImplicitRules         *string            `json:"implicitRules,omitempty"`
	Language              *string            `json:"language,omitempty"`
	Text                  *Narrative         `json:"text,omitempty"`
	Contained             []Resource         `json:"contained,omitempty"`
	Extension             []Extension        `json:"extension,omitempty"`
	ModifierExtension     []Extension        `json:"modifierExtension,omitempty"`
	Identifier            []Identifier       `json:"identifier,omitempty"`
	InstantiatesCanonical []string           `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string           `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference        `json:"basedOn,omitempty"`
	Replaces              []Reference        `json:"replaces,omitempty"`
	PartOf                []Reference        `json:"partOf,omitempty"`
	Status                string             `json:"status"`
	Intent                string             `json:"intent"`
	Category              []CodeableConcept  `json:"category,omitempty"`
	Title                 *string            `json:"title,omitempty"`
	Description           *string            `json:"description,omitempty"`
	Subject               Reference          `json:"subject"`
	Encounter             *Reference         `json:"encounter,omitempty"`
	Period                *Period            `json:"period,omitempty"`
	Created               *time.Time         `json:"created,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Author                *Reference         `json:"author,omitempty"`
	Contributor           []Reference        `json:"contributor,omitempty"`
	CareTeam              []Reference        `json:"careTeam,omitempty"`
	Addresses             []Reference        `json:"addresses,omitempty"`
	SupportingInfo        []Reference        `json:"supportingInfo,omitempty"`
	Goal                  []Reference        `json:"goal,omitempty"`
	Activity              []CarePlanActivity `json:"activity,omitempty"`
	Note                  []Annotation       `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CarePlan
type CarePlanActivity struct {
	Id                     *string                 `json:"id,omitempty"`
	Extension              []Extension             `json:"extension,omitempty"`
	ModifierExtension      []Extension             `json:"modifierExtension,omitempty"`
	OutcomeCodeableConcept []CodeableConcept       `json:"outcomeCodeableConcept,omitempty"`
	OutcomeReference       []Reference             `json:"outcomeReference,omitempty"`
	Progress               []Annotation            `json:"progress,omitempty"`
	Reference              *Reference              `json:"reference,omitempty"`
	Detail                 *CarePlanActivityDetail `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CarePlan
type CarePlanActivityDetail struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Kind                   *string           `json:"kind,omitempty"`
	InstantiatesCanonical  []string          `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri        []string          `json:"instantiatesUri,omitempty"`
	Code                   *CodeableConcept  `json:"code,omitempty"`
	ReasonCode             []CodeableConcept `json:"reasonCode,omitempty"`
	ReasonReference        []Reference       `json:"reasonReference,omitempty"`
	Goal                   []Reference       `json:"goal,omitempty"`
	Status                 string            `json:"status"`
	StatusReason           *CodeableConcept  `json:"statusReason,omitempty"`
	DoNotPerform           *bool             `json:"doNotPerform,omitempty"`
	ScheduledTiming        *Timing           `json:"scheduledTiming,omitempty"`
	ScheduledPeriod        *Period           `json:"scheduledPeriod,omitempty"`
	ScheduledString        *string           `json:"scheduledString,omitempty"`
	Location               *Reference        `json:"location,omitempty"`
	Performer              []Reference       `json:"performer,omitempty"`
	ProductCodeableConcept *CodeableConcept  `json:"productCodeableConcept,omitempty"`
	ProductReference       *Reference        `json:"productReference,omitempty"`
	DailyAmount            *Quantity         `json:"dailyAmount,omitempty"`
	Quantity               *Quantity         `json:"quantity,omitempty"`
	Description            *string           `json:"description,omitempty"`
}

type OtherCarePlan CarePlan

// on convert struct to json, automatically add resourceType=CarePlan
func (r CarePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCarePlan
		ResourceType string `json:"resourceType"`
	}{
		OtherCarePlan: OtherCarePlan(r),
		ResourceType:  "CarePlan",
	})
}
func (r CarePlan) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "CarePlan/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "CarePlan"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *CarePlan) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs string) templ.Component {

	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("CarePlan.InstantiatesCanonical."+strconv.Itoa(numInstantiatesCanonical)+".", nil, htmlAttrs)
	}
	return StringInput("CarePlan.InstantiatesCanonical."+strconv.Itoa(numInstantiatesCanonical)+".", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *CarePlan) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs string) templ.Component {

	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("CarePlan.InstantiatesUri."+strconv.Itoa(numInstantiatesUri)+".", nil, htmlAttrs)
	}
	return StringInput("CarePlan.InstantiatesUri."+strconv.Itoa(numInstantiatesUri)+".", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *CarePlan) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("CarePlan.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CarePlan.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_Intent(htmlAttrs string) templ.Component {
	optionsValueSet := VSCare_plan_intent

	if resource == nil {
		return CodeSelect("CarePlan.Intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CarePlan.Intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("CarePlan.Category."+strconv.Itoa(numCategory)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CarePlan.Category."+strconv.Itoa(numCategory)+".", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CarePlan.Title", nil, htmlAttrs)
	}
	return StringInput("CarePlan.Title", resource.Title, htmlAttrs)
}
func (resource *CarePlan) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CarePlan.Description", nil, htmlAttrs)
	}
	return StringInput("CarePlan.Description", resource.Description, htmlAttrs)
}
func (resource *CarePlan) T_Created(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("CarePlan.Created", nil, htmlAttrs)
	}
	return DateTimeInput("CarePlan.Created", resource.Created, htmlAttrs)
}
func (resource *CarePlan) T_Note(numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("CarePlan.Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("CarePlan.Note."+strconv.Itoa(numNote)+".", &resource.Note[numNote], htmlAttrs)
}
func (resource *CarePlan) T_ActivityOutcomeCodeableConcept(numActivity int, numOutcomeCodeableConcept int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numActivity >= len(resource.Activity) || numOutcomeCodeableConcept >= len(resource.Activity[numActivity].OutcomeCodeableConcept) {
		return CodeableConceptSelect("CarePlan.Activity."+strconv.Itoa(numActivity)+"..OutcomeCodeableConcept."+strconv.Itoa(numOutcomeCodeableConcept)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CarePlan.Activity."+strconv.Itoa(numActivity)+"..OutcomeCodeableConcept."+strconv.Itoa(numOutcomeCodeableConcept)+".", &resource.Activity[numActivity].OutcomeCodeableConcept[numOutcomeCodeableConcept], optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_ActivityProgress(numActivity int, numProgress int, htmlAttrs string) templ.Component {

	if resource == nil || numActivity >= len(resource.Activity) || numProgress >= len(resource.Activity[numActivity].Progress) {
		return AnnotationTextArea("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Progress."+strconv.Itoa(numProgress)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Progress."+strconv.Itoa(numProgress)+".", &resource.Activity[numActivity].Progress[numProgress], htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailKind(numActivity int, htmlAttrs string) templ.Component {
	optionsValueSet := VSCare_plan_activity_kind

	if resource == nil || numActivity >= len(resource.Activity) {
		return CodeSelect("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.Kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.Kind", resource.Activity[numActivity].Detail.Kind, optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailInstantiatesCanonical(numActivity int, numInstantiatesCanonical int, htmlAttrs string) templ.Component {

	if resource == nil || numActivity >= len(resource.Activity) || numInstantiatesCanonical >= len(resource.Activity[numActivity].Detail.InstantiatesCanonical) {
		return StringInput("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.InstantiatesCanonical."+strconv.Itoa(numInstantiatesCanonical)+".", nil, htmlAttrs)
	}
	return StringInput("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.InstantiatesCanonical."+strconv.Itoa(numInstantiatesCanonical)+".", &resource.Activity[numActivity].Detail.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailInstantiatesUri(numActivity int, numInstantiatesUri int, htmlAttrs string) templ.Component {

	if resource == nil || numActivity >= len(resource.Activity) || numInstantiatesUri >= len(resource.Activity[numActivity].Detail.InstantiatesUri) {
		return StringInput("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.InstantiatesUri."+strconv.Itoa(numInstantiatesUri)+".", nil, htmlAttrs)
	}
	return StringInput("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.InstantiatesUri."+strconv.Itoa(numInstantiatesUri)+".", &resource.Activity[numActivity].Detail.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailCode(numActivity int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numActivity >= len(resource.Activity) {
		return CodeableConceptSelect("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.Code", resource.Activity[numActivity].Detail.Code, optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailReasonCode(numActivity int, numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numActivity >= len(resource.Activity) || numReasonCode >= len(resource.Activity[numActivity].Detail.ReasonCode) {
		return CodeableConceptSelect("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.ReasonCode."+strconv.Itoa(numReasonCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.ReasonCode."+strconv.Itoa(numReasonCode)+".", &resource.Activity[numActivity].Detail.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailStatus(numActivity int, htmlAttrs string) templ.Component {
	optionsValueSet := VSCare_plan_activity_status

	if resource == nil || numActivity >= len(resource.Activity) {
		return CodeSelect("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.Status", &resource.Activity[numActivity].Detail.Status, optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailStatusReason(numActivity int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numActivity >= len(resource.Activity) {
		return CodeableConceptSelect("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.StatusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.StatusReason", resource.Activity[numActivity].Detail.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailDoNotPerform(numActivity int, htmlAttrs string) templ.Component {

	if resource == nil || numActivity >= len(resource.Activity) {
		return BoolInput("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.DoNotPerform", nil, htmlAttrs)
	}
	return BoolInput("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.DoNotPerform", resource.Activity[numActivity].Detail.DoNotPerform, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailScheduledString(numActivity int, htmlAttrs string) templ.Component {

	if resource == nil || numActivity >= len(resource.Activity) {
		return StringInput("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.ScheduledString", nil, htmlAttrs)
	}
	return StringInput("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.ScheduledString", resource.Activity[numActivity].Detail.ScheduledString, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailProductCodeableConcept(numActivity int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numActivity >= len(resource.Activity) {
		return CodeableConceptSelect("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.ProductCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.ProductCodeableConcept", resource.Activity[numActivity].Detail.ProductCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailDescription(numActivity int, htmlAttrs string) templ.Component {

	if resource == nil || numActivity >= len(resource.Activity) {
		return StringInput("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.Description", nil, htmlAttrs)
	}
	return StringInput("CarePlan.Activity."+strconv.Itoa(numActivity)+"..Detail.Description", resource.Activity[numActivity].Detail.Description, htmlAttrs)
}
