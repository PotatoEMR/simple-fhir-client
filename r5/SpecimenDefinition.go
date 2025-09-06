package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
func (resource *SpecimenDefinition) T_Url() templ.Component {

	if resource == nil {
		return StringInput("SpecimenDefinition.Url", nil)
	}
	return StringInput("SpecimenDefinition.Url", resource.Url)
}
func (resource *SpecimenDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("SpecimenDefinition.Version", nil)
	}
	return StringInput("SpecimenDefinition.Version", resource.Version)
}
func (resource *SpecimenDefinition) T_Name() templ.Component {

	if resource == nil {
		return StringInput("SpecimenDefinition.Name", nil)
	}
	return StringInput("SpecimenDefinition.Name", resource.Name)
}
func (resource *SpecimenDefinition) T_Title() templ.Component {

	if resource == nil {
		return StringInput("SpecimenDefinition.Title", nil)
	}
	return StringInput("SpecimenDefinition.Title", resource.Title)
}
func (resource *SpecimenDefinition) T_DerivedFromCanonical(numDerivedFromCanonical int) templ.Component {

	if resource == nil || len(resource.DerivedFromCanonical) >= numDerivedFromCanonical {
		return StringInput("SpecimenDefinition.DerivedFromCanonical["+strconv.Itoa(numDerivedFromCanonical)+"]", nil)
	}
	return StringInput("SpecimenDefinition.DerivedFromCanonical["+strconv.Itoa(numDerivedFromCanonical)+"]", &resource.DerivedFromCanonical[numDerivedFromCanonical])
}
func (resource *SpecimenDefinition) T_DerivedFromUri(numDerivedFromUri int) templ.Component {

	if resource == nil || len(resource.DerivedFromUri) >= numDerivedFromUri {
		return StringInput("SpecimenDefinition.DerivedFromUri["+strconv.Itoa(numDerivedFromUri)+"]", nil)
	}
	return StringInput("SpecimenDefinition.DerivedFromUri["+strconv.Itoa(numDerivedFromUri)+"]", &resource.DerivedFromUri[numDerivedFromUri])
}
func (resource *SpecimenDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("SpecimenDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("SpecimenDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *SpecimenDefinition) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("SpecimenDefinition.Experimental", nil)
	}
	return BoolInput("SpecimenDefinition.Experimental", resource.Experimental)
}
func (resource *SpecimenDefinition) T_Date() templ.Component {

	if resource == nil {
		return StringInput("SpecimenDefinition.Date", nil)
	}
	return StringInput("SpecimenDefinition.Date", resource.Date)
}
func (resource *SpecimenDefinition) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("SpecimenDefinition.Publisher", nil)
	}
	return StringInput("SpecimenDefinition.Publisher", resource.Publisher)
}
func (resource *SpecimenDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("SpecimenDefinition.Description", nil)
	}
	return StringInput("SpecimenDefinition.Description", resource.Description)
}
func (resource *SpecimenDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("SpecimenDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SpecimenDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *SpecimenDefinition) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("SpecimenDefinition.Purpose", nil)
	}
	return StringInput("SpecimenDefinition.Purpose", resource.Purpose)
}
func (resource *SpecimenDefinition) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("SpecimenDefinition.Copyright", nil)
	}
	return StringInput("SpecimenDefinition.Copyright", resource.Copyright)
}
func (resource *SpecimenDefinition) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("SpecimenDefinition.CopyrightLabel", nil)
	}
	return StringInput("SpecimenDefinition.CopyrightLabel", resource.CopyrightLabel)
}
func (resource *SpecimenDefinition) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("SpecimenDefinition.ApprovalDate", nil)
	}
	return StringInput("SpecimenDefinition.ApprovalDate", resource.ApprovalDate)
}
func (resource *SpecimenDefinition) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("SpecimenDefinition.LastReviewDate", nil)
	}
	return StringInput("SpecimenDefinition.LastReviewDate", resource.LastReviewDate)
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
func (resource *SpecimenDefinition) T_TypeTestedSingleUse(numTypeTested int) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested {
		return BoolInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].SingleUse", nil)
	}
	return BoolInput("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].SingleUse", resource.TypeTested[numTypeTested].SingleUse)
}
func (resource *SpecimenDefinition) T_TypeTestedRejectionCriterion(numTypeTested int, numRejectionCriterion int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested || len(resource.TypeTested[numTypeTested].RejectionCriterion) >= numRejectionCriterion {
		return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].RejectionCriterion["+strconv.Itoa(numRejectionCriterion)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].RejectionCriterion["+strconv.Itoa(numRejectionCriterion)+"]", &resource.TypeTested[numTypeTested].RejectionCriterion[numRejectionCriterion], optionsValueSet)
}
func (resource *SpecimenDefinition) T_TypeTestedTestingDestination(numTypeTested int, numTestingDestination int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.TypeTested) >= numTypeTested || len(resource.TypeTested[numTypeTested].TestingDestination) >= numTestingDestination {
		return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].TestingDestination["+strconv.Itoa(numTestingDestination)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SpecimenDefinition.TypeTested["+strconv.Itoa(numTypeTested)+"].TestingDestination["+strconv.Itoa(numTestingDestination)+"]", &resource.TypeTested[numTypeTested].TestingDestination[numTestingDestination], optionsValueSet)
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
