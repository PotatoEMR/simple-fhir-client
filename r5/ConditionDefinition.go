package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	VersionAlgorithmString *string                            `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                            `json:"versionAlgorithmCoding,omitempty"`
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
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
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
func (r ConditionDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ConditionDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ConditionDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ConditionDefinition) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *ConditionDefinition) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *ConditionDefinition) T_VersionAlgorithmString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *ConditionDefinition) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *ConditionDefinition) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *ConditionDefinition) T_Subtitle(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("subtitle", nil, htmlAttrs)
	}
	return StringInput("subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *ConditionDefinition) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *ConditionDefinition) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *ConditionDefinition) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *ConditionDefinition) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *ConditionDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_Severity(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("severity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("severity", resource.Severity, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_BodySite(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("bodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("bodySite", resource.BodySite, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_Stage(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("stage", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("stage", resource.Stage, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_HasSeverity(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("hasSeverity", nil, htmlAttrs)
	}
	return BoolInput("hasSeverity", resource.HasSeverity, htmlAttrs)
}
func (resource *ConditionDefinition) T_HasBodySite(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("hasBodySite", nil, htmlAttrs)
	}
	return BoolInput("hasBodySite", resource.HasBodySite, htmlAttrs)
}
func (resource *ConditionDefinition) T_HasStage(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("hasStage", nil, htmlAttrs)
	}
	return BoolInput("hasStage", resource.HasStage, htmlAttrs)
}
func (resource *ConditionDefinition) T_Definition(numDefinition int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDefinition >= len(resource.Definition) {
		return StringInput("definition["+strconv.Itoa(numDefinition)+"]", nil, htmlAttrs)
	}
	return StringInput("definition["+strconv.Itoa(numDefinition)+"]", &resource.Definition[numDefinition], htmlAttrs)
}
func (resource *ConditionDefinition) T_ObservationCategory(numObservation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numObservation >= len(resource.Observation) {
		return CodeableConceptSelect("observation["+strconv.Itoa(numObservation)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("observation["+strconv.Itoa(numObservation)+"].category", resource.Observation[numObservation].Category, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_ObservationCode(numObservation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numObservation >= len(resource.Observation) {
		return CodeableConceptSelect("observation["+strconv.Itoa(numObservation)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("observation["+strconv.Itoa(numObservation)+"].code", resource.Observation[numObservation].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_MedicationCategory(numMedication int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMedication >= len(resource.Medication) {
		return CodeableConceptSelect("medication["+strconv.Itoa(numMedication)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("medication["+strconv.Itoa(numMedication)+"].category", resource.Medication[numMedication].Category, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_MedicationCode(numMedication int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMedication >= len(resource.Medication) {
		return CodeableConceptSelect("medication["+strconv.Itoa(numMedication)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("medication["+strconv.Itoa(numMedication)+"].code", resource.Medication[numMedication].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_PreconditionType(numPrecondition int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCondition_precondition_type

	if resource == nil || numPrecondition >= len(resource.Precondition) {
		return CodeSelect("precondition["+strconv.Itoa(numPrecondition)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("precondition["+strconv.Itoa(numPrecondition)+"].type", &resource.Precondition[numPrecondition].Type, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_PreconditionCode(numPrecondition int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrecondition >= len(resource.Precondition) {
		return CodeableConceptSelect("precondition["+strconv.Itoa(numPrecondition)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("precondition["+strconv.Itoa(numPrecondition)+"].code", &resource.Precondition[numPrecondition].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_PreconditionValueCodeableConcept(numPrecondition int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrecondition >= len(resource.Precondition) {
		return CodeableConceptSelect("precondition["+strconv.Itoa(numPrecondition)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("precondition["+strconv.Itoa(numPrecondition)+"].valueCodeableConcept", resource.Precondition[numPrecondition].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_QuestionnairePurpose(numQuestionnaire int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCondition_questionnaire_purpose

	if resource == nil || numQuestionnaire >= len(resource.Questionnaire) {
		return CodeSelect("questionnaire["+strconv.Itoa(numQuestionnaire)+"].purpose", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("questionnaire["+strconv.Itoa(numQuestionnaire)+"].purpose", &resource.Questionnaire[numQuestionnaire].Purpose, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_PlanRole(numPlan int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) {
		return CodeableConceptSelect("plan["+strconv.Itoa(numPlan)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("plan["+strconv.Itoa(numPlan)+"].role", resource.Plan[numPlan].Role, optionsValueSet, htmlAttrs)
}
