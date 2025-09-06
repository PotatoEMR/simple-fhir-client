package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/Contract
type Contract struct {
	Id                       *string                    `json:"id,omitempty"`
	Meta                     *Meta                      `json:"meta,omitempty"`
	ImplicitRules            *string                    `json:"implicitRules,omitempty"`
	Language                 *string                    `json:"language,omitempty"`
	Text                     *Narrative                 `json:"text,omitempty"`
	Contained                []Resource                 `json:"contained,omitempty"`
	Extension                []Extension                `json:"extension,omitempty"`
	ModifierExtension        []Extension                `json:"modifierExtension,omitempty"`
	Identifier               []Identifier               `json:"identifier,omitempty"`
	Url                      *string                    `json:"url,omitempty"`
	Version                  *string                    `json:"version,omitempty"`
	Status                   *string                    `json:"status,omitempty"`
	LegalState               *CodeableConcept           `json:"legalState,omitempty"`
	InstantiatesCanonical    *Reference                 `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri          *string                    `json:"instantiatesUri,omitempty"`
	ContentDerivative        *CodeableConcept           `json:"contentDerivative,omitempty"`
	Issued                   *string                    `json:"issued,omitempty"`
	Applies                  *Period                    `json:"applies,omitempty"`
	ExpirationType           *CodeableConcept           `json:"expirationType,omitempty"`
	Subject                  []Reference                `json:"subject,omitempty"`
	Authority                []Reference                `json:"authority,omitempty"`
	Domain                   []Reference                `json:"domain,omitempty"`
	Site                     []Reference                `json:"site,omitempty"`
	Name                     *string                    `json:"name,omitempty"`
	Title                    *string                    `json:"title,omitempty"`
	Subtitle                 *string                    `json:"subtitle,omitempty"`
	Alias                    []string                   `json:"alias,omitempty"`
	Author                   *Reference                 `json:"author,omitempty"`
	Scope                    *CodeableConcept           `json:"scope,omitempty"`
	TopicCodeableConcept     *CodeableConcept           `json:"topicCodeableConcept,omitempty"`
	TopicReference           *Reference                 `json:"topicReference,omitempty"`
	Type                     *CodeableConcept           `json:"type,omitempty"`
	SubType                  []CodeableConcept          `json:"subType,omitempty"`
	ContentDefinition        *ContractContentDefinition `json:"contentDefinition,omitempty"`
	Term                     []ContractTerm             `json:"term,omitempty"`
	SupportingInfo           []Reference                `json:"supportingInfo,omitempty"`
	RelevantHistory          []Reference                `json:"relevantHistory,omitempty"`
	Signer                   []ContractSigner           `json:"signer,omitempty"`
	Friendly                 []ContractFriendly         `json:"friendly,omitempty"`
	Legal                    []ContractLegal            `json:"legal,omitempty"`
	Rule                     []ContractRule             `json:"rule,omitempty"`
	LegallyBindingAttachment *Attachment                `json:"legallyBindingAttachment,omitempty"`
	LegallyBindingReference  *Reference                 `json:"legallyBindingReference,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Contract
type ContractContentDefinition struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              CodeableConcept  `json:"type"`
	SubType           *CodeableConcept `json:"subType,omitempty"`
	Publisher         *Reference       `json:"publisher,omitempty"`
	PublicationDate   *string          `json:"publicationDate,omitempty"`
	PublicationStatus string           `json:"publicationStatus"`
	Copyright         *string          `json:"copyright,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Contract
type ContractTerm struct {
	Id                   *string                     `json:"id,omitempty"`
	Extension            []Extension                 `json:"extension,omitempty"`
	ModifierExtension    []Extension                 `json:"modifierExtension,omitempty"`
	Identifier           *Identifier                 `json:"identifier,omitempty"`
	Issued               *string                     `json:"issued,omitempty"`
	Applies              *Period                     `json:"applies,omitempty"`
	TopicCodeableConcept *CodeableConcept            `json:"topicCodeableConcept,omitempty"`
	TopicReference       *Reference                  `json:"topicReference,omitempty"`
	Type                 *CodeableConcept            `json:"type,omitempty"`
	SubType              *CodeableConcept            `json:"subType,omitempty"`
	Text                 *string                     `json:"text,omitempty"`
	SecurityLabel        []ContractTermSecurityLabel `json:"securityLabel,omitempty"`
	Offer                ContractTermOffer           `json:"offer"`
	Asset                []ContractTermAsset         `json:"asset,omitempty"`
	Action               []ContractTermAction        `json:"action,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Contract
type ContractTermSecurityLabel struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Number            []int       `json:"number,omitempty"`
	Classification    Coding      `json:"classification"`
	Category          []Coding    `json:"category,omitempty"`
	Control           []Coding    `json:"control,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Contract
type ContractTermOffer struct {
	Id                  *string                   `json:"id,omitempty"`
	Extension           []Extension               `json:"extension,omitempty"`
	ModifierExtension   []Extension               `json:"modifierExtension,omitempty"`
	Identifier          []Identifier              `json:"identifier,omitempty"`
	Party               []ContractTermOfferParty  `json:"party,omitempty"`
	Topic               *Reference                `json:"topic,omitempty"`
	Type                *CodeableConcept          `json:"type,omitempty"`
	Decision            *CodeableConcept          `json:"decision,omitempty"`
	DecisionMode        []CodeableConcept         `json:"decisionMode,omitempty"`
	Answer              []ContractTermOfferAnswer `json:"answer,omitempty"`
	Text                *string                   `json:"text,omitempty"`
	LinkId              []string                  `json:"linkId,omitempty"`
	SecurityLabelNumber []int                     `json:"securityLabelNumber,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Contract
type ContractTermOfferParty struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Reference         []Reference     `json:"reference"`
	Role              CodeableConcept `json:"role"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Contract
type ContractTermOfferAnswer struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ValueBoolean      bool        `json:"valueBoolean"`
	ValueDecimal      float64     `json:"valueDecimal"`
	ValueInteger      int         `json:"valueInteger"`
	ValueDate         string      `json:"valueDate"`
	ValueDateTime     string      `json:"valueDateTime"`
	ValueTime         string      `json:"valueTime"`
	ValueString       string      `json:"valueString"`
	ValueUri          string      `json:"valueUri"`
	ValueAttachment   Attachment  `json:"valueAttachment"`
	ValueCoding       Coding      `json:"valueCoding"`
	ValueQuantity     Quantity    `json:"valueQuantity"`
	ValueReference    Reference   `json:"valueReference"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Contract
type ContractTermAsset struct {
	Id                  *string                       `json:"id,omitempty"`
	Extension           []Extension                   `json:"extension,omitempty"`
	ModifierExtension   []Extension                   `json:"modifierExtension,omitempty"`
	Scope               *CodeableConcept              `json:"scope,omitempty"`
	Type                []CodeableConcept             `json:"type,omitempty"`
	TypeReference       []Reference                   `json:"typeReference,omitempty"`
	Subtype             []CodeableConcept             `json:"subtype,omitempty"`
	Relationship        *Coding                       `json:"relationship,omitempty"`
	Context             []ContractTermAssetContext    `json:"context,omitempty"`
	Condition           *string                       `json:"condition,omitempty"`
	PeriodType          []CodeableConcept             `json:"periodType,omitempty"`
	Period              []Period                      `json:"period,omitempty"`
	UsePeriod           []Period                      `json:"usePeriod,omitempty"`
	Text                *string                       `json:"text,omitempty"`
	LinkId              []string                      `json:"linkId,omitempty"`
	SecurityLabelNumber []int                         `json:"securityLabelNumber,omitempty"`
	ValuedItem          []ContractTermAssetValuedItem `json:"valuedItem,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Contract
type ContractTermAssetContext struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Reference         *Reference        `json:"reference,omitempty"`
	Code              []CodeableConcept `json:"code,omitempty"`
	Text              *string           `json:"text,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Contract
type ContractTermAssetValuedItem struct {
	Id                    *string          `json:"id,omitempty"`
	Extension             []Extension      `json:"extension,omitempty"`
	ModifierExtension     []Extension      `json:"modifierExtension,omitempty"`
	EntityCodeableConcept *CodeableConcept `json:"entityCodeableConcept,omitempty"`
	EntityReference       *Reference       `json:"entityReference,omitempty"`
	Identifier            *Identifier      `json:"identifier,omitempty"`
	EffectiveTime         *string          `json:"effectiveTime,omitempty"`
	Quantity              *Quantity        `json:"quantity,omitempty"`
	UnitPrice             *Money           `json:"unitPrice,omitempty"`
	Factor                *float64         `json:"factor,omitempty"`
	Points                *float64         `json:"points,omitempty"`
	Net                   *Money           `json:"net,omitempty"`
	Payment               *string          `json:"payment,omitempty"`
	PaymentDate           *string          `json:"paymentDate,omitempty"`
	Responsible           *Reference       `json:"responsible,omitempty"`
	Recipient             *Reference       `json:"recipient,omitempty"`
	LinkId                []string         `json:"linkId,omitempty"`
	SecurityLabelNumber   []int            `json:"securityLabelNumber,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Contract
type ContractTermAction struct {
	Id                  *string                     `json:"id,omitempty"`
	Extension           []Extension                 `json:"extension,omitempty"`
	ModifierExtension   []Extension                 `json:"modifierExtension,omitempty"`
	DoNotPerform        *bool                       `json:"doNotPerform,omitempty"`
	Type                CodeableConcept             `json:"type"`
	Subject             []ContractTermActionSubject `json:"subject,omitempty"`
	Intent              CodeableConcept             `json:"intent"`
	LinkId              []string                    `json:"linkId,omitempty"`
	Status              CodeableConcept             `json:"status"`
	Context             *Reference                  `json:"context,omitempty"`
	ContextLinkId       []string                    `json:"contextLinkId,omitempty"`
	OccurrenceDateTime  *string                     `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod    *Period                     `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming    *Timing                     `json:"occurrenceTiming,omitempty"`
	Requester           []Reference                 `json:"requester,omitempty"`
	RequesterLinkId     []string                    `json:"requesterLinkId,omitempty"`
	PerformerType       []CodeableConcept           `json:"performerType,omitempty"`
	PerformerRole       *CodeableConcept            `json:"performerRole,omitempty"`
	Performer           *Reference                  `json:"performer,omitempty"`
	PerformerLinkId     []string                    `json:"performerLinkId,omitempty"`
	ReasonCode          []CodeableConcept           `json:"reasonCode,omitempty"`
	ReasonReference     []Reference                 `json:"reasonReference,omitempty"`
	Reason              []string                    `json:"reason,omitempty"`
	ReasonLinkId        []string                    `json:"reasonLinkId,omitempty"`
	Note                []Annotation                `json:"note,omitempty"`
	SecurityLabelNumber []int                       `json:"securityLabelNumber,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Contract
type ContractTermActionSubject struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Reference         []Reference      `json:"reference"`
	Role              *CodeableConcept `json:"role,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Contract
type ContractSigner struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              Coding      `json:"type"`
	Party             Reference   `json:"party"`
	Signature         []Signature `json:"signature"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Contract
type ContractFriendly struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ContentAttachment Attachment  `json:"contentAttachment"`
	ContentReference  Reference   `json:"contentReference"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Contract
type ContractLegal struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ContentAttachment Attachment  `json:"contentAttachment"`
	ContentReference  Reference   `json:"contentReference"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Contract
type ContractRule struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ContentAttachment Attachment  `json:"contentAttachment"`
	ContentReference  Reference   `json:"contentReference"`
}

type OtherContract Contract

// on convert struct to json, automatically add resourceType=Contract
func (r Contract) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherContract
		ResourceType string `json:"resourceType"`
	}{
		OtherContract: OtherContract(r),
		ResourceType:  "Contract",
	})
}

func (resource *Contract) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Contract.Id", nil)
	}
	return StringInput("Contract.Id", resource.Id)
}
func (resource *Contract) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Contract.ImplicitRules", nil)
	}
	return StringInput("Contract.ImplicitRules", resource.ImplicitRules)
}
func (resource *Contract) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Contract.Language", nil, optionsValueSet)
	}
	return CodeSelect("Contract.Language", resource.Language, optionsValueSet)
}
func (resource *Contract) T_Url() templ.Component {

	if resource == nil {
		return StringInput("Contract.Url", nil)
	}
	return StringInput("Contract.Url", resource.Url)
}
func (resource *Contract) T_Version() templ.Component {

	if resource == nil {
		return StringInput("Contract.Version", nil)
	}
	return StringInput("Contract.Version", resource.Version)
}
func (resource *Contract) T_Status() templ.Component {
	optionsValueSet := VSContract_status

	if resource == nil {
		return CodeSelect("Contract.Status", nil, optionsValueSet)
	}
	return CodeSelect("Contract.Status", resource.Status, optionsValueSet)
}
func (resource *Contract) T_LegalState(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Contract.LegalState", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.LegalState", resource.LegalState, optionsValueSet)
}
func (resource *Contract) T_InstantiatesUri() templ.Component {

	if resource == nil {
		return StringInput("Contract.InstantiatesUri", nil)
	}
	return StringInput("Contract.InstantiatesUri", resource.InstantiatesUri)
}
func (resource *Contract) T_ContentDerivative(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Contract.ContentDerivative", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.ContentDerivative", resource.ContentDerivative, optionsValueSet)
}
func (resource *Contract) T_Issued() templ.Component {

	if resource == nil {
		return StringInput("Contract.Issued", nil)
	}
	return StringInput("Contract.Issued", resource.Issued)
}
func (resource *Contract) T_ExpirationType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Contract.ExpirationType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.ExpirationType", resource.ExpirationType, optionsValueSet)
}
func (resource *Contract) T_Name() templ.Component {

	if resource == nil {
		return StringInput("Contract.Name", nil)
	}
	return StringInput("Contract.Name", resource.Name)
}
func (resource *Contract) T_Title() templ.Component {

	if resource == nil {
		return StringInput("Contract.Title", nil)
	}
	return StringInput("Contract.Title", resource.Title)
}
func (resource *Contract) T_Subtitle() templ.Component {

	if resource == nil {
		return StringInput("Contract.Subtitle", nil)
	}
	return StringInput("Contract.Subtitle", resource.Subtitle)
}
func (resource *Contract) T_Alias(numAlias int) templ.Component {

	if resource == nil || len(resource.Alias) >= numAlias {
		return StringInput("Contract.Alias["+strconv.Itoa(numAlias)+"]", nil)
	}
	return StringInput("Contract.Alias["+strconv.Itoa(numAlias)+"]", &resource.Alias[numAlias])
}
func (resource *Contract) T_Scope(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Contract.Scope", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Scope", resource.Scope, optionsValueSet)
}
func (resource *Contract) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Contract.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Type", resource.Type, optionsValueSet)
}
func (resource *Contract) T_SubType(numSubType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.SubType) >= numSubType {
		return CodeableConceptSelect("Contract.SubType["+strconv.Itoa(numSubType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.SubType["+strconv.Itoa(numSubType)+"]", &resource.SubType[numSubType], optionsValueSet)
}
func (resource *Contract) T_ContentDefinitionId() templ.Component {

	if resource == nil {
		return StringInput("Contract.ContentDefinition.Id", nil)
	}
	return StringInput("Contract.ContentDefinition.Id", resource.ContentDefinition.Id)
}
func (resource *Contract) T_ContentDefinitionType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Contract.ContentDefinition.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.ContentDefinition.Type", &resource.ContentDefinition.Type, optionsValueSet)
}
func (resource *Contract) T_ContentDefinitionSubType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Contract.ContentDefinition.SubType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.ContentDefinition.SubType", resource.ContentDefinition.SubType, optionsValueSet)
}
func (resource *Contract) T_ContentDefinitionPublicationDate() templ.Component {

	if resource == nil {
		return StringInput("Contract.ContentDefinition.PublicationDate", nil)
	}
	return StringInput("Contract.ContentDefinition.PublicationDate", resource.ContentDefinition.PublicationDate)
}
func (resource *Contract) T_ContentDefinitionPublicationStatus() templ.Component {
	optionsValueSet := VSContract_publicationstatus

	if resource == nil {
		return CodeSelect("Contract.ContentDefinition.PublicationStatus", nil, optionsValueSet)
	}
	return CodeSelect("Contract.ContentDefinition.PublicationStatus", &resource.ContentDefinition.PublicationStatus, optionsValueSet)
}
func (resource *Contract) T_ContentDefinitionCopyright() templ.Component {

	if resource == nil {
		return StringInput("Contract.ContentDefinition.Copyright", nil)
	}
	return StringInput("Contract.ContentDefinition.Copyright", resource.ContentDefinition.Copyright)
}
func (resource *Contract) T_TermId(numTerm int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Id", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Id", resource.Term[numTerm].Id)
}
func (resource *Contract) T_TermIssued(numTerm int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Issued", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Issued", resource.Term[numTerm].Issued)
}
func (resource *Contract) T_TermType(numTerm int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Type", resource.Term[numTerm].Type, optionsValueSet)
}
func (resource *Contract) T_TermSubType(numTerm int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].SubType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].SubType", resource.Term[numTerm].SubType, optionsValueSet)
}
func (resource *Contract) T_TermText(numTerm int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Text", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Text", resource.Term[numTerm].Text)
}
func (resource *Contract) T_TermSecurityLabelId(numTerm int, numSecurityLabel int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].SecurityLabel) >= numSecurityLabel {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"].Id", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"].Id", resource.Term[numTerm].SecurityLabel[numSecurityLabel].Id)
}
func (resource *Contract) T_TermSecurityLabelNumber(numTerm int, numSecurityLabel int, numNumber int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].SecurityLabel) >= numSecurityLabel || len(resource.Term[numTerm].SecurityLabel[numSecurityLabel].Number) >= numNumber {
		return IntInput("Contract.Term["+strconv.Itoa(numTerm)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"].Number["+strconv.Itoa(numNumber)+"]", nil)
	}
	return IntInput("Contract.Term["+strconv.Itoa(numTerm)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"].Number["+strconv.Itoa(numNumber)+"]", &resource.Term[numTerm].SecurityLabel[numSecurityLabel].Number[numNumber])
}
func (resource *Contract) T_TermSecurityLabelClassification(numTerm int, numSecurityLabel int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].SecurityLabel) >= numSecurityLabel {
		return CodingSelect("Contract.Term["+strconv.Itoa(numTerm)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"].Classification", nil, optionsValueSet)
	}
	return CodingSelect("Contract.Term["+strconv.Itoa(numTerm)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"].Classification", &resource.Term[numTerm].SecurityLabel[numSecurityLabel].Classification, optionsValueSet)
}
func (resource *Contract) T_TermSecurityLabelCategory(numTerm int, numSecurityLabel int, numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].SecurityLabel) >= numSecurityLabel || len(resource.Term[numTerm].SecurityLabel[numSecurityLabel].Category) >= numCategory {
		return CodingSelect("Contract.Term["+strconv.Itoa(numTerm)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"].Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodingSelect("Contract.Term["+strconv.Itoa(numTerm)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"].Category["+strconv.Itoa(numCategory)+"]", &resource.Term[numTerm].SecurityLabel[numSecurityLabel].Category[numCategory], optionsValueSet)
}
func (resource *Contract) T_TermSecurityLabelControl(numTerm int, numSecurityLabel int, numControl int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].SecurityLabel) >= numSecurityLabel || len(resource.Term[numTerm].SecurityLabel[numSecurityLabel].Control) >= numControl {
		return CodingSelect("Contract.Term["+strconv.Itoa(numTerm)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"].Control["+strconv.Itoa(numControl)+"]", nil, optionsValueSet)
	}
	return CodingSelect("Contract.Term["+strconv.Itoa(numTerm)+"].SecurityLabel["+strconv.Itoa(numSecurityLabel)+"].Control["+strconv.Itoa(numControl)+"]", &resource.Term[numTerm].SecurityLabel[numSecurityLabel].Control[numControl], optionsValueSet)
}
func (resource *Contract) T_TermOfferId(numTerm int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.Id", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.Id", resource.Term[numTerm].Offer.Id)
}
func (resource *Contract) T_TermOfferType(numTerm int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.Type", resource.Term[numTerm].Offer.Type, optionsValueSet)
}
func (resource *Contract) T_TermOfferDecision(numTerm int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.Decision", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.Decision", resource.Term[numTerm].Offer.Decision, optionsValueSet)
}
func (resource *Contract) T_TermOfferDecisionMode(numTerm int, numDecisionMode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Offer.DecisionMode) >= numDecisionMode {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.DecisionMode["+strconv.Itoa(numDecisionMode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.DecisionMode["+strconv.Itoa(numDecisionMode)+"]", &resource.Term[numTerm].Offer.DecisionMode[numDecisionMode], optionsValueSet)
}
func (resource *Contract) T_TermOfferText(numTerm int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.Text", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.Text", resource.Term[numTerm].Offer.Text)
}
func (resource *Contract) T_TermOfferLinkId(numTerm int, numLinkId int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Offer.LinkId) >= numLinkId {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.LinkId["+strconv.Itoa(numLinkId)+"]", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.LinkId["+strconv.Itoa(numLinkId)+"]", &resource.Term[numTerm].Offer.LinkId[numLinkId])
}
func (resource *Contract) T_TermOfferSecurityLabelNumber(numTerm int, numSecurityLabelNumber int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Offer.SecurityLabelNumber) >= numSecurityLabelNumber {
		return IntInput("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.SecurityLabelNumber["+strconv.Itoa(numSecurityLabelNumber)+"]", nil)
	}
	return IntInput("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.SecurityLabelNumber["+strconv.Itoa(numSecurityLabelNumber)+"]", &resource.Term[numTerm].Offer.SecurityLabelNumber[numSecurityLabelNumber])
}
func (resource *Contract) T_TermOfferPartyId(numTerm int, numParty int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Offer.Party) >= numParty {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.Party["+strconv.Itoa(numParty)+"].Id", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.Party["+strconv.Itoa(numParty)+"].Id", resource.Term[numTerm].Offer.Party[numParty].Id)
}
func (resource *Contract) T_TermOfferPartyRole(numTerm int, numParty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Offer.Party) >= numParty {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.Party["+strconv.Itoa(numParty)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.Party["+strconv.Itoa(numParty)+"].Role", &resource.Term[numTerm].Offer.Party[numParty].Role, optionsValueSet)
}
func (resource *Contract) T_TermOfferAnswerId(numTerm int, numAnswer int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Offer.Answer) >= numAnswer {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.Answer["+strconv.Itoa(numAnswer)+"].Id", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Offer.Answer["+strconv.Itoa(numAnswer)+"].Id", resource.Term[numTerm].Offer.Answer[numAnswer].Id)
}
func (resource *Contract) T_TermAssetId(numTerm int, numAsset int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Id", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Id", resource.Term[numTerm].Asset[numAsset].Id)
}
func (resource *Contract) T_TermAssetScope(numTerm int, numAsset int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Scope", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Scope", resource.Term[numTerm].Asset[numAsset].Scope, optionsValueSet)
}
func (resource *Contract) T_TermAssetType(numTerm int, numAsset int, numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset || len(resource.Term[numTerm].Asset[numAsset].Type) >= numType {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Type["+strconv.Itoa(numType)+"]", &resource.Term[numTerm].Asset[numAsset].Type[numType], optionsValueSet)
}
func (resource *Contract) T_TermAssetSubtype(numTerm int, numAsset int, numSubtype int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset || len(resource.Term[numTerm].Asset[numAsset].Subtype) >= numSubtype {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Subtype["+strconv.Itoa(numSubtype)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Subtype["+strconv.Itoa(numSubtype)+"]", &resource.Term[numTerm].Asset[numAsset].Subtype[numSubtype], optionsValueSet)
}
func (resource *Contract) T_TermAssetRelationship(numTerm int, numAsset int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset {
		return CodingSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Relationship", nil, optionsValueSet)
	}
	return CodingSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Relationship", resource.Term[numTerm].Asset[numAsset].Relationship, optionsValueSet)
}
func (resource *Contract) T_TermAssetCondition(numTerm int, numAsset int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Condition", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Condition", resource.Term[numTerm].Asset[numAsset].Condition)
}
func (resource *Contract) T_TermAssetPeriodType(numTerm int, numAsset int, numPeriodType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset || len(resource.Term[numTerm].Asset[numAsset].PeriodType) >= numPeriodType {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].PeriodType["+strconv.Itoa(numPeriodType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].PeriodType["+strconv.Itoa(numPeriodType)+"]", &resource.Term[numTerm].Asset[numAsset].PeriodType[numPeriodType], optionsValueSet)
}
func (resource *Contract) T_TermAssetText(numTerm int, numAsset int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Text", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Text", resource.Term[numTerm].Asset[numAsset].Text)
}
func (resource *Contract) T_TermAssetLinkId(numTerm int, numAsset int, numLinkId int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset || len(resource.Term[numTerm].Asset[numAsset].LinkId) >= numLinkId {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].LinkId["+strconv.Itoa(numLinkId)+"]", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].LinkId["+strconv.Itoa(numLinkId)+"]", &resource.Term[numTerm].Asset[numAsset].LinkId[numLinkId])
}
func (resource *Contract) T_TermAssetSecurityLabelNumber(numTerm int, numAsset int, numSecurityLabelNumber int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset || len(resource.Term[numTerm].Asset[numAsset].SecurityLabelNumber) >= numSecurityLabelNumber {
		return IntInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].SecurityLabelNumber["+strconv.Itoa(numSecurityLabelNumber)+"]", nil)
	}
	return IntInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].SecurityLabelNumber["+strconv.Itoa(numSecurityLabelNumber)+"]", &resource.Term[numTerm].Asset[numAsset].SecurityLabelNumber[numSecurityLabelNumber])
}
func (resource *Contract) T_TermAssetContextId(numTerm int, numAsset int, numContext int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset || len(resource.Term[numTerm].Asset[numAsset].Context) >= numContext {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Context["+strconv.Itoa(numContext)+"].Id", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Context["+strconv.Itoa(numContext)+"].Id", resource.Term[numTerm].Asset[numAsset].Context[numContext].Id)
}
func (resource *Contract) T_TermAssetContextCode(numTerm int, numAsset int, numContext int, numCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset || len(resource.Term[numTerm].Asset[numAsset].Context) >= numContext || len(resource.Term[numTerm].Asset[numAsset].Context[numContext].Code) >= numCode {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Context["+strconv.Itoa(numContext)+"].Code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Context["+strconv.Itoa(numContext)+"].Code["+strconv.Itoa(numCode)+"]", &resource.Term[numTerm].Asset[numAsset].Context[numContext].Code[numCode], optionsValueSet)
}
func (resource *Contract) T_TermAssetContextText(numTerm int, numAsset int, numContext int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset || len(resource.Term[numTerm].Asset[numAsset].Context) >= numContext {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Context["+strconv.Itoa(numContext)+"].Text", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].Context["+strconv.Itoa(numContext)+"].Text", resource.Term[numTerm].Asset[numAsset].Context[numContext].Text)
}
func (resource *Contract) T_TermAssetValuedItemId(numTerm int, numAsset int, numValuedItem int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset || len(resource.Term[numTerm].Asset[numAsset].ValuedItem) >= numValuedItem {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].ValuedItem["+strconv.Itoa(numValuedItem)+"].Id", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].ValuedItem["+strconv.Itoa(numValuedItem)+"].Id", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].Id)
}
func (resource *Contract) T_TermAssetValuedItemEffectiveTime(numTerm int, numAsset int, numValuedItem int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset || len(resource.Term[numTerm].Asset[numAsset].ValuedItem) >= numValuedItem {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].ValuedItem["+strconv.Itoa(numValuedItem)+"].EffectiveTime", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].ValuedItem["+strconv.Itoa(numValuedItem)+"].EffectiveTime", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].EffectiveTime)
}
func (resource *Contract) T_TermAssetValuedItemFactor(numTerm int, numAsset int, numValuedItem int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset || len(resource.Term[numTerm].Asset[numAsset].ValuedItem) >= numValuedItem {
		return Float64Input("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].ValuedItem["+strconv.Itoa(numValuedItem)+"].Factor", nil)
	}
	return Float64Input("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].ValuedItem["+strconv.Itoa(numValuedItem)+"].Factor", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].Factor)
}
func (resource *Contract) T_TermAssetValuedItemPoints(numTerm int, numAsset int, numValuedItem int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset || len(resource.Term[numTerm].Asset[numAsset].ValuedItem) >= numValuedItem {
		return Float64Input("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].ValuedItem["+strconv.Itoa(numValuedItem)+"].Points", nil)
	}
	return Float64Input("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].ValuedItem["+strconv.Itoa(numValuedItem)+"].Points", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].Points)
}
func (resource *Contract) T_TermAssetValuedItemPayment(numTerm int, numAsset int, numValuedItem int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset || len(resource.Term[numTerm].Asset[numAsset].ValuedItem) >= numValuedItem {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].ValuedItem["+strconv.Itoa(numValuedItem)+"].Payment", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].ValuedItem["+strconv.Itoa(numValuedItem)+"].Payment", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].Payment)
}
func (resource *Contract) T_TermAssetValuedItemPaymentDate(numTerm int, numAsset int, numValuedItem int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset || len(resource.Term[numTerm].Asset[numAsset].ValuedItem) >= numValuedItem {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].ValuedItem["+strconv.Itoa(numValuedItem)+"].PaymentDate", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].ValuedItem["+strconv.Itoa(numValuedItem)+"].PaymentDate", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].PaymentDate)
}
func (resource *Contract) T_TermAssetValuedItemLinkId(numTerm int, numAsset int, numValuedItem int, numLinkId int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset || len(resource.Term[numTerm].Asset[numAsset].ValuedItem) >= numValuedItem || len(resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].LinkId) >= numLinkId {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].ValuedItem["+strconv.Itoa(numValuedItem)+"].LinkId["+strconv.Itoa(numLinkId)+"]", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].ValuedItem["+strconv.Itoa(numValuedItem)+"].LinkId["+strconv.Itoa(numLinkId)+"]", &resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].LinkId[numLinkId])
}
func (resource *Contract) T_TermAssetValuedItemSecurityLabelNumber(numTerm int, numAsset int, numValuedItem int, numSecurityLabelNumber int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Asset) >= numAsset || len(resource.Term[numTerm].Asset[numAsset].ValuedItem) >= numValuedItem || len(resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].SecurityLabelNumber) >= numSecurityLabelNumber {
		return IntInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].ValuedItem["+strconv.Itoa(numValuedItem)+"].SecurityLabelNumber["+strconv.Itoa(numSecurityLabelNumber)+"]", nil)
	}
	return IntInput("Contract.Term["+strconv.Itoa(numTerm)+"].Asset["+strconv.Itoa(numAsset)+"].ValuedItem["+strconv.Itoa(numValuedItem)+"].SecurityLabelNumber["+strconv.Itoa(numSecurityLabelNumber)+"]", &resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].SecurityLabelNumber[numSecurityLabelNumber])
}
func (resource *Contract) T_TermActionId(numTerm int, numAction int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].Id", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].Id", resource.Term[numTerm].Action[numAction].Id)
}
func (resource *Contract) T_TermActionDoNotPerform(numTerm int, numAction int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction {
		return BoolInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].DoNotPerform", nil)
	}
	return BoolInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].DoNotPerform", resource.Term[numTerm].Action[numAction].DoNotPerform)
}
func (resource *Contract) T_TermActionType(numTerm int, numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].Type", &resource.Term[numTerm].Action[numAction].Type, optionsValueSet)
}
func (resource *Contract) T_TermActionIntent(numTerm int, numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].Intent", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].Intent", &resource.Term[numTerm].Action[numAction].Intent, optionsValueSet)
}
func (resource *Contract) T_TermActionLinkId(numTerm int, numAction int, numLinkId int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction || len(resource.Term[numTerm].Action[numAction].LinkId) >= numLinkId {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].LinkId["+strconv.Itoa(numLinkId)+"]", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].LinkId["+strconv.Itoa(numLinkId)+"]", &resource.Term[numTerm].Action[numAction].LinkId[numLinkId])
}
func (resource *Contract) T_TermActionStatus(numTerm int, numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].Status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].Status", &resource.Term[numTerm].Action[numAction].Status, optionsValueSet)
}
func (resource *Contract) T_TermActionContextLinkId(numTerm int, numAction int, numContextLinkId int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction || len(resource.Term[numTerm].Action[numAction].ContextLinkId) >= numContextLinkId {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].ContextLinkId["+strconv.Itoa(numContextLinkId)+"]", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].ContextLinkId["+strconv.Itoa(numContextLinkId)+"]", &resource.Term[numTerm].Action[numAction].ContextLinkId[numContextLinkId])
}
func (resource *Contract) T_TermActionRequesterLinkId(numTerm int, numAction int, numRequesterLinkId int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction || len(resource.Term[numTerm].Action[numAction].RequesterLinkId) >= numRequesterLinkId {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].RequesterLinkId["+strconv.Itoa(numRequesterLinkId)+"]", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].RequesterLinkId["+strconv.Itoa(numRequesterLinkId)+"]", &resource.Term[numTerm].Action[numAction].RequesterLinkId[numRequesterLinkId])
}
func (resource *Contract) T_TermActionPerformerType(numTerm int, numAction int, numPerformerType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction || len(resource.Term[numTerm].Action[numAction].PerformerType) >= numPerformerType {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].PerformerType["+strconv.Itoa(numPerformerType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].PerformerType["+strconv.Itoa(numPerformerType)+"]", &resource.Term[numTerm].Action[numAction].PerformerType[numPerformerType], optionsValueSet)
}
func (resource *Contract) T_TermActionPerformerRole(numTerm int, numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].PerformerRole", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].PerformerRole", resource.Term[numTerm].Action[numAction].PerformerRole, optionsValueSet)
}
func (resource *Contract) T_TermActionPerformerLinkId(numTerm int, numAction int, numPerformerLinkId int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction || len(resource.Term[numTerm].Action[numAction].PerformerLinkId) >= numPerformerLinkId {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].PerformerLinkId["+strconv.Itoa(numPerformerLinkId)+"]", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].PerformerLinkId["+strconv.Itoa(numPerformerLinkId)+"]", &resource.Term[numTerm].Action[numAction].PerformerLinkId[numPerformerLinkId])
}
func (resource *Contract) T_TermActionReasonCode(numTerm int, numAction int, numReasonCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction || len(resource.Term[numTerm].Action[numAction].ReasonCode) >= numReasonCode {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].ReasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].ReasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.Term[numTerm].Action[numAction].ReasonCode[numReasonCode], optionsValueSet)
}
func (resource *Contract) T_TermActionReason(numTerm int, numAction int, numReason int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction || len(resource.Term[numTerm].Action[numAction].Reason) >= numReason {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].Reason["+strconv.Itoa(numReason)+"]", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].Reason["+strconv.Itoa(numReason)+"]", &resource.Term[numTerm].Action[numAction].Reason[numReason])
}
func (resource *Contract) T_TermActionReasonLinkId(numTerm int, numAction int, numReasonLinkId int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction || len(resource.Term[numTerm].Action[numAction].ReasonLinkId) >= numReasonLinkId {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].ReasonLinkId["+strconv.Itoa(numReasonLinkId)+"]", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].ReasonLinkId["+strconv.Itoa(numReasonLinkId)+"]", &resource.Term[numTerm].Action[numAction].ReasonLinkId[numReasonLinkId])
}
func (resource *Contract) T_TermActionSecurityLabelNumber(numTerm int, numAction int, numSecurityLabelNumber int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction || len(resource.Term[numTerm].Action[numAction].SecurityLabelNumber) >= numSecurityLabelNumber {
		return IntInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].SecurityLabelNumber["+strconv.Itoa(numSecurityLabelNumber)+"]", nil)
	}
	return IntInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].SecurityLabelNumber["+strconv.Itoa(numSecurityLabelNumber)+"]", &resource.Term[numTerm].Action[numAction].SecurityLabelNumber[numSecurityLabelNumber])
}
func (resource *Contract) T_TermActionSubjectId(numTerm int, numAction int, numSubject int) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction || len(resource.Term[numTerm].Action[numAction].Subject) >= numSubject {
		return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].Subject["+strconv.Itoa(numSubject)+"].Id", nil)
	}
	return StringInput("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].Subject["+strconv.Itoa(numSubject)+"].Id", resource.Term[numTerm].Action[numAction].Subject[numSubject].Id)
}
func (resource *Contract) T_TermActionSubjectRole(numTerm int, numAction int, numSubject int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Term) >= numTerm || len(resource.Term[numTerm].Action) >= numAction || len(resource.Term[numTerm].Action[numAction].Subject) >= numSubject {
		return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].Subject["+strconv.Itoa(numSubject)+"].Role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Contract.Term["+strconv.Itoa(numTerm)+"].Action["+strconv.Itoa(numAction)+"].Subject["+strconv.Itoa(numSubject)+"].Role", resource.Term[numTerm].Action[numAction].Subject[numSubject].Role, optionsValueSet)
}
func (resource *Contract) T_SignerId(numSigner int) templ.Component {

	if resource == nil || len(resource.Signer) >= numSigner {
		return StringInput("Contract.Signer["+strconv.Itoa(numSigner)+"].Id", nil)
	}
	return StringInput("Contract.Signer["+strconv.Itoa(numSigner)+"].Id", resource.Signer[numSigner].Id)
}
func (resource *Contract) T_SignerType(numSigner int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Signer) >= numSigner {
		return CodingSelect("Contract.Signer["+strconv.Itoa(numSigner)+"].Type", nil, optionsValueSet)
	}
	return CodingSelect("Contract.Signer["+strconv.Itoa(numSigner)+"].Type", &resource.Signer[numSigner].Type, optionsValueSet)
}
func (resource *Contract) T_FriendlyId(numFriendly int) templ.Component {

	if resource == nil || len(resource.Friendly) >= numFriendly {
		return StringInput("Contract.Friendly["+strconv.Itoa(numFriendly)+"].Id", nil)
	}
	return StringInput("Contract.Friendly["+strconv.Itoa(numFriendly)+"].Id", resource.Friendly[numFriendly].Id)
}
func (resource *Contract) T_LegalId(numLegal int) templ.Component {

	if resource == nil || len(resource.Legal) >= numLegal {
		return StringInput("Contract.Legal["+strconv.Itoa(numLegal)+"].Id", nil)
	}
	return StringInput("Contract.Legal["+strconv.Itoa(numLegal)+"].Id", resource.Legal[numLegal].Id)
}
func (resource *Contract) T_RuleId(numRule int) templ.Component {

	if resource == nil || len(resource.Rule) >= numRule {
		return StringInput("Contract.Rule["+strconv.Itoa(numRule)+"].Id", nil)
	}
	return StringInput("Contract.Rule["+strconv.Itoa(numRule)+"].Id", resource.Rule[numRule].Id)
}
