package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	ExpirationDate        *FhirDate                        `json:"expirationDate,omitempty"`
	Patient               Reference                        `json:"patient"`
	Encounter             *Reference                       `json:"encounter,omitempty"`
	SupportingInformation []Reference                      `json:"supportingInformation,omitempty"`
	OccurrenceDateTime    FhirDateTime                     `json:"occurrenceDateTime"`
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
	Date              *FhirDateTime      `json:"date,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
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
func (r Immunization) ResourceType() string {
	return "Immunization"
}

func (resource *Immunization) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
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
func (resource *Immunization) T_AdministeredProduct(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("administeredProduct", nil, htmlAttrs)
	}
	return CodeableReferenceInput("administeredProduct", resource.AdministeredProduct, htmlAttrs)
}
func (resource *Immunization) T_Manufacturer(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("manufacturer", nil, htmlAttrs)
	}
	return CodeableReferenceInput("manufacturer", resource.Manufacturer, htmlAttrs)
}
func (resource *Immunization) T_LotNumber(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("lotNumber", nil, htmlAttrs)
	}
	return StringInput("lotNumber", resource.LotNumber, htmlAttrs)
}
func (resource *Immunization) T_ExpirationDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("expirationDate", nil, htmlAttrs)
	}
	return FhirDateInput("expirationDate", resource.ExpirationDate, htmlAttrs)
}
func (resource *Immunization) T_Patient(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "patient", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "patient", &resource.Patient, htmlAttrs)
}
func (resource *Immunization) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *Immunization) T_SupportingInformation(frs []FhirResource, numSupportingInformation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInformation >= len(resource.SupportingInformation) {
		return ReferenceInput(frs, "supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", &resource.SupportingInformation[numSupportingInformation], htmlAttrs)
}
func (resource *Immunization) T_OccurrenceDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("occurrenceDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("occurrenceDateTime", &resource.OccurrenceDateTime, htmlAttrs)
}
func (resource *Immunization) T_OccurrenceString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("occurrenceString", nil, htmlAttrs)
	}
	return StringInput("occurrenceString", &resource.OccurrenceString, htmlAttrs)
}
func (resource *Immunization) T_PrimarySource(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("primarySource", nil, htmlAttrs)
	}
	return BoolInput("primarySource", resource.PrimarySource, htmlAttrs)
}
func (resource *Immunization) T_InformationSource(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("informationSource", nil, htmlAttrs)
	}
	return CodeableReferenceInput("informationSource", resource.InformationSource, htmlAttrs)
}
func (resource *Immunization) T_Location(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "location", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "location", resource.Location, htmlAttrs)
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
func (resource *Immunization) T_DoseQuantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("doseQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("doseQuantity", resource.DoseQuantity, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Immunization) T_Reason(numReason int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], htmlAttrs)
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
func (resource *Immunization) T_PerformerActor(frs []FhirResource, numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"].actor", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"].actor", &resource.Performer[numPerformer].Actor, htmlAttrs)
}
func (resource *Immunization) T_ProgramEligibilityProgram(numProgramEligibility int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProgramEligibility >= len(resource.ProgramEligibility) {
		return CodeableConceptSelect("programEligibility["+strconv.Itoa(numProgramEligibility)+"].program", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("programEligibility["+strconv.Itoa(numProgramEligibility)+"].program", &resource.ProgramEligibility[numProgramEligibility].Program, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_ProgramEligibilityProgramStatus(numProgramEligibility int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProgramEligibility >= len(resource.ProgramEligibility) {
		return CodeableConceptSelect("programEligibility["+strconv.Itoa(numProgramEligibility)+"].programStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("programEligibility["+strconv.Itoa(numProgramEligibility)+"].programStatus", &resource.ProgramEligibility[numProgramEligibility].ProgramStatus, optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_ReactionDate(numReaction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) {
		return FhirDateTimeInput("reaction["+strconv.Itoa(numReaction)+"].date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("reaction["+strconv.Itoa(numReaction)+"].date", resource.Reaction[numReaction].Date, htmlAttrs)
}
func (resource *Immunization) T_ReactionManifestation(numReaction int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReaction >= len(resource.Reaction) {
		return CodeableReferenceInput("reaction["+strconv.Itoa(numReaction)+"].manifestation", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reaction["+strconv.Itoa(numReaction)+"].manifestation", resource.Reaction[numReaction].Manifestation, htmlAttrs)
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
func (resource *Immunization) T_ProtocolAppliedAuthority(frs []FhirResource, numProtocolApplied int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProtocolApplied >= len(resource.ProtocolApplied) {
		return ReferenceInput(frs, "protocolApplied["+strconv.Itoa(numProtocolApplied)+"].authority", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "protocolApplied["+strconv.Itoa(numProtocolApplied)+"].authority", resource.ProtocolApplied[numProtocolApplied].Authority, htmlAttrs)
}
func (resource *Immunization) T_ProtocolAppliedTargetDisease(numProtocolApplied int, numTargetDisease int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProtocolApplied >= len(resource.ProtocolApplied) || numTargetDisease >= len(resource.ProtocolApplied[numProtocolApplied].TargetDisease) {
		return CodeableConceptSelect("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].targetDisease["+strconv.Itoa(numTargetDisease)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].targetDisease["+strconv.Itoa(numTargetDisease)+"]", &resource.ProtocolApplied[numProtocolApplied].TargetDisease[numTargetDisease], optionsValueSet, htmlAttrs)
}
func (resource *Immunization) T_ProtocolAppliedDoseNumber(numProtocolApplied int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProtocolApplied >= len(resource.ProtocolApplied) {
		return StringInput("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].doseNumber", nil, htmlAttrs)
	}
	return StringInput("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].doseNumber", &resource.ProtocolApplied[numProtocolApplied].DoseNumber, htmlAttrs)
}
func (resource *Immunization) T_ProtocolAppliedSeriesDoses(numProtocolApplied int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProtocolApplied >= len(resource.ProtocolApplied) {
		return StringInput("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].seriesDoses", nil, htmlAttrs)
	}
	return StringInput("protocolApplied["+strconv.Itoa(numProtocolApplied)+"].seriesDoses", resource.ProtocolApplied[numProtocolApplied].SeriesDoses, htmlAttrs)
}
