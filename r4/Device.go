package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/Device
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

// http://hl7.org/fhir/r4/StructureDefinition/Device
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

// http://hl7.org/fhir/r4/StructureDefinition/Device
type DeviceDeviceName struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	Type              string      `json:"type"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Device
type DeviceSpecialization struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	SystemType        CodeableConcept `json:"systemType"`
	Version           *string         `json:"version,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Device
type DeviceVersion struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Component         *Identifier      `json:"component,omitempty"`
	Value             string           `json:"value"`
}

// http://hl7.org/fhir/r4/StructureDefinition/Device
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

func (resource *Device) T_Id() templ.Component {

	if resource == nil {
		return StringInput("Device.Id", nil)
	}
	return StringInput("Device.Id", resource.Id)
}
func (resource *Device) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("Device.ImplicitRules", nil)
	}
	return StringInput("Device.ImplicitRules", resource.ImplicitRules)
}
func (resource *Device) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("Device.Language", nil, optionsValueSet)
	}
	return CodeSelect("Device.Language", resource.Language, optionsValueSet)
}
func (resource *Device) T_Status() templ.Component {
	optionsValueSet := VSDevice_status

	if resource == nil {
		return CodeSelect("Device.Status", nil, optionsValueSet)
	}
	return CodeSelect("Device.Status", resource.Status, optionsValueSet)
}
func (resource *Device) T_StatusReason(numStatusReason int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.StatusReason) >= numStatusReason {
		return CodeableConceptSelect("Device.StatusReason["+strconv.Itoa(numStatusReason)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Device.StatusReason["+strconv.Itoa(numStatusReason)+"]", &resource.StatusReason[numStatusReason], optionsValueSet)
}
func (resource *Device) T_DistinctIdentifier() templ.Component {

	if resource == nil {
		return StringInput("Device.DistinctIdentifier", nil)
	}
	return StringInput("Device.DistinctIdentifier", resource.DistinctIdentifier)
}
func (resource *Device) T_Manufacturer() templ.Component {

	if resource == nil {
		return StringInput("Device.Manufacturer", nil)
	}
	return StringInput("Device.Manufacturer", resource.Manufacturer)
}
func (resource *Device) T_ManufactureDate() templ.Component {

	if resource == nil {
		return StringInput("Device.ManufactureDate", nil)
	}
	return StringInput("Device.ManufactureDate", resource.ManufactureDate)
}
func (resource *Device) T_ExpirationDate() templ.Component {

	if resource == nil {
		return StringInput("Device.ExpirationDate", nil)
	}
	return StringInput("Device.ExpirationDate", resource.ExpirationDate)
}
func (resource *Device) T_LotNumber() templ.Component {

	if resource == nil {
		return StringInput("Device.LotNumber", nil)
	}
	return StringInput("Device.LotNumber", resource.LotNumber)
}
func (resource *Device) T_SerialNumber() templ.Component {

	if resource == nil {
		return StringInput("Device.SerialNumber", nil)
	}
	return StringInput("Device.SerialNumber", resource.SerialNumber)
}
func (resource *Device) T_ModelNumber() templ.Component {

	if resource == nil {
		return StringInput("Device.ModelNumber", nil)
	}
	return StringInput("Device.ModelNumber", resource.ModelNumber)
}
func (resource *Device) T_PartNumber() templ.Component {

	if resource == nil {
		return StringInput("Device.PartNumber", nil)
	}
	return StringInput("Device.PartNumber", resource.PartNumber)
}
func (resource *Device) T_Type(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Device.Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Device.Type", resource.Type, optionsValueSet)
}
func (resource *Device) T_Url() templ.Component {

	if resource == nil {
		return StringInput("Device.Url", nil)
	}
	return StringInput("Device.Url", resource.Url)
}
func (resource *Device) T_Safety(numSafety int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Safety) >= numSafety {
		return CodeableConceptSelect("Device.Safety["+strconv.Itoa(numSafety)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Device.Safety["+strconv.Itoa(numSafety)+"]", &resource.Safety[numSafety], optionsValueSet)
}
func (resource *Device) T_UdiCarrierId(numUdiCarrier int) templ.Component {

	if resource == nil || len(resource.UdiCarrier) >= numUdiCarrier {
		return StringInput("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].Id", nil)
	}
	return StringInput("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].Id", resource.UdiCarrier[numUdiCarrier].Id)
}
func (resource *Device) T_UdiCarrierDeviceIdentifier(numUdiCarrier int) templ.Component {

	if resource == nil || len(resource.UdiCarrier) >= numUdiCarrier {
		return StringInput("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].DeviceIdentifier", nil)
	}
	return StringInput("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].DeviceIdentifier", resource.UdiCarrier[numUdiCarrier].DeviceIdentifier)
}
func (resource *Device) T_UdiCarrierIssuer(numUdiCarrier int) templ.Component {

	if resource == nil || len(resource.UdiCarrier) >= numUdiCarrier {
		return StringInput("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].Issuer", nil)
	}
	return StringInput("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].Issuer", resource.UdiCarrier[numUdiCarrier].Issuer)
}
func (resource *Device) T_UdiCarrierJurisdiction(numUdiCarrier int) templ.Component {

	if resource == nil || len(resource.UdiCarrier) >= numUdiCarrier {
		return StringInput("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].Jurisdiction", nil)
	}
	return StringInput("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].Jurisdiction", resource.UdiCarrier[numUdiCarrier].Jurisdiction)
}
func (resource *Device) T_UdiCarrierCarrierAIDC(numUdiCarrier int) templ.Component {

	if resource == nil || len(resource.UdiCarrier) >= numUdiCarrier {
		return StringInput("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].CarrierAIDC", nil)
	}
	return StringInput("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].CarrierAIDC", resource.UdiCarrier[numUdiCarrier].CarrierAIDC)
}
func (resource *Device) T_UdiCarrierCarrierHRF(numUdiCarrier int) templ.Component {

	if resource == nil || len(resource.UdiCarrier) >= numUdiCarrier {
		return StringInput("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].CarrierHRF", nil)
	}
	return StringInput("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].CarrierHRF", resource.UdiCarrier[numUdiCarrier].CarrierHRF)
}
func (resource *Device) T_UdiCarrierEntryType(numUdiCarrier int) templ.Component {
	optionsValueSet := VSUdi_entry_type

	if resource == nil || len(resource.UdiCarrier) >= numUdiCarrier {
		return CodeSelect("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].EntryType", nil, optionsValueSet)
	}
	return CodeSelect("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].EntryType", resource.UdiCarrier[numUdiCarrier].EntryType, optionsValueSet)
}
func (resource *Device) T_DeviceNameId(numDeviceName int) templ.Component {

	if resource == nil || len(resource.DeviceName) >= numDeviceName {
		return StringInput("Device.DeviceName["+strconv.Itoa(numDeviceName)+"].Id", nil)
	}
	return StringInput("Device.DeviceName["+strconv.Itoa(numDeviceName)+"].Id", resource.DeviceName[numDeviceName].Id)
}
func (resource *Device) T_DeviceNameName(numDeviceName int) templ.Component {

	if resource == nil || len(resource.DeviceName) >= numDeviceName {
		return StringInput("Device.DeviceName["+strconv.Itoa(numDeviceName)+"].Name", nil)
	}
	return StringInput("Device.DeviceName["+strconv.Itoa(numDeviceName)+"].Name", &resource.DeviceName[numDeviceName].Name)
}
func (resource *Device) T_DeviceNameType(numDeviceName int) templ.Component {
	optionsValueSet := VSDevice_nametype

	if resource == nil || len(resource.DeviceName) >= numDeviceName {
		return CodeSelect("Device.DeviceName["+strconv.Itoa(numDeviceName)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("Device.DeviceName["+strconv.Itoa(numDeviceName)+"].Type", &resource.DeviceName[numDeviceName].Type, optionsValueSet)
}
func (resource *Device) T_SpecializationId(numSpecialization int) templ.Component {

	if resource == nil || len(resource.Specialization) >= numSpecialization {
		return StringInput("Device.Specialization["+strconv.Itoa(numSpecialization)+"].Id", nil)
	}
	return StringInput("Device.Specialization["+strconv.Itoa(numSpecialization)+"].Id", resource.Specialization[numSpecialization].Id)
}
func (resource *Device) T_SpecializationSystemType(numSpecialization int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Specialization) >= numSpecialization {
		return CodeableConceptSelect("Device.Specialization["+strconv.Itoa(numSpecialization)+"].SystemType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Device.Specialization["+strconv.Itoa(numSpecialization)+"].SystemType", &resource.Specialization[numSpecialization].SystemType, optionsValueSet)
}
func (resource *Device) T_SpecializationVersion(numSpecialization int) templ.Component {

	if resource == nil || len(resource.Specialization) >= numSpecialization {
		return StringInput("Device.Specialization["+strconv.Itoa(numSpecialization)+"].Version", nil)
	}
	return StringInput("Device.Specialization["+strconv.Itoa(numSpecialization)+"].Version", resource.Specialization[numSpecialization].Version)
}
func (resource *Device) T_VersionId(numVersion int) templ.Component {

	if resource == nil || len(resource.Version) >= numVersion {
		return StringInput("Device.Version["+strconv.Itoa(numVersion)+"].Id", nil)
	}
	return StringInput("Device.Version["+strconv.Itoa(numVersion)+"].Id", resource.Version[numVersion].Id)
}
func (resource *Device) T_VersionType(numVersion int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Version) >= numVersion {
		return CodeableConceptSelect("Device.Version["+strconv.Itoa(numVersion)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Device.Version["+strconv.Itoa(numVersion)+"].Type", resource.Version[numVersion].Type, optionsValueSet)
}
func (resource *Device) T_VersionValue(numVersion int) templ.Component {

	if resource == nil || len(resource.Version) >= numVersion {
		return StringInput("Device.Version["+strconv.Itoa(numVersion)+"].Value", nil)
	}
	return StringInput("Device.Version["+strconv.Itoa(numVersion)+"].Value", &resource.Version[numVersion].Value)
}
func (resource *Device) T_PropertyId(numProperty int) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return StringInput("Device.Property["+strconv.Itoa(numProperty)+"].Id", nil)
	}
	return StringInput("Device.Property["+strconv.Itoa(numProperty)+"].Id", resource.Property[numProperty].Id)
}
func (resource *Device) T_PropertyType(numProperty int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty {
		return CodeableConceptSelect("Device.Property["+strconv.Itoa(numProperty)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Device.Property["+strconv.Itoa(numProperty)+"].Type", &resource.Property[numProperty].Type, optionsValueSet)
}
func (resource *Device) T_PropertyValueCode(numProperty int, numValueCode int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Property) >= numProperty || len(resource.Property[numProperty].ValueCode) >= numValueCode {
		return CodeableConceptSelect("Device.Property["+strconv.Itoa(numProperty)+"].ValueCode["+strconv.Itoa(numValueCode)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Device.Property["+strconv.Itoa(numProperty)+"].ValueCode["+strconv.Itoa(numValueCode)+"]", &resource.Property[numProperty].ValueCode[numValueCode], optionsValueSet)
}
