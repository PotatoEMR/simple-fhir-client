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

// http://hl7.org/fhir/r5/StructureDefinition/MessageDefinition
type MessageDefinition struct {
	Id                     *string                            `json:"id,omitempty"`
	Meta                   *Meta                              `json:"meta,omitempty"`
	ImplicitRules          *string                            `json:"implicitRules,omitempty"`
	Language               *string                            `json:"language,omitempty"`
	Text                   *Narrative                         `json:"text,omitempty"`
	Contained              []Resource                         `json:"contained,omitempty"`
	Extension              []Extension                        `json:"extension,omitempty"`
	ModifierExtension      []Extension                        `json:"modifierExtension,omitempty"`
	Url                    *string                            `json:"url,omitempty"`
	Identifier             []Identifier                       `json:"identifier,omitempty"`
	Version                *string                            `json:"version,omitempty"`
	VersionAlgorithmString *string                            `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                            `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                            `json:"name,omitempty"`
	Title                  *string                            `json:"title,omitempty"`
	Replaces               []string                           `json:"replaces,omitempty"`
	Status                 string                             `json:"status"`
	Experimental           *bool                              `json:"experimental,omitempty"`
	Date                   time.Time                          `json:"date,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher              *string                            `json:"publisher,omitempty"`
	Contact                []ContactDetail                    `json:"contact,omitempty"`
	Description            *string                            `json:"description,omitempty"`
	UseContext             []UsageContext                     `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                  `json:"jurisdiction,omitempty"`
	Purpose                *string                            `json:"purpose,omitempty"`
	Copyright              *string                            `json:"copyright,omitempty"`
	CopyrightLabel         *string                            `json:"copyrightLabel,omitempty"`
	Base                   *string                            `json:"base,omitempty"`
	Parent                 []string                           `json:"parent,omitempty"`
	EventCoding            Coding                             `json:"eventCoding"`
	EventUri               string                             `json:"eventUri"`
	Category               *string                            `json:"category,omitempty"`
	Focus                  []MessageDefinitionFocus           `json:"focus,omitempty"`
	ResponseRequired       *string                            `json:"responseRequired,omitempty"`
	AllowedResponse        []MessageDefinitionAllowedResponse `json:"allowedResponse,omitempty"`
	Graph                  *string                            `json:"graph,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MessageDefinition
type MessageDefinitionFocus struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Profile           *string     `json:"profile,omitempty"`
	Min               int         `json:"min"`
	Max               *string     `json:"max,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MessageDefinition
type MessageDefinitionAllowedResponse struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Message           string      `json:"message"`
	Situation         *string     `json:"situation,omitempty"`
}

type OtherMessageDefinition MessageDefinition

// on convert struct to json, automatically add resourceType=MessageDefinition
func (r MessageDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMessageDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherMessageDefinition: OtherMessageDefinition(r),
		ResourceType:           "MessageDefinition",
	})
}
func (r MessageDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MessageDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "MessageDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MessageDefinition) T_Url(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Url", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.Url", resource.Url, htmlAttrs)
}
func (resource *MessageDefinition) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Version", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.Version", resource.Version, htmlAttrs)
}
func (resource *MessageDefinition) T_VersionAlgorithmString(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *MessageDefinition) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodingSelect("MessageDefinition.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("MessageDefinition.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *MessageDefinition) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Name", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.Name", resource.Name, htmlAttrs)
}
func (resource *MessageDefinition) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Title", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.Title", resource.Title, htmlAttrs)
}
func (resource *MessageDefinition) T_Replaces(numReplaces int, htmlAttrs string) templ.Component {

	if resource == nil || numReplaces >= len(resource.Replaces) {
		return StringInput("MessageDefinition.Replaces."+strconv.Itoa(numReplaces)+".", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.Replaces."+strconv.Itoa(numReplaces)+".", &resource.Replaces[numReplaces], htmlAttrs)
}
func (resource *MessageDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("MessageDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("MessageDefinition.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MessageDefinition) T_Experimental(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("MessageDefinition.Experimental", nil, htmlAttrs)
	}
	return BoolInput("MessageDefinition.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *MessageDefinition) T_Date(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("MessageDefinition.Date", nil, htmlAttrs)
	}
	return DateTimeInput("MessageDefinition.Date", &resource.Date, htmlAttrs)
}
func (resource *MessageDefinition) T_Publisher(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Publisher", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *MessageDefinition) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *MessageDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("MessageDefinition.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MessageDefinition.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *MessageDefinition) T_Purpose(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Purpose", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *MessageDefinition) T_Copyright(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Copyright", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *MessageDefinition) T_CopyrightLabel(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *MessageDefinition) T_Base(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Base", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.Base", resource.Base, htmlAttrs)
}
func (resource *MessageDefinition) T_Parent(numParent int, htmlAttrs string) templ.Component {

	if resource == nil || numParent >= len(resource.Parent) {
		return StringInput("MessageDefinition.Parent."+strconv.Itoa(numParent)+".", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.Parent."+strconv.Itoa(numParent)+".", &resource.Parent[numParent], htmlAttrs)
}
func (resource *MessageDefinition) T_EventCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodingSelect("MessageDefinition.EventCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("MessageDefinition.EventCoding", &resource.EventCoding, optionsValueSet, htmlAttrs)
}
func (resource *MessageDefinition) T_EventUri(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.EventUri", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.EventUri", &resource.EventUri, htmlAttrs)
}
func (resource *MessageDefinition) T_Category(htmlAttrs string) templ.Component {
	optionsValueSet := VSMessage_significance_category

	if resource == nil {
		return CodeSelect("MessageDefinition.Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("MessageDefinition.Category", resource.Category, optionsValueSet, htmlAttrs)
}
func (resource *MessageDefinition) T_ResponseRequired(htmlAttrs string) templ.Component {
	optionsValueSet := VSMessageheader_response_request

	if resource == nil {
		return CodeSelect("MessageDefinition.ResponseRequired", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("MessageDefinition.ResponseRequired", resource.ResponseRequired, optionsValueSet, htmlAttrs)
}
func (resource *MessageDefinition) T_Graph(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Graph", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.Graph", resource.Graph, htmlAttrs)
}
func (resource *MessageDefinition) T_FocusCode(numFocus int, htmlAttrs string) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numFocus >= len(resource.Focus) {
		return CodeSelect("MessageDefinition.Focus."+strconv.Itoa(numFocus)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("MessageDefinition.Focus."+strconv.Itoa(numFocus)+"..Code", &resource.Focus[numFocus].Code, optionsValueSet, htmlAttrs)
}
func (resource *MessageDefinition) T_FocusProfile(numFocus int, htmlAttrs string) templ.Component {

	if resource == nil || numFocus >= len(resource.Focus) {
		return StringInput("MessageDefinition.Focus."+strconv.Itoa(numFocus)+"..Profile", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.Focus."+strconv.Itoa(numFocus)+"..Profile", resource.Focus[numFocus].Profile, htmlAttrs)
}
func (resource *MessageDefinition) T_FocusMin(numFocus int, htmlAttrs string) templ.Component {

	if resource == nil || numFocus >= len(resource.Focus) {
		return IntInput("MessageDefinition.Focus."+strconv.Itoa(numFocus)+"..Min", nil, htmlAttrs)
	}
	return IntInput("MessageDefinition.Focus."+strconv.Itoa(numFocus)+"..Min", &resource.Focus[numFocus].Min, htmlAttrs)
}
func (resource *MessageDefinition) T_FocusMax(numFocus int, htmlAttrs string) templ.Component {

	if resource == nil || numFocus >= len(resource.Focus) {
		return StringInput("MessageDefinition.Focus."+strconv.Itoa(numFocus)+"..Max", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.Focus."+strconv.Itoa(numFocus)+"..Max", resource.Focus[numFocus].Max, htmlAttrs)
}
func (resource *MessageDefinition) T_AllowedResponseMessage(numAllowedResponse int, htmlAttrs string) templ.Component {

	if resource == nil || numAllowedResponse >= len(resource.AllowedResponse) {
		return StringInput("MessageDefinition.AllowedResponse."+strconv.Itoa(numAllowedResponse)+"..Message", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.AllowedResponse."+strconv.Itoa(numAllowedResponse)+"..Message", &resource.AllowedResponse[numAllowedResponse].Message, htmlAttrs)
}
func (resource *MessageDefinition) T_AllowedResponseSituation(numAllowedResponse int, htmlAttrs string) templ.Component {

	if resource == nil || numAllowedResponse >= len(resource.AllowedResponse) {
		return StringInput("MessageDefinition.AllowedResponse."+strconv.Itoa(numAllowedResponse)+"..Situation", nil, htmlAttrs)
	}
	return StringInput("MessageDefinition.AllowedResponse."+strconv.Itoa(numAllowedResponse)+"..Situation", resource.AllowedResponse[numAllowedResponse].Situation, htmlAttrs)
}
