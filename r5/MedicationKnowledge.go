package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledge struct {
	Id                         *string                                         `json:"id,omitempty"`
	Meta                       *Meta                                           `json:"meta,omitempty"`
	ImplicitRules              *string                                         `json:"implicitRules,omitempty"`
	Language                   *string                                         `json:"language,omitempty"`
	Text                       *Narrative                                      `json:"text,omitempty"`
	Contained                  []Resource                                      `json:"contained,omitempty"`
	Extension                  []Extension                                     `json:"extension,omitempty"`
	ModifierExtension          []Extension                                     `json:"modifierExtension,omitempty"`
	Identifier                 []Identifier                                    `json:"identifier,omitempty"`
	Code                       *CodeableConcept                                `json:"code,omitempty"`
	Status                     *string                                         `json:"status,omitempty"`
	Author                     *Reference                                      `json:"author,omitempty"`
	IntendedJurisdiction       []CodeableConcept                               `json:"intendedJurisdiction,omitempty"`
	Name                       []string                                        `json:"name,omitempty"`
	RelatedMedicationKnowledge []MedicationKnowledgeRelatedMedicationKnowledge `json:"relatedMedicationKnowledge,omitempty"`
	AssociatedMedication       []Reference                                     `json:"associatedMedication,omitempty"`
	ProductType                []CodeableConcept                               `json:"productType,omitempty"`
	Monograph                  []MedicationKnowledgeMonograph                  `json:"monograph,omitempty"`
	PreparationInstruction     *string                                         `json:"preparationInstruction,omitempty"`
	Cost                       []MedicationKnowledgeCost                       `json:"cost,omitempty"`
	MonitoringProgram          []MedicationKnowledgeMonitoringProgram          `json:"monitoringProgram,omitempty"`
	IndicationGuideline        []MedicationKnowledgeIndicationGuideline        `json:"indicationGuideline,omitempty"`
	MedicineClassification     []MedicationKnowledgeMedicineClassification     `json:"medicineClassification,omitempty"`
	Packaging                  []MedicationKnowledgePackaging                  `json:"packaging,omitempty"`
	ClinicalUseIssue           []Reference                                     `json:"clinicalUseIssue,omitempty"`
	StorageGuideline           []MedicationKnowledgeStorageGuideline           `json:"storageGuideline,omitempty"`
	Regulatory                 []MedicationKnowledgeRegulatory                 `json:"regulatory,omitempty"`
	Definitional               *MedicationKnowledgeDefinitional                `json:"definitional,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeRelatedMedicationKnowledge struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Reference         []Reference     `json:"reference"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeMonograph struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Source            *Reference       `json:"source,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeCost struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	EffectiveDate       []Period        `json:"effectiveDate,omitempty"`
	Type                CodeableConcept `json:"type"`
	Source              *string         `json:"source,omitempty"`
	CostMoney           Money           `json:"costMoney"`
	CostCodeableConcept CodeableConcept `json:"costCodeableConcept"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeMonitoringProgram struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Name              *string          `json:"name,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeIndicationGuideline struct {
	Id                *string                                                 `json:"id,omitempty"`
	Extension         []Extension                                             `json:"extension,omitempty"`
	ModifierExtension []Extension                                             `json:"modifierExtension,omitempty"`
	Indication        []CodeableReference                                     `json:"indication,omitempty"`
	DosingGuideline   []MedicationKnowledgeIndicationGuidelineDosingGuideline `json:"dosingGuideline,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeIndicationGuidelineDosingGuideline struct {
	Id                      *string                                                                      `json:"id,omitempty"`
	Extension               []Extension                                                                  `json:"extension,omitempty"`
	ModifierExtension       []Extension                                                                  `json:"modifierExtension,omitempty"`
	TreatmentIntent         *CodeableConcept                                                             `json:"treatmentIntent,omitempty"`
	Dosage                  []MedicationKnowledgeIndicationGuidelineDosingGuidelineDosage                `json:"dosage,omitempty"`
	AdministrationTreatment *CodeableConcept                                                             `json:"administrationTreatment,omitempty"`
	PatientCharacteristic   []MedicationKnowledgeIndicationGuidelineDosingGuidelinePatientCharacteristic `json:"patientCharacteristic,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeIndicationGuidelineDosingGuidelineDosage struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Dosage            []Dosage        `json:"dosage"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeIndicationGuidelineDosingGuidelinePatientCharacteristic struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept  `json:"type"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range           `json:"valueRange,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeMedicineClassification struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	SourceString      *string           `json:"sourceString,omitempty"`
	SourceUri         *string           `json:"sourceUri,omitempty"`
	Classification    []CodeableConcept `json:"classification,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgePackaging struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	PackagedProduct   *Reference  `json:"packagedProduct,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeStorageGuideline struct {
	Id                   *string                                                   `json:"id,omitempty"`
	Extension            []Extension                                               `json:"extension,omitempty"`
	ModifierExtension    []Extension                                               `json:"modifierExtension,omitempty"`
	Reference            *string                                                   `json:"reference,omitempty"`
	Note                 []Annotation                                              `json:"note,omitempty"`
	StabilityDuration    *Duration                                                 `json:"stabilityDuration,omitempty"`
	EnvironmentalSetting []MedicationKnowledgeStorageGuidelineEnvironmentalSetting `json:"environmentalSetting,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeStorageGuidelineEnvironmentalSetting struct {
	Id                   *string         `json:"id,omitempty"`
	Extension            []Extension     `json:"extension,omitempty"`
	ModifierExtension    []Extension     `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept `json:"type"`
	ValueQuantity        Quantity        `json:"valueQuantity"`
	ValueRange           Range           `json:"valueRange"`
	ValueCodeableConcept CodeableConcept `json:"valueCodeableConcept"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeRegulatory struct {
	Id                  *string                                     `json:"id,omitempty"`
	Extension           []Extension                                 `json:"extension,omitempty"`
	ModifierExtension   []Extension                                 `json:"modifierExtension,omitempty"`
	RegulatoryAuthority Reference                                   `json:"regulatoryAuthority"`
	Substitution        []MedicationKnowledgeRegulatorySubstitution `json:"substitution,omitempty"`
	Schedule            []CodeableConcept                           `json:"schedule,omitempty"`
	MaxDispense         *MedicationKnowledgeRegulatoryMaxDispense   `json:"maxDispense,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeRegulatorySubstitution struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Allowed           bool            `json:"allowed"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeRegulatoryMaxDispense struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Quantity          Quantity    `json:"quantity"`
	Period            *Duration   `json:"period,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeDefinitional struct {
	Id                 *string                                             `json:"id,omitempty"`
	Extension          []Extension                                         `json:"extension,omitempty"`
	ModifierExtension  []Extension                                         `json:"modifierExtension,omitempty"`
	Definition         []Reference                                         `json:"definition,omitempty"`
	DoseForm           *CodeableConcept                                    `json:"doseForm,omitempty"`
	IntendedRoute      []CodeableConcept                                   `json:"intendedRoute,omitempty"`
	Ingredient         []MedicationKnowledgeDefinitionalIngredient         `json:"ingredient,omitempty"`
	DrugCharacteristic []MedicationKnowledgeDefinitionalDrugCharacteristic `json:"drugCharacteristic,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeDefinitionalIngredient struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Item                    CodeableReference `json:"item"`
	Type                    *CodeableConcept  `json:"type,omitempty"`
	StrengthRatio           *Ratio            `json:"strengthRatio,omitempty"`
	StrengthCodeableConcept *CodeableConcept  `json:"strengthCodeableConcept,omitempty"`
	StrengthQuantity        *Quantity         `json:"strengthQuantity,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeDefinitionalDrugCharacteristic struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Type                 *CodeableConcept `json:"type,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string          `json:"valueString,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueBase64Binary    *string          `json:"valueBase64Binary,omitempty"`
	ValueAttachment      *Attachment      `json:"valueAttachment,omitempty"`
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
func (resource *MedicationKnowledge) T_IntendedJurisdiction(numIntendedJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.IntendedJurisdiction) >= numIntendedJurisdiction {
		return CodeableConceptSelect("MedicationKnowledge.IntendedJurisdiction["+strconv.Itoa(numIntendedJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.IntendedJurisdiction["+strconv.Itoa(numIntendedJurisdiction)+"]", &resource.IntendedJurisdiction[numIntendedJurisdiction], optionsValueSet)
}
func (resource *MedicationKnowledge) T_Name(numName int) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return StringInput("MedicationKnowledge.Name["+strconv.Itoa(numName)+"]", nil)
	}
	return StringInput("MedicationKnowledge.Name["+strconv.Itoa(numName)+"]", &resource.Name[numName])
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
func (resource *MedicationKnowledge) T_IndicationGuidelineId(numIndicationGuideline int) templ.Component {

	if resource == nil || len(resource.IndicationGuideline) >= numIndicationGuideline {
		return StringInput("MedicationKnowledge.IndicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.IndicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].Id", resource.IndicationGuideline[numIndicationGuideline].Id)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelineId(numIndicationGuideline int, numDosingGuideline int) templ.Component {

	if resource == nil || len(resource.IndicationGuideline) >= numIndicationGuideline || len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) >= numDosingGuideline {
		return StringInput("MedicationKnowledge.IndicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].DosingGuideline["+strconv.Itoa(numDosingGuideline)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.IndicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].DosingGuideline["+strconv.Itoa(numDosingGuideline)+"].Id", resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].Id)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelineTreatmentIntent(numIndicationGuideline int, numDosingGuideline int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.IndicationGuideline) >= numIndicationGuideline || len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) >= numDosingGuideline {
		return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].DosingGuideline["+strconv.Itoa(numDosingGuideline)+"].TreatmentIntent", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].DosingGuideline["+strconv.Itoa(numDosingGuideline)+"].TreatmentIntent", resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].TreatmentIntent, optionsValueSet)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelineAdministrationTreatment(numIndicationGuideline int, numDosingGuideline int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.IndicationGuideline) >= numIndicationGuideline || len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) >= numDosingGuideline {
		return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].DosingGuideline["+strconv.Itoa(numDosingGuideline)+"].AdministrationTreatment", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].DosingGuideline["+strconv.Itoa(numDosingGuideline)+"].AdministrationTreatment", resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].AdministrationTreatment, optionsValueSet)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelineDosageId(numIndicationGuideline int, numDosingGuideline int, numDosage int) templ.Component {

	if resource == nil || len(resource.IndicationGuideline) >= numIndicationGuideline || len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) >= numDosingGuideline || len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].Dosage) >= numDosage {
		return StringInput("MedicationKnowledge.IndicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].DosingGuideline["+strconv.Itoa(numDosingGuideline)+"].Dosage["+strconv.Itoa(numDosage)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.IndicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].DosingGuideline["+strconv.Itoa(numDosingGuideline)+"].Dosage["+strconv.Itoa(numDosage)+"].Id", resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].Dosage[numDosage].Id)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelineDosageType(numIndicationGuideline int, numDosingGuideline int, numDosage int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.IndicationGuideline) >= numIndicationGuideline || len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) >= numDosingGuideline || len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].Dosage) >= numDosage {
		return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].DosingGuideline["+strconv.Itoa(numDosingGuideline)+"].Dosage["+strconv.Itoa(numDosage)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].DosingGuideline["+strconv.Itoa(numDosingGuideline)+"].Dosage["+strconv.Itoa(numDosage)+"].Type", &resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].Dosage[numDosage].Type, optionsValueSet)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelinePatientCharacteristicId(numIndicationGuideline int, numDosingGuideline int, numPatientCharacteristic int) templ.Component {

	if resource == nil || len(resource.IndicationGuideline) >= numIndicationGuideline || len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) >= numDosingGuideline || len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].PatientCharacteristic) >= numPatientCharacteristic {
		return StringInput("MedicationKnowledge.IndicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].DosingGuideline["+strconv.Itoa(numDosingGuideline)+"].PatientCharacteristic["+strconv.Itoa(numPatientCharacteristic)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.IndicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].DosingGuideline["+strconv.Itoa(numDosingGuideline)+"].PatientCharacteristic["+strconv.Itoa(numPatientCharacteristic)+"].Id", resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].PatientCharacteristic[numPatientCharacteristic].Id)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelinePatientCharacteristicType(numIndicationGuideline int, numDosingGuideline int, numPatientCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.IndicationGuideline) >= numIndicationGuideline || len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) >= numDosingGuideline || len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].PatientCharacteristic) >= numPatientCharacteristic {
		return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].DosingGuideline["+strconv.Itoa(numDosingGuideline)+"].PatientCharacteristic["+strconv.Itoa(numPatientCharacteristic)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].DosingGuideline["+strconv.Itoa(numDosingGuideline)+"].PatientCharacteristic["+strconv.Itoa(numPatientCharacteristic)+"].Type", &resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].PatientCharacteristic[numPatientCharacteristic].Type, optionsValueSet)
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
func (resource *MedicationKnowledge) T_PackagingId(numPackaging int) templ.Component {

	if resource == nil || len(resource.Packaging) >= numPackaging {
		return StringInput("MedicationKnowledge.Packaging["+strconv.Itoa(numPackaging)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.Packaging["+strconv.Itoa(numPackaging)+"].Id", resource.Packaging[numPackaging].Id)
}
func (resource *MedicationKnowledge) T_StorageGuidelineId(numStorageGuideline int) templ.Component {

	if resource == nil || len(resource.StorageGuideline) >= numStorageGuideline {
		return StringInput("MedicationKnowledge.StorageGuideline["+strconv.Itoa(numStorageGuideline)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.StorageGuideline["+strconv.Itoa(numStorageGuideline)+"].Id", resource.StorageGuideline[numStorageGuideline].Id)
}
func (resource *MedicationKnowledge) T_StorageGuidelineReference(numStorageGuideline int) templ.Component {

	if resource == nil || len(resource.StorageGuideline) >= numStorageGuideline {
		return StringInput("MedicationKnowledge.StorageGuideline["+strconv.Itoa(numStorageGuideline)+"].Reference", nil)
	}
	return StringInput("MedicationKnowledge.StorageGuideline["+strconv.Itoa(numStorageGuideline)+"].Reference", resource.StorageGuideline[numStorageGuideline].Reference)
}
func (resource *MedicationKnowledge) T_StorageGuidelineEnvironmentalSettingId(numStorageGuideline int, numEnvironmentalSetting int) templ.Component {

	if resource == nil || len(resource.StorageGuideline) >= numStorageGuideline || len(resource.StorageGuideline[numStorageGuideline].EnvironmentalSetting) >= numEnvironmentalSetting {
		return StringInput("MedicationKnowledge.StorageGuideline["+strconv.Itoa(numStorageGuideline)+"].EnvironmentalSetting["+strconv.Itoa(numEnvironmentalSetting)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.StorageGuideline["+strconv.Itoa(numStorageGuideline)+"].EnvironmentalSetting["+strconv.Itoa(numEnvironmentalSetting)+"].Id", resource.StorageGuideline[numStorageGuideline].EnvironmentalSetting[numEnvironmentalSetting].Id)
}
func (resource *MedicationKnowledge) T_StorageGuidelineEnvironmentalSettingType(numStorageGuideline int, numEnvironmentalSetting int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.StorageGuideline) >= numStorageGuideline || len(resource.StorageGuideline[numStorageGuideline].EnvironmentalSetting) >= numEnvironmentalSetting {
		return CodeableConceptSelect("MedicationKnowledge.StorageGuideline["+strconv.Itoa(numStorageGuideline)+"].EnvironmentalSetting["+strconv.Itoa(numEnvironmentalSetting)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.StorageGuideline["+strconv.Itoa(numStorageGuideline)+"].EnvironmentalSetting["+strconv.Itoa(numEnvironmentalSetting)+"].Type", &resource.StorageGuideline[numStorageGuideline].EnvironmentalSetting[numEnvironmentalSetting].Type, optionsValueSet)
}
func (resource *MedicationKnowledge) T_RegulatoryId(numRegulatory int) templ.Component {

	if resource == nil || len(resource.Regulatory) >= numRegulatory {
		return StringInput("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Id", resource.Regulatory[numRegulatory].Id)
}
func (resource *MedicationKnowledge) T_RegulatorySchedule(numRegulatory int, numSchedule int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Regulatory) >= numRegulatory || len(resource.Regulatory[numRegulatory].Schedule) >= numSchedule {
		return CodeableConceptSelect("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Schedule["+strconv.Itoa(numSchedule)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Schedule["+strconv.Itoa(numSchedule)+"]", &resource.Regulatory[numRegulatory].Schedule[numSchedule], optionsValueSet)
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
func (resource *MedicationKnowledge) T_RegulatoryMaxDispenseId(numRegulatory int) templ.Component {

	if resource == nil || len(resource.Regulatory) >= numRegulatory {
		return StringInput("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].MaxDispense.Id", nil)
	}
	return StringInput("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].MaxDispense.Id", resource.Regulatory[numRegulatory].MaxDispense.Id)
}
func (resource *MedicationKnowledge) T_DefinitionalId() templ.Component {

	if resource == nil {
		return StringInput("MedicationKnowledge.Definitional.Id", nil)
	}
	return StringInput("MedicationKnowledge.Definitional.Id", resource.Definitional.Id)
}
func (resource *MedicationKnowledge) T_DefinitionalDoseForm(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationKnowledge.Definitional.DoseForm", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.Definitional.DoseForm", resource.Definitional.DoseForm, optionsValueSet)
}
func (resource *MedicationKnowledge) T_DefinitionalIntendedRoute(numIntendedRoute int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Definitional.IntendedRoute) >= numIntendedRoute {
		return CodeableConceptSelect("MedicationKnowledge.Definitional.IntendedRoute["+strconv.Itoa(numIntendedRoute)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.Definitional.IntendedRoute["+strconv.Itoa(numIntendedRoute)+"]", &resource.Definitional.IntendedRoute[numIntendedRoute], optionsValueSet)
}
func (resource *MedicationKnowledge) T_DefinitionalIngredientId(numIngredient int) templ.Component {

	if resource == nil || len(resource.Definitional.Ingredient) >= numIngredient {
		return StringInput("MedicationKnowledge.Definitional.Ingredient["+strconv.Itoa(numIngredient)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.Definitional.Ingredient["+strconv.Itoa(numIngredient)+"].Id", resource.Definitional.Ingredient[numIngredient].Id)
}
func (resource *MedicationKnowledge) T_DefinitionalIngredientType(numIngredient int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Definitional.Ingredient) >= numIngredient {
		return CodeableConceptSelect("MedicationKnowledge.Definitional.Ingredient["+strconv.Itoa(numIngredient)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.Definitional.Ingredient["+strconv.Itoa(numIngredient)+"].Type", resource.Definitional.Ingredient[numIngredient].Type, optionsValueSet)
}
func (resource *MedicationKnowledge) T_DefinitionalDrugCharacteristicId(numDrugCharacteristic int) templ.Component {

	if resource == nil || len(resource.Definitional.DrugCharacteristic) >= numDrugCharacteristic {
		return StringInput("MedicationKnowledge.Definitional.DrugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].Id", nil)
	}
	return StringInput("MedicationKnowledge.Definitional.DrugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].Id", resource.Definitional.DrugCharacteristic[numDrugCharacteristic].Id)
}
func (resource *MedicationKnowledge) T_DefinitionalDrugCharacteristicType(numDrugCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Definitional.DrugCharacteristic) >= numDrugCharacteristic {
		return CodeableConceptSelect("MedicationKnowledge.Definitional.DrugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationKnowledge.Definitional.DrugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].Type", resource.Definitional.DrugCharacteristic[numDrugCharacteristic].Type, optionsValueSet)
}
