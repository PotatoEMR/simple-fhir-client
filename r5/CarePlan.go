package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/CarePlan
type CarePlan struct {
	Id                    *string             `json:"id,omitempty"`
	Meta                  *Meta               `json:"meta,omitempty"`
	ImplicitRules         *string             `json:"implicitRules,omitempty"`
	Language              *string             `json:"language,omitempty"`
	Text                  *Narrative          `json:"text,omitempty"`
	Contained             []Resource          `json:"contained,omitempty"`
	Extension             []Extension         `json:"extension,omitempty"`
	ModifierExtension     []Extension         `json:"modifierExtension,omitempty"`
	Identifier            []Identifier        `json:"identifier,omitempty"`
	InstantiatesCanonical []string            `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string            `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference         `json:"basedOn,omitempty"`
	Replaces              []Reference         `json:"replaces,omitempty"`
	PartOf                []Reference         `json:"partOf,omitempty"`
	Status                string              `json:"status"`
	Intent                string              `json:"intent"`
	Category              []CodeableConcept   `json:"category,omitempty"`
	Title                 *string             `json:"title,omitempty"`
	Description           *string             `json:"description,omitempty"`
	Subject               Reference           `json:"subject"`
	Encounter             *Reference          `json:"encounter,omitempty"`
	Period                *Period             `json:"period,omitempty"`
	Created               *FhirDateTime       `json:"created,omitempty"`
	Custodian             *Reference          `json:"custodian,omitempty"`
	Contributor           []Reference         `json:"contributor,omitempty"`
	CareTeam              []Reference         `json:"careTeam,omitempty"`
	Addresses             []CodeableReference `json:"addresses,omitempty"`
	SupportingInfo        []Reference         `json:"supportingInfo,omitempty"`
	Goal                  []Reference         `json:"goal,omitempty"`
	Activity              []CarePlanActivity  `json:"activity,omitempty"`
	Note                  []Annotation        `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CarePlan
type CarePlanActivity struct {
	Id                       *string             `json:"id,omitempty"`
	Extension                []Extension         `json:"extension,omitempty"`
	ModifierExtension        []Extension         `json:"modifierExtension,omitempty"`
	PerformedActivity        []CodeableReference `json:"performedActivity,omitempty"`
	Progress                 []Annotation        `json:"progress,omitempty"`
	PlannedActivityReference *Reference          `json:"plannedActivityReference,omitempty"`
}

type OtherCarePlan CarePlan

// struct -> json, automatically add resourceType=Patient
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
func (r CarePlan) ResourceType() string {
	return "CarePlan"
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
func (resource *CarePlan) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *CarePlan) T_Replaces(frs []FhirResource, numReplaces int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReplaces >= len(resource.Replaces) {
		return ReferenceInput(frs, "replaces["+strconv.Itoa(numReplaces)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "replaces["+strconv.Itoa(numReplaces)+"]", &resource.Replaces[numReplaces], htmlAttrs)
}
func (resource *CarePlan) T_PartOf(frs []FhirResource, numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
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
func (resource *CarePlan) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *CarePlan) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
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
func (resource *CarePlan) T_Custodian(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "custodian", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "custodian", resource.Custodian, htmlAttrs)
}
func (resource *CarePlan) T_Contributor(frs []FhirResource, numContributor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContributor >= len(resource.Contributor) {
		return ReferenceInput(frs, "contributor["+strconv.Itoa(numContributor)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "contributor["+strconv.Itoa(numContributor)+"]", &resource.Contributor[numContributor], htmlAttrs)
}
func (resource *CarePlan) T_CareTeam(frs []FhirResource, numCareTeam int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return ReferenceInput(frs, "careTeam["+strconv.Itoa(numCareTeam)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "careTeam["+strconv.Itoa(numCareTeam)+"]", &resource.CareTeam[numCareTeam], htmlAttrs)
}
func (resource *CarePlan) T_Addresses(numAddresses int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddresses >= len(resource.Addresses) {
		return CodeableReferenceInput("addresses["+strconv.Itoa(numAddresses)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("addresses["+strconv.Itoa(numAddresses)+"]", &resource.Addresses[numAddresses], htmlAttrs)
}
func (resource *CarePlan) T_SupportingInfo(frs []FhirResource, numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return ReferenceInput(frs, "supportingInfo["+strconv.Itoa(numSupportingInfo)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "supportingInfo["+strconv.Itoa(numSupportingInfo)+"]", &resource.SupportingInfo[numSupportingInfo], htmlAttrs)
}
func (resource *CarePlan) T_Goal(frs []FhirResource, numGoal int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGoal >= len(resource.Goal) {
		return ReferenceInput(frs, "goal["+strconv.Itoa(numGoal)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "goal["+strconv.Itoa(numGoal)+"]", &resource.Goal[numGoal], htmlAttrs)
}
func (resource *CarePlan) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *CarePlan) T_ActivityPerformedActivity(numActivity int, numPerformedActivity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) || numPerformedActivity >= len(resource.Activity[numActivity].PerformedActivity) {
		return CodeableReferenceInput("activity["+strconv.Itoa(numActivity)+"].performedActivity["+strconv.Itoa(numPerformedActivity)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("activity["+strconv.Itoa(numActivity)+"].performedActivity["+strconv.Itoa(numPerformedActivity)+"]", &resource.Activity[numActivity].PerformedActivity[numPerformedActivity], htmlAttrs)
}
func (resource *CarePlan) T_ActivityProgress(numActivity int, numProgress int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) || numProgress >= len(resource.Activity[numActivity].Progress) {
		return AnnotationTextArea("activity["+strconv.Itoa(numActivity)+"].progress["+strconv.Itoa(numProgress)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("activity["+strconv.Itoa(numActivity)+"].progress["+strconv.Itoa(numProgress)+"]", &resource.Activity[numActivity].Progress[numProgress], htmlAttrs)
}
func (resource *CarePlan) T_ActivityPlannedActivityReference(frs []FhirResource, numActivity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) {
		return ReferenceInput(frs, "activity["+strconv.Itoa(numActivity)+"].plannedActivityReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "activity["+strconv.Itoa(numActivity)+"].plannedActivityReference", resource.Activity[numActivity].PlannedActivityReference, htmlAttrs)
}
