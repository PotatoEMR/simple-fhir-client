package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *Appointment) AppointmentLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Appointment) AppointmentStatus() templ.Component {
	optionsValueSet := VSAppointmentstatus

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Appointment) AppointmentCancellationReason(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("cancellationReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("cancellationReason", resource.CancellationReason, optionsValueSet)
}
func (resource *Appointment) AppointmentClass(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("class", nil, optionsValueSet)
	}
	return CodeableConceptSelect("class", &resource.Class[0], optionsValueSet)
}
func (resource *Appointment) AppointmentServiceCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("serviceCategory", nil, optionsValueSet)
	}
	return CodeableConceptSelect("serviceCategory", &resource.ServiceCategory[0], optionsValueSet)
}
func (resource *Appointment) AppointmentSpecialty(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("specialty", nil, optionsValueSet)
	}
	return CodeableConceptSelect("specialty", &resource.Specialty[0], optionsValueSet)
}
func (resource *Appointment) AppointmentAppointmentType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("appointmentType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("appointmentType", resource.AppointmentType, optionsValueSet)
}
func (resource *Appointment) AppointmentPriority(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("priority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("priority", resource.Priority, optionsValueSet)
}
func (resource *Appointment) AppointmentParticipantType(numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Participant) >= numParticipant {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Participant[numParticipant].Type[0], optionsValueSet)
}
func (resource *Appointment) AppointmentParticipantStatus(numParticipant int) templ.Component {
	optionsValueSet := VSParticipationstatus

	if resource != nil && len(resource.Participant) >= numParticipant {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Participant[numParticipant].Status, optionsValueSet)
}
func (resource *Appointment) AppointmentRecurrenceTemplateTimezone(numRecurrenceTemplate int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return CodeableConceptSelect("timezone", nil, optionsValueSet)
	}
	return CodeableConceptSelect("timezone", resource.RecurrenceTemplate[numRecurrenceTemplate].Timezone, optionsValueSet)
}
func (resource *Appointment) AppointmentRecurrenceTemplateRecurrenceType(numRecurrenceTemplate int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return CodeableConceptSelect("recurrenceType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("recurrenceType", &resource.RecurrenceTemplate[numRecurrenceTemplate].RecurrenceType, optionsValueSet)
}
func (resource *Appointment) AppointmentRecurrenceTemplateMonthlyTemplateNthWeekOfMonth(numRecurrenceTemplate int) templ.Component {
	optionsValueSet := VSWeek_of_month

	if resource != nil && len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return CodingSelect("nthWeekOfMonth", nil, optionsValueSet)
	}
	return CodingSelect("nthWeekOfMonth", resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.NthWeekOfMonth, optionsValueSet)
}
func (resource *Appointment) AppointmentRecurrenceTemplateMonthlyTemplateDayOfWeek(numRecurrenceTemplate int) templ.Component {
	optionsValueSet := VSDays_of_week

	if resource != nil && len(resource.RecurrenceTemplate) >= numRecurrenceTemplate {
		return CodingSelect("dayOfWeek", nil, optionsValueSet)
	}
	return CodingSelect("dayOfWeek", resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.DayOfWeek, optionsValueSet)
}
