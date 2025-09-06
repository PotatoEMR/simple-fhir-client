package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	Date                      *string                          `json:"date,omitempty"`
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

func (resource *EvidenceVariable) T_Id() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Id", nil)
	}
	return StringInput("EvidenceVariable.Id", resource.Id)
}
func (resource *EvidenceVariable) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.ImplicitRules", nil)
	}
	return StringInput("EvidenceVariable.ImplicitRules", resource.ImplicitRules)
}
func (resource *EvidenceVariable) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("EvidenceVariable.Language", nil, optionsValueSet)
	}
	return CodeSelect("EvidenceVariable.Language", resource.Language, optionsValueSet)
}
func (resource *EvidenceVariable) T_Url() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Url", nil)
	}
	return StringInput("EvidenceVariable.Url", resource.Url)
}
func (resource *EvidenceVariable) T_Version() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Version", nil)
	}
	return StringInput("EvidenceVariable.Version", resource.Version)
}
func (resource *EvidenceVariable) T_Name() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Name", nil)
	}
	return StringInput("EvidenceVariable.Name", resource.Name)
}
func (resource *EvidenceVariable) T_Title() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Title", nil)
	}
	return StringInput("EvidenceVariable.Title", resource.Title)
}
func (resource *EvidenceVariable) T_ShortTitle() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.ShortTitle", nil)
	}
	return StringInput("EvidenceVariable.ShortTitle", resource.ShortTitle)
}
func (resource *EvidenceVariable) T_Subtitle() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Subtitle", nil)
	}
	return StringInput("EvidenceVariable.Subtitle", resource.Subtitle)
}
func (resource *EvidenceVariable) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("EvidenceVariable.Status", nil, optionsValueSet)
	}
	return CodeSelect("EvidenceVariable.Status", &resource.Status, optionsValueSet)
}
func (resource *EvidenceVariable) T_Date() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Date", nil)
	}
	return StringInput("EvidenceVariable.Date", resource.Date)
}
func (resource *EvidenceVariable) T_Description() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Description", nil)
	}
	return StringInput("EvidenceVariable.Description", resource.Description)
}
func (resource *EvidenceVariable) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("EvidenceVariable.Publisher", nil)
	}
	return StringInput("EvidenceVariable.Publisher", resource.Publisher)
}
func (resource *EvidenceVariable) T_Actual() templ.Component {

	if resource == nil {
		return BoolInput("EvidenceVariable.Actual", nil)
	}
	return BoolInput("EvidenceVariable.Actual", resource.Actual)
}
func (resource *EvidenceVariable) T_CharacteristicCombination() templ.Component {
	optionsValueSet := VSCharacteristic_combination

	if resource == nil {
		return CodeSelect("EvidenceVariable.CharacteristicCombination", nil, optionsValueSet)
	}
	return CodeSelect("EvidenceVariable.CharacteristicCombination", resource.CharacteristicCombination, optionsValueSet)
}
func (resource *EvidenceVariable) T_Handling() templ.Component {
	optionsValueSet := VSVariable_handling

	if resource == nil {
		return CodeSelect("EvidenceVariable.Handling", nil, optionsValueSet)
	}
	return CodeSelect("EvidenceVariable.Handling", resource.Handling, optionsValueSet)
}
func (resource *EvidenceVariable) T_CharacteristicId(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Id", nil)
	}
	return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Id", resource.Characteristic[numCharacteristic].Id)
}
func (resource *EvidenceVariable) T_CharacteristicDescription(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Description", nil)
	}
	return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Description", resource.Characteristic[numCharacteristic].Description)
}
func (resource *EvidenceVariable) T_CharacteristicMethod(numCharacteristic int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return CodeableConceptSelect("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Method", nil, optionsValueSet)
	}
	return CodeableConceptSelect("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Method", resource.Characteristic[numCharacteristic].Method, optionsValueSet)
}
func (resource *EvidenceVariable) T_CharacteristicExclude(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return BoolInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Exclude", nil)
	}
	return BoolInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].Exclude", resource.Characteristic[numCharacteristic].Exclude)
}
func (resource *EvidenceVariable) T_CharacteristicGroupMeasure(numCharacteristic int) templ.Component {
	optionsValueSet := VSGroup_measure

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return CodeSelect("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].GroupMeasure", nil, optionsValueSet)
	}
	return CodeSelect("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].GroupMeasure", resource.Characteristic[numCharacteristic].GroupMeasure, optionsValueSet)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromStartId(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].TimeFromStart.Id", nil)
	}
	return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].TimeFromStart.Id", resource.Characteristic[numCharacteristic].TimeFromStart.Id)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromStartDescription(numCharacteristic int) templ.Component {

	if resource == nil || len(resource.Characteristic) >= numCharacteristic {
		return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].TimeFromStart.Description", nil)
	}
	return StringInput("EvidenceVariable.Characteristic["+strconv.Itoa(numCharacteristic)+"].TimeFromStart.Description", resource.Characteristic[numCharacteristic].TimeFromStart.Description)
}
func (resource *EvidenceVariable) T_CategoryId(numCategory int) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return StringInput("EvidenceVariable.Category["+strconv.Itoa(numCategory)+"].Id", nil)
	}
	return StringInput("EvidenceVariable.Category["+strconv.Itoa(numCategory)+"].Id", resource.Category[numCategory].Id)
}
func (resource *EvidenceVariable) T_CategoryName(numCategory int) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return StringInput("EvidenceVariable.Category["+strconv.Itoa(numCategory)+"].Name", nil)
	}
	return StringInput("EvidenceVariable.Category["+strconv.Itoa(numCategory)+"].Name", resource.Category[numCategory].Name)
}
