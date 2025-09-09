package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/EpisodeOfCare
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

// http://hl7.org/fhir/r4b/StructureDefinition/EpisodeOfCare
type EpisodeOfCareStatusHistory struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Status            string      `json:"status"`
	Period            Period      `json:"period"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/EpisodeOfCare
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
func (r EpisodeOfCare) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "EpisodeOfCare/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "EpisodeOfCare"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *EpisodeOfCare) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSEpisode_of_care_status

	if resource == nil {
		return CodeSelect("EpisodeOfCare.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("EpisodeOfCare.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EpisodeOfCare) T_Type(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("EpisodeOfCare.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EpisodeOfCare.Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *EpisodeOfCare) T_StatusHistoryStatus(numStatusHistory int, htmlAttrs string) templ.Component {
	optionsValueSet := VSEpisode_of_care_status

	if resource == nil || numStatusHistory >= len(resource.StatusHistory) {
		return CodeSelect("EpisodeOfCare.StatusHistory["+strconv.Itoa(numStatusHistory)+"].Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("EpisodeOfCare.StatusHistory["+strconv.Itoa(numStatusHistory)+"].Status", &resource.StatusHistory[numStatusHistory].Status, optionsValueSet, htmlAttrs)
}
func (resource *EpisodeOfCare) T_DiagnosisRole(numDiagnosis int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("EpisodeOfCare.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EpisodeOfCare.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Role", resource.Diagnosis[numDiagnosis].Role, optionsValueSet, htmlAttrs)
}
func (resource *EpisodeOfCare) T_DiagnosisRank(numDiagnosis int, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return IntInput("EpisodeOfCare.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Rank", nil, htmlAttrs)
	}
	return IntInput("EpisodeOfCare.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Rank", resource.Diagnosis[numDiagnosis].Rank, htmlAttrs)
}
