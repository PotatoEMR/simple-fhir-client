package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Composition
type Composition struct {
	Id                *string                `json:"id,omitempty"`
	Meta              *Meta                  `json:"meta,omitempty"`
	ImplicitRules     *string                `json:"implicitRules,omitempty"`
	Language          *string                `json:"language,omitempty"`
	Text              *Narrative             `json:"text,omitempty"`
	Contained         []Resource             `json:"contained,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Identifier        *Identifier            `json:"identifier,omitempty"`
	Status            string                 `json:"status"`
	Type              CodeableConcept        `json:"type"`
	Category          []CodeableConcept      `json:"category,omitempty"`
	Subject           *Reference             `json:"subject,omitempty"`
	Encounter         *Reference             `json:"encounter,omitempty"`
	Date              FhirDateTime           `json:"date"`
	Author            []Reference            `json:"author"`
	Title             string                 `json:"title"`
	Confidentiality   *string                `json:"confidentiality,omitempty"`
	Attester          []CompositionAttester  `json:"attester,omitempty"`
	Custodian         *Reference             `json:"custodian,omitempty"`
	RelatesTo         []CompositionRelatesTo `json:"relatesTo,omitempty"`
	Event             []CompositionEvent     `json:"event,omitempty"`
	Section           []CompositionSection   `json:"section,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Composition
type CompositionAttester struct {
	Id                *string       `json:"id,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	Mode              string        `json:"mode"`
	Time              *FhirDateTime `json:"time,omitempty"`
	Party             *Reference    `json:"party,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Composition
type CompositionRelatesTo struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	TargetIdentifier  Identifier  `json:"targetIdentifier"`
	TargetReference   Reference   `json:"targetReference"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Composition
type CompositionEvent struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `json:"code,omitempty"`
	Period            *Period           `json:"period,omitempty"`
	Detail            []Reference       `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Composition
type CompositionSection struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Title             *string          `json:"title,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Author            []Reference      `json:"author,omitempty"`
	Focus             *Reference       `json:"focus,omitempty"`
	Text              *Narrative       `json:"text,omitempty"`
	Mode              *string          `json:"mode,omitempty"`
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
	ref.Identifier = r.Identifier
	rtype := "Composition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
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
func (resource *Composition) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *Composition) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *Composition) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", &resource.Date, htmlAttrs)
}
func (resource *Composition) T_Author(frs []FhirResource, numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthor >= len(resource.Author) {
		return ReferenceInput(frs, "author["+strconv.Itoa(numAuthor)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "author["+strconv.Itoa(numAuthor)+"]", &resource.Author[numAuthor], htmlAttrs)
}
func (resource *Composition) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", &resource.Title, htmlAttrs)
}
func (resource *Composition) T_Confidentiality(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeSelect("confidentiality", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("confidentiality", resource.Confidentiality, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_Custodian(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "custodian", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "custodian", resource.Custodian, htmlAttrs)
}
func (resource *Composition) T_AttesterMode(numAttester int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSComposition_attestation_mode

	if resource == nil || numAttester >= len(resource.Attester) {
		return CodeSelect("attester["+strconv.Itoa(numAttester)+"].mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("attester["+strconv.Itoa(numAttester)+"].mode", &resource.Attester[numAttester].Mode, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_AttesterTime(numAttester int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAttester >= len(resource.Attester) {
		return FhirDateTimeInput("attester["+strconv.Itoa(numAttester)+"].time", nil, htmlAttrs)
	}
	return FhirDateTimeInput("attester["+strconv.Itoa(numAttester)+"].time", resource.Attester[numAttester].Time, htmlAttrs)
}
func (resource *Composition) T_AttesterParty(frs []FhirResource, numAttester int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAttester >= len(resource.Attester) {
		return ReferenceInput(frs, "attester["+strconv.Itoa(numAttester)+"].party", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "attester["+strconv.Itoa(numAttester)+"].party", resource.Attester[numAttester].Party, htmlAttrs)
}
func (resource *Composition) T_RelatesToCode(numRelatesTo int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDocument_relationship_type

	if resource == nil || numRelatesTo >= len(resource.RelatesTo) {
		return CodeSelect("relatesTo["+strconv.Itoa(numRelatesTo)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("relatesTo["+strconv.Itoa(numRelatesTo)+"].code", &resource.RelatesTo[numRelatesTo].Code, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_RelatesToTargetIdentifier(numRelatesTo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.RelatesTo) {
		return IdentifierInput("relatesTo["+strconv.Itoa(numRelatesTo)+"].targetIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("relatesTo["+strconv.Itoa(numRelatesTo)+"].targetIdentifier", &resource.RelatesTo[numRelatesTo].TargetIdentifier, htmlAttrs)
}
func (resource *Composition) T_RelatesToTargetReference(frs []FhirResource, numRelatesTo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatesTo >= len(resource.RelatesTo) {
		return ReferenceInput(frs, "relatesTo["+strconv.Itoa(numRelatesTo)+"].targetReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "relatesTo["+strconv.Itoa(numRelatesTo)+"].targetReference", &resource.RelatesTo[numRelatesTo].TargetReference, htmlAttrs)
}
func (resource *Composition) T_EventCode(numEvent int, numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) || numCode >= len(resource.Event[numEvent].Code) {
		return CodeableConceptSelect("event["+strconv.Itoa(numEvent)+"].code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("event["+strconv.Itoa(numEvent)+"].code["+strconv.Itoa(numCode)+"]", &resource.Event[numEvent].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_EventPeriod(numEvent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return PeriodInput("event["+strconv.Itoa(numEvent)+"].period", nil, htmlAttrs)
	}
	return PeriodInput("event["+strconv.Itoa(numEvent)+"].period", resource.Event[numEvent].Period, htmlAttrs)
}
func (resource *Composition) T_EventDetail(frs []FhirResource, numEvent int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) || numDetail >= len(resource.Event[numEvent].Detail) {
		return ReferenceInput(frs, "event["+strconv.Itoa(numEvent)+"].detail["+strconv.Itoa(numDetail)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "event["+strconv.Itoa(numEvent)+"].detail["+strconv.Itoa(numDetail)+"]", &resource.Event[numEvent].Detail[numDetail], htmlAttrs)
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
func (resource *Composition) T_SectionAuthor(frs []FhirResource, numSection int, numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) || numAuthor >= len(resource.Section[numSection].Author) {
		return ReferenceInput(frs, "section["+strconv.Itoa(numSection)+"].author["+strconv.Itoa(numAuthor)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "section["+strconv.Itoa(numSection)+"].author["+strconv.Itoa(numAuthor)+"]", &resource.Section[numSection].Author[numAuthor], htmlAttrs)
}
func (resource *Composition) T_SectionFocus(frs []FhirResource, numSection int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return ReferenceInput(frs, "section["+strconv.Itoa(numSection)+"].focus", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "section["+strconv.Itoa(numSection)+"].focus", resource.Section[numSection].Focus, htmlAttrs)
}
func (resource *Composition) T_SectionMode(numSection int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSList_mode

	if resource == nil || numSection >= len(resource.Section) {
		return CodeSelect("section["+strconv.Itoa(numSection)+"].mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("section["+strconv.Itoa(numSection)+"].mode", resource.Section[numSection].Mode, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_SectionOrderedBy(numSection int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].orderedBy", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].orderedBy", resource.Section[numSection].OrderedBy, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_SectionEntry(frs []FhirResource, numSection int, numEntry int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) || numEntry >= len(resource.Section[numSection].Entry) {
		return ReferenceInput(frs, "section["+strconv.Itoa(numSection)+"].entry["+strconv.Itoa(numEntry)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "section["+strconv.Itoa(numSection)+"].entry["+strconv.Itoa(numEntry)+"]", &resource.Section[numSection].Entry[numEntry], htmlAttrs)
}
func (resource *Composition) T_SectionEmptyReason(numSection int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].emptyReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("section["+strconv.Itoa(numSection)+"].emptyReason", resource.Section[numSection].EmptyReason, optionsValueSet, htmlAttrs)
}
