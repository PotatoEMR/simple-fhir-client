package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	OccurredPeriod    *Period            `json:"occurredPeriod,omitempty"`
	OccurredDateTime  *string            `json:"occurredDateTime,omitempty"`
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
	NetworkReference  *Reference        `json:"networkReference,omitempty"`
	NetworkUri        *string           `json:"networkUri,omitempty"`
	NetworkString     *string           `json:"networkString,omitempty"`
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
func (resource *AuditEvent) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("AuditEvent.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AuditEvent.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *AuditEvent) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("AuditEvent.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AuditEvent.Code", &resource.Code, optionsValueSet)
}
func (resource *AuditEvent) T_Action() templ.Component {
	optionsValueSet := VSAudit_event_action

	if resource == nil {
		return CodeSelect("AuditEvent.Action", nil, optionsValueSet)
	}
	return CodeSelect("AuditEvent.Action", resource.Action, optionsValueSet)
}
func (resource *AuditEvent) T_Severity() templ.Component {
	optionsValueSet := VSAudit_event_severity

	if resource == nil {
		return CodeSelect("AuditEvent.Severity", nil, optionsValueSet)
	}
	return CodeSelect("AuditEvent.Severity", resource.Severity, optionsValueSet)
}
func (resource *AuditEvent) T_Recorded() templ.Component {

	if resource == nil {
		return StringInput("AuditEvent.Recorded", nil)
	}
	return StringInput("AuditEvent.Recorded", &resource.Recorded)
}
func (resource *AuditEvent) T_Authorization(numAuthorization int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Authorization) >= numAuthorization {
		return CodeableConceptSelect("AuditEvent.Authorization["+strconv.Itoa(numAuthorization)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AuditEvent.Authorization["+strconv.Itoa(numAuthorization)+"]", &resource.Authorization[numAuthorization], optionsValueSet)
}
func (resource *AuditEvent) T_OutcomeId() templ.Component {

	if resource == nil {
		return StringInput("AuditEvent.Outcome.Id", nil)
	}
	return StringInput("AuditEvent.Outcome.Id", resource.Outcome.Id)
}
func (resource *AuditEvent) T_OutcomeCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("AuditEvent.Outcome.Code", nil, optionsValueSet)
	}
	return CodingSelect("AuditEvent.Outcome.Code", &resource.Outcome.Code, optionsValueSet)
}
func (resource *AuditEvent) T_OutcomeDetail(numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Outcome.Detail) >= numDetail {
		return CodeableConceptSelect("AuditEvent.Outcome.Detail["+strconv.Itoa(numDetail)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AuditEvent.Outcome.Detail["+strconv.Itoa(numDetail)+"]", &resource.Outcome.Detail[numDetail], optionsValueSet)
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
func (resource *AuditEvent) T_AgentRequestor(numAgent int) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent {
		return BoolInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Requestor", nil)
	}
	return BoolInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Requestor", resource.Agent[numAgent].Requestor)
}
func (resource *AuditEvent) T_AgentPolicy(numAgent int, numPolicy int) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent || len(resource.Agent[numAgent].Policy) >= numPolicy {
		return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Policy["+strconv.Itoa(numPolicy)+"]", nil)
	}
	return StringInput("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Policy["+strconv.Itoa(numPolicy)+"]", &resource.Agent[numAgent].Policy[numPolicy])
}
func (resource *AuditEvent) T_AgentAuthorization(numAgent int, numAuthorization int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Agent) >= numAgent || len(resource.Agent[numAgent].Authorization) >= numAuthorization {
		return CodeableConceptSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Authorization["+strconv.Itoa(numAuthorization)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AuditEvent.Agent["+strconv.Itoa(numAgent)+"].Authorization["+strconv.Itoa(numAuthorization)+"]", &resource.Agent[numAgent].Authorization[numAuthorization], optionsValueSet)
}
func (resource *AuditEvent) T_SourceId() templ.Component {

	if resource == nil {
		return StringInput("AuditEvent.Source.Id", nil)
	}
	return StringInput("AuditEvent.Source.Id", resource.Source.Id)
}
func (resource *AuditEvent) T_SourceType(numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Source.Type) >= numType {
		return CodeableConceptSelect("AuditEvent.Source.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AuditEvent.Source.Type["+strconv.Itoa(numType)+"]", &resource.Source.Type[numType], optionsValueSet)
}
func (resource *AuditEvent) T_EntityId(numEntity int) templ.Component {

	if resource == nil || len(resource.Entity) >= numEntity {
		return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Id", nil)
	}
	return StringInput("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Id", resource.Entity[numEntity].Id)
}
func (resource *AuditEvent) T_EntityRole(numEntity int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Entity) >= numEntity {
		return CodeableConceptSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Role", resource.Entity[numEntity].Role, optionsValueSet)
}
func (resource *AuditEvent) T_EntitySecurityLabel(numEntity int, numSecurityLabel int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Entity) >= numEntity || len(resource.Entity[numEntity].SecurityLabel) >= numSecurityLabel {
		return CodeableConceptSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", &resource.Entity[numEntity].SecurityLabel[numSecurityLabel], optionsValueSet)
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
func (resource *AuditEvent) T_EntityDetailType(numEntity int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Entity) >= numEntity || len(resource.Entity[numEntity].Detail) >= numDetail {
		return CodeableConceptSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Detail["+strconv.Itoa(numDetail)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("AuditEvent.Entity["+strconv.Itoa(numEntity)+"].Detail["+strconv.Itoa(numDetail)+"].Type", &resource.Entity[numEntity].Detail[numDetail].Type, optionsValueSet)
}
