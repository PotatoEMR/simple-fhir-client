package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/SubscriptionStatus
type SubscriptionStatus struct {
	Id                           *string                               `json:"id,omitempty"`
	Meta                         *Meta                                 `json:"meta,omitempty"`
	ImplicitRules                *string                               `json:"implicitRules,omitempty"`
	Language                     *string                               `json:"language,omitempty"`
	Text                         *Narrative                            `json:"text,omitempty"`
	Contained                    []Resource                            `json:"contained,omitempty"`
	Extension                    []Extension                           `json:"extension,omitempty"`
	ModifierExtension            []Extension                           `json:"modifierExtension,omitempty"`
	Status                       *string                               `json:"status,omitempty"`
	Type                         string                                `json:"type"`
	EventsSinceSubscriptionStart *string                               `json:"eventsSinceSubscriptionStart,omitempty"`
	NotificationEvent            []SubscriptionStatusNotificationEvent `json:"notificationEvent,omitempty"`
	Subscription                 Reference                             `json:"subscription"`
	Topic                        *string                               `json:"topic,omitempty"`
	Error                        []CodeableConcept                     `json:"error,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SubscriptionStatus
type SubscriptionStatusNotificationEvent struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	EventNumber       string      `json:"eventNumber"`
	Timestamp         *string     `json:"timestamp,omitempty"`
	Focus             *Reference  `json:"focus,omitempty"`
	AdditionalContext []Reference `json:"additionalContext,omitempty"`
}

type OtherSubscriptionStatus SubscriptionStatus

// on convert struct to json, automatically add resourceType=SubscriptionStatus
func (r SubscriptionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubscriptionStatus
		ResourceType string `json:"resourceType"`
	}{
		OtherSubscriptionStatus: OtherSubscriptionStatus(r),
		ResourceType:            "SubscriptionStatus",
	})
}

func (resource *SubscriptionStatus) T_Id() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionStatus.Id", nil)
	}
	return StringInput("SubscriptionStatus.Id", resource.Id)
}
func (resource *SubscriptionStatus) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionStatus.ImplicitRules", nil)
	}
	return StringInput("SubscriptionStatus.ImplicitRules", resource.ImplicitRules)
}
func (resource *SubscriptionStatus) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("SubscriptionStatus.Language", nil, optionsValueSet)
	}
	return CodeSelect("SubscriptionStatus.Language", resource.Language, optionsValueSet)
}
func (resource *SubscriptionStatus) T_Status() templ.Component {
	optionsValueSet := VSSubscription_status

	if resource == nil {
		return CodeSelect("SubscriptionStatus.Status", nil, optionsValueSet)
	}
	return CodeSelect("SubscriptionStatus.Status", resource.Status, optionsValueSet)
}
func (resource *SubscriptionStatus) T_Type() templ.Component {
	optionsValueSet := VSSubscription_notification_type

	if resource == nil {
		return CodeSelect("SubscriptionStatus.Type", nil, optionsValueSet)
	}
	return CodeSelect("SubscriptionStatus.Type", &resource.Type, optionsValueSet)
}
func (resource *SubscriptionStatus) T_EventsSinceSubscriptionStart() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionStatus.EventsSinceSubscriptionStart", nil)
	}
	return StringInput("SubscriptionStatus.EventsSinceSubscriptionStart", resource.EventsSinceSubscriptionStart)
}
func (resource *SubscriptionStatus) T_Topic() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionStatus.Topic", nil)
	}
	return StringInput("SubscriptionStatus.Topic", resource.Topic)
}
func (resource *SubscriptionStatus) T_Error(numError int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Error) >= numError {
		return CodeableConceptSelect("SubscriptionStatus.Error["+strconv.Itoa(numError)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubscriptionStatus.Error["+strconv.Itoa(numError)+"]", &resource.Error[numError], optionsValueSet)
}
func (resource *SubscriptionStatus) T_NotificationEventId(numNotificationEvent int) templ.Component {

	if resource == nil || len(resource.NotificationEvent) >= numNotificationEvent {
		return StringInput("SubscriptionStatus.NotificationEvent["+strconv.Itoa(numNotificationEvent)+"].Id", nil)
	}
	return StringInput("SubscriptionStatus.NotificationEvent["+strconv.Itoa(numNotificationEvent)+"].Id", resource.NotificationEvent[numNotificationEvent].Id)
}
func (resource *SubscriptionStatus) T_NotificationEventEventNumber(numNotificationEvent int) templ.Component {

	if resource == nil || len(resource.NotificationEvent) >= numNotificationEvent {
		return StringInput("SubscriptionStatus.NotificationEvent["+strconv.Itoa(numNotificationEvent)+"].EventNumber", nil)
	}
	return StringInput("SubscriptionStatus.NotificationEvent["+strconv.Itoa(numNotificationEvent)+"].EventNumber", &resource.NotificationEvent[numNotificationEvent].EventNumber)
}
func (resource *SubscriptionStatus) T_NotificationEventTimestamp(numNotificationEvent int) templ.Component {

	if resource == nil || len(resource.NotificationEvent) >= numNotificationEvent {
		return StringInput("SubscriptionStatus.NotificationEvent["+strconv.Itoa(numNotificationEvent)+"].Timestamp", nil)
	}
	return StringInput("SubscriptionStatus.NotificationEvent["+strconv.Itoa(numNotificationEvent)+"].Timestamp", resource.NotificationEvent[numNotificationEvent].Timestamp)
}
