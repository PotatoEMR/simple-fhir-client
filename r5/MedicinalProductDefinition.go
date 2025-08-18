//generated August 18 2025 with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

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
	StatusDate                     *string                                    `json:"statusDate,omitempty"`
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
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	ValueMarkdown        *string          `json:"valueMarkdown"`
	ValueQuantity        *Quantity        `json:"valueQuantity"`
	ValueInteger         *int             `json:"valueInteger"`
	ValueDate            *string          `json:"valueDate"`
	ValueBoolean         *bool            `json:"valueBoolean"`
	ValueAttachment      *Attachment      `json:"valueAttachment"`
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
