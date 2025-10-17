package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

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
	Date                         *FhirDateTime                    `json:"date,omitempty"`
	Publisher                    *string                          `json:"publisher,omitempty"`
	Contact                      []ContactDetail                  `json:"contact,omitempty"`
	Description                  *string                          `json:"description,omitempty"`
	UseContext                   []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction                 []CodeableConcept                `json:"jurisdiction,omitempty"`
	Purpose                      *string                          `json:"purpose,omitempty"`
	Usage                        *string                          `json:"usage,omitempty"`
	Copyright                    *string                          `json:"copyright,omitempty"`
	ApprovalDate                 *FhirDate                        `json:"approvalDate,omitempty"`
	LastReviewDate               *FhirDate                        `json:"lastReviewDate,omitempty"`
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
	TimingDateTime               *FhirDateTime                    `json:"timingDateTime,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
func (r ActivityDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherActivityDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherActivityDefinition: OtherActivityDefinition(r),
		ResourceType:            "ActivityDefinition",
	})
}

// json -> struct, first reject if resourceType != ActivityDefinition
func (r *ActivityDefinition) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "ActivityDefinition" {
		return errors.New("resourceType not ActivityDefinition")
	}
	return json.Unmarshal(data, (*OtherActivityDefinition)(r))
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
func (resource *ActivityDefinition) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *ActivityDefinition) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *ActivityDefinition) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *ActivityDefinition) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *ActivityDefinition) T_Subtitle(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("subtitle", nil, htmlAttrs)
	}
	return StringInput("subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *ActivityDefinition) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *ActivityDefinition) T_SubjectCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("subjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subjectCodeableConcept", resource.SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_SubjectReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subjectReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subjectReference", resource.SubjectReference, htmlAttrs)
}
func (resource *ActivityDefinition) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *ActivityDefinition) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *ActivityDefinition) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *ActivityDefinition) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *ActivityDefinition) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *ActivityDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *ActivityDefinition) T_Usage(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("usage", nil, htmlAttrs)
	}
	return StringInput("usage", resource.Usage, htmlAttrs)
}
func (resource *ActivityDefinition) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *ActivityDefinition) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("approvalDate", nil, htmlAttrs)
	}
	return FhirDateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *ActivityDefinition) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("lastReviewDate", nil, htmlAttrs)
	}
	return FhirDateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *ActivityDefinition) T_EffectivePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("effectivePeriod", resource.EffectivePeriod, htmlAttrs)
}
func (resource *ActivityDefinition) T_Topic(numTopic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTopic >= len(resource.Topic) {
		return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("topic["+strconv.Itoa(numTopic)+"]", &resource.Topic[numTopic], optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Author(numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthor >= len(resource.Author) {
		return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", &resource.Author[numAuthor], htmlAttrs)
}
func (resource *ActivityDefinition) T_Editor(numEditor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEditor >= len(resource.Editor) {
		return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", &resource.Editor[numEditor], htmlAttrs)
}
func (resource *ActivityDefinition) T_Reviewer(numReviewer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReviewer >= len(resource.Reviewer) {
		return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", &resource.Reviewer[numReviewer], htmlAttrs)
}
func (resource *ActivityDefinition) T_Endorser(numEndorser int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndorser >= len(resource.Endorser) {
		return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", &resource.Endorser[numEndorser], htmlAttrs)
}
func (resource *ActivityDefinition) T_RelatedArtifact(numRelatedArtifact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedArtifact >= len(resource.RelatedArtifact) {
		return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", &resource.RelatedArtifact[numRelatedArtifact], htmlAttrs)
}
func (resource *ActivityDefinition) T_Library(numLibrary int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLibrary >= len(resource.Library) {
		return StringInput("library["+strconv.Itoa(numLibrary)+"]", nil, htmlAttrs)
	}
	return StringInput("library["+strconv.Itoa(numLibrary)+"]", &resource.Library[numLibrary], htmlAttrs)
}
func (resource *ActivityDefinition) T_Kind(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_resource_types

	if resource == nil {
		return CodeSelect("kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("kind", resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Profile(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("profile", nil, htmlAttrs)
	}
	return StringInput("profile", resource.Profile, htmlAttrs)
}
func (resource *ActivityDefinition) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Intent(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_intent

	if resource == nil {
		return CodeSelect("intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("intent", resource.Intent, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Priority(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRequest_priority

	if resource == nil {
		return CodeSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_DoNotPerform(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("doNotPerform", nil, htmlAttrs)
	}
	return BoolInput("doNotPerform", resource.DoNotPerform, htmlAttrs)
}
func (resource *ActivityDefinition) T_TimingTiming(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return TimingInput("timingTiming", nil, htmlAttrs)
	}
	return TimingInput("timingTiming", resource.TimingTiming, htmlAttrs)
}
func (resource *ActivityDefinition) T_TimingDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("timingDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("timingDateTime", resource.TimingDateTime, htmlAttrs)
}
func (resource *ActivityDefinition) T_TimingAge(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return AgeInput("timingAge", nil, htmlAttrs)
	}
	return AgeInput("timingAge", resource.TimingAge, htmlAttrs)
}
func (resource *ActivityDefinition) T_TimingPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("timingPeriod", nil, htmlAttrs)
	}
	return PeriodInput("timingPeriod", resource.TimingPeriod, htmlAttrs)
}
func (resource *ActivityDefinition) T_TimingRange(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return RangeInput("timingRange", nil, htmlAttrs)
	}
	return RangeInput("timingRange", resource.TimingRange, htmlAttrs)
}
func (resource *ActivityDefinition) T_TimingDuration(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DurationInput("timingDuration", nil, htmlAttrs)
	}
	return DurationInput("timingDuration", resource.TimingDuration, htmlAttrs)
}
func (resource *ActivityDefinition) T_Location(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "location", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "location", resource.Location, htmlAttrs)
}
func (resource *ActivityDefinition) T_ProductReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "productReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "productReference", resource.ProductReference, htmlAttrs)
}
func (resource *ActivityDefinition) T_ProductCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("productCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("productCodeableConcept", resource.ProductCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Quantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("quantity", resource.Quantity, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_Dosage(numDosage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDosage >= len(resource.Dosage) {
		return DosageInput("dosage["+strconv.Itoa(numDosage)+"]", nil, htmlAttrs)
	}
	return DosageInput("dosage["+strconv.Itoa(numDosage)+"]", &resource.Dosage[numDosage], htmlAttrs)
}
func (resource *ActivityDefinition) T_BodySite(numBodySite int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBodySite >= len(resource.BodySite) {
		return CodeableConceptSelect("bodySite["+strconv.Itoa(numBodySite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("bodySite["+strconv.Itoa(numBodySite)+"]", &resource.BodySite[numBodySite], optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_SpecimenRequirement(frs []FhirResource, numSpecimenRequirement int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecimenRequirement >= len(resource.SpecimenRequirement) {
		return ReferenceInput(frs, "specimenRequirement["+strconv.Itoa(numSpecimenRequirement)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "specimenRequirement["+strconv.Itoa(numSpecimenRequirement)+"]", &resource.SpecimenRequirement[numSpecimenRequirement], htmlAttrs)
}
func (resource *ActivityDefinition) T_ObservationRequirement(frs []FhirResource, numObservationRequirement int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numObservationRequirement >= len(resource.ObservationRequirement) {
		return ReferenceInput(frs, "observationRequirement["+strconv.Itoa(numObservationRequirement)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "observationRequirement["+strconv.Itoa(numObservationRequirement)+"]", &resource.ObservationRequirement[numObservationRequirement], htmlAttrs)
}
func (resource *ActivityDefinition) T_ObservationResultRequirement(frs []FhirResource, numObservationResultRequirement int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numObservationResultRequirement >= len(resource.ObservationResultRequirement) {
		return ReferenceInput(frs, "observationResultRequirement["+strconv.Itoa(numObservationResultRequirement)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "observationResultRequirement["+strconv.Itoa(numObservationResultRequirement)+"]", &resource.ObservationResultRequirement[numObservationResultRequirement], htmlAttrs)
}
func (resource *ActivityDefinition) T_Transform(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("transform", nil, htmlAttrs)
	}
	return StringInput("transform", resource.Transform, htmlAttrs)
}
func (resource *ActivityDefinition) T_ParticipantType(numParticipant int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAction_participant_type

	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeSelect("participant["+strconv.Itoa(numParticipant)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("participant["+strconv.Itoa(numParticipant)+"].type", &resource.Participant[numParticipant].Type, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_ParticipantRole(numParticipant int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numParticipant >= len(resource.Participant) {
		return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("participant["+strconv.Itoa(numParticipant)+"].role", resource.Participant[numParticipant].Role, optionsValueSet, htmlAttrs)
}
func (resource *ActivityDefinition) T_DynamicValuePath(numDynamicValue int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDynamicValue >= len(resource.DynamicValue) {
		return StringInput("dynamicValue["+strconv.Itoa(numDynamicValue)+"].path", nil, htmlAttrs)
	}
	return StringInput("dynamicValue["+strconv.Itoa(numDynamicValue)+"].path", &resource.DynamicValue[numDynamicValue].Path, htmlAttrs)
}
func (resource *ActivityDefinition) T_DynamicValueExpression(numDynamicValue int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDynamicValue >= len(resource.DynamicValue) {
		return ExpressionInput("dynamicValue["+strconv.Itoa(numDynamicValue)+"].expression", nil, htmlAttrs)
	}
	return ExpressionInput("dynamicValue["+strconv.Itoa(numDynamicValue)+"].expression", &resource.DynamicValue[numDynamicValue].Expression, htmlAttrs)
}
