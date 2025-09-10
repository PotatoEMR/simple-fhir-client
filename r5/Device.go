package r5

//generated with command go run ./bultaoreune
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
func (resource *Device) T_DisplayName(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("displayName", nil, htmlAttrs)
	}
	return StringInput("displayName", resource.DisplayName, htmlAttrs)
}
func (resource *Device) T_Status(htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDevice_status

	if resource == nil {
		return CodeSelect("status", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("status", resource.Status, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_AvailabilityStatus(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("availabilityStatus", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("availabilityStatus", resource.AvailabilityStatus, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_Manufacturer(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("manufacturer", nil, htmlAttrs)
	}
	return StringInput("manufacturer", resource.Manufacturer, htmlAttrs)
}
func (resource *Device) T_ManufactureDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("manufactureDate", nil, htmlAttrs)
	}
	return DateTimeInput("manufactureDate", resource.ManufactureDate, htmlAttrs)
}
func (resource *Device) T_ExpirationDate(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return DateTimeInput("expirationDate", nil, htmlAttrs)
	}
	return DateTimeInput("expirationDate", resource.ExpirationDate, htmlAttrs)
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
func (resource *Device) T_Category(numCategory int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCategory >= len(resource.Category) {
		return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("category["+strconv.Itoa(numCategory)+"]", &resource.Category[numCategory], optionsValueSet, htmlAttrs)
}
func (resource *Device) T_Type(numType int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numType >= len(resource.Type) {
		return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("type["+strconv.Itoa(numType)+"]", &resource.Type[numType], optionsValueSet, htmlAttrs)
}
func (resource *Device) T_Mode(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("mode", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("mode", resource.Mode, optionsValueSet, htmlAttrs)
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
func (resource *Device) T_UdiCarrierDeviceIdentifier(numUdiCarrier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUdiCarrier >= len(resource.UdiCarrier) {
		return StringInput("udiCarrier["+strconv.Itoa(numUdiCarrier)+"].deviceIdentifier", nil, htmlAttrs)
	}
	return StringInput("udiCarrier["+strconv.Itoa(numUdiCarrier)+"].deviceIdentifier", &resource.UdiCarrier[numUdiCarrier].DeviceIdentifier, htmlAttrs)
}
func (resource *Device) T_UdiCarrierIssuer(numUdiCarrier int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numUdiCarrier >= len(resource.UdiCarrier) {
		return StringInput("udiCarrier["+strconv.Itoa(numUdiCarrier)+"].issuer", nil, htmlAttrs)
	}
	return StringInput("udiCarrier["+strconv.Itoa(numUdiCarrier)+"].issuer", &resource.UdiCarrier[numUdiCarrier].Issuer, htmlAttrs)
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
func (resource *Device) T_NameValue(numName int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return StringInput("name["+strconv.Itoa(numName)+"].value", nil, htmlAttrs)
	}
	return StringInput("name["+strconv.Itoa(numName)+"].value", &resource.Name[numName].Value, htmlAttrs)
}
func (resource *Device) T_NameType(numName int, htmlAttrs templ.Attributes) templ.Component {
	optionsValueSet := VSDevice_nametype

	if resource == nil || numName >= len(resource.Name) {
		return CodeSelect("name["+strconv.Itoa(numName)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeSelect("name["+strconv.Itoa(numName)+"].type", &resource.Name[numName].Type, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_NameDisplay(numName int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numName >= len(resource.Name) {
		return BoolInput("name["+strconv.Itoa(numName)+"].display", nil, htmlAttrs)
	}
	return BoolInput("name["+strconv.Itoa(numName)+"].display", resource.Name[numName].Display, htmlAttrs)
}
func (resource *Device) T_VersionType(numVersion int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVersion >= len(resource.Version) {
		return CodeableConceptSelect("version["+strconv.Itoa(numVersion)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("version["+strconv.Itoa(numVersion)+"].type", resource.Version[numVersion].Type, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_VersionInstallDate(numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVersion >= len(resource.Version) {
		return DateTimeInput("version["+strconv.Itoa(numVersion)+"].installDate", nil, htmlAttrs)
	}
	return DateTimeInput("version["+strconv.Itoa(numVersion)+"].installDate", resource.Version[numVersion].InstallDate, htmlAttrs)
}
func (resource *Device) T_VersionValue(numVersion int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numVersion >= len(resource.Version) {
		return StringInput("version["+strconv.Itoa(numVersion)+"].value", nil, htmlAttrs)
	}
	return StringInput("version["+strconv.Itoa(numVersion)+"].value", &resource.Version[numVersion].Value, htmlAttrs)
}
func (resource *Device) T_ConformsToCategory(numConformsTo int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConformsTo >= len(resource.ConformsTo) {
		return CodeableConceptSelect("conformsTo["+strconv.Itoa(numConformsTo)+"].category", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("conformsTo["+strconv.Itoa(numConformsTo)+"].category", resource.ConformsTo[numConformsTo].Category, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_ConformsToSpecification(numConformsTo int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConformsTo >= len(resource.ConformsTo) {
		return CodeableConceptSelect("conformsTo["+strconv.Itoa(numConformsTo)+"].specification", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("conformsTo["+strconv.Itoa(numConformsTo)+"].specification", &resource.ConformsTo[numConformsTo].Specification, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_ConformsToVersion(numConformsTo int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numConformsTo >= len(resource.ConformsTo) {
		return StringInput("conformsTo["+strconv.Itoa(numConformsTo)+"].version", nil, htmlAttrs)
	}
	return StringInput("conformsTo["+strconv.Itoa(numConformsTo)+"].version", resource.ConformsTo[numConformsTo].Version, htmlAttrs)
}
func (resource *Device) T_PropertyType(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].type", &resource.Property[numProperty].Type, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_PropertyValueCodeableConcept(numProperty int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("property["+strconv.Itoa(numProperty)+"].valueCodeableConcept", &resource.Property[numProperty].ValueCodeableConcept, optionsValueSet, htmlAttrs)
}
func (resource *Device) T_PropertyValueString(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return StringInput("property["+strconv.Itoa(numProperty)+"].valueString", nil, htmlAttrs)
	}
	return StringInput("property["+strconv.Itoa(numProperty)+"].valueString", &resource.Property[numProperty].ValueString, htmlAttrs)
}
func (resource *Device) T_PropertyValueBoolean(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", nil, htmlAttrs)
	}
	return BoolInput("property["+strconv.Itoa(numProperty)+"].valueBoolean", &resource.Property[numProperty].ValueBoolean, htmlAttrs)
}
func (resource *Device) T_PropertyValueInteger(numProperty int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numProperty >= len(resource.Property) {
		return IntInput("property["+strconv.Itoa(numProperty)+"].valueInteger", nil, htmlAttrs)
	}
	return IntInput("property["+strconv.Itoa(numProperty)+"].valueInteger", &resource.Property[numProperty].ValueInteger, htmlAttrs)
}
