package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/MedicationAdministration
type MedicationAdministration struct {
	Id                    *string                             `json:"id,omitempty"`
	Meta                  *Meta                               `json:"meta,omitempty"`
	ImplicitRules         *string                             `json:"implicitRules,omitempty"`
	Language              *string                             `json:"language,omitempty"`
	Text                  *Narrative                          `json:"text,omitempty"`
	Contained             []Resource                          `json:"contained,omitempty"`
	Extension             []Extension                         `json:"extension,omitempty"`
	ModifierExtension     []Extension                         `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                        `json:"identifier,omitempty"`
	BasedOn               []Reference                         `json:"basedOn,omitempty"`
	PartOf                []Reference                         `json:"partOf,omitempty"`
	Status                string                              `json:"status"`
	StatusReason          []CodeableConcept                   `json:"statusReason,omitempty"`
	Category              []CodeableConcept                   `json:"category,omitempty"`
	Medication            CodeableReference                   `json:"medication"`
	Subject               Reference                           `json:"subject"`
	Encounter             *Reference                          `json:"encounter,omitempty"`
	SupportingInformation []Reference                         `json:"supportingInformation,omitempty"`
	OccurenceDateTime     FhirDateTime                        `json:"occurenceDateTime"`
	OccurencePeriod       Period                              `json:"occurencePeriod"`
	OccurenceTiming       Timing                              `json:"occurenceTiming"`
	Recorded              *FhirDateTime                       `json:"recorded,omitempty"`
	IsSubPotent           *bool                               `json:"isSubPotent,omitempty"`
	SubPotentReason       []CodeableConcept                   `json:"subPotentReason,omitempty"`
	Performer             []MedicationAdministrationPerformer `json:"performer,omitempty"`
	Reason                []CodeableReference                 `json:"reason,omitempty"`
	Request               *Reference                          `json:"request,omitempty"`
	Device                []CodeableReference                 `json:"device,omitempty"`
	Note                  []Annotation                        `json:"note,omitempty"`
	Dosage                *MedicationAdministrationDosage     `json:"dosage,omitempty"`
	EventHistory          []Reference                         `json:"eventHistory,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationAdministration
type MedicationAdministrationPerformer struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept  `json:"function,omitempty"`
	Actor             CodeableReference `json:"actor"`
}

// http://hl7.org/fhir/r5/StructureDefinition/MedicationAdministration
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

// struct -> json, automatically add resourceType=Patient
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
func (r MedicationAdministration) ResourceType() string {
	return "MedicationAdministration"
}

func (resource *MedicationAdministration) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
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
func (resource *MedicationAdministration) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_Medication(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("medication", nil, htmlAttrs)
	}
	return CodeableReferenceInput("medication", &resource.Medication, htmlAttrs)
}
func (resource *MedicationAdministration) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", &resource.Subject, htmlAttrs)
}
func (resource *MedicationAdministration) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *MedicationAdministration) T_SupportingInformation(frs []FhirResource, numSupportingInformation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInformation >= len(resource.SupportingInformation) {
		return ReferenceInput(frs, "supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", &resource.SupportingInformation[numSupportingInformation], htmlAttrs)
}
func (resource *MedicationAdministration) T_OccurenceDateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("occurenceDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("occurenceDateTime", &resource.OccurenceDateTime, htmlAttrs)
}
func (resource *MedicationAdministration) T_OccurencePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("occurencePeriod", nil, htmlAttrs)
	}
	return PeriodInput("occurencePeriod", &resource.OccurencePeriod, htmlAttrs)
}
func (resource *MedicationAdministration) T_OccurenceTiming(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return TimingInput("occurenceTiming", nil, htmlAttrs)
	}
	return TimingInput("occurenceTiming", &resource.OccurenceTiming, htmlAttrs)
}
func (resource *MedicationAdministration) T_Recorded(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("recorded", nil, htmlAttrs)
	}
	return FhirDateTimeInput("recorded", resource.Recorded, htmlAttrs)
}
func (resource *MedicationAdministration) T_IsSubPotent(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("isSubPotent", nil, htmlAttrs)
	}
	return BoolInput("isSubPotent", resource.IsSubPotent, htmlAttrs)
}
func (resource *MedicationAdministration) T_SubPotentReason(numSubPotentReason int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubPotentReason >= len(resource.SubPotentReason) {
		return CodeableConceptSelect("subPotentReason["+strconv.Itoa(numSubPotentReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subPotentReason["+strconv.Itoa(numSubPotentReason)+"]", &resource.SubPotentReason[numSubPotentReason], optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_Reason(numReason int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReason >= len(resource.Reason) {
		return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("reason["+strconv.Itoa(numReason)+"]", &resource.Reason[numReason], htmlAttrs)
}
func (resource *MedicationAdministration) T_Request(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "request", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "request", resource.Request, htmlAttrs)
}
func (resource *MedicationAdministration) T_Device(numDevice int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDevice >= len(resource.Device) {
		return CodeableReferenceInput("device["+strconv.Itoa(numDevice)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("device["+strconv.Itoa(numDevice)+"]", &resource.Device[numDevice], htmlAttrs)
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
func (resource *MedicationAdministration) T_PerformerActor(numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableReferenceInput("performer["+strconv.Itoa(numPerformer)+"].actor", nil, htmlAttrs)
	}
	return CodeableReferenceInput("performer["+strconv.Itoa(numPerformer)+"].actor", &resource.Performer[numPerformer].Actor, htmlAttrs)
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
