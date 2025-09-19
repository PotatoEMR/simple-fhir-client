package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
func (resource *ResearchStudy) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *ResearchStudy) T_Protocol(numProtocol int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProtocol >= len(resource.Protocol) {
		return ReferenceInput("protocol["+strconv.Itoa(numProtocol)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("protocol["+strconv.Itoa(numProtocol)+"]", &resource.Protocol[numProtocol], htmlAttrs)
}
func (resource *ResearchStudy) T_PartOf(numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput("partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *ResearchStudy) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSResearch_study_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_PrimaryPurposeType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("primaryPurposeType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("primaryPurposeType", resource.PrimaryPurposeType, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Phase(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("phase", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("phase", resource.Phase, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Focus(numFocus int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFocus >= len(resource.Focus) {
		return CodeableConceptSelect("focus["+strconv.Itoa(numFocus)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("focus["+strconv.Itoa(numFocus)+"]", &resource.Focus[numFocus], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Condition(numCondition int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCondition >= len(resource.Condition) {
		return CodeableConceptSelect("condition["+strconv.Itoa(numCondition)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("condition["+strconv.Itoa(numCondition)+"]", &resource.Condition[numCondition], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *ResearchStudy) T_RelatedArtifact(numRelatedArtifact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedArtifact >= len(resource.RelatedArtifact) {
		return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", &resource.RelatedArtifact[numRelatedArtifact], htmlAttrs)
}
func (resource *ResearchStudy) T_Keyword(numKeyword int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numKeyword >= len(resource.Keyword) {
		return CodeableConceptSelect("keyword["+strconv.Itoa(numKeyword)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("keyword["+strconv.Itoa(numKeyword)+"]", &resource.Keyword[numKeyword], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Location(numLocation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLocation >= len(resource.Location) {
		return CodeableConceptSelect("location["+strconv.Itoa(numLocation)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("location["+strconv.Itoa(numLocation)+"]", &resource.Location[numLocation], optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *ResearchStudy) T_Enrollment(numEnrollment int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEnrollment >= len(resource.Enrollment) {
		return ReferenceInput("enrollment["+strconv.Itoa(numEnrollment)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("enrollment["+strconv.Itoa(numEnrollment)+"]", &resource.Enrollment[numEnrollment], htmlAttrs)
}
func (resource *ResearchStudy) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *ResearchStudy) T_Sponsor(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("sponsor", nil, htmlAttrs)
	}
	return ReferenceInput("sponsor", resource.Sponsor, htmlAttrs)
}
func (resource *ResearchStudy) T_PrincipalInvestigator(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("principalInvestigator", nil, htmlAttrs)
	}
	return ReferenceInput("principalInvestigator", resource.PrincipalInvestigator, htmlAttrs)
}
func (resource *ResearchStudy) T_Site(numSite int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSite >= len(resource.Site) {
		return ReferenceInput("site["+strconv.Itoa(numSite)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("site["+strconv.Itoa(numSite)+"]", &resource.Site[numSite], htmlAttrs)
}
func (resource *ResearchStudy) T_ReasonStopped(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("reasonStopped", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonStopped", resource.ReasonStopped, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *ResearchStudy) T_ArmName(numArm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numArm >= len(resource.Arm) {
		return StringInput("arm["+strconv.Itoa(numArm)+"].name", nil, htmlAttrs)
	}
	return StringInput("arm["+strconv.Itoa(numArm)+"].name", &resource.Arm[numArm].Name, htmlAttrs)
}
func (resource *ResearchStudy) T_ArmType(numArm int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numArm >= len(resource.Arm) {
		return CodeableConceptSelect("arm["+strconv.Itoa(numArm)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("arm["+strconv.Itoa(numArm)+"].type", resource.Arm[numArm].Type, optionsValueSet, htmlAttrs)
}
func (resource *ResearchStudy) T_ArmDescription(numArm int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numArm >= len(resource.Arm) {
		return StringInput("arm["+strconv.Itoa(numArm)+"].description", nil, htmlAttrs)
	}
	return StringInput("arm["+strconv.Itoa(numArm)+"].description", resource.Arm[numArm].Description, htmlAttrs)
}
func (resource *ResearchStudy) T_ObjectiveName(numObjective int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numObjective >= len(resource.Objective) {
		return StringInput("objective["+strconv.Itoa(numObjective)+"].name", nil, htmlAttrs)
	}
	return StringInput("objective["+strconv.Itoa(numObjective)+"].name", resource.Objective[numObjective].Name, htmlAttrs)
}
func (resource *ResearchStudy) T_ObjectiveType(numObjective int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numObjective >= len(resource.Objective) {
		return CodeableConceptSelect("objective["+strconv.Itoa(numObjective)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("objective["+strconv.Itoa(numObjective)+"].type", resource.Objective[numObjective].Type, optionsValueSet, htmlAttrs)
}
