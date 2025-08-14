//generated August 14 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

// http://hl7.org/fhir/r5/StructureDefinition/Composition
type Composition struct {
	Id                *string               `json:"id,omitempty"`
	Meta              *Meta                 `json:"meta,omitempty"`
	ImplicitRules     *string               `json:"implicitRules,omitempty"`
	Language          *string               `json:"language,omitempty"`
	Text              *Narrative            `json:"text,omitempty"`
	Contained         []Resource            `json:"contained,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Url               *string               `json:"url,omitempty"`
	Identifier        []Identifier          `json:"identifier,omitempty"`
	Version           *string               `json:"version,omitempty"`
	Status            string                `json:"status"`
	Type              CodeableConcept       `json:"type"`
	Category          []CodeableConcept     `json:"category,omitempty"`
	Subject           []Reference           `json:"subject,omitempty"`
	Encounter         *Reference            `json:"encounter,omitempty"`
	Date              string                `json:"date"`
	UseContext        []UsageContext        `json:"useContext,omitempty"`
	Author            []Reference           `json:"author"`
	Name              *string               `json:"name,omitempty"`
	Title             string                `json:"title"`
	Note              []Annotation          `json:"note,omitempty"`
	Attester          []CompositionAttester `json:"attester,omitempty"`
	Custodian         *Reference            `json:"custodian,omitempty"`
	RelatesTo         []RelatedArtifact     `json:"relatesTo,omitempty"`
	Event             []CompositionEvent    `json:"event,omitempty"`
	Section           []CompositionSection  `json:"section,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Composition
type CompositionAttester struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Mode              CodeableConcept `json:"mode"`
	Time              *string         `json:"time,omitempty"`
	Party             *Reference      `json:"party,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Composition
type CompositionEvent struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Period            *Period             `json:"period,omitempty"`
	Detail            []CodeableReference `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Composition
type CompositionSection struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Title             *string          `json:"title,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Author            []Reference      `json:"author,omitempty"`
	Focus             *Reference       `json:"focus,omitempty"`
	Text              *Narrative       `json:"text,omitempty"`
	OrderedBy         *CodeableConcept `json:"orderedBy,omitempty"`
	Entry             []Reference      `json:"entry,omitempty"`
	EmptyReason       *CodeableConcept `json:"emptyReason,omitempty"`
}

type OtherComposition Composition

// on convert struct to json, automatically add resourceType=Composition
func (r Composition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherComposition
		ResourceType string `json:"resourceType"`
	}{
		OtherComposition: OtherComposition(r),
		ResourceType:     "Composition",
	})
}
