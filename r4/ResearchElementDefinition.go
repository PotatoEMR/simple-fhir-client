package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/ResearchElementDefinition
type ResearchElementDefinition struct {
	Id                     *string                                   `json:"id,omitempty"`
	Meta                   *Meta                                     `json:"meta,omitempty"`
	ImplicitRules          *string                                   `json:"implicitRules,omitempty"`
	Language               *string                                   `json:"language,omitempty"`
	Text                   *Narrative                                `json:"text,omitempty"`
	Contained              []Resource                                `json:"contained,omitempty"`
	Extension              []Extension                               `json:"extension,omitempty"`
	ModifierExtension      []Extension                               `json:"modifierExtension,omitempty"`
	Url                    *string                                   `json:"url,omitempty"`
	Identifier             []Identifier                              `json:"identifier,omitempty"`
	Version                *string                                   `json:"version,omitempty"`
	Name                   *string                                   `json:"name,omitempty"`
	Title                  *string                                   `json:"title,omitempty"`
	ShortTitle             *string                                   `json:"shortTitle,omitempty"`
	Subtitle               *string                                   `json:"subtitle,omitempty"`
	Status                 string                                    `json:"status"`
	Experimental           *bool                                     `json:"experimental,omitempty"`
	SubjectCodeableConcept *CodeableConcept                          `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference                                `json:"subjectReference,omitempty"`
	Date                   *FhirDateTime                             `json:"date,omitempty"`
	Publisher              *string                                   `json:"publisher,omitempty"`
	Contact                []ContactDetail                           `json:"contact,omitempty"`
	Description            *string                                   `json:"description,omitempty"`
	Comment                []string                                  `json:"comment,omitempty"`
	UseContext             []UsageContext                            `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                         `json:"jurisdiction,omitempty"`
	Purpose                *string                                   `json:"purpose,omitempty"`
	Usage                  *string                                   `json:"usage,omitempty"`
	Copyright              *string                                   `json:"copyright,omitempty"`
	ApprovalDate           *FhirDate                                 `json:"approvalDate,omitempty"`
	LastReviewDate         *FhirDate                                 `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                                   `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept                         `json:"topic,omitempty"`
	Author                 []ContactDetail                           `json:"author,omitempty"`
	Editor                 []ContactDetail                           `json:"editor,omitempty"`
	Reviewer               []ContactDetail                           `json:"reviewer,omitempty"`
	Endorser               []ContactDetail                           `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact                         `json:"relatedArtifact,omitempty"`
	Library                []string                                  `json:"library,omitempty"`
	Type                   string                                    `json:"type"`
	VariableType           *string                                   `json:"variableType,omitempty"`
	Characteristic         []ResearchElementDefinitionCharacteristic `json:"characteristic"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ResearchElementDefinition
type ResearchElementDefinitionCharacteristic struct {
	Id                                *string          `json:"id,omitempty"`
	Extension                         []Extension      `json:"extension,omitempty"`
	ModifierExtension                 []Extension      `json:"modifierExtension,omitempty"`
	DefinitionCodeableConcept         CodeableConcept  `json:"definitionCodeableConcept"`
	DefinitionCanonical               string           `json:"definitionCanonical"`
	DefinitionExpression              Expression       `json:"definitionExpression"`
	DefinitionDataRequirement         DataRequirement  `json:"definitionDataRequirement"`
	UsageContext                      []UsageContext   `json:"usageContext,omitempty"`
	Exclude                           *bool            `json:"exclude,omitempty"`
	UnitOfMeasure                     *CodeableConcept `json:"unitOfMeasure,omitempty"`
	StudyEffectiveDescription         *string          `json:"studyEffectiveDescription,omitempty"`
	StudyEffectiveDateTime            *FhirDateTime    `json:"studyEffectiveDateTime,omitempty"`
	StudyEffectivePeriod              *Period          `json:"studyEffectivePeriod,omitempty"`
	StudyEffectiveDuration            *Duration        `json:"studyEffectiveDuration,omitempty"`
	StudyEffectiveTiming              *Timing          `json:"studyEffectiveTiming,omitempty"`
	StudyEffectiveTimeFromStart       *Duration        `json:"studyEffectiveTimeFromStart,omitempty"`
	StudyEffectiveGroupMeasure        *string          `json:"studyEffectiveGroupMeasure,omitempty"`
	ParticipantEffectiveDescription   *string          `json:"participantEffectiveDescription,omitempty"`
	ParticipantEffectiveDateTime      *FhirDateTime    `json:"participantEffectiveDateTime,omitempty"`
	ParticipantEffectivePeriod        *Period          `json:"participantEffectivePeriod,omitempty"`
	ParticipantEffectiveDuration      *Duration        `json:"participantEffectiveDuration,omitempty"`
	ParticipantEffectiveTiming        *Timing          `json:"participantEffectiveTiming,omitempty"`
	ParticipantEffectiveTimeFromStart *Duration        `json:"participantEffectiveTimeFromStart,omitempty"`
	ParticipantEffectiveGroupMeasure  *string          `json:"participantEffectiveGroupMeasure,omitempty"`
}

type OtherResearchElementDefinition ResearchElementDefinition

// on convert struct to json, automatically add resourceType=ResearchElementDefinition
func (r ResearchElementDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchElementDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchElementDefinition: OtherResearchElementDefinition(r),
		ResourceType:                   "ResearchElementDefinition",
	})
}
func (r ResearchElementDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ResearchElementDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ResearchElementDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ResearchElementDefinition) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_ShortTitle(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("shortTitle", nil, htmlAttrs)
	}
	return StringInput("shortTitle", resource.ShortTitle, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Subtitle(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("subtitle", nil, htmlAttrs)
	}
	return StringInput("subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_SubjectCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("subjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subjectCodeableConcept", resource.SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_SubjectReference(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("subjectReference", nil, htmlAttrs)
	}
	return ReferenceInput("subjectReference", resource.SubjectReference, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Comment(numComment int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComment >= len(resource.Comment) {
		return StringInput("comment["+strconv.Itoa(numComment)+"]", nil, htmlAttrs)
	}
	return StringInput("comment["+strconv.Itoa(numComment)+"]", &resource.Comment[numComment], htmlAttrs)
}
func (resource *ResearchElementDefinition) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Usage(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("usage", nil, htmlAttrs)
	}
	return StringInput("usage", resource.Usage, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("approvalDate", nil, htmlAttrs)
	}
	return FhirDateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("lastReviewDate", nil, htmlAttrs)
	}
	return FhirDateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_EffectivePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("effectivePeriod", resource.EffectivePeriod, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Author(numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthor >= len(resource.Author) {
		return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", &resource.Author[numAuthor], htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Editor(numEditor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEditor >= len(resource.Editor) {
		return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", &resource.Editor[numEditor], htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Reviewer(numReviewer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReviewer >= len(resource.Reviewer) {
		return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", &resource.Reviewer[numReviewer], htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Endorser(numEndorser int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndorser >= len(resource.Endorser) {
		return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", &resource.Endorser[numEndorser], htmlAttrs)
}
func (resource *ResearchElementDefinition) T_RelatedArtifact(numRelatedArtifact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedArtifact >= len(resource.RelatedArtifact) {
		return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", &resource.RelatedArtifact[numRelatedArtifact], htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Library(numLibrary int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLibrary >= len(resource.Library) {
		return StringInput("library["+strconv.Itoa(numLibrary)+"]", nil, htmlAttrs)
	}
	return StringInput("library["+strconv.Itoa(numLibrary)+"]", &resource.Library[numLibrary], htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Type(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSResearch_element_type

	if resource == nil {
		return CodeSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_VariableType(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSVariable_type

	if resource == nil {
		return CodeSelect("variableType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("variableType", resource.VariableType, optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicDefinitionCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionCodeableConcept", &resource.Characteristic[numCharacteristic].DefinitionCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicDefinitionCanonical(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionCanonical", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionCanonical", &resource.Characteristic[numCharacteristic].DefinitionCanonical, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicDefinitionExpression(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return ExpressionInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionExpression", nil, htmlAttrs)
	}
	return ExpressionInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionExpression", &resource.Characteristic[numCharacteristic].DefinitionExpression, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicDefinitionDataRequirement(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return DataRequirementInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionDataRequirement", nil, htmlAttrs)
	}
	return DataRequirementInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionDataRequirement", &resource.Characteristic[numCharacteristic].DefinitionDataRequirement, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicUsageContext(numCharacteristic int, numUsageContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numUsageContext >= len(resource.Characteristic[numCharacteristic].UsageContext) {
		return UsageContextInput("characteristic["+strconv.Itoa(numCharacteristic)+"].usageContext["+strconv.Itoa(numUsageContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("characteristic["+strconv.Itoa(numCharacteristic)+"].usageContext["+strconv.Itoa(numUsageContext)+"]", &resource.Characteristic[numCharacteristic].UsageContext[numUsageContext], htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicExclude(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].exclude", nil, htmlAttrs)
	}
	return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].exclude", resource.Characteristic[numCharacteristic].Exclude, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicUnitOfMeasure(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].unitOfMeasure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].unitOfMeasure", resource.Characteristic[numCharacteristic].UnitOfMeasure, optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicStudyEffectiveDescription(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].studyEffectiveDescription", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].studyEffectiveDescription", resource.Characteristic[numCharacteristic].StudyEffectiveDescription, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicStudyEffectiveDateTime(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return FhirDateTimeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].studyEffectiveDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].studyEffectiveDateTime", resource.Characteristic[numCharacteristic].StudyEffectiveDateTime, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicStudyEffectivePeriod(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return PeriodInput("characteristic["+strconv.Itoa(numCharacteristic)+"].studyEffectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("characteristic["+strconv.Itoa(numCharacteristic)+"].studyEffectivePeriod", resource.Characteristic[numCharacteristic].StudyEffectivePeriod, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicStudyEffectiveDuration(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return DurationInput("characteristic["+strconv.Itoa(numCharacteristic)+"].studyEffectiveDuration", nil, htmlAttrs)
	}
	return DurationInput("characteristic["+strconv.Itoa(numCharacteristic)+"].studyEffectiveDuration", resource.Characteristic[numCharacteristic].StudyEffectiveDuration, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicStudyEffectiveTiming(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return TimingInput("characteristic["+strconv.Itoa(numCharacteristic)+"].studyEffectiveTiming", nil, htmlAttrs)
	}
	return TimingInput("characteristic["+strconv.Itoa(numCharacteristic)+"].studyEffectiveTiming", resource.Characteristic[numCharacteristic].StudyEffectiveTiming, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicStudyEffectiveTimeFromStart(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return DurationInput("characteristic["+strconv.Itoa(numCharacteristic)+"].studyEffectiveTimeFromStart", nil, htmlAttrs)
	}
	return DurationInput("characteristic["+strconv.Itoa(numCharacteristic)+"].studyEffectiveTimeFromStart", resource.Characteristic[numCharacteristic].StudyEffectiveTimeFromStart, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicStudyEffectiveGroupMeasure(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSGroup_measure

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].studyEffectiveGroupMeasure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].studyEffectiveGroupMeasure", resource.Characteristic[numCharacteristic].StudyEffectiveGroupMeasure, optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicParticipantEffectiveDescription(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveDescription", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveDescription", resource.Characteristic[numCharacteristic].ParticipantEffectiveDescription, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicParticipantEffectiveDateTime(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return FhirDateTimeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveDateTime", resource.Characteristic[numCharacteristic].ParticipantEffectiveDateTime, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicParticipantEffectivePeriod(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return PeriodInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectivePeriod", resource.Characteristic[numCharacteristic].ParticipantEffectivePeriod, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicParticipantEffectiveDuration(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return DurationInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveDuration", nil, htmlAttrs)
	}
	return DurationInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveDuration", resource.Characteristic[numCharacteristic].ParticipantEffectiveDuration, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicParticipantEffectiveTiming(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return TimingInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveTiming", nil, htmlAttrs)
	}
	return TimingInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveTiming", resource.Characteristic[numCharacteristic].ParticipantEffectiveTiming, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicParticipantEffectiveTimeFromStart(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return DurationInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveTimeFromStart", nil, htmlAttrs)
	}
	return DurationInput("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveTimeFromStart", resource.Characteristic[numCharacteristic].ParticipantEffectiveTimeFromStart, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicParticipantEffectiveGroupMeasure(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSGroup_measure

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveGroupMeasure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].participantEffectiveGroupMeasure", resource.Characteristic[numCharacteristic].ParticipantEffectiveGroupMeasure, optionsValueSet, htmlAttrs)
}
