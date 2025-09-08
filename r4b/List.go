package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/List
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
	Date              *time.Time       `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Source            *Reference       `json:"source,omitempty"`
	OrderedBy         *CodeableConcept `json:"orderedBy,omitempty"`
	Note              []Annotation     `json:"note,omitempty"`
	Entry             []ListEntry      `json:"entry,omitempty"`
	EmptyReason       *CodeableConcept `json:"emptyReason,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/List
type ListEntry struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Flag              *CodeableConcept `json:"flag,omitempty"`
	Deleted           *bool            `json:"deleted,omitempty"`
	Date              *time.Time       `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r List) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "List/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "List"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *List) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSList_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *List) T_Mode(htmlAttrs string) templ.Component {
	optionsValueSet := VSList_mode

	if resource == nil {
		return CodeSelect("Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Mode", &resource.Mode, optionsValueSet, htmlAttrs)
}
func (resource *List) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *List) T_Code(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *List) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *List) T_OrderedBy(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("OrderedBy", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OrderedBy", resource.OrderedBy, optionsValueSet, htmlAttrs)
}
func (resource *List) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *List) T_EmptyReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("EmptyReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EmptyReason", resource.EmptyReason, optionsValueSet, htmlAttrs)
}
func (resource *List) T_EntryFlag(numEntry int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return CodeableConceptSelect("Entry["+strconv.Itoa(numEntry)+"]Flag", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Entry["+strconv.Itoa(numEntry)+"]Flag", resource.Entry[numEntry].Flag, optionsValueSet, htmlAttrs)
}
func (resource *List) T_EntryDeleted(numEntry int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return BoolInput("Entry["+strconv.Itoa(numEntry)+"]Deleted", nil, htmlAttrs)
	}
	return BoolInput("Entry["+strconv.Itoa(numEntry)+"]Deleted", resource.Entry[numEntry].Deleted, htmlAttrs)
}
func (resource *List) T_EntryDate(numEntry int, htmlAttrs string) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return DateTimeInput("Entry["+strconv.Itoa(numEntry)+"]Date", nil, htmlAttrs)
	}
	return DateTimeInput("Entry["+strconv.Itoa(numEntry)+"]Date", resource.Entry[numEntry].Date, htmlAttrs)
}
