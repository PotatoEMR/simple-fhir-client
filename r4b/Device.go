package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4b/StructureDefinition/Device
type Device struct {
	Id                 *string                `json:"id,omitempty"`
	Meta               *Meta                  `json:"meta,omitempty"`
	ImplicitRules      *string                `json:"implicitRules,omitempty"`
	Language           *string                `json:"language,omitempty"`
	Text               *Narrative             `json:"text,omitempty"`
	Contained          []Resource             `json:"contained,omitempty"`
	Extension          []Extension            `json:"extension,omitempty"`
	ModifierExtension  []Extension            `json:"modifierExtension,omitempty"`
	Identifier         []Identifier           `json:"identifier,omitempty"`
	Definition         *Reference             `json:"definition,omitempty"`
	UdiCarrier         []DeviceUdiCarrier     `json:"udiCarrier,omitempty"`
	Status             *string                `json:"status,omitempty"`
	StatusReason       []CodeableConcept      `json:"statusReason,omitempty"`
	DistinctIdentifier *string                `json:"distinctIdentifier,omitempty"`
	Manufacturer       *string                `json:"manufacturer,omitempty"`
	ManufactureDate    *string                `json:"manufactureDate,omitempty"`
	ExpirationDate     *string                `json:"expirationDate,omitempty"`
	LotNumber          *string                `json:"lotNumber,omitempty"`
	SerialNumber       *string                `json:"serialNumber,omitempty"`
	DeviceName         []DeviceDeviceName     `json:"deviceName,omitempty"`
	ModelNumber        *string                `json:"modelNumber,omitempty"`
	PartNumber         *string                `json:"partNumber,omitempty"`
	Type               *CodeableConcept       `json:"type,omitempty"`
	Specialization     []DeviceSpecialization `json:"specialization,omitempty"`
	Version            []DeviceVersion        `json:"version,omitempty"`
	Property           []DeviceProperty       `json:"property,omitempty"`
	Patient            *Reference             `json:"patient,omitempty"`
	Owner              *Reference             `json:"owner,omitempty"`
	Contact            []ContactPoint         `json:"contact,omitempty"`
	Location           *Reference             `json:"location,omitempty"`
	Url                *string                `json:"url,omitempty"`
	Note               []Annotation           `json:"note,omitempty"`
	Safety             []CodeableConcept      `json:"safety,omitempty"`
	Parent             *Reference             `json:"parent,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Device
type DeviceUdiCarrier struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	DeviceIdentifier  *string     `json:"deviceIdentifier,omitempty"`
	Issuer            *string     `json:"issuer,omitempty"`
	Jurisdiction      *string     `json:"jurisdiction,omitempty"`
	CarrierAIDC       *string     `json:"carrierAIDC,omitempty"`
	CarrierHRF        *string     `json:"carrierHRF,omitempty"`
	EntryType         *string     `json:"entryType,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Device
type DeviceDeviceName struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Type              string      `json:"type"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Device
type DeviceSpecialization struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	SystemType        CodeableConcept `json:"systemType"`
	Version           *string         `json:"version,omitempty"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Device
type DeviceVersion struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Component         *Identifier      `json:"component,omitempty"`
	Value             string           `json:"value"`
}

// http://hl7.org/fhir/r4b/StructureDefinition/Device
type DeviceProperty struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type"`
	ValueQuantity     []Quantity        `json:"valueQuantity,omitempty"`
	ValueCode         []CodeableConcept `json:"valueCode,omitempty"`
}

type OtherDevice Device

// on convert struct to json, automatically add resourceType=Device
func (r Device) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDevice
		ResourceType string `json:"resourceType"`
	}{
		OtherDevice:  OtherDevice(r),
		ResourceType: "Device",
	})
}

func (resource *Device) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *Device) T_Status() templ.Component {
	optionsValueSet := VSDevice_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet)
	}
	return CodeSelect("status", resource.Status, optionsValueSet)
}
func (resource *Device) T_StatusReason(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("statusReason", nil, optionsValueSet)
	}
	return CodeableConceptSelect("statusReason", &resource.StatusReason[0], optionsValueSet)
}
func (resource *Device) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet)
}
func (resource *Device) T_Safety(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("safety", nil, optionsValueSet)
	}
	return CodeableConceptSelect("safety", &resource.Safety[0], optionsValueSet)
}
func (resource *Device) T_UdiCarrierEntryType(numUdiCarrier int) templ.Component {
	optionsValueSet := VSUdi_entry_type

	if resource == nil && len(resource.UdiCarrier) >= numUdiCarrier {
		return CodeSelect("entryType", nil, optionsValueSet)
	}
	return CodeSelect("entryType", resource.UdiCarrier[numUdiCarrier].EntryType, optionsValueSet)
}
func (resource *Device) T_DeviceNameType(numDeviceName int) templ.Component {
	optionsValueSet := VSDevice_nametype

	if resource == nil && len(resource.DeviceName) >= numDeviceName {
		return CodeSelect("type", nil, optionsValueSet)
	}
	return CodeSelect("type", &resource.DeviceName[numDeviceName].Type, optionsValueSet)
}
func (resource *Device) T_SpecializationSystemType(numSpecialization int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Specialization) >= numSpecialization {
		return CodeableConceptSelect("systemType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("systemType", &resource.Specialization[numSpecialization].SystemType, optionsValueSet)
}
func (resource *Device) T_VersionType(numVersion int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Version) >= numVersion {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", resource.Version[numVersion].Type, optionsValueSet)
}
func (resource *Device) T_PropertyType(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Property) >= numProperty {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.Property[numProperty].Type, optionsValueSet)
}
func (resource *Device) T_PropertyValueCode(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.Property) >= numProperty {
		return CodeableConceptSelect("valueCode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("valueCode", &resource.Property[numProperty].ValueCode[0], optionsValueSet)
}
