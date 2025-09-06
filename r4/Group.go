package r4

//generated with command go run ./bultaoreune -nodownload
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

func (resource *Group) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Group.Id", nil)
	}
	return StringInput("Group.Id", resource.Id)
}
func (resource *Group) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Group.ImplicitRules", nil)
	}
	return StringInput("Group.ImplicitRules", resource.ImplicitRules)
}
func (resource *Group) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Group.Language", nil, optionsValueSet)
	}
	return CodeSelect("Group.Language", resource.Language, optionsValueSet)
}
func (resource *Group) T_Active() templ.Component {

	if resource == nil {
		return BoolInput("Group.Active", nil)
	}
	return BoolInput("Group.Active", resource.Active)
}
func (resource *Group) T_Type() templ.Component {
	optionsValueSet := VSGroup_type

	if resource == nil {
		return CodeSelect("Group.Type", nil, optionsValueSet)
	}
	return CodeSelect("Group.Type", &resource.Type, optionsValueSet)
}
func (resource *Group) T_Actual() templ.Component {

	if resource == nil {
		return BoolInput("Group.Actual", nil)
	}
	return BoolInput("Group.Actual", &resource.Actual)
}
func (resource *Group) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Group.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Group.Code", resource.Code, optionsValueSet)
}
func (resource *Group) T_Name() templ.Component {

	if resource == nil {
		return StringInput("Group.Name", nil)
	}
	return StringInput("Group.Name", resource.Name)
}
func (resource *Group) T_Quantity() templ.Component {

	if resource == nil {
		return IntInput("Group.Quantity", nil)
	}
	return IntInput("Group.Quantity", resource.Quantity)
}
func (resource *Group) T_CharacteristicId(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return StringInput("Group.Characteristic["+strconv.Itoa(numCharacteristic)+"].Id", nil)
	}
	return StringInput("Group.Characteristic["+strconv.Itoa(numCharacteristic)+"].Id", resource.Characteristic[numCharacteristic].Id)
}
func (resource *Group) T_CharacteristicCode(numCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return CodeableConceptSelect("Group.Characteristic["+strconv.Itoa(numCharacteristic)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Group.Characteristic["+strconv.Itoa(numCharacteristic)+"].Code", &resource.Characteristic[numCharacteristic].Code, optionsValueSet)
}
func (resource *Group) T_CharacteristicExclude(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return BoolInput("Group.Characteristic["+strconv.Itoa(numCharacteristic)+"].Exclude", nil)
	}
	return BoolInput("Group.Characteristic["+strconv.Itoa(numCharacteristic)+"].Exclude", &resource.Characteristic[numCharacteristic].Exclude)
}
func (resource *Group) T_MemberId(numMember int) templ.Component {

	if resource == nil || len(resource.Member) >= numMember {
		return StringInput("Group.Member["+strconv.Itoa(numMember)+"].Id", nil)
	}
	return StringInput("Group.Member["+strconv.Itoa(numMember)+"].Id", resource.Member[numMember].Id)
}
func (resource *Group) T_MemberInactive(numMember int) templ.Component {

	if resource == nil || len(resource.Member) >= numMember {
		return BoolInput("Group.Member["+strconv.Itoa(numMember)+"].Inactive", nil)
	}
	return BoolInput("Group.Member["+strconv.Itoa(numMember)+"].Inactive", resource.Member[numMember].Inactive)
}
