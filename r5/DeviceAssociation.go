package r5

//generated with command go run ./bultaoreune
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

// struct -> json, automatically add resourceType=Patient
func (r DeviceAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceAssociation
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceAssociation: OtherDeviceAssociation(r),
		ResourceType:           "DeviceAssociation",
	})
}

func (r DeviceAssociation) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "DeviceAssociation/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "DeviceAssociation"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r DeviceAssociation) ResourceType() string {
	return "DeviceAssociation"
}

func (resource *DeviceAssociation) T_Device(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "device", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "device", &resource.Device, htmlAttrs)
}
func (resource *DeviceAssociation) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *DeviceAssociation) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDeviceassociation_status

	if resource == nil {
		return CodeableConceptSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DeviceAssociation) T_StatusReason(numStatusReason int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDeviceassociation_status_reason

	if resource == nil || numStatusReason >= len(resource.StatusReason) {
		return CodeableConceptSelect("statusReason["+strconv.Itoa(numStatusReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statusReason["+strconv.Itoa(numStatusReason)+"]", &resource.StatusReason[numStatusReason], optionsValueSet, htmlAttrs)
}
func (resource *DeviceAssociation) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *DeviceAssociation) T_BodyStructure(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "bodyStructure", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "bodyStructure", resource.BodyStructure, htmlAttrs)
}
func (resource *DeviceAssociation) T_Period(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("period", nil, htmlAttrs)
	}
	return PeriodInput("period", resource.Period, htmlAttrs)
}
func (resource *DeviceAssociation) T_OperationStatus(numOperation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOperation >= len(resource.Operation) {
		return CodeableConceptSelect("operation["+strconv.Itoa(numOperation)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("operation["+strconv.Itoa(numOperation)+"].status", &resource.Operation[numOperation].Status, optionsValueSet, htmlAttrs)
}
func (resource *DeviceAssociation) T_OperationOperator(frs []FhirResource, numOperation int, numOperator int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOperation >= len(resource.Operation) || numOperator >= len(resource.Operation[numOperation].Operator) {
		return ReferenceInput(frs, "operation["+strconv.Itoa(numOperation)+"].operator["+strconv.Itoa(numOperator)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "operation["+strconv.Itoa(numOperation)+"].operator["+strconv.Itoa(numOperator)+"]", &resource.Operation[numOperation].Operator[numOperator], htmlAttrs)
}
func (resource *DeviceAssociation) T_OperationPeriod(numOperation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOperation >= len(resource.Operation) {
		return PeriodInput("operation["+strconv.Itoa(numOperation)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("operation["+strconv.Itoa(numOperation)+"].period", resource.Operation[numOperation].Period, htmlAttrs)
}
