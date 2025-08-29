package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/Group
type Group struct {
	Id                *string               `json:"id,omitempty"`
	Meta              *Meta                 `json:"meta,omitempty"`
	ImplicitRules     *string               `json:"implicitRules,omitempty"`
	Language          *string               `json:"language,omitempty"`
	Text              *Narrative            `json:"text,omitempty"`
	Contained         []Resource            `json:"contained,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `json:"identifier,omitempty"`
	Active            *bool                 `json:"active,omitempty"`
	Type              string                `json:"type"`
	Actual            bool                  `json:"actual"`
	Code              *CodeableConcept      `json:"code,omitempty"`
	Name              *string               `json:"name,omitempty"`
	Quantity          *int                  `json:"quantity,omitempty"`
	ManagingEntity    *Reference            `json:"managingEntity,omitempty"`
	Characteristic    []GroupCharacteristic `json:"characteristic,omitempty"`
	Member            []GroupMember         `json:"member,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Group
type GroupCharacteristic struct {
	Id                   *string         `json:"id,omitempty"`
	Extension            []Extension     `json:"extension,omitempty"`
	ModifierExtension    []Extension     `json:"modifierExtension,omitempty"`
	Code                 CodeableConcept `json:"code"`
	ValueCodeableConcept CodeableConcept `json:"valueCodeableConcept"`
	ValueBoolean         bool            `json:"valueBoolean"`
	ValueQuantity        Quantity        `json:"valueQuantity"`
	ValueRange           Range           `json:"valueRange"`
	ValueReference       Reference       `json:"valueReference"`
	Exclude              bool            `json:"exclude"`
	Period               *Period         `json:"period,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Group
type GroupMember struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Entity            Reference   `json:"entity"`
	Period            *Period     `json:"period,omitempty"`
	Inactive          *bool       `json:"inactive,omitempty"`
}

type OtherGroup Group

// on convert struct to json, automatically add resourceType=Group
func (r Group) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGroup
		ResourceType string `json:"resourceType"`
	}{
		OtherGroup:   OtherGroup(r),
		ResourceType: "Group",
	})
}

func (resource *Group) GroupLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Group) GroupType() templ.Component {
	optionsValueSet := VSGroup_type
	currentVal := ""
	if resource != nil {
		currentVal = resource.Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
