//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

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
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	ValueQuantity        *Quantity        `json:"valueQuantity"`
	ValueRange           *Range           `json:"valueRange"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeMedicineClassification struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	SourceString      *string           `json:"sourceString"`
	SourceUri         *string           `json:"sourceUri"`
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
	StrengthRatio           *Ratio            `json:"strengthRatio"`
	StrengthCodeableConcept *CodeableConcept  `json:"strengthCodeableConcept"`
	StrengthQuantity        *Quantity         `json:"strengthQuantity"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeDefinitionalDrugCharacteristic struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Type                 *CodeableConcept `json:"type,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	ValueString          *string          `json:"valueString"`
	ValueQuantity        *Quantity        `json:"valueQuantity"`
	ValueBase64Binary    *string          `json:"valueBase64Binary"`
	ValueAttachment      *Attachment      `json:"valueAttachment"`
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
