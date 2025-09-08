package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Contract
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
	Issued                   *time.Time                 `json:"issued,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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

// http://hl7.org/fhir/r5/StructureDefinition/Contract
type ContractContentDefinition struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              CodeableConcept  `json:"type"`
	SubType           *CodeableConcept `json:"subType,omitempty"`
	Publisher         *Reference       `json:"publisher,omitempty"`
	PublicationDate   *time.Time       `json:"publicationDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	PublicationStatus string           `json:"publicationStatus"`
	Copyright         *string          `json:"copyright,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Contract
type ContractTerm struct {
	Id                   *string                     `json:"id,omitempty"`
	Extension            []Extension                 `json:"extension,omitempty"`
	ModifierExtension    []Extension                 `json:"modifierExtension,omitempty"`
	Identifier           *Identifier                 `json:"identifier,omitempty"`
	Issued               *time.Time                  `json:"issued,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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

// http://hl7.org/fhir/r5/StructureDefinition/Contract
type ContractTermSecurityLabel struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Number            []int       `json:"number,omitempty"`
	Classification    Coding      `json:"classification"`
	Category          []Coding    `json:"category,omitempty"`
	Control           []Coding    `json:"control,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Contract
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

// http://hl7.org/fhir/r5/StructureDefinition/Contract
type ContractTermOfferParty struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Reference         []Reference     `json:"reference"`
	Role              CodeableConcept `json:"role"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Contract
type ContractTermOfferAnswer struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ValueBoolean      bool        `json:"valueBoolean"`
	ValueDecimal      float64     `json:"valueDecimal"`
	ValueInteger      int         `json:"valueInteger"`
	ValueDate         time.Time   `json:"valueDate,format:'2006-01-02'"`
	ValueDateTime     time.Time   `json:"valueDateTime,format:'2006-01-02T15:04:05Z07:00'"`
	ValueTime         string      `json:"valueTime"`
	ValueString       string      `json:"valueString"`
	ValueUri          string      `json:"valueUri"`
	ValueAttachment   Attachment  `json:"valueAttachment"`
	ValueCoding       Coding      `json:"valueCoding"`
	ValueQuantity     Quantity    `json:"valueQuantity"`
	ValueReference    Reference   `json:"valueReference"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Contract
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

// http://hl7.org/fhir/r5/StructureDefinition/Contract
type ContractTermAssetContext struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Reference         *Reference        `json:"reference,omitempty"`
	Code              []CodeableConcept `json:"code,omitempty"`
	Text              *string           `json:"text,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Contract
type ContractTermAssetValuedItem struct {
	Id                    *string          `json:"id,omitempty"`
	Extension             []Extension      `json:"extension,omitempty"`
	ModifierExtension     []Extension      `json:"modifierExtension,omitempty"`
	EntityCodeableConcept *CodeableConcept `json:"entityCodeableConcept,omitempty"`
	EntityReference       *Reference       `json:"entityReference,omitempty"`
	Identifier            *Identifier      `json:"identifier,omitempty"`
	EffectiveTime         *time.Time       `json:"effectiveTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Quantity              *Quantity        `json:"quantity,omitempty"`
	UnitPrice             *Money           `json:"unitPrice,omitempty"`
	Factor                *float64         `json:"factor,omitempty"`
	Points                *float64         `json:"points,omitempty"`
	Net                   *Money           `json:"net,omitempty"`
	Payment               *string          `json:"payment,omitempty"`
	PaymentDate           *time.Time       `json:"paymentDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	Responsible           *Reference       `json:"responsible,omitempty"`
	Recipient             *Reference       `json:"recipient,omitempty"`
	LinkId                []string         `json:"linkId,omitempty"`
	SecurityLabelNumber   []int            `json:"securityLabelNumber,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Contract
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
	OccurrenceDateTime  *time.Time                  `json:"occurrenceDateTime,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	OccurrencePeriod    *Period                     `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming    *Timing                     `json:"occurrenceTiming,omitempty"`
	Requester           []Reference                 `json:"requester,omitempty"`
	RequesterLinkId     []string                    `json:"requesterLinkId,omitempty"`
	PerformerType       []CodeableConcept           `json:"performerType,omitempty"`
	PerformerRole       *CodeableConcept            `json:"performerRole,omitempty"`
	Performer           *Reference                  `json:"performer,omitempty"`
	PerformerLinkId     []string                    `json:"performerLinkId,omitempty"`
	Reason              []CodeableReference         `json:"reason,omitempty"`
	ReasonLinkId        []string                    `json:"reasonLinkId,omitempty"`
	Note                []Annotation                `json:"note,omitempty"`
	SecurityLabelNumber []int                       `json:"securityLabelNumber,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Contract
type ContractTermActionSubject struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Reference         []Reference      `json:"reference"`
	Role              *CodeableConcept `json:"role,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Contract
type ContractSigner struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              Coding      `json:"type"`
	Party             Reference   `json:"party"`
	Signature         []Signature `json:"signature"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Contract
type ContractFriendly struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ContentAttachment Attachment  `json:"contentAttachment"`
	ContentReference  Reference   `json:"contentReference"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Contract
type ContractLegal struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ContentAttachment Attachment  `json:"contentAttachment"`
	ContentReference  Reference   `json:"contentReference"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Contract
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
func (r Contract) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Contract/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Contract"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Contract) T_Url(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Contract.Url", nil, htmlAttrs)
	}
	return StringInput("Contract.Url", resource.Url, htmlAttrs)
}
func (resource *Contract) T_Version(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Contract.Version", nil, htmlAttrs)
	}
	return StringInput("Contract.Version", resource.Version, htmlAttrs)
}
func (resource *Contract) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSContract_status

	if resource == nil {
		return CodeSelect("Contract.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Contract.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_LegalState(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Contract.LegalState", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.LegalState", resource.LegalState, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_InstantiatesUri(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Contract.InstantiatesUri", nil, htmlAttrs)
	}
	return StringInput("Contract.InstantiatesUri", resource.InstantiatesUri, htmlAttrs)
}
func (resource *Contract) T_ContentDerivative(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Contract.ContentDerivative", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.ContentDerivative", resource.ContentDerivative, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_Issued(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Contract.Issued", nil, htmlAttrs)
	}
	return DateTimeInput("Contract.Issued", resource.Issued, htmlAttrs)
}
func (resource *Contract) T_ExpirationType(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Contract.ExpirationType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.ExpirationType", resource.ExpirationType, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_Name(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Contract.Name", nil, htmlAttrs)
	}
	return StringInput("Contract.Name", resource.Name, htmlAttrs)
}
func (resource *Contract) T_Title(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Contract.Title", nil, htmlAttrs)
	}
	return StringInput("Contract.Title", resource.Title, htmlAttrs)
}
func (resource *Contract) T_Subtitle(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Contract.Subtitle", nil, htmlAttrs)
	}
	return StringInput("Contract.Subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *Contract) T_Alias(numAlias int, htmlAttrs string) templ.Component {

	if resource == nil || numAlias >= len(resource.Alias) {
		return StringInput("Contract.Alias."+strconv.Itoa(numAlias)+".", nil, htmlAttrs)
	}
	return StringInput("Contract.Alias."+strconv.Itoa(numAlias)+".", &resource.Alias[numAlias], htmlAttrs)
}
func (resource *Contract) T_Scope(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Contract.Scope", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Scope", resource.Scope, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TopicCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Contract.TopicCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.TopicCodeableConcept", resource.TopicCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Contract.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_SubType(numSubType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSubType >= len(resource.SubType) {
		return CodeableConceptSelect("Contract.SubType."+strconv.Itoa(numSubType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.SubType."+strconv.Itoa(numSubType)+".", &resource.SubType[numSubType], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_ContentDefinitionType(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Contract.ContentDefinition.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.ContentDefinition.Type", &resource.ContentDefinition.Type, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_ContentDefinitionSubType(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Contract.ContentDefinition.SubType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.ContentDefinition.SubType", resource.ContentDefinition.SubType, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_ContentDefinitionPublicationDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Contract.ContentDefinition.PublicationDate", nil, htmlAttrs)
	}
	return DateTimeInput("Contract.ContentDefinition.PublicationDate", resource.ContentDefinition.PublicationDate, htmlAttrs)
}
func (resource *Contract) T_ContentDefinitionPublicationStatus(htmlAttrs string) templ.Component {
	optionsValueSet := VSContract_publicationstatus

	if resource == nil {
		return CodeSelect("Contract.ContentDefinition.PublicationStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Contract.ContentDefinition.PublicationStatus", &resource.ContentDefinition.PublicationStatus, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_ContentDefinitionCopyright(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Contract.ContentDefinition.Copyright", nil, htmlAttrs)
	}
	return StringInput("Contract.ContentDefinition.Copyright", resource.ContentDefinition.Copyright, htmlAttrs)
}
func (resource *Contract) T_TermIssued(numTerm int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) {
		return DateTimeInput("Contract.Term."+strconv.Itoa(numTerm)+"..Issued", nil, htmlAttrs)
	}
	return DateTimeInput("Contract.Term."+strconv.Itoa(numTerm)+"..Issued", resource.Term[numTerm].Issued, htmlAttrs)
}
func (resource *Contract) T_TermTopicCodeableConcept(numTerm int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..TopicCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..TopicCodeableConcept", resource.Term[numTerm].TopicCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermType(numTerm int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Type", resource.Term[numTerm].Type, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermSubType(numTerm int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..SubType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..SubType", resource.Term[numTerm].SubType, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermText(numTerm int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Text", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Text", resource.Term[numTerm].Text, htmlAttrs)
}
func (resource *Contract) T_TermSecurityLabelNumber(numTerm int, numSecurityLabel int, numNumber int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numSecurityLabel >= len(resource.Term[numTerm].SecurityLabel) || numNumber >= len(resource.Term[numTerm].SecurityLabel[numSecurityLabel].Number) {
		return IntInput("Contract.Term."+strconv.Itoa(numTerm)+"..SecurityLabel."+strconv.Itoa(numSecurityLabel)+"..Number."+strconv.Itoa(numNumber)+".", nil, htmlAttrs)
	}
	return IntInput("Contract.Term."+strconv.Itoa(numTerm)+"..SecurityLabel."+strconv.Itoa(numSecurityLabel)+"..Number."+strconv.Itoa(numNumber)+".", &resource.Term[numTerm].SecurityLabel[numSecurityLabel].Number[numNumber], htmlAttrs)
}
func (resource *Contract) T_TermSecurityLabelClassification(numTerm int, numSecurityLabel int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numSecurityLabel >= len(resource.Term[numTerm].SecurityLabel) {
		return CodingSelect("Contract.Term."+strconv.Itoa(numTerm)+"..SecurityLabel."+strconv.Itoa(numSecurityLabel)+"..Classification", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Contract.Term."+strconv.Itoa(numTerm)+"..SecurityLabel."+strconv.Itoa(numSecurityLabel)+"..Classification", &resource.Term[numTerm].SecurityLabel[numSecurityLabel].Classification, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermSecurityLabelCategory(numTerm int, numSecurityLabel int, numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numSecurityLabel >= len(resource.Term[numTerm].SecurityLabel) || numCategory >= len(resource.Term[numTerm].SecurityLabel[numSecurityLabel].Category) {
		return CodingSelect("Contract.Term."+strconv.Itoa(numTerm)+"..SecurityLabel."+strconv.Itoa(numSecurityLabel)+"..Category."+strconv.Itoa(numCategory)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Contract.Term."+strconv.Itoa(numTerm)+"..SecurityLabel."+strconv.Itoa(numSecurityLabel)+"..Category."+strconv.Itoa(numCategory)+".", &resource.Term[numTerm].SecurityLabel[numSecurityLabel].Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermSecurityLabelControl(numTerm int, numSecurityLabel int, numControl int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numSecurityLabel >= len(resource.Term[numTerm].SecurityLabel) || numControl >= len(resource.Term[numTerm].SecurityLabel[numSecurityLabel].Control) {
		return CodingSelect("Contract.Term."+strconv.Itoa(numTerm)+"..SecurityLabel."+strconv.Itoa(numSecurityLabel)+"..Control."+strconv.Itoa(numControl)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Contract.Term."+strconv.Itoa(numTerm)+"..SecurityLabel."+strconv.Itoa(numSecurityLabel)+"..Control."+strconv.Itoa(numControl)+".", &resource.Term[numTerm].SecurityLabel[numSecurityLabel].Control[numControl], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermOfferType(numTerm int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Type", resource.Term[numTerm].Offer.Type, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermOfferDecision(numTerm int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Decision", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Decision", resource.Term[numTerm].Offer.Decision, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermOfferDecisionMode(numTerm int, numDecisionMode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numDecisionMode >= len(resource.Term[numTerm].Offer.DecisionMode) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.DecisionMode."+strconv.Itoa(numDecisionMode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.DecisionMode."+strconv.Itoa(numDecisionMode)+".", &resource.Term[numTerm].Offer.DecisionMode[numDecisionMode], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermOfferText(numTerm int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Text", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Text", resource.Term[numTerm].Offer.Text, htmlAttrs)
}
func (resource *Contract) T_TermOfferLinkId(numTerm int, numLinkId int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numLinkId >= len(resource.Term[numTerm].Offer.LinkId) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.LinkId."+strconv.Itoa(numLinkId)+".", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.LinkId."+strconv.Itoa(numLinkId)+".", &resource.Term[numTerm].Offer.LinkId[numLinkId], htmlAttrs)
}
func (resource *Contract) T_TermOfferSecurityLabelNumber(numTerm int, numSecurityLabelNumber int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numSecurityLabelNumber >= len(resource.Term[numTerm].Offer.SecurityLabelNumber) {
		return IntInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.SecurityLabelNumber."+strconv.Itoa(numSecurityLabelNumber)+".", nil, htmlAttrs)
	}
	return IntInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.SecurityLabelNumber."+strconv.Itoa(numSecurityLabelNumber)+".", &resource.Term[numTerm].Offer.SecurityLabelNumber[numSecurityLabelNumber], htmlAttrs)
}
func (resource *Contract) T_TermOfferPartyRole(numTerm int, numParty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numParty >= len(resource.Term[numTerm].Offer.Party) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Party."+strconv.Itoa(numParty)+"..Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Party."+strconv.Itoa(numParty)+"..Role", &resource.Term[numTerm].Offer.Party[numParty].Role, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueBoolean(numTerm int, numAnswer int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return BoolInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueBoolean", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueBoolean, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueDecimal(numTerm int, numAnswer int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return Float64Input("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueDecimal", nil, htmlAttrs)
	}
	return Float64Input("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueDecimal", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueDecimal, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueInteger(numTerm int, numAnswer int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return IntInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueInteger", nil, htmlAttrs)
	}
	return IntInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueInteger", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueInteger, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueDate(numTerm int, numAnswer int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return DateInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueDate", nil, htmlAttrs)
	}
	return DateInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueDate", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueDate, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueDateTime(numTerm int, numAnswer int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return DateTimeInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueDateTime", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueDateTime, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueTime(numTerm int, numAnswer int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueTime", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueTime", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueTime, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueString(numTerm int, numAnswer int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueString", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueString", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueString, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueUri(numTerm int, numAnswer int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueUri", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueUri", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueUri, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueCoding(numTerm int, numAnswer int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return CodingSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Offer.Answer."+strconv.Itoa(numAnswer)+"..ValueCoding", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermAssetScope(numTerm int, numAsset int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..Scope", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..Scope", resource.Term[numTerm].Asset[numAsset].Scope, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermAssetType(numTerm int, numAsset int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numType >= len(resource.Term[numTerm].Asset[numAsset].Type) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..Type."+strconv.Itoa(numType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..Type."+strconv.Itoa(numType)+".", &resource.Term[numTerm].Asset[numAsset].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermAssetSubtype(numTerm int, numAsset int, numSubtype int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numSubtype >= len(resource.Term[numTerm].Asset[numAsset].Subtype) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..Subtype."+strconv.Itoa(numSubtype)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..Subtype."+strconv.Itoa(numSubtype)+".", &resource.Term[numTerm].Asset[numAsset].Subtype[numSubtype], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermAssetRelationship(numTerm int, numAsset int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) {
		return CodingSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..Relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..Relationship", resource.Term[numTerm].Asset[numAsset].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermAssetCondition(numTerm int, numAsset int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..Condition", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..Condition", resource.Term[numTerm].Asset[numAsset].Condition, htmlAttrs)
}
func (resource *Contract) T_TermAssetPeriodType(numTerm int, numAsset int, numPeriodType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numPeriodType >= len(resource.Term[numTerm].Asset[numAsset].PeriodType) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..PeriodType."+strconv.Itoa(numPeriodType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..PeriodType."+strconv.Itoa(numPeriodType)+".", &resource.Term[numTerm].Asset[numAsset].PeriodType[numPeriodType], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermAssetText(numTerm int, numAsset int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..Text", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..Text", resource.Term[numTerm].Asset[numAsset].Text, htmlAttrs)
}
func (resource *Contract) T_TermAssetLinkId(numTerm int, numAsset int, numLinkId int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numLinkId >= len(resource.Term[numTerm].Asset[numAsset].LinkId) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..LinkId."+strconv.Itoa(numLinkId)+".", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..LinkId."+strconv.Itoa(numLinkId)+".", &resource.Term[numTerm].Asset[numAsset].LinkId[numLinkId], htmlAttrs)
}
func (resource *Contract) T_TermAssetSecurityLabelNumber(numTerm int, numAsset int, numSecurityLabelNumber int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numSecurityLabelNumber >= len(resource.Term[numTerm].Asset[numAsset].SecurityLabelNumber) {
		return IntInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..SecurityLabelNumber."+strconv.Itoa(numSecurityLabelNumber)+".", nil, htmlAttrs)
	}
	return IntInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..SecurityLabelNumber."+strconv.Itoa(numSecurityLabelNumber)+".", &resource.Term[numTerm].Asset[numAsset].SecurityLabelNumber[numSecurityLabelNumber], htmlAttrs)
}
func (resource *Contract) T_TermAssetContextCode(numTerm int, numAsset int, numContext int, numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numContext >= len(resource.Term[numTerm].Asset[numAsset].Context) || numCode >= len(resource.Term[numTerm].Asset[numAsset].Context[numContext].Code) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..Context."+strconv.Itoa(numContext)+"..Code."+strconv.Itoa(numCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..Context."+strconv.Itoa(numContext)+"..Code."+strconv.Itoa(numCode)+".", &resource.Term[numTerm].Asset[numAsset].Context[numContext].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermAssetContextText(numTerm int, numAsset int, numContext int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numContext >= len(resource.Term[numTerm].Asset[numAsset].Context) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..Context."+strconv.Itoa(numContext)+"..Text", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..Context."+strconv.Itoa(numContext)+"..Text", resource.Term[numTerm].Asset[numAsset].Context[numContext].Text, htmlAttrs)
}
func (resource *Contract) T_TermAssetValuedItemEntityCodeableConcept(numTerm int, numAsset int, numValuedItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numValuedItem >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..ValuedItem."+strconv.Itoa(numValuedItem)+"..EntityCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..ValuedItem."+strconv.Itoa(numValuedItem)+"..EntityCodeableConcept", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].EntityCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermAssetValuedItemEffectiveTime(numTerm int, numAsset int, numValuedItem int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numValuedItem >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem) {
		return DateTimeInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..ValuedItem."+strconv.Itoa(numValuedItem)+"..EffectiveTime", nil, htmlAttrs)
	}
	return DateTimeInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..ValuedItem."+strconv.Itoa(numValuedItem)+"..EffectiveTime", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].EffectiveTime, htmlAttrs)
}
func (resource *Contract) T_TermAssetValuedItemFactor(numTerm int, numAsset int, numValuedItem int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numValuedItem >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem) {
		return Float64Input("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..ValuedItem."+strconv.Itoa(numValuedItem)+"..Factor", nil, htmlAttrs)
	}
	return Float64Input("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..ValuedItem."+strconv.Itoa(numValuedItem)+"..Factor", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].Factor, htmlAttrs)
}
func (resource *Contract) T_TermAssetValuedItemPoints(numTerm int, numAsset int, numValuedItem int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numValuedItem >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem) {
		return Float64Input("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..ValuedItem."+strconv.Itoa(numValuedItem)+"..Points", nil, htmlAttrs)
	}
	return Float64Input("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..ValuedItem."+strconv.Itoa(numValuedItem)+"..Points", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].Points, htmlAttrs)
}
func (resource *Contract) T_TermAssetValuedItemPayment(numTerm int, numAsset int, numValuedItem int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numValuedItem >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..ValuedItem."+strconv.Itoa(numValuedItem)+"..Payment", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..ValuedItem."+strconv.Itoa(numValuedItem)+"..Payment", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].Payment, htmlAttrs)
}
func (resource *Contract) T_TermAssetValuedItemPaymentDate(numTerm int, numAsset int, numValuedItem int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numValuedItem >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem) {
		return DateTimeInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..ValuedItem."+strconv.Itoa(numValuedItem)+"..PaymentDate", nil, htmlAttrs)
	}
	return DateTimeInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..ValuedItem."+strconv.Itoa(numValuedItem)+"..PaymentDate", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].PaymentDate, htmlAttrs)
}
func (resource *Contract) T_TermAssetValuedItemLinkId(numTerm int, numAsset int, numValuedItem int, numLinkId int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numValuedItem >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem) || numLinkId >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].LinkId) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..ValuedItem."+strconv.Itoa(numValuedItem)+"..LinkId."+strconv.Itoa(numLinkId)+".", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..ValuedItem."+strconv.Itoa(numValuedItem)+"..LinkId."+strconv.Itoa(numLinkId)+".", &resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].LinkId[numLinkId], htmlAttrs)
}
func (resource *Contract) T_TermAssetValuedItemSecurityLabelNumber(numTerm int, numAsset int, numValuedItem int, numSecurityLabelNumber int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numValuedItem >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem) || numSecurityLabelNumber >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].SecurityLabelNumber) {
		return IntInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..ValuedItem."+strconv.Itoa(numValuedItem)+"..SecurityLabelNumber."+strconv.Itoa(numSecurityLabelNumber)+".", nil, htmlAttrs)
	}
	return IntInput("Contract.Term."+strconv.Itoa(numTerm)+"..Asset."+strconv.Itoa(numAsset)+"..ValuedItem."+strconv.Itoa(numValuedItem)+"..SecurityLabelNumber."+strconv.Itoa(numSecurityLabelNumber)+".", &resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].SecurityLabelNumber[numSecurityLabelNumber], htmlAttrs)
}
func (resource *Contract) T_TermActionDoNotPerform(numTerm int, numAction int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) {
		return BoolInput("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..DoNotPerform", nil, htmlAttrs)
	}
	return BoolInput("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..DoNotPerform", resource.Term[numTerm].Action[numAction].DoNotPerform, htmlAttrs)
}
func (resource *Contract) T_TermActionType(numTerm int, numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..Type", &resource.Term[numTerm].Action[numAction].Type, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermActionIntent(numTerm int, numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..Intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..Intent", &resource.Term[numTerm].Action[numAction].Intent, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermActionLinkId(numTerm int, numAction int, numLinkId int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numLinkId >= len(resource.Term[numTerm].Action[numAction].LinkId) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..LinkId."+strconv.Itoa(numLinkId)+".", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..LinkId."+strconv.Itoa(numLinkId)+".", &resource.Term[numTerm].Action[numAction].LinkId[numLinkId], htmlAttrs)
}
func (resource *Contract) T_TermActionStatus(numTerm int, numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..Status", &resource.Term[numTerm].Action[numAction].Status, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermActionContextLinkId(numTerm int, numAction int, numContextLinkId int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numContextLinkId >= len(resource.Term[numTerm].Action[numAction].ContextLinkId) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..ContextLinkId."+strconv.Itoa(numContextLinkId)+".", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..ContextLinkId."+strconv.Itoa(numContextLinkId)+".", &resource.Term[numTerm].Action[numAction].ContextLinkId[numContextLinkId], htmlAttrs)
}
func (resource *Contract) T_TermActionOccurrenceDateTime(numTerm int, numAction int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) {
		return DateTimeInput("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..OccurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..OccurrenceDateTime", resource.Term[numTerm].Action[numAction].OccurrenceDateTime, htmlAttrs)
}
func (resource *Contract) T_TermActionRequesterLinkId(numTerm int, numAction int, numRequesterLinkId int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numRequesterLinkId >= len(resource.Term[numTerm].Action[numAction].RequesterLinkId) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..RequesterLinkId."+strconv.Itoa(numRequesterLinkId)+".", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..RequesterLinkId."+strconv.Itoa(numRequesterLinkId)+".", &resource.Term[numTerm].Action[numAction].RequesterLinkId[numRequesterLinkId], htmlAttrs)
}
func (resource *Contract) T_TermActionPerformerType(numTerm int, numAction int, numPerformerType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numPerformerType >= len(resource.Term[numTerm].Action[numAction].PerformerType) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..PerformerType."+strconv.Itoa(numPerformerType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..PerformerType."+strconv.Itoa(numPerformerType)+".", &resource.Term[numTerm].Action[numAction].PerformerType[numPerformerType], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermActionPerformerRole(numTerm int, numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..PerformerRole", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..PerformerRole", resource.Term[numTerm].Action[numAction].PerformerRole, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermActionPerformerLinkId(numTerm int, numAction int, numPerformerLinkId int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numPerformerLinkId >= len(resource.Term[numTerm].Action[numAction].PerformerLinkId) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..PerformerLinkId."+strconv.Itoa(numPerformerLinkId)+".", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..PerformerLinkId."+strconv.Itoa(numPerformerLinkId)+".", &resource.Term[numTerm].Action[numAction].PerformerLinkId[numPerformerLinkId], htmlAttrs)
}
func (resource *Contract) T_TermActionReasonLinkId(numTerm int, numAction int, numReasonLinkId int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numReasonLinkId >= len(resource.Term[numTerm].Action[numAction].ReasonLinkId) {
		return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..ReasonLinkId."+strconv.Itoa(numReasonLinkId)+".", nil, htmlAttrs)
	}
	return StringInput("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..ReasonLinkId."+strconv.Itoa(numReasonLinkId)+".", &resource.Term[numTerm].Action[numAction].ReasonLinkId[numReasonLinkId], htmlAttrs)
}
func (resource *Contract) T_TermActionNote(numTerm int, numAction int, numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numNote >= len(resource.Term[numTerm].Action[numAction].Note) {
		return AnnotationTextArea("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..Note."+strconv.Itoa(numNote)+".", &resource.Term[numTerm].Action[numAction].Note[numNote], htmlAttrs)
}
func (resource *Contract) T_TermActionSecurityLabelNumber(numTerm int, numAction int, numSecurityLabelNumber int, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numSecurityLabelNumber >= len(resource.Term[numTerm].Action[numAction].SecurityLabelNumber) {
		return IntInput("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..SecurityLabelNumber."+strconv.Itoa(numSecurityLabelNumber)+".", nil, htmlAttrs)
	}
	return IntInput("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..SecurityLabelNumber."+strconv.Itoa(numSecurityLabelNumber)+".", &resource.Term[numTerm].Action[numAction].SecurityLabelNumber[numSecurityLabelNumber], htmlAttrs)
}
func (resource *Contract) T_TermActionSubjectRole(numTerm int, numAction int, numSubject int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numSubject >= len(resource.Term[numTerm].Action[numAction].Subject) {
		return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..Subject."+strconv.Itoa(numSubject)+"..Role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Contract.Term."+strconv.Itoa(numTerm)+"..Action."+strconv.Itoa(numAction)+"..Subject."+strconv.Itoa(numSubject)+"..Role", resource.Term[numTerm].Action[numAction].Subject[numSubject].Role, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_SignerType(numSigner int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSigner >= len(resource.Signer) {
		return CodingSelect("Contract.Signer."+strconv.Itoa(numSigner)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("Contract.Signer."+strconv.Itoa(numSigner)+"..Type", &resource.Signer[numSigner].Type, optionsValueSet, htmlAttrs)
}
