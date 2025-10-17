package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Group
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
	Membership        string                `json:"membership"`
	Code              *CodeableConcept      `json:"code,omitempty"`
	Name              *string               `json:"name,omitempty"`
	Description       *string               `json:"description,omitempty"`
	Quantity          *int                  `json:"quantity,omitempty"`
	ManagingEntity    *Reference            `json:"managingEntity,omitempty"`
	Characteristic    []GroupCharacteristic `json:"characteristic,omitempty"`
	Member            []GroupMember         `json:"member,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Group
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

// http://hl7.org/fhir/r5/StructureDefinition/Group
type GroupMember struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Entity            Reference   `json:"entity"`
	Period            *Period     `json:"period,omitempty"`
	Inactive          *bool       `json:"inactive,omitempty"`
}

type OtherGroup Group

// struct -> json, automatically add resourceType=Patient
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
func (r Group) ResourceType() string {
	return "Group"
}

func (resource *Group) T_Active(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("active", nil, htmlAttrs)
	}
	return BoolInput("active", resource.Active, htmlAttrs)
}
func (resource *Group) T_Type(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSGroup_type

	if resource == nil {
		return CodeSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Group) T_Membership(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSGroup_membership_basis

	if resource == nil {
		return CodeSelect("membership", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("membership", &resource.Membership, optionsValueSet, htmlAttrs)
}
func (resource *Group) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *Group) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *Group) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *Group) T_Quantity(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("quantity", nil, htmlAttrs)
	}
	return IntInput("quantity", resource.Quantity, htmlAttrs)
}
func (resource *Group) T_ManagingEntity(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "managingEntity", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "managingEntity", resource.ManagingEntity, htmlAttrs)
}
func (resource *Group) T_CharacteristicCode(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].code", &resource.Characteristic[numCharacteristic].Code, optionsValueSet, htmlAttrs)
}
func (resource *Group) T_CharacteristicValueCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].valueCodeableConcept", &resource.Characteristic[numCharacteristic].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Group) T_CharacteristicValueBoolean(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueBoolean", &resource.Characteristic[numCharacteristic].ValueBoolean, htmlAttrs)
}
func (resource *Group) T_CharacteristicValueQuantity(numCharacteristic int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return QuantityInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueQuantity", &resource.Characteristic[numCharacteristic].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *Group) T_CharacteristicValueRange(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return RangeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].valueRange", &resource.Characteristic[numCharacteristic].ValueRange, htmlAttrs)
}
func (resource *Group) T_CharacteristicValueReference(frs []FhirResource, numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return ReferenceInput(frs, "characteristic["+strconv.Itoa(numCharacteristic)+"].valueReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "characteristic["+strconv.Itoa(numCharacteristic)+"].valueReference", &resource.Characteristic[numCharacteristic].ValueReference, htmlAttrs)
}
func (resource *Group) T_CharacteristicExclude(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].exclude", nil, htmlAttrs)
	}
	return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].exclude", &resource.Characteristic[numCharacteristic].Exclude, htmlAttrs)
}
func (resource *Group) T_CharacteristicPeriod(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return PeriodInput("characteristic["+strconv.Itoa(numCharacteristic)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("characteristic["+strconv.Itoa(numCharacteristic)+"].period", resource.Characteristic[numCharacteristic].Period, htmlAttrs)
}
func (resource *Group) T_MemberEntity(frs []FhirResource, numMember int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMember >= len(resource.Member) {
		return ReferenceInput(frs, "member["+strconv.Itoa(numMember)+"].entity", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "member["+strconv.Itoa(numMember)+"].entity", &resource.Member[numMember].Entity, htmlAttrs)
}
func (resource *Group) T_MemberPeriod(numMember int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMember >= len(resource.Member) {
		return PeriodInput("member["+strconv.Itoa(numMember)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("member["+strconv.Itoa(numMember)+"].period", resource.Member[numMember].Period, htmlAttrs)
}
func (resource *Group) T_MemberInactive(numMember int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMember >= len(resource.Member) {
		return BoolInput("member["+strconv.Itoa(numMember)+"].inactive", nil, htmlAttrs)
	}
	return BoolInput("member["+strconv.Itoa(numMember)+"].inactive", resource.Member[numMember].Inactive, htmlAttrs)
}
