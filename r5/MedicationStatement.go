package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
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
	EffectiveDateTime          *FhirDateTime                 `json:"effectiveDateTime,omitempty"`
	EffectivePeriod            *Period                       `json:"effectivePeriod,omitempty"`
	EffectiveTiming            *Timing                       `json:"effectiveTiming,omitempty"`
	DateAsserted               *FhirDateTime                 `json:"dateAsserted,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
func (r MedicationStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationStatement
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationStatement: OtherMedicationStatement(r),
		ResourceType:             "MedicationStatement",
	})
}

// json -> struct, first reject if resourceType != MedicationStatement
func (r *MedicationStatement) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "MedicationStatement" {
		return errors.New("resourceType not MedicationStatement")
	}
	return json.Unmarshal(data, (*OtherMedicationStatement)(r))
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
func (resource *MedicationStatement) T_PartOf(frs []FhirResource, numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *MedicationStatement) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMedication_statement_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicationStatement) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *MedicationStatement) T_Medication(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("medication", nil, htmlAttrs)
	}
	return CodeableReferenceInput("medication", &resource.Medication, htmlAttrs)
}
func (resource *MedicationStatement) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *MedicationStatement) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
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
func (resource *MedicationStatement) T_EffectiveTiming(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return TimingInput("effectiveTiming", nil, htmlAttrs)
	}
	return TimingInput("effectiveTiming", resource.EffectiveTiming, htmlAttrs)
}
func (resource *MedicationStatement) T_DateAsserted(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("dateAsserted", nil, htmlAttrs)
	}
	return FhirDateTimeInput("dateAsserted", resource.DateAsserted, htmlAttrs)
}
func (resource *MedicationStatement) T_InformationSource(frs []FhirResource, numInformationSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInformationSource >= len(resource.InformationSource) {
		return ReferenceInput(frs, "informationSource["+strconv.Itoa(numInformationSource)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "informationSource["+strconv.Itoa(numInformationSource)+"]", &resource.InformationSource[numInformationSource], htmlAttrs)
}
func (resource *MedicationStatement) T_DerivedFrom(frs []FhirResource, numDerivedFrom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDerivedFrom >= len(resource.DerivedFrom) {
		return ReferenceInput(frs, "derivedFrom["+strconv.Itoa(numDerivedFrom)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "derivedFrom["+strconv.Itoa(numDerivedFrom)+"]", &resource.DerivedFrom[numDerivedFrom], htmlAttrs)
}
func (resource *MedicationStatement) T_Reason(numReason int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], htmlAttrs)
}
func (resource *MedicationStatement) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *MedicationStatement) T_RelatedClinicalInformation(frs []FhirResource, numRelatedClinicalInformation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedClinicalInformation >= len(resource.RelatedClinicalInformation) {
		return ReferenceInput(frs, "relatedClinicalInformation["+strconv.Itoa(numRelatedClinicalInformation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "relatedClinicalInformation["+strconv.Itoa(numRelatedClinicalInformation)+"]", &resource.RelatedClinicalInformation[numRelatedClinicalInformation], htmlAttrs)
}
func (resource *MedicationStatement) T_RenderedDosageInstruction(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("renderedDosageInstruction", nil, htmlAttrs)
	}
	return StringInput("renderedDosageInstruction", resource.RenderedDosageInstruction, htmlAttrs)
}
func (resource *MedicationStatement) T_Dosage(numDosage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDosage >= len(resource.Dosage) {
		return DosageInput("dosage["+strconv.Itoa(numDosage)+"]", nil, htmlAttrs)
	}
	return DosageInput("dosage["+strconv.Itoa(numDosage)+"]", &resource.Dosage[numDosage], htmlAttrs)
}
func (resource *MedicationStatement) T_AdherenceCode(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("adherence.code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("adherence.code", &resource.Adherence.Code, optionsValueSet, htmlAttrs)
}
func (resource *MedicationStatement) T_AdherenceReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("adherence.reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("adherence.reason", resource.Adherence.Reason, optionsValueSet, htmlAttrs)
}
