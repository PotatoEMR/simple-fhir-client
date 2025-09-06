package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/StructureDefinition
type StructureDefinition struct {
	Id                *string                          `json:"id,omitempty"`
	Meta              *Meta                            `json:"meta,omitempty"`
	ImplicitRules     *string                          `json:"implicitRules,omitempty"`
	Language          *string                          `json:"language,omitempty"`
	Text              *Narrative                       `json:"text,omitempty"`
	Contained         []Resource                       `json:"contained,omitempty"`
	Extension         []Extension                      `json:"extension,omitempty"`
	ModifierExtension []Extension                      `json:"modifierExtension,omitempty"`
	Url               string                           `json:"url"`
	Identifier        []Identifier                     `json:"identifier,omitempty"`
	Version           *string                          `json:"version,omitempty"`
	Name              string                           `json:"name"`
	Title             *string                          `json:"title,omitempty"`
	Status            string                           `json:"status"`
	Experimental      *bool                            `json:"experimental,omitempty"`
	Date              *string                          `json:"date,omitempty"`
	Publisher         *string                          `json:"publisher,omitempty"`
	Contact           []ContactDetail                  `json:"contact,omitempty"`
	Description       *string                          `json:"description,omitempty"`
	UseContext        []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                `json:"jurisdiction,omitempty"`
	Purpose           *string                          `json:"purpose,omitempty"`
	Copyright         *string                          `json:"copyright,omitempty"`
	Keyword           []Coding                         `json:"keyword,omitempty"`
	FhirVersion       *string                          `json:"fhirVersion,omitempty"`
	Mapping           []StructureDefinitionMapping     `json:"mapping,omitempty"`
	Kind              string                           `json:"kind"`
	Abstract          bool                             `json:"abstract"`
	Context           []StructureDefinitionContext     `json:"context,omitempty"`
	ContextInvariant  []string                         `json:"contextInvariant,omitempty"`
	Type              string                           `json:"type"`
	BaseDefinition    *string                          `json:"baseDefinition,omitempty"`
	Derivation        *string                          `json:"derivation,omitempty"`
	Snapshot          *StructureDefinitionSnapshot     `json:"snapshot,omitempty"`
	Differential      *StructureDefinitionDifferential `json:"differential,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/StructureDefinition
type StructureDefinitionMapping struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Identity          string      `json:"identity"`
	Uri               *string     `json:"uri,omitempty"`
	Name              *string     `json:"name,omitempty"`
	Comment           *string     `json:"comment,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/StructureDefinition
type StructureDefinitionContext struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Expression        string      `json:"expression"`
}

// http://hl7.org/fhir/r4/StructureDefinition/StructureDefinition
type StructureDefinitionSnapshot struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Element           []ElementDefinition `json:"element"`
}

// http://hl7.org/fhir/r4/StructureDefinition/StructureDefinition
type StructureDefinitionDifferential struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Element           []ElementDefinition `json:"element"`
}

type OtherStructureDefinition StructureDefinition

// on convert struct to json, automatically add resourceType=StructureDefinition
func (r StructureDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherStructureDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherStructureDefinition: OtherStructureDefinition(r),
		ResourceType:             "StructureDefinition",
	})
}

func (resource *StructureDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Id", nil)
	}
	return StringInput("StructureDefinition.Id", resource.Id)
}
func (resource *StructureDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.ImplicitRules", nil)
	}
	return StringInput("StructureDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *StructureDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("StructureDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("StructureDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *StructureDefinition) T_Url() templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Url", nil)
	}
	return StringInput("StructureDefinition.Url", &resource.Url)
}
func (resource *StructureDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Version", nil)
	}
	return StringInput("StructureDefinition.Version", resource.Version)
}
func (resource *StructureDefinition) T_Name() templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Name", nil)
	}
	return StringInput("StructureDefinition.Name", &resource.Name)
}
func (resource *StructureDefinition) T_Title() templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Title", nil)
	}
	return StringInput("StructureDefinition.Title", resource.Title)
}
func (resource *StructureDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("StructureDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("StructureDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *StructureDefinition) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("StructureDefinition.Experimental", nil)
	}
	return BoolInput("StructureDefinition.Experimental", resource.Experimental)
}
func (resource *StructureDefinition) T_Date() templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Date", nil)
	}
	return StringInput("StructureDefinition.Date", resource.Date)
}
func (resource *StructureDefinition) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Publisher", nil)
	}
	return StringInput("StructureDefinition.Publisher", resource.Publisher)
}
func (resource *StructureDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Description", nil)
	}
	return StringInput("StructureDefinition.Description", resource.Description)
}
func (resource *StructureDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("StructureDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("StructureDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *StructureDefinition) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Purpose", nil)
	}
	return StringInput("StructureDefinition.Purpose", resource.Purpose)
}
func (resource *StructureDefinition) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Copyright", nil)
	}
	return StringInput("StructureDefinition.Copyright", resource.Copyright)
}
func (resource *StructureDefinition) T_Keyword(numKeyword int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Keyword) >= numKeyword {
		return CodingSelect("StructureDefinition.Keyword["+strconv.Itoa(numKeyword)+"]", nil, optionsValueSet)
	}
	return CodingSelect("StructureDefinition.Keyword["+strconv.Itoa(numKeyword)+"]", &resource.Keyword[numKeyword], optionsValueSet)
}
func (resource *StructureDefinition) T_FhirVersion() templ.Component {
	optionsValueSet := VSFHIR_version

	if resource == nil {
		return CodeSelect("StructureDefinition.FhirVersion", nil, optionsValueSet)
	}
	return CodeSelect("StructureDefinition.FhirVersion", resource.FhirVersion, optionsValueSet)
}
func (resource *StructureDefinition) T_Kind() templ.Component {
	optionsValueSet := VSStructure_definition_kind

	if resource == nil {
		return CodeSelect("StructureDefinition.Kind", nil, optionsValueSet)
	}
	return CodeSelect("StructureDefinition.Kind", &resource.Kind, optionsValueSet)
}
func (resource *StructureDefinition) T_Abstract() templ.Component {

	if resource == nil {
		return BoolInput("StructureDefinition.Abstract", nil)
	}
	return BoolInput("StructureDefinition.Abstract", &resource.Abstract)
}
func (resource *StructureDefinition) T_ContextInvariant(numContextInvariant int) templ.Component {

	if resource == nil || len(resource.ContextInvariant) >= numContextInvariant {
		return StringInput("StructureDefinition.ContextInvariant["+strconv.Itoa(numContextInvariant)+"]", nil)
	}
	return StringInput("StructureDefinition.ContextInvariant["+strconv.Itoa(numContextInvariant)+"]", &resource.ContextInvariant[numContextInvariant])
}
func (resource *StructureDefinition) T_Type() templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Type", nil)
	}
	return StringInput("StructureDefinition.Type", &resource.Type)
}
func (resource *StructureDefinition) T_BaseDefinition() templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.BaseDefinition", nil)
	}
	return StringInput("StructureDefinition.BaseDefinition", resource.BaseDefinition)
}
func (resource *StructureDefinition) T_Derivation() templ.Component {
	optionsValueSet := VSType_derivation_rule

	if resource == nil {
		return CodeSelect("StructureDefinition.Derivation", nil, optionsValueSet)
	}
	return CodeSelect("StructureDefinition.Derivation", resource.Derivation, optionsValueSet)
}
func (resource *StructureDefinition) T_MappingId(numMapping int) templ.Component {

	if resource == nil || len(resource.Mapping) >= numMapping {
		return StringInput("StructureDefinition.Mapping["+strconv.Itoa(numMapping)+"].Id", nil)
	}
	return StringInput("StructureDefinition.Mapping["+strconv.Itoa(numMapping)+"].Id", resource.Mapping[numMapping].Id)
}
func (resource *StructureDefinition) T_MappingIdentity(numMapping int) templ.Component {

	if resource == nil || len(resource.Mapping) >= numMapping {
		return StringInput("StructureDefinition.Mapping["+strconv.Itoa(numMapping)+"].Identity", nil)
	}
	return StringInput("StructureDefinition.Mapping["+strconv.Itoa(numMapping)+"].Identity", &resource.Mapping[numMapping].Identity)
}
func (resource *StructureDefinition) T_MappingUri(numMapping int) templ.Component {

	if resource == nil || len(resource.Mapping) >= numMapping {
		return StringInput("StructureDefinition.Mapping["+strconv.Itoa(numMapping)+"].Uri", nil)
	}
	return StringInput("StructureDefinition.Mapping["+strconv.Itoa(numMapping)+"].Uri", resource.Mapping[numMapping].Uri)
}
func (resource *StructureDefinition) T_MappingName(numMapping int) templ.Component {

	if resource == nil || len(resource.Mapping) >= numMapping {
		return StringInput("StructureDefinition.Mapping["+strconv.Itoa(numMapping)+"].Name", nil)
	}
	return StringInput("StructureDefinition.Mapping["+strconv.Itoa(numMapping)+"].Name", resource.Mapping[numMapping].Name)
}
func (resource *StructureDefinition) T_MappingComment(numMapping int) templ.Component {

	if resource == nil || len(resource.Mapping) >= numMapping {
		return StringInput("StructureDefinition.Mapping["+strconv.Itoa(numMapping)+"].Comment", nil)
	}
	return StringInput("StructureDefinition.Mapping["+strconv.Itoa(numMapping)+"].Comment", resource.Mapping[numMapping].Comment)
}
func (resource *StructureDefinition) T_ContextId(numContext int) templ.Component {

	if resource == nil || len(resource.Context) >= numContext {
		return StringInput("StructureDefinition.Context["+strconv.Itoa(numContext)+"].Id", nil)
	}
	return StringInput("StructureDefinition.Context["+strconv.Itoa(numContext)+"].Id", resource.Context[numContext].Id)
}
func (resource *StructureDefinition) T_ContextType(numContext int) templ.Component {
	optionsValueSet := VSExtension_context_type

	if resource == nil || len(resource.Context) >= numContext {
		return CodeSelect("StructureDefinition.Context["+strconv.Itoa(numContext)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("StructureDefinition.Context["+strconv.Itoa(numContext)+"].Type", &resource.Context[numContext].Type, optionsValueSet)
}
func (resource *StructureDefinition) T_ContextExpression(numContext int) templ.Component {

	if resource == nil || len(resource.Context) >= numContext {
		return StringInput("StructureDefinition.Context["+strconv.Itoa(numContext)+"].Expression", nil)
	}
	return StringInput("StructureDefinition.Context["+strconv.Itoa(numContext)+"].Expression", &resource.Context[numContext].Expression)
}
func (resource *StructureDefinition) T_SnapshotId() templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Snapshot.Id", nil)
	}
	return StringInput("StructureDefinition.Snapshot.Id", resource.Snapshot.Id)
}
func (resource *StructureDefinition) T_DifferentialId() templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Differential.Id", nil)
	}
	return StringInput("StructureDefinition.Differential.Id", resource.Differential.Id)
}
