package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/QuestionnaireResponse
type QuestionnaireResponse struct {
	Id                *string                     `json:"id,omitempty"`
	Meta              *Meta                       `json:"meta,omitempty"`
	ImplicitRules     *string                     `json:"implicitRules,omitempty"`
	Language          *string                     `json:"language,omitempty"`
	Text              *Narrative                  `json:"text,omitempty"`
	Contained         []Resource                  `json:"contained,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Identifier        *Identifier                 `json:"identifier,omitempty"`
	BasedOn           []Reference                 `json:"basedOn,omitempty"`
	PartOf            []Reference                 `json:"partOf,omitempty"`
	Questionnaire     *string                     `json:"questionnaire,omitempty"`
	Status            string                      `json:"status"`
	Subject           *Reference                  `json:"subject,omitempty"`
	Encounter         *Reference                  `json:"encounter,omitempty"`
	Authored          *string                     `json:"authored,omitempty"`
	Author            *Reference                  `json:"author,omitempty"`
	Source            *Reference                  `json:"source,omitempty"`
	Item              []QuestionnaireResponseItem `json:"item,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/QuestionnaireResponse
type QuestionnaireResponseItem struct {
	Id                *string                           `json:"id,omitempty"`
	Extension         []Extension                       `json:"extension,omitempty"`
	ModifierExtension []Extension                       `json:"modifierExtension,omitempty"`
	LinkId            string                            `json:"linkId"`
	Definition        *string                           `json:"definition,omitempty"`
	Text              *string                           `json:"text,omitempty"`
	Answer            []QuestionnaireResponseItemAnswer `json:"answer,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/QuestionnaireResponse
type QuestionnaireResponseItemAnswer struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ValueBoolean      *bool       `json:"valueBoolean,omitempty"`
	ValueDecimal      *float64    `json:"valueDecimal,omitempty"`
	ValueInteger      *int        `json:"valueInteger,omitempty"`
	ValueDate         *string     `json:"valueDate,omitempty"`
	ValueDateTime     *string     `json:"valueDateTime,omitempty"`
	ValueTime         *string     `json:"valueTime,omitempty"`
	ValueString       *string     `json:"valueString,omitempty"`
	ValueUri          *string     `json:"valueUri,omitempty"`
	ValueAttachment   *Attachment `json:"valueAttachment,omitempty"`
	ValueCoding       *Coding     `json:"valueCoding,omitempty"`
	ValueQuantity     *Quantity   `json:"valueQuantity,omitempty"`
	ValueReference    *Reference  `json:"valueReference,omitempty"`
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

func (resource *QuestionnaireResponse) T_Id() templ.Component {

	if resource == nil {
		return StringInput("QuestionnaireResponse.Id", nil)
	}
	return StringInput("QuestionnaireResponse.Id", resource.Id)
}
func (resource *QuestionnaireResponse) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("QuestionnaireResponse.ImplicitRules", nil)
	}
	return StringInput("QuestionnaireResponse.ImplicitRules", resource.ImplicitRules)
}
func (resource *QuestionnaireResponse) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("QuestionnaireResponse.Language", nil, optionsValueSet)
	}
	return CodeSelect("QuestionnaireResponse.Language", resource.Language, optionsValueSet)
}
func (resource *QuestionnaireResponse) T_Questionnaire() templ.Component {

	if resource == nil {
		return StringInput("QuestionnaireResponse.Questionnaire", nil)
	}
	return StringInput("QuestionnaireResponse.Questionnaire", resource.Questionnaire)
}
func (resource *QuestionnaireResponse) T_Status() templ.Component {
	optionsValueSet := VSQuestionnaire_answers_status

	if resource == nil {
		return CodeSelect("QuestionnaireResponse.Status", nil, optionsValueSet)
	}
	return CodeSelect("QuestionnaireResponse.Status", &resource.Status, optionsValueSet)
}
func (resource *QuestionnaireResponse) T_Authored() templ.Component {

	if resource == nil {
		return StringInput("QuestionnaireResponse.Authored", nil)
	}
	return StringInput("QuestionnaireResponse.Authored", resource.Authored)
}
func (resource *QuestionnaireResponse) T_ItemId(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("QuestionnaireResponse.Item["+strconv.Itoa(numItem)+"].Id", nil)
	}
	return StringInput("QuestionnaireResponse.Item["+strconv.Itoa(numItem)+"].Id", resource.Item[numItem].Id)
}
func (resource *QuestionnaireResponse) T_ItemLinkId(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("QuestionnaireResponse.Item["+strconv.Itoa(numItem)+"].LinkId", nil)
	}
	return StringInput("QuestionnaireResponse.Item["+strconv.Itoa(numItem)+"].LinkId", &resource.Item[numItem].LinkId)
}
func (resource *QuestionnaireResponse) T_ItemDefinition(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("QuestionnaireResponse.Item["+strconv.Itoa(numItem)+"].Definition", nil)
	}
	return StringInput("QuestionnaireResponse.Item["+strconv.Itoa(numItem)+"].Definition", resource.Item[numItem].Definition)
}
func (resource *QuestionnaireResponse) T_ItemText(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("QuestionnaireResponse.Item["+strconv.Itoa(numItem)+"].Text", nil)
	}
	return StringInput("QuestionnaireResponse.Item["+strconv.Itoa(numItem)+"].Text", resource.Item[numItem].Text)
}
func (resource *QuestionnaireResponse) T_ItemAnswerId(numItem int, numAnswer int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Answer) >= numAnswer {
		return StringInput("QuestionnaireResponse.Item["+strconv.Itoa(numItem)+"].Answer["+strconv.Itoa(numAnswer)+"].Id", nil)
	}
	return StringInput("QuestionnaireResponse.Item["+strconv.Itoa(numItem)+"].Answer["+strconv.Itoa(numAnswer)+"].Id", resource.Item[numItem].Answer[numAnswer].Id)
}
