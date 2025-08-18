//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

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
