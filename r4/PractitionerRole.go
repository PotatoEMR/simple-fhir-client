package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *PractitionerRole) T_Id() templ.Component {

	if resource == nil {
		return StringInput("PractitionerRole.Id", nil)
	}
	return StringInput("PractitionerRole.Id", resource.Id)
}
func (resource *PractitionerRole) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("PractitionerRole.ImplicitRules", nil)
	}
	return StringInput("PractitionerRole.ImplicitRules", resource.ImplicitRules)
}
func (resource *PractitionerRole) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("PractitionerRole.Language", nil, optionsValueSet)
	}
	return CodeSelect("PractitionerRole.Language", resource.Language, optionsValueSet)
}
func (resource *PractitionerRole) T_Active() templ.Component {

	if resource == nil {
		return BoolInput("PractitionerRole.Active", nil)
	}
	return BoolInput("PractitionerRole.Active", resource.Active)
}
func (resource *PractitionerRole) T_Code(numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Code) >= numCode {
		return CodeableConceptSelect("PractitionerRole.Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PractitionerRole.Code["+strconv.Itoa(numCode)+"]", &resource.Code[numCode], optionsValueSet)
}
func (resource *PractitionerRole) T_Specialty(numSpecialty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Specialty) >= numSpecialty {
		return CodeableConceptSelect("PractitionerRole.Specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("PractitionerRole.Specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet)
}
func (resource *PractitionerRole) T_AvailabilityExceptions() templ.Component {

	if resource == nil {
		return StringInput("PractitionerRole.AvailabilityExceptions", nil)
	}
	return StringInput("PractitionerRole.AvailabilityExceptions", resource.AvailabilityExceptions)
}
func (resource *PractitionerRole) T_AvailableTimeId(numAvailableTime int) templ.Component {

	if resource == nil || len(resource.AvailableTime) >= numAvailableTime {
		return StringInput("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].Id", nil)
	}
	return StringInput("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].Id", resource.AvailableTime[numAvailableTime].Id)
}
func (resource *PractitionerRole) T_AvailableTimeDaysOfWeek(numAvailableTime int, numDaysOfWeek int) templ.Component {
	optionsValueSet := VSDays_of_week

	if resource == nil || len(resource.AvailableTime) >= numAvailableTime || len(resource.AvailableTime[numAvailableTime].DaysOfWeek) >= numDaysOfWeek {
		return CodeSelect("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].DaysOfWeek["+strconv.Itoa(numDaysOfWeek)+"]", nil, optionsValueSet)
	}
	return CodeSelect("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].DaysOfWeek["+strconv.Itoa(numDaysOfWeek)+"]", &resource.AvailableTime[numAvailableTime].DaysOfWeek[numDaysOfWeek], optionsValueSet)
}
func (resource *PractitionerRole) T_AvailableTimeAllDay(numAvailableTime int) templ.Component {

	if resource == nil || len(resource.AvailableTime) >= numAvailableTime {
		return BoolInput("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].AllDay", nil)
	}
	return BoolInput("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].AllDay", resource.AvailableTime[numAvailableTime].AllDay)
}
func (resource *PractitionerRole) T_AvailableTimeAvailableStartTime(numAvailableTime int) templ.Component {

	if resource == nil || len(resource.AvailableTime) >= numAvailableTime {
		return StringInput("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].AvailableStartTime", nil)
	}
	return StringInput("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].AvailableStartTime", resource.AvailableTime[numAvailableTime].AvailableStartTime)
}
func (resource *PractitionerRole) T_AvailableTimeAvailableEndTime(numAvailableTime int) templ.Component {

	if resource == nil || len(resource.AvailableTime) >= numAvailableTime {
		return StringInput("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].AvailableEndTime", nil)
	}
	return StringInput("PractitionerRole.AvailableTime["+strconv.Itoa(numAvailableTime)+"].AvailableEndTime", resource.AvailableTime[numAvailableTime].AvailableEndTime)
}
func (resource *PractitionerRole) T_NotAvailableId(numNotAvailable int) templ.Component {

	if resource == nil || len(resource.NotAvailable) >= numNotAvailable {
		return StringInput("PractitionerRole.NotAvailable["+strconv.Itoa(numNotAvailable)+"].Id", nil)
	}
	return StringInput("PractitionerRole.NotAvailable["+strconv.Itoa(numNotAvailable)+"].Id", resource.NotAvailable[numNotAvailable].Id)
}
func (resource *PractitionerRole) T_NotAvailableDescription(numNotAvailable int) templ.Component {

	if resource == nil || len(resource.NotAvailable) >= numNotAvailable {
		return StringInput("PractitionerRole.NotAvailable["+strconv.Itoa(numNotAvailable)+"].Description", nil)
	}
	return StringInput("PractitionerRole.NotAvailable["+strconv.Itoa(numNotAvailable)+"].Description", &resource.NotAvailable[numNotAvailable].Description)
}
