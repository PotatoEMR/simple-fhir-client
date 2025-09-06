package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/SubscriptionTopic
type SubscriptionTopic struct {
	Id                     *string                              `json:"id,omitempty"`
	Meta                   *Meta                                `json:"meta,omitempty"`
	ImplicitRules          *string                              `json:"implicitRules,omitempty"`
	Language               *string                              `json:"language,omitempty"`
	Text                   *Narrative                           `json:"text,omitempty"`
	Contained              []Resource                           `json:"contained,omitempty"`
	Extension              []Extension                          `json:"extension,omitempty"`
	ModifierExtension      []Extension                          `json:"modifierExtension,omitempty"`
	Url                    string                               `json:"url"`
	Identifier             []Identifier                         `json:"identifier,omitempty"`
	Version                *string                              `json:"version,omitempty"`
	VersionAlgorithmString *string                              `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                              `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                              `json:"name,omitempty"`
	Title                  *string                              `json:"title,omitempty"`
	DerivedFrom            []string                             `json:"derivedFrom,omitempty"`
	Status                 string                               `json:"status"`
	Experimental           *bool                                `json:"experimental,omitempty"`
	Date                   *string                              `json:"date,omitempty"`
	Publisher              *string                              `json:"publisher,omitempty"`
	Contact                []ContactDetail                      `json:"contact,omitempty"`
	Description            *string                              `json:"description,omitempty"`
	UseContext             []UsageContext                       `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                    `json:"jurisdiction,omitempty"`
	Purpose                *string                              `json:"purpose,omitempty"`
	Copyright              *string                              `json:"copyright,omitempty"`
	CopyrightLabel         *string                              `json:"copyrightLabel,omitempty"`
	ApprovalDate           *string                              `json:"approvalDate,omitempty"`
	LastReviewDate         *string                              `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                              `json:"effectivePeriod,omitempty"`
	ResourceTrigger        []SubscriptionTopicResourceTrigger   `json:"resourceTrigger,omitempty"`
	EventTrigger           []SubscriptionTopicEventTrigger      `json:"eventTrigger,omitempty"`
	CanFilterBy            []SubscriptionTopicCanFilterBy       `json:"canFilterBy,omitempty"`
	NotificationShape      []SubscriptionTopicNotificationShape `json:"notificationShape,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubscriptionTopic
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

// http://hl7.org/fhir/r5/StructureDefinition/SubscriptionTopic
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

// http://hl7.org/fhir/r5/StructureDefinition/SubscriptionTopic
type SubscriptionTopicEventTrigger struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Description       *string         `json:"description,omitempty"`
	Event             CodeableConcept `json:"event"`
	Resource          string          `json:"resource"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubscriptionTopic
type SubscriptionTopicCanFilterBy struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Resource          *string     `json:"resource,omitempty"`
	FilterParameter   string      `json:"filterParameter"`
	FilterDefinition  *string     `json:"filterDefinition,omitempty"`
	Comparator        []string    `json:"comparator,omitempty"`
	Modifier          []string    `json:"modifier,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubscriptionTopic
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

func (resource *SubscriptionTopic) T_Id() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Id", nil)
	}
	return StringInput("SubscriptionTopic.Id", resource.Id)
}
func (resource *SubscriptionTopic) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.ImplicitRules", nil)
	}
	return StringInput("SubscriptionTopic.ImplicitRules", resource.ImplicitRules)
}
func (resource *SubscriptionTopic) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("SubscriptionTopic.Language", nil, optionsValueSet)
	}
	return CodeSelect("SubscriptionTopic.Language", resource.Language, optionsValueSet)
}
func (resource *SubscriptionTopic) T_Url() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Url", nil)
	}
	return StringInput("SubscriptionTopic.Url", &resource.Url)
}
func (resource *SubscriptionTopic) T_Version() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Version", nil)
	}
	return StringInput("SubscriptionTopic.Version", resource.Version)
}
func (resource *SubscriptionTopic) T_Name() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Name", nil)
	}
	return StringInput("SubscriptionTopic.Name", resource.Name)
}
func (resource *SubscriptionTopic) T_Title() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Title", nil)
	}
	return StringInput("SubscriptionTopic.Title", resource.Title)
}
func (resource *SubscriptionTopic) T_DerivedFrom(numDerivedFrom int) templ.Component {

	if resource == nil || len(resource.DerivedFrom) >= numDerivedFrom {
		return StringInput("SubscriptionTopic.DerivedFrom["+strconv.Itoa(numDerivedFrom)+"]", nil)
	}
	return StringInput("SubscriptionTopic.DerivedFrom["+strconv.Itoa(numDerivedFrom)+"]", &resource.DerivedFrom[numDerivedFrom])
}
func (resource *SubscriptionTopic) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("SubscriptionTopic.Status", nil, optionsValueSet)
	}
	return CodeSelect("SubscriptionTopic.Status", &resource.Status, optionsValueSet)
}
func (resource *SubscriptionTopic) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("SubscriptionTopic.Experimental", nil)
	}
	return BoolInput("SubscriptionTopic.Experimental", resource.Experimental)
}
func (resource *SubscriptionTopic) T_Date() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Date", nil)
	}
	return StringInput("SubscriptionTopic.Date", resource.Date)
}
func (resource *SubscriptionTopic) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Publisher", nil)
	}
	return StringInput("SubscriptionTopic.Publisher", resource.Publisher)
}
func (resource *SubscriptionTopic) T_Description() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Description", nil)
	}
	return StringInput("SubscriptionTopic.Description", resource.Description)
}
func (resource *SubscriptionTopic) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("SubscriptionTopic.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubscriptionTopic.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *SubscriptionTopic) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Purpose", nil)
	}
	return StringInput("SubscriptionTopic.Purpose", resource.Purpose)
}
func (resource *SubscriptionTopic) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.Copyright", nil)
	}
	return StringInput("SubscriptionTopic.Copyright", resource.Copyright)
}
func (resource *SubscriptionTopic) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.CopyrightLabel", nil)
	}
	return StringInput("SubscriptionTopic.CopyrightLabel", resource.CopyrightLabel)
}
func (resource *SubscriptionTopic) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.ApprovalDate", nil)
	}
	return StringInput("SubscriptionTopic.ApprovalDate", resource.ApprovalDate)
}
func (resource *SubscriptionTopic) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("SubscriptionTopic.LastReviewDate", nil)
	}
	return StringInput("SubscriptionTopic.LastReviewDate", resource.LastReviewDate)
}
func (resource *SubscriptionTopic) T_ResourceTriggerId(numResourceTrigger int) templ.Component {

	if resource == nil || len(resource.ResourceTrigger) >= numResourceTrigger {
		return StringInput("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].Id", nil)
	}
	return StringInput("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].Id", resource.ResourceTrigger[numResourceTrigger].Id)
}
func (resource *SubscriptionTopic) T_ResourceTriggerDescription(numResourceTrigger int) templ.Component {

	if resource == nil || len(resource.ResourceTrigger) >= numResourceTrigger {
		return StringInput("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].Description", nil)
	}
	return StringInput("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].Description", resource.ResourceTrigger[numResourceTrigger].Description)
}
func (resource *SubscriptionTopic) T_ResourceTriggerResource(numResourceTrigger int) templ.Component {

	if resource == nil || len(resource.ResourceTrigger) >= numResourceTrigger {
		return StringInput("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].Resource", nil)
	}
	return StringInput("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].Resource", &resource.ResourceTrigger[numResourceTrigger].Resource)
}
func (resource *SubscriptionTopic) T_ResourceTriggerSupportedInteraction(numResourceTrigger int, numSupportedInteraction int) templ.Component {
	optionsValueSet := VSInteraction_trigger

	if resource == nil || len(resource.ResourceTrigger) >= numResourceTrigger || len(resource.ResourceTrigger[numResourceTrigger].SupportedInteraction) >= numSupportedInteraction {
		return CodeSelect("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].SupportedInteraction["+strconv.Itoa(numSupportedInteraction)+"]", nil, optionsValueSet)
	}
	return CodeSelect("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].SupportedInteraction["+strconv.Itoa(numSupportedInteraction)+"]", &resource.ResourceTrigger[numResourceTrigger].SupportedInteraction[numSupportedInteraction], optionsValueSet)
}
func (resource *SubscriptionTopic) T_ResourceTriggerFhirPathCriteria(numResourceTrigger int) templ.Component {

	if resource == nil || len(resource.ResourceTrigger) >= numResourceTrigger {
		return StringInput("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].FhirPathCriteria", nil)
	}
	return StringInput("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].FhirPathCriteria", resource.ResourceTrigger[numResourceTrigger].FhirPathCriteria)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaId(numResourceTrigger int) templ.Component {

	if resource == nil || len(resource.ResourceTrigger) >= numResourceTrigger {
		return StringInput("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].QueryCriteria.Id", nil)
	}
	return StringInput("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].QueryCriteria.Id", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.Id)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaPrevious(numResourceTrigger int) templ.Component {

	if resource == nil || len(resource.ResourceTrigger) >= numResourceTrigger {
		return StringInput("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].QueryCriteria.Previous", nil)
	}
	return StringInput("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].QueryCriteria.Previous", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.Previous)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaResultForCreate(numResourceTrigger int) templ.Component {
	optionsValueSet := VSSubscriptiontopic_cr_behavior

	if resource == nil || len(resource.ResourceTrigger) >= numResourceTrigger {
		return CodeSelect("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].QueryCriteria.ResultForCreate", nil, optionsValueSet)
	}
	return CodeSelect("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].QueryCriteria.ResultForCreate", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.ResultForCreate, optionsValueSet)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaCurrent(numResourceTrigger int) templ.Component {

	if resource == nil || len(resource.ResourceTrigger) >= numResourceTrigger {
		return StringInput("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].QueryCriteria.Current", nil)
	}
	return StringInput("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].QueryCriteria.Current", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.Current)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaResultForDelete(numResourceTrigger int) templ.Component {
	optionsValueSet := VSSubscriptiontopic_cr_behavior

	if resource == nil || len(resource.ResourceTrigger) >= numResourceTrigger {
		return CodeSelect("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].QueryCriteria.ResultForDelete", nil, optionsValueSet)
	}
	return CodeSelect("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].QueryCriteria.ResultForDelete", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.ResultForDelete, optionsValueSet)
}
func (resource *SubscriptionTopic) T_ResourceTriggerQueryCriteriaRequireBoth(numResourceTrigger int) templ.Component {

	if resource == nil || len(resource.ResourceTrigger) >= numResourceTrigger {
		return BoolInput("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].QueryCriteria.RequireBoth", nil)
	}
	return BoolInput("SubscriptionTopic.ResourceTrigger["+strconv.Itoa(numResourceTrigger)+"].QueryCriteria.RequireBoth", resource.ResourceTrigger[numResourceTrigger].QueryCriteria.RequireBoth)
}
func (resource *SubscriptionTopic) T_EventTriggerId(numEventTrigger int) templ.Component {

	if resource == nil || len(resource.EventTrigger) >= numEventTrigger {
		return StringInput("SubscriptionTopic.EventTrigger["+strconv.Itoa(numEventTrigger)+"].Id", nil)
	}
	return StringInput("SubscriptionTopic.EventTrigger["+strconv.Itoa(numEventTrigger)+"].Id", resource.EventTrigger[numEventTrigger].Id)
}
func (resource *SubscriptionTopic) T_EventTriggerDescription(numEventTrigger int) templ.Component {

	if resource == nil || len(resource.EventTrigger) >= numEventTrigger {
		return StringInput("SubscriptionTopic.EventTrigger["+strconv.Itoa(numEventTrigger)+"].Description", nil)
	}
	return StringInput("SubscriptionTopic.EventTrigger["+strconv.Itoa(numEventTrigger)+"].Description", resource.EventTrigger[numEventTrigger].Description)
}
func (resource *SubscriptionTopic) T_EventTriggerEvent(numEventTrigger int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.EventTrigger) >= numEventTrigger {
		return CodeableConceptSelect("SubscriptionTopic.EventTrigger["+strconv.Itoa(numEventTrigger)+"].Event", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubscriptionTopic.EventTrigger["+strconv.Itoa(numEventTrigger)+"].Event", &resource.EventTrigger[numEventTrigger].Event, optionsValueSet)
}
func (resource *SubscriptionTopic) T_EventTriggerResource(numEventTrigger int) templ.Component {

	if resource == nil || len(resource.EventTrigger) >= numEventTrigger {
		return StringInput("SubscriptionTopic.EventTrigger["+strconv.Itoa(numEventTrigger)+"].Resource", nil)
	}
	return StringInput("SubscriptionTopic.EventTrigger["+strconv.Itoa(numEventTrigger)+"].Resource", &resource.EventTrigger[numEventTrigger].Resource)
}
func (resource *SubscriptionTopic) T_CanFilterById(numCanFilterBy int) templ.Component {

	if resource == nil || len(resource.CanFilterBy) >= numCanFilterBy {
		return StringInput("SubscriptionTopic.CanFilterBy["+strconv.Itoa(numCanFilterBy)+"].Id", nil)
	}
	return StringInput("SubscriptionTopic.CanFilterBy["+strconv.Itoa(numCanFilterBy)+"].Id", resource.CanFilterBy[numCanFilterBy].Id)
}
func (resource *SubscriptionTopic) T_CanFilterByDescription(numCanFilterBy int) templ.Component {

	if resource == nil || len(resource.CanFilterBy) >= numCanFilterBy {
		return StringInput("SubscriptionTopic.CanFilterBy["+strconv.Itoa(numCanFilterBy)+"].Description", nil)
	}
	return StringInput("SubscriptionTopic.CanFilterBy["+strconv.Itoa(numCanFilterBy)+"].Description", resource.CanFilterBy[numCanFilterBy].Description)
}
func (resource *SubscriptionTopic) T_CanFilterByResource(numCanFilterBy int) templ.Component {

	if resource == nil || len(resource.CanFilterBy) >= numCanFilterBy {
		return StringInput("SubscriptionTopic.CanFilterBy["+strconv.Itoa(numCanFilterBy)+"].Resource", nil)
	}
	return StringInput("SubscriptionTopic.CanFilterBy["+strconv.Itoa(numCanFilterBy)+"].Resource", resource.CanFilterBy[numCanFilterBy].Resource)
}
func (resource *SubscriptionTopic) T_CanFilterByFilterParameter(numCanFilterBy int) templ.Component {

	if resource == nil || len(resource.CanFilterBy) >= numCanFilterBy {
		return StringInput("SubscriptionTopic.CanFilterBy["+strconv.Itoa(numCanFilterBy)+"].FilterParameter", nil)
	}
	return StringInput("SubscriptionTopic.CanFilterBy["+strconv.Itoa(numCanFilterBy)+"].FilterParameter", &resource.CanFilterBy[numCanFilterBy].FilterParameter)
}
func (resource *SubscriptionTopic) T_CanFilterByFilterDefinition(numCanFilterBy int) templ.Component {

	if resource == nil || len(resource.CanFilterBy) >= numCanFilterBy {
		return StringInput("SubscriptionTopic.CanFilterBy["+strconv.Itoa(numCanFilterBy)+"].FilterDefinition", nil)
	}
	return StringInput("SubscriptionTopic.CanFilterBy["+strconv.Itoa(numCanFilterBy)+"].FilterDefinition", resource.CanFilterBy[numCanFilterBy].FilterDefinition)
}
func (resource *SubscriptionTopic) T_CanFilterByComparator(numCanFilterBy int, numComparator int) templ.Component {
	optionsValueSet := VSSearch_comparator

	if resource == nil || len(resource.CanFilterBy) >= numCanFilterBy || len(resource.CanFilterBy[numCanFilterBy].Comparator) >= numComparator {
		return CodeSelect("SubscriptionTopic.CanFilterBy["+strconv.Itoa(numCanFilterBy)+"].Comparator["+strconv.Itoa(numComparator)+"]", nil, optionsValueSet)
	}
	return CodeSelect("SubscriptionTopic.CanFilterBy["+strconv.Itoa(numCanFilterBy)+"].Comparator["+strconv.Itoa(numComparator)+"]", &resource.CanFilterBy[numCanFilterBy].Comparator[numComparator], optionsValueSet)
}
func (resource *SubscriptionTopic) T_CanFilterByModifier(numCanFilterBy int, numModifier int) templ.Component {
	optionsValueSet := VSSearch_modifier_code

	if resource == nil || len(resource.CanFilterBy) >= numCanFilterBy || len(resource.CanFilterBy[numCanFilterBy].Modifier) >= numModifier {
		return CodeSelect("SubscriptionTopic.CanFilterBy["+strconv.Itoa(numCanFilterBy)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet)
	}
	return CodeSelect("SubscriptionTopic.CanFilterBy["+strconv.Itoa(numCanFilterBy)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.CanFilterBy[numCanFilterBy].Modifier[numModifier], optionsValueSet)
}
func (resource *SubscriptionTopic) T_NotificationShapeId(numNotificationShape int) templ.Component {

	if resource == nil || len(resource.NotificationShape) >= numNotificationShape {
		return StringInput("SubscriptionTopic.NotificationShape["+strconv.Itoa(numNotificationShape)+"].Id", nil)
	}
	return StringInput("SubscriptionTopic.NotificationShape["+strconv.Itoa(numNotificationShape)+"].Id", resource.NotificationShape[numNotificationShape].Id)
}
func (resource *SubscriptionTopic) T_NotificationShapeResource(numNotificationShape int) templ.Component {

	if resource == nil || len(resource.NotificationShape) >= numNotificationShape {
		return StringInput("SubscriptionTopic.NotificationShape["+strconv.Itoa(numNotificationShape)+"].Resource", nil)
	}
	return StringInput("SubscriptionTopic.NotificationShape["+strconv.Itoa(numNotificationShape)+"].Resource", &resource.NotificationShape[numNotificationShape].Resource)
}
func (resource *SubscriptionTopic) T_NotificationShapeInclude(numNotificationShape int, numInclude int) templ.Component {

	if resource == nil || len(resource.NotificationShape) >= numNotificationShape || len(resource.NotificationShape[numNotificationShape].Include) >= numInclude {
		return StringInput("SubscriptionTopic.NotificationShape["+strconv.Itoa(numNotificationShape)+"].Include["+strconv.Itoa(numInclude)+"]", nil)
	}
	return StringInput("SubscriptionTopic.NotificationShape["+strconv.Itoa(numNotificationShape)+"].Include["+strconv.Itoa(numInclude)+"]", &resource.NotificationShape[numNotificationShape].Include[numInclude])
}
func (resource *SubscriptionTopic) T_NotificationShapeRevInclude(numNotificationShape int, numRevInclude int) templ.Component {

	if resource == nil || len(resource.NotificationShape) >= numNotificationShape || len(resource.NotificationShape[numNotificationShape].RevInclude) >= numRevInclude {
		return StringInput("SubscriptionTopic.NotificationShape["+strconv.Itoa(numNotificationShape)+"].RevInclude["+strconv.Itoa(numRevInclude)+"]", nil)
	}
	return StringInput("SubscriptionTopic.NotificationShape["+strconv.Itoa(numNotificationShape)+"].RevInclude["+strconv.Itoa(numRevInclude)+"]", &resource.NotificationShape[numNotificationShape].RevInclude[numRevInclude])
}
