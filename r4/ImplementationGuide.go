package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	ExampleBoolean    *bool       `json:"exampleBoolean,omitempty"`
	ExampleCanonical  *string     `json:"exampleCanonical,omitempty"`
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
	ExampleBoolean    *bool       `json:"exampleBoolean,omitempty"`
	ExampleCanonical  *string     `json:"exampleCanonical,omitempty"`
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
func (r ImplementationGuide) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ImplementationGuide/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "ImplementationGuide"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ImplementationGuide) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", &resource.Url, htmlAttrs)
}
func (resource *ImplementationGuide) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *ImplementationGuide) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", &resource.Name, htmlAttrs)
}
func (resource *ImplementationGuide) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *ImplementationGuide) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *ImplementationGuide) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *ImplementationGuide) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *ImplementationGuide) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *ImplementationGuide) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *ImplementationGuide) T_PackageId(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("packageId", nil, htmlAttrs)
	}
	return StringInput("packageId", &resource.PackageId, htmlAttrs)
}
func (resource *ImplementationGuide) T_License(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSpdx_license

	if resource == nil {
		return CodeSelect("license", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("license", resource.License, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_FhirVersion(numFhirVersion int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFHIR_version

	if resource == nil || numFhirVersion >= len(resource.FhirVersion) {
		return CodeSelect("fhirVersion["+strconv.Itoa(numFhirVersion)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("fhirVersion["+strconv.Itoa(numFhirVersion)+"]", &resource.FhirVersion[numFhirVersion], optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_DependsOnUri(numDependsOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDependsOn >= len(resource.DependsOn) {
		return StringInput("dependsOn["+strconv.Itoa(numDependsOn)+"].uri", nil, htmlAttrs)
	}
	return StringInput("dependsOn["+strconv.Itoa(numDependsOn)+"].uri", &resource.DependsOn[numDependsOn].Uri, htmlAttrs)
}
func (resource *ImplementationGuide) T_DependsOnPackageId(numDependsOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDependsOn >= len(resource.DependsOn) {
		return StringInput("dependsOn["+strconv.Itoa(numDependsOn)+"].packageId", nil, htmlAttrs)
	}
	return StringInput("dependsOn["+strconv.Itoa(numDependsOn)+"].packageId", resource.DependsOn[numDependsOn].PackageId, htmlAttrs)
}
func (resource *ImplementationGuide) T_DependsOnVersion(numDependsOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDependsOn >= len(resource.DependsOn) {
		return StringInput("dependsOn["+strconv.Itoa(numDependsOn)+"].version", nil, htmlAttrs)
	}
	return StringInput("dependsOn["+strconv.Itoa(numDependsOn)+"].version", resource.DependsOn[numDependsOn].Version, htmlAttrs)
}
func (resource *ImplementationGuide) T_GlobalType(numGlobal int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numGlobal >= len(resource.Global) {
		return CodeSelect("global["+strconv.Itoa(numGlobal)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("global["+strconv.Itoa(numGlobal)+"].type", &resource.Global[numGlobal].Type, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_GlobalProfile(numGlobal int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGlobal >= len(resource.Global) {
		return StringInput("global["+strconv.Itoa(numGlobal)+"].profile", nil, htmlAttrs)
	}
	return StringInput("global["+strconv.Itoa(numGlobal)+"].profile", &resource.Global[numGlobal].Profile, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionGroupingName(numGrouping int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGrouping >= len(resource.Definition.Grouping) {
		return StringInput("definition.grouping["+strconv.Itoa(numGrouping)+"].name", nil, htmlAttrs)
	}
	return StringInput("definition.grouping["+strconv.Itoa(numGrouping)+"].name", &resource.Definition.Grouping[numGrouping].Name, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionGroupingDescription(numGrouping int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numGrouping >= len(resource.Definition.Grouping) {
		return StringInput("definition.grouping["+strconv.Itoa(numGrouping)+"].description", nil, htmlAttrs)
	}
	return StringInput("definition.grouping["+strconv.Itoa(numGrouping)+"].description", resource.Definition.Grouping[numGrouping].Description, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceFhirVersion(numResource int, numFhirVersion int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFHIR_version

	if resource == nil || numResource >= len(resource.Definition.Resource) || numFhirVersion >= len(resource.Definition.Resource[numResource].FhirVersion) {
		return CodeSelect("definition.resource["+strconv.Itoa(numResource)+"].fhirVersion["+strconv.Itoa(numFhirVersion)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("definition.resource["+strconv.Itoa(numResource)+"].fhirVersion["+strconv.Itoa(numFhirVersion)+"]", &resource.Definition.Resource[numResource].FhirVersion[numFhirVersion], optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceName(numResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResource >= len(resource.Definition.Resource) {
		return StringInput("definition.resource["+strconv.Itoa(numResource)+"].name", nil, htmlAttrs)
	}
	return StringInput("definition.resource["+strconv.Itoa(numResource)+"].name", resource.Definition.Resource[numResource].Name, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceDescription(numResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResource >= len(resource.Definition.Resource) {
		return StringInput("definition.resource["+strconv.Itoa(numResource)+"].description", nil, htmlAttrs)
	}
	return StringInput("definition.resource["+strconv.Itoa(numResource)+"].description", resource.Definition.Resource[numResource].Description, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceExampleBoolean(numResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResource >= len(resource.Definition.Resource) {
		return BoolInput("definition.resource["+strconv.Itoa(numResource)+"].exampleBoolean", nil, htmlAttrs)
	}
	return BoolInput("definition.resource["+strconv.Itoa(numResource)+"].exampleBoolean", resource.Definition.Resource[numResource].ExampleBoolean, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceExampleCanonical(numResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResource >= len(resource.Definition.Resource) {
		return StringInput("definition.resource["+strconv.Itoa(numResource)+"].exampleCanonical", nil, htmlAttrs)
	}
	return StringInput("definition.resource["+strconv.Itoa(numResource)+"].exampleCanonical", resource.Definition.Resource[numResource].ExampleCanonical, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceGroupingId(numResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResource >= len(resource.Definition.Resource) {
		return StringInput("definition.resource["+strconv.Itoa(numResource)+"].groupingId", nil, htmlAttrs)
	}
	return StringInput("definition.resource["+strconv.Itoa(numResource)+"].groupingId", resource.Definition.Resource[numResource].GroupingId, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionPageNameUrl(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("definition.page.nameUrl", nil, htmlAttrs)
	}
	return StringInput("definition.page.nameUrl", &resource.Definition.Page.NameUrl, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionPageTitle(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("definition.page.title", nil, htmlAttrs)
	}
	return StringInput("definition.page.title", &resource.Definition.Page.Title, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionPageGeneration(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSGuide_page_generation

	if resource == nil {
		return CodeSelect("definition.page.generation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("definition.page.generation", &resource.Definition.Page.Generation, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionParameterCode(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSGuide_parameter_code

	if resource == nil || numParameter >= len(resource.Definition.Parameter) {
		return CodeSelect("definition.parameter["+strconv.Itoa(numParameter)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("definition.parameter["+strconv.Itoa(numParameter)+"].code", &resource.Definition.Parameter[numParameter].Code, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionParameterValue(numParameter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParameter >= len(resource.Definition.Parameter) {
		return StringInput("definition.parameter["+strconv.Itoa(numParameter)+"].value", nil, htmlAttrs)
	}
	return StringInput("definition.parameter["+strconv.Itoa(numParameter)+"].value", &resource.Definition.Parameter[numParameter].Value, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionTemplateCode(numTemplate int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTemplate >= len(resource.Definition.Template) {
		return CodeSelect("definition.template["+strconv.Itoa(numTemplate)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("definition.template["+strconv.Itoa(numTemplate)+"].code", &resource.Definition.Template[numTemplate].Code, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionTemplateSource(numTemplate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTemplate >= len(resource.Definition.Template) {
		return StringInput("definition.template["+strconv.Itoa(numTemplate)+"].source", nil, htmlAttrs)
	}
	return StringInput("definition.template["+strconv.Itoa(numTemplate)+"].source", &resource.Definition.Template[numTemplate].Source, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionTemplateScope(numTemplate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTemplate >= len(resource.Definition.Template) {
		return StringInput("definition.template["+strconv.Itoa(numTemplate)+"].scope", nil, htmlAttrs)
	}
	return StringInput("definition.template["+strconv.Itoa(numTemplate)+"].scope", resource.Definition.Template[numTemplate].Scope, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestRendering(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("manifest.rendering", nil, htmlAttrs)
	}
	return StringInput("manifest.rendering", resource.Manifest.Rendering, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestImage(numImage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numImage >= len(resource.Manifest.Image) {
		return StringInput("manifest.image["+strconv.Itoa(numImage)+"]", nil, htmlAttrs)
	}
	return StringInput("manifest.image["+strconv.Itoa(numImage)+"]", &resource.Manifest.Image[numImage], htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestOther(numOther int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOther >= len(resource.Manifest.Other) {
		return StringInput("manifest.other["+strconv.Itoa(numOther)+"]", nil, htmlAttrs)
	}
	return StringInput("manifest.other["+strconv.Itoa(numOther)+"]", &resource.Manifest.Other[numOther], htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestResourceExampleBoolean(numResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResource >= len(resource.Manifest.Resource) {
		return BoolInput("manifest.resource["+strconv.Itoa(numResource)+"].exampleBoolean", nil, htmlAttrs)
	}
	return BoolInput("manifest.resource["+strconv.Itoa(numResource)+"].exampleBoolean", resource.Manifest.Resource[numResource].ExampleBoolean, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestResourceExampleCanonical(numResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResource >= len(resource.Manifest.Resource) {
		return StringInput("manifest.resource["+strconv.Itoa(numResource)+"].exampleCanonical", nil, htmlAttrs)
	}
	return StringInput("manifest.resource["+strconv.Itoa(numResource)+"].exampleCanonical", resource.Manifest.Resource[numResource].ExampleCanonical, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestResourceRelativePath(numResource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numResource >= len(resource.Manifest.Resource) {
		return StringInput("manifest.resource["+strconv.Itoa(numResource)+"].relativePath", nil, htmlAttrs)
	}
	return StringInput("manifest.resource["+strconv.Itoa(numResource)+"].relativePath", resource.Manifest.Resource[numResource].RelativePath, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestPageName(numPage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPage >= len(resource.Manifest.Page) {
		return StringInput("manifest.page["+strconv.Itoa(numPage)+"].name", nil, htmlAttrs)
	}
	return StringInput("manifest.page["+strconv.Itoa(numPage)+"].name", &resource.Manifest.Page[numPage].Name, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestPageTitle(numPage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPage >= len(resource.Manifest.Page) {
		return StringInput("manifest.page["+strconv.Itoa(numPage)+"].title", nil, htmlAttrs)
	}
	return StringInput("manifest.page["+strconv.Itoa(numPage)+"].title", resource.Manifest.Page[numPage].Title, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestPageAnchor(numPage int, numAnchor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPage >= len(resource.Manifest.Page) || numAnchor >= len(resource.Manifest.Page[numPage].Anchor) {
		return StringInput("manifest.page["+strconv.Itoa(numPage)+"].anchor["+strconv.Itoa(numAnchor)+"]", nil, htmlAttrs)
	}
	return StringInput("manifest.page["+strconv.Itoa(numPage)+"].anchor["+strconv.Itoa(numAnchor)+"]", &resource.Manifest.Page[numPage].Anchor[numAnchor], htmlAttrs)
}
