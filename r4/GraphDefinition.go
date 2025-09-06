package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *GraphDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Id", nil)
	}
	return StringInput("GraphDefinition.Id", resource.Id)
}
func (resource *GraphDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.ImplicitRules", nil)
	}
	return StringInput("GraphDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *GraphDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("GraphDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("GraphDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *GraphDefinition) T_Url() templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Url", nil)
	}
	return StringInput("GraphDefinition.Url", resource.Url)
}
func (resource *GraphDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Version", nil)
	}
	return StringInput("GraphDefinition.Version", resource.Version)
}
func (resource *GraphDefinition) T_Name() templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Name", nil)
	}
	return StringInput("GraphDefinition.Name", &resource.Name)
}
func (resource *GraphDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("GraphDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("GraphDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *GraphDefinition) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("GraphDefinition.Experimental", nil)
	}
	return BoolInput("GraphDefinition.Experimental", resource.Experimental)
}
func (resource *GraphDefinition) T_Date() templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Date", nil)
	}
	return StringInput("GraphDefinition.Date", resource.Date)
}
func (resource *GraphDefinition) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Publisher", nil)
	}
	return StringInput("GraphDefinition.Publisher", resource.Publisher)
}
func (resource *GraphDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Description", nil)
	}
	return StringInput("GraphDefinition.Description", resource.Description)
}
func (resource *GraphDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("GraphDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("GraphDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *GraphDefinition) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Purpose", nil)
	}
	return StringInput("GraphDefinition.Purpose", resource.Purpose)
}
func (resource *GraphDefinition) T_Start() templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil {
		return CodeSelect("GraphDefinition.Start", nil, optionsValueSet)
	}
	return CodeSelect("GraphDefinition.Start", &resource.Start, optionsValueSet)
}
func (resource *GraphDefinition) T_Profile() templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Profile", nil)
	}
	return StringInput("GraphDefinition.Profile", resource.Profile)
}
func (resource *GraphDefinition) T_LinkId(numLink int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Id", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Id", resource.Link[numLink].Id)
}
func (resource *GraphDefinition) T_LinkPath(numLink int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Path", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Path", resource.Link[numLink].Path)
}
func (resource *GraphDefinition) T_LinkSliceName(numLink int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].SliceName", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].SliceName", resource.Link[numLink].SliceName)
}
func (resource *GraphDefinition) T_LinkMin(numLink int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink {
		return IntInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Min", nil)
	}
	return IntInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Min", resource.Link[numLink].Min)
}
func (resource *GraphDefinition) T_LinkMax(numLink int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Max", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Max", resource.Link[numLink].Max)
}
func (resource *GraphDefinition) T_LinkDescription(numLink int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Description", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Description", resource.Link[numLink].Description)
}
func (resource *GraphDefinition) T_LinkTargetId(numLink int, numTarget int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink || len(resource.Link[numLink].Target) >= numTarget {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Id", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Id", resource.Link[numLink].Target[numTarget].Id)
}
func (resource *GraphDefinition) T_LinkTargetType(numLink int, numTarget int) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || len(resource.Link) >= numLink || len(resource.Link[numLink].Target) >= numTarget {
		return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Type", &resource.Link[numLink].Target[numTarget].Type, optionsValueSet)
}
func (resource *GraphDefinition) T_LinkTargetParams(numLink int, numTarget int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink || len(resource.Link[numLink].Target) >= numTarget {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Params", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Params", resource.Link[numLink].Target[numTarget].Params)
}
func (resource *GraphDefinition) T_LinkTargetProfile(numLink int, numTarget int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink || len(resource.Link[numLink].Target) >= numTarget {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Profile", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Profile", resource.Link[numLink].Target[numTarget].Profile)
}
func (resource *GraphDefinition) T_LinkTargetCompartmentId(numLink int, numTarget int, numCompartment int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink || len(resource.Link[numLink].Target) >= numTarget || len(resource.Link[numLink].Target[numTarget].Compartment) >= numCompartment {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Id", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Id", resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Id)
}
func (resource *GraphDefinition) T_LinkTargetCompartmentUse(numLink int, numTarget int, numCompartment int) templ.Component {
	optionsValueSet := VSGraph_compartment_use

	if resource == nil || len(resource.Link) >= numLink || len(resource.Link[numLink].Target) >= numTarget || len(resource.Link[numLink].Target[numTarget].Compartment) >= numCompartment {
		return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Use", nil, optionsValueSet)
	}
	return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Use", &resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Use, optionsValueSet)
}
func (resource *GraphDefinition) T_LinkTargetCompartmentCode(numLink int, numTarget int, numCompartment int) templ.Component {
	optionsValueSet := VSCompartment_type

	if resource == nil || len(resource.Link) >= numLink || len(resource.Link[numLink].Target) >= numTarget || len(resource.Link[numLink].Target[numTarget].Compartment) >= numCompartment {
		return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Code", &resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Code, optionsValueSet)
}
func (resource *GraphDefinition) T_LinkTargetCompartmentRule(numLink int, numTarget int, numCompartment int) templ.Component {
	optionsValueSet := VSGraph_compartment_rule

	if resource == nil || len(resource.Link) >= numLink || len(resource.Link[numLink].Target) >= numTarget || len(resource.Link[numLink].Target[numTarget].Compartment) >= numCompartment {
		return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Rule", nil, optionsValueSet)
	}
	return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Rule", &resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Rule, optionsValueSet)
}
func (resource *GraphDefinition) T_LinkTargetCompartmentExpression(numLink int, numTarget int, numCompartment int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink || len(resource.Link[numLink].Target) >= numTarget || len(resource.Link[numLink].Target[numTarget].Compartment) >= numCompartment {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Expression", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Expression", resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Expression)
}
func (resource *GraphDefinition) T_LinkTargetCompartmentDescription(numLink int, numTarget int, numCompartment int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink || len(resource.Link[numLink].Target) >= numTarget || len(resource.Link[numLink].Target[numTarget].Compartment) >= numCompartment {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Description", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Target["+strconv.Itoa(numTarget)+"].Compartment["+strconv.Itoa(numCompartment)+"].Description", resource.Link[numLink].Target[numTarget].Compartment[numCompartment].Description)
}
