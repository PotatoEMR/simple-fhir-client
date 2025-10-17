package r4

//generated with command go run ./bultaoreune
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
	Date              *FhirDateTime                    `json:"date,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
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
func (r StructureDefinition) ResourceType() string {
	return "StructureDefinition"
}

func (resource *StructureDefinition) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", &resource.Url, htmlAttrs)
}
func (resource *StructureDefinition) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *StructureDefinition) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", &resource.Name, htmlAttrs)
}
func (resource *StructureDefinition) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *StructureDefinition) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *StructureDefinition) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *StructureDefinition) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *StructureDefinition) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *StructureDefinition) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *StructureDefinition) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *StructureDefinition) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *StructureDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *StructureDefinition) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *StructureDefinition) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *StructureDefinition) T_Keyword(numKeyword int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numKeyword >= len(resource.Keyword) {
		return CodingSelect("keyword["+strconv.Itoa(numKeyword)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("keyword["+strconv.Itoa(numKeyword)+"]", &resource.Keyword[numKeyword], optionsValueSet, htmlAttrs)
}
func (resource *StructureDefinition) T_FhirVersion(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFHIR_version

	if resource == nil {
		return CodeSelect("fhirVersion", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("fhirVersion", resource.FhirVersion, optionsValueSet, htmlAttrs)
}
func (resource *StructureDefinition) T_Kind(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSStructure_definition_kind

	if resource == nil {
		return CodeSelect("kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("kind", &resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *StructureDefinition) T_Abstract(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("abstract", nil, htmlAttrs)
	}
	return BoolInput("abstract", &resource.Abstract, htmlAttrs)
}
func (resource *StructureDefinition) T_ContextInvariant(numContextInvariant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContextInvariant >= len(resource.ContextInvariant) {
		return StringInput("contextInvariant["+strconv.Itoa(numContextInvariant)+"]", nil, htmlAttrs)
	}
	return StringInput("contextInvariant["+strconv.Itoa(numContextInvariant)+"]", &resource.ContextInvariant[numContextInvariant], htmlAttrs)
}
func (resource *StructureDefinition) T_Type(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("type", nil, htmlAttrs)
	}
	return StringInput("type", &resource.Type, htmlAttrs)
}
func (resource *StructureDefinition) T_BaseDefinition(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("baseDefinition", nil, htmlAttrs)
	}
	return StringInput("baseDefinition", resource.BaseDefinition, htmlAttrs)
}
func (resource *StructureDefinition) T_Derivation(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSType_derivation_rule

	if resource == nil {
		return CodeSelect("derivation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("derivation", resource.Derivation, optionsValueSet, htmlAttrs)
}
func (resource *StructureDefinition) T_MappingIdentity(numMapping int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMapping >= len(resource.Mapping) {
		return StringInput("mapping["+strconv.Itoa(numMapping)+"].identity", nil, htmlAttrs)
	}
	return StringInput("mapping["+strconv.Itoa(numMapping)+"].identity", &resource.Mapping[numMapping].Identity, htmlAttrs)
}
func (resource *StructureDefinition) T_MappingUri(numMapping int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMapping >= len(resource.Mapping) {
		return StringInput("mapping["+strconv.Itoa(numMapping)+"].uri", nil, htmlAttrs)
	}
	return StringInput("mapping["+strconv.Itoa(numMapping)+"].uri", resource.Mapping[numMapping].Uri, htmlAttrs)
}
func (resource *StructureDefinition) T_MappingName(numMapping int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMapping >= len(resource.Mapping) {
		return StringInput("mapping["+strconv.Itoa(numMapping)+"].name", nil, htmlAttrs)
	}
	return StringInput("mapping["+strconv.Itoa(numMapping)+"].name", resource.Mapping[numMapping].Name, htmlAttrs)
}
func (resource *StructureDefinition) T_MappingComment(numMapping int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMapping >= len(resource.Mapping) {
		return StringInput("mapping["+strconv.Itoa(numMapping)+"].comment", nil, htmlAttrs)
	}
	return StringInput("mapping["+strconv.Itoa(numMapping)+"].comment", resource.Mapping[numMapping].Comment, htmlAttrs)
}
func (resource *StructureDefinition) T_ContextType(numContext int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSExtension_context_type

	if resource == nil || numContext >= len(resource.Context) {
		return CodeSelect("context["+strconv.Itoa(numContext)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("context["+strconv.Itoa(numContext)+"].type", &resource.Context[numContext].Type, optionsValueSet, htmlAttrs)
}
func (resource *StructureDefinition) T_ContextExpression(numContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContext >= len(resource.Context) {
		return StringInput("context["+strconv.Itoa(numContext)+"].expression", nil, htmlAttrs)
	}
	return StringInput("context["+strconv.Itoa(numContext)+"].expression", &resource.Context[numContext].Expression, htmlAttrs)
}
