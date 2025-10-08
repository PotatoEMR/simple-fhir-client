package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
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
	Created               FhirDateTime                         `json:"created"`
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

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitRelated struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Claim             *Reference       `json:"claim,omitempty"`
	Relationship      *CodeableConcept `json:"relationship,omitempty"`
	Reference         *Identifier      `json:"reference,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitPayee struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Party             *Reference       `json:"party,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
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

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitSupportingInfo struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Sequence          int              `json:"sequence"`
	Category          CodeableConcept  `json:"category"`
	Code              *CodeableConcept `json:"code,omitempty"`
	TimingDate        *FhirDate        `json:"timingDate,omitempty"`
	TimingPeriod      *Period          `json:"timingPeriod,omitempty"`
	ValueBoolean      *bool            `json:"valueBoolean,omitempty"`
	ValueString       *string          `json:"valueString,omitempty"`
	ValueQuantity     *Quantity        `json:"valueQuantity,omitempty"`
	ValueAttachment   *Attachment      `json:"valueAttachment,omitempty"`
	ValueReference    *Reference       `json:"valueReference,omitempty"`
	Reason            *Coding          `json:"reason,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
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

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitProcedure struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Sequence                 int               `json:"sequence"`
	Type                     []CodeableConcept `json:"type,omitempty"`
	Date                     *FhirDateTime     `json:"date,omitempty"`
	ProcedureCodeableConcept CodeableConcept   `json:"procedureCodeableConcept"`
	ProcedureReference       Reference         `json:"procedureReference"`
	Udi                      []Reference       `json:"udi,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitInsurance struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Focal             bool        `json:"focal"`
	Coverage          Reference   `json:"coverage"`
	PreAuthRef        []string    `json:"preAuthRef,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitAccident struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Date              *FhirDate        `json:"date,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	LocationAddress   *Address         `json:"locationAddress,omitempty"`
	LocationReference *Reference       `json:"locationReference,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
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
	ServicedDate            *FhirDate                              `json:"servicedDate,omitempty"`
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

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitItemAdjudication struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Category          CodeableConcept  `json:"category"`
	Reason            *CodeableConcept `json:"reason,omitempty"`
	Amount            *Money           `json:"amount,omitempty"`
	Value             *float64         `json:"value,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
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

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
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

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
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
	ServicedDate            *FhirDate                           `json:"servicedDate,omitempty"`
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

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
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

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
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

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitTotal struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Category          CodeableConcept `json:"category"`
	Amount            Money           `json:"amount"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitPayment struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Adjustment        *Money           `json:"adjustment,omitempty"`
	AdjustmentReason  *CodeableConcept `json:"adjustmentReason,omitempty"`
	Date              *FhirDate        `json:"date,omitempty"`
	Amount            *Money           `json:"amount,omitempty"`
	Identifier        *Identifier      `json:"identifier,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitProcessNote struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Number            *int             `json:"number,omitempty"`
	Type              *string          `json:"type,omitempty"`
	Text              *string          `json:"text,omitempty"`
	Language          *CodeableConcept `json:"language,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
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

// http://hl7.org/fhir/r4/StructureDefinition/ExplanationOfBenefit
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
func (resource *ExplanationOfBenefit) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSExplanationofbenefit_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SubType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("subType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subType", resource.SubType, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Use(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSClaim_use

	if resource == nil {
		return CodeSelect("use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("use", &resource.Use, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Patient(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "patient", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "patient", &resource.Patient, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BillablePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("billablePeriod", nil, htmlAttrs)
	}
	return PeriodInput("billablePeriod", resource.BillablePeriod, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Created(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("created", nil, htmlAttrs)
	}
	return FhirDateTimeInput("created", &resource.Created, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Enterer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "enterer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "enterer", resource.Enterer, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Insurer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "insurer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "insurer", &resource.Insurer, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Provider(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "provider", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "provider", &resource.Provider, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Priority(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_FundsReserveRequested(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("fundsReserveRequested", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("fundsReserveRequested", resource.FundsReserveRequested, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_FundsReserve(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("fundsReserve", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("fundsReserve", resource.FundsReserve, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Prescription(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "prescription", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "prescription", resource.Prescription, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_OriginalPrescription(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "originalPrescription", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "originalPrescription", resource.OriginalPrescription, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Referral(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "referral", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "referral", resource.Referral, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Facility(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "facility", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "facility", resource.Facility, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Claim(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "claim", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "claim", resource.Claim, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ClaimResponse(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "claimResponse", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "claimResponse", resource.ClaimResponse, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Outcome(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSRemittance_outcome

	if resource == nil {
		return CodeSelect("outcome", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("outcome", &resource.Outcome, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Disposition(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("disposition", nil, htmlAttrs)
	}
	return StringInput("disposition", resource.Disposition, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PreAuthRef(numPreAuthRef int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPreAuthRef >= len(resource.PreAuthRef) {
		return StringInput("preAuthRef["+strconv.Itoa(numPreAuthRef)+"]", nil, htmlAttrs)
	}
	return StringInput("preAuthRef["+strconv.Itoa(numPreAuthRef)+"]", &resource.PreAuthRef[numPreAuthRef], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PreAuthRefPeriod(numPreAuthRefPeriod int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPreAuthRefPeriod >= len(resource.PreAuthRefPeriod) {
		return PeriodInput("preAuthRefPeriod["+strconv.Itoa(numPreAuthRefPeriod)+"]", nil, htmlAttrs)
	}
	return PeriodInput("preAuthRefPeriod["+strconv.Itoa(numPreAuthRefPeriod)+"]", &resource.PreAuthRefPeriod[numPreAuthRefPeriod], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Precedence(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("precedence", nil, htmlAttrs)
	}
	return IntInput("precedence", resource.Precedence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_FormCode(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("formCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("formCode", resource.FormCode, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_Form(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return AttachmentInput("form", nil, htmlAttrs)
	}
	return AttachmentInput("form", resource.Form, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("benefitPeriod", nil, htmlAttrs)
	}
	return PeriodInput("benefitPeriod", resource.BenefitPeriod, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_RelatedClaim(frs []FhirResource, numRelated int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelated >= len(resource.Related) {
		return ReferenceInput(frs, "related["+strconv.Itoa(numRelated)+"].claim", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "related["+strconv.Itoa(numRelated)+"].claim", resource.Related[numRelated].Claim, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_RelatedRelationship(numRelated int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelated >= len(resource.Related) {
		return CodeableConceptSelect("related["+strconv.Itoa(numRelated)+"].relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("related["+strconv.Itoa(numRelated)+"].relationship", resource.Related[numRelated].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_RelatedReference(numRelated int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelated >= len(resource.Related) {
		return IdentifierInput("related["+strconv.Itoa(numRelated)+"].reference", nil, htmlAttrs)
	}
	return IdentifierInput("related["+strconv.Itoa(numRelated)+"].reference", resource.Related[numRelated].Reference, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PayeeType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("payee.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payee.type", resource.Payee.Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PayeeParty(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "payee.party", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "payee.party", resource.Payee.Party, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_CareTeamSequence(numCareTeam int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return IntInput("careTeam["+strconv.Itoa(numCareTeam)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("careTeam["+strconv.Itoa(numCareTeam)+"].sequence", &resource.CareTeam[numCareTeam].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_CareTeamProvider(frs []FhirResource, numCareTeam int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return ReferenceInput(frs, "careTeam["+strconv.Itoa(numCareTeam)+"].provider", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "careTeam["+strconv.Itoa(numCareTeam)+"].provider", &resource.CareTeam[numCareTeam].Provider, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_CareTeamResponsible(numCareTeam int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return BoolInput("careTeam["+strconv.Itoa(numCareTeam)+"].responsible", nil, htmlAttrs)
	}
	return BoolInput("careTeam["+strconv.Itoa(numCareTeam)+"].responsible", resource.CareTeam[numCareTeam].Responsible, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_CareTeamRole(numCareTeam int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return CodeableConceptSelect("careTeam["+strconv.Itoa(numCareTeam)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("careTeam["+strconv.Itoa(numCareTeam)+"].role", resource.CareTeam[numCareTeam].Role, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_CareTeamQualification(numCareTeam int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return CodeableConceptSelect("careTeam["+strconv.Itoa(numCareTeam)+"].qualification", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("careTeam["+strconv.Itoa(numCareTeam)+"].qualification", resource.CareTeam[numCareTeam].Qualification, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoSequence(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return IntInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].sequence", &resource.SupportingInfo[numSupportingInfo].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoCategory(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodeableConceptSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].category", &resource.SupportingInfo[numSupportingInfo].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoCode(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodeableConceptSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].code", resource.SupportingInfo[numSupportingInfo].Code, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoTimingDate(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return FhirDateInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].timingDate", nil, htmlAttrs)
	}
	return FhirDateInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].timingDate", resource.SupportingInfo[numSupportingInfo].TimingDate, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoTimingPeriod(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return PeriodInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].timingPeriod", nil, htmlAttrs)
	}
	return PeriodInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].timingPeriod", resource.SupportingInfo[numSupportingInfo].TimingPeriod, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoValueBoolean(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return BoolInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueBoolean", resource.SupportingInfo[numSupportingInfo].ValueBoolean, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoValueString(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return StringInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueString", resource.SupportingInfo[numSupportingInfo].ValueString, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoValueQuantity(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return QuantityInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueQuantity", resource.SupportingInfo[numSupportingInfo].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoValueAttachment(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return AttachmentInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueAttachment", resource.SupportingInfo[numSupportingInfo].ValueAttachment, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoValueReference(frs []FhirResource, numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return ReferenceInput(frs, "supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueReference", resource.SupportingInfo[numSupportingInfo].ValueReference, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_SupportingInfoReason(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodingSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].reason", resource.SupportingInfo[numSupportingInfo].Reason, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_DiagnosisSequence(numDiagnosis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return IntInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].sequence", &resource.Diagnosis[numDiagnosis].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_DiagnosisDiagnosisCodeableConcept(numDiagnosis int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].diagnosisCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].diagnosisCodeableConcept", &resource.Diagnosis[numDiagnosis].DiagnosisCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_DiagnosisDiagnosisReference(frs []FhirResource, numDiagnosis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return ReferenceInput(frs, "diagnosis["+strconv.Itoa(numDiagnosis)+"].diagnosisReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "diagnosis["+strconv.Itoa(numDiagnosis)+"].diagnosisReference", &resource.Diagnosis[numDiagnosis].DiagnosisReference, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_DiagnosisType(numDiagnosis int, numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) || numType >= len(resource.Diagnosis[numDiagnosis].Type) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].type["+strconv.Itoa(numType)+"]", &resource.Diagnosis[numDiagnosis].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_DiagnosisOnAdmission(numDiagnosis int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].onAdmission", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].onAdmission", resource.Diagnosis[numDiagnosis].OnAdmission, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_DiagnosisPackageCode(numDiagnosis int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].packageCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].packageCode", resource.Diagnosis[numDiagnosis].PackageCode, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcedureSequence(numProcedure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return IntInput("procedure["+strconv.Itoa(numProcedure)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("procedure["+strconv.Itoa(numProcedure)+"].sequence", &resource.Procedure[numProcedure].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcedureType(numProcedure int, numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) || numType >= len(resource.Procedure[numProcedure].Type) {
		return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].type["+strconv.Itoa(numType)+"]", &resource.Procedure[numProcedure].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcedureDate(numProcedure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return FhirDateTimeInput("procedure["+strconv.Itoa(numProcedure)+"].date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("procedure["+strconv.Itoa(numProcedure)+"].date", resource.Procedure[numProcedure].Date, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcedureProcedureCodeableConcept(numProcedure int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].procedureCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].procedureCodeableConcept", &resource.Procedure[numProcedure].ProcedureCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcedureProcedureReference(frs []FhirResource, numProcedure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return ReferenceInput(frs, "procedure["+strconv.Itoa(numProcedure)+"].procedureReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "procedure["+strconv.Itoa(numProcedure)+"].procedureReference", &resource.Procedure[numProcedure].ProcedureReference, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcedureUdi(frs []FhirResource, numProcedure int, numUdi int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) || numUdi >= len(resource.Procedure[numProcedure].Udi) {
		return ReferenceInput(frs, "procedure["+strconv.Itoa(numProcedure)+"].udi["+strconv.Itoa(numUdi)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "procedure["+strconv.Itoa(numProcedure)+"].udi["+strconv.Itoa(numUdi)+"]", &resource.Procedure[numProcedure].Udi[numUdi], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_InsuranceFocal(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].focal", nil, htmlAttrs)
	}
	return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].focal", &resource.Insurance[numInsurance].Focal, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_InsuranceCoverage(frs []FhirResource, numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return ReferenceInput(frs, "insurance["+strconv.Itoa(numInsurance)+"].coverage", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "insurance["+strconv.Itoa(numInsurance)+"].coverage", &resource.Insurance[numInsurance].Coverage, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_InsurancePreAuthRef(numInsurance int, numPreAuthRef int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numPreAuthRef >= len(resource.Insurance[numInsurance].PreAuthRef) {
		return StringInput("insurance["+strconv.Itoa(numInsurance)+"].preAuthRef["+strconv.Itoa(numPreAuthRef)+"]", nil, htmlAttrs)
	}
	return StringInput("insurance["+strconv.Itoa(numInsurance)+"].preAuthRef["+strconv.Itoa(numPreAuthRef)+"]", &resource.Insurance[numInsurance].PreAuthRef[numPreAuthRef], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AccidentDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("accident.date", nil, htmlAttrs)
	}
	return FhirDateInput("accident.date", resource.Accident.Date, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AccidentType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("accident.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("accident.type", resource.Accident.Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AccidentLocationAddress(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return AddressInput("accident.locationAddress", nil, htmlAttrs)
	}
	return AddressInput("accident.locationAddress", resource.Accident.LocationAddress, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AccidentLocationReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "accident.locationReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "accident.locationReference", resource.Accident.LocationReference, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemSequence(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return IntInput("item["+strconv.Itoa(numItem)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].sequence", &resource.Item[numItem].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemCareTeamSequence(numItem int, numCareTeamSequence int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numCareTeamSequence >= len(resource.Item[numItem].CareTeamSequence) {
		return IntInput("item["+strconv.Itoa(numItem)+"].careTeamSequence["+strconv.Itoa(numCareTeamSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].careTeamSequence["+strconv.Itoa(numCareTeamSequence)+"]", &resource.Item[numItem].CareTeamSequence[numCareTeamSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDiagnosisSequence(numItem int, numDiagnosisSequence int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDiagnosisSequence >= len(resource.Item[numItem].DiagnosisSequence) {
		return IntInput("item["+strconv.Itoa(numItem)+"].diagnosisSequence["+strconv.Itoa(numDiagnosisSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].diagnosisSequence["+strconv.Itoa(numDiagnosisSequence)+"]", &resource.Item[numItem].DiagnosisSequence[numDiagnosisSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemProcedureSequence(numItem int, numProcedureSequence int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numProcedureSequence >= len(resource.Item[numItem].ProcedureSequence) {
		return IntInput("item["+strconv.Itoa(numItem)+"].procedureSequence["+strconv.Itoa(numProcedureSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].procedureSequence["+strconv.Itoa(numProcedureSequence)+"]", &resource.Item[numItem].ProcedureSequence[numProcedureSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemInformationSequence(numItem int, numInformationSequence int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numInformationSequence >= len(resource.Item[numItem].InformationSequence) {
		return IntInput("item["+strconv.Itoa(numItem)+"].informationSequence["+strconv.Itoa(numInformationSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].informationSequence["+strconv.Itoa(numInformationSequence)+"]", &resource.Item[numItem].InformationSequence[numInformationSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemRevenue(numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].revenue", resource.Item[numItem].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemCategory(numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].category", resource.Item[numItem].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemProductOrService(numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].productOrService", &resource.Item[numItem].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemModifier(numItem int, numModifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numModifier >= len(resource.Item[numItem].Modifier) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemProgramCode(numItem int, numProgramCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numProgramCode >= len(resource.Item[numItem].ProgramCode) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].programCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].programCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemServicedDate(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return FhirDateInput("item["+strconv.Itoa(numItem)+"].servicedDate", nil, htmlAttrs)
	}
	return FhirDateInput("item["+strconv.Itoa(numItem)+"].servicedDate", resource.Item[numItem].ServicedDate, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemServicedPeriod(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return PeriodInput("item["+strconv.Itoa(numItem)+"].servicedPeriod", nil, htmlAttrs)
	}
	return PeriodInput("item["+strconv.Itoa(numItem)+"].servicedPeriod", resource.Item[numItem].ServicedPeriod, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemLocationCodeableConcept(numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].locationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].locationCodeableConcept", resource.Item[numItem].LocationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemLocationAddress(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return AddressInput("item["+strconv.Itoa(numItem)+"].locationAddress", nil, htmlAttrs)
	}
	return AddressInput("item["+strconv.Itoa(numItem)+"].locationAddress", resource.Item[numItem].LocationAddress, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemLocationReference(frs []FhirResource, numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return ReferenceInput(frs, "item["+strconv.Itoa(numItem)+"].locationReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "item["+strconv.Itoa(numItem)+"].locationReference", resource.Item[numItem].LocationReference, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemQuantity(numItem int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return QuantityInput("item["+strconv.Itoa(numItem)+"].quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("item["+strconv.Itoa(numItem)+"].quantity", resource.Item[numItem].Quantity, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemUnitPrice(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].unitPrice", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].unitPrice", resource.Item[numItem].UnitPrice, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemFactor(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return Float64Input("item["+strconv.Itoa(numItem)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("item["+strconv.Itoa(numItem)+"].factor", resource.Item[numItem].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemNet(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].net", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].net", resource.Item[numItem].Net, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemUdi(frs []FhirResource, numItem int, numUdi int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numUdi >= len(resource.Item[numItem].Udi) {
		return ReferenceInput(frs, "item["+strconv.Itoa(numItem)+"].udi["+strconv.Itoa(numUdi)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "item["+strconv.Itoa(numItem)+"].udi["+strconv.Itoa(numUdi)+"]", &resource.Item[numItem].Udi[numUdi], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemBodySite(numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].bodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].bodySite", resource.Item[numItem].BodySite, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemSubSite(numItem int, numSubSite int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numSubSite >= len(resource.Item[numItem].SubSite) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].subSite["+strconv.Itoa(numSubSite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].subSite["+strconv.Itoa(numSubSite)+"]", &resource.Item[numItem].SubSite[numSubSite], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemEncounter(frs []FhirResource, numItem int, numEncounter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numEncounter >= len(resource.Item[numItem].Encounter) {
		return ReferenceInput(frs, "item["+strconv.Itoa(numItem)+"].encounter["+strconv.Itoa(numEncounter)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "item["+strconv.Itoa(numItem)+"].encounter["+strconv.Itoa(numEncounter)+"]", &resource.Item[numItem].Encounter[numEncounter], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemNoteNumber(numItem int, numNoteNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numNoteNumber >= len(resource.Item[numItem].NoteNumber) {
		return IntInput("item["+strconv.Itoa(numItem)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemAdjudicationCategory(numItem int, numAdjudication int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].category", &resource.Item[numItem].Adjudication[numAdjudication].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemAdjudicationReason(numItem int, numAdjudication int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].reason", resource.Item[numItem].Adjudication[numAdjudication].Reason, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemAdjudicationAmount(numItem int, numAdjudication int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].amount", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].amount", resource.Item[numItem].Adjudication[numAdjudication].Amount, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemAdjudicationValue(numItem int, numAdjudication int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numAdjudication >= len(resource.Item[numItem].Adjudication) {
		return Float64Input("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].value", nil, htmlAttrs)
	}
	return Float64Input("item["+strconv.Itoa(numItem)+"].adjudication["+strconv.Itoa(numAdjudication)+"].value", resource.Item[numItem].Adjudication[numAdjudication].Value, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSequence(numItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].sequence", &resource.Item[numItem].Detail[numDetail].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailRevenue(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].revenue", resource.Item[numItem].Detail[numDetail].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailCategory(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].category", resource.Item[numItem].Detail[numDetail].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailProductOrService(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrService", &resource.Item[numItem].Detail[numDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailModifier(numItem int, numDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numModifier >= len(resource.Item[numItem].Detail[numDetail].Modifier) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Detail[numDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailProgramCode(numItem int, numDetail int, numProgramCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numProgramCode >= len(resource.Item[numItem].Detail[numDetail].ProgramCode) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].programCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].programCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].Detail[numDetail].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailQuantity(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return QuantityInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].quantity", resource.Item[numItem].Detail[numDetail].Quantity, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailUnitPrice(numItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].unitPrice", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].unitPrice", resource.Item[numItem].Detail[numDetail].UnitPrice, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailFactor(numItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return Float64Input("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].factor", resource.Item[numItem].Detail[numDetail].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailNet(numItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].net", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].net", resource.Item[numItem].Detail[numDetail].Net, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailUdi(frs []FhirResource, numItem int, numDetail int, numUdi int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numUdi >= len(resource.Item[numItem].Detail[numDetail].Udi) {
		return ReferenceInput(frs, "item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].udi["+strconv.Itoa(numUdi)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].udi["+strconv.Itoa(numUdi)+"]", &resource.Item[numItem].Detail[numDetail].Udi[numUdi], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailNoteNumber(numItem int, numDetail int, numNoteNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numNoteNumber >= len(resource.Item[numItem].Detail[numDetail].NoteNumber) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].Detail[numDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailSequence(numItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].sequence", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Sequence, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailRevenue(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].revenue", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailCategory(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].category", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailProductOrService(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrService", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailModifier(numItem int, numDetail int, numSubDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numModifier >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Modifier) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailProgramCode(numItem int, numDetail int, numSubDetail int, numProgramCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numProgramCode >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProgramCode) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].programCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].programCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailQuantity(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return QuantityInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].quantity", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Quantity, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailUnitPrice(numItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].unitPrice", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].unitPrice", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].UnitPrice, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailFactor(numItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return Float64Input("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].factor", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailNet(numItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].net", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].net", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Net, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailUdi(frs []FhirResource, numItem int, numDetail int, numSubDetail int, numUdi int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numUdi >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Udi) {
		return ReferenceInput(frs, "item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].udi["+strconv.Itoa(numUdi)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].udi["+strconv.Itoa(numUdi)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Udi[numUdi], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ItemDetailSubDetailNoteNumber(numItem int, numDetail int, numSubDetail int, numNoteNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numNoteNumber >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemItemSequence(numAddItem int, numItemSequence int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numItemSequence >= len(resource.AddItem[numAddItem].ItemSequence) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].itemSequence["+strconv.Itoa(numItemSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].itemSequence["+strconv.Itoa(numItemSequence)+"]", &resource.AddItem[numAddItem].ItemSequence[numItemSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSequence(numAddItem int, numDetailSequence int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetailSequence >= len(resource.AddItem[numAddItem].DetailSequence) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detailSequence["+strconv.Itoa(numDetailSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detailSequence["+strconv.Itoa(numDetailSequence)+"]", &resource.AddItem[numAddItem].DetailSequence[numDetailSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemSubDetailSequence(numAddItem int, numSubDetailSequence int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numSubDetailSequence >= len(resource.AddItem[numAddItem].SubDetailSequence) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].subDetailSequence["+strconv.Itoa(numSubDetailSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].subDetailSequence["+strconv.Itoa(numSubDetailSequence)+"]", &resource.AddItem[numAddItem].SubDetailSequence[numSubDetailSequence], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemProvider(frs []FhirResource, numAddItem int, numProvider int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numProvider >= len(resource.AddItem[numAddItem].Provider) {
		return ReferenceInput(frs, "addItem["+strconv.Itoa(numAddItem)+"].provider["+strconv.Itoa(numProvider)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "addItem["+strconv.Itoa(numAddItem)+"].provider["+strconv.Itoa(numProvider)+"]", &resource.AddItem[numAddItem].Provider[numProvider], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemProductOrService(numAddItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].productOrService", &resource.AddItem[numAddItem].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemModifier(numAddItem int, numModifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numModifier >= len(resource.AddItem[numAddItem].Modifier) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemProgramCode(numAddItem int, numProgramCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numProgramCode >= len(resource.AddItem[numAddItem].ProgramCode) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].programCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].programCode["+strconv.Itoa(numProgramCode)+"]", &resource.AddItem[numAddItem].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemServicedDate(numAddItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return FhirDateInput("addItem["+strconv.Itoa(numAddItem)+"].servicedDate", nil, htmlAttrs)
	}
	return FhirDateInput("addItem["+strconv.Itoa(numAddItem)+"].servicedDate", resource.AddItem[numAddItem].ServicedDate, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemServicedPeriod(numAddItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return PeriodInput("addItem["+strconv.Itoa(numAddItem)+"].servicedPeriod", nil, htmlAttrs)
	}
	return PeriodInput("addItem["+strconv.Itoa(numAddItem)+"].servicedPeriod", resource.AddItem[numAddItem].ServicedPeriod, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemLocationCodeableConcept(numAddItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].locationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].locationCodeableConcept", resource.AddItem[numAddItem].LocationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemLocationAddress(numAddItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return AddressInput("addItem["+strconv.Itoa(numAddItem)+"].locationAddress", nil, htmlAttrs)
	}
	return AddressInput("addItem["+strconv.Itoa(numAddItem)+"].locationAddress", resource.AddItem[numAddItem].LocationAddress, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemLocationReference(frs []FhirResource, numAddItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return ReferenceInput(frs, "addItem["+strconv.Itoa(numAddItem)+"].locationReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "addItem["+strconv.Itoa(numAddItem)+"].locationReference", resource.AddItem[numAddItem].LocationReference, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemQuantity(numAddItem int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return QuantityInput("addItem["+strconv.Itoa(numAddItem)+"].quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("addItem["+strconv.Itoa(numAddItem)+"].quantity", resource.AddItem[numAddItem].Quantity, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemUnitPrice(numAddItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].unitPrice", nil, htmlAttrs)
	}
	return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].unitPrice", resource.AddItem[numAddItem].UnitPrice, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemFactor(numAddItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].factor", resource.AddItem[numAddItem].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemNet(numAddItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].net", nil, htmlAttrs)
	}
	return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].net", resource.AddItem[numAddItem].Net, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemBodySite(numAddItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].bodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].bodySite", resource.AddItem[numAddItem].BodySite, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemSubSite(numAddItem int, numSubSite int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numSubSite >= len(resource.AddItem[numAddItem].SubSite) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].subSite["+strconv.Itoa(numSubSite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].subSite["+strconv.Itoa(numSubSite)+"]", &resource.AddItem[numAddItem].SubSite[numSubSite], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemNoteNumber(numAddItem int, numNoteNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numNoteNumber >= len(resource.AddItem[numAddItem].NoteNumber) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailProductOrService(numAddItem int, numDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrService", &resource.AddItem[numAddItem].Detail[numDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailModifier(numAddItem int, numDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numModifier >= len(resource.AddItem[numAddItem].Detail[numDetail].Modifier) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Detail[numDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailQuantity(numAddItem int, numDetail int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return QuantityInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].quantity", resource.AddItem[numAddItem].Detail[numDetail].Quantity, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailUnitPrice(numAddItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].unitPrice", nil, htmlAttrs)
	}
	return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].unitPrice", resource.AddItem[numAddItem].Detail[numDetail].UnitPrice, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailFactor(numAddItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].factor", resource.AddItem[numAddItem].Detail[numDetail].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailNet(numAddItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) {
		return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].net", nil, htmlAttrs)
	}
	return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].net", resource.AddItem[numAddItem].Detail[numDetail].Net, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailNoteNumber(numAddItem int, numDetail int, numNoteNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numNoteNumber >= len(resource.AddItem[numAddItem].Detail[numDetail].NoteNumber) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailProductOrService(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrService", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailModifier(numAddItem int, numDetail int, numSubDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) || numModifier >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier) {
		return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailQuantity(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return QuantityInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].quantity", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Quantity, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailUnitPrice(numAddItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].unitPrice", nil, htmlAttrs)
	}
	return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].unitPrice", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].UnitPrice, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailFactor(numAddItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].factor", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Factor, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailNet(numAddItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) {
		return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].net", nil, htmlAttrs)
	}
	return MoneyInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].net", resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].Net, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailNoteNumber(numAddItem int, numDetail int, numSubDetail int, numNoteNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numAddItem >= len(resource.AddItem) || numDetail >= len(resource.AddItem[numAddItem].Detail) || numSubDetail >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) || numNoteNumber >= len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber) {
		return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("addItem["+strconv.Itoa(numAddItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].noteNumber["+strconv.Itoa(numNoteNumber)+"]", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].NoteNumber[numNoteNumber], htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_TotalCategory(numTotal int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTotal >= len(resource.Total) {
		return CodeableConceptSelect("total["+strconv.Itoa(numTotal)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("total["+strconv.Itoa(numTotal)+"].category", &resource.Total[numTotal].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_TotalAmount(numTotal int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTotal >= len(resource.Total) {
		return MoneyInput("total["+strconv.Itoa(numTotal)+"].amount", nil, htmlAttrs)
	}
	return MoneyInput("total["+strconv.Itoa(numTotal)+"].amount", &resource.Total[numTotal].Amount, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PaymentType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("payment.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payment.type", resource.Payment.Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PaymentAdjustment(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return MoneyInput("payment.adjustment", nil, htmlAttrs)
	}
	return MoneyInput("payment.adjustment", resource.Payment.Adjustment, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PaymentAdjustmentReason(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("payment.adjustmentReason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payment.adjustmentReason", resource.Payment.AdjustmentReason, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PaymentDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("payment.date", nil, htmlAttrs)
	}
	return FhirDateInput("payment.date", resource.Payment.Date, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_PaymentAmount(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return MoneyInput("payment.amount", nil, htmlAttrs)
	}
	return MoneyInput("payment.amount", resource.Payment.Amount, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcessNoteNumber(numProcessNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return IntInput("processNote["+strconv.Itoa(numProcessNote)+"].number", nil, htmlAttrs)
	}
	return IntInput("processNote["+strconv.Itoa(numProcessNote)+"].number", resource.ProcessNote[numProcessNote].Number, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_ProcessNoteType(numProcessNote int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSNote_type

	if resource == nil || numProcessNote >= len(resource.ProcessNote) {
		return CodeSelect("processNote["+strconv.Itoa(numProcessNote)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("processNote["+strconv.Itoa(numProcessNote)+"].type", resource.ProcessNote[numProcessNote].Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceCategory(numBenefitBalance int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].category", &resource.BenefitBalance[numBenefitBalance].Category, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceExcluded(numBenefitBalance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return BoolInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].excluded", nil, htmlAttrs)
	}
	return BoolInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].excluded", resource.BenefitBalance[numBenefitBalance].Excluded, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceName(numBenefitBalance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return StringInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].name", nil, htmlAttrs)
	}
	return StringInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].name", resource.BenefitBalance[numBenefitBalance].Name, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceDescription(numBenefitBalance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return StringInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].description", nil, htmlAttrs)
	}
	return StringInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].description", resource.BenefitBalance[numBenefitBalance].Description, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceNetwork(numBenefitBalance int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].network", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].network", resource.BenefitBalance[numBenefitBalance].Network, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceUnit(numBenefitBalance int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].unit", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].unit", resource.BenefitBalance[numBenefitBalance].Unit, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceTerm(numBenefitBalance int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) {
		return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].term", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].term", resource.BenefitBalance[numBenefitBalance].Term, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialType(numBenefitBalance int, numFinancial int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) || numFinancial >= len(resource.BenefitBalance[numBenefitBalance].Financial) {
		return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].type", &resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].Type, optionsValueSet, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialAllowedUnsignedInt(numBenefitBalance int, numFinancial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) || numFinancial >= len(resource.BenefitBalance[numBenefitBalance].Financial) {
		return IntInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].allowedUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].allowedUnsignedInt", resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].AllowedUnsignedInt, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialAllowedString(numBenefitBalance int, numFinancial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) || numFinancial >= len(resource.BenefitBalance[numBenefitBalance].Financial) {
		return StringInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].allowedString", nil, htmlAttrs)
	}
	return StringInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].allowedString", resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].AllowedString, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialAllowedMoney(numBenefitBalance int, numFinancial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) || numFinancial >= len(resource.BenefitBalance[numBenefitBalance].Financial) {
		return MoneyInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].allowedMoney", nil, htmlAttrs)
	}
	return MoneyInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].allowedMoney", resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].AllowedMoney, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialUsedUnsignedInt(numBenefitBalance int, numFinancial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) || numFinancial >= len(resource.BenefitBalance[numBenefitBalance].Financial) {
		return IntInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].usedUnsignedInt", nil, htmlAttrs)
	}
	return IntInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].usedUnsignedInt", resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].UsedUnsignedInt, htmlAttrs)
}
func (resource *ExplanationOfBenefit) T_BenefitBalanceFinancialUsedMoney(numBenefitBalance int, numFinancial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numBenefitBalance >= len(resource.BenefitBalance) || numFinancial >= len(resource.BenefitBalance[numBenefitBalance].Financial) {
		return MoneyInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].usedMoney", nil, htmlAttrs)
	}
	return MoneyInput("benefitBalance["+strconv.Itoa(numBenefitBalance)+"].financial["+strconv.Itoa(numFinancial)+"].usedMoney", resource.BenefitBalance[numBenefitBalance].Financial[numFinancial].UsedMoney, htmlAttrs)
}
