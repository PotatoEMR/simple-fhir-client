package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Provenance
type Provenance struct {
	Id                *string            `json:"id,omitempty"`
	Meta              *Meta              `json:"meta,omitempty"`
	ImplicitRules     *string            `json:"implicitRules,omitempty"`
	Language          *string            `json:"language,omitempty"`
	Text              *Narrative         `json:"text,omitempty"`
	Contained         []Resource         `json:"contained,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Target            []Reference        `json:"target"`
	OccurredPeriod    *Period            `json:"occurredPeriod,omitempty"`
	OccurredDateTime  *time.Time         `json:"occurredDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Recorded          string             `json:"recorded"`
	Policy            []string           `json:"policy,omitempty"`
	Location          *Reference         `json:"location,omitempty"`
	Reason            []CodeableConcept  `json:"reason,omitempty"`
	Activity          *CodeableConcept   `json:"activity,omitempty"`
	Agent             []ProvenanceAgent  `json:"agent"`
	Entity            []ProvenanceEntity `json:"entity,omitempty"`
	Signature         []Signature        `json:"signature,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Provenance
type ProvenanceAgent struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Role              []CodeableConcept `json:"role,omitempty"`
	Who               Reference         `json:"who"`
	OnBehalfOf        *Reference        `json:"onBehalfOf,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Provenance
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
func (r Provenance) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Provenance/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "Provenance"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Provenance) T_OccurredDateTime(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Provenance.OccurredDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Provenance.OccurredDateTime", resource.OccurredDateTime, htmlAttrs)
}
func (resource *Provenance) T_Recorded(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Provenance.Recorded", nil, htmlAttrs)
	}
	return StringInput("Provenance.Recorded", &resource.Recorded, htmlAttrs)
}
func (resource *Provenance) T_Policy(numPolicy int, htmlAttrs string) templ.Component {

	if resource == nil || numPolicy >= len(resource.Policy) {
		return StringInput("Provenance.Policy."+strconv.Itoa(numPolicy)+".", nil, htmlAttrs)
	}
	return StringInput("Provenance.Policy."+strconv.Itoa(numPolicy)+".", &resource.Policy[numPolicy], htmlAttrs)
}
func (resource *Provenance) T_Reason(numReason int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableConceptSelect("Provenance.Reason."+strconv.Itoa(numReason)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Provenance.Reason."+strconv.Itoa(numReason)+".", &resource.Reason[numReason], optionsValueSet, htmlAttrs)
}
func (resource *Provenance) T_Activity(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Provenance.Activity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Provenance.Activity", resource.Activity, optionsValueSet, htmlAttrs)
}
func (resource *Provenance) T_AgentType(numAgent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numAgent >= len(resource.Agent) {
		return CodeableConceptSelect("Provenance.Agent."+strconv.Itoa(numAgent)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Provenance.Agent."+strconv.Itoa(numAgent)+"..Type", resource.Agent[numAgent].Type, optionsValueSet, htmlAttrs)
}
func (resource *Provenance) T_AgentRole(numAgent int, numRole int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numAgent >= len(resource.Agent) || numRole >= len(resource.Agent[numAgent].Role) {
		return CodeableConceptSelect("Provenance.Agent."+strconv.Itoa(numAgent)+"..Role."+strconv.Itoa(numRole)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Provenance.Agent."+strconv.Itoa(numAgent)+"..Role."+strconv.Itoa(numRole)+".", &resource.Agent[numAgent].Role[numRole], optionsValueSet, htmlAttrs)
}
func (resource *Provenance) T_EntityRole(numEntity int, htmlAttrs string) templ.Component {
	optionsValueSet := VSProvenance_entity_role

	if resource == nil || numEntity >= len(resource.Entity) {
		return CodeSelect("Provenance.Entity."+strconv.Itoa(numEntity)+"..Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Provenance.Entity."+strconv.Itoa(numEntity)+"..Role", &resource.Entity[numEntity].Role, optionsValueSet, htmlAttrs)
}
