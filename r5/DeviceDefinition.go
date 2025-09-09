package r5

//generated with command go run ./bultaoreune
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
func (r DeviceDefinition) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "DeviceDefinition/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "DeviceDefinition"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *DeviceDefinition) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DeviceDefinition.Description", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.Description", resource.Description, htmlAttrs)
}
func (resource *DeviceDefinition) T_PartNumber(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DeviceDefinition.PartNumber", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.PartNumber", resource.PartNumber, htmlAttrs)
}
func (resource *DeviceDefinition) T_ModelNumber(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DeviceDefinition.ModelNumber", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.ModelNumber", resource.ModelNumber, htmlAttrs)
}
func (resource *DeviceDefinition) T_Safety(numSafety int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSafety >= len(resource.Safety) {
		return CodeableConceptSelect("DeviceDefinition.Safety["+strconv.Itoa(numSafety)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.Safety["+strconv.Itoa(numSafety)+"]", &resource.Safety[numSafety], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_LanguageCode(numLanguageCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLanguageCode >= len(resource.LanguageCode) {
		return CodeableConceptSelect("DeviceDefinition.LanguageCode["+strconv.Itoa(numLanguageCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.LanguageCode["+strconv.Itoa(numLanguageCode)+"]", &resource.LanguageCode[numLanguageCode], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("DeviceDefinition.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("DeviceDefinition.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *DeviceDefinition) T_ProductionIdentifierInUDI(numProductionIdentifierInUDI int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDevice_productidentifierinudi

	if resource == nil || numProductionIdentifierInUDI >= len(resource.ProductionIdentifierInUDI) {
		return CodeSelect("DeviceDefinition.ProductionIdentifierInUDI["+strconv.Itoa(numProductionIdentifierInUDI)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DeviceDefinition.ProductionIdentifierInUDI["+strconv.Itoa(numProductionIdentifierInUDI)+"]", &resource.ProductionIdentifierInUDI[numProductionIdentifierInUDI], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierDeviceIdentifier(numUdiDeviceIdentifier int, htmlAttrs string) templ.Component {
	if resource == nil || numUdiDeviceIdentifier >= len(resource.UdiDeviceIdentifier) {
		return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].DeviceIdentifier", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].DeviceIdentifier", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].DeviceIdentifier, htmlAttrs)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierIssuer(numUdiDeviceIdentifier int, htmlAttrs string) templ.Component {
	if resource == nil || numUdiDeviceIdentifier >= len(resource.UdiDeviceIdentifier) {
		return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].Issuer", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].Issuer", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].Issuer, htmlAttrs)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierJurisdiction(numUdiDeviceIdentifier int, htmlAttrs string) templ.Component {
	if resource == nil || numUdiDeviceIdentifier >= len(resource.UdiDeviceIdentifier) {
		return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].Jurisdiction", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].Jurisdiction", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].Jurisdiction, htmlAttrs)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierMarketDistributionSubJurisdiction(numUdiDeviceIdentifier int, numMarketDistribution int, htmlAttrs string) templ.Component {
	if resource == nil || numUdiDeviceIdentifier >= len(resource.UdiDeviceIdentifier) || numMarketDistribution >= len(resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].MarketDistribution) {
		return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].MarketDistribution["+strconv.Itoa(numMarketDistribution)+"].SubJurisdiction", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.UdiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].MarketDistribution["+strconv.Itoa(numMarketDistribution)+"].SubJurisdiction", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].MarketDistribution[numMarketDistribution].SubJurisdiction, htmlAttrs)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierType(numRegulatoryIdentifier int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDevicedefinition_regulatory_identifier_type

	if resource == nil || numRegulatoryIdentifier >= len(resource.RegulatoryIdentifier) {
		return CodeSelect("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].Type", &resource.RegulatoryIdentifier[numRegulatoryIdentifier].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierDeviceIdentifier(numRegulatoryIdentifier int, htmlAttrs string) templ.Component {
	if resource == nil || numRegulatoryIdentifier >= len(resource.RegulatoryIdentifier) {
		return StringInput("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].DeviceIdentifier", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].DeviceIdentifier", &resource.RegulatoryIdentifier[numRegulatoryIdentifier].DeviceIdentifier, htmlAttrs)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierIssuer(numRegulatoryIdentifier int, htmlAttrs string) templ.Component {
	if resource == nil || numRegulatoryIdentifier >= len(resource.RegulatoryIdentifier) {
		return StringInput("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].Issuer", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].Issuer", &resource.RegulatoryIdentifier[numRegulatoryIdentifier].Issuer, htmlAttrs)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierJurisdiction(numRegulatoryIdentifier int, htmlAttrs string) templ.Component {
	if resource == nil || numRegulatoryIdentifier >= len(resource.RegulatoryIdentifier) {
		return StringInput("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].Jurisdiction", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.RegulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].Jurisdiction", &resource.RegulatoryIdentifier[numRegulatoryIdentifier].Jurisdiction, htmlAttrs)
}
func (resource *DeviceDefinition) T_DeviceNameName(numDeviceName int, htmlAttrs string) templ.Component {
	if resource == nil || numDeviceName >= len(resource.DeviceName) {
		return StringInput("DeviceDefinition.DeviceName["+strconv.Itoa(numDeviceName)+"].Name", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.DeviceName["+strconv.Itoa(numDeviceName)+"].Name", &resource.DeviceName[numDeviceName].Name, htmlAttrs)
}
func (resource *DeviceDefinition) T_DeviceNameType(numDeviceName int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDevice_nametype

	if resource == nil || numDeviceName >= len(resource.DeviceName) {
		return CodeSelect("DeviceDefinition.DeviceName["+strconv.Itoa(numDeviceName)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DeviceDefinition.DeviceName["+strconv.Itoa(numDeviceName)+"].Type", &resource.DeviceName[numDeviceName].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_ClassificationType(numClassification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numClassification >= len(resource.Classification) {
		return CodeableConceptSelect("DeviceDefinition.Classification["+strconv.Itoa(numClassification)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.Classification["+strconv.Itoa(numClassification)+"].Type", &resource.Classification[numClassification].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_ConformsToCategory(numConformsTo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numConformsTo >= len(resource.ConformsTo) {
		return CodeableConceptSelect("DeviceDefinition.ConformsTo["+strconv.Itoa(numConformsTo)+"].Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.ConformsTo["+strconv.Itoa(numConformsTo)+"].Category", resource.ConformsTo[numConformsTo].Category, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_ConformsToSpecification(numConformsTo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numConformsTo >= len(resource.ConformsTo) {
		return CodeableConceptSelect("DeviceDefinition.ConformsTo["+strconv.Itoa(numConformsTo)+"].Specification", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.ConformsTo["+strconv.Itoa(numConformsTo)+"].Specification", &resource.ConformsTo[numConformsTo].Specification, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_ConformsToVersion(numConformsTo int, numVersion int, htmlAttrs string) templ.Component {
	if resource == nil || numConformsTo >= len(resource.ConformsTo) || numVersion >= len(resource.ConformsTo[numConformsTo].Version) {
		return StringInput("DeviceDefinition.ConformsTo["+strconv.Itoa(numConformsTo)+"].Version["+strconv.Itoa(numVersion)+"]", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.ConformsTo["+strconv.Itoa(numConformsTo)+"].Version["+strconv.Itoa(numVersion)+"]", &resource.ConformsTo[numConformsTo].Version[numVersion], htmlAttrs)
}
func (resource *DeviceDefinition) T_HasPartCount(numHasPart int, htmlAttrs string) templ.Component {
	if resource == nil || numHasPart >= len(resource.HasPart) {
		return IntInput("DeviceDefinition.HasPart["+strconv.Itoa(numHasPart)+"].Count", nil, htmlAttrs)
	}
	return IntInput("DeviceDefinition.HasPart["+strconv.Itoa(numHasPart)+"].Count", resource.HasPart[numHasPart].Count, htmlAttrs)
}
func (resource *DeviceDefinition) T_PackagingType(numPackaging int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPackaging >= len(resource.Packaging) {
		return CodeableConceptSelect("DeviceDefinition.Packaging["+strconv.Itoa(numPackaging)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.Packaging["+strconv.Itoa(numPackaging)+"].Type", resource.Packaging[numPackaging].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_PackagingCount(numPackaging int, htmlAttrs string) templ.Component {
	if resource == nil || numPackaging >= len(resource.Packaging) {
		return IntInput("DeviceDefinition.Packaging["+strconv.Itoa(numPackaging)+"].Count", nil, htmlAttrs)
	}
	return IntInput("DeviceDefinition.Packaging["+strconv.Itoa(numPackaging)+"].Count", resource.Packaging[numPackaging].Count, htmlAttrs)
}
func (resource *DeviceDefinition) T_PackagingDistributorName(numPackaging int, numDistributor int, htmlAttrs string) templ.Component {
	if resource == nil || numPackaging >= len(resource.Packaging) || numDistributor >= len(resource.Packaging[numPackaging].Distributor) {
		return StringInput("DeviceDefinition.Packaging["+strconv.Itoa(numPackaging)+"].Distributor["+strconv.Itoa(numDistributor)+"].Name", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.Packaging["+strconv.Itoa(numPackaging)+"].Distributor["+strconv.Itoa(numDistributor)+"].Name", resource.Packaging[numPackaging].Distributor[numDistributor].Name, htmlAttrs)
}
func (resource *DeviceDefinition) T_VersionType(numVersion int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numVersion >= len(resource.Version) {
		return CodeableConceptSelect("DeviceDefinition.Version["+strconv.Itoa(numVersion)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.Version["+strconv.Itoa(numVersion)+"].Type", resource.Version[numVersion].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_VersionValue(numVersion int, htmlAttrs string) templ.Component {
	if resource == nil || numVersion >= len(resource.Version) {
		return StringInput("DeviceDefinition.Version["+strconv.Itoa(numVersion)+"].Value", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.Version["+strconv.Itoa(numVersion)+"].Value", &resource.Version[numVersion].Value, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].Type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].ValueCodeableConcept", &resource.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueString(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].ValueString", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].ValueString", &resource.Property[numProperty].ValueString, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueBoolean(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return BoolInput("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].ValueBoolean", &resource.Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueInteger(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return IntInput("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].ValueInteger", nil, htmlAttrs)
	}
	return IntInput("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].ValueInteger", &resource.Property[numProperty].ValueInteger, htmlAttrs)
}
func (resource *DeviceDefinition) T_LinkRelation(numLink int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return CodingSelect("DeviceDefinition.Link["+strconv.Itoa(numLink)+"].Relation", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("DeviceDefinition.Link["+strconv.Itoa(numLink)+"].Relation", &resource.Link[numLink].Relation, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_MaterialSubstance(numMaterial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMaterial >= len(resource.Material) {
		return CodeableConceptSelect("DeviceDefinition.Material["+strconv.Itoa(numMaterial)+"].Substance", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.Material["+strconv.Itoa(numMaterial)+"].Substance", &resource.Material[numMaterial].Substance, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_MaterialAlternate(numMaterial int, htmlAttrs string) templ.Component {
	if resource == nil || numMaterial >= len(resource.Material) {
		return BoolInput("DeviceDefinition.Material["+strconv.Itoa(numMaterial)+"].Alternate", nil, htmlAttrs)
	}
	return BoolInput("DeviceDefinition.Material["+strconv.Itoa(numMaterial)+"].Alternate", resource.Material[numMaterial].Alternate, htmlAttrs)
}
func (resource *DeviceDefinition) T_MaterialAllergenicIndicator(numMaterial int, htmlAttrs string) templ.Component {
	if resource == nil || numMaterial >= len(resource.Material) {
		return BoolInput("DeviceDefinition.Material["+strconv.Itoa(numMaterial)+"].AllergenicIndicator", nil, htmlAttrs)
	}
	return BoolInput("DeviceDefinition.Material["+strconv.Itoa(numMaterial)+"].AllergenicIndicator", resource.Material[numMaterial].AllergenicIndicator, htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineUsageInstruction(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DeviceDefinition.Guideline.UsageInstruction", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.Guideline.UsageInstruction", resource.Guideline.UsageInstruction, htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineIndication(numIndication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numIndication >= len(resource.Guideline.Indication) {
		return CodeableConceptSelect("DeviceDefinition.Guideline.Indication["+strconv.Itoa(numIndication)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.Guideline.Indication["+strconv.Itoa(numIndication)+"]", &resource.Guideline.Indication[numIndication], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineContraindication(numContraindication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numContraindication >= len(resource.Guideline.Contraindication) {
		return CodeableConceptSelect("DeviceDefinition.Guideline.Contraindication["+strconv.Itoa(numContraindication)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.Guideline.Contraindication["+strconv.Itoa(numContraindication)+"]", &resource.Guideline.Contraindication[numContraindication], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineWarning(numWarning int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numWarning >= len(resource.Guideline.Warning) {
		return CodeableConceptSelect("DeviceDefinition.Guideline.Warning["+strconv.Itoa(numWarning)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.Guideline.Warning["+strconv.Itoa(numWarning)+"]", &resource.Guideline.Warning[numWarning], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineIntendedUse(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DeviceDefinition.Guideline.IntendedUse", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.Guideline.IntendedUse", resource.Guideline.IntendedUse, htmlAttrs)
}
func (resource *DeviceDefinition) T_CorrectiveActionRecall(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("DeviceDefinition.CorrectiveAction.Recall", nil, htmlAttrs)
	}
	return BoolInput("DeviceDefinition.CorrectiveAction.Recall", &resource.CorrectiveAction.Recall, htmlAttrs)
}
func (resource *DeviceDefinition) T_CorrectiveActionScope(htmlAttrs string) templ.Component {
	optionsValueSet := VSDevice_correctiveactionscope

	if resource == nil {
		return CodeSelect("DeviceDefinition.CorrectiveAction.Scope", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("DeviceDefinition.CorrectiveAction.Scope", resource.CorrectiveAction.Scope, optionsValueSet, htmlAttrs)
}
