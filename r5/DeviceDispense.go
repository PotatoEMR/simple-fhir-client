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

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDispense
type DeviceDispense struct {
	Id                    *string                   `json:"id,omitempty"`
	Meta                  *Meta                     `json:"meta,omitempty"`
	ImplicitRules         *string                   `json:"implicitRules,omitempty"`
	Language              *string                   `json:"language,omitempty"`
	Text                  *Narrative                `json:"text,omitempty"`
	Contained             []Resource                `json:"contained,omitempty"`
	Extension             []Extension               `json:"extension,omitempty"`
	ModifierExtension     []Extension               `json:"modifierExtension,omitempty"`
	Identifier            []Identifier              `json:"identifier,omitempty"`
	BasedOn               []Reference               `json:"basedOn,omitempty"`
	PartOf                []Reference               `json:"partOf,omitempty"`
	Status                string                    `json:"status"`
	StatusReason          *CodeableReference        `json:"statusReason,omitempty"`
	Category              []CodeableConcept         `json:"category,omitempty"`
	Device                CodeableReference         `json:"device"`
	Subject               Reference                 `json:"subject"`
	Receiver              *Reference                `json:"receiver,omitempty"`
	Encounter             *Reference                `json:"encounter,omitempty"`
	SupportingInformation []Reference               `json:"supportingInformation,omitempty"`
	Performer             []DeviceDispensePerformer `json:"performer,omitempty"`
	Location              *Reference                `json:"location,omitempty"`
	Type                  *CodeableConcept          `json:"type,omitempty"`
	Quantity              *Quantity                 `json:"quantity,omitempty"`
	PreparedDate          *time.Time                `json:"preparedDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	WhenHandedOver        *time.Time                `json:"whenHandedOver,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Destination           *Reference                `json:"destination,omitempty"`
	Note                  []Annotation              `json:"note,omitempty"`
	UsageInstruction      *string                   `json:"usageInstruction,omitempty"`
	EventHistory          []Reference               `json:"eventHistory,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDispense
type DeviceDispensePerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

type OtherDeviceDispense DeviceDispense

// on convert struct to json, automatically add resourceType=DeviceDispense
func (r DeviceDispense) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceDispense
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceDispense: OtherDeviceDispense(r),
		ResourceType:        "DeviceDispense",
	})
}
func (r DeviceDispense) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "DeviceDispense/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "DeviceDispense"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *DeviceDispense) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSDevicedispense_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDispense) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDispense) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDispense) T_PreparedDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("PreparedDate", nil, htmlAttrs)
	}
	return DateTimeInput("PreparedDate", resource.PreparedDate, htmlAttrs)
}
func (resource *DeviceDispense) T_WhenHandedOver(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("WhenHandedOver", nil, htmlAttrs)
	}
	return DateTimeInput("WhenHandedOver", resource.WhenHandedOver, htmlAttrs)
}
func (resource *DeviceDispense) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *DeviceDispense) T_UsageInstruction(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("UsageInstruction", nil, htmlAttrs)
	}
	return StringInput("UsageInstruction", resource.UsageInstruction, htmlAttrs)
}
func (resource *DeviceDispense) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("Performer["+strconv.Itoa(numPerformer)+"]Function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Performer["+strconv.Itoa(numPerformer)+"]Function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
