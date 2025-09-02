package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/Immunization
type Immunization struct {
	Id                    *string                          `json:"id,omitempty"`
	Meta                  *Meta                            `json:"meta,omitempty"`
	ImplicitRules         *string                          `json:"implicitRules,omitempty"`
	Language              *string                          `json:"language,omitempty"`
	Text                  *Narrative                       `json:"text,omitempty"`
	Contained             []Resource                       `json:"contained,omitempty"`
	Extension             []Extension                      `json:"extension,omitempty"`
	ModifierExtension     []Extension                      `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                     `json:"identifier,omitempty"`
	BasedOn               []Reference                      `json:"basedOn,omitempty"`
	Status                string                           `json:"status"`
	StatusReason          *CodeableConcept                 `json:"statusReason,omitempty"`
	VaccineCode           CodeableConcept                  `json:"vaccineCode"`
	AdministeredProduct   *CodeableReference               `json:"administeredProduct,omitempty"`
	Manufacturer          *CodeableReference               `json:"manufacturer,omitempty"`
	LotNumber             *string                          `json:"lotNumber,omitempty"`
	ExpirationDate        *string                          `json:"expirationDate,omitempty"`
	Patient               Reference                        `json:"patient"`
	Encounter             *Reference                       `json:"encounter,omitempty"`
	SupportingInformation []Reference                      `json:"supportingInformation,omitempty"`
	OccurrenceDateTime    string                           `json:"occurrenceDateTime"`
	OccurrenceString      string                           `json:"occurrenceString"`
	PrimarySource         *bool                            `json:"primarySource,omitempty"`
	InformationSource     *CodeableReference               `json:"informationSource,omitempty"`
	Location              *Reference                       `json:"location,omitempty"`
	Site                  *CodeableConcept                 `json:"site,omitempty"`
	Route                 *CodeableConcept                 `json:"route,omitempty"`
	DoseQuantity          *Quantity                        `json:"doseQuantity,omitempty"`
	Performer             []ImmunizationPerformer          `json:"performer,omitempty"`
	Note                  []Annotation                     `json:"note,omitempty"`
	Reason                []CodeableReference              `json:"reason,omitempty"`
	IsSubpotent           *bool                            `json:"isSubpotent,omitempty"`
	SubpotentReason       []CodeableConcept                `json:"subpotentReason,omitempty"`
	ProgramEligibility    []ImmunizationProgramEligibility `json:"programEligibility,omitempty"`
	FundingSource         *CodeableConcept                 `json:"fundingSource,omitempty"`
	Reaction              []ImmunizationReaction           `json:"reaction,omitempty"`
	ProtocolApplied       []ImmunizationProtocolApplied    `json:"protocolApplied,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Immunization
type ImmunizationPerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Immunization
type ImmunizationProgramEligibility struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Program           CodeableConcept `json:"program"`
	ProgramStatus     CodeableConcept `json:"programStatus"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Immunization
type ImmunizationReaction struct {
	Id                *string            `json:"id,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Date              *string            `json:"date,omitempty"`
	Manifestation     *CodeableReference `json:"manifestation,omitempty"`
	Reported          *bool              `json:"reported,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Immunization
type ImmunizationProtocolApplied struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Series            *string           `json:"series,omitempty"`
	Authority         *Reference        `json:"authority,omitempty"`
	TargetDisease     []CodeableConcept `json:"targetDisease,omitempty"`
	DoseNumber        string            `json:"doseNumber"`
	SeriesDoses       *string           `json:"seriesDoses,omitempty"`
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

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Immunization) ImmunizationStatus() templ.Component {
	optionsValueSet := VSImmunization_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Immunization) ImmunizationStatusReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("statusReason", resource.StatusReason, optionsValueSet)
}
func (resource *Immunization) ImmunizationVaccineCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("vaccineCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("vaccineCode", &resource.VaccineCode, optionsValueSet)
}
func (resource *Immunization) ImmunizationSite(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("site", nil, optionsValueSet)
	}
	return CodeableConceptSelect("site", resource.Site, optionsValueSet)
}
func (resource *Immunization) ImmunizationRoute(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("route", nil, optionsValueSet)
	}
	return CodeableConceptSelect("route", resource.Route, optionsValueSet)
}
func (resource *Immunization) ImmunizationSubpotentReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("subpotentReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subpotentReason", &resource.SubpotentReason[0], optionsValueSet)
}
func (resource *Immunization) ImmunizationFundingSource(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("fundingSource", nil, optionsValueSet)
	}
	return CodeableConceptSelect("fundingSource", resource.FundingSource, optionsValueSet)
}
func (resource *Immunization) ImmunizationPerformerFunction(numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Performer) >= numPerformer {
		return CodeableConceptSelect("function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("function", resource.Performer[numPerformer].Function, optionsValueSet)
}
func (resource *Immunization) ImmunizationProgramEligibilityProgram(numProgramEligibility int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.ProgramEligibility) >= numProgramEligibility {
		return CodeableConceptSelect("program", nil, optionsValueSet)
	}
	return CodeableConceptSelect("program", &resource.ProgramEligibility[numProgramEligibility].Program, optionsValueSet)
}
func (resource *Immunization) ImmunizationProgramEligibilityProgramStatus(numProgramEligibility int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.ProgramEligibility) >= numProgramEligibility {
		return CodeableConceptSelect("programStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("programStatus", &resource.ProgramEligibility[numProgramEligibility].ProgramStatus, optionsValueSet)
}
func (resource *Immunization) ImmunizationProtocolAppliedTargetDisease(numProtocolApplied int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.ProtocolApplied) >= numProtocolApplied {
		return CodeableConceptSelect("targetDisease", nil, optionsValueSet)
	}
	return CodeableConceptSelect("targetDisease", &resource.ProtocolApplied[numProtocolApplied].TargetDisease[0], optionsValueSet)
}
