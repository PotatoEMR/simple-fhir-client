package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/ClinicalUseDefinition
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
	UndesirableEffect *ClinicalUseDefinitionUndesirableEffect `json:"undesirableEffect,omitempty"`
	Warning           *ClinicalUseDefinitionWarning           `json:"warning,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClinicalUseDefinition
type ClinicalUseDefinitionContraindication struct {
	Id                      *string                                             `json:"id,omitempty"`
	Extension               []Extension                                         `json:"extension,omitempty"`
	ModifierExtension       []Extension                                         `json:"modifierExtension,omitempty"`
	DiseaseSymptomProcedure *CodeableReference                                  `json:"diseaseSymptomProcedure,omitempty"`
	DiseaseStatus           *CodeableReference                                  `json:"diseaseStatus,omitempty"`
	Comorbidity             []CodeableReference                                 `json:"comorbidity,omitempty"`
	Indication              []Reference                                         `json:"indication,omitempty"`
	OtherTherapy            []ClinicalUseDefinitionContraindicationOtherTherapy `json:"otherTherapy,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClinicalUseDefinition
type ClinicalUseDefinitionContraindicationOtherTherapy struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	RelationshipType  CodeableConcept   `json:"relationshipType"`
	Therapy           CodeableReference `json:"therapy"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClinicalUseDefinition
type ClinicalUseDefinitionIndication struct {
	Id                      *string             `json:"id,omitempty"`
	Extension               []Extension         `json:"extension,omitempty"`
	ModifierExtension       []Extension         `json:"modifierExtension,omitempty"`
	DiseaseSymptomProcedure *CodeableReference  `json:"diseaseSymptomProcedure,omitempty"`
	DiseaseStatus           *CodeableReference  `json:"diseaseStatus,omitempty"`
	Comorbidity             []CodeableReference `json:"comorbidity,omitempty"`
	IntendedEffect          *CodeableReference  `json:"intendedEffect,omitempty"`
	DurationRange           *Range              `json:"durationRange,omitempty"`
	DurationString          *string             `json:"durationString,omitempty"`
	UndesirableEffect       []Reference         `json:"undesirableEffect,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClinicalUseDefinition
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

// http://hl7.org/fhir/r4b/StructureDefinition/ClinicalUseDefinition
type ClinicalUseDefinitionInteractionInteractant struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	ItemReference       Reference       `json:"itemReference"`
	ItemCodeableConcept CodeableConcept `json:"itemCodeableConcept"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClinicalUseDefinition
type ClinicalUseDefinitionUndesirableEffect struct {
	Id                     *string            `json:"id,omitempty"`
	Extension              []Extension        `json:"extension,omitempty"`
	ModifierExtension      []Extension        `json:"modifierExtension,omitempty"`
	SymptomConditionEffect *CodeableReference `json:"symptomConditionEffect,omitempty"`
	Classification         *CodeableConcept   `json:"classification,omitempty"`
	FrequencyOfOccurrence  *CodeableConcept   `json:"frequencyOfOccurrence,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClinicalUseDefinition
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

func (resource *ClinicalUseDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ClinicalUseDefinition.Id", nil)
	}
	return StringInput("ClinicalUseDefinition.Id", resource.Id)
}
func (resource *ClinicalUseDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ClinicalUseDefinition.ImplicitRules", nil)
	}
	return StringInput("ClinicalUseDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *ClinicalUseDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ClinicalUseDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("ClinicalUseDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *ClinicalUseDefinition) T_Type() templ.Component {
	optionsValueSet := VSClinical_use_definition_type

	if resource == nil {
		return CodeSelect("ClinicalUseDefinition.Type", nil, optionsValueSet)
	}
	return CodeSelect("ClinicalUseDefinition.Type", &resource.Type, optionsValueSet)
}
func (resource *ClinicalUseDefinition) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("ClinicalUseDefinition.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *ClinicalUseDefinition) T_Status(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ClinicalUseDefinition.Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.Status", resource.Status, optionsValueSet)
}
func (resource *ClinicalUseDefinition) T_ContraindicationId() templ.Component {

	if resource == nil {
		return StringInput("ClinicalUseDefinition.Contraindication.Id", nil)
	}
	return StringInput("ClinicalUseDefinition.Contraindication.Id", resource.Contraindication.Id)
}
func (resource *ClinicalUseDefinition) T_ContraindicationOtherTherapyId(numOtherTherapy int) templ.Component {

	if resource == nil || len(resource.Contraindication.OtherTherapy) >= numOtherTherapy {
		return StringInput("ClinicalUseDefinition.Contraindication.OtherTherapy["+strconv.Itoa(numOtherTherapy)+"].Id", nil)
	}
	return StringInput("ClinicalUseDefinition.Contraindication.OtherTherapy["+strconv.Itoa(numOtherTherapy)+"].Id", resource.Contraindication.OtherTherapy[numOtherTherapy].Id)
}
func (resource *ClinicalUseDefinition) T_ContraindicationOtherTherapyRelationshipType(numOtherTherapy int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Contraindication.OtherTherapy) >= numOtherTherapy {
		return CodeableConceptSelect("ClinicalUseDefinition.Contraindication.OtherTherapy["+strconv.Itoa(numOtherTherapy)+"].RelationshipType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.Contraindication.OtherTherapy["+strconv.Itoa(numOtherTherapy)+"].RelationshipType", &resource.Contraindication.OtherTherapy[numOtherTherapy].RelationshipType, optionsValueSet)
}
func (resource *ClinicalUseDefinition) T_IndicationId() templ.Component {

	if resource == nil {
		return StringInput("ClinicalUseDefinition.Indication.Id", nil)
	}
	return StringInput("ClinicalUseDefinition.Indication.Id", resource.Indication.Id)
}
func (resource *ClinicalUseDefinition) T_InteractionId() templ.Component {

	if resource == nil {
		return StringInput("ClinicalUseDefinition.Interaction.Id", nil)
	}
	return StringInput("ClinicalUseDefinition.Interaction.Id", resource.Interaction.Id)
}
func (resource *ClinicalUseDefinition) T_InteractionType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ClinicalUseDefinition.Interaction.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.Interaction.Type", resource.Interaction.Type, optionsValueSet)
}
func (resource *ClinicalUseDefinition) T_InteractionIncidence(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ClinicalUseDefinition.Interaction.Incidence", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.Interaction.Incidence", resource.Interaction.Incidence, optionsValueSet)
}
func (resource *ClinicalUseDefinition) T_InteractionManagement(numManagement int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Interaction.Management) >= numManagement {
		return CodeableConceptSelect("ClinicalUseDefinition.Interaction.Management["+strconv.Itoa(numManagement)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.Interaction.Management["+strconv.Itoa(numManagement)+"]", &resource.Interaction.Management[numManagement], optionsValueSet)
}
func (resource *ClinicalUseDefinition) T_InteractionInteractantId(numInteractant int) templ.Component {

	if resource == nil || len(resource.Interaction.Interactant) >= numInteractant {
		return StringInput("ClinicalUseDefinition.Interaction.Interactant["+strconv.Itoa(numInteractant)+"].Id", nil)
	}
	return StringInput("ClinicalUseDefinition.Interaction.Interactant["+strconv.Itoa(numInteractant)+"].Id", resource.Interaction.Interactant[numInteractant].Id)
}
func (resource *ClinicalUseDefinition) T_UndesirableEffectId() templ.Component {

	if resource == nil {
		return StringInput("ClinicalUseDefinition.UndesirableEffect.Id", nil)
	}
	return StringInput("ClinicalUseDefinition.UndesirableEffect.Id", resource.UndesirableEffect.Id)
}
func (resource *ClinicalUseDefinition) T_UndesirableEffectClassification(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ClinicalUseDefinition.UndesirableEffect.Classification", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.UndesirableEffect.Classification", resource.UndesirableEffect.Classification, optionsValueSet)
}
func (resource *ClinicalUseDefinition) T_UndesirableEffectFrequencyOfOccurrence(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ClinicalUseDefinition.UndesirableEffect.FrequencyOfOccurrence", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.UndesirableEffect.FrequencyOfOccurrence", resource.UndesirableEffect.FrequencyOfOccurrence, optionsValueSet)
}
func (resource *ClinicalUseDefinition) T_WarningId() templ.Component {

	if resource == nil {
		return StringInput("ClinicalUseDefinition.Warning.Id", nil)
	}
	return StringInput("ClinicalUseDefinition.Warning.Id", resource.Warning.Id)
}
func (resource *ClinicalUseDefinition) T_WarningDescription() templ.Component {

	if resource == nil {
		return StringInput("ClinicalUseDefinition.Warning.Description", nil)
	}
	return StringInput("ClinicalUseDefinition.Warning.Description", resource.Warning.Description)
}
func (resource *ClinicalUseDefinition) T_WarningCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ClinicalUseDefinition.Warning.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.Warning.Code", resource.Warning.Code, optionsValueSet)
}
