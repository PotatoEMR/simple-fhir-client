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
		return CodeSelect("Appointment.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Appointment.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_CancellationReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Appointment.CancellationReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Appointment.CancellationReason", resource.CancellationReason, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_Class(numClass int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numClass >= len(resource.Class) {
		return CodeableConceptSelect("Appointment.Class."+strconv.Itoa(numClass)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Appointment.Class."+strconv.Itoa(numClass)+".", &resource.Class[numClass], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_ServiceCategory(numServiceCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numServiceCategory >= len(resource.ServiceCategory) {
		return CodeableConceptSelect("Appointment.ServiceCategory."+strconv.Itoa(numServiceCategory)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Appointment.ServiceCategory."+strconv.Itoa(numServiceCategory)+".", &resource.ServiceCategory[numServiceCategory], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_Specialty(numSpecialty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSpecialty >= len(resource.Specialty) {
		return CodeableConceptSelect("Appointment.Specialty."+strconv.Itoa(numSpecialty)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Appointment.Specialty."+strconv.Itoa(numSpecialty)+".", &resource.Specialty[numSpecialty], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_AppointmentType(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Appointment.AppointmentType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Appointment.AppointmentType", resource.AppointmentType, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_Priority(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Appointment.Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Appointment.Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Appointment.Description", nil, htmlAttrs)
	}
	return StringInput("Appointment.Description", resource.Description, htmlAttrs)
}
func (resource *Appointment) T_Start(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Appointment.Start", nil, htmlAttrs)
	}
	return StringInput("Appointment.Start", resource.Start, htmlAttrs)
}
func (resource *Appointment) T_End(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Appointment.End", nil, htmlAttrs)
	}
	return StringInput("Appointment.End", resource.End, htmlAttrs)
}
func (resource *Appointment) T_MinutesDuration(htmlAttrs string) templ.Component {

	if resource == nil {
		return IntInput("Appointment.MinutesDuration", nil, htmlAttrs)
	}
	return IntInput("Appointment.MinutesDuration", resource.MinutesDuration, htmlAttrs)
}
func (resource *Appointment) T_Created(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Appointment.Created", nil, htmlAttrs)
	}
	return DateTimeInput("Appointment.Created", resource.Created, htmlAttrs)
}
func (resource *Appointment) T_CancellationDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Appointment.CancellationDate", nil, htmlAttrs)
	}
	return DateTimeInput("Appointment.CancellationDate", resource.CancellationDate, htmlAttrs)
}
func (resource *Appointment) T_Note(numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Appointment.Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("Appointment.Note."+strconv.Itoa(numNote)+".", &resource.Note[numNote], htmlAttrs)
}
func (resource *Appointment) T_RecurrenceId(htmlAttrs string) templ.Component {

	if resource == nil {
		return IntInput("Appointment.RecurrenceId", nil, htmlAttrs)
	}
	return IntInput("Appointment.RecurrenceId", resource.RecurrenceId, htmlAttrs)
}
func (resource *Appointment) T_OccurrenceChanged(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("Appointment.OccurrenceChanged", nil, htmlAttrs)
	}
	return BoolInput("Appointment.OccurrenceChanged", resource.OccurrenceChanged, htmlAttrs)
}
func (resource *Appointment) T_ParticipantType(numParticipant int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numParticipant >= len(resource.Participant) || numType >= len(resource.Participant[numParticipant].Type) {
		return CodeableConceptSelect("Appointment.Participant."+strconv.Itoa(numParticipant)+"..Type."+strconv.Itoa(numType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Appointment.Participant."+strconv.Itoa(numParticipant)+"..Type."+strconv.Itoa(numType)+".", &resource.Participant[numParticipant].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_ParticipantRequired(numParticipant int, htmlAttrs string) templ.Component {

	if resource == nil || numParticipant >= len(resource.Participant) {
		return BoolInput("Appointment.Participant."+strconv.Itoa(numParticipant)+"..Required", nil, htmlAttrs)
	}
	return BoolInput("Appointment.Participant."+strconv.Itoa(numParticipant)+"..Required", resource.Participant[numParticipant].Required, htmlAttrs)
}
func (resource *Appointment) T_ParticipantStatus(numParticipant int, htmlAttrs string) templ.Component {
	optionsValueSet := VSParticipationstatus

	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeSelect("Appointment.Participant."+strconv.Itoa(numParticipant)+"..Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Appointment.Participant."+strconv.Itoa(numParticipant)+"..Status", &resource.Participant[numParticipant].Status, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateTimezone(numRecurrenceTemplate int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return CodeableConceptSelect("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..Timezone", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..Timezone", resource.RecurrenceTemplate[numRecurrenceTemplate].Timezone, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateRecurrenceType(numRecurrenceTemplate int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return CodeableConceptSelect("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..RecurrenceType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..RecurrenceType", &resource.RecurrenceTemplate[numRecurrenceTemplate].RecurrenceType, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateLastOccurrenceDate(numRecurrenceTemplate int, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return DateInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..LastOccurrenceDate", nil, htmlAttrs)
	}
	return DateInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..LastOccurrenceDate", resource.RecurrenceTemplate[numRecurrenceTemplate].LastOccurrenceDate, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateOccurrenceCount(numRecurrenceTemplate int, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return IntInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..OccurrenceCount", nil, htmlAttrs)
	}
	return IntInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..OccurrenceCount", resource.RecurrenceTemplate[numRecurrenceTemplate].OccurrenceCount, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateOccurrenceDate(numRecurrenceTemplate int, numOccurrenceDate int, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) || numOccurrenceDate >= len(resource.RecurrenceTemplate[numRecurrenceTemplate].OccurrenceDate) {
		return DateInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..OccurrenceDate."+strconv.Itoa(numOccurrenceDate)+".", nil, htmlAttrs)
	}
	return DateInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..OccurrenceDate."+strconv.Itoa(numOccurrenceDate)+".", &resource.RecurrenceTemplate[numRecurrenceTemplate].OccurrenceDate[numOccurrenceDate], htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateExcludingDate(numRecurrenceTemplate int, numExcludingDate int, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) || numExcludingDate >= len(resource.RecurrenceTemplate[numRecurrenceTemplate].ExcludingDate) {
		return DateInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..ExcludingDate."+strconv.Itoa(numExcludingDate)+".", nil, htmlAttrs)
	}
	return DateInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..ExcludingDate."+strconv.Itoa(numExcludingDate)+".", &resource.RecurrenceTemplate[numRecurrenceTemplate].ExcludingDate[numExcludingDate], htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateExcludingRecurrenceId(numRecurrenceTemplate int, numExcludingRecurrenceId int, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) || numExcludingRecurrenceId >= len(resource.RecurrenceTemplate[numRecurrenceTemplate].ExcludingRecurrenceId) {
		return IntInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..ExcludingRecurrenceId."+strconv.Itoa(numExcludingRecurrenceId)+".", nil, htmlAttrs)
	}
	return IntInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..ExcludingRecurrenceId."+strconv.Itoa(numExcludingRecurrenceId)+".", &resource.RecurrenceTemplate[numRecurrenceTemplate].ExcludingRecurrenceId[numExcludingRecurrenceId], htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateMonday(numRecurrenceTemplate int, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..WeeklyTemplate.Monday", nil, htmlAttrs)
	}
	return BoolInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..WeeklyTemplate.Monday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Monday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateTuesday(numRecurrenceTemplate int, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..WeeklyTemplate.Tuesday", nil, htmlAttrs)
	}
	return BoolInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..WeeklyTemplate.Tuesday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Tuesday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateWednesday(numRecurrenceTemplate int, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..WeeklyTemplate.Wednesday", nil, htmlAttrs)
	}
	return BoolInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..WeeklyTemplate.Wednesday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Wednesday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateThursday(numRecurrenceTemplate int, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..WeeklyTemplate.Thursday", nil, htmlAttrs)
	}
	return BoolInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..WeeklyTemplate.Thursday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Thursday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateFriday(numRecurrenceTemplate int, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..WeeklyTemplate.Friday", nil, htmlAttrs)
	}
	return BoolInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..WeeklyTemplate.Friday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Friday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateSaturday(numRecurrenceTemplate int, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..WeeklyTemplate.Saturday", nil, htmlAttrs)
	}
	return BoolInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..WeeklyTemplate.Saturday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Saturday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateSunday(numRecurrenceTemplate int, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..WeeklyTemplate.Sunday", nil, htmlAttrs)
	}
	return BoolInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..WeeklyTemplate.Sunday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Sunday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateWeekInterval(numRecurrenceTemplate int, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return IntInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..WeeklyTemplate.WeekInterval", nil, htmlAttrs)
	}
	return IntInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..WeeklyTemplate.WeekInterval", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.WeekInterval, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateDayOfMonth(numRecurrenceTemplate int, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return IntInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..MonthlyTemplate.DayOfMonth", nil, htmlAttrs)
	}
	return IntInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..MonthlyTemplate.DayOfMonth", resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.DayOfMonth, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateNthWeekOfMonth(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	optionsValueSet := VSWeek_of_month

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return CodingSelect("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..MonthlyTemplate.NthWeekOfMonth", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..MonthlyTemplate.NthWeekOfMonth", resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.NthWeekOfMonth, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateDayOfWeek(numRecurrenceTemplate int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDays_of_week

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return CodingSelect("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..MonthlyTemplate.DayOfWeek", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..MonthlyTemplate.DayOfWeek", resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.DayOfWeek, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateMonthInterval(numRecurrenceTemplate int, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return IntInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..MonthlyTemplate.MonthInterval", nil, htmlAttrs)
	}
	return IntInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..MonthlyTemplate.MonthInterval", &resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.MonthInterval, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateYearlyTemplateYearInterval(numRecurrenceTemplate int, htmlAttrs string) templ.Component {

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return IntInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..YearlyTemplate.YearInterval", nil, htmlAttrs)
	}
	return IntInput("Appointment.RecurrenceTemplate."+strconv.Itoa(numRecurrenceTemplate)+"..YearlyTemplate.YearInterval", &resource.RecurrenceTemplate[numRecurrenceTemplate].YearlyTemplate.YearInterval, htmlAttrs)
}
