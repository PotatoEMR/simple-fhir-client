package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/OrganizationAffiliation
type OrganizationAffiliation struct {
	Id                        *string                 `json:"id,omitempty"`
	Meta                      *Meta                   `json:"meta,omitempty"`
	ImplicitRules             *string                 `json:"implicitRules,omitempty"`
	Language                  *string                 `json:"language,omitempty"`
	Text                      *Narrative              `json:"text,omitempty"`
	Contained                 []Resource              `json:"contained,omitempty"`
	Extension                 []Extension             `json:"extension,omitempty"`
	ModifierExtension         []Extension             `json:"modifierExtension,omitempty"`
	Identifier                []Identifier            `json:"identifier,omitempty"`
	Active                    *bool                   `json:"active,omitempty"`
	Period                    *Period                 `json:"period,omitempty"`
	Organization              *Reference              `json:"organization,omitempty"`
	ParticipatingOrganization *Reference              `json:"participatingOrganization,omitempty"`
	Network                   []Reference             `json:"network,omitempty"`
	Code                      []CodeableConcept       `json:"code,omitempty"`
	Specialty                 []CodeableConcept       `json:"specialty,omitempty"`
	Location                  []Reference             `json:"location,omitempty"`
	HealthcareService         []Reference             `json:"healthcareService,omitempty"`
	Contact                   []ExtendedContactDetail `json:"contact,omitempty"`
	Endpoint                  []Reference             `json:"endpoint,omitempty"`
}

type OtherOrganizationAffiliation OrganizationAffiliation

// on convert struct to json, automatically add resourceType=OrganizationAffiliation
func (r OrganizationAffiliation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOrganizationAffiliation
		ResourceType string `json:"resourceType"`
	}{
		OtherOrganizationAffiliation: OtherOrganizationAffiliation(r),
		ResourceType:                 "OrganizationAffiliation",
	})
}

func (resource *OrganizationAffiliation) T_Id() templ.Component {

	if resource == nil {
		return StringInput("OrganizationAffiliation.Id", nil)
	}
	return StringInput("OrganizationAffiliation.Id", resource.Id)
}
func (resource *OrganizationAffiliation) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("OrganizationAffiliation.ImplicitRules", nil)
	}
	return StringInput("OrganizationAffiliation.ImplicitRules", resource.ImplicitRules)
}
func (resource *OrganizationAffiliation) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("OrganizationAffiliation.Language", nil, optionsValueSet)
	}
	return CodeSelect("OrganizationAffiliation.Language", resource.Language, optionsValueSet)
}
func (resource *OrganizationAffiliation) T_Active() templ.Component {

	if resource == nil {
		return BoolInput("OrganizationAffiliation.Active", nil)
	}
	return BoolInput("OrganizationAffiliation.Active", resource.Active)
}
func (resource *OrganizationAffiliation) T_Code(numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Code) >= numCode {
		return CodeableConceptSelect("OrganizationAffiliation.Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("OrganizationAffiliation.Code["+strconv.Itoa(numCode)+"]", &resource.Code[numCode], optionsValueSet)
}
func (resource *OrganizationAffiliation) T_Specialty(numSpecialty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Specialty) >= numSpecialty {
		return CodeableConceptSelect("OrganizationAffiliation.Specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("OrganizationAffiliation.Specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet)
}
