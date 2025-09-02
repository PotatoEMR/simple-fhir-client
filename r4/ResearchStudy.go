package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/ResearchStudy
type ResearchStudy struct {
	Id                    *string                  `json:"id,omitempty"`
	Meta                  *Meta                    `json:"meta,omitempty"`
	ImplicitRules         *string                  `json:"implicitRules,omitempty"`
	Language              *string                  `json:"language,omitempty"`
	Text                  *Narrative               `json:"text,omitempty"`
	Contained             []Resource               `json:"contained,omitempty"`
	Extension             []Extension              `json:"extension,omitempty"`
	ModifierExtension     []Extension              `json:"modifierExtension,omitempty"`
	Identifier            []Identifier             `json:"identifier,omitempty"`
	Title                 *string                  `json:"title,omitempty"`
	Protocol              []Reference              `json:"protocol,omitempty"`
	PartOf                []Reference              `json:"partOf,omitempty"`
	Status                string                   `json:"status"`
	PrimaryPurposeType    *CodeableConcept         `json:"primaryPurposeType,omitempty"`
	Phase                 *CodeableConcept         `json:"phase,omitempty"`
	Category              []CodeableConcept        `json:"category,omitempty"`
	Focus                 []CodeableConcept        `json:"focus,omitempty"`
	Condition             []CodeableConcept        `json:"condition,omitempty"`
	Contact               []ContactDetail          `json:"contact,omitempty"`
	RelatedArtifact       []RelatedArtifact        `json:"relatedArtifact,omitempty"`
	Keyword               []CodeableConcept        `json:"keyword,omitempty"`
	Location              []CodeableConcept        `json:"location,omitempty"`
	Description           *string                  `json:"description,omitempty"`
	Enrollment            []Reference              `json:"enrollment,omitempty"`
	Period                *Period                  `json:"period,omitempty"`
	Sponsor               *Reference               `json:"sponsor,omitempty"`
	PrincipalInvestigator *Reference               `json:"principalInvestigator,omitempty"`
	Site                  []Reference              `json:"site,omitempty"`
	ReasonStopped         *CodeableConcept         `json:"reasonStopped,omitempty"`
	Note                  []Annotation             `json:"note,omitempty"`
	Arm                   []ResearchStudyArm       `json:"arm,omitempty"`
	Objective             []ResearchStudyObjective `json:"objective,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ResearchStudy
type ResearchStudyArm struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Name              string           `json:"name"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Description       *string          `json:"description,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ResearchStudy
type ResearchStudyObjective struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Name              *string          `json:"name,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
}

type OtherResearchStudy ResearchStudy

// on convert struct to json, automatically add resourceType=ResearchStudy
func (r ResearchStudy) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchStudy
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchStudy: OtherResearchStudy(r),
		ResourceType:       "ResearchStudy",
	})
}

func (resource *ResearchStudy) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ResearchStudy) T_Status() templ.Component {
	optionsValueSet := VSResearch_study_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ResearchStudy) T_PrimaryPurposeType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("primaryPurposeType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("primaryPurposeType", resource.PrimaryPurposeType, optionsValueSet)
}
func (resource *ResearchStudy) T_Phase(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("phase", nil, optionsValueSet)
	}
	return CodeableConceptSelect("phase", resource.Phase, optionsValueSet)
}
func (resource *ResearchStudy) T_Category(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *ResearchStudy) T_Focus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("focus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("focus", &resource.Focus[0], optionsValueSet)
}
func (resource *ResearchStudy) T_Condition(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("condition", nil, optionsValueSet)
	}
	return CodeableConceptSelect("condition", &resource.Condition[0], optionsValueSet)
}
func (resource *ResearchStudy) T_Keyword(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("keyword", nil, optionsValueSet)
	}
	return CodeableConceptSelect("keyword", &resource.Keyword[0], optionsValueSet)
}
func (resource *ResearchStudy) T_Location(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("location", nil, optionsValueSet)
	}
	return CodeableConceptSelect("location", &resource.Location[0], optionsValueSet)
}
func (resource *ResearchStudy) T_ReasonStopped(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("reasonStopped", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reasonStopped", resource.ReasonStopped, optionsValueSet)
}
func (resource *ResearchStudy) T_ArmType(numArm int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Arm) >= numArm {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Arm[numArm].Type, optionsValueSet)
}
func (resource *ResearchStudy) T_ObjectiveType(numObjective int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Objective) >= numObjective {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Objective[numObjective].Type, optionsValueSet)
}
