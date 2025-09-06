package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/MedicationStatement
type MedicationStatement struct {
	Id                        *string           `json:"id,omitempty"`
	Meta                      *Meta             `json:"meta,omitempty"`
	ImplicitRules             *string           `json:"implicitRules,omitempty"`
	Language                  *string           `json:"language,omitempty"`
	Text                      *Narrative        `json:"text,omitempty"`
	Contained                 []Resource        `json:"contained,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	ModifierExtension         []Extension       `json:"modifierExtension,omitempty"`
	Identifier                []Identifier      `json:"identifier,omitempty"`
	BasedOn                   []Reference       `json:"basedOn,omitempty"`
	PartOf                    []Reference       `json:"partOf,omitempty"`
	Status                    string            `json:"status"`
	StatusReason              []CodeableConcept `json:"statusReason,omitempty"`
	Category                  *CodeableConcept  `json:"category,omitempty"`
	MedicationCodeableConcept CodeableConcept   `json:"medicationCodeableConcept"`
	MedicationReference       Reference         `json:"medicationReference"`
	Subject                   Reference         `json:"subject"`
	Context                   *Reference        `json:"context,omitempty"`
	EffectiveDateTime         *string           `json:"effectiveDateTime,omitempty"`
	EffectivePeriod           *Period           `json:"effectivePeriod,omitempty"`
	DateAsserted              *string           `json:"dateAsserted,omitempty"`
	InformationSource         *Reference        `json:"informationSource,omitempty"`
	DerivedFrom               []Reference       `json:"derivedFrom,omitempty"`
	ReasonCode                []CodeableConcept `json:"reasonCode,omitempty"`
	ReasonReference           []Reference       `json:"reasonReference,omitempty"`
	Note                      []Annotation      `json:"note,omitempty"`
	Dosage                    []Dosage          `json:"dosage,omitempty"`
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
func (resource *MedicationStatement) T_StatusReason(numStatusReason int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.StatusReason) >= numStatusReason {
		return CodeableConceptSelect("MedicationStatement.StatusReason["+strconv.Itoa(numStatusReason)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationStatement.StatusReason["+strconv.Itoa(numStatusReason)+"]", &resource.StatusReason[numStatusReason], optionsValueSet)
}
func (resource *MedicationStatement) T_Category(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicationStatement.Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationStatement.Category", resource.Category, optionsValueSet)
}
func (resource *MedicationStatement) T_DateAsserted() templ.Component {

	if resource == nil {
		return StringInput("MedicationStatement.DateAsserted", nil)
	}
	return StringInput("MedicationStatement.DateAsserted", resource.DateAsserted)
}
func (resource *MedicationStatement) T_ReasonCode(numReasonCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ReasonCode) >= numReasonCode {
		return CodeableConceptSelect("MedicationStatement.ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicationStatement.ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet)
}
