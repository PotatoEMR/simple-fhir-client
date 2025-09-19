package r5

//generated with command go run ./bultaoreune
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
	OccurredDateTime  *FhirDateTime      `json:"occurredDateTime,omitempty"`
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
	ValueDateTime        FhirDateTime    `json:"valueDateTime"`
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
func (resource *AuditEvent) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_Action(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAudit_event_action

	if resource == nil {
		return CodeSelect("action", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("action", resource.Action, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_Severity(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAudit_event_severity

	if resource == nil {
		return CodeSelect("severity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("severity", resource.Severity, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_OccurredPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("occurredPeriod", nil, htmlAttrs)
	}
	return PeriodInput("occurredPeriod", resource.OccurredPeriod, htmlAttrs)
}
func (resource *AuditEvent) T_OccurredDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("occurredDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("occurredDateTime", resource.OccurredDateTime, htmlAttrs)
}
func (resource *AuditEvent) T_Recorded(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("recorded", nil, htmlAttrs)
	}
	return StringInput("recorded", &resource.Recorded, htmlAttrs)
}
func (resource *AuditEvent) T_Authorization(numAuthorization int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthorization >= len(resource.Authorization) {
		return CodeableConceptSelect("authorization["+strconv.Itoa(numAuthorization)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("authorization["+strconv.Itoa(numAuthorization)+"]", &resource.Authorization[numAuthorization], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_BasedOn(numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput("basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *AuditEvent) T_Patient(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("patient", nil, htmlAttrs)
	}
	return ReferenceInput("patient", resource.Patient, htmlAttrs)
}
func (resource *AuditEvent) T_Encounter(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("encounter", nil, htmlAttrs)
	}
	return ReferenceInput("encounter", resource.Encounter, htmlAttrs)
}
func (resource *AuditEvent) T_OutcomeCode(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("outcome.code", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("outcome.code", &resource.Outcome.Code, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_OutcomeDetail(numDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDetail >= len(resource.Outcome.Detail) {
		return CodeableConceptSelect("outcome.detail["+strconv.Itoa(numDetail)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("outcome.detail["+strconv.Itoa(numDetail)+"]", &resource.Outcome.Detail[numDetail], optionsValueSet, htmlAttrs)
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
func (resource *AuditEvent) T_AgentWho(numAgent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return ReferenceInput("agent["+strconv.Itoa(numAgent)+"].who", nil, htmlAttrs)
	}
	return ReferenceInput("agent["+strconv.Itoa(numAgent)+"].who", &resource.Agent[numAgent].Who, htmlAttrs)
}
func (resource *AuditEvent) T_AgentRequestor(numAgent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return BoolInput("agent["+strconv.Itoa(numAgent)+"].requestor", nil, htmlAttrs)
	}
	return BoolInput("agent["+strconv.Itoa(numAgent)+"].requestor", resource.Agent[numAgent].Requestor, htmlAttrs)
}
func (resource *AuditEvent) T_AgentLocation(numAgent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return ReferenceInput("agent["+strconv.Itoa(numAgent)+"].location", nil, htmlAttrs)
	}
	return ReferenceInput("agent["+strconv.Itoa(numAgent)+"].location", resource.Agent[numAgent].Location, htmlAttrs)
}
func (resource *AuditEvent) T_AgentPolicy(numAgent int, numPolicy int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) || numPolicy >= len(resource.Agent[numAgent].Policy) {
		return StringInput("agent["+strconv.Itoa(numAgent)+"].policy["+strconv.Itoa(numPolicy)+"]", nil, htmlAttrs)
	}
	return StringInput("agent["+strconv.Itoa(numAgent)+"].policy["+strconv.Itoa(numPolicy)+"]", &resource.Agent[numAgent].Policy[numPolicy], htmlAttrs)
}
func (resource *AuditEvent) T_AgentNetworkReference(numAgent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return ReferenceInput("agent["+strconv.Itoa(numAgent)+"].networkReference", nil, htmlAttrs)
	}
	return ReferenceInput("agent["+strconv.Itoa(numAgent)+"].networkReference", resource.Agent[numAgent].NetworkReference, htmlAttrs)
}
func (resource *AuditEvent) T_AgentNetworkUri(numAgent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return StringInput("agent["+strconv.Itoa(numAgent)+"].networkUri", nil, htmlAttrs)
	}
	return StringInput("agent["+strconv.Itoa(numAgent)+"].networkUri", resource.Agent[numAgent].NetworkUri, htmlAttrs)
}
func (resource *AuditEvent) T_AgentNetworkString(numAgent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) {
		return StringInput("agent["+strconv.Itoa(numAgent)+"].networkString", nil, htmlAttrs)
	}
	return StringInput("agent["+strconv.Itoa(numAgent)+"].networkString", resource.Agent[numAgent].NetworkString, htmlAttrs)
}
func (resource *AuditEvent) T_AgentAuthorization(numAgent int, numAuthorization int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAgent >= len(resource.Agent) || numAuthorization >= len(resource.Agent[numAgent].Authorization) {
		return CodeableConceptSelect("agent["+strconv.Itoa(numAgent)+"].authorization["+strconv.Itoa(numAuthorization)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("agent["+strconv.Itoa(numAgent)+"].authorization["+strconv.Itoa(numAuthorization)+"]", &resource.Agent[numAgent].Authorization[numAuthorization], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_SourceSite(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("source.site", nil, htmlAttrs)
	}
	return ReferenceInput("source.site", resource.Source.Site, htmlAttrs)
}
func (resource *AuditEvent) T_SourceObserver(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("source.observer", nil, htmlAttrs)
	}
	return ReferenceInput("source.observer", &resource.Source.Observer, htmlAttrs)
}
func (resource *AuditEvent) T_SourceType(numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numType >= len(resource.Source.Type) {
		return CodeableConceptSelect("source.type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("source.type["+strconv.Itoa(numType)+"]", &resource.Source.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntityWhat(numEntity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return ReferenceInput("entity["+strconv.Itoa(numEntity)+"].what", nil, htmlAttrs)
	}
	return ReferenceInput("entity["+strconv.Itoa(numEntity)+"].what", resource.Entity[numEntity].What, htmlAttrs)
}
func (resource *AuditEvent) T_EntityRole(numEntity int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return CodeableConceptSelect("entity["+strconv.Itoa(numEntity)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("entity["+strconv.Itoa(numEntity)+"].role", resource.Entity[numEntity].Role, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntitySecurityLabel(numEntity int, numSecurityLabel int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numSecurityLabel >= len(resource.Entity[numEntity].SecurityLabel) {
		return CodeableConceptSelect("entity["+strconv.Itoa(numEntity)+"].securityLabel["+strconv.Itoa(numSecurityLabel)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("entity["+strconv.Itoa(numEntity)+"].securityLabel["+strconv.Itoa(numSecurityLabel)+"]", &resource.Entity[numEntity].SecurityLabel[numSecurityLabel], optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntityQuery(numEntity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) {
		return StringInput("entity["+strconv.Itoa(numEntity)+"].query", nil, htmlAttrs)
	}
	return StringInput("entity["+strconv.Itoa(numEntity)+"].query", resource.Entity[numEntity].Query, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailType(numEntity int, numDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return CodeableConceptSelect("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].type", &resource.Entity[numEntity].Detail[numDetail].Type, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueQuantity(numEntity int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return QuantityInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueQuantity", nil, htmlAttrs)
	}
	return QuantityInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueQuantity", &resource.Entity[numEntity].Detail[numDetail].ValueQuantity, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueCodeableConcept(numEntity int, numDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return CodeableConceptSelect("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueCodeableConcept", &resource.Entity[numEntity].Detail[numDetail].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueString(numEntity int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return StringInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueString", &resource.Entity[numEntity].Detail[numDetail].ValueString, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueBoolean(numEntity int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return BoolInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueBoolean", &resource.Entity[numEntity].Detail[numDetail].ValueBoolean, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueInteger(numEntity int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return IntInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueInteger", &resource.Entity[numEntity].Detail[numDetail].ValueInteger, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueRange(numEntity int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return RangeInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueRange", &resource.Entity[numEntity].Detail[numDetail].ValueRange, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueRatio(numEntity int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return RatioInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueRatio", nil, htmlAttrs)
	}
	return RatioInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueRatio", &resource.Entity[numEntity].Detail[numDetail].ValueRatio, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueTime(numEntity int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return StringInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueTime", nil, htmlAttrs)
	}
	return StringInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueTime", &resource.Entity[numEntity].Detail[numDetail].ValueTime, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueDateTime(numEntity int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return FhirDateTimeInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueDateTime", &resource.Entity[numEntity].Detail[numDetail].ValueDateTime, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValuePeriod(numEntity int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return PeriodInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valuePeriod", nil, htmlAttrs)
	}
	return PeriodInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valuePeriod", &resource.Entity[numEntity].Detail[numDetail].ValuePeriod, htmlAttrs)
}
func (resource *AuditEvent) T_EntityDetailValueBase64Binary(numEntity int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntity >= len(resource.Entity) || numDetail >= len(resource.Entity[numEntity].Detail) {
		return StringInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("entity["+strconv.Itoa(numEntity)+"].detail["+strconv.Itoa(numDetail)+"].valueBase64Binary", &resource.Entity[numEntity].Detail[numDetail].ValueBase64Binary, htmlAttrs)
}
