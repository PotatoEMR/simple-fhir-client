package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *MedicinalProduct) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProduct) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *MedicinalProduct) T_Domain(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("domain", nil, optionsValueSet)
	}
	return CodingSelect("domain", resource.Domain, optionsValueSet)
}
func (resource *MedicinalProduct) T_CombinedPharmaceuticalDoseForm(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("combinedPharmaceuticalDoseForm", nil, optionsValueSet)
	}
	return CodeableConceptSelect("combinedPharmaceuticalDoseForm", resource.CombinedPharmaceuticalDoseForm, optionsValueSet)
}
func (resource *MedicinalProduct) T_LegalStatusOfSupply(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("legalStatusOfSupply", nil, optionsValueSet)
	}
	return CodeableConceptSelect("legalStatusOfSupply", resource.LegalStatusOfSupply, optionsValueSet)
}
func (resource *MedicinalProduct) T_AdditionalMonitoringIndicator(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("additionalMonitoringIndicator", nil, optionsValueSet)
	}
	return CodeableConceptSelect("additionalMonitoringIndicator", resource.AdditionalMonitoringIndicator, optionsValueSet)
}
func (resource *MedicinalProduct) T_PaediatricUseIndicator(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("paediatricUseIndicator", nil, optionsValueSet)
	}
	return CodeableConceptSelect("paediatricUseIndicator", resource.PaediatricUseIndicator, optionsValueSet)
}
func (resource *MedicinalProduct) T_ProductClassification(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("productClassification", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productClassification", &resource.ProductClassification[0], optionsValueSet)
}
func (resource *MedicinalProduct) T_NameNamePartType(numName int, numNamePart int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name[numName].NamePart) >= numNamePart {
		return CodingSelect("type", nil, optionsValueSet)
	}
	return CodingSelect("type", &resource.Name[numName].NamePart[numNamePart].Type, optionsValueSet)
}
func (resource *MedicinalProduct) T_NameCountryLanguageCountry(numName int, numCountryLanguage int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name[numName].CountryLanguage) >= numCountryLanguage {
		return CodeableConceptSelect("country", nil, optionsValueSet)
	}
	return CodeableConceptSelect("country", &resource.Name[numName].CountryLanguage[numCountryLanguage].Country, optionsValueSet)
}
func (resource *MedicinalProduct) T_NameCountryLanguageJurisdiction(numName int, numCountryLanguage int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name[numName].CountryLanguage) >= numCountryLanguage {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", resource.Name[numName].CountryLanguage[numCountryLanguage].Jurisdiction, optionsValueSet)
}
func (resource *MedicinalProduct) T_NameCountryLanguageLanguage(numName int, numCountryLanguage int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name[numName].CountryLanguage) >= numCountryLanguage {
		return CodeableConceptSelect("language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("language", &resource.Name[numName].CountryLanguage[numCountryLanguage].Language, optionsValueSet)
}
func (resource *MedicinalProduct) T_ManufacturingBusinessOperationOperationType(numManufacturingBusinessOperation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.ManufacturingBusinessOperation) >= numManufacturingBusinessOperation {
		return CodeableConceptSelect("operationType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("operationType", resource.ManufacturingBusinessOperation[numManufacturingBusinessOperation].OperationType, optionsValueSet)
}
func (resource *MedicinalProduct) T_ManufacturingBusinessOperationConfidentialityIndicator(numManufacturingBusinessOperation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.ManufacturingBusinessOperation) >= numManufacturingBusinessOperation {
		return CodeableConceptSelect("confidentialityIndicator", nil, optionsValueSet)
	}
	return CodeableConceptSelect("confidentialityIndicator", resource.ManufacturingBusinessOperation[numManufacturingBusinessOperation].ConfidentialityIndicator, optionsValueSet)
}
func (resource *MedicinalProduct) T_SpecialDesignationType(numSpecialDesignation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.SpecialDesignation) >= numSpecialDesignation {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.SpecialDesignation[numSpecialDesignation].Type, optionsValueSet)
}
func (resource *MedicinalProduct) T_SpecialDesignationIntendedUse(numSpecialDesignation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.SpecialDesignation) >= numSpecialDesignation {
		return CodeableConceptSelect("intendedUse", nil, optionsValueSet)
	}
	return CodeableConceptSelect("intendedUse", resource.SpecialDesignation[numSpecialDesignation].IntendedUse, optionsValueSet)
}
func (resource *MedicinalProduct) T_SpecialDesignationStatus(numSpecialDesignation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.SpecialDesignation) >= numSpecialDesignation {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", resource.SpecialDesignation[numSpecialDesignation].Status, optionsValueSet)
}
func (resource *MedicinalProduct) T_SpecialDesignationSpecies(numSpecialDesignation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.SpecialDesignation) >= numSpecialDesignation {
		return CodeableConceptSelect("species", nil, optionsValueSet)
	}
	return CodeableConceptSelect("species", resource.SpecialDesignation[numSpecialDesignation].Species, optionsValueSet)
}
