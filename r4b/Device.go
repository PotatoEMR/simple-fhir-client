package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

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
	ManufactureDate    *time.Time             `json:"manufactureDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	ExpirationDate     *time.Time             `json:"expirationDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (r Device) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "Device/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "Device"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *Device) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSDevice_status

	if resource == nil {
		return CodeSelect("Device.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Device.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_StatusReason(numStatusReason int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numStatusReason >= len(resource.StatusReason) {
		return CodeableConceptSelect("Device.StatusReason."+strconv.Itoa(numStatusReason)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Device.StatusReason."+strconv.Itoa(numStatusReason)+".", &resource.StatusReason[numStatusReason], optionsValueSet, htmlAttrs)
}
func (resource *Device) T_DistinctIdentifier(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Device.DistinctIdentifier", nil, htmlAttrs)
	}
	return StringInput("Device.DistinctIdentifier", resource.DistinctIdentifier, htmlAttrs)
}
func (resource *Device) T_Manufacturer(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Device.Manufacturer", nil, htmlAttrs)
	}
	return StringInput("Device.Manufacturer", resource.Manufacturer, htmlAttrs)
}
func (resource *Device) T_ManufactureDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Device.ManufactureDate", nil, htmlAttrs)
	}
	return DateTimeInput("Device.ManufactureDate", resource.ManufactureDate, htmlAttrs)
}
func (resource *Device) T_ExpirationDate(htmlAttrs string) templ.Component {

	if resource == nil {
		return DateTimeInput("Device.ExpirationDate", nil, htmlAttrs)
	}
	return DateTimeInput("Device.ExpirationDate", resource.ExpirationDate, htmlAttrs)
}
func (resource *Device) T_LotNumber(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Device.LotNumber", nil, htmlAttrs)
	}
	return StringInput("Device.LotNumber", resource.LotNumber, htmlAttrs)
}
func (resource *Device) T_SerialNumber(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Device.SerialNumber", nil, htmlAttrs)
	}
	return StringInput("Device.SerialNumber", resource.SerialNumber, htmlAttrs)
}
func (resource *Device) T_ModelNumber(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Device.ModelNumber", nil, htmlAttrs)
	}
	return StringInput("Device.ModelNumber", resource.ModelNumber, htmlAttrs)
}
func (resource *Device) T_PartNumber(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Device.PartNumber", nil, htmlAttrs)
	}
	return StringInput("Device.PartNumber", resource.PartNumber, htmlAttrs)
}
func (resource *Device) T_Type(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Device.Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Device.Type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_Url(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Device.Url", nil, htmlAttrs)
	}
	return StringInput("Device.Url", resource.Url, htmlAttrs)
}
func (resource *Device) T_Note(numNote int, htmlAttrs string) templ.Component {

	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("Device.Note."+strconv.Itoa(numNote)+".", nil, htmlAttrs)
	}
	return AnnotationTextArea("Device.Note."+strconv.Itoa(numNote)+".", &resource.Note[numNote], htmlAttrs)
}
func (resource *Device) T_Safety(numSafety int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSafety >= len(resource.Safety) {
		return CodeableConceptSelect("Device.Safety."+strconv.Itoa(numSafety)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Device.Safety."+strconv.Itoa(numSafety)+".", &resource.Safety[numSafety], optionsValueSet, htmlAttrs)
}
func (resource *Device) T_UdiCarrierDeviceIdentifier(numUdiCarrier int, htmlAttrs string) templ.Component {

	if resource == nil || numUdiCarrier >= len(resource.UdiCarrier) {
		return StringInput("Device.UdiCarrier."+strconv.Itoa(numUdiCarrier)+"..DeviceIdentifier", nil, htmlAttrs)
	}
	return StringInput("Device.UdiCarrier."+strconv.Itoa(numUdiCarrier)+"..DeviceIdentifier", resource.UdiCarrier[numUdiCarrier].DeviceIdentifier, htmlAttrs)
}
func (resource *Device) T_UdiCarrierIssuer(numUdiCarrier int, htmlAttrs string) templ.Component {

	if resource == nil || numUdiCarrier >= len(resource.UdiCarrier) {
		return StringInput("Device.UdiCarrier."+strconv.Itoa(numUdiCarrier)+"..Issuer", nil, htmlAttrs)
	}
	return StringInput("Device.UdiCarrier."+strconv.Itoa(numUdiCarrier)+"..Issuer", resource.UdiCarrier[numUdiCarrier].Issuer, htmlAttrs)
}
func (resource *Device) T_UdiCarrierJurisdiction(numUdiCarrier int, htmlAttrs string) templ.Component {

	if resource == nil || numUdiCarrier >= len(resource.UdiCarrier) {
		return StringInput("Device.UdiCarrier."+strconv.Itoa(numUdiCarrier)+"..Jurisdiction", nil, htmlAttrs)
	}
	return StringInput("Device.UdiCarrier."+strconv.Itoa(numUdiCarrier)+"..Jurisdiction", resource.UdiCarrier[numUdiCarrier].Jurisdiction, htmlAttrs)
}
func (resource *Device) T_UdiCarrierCarrierAIDC(numUdiCarrier int, htmlAttrs string) templ.Component {

	if resource == nil || numUdiCarrier >= len(resource.UdiCarrier) {
		return StringInput("Device.UdiCarrier."+strconv.Itoa(numUdiCarrier)+"..CarrierAIDC", nil, htmlAttrs)
	}
	return StringInput("Device.UdiCarrier."+strconv.Itoa(numUdiCarrier)+"..CarrierAIDC", resource.UdiCarrier[numUdiCarrier].CarrierAIDC, htmlAttrs)
}
func (resource *Device) T_UdiCarrierCarrierHRF(numUdiCarrier int, htmlAttrs string) templ.Component {

	if resource == nil || numUdiCarrier >= len(resource.UdiCarrier) {
		return StringInput("Device.UdiCarrier."+strconv.Itoa(numUdiCarrier)+"..CarrierHRF", nil, htmlAttrs)
	}
	return StringInput("Device.UdiCarrier."+strconv.Itoa(numUdiCarrier)+"..CarrierHRF", resource.UdiCarrier[numUdiCarrier].CarrierHRF, htmlAttrs)
}
func (resource *Device) T_UdiCarrierEntryType(numUdiCarrier int, htmlAttrs string) templ.Component {
	optionsValueSet := VSUdi_entry_type

	if resource == nil || numUdiCarrier >= len(resource.UdiCarrier) {
		return CodeSelect("Device.UdiCarrier."+strconv.Itoa(numUdiCarrier)+"..EntryType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Device.UdiCarrier."+strconv.Itoa(numUdiCarrier)+"..EntryType", resource.UdiCarrier[numUdiCarrier].EntryType, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_DeviceNameName(numDeviceName int, htmlAttrs string) templ.Component {

	if resource == nil || numDeviceName >= len(resource.DeviceName) {
		return StringInput("Device.DeviceName."+strconv.Itoa(numDeviceName)+"..Name", nil, htmlAttrs)
	}
	return StringInput("Device.DeviceName."+strconv.Itoa(numDeviceName)+"..Name", &resource.DeviceName[numDeviceName].Name, htmlAttrs)
}
func (resource *Device) T_DeviceNameType(numDeviceName int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDevice_nametype

	if resource == nil || numDeviceName >= len(resource.DeviceName) {
		return CodeSelect("Device.DeviceName."+strconv.Itoa(numDeviceName)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Device.DeviceName."+strconv.Itoa(numDeviceName)+"..Type", &resource.DeviceName[numDeviceName].Type, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_SpecializationSystemType(numSpecialization int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numSpecialization >= len(resource.Specialization) {
		return CodeableConceptSelect("Device.Specialization."+strconv.Itoa(numSpecialization)+"..SystemType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Device.Specialization."+strconv.Itoa(numSpecialization)+"..SystemType", &resource.Specialization[numSpecialization].SystemType, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_SpecializationVersion(numSpecialization int, htmlAttrs string) templ.Component {

	if resource == nil || numSpecialization >= len(resource.Specialization) {
		return StringInput("Device.Specialization."+strconv.Itoa(numSpecialization)+"..Version", nil, htmlAttrs)
	}
	return StringInput("Device.Specialization."+strconv.Itoa(numSpecialization)+"..Version", resource.Specialization[numSpecialization].Version, htmlAttrs)
}
func (resource *Device) T_VersionType(numVersion int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numVersion >= len(resource.Version) {
		return CodeableConceptSelect("Device.Version."+strconv.Itoa(numVersion)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Device.Version."+strconv.Itoa(numVersion)+"..Type", resource.Version[numVersion].Type, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_VersionValue(numVersion int, htmlAttrs string) templ.Component {

	if resource == nil || numVersion >= len(resource.Version) {
		return StringInput("Device.Version."+strconv.Itoa(numVersion)+"..Value", nil, htmlAttrs)
	}
	return StringInput("Device.Version."+strconv.Itoa(numVersion)+"..Value", &resource.Version[numVersion].Value, htmlAttrs)
}
func (resource *Device) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("Device.Property."+strconv.Itoa(numProperty)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Device.Property."+strconv.Itoa(numProperty)+"..Type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_PropertyValueCode(numProperty int, numValueCode int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) || numValueCode >= len(resource.Property[numProperty].ValueCode) {
		return CodeableConceptSelect("Device.Property."+strconv.Itoa(numProperty)+"..ValueCode."+strconv.Itoa(numValueCode)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Device.Property."+strconv.Itoa(numProperty)+"..ValueCode."+strconv.Itoa(numValueCode)+".", &resource.Property[numProperty].ValueCode[numValueCode], optionsValueSet, htmlAttrs)
}
