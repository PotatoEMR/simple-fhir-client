package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	SubjectCodeableConcept *CodeableConcept  `json:"subjectCodeableConcept"`
	SubjectReference       *Reference        `json:"subjectReference"`
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

func (resource *ResearchDefinition) ResearchDefinitionLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ResearchDefinition) ResearchDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ResearchDefinition) ResearchDefinitionJurisdiction(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *ResearchDefinition) ResearchDefinitionTopic(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("topic", nil, optionsValueSet)
	}
	return CodeableConceptSelect("topic", &resource.Topic[0], optionsValueSet)
}
