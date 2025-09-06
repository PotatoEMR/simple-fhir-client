package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/MedicinalProductDefinition
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
	ComprisedOf                    []Reference                                `json:"comprisedOf,omitempty"`
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

// http://hl7.org/fhir/r5/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinitionContact struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Contact           Reference        `json:"contact"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinitionName struct {
	Id                *string                               `json:"id,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	ProductName       string                                `json:"productName"`
	Type              *CodeableConcept                      `json:"type,omitempty"`
	Part              []MedicinalProductDefinitionNamePart  `json:"part,omitempty"`
	Usage             []MedicinalProductDefinitionNameUsage `json:"usage,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinitionNamePart struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Part              string          `json:"part"`
	Type              CodeableConcept `json:"type"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinitionNameUsage struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Country           CodeableConcept  `json:"country"`
	Jurisdiction      *CodeableConcept `json:"jurisdiction,omitempty"`
	Language          CodeableConcept  `json:"language"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinitionCrossReference struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Product           CodeableReference `json:"product"`
	Type              *CodeableConcept  `json:"type,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinitionOperation struct {
	Id                       *string            `json:"id,omitempty"`
	Extension                []Extension        `json:"extension,omitempty"`
	ModifierExtension        []Extension        `json:"modifierExtension,omitempty"`
	Type                     *CodeableReference `json:"type,omitempty"`
	EffectiveDate            *Period            `json:"effectiveDate,omitempty"`
	Organization             []Reference        `json:"organization,omitempty"`
	ConfidentialityIndicator *CodeableConcept   `json:"confidentialityIndicator,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinitionCharacteristic struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept  `json:"type"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueMarkdown        *string          `json:"valueMarkdown,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueInteger         *int             `json:"valueInteger,omitempty"`
	ValueDate            *string          `json:"valueDate,omitempty"`
	ValueBoolean         *bool            `json:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty"`
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

func (resource *MedicinalProductDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductDefinition.Id", nil)
	}
	return StringInput("MedicinalProductDefinition.Id", resource.Id)
}
func (resource *MedicinalProductDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductDefinition.ImplicitRules", nil)
	}
	return StringInput("MedicinalProductDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *MedicinalProductDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MedicinalProductDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("MedicinalProductDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductDefinition.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Type", resource.Type, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_Domain(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductDefinition.Domain", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Domain", resource.Domain, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductDefinition.Version", nil)
	}
	return StringInput("MedicinalProductDefinition.Version", resource.Version)
}
func (resource *MedicinalProductDefinition) T_Status(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductDefinition.Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Status", resource.Status, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_StatusDate() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductDefinition.StatusDate", nil)
	}
	return StringInput("MedicinalProductDefinition.StatusDate", resource.StatusDate)
}
func (resource *MedicinalProductDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductDefinition.Description", nil)
	}
	return StringInput("MedicinalProductDefinition.Description", resource.Description)
}
func (resource *MedicinalProductDefinition) T_CombinedPharmaceuticalDoseForm(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductDefinition.CombinedPharmaceuticalDoseForm", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.CombinedPharmaceuticalDoseForm", resource.CombinedPharmaceuticalDoseForm, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_Route(numRoute int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Route) >= numRoute {
		return CodeableConceptSelect("MedicinalProductDefinition.Route["+strconv.Itoa(numRoute)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Route["+strconv.Itoa(numRoute)+"]", &resource.Route[numRoute], optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_Indication() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductDefinition.Indication", nil)
	}
	return StringInput("MedicinalProductDefinition.Indication", resource.Indication)
}
func (resource *MedicinalProductDefinition) T_LegalStatusOfSupply(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductDefinition.LegalStatusOfSupply", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.LegalStatusOfSupply", resource.LegalStatusOfSupply, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_AdditionalMonitoringIndicator(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductDefinition.AdditionalMonitoringIndicator", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.AdditionalMonitoringIndicator", resource.AdditionalMonitoringIndicator, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_SpecialMeasures(numSpecialMeasures int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SpecialMeasures) >= numSpecialMeasures {
		return CodeableConceptSelect("MedicinalProductDefinition.SpecialMeasures["+strconv.Itoa(numSpecialMeasures)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.SpecialMeasures["+strconv.Itoa(numSpecialMeasures)+"]", &resource.SpecialMeasures[numSpecialMeasures], optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_PediatricUseIndicator(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductDefinition.PediatricUseIndicator", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.PediatricUseIndicator", resource.PediatricUseIndicator, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_Classification(numClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Classification) >= numClassification {
		return CodeableConceptSelect("MedicinalProductDefinition.Classification["+strconv.Itoa(numClassification)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Classification["+strconv.Itoa(numClassification)+"]", &resource.Classification[numClassification], optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_PackagedMedicinalProduct(numPackagedMedicinalProduct int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PackagedMedicinalProduct) >= numPackagedMedicinalProduct {
		return CodeableConceptSelect("MedicinalProductDefinition.PackagedMedicinalProduct["+strconv.Itoa(numPackagedMedicinalProduct)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.PackagedMedicinalProduct["+strconv.Itoa(numPackagedMedicinalProduct)+"]", &resource.PackagedMedicinalProduct[numPackagedMedicinalProduct], optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_Ingredient(numIngredient int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Ingredient) >= numIngredient {
		return CodeableConceptSelect("MedicinalProductDefinition.Ingredient["+strconv.Itoa(numIngredient)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Ingredient["+strconv.Itoa(numIngredient)+"]", &resource.Ingredient[numIngredient], optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_Code(numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Code) >= numCode {
		return CodingSelect("MedicinalProductDefinition.Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet)
	}
	return CodingSelect("MedicinalProductDefinition.Code["+strconv.Itoa(numCode)+"]", &resource.Code[numCode], optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_ContactId(numContact int) templ.Component {

	if resource == nil || len(resource.Contact) >= numContact {
		return StringInput("MedicinalProductDefinition.Contact["+strconv.Itoa(numContact)+"].Id", nil)
	}
	return StringInput("MedicinalProductDefinition.Contact["+strconv.Itoa(numContact)+"].Id", resource.Contact[numContact].Id)
}
func (resource *MedicinalProductDefinition) T_ContactType(numContact int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Contact) >= numContact {
		return CodeableConceptSelect("MedicinalProductDefinition.Contact["+strconv.Itoa(numContact)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Contact["+strconv.Itoa(numContact)+"].Type", resource.Contact[numContact].Type, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_NameId(numName int) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return StringInput("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Id", nil)
	}
	return StringInput("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Id", resource.Name[numName].Id)
}
func (resource *MedicinalProductDefinition) T_NameProductName(numName int) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return StringInput("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].ProductName", nil)
	}
	return StringInput("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].ProductName", &resource.Name[numName].ProductName)
}
func (resource *MedicinalProductDefinition) T_NameType(numName int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return CodeableConceptSelect("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Type", resource.Name[numName].Type, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_NamePartId(numName int, numPart int) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Part) >= numPart {
		return StringInput("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Part["+strconv.Itoa(numPart)+"].Id", nil)
	}
	return StringInput("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Part["+strconv.Itoa(numPart)+"].Id", resource.Name[numName].Part[numPart].Id)
}
func (resource *MedicinalProductDefinition) T_NamePartPart(numName int, numPart int) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Part) >= numPart {
		return StringInput("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Part["+strconv.Itoa(numPart)+"].Part", nil)
	}
	return StringInput("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Part["+strconv.Itoa(numPart)+"].Part", &resource.Name[numName].Part[numPart].Part)
}
func (resource *MedicinalProductDefinition) T_NamePartType(numName int, numPart int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Part) >= numPart {
		return CodeableConceptSelect("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Part["+strconv.Itoa(numPart)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Part["+strconv.Itoa(numPart)+"].Type", &resource.Name[numName].Part[numPart].Type, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_NameUsageId(numName int, numUsage int) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Usage) >= numUsage {
		return StringInput("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Usage["+strconv.Itoa(numUsage)+"].Id", nil)
	}
	return StringInput("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Usage["+strconv.Itoa(numUsage)+"].Id", resource.Name[numName].Usage[numUsage].Id)
}
func (resource *MedicinalProductDefinition) T_NameUsageCountry(numName int, numUsage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Usage) >= numUsage {
		return CodeableConceptSelect("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Usage["+strconv.Itoa(numUsage)+"].Country", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Usage["+strconv.Itoa(numUsage)+"].Country", &resource.Name[numName].Usage[numUsage].Country, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_NameUsageJurisdiction(numName int, numUsage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Usage) >= numUsage {
		return CodeableConceptSelect("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Usage["+strconv.Itoa(numUsage)+"].Jurisdiction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Usage["+strconv.Itoa(numUsage)+"].Jurisdiction", resource.Name[numName].Usage[numUsage].Jurisdiction, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_NameUsageLanguage(numName int, numUsage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Name) >= numName || len(resource.Name[numName].Usage) >= numUsage {
		return CodeableConceptSelect("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Usage["+strconv.Itoa(numUsage)+"].Language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Name["+strconv.Itoa(numName)+"].Usage["+strconv.Itoa(numUsage)+"].Language", &resource.Name[numName].Usage[numUsage].Language, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_CrossReferenceId(numCrossReference int) templ.Component {

	if resource == nil || len(resource.CrossReference) >= numCrossReference {
		return StringInput("MedicinalProductDefinition.CrossReference["+strconv.Itoa(numCrossReference)+"].Id", nil)
	}
	return StringInput("MedicinalProductDefinition.CrossReference["+strconv.Itoa(numCrossReference)+"].Id", resource.CrossReference[numCrossReference].Id)
}
func (resource *MedicinalProductDefinition) T_CrossReferenceType(numCrossReference int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CrossReference) >= numCrossReference {
		return CodeableConceptSelect("MedicinalProductDefinition.CrossReference["+strconv.Itoa(numCrossReference)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.CrossReference["+strconv.Itoa(numCrossReference)+"].Type", resource.CrossReference[numCrossReference].Type, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_OperationId(numOperation int) templ.Component {

	if resource == nil || len(resource.Operation) >= numOperation {
		return StringInput("MedicinalProductDefinition.Operation["+strconv.Itoa(numOperation)+"].Id", nil)
	}
	return StringInput("MedicinalProductDefinition.Operation["+strconv.Itoa(numOperation)+"].Id", resource.Operation[numOperation].Id)
}
func (resource *MedicinalProductDefinition) T_OperationConfidentialityIndicator(numOperation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Operation) >= numOperation {
		return CodeableConceptSelect("MedicinalProductDefinition.Operation["+strconv.Itoa(numOperation)+"].ConfidentialityIndicator", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Operation["+strconv.Itoa(numOperation)+"].ConfidentialityIndicator", resource.Operation[numOperation].ConfidentialityIndicator, optionsValueSet)
}
func (resource *MedicinalProductDefinition) T_CharacteristicId(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return StringInput("MedicinalProductDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].Id", nil)
	}
	return StringInput("MedicinalProductDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].Id", resource.Characteristic[numCharacteristic].Id)
}
func (resource *MedicinalProductDefinition) T_CharacteristicType(numCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return CodeableConceptSelect("MedicinalProductDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Characteristic["+strconv.Itoa(numCharacteristic)+"].Type", &resource.Characteristic[numCharacteristic].Type, optionsValueSet)
}
