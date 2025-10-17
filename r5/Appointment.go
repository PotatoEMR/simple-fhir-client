package r5

//generated with command go run ./bultaoreune
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
	Created                *FhirDateTime                   `json:"created,omitempty"`
	CancellationDate       *FhirDateTime                   `json:"cancellationDate,omitempty"`
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
	LastOccurrenceDate    *FhirDate                                     `json:"lastOccurrenceDate,omitempty"`
	OccurrenceCount       *int                                          `json:"occurrenceCount,omitempty"`
	OccurrenceDate        []FhirDate                                    `json:"occurrenceDate,omitempty"`
	WeeklyTemplate        *AppointmentRecurrenceTemplateWeeklyTemplate  `json:"weeklyTemplate,omitempty"`
	MonthlyTemplate       *AppointmentRecurrenceTemplateMonthlyTemplate `json:"monthlyTemplate,omitempty"`
	YearlyTemplate        *AppointmentRecurrenceTemplateYearlyTemplate  `json:"yearlyTemplate,omitempty"`
	ExcludingDate         []FhirDate                                    `json:"excludingDate,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
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
func (r Appointment) ResourceType() string {
	return "Appointment"
}

func (resource *Appointment) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAppointmentstatus

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_CancellationReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("cancellationReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("cancellationReason", resource.CancellationReason, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_Class(numClass int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClass >= len(resource.Class) {
		return CodeableConceptSelect("class["+strconv.Itoa(numClass)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("class["+strconv.Itoa(numClass)+"]", &resource.Class[numClass], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_ServiceCategory(numServiceCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numServiceCategory >= len(resource.ServiceCategory) {
		return CodeableConceptSelect("serviceCategory["+strconv.Itoa(numServiceCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("serviceCategory["+strconv.Itoa(numServiceCategory)+"]", &resource.ServiceCategory[numServiceCategory], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_ServiceType(numServiceType int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numServiceType >= len(resource.ServiceType) {
		return CodeableReferenceInput("serviceType["+strconv.Itoa(numServiceType)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("serviceType["+strconv.Itoa(numServiceType)+"]", &resource.ServiceType[numServiceType], htmlAttrs)
}
func (resource *Appointment) T_Specialty(numSpecialty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecialty >= len(resource.Specialty) {
		return CodeableConceptSelect("specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_AppointmentType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("appointmentType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("appointmentType", resource.AppointmentType, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_Reason(numReason int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], htmlAttrs)
}
func (resource *Appointment) T_Priority(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *Appointment) T_Replaces(frs []FhirResource, numReplaces int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReplaces >= len(resource.Replaces) {
		return ReferenceInput(frs, "replaces["+strconv.Itoa(numReplaces)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "replaces["+strconv.Itoa(numReplaces)+"]", &resource.Replaces[numReplaces], htmlAttrs)
}
func (resource *Appointment) T_VirtualService(numVirtualService int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVirtualService >= len(resource.VirtualService) {
		return VirtualServiceDetailInput("virtualService["+strconv.Itoa(numVirtualService)+"]", nil, htmlAttrs)
	}
	return VirtualServiceDetailInput("virtualService["+strconv.Itoa(numVirtualService)+"]", &resource.VirtualService[numVirtualService], htmlAttrs)
}
func (resource *Appointment) T_SupportingInformation(frs []FhirResource, numSupportingInformation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInformation >= len(resource.SupportingInformation) {
		return ReferenceInput(frs, "supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", &resource.SupportingInformation[numSupportingInformation], htmlAttrs)
}
func (resource *Appointment) T_PreviousAppointment(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "previousAppointment", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "previousAppointment", resource.PreviousAppointment, htmlAttrs)
}
func (resource *Appointment) T_OriginatingAppointment(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "originatingAppointment", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "originatingAppointment", resource.OriginatingAppointment, htmlAttrs)
}
func (resource *Appointment) T_Start(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("start", nil, htmlAttrs)
	}
	return StringInput("start", resource.Start, htmlAttrs)
}
func (resource *Appointment) T_End(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("end", nil, htmlAttrs)
	}
	return StringInput("end", resource.End, htmlAttrs)
}
func (resource *Appointment) T_MinutesDuration(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("minutesDuration", nil, htmlAttrs)
	}
	return IntInput("minutesDuration", resource.MinutesDuration, htmlAttrs)
}
func (resource *Appointment) T_RequestedPeriod(numRequestedPeriod int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRequestedPeriod >= len(resource.RequestedPeriod) {
		return PeriodInput("requestedPeriod["+strconv.Itoa(numRequestedPeriod)+"]", nil, htmlAttrs)
	}
	return PeriodInput("requestedPeriod["+strconv.Itoa(numRequestedPeriod)+"]", &resource.RequestedPeriod[numRequestedPeriod], htmlAttrs)
}
func (resource *Appointment) T_Slot(frs []FhirResource, numSlot int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSlot >= len(resource.Slot) {
		return ReferenceInput(frs, "slot["+strconv.Itoa(numSlot)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "slot["+strconv.Itoa(numSlot)+"]", &resource.Slot[numSlot], htmlAttrs)
}
func (resource *Appointment) T_Account(frs []FhirResource, numAccount int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAccount >= len(resource.Account) {
		return ReferenceInput(frs, "account["+strconv.Itoa(numAccount)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "account["+strconv.Itoa(numAccount)+"]", &resource.Account[numAccount], htmlAttrs)
}
func (resource *Appointment) T_Created(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("created", nil, htmlAttrs)
	}
	return FhirDateTimeInput("created", resource.Created, htmlAttrs)
}
func (resource *Appointment) T_CancellationDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("cancellationDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("cancellationDate", resource.CancellationDate, htmlAttrs)
}
func (resource *Appointment) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Appointment) T_PatientInstruction(numPatientInstruction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPatientInstruction >= len(resource.PatientInstruction) {
		return CodeableReferenceInput("patientInstruction["+strconv.Itoa(numPatientInstruction)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("patientInstruction["+strconv.Itoa(numPatientInstruction)+"]", &resource.PatientInstruction[numPatientInstruction], htmlAttrs)
}
func (resource *Appointment) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *Appointment) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceId(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("recurrenceId", nil, htmlAttrs)
	}
	return IntInput("recurrenceId", resource.RecurrenceId, htmlAttrs)
}
func (resource *Appointment) T_OccurrenceChanged(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("occurrenceChanged", nil, htmlAttrs)
	}
	return BoolInput("occurrenceChanged", resource.OccurrenceChanged, htmlAttrs)
}
func (resource *Appointment) T_ParticipantType(numParticipant int, numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) || numType >= len(resource.Participant[numParticipant].Type) {
		return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].type["+strconv.Itoa(numType)+"]", &resource.Participant[numParticipant].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_ParticipantPeriod(numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return PeriodInput("participant["+strconv.Itoa(numParticipant)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("participant["+strconv.Itoa(numParticipant)+"].period", resource.Participant[numParticipant].Period, htmlAttrs)
}
func (resource *Appointment) T_ParticipantActor(frs []FhirResource, numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return ReferenceInput(frs, "participant["+strconv.Itoa(numParticipant)+"].actor", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "participant["+strconv.Itoa(numParticipant)+"].actor", resource.Participant[numParticipant].Actor, htmlAttrs)
}
func (resource *Appointment) T_ParticipantRequired(numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return BoolInput("participant["+strconv.Itoa(numParticipant)+"].required", nil, htmlAttrs)
	}
	return BoolInput("participant["+strconv.Itoa(numParticipant)+"].required", resource.Participant[numParticipant].Required, htmlAttrs)
}
func (resource *Appointment) T_ParticipantStatus(numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSParticipationstatus

	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeSelect("participant["+strconv.Itoa(numParticipant)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("participant["+strconv.Itoa(numParticipant)+"].status", &resource.Participant[numParticipant].Status, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateTimezone(numRecurrenceTemplate int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return CodeableConceptSelect("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].timezone", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].timezone", resource.RecurrenceTemplate[numRecurrenceTemplate].Timezone, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateRecurrenceType(numRecurrenceTemplate int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return CodeableConceptSelect("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].recurrenceType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].recurrenceType", &resource.RecurrenceTemplate[numRecurrenceTemplate].RecurrenceType, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateLastOccurrenceDate(numRecurrenceTemplate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return FhirDateInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].lastOccurrenceDate", nil, htmlAttrs)
	}
	return FhirDateInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].lastOccurrenceDate", resource.RecurrenceTemplate[numRecurrenceTemplate].LastOccurrenceDate, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateOccurrenceCount(numRecurrenceTemplate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return IntInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].occurrenceCount", nil, htmlAttrs)
	}
	return IntInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].occurrenceCount", resource.RecurrenceTemplate[numRecurrenceTemplate].OccurrenceCount, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateOccurrenceDate(numRecurrenceTemplate int, numOccurrenceDate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) || numOccurrenceDate >= len(resource.RecurrenceTemplate[numRecurrenceTemplate].OccurrenceDate) {
		return FhirDateInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].occurrenceDate["+strconv.Itoa(numOccurrenceDate)+"]", nil, htmlAttrs)
	}
	return FhirDateInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].occurrenceDate["+strconv.Itoa(numOccurrenceDate)+"]", &resource.RecurrenceTemplate[numRecurrenceTemplate].OccurrenceDate[numOccurrenceDate], htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateExcludingDate(numRecurrenceTemplate int, numExcludingDate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) || numExcludingDate >= len(resource.RecurrenceTemplate[numRecurrenceTemplate].ExcludingDate) {
		return FhirDateInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].excludingDate["+strconv.Itoa(numExcludingDate)+"]", nil, htmlAttrs)
	}
	return FhirDateInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].excludingDate["+strconv.Itoa(numExcludingDate)+"]", &resource.RecurrenceTemplate[numRecurrenceTemplate].ExcludingDate[numExcludingDate], htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateExcludingRecurrenceId(numRecurrenceTemplate int, numExcludingRecurrenceId int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) || numExcludingRecurrenceId >= len(resource.RecurrenceTemplate[numRecurrenceTemplate].ExcludingRecurrenceId) {
		return IntInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].excludingRecurrenceId["+strconv.Itoa(numExcludingRecurrenceId)+"]", nil, htmlAttrs)
	}
	return IntInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].excludingRecurrenceId["+strconv.Itoa(numExcludingRecurrenceId)+"]", &resource.RecurrenceTemplate[numRecurrenceTemplate].ExcludingRecurrenceId[numExcludingRecurrenceId], htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateMonday(numRecurrenceTemplate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].weeklyTemplate.monday", nil, htmlAttrs)
	}
	return BoolInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].weeklyTemplate.monday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Monday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateTuesday(numRecurrenceTemplate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].weeklyTemplate.tuesday", nil, htmlAttrs)
	}
	return BoolInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].weeklyTemplate.tuesday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Tuesday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateWednesday(numRecurrenceTemplate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].weeklyTemplate.wednesday", nil, htmlAttrs)
	}
	return BoolInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].weeklyTemplate.wednesday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Wednesday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateThursday(numRecurrenceTemplate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].weeklyTemplate.thursday", nil, htmlAttrs)
	}
	return BoolInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].weeklyTemplate.thursday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Thursday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateFriday(numRecurrenceTemplate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].weeklyTemplate.friday", nil, htmlAttrs)
	}
	return BoolInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].weeklyTemplate.friday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Friday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateSaturday(numRecurrenceTemplate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].weeklyTemplate.saturday", nil, htmlAttrs)
	}
	return BoolInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].weeklyTemplate.saturday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Saturday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateSunday(numRecurrenceTemplate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return BoolInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].weeklyTemplate.sunday", nil, htmlAttrs)
	}
	return BoolInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].weeklyTemplate.sunday", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.Sunday, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateWeeklyTemplateWeekInterval(numRecurrenceTemplate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return IntInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].weeklyTemplate.weekInterval", nil, htmlAttrs)
	}
	return IntInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].weeklyTemplate.weekInterval", resource.RecurrenceTemplate[numRecurrenceTemplate].WeeklyTemplate.WeekInterval, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateDayOfMonth(numRecurrenceTemplate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return IntInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].monthlyTemplate.dayOfMonth", nil, htmlAttrs)
	}
	return IntInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].monthlyTemplate.dayOfMonth", resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.DayOfMonth, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateNthWeekOfMonth(numRecurrenceTemplate int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSWeek_of_month

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return CodingSelect("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].monthlyTemplate.nthWeekOfMonth", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].monthlyTemplate.nthWeekOfMonth", resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.NthWeekOfMonth, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateDayOfWeek(numRecurrenceTemplate int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDays_of_week

	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return CodingSelect("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].monthlyTemplate.dayOfWeek", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].monthlyTemplate.dayOfWeek", resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.DayOfWeek, optionsValueSet, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateMonthlyTemplateMonthInterval(numRecurrenceTemplate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return IntInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].monthlyTemplate.monthInterval", nil, htmlAttrs)
	}
	return IntInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].monthlyTemplate.monthInterval", &resource.RecurrenceTemplate[numRecurrenceTemplate].MonthlyTemplate.MonthInterval, htmlAttrs)
}
func (resource *Appointment) T_RecurrenceTemplateYearlyTemplateYearInterval(numRecurrenceTemplate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecurrenceTemplate >= len(resource.RecurrenceTemplate) {
		return IntInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].yearlyTemplate.yearInterval", nil, htmlAttrs)
	}
	return IntInput("recurrenceTemplate["+strconv.Itoa(numRecurrenceTemplate)+"].yearlyTemplate.yearInterval", &resource.RecurrenceTemplate[numRecurrenceTemplate].YearlyTemplate.YearInterval, htmlAttrs)
}
