//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

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
