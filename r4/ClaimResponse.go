package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	ServicedDate            *string                      `json:"servicedDate"`
	ServicedPeriod          *Period                      `json:"servicedPeriod"`
	LocationCodeableConcept *CodeableConcept             `json:"locationCodeableConcept"`
	LocationAddress         *Address                     `json:"locationAddress"`
	LocationReference       *Reference                   `json:"locationReference"`
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

func (resource *ClaimResponse) ClaimResponseLanguage(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseStatus() templ.Component {
	optionsValueSet := VSFm_status

	if resource != nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Type, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseSubType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("subType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subType", resource.SubType, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseUse() templ.Component {
	optionsValueSet := VSClaim_use

	if resource != nil {
		return CodeSelect("use", nil, optionsValueSet)
	}
	return CodeSelect("use", &resource.Use, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseOutcome() templ.Component {
	optionsValueSet := VSRemittance_outcome

	if resource != nil {
		return CodeSelect("outcome", nil, optionsValueSet)
	}
	return CodeSelect("outcome", &resource.Outcome, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponsePayeeType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("payeeType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("payeeType", resource.PayeeType, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseFundsReserve(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("fundsReserve", nil, optionsValueSet)
	}
	return CodeableConceptSelect("fundsReserve", resource.FundsReserve, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseFormCode(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("formCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("formCode", resource.FormCode, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseItemAdjudicationCategory(numItem int, numAdjudication int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Item[numItem].Adjudication) >= numAdjudication {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Item[numItem].Adjudication[numAdjudication].Category, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseItemAdjudicationReason(numItem int, numAdjudication int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Item[numItem].Adjudication) >= numAdjudication {
		return CodeableConceptSelect("reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reason", resource.Item[numItem].Adjudication[numAdjudication].Reason, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemProductOrService(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", &resource.AddItem[numAddItem].ProductOrService, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemModifier(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.AddItem[numAddItem].Modifier[0], optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemProgramCode(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("programCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("programCode", &resource.AddItem[numAddItem].ProgramCode[0], optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemBodySite(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("bodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("bodySite", resource.AddItem[numAddItem].BodySite, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemSubSite(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("subSite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subSite", &resource.AddItem[numAddItem].SubSite[0], optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemDetailProductOrService(numAddItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", &resource.AddItem[numAddItem].Detail[numDetail].ProductOrService, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemDetailModifier(numAddItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.AddItem[numAddItem].Detail[numDetail].Modifier[0], optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemDetailSubDetailProductOrService(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseAddItemDetailSubDetailModifier(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[0], optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseTotalCategory(numTotal int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Total) >= numTotal {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Total[numTotal].Category, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponsePaymentType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Payment.Type, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponsePaymentAdjustmentReason(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("adjustmentReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("adjustmentReason", resource.Payment.AdjustmentReason, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseProcessNoteType(numProcessNote int) templ.Component {
	optionsValueSet := VSNote_type

	if resource != nil && len(resource.ProcessNote) >= numProcessNote {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", resource.ProcessNote[numProcessNote].Type, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseProcessNoteLanguage(numProcessNote int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.ProcessNote) >= numProcessNote {
		return CodeableConceptSelect("language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("language", resource.ProcessNote[numProcessNote].Language, optionsValueSet)
}
func (resource *ClaimResponse) ClaimResponseErrorCode(numError int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Error) >= numError {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Error[numError].Code, optionsValueSet)
}
