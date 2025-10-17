package r4b

//generated with command go run ./bultaoreune
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
	DateTime          *FhirDateTime         `json:"dateTime,omitempty"`
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
	Id                *string       `json:"id,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	Verified          bool          `json:"verified"`
	VerifiedWith      *Reference    `json:"verifiedWith,omitempty"`
	VerificationDate  *FhirDateTime `json:"verificationDate,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
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
func (r Consent) ResourceType() string {
	return "Consent"
}

func (resource *Consent) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSConsent_state_codes

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_Scope(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("scope", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("scope", &resource.Scope, optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_Patient(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "patient", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "patient", resource.Patient, htmlAttrs)
}
func (resource *Consent) T_DateTime(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("dateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("dateTime", resource.DateTime, htmlAttrs)
}
func (resource *Consent) T_Performer(frs []FhirResource, numPerformer int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPerformer >= len(resource.Performer) {
		return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "performer["+strconv.Itoa(numPerformer)+"]", &resource.Performer[numPerformer], htmlAttrs)
}
func (resource *Consent) T_Organization(frs []FhirResource, numOrganization int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numOrganization >= len(resource.Organization) {
		return ReferenceInput(frs, "organization["+strconv.Itoa(numOrganization)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "organization["+strconv.Itoa(numOrganization)+"]", &resource.Organization[numOrganization], htmlAttrs)
}
func (resource *Consent) T_SourceAttachment(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return AttachmentInput("sourceAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("sourceAttachment", resource.SourceAttachment, htmlAttrs)
}
func (resource *Consent) T_SourceReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "sourceReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "sourceReference", resource.SourceReference, htmlAttrs)
}
func (resource *Consent) T_PolicyRule(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("policyRule", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("policyRule", resource.PolicyRule, optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_PolicyAuthority(numPolicy int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPolicy >= len(resource.Policy) {
		return StringInput("policy["+strconv.Itoa(numPolicy)+"].authority", nil, htmlAttrs)
	}
	return StringInput("policy["+strconv.Itoa(numPolicy)+"].authority", resource.Policy[numPolicy].Authority, htmlAttrs)
}
func (resource *Consent) T_PolicyUri(numPolicy int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPolicy >= len(resource.Policy) {
		return StringInput("policy["+strconv.Itoa(numPolicy)+"].uri", nil, htmlAttrs)
	}
	return StringInput("policy["+strconv.Itoa(numPolicy)+"].uri", resource.Policy[numPolicy].Uri, htmlAttrs)
}
func (resource *Consent) T_VerificationVerified(numVerification int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVerification >= len(resource.Verification) {
		return BoolInput("verification["+strconv.Itoa(numVerification)+"].verified", nil, htmlAttrs)
	}
	return BoolInput("verification["+strconv.Itoa(numVerification)+"].verified", &resource.Verification[numVerification].Verified, htmlAttrs)
}
func (resource *Consent) T_VerificationVerifiedWith(frs []FhirResource, numVerification int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVerification >= len(resource.Verification) {
		return ReferenceInput(frs, "verification["+strconv.Itoa(numVerification)+"].verifiedWith", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "verification["+strconv.Itoa(numVerification)+"].verifiedWith", resource.Verification[numVerification].VerifiedWith, htmlAttrs)
}
func (resource *Consent) T_VerificationVerificationDate(numVerification int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVerification >= len(resource.Verification) {
		return FhirDateTimeInput("verification["+strconv.Itoa(numVerification)+"].verificationDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("verification["+strconv.Itoa(numVerification)+"].verificationDate", resource.Verification[numVerification].VerificationDate, htmlAttrs)
}
func (resource *Consent) T_ProvisionType(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSConsent_provision_type

	if resource == nil {
		return CodeSelect("provision.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("provision.type", resource.Provision.Type, optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("provision.period", nil, htmlAttrs)
	}
	return PeriodInput("provision.period", resource.Provision.Period, htmlAttrs)
}
func (resource *Consent) T_ProvisionAction(numAction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAction >= len(resource.Provision.Action) {
		return CodeableConceptSelect("provision.action["+strconv.Itoa(numAction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("provision.action["+strconv.Itoa(numAction)+"]", &resource.Provision.Action[numAction], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionSecurityLabel(numSecurityLabel int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSecurityLabel >= len(resource.Provision.SecurityLabel) {
		return CodingSelect("provision.securityLabel["+strconv.Itoa(numSecurityLabel)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("provision.securityLabel["+strconv.Itoa(numSecurityLabel)+"]", &resource.Provision.SecurityLabel[numSecurityLabel], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionPurpose(numPurpose int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPurpose >= len(resource.Provision.Purpose) {
		return CodingSelect("provision.purpose["+strconv.Itoa(numPurpose)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("provision.purpose["+strconv.Itoa(numPurpose)+"]", &resource.Provision.Purpose[numPurpose], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionClass(numClass int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClass >= len(resource.Provision.Class) {
		return CodingSelect("provision.class["+strconv.Itoa(numClass)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("provision.class["+strconv.Itoa(numClass)+"]", &resource.Provision.Class[numClass], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionCode(numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCode >= len(resource.Provision.Code) {
		return CodeableConceptSelect("provision.code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("provision.code["+strconv.Itoa(numCode)+"]", &resource.Provision.Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionDataPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("provision.dataPeriod", nil, htmlAttrs)
	}
	return PeriodInput("provision.dataPeriod", resource.Provision.DataPeriod, htmlAttrs)
}
func (resource *Consent) T_ProvisionActorRole(numActor int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActor >= len(resource.Provision.Actor) {
		return CodeableConceptSelect("provision.actor["+strconv.Itoa(numActor)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("provision.actor["+strconv.Itoa(numActor)+"].role", &resource.Provision.Actor[numActor].Role, optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionActorReference(frs []FhirResource, numActor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numActor >= len(resource.Provision.Actor) {
		return ReferenceInput(frs, "provision.actor["+strconv.Itoa(numActor)+"].reference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "provision.actor["+strconv.Itoa(numActor)+"].reference", &resource.Provision.Actor[numActor].Reference, htmlAttrs)
}
func (resource *Consent) T_ProvisionDataMeaning(numData int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSConsent_data_meaning

	if resource == nil || numData >= len(resource.Provision.Data) {
		return CodeSelect("provision.data["+strconv.Itoa(numData)+"].meaning", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("provision.data["+strconv.Itoa(numData)+"].meaning", &resource.Provision.Data[numData].Meaning, optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionDataReference(frs []FhirResource, numData int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numData >= len(resource.Provision.Data) {
		return ReferenceInput(frs, "provision.data["+strconv.Itoa(numData)+"].reference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "provision.data["+strconv.Itoa(numData)+"].reference", &resource.Provision.Data[numData].Reference, htmlAttrs)
}
