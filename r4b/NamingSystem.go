package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/NamingSystem
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
	Date              string                 `json:"date"`
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

// http://hl7.org/fhir/r4b/StructureDefinition/NamingSystem
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
func (resource *NamingSystem) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", &resource.Name, htmlAttrs)
}
func (resource *NamingSystem) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Kind(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSNamingsystem_type

	if resource == nil {
		return CodeSelect("kind", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("kind", &resource.Kind, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", &resource.Date, htmlAttrs)
}
func (resource *NamingSystem) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *NamingSystem) T_Responsible(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("responsible", nil, htmlAttrs)
	}
	return StringInput("responsible", resource.Responsible, htmlAttrs)
}
func (resource *NamingSystem) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *NamingSystem) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_Usage(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("usage", nil, htmlAttrs)
	}
	return StringInput("usage", resource.Usage, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdType(numUniqueId int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSNamingsystem_identifier_type

	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return CodeSelect("uniqueId["+strconv.Itoa(numUniqueId)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("uniqueId["+strconv.Itoa(numUniqueId)+"].type", &resource.UniqueId[numUniqueId].Type, optionsValueSet, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdValue(numUniqueId int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return StringInput("uniqueId["+strconv.Itoa(numUniqueId)+"].value", nil, htmlAttrs)
	}
	return StringInput("uniqueId["+strconv.Itoa(numUniqueId)+"].value", &resource.UniqueId[numUniqueId].Value, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdPreferred(numUniqueId int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return BoolInput("uniqueId["+strconv.Itoa(numUniqueId)+"].preferred", nil, htmlAttrs)
	}
	return BoolInput("uniqueId["+strconv.Itoa(numUniqueId)+"].preferred", resource.UniqueId[numUniqueId].Preferred, htmlAttrs)
}
func (resource *NamingSystem) T_UniqueIdComment(numUniqueId int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUniqueId >= len(resource.UniqueId) {
		return StringInput("uniqueId["+strconv.Itoa(numUniqueId)+"].comment", nil, htmlAttrs)
	}
	return StringInput("uniqueId["+strconv.Itoa(numUniqueId)+"].comment", resource.UniqueId[numUniqueId].Comment, htmlAttrs)
}
