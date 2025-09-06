package r4

//generated with command go run ./bultaoreune -nodownload
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
	StatusDate                  *string                                                    `json:"statusDate,omitempty"`
	RestoreDate                 *string                                                    `json:"restoreDate,omitempty"`
	ValidityPeriod              *Period                                                    `json:"validityPeriod,omitempty"`
	DataExclusivityPeriod       *Period                                                    `json:"dataExclusivityPeriod,omitempty"`
	DateOfFirstAuthorization    *string                                                    `json:"dateOfFirstAuthorization,omitempty"`
	InternationalBirthDate      *string                                                    `json:"internationalBirthDate,omitempty"`
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
	DateDateTime      *string         `json:"dateDateTime,omitempty"`
}

type OtherMedicinalProductAuthorization MedicinalProductAuthorization

// on convert struct to json, automatically add resourceType=MedicinalProductAuthorization
func (r MedicinalProductAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductAuthorization
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductAuthorization: OtherMedicinalProductAuthorization(r),
		ResourceType:                       "MedicinalProductAuthorization",
	})
}

func (resource *MedicinalProductAuthorization) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductAuthorization.Id", nil)
	}
	return StringInput("MedicinalProductAuthorization.Id", resource.Id)
}
func (resource *MedicinalProductAuthorization) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductAuthorization.ImplicitRules", nil)
	}
	return StringInput("MedicinalProductAuthorization.ImplicitRules", resource.ImplicitRules)
}
func (resource *MedicinalProductAuthorization) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MedicinalProductAuthorization.Language", nil, optionsValueSet)
	}
	return CodeSelect("MedicinalProductAuthorization.Language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductAuthorization) T_Country(numCountry int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Country) >= numCountry {
		return CodeableConceptSelect("MedicinalProductAuthorization.Country["+strconv.Itoa(numCountry)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductAuthorization.Country["+strconv.Itoa(numCountry)+"]", &resource.Country[numCountry], optionsValueSet)
}
func (resource *MedicinalProductAuthorization) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("MedicinalProductAuthorization.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductAuthorization.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *MedicinalProductAuthorization) T_Status(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductAuthorization.Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductAuthorization.Status", resource.Status, optionsValueSet)
}
func (resource *MedicinalProductAuthorization) T_StatusDate() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductAuthorization.StatusDate", nil)
	}
	return StringInput("MedicinalProductAuthorization.StatusDate", resource.StatusDate)
}
func (resource *MedicinalProductAuthorization) T_RestoreDate() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductAuthorization.RestoreDate", nil)
	}
	return StringInput("MedicinalProductAuthorization.RestoreDate", resource.RestoreDate)
}
func (resource *MedicinalProductAuthorization) T_DateOfFirstAuthorization() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductAuthorization.DateOfFirstAuthorization", nil)
	}
	return StringInput("MedicinalProductAuthorization.DateOfFirstAuthorization", resource.DateOfFirstAuthorization)
}
func (resource *MedicinalProductAuthorization) T_InternationalBirthDate() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductAuthorization.InternationalBirthDate", nil)
	}
	return StringInput("MedicinalProductAuthorization.InternationalBirthDate", resource.InternationalBirthDate)
}
func (resource *MedicinalProductAuthorization) T_LegalBasis(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductAuthorization.LegalBasis", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductAuthorization.LegalBasis", resource.LegalBasis, optionsValueSet)
}
func (resource *MedicinalProductAuthorization) T_JurisdictionalAuthorizationId(numJurisdictionalAuthorization int) templ.Component {

	if resource == nil || len(resource.JurisdictionalAuthorization) >= numJurisdictionalAuthorization {
		return StringInput("MedicinalProductAuthorization.JurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].Id", nil)
	}
	return StringInput("MedicinalProductAuthorization.JurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].Id", resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].Id)
}
func (resource *MedicinalProductAuthorization) T_JurisdictionalAuthorizationCountry(numJurisdictionalAuthorization int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.JurisdictionalAuthorization) >= numJurisdictionalAuthorization {
		return CodeableConceptSelect("MedicinalProductAuthorization.JurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].Country", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductAuthorization.JurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].Country", resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].Country, optionsValueSet)
}
func (resource *MedicinalProductAuthorization) T_JurisdictionalAuthorizationJurisdiction(numJurisdictionalAuthorization int, numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.JurisdictionalAuthorization) >= numJurisdictionalAuthorization || len(resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("MedicinalProductAuthorization.JurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductAuthorization.JurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *MedicinalProductAuthorization) T_JurisdictionalAuthorizationLegalStatusOfSupply(numJurisdictionalAuthorization int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.JurisdictionalAuthorization) >= numJurisdictionalAuthorization {
		return CodeableConceptSelect("MedicinalProductAuthorization.JurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].LegalStatusOfSupply", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductAuthorization.JurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].LegalStatusOfSupply", resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].LegalStatusOfSupply, optionsValueSet)
}
func (resource *MedicinalProductAuthorization) T_ProcedureId() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductAuthorization.Procedure.Id", nil)
	}
	return StringInput("MedicinalProductAuthorization.Procedure.Id", resource.Procedure.Id)
}
func (resource *MedicinalProductAuthorization) T_ProcedureType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductAuthorization.Procedure.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductAuthorization.Procedure.Type", &resource.Procedure.Type, optionsValueSet)
}
