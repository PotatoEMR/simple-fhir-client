package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *ObservationDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.Id", nil)
	}
	return StringInput("ObservationDefinition.Id", resource.Id)
}
func (resource *ObservationDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.ImplicitRules", nil)
	}
	return StringInput("ObservationDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *ObservationDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ObservationDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("ObservationDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *ObservationDefinition) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("ObservationDefinition.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *ObservationDefinition) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ObservationDefinition.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.Code", &resource.Code, optionsValueSet)
}
func (resource *ObservationDefinition) T_PermittedDataType(numPermittedDataType int) templ.Component {
	optionsValueSet := VSPermitted_data_type

	if resource == nil || len(resource.PermittedDataType) >= numPermittedDataType {
		return CodeSelect("ObservationDefinition.PermittedDataType["+strconv.Itoa(numPermittedDataType)+"]", nil, optionsValueSet)
	}
	return CodeSelect("ObservationDefinition.PermittedDataType["+strconv.Itoa(numPermittedDataType)+"]", &resource.PermittedDataType[numPermittedDataType], optionsValueSet)
}
func (resource *ObservationDefinition) T_MultipleResultsAllowed() templ.Component {

	if resource == nil {
		return BoolInput("ObservationDefinition.MultipleResultsAllowed", nil)
	}
	return BoolInput("ObservationDefinition.MultipleResultsAllowed", resource.MultipleResultsAllowed)
}
func (resource *ObservationDefinition) T_Method(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ObservationDefinition.Method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.Method", resource.Method, optionsValueSet)
}
func (resource *ObservationDefinition) T_PreferredReportName() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.PreferredReportName", nil)
	}
	return StringInput("ObservationDefinition.PreferredReportName", resource.PreferredReportName)
}
func (resource *ObservationDefinition) T_QuantitativeDetailsId() templ.Component {

	if resource == nil {
		return StringInput("ObservationDefinition.QuantitativeDetails.Id", nil)
	}
	return StringInput("ObservationDefinition.QuantitativeDetails.Id", resource.QuantitativeDetails.Id)
}
func (resource *ObservationDefinition) T_QuantitativeDetailsCustomaryUnit(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ObservationDefinition.QuantitativeDetails.CustomaryUnit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.QuantitativeDetails.CustomaryUnit", resource.QuantitativeDetails.CustomaryUnit, optionsValueSet)
}
func (resource *ObservationDefinition) T_QuantitativeDetailsUnit(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ObservationDefinition.QuantitativeDetails.Unit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.QuantitativeDetails.Unit", resource.QuantitativeDetails.Unit, optionsValueSet)
}
func (resource *ObservationDefinition) T_QuantitativeDetailsConversionFactor() templ.Component {

	if resource == nil {
		return Float64Input("ObservationDefinition.QuantitativeDetails.ConversionFactor", nil)
	}
	return Float64Input("ObservationDefinition.QuantitativeDetails.ConversionFactor", resource.QuantitativeDetails.ConversionFactor)
}
func (resource *ObservationDefinition) T_QuantitativeDetailsDecimalPrecision() templ.Component {

	if resource == nil {
		return IntInput("ObservationDefinition.QuantitativeDetails.DecimalPrecision", nil)
	}
	return IntInput("ObservationDefinition.QuantitativeDetails.DecimalPrecision", resource.QuantitativeDetails.DecimalPrecision)
}
func (resource *ObservationDefinition) T_QualifiedIntervalId(numQualifiedInterval int) templ.Component {

	if resource == nil || len(resource.QualifiedInterval) >= numQualifiedInterval {
		return StringInput("ObservationDefinition.QualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].Id", nil)
	}
	return StringInput("ObservationDefinition.QualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].Id", resource.QualifiedInterval[numQualifiedInterval].Id)
}
func (resource *ObservationDefinition) T_QualifiedIntervalCategory(numQualifiedInterval int) templ.Component {
	optionsValueSet := VSObservation_range_category

	if resource == nil || len(resource.QualifiedInterval) >= numQualifiedInterval {
		return CodeSelect("ObservationDefinition.QualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].Category", nil, optionsValueSet)
	}
	return CodeSelect("ObservationDefinition.QualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].Category", resource.QualifiedInterval[numQualifiedInterval].Category, optionsValueSet)
}
func (resource *ObservationDefinition) T_QualifiedIntervalContext(numQualifiedInterval int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.QualifiedInterval) >= numQualifiedInterval {
		return CodeableConceptSelect("ObservationDefinition.QualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].Context", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.QualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].Context", resource.QualifiedInterval[numQualifiedInterval].Context, optionsValueSet)
}
func (resource *ObservationDefinition) T_QualifiedIntervalAppliesTo(numQualifiedInterval int, numAppliesTo int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.QualifiedInterval) >= numQualifiedInterval || len(resource.QualifiedInterval[numQualifiedInterval].AppliesTo) >= numAppliesTo {
		return CodeableConceptSelect("ObservationDefinition.QualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].AppliesTo["+strconv.Itoa(numAppliesTo)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ObservationDefinition.QualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].AppliesTo["+strconv.Itoa(numAppliesTo)+"]", &resource.QualifiedInterval[numQualifiedInterval].AppliesTo[numAppliesTo], optionsValueSet)
}
func (resource *ObservationDefinition) T_QualifiedIntervalGender(numQualifiedInterval int) templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource == nil || len(resource.QualifiedInterval) >= numQualifiedInterval {
		return CodeSelect("ObservationDefinition.QualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].Gender", nil, optionsValueSet)
	}
	return CodeSelect("ObservationDefinition.QualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].Gender", resource.QualifiedInterval[numQualifiedInterval].Gender, optionsValueSet)
}
func (resource *ObservationDefinition) T_QualifiedIntervalCondition(numQualifiedInterval int) templ.Component {

	if resource == nil || len(resource.QualifiedInterval) >= numQualifiedInterval {
		return StringInput("ObservationDefinition.QualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].Condition", nil)
	}
	return StringInput("ObservationDefinition.QualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].Condition", resource.QualifiedInterval[numQualifiedInterval].Condition)
}
