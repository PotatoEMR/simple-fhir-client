package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Date                   *string               `json:"date,omitempty"`
	Publisher              *string               `json:"publisher,omitempty"`
	Contact                []ContactDetail       `json:"contact,omitempty"`
	Description            *string               `json:"description,omitempty"`
	UseContext             []UsageContext        `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept     `json:"jurisdiction,omitempty"`
	Purpose                *string               `json:"purpose,omitempty"`
	Usage                  *string               `json:"usage,omitempty"`
	Copyright              *string               `json:"copyright,omitempty"`
	ApprovalDate           *string               `json:"approvalDate,omitempty"`
	LastReviewDate         *string               `json:"lastReviewDate,omitempty"`
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

func (resource *Library) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Library.Id", nil)
	}
	return StringInput("Library.Id", resource.Id)
}
func (resource *Library) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Library.ImplicitRules", nil)
	}
	return StringInput("Library.ImplicitRules", resource.ImplicitRules)
}
func (resource *Library) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Library.Language", nil, optionsValueSet)
	}
	return CodeSelect("Library.Language", resource.Language, optionsValueSet)
}
func (resource *Library) T_Url() templ.Component {

	if resource == nil {
		return StringInput("Library.Url", nil)
	}
	return StringInput("Library.Url", resource.Url)
}
func (resource *Library) T_Version() templ.Component {

	if resource == nil {
		return StringInput("Library.Version", nil)
	}
	return StringInput("Library.Version", resource.Version)
}
func (resource *Library) T_Name() templ.Component {

	if resource == nil {
		return StringInput("Library.Name", nil)
	}
	return StringInput("Library.Name", resource.Name)
}
func (resource *Library) T_Title() templ.Component {

	if resource == nil {
		return StringInput("Library.Title", nil)
	}
	return StringInput("Library.Title", resource.Title)
}
func (resource *Library) T_Subtitle() templ.Component {

	if resource == nil {
		return StringInput("Library.Subtitle", nil)
	}
	return StringInput("Library.Subtitle", resource.Subtitle)
}
func (resource *Library) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Library.Status", nil, optionsValueSet)
	}
	return CodeSelect("Library.Status", &resource.Status, optionsValueSet)
}
func (resource *Library) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("Library.Experimental", nil)
	}
	return BoolInput("Library.Experimental", resource.Experimental)
}
func (resource *Library) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Library.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Library.Type", &resource.Type, optionsValueSet)
}
func (resource *Library) T_Date() templ.Component {

	if resource == nil {
		return StringInput("Library.Date", nil)
	}
	return StringInput("Library.Date", resource.Date)
}
func (resource *Library) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("Library.Publisher", nil)
	}
	return StringInput("Library.Publisher", resource.Publisher)
}
func (resource *Library) T_Description() templ.Component {

	if resource == nil {
		return StringInput("Library.Description", nil)
	}
	return StringInput("Library.Description", resource.Description)
}
func (resource *Library) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("Library.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Library.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *Library) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("Library.Purpose", nil)
	}
	return StringInput("Library.Purpose", resource.Purpose)
}
func (resource *Library) T_Usage() templ.Component {

	if resource == nil {
		return StringInput("Library.Usage", nil)
	}
	return StringInput("Library.Usage", resource.Usage)
}
func (resource *Library) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("Library.Copyright", nil)
	}
	return StringInput("Library.Copyright", resource.Copyright)
}
func (resource *Library) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("Library.ApprovalDate", nil)
	}
	return StringInput("Library.ApprovalDate", resource.ApprovalDate)
}
func (resource *Library) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("Library.LastReviewDate", nil)
	}
	return StringInput("Library.LastReviewDate", resource.LastReviewDate)
}
func (resource *Library) T_Topic(numTopic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Topic) >= numTopic {
		return CodeableConceptSelect("Library.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Library.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet)
}
