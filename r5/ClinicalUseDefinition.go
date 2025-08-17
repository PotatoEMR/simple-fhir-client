//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/ClinicalUseDefinition
type ClinicalUseDefinition struct {
	Id                *string                                 `json:"id,omitempty"`
	Meta              *Meta                                   `json:"meta,omitempty"`
	ImplicitRules     *string                                 `json:"implicitRules,omitempty"`
	Language          *string                                 `json:"language,omitempty"`
	Text              *Narrative                              `json:"text,omitempty"`
	Contained         []Resource                              `json:"contained,omitempty"`
	Extension         []Extension                             `json:"extension,omitempty"`
	ModifierExtension []Extension                             `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                            `json:"identifier,omitempty"`
	Type              string                                  `json:"type"`
	Category          []CodeableConcept                       `json:"category,omitempty"`
	Subject           []Reference                             `json:"subject,omitempty"`
	Status            *CodeableConcept                        `json:"status,omitempty"`
	Contraindication  *ClinicalUseDefinitionContraindication  `json:"contraindication,omitempty"`
	Indication        *ClinicalUseDefinitionIndication        `json:"indication,omitempty"`
	Interaction       *ClinicalUseDefinitionInteraction       `json:"interaction,omitempty"`
	Population        []Reference                             `json:"population,omitempty"`
	Library           []string                                `json:"library,omitempty"`
	UndesirableEffect *ClinicalUseDefinitionUndesirableEffect `json:"undesirableEffect,omitempty"`
	Warning           *ClinicalUseDefinitionWarning           `json:"warning,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClinicalUseDefinition
type ClinicalUseDefinitionContraindication struct {
	Id                      *string                                             `json:"id,omitempty"`
	Extension               []Extension                                         `json:"extension,omitempty"`
	ModifierExtension       []Extension                                         `json:"modifierExtension,omitempty"`
	DiseaseSymptomProcedure *CodeableReference                                  `json:"diseaseSymptomProcedure,omitempty"`
	DiseaseStatus           *CodeableReference                                  `json:"diseaseStatus,omitempty"`
	Comorbidity             []CodeableReference                                 `json:"comorbidity,omitempty"`
	Indication              []Reference                                         `json:"indication,omitempty"`
	Applicability           *Expression                                         `json:"applicability,omitempty"`
	OtherTherapy            []ClinicalUseDefinitionContraindicationOtherTherapy `json:"otherTherapy,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClinicalUseDefinition
type ClinicalUseDefinitionContraindicationOtherTherapy struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	RelationshipType  CodeableConcept   `json:"relationshipType"`
	Treatment         CodeableReference `json:"treatment"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClinicalUseDefinition
type ClinicalUseDefinitionIndication struct {
	Id                      *string             `json:"id,omitempty"`
	Extension               []Extension         `json:"extension,omitempty"`
	ModifierExtension       []Extension         `json:"modifierExtension,omitempty"`
	DiseaseSymptomProcedure *CodeableReference  `json:"diseaseSymptomProcedure,omitempty"`
	DiseaseStatus           *CodeableReference  `json:"diseaseStatus,omitempty"`
	Comorbidity             []CodeableReference `json:"comorbidity,omitempty"`
	IntendedEffect          *CodeableReference  `json:"intendedEffect,omitempty"`
	DurationRange           *Range              `json:"durationRange"`
	DurationString          *string             `json:"durationString"`
	UndesirableEffect       []Reference         `json:"undesirableEffect,omitempty"`
	Applicability           *Expression         `json:"applicability,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClinicalUseDefinition
type ClinicalUseDefinitionInteraction struct {
	Id                *string                                       `json:"id,omitempty"`
	Extension         []Extension                                   `json:"extension,omitempty"`
	ModifierExtension []Extension                                   `json:"modifierExtension,omitempty"`
	Interactant       []ClinicalUseDefinitionInteractionInteractant `json:"interactant,omitempty"`
	Type              *CodeableConcept                              `json:"type,omitempty"`
	Effect            *CodeableReference                            `json:"effect,omitempty"`
	Incidence         *CodeableConcept                              `json:"incidence,omitempty"`
	Management        []CodeableConcept                             `json:"management,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClinicalUseDefinition
type ClinicalUseDefinitionInteractionInteractant struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	ItemReference       Reference       `json:"itemReference"`
	ItemCodeableConcept CodeableConcept `json:"itemCodeableConcept"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClinicalUseDefinition
type ClinicalUseDefinitionUndesirableEffect struct {
	Id                     *string            `json:"id,omitempty"`
	Extension              []Extension        `json:"extension,omitempty"`
	ModifierExtension      []Extension        `json:"modifierExtension,omitempty"`
	SymptomConditionEffect *CodeableReference `json:"symptomConditionEffect,omitempty"`
	Classification         *CodeableConcept   `json:"classification,omitempty"`
	FrequencyOfOccurrence  *CodeableConcept   `json:"frequencyOfOccurrence,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClinicalUseDefinition
type ClinicalUseDefinitionWarning struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
}

type OtherClinicalUseDefinition ClinicalUseDefinition

// on convert struct to json, automatically add resourceType=ClinicalUseDefinition
func (r ClinicalUseDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClinicalUseDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherClinicalUseDefinition: OtherClinicalUseDefinition(r),
		ResourceType:               "ClinicalUseDefinition",
	})
}
