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
func (resource *MedicationKnowledge) T_Manufacturer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "manufacturer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "manufacturer", resource.Manufacturer, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DoseForm(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("doseForm", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("doseForm", resource.DoseForm, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_Amount(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("amount", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("amount", resource.Amount, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_Synonym(numSynonym int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSynonym >= len(resource.Synonym) {
		return StringInput("synonym["+strconv.Itoa(numSynonym)+"]", nil, htmlAttrs)
	}
	return StringInput("synonym["+strconv.Itoa(numSynonym)+"]", &resource.Synonym[numSynonym], htmlAttrs)
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
func (resource *MedicationKnowledge) T_IntendedRoute(numIntendedRoute int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIntendedRoute >= len(resource.IntendedRoute) {
		return CodeableConceptSelect("intendedRoute["+strconv.Itoa(numIntendedRoute)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("intendedRoute["+strconv.Itoa(numIntendedRoute)+"]", &resource.IntendedRoute[numIntendedRoute], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_Contraindication(frs []FhirResource, numContraindication int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContraindication >= len(resource.Contraindication) {
		return ReferenceInput(frs, "contraindication["+strconv.Itoa(numContraindication)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "contraindication["+strconv.Itoa(numContraindication)+"]", &resource.Contraindication[numContraindication], htmlAttrs)
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
func (resource *MedicationKnowledge) T_IngredientItemCodeableConcept(numIngredient int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return CodeableConceptSelect("ingredient["+strconv.Itoa(numIngredient)+"].itemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ingredient["+strconv.Itoa(numIngredient)+"].itemCodeableConcept", &resource.Ingredient[numIngredient].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IngredientItemReference(frs []FhirResource, numIngredient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return ReferenceInput(frs, "ingredient["+strconv.Itoa(numIngredient)+"].itemReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "ingredient["+strconv.Itoa(numIngredient)+"].itemReference", &resource.Ingredient[numIngredient].ItemReference, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IngredientIsActive(numIngredient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return BoolInput("ingredient["+strconv.Itoa(numIngredient)+"].isActive", nil, htmlAttrs)
	}
	return BoolInput("ingredient["+strconv.Itoa(numIngredient)+"].isActive", resource.Ingredient[numIngredient].IsActive, htmlAttrs)
}
func (resource *MedicationKnowledge) T_IngredientStrength(numIngredient int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIngredient >= len(resource.Ingredient) {
		return RatioInput("ingredient["+strconv.Itoa(numIngredient)+"].strength", nil, htmlAttrs)
	}
	return RatioInput("ingredient["+strconv.Itoa(numIngredient)+"].strength", resource.Ingredient[numIngredient].Strength, htmlAttrs)
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
func (resource *MedicationKnowledge) T_CostCost(numCost int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCost >= len(resource.Cost) {
		return MoneyInput("cost["+strconv.Itoa(numCost)+"].cost", nil, htmlAttrs)
	}
	return MoneyInput("cost["+strconv.Itoa(numCost)+"].cost", &resource.Cost[numCost].Cost, htmlAttrs)
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
func (resource *MedicationKnowledge) T_AdministrationGuidelinesIndicationCodeableConcept(numAdministrationGuidelines int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdministrationGuidelines >= len(resource.AdministrationGuidelines) {
		return CodeableConceptSelect("administrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].indicationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("administrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].indicationCodeableConcept", resource.AdministrationGuidelines[numAdministrationGuidelines].IndicationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_AdministrationGuidelinesIndicationReference(frs []FhirResource, numAdministrationGuidelines int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdministrationGuidelines >= len(resource.AdministrationGuidelines) {
		return ReferenceInput(frs, "administrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].indicationReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "administrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].indicationReference", resource.AdministrationGuidelines[numAdministrationGuidelines].IndicationReference, htmlAttrs)
}
func (resource *MedicationKnowledge) T_AdministrationGuidelinesDosageType(numAdministrationGuidelines int, numDosage int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdministrationGuidelines >= len(resource.AdministrationGuidelines) || numDosage >= len(resource.AdministrationGuidelines[numAdministrationGuidelines].Dosage) {
		return CodeableConceptSelect("administrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].dosage["+strconv.Itoa(numDosage)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("administrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].dosage["+strconv.Itoa(numDosage)+"].type", &resource.AdministrationGuidelines[numAdministrationGuidelines].Dosage[numDosage].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_AdministrationGuidelinesPatientCharacteristicsCharacteristicCodeableConcept(numAdministrationGuidelines int, numPatientCharacteristics int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdministrationGuidelines >= len(resource.AdministrationGuidelines) || numPatientCharacteristics >= len(resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics) {
		return CodeableConceptSelect("administrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].patientCharacteristics["+strconv.Itoa(numPatientCharacteristics)+"].characteristicCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("administrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].patientCharacteristics["+strconv.Itoa(numPatientCharacteristics)+"].characteristicCodeableConcept", &resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics[numPatientCharacteristics].CharacteristicCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_AdministrationGuidelinesPatientCharacteristicsCharacteristicQuantity(numAdministrationGuidelines int, numPatientCharacteristics int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numAdministrationGuidelines >= len(resource.AdministrationGuidelines) || numPatientCharacteristics >= len(resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics) {
		return QuantityInput("administrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].patientCharacteristics["+strconv.Itoa(numPatientCharacteristics)+"].characteristicQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("administrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].patientCharacteristics["+strconv.Itoa(numPatientCharacteristics)+"].characteristicQuantity", &resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics[numPatientCharacteristics].CharacteristicQuantity, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_AdministrationGuidelinesPatientCharacteristicsValue(numAdministrationGuidelines int, numPatientCharacteristics int, numValue int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAdministrationGuidelines >= len(resource.AdministrationGuidelines) || numPatientCharacteristics >= len(resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics) || numValue >= len(resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics[numPatientCharacteristics].Value) {
		return StringInput("administrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].patientCharacteristics["+strconv.Itoa(numPatientCharacteristics)+"].value["+strconv.Itoa(numValue)+"]", nil, htmlAttrs)
	}
	return StringInput("administrationGuidelines["+strconv.Itoa(numAdministrationGuidelines)+"].patientCharacteristics["+strconv.Itoa(numPatientCharacteristics)+"].value["+strconv.Itoa(numValue)+"]", &resource.AdministrationGuidelines[numAdministrationGuidelines].PatientCharacteristics[numPatientCharacteristics].Value[numValue], htmlAttrs)
}
func (resource *MedicationKnowledge) T_MedicineClassificationType(numMedicineClassification int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMedicineClassification >= len(resource.MedicineClassification) {
		return CodeableConceptSelect("medicineClassification["+strconv.Itoa(numMedicineClassification)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("medicineClassification["+strconv.Itoa(numMedicineClassification)+"].type", &resource.MedicineClassification[numMedicineClassification].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_MedicineClassificationClassification(numMedicineClassification int, numClassification int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMedicineClassification >= len(resource.MedicineClassification) || numClassification >= len(resource.MedicineClassification[numMedicineClassification].Classification) {
		return CodeableConceptSelect("medicineClassification["+strconv.Itoa(numMedicineClassification)+"].classification["+strconv.Itoa(numClassification)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("medicineClassification["+strconv.Itoa(numMedicineClassification)+"].classification["+strconv.Itoa(numClassification)+"]", &resource.MedicineClassification[numMedicineClassification].Classification[numClassification], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_PackagingType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("packaging.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("packaging.type", resource.Packaging.Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_PackagingQuantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("packaging.quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("packaging.quantity", resource.Packaging.Quantity, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DrugCharacteristicType(numDrugCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDrugCharacteristic >= len(resource.DrugCharacteristic) {
		return CodeableConceptSelect("drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].type", resource.DrugCharacteristic[numDrugCharacteristic].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DrugCharacteristicValueCodeableConcept(numDrugCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDrugCharacteristic >= len(resource.DrugCharacteristic) {
		return CodeableConceptSelect("drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueCodeableConcept", resource.DrugCharacteristic[numDrugCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DrugCharacteristicValueString(numDrugCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDrugCharacteristic >= len(resource.DrugCharacteristic) {
		return StringInput("drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueString", resource.DrugCharacteristic[numDrugCharacteristic].ValueString, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DrugCharacteristicValueQuantity(numDrugCharacteristic int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numDrugCharacteristic >= len(resource.DrugCharacteristic) {
		return QuantityInput("drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueQuantity", resource.DrugCharacteristic[numDrugCharacteristic].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_DrugCharacteristicValueBase64Binary(numDrugCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDrugCharacteristic >= len(resource.DrugCharacteristic) {
		return StringInput("drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueBase64Binary", nil, htmlAttrs)
	}
	return StringInput("drugCharacteristic["+strconv.Itoa(numDrugCharacteristic)+"].valueBase64Binary", resource.DrugCharacteristic[numDrugCharacteristic].ValueBase64Binary, htmlAttrs)
}
func (resource *MedicationKnowledge) T_RegulatoryRegulatoryAuthority(frs []FhirResource, numRegulatory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRegulatory >= len(resource.Regulatory) {
		return ReferenceInput(frs, "regulatory["+strconv.Itoa(numRegulatory)+"].regulatoryAuthority", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "regulatory["+strconv.Itoa(numRegulatory)+"].regulatoryAuthority", &resource.Regulatory[numRegulatory].RegulatoryAuthority, htmlAttrs)
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
func (resource *MedicationKnowledge) T_RegulatoryScheduleSchedule(numRegulatory int, numSchedule int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRegulatory >= len(resource.Regulatory) || numSchedule >= len(resource.Regulatory[numRegulatory].Schedule) {
		return CodeableConceptSelect("regulatory["+strconv.Itoa(numRegulatory)+"].schedule["+strconv.Itoa(numSchedule)+"].schedule", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("regulatory["+strconv.Itoa(numRegulatory)+"].schedule["+strconv.Itoa(numSchedule)+"].schedule", &resource.Regulatory[numRegulatory].Schedule[numSchedule].Schedule, optionsValueSet, htmlAttrs)
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
func (resource *MedicationKnowledge) T_KineticsAreaUnderCurve(numKinetics int, numAreaUnderCurve int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numKinetics >= len(resource.Kinetics) || numAreaUnderCurve >= len(resource.Kinetics[numKinetics].AreaUnderCurve) {
		return QuantityInput("kinetics["+strconv.Itoa(numKinetics)+"].areaUnderCurve["+strconv.Itoa(numAreaUnderCurve)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("kinetics["+strconv.Itoa(numKinetics)+"].areaUnderCurve["+strconv.Itoa(numAreaUnderCurve)+"]", &resource.Kinetics[numKinetics].AreaUnderCurve[numAreaUnderCurve], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_KineticsLethalDose50(numKinetics int, numLethalDose50 int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numKinetics >= len(resource.Kinetics) || numLethalDose50 >= len(resource.Kinetics[numKinetics].LethalDose50) {
		return QuantityInput("kinetics["+strconv.Itoa(numKinetics)+"].lethalDose50["+strconv.Itoa(numLethalDose50)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("kinetics["+strconv.Itoa(numKinetics)+"].lethalDose50["+strconv.Itoa(numLethalDose50)+"]", &resource.Kinetics[numKinetics].LethalDose50[numLethalDose50], optionsValueSet, htmlAttrs)
}
func (resource *MedicationKnowledge) T_KineticsHalfLifePeriod(numKinetics int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numKinetics >= len(resource.Kinetics) {
		return DurationInput("kinetics["+strconv.Itoa(numKinetics)+"].halfLifePeriod", nil, htmlAttrs)
	}
	return DurationInput("kinetics["+strconv.Itoa(numKinetics)+"].halfLifePeriod", resource.Kinetics[numKinetics].HalfLifePeriod, htmlAttrs)
}
