//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

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
	Date              *string          `json:"date,omitempty"`
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
