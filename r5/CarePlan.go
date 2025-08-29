package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *CarePlan) CarePlanLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *CarePlan) CarePlanStatus() templ.Component {
	optionsValueSet := VSRequest_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *CarePlan) CarePlanIntent() templ.Component {
	optionsValueSet := VSCare_plan_intent
	currentVal := ""
	if resource != nil {
		currentVal = resource.Intent
	}
	return CodeSelect("intent", currentVal, optionsValueSet)
}
