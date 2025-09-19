package r4

//generated with command go run ./bultaoreune
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
	Date              *FhirDateTime    `json:"date,omitempty"`
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
	Date              *FhirDateTime    `json:"date,omitempty"`
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
func (resource *List) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSList_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *List) T_Mode(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSList_mode

	if resource == nil {
		return CodeSelect("mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("mode", &resource.Mode, optionsValueSet, htmlAttrs)
}
func (resource *List) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *List) T_Code(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet, htmlAttrs)
}
func (resource *List) T_Subject(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("subject", nil, htmlAttrs)
	}
	return ReferenceInput("subject", resource.Subject, htmlAttrs)
}
func (resource *List) T_Encounter(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("encounter", nil, htmlAttrs)
	}
	return ReferenceInput("encounter", resource.Encounter, htmlAttrs)
}
func (resource *List) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *List) T_Source(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("source", nil, htmlAttrs)
	}
	return ReferenceInput("source", resource.Source, htmlAttrs)
}
func (resource *List) T_OrderedBy(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("orderedBy", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("orderedBy", resource.OrderedBy, optionsValueSet, htmlAttrs)
}
func (resource *List) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *List) T_EmptyReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("emptyReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("emptyReason", resource.EmptyReason, optionsValueSet, htmlAttrs)
}
func (resource *List) T_EntryFlag(numEntry int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return CodeableConceptSelect("entry["+strconv.Itoa(numEntry)+"].flag", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("entry["+strconv.Itoa(numEntry)+"].flag", resource.Entry[numEntry].Flag, optionsValueSet, htmlAttrs)
}
func (resource *List) T_EntryDeleted(numEntry int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return BoolInput("entry["+strconv.Itoa(numEntry)+"].deleted", nil, htmlAttrs)
	}
	return BoolInput("entry["+strconv.Itoa(numEntry)+"].deleted", resource.Entry[numEntry].Deleted, htmlAttrs)
}
func (resource *List) T_EntryDate(numEntry int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return FhirDateTimeInput("entry["+strconv.Itoa(numEntry)+"].date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("entry["+strconv.Itoa(numEntry)+"].date", resource.Entry[numEntry].Date, htmlAttrs)
}
func (resource *List) T_EntryItem(numEntry int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEntry >= len(resource.Entry) {
		return ReferenceInput("entry["+strconv.Itoa(numEntry)+"].item", nil, htmlAttrs)
	}
	return ReferenceInput("entry["+strconv.Itoa(numEntry)+"].item", &resource.Entry[numEntry].Item, htmlAttrs)
}
