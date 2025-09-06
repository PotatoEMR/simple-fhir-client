package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinition struct {
	Id                        *string                                `json:"id,omitempty"`
	Meta                      *Meta                                  `json:"meta,omitempty"`
	ImplicitRules             *string                                `json:"implicitRules,omitempty"`
	Language                  *string                                `json:"language,omitempty"`
	Text                      *Narrative                             `json:"text,omitempty"`
	Contained                 []Resource                             `json:"contained,omitempty"`
	Extension                 []Extension                            `json:"extension,omitempty"`
	ModifierExtension         []Extension                            `json:"modifierExtension,omitempty"`
	Description               *string                                `json:"description,omitempty"`
	Identifier                []Identifier                           `json:"identifier,omitempty"`
	UdiDeviceIdentifier       []DeviceDefinitionUdiDeviceIdentifier  `json:"udiDeviceIdentifier,omitempty"`
	RegulatoryIdentifier      []DeviceDefinitionRegulatoryIdentifier `json:"regulatoryIdentifier,omitempty"`
	PartNumber                *string                                `json:"partNumber,omitempty"`
	Manufacturer              *Reference                             `json:"manufacturer,omitempty"`
	DeviceName                []DeviceDefinitionDeviceName           `json:"deviceName,omitempty"`
	ModelNumber               *string                                `json:"modelNumber,omitempty"`
	Classification            []DeviceDefinitionClassification       `json:"classification,omitempty"`
	ConformsTo                []DeviceDefinitionConformsTo           `json:"conformsTo,omitempty"`
	HasPart                   []DeviceDefinitionHasPart              `json:"hasPart,omitempty"`
	Packaging                 []DeviceDefinitionPackaging            `json:"packaging,omitempty"`
	Version                   []DeviceDefinitionVersion              `json:"version,omitempty"`
	Safety                    []CodeableConcept                      `json:"safety,omitempty"`
	ShelfLifeStorage          []ProductShelfLife                     `json:"shelfLifeStorage,omitempty"`
	LanguageCode              []CodeableConcept                      `json:"languageCode,omitempty"`
	Property                  []DeviceDefinitionProperty             `json:"property,omitempty"`
	Owner                     *Reference                             `json:"owner,omitempty"`
	Contact                   []ContactPoint                         `json:"contact,omitempty"`
	Link                      []DeviceDefinitionLink                 `json:"link,omitempty"`
	Note                      []Annotation                           `json:"note,omitempty"`
	Material                  []DeviceDefinitionMaterial             `json:"material,omitempty"`
	ProductionIdentifierInUDI []string                               `json:"productionIdentifierInUDI,omitempty"`
	Guideline                 *DeviceDefinitionGuideline             `json:"guideline,omitempty"`
	CorrectiveAction          *DeviceDefinitionCorrectiveAction      `json:"correctiveAction,omitempty"`
	ChargeItem                []DeviceDefinitionChargeItem           `json:"chargeItem,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinitionUdiDeviceIdentifier struct {
	Id                 *string                                                 `json:"id,omitempty"`
	Extension          []Extension                                             `json:"extension,omitempty"`
	ModifierExtension  []Extension                                             `json:"modifierExtension,omitempty"`
	DeviceIdentifier   string                                                  `json:"deviceIdentifier"`
	Issuer             string                                                  `json:"issuer"`
	Jurisdiction       string                                                  `json:"jurisdiction"`
	MarketDistribution []DeviceDefinitionUdiDeviceIdentifierMarketDistribution `json:"marketDistribution,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinitionUdiDeviceIdentifierMarketDistribution struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	MarketPeriod      Period      `json:"marketPeriod"`
	SubJurisdiction   string      `json:"subJurisdiction"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinitionRegulatoryIdentifier struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Type              string      `json:"type"`
	DeviceIdentifier  string      `json:"deviceIdentifier"`
	Issuer            string      `json:"issuer"`
	Jurisdiction      string      `json:"jurisdiction"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinitionDeviceName struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Type              string      `json:"type"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinitionClassification struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	Justification     []RelatedArtifact `json:"justification,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinitionConformsTo struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Category          *CodeableConcept  `json:"category,omitempty"`
	Specification     CodeableConcept   `json:"specification"`
	Version           []string          `json:"version,omitempty"`
	Source            []RelatedArtifact `json:"source,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinitionHasPart struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Reference         Reference   `json:"reference"`
	Count             *int        `json:"count,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinitionPackaging struct {
	Id                *string                                `json:"id,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Identifier        *Identifier                            `json:"identifier,omitempty"`
	Type              *CodeableConcept                       `json:"type,omitempty"`
	Count             *int                                   `json:"count,omitempty"`
	Distributor       []DeviceDefinitionPackagingDistributor `json:"distributor,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinitionPackagingDistributor struct {
	Id                    *string     `json:"id,omitempty"`
	Extension             []Extension `json:"extension,omitempty"`
	ModifierExtension     []Extension `json:"modifierExtension,omitempty"`
	Name                  *string     `json:"name,omitempty"`
	OrganizationReference []Reference `json:"organizationReference,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinitionVersion struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Component         *Identifier      `json:"component,omitempty"`
	Value             string           `json:"value"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinitionProperty struct {
	Id                   *string         `json:"id,omitempty"`
	Extension            []Extension     `json:"extension,omitempty"`
	ModifierExtension    []Extension     `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept `json:"type"`
	ValueQuantity        Quantity        `json:"valueQuantity"`
	ValueCodeableConcept CodeableConcept `json:"valueCodeableConcept"`
	ValueString          string          `json:"valueString"`
	ValueBoolean         bool            `json:"valueBoolean"`
	ValueInteger         int             `json:"valueInteger"`
	ValueRange           Range           `json:"valueRange"`
	ValueAttachment      Attachment      `json:"valueAttachment"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinitionLink struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Relation          Coding            `json:"relation"`
	RelatedDevice     CodeableReference `json:"relatedDevice"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinitionMaterial struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	Substance           CodeableConcept `json:"substance"`
	Alternate           *bool           `json:"alternate,omitempty"`
	AllergenicIndicator *bool           `json:"allergenicIndicator,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinitionGuideline struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	UseContext        []UsageContext    `json:"useContext,omitempty"`
	UsageInstruction  *string           `json:"usageInstruction,omitempty"`
	RelatedArtifact   []RelatedArtifact `json:"relatedArtifact,omitempty"`
	Indication        []CodeableConcept `json:"indication,omitempty"`
	Contraindication  []CodeableConcept `json:"contraindication,omitempty"`
	Warning           []CodeableConcept `json:"warning,omitempty"`
	IntendedUse       *string           `json:"intendedUse,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinitionCorrectiveAction struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Recall            bool        `json:"recall"`
	Scope             *string     `json:"scope,omitempty"`
	Period            Period      `json:"period"`
}

// http://hl7.org/fhir/r5/StructureDefinition/DeviceDefinition
type DeviceDefinitionChargeItem struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	ChargeItemCode    CodeableReference `json:"chargeItemCode"`
	Count             Quantity          `json:"count"`
	EffectivePeriod   *Period           `json:"effectivePeriod,omitempty"`
	UseContext        []UsageContext    `json:"useContext,omitempty"`
}

type OtherDeviceDefinition DeviceDefinition

// on convert struct to json, automatically add resourceType=DeviceDefinition
func (r DeviceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceDefinition: OtherDeviceDefinition(r),
		ResourceType:          "DeviceDefinition",
	})
}

func (resource *DeviceDefinition) T_Id() templ.Component {

	if resource == nil {
		return StringInput("DeviceDefinition.Id", nil)
	}
	return StringInput("DeviceDefinition.Id", resource.Id)
}
func (resource *DeviceDefinition) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("DeviceDefinition.ImplicitRules", nil)
	}
	return StringInput("DeviceDefinition.ImplicitRules", resource.ImplicitRules)
}
func (resource *DeviceDefinition) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("DeviceDefinition.Language", nil, optionsValueSet)
	}
	return CodeSelect("DeviceDefinition.Language", resource.Language, optionsValueSet)
}
func (resource *DeviceDefinition) T_Description() templ.Component {

	if resource == nil {
		return StringInput("DeviceDefinition.Description", nil)
	}
	return StringInput("DeviceDefinition.Description", resource.Description)
}
func (resource *DeviceDefinition) T_PartNumber() templ.Component {

	if resource == nil {
		return StringInput("DeviceDefinition.PartNumber", nil)
	}
	return StringInput("DeviceDefinition.PartNumber", resource.PartNumber)
}
func (resource *DeviceDefinition) T_ModelNumber() templ.Component {

	if resource == nil {
		return StringInput("DeviceDefinition.ModelNumber", nil)
	}
	return StringInput("DeviceDefinition.ModelNumber", resource.ModelNumber)
}
func (resource *DeviceDefinition) T_Safety(numSafety int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Safety) >= numSafety {
		return CodeableConceptSelect("DeviceDefinition.Safety["+strconv.Itoa(numSafety)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDefinition.Safety["+strconv.Itoa(numSafety)+"]", &resource.Safety[numSafety], optionsValueSet)
}
func (resource *DeviceDefinition) T_LanguageCode(numLanguageCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.LanguageCode) >= numLanguageCode {
		return CodeableConceptSelect("DeviceDefinition.LanguageCode["+strconv.Itoa(numLanguageCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDefinition.LanguageCode["+strconv.Itoa(numLanguageCode)+"]", &resource.LanguageCode[numLanguageCode], optionsValueSet)
}
func (resource *DeviceDefinition) T_ProductionIdentifierInUDI(numProductionIdentifierInUDI int) templ.Component {
	optionsValueSet := VSDevice_productidentifierinudi

	if resource == nil || len(resource.ProductionIdentifierInUDI) >= numProductionIdentifierInUDI {
		return CodeSelect("DeviceDefinition.ProductionIdentifierInUDI["+strconv.Itoa(numProductionIdentifierInUDI)+"]", nil, optionsValueSet)
	}
	return CodeSelect("DeviceDefinition.ProductionIdentifierInUDI["+strconv.Itoa(numProductionIdentifierInUDI)+"]", &resource.ProductionIdentifierInUDI[numProductionIdentifierInUDI], optionsValueSet)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierId(numUdiDeviceIdentifier int) templ.Component {

	if resource == nil || len(resource.UdiDeviceIdentifier) >= numUdiDeviceIdentifier {
		return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].Id", nil)
	}
	return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].Id", resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].Id)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierDeviceIdentifier(numUdiDeviceIdentifier int) templ.Component {

	if resource == nil || len(resource.UdiDeviceIdentifier) >= numUdiDeviceIdentifier {
		return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].DeviceIdentifier", nil)
	}
	return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].DeviceIdentifier", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].DeviceIdentifier)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierIssuer(numUdiDeviceIdentifier int) templ.Component {

	if resource == nil || len(resource.UdiDeviceIdentifier) >= numUdiDeviceIdentifier {
		return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].Issuer", nil)
	}
	return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].Issuer", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].Issuer)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierJurisdiction(numUdiDeviceIdentifier int) templ.Component {

	if resource == nil || len(resource.UdiDeviceIdentifier) >= numUdiDeviceIdentifier {
		return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].Jurisdiction", nil)
	}
	return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].Jurisdiction", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].Jurisdiction)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierMarketDistributionId(numUdiDeviceIdentifier int, numMarketDistribution int) templ.Component {

	if resource == nil || len(resource.UdiDeviceIdentifier) >= numUdiDeviceIdentifier || len(resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].MarketDistribution) >= numMarketDistribution {
		return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].MarketDistribution["+strconv.Itoa(numMarketDistribution)+"].Id", nil)
	}
	return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].MarketDistribution["+strconv.Itoa(numMarketDistribution)+"].Id", resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].MarketDistribution[numMarketDistribution].Id)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierMarketDistributionSubJurisdiction(numUdiDeviceIdentifier int, numMarketDistribution int) templ.Component {

	if resource == nil || len(resource.UdiDeviceIdentifier) >= numUdiDeviceIdentifier || len(resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].MarketDistribution) >= numMarketDistribution {
		return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].MarketDistribution["+strconv.Itoa(numMarketDistribution)+"].SubJurisdiction", nil)
	}
	return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].MarketDistribution["+strconv.Itoa(numMarketDistribution)+"].SubJurisdiction", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].MarketDistribution[numMarketDistribution].SubJurisdiction)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierId(numRegulatoryIdentifier int) templ.Component {

	if resource == nil || len(resource.RegulatoryIdentifier) >= numRegulatoryIdentifier {
		return StringInput("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].Id", nil)
	}
	return StringInput("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].Id", resource.RegulatoryIdentifier[numRegulatoryIdentifier].Id)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierType(numRegulatoryIdentifier int) templ.Component {
	optionsValueSet := VSDevicedefinition_regulatory_identifier_type

	if resource == nil || len(resource.RegulatoryIdentifier) >= numRegulatoryIdentifier {
		return CodeSelect("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].Type", &resource.RegulatoryIdentifier[numRegulatoryIdentifier].Type, optionsValueSet)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierDeviceIdentifier(numRegulatoryIdentifier int) templ.Component {

	if resource == nil || len(resource.RegulatoryIdentifier) >= numRegulatoryIdentifier {
		return StringInput("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].DeviceIdentifier", nil)
	}
	return StringInput("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].DeviceIdentifier", &resource.RegulatoryIdentifier[numRegulatoryIdentifier].DeviceIdentifier)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierIssuer(numRegulatoryIdentifier int) templ.Component {

	if resource == nil || len(resource.RegulatoryIdentifier) >= numRegulatoryIdentifier {
		return StringInput("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].Issuer", nil)
	}
	return StringInput("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].Issuer", &resource.RegulatoryIdentifier[numRegulatoryIdentifier].Issuer)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierJurisdiction(numRegulatoryIdentifier int) templ.Component {

	if resource == nil || len(resource.RegulatoryIdentifier) >= numRegulatoryIdentifier {
		return StringInput("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].Jurisdiction", nil)
	}
	return StringInput("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].Jurisdiction", &resource.RegulatoryIdentifier[numRegulatoryIdentifier].Jurisdiction)
}
func (resource *DeviceDefinition) T_DeviceNameId(numDeviceName int) templ.Component {

	if resource == nil || len(resource.DeviceName) >= numDeviceName {
		return StringInput("DeviceDefinition.DeviceName["+strconv.Itoa(numDeviceName)+"].Id", nil)
	}
	return StringInput("DeviceDefinition.DeviceName["+strconv.Itoa(numDeviceName)+"].Id", resource.DeviceName[numDeviceName].Id)
}
func (resource *DeviceDefinition) T_DeviceNameName(numDeviceName int) templ.Component {

	if resource == nil || len(resource.DeviceName) >= numDeviceName {
		return StringInput("DeviceDefinition.DeviceName["+strconv.Itoa(numDeviceName)+"].Name", nil)
	}
	return StringInput("DeviceDefinition.DeviceName["+strconv.Itoa(numDeviceName)+"].Name", &resource.DeviceName[numDeviceName].Name)
}
func (resource *DeviceDefinition) T_DeviceNameType(numDeviceName int) templ.Component {
	optionsValueSet := VSDevice_nametype

	if resource == nil || len(resource.DeviceName) >= numDeviceName {
		return CodeSelect("DeviceDefinition.DeviceName["+strconv.Itoa(numDeviceName)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("DeviceDefinition.DeviceName["+strconv.Itoa(numDeviceName)+"].Type", &resource.DeviceName[numDeviceName].Type, optionsValueSet)
}
func (resource *DeviceDefinition) T_ClassificationId(numClassification int) templ.Component {

	if resource == nil || len(resource.Classification) >= numClassification {
		return StringInput("DeviceDefinition.Classification["+strconv.Itoa(numClassification)+"].Id", nil)
	}
	return StringInput("DeviceDefinition.Classification["+strconv.Itoa(numClassification)+"].Id", resource.Classification[numClassification].Id)
}
func (resource *DeviceDefinition) T_ClassificationType(numClassification int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Classification) >= numClassification {
		return CodeableConceptSelect("DeviceDefinition.Classification["+strconv.Itoa(numClassification)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDefinition.Classification["+strconv.Itoa(numClassification)+"].Type", &resource.Classification[numClassification].Type, optionsValueSet)
}
func (resource *DeviceDefinition) T_ConformsToId(numConformsTo int) templ.Component {

	if resource == nil || len(resource.ConformsTo) >= numConformsTo {
		return StringInput("DeviceDefinition.ConformsTo["+strconv.Itoa(numConformsTo)+"].Id", nil)
	}
	return StringInput("DeviceDefinition.ConformsTo["+strconv.Itoa(numConformsTo)+"].Id", resource.ConformsTo[numConformsTo].Id)
}
func (resource *DeviceDefinition) T_ConformsToCategory(numConformsTo int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ConformsTo) >= numConformsTo {
		return CodeableConceptSelect("DeviceDefinition.ConformsTo["+strconv.Itoa(numConformsTo)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDefinition.ConformsTo["+strconv.Itoa(numConformsTo)+"].Category", resource.ConformsTo[numConformsTo].Category, optionsValueSet)
}
func (resource *DeviceDefinition) T_ConformsToSpecification(numConformsTo int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ConformsTo) >= numConformsTo {
		return CodeableConceptSelect("DeviceDefinition.ConformsTo["+strconv.Itoa(numConformsTo)+"].Specification", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDefinition.ConformsTo["+strconv.Itoa(numConformsTo)+"].Specification", &resource.ConformsTo[numConformsTo].Specification, optionsValueSet)
}
func (resource *DeviceDefinition) T_ConformsToVersion(numConformsTo int, numVersion int) templ.Component {

	if resource == nil || len(resource.ConformsTo) >= numConformsTo || len(resource.ConformsTo[numConformsTo].Version) >= numVersion {
		return StringInput("DeviceDefinition.ConformsTo["+strconv.Itoa(numConformsTo)+"].Version["+strconv.Itoa(numVersion)+"]", nil)
	}
	return StringInput("DeviceDefinition.ConformsTo["+strconv.Itoa(numConformsTo)+"].Version["+strconv.Itoa(numVersion)+"]", &resource.ConformsTo[numConformsTo].Version[numVersion])
}
func (resource *DeviceDefinition) T_HasPartId(numHasPart int) templ.Component {

	if resource == nil || len(resource.HasPart) >= numHasPart {
		return StringInput("DeviceDefinition.HasPart["+strconv.Itoa(numHasPart)+"].Id", nil)
	}
	return StringInput("DeviceDefinition.HasPart["+strconv.Itoa(numHasPart)+"].Id", resource.HasPart[numHasPart].Id)
}
func (resource *DeviceDefinition) T_HasPartCount(numHasPart int) templ.Component {

	if resource == nil || len(resource.HasPart) >= numHasPart {
		return IntInput("DeviceDefinition.HasPart["+strconv.Itoa(numHasPart)+"].Count", nil)
	}
	return IntInput("DeviceDefinition.HasPart["+strconv.Itoa(numHasPart)+"].Count", resource.HasPart[numHasPart].Count)
}
func (resource *DeviceDefinition) T_PackagingId(numPackaging int) templ.Component {

	if resource == nil || len(resource.Packaging) >= numPackaging {
		return StringInput("DeviceDefinition.Packaging["+strconv.Itoa(numPackaging)+"].Id", nil)
	}
	return StringInput("DeviceDefinition.Packaging["+strconv.Itoa(numPackaging)+"].Id", resource.Packaging[numPackaging].Id)
}
func (resource *DeviceDefinition) T_PackagingType(numPackaging int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Packaging) >= numPackaging {
		return CodeableConceptSelect("DeviceDefinition.Packaging["+strconv.Itoa(numPackaging)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDefinition.Packaging["+strconv.Itoa(numPackaging)+"].Type", resource.Packaging[numPackaging].Type, optionsValueSet)
}
func (resource *DeviceDefinition) T_PackagingCount(numPackaging int) templ.Component {

	if resource == nil || len(resource.Packaging) >= numPackaging {
		return IntInput("DeviceDefinition.Packaging["+strconv.Itoa(numPackaging)+"].Count", nil)
	}
	return IntInput("DeviceDefinition.Packaging["+strconv.Itoa(numPackaging)+"].Count", resource.Packaging[numPackaging].Count)
}
func (resource *DeviceDefinition) T_PackagingDistributorId(numPackaging int, numDistributor int) templ.Component {

	if resource == nil || len(resource.Packaging) >= numPackaging || len(resource.Packaging[numPackaging].Distributor) >= numDistributor {
		return StringInput("DeviceDefinition.Packaging["+strconv.Itoa(numPackaging)+"].Distributor["+strconv.Itoa(numDistributor)+"].Id", nil)
	}
	return StringInput("DeviceDefinition.Packaging["+strconv.Itoa(numPackaging)+"].Distributor["+strconv.Itoa(numDistributor)+"].Id", resource.Packaging[numPackaging].Distributor[numDistributor].Id)
}
func (resource *DeviceDefinition) T_PackagingDistributorName(numPackaging int, numDistributor int) templ.Component {

	if resource == nil || len(resource.Packaging) >= numPackaging || len(resource.Packaging[numPackaging].Distributor) >= numDistributor {
		return StringInput("DeviceDefinition.Packaging["+strconv.Itoa(numPackaging)+"].Distributor["+strconv.Itoa(numDistributor)+"].Name", nil)
	}
	return StringInput("DeviceDefinition.Packaging["+strconv.Itoa(numPackaging)+"].Distributor["+strconv.Itoa(numDistributor)+"].Name", resource.Packaging[numPackaging].Distributor[numDistributor].Name)
}
func (resource *DeviceDefinition) T_VersionId(numVersion int) templ.Component {

	if resource == nil || len(resource.Version) >= numVersion {
		return StringInput("DeviceDefinition.Version["+strconv.Itoa(numVersion)+"].Id", nil)
	}
	return StringInput("DeviceDefinition.Version["+strconv.Itoa(numVersion)+"].Id", resource.Version[numVersion].Id)
}
func (resource *DeviceDefinition) T_VersionType(numVersion int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Version) >= numVersion {
		return CodeableConceptSelect("DeviceDefinition.Version["+strconv.Itoa(numVersion)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDefinition.Version["+strconv.Itoa(numVersion)+"].Type", resource.Version[numVersion].Type, optionsValueSet)
}
func (resource *DeviceDefinition) T_VersionValue(numVersion int) templ.Component {

	if resource == nil || len(resource.Version) >= numVersion {
		return StringInput("DeviceDefinition.Version["+strconv.Itoa(numVersion)+"].Value", nil)
	}
	return StringInput("DeviceDefinition.Version["+strconv.Itoa(numVersion)+"].Value", &resource.Version[numVersion].Value)
}
func (resource *DeviceDefinition) T_PropertyId(numProperty int) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return StringInput("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].Id", nil)
	}
	return StringInput("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].Id", resource.Property[numProperty].Id)
}
func (resource *DeviceDefinition) T_PropertyType(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return CodeableConceptSelect("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].Type", &resource.Property[numProperty].Type, optionsValueSet)
}
func (resource *DeviceDefinition) T_LinkId(numLink int) templ.Component {

	if resource == nil || len(resource.Link) >= numLink {
		return StringInput("DeviceDefinition.Link["+strconv.Itoa(numLink)+"].Id", nil)
	}
	return StringInput("DeviceDefinition.Link["+strconv.Itoa(numLink)+"].Id", resource.Link[numLink].Id)
}
func (resource *DeviceDefinition) T_LinkRelation(numLink int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Link) >= numLink {
		return CodingSelect("DeviceDefinition.Link["+strconv.Itoa(numLink)+"].Relation", nil, optionsValueSet)
	}
	return CodingSelect("DeviceDefinition.Link["+strconv.Itoa(numLink)+"].Relation", &resource.Link[numLink].Relation, optionsValueSet)
}
func (resource *DeviceDefinition) T_MaterialId(numMaterial int) templ.Component {

	if resource == nil || len(resource.Material) >= numMaterial {
		return StringInput("DeviceDefinition.Material["+strconv.Itoa(numMaterial)+"].Id", nil)
	}
	return StringInput("DeviceDefinition.Material["+strconv.Itoa(numMaterial)+"].Id", resource.Material[numMaterial].Id)
}
func (resource *DeviceDefinition) T_MaterialSubstance(numMaterial int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Material) >= numMaterial {
		return CodeableConceptSelect("DeviceDefinition.Material["+strconv.Itoa(numMaterial)+"].Substance", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDefinition.Material["+strconv.Itoa(numMaterial)+"].Substance", &resource.Material[numMaterial].Substance, optionsValueSet)
}
func (resource *DeviceDefinition) T_MaterialAlternate(numMaterial int) templ.Component {

	if resource == nil || len(resource.Material) >= numMaterial {
		return BoolInput("DeviceDefinition.Material["+strconv.Itoa(numMaterial)+"].Alternate", nil)
	}
	return BoolInput("DeviceDefinition.Material["+strconv.Itoa(numMaterial)+"].Alternate", resource.Material[numMaterial].Alternate)
}
func (resource *DeviceDefinition) T_MaterialAllergenicIndicator(numMaterial int) templ.Component {

	if resource == nil || len(resource.Material) >= numMaterial {
		return BoolInput("DeviceDefinition.Material["+strconv.Itoa(numMaterial)+"].AllergenicIndicator", nil)
	}
	return BoolInput("DeviceDefinition.Material["+strconv.Itoa(numMaterial)+"].AllergenicIndicator", resource.Material[numMaterial].AllergenicIndicator)
}
func (resource *DeviceDefinition) T_GuidelineId() templ.Component {

	if resource == nil {
		return StringInput("DeviceDefinition.Guideline.Id", nil)
	}
	return StringInput("DeviceDefinition.Guideline.Id", resource.Guideline.Id)
}
func (resource *DeviceDefinition) T_GuidelineUsageInstruction() templ.Component {

	if resource == nil {
		return StringInput("DeviceDefinition.Guideline.UsageInstruction", nil)
	}
	return StringInput("DeviceDefinition.Guideline.UsageInstruction", resource.Guideline.UsageInstruction)
}
func (resource *DeviceDefinition) T_GuidelineIndication(numIndication int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Guideline.Indication) >= numIndication {
		return CodeableConceptSelect("DeviceDefinition.Guideline.Indication["+strconv.Itoa(numIndication)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDefinition.Guideline.Indication["+strconv.Itoa(numIndication)+"]", &resource.Guideline.Indication[numIndication], optionsValueSet)
}
func (resource *DeviceDefinition) T_GuidelineContraindication(numContraindication int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Guideline.Contraindication) >= numContraindication {
		return CodeableConceptSelect("DeviceDefinition.Guideline.Contraindication["+strconv.Itoa(numContraindication)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDefinition.Guideline.Contraindication["+strconv.Itoa(numContraindication)+"]", &resource.Guideline.Contraindication[numContraindication], optionsValueSet)
}
func (resource *DeviceDefinition) T_GuidelineWarning(numWarning int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Guideline.Warning) >= numWarning {
		return CodeableConceptSelect("DeviceDefinition.Guideline.Warning["+strconv.Itoa(numWarning)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDefinition.Guideline.Warning["+strconv.Itoa(numWarning)+"]", &resource.Guideline.Warning[numWarning], optionsValueSet)
}
func (resource *DeviceDefinition) T_GuidelineIntendedUse() templ.Component {

	if resource == nil {
		return StringInput("DeviceDefinition.Guideline.IntendedUse", nil)
	}
	return StringInput("DeviceDefinition.Guideline.IntendedUse", resource.Guideline.IntendedUse)
}
func (resource *DeviceDefinition) T_CorrectiveActionId() templ.Component {

	if resource == nil {
		return StringInput("DeviceDefinition.CorrectiveAction.Id", nil)
	}
	return StringInput("DeviceDefinition.CorrectiveAction.Id", resource.CorrectiveAction.Id)
}
func (resource *DeviceDefinition) T_CorrectiveActionRecall() templ.Component {

	if resource == nil {
		return BoolInput("DeviceDefinition.CorrectiveAction.Recall", nil)
	}
	return BoolInput("DeviceDefinition.CorrectiveAction.Recall", &resource.CorrectiveAction.Recall)
}
func (resource *DeviceDefinition) T_CorrectiveActionScope() templ.Component {
	optionsValueSet := VSDevice_correctiveactionscope

	if resource == nil {
		return CodeSelect("DeviceDefinition.CorrectiveAction.Scope", nil, optionsValueSet)
	}
	return CodeSelect("DeviceDefinition.CorrectiveAction.Scope", resource.CorrectiveAction.Scope, optionsValueSet)
}
func (resource *DeviceDefinition) T_ChargeItemId(numChargeItem int) templ.Component {

	if resource == nil || len(resource.ChargeItem) >= numChargeItem {
		return StringInput("DeviceDefinition.ChargeItem["+strconv.Itoa(numChargeItem)+"].Id", nil)
	}
	return StringInput("DeviceDefinition.ChargeItem["+strconv.Itoa(numChargeItem)+"].Id", resource.ChargeItem[numChargeItem].Id)
}
