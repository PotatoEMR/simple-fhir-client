package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Date                   *string                        `json:"date,omitempty"`
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

func (resource *ImplementationGuide) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Id", nil)
	}
	return StringInput("ImplementationGuide.Id", resource.Id)
}
func (resource *ImplementationGuide) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.ImplicitRules", nil)
	}
	return StringInput("ImplementationGuide.ImplicitRules", resource.ImplicitRules)
}
func (resource *ImplementationGuide) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ImplementationGuide.Language", nil, optionsValueSet)
	}
	return CodeSelect("ImplementationGuide.Language", resource.Language, optionsValueSet)
}
func (resource *ImplementationGuide) T_Url() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Url", nil)
	}
	return StringInput("ImplementationGuide.Url", &resource.Url)
}
func (resource *ImplementationGuide) T_Version() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Version", nil)
	}
	return StringInput("ImplementationGuide.Version", resource.Version)
}
func (resource *ImplementationGuide) T_Name() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Name", nil)
	}
	return StringInput("ImplementationGuide.Name", &resource.Name)
}
func (resource *ImplementationGuide) T_Title() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Title", nil)
	}
	return StringInput("ImplementationGuide.Title", resource.Title)
}
func (resource *ImplementationGuide) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ImplementationGuide.Status", nil, optionsValueSet)
	}
	return CodeSelect("ImplementationGuide.Status", &resource.Status, optionsValueSet)
}
func (resource *ImplementationGuide) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("ImplementationGuide.Experimental", nil)
	}
	return BoolInput("ImplementationGuide.Experimental", resource.Experimental)
}
func (resource *ImplementationGuide) T_Date() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Date", nil)
	}
	return StringInput("ImplementationGuide.Date", resource.Date)
}
func (resource *ImplementationGuide) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Publisher", nil)
	}
	return StringInput("ImplementationGuide.Publisher", resource.Publisher)
}
func (resource *ImplementationGuide) T_Description() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Description", nil)
	}
	return StringInput("ImplementationGuide.Description", resource.Description)
}
func (resource *ImplementationGuide) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("ImplementationGuide.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ImplementationGuide.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *ImplementationGuide) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Purpose", nil)
	}
	return StringInput("ImplementationGuide.Purpose", resource.Purpose)
}
func (resource *ImplementationGuide) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Copyright", nil)
	}
	return StringInput("ImplementationGuide.Copyright", resource.Copyright)
}
func (resource *ImplementationGuide) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.CopyrightLabel", nil)
	}
	return StringInput("ImplementationGuide.CopyrightLabel", resource.CopyrightLabel)
}
func (resource *ImplementationGuide) T_PackageId() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.PackageId", nil)
	}
	return StringInput("ImplementationGuide.PackageId", &resource.PackageId)
}
func (resource *ImplementationGuide) T_License() templ.Component {
	optionsValueSet := VSSpdx_license

	if resource == nil {
		return CodeSelect("ImplementationGuide.License", nil, optionsValueSet)
	}
	return CodeSelect("ImplementationGuide.License", resource.License, optionsValueSet)
}
func (resource *ImplementationGuide) T_FhirVersion(numFhirVersion int) templ.Component {
	optionsValueSet := VSFHIR_version

	if resource == nil || len(resource.FhirVersion) >= numFhirVersion {
		return CodeSelect("ImplementationGuide.FhirVersion["+strconv.Itoa(numFhirVersion)+"]", nil, optionsValueSet)
	}
	return CodeSelect("ImplementationGuide.FhirVersion["+strconv.Itoa(numFhirVersion)+"]", &resource.FhirVersion[numFhirVersion], optionsValueSet)
}
func (resource *ImplementationGuide) T_DependsOnId(numDependsOn int) templ.Component {

	if resource == nil || len(resource.DependsOn) >= numDependsOn {
		return StringInput("ImplementationGuide.DependsOn["+strconv.Itoa(numDependsOn)+"].Id", nil)
	}
	return StringInput("ImplementationGuide.DependsOn["+strconv.Itoa(numDependsOn)+"].Id", resource.DependsOn[numDependsOn].Id)
}
func (resource *ImplementationGuide) T_DependsOnUri(numDependsOn int) templ.Component {

	if resource == nil || len(resource.DependsOn) >= numDependsOn {
		return StringInput("ImplementationGuide.DependsOn["+strconv.Itoa(numDependsOn)+"].Uri", nil)
	}
	return StringInput("ImplementationGuide.DependsOn["+strconv.Itoa(numDependsOn)+"].Uri", &resource.DependsOn[numDependsOn].Uri)
}
func (resource *ImplementationGuide) T_DependsOnPackageId(numDependsOn int) templ.Component {

	if resource == nil || len(resource.DependsOn) >= numDependsOn {
		return StringInput("ImplementationGuide.DependsOn["+strconv.Itoa(numDependsOn)+"].PackageId", nil)
	}
	return StringInput("ImplementationGuide.DependsOn["+strconv.Itoa(numDependsOn)+"].PackageId", resource.DependsOn[numDependsOn].PackageId)
}
func (resource *ImplementationGuide) T_DependsOnVersion(numDependsOn int) templ.Component {

	if resource == nil || len(resource.DependsOn) >= numDependsOn {
		return StringInput("ImplementationGuide.DependsOn["+strconv.Itoa(numDependsOn)+"].Version", nil)
	}
	return StringInput("ImplementationGuide.DependsOn["+strconv.Itoa(numDependsOn)+"].Version", resource.DependsOn[numDependsOn].Version)
}
func (resource *ImplementationGuide) T_DependsOnReason(numDependsOn int) templ.Component {

	if resource == nil || len(resource.DependsOn) >= numDependsOn {
		return StringInput("ImplementationGuide.DependsOn["+strconv.Itoa(numDependsOn)+"].Reason", nil)
	}
	return StringInput("ImplementationGuide.DependsOn["+strconv.Itoa(numDependsOn)+"].Reason", resource.DependsOn[numDependsOn].Reason)
}
func (resource *ImplementationGuide) T_GlobalId(numGlobal int) templ.Component {

	if resource == nil || len(resource.Global) >= numGlobal {
		return StringInput("ImplementationGuide.Global["+strconv.Itoa(numGlobal)+"].Id", nil)
	}
	return StringInput("ImplementationGuide.Global["+strconv.Itoa(numGlobal)+"].Id", resource.Global[numGlobal].Id)
}
func (resource *ImplementationGuide) T_GlobalType(numGlobal int) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || len(resource.Global) >= numGlobal {
		return CodeSelect("ImplementationGuide.Global["+strconv.Itoa(numGlobal)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("ImplementationGuide.Global["+strconv.Itoa(numGlobal)+"].Type", &resource.Global[numGlobal].Type, optionsValueSet)
}
func (resource *ImplementationGuide) T_GlobalProfile(numGlobal int) templ.Component {

	if resource == nil || len(resource.Global) >= numGlobal {
		return StringInput("ImplementationGuide.Global["+strconv.Itoa(numGlobal)+"].Profile", nil)
	}
	return StringInput("ImplementationGuide.Global["+strconv.Itoa(numGlobal)+"].Profile", &resource.Global[numGlobal].Profile)
}
func (resource *ImplementationGuide) T_DefinitionId() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Definition.Id", nil)
	}
	return StringInput("ImplementationGuide.Definition.Id", resource.Definition.Id)
}
func (resource *ImplementationGuide) T_DefinitionGroupingId(numGrouping int) templ.Component {

	if resource == nil || len(resource.Definition.Grouping) >= numGrouping {
		return StringInput("ImplementationGuide.Definition.Grouping["+strconv.Itoa(numGrouping)+"].Id", nil)
	}
	return StringInput("ImplementationGuide.Definition.Grouping["+strconv.Itoa(numGrouping)+"].Id", resource.Definition.Grouping[numGrouping].Id)
}
func (resource *ImplementationGuide) T_DefinitionGroupingName(numGrouping int) templ.Component {

	if resource == nil || len(resource.Definition.Grouping) >= numGrouping {
		return StringInput("ImplementationGuide.Definition.Grouping["+strconv.Itoa(numGrouping)+"].Name", nil)
	}
	return StringInput("ImplementationGuide.Definition.Grouping["+strconv.Itoa(numGrouping)+"].Name", &resource.Definition.Grouping[numGrouping].Name)
}
func (resource *ImplementationGuide) T_DefinitionGroupingDescription(numGrouping int) templ.Component {

	if resource == nil || len(resource.Definition.Grouping) >= numGrouping {
		return StringInput("ImplementationGuide.Definition.Grouping["+strconv.Itoa(numGrouping)+"].Description", nil)
	}
	return StringInput("ImplementationGuide.Definition.Grouping["+strconv.Itoa(numGrouping)+"].Description", resource.Definition.Grouping[numGrouping].Description)
}
func (resource *ImplementationGuide) T_DefinitionResourceId(numResource int) templ.Component {

	if resource == nil || len(resource.Definition.Resource) >= numResource {
		return StringInput("ImplementationGuide.Definition.Resource["+strconv.Itoa(numResource)+"].Id", nil)
	}
	return StringInput("ImplementationGuide.Definition.Resource["+strconv.Itoa(numResource)+"].Id", resource.Definition.Resource[numResource].Id)
}
func (resource *ImplementationGuide) T_DefinitionResourceFhirVersion(numResource int, numFhirVersion int) templ.Component {
	optionsValueSet := VSFHIR_version

	if resource == nil || len(resource.Definition.Resource) >= numResource || len(resource.Definition.Resource[numResource].FhirVersion) >= numFhirVersion {
		return CodeSelect("ImplementationGuide.Definition.Resource["+strconv.Itoa(numResource)+"].FhirVersion["+strconv.Itoa(numFhirVersion)+"]", nil, optionsValueSet)
	}
	return CodeSelect("ImplementationGuide.Definition.Resource["+strconv.Itoa(numResource)+"].FhirVersion["+strconv.Itoa(numFhirVersion)+"]", &resource.Definition.Resource[numResource].FhirVersion[numFhirVersion], optionsValueSet)
}
func (resource *ImplementationGuide) T_DefinitionResourceName(numResource int) templ.Component {

	if resource == nil || len(resource.Definition.Resource) >= numResource {
		return StringInput("ImplementationGuide.Definition.Resource["+strconv.Itoa(numResource)+"].Name", nil)
	}
	return StringInput("ImplementationGuide.Definition.Resource["+strconv.Itoa(numResource)+"].Name", resource.Definition.Resource[numResource].Name)
}
func (resource *ImplementationGuide) T_DefinitionResourceDescription(numResource int) templ.Component {

	if resource == nil || len(resource.Definition.Resource) >= numResource {
		return StringInput("ImplementationGuide.Definition.Resource["+strconv.Itoa(numResource)+"].Description", nil)
	}
	return StringInput("ImplementationGuide.Definition.Resource["+strconv.Itoa(numResource)+"].Description", resource.Definition.Resource[numResource].Description)
}
func (resource *ImplementationGuide) T_DefinitionResourceIsExample(numResource int) templ.Component {

	if resource == nil || len(resource.Definition.Resource) >= numResource {
		return BoolInput("ImplementationGuide.Definition.Resource["+strconv.Itoa(numResource)+"].IsExample", nil)
	}
	return BoolInput("ImplementationGuide.Definition.Resource["+strconv.Itoa(numResource)+"].IsExample", resource.Definition.Resource[numResource].IsExample)
}
func (resource *ImplementationGuide) T_DefinitionResourceProfile(numResource int, numProfile int) templ.Component {

	if resource == nil || len(resource.Definition.Resource) >= numResource || len(resource.Definition.Resource[numResource].Profile) >= numProfile {
		return StringInput("ImplementationGuide.Definition.Resource["+strconv.Itoa(numResource)+"].Profile["+strconv.Itoa(numProfile)+"]", nil)
	}
	return StringInput("ImplementationGuide.Definition.Resource["+strconv.Itoa(numResource)+"].Profile["+strconv.Itoa(numProfile)+"]", &resource.Definition.Resource[numResource].Profile[numProfile])
}
func (resource *ImplementationGuide) T_DefinitionResourceGroupingId(numResource int) templ.Component {

	if resource == nil || len(resource.Definition.Resource) >= numResource {
		return StringInput("ImplementationGuide.Definition.Resource["+strconv.Itoa(numResource)+"].GroupingId", nil)
	}
	return StringInput("ImplementationGuide.Definition.Resource["+strconv.Itoa(numResource)+"].GroupingId", resource.Definition.Resource[numResource].GroupingId)
}
func (resource *ImplementationGuide) T_DefinitionPageId() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Definition.Page.Id", nil)
	}
	return StringInput("ImplementationGuide.Definition.Page.Id", resource.Definition.Page.Id)
}
func (resource *ImplementationGuide) T_DefinitionPageName() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Definition.Page.Name", nil)
	}
	return StringInput("ImplementationGuide.Definition.Page.Name", &resource.Definition.Page.Name)
}
func (resource *ImplementationGuide) T_DefinitionPageTitle() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Definition.Page.Title", nil)
	}
	return StringInput("ImplementationGuide.Definition.Page.Title", &resource.Definition.Page.Title)
}
func (resource *ImplementationGuide) T_DefinitionPageGeneration() templ.Component {
	optionsValueSet := VSGuide_page_generation

	if resource == nil {
		return CodeSelect("ImplementationGuide.Definition.Page.Generation", nil, optionsValueSet)
	}
	return CodeSelect("ImplementationGuide.Definition.Page.Generation", &resource.Definition.Page.Generation, optionsValueSet)
}
func (resource *ImplementationGuide) T_DefinitionParameterId(numParameter int) templ.Component {

	if resource == nil || len(resource.Definition.Parameter) >= numParameter {
		return StringInput("ImplementationGuide.Definition.Parameter["+strconv.Itoa(numParameter)+"].Id", nil)
	}
	return StringInput("ImplementationGuide.Definition.Parameter["+strconv.Itoa(numParameter)+"].Id", resource.Definition.Parameter[numParameter].Id)
}
func (resource *ImplementationGuide) T_DefinitionParameterCode(numParameter int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Definition.Parameter) >= numParameter {
		return CodingSelect("ImplementationGuide.Definition.Parameter["+strconv.Itoa(numParameter)+"].Code", nil, optionsValueSet)
	}
	return CodingSelect("ImplementationGuide.Definition.Parameter["+strconv.Itoa(numParameter)+"].Code", &resource.Definition.Parameter[numParameter].Code, optionsValueSet)
}
func (resource *ImplementationGuide) T_DefinitionParameterValue(numParameter int) templ.Component {

	if resource == nil || len(resource.Definition.Parameter) >= numParameter {
		return StringInput("ImplementationGuide.Definition.Parameter["+strconv.Itoa(numParameter)+"].Value", nil)
	}
	return StringInput("ImplementationGuide.Definition.Parameter["+strconv.Itoa(numParameter)+"].Value", &resource.Definition.Parameter[numParameter].Value)
}
func (resource *ImplementationGuide) T_DefinitionTemplateId(numTemplate int) templ.Component {

	if resource == nil || len(resource.Definition.Template) >= numTemplate {
		return StringInput("ImplementationGuide.Definition.Template["+strconv.Itoa(numTemplate)+"].Id", nil)
	}
	return StringInput("ImplementationGuide.Definition.Template["+strconv.Itoa(numTemplate)+"].Id", resource.Definition.Template[numTemplate].Id)
}
func (resource *ImplementationGuide) T_DefinitionTemplateCode(numTemplate int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Definition.Template) >= numTemplate {
		return CodeSelect("ImplementationGuide.Definition.Template["+strconv.Itoa(numTemplate)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("ImplementationGuide.Definition.Template["+strconv.Itoa(numTemplate)+"].Code", &resource.Definition.Template[numTemplate].Code, optionsValueSet)
}
func (resource *ImplementationGuide) T_DefinitionTemplateSource(numTemplate int) templ.Component {

	if resource == nil || len(resource.Definition.Template) >= numTemplate {
		return StringInput("ImplementationGuide.Definition.Template["+strconv.Itoa(numTemplate)+"].Source", nil)
	}
	return StringInput("ImplementationGuide.Definition.Template["+strconv.Itoa(numTemplate)+"].Source", &resource.Definition.Template[numTemplate].Source)
}
func (resource *ImplementationGuide) T_DefinitionTemplateScope(numTemplate int) templ.Component {

	if resource == nil || len(resource.Definition.Template) >= numTemplate {
		return StringInput("ImplementationGuide.Definition.Template["+strconv.Itoa(numTemplate)+"].Scope", nil)
	}
	return StringInput("ImplementationGuide.Definition.Template["+strconv.Itoa(numTemplate)+"].Scope", resource.Definition.Template[numTemplate].Scope)
}
func (resource *ImplementationGuide) T_ManifestId() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Manifest.Id", nil)
	}
	return StringInput("ImplementationGuide.Manifest.Id", resource.Manifest.Id)
}
func (resource *ImplementationGuide) T_ManifestRendering() templ.Component {

	if resource == nil {
		return StringInput("ImplementationGuide.Manifest.Rendering", nil)
	}
	return StringInput("ImplementationGuide.Manifest.Rendering", resource.Manifest.Rendering)
}
func (resource *ImplementationGuide) T_ManifestImage(numImage int) templ.Component {

	if resource == nil || len(resource.Manifest.Image) >= numImage {
		return StringInput("ImplementationGuide.Manifest.Image["+strconv.Itoa(numImage)+"]", nil)
	}
	return StringInput("ImplementationGuide.Manifest.Image["+strconv.Itoa(numImage)+"]", &resource.Manifest.Image[numImage])
}
func (resource *ImplementationGuide) T_ManifestOther(numOther int) templ.Component {

	if resource == nil || len(resource.Manifest.Other) >= numOther {
		return StringInput("ImplementationGuide.Manifest.Other["+strconv.Itoa(numOther)+"]", nil)
	}
	return StringInput("ImplementationGuide.Manifest.Other["+strconv.Itoa(numOther)+"]", &resource.Manifest.Other[numOther])
}
func (resource *ImplementationGuide) T_ManifestResourceId(numResource int) templ.Component {

	if resource == nil || len(resource.Manifest.Resource) >= numResource {
		return StringInput("ImplementationGuide.Manifest.Resource["+strconv.Itoa(numResource)+"].Id", nil)
	}
	return StringInput("ImplementationGuide.Manifest.Resource["+strconv.Itoa(numResource)+"].Id", resource.Manifest.Resource[numResource].Id)
}
func (resource *ImplementationGuide) T_ManifestResourceIsExample(numResource int) templ.Component {

	if resource == nil || len(resource.Manifest.Resource) >= numResource {
		return BoolInput("ImplementationGuide.Manifest.Resource["+strconv.Itoa(numResource)+"].IsExample", nil)
	}
	return BoolInput("ImplementationGuide.Manifest.Resource["+strconv.Itoa(numResource)+"].IsExample", resource.Manifest.Resource[numResource].IsExample)
}
func (resource *ImplementationGuide) T_ManifestResourceProfile(numResource int, numProfile int) templ.Component {

	if resource == nil || len(resource.Manifest.Resource) >= numResource || len(resource.Manifest.Resource[numResource].Profile) >= numProfile {
		return StringInput("ImplementationGuide.Manifest.Resource["+strconv.Itoa(numResource)+"].Profile["+strconv.Itoa(numProfile)+"]", nil)
	}
	return StringInput("ImplementationGuide.Manifest.Resource["+strconv.Itoa(numResource)+"].Profile["+strconv.Itoa(numProfile)+"]", &resource.Manifest.Resource[numResource].Profile[numProfile])
}
func (resource *ImplementationGuide) T_ManifestResourceRelativePath(numResource int) templ.Component {

	if resource == nil || len(resource.Manifest.Resource) >= numResource {
		return StringInput("ImplementationGuide.Manifest.Resource["+strconv.Itoa(numResource)+"].RelativePath", nil)
	}
	return StringInput("ImplementationGuide.Manifest.Resource["+strconv.Itoa(numResource)+"].RelativePath", resource.Manifest.Resource[numResource].RelativePath)
}
func (resource *ImplementationGuide) T_ManifestPageId(numPage int) templ.Component {

	if resource == nil || len(resource.Manifest.Page) >= numPage {
		return StringInput("ImplementationGuide.Manifest.Page["+strconv.Itoa(numPage)+"].Id", nil)
	}
	return StringInput("ImplementationGuide.Manifest.Page["+strconv.Itoa(numPage)+"].Id", resource.Manifest.Page[numPage].Id)
}
func (resource *ImplementationGuide) T_ManifestPageName(numPage int) templ.Component {

	if resource == nil || len(resource.Manifest.Page) >= numPage {
		return StringInput("ImplementationGuide.Manifest.Page["+strconv.Itoa(numPage)+"].Name", nil)
	}
	return StringInput("ImplementationGuide.Manifest.Page["+strconv.Itoa(numPage)+"].Name", &resource.Manifest.Page[numPage].Name)
}
func (resource *ImplementationGuide) T_ManifestPageTitle(numPage int) templ.Component {

	if resource == nil || len(resource.Manifest.Page) >= numPage {
		return StringInput("ImplementationGuide.Manifest.Page["+strconv.Itoa(numPage)+"].Title", nil)
	}
	return StringInput("ImplementationGuide.Manifest.Page["+strconv.Itoa(numPage)+"].Title", resource.Manifest.Page[numPage].Title)
}
func (resource *ImplementationGuide) T_ManifestPageAnchor(numPage int, numAnchor int) templ.Component {

	if resource == nil || len(resource.Manifest.Page) >= numPage || len(resource.Manifest.Page[numPage].Anchor) >= numAnchor {
		return StringInput("ImplementationGuide.Manifest.Page["+strconv.Itoa(numPage)+"].Anchor["+strconv.Itoa(numAnchor)+"]", nil)
	}
	return StringInput("ImplementationGuide.Manifest.Page["+strconv.Itoa(numPage)+"].Anchor["+strconv.Itoa(numAnchor)+"]", &resource.Manifest.Page[numPage].Anchor[numAnchor])
}
