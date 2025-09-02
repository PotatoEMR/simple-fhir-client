package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	Created              string                `json:"created"`
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
	TimingDate        *string          `json:"timingDate"`
	TimingPeriod      *Period          `json:"timingPeriod"`
	ValueBoolean      *bool            `json:"valueBoolean"`
	ValueString       *string          `json:"valueString"`
	ValueQuantity     *Quantity        `json:"valueQuantity"`
	ValueAttachment   *Attachment      `json:"valueAttachment"`
	ValueReference    *Reference       `json:"valueReference"`
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
	Date                     *string           `json:"date,omitempty"`
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
	Date              string           `json:"date"`
	Type              *CodeableConcept `json:"type,omitempty"`
	LocationAddress   *Address         `json:"locationAddress"`
	LocationReference *Reference       `json:"locationReference"`
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
	ServicedDate            *string           `json:"servicedDate"`
	ServicedPeriod          *Period           `json:"servicedPeriod"`
	LocationCodeableConcept *CodeableConcept  `json:"locationCodeableConcept"`
	LocationAddress         *Address          `json:"locationAddress"`
	LocationReference       *Reference        `json:"locationReference"`
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

func (resource *Claim) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Claim) T_Status() templ.Component {
	optionsValueSet := VSFm_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", &resource.Status, optionsValueSet)
}
func (resource *Claim) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Type, optionsValueSet)
}
func (resource *Claim) T_SubType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("subType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subType", resource.SubType, optionsValueSet)
}
func (resource *Claim) T_Use() templ.Component {
	optionsValueSet := VSClaim_use

	if resource == nil {
		return CodeSelect("use", nil, optionsValueSet)
	}
	return CodeSelect("use", &resource.Use, optionsValueSet)
}
func (resource *Claim) T_Priority(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("priority", nil, optionsValueSet)
	}
	return CodeableConceptSelect("priority", &resource.Priority, optionsValueSet)
}
func (resource *Claim) T_FundsReserve(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("fundsReserve", nil, optionsValueSet)
	}
	return CodeableConceptSelect("fundsReserve", resource.FundsReserve, optionsValueSet)
}
func (resource *Claim) T_RelatedRelationship(numRelated int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Related) >= numRelated {
		return CodeableConceptSelect("relationship", nil, optionsValueSet)
	}
	return CodeableConceptSelect("relationship", resource.Related[numRelated].Relationship, optionsValueSet)
}
func (resource *Claim) T_PayeeType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Payee.Type, optionsValueSet)
}
func (resource *Claim) T_CareTeamRole(numCareTeam int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CareTeam) >= numCareTeam {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", resource.CareTeam[numCareTeam].Role, optionsValueSet)
}
func (resource *Claim) T_CareTeamQualification(numCareTeam int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.CareTeam) >= numCareTeam {
		return CodeableConceptSelect("qualification", nil, optionsValueSet)
	}
	return CodeableConceptSelect("qualification", resource.CareTeam[numCareTeam].Qualification, optionsValueSet)
}
func (resource *Claim) T_SupportingInfoCategory(numSupportingInfo int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.SupportingInfo) >= numSupportingInfo {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", &resource.SupportingInfo[numSupportingInfo].Category, optionsValueSet)
}
func (resource *Claim) T_SupportingInfoCode(numSupportingInfo int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.SupportingInfo) >= numSupportingInfo {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", resource.SupportingInfo[numSupportingInfo].Code, optionsValueSet)
}
func (resource *Claim) T_SupportingInfoReason(numSupportingInfo int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.SupportingInfo) >= numSupportingInfo {
		return CodeableConceptSelect("reason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reason", resource.SupportingInfo[numSupportingInfo].Reason, optionsValueSet)
}
func (resource *Claim) T_DiagnosisType(numDiagnosis int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Diagnosis) >= numDiagnosis {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Diagnosis[numDiagnosis].Type[0], optionsValueSet)
}
func (resource *Claim) T_DiagnosisOnAdmission(numDiagnosis int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Diagnosis) >= numDiagnosis {
		return CodeableConceptSelect("onAdmission", nil, optionsValueSet)
	}
	return CodeableConceptSelect("onAdmission", resource.Diagnosis[numDiagnosis].OnAdmission, optionsValueSet)
}
func (resource *Claim) T_DiagnosisPackageCode(numDiagnosis int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Diagnosis) >= numDiagnosis {
		return CodeableConceptSelect("packageCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("packageCode", resource.Diagnosis[numDiagnosis].PackageCode, optionsValueSet)
}
func (resource *Claim) T_ProcedureType(numProcedure int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Procedure) >= numProcedure {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Procedure[numProcedure].Type[0], optionsValueSet)
}
func (resource *Claim) T_AccidentType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Accident.Type, optionsValueSet)
}
func (resource *Claim) T_ItemRevenue(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("revenue", resource.Item[numItem].Revenue, optionsValueSet)
}
func (resource *Claim) T_ItemCategory(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.Item[numItem].Category, optionsValueSet)
}
func (resource *Claim) T_ItemProductOrService(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", &resource.Item[numItem].ProductOrService, optionsValueSet)
}
func (resource *Claim) T_ItemModifier(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.Item[numItem].Modifier[0], optionsValueSet)
}
func (resource *Claim) T_ItemProgramCode(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("programCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("programCode", &resource.Item[numItem].ProgramCode[0], optionsValueSet)
}
func (resource *Claim) T_ItemBodySite(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("bodySite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("bodySite", resource.Item[numItem].BodySite, optionsValueSet)
}
func (resource *Claim) T_ItemSubSite(numItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item) >= numItem {
		return CodeableConceptSelect("subSite", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subSite", &resource.Item[numItem].SubSite[0], optionsValueSet)
}
func (resource *Claim) T_ItemDetailRevenue(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("revenue", resource.Item[numItem].Detail[numDetail].Revenue, optionsValueSet)
}
func (resource *Claim) T_ItemDetailCategory(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.Item[numItem].Detail[numDetail].Category, optionsValueSet)
}
func (resource *Claim) T_ItemDetailProductOrService(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", &resource.Item[numItem].Detail[numDetail].ProductOrService, optionsValueSet)
}
func (resource *Claim) T_ItemDetailModifier(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.Item[numItem].Detail[numDetail].Modifier[0], optionsValueSet)
}
func (resource *Claim) T_ItemDetailProgramCode(numItem int, numDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail) >= numDetail {
		return CodeableConceptSelect("programCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("programCode", &resource.Item[numItem].Detail[numDetail].ProgramCode[0], optionsValueSet)
}
func (resource *Claim) T_ItemDetailSubDetailRevenue(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("revenue", nil, optionsValueSet)
	}
	return CodeableConceptSelect("revenue", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Revenue, optionsValueSet)
}
func (resource *Claim) T_ItemDetailSubDetailCategory(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("category", resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Category, optionsValueSet)
}
func (resource *Claim) T_ItemDetailSubDetailProductOrService(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("productOrService", nil, optionsValueSet)
	}
	return CodeableConceptSelect("productOrService", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProductOrService, optionsValueSet)
}
func (resource *Claim) T_ItemDetailSubDetailModifier(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("modifier", nil, optionsValueSet)
	}
	return CodeableConceptSelect("modifier", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].Modifier[0], optionsValueSet)
}
func (resource *Claim) T_ItemDetailSubDetailProgramCode(numItem int, numDetail int, numSubDetail int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Item[numItem].Detail[numDetail].SubDetail) >= numSubDetail {
		return CodeableConceptSelect("programCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("programCode", &resource.Item[numItem].Detail[numDetail].SubDetail[numSubDetail].ProgramCode[0], optionsValueSet)
}
