//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProduct
type MedicinalProduct struct {
	Id                             *string                                          `json:"id,omitempty"`
	Meta                           *Meta                                            `json:"meta,omitempty"`
	ImplicitRules                  *string                                          `json:"implicitRules,omitempty"`
	Language                       *string                                          `json:"language,omitempty"`
	Text                           *Narrative                                       `json:"text,omitempty"`
	Contained                      []Resource                                       `json:"contained,omitempty"`
	Extension                      []Extension                                      `json:"extension,omitempty"`
	ModifierExtension              []Extension                                      `json:"modifierExtension,omitempty"`
	Identifier                     []Identifier                                     `json:"identifier,omitempty"`
	Type                           *CodeableConcept                                 `json:"type,omitempty"`
	Domain                         *Coding                                          `json:"domain,omitempty"`
	CombinedPharmaceuticalDoseForm *CodeableConcept                                 `json:"combinedPharmaceuticalDoseForm,omitempty"`
	LegalStatusOfSupply            *CodeableConcept                                 `json:"legalStatusOfSupply,omitempty"`
	AdditionalMonitoringIndicator  *CodeableConcept                                 `json:"additionalMonitoringIndicator,omitempty"`
	SpecialMeasures                []string                                         `json:"specialMeasures,omitempty"`
	PaediatricUseIndicator         *CodeableConcept                                 `json:"paediatricUseIndicator,omitempty"`
	ProductClassification          []CodeableConcept                                `json:"productClassification,omitempty"`
	MarketingStatus                []MarketingStatus                                `json:"marketingStatus,omitempty"`
	PharmaceuticalProduct          []Reference                                      `json:"pharmaceuticalProduct,omitempty"`
	PackagedMedicinalProduct       []Reference                                      `json:"packagedMedicinalProduct,omitempty"`
	AttachedDocument               []Reference                                      `json:"attachedDocument,omitempty"`
	MasterFile                     []Reference                                      `json:"masterFile,omitempty"`
	Contact                        []Reference                                      `json:"contact,omitempty"`
	ClinicalTrial                  []Reference                                      `json:"clinicalTrial,omitempty"`
	Name                           []MedicinalProductName                           `json:"name"`
	CrossReference                 []Identifier                                     `json:"crossReference,omitempty"`
	ManufacturingBusinessOperation []MedicinalProductManufacturingBusinessOperation `json:"manufacturingBusinessOperation,omitempty"`
	SpecialDesignation             []MedicinalProductSpecialDesignation             `json:"specialDesignation,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProduct
type MedicinalProductName struct {
	Id                *string                               `json:"id,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	ProductName       string                                `json:"productName"`
	NamePart          []MedicinalProductNameNamePart        `json:"namePart,omitempty"`
	CountryLanguage   []MedicinalProductNameCountryLanguage `json:"countryLanguage,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProduct
type MedicinalProductNameNamePart struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Part              string      `json:"part"`
	Type              Coding      `json:"type"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProduct
type MedicinalProductNameCountryLanguage struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Country           CodeableConcept  `json:"country"`
	Jurisdiction      *CodeableConcept `json:"jurisdiction,omitempty"`
	Language          CodeableConcept  `json:"language"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProduct
type MedicinalProductManufacturingBusinessOperation struct {
	Id                           *string          `json:"id,omitempty"`
	Extension                    []Extension      `json:"extension,omitempty"`
	ModifierExtension            []Extension      `json:"modifierExtension,omitempty"`
	OperationType                *CodeableConcept `json:"operationType,omitempty"`
	AuthorisationReferenceNumber *Identifier      `json:"authorisationReferenceNumber,omitempty"`
	EffectiveDate                *string          `json:"effectiveDate,omitempty"`
	ConfidentialityIndicator     *CodeableConcept `json:"confidentialityIndicator,omitempty"`
	Manufacturer                 []Reference      `json:"manufacturer,omitempty"`
	Regulator                    *Reference       `json:"regulator,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProduct
type MedicinalProductSpecialDesignation struct {
	Id                        *string          `json:"id,omitempty"`
	Extension                 []Extension      `json:"extension,omitempty"`
	ModifierExtension         []Extension      `json:"modifierExtension,omitempty"`
	Identifier                []Identifier     `json:"identifier,omitempty"`
	Type                      *CodeableConcept `json:"type,omitempty"`
	IntendedUse               *CodeableConcept `json:"intendedUse,omitempty"`
	IndicationCodeableConcept *CodeableConcept `json:"indicationCodeableConcept"`
	IndicationReference       *Reference       `json:"indicationReference"`
	Status                    *CodeableConcept `json:"status,omitempty"`
	Date                      *string          `json:"date,omitempty"`
	Species                   *CodeableConcept `json:"species,omitempty"`
}

type OtherMedicinalProduct MedicinalProduct

// on convert struct to json, automatically add resourceType=MedicinalProduct
func (r MedicinalProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProduct
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProduct: OtherMedicinalProduct(r),
		ResourceType:          "MedicinalProduct",
	})
}
