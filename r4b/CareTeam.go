package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/CareTeam
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
	Encounter            *Reference            `json:"encounter,omitempty"`
	Period               *Period               `json:"period,omitempty"`
	Participant          []CareTeamParticipant `json:"participant,omitempty"`
	ReasonCode           []CodeableConcept     `json:"reasonCode,omitempty"`
	ReasonReference      []Reference           `json:"reasonReference,omitempty"`
	ManagingOrganization []Reference           `json:"managingOrganization,omitempty"`
	Telecom              []ContactPoint        `json:"telecom,omitempty"`
	Note                 []Annotation          `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/CareTeam
type CareTeamParticipant struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Role              []CodeableConcept `json:"role,omitempty"`
	Member            *Reference        `json:"member,omitempty"`
	OnBehalfOf        *Reference        `json:"onBehalfOf,omitempty"`
	Period            *Period           `json:"period,omitempty"`
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

func (resource *CareTeam) T_Id() templ.Component {

	if resource == nil {
		return StringInput("CareTeam.Id", nil)
	}
	return StringInput("CareTeam.Id", resource.Id)
}
func (resource *CareTeam) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("CareTeam.ImplicitRules", nil)
	}
	return StringInput("CareTeam.ImplicitRules", resource.ImplicitRules)
}
func (resource *CareTeam) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("CareTeam.Language", nil, optionsValueSet)
	}
	return CodeSelect("CareTeam.Language", resource.Language, optionsValueSet)
}
func (resource *CareTeam) T_Status() templ.Component {
	optionsValueSet := VSCare_team_status

	if resource == nil {
		return CodeSelect("CareTeam.Status", nil, optionsValueSet)
	}
	return CodeSelect("CareTeam.Status", resource.Status, optionsValueSet)
}
func (resource *CareTeam) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("CareTeam.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CareTeam.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *CareTeam) T_Name() templ.Component {

	if resource == nil {
		return StringInput("CareTeam.Name", nil)
	}
	return StringInput("CareTeam.Name", resource.Name)
}
func (resource *CareTeam) T_ReasonCode(numReasonCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ReasonCode) >= numReasonCode {
		return CodeableConceptSelect("CareTeam.ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CareTeam.ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet)
}
func (resource *CareTeam) T_ParticipantId(numParticipant int) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return StringInput("CareTeam.Participant["+strconv.Itoa(numParticipant)+"].Id", nil)
	}
	return StringInput("CareTeam.Participant["+strconv.Itoa(numParticipant)+"].Id", resource.Participant[numParticipant].Id)
}
func (resource *CareTeam) T_ParticipantRole(numParticipant int, numRole int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant || len(resource.Participant[numParticipant].Role) >= numRole {
		return CodeableConceptSelect("CareTeam.Participant["+strconv.Itoa(numParticipant)+"].Role["+strconv.Itoa(numRole)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CareTeam.Participant["+strconv.Itoa(numParticipant)+"].Role["+strconv.Itoa(numRole)+"]", &resource.Participant[numParticipant].Role[numRole], optionsValueSet)
}
