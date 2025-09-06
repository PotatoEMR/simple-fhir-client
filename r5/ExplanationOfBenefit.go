package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
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
	TraceNumber           []Identifier                         `json:"traceNumber,omitempty"`
	Status                string                               `json:"status"`
	Type                  CodeableConcept                      `json:"type"`
	SubType               *CodeableConcept                     `json:"subType,omitempty"`
	Use                   string                               `json:"use"`
	Patient               Reference                            `json:"patient"`
	BillablePeriod        *Period                              `json:"billablePeriod,omitempty"`
	Created               string                               `json:"created"`
	Enterer               *Reference                           `json:"enterer,omitempty"`
	Insurer               *Reference                           `json:"insurer,omitempty"`
	Provider              *Reference                           `json:"provider,omitempty"`
	Priority              *CodeableConcept                     `json:"priority,omitempty"`
	FundsReserveRequested *CodeableConcept                     `json:"fundsReserveRequested,omitempty"`
	FundsReserve          *CodeableConcept                     `json:"fundsReserve,omitempty"`
	Related               []ExplanationOfBenefitRelated        `json:"related,omitempty"`
	Prescription          *Reference                           `json:"prescription,omitempty"`
	OriginalPrescription  *Reference                           `json:"originalPrescription,omitempty"`
	Event                 []ExplanationOfBenefitEvent          `json:"event,omitempty"`
	Payee                 *ExplanationOfBenefitPayee           `json:"payee,omitempty"`
	Referral              *Reference                           `json:"referral,omitempty"`
	Encounter             []Reference                          `json:"encounter,omitempty"`
	Facility              *Reference                           `json:"facility,omitempty"`
	Claim                 *Reference                           `json:"claim,omitempty"`
	ClaimResponse         *Reference                           `json:"claimResponse,omitempty"`
	Outcome               string                               `json:"outcome"`
	Decision              *CodeableConcept                     `json:"decision,omitempty"`
	Disposition           *string                              `json:"disposition,omitempty"`
	PreAuthRef            []string                             `json:"preAuthRef,omitempty"`
	PreAuthRefPeriod      []Period                             `json:"preAuthRefPeriod,omitempty"`
	DiagnosisRelatedGroup *CodeableConcept                     `json:"diagnosisRelatedGroup,omitempty"`
	CareTeam              []ExplanationOfBenefitCareTeam       `json:"careTeam,omitempty"`
	SupportingInfo        []ExplanationOfBenefitSupportingInfo `json:"supportingInfo,omitempty"`
	Diagnosis             []ExplanationOfBenefitDiagnosis      `json:"diagnosis,omitempty"`
	Procedure             []ExplanationOfBenefitProcedure      `json:"procedure,omitempty"`
	Precedence            *int                                 `json:"precedence,omitempty"`
	Insurance             []ExplanationOfBenefitInsurance      `json:"insurance,omitempty"`
	Accident              *ExplanationOfBenefitAccident        `json:"accident,omitempty"`
	PatientPaid           *Money                               `json:"patientPaid,omitempty"`
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

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitRelated struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Claim             *Reference       `json:"claim,omitempty"`
	Relationship      *CodeableConcept `json:"relationship,omitempty"`
	Reference         *Identifier      `json:"reference,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitEvent struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	WhenDateTime      string          `json:"whenDateTime"`
	WhenPeriod        Period          `json:"whenPeriod"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitPayee struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Party             *Reference       `json:"party,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitCareTeam struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Sequence          int              `json:"sequence"`
	Provider          Reference        `json:"provider"`
	Responsible       *bool            `json:"responsible,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Specialty         *CodeableConcept `json:"specialty,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitSupportingInfo struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Sequence          int              `json:"sequence"`
	Category          CodeableConcept  `json:"category"`
	Code              *CodeableConcept `json:"code,omitempty"`
	TimingDate        *string          `json:"timingDate,omitempty"`
	TimingPeriod      *Period          `json:"timingPeriod,omitempty"`
	ValueBoolean      *bool            `json:"valueBoolean,omitempty"`
	ValueString       *string          `json:"valueString,omitempty"`
	ValueQuantity     *Quantity        `json:"valueQuantity,omitempty"`
	ValueAttachment   *Attachment      `json:"valueAttachment,omitempty"`
	ValueReference    *Reference       `json:"valueReference,omitempty"`
	ValueIdentifier   *Identifier      `json:"valueIdentifier,omitempty"`
	Reason            *Coding          `json:"reason,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitDiagnosis struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Sequence                 int               `json:"sequence"`
	DiagnosisCodeableConcept CodeableConcept   `json:"diagnosisCodeableConcept"`
	DiagnosisReference       Reference         `json:"diagnosisReference"`
	Type                     []CodeableConcept `json:"type,omitempty"`
	OnAdmission              *CodeableConcept  `json:"onAdmission,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitProcedure struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Sequence                 int               `json:"sequence"`
	Type                     []CodeableConcept `json:"type,omitempty"`
	Date                     *string           `json:"date,omitempty"`
	ProcedureCodeableConcept CodeableConcept   `json:"procedureCodeableConcept"`
	ProcedureReference       Reference         `json:"procedureReference"`
	Udi                      []Reference       `json:"udi,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitInsurance struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Focal             bool        `json:"focal"`
	Coverage          Reference   `json:"coverage"`
	PreAuthRef        []string    `json:"preAuthRef,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitAccident struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Date              *string          `json:"date,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	LocationAddress   *Address         `json:"locationAddress,omitempty"`
	LocationReference *Reference       `json:"locationReference,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitItem struct {
	Id                      *string                                `json:"id,omitempty"`
	Extension               []Extension                            `json:"extension,omitempty"`
	ModifierExtension       []Extension                            `json:"modifierExtension,omitempty"`
	Sequence                int                                    `json:"sequence"`
	CareTeamSequence        []int                                  `json:"careTeamSequence,omitempty"`
	DiagnosisSequence       []int                                  `json:"diagnosisSequence,omitempty"`
	ProcedureSequence       []int                                  `json:"procedureSequence,omitempty"`
	InformationSequence     []int                                  `json:"informationSequence,omitempty"`
	TraceNumber             []Identifier                           `json:"traceNumber,omitempty"`
	Revenue                 *CodeableConcept                       `json:"revenue,omitempty"`
	Category                *CodeableConcept                       `json:"category,omitempty"`
	ProductOrService        *CodeableConcept                       `json:"productOrService,omitempty"`
	ProductOrServiceEnd     *CodeableConcept                       `json:"productOrServiceEnd,omitempty"`
	Request                 []Reference                            `json:"request,omitempty"`
	Modifier                []CodeableConcept                      `json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept                      `json:"programCode,omitempty"`
	ServicedDate            *string                                `json:"servicedDate,omitempty"`
	ServicedPeriod          *Period                                `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept                       `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                               `json:"locationAddress,omitempty"`
	LocationReference       *Reference                             `json:"locationReference,omitempty"`
	PatientPaid             *Money                                 `json:"patientPaid,omitempty"`
	Quantity                *Quantity                              `json:"quantity,omitempty"`
	UnitPrice               *Money                                 `json:"unitPrice,omitempty"`
	Factor                  *float64                               `json:"factor,omitempty"`
	Tax                     *Money                                 `json:"tax,omitempty"`
	Net                     *Money                                 `json:"net,omitempty"`
	Udi                     []Reference                            `json:"udi,omitempty"`
	BodySite                []ExplanationOfBenefitItemBodySite     `json:"bodySite,omitempty"`
	Encounter               []Reference                            `json:"encounter,omitempty"`
	NoteNumber              []int                                  `json:"noteNumber,omitempty"`
	ReviewOutcome           *ExplanationOfBenefitItemReviewOutcome `json:"reviewOutcome,omitempty"`
	Adjudication            []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`
	Detail                  []ExplanationOfBenefitItemDetail       `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitItemBodySite struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Site              []CodeableReference `json:"site"`
	SubSite           []CodeableConcept   `json:"subSite,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitItemReviewOutcome struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Decision          *CodeableConcept  `json:"decision,omitempty"`
	Reason            []CodeableConcept `json:"reason,omitempty"`
	PreAuthRef        *string           `json:"preAuthRef,omitempty"`
	PreAuthPeriod     *Period           `json:"preAuthPeriod,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitItemAdjudication struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Category          CodeableConcept  `json:"category"`
	Reason            *CodeableConcept `json:"reason,omitempty"`
	Amount            *Money           `json:"amount,omitempty"`
	Quantity          *Quantity        `json:"quantity,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitItemDetail struct {
	Id                  *string                                   `json:"id,omitempty"`
	Extension           []Extension                               `json:"extension,omitempty"`
	ModifierExtension   []Extension                               `json:"modifierExtension,omitempty"`
	Sequence            int                                       `json:"sequence"`
	TraceNumber         []Identifier                              `json:"traceNumber,omitempty"`
	Revenue             *CodeableConcept                          `json:"revenue,omitempty"`
	Category            *CodeableConcept                          `json:"category,omitempty"`
	ProductOrService    *CodeableConcept                          `json:"productOrService,omitempty"`
	ProductOrServiceEnd *CodeableConcept                          `json:"productOrServiceEnd,omitempty"`
	Modifier            []CodeableConcept                         `json:"modifier,omitempty"`
	ProgramCode         []CodeableConcept                         `json:"programCode,omitempty"`
	PatientPaid         *Money                                    `json:"patientPaid,omitempty"`
	Quantity            *Quantity                                 `json:"quantity,omitempty"`
	UnitPrice           *Money                                    `json:"unitPrice,omitempty"`
	Factor              *float64                                  `json:"factor,omitempty"`
	Tax                 *Money                                    `json:"tax,omitempty"`
	Net                 *Money                                    `json:"net,omitempty"`
	Udi                 []Reference                               `json:"udi,omitempty"`
	NoteNumber          []int                                     `json:"noteNumber,omitempty"`
	SubDetail           []ExplanationOfBenefitItemDetailSubDetail `json:"subDetail,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitItemDetailSubDetail struct {
	Id                  *string           `json:"id,omitempty"`
	Extension           []Extension       `json:"extension,omitempty"`
	ModifierExtension   []Extension       `json:"modifierExtension,omitempty"`
	Sequence            int               `json:"sequence"`
	TraceNumber         []Identifier      `json:"traceNumber,omitempty"`
	Revenue             *CodeableConcept  `json:"revenue,omitempty"`
	Category            *CodeableConcept  `json:"category,omitempty"`
	ProductOrService    *CodeableConcept  `json:"productOrService,omitempty"`
	ProductOrServiceEnd *CodeableConcept  `json:"productOrServiceEnd,omitempty"`
	Modifier            []CodeableConcept `json:"modifier,omitempty"`
	ProgramCode         []CodeableConcept `json:"programCode,omitempty"`
	PatientPaid         *Money            `json:"patientPaid,omitempty"`
	Quantity            *Quantity         `json:"quantity,omitempty"`
	UnitPrice           *Money            `json:"unitPrice,omitempty"`
	Factor              *float64          `json:"factor,omitempty"`
	Tax                 *Money            `json:"tax,omitempty"`
	Net                 *Money            `json:"net,omitempty"`
	Udi                 []Reference       `json:"udi,omitempty"`
	NoteNumber          []int             `json:"noteNumber,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitAddItem struct {
	Id                      *string                               `json:"id,omitempty"`
	Extension               []Extension                           `json:"extension,omitempty"`
	ModifierExtension       []Extension                           `json:"modifierExtension,omitempty"`
	ItemSequence            []int                                 `json:"itemSequence,omitempty"`
	DetailSequence          []int                                 `json:"detailSequence,omitempty"`
	SubDetailSequence       []int                                 `json:"subDetailSequence,omitempty"`
	TraceNumber             []Identifier                          `json:"traceNumber,omitempty"`
	Provider                []Reference                           `json:"provider,omitempty"`
	Revenue                 *CodeableConcept                      `json:"revenue,omitempty"`
	ProductOrService        *CodeableConcept                      `json:"productOrService,omitempty"`
	ProductOrServiceEnd     *CodeableConcept                      `json:"productOrServiceEnd,omitempty"`
	Request                 []Reference                           `json:"request,omitempty"`
	Modifier                []CodeableConcept                     `json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept                     `json:"programCode,omitempty"`
	ServicedDate            *string                               `json:"servicedDate,omitempty"`
	ServicedPeriod          *Period                               `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept                      `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                              `json:"locationAddress,omitempty"`
	LocationReference       *Reference                            `json:"locationReference,omitempty"`
	PatientPaid             *Money                                `json:"patientPaid,omitempty"`
	Quantity                *Quantity                             `json:"quantity,omitempty"`
	UnitPrice               *Money                                `json:"unitPrice,omitempty"`
	Factor                  *float64                              `json:"factor,omitempty"`
	Tax                     *Money                                `json:"tax,omitempty"`
	Net                     *Money                                `json:"net,omitempty"`
	BodySite                []ExplanationOfBenefitAddItemBodySite `json:"bodySite,omitempty"`
	NoteNumber              []int                                 `json:"noteNumber,omitempty"`
	Detail                  []ExplanationOfBenefitAddItemDetail   `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitAddItemBodySite struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Site              []CodeableReference `json:"site"`
	SubSite           []CodeableConcept   `json:"subSite,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitAddItemDetail struct {
	Id                  *string                                      `json:"id,omitempty"`
	Extension           []Extension                                  `json:"extension,omitempty"`
	ModifierExtension   []Extension                                  `json:"modifierExtension,omitempty"`
	TraceNumber         []Identifier                                 `json:"traceNumber,omitempty"`
	Revenue             *CodeableConcept                             `json:"revenue,omitempty"`
	ProductOrService    *CodeableConcept                             `json:"productOrService,omitempty"`
	ProductOrServiceEnd *CodeableConcept                             `json:"productOrServiceEnd,omitempty"`
	Modifier            []CodeableConcept                            `json:"modifier,omitempty"`
	PatientPaid         *Money                                       `json:"patientPaid,omitempty"`
	Quantity            *Quantity                                    `json:"quantity,omitempty"`
	UnitPrice           *Money                                       `json:"unitPrice,omitempty"`
	Factor              *float64                                     `json:"factor,omitempty"`
	Tax                 *Money                                       `json:"tax,omitempty"`
	Net                 *Money                                       `json:"net,omitempty"`
	NoteNumber          []int                                        `json:"noteNumber,omitempty"`
	SubDetail           []ExplanationOfBenefitAddItemDetailSubDetail `json:"subDetail,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitAddItemDetailSubDetail struct {
	Id                  *string           `json:"id,omitempty"`
	Extension           []Extension       `json:"extension,omitempty"`
	ModifierExtension   []Extension       `json:"modifierExtension,omitempty"`
	TraceNumber         []Identifier      `json:"traceNumber,omitempty"`
	Revenue             *CodeableConcept  `json:"revenue,omitempty"`
	ProductOrService    *CodeableConcept  `json:"productOrService,omitempty"`
	ProductOrServiceEnd *CodeableConcept  `json:"productOrServiceEnd,omitempty"`
	Modifier            []CodeableConcept `json:"modifier,omitempty"`
	PatientPaid         *Money            `json:"patientPaid,omitempty"`
	Quantity            *Quantity         `json:"quantity,omitempty"`
	UnitPrice           *Money            `json:"unitPrice,omitempty"`
	Factor              *float64          `json:"factor,omitempty"`
	Tax                 *Money            `json:"tax,omitempty"`
	Net                 *Money            `json:"net,omitempty"`
	NoteNumber          []int             `json:"noteNumber,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitTotal struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Category          CodeableConcept `json:"category"`
	Amount            Money           `json:"amount"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitPayment struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Adjustment        *Money           `json:"adjustment,omitempty"`
	AdjustmentReason  *CodeableConcept `json:"adjustmentReason,omitempty"`
	Date              *string          `json:"date,omitempty"`
	Amount            *Money           `json:"amount,omitempty"`
	Identifier        *Identifier      `json:"identifier,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitProcessNote struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Number            *int             `json:"number,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Text              *string          `json:"text,omitempty"`
	Language          *CodeableConcept `json:"language,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
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

// http://hl7.org/fhir/r5/StructureDefinition/ExplanationOfBenefit
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

func (resource *ExplanationOfBenefit) T_Id() templ.Component {

	if resource == nil {
		return StringInput("ExplanationOfBenefit.Id", nil)
	}
	return StringInput("ExplanationOfBenefit.Id", resource.Id)
}
func (resource *ExplanationOfBenefit) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("ExplanationOfBenefit.ImplicitRules", nil)
	}
	return StringInput("ExplanationOfBenefit.ImplicitRules", resource.ImplicitRules)
}
func (resource *ExplanationOfBenefit) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("ExplanationOfBenefit.Language", nil, optionsValueSet)
	}
	return CodeSelect("ExplanationOfBenefit.Language", resource.Language, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_Status() templ.Component {
	optionsValueSet := VSExplanationofbenefit_status

	if resource == nil {
		return CodeSelect("ExplanationOfBenefit.Status", nil, optionsValueSet)
	}
	return CodeSelect("ExplanationOfBenefit.Status", &resource.Status, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ExplanationOfBenefit.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Type", &resource.Type, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_SubType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ExplanationOfBenefit.SubType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.SubType", resource.SubType, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_Use() templ.Component {
	optionsValueSet := VSClaim_use

	if resource == nil {
		return CodeSelect("ExplanationOfBenefit.Use", nil, optionsValueSet)
	}
	return CodeSelect("ExplanationOfBenefit.Use", &resource.Use, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_Created() templ.Component {

	if resource == nil {
		return StringInput("ExplanationOfBenefit.Created", nil)
	}
	return StringInput("ExplanationOfBenefit.Created", &resource.Created)
}
func (resource *ExplanationOfBenefit) T_Priority(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ExplanationOfBenefit.Priority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Priority", resource.Priority, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_FundsReserveRequested(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ExplanationOfBenefit.FundsReserveRequested", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.FundsReserveRequested", resource.FundsReserveRequested, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_FundsReserve(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ExplanationOfBenefit.FundsReserve", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.FundsReserve", resource.FundsReserve, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_Outcome() templ.Component {
	optionsValueSet := VSClaim_outcome

	if resource == nil {
		return CodeSelect("ExplanationOfBenefit.Outcome", nil, optionsValueSet)
	}
	return CodeSelect("ExplanationOfBenefit.Outcome", &resource.Outcome, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_Decision(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ExplanationOfBenefit.Decision", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Decision", resource.Decision, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_Disposition() templ.Component {

	if resource == nil {
		return StringInput("ExplanationOfBenefit.Disposition", nil)
	}
	return StringInput("ExplanationOfBenefit.Disposition", resource.Disposition)
}
func (resource *ExplanationOfBenefit) T_PreAuthRef(numPreAuthRef int) templ.Component {

	if resource == nil || len(resource.PreAuthRef) >= numPreAuthRef {
		return StringInput("ExplanationOfBenefit.PreAuthRef["+strconv.Itoa(numPreAuthRef)+"]", nil)
	}
	return StringInput("ExplanationOfBenefit.PreAuthRef["+strconv.Itoa(numPreAuthRef)+"]", &resource.PreAuthRef[numPreAuthRef])
}
func (resource *ExplanationOfBenefit) T_DiagnosisRelatedGroup(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ExplanationOfBenefit.DiagnosisRelatedGroup", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.DiagnosisRelatedGroup", resource.DiagnosisRelatedGroup, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_Precedence() templ.Component {

	if resource == nil {
		return IntInput("ExplanationOfBenefit.Precedence", nil)
	}
	return IntInput("ExplanationOfBenefit.Precedence", resource.Precedence)
}
func (resource *ExplanationOfBenefit) T_FormCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ExplanationOfBenefit.FormCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.FormCode", resource.FormCode, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_RelatedId(numRelated int) templ.Component {

	if resource == nil || len(resource.Related) >= numRelated {
		return StringInput("ExplanationOfBenefit.Related["+strconv.Itoa(numRelated)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.Related["+strconv.Itoa(numRelated)+"].Id", resource.Related[numRelated].Id)
}
func (resource *ExplanationOfBenefit) T_RelatedRelationship(numRelated int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Related) >= numRelated {
		return CodeableConceptSelect("ExplanationOfBenefit.Related["+strconv.Itoa(numRelated)+"].Relationship", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Related["+strconv.Itoa(numRelated)+"].Relationship", resource.Related[numRelated].Relationship, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_EventId(numEvent int) templ.Component {

	if resource == nil || len(resource.Event) >= numEvent {
		return StringInput("ExplanationOfBenefit.Event["+strconv.Itoa(numEvent)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.Event["+strconv.Itoa(numEvent)+"].Id", resource.Event[numEvent].Id)
}
func (resource *ExplanationOfBenefit) T_EventType(numEvent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Event) >= numEvent {
		return CodeableConceptSelect("ExplanationOfBenefit.Event["+strconv.Itoa(numEvent)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Event["+strconv.Itoa(numEvent)+"].Type", &resource.Event[numEvent].Type, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_PayeeId() templ.Component {

	if resource == nil {
		return StringInput("ExplanationOfBenefit.Payee.Id", nil)
	}
	return StringInput("ExplanationOfBenefit.Payee.Id", resource.Payee.Id)
}
func (resource *ExplanationOfBenefit) T_PayeeType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ExplanationOfBenefit.Payee.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Payee.Type", resource.Payee.Type, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_CareTeamId(numCareTeam int) templ.Component {

	if resource == nil || len(resource.CareTeam) >= numCareTeam {
		return StringInput("ExplanationOfBenefit.CareTeam["+strconv.Itoa(numCareTeam)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.CareTeam["+strconv.Itoa(numCareTeam)+"].Id", resource.CareTeam[numCareTeam].Id)
}
func (resource *ExplanationOfBenefit) T_CareTeamSequence(numCareTeam int) templ.Component {

	if resource == nil || len(resource.CareTeam) >= numCareTeam {
		return IntInput("ExplanationOfBenefit.CareTeam["+strconv.Itoa(numCareTeam)+"].Sequence", nil)
	}
	return IntInput("ExplanationOfBenefit.CareTeam["+strconv.Itoa(numCareTeam)+"].Sequence", &resource.CareTeam[numCareTeam].Sequence)
}
func (resource *ExplanationOfBenefit) T_CareTeamResponsible(numCareTeam int) templ.Component {

	if resource == nil || len(resource.CareTeam) >= numCareTeam {
		return BoolInput("ExplanationOfBenefit.CareTeam["+strconv.Itoa(numCareTeam)+"].Responsible", nil)
	}
	return BoolInput("ExplanationOfBenefit.CareTeam["+strconv.Itoa(numCareTeam)+"].Responsible", resource.CareTeam[numCareTeam].Responsible)
}
func (resource *ExplanationOfBenefit) T_CareTeamRole(numCareTeam int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CareTeam) >= numCareTeam {
		return CodeableConceptSelect("ExplanationOfBenefit.CareTeam["+strconv.Itoa(numCareTeam)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.CareTeam["+strconv.Itoa(numCareTeam)+"].Role", resource.CareTeam[numCareTeam].Role, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_CareTeamSpecialty(numCareTeam int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CareTeam) >= numCareTeam {
		return CodeableConceptSelect("ExplanationOfBenefit.CareTeam["+strconv.Itoa(numCareTeam)+"].Specialty", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.CareTeam["+strconv.Itoa(numCareTeam)+"].Specialty", resource.CareTeam[numCareTeam].Specialty, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoId(numSupportingInfo int) templ.Component {

	if resource == nil || len(resource.SupportingInfo) >= numSupportingInfo {
		return StringInput("ExplanationOfBenefit.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Id", resource.SupportingInfo[numSupportingInfo].Id)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoSequence(numSupportingInfo int) templ.Component {

	if resource == nil || len(resource.SupportingInfo) >= numSupportingInfo {
		return IntInput("ExplanationOfBenefit.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Sequence", nil)
	}
	return IntInput("ExplanationOfBenefit.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Sequence", &resource.SupportingInfo[numSupportingInfo].Sequence)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoCategory(numSupportingInfo int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SupportingInfo) >= numSupportingInfo {
		return CodeableConceptSelect("ExplanationOfBenefit.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Category", &resource.SupportingInfo[numSupportingInfo].Category, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoCode(numSupportingInfo int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SupportingInfo) >= numSupportingInfo {
		return CodeableConceptSelect("ExplanationOfBenefit.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Code", resource.SupportingInfo[numSupportingInfo].Code, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoReason(numSupportingInfo int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SupportingInfo) >= numSupportingInfo {
		return CodingSelect("ExplanationOfBenefit.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Reason", nil, optionsValueSet)
	}
	return CodingSelect("ExplanationOfBenefit.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Reason", resource.SupportingInfo[numSupportingInfo].Reason, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_DiagnosisId(numDiagnosis int) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis {
		return StringInput("ExplanationOfBenefit.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Id", resource.Diagnosis[numDiagnosis].Id)
}
func (resource *ExplanationOfBenefit) T_DiagnosisSequence(numDiagnosis int) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis {
		return IntInput("ExplanationOfBenefit.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Sequence", nil)
	}
	return IntInput("ExplanationOfBenefit.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Sequence", &resource.Diagnosis[numDiagnosis].Sequence)
}
func (resource *ExplanationOfBenefit) T_DiagnosisType(numDiagnosis int, numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis || len(resource.Diagnosis[numDiagnosis].Type) >= numType {
		return CodeableConceptSelect("ExplanationOfBenefit.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Type["+strconv.Itoa(numType)+"]", &resource.Diagnosis[numDiagnosis].Type[numType], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_DiagnosisOnAdmission(numDiagnosis int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis {
		return CodeableConceptSelect("ExplanationOfBenefit.Diagnosis["+strconv.Itoa(numDiagnosis)+"].OnAdmission", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Diagnosis["+strconv.Itoa(numDiagnosis)+"].OnAdmission", resource.Diagnosis[numDiagnosis].OnAdmission, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ProcedureId(numProcedure int) templ.Component {

	if resource == nil || len(resource.Procedure) >= numProcedure {
		return StringInput("ExplanationOfBenefit.Procedure["+strconv.Itoa(numProcedure)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.Procedure["+strconv.Itoa(numProcedure)+"].Id", resource.Procedure[numProcedure].Id)
}
func (resource *ExplanationOfBenefit) T_ProcedureSequence(numProcedure int) templ.Component {

	if resource == nil || len(resource.Procedure) >= numProcedure {
		return IntInput("ExplanationOfBenefit.Procedure["+strconv.Itoa(numProcedure)+"].Sequence", nil)
	}
	return IntInput("ExplanationOfBenefit.Procedure["+strconv.Itoa(numProcedure)+"].Sequence", &resource.Procedure[numProcedure].Sequence)
}
func (resource *ExplanationOfBenefit) T_ProcedureType(numProcedure int, numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Procedure) >= numProcedure || len(resource.Procedure[numProcedure].Type) >= numType {
		return CodeableConceptSelect("ExplanationOfBenefit.Procedure["+strconv.Itoa(numProcedure)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Procedure["+strconv.Itoa(numProcedure)+"].Type["+strconv.Itoa(numType)+"]", &resource.Procedure[numProcedure].Type[numType], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ProcedureDate(numProcedure int) templ.Component {

	if resource == nil || len(resource.Procedure) >= numProcedure {
		return StringInput("ExplanationOfBenefit.Procedure["+strconv.Itoa(numProcedure)+"].Date", nil)
	}
	return StringInput("ExplanationOfBenefit.Procedure["+strconv.Itoa(numProcedure)+"].Date", resource.Procedure[numProcedure].Date)
}
func (resource *ExplanationOfBenefit) T_InsuranceId(numInsurance int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance {
		return StringInput("ExplanationOfBenefit.Insurance["+strconv.Itoa(numInsurance)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.Insurance["+strconv.Itoa(numInsurance)+"].Id", resource.Insurance[numInsurance].Id)
}
func (resource *ExplanationOfBenefit) T_InsuranceFocal(numInsurance int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance {
		return BoolInput("ExplanationOfBenefit.Insurance["+strconv.Itoa(numInsurance)+"].Focal", nil)
	}
	return BoolInput("ExplanationOfBenefit.Insurance["+strconv.Itoa(numInsurance)+"].Focal", &resource.Insurance[numInsurance].Focal)
}
func (resource *ExplanationOfBenefit) T_InsurancePreAuthRef(numInsurance int, numPreAuthRef int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].PreAuthRef) >= numPreAuthRef {
		return StringInput("ExplanationOfBenefit.Insurance["+strconv.Itoa(numInsurance)+"].PreAuthRef["+strconv.Itoa(numPreAuthRef)+"]", nil)
	}
	return StringInput("ExplanationOfBenefit.Insurance["+strconv.Itoa(numInsurance)+"].PreAuthRef["+strconv.Itoa(numPreAuthRef)+"]", &resource.Insurance[numInsurance].PreAuthRef[numPreAuthRef])
}
func (resource *ExplanationOfBenefit) T_AccidentId() templ.Component {

	if resource == nil {
		return StringInput("ExplanationOfBenefit.Accident.Id", nil)
	}
	return StringInput("ExplanationOfBenefit.Accident.Id", resource.Accident.Id)
}
func (resource *ExplanationOfBenefit) T_AccidentDate() templ.Component {

	if resource == nil {
		return StringInput("ExplanationOfBenefit.Accident.Date", nil)
	}
	return StringInput("ExplanationOfBenefit.Accident.Date", resource.Accident.Date)
}
func (resource *ExplanationOfBenefit) T_AccidentType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ExplanationOfBenefit.Accident.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Accident.Type", resource.Accident.Type, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemId(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Id", resource.Item[numItem].Id)
}
func (resource *ExplanationOfBenefit) T_ItemSequence(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Sequence", nil)
	}
	return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Sequence", &resource.Item[numItem].Sequence)
}
func (resource *ExplanationOfBenefit) T_ItemCareTeamSequence(numItem int, numCareTeamSequence int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].CareTeamSequence) >= numCareTeamSequence {
		return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].CareTeamSequence["+strconv.Itoa(numCareTeamSequence)+"]", nil)
	}
	return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].CareTeamSequence["+strconv.Itoa(numCareTeamSequence)+"]", &resource.Item[numItem].CareTeamSequence[numCareTeamSequence])
}
func (resource *ExplanationOfBenefit) T_ItemDiagnosisSequence(numItem int, numDiagnosisSequence int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].DiagnosisSequence) >= numDiagnosisSequence {
		return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].DiagnosisSequence["+strconv.Itoa(numDiagnosisSequence)+"]", nil)
	}
	return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].DiagnosisSequence["+strconv.Itoa(numDiagnosisSequence)+"]", &resource.Item[numItem].DiagnosisSequence[numDiagnosisSequence])
}
func (resource *ExplanationOfBenefit) T_ItemProcedureSequence(numItem int, numProcedureSequence int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].ProcedureSequence) >= numProcedureSequence {
		return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].ProcedureSequence["+strconv.Itoa(numProcedureSequence)+"]", nil)
	}
	return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].ProcedureSequence["+strconv.Itoa(numProcedureSequence)+"]", &resource.Item[numItem].ProcedureSequence[numProcedureSequence])
}
func (resource *ExplanationOfBenefit) T_ItemInformationSequence(numItem int, numInformationSequence int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].InformationSequence) >= numInformationSequence {
		return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].InformationSequence["+strconv.Itoa(numInformationSequence)+"]", nil)
	}
	return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].InformationSequence["+strconv.Itoa(numInformationSequence)+"]", &resource.Item[numItem].InformationSequence[numInformationSequence])
}
func (resource *ExplanationOfBenefit) T_ItemRevenue(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Revenue", resource.Item[numItem].Revenue, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemCategory(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Category", resource.Item[numItem].Category, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemProductOrService(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].ProductOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].ProductOrService", resource.Item[numItem].ProductOrService, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemProductOrServiceEnd(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].ProductOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].ProductOrServiceEnd", resource.Item[numItem].ProductOrServiceEnd, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemModifier(numItem int, numModifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Modifier) >= numModifier {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Modifier[numModifier], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemProgramCode(numItem int, numProgramCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].ProgramCode) >= numProgramCode {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].ProgramCode[numProgramCode], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemFactor(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return Float64Input("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Factor", nil)
	}
	return Float64Input("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Factor", resource.Item[numItem].Factor)
}
func (resource *ExplanationOfBenefit) T_ItemNoteNumber(numItem int, numNoteNumber int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].NoteNumber) >= numNoteNumber {
		return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil)
	}
	return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].NoteNumber[numNoteNumber])
}
func (resource *ExplanationOfBenefit) T_ItemBodySiteId(numItem int, numBodySite int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].BodySite) >= numBodySite {
		return StringInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].BodySite["+strconv.Itoa(numBodySite)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].BodySite["+strconv.Itoa(numBodySite)+"].Id", resource.Item[numItem].BodySite[numBodySite].Id)
}
func (resource *ExplanationOfBenefit) T_ItemBodySiteSubSite(numItem int, numBodySite int, numSubSite int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].BodySite) >= numBodySite || len(resource.Item[numItem].BodySite[numBodySite].SubSite) >= numSubSite {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].BodySite["+strconv.Itoa(numBodySite)+"].SubSite["+strconv.Itoa(numSubSite)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].BodySite["+strconv.Itoa(numBodySite)+"].SubSite["+strconv.Itoa(numSubSite)+"]", &resource.Item[numItem].BodySite[numBodySite].SubSite[numSubSite], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemReviewOutcomeId(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].ReviewOutcome.Id", nil)
	}
	return StringInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].ReviewOutcome.Id", resource.Item[numItem].ReviewOutcome.Id)
}
func (resource *ExplanationOfBenefit) T_ItemReviewOutcomeDecision(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].ReviewOutcome.Decision", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].ReviewOutcome.Decision", resource.Item[numItem].ReviewOutcome.Decision, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemReviewOutcomeReason(numItem int, numReason int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].ReviewOutcome.Reason) >= numReason {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].ReviewOutcome.Reason["+strconv.Itoa(numReason)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].ReviewOutcome.Reason["+strconv.Itoa(numReason)+"]", &resource.Item[numItem].ReviewOutcome.Reason[numReason], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemReviewOutcomePreAuthRef(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].ReviewOutcome.PreAuthRef", nil)
	}
	return StringInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].ReviewOutcome.PreAuthRef", resource.Item[numItem].ReviewOutcome.PreAuthRef)
}
func (resource *ExplanationOfBenefit) T_ItemAdjudicationId(numItem int, numAdjudication int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Adjudication) >= numAdjudication {
		return StringInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Adjudication["+strconv.Itoa(numAdjudication)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Adjudication["+strconv.Itoa(numAdjudication)+"].Id", resource.Item[numItem].Adjudication[numAdjudication].Id)
}
func (resource *ExplanationOfBenefit) T_ItemAdjudicationCategory(numItem int, numAdjudication int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Adjudication) >= numAdjudication {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Adjudication["+strconv.Itoa(numAdjudication)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Adjudication["+strconv.Itoa(numAdjudication)+"].Category", &resource.Item[numItem].Adjudication[numAdjudication].Category, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemAdjudicationReason(numItem int, numAdjudication int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Adjudication) >= numAdjudication {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Adjudication["+strconv.Itoa(numAdjudication)+"].Reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Adjudication["+strconv.Itoa(numAdjudication)+"].Reason", resource.Item[numItem].Adjudication[numAdjudication].Reason, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailId(numItem int, numDetail int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail {
		return StringInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Id", resource.Item[numItem].Detail[numDetail].Id)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSequence(numItem int, numDetail int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail {
		return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Sequence", nil)
	}
	return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Sequence", &resource.Item[numItem].Detail[numDetail].Sequence)
}
func (resource *ExplanationOfBenefit) T_ItemDetailRevenue(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Revenue", resource.Item[numItem].Detail[numDetail].Revenue, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailCategory(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Category", resource.Item[numItem].Detail[numDetail].Category, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailProductOrService(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProductOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProductOrService", resource.Item[numItem].Detail[numDetail].ProductOrService, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailProductOrServiceEnd(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProductOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProductOrServiceEnd", resource.Item[numItem].Detail[numDetail].ProductOrServiceEnd, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailModifier(numItem int, numDetail int, numModifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].Modifier) >= numModifier {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Detail[numDetail].Modifier[numModifier], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailProgramCode(numItem int, numDetail int, numProgramCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].ProgramCode) >= numProgramCode {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].Detail[numDetail].ProgramCode[numProgramCode], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailFactor(numItem int, numDetail int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail {
		return Float64Input("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Factor", nil)
	}
	return Float64Input("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Factor", resource.Item[numItem].Detail[numDetail].Factor)
}
func (resource *ExplanationOfBenefit) T_ItemDetailNoteNumber(numItem int, numDetail int, numNoteNumber int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].NoteNumber) >= numNoteNumber {
		return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil)
	}
	return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].Detail[numDetail].NoteNumber[numNoteNumber])
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailId(numItem int, numDetail int, numSubDetail int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return StringInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Id", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Id)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailSequence(numItem int, numDetail int, numSubDetail int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Sequence", nil)
	}
	return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Sequence", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Sequence)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailRevenue(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Revenue", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Revenue, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailCategory(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Category", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Category, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailProductOrService(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailProductOrServiceEnd(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrServiceEnd", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrServiceEnd, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailModifier(numItem int, numDetail int, numSubDetail int, numModifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Modifier) >= numModifier {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[numModifier], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailProgramCode(numItem int, numDetail int, numSubDetail int, numProgramCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProgramCode) >= numProgramCode {
		return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProgramCode[numProgramCode], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailFactor(numItem int, numDetail int, numSubDetail int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return Float64Input("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Factor", nil)
	}
	return Float64Input("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Factor", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Factor)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailNoteNumber(numItem int, numDetail int, numSubDetail int, numNoteNumber int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber) >= numNoteNumber {
		return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil)
	}
	return IntInput("ExplanationOfBenefit.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber[numNoteNumber])
}
func (resource *ExplanationOfBenefit) T_AddItemId(numAddItem int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem {
		return StringInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Id", resource.AddItem[numAddItem].Id)
}
func (resource *ExplanationOfBenefit) T_AddItemItemSequence(numAddItem int, numItemSequence int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].ItemSequence) >= numItemSequence {
		return IntInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].ItemSequence["+strconv.Itoa(numItemSequence)+"]", nil)
	}
	return IntInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].ItemSequence["+strconv.Itoa(numItemSequence)+"]", &resource.AddItem[numAddItem].ItemSequence[numItemSequence])
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSequence(numAddItem int, numDetailSequence int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].DetailSequence) >= numDetailSequence {
		return IntInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].DetailSequence["+strconv.Itoa(numDetailSequence)+"]", nil)
	}
	return IntInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].DetailSequence["+strconv.Itoa(numDetailSequence)+"]", &resource.AddItem[numAddItem].DetailSequence[numDetailSequence])
}
func (resource *ExplanationOfBenefit) T_AddItemSubDetailSequence(numAddItem int, numSubDetailSequence int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].SubDetailSequence) >= numSubDetailSequence {
		return IntInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].SubDetailSequence["+strconv.Itoa(numSubDetailSequence)+"]", nil)
	}
	return IntInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].SubDetailSequence["+strconv.Itoa(numSubDetailSequence)+"]", &resource.AddItem[numAddItem].SubDetailSequence[numSubDetailSequence])
}
func (resource *ExplanationOfBenefit) T_AddItemRevenue(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Revenue", resource.AddItem[numAddItem].Revenue, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemProductOrService(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].ProductOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].ProductOrService", resource.AddItem[numAddItem].ProductOrService, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemProductOrServiceEnd(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].ProductOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].ProductOrServiceEnd", resource.AddItem[numAddItem].ProductOrServiceEnd, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemModifier(numAddItem int, numModifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Modifier) >= numModifier {
		return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Modifier[numModifier], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemProgramCode(numAddItem int, numProgramCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].ProgramCode) >= numProgramCode {
		return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", &resource.AddItem[numAddItem].ProgramCode[numProgramCode], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemFactor(numAddItem int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem {
		return Float64Input("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Factor", nil)
	}
	return Float64Input("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Factor", resource.AddItem[numAddItem].Factor)
}
func (resource *ExplanationOfBenefit) T_AddItemNoteNumber(numAddItem int, numNoteNumber int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].NoteNumber) >= numNoteNumber {
		return IntInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil)
	}
	return IntInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].NoteNumber[numNoteNumber])
}
func (resource *ExplanationOfBenefit) T_AddItemBodySiteId(numAddItem int, numBodySite int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].BodySite) >= numBodySite {
		return StringInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].BodySite["+strconv.Itoa(numBodySite)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].BodySite["+strconv.Itoa(numBodySite)+"].Id", resource.AddItem[numAddItem].BodySite[numBodySite].Id)
}
func (resource *ExplanationOfBenefit) T_AddItemBodySiteSubSite(numAddItem int, numBodySite int, numSubSite int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].BodySite) >= numBodySite || len(resource.AddItem[numAddItem].BodySite[numBodySite].SubSite) >= numSubSite {
		return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].BodySite["+strconv.Itoa(numBodySite)+"].SubSite["+strconv.Itoa(numSubSite)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].BodySite["+strconv.Itoa(numBodySite)+"].SubSite["+strconv.Itoa(numSubSite)+"]", &resource.AddItem[numAddItem].BodySite[numBodySite].SubSite[numSubSite], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailId(numAddItem int, numDetail int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return StringInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].Id", resource.AddItem[numAddItem].Detail[numDetail].Id)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailRevenue(numAddItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].Revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].Revenue", resource.AddItem[numAddItem].Detail[numDetail].Revenue, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailProductOrService(numAddItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProductOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProductOrService", resource.AddItem[numAddItem].Detail[numDetail].ProductOrService, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailProductOrServiceEnd(numAddItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProductOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProductOrServiceEnd", resource.AddItem[numAddItem].Detail[numDetail].ProductOrServiceEnd, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailModifier(numAddItem int, numDetail int, numModifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail || len(resource.AddItem[numAddItem].Detail[numDetail].Modifier) >= numModifier {
		return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Detail[numDetail].Modifier[numModifier], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailFactor(numAddItem int, numDetail int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return Float64Input("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].Factor", nil)
	}
	return Float64Input("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].Factor", resource.AddItem[numAddItem].Detail[numDetail].Factor)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailNoteNumber(numAddItem int, numDetail int, numNoteNumber int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail || len(resource.AddItem[numAddItem].Detail[numDetail].NoteNumber) >= numNoteNumber {
		return IntInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil)
	}
	return IntInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].NoteNumber[numNoteNumber])
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailId(numAddItem int, numDetail int, numSubDetail int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail || len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return StringInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Id", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Id)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailRevenue(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail || len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Revenue", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Revenue, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailProductOrService(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail || len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailProductOrServiceEnd(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail || len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrServiceEnd", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrServiceEnd, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailModifier(numAddItem int, numDetail int, numSubDetail int, numModifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail || len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail || len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier) >= numModifier {
		return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[numModifier], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailFactor(numAddItem int, numDetail int, numSubDetail int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail || len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return Float64Input("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Factor", nil)
	}
	return Float64Input("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Factor", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Factor)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailNoteNumber(numAddItem int, numDetail int, numSubDetail int, numNoteNumber int) templ.Component {

	if resource == nil || len(resource.AddItem) >= numAddItem || len(resource.AddItem[numAddItem].Detail) >= numDetail || len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail || len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber) >= numNoteNumber {
		return IntInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", nil)
	}
	return IntInput("ExplanationOfBenefit.AddItem["+strconv.Itoa(numAddItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].NoteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber[numNoteNumber])
}
func (resource *ExplanationOfBenefit) T_TotalId(numTotal int) templ.Component {

	if resource == nil || len(resource.Total) >= numTotal {
		return StringInput("ExplanationOfBenefit.Total["+strconv.Itoa(numTotal)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.Total["+strconv.Itoa(numTotal)+"].Id", resource.Total[numTotal].Id)
}
func (resource *ExplanationOfBenefit) T_TotalCategory(numTotal int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Total) >= numTotal {
		return CodeableConceptSelect("ExplanationOfBenefit.Total["+strconv.Itoa(numTotal)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Total["+strconv.Itoa(numTotal)+"].Category", &resource.Total[numTotal].Category, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_PaymentId() templ.Component {

	if resource == nil {
		return StringInput("ExplanationOfBenefit.Payment.Id", nil)
	}
	return StringInput("ExplanationOfBenefit.Payment.Id", resource.Payment.Id)
}
func (resource *ExplanationOfBenefit) T_PaymentType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ExplanationOfBenefit.Payment.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Payment.Type", resource.Payment.Type, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_PaymentAdjustmentReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("ExplanationOfBenefit.Payment.AdjustmentReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.Payment.AdjustmentReason", resource.Payment.AdjustmentReason, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_PaymentDate() templ.Component {

	if resource == nil {
		return StringInput("ExplanationOfBenefit.Payment.Date", nil)
	}
	return StringInput("ExplanationOfBenefit.Payment.Date", resource.Payment.Date)
}
func (resource *ExplanationOfBenefit) T_ProcessNoteId(numProcessNote int) templ.Component {

	if resource == nil || len(resource.ProcessNote) >= numProcessNote {
		return StringInput("ExplanationOfBenefit.ProcessNote["+strconv.Itoa(numProcessNote)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.ProcessNote["+strconv.Itoa(numProcessNote)+"].Id", resource.ProcessNote[numProcessNote].Id)
}
func (resource *ExplanationOfBenefit) T_ProcessNoteNumber(numProcessNote int) templ.Component {

	if resource == nil || len(resource.ProcessNote) >= numProcessNote {
		return IntInput("ExplanationOfBenefit.ProcessNote["+strconv.Itoa(numProcessNote)+"].Number", nil)
	}
	return IntInput("ExplanationOfBenefit.ProcessNote["+strconv.Itoa(numProcessNote)+"].Number", resource.ProcessNote[numProcessNote].Number)
}
func (resource *ExplanationOfBenefit) T_ProcessNoteType(numProcessNote int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ProcessNote) >= numProcessNote {
		return CodeableConceptSelect("ExplanationOfBenefit.ProcessNote["+strconv.Itoa(numProcessNote)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.ProcessNote["+strconv.Itoa(numProcessNote)+"].Type", resource.ProcessNote[numProcessNote].Type, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ProcessNoteText(numProcessNote int) templ.Component {

	if resource == nil || len(resource.ProcessNote) >= numProcessNote {
		return StringInput("ExplanationOfBenefit.ProcessNote["+strconv.Itoa(numProcessNote)+"].Text", nil)
	}
	return StringInput("ExplanationOfBenefit.ProcessNote["+strconv.Itoa(numProcessNote)+"].Text", resource.ProcessNote[numProcessNote].Text)
}
func (resource *ExplanationOfBenefit) T_ProcessNoteLanguage(numProcessNote int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ProcessNote) >= numProcessNote {
		return CodeableConceptSelect("ExplanationOfBenefit.ProcessNote["+strconv.Itoa(numProcessNote)+"].Language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.ProcessNote["+strconv.Itoa(numProcessNote)+"].Language", resource.ProcessNote[numProcessNote].Language, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceId(numBenefitBalance int) templ.Component {

	if resource == nil || len(resource.BenefitBalance) >= numBenefitBalance {
		return StringInput("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Id", resource.BenefitBalance[numBenefitBalance].Id)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceCategory(numBenefitBalance int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.BenefitBalance) >= numBenefitBalance {
		return CodeableConceptSelect("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Category", &resource.BenefitBalance[numBenefitBalance].Category, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceExcluded(numBenefitBalance int) templ.Component {

	if resource == nil || len(resource.BenefitBalance) >= numBenefitBalance {
		return BoolInput("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Excluded", nil)
	}
	return BoolInput("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Excluded", resource.BenefitBalance[numBenefitBalance].Excluded)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceName(numBenefitBalance int) templ.Component {

	if resource == nil || len(resource.BenefitBalance) >= numBenefitBalance {
		return StringInput("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Name", nil)
	}
	return StringInput("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Name", resource.BenefitBalance[numBenefitBalance].Name)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceDescription(numBenefitBalance int) templ.Component {

	if resource == nil || len(resource.BenefitBalance) >= numBenefitBalance {
		return StringInput("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Description", nil)
	}
	return StringInput("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Description", resource.BenefitBalance[numBenefitBalance].Description)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceNetwork(numBenefitBalance int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.BenefitBalance) >= numBenefitBalance {
		return CodeableConceptSelect("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Network", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Network", resource.BenefitBalance[numBenefitBalance].Network, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceUnit(numBenefitBalance int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.BenefitBalance) >= numBenefitBalance {
		return CodeableConceptSelect("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Unit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Unit", resource.BenefitBalance[numBenefitBalance].Unit, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceTerm(numBenefitBalance int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.BenefitBalance) >= numBenefitBalance {
		return CodeableConceptSelect("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Term", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Term", resource.BenefitBalance[numBenefitBalance].Term, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialId(numBenefitBalance int, numFinancial int) templ.Component {

	if resource == nil || len(resource.BenefitBalance) >= numBenefitBalance || len(resource.BenefitBalance[numBenefitBalance].Financial) >= numFinancial {
		return StringInput("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Financial["+strconv.Itoa(numFinancial)+"].Id", nil)
	}
	return StringInput("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Financial["+strconv.Itoa(numFinancial)+"].Id", resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].Id)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialType(numBenefitBalance int, numFinancial int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.BenefitBalance) >= numBenefitBalance || len(resource.BenefitBalance[numBenefitBalance].Financial) >= numFinancial {
		return CodeableConceptSelect("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Financial["+strconv.Itoa(numFinancial)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("ExplanationOfBenefit.BenefitBalance["+strconv.Itoa(numBenefitBalance)+"].Financial["+strconv.Itoa(numFinancial)+"].Type", &resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].Type, optionsValueSet)
}
