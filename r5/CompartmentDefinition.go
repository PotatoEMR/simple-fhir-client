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
	VersionAlgorithmString *string                         `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                         `json:"versionAlgorithmCoding,omitempty"`
	Name                   string                          `json:"name"`
	Title                  *string                         `json:"title,omitempty"`
	Status                 string                          `json:"status"`
	Experimental           *bool                           `json:"experimental,omitempty"`
	Date                   *time.Time                      `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r CompartmentDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "CompartmentDefinition/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "CompartmentDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *CompartmentDefinition) T_Url(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.Url", nil, htmlAttrs)
	}
	return StringInput("CompartmentDefinition.Url", &resource.Url, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.Version", nil, htmlAttrs)
	}
	return StringInput("CompartmentDefinition.Version", resource.Version, htmlAttrs)
}
func (resource *CompartmentDefinition) T_VersionAlgorithmString(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("CompartmentDefinition.VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *CompartmentDefinition) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodingSelect("CompartmentDefinition.VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("CompartmentDefinition.VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.Name", nil, htmlAttrs)
	}
	return StringInput("CompartmentDefinition.Name", &resource.Name, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.Title", nil, htmlAttrs)
	}
	return StringInput("CompartmentDefinition.Title", resource.Title, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("CompartmentDefinition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CompartmentDefinition.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Experimental(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("CompartmentDefinition.Experimental", nil, htmlAttrs)
	}
	return BoolInput("CompartmentDefinition.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Date(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("CompartmentDefinition.Date", nil, htmlAttrs)
	}
	return DateTimeInput("CompartmentDefinition.Date", resource.Date, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Publisher(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.Publisher", nil, htmlAttrs)
	}
	return StringInput("CompartmentDefinition.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("CompartmentDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Purpose(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.Purpose", nil, htmlAttrs)
	}
	return StringInput("CompartmentDefinition.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Code(htmlAttrs string) templ.Component {
	optionsValueSet := VSCompartment_type

	if resource == nil {
		return CodeSelect("CompartmentDefinition.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CompartmentDefinition.Code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Search(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("CompartmentDefinition.Search", nil, htmlAttrs)
	}
	return BoolInput("CompartmentDefinition.Search", &resource.Search, htmlAttrs)
}
func (resource *CompartmentDefinition) T_ResourceCode(numResource int, htmlAttrs string) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numResource >= len(resource.Resource) {
		return CodeSelect("CompartmentDefinition.Resource."+strconv.Itoa(numResource)+"..Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("CompartmentDefinition.Resource."+strconv.Itoa(numResource)+"..Code", &resource.Resource[numResource].Code, optionsValueSet, htmlAttrs)
}
func (resource *CompartmentDefinition) T_ResourceParam(numResource int, numParam int, htmlAttrs string) templ.Component {

	if resource == nil || numResource >= len(resource.Resource) || numParam >= len(resource.Resource[numResource].Param) {
		return StringInput("CompartmentDefinition.Resource."+strconv.Itoa(numResource)+"..Param."+strconv.Itoa(numParam)+".", nil, htmlAttrs)
	}
	return StringInput("CompartmentDefinition.Resource."+strconv.Itoa(numResource)+"..Param."+strconv.Itoa(numParam)+".", &resource.Resource[numResource].Param[numParam], htmlAttrs)
}
func (resource *CompartmentDefinition) T_ResourceDocumentation(numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numResource >= len(resource.Resource) {
		return StringInput("CompartmentDefinition.Resource."+strconv.Itoa(numResource)+"..Documentation", nil, htmlAttrs)
	}
	return StringInput("CompartmentDefinition.Resource."+strconv.Itoa(numResource)+"..Documentation", resource.Resource[numResource].Documentation, htmlAttrs)
}
func (resource *CompartmentDefinition) T_ResourceStartParam(numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numResource >= len(resource.Resource) {
		return StringInput("CompartmentDefinition.Resource."+strconv.Itoa(numResource)+"..StartParam", nil, htmlAttrs)
	}
	return StringInput("CompartmentDefinition.Resource."+strconv.Itoa(numResource)+"..StartParam", resource.Resource[numResource].StartParam, htmlAttrs)
}
func (resource *CompartmentDefinition) T_ResourceEndParam(numResource int, htmlAttrs string) templ.Component {

	if resource == nil || numResource >= len(resource.Resource) {
		return StringInput("CompartmentDefinition.Resource."+strconv.Itoa(numResource)+"..EndParam", nil, htmlAttrs)
	}
	return StringInput("CompartmentDefinition.Resource."+strconv.Itoa(numResource)+"..EndParam", resource.Resource[numResource].EndParam, htmlAttrs)
}
