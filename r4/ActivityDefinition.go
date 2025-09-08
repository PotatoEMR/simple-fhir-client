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

// http://hl7.org/fhir/r4/StructureDefinition/ActivityDefinition
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

// http://hl7.org/fhir/r4/StructureDefinition/ActivityDefinition
type ActivityDefinitionParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              string           `json:"type"`
	Role              *CodeableConcept `json:"role,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ActivityDefinition
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
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *ActivityDefinition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *ActivityDefinition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *ActivityDefinition) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *ActivityDefinition) T_Subtitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Subtitle", nil, htmlAttrs)
	}
	return StringInput("Subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *ActivityDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Experimental", nil, htmlAttrs)
	}
	return BoolInput("Experimental", resource.Experimental, htmlAttrs)
}
func (resource *ActivityDefinition) T_SubjectCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SubjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubjectCodeableConcept", resource.SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *ActivityDefinition) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *ActivityDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *ActivityDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Purpose", nil, htmlAttrs)
	}
	return StringInput("Purpose", resource.Purpose, htmlAttrs)
}
func (resource *ActivityDefinition) T_Usage(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Usage", nil, htmlAttrs)
	}
	return StringInput("Usage", resource.Usage, htmlAttrs)
}
func (resource *ActivityDefinition) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Copyright", nil, htmlAttrs)
	}
	return StringInput("Copyright", resource.Copyright, htmlAttrs)
}
func (resource *ActivityDefinition) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *ActivityDefinition) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *ActivityDefinition) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Library(numLibrary int, htmlAttrs string) templ.Component {
	if resource == nil || numLibrary >= len(resource.Library) {
		return StringInput("Library["+strconv.Itoa(numLibrary)+"]", nil, htmlAttrs)
	}
	return StringInput("Library["+strconv.Itoa(numLibrary)+"]", &resource.Library[numLibrary], htmlAttrs)
}
func (resource *ActivityDefinition) T_Kind(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_resource_types

	if resource == nil {
		return CodeSelect("Kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Kind", resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Profile(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Profile", nil, htmlAttrs)
	}
	return StringInput("Profile", resource.Profile, htmlAttrs)
}
func (resource *ActivityDefinition) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Intent(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("Intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Intent", resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Priority(htmlAttrs string) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_DoNotPerform(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("DoNotPerform", nil, htmlAttrs)
	}
	return BoolInput("DoNotPerform", resource.DoNotPerform, htmlAttrs)
}
func (resource *ActivityDefinition) T_TimingDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("TimingDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("TimingDateTime", resource.TimingDateTime, htmlAttrs)
}
func (resource *ActivityDefinition) T_ProductCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ProductCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ProductCodeableConcept", resource.ProductCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_BodySite(numBodySite int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBodySite >= len(resource.BodySite) {
		return CodeableConceptSelect("BodySite["+strconv.Itoa(numBodySite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("BodySite["+strconv.Itoa(numBodySite)+"]", &resource.BodySite[numBodySite], optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Transform(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Transform", nil, htmlAttrs)
	}
	return StringInput("Transform", resource.Transform, htmlAttrs)
}
func (resource *ActivityDefinition) T_ParticipantType(numParticipant int, htmlAttrs string) templ.Component {
	optionsValueSet := VSAction_participant_type

	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeSelect("Participant["+strconv.Itoa(numParticipant)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Participant["+strconv.Itoa(numParticipant)+"]Type", &resource.Participant[numParticipant].Type, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_ParticipantRole(numParticipant int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeableConceptSelect("Participant["+strconv.Itoa(numParticipant)+"]Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Participant["+strconv.Itoa(numParticipant)+"]Role", resource.Participant[numParticipant].Role, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_DynamicValuePath(numDynamicValue int, htmlAttrs string) templ.Component {
	if resource == nil || numDynamicValue >= len(resource.DynamicValue) {
		return StringInput("DynamicValue["+strconv.Itoa(numDynamicValue)+"]Path", nil, htmlAttrs)
	}
	return StringInput("DynamicValue["+strconv.Itoa(numDynamicValue)+"]Path", &resource.DynamicValue[numDynamicValue].Path, htmlAttrs)
}
