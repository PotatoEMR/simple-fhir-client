//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/EvidenceVariable
type EvidenceVariable struct {
	Id                *string                          `json:"id,omitempty"`
	Meta              *Meta                            `json:"meta,omitempty"`
	ImplicitRules     *string                          `json:"implicitRules,omitempty"`
	Language          *string                          `json:"language,omitempty"`
	Text              *Narrative                       `json:"text,omitempty"`
	Contained         []Resource                       `json:"contained,omitempty"`
	Extension         []Extension                      `json:"extension,omitempty"`
	ModifierExtension []Extension                      `json:"modifierExtension,omitempty"`
	Url               *string                          `json:"url,omitempty"`
	Identifier        []Identifier                     `json:"identifier,omitempty"`
	Version           *string                          `json:"version,omitempty"`
	Name              *string                          `json:"name,omitempty"`
	Title             *string                          `json:"title,omitempty"`
	ShortTitle        *string                          `json:"shortTitle,omitempty"`
	Subtitle          *string                          `json:"subtitle,omitempty"`
	Status            string                           `json:"status"`
	Date              *string                          `json:"date,omitempty"`
	Publisher         *string                          `json:"publisher,omitempty"`
	Contact           []ContactDetail                  `json:"contact,omitempty"`
	Description       *string                          `json:"description,omitempty"`
	Note              []Annotation                     `json:"note,omitempty"`
	UseContext        []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                `json:"jurisdiction,omitempty"`
	Copyright         *string                          `json:"copyright,omitempty"`
	ApprovalDate      *string                          `json:"approvalDate,omitempty"`
	LastReviewDate    *string                          `json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period                          `json:"effectivePeriod,omitempty"`
	Topic             []CodeableConcept                `json:"topic,omitempty"`
	Author            []ContactDetail                  `json:"author,omitempty"`
	Editor            []ContactDetail                  `json:"editor,omitempty"`
	Reviewer          []ContactDetail                  `json:"reviewer,omitempty"`
	Endorser          []ContactDetail                  `json:"endorser,omitempty"`
	RelatedArtifact   []RelatedArtifact                `json:"relatedArtifact,omitempty"`
	Type              *string                          `json:"type,omitempty"`
	Characteristic    []EvidenceVariableCharacteristic `json:"characteristic"`
}

// http://hl7.org/fhir/r4/StructureDefinition/EvidenceVariable
type EvidenceVariableCharacteristic struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Description                  *string           `json:"description,omitempty"`
	DefinitionReference          Reference         `json:"definitionReference"`
	DefinitionCanonical          string            `json:"definitionCanonical"`
	DefinitionCodeableConcept    CodeableConcept   `json:"definitionCodeableConcept"`
	DefinitionExpression         Expression        `json:"definitionExpression"`
	DefinitionDataRequirement    DataRequirement   `json:"definitionDataRequirement"`
	DefinitionTriggerDefinition  TriggerDefinition `json:"definitionTriggerDefinition"`
	UsageContext                 []UsageContext    `json:"usageContext,omitempty"`
	Exclude                      *bool             `json:"exclude,omitempty"`
	ParticipantEffectiveDateTime *string           `json:"participantEffectiveDateTime"`
	ParticipantEffectivePeriod   *Period           `json:"participantEffectivePeriod"`
	ParticipantEffectiveDuration *Duration         `json:"participantEffectiveDuration"`
	ParticipantEffectiveTiming   *Timing           `json:"participantEffectiveTiming"`
	TimeFromStart                *Duration         `json:"timeFromStart,omitempty"`
	GroupMeasure                 *string           `json:"groupMeasure,omitempty"`
}

type OtherEvidenceVariable EvidenceVariable

// on convert struct to json, automatically add resourceType=EvidenceVariable
func (r EvidenceVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEvidenceVariable
		ResourceType string `json:"resourceType"`
	}{
		OtherEvidenceVariable: OtherEvidenceVariable(r),
		ResourceType:          "EvidenceVariable",
	})
}
