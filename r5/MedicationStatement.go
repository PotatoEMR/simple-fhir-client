package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	EffectiveDateTime          *string                       `json:"effectiveDateTime,omitempty"`
	EffectivePeriod            *Period                       `json:"effectivePeriod,omitempty"`
	EffectiveTiming            *Timing                       `json:"effectiveTiming,omitempty"`
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

func (resource *MedicationStatement) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MedicationStatement.Id", nil)
	}
	return StringInput("MedicationStatement.Id", resource.Id)
}
func (resource *MedicationStatement) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MedicationStatement.ImplicitRules", nil)
	}
	return StringInput("MedicationStatement.ImplicitRules", resource.ImplicitRules)
}
func (resource *MedicationStatement) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MedicationStatement.Language", nil, optionsValueSet)
	}
	return CodeSelect("MedicationStatement.Language", resource.Language, optionsValueSet)
}
func (resource *MedicationStatement) T_Status() templ.Component {
	optionsValueSet := VSMedication_statement_status

	if resource == nil {
		return CodeSelect("MedicationStatement.Status", nil, optionsValueSet)
	}
	return CodeSelect("MedicationStatement.Status", &resource.Status, optionsValueSet)
}
func (resource *MedicationStatement) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("MedicationStatement.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationStatement.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *MedicationStatement) T_DateAsserted() templ.Component {

	if resource == nil {
		return StringInput("MedicationStatement.DateAsserted", nil)
	}
	return StringInput("MedicationStatement.DateAsserted", resource.DateAsserted)
}
func (resource *MedicationStatement) T_RenderedDosageInstruction() templ.Component {

	if resource == nil {
		return StringInput("MedicationStatement.RenderedDosageInstruction", nil)
	}
	return StringInput("MedicationStatement.RenderedDosageInstruction", resource.RenderedDosageInstruction)
}
func (resource *MedicationStatement) T_AdherenceId() templ.Component {

	if resource == nil {
		return StringInput("MedicationStatement.Adherence.Id", nil)
	}
	return StringInput("MedicationStatement.Adherence.Id", resource.Adherence.Id)
}
func (resource *MedicationStatement) T_AdherenceCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationStatement.Adherence.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationStatement.Adherence.Code", &resource.Adherence.Code, optionsValueSet)
}
func (resource *MedicationStatement) T_AdherenceReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationStatement.Adherence.Reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationStatement.Adherence.Reason", resource.Adherence.Reason, optionsValueSet)
}
