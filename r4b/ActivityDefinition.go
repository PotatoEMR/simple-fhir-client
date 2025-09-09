package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Date                         *time.Time                       `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher                    *string                          `json:"publisher,omitempty"`
	Contact                      []ContactDetail                  `json:"contact,omitempty"`
	Description                  *string                          `json:"description,omitempty"`
	UseContext                   []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction                 []CodeableConcept                `json:"jurisdiction,omitempty"`
	Purpose                      *string                          `json:"purpose,omitempty"`
	Usage                        *string                          `json:"usage,omitempty"`
	Copyright                    *string                          `json:"copyright,omitempty"`
	ApprovalDate                 *time.Time                       `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate               *time.Time                       `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
	TimingDateTime               *time.Time                       `json:"timingDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r ActivityDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ActivityDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ActivityDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ActivityDefinition) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActivityDefinition.Url", nil, htmlAttrs)
	}
	return StringInput("ActivityDefinition.Url", resource.Url, htmlAttrs)
}
func (resource *ActivityDefinition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActivityDefinition.Version", nil, htmlAttrs)
	}
	return StringInput("ActivityDefinition.Version", resource.Version, htmlAttrs)
}
func (resource *ActivityDefinition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActivityDefinition.Name", nil, htmlAttrs)
	}
	return StringInput("ActivityDefinition.Name", resource.Name, htmlAttrs)
}
func (resource *ActivityDefinition) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActivityDefinition.Title", nil, htmlAttrs)
	}
	return StringInput("ActivityDefinition.Title", resource.Title, htmlAttrs)
}
func (resource *ActivityDefinition) T_Subtitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActivityDefinition.Subtitle", nil, htmlAttrs)
	}
	return StringInput("ActivityDefinition.Subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *ActivityDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ActivityDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ActivityDefinition.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("ActivityDefinition.Experimental", nil, htmlAttrs)
	}
	return BoolInput("ActivityDefinition.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *ActivityDefinition) T_SubjectCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ActivityDefinition.SubjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ActivityDefinition.SubjectCodeableConcept", resource.SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_SubjectCanonical(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActivityDefinition.SubjectCanonical", nil, htmlAttrs)
	}
	return StringInput("ActivityDefinition.SubjectCanonical", resource.SubjectCanonical, htmlAttrs)
}
func (resource *ActivityDefinition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("ActivityDefinition.Date", nil, htmlAttrs)
	}
	return DateTimeInput("ActivityDefinition.Date", resource.Date, htmlAttrs)
}
func (resource *ActivityDefinition) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActivityDefinition.Publisher", nil, htmlAttrs)
	}
	return StringInput("ActivityDefinition.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *ActivityDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActivityDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("ActivityDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *ActivityDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("ActivityDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ActivityDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActivityDefinition.Purpose", nil, htmlAttrs)
	}
	return StringInput("ActivityDefinition.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *ActivityDefinition) T_Usage(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActivityDefinition.Usage", nil, htmlAttrs)
	}
	return StringInput("ActivityDefinition.Usage", resource.Usage, htmlAttrs)
}
func (resource *ActivityDefinition) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActivityDefinition.Copyright", nil, htmlAttrs)
	}
	return StringInput("ActivityDefinition.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *ActivityDefinition) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ActivityDefinition.ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("ActivityDefinition.ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *ActivityDefinition) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ActivityDefinition.LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("ActivityDefinition.LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *ActivityDefinition) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("ActivityDefinition.Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ActivityDefinition.Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Library(numLibrary int, htmlAttrs string) templ.Component {
	if resource == nil || numLibrary >= len(resource.Library) {
		return StringInput("ActivityDefinition.Library["+strconv.Itoa(numLibrary)+"]", nil, htmlAttrs)
	}
	return StringInput("ActivityDefinition.Library["+strconv.Itoa(numLibrary)+"]", &resource.Library[numLibrary], htmlAttrs)
}
func (resource *ActivityDefinition) T_Kind(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_resource_types

	if resource == nil {
		return CodeSelect("ActivityDefinition.Kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ActivityDefinition.Kind", resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Profile(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActivityDefinition.Profile", nil, htmlAttrs)
	}
	return StringInput("ActivityDefinition.Profile", resource.Profile, htmlAttrs)
}
func (resource *ActivityDefinition) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ActivityDefinition.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ActivityDefinition.Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Intent(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("ActivityDefinition.Intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ActivityDefinition.Intent", resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Priority(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("ActivityDefinition.Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ActivityDefinition.Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_DoNotPerform(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("ActivityDefinition.DoNotPerform", nil, htmlAttrs)
	}
	return BoolInput("ActivityDefinition.DoNotPerform", resource.DoNotPerform, htmlAttrs)
}
func (resource *ActivityDefinition) T_TimingDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("ActivityDefinition.TimingDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("ActivityDefinition.TimingDateTime", resource.TimingDateTime, htmlAttrs)
}
func (resource *ActivityDefinition) T_ProductCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ActivityDefinition.ProductCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ActivityDefinition.ProductCodeableConcept", resource.ProductCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_BodySite(numBodySite int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBodySite >= len(resource.BodySite) {
		return CodeableConceptSelect("ActivityDefinition.BodySite["+strconv.Itoa(numBodySite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ActivityDefinition.BodySite["+strconv.Itoa(numBodySite)+"]", &resource.BodySite[numBodySite], optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Transform(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActivityDefinition.Transform", nil, htmlAttrs)
	}
	return StringInput("ActivityDefinition.Transform", resource.Transform, htmlAttrs)
}
func (resource *ActivityDefinition) T_ParticipantType(numParticipant int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_participant_type

	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeSelect("ActivityDefinition.Participant["+strconv.Itoa(numParticipant)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ActivityDefinition.Participant["+strconv.Itoa(numParticipant)+"].Type", &resource.Participant[numParticipant].Type, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_ParticipantRole(numParticipant int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeableConceptSelect("ActivityDefinition.Participant["+strconv.Itoa(numParticipant)+"].Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ActivityDefinition.Participant["+strconv.Itoa(numParticipant)+"].Role", resource.Participant[numParticipant].Role, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_DynamicValuePath(numDynamicValue int, htmlAttrs string) templ.Component {
	if resource == nil || numDynamicValue >= len(resource.DynamicValue) {
		return StringInput("ActivityDefinition.DynamicValue["+strconv.Itoa(numDynamicValue)+"].Path", nil, htmlAttrs)
	}
	return StringInput("ActivityDefinition.DynamicValue["+strconv.Itoa(numDynamicValue)+"].Path", &resource.DynamicValue[numDynamicValue].Path, htmlAttrs)
}
