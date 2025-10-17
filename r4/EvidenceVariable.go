package r4

//generated with command go run ./bultaoreune
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
	Date              *FhirDateTime                    `json:"date,omitempty"`
	Publisher         *string                          `json:"publisher,omitempty"`
	Contact           []ContactDetail                  `json:"contact,omitempty"`
	Description       *string                          `json:"description,omitempty"`
	Note              []Annotation                     `json:"note,omitempty"`
	UseContext        []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                `json:"jurisdiction,omitempty"`
	Copyright         *string                          `json:"copyright,omitempty"`
	ApprovalDate      *FhirDate                        `json:"approvalDate,omitempty"`
	LastReviewDate    *FhirDate                        `json:"lastReviewDate,omitempty"`
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
	ParticipantEffectiveDateTime *FhirDateTime     `json:"participantEffectiveDateTime,omitempty"`
	ParticipantEffectivePeriod   *Period           `json:"participantEffectivePeriod,omitempty"`
	ParticipantEffectiveDuration *Duration         `json:"participantEffectiveDuration,omitempty"`
	ParticipantEffectiveTiming   *Timing           `json:"participantEffectiveTiming,omitempty"`
	TimeFromStart                *Duration         `json:"timeFromStart,omitempty"`
	GroupMeasure                 *string           `json:"groupMeasure,omitempty"`
}

type OtherEvidenceVariable EvidenceVariable

// struct -> json, automatically add resourceType=Patient
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
func (r EvidenceVariable) ResourceType() string {
	return "EvidenceVariable"
}

func (resource *EvidenceVariable) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *EvidenceVariable) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *EvidenceVariable) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *EvidenceVariable) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *EvidenceVariable) T_ShortTitle(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("shortTitle", nil, htmlAttrs)
	}
	return StringInput("shortTitle", resource.ShortTitle, htmlAttrs)
}
func (resource *EvidenceVariable) T_Subtitle(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("subtitle", nil, htmlAttrs)
	}
	return StringInput("subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *EvidenceVariable) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *EvidenceVariable) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *EvidenceVariable) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *EvidenceVariable) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *EvidenceVariable) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *EvidenceVariable) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *EvidenceVariable) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *EvidenceVariable) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("approvalDate", nil, htmlAttrs)
	}
	return FhirDateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *EvidenceVariable) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("lastReviewDate", nil, htmlAttrs)
	}
	return FhirDateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *EvidenceVariable) T_EffectivePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("effectivePeriod", resource.EffectivePeriod, htmlAttrs)
}
func (resource *EvidenceVariable) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_Author(numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthor >= len(resource.Author) {
		return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", &resource.Author[numAuthor], htmlAttrs)
}
func (resource *EvidenceVariable) T_Editor(numEditor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEditor >= len(resource.Editor) {
		return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", &resource.Editor[numEditor], htmlAttrs)
}
func (resource *EvidenceVariable) T_Reviewer(numReviewer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReviewer >= len(resource.Reviewer) {
		return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", &resource.Reviewer[numReviewer], htmlAttrs)
}
func (resource *EvidenceVariable) T_Endorser(numEndorser int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndorser >= len(resource.Endorser) {
		return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", &resource.Endorser[numEndorser], htmlAttrs)
}
func (resource *EvidenceVariable) T_RelatedArtifact(numRelatedArtifact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedArtifact >= len(resource.RelatedArtifact) {
		return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", &resource.RelatedArtifact[numRelatedArtifact], htmlAttrs)
}
func (resource *EvidenceVariable) T_Type(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSVariable_type

	if resource == nil {
		return CodeSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDescription(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].description", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].description", resource.Characteristic[numCharacteristic].Description, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionReference(frs []FhirResource, numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return ReferenceInput(frs, "characteristic["+strconv.Itoa(numCharacteristic)+"].definitionReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "characteristic["+strconv.Itoa(numCharacteristic)+"].definitionReference", &resource.Characteristic[numCharacteristic].DefinitionReference, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionCanonical(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionCanonical", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionCanonical", &resource.Characteristic[numCharacteristic].DefinitionCanonical, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionCodeableConcept", &resource.Characteristic[numCharacteristic].DefinitionCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionExpression(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return ExpressionInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionExpression", nil, htmlAttrs)
	}
	return ExpressionInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionExpression", &resource.Characteristic[numCharacteristic].DefinitionExpression, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionDataRequirement(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return DataRequirementInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionDataRequirement", nil, htmlAttrs)
	}
	return DataRequirementInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionDataRequirement", &resource.Characteristic[numCharacteristic].DefinitionDataRequirement, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionTriggerDefinition(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return TriggerDefinitionInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionTriggerDefinition", nil, htmlAttrs)
	}
	return TriggerDefinitionInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionTriggerDefinition", &resource.Characteristic[numCharacteristic].DefinitionTriggerDefinition, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicUsageContext(numCharacteristic int, numUsageContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numUsageContext >= len(resource.Characteristic[numCharacteristic].UsageContext) {
		return UsageContextInput("characteristic["+strconv.Itoa(numCharacteristic)+"].usageContext["+strconv.Itoa(numUsageContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("characteristic["+strconv.Itoa(numCharacteristic)+"].usageContext["+strconv.Itoa(numUsageContext)+"]", &resource.Characteristic[numCharacteristic].UsageContext[numUsageContext], htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicExclude(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].exclude", nil, htmlAttrs)
	}
	return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].exclude", resource.Characteristic[numCharacteristic].Exclude, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicParticipantEffectiveDateTime(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return FhirDateTimeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveDateTime", resource.Characteristic[numCharacteristic].ParticipantEffectiveDateTime, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicParticipantEffectivePeriod(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return PeriodInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectivePeriod", resource.Characteristic[numCharacteristic].ParticipantEffectivePeriod, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicParticipantEffectiveDuration(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return DurationInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveDuration", nil, htmlAttrs)
	}
	return DurationInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveDuration", resource.Characteristic[numCharacteristic].ParticipantEffectiveDuration, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicParticipantEffectiveTiming(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return TimingInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveTiming", nil, htmlAttrs)
	}
	return TimingInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveTiming", resource.Characteristic[numCharacteristic].ParticipantEffectiveTiming, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromStart(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return DurationInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromStart", nil, htmlAttrs)
	}
	return DurationInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromStart", resource.Characteristic[numCharacteristic].TimeFromStart, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicGroupMeasure(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSGroup_measure

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].groupMeasure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].groupMeasure", resource.Characteristic[numCharacteristic].GroupMeasure, optionsValueSet, htmlAttrs)
}
