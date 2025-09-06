package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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

func (resource *OperationDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("OperationDefinition.Id", nil)
	}
	return StringInput("OperationDefinition.Id", resource.Id)
}
func (resource *OperationDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("OperationDefinition.ImplicitRules", nil)
	}
	return StringInput("OperationDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *OperationDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("OperationDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("OperationDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *OperationDefinition) T_Url() templ.Component {

	if resource == nil {
		return StringInput("OperationDefinition.Url", nil)
	}
	return StringInput("OperationDefinition.Url", resource.Url)
}
func (resource *OperationDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("OperationDefinition.Version", nil)
	}
	return StringInput("OperationDefinition.Version", resource.Version)
}
func (resource *OperationDefinition) T_Name() templ.Component {

	if resource == nil {
		return StringInput("OperationDefinition.Name", nil)
	}
	return StringInput("OperationDefinition.Name", &resource.Name)
}
func (resource *OperationDefinition) T_Title() templ.Component {

	if resource == nil {
		return StringInput("OperationDefinition.Title", nil)
	}
	return StringInput("OperationDefinition.Title", resource.Title)
}
func (resource *OperationDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("OperationDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("OperationDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *OperationDefinition) T_Kind() templ.Component {
	optionsValueSet := VSOperation_kind

	if resource == nil {
		return CodeSelect("OperationDefinition.Kind", nil, optionsValueSet)
	}
	return CodeSelect("OperationDefinition.Kind", &resource.Kind, optionsValueSet)
}
func (resource *OperationDefinition) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("OperationDefinition.Experimental", nil)
	}
	return BoolInput("OperationDefinition.Experimental", resource.Experimental)
}
func (resource *OperationDefinition) T_Date() templ.Component {

	if resource == nil {
		return StringInput("OperationDefinition.Date", nil)
	}
	return StringInput("OperationDefinition.Date", resource.Date)
}
func (resource *OperationDefinition) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("OperationDefinition.Publisher", nil)
	}
	return StringInput("OperationDefinition.Publisher", resource.Publisher)
}
func (resource *OperationDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("OperationDefinition.Description", nil)
	}
	return StringInput("OperationDefinition.Description", resource.Description)
}
func (resource *OperationDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("OperationDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("OperationDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *OperationDefinition) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("OperationDefinition.Purpose", nil)
	}
	return StringInput("OperationDefinition.Purpose", resource.Purpose)
}
func (resource *OperationDefinition) T_AffectsState() templ.Component {

	if resource == nil {
		return BoolInput("OperationDefinition.AffectsState", nil)
	}
	return BoolInput("OperationDefinition.AffectsState", resource.AffectsState)
}
func (resource *OperationDefinition) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("OperationDefinition.Code", nil, optionsValueSet)
	}
	return CodeSelect("OperationDefinition.Code", &resource.Code, optionsValueSet)
}
func (resource *OperationDefinition) T_Comment() templ.Component {

	if resource == nil {
		return StringInput("OperationDefinition.Comment", nil)
	}
	return StringInput("OperationDefinition.Comment", resource.Comment)
}
func (resource *OperationDefinition) T_Base() templ.Component {

	if resource == nil {
		return StringInput("OperationDefinition.Base", nil)
	}
	return StringInput("OperationDefinition.Base", resource.Base)
}
func (resource *OperationDefinition) T_Resource(numResource int) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || len(resource.Resource) >= numResource {
		return CodeSelect("OperationDefinition.Resource["+strconv.Itoa(numResource)+"]", nil, optionsValueSet)
	}
	return CodeSelect("OperationDefinition.Resource["+strconv.Itoa(numResource)+"]", &resource.Resource[numResource], optionsValueSet)
}
func (resource *OperationDefinition) T_System() templ.Component {

	if resource == nil {
		return BoolInput("OperationDefinition.System", nil)
	}
	return BoolInput("OperationDefinition.System", &resource.System)
}
func (resource *OperationDefinition) T_Type() templ.Component {

	if resource == nil {
		return BoolInput("OperationDefinition.Type", nil)
	}
	return BoolInput("OperationDefinition.Type", &resource.Type)
}
func (resource *OperationDefinition) T_Instance() templ.Component {

	if resource == nil {
		return BoolInput("OperationDefinition.Instance", nil)
	}
	return BoolInput("OperationDefinition.Instance", &resource.Instance)
}
func (resource *OperationDefinition) T_InputProfile() templ.Component {

	if resource == nil {
		return StringInput("OperationDefinition.InputProfile", nil)
	}
	return StringInput("OperationDefinition.InputProfile", resource.InputProfile)
}
func (resource *OperationDefinition) T_OutputProfile() templ.Component {

	if resource == nil {
		return StringInput("OperationDefinition.OutputProfile", nil)
	}
	return StringInput("OperationDefinition.OutputProfile", resource.OutputProfile)
}
func (resource *OperationDefinition) T_ParameterId(numParameter int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Id", nil)
	}
	return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Id", resource.Parameter[numParameter].Id)
}
func (resource *OperationDefinition) T_ParameterName(numParameter int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return CodeSelect("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Name", nil, optionsValueSet)
	}
	return CodeSelect("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Name", &resource.Parameter[numParameter].Name, optionsValueSet)
}
func (resource *OperationDefinition) T_ParameterUse(numParameter int) templ.Component {
	optionsValueSet := VSOperation_parameter_use

	if resource == nil || len(resource.Parameter) >= numParameter {
		return CodeSelect("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Use", nil, optionsValueSet)
	}
	return CodeSelect("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Use", &resource.Parameter[numParameter].Use, optionsValueSet)
}
func (resource *OperationDefinition) T_ParameterMin(numParameter int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return IntInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Min", nil)
	}
	return IntInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Min", &resource.Parameter[numParameter].Min)
}
func (resource *OperationDefinition) T_ParameterMax(numParameter int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Max", nil)
	}
	return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Max", &resource.Parameter[numParameter].Max)
}
func (resource *OperationDefinition) T_ParameterDocumentation(numParameter int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Documentation", nil)
	}
	return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Documentation", resource.Parameter[numParameter].Documentation)
}
func (resource *OperationDefinition) T_ParameterType(numParameter int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return CodeSelect("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Type", resource.Parameter[numParameter].Type, optionsValueSet)
}
func (resource *OperationDefinition) T_ParameterTargetProfile(numParameter int, numTargetProfile int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter || len(resource.Parameter[numParameter].TargetProfile) >= numTargetProfile {
		return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].TargetProfile["+strconv.Itoa(numTargetProfile)+"]", nil)
	}
	return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].TargetProfile["+strconv.Itoa(numTargetProfile)+"]", &resource.Parameter[numParameter].TargetProfile[numTargetProfile])
}
func (resource *OperationDefinition) T_ParameterSearchType(numParameter int) templ.Component {
	optionsValueSet := VSSearch_param_type

	if resource == nil || len(resource.Parameter) >= numParameter {
		return CodeSelect("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].SearchType", nil, optionsValueSet)
	}
	return CodeSelect("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].SearchType", resource.Parameter[numParameter].SearchType, optionsValueSet)
}
func (resource *OperationDefinition) T_ParameterBindingId(numParameter int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Binding.Id", nil)
	}
	return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Binding.Id", resource.Parameter[numParameter].Binding.Id)
}
func (resource *OperationDefinition) T_ParameterBindingStrength(numParameter int) templ.Component {
	optionsValueSet := VSBinding_strength

	if resource == nil || len(resource.Parameter) >= numParameter {
		return CodeSelect("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Binding.Strength", nil, optionsValueSet)
	}
	return CodeSelect("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Binding.Strength", &resource.Parameter[numParameter].Binding.Strength, optionsValueSet)
}
func (resource *OperationDefinition) T_ParameterBindingValueSet(numParameter int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter {
		return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Binding.ValueSet", nil)
	}
	return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].Binding.ValueSet", &resource.Parameter[numParameter].Binding.ValueSet)
}
func (resource *OperationDefinition) T_ParameterReferencedFromId(numParameter int, numReferencedFrom int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter || len(resource.Parameter[numParameter].ReferencedFrom) >= numReferencedFrom {
		return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].ReferencedFrom["+strconv.Itoa(numReferencedFrom)+"].Id", nil)
	}
	return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].ReferencedFrom["+strconv.Itoa(numReferencedFrom)+"].Id", resource.Parameter[numParameter].ReferencedFrom[numReferencedFrom].Id)
}
func (resource *OperationDefinition) T_ParameterReferencedFromSource(numParameter int, numReferencedFrom int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter || len(resource.Parameter[numParameter].ReferencedFrom) >= numReferencedFrom {
		return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].ReferencedFrom["+strconv.Itoa(numReferencedFrom)+"].Source", nil)
	}
	return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].ReferencedFrom["+strconv.Itoa(numReferencedFrom)+"].Source", &resource.Parameter[numParameter].ReferencedFrom[numReferencedFrom].Source)
}
func (resource *OperationDefinition) T_ParameterReferencedFromSourceId(numParameter int, numReferencedFrom int) templ.Component {

	if resource == nil || len(resource.Parameter) >= numParameter || len(resource.Parameter[numParameter].ReferencedFrom) >= numReferencedFrom {
		return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].ReferencedFrom["+strconv.Itoa(numReferencedFrom)+"].SourceId", nil)
	}
	return StringInput("OperationDefinition.Parameter["+strconv.Itoa(numParameter)+"].ReferencedFrom["+strconv.Itoa(numReferencedFrom)+"].SourceId", resource.Parameter[numParameter].ReferencedFrom[numReferencedFrom].SourceId)
}
func (resource *OperationDefinition) T_OverloadId(numOverload int) templ.Component {

	if resource == nil || len(resource.Overload) >= numOverload {
		return StringInput("OperationDefinition.Overload["+strconv.Itoa(numOverload)+"].Id", nil)
	}
	return StringInput("OperationDefinition.Overload["+strconv.Itoa(numOverload)+"].Id", resource.Overload[numOverload].Id)
}
func (resource *OperationDefinition) T_OverloadParameterName(numOverload int, numParameterName int) templ.Component {

	if resource == nil || len(resource.Overload) >= numOverload || len(resource.Overload[numOverload].ParameterName) >= numParameterName {
		return StringInput("OperationDefinition.Overload["+strconv.Itoa(numOverload)+"].ParameterName["+strconv.Itoa(numParameterName)+"]", nil)
	}
	return StringInput("OperationDefinition.Overload["+strconv.Itoa(numOverload)+"].ParameterName["+strconv.Itoa(numParameterName)+"]", &resource.Overload[numOverload].ParameterName[numParameterName])
}
func (resource *OperationDefinition) T_OverloadComment(numOverload int) templ.Component {

	if resource == nil || len(resource.Overload) >= numOverload {
		return StringInput("OperationDefinition.Overload["+strconv.Itoa(numOverload)+"].Comment", nil)
	}
	return StringInput("OperationDefinition.Overload["+strconv.Itoa(numOverload)+"].Comment", resource.Overload[numOverload].Comment)
}
