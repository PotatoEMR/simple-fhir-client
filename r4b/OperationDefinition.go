package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	Date              *FhirDateTime                  `json:"date,omitempty"`
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
func (r OperationDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "OperationDefinition/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "OperationDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *OperationDefinition) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *OperationDefinition) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *OperationDefinition) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", &resource.Name, htmlAttrs)
}
func (resource *OperationDefinition) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *OperationDefinition) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_Kind(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSOperation_kind

	if resource == nil {
		return CodeSelect("kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("kind", &resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *OperationDefinition) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *OperationDefinition) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *OperationDefinition) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *OperationDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *OperationDefinition) T_AffectsState(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("affectsState", nil, htmlAttrs)
	}
	return BoolInput("affectsState", resource.AffectsState, htmlAttrs)
}
func (resource *OperationDefinition) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_Comment(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("comment", nil, htmlAttrs)
	}
	return StringInput("comment", resource.Comment, htmlAttrs)
}
func (resource *OperationDefinition) T_Base(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("base", nil, htmlAttrs)
	}
	return StringInput("base", resource.Base, htmlAttrs)
}
func (resource *OperationDefinition) T_Resource(numResource int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numResource >= len(resource.Resource) {
		return CodeSelect("resource["+strconv.Itoa(numResource)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("resource["+strconv.Itoa(numResource)+"]", &resource.Resource[numResource], optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_System(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("system", nil, htmlAttrs)
	}
	return BoolInput("system", &resource.System, htmlAttrs)
}
func (resource *OperationDefinition) T_Type(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("type", nil, htmlAttrs)
	}
	return BoolInput("type", &resource.Type, htmlAttrs)
}
func (resource *OperationDefinition) T_Instance(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("instance", nil, htmlAttrs)
	}
	return BoolInput("instance", &resource.Instance, htmlAttrs)
}
func (resource *OperationDefinition) T_InputProfile(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("inputProfile", nil, htmlAttrs)
	}
	return StringInput("inputProfile", resource.InputProfile, htmlAttrs)
}
func (resource *OperationDefinition) T_OutputProfile(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("outputProfile", nil, htmlAttrs)
	}
	return StringInput("outputProfile", resource.OutputProfile, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterName(numParameter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeSelect("parameter["+strconv.Itoa(numParameter)+"].name", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("parameter["+strconv.Itoa(numParameter)+"].name", &resource.Parameter[numParameter].Name, optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterUse(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSOperation_parameter_use

	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeSelect("parameter["+strconv.Itoa(numParameter)+"].use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("parameter["+strconv.Itoa(numParameter)+"].use", &resource.Parameter[numParameter].Use, optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterMin(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return IntInput("parameter["+strconv.Itoa(numParameter)+"].min", nil, htmlAttrs)
	}
	return IntInput("parameter["+strconv.Itoa(numParameter)+"].min", &resource.Parameter[numParameter].Min, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterMax(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].max", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].max", &resource.Parameter[numParameter].Max, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterDocumentation(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].documentation", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].documentation", resource.Parameter[numParameter].Documentation, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterType(numParameter int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeSelect("parameter["+strconv.Itoa(numParameter)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("parameter["+strconv.Itoa(numParameter)+"].type", resource.Parameter[numParameter].Type, optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterTargetProfile(numParameter int, numTargetProfile int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) || numTargetProfile >= len(resource.Parameter[numParameter].TargetProfile) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].targetProfile["+strconv.Itoa(numTargetProfile)+"]", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].targetProfile["+strconv.Itoa(numTargetProfile)+"]", &resource.Parameter[numParameter].TargetProfile[numTargetProfile], htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterSearchType(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSearch_param_type

	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeSelect("parameter["+strconv.Itoa(numParameter)+"].searchType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("parameter["+strconv.Itoa(numParameter)+"].searchType", resource.Parameter[numParameter].SearchType, optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterBindingStrength(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSBinding_strength

	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeSelect("parameter["+strconv.Itoa(numParameter)+"].binding.strength", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("parameter["+strconv.Itoa(numParameter)+"].binding.strength", &resource.Parameter[numParameter].Binding.Strength, optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterBindingValueSet(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].binding.valueSet", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].binding.valueSet", &resource.Parameter[numParameter].Binding.ValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterReferencedFromSource(numParameter int, numReferencedFrom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) || numReferencedFrom >= len(resource.Parameter[numParameter].ReferencedFrom) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].referencedFrom["+strconv.Itoa(numReferencedFrom)+"].source", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].referencedFrom["+strconv.Itoa(numReferencedFrom)+"].source", &resource.Parameter[numParameter].ReferencedFrom[numReferencedFrom].Source, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterReferencedFromSourceId(numParameter int, numReferencedFrom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) || numReferencedFrom >= len(resource.Parameter[numParameter].ReferencedFrom) {
		return StringInput("parameter["+strconv.Itoa(numParameter)+"].referencedFrom["+strconv.Itoa(numReferencedFrom)+"].sourceId", nil, htmlAttrs)
	}
	return StringInput("parameter["+strconv.Itoa(numParameter)+"].referencedFrom["+strconv.Itoa(numReferencedFrom)+"].sourceId", resource.Parameter[numParameter].ReferencedFrom[numReferencedFrom].SourceId, htmlAttrs)
}
func (resource *OperationDefinition) T_OverloadParameterName(numOverload int, numParameterName int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOverload >= len(resource.Overload) || numParameterName >= len(resource.Overload[numOverload].ParameterName) {
		return StringInput("overload["+strconv.Itoa(numOverload)+"].parameterName["+strconv.Itoa(numParameterName)+"]", nil, htmlAttrs)
	}
	return StringInput("overload["+strconv.Itoa(numOverload)+"].parameterName["+strconv.Itoa(numParameterName)+"]", &resource.Overload[numOverload].ParameterName[numParameterName], htmlAttrs)
}
func (resource *OperationDefinition) T_OverloadComment(numOverload int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOverload >= len(resource.Overload) {
		return StringInput("overload["+strconv.Itoa(numOverload)+"].comment", nil, htmlAttrs)
	}
	return StringInput("overload["+strconv.Itoa(numOverload)+"].comment", resource.Overload[numOverload].Comment, htmlAttrs)
}
