package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	VersionAlgorithmString *string                    `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                    `json:"versionAlgorithmCoding"`
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

func (resource *SearchParameter) SearchParameterLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *SearchParameter) SearchParameterStatus() templ.Component {
	optionsValueSet := VSPublication_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *SearchParameter) SearchParameterJurisdiction(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *SearchParameter) SearchParameterCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("code", nil, optionsValueSet)
	}
	return CodeSelect("code", &resource.Code, optionsValueSet)
}
func (resource *SearchParameter) SearchParameterBase(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("base", nil, optionsValueSet)
	}
	return CodeSelect("base", &resource.Base[0], optionsValueSet)
}
func (resource *SearchParameter) SearchParameterType() templ.Component {
	optionsValueSet := VSSearch_param_type

	if resource != nil {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", &resource.Type, optionsValueSet)
}
func (resource *SearchParameter) SearchParameterProcessingMode() templ.Component {
	optionsValueSet := VSSearch_processingmode

	if resource != nil {
		return CodeSelect("processingMode", nil, optionsValueSet)
	}
	return CodeSelect("processingMode", resource.ProcessingMode, optionsValueSet)
}
func (resource *SearchParameter) SearchParameterTarget(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("target", nil, optionsValueSet)
	}
	return CodeSelect("target", &resource.Target[0], optionsValueSet)
}
func (resource *SearchParameter) SearchParameterComparator() templ.Component {
	optionsValueSet := VSSearch_comparator

	if resource != nil {
		return CodeSelect("comparator", nil, optionsValueSet)
	}
	return CodeSelect("comparator", &resource.Comparator[0], optionsValueSet)
}
func (resource *SearchParameter) SearchParameterModifier() templ.Component {
	optionsValueSet := VSSearch_modifier_code

	if resource != nil {
		return CodeSelect("modifier", nil, optionsValueSet)
	}
	return CodeSelect("modifier", &resource.Modifier[0], optionsValueSet)
}
