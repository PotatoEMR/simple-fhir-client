package r5

//generated with command go run ./bultaoreune
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
func (r Requirements) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Requirements/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Requirements"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Requirements) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *Requirements) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *Requirements) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *Requirements) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *Requirements) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *Requirements) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *Requirements) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Requirements) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *Requirements) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *Requirements) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *Requirements) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *Requirements) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *Requirements) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *Requirements) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *Requirements) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *Requirements) T_DerivedFrom(numDerivedFrom int, htmlAttrs string) templ.Component {
	if resource == nil || numDerivedFrom >= len(resource.DerivedFrom) {
		return StringInput("derivedFrom["+strconv.Itoa(numDerivedFrom)+"]", nil, htmlAttrs)
	}
	return StringInput("derivedFrom["+strconv.Itoa(numDerivedFrom)+"]", &resource.DerivedFrom[numDerivedFrom], htmlAttrs)
}
func (resource *Requirements) T_Reference(numReference int, htmlAttrs string) templ.Component {
	if resource == nil || numReference >= len(resource.Reference) {
		return StringInput("reference["+strconv.Itoa(numReference)+"]", nil, htmlAttrs)
	}
	return StringInput("reference["+strconv.Itoa(numReference)+"]", &resource.Reference[numReference], htmlAttrs)
}
func (resource *Requirements) T_Actor(numActor int, htmlAttrs string) templ.Component {
	if resource == nil || numActor >= len(resource.Actor) {
		return StringInput("actor["+strconv.Itoa(numActor)+"]", nil, htmlAttrs)
	}
	return StringInput("actor["+strconv.Itoa(numActor)+"]", &resource.Actor[numActor], htmlAttrs)
}
func (resource *Requirements) T_StatementKey(numStatement int, htmlAttrs string) templ.Component {
	if resource == nil || numStatement >= len(resource.Statement) {
		return StringInput("statement["+strconv.Itoa(numStatement)+"].key", nil, htmlAttrs)
	}
	return StringInput("statement["+strconv.Itoa(numStatement)+"].key", &resource.Statement[numStatement].Key, htmlAttrs)
}
func (resource *Requirements) T_StatementLabel(numStatement int, htmlAttrs string) templ.Component {
	if resource == nil || numStatement >= len(resource.Statement) {
		return StringInput("statement["+strconv.Itoa(numStatement)+"].label", nil, htmlAttrs)
	}
	return StringInput("statement["+strconv.Itoa(numStatement)+"].label", resource.Statement[numStatement].Label, htmlAttrs)
}
func (resource *Requirements) T_StatementConformance(numStatement int, numConformance int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConformance_expectation

	if resource == nil || numStatement >= len(resource.Statement) || numConformance >= len(resource.Statement[numStatement].Conformance) {
		return CodeSelect("statement["+strconv.Itoa(numStatement)+"].conformance["+strconv.Itoa(numConformance)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("statement["+strconv.Itoa(numStatement)+"].conformance["+strconv.Itoa(numConformance)+"]", &resource.Statement[numStatement].Conformance[numConformance], optionsValueSet, htmlAttrs)
}
func (resource *Requirements) T_StatementConditionality(numStatement int, htmlAttrs string) templ.Component {
	if resource == nil || numStatement >= len(resource.Statement) {
		return BoolInput("statement["+strconv.Itoa(numStatement)+"].conditionality", nil, htmlAttrs)
	}
	return BoolInput("statement["+strconv.Itoa(numStatement)+"].conditionality", resource.Statement[numStatement].Conditionality, htmlAttrs)
}
func (resource *Requirements) T_StatementRequirement(numStatement int, htmlAttrs string) templ.Component {
	if resource == nil || numStatement >= len(resource.Statement) {
		return StringInput("statement["+strconv.Itoa(numStatement)+"].requirement", nil, htmlAttrs)
	}
	return StringInput("statement["+strconv.Itoa(numStatement)+"].requirement", &resource.Statement[numStatement].Requirement, htmlAttrs)
}
func (resource *Requirements) T_StatementDerivedFrom(numStatement int, htmlAttrs string) templ.Component {
	if resource == nil || numStatement >= len(resource.Statement) {
		return StringInput("statement["+strconv.Itoa(numStatement)+"].derivedFrom", nil, htmlAttrs)
	}
	return StringInput("statement["+strconv.Itoa(numStatement)+"].derivedFrom", resource.Statement[numStatement].DerivedFrom, htmlAttrs)
}
func (resource *Requirements) T_StatementParent(numStatement int, htmlAttrs string) templ.Component {
	if resource == nil || numStatement >= len(resource.Statement) {
		return StringInput("statement["+strconv.Itoa(numStatement)+"].parent", nil, htmlAttrs)
	}
	return StringInput("statement["+strconv.Itoa(numStatement)+"].parent", resource.Statement[numStatement].Parent, htmlAttrs)
}
func (resource *Requirements) T_StatementSatisfiedBy(numStatement int, numSatisfiedBy int, htmlAttrs string) templ.Component {
	if resource == nil || numStatement >= len(resource.Statement) || numSatisfiedBy >= len(resource.Statement[numStatement].SatisfiedBy) {
		return StringInput("statement["+strconv.Itoa(numStatement)+"].satisfiedBy["+strconv.Itoa(numSatisfiedBy)+"]", nil, htmlAttrs)
	}
	return StringInput("statement["+strconv.Itoa(numStatement)+"].satisfiedBy["+strconv.Itoa(numSatisfiedBy)+"]", &resource.Statement[numStatement].SatisfiedBy[numSatisfiedBy], htmlAttrs)
}
func (resource *Requirements) T_StatementReference(numStatement int, numReference int, htmlAttrs string) templ.Component {
	if resource == nil || numStatement >= len(resource.Statement) || numReference >= len(resource.Statement[numStatement].Reference) {
		return StringInput("statement["+strconv.Itoa(numStatement)+"].reference["+strconv.Itoa(numReference)+"]", nil, htmlAttrs)
	}
	return StringInput("statement["+strconv.Itoa(numStatement)+"].reference["+strconv.Itoa(numReference)+"]", &resource.Statement[numStatement].Reference[numReference], htmlAttrs)
}
