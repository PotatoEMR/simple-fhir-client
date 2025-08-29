package r4

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *AuditEvent) AuditEventLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *AuditEvent) AuditEventAction() templ.Component {
	optionsValueSet := VSAudit_event_action
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Action
	}
	return CodeSelect("action", currentVal, optionsValueSet)
}
func (resource *AuditEvent) AuditEventOutcome() templ.Component {
	optionsValueSet := VSAudit_event_outcome
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Outcome
	}
	return CodeSelect("outcome", currentVal, optionsValueSet)
}
func (resource *AuditEvent) AuditEventAgentNetworkType(numAgent int) templ.Component {
	optionsValueSet := VSNetwork_type
	currentVal := ""
	if resource != nil && len(resource.Agent) >= numAgent {
		currentVal = *resource.Agent[numAgent].Network.Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
