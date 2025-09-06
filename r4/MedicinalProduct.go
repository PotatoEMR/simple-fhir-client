package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	IndicationCodeableConcept *CodeableConcept `json:"indicationCodeableConcept,omitempty"`
	IndicationReference       *Reference       `json:"indicationReference,omitempty"`
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

func (resource *MedicinalProduct) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProduct.Id", nil)
	}
	return StringInput("MedicinalProduct.Id", resource.Id)
}
func (resource *MedicinalProduct) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProduct.ImplicitRules", nil)
	}
	return StringInput("MedicinalProduct.ImplicitRules", resource.ImplicitRules)
}
func (resource *MedicinalProduct) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MedicinalProduct.Language", nil, optionsValueSet)
	}
	return CodeSelect("MedicinalProduct.Language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProduct) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProduct.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProduct.Type", resource.Type, optionsValueSet)
}
func (resource *MedicinalProduct) T_Domain(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("MedicinalProduct.Domain", nil, optionsValueSet)
	}
	return CodingSelect("MedicinalProduct.Domain", resource.Domain, optionsValueSet)
}
func (resource *MedicinalProduct) T_CombinedPharmaceuticalDoseForm(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProduct.CombinedPharmaceuticalDoseForm", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProduct.CombinedPharmaceuticalDoseForm", resource.CombinedPharmaceuticalDoseForm, optionsValueSet)
}
func (resource *MedicinalProduct) T_LegalStatusOfSupply(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProduct.LegalStatusOfSupply", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProduct.LegalStatusOfSupply", resource.LegalStatusOfSupply, optionsValueSet)
}
func (resource *MedicinalProduct) T_AdditionalMonitoringIndicator(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProduct.AdditionalMonitoringIndicator", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProduct.AdditionalMonitoringIndicator", resource.AdditionalMonitoringIndicator, optionsValueSet)
}
func (resource *MedicinalProduct) T_SpecialMeasures(numSpecialMeasures int) templ.Component {

	if resource == nil || len(resource.SpecialMeasures) >= numSpecialMeasures {
		return StringInput("MedicinalProduct.SpecialMeasures["+strconv.Itoa(numSpecialMeasures)+"]", nil)
	}
	return StringInput("MedicinalProduct.SpecialMeasures["+strconv.Itoa(numSpecialMeasures)+"]", &resource.SpecialMeasures[numSpecialMeasures])
}
func (resource *MedicinalProduct) T_PaediatricUseIndicator(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProduct.PaediatricUseIndicator", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProduct.PaediatricUseIndicator", resource.PaediatricUseIndicator, optionsValueSet)
}
func (resource *MedicinalProduct) T_ProductClassification(numProductClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ProductClassification) >= numProductClassification {
		return CodeableConceptSelect("MedicinalProduct.ProductClassification["+strconv.Itoa(numProductClassification)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProduct.ProductClassification["+strconv.Itoa(numProductClassification)+"]", &resource.ProductClassification[numProductClassification], optionsValueSet)
}
func (resource *MedicinalProduct) T_NameId(numName int) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return StringInput("MedicinalProduct.Name["+strconv.Itoa(numName)+"].Id", nil)
	}
	return StringInput("MedicinalProduct.Name["+strconv.Itoa(numName)+"].Id", resource.Name[numName].Id)
}
func (resource *MedicinalProduct) T_NameProductName(numName int) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return StringInput("MedicinalProduct.Name["+strconv.Itoa(numName)+"].ProductName", nil)
	}
	return StringInput("MedicinalProduct.Name["+strconv.Itoa(numName)+"].ProductName", &resource.Name[numName].ProductName)
}
func (resource *MedicinalProduct) T_NameNamePartId(numName int, numNamePart int) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].NamePart) >= numNamePart {
		return StringInput("MedicinalProduct.Name["+strconv.Itoa(numName)+"].NamePart["+strconv.Itoa(numNamePart)+"].Id", nil)
	}
	return StringInput("MedicinalProduct.Name["+strconv.Itoa(numName)+"].NamePart["+strconv.Itoa(numNamePart)+"].Id", resource.Name[numName].NamePart[numNamePart].Id)
}
func (resource *MedicinalProduct) T_NameNamePartPart(numName int, numNamePart int) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].NamePart) >= numNamePart {
		return StringInput("MedicinalProduct.Name["+strconv.Itoa(numName)+"].NamePart["+strconv.Itoa(numNamePart)+"].Part", nil)
	}
	return StringInput("MedicinalProduct.Name["+strconv.Itoa(numName)+"].NamePart["+strconv.Itoa(numNamePart)+"].Part", &resource.Name[numName].NamePart[numNamePart].Part)
}
func (resource *MedicinalProduct) T_NameNamePartType(numName int, numNamePart int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].NamePart) >= numNamePart {
		return CodingSelect("MedicinalProduct.Name["+strconv.Itoa(numName)+"].NamePart["+strconv.Itoa(numNamePart)+"].Type", nil, optionsValueSet)
	}
	return CodingSelect("MedicinalProduct.Name["+strconv.Itoa(numName)+"].NamePart["+strconv.Itoa(numNamePart)+"].Type", &resource.Name[numName].NamePart[numNamePart].Type, optionsValueSet)
}
func (resource *MedicinalProduct) T_NameCountryLanguageId(numName int, numCountryLanguage int) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].CountryLanguage) >= numCountryLanguage {
		return StringInput("MedicinalProduct.Name["+strconv.Itoa(numName)+"].CountryLanguage["+strconv.Itoa(numCountryLanguage)+"].Id", nil)
	}
	return StringInput("MedicinalProduct.Name["+strconv.Itoa(numName)+"].CountryLanguage["+strconv.Itoa(numCountryLanguage)+"].Id", resource.Name[numName].CountryLanguage[numCountryLanguage].Id)
}
func (resource *MedicinalProduct) T_NameCountryLanguageCountry(numName int, numCountryLanguage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].CountryLanguage) >= numCountryLanguage {
		return CodeableConceptSelect("MedicinalProduct.Name["+strconv.Itoa(numName)+"].CountryLanguage["+strconv.Itoa(numCountryLanguage)+"].Country", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProduct.Name["+strconv.Itoa(numName)+"].CountryLanguage["+strconv.Itoa(numCountryLanguage)+"].Country", &resource.Name[numName].CountryLanguage[numCountryLanguage].Country, optionsValueSet)
}
func (resource *MedicinalProduct) T_NameCountryLanguageJurisdiction(numName int, numCountryLanguage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].CountryLanguage) >= numCountryLanguage {
		return CodeableConceptSelect("MedicinalProduct.Name["+strconv.Itoa(numName)+"].CountryLanguage["+strconv.Itoa(numCountryLanguage)+"].Jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProduct.Name["+strconv.Itoa(numName)+"].CountryLanguage["+strconv.Itoa(numCountryLanguage)+"].Jurisdiction", resource.Name[numName].CountryLanguage[numCountryLanguage].Jurisdiction, optionsValueSet)
}
func (resource *MedicinalProduct) T_NameCountryLanguageLanguage(numName int, numCountryLanguage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].CountryLanguage) >= numCountryLanguage {
		return CodeableConceptSelect("MedicinalProduct.Name["+strconv.Itoa(numName)+"].CountryLanguage["+strconv.Itoa(numCountryLanguage)+"].Language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProduct.Name["+strconv.Itoa(numName)+"].CountryLanguage["+strconv.Itoa(numCountryLanguage)+"].Language", &resource.Name[numName].CountryLanguage[numCountryLanguage].Language, optionsValueSet)
}
func (resource *MedicinalProduct) T_ManufacturingBusinessOperationId(numManufacturingBusinessOperation int) templ.Component {

	if resource == nil || len(resource.ManufacturingBusinessOperation) >= numManufacturingBusinessOperation {
		return StringInput("MedicinalProduct.ManufacturingBusinessOperation["+strconv.Itoa(numManufacturingBusinessOperation)+"].Id", nil)
	}
	return StringInput("MedicinalProduct.ManufacturingBusinessOperation["+strconv.Itoa(numManufacturingBusinessOperation)+"].Id", resource.ManufacturingBusinessOperation[numManufacturingBusinessOperation].Id)
}
func (resource *MedicinalProduct) T_ManufacturingBusinessOperationOperationType(numManufacturingBusinessOperation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ManufacturingBusinessOperation) >= numManufacturingBusinessOperation {
		return CodeableConceptSelect("MedicinalProduct.ManufacturingBusinessOperation["+strconv.Itoa(numManufacturingBusinessOperation)+"].OperationType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProduct.ManufacturingBusinessOperation["+strconv.Itoa(numManufacturingBusinessOperation)+"].OperationType", resource.ManufacturingBusinessOperation[numManufacturingBusinessOperation].OperationType, optionsValueSet)
}
func (resource *MedicinalProduct) T_ManufacturingBusinessOperationEffectiveDate(numManufacturingBusinessOperation int) templ.Component {

	if resource == nil || len(resource.ManufacturingBusinessOperation) >= numManufacturingBusinessOperation {
		return StringInput("MedicinalProduct.ManufacturingBusinessOperation["+strconv.Itoa(numManufacturingBusinessOperation)+"].EffectiveDate", nil)
	}
	return StringInput("MedicinalProduct.ManufacturingBusinessOperation["+strconv.Itoa(numManufacturingBusinessOperation)+"].EffectiveDate", resource.ManufacturingBusinessOperation[numManufacturingBusinessOperation].EffectiveDate)
}
func (resource *MedicinalProduct) T_ManufacturingBusinessOperationConfidentialityIndicator(numManufacturingBusinessOperation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ManufacturingBusinessOperation) >= numManufacturingBusinessOperation {
		return CodeableConceptSelect("MedicinalProduct.ManufacturingBusinessOperation["+strconv.Itoa(numManufacturingBusinessOperation)+"].ConfidentialityIndicator", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProduct.ManufacturingBusinessOperation["+strconv.Itoa(numManufacturingBusinessOperation)+"].ConfidentialityIndicator", resource.ManufacturingBusinessOperation[numManufacturingBusinessOperation].ConfidentialityIndicator, optionsValueSet)
}
func (resource *MedicinalProduct) T_SpecialDesignationId(numSpecialDesignation int) templ.Component {

	if resource == nil || len(resource.SpecialDesignation) >= numSpecialDesignation {
		return StringInput("MedicinalProduct.SpecialDesignation["+strconv.Itoa(numSpecialDesignation)+"].Id", nil)
	}
	return StringInput("MedicinalProduct.SpecialDesignation["+strconv.Itoa(numSpecialDesignation)+"].Id", resource.SpecialDesignation[numSpecialDesignation].Id)
}
func (resource *MedicinalProduct) T_SpecialDesignationType(numSpecialDesignation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SpecialDesignation) >= numSpecialDesignation {
		return CodeableConceptSelect("MedicinalProduct.SpecialDesignation["+strconv.Itoa(numSpecialDesignation)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProduct.SpecialDesignation["+strconv.Itoa(numSpecialDesignation)+"].Type", resource.SpecialDesignation[numSpecialDesignation].Type, optionsValueSet)
}
func (resource *MedicinalProduct) T_SpecialDesignationIntendedUse(numSpecialDesignation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SpecialDesignation) >= numSpecialDesignation {
		return CodeableConceptSelect("MedicinalProduct.SpecialDesignation["+strconv.Itoa(numSpecialDesignation)+"].IntendedUse", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProduct.SpecialDesignation["+strconv.Itoa(numSpecialDesignation)+"].IntendedUse", resource.SpecialDesignation[numSpecialDesignation].IntendedUse, optionsValueSet)
}
func (resource *MedicinalProduct) T_SpecialDesignationStatus(numSpecialDesignation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SpecialDesignation) >= numSpecialDesignation {
		return CodeableConceptSelect("MedicinalProduct.SpecialDesignation["+strconv.Itoa(numSpecialDesignation)+"].Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProduct.SpecialDesignation["+strconv.Itoa(numSpecialDesignation)+"].Status", resource.SpecialDesignation[numSpecialDesignation].Status, optionsValueSet)
}
func (resource *MedicinalProduct) T_SpecialDesignationDate(numSpecialDesignation int) templ.Component {

	if resource == nil || len(resource.SpecialDesignation) >= numSpecialDesignation {
		return StringInput("MedicinalProduct.SpecialDesignation["+strconv.Itoa(numSpecialDesignation)+"].Date", nil)
	}
	return StringInput("MedicinalProduct.SpecialDesignation["+strconv.Itoa(numSpecialDesignation)+"].Date", resource.SpecialDesignation[numSpecialDesignation].Date)
}
func (resource *MedicinalProduct) T_SpecialDesignationSpecies(numSpecialDesignation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SpecialDesignation) >= numSpecialDesignation {
		return CodeableConceptSelect("MedicinalProduct.SpecialDesignation["+strconv.Itoa(numSpecialDesignation)+"].Species", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProduct.SpecialDesignation["+strconv.Itoa(numSpecialDesignation)+"].Species", resource.SpecialDesignation[numSpecialDesignation].Species, optionsValueSet)
}
