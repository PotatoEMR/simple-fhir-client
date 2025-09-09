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
	Date                   *time.Time                 `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
		return StringInput("SearchParameter.Url", nil, htmlAttrs)
	}
	return StringInput("SearchParameter.Url", &resource.Url, htmlAttrs)
}
func (resource *SearchParameter) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SearchParameter.Version", nil, htmlAttrs)
	}
	return StringInput("SearchParameter.Version", resource.Version, htmlAttrs)
}
func (resource *SearchParameter) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SearchParameter.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("SearchParameter.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *SearchParameter) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("SearchParameter.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("SearchParameter.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SearchParameter.Name", nil, htmlAttrs)
	}
	return StringInput("SearchParameter.Name", &resource.Name, htmlAttrs)
}
func (resource *SearchParameter) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SearchParameter.Title", nil, htmlAttrs)
	}
	return StringInput("SearchParameter.Title", resource.Title, htmlAttrs)
}
func (resource *SearchParameter) T_DerivedFrom(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SearchParameter.DerivedFrom", nil, htmlAttrs)
	}
	return StringInput("SearchParameter.DerivedFrom", resource.DerivedFrom, htmlAttrs)
}
func (resource *SearchParameter) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("SearchParameter.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SearchParameter.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("SearchParameter.Experimental", nil, htmlAttrs)
	}
	return BoolInput("SearchParameter.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *SearchParameter) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("SearchParameter.Date", nil, htmlAttrs)
	}
	return DateTimeInput("SearchParameter.Date", resource.Date, htmlAttrs)
}
func (resource *SearchParameter) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SearchParameter.Publisher", nil, htmlAttrs)
	}
	return StringInput("SearchParameter.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *SearchParameter) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SearchParameter.Description", nil, htmlAttrs)
	}
	return StringInput("SearchParameter.Description", &resource.Description, htmlAttrs)
}
func (resource *SearchParameter) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("SearchParameter.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SearchParameter.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SearchParameter.Purpose", nil, htmlAttrs)
	}
	return StringInput("SearchParameter.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *SearchParameter) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SearchParameter.Copyright", nil, htmlAttrs)
	}
	return StringInput("SearchParameter.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *SearchParameter) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SearchParameter.CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("SearchParameter.CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *SearchParameter) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeSelect("SearchParameter.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SearchParameter.Code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Base(numBase int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBase >= len(resource.Base) {
		return CodeSelect("SearchParameter.Base["+strconv.Itoa(numBase)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SearchParameter.Base["+strconv.Itoa(numBase)+"]", &resource.Base[numBase], optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Type(htmlAttrs string) templ.Component {
	optionsValueSet := VSSearch_param_type

	if resource == nil {
		return CodeSelect("SearchParameter.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SearchParameter.Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Expression(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SearchParameter.Expression", nil, htmlAttrs)
	}
	return StringInput("SearchParameter.Expression", resource.Expression, htmlAttrs)
}
func (resource *SearchParameter) T_ProcessingMode(htmlAttrs string) templ.Component {
	optionsValueSet := VSSearch_processingmode

	if resource == nil {
		return CodeSelect("SearchParameter.ProcessingMode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SearchParameter.ProcessingMode", resource.ProcessingMode, optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Constraint(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SearchParameter.Constraint", nil, htmlAttrs)
	}
	return StringInput("SearchParameter.Constraint", resource.Constraint, htmlAttrs)
}
func (resource *SearchParameter) T_Target(numTarget int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return CodeSelect("SearchParameter.Target["+strconv.Itoa(numTarget)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SearchParameter.Target["+strconv.Itoa(numTarget)+"]", &resource.Target[numTarget], optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_MultipleOr(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("SearchParameter.MultipleOr", nil, htmlAttrs)
	}
	return BoolInput("SearchParameter.MultipleOr", resource.MultipleOr, htmlAttrs)
}
func (resource *SearchParameter) T_MultipleAnd(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("SearchParameter.MultipleAnd", nil, htmlAttrs)
	}
	return BoolInput("SearchParameter.MultipleAnd", resource.MultipleAnd, htmlAttrs)
}
func (resource *SearchParameter) T_Comparator(numComparator int, htmlAttrs string) templ.Component {
	optionsValueSet := VSSearch_comparator

	if resource == nil || numComparator >= len(resource.Comparator) {
		return CodeSelect("SearchParameter.Comparator["+strconv.Itoa(numComparator)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SearchParameter.Comparator["+strconv.Itoa(numComparator)+"]", &resource.Comparator[numComparator], optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Modifier(numModifier int, htmlAttrs string) templ.Component {
	optionsValueSet := VSSearch_modifier_code

	if resource == nil || numModifier >= len(resource.Modifier) {
		return CodeSelect("SearchParameter.Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SearchParameter.Modifier["+strconv.Itoa(numModifier)+"]", &resource.Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Chain(numChain int, htmlAttrs string) templ.Component {
	if resource == nil || numChain >= len(resource.Chain) {
		return StringInput("SearchParameter.Chain["+strconv.Itoa(numChain)+"]", nil, htmlAttrs)
	}
	return StringInput("SearchParameter.Chain["+strconv.Itoa(numChain)+"]", &resource.Chain[numChain], htmlAttrs)
}
func (resource *SearchParameter) T_ComponentDefinition(numComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return StringInput("SearchParameter.Component["+strconv.Itoa(numComponent)+"].Definition", nil, htmlAttrs)
	}
	return StringInput("SearchParameter.Component["+strconv.Itoa(numComponent)+"].Definition", &resource.Component[numComponent].Definition, htmlAttrs)
}
func (resource *SearchParameter) T_ComponentExpression(numComponent int, htmlAttrs string) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return StringInput("SearchParameter.Component["+strconv.Itoa(numComponent)+"].Expression", nil, htmlAttrs)
	}
	return StringInput("SearchParameter.Component["+strconv.Itoa(numComponent)+"].Expression", &resource.Component[numComponent].Expression, htmlAttrs)
}
