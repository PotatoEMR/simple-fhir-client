package r5

//generated with command go run ./bultaoreune
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
		return BoolInput("Active", nil, htmlAttrs)
	}
	return BoolInput("Active", resource.Active, htmlAttrs)
}
func (resource *HealthcareService) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_Type(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_Specialty(numSpecialty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialty >= len(resource.Specialty) {
		return CodeableConceptSelect("Specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *HealthcareService) T_Comment(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Comment", nil, htmlAttrs)
	}
	return StringInput("Comment", resource.Comment, htmlAttrs)
}
func (resource *HealthcareService) T_ExtraDetails(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ExtraDetails", nil, htmlAttrs)
	}
	return StringInput("ExtraDetails", resource.ExtraDetails, htmlAttrs)
}
func (resource *HealthcareService) T_ServiceProvisionCode(numServiceProvisionCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numServiceProvisionCode >= len(resource.ServiceProvisionCode) {
		return CodeableConceptSelect("ServiceProvisionCode["+strconv.Itoa(numServiceProvisionCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ServiceProvisionCode["+strconv.Itoa(numServiceProvisionCode)+"]", &resource.ServiceProvisionCode[numServiceProvisionCode], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_Program(numProgram int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProgram >= len(resource.Program) {
		return CodeableConceptSelect("Program["+strconv.Itoa(numProgram)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Program["+strconv.Itoa(numProgram)+"]", &resource.Program[numProgram], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_Characteristic(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]", &resource.Characteristic[numCharacteristic], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_Communication(numCommunication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCommunication >= len(resource.Communication) {
		return CodeableConceptSelect("Communication["+strconv.Itoa(numCommunication)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Communication["+strconv.Itoa(numCommunication)+"]", &resource.Communication[numCommunication], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_ReferralMethod(numReferralMethod int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReferralMethod >= len(resource.ReferralMethod) {
		return CodeableConceptSelect("ReferralMethod["+strconv.Itoa(numReferralMethod)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ReferralMethod["+strconv.Itoa(numReferralMethod)+"]", &resource.ReferralMethod[numReferralMethod], optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_AppointmentRequired(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("AppointmentRequired", nil, htmlAttrs)
	}
	return BoolInput("AppointmentRequired", resource.AppointmentRequired, htmlAttrs)
}
func (resource *HealthcareService) T_EligibilityCode(numEligibility int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEligibility >= len(resource.Eligibility) {
		return CodeableConceptSelect("Eligibility["+strconv.Itoa(numEligibility)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Eligibility["+strconv.Itoa(numEligibility)+"]Code", resource.Eligibility[numEligibility].Code, optionsValueSet, htmlAttrs)
}
func (resource *HealthcareService) T_EligibilityComment(numEligibility int, htmlAttrs string) templ.Component {
	if resource == nil || numEligibility >= len(resource.Eligibility) {
		return StringInput("Eligibility["+strconv.Itoa(numEligibility)+"]Comment", nil, htmlAttrs)
	}
	return StringInput("Eligibility["+strconv.Itoa(numEligibility)+"]Comment", resource.Eligibility[numEligibility].Comment, htmlAttrs)
}
