//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/NamingSystem
type NamingSystem struct {
	Id                     *string                `json:"id,omitempty"`
	Meta                   *Meta                  `json:"meta,omitempty"`
	ImplicitRules          *string                `json:"implicitRules,omitempty"`
	Language               *string                `json:"language,omitempty"`
	Text                   *Narrative             `json:"text,omitempty"`
	Contained              []Resource             `json:"contained,omitempty"`
	Extension              []Extension            `json:"extension,omitempty"`
	ModifierExtension      []Extension            `json:"modifierExtension,omitempty"`
	Url                    *string                `json:"url,omitempty"`
	Identifier             []Identifier           `json:"identifier,omitempty"`
	Version                *string                `json:"version,omitempty"`
	VersionAlgorithmString *string                `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                `json:"versionAlgorithmCoding"`
	Name                   string                 `json:"name"`
	Title                  *string                `json:"title,omitempty"`
	Status                 string                 `json:"status"`
	Kind                   string                 `json:"kind"`
	Experimental           *bool                  `json:"experimental,omitempty"`
	Date                   string                 `json:"date"`
	Publisher              *string                `json:"publisher,omitempty"`
	Contact                []ContactDetail        `json:"contact,omitempty"`
	Responsible            *string                `json:"responsible,omitempty"`
	Type                   *CodeableConcept       `json:"type,omitempty"`
	Description            *string                `json:"description,omitempty"`
	UseContext             []UsageContext         `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept      `json:"jurisdiction,omitempty"`
	Purpose                *string                `json:"purpose,omitempty"`
	Copyright              *string                `json:"copyright,omitempty"`
	CopyrightLabel         *string                `json:"copyrightLabel,omitempty"`
	ApprovalDate           *string                `json:"approvalDate,omitempty"`
	LastReviewDate         *string                `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept      `json:"topic,omitempty"`
	Author                 []ContactDetail        `json:"author,omitempty"`
	Editor                 []ContactDetail        `json:"editor,omitempty"`
	Reviewer               []ContactDetail        `json:"reviewer,omitempty"`
	Endorser               []ContactDetail        `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact      `json:"relatedArtifact,omitempty"`
	Usage                  *string                `json:"usage,omitempty"`
	UniqueId               []NamingSystemUniqueId `json:"uniqueId"`
}

// http://hl7.org/fhir/r5/StructureDefinition/NamingSystem
type NamingSystemUniqueId struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Value             string      `json:"value"`
	Preferred         *bool       `json:"preferred,omitempty"`
	Comment           *string     `json:"comment,omitempty"`
	Period            *Period     `json:"period,omitempty"`
	Authoritative     *bool       `json:"authoritative,omitempty"`
}

type OtherNamingSystem NamingSystem

// on convert struct to json, automatically add resourceType=NamingSystem
func (r NamingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNamingSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherNamingSystem: OtherNamingSystem(r),
		ResourceType:      "NamingSystem",
	})
}
