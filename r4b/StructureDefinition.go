package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/StructureDefinition
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
	Date              *time.Time                       `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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

// http://hl7.org/fhir/r4b/StructureDefinition/StructureDefinition
type StructureDefinitionMapping struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Identity          string      `json:"identity"`
	Uri               *string     `json:"uri,omitempty"`
	Name              *string     `json:"name,omitempty"`
	Comment           *string     `json:"comment,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/StructureDefinition
type StructureDefinitionContext struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Expression        string      `json:"expression"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/StructureDefinition
type StructureDefinitionSnapshot struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Element           []ElementDefinition `json:"element"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/StructureDefinition
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
func (r StructureDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "StructureDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "StructureDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *StructureDefinition) T_Url(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Url", nil, htmlAttrs)
	}
	return StringInput("StructureDefinition.Url", &resource.Url, htmlAttrs)
}
func (resource *StructureDefinition) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Version", nil, htmlAttrs)
	}
	return StringInput("StructureDefinition.Version", resource.Version, htmlAttrs)
}
func (resource *StructureDefinition) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Name", nil, htmlAttrs)
	}
	return StringInput("StructureDefinition.Name", &resource.Name, htmlAttrs)
}
func (resource *StructureDefinition) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Title", nil, htmlAttrs)
	}
	return StringInput("StructureDefinition.Title", resource.Title, htmlAttrs)
}
func (resource *StructureDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("StructureDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureDefinition.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *StructureDefinition) T_Experimental(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("StructureDefinition.Experimental", nil, htmlAttrs)
	}
	return BoolInput("StructureDefinition.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *StructureDefinition) T_Date(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("StructureDefinition.Date", nil, htmlAttrs)
	}
	return DateTimeInput("StructureDefinition.Date", resource.Date, htmlAttrs)
}
func (resource *StructureDefinition) T_Publisher(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Publisher", nil, htmlAttrs)
	}
	return StringInput("StructureDefinition.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *StructureDefinition) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("StructureDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *StructureDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("StructureDefinition.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("StructureDefinition.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *StructureDefinition) T_Purpose(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Purpose", nil, htmlAttrs)
	}
	return StringInput("StructureDefinition.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *StructureDefinition) T_Copyright(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Copyright", nil, htmlAttrs)
	}
	return StringInput("StructureDefinition.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *StructureDefinition) T_Keyword(numKeyword int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numKeyword >= len(resource.Keyword) {
		return CodingSelect("StructureDefinition.Keyword."+strconv.Itoa(numKeyword)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("StructureDefinition.Keyword."+strconv.Itoa(numKeyword)+".", &resource.Keyword[numKeyword], optionsValueSet, htmlAttrs)
}
func (resource *StructureDefinition) T_FhirVersion(htmlAttrs string) templ.Component {
	optionsValueSet := VSFHIR_version

	if resource == nil {
		return CodeSelect("StructureDefinition.FhirVersion", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureDefinition.FhirVersion", resource.FhirVersion, optionsValueSet, htmlAttrs)
}
func (resource *StructureDefinition) T_Kind(htmlAttrs string) templ.Component {
	optionsValueSet := VSStructure_definition_kind

	if resource == nil {
		return CodeSelect("StructureDefinition.Kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureDefinition.Kind", &resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *StructureDefinition) T_Abstract(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("StructureDefinition.Abstract", nil, htmlAttrs)
	}
	return BoolInput("StructureDefinition.Abstract", &resource.Abstract, htmlAttrs)
}
func (resource *StructureDefinition) T_ContextInvariant(numContextInvariant int, htmlAttrs string) templ.Component {

	if resource == nil || numContextInvariant >= len(resource.ContextInvariant) {
		return StringInput("StructureDefinition.ContextInvariant."+strconv.Itoa(numContextInvariant)+".", nil, htmlAttrs)
	}
	return StringInput("StructureDefinition.ContextInvariant."+strconv.Itoa(numContextInvariant)+".", &resource.ContextInvariant[numContextInvariant], htmlAttrs)
}
func (resource *StructureDefinition) T_Type(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.Type", nil, htmlAttrs)
	}
	return StringInput("StructureDefinition.Type", &resource.Type, htmlAttrs)
}
func (resource *StructureDefinition) T_BaseDefinition(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("StructureDefinition.BaseDefinition", nil, htmlAttrs)
	}
	return StringInput("StructureDefinition.BaseDefinition", resource.BaseDefinition, htmlAttrs)
}
func (resource *StructureDefinition) T_Derivation(htmlAttrs string) templ.Component {
	optionsValueSet := VSType_derivation_rule

	if resource == nil {
		return CodeSelect("StructureDefinition.Derivation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureDefinition.Derivation", resource.Derivation, optionsValueSet, htmlAttrs)
}
func (resource *StructureDefinition) T_MappingIdentity(numMapping int, htmlAttrs string) templ.Component {

	if resource == nil || numMapping >= len(resource.Mapping) {
		return StringInput("StructureDefinition.Mapping."+strconv.Itoa(numMapping)+"..Identity", nil, htmlAttrs)
	}
	return StringInput("StructureDefinition.Mapping."+strconv.Itoa(numMapping)+"..Identity", &resource.Mapping[numMapping].Identity, htmlAttrs)
}
func (resource *StructureDefinition) T_MappingUri(numMapping int, htmlAttrs string) templ.Component {

	if resource == nil || numMapping >= len(resource.Mapping) {
		return StringInput("StructureDefinition.Mapping."+strconv.Itoa(numMapping)+"..Uri", nil, htmlAttrs)
	}
	return StringInput("StructureDefinition.Mapping."+strconv.Itoa(numMapping)+"..Uri", resource.Mapping[numMapping].Uri, htmlAttrs)
}
func (resource *StructureDefinition) T_MappingName(numMapping int, htmlAttrs string) templ.Component {

	if resource == nil || numMapping >= len(resource.Mapping) {
		return StringInput("StructureDefinition.Mapping."+strconv.Itoa(numMapping)+"..Name", nil, htmlAttrs)
	}
	return StringInput("StructureDefinition.Mapping."+strconv.Itoa(numMapping)+"..Name", resource.Mapping[numMapping].Name, htmlAttrs)
}
func (resource *StructureDefinition) T_MappingComment(numMapping int, htmlAttrs string) templ.Component {

	if resource == nil || numMapping >= len(resource.Mapping) {
		return StringInput("StructureDefinition.Mapping."+strconv.Itoa(numMapping)+"..Comment", nil, htmlAttrs)
	}
	return StringInput("StructureDefinition.Mapping."+strconv.Itoa(numMapping)+"..Comment", resource.Mapping[numMapping].Comment, htmlAttrs)
}
func (resource *StructureDefinition) T_ContextType(numContext int, htmlAttrs string) templ.Component {
	optionsValueSet := VSExtension_context_type

	if resource == nil || numContext >= len(resource.Context) {
		return CodeSelect("StructureDefinition.Context."+strconv.Itoa(numContext)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("StructureDefinition.Context."+strconv.Itoa(numContext)+"..Type", &resource.Context[numContext].Type, optionsValueSet, htmlAttrs)
}
func (resource *StructureDefinition) T_ContextExpression(numContext int, htmlAttrs string) templ.Component {

	if resource == nil || numContext >= len(resource.Context) {
		return StringInput("StructureDefinition.Context."+strconv.Itoa(numContext)+"..Expression", nil, htmlAttrs)
	}
	return StringInput("StructureDefinition.Context."+strconv.Itoa(numContext)+"..Expression", &resource.Context[numContext].Expression, htmlAttrs)
}
