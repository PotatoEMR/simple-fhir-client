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

// http://hl7.org/fhir/r4/StructureDefinition/Consent
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
	DateTime          *time.Time            `json:"dateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Performer         []Reference           `json:"performer,omitempty"`
	Organization      []Reference           `json:"organization,omitempty"`
	SourceAttachment  *Attachment           `json:"sourceAttachment,omitempty"`
	SourceReference   *Reference            `json:"sourceReference,omitempty"`
	Policy            []ConsentPolicy       `json:"policy,omitempty"`
	PolicyRule        *CodeableConcept      `json:"policyRule,omitempty"`
	Verification      []ConsentVerification `json:"verification,omitempty"`
	Provision         *ConsentProvision     `json:"provision,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Consent
type ConsentPolicy struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Authority         *string     `json:"authority,omitempty"`
	Uri               *string     `json:"uri,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Consent
type ConsentVerification struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Verified          bool        `json:"verified"`
	VerifiedWith      *Reference  `json:"verifiedWith,omitempty"`
	VerificationDate  *time.Time  `json:"verificationDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Consent
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

// http://hl7.org/fhir/r4/StructureDefinition/Consent
type ConsentProvisionActor struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Role              CodeableConcept `json:"role"`
	Reference         Reference       `json:"reference"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Consent
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
func (r Consent) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Consent/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Consent"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Consent) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSConsent_state_codes

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_Scope(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Scope", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Scope", &resource.Scope, optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_DateTime(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("DateTime", nil, htmlAttrs)
	}
	return DateTimeInput("DateTime", resource.DateTime, htmlAttrs)
}
func (resource *Consent) T_PolicyRule(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PolicyRule", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PolicyRule", resource.PolicyRule, optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_PolicyAuthority(numPolicy int, htmlAttrs string) templ.Component {
	if resource == nil || numPolicy >= len(resource.Policy) {
		return StringInput("Policy["+strconv.Itoa(numPolicy)+"]Authority", nil, htmlAttrs)
	}
	return StringInput("Policy["+strconv.Itoa(numPolicy)+"]Authority", resource.Policy[numPolicy].Authority, htmlAttrs)
}
func (resource *Consent) T_PolicyUri(numPolicy int, htmlAttrs string) templ.Component {
	if resource == nil || numPolicy >= len(resource.Policy) {
		return StringInput("Policy["+strconv.Itoa(numPolicy)+"]Uri", nil, htmlAttrs)
	}
	return StringInput("Policy["+strconv.Itoa(numPolicy)+"]Uri", resource.Policy[numPolicy].Uri, htmlAttrs)
}
func (resource *Consent) T_VerificationVerified(numVerification int, htmlAttrs string) templ.Component {
	if resource == nil || numVerification >= len(resource.Verification) {
		return BoolInput("Verification["+strconv.Itoa(numVerification)+"]Verified", nil, htmlAttrs)
	}
	return BoolInput("Verification["+strconv.Itoa(numVerification)+"]Verified", &resource.Verification[numVerification].Verified, htmlAttrs)
}
func (resource *Consent) T_VerificationVerificationDate(numVerification int, htmlAttrs string) templ.Component {
	if resource == nil || numVerification >= len(resource.Verification) {
		return DateTimeInput("Verification["+strconv.Itoa(numVerification)+"]VerificationDate", nil, htmlAttrs)
	}
	return DateTimeInput("Verification["+strconv.Itoa(numVerification)+"]VerificationDate", resource.Verification[numVerification].VerificationDate, htmlAttrs)
}
func (resource *Consent) T_ProvisionType(htmlAttrs string) templ.Component {
	optionsValueSet := VSConsent_provision_type

	if resource == nil {
		return CodeSelect("ProvisionType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ProvisionType", resource.Provision.Type, optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionAction(numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAction >= len(resource.Provision.Action) {
		return CodeableConceptSelect("ProvisionAction["+strconv.Itoa(numAction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ProvisionAction["+strconv.Itoa(numAction)+"]", &resource.Provision.Action[numAction], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionSecurityLabel(numSecurityLabel int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSecurityLabel >= len(resource.Provision.SecurityLabel) {
		return CodingSelect("ProvisionSecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("ProvisionSecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", &resource.Provision.SecurityLabel[numSecurityLabel], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionPurpose(numPurpose int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPurpose >= len(resource.Provision.Purpose) {
		return CodingSelect("ProvisionPurpose["+strconv.Itoa(numPurpose)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("ProvisionPurpose["+strconv.Itoa(numPurpose)+"]", &resource.Provision.Purpose[numPurpose], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionClass(numClass int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numClass >= len(resource.Provision.Class) {
		return CodingSelect("ProvisionClass["+strconv.Itoa(numClass)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("ProvisionClass["+strconv.Itoa(numClass)+"]", &resource.Provision.Class[numClass], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionCode(numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCode >= len(resource.Provision.Code) {
		return CodeableConceptSelect("ProvisionCode["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ProvisionCode["+strconv.Itoa(numCode)+"]", &resource.Provision.Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionActorRole(numActor int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numActor >= len(resource.Provision.Actor) {
		return CodeableConceptSelect("ProvisionActor["+strconv.Itoa(numActor)+"].Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ProvisionActor["+strconv.Itoa(numActor)+"].Role", &resource.Provision.Actor[numActor].Role, optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionDataMeaning(numData int, htmlAttrs string) templ.Component {
	optionsValueSet := VSConsent_data_meaning

	if resource == nil || numData >= len(resource.Provision.Data) {
		return CodeSelect("ProvisionData["+strconv.Itoa(numData)+"].Meaning", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ProvisionData["+strconv.Itoa(numData)+"].Meaning", &resource.Provision.Data[numData].Meaning, optionsValueSet, htmlAttrs)
}
