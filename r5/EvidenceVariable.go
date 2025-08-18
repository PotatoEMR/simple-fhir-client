//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceVariable
type EvidenceVariable struct {
	Id                     *string                          `json:"id,omitempty"`
	Meta                   *Meta                            `json:"meta,omitempty"`
	ImplicitRules          *string                          `json:"implicitRules,omitempty"`
	Language               *string                          `json:"language,omitempty"`
	Text                   *Narrative                       `json:"text,omitempty"`
	Contained              []Resource                       `json:"contained,omitempty"`
	Extension              []Extension                      `json:"extension,omitempty"`
	ModifierExtension      []Extension                      `json:"modifierExtension,omitempty"`
	Url                    *string                          `json:"url,omitempty"`
	Identifier             []Identifier                     `json:"identifier,omitempty"`
	Version                *string                          `json:"version,omitempty"`
	VersionAlgorithmString *string                          `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                          `json:"versionAlgorithmCoding"`
	Name                   *string                          `json:"name,omitempty"`
	Title                  *string                          `json:"title,omitempty"`
	ShortTitle             *string                          `json:"shortTitle,omitempty"`
	Status                 string                           `json:"status"`
	Experimental           *bool                            `json:"experimental,omitempty"`
	Date                   *string                          `json:"date,omitempty"`
	Publisher              *string                          `json:"publisher,omitempty"`
	Contact                []ContactDetail                  `json:"contact,omitempty"`
	Description            *string                          `json:"description,omitempty"`
	Note                   []Annotation                     `json:"note,omitempty"`
	UseContext             []UsageContext                   `json:"useContext,omitempty"`
	Purpose                *string                          `json:"purpose,omitempty"`
	Copyright              *string                          `json:"copyright,omitempty"`
	CopyrightLabel         *string                          `json:"copyrightLabel,omitempty"`
	ApprovalDate           *string                          `json:"approvalDate,omitempty"`
	LastReviewDate         *string                          `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                          `json:"effectivePeriod,omitempty"`
	Author                 []ContactDetail                  `json:"author,omitempty"`
	Editor                 []ContactDetail                  `json:"editor,omitempty"`
	Reviewer               []ContactDetail                  `json:"reviewer,omitempty"`
	Endorser               []ContactDetail                  `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact                `json:"relatedArtifact,omitempty"`
	Actual                 *bool                            `json:"actual,omitempty"`
	Characteristic         []EvidenceVariableCharacteristic `json:"characteristic,omitempty"`
	Handling               *string                          `json:"handling,omitempty"`
	Category               []EvidenceVariableCategory       `json:"category,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceVariable
type EvidenceVariableCharacteristic struct {
	Id                        *string                                                 `json:"id,omitempty"`
	Extension                 []Extension                                             `json:"extension,omitempty"`
	ModifierExtension         []Extension                                             `json:"modifierExtension,omitempty"`
	LinkId                    *string                                                 `json:"linkId,omitempty"`
	Description               *string                                                 `json:"description,omitempty"`
	Note                      []Annotation                                            `json:"note,omitempty"`
	Exclude                   *bool                                                   `json:"exclude,omitempty"`
	DefinitionReference       *Reference                                              `json:"definitionReference,omitempty"`
	DefinitionCanonical       *string                                                 `json:"definitionCanonical,omitempty"`
	DefinitionCodeableConcept *CodeableConcept                                        `json:"definitionCodeableConcept,omitempty"`
	DefinitionExpression      *Expression                                             `json:"definitionExpression,omitempty"`
	DefinitionId              *string                                                 `json:"definitionId,omitempty"`
	DefinitionByTypeAndValue  *EvidenceVariableCharacteristicDefinitionByTypeAndValue `json:"definitionByTypeAndValue,omitempty"`
	DefinitionByCombination   *EvidenceVariableCharacteristicDefinitionByCombination  `json:"definitionByCombination,omitempty"`
	InstancesQuantity         *Quantity                                               `json:"instancesQuantity"`
	InstancesRange            *Range                                                  `json:"instancesRange"`
	DurationQuantity          *Quantity                                               `json:"durationQuantity"`
	DurationRange             *Range                                                  `json:"durationRange"`
	TimeFromEvent             []EvidenceVariableCharacteristicTimeFromEvent           `json:"timeFromEvent,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceVariable
type EvidenceVariableCharacteristicDefinitionByTypeAndValue struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept   `json:"type"`
	Method               []CodeableConcept `json:"method,omitempty"`
	Device               *Reference        `json:"device,omitempty"`
	ValueCodeableConcept CodeableConcept   `json:"valueCodeableConcept"`
	ValueBoolean         bool              `json:"valueBoolean"`
	ValueQuantity        Quantity          `json:"valueQuantity"`
	ValueRange           Range             `json:"valueRange"`
	ValueReference       Reference         `json:"valueReference"`
	ValueId              string            `json:"valueId"`
	Offset               *CodeableConcept  `json:"offset,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceVariable
type EvidenceVariableCharacteristicDefinitionByCombination struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Threshold         *int        `json:"threshold,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceVariable
type EvidenceVariableCharacteristicTimeFromEvent struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Description          *string          `json:"description,omitempty"`
	Note                 []Annotation     `json:"note,omitempty"`
	EventCodeableConcept *CodeableConcept `json:"eventCodeableConcept"`
	EventReference       *Reference       `json:"eventReference"`
	EventDateTime        *string          `json:"eventDateTime"`
	EventId              *string          `json:"eventId"`
	Quantity             *Quantity        `json:"quantity,omitempty"`
	Range                *Range           `json:"range,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceVariable
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
