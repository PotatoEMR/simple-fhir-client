package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/DeviceDefinition
type DeviceDefinition struct {
	Id                      *string                               `json:"id,omitempty"`
	Meta                    *Meta                                 `json:"meta,omitempty"`
	ImplicitRules           *string                               `json:"implicitRules,omitempty"`
	Language                *string                               `json:"language,omitempty"`
	Text                    *Narrative                            `json:"text,omitempty"`
	Contained               []Resource                            `json:"contained,omitempty"`
	Extension               []Extension                           `json:"extension,omitempty"`
	ModifierExtension       []Extension                           `json:"modifierExtension,omitempty"`
	Identifier              []Identifier                          `json:"identifier,omitempty"`
	UdiDeviceIdentifier     []DeviceDefinitionUdiDeviceIdentifier `json:"udiDeviceIdentifier,omitempty"`
	ManufacturerString      *string                               `json:"manufacturerString,omitempty"`
	ManufacturerReference   *Reference                            `json:"manufacturerReference,omitempty"`
	DeviceName              []DeviceDefinitionDeviceName          `json:"deviceName,omitempty"`
	ModelNumber             *string                               `json:"modelNumber,omitempty"`
	Type                    *CodeableConcept                      `json:"type,omitempty"`
	Specialization          []DeviceDefinitionSpecialization      `json:"specialization,omitempty"`
	Version                 []string                              `json:"version,omitempty"`
	Safety                  []CodeableConcept                     `json:"safety,omitempty"`
	ShelfLifeStorage        []ProductShelfLife                    `json:"shelfLifeStorage,omitempty"`
	PhysicalCharacteristics *ProdCharacteristic                   `json:"physicalCharacteristics,omitempty"`
	LanguageCode            []CodeableConcept                     `json:"languageCode,omitempty"`
	Capability              []DeviceDefinitionCapability          `json:"capability,omitempty"`
	Property                []DeviceDefinitionProperty            `json:"property,omitempty"`
	Owner                   *Reference                            `json:"owner,omitempty"`
	Contact                 []ContactPoint                        `json:"contact,omitempty"`
	Url                     *string                               `json:"url,omitempty"`
	OnlineInformation       *string                               `json:"onlineInformation,omitempty"`
	Note                    []Annotation                          `json:"note,omitempty"`
	Quantity                *Quantity                             `json:"quantity,omitempty"`
	ParentDevice            *Reference                            `json:"parentDevice,omitempty"`
	Material                []DeviceDefinitionMaterial            `json:"material,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/DeviceDefinition
type DeviceDefinitionUdiDeviceIdentifier struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	DeviceIdentifier  string      `json:"deviceIdentifier"`
	Issuer            string      `json:"issuer"`
	Jurisdiction      string      `json:"jurisdiction"`
}

// http://hl7.org/fhir/r4/StructureDefinition/DeviceDefinition
type DeviceDefinitionDeviceName struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Type              string      `json:"type"`
}

// http://hl7.org/fhir/r4/StructureDefinition/DeviceDefinition
type DeviceDefinitionSpecialization struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	SystemType        string      `json:"systemType"`
	Version           *string     `json:"version,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/DeviceDefinition
type DeviceDefinitionCapability struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	Description       []CodeableConcept `json:"description,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/DeviceDefinition
type DeviceDefinitionProperty struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	ValueQuantity     []Quantity        `json:"valueQuantity,omitempty"`
	ValueCode         []CodeableConcept `json:"valueCode,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/DeviceDefinition
type DeviceDefinitionMaterial struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	Substance           CodeableConcept `json:"substance"`
	Alternate           *bool           `json:"alternate,omitempty"`
	AllergenicIndicator *bool           `json:"allergenicIndicator,omitempty"`
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

func (resource *DeviceDefinition) T_ManufacturerString(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("manufacturerString", nil, htmlAttrs)
	}
	return StringInput("manufacturerString", resource.ManufacturerString, htmlAttrs)
}
func (resource *DeviceDefinition) T_ManufacturerReference(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "manufacturerReference", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "manufacturerReference", resource.ManufacturerReference, htmlAttrs)
}
func (resource *DeviceDefinition) T_ModelNumber(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("modelNumber", nil, htmlAttrs)
	}
	return StringInput("modelNumber", resource.ModelNumber, htmlAttrs)
}
func (resource *DeviceDefinition) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_Version(numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVersion >= len(resource.Version) {
		return StringInput("version["+strconv.Itoa(numVersion)+"]", nil, htmlAttrs)
	}
	return StringInput("version["+strconv.Itoa(numVersion)+"]", &resource.Version[numVersion], htmlAttrs)
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
func (resource *DeviceDefinition) T_PhysicalCharacteristics(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ProdCharacteristicInput("physicalCharacteristics", nil, htmlAttrs)
	}
	return ProdCharacteristicInput("physicalCharacteristics", resource.PhysicalCharacteristics, htmlAttrs)
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
func (resource *DeviceDefinition) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *DeviceDefinition) T_OnlineInformation(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("onlineInformation", nil, htmlAttrs)
	}
	return StringInput("onlineInformation", resource.OnlineInformation, htmlAttrs)
}
func (resource *DeviceDefinition) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *DeviceDefinition) T_Quantity(optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil {
		return QuantityInput("quantity", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("quantity", resource.Quantity, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_ParentDevice(frs []FhirResource, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput(frs, "parentDevice", nil, htmlAttrs)
	}
	return ReferenceInput(frs, "parentDevice", resource.ParentDevice, htmlAttrs)
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
func (resource *DeviceDefinition) T_SpecializationSystemType(numSpecialization int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecialization >= len(resource.Specialization) {
		return StringInput("specialization["+strconv.Itoa(numSpecialization)+"].systemType", nil, htmlAttrs)
	}
	return StringInput("specialization["+strconv.Itoa(numSpecialization)+"].systemType", &resource.Specialization[numSpecialization].SystemType, htmlAttrs)
}
func (resource *DeviceDefinition) T_SpecializationVersion(numSpecialization int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecialization >= len(resource.Specialization) {
		return StringInput("specialization["+strconv.Itoa(numSpecialization)+"].version", nil, htmlAttrs)
	}
	return StringInput("specialization["+strconv.Itoa(numSpecialization)+"].version", resource.Specialization[numSpecialization].Version, htmlAttrs)
}
func (resource *DeviceDefinition) T_CapabilityType(numCapability int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCapability >= len(resource.Capability) {
		return CodeableConceptSelect("capability["+strconv.Itoa(numCapability)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("capability["+strconv.Itoa(numCapability)+"].type", &resource.Capability[numCapability].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_CapabilityDescription(numCapability int, numDescription int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCapability >= len(resource.Capability) || numDescription >= len(resource.Capability[numCapability].Description) {
		return CodeableConceptSelect("capability["+strconv.Itoa(numCapability)+"].description["+strconv.Itoa(numDescription)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("capability["+strconv.Itoa(numCapability)+"].description["+strconv.Itoa(numDescription)+"]", &resource.Capability[numCapability].Description[numDescription], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueQuantity(numProperty int, numValueQuantity int, optionsValueSet []Coding, htmlAttrs QuantityAttrs) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) || numValueQuantity >= len(resource.Property[numProperty].ValueQuantity) {
		return QuantityInput("property["+strconv.Itoa(numProperty)+"].valueQuantity["+strconv.Itoa(numValueQuantity)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return QuantityInput("property["+strconv.Itoa(numProperty)+"].valueQuantity["+strconv.Itoa(numValueQuantity)+"]", &resource.Property[numProperty].ValueQuantity[numValueQuantity], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueCode(numProperty int, numValueCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) || numValueCode >= len(resource.Property[numProperty].ValueCode) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCode["+strconv.Itoa(numValueCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCode["+strconv.Itoa(numValueCode)+"]", &resource.Property[numProperty].ValueCode[numValueCode], optionsValueSet, htmlAttrs)
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
