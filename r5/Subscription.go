package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Subscription
type Subscription struct {
	Id                *string                 `json:"id,omitempty"`
	Meta              *Meta                   `json:"meta,omitempty"`
	ImplicitRules     *string                 `json:"implicitRules,omitempty"`
	Language          *string                 `json:"language,omitempty"`
	Text              *Narrative              `json:"text,omitempty"`
	Contained         []Resource              `json:"contained,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Identifier        []Identifier            `json:"identifier,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Status            string                  `json:"status"`
	Topic             string                  `json:"topic"`
	Contact           []ContactPoint          `json:"contact,omitempty"`
	End               *string                 `json:"end,omitempty"`
	ManagingEntity    *Reference              `json:"managingEntity,omitempty"`
	Reason            *string                 `json:"reason,omitempty"`
	FilterBy          []SubscriptionFilterBy  `json:"filterBy,omitempty"`
	ChannelType       Coding                  `json:"channelType"`
	Endpoint          *string                 `json:"endpoint,omitempty"`
	Parameter         []SubscriptionParameter `json:"parameter,omitempty"`
	HeartbeatPeriod   *int                    `json:"heartbeatPeriod,omitempty"`
	Timeout           *int                    `json:"timeout,omitempty"`
	ContentType       *string                 `json:"contentType,omitempty"`
	Content           *string                 `json:"content,omitempty"`
	MaxCount          *int                    `json:"maxCount,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Subscription
type SubscriptionFilterBy struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ResourceType      *string     `json:"resourceType,omitempty"`
	FilterParameter   string      `json:"filterParameter"`
	Comparator        *string     `json:"comparator,omitempty"`
	Modifier          *string     `json:"modifier,omitempty"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Subscription
type SubscriptionParameter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Value             string      `json:"value"`
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
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Subscription"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Subscription) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *Subscription) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSSubscription_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Subscription) T_Topic(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("topic", nil, htmlAttrs)
	}
	return StringInput("topic", &resource.Topic, htmlAttrs)
}
func (resource *Subscription) T_End(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("end", nil, htmlAttrs)
	}
	return StringInput("end", resource.End, htmlAttrs)
}
func (resource *Subscription) T_Reason(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("reason", nil, htmlAttrs)
	}
	return StringInput("reason", resource.Reason, htmlAttrs)
}
func (resource *Subscription) T_ChannelType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("channelType", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("channelType", &resource.ChannelType, optionsValueSet, htmlAttrs)
}
func (resource *Subscription) T_Endpoint(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("endpoint", nil, htmlAttrs)
	}
	return StringInput("endpoint", resource.Endpoint, htmlAttrs)
}
func (resource *Subscription) T_HeartbeatPeriod(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("heartbeatPeriod", nil, htmlAttrs)
	}
	return IntInput("heartbeatPeriod", resource.HeartbeatPeriod, htmlAttrs)
}
func (resource *Subscription) T_Timeout(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("timeout", nil, htmlAttrs)
	}
	return IntInput("timeout", resource.Timeout, htmlAttrs)
}
func (resource *Subscription) T_ContentType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeSelect("contentType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("contentType", resource.ContentType, optionsValueSet, htmlAttrs)
}
func (resource *Subscription) T_Content(htmlAttrs string) templ.Component {
	optionsValueSet := VSSubscription_payload_content

	if resource == nil {
		return CodeSelect("content", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("content", resource.Content, optionsValueSet, htmlAttrs)
}
func (resource *Subscription) T_MaxCount(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("maxCount", nil, htmlAttrs)
	}
	return IntInput("maxCount", resource.MaxCount, htmlAttrs)
}
func (resource *Subscription) T_FilterByResourceType(numFilterBy int, htmlAttrs string) templ.Component {
	if resource == nil || numFilterBy >= len(resource.FilterBy) {
		return StringInput("filterBy["+strconv.Itoa(numFilterBy)+"].resourceType", nil, htmlAttrs)
	}
	return StringInput("filterBy["+strconv.Itoa(numFilterBy)+"].resourceType", resource.FilterBy[numFilterBy].ResourceType, htmlAttrs)
}
func (resource *Subscription) T_FilterByFilterParameter(numFilterBy int, htmlAttrs string) templ.Component {
	if resource == nil || numFilterBy >= len(resource.FilterBy) {
		return StringInput("filterBy["+strconv.Itoa(numFilterBy)+"].filterParameter", nil, htmlAttrs)
	}
	return StringInput("filterBy["+strconv.Itoa(numFilterBy)+"].filterParameter", &resource.FilterBy[numFilterBy].FilterParameter, htmlAttrs)
}
func (resource *Subscription) T_FilterByComparator(numFilterBy int, htmlAttrs string) templ.Component {
	optionsValueSet := VSSearch_comparator

	if resource == nil || numFilterBy >= len(resource.FilterBy) {
		return CodeSelect("filterBy["+strconv.Itoa(numFilterBy)+"].comparator", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("filterBy["+strconv.Itoa(numFilterBy)+"].comparator", resource.FilterBy[numFilterBy].Comparator, optionsValueSet, htmlAttrs)
}
func (resource *Subscription) T_FilterByModifier(numFilterBy int, htmlAttrs string) templ.Component {
	optionsValueSet := VSSearch_modifier_code

	if resource == nil || numFilterBy >= len(resource.FilterBy) {
		return CodeSelect("filterBy["+strconv.Itoa(numFilterBy)+"].modifier", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("filterBy["+strconv.Itoa(numFilterBy)+"].modifier", resource.FilterBy[numFilterBy].Modifier, optionsValueSet, htmlAttrs)
}
func (resource *Subscription) T_FilterByValue(numFilterBy int, htmlAttrs string) templ.Component {
	if resource == nil || numFilterBy >= len(resource.FilterBy) {
		return StringInput("filterBy["+strconv.Itoa(numFilterBy)+"].value", nil, htmlAttrs)
	}
	return StringInput("filterBy["+strconv.Itoa(numFilterBy)+"].value", &resource.FilterBy[numFilterBy].Value, htmlAttrs)
}
func (resource *Subscription) T_ParameterName(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].name", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].name", &resource.Parameter[numParameter].Name, htmlAttrs)
}
func (resource *Subscription) T_ParameterValue(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].value", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].value", &resource.Parameter[numParameter].Value, htmlAttrs)
}
