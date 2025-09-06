package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *CompartmentDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.Id", nil)
	}
	return StringInput("CompartmentDefinition.Id", resource.Id)
}
func (resource *CompartmentDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.ImplicitRules", nil)
	}
	return StringInput("CompartmentDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *CompartmentDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("CompartmentDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("CompartmentDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *CompartmentDefinition) T_Url() templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.Url", nil)
	}
	return StringInput("CompartmentDefinition.Url", &resource.Url)
}
func (resource *CompartmentDefinition) T_Version() templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.Version", nil)
	}
	return StringInput("CompartmentDefinition.Version", resource.Version)
}
func (resource *CompartmentDefinition) T_Name() templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.Name", nil)
	}
	return StringInput("CompartmentDefinition.Name", &resource.Name)
}
func (resource *CompartmentDefinition) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("CompartmentDefinition.Status", nil, optionsValueSet)
	}
	return CodeSelect("CompartmentDefinition.Status", &resource.Status, optionsValueSet)
}
func (resource *CompartmentDefinition) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("CompartmentDefinition.Experimental", nil)
	}
	return BoolInput("CompartmentDefinition.Experimental", resource.Experimental)
}
func (resource *CompartmentDefinition) T_Date() templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.Date", nil)
	}
	return StringInput("CompartmentDefinition.Date", resource.Date)
}
func (resource *CompartmentDefinition) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.Publisher", nil)
	}
	return StringInput("CompartmentDefinition.Publisher", resource.Publisher)
}
func (resource *CompartmentDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.Description", nil)
	}
	return StringInput("CompartmentDefinition.Description", resource.Description)
}
func (resource *CompartmentDefinition) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("CompartmentDefinition.Purpose", nil)
	}
	return StringInput("CompartmentDefinition.Purpose", resource.Purpose)
}
func (resource *CompartmentDefinition) T_Code() templ.Component {
	optionsValueSet := VSCompartment_type

	if resource == nil {
		return CodeSelect("CompartmentDefinition.Code", nil, optionsValueSet)
	}
	return CodeSelect("CompartmentDefinition.Code", &resource.Code, optionsValueSet)
}
func (resource *CompartmentDefinition) T_Search() templ.Component {

	if resource == nil {
		return BoolInput("CompartmentDefinition.Search", nil)
	}
	return BoolInput("CompartmentDefinition.Search", &resource.Search)
}
func (resource *CompartmentDefinition) T_ResourceId(numResource int) templ.Component {

	if resource == nil || len(resource.Resource) >= numResource {
		return StringInput("CompartmentDefinition.Resource["+strconv.Itoa(numResource)+"].Id", nil)
	}
	return StringInput("CompartmentDefinition.Resource["+strconv.Itoa(numResource)+"].Id", resource.Resource[numResource].Id)
}
func (resource *CompartmentDefinition) T_ResourceCode(numResource int) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || len(resource.Resource) >= numResource {
		return CodeSelect("CompartmentDefinition.Resource["+strconv.Itoa(numResource)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("CompartmentDefinition.Resource["+strconv.Itoa(numResource)+"].Code", &resource.Resource[numResource].Code, optionsValueSet)
}
func (resource *CompartmentDefinition) T_ResourceParam(numResource int, numParam int) templ.Component {

	if resource == nil || len(resource.Resource) >= numResource || len(resource.Resource[numResource].Param) >= numParam {
		return StringInput("CompartmentDefinition.Resource["+strconv.Itoa(numResource)+"].Param["+strconv.Itoa(numParam)+"]", nil)
	}
	return StringInput("CompartmentDefinition.Resource["+strconv.Itoa(numResource)+"].Param["+strconv.Itoa(numParam)+"]", &resource.Resource[numResource].Param[numParam])
}
func (resource *CompartmentDefinition) T_ResourceDocumentation(numResource int) templ.Component {

	if resource == nil || len(resource.Resource) >= numResource {
		return StringInput("CompartmentDefinition.Resource["+strconv.Itoa(numResource)+"].Documentation", nil)
	}
	return StringInput("CompartmentDefinition.Resource["+strconv.Itoa(numResource)+"].Documentation", resource.Resource[numResource].Documentation)
}
