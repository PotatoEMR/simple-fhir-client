package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/SearchParameter
type SearchParameter struct {
	Id                     *string                    `json:"id,omitempty"`
	Meta                   *Meta                      `json:"meta,omitempty"`
	ImplicitRules          *string                    `json:"implicitRules,omitempty"`
	Language               *string                    `json:"language,omitempty"`
	Text                   *Narrative                 `json:"text,omitempty"`
	Contained              []Resource                 `json:"contained,omitempty"`
	Extension              []Extension                `json:"extension,omitempty"`
	ModifierExtension      []Extension                `json:"modifierExtension,omitempty"`
	Url                    string                     `json:"url"`
	Identifier             []Identifier               `json:"identifier,omitempty"`
	Version                *string                    `json:"version,omitempty"`
	VersionAlgorithmString *string                    `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                    `json:"versionAlgorithmCoding,omitempty"`
	Name                   string                     `json:"name"`
	Title                  *string                    `json:"title,omitempty"`
	DerivedFrom            *string                    `json:"derivedFrom,omitempty"`
	Status                 string                     `json:"status"`
	Experimental           *bool                      `json:"experimental,omitempty"`
	Date                   *string                    `json:"date,omitempty"`
	Publisher              *string                    `json:"publisher,omitempty"`
	Contact                []ContactDetail            `json:"contact,omitempty"`
	Description            string                     `json:"description"`
	UseContext             []UsageContext             `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept          `json:"jurisdiction,omitempty"`
	Purpose                *string                    `json:"purpose,omitempty"`
	Copyright              *string                    `json:"copyright,omitempty"`
	CopyrightLabel         *string                    `json:"copyrightLabel,omitempty"`
	Code                   string                     `json:"code"`
	Base                   []string                   `json:"base"`
	Type                   string                     `json:"type"`
	Expression             *string                    `json:"expression,omitempty"`
	ProcessingMode         *string                    `json:"processingMode,omitempty"`
	Constraint             *string                    `json:"constraint,omitempty"`
	Target                 []string                   `json:"target,omitempty"`
	MultipleOr             *bool                      `json:"multipleOr,omitempty"`
	MultipleAnd            *bool                      `json:"multipleAnd,omitempty"`
	Comparator             []string                   `json:"comparator,omitempty"`
	Modifier               []string                   `json:"modifier,omitempty"`
	Chain                  []string                   `json:"chain,omitempty"`
	Component              []SearchParameterComponent `json:"component,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SearchParameter
type SearchParameterComponent struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Definition        string      `json:"definition"`
	Expression        string      `json:"expression"`
}

type OtherSearchParameter SearchParameter

// on convert struct to json, automatically add resourceType=SearchParameter
func (r SearchParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSearchParameter
		ResourceType string `json:"resourceType"`
	}{
		OtherSearchParameter: OtherSearchParameter(r),
		ResourceType:         "SearchParameter",
	})
}
func (r SearchParameter) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "SearchParameter/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "SearchParameter"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *SearchParameter) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", &resource.Url, htmlAttrs)
}
func (resource *SearchParameter) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *SearchParameter) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *SearchParameter) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", &resource.Name, htmlAttrs)
}
func (resource *SearchParameter) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *SearchParameter) T_DerivedFrom(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("derivedFrom", nil, htmlAttrs)
	}
	return StringInput("derivedFrom", resource.DerivedFrom, htmlAttrs)
}
func (resource *SearchParameter) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *SearchParameter) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *SearchParameter) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *SearchParameter) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", &resource.Description, htmlAttrs)
}
func (resource *SearchParameter) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *SearchParameter) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *SearchParameter) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *SearchParameter) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Base(numBase int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBase >= len(resource.Base) {
		return CodeSelect("base["+strconv.Itoa(numBase)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("base["+strconv.Itoa(numBase)+"]", &resource.Base[numBase], optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Type(htmlAttrs string) templ.Component {
	optionsValueSet := VSSearch_param_type

	if resource == nil {
		return CodeSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Expression(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("expression", nil, htmlAttrs)
	}
	return StringInput("expression", resource.Expression, htmlAttrs)
}
func (resource *SearchParameter) T_ProcessingMode(htmlAttrs string) templ.Component {
	optionsValueSet := VSSearch_processingmode

	if resource == nil {
		return CodeSelect("processingMode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("processingMode", resource.ProcessingMode, optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Constraint(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("constraint", nil, htmlAttrs)
	}
	return StringInput("constraint", resource.Constraint, htmlAttrs)
}
func (resource *SearchParameter) T_Target(numTarget int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return CodeSelect("target["+strconv.Itoa(numTarget)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("target["+strconv.Itoa(numTarget)+"]", &resource.Target[numTarget], optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_MultipleOr(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("multipleOr", nil, htmlAttrs)
	}
	return BoolInput("multipleOr", resource.MultipleOr, htmlAttrs)
}
func (resource *SearchParameter) T_MultipleAnd(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("multipleAnd", nil, htmlAttrs)
	}
	return BoolInput("multipleAnd", resource.MultipleAnd, htmlAttrs)
}
func (resource *SearchParameter) T_Comparator(numComparator int, htmlAttrs string) templ.Component {
	optionsValueSet := VSSearch_comparator

	if resource == nil || numComparator >= len(resource.Comparator) {
		return CodeSelect("comparator["+strconv.Itoa(numComparator)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("comparator["+strconv.Itoa(numComparator)+"]", &resource.Comparator[numComparator], optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Modifier(numModifier int, htmlAttrs string) templ.Component {
	optionsValueSet := VSSearch_modifier_code

	if resource == nil || numModifier >= len(resource.Modifier) {
		return CodeSelect("modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("modifier["+strconv.Itoa(numModifier)+"]", &resource.Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Chain(numChain int, htmlAttrs string) templ.Component {
	if resource == nil || numChain >= len(resource.Chain) {
		return StringInput("chain["+strconv.Itoa(numChain)+"]", nil, htmlAttrs)
	}
	return StringInput("chain["+strconv.Itoa(numChain)+"]", &resource.Chain[numChain], htmlAttrs)
}
func (resource *SearchParameter) T_ComponentDefinition(numComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return StringInput("component["+strconv.Itoa(numComponent)+"].definition", nil, htmlAttrs)
	}
	return StringInput("component["+strconv.Itoa(numComponent)+"].definition", &resource.Component[numComponent].Definition, htmlAttrs)
}
func (resource *SearchParameter) T_ComponentExpression(numComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return StringInput("component["+strconv.Itoa(numComponent)+"].expression", nil, htmlAttrs)
	}
	return StringInput("component["+strconv.Itoa(numComponent)+"].expression", &resource.Component[numComponent].Expression, htmlAttrs)
}
