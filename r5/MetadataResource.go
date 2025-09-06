package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/MetadataResource
type MetadataResource struct {
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
	VersionAlgorithmString *string           `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding           `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string           `json:"name,omitempty"`
	Title                  *string           `json:"title,omitempty"`
	Status                 string            `json:"status"`
	Experimental           *bool             `json:"experimental,omitempty"`
	Date                   *string           `json:"date,omitempty"`
	Publisher              *string           `json:"publisher,omitempty"`
	Contact                []ContactDetail   `json:"contact,omitempty"`
	Description            *string           `json:"description,omitempty"`
	UseContext             []UsageContext    `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept `json:"jurisdiction,omitempty"`
	Purpose                *string           `json:"purpose,omitempty"`
	Copyright              *string           `json:"copyright,omitempty"`
	CopyrightLabel         *string           `json:"copyrightLabel,omitempty"`
	ApprovalDate           *string           `json:"approvalDate,omitempty"`
	LastReviewDate         *string           `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period           `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept `json:"topic,omitempty"`
	Author                 []ContactDetail   `json:"author,omitempty"`
	Editor                 []ContactDetail   `json:"editor,omitempty"`
	Reviewer               []ContactDetail   `json:"reviewer,omitempty"`
	Endorser               []ContactDetail   `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact `json:"relatedArtifact,omitempty"`
}

type OtherMetadataResource MetadataResource

// on convert struct to json, automatically add resourceType=MetadataResource
func (r MetadataResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMetadataResource
		ResourceType string `json:"resourceType"`
	}{
		OtherMetadataResource: OtherMetadataResource(r),
		ResourceType:          "MetadataResource",
	})
}

func (resource *MetadataResource) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Id", nil)
	}
	return StringInput("MetadataResource.Id", resource.Id)
}
func (resource *MetadataResource) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.ImplicitRules", nil)
	}
	return StringInput("MetadataResource.ImplicitRules", resource.ImplicitRules)
}
func (resource *MetadataResource) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MetadataResource.Language", nil, optionsValueSet)
	}
	return CodeSelect("MetadataResource.Language", resource.Language, optionsValueSet)
}
func (resource *MetadataResource) T_Url() templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Url", nil)
	}
	return StringInput("MetadataResource.Url", resource.Url)
}
func (resource *MetadataResource) T_Version() templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Version", nil)
	}
	return StringInput("MetadataResource.Version", resource.Version)
}
func (resource *MetadataResource) T_Name() templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Name", nil)
	}
	return StringInput("MetadataResource.Name", resource.Name)
}
func (resource *MetadataResource) T_Title() templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Title", nil)
	}
	return StringInput("MetadataResource.Title", resource.Title)
}
func (resource *MetadataResource) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("MetadataResource.Status", nil, optionsValueSet)
	}
	return CodeSelect("MetadataResource.Status", &resource.Status, optionsValueSet)
}
func (resource *MetadataResource) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("MetadataResource.Experimental", nil)
	}
	return BoolInput("MetadataResource.Experimental", resource.Experimental)
}
func (resource *MetadataResource) T_Date() templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Date", nil)
	}
	return StringInput("MetadataResource.Date", resource.Date)
}
func (resource *MetadataResource) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Publisher", nil)
	}
	return StringInput("MetadataResource.Publisher", resource.Publisher)
}
func (resource *MetadataResource) T_Description() templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Description", nil)
	}
	return StringInput("MetadataResource.Description", resource.Description)
}
func (resource *MetadataResource) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("MetadataResource.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MetadataResource.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *MetadataResource) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Purpose", nil)
	}
	return StringInput("MetadataResource.Purpose", resource.Purpose)
}
func (resource *MetadataResource) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Copyright", nil)
	}
	return StringInput("MetadataResource.Copyright", resource.Copyright)
}
func (resource *MetadataResource) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.CopyrightLabel", nil)
	}
	return StringInput("MetadataResource.CopyrightLabel", resource.CopyrightLabel)
}
func (resource *MetadataResource) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.ApprovalDate", nil)
	}
	return StringInput("MetadataResource.ApprovalDate", resource.ApprovalDate)
}
func (resource *MetadataResource) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.LastReviewDate", nil)
	}
	return StringInput("MetadataResource.LastReviewDate", resource.LastReviewDate)
}
func (resource *MetadataResource) T_Topic(numTopic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Topic) >= numTopic {
		return CodeableConceptSelect("MetadataResource.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MetadataResource.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet)
}
