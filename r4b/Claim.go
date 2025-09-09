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

// http://hl7.org/fhir/r4b/StructureDefinition/Claim
type Claim struct {
	Id                   *string               `json:"id,omitempty"`
	Meta                 *Meta                 `json:"meta,omitempty"`
	ImplicitRules        *string               `json:"implicitRules,omitempty"`
	Language             *string               `json:"language,omitempty"`
	Text                 *Narrative            `json:"text,omitempty"`
	Contained            []Resource            `json:"contained,omitempty"`
	Extension            []Extension           `json:"extension,omitempty"`
	ModifierExtension    []Extension           `json:"modifierExtension,omitempty"`
	Identifier           []Identifier          `json:"identifier,omitempty"`
	Status               string                `json:"status"`
	Type                 CodeableConcept       `json:"type"`
	SubType              *CodeableConcept      `json:"subType,omitempty"`
	Use                  string                `json:"use"`
	Patient              Reference             `json:"patient"`
	BillablePeriod       *Period               `json:"billablePeriod,omitempty"`
	Created              time.Time             `json:"created,format:'2006-01-02T15:04:05Z07:00'"`
	Enterer              *Reference            `json:"enterer,omitempty"`
	Insurer              *Reference            `json:"insurer,omitempty"`
	Provider             Reference             `json:"provider"`
	Priority             CodeableConcept       `json:"priority"`
	FundsReserve         *CodeableConcept      `json:"fundsReserve,omitempty"`
	Related              []ClaimRelated        `json:"related,omitempty"`
	Prescription         *Reference            `json:"prescription,omitempty"`
	OriginalPrescription *Reference            `json:"originalPrescription,omitempty"`
	Payee                *ClaimPayee           `json:"payee,omitempty"`
	Referral             *Reference            `json:"referral,omitempty"`
	Facility             *Reference            `json:"facility,omitempty"`
	CareTeam             []ClaimCareTeam       `json:"careTeam,omitempty"`
	SupportingInfo       []ClaimSupportingInfo `json:"supportingInfo,omitempty"`
	Diagnosis            []ClaimDiagnosis      `json:"diagnosis,omitempty"`
	Procedure            []ClaimProcedure      `json:"procedure,omitempty"`
	Insurance            []ClaimInsurance      `json:"insurance"`
	Accident             *ClaimAccident        `json:"accident,omitempty"`
	Item                 []ClaimItem           `json:"item,omitempty"`
	Total                *Money                `json:"total,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Claim
type ClaimRelated struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Claim             *Reference       `json:"claim,omitempty"`
	Relationship      *CodeableConcept `json:"relationship,omitempty"`
	Reference         *Identifier      `json:"reference,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Claim
type ClaimPayee struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Party             *Reference      `json:"party,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Claim
type ClaimCareTeam struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Sequence          int              `json:"sequence"`
	Provider          Reference        `json:"provider"`
	Responsible       *bool            `json:"responsible,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Qualification     *CodeableConcept `json:"qualification,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Claim
type ClaimSupportingInfo struct {
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
	Reason            *CodeableConcept `json:"reason,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Claim
type ClaimDiagnosis struct {
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

// http://hl7.org/fhir/r4b/StructureDefinition/Claim
type ClaimProcedure struct {
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

// http://hl7.org/fhir/r4b/StructureDefinition/Claim
type ClaimInsurance struct {
	Id                  *string     `json:"id,omitempty"`
	Extension           []Extension `json:"extension,omitempty"`
	ModifierExtension   []Extension `json:"modifierExtension,omitempty"`
	Sequence            int         `json:"sequence"`
	Focal               bool        `json:"focal"`
	Identifier          *Identifier `json:"identifier,omitempty"`
	Coverage            Reference   `json:"coverage"`
	BusinessArrangement *string     `json:"businessArrangement,omitempty"`
	PreAuthRef          []string    `json:"preAuthRef,omitempty"`
	ClaimResponse       *Reference  `json:"claimResponse,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Claim
type ClaimAccident struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Date              time.Time        `json:"date,format:'2006-01-02'"`
	Type              *CodeableConcept `json:"type,omitempty"`
	LocationAddress   *Address         `json:"locationAddress,omitempty"`
	LocationReference *Reference       `json:"locationReference,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Claim
type ClaimItem struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Sequence                int               `json:"sequence"`
	CareTeamSequence        []int             `json:"careTeamSequence,omitempty"`
	DiagnosisSequence       []int             `json:"diagnosisSequence,omitempty"`
	ProcedureSequence       []int             `json:"procedureSequence,omitempty"`
	InformationSequence     []int             `json:"informationSequence,omitempty"`
	Revenue                 *CodeableConcept  `json:"revenue,omitempty"`
	Category                *CodeableConcept  `json:"category,omitempty"`
	ProductOrService        CodeableConcept   `json:"productOrService"`
	Modifier                []CodeableConcept `json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept `json:"programCode,omitempty"`
	ServicedDate            *time.Time        `json:"servicedDate,omitempty,format:'2006-01-02'"`
	ServicedPeriod          *Period           `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept  `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address          `json:"locationAddress,omitempty"`
	LocationReference       *Reference        `json:"locationReference,omitempty"`
	Quantity                *Quantity         `json:"quantity,omitempty"`
	UnitPrice               *Money            `json:"unitPrice,omitempty"`
	Factor                  *float64          `json:"factor,omitempty"`
	Net                     *Money            `json:"net,omitempty"`
	Udi                     []Reference       `json:"udi,omitempty"`
	BodySite                *CodeableConcept  `json:"bodySite,omitempty"`
	SubSite                 []CodeableConcept `json:"subSite,omitempty"`
	Encounter               []Reference       `json:"encounter,omitempty"`
	Detail                  []ClaimItemDetail `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Claim
type ClaimItemDetail struct {
	Id                *string                    `json:"id,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Sequence          int                        `json:"sequence"`
	Revenue           *CodeableConcept           `json:"revenue,omitempty"`
	Category          *CodeableConcept           `json:"category,omitempty"`
	ProductOrService  CodeableConcept            `json:"productOrService"`
	Modifier          []CodeableConcept          `json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept          `json:"programCode,omitempty"`
	Quantity          *Quantity                  `json:"quantity,omitempty"`
	UnitPrice         *Money                     `json:"unitPrice,omitempty"`
	Factor            *float64                   `json:"factor,omitempty"`
	Net               *Money                     `json:"net,omitempty"`
	Udi               []Reference                `json:"udi,omitempty"`
	SubDetail         []ClaimItemDetailSubDetail `json:"subDetail,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Claim
type ClaimItemDetailSubDetail struct {
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
}

type OtherClaim Claim

// on convert struct to json, automatically add resourceType=Claim
func (r Claim) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClaim
		ResourceType string `json:"resourceType"`
	}{
		OtherClaim:   OtherClaim(r),
		ResourceType: "Claim",
	})
}
func (r Claim) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Claim/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Claim"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Claim) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("Claim.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Claim.Status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Claim.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_SubType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Claim.SubType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.SubType", resource.SubType, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_Use(htmlAttrs string) templ.Component {
	optionsValueSet := VSClaim_use

	if resource == nil {
		return CodeSelect("Claim.Use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Claim.Use", &resource.Use, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_Created(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("Claim.Created", nil, htmlAttrs)
	}
	return DateTimeInput("Claim.Created", &resource.Created, htmlAttrs)
}
func (resource *Claim) T_Priority(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Claim.Priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Priority", &resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_FundsReserve(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Claim.FundsReserve", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.FundsReserve", resource.FundsReserve, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_RelatedRelationship(numRelated int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numRelated >= len(resource.Related) {
		return CodeableConceptSelect("Claim.Related["+strconv.Itoa(numRelated)+"].Relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Related["+strconv.Itoa(numRelated)+"].Relationship", resource.Related[numRelated].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_PayeeType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Claim.Payee.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Payee.Type", &resource.Payee.Type, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_CareTeamSequence(numCareTeam int, htmlAttrs string) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return IntInput("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Sequence", nil, htmlAttrs)
	}
	return IntInput("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Sequence", &resource.CareTeam[numCareTeam].Sequence, htmlAttrs)
}
func (resource *Claim) T_CareTeamResponsible(numCareTeam int, htmlAttrs string) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return BoolInput("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Responsible", nil, htmlAttrs)
	}
	return BoolInput("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Responsible", resource.CareTeam[numCareTeam].Responsible, htmlAttrs)
}
func (resource *Claim) T_CareTeamRole(numCareTeam int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return CodeableConceptSelect("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Role", resource.CareTeam[numCareTeam].Role, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_CareTeamQualification(numCareTeam int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return CodeableConceptSelect("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Qualification", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Qualification", resource.CareTeam[numCareTeam].Qualification, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoSequence(numSupportingInfo int, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return IntInput("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Sequence", nil, htmlAttrs)
	}
	return IntInput("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Sequence", &resource.SupportingInfo[numSupportingInfo].Sequence, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoCategory(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodeableConceptSelect("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Category", &resource.SupportingInfo[numSupportingInfo].Category, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoCode(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodeableConceptSelect("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Code", resource.SupportingInfo[numSupportingInfo].Code, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoTimingDate(numSupportingInfo int, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return DateInput("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].TimingDate", nil, htmlAttrs)
	}
	return DateInput("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].TimingDate", resource.SupportingInfo[numSupportingInfo].TimingDate, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoValueBoolean(numSupportingInfo int, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return BoolInput("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].ValueBoolean", resource.SupportingInfo[numSupportingInfo].ValueBoolean, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoValueString(numSupportingInfo int, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return StringInput("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].ValueString", nil, htmlAttrs)
	}
	return StringInput("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].ValueString", resource.SupportingInfo[numSupportingInfo].ValueString, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoReason(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodeableConceptSelect("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Reason", resource.SupportingInfo[numSupportingInfo].Reason, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_DiagnosisSequence(numDiagnosis int, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return IntInput("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Sequence", nil, htmlAttrs)
	}
	return IntInput("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Sequence", &resource.Diagnosis[numDiagnosis].Sequence, htmlAttrs)
}
func (resource *Claim) T_DiagnosisDiagnosisCodeableConcept(numDiagnosis int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].DiagnosisCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].DiagnosisCodeableConcept", &resource.Diagnosis[numDiagnosis].DiagnosisCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_DiagnosisType(numDiagnosis int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) || numType >= len(resource.Diagnosis[numDiagnosis].Type) {
		return CodeableConceptSelect("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Type["+strconv.Itoa(numType)+"]", &resource.Diagnosis[numDiagnosis].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_DiagnosisOnAdmission(numDiagnosis int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].OnAdmission", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].OnAdmission", resource.Diagnosis[numDiagnosis].OnAdmission, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_DiagnosisPackageCode(numDiagnosis int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].PackageCode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].PackageCode", resource.Diagnosis[numDiagnosis].PackageCode, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ProcedureSequence(numProcedure int, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return IntInput("Claim.Procedure["+strconv.Itoa(numProcedure)+"].Sequence", nil, htmlAttrs)
	}
	return IntInput("Claim.Procedure["+strconv.Itoa(numProcedure)+"].Sequence", &resource.Procedure[numProcedure].Sequence, htmlAttrs)
}
func (resource *Claim) T_ProcedureType(numProcedure int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) || numType >= len(resource.Procedure[numProcedure].Type) {
		return CodeableConceptSelect("Claim.Procedure["+strconv.Itoa(numProcedure)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Procedure["+strconv.Itoa(numProcedure)+"].Type["+strconv.Itoa(numType)+"]", &resource.Procedure[numProcedure].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ProcedureDate(numProcedure int, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return DateTimeInput("Claim.Procedure["+strconv.Itoa(numProcedure)+"].Date", nil, htmlAttrs)
	}
	return DateTimeInput("Claim.Procedure["+strconv.Itoa(numProcedure)+"].Date", resource.Procedure[numProcedure].Date, htmlAttrs)
}
func (resource *Claim) T_ProcedureProcedureCodeableConcept(numProcedure int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return CodeableConceptSelect("Claim.Procedure["+strconv.Itoa(numProcedure)+"].ProcedureCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Procedure["+strconv.Itoa(numProcedure)+"].ProcedureCodeableConcept", &resource.Procedure[numProcedure].ProcedureCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_InsuranceSequence(numInsurance int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return IntInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].Sequence", nil, htmlAttrs)
	}
	return IntInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].Sequence", &resource.Insurance[numInsurance].Sequence, htmlAttrs)
}
func (resource *Claim) T_InsuranceFocal(numInsurance int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return BoolInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].Focal", nil, htmlAttrs)
	}
	return BoolInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].Focal", &resource.Insurance[numInsurance].Focal, htmlAttrs)
}
func (resource *Claim) T_InsuranceBusinessArrangement(numInsurance int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return StringInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].BusinessArrangement", nil, htmlAttrs)
	}
	return StringInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].BusinessArrangement", resource.Insurance[numInsurance].BusinessArrangement, htmlAttrs)
}
func (resource *Claim) T_InsurancePreAuthRef(numInsurance int, numPreAuthRef int, htmlAttrs string) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numPreAuthRef >= len(resource.Insurance[numInsurance].PreAuthRef) {
		return StringInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].PreAuthRef["+strconv.Itoa(numPreAuthRef)+"]", nil, htmlAttrs)
	}
	return StringInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].PreAuthRef["+strconv.Itoa(numPreAuthRef)+"]", &resource.Insurance[numInsurance].PreAuthRef[numPreAuthRef], htmlAttrs)
}
func (resource *Claim) T_AccidentDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateInput("Claim.Accident.Date", nil, htmlAttrs)
	}
	return DateInput("Claim.Accident.Date", &resource.Accident.Date, htmlAttrs)
}
func (resource *Claim) T_AccidentType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("Claim.Accident.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Accident.Type", resource.Accident.Type, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemSequence(numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].Sequence", nil, htmlAttrs)
	}
	return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].Sequence", &resource.Item[numItem].Sequence, htmlAttrs)
}
func (resource *Claim) T_ItemCareTeamSequence(numItem int, numCareTeamSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numCareTeamSequence >= len(resource.Item[numItem].CareTeamSequence) {
		return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].CareTeamSequence["+strconv.Itoa(numCareTeamSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].CareTeamSequence["+strconv.Itoa(numCareTeamSequence)+"]", &resource.Item[numItem].CareTeamSequence[numCareTeamSequence], htmlAttrs)
}
func (resource *Claim) T_ItemDiagnosisSequence(numItem int, numDiagnosisSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDiagnosisSequence >= len(resource.Item[numItem].DiagnosisSequence) {
		return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].DiagnosisSequence["+strconv.Itoa(numDiagnosisSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].DiagnosisSequence["+strconv.Itoa(numDiagnosisSequence)+"]", &resource.Item[numItem].DiagnosisSequence[numDiagnosisSequence], htmlAttrs)
}
func (resource *Claim) T_ItemProcedureSequence(numItem int, numProcedureSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numProcedureSequence >= len(resource.Item[numItem].ProcedureSequence) {
		return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].ProcedureSequence["+strconv.Itoa(numProcedureSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].ProcedureSequence["+strconv.Itoa(numProcedureSequence)+"]", &resource.Item[numItem].ProcedureSequence[numProcedureSequence], htmlAttrs)
}
func (resource *Claim) T_ItemInformationSequence(numItem int, numInformationSequence int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numInformationSequence >= len(resource.Item[numItem].InformationSequence) {
		return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].InformationSequence["+strconv.Itoa(numInformationSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].InformationSequence["+strconv.Itoa(numInformationSequence)+"]", &resource.Item[numItem].InformationSequence[numInformationSequence], htmlAttrs)
}
func (resource *Claim) T_ItemRevenue(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Revenue", resource.Item[numItem].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemCategory(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Category", resource.Item[numItem].Category, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemProductOrService(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].ProductOrService", &resource.Item[numItem].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemModifier(numItem int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numModifier >= len(resource.Item[numItem].Modifier) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemProgramCode(numItem int, numProgramCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numProgramCode >= len(resource.Item[numItem].ProgramCode) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemServicedDate(numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return DateInput("Claim.Item["+strconv.Itoa(numItem)+"].ServicedDate", nil, htmlAttrs)
	}
	return DateInput("Claim.Item["+strconv.Itoa(numItem)+"].ServicedDate", resource.Item[numItem].ServicedDate, htmlAttrs)
}
func (resource *Claim) T_ItemLocationCodeableConcept(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].LocationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].LocationCodeableConcept", resource.Item[numItem].LocationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemFactor(numItem int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return Float64Input("Claim.Item["+strconv.Itoa(numItem)+"].Factor", nil, htmlAttrs)
	}
	return Float64Input("Claim.Item["+strconv.Itoa(numItem)+"].Factor", resource.Item[numItem].Factor, htmlAttrs)
}
func (resource *Claim) T_ItemBodySite(numItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].BodySite", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].BodySite", resource.Item[numItem].BodySite, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemSubSite(numItem int, numSubSite int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numSubSite >= len(resource.Item[numItem].SubSite) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].SubSite["+strconv.Itoa(numSubSite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].SubSite["+strconv.Itoa(numSubSite)+"]", &resource.Item[numItem].SubSite[numSubSite], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSequence(numItem int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Sequence", nil, htmlAttrs)
	}
	return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Sequence", &resource.Item[numItem].Detail[numDetail].Sequence, htmlAttrs)
}
func (resource *Claim) T_ItemDetailRevenue(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Revenue", resource.Item[numItem].Detail[numDetail].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailCategory(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Category", resource.Item[numItem].Detail[numDetail].Category, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailProductOrService(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProductOrService", &resource.Item[numItem].Detail[numDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailModifier(numItem int, numDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numModifier >= len(resource.Item[numItem].Detail[numDetail].Modifier) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Detail[numDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailProgramCode(numItem int, numDetail int, numProgramCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numProgramCode >= len(resource.Item[numItem].Detail[numDetail].ProgramCode) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].Detail[numDetail].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailFactor(numItem int, numDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return Float64Input("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Factor", nil, htmlAttrs)
	}
	return Float64Input("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Factor", resource.Item[numItem].Detail[numDetail].Factor, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailSequence(numItem int, numDetail int, numSubDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Sequence", nil, htmlAttrs)
	}
	return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Sequence", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Sequence, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailRevenue(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Revenue", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailCategory(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Category", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Category, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailProductOrService(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailModifier(numItem int, numDetail int, numSubDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numModifier >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Modifier) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailProgramCode(numItem int, numDetail int, numSubDetail int, numProgramCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numProgramCode >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProgramCode) {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailFactor(numItem int, numDetail int, numSubDetail int, htmlAttrs string) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return Float64Input("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Factor", nil, htmlAttrs)
	}
	return Float64Input("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Factor", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Factor, htmlAttrs)
}
