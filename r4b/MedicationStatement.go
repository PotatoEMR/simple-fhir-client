package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/MedicationStatement
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
	EffectiveDateTime         *string           `json:"effectiveDateTime"`
	EffectivePeriod           *Period           `json:"effectivePeriod"`
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

func (resource *MedicationStatement) MedicationStatementLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *MedicationStatement) MedicationStatementStatus() templ.Component {
	optionsValueSet := VSMedication_statement_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
