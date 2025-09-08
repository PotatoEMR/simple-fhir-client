package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/VerificationResult
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
	StatusDate        *time.Time                        `json:"statusDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	ValidationType    *CodeableConcept                  `json:"validationType,omitempty"`
	ValidationProcess []CodeableConcept                 `json:"validationProcess,omitempty"`
	Frequency         *Timing                           `json:"frequency,omitempty"`
	LastPerformed     *time.Time                        `json:"lastPerformed,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	NextScheduled     *time.Time                        `json:"nextScheduled,omitempty,format:'2006-01-02'"`
	FailureAction     *CodeableConcept                  `json:"failureAction,omitempty"`
	PrimarySource     []VerificationResultPrimarySource `json:"primarySource,omitempty"`
	Attestation       *VerificationResultAttestation    `json:"attestation,omitempty"`
	Validator         []VerificationResultValidator     `json:"validator,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/VerificationResult
type VerificationResultPrimarySource struct {
	Id                  *string           `json:"id,omitempty"`
	Extension           []Extension       `json:"extension,omitempty"`
	ModifierExtension   []Extension       `json:"modifierExtension,omitempty"`
	Who                 *Reference        `json:"who,omitempty"`
	Type                []CodeableConcept `json:"type,omitempty"`
	CommunicationMethod []CodeableConcept `json:"communicationMethod,omitempty"`
	ValidationStatus    *CodeableConcept  `json:"validationStatus,omitempty"`
	ValidationDate      *time.Time        `json:"validationDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	CanPushUpdates      *CodeableConcept  `json:"canPushUpdates,omitempty"`
	PushTypeAvailable   []CodeableConcept `json:"pushTypeAvailable,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/VerificationResult
type VerificationResultAttestation struct {
	Id                        *string          `json:"id,omitempty"`
	Extension                 []Extension      `json:"extension,omitempty"`
	ModifierExtension         []Extension      `json:"modifierExtension,omitempty"`
	Who                       *Reference       `json:"who,omitempty"`
	OnBehalfOf                *Reference       `json:"onBehalfOf,omitempty"`
	CommunicationMethod       *CodeableConcept `json:"communicationMethod,omitempty"`
	Date                      *time.Time       `json:"date,omitempty,format:'2006-01-02'"`
	SourceIdentityCertificate *string          `json:"sourceIdentityCertificate,omitempty"`
	ProxyIdentityCertificate  *string          `json:"proxyIdentityCertificate,omitempty"`
	ProxySignature            *Signature       `json:"proxySignature,omitempty"`
	SourceSignature           *Signature       `json:"sourceSignature,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/VerificationResult
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
func (resource *VerificationResult) T_TargetLocation(numTargetLocation int, htmlAttrs string) templ.Component {

	if resource == nil || numTargetLocation >= len(resource.TargetLocation) {
		return StringInput("VerificationResult.TargetLocation."+strconv.Itoa(numTargetLocation)+".", nil, htmlAttrs)
	}
	return StringInput("VerificationResult.TargetLocation."+strconv.Itoa(numTargetLocation)+".", &resource.TargetLocation[numTargetLocation], htmlAttrs)
}
func (resource *VerificationResult) T_Need(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("VerificationResult.Need", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("VerificationResult.Need", resource.Need, optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSVerificationresult_status

	if resource == nil {
		return CodeSelect("VerificationResult.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("VerificationResult.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_StatusDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("VerificationResult.StatusDate", nil, htmlAttrs)
	}
	return DateTimeInput("VerificationResult.StatusDate", resource.StatusDate, htmlAttrs)
}
func (resource *VerificationResult) T_ValidationType(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("VerificationResult.ValidationType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("VerificationResult.ValidationType", resource.ValidationType, optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_ValidationProcess(numValidationProcess int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numValidationProcess >= len(resource.ValidationProcess) {
		return CodeableConceptSelect("VerificationResult.ValidationProcess."+strconv.Itoa(numValidationProcess)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("VerificationResult.ValidationProcess."+strconv.Itoa(numValidationProcess)+".", &resource.ValidationProcess[numValidationProcess], optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_LastPerformed(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("VerificationResult.LastPerformed", nil, htmlAttrs)
	}
	return DateTimeInput("VerificationResult.LastPerformed", resource.LastPerformed, htmlAttrs)
}
func (resource *VerificationResult) T_NextScheduled(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("VerificationResult.NextScheduled", nil, htmlAttrs)
	}
	return DateInput("VerificationResult.NextScheduled", resource.NextScheduled, htmlAttrs)
}
func (resource *VerificationResult) T_FailureAction(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("VerificationResult.FailureAction", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("VerificationResult.FailureAction", resource.FailureAction, optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_PrimarySourceType(numPrimarySource int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numPrimarySource >= len(resource.PrimarySource) || numType >= len(resource.PrimarySource[numPrimarySource].Type) {
		return CodeableConceptSelect("VerificationResult.PrimarySource."+strconv.Itoa(numPrimarySource)+"..Type."+strconv.Itoa(numType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("VerificationResult.PrimarySource."+strconv.Itoa(numPrimarySource)+"..Type."+strconv.Itoa(numType)+".", &resource.PrimarySource[numPrimarySource].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_PrimarySourceCommunicationMethod(numPrimarySource int, numCommunicationMethod int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numPrimarySource >= len(resource.PrimarySource) || numCommunicationMethod >= len(resource.PrimarySource[numPrimarySource].CommunicationMethod) {
		return CodeableConceptSelect("VerificationResult.PrimarySource."+strconv.Itoa(numPrimarySource)+"..CommunicationMethod."+strconv.Itoa(numCommunicationMethod)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("VerificationResult.PrimarySource."+strconv.Itoa(numPrimarySource)+"..CommunicationMethod."+strconv.Itoa(numCommunicationMethod)+".", &resource.PrimarySource[numPrimarySource].CommunicationMethod[numCommunicationMethod], optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_PrimarySourceValidationStatus(numPrimarySource int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numPrimarySource >= len(resource.PrimarySource) {
		return CodeableConceptSelect("VerificationResult.PrimarySource."+strconv.Itoa(numPrimarySource)+"..ValidationStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("VerificationResult.PrimarySource."+strconv.Itoa(numPrimarySource)+"..ValidationStatus", resource.PrimarySource[numPrimarySource].ValidationStatus, optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_PrimarySourceValidationDate(numPrimarySource int, htmlAttrs string) templ.Component {

	if resource == nil || numPrimarySource >= len(resource.PrimarySource) {
		return DateTimeInput("VerificationResult.PrimarySource."+strconv.Itoa(numPrimarySource)+"..ValidationDate", nil, htmlAttrs)
	}
	return DateTimeInput("VerificationResult.PrimarySource."+strconv.Itoa(numPrimarySource)+"..ValidationDate", resource.PrimarySource[numPrimarySource].ValidationDate, htmlAttrs)
}
func (resource *VerificationResult) T_PrimarySourceCanPushUpdates(numPrimarySource int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numPrimarySource >= len(resource.PrimarySource) {
		return CodeableConceptSelect("VerificationResult.PrimarySource."+strconv.Itoa(numPrimarySource)+"..CanPushUpdates", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("VerificationResult.PrimarySource."+strconv.Itoa(numPrimarySource)+"..CanPushUpdates", resource.PrimarySource[numPrimarySource].CanPushUpdates, optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_PrimarySourcePushTypeAvailable(numPrimarySource int, numPushTypeAvailable int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numPrimarySource >= len(resource.PrimarySource) || numPushTypeAvailable >= len(resource.PrimarySource[numPrimarySource].PushTypeAvailable) {
		return CodeableConceptSelect("VerificationResult.PrimarySource."+strconv.Itoa(numPrimarySource)+"..PushTypeAvailable."+strconv.Itoa(numPushTypeAvailable)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("VerificationResult.PrimarySource."+strconv.Itoa(numPrimarySource)+"..PushTypeAvailable."+strconv.Itoa(numPushTypeAvailable)+".", &resource.PrimarySource[numPrimarySource].PushTypeAvailable[numPushTypeAvailable], optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_AttestationCommunicationMethod(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("VerificationResult.Attestation.CommunicationMethod", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("VerificationResult.Attestation.CommunicationMethod", resource.Attestation.CommunicationMethod, optionsValueSet, htmlAttrs)
}
func (resource *VerificationResult) T_AttestationDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateInput("VerificationResult.Attestation.Date", nil, htmlAttrs)
	}
	return DateInput("VerificationResult.Attestation.Date", resource.Attestation.Date, htmlAttrs)
}
func (resource *VerificationResult) T_AttestationSourceIdentityCertificate(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("VerificationResult.Attestation.SourceIdentityCertificate", nil, htmlAttrs)
	}
	return StringInput("VerificationResult.Attestation.SourceIdentityCertificate", resource.Attestation.SourceIdentityCertificate, htmlAttrs)
}
func (resource *VerificationResult) T_AttestationProxyIdentityCertificate(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("VerificationResult.Attestation.ProxyIdentityCertificate", nil, htmlAttrs)
	}
	return StringInput("VerificationResult.Attestation.ProxyIdentityCertificate", resource.Attestation.ProxyIdentityCertificate, htmlAttrs)
}
func (resource *VerificationResult) T_ValidatorIdentityCertificate(numValidator int, htmlAttrs string) templ.Component {

	if resource == nil || numValidator >= len(resource.Validator) {
		return StringInput("VerificationResult.Validator."+strconv.Itoa(numValidator)+"..IdentityCertificate", nil, htmlAttrs)
	}
	return StringInput("VerificationResult.Validator."+strconv.Itoa(numValidator)+"..IdentityCertificate", resource.Validator[numValidator].IdentityCertificate, htmlAttrs)
}
