package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	OccurredPeriod    *Period            `json:"occurredPeriod"`
	OccurredDateTime  *string            `json:"occurredDateTime"`
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

func (resource *Provenance) ProvenanceLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Provenance) ProvenanceReason(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reason", &resource.Reason[0], optionsValueSet)
}
func (resource *Provenance) ProvenanceActivity(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("activity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("activity", resource.Activity, optionsValueSet)
}
func (resource *Provenance) ProvenanceAgentType(numAgent int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Agent) >= numAgent {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Agent[numAgent].Type, optionsValueSet)
}
func (resource *Provenance) ProvenanceAgentRole(numAgent int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Agent) >= numAgent {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", &resource.Agent[numAgent].Role[0], optionsValueSet)
}
func (resource *Provenance) ProvenanceEntityRole(numEntity int) templ.Component {
	optionsValueSet := VSProvenance_entity_role

	if resource != nil && len(resource.Entity) >= numEntity {
		return CodeSelect("role", nil, optionsValueSet)
	}
	return CodeSelect("role", &resource.Entity[numEntity].Role, optionsValueSet)
}
