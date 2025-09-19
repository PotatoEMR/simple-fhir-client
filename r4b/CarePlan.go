package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/CarePlan
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
	Created               *FhirDateTime      `json:"created,omitempty"`
	Author                *Reference         `json:"author,omitempty"`
	Contributor           []Reference        `json:"contributor,omitempty"`
	CareTeam              []Reference        `json:"careTeam,omitempty"`
	Addresses             []Reference        `json:"addresses,omitempty"`
	SupportingInfo        []Reference        `json:"supportingInfo,omitempty"`
	Goal                  []Reference        `json:"goal,omitempty"`
	Activity              []CarePlanActivity `json:"activity,omitempty"`
	Note                  []Annotation       `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CarePlan
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

// http://hl7.org/fhir/r4b/StructureDefinition/CarePlan
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
func (resource *CarePlan) T_InstantiatesCanonical(numInstantiatesCanonical int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesCanonical >= len(resource.InstantiatesCanonical) {
		return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *CarePlan) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *CarePlan) T_BasedOn(numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput("basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *CarePlan) T_Replaces(numReplaces int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReplaces >= len(resource.Replaces) {
		return ReferenceInput("replaces["+strconv.Itoa(numReplaces)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("replaces["+strconv.Itoa(numReplaces)+"]", &resource.Replaces[numReplaces], htmlAttrs)
}
func (resource *CarePlan) T_PartOf(numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput("partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *CarePlan) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_Intent(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCare_plan_intent

	if resource == nil {
		return CodeSelect("intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *CarePlan) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *CarePlan) T_Subject(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("subject", nil, htmlAttrs)
	}
	return ReferenceInput("subject", &resource.Subject, htmlAttrs)
}
func (resource *CarePlan) T_Encounter(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("encounter", nil, htmlAttrs)
	}
	return ReferenceInput("encounter", resource.Encounter, htmlAttrs)
}
func (resource *CarePlan) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *CarePlan) T_Created(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("created", nil, htmlAttrs)
	}
	return FhirDateTimeInput("created", resource.Created, htmlAttrs)
}
func (resource *CarePlan) T_Author(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("author", nil, htmlAttrs)
	}
	return ReferenceInput("author", resource.Author, htmlAttrs)
}
func (resource *CarePlan) T_Contributor(numContributor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContributor >= len(resource.Contributor) {
		return ReferenceInput("contributor["+strconv.Itoa(numContributor)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("contributor["+strconv.Itoa(numContributor)+"]", &resource.Contributor[numContributor], htmlAttrs)
}
func (resource *CarePlan) T_CareTeam(numCareTeam int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return ReferenceInput("careTeam["+strconv.Itoa(numCareTeam)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("careTeam["+strconv.Itoa(numCareTeam)+"]", &resource.CareTeam[numCareTeam], htmlAttrs)
}
func (resource *CarePlan) T_Addresses(numAddresses int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddresses >= len(resource.Addresses) {
		return ReferenceInput("addresses["+strconv.Itoa(numAddresses)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("addresses["+strconv.Itoa(numAddresses)+"]", &resource.Addresses[numAddresses], htmlAttrs)
}
func (resource *CarePlan) T_SupportingInfo(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return ReferenceInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"]", &resource.SupportingInfo[numSupportingInfo], htmlAttrs)
}
func (resource *CarePlan) T_Goal(numGoal int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) {
		return ReferenceInput("goal["+strconv.Itoa(numGoal)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("goal["+strconv.Itoa(numGoal)+"]", &resource.Goal[numGoal], htmlAttrs)
}
func (resource *CarePlan) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *CarePlan) T_ActivityOutcomeCodeableConcept(numActivity int, numOutcomeCodeableConcept int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) || numOutcomeCodeableConcept >= len(resource.Activity[numActivity].OutcomeCodeableConcept) {
		return CodeableConceptSelect("activity["+strconv.Itoa(numActivity)+"].outcomeCodeableConcept["+strconv.Itoa(numOutcomeCodeableConcept)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("activity["+strconv.Itoa(numActivity)+"].outcomeCodeableConcept["+strconv.Itoa(numOutcomeCodeableConcept)+"]", &resource.Activity[numActivity].OutcomeCodeableConcept[numOutcomeCodeableConcept], optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_ActivityOutcomeReference(numActivity int, numOutcomeReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) || numOutcomeReference >= len(resource.Activity[numActivity].OutcomeReference) {
		return ReferenceInput("activity["+strconv.Itoa(numActivity)+"].outcomeReference["+strconv.Itoa(numOutcomeReference)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("activity["+strconv.Itoa(numActivity)+"].outcomeReference["+strconv.Itoa(numOutcomeReference)+"]", &resource.Activity[numActivity].OutcomeReference[numOutcomeReference], htmlAttrs)
}
func (resource *CarePlan) T_ActivityProgress(numActivity int, numProgress int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) || numProgress >= len(resource.Activity[numActivity].Progress) {
		return AnnotationTextArea("activity["+strconv.Itoa(numActivity)+"].progress["+strconv.Itoa(numProgress)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("activity["+strconv.Itoa(numActivity)+"].progress["+strconv.Itoa(numProgress)+"]", &resource.Activity[numActivity].Progress[numProgress], htmlAttrs)
}
func (resource *CarePlan) T_ActivityReference(numActivity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) {
		return ReferenceInput("activity["+strconv.Itoa(numActivity)+"].reference", nil, htmlAttrs)
	}
	return ReferenceInput("activity["+strconv.Itoa(numActivity)+"].reference", resource.Activity[numActivity].Reference, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailKind(numActivity int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCare_plan_activity_kind

	if resource == nil || numActivity >= len(resource.Activity) {
		return CodeSelect("activity["+strconv.Itoa(numActivity)+"].detail.kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("activity["+strconv.Itoa(numActivity)+"].detail.kind", resource.Activity[numActivity].Detail.Kind, optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailInstantiatesCanonical(numActivity int, numInstantiatesCanonical int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) || numInstantiatesCanonical >= len(resource.Activity[numActivity].Detail.InstantiatesCanonical) {
		return StringInput("activity["+strconv.Itoa(numActivity)+"].detail.instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("activity["+strconv.Itoa(numActivity)+"].detail.instantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.Activity[numActivity].Detail.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailInstantiatesUri(numActivity int, numInstantiatesUri int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) || numInstantiatesUri >= len(resource.Activity[numActivity].Detail.InstantiatesUri) {
		return StringInput("activity["+strconv.Itoa(numActivity)+"].detail.instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("activity["+strconv.Itoa(numActivity)+"].detail.instantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.Activity[numActivity].Detail.InstantiatesUri[numInstantiatesUri], htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailCode(numActivity int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) {
		return CodeableConceptSelect("activity["+strconv.Itoa(numActivity)+"].detail.code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("activity["+strconv.Itoa(numActivity)+"].detail.code", resource.Activity[numActivity].Detail.Code, optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailReasonCode(numActivity int, numReasonCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) || numReasonCode >= len(resource.Activity[numActivity].Detail.ReasonCode) {
		return CodeableConceptSelect("activity["+strconv.Itoa(numActivity)+"].detail.reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("activity["+strconv.Itoa(numActivity)+"].detail.reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.Activity[numActivity].Detail.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailReasonReference(numActivity int, numReasonReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) || numReasonReference >= len(resource.Activity[numActivity].Detail.ReasonReference) {
		return ReferenceInput("activity["+strconv.Itoa(numActivity)+"].detail.reasonReference["+strconv.Itoa(numReasonReference)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("activity["+strconv.Itoa(numActivity)+"].detail.reasonReference["+strconv.Itoa(numReasonReference)+"]", &resource.Activity[numActivity].Detail.ReasonReference[numReasonReference], htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailGoal(numActivity int, numGoal int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) || numGoal >= len(resource.Activity[numActivity].Detail.Goal) {
		return ReferenceInput("activity["+strconv.Itoa(numActivity)+"].detail.goal["+strconv.Itoa(numGoal)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("activity["+strconv.Itoa(numActivity)+"].detail.goal["+strconv.Itoa(numGoal)+"]", &resource.Activity[numActivity].Detail.Goal[numGoal], htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailStatus(numActivity int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCare_plan_activity_status

	if resource == nil || numActivity >= len(resource.Activity) {
		return CodeSelect("activity["+strconv.Itoa(numActivity)+"].detail.status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("activity["+strconv.Itoa(numActivity)+"].detail.status", &resource.Activity[numActivity].Detail.Status, optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailStatusReason(numActivity int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) {
		return CodeableConceptSelect("activity["+strconv.Itoa(numActivity)+"].detail.statusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("activity["+strconv.Itoa(numActivity)+"].detail.statusReason", resource.Activity[numActivity].Detail.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailDoNotPerform(numActivity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) {
		return BoolInput("activity["+strconv.Itoa(numActivity)+"].detail.doNotPerform", nil, htmlAttrs)
	}
	return BoolInput("activity["+strconv.Itoa(numActivity)+"].detail.doNotPerform", resource.Activity[numActivity].Detail.DoNotPerform, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailScheduledTiming(numActivity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) {
		return TimingInput("activity["+strconv.Itoa(numActivity)+"].detail.scheduledTiming", nil, htmlAttrs)
	}
	return TimingInput("activity["+strconv.Itoa(numActivity)+"].detail.scheduledTiming", resource.Activity[numActivity].Detail.ScheduledTiming, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailScheduledPeriod(numActivity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) {
		return PeriodInput("activity["+strconv.Itoa(numActivity)+"].detail.scheduledPeriod", nil, htmlAttrs)
	}
	return PeriodInput("activity["+strconv.Itoa(numActivity)+"].detail.scheduledPeriod", resource.Activity[numActivity].Detail.ScheduledPeriod, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailScheduledString(numActivity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) {
		return StringInput("activity["+strconv.Itoa(numActivity)+"].detail.scheduledString", nil, htmlAttrs)
	}
	return StringInput("activity["+strconv.Itoa(numActivity)+"].detail.scheduledString", resource.Activity[numActivity].Detail.ScheduledString, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailLocation(numActivity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) {
		return ReferenceInput("activity["+strconv.Itoa(numActivity)+"].detail.location", nil, htmlAttrs)
	}
	return ReferenceInput("activity["+strconv.Itoa(numActivity)+"].detail.location", resource.Activity[numActivity].Detail.Location, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailPerformer(numActivity int, numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) || numPerformer >= len(resource.Activity[numActivity].Detail.Performer) {
		return ReferenceInput("activity["+strconv.Itoa(numActivity)+"].detail.performer["+strconv.Itoa(numPerformer)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("activity["+strconv.Itoa(numActivity)+"].detail.performer["+strconv.Itoa(numPerformer)+"]", &resource.Activity[numActivity].Detail.Performer[numPerformer], htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailProductCodeableConcept(numActivity int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) {
		return CodeableConceptSelect("activity["+strconv.Itoa(numActivity)+"].detail.productCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("activity["+strconv.Itoa(numActivity)+"].detail.productCodeableConcept", resource.Activity[numActivity].Detail.ProductCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailProductReference(numActivity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) {
		return ReferenceInput("activity["+strconv.Itoa(numActivity)+"].detail.productReference", nil, htmlAttrs)
	}
	return ReferenceInput("activity["+strconv.Itoa(numActivity)+"].detail.productReference", resource.Activity[numActivity].Detail.ProductReference, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailDailyAmount(numActivity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) {
		return QuantityInput("activity["+strconv.Itoa(numActivity)+"].detail.dailyAmount", nil, htmlAttrs)
	}
	return QuantityInput("activity["+strconv.Itoa(numActivity)+"].detail.dailyAmount", resource.Activity[numActivity].Detail.DailyAmount, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailQuantity(numActivity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) {
		return QuantityInput("activity["+strconv.Itoa(numActivity)+"].detail.quantity", nil, htmlAttrs)
	}
	return QuantityInput("activity["+strconv.Itoa(numActivity)+"].detail.quantity", resource.Activity[numActivity].Detail.Quantity, htmlAttrs)
}
func (resource *CarePlan) T_ActivityDetailDescription(numActivity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) {
		return StringInput("activity["+strconv.Itoa(numActivity)+"].detail.description", nil, htmlAttrs)
	}
	return StringInput("activity["+strconv.Itoa(numActivity)+"].detail.description", resource.Activity[numActivity].Detail.Description, htmlAttrs)
}
