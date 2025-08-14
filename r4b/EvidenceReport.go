//generated August 14 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

// http://hl7.org/fhir/r4b/StructureDefinition/EvidenceReport
type EvidenceReport struct {
	Id                *string                   `json:"id,omitempty"`
	Meta              *Meta                     `json:"meta,omitempty"`
	ImplicitRules     *string                   `json:"implicitRules,omitempty"`
	Language          *string                   `json:"language,omitempty"`
	Text              *Narrative                `json:"text,omitempty"`
	Contained         []Resource                `json:"contained,omitempty"`
	Extension         []Extension               `json:"extension,omitempty"`
	ModifierExtension []Extension               `json:"modifierExtension,omitempty"`
	Url               *string                   `json:"url,omitempty"`
	Status            string                    `json:"status"`
	UseContext        []UsageContext            `json:"useContext,omitempty"`
	Identifier        []Identifier              `json:"identifier,omitempty"`
	RelatedIdentifier []Identifier              `json:"relatedIdentifier,omitempty"`
	CiteAsReference   *Reference                `json:"citeAsReference"`
	CiteAsMarkdown    *string                   `json:"citeAsMarkdown"`
	Type              *CodeableConcept          `json:"type,omitempty"`
	Note              []Annotation              `json:"note,omitempty"`
	RelatedArtifact   []RelatedArtifact         `json:"relatedArtifact,omitempty"`
	Subject           EvidenceReportSubject     `json:"subject"`
	Publisher         *string                   `json:"publisher,omitempty"`
	Contact           []ContactDetail           `json:"contact,omitempty"`
	Author            []ContactDetail           `json:"author,omitempty"`
	Editor            []ContactDetail           `json:"editor,omitempty"`
	Reviewer          []ContactDetail           `json:"reviewer,omitempty"`
	Endorser          []ContactDetail           `json:"endorser,omitempty"`
	RelatesTo         []EvidenceReportRelatesTo `json:"relatesTo,omitempty"`
	Section           []EvidenceReportSection   `json:"section,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/EvidenceReport
type EvidenceReportSubject struct {
	Id                *string                               `json:"id,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	Characteristic    []EvidenceReportSubjectCharacteristic `json:"characteristic,omitempty"`
	Note              []Annotation                          `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/EvidenceReport
type EvidenceReportSubjectCharacteristic struct {
	Id                   *string         `json:"id,omitempty"`
	Extension            []Extension     `json:"extension,omitempty"`
	ModifierExtension    []Extension     `json:"modifierExtension,omitempty"`
	Code                 CodeableConcept `json:"code"`
	ValueReference       Reference       `json:"valueReference"`
	ValueCodeableConcept CodeableConcept `json:"valueCodeableConcept"`
	ValueBoolean         bool            `json:"valueBoolean"`
	ValueQuantity        Quantity        `json:"valueQuantity"`
	ValueRange           Range           `json:"valueRange"`
	Exclude              *bool           `json:"exclude,omitempty"`
	Period               *Period         `json:"period,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/EvidenceReport
type EvidenceReportRelatesTo struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	TargetIdentifier  Identifier  `json:"targetIdentifier"`
	TargetReference   Reference   `json:"targetReference"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/EvidenceReport
type EvidenceReportSection struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Title             *string           `json:"title,omitempty"`
	Focus             *CodeableConcept  `json:"focus,omitempty"`
	FocusReference    *Reference        `json:"focusReference,omitempty"`
	Author            []Reference       `json:"author,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Mode              *string           `json:"mode,omitempty"`
	OrderedBy         *CodeableConcept  `json:"orderedBy,omitempty"`
	EntryClassifier   []CodeableConcept `json:"entryClassifier,omitempty"`
	EntryReference    []Reference       `json:"entryReference,omitempty"`
	EntryQuantity     []Quantity        `json:"entryQuantity,omitempty"`
	EmptyReason       *CodeableConcept  `json:"emptyReason,omitempty"`
}

type OtherEvidenceReport EvidenceReport

// on convert struct to json, automatically add resourceType=EvidenceReport
func (r EvidenceReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEvidenceReport
		ResourceType string `json:"resourceType"`
	}{
		OtherEvidenceReport: OtherEvidenceReport(r),
		ResourceType:        "EvidenceReport",
	})
}
