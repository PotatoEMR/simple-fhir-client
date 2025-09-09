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
	Created               *time.Time          `json:"created,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
		return StringInput("CarePlan.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("CarePlan.InstantiatesCanonical["+strconv.Itoa(numInstantiatesCanonical)+"]", &resource.InstantiatesCanonical[numInstantiatesCanonical], htmlAttrs)
}
func (resource *CarePlan) T_InstantiatesUri(numInstantiatesUri int, htmlAttrs string) templ.Component {
	if resource == nil || numInstantiatesUri >= len(resource.InstantiatesUri) {
		return StringInput("CarePlan.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", nil, htmlAttrs)
	}
	return StringInput("CarePlan.InstantiatesUri["+strconv.Itoa(numInstantiatesUri)+"]", &resource.InstantiatesUri[numInstantiatesUri], htmlAttrs)
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
		return CodeableConceptSelect("CarePlan.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CarePlan.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
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
		return AnnotationTextArea("CarePlan.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("CarePlan.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *CarePlan) T_ActivityProgress(numActivity int, numProgress int, htmlAttrs string) templ.Component {
	if resource == nil || numActivity >= len(resource.Activity) || numProgress >= len(resource.Activity[numActivity].Progress) {
		return AnnotationTextArea("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Progress["+strconv.Itoa(numProgress)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("CarePlan.Activity["+strconv.Itoa(numActivity)+"].Progress["+strconv.Itoa(numProgress)+"]", &resource.Activity[numActivity].Progress[numProgress], htmlAttrs)
}
