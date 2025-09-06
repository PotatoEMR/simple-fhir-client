package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/HealthcareService
type HealthcareService struct {
	Id                   *string                        `json:"id,omitempty"`
	Meta                 *Meta                          `json:"meta,omitempty"`
	ImplicitRules        *string                        `json:"implicitRules,omitempty"`
	Language             *string                        `json:"language,omitempty"`
	Text                 *Narrative                     `json:"text,omitempty"`
	Contained            []Resource                     `json:"contained,omitempty"`
	Extension            []Extension                    `json:"extension,omitempty"`
	ModifierExtension    []Extension                    `json:"modifierExtension,omitempty"`
	Identifier           []Identifier                   `json:"identifier,omitempty"`
	Active               *bool                          `json:"active,omitempty"`
	ProvidedBy           *Reference                     `json:"providedBy,omitempty"`
	OfferedIn            []Reference                    `json:"offeredIn,omitempty"`
	Category             []CodeableConcept              `json:"category,omitempty"`
	Type                 []CodeableConcept              `json:"type,omitempty"`
	Specialty            []CodeableConcept              `json:"specialty,omitempty"`
	Location             []Reference                    `json:"location,omitempty"`
	Name                 *string                        `json:"name,omitempty"`
	Comment              *string                        `json:"comment,omitempty"`
	ExtraDetails         *string                        `json:"extraDetails,omitempty"`
	Photo                *Attachment                    `json:"photo,omitempty"`
	Contact              []ExtendedContactDetail        `json:"contact,omitempty"`
	CoverageArea         []Reference                    `json:"coverageArea,omitempty"`
	ServiceProvisionCode []CodeableConcept              `json:"serviceProvisionCode,omitempty"`
	Eligibility          []HealthcareServiceEligibility `json:"eligibility,omitempty"`
	Program              []CodeableConcept              `json:"program,omitempty"`
	Characteristic       []CodeableConcept              `json:"characteristic,omitempty"`
	Communication        []CodeableConcept              `json:"communication,omitempty"`
	ReferralMethod       []CodeableConcept              `json:"referralMethod,omitempty"`
	AppointmentRequired  *bool                          `json:"appointmentRequired,omitempty"`
	Availability         []Availability                 `json:"availability,omitempty"`
	Endpoint             []Reference                    `json:"endpoint,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/HealthcareService
type HealthcareServiceEligibility struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Comment           *string          `json:"comment,omitempty"`
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

func (resource *HealthcareService) T_Id() templ.Component {

	if resource == nil {
		return StringInput("HealthcareService.Id", nil)
	}
	return StringInput("HealthcareService.Id", resource.Id)
}
func (resource *HealthcareService) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("HealthcareService.ImplicitRules", nil)
	}
	return StringInput("HealthcareService.ImplicitRules", resource.ImplicitRules)
}
func (resource *HealthcareService) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("HealthcareService.Language", nil, optionsValueSet)
	}
	return CodeSelect("HealthcareService.Language", resource.Language, optionsValueSet)
}
func (resource *HealthcareService) T_Active() templ.Component {

	if resource == nil {
		return BoolInput("HealthcareService.Active", nil)
	}
	return BoolInput("HealthcareService.Active", resource.Active)
}
func (resource *HealthcareService) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("HealthcareService.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("HealthcareService.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *HealthcareService) T_Type(numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Type) >= numType {
		return CodeableConceptSelect("HealthcareService.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("HealthcareService.Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet)
}
func (resource *HealthcareService) T_Specialty(numSpecialty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Specialty) >= numSpecialty {
		return CodeableConceptSelect("HealthcareService.Specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("HealthcareService.Specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet)
}
func (resource *HealthcareService) T_Name() templ.Component {

	if resource == nil {
		return StringInput("HealthcareService.Name", nil)
	}
	return StringInput("HealthcareService.Name", resource.Name)
}
func (resource *HealthcareService) T_Comment() templ.Component {

	if resource == nil {
		return StringInput("HealthcareService.Comment", nil)
	}
	return StringInput("HealthcareService.Comment", resource.Comment)
}
func (resource *HealthcareService) T_ExtraDetails() templ.Component {

	if resource == nil {
		return StringInput("HealthcareService.ExtraDetails", nil)
	}
	return StringInput("HealthcareService.ExtraDetails", resource.ExtraDetails)
}
func (resource *HealthcareService) T_ServiceProvisionCode(numServiceProvisionCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ServiceProvisionCode) >= numServiceProvisionCode {
		return CodeableConceptSelect("HealthcareService.ServiceProvisionCode["+strconv.Itoa(numServiceProvisionCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("HealthcareService.ServiceProvisionCode["+strconv.Itoa(numServiceProvisionCode)+"]", &resource.ServiceProvisionCode[numServiceProvisionCode], optionsValueSet)
}
func (resource *HealthcareService) T_Program(numProgram int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Program) >= numProgram {
		return CodeableConceptSelect("HealthcareService.Program["+strconv.Itoa(numProgram)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("HealthcareService.Program["+strconv.Itoa(numProgram)+"]", &resource.Program[numProgram], optionsValueSet)
}
func (resource *HealthcareService) T_Characteristic(numCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return CodeableConceptSelect("HealthcareService.Characteristic["+strconv.Itoa(numCharacteristic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("HealthcareService.Characteristic["+strconv.Itoa(numCharacteristic)+"]", &resource.Characteristic[numCharacteristic], optionsValueSet)
}
func (resource *HealthcareService) T_Communication(numCommunication int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Communication) >= numCommunication {
		return CodeableConceptSelect("HealthcareService.Communication["+strconv.Itoa(numCommunication)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("HealthcareService.Communication["+strconv.Itoa(numCommunication)+"]", &resource.Communication[numCommunication], optionsValueSet)
}
func (resource *HealthcareService) T_ReferralMethod(numReferralMethod int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ReferralMethod) >= numReferralMethod {
		return CodeableConceptSelect("HealthcareService.ReferralMethod["+strconv.Itoa(numReferralMethod)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("HealthcareService.ReferralMethod["+strconv.Itoa(numReferralMethod)+"]", &resource.ReferralMethod[numReferralMethod], optionsValueSet)
}
func (resource *HealthcareService) T_AppointmentRequired() templ.Component {

	if resource == nil {
		return BoolInput("HealthcareService.AppointmentRequired", nil)
	}
	return BoolInput("HealthcareService.AppointmentRequired", resource.AppointmentRequired)
}
func (resource *HealthcareService) T_EligibilityId(numEligibility int) templ.Component {

	if resource == nil || len(resource.Eligibility) >= numEligibility {
		return StringInput("HealthcareService.Eligibility["+strconv.Itoa(numEligibility)+"].Id", nil)
	}
	return StringInput("HealthcareService.Eligibility["+strconv.Itoa(numEligibility)+"].Id", resource.Eligibility[numEligibility].Id)
}
func (resource *HealthcareService) T_EligibilityCode(numEligibility int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Eligibility) >= numEligibility {
		return CodeableConceptSelect("HealthcareService.Eligibility["+strconv.Itoa(numEligibility)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("HealthcareService.Eligibility["+strconv.Itoa(numEligibility)+"].Code", resource.Eligibility[numEligibility].Code, optionsValueSet)
}
func (resource *HealthcareService) T_EligibilityComment(numEligibility int) templ.Component {

	if resource == nil || len(resource.Eligibility) >= numEligibility {
		return StringInput("HealthcareService.Eligibility["+strconv.Itoa(numEligibility)+"].Comment", nil)
	}
	return StringInput("HealthcareService.Eligibility["+strconv.Itoa(numEligibility)+"].Comment", resource.Eligibility[numEligibility].Comment)
}
