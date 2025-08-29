package r4

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
	ManufacturerString      *string                               `json:"manufacturerString"`
	ManufacturerReference   *Reference                            `json:"manufacturerReference"`
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

func (resource *DeviceDefinition) DeviceDefinitionLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *DeviceDefinition) DeviceDefinitionDeviceNameType(numDeviceName int) templ.Component {
	optionsValueSet := VSDevice_nametype
	currentVal := ""
	if resource != nil && len(resource.DeviceName) >= numDeviceName {
		currentVal = resource.DeviceName[numDeviceName].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
