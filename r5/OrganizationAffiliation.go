package r5

//generated with command go run ./bultaoreune
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
func (resource *OrganizationAffiliation) T_Active(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("active", nil, htmlAttrs)
	}
	return BoolInput("active", resource.Active, htmlAttrs)
}
func (resource *OrganizationAffiliation) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *OrganizationAffiliation) T_Organization(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "organization", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "organization", resource.Organization, htmlAttrs)
}
func (resource *OrganizationAffiliation) T_ParticipatingOrganization(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "participatingOrganization", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "participatingOrganization", resource.ParticipatingOrganization, htmlAttrs)
}
func (resource *OrganizationAffiliation) T_Network(frs []FhirResource, numNetwork int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNetwork >= len(resource.Network) {
		return ReferenceInput(frs, "network["+strconv.Itoa(numNetwork)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "network["+strconv.Itoa(numNetwork)+"]", &resource.Network[numNetwork], htmlAttrs)
}
func (resource *OrganizationAffiliation) T_Code(numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return CodeableConceptSelect("code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code["+strconv.Itoa(numCode)+"]", &resource.Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *OrganizationAffiliation) T_Specialty(numSpecialty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecialty >= len(resource.Specialty) {
		return CodeableConceptSelect("specialty["+strconv.Itoa(numSpecialty)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("specialty["+strconv.Itoa(numSpecialty)+"]", &resource.Specialty[numSpecialty], optionsValueSet, htmlAttrs)
}
func (resource *OrganizationAffiliation) T_Location(frs []FhirResource, numLocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLocation >= len(resource.Location) {
		return ReferenceInput(frs, "location["+strconv.Itoa(numLocation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "location["+strconv.Itoa(numLocation)+"]", &resource.Location[numLocation], htmlAttrs)
}
func (resource *OrganizationAffiliation) T_HealthcareService(frs []FhirResource, numHealthcareService int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numHealthcareService >= len(resource.HealthcareService) {
		return ReferenceInput(frs, "healthcareService["+strconv.Itoa(numHealthcareService)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "healthcareService["+strconv.Itoa(numHealthcareService)+"]", &resource.HealthcareService[numHealthcareService], htmlAttrs)
}
func (resource *OrganizationAffiliation) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ExtendedContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ExtendedContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *OrganizationAffiliation) T_Endpoint(frs []FhirResource, numEndpoint int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndpoint >= len(resource.Endpoint) {
		return ReferenceInput(frs, "endpoint["+strconv.Itoa(numEndpoint)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "endpoint["+strconv.Itoa(numEndpoint)+"]", &resource.Endpoint[numEndpoint], htmlAttrs)
}
