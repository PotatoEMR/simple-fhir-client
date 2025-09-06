package r5

//generated with command go run ./bultaoreune -nodownload
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
	Created               *string             `json:"created,omitempty"`
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
