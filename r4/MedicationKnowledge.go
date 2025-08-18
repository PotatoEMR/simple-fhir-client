//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

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
	IndicationCodeableConcept *CodeableConcept                                                    `json:"indicationCodeableConcept"`
	IndicationReference       *Reference                                                          `json:"indicationReference"`
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
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	ValueString          *string          `json:"valueString"`
	ValueQuantity        *Quantity        `json:"valueQuantity"`
	ValueBase64Binary    *string          `json:"valueBase64Binary"`
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
