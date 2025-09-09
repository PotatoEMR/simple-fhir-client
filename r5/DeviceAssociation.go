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
func (resource *DeviceAssociation) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("DeviceAssociation.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceAssociation.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *DeviceAssociation) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSDeviceassociation_status

	if resource == nil {
		return CodeableConceptSelect("DeviceAssociation.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceAssociation.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *DeviceAssociation) T_StatusReason(numStatusReason int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDeviceassociation_status_reason

	if resource == nil || numStatusReason >= len(resource.StatusReason) {
		return CodeableConceptSelect("DeviceAssociation.StatusReason["+strconv.Itoa(numStatusReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceAssociation.StatusReason["+strconv.Itoa(numStatusReason)+"]", &resource.StatusReason[numStatusReason], optionsValueSet, htmlAttrs)
}
func (resource *DeviceAssociation) T_OperationStatus(numOperation int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numOperation >= len(resource.Operation) {
		return CodeableConceptSelect("DeviceAssociation.Operation["+strconv.Itoa(numOperation)+"].Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceAssociation.Operation["+strconv.Itoa(numOperation)+"].Status", &resource.Operation[numOperation].Status, optionsValueSet, htmlAttrs)
}
