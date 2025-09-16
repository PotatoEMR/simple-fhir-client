package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/SubscriptionTopic
type SubscriptionTopic struct {
	Id                *string                              `json:"id,omitempty"`
	Meta              *Meta                                `json:"meta,omitempty"`
	ImplicitRules     *string                              `json:"implicitRules,omitempty"`
	Language          *string                              `json:"language,omitempty"`
	Text              *Narrative                           `json:"text,omitempty"`
	Contained         []Resource                           `json:"contained,omitempty"`
	Extension         []Extension                          `json:"extension,omitempty"`
	ModifierExtension []Extension                          `json:"modifierExtension,omitempty"`
	Url               string                               `json:"url"`
	Identifier        []Identifier                         `json:"identifier,omitempty"`
	Version           *string                              `json:"version,omitempty"`
	Title             *string                              `json:"title,omitempty"`
	DerivedFrom       []string                             `json:"derivedFrom,omitempty"`
	Status            string                               `json:"status"`
	Experimental      *bool                                `json:"experimental,omitempty"`
	Date              *FhirDateTime                        `json:"date,omitempty"`
	Publisher         *string                              `json:"publisher,omitempty"`
	Contact           []ContactDetail                      `json:"contact,omitempty"`
	Description       *string                              `json:"description,omitempty"`
	UseContext        []UsageContext                       `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                    `json:"jurisdiction,omitempty"`
	Purpose           *string                              `json:"purpose,omitempty"`
	Copyright         *string                              `json:"copyright,omitempty"`
	ApprovalDate      *FhirDate                            `json:"approvalDate,omitempty"`
	LastReviewDate    *FhirDate                            `json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period                              `json:"effectivePeriod,omitempty"`
	ResourceTrigger   []SubscriptionTopicResourceTrigger   `json:"resourceTrigger,omitempty"`
	EventTrigger      []SubscriptionTopicEventTrigger      `json:"eventTrigger,omitempty"`
	CanFilterBy       []SubscriptionTopicCanFilterBy       `json:"canFilterBy,omitempty"`
	NotificationShape []SubscriptionTopicNotificationShape `json:"notificationShape,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubscriptionTopic
type SubscriptionTopicResourceTrigger struct {
	Id                   *string                                        `json:"id,omitempty"`
	Extension            []Extension                                    `json:"extension,omitempty"`
	ModifierExtension    []Extension                                    `json:"modifierExtension,omitempty"`
	Description          *string                                        `json:"description,omitempty"`
	Resource             string                                         `json:"resource"`
	SupportedInteraction []string                                       `json:"supportedInteraction,omitempty"`
	QueryCriteria        *SubscriptionTopicResourceTriggerQueryCriteria `json:"queryCriteria,omitempty"`
	FhirPathCriteria     *string                                        `json:"fhirPathCriteria,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubscriptionTopic
type SubscriptionTopicResourceTriggerQueryCriteria struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Previous          *string     `json:"previous,omitempty"`
	ResultForCreate   *string     `json:"resultForCreate,omitempty"`
	Current           *string     `json:"current,omitempty"`
	ResultForDelete   *string     `json:"resultForDelete,omitempty"`
	RequireBoth       *bool       `json:"requireBoth,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubscriptionTopic
type SubscriptionTopicEventTrigger struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Description       *string         `json:"description,omitempty"`
	Event             CodeableConcept `json:"event"`
	Resource          string          `json:"resource"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubscriptionTopic
type SubscriptionTopicCanFilterBy struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Resource          *string     `json:"resource,omitempty"`
	FilterParameter   string      `json:"filterParameter"`
	FilterDefinition  *string     `json:"filterDefinition,omitempty"`
	Modifier          []string    `json:"modifier,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubscriptionTopic
type SubscriptionTopicNotificationShape struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Resource          string      `json:"resource"`
	Include           []string    `json:"include,omitempty"`
	RevInclude        []string    `json:"revInclude,omitempty"`
}

type OtherSubscriptionTopic SubscriptionTopic

// on convert struct to json, automatically add resourceType=SubscriptionTopic
func (r SubscriptionTopic) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubscriptionTopic
		ResourceType string `json:"resourceType"`
	}{
		OtherSubscriptionTopic: OtherSubscriptionTopic(r),
		ResourceType:           "SubscriptionTopic",
	})
}
func (r SubscriptionTopic) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "SubscriptionTopic/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "SubscriptionTopic"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *SubscriptionTopic) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", &resource.Url, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *SubscriptionTopic) T_DerivedFrom(numDerivedFrom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDerivedFrom >= len(resource.DerivedFrom) {
		return StringInput("derivedFrom["+strconv.Itoa(numDerivedFrom)+"]", nil, htmlAttrs)
	}
	return StringInput("derivedFrom["+strconv.Itoa(numDerivedFrom)+"]", &resource.DerivedFrom[numDerivedFrom], htmlAttrs)
}
func (resource *SubscriptionTopic) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateInput("approvalDate", nil, htmlAttrs)
	}
	return DateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *SubscriptionTopic) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateInput("lastReviewDate", nil, htmlAttrs)
	}
	return DateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerDescription(numResourceTrigger int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) {
		return StringInput("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].description", nil, htmlAttrs)
	}
	return StringInput("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].description", resource.ResourceTrigger[numResourceTrigger].Description, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerResource(numResourceTrigger int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) {
		return StringInput("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].resource", nil, htmlAttrs)
	}
	return StringInput("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].resource", &resource.ResourceTrigger[numResourceTrigger].Resource, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerSupportedInteraction(numResourceTrigger int, numSupportedInteraction int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSInteraction_trigger

	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) || numSupportedInteraction >= len(resource.ResourceTrigger[numResourceTrigger].SupportedInteraction) {
		return CodeSelect("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].supportedInteraction["+strconv.Itoa(numSupportedInteraction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].supportedInteraction["+strconv.Itoa(numSupportedInteraction)+"]", &resource.ResourceTrigger[numResourceTrigger].SupportedInteraction[numSupportedInteraction], optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerFhirPathCriteria(numResourceTrigger int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) {
		return StringInput("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].fhirPathCriteria", nil, htmlAttrs)
	}
	return StringInput("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].fhirPathCriteria", resource.ResourceTrigger[numResourceTrigger].FhirPathCriteria, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaPrevious(numResourceTrigger int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) {
		return StringInput("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].queryCriteria.previous", nil, htmlAttrs)
	}
	return StringInput("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].queryCriteria.previous", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.Previous, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaResultForCreate(numResourceTrigger int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSubscriptiontopic_cr_behavior

	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) {
		return CodeSelect("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].queryCriteria.resultForCreate", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].queryCriteria.resultForCreate", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.ResultForCreate, optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaCurrent(numResourceTrigger int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) {
		return StringInput("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].queryCriteria.current", nil, htmlAttrs)
	}
	return StringInput("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].queryCriteria.current", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.Current, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaResultForDelete(numResourceTrigger int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSubscriptiontopic_cr_behavior

	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) {
		return CodeSelect("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].queryCriteria.resultForDelete", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].queryCriteria.resultForDelete", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.ResultForDelete, optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaRequireBoth(numResourceTrigger int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) {
		return BoolInput("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].queryCriteria.requireBoth", nil, htmlAttrs)
	}
	return BoolInput("resourceTrigger["+strconv.Itoa(numResourceTrigger)+"].queryCriteria.requireBoth", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.RequireBoth, htmlAttrs)
}
func (resource *SubscriptionTopic) T_EventTriggerDescription(numEventTrigger int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEventTrigger >= len(resource.EventTrigger) {
		return StringInput("eventTrigger["+strconv.Itoa(numEventTrigger)+"].description", nil, htmlAttrs)
	}
	return StringInput("eventTrigger["+strconv.Itoa(numEventTrigger)+"].description", resource.EventTrigger[numEventTrigger].Description, htmlAttrs)
}
func (resource *SubscriptionTopic) T_EventTriggerEvent(numEventTrigger int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEventTrigger >= len(resource.EventTrigger) {
		return CodeableConceptSelect("eventTrigger["+strconv.Itoa(numEventTrigger)+"].event", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("eventTrigger["+strconv.Itoa(numEventTrigger)+"].event", &resource.EventTrigger[numEventTrigger].Event, optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionTopic) T_EventTriggerResource(numEventTrigger int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEventTrigger >= len(resource.EventTrigger) {
		return StringInput("eventTrigger["+strconv.Itoa(numEventTrigger)+"].resource", nil, htmlAttrs)
	}
	return StringInput("eventTrigger["+strconv.Itoa(numEventTrigger)+"].resource", &resource.EventTrigger[numEventTrigger].Resource, htmlAttrs)
}
func (resource *SubscriptionTopic) T_CanFilterByDescription(numCanFilterBy int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCanFilterBy >= len(resource.CanFilterBy) {
		return StringInput("canFilterBy["+strconv.Itoa(numCanFilterBy)+"].description", nil, htmlAttrs)
	}
	return StringInput("canFilterBy["+strconv.Itoa(numCanFilterBy)+"].description", resource.CanFilterBy[numCanFilterBy].Description, htmlAttrs)
}
func (resource *SubscriptionTopic) T_CanFilterByResource(numCanFilterBy int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCanFilterBy >= len(resource.CanFilterBy) {
		return StringInput("canFilterBy["+strconv.Itoa(numCanFilterBy)+"].resource", nil, htmlAttrs)
	}
	return StringInput("canFilterBy["+strconv.Itoa(numCanFilterBy)+"].resource", resource.CanFilterBy[numCanFilterBy].Resource, htmlAttrs)
}
func (resource *SubscriptionTopic) T_CanFilterByFilterParameter(numCanFilterBy int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCanFilterBy >= len(resource.CanFilterBy) {
		return StringInput("canFilterBy["+strconv.Itoa(numCanFilterBy)+"].filterParameter", nil, htmlAttrs)
	}
	return StringInput("canFilterBy["+strconv.Itoa(numCanFilterBy)+"].filterParameter", &resource.CanFilterBy[numCanFilterBy].FilterParameter, htmlAttrs)
}
func (resource *SubscriptionTopic) T_CanFilterByFilterDefinition(numCanFilterBy int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCanFilterBy >= len(resource.CanFilterBy) {
		return StringInput("canFilterBy["+strconv.Itoa(numCanFilterBy)+"].filterDefinition", nil, htmlAttrs)
	}
	return StringInput("canFilterBy["+strconv.Itoa(numCanFilterBy)+"].filterDefinition", resource.CanFilterBy[numCanFilterBy].FilterDefinition, htmlAttrs)
}
func (resource *SubscriptionTopic) T_CanFilterByModifier(numCanFilterBy int, numModifier int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSubscription_search_modifier

	if resource == nil || numCanFilterBy >= len(resource.CanFilterBy) || numModifier >= len(resource.CanFilterBy[numCanFilterBy].Modifier) {
		return CodeSelect("canFilterBy["+strconv.Itoa(numCanFilterBy)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("canFilterBy["+strconv.Itoa(numCanFilterBy)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.CanFilterBy[numCanFilterBy].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionTopic) T_NotificationShapeResource(numNotificationShape int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNotificationShape >= len(resource.NotificationShape) {
		return StringInput("notificationShape["+strconv.Itoa(numNotificationShape)+"].resource", nil, htmlAttrs)
	}
	return StringInput("notificationShape["+strconv.Itoa(numNotificationShape)+"].resource", &resource.NotificationShape[numNotificationShape].Resource, htmlAttrs)
}
func (resource *SubscriptionTopic) T_NotificationShapeInclude(numNotificationShape int, numInclude int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNotificationShape >= len(resource.NotificationShape) || numInclude >= len(resource.NotificationShape[numNotificationShape].Include) {
		return StringInput("notificationShape["+strconv.Itoa(numNotificationShape)+"].include["+strconv.Itoa(numInclude)+"]", nil, htmlAttrs)
	}
	return StringInput("notificationShape["+strconv.Itoa(numNotificationShape)+"].include["+strconv.Itoa(numInclude)+"]", &resource.NotificationShape[numNotificationShape].Include[numInclude], htmlAttrs)
}
func (resource *SubscriptionTopic) T_NotificationShapeRevInclude(numNotificationShape int, numRevInclude int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNotificationShape >= len(resource.NotificationShape) || numRevInclude >= len(resource.NotificationShape[numNotificationShape].RevInclude) {
		return StringInput("notificationShape["+strconv.Itoa(numNotificationShape)+"].revInclude["+strconv.Itoa(numRevInclude)+"]", nil, htmlAttrs)
	}
	return StringInput("notificationShape["+strconv.Itoa(numNotificationShape)+"].revInclude["+strconv.Itoa(numRevInclude)+"]", &resource.NotificationShape[numNotificationShape].RevInclude[numRevInclude], htmlAttrs)
}
