package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/RiskEvidenceSynthesis
type RiskEvidenceSynthesis struct {
	Id                *string                            `json:"id,omitempty"`
	Meta              *Meta                              `json:"meta,omitempty"`
	ImplicitRules     *string                            `json:"implicitRules,omitempty"`
	Language          *string                            `json:"language,omitempty"`
	Text              *Narrative                         `json:"text,omitempty"`
	Contained         []Resource                         `json:"contained,omitempty"`
	Extension         []Extension                        `json:"extension,omitempty"`
	ModifierExtension []Extension                        `json:"modifierExtension,omitempty"`
	Url               *string                            `json:"url,omitempty"`
	Identifier        []Identifier                       `json:"identifier,omitempty"`
	Version           *string                            `json:"version,omitempty"`
	Name              *string                            `json:"name,omitempty"`
	Title             *string                            `json:"title,omitempty"`
	Status            string                             `json:"status"`
	Date              *FhirDateTime                      `json:"date,omitempty"`
	Publisher         *string                            `json:"publisher,omitempty"`
	Contact           []ContactDetail                    `json:"contact,omitempty"`
	Description       *string                            `json:"description,omitempty"`
	Note              []Annotation                       `json:"note,omitempty"`
	UseContext        []UsageContext                     `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                  `json:"jurisdiction,omitempty"`
	Copyright         *string                            `json:"copyright,omitempty"`
	ApprovalDate      *FhirDate                          `json:"approvalDate,omitempty"`
	LastReviewDate    *FhirDate                          `json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period                            `json:"effectivePeriod,omitempty"`
	Topic             []CodeableConcept                  `json:"topic,omitempty"`
	Author            []ContactDetail                    `json:"author,omitempty"`
	Editor            []ContactDetail                    `json:"editor,omitempty"`
	Reviewer          []ContactDetail                    `json:"reviewer,omitempty"`
	Endorser          []ContactDetail                    `json:"endorser,omitempty"`
	RelatedArtifact   []RelatedArtifact                  `json:"relatedArtifact,omitempty"`
	SynthesisType     *CodeableConcept                   `json:"synthesisType,omitempty"`
	StudyType         *CodeableConcept                   `json:"studyType,omitempty"`
	Population        Reference                          `json:"population"`
	Exposure          *Reference                         `json:"exposure,omitempty"`
	Outcome           Reference                          `json:"outcome"`
	SampleSize        *RiskEvidenceSynthesisSampleSize   `json:"sampleSize,omitempty"`
	RiskEstimate      *RiskEvidenceSynthesisRiskEstimate `json:"riskEstimate,omitempty"`
	Certainty         []RiskEvidenceSynthesisCertainty   `json:"certainty,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/RiskEvidenceSynthesis
type RiskEvidenceSynthesisSampleSize struct {
	Id                   *string     `json:"id,omitempty"`
	Extension            []Extension `json:"extension,omitempty"`
	ModifierExtension    []Extension `json:"modifierExtension,omitempty"`
	Description          *string     `json:"description,omitempty"`
	NumberOfStudies      *int        `json:"numberOfStudies,omitempty"`
	NumberOfParticipants *int        `json:"numberOfParticipants,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/RiskEvidenceSynthesis
type RiskEvidenceSynthesisRiskEstimate struct {
	Id                *string                                              `json:"id,omitempty"`
	Extension         []Extension                                          `json:"extension,omitempty"`
	ModifierExtension []Extension                                          `json:"modifierExtension,omitempty"`
	Description       *string                                              `json:"description,omitempty"`
	Type              *CodeableConcept                                     `json:"type,omitempty"`
	Value             *float64                                             `json:"value,omitempty"`
	UnitOfMeasure     *CodeableConcept                                     `json:"unitOfMeasure,omitempty"`
	DenominatorCount  *int                                                 `json:"denominatorCount,omitempty"`
	NumeratorCount    *int                                                 `json:"numeratorCount,omitempty"`
	PrecisionEstimate []RiskEvidenceSynthesisRiskEstimatePrecisionEstimate `json:"precisionEstimate,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/RiskEvidenceSynthesis
type RiskEvidenceSynthesisRiskEstimatePrecisionEstimate struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Level             *float64         `json:"level,omitempty"`
	From              *float64         `json:"from,omitempty"`
	To                *float64         `json:"to,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/RiskEvidenceSynthesis
type RiskEvidenceSynthesisCertainty struct {
	Id                    *string                                               `json:"id,omitempty"`
	Extension             []Extension                                           `json:"extension,omitempty"`
	ModifierExtension     []Extension                                           `json:"modifierExtension,omitempty"`
	Rating                []CodeableConcept                                     `json:"rating,omitempty"`
	Note                  []Annotation                                          `json:"note,omitempty"`
	CertaintySubcomponent []RiskEvidenceSynthesisCertaintyCertaintySubcomponent `json:"certaintySubcomponent,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/RiskEvidenceSynthesis
type RiskEvidenceSynthesisCertaintyCertaintySubcomponent struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Rating            []CodeableConcept `json:"rating,omitempty"`
	Note              []Annotation      `json:"note,omitempty"`
}

type OtherRiskEvidenceSynthesis RiskEvidenceSynthesis

// on convert struct to json, automatically add resourceType=RiskEvidenceSynthesis
func (r RiskEvidenceSynthesis) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRiskEvidenceSynthesis
		ResourceType string `json:"resourceType"`
	}{
		OtherRiskEvidenceSynthesis: OtherRiskEvidenceSynthesis(r),
		ResourceType:               "RiskEvidenceSynthesis",
	})
}
func (r RiskEvidenceSynthesis) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "RiskEvidenceSynthesis/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "RiskEvidenceSynthesis"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *RiskEvidenceSynthesis) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateInput("approvalDate", nil, htmlAttrs)
	}
	return DateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateInput("lastReviewDate", nil, htmlAttrs)
	}
	return DateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_SynthesisType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("synthesisType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("synthesisType", resource.SynthesisType, optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_StudyType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("studyType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("studyType", resource.StudyType, optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_SampleSizeDescription(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("sampleSize.description", nil, htmlAttrs)
	}
	return StringInput("sampleSize.description", resource.SampleSize.Description, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_SampleSizeNumberOfStudies(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("sampleSize.numberOfStudies", nil, htmlAttrs)
	}
	return IntInput("sampleSize.numberOfStudies", resource.SampleSize.NumberOfStudies, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_SampleSizeNumberOfParticipants(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("sampleSize.numberOfParticipants", nil, htmlAttrs)
	}
	return IntInput("sampleSize.numberOfParticipants", resource.SampleSize.NumberOfParticipants, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateDescription(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("riskEstimate.description", nil, htmlAttrs)
	}
	return StringInput("riskEstimate.description", resource.RiskEstimate.Description, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("riskEstimate.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("riskEstimate.type", resource.RiskEstimate.Type, optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateValue(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return Float64Input("riskEstimate.value", nil, htmlAttrs)
	}
	return Float64Input("riskEstimate.value", resource.RiskEstimate.Value, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateUnitOfMeasure(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("riskEstimate.unitOfMeasure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("riskEstimate.unitOfMeasure", resource.RiskEstimate.UnitOfMeasure, optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateDenominatorCount(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("riskEstimate.denominatorCount", nil, htmlAttrs)
	}
	return IntInput("riskEstimate.denominatorCount", resource.RiskEstimate.DenominatorCount, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateNumeratorCount(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("riskEstimate.numeratorCount", nil, htmlAttrs)
	}
	return IntInput("riskEstimate.numeratorCount", resource.RiskEstimate.NumeratorCount, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimatePrecisionEstimateType(numPrecisionEstimate int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrecisionEstimate >= len(resource.RiskEstimate.PrecisionEstimate) {
		return CodeableConceptSelect("riskEstimate.precisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("riskEstimate.precisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].type", resource.RiskEstimate.PrecisionEstimate[numPrecisionEstimate].Type, optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimatePrecisionEstimateLevel(numPrecisionEstimate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrecisionEstimate >= len(resource.RiskEstimate.PrecisionEstimate) {
		return Float64Input("riskEstimate.precisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].level", nil, htmlAttrs)
	}
	return Float64Input("riskEstimate.precisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].level", resource.RiskEstimate.PrecisionEstimate[numPrecisionEstimate].Level, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimatePrecisionEstimateFrom(numPrecisionEstimate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrecisionEstimate >= len(resource.RiskEstimate.PrecisionEstimate) {
		return Float64Input("riskEstimate.precisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].from", nil, htmlAttrs)
	}
	return Float64Input("riskEstimate.precisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].from", resource.RiskEstimate.PrecisionEstimate[numPrecisionEstimate].From, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimatePrecisionEstimateTo(numPrecisionEstimate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrecisionEstimate >= len(resource.RiskEstimate.PrecisionEstimate) {
		return Float64Input("riskEstimate.precisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].to", nil, htmlAttrs)
	}
	return Float64Input("riskEstimate.precisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].to", resource.RiskEstimate.PrecisionEstimate[numPrecisionEstimate].To, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyRating(numCertainty int, numRating int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numRating >= len(resource.Certainty[numCertainty].Rating) {
		return CodeableConceptSelect("certainty["+strconv.Itoa(numCertainty)+"].rating["+strconv.Itoa(numRating)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("certainty["+strconv.Itoa(numCertainty)+"].rating["+strconv.Itoa(numRating)+"]", &resource.Certainty[numCertainty].Rating[numRating], optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyNote(numCertainty int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numNote >= len(resource.Certainty[numCertainty].Note) {
		return AnnotationTextArea("certainty["+strconv.Itoa(numCertainty)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("certainty["+strconv.Itoa(numCertainty)+"].note["+strconv.Itoa(numNote)+"]", &resource.Certainty[numCertainty].Note[numNote], htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyCertaintySubcomponentType(numCertainty int, numCertaintySubcomponent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numCertaintySubcomponent >= len(resource.Certainty[numCertainty].CertaintySubcomponent) {
		return CodeableConceptSelect("certainty["+strconv.Itoa(numCertainty)+"].certaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("certainty["+strconv.Itoa(numCertainty)+"].certaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].type", resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Type, optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyCertaintySubcomponentRating(numCertainty int, numCertaintySubcomponent int, numRating int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numCertaintySubcomponent >= len(resource.Certainty[numCertainty].CertaintySubcomponent) || numRating >= len(resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Rating) {
		return CodeableConceptSelect("certainty["+strconv.Itoa(numCertainty)+"].certaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].rating["+strconv.Itoa(numRating)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("certainty["+strconv.Itoa(numCertainty)+"].certaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].rating["+strconv.Itoa(numRating)+"]", &resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Rating[numRating], optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyCertaintySubcomponentNote(numCertainty int, numCertaintySubcomponent int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numCertaintySubcomponent >= len(resource.Certainty[numCertainty].CertaintySubcomponent) || numNote >= len(resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Note) {
		return AnnotationTextArea("certainty["+strconv.Itoa(numCertainty)+"].certaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("certainty["+strconv.Itoa(numCertainty)+"].certaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].note["+strconv.Itoa(numNote)+"]", &resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Note[numNote], htmlAttrs)
}
