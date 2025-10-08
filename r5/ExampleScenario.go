package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Date                   *FhirDateTime             `json:"date,omitempty"`
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
func (resource *ExampleScenario) T_VersionAlgorithmString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *ExampleScenario) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *ExampleScenario) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
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
func (resource *ExampleScenario) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
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
func (resource *ExampleScenario) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *ExampleScenario) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *ExampleScenario) T_CopyrightLabel(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *ExampleScenario) T_ActorKey(numActor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActor >= len(resource.Actor) {
		return StringInput("actor["+strconv.Itoa(numActor)+"].key", nil, htmlAttrs)
	}
	return StringInput("actor["+strconv.Itoa(numActor)+"].key", &resource.Actor[numActor].Key, htmlAttrs)
}
func (resource *ExampleScenario) T_ActorType(numActor int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSExamplescenario_actor_type

	if resource == nil || numActor >= len(resource.Actor) {
		return CodeSelect("actor["+strconv.Itoa(numActor)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("actor["+strconv.Itoa(numActor)+"].type", &resource.Actor[numActor].Type, optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_ActorTitle(numActor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActor >= len(resource.Actor) {
		return StringInput("actor["+strconv.Itoa(numActor)+"].title", nil, htmlAttrs)
	}
	return StringInput("actor["+strconv.Itoa(numActor)+"].title", &resource.Actor[numActor].Title, htmlAttrs)
}
func (resource *ExampleScenario) T_ActorDescription(numActor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActor >= len(resource.Actor) {
		return StringInput("actor["+strconv.Itoa(numActor)+"].description", nil, htmlAttrs)
	}
	return StringInput("actor["+strconv.Itoa(numActor)+"].description", resource.Actor[numActor].Description, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceKey(numInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].key", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].key", &resource.Instance[numInstance].Key, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceStructureType(numInstance int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return CodingSelect("instance["+strconv.Itoa(numInstance)+"].structureType", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("instance["+strconv.Itoa(numInstance)+"].structureType", &resource.Instance[numInstance].StructureType, optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceStructureVersion(numInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].structureVersion", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].structureVersion", resource.Instance[numInstance].StructureVersion, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceStructureProfileCanonical(numInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].structureProfileCanonical", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].structureProfileCanonical", resource.Instance[numInstance].StructureProfileCanonical, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceStructureProfileUri(numInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].structureProfileUri", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].structureProfileUri", resource.Instance[numInstance].StructureProfileUri, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceTitle(numInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].title", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].title", &resource.Instance[numInstance].Title, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceDescription(numInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].description", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].description", resource.Instance[numInstance].Description, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceContent(frs []FhirResource, numInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) {
		return ReferenceInput(frs, "instance["+strconv.Itoa(numInstance)+"].content", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "instance["+strconv.Itoa(numInstance)+"].content", resource.Instance[numInstance].Content, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceVersionKey(numInstance int, numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numVersion >= len(resource.Instance[numInstance].Version) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].version["+strconv.Itoa(numVersion)+"].key", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].version["+strconv.Itoa(numVersion)+"].key", &resource.Instance[numInstance].Version[numVersion].Key, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceVersionTitle(numInstance int, numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numVersion >= len(resource.Instance[numInstance].Version) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].version["+strconv.Itoa(numVersion)+"].title", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].version["+strconv.Itoa(numVersion)+"].title", &resource.Instance[numInstance].Version[numVersion].Title, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceVersionDescription(numInstance int, numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numVersion >= len(resource.Instance[numInstance].Version) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].version["+strconv.Itoa(numVersion)+"].description", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].version["+strconv.Itoa(numVersion)+"].description", resource.Instance[numInstance].Version[numVersion].Description, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceVersionContent(frs []FhirResource, numInstance int, numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numVersion >= len(resource.Instance[numInstance].Version) {
		return ReferenceInput(frs, "instance["+strconv.Itoa(numInstance)+"].version["+strconv.Itoa(numVersion)+"].content", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "instance["+strconv.Itoa(numInstance)+"].version["+strconv.Itoa(numVersion)+"].content", resource.Instance[numInstance].Version[numVersion].Content, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceContainedInstanceInstanceReference(numInstance int, numContainedInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numContainedInstance >= len(resource.Instance[numInstance].ContainedInstance) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].containedInstance["+strconv.Itoa(numContainedInstance)+"].instanceReference", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].containedInstance["+strconv.Itoa(numContainedInstance)+"].instanceReference", &resource.Instance[numInstance].ContainedInstance[numContainedInstance].InstanceReference, htmlAttrs)
}
func (resource *ExampleScenario) T_InstanceContainedInstanceVersionReference(numInstance int, numContainedInstance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstance >= len(resource.Instance) || numContainedInstance >= len(resource.Instance[numInstance].ContainedInstance) {
		return StringInput("instance["+strconv.Itoa(numInstance)+"].containedInstance["+strconv.Itoa(numContainedInstance)+"].versionReference", nil, htmlAttrs)
	}
	return StringInput("instance["+strconv.Itoa(numInstance)+"].containedInstance["+strconv.Itoa(numContainedInstance)+"].versionReference", resource.Instance[numInstance].ContainedInstance[numContainedInstance].VersionReference, htmlAttrs)
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
func (resource *ExampleScenario) T_ProcessStepNumber(numProcess int, numStep int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].number", nil, htmlAttrs)
	}
	return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].number", resource.Process[numProcess].Step[numStep].Number, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepWorkflow(numProcess int, numStep int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].workflow", nil, htmlAttrs)
	}
	return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].workflow", resource.Process[numProcess].Step[numStep].Workflow, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepPause(numProcess int, numStep int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return BoolInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].pause", nil, htmlAttrs)
	}
	return BoolInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].pause", resource.Process[numProcess].Step[numStep].Pause, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationType(numProcess int, numStep int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return CodingSelect("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.type", resource.Process[numProcess].Step[numStep].Operation.Type, optionsValueSet, htmlAttrs)
}
func (resource *ExampleScenario) T_ProcessStepOperationTitle(numProcess int, numStep int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcess >= len(resource.Process) || numStep >= len(resource.Process[numProcess].Step) {
		return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.title", nil, htmlAttrs)
	}
	return StringInput("process["+strconv.Itoa(numProcess)+"].step["+strconv.Itoa(numStep)+"].operation.title", &resource.Process[numProcess].Step[numStep].Operation.Title, htmlAttrs)
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
