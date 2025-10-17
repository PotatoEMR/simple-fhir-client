package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
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

// struct -> json, automatically add resourceType=Patient
func (r CareTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCareTeam
		ResourceType string `json:"resourceType"`
	}{
		OtherCareTeam: OtherCareTeam(r),
		ResourceType:  "CareTeam",
	})
}

// json -> struct, first reject if resourceType != CareTeam
func (r *CareTeam) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "CareTeam" {
		return errors.New("resourceType not CareTeam")
	}
	return json.Unmarshal(data, (*OtherCareTeam)(r))
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
func (resource *CareTeam) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *CareTeam) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *CareTeam) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *CareTeam) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *CareTeam) T_ReasonReference(frs []FhirResource, numReasonReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonReference >= len(resource.ReasonReference) {
		return ReferenceInput(frs, "reasonReference["+strconv.Itoa(numReasonReference)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "reasonReference["+strconv.Itoa(numReasonReference)+"]", &resource.ReasonReference[numReasonReference], htmlAttrs)
}
func (resource *CareTeam) T_ManagingOrganization(frs []FhirResource, numManagingOrganization int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numManagingOrganization >= len(resource.ManagingOrganization) {
		return ReferenceInput(frs, "managingOrganization["+strconv.Itoa(numManagingOrganization)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "managingOrganization["+strconv.Itoa(numManagingOrganization)+"]", &resource.ManagingOrganization[numManagingOrganization], htmlAttrs)
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
func (resource *CareTeam) T_ParticipantRole(numParticipant int, numRole int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) || numRole >= len(resource.Participant[numParticipant].Role) {
		return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].role["+strconv.Itoa(numRole)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].role["+strconv.Itoa(numRole)+"]", &resource.Participant[numParticipant].Role[numRole], optionsValueSet, htmlAttrs)
}
func (resource *CareTeam) T_ParticipantMember(frs []FhirResource, numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return ReferenceInput(frs, "participant["+strconv.Itoa(numParticipant)+"].member", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "participant["+strconv.Itoa(numParticipant)+"].member", resource.Participant[numParticipant].Member, htmlAttrs)
}
func (resource *CareTeam) T_ParticipantOnBehalfOf(frs []FhirResource, numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return ReferenceInput(frs, "participant["+strconv.Itoa(numParticipant)+"].onBehalfOf", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "participant["+strconv.Itoa(numParticipant)+"].onBehalfOf", resource.Participant[numParticipant].OnBehalfOf, htmlAttrs)
}
func (resource *CareTeam) T_ParticipantPeriod(numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return PeriodInput("participant["+strconv.Itoa(numParticipant)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("participant["+strconv.Itoa(numParticipant)+"].period", resource.Participant[numParticipant].Period, htmlAttrs)
}
