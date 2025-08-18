//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/MedicationStatement
type MedicationStatement struct {
	Id                         *string                       `json:"id,omitempty"`
	Meta                       *Meta                         `json:"meta,omitempty"`
	ImplicitRules              *string                       `json:"implicitRules,omitempty"`
	Language                   *string                       `json:"language,omitempty"`
	Text                       *Narrative                    `json:"text,omitempty"`
	Contained                  []Resource                    `json:"contained,omitempty"`
	Extension                  []Extension                   `json:"extension,omitempty"`
	ModifierExtension          []Extension                   `json:"modifierExtension,omitempty"`
	Identifier                 []Identifier                  `json:"identifier,omitempty"`
	PartOf                     []Reference                   `json:"partOf,omitempty"`
	Status                     string                        `json:"status"`
	Category                   []CodeableConcept             `json:"category,omitempty"`
	Medication                 CodeableReference             `json:"medication"`
	Subject                    Reference                     `json:"subject"`
	Encounter                  *Reference                    `json:"encounter,omitempty"`
	EffectiveDateTime          *string                       `json:"effectiveDateTime"`
	EffectivePeriod            *Period                       `json:"effectivePeriod"`
	EffectiveTiming            *Timing                       `json:"effectiveTiming"`
	DateAsserted               *string                       `json:"dateAsserted,omitempty"`
	InformationSource          []Reference                   `json:"informationSource,omitempty"`
	DerivedFrom                []Reference                   `json:"derivedFrom,omitempty"`
	Reason                     []CodeableReference           `json:"reason,omitempty"`
	Note                       []Annotation                  `json:"note,omitempty"`
	RelatedClinicalInformation []Reference                   `json:"relatedClinicalInformation,omitempty"`
	RenderedDosageInstruction  *string                       `json:"renderedDosageInstruction,omitempty"`
	Dosage                     []Dosage                      `json:"dosage,omitempty"`
	Adherence                  *MedicationStatementAdherence `json:"adherence,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationStatement
type MedicationStatementAdherence struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              CodeableConcept  `json:"code"`
	Reason            *CodeableConcept `json:"reason,omitempty"`
}

type OtherMedicationStatement MedicationStatement

// on convert struct to json, automatically add resourceType=MedicationStatement
func (r MedicationStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationStatement
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationStatement: OtherMedicationStatement(r),
		ResourceType:             "MedicationStatement",
	})
}
