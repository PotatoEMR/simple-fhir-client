package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
func (r ObservationDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ObservationDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ObservationDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ObservationDefinition) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_PermittedDataType(numPermittedDataType int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPermitted_data_type

	if resource == nil || numPermittedDataType >= len(resource.PermittedDataType) {
		return CodeSelect("permittedDataType["+strconv.Itoa(numPermittedDataType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("permittedDataType["+strconv.Itoa(numPermittedDataType)+"]", &resource.PermittedDataType[numPermittedDataType], optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_MultipleResultsAllowed(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("multipleResultsAllowed", nil, htmlAttrs)
	}
	return BoolInput("multipleResultsAllowed", resource.MultipleResultsAllowed, htmlAttrs)
}
func (resource *ObservationDefinition) T_Method(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("method", resource.Method, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_PreferredReportName(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("preferredReportName", nil, htmlAttrs)
	}
	return StringInput("preferredReportName", resource.PreferredReportName, htmlAttrs)
}
func (resource *ObservationDefinition) T_QuantitativeDetailsCustomaryUnit(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("quantitativeDetails.customaryUnit", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("quantitativeDetails.customaryUnit", resource.QuantitativeDetails.CustomaryUnit, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_QuantitativeDetailsUnit(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("quantitativeDetails.unit", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("quantitativeDetails.unit", resource.QuantitativeDetails.Unit, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_QuantitativeDetailsConversionFactor(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return Float64Input("quantitativeDetails.conversionFactor", nil, htmlAttrs)
	}
	return Float64Input("quantitativeDetails.conversionFactor", resource.QuantitativeDetails.ConversionFactor, htmlAttrs)
}
func (resource *ObservationDefinition) T_QuantitativeDetailsDecimalPrecision(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("quantitativeDetails.decimalPrecision", nil, htmlAttrs)
	}
	return IntInput("quantitativeDetails.decimalPrecision", resource.QuantitativeDetails.DecimalPrecision, htmlAttrs)
}
func (resource *ObservationDefinition) T_QualifiedIntervalCategory(numQualifiedInterval int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSObservation_range_category

	if resource == nil || numQualifiedInterval >= len(resource.QualifiedInterval) {
		return CodeSelect("qualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("qualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].category", resource.QualifiedInterval[numQualifiedInterval].Category, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_QualifiedIntervalContext(numQualifiedInterval int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numQualifiedInterval >= len(resource.QualifiedInterval) {
		return CodeableConceptSelect("qualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].context", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("qualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].context", resource.QualifiedInterval[numQualifiedInterval].Context, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_QualifiedIntervalAppliesTo(numQualifiedInterval int, numAppliesTo int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numQualifiedInterval >= len(resource.QualifiedInterval) || numAppliesTo >= len(resource.QualifiedInterval[numQualifiedInterval].AppliesTo) {
		return CodeableConceptSelect("qualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].appliesTo["+strconv.Itoa(numAppliesTo)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("qualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].appliesTo["+strconv.Itoa(numAppliesTo)+"]", &resource.QualifiedInterval[numQualifiedInterval].AppliesTo[numAppliesTo], optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_QualifiedIntervalGender(numQualifiedInterval int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSAdministrative_gender

	if resource == nil || numQualifiedInterval >= len(resource.QualifiedInterval) {
		return CodeSelect("qualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].gender", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("qualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].gender", resource.QualifiedInterval[numQualifiedInterval].Gender, optionsValueSet, htmlAttrs)
}
func (resource *ObservationDefinition) T_QualifiedIntervalCondition(numQualifiedInterval int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numQualifiedInterval >= len(resource.QualifiedInterval) {
		return StringInput("qualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].condition", nil, htmlAttrs)
	}
	return StringInput("qualifiedInterval["+strconv.Itoa(numQualifiedInterval)+"].condition", resource.QualifiedInterval[numQualifiedInterval].Condition, htmlAttrs)
}
