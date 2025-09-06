package r5

//generated with command go run ./bultaoreune -nodownload
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
	PreparedDate          *string                   `json:"preparedDate,omitempty"`
	WhenHandedOver        *string                   `json:"whenHandedOver,omitempty"`
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

func (resource *DeviceDispense) T_Id() templ.Component {

	if resource == nil {
		return StringInput("DeviceDispense.Id", nil)
	}
	return StringInput("DeviceDispense.Id", resource.Id)
}
func (resource *DeviceDispense) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("DeviceDispense.ImplicitRules", nil)
	}
	return StringInput("DeviceDispense.ImplicitRules", resource.ImplicitRules)
}
func (resource *DeviceDispense) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("DeviceDispense.Language", nil, optionsValueSet)
	}
	return CodeSelect("DeviceDispense.Language", resource.Language, optionsValueSet)
}
func (resource *DeviceDispense) T_Status() templ.Component {
	optionsValueSet := VSDevicedispense_status

	if resource == nil {
		return CodeSelect("DeviceDispense.Status", nil, optionsValueSet)
	}
	return CodeSelect("DeviceDispense.Status", &resource.Status, optionsValueSet)
}
func (resource *DeviceDispense) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("DeviceDispense.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDispense.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *DeviceDispense) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DeviceDispense.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDispense.Type", resource.Type, optionsValueSet)
}
func (resource *DeviceDispense) T_PreparedDate() templ.Component {

	if resource == nil {
		return StringInput("DeviceDispense.PreparedDate", nil)
	}
	return StringInput("DeviceDispense.PreparedDate", resource.PreparedDate)
}
func (resource *DeviceDispense) T_WhenHandedOver() templ.Component {

	if resource == nil {
		return StringInput("DeviceDispense.WhenHandedOver", nil)
	}
	return StringInput("DeviceDispense.WhenHandedOver", resource.WhenHandedOver)
}
func (resource *DeviceDispense) T_UsageInstruction() templ.Component {

	if resource == nil {
		return StringInput("DeviceDispense.UsageInstruction", nil)
	}
	return StringInput("DeviceDispense.UsageInstruction", resource.UsageInstruction)
}
func (resource *DeviceDispense) T_PerformerId(numPerformer int) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return StringInput("DeviceDispense.Performer["+strconv.Itoa(numPerformer)+"].Id", nil)
	}
	return StringInput("DeviceDispense.Performer["+strconv.Itoa(numPerformer)+"].Id", resource.Performer[numPerformer].Id)
}
func (resource *DeviceDispense) T_PerformerFunction(numPerformer int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Performer) >= numPerformer {
		return CodeableConceptSelect("DeviceDispense.Performer["+strconv.Itoa(numPerformer)+"].Function", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDispense.Performer["+strconv.Itoa(numPerformer)+"].Function", resource.Performer[numPerformer].Function, optionsValueSet)
}
