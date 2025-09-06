package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Requirements
type Requirements struct {
	Id                     *string                 `json:"id,omitempty"`
	Meta                   *Meta                   `json:"meta,omitempty"`
	ImplicitRules          *string                 `json:"implicitRules,omitempty"`
	Language               *string                 `json:"language,omitempty"`
	Text                   *Narrative              `json:"text,omitempty"`
	Contained              []Resource              `json:"contained,omitempty"`
	Extension              []Extension             `json:"extension,omitempty"`
	ModifierExtension      []Extension             `json:"modifierExtension,omitempty"`
	Url                    *string                 `json:"url,omitempty"`
	Identifier             []Identifier            `json:"identifier,omitempty"`
	Version                *string                 `json:"version,omitempty"`
	VersionAlgorithmString *string                 `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                 `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                 `json:"name,omitempty"`
	Title                  *string                 `json:"title,omitempty"`
	Status                 string                  `json:"status"`
	Experimental           *bool                   `json:"experimental,omitempty"`
	Date                   *string                 `json:"date,omitempty"`
	Publisher              *string                 `json:"publisher,omitempty"`
	Contact                []ContactDetail         `json:"contact,omitempty"`
	Description            *string                 `json:"description,omitempty"`
	UseContext             []UsageContext          `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept       `json:"jurisdiction,omitempty"`
	Purpose                *string                 `json:"purpose,omitempty"`
	Copyright              *string                 `json:"copyright,omitempty"`
	CopyrightLabel         *string                 `json:"copyrightLabel,omitempty"`
	DerivedFrom            []string                `json:"derivedFrom,omitempty"`
	Reference              []string                `json:"reference,omitempty"`
	Actor                  []string                `json:"actor,omitempty"`
	Statement              []RequirementsStatement `json:"statement,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Requirements
type RequirementsStatement struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Key               string      `json:"key"`
	Label             *string     `json:"label,omitempty"`
	Conformance       []string    `json:"conformance,omitempty"`
	Conditionality    *bool       `json:"conditionality,omitempty"`
	Requirement       string      `json:"requirement"`
	DerivedFrom       *string     `json:"derivedFrom,omitempty"`
	Parent            *string     `json:"parent,omitempty"`
	SatisfiedBy       []string    `json:"satisfiedBy,omitempty"`
	Reference         []string    `json:"reference,omitempty"`
	Source            []Reference `json:"source,omitempty"`
}

type OtherRequirements Requirements

// on convert struct to json, automatically add resourceType=Requirements
func (r Requirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRequirements
		ResourceType string `json:"resourceType"`
	}{
		OtherRequirements: OtherRequirements(r),
		ResourceType:      "Requirements",
	})
}

func (resource *Requirements) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Requirements.Id", nil)
	}
	return StringInput("Requirements.Id", resource.Id)
}
func (resource *Requirements) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Requirements.ImplicitRules", nil)
	}
	return StringInput("Requirements.ImplicitRules", resource.ImplicitRules)
}
func (resource *Requirements) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Requirements.Language", nil, optionsValueSet)
	}
	return CodeSelect("Requirements.Language", resource.Language, optionsValueSet)
}
func (resource *Requirements) T_Url() templ.Component {

	if resource == nil {
		return StringInput("Requirements.Url", nil)
	}
	return StringInput("Requirements.Url", resource.Url)
}
func (resource *Requirements) T_Version() templ.Component {

	if resource == nil {
		return StringInput("Requirements.Version", nil)
	}
	return StringInput("Requirements.Version", resource.Version)
}
func (resource *Requirements) T_Name() templ.Component {

	if resource == nil {
		return StringInput("Requirements.Name", nil)
	}
	return StringInput("Requirements.Name", resource.Name)
}
func (resource *Requirements) T_Title() templ.Component {

	if resource == nil {
		return StringInput("Requirements.Title", nil)
	}
	return StringInput("Requirements.Title", resource.Title)
}
func (resource *Requirements) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Requirements.Status", nil, optionsValueSet)
	}
	return CodeSelect("Requirements.Status", &resource.Status, optionsValueSet)
}
func (resource *Requirements) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("Requirements.Experimental", nil)
	}
	return BoolInput("Requirements.Experimental", resource.Experimental)
}
func (resource *Requirements) T_Date() templ.Component {

	if resource == nil {
		return StringInput("Requirements.Date", nil)
	}
	return StringInput("Requirements.Date", resource.Date)
}
func (resource *Requirements) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("Requirements.Publisher", nil)
	}
	return StringInput("Requirements.Publisher", resource.Publisher)
}
func (resource *Requirements) T_Description() templ.Component {

	if resource == nil {
		return StringInput("Requirements.Description", nil)
	}
	return StringInput("Requirements.Description", resource.Description)
}
func (resource *Requirements) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("Requirements.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Requirements.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *Requirements) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("Requirements.Purpose", nil)
	}
	return StringInput("Requirements.Purpose", resource.Purpose)
}
func (resource *Requirements) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("Requirements.Copyright", nil)
	}
	return StringInput("Requirements.Copyright", resource.Copyright)
}
func (resource *Requirements) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("Requirements.CopyrightLabel", nil)
	}
	return StringInput("Requirements.CopyrightLabel", resource.CopyrightLabel)
}
func (resource *Requirements) T_DerivedFrom(numDerivedFrom int) templ.Component {

	if resource == nil || len(resource.DerivedFrom) >= numDerivedFrom {
		return StringInput("Requirements.DerivedFrom["+strconv.Itoa(numDerivedFrom)+"]", nil)
	}
	return StringInput("Requirements.DerivedFrom["+strconv.Itoa(numDerivedFrom)+"]", &resource.DerivedFrom[numDerivedFrom])
}
func (resource *Requirements) T_Reference(numReference int) templ.Component {

	if resource == nil || len(resource.Reference) >= numReference {
		return StringInput("Requirements.Reference["+strconv.Itoa(numReference)+"]", nil)
	}
	return StringInput("Requirements.Reference["+strconv.Itoa(numReference)+"]", &resource.Reference[numReference])
}
func (resource *Requirements) T_Actor(numActor int) templ.Component {

	if resource == nil || len(resource.Actor) >= numActor {
		return StringInput("Requirements.Actor["+strconv.Itoa(numActor)+"]", nil)
	}
	return StringInput("Requirements.Actor["+strconv.Itoa(numActor)+"]", &resource.Actor[numActor])
}
func (resource *Requirements) T_StatementId(numStatement int) templ.Component {

	if resource == nil || len(resource.Statement) >= numStatement {
		return StringInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].Id", nil)
	}
	return StringInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].Id", resource.Statement[numStatement].Id)
}
func (resource *Requirements) T_StatementKey(numStatement int) templ.Component {

	if resource == nil || len(resource.Statement) >= numStatement {
		return StringInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].Key", nil)
	}
	return StringInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].Key", &resource.Statement[numStatement].Key)
}
func (resource *Requirements) T_StatementLabel(numStatement int) templ.Component {

	if resource == nil || len(resource.Statement) >= numStatement {
		return StringInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].Label", nil)
	}
	return StringInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].Label", resource.Statement[numStatement].Label)
}
func (resource *Requirements) T_StatementConformance(numStatement int, numConformance int) templ.Component {
	optionsValueSet := VSConformance_expectation

	if resource == nil || len(resource.Statement) >= numStatement || len(resource.Statement[numStatement].Conformance) >= numConformance {
		return CodeSelect("Requirements.Statement["+strconv.Itoa(numStatement)+"].Conformance["+strconv.Itoa(numConformance)+"]", nil, optionsValueSet)
	}
	return CodeSelect("Requirements.Statement["+strconv.Itoa(numStatement)+"].Conformance["+strconv.Itoa(numConformance)+"]", &resource.Statement[numStatement].Conformance[numConformance], optionsValueSet)
}
func (resource *Requirements) T_StatementConditionality(numStatement int) templ.Component {

	if resource == nil || len(resource.Statement) >= numStatement {
		return BoolInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].Conditionality", nil)
	}
	return BoolInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].Conditionality", resource.Statement[numStatement].Conditionality)
}
func (resource *Requirements) T_StatementRequirement(numStatement int) templ.Component {

	if resource == nil || len(resource.Statement) >= numStatement {
		return StringInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].Requirement", nil)
	}
	return StringInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].Requirement", &resource.Statement[numStatement].Requirement)
}
func (resource *Requirements) T_StatementDerivedFrom(numStatement int) templ.Component {

	if resource == nil || len(resource.Statement) >= numStatement {
		return StringInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].DerivedFrom", nil)
	}
	return StringInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].DerivedFrom", resource.Statement[numStatement].DerivedFrom)
}
func (resource *Requirements) T_StatementParent(numStatement int) templ.Component {

	if resource == nil || len(resource.Statement) >= numStatement {
		return StringInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].Parent", nil)
	}
	return StringInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].Parent", resource.Statement[numStatement].Parent)
}
func (resource *Requirements) T_StatementSatisfiedBy(numStatement int, numSatisfiedBy int) templ.Component {

	if resource == nil || len(resource.Statement) >= numStatement || len(resource.Statement[numStatement].SatisfiedBy) >= numSatisfiedBy {
		return StringInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].SatisfiedBy["+strconv.Itoa(numSatisfiedBy)+"]", nil)
	}
	return StringInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].SatisfiedBy["+strconv.Itoa(numSatisfiedBy)+"]", &resource.Statement[numStatement].SatisfiedBy[numSatisfiedBy])
}
func (resource *Requirements) T_StatementReference(numStatement int, numReference int) templ.Component {

	if resource == nil || len(resource.Statement) >= numStatement || len(resource.Statement[numStatement].Reference) >= numReference {
		return StringInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].Reference["+strconv.Itoa(numReference)+"]", nil)
	}
	return StringInput("Requirements.Statement["+strconv.Itoa(numStatement)+"].Reference["+strconv.Itoa(numReference)+"]", &resource.Statement[numStatement].Reference[numReference])
}
