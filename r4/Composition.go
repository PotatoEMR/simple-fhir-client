package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Date              time.Time              `json:"date,format:'2006-01-02T15:04:05Z07:00'"`
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
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Mode              string      `json:"mode"`
	Time              *time.Time  `json:"time,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Party             *Reference  `json:"party,omitempty"`
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
func (resource *Composition) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSComposition_status

	if resource == nil {
		return CodeSelect("Composition.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Composition.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Composition.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Composition.Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Composition.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Composition.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Composition.Date", nil, htmlAttrs)
	}
	return DateTimeInput("Composition.Date", &resource.Date, htmlAttrs)
}
func (resource *Composition) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Composition.Title", nil, htmlAttrs)
	}
	return StringInput("Composition.Title", &resource.Title, htmlAttrs)
}
func (resource *Composition) T_Confidentiality(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeSelect("Composition.Confidentiality", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Composition.Confidentiality", resource.Confidentiality, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_AttesterMode(numAttester int, htmlAttrs string) templ.Component {
	optionsValueSet := VSComposition_attestation_mode

	if resource == nil || numAttester >= len(resource.Attester) {
		return CodeSelect("Composition.Attester["+strconv.Itoa(numAttester)+"].Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Composition.Attester["+strconv.Itoa(numAttester)+"].Mode", &resource.Attester[numAttester].Mode, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_AttesterTime(numAttester int, htmlAttrs string) templ.Component {
	if resource == nil || numAttester >= len(resource.Attester) {
		return DateTimeInput("Composition.Attester["+strconv.Itoa(numAttester)+"].Time", nil, htmlAttrs)
	}
	return DateTimeInput("Composition.Attester["+strconv.Itoa(numAttester)+"].Time", resource.Attester[numAttester].Time, htmlAttrs)
}
func (resource *Composition) T_RelatesToCode(numRelatesTo int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDocument_relationship_type

	if resource == nil || numRelatesTo >= len(resource.RelatesTo) {
		return CodeSelect("Composition.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Composition.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Code", &resource.RelatesTo[numRelatesTo].Code, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_EventCode(numEvent int, numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) || numCode >= len(resource.Event[numEvent].Code) {
		return CodeableConceptSelect("Composition.Event["+strconv.Itoa(numEvent)+"].Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Composition.Event["+strconv.Itoa(numEvent)+"].Code["+strconv.Itoa(numCode)+"]", &resource.Event[numEvent].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_SectionTitle(numSection int, htmlAttrs string) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return StringInput("Composition.Section["+strconv.Itoa(numSection)+"].Title", nil, htmlAttrs)
	}
	return StringInput("Composition.Section["+strconv.Itoa(numSection)+"].Title", resource.Section[numSection].Title, htmlAttrs)
}
func (resource *Composition) T_SectionCode(numSection int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("Composition.Section["+strconv.Itoa(numSection)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Composition.Section["+strconv.Itoa(numSection)+"].Code", resource.Section[numSection].Code, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_SectionMode(numSection int, htmlAttrs string) templ.Component {
	optionsValueSet := VSList_mode

	if resource == nil || numSection >= len(resource.Section) {
		return CodeSelect("Composition.Section["+strconv.Itoa(numSection)+"].Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Composition.Section["+strconv.Itoa(numSection)+"].Mode", resource.Section[numSection].Mode, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_SectionOrderedBy(numSection int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("Composition.Section["+strconv.Itoa(numSection)+"].OrderedBy", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Composition.Section["+strconv.Itoa(numSection)+"].OrderedBy", resource.Section[numSection].OrderedBy, optionsValueSet, htmlAttrs)
}
func (resource *Composition) T_SectionEmptyReason(numSection int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSection >= len(resource.Section) {
		return CodeableConceptSelect("Composition.Section["+strconv.Itoa(numSection)+"].EmptyReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Composition.Section["+strconv.Itoa(numSection)+"].EmptyReason", resource.Section[numSection].EmptyReason, optionsValueSet, htmlAttrs)
}
