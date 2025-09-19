package r4b

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4b/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

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
	ManufactureDate    *FhirDateTime          `json:"manufactureDate,omitempty"`
	ExpirationDate     *FhirDateTime          `json:"expirationDate,omitempty"`
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
func (resource *Device) T_Definition(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("definition", nil, htmlAttrs)
	}
	return ReferenceInput("definition", resource.Definition, htmlAttrs)
}
func (resource *Device) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDevice_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_StatusReason(numStatusReason int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numStatusReason >= len(resource.StatusReason) {
		return CodeableConceptSelect("statusReason["+strconv.Itoa(numStatusReason)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("statusReason["+strconv.Itoa(numStatusReason)+"]", &resource.StatusReason[numStatusReason], optionsValueSet, htmlAttrs)
}
func (resource *Device) T_DistinctIdentifier(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("distinctIdentifier", nil, htmlAttrs)
	}
	return StringInput("distinctIdentifier", resource.DistinctIdentifier, htmlAttrs)
}
func (resource *Device) T_Manufacturer(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("manufacturer", nil, htmlAttrs)
	}
	return StringInput("manufacturer", resource.Manufacturer, htmlAttrs)
}
func (resource *Device) T_ManufactureDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("manufactureDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("manufactureDate", resource.ManufactureDate, htmlAttrs)
}
func (resource *Device) T_ExpirationDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return FhirDateTimeInput("expirationDate", nil, htmlAttrs)
	}
	return FhirDateTimeInput("expirationDate", resource.ExpirationDate, htmlAttrs)
}
func (resource *Device) T_LotNumber(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("lotNumber", nil, htmlAttrs)
	}
	return StringInput("lotNumber", resource.LotNumber, htmlAttrs)
}
func (resource *Device) T_SerialNumber(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("serialNumber", nil, htmlAttrs)
	}
	return StringInput("serialNumber", resource.SerialNumber, htmlAttrs)
}
func (resource *Device) T_ModelNumber(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("modelNumber", nil, htmlAttrs)
	}
	return StringInput("modelNumber", resource.ModelNumber, htmlAttrs)
}
func (resource *Device) T_PartNumber(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("partNumber", nil, htmlAttrs)
	}
	return StringInput("partNumber", resource.PartNumber, htmlAttrs)
}
func (resource *Device) T_Type(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type", resource.Type, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_Patient(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("patient", nil, htmlAttrs)
	}
	return ReferenceInput("patient", resource.Patient, htmlAttrs)
}
func (resource *Device) T_Owner(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("owner", nil, htmlAttrs)
	}
	return ReferenceInput("owner", resource.Owner, htmlAttrs)
}
func (resource *Device) T_Contact(numContact int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numContact >= len(resource.Contact) {
		return ContactPointInput("contact["+strconv.Itoa(numContact)+"]", nil, htmlAttrs)
	}
	return ContactPointInput("contact["+strconv.Itoa(numContact)+"]", &resource.Contact[numContact], htmlAttrs)
}
func (resource *Device) T_Location(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("location", nil, htmlAttrs)
	}
	return ReferenceInput("location", resource.Location, htmlAttrs)
}
func (resource *Device) T_Url(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("url", nil, htmlAttrs)
	}
	return StringInput("url", resource.Url, htmlAttrs)
}
func (resource *Device) T_Note(numNote int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numNote >= len(resource.Note) {
		return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", nil, htmlAttrs)
	}
	return AnnotationTextArea("note["+strconv.Itoa(numNote)+"]", &resource.Note[numNote], htmlAttrs)
}
func (resource *Device) T_Safety(numSafety int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSafety >= len(resource.Safety) {
		return CodeableConceptSelect("safety["+strconv.Itoa(numSafety)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("safety["+strconv.Itoa(numSafety)+"]", &resource.Safety[numSafety], optionsValueSet, htmlAttrs)
}
func (resource *Device) T_Parent(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return ReferenceInput("parent", nil, htmlAttrs)
	}
	return ReferenceInput("parent", resource.Parent, htmlAttrs)
}
func (resource *Device) T_UdiCarrierDeviceIdentifier(numUdiCarrier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUdiCarrier >= len(resource.UdiCarrier) {
		return StringInput("udiCarrier["+strconv.Itoa(numUdiCarrier)+"].deviceIdentifier", nil, htmlAttrs)
	}
	return StringInput("udiCarrier["+strconv.Itoa(numUdiCarrier)+"].deviceIdentifier", resource.UdiCarrier[numUdiCarrier].DeviceIdentifier, htmlAttrs)
}
func (resource *Device) T_UdiCarrierIssuer(numUdiCarrier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUdiCarrier >= len(resource.UdiCarrier) {
		return StringInput("udiCarrier["+strconv.Itoa(numUdiCarrier)+"].issuer", nil, htmlAttrs)
	}
	return StringInput("udiCarrier["+strconv.Itoa(numUdiCarrier)+"].issuer", resource.UdiCarrier[numUdiCarrier].Issuer, htmlAttrs)
}
func (resource *Device) T_UdiCarrierJurisdiction(numUdiCarrier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUdiCarrier >= len(resource.UdiCarrier) {
		return StringInput("udiCarrier["+strconv.Itoa(numUdiCarrier)+"].jurisdiction", nil, htmlAttrs)
	}
	return StringInput("udiCarrier["+strconv.Itoa(numUdiCarrier)+"].jurisdiction", resource.UdiCarrier[numUdiCarrier].Jurisdiction, htmlAttrs)
}
func (resource *Device) T_UdiCarrierCarrierAIDC(numUdiCarrier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUdiCarrier >= len(resource.UdiCarrier) {
		return StringInput("udiCarrier["+strconv.Itoa(numUdiCarrier)+"].carrierAIDC", nil, htmlAttrs)
	}
	return StringInput("udiCarrier["+strconv.Itoa(numUdiCarrier)+"].carrierAIDC", resource.UdiCarrier[numUdiCarrier].CarrierAIDC, htmlAttrs)
}
func (resource *Device) T_UdiCarrierCarrierHRF(numUdiCarrier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUdiCarrier >= len(resource.UdiCarrier) {
		return StringInput("udiCarrier["+strconv.Itoa(numUdiCarrier)+"].carrierHRF", nil, htmlAttrs)
	}
	return StringInput("udiCarrier["+strconv.Itoa(numUdiCarrier)+"].carrierHRF", resource.UdiCarrier[numUdiCarrier].CarrierHRF, htmlAttrs)
}
func (resource *Device) T_UdiCarrierEntryType(numUdiCarrier int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSUdi_entry_type

	if resource == nil || numUdiCarrier >= len(resource.UdiCarrier) {
		return CodeSelect("udiCarrier["+strconv.Itoa(numUdiCarrier)+"].entryType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("udiCarrier["+strconv.Itoa(numUdiCarrier)+"].entryType", resource.UdiCarrier[numUdiCarrier].EntryType, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_DeviceNameName(numDeviceName int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numDeviceName >= len(resource.DeviceName) {
		return StringInput("deviceName["+strconv.Itoa(numDeviceName)+"].name", nil, htmlAttrs)
	}
	return StringInput("deviceName["+strconv.Itoa(numDeviceName)+"].name", &resource.DeviceName[numDeviceName].Name, htmlAttrs)
}
func (resource *Device) T_DeviceNameType(numDeviceName int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDevice_nametype

	if resource == nil || numDeviceName >= len(resource.DeviceName) {
		return CodeSelect("deviceName["+strconv.Itoa(numDeviceName)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("deviceName["+strconv.Itoa(numDeviceName)+"].type", &resource.DeviceName[numDeviceName].Type, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_SpecializationSystemType(numSpecialization int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecialization >= len(resource.Specialization) {
		return CodeableConceptSelect("specialization["+strconv.Itoa(numSpecialization)+"].systemType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("specialization["+strconv.Itoa(numSpecialization)+"].systemType", &resource.Specialization[numSpecialization].SystemType, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_SpecializationVersion(numSpecialization int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSpecialization >= len(resource.Specialization) {
		return StringInput("specialization["+strconv.Itoa(numSpecialization)+"].version", nil, htmlAttrs)
	}
	return StringInput("specialization["+strconv.Itoa(numSpecialization)+"].version", resource.Specialization[numSpecialization].Version, htmlAttrs)
}
func (resource *Device) T_VersionType(numVersion int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVersion >= len(resource.Version) {
		return CodeableConceptSelect("version["+strconv.Itoa(numVersion)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("version["+strconv.Itoa(numVersion)+"].type", resource.Version[numVersion].Type, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_VersionComponent(numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVersion >= len(resource.Version) {
		return IdentifierInput("version["+strconv.Itoa(numVersion)+"].component", nil, htmlAttrs)
	}
	return IdentifierInput("version["+strconv.Itoa(numVersion)+"].component", resource.Version[numVersion].Component, htmlAttrs)
}
func (resource *Device) T_VersionValue(numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVersion >= len(resource.Version) {
		return StringInput("version["+strconv.Itoa(numVersion)+"].value", nil, htmlAttrs)
	}
	return StringInput("version["+strconv.Itoa(numVersion)+"].value", &resource.Version[numVersion].Value, htmlAttrs)
}
func (resource *Device) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_PropertyValueQuantity(numProperty int, numValueQuantity int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) || numValueQuantity >= len(resource.Property[numProperty].ValueQuantity) {
		return QuantityInput("property["+strconv.Itoa(numProperty)+"].valueQuantity["+strconv.Itoa(numValueQuantity)+"]", nil, htmlAttrs)
	}
	return QuantityInput("property["+strconv.Itoa(numProperty)+"].valueQuantity["+strconv.Itoa(numValueQuantity)+"]", &resource.Property[numProperty].ValueQuantity[numValueQuantity], htmlAttrs)
}
func (resource *Device) T_PropertyValueCode(numProperty int, numValueCode int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) || numValueCode >= len(resource.Property[numProperty].ValueCode) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCode["+strconv.Itoa(numValueCode)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCode["+strconv.Itoa(numValueCode)+"]", &resource.Property[numProperty].ValueCode[numValueCode], optionsValueSet, htmlAttrs)
}
