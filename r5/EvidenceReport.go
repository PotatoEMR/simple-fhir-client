package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceReport
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
	CiteAsReference   *Reference                `json:"citeAsReference,omitempty"`
	CiteAsMarkdown    *string                   `json:"citeAsMarkdown,omitempty"`
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

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceReport
type EvidenceReportSubject struct {
	Id                *string                               `json:"id,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	Characteristic    []EvidenceReportSubjectCharacteristic `json:"characteristic,omitempty"`
	Note              []Annotation                          `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceReport
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

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceReport
type EvidenceReportRelatesTo struct {
	Id                *string                       `json:"id,omitempty"`
	Extension         []Extension                   `json:"extension,omitempty"`
	ModifierExtension []Extension                   `json:"modifierExtension,omitempty"`
	Code              string                        `json:"code"`
	Target            EvidenceReportRelatesToTarget `json:"target"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceReport
type EvidenceReportRelatesToTarget struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Url               *string     `json:"url,omitempty"`
	Identifier        *Identifier `json:"identifier,omitempty"`
	Display           *string     `json:"display,omitempty"`
	Resource          *Reference  `json:"resource,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceReport
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
func (r EvidenceReport) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "EvidenceReport/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "EvidenceReport"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *EvidenceReport) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *EvidenceReport) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_CiteAsMarkdown(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("citeAsMarkdown", nil, htmlAttrs)
	}
	return StringInput("citeAsMarkdown", resource.CiteAsMarkdown, htmlAttrs)
}
func (resource *EvidenceReport) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *EvidenceReport) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *EvidenceReport) T_SubjectNote(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Subject.Note) {
		return AnnotationTextArea("subject.note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("subject.note["+strconv.Itoa(numNote)+"]", &resource.Subject.Note[numNote], htmlAttrs)
}
func (resource *EvidenceReport) T_SubjectCharacteristicCode(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Subject.Characteristic) {
		return CodeableConceptSelect("subject.characteristic["+strconv.Itoa(numCharacteristic)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subject.characteristic["+strconv.Itoa(numCharacteristic)+"].code", &resource.Subject.Characteristic[numCharacteristic].Code, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_SubjectCharacteristicValueCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Subject.Characteristic) {
		return CodeableConceptSelect("subject.characteristic["+strconv.Itoa(numCharacteristic)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subject.characteristic["+strconv.Itoa(numCharacteristic)+"].valueCodeableConcept", &resource.Subject.Characteristic[numCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_SubjectCharacteristicValueBoolean(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Subject.Characteristic) {
		return BoolInput("subject.characteristic["+strconv.Itoa(numCharacteristic)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("subject.characteristic["+strconv.Itoa(numCharacteristic)+"].valueBoolean", &resource.Subject.Characteristic[numCharacteristic].ValueBoolean, htmlAttrs)
}
func (resource *EvidenceReport) T_SubjectCharacteristicExclude(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Subject.Characteristic) {
		return BoolInput("subject.characteristic["+strconv.Itoa(numCharacteristic)+"].exclude", nil, htmlAttrs)
	}
	return BoolInput("subject.characteristic["+strconv.Itoa(numCharacteristic)+"].exclude", resource.Subject.Characteristic[numCharacteristic].Exclude, htmlAttrs)
}
func (resource *EvidenceReport) T_RelatesToCode(numRelatesTo int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSReport_relation_type

	if resource == nil || numRelatesTo >= len(resource.RelatesTo) {
		return CodeSelect("relatesTo["+strconv.Itoa(numRelatesTo)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("relatesTo["+strconv.Itoa(numRelatesTo)+"].code", &resource.RelatesTo[numRelatesTo].Code, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_RelatesToTargetUrl(numRelatesTo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.RelatesTo) {
		return StringInput("relatesTo["+strconv.Itoa(numRelatesTo)+"].target.url", nil, htmlAttrs)
	}
	return StringInput("relatesTo["+strconv.Itoa(numRelatesTo)+"].target.url", resource.RelatesTo[numRelatesTo].Target.Url, htmlAttrs)
}
func (resource *EvidenceReport) T_RelatesToTargetDisplay(numRelatesTo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.RelatesTo) {
		return StringInput("relatesTo["+strconv.Itoa(numRelatesTo)+"].target.display", nil, htmlAttrs)
	}
	return StringInput("relatesTo["+strconv.Itoa(numRelatesTo)+"].target.display", resource.RelatesTo[numRelatesTo].Target.Display, htmlAttrs)
}
func (resource *EvidenceReport) T_SectionTitle(numSection int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return StringInput("section["+strconv.Itoa(numSection)+"].title", nil, htmlAttrs)
	}
	return StringInput("section["+strconv.Itoa(numSection)+"].title", resource.Section[numSection].Title, htmlAttrs)
}
func (resource *EvidenceReport) T_SectionFocus(numSection int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].focus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].focus", resource.Section[numSection].Focus, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_SectionMode(numSection int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSList_mode

	if resource == nil || numSection >= len(resource.Section) {
		return CodeSelect("section["+strconv.Itoa(numSection)+"].mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("section["+strconv.Itoa(numSection)+"].mode", resource.Section[numSection].Mode, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_SectionOrderedBy(numSection int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].orderedBy", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].orderedBy", resource.Section[numSection].OrderedBy, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_SectionEntryClassifier(numSection int, numEntryClassifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) || numEntryClassifier >= len(resource.Section[numSection].EntryClassifier) {
		return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].entryClassifier["+strconv.Itoa(numEntryClassifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].entryClassifier["+strconv.Itoa(numEntryClassifier)+"]", &resource.Section[numSection].EntryClassifier[numEntryClassifier], optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_SectionEmptyReason(numSection int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].emptyReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].emptyReason", resource.Section[numSection].EmptyReason, optionsValueSet, htmlAttrs)
}
