package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Date              *time.Time                       `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher         *string                          `json:"publisher,omitempty"`
	Contact           []ContactDetail                  `json:"contact,omitempty"`
	Description       *string                          `json:"description,omitempty"`
	Note              []Annotation                     `json:"note,omitempty"`
	UseContext        []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                `json:"jurisdiction,omitempty"`
	Copyright         *string                          `json:"copyright,omitempty"`
	ApprovalDate      *time.Time                       `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate    *time.Time                       `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
	ParticipantEffectiveDateTime *time.Time        `json:"participantEffectiveDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r EvidenceVariable) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "EvidenceVariable/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "EvidenceVariable"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *EvidenceVariable) T_Url(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Url", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Url", resource.Url, htmlAttrs)
}
func (resource *EvidenceVariable) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Version", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Version", resource.Version, htmlAttrs)
}
func (resource *EvidenceVariable) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Name", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Name", resource.Name, htmlAttrs)
}
func (resource *EvidenceVariable) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Title", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Title", resource.Title, htmlAttrs)
}
func (resource *EvidenceVariable) T_ShortTitle(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.ShortTitle", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.ShortTitle", resource.ShortTitle, htmlAttrs)
}
func (resource *EvidenceVariable) T_Subtitle(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Subtitle", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *EvidenceVariable) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("EvidenceVariable.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("EvidenceVariable.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_Date(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("EvidenceVariable.Date", nil, htmlAttrs)
	}
	return DateTimeInput("EvidenceVariable.Date", resource.Date, htmlAttrs)
}
func (resource *EvidenceVariable) T_Publisher(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Publisher", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *EvidenceVariable) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Description", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Description", resource.Description, htmlAttrs)
}
func (resource *EvidenceVariable) T_Note(numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("EvidenceVariable.Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("EvidenceVariable.Note."+strconv.Itoa(numNote)+".", &resource.Note[numNote], htmlAttrs)
}
func (resource *EvidenceVariable) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("EvidenceVariable.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EvidenceVariable.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_Copyright(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Copyright", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *EvidenceVariable) T_ApprovalDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("EvidenceVariable.ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("EvidenceVariable.ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *EvidenceVariable) T_LastReviewDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("EvidenceVariable.LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("EvidenceVariable.LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *EvidenceVariable) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("EvidenceVariable.Topic."+strconv.Itoa(numTopic)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EvidenceVariable.Topic."+strconv.Itoa(numTopic)+".", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_Type(htmlAttrs string) templ.Component {
	optionsValueSet := VSVariable_type

	if resource == nil {
		return CodeSelect("EvidenceVariable.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("EvidenceVariable.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDescription(numCharacteristic int, htmlAttrs string) templ.Component {

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("EvidenceVariable.Characteristic."+strconv.Itoa(numCharacteristic)+"..Description", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Characteristic."+strconv.Itoa(numCharacteristic)+"..Description", resource.Characteristic[numCharacteristic].Description, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionCanonical(numCharacteristic int, htmlAttrs string) templ.Component {

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("EvidenceVariable.Characteristic."+strconv.Itoa(numCharacteristic)+"..DefinitionCanonical", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Characteristic."+strconv.Itoa(numCharacteristic)+"..DefinitionCanonical", &resource.Characteristic[numCharacteristic].DefinitionCanonical, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("EvidenceVariable.Characteristic."+strconv.Itoa(numCharacteristic)+"..DefinitionCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EvidenceVariable.Characteristic."+strconv.Itoa(numCharacteristic)+"..DefinitionCodeableConcept", &resource.Characteristic[numCharacteristic].DefinitionCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicExclude(numCharacteristic int, htmlAttrs string) templ.Component {

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("EvidenceVariable.Characteristic."+strconv.Itoa(numCharacteristic)+"..Exclude", nil, htmlAttrs)
	}
	return BoolInput("EvidenceVariable.Characteristic."+strconv.Itoa(numCharacteristic)+"..Exclude", resource.Characteristic[numCharacteristic].Exclude, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicParticipantEffectiveDateTime(numCharacteristic int, htmlAttrs string) templ.Component {

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return DateTimeInput("EvidenceVariable.Characteristic."+strconv.Itoa(numCharacteristic)+"..ParticipantEffectiveDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("EvidenceVariable.Characteristic."+strconv.Itoa(numCharacteristic)+"..ParticipantEffectiveDateTime", resource.Characteristic[numCharacteristic].ParticipantEffectiveDateTime, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicGroupMeasure(numCharacteristic int, htmlAttrs string) templ.Component {
	optionsValueSet := VSGroup_measure

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeSelect("EvidenceVariable.Characteristic."+strconv.Itoa(numCharacteristic)+"..GroupMeasure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("EvidenceVariable.Characteristic."+strconv.Itoa(numCharacteristic)+"..GroupMeasure", resource.Characteristic[numCharacteristic].GroupMeasure, optionsValueSet, htmlAttrs)
}
