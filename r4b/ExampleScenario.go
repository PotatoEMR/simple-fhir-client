package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
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
	Date              *FhirDateTime             `json:"date,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
func (r ExampleScenario) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherExampleScenario
		ResourceType string `json:"resourceType"`
	}{
		OtherExampleScenario: OtherExampleScenario(r),
		ResourceType:         "ExampleScenario",
	})
}

// json -> struct, first reject if resourceType != ExampleScenario
func (r *ExampleScenario) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "ExampleScenario" {
		return errors.New("resourceType not ExampleScenario")
	}
	return json.Unmarshal(data, (*OtherExampleScenario)(r))
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
func (resource *ExampleScenario) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *ExampleScenario) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *ExampleScenario) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *ExampleScenario) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *ExampleScenario) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *ExampleScenario) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *ExampleScenario) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *ExampleScenario) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *ExampleScenario) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *ExampleScenario) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *ExampleScenario) T_Workflow(numWorkflow int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numWorkflow >= len(resource.Workflow) {
		return StringInput("workflow["+strconv.Itoa(numWorkflow)+"]", nil, htmlAttrs)
	}
	return StringInput("workflow["+strconv.Itoa(numWorkflow)+"]", &resource.Workflow[numWorkflow], htmlAttrs)
}
func (resource *ExampleScenario) T_ActorActorId(numActor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActor >= len(resource.Actor) {
		return StringInput("actor["+strconv.Itoa(numActor)+"].actorId", nil, htmlAttrs)
	}
	return StringInput("actor["+strconv.Itoa(numActor)+"].actorId", &resource.Actor[numActor].ActorId, htmlAttrs)
}
func (resource *ExampleScenario) T_ActorType(numActor int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSExamplescenario_actor_type

	if resource == nil || numActor >= len(resource.Actor) {
		return CodeSelect("actor["+strconv.Itoa(numActor)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("actor["+strconv.Itoa(numActor)+"].type", &resource.Actor[numActor].Type, optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_ActorName(numActor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActor >= len(resource.Actor) {
		return StringInput("actor["+strconv.Itoa(numActor)+"].name", nil, htmlAttrs)
	}
	return StringInput("actor["+strconv.Itoa(numActor)+"].name", resource.Actor[numActor].Name, htmlAttrs)
}
func (resource *ExampleScenario) T_ActorDescription(numActor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActor >= len(resource.Actor) {
		return StringInput("actor["+strconv.Itoa(numActor)+"].description", nil, htmlAttrs)
	}
	return StringInput("actor["+strconv.Itoa(numActor)+"].description", resource.Actor[numActor].Description, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceResourceId(numInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].resourceId", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].resourceId", &resource.Instance[numInstance].ResourceId, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceResourceType(numInstance int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numInstance >= len(resource.Instance) {
		return CodeSelect("instance["+strconv.Itoa(numInstance)+"].resourceType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("instance["+strconv.Itoa(numInstance)+"].resourceType", &resource.Instance[numInstance].ResourceType, optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceName(numInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].name", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].name", resource.Instance[numInstance].Name, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceDescription(numInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].description", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].description", resource.Instance[numInstance].Description, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceVersionVersionId(numInstance int, numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numVersion >= len(resource.Instance[numInstance].Version) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].version["+strconv.Itoa(numVersion)+"].versionId", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].version["+strconv.Itoa(numVersion)+"].versionId", &resource.Instance[numInstance].Version[numVersion].VersionId, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceVersionDescription(numInstance int, numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numVersion >= len(resource.Instance[numInstance].Version) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].version["+strconv.Itoa(numVersion)+"].description", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].version["+strconv.Itoa(numVersion)+"].description", &resource.Instance[numInstance].Version[numVersion].Description, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceContainedInstanceResourceId(numInstance int, numContainedInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numContainedInstance >= len(resource.Instance[numInstance].ContainedInstance) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].containedInstance["+strconv.Itoa(numContainedInstance)+"].resourceId", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].containedInstance["+strconv.Itoa(numContainedInstance)+"].resourceId", &resource.Instance[numInstance].ContainedInstance[numContainedInstance].ResourceId, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceContainedInstanceVersionId(numInstance int, numContainedInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numContainedInstance >= len(resource.Instance[numInstance].ContainedInstance) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].containedInstance["+strconv.Itoa(numContainedInstance)+"].versionId", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].containedInstance["+strconv.Itoa(numContainedInstance)+"].versionId", resource.Instance[numInstance].ContainedInstance[numContainedInstance].VersionId, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessTitle(numProcess int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) {
		return StringInput("process["+strconv.Itoa(numProcess)+"].title", nil, htmlAttrs)
	}
	return StringInput("process["+strconv.Itoa(numProcess)+"].title", &resource.Process[numProcess].Title, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessDescription(numProcess int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) {
		return StringInput("process["+strconv.Itoa(numProcess)+"].description", nil, htmlAttrs)
	}
	return StringInput("process["+strconv.Itoa(numProcess)+"].description", resource.Process[numProcess].Description, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessPreConditions(numProcess int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) {
		return StringInput("process["+strconv.Itoa(numProcess)+"].preConditions", nil, htmlAttrs)
	}
	return StringInput("process["+strconv.Itoa(numProcess)+"].preConditions", resource.Process[numProcess].PreConditions, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessPostConditions(numProcess int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) {
		return StringInput("process["+strconv.Itoa(numProcess)+"].postConditions", nil, htmlAttrs)
	}
	return StringInput("process["+strconv.Itoa(numProcess)+"].postConditions", resource.Process[numProcess].PostConditions, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepPause(numProcess int, numStep int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return BoolInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].pause", nil, htmlAttrs)
	}
	return BoolInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].pause", resource.Process[numProcess].Step[numStep].Pause, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationNumber(numProcess int, numStep int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.number", nil, htmlAttrs)
	}
	return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.number", &resource.Process[numProcess].Step[numStep].Operation.Number, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationType(numProcess int, numStep int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.type", nil, htmlAttrs)
	}
	return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.type", resource.Process[numProcess].Step[numStep].Operation.Type, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationName(numProcess int, numStep int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.name", nil, htmlAttrs)
	}
	return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.name", resource.Process[numProcess].Step[numStep].Operation.Name, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationInitiator(numProcess int, numStep int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.initiator", nil, htmlAttrs)
	}
	return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.initiator", resource.Process[numProcess].Step[numStep].Operation.Initiator, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationReceiver(numProcess int, numStep int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.receiver", nil, htmlAttrs)
	}
	return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.receiver", resource.Process[numProcess].Step[numStep].Operation.Receiver, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationDescription(numProcess int, numStep int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.description", nil, htmlAttrs)
	}
	return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.description", resource.Process[numProcess].Step[numStep].Operation.Description, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationInitiatorActive(numProcess int, numStep int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return BoolInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.initiatorActive", nil, htmlAttrs)
	}
	return BoolInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.initiatorActive", resource.Process[numProcess].Step[numStep].Operation.InitiatorActive, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationReceiverActive(numProcess int, numStep int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return BoolInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.receiverActive", nil, htmlAttrs)
	}
	return BoolInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.receiverActive", resource.Process[numProcess].Step[numStep].Operation.ReceiverActive, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepAlternativeTitle(numProcess int, numStep int, numAlternative int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) || numAlternative >= len(resource.Process[numProcess].Step[numStep].Alternative) {
		return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].alternative["+strconv.Itoa(numAlternative)+"].title", nil, htmlAttrs)
	}
	return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].alternative["+strconv.Itoa(numAlternative)+"].title", &resource.Process[numProcess].Step[numStep].Alternative[numAlternative].Title, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepAlternativeDescription(numProcess int, numStep int, numAlternative int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) || numAlternative >= len(resource.Process[numProcess].Step[numStep].Alternative) {
		return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].alternative["+strconv.Itoa(numAlternative)+"].description", nil, htmlAttrs)
	}
	return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].alternative["+strconv.Itoa(numAlternative)+"].description", resource.Process[numProcess].Step[numStep].Alternative[numAlternative].Description, htmlAttrs)
}
