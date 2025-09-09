package r4b

//generated with command go run ./bultaoreune
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
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *Contract) T_Version(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("version", nil, htmlAttrs)
	}
	return StringInput("version", resource.Version, htmlAttrs)
}
func (resource *Contract) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSContract_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_LegalState(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("legalState", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("legalState", resource.LegalState, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_InstantiatesUri(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("instantiatesUri", nil, htmlAttrs)
	}
	return StringInput("instantiatesUri", resource.InstantiatesUri, htmlAttrs)
}
func (resource *Contract) T_ContentDerivative(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("contentDerivative", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("contentDerivative", resource.ContentDerivative, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_Issued(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("issued", nil, htmlAttrs)
	}
	return DateTimeInput("issued", resource.Issued, htmlAttrs)
}
func (resource *Contract) T_ExpirationType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("expirationType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("expirationType", resource.ExpirationType, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_Name(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("name", nil, htmlAttrs)
	}
	return StringInput("name", resource.Name, htmlAttrs)
}
func (resource *Contract) T_Title(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("title", nil, htmlAttrs)
	}
	return StringInput("title", resource.Title, htmlAttrs)
}
func (resource *Contract) T_Subtitle(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("subtitle", nil, htmlAttrs)
	}
	return StringInput("subtitle", resource.Subtitle, htmlAttrs)
}
func (resource *Contract) T_Alias(numAlias int, htmlAttrs string) templ.Component {
	if resource == nil || numAlias >= len(resource.Alias) {
		return StringInput("alias["+strconv.Itoa(numAlias)+"]", nil, htmlAttrs)
	}
	return StringInput("alias["+strconv.Itoa(numAlias)+"]", &resource.Alias[numAlias], htmlAttrs)
}
func (resource *Contract) T_Scope(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("scope", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("scope", resource.Scope, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TopicCodeableConcept(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("topicCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("topicCodeableConcept", resource.TopicCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_SubType(numSubType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSubType >= len(resource.SubType) {
		return CodeableConceptSelect("subType["+strconv.Itoa(numSubType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subType["+strconv.Itoa(numSubType)+"]", &resource.SubType[numSubType], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_ContentDefinitionType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("contentDefinition.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("contentDefinition.type", &resource.ContentDefinition.Type, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_ContentDefinitionSubType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("contentDefinition.subType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("contentDefinition.subType", resource.ContentDefinition.SubType, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_ContentDefinitionPublicationDate(htmlAttrs string) templ.Component {
	if resource == nil {
		return DateTimeInput("contentDefinition.publicationDate", nil, htmlAttrs)
	}
	return DateTimeInput("contentDefinition.publicationDate", resource.ContentDefinition.PublicationDate, htmlAttrs)
}
func (resource *Contract) T_ContentDefinitionPublicationStatus(htmlAttrs string) templ.Component {
	optionsValueSet := VSContract_publicationstatus

	if resource == nil {
		return CodeSelect("contentDefinition.publicationStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("contentDefinition.publicationStatus", &resource.ContentDefinition.PublicationStatus, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_ContentDefinitionCopyright(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("contentDefinition.copyright", nil, htmlAttrs)
	}
	return StringInput("contentDefinition.copyright", resource.ContentDefinition.Copyright, htmlAttrs)
}
func (resource *Contract) T_TermIssued(numTerm int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) {
		return DateTimeInput("term["+strconv.Itoa(numTerm)+"].issued", nil, htmlAttrs)
	}
	return DateTimeInput("term["+strconv.Itoa(numTerm)+"].issued", resource.Term[numTerm].Issued, htmlAttrs)
}
func (resource *Contract) T_TermTopicCodeableConcept(numTerm int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].topicCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].topicCodeableConcept", resource.Term[numTerm].TopicCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermType(numTerm int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].type", resource.Term[numTerm].Type, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermSubType(numTerm int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].subType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].subType", resource.Term[numTerm].SubType, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermText(numTerm int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].text", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].text", resource.Term[numTerm].Text, htmlAttrs)
}
func (resource *Contract) T_TermSecurityLabelNumber(numTerm int, numSecurityLabel int, numNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numSecurityLabel >= len(resource.Term[numTerm].SecurityLabel) || numNumber >= len(resource.Term[numTerm].SecurityLabel[numSecurityLabel].Number) {
		return IntInput("term["+strconv.Itoa(numTerm)+"].securityLabel["+strconv.Itoa(numSecurityLabel)+"].number["+strconv.Itoa(numNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("term["+strconv.Itoa(numTerm)+"].securityLabel["+strconv.Itoa(numSecurityLabel)+"].number["+strconv.Itoa(numNumber)+"]", &resource.Term[numTerm].SecurityLabel[numSecurityLabel].Number[numNumber], htmlAttrs)
}
func (resource *Contract) T_TermSecurityLabelClassification(numTerm int, numSecurityLabel int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numSecurityLabel >= len(resource.Term[numTerm].SecurityLabel) {
		return CodingSelect("term["+strconv.Itoa(numTerm)+"].securityLabel["+strconv.Itoa(numSecurityLabel)+"].classification", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("term["+strconv.Itoa(numTerm)+"].securityLabel["+strconv.Itoa(numSecurityLabel)+"].classification", &resource.Term[numTerm].SecurityLabel[numSecurityLabel].Classification, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermSecurityLabelCategory(numTerm int, numSecurityLabel int, numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numSecurityLabel >= len(resource.Term[numTerm].SecurityLabel) || numCategory >= len(resource.Term[numTerm].SecurityLabel[numSecurityLabel].Category) {
		return CodingSelect("term["+strconv.Itoa(numTerm)+"].securityLabel["+strconv.Itoa(numSecurityLabel)+"].category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("term["+strconv.Itoa(numTerm)+"].securityLabel["+strconv.Itoa(numSecurityLabel)+"].category["+strconv.Itoa(numCategory)+"]", &resource.Term[numTerm].SecurityLabel[numSecurityLabel].Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermSecurityLabelControl(numTerm int, numSecurityLabel int, numControl int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numSecurityLabel >= len(resource.Term[numTerm].SecurityLabel) || numControl >= len(resource.Term[numTerm].SecurityLabel[numSecurityLabel].Control) {
		return CodingSelect("term["+strconv.Itoa(numTerm)+"].securityLabel["+strconv.Itoa(numSecurityLabel)+"].control["+strconv.Itoa(numControl)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("term["+strconv.Itoa(numTerm)+"].securityLabel["+strconv.Itoa(numSecurityLabel)+"].control["+strconv.Itoa(numControl)+"]", &resource.Term[numTerm].SecurityLabel[numSecurityLabel].Control[numControl], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermOfferType(numTerm int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].offer.type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].offer.type", resource.Term[numTerm].Offer.Type, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermOfferDecision(numTerm int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].offer.decision", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].offer.decision", resource.Term[numTerm].Offer.Decision, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermOfferDecisionMode(numTerm int, numDecisionMode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numDecisionMode >= len(resource.Term[numTerm].Offer.DecisionMode) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].offer.decisionMode["+strconv.Itoa(numDecisionMode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].offer.decisionMode["+strconv.Itoa(numDecisionMode)+"]", &resource.Term[numTerm].Offer.DecisionMode[numDecisionMode], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermOfferText(numTerm int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].offer.text", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].offer.text", resource.Term[numTerm].Offer.Text, htmlAttrs)
}
func (resource *Contract) T_TermOfferLinkId(numTerm int, numLinkId int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numLinkId >= len(resource.Term[numTerm].Offer.LinkId) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].offer.linkId["+strconv.Itoa(numLinkId)+"]", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].offer.linkId["+strconv.Itoa(numLinkId)+"]", &resource.Term[numTerm].Offer.LinkId[numLinkId], htmlAttrs)
}
func (resource *Contract) T_TermOfferSecurityLabelNumber(numTerm int, numSecurityLabelNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numSecurityLabelNumber >= len(resource.Term[numTerm].Offer.SecurityLabelNumber) {
		return IntInput("term["+strconv.Itoa(numTerm)+"].offer.securityLabelNumber["+strconv.Itoa(numSecurityLabelNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("term["+strconv.Itoa(numTerm)+"].offer.securityLabelNumber["+strconv.Itoa(numSecurityLabelNumber)+"]", &resource.Term[numTerm].Offer.SecurityLabelNumber[numSecurityLabelNumber], htmlAttrs)
}
func (resource *Contract) T_TermOfferPartyRole(numTerm int, numParty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numParty >= len(resource.Term[numTerm].Offer.Party) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].offer.party["+strconv.Itoa(numParty)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].offer.party["+strconv.Itoa(numParty)+"].role", &resource.Term[numTerm].Offer.Party[numParty].Role, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueBoolean(numTerm int, numAnswer int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return BoolInput("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueBoolean", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueBoolean, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueDecimal(numTerm int, numAnswer int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return Float64Input("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueDecimal", nil, htmlAttrs)
	}
	return Float64Input("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueDecimal", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueDecimal, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueInteger(numTerm int, numAnswer int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return IntInput("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueInteger", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueInteger, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueDate(numTerm int, numAnswer int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return DateInput("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueDate", nil, htmlAttrs)
	}
	return DateInput("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueDate", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueDate, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueDateTime(numTerm int, numAnswer int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return DateTimeInput("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueDateTime", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueDateTime, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueTime(numTerm int, numAnswer int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueTime", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueTime", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueTime, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueString(numTerm int, numAnswer int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueString", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueString, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueUri(numTerm int, numAnswer int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueUri", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueUri", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueUri, htmlAttrs)
}
func (resource *Contract) T_TermOfferAnswerValueCoding(numTerm int, numAnswer int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAnswer >= len(resource.Term[numTerm].Offer.Answer) {
		return CodingSelect("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueCoding", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("term["+strconv.Itoa(numTerm)+"].offer.answer["+strconv.Itoa(numAnswer)+"].valueCoding", &resource.Term[numTerm].Offer.Answer[numAnswer].ValueCoding, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermAssetScope(numTerm int, numAsset int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].scope", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].scope", resource.Term[numTerm].Asset[numAsset].Scope, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermAssetType(numTerm int, numAsset int, numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numType >= len(resource.Term[numTerm].Asset[numAsset].Type) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].type["+strconv.Itoa(numType)+"]", &resource.Term[numTerm].Asset[numAsset].Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermAssetSubtype(numTerm int, numAsset int, numSubtype int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numSubtype >= len(resource.Term[numTerm].Asset[numAsset].Subtype) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].subtype["+strconv.Itoa(numSubtype)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].subtype["+strconv.Itoa(numSubtype)+"]", &resource.Term[numTerm].Asset[numAsset].Subtype[numSubtype], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermAssetRelationship(numTerm int, numAsset int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) {
		return CodingSelect("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].relationship", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].relationship", resource.Term[numTerm].Asset[numAsset].Relationship, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermAssetCondition(numTerm int, numAsset int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].condition", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].condition", resource.Term[numTerm].Asset[numAsset].Condition, htmlAttrs)
}
func (resource *Contract) T_TermAssetPeriodType(numTerm int, numAsset int, numPeriodType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numPeriodType >= len(resource.Term[numTerm].Asset[numAsset].PeriodType) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].periodType["+strconv.Itoa(numPeriodType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].periodType["+strconv.Itoa(numPeriodType)+"]", &resource.Term[numTerm].Asset[numAsset].PeriodType[numPeriodType], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermAssetText(numTerm int, numAsset int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].text", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].text", resource.Term[numTerm].Asset[numAsset].Text, htmlAttrs)
}
func (resource *Contract) T_TermAssetLinkId(numTerm int, numAsset int, numLinkId int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numLinkId >= len(resource.Term[numTerm].Asset[numAsset].LinkId) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].linkId["+strconv.Itoa(numLinkId)+"]", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].linkId["+strconv.Itoa(numLinkId)+"]", &resource.Term[numTerm].Asset[numAsset].LinkId[numLinkId], htmlAttrs)
}
func (resource *Contract) T_TermAssetSecurityLabelNumber(numTerm int, numAsset int, numSecurityLabelNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numSecurityLabelNumber >= len(resource.Term[numTerm].Asset[numAsset].SecurityLabelNumber) {
		return IntInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].securityLabelNumber["+strconv.Itoa(numSecurityLabelNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].securityLabelNumber["+strconv.Itoa(numSecurityLabelNumber)+"]", &resource.Term[numTerm].Asset[numAsset].SecurityLabelNumber[numSecurityLabelNumber], htmlAttrs)
}
func (resource *Contract) T_TermAssetContextCode(numTerm int, numAsset int, numContext int, numCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numContext >= len(resource.Term[numTerm].Asset[numAsset].Context) || numCode >= len(resource.Term[numTerm].Asset[numAsset].Context[numContext].Code) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].context["+strconv.Itoa(numContext)+"].code["+strconv.Itoa(numCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].context["+strconv.Itoa(numContext)+"].code["+strconv.Itoa(numCode)+"]", &resource.Term[numTerm].Asset[numAsset].Context[numContext].Code[numCode], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermAssetContextText(numTerm int, numAsset int, numContext int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numContext >= len(resource.Term[numTerm].Asset[numAsset].Context) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].context["+strconv.Itoa(numContext)+"].text", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].context["+strconv.Itoa(numContext)+"].text", resource.Term[numTerm].Asset[numAsset].Context[numContext].Text, htmlAttrs)
}
func (resource *Contract) T_TermAssetValuedItemEntityCodeableConcept(numTerm int, numAsset int, numValuedItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numValuedItem >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].valuedItem["+strconv.Itoa(numValuedItem)+"].entityCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].valuedItem["+strconv.Itoa(numValuedItem)+"].entityCodeableConcept", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].EntityCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermAssetValuedItemEffectiveTime(numTerm int, numAsset int, numValuedItem int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numValuedItem >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem) {
		return DateTimeInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].valuedItem["+strconv.Itoa(numValuedItem)+"].effectiveTime", nil, htmlAttrs)
	}
	return DateTimeInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].valuedItem["+strconv.Itoa(numValuedItem)+"].effectiveTime", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].EffectiveTime, htmlAttrs)
}
func (resource *Contract) T_TermAssetValuedItemFactor(numTerm int, numAsset int, numValuedItem int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numValuedItem >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem) {
		return Float64Input("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].valuedItem["+strconv.Itoa(numValuedItem)+"].factor", nil, htmlAttrs)
	}
	return Float64Input("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].valuedItem["+strconv.Itoa(numValuedItem)+"].factor", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].Factor, htmlAttrs)
}
func (resource *Contract) T_TermAssetValuedItemPoints(numTerm int, numAsset int, numValuedItem int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numValuedItem >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem) {
		return Float64Input("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].valuedItem["+strconv.Itoa(numValuedItem)+"].points", nil, htmlAttrs)
	}
	return Float64Input("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].valuedItem["+strconv.Itoa(numValuedItem)+"].points", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].Points, htmlAttrs)
}
func (resource *Contract) T_TermAssetValuedItemPayment(numTerm int, numAsset int, numValuedItem int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numValuedItem >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].valuedItem["+strconv.Itoa(numValuedItem)+"].payment", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].valuedItem["+strconv.Itoa(numValuedItem)+"].payment", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].Payment, htmlAttrs)
}
func (resource *Contract) T_TermAssetValuedItemPaymentDate(numTerm int, numAsset int, numValuedItem int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numValuedItem >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem) {
		return DateTimeInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].valuedItem["+strconv.Itoa(numValuedItem)+"].paymentDate", nil, htmlAttrs)
	}
	return DateTimeInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].valuedItem["+strconv.Itoa(numValuedItem)+"].paymentDate", resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].PaymentDate, htmlAttrs)
}
func (resource *Contract) T_TermAssetValuedItemLinkId(numTerm int, numAsset int, numValuedItem int, numLinkId int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numValuedItem >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem) || numLinkId >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].LinkId) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].valuedItem["+strconv.Itoa(numValuedItem)+"].linkId["+strconv.Itoa(numLinkId)+"]", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].valuedItem["+strconv.Itoa(numValuedItem)+"].linkId["+strconv.Itoa(numLinkId)+"]", &resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].LinkId[numLinkId], htmlAttrs)
}
func (resource *Contract) T_TermAssetValuedItemSecurityLabelNumber(numTerm int, numAsset int, numValuedItem int, numSecurityLabelNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAsset >= len(resource.Term[numTerm].Asset) || numValuedItem >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem) || numSecurityLabelNumber >= len(resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].SecurityLabelNumber) {
		return IntInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].valuedItem["+strconv.Itoa(numValuedItem)+"].securityLabelNumber["+strconv.Itoa(numSecurityLabelNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("term["+strconv.Itoa(numTerm)+"].asset["+strconv.Itoa(numAsset)+"].valuedItem["+strconv.Itoa(numValuedItem)+"].securityLabelNumber["+strconv.Itoa(numSecurityLabelNumber)+"]", &resource.Term[numTerm].Asset[numAsset].ValuedItem[numValuedItem].SecurityLabelNumber[numSecurityLabelNumber], htmlAttrs)
}
func (resource *Contract) T_TermActionDoNotPerform(numTerm int, numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) {
		return BoolInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].doNotPerform", nil, htmlAttrs)
	}
	return BoolInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].doNotPerform", resource.Term[numTerm].Action[numAction].DoNotPerform, htmlAttrs)
}
func (resource *Contract) T_TermActionType(numTerm int, numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].type", &resource.Term[numTerm].Action[numAction].Type, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermActionIntent(numTerm int, numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].intent", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].intent", &resource.Term[numTerm].Action[numAction].Intent, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermActionLinkId(numTerm int, numAction int, numLinkId int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numLinkId >= len(resource.Term[numTerm].Action[numAction].LinkId) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].linkId["+strconv.Itoa(numLinkId)+"]", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].linkId["+strconv.Itoa(numLinkId)+"]", &resource.Term[numTerm].Action[numAction].LinkId[numLinkId], htmlAttrs)
}
func (resource *Contract) T_TermActionStatus(numTerm int, numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].status", &resource.Term[numTerm].Action[numAction].Status, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermActionContextLinkId(numTerm int, numAction int, numContextLinkId int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numContextLinkId >= len(resource.Term[numTerm].Action[numAction].ContextLinkId) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].contextLinkId["+strconv.Itoa(numContextLinkId)+"]", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].contextLinkId["+strconv.Itoa(numContextLinkId)+"]", &resource.Term[numTerm].Action[numAction].ContextLinkId[numContextLinkId], htmlAttrs)
}
func (resource *Contract) T_TermActionOccurrenceDateTime(numTerm int, numAction int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) {
		return DateTimeInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].occurrenceDateTime", nil, htmlAttrs)
	}
	return DateTimeInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].occurrenceDateTime", resource.Term[numTerm].Action[numAction].OccurrenceDateTime, htmlAttrs)
}
func (resource *Contract) T_TermActionRequesterLinkId(numTerm int, numAction int, numRequesterLinkId int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numRequesterLinkId >= len(resource.Term[numTerm].Action[numAction].RequesterLinkId) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].requesterLinkId["+strconv.Itoa(numRequesterLinkId)+"]", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].requesterLinkId["+strconv.Itoa(numRequesterLinkId)+"]", &resource.Term[numTerm].Action[numAction].RequesterLinkId[numRequesterLinkId], htmlAttrs)
}
func (resource *Contract) T_TermActionPerformerType(numTerm int, numAction int, numPerformerType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numPerformerType >= len(resource.Term[numTerm].Action[numAction].PerformerType) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].performerType["+strconv.Itoa(numPerformerType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].performerType["+strconv.Itoa(numPerformerType)+"]", &resource.Term[numTerm].Action[numAction].PerformerType[numPerformerType], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermActionPerformerRole(numTerm int, numAction int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].performerRole", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].performerRole", resource.Term[numTerm].Action[numAction].PerformerRole, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermActionPerformerLinkId(numTerm int, numAction int, numPerformerLinkId int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numPerformerLinkId >= len(resource.Term[numTerm].Action[numAction].PerformerLinkId) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].performerLinkId["+strconv.Itoa(numPerformerLinkId)+"]", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].performerLinkId["+strconv.Itoa(numPerformerLinkId)+"]", &resource.Term[numTerm].Action[numAction].PerformerLinkId[numPerformerLinkId], htmlAttrs)
}
func (resource *Contract) T_TermActionReasonCode(numTerm int, numAction int, numReasonCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numReasonCode >= len(resource.Term[numTerm].Action[numAction].ReasonCode) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].reasonCode["+strconv.Itoa(numReasonCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].reasonCode["+strconv.Itoa(numReasonCode)+"]", &resource.Term[numTerm].Action[numAction].ReasonCode[numReasonCode], optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_TermActionReason(numTerm int, numAction int, numReason int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numReason >= len(resource.Term[numTerm].Action[numAction].Reason) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].reason["+strconv.Itoa(numReason)+"]", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].reason["+strconv.Itoa(numReason)+"]", &resource.Term[numTerm].Action[numAction].Reason[numReason], htmlAttrs)
}
func (resource *Contract) T_TermActionReasonLinkId(numTerm int, numAction int, numReasonLinkId int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numReasonLinkId >= len(resource.Term[numTerm].Action[numAction].ReasonLinkId) {
		return StringInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].reasonLinkId["+strconv.Itoa(numReasonLinkId)+"]", nil, htmlAttrs)
	}
	return StringInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].reasonLinkId["+strconv.Itoa(numReasonLinkId)+"]", &resource.Term[numTerm].Action[numAction].ReasonLinkId[numReasonLinkId], htmlAttrs)
}
func (resource *Contract) T_TermActionNote(numTerm int, numAction int, numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numNote >= len(resource.Term[numTerm].Action[numAction].Note) {
		return AnnotationTextArea("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].note["+strconv.Itoa(numNote)+"]", &resource.Term[numTerm].Action[numAction].Note[numNote], htmlAttrs)
}
func (resource *Contract) T_TermActionSecurityLabelNumber(numTerm int, numAction int, numSecurityLabelNumber int, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numSecurityLabelNumber >= len(resource.Term[numTerm].Action[numAction].SecurityLabelNumber) {
		return IntInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].securityLabelNumber["+strconv.Itoa(numSecurityLabelNumber)+"]", nil, htmlAttrs)
	}
	return IntInput("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].securityLabelNumber["+strconv.Itoa(numSecurityLabelNumber)+"]", &resource.Term[numTerm].Action[numAction].SecurityLabelNumber[numSecurityLabelNumber], htmlAttrs)
}
func (resource *Contract) T_TermActionSubjectRole(numTerm int, numAction int, numSubject int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numTerm >= len(resource.Term) || numAction >= len(resource.Term[numTerm].Action) || numSubject >= len(resource.Term[numTerm].Action[numAction].Subject) {
		return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].subject["+strconv.Itoa(numSubject)+"].role", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("term["+strconv.Itoa(numTerm)+"].action["+strconv.Itoa(numAction)+"].subject["+strconv.Itoa(numSubject)+"].role", resource.Term[numTerm].Action[numAction].Subject[numSubject].Role, optionsValueSet, htmlAttrs)
}
func (resource *Contract) T_SignerType(numSigner int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSigner >= len(resource.Signer) {
		return CodingSelect("signer["+strconv.Itoa(numSigner)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("signer["+strconv.Itoa(numSigner)+"].type", &resource.Signer[numSigner].Type, optionsValueSet, htmlAttrs)
}
