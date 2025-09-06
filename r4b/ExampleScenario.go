package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *ExampleScenario) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ExampleScenario.Id", nil)
	}
	return StringInput("ExampleScenario.Id", resource.Id)
}
func (resource *ExampleScenario) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ExampleScenario.ImplicitRules", nil)
	}
	return StringInput("ExampleScenario.ImplicitRules", resource.ImplicitRules)
}
func (resource *ExampleScenario) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ExampleScenario.Language", nil, optionsValueSet)
	}
	return CodeSelect("ExampleScenario.Language", resource.Language, optionsValueSet)
}
func (resource *ExampleScenario) T_Url() templ.Component {

	if resource == nil {
		return StringInput("ExampleScenario.Url", nil)
	}
	return StringInput("ExampleScenario.Url", resource.Url)
}
func (resource *ExampleScenario) T_Version() templ.Component {

	if resource == nil {
		return StringInput("ExampleScenario.Version", nil)
	}
	return StringInput("ExampleScenario.Version", resource.Version)
}
func (resource *ExampleScenario) T_Name() templ.Component {

	if resource == nil {
		return StringInput("ExampleScenario.Name", nil)
	}
	return StringInput("ExampleScenario.Name", resource.Name)
}
func (resource *ExampleScenario) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ExampleScenario.Status", nil, optionsValueSet)
	}
	return CodeSelect("ExampleScenario.Status", &resource.Status, optionsValueSet)
}
func (resource *ExampleScenario) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("ExampleScenario.Experimental", nil)
	}
	return BoolInput("ExampleScenario.Experimental", resource.Experimental)
}
func (resource *ExampleScenario) T_Date() templ.Component {

	if resource == nil {
		return StringInput("ExampleScenario.Date", nil)
	}
	return StringInput("ExampleScenario.Date", resource.Date)
}
func (resource *ExampleScenario) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("ExampleScenario.Publisher", nil)
	}
	return StringInput("ExampleScenario.Publisher", resource.Publisher)
}
func (resource *ExampleScenario) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("ExampleScenario.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExampleScenario.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *ExampleScenario) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("ExampleScenario.Copyright", nil)
	}
	return StringInput("ExampleScenario.Copyright", resource.Copyright)
}
func (resource *ExampleScenario) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("ExampleScenario.Purpose", nil)
	}
	return StringInput("ExampleScenario.Purpose", resource.Purpose)
}
func (resource *ExampleScenario) T_Workflow(numWorkflow int) templ.Component {

	if resource == nil || len(resource.Workflow) >= numWorkflow {
		return StringInput("ExampleScenario.Workflow["+strconv.Itoa(numWorkflow)+"]", nil)
	}
	return StringInput("ExampleScenario.Workflow["+strconv.Itoa(numWorkflow)+"]", &resource.Workflow[numWorkflow])
}
func (resource *ExampleScenario) T_ActorId(numActor int) templ.Component {

	if resource == nil || len(resource.Actor) >= numActor {
		return StringInput("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].Id", nil)
	}
	return StringInput("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].Id", resource.Actor[numActor].Id)
}
func (resource *ExampleScenario) T_ActorActorId(numActor int) templ.Component {

	if resource == nil || len(resource.Actor) >= numActor {
		return StringInput("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].ActorId", nil)
	}
	return StringInput("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].ActorId", &resource.Actor[numActor].ActorId)
}
func (resource *ExampleScenario) T_ActorType(numActor int) templ.Component {
	optionsValueSet := VSExamplescenario_actor_type

	if resource == nil || len(resource.Actor) >= numActor {
		return CodeSelect("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].Type", &resource.Actor[numActor].Type, optionsValueSet)
}
func (resource *ExampleScenario) T_ActorName(numActor int) templ.Component {

	if resource == nil || len(resource.Actor) >= numActor {
		return StringInput("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].Name", nil)
	}
	return StringInput("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].Name", resource.Actor[numActor].Name)
}
func (resource *ExampleScenario) T_ActorDescription(numActor int) templ.Component {

	if resource == nil || len(resource.Actor) >= numActor {
		return StringInput("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].Description", nil)
	}
	return StringInput("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].Description", resource.Actor[numActor].Description)
}
func (resource *ExampleScenario) T_InstanceId(numInstance int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Id", nil)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Id", resource.Instance[numInstance].Id)
}
func (resource *ExampleScenario) T_InstanceResourceId(numInstance int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ResourceId", nil)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ResourceId", &resource.Instance[numInstance].ResourceId)
}
func (resource *ExampleScenario) T_InstanceResourceType(numInstance int) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || len(resource.Instance) >= numInstance {
		return CodeSelect("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ResourceType", nil, optionsValueSet)
	}
	return CodeSelect("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ResourceType", &resource.Instance[numInstance].ResourceType, optionsValueSet)
}
func (resource *ExampleScenario) T_InstanceName(numInstance int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Name", nil)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Name", resource.Instance[numInstance].Name)
}
func (resource *ExampleScenario) T_InstanceDescription(numInstance int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Description", nil)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Description", resource.Instance[numInstance].Description)
}
func (resource *ExampleScenario) T_InstanceVersionId(numInstance int, numVersion int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance || len(resource.Instance[numInstance].Version) >= numVersion {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Version["+strconv.Itoa(numVersion)+"].Id", nil)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Version["+strconv.Itoa(numVersion)+"].Id", resource.Instance[numInstance].Version[numVersion].Id)
}
func (resource *ExampleScenario) T_InstanceVersionVersionId(numInstance int, numVersion int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance || len(resource.Instance[numInstance].Version) >= numVersion {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Version["+strconv.Itoa(numVersion)+"].VersionId", nil)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Version["+strconv.Itoa(numVersion)+"].VersionId", &resource.Instance[numInstance].Version[numVersion].VersionId)
}
func (resource *ExampleScenario) T_InstanceVersionDescription(numInstance int, numVersion int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance || len(resource.Instance[numInstance].Version) >= numVersion {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Version["+strconv.Itoa(numVersion)+"].Description", nil)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Version["+strconv.Itoa(numVersion)+"].Description", &resource.Instance[numInstance].Version[numVersion].Description)
}
func (resource *ExampleScenario) T_InstanceContainedInstanceId(numInstance int, numContainedInstance int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance || len(resource.Instance[numInstance].ContainedInstance) >= numContainedInstance {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ContainedInstance["+strconv.Itoa(numContainedInstance)+"].Id", nil)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ContainedInstance["+strconv.Itoa(numContainedInstance)+"].Id", resource.Instance[numInstance].ContainedInstance[numContainedInstance].Id)
}
func (resource *ExampleScenario) T_InstanceContainedInstanceResourceId(numInstance int, numContainedInstance int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance || len(resource.Instance[numInstance].ContainedInstance) >= numContainedInstance {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ContainedInstance["+strconv.Itoa(numContainedInstance)+"].ResourceId", nil)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ContainedInstance["+strconv.Itoa(numContainedInstance)+"].ResourceId", &resource.Instance[numInstance].ContainedInstance[numContainedInstance].ResourceId)
}
func (resource *ExampleScenario) T_InstanceContainedInstanceVersionId(numInstance int, numContainedInstance int) templ.Component {

	if resource == nil || len(resource.Instance) >= numInstance || len(resource.Instance[numInstance].ContainedInstance) >= numContainedInstance {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ContainedInstance["+strconv.Itoa(numContainedInstance)+"].VersionId", nil)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ContainedInstance["+strconv.Itoa(numContainedInstance)+"].VersionId", resource.Instance[numInstance].ContainedInstance[numContainedInstance].VersionId)
}
func (resource *ExampleScenario) T_ProcessId(numProcess int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Id", nil)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Id", resource.Process[numProcess].Id)
}
func (resource *ExampleScenario) T_ProcessTitle(numProcess int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Title", nil)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Title", &resource.Process[numProcess].Title)
}
func (resource *ExampleScenario) T_ProcessDescription(numProcess int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Description", nil)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Description", resource.Process[numProcess].Description)
}
func (resource *ExampleScenario) T_ProcessPreConditions(numProcess int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].PreConditions", nil)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].PreConditions", resource.Process[numProcess].PreConditions)
}
func (resource *ExampleScenario) T_ProcessPostConditions(numProcess int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].PostConditions", nil)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].PostConditions", resource.Process[numProcess].PostConditions)
}
func (resource *ExampleScenario) T_ProcessStepId(numProcess int, numStep int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess || len(resource.Process[numProcess].Step) >= numStep {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Id", nil)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Id", resource.Process[numProcess].Step[numStep].Id)
}
func (resource *ExampleScenario) T_ProcessStepPause(numProcess int, numStep int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess || len(resource.Process[numProcess].Step) >= numStep {
		return BoolInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Pause", nil)
	}
	return BoolInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Pause", resource.Process[numProcess].Step[numStep].Pause)
}
func (resource *ExampleScenario) T_ProcessStepOperationId(numProcess int, numStep int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess || len(resource.Process[numProcess].Step) >= numStep {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Id", nil)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Id", resource.Process[numProcess].Step[numStep].Operation.Id)
}
func (resource *ExampleScenario) T_ProcessStepOperationNumber(numProcess int, numStep int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess || len(resource.Process[numProcess].Step) >= numStep {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Number", nil)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Number", &resource.Process[numProcess].Step[numStep].Operation.Number)
}
func (resource *ExampleScenario) T_ProcessStepOperationType(numProcess int, numStep int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess || len(resource.Process[numProcess].Step) >= numStep {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Type", nil)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Type", resource.Process[numProcess].Step[numStep].Operation.Type)
}
func (resource *ExampleScenario) T_ProcessStepOperationName(numProcess int, numStep int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess || len(resource.Process[numProcess].Step) >= numStep {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Name", nil)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Name", resource.Process[numProcess].Step[numStep].Operation.Name)
}
func (resource *ExampleScenario) T_ProcessStepOperationInitiator(numProcess int, numStep int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess || len(resource.Process[numProcess].Step) >= numStep {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Initiator", nil)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Initiator", resource.Process[numProcess].Step[numStep].Operation.Initiator)
}
func (resource *ExampleScenario) T_ProcessStepOperationReceiver(numProcess int, numStep int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess || len(resource.Process[numProcess].Step) >= numStep {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Receiver", nil)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Receiver", resource.Process[numProcess].Step[numStep].Operation.Receiver)
}
func (resource *ExampleScenario) T_ProcessStepOperationDescription(numProcess int, numStep int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess || len(resource.Process[numProcess].Step) >= numStep {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Description", nil)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Description", resource.Process[numProcess].Step[numStep].Operation.Description)
}
func (resource *ExampleScenario) T_ProcessStepOperationInitiatorActive(numProcess int, numStep int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess || len(resource.Process[numProcess].Step) >= numStep {
		return BoolInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.InitiatorActive", nil)
	}
	return BoolInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.InitiatorActive", resource.Process[numProcess].Step[numStep].Operation.InitiatorActive)
}
func (resource *ExampleScenario) T_ProcessStepOperationReceiverActive(numProcess int, numStep int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess || len(resource.Process[numProcess].Step) >= numStep {
		return BoolInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.ReceiverActive", nil)
	}
	return BoolInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.ReceiverActive", resource.Process[numProcess].Step[numStep].Operation.ReceiverActive)
}
func (resource *ExampleScenario) T_ProcessStepAlternativeId(numProcess int, numStep int, numAlternative int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess || len(resource.Process[numProcess].Step) >= numStep || len(resource.Process[numProcess].Step[numStep].Alternative) >= numAlternative {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Alternative["+strconv.Itoa(numAlternative)+"].Id", nil)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Alternative["+strconv.Itoa(numAlternative)+"].Id", resource.Process[numProcess].Step[numStep].Alternative[numAlternative].Id)
}
func (resource *ExampleScenario) T_ProcessStepAlternativeTitle(numProcess int, numStep int, numAlternative int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess || len(resource.Process[numProcess].Step) >= numStep || len(resource.Process[numProcess].Step[numStep].Alternative) >= numAlternative {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Alternative["+strconv.Itoa(numAlternative)+"].Title", nil)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Alternative["+strconv.Itoa(numAlternative)+"].Title", &resource.Process[numProcess].Step[numStep].Alternative[numAlternative].Title)
}
func (resource *ExampleScenario) T_ProcessStepAlternativeDescription(numProcess int, numStep int, numAlternative int) templ.Component {

	if resource == nil || len(resource.Process) >= numProcess || len(resource.Process[numProcess].Step) >= numStep || len(resource.Process[numProcess].Step[numStep].Alternative) >= numAlternative {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Alternative["+strconv.Itoa(numAlternative)+"].Description", nil)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Alternative["+strconv.Itoa(numAlternative)+"].Description", resource.Process[numProcess].Step[numStep].Alternative[numAlternative].Description)
}
