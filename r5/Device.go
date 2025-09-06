package r5

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/Device
type Device struct {
	Id                    *string             `json:"id,omitempty"`
	Meta                  *Meta               `json:"meta,omitempty"`
	ImplicitRules         *string             `json:"implicitRules,omitempty"`
	Language              *string             `json:"language,omitempty"`
	Text                  *Narrative          `json:"text,omitempty"`
	Contained             []Resource          `json:"contained,omitempty"`
	Extension             []Extension         `json:"extension,omitempty"`
	ModifierExtension     []Extension         `json:"modifierExtension,omitempty"`
	Identifier            []Identifier        `json:"identifier,omitempty"`
	DisplayName           *string             `json:"displayName,omitempty"`
	Definition            *CodeableReference  `json:"definition,omitempty"`
	UdiCarrier            []DeviceUdiCarrier  `json:"udiCarrier,omitempty"`
	Status                *string             `json:"status,omitempty"`
	AvailabilityStatus    *CodeableConcept    `json:"availabilityStatus,omitempty"`
	BiologicalSourceEvent *Identifier         `json:"biologicalSourceEvent,omitempty"`
	Manufacturer          *string             `json:"manufacturer,omitempty"`
	ManufactureDate       *string             `json:"manufactureDate,omitempty"`
	ExpirationDate        *string             `json:"expirationDate,omitempty"`
	LotNumber             *string             `json:"lotNumber,omitempty"`
	SerialNumber          *string             `json:"serialNumber,omitempty"`
	Name                  []DeviceName        `json:"name,omitempty"`
	ModelNumber           *string             `json:"modelNumber,omitempty"`
	PartNumber            *string             `json:"partNumber,omitempty"`
	Category              []CodeableConcept   `json:"category,omitempty"`
	Type                  []CodeableConcept   `json:"type,omitempty"`
	Version               []DeviceVersion     `json:"version,omitempty"`
	ConformsTo            []DeviceConformsTo  `json:"conformsTo,omitempty"`
	Property              []DeviceProperty    `json:"property,omitempty"`
	Mode                  *CodeableConcept    `json:"mode,omitempty"`
	Cycle                 *Count              `json:"cycle,omitempty"`
	Duration              *Duration           `json:"duration,omitempty"`
	Owner                 *Reference          `json:"owner,omitempty"`
	Contact               []ContactPoint      `json:"contact,omitempty"`
	Location              *Reference          `json:"location,omitempty"`
	Url                   *string             `json:"url,omitempty"`
	Endpoint              []Reference         `json:"endpoint,omitempty"`
	Gateway               []CodeableReference `json:"gateway,omitempty"`
	Note                  []Annotation        `json:"note,omitempty"`
	Safety                []CodeableConcept   `json:"safety,omitempty"`
	Parent                *Reference          `json:"parent,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Device
type DeviceUdiCarrier struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	DeviceIdentifier  string      `json:"deviceIdentifier"`
	Issuer            string      `json:"issuer"`
	Jurisdiction      *string     `json:"jurisdiction,omitempty"`
	CarrierAIDC       *string     `json:"carrierAIDC,omitempty"`
	CarrierHRF        *string     `json:"carrierHRF,omitempty"`
	EntryType         *string     `json:"entryType,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Device
type DeviceName struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Value             string      `json:"value"`
	Type              string      `json:"type"`
	Display           *bool       `json:"display,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Device
type DeviceVersion struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Component         *Identifier      `json:"component,omitempty"`
	InstallDate       *string          `json:"installDate,omitempty"`
	Value             string           `json:"value"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Device
type DeviceConformsTo struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Category          *CodeableConcept `json:"category,omitempty"`
	Specification     CodeableConcept  `json:"specification"`
	Version           *string          `json:"version,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/Device
type DeviceProperty struct {
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
func (resource *Device) T_DisplayName() templ.Component {

	if resource == nil {
		return StringInput("Device.DisplayName", nil)
	}
	return StringInput("Device.DisplayName", resource.DisplayName)
}
func (resource *Device) T_Status() templ.Component {
	optionsValueSet := VSDevice_status

	if resource == nil {
		return CodeSelect("Device.Status", nil, optionsValueSet)
	}
	return CodeSelect("Device.Status", resource.Status, optionsValueSet)
}
func (resource *Device) T_AvailabilityStatus(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Device.AvailabilityStatus", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Device.AvailabilityStatus", resource.AvailabilityStatus, optionsValueSet)
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
func (resource *Device) T_Category(numCategory int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Category) >= numCategory {
		return CodeableConceptSelect("Device.Category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Device.Category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet)
}
func (resource *Device) T_Type(numType int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Type) >= numType {
		return CodeableConceptSelect("Device.Type["+strconv.Itoa(numType)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Device.Type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet)
}
func (resource *Device) T_Mode(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Device.Mode", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Device.Mode", resource.Mode, optionsValueSet)
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
	return StringInput("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].DeviceIdentifier", &resource.UdiCarrier[numUdiCarrier].DeviceIdentifier)
}
func (resource *Device) T_UdiCarrierIssuer(numUdiCarrier int) templ.Component {

	if resource == nil || len(resource.UdiCarrier) >= numUdiCarrier {
		return StringInput("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].Issuer", nil)
	}
	return StringInput("Device.UdiCarrier["+strconv.Itoa(numUdiCarrier)+"].Issuer", &resource.UdiCarrier[numUdiCarrier].Issuer)
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
func (resource *Device) T_NameId(numName int) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return StringInput("Device.Name["+strconv.Itoa(numName)+"].Id", nil)
	}
	return StringInput("Device.Name["+strconv.Itoa(numName)+"].Id", resource.Name[numName].Id)
}
func (resource *Device) T_NameValue(numName int) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return StringInput("Device.Name["+strconv.Itoa(numName)+"].Value", nil)
	}
	return StringInput("Device.Name["+strconv.Itoa(numName)+"].Value", &resource.Name[numName].Value)
}
func (resource *Device) T_NameType(numName int) templ.Component {
	optionsValueSet := VSDevice_nametype

	if resource == nil || len(resource.Name) >= numName {
		return CodeSelect("Device.Name["+strconv.Itoa(numName)+"].Type", nil, optionsValueSet)
	}
	return CodeSelect("Device.Name["+strconv.Itoa(numName)+"].Type", &resource.Name[numName].Type, optionsValueSet)
}
func (resource *Device) T_NameDisplay(numName int) templ.Component {

	if resource == nil || len(resource.Name) >= numName {
		return BoolInput("Device.Name["+strconv.Itoa(numName)+"].Display", nil)
	}
	return BoolInput("Device.Name["+strconv.Itoa(numName)+"].Display", resource.Name[numName].Display)
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
func (resource *Device) T_VersionInstallDate(numVersion int) templ.Component {

	if resource == nil || len(resource.Version) >= numVersion {
		return StringInput("Device.Version["+strconv.Itoa(numVersion)+"].InstallDate", nil)
	}
	return StringInput("Device.Version["+strconv.Itoa(numVersion)+"].InstallDate", resource.Version[numVersion].InstallDate)
}
func (resource *Device) T_VersionValue(numVersion int) templ.Component {

	if resource == nil || len(resource.Version) >= numVersion {
		return StringInput("Device.Version["+strconv.Itoa(numVersion)+"].Value", nil)
	}
	return StringInput("Device.Version["+strconv.Itoa(numVersion)+"].Value", &resource.Version[numVersion].Value)
}
func (resource *Device) T_ConformsToId(numConformsTo int) templ.Component {

	if resource == nil || len(resource.ConformsTo) >= numConformsTo {
		return StringInput("Device.ConformsTo["+strconv.Itoa(numConformsTo)+"].Id", nil)
	}
	return StringInput("Device.ConformsTo["+strconv.Itoa(numConformsTo)+"].Id", resource.ConformsTo[numConformsTo].Id)
}
func (resource *Device) T_ConformsToCategory(numConformsTo int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ConformsTo) >= numConformsTo {
		return CodeableConceptSelect("Device.ConformsTo["+strconv.Itoa(numConformsTo)+"].Category", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Device.ConformsTo["+strconv.Itoa(numConformsTo)+"].Category", resource.ConformsTo[numConformsTo].Category, optionsValueSet)
}
func (resource *Device) T_ConformsToSpecification(numConformsTo int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.ConformsTo) >= numConformsTo {
		return CodeableConceptSelect("Device.ConformsTo["+strconv.Itoa(numConformsTo)+"].Specification", nil, optionsValueSet)
	}
	return CodeableConceptSelect("Device.ConformsTo["+strconv.Itoa(numConformsTo)+"].Specification", &resource.ConformsTo[numConformsTo].Specification, optionsValueSet)
}
func (resource *Device) T_ConformsToVersion(numConformsTo int) templ.Component {

	if resource == nil || len(resource.ConformsTo) >= numConformsTo {
		return StringInput("Device.ConformsTo["+strconv.Itoa(numConformsTo)+"].Version", nil)
	}
	return StringInput("Device.ConformsTo["+strconv.Itoa(numConformsTo)+"].Version", resource.ConformsTo[numConformsTo].Version)
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
