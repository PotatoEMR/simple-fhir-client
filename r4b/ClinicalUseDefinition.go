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
func (resource *ClinicalUseDefinition) T_Type(htmlAttrs string) templ.Component {
	optionsValueSet := VSClinical_use_definition_type

	if resource == nil {
		return CodeSelect("ClinicalUseDefinition.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ClinicalUseDefinition.Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("ClinicalUseDefinition.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_Status(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ClinicalUseDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_ContraindicationOtherTherapyRelationshipType(numOtherTherapy int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOtherTherapy >= len(resource.Contraindication.OtherTherapy) {
		return CodeableConceptSelect("ClinicalUseDefinition.Contraindication.OtherTherapy["+strconv.Itoa(numOtherTherapy)+"].RelationshipType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.Contraindication.OtherTherapy["+strconv.Itoa(numOtherTherapy)+"].RelationshipType", &resource.Contraindication.OtherTherapy[numOtherTherapy].RelationshipType, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_IndicationDurationString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ClinicalUseDefinition.Indication.DurationString", nil, htmlAttrs)
	}
	return StringInput("ClinicalUseDefinition.Indication.DurationString", resource.Indication.DurationString, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_InteractionType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ClinicalUseDefinition.Interaction.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.Interaction.Type", resource.Interaction.Type, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_InteractionIncidence(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ClinicalUseDefinition.Interaction.Incidence", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.Interaction.Incidence", resource.Interaction.Incidence, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_InteractionManagement(numManagement int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numManagement >= len(resource.Interaction.Management) {
		return CodeableConceptSelect("ClinicalUseDefinition.Interaction.Management["+strconv.Itoa(numManagement)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.Interaction.Management["+strconv.Itoa(numManagement)+"]", &resource.Interaction.Management[numManagement], optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_InteractionInteractantItemCodeableConcept(numInteractant int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numInteractant >= len(resource.Interaction.Interactant) {
		return CodeableConceptSelect("ClinicalUseDefinition.Interaction.Interactant["+strconv.Itoa(numInteractant)+"].ItemCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.Interaction.Interactant["+strconv.Itoa(numInteractant)+"].ItemCodeableConcept", &resource.Interaction.Interactant[numInteractant].ItemCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_UndesirableEffectClassification(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ClinicalUseDefinition.UndesirableEffect.Classification", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.UndesirableEffect.Classification", resource.UndesirableEffect.Classification, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_UndesirableEffectFrequencyOfOccurrence(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ClinicalUseDefinition.UndesirableEffect.FrequencyOfOccurrence", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.UndesirableEffect.FrequencyOfOccurrence", resource.UndesirableEffect.FrequencyOfOccurrence, optionsValueSet, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_WarningDescription(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ClinicalUseDefinition.Warning.Description", nil, htmlAttrs)
	}
	return StringInput("ClinicalUseDefinition.Warning.Description", resource.Warning.Description, htmlAttrs)
}
func (resource *ClinicalUseDefinition) T_WarningCode(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("ClinicalUseDefinition.Warning.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ClinicalUseDefinition.Warning.Code", resource.Warning.Code, optionsValueSet, htmlAttrs)
}
