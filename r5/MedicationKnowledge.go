package r5

//generated with command go run ./bultaoreune
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
func (r MedicationKnowledge) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicationKnowledge/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "MedicationKnowledge"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MedicationKnowledge) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationKnowledge.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSMedicationknowledge_status

	if resource == nil {
		return CodeSelect("MedicationKnowledge.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("MedicationKnowledge.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IntendedJurisdiction(numIntendedJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numIntendedJurisdiction >= len(resource.IntendedJurisdiction) {
		return CodeableConceptSelect("MedicationKnowledge.IntendedJurisdiction."+strconv.Itoa(numIntendedJurisdiction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.IntendedJurisdiction."+strconv.Itoa(numIntendedJurisdiction)+".", &resource.IntendedJurisdiction[numIntendedJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_Name(numName int, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) {
		return StringInput("MedicationKnowledge.Name."+strconv.Itoa(numName)+".", nil, htmlAttrs)
	}
	return StringInput("MedicationKnowledge.Name."+strconv.Itoa(numName)+".", &resource.Name[numName], htmlAttrs)
}
func (resource *MedicationKnowledge) T_ProductType(numProductType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProductType >= len(resource.ProductType) {
		return CodeableConceptSelect("MedicationKnowledge.ProductType."+strconv.Itoa(numProductType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.ProductType."+strconv.Itoa(numProductType)+".", &resource.ProductType[numProductType], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_PreparationInstruction(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("MedicationKnowledge.PreparationInstruction", nil, htmlAttrs)
	}
	return StringInput("MedicationKnowledge.PreparationInstruction", resource.PreparationInstruction, htmlAttrs)
}
func (resource *MedicationKnowledge) T_RelatedMedicationKnowledgeType(numRelatedMedicationKnowledge int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numRelatedMedicationKnowledge >= len(resource.RelatedMedicationKnowledge) {
		return CodeableConceptSelect("MedicationKnowledge.RelatedMedicationKnowledge."+strconv.Itoa(numRelatedMedicationKnowledge)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.RelatedMedicationKnowledge."+strconv.Itoa(numRelatedMedicationKnowledge)+"..Type", &resource.RelatedMedicationKnowledge[numRelatedMedicationKnowledge].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MonographType(numMonograph int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numMonograph >= len(resource.Monograph) {
		return CodeableConceptSelect("MedicationKnowledge.Monograph."+strconv.Itoa(numMonograph)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Monograph."+strconv.Itoa(numMonograph)+"..Type", resource.Monograph[numMonograph].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_CostType(numCost int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCost >= len(resource.Cost) {
		return CodeableConceptSelect("MedicationKnowledge.Cost."+strconv.Itoa(numCost)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Cost."+strconv.Itoa(numCost)+"..Type", &resource.Cost[numCost].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_CostSource(numCost int, htmlAttrs string) templ.Component {

	if resource == nil || numCost >= len(resource.Cost) {
		return StringInput("MedicationKnowledge.Cost."+strconv.Itoa(numCost)+"..Source", nil, htmlAttrs)
	}
	return StringInput("MedicationKnowledge.Cost."+strconv.Itoa(numCost)+"..Source", resource.Cost[numCost].Source, htmlAttrs)
}
func (resource *MedicationKnowledge) T_CostCostCodeableConcept(numCost int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCost >= len(resource.Cost) {
		return CodeableConceptSelect("MedicationKnowledge.Cost."+strconv.Itoa(numCost)+"..CostCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Cost."+strconv.Itoa(numCost)+"..CostCodeableConcept", &resource.Cost[numCost].CostCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MonitoringProgramType(numMonitoringProgram int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numMonitoringProgram >= len(resource.MonitoringProgram) {
		return CodeableConceptSelect("MedicationKnowledge.MonitoringProgram."+strconv.Itoa(numMonitoringProgram)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.MonitoringProgram."+strconv.Itoa(numMonitoringProgram)+"..Type", resource.MonitoringProgram[numMonitoringProgram].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MonitoringProgramName(numMonitoringProgram int, htmlAttrs string) templ.Component {

	if resource == nil || numMonitoringProgram >= len(resource.MonitoringProgram) {
		return StringInput("MedicationKnowledge.MonitoringProgram."+strconv.Itoa(numMonitoringProgram)+"..Name", nil, htmlAttrs)
	}
	return StringInput("MedicationKnowledge.MonitoringProgram."+strconv.Itoa(numMonitoringProgram)+"..Name", resource.MonitoringProgram[numMonitoringProgram].Name, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelineTreatmentIntent(numIndicationGuideline int, numDosingGuideline int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numIndicationGuideline >= len(resource.IndicationGuideline) || numDosingGuideline >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) {
		return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline."+strconv.Itoa(numIndicationGuideline)+"..DosingGuideline."+strconv.Itoa(numDosingGuideline)+"..TreatmentIntent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline."+strconv.Itoa(numIndicationGuideline)+"..DosingGuideline."+strconv.Itoa(numDosingGuideline)+"..TreatmentIntent", resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].TreatmentIntent, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelineAdministrationTreatment(numIndicationGuideline int, numDosingGuideline int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numIndicationGuideline >= len(resource.IndicationGuideline) || numDosingGuideline >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) {
		return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline."+strconv.Itoa(numIndicationGuideline)+"..DosingGuideline."+strconv.Itoa(numDosingGuideline)+"..AdministrationTreatment", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline."+strconv.Itoa(numIndicationGuideline)+"..DosingGuideline."+strconv.Itoa(numDosingGuideline)+"..AdministrationTreatment", resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].AdministrationTreatment, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelineDosageType(numIndicationGuideline int, numDosingGuideline int, numDosage int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numIndicationGuideline >= len(resource.IndicationGuideline) || numDosingGuideline >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) || numDosage >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].Dosage) {
		return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline."+strconv.Itoa(numIndicationGuideline)+"..DosingGuideline."+strconv.Itoa(numDosingGuideline)+"..Dosage."+strconv.Itoa(numDosage)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline."+strconv.Itoa(numIndicationGuideline)+"..DosingGuideline."+strconv.Itoa(numDosingGuideline)+"..Dosage."+strconv.Itoa(numDosage)+"..Type", &resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].Dosage[numDosage].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelinePatientCharacteristicType(numIndicationGuideline int, numDosingGuideline int, numPatientCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numIndicationGuideline >= len(resource.IndicationGuideline) || numDosingGuideline >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) || numPatientCharacteristic >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].PatientCharacteristic) {
		return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline."+strconv.Itoa(numIndicationGuideline)+"..DosingGuideline."+strconv.Itoa(numDosingGuideline)+"..PatientCharacteristic."+strconv.Itoa(numPatientCharacteristic)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline."+strconv.Itoa(numIndicationGuideline)+"..DosingGuideline."+strconv.Itoa(numDosingGuideline)+"..PatientCharacteristic."+strconv.Itoa(numPatientCharacteristic)+"..Type", &resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].PatientCharacteristic[numPatientCharacteristic].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelinePatientCharacteristicValueCodeableConcept(numIndicationGuideline int, numDosingGuideline int, numPatientCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numIndicationGuideline >= len(resource.IndicationGuideline) || numDosingGuideline >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) || numPatientCharacteristic >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].PatientCharacteristic) {
		return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline."+strconv.Itoa(numIndicationGuideline)+"..DosingGuideline."+strconv.Itoa(numDosingGuideline)+"..PatientCharacteristic."+strconv.Itoa(numPatientCharacteristic)+"..ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.IndicationGuideline."+strconv.Itoa(numIndicationGuideline)+"..DosingGuideline."+strconv.Itoa(numDosingGuideline)+"..PatientCharacteristic."+strconv.Itoa(numPatientCharacteristic)+"..ValueCodeableConcept", resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].PatientCharacteristic[numPatientCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MedicineClassificationType(numMedicineClassification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numMedicineClassification >= len(resource.MedicineClassification) {
		return CodeableConceptSelect("MedicationKnowledge.MedicineClassification."+strconv.Itoa(numMedicineClassification)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.MedicineClassification."+strconv.Itoa(numMedicineClassification)+"..Type", &resource.MedicineClassification[numMedicineClassification].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MedicineClassificationSourceString(numMedicineClassification int, htmlAttrs string) templ.Component {

	if resource == nil || numMedicineClassification >= len(resource.MedicineClassification) {
		return StringInput("MedicationKnowledge.MedicineClassification."+strconv.Itoa(numMedicineClassification)+"..SourceString", nil, htmlAttrs)
	}
	return StringInput("MedicationKnowledge.MedicineClassification."+strconv.Itoa(numMedicineClassification)+"..SourceString", resource.MedicineClassification[numMedicineClassification].SourceString, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MedicineClassificationSourceUri(numMedicineClassification int, htmlAttrs string) templ.Component {

	if resource == nil || numMedicineClassification >= len(resource.MedicineClassification) {
		return StringInput("MedicationKnowledge.MedicineClassification."+strconv.Itoa(numMedicineClassification)+"..SourceUri", nil, htmlAttrs)
	}
	return StringInput("MedicationKnowledge.MedicineClassification."+strconv.Itoa(numMedicineClassification)+"..SourceUri", resource.MedicineClassification[numMedicineClassification].SourceUri, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MedicineClassificationClassification(numMedicineClassification int, numClassification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numMedicineClassification >= len(resource.MedicineClassification) || numClassification >= len(resource.MedicineClassification[numMedicineClassification].Classification) {
		return CodeableConceptSelect("MedicationKnowledge.MedicineClassification."+strconv.Itoa(numMedicineClassification)+"..Classification."+strconv.Itoa(numClassification)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.MedicineClassification."+strconv.Itoa(numMedicineClassification)+"..Classification."+strconv.Itoa(numClassification)+".", &resource.MedicineClassification[numMedicineClassification].Classification[numClassification], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_StorageGuidelineReference(numStorageGuideline int, htmlAttrs string) templ.Component {

	if resource == nil || numStorageGuideline >= len(resource.StorageGuideline) {
		return StringInput("MedicationKnowledge.StorageGuideline."+strconv.Itoa(numStorageGuideline)+"..Reference", nil, htmlAttrs)
	}
	return StringInput("MedicationKnowledge.StorageGuideline."+strconv.Itoa(numStorageGuideline)+"..Reference", resource.StorageGuideline[numStorageGuideline].Reference, htmlAttrs)
}
func (resource *MedicationKnowledge) T_StorageGuidelineNote(numStorageGuideline int, numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numStorageGuideline >= len(resource.StorageGuideline) || numNote >= len(resource.StorageGuideline[numStorageGuideline].Note) {
		return AnnotationTextArea("MedicationKnowledge.StorageGuideline."+strconv.Itoa(numStorageGuideline)+"..Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("MedicationKnowledge.StorageGuideline."+strconv.Itoa(numStorageGuideline)+"..Note."+strconv.Itoa(numNote)+".", &resource.StorageGuideline[numStorageGuideline].Note[numNote], htmlAttrs)
}
func (resource *MedicationKnowledge) T_StorageGuidelineEnvironmentalSettingType(numStorageGuideline int, numEnvironmentalSetting int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numStorageGuideline >= len(resource.StorageGuideline) || numEnvironmentalSetting >= len(resource.StorageGuideline[numStorageGuideline].EnvironmentalSetting) {
		return CodeableConceptSelect("MedicationKnowledge.StorageGuideline."+strconv.Itoa(numStorageGuideline)+"..EnvironmentalSetting."+strconv.Itoa(numEnvironmentalSetting)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.StorageGuideline."+strconv.Itoa(numStorageGuideline)+"..EnvironmentalSetting."+strconv.Itoa(numEnvironmentalSetting)+"..Type", &resource.StorageGuideline[numStorageGuideline].EnvironmentalSetting[numEnvironmentalSetting].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_StorageGuidelineEnvironmentalSettingValueCodeableConcept(numStorageGuideline int, numEnvironmentalSetting int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numStorageGuideline >= len(resource.StorageGuideline) || numEnvironmentalSetting >= len(resource.StorageGuideline[numStorageGuideline].EnvironmentalSetting) {
		return CodeableConceptSelect("MedicationKnowledge.StorageGuideline."+strconv.Itoa(numStorageGuideline)+"..EnvironmentalSetting."+strconv.Itoa(numEnvironmentalSetting)+"..ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.StorageGuideline."+strconv.Itoa(numStorageGuideline)+"..EnvironmentalSetting."+strconv.Itoa(numEnvironmentalSetting)+"..ValueCodeableConcept", &resource.StorageGuideline[numStorageGuideline].EnvironmentalSetting[numEnvironmentalSetting].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_RegulatorySchedule(numRegulatory int, numSchedule int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numRegulatory >= len(resource.Regulatory) || numSchedule >= len(resource.Regulatory[numRegulatory].Schedule) {
		return CodeableConceptSelect("MedicationKnowledge.Regulatory."+strconv.Itoa(numRegulatory)+"..Schedule."+strconv.Itoa(numSchedule)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Regulatory."+strconv.Itoa(numRegulatory)+"..Schedule."+strconv.Itoa(numSchedule)+".", &resource.Regulatory[numRegulatory].Schedule[numSchedule], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_RegulatorySubstitutionType(numRegulatory int, numSubstitution int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numRegulatory >= len(resource.Regulatory) || numSubstitution >= len(resource.Regulatory[numRegulatory].Substitution) {
		return CodeableConceptSelect("MedicationKnowledge.Regulatory."+strconv.Itoa(numRegulatory)+"..Substitution."+strconv.Itoa(numSubstitution)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Regulatory."+strconv.Itoa(numRegulatory)+"..Substitution."+strconv.Itoa(numSubstitution)+"..Type", &resource.Regulatory[numRegulatory].Substitution[numSubstitution].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_RegulatorySubstitutionAllowed(numRegulatory int, numSubstitution int, htmlAttrs string) templ.Component {

	if resource == nil || numRegulatory >= len(resource.Regulatory) || numSubstitution >= len(resource.Regulatory[numRegulatory].Substitution) {
		return BoolInput("MedicationKnowledge.Regulatory."+strconv.Itoa(numRegulatory)+"..Substitution."+strconv.Itoa(numSubstitution)+"..Allowed", nil, htmlAttrs)
	}
	return BoolInput("MedicationKnowledge.Regulatory."+strconv.Itoa(numRegulatory)+"..Substitution."+strconv.Itoa(numSubstitution)+"..Allowed", &resource.Regulatory[numRegulatory].Substitution[numSubstitution].Allowed, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalDoseForm(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationKnowledge.Definitional.DoseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Definitional.DoseForm", resource.Definitional.DoseForm, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalIntendedRoute(numIntendedRoute int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numIntendedRoute >= len(resource.Definitional.IntendedRoute) {
		return CodeableConceptSelect("MedicationKnowledge.Definitional.IntendedRoute."+strconv.Itoa(numIntendedRoute)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Definitional.IntendedRoute."+strconv.Itoa(numIntendedRoute)+".", &resource.Definitional.IntendedRoute[numIntendedRoute], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalIngredientType(numIngredient int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numIngredient >= len(resource.Definitional.Ingredient) {
		return CodeableConceptSelect("MedicationKnowledge.Definitional.Ingredient."+strconv.Itoa(numIngredient)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Definitional.Ingredient."+strconv.Itoa(numIngredient)+"..Type", resource.Definitional.Ingredient[numIngredient].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalIngredientStrengthCodeableConcept(numIngredient int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numIngredient >= len(resource.Definitional.Ingredient) {
		return CodeableConceptSelect("MedicationKnowledge.Definitional.Ingredient."+strconv.Itoa(numIngredient)+"..StrengthCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Definitional.Ingredient."+strconv.Itoa(numIngredient)+"..StrengthCodeableConcept", resource.Definitional.Ingredient[numIngredient].StrengthCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalDrugCharacteristicType(numDrugCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numDrugCharacteristic >= len(resource.Definitional.DrugCharacteristic) {
		return CodeableConceptSelect("MedicationKnowledge.Definitional.DrugCharacteristic."+strconv.Itoa(numDrugCharacteristic)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Definitional.DrugCharacteristic."+strconv.Itoa(numDrugCharacteristic)+"..Type", resource.Definitional.DrugCharacteristic[numDrugCharacteristic].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalDrugCharacteristicValueCodeableConcept(numDrugCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numDrugCharacteristic >= len(resource.Definitional.DrugCharacteristic) {
		return CodeableConceptSelect("MedicationKnowledge.Definitional.DrugCharacteristic."+strconv.Itoa(numDrugCharacteristic)+"..ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Definitional.DrugCharacteristic."+strconv.Itoa(numDrugCharacteristic)+"..ValueCodeableConcept", resource.Definitional.DrugCharacteristic[numDrugCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalDrugCharacteristicValueString(numDrugCharacteristic int, htmlAttrs string) templ.Component {

	if resource == nil || numDrugCharacteristic >= len(resource.Definitional.DrugCharacteristic) {
		return StringInput("MedicationKnowledge.Definitional.DrugCharacteristic."+strconv.Itoa(numDrugCharacteristic)+"..ValueString", nil, htmlAttrs)
	}
	return StringInput("MedicationKnowledge.Definitional.DrugCharacteristic."+strconv.Itoa(numDrugCharacteristic)+"..ValueString", resource.Definitional.DrugCharacteristic[numDrugCharacteristic].ValueString, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalDrugCharacteristicValueBase64Binary(numDrugCharacteristic int, htmlAttrs string) templ.Component {

	if resource == nil || numDrugCharacteristic >= len(resource.Definitional.DrugCharacteristic) {
		return StringInput("MedicationKnowledge.Definitional.DrugCharacteristic."+strconv.Itoa(numDrugCharacteristic)+"..ValueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("MedicationKnowledge.Definitional.DrugCharacteristic."+strconv.Itoa(numDrugCharacteristic)+"..ValueBase64Binary", resource.Definitional.DrugCharacteristic[numDrugCharacteristic].ValueBase64Binary, htmlAttrs)
}
