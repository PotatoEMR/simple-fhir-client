package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *AllergyIntolerance) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *AllergyIntolerance) T_ClinicalStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("clinicalStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("clinicalStatus", resource.ClinicalStatus, optionsValueSet)
}
func (resource *AllergyIntolerance) T_VerificationStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("verificationStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("verificationStatus", resource.VerificationStatus, optionsValueSet)
}
func (resource *AllergyIntolerance) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *AllergyIntolerance) T_Category() templ.Component {
	optionsValueSet := VSAllergy_intolerance_category

	if resource == nil {
		return CodeSelect("category", nil, optionsValueSet)
	}
	return CodeSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *AllergyIntolerance) T_Criticality() templ.Component {
	optionsValueSet := VSAllergy_intolerance_criticality

	if resource == nil {
		return CodeSelect("criticality", nil, optionsValueSet)
	}
	return CodeSelect("criticality", resource.Criticality, optionsValueSet)
}
func (resource *AllergyIntolerance) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet)
}
func (resource *AllergyIntolerance) T_ParticipantFunction(numParticipant int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Participant) >= numParticipant {
		return CodeableConceptSelect("function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("function", resource.Participant[numParticipant].Function, optionsValueSet)
}
func (resource *AllergyIntolerance) T_ReactionSubstance(numReaction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Reaction) >= numReaction {
		return CodeableConceptSelect("substance", nil, optionsValueSet)
	}
	return CodeableConceptSelect("substance", resource.Reaction[numReaction].Substance, optionsValueSet)
}
func (resource *AllergyIntolerance) T_ReactionSeverity(numReaction int) templ.Component {
	optionsValueSet := VSReaction_event_severity

	if resource == nil && len(resource.Reaction) >= numReaction {
		return CodeSelect("severity", nil, optionsValueSet)
	}
	return CodeSelect("severity", resource.Reaction[numReaction].Severity, optionsValueSet)
}
func (resource *AllergyIntolerance) T_ReactionExposureRoute(numReaction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Reaction) >= numReaction {
		return CodeableConceptSelect("exposureRoute", nil, optionsValueSet)
	}
	return CodeableConceptSelect("exposureRoute", resource.Reaction[numReaction].ExposureRoute, optionsValueSet)
}
