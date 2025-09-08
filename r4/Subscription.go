package r4

//generated with command go run ./bultaoreune
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
func (r Subscription) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Subscription/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "Subscription"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Subscription) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSSubscription_status

	if resource == nil {
		return CodeSelect("Subscription.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Subscription.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Subscription) T_End(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Subscription.End", nil, htmlAttrs)
	}
	return StringInput("Subscription.End", resource.End, htmlAttrs)
}
func (resource *Subscription) T_Reason(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Subscription.Reason", nil, htmlAttrs)
	}
	return StringInput("Subscription.Reason", &resource.Reason, htmlAttrs)
}
func (resource *Subscription) T_Criteria(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Subscription.Criteria", nil, htmlAttrs)
	}
	return StringInput("Subscription.Criteria", &resource.Criteria, htmlAttrs)
}
func (resource *Subscription) T_Error(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Subscription.Error", nil, htmlAttrs)
	}
	return StringInput("Subscription.Error", resource.Error, htmlAttrs)
}
func (resource *Subscription) T_ChannelType(htmlAttrs string) templ.Component {
	optionsValueSet := VSSubscription_channel_type

	if resource == nil {
		return CodeSelect("Subscription.Channel.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Subscription.Channel.Type", &resource.Channel.Type, optionsValueSet, htmlAttrs)
}
func (resource *Subscription) T_ChannelEndpoint(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Subscription.Channel.Endpoint", nil, htmlAttrs)
	}
	return StringInput("Subscription.Channel.Endpoint", resource.Channel.Endpoint, htmlAttrs)
}
func (resource *Subscription) T_ChannelPayload(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeSelect("Subscription.Channel.Payload", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Subscription.Channel.Payload", resource.Channel.Payload, optionsValueSet, htmlAttrs)
}
func (resource *Subscription) T_ChannelHeader(numHeader int, htmlAttrs string) templ.Component {

	if resource == nil || numHeader >= len(resource.Channel.Header) {
		return StringInput("Subscription.Channel.Header."+strconv.Itoa(numHeader)+".", nil, htmlAttrs)
	}
	return StringInput("Subscription.Channel.Header."+strconv.Itoa(numHeader)+".", &resource.Channel.Header[numHeader], htmlAttrs)
}
