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
	Date                   *time.Time                       `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher              *string                          `json:"publisher,omitempty"`
	Contact                []ContactDetail                  `json:"contact,omitempty"`
	Description            *string                          `json:"description,omitempty"`
	Note                   []Annotation                     `json:"note,omitempty"`
	UseContext             []UsageContext                   `json:"useContext,omitempty"`
	Purpose                *string                          `json:"purpose,omitempty"`
	Copyright              *string                          `json:"copyright,omitempty"`
	CopyrightLabel         *string                          `json:"copyrightLabel,omitempty"`
	ApprovalDate           *time.Time                       `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate         *time.Time                       `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
	EventDateTime        *time.Time       `json:"eventDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
		return StringInput("Url", nil, htmlAttrs)
	}
	return StringInput("Url", resource.Url, htmlAttrs)
}
func (resource *EvidenceVariable) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Version", nil, htmlAttrs)
	}
	return StringInput("Version", resource.Version, htmlAttrs)
}
func (resource *EvidenceVariable) T_VersionAlgorithmString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("VersionAlgorithmString", nil, htmlAttrs)
	}
	return StringInput("VersionAlgorithmString", resource.VersionAlgorithmString, htmlAttrs)
}
func (resource *EvidenceVariable) T_VersionAlgorithmCoding(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodingSelect("VersionAlgorithmCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("VersionAlgorithmCoding", resource.VersionAlgorithmCoding, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Name", nil, htmlAttrs)
	}
	return StringInput("Name", resource.Name, htmlAttrs)
}
func (resource *EvidenceVariable) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Title", nil, htmlAttrs)
	}
	return StringInput("Title", resource.Title, htmlAttrs)
}
func (resource *EvidenceVariable) T_ShortTitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("ShortTitle", nil, htmlAttrs)
	}
	return StringInput("ShortTitle", resource.ShortTitle, htmlAttrs)
}
func (resource *EvidenceVariable) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_Experimental(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Experimental", nil, htmlAttrs)
	}
	return BoolInput("Experimental", resource.Experimental, htmlAttrs)
}
func (resource *EvidenceVariable) T_Date(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Date", nil, htmlAttrs)
	}
	return DateTimeInput("Date", resource.Date, htmlAttrs)
}
func (resource *EvidenceVariable) T_Publisher(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Publisher", nil, htmlAttrs)
	}
	return StringInput("Publisher", resource.Publisher, htmlAttrs)
}
func (resource *EvidenceVariable) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Description", nil, htmlAttrs)
	}
	return StringInput("Description", resource.Description, htmlAttrs)
}
func (resource *EvidenceVariable) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *EvidenceVariable) T_Purpose(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Purpose", nil, htmlAttrs)
	}
	return StringInput("Purpose", resource.Purpose, htmlAttrs)
}
func (resource *EvidenceVariable) T_Copyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Copyright", nil, htmlAttrs)
	}
	return StringInput("Copyright", resource.Copyright, htmlAttrs)
}
func (resource *EvidenceVariable) T_CopyrightLabel(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("CopyrightLabel", nil, htmlAttrs)
	}
	return StringInput("CopyrightLabel", resource.CopyrightLabel, htmlAttrs)
}
func (resource *EvidenceVariable) T_ApprovalDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *EvidenceVariable) T_LastReviewDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *EvidenceVariable) T_Actual(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("Actual", nil, htmlAttrs)
	}
	return BoolInput("Actual", resource.Actual, htmlAttrs)
}
func (resource *EvidenceVariable) T_Handling(htmlAttrs string) templ.Component {
	optionsValueSet := VSVariable_handling

	if resource == nil {
		return CodeSelect("Handling", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Handling", resource.Handling, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicLinkId(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]LinkId", nil, htmlAttrs)
	}
	return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]LinkId", resource.Characteristic[numCharacteristic].LinkId, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDescription(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]Description", nil, htmlAttrs)
	}
	return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]Description", resource.Characteristic[numCharacteristic].Description, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicNote(numCharacteristic int, numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numNote >= len(resource.Characteristic[numCharacteristic].Note) {
		return AnnotationTextArea("Characteristic["+strconv.Itoa(numCharacteristic)+"]Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Characteristic["+strconv.Itoa(numCharacteristic)+"]Note["+strconv.Itoa(numNote)+"]", &resource.Characteristic[numCharacteristic].Note[numNote], htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicExclude(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]Exclude", nil, htmlAttrs)
	}
	return BoolInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]Exclude", resource.Characteristic[numCharacteristic].Exclude, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionCanonical(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionCanonical", nil, htmlAttrs)
	}
	return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionCanonical", resource.Characteristic[numCharacteristic].DefinitionCanonical, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionCodeableConcept", resource.Characteristic[numCharacteristic].DefinitionCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionId(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionId", nil, htmlAttrs)
	}
	return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionId", resource.Characteristic[numCharacteristic].DefinitionId, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByTypeAndValueType(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionByTypeAndValue.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionByTypeAndValue.Type", &resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.Type, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByTypeAndValueMethod(numCharacteristic int, numMethod int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numMethod >= len(resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.Method) {
		return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionByTypeAndValue.Method["+strconv.Itoa(numMethod)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionByTypeAndValue.Method["+strconv.Itoa(numMethod)+"]", &resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.Method[numMethod], optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByTypeAndValueValueCodeableConcept(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionByTypeAndValue.ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionByTypeAndValue.ValueCodeableConcept", &resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByTypeAndValueValueBoolean(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return BoolInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionByTypeAndValue.ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionByTypeAndValue.ValueBoolean", &resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.ValueBoolean, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByTypeAndValueValueId(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionByTypeAndValue.ValueId", nil, htmlAttrs)
	}
	return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionByTypeAndValue.ValueId", &resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.ValueId, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByTypeAndValueOffset(numCharacteristic int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionByTypeAndValue.Offset", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionByTypeAndValue.Offset", resource.Characteristic[numCharacteristic].DefinitionByTypeAndValue.Offset, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByCombinationCode(numCharacteristic int, htmlAttrs string) templ.Component {
	optionsValueSet := VSCharacteristic_combination

	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return CodeSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionByCombination.Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionByCombination.Code", &resource.Characteristic[numCharacteristic].DefinitionByCombination.Code, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicDefinitionByCombinationThreshold(numCharacteristic int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) {
		return IntInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionByCombination.Threshold", nil, htmlAttrs)
	}
	return IntInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]DefinitionByCombination.Threshold", resource.Characteristic[numCharacteristic].DefinitionByCombination.Threshold, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromEventDescription(numCharacteristic int, numTimeFromEvent int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numTimeFromEvent >= len(resource.Characteristic[numCharacteristic].TimeFromEvent) {
		return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]TimeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].Description", nil, htmlAttrs)
	}
	return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]TimeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].Description", resource.Characteristic[numCharacteristic].TimeFromEvent[numTimeFromEvent].Description, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromEventNote(numCharacteristic int, numTimeFromEvent int, numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numTimeFromEvent >= len(resource.Characteristic[numCharacteristic].TimeFromEvent) || numNote >= len(resource.Characteristic[numCharacteristic].TimeFromEvent[numTimeFromEvent].Note) {
		return AnnotationTextArea("Characteristic["+strconv.Itoa(numCharacteristic)+"]TimeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("Characteristic["+strconv.Itoa(numCharacteristic)+"]TimeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].Note["+strconv.Itoa(numNote)+"]", &resource.Characteristic[numCharacteristic].TimeFromEvent[numTimeFromEvent].Note[numNote], htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromEventEventCodeableConcept(numCharacteristic int, numTimeFromEvent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numTimeFromEvent >= len(resource.Characteristic[numCharacteristic].TimeFromEvent) {
		return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]TimeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].EventCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Characteristic["+strconv.Itoa(numCharacteristic)+"]TimeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].EventCodeableConcept", resource.Characteristic[numCharacteristic].TimeFromEvent[numTimeFromEvent].EventCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromEventEventDateTime(numCharacteristic int, numTimeFromEvent int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numTimeFromEvent >= len(resource.Characteristic[numCharacteristic].TimeFromEvent) {
		return DateTimeInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]TimeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].EventDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]TimeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].EventDateTime", resource.Characteristic[numCharacteristic].TimeFromEvent[numTimeFromEvent].EventDateTime, htmlAttrs)
}
func (resource *EvidenceVariable) T_CharacteristicTimeFromEventEventId(numCharacteristic int, numTimeFromEvent int, htmlAttrs string) templ.Component {
	if resource == nil || numCharacteristic >= len(resource.Characteristic) || numTimeFromEvent >= len(resource.Characteristic[numCharacteristic].TimeFromEvent) {
		return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]TimeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].EventId", nil, htmlAttrs)
	}
	return StringInput("Characteristic["+strconv.Itoa(numCharacteristic)+"]TimeFromEvent["+strconv.Itoa(numTimeFromEvent)+"].EventId", resource.Characteristic[numCharacteristic].TimeFromEvent[numTimeFromEvent].EventId, htmlAttrs)
}
func (resource *EvidenceVariable) T_CategoryName(numCategory int, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return StringInput("Category["+strconv.Itoa(numCategory)+"]Name", nil, htmlAttrs)
	}
	return StringInput("Category["+strconv.Itoa(numCategory)+"]Name", resource.Category[numCategory].Name, htmlAttrs)
}
func (resource *EvidenceVariable) T_CategoryValueCodeableConcept(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]ValueCodeableConcept", resource.Category[numCategory].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
