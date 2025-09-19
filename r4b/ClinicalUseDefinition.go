package r4b

//generated with command go run ./bultaoreune
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
func (r ClinicalUseDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ClinicalUseDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ClinicalUseDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ClinicalUseDefinition) T_Type(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSClinical_use_definition_type

	if resource == nil {
		return CodeSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_Subject(numSubject int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubject >= len(resource.Subject) {
		return ReferenceInput("subject["+strconv.Itoa(numSubject)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("subject["+strconv.Itoa(numSubject)+"]", &resource.Subject[numSubject], htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_Status(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_Population(numPopulation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPopulation >= len(resource.Population) {
		return ReferenceInput("population["+strconv.Itoa(numPopulation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("population["+strconv.Itoa(numPopulation)+"]", &resource.Population[numPopulation], htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_ContraindicationDiseaseSymptomProcedure(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("contraindication.diseaseSymptomProcedure", nil, htmlAttrs)
	}
	return CodeableReferenceInput("contraindication.diseaseSymptomProcedure", resource.Contraindication.DiseaseSymptomProcedure, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_ContraindicationDiseaseStatus(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("contraindication.diseaseStatus", nil, htmlAttrs)
	}
	return CodeableReferenceInput("contraindication.diseaseStatus", resource.Contraindication.DiseaseStatus, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_ContraindicationComorbidity(numComorbidity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComorbidity >= len(resource.Contraindication.Comorbidity) {
		return CodeableReferenceInput("contraindication.comorbidity["+strconv.Itoa(numComorbidity)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("contraindication.comorbidity["+strconv.Itoa(numComorbidity)+"]", &resource.Contraindication.Comorbidity[numComorbidity], htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_ContraindicationIndication(numIndication int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIndication >= len(resource.Contraindication.Indication) {
		return ReferenceInput("contraindication.indication["+strconv.Itoa(numIndication)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("contraindication.indication["+strconv.Itoa(numIndication)+"]", &resource.Contraindication.Indication[numIndication], htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_ContraindicationOtherTherapyRelationshipType(numOtherTherapy int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOtherTherapy >= len(resource.Contraindication.OtherTherapy) {
		return CodeableConceptSelect("contraindication.otherTherapy["+strconv.Itoa(numOtherTherapy)+"].relationshipType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("contraindication.otherTherapy["+strconv.Itoa(numOtherTherapy)+"].relationshipType", &resource.Contraindication.OtherTherapy[numOtherTherapy].RelationshipType, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_ContraindicationOtherTherapyTherapy(numOtherTherapy int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOtherTherapy >= len(resource.Contraindication.OtherTherapy) {
		return CodeableReferenceInput("contraindication.otherTherapy["+strconv.Itoa(numOtherTherapy)+"].therapy", nil, htmlAttrs)
	}
	return CodeableReferenceInput("contraindication.otherTherapy["+strconv.Itoa(numOtherTherapy)+"].therapy", &resource.Contraindication.OtherTherapy[numOtherTherapy].Therapy, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_IndicationDiseaseSymptomProcedure(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("indication.diseaseSymptomProcedure", nil, htmlAttrs)
	}
	return CodeableReferenceInput("indication.diseaseSymptomProcedure", resource.Indication.DiseaseSymptomProcedure, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_IndicationDiseaseStatus(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("indication.diseaseStatus", nil, htmlAttrs)
	}
	return CodeableReferenceInput("indication.diseaseStatus", resource.Indication.DiseaseStatus, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_IndicationComorbidity(numComorbidity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numComorbidity >= len(resource.Indication.Comorbidity) {
		return CodeableReferenceInput("indication.comorbidity["+strconv.Itoa(numComorbidity)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("indication.comorbidity["+strconv.Itoa(numComorbidity)+"]", &resource.Indication.Comorbidity[numComorbidity], htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_IndicationIntendedEffect(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("indication.intendedEffect", nil, htmlAttrs)
	}
	return CodeableReferenceInput("indication.intendedEffect", resource.Indication.IntendedEffect, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_IndicationDurationRange(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return RangeInput("indication.durationRange", nil, htmlAttrs)
	}
	return RangeInput("indication.durationRange", resource.Indication.DurationRange, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_IndicationDurationString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("indication.durationString", nil, htmlAttrs)
	}
	return StringInput("indication.durationString", resource.Indication.DurationString, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_IndicationUndesirableEffect(numUndesirableEffect int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUndesirableEffect >= len(resource.Indication.UndesirableEffect) {
		return ReferenceInput("indication.undesirableEffect["+strconv.Itoa(numUndesirableEffect)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("indication.undesirableEffect["+strconv.Itoa(numUndesirableEffect)+"]", &resource.Indication.UndesirableEffect[numUndesirableEffect], htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_InteractionType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("interaction.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("interaction.type", resource.Interaction.Type, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_InteractionEffect(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("interaction.effect", nil, htmlAttrs)
	}
	return CodeableReferenceInput("interaction.effect", resource.Interaction.Effect, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_InteractionIncidence(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("interaction.incidence", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("interaction.incidence", resource.Interaction.Incidence, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_InteractionManagement(numManagement int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numManagement >= len(resource.Interaction.Management) {
		return CodeableConceptSelect("interaction.management["+strconv.Itoa(numManagement)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("interaction.management["+strconv.Itoa(numManagement)+"]", &resource.Interaction.Management[numManagement], optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_InteractionInteractantItemReference(numInteractant int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInteractant >= len(resource.Interaction.Interactant) {
		return ReferenceInput("interaction.interactant["+strconv.Itoa(numInteractant)+"].itemReference", nil, htmlAttrs)
	}
	return ReferenceInput("interaction.interactant["+strconv.Itoa(numInteractant)+"].itemReference", &resource.Interaction.Interactant[numInteractant].ItemReference, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_InteractionInteractantItemCodeableConcept(numInteractant int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInteractant >= len(resource.Interaction.Interactant) {
		return CodeableConceptSelect("interaction.interactant["+strconv.Itoa(numInteractant)+"].itemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("interaction.interactant["+strconv.Itoa(numInteractant)+"].itemCodeableConcept", &resource.Interaction.Interactant[numInteractant].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_UndesirableEffectSymptomConditionEffect(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("undesirableEffect.symptomConditionEffect", nil, htmlAttrs)
	}
	return CodeableReferenceInput("undesirableEffect.symptomConditionEffect", resource.UndesirableEffect.SymptomConditionEffect, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_UndesirableEffectClassification(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("undesirableEffect.classification", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("undesirableEffect.classification", resource.UndesirableEffect.Classification, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_UndesirableEffectFrequencyOfOccurrence(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("undesirableEffect.frequencyOfOccurrence", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("undesirableEffect.frequencyOfOccurrence", resource.UndesirableEffect.FrequencyOfOccurrence, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_WarningDescription(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("warning.description", nil, htmlAttrs)
	}
	return StringInput("warning.description", resource.Warning.Description, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_WarningCode(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("warning.code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("warning.code", resource.Warning.Code, optionsValueSet, htmlAttrs)
}
