package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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

func (resource *NamingSystem) T_Id() templ.Component {

	if resource == nil {
		return StringInput("NamingSystem.Id", nil)
	}
	return StringInput("NamingSystem.Id", resource.Id)
}
func (resource *NamingSystem) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("NamingSystem.ImplicitRules", nil)
	}
	return StringInput("NamingSystem.ImplicitRules", resource.ImplicitRules)
}
func (resource *NamingSystem) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("NamingSystem.Language", nil, optionsValueSet)
	}
	return CodeSelect("NamingSystem.Language", resource.Language, optionsValueSet)
}
func (resource *NamingSystem) T_Name() templ.Component {

	if resource == nil {
		return StringInput("NamingSystem.Name", nil)
	}
	return StringInput("NamingSystem.Name", &resource.Name)
}
func (resource *NamingSystem) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("NamingSystem.Status", nil, optionsValueSet)
	}
	return CodeSelect("NamingSystem.Status", &resource.Status, optionsValueSet)
}
func (resource *NamingSystem) T_Kind() templ.Component {
	optionsValueSet := VSNamingsystem_type

	if resource == nil {
		return CodeSelect("NamingSystem.Kind", nil, optionsValueSet)
	}
	return CodeSelect("NamingSystem.Kind", &resource.Kind, optionsValueSet)
}
func (resource *NamingSystem) T_Date() templ.Component {

	if resource == nil {
		return StringInput("NamingSystem.Date", nil)
	}
	return StringInput("NamingSystem.Date", &resource.Date)
}
func (resource *NamingSystem) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("NamingSystem.Publisher", nil)
	}
	return StringInput("NamingSystem.Publisher", resource.Publisher)
}
func (resource *NamingSystem) T_Responsible() templ.Component {

	if resource == nil {
		return StringInput("NamingSystem.Responsible", nil)
	}
	return StringInput("NamingSystem.Responsible", resource.Responsible)
}
func (resource *NamingSystem) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("NamingSystem.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NamingSystem.Type", resource.Type, optionsValueSet)
}
func (resource *NamingSystem) T_Description() templ.Component {

	if resource == nil {
		return StringInput("NamingSystem.Description", nil)
	}
	return StringInput("NamingSystem.Description", resource.Description)
}
func (resource *NamingSystem) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("NamingSystem.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("NamingSystem.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *NamingSystem) T_Usage() templ.Component {

	if resource == nil {
		return StringInput("NamingSystem.Usage", nil)
	}
	return StringInput("NamingSystem.Usage", resource.Usage)
}
func (resource *NamingSystem) T_UniqueIdId(numUniqueId int) templ.Component {

	if resource == nil || len(resource.UniqueId) >= numUniqueId {
		return StringInput("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Id", nil)
	}
	return StringInput("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Id", resource.UniqueId[numUniqueId].Id)
}
func (resource *NamingSystem) T_UniqueIdType(numUniqueId int) templ.Component {
	optionsValueSet := VSNamingsystem_identifier_type

	if resource == nil || len(resource.UniqueId) >= numUniqueId {
		return CodeSelect("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Type", &resource.UniqueId[numUniqueId].Type, optionsValueSet)
}
func (resource *NamingSystem) T_UniqueIdValue(numUniqueId int) templ.Component {

	if resource == nil || len(resource.UniqueId) >= numUniqueId {
		return StringInput("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Value", nil)
	}
	return StringInput("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Value", &resource.UniqueId[numUniqueId].Value)
}
func (resource *NamingSystem) T_UniqueIdPreferred(numUniqueId int) templ.Component {

	if resource == nil || len(resource.UniqueId) >= numUniqueId {
		return BoolInput("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Preferred", nil)
	}
	return BoolInput("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Preferred", resource.UniqueId[numUniqueId].Preferred)
}
func (resource *NamingSystem) T_UniqueIdComment(numUniqueId int) templ.Component {

	if resource == nil || len(resource.UniqueId) >= numUniqueId {
		return StringInput("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Comment", nil)
	}
	return StringInput("NamingSystem.UniqueId["+strconv.Itoa(numUniqueId)+"].Comment", resource.UniqueId[numUniqueId].Comment)
}
