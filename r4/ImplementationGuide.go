package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/ImplementationGuide
type ImplementationGuide struct {
	Id                *string                        `json:"id,omitempty"`
	Meta              *Meta                          `json:"meta,omitempty"`
	ImplicitRules     *string                        `json:"implicitRules,omitempty"`
	Language          *string                        `json:"language,omitempty"`
	Text              *Narrative                     `json:"text,omitempty"`
	Contained         []Resource                     `json:"contained,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Url               string                         `json:"url"`
	Version           *string                        `json:"version,omitempty"`
	Name              string                         `json:"name"`
	Title             *string                        `json:"title,omitempty"`
	Status            string                         `json:"status"`
	Experimental      *bool                          `json:"experimental,omitempty"`
	Date              *string                        `json:"date,omitempty"`
	Publisher         *string                        `json:"publisher,omitempty"`
	Contact           []ContactDetail                `json:"contact,omitempty"`
	Description       *string                        `json:"description,omitempty"`
	UseContext        []UsageContext                 `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept              `json:"jurisdiction,omitempty"`
	Copyright         *string                        `json:"copyright,omitempty"`
	PackageId         string                         `json:"packageId"`
	License           *string                        `json:"license,omitempty"`
	FhirVersion       []string                       `json:"fhirVersion"`
	DependsOn         []ImplementationGuideDependsOn `json:"dependsOn,omitempty"`
	Global            []ImplementationGuideGlobal    `json:"global,omitempty"`
	Definition        *ImplementationGuideDefinition `json:"definition,omitempty"`
	Manifest          *ImplementationGuideManifest   `json:"manifest,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ImplementationGuide
type ImplementationGuideDependsOn struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Uri               string      `json:"uri"`
	PackageId         *string     `json:"packageId,omitempty"`
	Version           *string     `json:"version,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ImplementationGuide
type ImplementationGuideGlobal struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Profile           string      `json:"profile"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinition struct {
	Id                *string                                  `json:"id,omitempty"`
	Extension         []Extension                              `json:"extension,omitempty"`
	ModifierExtension []Extension                              `json:"modifierExtension,omitempty"`
	Grouping          []ImplementationGuideDefinitionGrouping  `json:"grouping,omitempty"`
	Resource          []ImplementationGuideDefinitionResource  `json:"resource"`
	Page              *ImplementationGuideDefinitionPage       `json:"page,omitempty"`
	Parameter         []ImplementationGuideDefinitionParameter `json:"parameter,omitempty"`
	Template          []ImplementationGuideDefinitionTemplate  `json:"template,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinitionGrouping struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Description       *string     `json:"description,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinitionResource struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Reference         Reference   `json:"reference"`
	FhirVersion       []string    `json:"fhirVersion,omitempty"`
	Name              *string     `json:"name,omitempty"`
	Description       *string     `json:"description,omitempty"`
	ExampleBoolean    *bool       `json:"exampleBoolean"`
	ExampleCanonical  *string     `json:"exampleCanonical"`
	GroupingId        *string     `json:"groupingId,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinitionPage struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	NameUrl           string      `json:"nameUrl"`
	NameReference     Reference   `json:"nameReference"`
	Title             string      `json:"title"`
	Generation        string      `json:"generation"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinitionParameter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinitionTemplate struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Source            string      `json:"source"`
	Scope             *string     `json:"scope,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ImplementationGuide
type ImplementationGuideManifest struct {
	Id                *string                               `json:"id,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	Rendering         *string                               `json:"rendering,omitempty"`
	Resource          []ImplementationGuideManifestResource `json:"resource"`
	Page              []ImplementationGuideManifestPage     `json:"page,omitempty"`
	Image             []string                              `json:"image,omitempty"`
	Other             []string                              `json:"other,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ImplementationGuide
type ImplementationGuideManifestResource struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Reference         Reference   `json:"reference"`
	ExampleBoolean    *bool       `json:"exampleBoolean"`
	ExampleCanonical  *string     `json:"exampleCanonical"`
	RelativePath      *string     `json:"relativePath,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ImplementationGuide
type ImplementationGuideManifestPage struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Title             *string     `json:"title,omitempty"`
	Anchor            []string    `json:"anchor,omitempty"`
}

type OtherImplementationGuide ImplementationGuide

// on convert struct to json, automatically add resourceType=ImplementationGuide
func (r ImplementationGuide) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImplementationGuide
		ResourceType string `json:"resourceType"`
	}{
		OtherImplementationGuide: OtherImplementationGuide(r),
		ResourceType:             "ImplementationGuide",
	})
}

func (resource *ImplementationGuide) ImplementationGuideLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *ImplementationGuide) ImplementationGuideStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *ImplementationGuide) ImplementationGuideLicense() templ.Component {
	optionsValueSet := VSSpdx_license
	currentVal := ""
	if resource != nil {
		currentVal = *resource.License
	}
	return CodeSelect("license", currentVal, optionsValueSet)
}
func (resource *ImplementationGuide) ImplementationGuideFhirVersion() templ.Component {
	optionsValueSet := VSFHIR_version
	currentVal := ""
	if resource != nil {
		currentVal = resource.FhirVersion[0]
	}
	return CodeSelect("fhirVersion", currentVal, optionsValueSet)
}
func (resource *ImplementationGuide) ImplementationGuideGlobalType(numGlobal int) templ.Component {
	optionsValueSet := VSResource_types
	currentVal := ""
	if resource != nil && len(resource.Global) >= numGlobal {
		currentVal = resource.Global[numGlobal].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
func (resource *ImplementationGuide) ImplementationGuideDefinitionResourceFhirVersion(numResource int) templ.Component {
	optionsValueSet := VSFHIR_version
	currentVal := ""
	if resource != nil && len(resource.Definition.Resource) >= numResource {
		currentVal = resource.Definition.Resource[numResource].FhirVersion[0]
	}
	return CodeSelect("fhirVersion", currentVal, optionsValueSet)
}
func (resource *ImplementationGuide) ImplementationGuideDefinitionPageGeneration() templ.Component {
	optionsValueSet := VSGuide_page_generation
	currentVal := ""
	if resource != nil {
		currentVal = resource.Definition.Page.Generation
	}
	return CodeSelect("generation", currentVal, optionsValueSet)
}
func (resource *ImplementationGuide) ImplementationGuideDefinitionParameterCode(numParameter int) templ.Component {
	optionsValueSet := VSGuide_parameter_code
	currentVal := ""
	if resource != nil && len(resource.Definition.Parameter) >= numParameter {
		currentVal = resource.Definition.Parameter[numParameter].Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
func (resource *ImplementationGuide) ImplementationGuideDefinitionTemplateCode(numTemplate int, optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil && len(resource.Definition.Template) >= numTemplate {
		currentVal = resource.Definition.Template[numTemplate].Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
