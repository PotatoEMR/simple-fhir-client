package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

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
	VersionAlgorithmString *string                        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                        `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                        `json:"name,omitempty"`
	Title                  *string                        `json:"title,omitempty"`
	DerivedFromCanonical   []string                       `json:"derivedFromCanonical,omitempty"`
	DerivedFromUri         []string                       `json:"derivedFromUri,omitempty"`
	Status                 string                         `json:"status"`
	Experimental           *bool                          `json:"experimental,omitempty"`
	SubjectCodeableConcept *CodeableConcept               `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference                     `json:"subjectReference,omitempty"`
	Date                   *time.Time                     `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher              *string                        `json:"publisher,omitempty"`
	Contact                []ContactDetail                `json:"contact,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	UseContext             []UsageContext                 `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept              `json:"jurisdiction,omitempty"`
	Purpose                *string                        `json:"purpose,omitempty"`
	Copyright              *string                        `json:"copyright,omitempty"`
	CopyrightLabel         *string                        `json:"copyrightLabel,omitempty"`
	ApprovalDate           *time.Time                     `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate         *time.Time                     `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
	MinimumVolumeQuantity *Quantity                                       `json:"minimumVolumeQuantity,omitempty"`
	MinimumVolumeString   *string                                         `json:"minimumVolumeString,omitempty"`
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
func (r SpecimenDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "SpecimenDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	ref.Identifier = r.Identifier
	rtype := "SpecimenDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *SpecimenDefinition) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SpecimenDefinition.Url", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.Url", resource.Url, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SpecimenDefinition.Version", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.Version", resource.Version, htmlAttrs)
}
func (resource *SpecimenDefinition) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SpecimenDefinition.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *SpecimenDefinition) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("SpecimenDefinition.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("SpecimenDefinition.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SpecimenDefinition.Name", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.Name", resource.Name, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SpecimenDefinition.Title", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.Title", resource.Title, htmlAttrs)
}
func (resource *SpecimenDefinition) T_DerivedFromCanonical(numDerivedFromCanonical int, htmlAttrs string) templ.Component {
	if resource == nil || numDerivedFromCanonical >= len(resource.DerivedFromCanonical) {
		return StringInput("SpecimenDefinition.DerivedFromCanonical["+strconv.Itoa(numDerivedFromCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.DerivedFromCanonical["+strconv.Itoa(numDerivedFromCanonical)+"]", &resource.DerivedFromCanonical[numDerivedFromCanonical], htmlAttrs)
}
func (resource *SpecimenDefinition) T_DerivedFromUri(numDerivedFromUri int, htmlAttrs string) templ.Component {
	if resource == nil || numDerivedFromUri >= len(resource.DerivedFromUri) {
		return StringInput("SpecimenDefinition.DerivedFromUri["+strconv.Itoa(numDerivedFromUri)+"]", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.DerivedFromUri["+strconv.Itoa(numDerivedFromUri)+"]", &resource.DerivedFromUri[numDerivedFromUri], htmlAttrs)
}
func (resource *SpecimenDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("SpecimenDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SpecimenDefinition.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("SpecimenDefinition.Experimental", nil, htmlAttrs)
	}
	return BoolInput("SpecimenDefinition.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *SpecimenDefinition) T_SubjectCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SpecimenDefinition.SubjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecimenDefinition.SubjectCodeableConcept", resource.SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("SpecimenDefinition.Date", nil, htmlAttrs)
	}
	return DateTimeInput("SpecimenDefinition.Date", resource.Date, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SpecimenDefinition.Publisher", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SpecimenDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("SpecimenDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecimenDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SpecimenDefinition.Purpose", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SpecimenDefinition.Copyright", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *SpecimenDefinition) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SpecimenDefinition.CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *SpecimenDefinition) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("SpecimenDefinition.ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("SpecimenDefinition.ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *SpecimenDefinition) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("SpecimenDefinition.LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("SpecimenDefinition.LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeCollected(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SpecimenDefinition.TypeCollected", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeCollected", resource.TypeCollected, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_PatientPreparation(numPatientPreparation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPatientPreparation >= len(resource.PatientPreparation) {
		return CodeableConceptSelect("SpecimenDefinition.PatientPreparation["+strconv.Itoa(numPatientPreparation)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecimenDefinition.PatientPreparation["+strconv.Itoa(numPatientPreparation)+"]", &resource.PatientPreparation[numPatientPreparation], optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TimeAspect(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("SpecimenDefinition.TimeAspect", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.TimeAspect", resource.TimeAspect, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Collection(numCollection int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCollection >= len(resource.Collection) {
		return CodeableConceptSelect("SpecimenDefinition.Collection["+strconv.Itoa(numCollection)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecimenDefinition.Collection["+strconv.Itoa(numCollection)+"]", &resource.Collection[numCollection], optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedIsDerived(numTypeTested int, htmlAttrs string) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return BoolInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].IsDerived", nil, htmlAttrs)
	}
	return BoolInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].IsDerived", resource.TypeTested[numTypeTested].IsDerived, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedType(numTypeTested int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Type", resource.TypeTested[numTypeTested].Type, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedPreference(numTypeTested int, htmlAttrs string) templ.Component {
	optionsValueSet := VSSpecimen_contained_preference

	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return CodeSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Preference", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Preference", &resource.TypeTested[numTypeTested].Preference, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedRequirement(numTypeTested int, htmlAttrs string) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Requirement", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Requirement", resource.TypeTested[numTypeTested].Requirement, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedSingleUse(numTypeTested int, htmlAttrs string) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return BoolInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].SingleUse", nil, htmlAttrs)
	}
	return BoolInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].SingleUse", resource.TypeTested[numTypeTested].SingleUse, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedRejectionCriterion(numTypeTested int, numRejectionCriterion int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) || numRejectionCriterion >= len(resource.TypeTested[numTypeTested].RejectionCriterion) {
		return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].RejectionCriterion["+strconv.Itoa(numRejectionCriterion)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].RejectionCriterion["+strconv.Itoa(numRejectionCriterion)+"]", &resource.TypeTested[numTypeTested].RejectionCriterion[numRejectionCriterion], optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedTestingDestination(numTypeTested int, numTestingDestination int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) || numTestingDestination >= len(resource.TypeTested[numTypeTested].TestingDestination) {
		return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].TestingDestination["+strconv.Itoa(numTestingDestination)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].TestingDestination["+strconv.Itoa(numTestingDestination)+"]", &resource.TypeTested[numTypeTested].TestingDestination[numTestingDestination], optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerMaterial(numTypeTested int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Material", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Material", resource.TypeTested[numTypeTested].Container.Material, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerType(numTypeTested int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Type", resource.TypeTested[numTypeTested].Container.Type, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerCap(numTypeTested int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Cap", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Cap", resource.TypeTested[numTypeTested].Container.Cap, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerDescription(numTypeTested int, htmlAttrs string) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Description", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Description", resource.TypeTested[numTypeTested].Container.Description, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerMinimumVolumeString(numTypeTested int, htmlAttrs string) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.MinimumVolumeString", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.MinimumVolumeString", resource.TypeTested[numTypeTested].Container.MinimumVolumeString, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerPreparation(numTypeTested int, htmlAttrs string) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Preparation", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Preparation", resource.TypeTested[numTypeTested].Container.Preparation, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerAdditiveAdditiveCodeableConcept(numTypeTested int, numAdditive int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) || numAdditive >= len(resource.TypeTested[numTypeTested].Container.Additive) {
		return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Additive["+strconv.Itoa(numAdditive)+"].AdditiveCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Container.Additive["+strconv.Itoa(numAdditive)+"].AdditiveCodeableConcept", &resource.TypeTested[numTypeTested].Container.Additive[numAdditive].AdditiveCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedHandlingTemperatureQualifier(numTypeTested int, numHandling int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) || numHandling >= len(resource.TypeTested[numTypeTested].Handling) {
		return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Handling["+strconv.Itoa(numHandling)+"].TemperatureQualifier", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Handling["+strconv.Itoa(numHandling)+"].TemperatureQualifier", resource.TypeTested[numTypeTested].Handling[numHandling].TemperatureQualifier, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedHandlingInstruction(numTypeTested int, numHandling int, htmlAttrs string) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) || numHandling >= len(resource.TypeTested[numTypeTested].Handling) {
		return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Handling["+strconv.Itoa(numHandling)+"].Instruction", nil, htmlAttrs)
	}
	return StringInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].Handling["+strconv.Itoa(numHandling)+"].Instruction", resource.TypeTested[numTypeTested].Handling[numHandling].Instruction, htmlAttrs)
}
