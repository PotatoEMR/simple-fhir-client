package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

// struct -> json, automatically add resourceType=Patient
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
func (r EpisodeOfCare) ResourceType() string {
	return "EpisodeOfCare"
}

func (resource *EpisodeOfCare) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSEpisode_of_care_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EpisodeOfCare) T_Type(numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *EpisodeOfCare) T_Patient(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "patient", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "patient", &resource.Patient, htmlAttrs)
}
func (resource *EpisodeOfCare) T_ManagingOrganization(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "managingOrganization", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "managingOrganization", resource.ManagingOrganization, htmlAttrs)
}
func (resource *EpisodeOfCare) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *EpisodeOfCare) T_ReferralRequest(frs []FhirResource, numReferralRequest int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReferralRequest >= len(resource.ReferralRequest) {
		return ReferenceInput(frs, "referralRequest["+strconv.Itoa(numReferralRequest)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "referralRequest["+strconv.Itoa(numReferralRequest)+"]", &resource.ReferralRequest[numReferralRequest], htmlAttrs)
}
func (resource *EpisodeOfCare) T_CareManager(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "careManager", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "careManager", resource.CareManager, htmlAttrs)
}
func (resource *EpisodeOfCare) T_Team(frs []FhirResource, numTeam int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTeam >= len(resource.Team) {
		return ReferenceInput(frs, "team["+strconv.Itoa(numTeam)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "team["+strconv.Itoa(numTeam)+"]", &resource.Team[numTeam], htmlAttrs)
}
func (resource *EpisodeOfCare) T_Account(frs []FhirResource, numAccount int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAccount >= len(resource.Account) {
		return ReferenceInput(frs, "account["+strconv.Itoa(numAccount)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "account["+strconv.Itoa(numAccount)+"]", &resource.Account[numAccount], htmlAttrs)
}
func (resource *EpisodeOfCare) T_StatusHistoryStatus(numStatusHistory int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSEpisode_of_care_status

	if resource == nil || numStatusHistory >= len(resource.StatusHistory) {
		return CodeSelect("statusHistory["+strconv.Itoa(numStatusHistory)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("statusHistory["+strconv.Itoa(numStatusHistory)+"].status", &resource.StatusHistory[numStatusHistory].Status, optionsValueSet, htmlAttrs)
}
func (resource *EpisodeOfCare) T_StatusHistoryPeriod(numStatusHistory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatusHistory >= len(resource.StatusHistory) {
		return PeriodInput("statusHistory["+strconv.Itoa(numStatusHistory)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("statusHistory["+strconv.Itoa(numStatusHistory)+"].period", &resource.StatusHistory[numStatusHistory].Period, htmlAttrs)
}
func (resource *EpisodeOfCare) T_DiagnosisCondition(frs []FhirResource, numDiagnosis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return ReferenceInput(frs, "diagnosis["+strconv.Itoa(numDiagnosis)+"].condition", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "diagnosis["+strconv.Itoa(numDiagnosis)+"].condition", &resource.Diagnosis[numDiagnosis].Condition, htmlAttrs)
}
func (resource *EpisodeOfCare) T_DiagnosisRole(numDiagnosis int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].role", resource.Diagnosis[numDiagnosis].Role, optionsValueSet, htmlAttrs)
}
func (resource *EpisodeOfCare) T_DiagnosisRank(numDiagnosis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return IntInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].rank", nil, htmlAttrs)
	}
	return IntInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].rank", resource.Diagnosis[numDiagnosis].Rank, htmlAttrs)
}
