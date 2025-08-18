//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

// http://hl7.org/fhir/r4b/StructureDefinition/ResearchDefinition
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
