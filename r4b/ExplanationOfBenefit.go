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

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefit struct {
	Id                    *string                              `json:"id,omitempty"`
	Meta                  *Meta                                `json:"meta,omitempty"`
	ImplicitRules         *string                              `json:"implicitRules,omitempty"`
	Language              *string                              `json:"language,omitempty"`
	Text                  *Narrative                           `json:"text,omitempty"`
	Contained             []Resource                           `json:"contained,omitempty"`
	Extension             []Extension                          `json:"extension,omitempty"`
	ModifierExtension     []Extension                          `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                         `json:"identifier,omitempty"`
	Status                string                               `json:"status"`
	Type                  CodeableConcept                      `json:"type"`
	SubType               *CodeableConcept                     `json:"subType,omitempty"`
	Use                   string                               `json:"use"`
	Patient               Reference                            `json:"patient"`
	BillablePeriod        *Period                              `json:"billablePeriod,omitempty"`
	Created               time.Time                            `json:"created,format:'2006-01-02T15:04:05Z07:00'"`
	Enterer               *Reference                           `json:"enterer,omitempty"`
	Insurer               Reference                            `json:"insurer"`
	Provider              Reference                            `json:"provider"`
	Priority              *CodeableConcept                     `json:"priority,omitempty"`
	FundsReserveRequested *CodeableConcept                     `json:"fundsReserveRequested,omitempty"`
	FundsReserve          *CodeableConcept                     `json:"fundsReserve,omitempty"`
	Related               []ExplanationOfBenefitRelated        `json:"related,omitempty"`
	Prescription          *Reference                           `json:"prescription,omitempty"`
	OriginalPrescription  *Reference                           `json:"originalPrescription,omitempty"`
	Payee                 *ExplanationOfBenefitPayee           `json:"payee,omitempty"`
	Referral              *Reference                           `json:"referral,omitempty"`
	Facility              *Reference                           `json:"facility,omitempty"`
	Claim                 *Reference                           `json:"claim,omitempty"`
	ClaimResponse         *Reference                           `json:"claimResponse,omitempty"`
	Outcome               string                               `json:"outcome"`
	Disposition           *string                              `json:"disposition,omitempty"`
	PreAuthRef            []string                             `json:"preAuthRef,omitempty"`
	PreAuthRefPeriod      []Period                             `json:"preAuthRefPeriod,omitempty"`
	CareTeam              []ExplanationOfBenefitCareTeam       `json:"careTeam,omitempty"`
	SupportingInfo        []ExplanationOfBenefitSupportingInfo `json:"supportingInfo,omitempty"`
	Diagnosis             []ExplanationOfBenefitDiagnosis      `json:"diagnosis,omitempty"`
	Procedure             []ExplanationOfBenefitProcedure      `json:"procedure,omitempty"`
	Precedence            *int                                 `json:"precedence,omitempty"`
	Insurance             []ExplanationOfBenefitInsurance      `json:"insurance"`
	Accident              *ExplanationOfBenefitAccident        `json:"accident,omitempty"`
	Item                  []ExplanationOfBenefitItem           `json:"item,omitempty"`
	AddItem               []ExplanationOfBenefitAddItem        `json:"addItem,omitempty"`
	Total                 []ExplanationOfBenefitTotal          `json:"total,omitempty"`
	Payment               *ExplanationOfBenefitPayment         `json:"payment,omitempty"`
	FormCode              *CodeableConcept                     `json:"formCode,omitempty"`
	Form                  *Attachment                          `json:"form,omitempty"`
	ProcessNote           []ExplanationOfBenefitProcessNote    `json:"processNote,omitempty"`
	BenefitPeriod         *Period                              `json:"benefitPeriod,omitempty"`
	BenefitBalance        []ExplanationOfBenefitBenefitBalance `json:"benefitBalance,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitRelated struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Claim             *Reference       `json:"claim,omitempty"`
	Relationship      *CodeableConcept `json:"relationship,omitempty"`
	Reference         *Identifier      `json:"reference,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitPayee struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Party             *Reference       `json:"party,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitCareTeam struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Sequence          int              `json:"sequence"`
	Provider          Reference        `json:"provider"`
	Responsible       *bool            `json:"responsible,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Qualification     *CodeableConcept `json:"qualification,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitSupportingInfo struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Sequence          int              `json:"sequence"`
	Category          CodeableConcept  `json:"category"`
	Code              *CodeableConcept `json:"code,omitempty"`
	TimingDate        *time.Time       `json:"timingDate,omitempty,format:'2006-01-02'"`
	TimingPeriod      *Period          `json:"timingPeriod,omitempty"`
	ValueBoolean      *bool            `json:"valueBoolean,omitempty"`
	ValueString       *string          `json:"valueString,omitempty"`
	ValueQuantity     *Quantity        `json:"valueQuantity,omitempty"`
	ValueAttachment   *Attachment      `json:"valueAttachment,omitempty"`
	ValueReference    *Reference       `json:"valueReference,omitempty"`
	Reason            *Coding          `json:"reason,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitDiagnosis struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Sequence                 int               `json:"sequence"`
	DiagnosisCodeableConcept CodeableConcept   `json:"diagnosisCodeableConcept"`
	DiagnosisReference       Reference         `json:"diagnosisReference"`
	Type                     []CodeableConcept `json:"type,omitempty"`
	OnAdmission              *CodeableConcept  `json:"onAdmission,omitempty"`
	PackageCode              *CodeableConcept  `json:"packageCode,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitProcedure struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Sequence                 int               `json:"sequence"`
	Type                     []CodeableConcept `json:"type,omitempty"`
	Date                     *time.Time        `json:"date,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	ProcedureCodeableConcept CodeableConcept   `json:"procedureCodeableConcept"`
	ProcedureReference       Reference         `json:"procedureReference"`
	Udi                      []Reference       `json:"udi,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitInsurance struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Focal             bool        `json:"focal"`
	Coverage          Reference   `json:"coverage"`
	PreAuthRef        []string    `json:"preAuthRef,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitAccident struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Date              *time.Time       `json:"date,omitempty,format:'2006-01-02'"`
	Type              *CodeableConcept `json:"type,omitempty"`
	LocationAddress   *Address         `json:"locationAddress,omitempty"`
	LocationReference *Reference       `json:"locationReference,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitItem struct {
	Id                      *string                                `json:"id,omitempty"`
	Extension               []Extension                            `json:"extension,omitempty"`
	ModifierExtension       []Extension                            `json:"modifierExtension,omitempty"`
	Sequence                int                                    `json:"sequence"`
	CareTeamSequence        []int                                  `json:"careTeamSequence,omitempty"`
	DiagnosisSequence       []int                                  `json:"diagnosisSequence,omitempty"`
	ProcedureSequence       []int                                  `json:"procedureSequence,omitempty"`
	InformationSequence     []int                                  `json:"informationSequence,omitempty"`
	Revenue                 *CodeableConcept                       `json:"revenue,omitempty"`
	Category                *CodeableConcept                       `json:"category,omitempty"`
	ProductOrService        CodeableConcept                        `json:"productOrService"`
	Modifier                []CodeableConcept                      `json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept                      `json:"programCode,omitempty"`
	ServicedDate            *time.Time                             `json:"servicedDate,omitempty,format:'2006-01-02'"`
	ServicedPeriod          *Period                                `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept                       `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                               `json:"locationAddress,omitempty"`
	LocationReference       *Reference                             `json:"locationReference,omitempty"`
	Quantity                *Quantity                              `json:"quantity,omitempty"`
	UnitPrice               *Money                                 `json:"unitPrice,omitempty"`
	Factor                  *float64                               `json:"factor,omitempty"`
	Net                     *Money                                 `json:"net,omitempty"`
	Udi                     []Reference                            `json:"udi,omitempty"`
	BodySite                *CodeableConcept                       `json:"bodySite,omitempty"`
	SubSite                 []CodeableConcept                      `json:"subSite,omitempty"`
	Encounter               []Reference                            `json:"encounter,omitempty"`
	NoteNumber              []int                                  `json:"noteNumber,omitempty"`
	Adjudication            []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`
	Detail                  []ExplanationOfBenefitItemDetail       `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitItemAdjudication struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Category          CodeableConcept  `json:"category"`
	Reason            *CodeableConcept `json:"reason,omitempty"`
	Amount            *Money           `json:"amount,omitempty"`
	Value             *float64         `json:"value,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitItemDetail struct {
	Id                *string                                   `json:"id,omitempty"`
	Extension         []Extension                               `json:"extension,omitempty"`
	ModifierExtension []Extension                               `json:"modifierExtension,omitempty"`
	Sequence          int                                       `json:"sequence"`
	Revenue           *CodeableConcept                          `json:"revenue,omitempty"`
	Category          *CodeableConcept                          `json:"category,omitempty"`
	ProductOrService  CodeableConcept                           `json:"productOrService"`
	Modifier          []CodeableConcept                         `json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept                         `json:"programCode,omitempty"`
	Quantity          *Quantity                                 `json:"quantity,omitempty"`
	UnitPrice         *Money                                    `json:"unitPrice,omitempty"`
	Factor            *float64                                  `json:"factor,omitempty"`
	Net               *Money                                    `json:"net,omitempty"`
	Udi               []Reference                               `json:"udi,omitempty"`
	NoteNumber        []int                                     `json:"noteNumber,omitempty"`
	SubDetail         []ExplanationOfBenefitItemDetailSubDetail `json:"subDetail,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitItemDetailSubDetail struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Sequence          int               `json:"sequence"`
	Revenue           *CodeableConcept  `json:"revenue,omitempty"`
	Category          *CodeableConcept  `json:"category,omitempty"`
	ProductOrService  CodeableConcept   `json:"productOrService"`
	Modifier          []CodeableConcept `json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept `json:"programCode,omitempty"`
	Quantity          *Quantity         `json:"quantity,omitempty"`
	UnitPrice         *Money            `json:"unitPrice,omitempty"`
	Factor            *float64          `json:"factor,omitempty"`
	Net               *Money            `json:"net,omitempty"`
	Udi               []Reference       `json:"udi,omitempty"`
	NoteNumber        []int             `json:"noteNumber,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitAddItem struct {
	Id                      *string                             `json:"id,omitempty"`
	Extension               []Extension                         `json:"extension,omitempty"`
	ModifierExtension       []Extension                         `json:"modifierExtension,omitempty"`
	ItemSequence            []int                               `json:"itemSequence,omitempty"`
	DetailSequence          []int                               `json:"detailSequence,omitempty"`
	SubDetailSequence       []int                               `json:"subDetailSequence,omitempty"`
	Provider                []Reference                         `json:"provider,omitempty"`
	ProductOrService        CodeableConcept                     `json:"productOrService"`
	Modifier                []CodeableConcept                   `json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept                   `json:"programCode,omitempty"`
	ServicedDate            *time.Time                          `json:"servicedDate,omitempty,format:'2006-01-02'"`
	ServicedPeriod          *Period                             `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept                    `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                            `json:"locationAddress,omitempty"`
	LocationReference       *Reference                          `json:"locationReference,omitempty"`
	Quantity                *Quantity                           `json:"quantity,omitempty"`
	UnitPrice               *Money                              `json:"unitPrice,omitempty"`
	Factor                  *float64                            `json:"factor,omitempty"`
	Net                     *Money                              `json:"net,omitempty"`
	BodySite                *CodeableConcept                    `json:"bodySite,omitempty"`
	SubSite                 []CodeableConcept                   `json:"subSite,omitempty"`
	NoteNumber              []int                               `json:"noteNumber,omitempty"`
	Detail                  []ExplanationOfBenefitAddItemDetail `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitAddItemDetail struct {
	Id                *string                                      `json:"id,omitempty"`
	Extension         []Extension                                  `json:"extension,omitempty"`
	ModifierExtension []Extension                                  `json:"modifierExtension,omitempty"`
	ProductOrService  CodeableConcept                              `json:"productOrService"`
	Modifier          []CodeableConcept                            `json:"modifier,omitempty"`
	Quantity          *Quantity                                    `json:"quantity,omitempty"`
	UnitPrice         *Money                                       `json:"unitPrice,omitempty"`
	Factor            *float64                                     `json:"factor,omitempty"`
	Net               *Money                                       `json:"net,omitempty"`
	NoteNumber        []int                                        `json:"noteNumber,omitempty"`
	SubDetail         []ExplanationOfBenefitAddItemDetailSubDetail `json:"subDetail,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitAddItemDetailSubDetail struct {
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

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitTotal struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Category          CodeableConcept `json:"category"`
	Amount            Money           `json:"amount"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitPayment struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Adjustment        *Money           `json:"adjustment,omitempty"`
	AdjustmentReason  *CodeableConcept `json:"adjustmentReason,omitempty"`
	Date              *time.Time       `json:"date,omitempty,format:'2006-01-02'"`
	Amount            *Money           `json:"amount,omitempty"`
	Identifier        *Identifier      `json:"identifier,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitProcessNote struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Number            *int             `json:"number,omitempty"`
	Type              *string          `json:"type,omitempty"`
	Text              *string          `json:"text,omitempty"`
	Language          *CodeableConcept `json:"language,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitBenefitBalance struct {
	Id                *string                                       `json:"id,omitempty"`
	Extension         []Extension                                   `json:"extension,omitempty"`
	ModifierExtension []Extension                                   `json:"modifierExtension,omitempty"`
	Category          CodeableConcept                               `json:"category"`
	Excluded          *bool                                         `json:"excluded,omitempty"`
	Name              *string                                       `json:"name,omitempty"`
	Description       *string                                       `json:"description,omitempty"`
	Network           *CodeableConcept                              `json:"network,omitempty"`
	Unit              *CodeableConcept                              `json:"unit,omitempty"`
	Term              *CodeableConcept                              `json:"term,omitempty"`
	Financial         []ExplanationOfBenefitBenefitBalanceFinancial `json:"financial,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitBenefitBalanceFinancial struct {
	Id                 *string         `json:"id,omitempty"`
	Extension          []Extension     `json:"extension,omitempty"`
	ModifierExtension  []Extension     `json:"modifierExtension,omitempty"`
	Type               CodeableConcept `json:"type"`
	AllowedUnsignedInt *int            `json:"allowedUnsignedInt,omitempty"`
	AllowedString      *string         `json:"allowedString,omitempty"`
	AllowedMoney       *Money          `json:"allowedMoney,omitempty"`
	UsedUnsignedInt    *int            `json:"usedUnsignedInt,omitempty"`
	UsedMoney          *Money          `json:"usedMoney,omitempty"`
}

type OtherExplanationOfBenefit ExplanationOfBenefit

// on convert struct to json, automatically add resourceType=ExplanationOfBenefit
func (r ExplanationOfBenefit) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherExplanationOfBenefit
		ResourceType string `json:"resourceType"`
	}{
		OtherExplanationOfBenefit: OtherExplanationOfBenefit(r),
		ResourceType:              "ExplanationOfBenefit",
	})
}
func (r ExplanationOfBenefit) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "ExplanationOfBenefit/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "ExplanationOfBenefit"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *ExplanationOfBenefit) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSExplanationofbenefit_status

	if resource == nil {
		return CodeSelect("Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SubType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SubType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubType", resource.SubType, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Use(htmlAttrs string) templ.Component {
	optionsValueSet := VSClaim_use

	if resource == nil {
		return CodeSelect("Use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Use", &resource.Use, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Created(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Created", nil, htmlAttrs)
	}
	return DateTimeInput("Created", &resource.Created, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Priority(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_FundsReserveRequested(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("FundsReserveRequested", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("FundsReserveRequested", resource.FundsReserveRequested, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_FundsReserve(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("FundsReserve", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("FundsReserve", resource.FundsReserve, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Outcome(htmlAttrs string) templ.Component {
	optionsValueSet := VSRemittance_outcome

	if resource == nil {
		return CodeSelect("Outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Outcome", &resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Disposition(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("Disposition", nil, htmlAttrs)
	}
	return StringInput("Disposition", resource.Disposition, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PreAuthRef(numPreAuthRef int, htmlAttrs string) templ.Component {
	if resource == nil || numPreAuthRef >= len(resource.PreAuthRef) {
		return StringInput("PreAuthRef["+strconv.Itoa(numPreAuthRef)+"]", nil, htmlAttrs)
	}
	return StringInput("PreAuthRef["+strconv.Itoa(numPreAuthRef)+"]", &resource.PreAuthRef[numPreAuthRef], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Precedence(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("Precedence", nil, htmlAttrs)
	}
	return IntInput("Precedence", resource.Precedence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_FormCode(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("FormCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("FormCode", resource.FormCode, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_RelatedRelationship(numRelated int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelated >= len(resource.Related) {
		return CodeableConceptSelect("Related["+strconv.Itoa(numRelated)+"]Relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Related["+strconv.Itoa(numRelated)+"]Relationship", resource.Related[numRelated].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PayeeType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PayeeType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PayeeType", resource.Payee.Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_CareTeamSequence(numCareTeam int, htmlAttrs string) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return IntInput("CareTeam["+strconv.Itoa(numCareTeam)+"]Sequence", nil, htmlAttrs)
	}
	return IntInput("CareTeam["+strconv.Itoa(numCareTeam)+"]Sequence", &resource.CareTeam[numCareTeam].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_CareTeamResponsible(numCareTeam int, htmlAttrs string) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return BoolInput("CareTeam["+strconv.Itoa(numCareTeam)+"]Responsible", nil, htmlAttrs)
	}
	return BoolInput("CareTeam["+strconv.Itoa(numCareTeam)+"]Responsible", resource.CareTeam[numCareTeam].Responsible, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_CareTeamRole(numCareTeam int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return CodeableConceptSelect("CareTeam["+strconv.Itoa(numCareTeam)+"]Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CareTeam["+strconv.Itoa(numCareTeam)+"]Role", resource.CareTeam[numCareTeam].Role, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_CareTeamQualification(numCareTeam int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return CodeableConceptSelect("CareTeam["+strconv.Itoa(numCareTeam)+"]Qualification", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("CareTeam["+strconv.Itoa(numCareTeam)+"]Qualification", resource.CareTeam[numCareTeam].Qualification, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoSequence(numSupportingInfo int, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return IntInput("SupportingInfo["+strconv.Itoa(numSupportingInfo)+"]Sequence", nil, htmlAttrs)
	}
	return IntInput("SupportingInfo["+strconv.Itoa(numSupportingInfo)+"]Sequence", &resource.SupportingInfo[numSupportingInfo].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoCategory(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodeableConceptSelect("SupportingInfo["+strconv.Itoa(numSupportingInfo)+"]Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SupportingInfo["+strconv.Itoa(numSupportingInfo)+"]Category", &resource.SupportingInfo[numSupportingInfo].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoCode(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodeableConceptSelect("SupportingInfo["+strconv.Itoa(numSupportingInfo)+"]Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SupportingInfo["+strconv.Itoa(numSupportingInfo)+"]Code", resource.SupportingInfo[numSupportingInfo].Code, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoTimingDate(numSupportingInfo int, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return DateInput("SupportingInfo["+strconv.Itoa(numSupportingInfo)+"]TimingDate", nil, htmlAttrs)
	}
	return DateInput("SupportingInfo["+strconv.Itoa(numSupportingInfo)+"]TimingDate", resource.SupportingInfo[numSupportingInfo].TimingDate, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoValueBoolean(numSupportingInfo int, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return BoolInput("SupportingInfo["+strconv.Itoa(numSupportingInfo)+"]ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("SupportingInfo["+strconv.Itoa(numSupportingInfo)+"]ValueBoolean", resource.SupportingInfo[numSupportingInfo].ValueBoolean, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoValueString(numSupportingInfo int, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return StringInput("SupportingInfo["+strconv.Itoa(numSupportingInfo)+"]ValueString", nil, htmlAttrs)
	}
	return StringInput("SupportingInfo["+strconv.Itoa(numSupportingInfo)+"]ValueString", resource.SupportingInfo[numSupportingInfo].ValueString, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoReason(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodingSelect("SupportingInfo["+strconv.Itoa(numSupportingInfo)+"]Reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("SupportingInfo["+strconv.Itoa(numSupportingInfo)+"]Reason", resource.SupportingInfo[numSupportingInfo].Reason, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_DiagnosisSequence(numDiagnosis int, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return IntInput("Diagnosis["+strconv.Itoa(numDiagnosis)+"]Sequence", nil, htmlAttrs)
	}
	return IntInput("Diagnosis["+strconv.Itoa(numDiagnosis)+"]Sequence", &resource.Diagnosis[numDiagnosis].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_DiagnosisDiagnosisCodeableConcept(numDiagnosis int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("Diagnosis["+strconv.Itoa(numDiagnosis)+"]DiagnosisCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Diagnosis["+strconv.Itoa(numDiagnosis)+"]DiagnosisCodeableConcept", &resource.Diagnosis[numDiagnosis].DiagnosisCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_DiagnosisType(numDiagnosis int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) || numType >= len(resource.Diagnosis[numDiagnosis].Type) {
		return CodeableConceptSelect("Diagnosis["+strconv.Itoa(numDiagnosis)+"]Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Diagnosis["+strconv.Itoa(numDiagnosis)+"]Type["+strconv.Itoa(numType)+"]", &resource.Diagnosis[numDiagnosis].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_DiagnosisOnAdmission(numDiagnosis int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("Diagnosis["+strconv.Itoa(numDiagnosis)+"]OnAdmission", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Diagnosis["+strconv.Itoa(numDiagnosis)+"]OnAdmission", resource.Diagnosis[numDiagnosis].OnAdmission, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_DiagnosisPackageCode(numDiagnosis int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("Diagnosis["+strconv.Itoa(numDiagnosis)+"]PackageCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Diagnosis["+strconv.Itoa(numDiagnosis)+"]PackageCode", resource.Diagnosis[numDiagnosis].PackageCode, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcedureSequence(numProcedure int, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return IntInput("Procedure["+strconv.Itoa(numProcedure)+"]Sequence", nil, htmlAttrs)
	}
	return IntInput("Procedure["+strconv.Itoa(numProcedure)+"]Sequence", &resource.Procedure[numProcedure].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcedureType(numProcedure int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) || numType >= len(resource.Procedure[numProcedure].Type) {
		return CodeableConceptSelect("Procedure["+strconv.Itoa(numProcedure)+"]Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Procedure["+strconv.Itoa(numProcedure)+"]Type["+strconv.Itoa(numType)+"]", &resource.Procedure[numProcedure].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcedureDate(numProcedure int, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return DateTimeInput("Procedure["+strconv.Itoa(numProcedure)+"]Date", nil, htmlAttrs)
	}
	return DateTimeInput("Procedure["+strconv.Itoa(numProcedure)+"]Date", resource.Procedure[numProcedure].Date, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcedureProcedureCodeableConcept(numProcedure int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return CodeableConceptSelect("Procedure["+strconv.Itoa(numProcedure)+"]ProcedureCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Procedure["+strconv.Itoa(numProcedure)+"]ProcedureCodeableConcept", &resource.Procedure[numProcedure].ProcedureCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_InsuranceFocal(numInsurance int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return BoolInput("Insurance["+strconv.Itoa(numInsurance)+"]Focal", nil, htmlAttrs)
	}
	return BoolInput("Insurance["+strconv.Itoa(numInsurance)+"]Focal", &resource.Insurance[numInsurance].Focal, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_InsurancePreAuthRef(numInsurance int, numPreAuthRef int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numPreAuthRef >= len(resource.Insurance[numInsurance].PreAuthRef) {
		return StringInput("Insurance["+strconv.Itoa(numInsurance)+"]PreAuthRef["+strconv.Itoa(numPreAuthRef)+"]", nil, htmlAttrs)
	}
	return StringInput("Insurance["+strconv.Itoa(numInsurance)+"]PreAuthRef["+strconv.Itoa(numPreAuthRef)+"]", &resource.Insurance[numInsurance].PreAuthRef[numPreAuthRef], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AccidentDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("AccidentDate", nil, htmlAttrs)
	}
	return DateInput("AccidentDate", resource.Accident.Date, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AccidentType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("AccidentType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AccidentType", resource.Accident.Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemSequence(numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return IntInput("Item["+strconv.Itoa(numItem)+"]Sequence", nil, htmlAttrs)
	}
	return IntInput("Item["+strconv.Itoa(numItem)+"]Sequence", &resource.Item[numItem].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemCareTeamSequence(numItem int, numCareTeamSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numCareTeamSequence >= len(resource.Item[numItem].CareTeamSequence) {
		return IntInput("Item["+strconv.Itoa(numItem)+"]CareTeamSequence["+strconv.Itoa(numCareTeamSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("Item["+strconv.Itoa(numItem)+"]CareTeamSequence["+strconv.Itoa(numCareTeamSequence)+"]", &resource.Item[numItem].CareTeamSequence[numCareTeamSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDiagnosisSequence(numItem int, numDiagnosisSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDiagnosisSequence >= len(resource.Item[numItem].DiagnosisSequence) {
		return IntInput("Item["+strconv.Itoa(numItem)+"]DiagnosisSequence["+strconv.Itoa(numDiagnosisSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("Item["+strconv.Itoa(numItem)+"]DiagnosisSequence["+strconv.Itoa(numDiagnosisSequence)+"]", &resource.Item[numItem].DiagnosisSequence[numDiagnosisSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemProcedureSequence(numItem int, numProcedureSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numProcedureSequence >= len(resource.Item[numItem].ProcedureSequence) {
		return IntInput("Item["+strconv.Itoa(numItem)+"]ProcedureSequence["+strconv.Itoa(numProcedureSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("Item["+strconv.Itoa(numItem)+"]ProcedureSequence["+strconv.Itoa(numProcedureSequence)+"]", &resource.Item[numItem].ProcedureSequence[numProcedureSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemInformationSequence(numItem int, numInformationSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numInformationSequence >= len(resource.Item[numItem].InformationSequence) {
		return IntInput("Item["+strconv.Itoa(numItem)+"]InformationSequence["+strconv.Itoa(numInformationSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("Item["+strconv.Itoa(numItem)+"]InformationSequence["+strconv.Itoa(numInformationSequence)+"]", &resource.Item[numItem].InformationSequence[numInformationSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemRevenue(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Revenue", resource.Item[numItem].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemCategory(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Category", resource.Item[numItem].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemProductOrService(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]ProductOrService", &resource.Item[numItem].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemModifier(numItem int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numModifier >= len(resource.Item[numItem].Modifier) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemProgramCode(numItem int, numProgramCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numProgramCode >= len(resource.Item[numItem].ProgramCode) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]ProgramCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]ProgramCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemServicedDate(numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return DateInput("Item["+strconv.Itoa(numItem)+"]ServicedDate", nil, htmlAttrs)
	}
	return DateInput("Item["+strconv.Itoa(numItem)+"]ServicedDate", resource.Item[numItem].ServicedDate, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemLocationCodeableConcept(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]LocationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]LocationCodeableConcept", resource.Item[numItem].LocationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemFactor(numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return Float64Input("Item["+strconv.Itoa(numItem)+"]Factor", nil, htmlAttrs)
	}
	return Float64Input("Item["+strconv.Itoa(numItem)+"]Factor", resource.Item[numItem].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemBodySite(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]BodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]BodySite", resource.Item[numItem].BodySite, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemSubSite(numItem int, numSubSite int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numSubSite >= len(resource.Item[numItem].SubSite) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]SubSite["+strconv.Itoa(numSubSite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]SubSite["+strconv.Itoa(numSubSite)+"]", &resource.Item[numItem].SubSite[numSubSite], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemNoteNumber(numItem int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numNoteNumber >= len(resource.Item[numItem].NoteNumber) {
		return IntInput("Item["+strconv.Itoa(numItem)+"]NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("Item["+strconv.Itoa(numItem)+"]NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemAdjudicationCategory(numItem int, numAdjudication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Adjudication["+strconv.Itoa(numAdjudication)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Adjudication["+strconv.Itoa(numAdjudication)+"].Category", &resource.Item[numItem].Adjudication[numAdjudication].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemAdjudicationReason(numItem int, numAdjudication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Adjudication["+strconv.Itoa(numAdjudication)+"].Reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Adjudication["+strconv.Itoa(numAdjudication)+"].Reason", resource.Item[numItem].Adjudication[numAdjudication].Reason, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemAdjudicationValue(numItem int, numAdjudication int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return Float64Input("Item["+strconv.Itoa(numItem)+"]Adjudication["+strconv.Itoa(numAdjudication)+"].Value", nil, htmlAttrs)
	}
	return Float64Input("Item["+strconv.Itoa(numItem)+"]Adjudication["+strconv.Itoa(numAdjudication)+"].Value", resource.Item[numItem].Adjudication[numAdjudication].Value, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSequence(numItem int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return IntInput("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].Sequence", nil, htmlAttrs)
	}
	return IntInput("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].Sequence", &resource.Item[numItem].Detail[numDetail].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailRevenue(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].Revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].Revenue", resource.Item[numItem].Detail[numDetail].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailCategory(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].Category", resource.Item[numItem].Detail[numDetail].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailProductOrService(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].ProductOrService", &resource.Item[numItem].Detail[numDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailModifier(numItem int, numDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numModifier >= len(resource.Item[numItem].Detail[numDetail].Modifier) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Detail[numDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailProgramCode(numItem int, numDetail int, numProgramCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numProgramCode >= len(resource.Item[numItem].Detail[numDetail].ProgramCode) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].Detail[numDetail].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailFactor(numItem int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return Float64Input("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].Factor", nil, htmlAttrs)
	}
	return Float64Input("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].Factor", resource.Item[numItem].Detail[numDetail].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailNoteNumber(numItem int, numDetail int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numNoteNumber >= len(resource.Item[numItem].Detail[numDetail].NoteNumber) {
		return IntInput("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].Detail[numDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailSequence(numItem int, numDetail int, numSubDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return IntInput("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Sequence", nil, htmlAttrs)
	}
	return IntInput("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Sequence", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailRevenue(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Revenue", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailCategory(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Category", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailProductOrService(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailModifier(numItem int, numDetail int, numSubDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numModifier >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Modifier) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailProgramCode(numItem int, numDetail int, numSubDetail int, numProgramCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numProgramCode >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProgramCode) {
		return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailFactor(numItem int, numDetail int, numSubDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return Float64Input("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Factor", nil, htmlAttrs)
	}
	return Float64Input("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Factor", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailNoteNumber(numItem int, numDetail int, numSubDetail int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numNoteNumber >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber) {
		return IntInput("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("Item["+strconv.Itoa(numItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemItemSequence(numAddItem int, numItemSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numItemSequence >= len(resource.AddItem[numAddItem].ItemSequence) {
		return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]ItemSequence["+strconv.Itoa(numItemSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]ItemSequence["+strconv.Itoa(numItemSequence)+"]", &resource.AddItem[numAddItem].ItemSequence[numItemSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSequence(numAddItem int, numDetailSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetailSequence >= len(resource.AddItem[numAddItem].DetailSequence) {
		return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]DetailSequence["+strconv.Itoa(numDetailSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]DetailSequence["+strconv.Itoa(numDetailSequence)+"]", &resource.AddItem[numAddItem].DetailSequence[numDetailSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemSubDetailSequence(numAddItem int, numSubDetailSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numSubDetailSequence >= len(resource.AddItem[numAddItem].SubDetailSequence) {
		return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]SubDetailSequence["+strconv.Itoa(numSubDetailSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]SubDetailSequence["+strconv.Itoa(numSubDetailSequence)+"]", &resource.AddItem[numAddItem].SubDetailSequence[numSubDetailSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemProductOrService(numAddItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]ProductOrService", &resource.AddItem[numAddItem].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemModifier(numAddItem int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numModifier >= len(resource.AddItem[numAddItem].Modifier) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemProgramCode(numAddItem int, numProgramCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numProgramCode >= len(resource.AddItem[numAddItem].ProgramCode) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]ProgramCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]ProgramCode["+strconv.Itoa(numProgramCode)+"]", &resource.AddItem[numAddItem].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemServicedDate(numAddItem int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return DateInput("AddItem["+strconv.Itoa(numAddItem)+"]ServicedDate", nil, htmlAttrs)
	}
	return DateInput("AddItem["+strconv.Itoa(numAddItem)+"]ServicedDate", resource.AddItem[numAddItem].ServicedDate, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemLocationCodeableConcept(numAddItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]LocationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]LocationCodeableConcept", resource.AddItem[numAddItem].LocationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemFactor(numAddItem int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return Float64Input("AddItem["+strconv.Itoa(numAddItem)+"]Factor", nil, htmlAttrs)
	}
	return Float64Input("AddItem["+strconv.Itoa(numAddItem)+"]Factor", resource.AddItem[numAddItem].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemBodySite(numAddItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]BodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]BodySite", resource.AddItem[numAddItem].BodySite, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemSubSite(numAddItem int, numSubSite int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numSubSite >= len(resource.AddItem[numAddItem].SubSite) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]SubSite["+strconv.Itoa(numSubSite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]SubSite["+strconv.Itoa(numSubSite)+"]", &resource.AddItem[numAddItem].SubSite[numSubSite], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemNoteNumber(numAddItem int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numNoteNumber >= len(resource.AddItem[numAddItem].NoteNumber) {
		return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailProductOrService(numAddItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].ProductOrService", &resource.AddItem[numAddItem].Detail[numDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailModifier(numAddItem int, numDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numModifier >= len(resource.AddItem[numAddItem].Detail[numDetail].Modifier) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Detail[numDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailFactor(numAddItem int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return Float64Input("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].Factor", nil, htmlAttrs)
	}
	return Float64Input("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].Factor", resource.AddItem[numAddItem].Detail[numDetail].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailNoteNumber(numAddItem int, numDetail int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numNoteNumber >= len(resource.AddItem[numAddItem].Detail[numDetail].NoteNumber) {
		return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailProductOrService(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailModifier(numAddItem int, numDetail int, numSubDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) || numModifier >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier) {
		return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailFactor(numAddItem int, numDetail int, numSubDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return Float64Input("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Factor", nil, htmlAttrs)
	}
	return Float64Input("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Factor", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailNoteNumber(numAddItem int, numDetail int, numSubDetail int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) || numNoteNumber >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber) {
		return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("AddItem["+strconv.Itoa(numAddItem)+"]Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_TotalCategory(numTotal int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTotal >= len(resource.Total) {
		return CodeableConceptSelect("Total["+strconv.Itoa(numTotal)+"]Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Total["+strconv.Itoa(numTotal)+"]Category", &resource.Total[numTotal].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PaymentType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PaymentType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PaymentType", resource.Payment.Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PaymentAdjustmentReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("PaymentAdjustmentReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("PaymentAdjustmentReason", resource.Payment.AdjustmentReason, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PaymentDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("PaymentDate", nil, htmlAttrs)
	}
	return DateInput("PaymentDate", resource.Payment.Date, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcessNoteNumber(numProcessNote int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return IntInput("ProcessNote["+strconv.Itoa(numProcessNote)+"]Number", nil, htmlAttrs)
	}
	return IntInput("ProcessNote["+strconv.Itoa(numProcessNote)+"]Number", resource.ProcessNote[numProcessNote].Number, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcessNoteType(numProcessNote int, htmlAttrs string) templ.Component {
	optionsValueSet := VSNote_type

	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return CodeSelect("ProcessNote["+strconv.Itoa(numProcessNote)+"]Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("ProcessNote["+strconv.Itoa(numProcessNote)+"]Type", resource.ProcessNote[numProcessNote].Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcessNoteText(numProcessNote int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return StringInput("ProcessNote["+strconv.Itoa(numProcessNote)+"]Text", nil, htmlAttrs)
	}
	return StringInput("ProcessNote["+strconv.Itoa(numProcessNote)+"]Text", resource.ProcessNote[numProcessNote].Text, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceCategory(numBenefitBalance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return CodeableConceptSelect("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Category", &resource.BenefitBalance[numBenefitBalance].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceExcluded(numBenefitBalance int, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return BoolInput("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Excluded", nil, htmlAttrs)
	}
	return BoolInput("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Excluded", resource.BenefitBalance[numBenefitBalance].Excluded, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceName(numBenefitBalance int, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return StringInput("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Name", nil, htmlAttrs)
	}
	return StringInput("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Name", resource.BenefitBalance[numBenefitBalance].Name, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceDescription(numBenefitBalance int, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return StringInput("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Description", nil, htmlAttrs)
	}
	return StringInput("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Description", resource.BenefitBalance[numBenefitBalance].Description, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceNetwork(numBenefitBalance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return CodeableConceptSelect("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Network", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Network", resource.BenefitBalance[numBenefitBalance].Network, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceUnit(numBenefitBalance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return CodeableConceptSelect("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Unit", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Unit", resource.BenefitBalance[numBenefitBalance].Unit, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceTerm(numBenefitBalance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return CodeableConceptSelect("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Term", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Term", resource.BenefitBalance[numBenefitBalance].Term, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialType(numBenefitBalance int, numFinancial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) || numFinancial >= len(resource.BenefitBalance[numBenefitBalance].Financial) {
		return CodeableConceptSelect("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Financial["+strconv.Itoa(numFinancial)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Financial["+strconv.Itoa(numFinancial)+"].Type", &resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialAllowedUnsignedInt(numBenefitBalance int, numFinancial int, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) || numFinancial >= len(resource.BenefitBalance[numBenefitBalance].Financial) {
		return IntInput("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Financial["+strconv.Itoa(numFinancial)+"].AllowedUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Financial["+strconv.Itoa(numFinancial)+"].AllowedUnsignedInt", resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].AllowedUnsignedInt, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialAllowedString(numBenefitBalance int, numFinancial int, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) || numFinancial >= len(resource.BenefitBalance[numBenefitBalance].Financial) {
		return StringInput("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Financial["+strconv.Itoa(numFinancial)+"].AllowedString", nil, htmlAttrs)
	}
	return StringInput("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Financial["+strconv.Itoa(numFinancial)+"].AllowedString", resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].AllowedString, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialUsedUnsignedInt(numBenefitBalance int, numFinancial int, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) || numFinancial >= len(resource.BenefitBalance[numBenefitBalance].Financial) {
		return IntInput("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Financial["+strconv.Itoa(numFinancial)+"].UsedUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("BenefitBalance["+strconv.Itoa(numBenefitBalance)+"]Financial["+strconv.Itoa(numFinancial)+"].UsedUnsignedInt", resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].UsedUnsignedInt, htmlAttrs)
}
