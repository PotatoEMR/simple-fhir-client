//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4

import "encoding/json"

// http://hl7.org/fhir/r4/StructureDefinition/ObservationDefinition
type ObservationDefinition struct {
	Id                     *string                                   `json:"id,omitempty"`
	Meta                   *Meta                                     `json:"meta,omitempty"`
	ImplicitRules          *string                                   `json:"implicitRules,omitempty"`
	Language               *string                                   `json:"language,omitempty"`
	Text                   *Narrative                                `json:"text,omitempty"`
	Contained              []Resource                                `json:"contained,omitempty"`
	Extension              []Extension                               `json:"extension,omitempty"`
	ModifierExtension      []Extension                               `json:"modifierExtension,omitempty"`
	Category               []CodeableConcept                         `json:"category,omitempty"`
	Code                   CodeableConcept                           `json:"code"`
	Identifier             []Identifier                              `json:"identifier,omitempty"`
	PermittedDataType      []string                                  `json:"permittedDataType,omitempty"`
	MultipleResultsAllowed *bool                                     `json:"multipleResultsAllowed,omitempty"`
	Method                 *CodeableConcept                          `json:"method,omitempty"`
	PreferredReportName    *string                                   `json:"preferredReportName,omitempty"`
	QuantitativeDetails    *ObservationDefinitionQuantitativeDetails `json:"quantitativeDetails,omitempty"`
	QualifiedInterval      []ObservationDefinitionQualifiedInterval  `json:"qualifiedInterval,omitempty"`
	ValidCodedValueSet     *Reference                                `json:"validCodedValueSet,omitempty"`
	NormalCodedValueSet    *Reference                                `json:"normalCodedValueSet,omitempty"`
	AbnormalCodedValueSet  *Reference                                `json:"abnormalCodedValueSet,omitempty"`
	CriticalCodedValueSet  *Reference                                `json:"criticalCodedValueSet,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ObservationDefinition
type ObservationDefinitionQuantitativeDetails struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	CustomaryUnit     *CodeableConcept `json:"customaryUnit,omitempty"`
	Unit              *CodeableConcept `json:"unit,omitempty"`
	ConversionFactor  *float64         `json:"conversionFactor,omitempty"`
	DecimalPrecision  *int             `json:"decimalPrecision,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ObservationDefinition
type ObservationDefinitionQualifiedInterval struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Category          *string           `json:"category,omitempty"`
	Range             *Range            `json:"range,omitempty"`
	Context           *CodeableConcept  `json:"context,omitempty"`
	AppliesTo         []CodeableConcept `json:"appliesTo,omitempty"`
	Gender            *string           `json:"gender,omitempty"`
	Age               *Range            `json:"age,omitempty"`
	GestationalAge    *Range            `json:"gestationalAge,omitempty"`
	Condition         *string           `json:"condition,omitempty"`
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
