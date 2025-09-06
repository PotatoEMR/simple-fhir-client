package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledge struct {
	Id                         *string                                         `json:"id,omitempty"`
	Meta                       *Meta                                           `json:"meta,omitempty"`
	ImplicitRules              *string                                         `json:"implicitRules,omitempty"`
	Language                   *string                                         `json:"language,omitempty"`
	Text                       *Narrative                                      `json:"text,omitempty"`
	Contained                  []Resource                                      `json:"contained,omitempty"`
	Extension                  []Extension                                     `json:"extension,omitempty"`
	ModifierExtension          []Extension                                     `json:"modifierExtension,omitempty"`
	Code                       *CodeableConcept                                `json:"code,omitempty"`
	Status                     *string                                         `json:"status,omitempty"`
	Manufacturer               *Reference                                      `json:"manufacturer,omitempty"`
	DoseForm                   *CodeableConcept                                `json:"doseForm,omitempty"`
	Amount                     *Quantity                                       `json:"amount,omitempty"`
	Synonym                    []string                                        `json:"synonym,omitempty"`
	RelatedMedicationKnowledge []MedicationKnowledgeRelatedMedicationKnowledge `json:"relatedMedicationKnowledge,omitempty"`
	AssociatedMedication       []Reference                                     `json:"associatedMedication,omitempty"`
	ProductType                []CodeableConcept                               `json:"productType,omitempty"`
	Monograph                  []MedicationKnowledgeMonograph                  `json:"monograph,omitempty"`
	Ingredient                 []MedicationKnowledgeIngredient                 `json:"ingredient,omitempty"`
	PreparationInstruction     *string                                         `json:"preparationInstruction,omitempty"`
	IntendedRoute              []CodeableConcept                               `json:"intendedRoute,omitempty"`
	Cost                       []MedicationKnowledgeCost                       `json:"cost,omitempty"`
	MonitoringProgram          []MedicationKnowledgeMonitoringProgram          `json:"monitoringProgram,omitempty"`
	AdministrationGuidelines   []MedicationKnowledgeAdministrationGuidelines   `json:"administrationGuidelines,omitempty"`
	MedicineClassification     []MedicationKnowledgeMedicineClassification     `json:"medicineClassification,omitempty"`
	Packaging                  *MedicationKnowledgePackaging                   `json:"packaging,omitempty"`
	DrugCharacteristic         []MedicationKnowledgeDrugCharacteristic         `json:"drugCharacteristic,omitempty"`
	Contraindication           []Reference                                     `json:"contraindication,omitempty"`
	Regulatory                 []MedicationKnowledgeRegulatory                 `json:"regulatory,omitempty"`
	Kinetics                   []MedicationKnowledgeKinetics                   `json:"kinetics,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeRelatedMedicationKnowledge struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Reference         []Reference     `json:"reference"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeMonograph struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Source            *Reference       `json:"source,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeIngredient struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	ItemCodeableConcept CodeableConcept `json:"itemCodeableConcept"`
	ItemReference       Reference       `json:"itemReference"`
	IsActive            *bool           `json:"isActive,omitempty"`
	Strength            *Ratio          `json:"strength,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeCost struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Source            *string         `json:"source,omitempty"`
	Cost              Money           `json:"cost"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeMonitoringProgram struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Name              *string          `json:"name,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeAdministrationGuidelines struct {
	Id                        *string                                                             `json:"id,omitempty"`
	Extension                 []Extension                                                         `json:"extension,omitempty"`
	ModifierExtension         []Extension                                                         `json:"modifierExtension,omitempty"`
	Dosage                    []MedicationKnowledgeAdministrationGuidelinesDosage                 `json:"dosage,omitempty"`
	IndicationCodeableConcept *CodeableConcept                                                    `json:"indicationCodeableConcept,omitempty"`
	IndicationReference       *Reference                                                          `json:"indicationReference,omitempty"`
	PatientCharacteristics    []MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics `json:"patientCharacteristics,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeAdministrationGuidelinesDosage struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Dosage            []Dosage        `json:"dosage"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics struct {
	Id                            *string         `json:"id,omitempty"`
	Extension                     []Extension     `json:"extension,omitempty"`
	ModifierExtension             []Extension     `json:"modifierExtension,omitempty"`
	CharacteristicCodeableConcept CodeableConcept `json:"characteristicCodeableConcept"`
	CharacteristicQuantity        Quantity        `json:"characteristicQuantity"`
	Value                         []string        `json:"value,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeMedicineClassification struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	Classification    []CodeableConcept `json:"classification,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledgePackaging struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Quantity          *Quantity        `json:"quantity,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeDrugCharacteristic struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Type                 *CodeableConcept `json:"type,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string          `json:"valueString,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueBase64Binary    *string          `json:"valueBase64Binary,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeRegulatory struct {
	Id                  *string                                     `json:"id,omitempty"`
	Extension           []Extension                                 `json:"extension,omitempty"`
	ModifierExtension   []Extension                                 `json:"modifierExtension,omitempty"`
	RegulatoryAuthority Reference                                   `json:"regulatoryAuthority"`
	Substitution        []MedicationKnowledgeRegulatorySubstitution `json:"substitution,omitempty"`
	Schedule            []MedicationKnowledgeRegulatorySchedule     `json:"schedule,omitempty"`
	MaxDispense         *MedicationKnowledgeRegulatoryMaxDispense   `json:"maxDispense,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeRegulatorySubstitution struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Allowed           bool            `json:"allowed"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeRegulatorySchedule struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Schedule          CodeableConcept `json:"schedule"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeRegulatoryMaxDispense struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Quantity          Quantity    `json:"quantity"`
	Period            *Duration   `json:"period,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeKinetics struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	AreaUnderCurve    []Quantity  `json:"areaUnderCurve,omitempty"`
	LethalDose50      []Quantity  `json:"lethalDose50,omitempty"`
	HalfLifePeriod    *Duration   `json:"halfLifePeriod,omitempty"`
}

type OtherMedicationKnowledge MedicationKnowledge

// on convert struct to json, automatically add resourceType=MedicationKnowledge
func (r MedicationKnowledge) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationKnowledge
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationKnowledge: OtherMedicationKnowledge(r),
		ResourceType:             "MedicationKnowledge",
	})
}

func (resource *MedicationKnowledge) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MedicationKnowledge.Id", nil)
	}
	return StringInput("MedicationKnowledge.Id", resource.Id)
}
func (resource *MedicationKnowledge) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MedicationKnowledge.ImplicitRules", nil)
	}
	return StringInput("MedicationKnowledge.ImplicitRules", resource.ImplicitRules)
}
func (resource *MedicationKnowledge) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MedicationKnowledge.Language", nil, optionsValueSet)
	}
	return CodeSelect("MedicationKnowledge.Language", resource.Language, optionsValueSet)
}
func (resource *MedicationKnowledge) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationKnowledge.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.Code", resource.Code, optionsValueSet)
}
func (resource *MedicationKnowledge) T_Status() templ.Component {
	optionsValueSet := VSMedicationknowledge_status

	if resource == nil {
		return CodeSelect("MedicationKnowledge.Status", nil, optionsValueSet)
	}
	return CodeSelect("MedicationKnowledge.Status", resource.Status, optionsValueSet)
}
func (resource *MedicationKnowledge) T_DoseForm(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationKnowledge.DoseForm", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.DoseForm", resource.DoseForm, optionsValueSet)
}
func (resource *MedicationKnowledge) T_Synonym(numSynonym int) templ.Component {

	if resource == nil || len(resource.Synonym) >= numSynonym {
		return StringInput("MedicationKnowledge.Synonym["+strconv.Itoa(numSynonym)+"]", nil)
	}
	return StringInput("MedicationKnowledge.Synonym["+strconv.Itoa(numSynonym)+"]", &resource.Synonym[numSynonym])
}
func (resource *MedicationKnowledge) T_ProductType(numProductType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ProductType) >= numProductType {
		return CodeableConceptSelect("MedicationKnowledge.ProductType["+strconv.Itoa(numProductType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.ProductType["+strconv.Itoa(numProductType)+"]", &resource.ProductType[numProductType], optionsValueSet)
}
func (resource *MedicationKnowledge) T_PreparationInstruction() templ.Component {

	if resource == nil {
		return StringInput("MedicationKnowledge.PreparationInstruction", nil)
	}
	return StringInput("MedicationKnowledge.PreparationInstruction", resource.PreparationInstruction)
}
func (resource *MedicationKnowledge) T_IntendedRoute(numIntendedRoute int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.IntendedRoute) >= numIntendedRoute {
		return CodeableConceptSelect("MedicationKnowledge.IntendedRoute["+strconv.Itoa(numIntendedRoute)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.IntendedRoute["+strconv.Itoa(numIntendedRoute)+"]", &resource.IntendedRoute[numIntendedRoute], optionsValueSet)
}
func (resource *MedicationKnowledge) T_RelatedMedicationKnowledgeId(numRelatedMedicationKnowledge int) templ.Component {

	if resource == nil || len(resource.RelatedMedicationKnowledge) >= numRelatedMedicationKnowledge {
		return StringInput("MedicationKnowledge.RelatedMedicationKnowledge["+strconv.Itoa(numRelatedMedicationKnowledge)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.RelatedMedicationKnowledge["+strconv.Itoa(numRelatedMedicationKnowledge)+"].Id", resource.RelatedMedicationKnowledge[numRelatedMedicationKnowledge].Id)
}
func (resource *MedicationKnowledge) T_RelatedMedicationKnowledgeType(numRelatedMedicationKnowledge int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.RelatedMedicationKnowledge) >= numRelatedMedicationKnowledge {
		return CodeableConceptSelect("MedicationKnowledge.RelatedMedicationKnowledge["+strconv.Itoa(numRelatedMedicationKnowledge)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.RelatedMedicationKnowledge["+strconv.Itoa(numRelatedMedicationKnowledge)+"].Type", &resource.RelatedMedicationKnowledge[numRelatedMedicationKnowledge].Type, optionsValueSet)
}
func (resource *MedicationKnowledge) T_MonographId(numMonograph int) templ.Component {

	if resource == nil || len(resource.Monograph) >= numMonograph {
		return StringInput("MedicationKnowledge.Monograph["+strconv.Itoa(numMonograph)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.Monograph["+strconv.Itoa(numMonograph)+"].Id", resource.Monograph[numMonograph].Id)
}
func (resource *MedicationKnowledge) T_MonographType(numMonograph int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Monograph) >= numMonograph {
		return CodeableConceptSelect("MedicationKnowledge.Monograph["+strconv.Itoa(numMonograph)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.Monograph["+strconv.Itoa(numMonograph)+"].Type", resource.Monograph[numMonograph].Type, optionsValueSet)
}
func (resource *MedicationKnowledge) T_IngredientId(numIngredient int) templ.Component {

	if resource == nil || len(resource.Ingredient) >= numIngredient {
		return StringInput("MedicationKnowledge.Ingredient["+strconv.Itoa(numIngredient)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.Ingredient["+strconv.Itoa(numIngredient)+"].Id", resource.Ingredient[numIngredient].Id)
}
func (resource *MedicationKnowledge) T_IngredientIsActive(numIngredient int) templ.Component {

	if resource == nil || len(resource.Ingredient) >= numIngredient {
		return BoolInput("MedicationKnowledge.Ingredient["+strconv.Itoa(numIngredient)+"].IsActive", nil)
	}
	return BoolInput("MedicationKnowledge.Ingredient["+strconv.Itoa(numIngredient)+"].IsActive", resource.Ingredient[numIngredient].IsActive)
}
func (resource *MedicationKnowledge) T_CostId(numCost int) templ.Component {

	if resource == nil || len(resource.Cost) >= numCost {
		return StringInput("MedicationKnowledge.Cost["+strconv.Itoa(numCost)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.Cost["+strconv.Itoa(numCost)+"].Id", resource.Cost[numCost].Id)
}
func (resource *MedicationKnowledge) T_CostType(numCost int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Cost) >= numCost {
		return CodeableConceptSelect("MedicationKnowledge.Cost["+strconv.Itoa(numCost)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.Cost["+strconv.Itoa(numCost)+"].Type", &resource.Cost[numCost].Type, optionsValueSet)
}
func (resource *MedicationKnowledge) T_CostSource(numCost int) templ.Component {

	if resource == nil || len(resource.Cost) >= numCost {
		return StringInput("MedicationKnowledge.Cost["+strconv.Itoa(numCost)+"].Source", nil)
	}
	return StringInput("MedicationKnowledge.Cost["+strconv.Itoa(numCost)+"].Source", resource.Cost[numCost].Source)
}
func (resource *MedicationKnowledge) T_MonitoringProgramId(numMonitoringProgram int) templ.Component {

	if resource == nil || len(resource.MonitoringProgram) >= numMonitoringProgram {
		return StringInput("MedicationKnowledge.MonitoringProgram["+strconv.Itoa(numMonitoringProgram)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.MonitoringProgram["+strconv.Itoa(numMonitoringProgram)+"].Id", resource.MonitoringProgram[numMonitoringProgram].Id)
}
func (resource *MedicationKnowledge) T_MonitoringProgramType(numMonitoringProgram int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.MonitoringProgram) >= numMonitoringProgram {
		return CodeableConceptSelect("MedicationKnowledge.MonitoringProgram["+strconv.Itoa(numMonitoringProgram)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.MonitoringProgram["+strconv.Itoa(numMonitoringProgram)+"].Type", resource.MonitoringProgram[numMonitoringProgram].Type, optionsValueSet)
}
func (resource *MedicationKnowledge) T_MonitoringProgramName(numMonitoringProgram int) templ.Component {

	if resource == nil || len(resource.MonitoringProgram) >= numMonitoringProgram {
		return StringInput("MedicationKnowledge.MonitoringProgram["+strconv.Itoa(numMonitoringProgram)+"].Name", nil)
	}
	return StringInput("MedicationKnowledge.MonitoringProgram["+strconv.Itoa(numMonitoringProgram)+"].Name", resource.MonitoringProgram[numMonitoringProgram].Name)
}
func (resource *MedicationKnowledge) T_AdministrationGuidelinesId(numAdministrationGuidelines int) templ.Component {

	if resource == nil || len(resource.AdministrationGuidelines) >= numAdministrationGuidelines {
		return StringInput("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].Id", resource.AdministrationGuidelines[numAdministrationGuidelines].Id)
}
func (resource *MedicationKnowledge) T_AdministrationGuidelinesDosageId(numAdministrationGuidelines int, numDosage int) templ.Component {

	if resource == nil || len(resource.AdministrationGuidelines) >= numAdministrationGuidelines || len(resource.AdministrationGuidelines[numAdministrationGuidelines].Dosage) >= numDosage {
		return StringInput("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].Dosage["+strconv.Itoa(numDosage)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].Dosage["+strconv.Itoa(numDosage)+"].Id", resource.AdministrationGuidelines[numAdministrationGuidelines].Dosage[numDosage].Id)
}
func (resource *MedicationKnowledge) T_AdministrationGuidelinesDosageType(numAdministrationGuidelines int, numDosage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AdministrationGuidelines) >= numAdministrationGuidelines || len(resource.AdministrationGuidelines[numAdministrationGuidelines].Dosage) >= numDosage {
		return CodeableConceptSelect("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].Dosage["+strconv.Itoa(numDosage)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].Dosage["+strconv.Itoa(numDosage)+"].Type", &resource.AdministrationGuidelines[numAdministrationGuidelines].Dosage[numDosage].Type, optionsValueSet)
}
func (resource *MedicationKnowledge) T_AdministrationGuidelinesPatientCharacteristicsId(numAdministrationGuidelines int, numPatientCharacteristics int) templ.Component {

	if resource == nil || len(resource.AdministrationGuidelines) >= numAdministrationGuidelines || len(resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics) >= numPatientCharacteristics {
		return StringInput("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].PatientCharacteristics["+strconv.Itoa(numPatientCharacteristics)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].PatientCharacteristics["+strconv.Itoa(numPatientCharacteristics)+"].Id", resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics[numPatientCharacteristics].Id)
}
func (resource *MedicationKnowledge) T_AdministrationGuidelinesPatientCharacteristicsValue(numAdministrationGuidelines int, numPatientCharacteristics int, numValue int) templ.Component {

	if resource == nil || len(resource.AdministrationGuidelines) >= numAdministrationGuidelines || len(resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics) >= numPatientCharacteristics || len(resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics[numPatientCharacteristics].Value) >= numValue {
		return StringInput("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].PatientCharacteristics["+strconv.Itoa(numPatientCharacteristics)+"].Value["+strconv.Itoa(numValue)+"]", nil)
	}
	return StringInput("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].PatientCharacteristics["+strconv.Itoa(numPatientCharacteristics)+"].Value["+strconv.Itoa(numValue)+"]", &resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics[numPatientCharacteristics].Value[numValue])
}
func (resource *MedicationKnowledge) T_MedicineClassificationId(numMedicineClassification int) templ.Component {

	if resource == nil || len(resource.MedicineClassification) >= numMedicineClassification {
		return StringInput("MedicationKnowledge.MedicineClassification["+strconv.Itoa(numMedicineClassification)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.MedicineClassification["+strconv.Itoa(numMedicineClassification)+"].Id", resource.MedicineClassification[numMedicineClassification].Id)
}
func (resource *MedicationKnowledge) T_MedicineClassificationType(numMedicineClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.MedicineClassification) >= numMedicineClassification {
		return CodeableConceptSelect("MedicationKnowledge.MedicineClassification["+strconv.Itoa(numMedicineClassification)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.MedicineClassification["+strconv.Itoa(numMedicineClassification)+"].Type", &resource.MedicineClassification[numMedicineClassification].Type, optionsValueSet)
}
func (resource *MedicationKnowledge) T_MedicineClassificationClassification(numMedicineClassification int, numClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.MedicineClassification) >= numMedicineClassification || len(resource.MedicineClassification[numMedicineClassification].Classification) >= numClassification {
		return CodeableConceptSelect("MedicationKnowledge.MedicineClassification["+strconv.Itoa(numMedicineClassification)+"].Classification["+strconv.Itoa(numClassification)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.MedicineClassification["+strconv.Itoa(numMedicineClassification)+"].Classification["+strconv.Itoa(numClassification)+"]", &resource.MedicineClassification[numMedicineClassification].Classification[numClassification], optionsValueSet)
}
func (resource *MedicationKnowledge) T_PackagingId() templ.Component {

	if resource == nil {
		return StringInput("MedicationKnowledge.Packaging.Id", nil)
	}
	return StringInput("MedicationKnowledge.Packaging.Id", resource.Packaging.Id)
}
func (resource *MedicationKnowledge) T_PackagingType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationKnowledge.Packaging.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.Packaging.Type", resource.Packaging.Type, optionsValueSet)
}
func (resource *MedicationKnowledge) T_DrugCharacteristicId(numDrugCharacteristic int) templ.Component {

	if resource == nil || len(resource.DrugCharacteristic) >= numDrugCharacteristic {
		return StringInput("MedicationKnowledge.DrugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.DrugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].Id", resource.DrugCharacteristic[numDrugCharacteristic].Id)
}
func (resource *MedicationKnowledge) T_DrugCharacteristicType(numDrugCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.DrugCharacteristic) >= numDrugCharacteristic {
		return CodeableConceptSelect("MedicationKnowledge.DrugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.DrugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].Type", resource.DrugCharacteristic[numDrugCharacteristic].Type, optionsValueSet)
}
func (resource *MedicationKnowledge) T_RegulatoryId(numRegulatory int) templ.Component {

	if resource == nil || len(resource.Regulatory) >= numRegulatory {
		return StringInput("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Id", resource.Regulatory[numRegulatory].Id)
}
func (resource *MedicationKnowledge) T_RegulatorySubstitutionId(numRegulatory int, numSubstitution int) templ.Component {

	if resource == nil || len(resource.Regulatory) >= numRegulatory || len(resource.Regulatory[numRegulatory].Substitution) >= numSubstitution {
		return StringInput("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Substitution["+strconv.Itoa(numSubstitution)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Substitution["+strconv.Itoa(numSubstitution)+"].Id", resource.Regulatory[numRegulatory].Substitution[numSubstitution].Id)
}
func (resource *MedicationKnowledge) T_RegulatorySubstitutionType(numRegulatory int, numSubstitution int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Regulatory) >= numRegulatory || len(resource.Regulatory[numRegulatory].Substitution) >= numSubstitution {
		return CodeableConceptSelect("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Substitution["+strconv.Itoa(numSubstitution)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Substitution["+strconv.Itoa(numSubstitution)+"].Type", &resource.Regulatory[numRegulatory].Substitution[numSubstitution].Type, optionsValueSet)
}
func (resource *MedicationKnowledge) T_RegulatorySubstitutionAllowed(numRegulatory int, numSubstitution int) templ.Component {

	if resource == nil || len(resource.Regulatory) >= numRegulatory || len(resource.Regulatory[numRegulatory].Substitution) >= numSubstitution {
		return BoolInput("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Substitution["+strconv.Itoa(numSubstitution)+"].Allowed", nil)
	}
	return BoolInput("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Substitution["+strconv.Itoa(numSubstitution)+"].Allowed", &resource.Regulatory[numRegulatory].Substitution[numSubstitution].Allowed)
}
func (resource *MedicationKnowledge) T_RegulatoryScheduleId(numRegulatory int, numSchedule int) templ.Component {

	if resource == nil || len(resource.Regulatory) >= numRegulatory || len(resource.Regulatory[numRegulatory].Schedule) >= numSchedule {
		return StringInput("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Schedule["+strconv.Itoa(numSchedule)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Schedule["+strconv.Itoa(numSchedule)+"].Id", resource.Regulatory[numRegulatory].Schedule[numSchedule].Id)
}
func (resource *MedicationKnowledge) T_RegulatoryScheduleSchedule(numRegulatory int, numSchedule int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Regulatory) >= numRegulatory || len(resource.Regulatory[numRegulatory].Schedule) >= numSchedule {
		return CodeableConceptSelect("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Schedule["+strconv.Itoa(numSchedule)+"].Schedule", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Schedule["+strconv.Itoa(numSchedule)+"].Schedule", &resource.Regulatory[numRegulatory].Schedule[numSchedule].Schedule, optionsValueSet)
}
func (resource *MedicationKnowledge) T_RegulatoryMaxDispenseId(numRegulatory int) templ.Component {

	if resource == nil || len(resource.Regulatory) >= numRegulatory {
		return StringInput("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].MaxDispense.Id", nil)
	}
	return StringInput("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].MaxDispense.Id", resource.Regulatory[numRegulatory].MaxDispense.Id)
}
func (resource *MedicationKnowledge) T_KineticsId(numKinetics int) templ.Component {

	if resource == nil || len(resource.Kinetics) >= numKinetics {
		return StringInput("MedicationKnowledge.Kinetics["+strconv.Itoa(numKinetics)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.Kinetics["+strconv.Itoa(numKinetics)+"].Id", resource.Kinetics[numKinetics].Id)
}
