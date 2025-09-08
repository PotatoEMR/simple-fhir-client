package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/OrganizationAffiliation
type OrganizationAffiliation struct {
	Id                        *string           `json:"id,omitempty"`
	Meta                      *Meta             `json:"meta,omitempty"`
	ImplicitRules             *string           `json:"implicitRules,omitempty"`
	Language                  *string           `json:"language,omitempty"`
	Text                      *Narrative        `json:"text,omitempty"`
	Contained                 []Resource        `json:"contained,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	ModifierExtension         []Extension       `json:"modifierExtension,omitempty"`
	Identifier                []Identifier      `json:"identifier,omitempty"`
	Active                    *bool             `json:"active,omitempty"`
	Period                    *Period           `json:"period,omitempty"`
	Organization              *Reference        `json:"organization,omitempty"`
	ParticipatingOrganization *Reference        `json:"participatingOrganization,omitempty"`
	Network                   []Reference       `json:"network,omitempty"`
	Code                      []CodeableConcept `json:"code,omitempty"`
	Specialty                 []CodeableConcept `json:"specialty,omitempty"`
	Location                  []Reference       `json:"location,omitempty"`
	HealthcareService         []Reference       `json:"healthcareService,omitempty"`
	Telecom                   []ContactPoint    `json:"telecom,omitempty"`
	Endpoint                  []Reference       `json:"endpoint,omitempty"`
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
func (r OrganizationAffiliation) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "OrganizationAffiliation/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "OrganizationAffiliation"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *OrganizationAffiliation) T_Active(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("OrganizationAffiliation.Active", nil, htmlAttrs)
	}
	return BoolInput("OrganizationAffiliation.Active", resource.Active, htmlAttrs)
}
func (resource *OrganizationAffiliation) T_Code(numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCode >= len(resource.Code) {
		return CodeableConceptSelect("OrganizationAffiliation.Code."+strconv.Itoa(numCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OrganizationAffiliation.Code."+strconv.Itoa(numCode)+".", &resource.Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *OrganizationAffiliation) T_Specialty(numSpecialty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSpecialty >= len(resource.Specialty) {
		return CodeableConceptSelect("OrganizationAffiliation.Specialty."+strconv.Itoa(numSpecialty)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OrganizationAffiliation.Specialty."+strconv.Itoa(numSpecialty)+".", &resource.Specialty[numSpecialty], optionsValueSet, htmlAttrs)
}
