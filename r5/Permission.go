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
	Date              []time.Time              `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r Permission) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Permission/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "Permission"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Permission) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSPermission_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Permission) T_Date(numDate int, htmlAttrs string) templ.Component {
	if resource == nil || numDate >= len(resource.Date) {
		return DateTimeInput("Date["+strconv.Itoa(numDate)+"]", nil, htmlAttrs)
	}
	return DateTimeInput("Date["+strconv.Itoa(numDate)+"]", &resource.Date[numDate], htmlAttrs)
}
func (resource *Permission) T_Combining(htmlAttrs string) templ.Component {
	optionsValueSet := VSPermission_rule_combining

	if resource == nil {
		return CodeSelect("Combining", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Combining", &resource.Combining, optionsValueSet, htmlAttrs)
}
func (resource *Permission) T_JustificationBasis(numBasis int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBasis >= len(resource.Justification.Basis) {
		return CodeableConceptSelect("JustificationBasis["+strconv.Itoa(numBasis)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("JustificationBasis["+strconv.Itoa(numBasis)+"]", &resource.Justification.Basis[numBasis], optionsValueSet, htmlAttrs)
}
func (resource *Permission) T_RuleType(numRule int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConsent_provision_type

	if resource == nil || numRule >= len(resource.Rule) {
		return CodeSelect("Rule["+strconv.Itoa(numRule)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Rule["+strconv.Itoa(numRule)+"]Type", resource.Rule[numRule].Type, optionsValueSet, htmlAttrs)
}
func (resource *Permission) T_RuleLimit(numRule int, numLimit int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRule >= len(resource.Rule) || numLimit >= len(resource.Rule[numRule].Limit) {
		return CodeableConceptSelect("Rule["+strconv.Itoa(numRule)+"]Limit["+strconv.Itoa(numLimit)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Rule["+strconv.Itoa(numRule)+"]Limit["+strconv.Itoa(numLimit)+"]", &resource.Rule[numRule].Limit[numLimit], optionsValueSet, htmlAttrs)
}
func (resource *Permission) T_RuleDataSecurity(numRule int, numData int, numSecurity int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRule >= len(resource.Rule) || numData >= len(resource.Rule[numRule].Data) || numSecurity >= len(resource.Rule[numRule].Data[numData].Security) {
		return CodingSelect("Rule["+strconv.Itoa(numRule)+"]Data["+strconv.Itoa(numData)+"].Security["+strconv.Itoa(numSecurity)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Rule["+strconv.Itoa(numRule)+"]Data["+strconv.Itoa(numData)+"].Security["+strconv.Itoa(numSecurity)+"]", &resource.Rule[numRule].Data[numData].Security[numSecurity], optionsValueSet, htmlAttrs)
}
func (resource *Permission) T_RuleDataResourceMeaning(numRule int, numData int, numResource int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConsent_data_meaning

	if resource == nil || numRule >= len(resource.Rule) || numData >= len(resource.Rule[numRule].Data) || numResource >= len(resource.Rule[numRule].Data[numData].Resource) {
		return CodeSelect("Rule["+strconv.Itoa(numRule)+"]Data["+strconv.Itoa(numData)+"].Resource["+strconv.Itoa(numResource)+"].Meaning", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Rule["+strconv.Itoa(numRule)+"]Data["+strconv.Itoa(numData)+"].Resource["+strconv.Itoa(numResource)+"].Meaning", &resource.Rule[numRule].Data[numData].Resource[numResource].Meaning, optionsValueSet, htmlAttrs)
}
func (resource *Permission) T_RuleActivityAction(numRule int, numActivity int, numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRule >= len(resource.Rule) || numActivity >= len(resource.Rule[numRule].Activity) || numAction >= len(resource.Rule[numRule].Activity[numActivity].Action) {
		return CodeableConceptSelect("Rule["+strconv.Itoa(numRule)+"]Activity["+strconv.Itoa(numActivity)+"].Action["+strconv.Itoa(numAction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Rule["+strconv.Itoa(numRule)+"]Activity["+strconv.Itoa(numActivity)+"].Action["+strconv.Itoa(numAction)+"]", &resource.Rule[numRule].Activity[numActivity].Action[numAction], optionsValueSet, htmlAttrs)
}
func (resource *Permission) T_RuleActivityPurpose(numRule int, numActivity int, numPurpose int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRule >= len(resource.Rule) || numActivity >= len(resource.Rule[numRule].Activity) || numPurpose >= len(resource.Rule[numRule].Activity[numActivity].Purpose) {
		return CodeableConceptSelect("Rule["+strconv.Itoa(numRule)+"]Activity["+strconv.Itoa(numActivity)+"].Purpose["+strconv.Itoa(numPurpose)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Rule["+strconv.Itoa(numRule)+"]Activity["+strconv.Itoa(numActivity)+"].Purpose["+strconv.Itoa(numPurpose)+"]", &resource.Rule[numRule].Activity[numActivity].Purpose[numPurpose], optionsValueSet, htmlAttrs)
}
