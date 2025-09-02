package r4b

//generated with command go run ./bultaoreune -nodownload
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

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *HealthcareService) HealthcareServiceCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *HealthcareService) HealthcareServiceType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Type[0], optionsValueSet)
}
func (resource *HealthcareService) HealthcareServiceSpecialty(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("specialty", nil, optionsValueSet)
	}
	return CodeableConceptSelect("specialty", &resource.Specialty[0], optionsValueSet)
}
func (resource *HealthcareService) HealthcareServiceServiceProvisionCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("serviceProvisionCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("serviceProvisionCode", &resource.ServiceProvisionCode[0], optionsValueSet)
}
func (resource *HealthcareService) HealthcareServiceProgram(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("program", nil, optionsValueSet)
	}
	return CodeableConceptSelect("program", &resource.Program[0], optionsValueSet)
}
func (resource *HealthcareService) HealthcareServiceCharacteristic(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("characteristic", nil, optionsValueSet)
	}
	return CodeableConceptSelect("characteristic", &resource.Characteristic[0], optionsValueSet)
}
func (resource *HealthcareService) HealthcareServiceCommunication(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("communication", nil, optionsValueSet)
	}
	return CodeableConceptSelect("communication", &resource.Communication[0], optionsValueSet)
}
func (resource *HealthcareService) HealthcareServiceReferralMethod(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("referralMethod", nil, optionsValueSet)
	}
	return CodeableConceptSelect("referralMethod", &resource.ReferralMethod[0], optionsValueSet)
}
func (resource *HealthcareService) HealthcareServiceEligibilityCode(numEligibility int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Eligibility) >= numEligibility {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Eligibility[numEligibility].Code, optionsValueSet)
}
func (resource *HealthcareService) HealthcareServiceAvailableTimeDaysOfWeek(numAvailableTime int) templ.Component {
	optionsValueSet := VSDays_of_week

	if resource != nil && len(resource.AvailableTime) >= numAvailableTime {
		return CodeSelect("daysOfWeek", nil, optionsValueSet)
	}
	return CodeSelect("daysOfWeek", &resource.AvailableTime[numAvailableTime].DaysOfWeek[0], optionsValueSet)
}
