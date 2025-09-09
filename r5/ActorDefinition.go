package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ActorDefinition
type ActorDefinition struct {
	Id                     *string           `json:"id,omitempty"`
	Meta                   *Meta             `json:"meta,omitempty"`
	ImplicitRules          *string           `json:"implicitRules,omitempty"`
	Language               *string           `json:"language,omitempty"`
	Text                   *Narrative        `json:"text,omitempty"`
	Contained              []Resource        `json:"contained,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Url                    *string           `json:"url,omitempty"`
	Identifier             []Identifier      `json:"identifier,omitempty"`
	Version                *string           `json:"version,omitempty"`
	VersionAlgorithmString *string           `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding           `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string           `json:"name,omitempty"`
	Title                  *string           `json:"title,omitempty"`
	Status                 string            `json:"status"`
	Experimental           *bool             `json:"experimental,omitempty"`
	Date                   *time.Time        `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher              *string           `json:"publisher,omitempty"`
	Contact                []ContactDetail   `json:"contact,omitempty"`
	Description            *string           `json:"description,omitempty"`
	UseContext             []UsageContext    `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept `json:"jurisdiction,omitempty"`
	Purpose                *string           `json:"purpose,omitempty"`
	Copyright              *string           `json:"copyright,omitempty"`
	CopyrightLabel         *string           `json:"copyrightLabel,omitempty"`
	Type                   string            `json:"type"`
	Documentation          *string           `json:"documentation,omitempty"`
	Reference              []string          `json:"reference,omitempty"`
	Capabilities           *string           `json:"capabilities,omitempty"`
	DerivedFrom            []string          `json:"derivedFrom,omitempty"`
}

type OtherActorDefinition ActorDefinition

// on convert struct to json, automatically add resourceType=ActorDefinition
func (r ActorDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherActorDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherActorDefinition: OtherActorDefinition(r),
		ResourceType:         "ActorDefinition",
	})
}
func (r ActorDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ActorDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ActorDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ActorDefinition) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActorDefinition.Url", nil, htmlAttrs)
	}
	return StringInput("ActorDefinition.Url", resource.Url, htmlAttrs)
}
func (resource *ActorDefinition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActorDefinition.Version", nil, htmlAttrs)
	}
	return StringInput("ActorDefinition.Version", resource.Version, htmlAttrs)
}
func (resource *ActorDefinition) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActorDefinition.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("ActorDefinition.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *ActorDefinition) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("ActorDefinition.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("ActorDefinition.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *ActorDefinition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActorDefinition.Name", nil, htmlAttrs)
	}
	return StringInput("ActorDefinition.Name", resource.Name, htmlAttrs)
}
func (resource *ActorDefinition) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActorDefinition.Title", nil, htmlAttrs)
	}
	return StringInput("ActorDefinition.Title", resource.Title, htmlAttrs)
}
func (resource *ActorDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ActorDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ActorDefinition.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ActorDefinition) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("ActorDefinition.Experimental", nil, htmlAttrs)
	}
	return BoolInput("ActorDefinition.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *ActorDefinition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("ActorDefinition.Date", nil, htmlAttrs)
	}
	return DateTimeInput("ActorDefinition.Date", resource.Date, htmlAttrs)
}
func (resource *ActorDefinition) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActorDefinition.Publisher", nil, htmlAttrs)
	}
	return StringInput("ActorDefinition.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *ActorDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActorDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("ActorDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *ActorDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("ActorDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ActorDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *ActorDefinition) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActorDefinition.Purpose", nil, htmlAttrs)
	}
	return StringInput("ActorDefinition.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *ActorDefinition) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActorDefinition.Copyright", nil, htmlAttrs)
	}
	return StringInput("ActorDefinition.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *ActorDefinition) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActorDefinition.CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("ActorDefinition.CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *ActorDefinition) T_Type(htmlAttrs string) templ.Component {
	optionsValueSet := VSExamplescenario_actor_type

	if resource == nil {
		return CodeSelect("ActorDefinition.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ActorDefinition.Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *ActorDefinition) T_Documentation(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActorDefinition.Documentation", nil, htmlAttrs)
	}
	return StringInput("ActorDefinition.Documentation", resource.Documentation, htmlAttrs)
}
func (resource *ActorDefinition) T_Reference(numReference int, htmlAttrs string) templ.Component {
	if resource == nil || numReference >= len(resource.Reference) {
		return StringInput("ActorDefinition.Reference["+strconv.Itoa(numReference)+"]", nil, htmlAttrs)
	}
	return StringInput("ActorDefinition.Reference["+strconv.Itoa(numReference)+"]", &resource.Reference[numReference], htmlAttrs)
}
func (resource *ActorDefinition) T_Capabilities(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ActorDefinition.Capabilities", nil, htmlAttrs)
	}
	return StringInput("ActorDefinition.Capabilities", resource.Capabilities, htmlAttrs)
}
func (resource *ActorDefinition) T_DerivedFrom(numDerivedFrom int, htmlAttrs string) templ.Component {
	if resource == nil || numDerivedFrom >= len(resource.DerivedFrom) {
		return StringInput("ActorDefinition.DerivedFrom["+strconv.Itoa(numDerivedFrom)+"]", nil, htmlAttrs)
	}
	return StringInput("ActorDefinition.DerivedFrom["+strconv.Itoa(numDerivedFrom)+"]", &resource.DerivedFrom[numDerivedFrom], htmlAttrs)
}
