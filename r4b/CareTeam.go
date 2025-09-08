package r4b

//generated with command go run ./bultaoreune
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
func (resource *CareTeam) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSCare_team_status

	if resource == nil {
		return CodeSelect("CareTeam.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CareTeam.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CareTeam) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("CareTeam.Category."+strconv.Itoa(numCategory)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CareTeam.Category."+strconv.Itoa(numCategory)+".", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *CareTeam) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CareTeam.Name", nil, htmlAttrs)
	}
	return StringInput("CareTeam.Name", resource.Name, htmlAttrs)
}
func (resource *CareTeam) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("CareTeam.ReasonCode."+strconv.Itoa(numReasonCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CareTeam.ReasonCode."+strconv.Itoa(numReasonCode)+".", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *CareTeam) T_Note(numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("CareTeam.Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("CareTeam.Note."+strconv.Itoa(numNote)+".", &resource.Note[numNote], htmlAttrs)
}
func (resource *CareTeam) T_ParticipantRole(numParticipant int, numRole int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numParticipant >= len(resource.Participant) || numRole >= len(resource.Participant[numParticipant].Role) {
		return CodeableConceptSelect("CareTeam.Participant."+strconv.Itoa(numParticipant)+"..Role."+strconv.Itoa(numRole)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CareTeam.Participant."+strconv.Itoa(numParticipant)+"..Role."+strconv.Itoa(numRole)+".", &resource.Participant[numParticipant].Role[numRole], optionsValueSet, htmlAttrs)
}
