package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	Created               string                               `json:"created"`
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
	TimingDate        *string          `json:"timingDate"`
	TimingPeriod      *Period          `json:"timingPeriod"`
	ValueBoolean      *bool            `json:"valueBoolean"`
	ValueString       *string          `json:"valueString"`
	ValueQuantity     *Quantity        `json:"valueQuantity"`
	ValueAttachment   *Attachment      `json:"valueAttachment"`
	ValueReference    *Reference       `json:"valueReference"`
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
	Date                     *string           `json:"date,omitempty"`
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
	Date              *string          `json:"date,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	LocationAddress   *Address         `json:"locationAddress"`
	LocationReference *Reference       `json:"locationReference"`
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
	ServicedDate            *string                                `json:"servicedDate"`
	ServicedPeriod          *Period                                `json:"servicedPeriod"`
	LocationCodeableConcept *CodeableConcept                       `json:"locationCodeableConcept"`
	LocationAddress         *Address                               `json:"locationAddress"`
	LocationReference       *Reference                             `json:"locationReference"`
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
	ServicedDate            *string                             `json:"servicedDate"`
	ServicedPeriod          *Period                             `json:"servicedPeriod"`
	LocationCodeableConcept *CodeableConcept                    `json:"locationCodeableConcept"`
	LocationAddress         *Address                            `json:"locationAddress"`
	LocationReference       *Reference                          `json:"locationReference"`
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
	Date              *string          `json:"date,omitempty"`
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
	optionsValueSet := VSRemittance_outcome

	if resource == nil {
		return CodeSelect("outcome", nil, optionsValueSet)
	}
	return CodeSelect("outcome", &resource.Outcome, optionsValueSet)
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
func (resource *ExplanationOfBenefit) T_CareTeamQualification(numCareTeam int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CareTeam) >= numCareTeam {
		return CodeableConceptSelect("qualification", nil, optionsValueSet)
	}
	return CodeableConceptSelect("qualification", resource.CareTeam[numCareTeam].Qualification, optionsValueSet)
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
func (resource *ExplanationOfBenefit) T_DiagnosisPackageCode(numDiagnosis int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Diagnosis) >= numDiagnosis {
		return CodeableConceptSelect("packageCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("packageCode", resource.Diagnosis[numDiagnosis].PackageCode, optionsValueSet)
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
	return CodeableConceptSelect("productOrService", &resource.Item[numItem].ProductOrService, optionsValueSet)
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
func (resource *ExplanationOfBenefit) T_ItemBodySite(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("bodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("bodySite", resource.Item[numItem].BodySite, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_ItemSubSite(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("subSite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subSite", &resource.Item[numItem].SubSite[0], optionsValueSet)
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
	return CodeableConceptSelect("productOrService", &resource.Item[numItem].Detail[numDetail].ProductOrService, optionsValueSet)
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
	return CodeableConceptSelect("productOrService", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet)
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
func (resource *ExplanationOfBenefit) T_AddItemProductOrService(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", &resource.AddItem[numAddItem].ProductOrService, optionsValueSet)
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
func (resource *ExplanationOfBenefit) T_AddItemBodySite(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("bodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("bodySite", resource.AddItem[numAddItem].BodySite, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemSubSite(numAddItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem) >= numAddItem {
		return CodeableConceptSelect("subSite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subSite", &resource.AddItem[numAddItem].SubSite[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailProductOrService(numAddItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", &resource.AddItem[numAddItem].Detail[numDetail].ProductOrService, optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailModifier(numAddItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail) >= numDetail {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.AddItem[numAddItem].Detail[numDetail].Modifier[0], optionsValueSet)
}
func (resource *ExplanationOfBenefit) T_AddItemDetailSubDetailProductOrService(numAddItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.AddItem[numAddItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", &resource.AddItem[numAddItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet)
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
func (resource *ExplanationOfBenefit) T_ProcessNoteType(numProcessNote int) templ.Component {
	optionsValueSet := VSNote_type

	if resource == nil && len(resource.ProcessNote) >= numProcessNote {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", resource.ProcessNote[numProcessNote].Type, optionsValueSet)
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
