package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/SpecimenDefinition
type SpecimenDefinition struct {
	Id                 *string                        `json:"id,omitempty"`
	Meta               *Meta                          `json:"meta,omitempty"`
	ImplicitRules      *string                        `json:"implicitRules,omitempty"`
	Language           *string                        `json:"language,omitempty"`
	Text               *Narrative                     `json:"text,omitempty"`
	Contained          []Resource                     `json:"contained,omitempty"`
	Extension          []Extension                    `json:"extension,omitempty"`
	ModifierExtension  []Extension                    `json:"modifierExtension,omitempty"`
	Identifier         *Identifier                    `json:"identifier,omitempty"`
	TypeCollected      *CodeableConcept               `json:"typeCollected,omitempty"`
	PatientPreparation []CodeableConcept              `json:"patientPreparation,omitempty"`
	TimeAspect         *string                        `json:"timeAspect,omitempty"`
	Collection         []CodeableConcept              `json:"collection,omitempty"`
	TypeTested         []SpecimenDefinitionTypeTested `json:"typeTested,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SpecimenDefinition
type SpecimenDefinitionTypeTested struct {
	Id                 *string                                `json:"id,omitempty"`
	Extension          []Extension                            `json:"extension,omitempty"`
	ModifierExtension  []Extension                            `json:"modifierExtension,omitempty"`
	IsDerived          *bool                                  `json:"isDerived,omitempty"`
	Type               *CodeableConcept                       `json:"type,omitempty"`
	Preference         string                                 `json:"preference"`
	Container          *SpecimenDefinitionTypeTestedContainer `json:"container,omitempty"`
	Requirement        *string                                `json:"requirement,omitempty"`
	RetentionTime      *Duration                              `json:"retentionTime,omitempty"`
	RejectionCriterion []CodeableConcept                      `json:"rejectionCriterion,omitempty"`
	Handling           []SpecimenDefinitionTypeTestedHandling `json:"handling,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SpecimenDefinition
type SpecimenDefinitionTypeTestedContainer struct {
	Id                    *string                                         `json:"id,omitempty"`
	Extension             []Extension                                     `json:"extension,omitempty"`
	ModifierExtension     []Extension                                     `json:"modifierExtension,omitempty"`
	Material              *CodeableConcept                                `json:"material,omitempty"`
	Type                  *CodeableConcept                                `json:"type,omitempty"`
	Cap                   *CodeableConcept                                `json:"cap,omitempty"`
	Description           *string                                         `json:"description,omitempty"`
	Capacity              *Quantity                                       `json:"capacity,omitempty"`
	MinimumVolumeQuantity *Quantity                                       `json:"minimumVolumeQuantity,omitempty"`
	MinimumVolumeString   *string                                         `json:"minimumVolumeString,omitempty"`
	Additive              []SpecimenDefinitionTypeTestedContainerAdditive `json:"additive,omitempty"`
	Preparation           *string                                         `json:"preparation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SpecimenDefinition
type SpecimenDefinitionTypeTestedContainerAdditive struct {
	Id                      *string         `json:"id,omitempty"`
	Extension               []Extension     `json:"extension,omitempty"`
	ModifierExtension       []Extension     `json:"modifierExtension,omitempty"`
	AdditiveCodeableConcept CodeableConcept `json:"additiveCodeableConcept"`
	AdditiveReference       Reference       `json:"additiveReference"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/SpecimenDefinition
type SpecimenDefinitionTypeTestedHandling struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	TemperatureQualifier *CodeableConcept `json:"temperatureQualifier,omitempty"`
	TemperatureRange     *Range           `json:"temperatureRange,omitempty"`
	MaxDuration          *Duration        `json:"maxDuration,omitempty"`
	Instruction          *string          `json:"instruction,omitempty"`
}

type OtherSpecimenDefinition SpecimenDefinition

// on convert struct to json, automatically add resourceType=SpecimenDefinition
func (r SpecimenDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSpecimenDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherSpecimenDefinition: OtherSpecimenDefinition(r),
		ResourceType:            "SpecimenDefinition",
	})
}

func (resource *SpecimenDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("SpecimenDefinition.Id", nil)
	}
	return StringInput("SpecimenDefinition.Id", resource.Id)
}
func (resource *SpecimenDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("SpecimenDefinition.ImplicitRules", nil)
	}
	return StringInput("SpecimenDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *SpecimenDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("SpecimenDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("SpecimenDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeCollected(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SpecimenDefinition.TypeCollected", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeCollected", resource.TypeCollected, optionsValueSet)
}
func (resource *SpecimenDefinition) T_PatientPreparation(numPatientPreparation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PatientPreparation) >= numPatientPreparation {
		return CodeableConceptSelect("SpecimenDefinition.PatientPreparation["+strconv.Itoa(numPatientPreparation)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SpecimenDefinition.PatientPreparation["+strconv.Itoa(numPatientPreparation)+"]", &resource.PatientPreparation[numPatientPreparation], optionsValueSet)
}
func (resource *SpecimenDefinition) T_TimeAspect() templ.Component {

	if resource == nil {
		return StringInput("SpecimenDefinition.TimeAspect", nil)
	}
	return StringInput("SpecimenDefinition.TimeAspect", resource.TimeAspect)
}
func (resource *SpecimenDefinition) T_Collection(numCollection int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Collection) >= numCollection {
		return CodeableConceptSelect("SpecimenDefinition.Collection["+strconv.Itoa(numCollection)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SpecimenDefinition.Collection["+strconv.Itoa(numCollection)+"]", &resource.Collection[numCollection], optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeTestedId(numTypeTested int) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested {
		return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Id", nil)
	}
	return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Id", resource.TypeTested[numTypeTested].Id)
}
func (resource *SpecimenDefinition) T_TypeTestedIsDerived(numTypeTested int) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested {
		return BoolInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].IsDerived", nil)
	}
	return BoolInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].IsDerived", resource.TypeTested[numTypeTested].IsDerived)
}
func (resource *SpecimenDefinition) T_TypeTestedType(numTypeTested int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested {
		return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Type", resource.TypeTested[numTypeTested].Type, optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeTestedPreference(numTypeTested int) templ.Component {
	optionsValueSet := VSSpecimen_contained_preference

	if resource == nil || len(resource.TypeTested) >= numTypeTested {
		return CodeSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Preference", nil, optionsValueSet)
	}
	return CodeSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Preference", &resource.TypeTested[numTypeTested].Preference, optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeTestedRequirement(numTypeTested int) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested {
		return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Requirement", nil)
	}
	return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Requirement", resource.TypeTested[numTypeTested].Requirement)
}
func (resource *SpecimenDefinition) T_TypeTestedRejectionCriterion(numTypeTested int, numRejectionCriterion int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested || len(resource.TypeTested[numTypeTested].RejectionCriterion) >= numRejectionCriterion {
		return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].RejectionCriterion["+strconv.Itoa(numRejectionCriterion)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].RejectionCriterion["+strconv.Itoa(numRejectionCriterion)+"]", &resource.TypeTested[numTypeTested].RejectionCriterion[numRejectionCriterion], optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerId(numTypeTested int) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested {
		return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Id", nil)
	}
	return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Id", resource.TypeTested[numTypeTested].Container.Id)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerMaterial(numTypeTested int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested {
		return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Material", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Material", resource.TypeTested[numTypeTested].Container.Material, optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerType(numTypeTested int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested {
		return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Type", resource.TypeTested[numTypeTested].Container.Type, optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerCap(numTypeTested int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested {
		return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Cap", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Cap", resource.TypeTested[numTypeTested].Container.Cap, optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerDescription(numTypeTested int) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested {
		return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Description", nil)
	}
	return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Description", resource.TypeTested[numTypeTested].Container.Description)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerPreparation(numTypeTested int) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested {
		return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Preparation", nil)
	}
	return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Preparation", resource.TypeTested[numTypeTested].Container.Preparation)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerAdditiveId(numTypeTested int, numAdditive int) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested || len(resource.TypeTested[numTypeTested].Container.Additive) >= numAdditive {
		return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Additive["+strconv.Itoa(numAdditive)+"].Id", nil)
	}
	return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Additive["+strconv.Itoa(numAdditive)+"].Id", resource.TypeTested[numTypeTested].Container.Additive[numAdditive].Id)
}
func (resource *SpecimenDefinition) T_TypeTestedHandlingId(numTypeTested int, numHandling int) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested || len(resource.TypeTested[numTypeTested].Handling) >= numHandling {
		return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Handling["+strconv.Itoa(numHandling)+"].Id", nil)
	}
	return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Handling["+strconv.Itoa(numHandling)+"].Id", resource.TypeTested[numTypeTested].Handling[numHandling].Id)
}
func (resource *SpecimenDefinition) T_TypeTestedHandlingTemperatureQualifier(numTypeTested int, numHandling int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested || len(resource.TypeTested[numTypeTested].Handling) >= numHandling {
		return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Handling["+strconv.Itoa(numHandling)+"].TemperatureQualifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Handling["+strconv.Itoa(numHandling)+"].TemperatureQualifier", resource.TypeTested[numTypeTested].Handling[numHandling].TemperatureQualifier, optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeTestedHandlingInstruction(numTypeTested int, numHandling int) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested || len(resource.TypeTested[numTypeTested].Handling) >= numHandling {
		return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Handling["+strconv.Itoa(numHandling)+"].Instruction", nil)
	}
	return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Handling["+strconv.Itoa(numHandling)+"].Instruction", resource.TypeTested[numTypeTested].Handling[numHandling].Instruction)
}
