package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/CompartmentDefinition
type CompartmentDefinition struct {
	Id                *string                         `json:"id,omitempty"`
	Meta              *Meta                           `json:"meta,omitempty"`
	ImplicitRules     *string                         `json:"implicitRules,omitempty"`
	Language          *string                         `json:"language,omitempty"`
	Text              *Narrative                      `json:"text,omitempty"`
	Contained         []Resource                      `json:"contained,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Url               string                          `json:"url"`
	Version           *string                         `json:"version,omitempty"`
	Name              string                          `json:"name"`
	Status            string                          `json:"status"`
	Experimental      *bool                           `json:"experimental,omitempty"`
	Date              *string                         `json:"date,omitempty"`
	Publisher         *string                         `json:"publisher,omitempty"`
	Contact           []ContactDetail                 `json:"contact,omitempty"`
	Description       *string                         `json:"description,omitempty"`
	UseContext        []UsageContext                  `json:"useContext,omitempty"`
	Purpose           *string                         `json:"purpose,omitempty"`
	Code              string                          `json:"code"`
	Search            bool                            `json:"search"`
	Resource          []CompartmentDefinitionResource `json:"resource,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/CompartmentDefinition
type CompartmentDefinitionResource struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Param             []string    `json:"param,omitempty"`
	Documentation     *string     `json:"documentation,omitempty"`
}

type OtherCompartmentDefinition CompartmentDefinition

// on convert struct to json, automatically add resourceType=CompartmentDefinition
func (r CompartmentDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCompartmentDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherCompartmentDefinition: OtherCompartmentDefinition(r),
		ResourceType:               "CompartmentDefinition",
	})
}

func (resource *CompartmentDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *CompartmentDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *CompartmentDefinition) T_Code() templ.Component {
	optionsValueSet := VSCompartment_type

	if resource == nil {
		return CodeSelect("code", nil, optionsValueSet)
	}
	return CodeSelect("code", &resource.Code, optionsValueSet)
}
func (resource *CompartmentDefinition) T_ResourceCode(numResource int) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil && len(resource.Resource) >= numResource {
		return CodeSelect("code", nil, optionsValueSet)
	}
	return CodeSelect("code", &resource.Resource[numResource].Code, optionsValueSet)
}
