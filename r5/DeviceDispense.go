//generated August 15 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

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
