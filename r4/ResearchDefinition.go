package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/ResearchDefinition
type ResearchDefinition struct {
	Id                     *string           `json:"id,omitempty"`
	Meta                   *Meta             `json:"meta,omitempty"`
	ImplicitRules          *string           `json:"implicitRules,omitempty"`
	Language               *string           `json:"language,omitempty"`
	Text                   *Narrative        `json:"text,omitempty"`
	Contained              []Resource        `json:"contained,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Url                    *string           `json:"url,omitempty"`
	Identifier             []Identifier      `json:"identifier,omitempty"`
	Version                *string           `json:"version,omitempty"`
	Name                   *string           `json:"name,omitempty"`
	Title                  *string           `json:"title,omitempty"`
	ShortTitle             *string           `json:"shortTitle,omitempty"`
	Subtitle               *string           `json:"subtitle,omitempty"`
	Status                 string            `json:"status"`
	Experimental           *bool             `json:"experimental,omitempty"`
	SubjectCodeableConcept *CodeableConcept  `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference        `json:"subjectReference,omitempty"`
	Date                   *string           `json:"date,omitempty"`
	Publisher              *string           `json:"publisher,omitempty"`
	Contact                []ContactDetail   `json:"contact,omitempty"`
	Description            *string           `json:"description,omitempty"`
	Comment                []string          `json:"comment,omitempty"`
	UseContext             []UsageContext    `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept `json:"jurisdiction,omitempty"`
	Purpose                *string           `json:"purpose,omitempty"`
	Usage                  *string           `json:"usage,omitempty"`
	Copyright              *string           `json:"copyright,omitempty"`
	ApprovalDate           *string           `json:"approvalDate,omitempty"`
	LastReviewDate         *string           `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period           `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept `json:"topic,omitempty"`
	Author                 []ContactDetail   `json:"author,omitempty"`
	Editor                 []ContactDetail   `json:"editor,omitempty"`
	Reviewer               []ContactDetail   `json:"reviewer,omitempty"`
	Endorser               []ContactDetail   `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact `json:"relatedArtifact,omitempty"`
	Library                []string          `json:"library,omitempty"`
	Population             Reference         `json:"population"`
	Exposure               *Reference        `json:"exposure,omitempty"`
	ExposureAlternative    *Reference        `json:"exposureAlternative,omitempty"`
	Outcome                *Reference        `json:"outcome,omitempty"`
}

type OtherResearchDefinition ResearchDefinition

// on convert struct to json, automatically add resourceType=ResearchDefinition
func (r ResearchDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchDefinition: OtherResearchDefinition(r),
		ResourceType:            "ResearchDefinition",
	})
}

func (resource *ResearchDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ResearchDefinition.Id", nil)
	}
	return StringInput("ResearchDefinition.Id", resource.Id)
}
func (resource *ResearchDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ResearchDefinition.ImplicitRules", nil)
	}
	return StringInput("ResearchDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *ResearchDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ResearchDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("ResearchDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *ResearchDefinition) T_Url() templ.Component {

	if resource == nil {
		return StringInput("ResearchDefinition.Url", nil)
	}
	return StringInput("ResearchDefinition.Url", resource.Url)
}
func (resource *ResearchDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("ResearchDefinition.Version", nil)
	}
	return StringInput("ResearchDefinition.Version", resource.Version)
}
func (resource *ResearchDefinition) T_Name() templ.Component {

	if resource == nil {
		return StringInput("ResearchDefinition.Name", nil)
	}
	return StringInput("ResearchDefinition.Name", resource.Name)
}
func (resource *ResearchDefinition) T_Title() templ.Component {

	if resource == nil {
		return StringInput("ResearchDefinition.Title", nil)
	}
	return StringInput("ResearchDefinition.Title", resource.Title)
}
func (resource *ResearchDefinition) T_ShortTitle() templ.Component {

	if resource == nil {
		return StringInput("ResearchDefinition.ShortTitle", nil)
	}
	return StringInput("ResearchDefinition.ShortTitle", resource.ShortTitle)
}
func (resource *ResearchDefinition) T_Subtitle() templ.Component {

	if resource == nil {
		return StringInput("ResearchDefinition.Subtitle", nil)
	}
	return StringInput("ResearchDefinition.Subtitle", resource.Subtitle)
}
func (resource *ResearchDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ResearchDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("ResearchDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *ResearchDefinition) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("ResearchDefinition.Experimental", nil)
	}
	return BoolInput("ResearchDefinition.Experimental", resource.Experimental)
}
func (resource *ResearchDefinition) T_Date() templ.Component {

	if resource == nil {
		return StringInput("ResearchDefinition.Date", nil)
	}
	return StringInput("ResearchDefinition.Date", resource.Date)
}
func (resource *ResearchDefinition) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("ResearchDefinition.Publisher", nil)
	}
	return StringInput("ResearchDefinition.Publisher", resource.Publisher)
}
func (resource *ResearchDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("ResearchDefinition.Description", nil)
	}
	return StringInput("ResearchDefinition.Description", resource.Description)
}
func (resource *ResearchDefinition) T_Comment(numComment int) templ.Component {

	if resource == nil || len(resource.Comment) >= numComment {
		return StringInput("ResearchDefinition.Comment["+strconv.Itoa(numComment)+"]", nil)
	}
	return StringInput("ResearchDefinition.Comment["+strconv.Itoa(numComment)+"]", &resource.Comment[numComment])
}
func (resource *ResearchDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("ResearchDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *ResearchDefinition) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("ResearchDefinition.Purpose", nil)
	}
	return StringInput("ResearchDefinition.Purpose", resource.Purpose)
}
func (resource *ResearchDefinition) T_Usage() templ.Component {

	if resource == nil {
		return StringInput("ResearchDefinition.Usage", nil)
	}
	return StringInput("ResearchDefinition.Usage", resource.Usage)
}
func (resource *ResearchDefinition) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("ResearchDefinition.Copyright", nil)
	}
	return StringInput("ResearchDefinition.Copyright", resource.Copyright)
}
func (resource *ResearchDefinition) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("ResearchDefinition.ApprovalDate", nil)
	}
	return StringInput("ResearchDefinition.ApprovalDate", resource.ApprovalDate)
}
func (resource *ResearchDefinition) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("ResearchDefinition.LastReviewDate", nil)
	}
	return StringInput("ResearchDefinition.LastReviewDate", resource.LastReviewDate)
}
func (resource *ResearchDefinition) T_Topic(numTopic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Topic) >= numTopic {
		return CodeableConceptSelect("ResearchDefinition.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchDefinition.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet)
}
func (resource *ResearchDefinition) T_Library(numLibrary int) templ.Component {

	if resource == nil || len(resource.Library) >= numLibrary {
		return StringInput("ResearchDefinition.Library["+strconv.Itoa(numLibrary)+"]", nil)
	}
	return StringInput("ResearchDefinition.Library["+strconv.Itoa(numLibrary)+"]", &resource.Library[numLibrary])
}
