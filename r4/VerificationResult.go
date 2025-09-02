package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/VerificationResult
type VerificationResult struct {
	Id                *string                           `json:"id,omitempty"`
	Meta              *Meta                             `json:"meta,omitempty"`
	ImplicitRules     *string                           `json:"implicitRules,omitempty"`
	Language          *string                           `json:"language,omitempty"`
	Text              *Narrative                        `json:"text,omitempty"`
	Contained         []Resource                        `json:"contained,omitempty"`
	Extension         []Extension                       `json:"extension,omitempty"`
	ModifierExtension []Extension                       `json:"modifierExtension,omitempty"`
	Target            []Reference                       `json:"target,omitempty"`
	TargetLocation    []string                          `json:"targetLocation,omitempty"`
	Need              *CodeableConcept                  `json:"need,omitempty"`
	Status            string                            `json:"status"`
	StatusDate        *string                           `json:"statusDate,omitempty"`
	ValidationType    *CodeableConcept                  `json:"validationType,omitempty"`
	ValidationProcess []CodeableConcept                 `json:"validationProcess,omitempty"`
	Frequency         *Timing                           `json:"frequency,omitempty"`
	LastPerformed     *string                           `json:"lastPerformed,omitempty"`
	NextScheduled     *string                           `json:"nextScheduled,omitempty"`
	FailureAction     *CodeableConcept                  `json:"failureAction,omitempty"`
	PrimarySource     []VerificationResultPrimarySource `json:"primarySource,omitempty"`
	Attestation       *VerificationResultAttestation    `json:"attestation,omitempty"`
	Validator         []VerificationResultValidator     `json:"validator,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/VerificationResult
type VerificationResultPrimarySource struct {
	Id                  *string           `json:"id,omitempty"`
	Extension           []Extension       `json:"extension,omitempty"`
	ModifierExtension   []Extension       `json:"modifierExtension,omitempty"`
	Who                 *Reference        `json:"who,omitempty"`
	Type                []CodeableConcept `json:"type,omitempty"`
	CommunicationMethod []CodeableConcept `json:"communicationMethod,omitempty"`
	ValidationStatus    *CodeableConcept  `json:"validationStatus,omitempty"`
	ValidationDate      *string           `json:"validationDate,omitempty"`
	CanPushUpdates      *CodeableConcept  `json:"canPushUpdates,omitempty"`
	PushTypeAvailable   []CodeableConcept `json:"pushTypeAvailable,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/VerificationResult
type VerificationResultAttestation struct {
	Id                        *string          `json:"id,omitempty"`
	Extension                 []Extension      `json:"extension,omitempty"`
	ModifierExtension         []Extension      `json:"modifierExtension,omitempty"`
	Who                       *Reference       `json:"who,omitempty"`
	OnBehalfOf                *Reference       `json:"onBehalfOf,omitempty"`
	CommunicationMethod       *CodeableConcept `json:"communicationMethod,omitempty"`
	Date                      *string          `json:"date,omitempty"`
	SourceIdentityCertificate *string          `json:"sourceIdentityCertificate,omitempty"`
	ProxyIdentityCertificate  *string          `json:"proxyIdentityCertificate,omitempty"`
	ProxySignature            *Signature       `json:"proxySignature,omitempty"`
	SourceSignature           *Signature       `json:"sourceSignature,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/VerificationResult
type VerificationResultValidator struct {
	Id                   *string     `json:"id,omitempty"`
	Extension            []Extension `json:"extension,omitempty"`
	ModifierExtension    []Extension `json:"modifierExtension,omitempty"`
	Organization         Reference   `json:"organization"`
	IdentityCertificate  *string     `json:"identityCertificate,omitempty"`
	AttestationSignature *Signature  `json:"attestationSignature,omitempty"`
}

type OtherVerificationResult VerificationResult

// on convert struct to json, automatically add resourceType=VerificationResult
func (r VerificationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherVerificationResult
		ResourceType string `json:"resourceType"`
	}{
		OtherVerificationResult: OtherVerificationResult(r),
		ResourceType:            "VerificationResult",
	})
}

func (resource *VerificationResult) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *VerificationResult) T_Need(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("need", nil, optionsValueSet)
	}
	return CodeableConceptSelect("need", resource.Need, optionsValueSet)
}
func (resource *VerificationResult) T_Status(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *VerificationResult) T_ValidationType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("validationType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("validationType", resource.ValidationType, optionsValueSet)
}
func (resource *VerificationResult) T_ValidationProcess(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("validationProcess", nil, optionsValueSet)
	}
	return CodeableConceptSelect("validationProcess", &resource.ValidationProcess[0], optionsValueSet)
}
func (resource *VerificationResult) T_FailureAction(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("failureAction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("failureAction", resource.FailureAction, optionsValueSet)
}
func (resource *VerificationResult) T_PrimarySourceType(numPrimarySource int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.PrimarySource) >= numPrimarySource {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.PrimarySource[numPrimarySource].Type[0], optionsValueSet)
}
func (resource *VerificationResult) T_PrimarySourceCommunicationMethod(numPrimarySource int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.PrimarySource) >= numPrimarySource {
		return CodeableConceptSelect("communicationMethod", nil, optionsValueSet)
	}
	return CodeableConceptSelect("communicationMethod", &resource.PrimarySource[numPrimarySource].CommunicationMethod[0], optionsValueSet)
}
func (resource *VerificationResult) T_PrimarySourceValidationStatus(numPrimarySource int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.PrimarySource) >= numPrimarySource {
		return CodeableConceptSelect("validationStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("validationStatus", resource.PrimarySource[numPrimarySource].ValidationStatus, optionsValueSet)
}
func (resource *VerificationResult) T_PrimarySourceCanPushUpdates(numPrimarySource int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.PrimarySource) >= numPrimarySource {
		return CodeableConceptSelect("canPushUpdates", nil, optionsValueSet)
	}
	return CodeableConceptSelect("canPushUpdates", resource.PrimarySource[numPrimarySource].CanPushUpdates, optionsValueSet)
}
func (resource *VerificationResult) T_PrimarySourcePushTypeAvailable(numPrimarySource int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.PrimarySource) >= numPrimarySource {
		return CodeableConceptSelect("pushTypeAvailable", nil, optionsValueSet)
	}
	return CodeableConceptSelect("pushTypeAvailable", &resource.PrimarySource[numPrimarySource].PushTypeAvailable[0], optionsValueSet)
}
func (resource *VerificationResult) T_AttestationCommunicationMethod(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("communicationMethod", nil, optionsValueSet)
	}
	return CodeableConceptSelect("communicationMethod", resource.Attestation.CommunicationMethod, optionsValueSet)
}
