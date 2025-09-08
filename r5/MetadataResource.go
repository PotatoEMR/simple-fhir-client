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
	Date                   *time.Time        `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher              *string           `json:"publisher,omitempty"`
	Contact                []ContactDetail   `json:"contact,omitempty"`
	Description            *string           `json:"description,omitempty"`
	UseContext             []UsageContext    `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept `json:"jurisdiction,omitempty"`
	Purpose                *string           `json:"purpose,omitempty"`
	Copyright              *string           `json:"copyright,omitempty"`
	CopyrightLabel         *string           `json:"copyrightLabel,omitempty"`
	ApprovalDate           *time.Time        `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate         *time.Time        `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
func (r MetadataResource) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MetadataResource/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "MetadataResource"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MetadataResource) T_Url(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Url", nil, htmlAttrs)
	}
	return StringInput("MetadataResource.Url", resource.Url, htmlAttrs)
}
func (resource *MetadataResource) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Version", nil, htmlAttrs)
	}
	return StringInput("MetadataResource.Version", resource.Version, htmlAttrs)
}
func (resource *MetadataResource) T_VersionAlgorithmString(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("MetadataResource.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *MetadataResource) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodingSelect("MetadataResource.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("MetadataResource.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *MetadataResource) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Name", nil, htmlAttrs)
	}
	return StringInput("MetadataResource.Name", resource.Name, htmlAttrs)
}
func (resource *MetadataResource) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Title", nil, htmlAttrs)
	}
	return StringInput("MetadataResource.Title", resource.Title, htmlAttrs)
}
func (resource *MetadataResource) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("MetadataResource.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("MetadataResource.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MetadataResource) T_Experimental(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("MetadataResource.Experimental", nil, htmlAttrs)
	}
	return BoolInput("MetadataResource.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *MetadataResource) T_Date(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("MetadataResource.Date", nil, htmlAttrs)
	}
	return DateTimeInput("MetadataResource.Date", resource.Date, htmlAttrs)
}
func (resource *MetadataResource) T_Publisher(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Publisher", nil, htmlAttrs)
	}
	return StringInput("MetadataResource.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *MetadataResource) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Description", nil, htmlAttrs)
	}
	return StringInput("MetadataResource.Description", resource.Description, htmlAttrs)
}
func (resource *MetadataResource) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("MetadataResource.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MetadataResource.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *MetadataResource) T_Purpose(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Purpose", nil, htmlAttrs)
	}
	return StringInput("MetadataResource.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *MetadataResource) T_Copyright(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.Copyright", nil, htmlAttrs)
	}
	return StringInput("MetadataResource.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *MetadataResource) T_CopyrightLabel(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MetadataResource.CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("MetadataResource.CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *MetadataResource) T_ApprovalDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("MetadataResource.ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("MetadataResource.ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *MetadataResource) T_LastReviewDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("MetadataResource.LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("MetadataResource.LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *MetadataResource) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("MetadataResource.Topic."+strconv.Itoa(numTopic)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MetadataResource.Topic."+strconv.Itoa(numTopic)+".", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
