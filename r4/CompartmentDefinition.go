package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Date              *time.Time                      `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", &resource.Url, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", &resource.Name, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Experimental", nil, htmlAttrs)
	}
	return BoolInput("Experimental", resource.Experimental, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Purpose", nil, htmlAttrs)
	}
	return StringInput("Purpose", resource.Purpose, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Code(htmlAttrs string) templ.Component {
	optionsValueSet := VSCompartment_type

	if resource == nil {
		return CodeSelect("Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Code", &resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *CompartmentDefinition) T_Search(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Search", nil, htmlAttrs)
	}
	return BoolInput("Search", &resource.Search, htmlAttrs)
}
func (resource *CompartmentDefinition) T_ResourceCode(numResource int, htmlAttrs string) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numResource >= len(resource.Resource) {
		return CodeSelect("Resource["+strconv.Itoa(numResource)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Resource["+strconv.Itoa(numResource)+"]Code", &resource.Resource[numResource].Code, optionsValueSet, htmlAttrs)
}
func (resource *CompartmentDefinition) T_ResourceParam(numResource int, numParam int, htmlAttrs string) templ.Component {
	if resource == nil || numResource >= len(resource.Resource) || numParam >= len(resource.Resource[numResource].Param) {
		return StringInput("Resource["+strconv.Itoa(numResource)+"]Param["+strconv.Itoa(numParam)+"]", nil, htmlAttrs)
	}
	return StringInput("Resource["+strconv.Itoa(numResource)+"]Param["+strconv.Itoa(numParam)+"]", &resource.Resource[numResource].Param[numParam], htmlAttrs)
}
func (resource *CompartmentDefinition) T_ResourceDocumentation(numResource int, htmlAttrs string) templ.Component {
	if resource == nil || numResource >= len(resource.Resource) {
		return StringInput("Resource["+strconv.Itoa(numResource)+"]Documentation", nil, htmlAttrs)
	}
	return StringInput("Resource["+strconv.Itoa(numResource)+"]Documentation", resource.Resource[numResource].Documentation, htmlAttrs)
}
