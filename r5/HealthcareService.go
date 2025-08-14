//generated August 14 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

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
