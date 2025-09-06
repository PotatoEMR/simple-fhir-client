package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Subscription
type Subscription struct {
	Id                *string             `json:"id,omitempty"`
	Meta              *Meta               `json:"meta,omitempty"`
	ImplicitRules     *string             `json:"implicitRules,omitempty"`
	Language          *string             `json:"language,omitempty"`
	Text              *Narrative          `json:"text,omitempty"`
	Contained         []Resource          `json:"contained,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Status            string              `json:"status"`
	Contact           []ContactPoint      `json:"contact,omitempty"`
	End               *string             `json:"end,omitempty"`
	Reason            string              `json:"reason"`
	Criteria          string              `json:"criteria"`
	Error             *string             `json:"error,omitempty"`
	Channel           SubscriptionChannel `json:"channel"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Subscription
type SubscriptionChannel struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Endpoint          *string     `json:"endpoint,omitempty"`
	Payload           *string     `json:"payload,omitempty"`
	Header            []string    `json:"header,omitempty"`
}

type OtherSubscription Subscription

// on convert struct to json, automatically add resourceType=Subscription
func (r Subscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubscription
		ResourceType string `json:"resourceType"`
	}{
		OtherSubscription: OtherSubscription(r),
		ResourceType:      "Subscription",
	})
}

func (resource *Subscription) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Subscription.Id", nil)
	}
	return StringInput("Subscription.Id", resource.Id)
}
func (resource *Subscription) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Subscription.ImplicitRules", nil)
	}
	return StringInput("Subscription.ImplicitRules", resource.ImplicitRules)
}
func (resource *Subscription) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Subscription.Language", nil, optionsValueSet)
	}
	return CodeSelect("Subscription.Language", resource.Language, optionsValueSet)
}
func (resource *Subscription) T_Status() templ.Component {
	optionsValueSet := VSSubscription_status

	if resource == nil {
		return CodeSelect("Subscription.Status", nil, optionsValueSet)
	}
	return CodeSelect("Subscription.Status", &resource.Status, optionsValueSet)
}
func (resource *Subscription) T_End() templ.Component {

	if resource == nil {
		return StringInput("Subscription.End", nil)
	}
	return StringInput("Subscription.End", resource.End)
}
func (resource *Subscription) T_Reason() templ.Component {

	if resource == nil {
		return StringInput("Subscription.Reason", nil)
	}
	return StringInput("Subscription.Reason", &resource.Reason)
}
func (resource *Subscription) T_Criteria() templ.Component {

	if resource == nil {
		return StringInput("Subscription.Criteria", nil)
	}
	return StringInput("Subscription.Criteria", &resource.Criteria)
}
func (resource *Subscription) T_Error() templ.Component {

	if resource == nil {
		return StringInput("Subscription.Error", nil)
	}
	return StringInput("Subscription.Error", resource.Error)
}
func (resource *Subscription) T_ChannelId() templ.Component {

	if resource == nil {
		return StringInput("Subscription.Channel.Id", nil)
	}
	return StringInput("Subscription.Channel.Id", resource.Channel.Id)
}
func (resource *Subscription) T_ChannelType() templ.Component {
	optionsValueSet := VSSubscription_channel_type

	if resource == nil {
		return CodeSelect("Subscription.Channel.Type", nil, optionsValueSet)
	}
	return CodeSelect("Subscription.Channel.Type", &resource.Channel.Type, optionsValueSet)
}
func (resource *Subscription) T_ChannelEndpoint() templ.Component {

	if resource == nil {
		return StringInput("Subscription.Channel.Endpoint", nil)
	}
	return StringInput("Subscription.Channel.Endpoint", resource.Channel.Endpoint)
}
func (resource *Subscription) T_ChannelPayload(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Subscription.Channel.Payload", nil, optionsValueSet)
	}
	return CodeSelect("Subscription.Channel.Payload", resource.Channel.Payload, optionsValueSet)
}
func (resource *Subscription) T_ChannelHeader(numHeader int) templ.Component {

	if resource == nil || len(resource.Channel.Header) >= numHeader {
		return StringInput("Subscription.Channel.Header["+strconv.Itoa(numHeader)+"]", nil)
	}
	return StringInput("Subscription.Channel.Header["+strconv.Itoa(numHeader)+"]", &resource.Channel.Header[numHeader])
}
