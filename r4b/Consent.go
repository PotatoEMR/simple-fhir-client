package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/Consent
type Consent struct {
	Id                *string               `json:"id,omitempty"`
	Meta              *Meta                 `json:"meta,omitempty"`
	ImplicitRules     *string               `json:"implicitRules,omitempty"`
	Language          *string               `json:"language,omitempty"`
	Text              *Narrative            `json:"text,omitempty"`
	Contained         []Resource            `json:"contained,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `json:"identifier,omitempty"`
	Status            string                `json:"status"`
	Scope             CodeableConcept       `json:"scope"`
	Category          []CodeableConcept     `json:"category"`
	Patient           *Reference            `json:"patient,omitempty"`
	DateTime          *string               `json:"dateTime,omitempty"`
	Performer         []Reference           `json:"performer,omitempty"`
	Organization      []Reference           `json:"organization,omitempty"`
	SourceAttachment  *Attachment           `json:"sourceAttachment"`
	SourceReference   *Reference            `json:"sourceReference"`
	Policy            []ConsentPolicy       `json:"policy,omitempty"`
	PolicyRule        *CodeableConcept      `json:"policyRule,omitempty"`
	Verification      []ConsentVerification `json:"verification,omitempty"`
	Provision         *ConsentProvision     `json:"provision,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Consent
type ConsentPolicy struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Authority         *string     `json:"authority,omitempty"`
	Uri               *string     `json:"uri,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Consent
type ConsentVerification struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Verified          bool        `json:"verified"`
	VerifiedWith      *Reference  `json:"verifiedWith,omitempty"`
	VerificationDate  *string     `json:"verificationDate,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Consent
type ConsentProvision struct {
	Id                *string                 `json:"id,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Type              *string                 `json:"type,omitempty"`
	Period            *Period                 `json:"period,omitempty"`
	Actor             []ConsentProvisionActor `json:"actor,omitempty"`
	Action            []CodeableConcept       `json:"action,omitempty"`
	SecurityLabel     []Coding                `json:"securityLabel,omitempty"`
	Purpose           []Coding                `json:"purpose,omitempty"`
	Class             []Coding                `json:"class,omitempty"`
	Code              []CodeableConcept       `json:"code,omitempty"`
	DataPeriod        *Period                 `json:"dataPeriod,omitempty"`
	Data              []ConsentProvisionData  `json:"data,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Consent
type ConsentProvisionActor struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Role              CodeableConcept `json:"role"`
	Reference         Reference       `json:"reference"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Consent
type ConsentProvisionData struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Meaning           string      `json:"meaning"`
	Reference         Reference   `json:"reference"`
}

type OtherConsent Consent

// on convert struct to json, automatically add resourceType=Consent
func (r Consent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherConsent
		ResourceType string `json:"resourceType"`
	}{
		OtherConsent: OtherConsent(r),
		ResourceType: "Consent",
	})
}

func (resource *Consent) ConsentLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Consent) ConsentStatus() templ.Component {
	optionsValueSet := VSConsent_state_codes

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Consent) ConsentScope(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("scope", nil, optionsValueSet)
	}
	return CodeableConceptSelect("scope", &resource.Scope, optionsValueSet)
}
func (resource *Consent) ConsentCategory(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *Consent) ConsentPolicyRule(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("policyRule", nil, optionsValueSet)
	}
	return CodeableConceptSelect("policyRule", resource.PolicyRule, optionsValueSet)
}
func (resource *Consent) ConsentProvisionType() templ.Component {
	optionsValueSet := VSConsent_provision_type

	if resource != nil {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", resource.Provision.Type, optionsValueSet)
}
func (resource *Consent) ConsentProvisionAction(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("action", nil, optionsValueSet)
	}
	return CodeableConceptSelect("action", &resource.Provision.Action[0], optionsValueSet)
}
func (resource *Consent) ConsentProvisionSecurityLabel(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodingSelect("securityLabel", nil, optionsValueSet)
	}
	return CodingSelect("securityLabel", &resource.Provision.SecurityLabel[0], optionsValueSet)
}
func (resource *Consent) ConsentProvisionPurpose(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodingSelect("purpose", nil, optionsValueSet)
	}
	return CodingSelect("purpose", &resource.Provision.Purpose[0], optionsValueSet)
}
func (resource *Consent) ConsentProvisionClass(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodingSelect("class", nil, optionsValueSet)
	}
	return CodingSelect("class", &resource.Provision.Class[0], optionsValueSet)
}
func (resource *Consent) ConsentProvisionCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Provision.Code[0], optionsValueSet)
}
func (resource *Consent) ConsentProvisionActorRole(numActor int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Provision.Actor) >= numActor {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", &resource.Provision.Actor[numActor].Role, optionsValueSet)
}
func (resource *Consent) ConsentProvisionDataMeaning(numData int) templ.Component {
	optionsValueSet := VSConsent_data_meaning

	if resource != nil && len(resource.Provision.Data) >= numData {
		return CodeSelect("meaning", nil, optionsValueSet)
	}
	return CodeSelect("meaning", &resource.Provision.Data[numData].Meaning, optionsValueSet)
}
