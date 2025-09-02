package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	ScheduledTiming        *Timing           `json:"scheduledTiming"`
	ScheduledPeriod        *Period           `json:"scheduledPeriod"`
	ScheduledString        *string           `json:"scheduledString"`
	Location               *Reference        `json:"location,omitempty"`
	Performer              []Reference       `json:"performer,omitempty"`
	ProductCodeableConcept *CodeableConcept  `json:"productCodeableConcept"`
	ProductReference       *Reference        `json:"productReference"`
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

func (resource *CarePlan) CarePlanLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *CarePlan) CarePlanStatus() templ.Component {
	optionsValueSet := VSRequest_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *CarePlan) CarePlanIntent() templ.Component {
	optionsValueSet := VSCare_plan_intent

	if resource != nil {
		return CodeSelect("intent", nil, optionsValueSet)
	}
	return CodeSelect("intent", &resource.Intent, optionsValueSet)
}
func (resource *CarePlan) CarePlanCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *CarePlan) CarePlanActivityOutcomeCodeableConcept(numActivity int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Activity) >= numActivity {
		return CodeableConceptSelect("outcomeCodeableConcept", nil, optionsValueSet)
	}
	return CodeableConceptSelect("outcomeCodeableConcept", &resource.Activity[numActivity].OutcomeCodeableConcept[0], optionsValueSet)
}
func (resource *CarePlan) CarePlanActivityDetailKind(numActivity int) templ.Component {
	optionsValueSet := VSCare_plan_activity_kind

	if resource != nil && len(resource.Activity) >= numActivity {
		return CodeSelect("kind", nil, optionsValueSet)
	}
	return CodeSelect("kind", resource.Activity[numActivity].Detail.Kind, optionsValueSet)
}
func (resource *CarePlan) CarePlanActivityDetailCode(numActivity int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Activity) >= numActivity {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Activity[numActivity].Detail.Code, optionsValueSet)
}
func (resource *CarePlan) CarePlanActivityDetailReasonCode(numActivity int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Activity) >= numActivity {
		return CodeableConceptSelect("reasonCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reasonCode", &resource.Activity[numActivity].Detail.ReasonCode[0], optionsValueSet)
}
func (resource *CarePlan) CarePlanActivityDetailStatus(numActivity int) templ.Component {
	optionsValueSet := VSCare_plan_activity_status

	if resource != nil && len(resource.Activity) >= numActivity {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Activity[numActivity].Detail.Status, optionsValueSet)
}
func (resource *CarePlan) CarePlanActivityDetailStatusReason(numActivity int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Activity) >= numActivity {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("statusReason", resource.Activity[numActivity].Detail.StatusReason, optionsValueSet)
}
