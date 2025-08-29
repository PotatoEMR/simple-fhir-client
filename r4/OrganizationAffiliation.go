package r4

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *OrganizationAffiliation) OrganizationAffiliationLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
