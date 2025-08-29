package r4b

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *HealthcareService) HealthcareServiceLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *HealthcareService) HealthcareServiceAvailableTimeDaysOfWeek(numAvailableTime int) templ.Component {
	optionsValueSet := VSDays_of_week
	currentVal := ""
	if resource != nil && len(resource.AvailableTime) >= numAvailableTime {
		currentVal = resource.AvailableTime[numAvailableTime].DaysOfWeek[0]
	}
	return CodeSelect("daysOfWeek", currentVal, optionsValueSet)
}
