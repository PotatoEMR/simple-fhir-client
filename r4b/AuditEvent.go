package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *AuditEvent) T_Id() templ.Component {

	if resource == nil {
		return StringInput("AuditEvent.Id", nil)
	}
	return StringInput("AuditEvent.Id", resource.Id)
}
func (resource *AuditEvent) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("AuditEvent.ImplicitRules", nil)
	}
	return StringInput("AuditEvent.ImplicitRules", resource.ImplicitRules)
}
func (resource *AuditEvent) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("AuditEvent.Language", nil, optionsValueSet)
	}
	return CodeSelect("AuditEvent.Language", resource.Language, optionsValueSet)
}
func (resource *AuditEvent) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("AuditEvent.Type", nil, optionsValueSet)
	}
	return CodingSelect("AuditEvent.Type", &resource.Type, optionsValueSet)
}
func (resource *AuditEvent) T_Subtype(numSubtype int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Subtype) >= numSubtype {
		return CodingSelect("AuditEvent.Subtype["+strconv.Itoa(numSubtype)+"]", nil, optionsValueSet)
	}
	return CodingSelect("AuditEvent.Subtype["+strconv.Itoa(numSubtype)+"]", &resource.Subtype[numSubtype], optionsValueSet)
}
func (resource *AuditEvent) T_Action() templ.Component {
	optionsValueSet := VSAudit_event_action

	if resource == nil {
		return CodeSelect("AuditEvent.Action", nil, optionsValueSet)
	}
	return CodeSelect("AuditEvent.Action", resource.Action, optionsValueSet)
}
func (resource *AuditEvent) T_Recorded() templ.Component {

	if resource == nil {
		return StringInput("AuditEvent.Recorded", nil)
	}
	return StringInput("AuditEvent.Recorded", &resource.Recorded)
}
func (resource *AuditEvent) T_Outcome() templ.Component {
	optionsValueSet := VSAudit_event_outcome

	if resource == nil {
		return CodeSelect("AuditEvent.Outcome", nil, optionsValueSet)
	}
	return CodeSelect("AuditEvent.Outcome", resource.Outcome, optionsValueSet)
}
func (resource *AuditEvent) T_OutcomeDesc() templ.Component {

	if resource == nil {
		return StringInput("AuditEvent.OutcomeDesc", nil)
	}
	return StringInput("AuditEvent.OutcomeDesc", resource.OutcomeDesc)
}
func (resource *AuditEvent) T_PurposeOfEvent(numPurposeOfEvent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PurposeOfEvent) >= numPurposeOfEvent {
		return CodeableConceptSelect("AuditEvent.PurposeOfEvent["+strconv.Itoa(numPurposeOfEvent)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AuditEvent.PurposeOfEvent["+strconv.Itoa(numPurposeOfEvent)+"]", &resource.PurposeOfEvent[numPurposeOfEvent], optionsValueSet)
}
func (resource *AuditEvent) T_AgentId(numAgent int) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent {
		return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Id", nil)
	}
	return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Id", resource.Agent[numAgent].Id)
}
func (resource *AuditEvent) T_AgentType(numAgent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent {
		return CodeableConceptSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Type", resource.Agent[numAgent].Type, optionsValueSet)
}
func (resource *AuditEvent) T_AgentRole(numAgent int, numRole int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent || len(resource.Agent[numAgent].Role) >= numRole {
		return CodeableConceptSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Role["+strconv.Itoa(numRole)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Role["+strconv.Itoa(numRole)+"]", &resource.Agent[numAgent].Role[numRole], optionsValueSet)
}
func (resource *AuditEvent) T_AgentAltId(numAgent int) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent {
		return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].AltId", nil)
	}
	return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].AltId", resource.Agent[numAgent].AltId)
}
func (resource *AuditEvent) T_AgentName(numAgent int) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent {
		return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Name", nil)
	}
	return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Name", resource.Agent[numAgent].Name)
}
func (resource *AuditEvent) T_AgentRequestor(numAgent int) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent {
		return BoolInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Requestor", nil)
	}
	return BoolInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Requestor", &resource.Agent[numAgent].Requestor)
}
func (resource *AuditEvent) T_AgentPolicy(numAgent int, numPolicy int) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent || len(resource.Agent[numAgent].Policy) >= numPolicy {
		return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Policy["+strconv.Itoa(numPolicy)+"]", nil)
	}
	return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Policy["+strconv.Itoa(numPolicy)+"]", &resource.Agent[numAgent].Policy[numPolicy])
}
func (resource *AuditEvent) T_AgentMedia(numAgent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent {
		return CodingSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Media", nil, optionsValueSet)
	}
	return CodingSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Media", resource.Agent[numAgent].Media, optionsValueSet)
}
func (resource *AuditEvent) T_AgentPurposeOfUse(numAgent int, numPurposeOfUse int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent || len(resource.Agent[numAgent].PurposeOfUse) >= numPurposeOfUse {
		return CodeableConceptSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].PurposeOfUse["+strconv.Itoa(numPurposeOfUse)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].PurposeOfUse["+strconv.Itoa(numPurposeOfUse)+"]", &resource.Agent[numAgent].PurposeOfUse[numPurposeOfUse], optionsValueSet)
}
func (resource *AuditEvent) T_AgentNetworkId(numAgent int) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent {
		return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Network.Id", nil)
	}
	return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Network.Id", resource.Agent[numAgent].Network.Id)
}
func (resource *AuditEvent) T_AgentNetworkAddress(numAgent int) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent {
		return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Network.Address", nil)
	}
	return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Network.Address", resource.Agent[numAgent].Network.Address)
}
func (resource *AuditEvent) T_AgentNetworkType(numAgent int) templ.Component {
	optionsValueSet := VSNetwork_type

	if resource == nil || len(resource.Agent) >= numAgent {
		return CodeSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Network.Type", nil, optionsValueSet)
	}
	return CodeSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Network.Type", resource.Agent[numAgent].Network.Type, optionsValueSet)
}
func (resource *AuditEvent) T_SourceId() templ.Component {

	if resource == nil {
		return StringInput("AuditEvent.Source.Id", nil)
	}
	return StringInput("AuditEvent.Source.Id", resource.Source.Id)
}
func (resource *AuditEvent) T_SourceSite() templ.Component {

	if resource == nil {
		return StringInput("AuditEvent.Source.Site", nil)
	}
	return StringInput("AuditEvent.Source.Site", resource.Source.Site)
}
func (resource *AuditEvent) T_SourceType(numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Source.Type) >= numType {
		return CodingSelect("AuditEvent.Source.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodingSelect("AuditEvent.Source.Type["+strconv.Itoa(numType)+"]", &resource.Source.Type[numType], optionsValueSet)
}
func (resource *AuditEvent) T_EntityId(numEntity int) templ.Component {

	if resource == nil || len(resource.Entity) >= numEntity {
		return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Id", nil)
	}
	return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Id", resource.Entity[numEntity].Id)
}
func (resource *AuditEvent) T_EntityType(numEntity int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Entity) >= numEntity {
		return CodingSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Type", nil, optionsValueSet)
	}
	return CodingSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Type", resource.Entity[numEntity].Type, optionsValueSet)
}
func (resource *AuditEvent) T_EntityRole(numEntity int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Entity) >= numEntity {
		return CodingSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Role", nil, optionsValueSet)
	}
	return CodingSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Role", resource.Entity[numEntity].Role, optionsValueSet)
}
func (resource *AuditEvent) T_EntityLifecycle(numEntity int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Entity) >= numEntity {
		return CodingSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Lifecycle", nil, optionsValueSet)
	}
	return CodingSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Lifecycle", resource.Entity[numEntity].Lifecycle, optionsValueSet)
}
func (resource *AuditEvent) T_EntitySecurityLabel(numEntity int, numSecurityLabel int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Entity) >= numEntity || len(resource.Entity[numEntity].SecurityLabel) >= numSecurityLabel {
		return CodingSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", nil, optionsValueSet)
	}
	return CodingSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", &resource.Entity[numEntity].SecurityLabel[numSecurityLabel], optionsValueSet)
}
func (resource *AuditEvent) T_EntityName(numEntity int) templ.Component {

	if resource == nil || len(resource.Entity) >= numEntity {
		return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Name", nil)
	}
	return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Name", resource.Entity[numEntity].Name)
}
func (resource *AuditEvent) T_EntityDescription(numEntity int) templ.Component {

	if resource == nil || len(resource.Entity) >= numEntity {
		return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Description", nil)
	}
	return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Description", resource.Entity[numEntity].Description)
}
func (resource *AuditEvent) T_EntityQuery(numEntity int) templ.Component {

	if resource == nil || len(resource.Entity) >= numEntity {
		return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Query", nil)
	}
	return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Query", resource.Entity[numEntity].Query)
}
func (resource *AuditEvent) T_EntityDetailId(numEntity int, numDetail int) templ.Component {

	if resource == nil || len(resource.Entity) >= numEntity || len(resource.Entity[numEntity].Detail) >= numDetail {
		return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Detail["+strconv.Itoa(numDetail)+"].Id", nil)
	}
	return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Detail["+strconv.Itoa(numDetail)+"].Id", resource.Entity[numEntity].Detail[numDetail].Id)
}
func (resource *AuditEvent) T_EntityDetailType(numEntity int, numDetail int) templ.Component {

	if resource == nil || len(resource.Entity) >= numEntity || len(resource.Entity[numEntity].Detail) >= numDetail {
		return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Detail["+strconv.Itoa(numDetail)+"].Type", nil)
	}
	return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Detail["+strconv.Itoa(numDetail)+"].Type", &resource.Entity[numEntity].Detail[numDetail].Type)
}
