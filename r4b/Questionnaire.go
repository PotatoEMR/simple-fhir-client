package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Questionnaire
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

// http://hl7.org/fhir/r4b/StructureDefinition/Questionnaire
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

// http://hl7.org/fhir/r4b/StructureDefinition/Questionnaire
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

// http://hl7.org/fhir/r4b/StructureDefinition/Questionnaire
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

// http://hl7.org/fhir/r4b/StructureDefinition/Questionnaire
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

func (resource *Questionnaire) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Id", nil)
	}
	return StringInput("Questionnaire.Id", resource.Id)
}
func (resource *Questionnaire) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.ImplicitRules", nil)
	}
	return StringInput("Questionnaire.ImplicitRules", resource.ImplicitRules)
}
func (resource *Questionnaire) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Questionnaire.Language", nil, optionsValueSet)
	}
	return CodeSelect("Questionnaire.Language", resource.Language, optionsValueSet)
}
func (resource *Questionnaire) T_Url() templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Url", nil)
	}
	return StringInput("Questionnaire.Url", resource.Url)
}
func (resource *Questionnaire) T_Version() templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Version", nil)
	}
	return StringInput("Questionnaire.Version", resource.Version)
}
func (resource *Questionnaire) T_Name() templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Name", nil)
	}
	return StringInput("Questionnaire.Name", resource.Name)
}
func (resource *Questionnaire) T_Title() templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Title", nil)
	}
	return StringInput("Questionnaire.Title", resource.Title)
}
func (resource *Questionnaire) T_DerivedFrom(numDerivedFrom int) templ.Component {

	if resource == nil || len(resource.DerivedFrom) >= numDerivedFrom {
		return StringInput("Questionnaire.DerivedFrom["+strconv.Itoa(numDerivedFrom)+"]", nil)
	}
	return StringInput("Questionnaire.DerivedFrom["+strconv.Itoa(numDerivedFrom)+"]", &resource.DerivedFrom[numDerivedFrom])
}
func (resource *Questionnaire) T_Status() templ.Component {
	optionsValueSet := VSPublication_status

	if resource == nil {
		return CodeSelect("Questionnaire.Status", nil, optionsValueSet)
	}
	return CodeSelect("Questionnaire.Status", &resource.Status, optionsValueSet)
}
func (resource *Questionnaire) T_Experimental() templ.Component {

	if resource == nil {
		return BoolInput("Questionnaire.Experimental", nil)
	}
	return BoolInput("Questionnaire.Experimental", resource.Experimental)
}
func (resource *Questionnaire) T_SubjectType(numSubjectType int) templ.Component {
	optionsValueSet := VSResource_types

	if resource == nil || len(resource.SubjectType) >= numSubjectType {
		return CodeSelect("Questionnaire.SubjectType["+strconv.Itoa(numSubjectType)+"]", nil, optionsValueSet)
	}
	return CodeSelect("Questionnaire.SubjectType["+strconv.Itoa(numSubjectType)+"]", &resource.SubjectType[numSubjectType], optionsValueSet)
}
func (resource *Questionnaire) T_Date() templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Date", nil)
	}
	return StringInput("Questionnaire.Date", resource.Date)
}
func (resource *Questionnaire) T_Publisher() templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Publisher", nil)
	}
	return StringInput("Questionnaire.Publisher", resource.Publisher)
}
func (resource *Questionnaire) T_Description() templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Description", nil)
	}
	return StringInput("Questionnaire.Description", resource.Description)
}
func (resource *Questionnaire) T_Jurisdiction(numJurisdiction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Jurisdiction) >= numJurisdiction {
		return CodeableConceptSelect("Questionnaire.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Questionnaire.Jurisdiction["+strconv.Itoa(numJurisdiction)+"]", &resource.Jurisdiction[numJurisdiction], optionsValueSet)
}
func (resource *Questionnaire) T_Purpose() templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Purpose", nil)
	}
	return StringInput("Questionnaire.Purpose", resource.Purpose)
}
func (resource *Questionnaire) T_Copyright() templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.Copyright", nil)
	}
	return StringInput("Questionnaire.Copyright", resource.Copyright)
}
func (resource *Questionnaire) T_ApprovalDate() templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.ApprovalDate", nil)
	}
	return StringInput("Questionnaire.ApprovalDate", resource.ApprovalDate)
}
func (resource *Questionnaire) T_LastReviewDate() templ.Component {

	if resource == nil {
		return StringInput("Questionnaire.LastReviewDate", nil)
	}
	return StringInput("Questionnaire.LastReviewDate", resource.LastReviewDate)
}
func (resource *Questionnaire) T_Code(numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Code) >= numCode {
		return CodingSelect("Questionnaire.Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet)
	}
	return CodingSelect("Questionnaire.Code["+strconv.Itoa(numCode)+"]", &resource.Code[numCode], optionsValueSet)
}
func (resource *Questionnaire) T_ItemId(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].Id", nil)
	}
	return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].Id", resource.Item[numItem].Id)
}
func (resource *Questionnaire) T_ItemLinkId(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].LinkId", nil)
	}
	return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].LinkId", &resource.Item[numItem].LinkId)
}
func (resource *Questionnaire) T_ItemDefinition(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].Definition", nil)
	}
	return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].Definition", resource.Item[numItem].Definition)
}
func (resource *Questionnaire) T_ItemCode(numItem int, numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Code) >= numCode {
		return CodingSelect("Questionnaire.Item["+strconv.Itoa(numItem)+"].Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet)
	}
	return CodingSelect("Questionnaire.Item["+strconv.Itoa(numItem)+"].Code["+strconv.Itoa(numCode)+"]", &resource.Item[numItem].Code[numCode], optionsValueSet)
}
func (resource *Questionnaire) T_ItemPrefix(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].Prefix", nil)
	}
	return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].Prefix", resource.Item[numItem].Prefix)
}
func (resource *Questionnaire) T_ItemText(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].Text", nil)
	}
	return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].Text", resource.Item[numItem].Text)
}
func (resource *Questionnaire) T_ItemType(numItem int) templ.Component {
	optionsValueSet := VSItem_type

	if resource == nil || len(resource.Item) >= numItem {
		return CodeSelect("Questionnaire.Item["+strconv.Itoa(numItem)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("Questionnaire.Item["+strconv.Itoa(numItem)+"].Type", &resource.Item[numItem].Type, optionsValueSet)
}
func (resource *Questionnaire) T_ItemEnableBehavior(numItem int) templ.Component {
	optionsValueSet := VSQuestionnaire_enable_behavior

	if resource == nil || len(resource.Item) >= numItem {
		return CodeSelect("Questionnaire.Item["+strconv.Itoa(numItem)+"].EnableBehavior", nil, optionsValueSet)
	}
	return CodeSelect("Questionnaire.Item["+strconv.Itoa(numItem)+"].EnableBehavior", resource.Item[numItem].EnableBehavior, optionsValueSet)
}
func (resource *Questionnaire) T_ItemRequired(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return BoolInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].Required", nil)
	}
	return BoolInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].Required", resource.Item[numItem].Required)
}
func (resource *Questionnaire) T_ItemRepeats(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return BoolInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].Repeats", nil)
	}
	return BoolInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].Repeats", resource.Item[numItem].Repeats)
}
func (resource *Questionnaire) T_ItemReadOnly(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return BoolInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].ReadOnly", nil)
	}
	return BoolInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].ReadOnly", resource.Item[numItem].ReadOnly)
}
func (resource *Questionnaire) T_ItemMaxLength(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return IntInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].MaxLength", nil)
	}
	return IntInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].MaxLength", resource.Item[numItem].MaxLength)
}
func (resource *Questionnaire) T_ItemAnswerValueSet(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].AnswerValueSet", nil)
	}
	return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].AnswerValueSet", resource.Item[numItem].AnswerValueSet)
}
func (resource *Questionnaire) T_ItemEnableWhenId(numItem int, numEnableWhen int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].EnableWhen) >= numEnableWhen {
		return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].EnableWhen["+strconv.Itoa(numEnableWhen)+"].Id", nil)
	}
	return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].EnableWhen["+strconv.Itoa(numEnableWhen)+"].Id", resource.Item[numItem].EnableWhen[numEnableWhen].Id)
}
func (resource *Questionnaire) T_ItemEnableWhenQuestion(numItem int, numEnableWhen int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].EnableWhen) >= numEnableWhen {
		return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].EnableWhen["+strconv.Itoa(numEnableWhen)+"].Question", nil)
	}
	return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].EnableWhen["+strconv.Itoa(numEnableWhen)+"].Question", &resource.Item[numItem].EnableWhen[numEnableWhen].Question)
}
func (resource *Questionnaire) T_ItemEnableWhenOperator(numItem int, numEnableWhen int) templ.Component {
	optionsValueSet := VSQuestionnaire_enable_operator

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].EnableWhen) >= numEnableWhen {
		return CodeSelect("Questionnaire.Item["+strconv.Itoa(numItem)+"].EnableWhen["+strconv.Itoa(numEnableWhen)+"].Operator", nil, optionsValueSet)
	}
	return CodeSelect("Questionnaire.Item["+strconv.Itoa(numItem)+"].EnableWhen["+strconv.Itoa(numEnableWhen)+"].Operator", &resource.Item[numItem].EnableWhen[numEnableWhen].Operator, optionsValueSet)
}
func (resource *Questionnaire) T_ItemAnswerOptionId(numItem int, numAnswerOption int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].AnswerOption) >= numAnswerOption {
		return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].AnswerOption["+strconv.Itoa(numAnswerOption)+"].Id", nil)
	}
	return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].AnswerOption["+strconv.Itoa(numAnswerOption)+"].Id", resource.Item[numItem].AnswerOption[numAnswerOption].Id)
}
func (resource *Questionnaire) T_ItemAnswerOptionInitialSelected(numItem int, numAnswerOption int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].AnswerOption) >= numAnswerOption {
		return BoolInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].AnswerOption["+strconv.Itoa(numAnswerOption)+"].InitialSelected", nil)
	}
	return BoolInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].AnswerOption["+strconv.Itoa(numAnswerOption)+"].InitialSelected", resource.Item[numItem].AnswerOption[numAnswerOption].InitialSelected)
}
func (resource *Questionnaire) T_ItemInitialId(numItem int, numInitial int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Initial) >= numInitial {
		return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].Initial["+strconv.Itoa(numInitial)+"].Id", nil)
	}
	return StringInput("Questionnaire.Item["+strconv.Itoa(numItem)+"].Initial["+strconv.Itoa(numInitial)+"].Id", resource.Item[numItem].Initial[numInitial].Id)
}
