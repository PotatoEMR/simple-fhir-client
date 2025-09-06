package r4

//generated with command go run ./bultaoreune -nodownload
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

func (resource *ClaimResponse) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ClaimResponse.Id", nil)
	}
	return StringInput("ClaimResponse.Id", resource.Id)
}
func (resource *ClaimResponse) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ClaimResponse.ImplicitRules", nil)
	}
	return StringInput("ClaimResponse.ImplicitRules", resource.ImplicitRules)
}
func (resource *ClaimResponse) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ClaimResponse.Language", nil, optionsValueSet)
	}
	return CodeSelect("ClaimResponse.Language", resource.Language, optionsValueSet)
}
func (resource *ClaimResponse) T_Status() templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("ClaimResponse.Status", nil, optionsValueSet)
	}
	return CodeSelect("ClaimResponse.Status", &resource.Status, optionsValueSet)
}
func (resource *ClaimResponse) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ClaimResponse.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.Type", &resource.Type, optionsValueSet)
}
func (resource *ClaimResponse) T_SubType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ClaimResponse.SubType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.SubType", resource.SubType, optionsValueSet)
}
func (resource *ClaimResponse) T_Use() templ.Component {
	optionsValueSet := VSClaim_use

	if resource == nil {
		return CodeSelect("ClaimResponse.Use", nil, optionsValueSet)
	}
	return CodeSelect("ClaimResponse.Use", &resource.Use, optionsValueSet)
}
func (resource *ClaimResponse) T_Created() templ.Component {

	if resource == nil {
		return StringInput("ClaimResponse.Created", nil)
	}
	return StringInput("ClaimResponse.Created", &resource.Created)
}
func (resource *ClaimResponse) T_Outcome() templ.Component {
	optionsValueSet := VSRemittance_outcome

	if resource == nil {
		return CodeSelect("ClaimResponse.Outcome", nil, optionsValueSet)
	}
	return CodeSelect("ClaimResponse.Outcome", &resource.Outcome, optionsValueSet)
}
func (resource *ClaimResponse) T_Disposition() templ.Component {

	if resource == nil {
		return StringInput("ClaimResponse.Disposition", nil)
	}
	return StringInput("ClaimResponse.Disposition", resource.Disposition)
}
func (resource *ClaimResponse) T_PreAuthRef() templ.Component {

	if resource == nil {
		return StringInput("ClaimResponse.PreAuthRef", nil)
	}
	return StringInput("ClaimResponse.PreAuthRef", resource.PreAuthRef)
}
func (resource *ClaimResponse) T_PayeeType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ClaimResponse.PayeeType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.PayeeType", resource.PayeeType, optionsValueSet)
}
func (resource *ClaimResponse) T_FundsReserve(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ClaimResponse.FundsReserve", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.FundsReserve", resource.FundsReserve, optionsValueSet)
}
func (resource *ClaimResponse) T_FormCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ClaimResponse.FormCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.FormCode", resource.FormCode, optionsValueSet)
}
func (resource *ClaimResponse) T_ItemId(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Id", nil)
	}
	return StringInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Id", resource.Item[numItem].Id)
}
func (resource *ClaimResponse) T_ItemItemSequence(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return IntInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].ItemSequence", nil)
	}
	return IntInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].ItemSequence", &resource.Item[numItem].ItemSequence)
}
func (resource *ClaimResponse) T_ItemNoteNumber(numItem int, numNoteNumber int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].NoteNumber) >= numNoteNumber {
		return IntInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil)
	}
	return IntInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].NoteNumber[numNoteNumber])
}
func (resource *ClaimResponse) T_ItemAdjudicationId(numItem int, numAdjudication int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Adjudication) >= numAdjudication {
		return StringInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Adjudication["+strconv.Itoa(numAdjudication)+"].Id", nil)
	}
	return StringInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Adjudication["+strconv.Itoa(numAdjudication)+"].Id", resource.Item[numItem].Adjudication[numAdjudication].Id)
}
func (resource *ClaimResponse) T_ItemAdjudicationCategory(numItem int, numAdjudication int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Adjudication) >= numAdjudication {
		return CodeableConceptSelect("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Adjudication["+strconv.Itoa(numAdjudication)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Adjudication["+strconv.Itoa(numAdjudication)+"].Category", &resource.Item[numItem].Adjudication[numAdjudication].Category, optionsValueSet)
}
func (resource *ClaimResponse) T_ItemAdjudicationReason(numItem int, numAdjudication int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Adjudication) >= numAdjudication {
		return CodeableConceptSelect("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Adjudication["+strconv.Itoa(numAdjudication)+"].Reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Adjudication["+strconv.Itoa(numAdjudication)+"].Reason", resource.Item[numItem].Adjudication[numAdjudication].Reason, optionsValueSet)
}
func (resource *ClaimResponse) T_ItemAdjudicationValue(numItem int, numAdjudication int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Adjudication) >= numAdjudication {
		return Float64Input("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Adjudication["+strconv.Itoa(numAdjudication)+"].Value", nil)
	}
	return Float64Input("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Adjudication["+strconv.Itoa(numAdjudication)+"].Value", resource.Item[numItem].Adjudication[numAdjudication].Value)
}
func (resource *ClaimResponse) T_ItemDetailId(numItem int, numDetail int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail {
		return StringInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Id", nil)
	}
	return StringInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Id", resource.Item[numItem].Detail[numDetail].Id)
}
func (resource *ClaimResponse) T_ItemDetailDetailSequence(numItem int, numDetail int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail {
		return IntInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].DetailSequence", nil)
	}
	return IntInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].DetailSequence", &resource.Item[numItem].Detail[numDetail].DetailSequence)
}
func (resource *ClaimResponse) T_ItemDetailNoteNumber(numItem int, numDetail int, numNoteNumber int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].NoteNumber) >= numNoteNumber {
		return IntInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil)
	}
	return IntInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].Detail[numDetail].NoteNumber[numNoteNumber])
}
func (resource *ClaimResponse) T_ItemDetailSubDetailId(numItem int, numDetail int, numSubDetail int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return StringInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Id", nil)
	}
	return StringInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Id", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Id)
}
func (resource *ClaimResponse) T_ItemDetailSubDetailSubDetailSequence(numItem int, numDetail int, numSubDetail int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return IntInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].SubDetailSequence", nil)
	}
	return IntInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].SubDetailSequence", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].SubDetailSequence)
}
func (resource *ClaimResponse) T_ItemDetailSubDetailNoteNumber(numItem int, numDetail int, numSubDetail int, numNoteNumber int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber) >= numNoteNumber {
		return IntInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil)
	}
	return IntInput("ClaimResponse.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber[numNoteNumber])
}
func (resource *ClaimResponse) T_AddItemId(numAddItem int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem {
		return StringInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Id", nil)
	}
	return StringInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Id", resource.AddItem[numAddItem].Id)
}
func (resource *ClaimResponse) T_AddItemItemSequence(numAddItem int, numItemSequence int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].ItemSequence) >= numItemSequence {
		return IntInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].ItemSequence["+strconv.Itoa(numItemSequence)+"]", nil)
	}
	return IntInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].ItemSequence["+strconv.Itoa(numItemSequence)+"]", &resource.AddItem[numAddItem].ItemSequence[numItemSequence])
}
func (resource *ClaimResponse) T_AddItemDetailSequence(numAddItem int, numDetailSequence int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].DetailSequence) >= numDetailSequence {
		return IntInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].DetailSequence["+strconv.Itoa(numDetailSequence)+"]", nil)
	}
	return IntInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].DetailSequence["+strconv.Itoa(numDetailSequence)+"]", &resource.AddItem[numAddItem].DetailSequence[numDetailSequence])
}
func (resource *ClaimResponse) T_AddItemSubdetailSequence(numAddItem int, numSubdetailSequence int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].SubdetailSequence) >= numSubdetailSequence {
		return IntInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].SubdetailSequence["+strconv.Itoa(numSubdetailSequence)+"]", nil)
	}
	return IntInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].SubdetailSequence["+strconv.Itoa(numSubdetailSequence)+"]", &resource.AddItem[numAddItem].SubdetailSequence[numSubdetailSequence])
}
func (resource *ClaimResponse) T_AddItemProductOrService(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].ProductOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].ProductOrService", &resource.AddItem[numAddItem].ProductOrService, optionsValueSet)
}
func (resource *ClaimResponse) T_AddItemModifier(numAddItem int, numModifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Modifier) >= numModifier {
		return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Modifier[numModifier], optionsValueSet)
}
func (resource *ClaimResponse) T_AddItemProgramCode(numAddItem int, numProgramCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].ProgramCode) >= numProgramCode {
		return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", &resource.AddItem[numAddItem].ProgramCode[numProgramCode], optionsValueSet)
}
func (resource *ClaimResponse) T_AddItemFactor(numAddItem int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem {
		return Float64Input("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Factor", nil)
	}
	return Float64Input("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Factor", resource.AddItem[numAddItem].Factor)
}
func (resource *ClaimResponse) T_AddItemBodySite(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].BodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].BodySite", resource.AddItem[numAddItem].BodySite, optionsValueSet)
}
func (resource *ClaimResponse) T_AddItemSubSite(numAddItem int, numSubSite int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].SubSite) >= numSubSite {
		return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].SubSite["+strconv.Itoa(numSubSite)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].SubSite["+strconv.Itoa(numSubSite)+"]", &resource.AddItem[numAddItem].SubSite[numSubSite], optionsValueSet)
}
func (resource *ClaimResponse) T_AddItemNoteNumber(numAddItem int, numNoteNumber int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].NoteNumber) >= numNoteNumber {
		return IntInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil)
	}
	return IntInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].NoteNumber[numNoteNumber])
}
func (resource *ClaimResponse) T_AddItemDetailId(numAddItem int, numDetail int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return StringInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].Id", nil)
	}
	return StringInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].Id", resource.AddItem[numAddItem].Detail[numDetail].Id)
}
func (resource *ClaimResponse) T_AddItemDetailProductOrService(numAddItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProductOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProductOrService", &resource.AddItem[numAddItem].Detail[numDetail].ProductOrService, optionsValueSet)
}
func (resource *ClaimResponse) T_AddItemDetailModifier(numAddItem int, numDetail int, numModifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail || len(resource.AddItem[numAddItem].Detail[numDetail].Modifier) >= numModifier {
		return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Detail[numDetail].Modifier[numModifier], optionsValueSet)
}
func (resource *ClaimResponse) T_AddItemDetailFactor(numAddItem int, numDetail int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return Float64Input("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].Factor", nil)
	}
	return Float64Input("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].Factor", resource.AddItem[numAddItem].Detail[numDetail].Factor)
}
func (resource *ClaimResponse) T_AddItemDetailNoteNumber(numAddItem int, numDetail int, numNoteNumber int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail || len(resource.AddItem[numAddItem].Detail[numDetail].NoteNumber) >= numNoteNumber {
		return IntInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil)
	}
	return IntInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].NoteNumber[numNoteNumber])
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailId(numAddItem int, numDetail int, numSubDetail int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail || len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return StringInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Id", nil)
	}
	return StringInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Id", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Id)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailProductOrService(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail || len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailModifier(numAddItem int, numDetail int, numSubDetail int, numModifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail || len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail || len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier) >= numModifier {
		return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[numModifier], optionsValueSet)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailFactor(numAddItem int, numDetail int, numSubDetail int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail || len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return Float64Input("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Factor", nil)
	}
	return Float64Input("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Factor", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Factor)
}
func (resource *ClaimResponse) T_AddItemDetailSubDetailNoteNumber(numAddItem int, numDetail int, numSubDetail int, numNoteNumber int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail || len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail || len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber) >= numNoteNumber {
		return IntInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil)
	}
	return IntInput("ClaimResponse.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber[numNoteNumber])
}
func (resource *ClaimResponse) T_TotalId(numTotal int) templ.Component {

	if resource == nil || len(resource.Total) >= numTotal {
		return StringInput("ClaimResponse.Total["+strconv.Itoa(numTotal)+"].Id", nil)
	}
	return StringInput("ClaimResponse.Total["+strconv.Itoa(numTotal)+"].Id", resource.Total[numTotal].Id)
}
func (resource *ClaimResponse) T_TotalCategory(numTotal int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Total) >= numTotal {
		return CodeableConceptSelect("ClaimResponse.Total["+strconv.Itoa(numTotal)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.Total["+strconv.Itoa(numTotal)+"].Category", &resource.Total[numTotal].Category, optionsValueSet)
}
func (resource *ClaimResponse) T_PaymentId() templ.Component {

	if resource == nil {
		return StringInput("ClaimResponse.Payment.Id", nil)
	}
	return StringInput("ClaimResponse.Payment.Id", resource.Payment.Id)
}
func (resource *ClaimResponse) T_PaymentType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ClaimResponse.Payment.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.Payment.Type", &resource.Payment.Type, optionsValueSet)
}
func (resource *ClaimResponse) T_PaymentAdjustmentReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ClaimResponse.Payment.AdjustmentReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.Payment.AdjustmentReason", resource.Payment.AdjustmentReason, optionsValueSet)
}
func (resource *ClaimResponse) T_PaymentDate() templ.Component {

	if resource == nil {
		return StringInput("ClaimResponse.Payment.Date", nil)
	}
	return StringInput("ClaimResponse.Payment.Date", resource.Payment.Date)
}
func (resource *ClaimResponse) T_ProcessNoteId(numProcessNote int) templ.Component {

	if resource == nil || len(resource.ProcessNote) >= numProcessNote {
		return StringInput("ClaimResponse.ProcessNote["+strconv.Itoa(numProcessNote)+"].Id", nil)
	}
	return StringInput("ClaimResponse.ProcessNote["+strconv.Itoa(numProcessNote)+"].Id", resource.ProcessNote[numProcessNote].Id)
}
func (resource *ClaimResponse) T_ProcessNoteNumber(numProcessNote int) templ.Component {

	if resource == nil || len(resource.ProcessNote) >= numProcessNote {
		return IntInput("ClaimResponse.ProcessNote["+strconv.Itoa(numProcessNote)+"].Number", nil)
	}
	return IntInput("ClaimResponse.ProcessNote["+strconv.Itoa(numProcessNote)+"].Number", resource.ProcessNote[numProcessNote].Number)
}
func (resource *ClaimResponse) T_ProcessNoteType(numProcessNote int) templ.Component {
	optionsValueSet := VSNote_type

	if resource == nil || len(resource.ProcessNote) >= numProcessNote {
		return CodeSelect("ClaimResponse.ProcessNote["+strconv.Itoa(numProcessNote)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("ClaimResponse.ProcessNote["+strconv.Itoa(numProcessNote)+"].Type", resource.ProcessNote[numProcessNote].Type, optionsValueSet)
}
func (resource *ClaimResponse) T_ProcessNoteText(numProcessNote int) templ.Component {

	if resource == nil || len(resource.ProcessNote) >= numProcessNote {
		return StringInput("ClaimResponse.ProcessNote["+strconv.Itoa(numProcessNote)+"].Text", nil)
	}
	return StringInput("ClaimResponse.ProcessNote["+strconv.Itoa(numProcessNote)+"].Text", &resource.ProcessNote[numProcessNote].Text)
}
func (resource *ClaimResponse) T_ProcessNoteLanguage(numProcessNote int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ProcessNote) >= numProcessNote {
		return CodeableConceptSelect("ClaimResponse.ProcessNote["+strconv.Itoa(numProcessNote)+"].Language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.ProcessNote["+strconv.Itoa(numProcessNote)+"].Language", resource.ProcessNote[numProcessNote].Language, optionsValueSet)
}
func (resource *ClaimResponse) T_InsuranceId(numInsurance int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance {
		return StringInput("ClaimResponse.Insurance["+strconv.Itoa(numInsurance)+"].Id", nil)
	}
	return StringInput("ClaimResponse.Insurance["+strconv.Itoa(numInsurance)+"].Id", resource.Insurance[numInsurance].Id)
}
func (resource *ClaimResponse) T_InsuranceSequence(numInsurance int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance {
		return IntInput("ClaimResponse.Insurance["+strconv.Itoa(numInsurance)+"].Sequence", nil)
	}
	return IntInput("ClaimResponse.Insurance["+strconv.Itoa(numInsurance)+"].Sequence", &resource.Insurance[numInsurance].Sequence)
}
func (resource *ClaimResponse) T_InsuranceFocal(numInsurance int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance {
		return BoolInput("ClaimResponse.Insurance["+strconv.Itoa(numInsurance)+"].Focal", nil)
	}
	return BoolInput("ClaimResponse.Insurance["+strconv.Itoa(numInsurance)+"].Focal", &resource.Insurance[numInsurance].Focal)
}
func (resource *ClaimResponse) T_InsuranceBusinessArrangement(numInsurance int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance {
		return StringInput("ClaimResponse.Insurance["+strconv.Itoa(numInsurance)+"].BusinessArrangement", nil)
	}
	return StringInput("ClaimResponse.Insurance["+strconv.Itoa(numInsurance)+"].BusinessArrangement", resource.Insurance[numInsurance].BusinessArrangement)
}
func (resource *ClaimResponse) T_ErrorId(numError int) templ.Component {

	if resource == nil || len(resource.Error) >= numError {
		return StringInput("ClaimResponse.Error["+strconv.Itoa(numError)+"].Id", nil)
	}
	return StringInput("ClaimResponse.Error["+strconv.Itoa(numError)+"].Id", resource.Error[numError].Id)
}
func (resource *ClaimResponse) T_ErrorItemSequence(numError int) templ.Component {

	if resource == nil || len(resource.Error) >= numError {
		return IntInput("ClaimResponse.Error["+strconv.Itoa(numError)+"].ItemSequence", nil)
	}
	return IntInput("ClaimResponse.Error["+strconv.Itoa(numError)+"].ItemSequence", resource.Error[numError].ItemSequence)
}
func (resource *ClaimResponse) T_ErrorDetailSequence(numError int) templ.Component {

	if resource == nil || len(resource.Error) >= numError {
		return IntInput("ClaimResponse.Error["+strconv.Itoa(numError)+"].DetailSequence", nil)
	}
	return IntInput("ClaimResponse.Error["+strconv.Itoa(numError)+"].DetailSequence", resource.Error[numError].DetailSequence)
}
func (resource *ClaimResponse) T_ErrorSubDetailSequence(numError int) templ.Component {

	if resource == nil || len(resource.Error) >= numError {
		return IntInput("ClaimResponse.Error["+strconv.Itoa(numError)+"].SubDetailSequence", nil)
	}
	return IntInput("ClaimResponse.Error["+strconv.Itoa(numError)+"].SubDetailSequence", resource.Error[numError].SubDetailSequence)
}
func (resource *ClaimResponse) T_ErrorCode(numError int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Error) >= numError {
		return CodeableConceptSelect("ClaimResponse.Error["+strconv.Itoa(numError)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ClaimResponse.Error["+strconv.Itoa(numError)+"].Code", &resource.Error[numError].Code, optionsValueSet)
}
