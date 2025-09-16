package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Provenance
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
	OccurredDateTime  *FhirDateTime      `json:"occurredDateTime,omitempty"`
	Recorded          string             `json:"recorded"`
	Policy            []string           `json:"policy,omitempty"`
	Location          *Reference         `json:"location,omitempty"`
	Reason            []CodeableConcept  `json:"reason,omitempty"`
	Activity          *CodeableConcept   `json:"activity,omitempty"`
	Agent             []ProvenanceAgent  `json:"agent"`
	Entity            []ProvenanceEntity `json:"entity,omitempty"`
	Signature         []Signature        `json:"signature,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Provenance
type ProvenanceAgent struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Role              []CodeableConcept `json:"role,omitempty"`
	Who               Reference         `json:"who"`
	OnBehalfOf        *Reference        `json:"onBehalfOf,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Provenance
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
func (resource *Provenance) T_OccurredDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("occurredDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("occurredDateTime", resource.OccurredDateTime, htmlAttrs)
}
func (resource *Provenance) T_Recorded(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("recorded", nil, htmlAttrs)
	}
	return StringInput("recorded", &resource.Recorded, htmlAttrs)
}
func (resource *Provenance) T_Policy(numPolicy int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPolicy >= len(resource.Policy) {
		return StringInput("policy["+strconv.Itoa(numPolicy)+"]", nil, htmlAttrs)
	}
	return StringInput("policy["+strconv.Itoa(numPolicy)+"]", &resource.Policy[numPolicy], htmlAttrs)
}
func (resource *Provenance) T_Reason(numReason int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableConceptSelect("reason["+strconv.Itoa(numReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], optionsValueSet, htmlAttrs)
}
func (resource *Provenance) T_Activity(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("activity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("activity", resource.Activity, optionsValueSet, htmlAttrs)
}
func (resource *Provenance) T_AgentType(numAgent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return CodeableConceptSelect("agent["+strconv.Itoa(numAgent)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("agent["+strconv.Itoa(numAgent)+"].type", resource.Agent[numAgent].Type, optionsValueSet, htmlAttrs)
}
func (resource *Provenance) T_AgentRole(numAgent int, numRole int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) || numRole >= len(resource.Agent[numAgent].Role) {
		return CodeableConceptSelect("agent["+strconv.Itoa(numAgent)+"].role["+strconv.Itoa(numRole)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("agent["+strconv.Itoa(numAgent)+"].role["+strconv.Itoa(numRole)+"]", &resource.Agent[numAgent].Role[numRole], optionsValueSet, htmlAttrs)
}
func (resource *Provenance) T_EntityRole(numEntity int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSProvenance_entity_role

	if resource == nil || numEntity >= len(resource.Entity) {
		return CodeSelect("entity["+strconv.Itoa(numEntity)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("entity["+strconv.Itoa(numEntity)+"].role", &resource.Entity[numEntity].Role, optionsValueSet, htmlAttrs)
}
