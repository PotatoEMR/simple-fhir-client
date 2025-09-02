package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinition struct {
	Id                             *string                                    `json:"id,omitempty"`
	Meta                           *Meta                                      `json:"meta,omitempty"`
	ImplicitRules                  *string                                    `json:"implicitRules,omitempty"`
	Language                       *string                                    `json:"language,omitempty"`
	Text                           *Narrative                                 `json:"text,omitempty"`
	Contained                      []Resource                                 `json:"contained,omitempty"`
	Extension                      []Extension                                `json:"extension,omitempty"`
	ModifierExtension              []Extension                                `json:"modifierExtension,omitempty"`
	Identifier                     []Identifier                               `json:"identifier,omitempty"`
	Type                           *CodeableConcept                           `json:"type,omitempty"`
	Domain                         *CodeableConcept                           `json:"domain,omitempty"`
	Version                        *string                                    `json:"version,omitempty"`
	Status                         *CodeableConcept                           `json:"status,omitempty"`
	StatusDate                     *string                                    `json:"statusDate,omitempty"`
	Description                    *string                                    `json:"description,omitempty"`
	CombinedPharmaceuticalDoseForm *CodeableConcept                           `json:"combinedPharmaceuticalDoseForm,omitempty"`
	Route                          []CodeableConcept                          `json:"route,omitempty"`
	Indication                     *string                                    `json:"indication,omitempty"`
	LegalStatusOfSupply            *CodeableConcept                           `json:"legalStatusOfSupply,omitempty"`
	AdditionalMonitoringIndicator  *CodeableConcept                           `json:"additionalMonitoringIndicator,omitempty"`
	SpecialMeasures                []CodeableConcept                          `json:"specialMeasures,omitempty"`
	PediatricUseIndicator          *CodeableConcept                           `json:"pediatricUseIndicator,omitempty"`
	Classification                 []CodeableConcept                          `json:"classification,omitempty"`
	MarketingStatus                []MarketingStatus                          `json:"marketingStatus,omitempty"`
	PackagedMedicinalProduct       []CodeableConcept                          `json:"packagedMedicinalProduct,omitempty"`
	Ingredient                     []CodeableConcept                          `json:"ingredient,omitempty"`
	Impurity                       []CodeableReference                        `json:"impurity,omitempty"`
	AttachedDocument               []Reference                                `json:"attachedDocument,omitempty"`
	MasterFile                     []Reference                                `json:"masterFile,omitempty"`
	Contact                        []MedicinalProductDefinitionContact        `json:"contact,omitempty"`
	ClinicalTrial                  []Reference                                `json:"clinicalTrial,omitempty"`
	Code                           []Coding                                   `json:"code,omitempty"`
	Name                           []MedicinalProductDefinitionName           `json:"name"`
	CrossReference                 []MedicinalProductDefinitionCrossReference `json:"crossReference,omitempty"`
	Operation                      []MedicinalProductDefinitionOperation      `json:"operation,omitempty"`
	Characteristic                 []MedicinalProductDefinitionCharacteristic `json:"characteristic,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinitionContact struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Contact           Reference        `json:"contact"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinitionName struct {
	Id                *string                                         `json:"id,omitempty"`
	Extension         []Extension                                     `json:"extension,omitempty"`
	ModifierExtension []Extension                                     `json:"modifierExtension,omitempty"`
	ProductName       string                                          `json:"productName"`
	Type              *CodeableConcept                                `json:"type,omitempty"`
	NamePart          []MedicinalProductDefinitionNameNamePart        `json:"namePart,omitempty"`
	CountryLanguage   []MedicinalProductDefinitionNameCountryLanguage `json:"countryLanguage,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinitionNameNamePart struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Part              string          `json:"part"`
	Type              CodeableConcept `json:"type"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinitionNameCountryLanguage struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Country           CodeableConcept  `json:"country"`
	Jurisdiction      *CodeableConcept `json:"jurisdiction,omitempty"`
	Language          CodeableConcept  `json:"language"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinitionCrossReference struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Product           CodeableReference `json:"product"`
	Type              *CodeableConcept  `json:"type,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinitionOperation struct {
	Id                       *string            `json:"id,omitempty"`
	Extension                []Extension        `json:"extension,omitempty"`
	ModifierExtension        []Extension        `json:"modifierExtension,omitempty"`
	Type                     *CodeableReference `json:"type,omitempty"`
	EffectiveDate            *Period            `json:"effectiveDate,omitempty"`
	Organization             []Reference        `json:"organization,omitempty"`
	ConfidentialityIndicator *CodeableConcept   `json:"confidentialityIndicator,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinitionCharacteristic struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept  `json:"type"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	ValueQuantity        *Quantity        `json:"valueQuantity"`
	ValueDate            *string          `json:"valueDate"`
	ValueBoolean         *bool            `json:"valueBoolean"`
	ValueAttachment      *Attachment      `json:"valueAttachment"`
}

type OtherMedicinalProductDefinition MedicinalProductDefinition

// on convert struct to json, automatically add resourceType=MedicinalProductDefinition
func (r MedicinalProductDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductDefinition: OtherMedicinalProductDefinition(r),
		ResourceType:                    "MedicinalProductDefinition",
	})
}

func (resource *MedicinalProductDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_Domain(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("domain", nil, optionsValueSet)
	}
	return CodeableConceptSelect("domain", resource.Domain, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_Status(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", resource.Status, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_CombinedPharmaceuticalDoseForm(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("combinedPharmaceuticalDoseForm", nil, optionsValueSet)
	}
	return CodeableConceptSelect("combinedPharmaceuticalDoseForm", resource.CombinedPharmaceuticalDoseForm, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_Route(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("route", nil, optionsValueSet)
	}
	return CodeableConceptSelect("route", &resource.Route[0], optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_LegalStatusOfSupply(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("legalStatusOfSupply", nil, optionsValueSet)
	}
	return CodeableConceptSelect("legalStatusOfSupply", resource.LegalStatusOfSupply, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_AdditionalMonitoringIndicator(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("additionalMonitoringIndicator", nil, optionsValueSet)
	}
	return CodeableConceptSelect("additionalMonitoringIndicator", resource.AdditionalMonitoringIndicator, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_SpecialMeasures(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("specialMeasures", nil, optionsValueSet)
	}
	return CodeableConceptSelect("specialMeasures", &resource.SpecialMeasures[0], optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_PediatricUseIndicator(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("pediatricUseIndicator", nil, optionsValueSet)
	}
	return CodeableConceptSelect("pediatricUseIndicator", resource.PediatricUseIndicator, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_Classification(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("classification", nil, optionsValueSet)
	}
	return CodeableConceptSelect("classification", &resource.Classification[0], optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_PackagedMedicinalProduct(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("packagedMedicinalProduct", nil, optionsValueSet)
	}
	return CodeableConceptSelect("packagedMedicinalProduct", &resource.PackagedMedicinalProduct[0], optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_Ingredient(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ingredient", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ingredient", &resource.Ingredient[0], optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodingSelect("code", nil, optionsValueSet)
	}
	return CodingSelect("code", &resource.Code[0], optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_ContactType(numContact int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Contact) >= numContact {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Contact[numContact].Type, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_NameType(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name) >= numName {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Name[numName].Type, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_NameNamePartType(numName int, numNamePart int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name[numName].NamePart) >= numNamePart {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Name[numName].NamePart[numNamePart].Type, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_NameCountryLanguageCountry(numName int, numCountryLanguage int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name[numName].CountryLanguage) >= numCountryLanguage {
		return CodeableConceptSelect("country", nil, optionsValueSet)
	}
	return CodeableConceptSelect("country", &resource.Name[numName].CountryLanguage[numCountryLanguage].Country, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_NameCountryLanguageJurisdiction(numName int, numCountryLanguage int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name[numName].CountryLanguage) >= numCountryLanguage {
		return CodeableConceptSelect("jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("jurisdiction", resource.Name[numName].CountryLanguage[numCountryLanguage].Jurisdiction, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_NameCountryLanguageLanguage(numName int, numCountryLanguage int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Name[numName].CountryLanguage) >= numCountryLanguage {
		return CodeableConceptSelect("language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("language", &resource.Name[numName].CountryLanguage[numCountryLanguage].Language, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_CrossReferenceType(numCrossReference int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CrossReference) >= numCrossReference {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.CrossReference[numCrossReference].Type, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_OperationConfidentialityIndicator(numOperation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Operation) >= numOperation {
		return CodeableConceptSelect("confidentialityIndicator", nil, optionsValueSet)
	}
	return CodeableConceptSelect("confidentialityIndicator", resource.Operation[numOperation].ConfidentialityIndicator, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_CharacteristicType(numCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Characteristic) >= numCharacteristic {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Characteristic[numCharacteristic].Type, optionsValueSet)
}
