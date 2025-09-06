package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Created               *string            `json:"created,omitempty"`
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

func (resource *CarePlan) T_Id() templ.Component {

	if resource == nil {
		return StringInput("CarePlan.Id", nil)
	}
	return StringInput("CarePlan.Id", resource.Id)
}
func (resource *CarePlan) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("CarePlan.ImplicitRules", nil)
	}
	return StringInput("CarePlan.ImplicitRules", resource.ImplicitRules)
}
func (resource *CarePlan) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("CarePlan.Language", nil, optionsValueSet)
	}
	return CodeSelect("CarePlan.Language", resource.Language, optionsValueSet)
}
func (resource *CarePlan) T_InstantiatesCanonical(numInstantiatesCanonical int) templ.Component {

	if resource == nil || len(resource.InstantiatesCanonical) >= numInstantiatesCanonical {
		return StringInput("CarePlan.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil)
	}
	return StringInput("CarePlan.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical])
}
func (resource *CarePlan) T_InstantiatesUri(numInstantiatesUri int) templ.Component {

	if resource == nil || len(resource.InstantiatesUri) >= numInstantiatesUri {
		return StringInput("CarePlan.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil)
	}
	return StringInput("CarePlan.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri])
}
func (resource *CarePlan) T_Status() templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("CarePlan.Status", nil, optionsValueSet)
	}
	return CodeSelect("CarePlan.Status", &resource.Status, optionsValueSet)
}
func (resource *CarePlan) T_Intent() templ.Component {
	optionsValueSet := VSCare_plan_intent

	if resource == nil {
		return CodeSelect("CarePlan.Intent", nil, optionsValueSet)
	}
	return CodeSelect("CarePlan.Intent", &resource.Intent, optionsValueSet)
}
func (resource *CarePlan) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("CarePlan.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CarePlan.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *CarePlan) T_Title() templ.Component {

	if resource == nil {
		return StringInput("CarePlan.Title", nil)
	}
	return StringInput("CarePlan.Title", resource.Title)
}
func (resource *CarePlan) T_Description() templ.Component {

	if resource == nil {
		return StringInput("CarePlan.Description", nil)
	}
	return StringInput("CarePlan.Description", resource.Description)
}
func (resource *CarePlan) T_Created() templ.Component {

	if resource == nil {
		return StringInput("CarePlan.Created", nil)
	}
	return StringInput("CarePlan.Created", resource.Created)
}
func (resource *CarePlan) T_ActivityId(numActivity int) templ.Component {

	if resource == nil || len(resource.Activity) >= numActivity {
		return StringInput("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Id", nil)
	}
	return StringInput("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Id", resource.Activity[numActivity].Id)
}
func (resource *CarePlan) T_ActivityOutcomeCodeableConcept(numActivity int, numOutcomeCodeableConcept int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Activity) >= numActivity || len(resource.Activity[numActivity].OutcomeCodeableConcept) >= numOutcomeCodeableConcept {
		return CodeableConceptSelect("CarePlan.Activity["+strconv.Itoa(numActivity)+"].OutcomeCodeableConcept["+strconv.Itoa(numOutcomeCodeableConcept)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CarePlan.Activity["+strconv.Itoa(numActivity)+"].OutcomeCodeableConcept["+strconv.Itoa(numOutcomeCodeableConcept)+"]", &resource.Activity[numActivity].OutcomeCodeableConcept[numOutcomeCodeableConcept], optionsValueSet)
}
func (resource *CarePlan) T_ActivityDetailId(numActivity int) templ.Component {

	if resource == nil || len(resource.Activity) >= numActivity {
		return StringInput("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.Id", nil)
	}
	return StringInput("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.Id", resource.Activity[numActivity].Detail.Id)
}
func (resource *CarePlan) T_ActivityDetailKind(numActivity int) templ.Component {
	optionsValueSet := VSCare_plan_activity_kind

	if resource == nil || len(resource.Activity) >= numActivity {
		return CodeSelect("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.Kind", nil, optionsValueSet)
	}
	return CodeSelect("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.Kind", resource.Activity[numActivity].Detail.Kind, optionsValueSet)
}
func (resource *CarePlan) T_ActivityDetailInstantiatesCanonical(numActivity int, numInstantiatesCanonical int) templ.Component {

	if resource == nil || len(resource.Activity) >= numActivity || len(resource.Activity[numActivity].Detail.InstantiatesCanonical) >= numInstantiatesCanonical {
		return StringInput("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil)
	}
	return StringInput("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.Activity[numActivity].Detail.InstantiatesCanonical[numInstantiatesCanonical])
}
func (resource *CarePlan) T_ActivityDetailInstantiatesUri(numActivity int, numInstantiatesUri int) templ.Component {

	if resource == nil || len(resource.Activity) >= numActivity || len(resource.Activity[numActivity].Detail.InstantiatesUri) >= numInstantiatesUri {
		return StringInput("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil)
	}
	return StringInput("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.Activity[numActivity].Detail.InstantiatesUri[numInstantiatesUri])
}
func (resource *CarePlan) T_ActivityDetailCode(numActivity int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Activity) >= numActivity {
		return CodeableConceptSelect("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.Code", resource.Activity[numActivity].Detail.Code, optionsValueSet)
}
func (resource *CarePlan) T_ActivityDetailReasonCode(numActivity int, numReasonCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Activity) >= numActivity || len(resource.Activity[numActivity].Detail.ReasonCode) >= numReasonCode {
		return CodeableConceptSelect("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.Activity[numActivity].Detail.ReasonCode[numReasonCode], optionsValueSet)
}
func (resource *CarePlan) T_ActivityDetailStatus(numActivity int) templ.Component {
	optionsValueSet := VSCare_plan_activity_status

	if resource == nil || len(resource.Activity) >= numActivity {
		return CodeSelect("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.Status", nil, optionsValueSet)
	}
	return CodeSelect("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.Status", &resource.Activity[numActivity].Detail.Status, optionsValueSet)
}
func (resource *CarePlan) T_ActivityDetailStatusReason(numActivity int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Activity) >= numActivity {
		return CodeableConceptSelect("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.StatusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.StatusReason", resource.Activity[numActivity].Detail.StatusReason, optionsValueSet)
}
func (resource *CarePlan) T_ActivityDetailDoNotPerform(numActivity int) templ.Component {

	if resource == nil || len(resource.Activity) >= numActivity {
		return BoolInput("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.DoNotPerform", nil)
	}
	return BoolInput("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.DoNotPerform", resource.Activity[numActivity].Detail.DoNotPerform)
}
func (resource *CarePlan) T_ActivityDetailDescription(numActivity int) templ.Component {

	if resource == nil || len(resource.Activity) >= numActivity {
		return StringInput("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.Description", nil)
	}
	return StringInput("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Detail.Description", resource.Activity[numActivity].Detail.Description)
}
