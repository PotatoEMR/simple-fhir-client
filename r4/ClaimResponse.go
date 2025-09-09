package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/ClaimResponse
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
	Created              string                     `json:"created"`
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

// http://hl7.org/fhir/r4/StructureDefinition/ClaimResponse
type ClaimResponseItem struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	ItemSequence      int                             `json:"itemSequence"`
	NoteNumber        []int                           `json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `json:"adjudication"`
	Detail            []ClaimResponseItemDetail       `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ClaimResponse
type ClaimResponseItemAdjudication struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Category          CodeableConcept  `json:"category"`
	Reason            *CodeableConcept `json:"reason,omitempty"`
	Amount            *Money           `json:"amount,omitempty"`
	Value             *float64         `json:"value,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ClaimResponse
type ClaimResponseItemDetail struct {
	Id                *string                            `json:"id,omitempty"`
	Extension         []Extension                        `json:"extension,omitempty"`
	ModifierExtension []Extension                        `json:"modifierExtension,omitempty"`
	DetailSequence    int                                `json:"detailSequence"`
	NoteNumber        []int                              `json:"noteNumber,omitempty"`
	SubDetail         []ClaimResponseItemDetailSubDetail `json:"subDetail,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ClaimResponse
type ClaimResponseItemDetailSubDetail struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	SubDetailSequence int         `json:"subDetailSequence"`
	NoteNumber        []int       `json:"noteNumber,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ClaimResponse
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
	ServicedDate            *string                      `json:"servicedDate,omitempty"`
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

// http://hl7.org/fhir/r4/StructureDefinition/ClaimResponse
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

// http://hl7.org/fhir/r4/StructureDefinition/ClaimResponse
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

// http://hl7.org/fhir/r4/StructureDefinition/ClaimResponse
type ClaimResponseTotal struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Category          CodeableConcept `json:"category"`
	Amount            Money           `json:"amount"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ClaimResponse
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

// http://hl7.org/fhir/r4/StructureDefinition/ClaimResponse
type ClaimResponseProcessNote struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Number            *int             `json:"number,omitempty"`
	Type              *string          `json:"type,omitempty"`
	Text              string           `json:"text"`
	Language          *CodeableConcept `json:"language,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ClaimResponse
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

// http://hl7.org/fhir/r4/StructureDefinition/ClaimResponse
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
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_SubType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("subType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subType", resource.SubType, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_Use(htmlAttrs string) templ.Component {
	optionsValueSet := VSClaim_use

	if resource == nil {
		return CodeSelect("use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("use", &resource.Use, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_Created(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("created", nil, htmlAttrs)
	}
	return DateTimeInput("created", &resource.Created, htmlAttrs)
}
func (resource *ClaimResponse) T_Outcome(htmlAttrs string) templ.Component {
	optionsValueSet := VSRemittance_outcome

	if resource == nil {
		return CodeSelect("outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("outcome", &resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_Disposition(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("disposition", nil, htmlAttrs)
	}
	return StringInput("disposition", resource.Disposition, htmlAttrs)
}
func (resource *ClaimResponse) T_PreAuthRef(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("preAuthRef", nil, htmlAttrs)
	}
	return StringInput("preAuthRef", resource.PreAuthRef, htmlAttrs)
}
func (resource *ClaimResponse) T_PayeeType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("payeeType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payeeType", resource.PayeeType, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_FundsReserve(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("fundsReserve", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("fundsReserve", resource.FundsReserve, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_FormCode(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("formCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("formCode", resource.FormCode, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemItemSequence(numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return IntInput("item["+strconv.Itoa(numItem)+"].itemSequence", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].itemSequence", &resource.Item[numItem].ItemSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemNoteNumber(numItem int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numNoteNumber >= len(resource.Item[numItem].NoteNumber) {
		return IntInput("item["+strconv.Itoa(numItem)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_ItemAdjudicationCategory(numItem int, numAdjudication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].category", &resource.Item[numItem].Adjudication[numAdjudication].Category, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemAdjudicationReason(numItem int, numAdjudication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].reason", resource.Item[numItem].Adjudication[numAdjudication].Reason, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemAdjudicationValue(numItem int, numAdjudication int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return Float64Input("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].value", nil, htmlAttrs)
	}
	return Float64Input("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].value", resource.Item[numItem].Adjudication[numAdjudication].Value, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemDetailDetailSequence(numItem int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].detailSequence", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].detailSequence", &resource.Item[numItem].Detail[numDetail].DetailSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemDetailNoteNumber(numItem int, numDetail int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numNoteNumber >= len(resource.Item[numItem].Detail[numDetail].NoteNumber) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].Detail[numDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_ItemDetailSubDetailSubDetailSequence(numItem int, numDetail int, numSubDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].subDetailSequence", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].subDetailSequence", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].SubDetailSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ItemDetailSubDetailNoteNumber(numItem int, numDetail int, numSubDetail int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numNoteNumber >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemItemSequence(numAddItem int, numItemSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numItemSequence >= len(resource.AddItem[numAddItem].ItemSequence) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].itemSequence["+strconv.Itoa(numItemSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].itemSequence["+strconv.Itoa(numItemSequence)+"]", &resource.AddItem[numAddItem].ItemSequence[numItemSequence], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSequence(numAddItem int, numDetailSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetailSequence >= len(resource.AddItem[numAddItem].DetailSequence) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detailSequence["+strconv.Itoa(numDetailSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detailSequence["+strconv.Itoa(numDetailSequence)+"]", &resource.AddItem[numAddItem].DetailSequence[numDetailSequence], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemSubdetailSequence(numAddItem int, numSubdetailSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numSubdetailSequence >= len(resource.AddItem[numAddItem].SubdetailSequence) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].subdetailSequence["+strconv.Itoa(numSubdetailSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].subdetailSequence["+strconv.Itoa(numSubdetailSequence)+"]", &resource.AddItem[numAddItem].SubdetailSequence[numSubdetailSequence], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemProductOrService(numAddItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].productOrService", &resource.AddItem[numAddItem].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemModifier(numAddItem int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numModifier >= len(resource.AddItem[numAddItem].Modifier) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemProgramCode(numAddItem int, numProgramCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numProgramCode >= len(resource.AddItem[numAddItem].ProgramCode) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].programCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].programCode["+strconv.Itoa(numProgramCode)+"]", &resource.AddItem[numAddItem].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemServicedDate(numAddItem int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return DateInput("addItem["+strconv.Itoa(numAddItem)+"].servicedDate", nil, htmlAttrs)
	}
	return DateInput("addItem["+strconv.Itoa(numAddItem)+"].servicedDate", resource.AddItem[numAddItem].ServicedDate, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemLocationCodeableConcept(numAddItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].locationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].locationCodeableConcept", resource.AddItem[numAddItem].LocationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemFactor(numAddItem int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].factor", resource.AddItem[numAddItem].Factor, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemBodySite(numAddItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].bodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].bodySite", resource.AddItem[numAddItem].BodySite, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemSubSite(numAddItem int, numSubSite int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numSubSite >= len(resource.AddItem[numAddItem].SubSite) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].subSite["+strconv.Itoa(numSubSite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].subSite["+strconv.Itoa(numSubSite)+"]", &resource.AddItem[numAddItem].SubSite[numSubSite], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemNoteNumber(numAddItem int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numNoteNumber >= len(resource.AddItem[numAddItem].NoteNumber) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailProductOrService(numAddItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrService", &resource.AddItem[numAddItem].Detail[numDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailModifier(numAddItem int, numDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numModifier >= len(resource.AddItem[numAddItem].Detail[numDetail].Modifier) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Detail[numDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailFactor(numAddItem int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].factor", resource.AddItem[numAddItem].Detail[numDetail].Factor, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailNoteNumber(numAddItem int, numDetail int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numNoteNumber >= len(resource.AddItem[numAddItem].Detail[numDetail].NoteNumber) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailProductOrService(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrService", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailModifier(numAddItem int, numDetail int, numSubDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) || numModifier >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailFactor(numAddItem int, numDetail int, numSubDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].factor", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Factor, htmlAttrs)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailNoteNumber(numAddItem int, numDetail int, numSubDetail int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) || numNoteNumber >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ClaimResponse) T_TotalCategory(numTotal int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTotal >= len(resource.Total) {
		return CodeableConceptSelect("total["+strconv.Itoa(numTotal)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("total["+strconv.Itoa(numTotal)+"].category", &resource.Total[numTotal].Category, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_PaymentType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("payment.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payment.type", &resource.Payment.Type, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_PaymentAdjustmentReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("payment.adjustmentReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payment.adjustmentReason", resource.Payment.AdjustmentReason, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_PaymentDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("payment.date", nil, htmlAttrs)
	}
	return DateInput("payment.date", resource.Payment.Date, htmlAttrs)
}
func (resource *ClaimResponse) T_ProcessNoteNumber(numProcessNote int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return IntInput("processNote["+strconv.Itoa(numProcessNote)+"].number", nil, htmlAttrs)
	}
	return IntInput("processNote["+strconv.Itoa(numProcessNote)+"].number", resource.ProcessNote[numProcessNote].Number, htmlAttrs)
}
func (resource *ClaimResponse) T_ProcessNoteType(numProcessNote int, htmlAttrs string) templ.Component {
	optionsValueSet := VSNote_type

	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return CodeSelect("processNote["+strconv.Itoa(numProcessNote)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("processNote["+strconv.Itoa(numProcessNote)+"].type", resource.ProcessNote[numProcessNote].Type, optionsValueSet, htmlAttrs)
}
func (resource *ClaimResponse) T_ProcessNoteText(numProcessNote int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return StringInput("processNote["+strconv.Itoa(numProcessNote)+"].text", nil, htmlAttrs)
	}
	return StringInput("processNote["+strconv.Itoa(numProcessNote)+"].text", &resource.ProcessNote[numProcessNote].Text, htmlAttrs)
}
func (resource *ClaimResponse) T_InsuranceSequence(numInsurance int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return IntInput("insurance["+strconv.Itoa(numInsurance)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("insurance["+strconv.Itoa(numInsurance)+"].sequence", &resource.Insurance[numInsurance].Sequence, htmlAttrs)
}
func (resource *ClaimResponse) T_InsuranceFocal(numInsurance int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].focal", nil, htmlAttrs)
	}
	return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].focal", &resource.Insurance[numInsurance].Focal, htmlAttrs)
}
func (resource *ClaimResponse) T_InsuranceBusinessArrangement(numInsurance int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return StringInput("insurance["+strconv.Itoa(numInsurance)+"].businessArrangement", nil, htmlAttrs)
	}
	return StringInput("insurance["+strconv.Itoa(numInsurance)+"].businessArrangement", resource.Insurance[numInsurance].BusinessArrangement, htmlAttrs)
}
func (resource *ClaimResponse) T_ErrorItemSequence(numError int, htmlAttrs string) templ.Component {
	if resource == nil || numError >= len(resource.Error) {
		return IntInput("error["+strconv.Itoa(numError)+"].itemSequence", nil, htmlAttrs)
	}
	return IntInput("error["+strconv.Itoa(numError)+"].itemSequence", resource.Error[numError].ItemSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ErrorDetailSequence(numError int, htmlAttrs string) templ.Component {
	if resource == nil || numError >= len(resource.Error) {
		return IntInput("error["+strconv.Itoa(numError)+"].detailSequence", nil, htmlAttrs)
	}
	return IntInput("error["+strconv.Itoa(numError)+"].detailSequence", resource.Error[numError].DetailSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ErrorSubDetailSequence(numError int, htmlAttrs string) templ.Component {
	if resource == nil || numError >= len(resource.Error) {
		return IntInput("error["+strconv.Itoa(numError)+"].subDetailSequence", nil, htmlAttrs)
	}
	return IntInput("error["+strconv.Itoa(numError)+"].subDetailSequence", resource.Error[numError].SubDetailSequence, htmlAttrs)
}
func (resource *ClaimResponse) T_ErrorCode(numError int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numError >= len(resource.Error) {
		return CodeableConceptSelect("error["+strconv.Itoa(numError)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("error["+strconv.Itoa(numError)+"].code", &resource.Error[numError].Code, optionsValueSet, htmlAttrs)
}
