package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Appointment
type Appointment struct {
	Id                     *string                         `json:"id,omitempty"`
	Meta                   *Meta                           `json:"meta,omitempty"`
	ImplicitRules          *string                         `json:"implicitRules,omitempty"`
	Language               *string                         `json:"language,omitempty"`
	Text                   *Narrative                      `json:"text,omitempty"`
	Contained              []Resource                      `json:"contained,omitempty"`
	Extension              []Extension                     `json:"extension,omitempty"`
	ModifierExtension      []Extension                     `json:"modifierExtension,omitempty"`
	Identifier             []Identifier                    `json:"identifier,omitempty"`
	Status                 string                          `json:"status"`
	CancellationReason     *CodeableConcept                `json:"cancellationReason,omitempty"`
	Class                  []CodeableConcept               `json:"class,omitempty"`
	ServiceCategory        []CodeableConcept               `json:"serviceCategory,omitempty"`
	ServiceType            []CodeableReference             `json:"serviceType,omitempty"`
	Specialty              []CodeableConcept               `json:"specialty,omitempty"`
	AppointmentType        *CodeableConcept                `json:"appointmentType,omitempty"`
	Reason                 []CodeableReference             `json:"reason,omitempty"`
	Priority               *CodeableConcept                `json:"priority,omitempty"`
	Description            *string                         `json:"description,omitempty"`
	Replaces               []Reference                     `json:"replaces,omitempty"`
	VirtualService         []VirtualServiceDetail          `json:"virtualService,omitempty"`
	SupportingInformation  []Reference                     `json:"supportingInformation,omitempty"`
	PreviousAppointment    *Reference                      `json:"previousAppointment,omitempty"`
	OriginatingAppointment *Reference                      `json:"originatingAppointment,omitempty"`
	Start                  *string                         `json:"start,omitempty"`
	End                    *string                         `json:"end,omitempty"`
	MinutesDuration        *int                            `json:"minutesDuration,omitempty"`
	RequestedPeriod        []Period                        `json:"requestedPeriod,omitempty"`
	Slot                   []Reference                     `json:"slot,omitempty"`
	Account                []Reference                     `json:"account,omitempty"`
	Created                *string                         `json:"created,omitempty"`
	CancellationDate       *string                         `json:"cancellationDate,omitempty"`
	Note                   []Annotation                    `json:"note,omitempty"`
	PatientInstruction     []CodeableReference             `json:"patientInstruction,omitempty"`
	BasedOn                []Reference                     `json:"basedOn,omitempty"`
	Subject                *Reference                      `json:"subject,omitempty"`
	Participant            []AppointmentParticipant        `json:"participant"`
	RecurrenceId           *int                            `json:"recurrenceId,omitempty"`
	OccurrenceChanged      *bool                           `json:"occurrenceChanged,omitempty"`
	RecurrenceTemplate     []AppointmentRecurrenceTemplate `json:"recurrenceTemplate,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Appointment
type AppointmentParticipant struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              []CodeableConcept `json:"type,omitempty"`
	Period            *Period           `json:"period,omitempty"`
	Actor             *Reference        `json:"actor,omitempty"`
	Required          *bool             `json:"required,omitempty"`
	Status            string            `json:"status"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Appointment
type AppointmentRecurrenceTemplate struct {
	Id                    *string                                       `json:"id,omitempty"`
	Extension             []Extension                                   `json:"extension,omitempty"`
	ModifierExtension     []Extension                                   `json:"modifierExtension,omitempty"`
	Timezone              *CodeableConcept                              `json:"timezone,omitempty"`
	RecurrenceType        CodeableConcept                               `json:"recurrenceType"`
	LastOccurrenceDate    *string                                       `json:"lastOccurrenceDate,omitempty"`
	OccurrenceCount       *int                                          `json:"occurrenceCount,omitempty"`
	OccurrenceDate        []string                                      `json:"occurrenceDate,omitempty"`
	WeeklyTemplate        *AppointmentRecurrenceTemplateWeeklyTemplate  `json:"weeklyTemplate,omitempty"`
	MonthlyTemplate       *AppointmentRecurrenceTemplateMonthlyTemplate `json:"monthlyTemplate,omitempty"`
	YearlyTemplate        *AppointmentRecurrenceTemplateYearlyTemplate  `json:"yearlyTemplate,omitempty"`
	ExcludingDate         []string                                      `json:"excludingDate,omitempty"`
	ExcludingRecurrenceId []int                                         `json:"excludingRecurrenceId,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Appointment
type AppointmentRecurrenceTemplateWeeklyTemplate struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Monday            *bool       `json:"monday,omitempty"`
	Tuesday           *bool       `json:"tuesday,omitempty"`
	Wednesday         *bool       `json:"wednesday,omitempty"`
	Thursday          *bool       `json:"thursday,omitempty"`
	Friday            *bool       `json:"friday,omitempty"`
	Saturday          *bool       `json:"saturday,omitempty"`
	Sunday            *bool       `json:"sunday,omitempty"`
	WeekInterval      *int        `json:"weekInterval,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Appointment
type AppointmentRecurrenceTemplateMonthlyTemplate struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	DayOfMonth        *int        `json:"dayOfMonth,omitempty"`
	NthWeekOfMonth    *Coding     `json:"nthWeekOfMonth,omitempty"`
	DayOfWeek         *Coding     `json:"dayOfWeek,omitempty"`
	MonthInterval     int         `json:"monthInterval"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Appointment
type AppointmentRecurrenceTemplateYearlyTemplate struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	YearInterval      int         `json:"yearInterval"`
}

type OtherAppointment Appointment

// on convert struct to json, automatically add resourceType=Appointment
func (r Appointment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAppointment
		ResourceType string `json:"resourceType"`
	}{
		OtherAppointment: OtherAppointment(r),
		ResourceType:     "Appointment",
	})
}

func (resource *Appointment) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Appointment.Id", nil)
	}
	return StringInput("Appointment.Id", resource.Id)
}
func (resource *Appointment) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Appointment.ImplicitRules", nil)
	}
	return StringInput("Appointment.ImplicitRules", resource.ImplicitRules)
}
func (resource *Appointment) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Appointment.Language", nil, optionsValueSet)
	}
	return CodeSelect("Appointment.Language", resource.Language, optionsValueSet)
}
func (resource *Appointment) T_Status() templ.Component {
	optionsValueSet := VSAppointmentstatus

	if resource == nil {
		return CodeSelect("Appointment.Status", nil, optionsValueSet)
	}
	return CodeSelect("Appointment.Status", &resource.Status, optionsValueSet)
}
func (resource *Appointment) T_CancellationReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Appointment.CancellationReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Appointment.CancellationReason", resource.CancellationReason, optionsValueSet)
}
func (resource *Appointment) T_Class(numClass int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Class) >= numClass {
		return CodeableConceptSelect("Appointment.Class["+strconv.Itoa(numClass)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Appointment.Class["+strconv.Itoa(numClass)+"]", &resource.Class[numClass], optionsValueSet)
}
func (resource *Appointment) T_ServiceCategory(numServiceCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ServiceCategory) >= numServiceCategory {
		return CodeableConceptSelect("Appointment.ServiceCategory["+strconv.Itoa(numServiceCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Appointment.ServiceCategory["+strconv.Itoa(numServiceCategory)+"]", &resource.ServiceCategory[numServiceCategory], optionsValueSet)
}
func (resource *Appointment) T_Specialty(numSpecialty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Specialty) >= numSpecialty {
		return CodeableConceptSelect("Appointment.Specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Appointment.Specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet)
}
func (resource *Appointment) T_AppointmentType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Appointment.AppointmentType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Appointment.AppointmentType", resource.AppointmentType, optionsValueSet)
}
func (resource *Appointment) T_Priority(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Appointment.Priority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Appointment.Priority", resource.Priority, optionsValueSet)
}
func (resource *Appointment) T_Description() templ.Component {

	if resource == nil {
		return StringInput("Appointment.Description", nil)
	}
	return StringInput("Appointment.Description", resource.Description)
}
func (resource *Appointment) T_Start() templ.Component {

	if resource == nil {
		return StringInput("Appointment.Start", nil)
	}
	return StringInput("Appointment.Start", resource.Start)
}
func (resource *Appointment) T_End() templ.Component {

	if resource == nil {
		return StringInput("Appointment.End", nil)
	}
	return StringInput("Appointment.End", resource.End)
}
func (resource *Appointment) T_MinutesDuration() templ.Component {

	if resource == nil {
		return IntInput("Appointment.MinutesDuration", nil)
	}
	return IntInput("Appointment.MinutesDuration", resource.MinutesDuration)
}
func (resource *Appointment) T_Created() templ.Component {

	if resource == nil {
		return StringInput("Appointment.Created", nil)
	}
	return StringInput("Appointment.Created", resource.Created)
}
func (resource *Appointment) T_CancellationDate() templ.Component {

	if resource == nil {
		return StringInput("Appointment.CancellationDate", nil)
	}
	return StringInput("Appointment.CancellationDate", resource.CancellationDate)
}
func (resource *Appointment) T_RecurrenceId() templ.Component {

	if resource == nil {
		return IntInput("Appointment.RecurrenceId", nil)
	}
	return IntInput("Appointment.RecurrenceId", resource.RecurrenceId)
}
func (resource *Appointment) T_OccurrenceChanged() templ.Component {

	if resource == nil {
		return BoolInput("Appointment.OccurrenceChanged", nil)
	}
	return BoolInput("Appointment.OccurrenceChanged", resource.OccurrenceChanged)
}
func (resource *Appointment) T_ParticipantId(numParticipant int) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return StringInput("Appointment.Participant["+strconv.Itoa(numParticipant)+"].Id", nil)
	}
	return StringInput("Appointment.Participant["+strconv.Itoa(numParticipant)+"].Id", resource.Participant[numParticipant].Id)
}
func (resource *Appointment) T_ParticipantType(numParticipant int, numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant || len(resource.Participant[numParticipant].Type) >= numType {
		return CodeableConceptSelect("Appointment.Participant["+strconv.Itoa(numParticipant)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Appointment.Participant["+strconv.Itoa(numParticipant)+"].Type["+strconv.Itoa(numType)+"]", &resource.Participant[numParticipant].Type[numType], optionsValueSet)
}
func (resource *Appointment) T_ParticipantRequired(numParticipant int) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return BoolInput("Appointment.Participant["+strconv.Itoa(numParticipant)+"].Required", nil)
	}
	return BoolInput("Appointment.Participant["+strconv.Itoa(numParticipant)+"].Required", resource.Participant[numParticipant].Required)
}
func (resource *Appointment) T_ParticipantStatus(numParticipant int) templ.Component {
	optionsValueSet := VSParticipationstatus

	if resource == nil || len(resource.Participant) >= numParticipant {
		return CodeSelect("Appointment.Participant["+strconv.Itoa(numParticipant)+"].Status", nil, optionsValueSet)
	}
	return CodeSelect("Appointment.Participant["+strconv.Itoa(numParticipant)+"].Status", &resource.Participant[numParticipant].Status, optionsValueSet)
}
func (resource *Appointment) T_RecurrenceTemplateId(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return StringInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].Id", nil)
	}
	return StringInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].Id", resource.RecurrenceTemplate[numRecurrenceTemplate].Id)
}
func (resource *Appointment) T_RecurrenceTemplateTimezone(numRecurrenceTemplate int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return CodeableConceptSelect("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].Timezone", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].Timezone", resource.RecurrenceTemplate[numRecurrenceTemplate].Timezone, optionsValueSet)
}
func (resource *Appointment) T_RecurrenceTemplateRecurrenceType(numRecurrenceTemplate int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return CodeableConceptSelect("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].RecurrenceType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].RecurrenceType", &resource.RecurrenceTemplate[numRecurrenceTemplate].RecurrenceType, optionsValueSet)
}
func (resource *Appointment) T_RecurrenceTemplateLastOccurrenceDate(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return StringInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].LastOccurrenceDate", nil)
	}
	return StringInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].LastOccurrenceDate", resource.RecurrenceTemplate[numRecurrenceTemplate].LastOccurrenceDate)
}
func (resource *Appointment) T_RecurrenceTemplateOccurrenceCount(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return IntInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].OccurrenceCount", nil)
	}
	return IntInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].OccurrenceCount", resource.RecurrenceTemplate[numRecurrenceTemplate].OccurrenceCount)
}
func (resource *Appointment) T_RecurrenceTemplateOccurrenceDate(numRecurrenceTemplate int, numOccurrenceDate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate || len(resource.RecurrenceTemplate[numRecurrenceTemplate].OccurrenceDate) >= numOccurrenceDate {
		return StringInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].OccurrenceDate["+strconv.Itoa(numOccurrenceDate)+"]", nil)
	}
	return StringInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].OccurrenceDate["+strconv.Itoa(numOccurrenceDate)+"]", &resource.RecurrenceTemplate[numRecurrenceTemplate].OccurrenceDate[numOccurrenceDate])
}
func (resource *Appointment) T_RecurrenceTemplateExcludingDate(numRecurrenceTemplate int, numExcludingDate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate || len(resource.RecurrenceTemplate[numRecurrenceTemplate].ExcludingDate) >= numExcludingDate {
		return StringInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].ExcludingDate["+strconv.Itoa(numExcludingDate)+"]", nil)
	}
	return StringInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].ExcludingDate["+strconv.Itoa(numExcludingDate)+"]", &resource.RecurrenceTemplate[numRecurrenceTemplate].ExcludingDate[numExcludingDate])
}
func (resource *Appointment) T_RecurrenceTemplateExcludingRecurrenceId(numRecurrenceTemplate int, numExcludingRecurrenceId int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate || len(resource.RecurrenceTemplate[numRecurrenceTemplate].ExcludingRecurrenceId) >= numExcludingRecurrenceId {
		return IntInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].ExcludingRecurrenceId["+strconv.Itoa(numExcludingRecurrenceId)+"]", nil)
	}
	return IntInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].ExcludingRecurrenceId["+strconv.Itoa(numExcludingRecurrenceId)+"]", &resource.RecurrenceTemplate[numRecurrenceTemplate].ExcludingRecurrenceId[numExcludingRecurrenceId])
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateId(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return StringInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.Id", nil)
	}
	return StringInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.Id", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Id)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateMonday(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return BoolInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.Monday", nil)
	}
	return BoolInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.Monday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Monday)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateTuesday(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return BoolInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.Tuesday", nil)
	}
	return BoolInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.Tuesday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Tuesday)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateWednesday(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return BoolInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.Wednesday", nil)
	}
	return BoolInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.Wednesday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Wednesday)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateThursday(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return BoolInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.Thursday", nil)
	}
	return BoolInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.Thursday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Thursday)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateFriday(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return BoolInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.Friday", nil)
	}
	return BoolInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.Friday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Friday)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateSaturday(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return BoolInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.Saturday", nil)
	}
	return BoolInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.Saturday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Saturday)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateSunday(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return BoolInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.Sunday", nil)
	}
	return BoolInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.Sunday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Sunday)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateWeekInterval(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return IntInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.WeekInterval", nil)
	}
	return IntInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].WeeklyTemplate.WeekInterval", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.WeekInterval)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateId(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return StringInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].MonthlyTemplate.Id", nil)
	}
	return StringInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].MonthlyTemplate.Id", resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.Id)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateDayOfMonth(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return IntInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].MonthlyTemplate.DayOfMonth", nil)
	}
	return IntInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].MonthlyTemplate.DayOfMonth", resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.DayOfMonth)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateNthWeekOfMonth(numRecurrenceTemplate int) templ.Component {
	optionsValueSet := VSWeek_of_month

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return CodingSelect("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].MonthlyTemplate.NthWeekOfMonth", nil, optionsValueSet)
	}
	return CodingSelect("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].MonthlyTemplate.NthWeekOfMonth", resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.NthWeekOfMonth, optionsValueSet)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateDayOfWeek(numRecurrenceTemplate int) templ.Component {
	optionsValueSet := VSDays_of_week

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return CodingSelect("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].MonthlyTemplate.DayOfWeek", nil, optionsValueSet)
	}
	return CodingSelect("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].MonthlyTemplate.DayOfWeek", resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.DayOfWeek, optionsValueSet)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateMonthInterval(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return IntInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].MonthlyTemplate.MonthInterval", nil)
	}
	return IntInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].MonthlyTemplate.MonthInterval", &resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.MonthInterval)
}
func (resource *Appointment) T_RecurrenceTemplateYearlyTemplateId(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return StringInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].YearlyTemplate.Id", nil)
	}
	return StringInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].YearlyTemplate.Id", resource.RecurrenceTemplate[numRecurrenceTemplate].YearlyTemplate.Id)
}
func (resource *Appointment) T_RecurrenceTemplateYearlyTemplateYearInterval(numRecurrenceTemplate int) templ.Component {

	if resource == nil || len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return IntInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].YearlyTemplate.YearInterval", nil)
	}
	return IntInput("Appointment.RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].YearlyTemplate.YearInterval", &resource.RecurrenceTemplate[numRecurrenceTemplate].YearlyTemplate.YearInterval)
}
