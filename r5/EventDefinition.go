package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/EventDefinition
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
	VersionAlgorithmString *string             `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding             `json:"versionAlgorithmCoding,omitempty"`
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
	CopyrightLabel         *string             `json:"copyrightLabel,omitempty"`
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
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *EventDefinition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *EventDefinition) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *EventDefinition) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *EventDefinition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *EventDefinition) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *EventDefinition) T_Subtitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Subtitle", nil, htmlAttrs)
	}
	return StringInput("Subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *EventDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EventDefinition) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Experimental", nil, htmlAttrs)
	}
	return BoolInput("Experimental", resource.Experimental, htmlAttrs)
}
func (resource *EventDefinition) T_SubjectCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SubjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubjectCodeableConcept", resource.SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *EventDefinition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *EventDefinition) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *EventDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *EventDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *EventDefinition) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Purpose", nil, htmlAttrs)
	}
	return StringInput("Purpose", resource.Purpose, htmlAttrs)
}
func (resource *EventDefinition) T_Usage(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Usage", nil, htmlAttrs)
	}
	return StringInput("Usage", resource.Usage, htmlAttrs)
}
func (resource *EventDefinition) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Copyright", nil, htmlAttrs)
	}
	return StringInput("Copyright", resource.Copyright, htmlAttrs)
}
func (resource *EventDefinition) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *EventDefinition) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *EventDefinition) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *EventDefinition) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
