package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/CompartmentDefinition
type CompartmentDefinition struct {
	Id                     *string                         `json:"id,omitempty"`
	Meta                   *Meta                           `json:"meta,omitempty"`
	ImplicitRules          *string                         `json:"implicitRules,omitempty"`
	Language               *string                         `json:"language,omitempty"`
	Text                   *Narrative                      `json:"text,omitempty"`
	Contained              []Resource                      `json:"contained,omitempty"`
	Extension              []Extension                     `json:"extension,omitempty"`
	ModifierExtension      []Extension                     `json:"modifierExtension,omitempty"`
	Url                    string                          `json:"url"`
	Version                *string                         `json:"version,omitempty"`
	VersionAlgorithmString *string                         `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding                         `json:"versionAlgorithmCoding"`
	Name                   string                          `json:"name"`
	Title                  *string                         `json:"title,omitempty"`
	Status                 string                          `json:"status"`
	Experimental           *bool                           `json:"experimental,omitempty"`
	Date                   *string                         `json:"date,omitempty"`
	Publisher              *string                         `json:"publisher,omitempty"`
	Contact                []ContactDetail                 `json:"contact,omitempty"`
	Description            *string                         `json:"description,omitempty"`
	UseContext             []UsageContext                  `json:"useContext,omitempty"`
	Purpose                *string                         `json:"purpose,omitempty"`
	Code                   string                          `json:"code"`
	Search                 bool                            `json:"search"`
	Resource               []CompartmentDefinitionResource `json:"resource,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/CompartmentDefinition
type CompartmentDefinitionResource struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Param             []string    `json:"param,omitempty"`
	Documentation     *string     `json:"documentation,omitempty"`
	StartParam        *string     `json:"startParam,omitempty"`
	EndParam          *string     `json:"endParam,omitempty"`
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
func (resource *CompartmentDefinition) CompartmentDefinitionLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *CompartmentDefinition) CompartmentDefinitionStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *CompartmentDefinition) CompartmentDefinitionCode() templ.Component {
	optionsValueSet := VSCompartment_type
	currentVal := ""
	if resource != nil {
		currentVal = resource.Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
func (resource *CompartmentDefinition) CompartmentDefinitionResourceCode(numResource int) templ.Component {
	optionsValueSet := VSResource_types
	currentVal := ""
	if resource != nil && len(resource.Resource) >= numResource {
		currentVal = resource.Resource[numResource].Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
