package r5

//generated with command go run ./bultaoreune -nodownload
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
func (resource *Subscription) T_Name() templ.Component {

	if resource == nil {
		return StringInput("Subscription.Name", nil)
	}
	return StringInput("Subscription.Name", resource.Name)
}
func (resource *Subscription) T_Status() templ.Component {
	optionsValueSet := VSSubscription_status

	if resource == nil {
		return CodeSelect("Subscription.Status", nil, optionsValueSet)
	}
	return CodeSelect("Subscription.Status", &resource.Status, optionsValueSet)
}
func (resource *Subscription) T_Topic() templ.Component {

	if resource == nil {
		return StringInput("Subscription.Topic", nil)
	}
	return StringInput("Subscription.Topic", &resource.Topic)
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
	return StringInput("Subscription.Reason", resource.Reason)
}
func (resource *Subscription) T_ChannelType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("Subscription.ChannelType", nil, optionsValueSet)
	}
	return CodingSelect("Subscription.ChannelType", &resource.ChannelType, optionsValueSet)
}
func (resource *Subscription) T_Endpoint() templ.Component {

	if resource == nil {
		return StringInput("Subscription.Endpoint", nil)
	}
	return StringInput("Subscription.Endpoint", resource.Endpoint)
}
func (resource *Subscription) T_HeartbeatPeriod() templ.Component {

	if resource == nil {
		return IntInput("Subscription.HeartbeatPeriod", nil)
	}
	return IntInput("Subscription.HeartbeatPeriod", resource.HeartbeatPeriod)
}
func (resource *Subscription) T_Timeout() templ.Component {

	if resource == nil {
		return IntInput("Subscription.Timeout", nil)
	}
	return IntInput("Subscription.Timeout", resource.Timeout)
}
func (resource *Subscription) T_ContentType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Subscription.ContentType", nil, optionsValueSet)
	}
	return CodeSelect("Subscription.ContentType", resource.ContentType, optionsValueSet)
}
func (resource *Subscription) T_Content() templ.Component {
	optionsValueSet := VSSubscription_payload_content

	if resource == nil {
		return CodeSelect("Subscription.Content", nil, optionsValueSet)
	}
	return CodeSelect("Subscription.Content", resource.Content, optionsValueSet)
}
func (resource *Subscription) T_MaxCount() templ.Component {

	if resource == nil {
		return IntInput("Subscription.MaxCount", nil)
	}
	return IntInput("Subscription.MaxCount", resource.MaxCount)
}
func (resource *Subscription) T_FilterById(numFilterBy int) templ.Component {

	if resource == nil || len(resource.FilterBy) >= numFilterBy {
		return StringInput("Subscription.FilterBy["+strconv.Itoa(numFilterBy)+"].Id", nil)
	}
	return StringInput("Subscription.FilterBy["+strconv.Itoa(numFilterBy)+"].Id", resource.FilterBy[numFilterBy].Id)
}
func (resource *Subscription) T_FilterByResourceType(numFilterBy int) templ.Component {

	if resource == nil || len(resource.FilterBy) >= numFilterBy {
		return StringInput("Subscription.FilterBy["+strconv.Itoa(numFilterBy)+"].ResourceType", nil)
	}
	return StringInput("Subscription.FilterBy["+strconv.Itoa(numFilterBy)+"].ResourceType", resource.FilterBy[numFilterBy].ResourceType)
}
func (resource *Subscription) T_FilterByFilterParameter(numFilterBy int) templ.Component {

	if resource == nil || len(resource.FilterBy) >= numFilterBy {
		return StringInput("Subscription.FilterBy["+strconv.Itoa(numFilterBy)+"].FilterParameter", nil)
	}
	return StringInput("Subscription.FilterBy["+strconv.Itoa(numFilterBy)+"].FilterParameter", &resource.FilterBy[numFilterBy].FilterParameter)
}
func (resource *Subscription) T_FilterByComparator(numFilterBy int) templ.Component {
	optionsValueSet := VSSearch_comparator

	if resource == nil || len(resource.FilterBy) >= numFilterBy {
		return CodeSelect("Subscription.FilterBy["+strconv.Itoa(numFilterBy)+"].Comparator", nil, optionsValueSet)
	}
	return CodeSelect("Subscription.FilterBy["+strconv.Itoa(numFilterBy)+"].Comparator", resource.FilterBy[numFilterBy].Comparator, optionsValueSet)
}
func (resource *Subscription) T_FilterByModifier(numFilterBy int) templ.Component {
	optionsValueSet := VSSearch_modifier_code

	if resource == nil || len(resource.FilterBy) >= numFilterBy {
		return CodeSelect("Subscription.FilterBy["+strconv.Itoa(numFilterBy)+"].Modifier", nil, optionsValueSet)
	}
	return CodeSelect("Subscription.FilterBy["+strconv.Itoa(numFilterBy)+"].Modifier", resource.FilterBy[numFilterBy].Modifier, optionsValueSet)
}
func (resource *Subscription) T_FilterByValue(numFilterBy int) templ.Component {

	if resource == nil || len(resource.FilterBy) >= numFilterBy {
		return StringInput("Subscription.FilterBy["+strconv.Itoa(numFilterBy)+"].Value", nil)
	}
	return StringInput("Subscription.FilterBy["+strconv.Itoa(numFilterBy)+"].Value", &resource.FilterBy[numFilterBy].Value)
}
func (resource *Subscription) T_ParameterId(numParameter int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return StringInput("Subscription.Parameter["+strconv.Itoa(numParameter)+"].Id", nil)
	}
	return StringInput("Subscription.Parameter["+strconv.Itoa(numParameter)+"].Id", resource.Parameter[numParameter].Id)
}
func (resource *Subscription) T_ParameterName(numParameter int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return StringInput("Subscription.Parameter["+strconv.Itoa(numParameter)+"].Name", nil)
	}
	return StringInput("Subscription.Parameter["+strconv.Itoa(numParameter)+"].Name", &resource.Parameter[numParameter].Name)
}
func (resource *Subscription) T_ParameterValue(numParameter int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return StringInput("Subscription.Parameter["+strconv.Itoa(numParameter)+"].Value", nil)
	}
	return StringInput("Subscription.Parameter["+strconv.Itoa(numParameter)+"].Value", &resource.Parameter[numParameter].Value)
}
