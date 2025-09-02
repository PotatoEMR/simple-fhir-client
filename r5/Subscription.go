package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *Subscription) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Subscription) T_Status() templ.Component {
	optionsValueSet := VSSubscription_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Subscription) T_ChannelType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("channelType", nil, optionsValueSet)
	}
	return CodingSelect("channelType", &resource.ChannelType, optionsValueSet)
}
func (resource *Subscription) T_ContentType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("contentType", nil, optionsValueSet)
	}
	return CodeSelect("contentType", resource.ContentType, optionsValueSet)
}
func (resource *Subscription) T_Content() templ.Component {
	optionsValueSet := VSSubscription_payload_content

	if resource == nil {
		return CodeSelect("content", nil, optionsValueSet)
	}
	return CodeSelect("content", resource.Content, optionsValueSet)
}
func (resource *Subscription) T_FilterByComparator(numFilterBy int) templ.Component {
	optionsValueSet := VSSearch_comparator

	if resource == nil && len(resource.FilterBy) >= numFilterBy {
		return CodeSelect("comparator", nil, optionsValueSet)
	}
	return CodeSelect("comparator", resource.FilterBy[numFilterBy].Comparator, optionsValueSet)
}
func (resource *Subscription) T_FilterByModifier(numFilterBy int) templ.Component {
	optionsValueSet := VSSearch_modifier_code

	if resource == nil && len(resource.FilterBy) >= numFilterBy {
		return CodeSelect("modifier", nil, optionsValueSet)
	}
	return CodeSelect("modifier", resource.FilterBy[numFilterBy].Modifier, optionsValueSet)
}
