package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Evidence
type Evidence struct {
	Id                 *string           `json:"id,omitempty"`
	Meta               *Meta             `json:"meta,omitempty"`
	ImplicitRules      *string           `json:"implicitRules,omitempty"`
	Language           *string           `json:"language,omitempty"`
	Text               *Narrative        `json:"text,omitempty"`
	Contained          []Resource        `json:"contained,omitempty"`
	Extension          []Extension       `json:"extension,omitempty"`
	ModifierExtension  []Extension       `json:"modifierExtension,omitempty"`
	Url                *string           `json:"url,omitempty"`
	Identifier         []Identifier      `json:"identifier,omitempty"`
	Version            *string           `json:"version,omitempty"`
	Name               *string           `json:"name,omitempty"`
	Title              *string           `json:"title,omitempty"`
	ShortTitle         *string           `json:"shortTitle,omitempty"`
	Subtitle           *string           `json:"subtitle,omitempty"`
	Status             string            `json:"status"`
	Date               *string           `json:"date,omitempty"`
	Publisher          *string           `json:"publisher,omitempty"`
	Contact            []ContactDetail   `json:"contact,omitempty"`
	Description        *string           `json:"description,omitempty"`
	Note               []Annotation      `json:"note,omitempty"`
	UseContext         []UsageContext    `json:"useContext,omitempty"`
	Jurisdiction       []CodeableConcept `json:"jurisdiction,omitempty"`
	Copyright          *string           `json:"copyright,omitempty"`
	ApprovalDate       *string           `json:"approvalDate,omitempty"`
	LastReviewDate     *string           `json:"lastReviewDate,omitempty"`
	EffectivePeriod    *Period           `json:"effectivePeriod,omitempty"`
	Topic              []CodeableConcept `json:"topic,omitempty"`
	Author             []ContactDetail   `json:"author,omitempty"`
	Editor             []ContactDetail   `json:"editor,omitempty"`
	Reviewer           []ContactDetail   `json:"reviewer,omitempty"`
	Endorser           []ContactDetail   `json:"endorser,omitempty"`
	RelatedArtifact    []RelatedArtifact `json:"relatedArtifact,omitempty"`
	ExposureBackground Reference         `json:"exposureBackground"`
	ExposureVariant    []Reference       `json:"exposureVariant,omitempty"`
	Outcome            []Reference       `json:"outcome,omitempty"`
}

type OtherEvidence Evidence

// on convert struct to json, automatically add resourceType=Evidence
func (r Evidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEvidence
		ResourceType string `json:"resourceType"`
	}{
		OtherEvidence: OtherEvidence(r),
		ResourceType:  "Evidence",
	})
}

func (resource *Evidence) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Id", nil)
	}
	return StringInput("Evidence.Id", resource.Id)
}
func (resource *Evidence) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Evidence.ImplicitRules", nil)
	}
	return StringInput("Evidence.ImplicitRules", resource.ImplicitRules)
}
func (resource *Evidence) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Evidence.Language", nil, optionsValueSet)
	}
	return CodeSelect("Evidence.Language", resource.Language, optionsValueSet)
}
func (resource *Evidence) T_Url() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Url", nil)
	}
	return StringInput("Evidence.Url", resource.Url)
}
func (resource *Evidence) T_Version() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Version", nil)
	}
	return StringInput("Evidence.Version", resource.Version)
}
func (resource *Evidence) T_Name() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Name", nil)
	}
	return StringInput("Evidence.Name", resource.Name)
}
func (resource *Evidence) T_Title() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Title", nil)
	}
	return StringInput("Evidence.Title", resource.Title)
}
func (resource *Evidence) T_ShortTitle() templ.Component {

	if resource == nil {
		return StringInput("Evidence.ShortTitle", nil)
	}
	return StringInput("Evidence.ShortTitle", resource.ShortTitle)
}
func (resource *Evidence) T_Subtitle() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Subtitle", nil)
	}
	return StringInput("Evidence.Subtitle", resource.Subtitle)
}
func (resource *Evidence) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Evidence.Status", nil, optionsValueSet)
	}
	return CodeSelect("Evidence.Status", &resource.Status, optionsValueSet)
}
func (resource *Evidence) T_Date() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Date", nil)
	}
	return StringInput("Evidence.Date", resource.Date)
}
func (resource *Evidence) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Publisher", nil)
	}
	return StringInput("Evidence.Publisher", resource.Publisher)
}
func (resource *Evidence) T_Description() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Description", nil)
	}
	return StringInput("Evidence.Description", resource.Description)
}
func (resource *Evidence) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("Evidence.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Evidence.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *Evidence) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("Evidence.Copyright", nil)
	}
	return StringInput("Evidence.Copyright", resource.Copyright)
}
func (resource *Evidence) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("Evidence.ApprovalDate", nil)
	}
	return StringInput("Evidence.ApprovalDate", resource.ApprovalDate)
}
func (resource *Evidence) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("Evidence.LastReviewDate", nil)
	}
	return StringInput("Evidence.LastReviewDate", resource.LastReviewDate)
}
func (resource *Evidence) T_Topic(numTopic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Topic) >= numTopic {
		return CodeableConceptSelect("Evidence.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Evidence.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet)
}
