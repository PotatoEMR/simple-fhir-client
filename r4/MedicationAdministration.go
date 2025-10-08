package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/MedicationAdministration
type MedicationAdministration struct {
	Id                        *string                             `json:"id,omitempty"`
	Meta                      *Meta                               `json:"meta,omitempty"`
	ImplicitRules             *string                             `json:"implicitRules,omitempty"`
	Language                  *string                             `json:"language,omitempty"`
	Text                      *Narrative                          `json:"text,omitempty"`
	Contained                 []Resource                          `json:"contained,omitempty"`
	Extension                 []Extension                         `json:"extension,omitempty"`
	ModifierExtension         []Extension                         `json:"modifierExtension,omitempty"`
	Identifier                []Identifier                        `json:"identifier,omitempty"`
	Instantiates              []string                            `json:"instantiates,omitempty"`
	PartOf                    []Reference                         `json:"partOf,omitempty"`
	Status                    string                              `json:"status"`
	StatusReason              []CodeableConcept                   `json:"statusReason,omitempty"`
	Category                  *CodeableConcept                    `json:"category,omitempty"`
	MedicationCodeableConcept CodeableConcept                     `json:"medicationCodeableConcept"`
	MedicationReference       Reference                           `json:"medicationReference"`
	Subject                   Reference                           `json:"subject"`
	Context                   *Reference                          `json:"context,omitempty"`
	SupportingInformation     []Reference                         `json:"supportingInformation,omitempty"`
	EffectiveDateTime         FhirDateTime                        `json:"effectiveDateTime"`
	EffectivePeriod           Period                              `json:"effectivePeriod"`
	Performer                 []MedicationAdministrationPerformer `json:"performer,omitempty"`
	ReasonCode                []CodeableConcept                   `json:"reasonCode,omitempty"`
	ReasonReference           []Reference                         `json:"reasonReference,omitempty"`
	Request                   *Reference                          `json:"request,omitempty"`
	Device                    []Reference                         `json:"device,omitempty"`
	Note                      []Annotation                        `json:"note,omitempty"`
	Dosage                    *MedicationAdministrationDosage     `json:"dosage,omitempty"`
	EventHistory              []Reference                         `json:"eventHistory,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationAdministration
type MedicationAdministrationPerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicationAdministration
type MedicationAdministrationDosage struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Text              *string          `json:"text,omitempty"`
	Site              *CodeableConcept `json:"site,omitempty"`
	Route             *CodeableConcept `json:"route,omitempty"`
	Method            *CodeableConcept `json:"method,omitempty"`
	Dose              *Quantity        `json:"dose,omitempty"`
	RateRatio         *Ratio           `json:"rateRatio,omitempty"`
	RateQuantity      *Quantity        `json:"rateQuantity,omitempty"`
}

type OtherMedicationAdministration MedicationAdministration

// on convert struct to json, automatically add resourceType=MedicationAdministration
func (r MedicationAdministration) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationAdministration
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationAdministration: OtherMedicationAdministration(r),
		ResourceType:                  "MedicationAdministration",
	})
}
func (r MedicationAdministration) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicationAdministration/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "MedicationAdministration"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MedicationAdministration) T_Instantiates(numInstantiates int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInstantiates >= len(resource.Instantiates) {
		return StringInput("instantiates["+strconv.Itoa(numInstantiates)+"]", nil, htmlAttrs)
	}
	return StringInput("instantiates["+strconv.Itoa(numInstantiates)+"]", &resource.Instantiates[numInstantiates], htmlAttrs)
}
func (resource *MedicationAdministration) T_PartOf(frs []FhirResource, numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *MedicationAdministration) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSMedication_admin_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_StatusReason(numStatusReason int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatusReason >= len(resource.StatusReason) {
		return CodeableConceptSelect("statusReason["+strconv.Itoa(numStatusReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statusReason["+strconv.Itoa(numStatusReason)+"]", &resource.StatusReason[numStatusReason], optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_Category(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category", resource.Category, optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_MedicationCodeableConcept(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("medicationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("medicationCodeableConcept", &resource.MedicationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_MedicationReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "medicationReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "medicationReference", &resource.MedicationReference, htmlAttrs)
}
func (resource *MedicationAdministration) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *MedicationAdministration) T_Context(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "context", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "context", resource.Context, htmlAttrs)
}
func (resource *MedicationAdministration) T_SupportingInformation(frs []FhirResource, numSupportingInformation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInformation >= len(resource.SupportingInformation) {
		return ReferenceInput(frs, "supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", &resource.SupportingInformation[numSupportingInformation], htmlAttrs)
}
func (resource *MedicationAdministration) T_EffectiveDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("effectiveDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("effectiveDateTime", &resource.EffectiveDateTime, htmlAttrs)
}
func (resource *MedicationAdministration) T_EffectivePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("effectivePeriod", &resource.EffectivePeriod, htmlAttrs)
}
func (resource *MedicationAdministration) T_ReasonCode(numReasonCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonCode >= len(resource.ReasonCode) {
		return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_ReasonReference(frs []FhirResource, numReasonReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReasonReference >= len(resource.ReasonReference) {
		return ReferenceInput(frs, "reasonReference["+strconv.Itoa(numReasonReference)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "reasonReference["+strconv.Itoa(numReasonReference)+"]", &resource.ReasonReference[numReasonReference], htmlAttrs)
}
func (resource *MedicationAdministration) T_Request(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "request", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "request", resource.Request, htmlAttrs)
}
func (resource *MedicationAdministration) T_Device(frs []FhirResource, numDevice int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDevice >= len(resource.Device) {
		return ReferenceInput(frs, "device["+strconv.Itoa(numDevice)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "device["+strconv.Itoa(numDevice)+"]", &resource.Device[numDevice], htmlAttrs)
}
func (resource *MedicationAdministration) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *MedicationAdministration) T_EventHistory(frs []FhirResource, numEventHistory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEventHistory >= len(resource.EventHistory) {
		return ReferenceInput(frs, "eventHistory["+strconv.Itoa(numEventHistory)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "eventHistory["+strconv.Itoa(numEventHistory)+"]", &resource.EventHistory[numEventHistory], htmlAttrs)
}
func (resource *MedicationAdministration) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_PerformerActor(frs []FhirResource, numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"].actor", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"].actor", &resource.Performer[numPerformer].Actor, htmlAttrs)
}
func (resource *MedicationAdministration) T_DosageSite(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("dosage.site", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("dosage.site", resource.Dosage.Site, optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_DosageRoute(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("dosage.route", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("dosage.route", resource.Dosage.Route, optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_DosageMethod(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("dosage.method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("dosage.method", resource.Dosage.Method, optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_DosageDose(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("dosage.dose", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("dosage.dose", resource.Dosage.Dose, optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_DosageRateRatio(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return RatioInput("dosage.rateRatio", nil, htmlAttrs)
	}
	return RatioInput("dosage.rateRatio", resource.Dosage.RateRatio, htmlAttrs)
}
func (resource *MedicationAdministration) T_DosageRateQuantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("dosage.rateQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("dosage.rateQuantity", resource.Dosage.RateQuantity, optionsValueSet, htmlAttrs)
}
