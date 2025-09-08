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
func (resource *EvidenceReport) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *EvidenceReport) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_CiteAsMarkdown(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("CiteAsMarkdown", nil, htmlAttrs)
	}
	return StringInput("CiteAsMarkdown", resource.CiteAsMarkdown, htmlAttrs)
}
func (resource *EvidenceReport) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *EvidenceReport) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *EvidenceReport) T_SubjectNote(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Subject.Note) {
		return AnnotationTextArea("SubjectNote["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("SubjectNote["+strconv.Itoa(numNote)+"]", &resource.Subject.Note[numNote], htmlAttrs)
}
func (resource *EvidenceReport) T_SubjectCharacteristicCode(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Subject.Characteristic) {
		return CodeableConceptSelect("SubjectCharacteristic["+strconv.Itoa(numCharacteristic)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubjectCharacteristic["+strconv.Itoa(numCharacteristic)+"].Code", &resource.Subject.Characteristic[numCharacteristic].Code, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_SubjectCharacteristicValueCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Subject.Characteristic) {
		return CodeableConceptSelect("SubjectCharacteristic["+strconv.Itoa(numCharacteristic)+"].ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubjectCharacteristic["+strconv.Itoa(numCharacteristic)+"].ValueCodeableConcept", &resource.Subject.Characteristic[numCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_SubjectCharacteristicValueBoolean(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Subject.Characteristic) {
		return BoolInput("SubjectCharacteristic["+strconv.Itoa(numCharacteristic)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("SubjectCharacteristic["+strconv.Itoa(numCharacteristic)+"].ValueBoolean", &resource.Subject.Characteristic[numCharacteristic].ValueBoolean, htmlAttrs)
}
func (resource *EvidenceReport) T_SubjectCharacteristicExclude(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Subject.Characteristic) {
		return BoolInput("SubjectCharacteristic["+strconv.Itoa(numCharacteristic)+"].Exclude", nil, htmlAttrs)
	}
	return BoolInput("SubjectCharacteristic["+strconv.Itoa(numCharacteristic)+"].Exclude", resource.Subject.Characteristic[numCharacteristic].Exclude, htmlAttrs)
}
func (resource *EvidenceReport) T_RelatesToCode(numRelatesTo int, htmlAttrs string) templ.Component {
	optionsValueSet := VSReport_relation_type

	if resource == nil || numRelatesTo >= len(resource.RelatesTo) {
		return CodeSelect("RelatesTo["+strconv.Itoa(numRelatesTo)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("RelatesTo["+strconv.Itoa(numRelatesTo)+"]Code", &resource.RelatesTo[numRelatesTo].Code, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_RelatesToTargetUrl(numRelatesTo int, htmlAttrs string) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.RelatesTo) {
		return StringInput("RelatesTo["+strconv.Itoa(numRelatesTo)+"]Target.Url", nil, htmlAttrs)
	}
	return StringInput("RelatesTo["+strconv.Itoa(numRelatesTo)+"]Target.Url", resource.RelatesTo[numRelatesTo].Target.Url, htmlAttrs)
}
func (resource *EvidenceReport) T_RelatesToTargetDisplay(numRelatesTo int, htmlAttrs string) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.RelatesTo) {
		return StringInput("RelatesTo["+strconv.Itoa(numRelatesTo)+"]Target.Display", nil, htmlAttrs)
	}
	return StringInput("RelatesTo["+strconv.Itoa(numRelatesTo)+"]Target.Display", resource.RelatesTo[numRelatesTo].Target.Display, htmlAttrs)
}
func (resource *EvidenceReport) T_SectionTitle(numSection int, htmlAttrs string) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return StringInput("Section["+strconv.Itoa(numSection)+"]Title", nil, htmlAttrs)
	}
	return StringInput("Section["+strconv.Itoa(numSection)+"]Title", resource.Section[numSection].Title, htmlAttrs)
}
func (resource *EvidenceReport) T_SectionFocus(numSection int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("Section["+strconv.Itoa(numSection)+"]Focus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Section["+strconv.Itoa(numSection)+"]Focus", resource.Section[numSection].Focus, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_SectionMode(numSection int, htmlAttrs string) templ.Component {
	optionsValueSet := VSList_mode

	if resource == nil || numSection >= len(resource.Section) {
		return CodeSelect("Section["+strconv.Itoa(numSection)+"]Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Section["+strconv.Itoa(numSection)+"]Mode", resource.Section[numSection].Mode, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_SectionOrderedBy(numSection int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("Section["+strconv.Itoa(numSection)+"]OrderedBy", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Section["+strconv.Itoa(numSection)+"]OrderedBy", resource.Section[numSection].OrderedBy, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_SectionEntryClassifier(numSection int, numEntryClassifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSection >= len(resource.Section) || numEntryClassifier >= len(resource.Section[numSection].EntryClassifier) {
		return CodeableConceptSelect("Section["+strconv.Itoa(numSection)+"]EntryClassifier["+strconv.Itoa(numEntryClassifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Section["+strconv.Itoa(numSection)+"]EntryClassifier["+strconv.Itoa(numEntryClassifier)+"]", &resource.Section[numSection].EntryClassifier[numEntryClassifier], optionsValueSet, htmlAttrs)
}
func (resource *EvidenceReport) T_SectionEmptyReason(numSection int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("Section["+strconv.Itoa(numSection)+"]EmptyReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Section["+strconv.Itoa(numSection)+"]EmptyReason", resource.Section[numSection].EmptyReason, optionsValueSet, htmlAttrs)
}
