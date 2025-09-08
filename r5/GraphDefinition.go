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
	Date                   *time.Time            `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
		return StringInput("GraphDefinition.Url", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Url", resource.Url, htmlAttrs)
}
func (resource *GraphDefinition) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Version", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Version", resource.Version, htmlAttrs)
}
func (resource *GraphDefinition) T_VersionAlgorithmString(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *GraphDefinition) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodingSelect("GraphDefinition.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("GraphDefinition.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Name", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Name", &resource.Name, htmlAttrs)
}
func (resource *GraphDefinition) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Title", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Title", resource.Title, htmlAttrs)
}
func (resource *GraphDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("GraphDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("GraphDefinition.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_Experimental(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("GraphDefinition.Experimental", nil, htmlAttrs)
	}
	return BoolInput("GraphDefinition.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *GraphDefinition) T_Date(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("GraphDefinition.Date", nil, htmlAttrs)
	}
	return DateTimeInput("GraphDefinition.Date", resource.Date, htmlAttrs)
}
func (resource *GraphDefinition) T_Publisher(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Publisher", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *GraphDefinition) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *GraphDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("GraphDefinition.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("GraphDefinition.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_Purpose(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Purpose", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *GraphDefinition) T_Copyright(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Copyright", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *GraphDefinition) T_CopyrightLabel(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *GraphDefinition) T_Start(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("GraphDefinition.Start", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Start", resource.Start, htmlAttrs)
}
func (resource *GraphDefinition) T_NodeNodeId(numNode int, htmlAttrs string) templ.Component {

	if resource == nil || numNode >= len(resource.Node) {
		return StringInput("GraphDefinition.Node."+strconv.Itoa(numNode)+"..NodeId", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Node."+strconv.Itoa(numNode)+"..NodeId", &resource.Node[numNode].NodeId, htmlAttrs)
}
func (resource *GraphDefinition) T_NodeDescription(numNode int, htmlAttrs string) templ.Component {

	if resource == nil || numNode >= len(resource.Node) {
		return StringInput("GraphDefinition.Node."+strconv.Itoa(numNode)+"..Description", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Node."+strconv.Itoa(numNode)+"..Description", resource.Node[numNode].Description, htmlAttrs)
}
func (resource *GraphDefinition) T_NodeType(numNode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numNode >= len(resource.Node) {
		return CodeSelect("GraphDefinition.Node."+strconv.Itoa(numNode)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("GraphDefinition.Node."+strconv.Itoa(numNode)+"..Type", &resource.Node[numNode].Type, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_NodeProfile(numNode int, htmlAttrs string) templ.Component {

	if resource == nil || numNode >= len(resource.Node) {
		return StringInput("GraphDefinition.Node."+strconv.Itoa(numNode)+"..Profile", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Node."+strconv.Itoa(numNode)+"..Profile", resource.Node[numNode].Profile, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkDescription(numLink int, htmlAttrs string) templ.Component {

	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Description", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Description", resource.Link[numLink].Description, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkMin(numLink int, htmlAttrs string) templ.Component {

	if resource == nil || numLink >= len(resource.Link) {
		return IntInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Min", nil, htmlAttrs)
	}
	return IntInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Min", resource.Link[numLink].Min, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkMax(numLink int, htmlAttrs string) templ.Component {

	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Max", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Max", resource.Link[numLink].Max, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkSourceId(numLink int, htmlAttrs string) templ.Component {

	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..SourceId", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..SourceId", &resource.Link[numLink].SourceId, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkPath(numLink int, htmlAttrs string) templ.Component {

	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Path", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Path", resource.Link[numLink].Path, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkSliceName(numLink int, htmlAttrs string) templ.Component {

	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..SliceName", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..SliceName", resource.Link[numLink].SliceName, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkTargetId(numLink int, htmlAttrs string) templ.Component {

	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..TargetId", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..TargetId", &resource.Link[numLink].TargetId, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkParams(numLink int, htmlAttrs string) templ.Component {

	if resource == nil || numLink >= len(resource.Link) {
		return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Params", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Params", resource.Link[numLink].Params, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkCompartmentUse(numLink int, numCompartment int, htmlAttrs string) templ.Component {
	optionsValueSet := VSGraph_compartment_use

	if resource == nil || numLink >= len(resource.Link) || numCompartment >= len(resource.Link[numLink].Compartment) {
		return CodeSelect("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Compartment."+strconv.Itoa(numCompartment)+"..Use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Compartment."+strconv.Itoa(numCompartment)+"..Use", &resource.Link[numLink].Compartment[numCompartment].Use, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkCompartmentRule(numLink int, numCompartment int, htmlAttrs string) templ.Component {
	optionsValueSet := VSGraph_compartment_rule

	if resource == nil || numLink >= len(resource.Link) || numCompartment >= len(resource.Link[numLink].Compartment) {
		return CodeSelect("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Compartment."+strconv.Itoa(numCompartment)+"..Rule", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Compartment."+strconv.Itoa(numCompartment)+"..Rule", &resource.Link[numLink].Compartment[numCompartment].Rule, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkCompartmentCode(numLink int, numCompartment int, htmlAttrs string) templ.Component {
	optionsValueSet := VSCompartment_type

	if resource == nil || numLink >= len(resource.Link) || numCompartment >= len(resource.Link[numLink].Compartment) {
		return CodeSelect("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Compartment."+strconv.Itoa(numCompartment)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Compartment."+strconv.Itoa(numCompartment)+"..Code", &resource.Link[numLink].Compartment[numCompartment].Code, optionsValueSet, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkCompartmentExpression(numLink int, numCompartment int, htmlAttrs string) templ.Component {

	if resource == nil || numLink >= len(resource.Link) || numCompartment >= len(resource.Link[numLink].Compartment) {
		return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Compartment."+strconv.Itoa(numCompartment)+"..Expression", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Compartment."+strconv.Itoa(numCompartment)+"..Expression", resource.Link[numLink].Compartment[numCompartment].Expression, htmlAttrs)
}
func (resource *GraphDefinition) T_LinkCompartmentDescription(numLink int, numCompartment int, htmlAttrs string) templ.Component {

	if resource == nil || numLink >= len(resource.Link) || numCompartment >= len(resource.Link[numLink].Compartment) {
		return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Compartment."+strconv.Itoa(numCompartment)+"..Description", nil, htmlAttrs)
	}
	return StringInput("GraphDefinition.Link."+strconv.Itoa(numLink)+"..Compartment."+strconv.Itoa(numCompartment)+"..Description", resource.Link[numLink].Compartment[numCompartment].Description, htmlAttrs)
}
