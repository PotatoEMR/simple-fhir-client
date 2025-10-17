package r5

//generated with command go run ./bultaoreune
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
	Date              FhirDateTime                               `json:"date"`
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
	Value             FhirDateTime    `json:"value"`
}

type OtherImmunizationRecommendation ImmunizationRecommendation

// struct -> json, automatically add resourceType=Patient
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
func (r ImmunizationRecommendation) ResourceType() string {
	return "ImmunizationRecommendation"
}

func (resource *ImmunizationRecommendation) T_Patient(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "patient", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "patient", &resource.Patient, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", &resource.Date, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_Authority(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "authority", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "authority", resource.Authority, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationVaccineCode(numRecommendation int, numVaccineCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numVaccineCode >= len(resource.Recommendation[numRecommendation].VaccineCode) {
		return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].vaccineCode["+strconv.Itoa(numVaccineCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].vaccineCode["+strconv.Itoa(numVaccineCode)+"]", &resource.Recommendation[numRecommendation].VaccineCode[numVaccineCode], optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationTargetDisease(numRecommendation int, numTargetDisease int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numTargetDisease >= len(resource.Recommendation[numRecommendation].TargetDisease) {
		return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].targetDisease["+strconv.Itoa(numTargetDisease)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].targetDisease["+strconv.Itoa(numTargetDisease)+"]", &resource.Recommendation[numRecommendation].TargetDisease[numTargetDisease], optionsValueSet, htmlAttrs)
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
func (resource *ImmunizationRecommendation) T_RecommendationDoseNumber(numRecommendation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) {
		return StringInput("recommendation["+strconv.Itoa(numRecommendation)+"].doseNumber", nil, htmlAttrs)
	}
	return StringInput("recommendation["+strconv.Itoa(numRecommendation)+"].doseNumber", resource.Recommendation[numRecommendation].DoseNumber, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationSeriesDoses(numRecommendation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) {
		return StringInput("recommendation["+strconv.Itoa(numRecommendation)+"].seriesDoses", nil, htmlAttrs)
	}
	return StringInput("recommendation["+strconv.Itoa(numRecommendation)+"].seriesDoses", resource.Recommendation[numRecommendation].SeriesDoses, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationSupportingImmunization(frs []FhirResource, numRecommendation int, numSupportingImmunization int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numSupportingImmunization >= len(resource.Recommendation[numRecommendation].SupportingImmunization) {
		return ReferenceInput(frs, "recommendation["+strconv.Itoa(numRecommendation)+"].supportingImmunization["+strconv.Itoa(numSupportingImmunization)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "recommendation["+strconv.Itoa(numRecommendation)+"].supportingImmunization["+strconv.Itoa(numSupportingImmunization)+"]", &resource.Recommendation[numRecommendation].SupportingImmunization[numSupportingImmunization], htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationSupportingPatientInformation(frs []FhirResource, numRecommendation int, numSupportingPatientInformation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numSupportingPatientInformation >= len(resource.Recommendation[numRecommendation].SupportingPatientInformation) {
		return ReferenceInput(frs, "recommendation["+strconv.Itoa(numRecommendation)+"].supportingPatientInformation["+strconv.Itoa(numSupportingPatientInformation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "recommendation["+strconv.Itoa(numRecommendation)+"].supportingPatientInformation["+strconv.Itoa(numSupportingPatientInformation)+"]", &resource.Recommendation[numRecommendation].SupportingPatientInformation[numSupportingPatientInformation], htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationDateCriterionCode(numRecommendation int, numDateCriterion int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numDateCriterion >= len(resource.Recommendation[numRecommendation].DateCriterion) {
		return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].dateCriterion["+strconv.Itoa(numDateCriterion)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("recommendation["+strconv.Itoa(numRecommendation)+"].dateCriterion["+strconv.Itoa(numDateCriterion)+"].code", &resource.Recommendation[numRecommendation].DateCriterion[numDateCriterion].Code, optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationDateCriterionValue(numRecommendation int, numDateCriterion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numDateCriterion >= len(resource.Recommendation[numRecommendation].DateCriterion) {
		return FhirDateTimeInput("recommendation["+strconv.Itoa(numRecommendation)+"].dateCriterion["+strconv.Itoa(numDateCriterion)+"].value", nil, htmlAttrs)
	}
	return FhirDateTimeInput("recommendation["+strconv.Itoa(numRecommendation)+"].dateCriterion["+strconv.Itoa(numDateCriterion)+"].value", &resource.Recommendation[numRecommendation].DateCriterion[numDateCriterion].Value, htmlAttrs)
}
