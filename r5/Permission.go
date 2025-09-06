package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Permission
type Permission struct {
	Id                *string                  `json:"id,omitempty"`
	Meta              *Meta                    `json:"meta,omitempty"`
	ImplicitRules     *string                  `json:"implicitRules,omitempty"`
	Language          *string                  `json:"language,omitempty"`
	Text              *Narrative               `json:"text,omitempty"`
	Contained         []Resource               `json:"contained,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Status            string                   `json:"status"`
	Asserter          *Reference               `json:"asserter,omitempty"`
	Date              []string                 `json:"date,omitempty"`
	Validity          *Period                  `json:"validity,omitempty"`
	Justification     *PermissionJustification `json:"justification,omitempty"`
	Combining         string                   `json:"combining"`
	Rule              []PermissionRule         `json:"rule,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Permission
type PermissionJustification struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Basis             []CodeableConcept `json:"basis,omitempty"`
	Evidence          []Reference       `json:"evidence,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Permission
type PermissionRule struct {
	Id                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Type              *string                  `json:"type,omitempty"`
	Data              []PermissionRuleData     `json:"data,omitempty"`
	Activity          []PermissionRuleActivity `json:"activity,omitempty"`
	Limit             []CodeableConcept        `json:"limit,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Permission
type PermissionRuleData struct {
	Id                *string                      `json:"id,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Resource          []PermissionRuleDataResource `json:"resource,omitempty"`
	Security          []Coding                     `json:"security,omitempty"`
	Period            []Period                     `json:"period,omitempty"`
	Expression        *Expression                  `json:"expression,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Permission
type PermissionRuleDataResource struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Meaning           string      `json:"meaning"`
	Reference         Reference   `json:"reference"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Permission
type PermissionRuleActivity struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Actor             []Reference       `json:"actor,omitempty"`
	Action            []CodeableConcept `json:"action,omitempty"`
	Purpose           []CodeableConcept `json:"purpose,omitempty"`
}

type OtherPermission Permission

// on convert struct to json, automatically add resourceType=Permission
func (r Permission) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPermission
		ResourceType string `json:"resourceType"`
	}{
		OtherPermission: OtherPermission(r),
		ResourceType:    "Permission",
	})
}

func (resource *Permission) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Permission.Id", nil)
	}
	return StringInput("Permission.Id", resource.Id)
}
func (resource *Permission) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Permission.ImplicitRules", nil)
	}
	return StringInput("Permission.ImplicitRules", resource.ImplicitRules)
}
func (resource *Permission) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Permission.Language", nil, optionsValueSet)
	}
	return CodeSelect("Permission.Language", resource.Language, optionsValueSet)
}
func (resource *Permission) T_Status() templ.Component {
	optionsValueSet := VSPermission_status

	if resource == nil {
		return CodeSelect("Permission.Status", nil, optionsValueSet)
	}
	return CodeSelect("Permission.Status", &resource.Status, optionsValueSet)
}
func (resource *Permission) T_Date(numDate int) templ.Component {

	if resource == nil || len(resource.Date) >= numDate {
		return StringInput("Permission.Date["+strconv.Itoa(numDate)+"]", nil)
	}
	return StringInput("Permission.Date["+strconv.Itoa(numDate)+"]", &resource.Date[numDate])
}
func (resource *Permission) T_Combining() templ.Component {
	optionsValueSet := VSPermission_rule_combining

	if resource == nil {
		return CodeSelect("Permission.Combining", nil, optionsValueSet)
	}
	return CodeSelect("Permission.Combining", &resource.Combining, optionsValueSet)
}
func (resource *Permission) T_JustificationId() templ.Component {

	if resource == nil {
		return StringInput("Permission.Justification.Id", nil)
	}
	return StringInput("Permission.Justification.Id", resource.Justification.Id)
}
func (resource *Permission) T_JustificationBasis(numBasis int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Justification.Basis) >= numBasis {
		return CodeableConceptSelect("Permission.Justification.Basis["+strconv.Itoa(numBasis)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Permission.Justification.Basis["+strconv.Itoa(numBasis)+"]", &resource.Justification.Basis[numBasis], optionsValueSet)
}
func (resource *Permission) T_RuleId(numRule int) templ.Component {

	if resource == nil || len(resource.Rule) >= numRule {
		return StringInput("Permission.Rule["+strconv.Itoa(numRule)+"].Id", nil)
	}
	return StringInput("Permission.Rule["+strconv.Itoa(numRule)+"].Id", resource.Rule[numRule].Id)
}
func (resource *Permission) T_RuleType(numRule int) templ.Component {
	optionsValueSet := VSConsent_provision_type

	if resource == nil || len(resource.Rule) >= numRule {
		return CodeSelect("Permission.Rule["+strconv.Itoa(numRule)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("Permission.Rule["+strconv.Itoa(numRule)+"].Type", resource.Rule[numRule].Type, optionsValueSet)
}
func (resource *Permission) T_RuleLimit(numRule int, numLimit int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Rule) >= numRule || len(resource.Rule[numRule].Limit) >= numLimit {
		return CodeableConceptSelect("Permission.Rule["+strconv.Itoa(numRule)+"].Limit["+strconv.Itoa(numLimit)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Permission.Rule["+strconv.Itoa(numRule)+"].Limit["+strconv.Itoa(numLimit)+"]", &resource.Rule[numRule].Limit[numLimit], optionsValueSet)
}
func (resource *Permission) T_RuleDataId(numRule int, numData int) templ.Component {

	if resource == nil || len(resource.Rule) >= numRule || len(resource.Rule[numRule].Data) >= numData {
		return StringInput("Permission.Rule["+strconv.Itoa(numRule)+"].Data["+strconv.Itoa(numData)+"].Id", nil)
	}
	return StringInput("Permission.Rule["+strconv.Itoa(numRule)+"].Data["+strconv.Itoa(numData)+"].Id", resource.Rule[numRule].Data[numData].Id)
}
func (resource *Permission) T_RuleDataSecurity(numRule int, numData int, numSecurity int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Rule) >= numRule || len(resource.Rule[numRule].Data) >= numData || len(resource.Rule[numRule].Data[numData].Security) >= numSecurity {
		return CodingSelect("Permission.Rule["+strconv.Itoa(numRule)+"].Data["+strconv.Itoa(numData)+"].Security["+strconv.Itoa(numSecurity)+"]", nil, optionsValueSet)
	}
	return CodingSelect("Permission.Rule["+strconv.Itoa(numRule)+"].Data["+strconv.Itoa(numData)+"].Security["+strconv.Itoa(numSecurity)+"]", &resource.Rule[numRule].Data[numData].Security[numSecurity], optionsValueSet)
}
func (resource *Permission) T_RuleDataResourceId(numRule int, numData int, numResource int) templ.Component {

	if resource == nil || len(resource.Rule) >= numRule || len(resource.Rule[numRule].Data) >= numData || len(resource.Rule[numRule].Data[numData].Resource) >= numResource {
		return StringInput("Permission.Rule["+strconv.Itoa(numRule)+"].Data["+strconv.Itoa(numData)+"].Resource["+strconv.Itoa(numResource)+"].Id", nil)
	}
	return StringInput("Permission.Rule["+strconv.Itoa(numRule)+"].Data["+strconv.Itoa(numData)+"].Resource["+strconv.Itoa(numResource)+"].Id", resource.Rule[numRule].Data[numData].Resource[numResource].Id)
}
func (resource *Permission) T_RuleDataResourceMeaning(numRule int, numData int, numResource int) templ.Component {
	optionsValueSet := VSConsent_data_meaning

	if resource == nil || len(resource.Rule) >= numRule || len(resource.Rule[numRule].Data) >= numData || len(resource.Rule[numRule].Data[numData].Resource) >= numResource {
		return CodeSelect("Permission.Rule["+strconv.Itoa(numRule)+"].Data["+strconv.Itoa(numData)+"].Resource["+strconv.Itoa(numResource)+"].Meaning", nil, optionsValueSet)
	}
	return CodeSelect("Permission.Rule["+strconv.Itoa(numRule)+"].Data["+strconv.Itoa(numData)+"].Resource["+strconv.Itoa(numResource)+"].Meaning", &resource.Rule[numRule].Data[numData].Resource[numResource].Meaning, optionsValueSet)
}
func (resource *Permission) T_RuleActivityId(numRule int, numActivity int) templ.Component {

	if resource == nil || len(resource.Rule) >= numRule || len(resource.Rule[numRule].Activity) >= numActivity {
		return StringInput("Permission.Rule["+strconv.Itoa(numRule)+"].Activity["+strconv.Itoa(numActivity)+"].Id", nil)
	}
	return StringInput("Permission.Rule["+strconv.Itoa(numRule)+"].Activity["+strconv.Itoa(numActivity)+"].Id", resource.Rule[numRule].Activity[numActivity].Id)
}
func (resource *Permission) T_RuleActivityAction(numRule int, numActivity int, numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Rule) >= numRule || len(resource.Rule[numRule].Activity) >= numActivity || len(resource.Rule[numRule].Activity[numActivity].Action) >= numAction {
		return CodeableConceptSelect("Permission.Rule["+strconv.Itoa(numRule)+"].Activity["+strconv.Itoa(numActivity)+"].Action["+strconv.Itoa(numAction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Permission.Rule["+strconv.Itoa(numRule)+"].Activity["+strconv.Itoa(numActivity)+"].Action["+strconv.Itoa(numAction)+"]", &resource.Rule[numRule].Activity[numActivity].Action[numAction], optionsValueSet)
}
func (resource *Permission) T_RuleActivityPurpose(numRule int, numActivity int, numPurpose int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Rule) >= numRule || len(resource.Rule[numRule].Activity) >= numActivity || len(resource.Rule[numRule].Activity[numActivity].Purpose) >= numPurpose {
		return CodeableConceptSelect("Permission.Rule["+strconv.Itoa(numRule)+"].Activity["+strconv.Itoa(numActivity)+"].Purpose["+strconv.Itoa(numPurpose)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Permission.Rule["+strconv.Itoa(numRule)+"].Activity["+strconv.Itoa(numActivity)+"].Purpose["+strconv.Itoa(numPurpose)+"]", &resource.Rule[numRule].Activity[numActivity].Purpose[numPurpose], optionsValueSet)
}
