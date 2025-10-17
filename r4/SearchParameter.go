package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/SearchParameter
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
	Date              *FhirDateTime              `json:"date,omitempty"`
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

// http://hl7.org/fhir/r4/StructureDefinition/SearchParameter
type SearchParameterComponent struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Definition        string      `json:"definition"`
	Expression        string      `json:"expression"`
}

type OtherSearchParameter SearchParameter

// struct -> json, automatically add resourceType=Patient
func (r SearchParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSearchParameter
		ResourceType string `json:"resourceType"`
	}{
		OtherSearchParameter: OtherSearchParameter(r),
		ResourceType:         "SearchParameter",
	})
}

// json -> struct, first reject if resourceType != SearchParameter
func (r *SearchParameter) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "SearchParameter" {
		return errors.New("resourceType not SearchParameter")
	}
	return json.Unmarshal(data, (*OtherSearchParameter)(r))
}

func (r SearchParameter) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "SearchParameter/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "SearchParameter"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *SearchParameter) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", &resource.Url, htmlAttrs)
}
func (resource *SearchParameter) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *SearchParameter) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", &resource.Name, htmlAttrs)
}
func (resource *SearchParameter) T_DerivedFrom(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("derivedFrom", nil, htmlAttrs)
	}
	return StringInput("derivedFrom", resource.DerivedFrom, htmlAttrs)
}
func (resource *SearchParameter) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *SearchParameter) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *SearchParameter) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *SearchParameter) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *SearchParameter) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", &resource.Description, htmlAttrs)
}
func (resource *SearchParameter) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *SearchParameter) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *SearchParameter) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Base(numBase int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numBase >= len(resource.Base) {
		return CodeSelect("base["+strconv.Itoa(numBase)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("base["+strconv.Itoa(numBase)+"]", &resource.Base[numBase], optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Type(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSearch_param_type

	if resource == nil {
		return CodeSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Expression(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("expression", nil, htmlAttrs)
	}
	return StringInput("expression", resource.Expression, htmlAttrs)
}
func (resource *SearchParameter) T_Xpath(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("xpath", nil, htmlAttrs)
	}
	return StringInput("xpath", resource.Xpath, htmlAttrs)
}
func (resource *SearchParameter) T_XpathUsage(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSearch_xpath_usage

	if resource == nil {
		return CodeSelect("xpathUsage", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("xpathUsage", resource.XpathUsage, optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Target(numTarget int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numTarget >= len(resource.Target) {
		return CodeSelect("target["+strconv.Itoa(numTarget)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("target["+strconv.Itoa(numTarget)+"]", &resource.Target[numTarget], optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_MultipleOr(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("multipleOr", nil, htmlAttrs)
	}
	return BoolInput("multipleOr", resource.MultipleOr, htmlAttrs)
}
func (resource *SearchParameter) T_MultipleAnd(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("multipleAnd", nil, htmlAttrs)
	}
	return BoolInput("multipleAnd", resource.MultipleAnd, htmlAttrs)
}
func (resource *SearchParameter) T_Comparator(numComparator int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSearch_comparator

	if resource == nil || numComparator >= len(resource.Comparator) {
		return CodeSelect("comparator["+strconv.Itoa(numComparator)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("comparator["+strconv.Itoa(numComparator)+"]", &resource.Comparator[numComparator], optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Modifier(numModifier int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSearch_modifier_code

	if resource == nil || numModifier >= len(resource.Modifier) {
		return CodeSelect("modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("modifier["+strconv.Itoa(numModifier)+"]", &resource.Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *SearchParameter) T_Chain(numChain int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numChain >= len(resource.Chain) {
		return StringInput("chain["+strconv.Itoa(numChain)+"]", nil, htmlAttrs)
	}
	return StringInput("chain["+strconv.Itoa(numChain)+"]", &resource.Chain[numChain], htmlAttrs)
}
func (resource *SearchParameter) T_ComponentDefinition(numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return StringInput("component["+strconv.Itoa(numComponent)+"].definition", nil, htmlAttrs)
	}
	return StringInput("component["+strconv.Itoa(numComponent)+"].definition", &resource.Component[numComponent].Definition, htmlAttrs)
}
func (resource *SearchParameter) T_ComponentExpression(numComponent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComponent >= len(resource.Component) {
		return StringInput("component["+strconv.Itoa(numComponent)+"].expression", nil, htmlAttrs)
	}
	return StringInput("component["+strconv.Itoa(numComponent)+"].expression", &resource.Component[numComponent].Expression, htmlAttrs)
}
