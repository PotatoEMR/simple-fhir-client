package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Claim
type Claim struct {
	Id                    *string               `json:"id,omitempty"`
	Meta                  *Meta                 `json:"meta,omitempty"`
	ImplicitRules         *string               `json:"implicitRules,omitempty"`
	Language              *string               `json:"language,omitempty"`
	Text                  *Narrative            `json:"text,omitempty"`
	Contained             []Resource            `json:"contained,omitempty"`
	Extension             []Extension           `json:"extension,omitempty"`
	ModifierExtension     []Extension           `json:"modifierExtension,omitempty"`
	Identifier            []Identifier          `json:"identifier,omitempty"`
	TraceNumber           []Identifier          `json:"traceNumber,omitempty"`
	Status                string                `json:"status"`
	Type                  CodeableConcept       `json:"type"`
	SubType               *CodeableConcept      `json:"subType,omitempty"`
	Use                   string                `json:"use"`
	Patient               Reference             `json:"patient"`
	BillablePeriod        *Period               `json:"billablePeriod,omitempty"`
	Created               string                `json:"created"`
	Enterer               *Reference            `json:"enterer,omitempty"`
	Insurer               *Reference            `json:"insurer,omitempty"`
	Provider              *Reference            `json:"provider,omitempty"`
	Priority              *CodeableConcept      `json:"priority,omitempty"`
	FundsReserve          *CodeableConcept      `json:"fundsReserve,omitempty"`
	Related               []ClaimRelated        `json:"related,omitempty"`
	Prescription          *Reference            `json:"prescription,omitempty"`
	OriginalPrescription  *Reference            `json:"originalPrescription,omitempty"`
	Payee                 *ClaimPayee           `json:"payee,omitempty"`
	Referral              *Reference            `json:"referral,omitempty"`
	Encounter             []Reference           `json:"encounter,omitempty"`
	Facility              *Reference            `json:"facility,omitempty"`
	DiagnosisRelatedGroup *CodeableConcept      `json:"diagnosisRelatedGroup,omitempty"`
	Event                 []ClaimEvent          `json:"event,omitempty"`
	CareTeam              []ClaimCareTeam       `json:"careTeam,omitempty"`
	SupportingInfo        []ClaimSupportingInfo `json:"supportingInfo,omitempty"`
	Diagnosis             []ClaimDiagnosis      `json:"diagnosis,omitempty"`
	Procedure             []ClaimProcedure      `json:"procedure,omitempty"`
	Insurance             []ClaimInsurance      `json:"insurance,omitempty"`
	Accident              *ClaimAccident        `json:"accident,omitempty"`
	PatientPaid           *Money                `json:"patientPaid,omitempty"`
	Item                  []ClaimItem           `json:"item,omitempty"`
	Total                 *Money                `json:"total,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Claim
type ClaimRelated struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Claim             *Reference       `json:"claim,omitempty"`
	Relationship      *CodeableConcept `json:"relationship,omitempty"`
	Reference         *Identifier      `json:"reference,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Claim
type ClaimPayee struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Party             *Reference      `json:"party,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Claim
type ClaimEvent struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	WhenDateTime      string          `json:"whenDateTime"`
	WhenPeriod        Period          `json:"whenPeriod"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Claim
type ClaimCareTeam struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Sequence          int              `json:"sequence"`
	Provider          Reference        `json:"provider"`
	Responsible       *bool            `json:"responsible,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Specialty         *CodeableConcept `json:"specialty,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Claim
type ClaimSupportingInfo struct {
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
	Reason            *CodeableConcept `json:"reason,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Claim
type ClaimDiagnosis struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Sequence                 int               `json:"sequence"`
	DiagnosisCodeableConcept CodeableConcept   `json:"diagnosisCodeableConcept"`
	DiagnosisReference       Reference         `json:"diagnosisReference"`
	Type                     []CodeableConcept `json:"type,omitempty"`
	OnAdmission              *CodeableConcept  `json:"onAdmission,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Claim
type ClaimProcedure struct {
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

// http://hl7.org/fhir/r5/StructureDefinition/Claim
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

// http://hl7.org/fhir/r5/StructureDefinition/Claim
type ClaimAccident struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Date              string           `json:"date"`
	Type              *CodeableConcept `json:"type,omitempty"`
	LocationAddress   *Address         `json:"locationAddress,omitempty"`
	LocationReference *Reference       `json:"locationReference,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Claim
type ClaimItem struct {
	Id                      *string             `json:"id,omitempty"`
	Extension               []Extension         `json:"extension,omitempty"`
	ModifierExtension       []Extension         `json:"modifierExtension,omitempty"`
	Sequence                int                 `json:"sequence"`
	TraceNumber             []Identifier        `json:"traceNumber,omitempty"`
	CareTeamSequence        []int               `json:"careTeamSequence,omitempty"`
	DiagnosisSequence       []int               `json:"diagnosisSequence,omitempty"`
	ProcedureSequence       []int               `json:"procedureSequence,omitempty"`
	InformationSequence     []int               `json:"informationSequence,omitempty"`
	Revenue                 *CodeableConcept    `json:"revenue,omitempty"`
	Category                *CodeableConcept    `json:"category,omitempty"`
	ProductOrService        *CodeableConcept    `json:"productOrService,omitempty"`
	ProductOrServiceEnd     *CodeableConcept    `json:"productOrServiceEnd,omitempty"`
	Request                 []Reference         `json:"request,omitempty"`
	Modifier                []CodeableConcept   `json:"modifier,omitempty"`
	ProgramCode             []CodeableConcept   `json:"programCode,omitempty"`
	ServicedDate            *string             `json:"servicedDate,omitempty"`
	ServicedPeriod          *Period             `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept    `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address            `json:"locationAddress,omitempty"`
	LocationReference       *Reference          `json:"locationReference,omitempty"`
	PatientPaid             *Money              `json:"patientPaid,omitempty"`
	Quantity                *Quantity           `json:"quantity,omitempty"`
	UnitPrice               *Money              `json:"unitPrice,omitempty"`
	Factor                  *float64            `json:"factor,omitempty"`
	Tax                     *Money              `json:"tax,omitempty"`
	Net                     *Money              `json:"net,omitempty"`
	Udi                     []Reference         `json:"udi,omitempty"`
	BodySite                []ClaimItemBodySite `json:"bodySite,omitempty"`
	Encounter               []Reference         `json:"encounter,omitempty"`
	Detail                  []ClaimItemDetail   `json:"detail,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Claim
type ClaimItemBodySite struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Site              []CodeableReference `json:"site"`
	SubSite           []CodeableConcept   `json:"subSite,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Claim
type ClaimItemDetail struct {
	Id                  *string                    `json:"id,omitempty"`
	Extension           []Extension                `json:"extension,omitempty"`
	ModifierExtension   []Extension                `json:"modifierExtension,omitempty"`
	Sequence            int                        `json:"sequence"`
	TraceNumber         []Identifier               `json:"traceNumber,omitempty"`
	Revenue             *CodeableConcept           `json:"revenue,omitempty"`
	Category            *CodeableConcept           `json:"category,omitempty"`
	ProductOrService    *CodeableConcept           `json:"productOrService,omitempty"`
	ProductOrServiceEnd *CodeableConcept           `json:"productOrServiceEnd,omitempty"`
	Modifier            []CodeableConcept          `json:"modifier,omitempty"`
	ProgramCode         []CodeableConcept          `json:"programCode,omitempty"`
	PatientPaid         *Money                     `json:"patientPaid,omitempty"`
	Quantity            *Quantity                  `json:"quantity,omitempty"`
	UnitPrice           *Money                     `json:"unitPrice,omitempty"`
	Factor              *float64                   `json:"factor,omitempty"`
	Tax                 *Money                     `json:"tax,omitempty"`
	Net                 *Money                     `json:"net,omitempty"`
	Udi                 []Reference                `json:"udi,omitempty"`
	SubDetail           []ClaimItemDetailSubDetail `json:"subDetail,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Claim
type ClaimItemDetailSubDetail struct {
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

func (resource *Claim) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Claim.Id", nil)
	}
	return StringInput("Claim.Id", resource.Id)
}
func (resource *Claim) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Claim.ImplicitRules", nil)
	}
	return StringInput("Claim.ImplicitRules", resource.ImplicitRules)
}
func (resource *Claim) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Claim.Language", nil, optionsValueSet)
	}
	return CodeSelect("Claim.Language", resource.Language, optionsValueSet)
}
func (resource *Claim) T_Status() templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("Claim.Status", nil, optionsValueSet)
	}
	return CodeSelect("Claim.Status", &resource.Status, optionsValueSet)
}
func (resource *Claim) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Claim.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Type", &resource.Type, optionsValueSet)
}
func (resource *Claim) T_SubType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Claim.SubType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.SubType", resource.SubType, optionsValueSet)
}
func (resource *Claim) T_Use() templ.Component {
	optionsValueSet := VSClaim_use

	if resource == nil {
		return CodeSelect("Claim.Use", nil, optionsValueSet)
	}
	return CodeSelect("Claim.Use", &resource.Use, optionsValueSet)
}
func (resource *Claim) T_Created() templ.Component {

	if resource == nil {
		return StringInput("Claim.Created", nil)
	}
	return StringInput("Claim.Created", &resource.Created)
}
func (resource *Claim) T_Priority(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Claim.Priority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Priority", resource.Priority, optionsValueSet)
}
func (resource *Claim) T_FundsReserve(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Claim.FundsReserve", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.FundsReserve", resource.FundsReserve, optionsValueSet)
}
func (resource *Claim) T_DiagnosisRelatedGroup(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Claim.DiagnosisRelatedGroup", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.DiagnosisRelatedGroup", resource.DiagnosisRelatedGroup, optionsValueSet)
}
func (resource *Claim) T_RelatedId(numRelated int) templ.Component {

	if resource == nil || len(resource.Related) >= numRelated {
		return StringInput("Claim.Related["+strconv.Itoa(numRelated)+"].Id", nil)
	}
	return StringInput("Claim.Related["+strconv.Itoa(numRelated)+"].Id", resource.Related[numRelated].Id)
}
func (resource *Claim) T_RelatedRelationship(numRelated int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Related) >= numRelated {
		return CodeableConceptSelect("Claim.Related["+strconv.Itoa(numRelated)+"].Relationship", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Related["+strconv.Itoa(numRelated)+"].Relationship", resource.Related[numRelated].Relationship, optionsValueSet)
}
func (resource *Claim) T_PayeeId() templ.Component {

	if resource == nil {
		return StringInput("Claim.Payee.Id", nil)
	}
	return StringInput("Claim.Payee.Id", resource.Payee.Id)
}
func (resource *Claim) T_PayeeType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Claim.Payee.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Payee.Type", &resource.Payee.Type, optionsValueSet)
}
func (resource *Claim) T_EventId(numEvent int) templ.Component {

	if resource == nil || len(resource.Event) >= numEvent {
		return StringInput("Claim.Event["+strconv.Itoa(numEvent)+"].Id", nil)
	}
	return StringInput("Claim.Event["+strconv.Itoa(numEvent)+"].Id", resource.Event[numEvent].Id)
}
func (resource *Claim) T_EventType(numEvent int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Event) >= numEvent {
		return CodeableConceptSelect("Claim.Event["+strconv.Itoa(numEvent)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Event["+strconv.Itoa(numEvent)+"].Type", &resource.Event[numEvent].Type, optionsValueSet)
}
func (resource *Claim) T_CareTeamId(numCareTeam int) templ.Component {

	if resource == nil || len(resource.CareTeam) >= numCareTeam {
		return StringInput("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Id", nil)
	}
	return StringInput("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Id", resource.CareTeam[numCareTeam].Id)
}
func (resource *Claim) T_CareTeamSequence(numCareTeam int) templ.Component {

	if resource == nil || len(resource.CareTeam) >= numCareTeam {
		return IntInput("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Sequence", nil)
	}
	return IntInput("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Sequence", &resource.CareTeam[numCareTeam].Sequence)
}
func (resource *Claim) T_CareTeamResponsible(numCareTeam int) templ.Component {

	if resource == nil || len(resource.CareTeam) >= numCareTeam {
		return BoolInput("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Responsible", nil)
	}
	return BoolInput("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Responsible", resource.CareTeam[numCareTeam].Responsible)
}
func (resource *Claim) T_CareTeamRole(numCareTeam int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CareTeam) >= numCareTeam {
		return CodeableConceptSelect("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Role", resource.CareTeam[numCareTeam].Role, optionsValueSet)
}
func (resource *Claim) T_CareTeamSpecialty(numCareTeam int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.CareTeam) >= numCareTeam {
		return CodeableConceptSelect("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Specialty", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.CareTeam["+strconv.Itoa(numCareTeam)+"].Specialty", resource.CareTeam[numCareTeam].Specialty, optionsValueSet)
}
func (resource *Claim) T_SupportingInfoId(numSupportingInfo int) templ.Component {

	if resource == nil || len(resource.SupportingInfo) >= numSupportingInfo {
		return StringInput("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Id", nil)
	}
	return StringInput("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Id", resource.SupportingInfo[numSupportingInfo].Id)
}
func (resource *Claim) T_SupportingInfoSequence(numSupportingInfo int) templ.Component {

	if resource == nil || len(resource.SupportingInfo) >= numSupportingInfo {
		return IntInput("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Sequence", nil)
	}
	return IntInput("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Sequence", &resource.SupportingInfo[numSupportingInfo].Sequence)
}
func (resource *Claim) T_SupportingInfoCategory(numSupportingInfo int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SupportingInfo) >= numSupportingInfo {
		return CodeableConceptSelect("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Category", &resource.SupportingInfo[numSupportingInfo].Category, optionsValueSet)
}
func (resource *Claim) T_SupportingInfoCode(numSupportingInfo int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SupportingInfo) >= numSupportingInfo {
		return CodeableConceptSelect("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Code", resource.SupportingInfo[numSupportingInfo].Code, optionsValueSet)
}
func (resource *Claim) T_SupportingInfoReason(numSupportingInfo int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SupportingInfo) >= numSupportingInfo {
		return CodeableConceptSelect("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.SupportingInfo["+strconv.Itoa(numSupportingInfo)+"].Reason", resource.SupportingInfo[numSupportingInfo].Reason, optionsValueSet)
}
func (resource *Claim) T_DiagnosisId(numDiagnosis int) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis {
		return StringInput("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Id", nil)
	}
	return StringInput("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Id", resource.Diagnosis[numDiagnosis].Id)
}
func (resource *Claim) T_DiagnosisSequence(numDiagnosis int) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis {
		return IntInput("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Sequence", nil)
	}
	return IntInput("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Sequence", &resource.Diagnosis[numDiagnosis].Sequence)
}
func (resource *Claim) T_DiagnosisType(numDiagnosis int, numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis || len(resource.Diagnosis[numDiagnosis].Type) >= numType {
		return CodeableConceptSelect("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].Type["+strconv.Itoa(numType)+"]", &resource.Diagnosis[numDiagnosis].Type[numType], optionsValueSet)
}
func (resource *Claim) T_DiagnosisOnAdmission(numDiagnosis int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Diagnosis) >= numDiagnosis {
		return CodeableConceptSelect("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].OnAdmission", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Diagnosis["+strconv.Itoa(numDiagnosis)+"].OnAdmission", resource.Diagnosis[numDiagnosis].OnAdmission, optionsValueSet)
}
func (resource *Claim) T_ProcedureId(numProcedure int) templ.Component {

	if resource == nil || len(resource.Procedure) >= numProcedure {
		return StringInput("Claim.Procedure["+strconv.Itoa(numProcedure)+"].Id", nil)
	}
	return StringInput("Claim.Procedure["+strconv.Itoa(numProcedure)+"].Id", resource.Procedure[numProcedure].Id)
}
func (resource *Claim) T_ProcedureSequence(numProcedure int) templ.Component {

	if resource == nil || len(resource.Procedure) >= numProcedure {
		return IntInput("Claim.Procedure["+strconv.Itoa(numProcedure)+"].Sequence", nil)
	}
	return IntInput("Claim.Procedure["+strconv.Itoa(numProcedure)+"].Sequence", &resource.Procedure[numProcedure].Sequence)
}
func (resource *Claim) T_ProcedureType(numProcedure int, numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Procedure) >= numProcedure || len(resource.Procedure[numProcedure].Type) >= numType {
		return CodeableConceptSelect("Claim.Procedure["+strconv.Itoa(numProcedure)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Procedure["+strconv.Itoa(numProcedure)+"].Type["+strconv.Itoa(numType)+"]", &resource.Procedure[numProcedure].Type[numType], optionsValueSet)
}
func (resource *Claim) T_ProcedureDate(numProcedure int) templ.Component {

	if resource == nil || len(resource.Procedure) >= numProcedure {
		return StringInput("Claim.Procedure["+strconv.Itoa(numProcedure)+"].Date", nil)
	}
	return StringInput("Claim.Procedure["+strconv.Itoa(numProcedure)+"].Date", resource.Procedure[numProcedure].Date)
}
func (resource *Claim) T_InsuranceId(numInsurance int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance {
		return StringInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].Id", nil)
	}
	return StringInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].Id", resource.Insurance[numInsurance].Id)
}
func (resource *Claim) T_InsuranceSequence(numInsurance int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance {
		return IntInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].Sequence", nil)
	}
	return IntInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].Sequence", &resource.Insurance[numInsurance].Sequence)
}
func (resource *Claim) T_InsuranceFocal(numInsurance int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance {
		return BoolInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].Focal", nil)
	}
	return BoolInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].Focal", &resource.Insurance[numInsurance].Focal)
}
func (resource *Claim) T_InsuranceBusinessArrangement(numInsurance int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance {
		return StringInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].BusinessArrangement", nil)
	}
	return StringInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].BusinessArrangement", resource.Insurance[numInsurance].BusinessArrangement)
}
func (resource *Claim) T_InsurancePreAuthRef(numInsurance int, numPreAuthRef int) templ.Component {

	if resource == nil || len(resource.Insurance) >= numInsurance || len(resource.Insurance[numInsurance].PreAuthRef) >= numPreAuthRef {
		return StringInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].PreAuthRef["+strconv.Itoa(numPreAuthRef)+"]", nil)
	}
	return StringInput("Claim.Insurance["+strconv.Itoa(numInsurance)+"].PreAuthRef["+strconv.Itoa(numPreAuthRef)+"]", &resource.Insurance[numInsurance].PreAuthRef[numPreAuthRef])
}
func (resource *Claim) T_AccidentId() templ.Component {

	if resource == nil {
		return StringInput("Claim.Accident.Id", nil)
	}
	return StringInput("Claim.Accident.Id", resource.Accident.Id)
}
func (resource *Claim) T_AccidentDate() templ.Component {

	if resource == nil {
		return StringInput("Claim.Accident.Date", nil)
	}
	return StringInput("Claim.Accident.Date", &resource.Accident.Date)
}
func (resource *Claim) T_AccidentType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Claim.Accident.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Accident.Type", resource.Accident.Type, optionsValueSet)
}
func (resource *Claim) T_ItemId(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return StringInput("Claim.Item["+strconv.Itoa(numItem)+"].Id", nil)
	}
	return StringInput("Claim.Item["+strconv.Itoa(numItem)+"].Id", resource.Item[numItem].Id)
}
func (resource *Claim) T_ItemSequence(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].Sequence", nil)
	}
	return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].Sequence", &resource.Item[numItem].Sequence)
}
func (resource *Claim) T_ItemCareTeamSequence(numItem int, numCareTeamSequence int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].CareTeamSequence) >= numCareTeamSequence {
		return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].CareTeamSequence["+strconv.Itoa(numCareTeamSequence)+"]", nil)
	}
	return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].CareTeamSequence["+strconv.Itoa(numCareTeamSequence)+"]", &resource.Item[numItem].CareTeamSequence[numCareTeamSequence])
}
func (resource *Claim) T_ItemDiagnosisSequence(numItem int, numDiagnosisSequence int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].DiagnosisSequence) >= numDiagnosisSequence {
		return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].DiagnosisSequence["+strconv.Itoa(numDiagnosisSequence)+"]", nil)
	}
	return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].DiagnosisSequence["+strconv.Itoa(numDiagnosisSequence)+"]", &resource.Item[numItem].DiagnosisSequence[numDiagnosisSequence])
}
func (resource *Claim) T_ItemProcedureSequence(numItem int, numProcedureSequence int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].ProcedureSequence) >= numProcedureSequence {
		return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].ProcedureSequence["+strconv.Itoa(numProcedureSequence)+"]", nil)
	}
	return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].ProcedureSequence["+strconv.Itoa(numProcedureSequence)+"]", &resource.Item[numItem].ProcedureSequence[numProcedureSequence])
}
func (resource *Claim) T_ItemInformationSequence(numItem int, numInformationSequence int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].InformationSequence) >= numInformationSequence {
		return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].InformationSequence["+strconv.Itoa(numInformationSequence)+"]", nil)
	}
	return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].InformationSequence["+strconv.Itoa(numInformationSequence)+"]", &resource.Item[numItem].InformationSequence[numInformationSequence])
}
func (resource *Claim) T_ItemRevenue(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Revenue", resource.Item[numItem].Revenue, optionsValueSet)
}
func (resource *Claim) T_ItemCategory(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Category", resource.Item[numItem].Category, optionsValueSet)
}
func (resource *Claim) T_ItemProductOrService(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].ProductOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].ProductOrService", resource.Item[numItem].ProductOrService, optionsValueSet)
}
func (resource *Claim) T_ItemProductOrServiceEnd(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].ProductOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].ProductOrServiceEnd", resource.Item[numItem].ProductOrServiceEnd, optionsValueSet)
}
func (resource *Claim) T_ItemModifier(numItem int, numModifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Modifier) >= numModifier {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Modifier[numModifier], optionsValueSet)
}
func (resource *Claim) T_ItemProgramCode(numItem int, numProgramCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].ProgramCode) >= numProgramCode {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].ProgramCode[numProgramCode], optionsValueSet)
}
func (resource *Claim) T_ItemFactor(numItem int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem {
		return Float64Input("Claim.Item["+strconv.Itoa(numItem)+"].Factor", nil)
	}
	return Float64Input("Claim.Item["+strconv.Itoa(numItem)+"].Factor", resource.Item[numItem].Factor)
}
func (resource *Claim) T_ItemBodySiteId(numItem int, numBodySite int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].BodySite) >= numBodySite {
		return StringInput("Claim.Item["+strconv.Itoa(numItem)+"].BodySite["+strconv.Itoa(numBodySite)+"].Id", nil)
	}
	return StringInput("Claim.Item["+strconv.Itoa(numItem)+"].BodySite["+strconv.Itoa(numBodySite)+"].Id", resource.Item[numItem].BodySite[numBodySite].Id)
}
func (resource *Claim) T_ItemBodySiteSubSite(numItem int, numBodySite int, numSubSite int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].BodySite) >= numBodySite || len(resource.Item[numItem].BodySite[numBodySite].SubSite) >= numSubSite {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].BodySite["+strconv.Itoa(numBodySite)+"].SubSite["+strconv.Itoa(numSubSite)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].BodySite["+strconv.Itoa(numBodySite)+"].SubSite["+strconv.Itoa(numSubSite)+"]", &resource.Item[numItem].BodySite[numBodySite].SubSite[numSubSite], optionsValueSet)
}
func (resource *Claim) T_ItemDetailId(numItem int, numDetail int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail {
		return StringInput("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Id", nil)
	}
	return StringInput("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Id", resource.Item[numItem].Detail[numDetail].Id)
}
func (resource *Claim) T_ItemDetailSequence(numItem int, numDetail int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail {
		return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Sequence", nil)
	}
	return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Sequence", &resource.Item[numItem].Detail[numDetail].Sequence)
}
func (resource *Claim) T_ItemDetailRevenue(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Revenue", resource.Item[numItem].Detail[numDetail].Revenue, optionsValueSet)
}
func (resource *Claim) T_ItemDetailCategory(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Category", resource.Item[numItem].Detail[numDetail].Category, optionsValueSet)
}
func (resource *Claim) T_ItemDetailProductOrService(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProductOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProductOrService", resource.Item[numItem].Detail[numDetail].ProductOrService, optionsValueSet)
}
func (resource *Claim) T_ItemDetailProductOrServiceEnd(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProductOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProductOrServiceEnd", resource.Item[numItem].Detail[numDetail].ProductOrServiceEnd, optionsValueSet)
}
func (resource *Claim) T_ItemDetailModifier(numItem int, numDetail int, numModifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].Modifier) >= numModifier {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Detail[numDetail].Modifier[numModifier], optionsValueSet)
}
func (resource *Claim) T_ItemDetailProgramCode(numItem int, numDetail int, numProgramCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].ProgramCode) >= numProgramCode {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].Detail[numDetail].ProgramCode[numProgramCode], optionsValueSet)
}
func (resource *Claim) T_ItemDetailFactor(numItem int, numDetail int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail {
		return Float64Input("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Factor", nil)
	}
	return Float64Input("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].Factor", resource.Item[numItem].Detail[numDetail].Factor)
}
func (resource *Claim) T_ItemDetailSubDetailId(numItem int, numDetail int, numSubDetail int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return StringInput("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Id", nil)
	}
	return StringInput("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Id", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Id)
}
func (resource *Claim) T_ItemDetailSubDetailSequence(numItem int, numDetail int, numSubDetail int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Sequence", nil)
	}
	return IntInput("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Sequence", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Sequence)
}
func (resource *Claim) T_ItemDetailSubDetailRevenue(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Revenue", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Revenue, optionsValueSet)
}
func (resource *Claim) T_ItemDetailSubDetailCategory(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Category", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Category, optionsValueSet)
}
func (resource *Claim) T_ItemDetailSubDetailProductOrService(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrService", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet)
}
func (resource *Claim) T_ItemDetailSubDetailProductOrServiceEnd(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrServiceEnd", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProductOrServiceEnd", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrServiceEnd, optionsValueSet)
}
func (resource *Claim) T_ItemDetailSubDetailModifier(numItem int, numDetail int, numSubDetail int, numModifier int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Modifier) >= numModifier {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[numModifier], optionsValueSet)
}
func (resource *Claim) T_ItemDetailSubDetailProgramCode(numItem int, numDetail int, numSubDetail int, numProgramCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProgramCode) >= numProgramCode {
		return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].ProgramCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProgramCode[numProgramCode], optionsValueSet)
}
func (resource *Claim) T_ItemDetailSubDetailFactor(numItem int, numDetail int, numSubDetail int) templ.Component {

	if resource == nil || len(resource.Item) >= numItem || len(resource.Item[numItem].Detail) >= numDetail || len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return Float64Input("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Factor", nil)
	}
	return Float64Input("Claim.Item["+strconv.Itoa(numItem)+"].Detail["+strconv.Itoa(numDetail)+"].SubDetail["+strconv.Itoa(numSubDetail)+"].Factor", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Factor)
}
