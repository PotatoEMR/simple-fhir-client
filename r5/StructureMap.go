package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMap struct {
	Id                     *string                 `json:"id,omitempty"`
	Meta                   *Meta                   `json:"meta,omitempty"`
	ImplicitRules          *string                 `json:"implicitRules,omitempty"`
	Language               *string                 `json:"language,omitempty"`
	Text                   *Narrative              `json:"text,omitempty"`
	Contained              []Resource              `json:"contained,omitempty"`
	Extension              []Extension             `json:"extension,omitempty"`
	ModifierExtension      []Extension             `json:"modifierExtension,omitempty"`
	Url                    string                  `json:"url"`
	Identifier             []Identifier            `json:"identifier,omitempty"`
	Version                *string                 `json:"version,omitempty"`
	VersionAlgorithmString *string                 `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                 `json:"versionAlgorithmCoding"`
	Name                   string                  `json:"name"`
	Title                  *string                 `json:"title,omitempty"`
	Status                 string                  `json:"status"`
	Experimental           *bool                   `json:"experimental,omitempty"`
	Date                   *string                 `json:"date,omitempty"`
	Publisher              *string                 `json:"publisher,omitempty"`
	Contact                []ContactDetail         `json:"contact,omitempty"`
	Description            *string                 `json:"description,omitempty"`
	UseContext             []UsageContext          `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept       `json:"jurisdiction,omitempty"`
	Purpose                *string                 `json:"purpose,omitempty"`
	Copyright              *string                 `json:"copyright,omitempty"`
	CopyrightLabel         *string                 `json:"copyrightLabel,omitempty"`
	Structure              []StructureMapStructure `json:"structure,omitempty"`
	Import                 []string                `json:"import,omitempty"`
	Const                  []StructureMapConst     `json:"const,omitempty"`
	Group                  []StructureMapGroup     `json:"group"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapStructure struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Url               string      `json:"url"`
	Mode              string      `json:"mode"`
	Alias             *string     `json:"alias,omitempty"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapConst struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              *string     `json:"name,omitempty"`
	Value             *string     `json:"value,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapGroup struct {
	Id                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Name              string                   `json:"name"`
	Extends           *string                  `json:"extends,omitempty"`
	TypeMode          *string                  `json:"typeMode,omitempty"`
	Documentation     *string                  `json:"documentation,omitempty"`
	Input             []StructureMapGroupInput `json:"input"`
	Rule              []StructureMapGroupRule  `json:"rule,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapGroupInput struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Type              *string     `json:"type,omitempty"`
	Mode              string      `json:"mode"`
	Documentation     *string     `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapGroupRule struct {
	Id                *string                          `json:"id,omitempty"`
	Extension         []Extension                      `json:"extension,omitempty"`
	ModifierExtension []Extension                      `json:"modifierExtension,omitempty"`
	Name              *string                          `json:"name,omitempty"`
	Source            []StructureMapGroupRuleSource    `json:"source"`
	Target            []StructureMapGroupRuleTarget    `json:"target,omitempty"`
	Dependent         []StructureMapGroupRuleDependent `json:"dependent,omitempty"`
	Documentation     *string                          `json:"documentation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapGroupRuleSource struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Context           string      `json:"context"`
	Min               *int        `json:"min,omitempty"`
	Max               *string     `json:"max,omitempty"`
	Type              *string     `json:"type,omitempty"`
	DefaultValue      *string     `json:"defaultValue,omitempty"`
	Element           *string     `json:"element,omitempty"`
	ListMode          *string     `json:"listMode,omitempty"`
	Variable          *string     `json:"variable,omitempty"`
	Condition         *string     `json:"condition,omitempty"`
	Check             *string     `json:"check,omitempty"`
	LogMessage        *string     `json:"logMessage,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapGroupRuleTarget struct {
	Id                *string                                `json:"id,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Context           *string                                `json:"context,omitempty"`
	Element           *string                                `json:"element,omitempty"`
	Variable          *string                                `json:"variable,omitempty"`
	ListMode          []string                               `json:"listMode,omitempty"`
	ListRuleId        *string                                `json:"listRuleId,omitempty"`
	Transform         *string                                `json:"transform,omitempty"`
	Parameter         []StructureMapGroupRuleTargetParameter `json:"parameter,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapGroupRuleTargetParameter struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ValueId           string      `json:"valueId"`
	ValueString       string      `json:"valueString"`
	ValueBoolean      bool        `json:"valueBoolean"`
	ValueInteger      int         `json:"valueInteger"`
	ValueDecimal      float64     `json:"valueDecimal"`
	ValueDate         string      `json:"valueDate"`
	ValueTime         string      `json:"valueTime"`
	ValueDateTime     string      `json:"valueDateTime"`
}

// http://hl7.org/fhir/r5/StructureDefinition/StructureMap
type StructureMapGroupRuleDependent struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
}

type OtherStructureMap StructureMap

// on convert struct to json, automatically add resourceType=StructureMap
func (r StructureMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherStructureMap
		ResourceType string `json:"resourceType"`
	}{
		OtherStructureMap: OtherStructureMap(r),
		ResourceType:      "StructureMap",
	})
}

func (resource *StructureMap) StructureMapLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *StructureMap) StructureMapStatus() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *StructureMap) StructureMapJurisdiction(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *StructureMap) StructureMapStructureMode(numStructure int) templ.Component {
	optionsValueSet := VSMap_model_mode

	if resource == nil && len(resource.Structure) >= numStructure {
		return CodeSelect("mode", nil, optionsValueSet)
	}
	return CodeSelect("mode", &resource.Structure[numStructure].Mode, optionsValueSet)
}
func (resource *StructureMap) StructureMapGroupTypeMode(numGroup int) templ.Component {
	optionsValueSet := VSMap_group_type_mode

	if resource == nil && len(resource.Group) >= numGroup {
		return CodeSelect("typeMode", nil, optionsValueSet)
	}
	return CodeSelect("typeMode", resource.Group[numGroup].TypeMode, optionsValueSet)
}
func (resource *StructureMap) StructureMapGroupInputMode(numGroup int, numInput int) templ.Component {
	optionsValueSet := VSMap_input_mode

	if resource == nil && len(resource.Group[numGroup].Input) >= numInput {
		return CodeSelect("mode", nil, optionsValueSet)
	}
	return CodeSelect("mode", &resource.Group[numGroup].Input[numInput].Mode, optionsValueSet)
}
func (resource *StructureMap) StructureMapGroupRuleSourceListMode(numGroup int, numRule int, numSource int) templ.Component {
	optionsValueSet := VSMap_source_list_mode

	if resource == nil && len(resource.Group[numGroup].Rule[numRule].Source) >= numSource {
		return CodeSelect("listMode", nil, optionsValueSet)
	}
	return CodeSelect("listMode", resource.Group[numGroup].Rule[numRule].Source[numSource].ListMode, optionsValueSet)
}
func (resource *StructureMap) StructureMapGroupRuleTargetListMode(numGroup int, numRule int, numTarget int) templ.Component {
	optionsValueSet := VSMap_target_list_mode

	if resource == nil && len(resource.Group[numGroup].Rule[numRule].Target) >= numTarget {
		return CodeSelect("listMode", nil, optionsValueSet)
	}
	return CodeSelect("listMode", &resource.Group[numGroup].Rule[numRule].Target[numTarget].ListMode[0], optionsValueSet)
}
func (resource *StructureMap) StructureMapGroupRuleTargetTransform(numGroup int, numRule int, numTarget int) templ.Component {
	optionsValueSet := VSMap_transform

	if resource == nil && len(resource.Group[numGroup].Rule[numRule].Target) >= numTarget {
		return CodeSelect("transform", nil, optionsValueSet)
	}
	return CodeSelect("transform", resource.Group[numGroup].Rule[numRule].Target[numTarget].Transform, optionsValueSet)
}
