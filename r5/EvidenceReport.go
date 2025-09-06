package r5

//generated with command go run ./bultaoreune -nodownload
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

func (resource *EvidenceReport) T_Id() templ.Component {

	if resource == nil {
		return StringInput("EvidenceReport.Id", nil)
	}
	return StringInput("EvidenceReport.Id", resource.Id)
}
func (resource *EvidenceReport) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("EvidenceReport.ImplicitRules", nil)
	}
	return StringInput("EvidenceReport.ImplicitRules", resource.ImplicitRules)
}
func (resource *EvidenceReport) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("EvidenceReport.Language", nil, optionsValueSet)
	}
	return CodeSelect("EvidenceReport.Language", resource.Language, optionsValueSet)
}
func (resource *EvidenceReport) T_Url() templ.Component {

	if resource == nil {
		return StringInput("EvidenceReport.Url", nil)
	}
	return StringInput("EvidenceReport.Url", resource.Url)
}
func (resource *EvidenceReport) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("EvidenceReport.Status", nil, optionsValueSet)
	}
	return CodeSelect("EvidenceReport.Status", &resource.Status, optionsValueSet)
}
func (resource *EvidenceReport) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("EvidenceReport.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EvidenceReport.Type", resource.Type, optionsValueSet)
}
func (resource *EvidenceReport) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("EvidenceReport.Publisher", nil)
	}
	return StringInput("EvidenceReport.Publisher", resource.Publisher)
}
func (resource *EvidenceReport) T_SubjectId() templ.Component {

	if resource == nil {
		return StringInput("EvidenceReport.Subject.Id", nil)
	}
	return StringInput("EvidenceReport.Subject.Id", resource.Subject.Id)
}
func (resource *EvidenceReport) T_SubjectCharacteristicId(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Subject.Characteristic) >= numCharacteristic {
		return StringInput("EvidenceReport.Subject.Characteristic["+strconv.Itoa(numCharacteristic)+"].Id", nil)
	}
	return StringInput("EvidenceReport.Subject.Characteristic["+strconv.Itoa(numCharacteristic)+"].Id", resource.Subject.Characteristic[numCharacteristic].Id)
}
func (resource *EvidenceReport) T_SubjectCharacteristicCode(numCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Subject.Characteristic) >= numCharacteristic {
		return CodeableConceptSelect("EvidenceReport.Subject.Characteristic["+strconv.Itoa(numCharacteristic)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EvidenceReport.Subject.Characteristic["+strconv.Itoa(numCharacteristic)+"].Code", &resource.Subject.Characteristic[numCharacteristic].Code, optionsValueSet)
}
func (resource *EvidenceReport) T_SubjectCharacteristicExclude(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Subject.Characteristic) >= numCharacteristic {
		return BoolInput("EvidenceReport.Subject.Characteristic["+strconv.Itoa(numCharacteristic)+"].Exclude", nil)
	}
	return BoolInput("EvidenceReport.Subject.Characteristic["+strconv.Itoa(numCharacteristic)+"].Exclude", resource.Subject.Characteristic[numCharacteristic].Exclude)
}
func (resource *EvidenceReport) T_RelatesToId(numRelatesTo int) templ.Component {

	if resource == nil || len(resource.RelatesTo) >= numRelatesTo {
		return StringInput("EvidenceReport.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Id", nil)
	}
	return StringInput("EvidenceReport.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Id", resource.RelatesTo[numRelatesTo].Id)
}
func (resource *EvidenceReport) T_RelatesToCode(numRelatesTo int) templ.Component {
	optionsValueSet := VSReport_relation_type

	if resource == nil || len(resource.RelatesTo) >= numRelatesTo {
		return CodeSelect("EvidenceReport.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("EvidenceReport.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Code", &resource.RelatesTo[numRelatesTo].Code, optionsValueSet)
}
func (resource *EvidenceReport) T_RelatesToTargetId(numRelatesTo int) templ.Component {

	if resource == nil || len(resource.RelatesTo) >= numRelatesTo {
		return StringInput("EvidenceReport.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Target.Id", nil)
	}
	return StringInput("EvidenceReport.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Target.Id", resource.RelatesTo[numRelatesTo].Target.Id)
}
func (resource *EvidenceReport) T_RelatesToTargetUrl(numRelatesTo int) templ.Component {

	if resource == nil || len(resource.RelatesTo) >= numRelatesTo {
		return StringInput("EvidenceReport.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Target.Url", nil)
	}
	return StringInput("EvidenceReport.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Target.Url", resource.RelatesTo[numRelatesTo].Target.Url)
}
func (resource *EvidenceReport) T_RelatesToTargetDisplay(numRelatesTo int) templ.Component {

	if resource == nil || len(resource.RelatesTo) >= numRelatesTo {
		return StringInput("EvidenceReport.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Target.Display", nil)
	}
	return StringInput("EvidenceReport.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Target.Display", resource.RelatesTo[numRelatesTo].Target.Display)
}
func (resource *EvidenceReport) T_SectionId(numSection int) templ.Component {

	if resource == nil || len(resource.Section) >= numSection {
		return StringInput("EvidenceReport.Section["+strconv.Itoa(numSection)+"].Id", nil)
	}
	return StringInput("EvidenceReport.Section["+strconv.Itoa(numSection)+"].Id", resource.Section[numSection].Id)
}
func (resource *EvidenceReport) T_SectionTitle(numSection int) templ.Component {

	if resource == nil || len(resource.Section) >= numSection {
		return StringInput("EvidenceReport.Section["+strconv.Itoa(numSection)+"].Title", nil)
	}
	return StringInput("EvidenceReport.Section["+strconv.Itoa(numSection)+"].Title", resource.Section[numSection].Title)
}
func (resource *EvidenceReport) T_SectionFocus(numSection int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Section) >= numSection {
		return CodeableConceptSelect("EvidenceReport.Section["+strconv.Itoa(numSection)+"].Focus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EvidenceReport.Section["+strconv.Itoa(numSection)+"].Focus", resource.Section[numSection].Focus, optionsValueSet)
}
func (resource *EvidenceReport) T_SectionMode(numSection int) templ.Component {
	optionsValueSet := VSList_mode

	if resource == nil || len(resource.Section) >= numSection {
		return CodeSelect("EvidenceReport.Section["+strconv.Itoa(numSection)+"].Mode", nil, optionsValueSet)
	}
	return CodeSelect("EvidenceReport.Section["+strconv.Itoa(numSection)+"].Mode", resource.Section[numSection].Mode, optionsValueSet)
}
func (resource *EvidenceReport) T_SectionOrderedBy(numSection int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Section) >= numSection {
		return CodeableConceptSelect("EvidenceReport.Section["+strconv.Itoa(numSection)+"].OrderedBy", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EvidenceReport.Section["+strconv.Itoa(numSection)+"].OrderedBy", resource.Section[numSection].OrderedBy, optionsValueSet)
}
func (resource *EvidenceReport) T_SectionEntryClassifier(numSection int, numEntryClassifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Section) >= numSection || len(resource.Section[numSection].EntryClassifier) >= numEntryClassifier {
		return CodeableConceptSelect("EvidenceReport.Section["+strconv.Itoa(numSection)+"].EntryClassifier["+strconv.Itoa(numEntryClassifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EvidenceReport.Section["+strconv.Itoa(numSection)+"].EntryClassifier["+strconv.Itoa(numEntryClassifier)+"]", &resource.Section[numSection].EntryClassifier[numEntryClassifier], optionsValueSet)
}
func (resource *EvidenceReport) T_SectionEmptyReason(numSection int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Section) >= numSection {
		return CodeableConceptSelect("EvidenceReport.Section["+strconv.Itoa(numSection)+"].EmptyReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EvidenceReport.Section["+strconv.Itoa(numSection)+"].EmptyReason", resource.Section[numSection].EmptyReason, optionsValueSet)
}
