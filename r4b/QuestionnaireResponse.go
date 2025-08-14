//generated August 14 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r4b

import "encoding/json"

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
	ValueBoolean      *bool       `json:"valueBoolean"`
	ValueDecimal      *float64    `json:"valueDecimal"`
	ValueInteger      *int        `json:"valueInteger"`
	ValueDate         *string     `json:"valueDate"`
	ValueDateTime     *string     `json:"valueDateTime"`
	ValueTime         *string     `json:"valueTime"`
	ValueString       *string     `json:"valueString"`
	ValueUri          *string     `json:"valueUri"`
	ValueAttachment   *Attachment `json:"valueAttachment"`
	ValueCoding       *Coding     `json:"valueCoding"`
	ValueQuantity     *Quantity   `json:"valueQuantity"`
	ValueReference    *Reference  `json:"valueReference"`
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
