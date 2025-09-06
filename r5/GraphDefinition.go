package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/GraphDefinition
type GraphDefinition struct {
	Id                     *string               `json:"id,omitempty"`
	Meta                   *Meta                 `json:"meta,omitempty"`
	ImplicitRules          *string               `json:"implicitRules,omitempty"`
	Language               *string               `json:"language,omitempty"`
	Text                   *Narrative            `json:"text,omitempty"`
	Contained              []Resource            `json:"contained,omitempty"`
	Extension              []Extension           `json:"extension,omitempty"`
	ModifierExtension      []Extension           `json:"modifierExtension,omitempty"`
	Url                    *string               `json:"url,omitempty"`
	Identifier             []Identifier          `json:"identifier,omitempty"`
	Version                *string               `json:"version,omitempty"`
	VersionAlgorithmString *string               `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding               `json:"versionAlgorithmCoding,omitempty"`
	Name                   string                `json:"name"`
	Title                  *string               `json:"title,omitempty"`
	Status                 string                `json:"status"`
	Experimental           *bool                 `json:"experimental,omitempty"`
	Date                   *string               `json:"date,omitempty"`
	Publisher              *string               `json:"publisher,omitempty"`
	Contact                []ContactDetail       `json:"contact,omitempty"`
	Description            *string               `json:"description,omitempty"`
	UseContext             []UsageContext        `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept     `json:"jurisdiction,omitempty"`
	Purpose                *string               `json:"purpose,omitempty"`
	Copyright              *string               `json:"copyright,omitempty"`
	CopyrightLabel         *string               `json:"copyrightLabel,omitempty"`
	Start                  *string               `json:"start,omitempty"`
	Node                   []GraphDefinitionNode `json:"node,omitempty"`
	Link                   []GraphDefinitionLink `json:"link,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/GraphDefinition
type GraphDefinitionNode struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	NodeId            string      `json:"nodeId"`
	Description       *string     `json:"description,omitempty"`
	Type              string      `json:"type"`
	Profile           *string     `json:"profile,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/GraphDefinition
type GraphDefinitionLink struct {
	Id                *string                          `json:"id,omitempty"`
	Extension         []Extension                      `json:"extension,omitempty"`
	ModifierExtension []Extension                      `json:"modifierExtension,omitempty"`
	Description       *string                          `json:"description,omitempty"`
	Min               *int                             `json:"min,omitempty"`
	Max               *string                          `json:"max,omitempty"`
	SourceId          string                           `json:"sourceId"`
	Path              *string                          `json:"path,omitempty"`
	SliceName         *string                          `json:"sliceName,omitempty"`
	TargetId          string                           `json:"targetId"`
	Params            *string                          `json:"params,omitempty"`
	Compartment       []GraphDefinitionLinkCompartment `json:"compartment,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/GraphDefinition
type GraphDefinitionLinkCompartment struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Use               string      `json:"use"`
	Rule              string      `json:"rule"`
	Code              string      `json:"code"`
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
func (resource *GraphDefinition) T_Title() templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Title", nil)
	}
	return StringInput("GraphDefinition.Title", resource.Title)
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
func (resource *GraphDefinition) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Copyright", nil)
	}
	return StringInput("GraphDefinition.Copyright", resource.Copyright)
}
func (resource *GraphDefinition) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.CopyrightLabel", nil)
	}
	return StringInput("GraphDefinition.CopyrightLabel", resource.CopyrightLabel)
}
func (resource *GraphDefinition) T_Start() templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Start", nil)
	}
	return StringInput("GraphDefinition.Start", resource.Start)
}
func (resource *GraphDefinition) T_NodeId(numNode int) templ.Component {

	if resource == nil || len(resource.Node) >= numNode {
		return StringInput("GraphDefinition.Node["+strconv.Itoa(numNode)+"].Id", nil)
	}
	return StringInput("GraphDefinition.Node["+strconv.Itoa(numNode)+"].Id", resource.Node[numNode].Id)
}
func (resource *GraphDefinition) T_NodeNodeId(numNode int) templ.Component {

	if resource == nil || len(resource.Node) >= numNode {
		return StringInput("GraphDefinition.Node["+strconv.Itoa(numNode)+"].NodeId", nil)
	}
	return StringInput("GraphDefinition.Node["+strconv.Itoa(numNode)+"].NodeId", &resource.Node[numNode].NodeId)
}
func (resource *GraphDefinition) T_NodeDescription(numNode int) templ.Component {

	if resource == nil || len(resource.Node) >= numNode {
		return StringInput("GraphDefinition.Node["+strconv.Itoa(numNode)+"].Description", nil)
	}
	return StringInput("GraphDefinition.Node["+strconv.Itoa(numNode)+"].Description", resource.Node[numNode].Description)
}
func (resource *GraphDefinition) T_NodeType(numNode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Node) >= numNode {
		return CodeSelect("GraphDefinition.Node["+strconv.Itoa(numNode)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("GraphDefinition.Node["+strconv.Itoa(numNode)+"].Type", &resource.Node[numNode].Type, optionsValueSet)
}
func (resource *GraphDefinition) T_NodeProfile(numNode int) templ.Component {

	if resource == nil || len(resource.Node) >= numNode {
		return StringInput("GraphDefinition.Node["+strconv.Itoa(numNode)+"].Profile", nil)
	}
	return StringInput("GraphDefinition.Node["+strconv.Itoa(numNode)+"].Profile", resource.Node[numNode].Profile)
}
func (resource *GraphDefinition) T_LinkId(numLink int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Id", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Id", resource.Link[numLink].Id)
}
func (resource *GraphDefinition) T_LinkDescription(numLink int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Description", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Description", resource.Link[numLink].Description)
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
func (resource *GraphDefinition) T_LinkSourceId(numLink int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].SourceId", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].SourceId", &resource.Link[numLink].SourceId)
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
func (resource *GraphDefinition) T_LinkTargetId(numLink int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].TargetId", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].TargetId", &resource.Link[numLink].TargetId)
}
func (resource *GraphDefinition) T_LinkParams(numLink int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Params", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Params", resource.Link[numLink].Params)
}
func (resource *GraphDefinition) T_LinkCompartmentId(numLink int, numCompartment int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink || len(resource.Link[numLink].Compartment) >= numCompartment {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Compartment["+strconv.Itoa(numCompartment)+"].Id", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Compartment["+strconv.Itoa(numCompartment)+"].Id", resource.Link[numLink].Compartment[numCompartment].Id)
}
func (resource *GraphDefinition) T_LinkCompartmentUse(numLink int, numCompartment int) templ.Component {
	optionsValueSet := VSGraph_compartment_use

	if resource == nil || len(resource.Link) >= numLink || len(resource.Link[numLink].Compartment) >= numCompartment {
		return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Compartment["+strconv.Itoa(numCompartment)+"].Use", nil, optionsValueSet)
	}
	return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Compartment["+strconv.Itoa(numCompartment)+"].Use", &resource.Link[numLink].Compartment[numCompartment].Use, optionsValueSet)
}
func (resource *GraphDefinition) T_LinkCompartmentRule(numLink int, numCompartment int) templ.Component {
	optionsValueSet := VSGraph_compartment_rule

	if resource == nil || len(resource.Link) >= numLink || len(resource.Link[numLink].Compartment) >= numCompartment {
		return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Compartment["+strconv.Itoa(numCompartment)+"].Rule", nil, optionsValueSet)
	}
	return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Compartment["+strconv.Itoa(numCompartment)+"].Rule", &resource.Link[numLink].Compartment[numCompartment].Rule, optionsValueSet)
}
func (resource *GraphDefinition) T_LinkCompartmentCode(numLink int, numCompartment int) templ.Component {
	optionsValueSet := VSCompartment_type

	if resource == nil || len(resource.Link) >= numLink || len(resource.Link[numLink].Compartment) >= numCompartment {
		return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Compartment["+strconv.Itoa(numCompartment)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Compartment["+strconv.Itoa(numCompartment)+"].Code", &resource.Link[numLink].Compartment[numCompartment].Code, optionsValueSet)
}
func (resource *GraphDefinition) T_LinkCompartmentExpression(numLink int, numCompartment int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink || len(resource.Link[numLink].Compartment) >= numCompartment {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Compartment["+strconv.Itoa(numCompartment)+"].Expression", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Compartment["+strconv.Itoa(numCompartment)+"].Expression", resource.Link[numLink].Compartment[numCompartment].Expression)
}
func (resource *GraphDefinition) T_LinkCompartmentDescription(numLink int, numCompartment int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink || len(resource.Link[numLink].Compartment) >= numCompartment {
		return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Compartment["+strconv.Itoa(numCompartment)+"].Description", nil)
	}
	return StringInput("GraphDefinition.Link["+strconv.Itoa(numLink)+"].Compartment["+strconv.Itoa(numCompartment)+"].Description", resource.Link[numLink].Compartment[numCompartment].Description)
}
