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
	Date                *time.Time                                 `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher           *string                                    `json:"publisher,omitempty"`
	Contact             []ContactDetail                            `json:"contact,omitempty"`
	Description         *string                                    `json:"description,omitempty"`
	Note                []Annotation                               `json:"note,omitempty"`
	UseContext          []UsageContext                             `json:"useContext,omitempty"`
	Jurisdiction        []CodeableConcept                          `json:"jurisdiction,omitempty"`
	Copyright           *string                                    `json:"copyright,omitempty"`
	ApprovalDate        *time.Time                                 `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate      *time.Time                                 `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
func (resource *EffectEvidenceSynthesis) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Copyright", nil, htmlAttrs)
	}
	return StringInput("Copyright", resource.Copyright, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_SynthesisType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SynthesisType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SynthesisType", resource.SynthesisType, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_StudyType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("StudyType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("StudyType", resource.StudyType, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_SampleSizeDescription(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SampleSizeDescription", nil, htmlAttrs)
	}
	return StringInput("SampleSizeDescription", resource.SampleSize.Description, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_SampleSizeNumberOfStudies(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("SampleSizeNumberOfStudies", nil, htmlAttrs)
	}
	return IntInput("SampleSizeNumberOfStudies", resource.SampleSize.NumberOfStudies, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_SampleSizeNumberOfParticipants(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("SampleSizeNumberOfParticipants", nil, htmlAttrs)
	}
	return IntInput("SampleSizeNumberOfParticipants", resource.SampleSize.NumberOfParticipants, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_ResultsByExposureDescription(numResultsByExposure int, htmlAttrs string) templ.Component {
	if resource == nil || numResultsByExposure >= len(resource.ResultsByExposure) {
		return StringInput("ResultsByExposure["+strconv.Itoa(numResultsByExposure)+"]Description", nil, htmlAttrs)
	}
	return StringInput("ResultsByExposure["+strconv.Itoa(numResultsByExposure)+"]Description", resource.ResultsByExposure[numResultsByExposure].Description, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_ResultsByExposureExposureState(numResultsByExposure int, htmlAttrs string) templ.Component {
	optionsValueSet := VSExposure_state

	if resource == nil || numResultsByExposure >= len(resource.ResultsByExposure) {
		return CodeSelect("ResultsByExposure["+strconv.Itoa(numResultsByExposure)+"]ExposureState", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ResultsByExposure["+strconv.Itoa(numResultsByExposure)+"]ExposureState", resource.ResultsByExposure[numResultsByExposure].ExposureState, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_ResultsByExposureVariantState(numResultsByExposure int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numResultsByExposure >= len(resource.ResultsByExposure) {
		return CodeableConceptSelect("ResultsByExposure["+strconv.Itoa(numResultsByExposure)+"]VariantState", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ResultsByExposure["+strconv.Itoa(numResultsByExposure)+"]VariantState", resource.ResultsByExposure[numResultsByExposure].VariantState, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimateDescription(numEffectEstimate int, htmlAttrs string) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) {
		return StringInput("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]Description", nil, htmlAttrs)
	}
	return StringInput("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]Description", resource.EffectEstimate[numEffectEstimate].Description, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimateType(numEffectEstimate int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) {
		return CodeableConceptSelect("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]Type", resource.EffectEstimate[numEffectEstimate].Type, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimateVariantState(numEffectEstimate int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) {
		return CodeableConceptSelect("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]VariantState", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]VariantState", resource.EffectEstimate[numEffectEstimate].VariantState, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimateValue(numEffectEstimate int, htmlAttrs string) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) {
		return Float64Input("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]Value", nil, htmlAttrs)
	}
	return Float64Input("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]Value", resource.EffectEstimate[numEffectEstimate].Value, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimateUnitOfMeasure(numEffectEstimate int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) {
		return CodeableConceptSelect("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]UnitOfMeasure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]UnitOfMeasure", resource.EffectEstimate[numEffectEstimate].UnitOfMeasure, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimatePrecisionEstimateType(numEffectEstimate int, numPrecisionEstimate int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) || numPrecisionEstimate >= len(resource.EffectEstimate[numEffectEstimate].PrecisionEstimate) {
		return CodeableConceptSelect("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Type", resource.EffectEstimate[numEffectEstimate].PrecisionEstimate[numPrecisionEstimate].Type, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimatePrecisionEstimateLevel(numEffectEstimate int, numPrecisionEstimate int, htmlAttrs string) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) || numPrecisionEstimate >= len(resource.EffectEstimate[numEffectEstimate].PrecisionEstimate) {
		return Float64Input("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Level", nil, htmlAttrs)
	}
	return Float64Input("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].Level", resource.EffectEstimate[numEffectEstimate].PrecisionEstimate[numPrecisionEstimate].Level, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimatePrecisionEstimateFrom(numEffectEstimate int, numPrecisionEstimate int, htmlAttrs string) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) || numPrecisionEstimate >= len(resource.EffectEstimate[numEffectEstimate].PrecisionEstimate) {
		return Float64Input("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].From", nil, htmlAttrs)
	}
	return Float64Input("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].From", resource.EffectEstimate[numEffectEstimate].PrecisionEstimate[numPrecisionEstimate].From, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_EffectEstimatePrecisionEstimateTo(numEffectEstimate int, numPrecisionEstimate int, htmlAttrs string) templ.Component {
	if resource == nil || numEffectEstimate >= len(resource.EffectEstimate) || numPrecisionEstimate >= len(resource.EffectEstimate[numEffectEstimate].PrecisionEstimate) {
		return Float64Input("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].To", nil, htmlAttrs)
	}
	return Float64Input("EffectEstimate["+strconv.Itoa(numEffectEstimate)+"]PrecisionEstimate["+strconv.Itoa(numPrecisionEstimate)+"].To", resource.EffectEstimate[numEffectEstimate].PrecisionEstimate[numPrecisionEstimate].To, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_CertaintyRating(numCertainty int, numRating int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numRating >= len(resource.Certainty[numCertainty].Rating) {
		return CodeableConceptSelect("Certainty["+strconv.Itoa(numCertainty)+"]Rating["+strconv.Itoa(numRating)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Certainty["+strconv.Itoa(numCertainty)+"]Rating["+strconv.Itoa(numRating)+"]", &resource.Certainty[numCertainty].Rating[numRating], optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_CertaintyNote(numCertainty int, numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numNote >= len(resource.Certainty[numCertainty].Note) {
		return AnnotationTextArea("Certainty["+strconv.Itoa(numCertainty)+"]Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Certainty["+strconv.Itoa(numCertainty)+"]Note["+strconv.Itoa(numNote)+"]", &resource.Certainty[numCertainty].Note[numNote], htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_CertaintyCertaintySubcomponentType(numCertainty int, numCertaintySubcomponent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numCertaintySubcomponent >= len(resource.Certainty[numCertainty].CertaintySubcomponent) {
		return CodeableConceptSelect("Certainty["+strconv.Itoa(numCertainty)+"]CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Certainty["+strconv.Itoa(numCertainty)+"]CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Type", resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Type, optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_CertaintyCertaintySubcomponentRating(numCertainty int, numCertaintySubcomponent int, numRating int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numCertaintySubcomponent >= len(resource.Certainty[numCertainty].CertaintySubcomponent) || numRating >= len(resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Rating) {
		return CodeableConceptSelect("Certainty["+strconv.Itoa(numCertainty)+"]CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Rating["+strconv.Itoa(numRating)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Certainty["+strconv.Itoa(numCertainty)+"]CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Rating["+strconv.Itoa(numRating)+"]", &resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Rating[numRating], optionsValueSet, htmlAttrs)
}
func (resource *EffectEvidenceSynthesis) T_CertaintyCertaintySubcomponentNote(numCertainty int, numCertaintySubcomponent int, numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numCertainty >= len(resource.Certainty) || numCertaintySubcomponent >= len(resource.Certainty[numCertainty].CertaintySubcomponent) || numNote >= len(resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Note) {
		return AnnotationTextArea("Certainty["+strconv.Itoa(numCertainty)+"]CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Certainty["+strconv.Itoa(numCertainty)+"]CertaintySubcomponent["+strconv.Itoa(numCertaintySubcomponent)+"].Note["+strconv.Itoa(numNote)+"]", &resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Note[numNote], htmlAttrs)
}
