package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductAuthorization
type MedicinalProductAuthorization struct {
	Id                          *string                                                    `json:"id,omitempty"`
	Meta                        *Meta                                                      `json:"meta,omitempty"`
	ImplicitRules               *string                                                    `json:"implicitRules,omitempty"`
	Language                    *string                                                    `json:"language,omitempty"`
	Text                        *Narrative                                                 `json:"text,omitempty"`
	Contained                   []Resource                                                 `json:"contained,omitempty"`
	Extension                   []Extension                                                `json:"extension,omitempty"`
	ModifierExtension           []Extension                                                `json:"modifierExtension,omitempty"`
	Identifier                  []Identifier                                               `json:"identifier,omitempty"`
	Subject                     *Reference                                                 `json:"subject,omitempty"`
	Country                     []CodeableConcept                                          `json:"country,omitempty"`
	Jurisdiction                []CodeableConcept                                          `json:"jurisdiction,omitempty"`
	Status                      *CodeableConcept                                           `json:"status,omitempty"`
	StatusDate                  *FhirDateTime                                              `json:"statusDate,omitempty"`
	RestoreDate                 *FhirDateTime                                              `json:"restoreDate,omitempty"`
	ValidityPeriod              *Period                                                    `json:"validityPeriod,omitempty"`
	DataExclusivityPeriod       *Period                                                    `json:"dataExclusivityPeriod,omitempty"`
	DateOfFirstAuthorization    *FhirDateTime                                              `json:"dateOfFirstAuthorization,omitempty"`
	InternationalBirthDate      *FhirDateTime                                              `json:"internationalBirthDate,omitempty"`
	LegalBasis                  *CodeableConcept                                           `json:"legalBasis,omitempty"`
	JurisdictionalAuthorization []MedicinalProductAuthorizationJurisdictionalAuthorization `json:"jurisdictionalAuthorization,omitempty"`
	Holder                      *Reference                                                 `json:"holder,omitempty"`
	Regulator                   *Reference                                                 `json:"regulator,omitempty"`
	Procedure                   *MedicinalProductAuthorizationProcedure                    `json:"procedure,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductAuthorization
type MedicinalProductAuthorizationJurisdictionalAuthorization struct {
	Id                  *string           `json:"id,omitempty"`
	Extension           []Extension       `json:"extension,omitempty"`
	ModifierExtension   []Extension       `json:"modifierExtension,omitempty"`
	Identifier          []Identifier      `json:"identifier,omitempty"`
	Country             *CodeableConcept  `json:"country,omitempty"`
	Jurisdiction        []CodeableConcept `json:"jurisdiction,omitempty"`
	LegalStatusOfSupply *CodeableConcept  `json:"legalStatusOfSupply,omitempty"`
	ValidityPeriod      *Period           `json:"validityPeriod,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductAuthorization
type MedicinalProductAuthorizationProcedure struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Identifier        *Identifier     `json:"identifier,omitempty"`
	Type              CodeableConcept `json:"type"`
	DatePeriod        *Period         `json:"datePeriod,omitempty"`
	DateDateTime      *FhirDateTime   `json:"dateDateTime,omitempty"`
}

type OtherMedicinalProductAuthorization MedicinalProductAuthorization

// struct -> json, automatically add resourceType=Patient
func (r MedicinalProductAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductAuthorization
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductAuthorization: OtherMedicinalProductAuthorization(r),
		ResourceType:                       "MedicinalProductAuthorization",
	})
}

func (r MedicinalProductAuthorization) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicinalProductAuthorization/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "MedicinalProductAuthorization"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r MedicinalProductAuthorization) ResourceType() string {
	return "MedicinalProductAuthorization"
}

func (resource *MedicinalProductAuthorization) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_Country(numCountry int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCountry >= len(resource.Country) {
		return CodeableConceptSelect("country["+strconv.Itoa(numCountry)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("country["+strconv.Itoa(numCountry)+"]", &resource.Country[numCountry], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_Status(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_StatusDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("statusDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("statusDate", resource.StatusDate, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_RestoreDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("restoreDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("restoreDate", resource.RestoreDate, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_ValidityPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("validityPeriod", nil, htmlAttrs)
	}
	return PeriodInput("validityPeriod", resource.ValidityPeriod, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_DataExclusivityPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("dataExclusivityPeriod", nil, htmlAttrs)
	}
	return PeriodInput("dataExclusivityPeriod", resource.DataExclusivityPeriod, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_DateOfFirstAuthorization(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("dateOfFirstAuthorization", nil, htmlAttrs)
	}
	return FhirDateTimeInput("dateOfFirstAuthorization", resource.DateOfFirstAuthorization, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_InternationalBirthDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("internationalBirthDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("internationalBirthDate", resource.InternationalBirthDate, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_LegalBasis(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("legalBasis", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("legalBasis", resource.LegalBasis, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_Holder(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "holder", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "holder", resource.Holder, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_Regulator(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "regulator", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "regulator", resource.Regulator, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_JurisdictionalAuthorizationCountry(numJurisdictionalAuthorization int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdictionalAuthorization >= len(resource.JurisdictionalAuthorization) {
		return CodeableConceptSelect("jurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].country", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].country", resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].Country, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_JurisdictionalAuthorizationJurisdiction(numJurisdictionalAuthorization int, numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdictionalAuthorization >= len(resource.JurisdictionalAuthorization) || numJurisdiction >= len(resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].Jurisdiction) {
		return CodeableConceptSelect("jurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_JurisdictionalAuthorizationLegalStatusOfSupply(numJurisdictionalAuthorization int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdictionalAuthorization >= len(resource.JurisdictionalAuthorization) {
		return CodeableConceptSelect("jurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].legalStatusOfSupply", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].legalStatusOfSupply", resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].LegalStatusOfSupply, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_JurisdictionalAuthorizationValidityPeriod(numJurisdictionalAuthorization int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdictionalAuthorization >= len(resource.JurisdictionalAuthorization) {
		return PeriodInput("jurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].validityPeriod", nil, htmlAttrs)
	}
	return PeriodInput("jurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].validityPeriod", resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].ValidityPeriod, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_ProcedureType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("procedure.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("procedure.type", &resource.Procedure.Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_ProcedureDatePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("procedure.datePeriod", nil, htmlAttrs)
	}
	return PeriodInput("procedure.datePeriod", resource.Procedure.DatePeriod, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_ProcedureDateDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("procedure.dateDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("procedure.dateDateTime", resource.Procedure.DateDateTime, htmlAttrs)
}
