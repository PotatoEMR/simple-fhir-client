package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
	SourceAttachment  *Attachment           `json:"sourceAttachment,omitempty"`
	SourceReference   *Reference            `json:"sourceReference,omitempty"`
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

func (resource *Consent) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Consent.Id", nil)
	}
	return StringInput("Consent.Id", resource.Id)
}
func (resource *Consent) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Consent.ImplicitRules", nil)
	}
	return StringInput("Consent.ImplicitRules", resource.ImplicitRules)
}
func (resource *Consent) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Consent.Language", nil, optionsValueSet)
	}
	return CodeSelect("Consent.Language", resource.Language, optionsValueSet)
}
func (resource *Consent) T_Status() templ.Component {
	optionsValueSet := VSConsent_state_codes

	if resource == nil {
		return CodeSelect("Consent.Status", nil, optionsValueSet)
	}
	return CodeSelect("Consent.Status", &resource.Status, optionsValueSet)
}
func (resource *Consent) T_Scope(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Consent.Scope", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Consent.Scope", &resource.Scope, optionsValueSet)
}
func (resource *Consent) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("Consent.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Consent.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *Consent) T_DateTime() templ.Component {

	if resource == nil {
		return StringInput("Consent.DateTime", nil)
	}
	return StringInput("Consent.DateTime", resource.DateTime)
}
func (resource *Consent) T_PolicyRule(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Consent.PolicyRule", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Consent.PolicyRule", resource.PolicyRule, optionsValueSet)
}
func (resource *Consent) T_PolicyId(numPolicy int) templ.Component {

	if resource == nil || len(resource.Policy) >= numPolicy {
		return StringInput("Consent.Policy["+strconv.Itoa(numPolicy)+"].Id", nil)
	}
	return StringInput("Consent.Policy["+strconv.Itoa(numPolicy)+"].Id", resource.Policy[numPolicy].Id)
}
func (resource *Consent) T_PolicyAuthority(numPolicy int) templ.Component {

	if resource == nil || len(resource.Policy) >= numPolicy {
		return StringInput("Consent.Policy["+strconv.Itoa(numPolicy)+"].Authority", nil)
	}
	return StringInput("Consent.Policy["+strconv.Itoa(numPolicy)+"].Authority", resource.Policy[numPolicy].Authority)
}
func (resource *Consent) T_PolicyUri(numPolicy int) templ.Component {

	if resource == nil || len(resource.Policy) >= numPolicy {
		return StringInput("Consent.Policy["+strconv.Itoa(numPolicy)+"].Uri", nil)
	}
	return StringInput("Consent.Policy["+strconv.Itoa(numPolicy)+"].Uri", resource.Policy[numPolicy].Uri)
}
func (resource *Consent) T_VerificationId(numVerification int) templ.Component {

	if resource == nil || len(resource.Verification) >= numVerification {
		return StringInput("Consent.Verification["+strconv.Itoa(numVerification)+"].Id", nil)
	}
	return StringInput("Consent.Verification["+strconv.Itoa(numVerification)+"].Id", resource.Verification[numVerification].Id)
}
func (resource *Consent) T_VerificationVerified(numVerification int) templ.Component {

	if resource == nil || len(resource.Verification) >= numVerification {
		return BoolInput("Consent.Verification["+strconv.Itoa(numVerification)+"].Verified", nil)
	}
	return BoolInput("Consent.Verification["+strconv.Itoa(numVerification)+"].Verified", &resource.Verification[numVerification].Verified)
}
func (resource *Consent) T_VerificationVerificationDate(numVerification int) templ.Component {

	if resource == nil || len(resource.Verification) >= numVerification {
		return StringInput("Consent.Verification["+strconv.Itoa(numVerification)+"].VerificationDate", nil)
	}
	return StringInput("Consent.Verification["+strconv.Itoa(numVerification)+"].VerificationDate", resource.Verification[numVerification].VerificationDate)
}
func (resource *Consent) T_ProvisionId() templ.Component {

	if resource == nil {
		return StringInput("Consent.Provision.Id", nil)
	}
	return StringInput("Consent.Provision.Id", resource.Provision.Id)
}
func (resource *Consent) T_ProvisionType() templ.Component {
	optionsValueSet := VSConsent_provision_type

	if resource == nil {
		return CodeSelect("Consent.Provision.Type", nil, optionsValueSet)
	}
	return CodeSelect("Consent.Provision.Type", resource.Provision.Type, optionsValueSet)
}
func (resource *Consent) T_ProvisionAction(numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Provision.Action) >= numAction {
		return CodeableConceptSelect("Consent.Provision.Action["+strconv.Itoa(numAction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Consent.Provision.Action["+strconv.Itoa(numAction)+"]", &resource.Provision.Action[numAction], optionsValueSet)
}
func (resource *Consent) T_ProvisionSecurityLabel(numSecurityLabel int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Provision.SecurityLabel) >= numSecurityLabel {
		return CodingSelect("Consent.Provision.SecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", nil, optionsValueSet)
	}
	return CodingSelect("Consent.Provision.SecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", &resource.Provision.SecurityLabel[numSecurityLabel], optionsValueSet)
}
func (resource *Consent) T_ProvisionPurpose(numPurpose int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Provision.Purpose) >= numPurpose {
		return CodingSelect("Consent.Provision.Purpose["+strconv.Itoa(numPurpose)+"]", nil, optionsValueSet)
	}
	return CodingSelect("Consent.Provision.Purpose["+strconv.Itoa(numPurpose)+"]", &resource.Provision.Purpose[numPurpose], optionsValueSet)
}
func (resource *Consent) T_ProvisionClass(numClass int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Provision.Class) >= numClass {
		return CodingSelect("Consent.Provision.Class["+strconv.Itoa(numClass)+"]", nil, optionsValueSet)
	}
	return CodingSelect("Consent.Provision.Class["+strconv.Itoa(numClass)+"]", &resource.Provision.Class[numClass], optionsValueSet)
}
func (resource *Consent) T_ProvisionCode(numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Provision.Code) >= numCode {
		return CodeableConceptSelect("Consent.Provision.Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Consent.Provision.Code["+strconv.Itoa(numCode)+"]", &resource.Provision.Code[numCode], optionsValueSet)
}
func (resource *Consent) T_ProvisionActorId(numActor int) templ.Component {

	if resource == nil || len(resource.Provision.Actor) >= numActor {
		return StringInput("Consent.Provision.Actor["+strconv.Itoa(numActor)+"].Id", nil)
	}
	return StringInput("Consent.Provision.Actor["+strconv.Itoa(numActor)+"].Id", resource.Provision.Actor[numActor].Id)
}
func (resource *Consent) T_ProvisionActorRole(numActor int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Provision.Actor) >= numActor {
		return CodeableConceptSelect("Consent.Provision.Actor["+strconv.Itoa(numActor)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Consent.Provision.Actor["+strconv.Itoa(numActor)+"].Role", &resource.Provision.Actor[numActor].Role, optionsValueSet)
}
func (resource *Consent) T_ProvisionDataId(numData int) templ.Component {

	if resource == nil || len(resource.Provision.Data) >= numData {
		return StringInput("Consent.Provision.Data["+strconv.Itoa(numData)+"].Id", nil)
	}
	return StringInput("Consent.Provision.Data["+strconv.Itoa(numData)+"].Id", resource.Provision.Data[numData].Id)
}
func (resource *Consent) T_ProvisionDataMeaning(numData int) templ.Component {
	optionsValueSet := VSConsent_data_meaning

	if resource == nil || len(resource.Provision.Data) >= numData {
		return CodeSelect("Consent.Provision.Data["+strconv.Itoa(numData)+"].Meaning", nil, optionsValueSet)
	}
	return CodeSelect("Consent.Provision.Data["+strconv.Itoa(numData)+"].Meaning", &resource.Provision.Data[numData].Meaning, optionsValueSet)
}
