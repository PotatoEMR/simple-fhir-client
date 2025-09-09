package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/ResearchStudy
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

// http://hl7.org/fhir/r4b/StructureDefinition/ResearchStudy
type ResearchStudyArm struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Name              string           `json:"name"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Description       *string          `json:"description,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ResearchStudy
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
func (r ResearchStudy) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ResearchStudy/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ResearchStudy"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ResearchStudy) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ResearchStudy.Title", nil, htmlAttrs)
	}
	return StringInput("ResearchStudy.Title", resource.Title, htmlAttrs)
}
func (resource *ResearchStudy) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSResearch_study_status

	if resource == nil {
		return CodeSelect("ResearchStudy.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ResearchStudy.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_PrimaryPurposeType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ResearchStudy.PrimaryPurposeType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ResearchStudy.PrimaryPurposeType", resource.PrimaryPurposeType, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Phase(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ResearchStudy.Phase", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ResearchStudy.Phase", resource.Phase, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("ResearchStudy.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ResearchStudy.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Focus(numFocus int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numFocus >= len(resource.Focus) {
		return CodeableConceptSelect("ResearchStudy.Focus["+strconv.Itoa(numFocus)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ResearchStudy.Focus["+strconv.Itoa(numFocus)+"]", &resource.Focus[numFocus], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Condition(numCondition int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return CodeableConceptSelect("ResearchStudy.Condition["+strconv.Itoa(numCondition)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ResearchStudy.Condition["+strconv.Itoa(numCondition)+"]", &resource.Condition[numCondition], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Keyword(numKeyword int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numKeyword >= len(resource.Keyword) {
		return CodeableConceptSelect("ResearchStudy.Keyword["+strconv.Itoa(numKeyword)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ResearchStudy.Keyword["+strconv.Itoa(numKeyword)+"]", &resource.Keyword[numKeyword], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Location(numLocation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLocation >= len(resource.Location) {
		return CodeableConceptSelect("ResearchStudy.Location["+strconv.Itoa(numLocation)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ResearchStudy.Location["+strconv.Itoa(numLocation)+"]", &resource.Location[numLocation], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ResearchStudy.Description", nil, htmlAttrs)
	}
	return StringInput("ResearchStudy.Description", resource.Description, htmlAttrs)
}
func (resource *ResearchStudy) T_ReasonStopped(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ResearchStudy.ReasonStopped", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ResearchStudy.ReasonStopped", resource.ReasonStopped, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("ResearchStudy.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("ResearchStudy.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *ResearchStudy) T_ArmName(numArm int, htmlAttrs string) templ.Component {
	if resource == nil || numArm >= len(resource.Arm) {
		return StringInput("ResearchStudy.Arm["+strconv.Itoa(numArm)+"].Name", nil, htmlAttrs)
	}
	return StringInput("ResearchStudy.Arm["+strconv.Itoa(numArm)+"].Name", &resource.Arm[numArm].Name, htmlAttrs)
}
func (resource *ResearchStudy) T_ArmType(numArm int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numArm >= len(resource.Arm) {
		return CodeableConceptSelect("ResearchStudy.Arm["+strconv.Itoa(numArm)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ResearchStudy.Arm["+strconv.Itoa(numArm)+"].Type", resource.Arm[numArm].Type, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_ArmDescription(numArm int, htmlAttrs string) templ.Component {
	if resource == nil || numArm >= len(resource.Arm) {
		return StringInput("ResearchStudy.Arm["+strconv.Itoa(numArm)+"].Description", nil, htmlAttrs)
	}
	return StringInput("ResearchStudy.Arm["+strconv.Itoa(numArm)+"].Description", resource.Arm[numArm].Description, htmlAttrs)
}
func (resource *ResearchStudy) T_ObjectiveName(numObjective int, htmlAttrs string) templ.Component {
	if resource == nil || numObjective >= len(resource.Objective) {
		return StringInput("ResearchStudy.Objective["+strconv.Itoa(numObjective)+"].Name", nil, htmlAttrs)
	}
	return StringInput("ResearchStudy.Objective["+strconv.Itoa(numObjective)+"].Name", resource.Objective[numObjective].Name, htmlAttrs)
}
func (resource *ResearchStudy) T_ObjectiveType(numObjective int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numObjective >= len(resource.Objective) {
		return CodeableConceptSelect("ResearchStudy.Objective["+strconv.Itoa(numObjective)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ResearchStudy.Objective["+strconv.Itoa(numObjective)+"].Type", resource.Objective[numObjective].Type, optionsValueSet, htmlAttrs)
}
