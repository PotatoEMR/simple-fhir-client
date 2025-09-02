package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/List
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
	Subject           []Reference      `json:"subject,omitempty"`
	Encounter         *Reference       `json:"encounter,omitempty"`
	Date              *string          `json:"date,omitempty"`
	Source            *Reference       `json:"source,omitempty"`
	OrderedBy         *CodeableConcept `json:"orderedBy,omitempty"`
	Note              []Annotation     `json:"note,omitempty"`
	Entry             []ListEntry      `json:"entry,omitempty"`
	EmptyReason       *CodeableConcept `json:"emptyReason,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/List
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

func (resource *List) ListLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *List) ListStatus() templ.Component {
	optionsValueSet := VSList_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *List) ListMode() templ.Component {
	optionsValueSet := VSList_mode

	if resource == nil {
		return CodeSelect("mode", nil, optionsValueSet)
	}
	return CodeSelect("mode", &resource.Mode, optionsValueSet)
}
func (resource *List) ListCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.Code, optionsValueSet)
}
func (resource *List) ListOrderedBy(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("orderedBy", nil, optionsValueSet)
	}
	return CodeableConceptSelect("orderedBy", resource.OrderedBy, optionsValueSet)
}
func (resource *List) ListEmptyReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("emptyReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("emptyReason", resource.EmptyReason, optionsValueSet)
}
func (resource *List) ListEntryFlag(numEntry int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Entry) >= numEntry {
		return CodeableConceptSelect("flag", nil, optionsValueSet)
	}
	return CodeableConceptSelect("flag", resource.Entry[numEntry].Flag, optionsValueSet)
}
