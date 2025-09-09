package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	OccurenceDateTime     time.Time                           `json:"occurenceDateTime,format:'2006-01-02T15:04:05Z07:00'"`
	OccurencePeriod       Period                              `json:"occurencePeriod"`
	OccurenceTiming       Timing                              `json:"occurenceTiming"`
	Recorded              *time.Time                          `json:"recorded,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *MedicationAdministration) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSMedication_admin_status

	if resource == nil {
		return CodeSelect("MedicationAdministration.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("MedicationAdministration.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_StatusReason(numStatusReason int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numStatusReason >= len(resource.StatusReason) {
		return CodeableConceptSelect("MedicationAdministration.StatusReason["+strconv.Itoa(numStatusReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationAdministration.StatusReason["+strconv.Itoa(numStatusReason)+"]", &resource.StatusReason[numStatusReason], optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("MedicationAdministration.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationAdministration.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_OccurenceDateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("MedicationAdministration.OccurenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("MedicationAdministration.OccurenceDateTime", &resource.OccurenceDateTime, htmlAttrs)
}
func (resource *MedicationAdministration) T_Recorded(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("MedicationAdministration.Recorded", nil, htmlAttrs)
	}
	return DateTimeInput("MedicationAdministration.Recorded", resource.Recorded, htmlAttrs)
}
func (resource *MedicationAdministration) T_IsSubPotent(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("MedicationAdministration.IsSubPotent", nil, htmlAttrs)
	}
	return BoolInput("MedicationAdministration.IsSubPotent", resource.IsSubPotent, htmlAttrs)
}
func (resource *MedicationAdministration) T_SubPotentReason(numSubPotentReason int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSubPotentReason >= len(resource.SubPotentReason) {
		return CodeableConceptSelect("MedicationAdministration.SubPotentReason["+strconv.Itoa(numSubPotentReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationAdministration.SubPotentReason["+strconv.Itoa(numSubPotentReason)+"]", &resource.SubPotentReason[numSubPotentReason], optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("MedicationAdministration.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("MedicationAdministration.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *MedicationAdministration) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("MedicationAdministration.Performer["+strconv.Itoa(numPerformer)+"].Function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationAdministration.Performer["+strconv.Itoa(numPerformer)+"].Function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_DosageText(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("MedicationAdministration.Dosage.Text", nil, htmlAttrs)
	}
	return StringInput("MedicationAdministration.Dosage.Text", resource.Dosage.Text, htmlAttrs)
}
func (resource *MedicationAdministration) T_DosageSite(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("MedicationAdministration.Dosage.Site", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationAdministration.Dosage.Site", resource.Dosage.Site, optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_DosageRoute(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("MedicationAdministration.Dosage.Route", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationAdministration.Dosage.Route", resource.Dosage.Route, optionsValueSet, htmlAttrs)
}
func (resource *MedicationAdministration) T_DosageMethod(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("MedicationAdministration.Dosage.Method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("MedicationAdministration.Dosage.Method", resource.Dosage.Method, optionsValueSet, htmlAttrs)
}
