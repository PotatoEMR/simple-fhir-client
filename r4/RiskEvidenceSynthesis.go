package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Date              *time.Time                         `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher         *string                            `json:"publisher,omitempty"`
	Contact           []ContactDetail                    `json:"contact,omitempty"`
	Description       *string                            `json:"description,omitempty"`
	Note              []Annotation                       `json:"note,omitempty"`
	UseContext        []UsageContext                     `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                  `json:"jurisdiction,omitempty"`
	Copyright         *string                            `json:"copyright,omitempty"`
	ApprovalDate      *time.Time                         `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate    *time.Time                         `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
func (resource *RiskEvidenceSynthesis) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Copyright", nil, htmlAttrs)
	}
	return StringInput("Copyright", resource.Copyright, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_SynthesisType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SynthesisType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SynthesisType", resource.SynthesisType, optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_StudyType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("StudyType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("StudyType", resource.StudyType, optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_SampleSizeDescription(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SampleSizeDescription", nil, htmlAttrs)
	}
	return StringInput("SampleSizeDescription", resource.SampleSize.Description, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_SampleSizeNumberOfStudies(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("SampleSizeNumberOfStudies", nil, htmlAttrs)
	}
	return IntInput("SampleSizeNumberOfStudies", resource.SampleSize.NumberOfStudies, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_SampleSizeNumberOfParticipants(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("SampleSizeNumberOfParticipants", nil, htmlAttrs)
	}
	return IntInput("SampleSizeNumberOfParticipants", resource.SampleSize.NumberOfParticipants, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateDescription(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("RiskEstimateDescription", nil, htmlAttrs)
	}
	return StringInput("RiskEstimateDescription", resource.RiskEstimate.Description, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("RiskEstimateType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("RiskEstimateType", resource.RiskEstimate.Type, optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateValue(htmlAttrs string) templ.Component {
	if resource == nil {
		return Float64Input("RiskEstimateValue", nil, htmlAttrs)
	}
	return Float64Input("RiskEstimateValue", resource.RiskEstimate.Value, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateUnitOfMeasure(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("RiskEstimateUnitOfMeasure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("RiskEstimateUnitOfMeasure", resource.RiskEstimate.UnitOfMeasure, optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateDenominatorCount(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("RiskEstimateDenominatorCount", nil, htmlAttrs)
	}
	return IntInput("RiskEstimateDenominatorCount", resource.RiskEstimate.DenominatorCount, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateNumeratorCount(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("RiskEstimateNumeratorCount", nil, htmlAttrs)
	}
	return IntInput("RiskEstimateNumeratorCount", resource.RiskEstimate.NumeratorCount, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimatePrecisionEstimateType(numPrecisionEstimate int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPrecisionEstimate >= len(resource.RiskEstimate.PrecisionEstimate) {
		return CodeableConceptSelect("RiskEstimatePrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("RiskEstimatePrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Type", resource.RiskEstimate.PrecisionEstimate[numPrecisionEstimate].Type, optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimatePrecisionEstimateLevel(numPrecisionEstimate int, htmlAttrs string) templ.Component {
	if resource == nil || numPrecisionEstimate >= len(resource.RiskEstimate.PrecisionEstimate) {
		return Float64Input("RiskEstimatePrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Level", nil, htmlAttrs)
	}
	return Float64Input("RiskEstimatePrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Level", resource.RiskEstimate.PrecisionEstimate[numPrecisionEstimate].Level, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimatePrecisionEstimateFrom(numPrecisionEstimate int, htmlAttrs string) templ.Component {
	if resource == nil || numPrecisionEstimate >= len(resource.RiskEstimate.PrecisionEstimate) {
		return Float64Input("RiskEstimatePrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].From", nil, htmlAttrs)
	}
	return Float64Input("RiskEstimatePrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].From", resource.RiskEstimate.PrecisionEstimate[numPrecisionEstimate].From, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimatePrecisionEstimateTo(numPrecisionEstimate int, htmlAttrs string) templ.Component {
	if resource == nil || numPrecisionEstimate >= len(resource.RiskEstimate.PrecisionEstimate) {
		return Float64Input("RiskEstimatePrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].To", nil, htmlAttrs)
	}
	return Float64Input("RiskEstimatePrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].To", resource.RiskEstimate.PrecisionEstimate[numPrecisionEstimate].To, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyRating(numCertainty int, numRating int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numRating >= len(resource.Certainty[numCertainty].Rating) {
		return CodeableConceptSelect("Certainty["+strconv.Itoa(numCertainty)+"]Rating["+strconv.Itoa(numRating)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Certainty["+strconv.Itoa(numCertainty)+"]Rating["+strconv.Itoa(numRating)+"]", &resource.Certainty[numCertainty].Rating[numRating], optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyNote(numCertainty int, numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numNote >= len(resource.Certainty[numCertainty].Note) {
		return AnnotationTextArea("Certainty["+strconv.Itoa(numCertainty)+"]Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Certainty["+strconv.Itoa(numCertainty)+"]Note["+strconv.Itoa(numNote)+"]", &resource.Certainty[numCertainty].Note[numNote], htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyCertaintySubcomponentType(numCertainty int, numCertaintySubcomponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numCertaintySubcomponent >= len(resource.Certainty[numCertainty].CertaintySubcomponent) {
		return CodeableConceptSelect("Certainty["+strconv.Itoa(numCertainty)+"]CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Certainty["+strconv.Itoa(numCertainty)+"]CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Type", resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Type, optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyCertaintySubcomponentRating(numCertainty int, numCertaintySubcomponent int, numRating int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numCertaintySubcomponent >= len(resource.Certainty[numCertainty].CertaintySubcomponent) || numRating >= len(resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Rating) {
		return CodeableConceptSelect("Certainty["+strconv.Itoa(numCertainty)+"]CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Rating["+strconv.Itoa(numRating)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Certainty["+strconv.Itoa(numCertainty)+"]CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Rating["+strconv.Itoa(numRating)+"]", &resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Rating[numRating], optionsValueSet, htmlAttrs)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyCertaintySubcomponentNote(numCertainty int, numCertaintySubcomponent int, numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numCertaintySubcomponent >= len(resource.Certainty[numCertainty].CertaintySubcomponent) || numNote >= len(resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Note) {
		return AnnotationTextArea("Certainty["+strconv.Itoa(numCertainty)+"]CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Certainty["+strconv.Itoa(numCertainty)+"]CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Note["+strconv.Itoa(numNote)+"]", &resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Note[numNote], htmlAttrs)
}
