package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/AuditEvent
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

// http://hl7.org/fhir/r4/StructureDefinition/AuditEvent
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

// http://hl7.org/fhir/r4/StructureDefinition/AuditEvent
type AuditEventAgentNetwork struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Address           *string     `json:"address,omitempty"`
	Type              *string     `json:"type,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/AuditEvent
type AuditEventSource struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Site              *string     `json:"site,omitempty"`
	Observer          Reference   `json:"observer"`
	Type              []Coding    `json:"type,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/AuditEvent
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

// http://hl7.org/fhir/r4/StructureDefinition/AuditEvent
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
func (resource *AuditEvent) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("AuditEvent.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("AuditEvent.Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_Subtype(numSubtype int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSubtype >= len(resource.Subtype) {
		return CodingSelect("AuditEvent.Subtype["+strconv.Itoa(numSubtype)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("AuditEvent.Subtype["+strconv.Itoa(numSubtype)+"]", &resource.Subtype[numSubtype], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_Action(htmlAttrs string) templ.Component {
	optionsValueSet := VSAudit_event_action

	if resource == nil {
		return CodeSelect("AuditEvent.Action", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("AuditEvent.Action", resource.Action, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_Recorded(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("AuditEvent.Recorded", nil, htmlAttrs)
	}
	return StringInput("AuditEvent.Recorded", &resource.Recorded, htmlAttrs)
}
func (resource *AuditEvent) T_Outcome(htmlAttrs string) templ.Component {
	optionsValueSet := VSAudit_event_outcome

	if resource == nil {
		return CodeSelect("AuditEvent.Outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("AuditEvent.Outcome", resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_OutcomeDesc(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("AuditEvent.OutcomeDesc", nil, htmlAttrs)
	}
	return StringInput("AuditEvent.OutcomeDesc", resource.OutcomeDesc, htmlAttrs)
}
func (resource *AuditEvent) T_PurposeOfEvent(numPurposeOfEvent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPurposeOfEvent >= len(resource.PurposeOfEvent) {
		return CodeableConceptSelect("AuditEvent.PurposeOfEvent["+strconv.Itoa(numPurposeOfEvent)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AuditEvent.PurposeOfEvent["+strconv.Itoa(numPurposeOfEvent)+"]", &resource.PurposeOfEvent[numPurposeOfEvent], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_AgentType(numAgent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return CodeableConceptSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Type", resource.Agent[numAgent].Type, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_AgentRole(numAgent int, numRole int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) || numRole >= len(resource.Agent[numAgent].Role) {
		return CodeableConceptSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Role["+strconv.Itoa(numRole)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Role["+strconv.Itoa(numRole)+"]", &resource.Agent[numAgent].Role[numRole], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_AgentAltId(numAgent int, htmlAttrs string) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].AltId", nil, htmlAttrs)
	}
	return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].AltId", resource.Agent[numAgent].AltId, htmlAttrs)
}
func (resource *AuditEvent) T_AgentName(numAgent int, htmlAttrs string) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Name", nil, htmlAttrs)
	}
	return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Name", resource.Agent[numAgent].Name, htmlAttrs)
}
func (resource *AuditEvent) T_AgentRequestor(numAgent int, htmlAttrs string) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return BoolInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Requestor", nil, htmlAttrs)
	}
	return BoolInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Requestor", &resource.Agent[numAgent].Requestor, htmlAttrs)
}
func (resource *AuditEvent) T_AgentPolicy(numAgent int, numPolicy int, htmlAttrs string) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) || numPolicy >= len(resource.Agent[numAgent].Policy) {
		return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Policy["+strconv.Itoa(numPolicy)+"]", nil, htmlAttrs)
	}
	return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Policy["+strconv.Itoa(numPolicy)+"]", &resource.Agent[numAgent].Policy[numPolicy], htmlAttrs)
}
func (resource *AuditEvent) T_AgentMedia(numAgent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return CodingSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Media", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Media", resource.Agent[numAgent].Media, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_AgentPurposeOfUse(numAgent int, numPurposeOfUse int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) || numPurposeOfUse >= len(resource.Agent[numAgent].PurposeOfUse) {
		return CodeableConceptSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].PurposeOfUse["+strconv.Itoa(numPurposeOfUse)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].PurposeOfUse["+strconv.Itoa(numPurposeOfUse)+"]", &resource.Agent[numAgent].PurposeOfUse[numPurposeOfUse], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_AgentNetworkAddress(numAgent int, htmlAttrs string) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Network.Address", nil, htmlAttrs)
	}
	return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Network.Address", resource.Agent[numAgent].Network.Address, htmlAttrs)
}
func (resource *AuditEvent) T_AgentNetworkType(numAgent int, htmlAttrs string) templ.Component {
	optionsValueSet := VSNetwork_type

	if resource == nil || numAgent >= len(resource.Agent) {
		return CodeSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Network.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Network.Type", resource.Agent[numAgent].Network.Type, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_SourceSite(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("AuditEvent.Source.Site", nil, htmlAttrs)
	}
	return StringInput("AuditEvent.Source.Site", resource.Source.Site, htmlAttrs)
}
func (resource *AuditEvent) T_SourceType(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numType >= len(resource.Source.Type) {
		return CodingSelect("AuditEvent.Source.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("AuditEvent.Source.Type["+strconv.Itoa(numType)+"]", &resource.Source.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntityType(numEntity int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return CodingSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Type", resource.Entity[numEntity].Type, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntityRole(numEntity int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return CodingSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Role", resource.Entity[numEntity].Role, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntityLifecycle(numEntity int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return CodingSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Lifecycle", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Lifecycle", resource.Entity[numEntity].Lifecycle, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntitySecurityLabel(numEntity int, numSecurityLabel int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numSecurityLabel >= len(resource.Entity[numEntity].SecurityLabel) {
		return CodingSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", &resource.Entity[numEntity].SecurityLabel[numSecurityLabel], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntityName(numEntity int, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Name", nil, htmlAttrs)
	}
	return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Name", resource.Entity[numEntity].Name, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDescription(numEntity int, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Description", nil, htmlAttrs)
	}
	return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Description", resource.Entity[numEntity].Description, htmlAttrs)
}
func (resource *AuditEvent) T_EntityQuery(numEntity int, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Query", nil, htmlAttrs)
	}
	return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Query", resource.Entity[numEntity].Query, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailType(numEntity int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Detail["+strconv.Itoa(numDetail)+"].Type", nil, htmlAttrs)
	}
	return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Detail["+strconv.Itoa(numDetail)+"].Type", &resource.Entity[numEntity].Detail[numDetail].Type, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueString(numEntity int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Detail["+strconv.Itoa(numDetail)+"].ValueString", nil, htmlAttrs)
	}
	return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Detail["+strconv.Itoa(numDetail)+"].ValueString", &resource.Entity[numEntity].Detail[numDetail].ValueString, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueBase64Binary(numEntity int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Detail["+strconv.Itoa(numDetail)+"].ValueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Detail["+strconv.Itoa(numDetail)+"].ValueBase64Binary", &resource.Entity[numEntity].Detail[numDetail].ValueBase64Binary, htmlAttrs)
}
