package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/OperationDefinition
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
	Date              *time.Time                     `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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

// http://hl7.org/fhir/r4/StructureDefinition/OperationDefinition
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

// http://hl7.org/fhir/r4/StructureDefinition/OperationDefinition
type OperationDefinitionParameterBinding struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Strength          string      `json:"strength"`
	ValueSet          string      `json:"valueSet"`
}

// http://hl7.org/fhir/r4/StructureDefinition/OperationDefinition
type OperationDefinitionParameterReferencedFrom struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Source            string      `json:"source"`
	SourceId          *string     `json:"sourceId,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/OperationDefinition
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
func (resource *OperationDefinition) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *OperationDefinition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *OperationDefinition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", &resource.Name, htmlAttrs)
}
func (resource *OperationDefinition) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *OperationDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_Kind(htmlAttrs string) templ.Component {
	optionsValueSet := VSOperation_kind

	if resource == nil {
		return CodeSelect("Kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Kind", &resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Experimental", nil, htmlAttrs)
	}
	return BoolInput("Experimental", resource.Experimental, htmlAttrs)
}
func (resource *OperationDefinition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *OperationDefinition) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *OperationDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *OperationDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Purpose", nil, htmlAttrs)
	}
	return StringInput("Purpose", resource.Purpose, htmlAttrs)
}
func (resource *OperationDefinition) T_AffectsState(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("AffectsState", nil, htmlAttrs)
	}
	return BoolInput("AffectsState", resource.AffectsState, htmlAttrs)
}
func (resource *OperationDefinition) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeSelect("Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_Comment(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Comment", nil, htmlAttrs)
	}
	return StringInput("Comment", resource.Comment, htmlAttrs)
}
func (resource *OperationDefinition) T_Base(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Base", nil, htmlAttrs)
	}
	return StringInput("Base", resource.Base, htmlAttrs)
}
func (resource *OperationDefinition) T_Resource(numResource int, htmlAttrs string) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numResource >= len(resource.Resource) {
		return CodeSelect("Resource["+strconv.Itoa(numResource)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Resource["+strconv.Itoa(numResource)+"]", &resource.Resource[numResource], optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_System(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("System", nil, htmlAttrs)
	}
	return BoolInput("System", &resource.System, htmlAttrs)
}
func (resource *OperationDefinition) T_Type(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Type", nil, htmlAttrs)
	}
	return BoolInput("Type", &resource.Type, htmlAttrs)
}
func (resource *OperationDefinition) T_Instance(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Instance", nil, htmlAttrs)
	}
	return BoolInput("Instance", &resource.Instance, htmlAttrs)
}
func (resource *OperationDefinition) T_InputProfile(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("InputProfile", nil, htmlAttrs)
	}
	return StringInput("InputProfile", resource.InputProfile, htmlAttrs)
}
func (resource *OperationDefinition) T_OutputProfile(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("OutputProfile", nil, htmlAttrs)
	}
	return StringInput("OutputProfile", resource.OutputProfile, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterName(numParameter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeSelect("Parameter["+strconv.Itoa(numParameter)+"]Name", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Parameter["+strconv.Itoa(numParameter)+"]Name", &resource.Parameter[numParameter].Name, optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterUse(numParameter int, htmlAttrs string) templ.Component {
	optionsValueSet := VSOperation_parameter_use

	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeSelect("Parameter["+strconv.Itoa(numParameter)+"]Use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Parameter["+strconv.Itoa(numParameter)+"]Use", &resource.Parameter[numParameter].Use, optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterMin(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return IntInput("Parameter["+strconv.Itoa(numParameter)+"]Min", nil, htmlAttrs)
	}
	return IntInput("Parameter["+strconv.Itoa(numParameter)+"]Min", &resource.Parameter[numParameter].Min, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterMax(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]Max", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]Max", &resource.Parameter[numParameter].Max, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterDocumentation(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]Documentation", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]Documentation", resource.Parameter[numParameter].Documentation, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterType(numParameter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeSelect("Parameter["+strconv.Itoa(numParameter)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Parameter["+strconv.Itoa(numParameter)+"]Type", resource.Parameter[numParameter].Type, optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterTargetProfile(numParameter int, numTargetProfile int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) || numTargetProfile >= len(resource.Parameter[numParameter].TargetProfile) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]TargetProfile["+strconv.Itoa(numTargetProfile)+"]", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]TargetProfile["+strconv.Itoa(numTargetProfile)+"]", &resource.Parameter[numParameter].TargetProfile[numTargetProfile], htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterSearchType(numParameter int, htmlAttrs string) templ.Component {
	optionsValueSet := VSSearch_param_type

	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeSelect("Parameter["+strconv.Itoa(numParameter)+"]SearchType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Parameter["+strconv.Itoa(numParameter)+"]SearchType", resource.Parameter[numParameter].SearchType, optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterBindingStrength(numParameter int, htmlAttrs string) templ.Component {
	optionsValueSet := VSBinding_strength

	if resource == nil || numParameter >= len(resource.Parameter) {
		return CodeSelect("Parameter["+strconv.Itoa(numParameter)+"]Binding.Strength", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Parameter["+strconv.Itoa(numParameter)+"]Binding.Strength", &resource.Parameter[numParameter].Binding.Strength, optionsValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterBindingValueSet(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]Binding.ValueSet", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]Binding.ValueSet", &resource.Parameter[numParameter].Binding.ValueSet, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterReferencedFromSource(numParameter int, numReferencedFrom int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) || numReferencedFrom >= len(resource.Parameter[numParameter].ReferencedFrom) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ReferencedFrom["+strconv.Itoa(numReferencedFrom)+"].Source", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ReferencedFrom["+strconv.Itoa(numReferencedFrom)+"].Source", &resource.Parameter[numParameter].ReferencedFrom[numReferencedFrom].Source, htmlAttrs)
}
func (resource *OperationDefinition) T_ParameterReferencedFromSourceId(numParameter int, numReferencedFrom int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Parameter) || numReferencedFrom >= len(resource.Parameter[numParameter].ReferencedFrom) {
		return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ReferencedFrom["+strconv.Itoa(numReferencedFrom)+"].SourceId", nil, htmlAttrs)
	}
	return StringInput("Parameter["+strconv.Itoa(numParameter)+"]ReferencedFrom["+strconv.Itoa(numReferencedFrom)+"].SourceId", resource.Parameter[numParameter].ReferencedFrom[numReferencedFrom].SourceId, htmlAttrs)
}
func (resource *OperationDefinition) T_OverloadParameterName(numOverload int, numParameterName int, htmlAttrs string) templ.Component {
	if resource == nil || numOverload >= len(resource.Overload) || numParameterName >= len(resource.Overload[numOverload].ParameterName) {
		return StringInput("Overload["+strconv.Itoa(numOverload)+"]ParameterName["+strconv.Itoa(numParameterName)+"]", nil, htmlAttrs)
	}
	return StringInput("Overload["+strconv.Itoa(numOverload)+"]ParameterName["+strconv.Itoa(numParameterName)+"]", &resource.Overload[numOverload].ParameterName[numParameterName], htmlAttrs)
}
func (resource *OperationDefinition) T_OverloadComment(numOverload int, htmlAttrs string) templ.Component {
	if resource == nil || numOverload >= len(resource.Overload) {
		return StringInput("Overload["+strconv.Itoa(numOverload)+"]Comment", nil, htmlAttrs)
	}
	return StringInput("Overload["+strconv.Itoa(numOverload)+"]Comment", resource.Overload[numOverload].Comment, htmlAttrs)
}
