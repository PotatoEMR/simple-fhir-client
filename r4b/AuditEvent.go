package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/AuditEvent
type AuditEvent struct {
	Id                *string            `json:"id,omitempty"`
	Meta              *Meta              `json:"meta,omitempty"`
	ImplicitRules     *string            `json:"implicitRules,omitempty"`
	Language          *string            `json:"language,omitempty"`
	Text              *Narrative         `json:"text,omitempty"`
	Contained         []Resource         `json:"contained,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Type              Coding             `json:"type"`
	Subtype           []Coding           `json:"subtype,omitempty"`
	Action            *string            `json:"action,omitempty"`
	Period            *Period            `json:"period,omitempty"`
	Recorded          string             `json:"recorded"`
	Outcome           *string            `json:"outcome,omitempty"`
	OutcomeDesc       *string            `json:"outcomeDesc,omitempty"`
	PurposeOfEvent    []CodeableConcept  `json:"purposeOfEvent,omitempty"`
	Agent             []AuditEventAgent  `json:"agent"`
	Source            AuditEventSource   `json:"source"`
	Entity            []AuditEventEntity `json:"entity,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/AuditEvent
type AuditEventAgent struct {
	Id                *string                 `json:"id,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept        `json:"type,omitempty"`
	Role              []CodeableConcept       `json:"role,omitempty"`
	Who               *Reference              `json:"who,omitempty"`
	AltId             *string                 `json:"altId,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Requestor         bool                    `json:"requestor"`
	Location          *Reference              `json:"location,omitempty"`
	Policy            []string                `json:"policy,omitempty"`
	Media             *Coding                 `json:"media,omitempty"`
	Network           *AuditEventAgentNetwork `json:"network,omitempty"`
	PurposeOfUse      []CodeableConcept       `json:"purposeOfUse,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/AuditEvent
type AuditEventAgentNetwork struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Address           *string     `json:"address,omitempty"`
	Type              *string     `json:"type,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/AuditEvent
type AuditEventSource struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Site              *string     `json:"site,omitempty"`
	Observer          Reference   `json:"observer"`
	Type              []Coding    `json:"type,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/AuditEvent
type AuditEventEntity struct {
	Id                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	What              *Reference               `json:"what,omitempty"`
	Type              *Coding                  `json:"type,omitempty"`
	Role              *Coding                  `json:"role,omitempty"`
	Lifecycle         *Coding                  `json:"lifecycle,omitempty"`
	SecurityLabel     []Coding                 `json:"securityLabel,omitempty"`
	Name              *string                  `json:"name,omitempty"`
	Description       *string                  `json:"description,omitempty"`
	Query             *string                  `json:"query,omitempty"`
	Detail            []AuditEventEntityDetail `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/AuditEvent
type AuditEventEntityDetail struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	ValueString       string      `json:"valueString"`
	ValueBase64Binary string      `json:"valueBase64Binary"`
}

type OtherAuditEvent AuditEvent

// on convert struct to json, automatically add resourceType=AuditEvent
func (r AuditEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAuditEvent
		ResourceType string `json:"resourceType"`
	}{
		OtherAuditEvent: OtherAuditEvent(r),
		ResourceType:    "AuditEvent",
	})
}

func (resource *AuditEvent) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *AuditEvent) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("type", nil, optionsValueSet)
	}
	return CodingSelect("type", &resource.Type, optionsValueSet)
}
func (resource *AuditEvent) T_Subtype(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("subtype", nil, optionsValueSet)
	}
	return CodingSelect("subtype", &resource.Subtype[0], optionsValueSet)
}
func (resource *AuditEvent) T_Action() templ.Component {
	optionsValueSet := VSAudit_event_action

	if resource == nil {
		return CodeSelect("action", nil, optionsValueSet)
	}
	return CodeSelect("action", resource.Action, optionsValueSet)
}
func (resource *AuditEvent) T_Outcome() templ.Component {
	optionsValueSet := VSAudit_event_outcome

	if resource == nil {
		return CodeSelect("outcome", nil, optionsValueSet)
	}
	return CodeSelect("outcome", resource.Outcome, optionsValueSet)
}
func (resource *AuditEvent) T_PurposeOfEvent(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("purposeOfEvent", nil, optionsValueSet)
	}
	return CodeableConceptSelect("purposeOfEvent", &resource.PurposeOfEvent[0], optionsValueSet)
}
func (resource *AuditEvent) T_AgentType(numAgent int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Agent) >= numAgent {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Agent[numAgent].Type, optionsValueSet)
}
func (resource *AuditEvent) T_AgentRole(numAgent int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Agent) >= numAgent {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", &resource.Agent[numAgent].Role[0], optionsValueSet)
}
func (resource *AuditEvent) T_AgentMedia(numAgent int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Agent) >= numAgent {
		return CodingSelect("media", nil, optionsValueSet)
	}
	return CodingSelect("media", resource.Agent[numAgent].Media, optionsValueSet)
}
func (resource *AuditEvent) T_AgentPurposeOfUse(numAgent int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Agent) >= numAgent {
		return CodeableConceptSelect("purposeOfUse", nil, optionsValueSet)
	}
	return CodeableConceptSelect("purposeOfUse", &resource.Agent[numAgent].PurposeOfUse[0], optionsValueSet)
}
func (resource *AuditEvent) T_AgentNetworkType(numAgent int) templ.Component {
	optionsValueSet := VSNetwork_type

	if resource == nil && len(resource.Agent) >= numAgent {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", resource.Agent[numAgent].Network.Type, optionsValueSet)
}
func (resource *AuditEvent) T_SourceType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("type", nil, optionsValueSet)
	}
	return CodingSelect("type", &resource.Source.Type[0], optionsValueSet)
}
func (resource *AuditEvent) T_EntityType(numEntity int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Entity) >= numEntity {
		return CodingSelect("type", nil, optionsValueSet)
	}
	return CodingSelect("type", resource.Entity[numEntity].Type, optionsValueSet)
}
func (resource *AuditEvent) T_EntityRole(numEntity int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Entity) >= numEntity {
		return CodingSelect("role", nil, optionsValueSet)
	}
	return CodingSelect("role", resource.Entity[numEntity].Role, optionsValueSet)
}
func (resource *AuditEvent) T_EntityLifecycle(numEntity int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Entity) >= numEntity {
		return CodingSelect("lifecycle", nil, optionsValueSet)
	}
	return CodingSelect("lifecycle", resource.Entity[numEntity].Lifecycle, optionsValueSet)
}
func (resource *AuditEvent) T_EntitySecurityLabel(numEntity int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Entity) >= numEntity {
		return CodingSelect("securityLabel", nil, optionsValueSet)
	}
	return CodingSelect("securityLabel", &resource.Entity[numEntity].SecurityLabel[0], optionsValueSet)
}
