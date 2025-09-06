package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/EventDefinition
type EventDefinition struct {
	Id                     *string             `json:"id,omitempty"`
	Meta                   *Meta               `json:"meta,omitempty"`
	ImplicitRules          *string             `json:"implicitRules,omitempty"`
	Language               *string             `json:"language,omitempty"`
	Text                   *Narrative          `json:"text,omitempty"`
	Contained              []Resource          `json:"contained,omitempty"`
	Extension              []Extension         `json:"extension,omitempty"`
	ModifierExtension      []Extension         `json:"modifierExtension,omitempty"`
	Url                    *string             `json:"url,omitempty"`
	Identifier             []Identifier        `json:"identifier,omitempty"`
	Version                *string             `json:"version,omitempty"`
	Name                   *string             `json:"name,omitempty"`
	Title                  *string             `json:"title,omitempty"`
	Subtitle               *string             `json:"subtitle,omitempty"`
	Status                 string              `json:"status"`
	Experimental           *bool               `json:"experimental,omitempty"`
	SubjectCodeableConcept *CodeableConcept    `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference          `json:"subjectReference,omitempty"`
	Date                   *string             `json:"date,omitempty"`
	Publisher              *string             `json:"publisher,omitempty"`
	Contact                []ContactDetail     `json:"contact,omitempty"`
	Description            *string             `json:"description,omitempty"`
	UseContext             []UsageContext      `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept   `json:"jurisdiction,omitempty"`
	Purpose                *string             `json:"purpose,omitempty"`
	Usage                  *string             `json:"usage,omitempty"`
	Copyright              *string             `json:"copyright,omitempty"`
	ApprovalDate           *string             `json:"approvalDate,omitempty"`
	LastReviewDate         *string             `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period             `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept   `json:"topic,omitempty"`
	Author                 []ContactDetail     `json:"author,omitempty"`
	Editor                 []ContactDetail     `json:"editor,omitempty"`
	Reviewer               []ContactDetail     `json:"reviewer,omitempty"`
	Endorser               []ContactDetail     `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact   `json:"relatedArtifact,omitempty"`
	Trigger                []TriggerDefinition `json:"trigger"`
}

type OtherEventDefinition EventDefinition

// on convert struct to json, automatically add resourceType=EventDefinition
func (r EventDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEventDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherEventDefinition: OtherEventDefinition(r),
		ResourceType:         "EventDefinition",
	})
}

func (resource *EventDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("EventDefinition.Id", nil)
	}
	return StringInput("EventDefinition.Id", resource.Id)
}
func (resource *EventDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("EventDefinition.ImplicitRules", nil)
	}
	return StringInput("EventDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *EventDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("EventDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("EventDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *EventDefinition) T_Url() templ.Component {

	if resource == nil {
		return StringInput("EventDefinition.Url", nil)
	}
	return StringInput("EventDefinition.Url", resource.Url)
}
func (resource *EventDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("EventDefinition.Version", nil)
	}
	return StringInput("EventDefinition.Version", resource.Version)
}
func (resource *EventDefinition) T_Name() templ.Component {

	if resource == nil {
		return StringInput("EventDefinition.Name", nil)
	}
	return StringInput("EventDefinition.Name", resource.Name)
}
func (resource *EventDefinition) T_Title() templ.Component {

	if resource == nil {
		return StringInput("EventDefinition.Title", nil)
	}
	return StringInput("EventDefinition.Title", resource.Title)
}
func (resource *EventDefinition) T_Subtitle() templ.Component {

	if resource == nil {
		return StringInput("EventDefinition.Subtitle", nil)
	}
	return StringInput("EventDefinition.Subtitle", resource.Subtitle)
}
func (resource *EventDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("EventDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("EventDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *EventDefinition) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("EventDefinition.Experimental", nil)
	}
	return BoolInput("EventDefinition.Experimental", resource.Experimental)
}
func (resource *EventDefinition) T_Date() templ.Component {

	if resource == nil {
		return StringInput("EventDefinition.Date", nil)
	}
	return StringInput("EventDefinition.Date", resource.Date)
}
func (resource *EventDefinition) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("EventDefinition.Publisher", nil)
	}
	return StringInput("EventDefinition.Publisher", resource.Publisher)
}
func (resource *EventDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("EventDefinition.Description", nil)
	}
	return StringInput("EventDefinition.Description", resource.Description)
}
func (resource *EventDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("EventDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EventDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *EventDefinition) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("EventDefinition.Purpose", nil)
	}
	return StringInput("EventDefinition.Purpose", resource.Purpose)
}
func (resource *EventDefinition) T_Usage() templ.Component {

	if resource == nil {
		return StringInput("EventDefinition.Usage", nil)
	}
	return StringInput("EventDefinition.Usage", resource.Usage)
}
func (resource *EventDefinition) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("EventDefinition.Copyright", nil)
	}
	return StringInput("EventDefinition.Copyright", resource.Copyright)
}
func (resource *EventDefinition) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("EventDefinition.ApprovalDate", nil)
	}
	return StringInput("EventDefinition.ApprovalDate", resource.ApprovalDate)
}
func (resource *EventDefinition) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("EventDefinition.LastReviewDate", nil)
	}
	return StringInput("EventDefinition.LastReviewDate", resource.LastReviewDate)
}
func (resource *EventDefinition) T_Topic(numTopic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Topic) >= numTopic {
		return CodeableConceptSelect("EventDefinition.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EventDefinition.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet)
}
