package r4

//generated with command go run ./bultaoreune
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
	StatusDate        *FhirDateTime                     `json:"statusDate,omitempty"`
	ValidationType    *CodeableConcept                  `json:"validationType,omitempty"`
	ValidationProcess []CodeableConcept                 `json:"validationProcess,omitempty"`
	Frequency         *Timing                           `json:"frequency,omitempty"`
	LastPerformed     *FhirDateTime                     `json:"lastPerformed,omitempty"`
	NextScheduled     *FhirDate                         `json:"nextScheduled,omitempty"`
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
	ValidationDate      *FhirDateTime     `json:"validationDate,omitempty"`
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
	Date                      *FhirDate        `json:"date,omitempty"`
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

// struct -> json, automatically add resourceType=Patient
func (r VerificationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherVerificationResult
		ResourceType string `json:"resourceType"`
	}{
		OtherVerificationResult: OtherVerificationResult(r),
		ResourceType:            "VerificationResult",
	})
}

func (r VerificationResult) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "VerificationResult/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "VerificationResult"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (r VerificationResult) ResourceType() string {
	return "VerificationResult"
}

func (resource *VerificationResult) T_Target(frs []FhirResource, numTarget int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTarget >= len(resource.Target) {
		return ReferenceInput(frs, "target["+strconv.Itoa(numTarget)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "target["+strconv.Itoa(numTarget)+"]", &resource.Target[numTarget], htmlAttrs)
}
func (resource *VerificationResult) T_TargetLocation(numTargetLocation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTargetLocation >= len(resource.TargetLocation) {
		return StringInput("targetLocation["+strconv.Itoa(numTargetLocation)+"]", nil, htmlAttrs)
	}
	return StringInput("targetLocation["+strconv.Itoa(numTargetLocation)+"]", &resource.TargetLocation[numTargetLocation], htmlAttrs)
}
func (resource *VerificationResult) T_Need(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("need", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("need", resource.Need, optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_Status(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_StatusDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("statusDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("statusDate", resource.StatusDate, htmlAttrs)
}
func (resource *VerificationResult) T_ValidationType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("validationType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("validationType", resource.ValidationType, optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_ValidationProcess(numValidationProcess int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numValidationProcess >= len(resource.ValidationProcess) {
		return CodeableConceptSelect("validationProcess["+strconv.Itoa(numValidationProcess)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("validationProcess["+strconv.Itoa(numValidationProcess)+"]", &resource.ValidationProcess[numValidationProcess], optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_Frequency(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return TimingInput("frequency", nil, htmlAttrs)
	}
	return TimingInput("frequency", resource.Frequency, htmlAttrs)
}
func (resource *VerificationResult) T_LastPerformed(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("lastPerformed", nil, htmlAttrs)
	}
	return FhirDateTimeInput("lastPerformed", resource.LastPerformed, htmlAttrs)
}
func (resource *VerificationResult) T_NextScheduled(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("nextScheduled", nil, htmlAttrs)
	}
	return FhirDateInput("nextScheduled", resource.NextScheduled, htmlAttrs)
}
func (resource *VerificationResult) T_FailureAction(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("failureAction", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("failureAction", resource.FailureAction, optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_PrimarySourceWho(frs []FhirResource, numPrimarySource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrimarySource >= len(resource.PrimarySource) {
		return ReferenceInput(frs, "primarySource["+strconv.Itoa(numPrimarySource)+"].who", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "primarySource["+strconv.Itoa(numPrimarySource)+"].who", resource.PrimarySource[numPrimarySource].Who, htmlAttrs)
}
func (resource *VerificationResult) T_PrimarySourceType(numPrimarySource int, numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrimarySource >= len(resource.PrimarySource) || numType >= len(resource.PrimarySource[numPrimarySource].Type) {
		return CodeableConceptSelect("primarySource["+strconv.Itoa(numPrimarySource)+"].type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("primarySource["+strconv.Itoa(numPrimarySource)+"].type["+strconv.Itoa(numType)+"]", &resource.PrimarySource[numPrimarySource].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_PrimarySourceCommunicationMethod(numPrimarySource int, numCommunicationMethod int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrimarySource >= len(resource.PrimarySource) || numCommunicationMethod >= len(resource.PrimarySource[numPrimarySource].CommunicationMethod) {
		return CodeableConceptSelect("primarySource["+strconv.Itoa(numPrimarySource)+"].communicationMethod["+strconv.Itoa(numCommunicationMethod)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("primarySource["+strconv.Itoa(numPrimarySource)+"].communicationMethod["+strconv.Itoa(numCommunicationMethod)+"]", &resource.PrimarySource[numPrimarySource].CommunicationMethod[numCommunicationMethod], optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_PrimarySourceValidationStatus(numPrimarySource int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrimarySource >= len(resource.PrimarySource) {
		return CodeableConceptSelect("primarySource["+strconv.Itoa(numPrimarySource)+"].validationStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("primarySource["+strconv.Itoa(numPrimarySource)+"].validationStatus", resource.PrimarySource[numPrimarySource].ValidationStatus, optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_PrimarySourceValidationDate(numPrimarySource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrimarySource >= len(resource.PrimarySource) {
		return FhirDateTimeInput("primarySource["+strconv.Itoa(numPrimarySource)+"].validationDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("primarySource["+strconv.Itoa(numPrimarySource)+"].validationDate", resource.PrimarySource[numPrimarySource].ValidationDate, htmlAttrs)
}
func (resource *VerificationResult) T_PrimarySourceCanPushUpdates(numPrimarySource int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrimarySource >= len(resource.PrimarySource) {
		return CodeableConceptSelect("primarySource["+strconv.Itoa(numPrimarySource)+"].canPushUpdates", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("primarySource["+strconv.Itoa(numPrimarySource)+"].canPushUpdates", resource.PrimarySource[numPrimarySource].CanPushUpdates, optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_PrimarySourcePushTypeAvailable(numPrimarySource int, numPushTypeAvailable int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPrimarySource >= len(resource.PrimarySource) || numPushTypeAvailable >= len(resource.PrimarySource[numPrimarySource].PushTypeAvailable) {
		return CodeableConceptSelect("primarySource["+strconv.Itoa(numPrimarySource)+"].pushTypeAvailable["+strconv.Itoa(numPushTypeAvailable)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("primarySource["+strconv.Itoa(numPrimarySource)+"].pushTypeAvailable["+strconv.Itoa(numPushTypeAvailable)+"]", &resource.PrimarySource[numPrimarySource].PushTypeAvailable[numPushTypeAvailable], optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_AttestationWho(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "attestation.who", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "attestation.who", resource.Attestation.Who, htmlAttrs)
}
func (resource *VerificationResult) T_AttestationOnBehalfOf(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "attestation.onBehalfOf", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "attestation.onBehalfOf", resource.Attestation.OnBehalfOf, htmlAttrs)
}
func (resource *VerificationResult) T_AttestationCommunicationMethod(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("attestation.communicationMethod", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("attestation.communicationMethod", resource.Attestation.CommunicationMethod, optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_AttestationDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("attestation.date", nil, htmlAttrs)
	}
	return FhirDateInput("attestation.date", resource.Attestation.Date, htmlAttrs)
}
func (resource *VerificationResult) T_AttestationSourceIdentityCertificate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("attestation.sourceIdentityCertificate", nil, htmlAttrs)
	}
	return StringInput("attestation.sourceIdentityCertificate", resource.Attestation.SourceIdentityCertificate, htmlAttrs)
}
func (resource *VerificationResult) T_AttestationProxyIdentityCertificate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("attestation.proxyIdentityCertificate", nil, htmlAttrs)
	}
	return StringInput("attestation.proxyIdentityCertificate", resource.Attestation.ProxyIdentityCertificate, htmlAttrs)
}
func (resource *VerificationResult) T_AttestationProxySignature(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return SignatureInput("attestation.proxySignature", nil, htmlAttrs)
	}
	return SignatureInput("attestation.proxySignature", resource.Attestation.ProxySignature, htmlAttrs)
}
func (resource *VerificationResult) T_AttestationSourceSignature(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return SignatureInput("attestation.sourceSignature", nil, htmlAttrs)
	}
	return SignatureInput("attestation.sourceSignature", resource.Attestation.SourceSignature, htmlAttrs)
}
func (resource *VerificationResult) T_ValidatorOrganization(frs []FhirResource, numValidator int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numValidator >= len(resource.Validator) {
		return ReferenceInput(frs, "validator["+strconv.Itoa(numValidator)+"].organization", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "validator["+strconv.Itoa(numValidator)+"].organization", &resource.Validator[numValidator].Organization, htmlAttrs)
}
func (resource *VerificationResult) T_ValidatorIdentityCertificate(numValidator int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numValidator >= len(resource.Validator) {
		return StringInput("validator["+strconv.Itoa(numValidator)+"].identityCertificate", nil, htmlAttrs)
	}
	return StringInput("validator["+strconv.Itoa(numValidator)+"].identityCertificate", resource.Validator[numValidator].IdentityCertificate, htmlAttrs)
}
func (resource *VerificationResult) T_ValidatorAttestationSignature(numValidator int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numValidator >= len(resource.Validator) {
		return SignatureInput("validator["+strconv.Itoa(numValidator)+"].attestationSignature", nil, htmlAttrs)
	}
	return SignatureInput("validator["+strconv.Itoa(numValidator)+"].attestationSignature", resource.Validator[numValidator].AttestationSignature, htmlAttrs)
}
