package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"
	"time"

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
	ManufactureDate       *time.Time          `json:"manufactureDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
	ExpirationDate        *time.Time          `json:"expirationDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
	InstallDate       *time.Time       `json:"installDate,omitempty,format:'2006-01-02T15:04:05Z07:00'"`
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
func (resource *Device) T_DisplayName(htmlAttrs string) templ.Component {

	if resource == nil {
		return StringInput("Device.DisplayName", nil, htmlAttrs)
	}
	return StringInput("Device.DisplayName", resource.DisplayName, htmlAttrs)
}
func (resource *Device) T_Status(htmlAttrs string) templ.Component {
	optionsValueSet := VSDevice_status

	if resource == nil {
		return CodeSelect("Device.Status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Device.Status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_AvailabilityStatus(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Device.AvailabilityStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Device.AvailabilityStatus", resource.AvailabilityStatus, optionsValueSet, htmlAttrs)
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
func (resource *Device) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("Device.Category."+strconv.Itoa(numCategory)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Device.Category."+strconv.Itoa(numCategory)+".", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Device) T_Type(numType int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("Device.Type."+strconv.Itoa(numType)+".", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Device.Type."+strconv.Itoa(numType)+".", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Device) T_Mode(optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("Device.Mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Device.Mode", resource.Mode, optionsValueSet, htmlAttrs)
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
	return StringInput("Device.UdiCarrier."+strconv.Itoa(numUdiCarrier)+"..DeviceIdentifier", &resource.UdiCarrier[numUdiCarrier].DeviceIdentifier, htmlAttrs)
}
func (resource *Device) T_UdiCarrierIssuer(numUdiCarrier int, htmlAttrs string) templ.Component {

	if resource == nil || numUdiCarrier >= len(resource.UdiCarrier) {
		return StringInput("Device.UdiCarrier."+strconv.Itoa(numUdiCarrier)+"..Issuer", nil, htmlAttrs)
	}
	return StringInput("Device.UdiCarrier."+strconv.Itoa(numUdiCarrier)+"..Issuer", &resource.UdiCarrier[numUdiCarrier].Issuer, htmlAttrs)
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
func (resource *Device) T_NameValue(numName int, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) {
		return StringInput("Device.Name."+strconv.Itoa(numName)+"..Value", nil, htmlAttrs)
	}
	return StringInput("Device.Name."+strconv.Itoa(numName)+"..Value", &resource.Name[numName].Value, htmlAttrs)
}
func (resource *Device) T_NameType(numName int, htmlAttrs string) templ.Component {
	optionsValueSet := VSDevice_nametype

	if resource == nil || numName >= len(resource.Name) {
		return CodeSelect("Device.Name."+strconv.Itoa(numName)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("Device.Name."+strconv.Itoa(numName)+"..Type", &resource.Name[numName].Type, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_NameDisplay(numName int, htmlAttrs string) templ.Component {

	if resource == nil || numName >= len(resource.Name) {
		return BoolInput("Device.Name."+strconv.Itoa(numName)+"..Display", nil, htmlAttrs)
	}
	return BoolInput("Device.Name."+strconv.Itoa(numName)+"..Display", resource.Name[numName].Display, htmlAttrs)
}
func (resource *Device) T_VersionType(numVersion int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numVersion >= len(resource.Version) {
		return CodeableConceptSelect("Device.Version."+strconv.Itoa(numVersion)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Device.Version."+strconv.Itoa(numVersion)+"..Type", resource.Version[numVersion].Type, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_VersionInstallDate(numVersion int, htmlAttrs string) templ.Component {

	if resource == nil || numVersion >= len(resource.Version) {
		return DateTimeInput("Device.Version."+strconv.Itoa(numVersion)+"..InstallDate", nil, htmlAttrs)
	}
	return DateTimeInput("Device.Version."+strconv.Itoa(numVersion)+"..InstallDate", resource.Version[numVersion].InstallDate, htmlAttrs)
}
func (resource *Device) T_VersionValue(numVersion int, htmlAttrs string) templ.Component {

	if resource == nil || numVersion >= len(resource.Version) {
		return StringInput("Device.Version."+strconv.Itoa(numVersion)+"..Value", nil, htmlAttrs)
	}
	return StringInput("Device.Version."+strconv.Itoa(numVersion)+"..Value", &resource.Version[numVersion].Value, htmlAttrs)
}
func (resource *Device) T_ConformsToCategory(numConformsTo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numConformsTo >= len(resource.ConformsTo) {
		return CodeableConceptSelect("Device.ConformsTo."+strconv.Itoa(numConformsTo)+"..Category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Device.ConformsTo."+strconv.Itoa(numConformsTo)+"..Category", resource.ConformsTo[numConformsTo].Category, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_ConformsToSpecification(numConformsTo int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numConformsTo >= len(resource.ConformsTo) {
		return CodeableConceptSelect("Device.ConformsTo."+strconv.Itoa(numConformsTo)+"..Specification", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Device.ConformsTo."+strconv.Itoa(numConformsTo)+"..Specification", &resource.ConformsTo[numConformsTo].Specification, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_ConformsToVersion(numConformsTo int, htmlAttrs string) templ.Component {

	if resource == nil || numConformsTo >= len(resource.ConformsTo) {
		return StringInput("Device.ConformsTo."+strconv.Itoa(numConformsTo)+"..Version", nil, htmlAttrs)
	}
	return StringInput("Device.ConformsTo."+strconv.Itoa(numConformsTo)+"..Version", resource.ConformsTo[numConformsTo].Version, htmlAttrs)
}
func (resource *Device) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("Device.Property."+strconv.Itoa(numProperty)+"..Type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Device.Property."+strconv.Itoa(numProperty)+"..Type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_PropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("Device.Property."+strconv.Itoa(numProperty)+"..ValueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Device.Property."+strconv.Itoa(numProperty)+"..ValueCodeableConcept", &resource.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_PropertyValueString(numProperty int, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("Device.Property."+strconv.Itoa(numProperty)+"..ValueString", nil, htmlAttrs)
	}
	return StringInput("Device.Property."+strconv.Itoa(numProperty)+"..ValueString", &resource.Property[numProperty].ValueString, htmlAttrs)
}
func (resource *Device) T_PropertyValueBoolean(numProperty int, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return BoolInput("Device.Property."+strconv.Itoa(numProperty)+"..ValueBoolean", nil, htmlAttrs)
	}
	return BoolInput("Device.Property."+strconv.Itoa(numProperty)+"..ValueBoolean", &resource.Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *Device) T_PropertyValueInteger(numProperty int, htmlAttrs string) templ.Component {

	if resource == nil || numProperty >= len(resource.Property) {
		return IntInput("Device.Property."+strconv.Itoa(numProperty)+"..ValueInteger", nil, htmlAttrs)
	}
	return IntInput("Device.Property."+strconv.Itoa(numProperty)+"..ValueInteger", &resource.Property[numProperty].ValueInteger, htmlAttrs)
}
