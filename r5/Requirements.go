//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/Requirements
type Requirements struct {
	Id                     *string                 `json:"id,omitempty"`
	Meta                   *Meta                   `json:"meta,omitempty"`
	ImplicitRules          *string                 `json:"implicitRules,omitempty"`
	Language               *string                 `json:"language,omitempty"`
	Text                   *Narrative              `json:"text,omitempty"`
	Contained              []Resource              `json:"contained,omitempty"`
	Extension              []Extension             `json:"extension,omitempty"`
	ModifierExtension      []Extension             `json:"modifierExtension,omitempty"`
	Url                    *string                 `json:"url,omitempty"`
	Identifier             []Identifier            `json:"identifier,omitempty"`
	Version                *string                 `json:"version,omitempty"`
	VersionAlgorithmString *string                 `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                 `json:"versionAlgorithmCoding"`
	Name                   *string                 `json:"name,omitempty"`
	Title                  *string                 `json:"title,omitempty"`
	Status                 string                  `json:"status"`
	Experimental           *bool                   `json:"experimental,omitempty"`
	Date                   *string                 `json:"date,omitempty"`
	Publisher              *string                 `json:"publisher,omitempty"`
	Contact                []ContactDetail         `json:"contact,omitempty"`
	Description            *string                 `json:"description,omitempty"`
	UseContext             []UsageContext          `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept       `json:"jurisdiction,omitempty"`
	Purpose                *string                 `json:"purpose,omitempty"`
	Copyright              *string                 `json:"copyright,omitempty"`
	CopyrightLabel         *string                 `json:"copyrightLabel,omitempty"`
	DerivedFrom            []string                `json:"derivedFrom,omitempty"`
	Reference              []string                `json:"reference,omitempty"`
	Actor                  []string                `json:"actor,omitempty"`
	Statement              []RequirementsStatement `json:"statement,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Requirements
type RequirementsStatement struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Key               string      `json:"key"`
	Label             *string     `json:"label,omitempty"`
	Conformance       []string    `json:"conformance,omitempty"`
	Conditionality    *bool       `json:"conditionality,omitempty"`
	Requirement       string      `json:"requirement"`
	DerivedFrom       *string     `json:"derivedFrom,omitempty"`
	Parent            *string     `json:"parent,omitempty"`
	SatisfiedBy       []string    `json:"satisfiedBy,omitempty"`
	Reference         []string    `json:"reference,omitempty"`
	Source            []Reference `json:"source,omitempty"`
}

type OtherRequirements Requirements

// on convert struct to json, automatically add resourceType=Requirements
func (r Requirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRequirements
		ResourceType string `json:"resourceType"`
	}{
		OtherRequirements: OtherRequirements(r),
		ResourceType:      "Requirements",
	})
}
