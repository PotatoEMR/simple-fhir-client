package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/Evidence
type Evidence struct {
	Id                     *string                      `json:"id,omitempty"`
	Meta                   *Meta                        `json:"meta,omitempty"`
	ImplicitRules          *string                      `json:"implicitRules,omitempty"`
	Language               *string                      `json:"language,omitempty"`
	Text                   *Narrative                   `json:"text,omitempty"`
	Contained              []Resource                   `json:"contained,omitempty"`
	Extension              []Extension                  `json:"extension,omitempty"`
	ModifierExtension      []Extension                  `json:"modifierExtension,omitempty"`
	Url                    *string                      `json:"url,omitempty"`
	Identifier             []Identifier                 `json:"identifier,omitempty"`
	Version                *string                      `json:"version,omitempty"`
	VersionAlgorithmString *string                      `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                      `json:"versionAlgorithmCoding"`
	Name                   *string                      `json:"name,omitempty"`
	Title                  *string                      `json:"title,omitempty"`
	CiteAsReference        *Reference                   `json:"citeAsReference"`
	CiteAsMarkdown         *string                      `json:"citeAsMarkdown"`
	Status                 string                       `json:"status"`
	Experimental           *bool                        `json:"experimental,omitempty"`
	Date                   *string                      `json:"date,omitempty"`
	ApprovalDate           *string                      `json:"approvalDate,omitempty"`
	LastReviewDate         *string                      `json:"lastReviewDate,omitempty"`
	Publisher              *string                      `json:"publisher,omitempty"`
	Contact                []ContactDetail              `json:"contact,omitempty"`
	Author                 []ContactDetail              `json:"author,omitempty"`
	Editor                 []ContactDetail              `json:"editor,omitempty"`
	Reviewer               []ContactDetail              `json:"reviewer,omitempty"`
	Endorser               []ContactDetail              `json:"endorser,omitempty"`
	UseContext             []UsageContext               `json:"useContext,omitempty"`
	Purpose                *string                      `json:"purpose,omitempty"`
	Copyright              *string                      `json:"copyright,omitempty"`
	CopyrightLabel         *string                      `json:"copyrightLabel,omitempty"`
	RelatedArtifact        []RelatedArtifact            `json:"relatedArtifact,omitempty"`
	Description            *string                      `json:"description,omitempty"`
	Assertion              *string                      `json:"assertion,omitempty"`
	Note                   []Annotation                 `json:"note,omitempty"`
	VariableDefinition     []EvidenceVariableDefinition `json:"variableDefinition"`
	SynthesisType          *CodeableConcept             `json:"synthesisType,omitempty"`
	StudyDesign            []CodeableConcept            `json:"studyDesign,omitempty"`
	Statistic              []EvidenceStatistic          `json:"statistic,omitempty"`
	Certainty              []EvidenceCertainty          `json:"certainty,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Evidence
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

// http://hl7.org/fhir/r5/StructureDefinition/Evidence
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

// http://hl7.org/fhir/r5/StructureDefinition/Evidence
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

// http://hl7.org/fhir/r5/StructureDefinition/Evidence
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

// http://hl7.org/fhir/r5/StructureDefinition/Evidence
type EvidenceStatisticModelCharacteristic struct {
	Id                *string                                        `json:"id,omitempty"`
	Extension         []Extension                                    `json:"extension,omitempty"`
	ModifierExtension []Extension                                    `json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                `json:"code"`
	Value             *Quantity                                      `json:"value,omitempty"`
	Variable          []EvidenceStatisticModelCharacteristicVariable `json:"variable,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Evidence
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

// http://hl7.org/fhir/r5/StructureDefinition/Evidence
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

func (resource *Evidence) EvidenceLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Evidence) EvidenceStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *Evidence) EvidenceStatisticModelCharacteristicVariableHandling(numStatistic int, numModelCharacteristic int, numVariable int) templ.Component {
	optionsValueSet := VSVariable_handling
	currentVal := ""
	if resource != nil && len(resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable) >= numVariable {
		currentVal = *resource.Statistic[numStatistic].ModelCharacteristic[numModelCharacteristic].Variable[numVariable].Handling
	}
	return CodeSelect("handling", currentVal, optionsValueSet)
}
