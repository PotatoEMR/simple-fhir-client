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

// http://hl7.org/fhir/r4b/StructureDefinition/ImplementationGuide
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
	Date              *time.Time                     `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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

// http://hl7.org/fhir/r4b/StructureDefinition/ImplementationGuide
type ImplementationGuideDependsOn struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Uri               string      `json:"uri"`
	PackageId         *string     `json:"packageId,omitempty"`
	Version           *string     `json:"version,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ImplementationGuide
type ImplementationGuideGlobal struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Profile           string      `json:"profile"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ImplementationGuide
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

// http://hl7.org/fhir/r4b/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinitionGrouping struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Description       *string     `json:"description,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ImplementationGuide
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

// http://hl7.org/fhir/r4b/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinitionPage struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	NameUrl           string      `json:"nameUrl"`
	NameReference     Reference   `json:"nameReference"`
	Title             string      `json:"title"`
	Generation        string      `json:"generation"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinitionParameter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinitionTemplate struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Source            string      `json:"source"`
	Scope             *string     `json:"scope,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ImplementationGuide
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

// http://hl7.org/fhir/r4b/StructureDefinition/ImplementationGuide
type ImplementationGuideManifestResource struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Reference         Reference   `json:"reference"`
	ExampleBoolean    *bool       `json:"exampleBoolean,omitempty"`
	ExampleCanonical  *string     `json:"exampleCanonical,omitempty"`
	RelativePath      *string     `json:"relativePath,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ImplementationGuide
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
func (resource *ImplementationGuide) T_Url(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Url", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Url", &resource.Url, htmlAttrs)
}
func (resource *ImplementationGuide) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Version", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Version", resource.Version, htmlAttrs)
}
func (resource *ImplementationGuide) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Name", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Name", &resource.Name, htmlAttrs)
}
func (resource *ImplementationGuide) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Title", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Title", resource.Title, htmlAttrs)
}
func (resource *ImplementationGuide) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ImplementationGuide.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ImplementationGuide.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_Experimental(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("ImplementationGuide.Experimental", nil, htmlAttrs)
	}
	return BoolInput("ImplementationGuide.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *ImplementationGuide) T_Date(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("ImplementationGuide.Date", nil, htmlAttrs)
	}
	return DateTimeInput("ImplementationGuide.Date", resource.Date, htmlAttrs)
}
func (resource *ImplementationGuide) T_Publisher(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Publisher", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *ImplementationGuide) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Description", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Description", resource.Description, htmlAttrs)
}
func (resource *ImplementationGuide) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("ImplementationGuide.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ImplementationGuide.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_Copyright(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Copyright", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *ImplementationGuide) T_PackageId(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.PackageId", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.PackageId", &resource.PackageId, htmlAttrs)
}
func (resource *ImplementationGuide) T_License(htmlAttrs string) templ.Component {
	optionsValueSet := VSSpdx_license

	if resource == nil {
		return CodeSelect("ImplementationGuide.License", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ImplementationGuide.License", resource.License, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_FhirVersion(numFhirVersion int, htmlAttrs string) templ.Component {
	optionsValueSet := VSFHIR_version

	if resource == nil || numFhirVersion >= len(resource.FhirVersion) {
		return CodeSelect("ImplementationGuide.FhirVersion."+strconv.Itoa(numFhirVersion)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ImplementationGuide.FhirVersion."+strconv.Itoa(numFhirVersion)+".", &resource.FhirVersion[numFhirVersion], optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_DependsOnUri(numDependsOn int, htmlAttrs string) templ.Component {

	if resource == nil || numDependsOn >= len(resource.DependsOn) {
		return StringInput("ImplementationGuide.DependsOn."+strconv.Itoa(numDependsOn)+"..Uri", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.DependsOn."+strconv.Itoa(numDependsOn)+"..Uri", &resource.DependsOn[numDependsOn].Uri, htmlAttrs)
}
func (resource *ImplementationGuide) T_DependsOnPackageId(numDependsOn int, htmlAttrs string) templ.Component {

	if resource == nil || numDependsOn >= len(resource.DependsOn) {
		return StringInput("ImplementationGuide.DependsOn."+strconv.Itoa(numDependsOn)+"..PackageId", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.DependsOn."+strconv.Itoa(numDependsOn)+"..PackageId", resource.DependsOn[numDependsOn].PackageId, htmlAttrs)
}
func (resource *ImplementationGuide) T_DependsOnVersion(numDependsOn int, htmlAttrs string) templ.Component {

	if resource == nil || numDependsOn >= len(resource.DependsOn) {
		return StringInput("ImplementationGuide.DependsOn."+strconv.Itoa(numDependsOn)+"..Version", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.DependsOn."+strconv.Itoa(numDependsOn)+"..Version", resource.DependsOn[numDependsOn].Version, htmlAttrs)
}
func (resource *ImplementationGuide) T_GlobalType(numGlobal int, htmlAttrs string) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numGlobal >= len(resource.Global) {
		return CodeSelect("ImplementationGuide.Global."+strconv.Itoa(numGlobal)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ImplementationGuide.Global."+strconv.Itoa(numGlobal)+"..Type", &resource.Global[numGlobal].Type, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_GlobalProfile(numGlobal int, htmlAttrs string) templ.Component {

	if resource == nil || numGlobal >= len(resource.Global) {
		return StringInput("ImplementationGuide.Global."+strconv.Itoa(numGlobal)+"..Profile", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Global."+strconv.Itoa(numGlobal)+"..Profile", &resource.Global[numGlobal].Profile, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionGroupingName(numGrouping int, htmlAttrs string) templ.Component {

	if resource == nil || numGrouping >= len(resource.Definition.Grouping) {
		return StringInput("ImplementationGuide.Definition.Grouping."+strconv.Itoa(numGrouping)+"..Name", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Definition.Grouping."+strconv.Itoa(numGrouping)+"..Name", &resource.Definition.Grouping[numGrouping].Name, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionGroupingDescription(numGrouping int, htmlAttrs string) templ.Component {

	if resource == nil || numGrouping >= len(resource.Definition.Grouping) {
		return StringInput("ImplementationGuide.Definition.Grouping."+strconv.Itoa(numGrouping)+"..Description", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Definition.Grouping."+strconv.Itoa(numGrouping)+"..Description", resource.Definition.Grouping[numGrouping].Description, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceFhirVersion(numResource int, numFhirVersion int, htmlAttrs string) templ.Component {
	optionsValueSet := VSFHIR_version

	if resource == nil || numResource >= len(resource.Definition.Resource) || numFhirVersion >= len(resource.Definition.Resource[numResource].FhirVersion) {
		return CodeSelect("ImplementationGuide.Definition.Resource."+strconv.Itoa(numResource)+"..FhirVersion."+strconv.Itoa(numFhirVersion)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ImplementationGuide.Definition.Resource."+strconv.Itoa(numResource)+"..FhirVersion."+strconv.Itoa(numFhirVersion)+".", &resource.Definition.Resource[numResource].FhirVersion[numFhirVersion], optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceName(numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numResource >= len(resource.Definition.Resource) {
		return StringInput("ImplementationGuide.Definition.Resource."+strconv.Itoa(numResource)+"..Name", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Definition.Resource."+strconv.Itoa(numResource)+"..Name", resource.Definition.Resource[numResource].Name, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceDescription(numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numResource >= len(resource.Definition.Resource) {
		return StringInput("ImplementationGuide.Definition.Resource."+strconv.Itoa(numResource)+"..Description", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Definition.Resource."+strconv.Itoa(numResource)+"..Description", resource.Definition.Resource[numResource].Description, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceExampleBoolean(numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numResource >= len(resource.Definition.Resource) {
		return BoolInput("ImplementationGuide.Definition.Resource."+strconv.Itoa(numResource)+"..ExampleBoolean", nil, htmlAttrs)
	}
	return BoolInput("ImplementationGuide.Definition.Resource."+strconv.Itoa(numResource)+"..ExampleBoolean", resource.Definition.Resource[numResource].ExampleBoolean, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceExampleCanonical(numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numResource >= len(resource.Definition.Resource) {
		return StringInput("ImplementationGuide.Definition.Resource."+strconv.Itoa(numResource)+"..ExampleCanonical", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Definition.Resource."+strconv.Itoa(numResource)+"..ExampleCanonical", resource.Definition.Resource[numResource].ExampleCanonical, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceGroupingId(numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numResource >= len(resource.Definition.Resource) {
		return StringInput("ImplementationGuide.Definition.Resource."+strconv.Itoa(numResource)+"..GroupingId", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Definition.Resource."+strconv.Itoa(numResource)+"..GroupingId", resource.Definition.Resource[numResource].GroupingId, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionPageNameUrl(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Definition.Page.NameUrl", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Definition.Page.NameUrl", &resource.Definition.Page.NameUrl, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionPageTitle(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Definition.Page.Title", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Definition.Page.Title", &resource.Definition.Page.Title, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionPageGeneration(htmlAttrs string) templ.Component {
	optionsValueSet := VSGuide_page_generation

	if resource == nil {
		return CodeSelect("ImplementationGuide.Definition.Page.Generation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ImplementationGuide.Definition.Page.Generation", &resource.Definition.Page.Generation, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionParameterCode(numParameter int, htmlAttrs string) templ.Component {
	optionsValueSet := VSGuide_parameter_code

	if resource == nil || numParameter >= len(resource.Definition.Parameter) {
		return CodeSelect("ImplementationGuide.Definition.Parameter."+strconv.Itoa(numParameter)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ImplementationGuide.Definition.Parameter."+strconv.Itoa(numParameter)+"..Code", &resource.Definition.Parameter[numParameter].Code, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionParameterValue(numParameter int, htmlAttrs string) templ.Component {

	if resource == nil || numParameter >= len(resource.Definition.Parameter) {
		return StringInput("ImplementationGuide.Definition.Parameter."+strconv.Itoa(numParameter)+"..Value", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Definition.Parameter."+strconv.Itoa(numParameter)+"..Value", &resource.Definition.Parameter[numParameter].Value, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionTemplateCode(numTemplate int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTemplate >= len(resource.Definition.Template) {
		return CodeSelect("ImplementationGuide.Definition.Template."+strconv.Itoa(numTemplate)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ImplementationGuide.Definition.Template."+strconv.Itoa(numTemplate)+"..Code", &resource.Definition.Template[numTemplate].Code, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionTemplateSource(numTemplate int, htmlAttrs string) templ.Component {

	if resource == nil || numTemplate >= len(resource.Definition.Template) {
		return StringInput("ImplementationGuide.Definition.Template."+strconv.Itoa(numTemplate)+"..Source", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Definition.Template."+strconv.Itoa(numTemplate)+"..Source", &resource.Definition.Template[numTemplate].Source, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionTemplateScope(numTemplate int, htmlAttrs string) templ.Component {

	if resource == nil || numTemplate >= len(resource.Definition.Template) {
		return StringInput("ImplementationGuide.Definition.Template."+strconv.Itoa(numTemplate)+"..Scope", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Definition.Template."+strconv.Itoa(numTemplate)+"..Scope", resource.Definition.Template[numTemplate].Scope, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestRendering(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Manifest.Rendering", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Manifest.Rendering", resource.Manifest.Rendering, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestImage(numImage int, htmlAttrs string) templ.Component {

	if resource == nil || numImage >= len(resource.Manifest.Image) {
		return StringInput("ImplementationGuide.Manifest.Image."+strconv.Itoa(numImage)+".", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Manifest.Image."+strconv.Itoa(numImage)+".", &resource.Manifest.Image[numImage], htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestOther(numOther int, htmlAttrs string) templ.Component {

	if resource == nil || numOther >= len(resource.Manifest.Other) {
		return StringInput("ImplementationGuide.Manifest.Other."+strconv.Itoa(numOther)+".", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Manifest.Other."+strconv.Itoa(numOther)+".", &resource.Manifest.Other[numOther], htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestResourceExampleBoolean(numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numResource >= len(resource.Manifest.Resource) {
		return BoolInput("ImplementationGuide.Manifest.Resource."+strconv.Itoa(numResource)+"..ExampleBoolean", nil, htmlAttrs)
	}
	return BoolInput("ImplementationGuide.Manifest.Resource."+strconv.Itoa(numResource)+"..ExampleBoolean", resource.Manifest.Resource[numResource].ExampleBoolean, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestResourceExampleCanonical(numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numResource >= len(resource.Manifest.Resource) {
		return StringInput("ImplementationGuide.Manifest.Resource."+strconv.Itoa(numResource)+"..ExampleCanonical", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Manifest.Resource."+strconv.Itoa(numResource)+"..ExampleCanonical", resource.Manifest.Resource[numResource].ExampleCanonical, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestResourceRelativePath(numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numResource >= len(resource.Manifest.Resource) {
		return StringInput("ImplementationGuide.Manifest.Resource."+strconv.Itoa(numResource)+"..RelativePath", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Manifest.Resource."+strconv.Itoa(numResource)+"..RelativePath", resource.Manifest.Resource[numResource].RelativePath, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestPageName(numPage int, htmlAttrs string) templ.Component {

	if resource == nil || numPage >= len(resource.Manifest.Page) {
		return StringInput("ImplementationGuide.Manifest.Page."+strconv.Itoa(numPage)+"..Name", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Manifest.Page."+strconv.Itoa(numPage)+"..Name", &resource.Manifest.Page[numPage].Name, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestPageTitle(numPage int, htmlAttrs string) templ.Component {

	if resource == nil || numPage >= len(resource.Manifest.Page) {
		return StringInput("ImplementationGuide.Manifest.Page."+strconv.Itoa(numPage)+"..Title", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Manifest.Page."+strconv.Itoa(numPage)+"..Title", resource.Manifest.Page[numPage].Title, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestPageAnchor(numPage int, numAnchor int, htmlAttrs string) templ.Component {

	if resource == nil || numPage >= len(resource.Manifest.Page) || numAnchor >= len(resource.Manifest.Page[numPage].Anchor) {
		return StringInput("ImplementationGuide.Manifest.Page."+strconv.Itoa(numPage)+"..Anchor."+strconv.Itoa(numAnchor)+".", nil, htmlAttrs)
	}
	return StringInput("ImplementationGuide.Manifest.Page."+strconv.Itoa(numPage)+"..Anchor."+strconv.Itoa(numAnchor)+".", &resource.Manifest.Page[numPage].Anchor[numAnchor], htmlAttrs)
}
