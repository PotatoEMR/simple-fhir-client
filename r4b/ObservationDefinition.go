package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/ObservationDefinition
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

// http://hl7.org/fhir/r4b/StructureDefinition/ObservationDefinition
type ObservationDefinitionQuantitativeDetails struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	CustomaryUnit     *CodeableConcept `json:"customaryUnit,omitempty"`
	Unit              *CodeableConcept `json:"unit,omitempty"`
	ConversionFactor  *float64         `json:"conversionFactor,omitempty"`
	DecimalPrecision  *int             `json:"decimalPrecision,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ObservationDefinition
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

func (resource *ObservationDefinition) ObservationDefinitionLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ObservationDefinition) ObservationDefinitionCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *ObservationDefinition) ObservationDefinitionCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Code, optionsValueSet)
}
func (resource *ObservationDefinition) ObservationDefinitionPermittedDataType() templ.Component {
	optionsValueSet := VSPermitted_data_type

	if resource != nil {
		return CodeSelect("permittedDataType", nil, optionsValueSet)
	}
	return CodeSelect("permittedDataType", &resource.PermittedDataType[0], optionsValueSet)
}
func (resource *ObservationDefinition) ObservationDefinitionMethod(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("method", resource.Method, optionsValueSet)
}
func (resource *ObservationDefinition) ObservationDefinitionQuantitativeDetailsCustomaryUnit(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("customaryUnit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("customaryUnit", resource.QuantitativeDetails.CustomaryUnit, optionsValueSet)
}
func (resource *ObservationDefinition) ObservationDefinitionQuantitativeDetailsUnit(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("unit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("unit", resource.QuantitativeDetails.Unit, optionsValueSet)
}
func (resource *ObservationDefinition) ObservationDefinitionQualifiedIntervalCategory(numQualifiedInterval int) templ.Component {
	optionsValueSet := VSObservation_range_category

	if resource != nil && len(resource.QualifiedInterval) >= numQualifiedInterval {
		return CodeSelect("category", nil, optionsValueSet)
	}
	return CodeSelect("category", resource.QualifiedInterval[numQualifiedInterval].Category, optionsValueSet)
}
func (resource *ObservationDefinition) ObservationDefinitionQualifiedIntervalContext(numQualifiedInterval int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.QualifiedInterval) >= numQualifiedInterval {
		return CodeableConceptSelect("context", nil, optionsValueSet)
	}
	return CodeableConceptSelect("context", resource.QualifiedInterval[numQualifiedInterval].Context, optionsValueSet)
}
func (resource *ObservationDefinition) ObservationDefinitionQualifiedIntervalAppliesTo(numQualifiedInterval int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.QualifiedInterval) >= numQualifiedInterval {
		return CodeableConceptSelect("appliesTo", nil, optionsValueSet)
	}
	return CodeableConceptSelect("appliesTo", &resource.QualifiedInterval[numQualifiedInterval].AppliesTo[0], optionsValueSet)
}
func (resource *ObservationDefinition) ObservationDefinitionQualifiedIntervalGender(numQualifiedInterval int) templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource != nil && len(resource.QualifiedInterval) >= numQualifiedInterval {
		return CodeSelect("gender", nil, optionsValueSet)
	}
	return CodeSelect("gender", resource.QualifiedInterval[numQualifiedInterval].Gender, optionsValueSet)
}
