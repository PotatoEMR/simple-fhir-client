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
	StatusDate                  *time.Time                                                 `json:"statusDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	RestoreDate                 *time.Time                                                 `json:"restoreDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	ValidityPeriod              *Period                                                    `json:"validityPeriod,omitempty"`
	DataExclusivityPeriod       *Period                                                    `json:"dataExclusivityPeriod,omitempty"`
	DateOfFirstAuthorization    *time.Time                                                 `json:"dateOfFirstAuthorization,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	InternationalBirthDate      *time.Time                                                 `json:"internationalBirthDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	DateDateTime      *time.Time      `json:"dateDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *MedicinalProductAuthorization) T_Country(numCountry int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCountry >= len(resource.Country) {
		return CodeableConceptSelect("MedicinalProductAuthorization.Country["+strconv.Itoa(numCountry)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductAuthorization.Country["+strconv.Itoa(numCountry)+"]", &resource.Country[numCountry], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("MedicinalProductAuthorization.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductAuthorization.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_Status(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("MedicinalProductAuthorization.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductAuthorization.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_StatusDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("MedicinalProductAuthorization.StatusDate", nil, htmlAttrs)
	}
	return DateTimeInput("MedicinalProductAuthorization.StatusDate", resource.StatusDate, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_RestoreDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("MedicinalProductAuthorization.RestoreDate", nil, htmlAttrs)
	}
	return DateTimeInput("MedicinalProductAuthorization.RestoreDate", resource.RestoreDate, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_DateOfFirstAuthorization(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("MedicinalProductAuthorization.DateOfFirstAuthorization", nil, htmlAttrs)
	}
	return DateTimeInput("MedicinalProductAuthorization.DateOfFirstAuthorization", resource.DateOfFirstAuthorization, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_InternationalBirthDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("MedicinalProductAuthorization.InternationalBirthDate", nil, htmlAttrs)
	}
	return DateTimeInput("MedicinalProductAuthorization.InternationalBirthDate", resource.InternationalBirthDate, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_LegalBasis(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("MedicinalProductAuthorization.LegalBasis", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductAuthorization.LegalBasis", resource.LegalBasis, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_JurisdictionalAuthorizationCountry(numJurisdictionalAuthorization int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdictionalAuthorization >= len(resource.JurisdictionalAuthorization) {
		return CodeableConceptSelect("MedicinalProductAuthorization.JurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].Country", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductAuthorization.JurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].Country", resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].Country, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_JurisdictionalAuthorizationJurisdiction(numJurisdictionalAuthorization int, numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdictionalAuthorization >= len(resource.JurisdictionalAuthorization) || numJurisdiction >= len(resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].Jurisdiction) {
		return CodeableConceptSelect("MedicinalProductAuthorization.JurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductAuthorization.JurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_JurisdictionalAuthorizationLegalStatusOfSupply(numJurisdictionalAuthorization int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdictionalAuthorization >= len(resource.JurisdictionalAuthorization) {
		return CodeableConceptSelect("MedicinalProductAuthorization.JurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].LegalStatusOfSupply", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductAuthorization.JurisdictionalAuthorization["+strconv.Itoa(numJurisdictionalAuthorization)+"].LegalStatusOfSupply", resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].LegalStatusOfSupply, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_ProcedureType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("MedicinalProductAuthorization.Procedure.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductAuthorization.Procedure.Type", &resource.Procedure.Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductAuthorization) T_ProcedureDateDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("MedicinalProductAuthorization.Procedure.DateDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("MedicinalProductAuthorization.Procedure.DateDateTime", resource.Procedure.DateDateTime, htmlAttrs)
}
