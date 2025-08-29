package r4

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/GraphDefinition
type GraphDefinition struct {
	Id                *string               `json:"id,omitempty"`
	Meta              *Meta                 `json:"meta,omitempty"`
	ImplicitRules     *string               `json:"implicitRules,omitempty"`
	Language          *string               `json:"language,omitempty"`
	Text              *Narrative            `json:"text,omitempty"`
	Contained         []Resource            `json:"contained,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Url               *string               `json:"url,omitempty"`
	Version           *string               `json:"version,omitempty"`
	Name              string                `json:"name"`
	Status            string                `json:"status"`
	Experimental      *bool                 `json:"experimental,omitempty"`
	Date              *string               `json:"date,omitempty"`
	Publisher         *string               `json:"publisher,omitempty"`
	Contact           []ContactDetail       `json:"contact,omitempty"`
	Description       *string               `json:"description,omitempty"`
	UseContext        []UsageContext        `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept     `json:"jurisdiction,omitempty"`
	Purpose           *string               `json:"purpose,omitempty"`
	Start             string                `json:"start"`
	Profile           *string               `json:"profile,omitempty"`
	Link              []GraphDefinitionLink `json:"link,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/GraphDefinition
type GraphDefinitionLink struct {
	Id                *string                     `json:"id,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Path              *string                     `json:"path,omitempty"`
	SliceName         *string                     `json:"sliceName,omitempty"`
	Min               *int                        `json:"min,omitempty"`
	Max               *string                     `json:"max,omitempty"`
	Description       *string                     `json:"description,omitempty"`
	Target            []GraphDefinitionLinkTarget `json:"target,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/GraphDefinition
type GraphDefinitionLinkTarget struct {
	Id                *string                                `json:"id,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Type              string                                 `json:"type"`
	Params            *string                                `json:"params,omitempty"`
	Profile           *string                                `json:"profile,omitempty"`
	Compartment       []GraphDefinitionLinkTargetCompartment `json:"compartment,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/GraphDefinition
type GraphDefinitionLinkTargetCompartment struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Use               string      `json:"use"`
	Code              string      `json:"code"`
	Rule              string      `json:"rule"`
	Expression        *string     `json:"expression,omitempty"`
	Description       *string     `json:"description,omitempty"`
}

type OtherGraphDefinition GraphDefinition

// on convert struct to json, automatically add resourceType=GraphDefinition
func (r GraphDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGraphDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherGraphDefinition: OtherGraphDefinition(r),
		ResourceType:         "GraphDefinition",
	})
}
func (resource *GraphDefinition) GraphDefinitionLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *GraphDefinition) GraphDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *GraphDefinition) GraphDefinitionStart() templ.Component {
	optionsValueSet := VSResource_types
	currentVal := ""
	if resource != nil {
		currentVal = resource.Start
	}
	return CodeSelect("start", currentVal, optionsValueSet)
}
func (resource *GraphDefinition) GraphDefinitionLinkTargetType(numLink int, numTarget int) templ.Component {
	optionsValueSet := VSResource_types
	currentVal := ""
	if resource != nil && len(resource.Link[numLink].Target) >= numTarget {
		currentVal = resource.Link[numLink].Target[numTarget].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
func (resource *GraphDefinition) GraphDefinitionLinkTargetCompartmentUse(numLink int, numTarget int, numCompartment int) templ.Component {
	optionsValueSet := VSGraph_compartment_use
	currentVal := ""
	if resource != nil && len(resource.Link[numLink].Target[numTarget].Compartment) >= numCompartment {
		currentVal = resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Use
	}
	return CodeSelect("use", currentVal, optionsValueSet)
}
func (resource *GraphDefinition) GraphDefinitionLinkTargetCompartmentCode(numLink int, numTarget int, numCompartment int) templ.Component {
	optionsValueSet := VSCompartment_type
	currentVal := ""
	if resource != nil && len(resource.Link[numLink].Target[numTarget].Compartment) >= numCompartment {
		currentVal = resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
func (resource *GraphDefinition) GraphDefinitionLinkTargetCompartmentRule(numLink int, numTarget int, numCompartment int) templ.Component {
	optionsValueSet := VSGraph_compartment_rule
	currentVal := ""
	if resource != nil && len(resource.Link[numLink].Target[numTarget].Compartment) >= numCompartment {
		currentVal = resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Rule
	}
	return CodeSelect("rule", currentVal, optionsValueSet)
}
