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

// http://hl7.org/fhir/r5/StructureDefinition/QuestionnaireResponse
type QuestionnaireResponse struct {
	Id                *string                     `json:"id,omitempty"`
	Meta              *Meta                       `json:"meta,omitempty"`
	ImplicitRules     *string                     `json:"implicitRules,omitempty"`
	Language          *string                     `json:"language,omitempty"`
	Text              *Narrative                  `json:"text,omitempty"`
	Contained         []Resource                  `json:"contained,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `json:"identifier,omitempty"`
	BasedOn           []Reference                 `json:"basedOn,omitempty"`
	PartOf            []Reference                 `json:"partOf,omitempty"`
	Questionnaire     string                      `json:"questionnaire"`
	Status            string                      `json:"status"`
	Subject           *Reference                  `json:"subject,omitempty"`
	Encounter         *Reference                  `json:"encounter,omitempty"`
	Authored          *FhirDateTime               `json:"authored,omitempty"`
	Author            *Reference                  `json:"author,omitempty"`
	Source            *Reference                  `json:"source,omitempty"`
	Item              []QuestionnaireResponseItem `json:"item,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/QuestionnaireResponse
type QuestionnaireResponseItem struct {
	Id                *string                           `json:"id,omitempty"`
	Extension         []Extension                       `json:"extension,omitempty"`
	ModifierExtension []Extension                       `json:"modifierExtension,omitempty"`
	LinkId            string                            `json:"linkId"`
	Definition        *string                           `json:"definition,omitempty"`
	Text              *string                           `json:"text,omitempty"`
	Answer            []QuestionnaireResponseItemAnswer `json:"answer,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/QuestionnaireResponse
type QuestionnaireResponseItemAnswer struct {
	Id                *string      `json:"id,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	ValueBoolean      bool         `json:"valueBoolean"`
	ValueDecimal      float64      `json:"valueDecimal"`
	ValueInteger      int          `json:"valueInteger"`
	ValueDate         FhirDate     `json:"valueDate"`
	ValueDateTime     FhirDateTime `json:"valueDateTime"`
	ValueTime         string       `json:"valueTime"`
	ValueString       string       `json:"valueString"`
	ValueUri          string       `json:"valueUri"`
	ValueAttachment   Attachment   `json:"valueAttachment"`
	ValueCoding       Coding       `json:"valueCoding"`
	ValueQuantity     Quantity     `json:"valueQuantity"`
	ValueReference    Reference    `json:"valueReference"`
}

type OtherQuestionnaireResponse QuestionnaireResponse

// struct -> json, automatically add resourceType=Patient
func (r QuestionnaireResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherQuestionnaireResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherQuestionnaireResponse: OtherQuestionnaireResponse(r),
		ResourceType:               "QuestionnaireResponse",
	})
}

// json -> struct, first reject if resourceType != QuestionnaireResponse
func (r *QuestionnaireResponse) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "QuestionnaireResponse" {
		return errors.New("resourceType not QuestionnaireResponse")
	}
	return json.Unmarshal(data, (*OtherQuestionnaireResponse)(r))
}

func (r QuestionnaireResponse) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "QuestionnaireResponse/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "QuestionnaireResponse"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *QuestionnaireResponse) T_BasedOn(frs []FhirResource, numBasedOn int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBasedOn >= len(resource.BasedOn) {
		return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "basedOn["+strconv.Itoa(numBasedOn)+"]", &resource.BasedOn[numBasedOn], htmlAttrs)
}
func (resource *QuestionnaireResponse) T_PartOf(frs []FhirResource, numPartOf int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPartOf >= len(resource.PartOf) {
		return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "partOf["+strconv.Itoa(numPartOf)+"]", &resource.PartOf[numPartOf], htmlAttrs)
}
func (resource *QuestionnaireResponse) T_Questionnaire(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("questionnaire", nil, htmlAttrs)
	}
	return StringInput("questionnaire", &resource.Questionnaire, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSQuestionnaire_answers_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_Subject(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "subject", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "subject", resource.Subject, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_Encounter(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "encounter", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "encounter", resource.Encounter, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_Authored(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("authored", nil, htmlAttrs)
	}
	return FhirDateTimeInput("authored", resource.Authored, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_Author(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "author", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "author", resource.Author, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_Source(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "source", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "source", resource.Source, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemLinkId(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("item["+strconv.Itoa(numItem)+"].linkId", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].linkId", &resource.Item[numItem].LinkId, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemDefinition(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("item["+strconv.Itoa(numItem)+"].definition", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].definition", resource.Item[numItem].Definition, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueBoolean(numItem int, numAnswer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return BoolInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueBoolean", &resource.Item[numItem].Answer[numAnswer].ValueBoolean, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueDecimal(numItem int, numAnswer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return Float64Input("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueDecimal", nil, htmlAttrs)
	}
	return Float64Input("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueDecimal", &resource.Item[numItem].Answer[numAnswer].ValueDecimal, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueInteger(numItem int, numAnswer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return IntInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueInteger", &resource.Item[numItem].Answer[numAnswer].ValueInteger, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueDate(numItem int, numAnswer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return FhirDateInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueDate", nil, htmlAttrs)
	}
	return FhirDateInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueDate", &resource.Item[numItem].Answer[numAnswer].ValueDate, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueDateTime(numItem int, numAnswer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return FhirDateTimeInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueDateTime", &resource.Item[numItem].Answer[numAnswer].ValueDateTime, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueTime(numItem int, numAnswer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return StringInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueTime", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueTime", &resource.Item[numItem].Answer[numAnswer].ValueTime, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueString(numItem int, numAnswer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return StringInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueString", &resource.Item[numItem].Answer[numAnswer].ValueString, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueUri(numItem int, numAnswer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return StringInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueUri", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueUri", &resource.Item[numItem].Answer[numAnswer].ValueUri, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueAttachment(numItem int, numAnswer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return AttachmentInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueAttachment", &resource.Item[numItem].Answer[numAnswer].ValueAttachment, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueCoding(numItem int, numAnswer int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return CodingSelect("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueCoding", &resource.Item[numItem].Answer[numAnswer].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueQuantity(numItem int, numAnswer int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return QuantityInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueQuantity", &resource.Item[numItem].Answer[numAnswer].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueReference(frs []FhirResource, numItem int, numAnswer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return ReferenceInput(frs, "item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "item["+strconv.Itoa(numItem)+"].answer["+strconv.Itoa(numAnswer)+"].valueReference", &resource.Item[numItem].Answer[numAnswer].ValueReference, htmlAttrs)
}
