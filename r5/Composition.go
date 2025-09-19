package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Date              FhirDateTime          `json:"date"`
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
	Time              *FhirDateTime   `json:"time,omitempty"`
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
func (resource *Composition) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *Composition) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *Composition) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSComposition_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_Subject(numSubject int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubject >= len(resource.Subject) {
		return ReferenceInput("subject["+strconv.Itoa(numSubject)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("subject["+strconv.Itoa(numSubject)+"]", &resource.Subject[numSubject], htmlAttrs)
}
func (resource *Composition) T_Encounter(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("encounter", nil, htmlAttrs)
	}
	return ReferenceInput("encounter", resource.Encounter, htmlAttrs)
}
func (resource *Composition) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", &resource.Date, htmlAttrs)
}
func (resource *Composition) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *Composition) T_Author(numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthor >= len(resource.Author) {
		return ReferenceInput("author["+strconv.Itoa(numAuthor)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("author["+strconv.Itoa(numAuthor)+"]", &resource.Author[numAuthor], htmlAttrs)
}
func (resource *Composition) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *Composition) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", &resource.Title, htmlAttrs)
}
func (resource *Composition) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Composition) T_Custodian(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("custodian", nil, htmlAttrs)
	}
	return ReferenceInput("custodian", resource.Custodian, htmlAttrs)
}
func (resource *Composition) T_RelatesTo(numRelatesTo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.RelatesTo) {
		return RelatedArtifactInput("relatesTo["+strconv.Itoa(numRelatesTo)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("relatesTo["+strconv.Itoa(numRelatesTo)+"]", &resource.RelatesTo[numRelatesTo], htmlAttrs)
}
func (resource *Composition) T_AttesterMode(numAttester int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAttester >= len(resource.Attester) {
		return CodeableConceptSelect("attester["+strconv.Itoa(numAttester)+"].mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("attester["+strconv.Itoa(numAttester)+"].mode", &resource.Attester[numAttester].Mode, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_AttesterTime(numAttester int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAttester >= len(resource.Attester) {
		return FhirDateTimeInput("attester["+strconv.Itoa(numAttester)+"].time", nil, htmlAttrs)
	}
	return FhirDateTimeInput("attester["+strconv.Itoa(numAttester)+"].time", resource.Attester[numAttester].Time, htmlAttrs)
}
func (resource *Composition) T_AttesterParty(numAttester int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAttester >= len(resource.Attester) {
		return ReferenceInput("attester["+strconv.Itoa(numAttester)+"].party", nil, htmlAttrs)
	}
	return ReferenceInput("attester["+strconv.Itoa(numAttester)+"].party", resource.Attester[numAttester].Party, htmlAttrs)
}
func (resource *Composition) T_EventPeriod(numEvent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return PeriodInput("event["+strconv.Itoa(numEvent)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("event["+strconv.Itoa(numEvent)+"].period", resource.Event[numEvent].Period, htmlAttrs)
}
func (resource *Composition) T_EventDetail(numEvent int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) || numDetail >= len(resource.Event[numEvent].Detail) {
		return CodeableReferenceInput("event["+strconv.Itoa(numEvent)+"].detail["+strconv.Itoa(numDetail)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("event["+strconv.Itoa(numEvent)+"].detail["+strconv.Itoa(numDetail)+"]", &resource.Event[numEvent].Detail[numDetail], htmlAttrs)
}
func (resource *Composition) T_SectionTitle(numSection int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return StringInput("section["+strconv.Itoa(numSection)+"].title", nil, htmlAttrs)
	}
	return StringInput("section["+strconv.Itoa(numSection)+"].title", resource.Section[numSection].Title, htmlAttrs)
}
func (resource *Composition) T_SectionCode(numSection int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].code", resource.Section[numSection].Code, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_SectionAuthor(numSection int, numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) || numAuthor >= len(resource.Section[numSection].Author) {
		return ReferenceInput("section["+strconv.Itoa(numSection)+"].author["+strconv.Itoa(numAuthor)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("section["+strconv.Itoa(numSection)+"].author["+strconv.Itoa(numAuthor)+"]", &resource.Section[numSection].Author[numAuthor], htmlAttrs)
}
func (resource *Composition) T_SectionFocus(numSection int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return ReferenceInput("section["+strconv.Itoa(numSection)+"].focus", nil, htmlAttrs)
	}
	return ReferenceInput("section["+strconv.Itoa(numSection)+"].focus", resource.Section[numSection].Focus, htmlAttrs)
}
func (resource *Composition) T_SectionOrderedBy(numSection int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].orderedBy", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].orderedBy", resource.Section[numSection].OrderedBy, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_SectionEntry(numSection int, numEntry int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) || numEntry >= len(resource.Section[numSection].Entry) {
		return ReferenceInput("section["+strconv.Itoa(numSection)+"].entry["+strconv.Itoa(numEntry)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("section["+strconv.Itoa(numSection)+"].entry["+strconv.Itoa(numEntry)+"]", &resource.Section[numSection].Entry[numEntry], htmlAttrs)
}
func (resource *Composition) T_SectionEmptyReason(numSection int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].emptyReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].emptyReason", resource.Section[numSection].EmptyReason, optionsValueSet, htmlAttrs)
}
