package r4b

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/OperationDefinition
type OperationDefinition struct {
	Id                *string                        `json:"id,omitempty"`
	Meta              *Meta                          `json:"meta,omitempty"`
	ImplicitRules     *string                        `json:"implicitRules,omitempty"`
	Language          *string                        `json:"language,omitempty"`
	Text              *Narrative                     `json:"text,omitempty"`
	Contained         []Resource                     `json:"contained,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Url               *string                        `json:"url,omitempty"`
	Version           *string                        `json:"version,omitempty"`
	Name              string                         `json:"name"`
	Title             *string                        `json:"title,omitempty"`
	Status            string                         `json:"status"`
	Kind              string                         `json:"kind"`
	Experimental      *bool                          `json:"experimental,omitempty"`
	Date              *string                        `json:"date,omitempty"`
	Publisher         *string                        `json:"publisher,omitempty"`
	Contact           []ContactDetail                `json:"contact,omitempty"`
	Description       *string                        `json:"description,omitempty"`
	UseContext        []UsageContext                 `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept              `json:"jurisdiction,omitempty"`
	Purpose           *string                        `json:"purpose,omitempty"`
	AffectsState      *bool                          `json:"affectsState,omitempty"`
	Code              string                         `json:"code"`
	Comment           *string                        `json:"comment,omitempty"`
	Base              *string                        `json:"base,omitempty"`
	Resource          []string                       `json:"resource,omitempty"`
	System            bool                           `json:"system"`
	Type              bool                           `json:"type"`
	Instance          bool                           `json:"instance"`
	InputProfile      *string                        `json:"inputProfile,omitempty"`
	OutputProfile     *string                        `json:"outputProfile,omitempty"`
	Parameter         []OperationDefinitionParameter `json:"parameter,omitempty"`
	Overload          []OperationDefinitionOverload  `json:"overload,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/OperationDefinition
type OperationDefinitionParameter struct {
	Id                *string                                      `json:"id,omitempty"`
	Extension         []Extension                                  `json:"extension,omitempty"`
	ModifierExtension []Extension                                  `json:"modifierExtension,omitempty"`
	Name              string                                       `json:"name"`
	Use               string                                       `json:"use"`
	Min               int                                          `json:"min"`
	Max               string                                       `json:"max"`
	Documentation     *string                                      `json:"documentation,omitempty"`
	Type              *string                                      `json:"type,omitempty"`
	TargetProfile     []string                                     `json:"targetProfile,omitempty"`
	SearchType        *string                                      `json:"searchType,omitempty"`
	Binding           *OperationDefinitionParameterBinding         `json:"binding,omitempty"`
	ReferencedFrom    []OperationDefinitionParameterReferencedFrom `json:"referencedFrom,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/OperationDefinition
type OperationDefinitionParameterBinding struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Strength          string      `json:"strength"`
	ValueSet          string      `json:"valueSet"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/OperationDefinition
type OperationDefinitionParameterReferencedFrom struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Source            string      `json:"source"`
	SourceId          *string     `json:"sourceId,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/OperationDefinition
type OperationDefinitionOverload struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ParameterName     []string    `json:"parameterName,omitempty"`
	Comment           *string     `json:"comment,omitempty"`
}

type OtherOperationDefinition OperationDefinition

// on convert struct to json, automatically add resourceType=OperationDefinition
func (r OperationDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOperationDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherOperationDefinition: OtherOperationDefinition(r),
		ResourceType:             "OperationDefinition",
	})
}

func (resource *OperationDefinition) OperationDefinitionLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *OperationDefinition) OperationDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *OperationDefinition) OperationDefinitionKind() templ.Component {
	optionsValueSet := VSOperation_kind
	currentVal := ""
	if resource != nil {
		currentVal = resource.Kind
	}
	return CodeSelect("kind", currentVal, optionsValueSet)
}
func (resource *OperationDefinition) OperationDefinitionCode(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = resource.Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
func (resource *OperationDefinition) OperationDefinitionResource() templ.Component {
	optionsValueSet := VSResource_types
	currentVal := ""
	if resource != nil {
		currentVal = resource.Resource[0]
	}
	return CodeSelect("resource", currentVal, optionsValueSet)
}
func (resource *OperationDefinition) OperationDefinitionParameterName(numParameter int, optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil && len(resource.Parameter) >= numParameter {
		currentVal = resource.Parameter[numParameter].Name
	}
	return CodeSelect("name", currentVal, optionsValueSet)
}
func (resource *OperationDefinition) OperationDefinitionParameterUse(numParameter int) templ.Component {
	optionsValueSet := VSOperation_parameter_use
	currentVal := ""
	if resource != nil && len(resource.Parameter) >= numParameter {
		currentVal = resource.Parameter[numParameter].Use
	}
	return CodeSelect("use", currentVal, optionsValueSet)
}
func (resource *OperationDefinition) OperationDefinitionParameterType(numParameter int, optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil && len(resource.Parameter) >= numParameter {
		currentVal = *resource.Parameter[numParameter].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
func (resource *OperationDefinition) OperationDefinitionParameterSearchType(numParameter int) templ.Component {
	optionsValueSet := VSSearch_param_type
	currentVal := ""
	if resource != nil && len(resource.Parameter) >= numParameter {
		currentVal = *resource.Parameter[numParameter].SearchType
	}
	return CodeSelect("searchType", currentVal, optionsValueSet)
}
func (resource *OperationDefinition) OperationDefinitionParameterBindingStrength(numParameter int) templ.Component {
	optionsValueSet := VSBinding_strength
	currentVal := ""
	if resource != nil && len(resource.Parameter) >= numParameter {
		currentVal = resource.Parameter[numParameter].Binding.Strength
	}
	return CodeSelect("strength", currentVal, optionsValueSet)
}
