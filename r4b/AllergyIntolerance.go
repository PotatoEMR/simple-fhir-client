package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/AllergyIntolerance
type AllergyIntolerance struct {
	Id                 *string                      `json:"id,omitempty"`
	Meta               *Meta                        `json:"meta,omitempty"`
	ImplicitRules      *string                      `json:"implicitRules,omitempty"`
	Language           *string                      `json:"language,omitempty"`
	Text               *Narrative                   `json:"text,omitempty"`
	Contained          []Resource                   `json:"contained,omitempty"`
	Extension          []Extension                  `json:"extension,omitempty"`
	ModifierExtension  []Extension                  `json:"modifierExtension,omitempty"`
	Identifier         []Identifier                 `json:"identifier,omitempty"`
	ClinicalStatus     *CodeableConcept             `json:"clinicalStatus,omitempty"`
	VerificationStatus *CodeableConcept             `json:"verificationStatus,omitempty"`
	Type               *string                      `json:"type,omitempty"`
	Category           []string                     `json:"category,omitempty"`
	Criticality        *string                      `json:"criticality,omitempty"`
	Code               *CodeableConcept             `json:"code,omitempty"`
	Patient            Reference                    `json:"patient"`
	Encounter          *Reference                   `json:"encounter,omitempty"`
	OnsetDateTime      *string                      `json:"onsetDateTime"`
	OnsetAge           *Age                         `json:"onsetAge"`
	OnsetPeriod        *Period                      `json:"onsetPeriod"`
	OnsetRange         *Range                       `json:"onsetRange"`
	OnsetString        *string                      `json:"onsetString"`
	RecordedDate       *string                      `json:"recordedDate,omitempty"`
	Recorder           *Reference                   `json:"recorder,omitempty"`
	Asserter           *Reference                   `json:"asserter,omitempty"`
	LastOccurrence     *string                      `json:"lastOccurrence,omitempty"`
	Note               []Annotation                 `json:"note,omitempty"`
	Reaction           []AllergyIntoleranceReaction `json:"reaction,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/AllergyIntolerance
type AllergyIntoleranceReaction struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Substance         *CodeableConcept  `json:"substance,omitempty"`
	Manifestation     []CodeableConcept `json:"manifestation"`
	Description       *string           `json:"description,omitempty"`
	Onset             *string           `json:"onset,omitempty"`
	Severity          *string           `json:"severity,omitempty"`
	ExposureRoute     *CodeableConcept  `json:"exposureRoute,omitempty"`
	Note              []Annotation      `json:"note,omitempty"`
}

type OtherAllergyIntolerance AllergyIntolerance

// on convert struct to json, automatically add resourceType=AllergyIntolerance
func (r AllergyIntolerance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAllergyIntolerance
		ResourceType string `json:"resourceType"`
	}{
		OtherAllergyIntolerance: OtherAllergyIntolerance(r),
		ResourceType:            "AllergyIntolerance",
	})
}

func (resource *AllergyIntolerance) AllergyIntoleranceLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *AllergyIntolerance) AllergyIntoleranceClinicalStatus() templ.Component {
	optionsValueSet := VSAllergyintolerance_clinical

	if resource == nil {
		return CodeableConceptSelect("clinicalStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("clinicalStatus", resource.ClinicalStatus, optionsValueSet)
}
func (resource *AllergyIntolerance) AllergyIntoleranceVerificationStatus() templ.Component {
	optionsValueSet := VSAllergyintolerance_verification

	if resource == nil {
		return CodeableConceptSelect("verificationStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("verificationStatus", resource.VerificationStatus, optionsValueSet)
}
func (resource *AllergyIntolerance) AllergyIntoleranceType() templ.Component {
	optionsValueSet := VSAllergy_intolerance_type

	if resource == nil {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", resource.Type, optionsValueSet)
}
func (resource *AllergyIntolerance) AllergyIntoleranceCategory() templ.Component {
	optionsValueSet := VSAllergy_intolerance_category

	if resource == nil {
		return CodeSelect("category", nil, optionsValueSet)
	}
	return CodeSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *AllergyIntolerance) AllergyIntoleranceCriticality() templ.Component {
	optionsValueSet := VSAllergy_intolerance_criticality

	if resource == nil {
		return CodeSelect("criticality", nil, optionsValueSet)
	}
	return CodeSelect("criticality", resource.Criticality, optionsValueSet)
}
func (resource *AllergyIntolerance) AllergyIntoleranceCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet)
}
func (resource *AllergyIntolerance) AllergyIntoleranceReactionSubstance(numReaction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Reaction) >= numReaction {
		return CodeableConceptSelect("substance", nil, optionsValueSet)
	}
	return CodeableConceptSelect("substance", resource.Reaction[numReaction].Substance, optionsValueSet)
}
func (resource *AllergyIntolerance) AllergyIntoleranceReactionManifestation(numReaction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Reaction) >= numReaction {
		return CodeableConceptSelect("manifestation", nil, optionsValueSet)
	}
	return CodeableConceptSelect("manifestation", &resource.Reaction[numReaction].Manifestation[0], optionsValueSet)
}
func (resource *AllergyIntolerance) AllergyIntoleranceReactionSeverity(numReaction int) templ.Component {
	optionsValueSet := VSReaction_event_severity

	if resource == nil && len(resource.Reaction) >= numReaction {
		return CodeSelect("severity", nil, optionsValueSet)
	}
	return CodeSelect("severity", resource.Reaction[numReaction].Severity, optionsValueSet)
}
func (resource *AllergyIntolerance) AllergyIntoleranceReactionExposureRoute(numReaction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Reaction) >= numReaction {
		return CodeableConceptSelect("exposureRoute", nil, optionsValueSet)
	}
	return CodeableConceptSelect("exposureRoute", resource.Reaction[numReaction].ExposureRoute, optionsValueSet)
}
