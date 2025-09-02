package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	CiteAsReference    *Reference                   `json:"citeAsReference"`
	CiteAsMarkdown     *string                      `json:"citeAsMarkdown"`
	Status             string                       `json:"status"`
	Date               *string                      `json:"date,omitempty"`
	UseContext         []UsageContext               `json:"useContext,omitempty"`
	ApprovalDate       *string                      `json:"approvalDate,omitempty"`
	LastReviewDate     *string                      `json:"lastReviewDate,omitempty"`
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

func (resource *Evidence) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Evidence) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Evidence) T_SynthesisType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("synthesisType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("synthesisType", resource.SynthesisType, optionsValueSet)
}
func (resource *Evidence) T_StudyType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("studyType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("studyType", resource.StudyType, optionsValueSet)
}
func (resource *Evidence) T_VariableDefinitionVariableRole(numVariableDefinition int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.VariableDefinition) >= numVariableDefinition {
		return CodeableConceptSelect("variableRole", nil, optionsValueSet)
	}
	return CodeableConceptSelect("variableRole", &resource.VariableDefinition[numVariableDefinition].VariableRole, optionsValueSet)
}
func (resource *Evidence) T_VariableDefinitionDirectnessMatch(numVariableDefinition int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.VariableDefinition) >= numVariableDefinition {
		return CodeableConceptSelect("directnessMatch", nil, optionsValueSet)
	}
	return CodeableConceptSelect("directnessMatch", resource.VariableDefinition[numVariableDefinition].DirectnessMatch, optionsValueSet)
}
func (resource *Evidence) T_StatisticStatisticType(numStatistic int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Statistic) >= numStatistic {
		return CodeableConceptSelect("statisticType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("statisticType", resource.Statistic[numStatistic].StatisticType, optionsValueSet)
}
func (resource *Evidence) T_StatisticCategory(numStatistic int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Statistic) >= numStatistic {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.Statistic[numStatistic].Category, optionsValueSet)
}
func (resource *Evidence) T_StatisticAttributeEstimateType(numStatistic int, numAttributeEstimate int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Statistic[numStatistic].AttributeEstimate) >= numAttributeEstimate {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Statistic[numStatistic].AttributeEstimate[numAttributeEstimate].Type, optionsValueSet)
}
func (resource *Evidence) T_StatisticModelCharacteristicCode(numStatistic int, numModelCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Statistic[numStatistic].ModelCharacteristic) >= numModelCharacteristic {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Code, optionsValueSet)
}
func (resource *Evidence) T_StatisticModelCharacteristicVariableHandling(numStatistic int, numModelCharacteristic int, numVariable int) templ.Component {
	optionsValueSet := VSVariable_handling

	if resource == nil && len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable) >= numVariable {
		return CodeSelect("handling", nil, optionsValueSet)
	}
	return CodeSelect("handling", resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].Handling, optionsValueSet)
}
func (resource *Evidence) T_StatisticModelCharacteristicVariableValueCategory(numStatistic int, numModelCharacteristic int, numVariable int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable) >= numVariable {
		return CodeableConceptSelect("valueCategory", nil, optionsValueSet)
	}
	return CodeableConceptSelect("valueCategory", &resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].ValueCategory[0], optionsValueSet)
}
func (resource *Evidence) T_CertaintyType(numCertainty int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Certainty) >= numCertainty {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Certainty[numCertainty].Type, optionsValueSet)
}
func (resource *Evidence) T_CertaintyRating(numCertainty int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Certainty) >= numCertainty {
		return CodeableConceptSelect("rating", nil, optionsValueSet)
	}
	return CodeableConceptSelect("rating", resource.Certainty[numCertainty].Rating, optionsValueSet)
}
