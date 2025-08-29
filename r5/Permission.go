package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
func (resource *Permission) PermissionLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Permission) PermissionStatus() templ.Component {
	optionsValueSet := VSPermission_status
	currentVal := ""
	if resource != nil {
		currentVal = resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *Permission) PermissionCombining() templ.Component {
	optionsValueSet := VSPermission_rule_combining
	currentVal := ""
	if resource != nil {
		currentVal = resource.Combining
	}
	return CodeSelect("combining", currentVal, optionsValueSet)
}
func (resource *Permission) PermissionRuleType(numRule int) templ.Component {
	optionsValueSet := VSConsent_provision_type
	currentVal := ""
	if resource != nil && len(resource.Rule) >= numRule {
		currentVal = *resource.Rule[numRule].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
func (resource *Permission) PermissionRuleDataResourceMeaning(numRule int, numData int, numResource int) templ.Component {
	optionsValueSet := VSConsent_data_meaning
	currentVal := ""
	if resource != nil && len(resource.Rule[numRule].Data[numData].Resource) >= numResource {
		currentVal = resource.Rule[numRule].Data[numData].Resource[numResource].Meaning
	}
	return CodeSelect("meaning", currentVal, optionsValueSet)
}
