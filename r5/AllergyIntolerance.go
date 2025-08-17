//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/AllergyIntolerance
type AllergyIntolerance struct {
	Id                 *string                         `json:"id,omitempty"`
	Meta               *Meta                           `json:"meta,omitempty"`
	ImplicitRules      *string                         `json:"implicitRules,omitempty"`
	Language           *string                         `json:"language,omitempty"`
	Text               *Narrative                      `json:"text,omitempty"`
	Contained          []Resource                      `json:"contained,omitempty"`
	Extension          []Extension                     `json:"extension,omitempty"`
	ModifierExtension  []Extension                     `json:"modifierExtension,omitempty"`
	Identifier         []Identifier                    `json:"identifier,omitempty"`
	ClinicalStatus     *CodeableConcept                `json:"clinicalStatus,omitempty"`
	VerificationStatus *CodeableConcept                `json:"verificationStatus,omitempty"`
	Type               *CodeableConcept                `json:"type,omitempty"`
	Category           []string                        `json:"category,omitempty"`
	Criticality        *string                         `json:"criticality,omitempty"`
	Code               *CodeableConcept                `json:"code,omitempty"`
	Patient            Reference                       `json:"patient"`
	Encounter          *Reference                      `json:"encounter,omitempty"`
	OnsetDateTime      *string                         `json:"onsetDateTime"`
	OnsetAge           *Age                            `json:"onsetAge"`
	OnsetPeriod        *Period                         `json:"onsetPeriod"`
	OnsetRange         *Range                          `json:"onsetRange"`
	OnsetString        *string                         `json:"onsetString"`
	RecordedDate       *string                         `json:"recordedDate,omitempty"`
	Participant        []AllergyIntoleranceParticipant `json:"participant,omitempty"`
	LastOccurrence     *string                         `json:"lastOccurrence,omitempty"`
	Note               []Annotation                    `json:"note,omitempty"`
	Reaction           []AllergyIntoleranceReaction    `json:"reaction,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AllergyIntolerance
type AllergyIntoleranceParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r5/StructureDefinition/AllergyIntolerance
type AllergyIntoleranceReaction struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Substance         *CodeableConcept    `json:"substance,omitempty"`
	Manifestation     []CodeableReference `json:"manifestation"`
	Description       *string             `json:"description,omitempty"`
	Onset             *string             `json:"onset,omitempty"`
	Severity          *string             `json:"severity,omitempty"`
	ExposureRoute     *CodeableConcept    `json:"exposureRoute,omitempty"`
	Note              []Annotation        `json:"note,omitempty"`
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
