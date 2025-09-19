package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ClaimResponse
type ClaimResponse struct {
	Id                    *string                    `json:"id,omitempty"`
	Meta                  *Meta                      `json:"meta,omitempty"`
	ImplicitRules         *string                    `json:"implicitRules,omitempty"`
	Language              *string                    `json:"language,omitempty"`
	Text                  *Narrative                 `json:"text,omitempty"`
	Contained             []Resource                 `json:"contained,omitempty"`
	Extension             []Extension                `json:"extension,omitempty"`
	ModifierExtension     []Extension                `json:"modifierExtension,omitempty"`
	Identifier            []Identifier               `json:"identifier,omitempty"`
	TraceNumber           []Identifier               `json:"traceNumber,omitempty"`
	Status                string                     `json:"status"`
	Type                  CodeableConcept            `json:"type"`
	SubType               *CodeableConcept           `json:"subType,omitempty"`
	Use                   string                     `json:"use"`
	Patient               Reference                  `json:"patient"`
	Created               FhirDateTime               `json:"created"`
	Insurer               *Reference                 `json:"insurer,omitempty"`
	Requestor             *Reference                 `json:"requestor,omitempty"`
	Request               *Reference                 `json:"request,omitempty"`
	Outcome               string                     `json:"outcome"`
	Decision              *CodeableConcept           `json:"decision,omitempty"`
	Disposition           *string                    `json:"disposition,omitempty"`
	PreAuthRef            *string                    `json:"preAuthRef,omitempty"`
	PreAuthPeriod         *Period                    `json:"preAuthPeriod,omitempty"`
	Event                 []ClaimResponseEvent       `json:"event,omitempty"`
	PayeeType             *CodeableConcept           `json:"payeeType,omitempty"`
	Encounter             []Reference                `json:"encounter,omitempty"`
	DiagnosisRelatedGroup *CodeableConcept           `json:"diagnosisRelatedGroup,omitempty"`
	Item                  []ClaimResponseItem        `json:"item,omitempty"`
	AddItem               []ClaimResponseAddItem     `json:"addItem,omitempty"`
	Total                 []ClaimResponseTotal       `json:"total,omitempty"`
	Payment               *ClaimResponsePayment      `json:"payment,omitempty"`
	FundsReserve          *CodeableConcept           `json:"fundsReserve,omitempty"`
	FormCode              *CodeableConcept           `json:"formCode,omitempty"`
	Form                  *Attachment                `json:"form,omitempty"`
	ProcessNote           []ClaimResponseProcessNote `json:"processNote,omitempty"`
	CommunicationRequest  []Reference                `json:"communicationRequest,omitempty"`
	Insurance             []ClaimResponseInsurance   `json:"insurance,omitempty"`
	Error                 []ClaimResponseError       `json:"error,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClaimResponse
type ClaimResponseEvent struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	WhenDateTime      FhirDateTime    `json:"whenDateTime"`
	WhenPeriod        Period          `json:"whenPeriod"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClaimResponse
type ClaimResponseItem struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	ItemSequence      int                             `json:"itemSequence"`
	TraceNumber       []Identifier                    `json:"traceNumber,omitempty"`
	NoteNumber        []int                           `json:"noteNumber,omitempty"`
	ReviewOutcome     *ClaimResponseItemReviewOutcome `json:"reviewOutcome,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`
	Detail            []ClaimResponseItemDetail       `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClaimResponse
type ClaimResponseItemReviewOutcome struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Decision          *CodeableConcept  `json:"decision,omitempty"`
	Reason            []CodeableConcept `json:"reason,omitempty"`
	PreAuthRef        *string           `json:"preAuthRef,omitempty"`
	PreAuthPeriod     *Period           `json:"preAuthPeriod,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClaimResponse
type ClaimResponseItemAdjudication struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Category          CodeableConcept  `json:"category"`
	Reason            *CodeableConcept `json:"reason,omitempty"`
	Amount            *Money           `json:"amount,omitempty"`
	Quantity          *Quantity        `json:"quantity,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClaimResponse
type ClaimResponseItemDetail struct {
	Id                *string                            `json:"id,omitempty"`
	Extension         []Extension                        `json:"extension,omitempty"`
	ModifierExtension []Extension                        `json:"modifierExtension,omitempty"`
	DetailSequence    int                                `json:"detailSequence"`
	TraceNumber       []Identifier                       `json:"traceNumber,omitempty"`
	NoteNumber        []int                              `json:"noteNumber,omitempty"`
	SubDetail         []ClaimResponseItemDetailSubDetail `json:"subDetail,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClaimResponse
type ClaimResponseItemDetailSubDetail struct {
	Id                *string      `json:"id,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	SubDetailSequence int          `json:"subDetailSequence"`
	TraceNumber       []Identifier `json:"traceNumber,omitempty"`
	NoteNumber        []int        `json:"noteNumber,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClaimResponse
type ClaimResponseAddItem struct {
	Id                      *string                        `json:"id,omitempty"`
	Extension               []Extension                    `json:"extension,omitempty"`
	ModifierExtension       []Extension                    `json:"modifierExtension,omitempty"`
	ItemSequence            []int                          `json:"itemSequence,omitempty"`
	DetailSequence          []int                          `json:"detailSequence,omitempty"`
	SubdetailSequence       []int                          `json:"subdetailSequence,omitempty"`
	TraceNumber             []Identifier                   `json:"traceNumber,omitempty"`
	Provider                []Reference                    `json:"provider,omitempty"`
	Revenue                 *CodeableConcept               `json:"revenue,omitempty"`
	ProductOrService        *CodeableConcept               `json:"productOrService,omitempty"`
	ProductOrServiceEnd     *CodeableConcept               `json:"productOrServiceEnd,omitempty"`
	Request                 []Reference                    `json:"request,omitempty"`
	Modifier                []CodeableConcept              `json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept              `json:"programCode,omitempty"`
	ServicedDate            *FhirDate                      `json:"servicedDate,omitempty"`
	ServicedPeriod          *Period                        `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept               `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                       `json:"locationAddress,omitempty"`
	LocationReference       *Reference                     `json:"locationReference,omitempty"`
	Quantity                *Quantity                      `json:"quantity,omitempty"`
	UnitPrice               *Money                         `json:"unitPrice,omitempty"`
	Factor                  *float64                       `json:"factor,omitempty"`
	Tax                     *Money                         `json:"tax,omitempty"`
	Net                     *Money                         `json:"net,omitempty"`
	BodySite                []ClaimResponseAddItemBodySite `json:"bodySite,omitempty"`
	NoteNumber              []int                          `json:"noteNumber,omitempty"`
	Detail                  []ClaimResponseAddItemDetail   `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClaimResponse
type ClaimResponseAddItemBodySite struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Site              []CodeableReference `json:"site"`
	SubSite           []CodeableConcept   `json:"subSite,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClaimResponse
type ClaimResponseAddItemDetail struct {
	Id                  *string                               `json:"id,omitempty"`
	Extension           []Extension                           `json:"extension,omitempty"`
	ModifierExtension   []Extension                           `json:"modifierExtension,omitempty"`
	TraceNumber         []Identifier                          `json:"traceNumber,omitempty"`
	Revenue             *CodeableConcept                      `json:"revenue,omitempty"`
	ProductOrService    *CodeableConcept                      `json:"productOrService,omitempty"`
	ProductOrServiceEnd *CodeableConcept                      `json:"productOrServiceEnd,omitempty"`
	Modifier            []CodeableConcept                     `json:"modifier,omitempty"`
	Quantity            *Quantity                             `json:"quantity,omitempty"`
	UnitPrice           *Money                                `json:"unitPrice,omitempty"`
	Factor              *float64                              `json:"factor,omitempty"`
	Tax                 *Money                                `json:"tax,omitempty"`
	Net                 *Money                                `json:"net,omitempty"`
	NoteNumber          []int                                 `json:"noteNumber,omitempty"`
	SubDetail           []ClaimResponseAddItemDetailSubDetail `json:"subDetail,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClaimResponse
type ClaimResponseAddItemDetailSubDetail struct {
	Id                  *string           `json:"id,omitempty"`
	Extension           []Extension       `json:"extension,omitempty"`
	ModifierExtension   []Extension       `json:"modifierExtension,omitempty"`
	TraceNumber         []Identifier      `json:"traceNumber,omitempty"`
	Revenue             *CodeableConcept  `json:"revenue,omitempty"`
	ProductOrService    *CodeableConcept  `json:"productOrService,omitempty"`
	ProductOrServiceEnd *CodeableConcept  `json:"productOrServiceEnd,omitempty"`
	Modifier            []CodeableConcept `json:"modifier,omitempty"`
	Quantity            *Quantity         `json:"quantity,omitempty"`
	UnitPrice           *Money            `json:"unitPrice,omitempty"`
	Factor              *float64          `json:"factor,omitempty"`
	Tax                 *Money            `json:"tax,omitempty"`
	Net                 *Money            `json:"net,omitempty"`
	NoteNumber          []int             `json:"noteNumber,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClaimResponse
type ClaimResponseTotal struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Category          CodeableConcept `json:"category"`
	Amount            Money           `json:"amount"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClaimResponse
type ClaimResponsePayment struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              CodeableConcept  `json:"type"`
	Adjustment        *Money           `json:"adjustment,omitempty"`
	AdjustmentReason  *CodeableConcept `json:"adjustmentReason,omitempty"`
	Date              *FhirDate        `json:"date,omitempty"`
	Amount            Money            `json:"amount"`
	Identifier        *Identifier      `json:"identifier,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClaimResponse
type ClaimResponseProcessNote struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Number            *int             `json:"number,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Text              string           `json:"text"`
	Language          *CodeableConcept `json:"language,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClaimResponse
type ClaimResponseInsurance struct {
	Id                  *string     `json:"id,omitempty"`
	Extension           []Extension `json:"extension,omitempty"`
	ModifierExtension   []Extension `json:"modifierExtension,omitempty"`
	Sequence            int         `json:"sequence"`
	Focal               bool        `json:"focal"`
	Coverage            Reference   `json:"coverage"`
	BusinessArrangement *string     `json:"businessArrangement,omitempty"`
	ClaimResponse       *Reference  `json:"claimResponse,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ClaimResponse
type ClaimResponseError struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	ItemSequence      *int            `json:"itemSequence,omitempty"`
	DetailSequence    *int            `json:"detailSequence,omitempty"`
	SubDetailSequence *int            `json:"subDetailSequence,omitempty"`
	Code              CodeableConcept `json:"code"`
	Expression        []string        `json:"expression,omitempty"`
}

type OtherClaimResponse ClaimResponse

// on convert struct to json, automatically add resourceType=ClaimResponse
func (r ClaimResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClaimResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherClaimResponse: OtherClaimResponse(r),
		ResourceType:       "ClaimResponse",
	})
}
func (r ClaimResponse) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ClaimResponse/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ClaimResponse"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ClaimResponse) T_TraceNumber(numTraceNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTraceNumber >= len(resource.TraceNumber) {
		return IdentifierInput("traceNumber["+strconv.Itoa(numTraceNumber)+"]", nil, htmlAttrs)
	}
	return IdentifierInput("traceNumber["+strconv.Itoa(numTraceNumber)+"]", &resource.TraceNumber[numTraceNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_SubType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("subType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subType", resource.SubType, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_Use(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSClaim_use

	if resource == nil {
		return CodeSelect("use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("use", &resource.Use, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_Patient(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("patient", nil, htmlAttrs)
	}
	return ReferenceInput("patient", &resource.Patient, htmlAttrs)
}
func (resource *ClaimResponse) T_Created(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("created", nil, htmlAttrs)
	}
	return FhirDateTimeInput("created", &resource.Created, htmlAttrs)
}
func (resource *ClaimResponse) T_Insurer(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("insurer", nil, htmlAttrs)
	}
	return ReferenceInput("insurer", resource.Insurer, htmlAttrs)
}
func (resource *ClaimResponse) T_Requestor(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("requestor", nil, htmlAttrs)
	}
	return ReferenceInput("requestor", resource.Requestor, htmlAttrs)
}
func (resource *ClaimResponse) T_Request(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("request", nil, htmlAttrs)
	}
	return ReferenceInput("request", resource.Request, htmlAttrs)
}
func (resource *ClaimResponse) T_Outcome(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSClaim_outcome

	if resource == nil {
		return CodeSelect("outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("outcome", &resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_Decision(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("decision", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("decision", resource.Decision, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_Disposition(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("disposition", nil, htmlAttrs)
	}
	return StringInput("disposition", resource.Disposition, htmlAttrs)
}
func (resource *ClaimResponse) T_PreAuthRef(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("preAuthRef", nil, htmlAttrs)
	}
	return StringInput("preAuthRef", resource.PreAuthRef, htmlAttrs)
}
func (resource *ClaimResponse) T_PreAuthPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("preAuthPeriod", nil, htmlAttrs)
	}
	return PeriodInput("preAuthPeriod", resource.PreAuthPeriod, htmlAttrs)
}
func (resource *ClaimResponse) T_PayeeType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("payeeType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payeeType", resource.PayeeType, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_Encounter(numEncounter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEncounter >= len(resource.Encounter) {
		return ReferenceInput("encounter["+strconv.Itoa(numEncounter)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("encounter["+strconv.Itoa(numEncounter)+"]", &resource.Encounter[numEncounter], htmlAttrs)
}
func (resource *ClaimResponse) T_DiagnosisRelatedGroup(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("diagnosisRelatedGroup", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosisRelatedGroup", resource.DiagnosisRelatedGroup, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_FundsReserve(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("fundsReserve", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("fundsReserve", resource.FundsReserve, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_FormCode(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("formCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("formCode", resource.FormCode, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_Form(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return AttachmentInput("form", nil, htmlAttrs)
	}
	return AttachmentInput("form", resource.Form, htmlAttrs)
}
func (resource *ClaimResponse) T_CommunicationRequest(numCommunicationRequest int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCommunicationRequest >= len(resource.CommunicationRequest) {
		return ReferenceInput("communicationRequest["+strconv.Itoa(numCommunicationRequest)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("communicationRequest["+strconv.Itoa(numCommunicationRequest)+"]", &resource.CommunicationRequest[numCommunicationRequest], htmlAttrs)
}
func (resource *ClaimResponse) T_EventType(numEvent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return CodeableConceptSelect("event["+strconv.Itoa(numEvent)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("event["+strconv.Itoa(numEvent)+"].type", &resource.Event[numEvent].Type, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_EventWhenDateTime(numEvent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return FhirDateTimeInput("event["+strconv.Itoa(numEvent)+"].whenDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("event["+strconv.Itoa(numEvent)+"].whenDateTime", &resource.Event[numEvent].WhenDateTime, htmlAttrs)
}
func (resource *ClaimResponse) T_EventWhenPeriod(numEvent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return PeriodInput("event["+strconv.Itoa(numEvent)+"].whenPeriod", nil, htmlAttrs)
	}
	return PeriodInput("event["+strconv.Itoa(numEvent)+"].whenPeriod", &resource.Event[numEvent].WhenPeriod, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemItemSequence(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return IntInput("item["+strconv.Itoa(numItem)+"].itemSequence", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].itemSequence", &resource.Item[numItem].ItemSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemTraceNumber(numItem int, numTraceNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numTraceNumber >= len(resource.Item[numItem].TraceNumber) {
		return IdentifierInput("item["+strconv.Itoa(numItem)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", nil, htmlAttrs)
	}
	return IdentifierInput("item["+strconv.Itoa(numItem)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", &resource.Item[numItem].TraceNumber[numTraceNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_ItemNoteNumber(numItem int, numNoteNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numNoteNumber >= len(resource.Item[numItem].NoteNumber) {
		return IntInput("item["+strconv.Itoa(numItem)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_ItemReviewOutcomeDecision(numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].reviewOutcome.decision", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].reviewOutcome.decision", resource.Item[numItem].ReviewOutcome.Decision, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemReviewOutcomeReason(numItem int, numReason int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numReason >= len(resource.Item[numItem].ReviewOutcome.Reason) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].reviewOutcome.reason["+strconv.Itoa(numReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].reviewOutcome.reason["+strconv.Itoa(numReason)+"]", &resource.Item[numItem].ReviewOutcome.Reason[numReason], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemReviewOutcomePreAuthRef(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("item["+strconv.Itoa(numItem)+"].reviewOutcome.preAuthRef", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].reviewOutcome.preAuthRef", resource.Item[numItem].ReviewOutcome.PreAuthRef, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemReviewOutcomePreAuthPeriod(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return PeriodInput("item["+strconv.Itoa(numItem)+"].reviewOutcome.preAuthPeriod", nil, htmlAttrs)
	}
	return PeriodInput("item["+strconv.Itoa(numItem)+"].reviewOutcome.preAuthPeriod", resource.Item[numItem].ReviewOutcome.PreAuthPeriod, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemAdjudicationCategory(numItem int, numAdjudication int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].category", &resource.Item[numItem].Adjudication[numAdjudication].Category, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemAdjudicationReason(numItem int, numAdjudication int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].reason", resource.Item[numItem].Adjudication[numAdjudication].Reason, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemAdjudicationAmount(numItem int, numAdjudication int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].amount", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].amount", resource.Item[numItem].Adjudication[numAdjudication].Amount, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemAdjudicationQuantity(numItem int, numAdjudication int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return QuantityInput("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].quantity", nil, htmlAttrs)
	}
	return QuantityInput("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].quantity", resource.Item[numItem].Adjudication[numAdjudication].Quantity, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemDetailDetailSequence(numItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].detailSequence", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].detailSequence", &resource.Item[numItem].Detail[numDetail].DetailSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemDetailTraceNumber(numItem int, numDetail int, numTraceNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numTraceNumber >= len(resource.Item[numItem].Detail[numDetail].TraceNumber) {
		return IdentifierInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", nil, htmlAttrs)
	}
	return IdentifierInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", &resource.Item[numItem].Detail[numDetail].TraceNumber[numTraceNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_ItemDetailNoteNumber(numItem int, numDetail int, numNoteNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numNoteNumber >= len(resource.Item[numItem].Detail[numDetail].NoteNumber) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].Detail[numDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_ItemDetailSubDetailSubDetailSequence(numItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].subDetailSequence", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].subDetailSequence", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].SubDetailSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemDetailSubDetailTraceNumber(numItem int, numDetail int, numSubDetail int, numTraceNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numTraceNumber >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].TraceNumber) {
		return IdentifierInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", nil, htmlAttrs)
	}
	return IdentifierInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].TraceNumber[numTraceNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_ItemDetailSubDetailNoteNumber(numItem int, numDetail int, numSubDetail int, numNoteNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numNoteNumber >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemItemSequence(numAddItem int, numItemSequence int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numItemSequence >= len(resource.AddItem[numAddItem].ItemSequence) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].itemSequence["+strconv.Itoa(numItemSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].itemSequence["+strconv.Itoa(numItemSequence)+"]", &resource.AddItem[numAddItem].ItemSequence[numItemSequence], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSequence(numAddItem int, numDetailSequence int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetailSequence >= len(resource.AddItem[numAddItem].DetailSequence) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detailSequence["+strconv.Itoa(numDetailSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detailSequence["+strconv.Itoa(numDetailSequence)+"]", &resource.AddItem[numAddItem].DetailSequence[numDetailSequence], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemSubdetailSequence(numAddItem int, numSubdetailSequence int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numSubdetailSequence >= len(resource.AddItem[numAddItem].SubdetailSequence) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].subdetailSequence["+strconv.Itoa(numSubdetailSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].subdetailSequence["+strconv.Itoa(numSubdetailSequence)+"]", &resource.AddItem[numAddItem].SubdetailSequence[numSubdetailSequence], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemTraceNumber(numAddItem int, numTraceNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numTraceNumber >= len(resource.AddItem[numAddItem].TraceNumber) {
		return IdentifierInput("addItem["+strconv.Itoa(numAddItem)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", nil, htmlAttrs)
	}
	return IdentifierInput("addItem["+strconv.Itoa(numAddItem)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", &resource.AddItem[numAddItem].TraceNumber[numTraceNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemProvider(numAddItem int, numProvider int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numProvider >= len(resource.AddItem[numAddItem].Provider) {
		return ReferenceInput("addItem["+strconv.Itoa(numAddItem)+"].provider["+strconv.Itoa(numProvider)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("addItem["+strconv.Itoa(numAddItem)+"].provider["+strconv.Itoa(numProvider)+"]", &resource.AddItem[numAddItem].Provider[numProvider], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemRevenue(numAddItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].revenue", resource.AddItem[numAddItem].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemProductOrService(numAddItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].productOrService", resource.AddItem[numAddItem].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemProductOrServiceEnd(numAddItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].productOrServiceEnd", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].productOrServiceEnd", resource.AddItem[numAddItem].ProductOrServiceEnd, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemRequest(numAddItem int, numRequest int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numRequest >= len(resource.AddItem[numAddItem].Request) {
		return ReferenceInput("addItem["+strconv.Itoa(numAddItem)+"].request["+strconv.Itoa(numRequest)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("addItem["+strconv.Itoa(numAddItem)+"].request["+strconv.Itoa(numRequest)+"]", &resource.AddItem[numAddItem].Request[numRequest], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemModifier(numAddItem int, numModifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numModifier >= len(resource.AddItem[numAddItem].Modifier) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemProgramCode(numAddItem int, numProgramCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numProgramCode >= len(resource.AddItem[numAddItem].ProgramCode) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].programCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].programCode["+strconv.Itoa(numProgramCode)+"]", &resource.AddItem[numAddItem].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemServicedDate(numAddItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return FhirDateInput("addItem["+strconv.Itoa(numAddItem)+"].servicedDate", nil, htmlAttrs)
	}
	return FhirDateInput("addItem["+strconv.Itoa(numAddItem)+"].servicedDate", resource.AddItem[numAddItem].ServicedDate, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemServicedPeriod(numAddItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return PeriodInput("addItem["+strconv.Itoa(numAddItem)+"].servicedPeriod", nil, htmlAttrs)
	}
	return PeriodInput("addItem["+strconv.Itoa(numAddItem)+"].servicedPeriod", resource.AddItem[numAddItem].ServicedPeriod, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemLocationCodeableConcept(numAddItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].locationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].locationCodeableConcept", resource.AddItem[numAddItem].LocationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemLocationAddress(numAddItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return AddressInput("addItem["+strconv.Itoa(numAddItem)+"].locationAddress", nil, htmlAttrs)
	}
	return AddressInput("addItem["+strconv.Itoa(numAddItem)+"].locationAddress", resource.AddItem[numAddItem].LocationAddress, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemLocationReference(numAddItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return ReferenceInput("addItem["+strconv.Itoa(numAddItem)+"].locationReference", nil, htmlAttrs)
	}
	return ReferenceInput("addItem["+strconv.Itoa(numAddItem)+"].locationReference", resource.AddItem[numAddItem].LocationReference, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemQuantity(numAddItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return QuantityInput("addItem["+strconv.Itoa(numAddItem)+"].quantity", nil, htmlAttrs)
	}
	return QuantityInput("addItem["+strconv.Itoa(numAddItem)+"].quantity", resource.AddItem[numAddItem].Quantity, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemUnitPrice(numAddItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].unitPrice", nil, htmlAttrs)
	}
	return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].unitPrice", resource.AddItem[numAddItem].UnitPrice, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemFactor(numAddItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].factor", resource.AddItem[numAddItem].Factor, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemTax(numAddItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].tax", nil, htmlAttrs)
	}
	return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].tax", resource.AddItem[numAddItem].Tax, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemNet(numAddItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].net", nil, htmlAttrs)
	}
	return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].net", resource.AddItem[numAddItem].Net, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemNoteNumber(numAddItem int, numNoteNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numNoteNumber >= len(resource.AddItem[numAddItem].NoteNumber) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemBodySiteSite(numAddItem int, numBodySite int, numSite int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numBodySite >= len(resource.AddItem[numAddItem].BodySite) || numSite >= len(resource.AddItem[numAddItem].BodySite[numBodySite].Site) {
		return CodeableReferenceInput("addItem["+strconv.Itoa(numAddItem)+"].bodySite["+strconv.Itoa(numBodySite)+"].site["+strconv.Itoa(numSite)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("addItem["+strconv.Itoa(numAddItem)+"].bodySite["+strconv.Itoa(numBodySite)+"].site["+strconv.Itoa(numSite)+"]", &resource.AddItem[numAddItem].BodySite[numBodySite].Site[numSite], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemBodySiteSubSite(numAddItem int, numBodySite int, numSubSite int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numBodySite >= len(resource.AddItem[numAddItem].BodySite) || numSubSite >= len(resource.AddItem[numAddItem].BodySite[numBodySite].SubSite) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].bodySite["+strconv.Itoa(numBodySite)+"].subSite["+strconv.Itoa(numSubSite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].bodySite["+strconv.Itoa(numBodySite)+"].subSite["+strconv.Itoa(numSubSite)+"]", &resource.AddItem[numAddItem].BodySite[numBodySite].SubSite[numSubSite], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailTraceNumber(numAddItem int, numDetail int, numTraceNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numTraceNumber >= len(resource.AddItem[numAddItem].Detail[numDetail].TraceNumber) {
		return IdentifierInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", nil, htmlAttrs)
	}
	return IdentifierInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].TraceNumber[numTraceNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailRevenue(numAddItem int, numDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].revenue", resource.AddItem[numAddItem].Detail[numDetail].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailProductOrService(numAddItem int, numDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrService", resource.AddItem[numAddItem].Detail[numDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailProductOrServiceEnd(numAddItem int, numDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrServiceEnd", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrServiceEnd", resource.AddItem[numAddItem].Detail[numDetail].ProductOrServiceEnd, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailModifier(numAddItem int, numDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numModifier >= len(resource.AddItem[numAddItem].Detail[numDetail].Modifier) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Detail[numDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailQuantity(numAddItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return QuantityInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].quantity", nil, htmlAttrs)
	}
	return QuantityInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].quantity", resource.AddItem[numAddItem].Detail[numDetail].Quantity, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailUnitPrice(numAddItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].unitPrice", nil, htmlAttrs)
	}
	return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].unitPrice", resource.AddItem[numAddItem].Detail[numDetail].UnitPrice, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailFactor(numAddItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].factor", resource.AddItem[numAddItem].Detail[numDetail].Factor, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailTax(numAddItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].tax", nil, htmlAttrs)
	}
	return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].tax", resource.AddItem[numAddItem].Detail[numDetail].Tax, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailNet(numAddItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].net", nil, htmlAttrs)
	}
	return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].net", resource.AddItem[numAddItem].Detail[numDetail].Net, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailNoteNumber(numAddItem int, numDetail int, numNoteNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numNoteNumber >= len(resource.AddItem[numAddItem].Detail[numDetail].NoteNumber) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailTraceNumber(numAddItem int, numDetail int, numSubDetail int, numTraceNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) || numTraceNumber >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].TraceNumber) {
		return IdentifierInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", nil, htmlAttrs)
	}
	return IdentifierInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].TraceNumber[numTraceNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailRevenue(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].revenue", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailProductOrService(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrService", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailProductOrServiceEnd(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrServiceEnd", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrServiceEnd", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrServiceEnd, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailModifier(numAddItem int, numDetail int, numSubDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) || numModifier >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailQuantity(numAddItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return QuantityInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].quantity", nil, htmlAttrs)
	}
	return QuantityInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].quantity", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Quantity, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailUnitPrice(numAddItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].unitPrice", nil, htmlAttrs)
	}
	return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].unitPrice", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].UnitPrice, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailFactor(numAddItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].factor", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Factor, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailTax(numAddItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].tax", nil, htmlAttrs)
	}
	return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].tax", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Tax, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailNet(numAddItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].net", nil, htmlAttrs)
	}
	return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].net", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Net, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailNoteNumber(numAddItem int, numDetail int, numSubDetail int, numNoteNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) || numNoteNumber >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_TotalCategory(numTotal int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTotal >= len(resource.Total) {
		return CodeableConceptSelect("total["+strconv.Itoa(numTotal)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("total["+strconv.Itoa(numTotal)+"].category", &resource.Total[numTotal].Category, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_TotalAmount(numTotal int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTotal >= len(resource.Total) {
		return MoneyInput("total["+strconv.Itoa(numTotal)+"].amount", nil, htmlAttrs)
	}
	return MoneyInput("total["+strconv.Itoa(numTotal)+"].amount", &resource.Total[numTotal].Amount, htmlAttrs)
}
func (resource *ClaimResponse) T_PaymentType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("payment.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payment.type", &resource.Payment.Type, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_PaymentAdjustment(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return MoneyInput("payment.adjustment", nil, htmlAttrs)
	}
	return MoneyInput("payment.adjustment", resource.Payment.Adjustment, htmlAttrs)
}
func (resource *ClaimResponse) T_PaymentAdjustmentReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("payment.adjustmentReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payment.adjustmentReason", resource.Payment.AdjustmentReason, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_PaymentDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("payment.date", nil, htmlAttrs)
	}
	return FhirDateInput("payment.date", resource.Payment.Date, htmlAttrs)
}
func (resource *ClaimResponse) T_PaymentAmount(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return MoneyInput("payment.amount", nil, htmlAttrs)
	}
	return MoneyInput("payment.amount", &resource.Payment.Amount, htmlAttrs)
}
func (resource *ClaimResponse) T_ProcessNoteNumber(numProcessNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return IntInput("processNote["+strconv.Itoa(numProcessNote)+"].number", nil, htmlAttrs)
	}
	return IntInput("processNote["+strconv.Itoa(numProcessNote)+"].number", resource.ProcessNote[numProcessNote].Number, htmlAttrs)
}
func (resource *ClaimResponse) T_ProcessNoteType(numProcessNote int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return CodeableConceptSelect("processNote["+strconv.Itoa(numProcessNote)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("processNote["+strconv.Itoa(numProcessNote)+"].type", resource.ProcessNote[numProcessNote].Type, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_InsuranceSequence(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return IntInput("insurance["+strconv.Itoa(numInsurance)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("insurance["+strconv.Itoa(numInsurance)+"].sequence", &resource.Insurance[numInsurance].Sequence, htmlAttrs)
}
func (resource *ClaimResponse) T_InsuranceFocal(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].focal", nil, htmlAttrs)
	}
	return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].focal", &resource.Insurance[numInsurance].Focal, htmlAttrs)
}
func (resource *ClaimResponse) T_InsuranceCoverage(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return ReferenceInput("insurance["+strconv.Itoa(numInsurance)+"].coverage", nil, htmlAttrs)
	}
	return ReferenceInput("insurance["+strconv.Itoa(numInsurance)+"].coverage", &resource.Insurance[numInsurance].Coverage, htmlAttrs)
}
func (resource *ClaimResponse) T_InsuranceBusinessArrangement(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return StringInput("insurance["+strconv.Itoa(numInsurance)+"].businessArrangement", nil, htmlAttrs)
	}
	return StringInput("insurance["+strconv.Itoa(numInsurance)+"].businessArrangement", resource.Insurance[numInsurance].BusinessArrangement, htmlAttrs)
}
func (resource *ClaimResponse) T_InsuranceClaimResponse(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return ReferenceInput("insurance["+strconv.Itoa(numInsurance)+"].claimResponse", nil, htmlAttrs)
	}
	return ReferenceInput("insurance["+strconv.Itoa(numInsurance)+"].claimResponse", resource.Insurance[numInsurance].ClaimResponse, htmlAttrs)
}
func (resource *ClaimResponse) T_ErrorItemSequence(numError int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numError >= len(resource.Error) {
		return IntInput("error["+strconv.Itoa(numError)+"].itemSequence", nil, htmlAttrs)
	}
	return IntInput("error["+strconv.Itoa(numError)+"].itemSequence", resource.Error[numError].ItemSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ErrorDetailSequence(numError int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numError >= len(resource.Error) {
		return IntInput("error["+strconv.Itoa(numError)+"].detailSequence", nil, htmlAttrs)
	}
	return IntInput("error["+strconv.Itoa(numError)+"].detailSequence", resource.Error[numError].DetailSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ErrorSubDetailSequence(numError int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numError >= len(resource.Error) {
		return IntInput("error["+strconv.Itoa(numError)+"].subDetailSequence", nil, htmlAttrs)
	}
	return IntInput("error["+strconv.Itoa(numError)+"].subDetailSequence", resource.Error[numError].SubDetailSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ErrorCode(numError int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numError >= len(resource.Error) {
		return CodeableConceptSelect("error["+strconv.Itoa(numError)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("error["+strconv.Itoa(numError)+"].code", &resource.Error[numError].Code, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_ErrorExpression(numError int, numExpression int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numError >= len(resource.Error) || numExpression >= len(resource.Error[numError].Expression) {
		return StringInput("error["+strconv.Itoa(numError)+"].expression["+strconv.Itoa(numExpression)+"]", nil, htmlAttrs)
	}
	return StringInput("error["+strconv.Itoa(numError)+"].expression["+strconv.Itoa(numExpression)+"]", &resource.Error[numError].Expression[numExpression], htmlAttrs)
}
