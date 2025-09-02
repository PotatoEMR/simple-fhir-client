package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *ExampleScenario) ExampleScenarioLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ExampleScenario) ExampleScenarioStatus() templ.Component {
	optionsValueSet := VSPublication_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ExampleScenario) ExampleScenarioJurisdiction(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *ExampleScenario) ExampleScenarioActorType(numActor int) templ.Component {
	optionsValueSet := VSExamplescenario_actor_type

	if resource != nil && len(resource.Actor) >= numActor {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", &resource.Actor[numActor].Type, optionsValueSet)
}
func (resource *ExampleScenario) ExampleScenarioInstanceStructureType(numInstance int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Instance) >= numInstance {
		return CodingSelect("structureType", nil, optionsValueSet)
	}
	return CodingSelect("structureType", &resource.Instance[numInstance].StructureType, optionsValueSet)
}
func (resource *ExampleScenario) ExampleScenarioProcessStepOperationType(numProcess int, numStep int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Process[numProcess].Step) >= numStep {
		return CodingSelect("type", nil, optionsValueSet)
	}
	return CodingSelect("type", resource.Process[numProcess].Step[numStep].Operation.Type, optionsValueSet)
}
