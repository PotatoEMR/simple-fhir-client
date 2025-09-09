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
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *DeviceDefinition) T_PartNumber(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("partNumber", nil, htmlAttrs)
	}
	return StringInput("partNumber", resource.PartNumber, htmlAttrs)
}
func (resource *DeviceDefinition) T_ModelNumber(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("modelNumber", nil, htmlAttrs)
	}
	return StringInput("modelNumber", resource.ModelNumber, htmlAttrs)
}
func (resource *DeviceDefinition) T_Safety(numSafety int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSafety >= len(resource.Safety) {
		return CodeableConceptSelect("safety["+strconv.Itoa(numSafety)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("safety["+strconv.Itoa(numSafety)+"]", &resource.Safety[numSafety], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_LanguageCode(numLanguageCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLanguageCode >= len(resource.LanguageCode) {
		return CodeableConceptSelect("languageCode["+strconv.Itoa(numLanguageCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("languageCode["+strconv.Itoa(numLanguageCode)+"]", &resource.LanguageCode[numLanguageCode], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *DeviceDefinition) T_ProductionIdentifierInUDI(numProductionIdentifierInUDI int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDevice_productidentifierinudi

	if resource == nil || numProductionIdentifierInUDI >= len(resource.ProductionIdentifierInUDI) {
		return CodeSelect("productionIdentifierInUDI["+strconv.Itoa(numProductionIdentifierInUDI)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("productionIdentifierInUDI["+strconv.Itoa(numProductionIdentifierInUDI)+"]", &resource.ProductionIdentifierInUDI[numProductionIdentifierInUDI], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierDeviceIdentifier(numUdiDeviceIdentifier int, htmlAttrs string) templ.Component {
	if resource == nil || numUdiDeviceIdentifier >= len(resource.UdiDeviceIdentifier) {
		return StringInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].deviceIdentifier", nil, htmlAttrs)
	}
	return StringInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].deviceIdentifier", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].DeviceIdentifier, htmlAttrs)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierIssuer(numUdiDeviceIdentifier int, htmlAttrs string) templ.Component {
	if resource == nil || numUdiDeviceIdentifier >= len(resource.UdiDeviceIdentifier) {
		return StringInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].issuer", nil, htmlAttrs)
	}
	return StringInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].issuer", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].Issuer, htmlAttrs)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierJurisdiction(numUdiDeviceIdentifier int, htmlAttrs string) templ.Component {
	if resource == nil || numUdiDeviceIdentifier >= len(resource.UdiDeviceIdentifier) {
		return StringInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].jurisdiction", nil, htmlAttrs)
	}
	return StringInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].jurisdiction", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].Jurisdiction, htmlAttrs)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierMarketDistributionSubJurisdiction(numUdiDeviceIdentifier int, numMarketDistribution int, htmlAttrs string) templ.Component {
	if resource == nil || numUdiDeviceIdentifier >= len(resource.UdiDeviceIdentifier) || numMarketDistribution >= len(resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].MarketDistribution) {
		return StringInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].marketDistribution["+strconv.Itoa(numMarketDistribution)+"].subJurisdiction", nil, htmlAttrs)
	}
	return StringInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].marketDistribution["+strconv.Itoa(numMarketDistribution)+"].subJurisdiction", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].MarketDistribution[numMarketDistribution].SubJurisdiction, htmlAttrs)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierType(numRegulatoryIdentifier int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDevicedefinition_regulatory_identifier_type

	if resource == nil || numRegulatoryIdentifier >= len(resource.RegulatoryIdentifier) {
		return CodeSelect("regulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("regulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].type", &resource.RegulatoryIdentifier[numRegulatoryIdentifier].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierDeviceIdentifier(numRegulatoryIdentifier int, htmlAttrs string) templ.Component {
	if resource == nil || numRegulatoryIdentifier >= len(resource.RegulatoryIdentifier) {
		return StringInput("regulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].deviceIdentifier", nil, htmlAttrs)
	}
	return StringInput("regulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].deviceIdentifier", &resource.RegulatoryIdentifier[numRegulatoryIdentifier].DeviceIdentifier, htmlAttrs)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierIssuer(numRegulatoryIdentifier int, htmlAttrs string) templ.Component {
	if resource == nil || numRegulatoryIdentifier >= len(resource.RegulatoryIdentifier) {
		return StringInput("regulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].issuer", nil, htmlAttrs)
	}
	return StringInput("regulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].issuer", &resource.RegulatoryIdentifier[numRegulatoryIdentifier].Issuer, htmlAttrs)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierJurisdiction(numRegulatoryIdentifier int, htmlAttrs string) templ.Component {
	if resource == nil || numRegulatoryIdentifier >= len(resource.RegulatoryIdentifier) {
		return StringInput("regulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].jurisdiction", nil, htmlAttrs)
	}
	return StringInput("regulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].jurisdiction", &resource.RegulatoryIdentifier[numRegulatoryIdentifier].Jurisdiction, htmlAttrs)
}
func (resource *DeviceDefinition) T_DeviceNameName(numDeviceName int, htmlAttrs string) templ.Component {
	if resource == nil || numDeviceName >= len(resource.DeviceName) {
		return StringInput("deviceName["+strconv.Itoa(numDeviceName)+"].name", nil, htmlAttrs)
	}
	return StringInput("deviceName["+strconv.Itoa(numDeviceName)+"].name", &resource.DeviceName[numDeviceName].Name, htmlAttrs)
}
func (resource *DeviceDefinition) T_DeviceNameType(numDeviceName int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDevice_nametype

	if resource == nil || numDeviceName >= len(resource.DeviceName) {
		return CodeSelect("deviceName["+strconv.Itoa(numDeviceName)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("deviceName["+strconv.Itoa(numDeviceName)+"].type", &resource.DeviceName[numDeviceName].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_ClassificationType(numClassification int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numClassification >= len(resource.Classification) {
		return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"].type", &resource.Classification[numClassification].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_ConformsToCategory(numConformsTo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numConformsTo >= len(resource.ConformsTo) {
		return CodeableConceptSelect("conformsTo["+strconv.Itoa(numConformsTo)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("conformsTo["+strconv.Itoa(numConformsTo)+"].category", resource.ConformsTo[numConformsTo].Category, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_ConformsToSpecification(numConformsTo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numConformsTo >= len(resource.ConformsTo) {
		return CodeableConceptSelect("conformsTo["+strconv.Itoa(numConformsTo)+"].specification", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("conformsTo["+strconv.Itoa(numConformsTo)+"].specification", &resource.ConformsTo[numConformsTo].Specification, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_ConformsToVersion(numConformsTo int, numVersion int, htmlAttrs string) templ.Component {
	if resource == nil || numConformsTo >= len(resource.ConformsTo) || numVersion >= len(resource.ConformsTo[numConformsTo].Version) {
		return StringInput("conformsTo["+strconv.Itoa(numConformsTo)+"].version["+strconv.Itoa(numVersion)+"]", nil, htmlAttrs)
	}
	return StringInput("conformsTo["+strconv.Itoa(numConformsTo)+"].version["+strconv.Itoa(numVersion)+"]", &resource.ConformsTo[numConformsTo].Version[numVersion], htmlAttrs)
}
func (resource *DeviceDefinition) T_HasPartCount(numHasPart int, htmlAttrs string) templ.Component {
	if resource == nil || numHasPart >= len(resource.HasPart) {
		return IntInput("hasPart["+strconv.Itoa(numHasPart)+"].count", nil, htmlAttrs)
	}
	return IntInput("hasPart["+strconv.Itoa(numHasPart)+"].count", resource.HasPart[numHasPart].Count, htmlAttrs)
}
func (resource *DeviceDefinition) T_PackagingType(numPackaging int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPackaging >= len(resource.Packaging) {
		return CodeableConceptSelect("packaging["+strconv.Itoa(numPackaging)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("packaging["+strconv.Itoa(numPackaging)+"].type", resource.Packaging[numPackaging].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_PackagingCount(numPackaging int, htmlAttrs string) templ.Component {
	if resource == nil || numPackaging >= len(resource.Packaging) {
		return IntInput("packaging["+strconv.Itoa(numPackaging)+"].count", nil, htmlAttrs)
	}
	return IntInput("packaging["+strconv.Itoa(numPackaging)+"].count", resource.Packaging[numPackaging].Count, htmlAttrs)
}
func (resource *DeviceDefinition) T_PackagingDistributorName(numPackaging int, numDistributor int, htmlAttrs string) templ.Component {
	if resource == nil || numPackaging >= len(resource.Packaging) || numDistributor >= len(resource.Packaging[numPackaging].Distributor) {
		return StringInput("packaging["+strconv.Itoa(numPackaging)+"].distributor["+strconv.Itoa(numDistributor)+"].name", nil, htmlAttrs)
	}
	return StringInput("packaging["+strconv.Itoa(numPackaging)+"].distributor["+strconv.Itoa(numDistributor)+"].name", resource.Packaging[numPackaging].Distributor[numDistributor].Name, htmlAttrs)
}
func (resource *DeviceDefinition) T_VersionType(numVersion int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numVersion >= len(resource.Version) {
		return CodeableConceptSelect("version["+strconv.Itoa(numVersion)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("version["+strconv.Itoa(numVersion)+"].type", resource.Version[numVersion].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_VersionValue(numVersion int, htmlAttrs string) templ.Component {
	if resource == nil || numVersion >= len(resource.Version) {
		return StringInput("version["+strconv.Itoa(numVersion)+"].value", nil, htmlAttrs)
	}
	return StringInput("version["+strconv.Itoa(numVersion)+"].value", &resource.Version[numVersion].Value, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", &resource.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueString(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("property["+strconv.Itoa(numProperty)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("property["+strconv.Itoa(numProperty)+"].valueString", &resource.Property[numProperty].ValueString, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueBoolean(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", &resource.Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueInteger(numProperty int, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return IntInput("property["+strconv.Itoa(numProperty)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("property["+strconv.Itoa(numProperty)+"].valueInteger", &resource.Property[numProperty].ValueInteger, htmlAttrs)
}
func (resource *DeviceDefinition) T_LinkRelation(numLink int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return CodingSelect("link["+strconv.Itoa(numLink)+"].relation", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("link["+strconv.Itoa(numLink)+"].relation", &resource.Link[numLink].Relation, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_MaterialSubstance(numMaterial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numMaterial >= len(resource.Material) {
		return CodeableConceptSelect("material["+strconv.Itoa(numMaterial)+"].substance", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("material["+strconv.Itoa(numMaterial)+"].substance", &resource.Material[numMaterial].Substance, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_MaterialAlternate(numMaterial int, htmlAttrs string) templ.Component {
	if resource == nil || numMaterial >= len(resource.Material) {
		return BoolInput("material["+strconv.Itoa(numMaterial)+"].alternate", nil, htmlAttrs)
	}
	return BoolInput("material["+strconv.Itoa(numMaterial)+"].alternate", resource.Material[numMaterial].Alternate, htmlAttrs)
}
func (resource *DeviceDefinition) T_MaterialAllergenicIndicator(numMaterial int, htmlAttrs string) templ.Component {
	if resource == nil || numMaterial >= len(resource.Material) {
		return BoolInput("material["+strconv.Itoa(numMaterial)+"].allergenicIndicator", nil, htmlAttrs)
	}
	return BoolInput("material["+strconv.Itoa(numMaterial)+"].allergenicIndicator", resource.Material[numMaterial].AllergenicIndicator, htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineUsageInstruction(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("guideline.usageInstruction", nil, htmlAttrs)
	}
	return StringInput("guideline.usageInstruction", resource.Guideline.UsageInstruction, htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineIndication(numIndication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numIndication >= len(resource.Guideline.Indication) {
		return CodeableConceptSelect("guideline.indication["+strconv.Itoa(numIndication)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("guideline.indication["+strconv.Itoa(numIndication)+"]", &resource.Guideline.Indication[numIndication], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineContraindication(numContraindication int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numContraindication >= len(resource.Guideline.Contraindication) {
		return CodeableConceptSelect("guideline.contraindication["+strconv.Itoa(numContraindication)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("guideline.contraindication["+strconv.Itoa(numContraindication)+"]", &resource.Guideline.Contraindication[numContraindication], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineWarning(numWarning int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numWarning >= len(resource.Guideline.Warning) {
		return CodeableConceptSelect("guideline.warning["+strconv.Itoa(numWarning)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("guideline.warning["+strconv.Itoa(numWarning)+"]", &resource.Guideline.Warning[numWarning], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineIntendedUse(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("guideline.intendedUse", nil, htmlAttrs)
	}
	return StringInput("guideline.intendedUse", resource.Guideline.IntendedUse, htmlAttrs)
}
func (resource *DeviceDefinition) T_CorrectiveActionRecall(htmlAttrs string) templ.Component {
	if resource == nil {
		return BoolInput("correctiveAction.recall", nil, htmlAttrs)
	}
	return BoolInput("correctiveAction.recall", &resource.CorrectiveAction.Recall, htmlAttrs)
}
func (resource *DeviceDefinition) T_CorrectiveActionScope(htmlAttrs string) templ.Component {
	optionsValueSet := VSDevice_correctiveactionscope

	if resource == nil {
		return CodeSelect("correctiveAction.scope", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("correctiveAction.scope", resource.CorrectiveAction.Scope, optionsValueSet, htmlAttrs)
}
