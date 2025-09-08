package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ImplementationGuide
type ImplementationGuide struct {
	Id                     *string                        `json:"id,omitempty"`
	Meta                   *Meta                          `json:"meta,omitempty"`
	ImplicitRules          *string                        `json:"implicitRules,omitempty"`
	Language               *string                        `json:"language,omitempty"`
	Text                   *Narrative                     `json:"text,omitempty"`
	Contained              []Resource                     `json:"contained,omitempty"`
	Extension              []Extension                    `json:"extension,omitempty"`
	ModifierExtension      []Extension                    `json:"modifierExtension,omitempty"`
	Url                    string                         `json:"url"`
	Identifier             []Identifier                   `json:"identifier,omitempty"`
	Version                *string                        `json:"version,omitempty"`
	VersionAlgorithmString *string                        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                        `json:"versionAlgorithmCoding,omitempty"`
	Name                   string                         `json:"name"`
	Title                  *string                        `json:"title,omitempty"`
	Status                 string                         `json:"status"`
	Experimental           *bool                          `json:"experimental,omitempty"`
	Date                   *time.Time                     `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher              *string                        `json:"publisher,omitempty"`
	Contact                []ContactDetail                `json:"contact,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	UseContext             []UsageContext                 `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept              `json:"jurisdiction,omitempty"`
	Purpose                *string                        `json:"purpose,omitempty"`
	Copyright              *string                        `json:"copyright,omitempty"`
	CopyrightLabel         *string                        `json:"copyrightLabel,omitempty"`
	PackageId              string                         `json:"packageId"`
	License                *string                        `json:"license,omitempty"`
	FhirVersion            []string                       `json:"fhirVersion"`
	DependsOn              []ImplementationGuideDependsOn `json:"dependsOn,omitempty"`
	Global                 []ImplementationGuideGlobal    `json:"global,omitempty"`
	Definition             *ImplementationGuideDefinition `json:"definition,omitempty"`
	Manifest               *ImplementationGuideManifest   `json:"manifest,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImplementationGuide
type ImplementationGuideDependsOn struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Uri               string      `json:"uri"`
	PackageId         *string     `json:"packageId,omitempty"`
	Version           *string     `json:"version,omitempty"`
	Reason            *string     `json:"reason,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImplementationGuide
type ImplementationGuideGlobal struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Profile           string      `json:"profile"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinition struct {
	Id                *string                                  `json:"id,omitempty"`
	Extension         []Extension                              `json:"extension,omitempty"`
	ModifierExtension []Extension                              `json:"modifierExtension,omitempty"`
	Grouping          []ImplementationGuideDefinitionGrouping  `json:"grouping,omitempty"`
	Resource          []ImplementationGuideDefinitionResource  `json:"resource,omitempty"`
	Page              *ImplementationGuideDefinitionPage       `json:"page,omitempty"`
	Parameter         []ImplementationGuideDefinitionParameter `json:"parameter,omitempty"`
	Template          []ImplementationGuideDefinitionTemplate  `json:"template,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinitionGrouping struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Description       *string     `json:"description,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinitionResource struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Reference         Reference   `json:"reference"`
	FhirVersion       []string    `json:"fhirVersion,omitempty"`
	Name              *string     `json:"name,omitempty"`
	Description       *string     `json:"description,omitempty"`
	IsExample         *bool       `json:"isExample,omitempty"`
	Profile           []string    `json:"profile,omitempty"`
	GroupingId        *string     `json:"groupingId,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinitionPage struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	SourceUrl         *string     `json:"sourceUrl,omitempty"`
	SourceString      *string     `json:"sourceString,omitempty"`
	SourceMarkdown    *string     `json:"sourceMarkdown,omitempty"`
	Name              string      `json:"name"`
	Title             string      `json:"title"`
	Generation        string      `json:"generation"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinitionParameter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              Coding      `json:"code"`
	Value             string      `json:"value"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinitionTemplate struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Source            string      `json:"source"`
	Scope             *string     `json:"scope,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImplementationGuide
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

// http://hl7.org/fhir/r5/StructureDefinition/ImplementationGuide
type ImplementationGuideManifestResource struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Reference         Reference   `json:"reference"`
	IsExample         *bool       `json:"isExample,omitempty"`
	Profile           []string    `json:"profile,omitempty"`
	RelativePath      *string     `json:"relativePath,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImplementationGuide
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
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ImplementationGuide"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ImplementationGuide) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", &resource.Url, htmlAttrs)
}
func (resource *ImplementationGuide) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *ImplementationGuide) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *ImplementationGuide) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", &resource.Name, htmlAttrs)
}
func (resource *ImplementationGuide) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *ImplementationGuide) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Experimental", nil, htmlAttrs)
	}
	return BoolInput("Experimental", resource.Experimental, htmlAttrs)
}
func (resource *ImplementationGuide) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *ImplementationGuide) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *ImplementationGuide) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *ImplementationGuide) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Purpose", nil, htmlAttrs)
	}
	return StringInput("Purpose", resource.Purpose, htmlAttrs)
}
func (resource *ImplementationGuide) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Copyright", nil, htmlAttrs)
	}
	return StringInput("Copyright", resource.Copyright, htmlAttrs)
}
func (resource *ImplementationGuide) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *ImplementationGuide) T_PackageId(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PackageId", nil, htmlAttrs)
	}
	return StringInput("PackageId", &resource.PackageId, htmlAttrs)
}
func (resource *ImplementationGuide) T_License(htmlAttrs string) templ.Component {
	optionsValueSet := VSSpdx_license

	if resource == nil {
		return CodeSelect("License", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("License", resource.License, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_FhirVersion(numFhirVersion int, htmlAttrs string) templ.Component {
	optionsValueSet := VSFHIR_version

	if resource == nil || numFhirVersion >= len(resource.FhirVersion) {
		return CodeSelect("FhirVersion["+strconv.Itoa(numFhirVersion)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("FhirVersion["+strconv.Itoa(numFhirVersion)+"]", &resource.FhirVersion[numFhirVersion], optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_DependsOnUri(numDependsOn int, htmlAttrs string) templ.Component {
	if resource == nil || numDependsOn >= len(resource.DependsOn) {
		return StringInput("DependsOn["+strconv.Itoa(numDependsOn)+"]Uri", nil, htmlAttrs)
	}
	return StringInput("DependsOn["+strconv.Itoa(numDependsOn)+"]Uri", &resource.DependsOn[numDependsOn].Uri, htmlAttrs)
}
func (resource *ImplementationGuide) T_DependsOnPackageId(numDependsOn int, htmlAttrs string) templ.Component {
	if resource == nil || numDependsOn >= len(resource.DependsOn) {
		return StringInput("DependsOn["+strconv.Itoa(numDependsOn)+"]PackageId", nil, htmlAttrs)
	}
	return StringInput("DependsOn["+strconv.Itoa(numDependsOn)+"]PackageId", resource.DependsOn[numDependsOn].PackageId, htmlAttrs)
}
func (resource *ImplementationGuide) T_DependsOnVersion(numDependsOn int, htmlAttrs string) templ.Component {
	if resource == nil || numDependsOn >= len(resource.DependsOn) {
		return StringInput("DependsOn["+strconv.Itoa(numDependsOn)+"]Version", nil, htmlAttrs)
	}
	return StringInput("DependsOn["+strconv.Itoa(numDependsOn)+"]Version", resource.DependsOn[numDependsOn].Version, htmlAttrs)
}
func (resource *ImplementationGuide) T_DependsOnReason(numDependsOn int, htmlAttrs string) templ.Component {
	if resource == nil || numDependsOn >= len(resource.DependsOn) {
		return StringInput("DependsOn["+strconv.Itoa(numDependsOn)+"]Reason", nil, htmlAttrs)
	}
	return StringInput("DependsOn["+strconv.Itoa(numDependsOn)+"]Reason", resource.DependsOn[numDependsOn].Reason, htmlAttrs)
}
func (resource *ImplementationGuide) T_GlobalType(numGlobal int, htmlAttrs string) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numGlobal >= len(resource.Global) {
		return CodeSelect("Global["+strconv.Itoa(numGlobal)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Global["+strconv.Itoa(numGlobal)+"]Type", &resource.Global[numGlobal].Type, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_GlobalProfile(numGlobal int, htmlAttrs string) templ.Component {
	if resource == nil || numGlobal >= len(resource.Global) {
		return StringInput("Global["+strconv.Itoa(numGlobal)+"]Profile", nil, htmlAttrs)
	}
	return StringInput("Global["+strconv.Itoa(numGlobal)+"]Profile", &resource.Global[numGlobal].Profile, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionGroupingName(numGrouping int, htmlAttrs string) templ.Component {
	if resource == nil || numGrouping >= len(resource.Definition.Grouping) {
		return StringInput("DefinitionGrouping["+strconv.Itoa(numGrouping)+"].Name", nil, htmlAttrs)
	}
	return StringInput("DefinitionGrouping["+strconv.Itoa(numGrouping)+"].Name", &resource.Definition.Grouping[numGrouping].Name, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionGroupingDescription(numGrouping int, htmlAttrs string) templ.Component {
	if resource == nil || numGrouping >= len(resource.Definition.Grouping) {
		return StringInput("DefinitionGrouping["+strconv.Itoa(numGrouping)+"].Description", nil, htmlAttrs)
	}
	return StringInput("DefinitionGrouping["+strconv.Itoa(numGrouping)+"].Description", resource.Definition.Grouping[numGrouping].Description, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceFhirVersion(numResource int, numFhirVersion int, htmlAttrs string) templ.Component {
	optionsValueSet := VSFHIR_version

	if resource == nil || numResource >= len(resource.Definition.Resource) || numFhirVersion >= len(resource.Definition.Resource[numResource].FhirVersion) {
		return CodeSelect("DefinitionResource["+strconv.Itoa(numResource)+"].FhirVersion["+strconv.Itoa(numFhirVersion)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DefinitionResource["+strconv.Itoa(numResource)+"].FhirVersion["+strconv.Itoa(numFhirVersion)+"]", &resource.Definition.Resource[numResource].FhirVersion[numFhirVersion], optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceName(numResource int, htmlAttrs string) templ.Component {
	if resource == nil || numResource >= len(resource.Definition.Resource) {
		return StringInput("DefinitionResource["+strconv.Itoa(numResource)+"].Name", nil, htmlAttrs)
	}
	return StringInput("DefinitionResource["+strconv.Itoa(numResource)+"].Name", resource.Definition.Resource[numResource].Name, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceDescription(numResource int, htmlAttrs string) templ.Component {
	if resource == nil || numResource >= len(resource.Definition.Resource) {
		return StringInput("DefinitionResource["+strconv.Itoa(numResource)+"].Description", nil, htmlAttrs)
	}
	return StringInput("DefinitionResource["+strconv.Itoa(numResource)+"].Description", resource.Definition.Resource[numResource].Description, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceIsExample(numResource int, htmlAttrs string) templ.Component {
	if resource == nil || numResource >= len(resource.Definition.Resource) {
		return BoolInput("DefinitionResource["+strconv.Itoa(numResource)+"].IsExample", nil, htmlAttrs)
	}
	return BoolInput("DefinitionResource["+strconv.Itoa(numResource)+"].IsExample", resource.Definition.Resource[numResource].IsExample, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceProfile(numResource int, numProfile int, htmlAttrs string) templ.Component {
	if resource == nil || numResource >= len(resource.Definition.Resource) || numProfile >= len(resource.Definition.Resource[numResource].Profile) {
		return StringInput("DefinitionResource["+strconv.Itoa(numResource)+"].Profile["+strconv.Itoa(numProfile)+"]", nil, htmlAttrs)
	}
	return StringInput("DefinitionResource["+strconv.Itoa(numResource)+"].Profile["+strconv.Itoa(numProfile)+"]", &resource.Definition.Resource[numResource].Profile[numProfile], htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionResourceGroupingId(numResource int, htmlAttrs string) templ.Component {
	if resource == nil || numResource >= len(resource.Definition.Resource) {
		return StringInput("DefinitionResource["+strconv.Itoa(numResource)+"].GroupingId", nil, htmlAttrs)
	}
	return StringInput("DefinitionResource["+strconv.Itoa(numResource)+"].GroupingId", resource.Definition.Resource[numResource].GroupingId, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionPageSourceUrl(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DefinitionPage.SourceUrl", nil, htmlAttrs)
	}
	return StringInput("DefinitionPage.SourceUrl", resource.Definition.Page.SourceUrl, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionPageSourceString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DefinitionPage.SourceString", nil, htmlAttrs)
	}
	return StringInput("DefinitionPage.SourceString", resource.Definition.Page.SourceString, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionPageSourceMarkdown(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DefinitionPage.SourceMarkdown", nil, htmlAttrs)
	}
	return StringInput("DefinitionPage.SourceMarkdown", resource.Definition.Page.SourceMarkdown, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionPageName(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DefinitionPage.Name", nil, htmlAttrs)
	}
	return StringInput("DefinitionPage.Name", &resource.Definition.Page.Name, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionPageTitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DefinitionPage.Title", nil, htmlAttrs)
	}
	return StringInput("DefinitionPage.Title", &resource.Definition.Page.Title, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionPageGeneration(htmlAttrs string) templ.Component {
	optionsValueSet := VSGuide_page_generation

	if resource == nil {
		return CodeSelect("DefinitionPage.Generation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DefinitionPage.Generation", &resource.Definition.Page.Generation, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionParameterCode(numParameter int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Definition.Parameter) {
		return CodingSelect("DefinitionParameter["+strconv.Itoa(numParameter)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("DefinitionParameter["+strconv.Itoa(numParameter)+"].Code", &resource.Definition.Parameter[numParameter].Code, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionParameterValue(numParameter int, htmlAttrs string) templ.Component {
	if resource == nil || numParameter >= len(resource.Definition.Parameter) {
		return StringInput("DefinitionParameter["+strconv.Itoa(numParameter)+"].Value", nil, htmlAttrs)
	}
	return StringInput("DefinitionParameter["+strconv.Itoa(numParameter)+"].Value", &resource.Definition.Parameter[numParameter].Value, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionTemplateCode(numTemplate int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTemplate >= len(resource.Definition.Template) {
		return CodeSelect("DefinitionTemplate["+strconv.Itoa(numTemplate)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DefinitionTemplate["+strconv.Itoa(numTemplate)+"].Code", &resource.Definition.Template[numTemplate].Code, optionsValueSet, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionTemplateSource(numTemplate int, htmlAttrs string) templ.Component {
	if resource == nil || numTemplate >= len(resource.Definition.Template) {
		return StringInput("DefinitionTemplate["+strconv.Itoa(numTemplate)+"].Source", nil, htmlAttrs)
	}
	return StringInput("DefinitionTemplate["+strconv.Itoa(numTemplate)+"].Source", &resource.Definition.Template[numTemplate].Source, htmlAttrs)
}
func (resource *ImplementationGuide) T_DefinitionTemplateScope(numTemplate int, htmlAttrs string) templ.Component {
	if resource == nil || numTemplate >= len(resource.Definition.Template) {
		return StringInput("DefinitionTemplate["+strconv.Itoa(numTemplate)+"].Scope", nil, htmlAttrs)
	}
	return StringInput("DefinitionTemplate["+strconv.Itoa(numTemplate)+"].Scope", resource.Definition.Template[numTemplate].Scope, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestRendering(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ManifestRendering", nil, htmlAttrs)
	}
	return StringInput("ManifestRendering", resource.Manifest.Rendering, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestImage(numImage int, htmlAttrs string) templ.Component {
	if resource == nil || numImage >= len(resource.Manifest.Image) {
		return StringInput("ManifestImage["+strconv.Itoa(numImage)+"]", nil, htmlAttrs)
	}
	return StringInput("ManifestImage["+strconv.Itoa(numImage)+"]", &resource.Manifest.Image[numImage], htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestOther(numOther int, htmlAttrs string) templ.Component {
	if resource == nil || numOther >= len(resource.Manifest.Other) {
		return StringInput("ManifestOther["+strconv.Itoa(numOther)+"]", nil, htmlAttrs)
	}
	return StringInput("ManifestOther["+strconv.Itoa(numOther)+"]", &resource.Manifest.Other[numOther], htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestResourceIsExample(numResource int, htmlAttrs string) templ.Component {
	if resource == nil || numResource >= len(resource.Manifest.Resource) {
		return BoolInput("ManifestResource["+strconv.Itoa(numResource)+"].IsExample", nil, htmlAttrs)
	}
	return BoolInput("ManifestResource["+strconv.Itoa(numResource)+"].IsExample", resource.Manifest.Resource[numResource].IsExample, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestResourceProfile(numResource int, numProfile int, htmlAttrs string) templ.Component {
	if resource == nil || numResource >= len(resource.Manifest.Resource) || numProfile >= len(resource.Manifest.Resource[numResource].Profile) {
		return StringInput("ManifestResource["+strconv.Itoa(numResource)+"].Profile["+strconv.Itoa(numProfile)+"]", nil, htmlAttrs)
	}
	return StringInput("ManifestResource["+strconv.Itoa(numResource)+"].Profile["+strconv.Itoa(numProfile)+"]", &resource.Manifest.Resource[numResource].Profile[numProfile], htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestResourceRelativePath(numResource int, htmlAttrs string) templ.Component {
	if resource == nil || numResource >= len(resource.Manifest.Resource) {
		return StringInput("ManifestResource["+strconv.Itoa(numResource)+"].RelativePath", nil, htmlAttrs)
	}
	return StringInput("ManifestResource["+strconv.Itoa(numResource)+"].RelativePath", resource.Manifest.Resource[numResource].RelativePath, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestPageName(numPage int, htmlAttrs string) templ.Component {
	if resource == nil || numPage >= len(resource.Manifest.Page) {
		return StringInput("ManifestPage["+strconv.Itoa(numPage)+"].Name", nil, htmlAttrs)
	}
	return StringInput("ManifestPage["+strconv.Itoa(numPage)+"].Name", &resource.Manifest.Page[numPage].Name, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestPageTitle(numPage int, htmlAttrs string) templ.Component {
	if resource == nil || numPage >= len(resource.Manifest.Page) {
		return StringInput("ManifestPage["+strconv.Itoa(numPage)+"].Title", nil, htmlAttrs)
	}
	return StringInput("ManifestPage["+strconv.Itoa(numPage)+"].Title", resource.Manifest.Page[numPage].Title, htmlAttrs)
}
func (resource *ImplementationGuide) T_ManifestPageAnchor(numPage int, numAnchor int, htmlAttrs string) templ.Component {
	if resource == nil || numPage >= len(resource.Manifest.Page) || numAnchor >= len(resource.Manifest.Page[numPage].Anchor) {
		return StringInput("ManifestPage["+strconv.Itoa(numPage)+"].Anchor["+strconv.Itoa(numAnchor)+"]", nil, htmlAttrs)
	}
	return StringInput("ManifestPage["+strconv.Itoa(numPage)+"].Anchor["+strconv.Itoa(numAnchor)+"]", &resource.Manifest.Page[numPage].Anchor[numAnchor], htmlAttrs)
}
