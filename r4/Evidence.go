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
	Date               *time.Time        `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher          *string           `json:"publisher,omitempty"`
	Contact            []ContactDetail   `json:"contact,omitempty"`
	Description        *string           `json:"description,omitempty"`
	Note               []Annotation      `json:"note,omitempty"`
	UseContext         []UsageContext    `json:"useContext,omitempty"`
	Jurisdiction       []CodeableConcept `json:"jurisdiction,omitempty"`
	Copyright          *string           `json:"copyright,omitempty"`
	ApprovalDate       *time.Time        `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate     *time.Time        `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
func (r Evidence) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Evidence/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Evidence"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Evidence) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *Evidence) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *Evidence) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *Evidence) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *Evidence) T_ShortTitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ShortTitle", nil, htmlAttrs)
	}
	return StringInput("ShortTitle", resource.ShortTitle, htmlAttrs)
}
func (resource *Evidence) T_Subtitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Subtitle", nil, htmlAttrs)
	}
	return StringInput("Subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *Evidence) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *Evidence) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *Evidence) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *Evidence) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Evidence) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Copyright", nil, htmlAttrs)
	}
	return StringInput("Copyright", resource.Copyright, htmlAttrs)
}
func (resource *Evidence) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *Evidence) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *Evidence) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
