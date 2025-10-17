package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
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

// struct -> json, automatically add resourceType=Patient
func (r Subscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubscription
		ResourceType string `json:"resourceType"`
	}{
		OtherSubscription: OtherSubscription(r),
		ResourceType:      "Subscription",
	})
}

// json -> struct, first reject if resourceType != Subscription
func (r *Subscription) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "Subscription" {
		return errors.New("resourceType not Subscription")
	}
	return json.Unmarshal(data, (*OtherSubscription)(r))
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
func (resource *Subscription) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSubscription_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Subscription) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactPointInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactPointInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *Subscription) T_End(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("end", nil, htmlAttrs)
	}
	return StringInput("end", resource.End, htmlAttrs)
}
func (resource *Subscription) T_Reason(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("reason", nil, htmlAttrs)
	}
	return StringInput("reason", &resource.Reason, htmlAttrs)
}
func (resource *Subscription) T_Criteria(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("criteria", nil, htmlAttrs)
	}
	return StringInput("criteria", &resource.Criteria, htmlAttrs)
}
func (resource *Subscription) T_Error(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("error", nil, htmlAttrs)
	}
	return StringInput("error", resource.Error, htmlAttrs)
}
func (resource *Subscription) T_ChannelType(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSubscription_channel_type

	if resource == nil {
		return CodeSelect("channel.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("channel.type", &resource.Channel.Type, optionsValueSet, htmlAttrs)
}
func (resource *Subscription) T_ChannelEndpoint(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("channel.endpoint", nil, htmlAttrs)
	}
	return StringInput("channel.endpoint", resource.Channel.Endpoint, htmlAttrs)
}
func (resource *Subscription) T_ChannelPayload(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeSelect("channel.payload", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("channel.payload", resource.Channel.Payload, optionsValueSet, htmlAttrs)
}
func (resource *Subscription) T_ChannelHeader(numHeader int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numHeader >= len(resource.Channel.Header) {
		return StringInput("channel.header["+strconv.Itoa(numHeader)+"]", nil, htmlAttrs)
	}
	return StringInput("channel.header["+strconv.Itoa(numHeader)+"]", &resource.Channel.Header[numHeader], htmlAttrs)
}
