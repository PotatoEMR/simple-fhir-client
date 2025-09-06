package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Immunization
type Immunization struct {
	Id                 *string                       `json:"id,omitempty"`
	Meta               *Meta                         `json:"meta,omitempty"`
	ImplicitRules      *string                       `json:"implicitRules,omitempty"`
	Language           *string                       `json:"language,omitempty"`
	Text               *Narrative                    `json:"text,omitempty"`
	Contained          []Resource                    `json:"contained,omitempty"`
	Extension          []Extension                   `json:"extension,omitempty"`
	ModifierExtension  []Extension                   `json:"modifierExtension,omitempty"`
	Identifier         []Identifier                  `json:"identifier,omitempty"`
	Status             string                        `json:"status"`
	StatusReason       *CodeableConcept              `json:"statusReason,omitempty"`
	VaccineCode        CodeableConcept               `json:"vaccineCode"`
	Patient            Reference                     `json:"patient"`
	Encounter          *Reference                    `json:"encounter,omitempty"`
	OccurrenceDateTime string                        `json:"occurrenceDateTime"`
	OccurrenceString   string                        `json:"occurrenceString"`
	Recorded           *string                       `json:"recorded,omitempty"`
	PrimarySource      *bool                         `json:"primarySource,omitempty"`
	ReportOrigin       *CodeableConcept              `json:"reportOrigin,omitempty"`
	Location           *Reference                    `json:"location,omitempty"`
	Manufacturer       *Reference                    `json:"manufacturer,omitempty"`
	LotNumber          *string                       `json:"lotNumber,omitempty"`
	ExpirationDate     *string                       `json:"expirationDate,omitempty"`
	Site               *CodeableConcept              `json:"site,omitempty"`
	Route              *CodeableConcept              `json:"route,omitempty"`
	DoseQuantity       *Quantity                     `json:"doseQuantity,omitempty"`
	Performer          []ImmunizationPerformer       `json:"performer,omitempty"`
	Note               []Annotation                  `json:"note,omitempty"`
	ReasonCode         []CodeableConcept             `json:"reasonCode,omitempty"`
	ReasonReference    []Reference                   `json:"reasonReference,omitempty"`
	IsSubpotent        *bool                         `json:"isSubpotent,omitempty"`
	SubpotentReason    []CodeableConcept             `json:"subpotentReason,omitempty"`
	Education          []ImmunizationEducation       `json:"education,omitempty"`
	ProgramEligibility []CodeableConcept             `json:"programEligibility,omitempty"`
	FundingSource      *CodeableConcept              `json:"fundingSource,omitempty"`
	Reaction           []ImmunizationReaction        `json:"reaction,omitempty"`
	ProtocolApplied    []ImmunizationProtocolApplied `json:"protocolApplied,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Immunization
type ImmunizationPerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Immunization
type ImmunizationEducation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	DocumentType      *string     `json:"documentType,omitempty"`
	Reference         *string     `json:"reference,omitempty"`
	PublicationDate   *string     `json:"publicationDate,omitempty"`
	PresentationDate  *string     `json:"presentationDate,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Immunization
type ImmunizationReaction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Date              *string     `json:"date,omitempty"`
	Detail            *Reference  `json:"detail,omitempty"`
	Reported          *bool       `json:"reported,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Immunization
type ImmunizationProtocolApplied struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Series                 *string           `json:"series,omitempty"`
	Authority              *Reference        `json:"authority,omitempty"`
	TargetDisease          []CodeableConcept `json:"targetDisease,omitempty"`
	DoseNumberPositiveInt  int               `json:"doseNumberPositiveInt"`
	DoseNumberString       string            `json:"doseNumberString"`
	SeriesDosesPositiveInt *int              `json:"seriesDosesPositiveInt,omitempty"`
	SeriesDosesString      *string           `json:"seriesDosesString,omitempty"`
}

type OtherImmunization Immunization

// on convert struct to json, automatically add resourceType=Immunization
func (r Immunization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImmunization
		ResourceType string `json:"resourceType"`
	}{
		OtherImmunization: OtherImmunization(r),
		ResourceType:      "Immunization",
	})
}

func (resource *Immunization) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Immunization.Id", nil)
	}
	return StringInput("Immunization.Id", resource.Id)
}
func (resource *Immunization) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Immunization.ImplicitRules", nil)
	}
	return StringInput("Immunization.ImplicitRules", resource.ImplicitRules)
}
func (resource *Immunization) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Immunization.Language", nil, optionsValueSet)
	}
	return CodeSelect("Immunization.Language", resource.Language, optionsValueSet)
}
func (resource *Immunization) T_Status() templ.Component {
	optionsValueSet := VSImmunization_status

	if resource == nil {
		return CodeSelect("Immunization.Status", nil, optionsValueSet)
	}
	return CodeSelect("Immunization.Status", &resource.Status, optionsValueSet)
}
func (resource *Immunization) T_StatusReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Immunization.StatusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Immunization.StatusReason", resource.StatusReason, optionsValueSet)
}
func (resource *Immunization) T_VaccineCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Immunization.VaccineCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Immunization.VaccineCode", &resource.VaccineCode, optionsValueSet)
}
func (resource *Immunization) T_Recorded() templ.Component {

	if resource == nil {
		return StringInput("Immunization.Recorded", nil)
	}
	return StringInput("Immunization.Recorded", resource.Recorded)
}
func (resource *Immunization) T_PrimarySource() templ.Component {

	if resource == nil {
		return BoolInput("Immunization.PrimarySource", nil)
	}
	return BoolInput("Immunization.PrimarySource", resource.PrimarySource)
}
func (resource *Immunization) T_ReportOrigin(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Immunization.ReportOrigin", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Immunization.ReportOrigin", resource.ReportOrigin, optionsValueSet)
}
func (resource *Immunization) T_LotNumber() templ.Component {

	if resource == nil {
		return StringInput("Immunization.LotNumber", nil)
	}
	return StringInput("Immunization.LotNumber", resource.LotNumber)
}
func (resource *Immunization) T_ExpirationDate() templ.Component {

	if resource == nil {
		return StringInput("Immunization.ExpirationDate", nil)
	}
	return StringInput("Immunization.ExpirationDate", resource.ExpirationDate)
}
func (resource *Immunization) T_Site(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Immunization.Site", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Immunization.Site", resource.Site, optionsValueSet)
}
func (resource *Immunization) T_Route(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Immunization.Route", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Immunization.Route", resource.Route, optionsValueSet)
}
func (resource *Immunization) T_ReasonCode(numReasonCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ReasonCode) >= numReasonCode {
		return CodeableConceptSelect("Immunization.ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Immunization.ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet)
}
func (resource *Immunization) T_IsSubpotent() templ.Component {

	if resource == nil {
		return BoolInput("Immunization.IsSubpotent", nil)
	}
	return BoolInput("Immunization.IsSubpotent", resource.IsSubpotent)
}
func (resource *Immunization) T_SubpotentReason(numSubpotentReason int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SubpotentReason) >= numSubpotentReason {
		return CodeableConceptSelect("Immunization.SubpotentReason["+strconv.Itoa(numSubpotentReason)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Immunization.SubpotentReason["+strconv.Itoa(numSubpotentReason)+"]", &resource.SubpotentReason[numSubpotentReason], optionsValueSet)
}
func (resource *Immunization) T_ProgramEligibility(numProgramEligibility int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ProgramEligibility) >= numProgramEligibility {
		return CodeableConceptSelect("Immunization.ProgramEligibility["+strconv.Itoa(numProgramEligibility)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Immunization.ProgramEligibility["+strconv.Itoa(numProgramEligibility)+"]", &resource.ProgramEligibility[numProgramEligibility], optionsValueSet)
}
func (resource *Immunization) T_FundingSource(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Immunization.FundingSource", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Immunization.FundingSource", resource.FundingSource, optionsValueSet)
}
func (resource *Immunization) T_PerformerId(numPerformer int) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return StringInput("Immunization.Performer["+strconv.Itoa(numPerformer)+"].Id", nil)
	}
	return StringInput("Immunization.Performer["+strconv.Itoa(numPerformer)+"].Id", resource.Performer[numPerformer].Id)
}
func (resource *Immunization) T_PerformerFunction(numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return CodeableConceptSelect("Immunization.Performer["+strconv.Itoa(numPerformer)+"].Function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Immunization.Performer["+strconv.Itoa(numPerformer)+"].Function", resource.Performer[numPerformer].Function, optionsValueSet)
}
func (resource *Immunization) T_EducationId(numEducation int) templ.Component {

	if resource == nil || len(resource.Education) >= numEducation {
		return StringInput("Immunization.Education["+strconv.Itoa(numEducation)+"].Id", nil)
	}
	return StringInput("Immunization.Education["+strconv.Itoa(numEducation)+"].Id", resource.Education[numEducation].Id)
}
func (resource *Immunization) T_EducationDocumentType(numEducation int) templ.Component {

	if resource == nil || len(resource.Education) >= numEducation {
		return StringInput("Immunization.Education["+strconv.Itoa(numEducation)+"].DocumentType", nil)
	}
	return StringInput("Immunization.Education["+strconv.Itoa(numEducation)+"].DocumentType", resource.Education[numEducation].DocumentType)
}
func (resource *Immunization) T_EducationReference(numEducation int) templ.Component {

	if resource == nil || len(resource.Education) >= numEducation {
		return StringInput("Immunization.Education["+strconv.Itoa(numEducation)+"].Reference", nil)
	}
	return StringInput("Immunization.Education["+strconv.Itoa(numEducation)+"].Reference", resource.Education[numEducation].Reference)
}
func (resource *Immunization) T_EducationPublicationDate(numEducation int) templ.Component {

	if resource == nil || len(resource.Education) >= numEducation {
		return StringInput("Immunization.Education["+strconv.Itoa(numEducation)+"].PublicationDate", nil)
	}
	return StringInput("Immunization.Education["+strconv.Itoa(numEducation)+"].PublicationDate", resource.Education[numEducation].PublicationDate)
}
func (resource *Immunization) T_EducationPresentationDate(numEducation int) templ.Component {

	if resource == nil || len(resource.Education) >= numEducation {
		return StringInput("Immunization.Education["+strconv.Itoa(numEducation)+"].PresentationDate", nil)
	}
	return StringInput("Immunization.Education["+strconv.Itoa(numEducation)+"].PresentationDate", resource.Education[numEducation].PresentationDate)
}
func (resource *Immunization) T_ReactionId(numReaction int) templ.Component {

	if resource == nil || len(resource.Reaction) >= numReaction {
		return StringInput("Immunization.Reaction["+strconv.Itoa(numReaction)+"].Id", nil)
	}
	return StringInput("Immunization.Reaction["+strconv.Itoa(numReaction)+"].Id", resource.Reaction[numReaction].Id)
}
func (resource *Immunization) T_ReactionDate(numReaction int) templ.Component {

	if resource == nil || len(resource.Reaction) >= numReaction {
		return StringInput("Immunization.Reaction["+strconv.Itoa(numReaction)+"].Date", nil)
	}
	return StringInput("Immunization.Reaction["+strconv.Itoa(numReaction)+"].Date", resource.Reaction[numReaction].Date)
}
func (resource *Immunization) T_ReactionReported(numReaction int) templ.Component {

	if resource == nil || len(resource.Reaction) >= numReaction {
		return BoolInput("Immunization.Reaction["+strconv.Itoa(numReaction)+"].Reported", nil)
	}
	return BoolInput("Immunization.Reaction["+strconv.Itoa(numReaction)+"].Reported", resource.Reaction[numReaction].Reported)
}
func (resource *Immunization) T_ProtocolAppliedId(numProtocolApplied int) templ.Component {

	if resource == nil || len(resource.ProtocolApplied) >= numProtocolApplied {
		return StringInput("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].Id", nil)
	}
	return StringInput("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].Id", resource.ProtocolApplied[numProtocolApplied].Id)
}
func (resource *Immunization) T_ProtocolAppliedSeries(numProtocolApplied int) templ.Component {

	if resource == nil || len(resource.ProtocolApplied) >= numProtocolApplied {
		return StringInput("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].Series", nil)
	}
	return StringInput("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].Series", resource.ProtocolApplied[numProtocolApplied].Series)
}
func (resource *Immunization) T_ProtocolAppliedTargetDisease(numProtocolApplied int, numTargetDisease int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ProtocolApplied) >= numProtocolApplied || len(resource.ProtocolApplied[numProtocolApplied].TargetDisease) >= numTargetDisease {
		return CodeableConceptSelect("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].TargetDisease["+strconv.Itoa(numTargetDisease)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].TargetDisease["+strconv.Itoa(numTargetDisease)+"]", &resource.ProtocolApplied[numProtocolApplied].TargetDisease[numTargetDisease], optionsValueSet)
}
