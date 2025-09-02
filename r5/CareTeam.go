package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/CareTeam
type CareTeam struct {
	Id                   *string               `json:"id,omitempty"`
	Meta                 *Meta                 `json:"meta,omitempty"`
	ImplicitRules        *string               `json:"implicitRules,omitempty"`
	Language             *string               `json:"language,omitempty"`
	Text                 *Narrative            `json:"text,omitempty"`
	Contained            []Resource            `json:"contained,omitempty"`
	Extension            []Extension           `json:"extension,omitempty"`
	ModifierExtension    []Extension           `json:"modifierExtension,omitempty"`
	Identifier           []Identifier          `json:"identifier,omitempty"`
	Status               *string               `json:"status,omitempty"`
	Category             []CodeableConcept     `json:"category,omitempty"`
	Name                 *string               `json:"name,omitempty"`
	Subject              *Reference            `json:"subject,omitempty"`
	Period               *Period               `json:"period,omitempty"`
	Participant          []CareTeamParticipant `json:"participant,omitempty"`
	Reason               []CodeableReference   `json:"reason,omitempty"`
	ManagingOrganization []Reference           `json:"managingOrganization,omitempty"`
	Telecom              []ContactPoint        `json:"telecom,omitempty"`
	Note                 []Annotation          `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CareTeam
type CareTeamParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Member            *Reference       `json:"member,omitempty"`
	OnBehalfOf        *Reference       `json:"onBehalfOf,omitempty"`
	CoveragePeriod    *Period          `json:"coveragePeriod"`
	CoverageTiming    *Timing          `json:"coverageTiming"`
}

type OtherCareTeam CareTeam

// on convert struct to json, automatically add resourceType=CareTeam
func (r CareTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCareTeam
		ResourceType string `json:"resourceType"`
	}{
		OtherCareTeam: OtherCareTeam(r),
		ResourceType:  "CareTeam",
	})
}

func (resource *CareTeam) CareTeamLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *CareTeam) CareTeamStatus() templ.Component {
	optionsValueSet := VSCare_team_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", resource.Status, optionsValueSet)
}
func (resource *CareTeam) CareTeamCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *CareTeam) CareTeamParticipantRole(numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Participant) >= numParticipant {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", resource.Participant[numParticipant].Role, optionsValueSet)
}
