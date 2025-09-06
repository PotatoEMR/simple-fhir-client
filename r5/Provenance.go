package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Provenance
type Provenance struct {
	Id                *string             `json:"id,omitempty"`
	Meta              *Meta               `json:"meta,omitempty"`
	ImplicitRules     *string             `json:"implicitRules,omitempty"`
	Language          *string             `json:"language,omitempty"`
	Text              *Narrative          `json:"text,omitempty"`
	Contained         []Resource          `json:"contained,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Target            []Reference         `json:"target"`
	OccurredPeriod    *Period             `json:"occurredPeriod,omitempty"`
	OccurredDateTime  *string             `json:"occurredDateTime,omitempty"`
	Recorded          *string             `json:"recorded,omitempty"`
	Policy            []string            `json:"policy,omitempty"`
	Location          *Reference          `json:"location,omitempty"`
	Authorization     []CodeableReference `json:"authorization,omitempty"`
	Activity          *CodeableConcept    `json:"activity,omitempty"`
	BasedOn           []Reference         `json:"basedOn,omitempty"`
	Patient           *Reference          `json:"patient,omitempty"`
	Encounter         *Reference          `json:"encounter,omitempty"`
	Agent             []ProvenanceAgent   `json:"agent"`
	Entity            []ProvenanceEntity  `json:"entity,omitempty"`
	Signature         []Signature         `json:"signature,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Provenance
type ProvenanceAgent struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Role              []CodeableConcept `json:"role,omitempty"`
	Who               Reference         `json:"who"`
	OnBehalfOf        *Reference        `json:"onBehalfOf,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Provenance
type ProvenanceEntity struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Role              string      `json:"role"`
	What              Reference   `json:"what"`
}

type OtherProvenance Provenance

// on convert struct to json, automatically add resourceType=Provenance
func (r Provenance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProvenance
		ResourceType string `json:"resourceType"`
	}{
		OtherProvenance: OtherProvenance(r),
		ResourceType:    "Provenance",
	})
}

func (resource *Provenance) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Provenance.Id", nil)
	}
	return StringInput("Provenance.Id", resource.Id)
}
func (resource *Provenance) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Provenance.ImplicitRules", nil)
	}
	return StringInput("Provenance.ImplicitRules", resource.ImplicitRules)
}
func (resource *Provenance) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Provenance.Language", nil, optionsValueSet)
	}
	return CodeSelect("Provenance.Language", resource.Language, optionsValueSet)
}
func (resource *Provenance) T_Recorded() templ.Component {

	if resource == nil {
		return StringInput("Provenance.Recorded", nil)
	}
	return StringInput("Provenance.Recorded", resource.Recorded)
}
func (resource *Provenance) T_Policy(numPolicy int) templ.Component {

	if resource == nil || len(resource.Policy) >= numPolicy {
		return StringInput("Provenance.Policy["+strconv.Itoa(numPolicy)+"]", nil)
	}
	return StringInput("Provenance.Policy["+strconv.Itoa(numPolicy)+"]", &resource.Policy[numPolicy])
}
func (resource *Provenance) T_Activity(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Provenance.Activity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Provenance.Activity", resource.Activity, optionsValueSet)
}
func (resource *Provenance) T_AgentId(numAgent int) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent {
		return StringInput("Provenance.Agent["+strconv.Itoa(numAgent)+"].Id", nil)
	}
	return StringInput("Provenance.Agent["+strconv.Itoa(numAgent)+"].Id", resource.Agent[numAgent].Id)
}
func (resource *Provenance) T_AgentType(numAgent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent {
		return CodeableConceptSelect("Provenance.Agent["+strconv.Itoa(numAgent)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Provenance.Agent["+strconv.Itoa(numAgent)+"].Type", resource.Agent[numAgent].Type, optionsValueSet)
}
func (resource *Provenance) T_AgentRole(numAgent int, numRole int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent || len(resource.Agent[numAgent].Role) >= numRole {
		return CodeableConceptSelect("Provenance.Agent["+strconv.Itoa(numAgent)+"].Role["+strconv.Itoa(numRole)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Provenance.Agent["+strconv.Itoa(numAgent)+"].Role["+strconv.Itoa(numRole)+"]", &resource.Agent[numAgent].Role[numRole], optionsValueSet)
}
func (resource *Provenance) T_EntityId(numEntity int) templ.Component {

	if resource == nil || len(resource.Entity) >= numEntity {
		return StringInput("Provenance.Entity["+strconv.Itoa(numEntity)+"].Id", nil)
	}
	return StringInput("Provenance.Entity["+strconv.Itoa(numEntity)+"].Id", resource.Entity[numEntity].Id)
}
func (resource *Provenance) T_EntityRole(numEntity int) templ.Component {
	optionsValueSet := VSProvenance_entity_role

	if resource == nil || len(resource.Entity) >= numEntity {
		return CodeSelect("Provenance.Entity["+strconv.Itoa(numEntity)+"].Role", nil, optionsValueSet)
	}
	return CodeSelect("Provenance.Entity["+strconv.Itoa(numEntity)+"].Role", &resource.Entity[numEntity].Role, optionsValueSet)
}
