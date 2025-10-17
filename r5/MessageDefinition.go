package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

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
	Date                   FhirDateTime                       `json:"date"`
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

// struct -> json, automatically add resourceType=Patient
func (r MessageDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMessageDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherMessageDefinition: OtherMessageDefinition(r),
		ResourceType:           "MessageDefinition",
	})
}

// json -> struct, first reject if resourceType != MessageDefinition
func (r *MessageDefinition) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "MessageDefinition" {
		return errors.New("resourceType not MessageDefinition")
	}
	return json.Unmarshal(data, (*OtherMessageDefinition)(r))
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
func (resource *MessageDefinition) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *MessageDefinition) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *MessageDefinition) T_VersionAlgorithmString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *MessageDefinition) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *MessageDefinition) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *MessageDefinition) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *MessageDefinition) T_Replaces(numReplaces int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReplaces >= len(resource.Replaces) {
		return StringInput("replaces["+strconv.Itoa(numReplaces)+"]", nil, htmlAttrs)
	}
	return StringInput("replaces["+strconv.Itoa(numReplaces)+"]", &resource.Replaces[numReplaces], htmlAttrs)
}
func (resource *MessageDefinition) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MessageDefinition) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *MessageDefinition) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", &resource.Date, htmlAttrs)
}
func (resource *MessageDefinition) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *MessageDefinition) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *MessageDefinition) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *MessageDefinition) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *MessageDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *MessageDefinition) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *MessageDefinition) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *MessageDefinition) T_CopyrightLabel(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *MessageDefinition) T_Base(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("base", nil, htmlAttrs)
	}
	return StringInput("base", resource.Base, htmlAttrs)
}
func (resource *MessageDefinition) T_Parent(numParent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParent >= len(resource.Parent) {
		return StringInput("parent["+strconv.Itoa(numParent)+"]", nil, htmlAttrs)
	}
	return StringInput("parent["+strconv.Itoa(numParent)+"]", &resource.Parent[numParent], htmlAttrs)
}
func (resource *MessageDefinition) T_EventCoding(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("eventCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("eventCoding", &resource.EventCoding, optionsValueSet, htmlAttrs)
}
func (resource *MessageDefinition) T_EventUri(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("eventUri", nil, htmlAttrs)
	}
	return StringInput("eventUri", &resource.EventUri, htmlAttrs)
}
func (resource *MessageDefinition) T_Category(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMessage_significance_category

	if resource == nil {
		return CodeSelect("category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("category", resource.Category, optionsValueSet, htmlAttrs)
}
func (resource *MessageDefinition) T_ResponseRequired(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMessageheader_response_request

	if resource == nil {
		return CodeSelect("responseRequired", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("responseRequired", resource.ResponseRequired, optionsValueSet, htmlAttrs)
}
func (resource *MessageDefinition) T_Graph(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("graph", nil, htmlAttrs)
	}
	return StringInput("graph", resource.Graph, htmlAttrs)
}
func (resource *MessageDefinition) T_FocusCode(numFocus int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numFocus >= len(resource.Focus) {
		return CodeSelect("focus["+strconv.Itoa(numFocus)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("focus["+strconv.Itoa(numFocus)+"].code", &resource.Focus[numFocus].Code, optionsValueSet, htmlAttrs)
}
func (resource *MessageDefinition) T_FocusProfile(numFocus int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFocus >= len(resource.Focus) {
		return StringInput("focus["+strconv.Itoa(numFocus)+"].profile", nil, htmlAttrs)
	}
	return StringInput("focus["+strconv.Itoa(numFocus)+"].profile", resource.Focus[numFocus].Profile, htmlAttrs)
}
func (resource *MessageDefinition) T_FocusMin(numFocus int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFocus >= len(resource.Focus) {
		return IntInput("focus["+strconv.Itoa(numFocus)+"].min", nil, htmlAttrs)
	}
	return IntInput("focus["+strconv.Itoa(numFocus)+"].min", &resource.Focus[numFocus].Min, htmlAttrs)
}
func (resource *MessageDefinition) T_FocusMax(numFocus int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numFocus >= len(resource.Focus) {
		return StringInput("focus["+strconv.Itoa(numFocus)+"].max", nil, htmlAttrs)
	}
	return StringInput("focus["+strconv.Itoa(numFocus)+"].max", resource.Focus[numFocus].Max, htmlAttrs)
}
func (resource *MessageDefinition) T_AllowedResponseMessage(numAllowedResponse int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllowedResponse >= len(resource.AllowedResponse) {
		return StringInput("allowedResponse["+strconv.Itoa(numAllowedResponse)+"].message", nil, htmlAttrs)
	}
	return StringInput("allowedResponse["+strconv.Itoa(numAllowedResponse)+"].message", &resource.AllowedResponse[numAllowedResponse].Message, htmlAttrs)
}
func (resource *MessageDefinition) T_AllowedResponseSituation(numAllowedResponse int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAllowedResponse >= len(resource.AllowedResponse) {
		return StringInput("allowedResponse["+strconv.Itoa(numAllowedResponse)+"].situation", nil, htmlAttrs)
	}
	return StringInput("allowedResponse["+strconv.Itoa(numAllowedResponse)+"].situation", resource.AllowedResponse[numAllowedResponse].Situation, htmlAttrs)
}
