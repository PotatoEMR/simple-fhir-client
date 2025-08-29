package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/ConditionDefinition
type ConditionDefinition struct {
	Id                     *string                            `json:"id,omitempty"`
	Meta                   *Meta                              `json:"meta,omitempty"`
	ImplicitRules          *string                            `json:"implicitRules,omitempty"`
	Language               *string                            `json:"language,omitempty"`
	Text                   *Narrative                         `json:"text,omitempty"`
	Contained              []Resource                         `json:"contained,omitempty"`
	Extension              []Extension                        `json:"extension,omitempty"`
	ModifierExtension      []Extension                        `json:"modifierExtension,omitempty"`
	Url                    *string                            `json:"url,omitempty"`
	Identifier             []Identifier                       `json:"identifier,omitempty"`
	Version                *string                            `json:"version,omitempty"`
	VersionAlgorithmString *string                            `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                            `json:"versionAlgorithmCoding"`
	Name                   *string                            `json:"name,omitempty"`
	Title                  *string                            `json:"title,omitempty"`
	Subtitle               *string                            `json:"subtitle,omitempty"`
	Status                 string                             `json:"status"`
	Experimental           *bool                              `json:"experimental,omitempty"`
	Date                   *string                            `json:"date,omitempty"`
	Publisher              *string                            `json:"publisher,omitempty"`
	Contact                []ContactDetail                    `json:"contact,omitempty"`
	Description            *string                            `json:"description,omitempty"`
	UseContext             []UsageContext                     `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                  `json:"jurisdiction,omitempty"`
	Code                   CodeableConcept                    `json:"code"`
	Severity               *CodeableConcept                   `json:"severity,omitempty"`
	BodySite               *CodeableConcept                   `json:"bodySite,omitempty"`
	Stage                  *CodeableConcept                   `json:"stage,omitempty"`
	HasSeverity            *bool                              `json:"hasSeverity,omitempty"`
	HasBodySite            *bool                              `json:"hasBodySite,omitempty"`
	HasStage               *bool                              `json:"hasStage,omitempty"`
	Definition             []string                           `json:"definition,omitempty"`
	Observation            []ConditionDefinitionObservation   `json:"observation,omitempty"`
	Medication             []ConditionDefinitionMedication    `json:"medication,omitempty"`
	Precondition           []ConditionDefinitionPrecondition  `json:"precondition,omitempty"`
	Team                   []Reference                        `json:"team,omitempty"`
	Questionnaire          []ConditionDefinitionQuestionnaire `json:"questionnaire,omitempty"`
	Plan                   []ConditionDefinitionPlan          `json:"plan,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConditionDefinition
type ConditionDefinitionObservation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Category          *CodeableConcept `json:"category,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConditionDefinition
type ConditionDefinitionMedication struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Category          *CodeableConcept `json:"category,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConditionDefinition
type ConditionDefinitionPrecondition struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Type                 string           `json:"type"`
	Code                 CodeableConcept  `json:"code"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	ValueQuantity        *Quantity        `json:"valueQuantity"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConditionDefinition
type ConditionDefinitionQuestionnaire struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Purpose           string      `json:"purpose"`
	Reference         Reference   `json:"reference"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ConditionDefinition
type ConditionDefinitionPlan struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Reference         Reference        `json:"reference"`
}

type OtherConditionDefinition ConditionDefinition

// on convert struct to json, automatically add resourceType=ConditionDefinition
func (r ConditionDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherConditionDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherConditionDefinition: OtherConditionDefinition(r),
		ResourceType:             "ConditionDefinition",
	})
}

func (resource *ConditionDefinition) ConditionDefinitionLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionPreconditionType(numPrecondition int) templ.Component {
	optionsValueSet := VSCondition_precondition_type
	currentVal := ""
	if resource != nil && len(resource.Precondition) >= numPrecondition {
		currentVal = resource.Precondition[numPrecondition].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionQuestionnairePurpose(numQuestionnaire int) templ.Component {
	optionsValueSet := VSCondition_questionnaire_purpose
	currentVal := ""
	if resource != nil && len(resource.Questionnaire) >= numQuestionnaire {
		currentVal = resource.Questionnaire[numQuestionnaire].Purpose
	}
	return CodeSelect("purpose", currentVal, optionsValueSet)
}
