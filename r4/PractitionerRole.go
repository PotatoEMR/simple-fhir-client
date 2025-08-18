//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/PractitionerRole
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

// http://hl7.org/fhir/r4/StructureDefinition/PractitionerRole
type PractitionerRoleAvailableTime struct {
	Id                 *string     `json:"id,omitempty"`
	Extension          []Extension `json:"extension,omitempty"`
	ModifierExtension  []Extension `json:"modifierExtension,omitempty"`
	DaysOfWeek         []string    `json:"daysOfWeek,omitempty"`
	AllDay             *bool       `json:"allDay,omitempty"`
	AvailableStartTime *string     `json:"availableStartTime,omitempty"`
	AvailableEndTime   *string     `json:"availableEndTime,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/PractitionerRole
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
