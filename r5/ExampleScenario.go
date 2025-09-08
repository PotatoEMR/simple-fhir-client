package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

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
	VersionAlgorithmString *string                   `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                   `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                   `json:"name,omitempty"`
	Title                  *string                   `json:"title,omitempty"`
	Status                 string                    `json:"status"`
	Experimental           *bool                     `json:"experimental,omitempty"`
	Date                   *time.Time                `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	StructureProfileCanonical *string                                    `json:"structureProfileCanonical,omitempty"`
	StructureProfileUri       *string                                    `json:"structureProfileUri,omitempty"`
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
func (resource *ExampleScenario) T_VersionAlgorithmString(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ExampleScenario.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *ExampleScenario) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodingSelect("ExampleScenario.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("ExampleScenario.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ExampleScenario.Name", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Name", resource.Name, htmlAttrs)
}
func (resource *ExampleScenario) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ExampleScenario.Title", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Title", resource.Title, htmlAttrs)
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
func (resource *ExampleScenario) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ExampleScenario.Description", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Description", resource.Description, htmlAttrs)
}
func (resource *ExampleScenario) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("ExampleScenario.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ExampleScenario.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_Purpose(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ExampleScenario.Purpose", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *ExampleScenario) T_Copyright(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ExampleScenario.Copyright", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *ExampleScenario) T_CopyrightLabel(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("ExampleScenario.CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *ExampleScenario) T_ActorKey(numActor int, htmlAttrs string) templ.Component {

	if resource == nil || numActor >= len(resource.Actor) {
		return StringInput("ExampleScenario.Actor."+strconv.Itoa(numActor)+"..Key", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Actor."+strconv.Itoa(numActor)+"..Key", &resource.Actor[numActor].Key, htmlAttrs)
}
func (resource *ExampleScenario) T_ActorType(numActor int, htmlAttrs string) templ.Component {
	optionsValueSet := VSExamplescenario_actor_type

	if resource == nil || numActor >= len(resource.Actor) {
		return CodeSelect("ExampleScenario.Actor."+strconv.Itoa(numActor)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ExampleScenario.Actor."+strconv.Itoa(numActor)+"..Type", &resource.Actor[numActor].Type, optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_ActorTitle(numActor int, htmlAttrs string) templ.Component {

	if resource == nil || numActor >= len(resource.Actor) {
		return StringInput("ExampleScenario.Actor."+strconv.Itoa(numActor)+"..Title", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Actor."+strconv.Itoa(numActor)+"..Title", &resource.Actor[numActor].Title, htmlAttrs)
}
func (resource *ExampleScenario) T_ActorDescription(numActor int, htmlAttrs string) templ.Component {

	if resource == nil || numActor >= len(resource.Actor) {
		return StringInput("ExampleScenario.Actor."+strconv.Itoa(numActor)+"..Description", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Actor."+strconv.Itoa(numActor)+"..Description", resource.Actor[numActor].Description, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceKey(numInstance int, htmlAttrs string) templ.Component {

	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..Key", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..Key", &resource.Instance[numInstance].Key, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceStructureType(numInstance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numInstance >= len(resource.Instance) {
		return CodingSelect("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..StructureType", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..StructureType", &resource.Instance[numInstance].StructureType, optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceStructureVersion(numInstance int, htmlAttrs string) templ.Component {

	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..StructureVersion", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..StructureVersion", resource.Instance[numInstance].StructureVersion, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceStructureProfileCanonical(numInstance int, htmlAttrs string) templ.Component {

	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..StructureProfileCanonical", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..StructureProfileCanonical", resource.Instance[numInstance].StructureProfileCanonical, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceStructureProfileUri(numInstance int, htmlAttrs string) templ.Component {

	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..StructureProfileUri", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..StructureProfileUri", resource.Instance[numInstance].StructureProfileUri, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceTitle(numInstance int, htmlAttrs string) templ.Component {

	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..Title", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..Title", &resource.Instance[numInstance].Title, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceDescription(numInstance int, htmlAttrs string) templ.Component {

	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..Description", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..Description", resource.Instance[numInstance].Description, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceVersionKey(numInstance int, numVersion int, htmlAttrs string) templ.Component {

	if resource == nil || numInstance >= len(resource.Instance) || numVersion >= len(resource.Instance[numInstance].Version) {
		return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..Version."+strconv.Itoa(numVersion)+"..Key", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..Version."+strconv.Itoa(numVersion)+"..Key", &resource.Instance[numInstance].Version[numVersion].Key, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceVersionTitle(numInstance int, numVersion int, htmlAttrs string) templ.Component {

	if resource == nil || numInstance >= len(resource.Instance) || numVersion >= len(resource.Instance[numInstance].Version) {
		return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..Version."+strconv.Itoa(numVersion)+"..Title", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..Version."+strconv.Itoa(numVersion)+"..Title", &resource.Instance[numInstance].Version[numVersion].Title, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceVersionDescription(numInstance int, numVersion int, htmlAttrs string) templ.Component {

	if resource == nil || numInstance >= len(resource.Instance) || numVersion >= len(resource.Instance[numInstance].Version) {
		return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..Version."+strconv.Itoa(numVersion)+"..Description", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..Version."+strconv.Itoa(numVersion)+"..Description", resource.Instance[numInstance].Version[numVersion].Description, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceContainedInstanceInstanceReference(numInstance int, numContainedInstance int, htmlAttrs string) templ.Component {

	if resource == nil || numInstance >= len(resource.Instance) || numContainedInstance >= len(resource.Instance[numInstance].ContainedInstance) {
		return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..ContainedInstance."+strconv.Itoa(numContainedInstance)+"..InstanceReference", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..ContainedInstance."+strconv.Itoa(numContainedInstance)+"..InstanceReference", &resource.Instance[numInstance].ContainedInstance[numContainedInstance].InstanceReference, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceContainedInstanceVersionReference(numInstance int, numContainedInstance int, htmlAttrs string) templ.Component {

	if resource == nil || numInstance >= len(resource.Instance) || numContainedInstance >= len(resource.Instance[numInstance].ContainedInstance) {
		return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..ContainedInstance."+strconv.Itoa(numContainedInstance)+"..VersionReference", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Instance."+strconv.Itoa(numInstance)+"..ContainedInstance."+strconv.Itoa(numContainedInstance)+"..VersionReference", resource.Instance[numInstance].ContainedInstance[numContainedInstance].VersionReference, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessTitle(numProcess int, htmlAttrs string) templ.Component {

	if resource == nil || numProcess >= len(resource.Process) {
		return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Title", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Title", &resource.Process[numProcess].Title, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessDescription(numProcess int, htmlAttrs string) templ.Component {

	if resource == nil || numProcess >= len(resource.Process) {
		return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Description", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Description", resource.Process[numProcess].Description, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessPreConditions(numProcess int, htmlAttrs string) templ.Component {

	if resource == nil || numProcess >= len(resource.Process) {
		return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..PreConditions", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..PreConditions", resource.Process[numProcess].PreConditions, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessPostConditions(numProcess int, htmlAttrs string) templ.Component {

	if resource == nil || numProcess >= len(resource.Process) {
		return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..PostConditions", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..PostConditions", resource.Process[numProcess].PostConditions, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepNumber(numProcess int, numStep int, htmlAttrs string) templ.Component {

	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Number", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Number", resource.Process[numProcess].Step[numStep].Number, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepWorkflow(numProcess int, numStep int, htmlAttrs string) templ.Component {

	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Workflow", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Workflow", resource.Process[numProcess].Step[numStep].Workflow, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepPause(numProcess int, numStep int, htmlAttrs string) templ.Component {

	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return BoolInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Pause", nil, htmlAttrs)
	}
	return BoolInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Pause", resource.Process[numProcess].Step[numStep].Pause, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationType(numProcess int, numStep int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return CodingSelect("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Operation.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Operation.Type", resource.Process[numProcess].Step[numStep].Operation.Type, optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationTitle(numProcess int, numStep int, htmlAttrs string) templ.Component {

	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Operation.Title", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Operation.Title", &resource.Process[numProcess].Step[numStep].Operation.Title, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationInitiator(numProcess int, numStep int, htmlAttrs string) templ.Component {

	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Operation.Initiator", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Operation.Initiator", resource.Process[numProcess].Step[numStep].Operation.Initiator, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationReceiver(numProcess int, numStep int, htmlAttrs string) templ.Component {

	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Operation.Receiver", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Operation.Receiver", resource.Process[numProcess].Step[numStep].Operation.Receiver, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationDescription(numProcess int, numStep int, htmlAttrs string) templ.Component {

	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Operation.Description", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Operation.Description", resource.Process[numProcess].Step[numStep].Operation.Description, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationInitiatorActive(numProcess int, numStep int, htmlAttrs string) templ.Component {

	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return BoolInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Operation.InitiatorActive", nil, htmlAttrs)
	}
	return BoolInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Operation.InitiatorActive", resource.Process[numProcess].Step[numStep].Operation.InitiatorActive, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationReceiverActive(numProcess int, numStep int, htmlAttrs string) templ.Component {

	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return BoolInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Operation.ReceiverActive", nil, htmlAttrs)
	}
	return BoolInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Operation.ReceiverActive", resource.Process[numProcess].Step[numStep].Operation.ReceiverActive, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepAlternativeTitle(numProcess int, numStep int, numAlternative int, htmlAttrs string) templ.Component {

	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) || numAlternative >= len(resource.Process[numProcess].Step[numStep].Alternative) {
		return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Alternative."+strconv.Itoa(numAlternative)+"..Title", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Alternative."+strconv.Itoa(numAlternative)+"..Title", &resource.Process[numProcess].Step[numStep].Alternative[numAlternative].Title, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepAlternativeDescription(numProcess int, numStep int, numAlternative int, htmlAttrs string) templ.Component {

	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) || numAlternative >= len(resource.Process[numProcess].Step[numStep].Alternative) {
		return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Alternative."+strconv.Itoa(numAlternative)+"..Description", nil, htmlAttrs)
	}
	return StringInput("ExampleScenario.Process."+strconv.Itoa(numProcess)+"..Step."+strconv.Itoa(numStep)+"..Alternative."+strconv.Itoa(numAlternative)+"..Description", resource.Process[numProcess].Step[numStep].Alternative[numAlternative].Description, htmlAttrs)
}
