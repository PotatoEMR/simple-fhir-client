package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
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

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeRelatedMedicationKnowledge struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Reference         []Reference     `json:"reference"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeMonograph struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Source            *Reference       `json:"source,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeIngredient struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	ItemCodeableConcept CodeableConcept `json:"itemCodeableConcept"`
	ItemReference       Reference       `json:"itemReference"`
	IsActive            *bool           `json:"isActive,omitempty"`
	Strength            *Ratio          `json:"strength,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeCost struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Source            *string         `json:"source,omitempty"`
	Cost              Money           `json:"cost"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeMonitoringProgram struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Name              *string          `json:"name,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeAdministrationGuidelines struct {
	Id                        *string                                                             `json:"id,omitempty"`
	Extension                 []Extension                                                         `json:"extension,omitempty"`
	ModifierExtension         []Extension                                                         `json:"modifierExtension,omitempty"`
	Dosage                    []MedicationKnowledgeAdministrationGuidelinesDosage                 `json:"dosage,omitempty"`
	IndicationCodeableConcept *CodeableConcept                                                    `json:"indicationCodeableConcept,omitempty"`
	IndicationReference       *Reference                                                          `json:"indicationReference,omitempty"`
	PatientCharacteristics    []MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics `json:"patientCharacteristics,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeAdministrationGuidelinesDosage struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Dosage            []Dosage        `json:"dosage"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics struct {
	Id                            *string         `json:"id,omitempty"`
	Extension                     []Extension     `json:"extension,omitempty"`
	ModifierExtension             []Extension     `json:"modifierExtension,omitempty"`
	CharacteristicCodeableConcept CodeableConcept `json:"characteristicCodeableConcept"`
	CharacteristicQuantity        Quantity        `json:"characteristicQuantity"`
	Value                         []string        `json:"value,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeMedicineClassification struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	Classification    []CodeableConcept `json:"classification,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
type MedicationKnowledgePackaging struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Quantity          *Quantity        `json:"quantity,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
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

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeRegulatory struct {
	Id                  *string                                     `json:"id,omitempty"`
	Extension           []Extension                                 `json:"extension,omitempty"`
	ModifierExtension   []Extension                                 `json:"modifierExtension,omitempty"`
	RegulatoryAuthority Reference                                   `json:"regulatoryAuthority"`
	Substitution        []MedicationKnowledgeRegulatorySubstitution `json:"substitution,omitempty"`
	Schedule            []MedicationKnowledgeRegulatorySchedule     `json:"schedule,omitempty"`
	MaxDispense         *MedicationKnowledgeRegulatoryMaxDispense   `json:"maxDispense,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeRegulatorySubstitution struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Allowed           bool            `json:"allowed"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeRegulatorySchedule struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Schedule          CodeableConcept `json:"schedule"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeRegulatoryMaxDispense struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Quantity          Quantity    `json:"quantity"`
	Period            *Duration   `json:"period,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationKnowledge
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
func (r MedicationKnowledge) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicationKnowledge/" + *r.Id
		ref.Reference = &refStr
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
func (resource *MedicationKnowledge) T_DoseForm(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("MedicationKnowledge.DoseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.DoseForm", resource.DoseForm, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_Synonym(numSynonym int, htmlAttrs string) templ.Component {
	if resource == nil || numSynonym >= len(resource.Synonym) {
		return StringInput("MedicationKnowledge.Synonym["+strconv.Itoa(numSynonym)+"]", nil, htmlAttrs)
	}
	return StringInput("MedicationKnowledge.Synonym["+strconv.Itoa(numSynonym)+"]", &resource.Synonym[numSynonym], htmlAttrs)
}
func (resource *MedicationKnowledge) T_ProductType(numProductType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProductType >= len(resource.ProductType) {
		return CodeableConceptSelect("MedicationKnowledge.ProductType["+strconv.Itoa(numProductType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.ProductType["+strconv.Itoa(numProductType)+"]", &resource.ProductType[numProductType], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_PreparationInstruction(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("MedicationKnowledge.PreparationInstruction", nil, htmlAttrs)
	}
	return StringInput("MedicationKnowledge.PreparationInstruction", resource.PreparationInstruction, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IntendedRoute(numIntendedRoute int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numIntendedRoute >= len(resource.IntendedRoute) {
		return CodeableConceptSelect("MedicationKnowledge.IntendedRoute["+strconv.Itoa(numIntendedRoute)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.IntendedRoute["+strconv.Itoa(numIntendedRoute)+"]", &resource.IntendedRoute[numIntendedRoute], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_RelatedMedicationKnowledgeType(numRelatedMedicationKnowledge int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelatedMedicationKnowledge >= len(resource.RelatedMedicationKnowledge) {
		return CodeableConceptSelect("MedicationKnowledge.RelatedMedicationKnowledge["+strconv.Itoa(numRelatedMedicationKnowledge)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.RelatedMedicationKnowledge["+strconv.Itoa(numRelatedMedicationKnowledge)+"].Type", &resource.RelatedMedicationKnowledge[numRelatedMedicationKnowledge].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MonographType(numMonograph int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMonograph >= len(resource.Monograph) {
		return CodeableConceptSelect("MedicationKnowledge.Monograph["+strconv.Itoa(numMonograph)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Monograph["+strconv.Itoa(numMonograph)+"].Type", resource.Monograph[numMonograph].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IngredientItemCodeableConcept(numIngredient int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableConceptSelect("MedicationKnowledge.Ingredient["+strconv.Itoa(numIngredient)+"].ItemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Ingredient["+strconv.Itoa(numIngredient)+"].ItemCodeableConcept", &resource.Ingredient[numIngredient].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IngredientIsActive(numIngredient int, htmlAttrs string) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return BoolInput("MedicationKnowledge.Ingredient["+strconv.Itoa(numIngredient)+"].IsActive", nil, htmlAttrs)
	}
	return BoolInput("MedicationKnowledge.Ingredient["+strconv.Itoa(numIngredient)+"].IsActive", resource.Ingredient[numIngredient].IsActive, htmlAttrs)
}
func (resource *MedicationKnowledge) T_CostType(numCost int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCost >= len(resource.Cost) {
		return CodeableConceptSelect("MedicationKnowledge.Cost["+strconv.Itoa(numCost)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Cost["+strconv.Itoa(numCost)+"].Type", &resource.Cost[numCost].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_CostSource(numCost int, htmlAttrs string) templ.Component {
	if resource == nil || numCost >= len(resource.Cost) {
		return StringInput("MedicationKnowledge.Cost["+strconv.Itoa(numCost)+"].Source", nil, htmlAttrs)
	}
	return StringInput("MedicationKnowledge.Cost["+strconv.Itoa(numCost)+"].Source", resource.Cost[numCost].Source, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MonitoringProgramType(numMonitoringProgram int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMonitoringProgram >= len(resource.MonitoringProgram) {
		return CodeableConceptSelect("MedicationKnowledge.MonitoringProgram["+strconv.Itoa(numMonitoringProgram)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.MonitoringProgram["+strconv.Itoa(numMonitoringProgram)+"].Type", resource.MonitoringProgram[numMonitoringProgram].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MonitoringProgramName(numMonitoringProgram int, htmlAttrs string) templ.Component {
	if resource == nil || numMonitoringProgram >= len(resource.MonitoringProgram) {
		return StringInput("MedicationKnowledge.MonitoringProgram["+strconv.Itoa(numMonitoringProgram)+"].Name", nil, htmlAttrs)
	}
	return StringInput("MedicationKnowledge.MonitoringProgram["+strconv.Itoa(numMonitoringProgram)+"].Name", resource.MonitoringProgram[numMonitoringProgram].Name, htmlAttrs)
}
func (resource *MedicationKnowledge) T_AdministrationGuidelinesIndicationCodeableConcept(numAdministrationGuidelines int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAdministrationGuidelines >= len(resource.AdministrationGuidelines) {
		return CodeableConceptSelect("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].IndicationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].IndicationCodeableConcept", resource.AdministrationGuidelines[numAdministrationGuidelines].IndicationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_AdministrationGuidelinesDosageType(numAdministrationGuidelines int, numDosage int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAdministrationGuidelines >= len(resource.AdministrationGuidelines) || numDosage >= len(resource.AdministrationGuidelines[numAdministrationGuidelines].Dosage) {
		return CodeableConceptSelect("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].Dosage["+strconv.Itoa(numDosage)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].Dosage["+strconv.Itoa(numDosage)+"].Type", &resource.AdministrationGuidelines[numAdministrationGuidelines].Dosage[numDosage].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_AdministrationGuidelinesPatientCharacteristicsCharacteristicCodeableConcept(numAdministrationGuidelines int, numPatientCharacteristics int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAdministrationGuidelines >= len(resource.AdministrationGuidelines) || numPatientCharacteristics >= len(resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics) {
		return CodeableConceptSelect("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].PatientCharacteristics["+strconv.Itoa(numPatientCharacteristics)+"].CharacteristicCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].PatientCharacteristics["+strconv.Itoa(numPatientCharacteristics)+"].CharacteristicCodeableConcept", &resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics[numPatientCharacteristics].CharacteristicCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_AdministrationGuidelinesPatientCharacteristicsValue(numAdministrationGuidelines int, numPatientCharacteristics int, numValue int, htmlAttrs string) templ.Component {
	if resource == nil || numAdministrationGuidelines >= len(resource.AdministrationGuidelines) || numPatientCharacteristics >= len(resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics) || numValue >= len(resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics[numPatientCharacteristics].Value) {
		return StringInput("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].PatientCharacteristics["+strconv.Itoa(numPatientCharacteristics)+"].Value["+strconv.Itoa(numValue)+"]", nil, htmlAttrs)
	}
	return StringInput("MedicationKnowledge.AdministrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].PatientCharacteristics["+strconv.Itoa(numPatientCharacteristics)+"].Value["+strconv.Itoa(numValue)+"]", &resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics[numPatientCharacteristics].Value[numValue], htmlAttrs)
}
func (resource *MedicationKnowledge) T_MedicineClassificationType(numMedicineClassification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMedicineClassification >= len(resource.MedicineClassification) {
		return CodeableConceptSelect("MedicationKnowledge.MedicineClassification["+strconv.Itoa(numMedicineClassification)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.MedicineClassification["+strconv.Itoa(numMedicineClassification)+"].Type", &resource.MedicineClassification[numMedicineClassification].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MedicineClassificationClassification(numMedicineClassification int, numClassification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMedicineClassification >= len(resource.MedicineClassification) || numClassification >= len(resource.MedicineClassification[numMedicineClassification].Classification) {
		return CodeableConceptSelect("MedicationKnowledge.MedicineClassification["+strconv.Itoa(numMedicineClassification)+"].Classification["+strconv.Itoa(numClassification)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.MedicineClassification["+strconv.Itoa(numMedicineClassification)+"].Classification["+strconv.Itoa(numClassification)+"]", &resource.MedicineClassification[numMedicineClassification].Classification[numClassification], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_PackagingType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("MedicationKnowledge.Packaging.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Packaging.Type", resource.Packaging.Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DrugCharacteristicType(numDrugCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDrugCharacteristic >= len(resource.DrugCharacteristic) {
		return CodeableConceptSelect("MedicationKnowledge.DrugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.DrugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].Type", resource.DrugCharacteristic[numDrugCharacteristic].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DrugCharacteristicValueCodeableConcept(numDrugCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDrugCharacteristic >= len(resource.DrugCharacteristic) {
		return CodeableConceptSelect("MedicationKnowledge.DrugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.DrugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].ValueCodeableConcept", resource.DrugCharacteristic[numDrugCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DrugCharacteristicValueString(numDrugCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numDrugCharacteristic >= len(resource.DrugCharacteristic) {
		return StringInput("MedicationKnowledge.DrugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].ValueString", nil, htmlAttrs)
	}
	return StringInput("MedicationKnowledge.DrugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].ValueString", resource.DrugCharacteristic[numDrugCharacteristic].ValueString, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DrugCharacteristicValueBase64Binary(numDrugCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numDrugCharacteristic >= len(resource.DrugCharacteristic) {
		return StringInput("MedicationKnowledge.DrugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].ValueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("MedicationKnowledge.DrugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].ValueBase64Binary", resource.DrugCharacteristic[numDrugCharacteristic].ValueBase64Binary, htmlAttrs)
}
func (resource *MedicationKnowledge) T_RegulatorySubstitutionType(numRegulatory int, numSubstitution int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRegulatory >= len(resource.Regulatory) || numSubstitution >= len(resource.Regulatory[numRegulatory].Substitution) {
		return CodeableConceptSelect("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Substitution["+strconv.Itoa(numSubstitution)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Substitution["+strconv.Itoa(numSubstitution)+"].Type", &resource.Regulatory[numRegulatory].Substitution[numSubstitution].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_RegulatorySubstitutionAllowed(numRegulatory int, numSubstitution int, htmlAttrs string) templ.Component {
	if resource == nil || numRegulatory >= len(resource.Regulatory) || numSubstitution >= len(resource.Regulatory[numRegulatory].Substitution) {
		return BoolInput("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Substitution["+strconv.Itoa(numSubstitution)+"].Allowed", nil, htmlAttrs)
	}
	return BoolInput("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Substitution["+strconv.Itoa(numSubstitution)+"].Allowed", &resource.Regulatory[numRegulatory].Substitution[numSubstitution].Allowed, htmlAttrs)
}
func (resource *MedicationKnowledge) T_RegulatoryScheduleSchedule(numRegulatory int, numSchedule int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRegulatory >= len(resource.Regulatory) || numSchedule >= len(resource.Regulatory[numRegulatory].Schedule) {
		return CodeableConceptSelect("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Schedule["+strconv.Itoa(numSchedule)+"].Schedule", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationKnowledge.Regulatory["+strconv.Itoa(numRegulatory)+"].Schedule["+strconv.Itoa(numSchedule)+"].Schedule", &resource.Regulatory[numRegulatory].Schedule[numSchedule].Schedule, optionsValueSet, htmlAttrs)
}
