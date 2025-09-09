package r4b

//generated with command go run ./bultaoreune
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
func (resource *DeviceDefinition) T_ManufacturerString(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DeviceDefinition.ManufacturerString", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.ManufacturerString", resource.ManufacturerString, htmlAttrs)
}
func (resource *DeviceDefinition) T_ModelNumber(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DeviceDefinition.ModelNumber", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.ModelNumber", resource.ModelNumber, htmlAttrs)
}
func (resource *DeviceDefinition) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("DeviceDefinition.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_Version(numVersion int, htmlAttrs string) templ.Component {
	if resource == nil || numVersion >= len(resource.Version) {
		return StringInput("DeviceDefinition.Version["+strconv.Itoa(numVersion)+"]", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.Version["+strconv.Itoa(numVersion)+"]", &resource.Version[numVersion], htmlAttrs)
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
func (resource *DeviceDefinition) T_Url(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DeviceDefinition.Url", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.Url", resource.Url, htmlAttrs)
}
func (resource *DeviceDefinition) T_OnlineInformation(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("DeviceDefinition.OnlineInformation", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.OnlineInformation", resource.OnlineInformation, htmlAttrs)
}
func (resource *DeviceDefinition) T_Note(numNote int, htmlAttrs string) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("DeviceDefinition.Note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("DeviceDefinition.Note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
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
func (resource *DeviceDefinition) T_SpecializationSystemType(numSpecialization int, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialization >= len(resource.Specialization) {
		return StringInput("DeviceDefinition.Specialization["+strconv.Itoa(numSpecialization)+"].SystemType", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.Specialization["+strconv.Itoa(numSpecialization)+"].SystemType", &resource.Specialization[numSpecialization].SystemType, htmlAttrs)
}
func (resource *DeviceDefinition) T_SpecializationVersion(numSpecialization int, htmlAttrs string) templ.Component {
	if resource == nil || numSpecialization >= len(resource.Specialization) {
		return StringInput("DeviceDefinition.Specialization["+strconv.Itoa(numSpecialization)+"].Version", nil, htmlAttrs)
	}
	return StringInput("DeviceDefinition.Specialization["+strconv.Itoa(numSpecialization)+"].Version", resource.Specialization[numSpecialization].Version, htmlAttrs)
}
func (resource *DeviceDefinition) T_CapabilityType(numCapability int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCapability >= len(resource.Capability) {
		return CodeableConceptSelect("DeviceDefinition.Capability["+strconv.Itoa(numCapability)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.Capability["+strconv.Itoa(numCapability)+"].Type", &resource.Capability[numCapability].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_CapabilityDescription(numCapability int, numDescription int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numCapability >= len(resource.Capability) || numDescription >= len(resource.Capability[numCapability].Description) {
		return CodeableConceptSelect("DeviceDefinition.Capability["+strconv.Itoa(numCapability)+"].Description["+strconv.Itoa(numDescription)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.Capability["+strconv.Itoa(numCapability)+"].Description["+strconv.Itoa(numDescription)+"]", &resource.Capability[numCapability].Description[numDescription], optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].Type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *DeviceDefinition) T_PropertyValueCode(numProperty int, numValueCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) || numValueCode >= len(resource.Property[numProperty].ValueCode) {
		return CodeableConceptSelect("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].ValueCode["+strconv.Itoa(numValueCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("DeviceDefinition.Property["+strconv.Itoa(numProperty)+"].ValueCode["+strconv.Itoa(numValueCode)+"]", &resource.Property[numProperty].ValueCode[numValueCode], optionsValueSet, htmlAttrs)
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
