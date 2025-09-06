package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ImmunizationRecommendation
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

// http://hl7.org/fhir/r5/StructureDefinition/ImmunizationRecommendation
type ImmunizationRecommendationRecommendation struct {
	Id                           *string                                                 `json:"id,omitempty"`
	Extension                    []Extension                                             `json:"extension,omitempty"`
	ModifierExtension            []Extension                                             `json:"modifierExtension,omitempty"`
	VaccineCode                  []CodeableConcept                                       `json:"vaccineCode,omitempty"`
	TargetDisease                []CodeableConcept                                       `json:"targetDisease,omitempty"`
	ContraindicatedVaccineCode   []CodeableConcept                                       `json:"contraindicatedVaccineCode,omitempty"`
	ForecastStatus               CodeableConcept                                         `json:"forecastStatus"`
	ForecastReason               []CodeableConcept                                       `json:"forecastReason,omitempty"`
	DateCriterion                []ImmunizationRecommendationRecommendationDateCriterion `json:"dateCriterion,omitempty"`
	Description                  *string                                                 `json:"description,omitempty"`
	Series                       *string                                                 `json:"series,omitempty"`
	DoseNumber                   *string                                                 `json:"doseNumber,omitempty"`
	SeriesDoses                  *string                                                 `json:"seriesDoses,omitempty"`
	SupportingImmunization       []Reference                                             `json:"supportingImmunization,omitempty"`
	SupportingPatientInformation []Reference                                             `json:"supportingPatientInformation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ImmunizationRecommendation
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

func (resource *ImmunizationRecommendation) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ImmunizationRecommendation.Id", nil)
	}
	return StringInput("ImmunizationRecommendation.Id", resource.Id)
}
func (resource *ImmunizationRecommendation) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ImmunizationRecommendation.ImplicitRules", nil)
	}
	return StringInput("ImmunizationRecommendation.ImplicitRules", resource.ImplicitRules)
}
func (resource *ImmunizationRecommendation) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ImmunizationRecommendation.Language", nil, optionsValueSet)
	}
	return CodeSelect("ImmunizationRecommendation.Language", resource.Language, optionsValueSet)
}
func (resource *ImmunizationRecommendation) T_Date() templ.Component {

	if resource == nil {
		return StringInput("ImmunizationRecommendation.Date", nil)
	}
	return StringInput("ImmunizationRecommendation.Date", &resource.Date)
}
func (resource *ImmunizationRecommendation) T_RecommendationId(numRecommendation int) templ.Component {

	if resource == nil || len(resource.Recommendation) >= numRecommendation {
		return StringInput("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].Id", nil)
	}
	return StringInput("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].Id", resource.Recommendation[numRecommendation].Id)
}
func (resource *ImmunizationRecommendation) T_RecommendationVaccineCode(numRecommendation int, numVaccineCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Recommendation) >= numRecommendation || len(resource.Recommendation[numRecommendation].VaccineCode) >= numVaccineCode {
		return CodeableConceptSelect("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].VaccineCode["+strconv.Itoa(numVaccineCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].VaccineCode["+strconv.Itoa(numVaccineCode)+"]", &resource.Recommendation[numRecommendation].VaccineCode[numVaccineCode], optionsValueSet)
}
func (resource *ImmunizationRecommendation) T_RecommendationTargetDisease(numRecommendation int, numTargetDisease int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Recommendation) >= numRecommendation || len(resource.Recommendation[numRecommendation].TargetDisease) >= numTargetDisease {
		return CodeableConceptSelect("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].TargetDisease["+strconv.Itoa(numTargetDisease)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].TargetDisease["+strconv.Itoa(numTargetDisease)+"]", &resource.Recommendation[numRecommendation].TargetDisease[numTargetDisease], optionsValueSet)
}
func (resource *ImmunizationRecommendation) T_RecommendationContraindicatedVaccineCode(numRecommendation int, numContraindicatedVaccineCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Recommendation) >= numRecommendation || len(resource.Recommendation[numRecommendation].ContraindicatedVaccineCode) >= numContraindicatedVaccineCode {
		return CodeableConceptSelect("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].ContraindicatedVaccineCode["+strconv.Itoa(numContraindicatedVaccineCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].ContraindicatedVaccineCode["+strconv.Itoa(numContraindicatedVaccineCode)+"]", &resource.Recommendation[numRecommendation].ContraindicatedVaccineCode[numContraindicatedVaccineCode], optionsValueSet)
}
func (resource *ImmunizationRecommendation) T_RecommendationForecastStatus(numRecommendation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Recommendation) >= numRecommendation {
		return CodeableConceptSelect("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].ForecastStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].ForecastStatus", &resource.Recommendation[numRecommendation].ForecastStatus, optionsValueSet)
}
func (resource *ImmunizationRecommendation) T_RecommendationForecastReason(numRecommendation int, numForecastReason int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Recommendation) >= numRecommendation || len(resource.Recommendation[numRecommendation].ForecastReason) >= numForecastReason {
		return CodeableConceptSelect("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].ForecastReason["+strconv.Itoa(numForecastReason)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].ForecastReason["+strconv.Itoa(numForecastReason)+"]", &resource.Recommendation[numRecommendation].ForecastReason[numForecastReason], optionsValueSet)
}
func (resource *ImmunizationRecommendation) T_RecommendationDescription(numRecommendation int) templ.Component {

	if resource == nil || len(resource.Recommendation) >= numRecommendation {
		return StringInput("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].Description", nil)
	}
	return StringInput("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].Description", resource.Recommendation[numRecommendation].Description)
}
func (resource *ImmunizationRecommendation) T_RecommendationSeries(numRecommendation int) templ.Component {

	if resource == nil || len(resource.Recommendation) >= numRecommendation {
		return StringInput("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].Series", nil)
	}
	return StringInput("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].Series", resource.Recommendation[numRecommendation].Series)
}
func (resource *ImmunizationRecommendation) T_RecommendationDoseNumber(numRecommendation int) templ.Component {

	if resource == nil || len(resource.Recommendation) >= numRecommendation {
		return StringInput("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].DoseNumber", nil)
	}
	return StringInput("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].DoseNumber", resource.Recommendation[numRecommendation].DoseNumber)
}
func (resource *ImmunizationRecommendation) T_RecommendationSeriesDoses(numRecommendation int) templ.Component {

	if resource == nil || len(resource.Recommendation) >= numRecommendation {
		return StringInput("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].SeriesDoses", nil)
	}
	return StringInput("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].SeriesDoses", resource.Recommendation[numRecommendation].SeriesDoses)
}
func (resource *ImmunizationRecommendation) T_RecommendationDateCriterionId(numRecommendation int, numDateCriterion int) templ.Component {

	if resource == nil || len(resource.Recommendation) >= numRecommendation || len(resource.Recommendation[numRecommendation].DateCriterion) >= numDateCriterion {
		return StringInput("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].DateCriterion["+strconv.Itoa(numDateCriterion)+"].Id", nil)
	}
	return StringInput("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].DateCriterion["+strconv.Itoa(numDateCriterion)+"].Id", resource.Recommendation[numRecommendation].DateCriterion[numDateCriterion].Id)
}
func (resource *ImmunizationRecommendation) T_RecommendationDateCriterionCode(numRecommendation int, numDateCriterion int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Recommendation) >= numRecommendation || len(resource.Recommendation[numRecommendation].DateCriterion) >= numDateCriterion {
		return CodeableConceptSelect("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].DateCriterion["+strconv.Itoa(numDateCriterion)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].DateCriterion["+strconv.Itoa(numDateCriterion)+"].Code", &resource.Recommendation[numRecommendation].DateCriterion[numDateCriterion].Code, optionsValueSet)
}
func (resource *ImmunizationRecommendation) T_RecommendationDateCriterionValue(numRecommendation int, numDateCriterion int) templ.Component {

	if resource == nil || len(resource.Recommendation) >= numRecommendation || len(resource.Recommendation[numRecommendation].DateCriterion) >= numDateCriterion {
		return StringInput("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].DateCriterion["+strconv.Itoa(numDateCriterion)+"].Value", nil)
	}
	return StringInput("ImmunizationRecommendation.Recommendation["+strconv.Itoa(numRecommendation)+"].DateCriterion["+strconv.Itoa(numDateCriterion)+"].Value", &resource.Recommendation[numRecommendation].DateCriterion[numDateCriterion].Value)
}
