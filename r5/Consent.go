package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Consent
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
	Category          []CodeableConcept     `json:"category,omitempty"`
	Subject           *Reference            `json:"subject,omitempty"`
	Date              *FhirDate             `json:"date,omitempty"`
	Period            *Period               `json:"period,omitempty"`
	Grantor           []Reference           `json:"grantor,omitempty"`
	Grantee           []Reference           `json:"grantee,omitempty"`
	Manager           []Reference           `json:"manager,omitempty"`
	Controller        []Reference           `json:"controller,omitempty"`
	SourceAttachment  []Attachment          `json:"sourceAttachment,omitempty"`
	SourceReference   []Reference           `json:"sourceReference,omitempty"`
	RegulatoryBasis   []CodeableConcept     `json:"regulatoryBasis,omitempty"`
	PolicyBasis       *ConsentPolicyBasis   `json:"policyBasis,omitempty"`
	PolicyText        []Reference           `json:"policyText,omitempty"`
	Verification      []ConsentVerification `json:"verification,omitempty"`
	Decision          *string               `json:"decision,omitempty"`
	Provision         []ConsentProvision    `json:"provision,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Consent
type ConsentPolicyBasis struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Reference         *Reference  `json:"reference,omitempty"`
	Url               *string     `json:"url,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Consent
type ConsentVerification struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Verified          bool             `json:"verified"`
	VerificationType  *CodeableConcept `json:"verificationType,omitempty"`
	VerifiedBy        *Reference       `json:"verifiedBy,omitempty"`
	VerifiedWith      *Reference       `json:"verifiedWith,omitempty"`
	VerificationDate  []FhirDateTime   `json:"verificationDate,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Consent
type ConsentProvision struct {
	Id                *string                 `json:"id,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Period            *Period                 `json:"period,omitempty"`
	Actor             []ConsentProvisionActor `json:"actor,omitempty"`
	Action            []CodeableConcept       `json:"action,omitempty"`
	SecurityLabel     []Coding                `json:"securityLabel,omitempty"`
	Purpose           []Coding                `json:"purpose,omitempty"`
	DocumentType      []Coding                `json:"documentType,omitempty"`
	ResourceType      []Coding                `json:"resourceType,omitempty"`
	Code              []CodeableConcept       `json:"code,omitempty"`
	DataPeriod        *Period                 `json:"dataPeriod,omitempty"`
	Data              []ConsentProvisionData  `json:"data,omitempty"`
	Expression        *Expression             `json:"expression,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Consent
type ConsentProvisionActor struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Reference         *Reference       `json:"reference,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Consent
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
func (resource *Consent) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSConsent_state_codes

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_Date(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateInput("date", nil, htmlAttrs)
	}
	return DateInput("date", resource.Date, htmlAttrs)
}
func (resource *Consent) T_RegulatoryBasis(numRegulatoryBasis int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRegulatoryBasis >= len(resource.RegulatoryBasis) {
		return CodeableConceptSelect("regulatoryBasis["+strconv.Itoa(numRegulatoryBasis)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("regulatoryBasis["+strconv.Itoa(numRegulatoryBasis)+"]", &resource.RegulatoryBasis[numRegulatoryBasis], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_Decision(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSConsent_provision_type

	if resource == nil {
		return CodeSelect("decision", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("decision", resource.Decision, optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_PolicyBasisUrl(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("policyBasis.url", nil, htmlAttrs)
	}
	return StringInput("policyBasis.url", resource.PolicyBasis.Url, htmlAttrs)
}
func (resource *Consent) T_VerificationVerified(numVerification int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVerification >= len(resource.Verification) {
		return BoolInput("verification["+strconv.Itoa(numVerification)+"].verified", nil, htmlAttrs)
	}
	return BoolInput("verification["+strconv.Itoa(numVerification)+"].verified", &resource.Verification[numVerification].Verified, htmlAttrs)
}
func (resource *Consent) T_VerificationVerificationType(numVerification int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVerification >= len(resource.Verification) {
		return CodeableConceptSelect("verification["+strconv.Itoa(numVerification)+"].verificationType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("verification["+strconv.Itoa(numVerification)+"].verificationType", resource.Verification[numVerification].VerificationType, optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_VerificationVerificationDate(numVerification int, numVerificationDate int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVerification >= len(resource.Verification) || numVerificationDate >= len(resource.Verification[numVerification].VerificationDate) {
		return DateTimeInput("verification["+strconv.Itoa(numVerification)+"].verificationDate["+strconv.Itoa(numVerificationDate)+"]", nil, htmlAttrs)
	}
	return DateTimeInput("verification["+strconv.Itoa(numVerification)+"].verificationDate["+strconv.Itoa(numVerificationDate)+"]", &resource.Verification[numVerification].VerificationDate[numVerificationDate], htmlAttrs)
}
func (resource *Consent) T_ProvisionAction(numProvision int, numAction int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProvision >= len(resource.Provision) || numAction >= len(resource.Provision[numProvision].Action) {
		return CodeableConceptSelect("provision["+strconv.Itoa(numProvision)+"].action["+strconv.Itoa(numAction)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("provision["+strconv.Itoa(numProvision)+"].action["+strconv.Itoa(numAction)+"]", &resource.Provision[numProvision].Action[numAction], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionSecurityLabel(numProvision int, numSecurityLabel int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProvision >= len(resource.Provision) || numSecurityLabel >= len(resource.Provision[numProvision].SecurityLabel) {
		return CodingSelect("provision["+strconv.Itoa(numProvision)+"].securityLabel["+strconv.Itoa(numSecurityLabel)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("provision["+strconv.Itoa(numProvision)+"].securityLabel["+strconv.Itoa(numSecurityLabel)+"]", &resource.Provision[numProvision].SecurityLabel[numSecurityLabel], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionPurpose(numProvision int, numPurpose int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProvision >= len(resource.Provision) || numPurpose >= len(resource.Provision[numProvision].Purpose) {
		return CodingSelect("provision["+strconv.Itoa(numProvision)+"].purpose["+strconv.Itoa(numPurpose)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("provision["+strconv.Itoa(numProvision)+"].purpose["+strconv.Itoa(numPurpose)+"]", &resource.Provision[numProvision].Purpose[numPurpose], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionDocumentType(numProvision int, numDocumentType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProvision >= len(resource.Provision) || numDocumentType >= len(resource.Provision[numProvision].DocumentType) {
		return CodingSelect("provision["+strconv.Itoa(numProvision)+"].documentType["+strconv.Itoa(numDocumentType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("provision["+strconv.Itoa(numProvision)+"].documentType["+strconv.Itoa(numDocumentType)+"]", &resource.Provision[numProvision].DocumentType[numDocumentType], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionResourceType(numProvision int, numResourceType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProvision >= len(resource.Provision) || numResourceType >= len(resource.Provision[numProvision].ResourceType) {
		return CodingSelect("provision["+strconv.Itoa(numProvision)+"].resourceType["+strconv.Itoa(numResourceType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("provision["+strconv.Itoa(numProvision)+"].resourceType["+strconv.Itoa(numResourceType)+"]", &resource.Provision[numProvision].ResourceType[numResourceType], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionCode(numProvision int, numCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProvision >= len(resource.Provision) || numCode >= len(resource.Provision[numProvision].Code) {
		return CodeableConceptSelect("provision["+strconv.Itoa(numProvision)+"].code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("provision["+strconv.Itoa(numProvision)+"].code["+strconv.Itoa(numCode)+"]", &resource.Provision[numProvision].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionActorRole(numProvision int, numActor int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProvision >= len(resource.Provision) || numActor >= len(resource.Provision[numProvision].Actor) {
		return CodeableConceptSelect("provision["+strconv.Itoa(numProvision)+"].actor["+strconv.Itoa(numActor)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("provision["+strconv.Itoa(numProvision)+"].actor["+strconv.Itoa(numActor)+"].role", resource.Provision[numProvision].Actor[numActor].Role, optionsValueSet, htmlAttrs)
}
func (resource *Consent) T_ProvisionDataMeaning(numProvision int, numData int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSConsent_data_meaning

	if resource == nil || numProvision >= len(resource.Provision) || numData >= len(resource.Provision[numProvision].Data) {
		return CodeSelect("provision["+strconv.Itoa(numProvision)+"].data["+strconv.Itoa(numData)+"].meaning", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("provision["+strconv.Itoa(numProvision)+"].data["+strconv.Itoa(numData)+"].meaning", &resource.Provision[numProvision].Data[numData].Meaning, optionsValueSet, htmlAttrs)
}
