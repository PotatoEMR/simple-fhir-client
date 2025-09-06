package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/ResearchElementDefinition
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
	Date                   *string                                   `json:"date,omitempty"`
	Publisher              *string                                   `json:"publisher,omitempty"`
	Contact                []ContactDetail                           `json:"contact,omitempty"`
	Description            *string                                   `json:"description,omitempty"`
	Comment                []string                                  `json:"comment,omitempty"`
	UseContext             []UsageContext                            `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                         `json:"jurisdiction,omitempty"`
	Purpose                *string                                   `json:"purpose,omitempty"`
	Usage                  *string                                   `json:"usage,omitempty"`
	Copyright              *string                                   `json:"copyright,omitempty"`
	ApprovalDate           *string                                   `json:"approvalDate,omitempty"`
	LastReviewDate         *string                                   `json:"lastReviewDate,omitempty"`
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

// http://hl7.org/fhir/r4b/StructureDefinition/ResearchElementDefinition
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
	StudyEffectiveDateTime            *string          `json:"studyEffectiveDateTime,omitempty"`
	StudyEffectivePeriod              *Period          `json:"studyEffectivePeriod,omitempty"`
	StudyEffectiveDuration            *Duration        `json:"studyEffectiveDuration,omitempty"`
	StudyEffectiveTiming              *Timing          `json:"studyEffectiveTiming,omitempty"`
	StudyEffectiveTimeFromStart       *Duration        `json:"studyEffectiveTimeFromStart,omitempty"`
	StudyEffectiveGroupMeasure        *string          `json:"studyEffectiveGroupMeasure,omitempty"`
	ParticipantEffectiveDescription   *string          `json:"participantEffectiveDescription,omitempty"`
	ParticipantEffectiveDateTime      *string          `json:"participantEffectiveDateTime,omitempty"`
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

func (resource *ResearchElementDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ResearchElementDefinition.Id", nil)
	}
	return StringInput("ResearchElementDefinition.Id", resource.Id)
}
func (resource *ResearchElementDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ResearchElementDefinition.ImplicitRules", nil)
	}
	return StringInput("ResearchElementDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *ResearchElementDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ResearchElementDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("ResearchElementDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *ResearchElementDefinition) T_Url() templ.Component {

	if resource == nil {
		return StringInput("ResearchElementDefinition.Url", nil)
	}
	return StringInput("ResearchElementDefinition.Url", resource.Url)
}
func (resource *ResearchElementDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("ResearchElementDefinition.Version", nil)
	}
	return StringInput("ResearchElementDefinition.Version", resource.Version)
}
func (resource *ResearchElementDefinition) T_Name() templ.Component {

	if resource == nil {
		return StringInput("ResearchElementDefinition.Name", nil)
	}
	return StringInput("ResearchElementDefinition.Name", resource.Name)
}
func (resource *ResearchElementDefinition) T_Title() templ.Component {

	if resource == nil {
		return StringInput("ResearchElementDefinition.Title", nil)
	}
	return StringInput("ResearchElementDefinition.Title", resource.Title)
}
func (resource *ResearchElementDefinition) T_ShortTitle() templ.Component {

	if resource == nil {
		return StringInput("ResearchElementDefinition.ShortTitle", nil)
	}
	return StringInput("ResearchElementDefinition.ShortTitle", resource.ShortTitle)
}
func (resource *ResearchElementDefinition) T_Subtitle() templ.Component {

	if resource == nil {
		return StringInput("ResearchElementDefinition.Subtitle", nil)
	}
	return StringInput("ResearchElementDefinition.Subtitle", resource.Subtitle)
}
func (resource *ResearchElementDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ResearchElementDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("ResearchElementDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *ResearchElementDefinition) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("ResearchElementDefinition.Experimental", nil)
	}
	return BoolInput("ResearchElementDefinition.Experimental", resource.Experimental)
}
func (resource *ResearchElementDefinition) T_Date() templ.Component {

	if resource == nil {
		return StringInput("ResearchElementDefinition.Date", nil)
	}
	return StringInput("ResearchElementDefinition.Date", resource.Date)
}
func (resource *ResearchElementDefinition) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("ResearchElementDefinition.Publisher", nil)
	}
	return StringInput("ResearchElementDefinition.Publisher", resource.Publisher)
}
func (resource *ResearchElementDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("ResearchElementDefinition.Description", nil)
	}
	return StringInput("ResearchElementDefinition.Description", resource.Description)
}
func (resource *ResearchElementDefinition) T_Comment(numComment int) templ.Component {

	if resource == nil || len(resource.Comment) >= numComment {
		return StringInput("ResearchElementDefinition.Comment["+strconv.Itoa(numComment)+"]", nil)
	}
	return StringInput("ResearchElementDefinition.Comment["+strconv.Itoa(numComment)+"]", &resource.Comment[numComment])
}
func (resource *ResearchElementDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("ResearchElementDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchElementDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *ResearchElementDefinition) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("ResearchElementDefinition.Purpose", nil)
	}
	return StringInput("ResearchElementDefinition.Purpose", resource.Purpose)
}
func (resource *ResearchElementDefinition) T_Usage() templ.Component {

	if resource == nil {
		return StringInput("ResearchElementDefinition.Usage", nil)
	}
	return StringInput("ResearchElementDefinition.Usage", resource.Usage)
}
func (resource *ResearchElementDefinition) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("ResearchElementDefinition.Copyright", nil)
	}
	return StringInput("ResearchElementDefinition.Copyright", resource.Copyright)
}
func (resource *ResearchElementDefinition) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("ResearchElementDefinition.ApprovalDate", nil)
	}
	return StringInput("ResearchElementDefinition.ApprovalDate", resource.ApprovalDate)
}
func (resource *ResearchElementDefinition) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("ResearchElementDefinition.LastReviewDate", nil)
	}
	return StringInput("ResearchElementDefinition.LastReviewDate", resource.LastReviewDate)
}
func (resource *ResearchElementDefinition) T_Topic(numTopic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Topic) >= numTopic {
		return CodeableConceptSelect("ResearchElementDefinition.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchElementDefinition.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet)
}
func (resource *ResearchElementDefinition) T_Library(numLibrary int) templ.Component {

	if resource == nil || len(resource.Library) >= numLibrary {
		return StringInput("ResearchElementDefinition.Library["+strconv.Itoa(numLibrary)+"]", nil)
	}
	return StringInput("ResearchElementDefinition.Library["+strconv.Itoa(numLibrary)+"]", &resource.Library[numLibrary])
}
func (resource *ResearchElementDefinition) T_Type() templ.Component {
	optionsValueSet := VSResearch_element_type

	if resource == nil {
		return CodeSelect("ResearchElementDefinition.Type", nil, optionsValueSet)
	}
	return CodeSelect("ResearchElementDefinition.Type", &resource.Type, optionsValueSet)
}
func (resource *ResearchElementDefinition) T_VariableType() templ.Component {
	optionsValueSet := VSVariable_type

	if resource == nil {
		return CodeSelect("ResearchElementDefinition.VariableType", nil, optionsValueSet)
	}
	return CodeSelect("ResearchElementDefinition.VariableType", resource.VariableType, optionsValueSet)
}
func (resource *ResearchElementDefinition) T_CharacteristicId(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return StringInput("ResearchElementDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].Id", nil)
	}
	return StringInput("ResearchElementDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].Id", resource.Characteristic[numCharacteristic].Id)
}
func (resource *ResearchElementDefinition) T_CharacteristicExclude(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return BoolInput("ResearchElementDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].Exclude", nil)
	}
	return BoolInput("ResearchElementDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].Exclude", resource.Characteristic[numCharacteristic].Exclude)
}
func (resource *ResearchElementDefinition) T_CharacteristicUnitOfMeasure(numCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return CodeableConceptSelect("ResearchElementDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].UnitOfMeasure", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ResearchElementDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].UnitOfMeasure", resource.Characteristic[numCharacteristic].UnitOfMeasure, optionsValueSet)
}
func (resource *ResearchElementDefinition) T_CharacteristicStudyEffectiveDescription(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return StringInput("ResearchElementDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].StudyEffectiveDescription", nil)
	}
	return StringInput("ResearchElementDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].StudyEffectiveDescription", resource.Characteristic[numCharacteristic].StudyEffectiveDescription)
}
func (resource *ResearchElementDefinition) T_CharacteristicStudyEffectiveGroupMeasure(numCharacteristic int) templ.Component {
	optionsValueSet := VSGroup_measure

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return CodeSelect("ResearchElementDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].StudyEffectiveGroupMeasure", nil, optionsValueSet)
	}
	return CodeSelect("ResearchElementDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].StudyEffectiveGroupMeasure", resource.Characteristic[numCharacteristic].StudyEffectiveGroupMeasure, optionsValueSet)
}
func (resource *ResearchElementDefinition) T_CharacteristicParticipantEffectiveDescription(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return StringInput("ResearchElementDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].ParticipantEffectiveDescription", nil)
	}
	return StringInput("ResearchElementDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].ParticipantEffectiveDescription", resource.Characteristic[numCharacteristic].ParticipantEffectiveDescription)
}
func (resource *ResearchElementDefinition) T_CharacteristicParticipantEffectiveGroupMeasure(numCharacteristic int) templ.Component {
	optionsValueSet := VSGroup_measure

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return CodeSelect("ResearchElementDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].ParticipantEffectiveGroupMeasure", nil, optionsValueSet)
	}
	return CodeSelect("ResearchElementDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].ParticipantEffectiveGroupMeasure", resource.Characteristic[numCharacteristic].ParticipantEffectiveGroupMeasure, optionsValueSet)
}
