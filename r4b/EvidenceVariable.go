package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/EvidenceVariable
type EvidenceVariable struct {
	Id                        *string                          `json:"id,omitempty"`
	Meta                      *Meta                            `json:"meta,omitempty"`
	ImplicitRules             *string                          `json:"implicitRules,omitempty"`
	Language                  *string                          `json:"language,omitempty"`
	Text                      *Narrative                       `json:"text,omitempty"`
	Contained                 []Resource                       `json:"contained,omitempty"`
	Extension                 []Extension                      `json:"extension,omitempty"`
	ModifierExtension         []Extension                      `json:"modifierExtension,omitempty"`
	Url                       *string                          `json:"url,omitempty"`
	Identifier                []Identifier                     `json:"identifier,omitempty"`
	Version                   *string                          `json:"version,omitempty"`
	Name                      *string                          `json:"name,omitempty"`
	Title                     *string                          `json:"title,omitempty"`
	ShortTitle                *string                          `json:"shortTitle,omitempty"`
	Subtitle                  *string                          `json:"subtitle,omitempty"`
	Status                    string                           `json:"status"`
	Date                      *time.Time                       `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Description               *string                          `json:"description,omitempty"`
	Note                      []Annotation                     `json:"note,omitempty"`
	UseContext                []UsageContext                   `json:"useContext,omitempty"`
	Publisher                 *string                          `json:"publisher,omitempty"`
	Contact                   []ContactDetail                  `json:"contact,omitempty"`
	Author                    []ContactDetail                  `json:"author,omitempty"`
	Editor                    []ContactDetail                  `json:"editor,omitempty"`
	Reviewer                  []ContactDetail                  `json:"reviewer,omitempty"`
	Endorser                  []ContactDetail                  `json:"endorser,omitempty"`
	RelatedArtifact           []RelatedArtifact                `json:"relatedArtifact,omitempty"`
	Actual                    *bool                            `json:"actual,omitempty"`
	CharacteristicCombination *string                          `json:"characteristicCombination,omitempty"`
	Characteristic            []EvidenceVariableCharacteristic `json:"characteristic,omitempty"`
	Handling                  *string                          `json:"handling,omitempty"`
	Category                  []EvidenceVariableCategory       `json:"category,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/EvidenceVariable
type EvidenceVariableCharacteristic struct {
	Id                        *string                                      `json:"id,omitempty"`
	Extension                 []Extension                                  `json:"extension,omitempty"`
	ModifierExtension         []Extension                                  `json:"modifierExtension,omitempty"`
	Description               *string                                      `json:"description,omitempty"`
	DefinitionReference       Reference                                    `json:"definitionReference"`
	DefinitionCanonical       string                                       `json:"definitionCanonical"`
	DefinitionCodeableConcept CodeableConcept                              `json:"definitionCodeableConcept"`
	DefinitionExpression      Expression                                   `json:"definitionExpression"`
	Method                    *CodeableConcept                             `json:"method,omitempty"`
	Device                    *Reference                                   `json:"device,omitempty"`
	Exclude                   *bool                                        `json:"exclude,omitempty"`
	TimeFromStart             *EvidenceVariableCharacteristicTimeFromStart `json:"timeFromStart,omitempty"`
	GroupMeasure              *string                                      `json:"groupMeasure,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/EvidenceVariable
type EvidenceVariableCharacteristicTimeFromStart struct {
	Id                *string      `json:"id,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Description       *string      `json:"description,omitempty"`
	Quantity          *Quantity    `json:"quantity,omitempty"`
	Range             *Range       `json:"range,omitempty"`
	Note              []Annotation `json:"note,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/EvidenceVariable
type EvidenceVariableCategory struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Name                 *string          `json:"name,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range           `json:"valueRange,omitempty"`
}

type OtherEvidenceVariable EvidenceVariable

// on convert struct to json, automatically add resourceType=EvidenceVariable
func (r EvidenceVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEvidenceVariable
		ResourceType string `json:"resourceType"`
	}{
		OtherEvidenceVariable: OtherEvidenceVariable(r),
		ResourceType:          "EvidenceVariable",
	})
}
func (r EvidenceVariable) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "EvidenceVariable/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "EvidenceVariable"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *EvidenceVariable) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EvidenceVariable.Url", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Url", resource.Url, htmlAttrs)
}
func (resource *EvidenceVariable) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EvidenceVariable.Version", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Version", resource.Version, htmlAttrs)
}
func (resource *EvidenceVariable) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EvidenceVariable.Name", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Name", resource.Name, htmlAttrs)
}
func (resource *EvidenceVariable) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EvidenceVariable.Title", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Title", resource.Title, htmlAttrs)
}
func (resource *EvidenceVariable) T_ShortTitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EvidenceVariable.ShortTitle", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.ShortTitle", resource.ShortTitle, htmlAttrs)
}
func (resource *EvidenceVariable) T_Subtitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EvidenceVariable.Subtitle", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *EvidenceVariable) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("EvidenceVariable.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("EvidenceVariable.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("EvidenceVariable.Date", nil, htmlAttrs)
	}
	return DateTimeInput("EvidenceVariable.Date", resource.Date, htmlAttrs)
}
func (resource *EvidenceVariable) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EvidenceVariable.Description", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Description", resource.Description, htmlAttrs)
}
func (resource *EvidenceVariable) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("EvidenceVariable.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("EvidenceVariable.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *EvidenceVariable) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("EvidenceVariable.Publisher", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *EvidenceVariable) T_Actual(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("EvidenceVariable.Actual", nil, htmlAttrs)
	}
	return BoolInput("EvidenceVariable.Actual", resource.Actual, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicCombination(htmlAttrs string) templ.Component {
	optionsValueSet := VSCharacteristic_combination

	if resource == nil {
		return CodeSelect("EvidenceVariable.CharacteristicCombination", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("EvidenceVariable.CharacteristicCombination", resource.CharacteristicCombination, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_Handling(htmlAttrs string) templ.Component {
	optionsValueSet := VSVariable_handling

	if resource == nil {
		return CodeSelect("EvidenceVariable.Handling", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("EvidenceVariable.Handling", resource.Handling, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDescription(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Description", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Description", resource.Characteristic[numCharacteristic].Description, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionCanonical(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].DefinitionCanonical", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].DefinitionCanonical", &resource.Characteristic[numCharacteristic].DefinitionCanonical, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].DefinitionCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].DefinitionCodeableConcept", &resource.Characteristic[numCharacteristic].DefinitionCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicMethod(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Method", resource.Characteristic[numCharacteristic].Method, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicExclude(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Exclude", nil, htmlAttrs)
	}
	return BoolInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Exclude", resource.Characteristic[numCharacteristic].Exclude, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicGroupMeasure(numCharacteristic int, htmlAttrs string) templ.Component {
	optionsValueSet := VSGroup_measure

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeSelect("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].GroupMeasure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].GroupMeasure", resource.Characteristic[numCharacteristic].GroupMeasure, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromStartDescription(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].TimeFromStart.Description", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].TimeFromStart.Description", resource.Characteristic[numCharacteristic].TimeFromStart.Description, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromStartNote(numCharacteristic int, numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numNote >= len(resource.Characteristic[numCharacteristic].TimeFromStart.Note) {
		return AnnotationTextArea("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].TimeFromStart.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].TimeFromStart.Note["+strconv.Itoa(numNote)+"]", &resource.Characteristic[numCharacteristic].TimeFromStart.Note[numNote], htmlAttrs)
}
func (resource *EvidenceVariable) T_CategoryName(numCategory int, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return StringInput("EvidenceVariable.Category["+strconv.Itoa(numCategory)+"].Name", nil, htmlAttrs)
	}
	return StringInput("EvidenceVariable.Category["+strconv.Itoa(numCategory)+"].Name", resource.Category[numCategory].Name, htmlAttrs)
}
func (resource *EvidenceVariable) T_CategoryValueCodeableConcept(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("EvidenceVariable.Category["+strconv.Itoa(numCategory)+"].ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("EvidenceVariable.Category["+strconv.Itoa(numCategory)+"].ValueCodeableConcept", resource.Category[numCategory].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
