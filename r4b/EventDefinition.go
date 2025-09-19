package r4b

//generated with command go run ./bultaoreune
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
	Date                   *FhirDateTime       `json:"date,omitempty"`
	Publisher              *string             `json:"publisher,omitempty"`
	Contact                []ContactDetail     `json:"contact,omitempty"`
	Description            *string             `json:"description,omitempty"`
	UseContext             []UsageContext      `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept   `json:"jurisdiction,omitempty"`
	Purpose                *string             `json:"purpose,omitempty"`
	Usage                  *string             `json:"usage,omitempty"`
	Copyright              *string             `json:"copyright,omitempty"`
	ApprovalDate           *FhirDate           `json:"approvalDate,omitempty"`
	LastReviewDate         *FhirDate           `json:"lastReviewDate,omitempty"`
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
func (resource *EventDefinition) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *EventDefinition) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *EventDefinition) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *EventDefinition) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *EventDefinition) T_Subtitle(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("subtitle", nil, htmlAttrs)
	}
	return StringInput("subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *EventDefinition) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EventDefinition) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *EventDefinition) T_SubjectCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("subjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subjectCodeableConcept", resource.SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *EventDefinition) T_SubjectReference(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("subjectReference", nil, htmlAttrs)
	}
	return ReferenceInput("subjectReference", resource.SubjectReference, htmlAttrs)
}
func (resource *EventDefinition) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *EventDefinition) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *EventDefinition) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *EventDefinition) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *EventDefinition) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *EventDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *EventDefinition) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *EventDefinition) T_Usage(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("usage", nil, htmlAttrs)
	}
	return StringInput("usage", resource.Usage, htmlAttrs)
}
func (resource *EventDefinition) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *EventDefinition) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("approvalDate", nil, htmlAttrs)
	}
	return FhirDateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *EventDefinition) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("lastReviewDate", nil, htmlAttrs)
	}
	return FhirDateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *EventDefinition) T_EffectivePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("effectivePeriod", resource.EffectivePeriod, htmlAttrs)
}
func (resource *EventDefinition) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *EventDefinition) T_Author(numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthor >= len(resource.Author) {
		return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", &resource.Author[numAuthor], htmlAttrs)
}
func (resource *EventDefinition) T_Editor(numEditor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEditor >= len(resource.Editor) {
		return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", &resource.Editor[numEditor], htmlAttrs)
}
func (resource *EventDefinition) T_Reviewer(numReviewer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReviewer >= len(resource.Reviewer) {
		return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", &resource.Reviewer[numReviewer], htmlAttrs)
}
func (resource *EventDefinition) T_Endorser(numEndorser int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndorser >= len(resource.Endorser) {
		return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", &resource.Endorser[numEndorser], htmlAttrs)
}
func (resource *EventDefinition) T_RelatedArtifact(numRelatedArtifact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedArtifact >= len(resource.RelatedArtifact) {
		return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", &resource.RelatedArtifact[numRelatedArtifact], htmlAttrs)
}
func (resource *EventDefinition) T_Trigger(numTrigger int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTrigger >= len(resource.Trigger) {
		return TriggerDefinitionInput("trigger["+strconv.Itoa(numTrigger)+"]", nil, htmlAttrs)
	}
	return TriggerDefinitionInput("trigger["+strconv.Itoa(numTrigger)+"]", &resource.Trigger[numTrigger], htmlAttrs)
}
