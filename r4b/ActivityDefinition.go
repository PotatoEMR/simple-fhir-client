package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/ActivityDefinition
type ActivityDefinition struct {
	Id                           *string                          `json:"id,omitempty"`
	Meta                         *Meta                            `json:"meta,omitempty"`
	ImplicitRules                *string                          `json:"implicitRules,omitempty"`
	Language                     *string                          `json:"language,omitempty"`
	Text                         *Narrative                       `json:"text,omitempty"`
	Contained                    []Resource                       `json:"contained,omitempty"`
	Extension                    []Extension                      `json:"extension,omitempty"`
	ModifierExtension            []Extension                      `json:"modifierExtension,omitempty"`
	Url                          *string                          `json:"url,omitempty"`
	Identifier                   []Identifier                     `json:"identifier,omitempty"`
	Version                      *string                          `json:"version,omitempty"`
	Name                         *string                          `json:"name,omitempty"`
	Title                        *string                          `json:"title,omitempty"`
	Subtitle                     *string                          `json:"subtitle,omitempty"`
	Status                       string                           `json:"status"`
	Experimental                 *bool                            `json:"experimental,omitempty"`
	SubjectCodeableConcept       *CodeableConcept                 `json:"subjectCodeableConcept,omitempty"`
	SubjectReference             *Reference                       `json:"subjectReference,omitempty"`
	SubjectCanonical             *string                          `json:"subjectCanonical,omitempty"`
	Date                         *string                          `json:"date,omitempty"`
	Publisher                    *string                          `json:"publisher,omitempty"`
	Contact                      []ContactDetail                  `json:"contact,omitempty"`
	Description                  *string                          `json:"description,omitempty"`
	UseContext                   []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction                 []CodeableConcept                `json:"jurisdiction,omitempty"`
	Purpose                      *string                          `json:"purpose,omitempty"`
	Usage                        *string                          `json:"usage,omitempty"`
	Copyright                    *string                          `json:"copyright,omitempty"`
	ApprovalDate                 *string                          `json:"approvalDate,omitempty"`
	LastReviewDate               *string                          `json:"lastReviewDate,omitempty"`
	EffectivePeriod              *Period                          `json:"effectivePeriod,omitempty"`
	Topic                        []CodeableConcept                `json:"topic,omitempty"`
	Author                       []ContactDetail                  `json:"author,omitempty"`
	Editor                       []ContactDetail                  `json:"editor,omitempty"`
	Reviewer                     []ContactDetail                  `json:"reviewer,omitempty"`
	Endorser                     []ContactDetail                  `json:"endorser,omitempty"`
	RelatedArtifact              []RelatedArtifact                `json:"relatedArtifact,omitempty"`
	Library                      []string                         `json:"library,omitempty"`
	Kind                         *string                          `json:"kind,omitempty"`
	Profile                      *string                          `json:"profile,omitempty"`
	Code                         *CodeableConcept                 `json:"code,omitempty"`
	Intent                       *string                          `json:"intent,omitempty"`
	Priority                     *string                          `json:"priority,omitempty"`
	DoNotPerform                 *bool                            `json:"doNotPerform,omitempty"`
	TimingTiming                 *Timing                          `json:"timingTiming,omitempty"`
	TimingDateTime               *string                          `json:"timingDateTime,omitempty"`
	TimingAge                    *Age                             `json:"timingAge,omitempty"`
	TimingPeriod                 *Period                          `json:"timingPeriod,omitempty"`
	TimingRange                  *Range                           `json:"timingRange,omitempty"`
	TimingDuration               *Duration                        `json:"timingDuration,omitempty"`
	Location                     *Reference                       `json:"location,omitempty"`
	Participant                  []ActivityDefinitionParticipant  `json:"participant,omitempty"`
	ProductReference             *Reference                       `json:"productReference,omitempty"`
	ProductCodeableConcept       *CodeableConcept                 `json:"productCodeableConcept,omitempty"`
	Quantity                     *Quantity                        `json:"quantity,omitempty"`
	Dosage                       []Dosage                         `json:"dosage,omitempty"`
	BodySite                     []CodeableConcept                `json:"bodySite,omitempty"`
	SpecimenRequirement          []Reference                      `json:"specimenRequirement,omitempty"`
	ObservationRequirement       []Reference                      `json:"observationRequirement,omitempty"`
	ObservationResultRequirement []Reference                      `json:"observationResultRequirement,omitempty"`
	Transform                    *string                          `json:"transform,omitempty"`
	DynamicValue                 []ActivityDefinitionDynamicValue `json:"dynamicValue,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ActivityDefinition
type ActivityDefinitionParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              string           `json:"type"`
	Role              *CodeableConcept `json:"role,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ActivityDefinition
type ActivityDefinitionDynamicValue struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Path              string      `json:"path"`
	Expression        Expression  `json:"expression"`
}

type OtherActivityDefinition ActivityDefinition

// on convert struct to json, automatically add resourceType=ActivityDefinition
func (r ActivityDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherActivityDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherActivityDefinition: OtherActivityDefinition(r),
		ResourceType:            "ActivityDefinition",
	})
}

func (resource *ActivityDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.Id", nil)
	}
	return StringInput("ActivityDefinition.Id", resource.Id)
}
func (resource *ActivityDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.ImplicitRules", nil)
	}
	return StringInput("ActivityDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *ActivityDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ActivityDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("ActivityDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *ActivityDefinition) T_Url() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.Url", nil)
	}
	return StringInput("ActivityDefinition.Url", resource.Url)
}
func (resource *ActivityDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.Version", nil)
	}
	return StringInput("ActivityDefinition.Version", resource.Version)
}
func (resource *ActivityDefinition) T_Name() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.Name", nil)
	}
	return StringInput("ActivityDefinition.Name", resource.Name)
}
func (resource *ActivityDefinition) T_Title() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.Title", nil)
	}
	return StringInput("ActivityDefinition.Title", resource.Title)
}
func (resource *ActivityDefinition) T_Subtitle() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.Subtitle", nil)
	}
	return StringInput("ActivityDefinition.Subtitle", resource.Subtitle)
}
func (resource *ActivityDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ActivityDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("ActivityDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *ActivityDefinition) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("ActivityDefinition.Experimental", nil)
	}
	return BoolInput("ActivityDefinition.Experimental", resource.Experimental)
}
func (resource *ActivityDefinition) T_Date() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.Date", nil)
	}
	return StringInput("ActivityDefinition.Date", resource.Date)
}
func (resource *ActivityDefinition) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.Publisher", nil)
	}
	return StringInput("ActivityDefinition.Publisher", resource.Publisher)
}
func (resource *ActivityDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.Description", nil)
	}
	return StringInput("ActivityDefinition.Description", resource.Description)
}
func (resource *ActivityDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("ActivityDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ActivityDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *ActivityDefinition) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.Purpose", nil)
	}
	return StringInput("ActivityDefinition.Purpose", resource.Purpose)
}
func (resource *ActivityDefinition) T_Usage() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.Usage", nil)
	}
	return StringInput("ActivityDefinition.Usage", resource.Usage)
}
func (resource *ActivityDefinition) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.Copyright", nil)
	}
	return StringInput("ActivityDefinition.Copyright", resource.Copyright)
}
func (resource *ActivityDefinition) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.ApprovalDate", nil)
	}
	return StringInput("ActivityDefinition.ApprovalDate", resource.ApprovalDate)
}
func (resource *ActivityDefinition) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.LastReviewDate", nil)
	}
	return StringInput("ActivityDefinition.LastReviewDate", resource.LastReviewDate)
}
func (resource *ActivityDefinition) T_Topic(numTopic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Topic) >= numTopic {
		return CodeableConceptSelect("ActivityDefinition.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ActivityDefinition.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet)
}
func (resource *ActivityDefinition) T_Library(numLibrary int) templ.Component {

	if resource == nil || len(resource.Library) >= numLibrary {
		return StringInput("ActivityDefinition.Library["+strconv.Itoa(numLibrary)+"]", nil)
	}
	return StringInput("ActivityDefinition.Library["+strconv.Itoa(numLibrary)+"]", &resource.Library[numLibrary])
}
func (resource *ActivityDefinition) T_Kind() templ.Component {
	optionsValueSet := VSRequest_resource_types

	if resource == nil {
		return CodeSelect("ActivityDefinition.Kind", nil, optionsValueSet)
	}
	return CodeSelect("ActivityDefinition.Kind", resource.Kind, optionsValueSet)
}
func (resource *ActivityDefinition) T_Profile() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.Profile", nil)
	}
	return StringInput("ActivityDefinition.Profile", resource.Profile)
}
func (resource *ActivityDefinition) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ActivityDefinition.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ActivityDefinition.Code", resource.Code, optionsValueSet)
}
func (resource *ActivityDefinition) T_Intent() templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("ActivityDefinition.Intent", nil, optionsValueSet)
	}
	return CodeSelect("ActivityDefinition.Intent", resource.Intent, optionsValueSet)
}
func (resource *ActivityDefinition) T_Priority() templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("ActivityDefinition.Priority", nil, optionsValueSet)
	}
	return CodeSelect("ActivityDefinition.Priority", resource.Priority, optionsValueSet)
}
func (resource *ActivityDefinition) T_DoNotPerform() templ.Component {

	if resource == nil {
		return BoolInput("ActivityDefinition.DoNotPerform", nil)
	}
	return BoolInput("ActivityDefinition.DoNotPerform", resource.DoNotPerform)
}
func (resource *ActivityDefinition) T_BodySite(numBodySite int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.BodySite) >= numBodySite {
		return CodeableConceptSelect("ActivityDefinition.BodySite["+strconv.Itoa(numBodySite)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ActivityDefinition.BodySite["+strconv.Itoa(numBodySite)+"]", &resource.BodySite[numBodySite], optionsValueSet)
}
func (resource *ActivityDefinition) T_Transform() templ.Component {

	if resource == nil {
		return StringInput("ActivityDefinition.Transform", nil)
	}
	return StringInput("ActivityDefinition.Transform", resource.Transform)
}
func (resource *ActivityDefinition) T_ParticipantId(numParticipant int) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return StringInput("ActivityDefinition.Participant["+strconv.Itoa(numParticipant)+"].Id", nil)
	}
	return StringInput("ActivityDefinition.Participant["+strconv.Itoa(numParticipant)+"].Id", resource.Participant[numParticipant].Id)
}
func (resource *ActivityDefinition) T_ParticipantType(numParticipant int) templ.Component {
	optionsValueSet := VSAction_participant_type

	if resource == nil || len(resource.Participant) >= numParticipant {
		return CodeSelect("ActivityDefinition.Participant["+strconv.Itoa(numParticipant)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("ActivityDefinition.Participant["+strconv.Itoa(numParticipant)+"].Type", &resource.Participant[numParticipant].Type, optionsValueSet)
}
func (resource *ActivityDefinition) T_ParticipantRole(numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Participant) >= numParticipant {
		return CodeableConceptSelect("ActivityDefinition.Participant["+strconv.Itoa(numParticipant)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ActivityDefinition.Participant["+strconv.Itoa(numParticipant)+"].Role", resource.Participant[numParticipant].Role, optionsValueSet)
}
func (resource *ActivityDefinition) T_DynamicValueId(numDynamicValue int) templ.Component {

	if resource == nil || len(resource.DynamicValue) >= numDynamicValue {
		return StringInput("ActivityDefinition.DynamicValue["+strconv.Itoa(numDynamicValue)+"].Id", nil)
	}
	return StringInput("ActivityDefinition.DynamicValue["+strconv.Itoa(numDynamicValue)+"].Id", resource.DynamicValue[numDynamicValue].Id)
}
func (resource *ActivityDefinition) T_DynamicValuePath(numDynamicValue int) templ.Component {

	if resource == nil || len(resource.DynamicValue) >= numDynamicValue {
		return StringInput("ActivityDefinition.DynamicValue["+strconv.Itoa(numDynamicValue)+"].Path", nil)
	}
	return StringInput("ActivityDefinition.DynamicValue["+strconv.Itoa(numDynamicValue)+"].Path", &resource.DynamicValue[numDynamicValue].Path)
}
