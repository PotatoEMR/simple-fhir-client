package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/EpisodeOfCare
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
	Reason               []EpisodeOfCareReason        `json:"reason,omitempty"`
	Diagnosis            []EpisodeOfCareDiagnosis     `json:"diagnosis,omitempty"`
	Patient              Reference                    `json:"patient"`
	ManagingOrganization *Reference                   `json:"managingOrganization,omitempty"`
	Period               *Period                      `json:"period,omitempty"`
	ReferralRequest      []Reference                  `json:"referralRequest,omitempty"`
	CareManager          *Reference                   `json:"careManager,omitempty"`
	CareTeam             []Reference                  `json:"careTeam,omitempty"`
	Account              []Reference                  `json:"account,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EpisodeOfCare
type EpisodeOfCareStatusHistory struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Status            string      `json:"status"`
	Period            Period      `json:"period"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EpisodeOfCare
type EpisodeOfCareReason struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Use               *CodeableConcept    `json:"use,omitempty"`
	Value             []CodeableReference `json:"value,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EpisodeOfCare
type EpisodeOfCareDiagnosis struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Condition         []CodeableReference `json:"condition,omitempty"`
	Use               *CodeableConcept    `json:"use,omitempty"`
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

func (resource *EpisodeOfCare) T_Id() templ.Component {

	if resource == nil {
		return StringInput("EpisodeOfCare.Id", nil)
	}
	return StringInput("EpisodeOfCare.Id", resource.Id)
}
func (resource *EpisodeOfCare) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("EpisodeOfCare.ImplicitRules", nil)
	}
	return StringInput("EpisodeOfCare.ImplicitRules", resource.ImplicitRules)
}
func (resource *EpisodeOfCare) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("EpisodeOfCare.Language", nil, optionsValueSet)
	}
	return CodeSelect("EpisodeOfCare.Language", resource.Language, optionsValueSet)
}
func (resource *EpisodeOfCare) T_Status() templ.Component {
	optionsValueSet := VSEpisode_of_care_status

	if resource == nil {
		return CodeSelect("EpisodeOfCare.Status", nil, optionsValueSet)
	}
	return CodeSelect("EpisodeOfCare.Status", &resource.Status, optionsValueSet)
}
func (resource *EpisodeOfCare) T_Type(numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Type) >= numType {
		return CodeableConceptSelect("EpisodeOfCare.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EpisodeOfCare.Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet)
}
func (resource *EpisodeOfCare) T_StatusHistoryId(numStatusHistory int) templ.Component {

	if resource == nil || len(resource.StatusHistory) >= numStatusHistory {
		return StringInput("EpisodeOfCare.StatusHistory["+strconv.Itoa(numStatusHistory)+"].Id", nil)
	}
	return StringInput("EpisodeOfCare.StatusHistory["+strconv.Itoa(numStatusHistory)+"].Id", resource.StatusHistory[numStatusHistory].Id)
}
func (resource *EpisodeOfCare) T_StatusHistoryStatus(numStatusHistory int) templ.Component {
	optionsValueSet := VSEpisode_of_care_status

	if resource == nil || len(resource.StatusHistory) >= numStatusHistory {
		return CodeSelect("EpisodeOfCare.StatusHistory["+strconv.Itoa(numStatusHistory)+"].Status", nil, optionsValueSet)
	}
	return CodeSelect("EpisodeOfCare.StatusHistory["+strconv.Itoa(numStatusHistory)+"].Status", &resource.StatusHistory[numStatusHistory].Status, optionsValueSet)
}
func (resource *EpisodeOfCare) T_ReasonId(numReason int) templ.Component {

	if resource == nil || len(resource.Reason) >= numReason {
		return StringInput("EpisodeOfCare.Reason["+strconv.Itoa(numReason)+"].Id", nil)
	}
	return StringInput("EpisodeOfCare.Reason["+strconv.Itoa(numReason)+"].Id", resource.Reason[numReason].Id)
}
func (resource *EpisodeOfCare) T_ReasonUse(numReason int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Reason) >= numReason {
		return CodeableConceptSelect("EpisodeOfCare.Reason["+strconv.Itoa(numReason)+"].Use", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EpisodeOfCare.Reason["+strconv.Itoa(numReason)+"].Use", resource.Reason[numReason].Use, optionsValueSet)
}
func (resource *EpisodeOfCare) T_DiagnosisId(numDiagnosis int) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis {
		return StringInput("EpisodeOfCare.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Id", nil)
	}
	return StringInput("EpisodeOfCare.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Id", resource.Diagnosis[numDiagnosis].Id)
}
func (resource *EpisodeOfCare) T_DiagnosisUse(numDiagnosis int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis {
		return CodeableConceptSelect("EpisodeOfCare.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Use", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EpisodeOfCare.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Use", resource.Diagnosis[numDiagnosis].Use, optionsValueSet)
}
