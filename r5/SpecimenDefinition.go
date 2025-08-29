package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/SpecimenDefinition
type SpecimenDefinition struct {
	Id                     *string                        `json:"id,omitempty"`
	Meta                   *Meta                          `json:"meta,omitempty"`
	ImplicitRules          *string                        `json:"implicitRules,omitempty"`
	Language               *string                        `json:"language,omitempty"`
	Text                   *Narrative                     `json:"text,omitempty"`
	Contained              []Resource                     `json:"contained,omitempty"`
	Extension              []Extension                    `json:"extension,omitempty"`
	ModifierExtension      []Extension                    `json:"modifierExtension,omitempty"`
	Url                    *string                        `json:"url,omitempty"`
	Identifier             *Identifier                    `json:"identifier,omitempty"`
	Version                *string                        `json:"version,omitempty"`
	VersionAlgorithmString *string                        `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                        `json:"versionAlgorithmCoding"`
	Name                   *string                        `json:"name,omitempty"`
	Title                  *string                        `json:"title,omitempty"`
	DerivedFromCanonical   []string                       `json:"derivedFromCanonical,omitempty"`
	DerivedFromUri         []string                       `json:"derivedFromUri,omitempty"`
	Status                 string                         `json:"status"`
	Experimental           *bool                          `json:"experimental,omitempty"`
	SubjectCodeableConcept *CodeableConcept               `json:"subjectCodeableConcept"`
	SubjectReference       *Reference                     `json:"subjectReference"`
	Date                   *string                        `json:"date,omitempty"`
	Publisher              *string                        `json:"publisher,omitempty"`
	Contact                []ContactDetail                `json:"contact,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	UseContext             []UsageContext                 `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept              `json:"jurisdiction,omitempty"`
	Purpose                *string                        `json:"purpose,omitempty"`
	Copyright              *string                        `json:"copyright,omitempty"`
	CopyrightLabel         *string                        `json:"copyrightLabel,omitempty"`
	ApprovalDate           *string                        `json:"approvalDate,omitempty"`
	LastReviewDate         *string                        `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                        `json:"effectivePeriod,omitempty"`
	TypeCollected          *CodeableConcept               `json:"typeCollected,omitempty"`
	PatientPreparation     []CodeableConcept              `json:"patientPreparation,omitempty"`
	TimeAspect             *string                        `json:"timeAspect,omitempty"`
	Collection             []CodeableConcept              `json:"collection,omitempty"`
	TypeTested             []SpecimenDefinitionTypeTested `json:"typeTested,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SpecimenDefinition
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
	SingleUse          *bool                                  `json:"singleUse,omitempty"`
	RejectionCriterion []CodeableConcept                      `json:"rejectionCriterion,omitempty"`
	Handling           []SpecimenDefinitionTypeTestedHandling `json:"handling,omitempty"`
	TestingDestination []CodeableConcept                      `json:"testingDestination,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SpecimenDefinition
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

// http://hl7.org/fhir/r5/StructureDefinition/SpecimenDefinition
type SpecimenDefinitionTypeTestedContainerAdditive struct {
	Id                      *string         `json:"id,omitempty"`
	Extension               []Extension     `json:"extension,omitempty"`
	ModifierExtension       []Extension     `json:"modifierExtension,omitempty"`
	AdditiveCodeableConcept CodeableConcept `json:"additiveCodeableConcept"`
	AdditiveReference       Reference       `json:"additiveReference"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SpecimenDefinition
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

func (resource *SpecimenDefinition) SpecimenDefinitionLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *SpecimenDefinition) SpecimenDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *SpecimenDefinition) SpecimenDefinitionTypeTestedPreference(numTypeTested int) templ.Component {
	optionsValueSet := VSSpecimen_contained_preference
	currentVal := ""
	if resource != nil && len(resource.TypeTested) >= numTypeTested {
		currentVal = resource.TypeTested[numTypeTested].Preference
	}
	return CodeSelect("preference", currentVal, optionsValueSet)
}
