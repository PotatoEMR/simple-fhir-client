package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Date              *time.Time                `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r ExampleScenario) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ExampleScenario/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ExampleScenario"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ExampleScenario) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ExampleScenario.Url", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Url", resource.Url, htmlAttrs)
}
func (resource *ExampleScenario) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ExampleScenario.Version", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Version", resource.Version, htmlAttrs)
}
func (resource *ExampleScenario) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ExampleScenario.Name", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Name", resource.Name, htmlAttrs)
}
func (resource *ExampleScenario) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ExampleScenario.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ExampleScenario.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("ExampleScenario.Experimental", nil, htmlAttrs)
	}
	return BoolInput("ExampleScenario.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *ExampleScenario) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("ExampleScenario.Date", nil, htmlAttrs)
	}
	return DateTimeInput("ExampleScenario.Date", resource.Date, htmlAttrs)
}
func (resource *ExampleScenario) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ExampleScenario.Publisher", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *ExampleScenario) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("ExampleScenario.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ExampleScenario.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ExampleScenario.Copyright", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *ExampleScenario) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ExampleScenario.Purpose", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *ExampleScenario) T_Workflow(numWorkflow int, htmlAttrs string) templ.Component {
	if resource == nil || numWorkflow >= len(resource.Workflow) {
		return StringInput("ExampleScenario.Workflow["+strconv.Itoa(numWorkflow)+"]", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Workflow["+strconv.Itoa(numWorkflow)+"]", &resource.Workflow[numWorkflow], htmlAttrs)
}
func (resource *ExampleScenario) T_ActorActorId(numActor int, htmlAttrs string) templ.Component {
	if resource == nil || numActor >= len(resource.Actor) {
		return StringInput("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].ActorId", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].ActorId", &resource.Actor[numActor].ActorId, htmlAttrs)
}
func (resource *ExampleScenario) T_ActorType(numActor int, htmlAttrs string) templ.Component {
	optionsValueSet := VSExamplescenario_actor_type

	if resource == nil || numActor >= len(resource.Actor) {
		return CodeSelect("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].Type", &resource.Actor[numActor].Type, optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_ActorName(numActor int, htmlAttrs string) templ.Component {
	if resource == nil || numActor >= len(resource.Actor) {
		return StringInput("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].Name", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].Name", resource.Actor[numActor].Name, htmlAttrs)
}
func (resource *ExampleScenario) T_ActorDescription(numActor int, htmlAttrs string) templ.Component {
	if resource == nil || numActor >= len(resource.Actor) {
		return StringInput("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].Description", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Actor["+strconv.Itoa(numActor)+"].Description", resource.Actor[numActor].Description, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceResourceId(numInstance int, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ResourceId", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ResourceId", &resource.Instance[numInstance].ResourceId, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceResourceType(numInstance int, htmlAttrs string) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numInstance >= len(resource.Instance) {
		return CodeSelect("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ResourceType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ResourceType", &resource.Instance[numInstance].ResourceType, optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceName(numInstance int, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Name", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Name", resource.Instance[numInstance].Name, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceDescription(numInstance int, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Description", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Description", resource.Instance[numInstance].Description, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceVersionVersionId(numInstance int, numVersion int, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numVersion >= len(resource.Instance[numInstance].Version) {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Version["+strconv.Itoa(numVersion)+"].VersionId", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Version["+strconv.Itoa(numVersion)+"].VersionId", &resource.Instance[numInstance].Version[numVersion].VersionId, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceVersionDescription(numInstance int, numVersion int, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numVersion >= len(resource.Instance[numInstance].Version) {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Version["+strconv.Itoa(numVersion)+"].Description", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].Version["+strconv.Itoa(numVersion)+"].Description", &resource.Instance[numInstance].Version[numVersion].Description, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceContainedInstanceResourceId(numInstance int, numContainedInstance int, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numContainedInstance >= len(resource.Instance[numInstance].ContainedInstance) {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ContainedInstance["+strconv.Itoa(numContainedInstance)+"].ResourceId", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ContainedInstance["+strconv.Itoa(numContainedInstance)+"].ResourceId", &resource.Instance[numInstance].ContainedInstance[numContainedInstance].ResourceId, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceContainedInstanceVersionId(numInstance int, numContainedInstance int, htmlAttrs string) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numContainedInstance >= len(resource.Instance[numInstance].ContainedInstance) {
		return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ContainedInstance["+strconv.Itoa(numContainedInstance)+"].VersionId", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance["+strconv.Itoa(numInstance)+"].ContainedInstance["+strconv.Itoa(numContainedInstance)+"].VersionId", resource.Instance[numInstance].ContainedInstance[numContainedInstance].VersionId, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessTitle(numProcess int, htmlAttrs string) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Title", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Title", &resource.Process[numProcess].Title, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessDescription(numProcess int, htmlAttrs string) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Description", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Description", resource.Process[numProcess].Description, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessPreConditions(numProcess int, htmlAttrs string) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].PreConditions", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].PreConditions", resource.Process[numProcess].PreConditions, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessPostConditions(numProcess int, htmlAttrs string) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].PostConditions", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].PostConditions", resource.Process[numProcess].PostConditions, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepPause(numProcess int, numStep int, htmlAttrs string) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return BoolInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Pause", nil, htmlAttrs)
	}
	return BoolInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Pause", resource.Process[numProcess].Step[numStep].Pause, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationNumber(numProcess int, numStep int, htmlAttrs string) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Number", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Number", &resource.Process[numProcess].Step[numStep].Operation.Number, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationType(numProcess int, numStep int, htmlAttrs string) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Type", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Type", resource.Process[numProcess].Step[numStep].Operation.Type, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationName(numProcess int, numStep int, htmlAttrs string) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Name", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Name", resource.Process[numProcess].Step[numStep].Operation.Name, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationInitiator(numProcess int, numStep int, htmlAttrs string) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Initiator", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Initiator", resource.Process[numProcess].Step[numStep].Operation.Initiator, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationReceiver(numProcess int, numStep int, htmlAttrs string) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Receiver", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Receiver", resource.Process[numProcess].Step[numStep].Operation.Receiver, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationDescription(numProcess int, numStep int, htmlAttrs string) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Description", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.Description", resource.Process[numProcess].Step[numStep].Operation.Description, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationInitiatorActive(numProcess int, numStep int, htmlAttrs string) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return BoolInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.InitiatorActive", nil, htmlAttrs)
	}
	return BoolInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.InitiatorActive", resource.Process[numProcess].Step[numStep].Operation.InitiatorActive, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationReceiverActive(numProcess int, numStep int, htmlAttrs string) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return BoolInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.ReceiverActive", nil, htmlAttrs)
	}
	return BoolInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Operation.ReceiverActive", resource.Process[numProcess].Step[numStep].Operation.ReceiverActive, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepAlternativeTitle(numProcess int, numStep int, numAlternative int, htmlAttrs string) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) || numAlternative >= len(resource.Process[numProcess].Step[numStep].Alternative) {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Alternative["+strconv.Itoa(numAlternative)+"].Title", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Alternative["+strconv.Itoa(numAlternative)+"].Title", &resource.Process[numProcess].Step[numStep].Alternative[numAlternative].Title, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepAlternativeDescription(numProcess int, numStep int, numAlternative int, htmlAttrs string) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) || numAlternative >= len(resource.Process[numProcess].Step[numStep].Alternative) {
		return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Alternative["+strconv.Itoa(numAlternative)+"].Description", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process["+strconv.Itoa(numProcess)+"].Step["+strconv.Itoa(numStep)+"].Alternative["+strconv.Itoa(numAlternative)+"].Description", resource.Process[numProcess].Step[numStep].Alternative[numAlternative].Description, htmlAttrs)
}
