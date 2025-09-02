package r4b

//generated with command go run ./bultaoreune
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

func (resource *EvidenceReport) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *EvidenceReport) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *EvidenceReport) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *EvidenceReport) T_SubjectCharacteristicCode(numCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Subject.Characteristic) >= numCharacteristic {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Subject.Characteristic[numCharacteristic].Code, optionsValueSet)
}
func (resource *EvidenceReport) T_RelatesToCode(numRelatesTo int) templ.Component {
	optionsValueSet := VSReport_relation_type

	if resource == nil && len(resource.RelatesTo) >= numRelatesTo {
		return CodeSelect("code", nil, optionsValueSet)
	}
	return CodeSelect("code", &resource.RelatesTo[numRelatesTo].Code, optionsValueSet)
}
func (resource *EvidenceReport) T_SectionFocus(numSection int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Section) >= numSection {
		return CodeableConceptSelect("focus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("focus", resource.Section[numSection].Focus, optionsValueSet)
}
func (resource *EvidenceReport) T_SectionMode(numSection int) templ.Component {
	optionsValueSet := VSList_mode

	if resource == nil && len(resource.Section) >= numSection {
		return CodeSelect("mode", nil, optionsValueSet)
	}
	return CodeSelect("mode", resource.Section[numSection].Mode, optionsValueSet)
}
func (resource *EvidenceReport) T_SectionOrderedBy(numSection int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Section) >= numSection {
		return CodeableConceptSelect("orderedBy", nil, optionsValueSet)
	}
	return CodeableConceptSelect("orderedBy", resource.Section[numSection].OrderedBy, optionsValueSet)
}
func (resource *EvidenceReport) T_SectionEntryClassifier(numSection int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Section) >= numSection {
		return CodeableConceptSelect("entryClassifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("entryClassifier", &resource.Section[numSection].EntryClassifier[0], optionsValueSet)
}
func (resource *EvidenceReport) T_SectionEmptyReason(numSection int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Section) >= numSection {
		return CodeableConceptSelect("emptyReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("emptyReason", resource.Section[numSection].EmptyReason, optionsValueSet)
}
