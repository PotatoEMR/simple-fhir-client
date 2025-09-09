package r5

//generated with command go run ./bultaoreune
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
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SubType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("subType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subType", resource.SubType, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Use(htmlAttrs string) templ.Component {
	optionsValueSet := VSClaim_use

	if resource == nil {
		return CodeSelect("use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("use", &resource.Use, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Created(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("created", nil, htmlAttrs)
	}
	return DateTimeInput("created", &resource.Created, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Priority(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_FundsReserveRequested(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("fundsReserveRequested", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("fundsReserveRequested", resource.FundsReserveRequested, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_FundsReserve(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("fundsReserve", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("fundsReserve", resource.FundsReserve, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Outcome(htmlAttrs string) templ.Component {
	optionsValueSet := VSClaim_outcome

	if resource == nil {
		return CodeSelect("outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("outcome", &resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Decision(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("decision", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("decision", resource.Decision, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Disposition(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("disposition", nil, htmlAttrs)
	}
	return StringInput("disposition", resource.Disposition, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PreAuthRef(numPreAuthRef int, htmlAttrs string) templ.Component {
	if resource == nil || numPreAuthRef >= len(resource.PreAuthRef) {
		return StringInput("preAuthRef["+strconv.Itoa(numPreAuthRef)+"]", nil, htmlAttrs)
	}
	return StringInput("preAuthRef["+strconv.Itoa(numPreAuthRef)+"]", &resource.PreAuthRef[numPreAuthRef], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_DiagnosisRelatedGroup(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("diagnosisRelatedGroup", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosisRelatedGroup", resource.DiagnosisRelatedGroup, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Precedence(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("precedence", nil, htmlAttrs)
	}
	return IntInput("precedence", resource.Precedence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_FormCode(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("formCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("formCode", resource.FormCode, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_RelatedRelationship(numRelated int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelated >= len(resource.Related) {
		return CodeableConceptSelect("related["+strconv.Itoa(numRelated)+"].relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("related["+strconv.Itoa(numRelated)+"].relationship", resource.Related[numRelated].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_EventType(numEvent int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return CodeableConceptSelect("event["+strconv.Itoa(numEvent)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("event["+strconv.Itoa(numEvent)+"].type", &resource.Event[numEvent].Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_EventWhenDateTime(numEvent int, htmlAttrs string) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return DateTimeInput("event["+strconv.Itoa(numEvent)+"].whenDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("event["+strconv.Itoa(numEvent)+"].whenDateTime", &resource.Event[numEvent].WhenDateTime, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PayeeType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("payee.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payee.type", resource.Payee.Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_CareTeamSequence(numCareTeam int, htmlAttrs string) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return IntInput("careTeam["+strconv.Itoa(numCareTeam)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("careTeam["+strconv.Itoa(numCareTeam)+"].sequence", &resource.CareTeam[numCareTeam].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_CareTeamResponsible(numCareTeam int, htmlAttrs string) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return BoolInput("careTeam["+strconv.Itoa(numCareTeam)+"].responsible", nil, htmlAttrs)
	}
	return BoolInput("careTeam["+strconv.Itoa(numCareTeam)+"].responsible", resource.CareTeam[numCareTeam].Responsible, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_CareTeamRole(numCareTeam int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return CodeableConceptSelect("careTeam["+strconv.Itoa(numCareTeam)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("careTeam["+strconv.Itoa(numCareTeam)+"].role", resource.CareTeam[numCareTeam].Role, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_CareTeamSpecialty(numCareTeam int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return CodeableConceptSelect("careTeam["+strconv.Itoa(numCareTeam)+"].specialty", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("careTeam["+strconv.Itoa(numCareTeam)+"].specialty", resource.CareTeam[numCareTeam].Specialty, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoSequence(numSupportingInfo int, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return IntInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].sequence", &resource.SupportingInfo[numSupportingInfo].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoCategory(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodeableConceptSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].category", &resource.SupportingInfo[numSupportingInfo].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoCode(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodeableConceptSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].code", resource.SupportingInfo[numSupportingInfo].Code, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoTimingDate(numSupportingInfo int, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return DateInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].timingDate", nil, htmlAttrs)
	}
	return DateInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].timingDate", resource.SupportingInfo[numSupportingInfo].TimingDate, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoValueBoolean(numSupportingInfo int, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return BoolInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueBoolean", resource.SupportingInfo[numSupportingInfo].ValueBoolean, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoValueString(numSupportingInfo int, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return StringInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueString", resource.SupportingInfo[numSupportingInfo].ValueString, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoReason(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodingSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].reason", resource.SupportingInfo[numSupportingInfo].Reason, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_DiagnosisSequence(numDiagnosis int, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return IntInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].sequence", &resource.Diagnosis[numDiagnosis].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_DiagnosisDiagnosisCodeableConcept(numDiagnosis int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].diagnosisCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].diagnosisCodeableConcept", &resource.Diagnosis[numDiagnosis].DiagnosisCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_DiagnosisType(numDiagnosis int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) || numType >= len(resource.Diagnosis[numDiagnosis].Type) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].type["+strconv.Itoa(numType)+"]", &resource.Diagnosis[numDiagnosis].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_DiagnosisOnAdmission(numDiagnosis int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].onAdmission", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].onAdmission", resource.Diagnosis[numDiagnosis].OnAdmission, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcedureSequence(numProcedure int, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return IntInput("procedure["+strconv.Itoa(numProcedure)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("procedure["+strconv.Itoa(numProcedure)+"].sequence", &resource.Procedure[numProcedure].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcedureType(numProcedure int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) || numType >= len(resource.Procedure[numProcedure].Type) {
		return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].type["+strconv.Itoa(numType)+"]", &resource.Procedure[numProcedure].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcedureDate(numProcedure int, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return DateTimeInput("procedure["+strconv.Itoa(numProcedure)+"].date", nil, htmlAttrs)
	}
	return DateTimeInput("procedure["+strconv.Itoa(numProcedure)+"].date", resource.Procedure[numProcedure].Date, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcedureProcedureCodeableConcept(numProcedure int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].procedureCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].procedureCodeableConcept", &resource.Procedure[numProcedure].ProcedureCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_InsuranceFocal(numInsurance int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].focal", nil, htmlAttrs)
	}
	return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].focal", &resource.Insurance[numInsurance].Focal, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_InsurancePreAuthRef(numInsurance int, numPreAuthRef int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numPreAuthRef >= len(resource.Insurance[numInsurance].PreAuthRef) {
		return StringInput("insurance["+strconv.Itoa(numInsurance)+"].preAuthRef["+strconv.Itoa(numPreAuthRef)+"]", nil, htmlAttrs)
	}
	return StringInput("insurance["+strconv.Itoa(numInsurance)+"].preAuthRef["+strconv.Itoa(numPreAuthRef)+"]", &resource.Insurance[numInsurance].PreAuthRef[numPreAuthRef], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AccidentDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("accident.date", nil, htmlAttrs)
	}
	return DateInput("accident.date", resource.Accident.Date, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AccidentType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("accident.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("accident.type", resource.Accident.Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemSequence(numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return IntInput("item["+strconv.Itoa(numItem)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].sequence", &resource.Item[numItem].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemCareTeamSequence(numItem int, numCareTeamSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numCareTeamSequence >= len(resource.Item[numItem].CareTeamSequence) {
		return IntInput("item["+strconv.Itoa(numItem)+"].careTeamSequence["+strconv.Itoa(numCareTeamSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].careTeamSequence["+strconv.Itoa(numCareTeamSequence)+"]", &resource.Item[numItem].CareTeamSequence[numCareTeamSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDiagnosisSequence(numItem int, numDiagnosisSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDiagnosisSequence >= len(resource.Item[numItem].DiagnosisSequence) {
		return IntInput("item["+strconv.Itoa(numItem)+"].diagnosisSequence["+strconv.Itoa(numDiagnosisSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].diagnosisSequence["+strconv.Itoa(numDiagnosisSequence)+"]", &resource.Item[numItem].DiagnosisSequence[numDiagnosisSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemProcedureSequence(numItem int, numProcedureSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numProcedureSequence >= len(resource.Item[numItem].ProcedureSequence) {
		return IntInput("item["+strconv.Itoa(numItem)+"].procedureSequence["+strconv.Itoa(numProcedureSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].procedureSequence["+strconv.Itoa(numProcedureSequence)+"]", &resource.Item[numItem].ProcedureSequence[numProcedureSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemInformationSequence(numItem int, numInformationSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numInformationSequence >= len(resource.Item[numItem].InformationSequence) {
		return IntInput("item["+strconv.Itoa(numItem)+"].informationSequence["+strconv.Itoa(numInformationSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].informationSequence["+strconv.Itoa(numInformationSequence)+"]", &resource.Item[numItem].InformationSequence[numInformationSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemRevenue(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].revenue", resource.Item[numItem].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemCategory(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].category", resource.Item[numItem].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemProductOrService(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].productOrService", resource.Item[numItem].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemProductOrServiceEnd(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].productOrServiceEnd", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].productOrServiceEnd", resource.Item[numItem].ProductOrServiceEnd, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemModifier(numItem int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numModifier >= len(resource.Item[numItem].Modifier) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemProgramCode(numItem int, numProgramCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numProgramCode >= len(resource.Item[numItem].ProgramCode) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].programCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].programCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemServicedDate(numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return DateInput("item["+strconv.Itoa(numItem)+"].servicedDate", nil, htmlAttrs)
	}
	return DateInput("item["+strconv.Itoa(numItem)+"].servicedDate", resource.Item[numItem].ServicedDate, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemLocationCodeableConcept(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].locationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].locationCodeableConcept", resource.Item[numItem].LocationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemFactor(numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return Float64Input("item["+strconv.Itoa(numItem)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("item["+strconv.Itoa(numItem)+"].factor", resource.Item[numItem].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemNoteNumber(numItem int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numNoteNumber >= len(resource.Item[numItem].NoteNumber) {
		return IntInput("item["+strconv.Itoa(numItem)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemBodySiteSubSite(numItem int, numBodySite int, numSubSite int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numBodySite >= len(resource.Item[numItem].BodySite) || numSubSite >= len(resource.Item[numItem].BodySite[numBodySite].SubSite) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].bodySite["+strconv.Itoa(numBodySite)+"].subSite["+strconv.Itoa(numSubSite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].bodySite["+strconv.Itoa(numBodySite)+"].subSite["+strconv.Itoa(numSubSite)+"]", &resource.Item[numItem].BodySite[numBodySite].SubSite[numSubSite], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemReviewOutcomeDecision(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].reviewOutcome.decision", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].reviewOutcome.decision", resource.Item[numItem].ReviewOutcome.Decision, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemReviewOutcomeReason(numItem int, numReason int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numReason >= len(resource.Item[numItem].ReviewOutcome.Reason) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].reviewOutcome.reason["+strconv.Itoa(numReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].reviewOutcome.reason["+strconv.Itoa(numReason)+"]", &resource.Item[numItem].ReviewOutcome.Reason[numReason], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemReviewOutcomePreAuthRef(numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return StringInput("item["+strconv.Itoa(numItem)+"].reviewOutcome.preAuthRef", nil, htmlAttrs)
	}
	return StringInput("item["+strconv.Itoa(numItem)+"].reviewOutcome.preAuthRef", resource.Item[numItem].ReviewOutcome.PreAuthRef, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemAdjudicationCategory(numItem int, numAdjudication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].category", &resource.Item[numItem].Adjudication[numAdjudication].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemAdjudicationReason(numItem int, numAdjudication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].reason", resource.Item[numItem].Adjudication[numAdjudication].Reason, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSequence(numItem int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].sequence", &resource.Item[numItem].Detail[numDetail].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailRevenue(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].revenue", resource.Item[numItem].Detail[numDetail].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailCategory(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].category", resource.Item[numItem].Detail[numDetail].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailProductOrService(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrService", resource.Item[numItem].Detail[numDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailProductOrServiceEnd(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrServiceEnd", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrServiceEnd", resource.Item[numItem].Detail[numDetail].ProductOrServiceEnd, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailModifier(numItem int, numDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numModifier >= len(resource.Item[numItem].Detail[numDetail].Modifier) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Detail[numDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailProgramCode(numItem int, numDetail int, numProgramCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numProgramCode >= len(resource.Item[numItem].Detail[numDetail].ProgramCode) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].programCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].programCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].Detail[numDetail].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailFactor(numItem int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return Float64Input("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].factor", resource.Item[numItem].Detail[numDetail].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailNoteNumber(numItem int, numDetail int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numNoteNumber >= len(resource.Item[numItem].Detail[numDetail].NoteNumber) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].Detail[numDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailSequence(numItem int, numDetail int, numSubDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].sequence", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailRevenue(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].revenue", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailCategory(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].category", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailProductOrService(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrService", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailProductOrServiceEnd(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrServiceEnd", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrServiceEnd", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrServiceEnd, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailModifier(numItem int, numDetail int, numSubDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numModifier >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Modifier) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailProgramCode(numItem int, numDetail int, numSubDetail int, numProgramCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numProgramCode >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProgramCode) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].programCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].programCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailFactor(numItem int, numDetail int, numSubDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return Float64Input("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].factor", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailNoteNumber(numItem int, numDetail int, numSubDetail int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numNoteNumber >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemItemSequence(numAddItem int, numItemSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numItemSequence >= len(resource.AddItem[numAddItem].ItemSequence) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].itemSequence["+strconv.Itoa(numItemSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].itemSequence["+strconv.Itoa(numItemSequence)+"]", &resource.AddItem[numAddItem].ItemSequence[numItemSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSequence(numAddItem int, numDetailSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetailSequence >= len(resource.AddItem[numAddItem].DetailSequence) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detailSequence["+strconv.Itoa(numDetailSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detailSequence["+strconv.Itoa(numDetailSequence)+"]", &resource.AddItem[numAddItem].DetailSequence[numDetailSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemSubDetailSequence(numAddItem int, numSubDetailSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numSubDetailSequence >= len(resource.AddItem[numAddItem].SubDetailSequence) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].subDetailSequence["+strconv.Itoa(numSubDetailSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].subDetailSequence["+strconv.Itoa(numSubDetailSequence)+"]", &resource.AddItem[numAddItem].SubDetailSequence[numSubDetailSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemRevenue(numAddItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].revenue", resource.AddItem[numAddItem].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemProductOrService(numAddItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].productOrService", resource.AddItem[numAddItem].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemProductOrServiceEnd(numAddItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].productOrServiceEnd", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].productOrServiceEnd", resource.AddItem[numAddItem].ProductOrServiceEnd, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemModifier(numAddItem int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numModifier >= len(resource.AddItem[numAddItem].Modifier) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemProgramCode(numAddItem int, numProgramCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numProgramCode >= len(resource.AddItem[numAddItem].ProgramCode) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].programCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].programCode["+strconv.Itoa(numProgramCode)+"]", &resource.AddItem[numAddItem].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemServicedDate(numAddItem int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return DateInput("addItem["+strconv.Itoa(numAddItem)+"].servicedDate", nil, htmlAttrs)
	}
	return DateInput("addItem["+strconv.Itoa(numAddItem)+"].servicedDate", resource.AddItem[numAddItem].ServicedDate, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemLocationCodeableConcept(numAddItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].locationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].locationCodeableConcept", resource.AddItem[numAddItem].LocationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemFactor(numAddItem int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].factor", resource.AddItem[numAddItem].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemNoteNumber(numAddItem int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numNoteNumber >= len(resource.AddItem[numAddItem].NoteNumber) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemBodySiteSubSite(numAddItem int, numBodySite int, numSubSite int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numBodySite >= len(resource.AddItem[numAddItem].BodySite) || numSubSite >= len(resource.AddItem[numAddItem].BodySite[numBodySite].SubSite) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].bodySite["+strconv.Itoa(numBodySite)+"].subSite["+strconv.Itoa(numSubSite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].bodySite["+strconv.Itoa(numBodySite)+"].subSite["+strconv.Itoa(numSubSite)+"]", &resource.AddItem[numAddItem].BodySite[numBodySite].SubSite[numSubSite], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailRevenue(numAddItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].revenue", resource.AddItem[numAddItem].Detail[numDetail].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailProductOrService(numAddItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrService", resource.AddItem[numAddItem].Detail[numDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailProductOrServiceEnd(numAddItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrServiceEnd", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrServiceEnd", resource.AddItem[numAddItem].Detail[numDetail].ProductOrServiceEnd, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailModifier(numAddItem int, numDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numModifier >= len(resource.AddItem[numAddItem].Detail[numDetail].Modifier) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Detail[numDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailFactor(numAddItem int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].factor", resource.AddItem[numAddItem].Detail[numDetail].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailNoteNumber(numAddItem int, numDetail int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numNoteNumber >= len(resource.AddItem[numAddItem].Detail[numDetail].NoteNumber) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailRevenue(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].revenue", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailProductOrService(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrService", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailProductOrServiceEnd(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrServiceEnd", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrServiceEnd", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrServiceEnd, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailModifier(numAddItem int, numDetail int, numSubDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) || numModifier >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailFactor(numAddItem int, numDetail int, numSubDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].factor", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailNoteNumber(numAddItem int, numDetail int, numSubDetail int, numNoteNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) || numNoteNumber >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_TotalCategory(numTotal int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTotal >= len(resource.Total) {
		return CodeableConceptSelect("total["+strconv.Itoa(numTotal)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("total["+strconv.Itoa(numTotal)+"].category", &resource.Total[numTotal].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PaymentType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("payment.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payment.type", resource.Payment.Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PaymentAdjustmentReason(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("payment.adjustmentReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payment.adjustmentReason", resource.Payment.AdjustmentReason, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PaymentDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("payment.date", nil, htmlAttrs)
	}
	return DateInput("payment.date", resource.Payment.Date, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcessNoteNumber(numProcessNote int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return IntInput("processNote["+strconv.Itoa(numProcessNote)+"].number", nil, htmlAttrs)
	}
	return IntInput("processNote["+strconv.Itoa(numProcessNote)+"].number", resource.ProcessNote[numProcessNote].Number, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcessNoteType(numProcessNote int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return CodeableConceptSelect("processNote["+strconv.Itoa(numProcessNote)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("processNote["+strconv.Itoa(numProcessNote)+"].type", resource.ProcessNote[numProcessNote].Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcessNoteText(numProcessNote int, htmlAttrs string) templ.Component {
	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return StringInput("processNote["+strconv.Itoa(numProcessNote)+"].text", nil, htmlAttrs)
	}
	return StringInput("processNote["+strconv.Itoa(numProcessNote)+"].text", resource.ProcessNote[numProcessNote].Text, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceCategory(numBenefitBalance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].category", &resource.BenefitBalance[numBenefitBalance].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceExcluded(numBenefitBalance int, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return BoolInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].excluded", nil, htmlAttrs)
	}
	return BoolInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].excluded", resource.BenefitBalance[numBenefitBalance].Excluded, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceName(numBenefitBalance int, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return StringInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].name", nil, htmlAttrs)
	}
	return StringInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].name", resource.BenefitBalance[numBenefitBalance].Name, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceDescription(numBenefitBalance int, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return StringInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].description", nil, htmlAttrs)
	}
	return StringInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].description", resource.BenefitBalance[numBenefitBalance].Description, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceNetwork(numBenefitBalance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].network", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].network", resource.BenefitBalance[numBenefitBalance].Network, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceUnit(numBenefitBalance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].unit", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].unit", resource.BenefitBalance[numBenefitBalance].Unit, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceTerm(numBenefitBalance int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].term", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].term", resource.BenefitBalance[numBenefitBalance].Term, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialType(numBenefitBalance int, numFinancial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) || numFinancial >= len(resource.BenefitBalance[numBenefitBalance].Financial) {
		return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].type", &resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialAllowedUnsignedInt(numBenefitBalance int, numFinancial int, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) || numFinancial >= len(resource.BenefitBalance[numBenefitBalance].Financial) {
		return IntInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].allowedUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].allowedUnsignedInt", resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].AllowedUnsignedInt, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialAllowedString(numBenefitBalance int, numFinancial int, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) || numFinancial >= len(resource.BenefitBalance[numBenefitBalance].Financial) {
		return StringInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].allowedString", nil, htmlAttrs)
	}
	return StringInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].allowedString", resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].AllowedString, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialUsedUnsignedInt(numBenefitBalance int, numFinancial int, htmlAttrs string) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) || numFinancial >= len(resource.BenefitBalance[numBenefitBalance].Financial) {
		return IntInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].usedUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].usedUnsignedInt", resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].UsedUnsignedInt, htmlAttrs)
}
