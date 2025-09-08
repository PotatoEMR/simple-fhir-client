package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	Created               time.Time                  `json:"created,format:'2006-01-02T15:04:05Z07:00'"`
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
	WhenDateTime      time.Time       `json:"whenDateTime,format:'2006-01-02T15:04:05Z07:00'"`
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
	ServicedDate            *time.Time                     `json:"servicedDate,omitempty,format:'2006-01-02'"`
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
	Date              *time.Time       `json:"date,omitempty,format:'2006-01-02'"`
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
func (resource *ClaimResponse) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_SubType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SubType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubType", resource.SubType, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_Use(htmlAttrs string) templ.Component {
	optionsValueSet := VSClaim_use

	if resource == nil {
		return CodeSelect("Use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Use", &resource.Use, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_Created(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Created", nil, htmlAttrs)
	}
	return DateTimeInput("Created", &resource.Created, htmlAttrs)
}
func (resource *ClaimResponse) T_Outcome(htmlAttrs string) templ.Component {
	optionsValueSet := VSClaim_outcome

	if resource == nil {
		return CodeSelect("Outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Outcome", &resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_Decision(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Decision", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Decision", resource.Decision, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_Disposition(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Disposition", nil, htmlAttrs)
	}
	return StringInput("Disposition", resource.Disposition, htmlAttrs)
}
func (resource *ClaimResponse) T_PreAuthRef(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("PreAuthRef", nil, htmlAttrs)
	}
	return StringInput("PreAuthRef", resource.PreAuthRef, htmlAttrs)
}
func (resource *ClaimResponse) T_PayeeType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PayeeType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PayeeType", resource.PayeeType, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_DiagnosisRelatedGroup(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("DiagnosisRelatedGroup", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DiagnosisRelatedGroup", resource.DiagnosisRelatedGroup, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_FundsReserve(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("FundsReserve", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("FundsReserve", resource.FundsReserve, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_FormCode(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("FormCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("FormCode", resource.FormCode, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_EventType(numEvent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return CodeableConceptSelect("Event["+strconv.Itoa(numEvent)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Event["+strconv.Itoa(numEvent)+"]Type", &resource.Event[numEvent].Type, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_EventWhenDateTime(numEvent int, htmlAttrs string) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return DateTimeInput("Event["+strconv.Itoa(numEvent)+"]WhenDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Event["+strconv.Itoa(numEvent)+"]WhenDateTime", &resource.Event[numEvent].WhenDateTime, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemItemSequence(numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return IntInput("Item["+strconv.Itoa(numItem)+"]ItemSequence", nil, htmlAttrs)
	}
	return IntInput("Item["+strconv.Itoa(numItem)+"]ItemSequence", &resource.Item[numItem].ItemSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemNoteNumber(numItem int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numNoteNumber >= len(resource.Item[numItem].NoteNumber) {
		return IntInput("Item["+strconv.Itoa(numItem)+"]NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("Item["+strconv.Itoa(numItem)+"]NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_ItemReviewOutcomeDecision(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]ReviewOutcome.Decision", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]ReviewOutcome.Decision", resource.Item[numItem].ReviewOutcome.Decision, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemReviewOutcomeReason(numItem int, numReason int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numReason >= len(resource.Item[numItem].ReviewOutcome.Reason) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]ReviewOutcome.Reason["+strconv.Itoa(numReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]ReviewOutcome.Reason["+strconv.Itoa(numReason)+"]", &resource.Item[numItem].ReviewOutcome.Reason[numReason], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemReviewOutcomePreAuthRef(numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("Item["+strconv.Itoa(numItem)+"]ReviewOutcome.PreAuthRef", nil, htmlAttrs)
	}
	return StringInput("Item["+strconv.Itoa(numItem)+"]ReviewOutcome.PreAuthRef", resource.Item[numItem].ReviewOutcome.PreAuthRef, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemAdjudicationCategory(numItem int, numAdjudication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Adjudication["+strconv.Itoa(numAdjudication)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Adjudication["+strconv.Itoa(numAdjudication)+"].Category", &resource.Item[numItem].Adjudication[numAdjudication].Category, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemAdjudicationReason(numItem int, numAdjudication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Adjudication["+strconv.Itoa(numAdjudication)+"].Reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Adjudication["+strconv.Itoa(numAdjudication)+"].Reason", resource.Item[numItem].Adjudication[numAdjudication].Reason, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemDetailDetailSequence(numItem int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return IntInput("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].DetailSequence", nil, htmlAttrs)
	}
	return IntInput("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].DetailSequence", &resource.Item[numItem].Detail[numDetail].DetailSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemDetailNoteNumber(numItem int, numDetail int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numNoteNumber >= len(resource.Item[numItem].Detail[numDetail].NoteNumber) {
		return IntInput("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].Detail[numDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_ItemDetailSubDetailSubDetailSequence(numItem int, numDetail int, numSubDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return IntInput("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].SubDetailSequence", nil, htmlAttrs)
	}
	return IntInput("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].SubDetailSequence", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].SubDetailSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemDetailSubDetailNoteNumber(numItem int, numDetail int, numSubDetail int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numNoteNumber >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber) {
		return IntInput("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemItemSequence(numAddItem int, numItemSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numItemSequence >= len(resource.AddItem[numAddItem].ItemSequence) {
		return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]ItemSequence["+strconv.Itoa(numItemSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]ItemSequence["+strconv.Itoa(numItemSequence)+"]", &resource.AddItem[numAddItem].ItemSequence[numItemSequence], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSequence(numAddItem int, numDetailSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetailSequence >= len(resource.AddItem[numAddItem].DetailSequence) {
		return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]DetailSequence["+strconv.Itoa(numDetailSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]DetailSequence["+strconv.Itoa(numDetailSequence)+"]", &resource.AddItem[numAddItem].DetailSequence[numDetailSequence], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemSubdetailSequence(numAddItem int, numSubdetailSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numSubdetailSequence >= len(resource.AddItem[numAddItem].SubdetailSequence) {
		return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]SubdetailSequence["+strconv.Itoa(numSubdetailSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]SubdetailSequence["+strconv.Itoa(numSubdetailSequence)+"]", &resource.AddItem[numAddItem].SubdetailSequence[numSubdetailSequence], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemRevenue(numAddItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Revenue", resource.AddItem[numAddItem].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemProductOrService(numAddItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]ProductOrService", resource.AddItem[numAddItem].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemProductOrServiceEnd(numAddItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]ProductOrServiceEnd", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]ProductOrServiceEnd", resource.AddItem[numAddItem].ProductOrServiceEnd, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemModifier(numAddItem int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numModifier >= len(resource.AddItem[numAddItem].Modifier) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemProgramCode(numAddItem int, numProgramCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numProgramCode >= len(resource.AddItem[numAddItem].ProgramCode) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]ProgramCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]ProgramCode["+strconv.Itoa(numProgramCode)+"]", &resource.AddItem[numAddItem].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemServicedDate(numAddItem int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return DateInput("AddItem["+strconv.Itoa(numAddItem)+"]ServicedDate", nil, htmlAttrs)
	}
	return DateInput("AddItem["+strconv.Itoa(numAddItem)+"]ServicedDate", resource.AddItem[numAddItem].ServicedDate, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemLocationCodeableConcept(numAddItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]LocationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]LocationCodeableConcept", resource.AddItem[numAddItem].LocationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemFactor(numAddItem int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return Float64Input("AddItem["+strconv.Itoa(numAddItem)+"]Factor", nil, htmlAttrs)
	}
	return Float64Input("AddItem["+strconv.Itoa(numAddItem)+"]Factor", resource.AddItem[numAddItem].Factor, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemNoteNumber(numAddItem int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numNoteNumber >= len(resource.AddItem[numAddItem].NoteNumber) {
		return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemBodySiteSubSite(numAddItem int, numBodySite int, numSubSite int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numBodySite >= len(resource.AddItem[numAddItem].BodySite) || numSubSite >= len(resource.AddItem[numAddItem].BodySite[numBodySite].SubSite) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]BodySite["+strconv.Itoa(numBodySite)+"].SubSite["+strconv.Itoa(numSubSite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]BodySite["+strconv.Itoa(numBodySite)+"].SubSite["+strconv.Itoa(numSubSite)+"]", &resource.AddItem[numAddItem].BodySite[numBodySite].SubSite[numSubSite], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailRevenue(numAddItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].Revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].Revenue", resource.AddItem[numAddItem].Detail[numDetail].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailProductOrService(numAddItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].ProductOrService", resource.AddItem[numAddItem].Detail[numDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailProductOrServiceEnd(numAddItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].ProductOrServiceEnd", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].ProductOrServiceEnd", resource.AddItem[numAddItem].Detail[numDetail].ProductOrServiceEnd, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailModifier(numAddItem int, numDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numModifier >= len(resource.AddItem[numAddItem].Detail[numDetail].Modifier) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Detail[numDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailFactor(numAddItem int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return Float64Input("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].Factor", nil, htmlAttrs)
	}
	return Float64Input("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].Factor", resource.AddItem[numAddItem].Detail[numDetail].Factor, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailNoteNumber(numAddItem int, numDetail int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numNoteNumber >= len(resource.AddItem[numAddItem].Detail[numDetail].NoteNumber) {
		return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailRevenue(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Revenue", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailProductOrService(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailProductOrServiceEnd(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrServiceEnd", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrServiceEnd", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrServiceEnd, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailModifier(numAddItem int, numDetail int, numSubDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) || numModifier >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailFactor(numAddItem int, numDetail int, numSubDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return Float64Input("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Factor", nil, htmlAttrs)
	}
	return Float64Input("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Factor", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Factor, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailNoteNumber(numAddItem int, numDetail int, numSubDetail int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) || numNoteNumber >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber) {
		return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_TotalCategory(numTotal int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTotal >= len(resource.Total) {
		return CodeableConceptSelect("Total["+strconv.Itoa(numTotal)+"]Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Total["+strconv.Itoa(numTotal)+"]Category", &resource.Total[numTotal].Category, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_PaymentType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PaymentType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PaymentType", &resource.Payment.Type, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_PaymentAdjustmentReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PaymentAdjustmentReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PaymentAdjustmentReason", resource.Payment.AdjustmentReason, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_PaymentDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("PaymentDate", nil, htmlAttrs)
	}
	return DateInput("PaymentDate", resource.Payment.Date, htmlAttrs)
}
func (resource *ClaimResponse) T_ProcessNoteNumber(numProcessNote int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return IntInput("ProcessNote["+strconv.Itoa(numProcessNote)+"]Number", nil, htmlAttrs)
	}
	return IntInput("ProcessNote["+strconv.Itoa(numProcessNote)+"]Number", resource.ProcessNote[numProcessNote].Number, htmlAttrs)
}
func (resource *ClaimResponse) T_ProcessNoteType(numProcessNote int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return CodeableConceptSelect("ProcessNote["+strconv.Itoa(numProcessNote)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("ProcessNote["+strconv.Itoa(numProcessNote)+"]Type", resource.ProcessNote[numProcessNote].Type, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_ProcessNoteText(numProcessNote int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return StringInput("ProcessNote["+strconv.Itoa(numProcessNote)+"]Text", nil, htmlAttrs)
	}
	return StringInput("ProcessNote["+strconv.Itoa(numProcessNote)+"]Text", &resource.ProcessNote[numProcessNote].Text, htmlAttrs)
}
func (resource *ClaimResponse) T_InsuranceSequence(numInsurance int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return IntInput("Insurance["+strconv.Itoa(numInsurance)+"]Sequence", nil, htmlAttrs)
	}
	return IntInput("Insurance["+strconv.Itoa(numInsurance)+"]Sequence", &resource.Insurance[numInsurance].Sequence, htmlAttrs)
}
func (resource *ClaimResponse) T_InsuranceFocal(numInsurance int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return BoolInput("Insurance["+strconv.Itoa(numInsurance)+"]Focal", nil, htmlAttrs)
	}
	return BoolInput("Insurance["+strconv.Itoa(numInsurance)+"]Focal", &resource.Insurance[numInsurance].Focal, htmlAttrs)
}
func (resource *ClaimResponse) T_InsuranceBusinessArrangement(numInsurance int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return StringInput("Insurance["+strconv.Itoa(numInsurance)+"]BusinessArrangement", nil, htmlAttrs)
	}
	return StringInput("Insurance["+strconv.Itoa(numInsurance)+"]BusinessArrangement", resource.Insurance[numInsurance].BusinessArrangement, htmlAttrs)
}
func (resource *ClaimResponse) T_ErrorItemSequence(numError int, htmlAttrs string) templ.Component {
	if resource == nil || numError >= len(resource.Error) {
		return IntInput("Error["+strconv.Itoa(numError)+"]ItemSequence", nil, htmlAttrs)
	}
	return IntInput("Error["+strconv.Itoa(numError)+"]ItemSequence", resource.Error[numError].ItemSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ErrorDetailSequence(numError int, htmlAttrs string) templ.Component {
	if resource == nil || numError >= len(resource.Error) {
		return IntInput("Error["+strconv.Itoa(numError)+"]DetailSequence", nil, htmlAttrs)
	}
	return IntInput("Error["+strconv.Itoa(numError)+"]DetailSequence", resource.Error[numError].DetailSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ErrorSubDetailSequence(numError int, htmlAttrs string) templ.Component {
	if resource == nil || numError >= len(resource.Error) {
		return IntInput("Error["+strconv.Itoa(numError)+"]SubDetailSequence", nil, htmlAttrs)
	}
	return IntInput("Error["+strconv.Itoa(numError)+"]SubDetailSequence", resource.Error[numError].SubDetailSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ErrorCode(numError int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numError >= len(resource.Error) {
		return CodeableConceptSelect("Error["+strconv.Itoa(numError)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Error["+strconv.Itoa(numError)+"]Code", &resource.Error[numError].Code, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_ErrorExpression(numError int, numExpression int, htmlAttrs string) templ.Component {
	if resource == nil || numError >= len(resource.Error) || numExpression >= len(resource.Error[numError].Expression) {
		return StringInput("Error["+strconv.Itoa(numError)+"]Expression["+strconv.Itoa(numExpression)+"]", nil, htmlAttrs)
	}
	return StringInput("Error["+strconv.Itoa(numError)+"]Expression["+strconv.Itoa(numExpression)+"]", &resource.Error[numError].Expression[numExpression], htmlAttrs)
}
