package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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

func (resource *VerificationResult) T_Id() templ.Component {

	if resource == nil {
		return StringInput("VerificationResult.Id", nil)
	}
	return StringInput("VerificationResult.Id", resource.Id)
}
func (resource *VerificationResult) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("VerificationResult.ImplicitRules", nil)
	}
	return StringInput("VerificationResult.ImplicitRules", resource.ImplicitRules)
}
func (resource *VerificationResult) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("VerificationResult.Language", nil, optionsValueSet)
	}
	return CodeSelect("VerificationResult.Language", resource.Language, optionsValueSet)
}
func (resource *VerificationResult) T_TargetLocation(numTargetLocation int) templ.Component {

	if resource == nil || len(resource.TargetLocation) >= numTargetLocation {
		return StringInput("VerificationResult.TargetLocation["+strconv.Itoa(numTargetLocation)+"]", nil)
	}
	return StringInput("VerificationResult.TargetLocation["+strconv.Itoa(numTargetLocation)+"]", &resource.TargetLocation[numTargetLocation])
}
func (resource *VerificationResult) T_Need(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("VerificationResult.Need", nil, optionsValueSet)
	}
	return CodeableConceptSelect("VerificationResult.Need", resource.Need, optionsValueSet)
}
func (resource *VerificationResult) T_Status(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("VerificationResult.Status", nil, optionsValueSet)
	}
	return CodeSelect("VerificationResult.Status", &resource.Status, optionsValueSet)
}
func (resource *VerificationResult) T_StatusDate() templ.Component {

	if resource == nil {
		return StringInput("VerificationResult.StatusDate", nil)
	}
	return StringInput("VerificationResult.StatusDate", resource.StatusDate)
}
func (resource *VerificationResult) T_ValidationType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("VerificationResult.ValidationType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("VerificationResult.ValidationType", resource.ValidationType, optionsValueSet)
}
func (resource *VerificationResult) T_ValidationProcess(numValidationProcess int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ValidationProcess) >= numValidationProcess {
		return CodeableConceptSelect("VerificationResult.ValidationProcess["+strconv.Itoa(numValidationProcess)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("VerificationResult.ValidationProcess["+strconv.Itoa(numValidationProcess)+"]", &resource.ValidationProcess[numValidationProcess], optionsValueSet)
}
func (resource *VerificationResult) T_LastPerformed() templ.Component {

	if resource == nil {
		return StringInput("VerificationResult.LastPerformed", nil)
	}
	return StringInput("VerificationResult.LastPerformed", resource.LastPerformed)
}
func (resource *VerificationResult) T_NextScheduled() templ.Component {

	if resource == nil {
		return StringInput("VerificationResult.NextScheduled", nil)
	}
	return StringInput("VerificationResult.NextScheduled", resource.NextScheduled)
}
func (resource *VerificationResult) T_FailureAction(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("VerificationResult.FailureAction", nil, optionsValueSet)
	}
	return CodeableConceptSelect("VerificationResult.FailureAction", resource.FailureAction, optionsValueSet)
}
func (resource *VerificationResult) T_PrimarySourceId(numPrimarySource int) templ.Component {

	if resource == nil || len(resource.PrimarySource) >= numPrimarySource {
		return StringInput("VerificationResult.PrimarySource["+strconv.Itoa(numPrimarySource)+"].Id", nil)
	}
	return StringInput("VerificationResult.PrimarySource["+strconv.Itoa(numPrimarySource)+"].Id", resource.PrimarySource[numPrimarySource].Id)
}
func (resource *VerificationResult) T_PrimarySourceType(numPrimarySource int, numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PrimarySource) >= numPrimarySource || len(resource.PrimarySource[numPrimarySource].Type) >= numType {
		return CodeableConceptSelect("VerificationResult.PrimarySource["+strconv.Itoa(numPrimarySource)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("VerificationResult.PrimarySource["+strconv.Itoa(numPrimarySource)+"].Type["+strconv.Itoa(numType)+"]", &resource.PrimarySource[numPrimarySource].Type[numType], optionsValueSet)
}
func (resource *VerificationResult) T_PrimarySourceCommunicationMethod(numPrimarySource int, numCommunicationMethod int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PrimarySource) >= numPrimarySource || len(resource.PrimarySource[numPrimarySource].CommunicationMethod) >= numCommunicationMethod {
		return CodeableConceptSelect("VerificationResult.PrimarySource["+strconv.Itoa(numPrimarySource)+"].CommunicationMethod["+strconv.Itoa(numCommunicationMethod)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("VerificationResult.PrimarySource["+strconv.Itoa(numPrimarySource)+"].CommunicationMethod["+strconv.Itoa(numCommunicationMethod)+"]", &resource.PrimarySource[numPrimarySource].CommunicationMethod[numCommunicationMethod], optionsValueSet)
}
func (resource *VerificationResult) T_PrimarySourceValidationStatus(numPrimarySource int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PrimarySource) >= numPrimarySource {
		return CodeableConceptSelect("VerificationResult.PrimarySource["+strconv.Itoa(numPrimarySource)+"].ValidationStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("VerificationResult.PrimarySource["+strconv.Itoa(numPrimarySource)+"].ValidationStatus", resource.PrimarySource[numPrimarySource].ValidationStatus, optionsValueSet)
}
func (resource *VerificationResult) T_PrimarySourceValidationDate(numPrimarySource int) templ.Component {

	if resource == nil || len(resource.PrimarySource) >= numPrimarySource {
		return StringInput("VerificationResult.PrimarySource["+strconv.Itoa(numPrimarySource)+"].ValidationDate", nil)
	}
	return StringInput("VerificationResult.PrimarySource["+strconv.Itoa(numPrimarySource)+"].ValidationDate", resource.PrimarySource[numPrimarySource].ValidationDate)
}
func (resource *VerificationResult) T_PrimarySourceCanPushUpdates(numPrimarySource int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PrimarySource) >= numPrimarySource {
		return CodeableConceptSelect("VerificationResult.PrimarySource["+strconv.Itoa(numPrimarySource)+"].CanPushUpdates", nil, optionsValueSet)
	}
	return CodeableConceptSelect("VerificationResult.PrimarySource["+strconv.Itoa(numPrimarySource)+"].CanPushUpdates", resource.PrimarySource[numPrimarySource].CanPushUpdates, optionsValueSet)
}
func (resource *VerificationResult) T_PrimarySourcePushTypeAvailable(numPrimarySource int, numPushTypeAvailable int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PrimarySource) >= numPrimarySource || len(resource.PrimarySource[numPrimarySource].PushTypeAvailable) >= numPushTypeAvailable {
		return CodeableConceptSelect("VerificationResult.PrimarySource["+strconv.Itoa(numPrimarySource)+"].PushTypeAvailable["+strconv.Itoa(numPushTypeAvailable)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("VerificationResult.PrimarySource["+strconv.Itoa(numPrimarySource)+"].PushTypeAvailable["+strconv.Itoa(numPushTypeAvailable)+"]", &resource.PrimarySource[numPrimarySource].PushTypeAvailable[numPushTypeAvailable], optionsValueSet)
}
func (resource *VerificationResult) T_AttestationId() templ.Component {

	if resource == nil {
		return StringInput("VerificationResult.Attestation.Id", nil)
	}
	return StringInput("VerificationResult.Attestation.Id", resource.Attestation.Id)
}
func (resource *VerificationResult) T_AttestationCommunicationMethod(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("VerificationResult.Attestation.CommunicationMethod", nil, optionsValueSet)
	}
	return CodeableConceptSelect("VerificationResult.Attestation.CommunicationMethod", resource.Attestation.CommunicationMethod, optionsValueSet)
}
func (resource *VerificationResult) T_AttestationDate() templ.Component {

	if resource == nil {
		return StringInput("VerificationResult.Attestation.Date", nil)
	}
	return StringInput("VerificationResult.Attestation.Date", resource.Attestation.Date)
}
func (resource *VerificationResult) T_AttestationSourceIdentityCertificate() templ.Component {

	if resource == nil {
		return StringInput("VerificationResult.Attestation.SourceIdentityCertificate", nil)
	}
	return StringInput("VerificationResult.Attestation.SourceIdentityCertificate", resource.Attestation.SourceIdentityCertificate)
}
func (resource *VerificationResult) T_AttestationProxyIdentityCertificate() templ.Component {

	if resource == nil {
		return StringInput("VerificationResult.Attestation.ProxyIdentityCertificate", nil)
	}
	return StringInput("VerificationResult.Attestation.ProxyIdentityCertificate", resource.Attestation.ProxyIdentityCertificate)
}
func (resource *VerificationResult) T_ValidatorId(numValidator int) templ.Component {

	if resource == nil || len(resource.Validator) >= numValidator {
		return StringInput("VerificationResult.Validator["+strconv.Itoa(numValidator)+"].Id", nil)
	}
	return StringInput("VerificationResult.Validator["+strconv.Itoa(numValidator)+"].Id", resource.Validator[numValidator].Id)
}
func (resource *VerificationResult) T_ValidatorIdentityCertificate(numValidator int) templ.Component {

	if resource == nil || len(resource.Validator) >= numValidator {
		return StringInput("VerificationResult.Validator["+strconv.Itoa(numValidator)+"].IdentityCertificate", nil)
	}
	return StringInput("VerificationResult.Validator["+strconv.Itoa(numValidator)+"].IdentityCertificate", resource.Validator[numValidator].IdentityCertificate)
}
