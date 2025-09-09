package r5

//generated with command go run ./bultaoreune
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
func (r GraphDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "GraphDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "GraphDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *GraphDefinition) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *GraphDefinition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *GraphDefinition) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *GraphDefinition) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", &resource.Name, htmlAttrs)
}
func (resource *GraphDefinition) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *GraphDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *GraphDefinition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *GraphDefinition) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *GraphDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *GraphDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *GraphDefinition) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *GraphDefinition) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *GraphDefinition) T_Start(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("start", nil, htmlAttrs)
	}
	return StringInput("start", resource.Start, htmlAttrs)
}
func (resource *GraphDefinition) T_NodeNodeId(numNode int, htmlAttrs string) templ.Component {
	if resource == nil || numNode >= len(resource.Node) {
		return StringInput("node["+strconv.Itoa(numNode)+"].nodeId", nil, htmlAttrs)
	}
	return StringInput("node["+strconv.Itoa(numNode)+"].nodeId", &resource.Node[numNode].NodeId, htmlAttrs)
}
func (resource *GraphDefinition) T_NodeDescription(numNode int, htmlAttrs string) templ.Component {
	if resource == nil || numNode >= len(resource.Node) {
		return StringInput("node["+strconv.Itoa(numNode)+"].description", nil, htmlAttrs)
	}
	return StringInput("node["+strconv.Itoa(numNode)+"].description", resource.Node[numNode].Description, htmlAttrs)
}
func (resource *GraphDefinition) T_NodeType(numNode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numNode >= len(resource.Node) {
		return CodeSelect("node["+strconv.Itoa(numNode)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("node["+strconv.Itoa(numNode)+"].type", &resource.Node[numNode].Type, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_NodeProfile(numNode int, htmlAttrs string) templ.Component {
	if resource == nil || numNode >= len(resource.Node) {
		return StringInput("node["+strconv.Itoa(numNode)+"].profile", nil, htmlAttrs)
	}
	return StringInput("node["+strconv.Itoa(numNode)+"].profile", resource.Node[numNode].Profile, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkDescription(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("link["+strconv.Itoa(numLink)+"].description", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].description", resource.Link[numLink].Description, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkMin(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return IntInput("link["+strconv.Itoa(numLink)+"].min", nil, htmlAttrs)
	}
	return IntInput("link["+strconv.Itoa(numLink)+"].min", resource.Link[numLink].Min, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkMax(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("link["+strconv.Itoa(numLink)+"].max", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].max", resource.Link[numLink].Max, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkSourceId(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("link["+strconv.Itoa(numLink)+"].sourceId", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].sourceId", &resource.Link[numLink].SourceId, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkPath(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("link["+strconv.Itoa(numLink)+"].path", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].path", resource.Link[numLink].Path, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkSliceName(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("link["+strconv.Itoa(numLink)+"].sliceName", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].sliceName", resource.Link[numLink].SliceName, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetId(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("link["+strconv.Itoa(numLink)+"].targetId", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].targetId", &resource.Link[numLink].TargetId, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkParams(numLink int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("link["+strconv.Itoa(numLink)+"].params", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].params", resource.Link[numLink].Params, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkCompartmentUse(numLink int, numCompartment int, htmlAttrs string) templ.Component {
	optionsValueSet := VSGraph_compartment_use

	if resource == nil || numLink >= len(resource.Link) || numCompartment >= len(resource.Link[numLink].Compartment) {
		return CodeSelect("link["+strconv.Itoa(numLink)+"].compartment["+strconv.Itoa(numCompartment)+"].use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("link["+strconv.Itoa(numLink)+"].compartment["+strconv.Itoa(numCompartment)+"].use", &resource.Link[numLink].Compartment[numCompartment].Use, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkCompartmentRule(numLink int, numCompartment int, htmlAttrs string) templ.Component {
	optionsValueSet := VSGraph_compartment_rule

	if resource == nil || numLink >= len(resource.Link) || numCompartment >= len(resource.Link[numLink].Compartment) {
		return CodeSelect("link["+strconv.Itoa(numLink)+"].compartment["+strconv.Itoa(numCompartment)+"].rule", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("link["+strconv.Itoa(numLink)+"].compartment["+strconv.Itoa(numCompartment)+"].rule", &resource.Link[numLink].Compartment[numCompartment].Rule, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkCompartmentCode(numLink int, numCompartment int, htmlAttrs string) templ.Component {
	optionsValueSet := VSCompartment_type

	if resource == nil || numLink >= len(resource.Link) || numCompartment >= len(resource.Link[numLink].Compartment) {
		return CodeSelect("link["+strconv.Itoa(numLink)+"].compartment["+strconv.Itoa(numCompartment)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("link["+strconv.Itoa(numLink)+"].compartment["+strconv.Itoa(numCompartment)+"].code", &resource.Link[numLink].Compartment[numCompartment].Code, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkCompartmentExpression(numLink int, numCompartment int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) || numCompartment >= len(resource.Link[numLink].Compartment) {
		return StringInput("link["+strconv.Itoa(numLink)+"].compartment["+strconv.Itoa(numCompartment)+"].expression", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].compartment["+strconv.Itoa(numCompartment)+"].expression", resource.Link[numLink].Compartment[numCompartment].Expression, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkCompartmentDescription(numLink int, numCompartment int, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) || numCompartment >= len(resource.Link[numLink].Compartment) {
		return StringInput("link["+strconv.Itoa(numLink)+"].compartment["+strconv.Itoa(numCompartment)+"].description", nil, htmlAttrs)
	}
	return StringInput("link["+strconv.Itoa(numLink)+"].compartment["+strconv.Itoa(numCompartment)+"].description", resource.Link[numLink].Compartment[numCompartment].Description, htmlAttrs)
}
