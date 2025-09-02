package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *StructureDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *StructureDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *StructureDefinition) T_Jurisdiction(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *StructureDefinition) T_Keyword(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("keyword", nil, optionsValueSet)
	}
	return CodingSelect("keyword", &resource.Keyword[0], optionsValueSet)
}
func (resource *StructureDefinition) T_FhirVersion() templ.Component {
	optionsValueSet := VSFHIR_version

	if resource == nil {
		return CodeSelect("fhirVersion", nil, optionsValueSet)
	}
	return CodeSelect("fhirVersion", resource.FhirVersion, optionsValueSet)
}
func (resource *StructureDefinition) T_Kind() templ.Component {
	optionsValueSet := VSStructure_definition_kind

	if resource == nil {
		return CodeSelect("kind", nil, optionsValueSet)
	}
	return CodeSelect("kind", &resource.Kind, optionsValueSet)
}
func (resource *StructureDefinition) T_Derivation() templ.Component {
	optionsValueSet := VSType_derivation_rule

	if resource == nil {
		return CodeSelect("derivation", nil, optionsValueSet)
	}
	return CodeSelect("derivation", resource.Derivation, optionsValueSet)
}
func (resource *StructureDefinition) T_ContextType(numContext int) templ.Component {
	optionsValueSet := VSExtension_context_type

	if resource == nil && len(resource.Context) >= numContext {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", &resource.Context[numContext].Type, optionsValueSet)
}
