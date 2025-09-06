package r4b

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4b/StructureDefinition/DeviceDefinition
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

// http://hl7.org/fhir/r4b/StructureDefinition/DeviceDefinition
type DeviceDefinitionUdiDeviceIdentifier struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	DeviceIdentifier  string      `json:"deviceIdentifier"`
	Issuer            string      `json:"issuer"`
	Jurisdiction      string      `json:"jurisdiction"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/DeviceDefinition
type DeviceDefinitionDeviceName struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Type              string      `json:"type"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/DeviceDefinition
type DeviceDefinitionSpecialization struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	SystemType        string      `json:"systemType"`
	Version           *string     `json:"version,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/DeviceDefinition
type DeviceDefinitionCapability struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	Description       []CodeableConcept `json:"description,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/DeviceDefinition
type DeviceDefinitionProperty struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	ValueQuantity     []Quantity        `json:"valueQuantity,omitempty"`
	ValueCode         []CodeableConcept `json:"valueCode,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/DeviceDefinition
type DeviceDefinitionMaterial struct {
	Id                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	Substance           CodeableConcept `json:"substance"`
	Alternate           *bool           `json:"alternate,omitempty"`
	AllergenicIndicator *bool           `json:"allergenicIndicator,omitempty"`
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
func (resource *DeviceDefinition) T_ModelNumber() templ.Component {

	if resource == nil {
		return StringInput("DeviceDefinition.ModelNumber", nil)
	}
	return StringInput("DeviceDefinition.ModelNumber", resource.ModelNumber)
}
func (resource *DeviceDefinition) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("DeviceDefinition.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDefinition.Type", resource.Type, optionsValueSet)
}
func (resource *DeviceDefinition) T_Version(numVersion int) templ.Component {

	if resource == nil || len(resource.Version) >= numVersion {
		return StringInput("DeviceDefinition.Version["+strconv.Itoa(numVersion)+"]", nil)
	}
	return StringInput("DeviceDefinition.Version["+strconv.Itoa(numVersion)+"]", &resource.Version[numVersion])
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
func (resource *DeviceDefinition) T_Url() templ.Component {

	if resource == nil {
		return StringInput("DeviceDefinition.Url", nil)
	}
	return StringInput("DeviceDefinition.Url", resource.Url)
}
func (resource *DeviceDefinition) T_OnlineInformation() templ.Component {

	if resource == nil {
		return StringInput("DeviceDefinition.OnlineInformation", nil)
	}
	return StringInput("DeviceDefinition.OnlineInformation", resource.OnlineInformation)
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
func (resource *DeviceDefinition) T_SpecializationId(numSpecialization int) templ.Component {

	if resource == nil || len(resource.Specialization) >= numSpecialization {
		return StringInput("DeviceDefinition.Specialization["+strconv.Itoa(numSpecialization)+"].Id", nil)
	}
	return StringInput("DeviceDefinition.Specialization["+strconv.Itoa(numSpecialization)+"].Id", resource.Specialization[numSpecialization].Id)
}
func (resource *DeviceDefinition) T_SpecializationSystemType(numSpecialization int) templ.Component {

	if resource == nil || len(resource.Specialization) >= numSpecialization {
		return StringInput("DeviceDefinition.Specialization["+strconv.Itoa(numSpecialization)+"].SystemType", nil)
	}
	return StringInput("DeviceDefinition.Specialization["+strconv.Itoa(numSpecialization)+"].SystemType", &resource.Specialization[numSpecialization].SystemType)
}
func (resource *DeviceDefinition) T_SpecializationVersion(numSpecialization int) templ.Component {

	if resource == nil || len(resource.Specialization) >= numSpecialization {
		return StringInput("DeviceDefinition.Specialization["+strconv.Itoa(numSpecialization)+"].Version", nil)
	}
	return StringInput("DeviceDefinition.Specialization["+strconv.Itoa(numSpecialization)+"].Version", resource.Specialization[numSpecialization].Version)
}
func (resource *DeviceDefinition) T_CapabilityId(numCapability int) templ.Component {

	if resource == nil || len(resource.Capability) >= numCapability {
		return StringInput("DeviceDefinition.Capability["+strconv.Itoa(numCapability)+"].Id", nil)
	}
	return StringInput("DeviceDefinition.Capability["+strconv.Itoa(numCapability)+"].Id", resource.Capability[numCapability].Id)
}
func (resource *DeviceDefinition) T_CapabilityType(numCapability int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Capability) >= numCapability {
		return CodeableConceptSelect("DeviceDefinition.Capability["+strconv.Itoa(numCapability)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDefinition.Capability["+strconv.Itoa(numCapability)+"].Type", &resource.Capability[numCapability].Type, optionsValueSet)
}
func (resource *DeviceDefinition) T_CapabilityDescription(numCapability int, numDescription int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Capability) >= numCapability || len(resource.Capability[numCapability].Description) >= numDescription {
		return CodeableConceptSelect("DeviceDefinition.Capability["+strconv.Itoa(numCapability)+"].Description["+strconv.Itoa(numDescription)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDefinition.Capability["+strconv.Itoa(numCapability)+"].Description["+strconv.Itoa(numDescription)+"]", &resource.Capability[numCapability].Description[numDescription], optionsValueSet)
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
func (resource *DeviceDefinition) T_PropertyValueCode(numProperty int, numValueCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty || len(resource.Property[numProperty].ValueCode) >= numValueCode {
		return CodeableConceptSelect("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].ValueCode["+strconv.Itoa(numValueCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].ValueCode["+strconv.Itoa(numValueCode)+"]", &resource.Property[numProperty].ValueCode[numValueCode], optionsValueSet)
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
