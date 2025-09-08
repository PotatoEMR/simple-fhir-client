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
	Created                *time.Time                      `json:"created,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	CancellationDate       *time.Time                      `json:"cancellationDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	LastOccurrenceDate    *time.Time                                    `json:"lastOccurrenceDate,omitempty,format:'2006-01-02'"`
	OccurrenceCount       *int                                          `json:"occurrenceCount,omitempty"`
	OccurrenceDate        []time.Time                                   `json:"occurrenceDate,omitempty,format:'2006-01-02'"`
	WeeklyTemplate        *AppointmentRecurrenceTemplateWeeklyTemplate  `json:"weeklyTemplate,omitempty"`
	MonthlyTemplate       *AppointmentRecurrenceTemplateMonthlyTemplate `json:"monthlyTemplate,omitempty"`
	YearlyTemplate        *AppointmentRecurrenceTemplateYearlyTemplate  `json:"yearlyTemplate,omitempty"`
	ExcludingDate         []time.Time                                   `json:"excludingDate,omitempty,format:'2006-01-02'"`
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
func (r Appointment) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Appointment/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Appointment"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Appointment) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSAppointmentstatus

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_CancellationReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("CancellationReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CancellationReason", resource.CancellationReason, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_Class(numClass int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numClass >= len(resource.Class) {
		return CodeableConceptSelect("Class["+strconv.Itoa(numClass)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Class["+strconv.Itoa(numClass)+"]", &resource.Class[numClass], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_ServiceCategory(numServiceCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numServiceCategory >= len(resource.ServiceCategory) {
		return CodeableConceptSelect("ServiceCategory["+strconv.Itoa(numServiceCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ServiceCategory["+strconv.Itoa(numServiceCategory)+"]", &resource.ServiceCategory[numServiceCategory], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_Specialty(numSpecialty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialty >= len(resource.Specialty) {
		return CodeableConceptSelect("Specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_AppointmentType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("AppointmentType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AppointmentType", resource.AppointmentType, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_Priority(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *Appointment) T_Start(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Start", nil, htmlAttrs)
	}
	return StringInput("Start", resource.Start, htmlAttrs)
}
func (resource *Appointment) T_End(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("End", nil, htmlAttrs)
	}
	return StringInput("End", resource.End, htmlAttrs)
}
func (resource *Appointment) T_MinutesDuration(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("MinutesDuration", nil, htmlAttrs)
	}
	return IntInput("MinutesDuration", resource.MinutesDuration, htmlAttrs)
}
func (resource *Appointment) T_Created(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Created", nil, htmlAttrs)
	}
	return DateTimeInput("Created", resource.Created, htmlAttrs)
}
func (resource *Appointment) T_CancellationDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("CancellationDate", nil, htmlAttrs)
	}
	return DateTimeInput("CancellationDate", resource.CancellationDate, htmlAttrs)
}
func (resource *Appointment) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Appointment) T_RecurrenceId(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("RecurrenceId", nil, htmlAttrs)
	}
	return IntInput("RecurrenceId", resource.RecurrenceId, htmlAttrs)
}
func (resource *Appointment) T_OccurrenceChanged(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("OccurrenceChanged", nil, htmlAttrs)
	}
	return BoolInput("OccurrenceChanged", resource.OccurrenceChanged, htmlAttrs)
}
func (resource *Appointment) T_ParticipantType(numParticipant int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) || numType >= len(resource.Participant[numParticipant].Type) {
		return CodeableConceptSelect("Participant["+strconv.Itoa(numParticipant)+"]Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Participant["+strconv.Itoa(numParticipant)+"]Type["+strconv.Itoa(numType)+"]", &resource.Participant[numParticipant].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_ParticipantRequired(numParticipant int, htmlAttrs string) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return BoolInput("Participant["+strconv.Itoa(numParticipant)+"]Required", nil, htmlAttrs)
	}
	return BoolInput("Participant["+strconv.Itoa(numParticipant)+"]Required", resource.Participant[numParticipant].Required, htmlAttrs)
}
func (resource *Appointment) T_ParticipantStatus(numParticipant int, htmlAttrs string) templ.Component {
	optionsValueSet := VSParticipationstatus

	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeSelect("Participant["+strconv.Itoa(numParticipant)+"]Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Participant["+strconv.Itoa(numParticipant)+"]Status", &resource.Participant[numParticipant].Status, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateTimezone(numRecurrenceTemplate int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return CodeableConceptSelect("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]Timezone", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]Timezone", resource.RecurrenceTemplate[numRecurrenceTemplate].Timezone, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateRecurrenceType(numRecurrenceTemplate int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return CodeableConceptSelect("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]RecurrenceType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]RecurrenceType", &resource.RecurrenceTemplate[numRecurrenceTemplate].RecurrenceType, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateLastOccurrenceDate(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return DateInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]LastOccurrenceDate", nil, htmlAttrs)
	}
	return DateInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]LastOccurrenceDate", resource.RecurrenceTemplate[numRecurrenceTemplate].LastOccurrenceDate, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateOccurrenceCount(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return IntInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]OccurrenceCount", nil, htmlAttrs)
	}
	return IntInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]OccurrenceCount", resource.RecurrenceTemplate[numRecurrenceTemplate].OccurrenceCount, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateOccurrenceDate(numRecurrenceTemplate int, numOccurrenceDate int, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) || numOccurrenceDate >= len(resource.RecurrenceTemplate[numRecurrenceTemplate].OccurrenceDate) {
		return DateInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]OccurrenceDate["+strconv.Itoa(numOccurrenceDate)+"]", nil, htmlAttrs)
	}
	return DateInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]OccurrenceDate["+strconv.Itoa(numOccurrenceDate)+"]", &resource.RecurrenceTemplate[numRecurrenceTemplate].OccurrenceDate[numOccurrenceDate], htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateExcludingDate(numRecurrenceTemplate int, numExcludingDate int, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) || numExcludingDate >= len(resource.RecurrenceTemplate[numRecurrenceTemplate].ExcludingDate) {
		return DateInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]ExcludingDate["+strconv.Itoa(numExcludingDate)+"]", nil, htmlAttrs)
	}
	return DateInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]ExcludingDate["+strconv.Itoa(numExcludingDate)+"]", &resource.RecurrenceTemplate[numRecurrenceTemplate].ExcludingDate[numExcludingDate], htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateExcludingRecurrenceId(numRecurrenceTemplate int, numExcludingRecurrenceId int, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) || numExcludingRecurrenceId >= len(resource.RecurrenceTemplate[numRecurrenceTemplate].ExcludingRecurrenceId) {
		return IntInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]ExcludingRecurrenceId["+strconv.Itoa(numExcludingRecurrenceId)+"]", nil, htmlAttrs)
	}
	return IntInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]ExcludingRecurrenceId["+strconv.Itoa(numExcludingRecurrenceId)+"]", &resource.RecurrenceTemplate[numRecurrenceTemplate].ExcludingRecurrenceId[numExcludingRecurrenceId], htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateMonday(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]WeeklyTemplate.Monday", nil, htmlAttrs)
	}
	return BoolInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]WeeklyTemplate.Monday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Monday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateTuesday(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]WeeklyTemplate.Tuesday", nil, htmlAttrs)
	}
	return BoolInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]WeeklyTemplate.Tuesday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Tuesday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateWednesday(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]WeeklyTemplate.Wednesday", nil, htmlAttrs)
	}
	return BoolInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]WeeklyTemplate.Wednesday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Wednesday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateThursday(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]WeeklyTemplate.Thursday", nil, htmlAttrs)
	}
	return BoolInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]WeeklyTemplate.Thursday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Thursday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateFriday(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]WeeklyTemplate.Friday", nil, htmlAttrs)
	}
	return BoolInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]WeeklyTemplate.Friday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Friday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateSaturday(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]WeeklyTemplate.Saturday", nil, htmlAttrs)
	}
	return BoolInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]WeeklyTemplate.Saturday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Saturday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateSunday(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]WeeklyTemplate.Sunday", nil, htmlAttrs)
	}
	return BoolInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]WeeklyTemplate.Sunday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Sunday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateWeekInterval(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return IntInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]WeeklyTemplate.WeekInterval", nil, htmlAttrs)
	}
	return IntInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]WeeklyTemplate.WeekInterval", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.WeekInterval, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateDayOfMonth(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return IntInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]MonthlyTemplate.DayOfMonth", nil, htmlAttrs)
	}
	return IntInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]MonthlyTemplate.DayOfMonth", resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.DayOfMonth, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateNthWeekOfMonth(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	optionsValueSet := VSWeek_of_month

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return CodingSelect("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]MonthlyTemplate.NthWeekOfMonth", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]MonthlyTemplate.NthWeekOfMonth", resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.NthWeekOfMonth, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateDayOfWeek(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDays_of_week

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return CodingSelect("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]MonthlyTemplate.DayOfWeek", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]MonthlyTemplate.DayOfWeek", resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.DayOfWeek, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateMonthInterval(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return IntInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]MonthlyTemplate.MonthInterval", nil, htmlAttrs)
	}
	return IntInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]MonthlyTemplate.MonthInterval", &resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.MonthInterval, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateYearlyTemplateYearInterval(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return IntInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]YearlyTemplate.YearInterval", nil, htmlAttrs)
	}
	return IntInput("RecurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"]YearlyTemplate.YearInterval", &resource.RecurrenceTemplate[numRecurrenceTemplate].YearlyTemplate.YearInterval, htmlAttrs)
}
