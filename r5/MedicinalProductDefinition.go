package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	StatusDate                     *time.Time                                 `json:"statusDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	ValueDate            *time.Time       `json:"valueDate,omitempty,format:'2006-01-02'"`
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
func (resource *MedicinalProductDefinition) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductDefinition.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Domain(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductDefinition.Domain", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Domain", resource.Domain, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductDefinition.Version", nil, htmlAttrs)
	}
	return StringInput("MedicinalProductDefinition.Version", resource.Version, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Status(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_StatusDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("MedicinalProductDefinition.StatusDate", nil, htmlAttrs)
	}
	return DateTimeInput("MedicinalProductDefinition.StatusDate", resource.StatusDate, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("MedicinalProductDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CombinedPharmaceuticalDoseForm(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductDefinition.CombinedPharmaceuticalDoseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.CombinedPharmaceuticalDoseForm", resource.CombinedPharmaceuticalDoseForm, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Route(numRoute int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numRoute >= len(resource.Route) {
		return CodeableConceptSelect("MedicinalProductDefinition.Route."+strconv.Itoa(numRoute)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Route."+strconv.Itoa(numRoute)+".", &resource.Route[numRoute], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Indication(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductDefinition.Indication", nil, htmlAttrs)
	}
	return StringInput("MedicinalProductDefinition.Indication", resource.Indication, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_LegalStatusOfSupply(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductDefinition.LegalStatusOfSupply", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.LegalStatusOfSupply", resource.LegalStatusOfSupply, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_AdditionalMonitoringIndicator(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductDefinition.AdditionalMonitoringIndicator", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.AdditionalMonitoringIndicator", resource.AdditionalMonitoringIndicator, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_SpecialMeasures(numSpecialMeasures int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSpecialMeasures >= len(resource.SpecialMeasures) {
		return CodeableConceptSelect("MedicinalProductDefinition.SpecialMeasures."+strconv.Itoa(numSpecialMeasures)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.SpecialMeasures."+strconv.Itoa(numSpecialMeasures)+".", &resource.SpecialMeasures[numSpecialMeasures], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_PediatricUseIndicator(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductDefinition.PediatricUseIndicator", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.PediatricUseIndicator", resource.PediatricUseIndicator, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Classification(numClassification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numClassification >= len(resource.Classification) {
		return CodeableConceptSelect("MedicinalProductDefinition.Classification."+strconv.Itoa(numClassification)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Classification."+strconv.Itoa(numClassification)+".", &resource.Classification[numClassification], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_PackagedMedicinalProduct(numPackagedMedicinalProduct int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numPackagedMedicinalProduct >= len(resource.PackagedMedicinalProduct) {
		return CodeableConceptSelect("MedicinalProductDefinition.PackagedMedicinalProduct."+strconv.Itoa(numPackagedMedicinalProduct)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.PackagedMedicinalProduct."+strconv.Itoa(numPackagedMedicinalProduct)+".", &resource.PackagedMedicinalProduct[numPackagedMedicinalProduct], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Ingredient(numIngredient int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableConceptSelect("MedicinalProductDefinition.Ingredient."+strconv.Itoa(numIngredient)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Ingredient."+strconv.Itoa(numIngredient)+".", &resource.Ingredient[numIngredient], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_Code(numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCode >= len(resource.Code) {
		return CodingSelect("MedicinalProductDefinition.Code."+strconv.Itoa(numCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("MedicinalProductDefinition.Code."+strconv.Itoa(numCode)+".", &resource.Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_ContactType(numContact int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numContact >= len(resource.Contact) {
		return CodeableConceptSelect("MedicinalProductDefinition.Contact."+strconv.Itoa(numContact)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Contact."+strconv.Itoa(numContact)+"..Type", resource.Contact[numContact].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_NameProductName(numName int, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) {
		return StringInput("MedicinalProductDefinition.Name."+strconv.Itoa(numName)+"..ProductName", nil, htmlAttrs)
	}
	return StringInput("MedicinalProductDefinition.Name."+strconv.Itoa(numName)+"..ProductName", &resource.Name[numName].ProductName, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_NameType(numName int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) {
		return CodeableConceptSelect("MedicinalProductDefinition.Name."+strconv.Itoa(numName)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Name."+strconv.Itoa(numName)+"..Type", resource.Name[numName].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_NamePartPart(numName int, numPart int, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) || numPart >= len(resource.Name[numName].Part) {
		return StringInput("MedicinalProductDefinition.Name."+strconv.Itoa(numName)+"..Part."+strconv.Itoa(numPart)+"..Part", nil, htmlAttrs)
	}
	return StringInput("MedicinalProductDefinition.Name."+strconv.Itoa(numName)+"..Part."+strconv.Itoa(numPart)+"..Part", &resource.Name[numName].Part[numPart].Part, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_NamePartType(numName int, numPart int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) || numPart >= len(resource.Name[numName].Part) {
		return CodeableConceptSelect("MedicinalProductDefinition.Name."+strconv.Itoa(numName)+"..Part."+strconv.Itoa(numPart)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Name."+strconv.Itoa(numName)+"..Part."+strconv.Itoa(numPart)+"..Type", &resource.Name[numName].Part[numPart].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_NameUsageCountry(numName int, numUsage int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) || numUsage >= len(resource.Name[numName].Usage) {
		return CodeableConceptSelect("MedicinalProductDefinition.Name."+strconv.Itoa(numName)+"..Usage."+strconv.Itoa(numUsage)+"..Country", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Name."+strconv.Itoa(numName)+"..Usage."+strconv.Itoa(numUsage)+"..Country", &resource.Name[numName].Usage[numUsage].Country, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_NameUsageJurisdiction(numName int, numUsage int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) || numUsage >= len(resource.Name[numName].Usage) {
		return CodeableConceptSelect("MedicinalProductDefinition.Name."+strconv.Itoa(numName)+"..Usage."+strconv.Itoa(numUsage)+"..Jurisdiction", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Name."+strconv.Itoa(numName)+"..Usage."+strconv.Itoa(numUsage)+"..Jurisdiction", resource.Name[numName].Usage[numUsage].Jurisdiction, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CrossReferenceType(numCrossReference int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCrossReference >= len(resource.CrossReference) {
		return CodeableConceptSelect("MedicinalProductDefinition.CrossReference."+strconv.Itoa(numCrossReference)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.CrossReference."+strconv.Itoa(numCrossReference)+"..Type", resource.CrossReference[numCrossReference].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_OperationConfidentialityIndicator(numOperation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numOperation >= len(resource.Operation) {
		return CodeableConceptSelect("MedicinalProductDefinition.Operation."+strconv.Itoa(numOperation)+"..ConfidentialityIndicator", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Operation."+strconv.Itoa(numOperation)+"..ConfidentialityIndicator", resource.Operation[numOperation].ConfidentialityIndicator, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CharacteristicType(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("MedicinalProductDefinition.Characteristic."+strconv.Itoa(numCharacteristic)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Characteristic."+strconv.Itoa(numCharacteristic)+"..Type", &resource.Characteristic[numCharacteristic].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CharacteristicValueCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("MedicinalProductDefinition.Characteristic."+strconv.Itoa(numCharacteristic)+"..ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicinalProductDefinition.Characteristic."+strconv.Itoa(numCharacteristic)+"..ValueCodeableConcept", resource.Characteristic[numCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CharacteristicValueMarkdown(numCharacteristic int, htmlAttrs string) templ.Component {

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("MedicinalProductDefinition.Characteristic."+strconv.Itoa(numCharacteristic)+"..ValueMarkdown", nil, htmlAttrs)
	}
	return StringInput("MedicinalProductDefinition.Characteristic."+strconv.Itoa(numCharacteristic)+"..ValueMarkdown", resource.Characteristic[numCharacteristic].ValueMarkdown, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CharacteristicValueInteger(numCharacteristic int, htmlAttrs string) templ.Component {

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return IntInput("MedicinalProductDefinition.Characteristic."+strconv.Itoa(numCharacteristic)+"..ValueInteger", nil, htmlAttrs)
	}
	return IntInput("MedicinalProductDefinition.Characteristic."+strconv.Itoa(numCharacteristic)+"..ValueInteger", resource.Characteristic[numCharacteristic].ValueInteger, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CharacteristicValueDate(numCharacteristic int, htmlAttrs string) templ.Component {

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return DateInput("MedicinalProductDefinition.Characteristic."+strconv.Itoa(numCharacteristic)+"..ValueDate", nil, htmlAttrs)
	}
	return DateInput("MedicinalProductDefinition.Characteristic."+strconv.Itoa(numCharacteristic)+"..ValueDate", resource.Characteristic[numCharacteristic].ValueDate, htmlAttrs)
}
func (resource *MedicinalProductDefinition) T_CharacteristicValueBoolean(numCharacteristic int, htmlAttrs string) templ.Component {

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("MedicinalProductDefinition.Characteristic."+strconv.Itoa(numCharacteristic)+"..ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("MedicinalProductDefinition.Characteristic."+strconv.Itoa(numCharacteristic)+"..ValueBoolean", resource.Characteristic[numCharacteristic].ValueBoolean, htmlAttrs)
}
