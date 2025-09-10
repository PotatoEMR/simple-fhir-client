package r5

//generated with command go run ./bultaoreune
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
func (resource *EpisodeOfCare) T_StatusHistoryStatus(numStatusHistory int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSEpisode_of_care_status

	if resource == nil || numStatusHistory >= len(resource.StatusHistory) {
		return CodeSelect("statusHistory["+strconv.Itoa(numStatusHistory)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("statusHistory["+strconv.Itoa(numStatusHistory)+"].status", &resource.StatusHistory[numStatusHistory].Status, optionsValueSet, htmlAttrs)
}
func (resource *EpisodeOfCare) T_ReasonUse(numReason int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableConceptSelect("reason["+strconv.Itoa(numReason)+"].use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reason["+strconv.Itoa(numReason)+"].use", resource.Reason[numReason].Use, optionsValueSet, htmlAttrs)
}
func (resource *EpisodeOfCare) T_DiagnosisUse(numDiagnosis int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].use", resource.Diagnosis[numDiagnosis].Use, optionsValueSet, htmlAttrs)
}
