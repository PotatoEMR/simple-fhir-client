package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *Consent) ConsentLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Consent) ConsentStatus() templ.Component {
	optionsValueSet := VSConsent_state_codes

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Consent) ConsentCategory(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Category[0], optionsValueSet)
}
func (resource *Consent) ConsentRegulatoryBasis(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("regulatoryBasis", nil, optionsValueSet)
	}
	return CodeableConceptSelect("regulatoryBasis", &resource.RegulatoryBasis[0], optionsValueSet)
}
func (resource *Consent) ConsentDecision() templ.Component {
	optionsValueSet := VSConsent_provision_type

	if resource == nil {
		return CodeSelect("decision", nil, optionsValueSet)
	}
	return CodeSelect("decision", resource.Decision, optionsValueSet)
}
func (resource *Consent) ConsentVerificationVerificationType(numVerification int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Verification) >= numVerification {
		return CodeableConceptSelect("verificationType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("verificationType", resource.Verification[numVerification].VerificationType, optionsValueSet)
}
func (resource *Consent) ConsentProvisionAction(numProvision int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Provision) >= numProvision {
		return CodeableConceptSelect("action", nil, optionsValueSet)
	}
	return CodeableConceptSelect("action", &resource.Provision[numProvision].Action[0], optionsValueSet)
}
func (resource *Consent) ConsentProvisionSecurityLabel(numProvision int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Provision) >= numProvision {
		return CodingSelect("securityLabel", nil, optionsValueSet)
	}
	return CodingSelect("securityLabel", &resource.Provision[numProvision].SecurityLabel[0], optionsValueSet)
}
func (resource *Consent) ConsentProvisionPurpose(numProvision int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Provision) >= numProvision {
		return CodingSelect("purpose", nil, optionsValueSet)
	}
	return CodingSelect("purpose", &resource.Provision[numProvision].Purpose[0], optionsValueSet)
}
func (resource *Consent) ConsentProvisionDocumentType(numProvision int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Provision) >= numProvision {
		return CodingSelect("documentType", nil, optionsValueSet)
	}
	return CodingSelect("documentType", &resource.Provision[numProvision].DocumentType[0], optionsValueSet)
}
func (resource *Consent) ConsentProvisionResourceType(numProvision int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Provision) >= numProvision {
		return CodingSelect("resourceType", nil, optionsValueSet)
	}
	return CodingSelect("resourceType", &resource.Provision[numProvision].ResourceType[0], optionsValueSet)
}
func (resource *Consent) ConsentProvisionCode(numProvision int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Provision) >= numProvision {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Provision[numProvision].Code[0], optionsValueSet)
}
func (resource *Consent) ConsentProvisionActorRole(numProvision int, numActor int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Provision[numProvision].Actor) >= numActor {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", resource.Provision[numProvision].Actor[numActor].Role, optionsValueSet)
}
func (resource *Consent) ConsentProvisionDataMeaning(numProvision int, numData int) templ.Component {
	optionsValueSet := VSConsent_data_meaning

	if resource == nil && len(resource.Provision[numProvision].Data) >= numData {
		return CodeSelect("meaning", nil, optionsValueSet)
	}
	return CodeSelect("meaning", &resource.Provision[numProvision].Data[numData].Meaning, optionsValueSet)
}
