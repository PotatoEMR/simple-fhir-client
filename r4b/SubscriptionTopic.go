package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Date              *time.Time                           `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher         *string                              `json:"publisher,omitempty"`
	Contact           []ContactDetail                      `json:"contact,omitempty"`
	Description       *string                              `json:"description,omitempty"`
	UseContext        []UsageContext                       `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                    `json:"jurisdiction,omitempty"`
	Purpose           *string                              `json:"purpose,omitempty"`
	Copyright         *string                              `json:"copyright,omitempty"`
	ApprovalDate      *time.Time                           `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate    *time.Time                           `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
func (resource *SubscriptionTopic) T_Url(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Url", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.Url", &resource.Url, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Version", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.Version", resource.Version, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Title", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.Title", resource.Title, htmlAttrs)
}
func (resource *SubscriptionTopic) T_DerivedFrom(numDerivedFrom int, htmlAttrs string) templ.Component {

	if resource == nil || numDerivedFrom >= len(resource.DerivedFrom) {
		return StringInput("SubscriptionTopic.DerivedFrom."+strconv.Itoa(numDerivedFrom)+".", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.DerivedFrom."+strconv.Itoa(numDerivedFrom)+".", &resource.DerivedFrom[numDerivedFrom], htmlAttrs)
}
func (resource *SubscriptionTopic) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("SubscriptionTopic.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SubscriptionTopic.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Experimental(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("SubscriptionTopic.Experimental", nil, htmlAttrs)
	}
	return BoolInput("SubscriptionTopic.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Date(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("SubscriptionTopic.Date", nil, htmlAttrs)
	}
	return DateTimeInput("SubscriptionTopic.Date", resource.Date, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Publisher(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Publisher", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Description", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.Description", resource.Description, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("SubscriptionTopic.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubscriptionTopic.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Purpose(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Purpose", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *SubscriptionTopic) T_Copyright(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Copyright", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ApprovalDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("SubscriptionTopic.ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("SubscriptionTopic.ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *SubscriptionTopic) T_LastReviewDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("SubscriptionTopic.LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("SubscriptionTopic.LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerDescription(numResourceTrigger int, htmlAttrs string) templ.Component {

	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) {
		return StringInput("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..Description", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..Description", resource.ResourceTrigger[numResourceTrigger].Description, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerResource(numResourceTrigger int, htmlAttrs string) templ.Component {

	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) {
		return StringInput("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..Resource", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..Resource", &resource.ResourceTrigger[numResourceTrigger].Resource, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerSupportedInteraction(numResourceTrigger int, numSupportedInteraction int, htmlAttrs string) templ.Component {
	optionsValueSet := VSInteraction_trigger

	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) || numSupportedInteraction >= len(resource.ResourceTrigger[numResourceTrigger].SupportedInteraction) {
		return CodeSelect("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..SupportedInteraction."+strconv.Itoa(numSupportedInteraction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..SupportedInteraction."+strconv.Itoa(numSupportedInteraction)+".", &resource.ResourceTrigger[numResourceTrigger].SupportedInteraction[numSupportedInteraction], optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerFhirPathCriteria(numResourceTrigger int, htmlAttrs string) templ.Component {

	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) {
		return StringInput("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..FhirPathCriteria", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..FhirPathCriteria", resource.ResourceTrigger[numResourceTrigger].FhirPathCriteria, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaPrevious(numResourceTrigger int, htmlAttrs string) templ.Component {

	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) {
		return StringInput("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..QueryCriteria.Previous", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..QueryCriteria.Previous", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.Previous, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaResultForCreate(numResourceTrigger int, htmlAttrs string) templ.Component {
	optionsValueSet := VSSubscriptiontopic_cr_behavior

	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) {
		return CodeSelect("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..QueryCriteria.ResultForCreate", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..QueryCriteria.ResultForCreate", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.ResultForCreate, optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaCurrent(numResourceTrigger int, htmlAttrs string) templ.Component {

	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) {
		return StringInput("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..QueryCriteria.Current", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..QueryCriteria.Current", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.Current, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaResultForDelete(numResourceTrigger int, htmlAttrs string) templ.Component {
	optionsValueSet := VSSubscriptiontopic_cr_behavior

	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) {
		return CodeSelect("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..QueryCriteria.ResultForDelete", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..QueryCriteria.ResultForDelete", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.ResultForDelete, optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaRequireBoth(numResourceTrigger int, htmlAttrs string) templ.Component {

	if resource == nil || numResourceTrigger >= len(resource.ResourceTrigger) {
		return BoolInput("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..QueryCriteria.RequireBoth", nil, htmlAttrs)
	}
	return BoolInput("SubscriptionTopic.ResourceTrigger."+strconv.Itoa(numResourceTrigger)+"..QueryCriteria.RequireBoth", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.RequireBoth, htmlAttrs)
}
func (resource *SubscriptionTopic) T_EventTriggerDescription(numEventTrigger int, htmlAttrs string) templ.Component {

	if resource == nil || numEventTrigger >= len(resource.EventTrigger) {
		return StringInput("SubscriptionTopic.EventTrigger."+strconv.Itoa(numEventTrigger)+"..Description", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.EventTrigger."+strconv.Itoa(numEventTrigger)+"..Description", resource.EventTrigger[numEventTrigger].Description, htmlAttrs)
}
func (resource *SubscriptionTopic) T_EventTriggerEvent(numEventTrigger int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numEventTrigger >= len(resource.EventTrigger) {
		return CodeableConceptSelect("SubscriptionTopic.EventTrigger."+strconv.Itoa(numEventTrigger)+"..Event", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubscriptionTopic.EventTrigger."+strconv.Itoa(numEventTrigger)+"..Event", &resource.EventTrigger[numEventTrigger].Event, optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionTopic) T_EventTriggerResource(numEventTrigger int, htmlAttrs string) templ.Component {

	if resource == nil || numEventTrigger >= len(resource.EventTrigger) {
		return StringInput("SubscriptionTopic.EventTrigger."+strconv.Itoa(numEventTrigger)+"..Resource", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.EventTrigger."+strconv.Itoa(numEventTrigger)+"..Resource", &resource.EventTrigger[numEventTrigger].Resource, htmlAttrs)
}
func (resource *SubscriptionTopic) T_CanFilterByDescription(numCanFilterBy int, htmlAttrs string) templ.Component {

	if resource == nil || numCanFilterBy >= len(resource.CanFilterBy) {
		return StringInput("SubscriptionTopic.CanFilterBy."+strconv.Itoa(numCanFilterBy)+"..Description", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.CanFilterBy."+strconv.Itoa(numCanFilterBy)+"..Description", resource.CanFilterBy[numCanFilterBy].Description, htmlAttrs)
}
func (resource *SubscriptionTopic) T_CanFilterByResource(numCanFilterBy int, htmlAttrs string) templ.Component {

	if resource == nil || numCanFilterBy >= len(resource.CanFilterBy) {
		return StringInput("SubscriptionTopic.CanFilterBy."+strconv.Itoa(numCanFilterBy)+"..Resource", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.CanFilterBy."+strconv.Itoa(numCanFilterBy)+"..Resource", resource.CanFilterBy[numCanFilterBy].Resource, htmlAttrs)
}
func (resource *SubscriptionTopic) T_CanFilterByFilterParameter(numCanFilterBy int, htmlAttrs string) templ.Component {

	if resource == nil || numCanFilterBy >= len(resource.CanFilterBy) {
		return StringInput("SubscriptionTopic.CanFilterBy."+strconv.Itoa(numCanFilterBy)+"..FilterParameter", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.CanFilterBy."+strconv.Itoa(numCanFilterBy)+"..FilterParameter", &resource.CanFilterBy[numCanFilterBy].FilterParameter, htmlAttrs)
}
func (resource *SubscriptionTopic) T_CanFilterByFilterDefinition(numCanFilterBy int, htmlAttrs string) templ.Component {

	if resource == nil || numCanFilterBy >= len(resource.CanFilterBy) {
		return StringInput("SubscriptionTopic.CanFilterBy."+strconv.Itoa(numCanFilterBy)+"..FilterDefinition", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.CanFilterBy."+strconv.Itoa(numCanFilterBy)+"..FilterDefinition", resource.CanFilterBy[numCanFilterBy].FilterDefinition, htmlAttrs)
}
func (resource *SubscriptionTopic) T_CanFilterByModifier(numCanFilterBy int, numModifier int, htmlAttrs string) templ.Component {
	optionsValueSet := VSSubscription_search_modifier

	if resource == nil || numCanFilterBy >= len(resource.CanFilterBy) || numModifier >= len(resource.CanFilterBy[numCanFilterBy].Modifier) {
		return CodeSelect("SubscriptionTopic.CanFilterBy."+strconv.Itoa(numCanFilterBy)+"..Modifier."+strconv.Itoa(numModifier)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SubscriptionTopic.CanFilterBy."+strconv.Itoa(numCanFilterBy)+"..Modifier."+strconv.Itoa(numModifier)+".", &resource.CanFilterBy[numCanFilterBy].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionTopic) T_NotificationShapeResource(numNotificationShape int, htmlAttrs string) templ.Component {

	if resource == nil || numNotificationShape >= len(resource.NotificationShape) {
		return StringInput("SubscriptionTopic.NotificationShape."+strconv.Itoa(numNotificationShape)+"..Resource", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.NotificationShape."+strconv.Itoa(numNotificationShape)+"..Resource", &resource.NotificationShape[numNotificationShape].Resource, htmlAttrs)
}
func (resource *SubscriptionTopic) T_NotificationShapeInclude(numNotificationShape int, numInclude int, htmlAttrs string) templ.Component {

	if resource == nil || numNotificationShape >= len(resource.NotificationShape) || numInclude >= len(resource.NotificationShape[numNotificationShape].Include) {
		return StringInput("SubscriptionTopic.NotificationShape."+strconv.Itoa(numNotificationShape)+"..Include."+strconv.Itoa(numInclude)+".", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.NotificationShape."+strconv.Itoa(numNotificationShape)+"..Include."+strconv.Itoa(numInclude)+".", &resource.NotificationShape[numNotificationShape].Include[numInclude], htmlAttrs)
}
func (resource *SubscriptionTopic) T_NotificationShapeRevInclude(numNotificationShape int, numRevInclude int, htmlAttrs string) templ.Component {

	if resource == nil || numNotificationShape >= len(resource.NotificationShape) || numRevInclude >= len(resource.NotificationShape[numNotificationShape].RevInclude) {
		return StringInput("SubscriptionTopic.NotificationShape."+strconv.Itoa(numNotificationShape)+"..RevInclude."+strconv.Itoa(numRevInclude)+".", nil, htmlAttrs)
	}
	return StringInput("SubscriptionTopic.NotificationShape."+strconv.Itoa(numNotificationShape)+"..RevInclude."+strconv.Itoa(numRevInclude)+".", &resource.NotificationShape[numNotificationShape].RevInclude[numRevInclude], htmlAttrs)
}
