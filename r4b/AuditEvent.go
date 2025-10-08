package r4b

//generated with command go run ./bultaoreune
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
func (r AuditEvent) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "AuditEvent/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "AuditEvent"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *AuditEvent) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_Subtype(numSubtype int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubtype >= len(resource.Subtype) {
		return CodingSelect("subtype["+strconv.Itoa(numSubtype)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("subtype["+strconv.Itoa(numSubtype)+"]", &resource.Subtype[numSubtype], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_Action(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAudit_event_action

	if resource == nil {
		return CodeSelect("action", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action", resource.Action, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *AuditEvent) T_Recorded(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("recorded", nil, htmlAttrs)
	}
	return StringInput("recorded", &resource.Recorded, htmlAttrs)
}
func (resource *AuditEvent) T_Outcome(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAudit_event_outcome

	if resource == nil {
		return CodeSelect("outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("outcome", resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_OutcomeDesc(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("outcomeDesc", nil, htmlAttrs)
	}
	return StringInput("outcomeDesc", resource.OutcomeDesc, htmlAttrs)
}
func (resource *AuditEvent) T_PurposeOfEvent(numPurposeOfEvent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPurposeOfEvent >= len(resource.PurposeOfEvent) {
		return CodeableConceptSelect("purposeOfEvent["+strconv.Itoa(numPurposeOfEvent)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("purposeOfEvent["+strconv.Itoa(numPurposeOfEvent)+"]", &resource.PurposeOfEvent[numPurposeOfEvent], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_AgentType(numAgent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return CodeableConceptSelect("agent["+strconv.Itoa(numAgent)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("agent["+strconv.Itoa(numAgent)+"].type", resource.Agent[numAgent].Type, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_AgentRole(numAgent int, numRole int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) || numRole >= len(resource.Agent[numAgent].Role) {
		return CodeableConceptSelect("agent["+strconv.Itoa(numAgent)+"].role["+strconv.Itoa(numRole)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("agent["+strconv.Itoa(numAgent)+"].role["+strconv.Itoa(numRole)+"]", &resource.Agent[numAgent].Role[numRole], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_AgentWho(frs []FhirResource, numAgent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return ReferenceInput(frs, "agent["+strconv.Itoa(numAgent)+"].who", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "agent["+strconv.Itoa(numAgent)+"].who", resource.Agent[numAgent].Who, htmlAttrs)
}
func (resource *AuditEvent) T_AgentAltId(numAgent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return StringInput("agent["+strconv.Itoa(numAgent)+"].altId", nil, htmlAttrs)
	}
	return StringInput("agent["+strconv.Itoa(numAgent)+"].altId", resource.Agent[numAgent].AltId, htmlAttrs)
}
func (resource *AuditEvent) T_AgentName(numAgent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return StringInput("agent["+strconv.Itoa(numAgent)+"].name", nil, htmlAttrs)
	}
	return StringInput("agent["+strconv.Itoa(numAgent)+"].name", resource.Agent[numAgent].Name, htmlAttrs)
}
func (resource *AuditEvent) T_AgentRequestor(numAgent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return BoolInput("agent["+strconv.Itoa(numAgent)+"].requestor", nil, htmlAttrs)
	}
	return BoolInput("agent["+strconv.Itoa(numAgent)+"].requestor", &resource.Agent[numAgent].Requestor, htmlAttrs)
}
func (resource *AuditEvent) T_AgentLocation(frs []FhirResource, numAgent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return ReferenceInput(frs, "agent["+strconv.Itoa(numAgent)+"].location", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "agent["+strconv.Itoa(numAgent)+"].location", resource.Agent[numAgent].Location, htmlAttrs)
}
func (resource *AuditEvent) T_AgentPolicy(numAgent int, numPolicy int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) || numPolicy >= len(resource.Agent[numAgent].Policy) {
		return StringInput("agent["+strconv.Itoa(numAgent)+"].policy["+strconv.Itoa(numPolicy)+"]", nil, htmlAttrs)
	}
	return StringInput("agent["+strconv.Itoa(numAgent)+"].policy["+strconv.Itoa(numPolicy)+"]", &resource.Agent[numAgent].Policy[numPolicy], htmlAttrs)
}
func (resource *AuditEvent) T_AgentMedia(numAgent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return CodingSelect("agent["+strconv.Itoa(numAgent)+"].media", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("agent["+strconv.Itoa(numAgent)+"].media", resource.Agent[numAgent].Media, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_AgentPurposeOfUse(numAgent int, numPurposeOfUse int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) || numPurposeOfUse >= len(resource.Agent[numAgent].PurposeOfUse) {
		return CodeableConceptSelect("agent["+strconv.Itoa(numAgent)+"].purposeOfUse["+strconv.Itoa(numPurposeOfUse)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("agent["+strconv.Itoa(numAgent)+"].purposeOfUse["+strconv.Itoa(numPurposeOfUse)+"]", &resource.Agent[numAgent].PurposeOfUse[numPurposeOfUse], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_AgentNetworkAddress(numAgent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return StringInput("agent["+strconv.Itoa(numAgent)+"].network.address", nil, htmlAttrs)
	}
	return StringInput("agent["+strconv.Itoa(numAgent)+"].network.address", resource.Agent[numAgent].Network.Address, htmlAttrs)
}
func (resource *AuditEvent) T_AgentNetworkType(numAgent int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSNetwork_type

	if resource == nil || numAgent >= len(resource.Agent) {
		return CodeSelect("agent["+strconv.Itoa(numAgent)+"].network.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("agent["+strconv.Itoa(numAgent)+"].network.type", resource.Agent[numAgent].Network.Type, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_SourceSite(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("source.site", nil, htmlAttrs)
	}
	return StringInput("source.site", resource.Source.Site, htmlAttrs)
}
func (resource *AuditEvent) T_SourceObserver(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "source.observer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "source.observer", &resource.Source.Observer, htmlAttrs)
}
func (resource *AuditEvent) T_SourceType(numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numType >= len(resource.Source.Type) {
		return CodingSelect("source.type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("source.type["+strconv.Itoa(numType)+"]", &resource.Source.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntityWhat(frs []FhirResource, numEntity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return ReferenceInput(frs, "entity["+strconv.Itoa(numEntity)+"].what", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "entity["+strconv.Itoa(numEntity)+"].what", resource.Entity[numEntity].What, htmlAttrs)
}
func (resource *AuditEvent) T_EntityType(numEntity int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return CodingSelect("entity["+strconv.Itoa(numEntity)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("entity["+strconv.Itoa(numEntity)+"].type", resource.Entity[numEntity].Type, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntityRole(numEntity int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return CodingSelect("entity["+strconv.Itoa(numEntity)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("entity["+strconv.Itoa(numEntity)+"].role", resource.Entity[numEntity].Role, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntityLifecycle(numEntity int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return CodingSelect("entity["+strconv.Itoa(numEntity)+"].lifecycle", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("entity["+strconv.Itoa(numEntity)+"].lifecycle", resource.Entity[numEntity].Lifecycle, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntitySecurityLabel(numEntity int, numSecurityLabel int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numSecurityLabel >= len(resource.Entity[numEntity].SecurityLabel) {
		return CodingSelect("entity["+strconv.Itoa(numEntity)+"].securityLabel["+strconv.Itoa(numSecurityLabel)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("entity["+strconv.Itoa(numEntity)+"].securityLabel["+strconv.Itoa(numSecurityLabel)+"]", &resource.Entity[numEntity].SecurityLabel[numSecurityLabel], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntityName(numEntity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return StringInput("entity["+strconv.Itoa(numEntity)+"].name", nil, htmlAttrs)
	}
	return StringInput("entity["+strconv.Itoa(numEntity)+"].name", resource.Entity[numEntity].Name, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDescription(numEntity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return StringInput("entity["+strconv.Itoa(numEntity)+"].description", nil, htmlAttrs)
	}
	return StringInput("entity["+strconv.Itoa(numEntity)+"].description", resource.Entity[numEntity].Description, htmlAttrs)
}
func (resource *AuditEvent) T_EntityQuery(numEntity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return StringInput("entity["+strconv.Itoa(numEntity)+"].query", nil, htmlAttrs)
	}
	return StringInput("entity["+strconv.Itoa(numEntity)+"].query", resource.Entity[numEntity].Query, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailType(numEntity int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return StringInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].type", nil, htmlAttrs)
	}
	return StringInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].type", &resource.Entity[numEntity].Detail[numDetail].Type, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueString(numEntity int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return StringInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueString", &resource.Entity[numEntity].Detail[numDetail].ValueString, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueBase64Binary(numEntity int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return StringInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueBase64Binary", &resource.Entity[numEntity].Detail[numDetail].ValueBase64Binary, htmlAttrs)
}
