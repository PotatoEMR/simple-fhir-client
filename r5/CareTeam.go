package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	CoveragePeriod    *Period          `json:"coveragePeriod,omitempty"`
	CoverageTiming    *Timing          `json:"coverageTiming,omitempty"`
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
func (r CareTeam) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "CareTeam/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "CareTeam"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *CareTeam) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCare_team_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CareTeam) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *CareTeam) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *CareTeam) T_Subject(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("subject", nil, htmlAttrs)
	}
	return ReferenceInput("subject", resource.Subject, htmlAttrs)
}
func (resource *CareTeam) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *CareTeam) T_Reason(numReason int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], htmlAttrs)
}
func (resource *CareTeam) T_ManagingOrganization(numManagingOrganization int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numManagingOrganization >= len(resource.ManagingOrganization) {
		return ReferenceInput("managingOrganization["+strconv.Itoa(numManagingOrganization)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("managingOrganization["+strconv.Itoa(numManagingOrganization)+"]", &resource.ManagingOrganization[numManagingOrganization], htmlAttrs)
}
func (resource *CareTeam) T_Telecom(numTelecom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTelecom >= len(resource.Telecom) {
		return ContactPointInput("telecom["+strconv.Itoa(numTelecom)+"]", nil, htmlAttrs)
	}
	return ContactPointInput("telecom["+strconv.Itoa(numTelecom)+"]", &resource.Telecom[numTelecom], htmlAttrs)
}
func (resource *CareTeam) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *CareTeam) T_ParticipantRole(numParticipant int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].role", resource.Participant[numParticipant].Role, optionsValueSet, htmlAttrs)
}
func (resource *CareTeam) T_ParticipantMember(numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return ReferenceInput("participant["+strconv.Itoa(numParticipant)+"].member", nil, htmlAttrs)
	}
	return ReferenceInput("participant["+strconv.Itoa(numParticipant)+"].member", resource.Participant[numParticipant].Member, htmlAttrs)
}
func (resource *CareTeam) T_ParticipantOnBehalfOf(numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return ReferenceInput("participant["+strconv.Itoa(numParticipant)+"].onBehalfOf", nil, htmlAttrs)
	}
	return ReferenceInput("participant["+strconv.Itoa(numParticipant)+"].onBehalfOf", resource.Participant[numParticipant].OnBehalfOf, htmlAttrs)
}
func (resource *CareTeam) T_ParticipantCoveragePeriod(numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return PeriodInput("participant["+strconv.Itoa(numParticipant)+"].coveragePeriod", nil, htmlAttrs)
	}
	return PeriodInput("participant["+strconv.Itoa(numParticipant)+"].coveragePeriod", resource.Participant[numParticipant].CoveragePeriod, htmlAttrs)
}
func (resource *CareTeam) T_ParticipantCoverageTiming(numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return TimingInput("participant["+strconv.Itoa(numParticipant)+"].coverageTiming", nil, htmlAttrs)
	}
	return TimingInput("participant["+strconv.Itoa(numParticipant)+"].coverageTiming", resource.Participant[numParticipant].CoverageTiming, htmlAttrs)
}
