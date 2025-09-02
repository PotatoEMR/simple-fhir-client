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

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionJurisdiction(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Code, optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionSeverity(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("severity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("severity", resource.Severity, optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionBodySite(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("bodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("bodySite", resource.BodySite, optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionStage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("stage", nil, optionsValueSet)
	}
	return CodeableConceptSelect("stage", resource.Stage, optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionObservationCategory(numObservation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Observation) >= numObservation {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.Observation[numObservation].Category, optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionObservationCode(numObservation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Observation) >= numObservation {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Observation[numObservation].Code, optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionMedicationCategory(numMedication int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Medication) >= numMedication {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.Medication[numMedication].Category, optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionMedicationCode(numMedication int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Medication) >= numMedication {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Medication[numMedication].Code, optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionPreconditionType(numPrecondition int) templ.Component {
	optionsValueSet := VSCondition_precondition_type

	if resource == nil && len(resource.Precondition) >= numPrecondition {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", &resource.Precondition[numPrecondition].Type, optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionPreconditionCode(numPrecondition int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Precondition) >= numPrecondition {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Precondition[numPrecondition].Code, optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionQuestionnairePurpose(numQuestionnaire int) templ.Component {
	optionsValueSet := VSCondition_questionnaire_purpose

	if resource == nil && len(resource.Questionnaire) >= numQuestionnaire {
		return CodeSelect("purpose", nil, optionsValueSet)
	}
	return CodeSelect("purpose", &resource.Questionnaire[numQuestionnaire].Purpose, optionsValueSet)
}
func (resource *ConditionDefinition) ConditionDefinitionPlanRole(numPlan int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Plan) >= numPlan {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", resource.Plan[numPlan].Role, optionsValueSet)
}
