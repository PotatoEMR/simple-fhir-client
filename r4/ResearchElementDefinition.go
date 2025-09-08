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
	Date                   *time.Time                                `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher              *string                                   `json:"publisher,omitempty"`
	Contact                []ContactDetail                           `json:"contact,omitempty"`
	Description            *string                                   `json:"description,omitempty"`
	Comment                []string                                  `json:"comment,omitempty"`
	UseContext             []UsageContext                            `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                         `json:"jurisdiction,omitempty"`
	Purpose                *string                                   `json:"purpose,omitempty"`
	Usage                  *string                                   `json:"usage,omitempty"`
	Copyright              *string                                   `json:"copyright,omitempty"`
	ApprovalDate           *time.Time                                `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate         *time.Time                                `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
	StudyEffectiveDateTime            *time.Time       `json:"studyEffectiveDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	StudyEffectivePeriod              *Period          `json:"studyEffectivePeriod,omitempty"`
	StudyEffectiveDuration            *Duration        `json:"studyEffectiveDuration,omitempty"`
	StudyEffectiveTiming              *Timing          `json:"studyEffectiveTiming,omitempty"`
	StudyEffectiveTimeFromStart       *Duration        `json:"studyEffectiveTimeFromStart,omitempty"`
	StudyEffectiveGroupMeasure        *string          `json:"studyEffectiveGroupMeasure,omitempty"`
	ParticipantEffectiveDescription   *string          `json:"participantEffectiveDescription,omitempty"`
	ParticipantEffectiveDateTime      *time.Time       `json:"participantEffectiveDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *ResearchElementDefinition) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_ShortTitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ShortTitle", nil, htmlAttrs)
	}
	return StringInput("ShortTitle", resource.ShortTitle, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Subtitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Subtitle", nil, htmlAttrs)
	}
	return StringInput("Subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Experimental", nil, htmlAttrs)
	}
	return BoolInput("Experimental", resource.Experimental, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_SubjectCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SubjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubjectCodeableConcept", resource.SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Comment(numComment int, htmlAttrs string) templ.Component {
	if resource == nil || numComment >= len(resource.Comment) {
		return StringInput("Comment["+strconv.Itoa(numComment)+"]", nil, htmlAttrs)
	}
	return StringInput("Comment["+strconv.Itoa(numComment)+"]", &resource.Comment[numComment], htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Purpose", nil, htmlAttrs)
	}
	return StringInput("Purpose", resource.Purpose, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Usage(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Usage", nil, htmlAttrs)
	}
	return StringInput("Usage", resource.Usage, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Copyright", nil, htmlAttrs)
	}
	return StringInput("Copyright", resource.Copyright, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Library(numLibrary int, htmlAttrs string) templ.Component {
	if resource == nil || numLibrary >= len(resource.Library) {
		return StringInput("Library["+strconv.Itoa(numLibrary)+"]", nil, htmlAttrs)
	}
	return StringInput("Library["+strconv.Itoa(numLibrary)+"]", &resource.Library[numLibrary], htmlAttrs)
}
func (resource *ResearchElementDefinition) T_Type(htmlAttrs string) templ.Component {
	optionsValueSet := VSResearch_element_type

	if resource == nil {
		return CodeSelect("Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_VariableType(htmlAttrs string) templ.Component {
	optionsValueSet := VSVariable_type

	if resource == nil {
		return CodeSelect("VariableType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("VariableType", resource.VariableType, optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicDefinitionCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionCodeableConcept", &resource.Characteristic[numCharacteristic].DefinitionCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicDefinitionCanonical(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionCanonical", nil, htmlAttrs)
	}
	return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionCanonical", &resource.Characteristic[numCharacteristic].DefinitionCanonical, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicExclude(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]Exclude", nil, htmlAttrs)
	}
	return BoolInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]Exclude", resource.Characteristic[numCharacteristic].Exclude, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicUnitOfMeasure(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]UnitOfMeasure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]UnitOfMeasure", resource.Characteristic[numCharacteristic].UnitOfMeasure, optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicStudyEffectiveDescription(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]StudyEffectiveDescription", nil, htmlAttrs)
	}
	return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]StudyEffectiveDescription", resource.Characteristic[numCharacteristic].StudyEffectiveDescription, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicStudyEffectiveDateTime(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return DateTimeInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]StudyEffectiveDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]StudyEffectiveDateTime", resource.Characteristic[numCharacteristic].StudyEffectiveDateTime, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicStudyEffectiveGroupMeasure(numCharacteristic int, htmlAttrs string) templ.Component {
	optionsValueSet := VSGroup_measure

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]StudyEffectiveGroupMeasure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]StudyEffectiveGroupMeasure", resource.Characteristic[numCharacteristic].StudyEffectiveGroupMeasure, optionsValueSet, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicParticipantEffectiveDescription(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]ParticipantEffectiveDescription", nil, htmlAttrs)
	}
	return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]ParticipantEffectiveDescription", resource.Characteristic[numCharacteristic].ParticipantEffectiveDescription, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicParticipantEffectiveDateTime(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return DateTimeInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]ParticipantEffectiveDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]ParticipantEffectiveDateTime", resource.Characteristic[numCharacteristic].ParticipantEffectiveDateTime, htmlAttrs)
}
func (resource *ResearchElementDefinition) T_CharacteristicParticipantEffectiveGroupMeasure(numCharacteristic int, htmlAttrs string) templ.Component {
	optionsValueSet := VSGroup_measure

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]ParticipantEffectiveGroupMeasure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]ParticipantEffectiveGroupMeasure", resource.Characteristic[numCharacteristic].ParticipantEffectiveGroupMeasure, optionsValueSet, htmlAttrs)
}
