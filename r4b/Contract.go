package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	TopicCodeableConcept     *CodeableConcept           `json:"topicCodeableConcept"`
	TopicReference           *Reference                 `json:"topicReference"`
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
	LegallyBindingAttachment *Attachment                `json:"legallyBindingAttachment"`
	LegallyBindingReference  *Reference                 `json:"legallyBindingReference"`
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
	TopicCodeableConcept *CodeableConcept            `json:"topicCodeableConcept"`
	TopicReference       *Reference                  `json:"topicReference"`
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
	EntityCodeableConcept *CodeableConcept `json:"entityCodeableConcept"`
	EntityReference       *Reference       `json:"entityReference"`
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
	OccurrenceDateTime  *string                     `json:"occurrenceDateTime"`
	OccurrencePeriod    *Period                     `json:"occurrencePeriod"`
	OccurrenceTiming    *Timing                     `json:"occurrenceTiming"`
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

func (resource *Contract) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Contract) T_Status() templ.Component {
	optionsValueSet := VSContract_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", resource.Status, optionsValueSet)
}
func (resource *Contract) T_LegalState(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("legalState", nil, optionsValueSet)
	}
	return CodeableConceptSelect("legalState", resource.LegalState, optionsValueSet)
}
func (resource *Contract) T_ContentDerivative(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("contentDerivative", nil, optionsValueSet)
	}
	return CodeableConceptSelect("contentDerivative", resource.ContentDerivative, optionsValueSet)
}
func (resource *Contract) T_ExpirationType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("expirationType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("expirationType", resource.ExpirationType, optionsValueSet)
}
func (resource *Contract) T_Scope(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("scope", nil, optionsValueSet)
	}
	return CodeableConceptSelect("scope", resource.Scope, optionsValueSet)
}
func (resource *Contract) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *Contract) T_SubType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("subType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subType", &resource.SubType[0], optionsValueSet)
}
func (resource *Contract) T_ContentDefinitionType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.ContentDefinition.Type, optionsValueSet)
}
func (resource *Contract) T_ContentDefinitionSubType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("subType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subType", resource.ContentDefinition.SubType, optionsValueSet)
}
func (resource *Contract) T_ContentDefinitionPublicationStatus() templ.Component {
	optionsValueSet := VSContract_publicationstatus

	if resource == nil {
		return CodeSelect("publicationStatus", nil, optionsValueSet)
	}
	return CodeSelect("publicationStatus", &resource.ContentDefinition.PublicationStatus, optionsValueSet)
}
func (resource *Contract) T_TermType(numTerm int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term) >= numTerm {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Term[numTerm].Type, optionsValueSet)
}
func (resource *Contract) T_TermSubType(numTerm int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term) >= numTerm {
		return CodeableConceptSelect("subType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subType", resource.Term[numTerm].SubType, optionsValueSet)
}
func (resource *Contract) T_TermSecurityLabelClassification(numTerm int, numSecurityLabel int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].SecurityLabel) >= numSecurityLabel {
		return CodingSelect("classification", nil, optionsValueSet)
	}
	return CodingSelect("classification", &resource.Term[numTerm].SecurityLabel[numSecurityLabel].Classification, optionsValueSet)
}
func (resource *Contract) T_TermSecurityLabelCategory(numTerm int, numSecurityLabel int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].SecurityLabel) >= numSecurityLabel {
		return CodingSelect("category", nil, optionsValueSet)
	}
	return CodingSelect("category", &resource.Term[numTerm].SecurityLabel[numSecurityLabel].Category[0], optionsValueSet)
}
func (resource *Contract) T_TermSecurityLabelControl(numTerm int, numSecurityLabel int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].SecurityLabel) >= numSecurityLabel {
		return CodingSelect("control", nil, optionsValueSet)
	}
	return CodingSelect("control", &resource.Term[numTerm].SecurityLabel[numSecurityLabel].Control[0], optionsValueSet)
}
func (resource *Contract) T_TermOfferType(numTerm int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term) >= numTerm {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Term[numTerm].Offer.Type, optionsValueSet)
}
func (resource *Contract) T_TermOfferDecision(numTerm int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term) >= numTerm {
		return CodeableConceptSelect("decision", nil, optionsValueSet)
	}
	return CodeableConceptSelect("decision", resource.Term[numTerm].Offer.Decision, optionsValueSet)
}
func (resource *Contract) T_TermOfferDecisionMode(numTerm int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term) >= numTerm {
		return CodeableConceptSelect("decisionMode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("decisionMode", &resource.Term[numTerm].Offer.DecisionMode[0], optionsValueSet)
}
func (resource *Contract) T_TermOfferPartyRole(numTerm int, numParty int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].Offer.Party) >= numParty {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", &resource.Term[numTerm].Offer.Party[numParty].Role, optionsValueSet)
}
func (resource *Contract) T_TermAssetScope(numTerm int, numAsset int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].Asset) >= numAsset {
		return CodeableConceptSelect("scope", nil, optionsValueSet)
	}
	return CodeableConceptSelect("scope", resource.Term[numTerm].Asset[numAsset].Scope, optionsValueSet)
}
func (resource *Contract) T_TermAssetType(numTerm int, numAsset int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].Asset) >= numAsset {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Term[numTerm].Asset[numAsset].Type[0], optionsValueSet)
}
func (resource *Contract) T_TermAssetSubtype(numTerm int, numAsset int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].Asset) >= numAsset {
		return CodeableConceptSelect("subtype", nil, optionsValueSet)
	}
	return CodeableConceptSelect("subtype", &resource.Term[numTerm].Asset[numAsset].Subtype[0], optionsValueSet)
}
func (resource *Contract) T_TermAssetRelationship(numTerm int, numAsset int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].Asset) >= numAsset {
		return CodingSelect("relationship", nil, optionsValueSet)
	}
	return CodingSelect("relationship", resource.Term[numTerm].Asset[numAsset].Relationship, optionsValueSet)
}
func (resource *Contract) T_TermAssetPeriodType(numTerm int, numAsset int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].Asset) >= numAsset {
		return CodeableConceptSelect("periodType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("periodType", &resource.Term[numTerm].Asset[numAsset].PeriodType[0], optionsValueSet)
}
func (resource *Contract) T_TermAssetContextCode(numTerm int, numAsset int, numContext int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].Asset[numAsset].Context) >= numContext {
		return CodeableConceptSelect("code", nil, optionsValueSet)
	}
	return CodeableConceptSelect("code", &resource.Term[numTerm].Asset[numAsset].Context[numContext].Code[0], optionsValueSet)
}
func (resource *Contract) T_TermActionType(numTerm int, numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].Action) >= numAction {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Term[numTerm].Action[numAction].Type, optionsValueSet)
}
func (resource *Contract) T_TermActionIntent(numTerm int, numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].Action) >= numAction {
		return CodeableConceptSelect("intent", nil, optionsValueSet)
	}
	return CodeableConceptSelect("intent", &resource.Term[numTerm].Action[numAction].Intent, optionsValueSet)
}
func (resource *Contract) T_TermActionStatus(numTerm int, numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].Action) >= numAction {
		return CodeableConceptSelect("status", nil, optionsValueSet)
	}
	return CodeableConceptSelect("status", &resource.Term[numTerm].Action[numAction].Status, optionsValueSet)
}
func (resource *Contract) T_TermActionPerformerType(numTerm int, numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].Action) >= numAction {
		return CodeableConceptSelect("performerType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("performerType", &resource.Term[numTerm].Action[numAction].PerformerType[0], optionsValueSet)
}
func (resource *Contract) T_TermActionPerformerRole(numTerm int, numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].Action) >= numAction {
		return CodeableConceptSelect("performerRole", nil, optionsValueSet)
	}
	return CodeableConceptSelect("performerRole", resource.Term[numTerm].Action[numAction].PerformerRole, optionsValueSet)
}
func (resource *Contract) T_TermActionReasonCode(numTerm int, numAction int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].Action) >= numAction {
		return CodeableConceptSelect("reasonCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("reasonCode", &resource.Term[numTerm].Action[numAction].ReasonCode[0], optionsValueSet)
}
func (resource *Contract) T_TermActionSubjectRole(numTerm int, numAction int, numSubject int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Term[numTerm].Action[numAction].Subject) >= numSubject {
		return CodeableConceptSelect("role", nil, optionsValueSet)
	}
	return CodeableConceptSelect("role", resource.Term[numTerm].Action[numAction].Subject[numSubject].Role, optionsValueSet)
}
func (resource *Contract) T_SignerType(numSigner int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Signer) >= numSigner {
		return CodingSelect("type", nil, optionsValueSet)
	}
	return CodingSelect("type", &resource.Signer[numSigner].Type, optionsValueSet)
}
