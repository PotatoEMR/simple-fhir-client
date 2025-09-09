package r4

//generated with command go run ./bultaoreune
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
func (r MedicinalProduct) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicinalProduct/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "MedicinalProduct"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MedicinalProduct) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProduct) T_Domain(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("domain", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("domain", resource.Domain, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProduct) T_CombinedPharmaceuticalDoseForm(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("combinedPharmaceuticalDoseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("combinedPharmaceuticalDoseForm", resource.CombinedPharmaceuticalDoseForm, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProduct) T_LegalStatusOfSupply(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("legalStatusOfSupply", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("legalStatusOfSupply", resource.LegalStatusOfSupply, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProduct) T_AdditionalMonitoringIndicator(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("additionalMonitoringIndicator", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("additionalMonitoringIndicator", resource.AdditionalMonitoringIndicator, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProduct) T_SpecialMeasures(numSpecialMeasures int, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialMeasures >= len(resource.SpecialMeasures) {
		return StringInput("specialMeasures["+strconv.Itoa(numSpecialMeasures)+"]", nil, htmlAttrs)
	}
	return StringInput("specialMeasures["+strconv.Itoa(numSpecialMeasures)+"]", &resource.SpecialMeasures[numSpecialMeasures], htmlAttrs)
}
func (resource *MedicinalProduct) T_PaediatricUseIndicator(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("paediatricUseIndicator", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("paediatricUseIndicator", resource.PaediatricUseIndicator, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProduct) T_ProductClassification(numProductClassification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProductClassification >= len(resource.ProductClassification) {
		return CodeableConceptSelect("productClassification["+strconv.Itoa(numProductClassification)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("productClassification["+strconv.Itoa(numProductClassification)+"]", &resource.ProductClassification[numProductClassification], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProduct) T_NameProductName(numName int, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return StringInput("name["+strconv.Itoa(numName)+"].productName", nil, htmlAttrs)
	}
	return StringInput("name["+strconv.Itoa(numName)+"].productName", &resource.Name[numName].ProductName, htmlAttrs)
}
func (resource *MedicinalProduct) T_NameNamePartPart(numName int, numNamePart int, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numNamePart >= len(resource.Name[numName].NamePart) {
		return StringInput("name["+strconv.Itoa(numName)+"].namePart["+strconv.Itoa(numNamePart)+"].part", nil, htmlAttrs)
	}
	return StringInput("name["+strconv.Itoa(numName)+"].namePart["+strconv.Itoa(numNamePart)+"].part", &resource.Name[numName].NamePart[numNamePart].Part, htmlAttrs)
}
func (resource *MedicinalProduct) T_NameNamePartType(numName int, numNamePart int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numNamePart >= len(resource.Name[numName].NamePart) {
		return CodingSelect("name["+strconv.Itoa(numName)+"].namePart["+strconv.Itoa(numNamePart)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("name["+strconv.Itoa(numName)+"].namePart["+strconv.Itoa(numNamePart)+"].type", &resource.Name[numName].NamePart[numNamePart].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProduct) T_NameCountryLanguageCountry(numName int, numCountryLanguage int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numCountryLanguage >= len(resource.Name[numName].CountryLanguage) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].countryLanguage["+strconv.Itoa(numCountryLanguage)+"].country", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].countryLanguage["+strconv.Itoa(numCountryLanguage)+"].country", &resource.Name[numName].CountryLanguage[numCountryLanguage].Country, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProduct) T_NameCountryLanguageJurisdiction(numName int, numCountryLanguage int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numCountryLanguage >= len(resource.Name[numName].CountryLanguage) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].countryLanguage["+strconv.Itoa(numCountryLanguage)+"].jurisdiction", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].countryLanguage["+strconv.Itoa(numCountryLanguage)+"].jurisdiction", resource.Name[numName].CountryLanguage[numCountryLanguage].Jurisdiction, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProduct) T_ManufacturingBusinessOperationOperationType(numManufacturingBusinessOperation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numManufacturingBusinessOperation >= len(resource.ManufacturingBusinessOperation) {
		return CodeableConceptSelect("manufacturingBusinessOperation["+strconv.Itoa(numManufacturingBusinessOperation)+"].operationType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("manufacturingBusinessOperation["+strconv.Itoa(numManufacturingBusinessOperation)+"].operationType", resource.ManufacturingBusinessOperation[numManufacturingBusinessOperation].OperationType, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProduct) T_ManufacturingBusinessOperationEffectiveDate(numManufacturingBusinessOperation int, htmlAttrs string) templ.Component {
	if resource == nil || numManufacturingBusinessOperation >= len(resource.ManufacturingBusinessOperation) {
		return DateTimeInput("manufacturingBusinessOperation["+strconv.Itoa(numManufacturingBusinessOperation)+"].effectiveDate", nil, htmlAttrs)
	}
	return DateTimeInput("manufacturingBusinessOperation["+strconv.Itoa(numManufacturingBusinessOperation)+"].effectiveDate", resource.ManufacturingBusinessOperation[numManufacturingBusinessOperation].EffectiveDate, htmlAttrs)
}
func (resource *MedicinalProduct) T_ManufacturingBusinessOperationConfidentialityIndicator(numManufacturingBusinessOperation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numManufacturingBusinessOperation >= len(resource.ManufacturingBusinessOperation) {
		return CodeableConceptSelect("manufacturingBusinessOperation["+strconv.Itoa(numManufacturingBusinessOperation)+"].confidentialityIndicator", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("manufacturingBusinessOperation["+strconv.Itoa(numManufacturingBusinessOperation)+"].confidentialityIndicator", resource.ManufacturingBusinessOperation[numManufacturingBusinessOperation].ConfidentialityIndicator, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProduct) T_SpecialDesignationType(numSpecialDesignation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialDesignation >= len(resource.SpecialDesignation) {
		return CodeableConceptSelect("specialDesignation["+strconv.Itoa(numSpecialDesignation)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("specialDesignation["+strconv.Itoa(numSpecialDesignation)+"].type", resource.SpecialDesignation[numSpecialDesignation].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProduct) T_SpecialDesignationIntendedUse(numSpecialDesignation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialDesignation >= len(resource.SpecialDesignation) {
		return CodeableConceptSelect("specialDesignation["+strconv.Itoa(numSpecialDesignation)+"].intendedUse", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("specialDesignation["+strconv.Itoa(numSpecialDesignation)+"].intendedUse", resource.SpecialDesignation[numSpecialDesignation].IntendedUse, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProduct) T_SpecialDesignationIndicationCodeableConcept(numSpecialDesignation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialDesignation >= len(resource.SpecialDesignation) {
		return CodeableConceptSelect("specialDesignation["+strconv.Itoa(numSpecialDesignation)+"].indicationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("specialDesignation["+strconv.Itoa(numSpecialDesignation)+"].indicationCodeableConcept", resource.SpecialDesignation[numSpecialDesignation].IndicationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProduct) T_SpecialDesignationStatus(numSpecialDesignation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialDesignation >= len(resource.SpecialDesignation) {
		return CodeableConceptSelect("specialDesignation["+strconv.Itoa(numSpecialDesignation)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("specialDesignation["+strconv.Itoa(numSpecialDesignation)+"].status", resource.SpecialDesignation[numSpecialDesignation].Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProduct) T_SpecialDesignationDate(numSpecialDesignation int, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialDesignation >= len(resource.SpecialDesignation) {
		return DateTimeInput("specialDesignation["+strconv.Itoa(numSpecialDesignation)+"].date", nil, htmlAttrs)
	}
	return DateTimeInput("specialDesignation["+strconv.Itoa(numSpecialDesignation)+"].date", resource.SpecialDesignation[numSpecialDesignation].Date, htmlAttrs)
}
func (resource *MedicinalProduct) T_SpecialDesignationSpecies(numSpecialDesignation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialDesignation >= len(resource.SpecialDesignation) {
		return CodeableConceptSelect("specialDesignation["+strconv.Itoa(numSpecialDesignation)+"].species", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("specialDesignation["+strconv.Itoa(numSpecialDesignation)+"].species", resource.SpecialDesignation[numSpecialDesignation].Species, optionsValueSet, htmlAttrs)
}
