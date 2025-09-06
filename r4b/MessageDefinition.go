package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/MessageDefinition
type MessageDefinition struct {
	Id                *string                            `json:"id,omitempty"`
	Meta              *Meta                              `json:"meta,omitempty"`
	ImplicitRules     *string                            `json:"implicitRules,omitempty"`
	Language          *string                            `json:"language,omitempty"`
	Text              *Narrative                         `json:"text,omitempty"`
	Contained         []Resource                         `json:"contained,omitempty"`
	Extension         []Extension                        `json:"extension,omitempty"`
	ModifierExtension []Extension                        `json:"modifierExtension,omitempty"`
	Url               *string                            `json:"url,omitempty"`
	Identifier        []Identifier                       `json:"identifier,omitempty"`
	Version           *string                            `json:"version,omitempty"`
	Name              *string                            `json:"name,omitempty"`
	Title             *string                            `json:"title,omitempty"`
	Replaces          []string                           `json:"replaces,omitempty"`
	Status            string                             `json:"status"`
	Experimental      *bool                              `json:"experimental,omitempty"`
	Date              string                             `json:"date"`
	Publisher         *string                            `json:"publisher,omitempty"`
	Contact           []ContactDetail                    `json:"contact,omitempty"`
	Description       *string                            `json:"description,omitempty"`
	UseContext        []UsageContext                     `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                  `json:"jurisdiction,omitempty"`
	Purpose           *string                            `json:"purpose,omitempty"`
	Copyright         *string                            `json:"copyright,omitempty"`
	Base              *string                            `json:"base,omitempty"`
	Parent            []string                           `json:"parent,omitempty"`
	EventCoding       Coding                             `json:"eventCoding"`
	EventUri          string                             `json:"eventUri"`
	Category          *string                            `json:"category,omitempty"`
	Focus             []MessageDefinitionFocus           `json:"focus,omitempty"`
	ResponseRequired  *string                            `json:"responseRequired,omitempty"`
	AllowedResponse   []MessageDefinitionAllowedResponse `json:"allowedResponse,omitempty"`
	Graph             []string                           `json:"graph,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MessageDefinition
type MessageDefinitionFocus struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Profile           *string     `json:"profile,omitempty"`
	Min               int         `json:"min"`
	Max               *string     `json:"max,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MessageDefinition
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

func (resource *MessageDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Id", nil)
	}
	return StringInput("MessageDefinition.Id", resource.Id)
}
func (resource *MessageDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.ImplicitRules", nil)
	}
	return StringInput("MessageDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *MessageDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MessageDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("MessageDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *MessageDefinition) T_Url() templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Url", nil)
	}
	return StringInput("MessageDefinition.Url", resource.Url)
}
func (resource *MessageDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Version", nil)
	}
	return StringInput("MessageDefinition.Version", resource.Version)
}
func (resource *MessageDefinition) T_Name() templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Name", nil)
	}
	return StringInput("MessageDefinition.Name", resource.Name)
}
func (resource *MessageDefinition) T_Title() templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Title", nil)
	}
	return StringInput("MessageDefinition.Title", resource.Title)
}
func (resource *MessageDefinition) T_Replaces(numReplaces int) templ.Component {

	if resource == nil || len(resource.Replaces) >= numReplaces {
		return StringInput("MessageDefinition.Replaces["+strconv.Itoa(numReplaces)+"]", nil)
	}
	return StringInput("MessageDefinition.Replaces["+strconv.Itoa(numReplaces)+"]", &resource.Replaces[numReplaces])
}
func (resource *MessageDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("MessageDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("MessageDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *MessageDefinition) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("MessageDefinition.Experimental", nil)
	}
	return BoolInput("MessageDefinition.Experimental", resource.Experimental)
}
func (resource *MessageDefinition) T_Date() templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Date", nil)
	}
	return StringInput("MessageDefinition.Date", &resource.Date)
}
func (resource *MessageDefinition) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Publisher", nil)
	}
	return StringInput("MessageDefinition.Publisher", resource.Publisher)
}
func (resource *MessageDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Description", nil)
	}
	return StringInput("MessageDefinition.Description", resource.Description)
}
func (resource *MessageDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("MessageDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MessageDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *MessageDefinition) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Purpose", nil)
	}
	return StringInput("MessageDefinition.Purpose", resource.Purpose)
}
func (resource *MessageDefinition) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Copyright", nil)
	}
	return StringInput("MessageDefinition.Copyright", resource.Copyright)
}
func (resource *MessageDefinition) T_Base() templ.Component {

	if resource == nil {
		return StringInput("MessageDefinition.Base", nil)
	}
	return StringInput("MessageDefinition.Base", resource.Base)
}
func (resource *MessageDefinition) T_Parent(numParent int) templ.Component {

	if resource == nil || len(resource.Parent) >= numParent {
		return StringInput("MessageDefinition.Parent["+strconv.Itoa(numParent)+"]", nil)
	}
	return StringInput("MessageDefinition.Parent["+strconv.Itoa(numParent)+"]", &resource.Parent[numParent])
}
func (resource *MessageDefinition) T_Category() templ.Component {
	optionsValueSet := VSMessage_significance_category

	if resource == nil {
		return CodeSelect("MessageDefinition.Category", nil, optionsValueSet)
	}
	return CodeSelect("MessageDefinition.Category", resource.Category, optionsValueSet)
}
func (resource *MessageDefinition) T_ResponseRequired() templ.Component {
	optionsValueSet := VSMessageheader_response_request

	if resource == nil {
		return CodeSelect("MessageDefinition.ResponseRequired", nil, optionsValueSet)
	}
	return CodeSelect("MessageDefinition.ResponseRequired", resource.ResponseRequired, optionsValueSet)
}
func (resource *MessageDefinition) T_Graph(numGraph int) templ.Component {

	if resource == nil || len(resource.Graph) >= numGraph {
		return StringInput("MessageDefinition.Graph["+strconv.Itoa(numGraph)+"]", nil)
	}
	return StringInput("MessageDefinition.Graph["+strconv.Itoa(numGraph)+"]", &resource.Graph[numGraph])
}
func (resource *MessageDefinition) T_FocusId(numFocus int) templ.Component {

	if resource == nil || len(resource.Focus) >= numFocus {
		return StringInput("MessageDefinition.Focus["+strconv.Itoa(numFocus)+"].Id", nil)
	}
	return StringInput("MessageDefinition.Focus["+strconv.Itoa(numFocus)+"].Id", resource.Focus[numFocus].Id)
}
func (resource *MessageDefinition) T_FocusCode(numFocus int) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || len(resource.Focus) >= numFocus {
		return CodeSelect("MessageDefinition.Focus["+strconv.Itoa(numFocus)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("MessageDefinition.Focus["+strconv.Itoa(numFocus)+"].Code", &resource.Focus[numFocus].Code, optionsValueSet)
}
func (resource *MessageDefinition) T_FocusProfile(numFocus int) templ.Component {

	if resource == nil || len(resource.Focus) >= numFocus {
		return StringInput("MessageDefinition.Focus["+strconv.Itoa(numFocus)+"].Profile", nil)
	}
	return StringInput("MessageDefinition.Focus["+strconv.Itoa(numFocus)+"].Profile", resource.Focus[numFocus].Profile)
}
func (resource *MessageDefinition) T_FocusMin(numFocus int) templ.Component {

	if resource == nil || len(resource.Focus) >= numFocus {
		return IntInput("MessageDefinition.Focus["+strconv.Itoa(numFocus)+"].Min", nil)
	}
	return IntInput("MessageDefinition.Focus["+strconv.Itoa(numFocus)+"].Min", &resource.Focus[numFocus].Min)
}
func (resource *MessageDefinition) T_FocusMax(numFocus int) templ.Component {

	if resource == nil || len(resource.Focus) >= numFocus {
		return StringInput("MessageDefinition.Focus["+strconv.Itoa(numFocus)+"].Max", nil)
	}
	return StringInput("MessageDefinition.Focus["+strconv.Itoa(numFocus)+"].Max", resource.Focus[numFocus].Max)
}
func (resource *MessageDefinition) T_AllowedResponseId(numAllowedResponse int) templ.Component {

	if resource == nil || len(resource.AllowedResponse) >= numAllowedResponse {
		return StringInput("MessageDefinition.AllowedResponse["+strconv.Itoa(numAllowedResponse)+"].Id", nil)
	}
	return StringInput("MessageDefinition.AllowedResponse["+strconv.Itoa(numAllowedResponse)+"].Id", resource.AllowedResponse[numAllowedResponse].Id)
}
func (resource *MessageDefinition) T_AllowedResponseMessage(numAllowedResponse int) templ.Component {

	if resource == nil || len(resource.AllowedResponse) >= numAllowedResponse {
		return StringInput("MessageDefinition.AllowedResponse["+strconv.Itoa(numAllowedResponse)+"].Message", nil)
	}
	return StringInput("MessageDefinition.AllowedResponse["+strconv.Itoa(numAllowedResponse)+"].Message", &resource.AllowedResponse[numAllowedResponse].Message)
}
func (resource *MessageDefinition) T_AllowedResponseSituation(numAllowedResponse int) templ.Component {

	if resource == nil || len(resource.AllowedResponse) >= numAllowedResponse {
		return StringInput("MessageDefinition.AllowedResponse["+strconv.Itoa(numAllowedResponse)+"].Situation", nil)
	}
	return StringInput("MessageDefinition.AllowedResponse["+strconv.Itoa(numAllowedResponse)+"].Situation", resource.AllowedResponse[numAllowedResponse].Situation)
}
