package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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
func (resource *Device) DeviceLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
func (resource *Device) DeviceStatus() templ.Component {
	optionsValueSet := VSDevice_status
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Status
	}
	return CodeSelect("status", currentVal, optionsValueSet)
}
func (resource *Device) DeviceUdiCarrierEntryType(numUdiCarrier int) templ.Component {
	optionsValueSet := VSUdi_entry_type
	currentVal := ""
	if resource != nil && len(resource.UdiCarrier) >= numUdiCarrier {
		currentVal = *resource.UdiCarrier[numUdiCarrier].EntryType
	}
	return CodeSelect("entryType", currentVal, optionsValueSet)
}
func (resource *Device) DeviceNameType(numName int) templ.Component {
	optionsValueSet := VSDevice_nametype
	currentVal := ""
	if resource != nil && len(resource.Name) >= numName {
		currentVal = resource.Name[numName].Type
	}
	return CodeSelect("type", currentVal, optionsValueSet)
}
