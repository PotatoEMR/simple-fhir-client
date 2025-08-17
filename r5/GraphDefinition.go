//generated August 17 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/GraphDefinition
type GraphDefinition struct {
	Id                     *string               `json:"id,omitempty"`
	Meta                   *Meta                 `json:"meta,omitempty"`
	ImplicitRules          *string               `json:"implicitRules,omitempty"`
	Language               *string               `json:"language,omitempty"`
	Text                   *Narrative            `json:"text,omitempty"`
	Contained              []Resource            `json:"contained,omitempty"`
	Extension              []Extension           `json:"extension,omitempty"`
	ModifierExtension      []Extension           `json:"modifierExtension,omitempty"`
	Url                    *string               `json:"url,omitempty"`
	Identifier             []Identifier          `json:"identifier,omitempty"`
	Version                *string               `json:"version,omitempty"`
	VersionAlgorithmString *string               `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding               `json:"versionAlgorithmCoding"`
	Name                   string                `json:"name"`
	Title                  *string               `json:"title,omitempty"`
	Status                 string                `json:"status"`
	Experimental           *bool                 `json:"experimental,omitempty"`
	Date                   *string               `json:"date,omitempty"`
	Publisher              *string               `json:"publisher,omitempty"`
	Contact                []ContactDetail       `json:"contact,omitempty"`
	Description            *string               `json:"description,omitempty"`
	UseContext             []UsageContext        `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept     `json:"jurisdiction,omitempty"`
	Purpose                *string               `json:"purpose,omitempty"`
	Copyright              *string               `json:"copyright,omitempty"`
	CopyrightLabel         *string               `json:"copyrightLabel,omitempty"`
	Start                  *string               `json:"start,omitempty"`
	Node                   []GraphDefinitionNode `json:"node,omitempty"`
	Link                   []GraphDefinitionLink `json:"link,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/GraphDefinition
type GraphDefinitionNode struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	NodeId            string      `json:"nodeId"`
	Description       *string     `json:"description,omitempty"`
	Type              string      `json:"type"`
	Profile           *string     `json:"profile,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/GraphDefinition
type GraphDefinitionLink struct {
	Id                *string                          `json:"id,omitempty"`
	Extension         []Extension                      `json:"extension,omitempty"`
	ModifierExtension []Extension                      `json:"modifierExtension,omitempty"`
	Description       *string                          `json:"description,omitempty"`
	Min               *int                             `json:"min,omitempty"`
	Max               *string                          `json:"max,omitempty"`
	SourceId          string                           `json:"sourceId"`
	Path              *string                          `json:"path,omitempty"`
	SliceName         *string                          `json:"sliceName,omitempty"`
	TargetId          string                           `json:"targetId"`
	Params            *string                          `json:"params,omitempty"`
	Compartment       []GraphDefinitionLinkCompartment `json:"compartment,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/GraphDefinition
type GraphDefinitionLinkCompartment struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Use               string      `json:"use"`
	Rule              string      `json:"rule"`
	Code              string      `json:"code"`
	Expression        *string     `json:"expression,omitempty"`
	Description       *string     `json:"description,omitempty"`
}

type OtherGraphDefinition GraphDefinition

// on convert struct to json, automatically add resourceType=GraphDefinition
func (r GraphDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGraphDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherGraphDefinition: OtherGraphDefinition(r),
		ResourceType:         "GraphDefinition",
	})
}
