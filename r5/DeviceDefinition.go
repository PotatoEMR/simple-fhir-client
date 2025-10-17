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

// struct -> json, automatically add resourceType=Patient
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
func (r DeviceDefinition) ResourceType() string {
	return "DeviceDefinition"
}

func (resource *DeviceDefinition) T_Description(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *DeviceDefinition) T_PartNumber(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("partNumber", nil, htmlAttrs)
	}
	return StringInput("partNumber", resource.PartNumber, htmlAttrs)
}
func (resource *DeviceDefinition) T_Manufacturer(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "manufacturer", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "manufacturer", resource.Manufacturer, htmlAttrs)
}
func (resource *DeviceDefinition) T_ModelNumber(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("modelNumber", nil, htmlAttrs)
	}
	return StringInput("modelNumber", resource.ModelNumber, htmlAttrs)
}
func (resource *DeviceDefinition) T_Safety(numSafety int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSafety >= len(resource.Safety) {
		return CodeableConceptSelect("safety["+strconv.Itoa(numSafety)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("safety["+strconv.Itoa(numSafety)+"]", &resource.Safety[numSafety], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_ShelfLifeStorage(numShelfLifeStorage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numShelfLifeStorage >= len(resource.ShelfLifeStorage) {
		return ProductShelfLifeInput("shelfLifeStorage["+strconv.Itoa(numShelfLifeStorage)+"]", nil, htmlAttrs)
	}
	return ProductShelfLifeInput("shelfLifeStorage["+strconv.Itoa(numShelfLifeStorage)+"]", &resource.ShelfLifeStorage[numShelfLifeStorage], htmlAttrs)
}
func (resource *DeviceDefinition) T_LanguageCode(numLanguageCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLanguageCode >= len(resource.LanguageCode) {
		return CodeableConceptSelect("languageCode["+strconv.Itoa(numLanguageCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("languageCode["+strconv.Itoa(numLanguageCode)+"]", &resource.LanguageCode[numLanguageCode], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_Owner(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "owner", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "owner", resource.Owner, htmlAttrs)
}
func (resource *DeviceDefinition) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactPointInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactPointInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *DeviceDefinition) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *DeviceDefinition) T_ProductionIdentifierInUDI(numProductionIdentifierInUDI int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDevice_productidentifierinudi

	if resource == nil || numProductionIdentifierInUDI >= len(resource.ProductionIdentifierInUDI) {
		return CodeSelect("productionIdentifierInUDI["+strconv.Itoa(numProductionIdentifierInUDI)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("productionIdentifierInUDI["+strconv.Itoa(numProductionIdentifierInUDI)+"]", &resource.ProductionIdentifierInUDI[numProductionIdentifierInUDI], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierDeviceIdentifier(numUdiDeviceIdentifier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUdiDeviceIdentifier >= len(resource.UdiDeviceIdentifier) {
		return StringInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].deviceIdentifier", nil, htmlAttrs)
	}
	return StringInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].deviceIdentifier", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].DeviceIdentifier, htmlAttrs)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierIssuer(numUdiDeviceIdentifier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUdiDeviceIdentifier >= len(resource.UdiDeviceIdentifier) {
		return StringInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].issuer", nil, htmlAttrs)
	}
	return StringInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].issuer", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].Issuer, htmlAttrs)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierJurisdiction(numUdiDeviceIdentifier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUdiDeviceIdentifier >= len(resource.UdiDeviceIdentifier) {
		return StringInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].jurisdiction", nil, htmlAttrs)
	}
	return StringInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].jurisdiction", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].Jurisdiction, htmlAttrs)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierMarketDistributionMarketPeriod(numUdiDeviceIdentifier int, numMarketDistribution int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUdiDeviceIdentifier >= len(resource.UdiDeviceIdentifier) || numMarketDistribution >= len(resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].MarketDistribution) {
		return PeriodInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].marketDistribution["+strconv.Itoa(numMarketDistribution)+"].marketPeriod", nil, htmlAttrs)
	}
	return PeriodInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].marketDistribution["+strconv.Itoa(numMarketDistribution)+"].marketPeriod", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].MarketDistribution[numMarketDistribution].MarketPeriod, htmlAttrs)
}
func (resource *DeviceDefinition) T_UdiDeviceIdentifierMarketDistributionSubJurisdiction(numUdiDeviceIdentifier int, numMarketDistribution int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUdiDeviceIdentifier >= len(resource.UdiDeviceIdentifier) || numMarketDistribution >= len(resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].MarketDistribution) {
		return StringInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].marketDistribution["+strconv.Itoa(numMarketDistribution)+"].subJurisdiction", nil, htmlAttrs)
	}
	return StringInput("udiDeviceIdentifier["+strconv.Itoa(numUdiDeviceIdentifier)+"].marketDistribution["+strconv.Itoa(numMarketDistribution)+"].subJurisdiction", &resource.UdiDeviceIdentifier[numUdiDeviceIdentifier].MarketDistribution[numMarketDistribution].SubJurisdiction, htmlAttrs)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierType(numRegulatoryIdentifier int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDevicedefinition_regulatory_identifier_type

	if resource == nil || numRegulatoryIdentifier >= len(resource.RegulatoryIdentifier) {
		return CodeSelect("regulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("regulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].type", &resource.RegulatoryIdentifier[numRegulatoryIdentifier].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierDeviceIdentifier(numRegulatoryIdentifier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRegulatoryIdentifier >= len(resource.RegulatoryIdentifier) {
		return StringInput("regulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].deviceIdentifier", nil, htmlAttrs)
	}
	return StringInput("regulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].deviceIdentifier", &resource.RegulatoryIdentifier[numRegulatoryIdentifier].DeviceIdentifier, htmlAttrs)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierIssuer(numRegulatoryIdentifier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRegulatoryIdentifier >= len(resource.RegulatoryIdentifier) {
		return StringInput("regulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].issuer", nil, htmlAttrs)
	}
	return StringInput("regulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].issuer", &resource.RegulatoryIdentifier[numRegulatoryIdentifier].Issuer, htmlAttrs)
}
func (resource *DeviceDefinition) T_RegulatoryIdentifierJurisdiction(numRegulatoryIdentifier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRegulatoryIdentifier >= len(resource.RegulatoryIdentifier) {
		return StringInput("regulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].jurisdiction", nil, htmlAttrs)
	}
	return StringInput("regulatoryIdentifier["+strconv.Itoa(numRegulatoryIdentifier)+"].jurisdiction", &resource.RegulatoryIdentifier[numRegulatoryIdentifier].Jurisdiction, htmlAttrs)
}
func (resource *DeviceDefinition) T_DeviceNameName(numDeviceName int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDeviceName >= len(resource.DeviceName) {
		return StringInput("deviceName["+strconv.Itoa(numDeviceName)+"].name", nil, htmlAttrs)
	}
	return StringInput("deviceName["+strconv.Itoa(numDeviceName)+"].name", &resource.DeviceName[numDeviceName].Name, htmlAttrs)
}
func (resource *DeviceDefinition) T_DeviceNameType(numDeviceName int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDevice_nametype

	if resource == nil || numDeviceName >= len(resource.DeviceName) {
		return CodeSelect("deviceName["+strconv.Itoa(numDeviceName)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("deviceName["+strconv.Itoa(numDeviceName)+"].type", &resource.DeviceName[numDeviceName].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_ClassificationType(numClassification int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClassification >= len(resource.Classification) {
		return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("classification["+strconv.Itoa(numClassification)+"].type", &resource.Classification[numClassification].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_ClassificationJustification(numClassification int, numJustification int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numClassification >= len(resource.Classification) || numJustification >= len(resource.Classification[numClassification].Justification) {
		return RelatedArtifactInput("classification["+strconv.Itoa(numClassification)+"].justification["+strconv.Itoa(numJustification)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("classification["+strconv.Itoa(numClassification)+"].justification["+strconv.Itoa(numJustification)+"]", &resource.Classification[numClassification].Justification[numJustification], htmlAttrs)
}
func (resource *DeviceDefinition) T_ConformsToCategory(numConformsTo int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConformsTo >= len(resource.ConformsTo) {
		return CodeableConceptSelect("conformsTo["+strconv.Itoa(numConformsTo)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("conformsTo["+strconv.Itoa(numConformsTo)+"].category", resource.ConformsTo[numConformsTo].Category, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_ConformsToSpecification(numConformsTo int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConformsTo >= len(resource.ConformsTo) {
		return CodeableConceptSelect("conformsTo["+strconv.Itoa(numConformsTo)+"].specification", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("conformsTo["+strconv.Itoa(numConformsTo)+"].specification", &resource.ConformsTo[numConformsTo].Specification, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_ConformsToVersion(numConformsTo int, numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConformsTo >= len(resource.ConformsTo) || numVersion >= len(resource.ConformsTo[numConformsTo].Version) {
		return StringInput("conformsTo["+strconv.Itoa(numConformsTo)+"].version["+strconv.Itoa(numVersion)+"]", nil, htmlAttrs)
	}
	return StringInput("conformsTo["+strconv.Itoa(numConformsTo)+"].version["+strconv.Itoa(numVersion)+"]", &resource.ConformsTo[numConformsTo].Version[numVersion], htmlAttrs)
}
func (resource *DeviceDefinition) T_ConformsToSource(numConformsTo int, numSource int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConformsTo >= len(resource.ConformsTo) || numSource >= len(resource.ConformsTo[numConformsTo].Source) {
		return RelatedArtifactInput("conformsTo["+strconv.Itoa(numConformsTo)+"].source["+strconv.Itoa(numSource)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("conformsTo["+strconv.Itoa(numConformsTo)+"].source["+strconv.Itoa(numSource)+"]", &resource.ConformsTo[numConformsTo].Source[numSource], htmlAttrs)
}
func (resource *DeviceDefinition) T_HasPartReference(frs []FhirResource, numHasPart int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numHasPart >= len(resource.HasPart) {
		return ReferenceInput(frs, "hasPart["+strconv.Itoa(numHasPart)+"].reference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "hasPart["+strconv.Itoa(numHasPart)+"].reference", &resource.HasPart[numHasPart].Reference, htmlAttrs)
}
func (resource *DeviceDefinition) T_HasPartCount(numHasPart int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numHasPart >= len(resource.HasPart) {
		return IntInput("hasPart["+strconv.Itoa(numHasPart)+"].count", nil, htmlAttrs)
	}
	return IntInput("hasPart["+strconv.Itoa(numHasPart)+"].count", resource.HasPart[numHasPart].Count, htmlAttrs)
}
func (resource *DeviceDefinition) T_PackagingType(numPackaging int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPackaging >= len(resource.Packaging) {
		return CodeableConceptSelect("packaging["+strconv.Itoa(numPackaging)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("packaging["+strconv.Itoa(numPackaging)+"].type", resource.Packaging[numPackaging].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_PackagingCount(numPackaging int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPackaging >= len(resource.Packaging) {
		return IntInput("packaging["+strconv.Itoa(numPackaging)+"].count", nil, htmlAttrs)
	}
	return IntInput("packaging["+strconv.Itoa(numPackaging)+"].count", resource.Packaging[numPackaging].Count, htmlAttrs)
}
func (resource *DeviceDefinition) T_PackagingDistributorName(numPackaging int, numDistributor int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPackaging >= len(resource.Packaging) || numDistributor >= len(resource.Packaging[numPackaging].Distributor) {
		return StringInput("packaging["+strconv.Itoa(numPackaging)+"].distributor["+strconv.Itoa(numDistributor)+"].name", nil, htmlAttrs)
	}
	return StringInput("packaging["+strconv.Itoa(numPackaging)+"].distributor["+strconv.Itoa(numDistributor)+"].name", resource.Packaging[numPackaging].Distributor[numDistributor].Name, htmlAttrs)
}
func (resource *DeviceDefinition) T_PackagingDistributorOrganizationReference(frs []FhirResource, numPackaging int, numDistributor int, numOrganizationReference int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numPackaging >= len(resource.Packaging) || numDistributor >= len(resource.Packaging[numPackaging].Distributor) || numOrganizationReference >= len(resource.Packaging[numPackaging].Distributor[numDistributor].OrganizationReference) {
		return ReferenceInput(frs, "packaging["+strconv.Itoa(numPackaging)+"].distributor["+strconv.Itoa(numDistributor)+"].organizationReference["+strconv.Itoa(numOrganizationReference)+"]", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "packaging["+strconv.Itoa(numPackaging)+"].distributor["+strconv.Itoa(numDistributor)+"].organizationReference["+strconv.Itoa(numOrganizationReference)+"]", &resource.Packaging[numPackaging].Distributor[numDistributor].OrganizationReference[numOrganizationReference], htmlAttrs)
}
func (resource *DeviceDefinition) T_VersionType(numVersion int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVersion >= len(resource.Version) {
		return CodeableConceptSelect("version["+strconv.Itoa(numVersion)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("version["+strconv.Itoa(numVersion)+"].type", resource.Version[numVersion].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_VersionComponent(numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVersion >= len(resource.Version) {
		return IdentifierInput("version["+strconv.Itoa(numVersion)+"].component", nil, htmlAttrs)
	}
	return IdentifierInput("version["+strconv.Itoa(numVersion)+"].component", resource.Version[numVersion].Component, htmlAttrs)
}
func (resource *DeviceDefinition) T_VersionValue(numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVersion >= len(resource.Version) {
		return StringInput("version["+strconv.Itoa(numVersion)+"].value", nil, htmlAttrs)
	}
	return StringInput("version["+strconv.Itoa(numVersion)+"].value", &resource.Version[numVersion].Value, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueQuantity(numProperty int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return QuantityInput("property["+strconv.Itoa(numProperty)+"].valueQuantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("property["+strconv.Itoa(numProperty)+"].valueQuantity", &resource.Property[numProperty].ValueQuantity, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", &resource.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueString(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("property["+strconv.Itoa(numProperty)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("property["+strconv.Itoa(numProperty)+"].valueString", &resource.Property[numProperty].ValueString, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueBoolean(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", &resource.Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueInteger(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return IntInput("property["+strconv.Itoa(numProperty)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("property["+strconv.Itoa(numProperty)+"].valueInteger", &resource.Property[numProperty].ValueInteger, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueRange(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return RangeInput("property["+strconv.Itoa(numProperty)+"].valueRange", nil, htmlAttrs)
	}
	return RangeInput("property["+strconv.Itoa(numProperty)+"].valueRange", &resource.Property[numProperty].ValueRange, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueAttachment(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return AttachmentInput("property["+strconv.Itoa(numProperty)+"].valueAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("property["+strconv.Itoa(numProperty)+"].valueAttachment", &resource.Property[numProperty].ValueAttachment, htmlAttrs)
}
func (resource *DeviceDefinition) T_LinkRelation(numLink int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return CodingSelect("link["+strconv.Itoa(numLink)+"].relation", nil, optionsValueSet, htmlAttrs)
	}
	return CodingSelect("link["+strconv.Itoa(numLink)+"].relation", &resource.Link[numLink].Relation, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_LinkRelatedDevice(numLink int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numLink >= len(resource.Link) {
		return CodeableReferenceInput("link["+strconv.Itoa(numLink)+"].relatedDevice", nil, htmlAttrs)
	}
	return CodeableReferenceInput("link["+strconv.Itoa(numLink)+"].relatedDevice", &resource.Link[numLink].RelatedDevice, htmlAttrs)
}
func (resource *DeviceDefinition) T_MaterialSubstance(numMaterial int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMaterial >= len(resource.Material) {
		return CodeableConceptSelect("material["+strconv.Itoa(numMaterial)+"].substance", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("material["+strconv.Itoa(numMaterial)+"].substance", &resource.Material[numMaterial].Substance, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_MaterialAlternate(numMaterial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMaterial >= len(resource.Material) {
		return BoolInput("material["+strconv.Itoa(numMaterial)+"].alternate", nil, htmlAttrs)
	}
	return BoolInput("material["+strconv.Itoa(numMaterial)+"].alternate", resource.Material[numMaterial].Alternate, htmlAttrs)
}
func (resource *DeviceDefinition) T_MaterialAllergenicIndicator(numMaterial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMaterial >= len(resource.Material) {
		return BoolInput("material["+strconv.Itoa(numMaterial)+"].allergenicIndicator", nil, htmlAttrs)
	}
	return BoolInput("material["+strconv.Itoa(numMaterial)+"].allergenicIndicator", resource.Material[numMaterial].AllergenicIndicator, htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineUseContext(numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUseContext >= len(resource.Guideline.UseContext) {
		return UsageContextInput("guideline.useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("guideline.useContext["+strconv.Itoa(numUseContext)+"]", &resource.Guideline.UseContext[numUseContext], htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineUsageInstruction(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("guideline.usageInstruction", nil, htmlAttrs)
	}
	return StringInput("guideline.usageInstruction", resource.Guideline.UsageInstruction, htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineRelatedArtifact(numRelatedArtifact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRelatedArtifact >= len(resource.Guideline.RelatedArtifact) {
		return RelatedArtifactInput("guideline.relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", nil, htmlAttrs)
	}
	return RelatedArtifactInput("guideline.relatedArtifact["+strconv.Itoa(numRelatedArtifact)+"]", &resource.Guideline.RelatedArtifact[numRelatedArtifact], htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineIndication(numIndication int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numIndication >= len(resource.Guideline.Indication) {
		return CodeableConceptSelect("guideline.indication["+strconv.Itoa(numIndication)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("guideline.indication["+strconv.Itoa(numIndication)+"]", &resource.Guideline.Indication[numIndication], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineContraindication(numContraindication int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContraindication >= len(resource.Guideline.Contraindication) {
		return CodeableConceptSelect("guideline.contraindication["+strconv.Itoa(numContraindication)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("guideline.contraindication["+strconv.Itoa(numContraindication)+"]", &resource.Guideline.Contraindication[numContraindication], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineWarning(numWarning int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numWarning >= len(resource.Guideline.Warning) {
		return CodeableConceptSelect("guideline.warning["+strconv.Itoa(numWarning)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("guideline.warning["+strconv.Itoa(numWarning)+"]", &resource.Guideline.Warning[numWarning], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_GuidelineIntendedUse(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("guideline.intendedUse", nil, htmlAttrs)
	}
	return StringInput("guideline.intendedUse", resource.Guideline.IntendedUse, htmlAttrs)
}
func (resource *DeviceDefinition) T_CorrectiveActionRecall(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return BoolInput("correctiveAction.recall", nil, htmlAttrs)
	}
	return BoolInput("correctiveAction.recall", &resource.CorrectiveAction.Recall, htmlAttrs)
}
func (resource *DeviceDefinition) T_CorrectiveActionScope(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDevice_correctiveactionscope

	if resource == nil {
		return CodeSelect("correctiveAction.scope", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("correctiveAction.scope", resource.CorrectiveAction.Scope, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_CorrectiveActionPeriod(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return PeriodInput("correctiveAction.period", nil, htmlAttrs)
	}
	return PeriodInput("correctiveAction.period", &resource.CorrectiveAction.Period, htmlAttrs)
}
func (resource *DeviceDefinition) T_ChargeItemChargeItemCode(numChargeItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numChargeItem >= len(resource.ChargeItem) {
		return CodeableReferenceInput("chargeItem["+strconv.Itoa(numChargeItem)+"].chargeItemCode", nil, htmlAttrs)
	}
	return CodeableReferenceInput("chargeItem["+strconv.Itoa(numChargeItem)+"].chargeItemCode", &resource.ChargeItem[numChargeItem].ChargeItemCode, htmlAttrs)
}
func (resource *DeviceDefinition) T_ChargeItemCount(numChargeItem int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numChargeItem >= len(resource.ChargeItem) {
		return QuantityInput("chargeItem["+strconv.Itoa(numChargeItem)+"].count", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("chargeItem["+strconv.Itoa(numChargeItem)+"].count", &resource.ChargeItem[numChargeItem].Count, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_ChargeItemEffectivePeriod(numChargeItem int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numChargeItem >= len(resource.ChargeItem) {
		return PeriodInput("chargeItem["+strconv.Itoa(numChargeItem)+"].effectivePeriod", nil, htmlAttrs)
	}
	return PeriodInput("chargeItem["+strconv.Itoa(numChargeItem)+"].effectivePeriod", resource.ChargeItem[numChargeItem].EffectivePeriod, htmlAttrs)
}
func (resource *DeviceDefinition) T_ChargeItemUseContext(numChargeItem int, numUseContext int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numChargeItem >= len(resource.ChargeItem) || numUseContext >= len(resource.ChargeItem[numChargeItem].UseContext) {
		return UsageContextInput("chargeItem["+strconv.Itoa(numChargeItem)+"].useContext["+strconv.Itoa(numUseContext)+"]", nil, htmlAttrs)
	}
	return UsageContextInput("chargeItem["+strconv.Itoa(numChargeItem)+"].useContext["+strconv.Itoa(numUseContext)+"]", &resource.ChargeItem[numChargeItem].UseContext[numUseContext], htmlAttrs)
}
