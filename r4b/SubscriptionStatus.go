package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *SubscriptionStatus) SubscriptionStatusLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *SubscriptionStatus) SubscriptionStatusStatus() templ.Component {
	optionsValueSet := VSSubscription_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", resource.Status, optionsValueSet)
}
func (resource *SubscriptionStatus) SubscriptionStatusType() templ.Component {
	optionsValueSet := VSSubscription_notification_type

	if resource == nil {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", &resource.Type, optionsValueSet)
}
func (resource *SubscriptionStatus) SubscriptionStatusError(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("error", nil, optionsValueSet)
	}
	return CodeableConceptSelect("error", &resource.Error[0], optionsValueSet)
}
