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

// struct -> json, automatically add resourceType=Patient
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
func (r MedicationKnowledge) ResourceType() string {
	return "MedicationKnowledge"
}

func (resource *MedicationKnowledge) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMedicationknowledge_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_Author(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "author", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "author", resource.Author, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IntendedJurisdiction(numIntendedJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIntendedJurisdiction >= len(resource.IntendedJurisdiction) {
		return CodeableConceptSelect("intendedJurisdiction["+strconv.Itoa(numIntendedJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("intendedJurisdiction["+strconv.Itoa(numIntendedJurisdiction)+"]", &resource.IntendedJurisdiction[numIntendedJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_Name(numName int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return StringInput("name["+strconv.Itoa(numName)+"]", nil, htmlAttrs)
	}
	return StringInput("name["+strconv.Itoa(numName)+"]", &resource.Name[numName], htmlAttrs)
}
func (resource *MedicationKnowledge) T_AssociatedMedication(frs []FhirResource, numAssociatedMedication int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAssociatedMedication >= len(resource.AssociatedMedication) {
		return ReferenceInput(frs, "associatedMedication["+strconv.Itoa(numAssociatedMedication)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "associatedMedication["+strconv.Itoa(numAssociatedMedication)+"]", &resource.AssociatedMedication[numAssociatedMedication], htmlAttrs)
}
func (resource *MedicationKnowledge) T_ProductType(numProductType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProductType >= len(resource.ProductType) {
		return CodeableConceptSelect("productType["+strconv.Itoa(numProductType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("productType["+strconv.Itoa(numProductType)+"]", &resource.ProductType[numProductType], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_PreparationInstruction(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("preparationInstruction", nil, htmlAttrs)
	}
	return StringInput("preparationInstruction", resource.PreparationInstruction, htmlAttrs)
}
func (resource *MedicationKnowledge) T_ClinicalUseIssue(frs []FhirResource, numClinicalUseIssue int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClinicalUseIssue >= len(resource.ClinicalUseIssue) {
		return ReferenceInput(frs, "clinicalUseIssue["+strconv.Itoa(numClinicalUseIssue)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "clinicalUseIssue["+strconv.Itoa(numClinicalUseIssue)+"]", &resource.ClinicalUseIssue[numClinicalUseIssue], htmlAttrs)
}
func (resource *MedicationKnowledge) T_RelatedMedicationKnowledgeType(numRelatedMedicationKnowledge int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedMedicationKnowledge >= len(resource.RelatedMedicationKnowledge) {
		return CodeableConceptSelect("relatedMedicationKnowledge["+strconv.Itoa(numRelatedMedicationKnowledge)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("relatedMedicationKnowledge["+strconv.Itoa(numRelatedMedicationKnowledge)+"].type", &resource.RelatedMedicationKnowledge[numRelatedMedicationKnowledge].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_RelatedMedicationKnowledgeReference(frs []FhirResource, numRelatedMedicationKnowledge int, numReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedMedicationKnowledge >= len(resource.RelatedMedicationKnowledge) || numReference >= len(resource.RelatedMedicationKnowledge[numRelatedMedicationKnowledge].Reference) {
		return ReferenceInput(frs, "relatedMedicationKnowledge["+strconv.Itoa(numRelatedMedicationKnowledge)+"].reference["+strconv.Itoa(numReference)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "relatedMedicationKnowledge["+strconv.Itoa(numRelatedMedicationKnowledge)+"].reference["+strconv.Itoa(numReference)+"]", &resource.RelatedMedicationKnowledge[numRelatedMedicationKnowledge].Reference[numReference], htmlAttrs)
}
func (resource *MedicationKnowledge) T_MonographType(numMonograph int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMonograph >= len(resource.Monograph) {
		return CodeableConceptSelect("monograph["+strconv.Itoa(numMonograph)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("monograph["+strconv.Itoa(numMonograph)+"].type", resource.Monograph[numMonograph].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MonographSource(frs []FhirResource, numMonograph int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMonograph >= len(resource.Monograph) {
		return ReferenceInput(frs, "monograph["+strconv.Itoa(numMonograph)+"].source", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "monograph["+strconv.Itoa(numMonograph)+"].source", resource.Monograph[numMonograph].Source, htmlAttrs)
}
func (resource *MedicationKnowledge) T_CostEffectiveDate(numCost int, numEffectiveDate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCost >= len(resource.Cost) || numEffectiveDate >= len(resource.Cost[numCost].EffectiveDate) {
		return PeriodInput("cost["+strconv.Itoa(numCost)+"].effectiveDate["+strconv.Itoa(numEffectiveDate)+"]", nil, htmlAttrs)
	}
	return PeriodInput("cost["+strconv.Itoa(numCost)+"].effectiveDate["+strconv.Itoa(numEffectiveDate)+"]", &resource.Cost[numCost].EffectiveDate[numEffectiveDate], htmlAttrs)
}
func (resource *MedicationKnowledge) T_CostType(numCost int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCost >= len(resource.Cost) {
		return CodeableConceptSelect("cost["+strconv.Itoa(numCost)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("cost["+strconv.Itoa(numCost)+"].type", &resource.Cost[numCost].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_CostSource(numCost int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCost >= len(resource.Cost) {
		return StringInput("cost["+strconv.Itoa(numCost)+"].source", nil, htmlAttrs)
	}
	return StringInput("cost["+strconv.Itoa(numCost)+"].source", resource.Cost[numCost].Source, htmlAttrs)
}
func (resource *MedicationKnowledge) T_CostCostMoney(numCost int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCost >= len(resource.Cost) {
		return MoneyInput("cost["+strconv.Itoa(numCost)+"].costMoney", nil, htmlAttrs)
	}
	return MoneyInput("cost["+strconv.Itoa(numCost)+"].costMoney", &resource.Cost[numCost].CostMoney, htmlAttrs)
}
func (resource *MedicationKnowledge) T_CostCostCodeableConcept(numCost int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCost >= len(resource.Cost) {
		return CodeableConceptSelect("cost["+strconv.Itoa(numCost)+"].costCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("cost["+strconv.Itoa(numCost)+"].costCodeableConcept", &resource.Cost[numCost].CostCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MonitoringProgramType(numMonitoringProgram int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMonitoringProgram >= len(resource.MonitoringProgram) {
		return CodeableConceptSelect("monitoringProgram["+strconv.Itoa(numMonitoringProgram)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("monitoringProgram["+strconv.Itoa(numMonitoringProgram)+"].type", resource.MonitoringProgram[numMonitoringProgram].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MonitoringProgramName(numMonitoringProgram int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMonitoringProgram >= len(resource.MonitoringProgram) {
		return StringInput("monitoringProgram["+strconv.Itoa(numMonitoringProgram)+"].name", nil, htmlAttrs)
	}
	return StringInput("monitoringProgram["+strconv.Itoa(numMonitoringProgram)+"].name", resource.MonitoringProgram[numMonitoringProgram].Name, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineIndication(numIndicationGuideline int, numIndication int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIndicationGuideline >= len(resource.IndicationGuideline) || numIndication >= len(resource.IndicationGuideline[numIndicationGuideline].Indication) {
		return CodeableReferenceInput("indicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].indication["+strconv.Itoa(numIndication)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("indicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].indication["+strconv.Itoa(numIndication)+"]", &resource.IndicationGuideline[numIndicationGuideline].Indication[numIndication], htmlAttrs)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelineTreatmentIntent(numIndicationGuideline int, numDosingGuideline int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIndicationGuideline >= len(resource.IndicationGuideline) || numDosingGuideline >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) {
		return CodeableConceptSelect("indicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].dosingGuideline["+strconv.Itoa(numDosingGuideline)+"].treatmentIntent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("indicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].dosingGuideline["+strconv.Itoa(numDosingGuideline)+"].treatmentIntent", resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].TreatmentIntent, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelineAdministrationTreatment(numIndicationGuideline int, numDosingGuideline int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIndicationGuideline >= len(resource.IndicationGuideline) || numDosingGuideline >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) {
		return CodeableConceptSelect("indicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].dosingGuideline["+strconv.Itoa(numDosingGuideline)+"].administrationTreatment", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("indicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].dosingGuideline["+strconv.Itoa(numDosingGuideline)+"].administrationTreatment", resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].AdministrationTreatment, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelineDosageType(numIndicationGuideline int, numDosingGuideline int, numDosage int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIndicationGuideline >= len(resource.IndicationGuideline) || numDosingGuideline >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) || numDosage >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].Dosage) {
		return CodeableConceptSelect("indicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].dosingGuideline["+strconv.Itoa(numDosingGuideline)+"].dosage["+strconv.Itoa(numDosage)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("indicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].dosingGuideline["+strconv.Itoa(numDosingGuideline)+"].dosage["+strconv.Itoa(numDosage)+"].type", &resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].Dosage[numDosage].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelinePatientCharacteristicType(numIndicationGuideline int, numDosingGuideline int, numPatientCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIndicationGuideline >= len(resource.IndicationGuideline) || numDosingGuideline >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) || numPatientCharacteristic >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].PatientCharacteristic) {
		return CodeableConceptSelect("indicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].dosingGuideline["+strconv.Itoa(numDosingGuideline)+"].patientCharacteristic["+strconv.Itoa(numPatientCharacteristic)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("indicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].dosingGuideline["+strconv.Itoa(numDosingGuideline)+"].patientCharacteristic["+strconv.Itoa(numPatientCharacteristic)+"].type", &resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].PatientCharacteristic[numPatientCharacteristic].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelinePatientCharacteristicValueCodeableConcept(numIndicationGuideline int, numDosingGuideline int, numPatientCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIndicationGuideline >= len(resource.IndicationGuideline) || numDosingGuideline >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) || numPatientCharacteristic >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].PatientCharacteristic) {
		return CodeableConceptSelect("indicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].dosingGuideline["+strconv.Itoa(numDosingGuideline)+"].patientCharacteristic["+strconv.Itoa(numPatientCharacteristic)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("indicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].dosingGuideline["+strconv.Itoa(numDosingGuideline)+"].patientCharacteristic["+strconv.Itoa(numPatientCharacteristic)+"].valueCodeableConcept", resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].PatientCharacteristic[numPatientCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelinePatientCharacteristicValueQuantity(numIndicationGuideline int, numDosingGuideline int, numPatientCharacteristic int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numIndicationGuideline >= len(resource.IndicationGuideline) || numDosingGuideline >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) || numPatientCharacteristic >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].PatientCharacteristic) {
		return QuantityInput("indicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].dosingGuideline["+strconv.Itoa(numDosingGuideline)+"].patientCharacteristic["+strconv.Itoa(numPatientCharacteristic)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("indicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].dosingGuideline["+strconv.Itoa(numDosingGuideline)+"].patientCharacteristic["+strconv.Itoa(numPatientCharacteristic)+"].valueQuantity", resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].PatientCharacteristic[numPatientCharacteristic].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IndicationGuidelineDosingGuidelinePatientCharacteristicValueRange(numIndicationGuideline int, numDosingGuideline int, numPatientCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIndicationGuideline >= len(resource.IndicationGuideline) || numDosingGuideline >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline) || numPatientCharacteristic >= len(resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].PatientCharacteristic) {
		return RangeInput("indicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].dosingGuideline["+strconv.Itoa(numDosingGuideline)+"].patientCharacteristic["+strconv.Itoa(numPatientCharacteristic)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("indicationGuideline["+strconv.Itoa(numIndicationGuideline)+"].dosingGuideline["+strconv.Itoa(numDosingGuideline)+"].patientCharacteristic["+strconv.Itoa(numPatientCharacteristic)+"].valueRange", resource.IndicationGuideline[numIndicationGuideline].DosingGuideline[numDosingGuideline].PatientCharacteristic[numPatientCharacteristic].ValueRange, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MedicineClassificationType(numMedicineClassification int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMedicineClassification >= len(resource.MedicineClassification) {
		return CodeableConceptSelect("medicineClassification["+strconv.Itoa(numMedicineClassification)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("medicineClassification["+strconv.Itoa(numMedicineClassification)+"].type", &resource.MedicineClassification[numMedicineClassification].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MedicineClassificationSourceString(numMedicineClassification int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMedicineClassification >= len(resource.MedicineClassification) {
		return StringInput("medicineClassification["+strconv.Itoa(numMedicineClassification)+"].sourceString", nil, htmlAttrs)
	}
	return StringInput("medicineClassification["+strconv.Itoa(numMedicineClassification)+"].sourceString", resource.MedicineClassification[numMedicineClassification].SourceString, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MedicineClassificationSourceUri(numMedicineClassification int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMedicineClassification >= len(resource.MedicineClassification) {
		return StringInput("medicineClassification["+strconv.Itoa(numMedicineClassification)+"].sourceUri", nil, htmlAttrs)
	}
	return StringInput("medicineClassification["+strconv.Itoa(numMedicineClassification)+"].sourceUri", resource.MedicineClassification[numMedicineClassification].SourceUri, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MedicineClassificationClassification(numMedicineClassification int, numClassification int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMedicineClassification >= len(resource.MedicineClassification) || numClassification >= len(resource.MedicineClassification[numMedicineClassification].Classification) {
		return CodeableConceptSelect("medicineClassification["+strconv.Itoa(numMedicineClassification)+"].classification["+strconv.Itoa(numClassification)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("medicineClassification["+strconv.Itoa(numMedicineClassification)+"].classification["+strconv.Itoa(numClassification)+"]", &resource.MedicineClassification[numMedicineClassification].Classification[numClassification], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_PackagingPackagedProduct(frs []FhirResource, numPackaging int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPackaging >= len(resource.Packaging) {
		return ReferenceInput(frs, "packaging["+strconv.Itoa(numPackaging)+"].packagedProduct", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "packaging["+strconv.Itoa(numPackaging)+"].packagedProduct", resource.Packaging[numPackaging].PackagedProduct, htmlAttrs)
}
func (resource *MedicationKnowledge) T_StorageGuidelineReference(numStorageGuideline int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStorageGuideline >= len(resource.StorageGuideline) {
		return StringInput("storageGuideline["+strconv.Itoa(numStorageGuideline)+"].reference", nil, htmlAttrs)
	}
	return StringInput("storageGuideline["+strconv.Itoa(numStorageGuideline)+"].reference", resource.StorageGuideline[numStorageGuideline].Reference, htmlAttrs)
}
func (resource *MedicationKnowledge) T_StorageGuidelineNote(numStorageGuideline int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStorageGuideline >= len(resource.StorageGuideline) || numNote >= len(resource.StorageGuideline[numStorageGuideline].Note) {
		return AnnotationTextArea("storageGuideline["+strconv.Itoa(numStorageGuideline)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("storageGuideline["+strconv.Itoa(numStorageGuideline)+"].note["+strconv.Itoa(numNote)+"]", &resource.StorageGuideline[numStorageGuideline].Note[numNote], htmlAttrs)
}
func (resource *MedicationKnowledge) T_StorageGuidelineStabilityDuration(numStorageGuideline int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStorageGuideline >= len(resource.StorageGuideline) {
		return DurationInput("storageGuideline["+strconv.Itoa(numStorageGuideline)+"].stabilityDuration", nil, htmlAttrs)
	}
	return DurationInput("storageGuideline["+strconv.Itoa(numStorageGuideline)+"].stabilityDuration", resource.StorageGuideline[numStorageGuideline].StabilityDuration, htmlAttrs)
}
func (resource *MedicationKnowledge) T_StorageGuidelineEnvironmentalSettingType(numStorageGuideline int, numEnvironmentalSetting int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStorageGuideline >= len(resource.StorageGuideline) || numEnvironmentalSetting >= len(resource.StorageGuideline[numStorageGuideline].EnvironmentalSetting) {
		return CodeableConceptSelect("storageGuideline["+strconv.Itoa(numStorageGuideline)+"].environmentalSetting["+strconv.Itoa(numEnvironmentalSetting)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("storageGuideline["+strconv.Itoa(numStorageGuideline)+"].environmentalSetting["+strconv.Itoa(numEnvironmentalSetting)+"].type", &resource.StorageGuideline[numStorageGuideline].EnvironmentalSetting[numEnvironmentalSetting].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_StorageGuidelineEnvironmentalSettingValueQuantity(numStorageGuideline int, numEnvironmentalSetting int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numStorageGuideline >= len(resource.StorageGuideline) || numEnvironmentalSetting >= len(resource.StorageGuideline[numStorageGuideline].EnvironmentalSetting) {
		return QuantityInput("storageGuideline["+strconv.Itoa(numStorageGuideline)+"].environmentalSetting["+strconv.Itoa(numEnvironmentalSetting)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("storageGuideline["+strconv.Itoa(numStorageGuideline)+"].environmentalSetting["+strconv.Itoa(numEnvironmentalSetting)+"].valueQuantity", &resource.StorageGuideline[numStorageGuideline].EnvironmentalSetting[numEnvironmentalSetting].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_StorageGuidelineEnvironmentalSettingValueRange(numStorageGuideline int, numEnvironmentalSetting int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStorageGuideline >= len(resource.StorageGuideline) || numEnvironmentalSetting >= len(resource.StorageGuideline[numStorageGuideline].EnvironmentalSetting) {
		return RangeInput("storageGuideline["+strconv.Itoa(numStorageGuideline)+"].environmentalSetting["+strconv.Itoa(numEnvironmentalSetting)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("storageGuideline["+strconv.Itoa(numStorageGuideline)+"].environmentalSetting["+strconv.Itoa(numEnvironmentalSetting)+"].valueRange", &resource.StorageGuideline[numStorageGuideline].EnvironmentalSetting[numEnvironmentalSetting].ValueRange, htmlAttrs)
}
func (resource *MedicationKnowledge) T_StorageGuidelineEnvironmentalSettingValueCodeableConcept(numStorageGuideline int, numEnvironmentalSetting int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStorageGuideline >= len(resource.StorageGuideline) || numEnvironmentalSetting >= len(resource.StorageGuideline[numStorageGuideline].EnvironmentalSetting) {
		return CodeableConceptSelect("storageGuideline["+strconv.Itoa(numStorageGuideline)+"].environmentalSetting["+strconv.Itoa(numEnvironmentalSetting)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("storageGuideline["+strconv.Itoa(numStorageGuideline)+"].environmentalSetting["+strconv.Itoa(numEnvironmentalSetting)+"].valueCodeableConcept", &resource.StorageGuideline[numStorageGuideline].EnvironmentalSetting[numEnvironmentalSetting].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_RegulatoryRegulatoryAuthority(frs []FhirResource, numRegulatory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRegulatory >= len(resource.Regulatory) {
		return ReferenceInput(frs, "regulatory["+strconv.Itoa(numRegulatory)+"].regulatoryAuthority", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "regulatory["+strconv.Itoa(numRegulatory)+"].regulatoryAuthority", &resource.Regulatory[numRegulatory].RegulatoryAuthority, htmlAttrs)
}
func (resource *MedicationKnowledge) T_RegulatorySchedule(numRegulatory int, numSchedule int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRegulatory >= len(resource.Regulatory) || numSchedule >= len(resource.Regulatory[numRegulatory].Schedule) {
		return CodeableConceptSelect("regulatory["+strconv.Itoa(numRegulatory)+"].schedule["+strconv.Itoa(numSchedule)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("regulatory["+strconv.Itoa(numRegulatory)+"].schedule["+strconv.Itoa(numSchedule)+"]", &resource.Regulatory[numRegulatory].Schedule[numSchedule], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_RegulatorySubstitutionType(numRegulatory int, numSubstitution int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRegulatory >= len(resource.Regulatory) || numSubstitution >= len(resource.Regulatory[numRegulatory].Substitution) {
		return CodeableConceptSelect("regulatory["+strconv.Itoa(numRegulatory)+"].substitution["+strconv.Itoa(numSubstitution)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("regulatory["+strconv.Itoa(numRegulatory)+"].substitution["+strconv.Itoa(numSubstitution)+"].type", &resource.Regulatory[numRegulatory].Substitution[numSubstitution].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_RegulatorySubstitutionAllowed(numRegulatory int, numSubstitution int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRegulatory >= len(resource.Regulatory) || numSubstitution >= len(resource.Regulatory[numRegulatory].Substitution) {
		return BoolInput("regulatory["+strconv.Itoa(numRegulatory)+"].substitution["+strconv.Itoa(numSubstitution)+"].allowed", nil, htmlAttrs)
	}
	return BoolInput("regulatory["+strconv.Itoa(numRegulatory)+"].substitution["+strconv.Itoa(numSubstitution)+"].allowed", &resource.Regulatory[numRegulatory].Substitution[numSubstitution].Allowed, htmlAttrs)
}
func (resource *MedicationKnowledge) T_RegulatoryMaxDispenseQuantity(numRegulatory int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numRegulatory >= len(resource.Regulatory) {
		return QuantityInput("regulatory["+strconv.Itoa(numRegulatory)+"].maxDispense.quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("regulatory["+strconv.Itoa(numRegulatory)+"].maxDispense.quantity", &resource.Regulatory[numRegulatory].MaxDispense.Quantity, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_RegulatoryMaxDispensePeriod(numRegulatory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRegulatory >= len(resource.Regulatory) {
		return DurationInput("regulatory["+strconv.Itoa(numRegulatory)+"].maxDispense.period", nil, htmlAttrs)
	}
	return DurationInput("regulatory["+strconv.Itoa(numRegulatory)+"].maxDispense.period", resource.Regulatory[numRegulatory].MaxDispense.Period, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalDefinition(frs []FhirResource, numDefinition int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDefinition >= len(resource.Definitional.Definition) {
		return ReferenceInput(frs, "definitional.definition["+strconv.Itoa(numDefinition)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "definitional.definition["+strconv.Itoa(numDefinition)+"]", &resource.Definitional.Definition[numDefinition], htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalDoseForm(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("definitional.doseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("definitional.doseForm", resource.Definitional.DoseForm, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalIntendedRoute(numIntendedRoute int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIntendedRoute >= len(resource.Definitional.IntendedRoute) {
		return CodeableConceptSelect("definitional.intendedRoute["+strconv.Itoa(numIntendedRoute)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("definitional.intendedRoute["+strconv.Itoa(numIntendedRoute)+"]", &resource.Definitional.IntendedRoute[numIntendedRoute], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalIngredientItem(numIngredient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Definitional.Ingredient) {
		return CodeableReferenceInput("definitional.ingredient["+strconv.Itoa(numIngredient)+"].item", nil, htmlAttrs)
	}
	return CodeableReferenceInput("definitional.ingredient["+strconv.Itoa(numIngredient)+"].item", &resource.Definitional.Ingredient[numIngredient].Item, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalIngredientType(numIngredient int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Definitional.Ingredient) {
		return CodeableConceptSelect("definitional.ingredient["+strconv.Itoa(numIngredient)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("definitional.ingredient["+strconv.Itoa(numIngredient)+"].type", resource.Definitional.Ingredient[numIngredient].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalIngredientStrengthRatio(numIngredient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Definitional.Ingredient) {
		return RatioInput("definitional.ingredient["+strconv.Itoa(numIngredient)+"].strengthRatio", nil, htmlAttrs)
	}
	return RatioInput("definitional.ingredient["+strconv.Itoa(numIngredient)+"].strengthRatio", resource.Definitional.Ingredient[numIngredient].StrengthRatio, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalIngredientStrengthCodeableConcept(numIngredient int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Definitional.Ingredient) {
		return CodeableConceptSelect("definitional.ingredient["+strconv.Itoa(numIngredient)+"].strengthCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("definitional.ingredient["+strconv.Itoa(numIngredient)+"].strengthCodeableConcept", resource.Definitional.Ingredient[numIngredient].StrengthCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalIngredientStrengthQuantity(numIngredient int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numIngredient >= len(resource.Definitional.Ingredient) {
		return QuantityInput("definitional.ingredient["+strconv.Itoa(numIngredient)+"].strengthQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("definitional.ingredient["+strconv.Itoa(numIngredient)+"].strengthQuantity", resource.Definitional.Ingredient[numIngredient].StrengthQuantity, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalDrugCharacteristicType(numDrugCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDrugCharacteristic >= len(resource.Definitional.DrugCharacteristic) {
		return CodeableConceptSelect("definitional.drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("definitional.drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].type", resource.Definitional.DrugCharacteristic[numDrugCharacteristic].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalDrugCharacteristicValueCodeableConcept(numDrugCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDrugCharacteristic >= len(resource.Definitional.DrugCharacteristic) {
		return CodeableConceptSelect("definitional.drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("definitional.drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueCodeableConcept", resource.Definitional.DrugCharacteristic[numDrugCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalDrugCharacteristicValueString(numDrugCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDrugCharacteristic >= len(resource.Definitional.DrugCharacteristic) {
		return StringInput("definitional.drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("definitional.drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueString", resource.Definitional.DrugCharacteristic[numDrugCharacteristic].ValueString, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalDrugCharacteristicValueQuantity(numDrugCharacteristic int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numDrugCharacteristic >= len(resource.Definitional.DrugCharacteristic) {
		return QuantityInput("definitional.drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("definitional.drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueQuantity", resource.Definitional.DrugCharacteristic[numDrugCharacteristic].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalDrugCharacteristicValueBase64Binary(numDrugCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDrugCharacteristic >= len(resource.Definitional.DrugCharacteristic) {
		return StringInput("definitional.drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("definitional.drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueBase64Binary", resource.Definitional.DrugCharacteristic[numDrugCharacteristic].ValueBase64Binary, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DefinitionalDrugCharacteristicValueAttachment(numDrugCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDrugCharacteristic >= len(resource.Definitional.DrugCharacteristic) {
		return AttachmentInput("definitional.drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("definitional.drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueAttachment", resource.Definitional.DrugCharacteristic[numDrugCharacteristic].ValueAttachment, htmlAttrs)
}
