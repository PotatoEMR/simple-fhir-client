package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	Created               string                     `json:"created"`
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
	WhenDateTime      string          `json:"whenDateTime"`
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
	ServicedDate            *string                        `json:"servicedDate"`
	ServicedPeriod          *Period                        `json:"servicedPeriod"`
	LocationCodeableConcept *CodeableConcept               `json:"locationCodeableConcept"`
	LocationAddress         *Address                       `json:"locationAddress"`
	LocationReference       *Reference                     `json:"locationReference"`
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
	Date              *string          `json:"date,omitempty"`
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

func (resource *ClaimResponse) ClaimResponseLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseStatus() templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Type, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseSubType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("subType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subType", resource.SubType, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseUse() templ.Component {
	optionsValueSet := VSClaim_use

	if resource == nil {
		return CodeSelect("use", nil, optionsValueSet)
	}
	return CodeSelect("use", &resource.Use, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseOutcome() templ.Component {
	optionsValueSet := VSClaim_outcome

	if resource == nil {
		return CodeSelect("outcome", nil, optionsValueSet)
	}
	return CodeSelect("outcome", &resource.Outcome, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseDecision(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("decision", nil, optionsValueSet)
	}
	return CodeableConceptSelect("decision", resource.Decision, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponsePayeeType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("payeeType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("payeeType", resource.PayeeType, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseDiagnosisRelatedGroup(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("diagnosisRelatedGroup", nil, optionsValueSet)
	}
	return CodeableConceptSelect("diagnosisRelatedGroup", resource.DiagnosisRelatedGroup, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseFundsReserve(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("fundsReserve", nil, optionsValueSet)
	}
	return CodeableConceptSelect("fundsReserve", resource.FundsReserve, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseFormCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("formCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("formCode", resource.FormCode, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseEventType(numEvent int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Event) >= numEvent {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Event[numEvent].Type, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseItemReviewOutcomeDecision(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("decision", nil, optionsValueSet)
	}
	return CodeableConceptSelect("decision", resource.Item[numItem].ReviewOutcome.Decision, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseItemReviewOutcomeReason(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reason", &resource.Item[numItem].ReviewOutcome.Reason[0], optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseItemAdjudicationCategory(numItem int, numAdjudication int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Adjudication) >= numAdjudication {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Item[numItem].Adjudication[numAdjudication].Category, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseItemAdjudicationReason(numItem int, numAdjudication int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Adjudication) >= numAdjudication {
		return CodeableConceptSelect("reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reason", resource.Item[numItem].Adjudication[numAdjudication].Reason, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemRevenue(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("revenue", resource.AddItem[numAddItem].Revenue, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemProductOrService(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", resource.AddItem[numAddItem].ProductOrService, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemProductOrServiceEnd(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("productOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrServiceEnd", resource.AddItem[numAddItem].ProductOrServiceEnd, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemModifier(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.AddItem[numAddItem].Modifier[0], optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemProgramCode(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("programCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("programCode", &resource.AddItem[numAddItem].ProgramCode[0], optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemBodySiteSubSite(numAddItem int, numBodySite int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].BodySite) >= numBodySite {
		return CodeableConceptSelect("subSite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subSite", &resource.AddItem[numAddItem].BodySite[numBodySite].SubSite[0], optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemDetailRevenue(numAddItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return CodeableConceptSelect("revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("revenue", resource.AddItem[numAddItem].Detail[numDetail].Revenue, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemDetailProductOrService(numAddItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", resource.AddItem[numAddItem].Detail[numDetail].ProductOrService, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemDetailProductOrServiceEnd(numAddItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return CodeableConceptSelect("productOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrServiceEnd", resource.AddItem[numAddItem].Detail[numDetail].ProductOrServiceEnd, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemDetailModifier(numAddItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.AddItem[numAddItem].Detail[numDetail].Modifier[0], optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemDetailSubDetailRevenue(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("revenue", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Revenue, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemDetailSubDetailProductOrService(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemDetailSubDetailProductOrServiceEnd(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("productOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrServiceEnd", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrServiceEnd, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemDetailSubDetailModifier(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[0], optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseTotalCategory(numTotal int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Total) >= numTotal {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Total[numTotal].Category, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponsePaymentType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Payment.Type, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponsePaymentAdjustmentReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("adjustmentReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("adjustmentReason", resource.Payment.AdjustmentReason, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseProcessNoteType(numProcessNote int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.ProcessNote) >= numProcessNote {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.ProcessNote[numProcessNote].Type, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseProcessNoteLanguage(numProcessNote int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.ProcessNote) >= numProcessNote {
		return CodeableConceptSelect("language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("language", resource.ProcessNote[numProcessNote].Language, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseErrorCode(numError int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Error) >= numError {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Error[numError].Code, optionsValueSet)
}
