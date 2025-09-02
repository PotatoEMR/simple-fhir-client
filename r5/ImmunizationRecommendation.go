package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *ImmunizationRecommendation) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ImmunizationRecommendation) T_RecommendationVaccineCode(numRecommendation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Recommendation) >= numRecommendation {
		return CodeableConceptSelect("vaccineCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("vaccineCode", &resource.Recommendation[numRecommendation].VaccineCode[0], optionsValueSet)
}
func (resource *ImmunizationRecommendation) T_RecommendationTargetDisease(numRecommendation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Recommendation) >= numRecommendation {
		return CodeableConceptSelect("targetDisease", nil, optionsValueSet)
	}
	return CodeableConceptSelect("targetDisease", &resource.Recommendation[numRecommendation].TargetDisease[0], optionsValueSet)
}
func (resource *ImmunizationRecommendation) T_RecommendationContraindicatedVaccineCode(numRecommendation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Recommendation) >= numRecommendation {
		return CodeableConceptSelect("contraindicatedVaccineCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("contraindicatedVaccineCode", &resource.Recommendation[numRecommendation].ContraindicatedVaccineCode[0], optionsValueSet)
}
func (resource *ImmunizationRecommendation) T_RecommendationForecastStatus(numRecommendation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Recommendation) >= numRecommendation {
		return CodeableConceptSelect("forecastStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("forecastStatus", &resource.Recommendation[numRecommendation].ForecastStatus, optionsValueSet)
}
func (resource *ImmunizationRecommendation) T_RecommendationForecastReason(numRecommendation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Recommendation) >= numRecommendation {
		return CodeableConceptSelect("forecastReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("forecastReason", &resource.Recommendation[numRecommendation].ForecastReason[0], optionsValueSet)
}
func (resource *ImmunizationRecommendation) T_RecommendationDateCriterionCode(numRecommendation int, numDateCriterion int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Recommendation[numRecommendation].DateCriterion) >= numDateCriterion {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Recommendation[numRecommendation].DateCriterion[numDateCriterion].Code, optionsValueSet)
}
