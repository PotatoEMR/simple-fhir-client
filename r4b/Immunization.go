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
	OccurrenceDateTime time.Time                     `json:"occurrenceDateTime,format:'2006-01-02T15:04:05Z07:00'"`
	OccurrenceString   string                        `json:"occurrenceString"`
	Recorded           *time.Time                    `json:"recorded,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	PrimarySource      *bool                         `json:"primarySource,omitempty"`
	ReportOrigin       *CodeableConcept              `json:"reportOrigin,omitempty"`
	Location           *Reference                    `json:"location,omitempty"`
	Manufacturer       *Reference                    `json:"manufacturer,omitempty"`
	LotNumber          *string                       `json:"lotNumber,omitempty"`
	ExpirationDate     *time.Time                    `json:"expirationDate,omitempty,format:'2006-01-02'"`
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
	PublicationDate   *time.Time  `json:"publicationDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	PresentationDate  *time.Time  `json:"presentationDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Immunization
type ImmunizationReaction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Date              *time.Time  `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r Immunization) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Immunization/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Immunization"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Immunization) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSImmunization_status

	if resource == nil {
		return CodeSelect("Immunization.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Immunization.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_StatusReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Immunization.StatusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Immunization.StatusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_VaccineCode(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Immunization.VaccineCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Immunization.VaccineCode", &resource.VaccineCode, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_OccurrenceDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Immunization.OccurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Immunization.OccurrenceDateTime", &resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *Immunization) T_OccurrenceString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Immunization.OccurrenceString", nil, htmlAttrs)
	}
	return StringInput("Immunization.OccurrenceString", &resource.OccurrenceString, htmlAttrs)
}
func (resource *Immunization) T_Recorded(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Immunization.Recorded", nil, htmlAttrs)
	}
	return DateTimeInput("Immunization.Recorded", resource.Recorded, htmlAttrs)
}
func (resource *Immunization) T_PrimarySource(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Immunization.PrimarySource", nil, htmlAttrs)
	}
	return BoolInput("Immunization.PrimarySource", resource.PrimarySource, htmlAttrs)
}
func (resource *Immunization) T_ReportOrigin(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Immunization.ReportOrigin", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Immunization.ReportOrigin", resource.ReportOrigin, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_LotNumber(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Immunization.LotNumber", nil, htmlAttrs)
	}
	return StringInput("Immunization.LotNumber", resource.LotNumber, htmlAttrs)
}
func (resource *Immunization) T_ExpirationDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("Immunization.ExpirationDate", nil, htmlAttrs)
	}
	return DateInput("Immunization.ExpirationDate", resource.ExpirationDate, htmlAttrs)
}
func (resource *Immunization) T_Site(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Immunization.Site", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Immunization.Site", resource.Site, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_Route(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Immunization.Route", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Immunization.Route", resource.Route, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Immunization.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Immunization.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Immunization) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("Immunization.ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Immunization.ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_IsSubpotent(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Immunization.IsSubpotent", nil, htmlAttrs)
	}
	return BoolInput("Immunization.IsSubpotent", resource.IsSubpotent, htmlAttrs)
}
func (resource *Immunization) T_SubpotentReason(numSubpotentReason int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSubpotentReason >= len(resource.SubpotentReason) {
		return CodeableConceptSelect("Immunization.SubpotentReason["+strconv.Itoa(numSubpotentReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Immunization.SubpotentReason["+strconv.Itoa(numSubpotentReason)+"]", &resource.SubpotentReason[numSubpotentReason], optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_ProgramEligibility(numProgramEligibility int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProgramEligibility >= len(resource.ProgramEligibility) {
		return CodeableConceptSelect("Immunization.ProgramEligibility["+strconv.Itoa(numProgramEligibility)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Immunization.ProgramEligibility["+strconv.Itoa(numProgramEligibility)+"]", &resource.ProgramEligibility[numProgramEligibility], optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_FundingSource(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Immunization.FundingSource", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Immunization.FundingSource", resource.FundingSource, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("Immunization.Performer["+strconv.Itoa(numPerformer)+"].Function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Immunization.Performer["+strconv.Itoa(numPerformer)+"].Function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_EducationDocumentType(numEducation int, htmlAttrs string) templ.Component {
	if resource == nil || numEducation >= len(resource.Education) {
		return StringInput("Immunization.Education["+strconv.Itoa(numEducation)+"].DocumentType", nil, htmlAttrs)
	}
	return StringInput("Immunization.Education["+strconv.Itoa(numEducation)+"].DocumentType", resource.Education[numEducation].DocumentType, htmlAttrs)
}
func (resource *Immunization) T_EducationReference(numEducation int, htmlAttrs string) templ.Component {
	if resource == nil || numEducation >= len(resource.Education) {
		return StringInput("Immunization.Education["+strconv.Itoa(numEducation)+"].Reference", nil, htmlAttrs)
	}
	return StringInput("Immunization.Education["+strconv.Itoa(numEducation)+"].Reference", resource.Education[numEducation].Reference, htmlAttrs)
}
func (resource *Immunization) T_EducationPublicationDate(numEducation int, htmlAttrs string) templ.Component {
	if resource == nil || numEducation >= len(resource.Education) {
		return DateTimeInput("Immunization.Education["+strconv.Itoa(numEducation)+"].PublicationDate", nil, htmlAttrs)
	}
	return DateTimeInput("Immunization.Education["+strconv.Itoa(numEducation)+"].PublicationDate", resource.Education[numEducation].PublicationDate, htmlAttrs)
}
func (resource *Immunization) T_EducationPresentationDate(numEducation int, htmlAttrs string) templ.Component {
	if resource == nil || numEducation >= len(resource.Education) {
		return DateTimeInput("Immunization.Education["+strconv.Itoa(numEducation)+"].PresentationDate", nil, htmlAttrs)
	}
	return DateTimeInput("Immunization.Education["+strconv.Itoa(numEducation)+"].PresentationDate", resource.Education[numEducation].PresentationDate, htmlAttrs)
}
func (resource *Immunization) T_ReactionDate(numReaction int, htmlAttrs string) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) {
		return DateTimeInput("Immunization.Reaction["+strconv.Itoa(numReaction)+"].Date", nil, htmlAttrs)
	}
	return DateTimeInput("Immunization.Reaction["+strconv.Itoa(numReaction)+"].Date", resource.Reaction[numReaction].Date, htmlAttrs)
}
func (resource *Immunization) T_ReactionReported(numReaction int, htmlAttrs string) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) {
		return BoolInput("Immunization.Reaction["+strconv.Itoa(numReaction)+"].Reported", nil, htmlAttrs)
	}
	return BoolInput("Immunization.Reaction["+strconv.Itoa(numReaction)+"].Reported", resource.Reaction[numReaction].Reported, htmlAttrs)
}
func (resource *Immunization) T_ProtocolAppliedSeries(numProtocolApplied int, htmlAttrs string) templ.Component {
	if resource == nil || numProtocolApplied >= len(resource.ProtocolApplied) {
		return StringInput("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].Series", nil, htmlAttrs)
	}
	return StringInput("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].Series", resource.ProtocolApplied[numProtocolApplied].Series, htmlAttrs)
}
func (resource *Immunization) T_ProtocolAppliedTargetDisease(numProtocolApplied int, numTargetDisease int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProtocolApplied >= len(resource.ProtocolApplied) || numTargetDisease >= len(resource.ProtocolApplied[numProtocolApplied].TargetDisease) {
		return CodeableConceptSelect("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].TargetDisease["+strconv.Itoa(numTargetDisease)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].TargetDisease["+strconv.Itoa(numTargetDisease)+"]", &resource.ProtocolApplied[numProtocolApplied].TargetDisease[numTargetDisease], optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_ProtocolAppliedDoseNumberPositiveInt(numProtocolApplied int, htmlAttrs string) templ.Component {
	if resource == nil || numProtocolApplied >= len(resource.ProtocolApplied) {
		return IntInput("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].DoseNumberPositiveInt", nil, htmlAttrs)
	}
	return IntInput("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].DoseNumberPositiveInt", &resource.ProtocolApplied[numProtocolApplied].DoseNumberPositiveInt, htmlAttrs)
}
func (resource *Immunization) T_ProtocolAppliedDoseNumberString(numProtocolApplied int, htmlAttrs string) templ.Component {
	if resource == nil || numProtocolApplied >= len(resource.ProtocolApplied) {
		return StringInput("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].DoseNumberString", nil, htmlAttrs)
	}
	return StringInput("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].DoseNumberString", &resource.ProtocolApplied[numProtocolApplied].DoseNumberString, htmlAttrs)
}
func (resource *Immunization) T_ProtocolAppliedSeriesDosesPositiveInt(numProtocolApplied int, htmlAttrs string) templ.Component {
	if resource == nil || numProtocolApplied >= len(resource.ProtocolApplied) {
		return IntInput("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].SeriesDosesPositiveInt", nil, htmlAttrs)
	}
	return IntInput("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].SeriesDosesPositiveInt", resource.ProtocolApplied[numProtocolApplied].SeriesDosesPositiveInt, htmlAttrs)
}
func (resource *Immunization) T_ProtocolAppliedSeriesDosesString(numProtocolApplied int, htmlAttrs string) templ.Component {
	if resource == nil || numProtocolApplied >= len(resource.ProtocolApplied) {
		return StringInput("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].SeriesDosesString", nil, htmlAttrs)
	}
	return StringInput("Immunization.ProtocolApplied["+strconv.Itoa(numProtocolApplied)+"].SeriesDosesString", resource.ProtocolApplied[numProtocolApplied].SeriesDosesString, htmlAttrs)
}
