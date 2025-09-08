package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	OccurredDateTime  *time.Time         `json:"occurredDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	ValueDateTime        time.Time       `json:"valueDateTime,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *AuditEvent) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_Action(htmlAttrs string) templ.Component {
	optionsValueSet := VSAudit_event_action

	if resource == nil {
		return CodeSelect("Action", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Action", resource.Action, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_Severity(htmlAttrs string) templ.Component {
	optionsValueSet := VSAudit_event_severity

	if resource == nil {
		return CodeSelect("Severity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Severity", resource.Severity, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_OccurredDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("OccurredDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("OccurredDateTime", resource.OccurredDateTime, htmlAttrs)
}
func (resource *AuditEvent) T_Recorded(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Recorded", nil, htmlAttrs)
	}
	return StringInput("Recorded", &resource.Recorded, htmlAttrs)
}
func (resource *AuditEvent) T_Authorization(numAuthorization int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAuthorization >= len(resource.Authorization) {
		return CodeableConceptSelect("Authorization["+strconv.Itoa(numAuthorization)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Authorization["+strconv.Itoa(numAuthorization)+"]", &resource.Authorization[numAuthorization], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_OutcomeCode(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("OutcomeCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("OutcomeCode", &resource.Outcome.Code, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_OutcomeDetail(numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDetail >= len(resource.Outcome.Detail) {
		return CodeableConceptSelect("OutcomeDetail["+strconv.Itoa(numDetail)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OutcomeDetail["+strconv.Itoa(numDetail)+"]", &resource.Outcome.Detail[numDetail], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_AgentType(numAgent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return CodeableConceptSelect("Agent["+strconv.Itoa(numAgent)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Agent["+strconv.Itoa(numAgent)+"]Type", resource.Agent[numAgent].Type, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_AgentRole(numAgent int, numRole int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) || numRole >= len(resource.Agent[numAgent].Role) {
		return CodeableConceptSelect("Agent["+strconv.Itoa(numAgent)+"]Role["+strconv.Itoa(numRole)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Agent["+strconv.Itoa(numAgent)+"]Role["+strconv.Itoa(numRole)+"]", &resource.Agent[numAgent].Role[numRole], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_AgentRequestor(numAgent int, htmlAttrs string) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return BoolInput("Agent["+strconv.Itoa(numAgent)+"]Requestor", nil, htmlAttrs)
	}
	return BoolInput("Agent["+strconv.Itoa(numAgent)+"]Requestor", resource.Agent[numAgent].Requestor, htmlAttrs)
}
func (resource *AuditEvent) T_AgentPolicy(numAgent int, numPolicy int, htmlAttrs string) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) || numPolicy >= len(resource.Agent[numAgent].Policy) {
		return StringInput("Agent["+strconv.Itoa(numAgent)+"]Policy["+strconv.Itoa(numPolicy)+"]", nil, htmlAttrs)
	}
	return StringInput("Agent["+strconv.Itoa(numAgent)+"]Policy["+strconv.Itoa(numPolicy)+"]", &resource.Agent[numAgent].Policy[numPolicy], htmlAttrs)
}
func (resource *AuditEvent) T_AgentNetworkUri(numAgent int, htmlAttrs string) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return StringInput("Agent["+strconv.Itoa(numAgent)+"]NetworkUri", nil, htmlAttrs)
	}
	return StringInput("Agent["+strconv.Itoa(numAgent)+"]NetworkUri", resource.Agent[numAgent].NetworkUri, htmlAttrs)
}
func (resource *AuditEvent) T_AgentNetworkString(numAgent int, htmlAttrs string) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return StringInput("Agent["+strconv.Itoa(numAgent)+"]NetworkString", nil, htmlAttrs)
	}
	return StringInput("Agent["+strconv.Itoa(numAgent)+"]NetworkString", resource.Agent[numAgent].NetworkString, htmlAttrs)
}
func (resource *AuditEvent) T_AgentAuthorization(numAgent int, numAuthorization int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) || numAuthorization >= len(resource.Agent[numAgent].Authorization) {
		return CodeableConceptSelect("Agent["+strconv.Itoa(numAgent)+"]Authorization["+strconv.Itoa(numAuthorization)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Agent["+strconv.Itoa(numAgent)+"]Authorization["+strconv.Itoa(numAuthorization)+"]", &resource.Agent[numAgent].Authorization[numAuthorization], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_SourceType(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numType >= len(resource.Source.Type) {
		return CodeableConceptSelect("SourceType["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SourceType["+strconv.Itoa(numType)+"]", &resource.Source.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntityRole(numEntity int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return CodeableConceptSelect("Entity["+strconv.Itoa(numEntity)+"]Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Entity["+strconv.Itoa(numEntity)+"]Role", resource.Entity[numEntity].Role, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntitySecurityLabel(numEntity int, numSecurityLabel int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numSecurityLabel >= len(resource.Entity[numEntity].SecurityLabel) {
		return CodeableConceptSelect("Entity["+strconv.Itoa(numEntity)+"]SecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Entity["+strconv.Itoa(numEntity)+"]SecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", &resource.Entity[numEntity].SecurityLabel[numSecurityLabel], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntityQuery(numEntity int, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return StringInput("Entity["+strconv.Itoa(numEntity)+"]Query", nil, htmlAttrs)
	}
	return StringInput("Entity["+strconv.Itoa(numEntity)+"]Query", resource.Entity[numEntity].Query, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailType(numEntity int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return CodeableConceptSelect("Entity["+strconv.Itoa(numEntity)+"]Detail["+strconv.Itoa(numDetail)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Entity["+strconv.Itoa(numEntity)+"]Detail["+strconv.Itoa(numDetail)+"].Type", &resource.Entity[numEntity].Detail[numDetail].Type, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueCodeableConcept(numEntity int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return CodeableConceptSelect("Entity["+strconv.Itoa(numEntity)+"]Detail["+strconv.Itoa(numDetail)+"].ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Entity["+strconv.Itoa(numEntity)+"]Detail["+strconv.Itoa(numDetail)+"].ValueCodeableConcept", &resource.Entity[numEntity].Detail[numDetail].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueString(numEntity int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return StringInput("Entity["+strconv.Itoa(numEntity)+"]Detail["+strconv.Itoa(numDetail)+"].ValueString", nil, htmlAttrs)
	}
	return StringInput("Entity["+strconv.Itoa(numEntity)+"]Detail["+strconv.Itoa(numDetail)+"].ValueString", &resource.Entity[numEntity].Detail[numDetail].ValueString, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueBoolean(numEntity int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return BoolInput("Entity["+strconv.Itoa(numEntity)+"]Detail["+strconv.Itoa(numDetail)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Entity["+strconv.Itoa(numEntity)+"]Detail["+strconv.Itoa(numDetail)+"].ValueBoolean", &resource.Entity[numEntity].Detail[numDetail].ValueBoolean, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueInteger(numEntity int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return IntInput("Entity["+strconv.Itoa(numEntity)+"]Detail["+strconv.Itoa(numDetail)+"].ValueInteger", nil, htmlAttrs)
	}
	return IntInput("Entity["+strconv.Itoa(numEntity)+"]Detail["+strconv.Itoa(numDetail)+"].ValueInteger", &resource.Entity[numEntity].Detail[numDetail].ValueInteger, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueTime(numEntity int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return StringInput("Entity["+strconv.Itoa(numEntity)+"]Detail["+strconv.Itoa(numDetail)+"].ValueTime", nil, htmlAttrs)
	}
	return StringInput("Entity["+strconv.Itoa(numEntity)+"]Detail["+strconv.Itoa(numDetail)+"].ValueTime", &resource.Entity[numEntity].Detail[numDetail].ValueTime, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueDateTime(numEntity int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return DateTimeInput("Entity["+strconv.Itoa(numEntity)+"]Detail["+strconv.Itoa(numDetail)+"].ValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Entity["+strconv.Itoa(numEntity)+"]Detail["+strconv.Itoa(numDetail)+"].ValueDateTime", &resource.Entity[numEntity].Detail[numDetail].ValueDateTime, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueBase64Binary(numEntity int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return StringInput("Entity["+strconv.Itoa(numEntity)+"]Detail["+strconv.Itoa(numDetail)+"].ValueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("Entity["+strconv.Itoa(numEntity)+"]Detail["+strconv.Itoa(numDetail)+"].ValueBase64Binary", &resource.Entity[numEntity].Detail[numDetail].ValueBase64Binary, htmlAttrs)
}
