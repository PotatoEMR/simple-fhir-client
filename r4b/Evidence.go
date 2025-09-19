package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Evidence
type Evidence struct {
	Id                 *string                      `json:"id,omitempty"`
	Meta               *Meta                        `json:"meta,omitempty"`
	ImplicitRules      *string                      `json:"implicitRules,omitempty"`
	Language           *string                      `json:"language,omitempty"`
	Text               *Narrative                   `json:"text,omitempty"`
	Contained          []Resource                   `json:"contained,omitempty"`
	Extension          []Extension                  `json:"extension,omitempty"`
	ModifierExtension  []Extension                  `json:"modifierExtension,omitempty"`
	Url                *string                      `json:"url,omitempty"`
	Identifier         []Identifier                 `json:"identifier,omitempty"`
	Version            *string                      `json:"version,omitempty"`
	Title              *string                      `json:"title,omitempty"`
	CiteAsReference    *Reference                   `json:"citeAsReference,omitempty"`
	CiteAsMarkdown     *string                      `json:"citeAsMarkdown,omitempty"`
	Status             string                       `json:"status"`
	Date               *FhirDateTime                `json:"date,omitempty"`
	UseContext         []UsageContext               `json:"useContext,omitempty"`
	ApprovalDate       *FhirDate                    `json:"approvalDate,omitempty"`
	LastReviewDate     *FhirDate                    `json:"lastReviewDate,omitempty"`
	Publisher          *string                      `json:"publisher,omitempty"`
	Contact            []ContactDetail              `json:"contact,omitempty"`
	Author             []ContactDetail              `json:"author,omitempty"`
	Editor             []ContactDetail              `json:"editor,omitempty"`
	Reviewer           []ContactDetail              `json:"reviewer,omitempty"`
	Endorser           []ContactDetail              `json:"endorser,omitempty"`
	RelatedArtifact    []RelatedArtifact            `json:"relatedArtifact,omitempty"`
	Description        *string                      `json:"description,omitempty"`
	Assertion          *string                      `json:"assertion,omitempty"`
	Note               []Annotation                 `json:"note,omitempty"`
	VariableDefinition []EvidenceVariableDefinition `json:"variableDefinition"`
	SynthesisType      *CodeableConcept             `json:"synthesisType,omitempty"`
	StudyType          *CodeableConcept             `json:"studyType,omitempty"`
	Statistic          []EvidenceStatistic          `json:"statistic,omitempty"`
	Certainty          []EvidenceCertainty          `json:"certainty,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Evidence
type EvidenceVariableDefinition struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Note              []Annotation     `json:"note,omitempty"`
	VariableRole      CodeableConcept  `json:"variableRole"`
	Observed          *Reference       `json:"observed,omitempty"`
	Intended          *Reference       `json:"intended,omitempty"`
	DirectnessMatch   *CodeableConcept `json:"directnessMatch,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Evidence
type EvidenceStatistic struct {
	Id                  *string                                `json:"id,omitempty"`
	Extension           []Extension                            `json:"extension,omitempty"`
	ModifierExtension   []Extension                            `json:"modifierExtension,omitempty"`
	Description         *string                                `json:"description,omitempty"`
	Note                []Annotation                           `json:"note,omitempty"`
	StatisticType       *CodeableConcept                       `json:"statisticType,omitempty"`
	Category            *CodeableConcept                       `json:"category,omitempty"`
	Quantity            *Quantity                              `json:"quantity,omitempty"`
	NumberOfEvents      *int                                   `json:"numberOfEvents,omitempty"`
	NumberAffected      *int                                   `json:"numberAffected,omitempty"`
	SampleSize          *EvidenceStatisticSampleSize           `json:"sampleSize,omitempty"`
	AttributeEstimate   []EvidenceStatisticAttributeEstimate   `json:"attributeEstimate,omitempty"`
	ModelCharacteristic []EvidenceStatisticModelCharacteristic `json:"modelCharacteristic,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Evidence
type EvidenceStatisticSampleSize struct {
	Id                   *string      `json:"id,omitempty"`
	Extension            []Extension  `json:"extension,omitempty"`
	ModifierExtension    []Extension  `json:"modifierExtension,omitempty"`
	Description          *string      `json:"description,omitempty"`
	Note                 []Annotation `json:"note,omitempty"`
	NumberOfStudies      *int         `json:"numberOfStudies,omitempty"`
	NumberOfParticipants *int         `json:"numberOfParticipants,omitempty"`
	KnownDataCount       *int         `json:"knownDataCount,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Evidence
type EvidenceStatisticAttributeEstimate struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Note              []Annotation     `json:"note,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Quantity          *Quantity        `json:"quantity,omitempty"`
	Level             *float64         `json:"level,omitempty"`
	Range             *Range           `json:"range,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Evidence
type EvidenceStatisticModelCharacteristic struct {
	Id                *string                                        `json:"id,omitempty"`
	Extension         []Extension                                    `json:"extension,omitempty"`
	ModifierExtension []Extension                                    `json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                `json:"code"`
	Value             *Quantity                                      `json:"value,omitempty"`
	Variable          []EvidenceStatisticModelCharacteristicVariable `json:"variable,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Evidence
type EvidenceStatisticModelCharacteristicVariable struct {
	Id                 *string           `json:"id,omitempty"`
	Extension          []Extension       `json:"extension,omitempty"`
	ModifierExtension  []Extension       `json:"modifierExtension,omitempty"`
	VariableDefinition Reference         `json:"variableDefinition"`
	Handling           *string           `json:"handling,omitempty"`
	ValueCategory      []CodeableConcept `json:"valueCategory,omitempty"`
	ValueQuantity      []Quantity        `json:"valueQuantity,omitempty"`
	ValueRange         []Range           `json:"valueRange,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Evidence
type EvidenceCertainty struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Note              []Annotation     `json:"note,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Rating            *CodeableConcept `json:"rating,omitempty"`
	Rater             *string          `json:"rater,omitempty"`
}

type OtherEvidence Evidence

// on convert struct to json, automatically add resourceType=Evidence
func (r Evidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEvidence
		ResourceType string `json:"resourceType"`
	}{
		OtherEvidence: OtherEvidence(r),
		ResourceType:  "Evidence",
	})
}
func (r Evidence) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Evidence/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Evidence"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Evidence) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *Evidence) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *Evidence) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *Evidence) T_CiteAsReference(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("citeAsReference", nil, htmlAttrs)
	}
	return ReferenceInput("citeAsReference", resource.CiteAsReference, htmlAttrs)
}
func (resource *Evidence) T_CiteAsMarkdown(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("citeAsMarkdown", nil, htmlAttrs)
	}
	return StringInput("citeAsMarkdown", resource.CiteAsMarkdown, htmlAttrs)
}
func (resource *Evidence) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *Evidence) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *Evidence) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("approvalDate", nil, htmlAttrs)
	}
	return FhirDateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *Evidence) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("lastReviewDate", nil, htmlAttrs)
	}
	return FhirDateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *Evidence) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *Evidence) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *Evidence) T_Author(numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthor >= len(resource.Author) {
		return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", &resource.Author[numAuthor], htmlAttrs)
}
func (resource *Evidence) T_Editor(numEditor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEditor >= len(resource.Editor) {
		return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", &resource.Editor[numEditor], htmlAttrs)
}
func (resource *Evidence) T_Reviewer(numReviewer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReviewer >= len(resource.Reviewer) {
		return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", &resource.Reviewer[numReviewer], htmlAttrs)
}
func (resource *Evidence) T_Endorser(numEndorser int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndorser >= len(resource.Endorser) {
		return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", &resource.Endorser[numEndorser], htmlAttrs)
}
func (resource *Evidence) T_RelatedArtifact(numRelatedArtifact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedArtifact >= len(resource.RelatedArtifact) {
		return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", &resource.RelatedArtifact[numRelatedArtifact], htmlAttrs)
}
func (resource *Evidence) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *Evidence) T_Assertion(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("assertion", nil, htmlAttrs)
	}
	return StringInput("assertion", resource.Assertion, htmlAttrs)
}
func (resource *Evidence) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Evidence) T_SynthesisType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("synthesisType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("synthesisType", resource.SynthesisType, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_StudyType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("studyType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("studyType", resource.StudyType, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_VariableDefinitionDescription(numVariableDefinition int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVariableDefinition >= len(resource.VariableDefinition) {
		return StringInput("variableDefinition["+strconv.Itoa(numVariableDefinition)+"].description", nil, htmlAttrs)
	}
	return StringInput("variableDefinition["+strconv.Itoa(numVariableDefinition)+"].description", resource.VariableDefinition[numVariableDefinition].Description, htmlAttrs)
}
func (resource *Evidence) T_VariableDefinitionNote(numVariableDefinition int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVariableDefinition >= len(resource.VariableDefinition) || numNote >= len(resource.VariableDefinition[numVariableDefinition].Note) {
		return AnnotationTextArea("variableDefinition["+strconv.Itoa(numVariableDefinition)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("variableDefinition["+strconv.Itoa(numVariableDefinition)+"].note["+strconv.Itoa(numNote)+"]", &resource.VariableDefinition[numVariableDefinition].Note[numNote], htmlAttrs)
}
func (resource *Evidence) T_VariableDefinitionVariableRole(numVariableDefinition int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVariableDefinition >= len(resource.VariableDefinition) {
		return CodeableConceptSelect("variableDefinition["+strconv.Itoa(numVariableDefinition)+"].variableRole", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("variableDefinition["+strconv.Itoa(numVariableDefinition)+"].variableRole", &resource.VariableDefinition[numVariableDefinition].VariableRole, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_VariableDefinitionObserved(numVariableDefinition int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVariableDefinition >= len(resource.VariableDefinition) {
		return ReferenceInput("variableDefinition["+strconv.Itoa(numVariableDefinition)+"].observed", nil, htmlAttrs)
	}
	return ReferenceInput("variableDefinition["+strconv.Itoa(numVariableDefinition)+"].observed", resource.VariableDefinition[numVariableDefinition].Observed, htmlAttrs)
}
func (resource *Evidence) T_VariableDefinitionIntended(numVariableDefinition int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVariableDefinition >= len(resource.VariableDefinition) {
		return ReferenceInput("variableDefinition["+strconv.Itoa(numVariableDefinition)+"].intended", nil, htmlAttrs)
	}
	return ReferenceInput("variableDefinition["+strconv.Itoa(numVariableDefinition)+"].intended", resource.VariableDefinition[numVariableDefinition].Intended, htmlAttrs)
}
func (resource *Evidence) T_VariableDefinitionDirectnessMatch(numVariableDefinition int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVariableDefinition >= len(resource.VariableDefinition) {
		return CodeableConceptSelect("variableDefinition["+strconv.Itoa(numVariableDefinition)+"].directnessMatch", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("variableDefinition["+strconv.Itoa(numVariableDefinition)+"].directnessMatch", resource.VariableDefinition[numVariableDefinition].DirectnessMatch, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_StatisticDescription(numStatistic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) {
		return StringInput("statistic["+strconv.Itoa(numStatistic)+"].description", nil, htmlAttrs)
	}
	return StringInput("statistic["+strconv.Itoa(numStatistic)+"].description", resource.Statistic[numStatistic].Description, htmlAttrs)
}
func (resource *Evidence) T_StatisticNote(numStatistic int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) || numNote >= len(resource.Statistic[numStatistic].Note) {
		return AnnotationTextArea("statistic["+strconv.Itoa(numStatistic)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("statistic["+strconv.Itoa(numStatistic)+"].note["+strconv.Itoa(numNote)+"]", &resource.Statistic[numStatistic].Note[numNote], htmlAttrs)
}
func (resource *Evidence) T_StatisticStatisticType(numStatistic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) {
		return CodeableConceptSelect("statistic["+strconv.Itoa(numStatistic)+"].statisticType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statistic["+strconv.Itoa(numStatistic)+"].statisticType", resource.Statistic[numStatistic].StatisticType, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_StatisticCategory(numStatistic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) {
		return CodeableConceptSelect("statistic["+strconv.Itoa(numStatistic)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statistic["+strconv.Itoa(numStatistic)+"].category", resource.Statistic[numStatistic].Category, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_StatisticQuantity(numStatistic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) {
		return QuantityInput("statistic["+strconv.Itoa(numStatistic)+"].quantity", nil, htmlAttrs)
	}
	return QuantityInput("statistic["+strconv.Itoa(numStatistic)+"].quantity", resource.Statistic[numStatistic].Quantity, htmlAttrs)
}
func (resource *Evidence) T_StatisticNumberOfEvents(numStatistic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) {
		return IntInput("statistic["+strconv.Itoa(numStatistic)+"].numberOfEvents", nil, htmlAttrs)
	}
	return IntInput("statistic["+strconv.Itoa(numStatistic)+"].numberOfEvents", resource.Statistic[numStatistic].NumberOfEvents, htmlAttrs)
}
func (resource *Evidence) T_StatisticNumberAffected(numStatistic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) {
		return IntInput("statistic["+strconv.Itoa(numStatistic)+"].numberAffected", nil, htmlAttrs)
	}
	return IntInput("statistic["+strconv.Itoa(numStatistic)+"].numberAffected", resource.Statistic[numStatistic].NumberAffected, htmlAttrs)
}
func (resource *Evidence) T_StatisticSampleSizeDescription(numStatistic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) {
		return StringInput("statistic["+strconv.Itoa(numStatistic)+"].sampleSize.description", nil, htmlAttrs)
	}
	return StringInput("statistic["+strconv.Itoa(numStatistic)+"].sampleSize.description", resource.Statistic[numStatistic].SampleSize.Description, htmlAttrs)
}
func (resource *Evidence) T_StatisticSampleSizeNote(numStatistic int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) || numNote >= len(resource.Statistic[numStatistic].SampleSize.Note) {
		return AnnotationTextArea("statistic["+strconv.Itoa(numStatistic)+"].sampleSize.note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("statistic["+strconv.Itoa(numStatistic)+"].sampleSize.note["+strconv.Itoa(numNote)+"]", &resource.Statistic[numStatistic].SampleSize.Note[numNote], htmlAttrs)
}
func (resource *Evidence) T_StatisticSampleSizeNumberOfStudies(numStatistic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) {
		return IntInput("statistic["+strconv.Itoa(numStatistic)+"].sampleSize.numberOfStudies", nil, htmlAttrs)
	}
	return IntInput("statistic["+strconv.Itoa(numStatistic)+"].sampleSize.numberOfStudies", resource.Statistic[numStatistic].SampleSize.NumberOfStudies, htmlAttrs)
}
func (resource *Evidence) T_StatisticSampleSizeNumberOfParticipants(numStatistic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) {
		return IntInput("statistic["+strconv.Itoa(numStatistic)+"].sampleSize.numberOfParticipants", nil, htmlAttrs)
	}
	return IntInput("statistic["+strconv.Itoa(numStatistic)+"].sampleSize.numberOfParticipants", resource.Statistic[numStatistic].SampleSize.NumberOfParticipants, htmlAttrs)
}
func (resource *Evidence) T_StatisticSampleSizeKnownDataCount(numStatistic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) {
		return IntInput("statistic["+strconv.Itoa(numStatistic)+"].sampleSize.knownDataCount", nil, htmlAttrs)
	}
	return IntInput("statistic["+strconv.Itoa(numStatistic)+"].sampleSize.knownDataCount", resource.Statistic[numStatistic].SampleSize.KnownDataCount, htmlAttrs)
}
func (resource *Evidence) T_StatisticAttributeEstimateDescription(numStatistic int, numAttributeEstimate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) || numAttributeEstimate >= len(resource.Statistic[numStatistic].AttributeEstimate) {
		return StringInput("statistic["+strconv.Itoa(numStatistic)+"].attributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].description", nil, htmlAttrs)
	}
	return StringInput("statistic["+strconv.Itoa(numStatistic)+"].attributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].description", resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Description, htmlAttrs)
}
func (resource *Evidence) T_StatisticAttributeEstimateNote(numStatistic int, numAttributeEstimate int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) || numAttributeEstimate >= len(resource.Statistic[numStatistic].AttributeEstimate) || numNote >= len(resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Note) {
		return AnnotationTextArea("statistic["+strconv.Itoa(numStatistic)+"].attributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("statistic["+strconv.Itoa(numStatistic)+"].attributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].note["+strconv.Itoa(numNote)+"]", &resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Note[numNote], htmlAttrs)
}
func (resource *Evidence) T_StatisticAttributeEstimateType(numStatistic int, numAttributeEstimate int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) || numAttributeEstimate >= len(resource.Statistic[numStatistic].AttributeEstimate) {
		return CodeableConceptSelect("statistic["+strconv.Itoa(numStatistic)+"].attributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statistic["+strconv.Itoa(numStatistic)+"].attributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].type", resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Type, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_StatisticAttributeEstimateQuantity(numStatistic int, numAttributeEstimate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) || numAttributeEstimate >= len(resource.Statistic[numStatistic].AttributeEstimate) {
		return QuantityInput("statistic["+strconv.Itoa(numStatistic)+"].attributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].quantity", nil, htmlAttrs)
	}
	return QuantityInput("statistic["+strconv.Itoa(numStatistic)+"].attributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].quantity", resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Quantity, htmlAttrs)
}
func (resource *Evidence) T_StatisticAttributeEstimateLevel(numStatistic int, numAttributeEstimate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) || numAttributeEstimate >= len(resource.Statistic[numStatistic].AttributeEstimate) {
		return Float64Input("statistic["+strconv.Itoa(numStatistic)+"].attributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].level", nil, htmlAttrs)
	}
	return Float64Input("statistic["+strconv.Itoa(numStatistic)+"].attributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].level", resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Level, htmlAttrs)
}
func (resource *Evidence) T_StatisticAttributeEstimateRange(numStatistic int, numAttributeEstimate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) || numAttributeEstimate >= len(resource.Statistic[numStatistic].AttributeEstimate) {
		return RangeInput("statistic["+strconv.Itoa(numStatistic)+"].attributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].range", nil, htmlAttrs)
	}
	return RangeInput("statistic["+strconv.Itoa(numStatistic)+"].attributeEstimate["+strconv.Itoa(numAttributeEstimate)+"].range", resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Range, htmlAttrs)
}
func (resource *Evidence) T_StatisticModelCharacteristicCode(numStatistic int, numModelCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) || numModelCharacteristic >= len(resource.Statistic[numStatistic].ModelCharacteristic) {
		return CodeableConceptSelect("statistic["+strconv.Itoa(numStatistic)+"].modelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statistic["+strconv.Itoa(numStatistic)+"].modelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].code", &resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Code, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_StatisticModelCharacteristicValue(numStatistic int, numModelCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) || numModelCharacteristic >= len(resource.Statistic[numStatistic].ModelCharacteristic) {
		return QuantityInput("statistic["+strconv.Itoa(numStatistic)+"].modelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].value", nil, htmlAttrs)
	}
	return QuantityInput("statistic["+strconv.Itoa(numStatistic)+"].modelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].value", resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Value, htmlAttrs)
}
func (resource *Evidence) T_StatisticModelCharacteristicVariableVariableDefinition(numStatistic int, numModelCharacteristic int, numVariable int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) || numModelCharacteristic >= len(resource.Statistic[numStatistic].ModelCharacteristic) || numVariable >= len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable) {
		return ReferenceInput("statistic["+strconv.Itoa(numStatistic)+"].modelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].variable["+strconv.Itoa(numVariable)+"].variableDefinition", nil, htmlAttrs)
	}
	return ReferenceInput("statistic["+strconv.Itoa(numStatistic)+"].modelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].variable["+strconv.Itoa(numVariable)+"].variableDefinition", &resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].VariableDefinition, htmlAttrs)
}
func (resource *Evidence) T_StatisticModelCharacteristicVariableHandling(numStatistic int, numModelCharacteristic int, numVariable int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSVariable_handling

	if resource == nil || numStatistic >= len(resource.Statistic) || numModelCharacteristic >= len(resource.Statistic[numStatistic].ModelCharacteristic) || numVariable >= len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable) {
		return CodeSelect("statistic["+strconv.Itoa(numStatistic)+"].modelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].variable["+strconv.Itoa(numVariable)+"].handling", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("statistic["+strconv.Itoa(numStatistic)+"].modelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].variable["+strconv.Itoa(numVariable)+"].handling", resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].Handling, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_StatisticModelCharacteristicVariableValueCategory(numStatistic int, numModelCharacteristic int, numVariable int, numValueCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) || numModelCharacteristic >= len(resource.Statistic[numStatistic].ModelCharacteristic) || numVariable >= len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable) || numValueCategory >= len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].ValueCategory) {
		return CodeableConceptSelect("statistic["+strconv.Itoa(numStatistic)+"].modelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].variable["+strconv.Itoa(numVariable)+"].valueCategory["+strconv.Itoa(numValueCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statistic["+strconv.Itoa(numStatistic)+"].modelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].variable["+strconv.Itoa(numVariable)+"].valueCategory["+strconv.Itoa(numValueCategory)+"]", &resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].ValueCategory[numValueCategory], optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_StatisticModelCharacteristicVariableValueQuantity(numStatistic int, numModelCharacteristic int, numVariable int, numValueQuantity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) || numModelCharacteristic >= len(resource.Statistic[numStatistic].ModelCharacteristic) || numVariable >= len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable) || numValueQuantity >= len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].ValueQuantity) {
		return QuantityInput("statistic["+strconv.Itoa(numStatistic)+"].modelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].variable["+strconv.Itoa(numVariable)+"].valueQuantity["+strconv.Itoa(numValueQuantity)+"]", nil, htmlAttrs)
	}
	return QuantityInput("statistic["+strconv.Itoa(numStatistic)+"].modelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].variable["+strconv.Itoa(numVariable)+"].valueQuantity["+strconv.Itoa(numValueQuantity)+"]", &resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].ValueQuantity[numValueQuantity], htmlAttrs)
}
func (resource *Evidence) T_StatisticModelCharacteristicVariableValueRange(numStatistic int, numModelCharacteristic int, numVariable int, numValueRange int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatistic >= len(resource.Statistic) || numModelCharacteristic >= len(resource.Statistic[numStatistic].ModelCharacteristic) || numVariable >= len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable) || numValueRange >= len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].ValueRange) {
		return RangeInput("statistic["+strconv.Itoa(numStatistic)+"].modelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].variable["+strconv.Itoa(numVariable)+"].valueRange["+strconv.Itoa(numValueRange)+"]", nil, htmlAttrs)
	}
	return RangeInput("statistic["+strconv.Itoa(numStatistic)+"].modelCharacteristic["+strconv.Itoa(numModelCharacteristic)+"].variable["+strconv.Itoa(numVariable)+"].valueRange["+strconv.Itoa(numValueRange)+"]", &resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].ValueRange[numValueRange], htmlAttrs)
}
func (resource *Evidence) T_CertaintyDescription(numCertainty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) {
		return StringInput("certainty["+strconv.Itoa(numCertainty)+"].description", nil, htmlAttrs)
	}
	return StringInput("certainty["+strconv.Itoa(numCertainty)+"].description", resource.Certainty[numCertainty].Description, htmlAttrs)
}
func (resource *Evidence) T_CertaintyNote(numCertainty int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numNote >= len(resource.Certainty[numCertainty].Note) {
		return AnnotationTextArea("certainty["+strconv.Itoa(numCertainty)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("certainty["+strconv.Itoa(numCertainty)+"].note["+strconv.Itoa(numNote)+"]", &resource.Certainty[numCertainty].Note[numNote], htmlAttrs)
}
func (resource *Evidence) T_CertaintyType(numCertainty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) {
		return CodeableConceptSelect("certainty["+strconv.Itoa(numCertainty)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("certainty["+strconv.Itoa(numCertainty)+"].type", resource.Certainty[numCertainty].Type, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_CertaintyRating(numCertainty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) {
		return CodeableConceptSelect("certainty["+strconv.Itoa(numCertainty)+"].rating", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("certainty["+strconv.Itoa(numCertainty)+"].rating", resource.Certainty[numCertainty].Rating, optionsValueSet, htmlAttrs)
}
func (resource *Evidence) T_CertaintyRater(numCertainty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) {
		return StringInput("certainty["+strconv.Itoa(numCertainty)+"].rater", nil, htmlAttrs)
	}
	return StringInput("certainty["+strconv.Itoa(numCertainty)+"].rater", resource.Certainty[numCertainty].Rater, htmlAttrs)
}
