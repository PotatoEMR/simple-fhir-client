package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/DeviceAssociation
type DeviceAssociation struct {
	Id                *string                      `json:"id,omitempty"`
	Meta              *Meta                        `json:"meta,omitempty"`
	ImplicitRules     *string                      `json:"implicitRules,omitempty"`
	Language          *string                      `json:"language,omitempty"`
	Text              *Narrative                   `json:"text,omitempty"`
	Contained         []Resource                   `json:"contained,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                 `json:"identifier,omitempty"`
	Device            Reference                    `json:"device"`
	Category          []CodeableConcept            `json:"category,omitempty"`
	Status            CodeableConcept              `json:"status"`
	StatusReason      []CodeableConcept            `json:"statusReason,omitempty"`
	Subject           *Reference                   `json:"subject,omitempty"`
	BodyStructure     *Reference                   `json:"bodyStructure,omitempty"`
	Period            *Period                      `json:"period,omitempty"`
	Operation         []DeviceAssociationOperation `json:"operation,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceAssociation
type DeviceAssociationOperation struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Status            CodeableConcept `json:"status"`
	Operator          []Reference     `json:"operator,omitempty"`
	Period            *Period         `json:"period,omitempty"`
}

type OtherDeviceAssociation DeviceAssociation

// on convert struct to json, automatically add resourceType=DeviceAssociation
func (r DeviceAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceAssociation
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceAssociation: OtherDeviceAssociation(r),
		ResourceType:           "DeviceAssociation",
	})
}

func (resource *DeviceAssociation) T_Id() templ.Component {

	if resource == nil {
		return StringInput("DeviceAssociation.Id", nil)
	}
	return StringInput("DeviceAssociation.Id", resource.Id)
}
func (resource *DeviceAssociation) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("DeviceAssociation.ImplicitRules", nil)
	}
	return StringInput("DeviceAssociation.ImplicitRules", resource.ImplicitRules)
}
func (resource *DeviceAssociation) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("DeviceAssociation.Language", nil, optionsValueSet)
	}
	return CodeSelect("DeviceAssociation.Language", resource.Language, optionsValueSet)
}
func (resource *DeviceAssociation) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("DeviceAssociation.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceAssociation.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *DeviceAssociation) T_Status() templ.Component {
	optionsValueSet := VSDeviceassociation_status

	if resource == nil {
		return CodeableConceptSelect("DeviceAssociation.Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceAssociation.Status", &resource.Status, optionsValueSet)
}
func (resource *DeviceAssociation) T_StatusReason(numStatusReason int) templ.Component {
	optionsValueSet := VSDeviceassociation_status_reason

	if resource == nil || len(resource.StatusReason) >= numStatusReason {
		return CodeableConceptSelect("DeviceAssociation.StatusReason["+strconv.Itoa(numStatusReason)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceAssociation.StatusReason["+strconv.Itoa(numStatusReason)+"]", &resource.StatusReason[numStatusReason], optionsValueSet)
}
func (resource *DeviceAssociation) T_OperationId(numOperation int) templ.Component {

	if resource == nil || len(resource.Operation) >= numOperation {
		return StringInput("DeviceAssociation.Operation["+strconv.Itoa(numOperation)+"].Id", nil)
	}
	return StringInput("DeviceAssociation.Operation["+strconv.Itoa(numOperation)+"].Id", resource.Operation[numOperation].Id)
}
func (resource *DeviceAssociation) T_OperationStatus(numOperation int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Operation) >= numOperation {
		return CodeableConceptSelect("DeviceAssociation.Operation["+strconv.Itoa(numOperation)+"].Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceAssociation.Operation["+strconv.Itoa(numOperation)+"].Status", &resource.Operation[numOperation].Status, optionsValueSet)
}
