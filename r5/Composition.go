package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

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
	Date              time.Time             `json:"date,format:'2006-01-02T15:04:05Z07:00'"`
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
	Time              *time.Time      `json:"time,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r Composition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Composition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Composition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Composition) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *Composition) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *Composition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSComposition_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", &resource.Date, htmlAttrs)
}
func (resource *Composition) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *Composition) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", &resource.Title, htmlAttrs)
}
func (resource *Composition) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Composition) T_AttesterMode(numAttester int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAttester >= len(resource.Attester) {
		return CodeableConceptSelect("Attester["+strconv.Itoa(numAttester)+"]Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Attester["+strconv.Itoa(numAttester)+"]Mode", &resource.Attester[numAttester].Mode, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_AttesterTime(numAttester int, htmlAttrs string) templ.Component {
	if resource == nil || numAttester >= len(resource.Attester) {
		return DateTimeInput("Attester["+strconv.Itoa(numAttester)+"]Time", nil, htmlAttrs)
	}
	return DateTimeInput("Attester["+strconv.Itoa(numAttester)+"]Time", resource.Attester[numAttester].Time, htmlAttrs)
}
func (resource *Composition) T_SectionTitle(numSection int, htmlAttrs string) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return StringInput("Section["+strconv.Itoa(numSection)+"]Title", nil, htmlAttrs)
	}
	return StringInput("Section["+strconv.Itoa(numSection)+"]Title", resource.Section[numSection].Title, htmlAttrs)
}
func (resource *Composition) T_SectionCode(numSection int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("Section["+strconv.Itoa(numSection)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Section["+strconv.Itoa(numSection)+"]Code", resource.Section[numSection].Code, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_SectionOrderedBy(numSection int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("Section["+strconv.Itoa(numSection)+"]OrderedBy", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Section["+strconv.Itoa(numSection)+"]OrderedBy", resource.Section[numSection].OrderedBy, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_SectionEmptyReason(numSection int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("Section["+strconv.Itoa(numSection)+"]EmptyReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Section["+strconv.Itoa(numSection)+"]EmptyReason", resource.Section[numSection].EmptyReason, optionsValueSet, htmlAttrs)
}
