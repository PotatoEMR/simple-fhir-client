package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *Composition) CompositionLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Composition) CompositionStatus() templ.Component {
	optionsValueSet := VSComposition_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *Composition) CompositionConfidentiality(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Confidentiality
	}
	return CodeSelect("confidentiality", currentVal, optionsValueSet)
}
func (resource *Composition) CompositionAttesterMode(numAttester int) templ.Component {
	optionsValueSet := VSComposition_attestation_mode
	currentVal := ""
	if resource != nil && len(resource.Attester) >= numAttester {
		currentVal = resource.Attester[numAttester].Mode
	}
	return CodeSelect("mode", currentVal, optionsValueSet)
}
func (resource *Composition) CompositionRelatesToCode(numRelatesTo int) templ.Component {
	optionsValueSet := VSDocument_relationship_type
	currentVal := ""
	if resource != nil && len(resource.RelatesTo) >= numRelatesTo {
		currentVal = resource.RelatesTo[numRelatesTo].Code
	}
	return CodeSelect("code", currentVal, optionsValueSet)
}
func (resource *Composition) CompositionSectionMode(numSection int) templ.Component {
	optionsValueSet := VSList_mode
	currentVal := ""
	if resource != nil && len(resource.Section) >= numSection {
		currentVal = *resource.Section[numSection].Mode
	}
	return CodeSelect("mode", currentVal, optionsValueSet)
}
