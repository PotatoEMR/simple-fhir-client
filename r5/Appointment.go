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
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Appointment) AppointmentStatus() templ.Component {
	optionsValueSet := VSAppointmentstatus
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *Appointment) AppointmentParticipantStatus(numParticipant int) templ.Component {
	optionsValueSet := VSParticipationstatus
	currentVal := ""
	if resource != nil && len(resource.Participant) >= numParticipant {
		currentVal = resource.Participant[numParticipant].Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
