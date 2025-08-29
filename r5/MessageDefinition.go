package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	VersionAlgorithmString *string                            `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                            `json:"versionAlgorithmCoding"`
	Name                   *string                            `json:"name,omitempty"`
	Title                  *string                            `json:"title,omitempty"`
	Replaces               []string                           `json:"replaces,omitempty"`
	Status                 string                             `json:"status"`
	Experimental           *bool                              `json:"experimental,omitempty"`
	Date                   string                             `json:"date"`
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

func (resource *MessageDefinition) MessageDefinitionLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *MessageDefinition) MessageDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *MessageDefinition) MessageDefinitionCategory() templ.Component {
	optionsValueSet := VSMessage_significance_category
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Category
	}
	return CodeSelect("category", currentVal, optionsValueSet)
}
func (resource *MessageDefinition) MessageDefinitionResponseRequired() templ.Component {
	optionsValueSet := VSMessageheader_response_request
	currentVal := ""
	if resource != nil {
		currentVal = *resource.ResponseRequired
	}
	return CodeSelect("responseRequired", currentVal, optionsValueSet)
}
func (resource *MessageDefinition) MessageDefinitionFocusCode(numFocus int) templ.Component {
	optionsValueSet := VSResource_types
	currentVal := ""
	if resource != nil && len(resource.Focus) >= numFocus {
		currentVal = resource.Focus[numFocus].Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
