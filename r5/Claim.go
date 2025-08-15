//generated August 15 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

package r5

import "encoding/json"

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
	TimingDate        *string          `json:"timingDate"`
	TimingPeriod      *Period          `json:"timingPeriod"`
	ValueBoolean      *bool            `json:"valueBoolean"`
	ValueString       *string          `json:"valueString"`
	ValueQuantity     *Quantity        `json:"valueQuantity"`
	ValueAttachment   *Attachment      `json:"valueAttachment"`
	ValueReference    *Reference       `json:"valueReference"`
	ValueIdentifier   *Identifier      `json:"valueIdentifier"`
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
	LocationAddress   *Address         `json:"locationAddress"`
	LocationReference *Reference       `json:"locationReference"`
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
	ServicedDate            *string             `json:"servicedDate"`
	ServicedPeriod          *Period             `json:"servicedPeriod"`
	LocationCodeableConcept *CodeableConcept    `json:"locationCodeableConcept"`
	LocationAddress         *Address            `json:"locationAddress"`
	LocationReference       *Reference          `json:"locationReference"`
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
