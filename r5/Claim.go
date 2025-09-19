package r5

//generated with command go run ./bultaoreune
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
	Created               FhirDateTime          `json:"created"`
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
	WhenDateTime      FhirDateTime    `json:"whenDateTime"`
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
	TimingDate        *FhirDate        `json:"timingDate,omitempty"`
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
	Date                     *FhirDateTime     `json:"date,omitempty"`
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
	Date              FhirDate         `json:"date"`
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
	ServicedDate            *FhirDate           `json:"servicedDate,omitempty"`
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
func (resource *Claim) T_TraceNumber(numTraceNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numTraceNumber >= len(resource.TraceNumber) {
		return IdentifierInput("traceNumber["+strconv.Itoa(numTraceNumber)+"]", nil, htmlAttrs)
	}
	return IdentifierInput("traceNumber["+strconv.Itoa(numTraceNumber)+"]", &resource.TraceNumber[numTraceNumber], htmlAttrs)
}
func (resource *Claim) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", &resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_SubType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("subType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subType", resource.SubType, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_Use(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSClaim_use

	if resource == nil {
		return CodeSelect("use", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("use", &resource.Use, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_Patient(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("patient", nil, htmlAttrs)
	}
	return ReferenceInput("patient", &resource.Patient, htmlAttrs)
}
func (resource *Claim) T_BillablePeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("billablePeriod", nil, htmlAttrs)
	}
	return PeriodInput("billablePeriod", resource.BillablePeriod, htmlAttrs)
}
func (resource *Claim) T_Created(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("created", nil, htmlAttrs)
	}
	return FhirDateTimeInput("created", &resource.Created, htmlAttrs)
}
func (resource *Claim) T_Enterer(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("enterer", nil, htmlAttrs)
	}
	return ReferenceInput("enterer", resource.Enterer, htmlAttrs)
}
func (resource *Claim) T_Insurer(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("insurer", nil, htmlAttrs)
	}
	return ReferenceInput("insurer", resource.Insurer, htmlAttrs)
}
func (resource *Claim) T_Provider(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("provider", nil, htmlAttrs)
	}
	return ReferenceInput("provider", resource.Provider, htmlAttrs)
}
func (resource *Claim) T_Priority(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("priority", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("priority", resource.Priority, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_FundsReserve(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("fundsReserve", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("fundsReserve", resource.FundsReserve, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_Prescription(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("prescription", nil, htmlAttrs)
	}
	return ReferenceInput("prescription", resource.Prescription, htmlAttrs)
}
func (resource *Claim) T_OriginalPrescription(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("originalPrescription", nil, htmlAttrs)
	}
	return ReferenceInput("originalPrescription", resource.OriginalPrescription, htmlAttrs)
}
func (resource *Claim) T_Referral(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("referral", nil, htmlAttrs)
	}
	return ReferenceInput("referral", resource.Referral, htmlAttrs)
}
func (resource *Claim) T_Encounter(numEncounter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEncounter >= len(resource.Encounter) {
		return ReferenceInput("encounter["+strconv.Itoa(numEncounter)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("encounter["+strconv.Itoa(numEncounter)+"]", &resource.Encounter[numEncounter], htmlAttrs)
}
func (resource *Claim) T_Facility(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("facility", nil, htmlAttrs)
	}
	return ReferenceInput("facility", resource.Facility, htmlAttrs)
}
func (resource *Claim) T_DiagnosisRelatedGroup(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("diagnosisRelatedGroup", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosisRelatedGroup", resource.DiagnosisRelatedGroup, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_PatientPaid(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return MoneyInput("patientPaid", nil, htmlAttrs)
	}
	return MoneyInput("patientPaid", resource.PatientPaid, htmlAttrs)
}
func (resource *Claim) T_Total(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return MoneyInput("total", nil, htmlAttrs)
	}
	return MoneyInput("total", resource.Total, htmlAttrs)
}
func (resource *Claim) T_RelatedClaim(numRelated int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelated >= len(resource.Related) {
		return ReferenceInput("related["+strconv.Itoa(numRelated)+"].claim", nil, htmlAttrs)
	}
	return ReferenceInput("related["+strconv.Itoa(numRelated)+"].claim", resource.Related[numRelated].Claim, htmlAttrs)
}
func (resource *Claim) T_RelatedRelationship(numRelated int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelated >= len(resource.Related) {
		return CodeableConceptSelect("related["+strconv.Itoa(numRelated)+"].relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("related["+strconv.Itoa(numRelated)+"].relationship", resource.Related[numRelated].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_RelatedReference(numRelated int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelated >= len(resource.Related) {
		return IdentifierInput("related["+strconv.Itoa(numRelated)+"].reference", nil, htmlAttrs)
	}
	return IdentifierInput("related["+strconv.Itoa(numRelated)+"].reference", resource.Related[numRelated].Reference, htmlAttrs)
}
func (resource *Claim) T_PayeeType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("payee.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("payee.type", &resource.Payee.Type, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_PayeeParty(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("payee.party", nil, htmlAttrs)
	}
	return ReferenceInput("payee.party", resource.Payee.Party, htmlAttrs)
}
func (resource *Claim) T_EventType(numEvent int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return CodeableConceptSelect("event["+strconv.Itoa(numEvent)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("event["+strconv.Itoa(numEvent)+"].type", &resource.Event[numEvent].Type, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_EventWhenDateTime(numEvent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return FhirDateTimeInput("event["+strconv.Itoa(numEvent)+"].whenDateTime", nil, htmlAttrs)
	}
	return FhirDateTimeInput("event["+strconv.Itoa(numEvent)+"].whenDateTime", &resource.Event[numEvent].WhenDateTime, htmlAttrs)
}
func (resource *Claim) T_EventWhenPeriod(numEvent int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numEvent >= len(resource.Event) {
		return PeriodInput("event["+strconv.Itoa(numEvent)+"].whenPeriod", nil, htmlAttrs)
	}
	return PeriodInput("event["+strconv.Itoa(numEvent)+"].whenPeriod", &resource.Event[numEvent].WhenPeriod, htmlAttrs)
}
func (resource *Claim) T_CareTeamSequence(numCareTeam int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return IntInput("careTeam["+strconv.Itoa(numCareTeam)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("careTeam["+strconv.Itoa(numCareTeam)+"].sequence", &resource.CareTeam[numCareTeam].Sequence, htmlAttrs)
}
func (resource *Claim) T_CareTeamProvider(numCareTeam int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return ReferenceInput("careTeam["+strconv.Itoa(numCareTeam)+"].provider", nil, htmlAttrs)
	}
	return ReferenceInput("careTeam["+strconv.Itoa(numCareTeam)+"].provider", &resource.CareTeam[numCareTeam].Provider, htmlAttrs)
}
func (resource *Claim) T_CareTeamResponsible(numCareTeam int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return BoolInput("careTeam["+strconv.Itoa(numCareTeam)+"].responsible", nil, htmlAttrs)
	}
	return BoolInput("careTeam["+strconv.Itoa(numCareTeam)+"].responsible", resource.CareTeam[numCareTeam].Responsible, htmlAttrs)
}
func (resource *Claim) T_CareTeamRole(numCareTeam int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return CodeableConceptSelect("careTeam["+strconv.Itoa(numCareTeam)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("careTeam["+strconv.Itoa(numCareTeam)+"].role", resource.CareTeam[numCareTeam].Role, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_CareTeamSpecialty(numCareTeam int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCareTeam >= len(resource.CareTeam) {
		return CodeableConceptSelect("careTeam["+strconv.Itoa(numCareTeam)+"].specialty", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("careTeam["+strconv.Itoa(numCareTeam)+"].specialty", resource.CareTeam[numCareTeam].Specialty, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoSequence(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return IntInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].sequence", &resource.SupportingInfo[numSupportingInfo].Sequence, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoCategory(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodeableConceptSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].category", &resource.SupportingInfo[numSupportingInfo].Category, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoCode(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodeableConceptSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].code", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].code", resource.SupportingInfo[numSupportingInfo].Code, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoTimingDate(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return FhirDateInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].timingDate", nil, htmlAttrs)
	}
	return FhirDateInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].timingDate", resource.SupportingInfo[numSupportingInfo].TimingDate, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoTimingPeriod(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return PeriodInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].timingPeriod", nil, htmlAttrs)
	}
	return PeriodInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].timingPeriod", resource.SupportingInfo[numSupportingInfo].TimingPeriod, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoValueBoolean(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return BoolInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueBoolean", resource.SupportingInfo[numSupportingInfo].ValueBoolean, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoValueString(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return StringInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueString", resource.SupportingInfo[numSupportingInfo].ValueString, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoValueQuantity(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return QuantityInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueQuantity", nil, htmlAttrs)
	}
	return QuantityInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueQuantity", resource.SupportingInfo[numSupportingInfo].ValueQuantity, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoValueAttachment(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return AttachmentInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueAttachment", resource.SupportingInfo[numSupportingInfo].ValueAttachment, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoValueReference(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return ReferenceInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueReference", nil, htmlAttrs)
	}
	return ReferenceInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueReference", resource.SupportingInfo[numSupportingInfo].ValueReference, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoValueIdentifier(numSupportingInfo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return IdentifierInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueIdentifier", nil, htmlAttrs)
	}
	return IdentifierInput("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].valueIdentifier", resource.SupportingInfo[numSupportingInfo].ValueIdentifier, htmlAttrs)
}
func (resource *Claim) T_SupportingInfoReason(numSupportingInfo int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSupportingInfo >= len(resource.SupportingInfo) {
		return CodeableConceptSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].reason", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("supportingInfo["+strconv.Itoa(numSupportingInfo)+"].reason", resource.SupportingInfo[numSupportingInfo].Reason, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_DiagnosisSequence(numDiagnosis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return IntInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].sequence", &resource.Diagnosis[numDiagnosis].Sequence, htmlAttrs)
}
func (resource *Claim) T_DiagnosisDiagnosisCodeableConcept(numDiagnosis int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].diagnosisCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].diagnosisCodeableConcept", &resource.Diagnosis[numDiagnosis].DiagnosisCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_DiagnosisDiagnosisReference(numDiagnosis int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return ReferenceInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].diagnosisReference", nil, htmlAttrs)
	}
	return ReferenceInput("diagnosis["+strconv.Itoa(numDiagnosis)+"].diagnosisReference", &resource.Diagnosis[numDiagnosis].DiagnosisReference, htmlAttrs)
}
func (resource *Claim) T_DiagnosisType(numDiagnosis int, numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) || numType >= len(resource.Diagnosis[numDiagnosis].Type) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].type["+strconv.Itoa(numType)+"]", &resource.Diagnosis[numDiagnosis].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_DiagnosisOnAdmission(numDiagnosis int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDiagnosis >= len(resource.Diagnosis) {
		return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].onAdmission", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("diagnosis["+strconv.Itoa(numDiagnosis)+"].onAdmission", resource.Diagnosis[numDiagnosis].OnAdmission, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ProcedureSequence(numProcedure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return IntInput("procedure["+strconv.Itoa(numProcedure)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("procedure["+strconv.Itoa(numProcedure)+"].sequence", &resource.Procedure[numProcedure].Sequence, htmlAttrs)
}
func (resource *Claim) T_ProcedureType(numProcedure int, numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) || numType >= len(resource.Procedure[numProcedure].Type) {
		return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].type["+strconv.Itoa(numType)+"]", &resource.Procedure[numProcedure].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ProcedureDate(numProcedure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return FhirDateTimeInput("procedure["+strconv.Itoa(numProcedure)+"].date", nil, htmlAttrs)
	}
	return FhirDateTimeInput("procedure["+strconv.Itoa(numProcedure)+"].date", resource.Procedure[numProcedure].Date, htmlAttrs)
}
func (resource *Claim) T_ProcedureProcedureCodeableConcept(numProcedure int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].procedureCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("procedure["+strconv.Itoa(numProcedure)+"].procedureCodeableConcept", &resource.Procedure[numProcedure].ProcedureCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ProcedureProcedureReference(numProcedure int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) {
		return ReferenceInput("procedure["+strconv.Itoa(numProcedure)+"].procedureReference", nil, htmlAttrs)
	}
	return ReferenceInput("procedure["+strconv.Itoa(numProcedure)+"].procedureReference", &resource.Procedure[numProcedure].ProcedureReference, htmlAttrs)
}
func (resource *Claim) T_ProcedureUdi(numProcedure int, numUdi int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProcedure >= len(resource.Procedure) || numUdi >= len(resource.Procedure[numProcedure].Udi) {
		return ReferenceInput("procedure["+strconv.Itoa(numProcedure)+"].udi["+strconv.Itoa(numUdi)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("procedure["+strconv.Itoa(numProcedure)+"].udi["+strconv.Itoa(numUdi)+"]", &resource.Procedure[numProcedure].Udi[numUdi], htmlAttrs)
}
func (resource *Claim) T_InsuranceSequence(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return IntInput("insurance["+strconv.Itoa(numInsurance)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("insurance["+strconv.Itoa(numInsurance)+"].sequence", &resource.Insurance[numInsurance].Sequence, htmlAttrs)
}
func (resource *Claim) T_InsuranceFocal(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].focal", nil, htmlAttrs)
	}
	return BoolInput("insurance["+strconv.Itoa(numInsurance)+"].focal", &resource.Insurance[numInsurance].Focal, htmlAttrs)
}
func (resource *Claim) T_InsuranceCoverage(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return ReferenceInput("insurance["+strconv.Itoa(numInsurance)+"].coverage", nil, htmlAttrs)
	}
	return ReferenceInput("insurance["+strconv.Itoa(numInsurance)+"].coverage", &resource.Insurance[numInsurance].Coverage, htmlAttrs)
}
func (resource *Claim) T_InsuranceBusinessArrangement(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return StringInput("insurance["+strconv.Itoa(numInsurance)+"].businessArrangement", nil, htmlAttrs)
	}
	return StringInput("insurance["+strconv.Itoa(numInsurance)+"].businessArrangement", resource.Insurance[numInsurance].BusinessArrangement, htmlAttrs)
}
func (resource *Claim) T_InsurancePreAuthRef(numInsurance int, numPreAuthRef int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) || numPreAuthRef >= len(resource.Insurance[numInsurance].PreAuthRef) {
		return StringInput("insurance["+strconv.Itoa(numInsurance)+"].preAuthRef["+strconv.Itoa(numPreAuthRef)+"]", nil, htmlAttrs)
	}
	return StringInput("insurance["+strconv.Itoa(numInsurance)+"].preAuthRef["+strconv.Itoa(numPreAuthRef)+"]", &resource.Insurance[numInsurance].PreAuthRef[numPreAuthRef], htmlAttrs)
}
func (resource *Claim) T_InsuranceClaimResponse(numInsurance int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numInsurance >= len(resource.Insurance) {
		return ReferenceInput("insurance["+strconv.Itoa(numInsurance)+"].claimResponse", nil, htmlAttrs)
	}
	return ReferenceInput("insurance["+strconv.Itoa(numInsurance)+"].claimResponse", resource.Insurance[numInsurance].ClaimResponse, htmlAttrs)
}
func (resource *Claim) T_AccidentDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateInput("accident.date", nil, htmlAttrs)
	}
	return FhirDateInput("accident.date", &resource.Accident.Date, htmlAttrs)
}
func (resource *Claim) T_AccidentType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("accident.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("accident.type", resource.Accident.Type, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_AccidentLocationAddress(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return AddressInput("accident.locationAddress", nil, htmlAttrs)
	}
	return AddressInput("accident.locationAddress", resource.Accident.LocationAddress, htmlAttrs)
}
func (resource *Claim) T_AccidentLocationReference(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("accident.locationReference", nil, htmlAttrs)
	}
	return ReferenceInput("accident.locationReference", resource.Accident.LocationReference, htmlAttrs)
}
func (resource *Claim) T_ItemSequence(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return IntInput("item["+strconv.Itoa(numItem)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].sequence", &resource.Item[numItem].Sequence, htmlAttrs)
}
func (resource *Claim) T_ItemTraceNumber(numItem int, numTraceNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numTraceNumber >= len(resource.Item[numItem].TraceNumber) {
		return IdentifierInput("item["+strconv.Itoa(numItem)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", nil, htmlAttrs)
	}
	return IdentifierInput("item["+strconv.Itoa(numItem)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", &resource.Item[numItem].TraceNumber[numTraceNumber], htmlAttrs)
}
func (resource *Claim) T_ItemCareTeamSequence(numItem int, numCareTeamSequence int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numCareTeamSequence >= len(resource.Item[numItem].CareTeamSequence) {
		return IntInput("item["+strconv.Itoa(numItem)+"].careTeamSequence["+strconv.Itoa(numCareTeamSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].careTeamSequence["+strconv.Itoa(numCareTeamSequence)+"]", &resource.Item[numItem].CareTeamSequence[numCareTeamSequence], htmlAttrs)
}
func (resource *Claim) T_ItemDiagnosisSequence(numItem int, numDiagnosisSequence int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDiagnosisSequence >= len(resource.Item[numItem].DiagnosisSequence) {
		return IntInput("item["+strconv.Itoa(numItem)+"].diagnosisSequence["+strconv.Itoa(numDiagnosisSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].diagnosisSequence["+strconv.Itoa(numDiagnosisSequence)+"]", &resource.Item[numItem].DiagnosisSequence[numDiagnosisSequence], htmlAttrs)
}
func (resource *Claim) T_ItemProcedureSequence(numItem int, numProcedureSequence int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numProcedureSequence >= len(resource.Item[numItem].ProcedureSequence) {
		return IntInput("item["+strconv.Itoa(numItem)+"].procedureSequence["+strconv.Itoa(numProcedureSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].procedureSequence["+strconv.Itoa(numProcedureSequence)+"]", &resource.Item[numItem].ProcedureSequence[numProcedureSequence], htmlAttrs)
}
func (resource *Claim) T_ItemInformationSequence(numItem int, numInformationSequence int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numInformationSequence >= len(resource.Item[numItem].InformationSequence) {
		return IntInput("item["+strconv.Itoa(numItem)+"].informationSequence["+strconv.Itoa(numInformationSequence)+"]", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].informationSequence["+strconv.Itoa(numInformationSequence)+"]", &resource.Item[numItem].InformationSequence[numInformationSequence], htmlAttrs)
}
func (resource *Claim) T_ItemRevenue(numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].revenue", resource.Item[numItem].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemCategory(numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].category", resource.Item[numItem].Category, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemProductOrService(numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].productOrService", resource.Item[numItem].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemProductOrServiceEnd(numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].productOrServiceEnd", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].productOrServiceEnd", resource.Item[numItem].ProductOrServiceEnd, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemRequest(numItem int, numRequest int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numRequest >= len(resource.Item[numItem].Request) {
		return ReferenceInput("item["+strconv.Itoa(numItem)+"].request["+strconv.Itoa(numRequest)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("item["+strconv.Itoa(numItem)+"].request["+strconv.Itoa(numRequest)+"]", &resource.Item[numItem].Request[numRequest], htmlAttrs)
}
func (resource *Claim) T_ItemModifier(numItem int, numModifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numModifier >= len(resource.Item[numItem].Modifier) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemProgramCode(numItem int, numProgramCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numProgramCode >= len(resource.Item[numItem].ProgramCode) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].programCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].programCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemServicedDate(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return FhirDateInput("item["+strconv.Itoa(numItem)+"].servicedDate", nil, htmlAttrs)
	}
	return FhirDateInput("item["+strconv.Itoa(numItem)+"].servicedDate", resource.Item[numItem].ServicedDate, htmlAttrs)
}
func (resource *Claim) T_ItemServicedPeriod(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return PeriodInput("item["+strconv.Itoa(numItem)+"].servicedPeriod", nil, htmlAttrs)
	}
	return PeriodInput("item["+strconv.Itoa(numItem)+"].servicedPeriod", resource.Item[numItem].ServicedPeriod, htmlAttrs)
}
func (resource *Claim) T_ItemLocationCodeableConcept(numItem int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].locationCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].locationCodeableConcept", resource.Item[numItem].LocationCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemLocationAddress(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return AddressInput("item["+strconv.Itoa(numItem)+"].locationAddress", nil, htmlAttrs)
	}
	return AddressInput("item["+strconv.Itoa(numItem)+"].locationAddress", resource.Item[numItem].LocationAddress, htmlAttrs)
}
func (resource *Claim) T_ItemLocationReference(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return ReferenceInput("item["+strconv.Itoa(numItem)+"].locationReference", nil, htmlAttrs)
	}
	return ReferenceInput("item["+strconv.Itoa(numItem)+"].locationReference", resource.Item[numItem].LocationReference, htmlAttrs)
}
func (resource *Claim) T_ItemPatientPaid(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].patientPaid", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].patientPaid", resource.Item[numItem].PatientPaid, htmlAttrs)
}
func (resource *Claim) T_ItemQuantity(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return QuantityInput("item["+strconv.Itoa(numItem)+"].quantity", nil, htmlAttrs)
	}
	return QuantityInput("item["+strconv.Itoa(numItem)+"].quantity", resource.Item[numItem].Quantity, htmlAttrs)
}
func (resource *Claim) T_ItemUnitPrice(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].unitPrice", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].unitPrice", resource.Item[numItem].UnitPrice, htmlAttrs)
}
func (resource *Claim) T_ItemFactor(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return Float64Input("item["+strconv.Itoa(numItem)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("item["+strconv.Itoa(numItem)+"].factor", resource.Item[numItem].Factor, htmlAttrs)
}
func (resource *Claim) T_ItemTax(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].tax", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].tax", resource.Item[numItem].Tax, htmlAttrs)
}
func (resource *Claim) T_ItemNet(numItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].net", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].net", resource.Item[numItem].Net, htmlAttrs)
}
func (resource *Claim) T_ItemUdi(numItem int, numUdi int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numUdi >= len(resource.Item[numItem].Udi) {
		return ReferenceInput("item["+strconv.Itoa(numItem)+"].udi["+strconv.Itoa(numUdi)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("item["+strconv.Itoa(numItem)+"].udi["+strconv.Itoa(numUdi)+"]", &resource.Item[numItem].Udi[numUdi], htmlAttrs)
}
func (resource *Claim) T_ItemEncounter(numItem int, numEncounter int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numEncounter >= len(resource.Item[numItem].Encounter) {
		return ReferenceInput("item["+strconv.Itoa(numItem)+"].encounter["+strconv.Itoa(numEncounter)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("item["+strconv.Itoa(numItem)+"].encounter["+strconv.Itoa(numEncounter)+"]", &resource.Item[numItem].Encounter[numEncounter], htmlAttrs)
}
func (resource *Claim) T_ItemBodySiteSite(numItem int, numBodySite int, numSite int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numBodySite >= len(resource.Item[numItem].BodySite) || numSite >= len(resource.Item[numItem].BodySite[numBodySite].Site) {
		return CodeableReferenceInput("item["+strconv.Itoa(numItem)+"].bodySite["+strconv.Itoa(numBodySite)+"].site["+strconv.Itoa(numSite)+"]", nil, htmlAttrs)
	}
	return CodeableReferenceInput("item["+strconv.Itoa(numItem)+"].bodySite["+strconv.Itoa(numBodySite)+"].site["+strconv.Itoa(numSite)+"]", &resource.Item[numItem].BodySite[numBodySite].Site[numSite], htmlAttrs)
}
func (resource *Claim) T_ItemBodySiteSubSite(numItem int, numBodySite int, numSubSite int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numBodySite >= len(resource.Item[numItem].BodySite) || numSubSite >= len(resource.Item[numItem].BodySite[numBodySite].SubSite) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].bodySite["+strconv.Itoa(numBodySite)+"].subSite["+strconv.Itoa(numSubSite)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].bodySite["+strconv.Itoa(numBodySite)+"].subSite["+strconv.Itoa(numSubSite)+"]", &resource.Item[numItem].BodySite[numBodySite].SubSite[numSubSite], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSequence(numItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].sequence", &resource.Item[numItem].Detail[numDetail].Sequence, htmlAttrs)
}
func (resource *Claim) T_ItemDetailTraceNumber(numItem int, numDetail int, numTraceNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numTraceNumber >= len(resource.Item[numItem].Detail[numDetail].TraceNumber) {
		return IdentifierInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", nil, htmlAttrs)
	}
	return IdentifierInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", &resource.Item[numItem].Detail[numDetail].TraceNumber[numTraceNumber], htmlAttrs)
}
func (resource *Claim) T_ItemDetailRevenue(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].revenue", resource.Item[numItem].Detail[numDetail].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailCategory(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].category", resource.Item[numItem].Detail[numDetail].Category, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailProductOrService(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrService", resource.Item[numItem].Detail[numDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailProductOrServiceEnd(numItem int, numDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrServiceEnd", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].productOrServiceEnd", resource.Item[numItem].Detail[numDetail].ProductOrServiceEnd, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailModifier(numItem int, numDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numModifier >= len(resource.Item[numItem].Detail[numDetail].Modifier) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Detail[numDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailProgramCode(numItem int, numDetail int, numProgramCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numProgramCode >= len(resource.Item[numItem].Detail[numDetail].ProgramCode) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].programCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].programCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].Detail[numDetail].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailPatientPaid(numItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].patientPaid", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].patientPaid", resource.Item[numItem].Detail[numDetail].PatientPaid, htmlAttrs)
}
func (resource *Claim) T_ItemDetailQuantity(numItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return QuantityInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].quantity", nil, htmlAttrs)
	}
	return QuantityInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].quantity", resource.Item[numItem].Detail[numDetail].Quantity, htmlAttrs)
}
func (resource *Claim) T_ItemDetailUnitPrice(numItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].unitPrice", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].unitPrice", resource.Item[numItem].Detail[numDetail].UnitPrice, htmlAttrs)
}
func (resource *Claim) T_ItemDetailFactor(numItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return Float64Input("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].factor", resource.Item[numItem].Detail[numDetail].Factor, htmlAttrs)
}
func (resource *Claim) T_ItemDetailTax(numItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].tax", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].tax", resource.Item[numItem].Detail[numDetail].Tax, htmlAttrs)
}
func (resource *Claim) T_ItemDetailNet(numItem int, numDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].net", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].net", resource.Item[numItem].Detail[numDetail].Net, htmlAttrs)
}
func (resource *Claim) T_ItemDetailUdi(numItem int, numDetail int, numUdi int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numUdi >= len(resource.Item[numItem].Detail[numDetail].Udi) {
		return ReferenceInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].udi["+strconv.Itoa(numUdi)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].udi["+strconv.Itoa(numUdi)+"]", &resource.Item[numItem].Detail[numDetail].Udi[numUdi], htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailSequence(numItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].sequence", nil, htmlAttrs)
	}
	return IntInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].sequence", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Sequence, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailTraceNumber(numItem int, numDetail int, numSubDetail int, numTraceNumber int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numTraceNumber >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].TraceNumber) {
		return IdentifierInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", nil, htmlAttrs)
	}
	return IdentifierInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].traceNumber["+strconv.Itoa(numTraceNumber)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].TraceNumber[numTraceNumber], htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailRevenue(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].revenue", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].revenue", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Revenue, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailCategory(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].category", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Category, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailProductOrService(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrService", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrService", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailProductOrServiceEnd(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrServiceEnd", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].productOrServiceEnd", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrServiceEnd, optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailModifier(numItem int, numDetail int, numSubDetail int, numModifier int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numModifier >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Modifier) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].modifier["+strconv.Itoa(numModifier)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[numModifier], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailProgramCode(numItem int, numDetail int, numSubDetail int, numProgramCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numProgramCode >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProgramCode) {
		return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].programCode["+strconv.Itoa(numProgramCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].programCode["+strconv.Itoa(numProgramCode)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProgramCode[numProgramCode], optionsValueSet, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailPatientPaid(numItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].patientPaid", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].patientPaid", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].PatientPaid, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailQuantity(numItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return QuantityInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].quantity", nil, htmlAttrs)
	}
	return QuantityInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].quantity", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Quantity, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailUnitPrice(numItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].unitPrice", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].unitPrice", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].UnitPrice, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailFactor(numItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return Float64Input("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].factor", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Factor, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailTax(numItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].tax", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].tax", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Tax, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailNet(numItem int, numDetail int, numSubDetail int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) {
		return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].net", nil, htmlAttrs)
	}
	return MoneyInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].net", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Net, htmlAttrs)
}
func (resource *Claim) T_ItemDetailSubDetailUdi(numItem int, numDetail int, numSubDetail int, numUdi int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numItem >= len(resource.Item) || numDetail >= len(resource.Item[numItem].Detail) || numSubDetail >= len(resource.Item[numItem].Detail[numDetail].SubDetail) || numUdi >= len(resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Udi) {
		return ReferenceInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].udi["+strconv.Itoa(numUdi)+"]", nil, htmlAttrs)
	}
	return ReferenceInput("item["+strconv.Itoa(numItem)+"].detail["+strconv.Itoa(numDetail)+"].subDetail["+strconv.Itoa(numSubDetail)+"].udi["+strconv.Itoa(numUdi)+"]", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Udi[numUdi], htmlAttrs)
}
