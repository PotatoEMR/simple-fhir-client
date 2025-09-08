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
	Authored          *time.Time                  `json:"authored,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ValueBoolean      bool        `json:"valueBoolean"`
	ValueDecimal      float64     `json:"valueDecimal"`
	ValueInteger      int         `json:"valueInteger"`
	ValueDate         time.Time   `json:"valueDate,format:'2006-01-02'"`
	ValueDateTime     time.Time   `json:"valueDateTime,format:'2006-01-02T15:04:05Z07:00'"`
	ValueTime         string      `json:"valueTime"`
	ValueString       string      `json:"valueString"`
	ValueUri          string      `json:"valueUri"`
	ValueAttachment   Attachment  `json:"valueAttachment"`
	ValueCoding       Coding      `json:"valueCoding"`
	ValueQuantity     Quantity    `json:"valueQuantity"`
	ValueReference    Reference   `json:"valueReference"`
}

type OtherQuestionnaireResponse QuestionnaireResponse

// on convert struct to json, automatically add resourceType=QuestionnaireResponse
func (r QuestionnaireResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherQuestionnaireResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherQuestionnaireResponse: OtherQuestionnaireResponse(r),
		ResourceType:               "QuestionnaireResponse",
	})
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
func (resource *QuestionnaireResponse) T_Questionnaire(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("QuestionnaireResponse.Questionnaire", nil, htmlAttrs)
	}
	return StringInput("QuestionnaireResponse.Questionnaire", &resource.Questionnaire, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSQuestionnaire_answers_status

	if resource == nil {
		return CodeSelect("QuestionnaireResponse.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("QuestionnaireResponse.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_Authored(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("QuestionnaireResponse.Authored", nil, htmlAttrs)
	}
	return DateTimeInput("QuestionnaireResponse.Authored", resource.Authored, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemLinkId(numItem int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..LinkId", nil, htmlAttrs)
	}
	return StringInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..LinkId", &resource.Item[numItem].LinkId, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemDefinition(numItem int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Definition", nil, htmlAttrs)
	}
	return StringInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Definition", resource.Item[numItem].Definition, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemText(numItem int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Text", nil, htmlAttrs)
	}
	return StringInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Text", resource.Item[numItem].Text, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueBoolean(numItem int, numAnswer int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return BoolInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueBoolean", &resource.Item[numItem].Answer[numAnswer].ValueBoolean, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueDecimal(numItem int, numAnswer int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return Float64Input("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueDecimal", nil, htmlAttrs)
	}
	return Float64Input("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueDecimal", &resource.Item[numItem].Answer[numAnswer].ValueDecimal, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueInteger(numItem int, numAnswer int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return IntInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueInteger", nil, htmlAttrs)
	}
	return IntInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueInteger", &resource.Item[numItem].Answer[numAnswer].ValueInteger, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueDate(numItem int, numAnswer int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return DateInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueDate", nil, htmlAttrs)
	}
	return DateInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueDate", &resource.Item[numItem].Answer[numAnswer].ValueDate, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueDateTime(numItem int, numAnswer int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return DateTimeInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueDateTime", &resource.Item[numItem].Answer[numAnswer].ValueDateTime, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueTime(numItem int, numAnswer int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return StringInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueTime", nil, htmlAttrs)
	}
	return StringInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueTime", &resource.Item[numItem].Answer[numAnswer].ValueTime, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueString(numItem int, numAnswer int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return StringInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueString", nil, htmlAttrs)
	}
	return StringInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueString", &resource.Item[numItem].Answer[numAnswer].ValueString, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueUri(numItem int, numAnswer int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return StringInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueUri", nil, htmlAttrs)
	}
	return StringInput("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueUri", &resource.Item[numItem].Answer[numAnswer].ValueUri, htmlAttrs)
}
func (resource *QuestionnaireResponse) T_ItemAnswerValueCoding(numItem int, numAnswer int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numAnswer >= len(resource.Item[numItem].Answer) {
		return CodingSelect("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("QuestionnaireResponse.Item."+strconv.Itoa(numItem)+"..Answer."+strconv.Itoa(numAnswer)+"..ValueCoding", &resource.Item[numItem].Answer[numAnswer].ValueCoding, optionsValueSet, htmlAttrs)
}
