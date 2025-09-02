package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/OperationDefinition
type OperationDefinition struct {
	Id                     *string                        `json:"id,omitempty"`
	Meta                   *Meta                          `json:"meta,omitempty"`
	ImplicitRules          *string                        `json:"implicitRules,omitempty"`
	Language               *string                        `json:"language,omitempty"`
	Text                   *Narrative                     `json:"text,omitempty"`
	Contained              []Resource                     `json:"contained,omitempty"`
	Extension              []Extension                    `json:"extension,omitempty"`
	ModifierExtension      []Extension                    `json:"modifierExtension,omitempty"`
	Url                    *string                        `json:"url,omitempty"`
	Identifier             []Identifier                   `json:"identifier,omitempty"`
	Version                *string                        `json:"version,omitempty"`
	VersionAlgorithmString *string                        `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                        `json:"versionAlgorithmCoding"`
	Name                   string                         `json:"name"`
	Title                  *string                        `json:"title,omitempty"`
	Status                 string                         `json:"status"`
	Kind                   string                         `json:"kind"`
	Experimental           *bool                          `json:"experimental,omitempty"`
	Date                   *string                        `json:"date,omitempty"`
	Publisher              *string                        `json:"publisher,omitempty"`
	Contact                []ContactDetail                `json:"contact,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	UseContext             []UsageContext                 `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept              `json:"jurisdiction,omitempty"`
	Purpose                *string                        `json:"purpose,omitempty"`
	Copyright              *string                        `json:"copyright,omitempty"`
	CopyrightLabel         *string                        `json:"copyrightLabel,omitempty"`
	AffectsState           *bool                          `json:"affectsState,omitempty"`
	Code                   string                         `json:"code"`
	Comment                *string                        `json:"comment,omitempty"`
	Base                   *string                        `json:"base,omitempty"`
	Resource               []string                       `json:"resource,omitempty"`
	System                 bool                           `json:"system"`
	Type                   bool                           `json:"type"`
	Instance               bool                           `json:"instance"`
	InputProfile           *string                        `json:"inputProfile,omitempty"`
	OutputProfile          *string                        `json:"outputProfile,omitempty"`
	Parameter              []OperationDefinitionParameter `json:"parameter,omitempty"`
	Overload               []OperationDefinitionOverload  `json:"overload,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/OperationDefinition
type OperationDefinitionParameter struct {
	Id                *string                                      `json:"id,omitempty"`
	Extension         []Extension                                  `json:"extension,omitempty"`
	ModifierExtension []Extension                                  `json:"modifierExtension,omitempty"`
	Name              string                                       `json:"name"`
	Use               string                                       `json:"use"`
	Scope             []string                                     `json:"scope,omitempty"`
	Min               int                                          `json:"min"`
	Max               string                                       `json:"max"`
	Documentation     *string                                      `json:"documentation,omitempty"`
	Type              *string                                      `json:"type,omitempty"`
	AllowedType       []string                                     `json:"allowedType,omitempty"`
	TargetProfile     []string                                     `json:"targetProfile,omitempty"`
	SearchType        *string                                      `json:"searchType,omitempty"`
	Binding           *OperationDefinitionParameterBinding         `json:"binding,omitempty"`
	ReferencedFrom    []OperationDefinitionParameterReferencedFrom `json:"referencedFrom,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/OperationDefinition
type OperationDefinitionParameterBinding struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Strength          string      `json:"strength"`
	ValueSet          string      `json:"valueSet"`
}

// http://hl7.org/fhir/r5/StructureDefinition/OperationDefinition
type OperationDefinitionParameterReferencedFrom struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Source            string      `json:"source"`
	SourceId          *string     `json:"sourceId,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/OperationDefinition
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

func (resource *OperationDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *OperationDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *OperationDefinition) T_Kind() templ.Component {
	optionsValueSet := VSOperation_kind

	if resource == nil {
		return CodeSelect("kind", nil, optionsValueSet)
	}
	return CodeSelect("kind", &resource.Kind, optionsValueSet)
}
func (resource *OperationDefinition) T_Jurisdiction(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *OperationDefinition) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("code", nil, optionsValueSet)
	}
	return CodeSelect("code", &resource.Code, optionsValueSet)
}
func (resource *OperationDefinition) T_Resource(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("resource", nil, optionsValueSet)
	}
	return CodeSelect("resource", &resource.Resource[0], optionsValueSet)
}
func (resource *OperationDefinition) T_ParameterName(numParameter int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Parameter) >= numParameter {
		return CodeSelect("name", nil, optionsValueSet)
	}
	return CodeSelect("name", &resource.Parameter[numParameter].Name, optionsValueSet)
}
func (resource *OperationDefinition) T_ParameterUse(numParameter int) templ.Component {
	optionsValueSet := VSOperation_parameter_use

	if resource == nil && len(resource.Parameter) >= numParameter {
		return CodeSelect("use", nil, optionsValueSet)
	}
	return CodeSelect("use", &resource.Parameter[numParameter].Use, optionsValueSet)
}
func (resource *OperationDefinition) T_ParameterScope(numParameter int) templ.Component {
	optionsValueSet := VSOperation_parameter_scope

	if resource == nil && len(resource.Parameter) >= numParameter {
		return CodeSelect("scope", nil, optionsValueSet)
	}
	return CodeSelect("scope", &resource.Parameter[numParameter].Scope[0], optionsValueSet)
}
func (resource *OperationDefinition) T_ParameterType(numParameter int) templ.Component {
	optionsValueSet := VSFhir_types

	if resource == nil && len(resource.Parameter) >= numParameter {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", resource.Parameter[numParameter].Type, optionsValueSet)
}
func (resource *OperationDefinition) T_ParameterAllowedType(numParameter int) templ.Component {
	optionsValueSet := VSFhir_types

	if resource == nil && len(resource.Parameter) >= numParameter {
		return CodeSelect("allowedType", nil, optionsValueSet)
	}
	return CodeSelect("allowedType", &resource.Parameter[numParameter].AllowedType[0], optionsValueSet)
}
func (resource *OperationDefinition) T_ParameterSearchType(numParameter int) templ.Component {
	optionsValueSet := VSSearch_param_type

	if resource == nil && len(resource.Parameter) >= numParameter {
		return CodeSelect("searchType", nil, optionsValueSet)
	}
	return CodeSelect("searchType", resource.Parameter[numParameter].SearchType, optionsValueSet)
}
func (resource *OperationDefinition) T_ParameterBindingStrength(numParameter int) templ.Component {
	optionsValueSet := VSBinding_strength

	if resource == nil && len(resource.Parameter) >= numParameter {
		return CodeSelect("strength", nil, optionsValueSet)
	}
	return CodeSelect("strength", &resource.Parameter[numParameter].Binding.Strength, optionsValueSet)
}
