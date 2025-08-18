//generated August 18 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

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
	Date              *string               `json:"date,omitempty"`
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
	VerificationDate  []string         `json:"verificationDate,omitempty"`
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
