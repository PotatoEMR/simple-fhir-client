package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Immunization
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
	OccurrenceDateTime FhirDateTime                  `json:"occurrenceDateTime"`
	OccurrenceString   string                        `json:"occurrenceString"`
	Recorded           *FhirDateTime                 `json:"recorded,omitempty"`
	PrimarySource      *bool                         `json:"primarySource,omitempty"`
	ReportOrigin       *CodeableConcept              `json:"reportOrigin,omitempty"`
	Location           *Reference                    `json:"location,omitempty"`
	Manufacturer       *Reference                    `json:"manufacturer,omitempty"`
	LotNumber          *string                       `json:"lotNumber,omitempty"`
	ExpirationDate     *FhirDate                     `json:"expirationDate,omitempty"`
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

// http://hl7.org/fhir/r4/StructureDefinition/Immunization
type ImmunizationPerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Immunization
type ImmunizationEducation struct {
	Id                *string       `json:"id,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	DocumentType      *string       `json:"documentType,omitempty"`
	Reference         *string       `json:"reference,omitempty"`
	PublicationDate   *FhirDateTime `json:"publicationDate,omitempty"`
	PresentationDate  *FhirDateTime `json:"presentationDate,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Immunization
type ImmunizationReaction struct {
	Id                *string       `json:"id,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	Date              *FhirDateTime `json:"date,omitempty"`
	Detail            *Reference    `json:"detail,omitempty"`
	Reported          *bool         `json:"reported,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Immunization
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
func (resource *Immunization) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSImmunization_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_StatusReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statusReason", resource.StatusReason, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_VaccineCode(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("vaccineCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("vaccineCode", &resource.VaccineCode, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_OccurrenceDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("occurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("occurrenceDateTime", &resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *Immunization) T_OccurrenceString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("occurrenceString", nil, htmlAttrs)
	}
	return StringInput("occurrenceString", &resource.OccurrenceString, htmlAttrs)
}
func (resource *Immunization) T_Recorded(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("recorded", nil, htmlAttrs)
	}
	return DateTimeInput("recorded", resource.Recorded, htmlAttrs)
}
func (resource *Immunization) T_PrimarySource(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("primarySource", nil, htmlAttrs)
	}
	return BoolInput("primarySource", resource.PrimarySource, htmlAttrs)
}
func (resource *Immunization) T_ReportOrigin(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("reportOrigin", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reportOrigin", resource.ReportOrigin, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_LotNumber(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("lotNumber", nil, htmlAttrs)
	}
	return StringInput("lotNumber", resource.LotNumber, htmlAttrs)
}
func (resource *Immunization) T_ExpirationDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateInput("expirationDate", nil, htmlAttrs)
	}
	return DateInput("expirationDate", resource.ExpirationDate, htmlAttrs)
}
func (resource *Immunization) T_Site(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("site", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("site", resource.Site, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_Route(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("route", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("route", resource.Route, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Immunization) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_IsSubpotent(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("isSubpotent", nil, htmlAttrs)
	}
	return BoolInput("isSubpotent", resource.IsSubpotent, htmlAttrs)
}
func (resource *Immunization) T_SubpotentReason(numSubpotentReason int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubpotentReason >= len(resource.SubpotentReason) {
		return CodeableConceptSelect("subpotentReason["+strconv.Itoa(numSubpotentReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subpotentReason["+strconv.Itoa(numSubpotentReason)+"]", &resource.SubpotentReason[numSubpotentReason], optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_ProgramEligibility(numProgramEligibility int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProgramEligibility >= len(resource.ProgramEligibility) {
		return CodeableConceptSelect("programEligibility["+strconv.Itoa(numProgramEligibility)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("programEligibility["+strconv.Itoa(numProgramEligibility)+"]", &resource.ProgramEligibility[numProgramEligibility], optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_FundingSource(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("fundingSource", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("fundingSource", resource.FundingSource, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_EducationDocumentType(numEducation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEducation >= len(resource.Education) {
		return StringInput("education["+strconv.Itoa(numEducation)+"].documentType", nil, htmlAttrs)
	}
	return StringInput("education["+strconv.Itoa(numEducation)+"].documentType", resource.Education[numEducation].DocumentType, htmlAttrs)
}
func (resource *Immunization) T_EducationReference(numEducation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEducation >= len(resource.Education) {
		return StringInput("education["+strconv.Itoa(numEducation)+"].reference", nil, htmlAttrs)
	}
	return StringInput("education["+strconv.Itoa(numEducation)+"].reference", resource.Education[numEducation].Reference, htmlAttrs)
}
func (resource *Immunization) T_EducationPublicationDate(numEducation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEducation >= len(resource.Education) {
		return DateTimeInput("education["+strconv.Itoa(numEducation)+"].publicationDate", nil, htmlAttrs)
	}
	return DateTimeInput("education["+strconv.Itoa(numEducation)+"].publicationDate", resource.Education[numEducation].PublicationDate, htmlAttrs)
}
func (resource *Immunization) T_EducationPresentationDate(numEducation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEducation >= len(resource.Education) {
		return DateTimeInput("education["+strconv.Itoa(numEducation)+"].presentationDate", nil, htmlAttrs)
	}
	return DateTimeInput("education["+strconv.Itoa(numEducation)+"].presentationDate", resource.Education[numEducation].PresentationDate, htmlAttrs)
}
func (resource *Immunization) T_ReactionDate(numReaction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) {
		return DateTimeInput("reaction["+strconv.Itoa(numReaction)+"].date", nil, htmlAttrs)
	}
	return DateTimeInput("reaction["+strconv.Itoa(numReaction)+"].date", resource.Reaction[numReaction].Date, htmlAttrs)
}
func (resource *Immunization) T_ReactionReported(numReaction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) {
		return BoolInput("reaction["+strconv.Itoa(numReaction)+"].reported", nil, htmlAttrs)
	}
	return BoolInput("reaction["+strconv.Itoa(numReaction)+"].reported", resource.Reaction[numReaction].Reported, htmlAttrs)
}
func (resource *Immunization) T_ProtocolAppliedSeries(numProtocolApplied int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProtocolApplied >= len(resource.ProtocolApplied) {
		return StringInput("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].series", nil, htmlAttrs)
	}
	return StringInput("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].series", resource.ProtocolApplied[numProtocolApplied].Series, htmlAttrs)
}
func (resource *Immunization) T_ProtocolAppliedTargetDisease(numProtocolApplied int, numTargetDisease int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProtocolApplied >= len(resource.ProtocolApplied) || numTargetDisease >= len(resource.ProtocolApplied[numProtocolApplied].TargetDisease) {
		return CodeableConceptSelect("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].targetDisease["+strconv.Itoa(numTargetDisease)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].targetDisease["+strconv.Itoa(numTargetDisease)+"]", &resource.ProtocolApplied[numProtocolApplied].TargetDisease[numTargetDisease], optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_ProtocolAppliedDoseNumberPositiveInt(numProtocolApplied int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProtocolApplied >= len(resource.ProtocolApplied) {
		return IntInput("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].doseNumberPositiveInt", nil, htmlAttrs)
	}
	return IntInput("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].doseNumberPositiveInt", &resource.ProtocolApplied[numProtocolApplied].DoseNumberPositiveInt, htmlAttrs)
}
func (resource *Immunization) T_ProtocolAppliedDoseNumberString(numProtocolApplied int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProtocolApplied >= len(resource.ProtocolApplied) {
		return StringInput("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].doseNumberString", nil, htmlAttrs)
	}
	return StringInput("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].doseNumberString", &resource.ProtocolApplied[numProtocolApplied].DoseNumberString, htmlAttrs)
}
func (resource *Immunization) T_ProtocolAppliedSeriesDosesPositiveInt(numProtocolApplied int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProtocolApplied >= len(resource.ProtocolApplied) {
		return IntInput("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].seriesDosesPositiveInt", nil, htmlAttrs)
	}
	return IntInput("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].seriesDosesPositiveInt", resource.ProtocolApplied[numProtocolApplied].SeriesDosesPositiveInt, htmlAttrs)
}
func (resource *Immunization) T_ProtocolAppliedSeriesDosesString(numProtocolApplied int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProtocolApplied >= len(resource.ProtocolApplied) {
		return StringInput("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].seriesDosesString", nil, htmlAttrs)
	}
	return StringInput("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].seriesDosesString", resource.ProtocolApplied[numProtocolApplied].SeriesDosesString, htmlAttrs)
}
