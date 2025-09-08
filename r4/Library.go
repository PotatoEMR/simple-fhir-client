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

// http://hl7.org/fhir/r4/StructureDefinition/Library
type Library struct {
	Id                     *string               `json:"id,omitempty"`
	Meta                   *Meta                 `json:"meta,omitempty"`
	ImplicitRules          *string               `json:"implicitRules,omitempty"`
	Language               *string               `json:"language,omitempty"`
	Text                   *Narrative            `json:"text,omitempty"`
	Contained              []Resource            `json:"contained,omitempty"`
	Extension              []Extension           `json:"extension,omitempty"`
	ModifierExtension      []Extension           `json:"modifierExtension,omitempty"`
	Url                    *string               `json:"url,omitempty"`
	Identifier             []Identifier          `json:"identifier,omitempty"`
	Version                *string               `json:"version,omitempty"`
	Name                   *string               `json:"name,omitempty"`
	Title                  *string               `json:"title,omitempty"`
	Subtitle               *string               `json:"subtitle,omitempty"`
	Status                 string                `json:"status"`
	Experimental           *bool                 `json:"experimental,omitempty"`
	Type                   CodeableConcept       `json:"type"`
	SubjectCodeableConcept *CodeableConcept      `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference            `json:"subjectReference,omitempty"`
	Date                   *time.Time            `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher              *string               `json:"publisher,omitempty"`
	Contact                []ContactDetail       `json:"contact,omitempty"`
	Description            *string               `json:"description,omitempty"`
	UseContext             []UsageContext        `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept     `json:"jurisdiction,omitempty"`
	Purpose                *string               `json:"purpose,omitempty"`
	Usage                  *string               `json:"usage,omitempty"`
	Copyright              *string               `json:"copyright,omitempty"`
	ApprovalDate           *time.Time            `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate         *time.Time            `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
	EffectivePeriod        *Period               `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept     `json:"topic,omitempty"`
	Author                 []ContactDetail       `json:"author,omitempty"`
	Editor                 []ContactDetail       `json:"editor,omitempty"`
	Reviewer               []ContactDetail       `json:"reviewer,omitempty"`
	Endorser               []ContactDetail       `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact     `json:"relatedArtifact,omitempty"`
	Parameter              []ParameterDefinition `json:"parameter,omitempty"`
	DataRequirement        []DataRequirement     `json:"dataRequirement,omitempty"`
	Content                []Attachment          `json:"content,omitempty"`
}

type OtherLibrary Library

// on convert struct to json, automatically add resourceType=Library
func (r Library) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherLibrary
		ResourceType string `json:"resourceType"`
	}{
		OtherLibrary: OtherLibrary(r),
		ResourceType: "Library",
	})
}
func (r Library) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Library/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Library"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Library) T_Url(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Library.Url", nil, htmlAttrs)
	}
	return StringInput("Library.Url", resource.Url, htmlAttrs)
}
func (resource *Library) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Library.Version", nil, htmlAttrs)
	}
	return StringInput("Library.Version", resource.Version, htmlAttrs)
}
func (resource *Library) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Library.Name", nil, htmlAttrs)
	}
	return StringInput("Library.Name", resource.Name, htmlAttrs)
}
func (resource *Library) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Library.Title", nil, htmlAttrs)
	}
	return StringInput("Library.Title", resource.Title, htmlAttrs)
}
func (resource *Library) T_Subtitle(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Library.Subtitle", nil, htmlAttrs)
	}
	return StringInput("Library.Subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *Library) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Library.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Library.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Library) T_Experimental(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("Library.Experimental", nil, htmlAttrs)
	}
	return BoolInput("Library.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *Library) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Library.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Library.Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Library) T_SubjectCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Library.SubjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Library.SubjectCodeableConcept", resource.SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Library) T_Date(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Library.Date", nil, htmlAttrs)
	}
	return DateTimeInput("Library.Date", resource.Date, htmlAttrs)
}
func (resource *Library) T_Publisher(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Library.Publisher", nil, htmlAttrs)
	}
	return StringInput("Library.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *Library) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Library.Description", nil, htmlAttrs)
	}
	return StringInput("Library.Description", resource.Description, htmlAttrs)
}
func (resource *Library) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Library.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Library.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *Library) T_Purpose(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Library.Purpose", nil, htmlAttrs)
	}
	return StringInput("Library.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *Library) T_Usage(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Library.Usage", nil, htmlAttrs)
	}
	return StringInput("Library.Usage", resource.Usage, htmlAttrs)
}
func (resource *Library) T_Copyright(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Library.Copyright", nil, htmlAttrs)
	}
	return StringInput("Library.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *Library) T_ApprovalDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("Library.ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("Library.ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *Library) T_LastReviewDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("Library.LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("Library.LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *Library) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("Library.Topic."+strconv.Itoa(numTopic)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Library.Topic."+strconv.Itoa(numTopic)+".", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
