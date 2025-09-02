package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	Date              *string                              `json:"date,omitempty"`
	Publisher         *string                              `json:"publisher,omitempty"`
	Contact           []ContactDetail                      `json:"contact,omitempty"`
	Description       *string                              `json:"description,omitempty"`
	UseContext        []UsageContext                       `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                    `json:"jurisdiction,omitempty"`
	Purpose           *string                              `json:"purpose,omitempty"`
	Copyright         *string                              `json:"copyright,omitempty"`
	ApprovalDate      *string                              `json:"approvalDate,omitempty"`
	LastReviewDate    *string                              `json:"lastReviewDate,omitempty"`
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

func (resource *SubscriptionTopic) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *SubscriptionTopic) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *SubscriptionTopic) T_Jurisdiction(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *SubscriptionTopic) T_ResourceTriggerSupportedInteraction(numResourceTrigger int) templ.Component {
	optionsValueSet := VSInteraction_trigger

	if resource == nil && len(resource.ResourceTrigger) >= numResourceTrigger {
		return CodeSelect("supportedInteraction", nil, optionsValueSet)
	}
	return CodeSelect("supportedInteraction", &resource.ResourceTrigger[numResourceTrigger].SupportedInteraction[0], optionsValueSet)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaResultForCreate(numResourceTrigger int) templ.Component {
	optionsValueSet := VSSubscriptiontopic_cr_behavior

	if resource == nil && len(resource.ResourceTrigger) >= numResourceTrigger {
		return CodeSelect("resultForCreate", nil, optionsValueSet)
	}
	return CodeSelect("resultForCreate", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.ResultForCreate, optionsValueSet)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaResultForDelete(numResourceTrigger int) templ.Component {
	optionsValueSet := VSSubscriptiontopic_cr_behavior

	if resource == nil && len(resource.ResourceTrigger) >= numResourceTrigger {
		return CodeSelect("resultForDelete", nil, optionsValueSet)
	}
	return CodeSelect("resultForDelete", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.ResultForDelete, optionsValueSet)
}
func (resource *SubscriptionTopic) T_EventTriggerEvent(numEventTrigger int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.EventTrigger) >= numEventTrigger {
		return CodeableConceptSelect("event", nil, optionsValueSet)
	}
	return CodeableConceptSelect("event", &resource.EventTrigger[numEventTrigger].Event, optionsValueSet)
}
func (resource *SubscriptionTopic) T_CanFilterByModifier(numCanFilterBy int) templ.Component {
	optionsValueSet := VSSubscription_search_modifier

	if resource == nil && len(resource.CanFilterBy) >= numCanFilterBy {
		return CodeSelect("modifier", nil, optionsValueSet)
	}
	return CodeSelect("modifier", &resource.CanFilterBy[numCanFilterBy].Modifier[0], optionsValueSet)
}
