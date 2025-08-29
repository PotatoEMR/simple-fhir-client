package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/Questionnaire
type Questionnaire struct {
	Id                     *string             `json:"id,omitempty"`
	Meta                   *Meta               `json:"meta,omitempty"`
	ImplicitRules          *string             `json:"implicitRules,omitempty"`
	Language               *string             `json:"language,omitempty"`
	Text                   *Narrative          `json:"text,omitempty"`
	Contained              []Resource          `json:"contained,omitempty"`
	Extension              []Extension         `json:"extension,omitempty"`
	ModifierExtension      []Extension         `json:"modifierExtension,omitempty"`
	Url                    *string             `json:"url,omitempty"`
	Identifier             []Identifier        `json:"identifier,omitempty"`
	Version                *string             `json:"version,omitempty"`
	VersionAlgorithmString *string             `json:"versionAlgorithmString"`
	VersionAlgorithmCoding *Coding             `json:"versionAlgorithmCoding"`
	Name                   *string             `json:"name,omitempty"`
	Title                  *string             `json:"title,omitempty"`
	DerivedFrom            []string            `json:"derivedFrom,omitempty"`
	Status                 string              `json:"status"`
	Experimental           *bool               `json:"experimental,omitempty"`
	SubjectType            []string            `json:"subjectType,omitempty"`
	Date                   *string             `json:"date,omitempty"`
	Publisher              *string             `json:"publisher,omitempty"`
	Contact                []ContactDetail     `json:"contact,omitempty"`
	Description            *string             `json:"description,omitempty"`
	UseContext             []UsageContext      `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept   `json:"jurisdiction,omitempty"`
	Purpose                *string             `json:"purpose,omitempty"`
	Copyright              *string             `json:"copyright,omitempty"`
	CopyrightLabel         *string             `json:"copyrightLabel,omitempty"`
	ApprovalDate           *string             `json:"approvalDate,omitempty"`
	LastReviewDate         *string             `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period             `json:"effectivePeriod,omitempty"`
	Code                   []Coding            `json:"code,omitempty"`
	Item                   []QuestionnaireItem `json:"item,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Questionnaire
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
	DisabledDisplay   *string                         `json:"disabledDisplay,omitempty"`
	Required          *bool                           `json:"required,omitempty"`
	Repeats           *bool                           `json:"repeats,omitempty"`
	ReadOnly          *bool                           `json:"readOnly,omitempty"`
	MaxLength         *int                            `json:"maxLength,omitempty"`
	AnswerConstraint  *string                         `json:"answerConstraint,omitempty"`
	AnswerValueSet    *string                         `json:"answerValueSet,omitempty"`
	AnswerOption      []QuestionnaireItemAnswerOption `json:"answerOption,omitempty"`
	Initial           []QuestionnaireItemInitial      `json:"initial,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Questionnaire
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

// http://hl7.org/fhir/r5/StructureDefinition/Questionnaire
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

// http://hl7.org/fhir/r5/StructureDefinition/Questionnaire
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
func (resource *Questionnaire) QuestionnaireLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Questionnaire) QuestionnaireStatus() templ.Component {
	optionsValueSet := VSPublication_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *Questionnaire) QuestionnaireSubjectType() templ.Component {
	optionsValueSet := VSResource_types
	currentVal := ""
	if resource != nil {
		currentVal = resource.SubjectType[0]
	}
	return CodeSelect("subjectType", currentVal, optionsValueSet)
}
func (resource *Questionnaire) QuestionnaireItemType(numItem int) templ.Component {
	optionsValueSet := VSItem_type
	currentVal := ""
	if resource != nil && len(resource.Item) >= numItem {
		currentVal = resource.Item[numItem].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
func (resource *Questionnaire) QuestionnaireItemEnableBehavior(numItem int) templ.Component {
	optionsValueSet := VSQuestionnaire_enable_behavior
	currentVal := ""
	if resource != nil && len(resource.Item) >= numItem {
		currentVal = *resource.Item[numItem].EnableBehavior
	}
	return CodeSelect("enableBehavior", currentVal, optionsValueSet)
}
func (resource *Questionnaire) QuestionnaireItemDisabledDisplay(numItem int) templ.Component {
	optionsValueSet := VSQuestionnaire_disabled_display
	currentVal := ""
	if resource != nil && len(resource.Item) >= numItem {
		currentVal = *resource.Item[numItem].DisabledDisplay
	}
	return CodeSelect("disabledDisplay", currentVal, optionsValueSet)
}
func (resource *Questionnaire) QuestionnaireItemAnswerConstraint(numItem int) templ.Component {
	optionsValueSet := VSQuestionnaire_answer_constraint
	currentVal := ""
	if resource != nil && len(resource.Item) >= numItem {
		currentVal = *resource.Item[numItem].AnswerConstraint
	}
	return CodeSelect("answerConstraint", currentVal, optionsValueSet)
}
func (resource *Questionnaire) QuestionnaireItemEnableWhenOperator(numItem int, numEnableWhen int) templ.Component {
	optionsValueSet := VSQuestionnaire_enable_operator
	currentVal := ""
	if resource != nil && len(resource.Item[numItem].EnableWhen) >= numEnableWhen {
		currentVal = resource.Item[numItem].EnableWhen[numEnableWhen].Operator
	}
	return CodeSelect("operator", currentVal, optionsValueSet)
}
