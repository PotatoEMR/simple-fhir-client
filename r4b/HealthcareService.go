package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/HealthcareService
type HealthcareService struct {
	Id                     *string                          `json:"id,omitempty"`
	Meta                   *Meta                            `json:"meta,omitempty"`
	ImplicitRules          *string                          `json:"implicitRules,omitempty"`
	Language               *string                          `json:"language,omitempty"`
	Text                   *Narrative                       `json:"text,omitempty"`
	Contained              []Resource                       `json:"contained,omitempty"`
	Extension              []Extension                      `json:"extension,omitempty"`
	ModifierExtension      []Extension                      `json:"modifierExtension,omitempty"`
	Identifier             []Identifier                     `json:"identifier,omitempty"`
	Active                 *bool                            `json:"active,omitempty"`
	ProvidedBy             *Reference                       `json:"providedBy,omitempty"`
	Category               []CodeableConcept                `json:"category,omitempty"`
	Type                   []CodeableConcept                `json:"type,omitempty"`
	Specialty              []CodeableConcept                `json:"specialty,omitempty"`
	Location               []Reference                      `json:"location,omitempty"`
	Name                   *string                          `json:"name,omitempty"`
	Comment                *string                          `json:"comment,omitempty"`
	ExtraDetails           *string                          `json:"extraDetails,omitempty"`
	Photo                  *Attachment                      `json:"photo,omitempty"`
	Telecom                []ContactPoint                   `json:"telecom,omitempty"`
	CoverageArea           []Reference                      `json:"coverageArea,omitempty"`
	ServiceProvisionCode   []CodeableConcept                `json:"serviceProvisionCode,omitempty"`
	Eligibility            []HealthcareServiceEligibility   `json:"eligibility,omitempty"`
	Program                []CodeableConcept                `json:"program,omitempty"`
	Characteristic         []CodeableConcept                `json:"characteristic,omitempty"`
	Communication          []CodeableConcept                `json:"communication,omitempty"`
	ReferralMethod         []CodeableConcept                `json:"referralMethod,omitempty"`
	AppointmentRequired    *bool                            `json:"appointmentRequired,omitempty"`
	AvailableTime          []HealthcareServiceAvailableTime `json:"availableTime,omitempty"`
	NotAvailable           []HealthcareServiceNotAvailable  `json:"notAvailable,omitempty"`
	AvailabilityExceptions *string                          `json:"availabilityExceptions,omitempty"`
	Endpoint               []Reference                      `json:"endpoint,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/HealthcareService
type HealthcareServiceEligibility struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Comment           *string          `json:"comment,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/HealthcareService
type HealthcareServiceAvailableTime struct {
	Id                 *string     `json:"id,omitempty"`
	Extension          []Extension `json:"extension,omitempty"`
	ModifierExtension  []Extension `json:"modifierExtension,omitempty"`
	DaysOfWeek         []string    `json:"daysOfWeek,omitempty"`
	AllDay             *bool       `json:"allDay,omitempty"`
	AvailableStartTime *string     `json:"availableStartTime,omitempty"`
	AvailableEndTime   *string     `json:"availableEndTime,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/HealthcareService
type HealthcareServiceNotAvailable struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       string      `json:"description"`
	During            *Period     `json:"during,omitempty"`
}

type OtherHealthcareService HealthcareService

// on convert struct to json, automatically add resourceType=HealthcareService
func (r HealthcareService) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherHealthcareService
		ResourceType string `json:"resourceType"`
	}{
		OtherHealthcareService: OtherHealthcareService(r),
		ResourceType:           "HealthcareService",
	})
}
func (r HealthcareService) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "HealthcareService/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "HealthcareService"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *HealthcareService) T_Active(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("active", nil, htmlAttrs)
	}
	return BoolInput("active", resource.Active, htmlAttrs)
}
func (resource *HealthcareService) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_Type(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_Specialty(numSpecialty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialty >= len(resource.Specialty) {
		return CodeableConceptSelect("specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *HealthcareService) T_Comment(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("comment", nil, htmlAttrs)
	}
	return StringInput("comment", resource.Comment, htmlAttrs)
}
func (resource *HealthcareService) T_ExtraDetails(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("extraDetails", nil, htmlAttrs)
	}
	return StringInput("extraDetails", resource.ExtraDetails, htmlAttrs)
}
func (resource *HealthcareService) T_ServiceProvisionCode(numServiceProvisionCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numServiceProvisionCode >= len(resource.ServiceProvisionCode) {
		return CodeableConceptSelect("serviceProvisionCode["+strconv.Itoa(numServiceProvisionCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("serviceProvisionCode["+strconv.Itoa(numServiceProvisionCode)+"]", &resource.ServiceProvisionCode[numServiceProvisionCode], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_Program(numProgram int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProgram >= len(resource.Program) {
		return CodeableConceptSelect("program["+strconv.Itoa(numProgram)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("program["+strconv.Itoa(numProgram)+"]", &resource.Program[numProgram], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_Characteristic(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"]", &resource.Characteristic[numCharacteristic], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_Communication(numCommunication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCommunication >= len(resource.Communication) {
		return CodeableConceptSelect("communication["+strconv.Itoa(numCommunication)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("communication["+strconv.Itoa(numCommunication)+"]", &resource.Communication[numCommunication], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_ReferralMethod(numReferralMethod int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReferralMethod >= len(resource.ReferralMethod) {
		return CodeableConceptSelect("referralMethod["+strconv.Itoa(numReferralMethod)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("referralMethod["+strconv.Itoa(numReferralMethod)+"]", &resource.ReferralMethod[numReferralMethod], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_AppointmentRequired(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("appointmentRequired", nil, htmlAttrs)
	}
	return BoolInput("appointmentRequired", resource.AppointmentRequired, htmlAttrs)
}
func (resource *HealthcareService) T_AvailabilityExceptions(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("availabilityExceptions", nil, htmlAttrs)
	}
	return StringInput("availabilityExceptions", resource.AvailabilityExceptions, htmlAttrs)
}
func (resource *HealthcareService) T_EligibilityCode(numEligibility int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEligibility >= len(resource.Eligibility) {
		return CodeableConceptSelect("eligibility["+strconv.Itoa(numEligibility)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("eligibility["+strconv.Itoa(numEligibility)+"].code", resource.Eligibility[numEligibility].Code, optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_EligibilityComment(numEligibility int, htmlAttrs string) templ.Component {
	if resource == nil || numEligibility >= len(resource.Eligibility) {
		return StringInput("eligibility["+strconv.Itoa(numEligibility)+"].comment", nil, htmlAttrs)
	}
	return StringInput("eligibility["+strconv.Itoa(numEligibility)+"].comment", resource.Eligibility[numEligibility].Comment, htmlAttrs)
}
func (resource *HealthcareService) T_AvailableTimeDaysOfWeek(numAvailableTime int, numDaysOfWeek int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDays_of_week

	if resource == nil || numAvailableTime >= len(resource.AvailableTime) || numDaysOfWeek >= len(resource.AvailableTime[numAvailableTime].DaysOfWeek) {
		return CodeSelect("availableTime["+strconv.Itoa(numAvailableTime)+"].daysOfWeek["+strconv.Itoa(numDaysOfWeek)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("availableTime["+strconv.Itoa(numAvailableTime)+"].daysOfWeek["+strconv.Itoa(numDaysOfWeek)+"]", &resource.AvailableTime[numAvailableTime].DaysOfWeek[numDaysOfWeek], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_AvailableTimeAllDay(numAvailableTime int, htmlAttrs string) templ.Component {
	if resource == nil || numAvailableTime >= len(resource.AvailableTime) {
		return BoolInput("availableTime["+strconv.Itoa(numAvailableTime)+"].allDay", nil, htmlAttrs)
	}
	return BoolInput("availableTime["+strconv.Itoa(numAvailableTime)+"].allDay", resource.AvailableTime[numAvailableTime].AllDay, htmlAttrs)
}
func (resource *HealthcareService) T_AvailableTimeAvailableStartTime(numAvailableTime int, htmlAttrs string) templ.Component {
	if resource == nil || numAvailableTime >= len(resource.AvailableTime) {
		return StringInput("availableTime["+strconv.Itoa(numAvailableTime)+"].availableStartTime", nil, htmlAttrs)
	}
	return StringInput("availableTime["+strconv.Itoa(numAvailableTime)+"].availableStartTime", resource.AvailableTime[numAvailableTime].AvailableStartTime, htmlAttrs)
}
func (resource *HealthcareService) T_AvailableTimeAvailableEndTime(numAvailableTime int, htmlAttrs string) templ.Component {
	if resource == nil || numAvailableTime >= len(resource.AvailableTime) {
		return StringInput("availableTime["+strconv.Itoa(numAvailableTime)+"].availableEndTime", nil, htmlAttrs)
	}
	return StringInput("availableTime["+strconv.Itoa(numAvailableTime)+"].availableEndTime", resource.AvailableTime[numAvailableTime].AvailableEndTime, htmlAttrs)
}
func (resource *HealthcareService) T_NotAvailableDescription(numNotAvailable int, htmlAttrs string) templ.Component {
	if resource == nil || numNotAvailable >= len(resource.NotAvailable) {
		return StringInput("notAvailable["+strconv.Itoa(numNotAvailable)+"].description", nil, htmlAttrs)
	}
	return StringInput("notAvailable["+strconv.Itoa(numNotAvailable)+"].description", &resource.NotAvailable[numNotAvailable].Description, htmlAttrs)
}
