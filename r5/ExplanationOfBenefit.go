package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	TimingDate        *string          `json:"timingDate"`
	TimingPeriod      *Period          `json:"timingPeriod"`
	ValueBoolean      *bool            `json:"valueBoolean"`
	ValueString       *string          `json:"valueString"`
	ValueQuantity     *Quantity        `json:"valueQuantity"`
	ValueAttachment   *Attachment      `json:"valueAttachment"`
	ValueReference    *Reference       `json:"valueReference"`
	ValueIdentifier   *Identifier      `json:"valueIdentifier"`
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
	LocationAddress   *Address         `json:"locationAddress"`
	LocationReference *Reference       `json:"locationReference"`
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
	ServicedDate            *string                                `json:"servicedDate"`
	ServicedPeriod          *Period                                `json:"servicedPeriod"`
	LocationCodeableConcept *CodeableConcept                       `json:"locationCodeableConcept"`
	LocationAddress         *Address                               `json:"locationAddress"`
	LocationReference       *Reference                             `json:"locationReference"`
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
	ServicedDate            *string                               `json:"servicedDate"`
	ServicedPeriod          *Period                               `json:"servicedPeriod"`
	LocationCodeableConcept *CodeableConcept                      `json:"locationCodeableConcept"`
	LocationAddress         *Address                              `json:"locationAddress"`
	LocationReference       *Reference                            `json:"locationReference"`
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
	AllowedUnsignedInt *int            `json:"allowedUnsignedInt"`
	AllowedString      *string         `json:"allowedString"`
	AllowedMoney       *Money          `json:"allowedMoney"`
	UsedUnsignedInt    *int            `json:"usedUnsignedInt"`
	UsedMoney          *Money          `json:"usedMoney"`
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

func (resource *ExplanationOfBenefit) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_Status() templ.Component {
	optionsValueSet := VSExplanationofbenefit_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Type, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_SubType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("subType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subType", resource.SubType, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_Use() templ.Component {
	optionsValueSet := VSClaim_use

	if resource == nil {
		return CodeSelect("use", nil, optionsValueSet)
	}
	return CodeSelect("use", &resource.Use, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_Priority(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("priority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("priority", resource.Priority, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_FundsReserveRequested(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("fundsReserveRequested", nil, optionsValueSet)
	}
	return CodeableConceptSelect("fundsReserveRequested", resource.FundsReserveRequested, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_FundsReserve(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("fundsReserve", nil, optionsValueSet)
	}
	return CodeableConceptSelect("fundsReserve", resource.FundsReserve, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_Outcome() templ.Component {
	optionsValueSet := VSClaim_outcome

	if resource == nil {
		return CodeSelect("outcome", nil, optionsValueSet)
	}
	return CodeSelect("outcome", &resource.Outcome, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_Decision(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("decision", nil, optionsValueSet)
	}
	return CodeableConceptSelect("decision", resource.Decision, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_DiagnosisRelatedGroup(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("diagnosisRelatedGroup", nil, optionsValueSet)
	}
	return CodeableConceptSelect("diagnosisRelatedGroup", resource.DiagnosisRelatedGroup, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_FormCode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("formCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("formCode", resource.FormCode, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_RelatedRelationship(numRelated int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Related) >= numRelated {
		return CodeableConceptSelect("relationship", nil, optionsValueSet)
	}
	return CodeableConceptSelect("relationship", resource.Related[numRelated].Relationship, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_EventType(numEvent int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Event) >= numEvent {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Event[numEvent].Type, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_PayeeType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Payee.Type, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_CareTeamRole(numCareTeam int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CareTeam) >= numCareTeam {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", resource.CareTeam[numCareTeam].Role, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_CareTeamSpecialty(numCareTeam int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CareTeam) >= numCareTeam {
		return CodeableConceptSelect("specialty", nil, optionsValueSet)
	}
	return CodeableConceptSelect("specialty", resource.CareTeam[numCareTeam].Specialty, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoCategory(numSupportingInfo int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.SupportingInfo) >= numSupportingInfo {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.SupportingInfo[numSupportingInfo].Category, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoCode(numSupportingInfo int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.SupportingInfo) >= numSupportingInfo {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.SupportingInfo[numSupportingInfo].Code, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoReason(numSupportingInfo int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.SupportingInfo) >= numSupportingInfo {
		return CodingSelect("reason", nil, optionsValueSet)
	}
	return CodingSelect("reason", resource.SupportingInfo[numSupportingInfo].Reason, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_DiagnosisType(numDiagnosis int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Diagnosis) >= numDiagnosis {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Diagnosis[numDiagnosis].Type[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_DiagnosisOnAdmission(numDiagnosis int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Diagnosis) >= numDiagnosis {
		return CodeableConceptSelect("onAdmission", nil, optionsValueSet)
	}
	return CodeableConceptSelect("onAdmission", resource.Diagnosis[numDiagnosis].OnAdmission, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ProcedureType(numProcedure int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Procedure) >= numProcedure {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Procedure[numProcedure].Type[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AccidentType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Accident.Type, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemRevenue(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("revenue", resource.Item[numItem].Revenue, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemCategory(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.Item[numItem].Category, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemProductOrService(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", resource.Item[numItem].ProductOrService, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemProductOrServiceEnd(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("productOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrServiceEnd", resource.Item[numItem].ProductOrServiceEnd, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemModifier(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.Item[numItem].Modifier[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemProgramCode(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("programCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("programCode", &resource.Item[numItem].ProgramCode[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemBodySiteSubSite(numItem int, numBodySite int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].BodySite) >= numBodySite {
		return CodeableConceptSelect("subSite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subSite", &resource.Item[numItem].BodySite[numBodySite].SubSite[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemReviewOutcomeDecision(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("decision", nil, optionsValueSet)
	}
	return CodeableConceptSelect("decision", resource.Item[numItem].ReviewOutcome.Decision, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemReviewOutcomeReason(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reason", &resource.Item[numItem].ReviewOutcome.Reason[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemAdjudicationCategory(numItem int, numAdjudication int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Adjudication) >= numAdjudication {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Item[numItem].Adjudication[numAdjudication].Category, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemAdjudicationReason(numItem int, numAdjudication int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Adjudication) >= numAdjudication {
		return CodeableConceptSelect("reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reason", resource.Item[numItem].Adjudication[numAdjudication].Reason, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailRevenue(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("revenue", resource.Item[numItem].Detail[numDetail].Revenue, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailCategory(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.Item[numItem].Detail[numDetail].Category, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailProductOrService(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", resource.Item[numItem].Detail[numDetail].ProductOrService, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailProductOrServiceEnd(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("productOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrServiceEnd", resource.Item[numItem].Detail[numDetail].ProductOrServiceEnd, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailModifier(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.Item[numItem].Detail[numDetail].Modifier[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailProgramCode(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("programCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("programCode", &resource.Item[numItem].Detail[numDetail].ProgramCode[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailRevenue(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("revenue", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Revenue, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailCategory(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Category, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailProductOrService(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailProductOrServiceEnd(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("productOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrServiceEnd", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrServiceEnd, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailModifier(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailProgramCode(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("programCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("programCode", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProgramCode[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemRevenue(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("revenue", resource.AddItem[numAddItem].Revenue, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemProductOrService(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", resource.AddItem[numAddItem].ProductOrService, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemProductOrServiceEnd(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("productOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrServiceEnd", resource.AddItem[numAddItem].ProductOrServiceEnd, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemModifier(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.AddItem[numAddItem].Modifier[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemProgramCode(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("programCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("programCode", &resource.AddItem[numAddItem].ProgramCode[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemBodySiteSubSite(numAddItem int, numBodySite int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].BodySite) >= numBodySite {
		return CodeableConceptSelect("subSite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subSite", &resource.AddItem[numAddItem].BodySite[numBodySite].SubSite[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailRevenue(numAddItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return CodeableConceptSelect("revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("revenue", resource.AddItem[numAddItem].Detail[numDetail].Revenue, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailProductOrService(numAddItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", resource.AddItem[numAddItem].Detail[numDetail].ProductOrService, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailProductOrServiceEnd(numAddItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return CodeableConceptSelect("productOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrServiceEnd", resource.AddItem[numAddItem].Detail[numDetail].ProductOrServiceEnd, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailModifier(numAddItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.AddItem[numAddItem].Detail[numDetail].Modifier[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailRevenue(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("revenue", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Revenue, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailProductOrService(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailProductOrServiceEnd(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("productOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrServiceEnd", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrServiceEnd, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailModifier(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_TotalCategory(numTotal int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Total) >= numTotal {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.Total[numTotal].Category, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_PaymentType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Payment.Type, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_PaymentAdjustmentReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("adjustmentReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("adjustmentReason", resource.Payment.AdjustmentReason, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ProcessNoteType(numProcessNote int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.ProcessNote) >= numProcessNote {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.ProcessNote[numProcessNote].Type, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ProcessNoteLanguage(numProcessNote int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.ProcessNote) >= numProcessNote {
		return CodeableConceptSelect("language", nil, optionsValueSet)
	}
	return CodeableConceptSelect("language", resource.ProcessNote[numProcessNote].Language, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceCategory(numBenefitBalance int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.BenefitBalance) >= numBenefitBalance {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.BenefitBalance[numBenefitBalance].Category, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceNetwork(numBenefitBalance int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.BenefitBalance) >= numBenefitBalance {
		return CodeableConceptSelect("network", nil, optionsValueSet)
	}
	return CodeableConceptSelect("network", resource.BenefitBalance[numBenefitBalance].Network, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceUnit(numBenefitBalance int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.BenefitBalance) >= numBenefitBalance {
		return CodeableConceptSelect("unit", nil, optionsValueSet)
	}
	return CodeableConceptSelect("unit", resource.BenefitBalance[numBenefitBalance].Unit, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceTerm(numBenefitBalance int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.BenefitBalance) >= numBenefitBalance {
		return CodeableConceptSelect("term", nil, optionsValueSet)
	}
	return CodeableConceptSelect("term", resource.BenefitBalance[numBenefitBalance].Term, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialType(numBenefitBalance int, numFinancial int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.BenefitBalance[numBenefitBalance].Financial) >= numFinancial {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].Type, optionsValueSet)
}
