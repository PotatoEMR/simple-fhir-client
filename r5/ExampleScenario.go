//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/ExampleScenario
type ExampleScenario struct {
	Id                     *string                   `json:"id,omitempty"`
	Meta                   *Meta                     `json:"meta,omitempty"`
	ImplicitRules          *string                   `json:"implicitRules,omitempty"`
	Language               *string                   `json:"language,omitempty"`
	Text                   *Narrative                `json:"text,omitempty"`
	Contained              []Resource                `json:"contained,omitempty"`
	Extension              []Extension               `json:"extension,omitempty"`
	ModifierExtension      []Extension               `json:"modifierExtension,omitempty"`
	Url                    *string                   `json:"url,omitempty"`
	Identifier             []Identifier              `json:"identifier,omitempty"`
	Version                *string                   `json:"version,omitempty"`
	VersionAlgorithmString *string                   `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                   `json:"versionAlgorithmCoding"`
	Name                   *string                   `json:"name,omitempty"`
	Title                  *string                   `json:"title,omitempty"`
	Status                 string                    `json:"status"`
	Experimental           *bool                     `json:"experimental,omitempty"`
	Date                   *string                   `json:"date,omitempty"`
	Publisher              *string                   `json:"publisher,omitempty"`
	Contact                []ContactDetail           `json:"contact,omitempty"`
	Description            *string                   `json:"description,omitempty"`
	UseContext             []UsageContext            `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept         `json:"jurisdiction,omitempty"`
	Purpose                *string                   `json:"purpose,omitempty"`
	Copyright              *string                   `json:"copyright,omitempty"`
	CopyrightLabel         *string                   `json:"copyrightLabel,omitempty"`
	Actor                  []ExampleScenarioActor    `json:"actor,omitempty"`
	Instance               []ExampleScenarioInstance `json:"instance,omitempty"`
	Process                []ExampleScenarioProcess  `json:"process,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExampleScenario
type ExampleScenarioActor struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Key               string      `json:"key"`
	Type              string      `json:"type"`
	Title             string      `json:"title"`
	Description       *string     `json:"description,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExampleScenario
type ExampleScenarioInstance struct {
	Id                        *string                                    `json:"id,omitempty"`
	Extension                 []Extension                                `json:"extension,omitempty"`
	ModifierExtension         []Extension                                `json:"modifierExtension,omitempty"`
	Key                       string                                     `json:"key"`
	StructureType             Coding                                     `json:"structureType"`
	StructureVersion          *string                                    `json:"structureVersion,omitempty"`
	StructureProfileCanonical *string                                    `json:"structureProfileCanonical"`
	StructureProfileUri       *string                                    `json:"structureProfileUri"`
	Title                     string                                     `json:"title"`
	Description               *string                                    `json:"description,omitempty"`
	Content                   *Reference                                 `json:"content,omitempty"`
	Version                   []ExampleScenarioInstanceVersion           `json:"version,omitempty"`
	ContainedInstance         []ExampleScenarioInstanceContainedInstance `json:"containedInstance,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExampleScenario
type ExampleScenarioInstanceVersion struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Key               string      `json:"key"`
	Title             string      `json:"title"`
	Description       *string     `json:"description,omitempty"`
	Content           *Reference  `json:"content,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExampleScenario
type ExampleScenarioInstanceContainedInstance struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	InstanceReference string      `json:"instanceReference"`
	VersionReference  *string     `json:"versionReference,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExampleScenario
type ExampleScenarioProcess struct {
	Id                *string                      `json:"id,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Title             string                       `json:"title"`
	Description       *string                      `json:"description,omitempty"`
	PreConditions     *string                      `json:"preConditions,omitempty"`
	PostConditions    *string                      `json:"postConditions,omitempty"`
	Step              []ExampleScenarioProcessStep `json:"step,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExampleScenario
type ExampleScenarioProcessStep struct {
	Id                *string                                 `json:"id,omitempty"`
	Extension         []Extension                             `json:"extension,omitempty"`
	ModifierExtension []Extension                             `json:"modifierExtension,omitempty"`
	Number            *string                                 `json:"number,omitempty"`
	Workflow          *string                                 `json:"workflow,omitempty"`
	Operation         *ExampleScenarioProcessStepOperation    `json:"operation,omitempty"`
	Alternative       []ExampleScenarioProcessStepAlternative `json:"alternative,omitempty"`
	Pause             *bool                                   `json:"pause,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExampleScenario
type ExampleScenarioProcessStepOperation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              *Coding     `json:"type,omitempty"`
	Title             string      `json:"title"`
	Initiator         *string     `json:"initiator,omitempty"`
	Receiver          *string     `json:"receiver,omitempty"`
	Description       *string     `json:"description,omitempty"`
	InitiatorActive   *bool       `json:"initiatorActive,omitempty"`
	ReceiverActive    *bool       `json:"receiverActive,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExampleScenario
type ExampleScenarioProcessStepAlternative struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Title             string      `json:"title"`
	Description       *string     `json:"description,omitempty"`
}

type OtherExampleScenario ExampleScenario

// on convert struct to json, automatically add resourceType=ExampleScenario
func (r ExampleScenario) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherExampleScenario
		ResourceType string `json:"resourceType"`
	}{
		OtherExampleScenario: OtherExampleScenario(r),
		ResourceType:         "ExampleScenario",
	})
}
