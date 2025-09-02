package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/PractitionerRole
type PractitionerRole struct {
	Id                     *string                         `json:"id,omitempty"`
	Meta                   *Meta                           `json:"meta,omitempty"`
	ImplicitRules          *string                         `json:"implicitRules,omitempty"`
	Language               *string                         `json:"language,omitempty"`
	Text                   *Narrative                      `json:"text,omitempty"`
	Contained              []Resource                      `json:"contained,omitempty"`
	Extension              []Extension                     `json:"extension,omitempty"`
	ModifierExtension      []Extension                     `json:"modifierExtension,omitempty"`
	Identifier             []Identifier                    `json:"identifier,omitempty"`
	Active                 *bool                           `json:"active,omitempty"`
	Period                 *Period                         `json:"period,omitempty"`
	Practitioner           *Reference                      `json:"practitioner,omitempty"`
	Organization           *Reference                      `json:"organization,omitempty"`
	Code                   []CodeableConcept               `json:"code,omitempty"`
	Specialty              []CodeableConcept               `json:"specialty,omitempty"`
	Location               []Reference                     `json:"location,omitempty"`
	HealthcareService      []Reference                     `json:"healthcareService,omitempty"`
	Telecom                []ContactPoint                  `json:"telecom,omitempty"`
	AvailableTime          []PractitionerRoleAvailableTime `json:"availableTime,omitempty"`
	NotAvailable           []PractitionerRoleNotAvailable  `json:"notAvailable,omitempty"`
	AvailabilityExceptions *string                         `json:"availabilityExceptions,omitempty"`
	Endpoint               []Reference                     `json:"endpoint,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/PractitionerRole
type PractitionerRoleAvailableTime struct {
	Id                 *string     `json:"id,omitempty"`
	Extension          []Extension `json:"extension,omitempty"`
	ModifierExtension  []Extension `json:"modifierExtension,omitempty"`
	DaysOfWeek         []string    `json:"daysOfWeek,omitempty"`
	AllDay             *bool       `json:"allDay,omitempty"`
	AvailableStartTime *string     `json:"availableStartTime,omitempty"`
	AvailableEndTime   *string     `json:"availableEndTime,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/PractitionerRole
type PractitionerRoleNotAvailable struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       string      `json:"description"`
	During            *Period     `json:"during,omitempty"`
}

type OtherPractitionerRole PractitionerRole

// on convert struct to json, automatically add resourceType=PractitionerRole
func (r PractitionerRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPractitionerRole
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitionerRole: OtherPractitionerRole(r),
		ResourceType:          "PractitionerRole",
	})
}

func (resource *PractitionerRole) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *PractitionerRole) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Code[0], optionsValueSet)
}
func (resource *PractitionerRole) T_Specialty(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("specialty", nil, optionsValueSet)
	}
	return CodeableConceptSelect("specialty", &resource.Specialty[0], optionsValueSet)
}
func (resource *PractitionerRole) T_AvailableTimeDaysOfWeek(numAvailableTime int) templ.Component {
	optionsValueSet := VSDays_of_week

	if resource == nil && len(resource.AvailableTime) >= numAvailableTime {
		return CodeSelect("daysOfWeek", nil, optionsValueSet)
	}
	return CodeSelect("daysOfWeek", &resource.AvailableTime[numAvailableTime].DaysOfWeek[0], optionsValueSet)
}
