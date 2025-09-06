package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/EvidenceVariable
type EvidenceVariable struct {
	Id                *string                          `json:"id,omitempty"`
	Meta              *Meta                            `json:"meta,omitempty"`
	ImplicitRules     *string                          `json:"implicitRules,omitempty"`
	Language          *string                          `json:"language,omitempty"`
	Text              *Narrative                       `json:"text,omitempty"`
	Contained         []Resource                       `json:"contained,omitempty"`
	Extension         []Extension                      `json:"extension,omitempty"`
	ModifierExtension []Extension                      `json:"modifierExtension,omitempty"`
	Url               *string                          `json:"url,omitempty"`
	Identifier        []Identifier                     `json:"identifier,omitempty"`
	Version           *string                          `json:"version,omitempty"`
	Name              *string                          `json:"name,omitempty"`
	Title             *string                          `json:"title,omitempty"`
	ShortTitle        *string                          `json:"shortTitle,omitempty"`
	Subtitle          *string                          `json:"subtitle,omitempty"`
	Status            string                           `json:"status"`
	Date              *string                          `json:"date,omitempty"`
	Publisher         *string                          `json:"publisher,omitempty"`
	Contact           []ContactDetail                  `json:"contact,omitempty"`
	Description       *string                          `json:"description,omitempty"`
	Note              []Annotation                     `json:"note,omitempty"`
	UseContext        []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                `json:"jurisdiction,omitempty"`
	Copyright         *string                          `json:"copyright,omitempty"`
	ApprovalDate      *string                          `json:"approvalDate,omitempty"`
	LastReviewDate    *string                          `json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period                          `json:"effectivePeriod,omitempty"`
	Topic             []CodeableConcept                `json:"topic,omitempty"`
	Author            []ContactDetail                  `json:"author,omitempty"`
	Editor            []ContactDetail                  `json:"editor,omitempty"`
	Reviewer          []ContactDetail                  `json:"reviewer,omitempty"`
	Endorser          []ContactDetail                  `json:"endorser,omitempty"`
	RelatedArtifact   []RelatedArtifact                `json:"relatedArtifact,omitempty"`
	Type              *string                          `json:"type,omitempty"`
	Characteristic    []EvidenceVariableCharacteristic `json:"characteristic"`
}

// http://hl7.org/fhir/r4/StructureDefinition/EvidenceVariable
type EvidenceVariableCharacteristic struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Description                  *string           `json:"description,omitempty"`
	DefinitionReference          Reference         `json:"definitionReference"`
	DefinitionCanonical          string            `json:"definitionCanonical"`
	DefinitionCodeableConcept    CodeableConcept   `json:"definitionCodeableConcept"`
	DefinitionExpression         Expression        `json:"definitionExpression"`
	DefinitionDataRequirement    DataRequirement   `json:"definitionDataRequirement"`
	DefinitionTriggerDefinition  TriggerDefinition `json:"definitionTriggerDefinition"`
	UsageContext                 []UsageContext    `json:"usageContext,omitempty"`
	Exclude                      *bool             `json:"exclude,omitempty"`
	ParticipantEffectiveDateTime *string           `json:"participantEffectiveDateTime,omitempty"`
	ParticipantEffectivePeriod   *Period           `json:"participantEffectivePeriod,omitempty"`
	ParticipantEffectiveDuration *Duration         `json:"participantEffectiveDuration,omitempty"`
	ParticipantEffectiveTiming   *Timing           `json:"participantEffectiveTiming,omitempty"`
	TimeFromStart                *Duration         `json:"timeFromStart,omitempty"`
	GroupMeasure                 *string           `json:"groupMeasure,omitempty"`
}

type OtherEvidenceVariable EvidenceVariable

// on convert struct to json, automatically add resourceType=EvidenceVariable
func (r EvidenceVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEvidenceVariable
		ResourceType string `json:"resourceType"`
	}{
		OtherEvidenceVariable: OtherEvidenceVariable(r),
		ResourceType:          "EvidenceVariable",
	})
}

func (resource *EvidenceVariable) T_Id() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Id", nil)
	}
	return StringInput("EvidenceVariable.Id", resource.Id)
}
func (resource *EvidenceVariable) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.ImplicitRules", nil)
	}
	return StringInput("EvidenceVariable.ImplicitRules", resource.ImplicitRules)
}
func (resource *EvidenceVariable) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("EvidenceVariable.Language", nil, optionsValueSet)
	}
	return CodeSelect("EvidenceVariable.Language", resource.Language, optionsValueSet)
}
func (resource *EvidenceVariable) T_Url() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Url", nil)
	}
	return StringInput("EvidenceVariable.Url", resource.Url)
}
func (resource *EvidenceVariable) T_Version() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Version", nil)
	}
	return StringInput("EvidenceVariable.Version", resource.Version)
}
func (resource *EvidenceVariable) T_Name() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Name", nil)
	}
	return StringInput("EvidenceVariable.Name", resource.Name)
}
func (resource *EvidenceVariable) T_Title() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Title", nil)
	}
	return StringInput("EvidenceVariable.Title", resource.Title)
}
func (resource *EvidenceVariable) T_ShortTitle() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.ShortTitle", nil)
	}
	return StringInput("EvidenceVariable.ShortTitle", resource.ShortTitle)
}
func (resource *EvidenceVariable) T_Subtitle() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Subtitle", nil)
	}
	return StringInput("EvidenceVariable.Subtitle", resource.Subtitle)
}
func (resource *EvidenceVariable) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("EvidenceVariable.Status", nil, optionsValueSet)
	}
	return CodeSelect("EvidenceVariable.Status", &resource.Status, optionsValueSet)
}
func (resource *EvidenceVariable) T_Date() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Date", nil)
	}
	return StringInput("EvidenceVariable.Date", resource.Date)
}
func (resource *EvidenceVariable) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Publisher", nil)
	}
	return StringInput("EvidenceVariable.Publisher", resource.Publisher)
}
func (resource *EvidenceVariable) T_Description() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Description", nil)
	}
	return StringInput("EvidenceVariable.Description", resource.Description)
}
func (resource *EvidenceVariable) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("EvidenceVariable.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EvidenceVariable.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *EvidenceVariable) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Copyright", nil)
	}
	return StringInput("EvidenceVariable.Copyright", resource.Copyright)
}
func (resource *EvidenceVariable) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.ApprovalDate", nil)
	}
	return StringInput("EvidenceVariable.ApprovalDate", resource.ApprovalDate)
}
func (resource *EvidenceVariable) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.LastReviewDate", nil)
	}
	return StringInput("EvidenceVariable.LastReviewDate", resource.LastReviewDate)
}
func (resource *EvidenceVariable) T_Topic(numTopic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Topic) >= numTopic {
		return CodeableConceptSelect("EvidenceVariable.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EvidenceVariable.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet)
}
func (resource *EvidenceVariable) T_Type() templ.Component {
	optionsValueSet := VSVariable_type

	if resource == nil {
		return CodeSelect("EvidenceVariable.Type", nil, optionsValueSet)
	}
	return CodeSelect("EvidenceVariable.Type", resource.Type, optionsValueSet)
}
func (resource *EvidenceVariable) T_CharacteristicId(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Id", nil)
	}
	return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Id", resource.Characteristic[numCharacteristic].Id)
}
func (resource *EvidenceVariable) T_CharacteristicDescription(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Description", nil)
	}
	return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Description", resource.Characteristic[numCharacteristic].Description)
}
func (resource *EvidenceVariable) T_CharacteristicExclude(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return BoolInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Exclude", nil)
	}
	return BoolInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Exclude", resource.Characteristic[numCharacteristic].Exclude)
}
func (resource *EvidenceVariable) T_CharacteristicGroupMeasure(numCharacteristic int) templ.Component {
	optionsValueSet := VSGroup_measure

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return CodeSelect("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].GroupMeasure", nil, optionsValueSet)
	}
	return CodeSelect("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].GroupMeasure", resource.Characteristic[numCharacteristic].GroupMeasure, optionsValueSet)
}
