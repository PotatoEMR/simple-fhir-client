package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Date                   *string           `json:"date,omitempty"`
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

func (resource *ActorDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ActorDefinition.Id", nil)
	}
	return StringInput("ActorDefinition.Id", resource.Id)
}
func (resource *ActorDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ActorDefinition.ImplicitRules", nil)
	}
	return StringInput("ActorDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *ActorDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ActorDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("ActorDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *ActorDefinition) T_Url() templ.Component {

	if resource == nil {
		return StringInput("ActorDefinition.Url", nil)
	}
	return StringInput("ActorDefinition.Url", resource.Url)
}
func (resource *ActorDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("ActorDefinition.Version", nil)
	}
	return StringInput("ActorDefinition.Version", resource.Version)
}
func (resource *ActorDefinition) T_Name() templ.Component {

	if resource == nil {
		return StringInput("ActorDefinition.Name", nil)
	}
	return StringInput("ActorDefinition.Name", resource.Name)
}
func (resource *ActorDefinition) T_Title() templ.Component {

	if resource == nil {
		return StringInput("ActorDefinition.Title", nil)
	}
	return StringInput("ActorDefinition.Title", resource.Title)
}
func (resource *ActorDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("ActorDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("ActorDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *ActorDefinition) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("ActorDefinition.Experimental", nil)
	}
	return BoolInput("ActorDefinition.Experimental", resource.Experimental)
}
func (resource *ActorDefinition) T_Date() templ.Component {

	if resource == nil {
		return StringInput("ActorDefinition.Date", nil)
	}
	return StringInput("ActorDefinition.Date", resource.Date)
}
func (resource *ActorDefinition) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("ActorDefinition.Publisher", nil)
	}
	return StringInput("ActorDefinition.Publisher", resource.Publisher)
}
func (resource *ActorDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("ActorDefinition.Description", nil)
	}
	return StringInput("ActorDefinition.Description", resource.Description)
}
func (resource *ActorDefinition) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("ActorDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ActorDefinition.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *ActorDefinition) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("ActorDefinition.Purpose", nil)
	}
	return StringInput("ActorDefinition.Purpose", resource.Purpose)
}
func (resource *ActorDefinition) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("ActorDefinition.Copyright", nil)
	}
	return StringInput("ActorDefinition.Copyright", resource.Copyright)
}
func (resource *ActorDefinition) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("ActorDefinition.CopyrightLabel", nil)
	}
	return StringInput("ActorDefinition.CopyrightLabel", resource.CopyrightLabel)
}
func (resource *ActorDefinition) T_Type() templ.Component {
	optionsValueSet := VSExamplescenario_actor_type

	if resource == nil {
		return CodeSelect("ActorDefinition.Type", nil, optionsValueSet)
	}
	return CodeSelect("ActorDefinition.Type", &resource.Type, optionsValueSet)
}
func (resource *ActorDefinition) T_Documentation() templ.Component {

	if resource == nil {
		return StringInput("ActorDefinition.Documentation", nil)
	}
	return StringInput("ActorDefinition.Documentation", resource.Documentation)
}
func (resource *ActorDefinition) T_Reference(numReference int) templ.Component {

	if resource == nil || len(resource.Reference) >= numReference {
		return StringInput("ActorDefinition.Reference["+strconv.Itoa(numReference)+"]", nil)
	}
	return StringInput("ActorDefinition.Reference["+strconv.Itoa(numReference)+"]", &resource.Reference[numReference])
}
func (resource *ActorDefinition) T_Capabilities() templ.Component {

	if resource == nil {
		return StringInput("ActorDefinition.Capabilities", nil)
	}
	return StringInput("ActorDefinition.Capabilities", resource.Capabilities)
}
func (resource *ActorDefinition) T_DerivedFrom(numDerivedFrom int) templ.Component {

	if resource == nil || len(resource.DerivedFrom) >= numDerivedFrom {
		return StringInput("ActorDefinition.DerivedFrom["+strconv.Itoa(numDerivedFrom)+"]", nil)
	}
	return StringInput("ActorDefinition.DerivedFrom["+strconv.Itoa(numDerivedFrom)+"]", &resource.DerivedFrom[numDerivedFrom])
}
