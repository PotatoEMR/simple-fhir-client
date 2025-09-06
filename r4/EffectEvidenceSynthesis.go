package r4

//generated with command go run ./bultaoreune -nodownload
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
	Date                *string                                    `json:"date,omitempty"`
	Publisher           *string                                    `json:"publisher,omitempty"`
	Contact             []ContactDetail                            `json:"contact,omitempty"`
	Description         *string                                    `json:"description,omitempty"`
	Note                []Annotation                               `json:"note,omitempty"`
	UseContext          []UsageContext                             `json:"useContext,omitempty"`
	Jurisdiction        []CodeableConcept                          `json:"jurisdiction,omitempty"`
	Copyright           *string                                    `json:"copyright,omitempty"`
	ApprovalDate        *string                                    `json:"approvalDate,omitempty"`
	LastReviewDate      *string                                    `json:"lastReviewDate,omitempty"`
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

func (resource *EffectEvidenceSynthesis) T_Id() templ.Component {

	if resource == nil {
		return StringInput("EffectEvidenceSynthesis.Id", nil)
	}
	return StringInput("EffectEvidenceSynthesis.Id", resource.Id)
}
func (resource *EffectEvidenceSynthesis) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("EffectEvidenceSynthesis.ImplicitRules", nil)
	}
	return StringInput("EffectEvidenceSynthesis.ImplicitRules", resource.ImplicitRules)
}
func (resource *EffectEvidenceSynthesis) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("EffectEvidenceSynthesis.Language", nil, optionsValueSet)
	}
	return CodeSelect("EffectEvidenceSynthesis.Language", resource.Language, optionsValueSet)
}
func (resource *EffectEvidenceSynthesis) T_Url() templ.Component {

	if resource == nil {
		return StringInput("EffectEvidenceSynthesis.Url", nil)
	}
	return StringInput("EffectEvidenceSynthesis.Url", resource.Url)
}
func (resource *EffectEvidenceSynthesis) T_Version() templ.Component {

	if resource == nil {
		return StringInput("EffectEvidenceSynthesis.Version", nil)
	}
	return StringInput("EffectEvidenceSynthesis.Version", resource.Version)
}
func (resource *EffectEvidenceSynthesis) T_Name() templ.Component {

	if resource == nil {
		return StringInput("EffectEvidenceSynthesis.Name", nil)
	}
	return StringInput("EffectEvidenceSynthesis.Name", resource.Name)
}
func (resource *EffectEvidenceSynthesis) T_Title() templ.Component {

	if resource == nil {
		return StringInput("EffectEvidenceSynthesis.Title", nil)
	}
	return StringInput("EffectEvidenceSynthesis.Title", resource.Title)
}
func (resource *EffectEvidenceSynthesis) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("EffectEvidenceSynthesis.Status", nil, optionsValueSet)
	}
	return CodeSelect("EffectEvidenceSynthesis.Status", &resource.Status, optionsValueSet)
}
func (resource *EffectEvidenceSynthesis) T_Date() templ.Component {

	if resource == nil {
		return StringInput("EffectEvidenceSynthesis.Date", nil)
	}
	return StringInput("EffectEvidenceSynthesis.Date", resource.Date)
}
func (resource *EffectEvidenceSynthesis) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("EffectEvidenceSynthesis.Publisher", nil)
	}
	return StringInput("EffectEvidenceSynthesis.Publisher", resource.Publisher)
}
func (resource *EffectEvidenceSynthesis) T_Description() templ.Component {

	if resource == nil {
		return StringInput("EffectEvidenceSynthesis.Description", nil)
	}
	return StringInput("EffectEvidenceSynthesis.Description", resource.Description)
}
func (resource *EffectEvidenceSynthesis) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("EffectEvidenceSynthesis.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EffectEvidenceSynthesis.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *EffectEvidenceSynthesis) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("EffectEvidenceSynthesis.Copyright", nil)
	}
	return StringInput("EffectEvidenceSynthesis.Copyright", resource.Copyright)
}
func (resource *EffectEvidenceSynthesis) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("EffectEvidenceSynthesis.ApprovalDate", nil)
	}
	return StringInput("EffectEvidenceSynthesis.ApprovalDate", resource.ApprovalDate)
}
func (resource *EffectEvidenceSynthesis) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("EffectEvidenceSynthesis.LastReviewDate", nil)
	}
	return StringInput("EffectEvidenceSynthesis.LastReviewDate", resource.LastReviewDate)
}
func (resource *EffectEvidenceSynthesis) T_Topic(numTopic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Topic) >= numTopic {
		return CodeableConceptSelect("EffectEvidenceSynthesis.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EffectEvidenceSynthesis.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet)
}
func (resource *EffectEvidenceSynthesis) T_SynthesisType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("EffectEvidenceSynthesis.SynthesisType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EffectEvidenceSynthesis.SynthesisType", resource.SynthesisType, optionsValueSet)
}
func (resource *EffectEvidenceSynthesis) T_StudyType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("EffectEvidenceSynthesis.StudyType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EffectEvidenceSynthesis.StudyType", resource.StudyType, optionsValueSet)
}
func (resource *EffectEvidenceSynthesis) T_SampleSizeId() templ.Component {

	if resource == nil {
		return StringInput("EffectEvidenceSynthesis.SampleSize.Id", nil)
	}
	return StringInput("EffectEvidenceSynthesis.SampleSize.Id", resource.SampleSize.Id)
}
func (resource *EffectEvidenceSynthesis) T_SampleSizeDescription() templ.Component {

	if resource == nil {
		return StringInput("EffectEvidenceSynthesis.SampleSize.Description", nil)
	}
	return StringInput("EffectEvidenceSynthesis.SampleSize.Description", resource.SampleSize.Description)
}
func (resource *EffectEvidenceSynthesis) T_SampleSizeNumberOfStudies() templ.Component {

	if resource == nil {
		return IntInput("EffectEvidenceSynthesis.SampleSize.NumberOfStudies", nil)
	}
	return IntInput("EffectEvidenceSynthesis.SampleSize.NumberOfStudies", resource.SampleSize.NumberOfStudies)
}
func (resource *EffectEvidenceSynthesis) T_SampleSizeNumberOfParticipants() templ.Component {

	if resource == nil {
		return IntInput("EffectEvidenceSynthesis.SampleSize.NumberOfParticipants", nil)
	}
	return IntInput("EffectEvidenceSynthesis.SampleSize.NumberOfParticipants", resource.SampleSize.NumberOfParticipants)
}
func (resource *EffectEvidenceSynthesis) T_ResultsByExposureId(numResultsByExposure int) templ.Component {

	if resource == nil || len(resource.ResultsByExposure) >= numResultsByExposure {
		return StringInput("EffectEvidenceSynthesis.ResultsByExposure["+strconv.Itoa(numResultsByExposure)+"].Id", nil)
	}
	return StringInput("EffectEvidenceSynthesis.ResultsByExposure["+strconv.Itoa(numResultsByExposure)+"].Id", resource.ResultsByExposure[numResultsByExposure].Id)
}
func (resource *EffectEvidenceSynthesis) T_ResultsByExposureDescription(numResultsByExposure int) templ.Component {

	if resource == nil || len(resource.ResultsByExposure) >= numResultsByExposure {
		return StringInput("EffectEvidenceSynthesis.ResultsByExposure["+strconv.Itoa(numResultsByExposure)+"].Description", nil)
	}
	return StringInput("EffectEvidenceSynthesis.ResultsByExposure["+strconv.Itoa(numResultsByExposure)+"].Description", resource.ResultsByExposure[numResultsByExposure].Description)
}
func (resource *EffectEvidenceSynthesis) T_ResultsByExposureExposureState(numResultsByExposure int) templ.Component {
	optionsValueSet := VSExposure_state

	if resource == nil || len(resource.ResultsByExposure) >= numResultsByExposure {
		return CodeSelect("EffectEvidenceSynthesis.ResultsByExposure["+strconv.Itoa(numResultsByExposure)+"].ExposureState", nil, optionsValueSet)
	}
	return CodeSelect("EffectEvidenceSynthesis.ResultsByExposure["+strconv.Itoa(numResultsByExposure)+"].ExposureState", resource.ResultsByExposure[numResultsByExposure].ExposureState, optionsValueSet)
}
func (resource *EffectEvidenceSynthesis) T_ResultsByExposureVariantState(numResultsByExposure int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ResultsByExposure) >= numResultsByExposure {
		return CodeableConceptSelect("EffectEvidenceSynthesis.ResultsByExposure["+strconv.Itoa(numResultsByExposure)+"].VariantState", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EffectEvidenceSynthesis.ResultsByExposure["+strconv.Itoa(numResultsByExposure)+"].VariantState", resource.ResultsByExposure[numResultsByExposure].VariantState, optionsValueSet)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimateId(numEffectEstimate int) templ.Component {

	if resource == nil || len(resource.EffectEstimate) >= numEffectEstimate {
		return StringInput("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].Id", nil)
	}
	return StringInput("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].Id", resource.EffectEstimate[numEffectEstimate].Id)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimateDescription(numEffectEstimate int) templ.Component {

	if resource == nil || len(resource.EffectEstimate) >= numEffectEstimate {
		return StringInput("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].Description", nil)
	}
	return StringInput("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].Description", resource.EffectEstimate[numEffectEstimate].Description)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimateType(numEffectEstimate int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.EffectEstimate) >= numEffectEstimate {
		return CodeableConceptSelect("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].Type", resource.EffectEstimate[numEffectEstimate].Type, optionsValueSet)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimateVariantState(numEffectEstimate int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.EffectEstimate) >= numEffectEstimate {
		return CodeableConceptSelect("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].VariantState", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].VariantState", resource.EffectEstimate[numEffectEstimate].VariantState, optionsValueSet)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimateValue(numEffectEstimate int) templ.Component {

	if resource == nil || len(resource.EffectEstimate) >= numEffectEstimate {
		return Float64Input("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].Value", nil)
	}
	return Float64Input("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].Value", resource.EffectEstimate[numEffectEstimate].Value)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimateUnitOfMeasure(numEffectEstimate int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.EffectEstimate) >= numEffectEstimate {
		return CodeableConceptSelect("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].UnitOfMeasure", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].UnitOfMeasure", resource.EffectEstimate[numEffectEstimate].UnitOfMeasure, optionsValueSet)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimatePrecisionEstimateId(numEffectEstimate int, numPrecisionEstimate int) templ.Component {

	if resource == nil || len(resource.EffectEstimate) >= numEffectEstimate || len(resource.EffectEstimate[numEffectEstimate].PrecisionEstimate) >= numPrecisionEstimate {
		return StringInput("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Id", nil)
	}
	return StringInput("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Id", resource.EffectEstimate[numEffectEstimate].PrecisionEstimate[numPrecisionEstimate].Id)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimatePrecisionEstimateType(numEffectEstimate int, numPrecisionEstimate int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.EffectEstimate) >= numEffectEstimate || len(resource.EffectEstimate[numEffectEstimate].PrecisionEstimate) >= numPrecisionEstimate {
		return CodeableConceptSelect("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Type", resource.EffectEstimate[numEffectEstimate].PrecisionEstimate[numPrecisionEstimate].Type, optionsValueSet)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimatePrecisionEstimateLevel(numEffectEstimate int, numPrecisionEstimate int) templ.Component {

	if resource == nil || len(resource.EffectEstimate) >= numEffectEstimate || len(resource.EffectEstimate[numEffectEstimate].PrecisionEstimate) >= numPrecisionEstimate {
		return Float64Input("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Level", nil)
	}
	return Float64Input("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Level", resource.EffectEstimate[numEffectEstimate].PrecisionEstimate[numPrecisionEstimate].Level)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimatePrecisionEstimateFrom(numEffectEstimate int, numPrecisionEstimate int) templ.Component {

	if resource == nil || len(resource.EffectEstimate) >= numEffectEstimate || len(resource.EffectEstimate[numEffectEstimate].PrecisionEstimate) >= numPrecisionEstimate {
		return Float64Input("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].From", nil)
	}
	return Float64Input("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].From", resource.EffectEstimate[numEffectEstimate].PrecisionEstimate[numPrecisionEstimate].From)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimatePrecisionEstimateTo(numEffectEstimate int, numPrecisionEstimate int) templ.Component {

	if resource == nil || len(resource.EffectEstimate) >= numEffectEstimate || len(resource.EffectEstimate[numEffectEstimate].PrecisionEstimate) >= numPrecisionEstimate {
		return Float64Input("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].To", nil)
	}
	return Float64Input("EffectEvidenceSynthesis.EffectEstimate["+strconv.Itoa(numEffectEstimate)+"].PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].To", resource.EffectEstimate[numEffectEstimate].PrecisionEstimate[numPrecisionEstimate].To)
}
func (resource *EffectEvidenceSynthesis) T_CertaintyId(numCertainty int) templ.Component {

	if resource == nil || len(resource.Certainty) >= numCertainty {
		return StringInput("EffectEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].Id", nil)
	}
	return StringInput("EffectEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].Id", resource.Certainty[numCertainty].Id)
}
func (resource *EffectEvidenceSynthesis) T_CertaintyRating(numCertainty int, numRating int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Certainty) >= numCertainty || len(resource.Certainty[numCertainty].Rating) >= numRating {
		return CodeableConceptSelect("EffectEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].Rating["+strconv.Itoa(numRating)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EffectEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].Rating["+strconv.Itoa(numRating)+"]", &resource.Certainty[numCertainty].Rating[numRating], optionsValueSet)
}
func (resource *EffectEvidenceSynthesis) T_CertaintyCertaintySubcomponentId(numCertainty int, numCertaintySubcomponent int) templ.Component {

	if resource == nil || len(resource.Certainty) >= numCertainty || len(resource.Certainty[numCertainty].CertaintySubcomponent) >= numCertaintySubcomponent {
		return StringInput("EffectEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Id", nil)
	}
	return StringInput("EffectEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Id", resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Id)
}
func (resource *EffectEvidenceSynthesis) T_CertaintyCertaintySubcomponentType(numCertainty int, numCertaintySubcomponent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Certainty) >= numCertainty || len(resource.Certainty[numCertainty].CertaintySubcomponent) >= numCertaintySubcomponent {
		return CodeableConceptSelect("EffectEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EffectEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Type", resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Type, optionsValueSet)
}
func (resource *EffectEvidenceSynthesis) T_CertaintyCertaintySubcomponentRating(numCertainty int, numCertaintySubcomponent int, numRating int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Certainty) >= numCertainty || len(resource.Certainty[numCertainty].CertaintySubcomponent) >= numCertaintySubcomponent || len(resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Rating) >= numRating {
		return CodeableConceptSelect("EffectEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Rating["+strconv.Itoa(numRating)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EffectEvidenceSynthesis.Certainty["+strconv.Itoa(numCertainty)+"].CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Rating["+strconv.Itoa(numRating)+"]", &resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Rating[numRating], optionsValueSet)
}
