//generated August 15 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/ObservationDefinition
type ObservationDefinition struct {
	Id                     *string                               `json:"id,omitempty"`
	Meta                   *Meta                                 `json:"meta,omitempty"`
	ImplicitRules          *string                               `json:"implicitRules,omitempty"`
	Language               *string                               `json:"language,omitempty"`
	Text                   *Narrative                            `json:"text,omitempty"`
	Contained              []Resource                            `json:"contained,omitempty"`
	Extension              []Extension                           `json:"extension,omitempty"`
	ModifierExtension      []Extension                           `json:"modifierExtension,omitempty"`
	Url                    *string                               `json:"url,omitempty"`
	Identifier             *Identifier                           `json:"identifier,omitempty"`
	Version                *string                               `json:"version,omitempty"`
	VersionAlgorithmString *string                               `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                               `json:"versionAlgorithmCoding"`
	Name                   *string                               `json:"name,omitempty"`
	Title                  *string                               `json:"title,omitempty"`
	Status                 string                                `json:"status"`
	Experimental           *bool                                 `json:"experimental,omitempty"`
	Date                   *string                               `json:"date,omitempty"`
	Publisher              *string                               `json:"publisher,omitempty"`
	Contact                []ContactDetail                       `json:"contact,omitempty"`
	Description            *string                               `json:"description,omitempty"`
	UseContext             []UsageContext                        `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                     `json:"jurisdiction,omitempty"`
	Purpose                *string                               `json:"purpose,omitempty"`
	Copyright              *string                               `json:"copyright,omitempty"`
	CopyrightLabel         *string                               `json:"copyrightLabel,omitempty"`
	ApprovalDate           *string                               `json:"approvalDate,omitempty"`
	LastReviewDate         *string                               `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                               `json:"effectivePeriod,omitempty"`
	DerivedFromCanonical   []string                              `json:"derivedFromCanonical,omitempty"`
	DerivedFromUri         []string                              `json:"derivedFromUri,omitempty"`
	Subject                []CodeableConcept                     `json:"subject,omitempty"`
	PerformerType          *CodeableConcept                      `json:"performerType,omitempty"`
	Category               []CodeableConcept                     `json:"category,omitempty"`
	Code                   CodeableConcept                       `json:"code"`
	PermittedDataType      []string                              `json:"permittedDataType,omitempty"`
	MultipleResultsAllowed *bool                                 `json:"multipleResultsAllowed,omitempty"`
	BodySite               *CodeableConcept                      `json:"bodySite,omitempty"`
	Method                 *CodeableConcept                      `json:"method,omitempty"`
	Specimen               []Reference                           `json:"specimen,omitempty"`
	Device                 []Reference                           `json:"device,omitempty"`
	PreferredReportName    *string                               `json:"preferredReportName,omitempty"`
	PermittedUnit          []Coding                              `json:"permittedUnit,omitempty"`
	QualifiedValue         []ObservationDefinitionQualifiedValue `json:"qualifiedValue,omitempty"`
	HasMember              []Reference                           `json:"hasMember,omitempty"`
	Component              []ObservationDefinitionComponent      `json:"component,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ObservationDefinition
type ObservationDefinitionQualifiedValue struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Context               *CodeableConcept  `json:"context,omitempty"`
	AppliesTo             []CodeableConcept `json:"appliesTo,omitempty"`
	Gender                *string           `json:"gender,omitempty"`
	Age                   *Range            `json:"age,omitempty"`
	GestationalAge        *Range            `json:"gestationalAge,omitempty"`
	Condition             *string           `json:"condition,omitempty"`
	RangeCategory         *string           `json:"rangeCategory,omitempty"`
	Range                 *Range            `json:"range,omitempty"`
	ValidCodedValueSet    *string           `json:"validCodedValueSet,omitempty"`
	NormalCodedValueSet   *string           `json:"normalCodedValueSet,omitempty"`
	AbnormalCodedValueSet *string           `json:"abnormalCodedValueSet,omitempty"`
	CriticalCodedValueSet *string           `json:"criticalCodedValueSet,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ObservationDefinition
type ObservationDefinitionComponent struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code"`
	PermittedDataType []string        `json:"permittedDataType,omitempty"`
	PermittedUnit     []Coding        `json:"permittedUnit,omitempty"`
}

type OtherObservationDefinition ObservationDefinition

// on convert struct to json, automatically add resourceType=ObservationDefinition
func (r ObservationDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherObservationDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherObservationDefinition: OtherObservationDefinition(r),
		ResourceType:               "ObservationDefinition",
	})
}
