package r4

//generated with command go run ./bultaoreune
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
	EffectiveDateTime         *FhirDateTime     `json:"effectiveDateTime,omitempty"`
	EffectivePeriod           *Period           `json:"effectivePeriod,omitempty"`
	DateAsserted              *FhirDateTime     `json:"dateAsserted,omitempty"`
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
func (r MedicationStatement) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicationStatement/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "MedicationStatement"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MedicationStatement) T_BasedOn(numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput("basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *MedicationStatement) T_PartOf(numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput("partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *MedicationStatement) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMedication_statement_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicationStatement) T_StatusReason(numStatusReason int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatusReason >= len(resource.StatusReason) {
		return CodeableConceptSelect("statusReason["+strconv.Itoa(numStatusReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statusReason["+strconv.Itoa(numStatusReason)+"]", &resource.StatusReason[numStatusReason], optionsValueSet, htmlAttrs)
}
func (resource *MedicationStatement) T_Category(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category", resource.Category, optionsValueSet, htmlAttrs)
}
func (resource *MedicationStatement) T_MedicationCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("medicationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("medicationCodeableConcept", &resource.MedicationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationStatement) T_MedicationReference(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("medicationReference", nil, htmlAttrs)
	}
	return ReferenceInput("medicationReference", &resource.MedicationReference, htmlAttrs)
}
func (resource *MedicationStatement) T_Subject(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("subject", nil, htmlAttrs)
	}
	return ReferenceInput("subject", &resource.Subject, htmlAttrs)
}
func (resource *MedicationStatement) T_Context(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("context", nil, htmlAttrs)
	}
	return ReferenceInput("context", resource.Context, htmlAttrs)
}
func (resource *MedicationStatement) T_EffectiveDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("effectiveDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("effectiveDateTime", resource.EffectiveDateTime, htmlAttrs)
}
func (resource *MedicationStatement) T_EffectivePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("effectivePeriod", resource.EffectivePeriod, htmlAttrs)
}
func (resource *MedicationStatement) T_DateAsserted(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("dateAsserted", nil, htmlAttrs)
	}
	return FhirDateTimeInput("dateAsserted", resource.DateAsserted, htmlAttrs)
}
func (resource *MedicationStatement) T_InformationSource(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("informationSource", nil, htmlAttrs)
	}
	return ReferenceInput("informationSource", resource.InformationSource, htmlAttrs)
}
func (resource *MedicationStatement) T_DerivedFrom(numDerivedFrom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDerivedFrom >= len(resource.DerivedFrom) {
		return ReferenceInput("derivedFrom["+strconv.Itoa(numDerivedFrom)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("derivedFrom["+strconv.Itoa(numDerivedFrom)+"]", &resource.DerivedFrom[numDerivedFrom], htmlAttrs)
}
func (resource *MedicationStatement) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *MedicationStatement) T_ReasonReference(numReasonReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonReference >= len(resource.ReasonReference) {
		return ReferenceInput("reasonReference["+strconv.Itoa(numReasonReference)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("reasonReference["+strconv.Itoa(numReasonReference)+"]", &resource.ReasonReference[numReasonReference], htmlAttrs)
}
func (resource *MedicationStatement) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *MedicationStatement) T_Dosage(numDosage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDosage >= len(resource.Dosage) {
		return DosageInput("dosage["+strconv.Itoa(numDosage)+"]", nil, htmlAttrs)
	}
	return DosageInput("dosage["+strconv.Itoa(numDosage)+"]", &resource.Dosage[numDosage], htmlAttrs)
}
