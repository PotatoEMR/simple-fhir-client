package r4b

//generated with command go run ./bultaoreune
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
	Date                      *FhirDateTime                    `json:"date,omitempty"`
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
func (resource *EvidenceVariable) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *EvidenceVariable) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *EvidenceVariable) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *EvidenceVariable) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *EvidenceVariable) T_ShortTitle(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("shortTitle", nil, htmlAttrs)
	}
	return StringInput("shortTitle", resource.ShortTitle, htmlAttrs)
}
func (resource *EvidenceVariable) T_Subtitle(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("subtitle", nil, htmlAttrs)
	}
	return StringInput("subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *EvidenceVariable) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *EvidenceVariable) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *EvidenceVariable) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *EvidenceVariable) T_UseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.UseContext) {
		return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("useContext["+strconv.Itoa(numUseContext)+"]", &resource.UseContext[numUseContext], htmlAttrs)
}
func (resource *EvidenceVariable) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *EvidenceVariable) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *EvidenceVariable) T_Author(numAuthor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAuthor >= len(resource.Author) {
		return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("author["+strconv.Itoa(numAuthor)+"]", &resource.Author[numAuthor], htmlAttrs)
}
func (resource *EvidenceVariable) T_Editor(numEditor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEditor >= len(resource.Editor) {
		return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("editor["+strconv.Itoa(numEditor)+"]", &resource.Editor[numEditor], htmlAttrs)
}
func (resource *EvidenceVariable) T_Reviewer(numReviewer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numReviewer >= len(resource.Reviewer) {
		return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("reviewer["+strconv.Itoa(numReviewer)+"]", &resource.Reviewer[numReviewer], htmlAttrs)
}
func (resource *EvidenceVariable) T_Endorser(numEndorser int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEndorser >= len(resource.Endorser) {
		return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", nil, htmlAttrs)
	}
	return ContactDetailInput("endorser["+strconv.Itoa(numEndorser)+"]", &resource.Endorser[numEndorser], htmlAttrs)
}
func (resource *EvidenceVariable) T_RelatedArtifact(numRelatedArtifact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedArtifact >= len(resource.RelatedArtifact) {
		return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", &resource.RelatedArtifact[numRelatedArtifact], htmlAttrs)
}
func (resource *EvidenceVariable) T_Actual(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("actual", nil, htmlAttrs)
	}
	return BoolInput("actual", resource.Actual, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicCombination(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCharacteristic_combination

	if resource == nil {
		return CodeSelect("characteristicCombination", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("characteristicCombination", resource.CharacteristicCombination, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_Handling(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSVariable_handling

	if resource == nil {
		return CodeSelect("handling", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("handling", resource.Handling, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDescription(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].description", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].description", resource.Characteristic[numCharacteristic].Description, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionReference(frs []FhirResource, numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return ReferenceInput(frs, "characteristic["+strconv.Itoa(numCharacteristic)+"].definitionReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "characteristic["+strconv.Itoa(numCharacteristic)+"].definitionReference", &resource.Characteristic[numCharacteristic].DefinitionReference, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionCanonical(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionCanonical", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionCanonical", &resource.Characteristic[numCharacteristic].DefinitionCanonical, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionCodeableConcept", &resource.Characteristic[numCharacteristic].DefinitionCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionExpression(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return ExpressionInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionExpression", nil, htmlAttrs)
	}
	return ExpressionInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionExpression", &resource.Characteristic[numCharacteristic].DefinitionExpression, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicMethod(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].method", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].method", resource.Characteristic[numCharacteristic].Method, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDevice(frs []FhirResource, numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return ReferenceInput(frs, "characteristic["+strconv.Itoa(numCharacteristic)+"].device", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "characteristic["+strconv.Itoa(numCharacteristic)+"].device", resource.Characteristic[numCharacteristic].Device, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicExclude(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].exclude", nil, htmlAttrs)
	}
	return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].exclude", resource.Characteristic[numCharacteristic].Exclude, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicGroupMeasure(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSGroup_measure

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].groupMeasure", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].groupMeasure", resource.Characteristic[numCharacteristic].GroupMeasure, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromStartDescription(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromStart.description", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromStart.description", resource.Characteristic[numCharacteristic].TimeFromStart.Description, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromStartQuantity(numCharacteristic int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return QuantityInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromStart.quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromStart.quantity", resource.Characteristic[numCharacteristic].TimeFromStart.Quantity, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromStartRange(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return RangeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromStart.range", nil, htmlAttrs)
	}
	return RangeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromStart.range", resource.Characteristic[numCharacteristic].TimeFromStart.Range, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromStartNote(numCharacteristic int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numNote >= len(resource.Characteristic[numCharacteristic].TimeFromStart.Note) {
		return AnnotationTextArea("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromStart.note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromStart.note["+strconv.Itoa(numNote)+"]", &resource.Characteristic[numCharacteristic].TimeFromStart.Note[numNote], htmlAttrs)
}
func (resource *EvidenceVariable) T_CategoryName(numCategory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return StringInput("category["+strconv.Itoa(numCategory)+"].name", nil, htmlAttrs)
	}
	return StringInput("category["+strconv.Itoa(numCategory)+"].name", resource.Category[numCategory].Name, htmlAttrs)
}
func (resource *EvidenceVariable) T_CategoryValueCodeableConcept(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"].valueCodeableConcept", resource.Category[numCategory].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CategoryValueQuantity(numCategory int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return QuantityInput("category["+strconv.Itoa(numCategory)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("category["+strconv.Itoa(numCategory)+"].valueQuantity", resource.Category[numCategory].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CategoryValueRange(numCategory int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return RangeInput("category["+strconv.Itoa(numCategory)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("category["+strconv.Itoa(numCategory)+"].valueRange", resource.Category[numCategory].ValueRange, htmlAttrs)
}
