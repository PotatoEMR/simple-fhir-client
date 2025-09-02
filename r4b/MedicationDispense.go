package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *MedicationDispense) MedicationDispenseLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *MedicationDispense) MedicationDispenseStatus() templ.Component {
	optionsValueSet := VSMedicationdispense_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *MedicationDispense) MedicationDispenseCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.Category, optionsValueSet)
}
func (resource *MedicationDispense) MedicationDispenseType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *MedicationDispense) MedicationDispensePerformerFunction(numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Performer) >= numPerformer {
		return CodeableConceptSelect("function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("function", resource.Performer[numPerformer].Function, optionsValueSet)
}
func (resource *MedicationDispense) MedicationDispenseSubstitutionType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Substitution.Type, optionsValueSet)
}
func (resource *MedicationDispense) MedicationDispenseSubstitutionReason(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reason", &resource.Substitution.Reason[0], optionsValueSet)
}
