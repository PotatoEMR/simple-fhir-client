package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	SeriesDosesPositiveInt *int              `json:"seriesDosesPositiveInt"`
	SeriesDosesString      *string           `json:"seriesDosesString"`
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

func (resource *Immunization) ImmunizationLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Immunization) ImmunizationStatus() templ.Component {
	optionsValueSet := VSImmunization_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Immunization) ImmunizationStatusReason(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("statusReason", resource.StatusReason, optionsValueSet)
}
func (resource *Immunization) ImmunizationVaccineCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("vaccineCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("vaccineCode", &resource.VaccineCode, optionsValueSet)
}
func (resource *Immunization) ImmunizationReportOrigin(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("reportOrigin", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reportOrigin", resource.ReportOrigin, optionsValueSet)
}
func (resource *Immunization) ImmunizationSite(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("site", nil, optionsValueSet)
	}
	return CodeableConceptSelect("site", resource.Site, optionsValueSet)
}
func (resource *Immunization) ImmunizationRoute(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("route", nil, optionsValueSet)
	}
	return CodeableConceptSelect("route", resource.Route, optionsValueSet)
}
func (resource *Immunization) ImmunizationReasonCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("reasonCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reasonCode", &resource.ReasonCode[0], optionsValueSet)
}
func (resource *Immunization) ImmunizationSubpotentReason(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("subpotentReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subpotentReason", &resource.SubpotentReason[0], optionsValueSet)
}
func (resource *Immunization) ImmunizationProgramEligibility(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("programEligibility", nil, optionsValueSet)
	}
	return CodeableConceptSelect("programEligibility", &resource.ProgramEligibility[0], optionsValueSet)
}
func (resource *Immunization) ImmunizationFundingSource(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("fundingSource", nil, optionsValueSet)
	}
	return CodeableConceptSelect("fundingSource", resource.FundingSource, optionsValueSet)
}
func (resource *Immunization) ImmunizationPerformerFunction(numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Performer) >= numPerformer {
		return CodeableConceptSelect("function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("function", resource.Performer[numPerformer].Function, optionsValueSet)
}
func (resource *Immunization) ImmunizationProtocolAppliedTargetDisease(numProtocolApplied int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.ProtocolApplied) >= numProtocolApplied {
		return CodeableConceptSelect("targetDisease", nil, optionsValueSet)
	}
	return CodeableConceptSelect("targetDisease", &resource.ProtocolApplied[numProtocolApplied].TargetDisease[0], optionsValueSet)
}
