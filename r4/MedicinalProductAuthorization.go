//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

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
	DatePeriod        *Period         `json:"datePeriod"`
	DateDateTime      *string         `json:"dateDateTime"`
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
