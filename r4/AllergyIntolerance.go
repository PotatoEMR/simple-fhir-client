package r4

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/AllergyIntolerance
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

// http://hl7.org/fhir/r4/StructureDefinition/AllergyIntolerance
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
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *AllergyIntolerance) AllergyIntoleranceType() templ.Component {
	optionsValueSet := VSAllergy_intolerance_type
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
func (resource *AllergyIntolerance) AllergyIntoleranceCategory() templ.Component {
	optionsValueSet := VSAllergy_intolerance_category
	currentVal := ""
	if resource != nil {
		currentVal = resource.Category[0]
	}
	return CodeSelect("category", currentVal, optionsValueSet)
}
func (resource *AllergyIntolerance) AllergyIntoleranceCriticality() templ.Component {
	optionsValueSet := VSAllergy_intolerance_criticality
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Criticality
	}
	return CodeSelect("criticality", currentVal, optionsValueSet)
}
func (resource *AllergyIntolerance) AllergyIntoleranceReactionSeverity(numReaction int) templ.Component {
	optionsValueSet := VSReaction_event_severity
	currentVal := ""
	if resource != nil && len(resource.Reaction) >= numReaction {
		currentVal = *resource.Reaction[numReaction].Severity
	}
	return CodeSelect("severity", currentVal, optionsValueSet)
}
