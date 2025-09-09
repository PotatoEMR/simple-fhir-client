package r4b

//generated with command go run ./bultaoreune
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
func (r SubscriptionStatus) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "SubscriptionStatus/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "SubscriptionStatus"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *SubscriptionStatus) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSSubscription_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionStatus) T_Type(htmlAttrs string) templ.Component {
	optionsValueSet := VSSubscription_notification_type

	if resource == nil {
		return CodeSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionStatus) T_EventsSinceSubscriptionStart(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("eventsSinceSubscriptionStart", nil, htmlAttrs)
	}
	return StringInput("eventsSinceSubscriptionStart", resource.EventsSinceSubscriptionStart, htmlAttrs)
}
func (resource *SubscriptionStatus) T_Topic(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("topic", nil, htmlAttrs)
	}
	return StringInput("topic", resource.Topic, htmlAttrs)
}
func (resource *SubscriptionStatus) T_Error(numError int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numError >= len(resource.Error) {
		return CodeableConceptSelect("error["+strconv.Itoa(numError)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("error["+strconv.Itoa(numError)+"]", &resource.Error[numError], optionsValueSet, htmlAttrs)
}
func (resource *SubscriptionStatus) T_NotificationEventEventNumber(numNotificationEvent int, htmlAttrs string) templ.Component {
	if resource == nil || numNotificationEvent >= len(resource.NotificationEvent) {
		return StringInput("notificationEvent["+strconv.Itoa(numNotificationEvent)+"].eventNumber", nil, htmlAttrs)
	}
	return StringInput("notificationEvent["+strconv.Itoa(numNotificationEvent)+"].eventNumber", &resource.NotificationEvent[numNotificationEvent].EventNumber, htmlAttrs)
}
func (resource *SubscriptionStatus) T_NotificationEventTimestamp(numNotificationEvent int, htmlAttrs string) templ.Component {
	if resource == nil || numNotificationEvent >= len(resource.NotificationEvent) {
		return StringInput("notificationEvent["+strconv.Itoa(numNotificationEvent)+"].timestamp", nil, htmlAttrs)
	}
	return StringInput("notificationEvent["+strconv.Itoa(numNotificationEvent)+"].timestamp", resource.NotificationEvent[numNotificationEvent].Timestamp, htmlAttrs)
}
