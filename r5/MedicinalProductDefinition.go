package r5

//generated with command go run ./bultaoreune
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
	StatusDate                     *FhirDateTime                              `json:"statusDate,omitempty"`
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
	ValueDate            *FhirDate        `json:"valueDate,omitempty"`
	ValueBoolean         *bool            `json:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty"`
}

type OtherMedicinalProductDefinition MedicinalProductDefinition

// struct -> json, automatically add resourceType=Patient
func (r MedicinalProductDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductDefinition: OtherMedicinalProductDefinition(r),
		ResourceType:                    "MedicinalProductDefinition",
	})
}

func (r MedicinalProductDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicinalProductDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "MedicinalProductDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r MedicinalProductDefinition) ResourceType() string {
	return "MedicinalProductDefinition"
}

func (resource *MedicinalProductDefinition) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Domain(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("domain", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("domain", resource.Domain, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Status(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_StatusDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("statusDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("statusDate", resource.StatusDate, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CombinedPharmaceuticalDoseForm(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("combinedPharmaceuticalDoseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("combinedPharmaceuticalDoseForm", resource.CombinedPharmaceuticalDoseForm, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Route(numRoute int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRoute >= len(resource.Route) {
		return CodeableConceptSelect("route["+strconv.Itoa(numRoute)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("route["+strconv.Itoa(numRoute)+"]", &resource.Route[numRoute], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Indication(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("indication", nil, htmlAttrs)
	}
	return StringInput("indication", resource.Indication, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_LegalStatusOfSupply(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("legalStatusOfSupply", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("legalStatusOfSupply", resource.LegalStatusOfSupply, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_AdditionalMonitoringIndicator(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("additionalMonitoringIndicator", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("additionalMonitoringIndicator", resource.AdditionalMonitoringIndicator, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_SpecialMeasures(numSpecialMeasures int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecialMeasures >= len(resource.SpecialMeasures) {
		return CodeableConceptSelect("specialMeasures["+strconv.Itoa(numSpecialMeasures)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("specialMeasures["+strconv.Itoa(numSpecialMeasures)+"]", &resource.SpecialMeasures[numSpecialMeasures], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_PediatricUseIndicator(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("pediatricUseIndicator", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("pediatricUseIndicator", resource.PediatricUseIndicator, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Classification(numClassification int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClassification >= len(resource.Classification) {
		return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"]", &resource.Classification[numClassification], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_MarketingStatus(numMarketingStatus int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMarketingStatus >= len(resource.MarketingStatus) {
		return MarketingStatusInput("marketingStatus["+strconv.Itoa(numMarketingStatus)+"]", nil, htmlAttrs)
	}
	return MarketingStatusInput("marketingStatus["+strconv.Itoa(numMarketingStatus)+"]", &resource.MarketingStatus[numMarketingStatus], htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_PackagedMedicinalProduct(numPackagedMedicinalProduct int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPackagedMedicinalProduct >= len(resource.PackagedMedicinalProduct) {
		return CodeableConceptSelect("packagedMedicinalProduct["+strconv.Itoa(numPackagedMedicinalProduct)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("packagedMedicinalProduct["+strconv.Itoa(numPackagedMedicinalProduct)+"]", &resource.PackagedMedicinalProduct[numPackagedMedicinalProduct], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_ComprisedOf(frs []FhirResource, numComprisedOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComprisedOf >= len(resource.ComprisedOf) {
		return ReferenceInput(frs, "comprisedOf["+strconv.Itoa(numComprisedOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "comprisedOf["+strconv.Itoa(numComprisedOf)+"]", &resource.ComprisedOf[numComprisedOf], htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Ingredient(numIngredient int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableConceptSelect("ingredient["+strconv.Itoa(numIngredient)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ingredient["+strconv.Itoa(numIngredient)+"]", &resource.Ingredient[numIngredient], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Impurity(numImpurity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numImpurity >= len(resource.Impurity) {
		return CodeableReferenceInput("impurity["+strconv.Itoa(numImpurity)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("impurity["+strconv.Itoa(numImpurity)+"]", &resource.Impurity[numImpurity], htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_AttachedDocument(frs []FhirResource, numAttachedDocument int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAttachedDocument >= len(resource.AttachedDocument) {
		return ReferenceInput(frs, "attachedDocument["+strconv.Itoa(numAttachedDocument)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "attachedDocument["+strconv.Itoa(numAttachedDocument)+"]", &resource.AttachedDocument[numAttachedDocument], htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_MasterFile(frs []FhirResource, numMasterFile int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMasterFile >= len(resource.MasterFile) {
		return ReferenceInput(frs, "masterFile["+strconv.Itoa(numMasterFile)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "masterFile["+strconv.Itoa(numMasterFile)+"]", &resource.MasterFile[numMasterFile], htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_ClinicalTrial(frs []FhirResource, numClinicalTrial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClinicalTrial >= len(resource.ClinicalTrial) {
		return ReferenceInput(frs, "clinicalTrial["+strconv.Itoa(numClinicalTrial)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "clinicalTrial["+strconv.Itoa(numClinicalTrial)+"]", &resource.ClinicalTrial[numClinicalTrial], htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Code(numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return CodingSelect("code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("code["+strconv.Itoa(numCode)+"]", &resource.Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_ContactType(numContact int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return CodeableConceptSelect("contact["+strconv.Itoa(numContact)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("contact["+strconv.Itoa(numContact)+"].type", resource.Contact[numContact].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_ContactContact(frs []FhirResource, numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ReferenceInput(frs, "contact["+strconv.Itoa(numContact)+"].contact", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "contact["+strconv.Itoa(numContact)+"].contact", &resource.Contact[numContact].Contact, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_NameProductName(numName int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return StringInput("name["+strconv.Itoa(numName)+"].productName", nil, htmlAttrs)
	}
	return StringInput("name["+strconv.Itoa(numName)+"].productName", &resource.Name[numName].ProductName, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_NameType(numName int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].type", resource.Name[numName].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_NamePartPart(numName int, numPart int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numPart >= len(resource.Name[numName].Part) {
		return StringInput("name["+strconv.Itoa(numName)+"].part["+strconv.Itoa(numPart)+"].part", nil, htmlAttrs)
	}
	return StringInput("name["+strconv.Itoa(numName)+"].part["+strconv.Itoa(numPart)+"].part", &resource.Name[numName].Part[numPart].Part, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_NamePartType(numName int, numPart int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numPart >= len(resource.Name[numName].Part) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].part["+strconv.Itoa(numPart)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].part["+strconv.Itoa(numPart)+"].type", &resource.Name[numName].Part[numPart].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_NameUsageCountry(numName int, numUsage int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numUsage >= len(resource.Name[numName].Usage) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].usage["+strconv.Itoa(numUsage)+"].country", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].usage["+strconv.Itoa(numUsage)+"].country", &resource.Name[numName].Usage[numUsage].Country, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_NameUsageJurisdiction(numName int, numUsage int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) || numUsage >= len(resource.Name[numName].Usage) {
		return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].usage["+strconv.Itoa(numUsage)+"].jurisdiction", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("name["+strconv.Itoa(numName)+"].usage["+strconv.Itoa(numUsage)+"].jurisdiction", resource.Name[numName].Usage[numUsage].Jurisdiction, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CrossReferenceProduct(numCrossReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCrossReference >= len(resource.CrossReference) {
		return CodeableReferenceInput("crossReference["+strconv.Itoa(numCrossReference)+"].product", nil, htmlAttrs)
	}
	return CodeableReferenceInput("crossReference["+strconv.Itoa(numCrossReference)+"].product", &resource.CrossReference[numCrossReference].Product, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CrossReferenceType(numCrossReference int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCrossReference >= len(resource.CrossReference) {
		return CodeableConceptSelect("crossReference["+strconv.Itoa(numCrossReference)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("crossReference["+strconv.Itoa(numCrossReference)+"].type", resource.CrossReference[numCrossReference].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_OperationType(numOperation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOperation >= len(resource.Operation) {
		return CodeableReferenceInput("operation["+strconv.Itoa(numOperation)+"].type", nil, htmlAttrs)
	}
	return CodeableReferenceInput("operation["+strconv.Itoa(numOperation)+"].type", resource.Operation[numOperation].Type, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_OperationEffectiveDate(numOperation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOperation >= len(resource.Operation) {
		return PeriodInput("operation["+strconv.Itoa(numOperation)+"].effectiveDate", nil, htmlAttrs)
	}
	return PeriodInput("operation["+strconv.Itoa(numOperation)+"].effectiveDate", resource.Operation[numOperation].EffectiveDate, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_OperationOrganization(frs []FhirResource, numOperation int, numOrganization int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOperation >= len(resource.Operation) || numOrganization >= len(resource.Operation[numOperation].Organization) {
		return ReferenceInput(frs, "operation["+strconv.Itoa(numOperation)+"].organization["+strconv.Itoa(numOrganization)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "operation["+strconv.Itoa(numOperation)+"].organization["+strconv.Itoa(numOrganization)+"]", &resource.Operation[numOperation].Organization[numOrganization], htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_OperationConfidentialityIndicator(numOperation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOperation >= len(resource.Operation) {
		return CodeableConceptSelect("operation["+strconv.Itoa(numOperation)+"].confidentialityIndicator", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("operation["+strconv.Itoa(numOperation)+"].confidentialityIndicator", resource.Operation[numOperation].ConfidentialityIndicator, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CharacteristicType(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].type", &resource.Characteristic[numCharacteristic].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CharacteristicValueCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].valueCodeableConcept", resource.Characteristic[numCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CharacteristicValueMarkdown(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueMarkdown", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueMarkdown", resource.Characteristic[numCharacteristic].ValueMarkdown, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CharacteristicValueQuantity(numCharacteristic int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return QuantityInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueQuantity", resource.Characteristic[numCharacteristic].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CharacteristicValueInteger(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return IntInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueInteger", resource.Characteristic[numCharacteristic].ValueInteger, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CharacteristicValueDate(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return FhirDateInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueDate", nil, htmlAttrs)
	}
	return FhirDateInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueDate", resource.Characteristic[numCharacteristic].ValueDate, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CharacteristicValueBoolean(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueBoolean", resource.Characteristic[numCharacteristic].ValueBoolean, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CharacteristicValueAttachment(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return AttachmentInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueAttachment", resource.Characteristic[numCharacteristic].ValueAttachment, htmlAttrs)
}
