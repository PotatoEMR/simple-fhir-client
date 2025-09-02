package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/AuditEvent
type AuditEvent struct {
	Id                *string            `json:"id,omitempty"`
	Meta              *Meta              `json:"meta,omitempty"`
	ImplicitRules     *string            `json:"implicitRules,omitempty"`
	Language          *string            `json:"language,omitempty"`
	Text              *Narrative         `json:"text,omitempty"`
	Contained         []Resource         `json:"contained,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Category          []CodeableConcept  `json:"category,omitempty"`
	Code              CodeableConcept    `json:"code"`
	Action            *string            `json:"action,omitempty"`
	Severity          *string            `json:"severity,omitempty"`
	OccurredPeriod    *Period            `json:"occurredPeriod"`
	OccurredDateTime  *string            `json:"occurredDateTime"`
	Recorded          string             `json:"recorded"`
	Outcome           *AuditEventOutcome `json:"outcome,omitempty"`
	Authorization     []CodeableConcept  `json:"authorization,omitempty"`
	BasedOn           []Reference        `json:"basedOn,omitempty"`
	Patient           *Reference         `json:"patient,omitempty"`
	Encounter         *Reference         `json:"encounter,omitempty"`
	Agent             []AuditEventAgent  `json:"agent"`
	Source            AuditEventSource   `json:"source"`
	Entity            []AuditEventEntity `json:"entity,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AuditEvent
type AuditEventOutcome struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Code              Coding            `json:"code"`
	Detail            []CodeableConcept `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AuditEvent
type AuditEventAgent struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Role              []CodeableConcept `json:"role,omitempty"`
	Who               Reference         `json:"who"`
	Requestor         *bool             `json:"requestor,omitempty"`
	Location          *Reference        `json:"location,omitempty"`
	Policy            []string          `json:"policy,omitempty"`
	NetworkReference  *Reference        `json:"networkReference"`
	NetworkUri        *string           `json:"networkUri"`
	NetworkString     *string           `json:"networkString"`
	Authorization     []CodeableConcept `json:"authorization,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AuditEvent
type AuditEventSource struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Site              *Reference        `json:"site,omitempty"`
	Observer          Reference         `json:"observer"`
	Type              []CodeableConcept `json:"type,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AuditEvent
type AuditEventEntity struct {
	Id                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	What              *Reference               `json:"what,omitempty"`
	Role              *CodeableConcept         `json:"role,omitempty"`
	SecurityLabel     []CodeableConcept        `json:"securityLabel,omitempty"`
	Query             *string                  `json:"query,omitempty"`
	Detail            []AuditEventEntityDetail `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AuditEvent
type AuditEventEntityDetail struct {
	Id                   *string         `json:"id,omitempty"`
	Extension            []Extension     `json:"extension,omitempty"`
	ModifierExtension    []Extension     `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept `json:"type"`
	ValueQuantity        Quantity        `json:"valueQuantity"`
	ValueCodeableConcept CodeableConcept `json:"valueCodeableConcept"`
	ValueString          string          `json:"valueString"`
	ValueBoolean         bool            `json:"valueBoolean"`
	ValueInteger         int             `json:"valueInteger"`
	ValueRange           Range           `json:"valueRange"`
	ValueRatio           Ratio           `json:"valueRatio"`
	ValueTime            string          `json:"valueTime"`
	ValueDateTime        string          `json:"valueDateTime"`
	ValuePeriod          Period          `json:"valuePeriod"`
	ValueBase64Binary    string          `json:"valueBase64Binary"`
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

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *AuditEvent) AuditEventCategory(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *AuditEvent) AuditEventCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Code, optionsValueSet)
}
func (resource *AuditEvent) AuditEventAction() templ.Component {
	optionsValueSet := VSAudit_event_action

	if resource == nil {
		return CodeSelect("action", nil, optionsValueSet)
	}
	return CodeSelect("action", resource.Action, optionsValueSet)
}
func (resource *AuditEvent) AuditEventSeverity() templ.Component {
	optionsValueSet := VSAudit_event_severity

	if resource == nil {
		return CodeSelect("severity", nil, optionsValueSet)
	}
	return CodeSelect("severity", resource.Severity, optionsValueSet)
}
func (resource *AuditEvent) AuditEventAuthorization(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("authorization", nil, optionsValueSet)
	}
	return CodeableConceptSelect("authorization", &resource.Authorization[0], optionsValueSet)
}
func (resource *AuditEvent) AuditEventOutcomeCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("code", nil, optionsValueSet)
	}
	return CodingSelect("code", &resource.Outcome.Code, optionsValueSet)
}
func (resource *AuditEvent) AuditEventOutcomeDetail(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("detail", nil, optionsValueSet)
	}
	return CodeableConceptSelect("detail", &resource.Outcome.Detail[0], optionsValueSet)
}
func (resource *AuditEvent) AuditEventAgentType(numAgent int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Agent) >= numAgent {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Agent[numAgent].Type, optionsValueSet)
}
func (resource *AuditEvent) AuditEventAgentRole(numAgent int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Agent) >= numAgent {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", &resource.Agent[numAgent].Role[0], optionsValueSet)
}
func (resource *AuditEvent) AuditEventAgentAuthorization(numAgent int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Agent) >= numAgent {
		return CodeableConceptSelect("authorization", nil, optionsValueSet)
	}
	return CodeableConceptSelect("authorization", &resource.Agent[numAgent].Authorization[0], optionsValueSet)
}
func (resource *AuditEvent) AuditEventSourceType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Source.Type[0], optionsValueSet)
}
func (resource *AuditEvent) AuditEventEntityRole(numEntity int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Entity) >= numEntity {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", resource.Entity[numEntity].Role, optionsValueSet)
}
func (resource *AuditEvent) AuditEventEntitySecurityLabel(numEntity int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Entity) >= numEntity {
		return CodeableConceptSelect("securityLabel", nil, optionsValueSet)
	}
	return CodeableConceptSelect("securityLabel", &resource.Entity[numEntity].SecurityLabel[0], optionsValueSet)
}
func (resource *AuditEvent) AuditEventEntityDetailType(numEntity int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Entity[numEntity].Detail) >= numDetail {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Entity[numEntity].Detail[numDetail].Type, optionsValueSet)
}
