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

// http://hl7.org/fhir/r4b/StructureDefinition/ClaimResponse
type ClaimResponse struct {
	Id                   *string                    `json:"id,omitempty"`
	Meta                 *Meta                      `json:"meta,omitempty"`
	ImplicitRules        *string                    `json:"implicitRules,omitempty"`
	Language             *string                    `json:"language,omitempty"`
	Text                 *Narrative                 `json:"text,omitempty"`
	Contained            []Resource                 `json:"contained,omitempty"`
	Extension            []Extension                `json:"extension,omitempty"`
	ModifierExtension    []Extension                `json:"modifierExtension,omitempty"`
	Identifier           []Identifier               `json:"identifier,omitempty"`
	Status               string                     `json:"status"`
	Type                 CodeableConcept            `json:"type"`
	SubType              *CodeableConcept           `json:"subType,omitempty"`
	Use                  string                     `json:"use"`
	Patient              Reference                  `json:"patient"`
	Created              time.Time                  `json:"created,format:'2006-01-02T15:04:05Z07:00'"`
	Insurer              Reference                  `json:"insurer"`
	Requestor            *Reference                 `json:"requestor,omitempty"`
	Request              *Reference                 `json:"request,omitempty"`
	Outcome              string                     `json:"outcome"`
	Disposition          *string                    `json:"disposition,omitempty"`
	PreAuthRef           *string                    `json:"preAuthRef,omitempty"`
	PreAuthPeriod        *Period                    `json:"preAuthPeriod,omitempty"`
	PayeeType            *CodeableConcept           `json:"payeeType,omitempty"`
	Item                 []ClaimResponseItem        `json:"item,omitempty"`
	AddItem              []ClaimResponseAddItem     `json:"addItem,omitempty"`
	Total                []ClaimResponseTotal       `json:"total,omitempty"`
	Payment              *ClaimResponsePayment      `json:"payment,omitempty"`
	FundsReserve         *CodeableConcept           `json:"fundsReserve,omitempty"`
	FormCode             *CodeableConcept           `json:"formCode,omitempty"`
	Form                 *Attachment                `json:"form,omitempty"`
	ProcessNote          []ClaimResponseProcessNote `json:"processNote,omitempty"`
	CommunicationRequest []Reference                `json:"communicationRequest,omitempty"`
	Insurance            []ClaimResponseInsurance   `json:"insurance,omitempty"`
	Error                []ClaimResponseError       `json:"error,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClaimResponse
type ClaimResponseItem struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	ItemSequence      int                             `json:"itemSequence"`
	NoteNumber        []int                           `json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `json:"adjudication"`
	Detail            []ClaimResponseItemDetail       `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClaimResponse
type ClaimResponseItemAdjudication struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Category          CodeableConcept  `json:"category"`
	Reason            *CodeableConcept `json:"reason,omitempty"`
	Amount            *Money           `json:"amount,omitempty"`
	Value             *float64         `json:"value,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClaimResponse
type ClaimResponseItemDetail struct {
	Id                *string                            `json:"id,omitempty"`
	Extension         []Extension                        `json:"extension,omitempty"`
	ModifierExtension []Extension                        `json:"modifierExtension,omitempty"`
	DetailSequence    int                                `json:"detailSequence"`
	NoteNumber        []int                              `json:"noteNumber,omitempty"`
	SubDetail         []ClaimResponseItemDetailSubDetail `json:"subDetail,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClaimResponse
type ClaimResponseItemDetailSubDetail struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	SubDetailSequence int         `json:"subDetailSequence"`
	NoteNumber        []int       `json:"noteNumber,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClaimResponse
type ClaimResponseAddItem struct {
	Id                      *string                      `json:"id,omitempty"`
	Extension               []Extension                  `json:"extension,omitempty"`
	ModifierExtension       []Extension                  `json:"modifierExtension,omitempty"`
	ItemSequence            []int                        `json:"itemSequence,omitempty"`
	DetailSequence          []int                        `json:"detailSequence,omitempty"`
	SubdetailSequence       []int                        `json:"subdetailSequence,omitempty"`
	Provider                []Reference                  `json:"provider,omitempty"`
	ProductOrService        CodeableConcept              `json:"productOrService"`
	Modifier                []CodeableConcept            `json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept            `json:"programCode,omitempty"`
	ServicedDate            *time.Time                   `json:"servicedDate,omitempty,format:'2006-01-02'"`
	ServicedPeriod          *Period                      `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept             `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                     `json:"locationAddress,omitempty"`
	LocationReference       *Reference                   `json:"locationReference,omitempty"`
	Quantity                *Quantity                    `json:"quantity,omitempty"`
	UnitPrice               *Money                       `json:"unitPrice,omitempty"`
	Factor                  *float64                     `json:"factor,omitempty"`
	Net                     *Money                       `json:"net,omitempty"`
	BodySite                *CodeableConcept             `json:"bodySite,omitempty"`
	SubSite                 []CodeableConcept            `json:"subSite,omitempty"`
	NoteNumber              []int                        `json:"noteNumber,omitempty"`
	Detail                  []ClaimResponseAddItemDetail `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClaimResponse
type ClaimResponseAddItemDetail struct {
	Id                *string                               `json:"id,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	ProductOrService  CodeableConcept                       `json:"productOrService"`
	Modifier          []CodeableConcept                     `json:"modifier,omitempty"`
	Quantity          *Quantity                             `json:"quantity,omitempty"`
	UnitPrice         *Money                                `json:"unitPrice,omitempty"`
	Factor            *float64                              `json:"factor,omitempty"`
	Net               *Money                                `json:"net,omitempty"`
	NoteNumber        []int                                 `json:"noteNumber,omitempty"`
	SubDetail         []ClaimResponseAddItemDetailSubDetail `json:"subDetail,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClaimResponse
type ClaimResponseAddItemDetailSubDetail struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	ProductOrService  CodeableConcept   `json:"productOrService"`
	Modifier          []CodeableConcept `json:"modifier,omitempty"`
	Quantity          *Quantity         `json:"quantity,omitempty"`
	UnitPrice         *Money            `json:"unitPrice,omitempty"`
	Factor            *float64          `json:"factor,omitempty"`
	Net               *Money            `json:"net,omitempty"`
	NoteNumber        []int             `json:"noteNumber,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClaimResponse
type ClaimResponseTotal struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Category          CodeableConcept `json:"category"`
	Amount            Money           `json:"amount"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClaimResponse
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

// http://hl7.org/fhir/r4b/StructureDefinition/ClaimResponse
type ClaimResponseProcessNote struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Number            *int             `json:"number,omitempty"`
	Type              *string          `json:"type,omitempty"`
	Text              string           `json:"text"`
	Language          *CodeableConcept `json:"language,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ClaimResponse
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

// http://hl7.org/fhir/r4b/StructureDefinition/ClaimResponse
type ClaimResponseError struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	ItemSequence      *int            `json:"itemSequence,omitempty"`
	DetailSequence    *int            `json:"detailSequence,omitempty"`
	SubDetailSequence *int            `json:"subDetailSequence,omitempty"`
	Code              CodeableConcept `json:"code"`
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
	optionsValueSet := VSRemittance_outcome

	if resource == nil {
		return CodeSelect("Outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Outcome", &resource.Outcome, optionsValueSet, htmlAttrs)
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
func (resource *ClaimResponse) T_ItemAdjudicationValue(numItem int, numAdjudication int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return Float64Input("Item["+strconv.Itoa(numItem)+"]Adjudication["+strconv.Itoa(numAdjudication)+"].Value", nil, htmlAttrs)
	}
	return Float64Input("Item["+strconv.Itoa(numItem)+"]Adjudication["+strconv.Itoa(numAdjudication)+"].Value", resource.Item[numItem].Adjudication[numAdjudication].Value, htmlAttrs)
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
func (resource *ClaimResponse) T_AddItemProductOrService(numAddItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]ProductOrService", &resource.AddItem[numAddItem].ProductOrService, optionsValueSet, htmlAttrs)
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
func (resource *ClaimResponse) T_AddItemBodySite(numAddItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]BodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]BodySite", resource.AddItem[numAddItem].BodySite, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemSubSite(numAddItem int, numSubSite int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numSubSite >= len(resource.AddItem[numAddItem].SubSite) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]SubSite["+strconv.Itoa(numSubSite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]SubSite["+strconv.Itoa(numSubSite)+"]", &resource.AddItem[numAddItem].SubSite[numSubSite], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemNoteNumber(numAddItem int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numNoteNumber >= len(resource.AddItem[numAddItem].NoteNumber) {
		return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailProductOrService(numAddItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].ProductOrService", &resource.AddItem[numAddItem].Detail[numDetail].ProductOrService, optionsValueSet, htmlAttrs)
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
func (resource *ClaimResponse) T_AddItemDetailSubDetailProductOrService(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet, htmlAttrs)
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
func (resource *ClaimResponse) T_ProcessNoteType(numProcessNote int, htmlAttrs string) templ.Component {
	optionsValueSet := VSNote_type

	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return CodeSelect("ProcessNote["+strconv.Itoa(numProcessNote)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ProcessNote["+strconv.Itoa(numProcessNote)+"]Type", resource.ProcessNote[numProcessNote].Type, optionsValueSet, htmlAttrs)
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
