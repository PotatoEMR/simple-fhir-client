package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/EffectEvidenceSynthesis
type EffectEvidenceSynthesis struct {
	Id                  *string                                    `json:"id,omitempty"`
	Meta                *Meta                                      `json:"meta,omitempty"`
	ImplicitRules       *string                                    `json:"implicitRules,omitempty"`
	Language            *string                                    `json:"language,omitempty"`
	Text                *Narrative                                 `json:"text,omitempty"`
	Contained           []Resource                                 `json:"contained,omitempty"`
	Extension           []Extension                                `json:"extension,omitempty"`
	ModifierExtension   []Extension                                `json:"modifierExtension,omitempty"`
	Url                 *string                                    `json:"url,omitempty"`
	Identifier          []Identifier                               `json:"identifier,omitempty"`
	Version             *string                                    `json:"version,omitempty"`
	Name                *string                                    `json:"name,omitempty"`
	Title               *string                                    `json:"title,omitempty"`
	Status              string                                     `json:"status"`
	Date                *FhirDateTime                              `json:"date,omitempty"`
	Publisher           *string                                    `json:"publisher,omitempty"`
	Contact             []ContactDetail                            `json:"contact,omitempty"`
	Description         *string                                    `json:"description,omitempty"`
	Note                []Annotation                               `json:"note,omitempty"`
	UseContext          []UsageContext                             `json:"useContext,omitempty"`
	Jurisdiction        []CodeableConcept                          `json:"jurisdiction,omitempty"`
	Copyright           *string                                    `json:"copyright,omitempty"`
	ApprovalDate        *FhirDate                                  `json:"approvalDate,omitempty"`
	LastReviewDate      *FhirDate                                  `json:"lastReviewDate,omitempty"`
	EffectivePeriod     *Period                                    `json:"effectivePeriod,omitempty"`
	Topic               []CodeableConcept                          `json:"topic,omitempty"`
	Author              []ContactDetail                            `json:"author,omitempty"`
	Editor              []ContactDetail                            `json:"editor,omitempty"`
	Reviewer            []ContactDetail                            `json:"reviewer,omitempty"`
	Endorser            []ContactDetail                            `json:"endorser,omitempty"`
	RelatedArtifact     []RelatedArtifact                          `json:"relatedArtifact,omitempty"`
	SynthesisType       *CodeableConcept                           `json:"synthesisType,omitempty"`
	StudyType           *CodeableConcept                           `json:"studyType,omitempty"`
	Population          Reference                                  `json:"population"`
	Exposure            Reference                                  `json:"exposure"`
	ExposureAlternative Reference                                  `json:"exposureAlternative"`
	Outcome             Reference                                  `json:"outcome"`
	SampleSize          *EffectEvidenceSynthesisSampleSize         `json:"sampleSize,omitempty"`
	ResultsByExposure   []EffectEvidenceSynthesisResultsByExposure `json:"resultsByExposure,omitempty"`
	EffectEstimate      []EffectEvidenceSynthesisEffectEstimate    `json:"effectEstimate,omitempty"`
	Certainty           []EffectEvidenceSynthesisCertainty         `json:"certainty,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/EffectEvidenceSynthesis
type EffectEvidenceSynthesisSampleSize struct {
	Id                   *string     `json:"id,omitempty"`
	Extension            []Extension `json:"extension,omitempty"`
	ModifierExtension    []Extension `json:"modifierExtension,omitempty"`
	Description          *string     `json:"description,omitempty"`
	NumberOfStudies      *int        `json:"numberOfStudies,omitempty"`
	NumberOfParticipants *int        `json:"numberOfParticipants,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/EffectEvidenceSynthesis
type EffectEvidenceSynthesisResultsByExposure struct {
	Id                    *string          `json:"id,omitempty"`
	Extension             []Extension      `json:"extension,omitempty"`
	ModifierExtension     []Extension      `json:"modifierExtension,omitempty"`
	Description           *string          `json:"description,omitempty"`
	ExposureState         *string          `json:"exposureState,omitempty"`
	VariantState          *CodeableConcept `json:"variantState,omitempty"`
	RiskEvidenceSynthesis Reference        `json:"riskEvidenceSynthesis"`
}

// http://hl7.org/fhir/r4/StructureDefinition/EffectEvidenceSynthesis
type EffectEvidenceSynthesisEffectEstimate struct {
	Id                *string                                                  `json:"id,omitempty"`
	Extension         []Extension                                              `json:"extension,omitempty"`
	ModifierExtension []Extension                                              `json:"modifierExtension,omitempty"`
	Description       *string                                                  `json:"description,omitempty"`
	Type              *CodeableConcept                                         `json:"type,omitempty"`
	VariantState      *CodeableConcept                                         `json:"variantState,omitempty"`
	Value             *float64                                                 `json:"value,omitempty"`
	UnitOfMeasure     *CodeableConcept                                         `json:"unitOfMeasure,omitempty"`
	PrecisionEstimate []EffectEvidenceSynthesisEffectEstimatePrecisionEstimate `json:"precisionEstimate,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/EffectEvidenceSynthesis
type EffectEvidenceSynthesisEffectEstimatePrecisionEstimate struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Level             *float64         `json:"level,omitempty"`
	From              *float64         `json:"from,omitempty"`
	To                *float64         `json:"to,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/EffectEvidenceSynthesis
type EffectEvidenceSynthesisCertainty struct {
	Id                    *string                                                 `json:"id,omitempty"`
	Extension             []Extension                                             `json:"extension,omitempty"`
	ModifierExtension     []Extension                                             `json:"modifierExtension,omitempty"`
	Rating                []CodeableConcept                                       `json:"rating,omitempty"`
	Note                  []Annotation                                            `json:"note,omitempty"`
	CertaintySubcomponent []EffectEvidenceSynthesisCertaintyCertaintySubcomponent `json:"certaintySubcomponent,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/EffectEvidenceSynthesis
type EffectEvidenceSynthesisCertaintyCertaintySubcomponent struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Rating            []CodeableConcept `json:"rating,omitempty"`
	Note              []Annotation      `json:"note,omitempty"`
}

type OtherEffectEvidenceSynthesis EffectEvidenceSynthesis

// on convert struct to json, automatically add resourceType=EffectEvidenceSynthesis
func (r EffectEvidenceSynthesis) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEffectEvidenceSynthesis
		ResourceType string `json:"resourceType"`
	}{
		OtherEffectEvidenceSynthesis: OtherEffectEvidenceSynthesis(r),
		ResourceType:                 "EffectEvidenceSynthesis",
	})
}
func (r EffectEvidenceSynthesis) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "EffectEvidenceSynthesis/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "EffectEvidenceSynthesis"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *EffectEvidenceSynthesis) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("approvalDate", nil, htmlAttrs)
	}
	return FhirDateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("lastReviewDate", nil, htmlAttrs)
	}
	return FhirDateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectivePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("effectivePeriod", resource.EffectivePeriod, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Author(numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthor >= len(resource.Author) {
		return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", &resource.Author[numAuthor], htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Editor(numEditor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEditor >= len(resource.Editor) {
		return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", &resource.Editor[numEditor], htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Reviewer(numReviewer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReviewer >= len(resource.Reviewer) {
		return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", &resource.Reviewer[numReviewer], htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Endorser(numEndorser int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndorser >= len(resource.Endorser) {
		return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", &resource.Endorser[numEndorser], htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_RelatedArtifact(numRelatedArtifact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedArtifact >= len(resource.RelatedArtifact) {
		return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", &resource.RelatedArtifact[numRelatedArtifact], htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_SynthesisType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("synthesisType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("synthesisType", resource.SynthesisType, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_StudyType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("studyType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("studyType", resource.StudyType, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Population(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "population", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "population", &resource.Population, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Exposure(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "exposure", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "exposure", &resource.Exposure, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_ExposureAlternative(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "exposureAlternative", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "exposureAlternative", &resource.ExposureAlternative, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Outcome(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "outcome", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "outcome", &resource.Outcome, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_SampleSizeDescription(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("sampleSize.description", nil, htmlAttrs)
	}
	return StringInput("sampleSize.description", resource.SampleSize.Description, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_SampleSizeNumberOfStudies(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("sampleSize.numberOfStudies", nil, htmlAttrs)
	}
	return IntInput("sampleSize.numberOfStudies", resource.SampleSize.NumberOfStudies, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_SampleSizeNumberOfParticipants(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("sampleSize.numberOfParticipants", nil, htmlAttrs)
	}
	return IntInput("sampleSize.numberOfParticipants", resource.SampleSize.NumberOfParticipants, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_ResultsByExposureDescription(numResultsByExposure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResultsByExposure >= len(resource.ResultsByExposure) {
		return StringInput("resultsByExposure["+strconv.Itoa(numResultsByExposure)+"].description", nil, htmlAttrs)
	}
	return StringInput("resultsByExposure["+strconv.Itoa(numResultsByExposure)+"].description", resource.ResultsByExposure[numResultsByExposure].Description, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_ResultsByExposureExposureState(numResultsByExposure int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSExposure_state

	if resource == nil || numResultsByExposure >= len(resource.ResultsByExposure) {
		return CodeSelect("resultsByExposure["+strconv.Itoa(numResultsByExposure)+"].exposureState", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("resultsByExposure["+strconv.Itoa(numResultsByExposure)+"].exposureState", resource.ResultsByExposure[numResultsByExposure].ExposureState, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_ResultsByExposureVariantState(numResultsByExposure int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResultsByExposure >= len(resource.ResultsByExposure) {
		return CodeableConceptSelect("resultsByExposure["+strconv.Itoa(numResultsByExposure)+"].variantState", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("resultsByExposure["+strconv.Itoa(numResultsByExposure)+"].variantState", resource.ResultsByExposure[numResultsByExposure].VariantState, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_ResultsByExposureRiskEvidenceSynthesis(frs []FhirResource, numResultsByExposure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResultsByExposure >= len(resource.ResultsByExposure) {
		return ReferenceInput(frs, "resultsByExposure["+strconv.Itoa(numResultsByExposure)+"].riskEvidenceSynthesis", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "resultsByExposure["+strconv.Itoa(numResultsByExposure)+"].riskEvidenceSynthesis", &resource.ResultsByExposure[numResultsByExposure].RiskEvidenceSynthesis, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimateDescription(numEffectEstimate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) {
		return StringInput("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].description", nil, htmlAttrs)
	}
	return StringInput("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].description", resource.EffectEstimate[numEffectEstimate].Description, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimateType(numEffectEstimate int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) {
		return CodeableConceptSelect("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].type", resource.EffectEstimate[numEffectEstimate].Type, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimateVariantState(numEffectEstimate int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) {
		return CodeableConceptSelect("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].variantState", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].variantState", resource.EffectEstimate[numEffectEstimate].VariantState, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimateValue(numEffectEstimate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) {
		return Float64Input("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].value", nil, htmlAttrs)
	}
	return Float64Input("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].value", resource.EffectEstimate[numEffectEstimate].Value, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimateUnitOfMeasure(numEffectEstimate int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) {
		return CodeableConceptSelect("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].unitOfMeasure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].unitOfMeasure", resource.EffectEstimate[numEffectEstimate].UnitOfMeasure, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimatePrecisionEstimateType(numEffectEstimate int, numPrecisionEstimate int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) || numPrecisionEstimate >= len(resource.EffectEstimate[numEffectEstimate].PrecisionEstimate) {
		return CodeableConceptSelect("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].precisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].precisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].type", resource.EffectEstimate[numEffectEstimate].PrecisionEstimate[numPrecisionEstimate].Type, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimatePrecisionEstimateLevel(numEffectEstimate int, numPrecisionEstimate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) || numPrecisionEstimate >= len(resource.EffectEstimate[numEffectEstimate].PrecisionEstimate) {
		return Float64Input("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].precisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].level", nil, htmlAttrs)
	}
	return Float64Input("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].precisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].level", resource.EffectEstimate[numEffectEstimate].PrecisionEstimate[numPrecisionEstimate].Level, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimatePrecisionEstimateFrom(numEffectEstimate int, numPrecisionEstimate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) || numPrecisionEstimate >= len(resource.EffectEstimate[numEffectEstimate].PrecisionEstimate) {
		return Float64Input("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].precisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].from", nil, htmlAttrs)
	}
	return Float64Input("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].precisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].from", resource.EffectEstimate[numEffectEstimate].PrecisionEstimate[numPrecisionEstimate].From, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimatePrecisionEstimateTo(numEffectEstimate int, numPrecisionEstimate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) || numPrecisionEstimate >= len(resource.EffectEstimate[numEffectEstimate].PrecisionEstimate) {
		return Float64Input("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].precisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].to", nil, htmlAttrs)
	}
	return Float64Input("effectEstimate["+strconv.Itoa(numEffectEstimate)+"].precisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].to", resource.EffectEstimate[numEffectEstimate].PrecisionEstimate[numPrecisionEstimate].To, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_CertaintyRating(numCertainty int, numRating int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numRating >= len(resource.Certainty[numCertainty].Rating) {
		return CodeableConceptSelect("certainty["+strconv.Itoa(numCertainty)+"].rating["+strconv.Itoa(numRating)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("certainty["+strconv.Itoa(numCertainty)+"].rating["+strconv.Itoa(numRating)+"]", &resource.Certainty[numCertainty].Rating[numRating], optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_CertaintyNote(numCertainty int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numNote >= len(resource.Certainty[numCertainty].Note) {
		return AnnotationTextArea("certainty["+strconv.Itoa(numCertainty)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("certainty["+strconv.Itoa(numCertainty)+"].note["+strconv.Itoa(numNote)+"]", &resource.Certainty[numCertainty].Note[numNote], htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_CertaintyCertaintySubcomponentType(numCertainty int, numCertaintySubcomponent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numCertaintySubcomponent >= len(resource.Certainty[numCertainty].CertaintySubcomponent) {
		return CodeableConceptSelect("certainty["+strconv.Itoa(numCertainty)+"].certaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("certainty["+strconv.Itoa(numCertainty)+"].certaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].type", resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Type, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_CertaintyCertaintySubcomponentRating(numCertainty int, numCertaintySubcomponent int, numRating int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numCertaintySubcomponent >= len(resource.Certainty[numCertainty].CertaintySubcomponent) || numRating >= len(resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Rating) {
		return CodeableConceptSelect("certainty["+strconv.Itoa(numCertainty)+"].certaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].rating["+strconv.Itoa(numRating)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("certainty["+strconv.Itoa(numCertainty)+"].certaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].rating["+strconv.Itoa(numRating)+"]", &resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Rating[numRating], optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_CertaintyCertaintySubcomponentNote(numCertainty int, numCertaintySubcomponent int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numCertaintySubcomponent >= len(resource.Certainty[numCertainty].CertaintySubcomponent) || numNote >= len(resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Note) {
		return AnnotationTextArea("certainty["+strconv.Itoa(numCertainty)+"].certaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("certainty["+strconv.Itoa(numCertainty)+"].certaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].note["+strconv.Itoa(numNote)+"]", &resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Note[numNote], htmlAttrs)
}
