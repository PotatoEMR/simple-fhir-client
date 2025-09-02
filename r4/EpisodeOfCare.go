package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/EpisodeOfCare
type EpisodeOfCare struct {
	Id                   *string                      `json:"id,omitempty"`
	Meta                 *Meta                        `json:"meta,omitempty"`
	ImplicitRules        *string                      `json:"implicitRules,omitempty"`
	Language             *string                      `json:"language,omitempty"`
	Text                 *Narrative                   `json:"text,omitempty"`
	Contained            []Resource                   `json:"contained,omitempty"`
	Extension            []Extension                  `json:"extension,omitempty"`
	ModifierExtension    []Extension                  `json:"modifierExtension,omitempty"`
	Identifier           []Identifier                 `json:"identifier,omitempty"`
	Status               string                       `json:"status"`
	StatusHistory        []EpisodeOfCareStatusHistory `json:"statusHistory,omitempty"`
	Type                 []CodeableConcept            `json:"type,omitempty"`
	Diagnosis            []EpisodeOfCareDiagnosis     `json:"diagnosis,omitempty"`
	Patient              Reference                    `json:"patient"`
	ManagingOrganization *Reference                   `json:"managingOrganization,omitempty"`
	Period               *Period                      `json:"period,omitempty"`
	ReferralRequest      []Reference                  `json:"referralRequest,omitempty"`
	CareManager          *Reference                   `json:"careManager,omitempty"`
	Team                 []Reference                  `json:"team,omitempty"`
	Account              []Reference                  `json:"account,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/EpisodeOfCare
type EpisodeOfCareStatusHistory struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Status            string      `json:"status"`
	Period            Period      `json:"period"`
}

// http://hl7.org/fhir/r4/StructureDefinition/EpisodeOfCare
type EpisodeOfCareDiagnosis struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Condition         Reference        `json:"condition"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Rank              *int             `json:"rank,omitempty"`
}

type OtherEpisodeOfCare EpisodeOfCare

// on convert struct to json, automatically add resourceType=EpisodeOfCare
func (r EpisodeOfCare) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEpisodeOfCare
		ResourceType string `json:"resourceType"`
	}{
		OtherEpisodeOfCare: OtherEpisodeOfCare(r),
		ResourceType:       "EpisodeOfCare",
	})
}

func (resource *EpisodeOfCare) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *EpisodeOfCare) T_Status() templ.Component {
	optionsValueSet := VSEpisode_of_care_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *EpisodeOfCare) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Type[0], optionsValueSet)
}
func (resource *EpisodeOfCare) T_StatusHistoryStatus(numStatusHistory int) templ.Component {
	optionsValueSet := VSEpisode_of_care_status

	if resource == nil && len(resource.StatusHistory) >= numStatusHistory {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.StatusHistory[numStatusHistory].Status, optionsValueSet)
}
func (resource *EpisodeOfCare) T_DiagnosisRole(numDiagnosis int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Diagnosis) >= numDiagnosis {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", resource.Diagnosis[numDiagnosis].Role, optionsValueSet)
}
