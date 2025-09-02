package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *RiskEvidenceSynthesis) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_Jurisdiction(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_Topic(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("topic", nil, optionsValueSet)
	}
	return CodeableConceptSelect("topic", &resource.Topic[0], optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_SynthesisType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("synthesisType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("synthesisType", resource.SynthesisType, optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_StudyType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("studyType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("studyType", resource.StudyType, optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.RiskEstimate.Type, optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimateUnitOfMeasure(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("unitOfMeasure", nil, optionsValueSet)
	}
	return CodeableConceptSelect("unitOfMeasure", resource.RiskEstimate.UnitOfMeasure, optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_RiskEstimatePrecisionEstimateType(numPrecisionEstimate int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.RiskEstimate.PrecisionEstimate) >= numPrecisionEstimate {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.RiskEstimate.PrecisionEstimate[numPrecisionEstimate].Type, optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyRating(numCertainty int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Certainty) >= numCertainty {
		return CodeableConceptSelect("rating", nil, optionsValueSet)
	}
	return CodeableConceptSelect("rating", &resource.Certainty[numCertainty].Rating[0], optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyCertaintySubcomponentType(numCertainty int, numCertaintySubcomponent int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Certainty[numCertainty].CertaintySubcomponent) >= numCertaintySubcomponent {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Type, optionsValueSet)
}
func (resource *RiskEvidenceSynthesis) T_CertaintyCertaintySubcomponentRating(numCertainty int, numCertaintySubcomponent int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Certainty[numCertainty].CertaintySubcomponent) >= numCertaintySubcomponent {
		return CodeableConceptSelect("rating", nil, optionsValueSet)
	}
	return CodeableConceptSelect("rating", &resource.Certainty[numCertainty].CertaintySubcomponent[numCertaintySubcomponent].Rating[0], optionsValueSet)
}
