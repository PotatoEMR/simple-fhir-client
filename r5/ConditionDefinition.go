package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Date                   *time.Time                         `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *ConditionDefinition) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConditionDefinition.Url", nil, htmlAttrs)
	}
	return StringInput("ConditionDefinition.Url", resource.Url, htmlAttrs)
}
func (resource *ConditionDefinition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConditionDefinition.Version", nil, htmlAttrs)
	}
	return StringInput("ConditionDefinition.Version", resource.Version, htmlAttrs)
}
func (resource *ConditionDefinition) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConditionDefinition.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("ConditionDefinition.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *ConditionDefinition) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("ConditionDefinition.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("ConditionDefinition.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConditionDefinition.Name", nil, htmlAttrs)
	}
	return StringInput("ConditionDefinition.Name", resource.Name, htmlAttrs)
}
func (resource *ConditionDefinition) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConditionDefinition.Title", nil, htmlAttrs)
	}
	return StringInput("ConditionDefinition.Title", resource.Title, htmlAttrs)
}
func (resource *ConditionDefinition) T_Subtitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConditionDefinition.Subtitle", nil, htmlAttrs)
	}
	return StringInput("ConditionDefinition.Subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *ConditionDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ConditionDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConditionDefinition.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("ConditionDefinition.Experimental", nil, htmlAttrs)
	}
	return BoolInput("ConditionDefinition.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *ConditionDefinition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("ConditionDefinition.Date", nil, htmlAttrs)
	}
	return DateTimeInput("ConditionDefinition.Date", resource.Date, htmlAttrs)
}
func (resource *ConditionDefinition) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConditionDefinition.Publisher", nil, htmlAttrs)
	}
	return StringInput("ConditionDefinition.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *ConditionDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ConditionDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("ConditionDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *ConditionDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("ConditionDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ConditionDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ConditionDefinition.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ConditionDefinition.Code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_Severity(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ConditionDefinition.Severity", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ConditionDefinition.Severity", resource.Severity, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_BodySite(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ConditionDefinition.BodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ConditionDefinition.BodySite", resource.BodySite, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_Stage(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ConditionDefinition.Stage", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ConditionDefinition.Stage", resource.Stage, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_HasSeverity(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("ConditionDefinition.HasSeverity", nil, htmlAttrs)
	}
	return BoolInput("ConditionDefinition.HasSeverity", resource.HasSeverity, htmlAttrs)
}
func (resource *ConditionDefinition) T_HasBodySite(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("ConditionDefinition.HasBodySite", nil, htmlAttrs)
	}
	return BoolInput("ConditionDefinition.HasBodySite", resource.HasBodySite, htmlAttrs)
}
func (resource *ConditionDefinition) T_HasStage(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("ConditionDefinition.HasStage", nil, htmlAttrs)
	}
	return BoolInput("ConditionDefinition.HasStage", resource.HasStage, htmlAttrs)
}
func (resource *ConditionDefinition) T_Definition(numDefinition int, htmlAttrs string) templ.Component {
	if resource == nil || numDefinition >= len(resource.Definition) {
		return StringInput("ConditionDefinition.Definition["+strconv.Itoa(numDefinition)+"]", nil, htmlAttrs)
	}
	return StringInput("ConditionDefinition.Definition["+strconv.Itoa(numDefinition)+"]", &resource.Definition[numDefinition], htmlAttrs)
}
func (resource *ConditionDefinition) T_ObservationCategory(numObservation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numObservation >= len(resource.Observation) {
		return CodeableConceptSelect("ConditionDefinition.Observation["+strconv.Itoa(numObservation)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ConditionDefinition.Observation["+strconv.Itoa(numObservation)+"].Category", resource.Observation[numObservation].Category, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_ObservationCode(numObservation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numObservation >= len(resource.Observation) {
		return CodeableConceptSelect("ConditionDefinition.Observation["+strconv.Itoa(numObservation)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ConditionDefinition.Observation["+strconv.Itoa(numObservation)+"].Code", resource.Observation[numObservation].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_MedicationCategory(numMedication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMedication >= len(resource.Medication) {
		return CodeableConceptSelect("ConditionDefinition.Medication["+strconv.Itoa(numMedication)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ConditionDefinition.Medication["+strconv.Itoa(numMedication)+"].Category", resource.Medication[numMedication].Category, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_MedicationCode(numMedication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMedication >= len(resource.Medication) {
		return CodeableConceptSelect("ConditionDefinition.Medication["+strconv.Itoa(numMedication)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ConditionDefinition.Medication["+strconv.Itoa(numMedication)+"].Code", resource.Medication[numMedication].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_PreconditionType(numPrecondition int, htmlAttrs string) templ.Component {
	optionsValueSet := VSCondition_precondition_type

	if resource == nil || numPrecondition >= len(resource.Precondition) {
		return CodeSelect("ConditionDefinition.Precondition["+strconv.Itoa(numPrecondition)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConditionDefinition.Precondition["+strconv.Itoa(numPrecondition)+"].Type", &resource.Precondition[numPrecondition].Type, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_PreconditionCode(numPrecondition int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPrecondition >= len(resource.Precondition) {
		return CodeableConceptSelect("ConditionDefinition.Precondition["+strconv.Itoa(numPrecondition)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ConditionDefinition.Precondition["+strconv.Itoa(numPrecondition)+"].Code", &resource.Precondition[numPrecondition].Code, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_PreconditionValueCodeableConcept(numPrecondition int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPrecondition >= len(resource.Precondition) {
		return CodeableConceptSelect("ConditionDefinition.Precondition["+strconv.Itoa(numPrecondition)+"].ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ConditionDefinition.Precondition["+strconv.Itoa(numPrecondition)+"].ValueCodeableConcept", resource.Precondition[numPrecondition].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_QuestionnairePurpose(numQuestionnaire int, htmlAttrs string) templ.Component {
	optionsValueSet := VSCondition_questionnaire_purpose

	if resource == nil || numQuestionnaire >= len(resource.Questionnaire) {
		return CodeSelect("ConditionDefinition.Questionnaire["+strconv.Itoa(numQuestionnaire)+"].Purpose", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ConditionDefinition.Questionnaire["+strconv.Itoa(numQuestionnaire)+"].Purpose", &resource.Questionnaire[numQuestionnaire].Purpose, optionsValueSet, htmlAttrs)
}
func (resource *ConditionDefinition) T_PlanRole(numPlan int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPlan >= len(resource.Plan) {
		return CodeableConceptSelect("ConditionDefinition.Plan["+strconv.Itoa(numPlan)+"].Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ConditionDefinition.Plan["+strconv.Itoa(numPlan)+"].Role", resource.Plan[numPlan].Role, optionsValueSet, htmlAttrs)
}
