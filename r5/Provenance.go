package r5

//generated with command go run ./bultaoreune
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
	OccurredDateTime  *FhirDateTime       `json:"occurredDateTime,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
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
func (r Provenance) ResourceType() string {
	return "Provenance"
}

func (resource *Provenance) T_Target(frs []FhirResource, numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return ReferenceInput(frs, "target["+strconv.Itoa(numTarget)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "target["+strconv.Itoa(numTarget)+"]", &resource.Target[numTarget], htmlAttrs)
}
func (resource *Provenance) T_OccurredPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("occurredPeriod", nil, htmlAttrs)
	}
	return PeriodInput("occurredPeriod", resource.OccurredPeriod, htmlAttrs)
}
func (resource *Provenance) T_OccurredDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("occurredDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("occurredDateTime", resource.OccurredDateTime, htmlAttrs)
}
func (resource *Provenance) T_Recorded(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("recorded", nil, htmlAttrs)
	}
	return StringInput("recorded", resource.Recorded, htmlAttrs)
}
func (resource *Provenance) T_Policy(numPolicy int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPolicy >= len(resource.Policy) {
		return StringInput("policy["+strconv.Itoa(numPolicy)+"]", nil, htmlAttrs)
	}
	return StringInput("policy["+strconv.Itoa(numPolicy)+"]", &resource.Policy[numPolicy], htmlAttrs)
}
func (resource *Provenance) T_Location(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "location", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "location", resource.Location, htmlAttrs)
}
func (resource *Provenance) T_Authorization(numAuthorization int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthorization >= len(resource.Authorization) {
		return CodeableReferenceInput("authorization["+strconv.Itoa(numAuthorization)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("authorization["+strconv.Itoa(numAuthorization)+"]", &resource.Authorization[numAuthorization], htmlAttrs)
}
func (resource *Provenance) T_Activity(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("activity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("activity", resource.Activity, optionsValueSet, htmlAttrs)
}
func (resource *Provenance) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *Provenance) T_Patient(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "patient", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "patient", resource.Patient, htmlAttrs)
}
func (resource *Provenance) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *Provenance) T_Signature(numSignature int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSignature >= len(resource.Signature) {
		return SignatureInput("signature["+strconv.Itoa(numSignature)+"]", nil, htmlAttrs)
	}
	return SignatureInput("signature["+strconv.Itoa(numSignature)+"]", &resource.Signature[numSignature], htmlAttrs)
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
func (resource *Provenance) T_AgentWho(frs []FhirResource, numAgent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return ReferenceInput(frs, "agent["+strconv.Itoa(numAgent)+"].who", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "agent["+strconv.Itoa(numAgent)+"].who", &resource.Agent[numAgent].Who, htmlAttrs)
}
func (resource *Provenance) T_AgentOnBehalfOf(frs []FhirResource, numAgent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return ReferenceInput(frs, "agent["+strconv.Itoa(numAgent)+"].onBehalfOf", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "agent["+strconv.Itoa(numAgent)+"].onBehalfOf", resource.Agent[numAgent].OnBehalfOf, htmlAttrs)
}
func (resource *Provenance) T_EntityRole(numEntity int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSProvenance_entity_role

	if resource == nil || numEntity >= len(resource.Entity) {
		return CodeSelect("entity["+strconv.Itoa(numEntity)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("entity["+strconv.Itoa(numEntity)+"].role", &resource.Entity[numEntity].Role, optionsValueSet, htmlAttrs)
}
func (resource *Provenance) T_EntityWhat(frs []FhirResource, numEntity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return ReferenceInput(frs, "entity["+strconv.Itoa(numEntity)+"].what", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "entity["+strconv.Itoa(numEntity)+"].what", &resource.Entity[numEntity].What, htmlAttrs)
}
