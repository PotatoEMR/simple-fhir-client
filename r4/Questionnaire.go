package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Date              *time.Time          `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Publisher         *string             `json:"publisher,omitempty"`
	Contact           []ContactDetail     `json:"contact,omitempty"`
	Description       *string             `json:"description,omitempty"`
	UseContext        []UsageContext      `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept   `json:"jurisdiction,omitempty"`
	Purpose           *string             `json:"purpose,omitempty"`
	Copyright         *string             `json:"copyright,omitempty"`
	ApprovalDate      *time.Time          `json:"approvalDate,omitempty,format:'2006-01-02'"`
	LastReviewDate    *time.Time          `json:"lastReviewDate,omitempty,format:'2006-01-02'"`
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
	AnswerDate        time.Time   `json:"answerDate,format:'2006-01-02'"`
	AnswerDateTime    time.Time   `json:"answerDateTime,format:'2006-01-02T15:04:05Z07:00'"`
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
	ValueDate         time.Time   `json:"valueDate,format:'2006-01-02'"`
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
func (resource *Questionnaire) T_Url(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Url", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Url", resource.Url, htmlAttrs)
}
func (resource *Questionnaire) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Version", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Version", resource.Version, htmlAttrs)
}
func (resource *Questionnaire) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Name", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Name", resource.Name, htmlAttrs)
}
func (resource *Questionnaire) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Title", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Title", resource.Title, htmlAttrs)
}
func (resource *Questionnaire) T_DerivedFrom(numDerivedFrom int, htmlAttrs string) templ.Component {

	if resource == nil || numDerivedFrom >= len(resource.DerivedFrom) {
		return StringInput("Questionnaire.DerivedFrom."+strconv.Itoa(numDerivedFrom)+".", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.DerivedFrom."+strconv.Itoa(numDerivedFrom)+".", &resource.DerivedFrom[numDerivedFrom], htmlAttrs)
}
func (resource *Questionnaire) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Questionnaire.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Questionnaire.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_Experimental(htmlAttrs string) templ.Component {

	if resource == nil {
		return BoolInput("Questionnaire.Experimental", nil, htmlAttrs)
	}
	return BoolInput("Questionnaire.Experimental", resource.Experimental, htmlAttrs)
}
func (resource *Questionnaire) T_SubjectType(numSubjectType int, htmlAttrs string) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || numSubjectType >= len(resource.SubjectType) {
		return CodeSelect("Questionnaire.SubjectType."+strconv.Itoa(numSubjectType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Questionnaire.SubjectType."+strconv.Itoa(numSubjectType)+".", &resource.SubjectType[numSubjectType], optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_Date(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Questionnaire.Date", nil, htmlAttrs)
	}
	return DateTimeInput("Questionnaire.Date", resource.Date, htmlAttrs)
}
func (resource *Questionnaire) T_Publisher(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Publisher", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Publisher", resource.Publisher, htmlAttrs)
}
func (resource *Questionnaire) T_Description(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Description", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Description", resource.Description, htmlAttrs)
}
func (resource *Questionnaire) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numJurisdiction >= len(resource.Jurisdiction) {
		return CodeableConceptSelect("Questionnaire.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Questionnaire.Jurisdiction."+strconv.Itoa(numJurisdiction)+".", &resource.Jurisdiction[numJurisdiction], optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_Purpose(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Purpose", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Purpose", resource.Purpose, htmlAttrs)
}
func (resource *Questionnaire) T_Copyright(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Copyright", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Copyright", resource.Copyright, htmlAttrs)
}
func (resource *Questionnaire) T_ApprovalDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("Questionnaire.ApprovalDate", nil, htmlAttrs)
	}
	return DateInput("Questionnaire.ApprovalDate", resource.ApprovalDate, htmlAttrs)
}
func (resource *Questionnaire) T_LastReviewDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("Questionnaire.LastReviewDate", nil, htmlAttrs)
	}
	return DateInput("Questionnaire.LastReviewDate", resource.LastReviewDate, htmlAttrs)
}
func (resource *Questionnaire) T_Code(numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCode >= len(resource.Code) {
		return CodingSelect("Questionnaire.Code."+strconv.Itoa(numCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Questionnaire.Code."+strconv.Itoa(numCode)+".", &resource.Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_ItemLinkId(numItem int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..LinkId", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..LinkId", &resource.Item[numItem].LinkId, htmlAttrs)
}
func (resource *Questionnaire) T_ItemDefinition(numItem int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Definition", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Definition", resource.Item[numItem].Definition, htmlAttrs)
}
func (resource *Questionnaire) T_ItemCode(numItem int, numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numCode >= len(resource.Item[numItem].Code) {
		return CodingSelect("Questionnaire.Item."+strconv.Itoa(numItem)+"..Code."+strconv.Itoa(numCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Questionnaire.Item."+strconv.Itoa(numItem)+"..Code."+strconv.Itoa(numCode)+".", &resource.Item[numItem].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_ItemPrefix(numItem int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Prefix", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Prefix", resource.Item[numItem].Prefix, htmlAttrs)
}
func (resource *Questionnaire) T_ItemText(numItem int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Text", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Text", resource.Item[numItem].Text, htmlAttrs)
}
func (resource *Questionnaire) T_ItemType(numItem int, htmlAttrs string) templ.Component {
	optionsValueSet := VSItem_type

	if resource == nil || numItem >= len(resource.Item) {
		return CodeSelect("Questionnaire.Item."+strconv.Itoa(numItem)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Questionnaire.Item."+strconv.Itoa(numItem)+"..Type", &resource.Item[numItem].Type, optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableBehavior(numItem int, htmlAttrs string) templ.Component {
	optionsValueSet := VSQuestionnaire_enable_behavior

	if resource == nil || numItem >= len(resource.Item) {
		return CodeSelect("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableBehavior", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableBehavior", resource.Item[numItem].EnableBehavior, optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_ItemRequired(numItem int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) {
		return BoolInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Required", nil, htmlAttrs)
	}
	return BoolInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Required", resource.Item[numItem].Required, htmlAttrs)
}
func (resource *Questionnaire) T_ItemRepeats(numItem int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) {
		return BoolInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Repeats", nil, htmlAttrs)
	}
	return BoolInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Repeats", resource.Item[numItem].Repeats, htmlAttrs)
}
func (resource *Questionnaire) T_ItemReadOnly(numItem int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) {
		return BoolInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..ReadOnly", nil, htmlAttrs)
	}
	return BoolInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..ReadOnly", resource.Item[numItem].ReadOnly, htmlAttrs)
}
func (resource *Questionnaire) T_ItemMaxLength(numItem int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) {
		return IntInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..MaxLength", nil, htmlAttrs)
	}
	return IntInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..MaxLength", resource.Item[numItem].MaxLength, htmlAttrs)
}
func (resource *Questionnaire) T_ItemAnswerValueSet(numItem int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..AnswerValueSet", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..AnswerValueSet", resource.Item[numItem].AnswerValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenQuestion(numItem int, numEnableWhen int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..Question", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..Question", &resource.Item[numItem].EnableWhen[numEnableWhen].Question, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenOperator(numItem int, numEnableWhen int, htmlAttrs string) templ.Component {
	optionsValueSet := VSQuestionnaire_enable_operator

	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return CodeSelect("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..Operator", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..Operator", &resource.Item[numItem].EnableWhen[numEnableWhen].Operator, optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenAnswerBoolean(numItem int, numEnableWhen int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return BoolInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..AnswerBoolean", nil, htmlAttrs)
	}
	return BoolInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..AnswerBoolean", &resource.Item[numItem].EnableWhen[numEnableWhen].AnswerBoolean, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenAnswerDecimal(numItem int, numEnableWhen int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return Float64Input("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..AnswerDecimal", nil, htmlAttrs)
	}
	return Float64Input("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..AnswerDecimal", &resource.Item[numItem].EnableWhen[numEnableWhen].AnswerDecimal, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenAnswerInteger(numItem int, numEnableWhen int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return IntInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..AnswerInteger", nil, htmlAttrs)
	}
	return IntInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..AnswerInteger", &resource.Item[numItem].EnableWhen[numEnableWhen].AnswerInteger, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenAnswerDate(numItem int, numEnableWhen int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return DateInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..AnswerDate", nil, htmlAttrs)
	}
	return DateInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..AnswerDate", &resource.Item[numItem].EnableWhen[numEnableWhen].AnswerDate, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenAnswerDateTime(numItem int, numEnableWhen int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return DateTimeInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..AnswerDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..AnswerDateTime", &resource.Item[numItem].EnableWhen[numEnableWhen].AnswerDateTime, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenAnswerTime(numItem int, numEnableWhen int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..AnswerTime", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..AnswerTime", &resource.Item[numItem].EnableWhen[numEnableWhen].AnswerTime, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenAnswerString(numItem int, numEnableWhen int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..AnswerString", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..AnswerString", &resource.Item[numItem].EnableWhen[numEnableWhen].AnswerString, htmlAttrs)
}
func (resource *Questionnaire) T_ItemEnableWhenAnswerCoding(numItem int, numEnableWhen int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numEnableWhen >= len(resource.Item[numItem].EnableWhen) {
		return CodingSelect("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..AnswerCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Questionnaire.Item."+strconv.Itoa(numItem)+"..EnableWhen."+strconv.Itoa(numEnableWhen)+"..AnswerCoding", &resource.Item[numItem].EnableWhen[numEnableWhen].AnswerCoding, optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_ItemAnswerOptionValueInteger(numItem int, numAnswerOption int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numAnswerOption >= len(resource.Item[numItem].AnswerOption) {
		return IntInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..AnswerOption."+strconv.Itoa(numAnswerOption)+"..ValueInteger", nil, htmlAttrs)
	}
	return IntInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..AnswerOption."+strconv.Itoa(numAnswerOption)+"..ValueInteger", &resource.Item[numItem].AnswerOption[numAnswerOption].ValueInteger, htmlAttrs)
}
func (resource *Questionnaire) T_ItemAnswerOptionValueDate(numItem int, numAnswerOption int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numAnswerOption >= len(resource.Item[numItem].AnswerOption) {
		return DateInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..AnswerOption."+strconv.Itoa(numAnswerOption)+"..ValueDate", nil, htmlAttrs)
	}
	return DateInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..AnswerOption."+strconv.Itoa(numAnswerOption)+"..ValueDate", &resource.Item[numItem].AnswerOption[numAnswerOption].ValueDate, htmlAttrs)
}
func (resource *Questionnaire) T_ItemAnswerOptionValueTime(numItem int, numAnswerOption int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numAnswerOption >= len(resource.Item[numItem].AnswerOption) {
		return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..AnswerOption."+strconv.Itoa(numAnswerOption)+"..ValueTime", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..AnswerOption."+strconv.Itoa(numAnswerOption)+"..ValueTime", &resource.Item[numItem].AnswerOption[numAnswerOption].ValueTime, htmlAttrs)
}
func (resource *Questionnaire) T_ItemAnswerOptionValueString(numItem int, numAnswerOption int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numAnswerOption >= len(resource.Item[numItem].AnswerOption) {
		return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..AnswerOption."+strconv.Itoa(numAnswerOption)+"..ValueString", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..AnswerOption."+strconv.Itoa(numAnswerOption)+"..ValueString", &resource.Item[numItem].AnswerOption[numAnswerOption].ValueString, htmlAttrs)
}
func (resource *Questionnaire) T_ItemAnswerOptionValueCoding(numItem int, numAnswerOption int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numAnswerOption >= len(resource.Item[numItem].AnswerOption) {
		return CodingSelect("Questionnaire.Item."+strconv.Itoa(numItem)+"..AnswerOption."+strconv.Itoa(numAnswerOption)+"..ValueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Questionnaire.Item."+strconv.Itoa(numItem)+"..AnswerOption."+strconv.Itoa(numAnswerOption)+"..ValueCoding", &resource.Item[numItem].AnswerOption[numAnswerOption].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *Questionnaire) T_ItemAnswerOptionInitialSelected(numItem int, numAnswerOption int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numAnswerOption >= len(resource.Item[numItem].AnswerOption) {
		return BoolInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..AnswerOption."+strconv.Itoa(numAnswerOption)+"..InitialSelected", nil, htmlAttrs)
	}
	return BoolInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..AnswerOption."+strconv.Itoa(numAnswerOption)+"..InitialSelected", resource.Item[numItem].AnswerOption[numAnswerOption].InitialSelected, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueBoolean(numItem int, numInitial int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return BoolInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueBoolean", &resource.Item[numItem].Initial[numInitial].ValueBoolean, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueDecimal(numItem int, numInitial int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return Float64Input("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueDecimal", nil, htmlAttrs)
	}
	return Float64Input("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueDecimal", &resource.Item[numItem].Initial[numInitial].ValueDecimal, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueInteger(numItem int, numInitial int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return IntInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueInteger", nil, htmlAttrs)
	}
	return IntInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueInteger", &resource.Item[numItem].Initial[numInitial].ValueInteger, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueDate(numItem int, numInitial int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return DateInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueDate", nil, htmlAttrs)
	}
	return DateInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueDate", &resource.Item[numItem].Initial[numInitial].ValueDate, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueDateTime(numItem int, numInitial int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return DateTimeInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueDateTime", &resource.Item[numItem].Initial[numInitial].ValueDateTime, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueTime(numItem int, numInitial int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueTime", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueTime", &resource.Item[numItem].Initial[numInitial].ValueTime, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueString(numItem int, numInitial int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueString", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueString", &resource.Item[numItem].Initial[numInitial].ValueString, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueUri(numItem int, numInitial int, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueUri", nil, htmlAttrs)
	}
	return StringInput("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueUri", &resource.Item[numItem].Initial[numInitial].ValueUri, htmlAttrs)
}
func (resource *Questionnaire) T_ItemInitialValueCoding(numItem int, numInitial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numItem >= len(resource.Item) || numInitial >= len(resource.Item[numItem].Initial) {
		return CodingSelect("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Questionnaire.Item."+strconv.Itoa(numItem)+"..Initial."+strconv.Itoa(numInitial)+"..ValueCoding", &resource.Item[numItem].Initial[numInitial].ValueCoding, optionsValueSet, htmlAttrs)
}
