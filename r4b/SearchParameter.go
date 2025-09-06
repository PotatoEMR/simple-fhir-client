package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/SearchParameter
type SearchParameter struct {
	Id                *string                    `json:"id,omitempty"`
	Meta              *Meta                      `json:"meta,omitempty"`
	ImplicitRules     *string                    `json:"implicitRules,omitempty"`
	Language          *string                    `json:"language,omitempty"`
	Text              *Narrative                 `json:"text,omitempty"`
	Contained         []Resource                 `json:"contained,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Url               string                     `json:"url"`
	Version           *string                    `json:"version,omitempty"`
	Name              string                     `json:"name"`
	DerivedFrom       *string                    `json:"derivedFrom,omitempty"`
	Status            string                     `json:"status"`
	Experimental      *bool                      `json:"experimental,omitempty"`
	Date              *string                    `json:"date,omitempty"`
	Publisher         *string                    `json:"publisher,omitempty"`
	Contact           []ContactDetail            `json:"contact,omitempty"`
	Description       string                     `json:"description"`
	UseContext        []UsageContext             `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept          `json:"jurisdiction,omitempty"`
	Purpose           *string                    `json:"purpose,omitempty"`
	Code              string                     `json:"code"`
	Base              []string                   `json:"base"`
	Type              string                     `json:"type"`
	Expression        *string                    `json:"expression,omitempty"`
	Xpath             *string                    `json:"xpath,omitempty"`
	XpathUsage        *string                    `json:"xpathUsage,omitempty"`
	Target            []string                   `json:"target,omitempty"`
	MultipleOr        *bool                      `json:"multipleOr,omitempty"`
	MultipleAnd       *bool                      `json:"multipleAnd,omitempty"`
	Comparator        []string                   `json:"comparator,omitempty"`
	Modifier          []string                   `json:"modifier,omitempty"`
	Chain             []string                   `json:"chain,omitempty"`
	Component         []SearchParameterComponent `json:"component,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SearchParameter
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

func (resource *SearchParameter) T_Id() templ.Component {

	if resource == nil {
		return StringInput("SearchParameter.Id", nil)
	}
	return StringInput("SearchParameter.Id", resource.Id)
}
func (resource *SearchParameter) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("SearchParameter.ImplicitRules", nil)
	}
	return StringInput("SearchParameter.ImplicitRules", resource.ImplicitRules)
}
func (resource *SearchParameter) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("SearchParameter.Language", nil, optionsValueSet)
	}
	return CodeSelect("SearchParameter.Language", resource.Language, optionsValueSet)
}
func (resource *SearchParameter) T_Url() templ.Component {

	if resource == nil {
		return StringInput("SearchParameter.Url", nil)
	}
	return StringInput("SearchParameter.Url", &resource.Url)
}
func (resource *SearchParameter) T_Version() templ.Component {

	if resource == nil {
		return StringInput("SearchParameter.Version", nil)
	}
	return StringInput("SearchParameter.Version", resource.Version)
}
func (resource *SearchParameter) T_Name() templ.Component {

	if resource == nil {
		return StringInput("SearchParameter.Name", nil)
	}
	return StringInput("SearchParameter.Name", &resource.Name)
}
func (resource *SearchParameter) T_DerivedFrom() templ.Component {

	if resource == nil {
		return StringInput("SearchParameter.DerivedFrom", nil)
	}
	return StringInput("SearchParameter.DerivedFrom", resource.DerivedFrom)
}
func (resource *SearchParameter) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("SearchParameter.Status", nil, optionsValueSet)
	}
	return CodeSelect("SearchParameter.Status", &resource.Status, optionsValueSet)
}
func (resource *SearchParameter) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("SearchParameter.Experimental", nil)
	}
	return BoolInput("SearchParameter.Experimental", resource.Experimental)
}
func (resource *SearchParameter) T_Date() templ.Component {

	if resource == nil {
		return StringInput("SearchParameter.Date", nil)
	}
	return StringInput("SearchParameter.Date", resource.Date)
}
func (resource *SearchParameter) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("SearchParameter.Publisher", nil)
	}
	return StringInput("SearchParameter.Publisher", resource.Publisher)
}
func (resource *SearchParameter) T_Description() templ.Component {

	if resource == nil {
		return StringInput("SearchParameter.Description", nil)
	}
	return StringInput("SearchParameter.Description", &resource.Description)
}
func (resource *SearchParameter) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("SearchParameter.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SearchParameter.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *SearchParameter) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("SearchParameter.Purpose", nil)
	}
	return StringInput("SearchParameter.Purpose", resource.Purpose)
}
func (resource *SearchParameter) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("SearchParameter.Code", nil, optionsValueSet)
	}
	return CodeSelect("SearchParameter.Code", &resource.Code, optionsValueSet)
}
func (resource *SearchParameter) T_Base(numBase int) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || len(resource.Base) >= numBase {
		return CodeSelect("SearchParameter.Base["+strconv.Itoa(numBase)+"]", nil, optionsValueSet)
	}
	return CodeSelect("SearchParameter.Base["+strconv.Itoa(numBase)+"]", &resource.Base[numBase], optionsValueSet)
}
func (resource *SearchParameter) T_Type() templ.Component {
	optionsValueSet := VSSearch_param_type

	if resource == nil {
		return CodeSelect("SearchParameter.Type", nil, optionsValueSet)
	}
	return CodeSelect("SearchParameter.Type", &resource.Type, optionsValueSet)
}
func (resource *SearchParameter) T_Expression() templ.Component {

	if resource == nil {
		return StringInput("SearchParameter.Expression", nil)
	}
	return StringInput("SearchParameter.Expression", resource.Expression)
}
func (resource *SearchParameter) T_Xpath() templ.Component {

	if resource == nil {
		return StringInput("SearchParameter.Xpath", nil)
	}
	return StringInput("SearchParameter.Xpath", resource.Xpath)
}
func (resource *SearchParameter) T_XpathUsage() templ.Component {
	optionsValueSet := VSSearch_xpath_usage

	if resource == nil {
		return CodeSelect("SearchParameter.XpathUsage", nil, optionsValueSet)
	}
	return CodeSelect("SearchParameter.XpathUsage", resource.XpathUsage, optionsValueSet)
}
func (resource *SearchParameter) T_Target(numTarget int) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || len(resource.Target) >= numTarget {
		return CodeSelect("SearchParameter.Target["+strconv.Itoa(numTarget)+"]", nil, optionsValueSet)
	}
	return CodeSelect("SearchParameter.Target["+strconv.Itoa(numTarget)+"]", &resource.Target[numTarget], optionsValueSet)
}
func (resource *SearchParameter) T_MultipleOr() templ.Component {

	if resource == nil {
		return BoolInput("SearchParameter.MultipleOr", nil)
	}
	return BoolInput("SearchParameter.MultipleOr", resource.MultipleOr)
}
func (resource *SearchParameter) T_MultipleAnd() templ.Component {

	if resource == nil {
		return BoolInput("SearchParameter.MultipleAnd", nil)
	}
	return BoolInput("SearchParameter.MultipleAnd", resource.MultipleAnd)
}
func (resource *SearchParameter) T_Comparator(numComparator int) templ.Component {
	optionsValueSet := VSSearch_comparator

	if resource == nil || len(resource.Comparator) >= numComparator {
		return CodeSelect("SearchParameter.Comparator["+strconv.Itoa(numComparator)+"]", nil, optionsValueSet)
	}
	return CodeSelect("SearchParameter.Comparator["+strconv.Itoa(numComparator)+"]", &resource.Comparator[numComparator], optionsValueSet)
}
func (resource *SearchParameter) T_Modifier(numModifier int) templ.Component {
	optionsValueSet := VSSearch_modifier_code

	if resource == nil || len(resource.Modifier) >= numModifier {
		return CodeSelect("SearchParameter.Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet)
	}
	return CodeSelect("SearchParameter.Modifier["+strconv.Itoa(numModifier)+"]", &resource.Modifier[numModifier], optionsValueSet)
}
func (resource *SearchParameter) T_Chain(numChain int) templ.Component {

	if resource == nil || len(resource.Chain) >= numChain {
		return StringInput("SearchParameter.Chain["+strconv.Itoa(numChain)+"]", nil)
	}
	return StringInput("SearchParameter.Chain["+strconv.Itoa(numChain)+"]", &resource.Chain[numChain])
}
func (resource *SearchParameter) T_ComponentId(numComponent int) templ.Component {

	if resource == nil || len(resource.Component) >= numComponent {
		return StringInput("SearchParameter.Component["+strconv.Itoa(numComponent)+"].Id", nil)
	}
	return StringInput("SearchParameter.Component["+strconv.Itoa(numComponent)+"].Id", resource.Component[numComponent].Id)
}
func (resource *SearchParameter) T_ComponentDefinition(numComponent int) templ.Component {

	if resource == nil || len(resource.Component) >= numComponent {
		return StringInput("SearchParameter.Component["+strconv.Itoa(numComponent)+"].Definition", nil)
	}
	return StringInput("SearchParameter.Component["+strconv.Itoa(numComponent)+"].Definition", &resource.Component[numComponent].Definition)
}
func (resource *SearchParameter) T_ComponentExpression(numComponent int) templ.Component {

	if resource == nil || len(resource.Component) >= numComponent {
		return StringInput("SearchParameter.Component["+strconv.Itoa(numComponent)+"].Expression", nil)
	}
	return StringInput("SearchParameter.Component["+strconv.Itoa(numComponent)+"].Expression", &resource.Component[numComponent].Expression)
}
