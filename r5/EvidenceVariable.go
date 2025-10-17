package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceVariable
type EvidenceVariable struct {
	Id                     *string                          `json:"id,omitempty"`
	Meta                   *Meta                            `json:"meta,omitempty"`
	ImplicitRules          *string                          `json:"implicitRules,omitempty"`
	Language               *string                          `json:"language,omitempty"`
	Text                   *Narrative                       `json:"text,omitempty"`
	Contained              []Resource                       `json:"contained,omitempty"`
	Extension              []Extension                      `json:"extension,omitempty"`
	ModifierExtension      []Extension                      `json:"modifierExtension,omitempty"`
	Url                    *string                          `json:"url,omitempty"`
	Identifier             []Identifier                     `json:"identifier,omitempty"`
	Version                *string                          `json:"version,omitempty"`
	VersionAlgorithmString *string                          `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                          `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                          `json:"name,omitempty"`
	Title                  *string                          `json:"title,omitempty"`
	ShortTitle             *string                          `json:"shortTitle,omitempty"`
	Status                 string                           `json:"status"`
	Experimental           *bool                            `json:"experimental,omitempty"`
	Date                   *FhirDateTime                    `json:"date,omitempty"`
	Publisher              *string                          `json:"publisher,omitempty"`
	Contact                []ContactDetail                  `json:"contact,omitempty"`
	Description            *string                          `json:"description,omitempty"`
	Note                   []Annotation                     `json:"note,omitempty"`
	UseContext             []UsageContext                   `json:"useContext,omitempty"`
	Purpose                *string                          `json:"purpose,omitempty"`
	Copyright              *string                          `json:"copyright,omitempty"`
	CopyrightLabel         *string                          `json:"copyrightLabel,omitempty"`
	ApprovalDate           *FhirDate                        `json:"approvalDate,omitempty"`
	LastReviewDate         *FhirDate                        `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                          `json:"effectivePeriod,omitempty"`
	Author                 []ContactDetail                  `json:"author,omitempty"`
	Editor                 []ContactDetail                  `json:"editor,omitempty"`
	Reviewer               []ContactDetail                  `json:"reviewer,omitempty"`
	Endorser               []ContactDetail                  `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact                `json:"relatedArtifact,omitempty"`
	Actual                 *bool                            `json:"actual,omitempty"`
	Characteristic         []EvidenceVariableCharacteristic `json:"characteristic,omitempty"`
	Handling               *string                          `json:"handling,omitempty"`
	Category               []EvidenceVariableCategory       `json:"category,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceVariable
type EvidenceVariableCharacteristic struct {
	Id                        *string                                                 `json:"id,omitempty"`
	Extension                 []Extension                                             `json:"extension,omitempty"`
	ModifierExtension         []Extension                                             `json:"modifierExtension,omitempty"`
	LinkId                    *string                                                 `json:"linkId,omitempty"`
	Description               *string                                                 `json:"description,omitempty"`
	Note                      []Annotation                                            `json:"note,omitempty"`
	Exclude                   *bool                                                   `json:"exclude,omitempty"`
	DefinitionReference       *Reference                                              `json:"definitionReference,omitempty"`
	DefinitionCanonical       *string                                                 `json:"definitionCanonical,omitempty"`
	DefinitionCodeableConcept *CodeableConcept                                        `json:"definitionCodeableConcept,omitempty"`
	DefinitionExpression      *Expression                                             `json:"definitionExpression,omitempty"`
	DefinitionId              *string                                                 `json:"definitionId,omitempty"`
	DefinitionByTypeAndValue  *EvidenceVariableCharacteristicDefinitionByTypeAndValue `json:"definitionByTypeAndValue,omitempty"`
	DefinitionByCombination   *EvidenceVariableCharacteristicDefinitionByCombination  `json:"definitionByCombination,omitempty"`
	InstancesQuantity         *Quantity                                               `json:"instancesQuantity,omitempty"`
	InstancesRange            *Range                                                  `json:"instancesRange,omitempty"`
	DurationQuantity          *Quantity                                               `json:"durationQuantity,omitempty"`
	DurationRange             *Range                                                  `json:"durationRange,omitempty"`
	TimeFromEvent             []EvidenceVariableCharacteristicTimeFromEvent           `json:"timeFromEvent,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceVariable
type EvidenceVariableCharacteristicDefinitionByTypeAndValue struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept   `json:"type"`
	Method               []CodeableConcept `json:"method,omitempty"`
	Device               *Reference        `json:"device,omitempty"`
	ValueCodeableConcept CodeableConcept   `json:"valueCodeableConcept"`
	ValueBoolean         bool              `json:"valueBoolean"`
	ValueQuantity        Quantity          `json:"valueQuantity"`
	ValueRange           Range             `json:"valueRange"`
	ValueReference       Reference         `json:"valueReference"`
	ValueId              string            `json:"valueId"`
	Offset               *CodeableConcept  `json:"offset,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceVariable
type EvidenceVariableCharacteristicDefinitionByCombination struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Code              string      `json:"code"`
	Threshold         *int        `json:"threshold,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceVariable
type EvidenceVariableCharacteristicTimeFromEvent struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Description          *string          `json:"description,omitempty"`
	Note                 []Annotation     `json:"note,omitempty"`
	EventCodeableConcept *CodeableConcept `json:"eventCodeableConcept,omitempty"`
	EventReference       *Reference       `json:"eventReference,omitempty"`
	EventDateTime        *FhirDateTime    `json:"eventDateTime,omitempty"`
	EventId              *string          `json:"eventId,omitempty"`
	Quantity             *Quantity        `json:"quantity,omitempty"`
	Range                *Range           `json:"range,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/EvidenceVariable
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

// struct -> json, automatically add resourceType=Patient
func (r EvidenceVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEvidenceVariable
		ResourceType string `json:"resourceType"`
	}{
		OtherEvidenceVariable: OtherEvidenceVariable(r),
		ResourceType:          "EvidenceVariable",
	})
}

// json -> struct, first reject if resourceType != EvidenceVariable
func (r *EvidenceVariable) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "EvidenceVariable" {
		return errors.New("resourceType not EvidenceVariable")
	}
	return json.Unmarshal(data, (*OtherEvidenceVariable)(r))
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
func (resource *EvidenceVariable) T_VersionAlgorithmString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("versionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("versionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *EvidenceVariable) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodingSelect("versionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("versionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
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
func (resource *EvidenceVariable) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *EvidenceVariable) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("date", resource.Date, htmlAttrs)
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
func (resource *EvidenceVariable) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *EvidenceVariable) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *EvidenceVariable) T_CopyrightLabel(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyrightLabel", nil, htmlAttrs)
	}
	return StringInput("copyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *EvidenceVariable) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("approvalDate", nil, htmlAttrs)
	}
	return FhirDateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *EvidenceVariable) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("lastReviewDate", nil, htmlAttrs)
	}
	return FhirDateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *EvidenceVariable) T_EffectivePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("effectivePeriod", resource.EffectivePeriod, htmlAttrs)
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
func (resource *EvidenceVariable) T_Handling(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSVariable_handling

	if resource == nil {
		return CodeSelect("handling", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("handling", resource.Handling, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicLinkId(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].linkId", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].linkId", resource.Characteristic[numCharacteristic].LinkId, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDescription(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].description", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].description", resource.Characteristic[numCharacteristic].Description, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicNote(numCharacteristic int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numNote >= len(resource.Characteristic[numCharacteristic].Note) {
		return AnnotationTextArea("characteristic["+strconv.Itoa(numCharacteristic)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("characteristic["+strconv.Itoa(numCharacteristic)+"].note["+strconv.Itoa(numNote)+"]", &resource.Characteristic[numCharacteristic].Note[numNote], htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicExclude(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].exclude", nil, htmlAttrs)
	}
	return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].exclude", resource.Characteristic[numCharacteristic].Exclude, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionReference(frs []FhirResource, numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return ReferenceInput(frs, "characteristic["+strconv.Itoa(numCharacteristic)+"].definitionReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "characteristic["+strconv.Itoa(numCharacteristic)+"].definitionReference", resource.Characteristic[numCharacteristic].DefinitionReference, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionCanonical(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionCanonical", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionCanonical", resource.Characteristic[numCharacteristic].DefinitionCanonical, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionCodeableConcept", resource.Characteristic[numCharacteristic].DefinitionCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionExpression(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return ExpressionInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionExpression", nil, htmlAttrs)
	}
	return ExpressionInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionExpression", resource.Characteristic[numCharacteristic].DefinitionExpression, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionId(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionId", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionId", resource.Characteristic[numCharacteristic].DefinitionId, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicInstancesQuantity(numCharacteristic int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return QuantityInput("characteristic["+strconv.Itoa(numCharacteristic)+"].instancesQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("characteristic["+strconv.Itoa(numCharacteristic)+"].instancesQuantity", resource.Characteristic[numCharacteristic].InstancesQuantity, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicInstancesRange(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return RangeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].instancesRange", nil, htmlAttrs)
	}
	return RangeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].instancesRange", resource.Characteristic[numCharacteristic].InstancesRange, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDurationQuantity(numCharacteristic int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return QuantityInput("characteristic["+strconv.Itoa(numCharacteristic)+"].durationQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("characteristic["+strconv.Itoa(numCharacteristic)+"].durationQuantity", resource.Characteristic[numCharacteristic].DurationQuantity, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDurationRange(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return RangeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].durationRange", nil, htmlAttrs)
	}
	return RangeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].durationRange", resource.Characteristic[numCharacteristic].DurationRange, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByTypeAndValueType(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.type", &resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.Type, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByTypeAndValueMethod(numCharacteristic int, numMethod int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numMethod >= len(resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.Method) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.method["+strconv.Itoa(numMethod)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.method["+strconv.Itoa(numMethod)+"]", &resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.Method[numMethod], optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByTypeAndValueDevice(frs []FhirResource, numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return ReferenceInput(frs, "characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.device", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.device", resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.Device, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByTypeAndValueValueCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.valueCodeableConcept", &resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByTypeAndValueValueBoolean(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.valueBoolean", &resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.ValueBoolean, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByTypeAndValueValueQuantity(numCharacteristic int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return QuantityInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.valueQuantity", &resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByTypeAndValueValueRange(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return RangeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.valueRange", nil, htmlAttrs)
	}
	return RangeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.valueRange", &resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.ValueRange, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByTypeAndValueValueReference(frs []FhirResource, numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return ReferenceInput(frs, "characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.valueReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.valueReference", &resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.ValueReference, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByTypeAndValueValueId(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.valueId", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.valueId", &resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.ValueId, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByTypeAndValueOffset(numCharacteristic int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.offset", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByTypeAndValue.offset", resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.Offset, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByCombinationCode(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSCharacteristic_combination

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByCombination.code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByCombination.code", &resource.Characteristic[numCharacteristic].DefinitionByCombination.Code, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByCombinationThreshold(numCharacteristic int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return IntInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByCombination.threshold", nil, htmlAttrs)
	}
	return IntInput("characteristic["+strconv.Itoa(numCharacteristic)+"].definitionByCombination.threshold", resource.Characteristic[numCharacteristic].DefinitionByCombination.Threshold, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromEventDescription(numCharacteristic int, numTimeFromEvent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numTimeFromEvent >= len(resource.Characteristic[numCharacteristic].TimeFromEvent) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].description", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].description", resource.Characteristic[numCharacteristic].TimeFromEvent[numTimeFromEvent].Description, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromEventNote(numCharacteristic int, numTimeFromEvent int, numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numTimeFromEvent >= len(resource.Characteristic[numCharacteristic].TimeFromEvent) || numNote >= len(resource.Characteristic[numCharacteristic].TimeFromEvent[numTimeFromEvent].Note) {
		return AnnotationTextArea("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].note["+strconv.Itoa(numNote)+"]", &resource.Characteristic[numCharacteristic].TimeFromEvent[numTimeFromEvent].Note[numNote], htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromEventEventCodeableConcept(numCharacteristic int, numTimeFromEvent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numTimeFromEvent >= len(resource.Characteristic[numCharacteristic].TimeFromEvent) {
		return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].eventCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].eventCodeableConcept", resource.Characteristic[numCharacteristic].TimeFromEvent[numTimeFromEvent].EventCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromEventEventReference(frs []FhirResource, numCharacteristic int, numTimeFromEvent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numTimeFromEvent >= len(resource.Characteristic[numCharacteristic].TimeFromEvent) {
		return ReferenceInput(frs, "characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].eventReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].eventReference", resource.Characteristic[numCharacteristic].TimeFromEvent[numTimeFromEvent].EventReference, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromEventEventDateTime(numCharacteristic int, numTimeFromEvent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numTimeFromEvent >= len(resource.Characteristic[numCharacteristic].TimeFromEvent) {
		return FhirDateTimeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].eventDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].eventDateTime", resource.Characteristic[numCharacteristic].TimeFromEvent[numTimeFromEvent].EventDateTime, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromEventEventId(numCharacteristic int, numTimeFromEvent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numTimeFromEvent >= len(resource.Characteristic[numCharacteristic].TimeFromEvent) {
		return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].eventId", nil, htmlAttrs)
	}
	return StringInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].eventId", resource.Characteristic[numCharacteristic].TimeFromEvent[numTimeFromEvent].EventId, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromEventQuantity(numCharacteristic int, numTimeFromEvent int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numTimeFromEvent >= len(resource.Characteristic[numCharacteristic].TimeFromEvent) {
		return QuantityInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].quantity", resource.Characteristic[numCharacteristic].TimeFromEvent[numTimeFromEvent].Quantity, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromEventRange(numCharacteristic int, numTimeFromEvent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numTimeFromEvent >= len(resource.Characteristic[numCharacteristic].TimeFromEvent) {
		return RangeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].range", nil, htmlAttrs)
	}
	return RangeInput("characteristic["+strconv.Itoa(numCharacteristic)+"].timeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].range", resource.Characteristic[numCharacteristic].TimeFromEvent[numTimeFromEvent].Range, htmlAttrs)
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
