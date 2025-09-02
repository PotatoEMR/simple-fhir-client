package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *MedicinalProductAuthorization) MedicinalProductAuthorizationLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductAuthorization) MedicinalProductAuthorizationCountry(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("country", nil, optionsValueSet)
	}
	return CodeableConceptSelect("country", &resource.Country[0], optionsValueSet)
}
func (resource *MedicinalProductAuthorization) MedicinalProductAuthorizationJurisdiction(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.Jurisdiction[0], optionsValueSet)
}
func (resource *MedicinalProductAuthorization) MedicinalProductAuthorizationStatus(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", resource.Status, optionsValueSet)
}
func (resource *MedicinalProductAuthorization) MedicinalProductAuthorizationLegalBasis(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("legalBasis", nil, optionsValueSet)
	}
	return CodeableConceptSelect("legalBasis", resource.LegalBasis, optionsValueSet)
}
func (resource *MedicinalProductAuthorization) MedicinalProductAuthorizationJurisdictionalAuthorizationCountry(numJurisdictionalAuthorization int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.JurisdictionalAuthorization) >= numJurisdictionalAuthorization {
		return CodeableConceptSelect("country", nil, optionsValueSet)
	}
	return CodeableConceptSelect("country", resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].Country, optionsValueSet)
}
func (resource *MedicinalProductAuthorization) MedicinalProductAuthorizationJurisdictionalAuthorizationJurisdiction(numJurisdictionalAuthorization int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.JurisdictionalAuthorization) >= numJurisdictionalAuthorization {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", &resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].Jurisdiction[0], optionsValueSet)
}
func (resource *MedicinalProductAuthorization) MedicinalProductAuthorizationJurisdictionalAuthorizationLegalStatusOfSupply(numJurisdictionalAuthorization int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.JurisdictionalAuthorization) >= numJurisdictionalAuthorization {
		return CodeableConceptSelect("legalStatusOfSupply", nil, optionsValueSet)
	}
	return CodeableConceptSelect("legalStatusOfSupply", resource.JurisdictionalAuthorization[numJurisdictionalAuthorization].LegalStatusOfSupply, optionsValueSet)
}
func (resource *MedicinalProductAuthorization) MedicinalProductAuthorizationProcedureType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Procedure.Type, optionsValueSet)
}
