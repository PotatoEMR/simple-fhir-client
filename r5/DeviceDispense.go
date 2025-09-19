package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	PreparedDate          *FhirDateTime             `json:"preparedDate,omitempty"`
	WhenHandedOver        *FhirDateTime             `json:"whenHandedOver,omitempty"`
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
func (resource *DeviceDispense) T_BasedOn(numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput("basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *DeviceDispense) T_PartOf(numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput("partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *DeviceDispense) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDevicedispense_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDispense) T_StatusReason(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("statusReason", nil, htmlAttrs)
	}
	return CodeableReferenceInput("statusReason", resource.StatusReason, htmlAttrs)
}
func (resource *DeviceDispense) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDispense) T_Device(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableReferenceInput("device", nil, htmlAttrs)
	}
	return CodeableReferenceInput("device", &resource.Device, htmlAttrs)
}
func (resource *DeviceDispense) T_Subject(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("subject", nil, htmlAttrs)
	}
	return ReferenceInput("subject", &resource.Subject, htmlAttrs)
}
func (resource *DeviceDispense) T_Receiver(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("receiver", nil, htmlAttrs)
	}
	return ReferenceInput("receiver", resource.Receiver, htmlAttrs)
}
func (resource *DeviceDispense) T_Encounter(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("encounter", nil, htmlAttrs)
	}
	return ReferenceInput("encounter", resource.Encounter, htmlAttrs)
}
func (resource *DeviceDispense) T_SupportingInformation(numSupportingInformation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInformation >= len(resource.SupportingInformation) {
		return ReferenceInput("supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("supportingInformation["+strconv.Itoa(numSupportingInformation)+"]", &resource.SupportingInformation[numSupportingInformation], htmlAttrs)
}
func (resource *DeviceDispense) T_Location(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("location", nil, htmlAttrs)
	}
	return ReferenceInput("location", resource.Location, htmlAttrs)
}
func (resource *DeviceDispense) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDispense) T_Quantity(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return QuantityInput("quantity", nil, htmlAttrs)
	}
	return QuantityInput("quantity", resource.Quantity, htmlAttrs)
}
func (resource *DeviceDispense) T_PreparedDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("preparedDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("preparedDate", resource.PreparedDate, htmlAttrs)
}
func (resource *DeviceDispense) T_WhenHandedOver(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("whenHandedOver", nil, htmlAttrs)
	}
	return FhirDateTimeInput("whenHandedOver", resource.WhenHandedOver, htmlAttrs)
}
func (resource *DeviceDispense) T_Destination(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("destination", nil, htmlAttrs)
	}
	return ReferenceInput("destination", resource.Destination, htmlAttrs)
}
func (resource *DeviceDispense) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *DeviceDispense) T_UsageInstruction(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("usageInstruction", nil, htmlAttrs)
	}
	return StringInput("usageInstruction", resource.UsageInstruction, htmlAttrs)
}
func (resource *DeviceDispense) T_EventHistory(numEventHistory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEventHistory >= len(resource.EventHistory) {
		return ReferenceInput("eventHistory["+strconv.Itoa(numEventHistory)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("eventHistory["+strconv.Itoa(numEventHistory)+"]", &resource.EventHistory[numEventHistory], htmlAttrs)
}
func (resource *DeviceDispense) T_PerformerFunction(numPerformer int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("performer["+strconv.Itoa(numPerformer)+"].function", resource.Performer[numPerformer].Function, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDispense) T_PerformerActor(numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return ReferenceInput("performer["+strconv.Itoa(numPerformer)+"].actor", nil, htmlAttrs)
	}
	return ReferenceInput("performer["+strconv.Itoa(numPerformer)+"].actor", &resource.Performer[numPerformer].Actor, htmlAttrs)
}
