package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/ImmunizationRecommendation
type ImmunizationRecommendation struct {
	Id                *string                                    `json:"id,omitempty"`
	Meta              *Meta                                      `json:"meta,omitempty"`
	ImplicitRules     *string                                    `json:"implicitRules,omitempty"`
	Language          *string                                    `json:"language,omitempty"`
	Text              *Narrative                                 `json:"text,omitempty"`
	Contained         []Resource                                 `json:"contained,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                               `json:"identifier,omitempty"`
	Patient           Reference                                  `json:"patient"`
	Date              string                                     `json:"date"`
	Authority         *Reference                                 `json:"authority,omitempty"`
	Recommendation    []ImmunizationRecommendationRecommendation `json:"recommendation"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ImmunizationRecommendation
type ImmunizationRecommendationRecommendation struct {
	Id                           *string                                                 `json:"id,omitempty"`
	Extension                    []Extension                                             `json:"extension,omitempty"`
	ModifierExtension            []Extension                                             `json:"modifierExtension,omitempty"`
	VaccineCode                  []CodeableConcept                                       `json:"vaccineCode,omitempty"`
	TargetDisease                *CodeableConcept                                        `json:"targetDisease,omitempty"`
	ContraindicatedVaccineCode   []CodeableConcept                                       `json:"contraindicatedVaccineCode,omitempty"`
	ForecastStatus               CodeableConcept                                         `json:"forecastStatus"`
	ForecastReason               []CodeableConcept                                       `json:"forecastReason,omitempty"`
	DateCriterion                []ImmunizationRecommendationRecommendationDateCriterion `json:"dateCriterion,omitempty"`
	Description                  *string                                                 `json:"description,omitempty"`
	Series                       *string                                                 `json:"series,omitempty"`
	DoseNumberPositiveInt        *int                                                    `json:"doseNumberPositiveInt,omitempty"`
	DoseNumberString             *string                                                 `json:"doseNumberString,omitempty"`
	SeriesDosesPositiveInt       *int                                                    `json:"seriesDosesPositiveInt,omitempty"`
	SeriesDosesString            *string                                                 `json:"seriesDosesString,omitempty"`
	SupportingImmunization       []Reference                                             `json:"supportingImmunization,omitempty"`
	SupportingPatientInformation []Reference                                             `json:"supportingPatientInformation,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ImmunizationRecommendation
type ImmunizationRecommendationRecommendationDateCriterion struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code"`
	Value             string          `json:"value"`
}

type OtherImmunizationRecommendation ImmunizationRecommendation

// on convert struct to json, automatically add resourceType=ImmunizationRecommendation
func (r ImmunizationRecommendation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImmunizationRecommendation
		ResourceType string `json:"resourceType"`
	}{
		OtherImmunizationRecommendation: OtherImmunizationRecommendation(r),
		ResourceType:                    "ImmunizationRecommendation",
	})
}
func (r ImmunizationRecommendation) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ImmunizationRecommendation/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ImmunizationRecommendation"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ImmunizationRecommendation) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", &resource.Date, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationVaccineCode(numRecommendation int, numVaccineCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numVaccineCode >= len(resource.Recommendation[numRecommendation].VaccineCode) {
		return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].vaccineCode["+strconv.Itoa(numVaccineCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].vaccineCode["+strconv.Itoa(numVaccineCode)+"]", &resource.Recommendation[numRecommendation].VaccineCode[numVaccineCode], optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationTargetDisease(numRecommendation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) {
		return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].targetDisease", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].targetDisease", resource.Recommendation[numRecommendation].TargetDisease, optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationContraindicatedVaccineCode(numRecommendation int, numContraindicatedVaccineCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numContraindicatedVaccineCode >= len(resource.Recommendation[numRecommendation].ContraindicatedVaccineCode) {
		return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].contraindicatedVaccineCode["+strconv.Itoa(numContraindicatedVaccineCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].contraindicatedVaccineCode["+strconv.Itoa(numContraindicatedVaccineCode)+"]", &resource.Recommendation[numRecommendation].ContraindicatedVaccineCode[numContraindicatedVaccineCode], optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationForecastStatus(numRecommendation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) {
		return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].forecastStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].forecastStatus", &resource.Recommendation[numRecommendation].ForecastStatus, optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationForecastReason(numRecommendation int, numForecastReason int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numForecastReason >= len(resource.Recommendation[numRecommendation].ForecastReason) {
		return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].forecastReason["+strconv.Itoa(numForecastReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].forecastReason["+strconv.Itoa(numForecastReason)+"]", &resource.Recommendation[numRecommendation].ForecastReason[numForecastReason], optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationDescription(numRecommendation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) {
		return StringInput("recommendation["+strconv.Itoa(numRecommendation)+"].description", nil, htmlAttrs)
	}
	return StringInput("recommendation["+strconv.Itoa(numRecommendation)+"].description", resource.Recommendation[numRecommendation].Description, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationSeries(numRecommendation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) {
		return StringInput("recommendation["+strconv.Itoa(numRecommendation)+"].series", nil, htmlAttrs)
	}
	return StringInput("recommendation["+strconv.Itoa(numRecommendation)+"].series", resource.Recommendation[numRecommendation].Series, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationDoseNumberPositiveInt(numRecommendation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) {
		return IntInput("recommendation["+strconv.Itoa(numRecommendation)+"].doseNumberPositiveInt", nil, htmlAttrs)
	}
	return IntInput("recommendation["+strconv.Itoa(numRecommendation)+"].doseNumberPositiveInt", resource.Recommendation[numRecommendation].DoseNumberPositiveInt, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationDoseNumberString(numRecommendation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) {
		return StringInput("recommendation["+strconv.Itoa(numRecommendation)+"].doseNumberString", nil, htmlAttrs)
	}
	return StringInput("recommendation["+strconv.Itoa(numRecommendation)+"].doseNumberString", resource.Recommendation[numRecommendation].DoseNumberString, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationSeriesDosesPositiveInt(numRecommendation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) {
		return IntInput("recommendation["+strconv.Itoa(numRecommendation)+"].seriesDosesPositiveInt", nil, htmlAttrs)
	}
	return IntInput("recommendation["+strconv.Itoa(numRecommendation)+"].seriesDosesPositiveInt", resource.Recommendation[numRecommendation].SeriesDosesPositiveInt, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationSeriesDosesString(numRecommendation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) {
		return StringInput("recommendation["+strconv.Itoa(numRecommendation)+"].seriesDosesString", nil, htmlAttrs)
	}
	return StringInput("recommendation["+strconv.Itoa(numRecommendation)+"].seriesDosesString", resource.Recommendation[numRecommendation].SeriesDosesString, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationDateCriterionCode(numRecommendation int, numDateCriterion int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numDateCriterion >= len(resource.Recommendation[numRecommendation].DateCriterion) {
		return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].dateCriterion["+strconv.Itoa(numDateCriterion)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].dateCriterion["+strconv.Itoa(numDateCriterion)+"].code", &resource.Recommendation[numRecommendation].DateCriterion[numDateCriterion].Code, optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationDateCriterionValue(numRecommendation int, numDateCriterion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numDateCriterion >= len(resource.Recommendation[numRecommendation].DateCriterion) {
		return DateTimeInput("recommendation["+strconv.Itoa(numRecommendation)+"].dateCriterion["+strconv.Itoa(numDateCriterion)+"].value", nil, htmlAttrs)
	}
	return DateTimeInput("recommendation["+strconv.Itoa(numRecommendation)+"].dateCriterion["+strconv.Itoa(numDateCriterion)+"].value", &resource.Recommendation[numRecommendation].DateCriterion[numDateCriterion].Value, htmlAttrs)
}
