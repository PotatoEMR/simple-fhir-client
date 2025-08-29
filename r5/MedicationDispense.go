package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/MedicationDispense
type MedicationDispense struct {
	Id                        *string                         `json:"id,omitempty"`
	Meta                      *Meta                           `json:"meta,omitempty"`
	ImplicitRules             *string                         `json:"implicitRules,omitempty"`
	Language                  *string                         `json:"language,omitempty"`
	Text                      *Narrative                      `json:"text,omitempty"`
	Contained                 []Resource                      `json:"contained,omitempty"`
	Extension                 []Extension                     `json:"extension,omitempty"`
	ModifierExtension         []Extension                     `json:"modifierExtension,omitempty"`
	Identifier                []Identifier                    `json:"identifier,omitempty"`
	BasedOn                   []Reference                     `json:"basedOn,omitempty"`
	PartOf                    []Reference                     `json:"partOf,omitempty"`
	Status                    string                          `json:"status"`
	NotPerformedReason        *CodeableReference              `json:"notPerformedReason,omitempty"`
	StatusChanged             *string                         `json:"statusChanged,omitempty"`
	Category                  []CodeableConcept               `json:"category,omitempty"`
	Medication                CodeableReference               `json:"medication"`
	Subject                   Reference                       `json:"subject"`
	Encounter                 *Reference                      `json:"encounter,omitempty"`
	SupportingInformation     []Reference                     `json:"supportingInformation,omitempty"`
	Performer                 []MedicationDispensePerformer   `json:"performer,omitempty"`
	Location                  *Reference                      `json:"location,omitempty"`
	AuthorizingPrescription   []Reference                     `json:"authorizingPrescription,omitempty"`
	Type                      *CodeableConcept                `json:"type,omitempty"`
	Quantity                  *Quantity                       `json:"quantity,omitempty"`
	DaysSupply                *Quantity                       `json:"daysSupply,omitempty"`
	Recorded                  *string                         `json:"recorded,omitempty"`
	WhenPrepared              *string                         `json:"whenPrepared,omitempty"`
	WhenHandedOver            *string                         `json:"whenHandedOver,omitempty"`
	Destination               *Reference                      `json:"destination,omitempty"`
	Receiver                  []Reference                     `json:"receiver,omitempty"`
	Note                      []Annotation                    `json:"note,omitempty"`
	RenderedDosageInstruction *string                         `json:"renderedDosageInstruction,omitempty"`
	DosageInstruction         []Dosage                        `json:"dosageInstruction,omitempty"`
	Substitution              *MedicationDispenseSubstitution `json:"substitution,omitempty"`
	EventHistory              []Reference                     `json:"eventHistory,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationDispense
type MedicationDispensePerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationDispense
type MedicationDispenseSubstitution struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	WasSubstituted    bool              `json:"wasSubstituted"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Reason            []CodeableConcept `json:"reason,omitempty"`
	ResponsibleParty  *Reference        `json:"responsibleParty,omitempty"`
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

func (resource *MedicationDispense) MedicationDispenseLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *MedicationDispense) MedicationDispenseStatus() templ.Component {
	optionsValueSet := VSMedicationdispense_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
