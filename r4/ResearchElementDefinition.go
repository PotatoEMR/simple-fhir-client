package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	SubjectCodeableConcept *CodeableConcept                          `json:"subjectCodeableConcept"`
	SubjectReference       *Reference                                `json:"subjectReference"`
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
	StudyEffectiveDateTime            *string          `json:"studyEffectiveDateTime"`
	StudyEffectivePeriod              *Period          `json:"studyEffectivePeriod"`
	StudyEffectiveDuration            *Duration        `json:"studyEffectiveDuration"`
	StudyEffectiveTiming              *Timing          `json:"studyEffectiveTiming"`
	StudyEffectiveTimeFromStart       *Duration        `json:"studyEffectiveTimeFromStart,omitempty"`
	StudyEffectiveGroupMeasure        *string          `json:"studyEffectiveGroupMeasure,omitempty"`
	ParticipantEffectiveDescription   *string          `json:"participantEffectiveDescription,omitempty"`
	ParticipantEffectiveDateTime      *string          `json:"participantEffectiveDateTime"`
	ParticipantEffectivePeriod        *Period          `json:"participantEffectivePeriod"`
	ParticipantEffectiveDuration      *Duration        `json:"participantEffectiveDuration"`
	ParticipantEffectiveTiming        *Timing          `json:"participantEffectiveTiming"`
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

func (resource *ResearchElementDefinition) ResearchElementDefinitionLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ResearchElementDefinition) ResearchElementDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ResearchElementDefinition) ResearchElementDefinitionJurisdiction(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *ResearchElementDefinition) ResearchElementDefinitionTopic(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("topic", nil, optionsValueSet)
	}
	return CodeableConceptSelect("topic", &resource.Topic[0], optionsValueSet)
}
func (resource *ResearchElementDefinition) ResearchElementDefinitionType() templ.Component {
	optionsValueSet := VSResearch_element_type

	if resource != nil {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", &resource.Type, optionsValueSet)
}
func (resource *ResearchElementDefinition) ResearchElementDefinitionVariableType() templ.Component {
	optionsValueSet := VSVariable_type

	if resource != nil {
		return CodeSelect("variableType", nil, optionsValueSet)
	}
	return CodeSelect("variableType", resource.VariableType, optionsValueSet)
}
func (resource *ResearchElementDefinition) ResearchElementDefinitionCharacteristicUnitOfMeasure(numCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Characteristic) >= numCharacteristic {
		return CodeableConceptSelect("unitOfMeasure", nil, optionsValueSet)
	}
	return CodeableConceptSelect("unitOfMeasure", resource.Characteristic[numCharacteristic].UnitOfMeasure, optionsValueSet)
}
func (resource *ResearchElementDefinition) ResearchElementDefinitionCharacteristicStudyEffectiveGroupMeasure(numCharacteristic int) templ.Component {
	optionsValueSet := VSGroup_measure

	if resource != nil && len(resource.Characteristic) >= numCharacteristic {
		return CodeSelect("studyEffectiveGroupMeasure", nil, optionsValueSet)
	}
	return CodeSelect("studyEffectiveGroupMeasure", resource.Characteristic[numCharacteristic].StudyEffectiveGroupMeasure, optionsValueSet)
}
func (resource *ResearchElementDefinition) ResearchElementDefinitionCharacteristicParticipantEffectiveGroupMeasure(numCharacteristic int) templ.Component {
	optionsValueSet := VSGroup_measure

	if resource != nil && len(resource.Characteristic) >= numCharacteristic {
		return CodeSelect("participantEffectiveGroupMeasure", nil, optionsValueSet)
	}
	return CodeSelect("participantEffectiveGroupMeasure", resource.Characteristic[numCharacteristic].ParticipantEffectiveGroupMeasure, optionsValueSet)
}
