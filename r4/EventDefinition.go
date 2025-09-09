package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/EventDefinition
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
	Date                   *time.Time          `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher              *string             `json:"publisher,omitempty"`
	Contact                []ContactDetail     `json:"contact,omitempty"`
	Description            *string             `json:"description,omitempty"`
	UseContext             []UsageContext      `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept   `json:"jurisdiction,omitempty"`
	Purpose                *string             `json:"purpose,omitempty"`
	Usage                  *string             `json:"usage,omitempty"`
	Copyright              *string             `json:"copyright,omitempty"`
	ApprovalDate           *time.Time          `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate         *time.Time          `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
func (r EventDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "EventDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "EventDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *EventDefinition) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EventDefinition.Url", nil, htmlAttrs)
	}
	return StringInput("EventDefinition.Url", resource.Url, htmlAttrs)
}
func (resource *EventDefinition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EventDefinition.Version", nil, htmlAttrs)
	}
	return StringInput("EventDefinition.Version", resource.Version, htmlAttrs)
}
func (resource *EventDefinition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EventDefinition.Name", nil, htmlAttrs)
	}
	return StringInput("EventDefinition.Name", resource.Name, htmlAttrs)
}
func (resource *EventDefinition) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EventDefinition.Title", nil, htmlAttrs)
	}
	return StringInput("EventDefinition.Title", resource.Title, htmlAttrs)
}
func (resource *EventDefinition) T_Subtitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EventDefinition.Subtitle", nil, htmlAttrs)
	}
	return StringInput("EventDefinition.Subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *EventDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("EventDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("EventDefinition.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EventDefinition) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("EventDefinition.Experimental", nil, htmlAttrs)
	}
	return BoolInput("EventDefinition.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *EventDefinition) T_SubjectCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("EventDefinition.SubjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EventDefinition.SubjectCodeableConcept", resource.SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *EventDefinition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("EventDefinition.Date", nil, htmlAttrs)
	}
	return DateTimeInput("EventDefinition.Date", resource.Date, htmlAttrs)
}
func (resource *EventDefinition) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EventDefinition.Publisher", nil, htmlAttrs)
	}
	return StringInput("EventDefinition.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *EventDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EventDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("EventDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *EventDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("EventDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EventDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *EventDefinition) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EventDefinition.Purpose", nil, htmlAttrs)
	}
	return StringInput("EventDefinition.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *EventDefinition) T_Usage(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EventDefinition.Usage", nil, htmlAttrs)
	}
	return StringInput("EventDefinition.Usage", resource.Usage, htmlAttrs)
}
func (resource *EventDefinition) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EventDefinition.Copyright", nil, htmlAttrs)
	}
	return StringInput("EventDefinition.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *EventDefinition) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("EventDefinition.ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("EventDefinition.ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *EventDefinition) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("EventDefinition.LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("EventDefinition.LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *EventDefinition) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("EventDefinition.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EventDefinition.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
