package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Group
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

// http://hl7.org/fhir/r4/StructureDefinition/Group
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

// http://hl7.org/fhir/r4/StructureDefinition/Group
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
func (r Group) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Group/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Group"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Group) T_Active(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Active", nil, htmlAttrs)
	}
	return BoolInput("Active", resource.Active, htmlAttrs)
}
func (resource *Group) T_Type(htmlAttrs string) templ.Component {
	optionsValueSet := VSGroup_type

	if resource == nil {
		return CodeSelect("Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Group) T_Actual(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Actual", nil, htmlAttrs)
	}
	return BoolInput("Actual", &resource.Actual, htmlAttrs)
}
func (resource *Group) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Group) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *Group) T_Quantity(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("Quantity", nil, htmlAttrs)
	}
	return IntInput("Quantity", resource.Quantity, htmlAttrs)
}
func (resource *Group) T_CharacteristicCode(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]Code", &resource.Characteristic[numCharacteristic].Code, optionsValueSet, htmlAttrs)
}
func (resource *Group) T_CharacteristicValueCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueCodeableConcept", &resource.Characteristic[numCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Group) T_CharacteristicValueBoolean(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]ValueBoolean", &resource.Characteristic[numCharacteristic].ValueBoolean, htmlAttrs)
}
func (resource *Group) T_CharacteristicExclude(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]Exclude", nil, htmlAttrs)
	}
	return BoolInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]Exclude", &resource.Characteristic[numCharacteristic].Exclude, htmlAttrs)
}
func (resource *Group) T_MemberInactive(numMember int, htmlAttrs string) templ.Component {
	if resource == nil || numMember >= len(resource.Member) {
		return BoolInput("Member["+strconv.Itoa(numMember)+"]Inactive", nil, htmlAttrs)
	}
	return BoolInput("Member["+strconv.Itoa(numMember)+"]Inactive", resource.Member[numMember].Inactive, htmlAttrs)
}
