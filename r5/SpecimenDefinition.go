package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
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
	Date                   *FhirDateTime                  `json:"date,omitempty"`
	Publisher              *string                        `json:"publisher,omitempty"`
	Contact                []ContactDetail                `json:"contact,omitempty"`
	Description            *string                        `json:"description,omitempty"`
	UseContext             []UsageContext                 `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept              `json:"jurisdiction,omitempty"`
	Purpose                *string                        `json:"purpose,omitempty"`
	Copyright              *string                        `json:"copyright,omitempty"`
	CopyrightLabel         *string                        `json:"copyrightLabel,omitempty"`
	ApprovalDate           *FhirDate                      `json:"approvalDate,omitempty"`
	LastReviewDate         *FhirDate                      `json:"lastReviewDate,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
func (r SpecimenDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSpecimenDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherSpecimenDefinition: OtherSpecimenDefinition(r),
		ResourceType:            "SpecimenDefinition",
	})
}

// json -> struct, first reject if resourceType != SpecimenDefinition
func (r *SpecimenDefinition) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "SpecimenDefinition" {
		return errors.New("resourceType not SpecimenDefinition")
	}
	return json.Unmarshal(data, (*OtherSpecimenDefinition)(r))
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
func (resource *SpecimenDefinition) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *SpecimenDefinition) T_VersionAlgorithmString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *SpecimenDefinition) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *SpecimenDefinition) T_DerivedFromCanonical(numDerivedFromCanonical int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDerivedFromCanonical >= len(resource.DerivedFromCanonical) {
		return StringInput("derivedFromCanonical["+strconv.Itoa(numDerivedFromCanonical)+"]", nil, htmlAttrs)
	}
	return StringInput("derivedFromCanonical["+strconv.Itoa(numDerivedFromCanonical)+"]", &resource.DerivedFromCanonical[numDerivedFromCanonical], htmlAttrs)
}
func (resource *SpecimenDefinition) T_DerivedFromUri(numDerivedFromUri int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDerivedFromUri >= len(resource.DerivedFromUri) {
		return StringInput("derivedFromUri["+strconv.Itoa(numDerivedFromUri)+"]", nil, htmlAttrs)
	}
	return StringInput("derivedFromUri["+strconv.Itoa(numDerivedFromUri)+"]", &resource.DerivedFromUri[numDerivedFromUri], htmlAttrs)
}
func (resource *SpecimenDefinition) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *SpecimenDefinition) T_SubjectCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("subjectCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subjectCodeableConcept", resource.SubjectCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_SubjectReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subjectReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subjectReference", resource.SubjectReference, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *SpecimenDefinition) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *SpecimenDefinition) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *SpecimenDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *SpecimenDefinition) T_CopyrightLabel(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *SpecimenDefinition) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("approvalDate", nil, htmlAttrs)
	}
	return FhirDateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *SpecimenDefinition) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("lastReviewDate", nil, htmlAttrs)
	}
	return FhirDateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *SpecimenDefinition) T_EffectivePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("effectivePeriod", resource.EffectivePeriod, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeCollected(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("typeCollected", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("typeCollected", resource.TypeCollected, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_PatientPreparation(numPatientPreparation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPatientPreparation >= len(resource.PatientPreparation) {
		return CodeableConceptSelect("patientPreparation["+strconv.Itoa(numPatientPreparation)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("patientPreparation["+strconv.Itoa(numPatientPreparation)+"]", &resource.PatientPreparation[numPatientPreparation], optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TimeAspect(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("timeAspect", nil, htmlAttrs)
	}
	return StringInput("timeAspect", resource.TimeAspect, htmlAttrs)
}
func (resource *SpecimenDefinition) T_Collection(numCollection int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCollection >= len(resource.Collection) {
		return CodeableConceptSelect("collection["+strconv.Itoa(numCollection)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("collection["+strconv.Itoa(numCollection)+"]", &resource.Collection[numCollection], optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedIsDerived(numTypeTested int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return BoolInput("typeTested["+strconv.Itoa(numTypeTested)+"].isDerived", nil, htmlAttrs)
	}
	return BoolInput("typeTested["+strconv.Itoa(numTypeTested)+"].isDerived", resource.TypeTested[numTypeTested].IsDerived, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedType(numTypeTested int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return CodeableConceptSelect("typeTested["+strconv.Itoa(numTypeTested)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("typeTested["+strconv.Itoa(numTypeTested)+"].type", resource.TypeTested[numTypeTested].Type, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedPreference(numTypeTested int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSSpecimen_contained_preference

	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return CodeSelect("typeTested["+strconv.Itoa(numTypeTested)+"].preference", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("typeTested["+strconv.Itoa(numTypeTested)+"].preference", &resource.TypeTested[numTypeTested].Preference, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedRequirement(numTypeTested int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return StringInput("typeTested["+strconv.Itoa(numTypeTested)+"].requirement", nil, htmlAttrs)
	}
	return StringInput("typeTested["+strconv.Itoa(numTypeTested)+"].requirement", resource.TypeTested[numTypeTested].Requirement, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedRetentionTime(numTypeTested int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return DurationInput("typeTested["+strconv.Itoa(numTypeTested)+"].retentionTime", nil, htmlAttrs)
	}
	return DurationInput("typeTested["+strconv.Itoa(numTypeTested)+"].retentionTime", resource.TypeTested[numTypeTested].RetentionTime, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedSingleUse(numTypeTested int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return BoolInput("typeTested["+strconv.Itoa(numTypeTested)+"].singleUse", nil, htmlAttrs)
	}
	return BoolInput("typeTested["+strconv.Itoa(numTypeTested)+"].singleUse", resource.TypeTested[numTypeTested].SingleUse, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedRejectionCriterion(numTypeTested int, numRejectionCriterion int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) || numRejectionCriterion >= len(resource.TypeTested[numTypeTested].RejectionCriterion) {
		return CodeableConceptSelect("typeTested["+strconv.Itoa(numTypeTested)+"].rejectionCriterion["+strconv.Itoa(numRejectionCriterion)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("typeTested["+strconv.Itoa(numTypeTested)+"].rejectionCriterion["+strconv.Itoa(numRejectionCriterion)+"]", &resource.TypeTested[numTypeTested].RejectionCriterion[numRejectionCriterion], optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedTestingDestination(numTypeTested int, numTestingDestination int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) || numTestingDestination >= len(resource.TypeTested[numTypeTested].TestingDestination) {
		return CodeableConceptSelect("typeTested["+strconv.Itoa(numTypeTested)+"].testingDestination["+strconv.Itoa(numTestingDestination)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("typeTested["+strconv.Itoa(numTypeTested)+"].testingDestination["+strconv.Itoa(numTestingDestination)+"]", &resource.TypeTested[numTypeTested].TestingDestination[numTestingDestination], optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerMaterial(numTypeTested int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return CodeableConceptSelect("typeTested["+strconv.Itoa(numTypeTested)+"].container.material", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("typeTested["+strconv.Itoa(numTypeTested)+"].container.material", resource.TypeTested[numTypeTested].Container.Material, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerType(numTypeTested int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return CodeableConceptSelect("typeTested["+strconv.Itoa(numTypeTested)+"].container.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("typeTested["+strconv.Itoa(numTypeTested)+"].container.type", resource.TypeTested[numTypeTested].Container.Type, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerCap(numTypeTested int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return CodeableConceptSelect("typeTested["+strconv.Itoa(numTypeTested)+"].container.cap", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("typeTested["+strconv.Itoa(numTypeTested)+"].container.cap", resource.TypeTested[numTypeTested].Container.Cap, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerDescription(numTypeTested int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return StringInput("typeTested["+strconv.Itoa(numTypeTested)+"].container.description", nil, htmlAttrs)
	}
	return StringInput("typeTested["+strconv.Itoa(numTypeTested)+"].container.description", resource.TypeTested[numTypeTested].Container.Description, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerCapacity(numTypeTested int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return QuantityInput("typeTested["+strconv.Itoa(numTypeTested)+"].container.capacity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("typeTested["+strconv.Itoa(numTypeTested)+"].container.capacity", resource.TypeTested[numTypeTested].Container.Capacity, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerMinimumVolumeQuantity(numTypeTested int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return QuantityInput("typeTested["+strconv.Itoa(numTypeTested)+"].container.minimumVolumeQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("typeTested["+strconv.Itoa(numTypeTested)+"].container.minimumVolumeQuantity", resource.TypeTested[numTypeTested].Container.MinimumVolumeQuantity, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerMinimumVolumeString(numTypeTested int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return StringInput("typeTested["+strconv.Itoa(numTypeTested)+"].container.minimumVolumeString", nil, htmlAttrs)
	}
	return StringInput("typeTested["+strconv.Itoa(numTypeTested)+"].container.minimumVolumeString", resource.TypeTested[numTypeTested].Container.MinimumVolumeString, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerPreparation(numTypeTested int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) {
		return StringInput("typeTested["+strconv.Itoa(numTypeTested)+"].container.preparation", nil, htmlAttrs)
	}
	return StringInput("typeTested["+strconv.Itoa(numTypeTested)+"].container.preparation", resource.TypeTested[numTypeTested].Container.Preparation, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerAdditiveAdditiveCodeableConcept(numTypeTested int, numAdditive int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) || numAdditive >= len(resource.TypeTested[numTypeTested].Container.Additive) {
		return CodeableConceptSelect("typeTested["+strconv.Itoa(numTypeTested)+"].container.additive["+strconv.Itoa(numAdditive)+"].additiveCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("typeTested["+strconv.Itoa(numTypeTested)+"].container.additive["+strconv.Itoa(numAdditive)+"].additiveCodeableConcept", &resource.TypeTested[numTypeTested].Container.Additive[numAdditive].AdditiveCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedContainerAdditiveAdditiveReference(frs []FhirResource, numTypeTested int, numAdditive int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) || numAdditive >= len(resource.TypeTested[numTypeTested].Container.Additive) {
		return ReferenceInput(frs, "typeTested["+strconv.Itoa(numTypeTested)+"].container.additive["+strconv.Itoa(numAdditive)+"].additiveReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "typeTested["+strconv.Itoa(numTypeTested)+"].container.additive["+strconv.Itoa(numAdditive)+"].additiveReference", &resource.TypeTested[numTypeTested].Container.Additive[numAdditive].AdditiveReference, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedHandlingTemperatureQualifier(numTypeTested int, numHandling int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) || numHandling >= len(resource.TypeTested[numTypeTested].Handling) {
		return CodeableConceptSelect("typeTested["+strconv.Itoa(numTypeTested)+"].handling["+strconv.Itoa(numHandling)+"].temperatureQualifier", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("typeTested["+strconv.Itoa(numTypeTested)+"].handling["+strconv.Itoa(numHandling)+"].temperatureQualifier", resource.TypeTested[numTypeTested].Handling[numHandling].TemperatureQualifier, optionsValueSet, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedHandlingTemperatureRange(numTypeTested int, numHandling int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) || numHandling >= len(resource.TypeTested[numTypeTested].Handling) {
		return RangeInput("typeTested["+strconv.Itoa(numTypeTested)+"].handling["+strconv.Itoa(numHandling)+"].temperatureRange", nil, htmlAttrs)
	}
	return RangeInput("typeTested["+strconv.Itoa(numTypeTested)+"].handling["+strconv.Itoa(numHandling)+"].temperatureRange", resource.TypeTested[numTypeTested].Handling[numHandling].TemperatureRange, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedHandlingMaxDuration(numTypeTested int, numHandling int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) || numHandling >= len(resource.TypeTested[numTypeTested].Handling) {
		return DurationInput("typeTested["+strconv.Itoa(numTypeTested)+"].handling["+strconv.Itoa(numHandling)+"].maxDuration", nil, htmlAttrs)
	}
	return DurationInput("typeTested["+strconv.Itoa(numTypeTested)+"].handling["+strconv.Itoa(numHandling)+"].maxDuration", resource.TypeTested[numTypeTested].Handling[numHandling].MaxDuration, htmlAttrs)
}
func (resource *SpecimenDefinition) T_TypeTestedHandlingInstruction(numTypeTested int, numHandling int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTypeTested >= len(resource.TypeTested) || numHandling >= len(resource.TypeTested[numTypeTested].Handling) {
		return StringInput("typeTested["+strconv.Itoa(numTypeTested)+"].handling["+strconv.Itoa(numHandling)+"].instruction", nil, htmlAttrs)
	}
	return StringInput("typeTested["+strconv.Itoa(numTypeTested)+"].handling["+strconv.Itoa(numHandling)+"].instruction", resource.TypeTested[numTypeTested].Handling[numHandling].Instruction, htmlAttrs)
}
