package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/List
type List struct {
	Id                *string          `json:"id,omitempty"`
	Meta              *Meta            `json:"meta,omitempty"`
	ImplicitRules     *string          `json:"implicitRules,omitempty"`
	Language          *string          `json:"language,omitempty"`
	Text              *Narrative       `json:"text,omitempty"`
	Contained         []Resource       `json:"contained,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Identifier        []Identifier     `json:"identifier,omitempty"`
	Status            string           `json:"status"`
	Mode              string           `json:"mode"`
	Title             *string          `json:"title,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Subject           *Reference       `json:"subject,omitempty"`
	Encounter         *Reference       `json:"encounter,omitempty"`
	Date              *string          `json:"date,omitempty"`
	Source            *Reference       `json:"source,omitempty"`
	OrderedBy         *CodeableConcept `json:"orderedBy,omitempty"`
	Note              []Annotation     `json:"note,omitempty"`
	Entry             []ListEntry      `json:"entry,omitempty"`
	EmptyReason       *CodeableConcept `json:"emptyReason,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/List
type ListEntry struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Flag              *CodeableConcept `json:"flag,omitempty"`
	Deleted           *bool            `json:"deleted,omitempty"`
	Date              *string          `json:"date,omitempty"`
	Item              Reference        `json:"item"`
}

type OtherList List

// on convert struct to json, automatically add resourceType=List
func (r List) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherList
		ResourceType string `json:"resourceType"`
	}{
		OtherList:    OtherList(r),
		ResourceType: "List",
	})
}

func (resource *List) T_Id() templ.Component {

	if resource == nil {
		return StringInput("List.Id", nil)
	}
	return StringInput("List.Id", resource.Id)
}
func (resource *List) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("List.ImplicitRules", nil)
	}
	return StringInput("List.ImplicitRules", resource.ImplicitRules)
}
func (resource *List) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("List.Language", nil, optionsValueSet)
	}
	return CodeSelect("List.Language", resource.Language, optionsValueSet)
}
func (resource *List) T_Status() templ.Component {
	optionsValueSet := VSList_status

	if resource == nil {
		return CodeSelect("List.Status", nil, optionsValueSet)
	}
	return CodeSelect("List.Status", &resource.Status, optionsValueSet)
}
func (resource *List) T_Mode() templ.Component {
	optionsValueSet := VSList_mode

	if resource == nil {
		return CodeSelect("List.Mode", nil, optionsValueSet)
	}
	return CodeSelect("List.Mode", &resource.Mode, optionsValueSet)
}
func (resource *List) T_Title() templ.Component {

	if resource == nil {
		return StringInput("List.Title", nil)
	}
	return StringInput("List.Title", resource.Title)
}
func (resource *List) T_Code(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("List.Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("List.Code", resource.Code, optionsValueSet)
}
func (resource *List) T_Date() templ.Component {

	if resource == nil {
		return StringInput("List.Date", nil)
	}
	return StringInput("List.Date", resource.Date)
}
func (resource *List) T_OrderedBy(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("List.OrderedBy", nil, optionsValueSet)
	}
	return CodeableConceptSelect("List.OrderedBy", resource.OrderedBy, optionsValueSet)
}
func (resource *List) T_EmptyReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("List.EmptyReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("List.EmptyReason", resource.EmptyReason, optionsValueSet)
}
func (resource *List) T_EntryId(numEntry int) templ.Component {

	if resource == nil || len(resource.Entry) >= numEntry {
		return StringInput("List.Entry["+strconv.Itoa(numEntry)+"].Id", nil)
	}
	return StringInput("List.Entry["+strconv.Itoa(numEntry)+"].Id", resource.Entry[numEntry].Id)
}
func (resource *List) T_EntryFlag(numEntry int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Entry) >= numEntry {
		return CodeableConceptSelect("List.Entry["+strconv.Itoa(numEntry)+"].Flag", nil, optionsValueSet)
	}
	return CodeableConceptSelect("List.Entry["+strconv.Itoa(numEntry)+"].Flag", resource.Entry[numEntry].Flag, optionsValueSet)
}
func (resource *List) T_EntryDeleted(numEntry int) templ.Component {

	if resource == nil || len(resource.Entry) >= numEntry {
		return BoolInput("List.Entry["+strconv.Itoa(numEntry)+"].Deleted", nil)
	}
	return BoolInput("List.Entry["+strconv.Itoa(numEntry)+"].Deleted", resource.Entry[numEntry].Deleted)
}
func (resource *List) T_EntryDate(numEntry int) templ.Component {

	if resource == nil || len(resource.Entry) >= numEntry {
		return StringInput("List.Entry["+strconv.Itoa(numEntry)+"].Date", nil)
	}
	return StringInput("List.Entry["+strconv.Itoa(numEntry)+"].Date", resource.Entry[numEntry].Date)
}
