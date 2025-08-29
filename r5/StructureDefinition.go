package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/StructureDefinition
type StructureDefinition struct {
	Id                     *string                          `json:"id,omitempty"`
	Meta                   *Meta                            `json:"meta,omitempty"`
	ImplicitRules          *string                          `json:"implicitRules,omitempty"`
	Language               *string                          `json:"language,omitempty"`
	Text                   *Narrative                       `json:"text,omitempty"`
	Contained              []Resource                       `json:"contained,omitempty"`
	Extension              []Extension                      `json:"extension,omitempty"`
	ModifierExtension      []Extension                      `json:"modifierExtension,omitempty"`
	Url                    string                           `json:"url"`
	Identifier             []Identifier                     `json:"identifier,omitempty"`
	Version                *string                          `json:"version,omitempty"`
	VersionAlgorithmString *string                          `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                          `json:"versionAlgorithmCoding"`
	Name                   string                           `json:"name"`
	Title                  *string                          `json:"title,omitempty"`
	Status                 string                           `json:"status"`
	Experimental           *bool                            `json:"experimental,omitempty"`
	Date                   *string                          `json:"date,omitempty"`
	Publisher              *string                          `json:"publisher,omitempty"`
	Contact                []ContactDetail                  `json:"contact,omitempty"`
	Description            *string                          `json:"description,omitempty"`
	UseContext             []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                `json:"jurisdiction,omitempty"`
	Purpose                *string                          `json:"purpose,omitempty"`
	Copyright              *string                          `json:"copyright,omitempty"`
	CopyrightLabel         *string                          `json:"copyrightLabel,omitempty"`
	Keyword                []Coding                         `json:"keyword,omitempty"`
	FhirVersion            *string                          `json:"fhirVersion,omitempty"`
	Mapping                []StructureDefinitionMapping     `json:"mapping,omitempty"`
	Kind                   string                           `json:"kind"`
	Abstract               bool                             `json:"abstract"`
	Context                []StructureDefinitionContext     `json:"context,omitempty"`
	ContextInvariant       []string                         `json:"contextInvariant,omitempty"`
	Type                   string                           `json:"type"`
	BaseDefinition         *string                          `json:"baseDefinition,omitempty"`
	Derivation             *string                          `json:"derivation,omitempty"`
	Snapshot               *StructureDefinitionSnapshot     `json:"snapshot,omitempty"`
	Differential           *StructureDefinitionDifferential `json:"differential,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureDefinition
type StructureDefinitionMapping struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Identity          string      `json:"identity"`
	Uri               *string     `json:"uri,omitempty"`
	Name              *string     `json:"name,omitempty"`
	Comment           *string     `json:"comment,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureDefinition
type StructureDefinitionContext struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Expression        string      `json:"expression"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureDefinition
type StructureDefinitionSnapshot struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Element           []ElementDefinition `json:"element"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureDefinition
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

func (resource *StructureDefinition) StructureDefinitionLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *StructureDefinition) StructureDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *StructureDefinition) StructureDefinitionFhirVersion() templ.Component {
	optionsValueSet := VSFHIR_version
	currentVal := ""
	if resource != nil {
		currentVal = *resource.FhirVersion
	}
	return CodeSelect("fhirVersion", currentVal, optionsValueSet)
}
func (resource *StructureDefinition) StructureDefinitionKind() templ.Component {
	optionsValueSet := VSStructure_definition_kind
	currentVal := ""
	if resource != nil {
		currentVal = resource.Kind
	}
	return CodeSelect("kind", currentVal, optionsValueSet)
}
func (resource *StructureDefinition) StructureDefinitionDerivation() templ.Component {
	optionsValueSet := VSType_derivation_rule
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Derivation
	}
	return CodeSelect("derivation", currentVal, optionsValueSet)
}
func (resource *StructureDefinition) StructureDefinitionContextType(numContext int) templ.Component {
	optionsValueSet := VSExtension_context_type
	currentVal := ""
	if resource != nil && len(resource.Context) >= numContext {
		currentVal = resource.Context[numContext].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
