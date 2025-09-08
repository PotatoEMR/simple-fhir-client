package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Date              time.Time                                  `json:"date,format:'2006-01-02T15:04:05Z07:00'"`
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
	Value             time.Time       `json:"value,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *ImmunizationRecommendation) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", &resource.Date, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationVaccineCode(numRecommendation int, numVaccineCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numVaccineCode >= len(resource.Recommendation[numRecommendation].VaccineCode) {
		return CodeableConceptSelect("Recommendation["+strconv.Itoa(numRecommendation)+"]VaccineCode["+strconv.Itoa(numVaccineCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Recommendation["+strconv.Itoa(numRecommendation)+"]VaccineCode["+strconv.Itoa(numVaccineCode)+"]", &resource.Recommendation[numRecommendation].VaccineCode[numVaccineCode], optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationTargetDisease(numRecommendation int, numTargetDisease int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numTargetDisease >= len(resource.Recommendation[numRecommendation].TargetDisease) {
		return CodeableConceptSelect("Recommendation["+strconv.Itoa(numRecommendation)+"]TargetDisease["+strconv.Itoa(numTargetDisease)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Recommendation["+strconv.Itoa(numRecommendation)+"]TargetDisease["+strconv.Itoa(numTargetDisease)+"]", &resource.Recommendation[numRecommendation].TargetDisease[numTargetDisease], optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationContraindicatedVaccineCode(numRecommendation int, numContraindicatedVaccineCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numContraindicatedVaccineCode >= len(resource.Recommendation[numRecommendation].ContraindicatedVaccineCode) {
		return CodeableConceptSelect("Recommendation["+strconv.Itoa(numRecommendation)+"]ContraindicatedVaccineCode["+strconv.Itoa(numContraindicatedVaccineCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Recommendation["+strconv.Itoa(numRecommendation)+"]ContraindicatedVaccineCode["+strconv.Itoa(numContraindicatedVaccineCode)+"]", &resource.Recommendation[numRecommendation].ContraindicatedVaccineCode[numContraindicatedVaccineCode], optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationForecastStatus(numRecommendation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) {
		return CodeableConceptSelect("Recommendation["+strconv.Itoa(numRecommendation)+"]ForecastStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Recommendation["+strconv.Itoa(numRecommendation)+"]ForecastStatus", &resource.Recommendation[numRecommendation].ForecastStatus, optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationForecastReason(numRecommendation int, numForecastReason int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numForecastReason >= len(resource.Recommendation[numRecommendation].ForecastReason) {
		return CodeableConceptSelect("Recommendation["+strconv.Itoa(numRecommendation)+"]ForecastReason["+strconv.Itoa(numForecastReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Recommendation["+strconv.Itoa(numRecommendation)+"]ForecastReason["+strconv.Itoa(numForecastReason)+"]", &resource.Recommendation[numRecommendation].ForecastReason[numForecastReason], optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationDescription(numRecommendation int, htmlAttrs string) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) {
		return StringInput("Recommendation["+strconv.Itoa(numRecommendation)+"]Description", nil, htmlAttrs)
	}
	return StringInput("Recommendation["+strconv.Itoa(numRecommendation)+"]Description", resource.Recommendation[numRecommendation].Description, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationSeries(numRecommendation int, htmlAttrs string) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) {
		return StringInput("Recommendation["+strconv.Itoa(numRecommendation)+"]Series", nil, htmlAttrs)
	}
	return StringInput("Recommendation["+strconv.Itoa(numRecommendation)+"]Series", resource.Recommendation[numRecommendation].Series, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationDoseNumber(numRecommendation int, htmlAttrs string) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) {
		return StringInput("Recommendation["+strconv.Itoa(numRecommendation)+"]DoseNumber", nil, htmlAttrs)
	}
	return StringInput("Recommendation["+strconv.Itoa(numRecommendation)+"]DoseNumber", resource.Recommendation[numRecommendation].DoseNumber, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationSeriesDoses(numRecommendation int, htmlAttrs string) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) {
		return StringInput("Recommendation["+strconv.Itoa(numRecommendation)+"]SeriesDoses", nil, htmlAttrs)
	}
	return StringInput("Recommendation["+strconv.Itoa(numRecommendation)+"]SeriesDoses", resource.Recommendation[numRecommendation].SeriesDoses, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationDateCriterionCode(numRecommendation int, numDateCriterion int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numDateCriterion >= len(resource.Recommendation[numRecommendation].DateCriterion) {
		return CodeableConceptSelect("Recommendation["+strconv.Itoa(numRecommendation)+"]DateCriterion["+strconv.Itoa(numDateCriterion)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Recommendation["+strconv.Itoa(numRecommendation)+"]DateCriterion["+strconv.Itoa(numDateCriterion)+"].Code", &resource.Recommendation[numRecommendation].DateCriterion[numDateCriterion].Code, optionsValueSet, htmlAttrs)
}
func (resource *ImmunizationRecommendation) T_RecommendationDateCriterionValue(numRecommendation int, numDateCriterion int, htmlAttrs string) templ.Component {
	if resource == nil || numRecommendation >= len(resource.Recommendation) || numDateCriterion >= len(resource.Recommendation[numRecommendation].DateCriterion) {
		return DateTimeInput("Recommendation["+strconv.Itoa(numRecommendation)+"]DateCriterion["+strconv.Itoa(numDateCriterion)+"].Value", nil, htmlAttrs)
	}
	return DateTimeInput("Recommendation["+strconv.Itoa(numRecommendation)+"]DateCriterion["+strconv.Itoa(numDateCriterion)+"].Value", &resource.Recommendation[numRecommendation].DateCriterion[numDateCriterion].Value, htmlAttrs)
}
