package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Questionnaire
type Questionnaire struct {
	Id                *string             `json:"id,omitempty"`
	Meta              *Meta               `json:"meta,omitempty"`
	ImplicitRules     *string             `json:"implicitRules,omitempty"`
	Language          *string             `json:"language,omitempty"`
	Text              *Narrative          `json:"text,omitempty"`
	Contained         []Resource          `json:"contained,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Url               *string             `json:"url,omitempty"`
	Identifier        []Identifier        `json:"identifier,omitempty"`
	Version           *string             `json:"version,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Title             *string             `json:"title,omitempty"`
	DerivedFrom       []string            `json:"derivedFrom,omitempty"`
	Status            string              `json:"status"`
	Experimental      *bool               `json:"experimental,omitempty"`
	SubjectType       []string            `json:"subjectType,omitempty"`
	Date              *string             `json:"date,omitempty"`
	Publisher         *string             `json:"publisher,omitempty"`
	Contact           []ContactDetail     `json:"contact,omitempty"`
	Description       *string             `json:"description,omitempty"`
	UseContext        []UsageContext      `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept   `json:"jurisdiction,omitempty"`
	Purpose           *string             `json:"purpose,omitempty"`
	Copyright         *string             `json:"copyright,omitempty"`
	ApprovalDate      *string             `json:"approvalDate,omitempty"`
	LastReviewDate    *string             `json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period             `json:"effectivePeriod,omitempty"`
	Code              []Coding            `json:"code,omitempty"`
	Item              []QuestionnaireItem `json:"item,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Questionnaire
type QuestionnaireItem struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	LinkId            string                          `json:"linkId"`
	Definition        *string                         `json:"definition,omitempty"`
	Code              []Coding                        `json:"code,omitempty"`
	Prefix            *string                         `json:"prefix,omitempty"`
	Text              *string                         `json:"text,omitempty"`
	Type              string                          `json:"type"`
	EnableWhen        []QuestionnaireItemEnableWhen   `json:"enableWhen,omitempty"`
	EnableBehavior    *string                         `json:"enableBehavior,omitempty"`
	Required          *bool                           `json:"required,omitempty"`
	Repeats           *bool                           `json:"repeats,omitempty"`
	ReadOnly          *bool                           `json:"readOnly,omitempty"`
	MaxLength         *int                            `json:"maxLength,omitempty"`
	AnswerValueSet    *string                         `json:"answerValueSet,omitempty"`
	AnswerOption      []QuestionnaireItemAnswerOption `json:"answerOption,omitempty"`
	Initial           []QuestionnaireItemInitial      `json:"initial,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Questionnaire
type QuestionnaireItemEnableWhen struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Question          string      `json:"question"`
	Operator          string      `json:"operator"`
	AnswerBoolean     bool        `json:"answerBoolean"`
	AnswerDecimal     float64     `json:"answerDecimal"`
	AnswerInteger     int         `json:"answerInteger"`
	AnswerDate        string      `json:"answerDate"`
	AnswerDateTime    string      `json:"answerDateTime"`
	AnswerTime        string      `json:"answerTime"`
	AnswerString      string      `json:"answerString"`
	AnswerCoding      Coding      `json:"answerCoding"`
	AnswerQuantity    Quantity    `json:"answerQuantity"`
	AnswerReference   Reference   `json:"answerReference"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Questionnaire
type QuestionnaireItemAnswerOption struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ValueInteger      int         `json:"valueInteger"`
	ValueDate         string      `json:"valueDate"`
	ValueTime         string      `json:"valueTime"`
	ValueString       string      `json:"valueString"`
	ValueCoding       Coding      `json:"valueCoding"`
	ValueReference    Reference   `json:"valueReference"`
	InitialSelected   *bool       `json:"initialSelected,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Questionnaire
type QuestionnaireItemInitial struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ValueBoolean      bool        `json:"valueBoolean"`
	ValueDecimal      float64     `json:"valueDecimal"`
	ValueInteger      int         `json:"valueInteger"`
	ValueDate         string      `json:"valueDate"`
	ValueDateTime     string      `json:"valueDateTime"`
	ValueTime         string      `json:"valueTime"`
	ValueString       string      `json:"valueString"`
	ValueUri          string      `json:"valueUri"`
	ValueAttachment   Attachment  `json:"valueAttachment"`
	ValueCoding       Coding      `json:"valueCoding"`
	ValueQuantity     Quantity    `json:"valueQuantity"`
	ValueReference    Reference   `json:"valueReference"`
}

type OtherQuestionnaire Questionnaire

// on convert struct to json, automatically add resourceType=Questionnaire
func (r Questionnaire) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherQuestionnaire
		ResourceType string `json:"resourceType"`
	}{
		OtherQuestionnaire: OtherQuestionnaire(r),
		ResourceType:       "Questionnaire",
	})
}
func (r Questionnaire) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Questionnaire/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Questionnaire"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Questionnaire) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *Questionnaire) T_Version(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *Questionnaire) T_Name(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *Questionnaire) T_Title(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *Questionnaire) T_DerivedFrom(numDerivedFrom int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDerivedFrom >= len(resource.DerivedFrom) {
		return StringInput("derivedFrom["+strconv.Itoa(numDerivedFrom)+"]", nil, htmlAttrs)
	}
	return StringInput("derivedFrom["+strconv.Itoa(numDerivedFrom)+"]", &resource.DerivedFrom[numDerivedFrom], htmlAttrs)
}
func (resource *Questionnaire) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_Experimental(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("experimental", nil, htmlAttrs)
	}
	return BoolInput("experimental", resource.Experimental, htmlAttrs)
}
func (resource *Questionnaire) T_SubjectType(numSubjectType int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numSubjectType >= len(resource.SubjectType) {
		return CodeSelect("subjectType["+strconv.Itoa(numSubjectType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("subjectType["+strconv.Itoa(numSubjectType)+"]", &resource.SubjectType[numSubjectType], optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("date", nil, htmlAttrs)
	}
	return DateTimeInput("date", resource.Date, htmlAttrs)
}
func (resource *Questionnaire) T_Publisher(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("publisher", nil, htmlAttrs)
	}
	return StringInput("publisher", resource.Publisher, htmlAttrs)
}
func (resource *Questionnaire) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *Questionnaire) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_Purpose(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("purpose", nil, htmlAttrs)
	}
	return StringInput("purpose", resource.Purpose, htmlAttrs)
}
func (resource *Questionnaire) T_Copyright(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("copyright", nil, htmlAttrs)
	}
	return StringInput("copyright", resource.Copyright, htmlAttrs)
}
func (resource *Questionnaire) T_ApprovalDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateInput("approvalDate", nil, htmlAttrs)
	}
	return DateInput("approvalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *Questionnaire) T_LastReviewDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateInput("lastReviewDate", nil, htmlAttrs)
	}
	return DateInput("lastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *Questionnaire) T_Code(numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCode >= len(resource.Code) {
		return CodingSelect("code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("code["+strconv.Itoa(numCode)+"]", &resource.Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_ItemLinkId(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("item["+strconv.Itoa(numItem)+"].linkId", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].linkId", &resource.Item[numItem].LinkId, htmlAttrs)
}
func (resource *Questionnaire) T_ItemDefinition(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("item["+strconv.Itoa(numItem)+"].definition", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].definition", resource.Item[numItem].Definition, htmlAttrs)
}
func (resource *Questionnaire) T_ItemCode(numItem int, numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numCode >= len(resource.Item[numItem].Code) {
		return CodingSelect("item["+strconv.Itoa(numItem)+"].code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("item["+strconv.Itoa(numItem)+"].code["+strconv.Itoa(numCode)+"]", &resource.Item[numItem].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_ItemPrefix(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("item["+strconv.Itoa(numItem)+"].prefix", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].prefix", resource.Item[numItem].Prefix, htmlAttrs)
}
func (resource *Questionnaire) T_ItemText(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("item["+strconv.Itoa(numItem)+"].text", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].text", resource.Item[numItem].Text, htmlAttrs)
}
func (resource *Questionnaire) T_ItemType(numItem int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSItem_type

	if resource == nil || numItem >= len(resource.Item) {
		return CodeSelect("item["+strconv.Itoa(numItem)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("item["+strconv.Itoa(numItem)+"].type", &resource.Item[numItem].Type, optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableBehavior(numItem int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSQuestionnaire_enable_behavior

	if resource == nil || numItem >= len(resource.Item) {
		return CodeSelect("item["+strconv.Itoa(numItem)+"].enableBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("item["+strconv.Itoa(numItem)+"].enableBehavior", resource.Item[numItem].EnableBehavior, optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_ItemRequired(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return BoolInput("item["+strconv.Itoa(numItem)+"].required", nil, htmlAttrs)
	}
	return BoolInput("item["+strconv.Itoa(numItem)+"].required", resource.Item[numItem].Required, htmlAttrs)
}
func (resource *Questionnaire) T_ItemRepeats(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return BoolInput("item["+strconv.Itoa(numItem)+"].repeats", nil, htmlAttrs)
	}
	return BoolInput("item["+strconv.Itoa(numItem)+"].repeats", resource.Item[numItem].Repeats, htmlAttrs)
}
func (resource *Questionnaire) T_ItemReadOnly(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return BoolInput("item["+strconv.Itoa(numItem)+"].readOnly", nil, htmlAttrs)
	}
	return BoolInput("item["+strconv.Itoa(numItem)+"].readOnly", resource.Item[numItem].ReadOnly, htmlAttrs)
}
func (resource *Questionnaire) T_ItemMaxLength(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return IntInput("item["+strconv.Itoa(numItem)+"].maxLength", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].maxLength", resource.Item[numItem].MaxLength, htmlAttrs)
}
func (resource *Questionnaire) T_ItemAnswerValueSet(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("item["+strconv.Itoa(numItem)+"].answerValueSet", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].answerValueSet", resource.Item[numItem].AnswerValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenQuestion(numItem int, numEnableWhen int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return StringInput("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].question", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].question", &resource.Item[numItem].EnableWhen[numEnableWhen].Question, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenOperator(numItem int, numEnableWhen int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSQuestionnaire_enable_operator

	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return CodeSelect("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].operator", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].operator", &resource.Item[numItem].EnableWhen[numEnableWhen].Operator, optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenAnswerBoolean(numItem int, numEnableWhen int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return BoolInput("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].answerBoolean", nil, htmlAttrs)
	}
	return BoolInput("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].answerBoolean", &resource.Item[numItem].EnableWhen[numEnableWhen].AnswerBoolean, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenAnswerDecimal(numItem int, numEnableWhen int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return Float64Input("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].answerDecimal", nil, htmlAttrs)
	}
	return Float64Input("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].answerDecimal", &resource.Item[numItem].EnableWhen[numEnableWhen].AnswerDecimal, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenAnswerInteger(numItem int, numEnableWhen int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return IntInput("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].answerInteger", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].answerInteger", &resource.Item[numItem].EnableWhen[numEnableWhen].AnswerInteger, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenAnswerDate(numItem int, numEnableWhen int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return DateInput("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].answerDate", nil, htmlAttrs)
	}
	return DateInput("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].answerDate", &resource.Item[numItem].EnableWhen[numEnableWhen].AnswerDate, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenAnswerDateTime(numItem int, numEnableWhen int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return DateTimeInput("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].answerDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].answerDateTime", &resource.Item[numItem].EnableWhen[numEnableWhen].AnswerDateTime, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenAnswerTime(numItem int, numEnableWhen int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return StringInput("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].answerTime", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].answerTime", &resource.Item[numItem].EnableWhen[numEnableWhen].AnswerTime, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenAnswerString(numItem int, numEnableWhen int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return StringInput("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].answerString", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].answerString", &resource.Item[numItem].EnableWhen[numEnableWhen].AnswerString, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenAnswerCoding(numItem int, numEnableWhen int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return CodingSelect("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].answerCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("item["+strconv.Itoa(numItem)+"].enableWhen["+strconv.Itoa(numEnableWhen)+"].answerCoding", &resource.Item[numItem].EnableWhen[numEnableWhen].AnswerCoding, optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_ItemAnswerOptionValueInteger(numItem int, numAnswerOption int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswerOption >= len(resource.Item[numItem].AnswerOption) {
		return IntInput("item["+strconv.Itoa(numItem)+"].answerOption["+strconv.Itoa(numAnswerOption)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].answerOption["+strconv.Itoa(numAnswerOption)+"].valueInteger", &resource.Item[numItem].AnswerOption[numAnswerOption].ValueInteger, htmlAttrs)
}
func (resource *Questionnaire) T_ItemAnswerOptionValueDate(numItem int, numAnswerOption int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswerOption >= len(resource.Item[numItem].AnswerOption) {
		return DateInput("item["+strconv.Itoa(numItem)+"].answerOption["+strconv.Itoa(numAnswerOption)+"].valueDate", nil, htmlAttrs)
	}
	return DateInput("item["+strconv.Itoa(numItem)+"].answerOption["+strconv.Itoa(numAnswerOption)+"].valueDate", &resource.Item[numItem].AnswerOption[numAnswerOption].ValueDate, htmlAttrs)
}
func (resource *Questionnaire) T_ItemAnswerOptionValueTime(numItem int, numAnswerOption int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswerOption >= len(resource.Item[numItem].AnswerOption) {
		return StringInput("item["+strconv.Itoa(numItem)+"].answerOption["+strconv.Itoa(numAnswerOption)+"].valueTime", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].answerOption["+strconv.Itoa(numAnswerOption)+"].valueTime", &resource.Item[numItem].AnswerOption[numAnswerOption].ValueTime, htmlAttrs)
}
func (resource *Questionnaire) T_ItemAnswerOptionValueString(numItem int, numAnswerOption int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswerOption >= len(resource.Item[numItem].AnswerOption) {
		return StringInput("item["+strconv.Itoa(numItem)+"].answerOption["+strconv.Itoa(numAnswerOption)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].answerOption["+strconv.Itoa(numAnswerOption)+"].valueString", &resource.Item[numItem].AnswerOption[numAnswerOption].ValueString, htmlAttrs)
}
func (resource *Questionnaire) T_ItemAnswerOptionValueCoding(numItem int, numAnswerOption int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswerOption >= len(resource.Item[numItem].AnswerOption) {
		return CodingSelect("item["+strconv.Itoa(numItem)+"].answerOption["+strconv.Itoa(numAnswerOption)+"].valueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("item["+strconv.Itoa(numItem)+"].answerOption["+strconv.Itoa(numAnswerOption)+"].valueCoding", &resource.Item[numItem].AnswerOption[numAnswerOption].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_ItemAnswerOptionInitialSelected(numItem int, numAnswerOption int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAnswerOption >= len(resource.Item[numItem].AnswerOption) {
		return BoolInput("item["+strconv.Itoa(numItem)+"].answerOption["+strconv.Itoa(numAnswerOption)+"].initialSelected", nil, htmlAttrs)
	}
	return BoolInput("item["+strconv.Itoa(numItem)+"].answerOption["+strconv.Itoa(numAnswerOption)+"].initialSelected", resource.Item[numItem].AnswerOption[numAnswerOption].InitialSelected, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueBoolean(numItem int, numInitial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return BoolInput("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueBoolean", &resource.Item[numItem].Initial[numInitial].ValueBoolean, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueDecimal(numItem int, numInitial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return Float64Input("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueDecimal", nil, htmlAttrs)
	}
	return Float64Input("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueDecimal", &resource.Item[numItem].Initial[numInitial].ValueDecimal, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueInteger(numItem int, numInitial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return IntInput("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueInteger", &resource.Item[numItem].Initial[numInitial].ValueInteger, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueDate(numItem int, numInitial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return DateInput("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueDate", nil, htmlAttrs)
	}
	return DateInput("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueDate", &resource.Item[numItem].Initial[numInitial].ValueDate, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueDateTime(numItem int, numInitial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return DateTimeInput("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueDateTime", &resource.Item[numItem].Initial[numInitial].ValueDateTime, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueTime(numItem int, numInitial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return StringInput("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueTime", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueTime", &resource.Item[numItem].Initial[numInitial].ValueTime, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueString(numItem int, numInitial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return StringInput("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueString", &resource.Item[numItem].Initial[numInitial].ValueString, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueUri(numItem int, numInitial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return StringInput("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueUri", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueUri", &resource.Item[numItem].Initial[numInitial].ValueUri, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueCoding(numItem int, numInitial int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return CodingSelect("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("item["+strconv.Itoa(numItem)+"].initial["+strconv.Itoa(numInitial)+"].valueCoding", &resource.Item[numItem].Initial[numInitial].ValueCoding, optionsValueSet, htmlAttrs)
}
