package r4

//generated with command go run ./bultaoreune -nodownload
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
	Date              *string                            `json:"date,omitempty"`
	Publisher         *string                            `json:"publisher,omitempty"`
	Contact           []ContactDetail                    `json:"contact,omitempty"`
	Description       *string                            `json:"description,omitempty"`
	Note              []Annotation                       `json:"note,omitempty"`
	UseContext        []UsageContext                     `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                  `json:"jurisdiction,omitempty"`
	Copyright         *string                            `json:"copyright,omitempty"`
	ApprovalDate      *string                            `json:"approvalDate,omitempty"`
	LastReviewDate    *string                            `json:"lastReviewDate,omitempty"`
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

func (resource *RiskEvidenceSynthesis) T_Id() templ.Component {

	if resource == nil {
		return StringInput("RiskEvidenceSynthesis.Id", nil)
	}
	return StringInput("RiskEvidenceSynthesis.Id", resource.Id)
}
func (resource *RiskEvidenceSynthesis) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("RiskEvidenceSynthesis.ImplicitRules", nil)
	}
	return StringInput("RiskEvidenceSynthesis.ImplicitRules", resource.ImplicitRules)
}
func (resource *RiskEvidenceSynthesis) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("RiskEvidenceSynthesis.Language", nil, optionsValueSet)
	}
	return CodeSelect("RiskEvidenceSynthesis.Language", resource.Language, optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_Url() templ.Component {

	if resource == nil {
		return StringInput("RiskEvidenceSynthesis.Url", nil)
	}
	return StringInput("RiskEvidenceSynthesis.Url", resource.Url)
}
func (resource *RiskEvidenceSynthesis) T_Version() templ.Component {

	if resource == nil {
		return StringInput("RiskEvidenceSynthesis.Version", nil)
	}
	return StringInput("RiskEvidenceSynthesis.Version", resource.Version)
}
func (resource *RiskEvidenceSynthesis) T_Name() templ.Component {

	if resource == nil {
		return StringInput("RiskEvidenceSynthesis.Name", nil)
	}
	return StringInput("RiskEvidenceSynthesis.Name", resource.Name)
}
func (resource *RiskEvidenceSynthesis) T_Title() templ.Component {

	if resource == nil {
		return StringInput("RiskEvidenceSynthesis.Title", nil)
	}
	return StringInput("RiskEvidenceSynthesis.Title", resource.Title)
}
func (resource *RiskEvidenceSynthesis) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("RiskEvidenceSynthesis.Status", nil, optionsValueSet)
	}
	return CodeSelect("RiskEvidenceSynthesis.Status", &resource.Status, optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_Date() templ.Component {

	if resource == nil {
		return StringInput("RiskEvidenceSynthesis.Date", nil)
	}
	return StringInput("RiskEvidenceSynthesis.Date", resource.Date)
}
func (resource *RiskEvidenceSynthesis) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("RiskEvidenceSynthesis.Publisher", nil)
	}
	return StringInput("RiskEvidenceSynthesis.Publisher", resource.Publisher)
}
func (resource *RiskEvidenceSynthesis) T_Description() templ.Component {

	if resource == nil {
		return StringInput("RiskEvidenceSynthesis.Description", nil)
	}
	return StringInput("RiskEvidenceSynthesis.Description", resource.Description)
}
func (resource *RiskEvidenceSynthesis) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("RiskEvidenceSynthesis.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RiskEvidenceSynthesis.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("RiskEvidenceSynthesis.Copyright", nil)
	}
	return StringInput("RiskEvidenceSynthesis.Copyright", resource.Copyright)
}
func (resource *RiskEvidenceSynthesis) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("RiskEvidenceSynthesis.ApprovalDate", nil)
	}
	return StringInput("RiskEvidenceSynthesis.ApprovalDate", resource.ApprovalDate)
}
func (resource *RiskEvidenceSynthesis) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("RiskEvidenceSynthesis.LastReviewDate", nil)
	}
	return StringInput("RiskEvidenceSynthesis.LastReviewDate", resource.LastReviewDate)
}
func (resource *RiskEvidenceSynthesis) T_Topic(numTopic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Topic) >= numTopic {
		return CodeableConceptSelect("RiskEvidenceSynthesis.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RiskEvidenceSynthesis.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_SynthesisType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("RiskEvidenceSynthesis.SynthesisType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RiskEvidenceSynthesis.SynthesisType", resource.SynthesisType, optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_StudyType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("RiskEvidenceSynthesis.StudyType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RiskEvidenceSynthesis.StudyType", resource.StudyType, optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_SampleSizeId() templ.Component {

	if resource == nil {
		return StringInput("RiskEvidenceSynthesis.SampleSize.Id", nil)
	}
	return StringInput("RiskEvidenceSynthesis.SampleSize.Id", resource.SampleSize.Id)
}
func (resource *RiskEvidenceSynthesis) T_SampleSizeDescription() templ.Component {

	if resource == nil {
		return StringInput("RiskEvidenceSynthesis.SampleSize.Description", nil)
	}
	return StringInput("RiskEvidenceSynthesis.SampleSize.Description", resource.SampleSize.Description)
}
func (resource *RiskEvidenceSynthesis) T_SampleSizeNumberOfStudies() templ.Component {

	if resource == nil {
		return IntInput("RiskEvidenceSynthesis.SampleSize.NumberOfStudies", nil)
	}
	return IntInput("RiskEvidenceSynthesis.SampleSize.NumberOfStudies", resource.SampleSize.NumberOfStudies)
}
func (resource *RiskEvidenceSynthesis) T_SampleSizeNumberOfParticipants() templ.Component {

	if resource == nil {
		return IntInput("RiskEvidenceSynthesis.SampleSize.NumberOfParticipants", nil)
	}
	return IntInput("RiskEvidenceSynthesis.SampleSize.NumberOfParticipants", resource.SampleSize.NumberOfParticipants)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateId() templ.Component {

	if resource == nil {
		return StringInput("RiskEvidenceSynthesis.RiskEstimate.Id", nil)
	}
	return StringInput("RiskEvidenceSynthesis.RiskEstimate.Id", resource.RiskEstimate.Id)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateDescription() templ.Component {

	if resource == nil {
		return StringInput("RiskEvidenceSynthesis.RiskEstimate.Description", nil)
	}
	return StringInput("RiskEvidenceSynthesis.RiskEstimate.Description", resource.RiskEstimate.Description)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("RiskEvidenceSynthesis.RiskEstimate.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RiskEvidenceSynthesis.RiskEstimate.Type", resource.RiskEstimate.Type, optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateValue() templ.Component {

	if resource == nil {
		return Float64Input("RiskEvidenceSynthesis.RiskEstimate.Value", nil)
	}
	return Float64Input("RiskEvidenceSynthesis.RiskEstimate.Value", resource.RiskEstimate.Value)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateUnitOfMeasure(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("RiskEvidenceSynthesis.RiskEstimate.UnitOfMeasure", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RiskEvidenceSynthesis.RiskEstimate.UnitOfMeasure", resource.RiskEstimate.UnitOfMeasure, optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateDenominatorCount() templ.Component {

	if resource == nil {
		return IntInput("RiskEvidenceSynthesis.RiskEstimate.DenominatorCount", nil)
	}
	return IntInput("RiskEvidenceSynthesis.RiskEstimate.DenominatorCount", resource.RiskEstimate.DenominatorCount)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateNumeratorCount() templ.Component {

	if resource == nil {
		return IntInput("RiskEvidenceSynthesis.RiskEstimate.NumeratorCount", nil)
	}
	return IntInput("RiskEvidenceSynthesis.RiskEstimate.NumeratorCount", resource.RiskEstimate.NumeratorCount)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimatePrecisionEstimateId(numPrecisionEstimate int) templ.Component {

	if resource == nil || len(resource.RiskEstimate.PrecisionEstimate) >= numPrecisionEstimate {
		return StringInput("RiskEvidenceSynthesis.RiskEstimate.PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Id", nil)
	}
	return StringInput("RiskEvidenceSynthesis.RiskEstimate.PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Id", resource.RiskEstimate.PrecisionEstimate[numPrecisionEstimate].Id)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimatePrecisionEstimateType(numPrecisionEstimate int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.RiskEstimate.PrecisionEstimate) >= numPrecisionEstimate {
		return CodeableConceptSelect("RiskEvidenceSynthesis.RiskEstimate.PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RiskEvidenceSynthesis.RiskEstimate.PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Type", resource.RiskEstimate.PrecisionEstimate[numPrecisionEstimate].Type, optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimatePrecisionEstimateLevel(numPrecisionEstimate int) templ.Component {

	if resource == nil || len(resource.RiskEstimate.PrecisionEstimate) >= numPrecisionEstimate {
		return Float64Input("RiskEvidenceSynthesis.RiskEstimate.PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Level", nil)
	}
	return Float64Input("RiskEvidenceSynthesis.RiskEstimate.PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Level", resource.RiskEstimate.PrecisionEstimate[numPrecisionEstimate].Level)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimatePrecisionEstimateFrom(numPrecisionEstimate int) templ.Component {

	if resource == nil || len(resource.RiskEstimate.PrecisionEstimate) >= numPrecisionEstimate {
		return Float64Input("RiskEvidenceSynthesis.RiskEstimate.PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].From", nil)
	}
	return Float64Input("RiskEvidenceSynthesis.RiskEstimate.PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].From", resource.RiskEstimate.PrecisionEstimate[numPrecisionEstimate].From)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimatePrecisionEstimateTo(numPrecisionEstimate int) templ.Component {

	if resource == nil || len(resource.RiskEstimate.PrecisionEstimate) >= numPrecisionEstimate {
		return Float64Input("RiskEvidenceSynthesis.RiskEstimate.PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].To", nil)
	}
	return Float64Input("RiskEvidenceSynthesis.RiskEstimate.PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].To", resource.RiskEstimate.PrecisionEstimate[numPrecisionEstimate].To)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyId(numCertainty int) templ.Component {

	if resource == nil || len(resource.Certainty) >= numCertainty {
		return StringInput("RiskEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].Id", nil)
	}
	return StringInput("RiskEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].Id", resource.Certainty[numCertainty].Id)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyRating(numCertainty int, numRating int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Certainty) >= numCertainty || len(resource.Certainty[numCertainty].Rating) >= numRating {
		return CodeableConceptSelect("RiskEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].Rating["+strconv.Itoa(numRating)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RiskEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].Rating["+strconv.Itoa(numRating)+"]", &resource.Certainty[numCertainty].Rating[numRating], optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyCertaintySubcomponentId(numCertainty int, numCertaintySubcomponent int) templ.Component {

	if resource == nil || len(resource.Certainty) >= numCertainty || len(resource.Certainty[numCertainty].CertaintySubcomponent) >= numCertaintySubcomponent {
		return StringInput("RiskEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Id", nil)
	}
	return StringInput("RiskEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Id", resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Id)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyCertaintySubcomponentType(numCertainty int, numCertaintySubcomponent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Certainty) >= numCertainty || len(resource.Certainty[numCertainty].CertaintySubcomponent) >= numCertaintySubcomponent {
		return CodeableConceptSelect("RiskEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RiskEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Type", resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Type, optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyCertaintySubcomponentRating(numCertainty int, numCertaintySubcomponent int, numRating int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Certainty) >= numCertainty || len(resource.Certainty[numCertainty].CertaintySubcomponent) >= numCertaintySubcomponent || len(resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Rating) >= numRating {
		return CodeableConceptSelect("RiskEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Rating["+strconv.Itoa(numRating)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("RiskEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Rating["+strconv.Itoa(numRating)+"]", &resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Rating[numRating], optionsValueSet)
}
