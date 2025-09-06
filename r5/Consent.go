package r5

//generated with command go run ./bultaoreune -nodownload
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
func (resource *Consent) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("Consent.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Consent.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *Consent) T_Date() templ.Component {

	if resource == nil {
		return StringInput("Consent.Date", nil)
	}
	return StringInput("Consent.Date", resource.Date)
}
func (resource *Consent) T_RegulatoryBasis(numRegulatoryBasis int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.RegulatoryBasis) >= numRegulatoryBasis {
		return CodeableConceptSelect("Consent.RegulatoryBasis["+strconv.Itoa(numRegulatoryBasis)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Consent.RegulatoryBasis["+strconv.Itoa(numRegulatoryBasis)+"]", &resource.RegulatoryBasis[numRegulatoryBasis], optionsValueSet)
}
func (resource *Consent) T_Decision() templ.Component {
	optionsValueSet := VSConsent_provision_type

	if resource == nil {
		return CodeSelect("Consent.Decision", nil, optionsValueSet)
	}
	return CodeSelect("Consent.Decision", resource.Decision, optionsValueSet)
}
func (resource *Consent) T_PolicyBasisId() templ.Component {

	if resource == nil {
		return StringInput("Consent.PolicyBasis.Id", nil)
	}
	return StringInput("Consent.PolicyBasis.Id", resource.PolicyBasis.Id)
}
func (resource *Consent) T_PolicyBasisUrl() templ.Component {

	if resource == nil {
		return StringInput("Consent.PolicyBasis.Url", nil)
	}
	return StringInput("Consent.PolicyBasis.Url", resource.PolicyBasis.Url)
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
func (resource *Consent) T_VerificationVerificationType(numVerification int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Verification) >= numVerification {
		return CodeableConceptSelect("Consent.Verification["+strconv.Itoa(numVerification)+"].VerificationType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Consent.Verification["+strconv.Itoa(numVerification)+"].VerificationType", resource.Verification[numVerification].VerificationType, optionsValueSet)
}
func (resource *Consent) T_VerificationVerificationDate(numVerification int, numVerificationDate int) templ.Component {

	if resource == nil || len(resource.Verification) >= numVerification || len(resource.Verification[numVerification].VerificationDate) >= numVerificationDate {
		return StringInput("Consent.Verification["+strconv.Itoa(numVerification)+"].VerificationDate["+strconv.Itoa(numVerificationDate)+"]", nil)
	}
	return StringInput("Consent.Verification["+strconv.Itoa(numVerification)+"].VerificationDate["+strconv.Itoa(numVerificationDate)+"]", &resource.Verification[numVerification].VerificationDate[numVerificationDate])
}
func (resource *Consent) T_ProvisionId(numProvision int) templ.Component {

	if resource == nil || len(resource.Provision) >= numProvision {
		return StringInput("Consent.Provision["+strconv.Itoa(numProvision)+"].Id", nil)
	}
	return StringInput("Consent.Provision["+strconv.Itoa(numProvision)+"].Id", resource.Provision[numProvision].Id)
}
func (resource *Consent) T_ProvisionAction(numProvision int, numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Provision) >= numProvision || len(resource.Provision[numProvision].Action) >= numAction {
		return CodeableConceptSelect("Consent.Provision["+strconv.Itoa(numProvision)+"].Action["+strconv.Itoa(numAction)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Consent.Provision["+strconv.Itoa(numProvision)+"].Action["+strconv.Itoa(numAction)+"]", &resource.Provision[numProvision].Action[numAction], optionsValueSet)
}
func (resource *Consent) T_ProvisionSecurityLabel(numProvision int, numSecurityLabel int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Provision) >= numProvision || len(resource.Provision[numProvision].SecurityLabel) >= numSecurityLabel {
		return CodingSelect("Consent.Provision["+strconv.Itoa(numProvision)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", nil, optionsValueSet)
	}
	return CodingSelect("Consent.Provision["+strconv.Itoa(numProvision)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"]", &resource.Provision[numProvision].SecurityLabel[numSecurityLabel], optionsValueSet)
}
func (resource *Consent) T_ProvisionPurpose(numProvision int, numPurpose int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Provision) >= numProvision || len(resource.Provision[numProvision].Purpose) >= numPurpose {
		return CodingSelect("Consent.Provision["+strconv.Itoa(numProvision)+"].Purpose["+strconv.Itoa(numPurpose)+"]", nil, optionsValueSet)
	}
	return CodingSelect("Consent.Provision["+strconv.Itoa(numProvision)+"].Purpose["+strconv.Itoa(numPurpose)+"]", &resource.Provision[numProvision].Purpose[numPurpose], optionsValueSet)
}
func (resource *Consent) T_ProvisionDocumentType(numProvision int, numDocumentType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Provision) >= numProvision || len(resource.Provision[numProvision].DocumentType) >= numDocumentType {
		return CodingSelect("Consent.Provision["+strconv.Itoa(numProvision)+"].DocumentType["+strconv.Itoa(numDocumentType)+"]", nil, optionsValueSet)
	}
	return CodingSelect("Consent.Provision["+strconv.Itoa(numProvision)+"].DocumentType["+strconv.Itoa(numDocumentType)+"]", &resource.Provision[numProvision].DocumentType[numDocumentType], optionsValueSet)
}
func (resource *Consent) T_ProvisionResourceType(numProvision int, numResourceType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Provision) >= numProvision || len(resource.Provision[numProvision].ResourceType) >= numResourceType {
		return CodingSelect("Consent.Provision["+strconv.Itoa(numProvision)+"].ResourceType["+strconv.Itoa(numResourceType)+"]", nil, optionsValueSet)
	}
	return CodingSelect("Consent.Provision["+strconv.Itoa(numProvision)+"].ResourceType["+strconv.Itoa(numResourceType)+"]", &resource.Provision[numProvision].ResourceType[numResourceType], optionsValueSet)
}
func (resource *Consent) T_ProvisionCode(numProvision int, numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Provision) >= numProvision || len(resource.Provision[numProvision].Code) >= numCode {
		return CodeableConceptSelect("Consent.Provision["+strconv.Itoa(numProvision)+"].Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Consent.Provision["+strconv.Itoa(numProvision)+"].Code["+strconv.Itoa(numCode)+"]", &resource.Provision[numProvision].Code[numCode], optionsValueSet)
}
func (resource *Consent) T_ProvisionActorId(numProvision int, numActor int) templ.Component {

	if resource == nil || len(resource.Provision) >= numProvision || len(resource.Provision[numProvision].Actor) >= numActor {
		return StringInput("Consent.Provision["+strconv.Itoa(numProvision)+"].Actor["+strconv.Itoa(numActor)+"].Id", nil)
	}
	return StringInput("Consent.Provision["+strconv.Itoa(numProvision)+"].Actor["+strconv.Itoa(numActor)+"].Id", resource.Provision[numProvision].Actor[numActor].Id)
}
func (resource *Consent) T_ProvisionActorRole(numProvision int, numActor int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Provision) >= numProvision || len(resource.Provision[numProvision].Actor) >= numActor {
		return CodeableConceptSelect("Consent.Provision["+strconv.Itoa(numProvision)+"].Actor["+strconv.Itoa(numActor)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Consent.Provision["+strconv.Itoa(numProvision)+"].Actor["+strconv.Itoa(numActor)+"].Role", resource.Provision[numProvision].Actor[numActor].Role, optionsValueSet)
}
func (resource *Consent) T_ProvisionDataId(numProvision int, numData int) templ.Component {

	if resource == nil || len(resource.Provision) >= numProvision || len(resource.Provision[numProvision].Data) >= numData {
		return StringInput("Consent.Provision["+strconv.Itoa(numProvision)+"].Data["+strconv.Itoa(numData)+"].Id", nil)
	}
	return StringInput("Consent.Provision["+strconv.Itoa(numProvision)+"].Data["+strconv.Itoa(numData)+"].Id", resource.Provision[numProvision].Data[numData].Id)
}
func (resource *Consent) T_ProvisionDataMeaning(numProvision int, numData int) templ.Component {
	optionsValueSet := VSConsent_data_meaning

	if resource == nil || len(resource.Provision) >= numProvision || len(resource.Provision[numProvision].Data) >= numData {
		return CodeSelect("Consent.Provision["+strconv.Itoa(numProvision)+"].Data["+strconv.Itoa(numData)+"].Meaning", nil, optionsValueSet)
	}
	return CodeSelect("Consent.Provision["+strconv.Itoa(numProvision)+"].Data["+strconv.Itoa(numData)+"].Meaning", &resource.Provision[numProvision].Data[numData].Meaning, optionsValueSet)
}
