package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/ActorDefinition
type ActorDefinition struct {
	Id                     *string           `json:"id,omitempty"`
	Meta                   *Meta             `json:"meta,omitempty"`
	ImplicitRules          *string           `json:"implicitRules,omitempty"`
	Language               *string           `json:"language,omitempty"`
	Text                   *Narrative        `json:"text,omitempty"`
	Contained              []Resource        `json:"contained,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Url                    *string           `json:"url,omitempty"`
	Identifier             []Identifier      `json:"identifier,omitempty"`
	Version                *string           `json:"version,omitempty"`
	VersionAlgorithmString *string           `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding           `json:"versionAlgorithmCoding"`
	Name                   *string           `json:"name,omitempty"`
	Title                  *string           `json:"title,omitempty"`
	Status                 string            `json:"status"`
	Experimental           *bool             `json:"experimental,omitempty"`
	Date                   *string           `json:"date,omitempty"`
	Publisher              *string           `json:"publisher,omitempty"`
	Contact                []ContactDetail   `json:"contact,omitempty"`
	Description            *string           `json:"description,omitempty"`
	UseContext             []UsageContext    `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept `json:"jurisdiction,omitempty"`
	Purpose                *string           `json:"purpose,omitempty"`
	Copyright              *string           `json:"copyright,omitempty"`
	CopyrightLabel         *string           `json:"copyrightLabel,omitempty"`
	Type                   string            `json:"type"`
	Documentation          *string           `json:"documentation,omitempty"`
	Reference              []string          `json:"reference,omitempty"`
	Capabilities           *string           `json:"capabilities,omitempty"`
	DerivedFrom            []string          `json:"derivedFrom,omitempty"`
}

type OtherActorDefinition ActorDefinition

// on convert struct to json, automatically add resourceType=ActorDefinition
func (r ActorDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherActorDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherActorDefinition: OtherActorDefinition(r),
		ResourceType:         "ActorDefinition",
	})
}
func (resource *ActorDefinition) ActorDefinitionLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *ActorDefinition) ActorDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *ActorDefinition) ActorDefinitionType() templ.Component {
	optionsValueSet := VSExamplescenario_actor_type
	currentVal := ""
	if resource != nil {
		currentVal = resource.Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
