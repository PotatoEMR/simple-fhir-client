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

// http://hl7.org/fhir/r4/StructureDefinition/NamingSystem
type NamingSystem struct {
	Id                *string                `json:"id,omitempty"`
	Meta              *Meta                  `json:"meta,omitempty"`
	ImplicitRules     *string                `json:"implicitRules,omitempty"`
	Language          *string                `json:"language,omitempty"`
	Text              *Narrative             `json:"text,omitempty"`
	Contained         []Resource             `json:"contained,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Name              string                 `json:"name"`
	Status            string                 `json:"status"`
	Kind              string                 `json:"kind"`
	Date              time.Time              `json:"date,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher         *string                `json:"publisher,omitempty"`
	Contact           []ContactDetail        `json:"contact,omitempty"`
	Responsible       *string                `json:"responsible,omitempty"`
	Type              *CodeableConcept       `json:"type,omitempty"`
	Description       *string                `json:"description,omitempty"`
	UseContext        []UsageContext         `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept      `json:"jurisdiction,omitempty"`
	Usage             *string                `json:"usage,omitempty"`
	UniqueId          []NamingSystemUniqueId `json:"uniqueId"`
}

// http://hl7.org/fhir/r4/StructureDefinition/NamingSystem
type NamingSystemUniqueId struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	Value             string      `json:"value"`
	Preferred         *bool       `json:"preferred,omitempty"`
	Comment           *string     `json:"comment,omitempty"`
	Period            *Period     `json:"period,omitempty"`
}

type OtherNamingSystem NamingSystem

// on convert struct to json, automatically add resourceType=NamingSystem
func (r NamingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNamingSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherNamingSystem: OtherNamingSystem(r),
		ResourceType:      "NamingSystem",
	})
}
func (r NamingSystem) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "NamingSystem/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "NamingSystem"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *NamingSystem) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("NamingSystem.Name", nil, htmlAttrs)
	}
	return StringInput("NamingSystem.Name", &resource.Name, htmlAttrs)
}
func (resource *NamingSystem) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("NamingSystem.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("NamingSystem.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Kind(htmlAttrs string) templ.Component {
	optionsValueSet := VSNamingsystem_type

	if resource == nil {
		return CodeSelect("NamingSystem.Kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("NamingSystem.Kind", &resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("NamingSystem.Date", nil, htmlAttrs)
	}
	return DateTimeInput("NamingSystem.Date", &resource.Date, htmlAttrs)
}
func (resource *NamingSystem) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("NamingSystem.Publisher", nil, htmlAttrs)
	}
	return StringInput("NamingSystem.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *NamingSystem) T_Responsible(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("NamingSystem.Responsible", nil, htmlAttrs)
	}
	return StringInput("NamingSystem.Responsible", resource.Responsible, htmlAttrs)
}
func (resource *NamingSystem) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("NamingSystem.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NamingSystem.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("NamingSystem.Description", nil, htmlAttrs)
	}
	return StringInput("NamingSystem.Description", resource.Description, htmlAttrs)
}
func (resource *NamingSystem) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("NamingSystem.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("NamingSystem.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Usage(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("NamingSystem.Usage", nil, htmlAttrs)
	}
	return StringInput("NamingSystem.Usage", resource.Usage, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdType(numUniqueId int, htmlAttrs string) templ.Component {
	optionsValueSet := VSNamingsystem_identifier_type

	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return CodeSelect("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Type", &resource.UniqueId[numUniqueId].Type, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdValue(numUniqueId int, htmlAttrs string) templ.Component {
	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return StringInput("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Value", nil, htmlAttrs)
	}
	return StringInput("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Value", &resource.UniqueId[numUniqueId].Value, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdPreferred(numUniqueId int, htmlAttrs string) templ.Component {
	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return BoolInput("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Preferred", nil, htmlAttrs)
	}
	return BoolInput("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Preferred", resource.UniqueId[numUniqueId].Preferred, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdComment(numUniqueId int, htmlAttrs string) templ.Component {
	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return StringInput("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Comment", nil, htmlAttrs)
	}
	return StringInput("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Comment", resource.UniqueId[numUniqueId].Comment, htmlAttrs)
}
