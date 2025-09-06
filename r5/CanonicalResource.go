package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/CanonicalResource
type CanonicalResource struct {
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
}

type OtherCanonicalResource CanonicalResource

// on convert struct to json, automatically add resourceType=CanonicalResource
func (r CanonicalResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCanonicalResource
		ResourceType string `json:"resourceType"`
	}{
		OtherCanonicalResource: OtherCanonicalResource(r),
		ResourceType:           "CanonicalResource",
	})
}

func (resource *CanonicalResource) T_Id() templ.Component {

	if resource == nil {
		return StringInput("CanonicalResource.Id", nil)
	}
	return StringInput("CanonicalResource.Id", resource.Id)
}
func (resource *CanonicalResource) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("CanonicalResource.ImplicitRules", nil)
	}
	return StringInput("CanonicalResource.ImplicitRules", resource.ImplicitRules)
}
func (resource *CanonicalResource) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("CanonicalResource.Language", nil, optionsValueSet)
	}
	return CodeSelect("CanonicalResource.Language", resource.Language, optionsValueSet)
}
func (resource *CanonicalResource) T_Url() templ.Component {

	if resource == nil {
		return StringInput("CanonicalResource.Url", nil)
	}
	return StringInput("CanonicalResource.Url", resource.Url)
}
func (resource *CanonicalResource) T_Version() templ.Component {

	if resource == nil {
		return StringInput("CanonicalResource.Version", nil)
	}
	return StringInput("CanonicalResource.Version", resource.Version)
}
func (resource *CanonicalResource) T_Name() templ.Component {

	if resource == nil {
		return StringInput("CanonicalResource.Name", nil)
	}
	return StringInput("CanonicalResource.Name", resource.Name)
}
func (resource *CanonicalResource) T_Title() templ.Component {

	if resource == nil {
		return StringInput("CanonicalResource.Title", nil)
	}
	return StringInput("CanonicalResource.Title", resource.Title)
}
func (resource *CanonicalResource) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("CanonicalResource.Status", nil, optionsValueSet)
	}
	return CodeSelect("CanonicalResource.Status", &resource.Status, optionsValueSet)
}
func (resource *CanonicalResource) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("CanonicalResource.Experimental", nil)
	}
	return BoolInput("CanonicalResource.Experimental", resource.Experimental)
}
func (resource *CanonicalResource) T_Date() templ.Component {

	if resource == nil {
		return StringInput("CanonicalResource.Date", nil)
	}
	return StringInput("CanonicalResource.Date", resource.Date)
}
func (resource *CanonicalResource) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("CanonicalResource.Publisher", nil)
	}
	return StringInput("CanonicalResource.Publisher", resource.Publisher)
}
func (resource *CanonicalResource) T_Description() templ.Component {

	if resource == nil {
		return StringInput("CanonicalResource.Description", nil)
	}
	return StringInput("CanonicalResource.Description", resource.Description)
}
func (resource *CanonicalResource) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("CanonicalResource.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("CanonicalResource.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *CanonicalResource) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("CanonicalResource.Purpose", nil)
	}
	return StringInput("CanonicalResource.Purpose", resource.Purpose)
}
func (resource *CanonicalResource) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("CanonicalResource.Copyright", nil)
	}
	return StringInput("CanonicalResource.Copyright", resource.Copyright)
}
func (resource *CanonicalResource) T_CopyrightLabel() templ.Component {

	if resource == nil {
		return StringInput("CanonicalResource.CopyrightLabel", nil)
	}
	return StringInput("CanonicalResource.CopyrightLabel", resource.CopyrightLabel)
}
