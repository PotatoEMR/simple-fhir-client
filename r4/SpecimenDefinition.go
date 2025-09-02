package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/SpecimenDefinition
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

// http://hl7.org/fhir/r4/StructureDefinition/SpecimenDefinition
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

// http://hl7.org/fhir/r4/StructureDefinition/SpecimenDefinition
type SpecimenDefinitionTypeTestedContainer struct {
	Id                    *string                                         `json:"id,omitempty"`
	Extension             []Extension                                     `json:"extension,omitempty"`
	ModifierExtension     []Extension                                     `json:"modifierExtension,omitempty"`
	Material              *CodeableConcept                                `json:"material,omitempty"`
	Type                  *CodeableConcept                                `json:"type,omitempty"`
	Cap                   *CodeableConcept                                `json:"cap,omitempty"`
	Description           *string                                         `json:"description,omitempty"`
	Capacity              *Quantity                                       `json:"capacity,omitempty"`
	MinimumVolumeQuantity *Quantity                                       `json:"minimumVolumeQuantity"`
	MinimumVolumeString   *string                                         `json:"minimumVolumeString"`
	Additive              []SpecimenDefinitionTypeTestedContainerAdditive `json:"additive,omitempty"`
	Preparation           *string                                         `json:"preparation,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SpecimenDefinition
type SpecimenDefinitionTypeTestedContainerAdditive struct {
	Id                      *string         `json:"id,omitempty"`
	Extension               []Extension     `json:"extension,omitempty"`
	ModifierExtension       []Extension     `json:"modifierExtension,omitempty"`
	AdditiveCodeableConcept CodeableConcept `json:"additiveCodeableConcept"`
	AdditiveReference       Reference       `json:"additiveReference"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SpecimenDefinition
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

func (resource *SpecimenDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeCollected(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("typeCollected", nil, optionsValueSet)
	}
	return CodeableConceptSelect("typeCollected", resource.TypeCollected, optionsValueSet)
}
func (resource *SpecimenDefinition) T_PatientPreparation(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("patientPreparation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("patientPreparation", &resource.PatientPreparation[0], optionsValueSet)
}
func (resource *SpecimenDefinition) T_Collection(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("collection", nil, optionsValueSet)
	}
	return CodeableConceptSelect("collection", &resource.Collection[0], optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeTestedType(numTypeTested int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.TypeTested) >= numTypeTested {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.TypeTested[numTypeTested].Type, optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeTestedPreference(numTypeTested int) templ.Component {
	optionsValueSet := VSSpecimen_contained_preference

	if resource == nil && len(resource.TypeTested) >= numTypeTested {
		return CodeSelect("preference", nil, optionsValueSet)
	}
	return CodeSelect("preference", &resource.TypeTested[numTypeTested].Preference, optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeTestedRejectionCriterion(numTypeTested int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.TypeTested) >= numTypeTested {
		return CodeableConceptSelect("rejectionCriterion", nil, optionsValueSet)
	}
	return CodeableConceptSelect("rejectionCriterion", &resource.TypeTested[numTypeTested].RejectionCriterion[0], optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerMaterial(numTypeTested int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.TypeTested) >= numTypeTested {
		return CodeableConceptSelect("material", nil, optionsValueSet)
	}
	return CodeableConceptSelect("material", resource.TypeTested[numTypeTested].Container.Material, optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerType(numTypeTested int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.TypeTested) >= numTypeTested {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.TypeTested[numTypeTested].Container.Type, optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerCap(numTypeTested int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.TypeTested) >= numTypeTested {
		return CodeableConceptSelect("cap", nil, optionsValueSet)
	}
	return CodeableConceptSelect("cap", resource.TypeTested[numTypeTested].Container.Cap, optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeTestedHandlingTemperatureQualifier(numTypeTested int, numHandling int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.TypeTested[numTypeTested].Handling) >= numHandling {
		return CodeableConceptSelect("temperatureQualifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("temperatureQualifier", resource.TypeTested[numTypeTested].Handling[numHandling].TemperatureQualifier, optionsValueSet)
}
