//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

// http://hl7.org/fhir/r4b/StructureDefinition/EvidenceVariable
type EvidenceVariable struct {
	Id                        *string                          `json:"id,omitempty"`
	Meta                      *Meta                            `json:"meta,omitempty"`
	ImplicitRules             *string                          `json:"implicitRules,omitempty"`
	Language                  *string                          `json:"language,omitempty"`
	Text                      *Narrative                       `json:"text,omitempty"`
	Contained                 []Resource                       `json:"contained,omitempty"`
	Extension                 []Extension                      `json:"extension,omitempty"`
	ModifierExtension         []Extension                      `json:"modifierExtension,omitempty"`
	Url                       *string                          `json:"url,omitempty"`
	Identifier                []Identifier                     `json:"identifier,omitempty"`
	Version                   *string                          `json:"version,omitempty"`
	Name                      *string                          `json:"name,omitempty"`
	Title                     *string                          `json:"title,omitempty"`
	ShortTitle                *string                          `json:"shortTitle,omitempty"`
	Subtitle                  *string                          `json:"subtitle,omitempty"`
	Status                    string                           `json:"status"`
	Date                      *string                          `json:"date,omitempty"`
	Description               *string                          `json:"description,omitempty"`
	Note                      []Annotation                     `json:"note,omitempty"`
	UseContext                []UsageContext                   `json:"useContext,omitempty"`
	Publisher                 *string                          `json:"publisher,omitempty"`
	Contact                   []ContactDetail                  `json:"contact,omitempty"`
	Author                    []ContactDetail                  `json:"author,omitempty"`
	Editor                    []ContactDetail                  `json:"editor,omitempty"`
	Reviewer                  []ContactDetail                  `json:"reviewer,omitempty"`
	Endorser                  []ContactDetail                  `json:"endorser,omitempty"`
	RelatedArtifact           []RelatedArtifact                `json:"relatedArtifact,omitempty"`
	Actual                    *bool                            `json:"actual,omitempty"`
	CharacteristicCombination *string                          `json:"characteristicCombination,omitempty"`
	Characteristic            []EvidenceVariableCharacteristic `json:"characteristic,omitempty"`
	Handling                  *string                          `json:"handling,omitempty"`
	Category                  []EvidenceVariableCategory       `json:"category,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/EvidenceVariable
type EvidenceVariableCharacteristic struct {
	Id                        *string                                      `json:"id,omitempty"`
	Extension                 []Extension                                  `json:"extension,omitempty"`
	ModifierExtension         []Extension                                  `json:"modifierExtension,omitempty"`
	Description               *string                                      `json:"description,omitempty"`
	DefinitionReference       Reference                                    `json:"definitionReference"`
	DefinitionCanonical       string                                       `json:"definitionCanonical"`
	DefinitionCodeableConcept CodeableConcept                              `json:"definitionCodeableConcept"`
	DefinitionExpression      Expression                                   `json:"definitionExpression"`
	Method                    *CodeableConcept                             `json:"method,omitempty"`
	Device                    *Reference                                   `json:"device,omitempty"`
	Exclude                   *bool                                        `json:"exclude,omitempty"`
	TimeFromStart             *EvidenceVariableCharacteristicTimeFromStart `json:"timeFromStart,omitempty"`
	GroupMeasure              *string                                      `json:"groupMeasure,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/EvidenceVariable
type EvidenceVariableCharacteristicTimeFromStart struct {
	Id                *string      `json:"id,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Description       *string      `json:"description,omitempty"`
	Quantity          *Quantity    `json:"quantity,omitempty"`
	Range             *Range       `json:"range,omitempty"`
	Note              []Annotation `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/EvidenceVariable
type EvidenceVariableCategory struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Name                 *string          `json:"name,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	ValueQuantity        *Quantity        `json:"valueQuantity"`
	ValueRange           *Range           `json:"valueRange"`
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
