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
func (resource *AuditEvent) AuditEventSeverity() templ.Component {
	optionsValueSet := VSAudit_event_severity
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Severity
	}
	return CodeSelect("severity", currentVal, optionsValueSet)
}
