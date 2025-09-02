package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/ExampleScenario
type ExampleScenario struct {
	Id                *string                   `json:"id,omitempty"`
	Meta              *Meta                     `json:"meta,omitempty"`
	ImplicitRules     *string                   `json:"implicitRules,omitempty"`
	Language          *string                   `json:"language,omitempty"`
	Text              *Narrative                `json:"text,omitempty"`
	Contained         []Resource                `json:"contained,omitempty"`
	Extension         []Extension               `json:"extension,omitempty"`
	ModifierExtension []Extension               `json:"modifierExtension,omitempty"`
	Url               *string                   `json:"url,omitempty"`
	Identifier        []Identifier              `json:"identifier,omitempty"`
	Version           *string                   `json:"version,omitempty"`
	Name              *string                   `json:"name,omitempty"`
	Status            string                    `json:"status"`
	Experimental      *bool                     `json:"experimental,omitempty"`
	Date              *string                   `json:"date,omitempty"`
	Publisher         *string                   `json:"publisher,omitempty"`
	Contact           []ContactDetail           `json:"contact,omitempty"`
	UseContext        []UsageContext            `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept         `json:"jurisdiction,omitempty"`
	Copyright         *string                   `json:"copyright,omitempty"`
	Purpose           *string                   `json:"purpose,omitempty"`
	Actor             []ExampleScenarioActor    `json:"actor,omitempty"`
	Instance          []ExampleScenarioInstance `json:"instance,omitempty"`
	Process           []ExampleScenarioProcess  `json:"process,omitempty"`
	Workflow          []string                  `json:"workflow,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExampleScenario
type ExampleScenarioActor struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ActorId           string      `json:"actorId"`
	Type              string      `json:"type"`
	Name              *string     `json:"name,omitempty"`
	Description       *string     `json:"description,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExampleScenario
type ExampleScenarioInstance struct {
	Id                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	ResourceId        string                                     `json:"resourceId"`
	ResourceType      string                                     `json:"resourceType"`
	Name              *string                                    `json:"name,omitempty"`
	Description       *string                                    `json:"description,omitempty"`
	Version           []ExampleScenarioInstanceVersion           `json:"version,omitempty"`
	ContainedInstance []ExampleScenarioInstanceContainedInstance `json:"containedInstance,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExampleScenario
type ExampleScenarioInstanceVersion struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	VersionId         string      `json:"versionId"`
	Description       string      `json:"description"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExampleScenario
type ExampleScenarioInstanceContainedInstance struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ResourceId        string      `json:"resourceId"`
	VersionId         *string     `json:"versionId,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExampleScenario
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

// http://hl7.org/fhir/r4b/StructureDefinition/ExampleScenario
type ExampleScenarioProcessStep struct {
	Id                *string                                 `json:"id,omitempty"`
	Extension         []Extension                             `json:"extension,omitempty"`
	ModifierExtension []Extension                             `json:"modifierExtension,omitempty"`
	Pause             *bool                                   `json:"pause,omitempty"`
	Operation         *ExampleScenarioProcessStepOperation    `json:"operation,omitempty"`
	Alternative       []ExampleScenarioProcessStepAlternative `json:"alternative,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExampleScenario
type ExampleScenarioProcessStepOperation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Number            string      `json:"number"`
	Type              *string     `json:"type,omitempty"`
	Name              *string     `json:"name,omitempty"`
	Initiator         *string     `json:"initiator,omitempty"`
	Receiver          *string     `json:"receiver,omitempty"`
	Description       *string     `json:"description,omitempty"`
	InitiatorActive   *bool       `json:"initiatorActive,omitempty"`
	ReceiverActive    *bool       `json:"receiverActive,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExampleScenario
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

func (resource *ExampleScenario) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ExampleScenario) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ExampleScenario) T_Jurisdiction(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *ExampleScenario) T_ActorType(numActor int) templ.Component {
	optionsValueSet := VSExamplescenario_actor_type

	if resource == nil && len(resource.Actor) >= numActor {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", &resource.Actor[numActor].Type, optionsValueSet)
}
func (resource *ExampleScenario) T_InstanceResourceType(numInstance int) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil && len(resource.Instance) >= numInstance {
		return CodeSelect("resourceType", nil, optionsValueSet)
	}
	return CodeSelect("resourceType", &resource.Instance[numInstance].ResourceType, optionsValueSet)
}
