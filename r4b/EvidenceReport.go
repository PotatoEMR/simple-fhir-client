package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *EvidenceReport) EvidenceReportLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *EvidenceReport) EvidenceReportStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *EvidenceReport) EvidenceReportRelatesToCode(numRelatesTo int) templ.Component {
	optionsValueSet := VSReport_relation_type
	currentVal := ""
	if resource != nil && len(resource.RelatesTo) >= numRelatesTo {
		currentVal = resource.RelatesTo[numRelatesTo].Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
func (resource *EvidenceReport) EvidenceReportSectionMode(numSection int) templ.Component {
	optionsValueSet := VSList_mode
	currentVal := ""
	if resource != nil && len(resource.Section) >= numSection {
		currentVal = *resource.Section[numSection].Mode
	}
	return CodeSelect("mode", currentVal, optionsValueSet)
}
