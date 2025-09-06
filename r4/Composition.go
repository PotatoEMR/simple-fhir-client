package r4

//generated with command go run ./bultaoreune -nodownload
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
	Date              string                 `json:"date"`
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
	Time              *string     `json:"time,omitempty"`
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

func (resource *Composition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Composition.Id", nil)
	}
	return StringInput("Composition.Id", resource.Id)
}
func (resource *Composition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Composition.ImplicitRules", nil)
	}
	return StringInput("Composition.ImplicitRules", resource.ImplicitRules)
}
func (resource *Composition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Composition.Language", nil, optionsValueSet)
	}
	return CodeSelect("Composition.Language", resource.Language, optionsValueSet)
}
func (resource *Composition) T_Status() templ.Component {
	optionsValueSet := VSComposition_status

	if resource == nil {
		return CodeSelect("Composition.Status", nil, optionsValueSet)
	}
	return CodeSelect("Composition.Status", &resource.Status, optionsValueSet)
}
func (resource *Composition) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Composition.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Composition.Type", &resource.Type, optionsValueSet)
}
func (resource *Composition) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("Composition.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Composition.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *Composition) T_Date() templ.Component {

	if resource == nil {
		return StringInput("Composition.Date", nil)
	}
	return StringInput("Composition.Date", &resource.Date)
}
func (resource *Composition) T_Title() templ.Component {

	if resource == nil {
		return StringInput("Composition.Title", nil)
	}
	return StringInput("Composition.Title", &resource.Title)
}
func (resource *Composition) T_Confidentiality(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Composition.Confidentiality", nil, optionsValueSet)
	}
	return CodeSelect("Composition.Confidentiality", resource.Confidentiality, optionsValueSet)
}
func (resource *Composition) T_AttesterId(numAttester int) templ.Component {

	if resource == nil || len(resource.Attester) >= numAttester {
		return StringInput("Composition.Attester["+strconv.Itoa(numAttester)+"].Id", nil)
	}
	return StringInput("Composition.Attester["+strconv.Itoa(numAttester)+"].Id", resource.Attester[numAttester].Id)
}
func (resource *Composition) T_AttesterMode(numAttester int) templ.Component {
	optionsValueSet := VSComposition_attestation_mode

	if resource == nil || len(resource.Attester) >= numAttester {
		return CodeSelect("Composition.Attester["+strconv.Itoa(numAttester)+"].Mode", nil, optionsValueSet)
	}
	return CodeSelect("Composition.Attester["+strconv.Itoa(numAttester)+"].Mode", &resource.Attester[numAttester].Mode, optionsValueSet)
}
func (resource *Composition) T_AttesterTime(numAttester int) templ.Component {

	if resource == nil || len(resource.Attester) >= numAttester {
		return StringInput("Composition.Attester["+strconv.Itoa(numAttester)+"].Time", nil)
	}
	return StringInput("Composition.Attester["+strconv.Itoa(numAttester)+"].Time", resource.Attester[numAttester].Time)
}
func (resource *Composition) T_RelatesToId(numRelatesTo int) templ.Component {

	if resource == nil || len(resource.RelatesTo) >= numRelatesTo {
		return StringInput("Composition.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Id", nil)
	}
	return StringInput("Composition.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Id", resource.RelatesTo[numRelatesTo].Id)
}
func (resource *Composition) T_RelatesToCode(numRelatesTo int) templ.Component {
	optionsValueSet := VSDocument_relationship_type

	if resource == nil || len(resource.RelatesTo) >= numRelatesTo {
		return CodeSelect("Composition.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Code", nil, optionsValueSet)
	}
	return CodeSelect("Composition.RelatesTo["+strconv.Itoa(numRelatesTo)+"].Code", &resource.RelatesTo[numRelatesTo].Code, optionsValueSet)
}
func (resource *Composition) T_EventId(numEvent int) templ.Component {

	if resource == nil || len(resource.Event) >= numEvent {
		return StringInput("Composition.Event["+strconv.Itoa(numEvent)+"].Id", nil)
	}
	return StringInput("Composition.Event["+strconv.Itoa(numEvent)+"].Id", resource.Event[numEvent].Id)
}
func (resource *Composition) T_EventCode(numEvent int, numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Event) >= numEvent || len(resource.Event[numEvent].Code) >= numCode {
		return CodeableConceptSelect("Composition.Event["+strconv.Itoa(numEvent)+"].Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Composition.Event["+strconv.Itoa(numEvent)+"].Code["+strconv.Itoa(numCode)+"]", &resource.Event[numEvent].Code[numCode], optionsValueSet)
}
func (resource *Composition) T_SectionId(numSection int) templ.Component {

	if resource == nil || len(resource.Section) >= numSection {
		return StringInput("Composition.Section["+strconv.Itoa(numSection)+"].Id", nil)
	}
	return StringInput("Composition.Section["+strconv.Itoa(numSection)+"].Id", resource.Section[numSection].Id)
}
func (resource *Composition) T_SectionTitle(numSection int) templ.Component {

	if resource == nil || len(resource.Section) >= numSection {
		return StringInput("Composition.Section["+strconv.Itoa(numSection)+"].Title", nil)
	}
	return StringInput("Composition.Section["+strconv.Itoa(numSection)+"].Title", resource.Section[numSection].Title)
}
func (resource *Composition) T_SectionCode(numSection int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Section) >= numSection {
		return CodeableConceptSelect("Composition.Section["+strconv.Itoa(numSection)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Composition.Section["+strconv.Itoa(numSection)+"].Code", resource.Section[numSection].Code, optionsValueSet)
}
func (resource *Composition) T_SectionMode(numSection int) templ.Component {
	optionsValueSet := VSList_mode

	if resource == nil || len(resource.Section) >= numSection {
		return CodeSelect("Composition.Section["+strconv.Itoa(numSection)+"].Mode", nil, optionsValueSet)
	}
	return CodeSelect("Composition.Section["+strconv.Itoa(numSection)+"].Mode", resource.Section[numSection].Mode, optionsValueSet)
}
func (resource *Composition) T_SectionOrderedBy(numSection int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Section) >= numSection {
		return CodeableConceptSelect("Composition.Section["+strconv.Itoa(numSection)+"].OrderedBy", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Composition.Section["+strconv.Itoa(numSection)+"].OrderedBy", resource.Section[numSection].OrderedBy, optionsValueSet)
}
func (resource *Composition) T_SectionEmptyReason(numSection int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Section) >= numSection {
		return CodeableConceptSelect("Composition.Section["+strconv.Itoa(numSection)+"].EmptyReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Composition.Section["+strconv.Itoa(numSection)+"].EmptyReason", resource.Section[numSection].EmptyReason, optionsValueSet)
}
