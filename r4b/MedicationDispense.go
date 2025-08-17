//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationDispense
type MedicationDispense struct {
	Id                          *string                         `json:"id,omitempty"`
	Meta                        *Meta                           `json:"meta,omitempty"`
	ImplicitRules               *string                         `json:"implicitRules,omitempty"`
	Language                    *string                         `json:"language,omitempty"`
	Text                        *Narrative                      `json:"text,omitempty"`
	Contained                   []Resource                      `json:"contained,omitempty"`
	Extension                   []Extension                     `json:"extension,omitempty"`
	ModifierExtension           []Extension                     `json:"modifierExtension,omitempty"`
	Identifier                  []Identifier                    `json:"identifier,omitempty"`
	PartOf                      []Reference                     `json:"partOf,omitempty"`
	Status                      string                          `json:"status"`
	StatusReasonCodeableConcept *CodeableConcept                `json:"statusReasonCodeableConcept"`
	StatusReasonReference       *Reference                      `json:"statusReasonReference"`
	Category                    *CodeableConcept                `json:"category,omitempty"`
	MedicationCodeableConcept   CodeableConcept                 `json:"medicationCodeableConcept"`
	MedicationReference         Reference                       `json:"medicationReference"`
	Subject                     *Reference                      `json:"subject,omitempty"`
	Context                     *Reference                      `json:"context,omitempty"`
	SupportingInformation       []Reference                     `json:"supportingInformation,omitempty"`
	Performer                   []MedicationDispensePerformer   `json:"performer,omitempty"`
	Location                    *Reference                      `json:"location,omitempty"`
	AuthorizingPrescription     []Reference                     `json:"authorizingPrescription,omitempty"`
	Type                        *CodeableConcept                `json:"type,omitempty"`
	Quantity                    *Quantity                       `json:"quantity,omitempty"`
	DaysSupply                  *Quantity                       `json:"daysSupply,omitempty"`
	WhenPrepared                *string                         `json:"whenPrepared,omitempty"`
	WhenHandedOver              *string                         `json:"whenHandedOver,omitempty"`
	Destination                 *Reference                      `json:"destination,omitempty"`
	Receiver                    []Reference                     `json:"receiver,omitempty"`
	Note                        []Annotation                    `json:"note,omitempty"`
	DosageInstruction           []Dosage                        `json:"dosageInstruction,omitempty"`
	Substitution                *MedicationDispenseSubstitution `json:"substitution,omitempty"`
	DetectedIssue               []Reference                     `json:"detectedIssue,omitempty"`
	EventHistory                []Reference                     `json:"eventHistory,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationDispense
type MedicationDispensePerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationDispense
type MedicationDispenseSubstitution struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	WasSubstituted    bool              `json:"wasSubstituted"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Reason            []CodeableConcept `json:"reason,omitempty"`
	ResponsibleParty  []Reference       `json:"responsibleParty,omitempty"`
}

type OtherMedicationDispense MedicationDispense

// on convert struct to json, automatically add resourceType=MedicationDispense
func (r MedicationDispense) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationDispense
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationDispense: OtherMedicationDispense(r),
		ResourceType:            "MedicationDispense",
	})
}
