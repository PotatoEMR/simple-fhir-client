package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *DeviceAssociation) DeviceAssociationLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *DeviceAssociation) DeviceAssociationCategory(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *DeviceAssociation) DeviceAssociationStatus() templ.Component {
	optionsValueSet := VSDeviceassociation_status

	if resource == nil {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", &resource.Status, optionsValueSet)
}
func (resource *DeviceAssociation) DeviceAssociationStatusReason() templ.Component {
	optionsValueSet := VSDeviceassociation_status_reason

	if resource == nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("statusReason", &resource.StatusReason[0], optionsValueSet)
}
func (resource *DeviceAssociation) DeviceAssociationOperationStatus(numOperation int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Operation) >= numOperation {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", &resource.Operation[numOperation].Status, optionsValueSet)
}
