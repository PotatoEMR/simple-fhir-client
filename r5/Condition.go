package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/Condition
type Condition struct {
	Id                 *string                `json:"id,omitempty"`
	Meta               *Meta                  `json:"meta,omitempty"`
	ImplicitRules      *string                `json:"implicitRules,omitempty"`
	Language           *string                `json:"language,omitempty"`
	Text               *Narrative             `json:"text,omitempty"`
	Contained          []Resource             `json:"contained,omitempty"`
	Extension          []Extension            `json:"extension,omitempty"`
	ModifierExtension  []Extension            `json:"modifierExtension,omitempty"`
	Identifier         []Identifier           `json:"identifier,omitempty"`
	ClinicalStatus     CodeableConcept        `json:"clinicalStatus"`
	VerificationStatus *CodeableConcept       `json:"verificationStatus,omitempty"`
	Category           []CodeableConcept      `json:"category,omitempty"`
	Severity           *CodeableConcept       `json:"severity,omitempty"`
	Code               *CodeableConcept       `json:"code,omitempty"`
	BodySite           []CodeableConcept      `json:"bodySite,omitempty"`
	Subject            Reference              `json:"subject"`
	Encounter          *Reference             `json:"encounter,omitempty"`
	OnsetDateTime      *string                `json:"onsetDateTime"`
	OnsetAge           *Age                   `json:"onsetAge"`
	OnsetPeriod        *Period                `json:"onsetPeriod"`
	OnsetRange         *Range                 `json:"onsetRange"`
	OnsetString        *string                `json:"onsetString"`
	AbatementDateTime  *string                `json:"abatementDateTime"`
	AbatementAge       *Age                   `json:"abatementAge"`
	AbatementPeriod    *Period                `json:"abatementPeriod"`
	AbatementRange     *Range                 `json:"abatementRange"`
	AbatementString    *string                `json:"abatementString"`
	RecordedDate       *string                `json:"recordedDate,omitempty"`
	Participant        []ConditionParticipant `json:"participant,omitempty"`
	Stage              []ConditionStage       `json:"stage,omitempty"`
	Evidence           []CodeableReference    `json:"evidence,omitempty"`
	Note               []Annotation           `json:"note,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Condition
type ConditionParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Condition
type ConditionStage struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Summary           *CodeableConcept `json:"summary,omitempty"`
	Assessment        []Reference      `json:"assessment,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
}

type OtherCondition Condition

// on convert struct to json, automatically add resourceType=Condition
func (r Condition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCondition
		ResourceType string `json:"resourceType"`
	}{
		OtherCondition: OtherCondition(r),
		ResourceType:   "Condition",
	})
}

func (resource *Condition) ConditionLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Condition) ConditionClinicalStatus(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("clinicalStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("clinicalStatus", &resource.ClinicalStatus, optionsValueSet)
}
func (resource *Condition) ConditionVerificationStatus(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("verificationStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("verificationStatus", resource.VerificationStatus, optionsValueSet)
}
func (resource *Condition) ConditionCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *Condition) ConditionSeverity(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("severity", nil, optionsValueSet)
	}
	return CodeableConceptSelect("severity", resource.Severity, optionsValueSet)
}
func (resource *Condition) ConditionCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet)
}
func (resource *Condition) ConditionBodySite(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("bodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("bodySite", &resource.BodySite[0], optionsValueSet)
}
func (resource *Condition) ConditionParticipantFunction(numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Participant) >= numParticipant {
		return CodeableConceptSelect("function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("function", resource.Participant[numParticipant].Function, optionsValueSet)
}
func (resource *Condition) ConditionStageSummary(numStage int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Stage) >= numStage {
		return CodeableConceptSelect("summary", nil, optionsValueSet)
	}
	return CodeableConceptSelect("summary", resource.Stage[numStage].Summary, optionsValueSet)
}
func (resource *Condition) ConditionStageType(numStage int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Stage) >= numStage {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Stage[numStage].Type, optionsValueSet)
}
