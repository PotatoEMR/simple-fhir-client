package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
func (r PractitionerRole) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "PractitionerRole/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "PractitionerRole"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *PractitionerRole) T_Active(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("PractitionerRole.Active", nil, htmlAttrs)
	}
	return BoolInput("PractitionerRole.Active", resource.Active, htmlAttrs)
}
func (resource *PractitionerRole) T_Code(numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return CodeableConceptSelect("PractitionerRole.Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PractitionerRole.Code["+strconv.Itoa(numCode)+"]", &resource.Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *PractitionerRole) T_Specialty(numSpecialty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialty >= len(resource.Specialty) {
		return CodeableConceptSelect("PractitionerRole.Specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PractitionerRole.Specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet, htmlAttrs)
}
func (resource *PractitionerRole) T_AvailabilityExceptions(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PractitionerRole.AvailabilityExceptions", nil, htmlAttrs)
	}
	return StringInput("PractitionerRole.AvailabilityExceptions", resource.AvailabilityExceptions, htmlAttrs)
}
func (resource *PractitionerRole) T_AvailableTimeDaysOfWeek(numAvailableTime int, numDaysOfWeek int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDays_of_week

	if resource == nil || numAvailableTime >= len(resource.AvailableTime) || numDaysOfWeek >= len(resource.AvailableTime[numAvailableTime].DaysOfWeek) {
		return CodeSelect("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].DaysOfWeek["+strconv.Itoa(numDaysOfWeek)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].DaysOfWeek["+strconv.Itoa(numDaysOfWeek)+"]", &resource.AvailableTime[numAvailableTime].DaysOfWeek[numDaysOfWeek], optionsValueSet, htmlAttrs)
}
func (resource *PractitionerRole) T_AvailableTimeAllDay(numAvailableTime int, htmlAttrs string) templ.Component {
	if resource == nil || numAvailableTime >= len(resource.AvailableTime) {
		return BoolInput("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].AllDay", nil, htmlAttrs)
	}
	return BoolInput("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].AllDay", resource.AvailableTime[numAvailableTime].AllDay, htmlAttrs)
}
func (resource *PractitionerRole) T_AvailableTimeAvailableStartTime(numAvailableTime int, htmlAttrs string) templ.Component {
	if resource == nil || numAvailableTime >= len(resource.AvailableTime) {
		return StringInput("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].AvailableStartTime", nil, htmlAttrs)
	}
	return StringInput("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].AvailableStartTime", resource.AvailableTime[numAvailableTime].AvailableStartTime, htmlAttrs)
}
func (resource *PractitionerRole) T_AvailableTimeAvailableEndTime(numAvailableTime int, htmlAttrs string) templ.Component {
	if resource == nil || numAvailableTime >= len(resource.AvailableTime) {
		return StringInput("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].AvailableEndTime", nil, htmlAttrs)
	}
	return StringInput("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].AvailableEndTime", resource.AvailableTime[numAvailableTime].AvailableEndTime, htmlAttrs)
}
func (resource *PractitionerRole) T_NotAvailableDescription(numNotAvailable int, htmlAttrs string) templ.Component {
	if resource == nil || numNotAvailable >= len(resource.NotAvailable) {
		return StringInput("PractitionerRole.NotAvailable["+strconv.Itoa(numNotAvailable)+"].Description", nil, htmlAttrs)
	}
	return StringInput("PractitionerRole.NotAvailable["+strconv.Itoa(numNotAvailable)+"].Description", &resource.NotAvailable[numNotAvailable].Description, htmlAttrs)
}
