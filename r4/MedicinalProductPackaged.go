package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductPackaged
type MedicinalProductPackaged struct {
	Id                     *string                                   `json:"id,omitempty"`
	Meta                   *Meta                                     `json:"meta,omitempty"`
	ImplicitRules          *string                                   `json:"implicitRules,omitempty"`
	Language               *string                                   `json:"language,omitempty"`
	Text                   *Narrative                                `json:"text,omitempty"`
	Contained              []Resource                                `json:"contained,omitempty"`
	Extension              []Extension                               `json:"extension,omitempty"`
	ModifierExtension      []Extension                               `json:"modifierExtension,omitempty"`
	Identifier             []Identifier                              `json:"identifier,omitempty"`
	Subject                []Reference                               `json:"subject,omitempty"`
	Description            *string                                   `json:"description,omitempty"`
	LegalStatusOfSupply    *CodeableConcept                          `json:"legalStatusOfSupply,omitempty"`
	MarketingStatus        []MarketingStatus                         `json:"marketingStatus,omitempty"`
	MarketingAuthorization *Reference                                `json:"marketingAuthorization,omitempty"`
	Manufacturer           []Reference                               `json:"manufacturer,omitempty"`
	BatchIdentifier        []MedicinalProductPackagedBatchIdentifier `json:"batchIdentifier,omitempty"`
	PackageItem            []MedicinalProductPackagedPackageItem     `json:"packageItem"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductPackaged
type MedicinalProductPackagedBatchIdentifier struct {
	Id                 *string     `json:"id,omitempty"`
	Extension          []Extension `json:"extension,omitempty"`
	ModifierExtension  []Extension `json:"modifierExtension,omitempty"`
	OuterPackaging     Identifier  `json:"outerPackaging"`
	ImmediatePackaging *Identifier `json:"immediatePackaging,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/MedicinalProductPackaged
type MedicinalProductPackagedPackageItem struct {
	Id                      *string             `json:"id,omitempty"`
	Extension               []Extension         `json:"extension,omitempty"`
	ModifierExtension       []Extension         `json:"modifierExtension,omitempty"`
	Identifier              []Identifier        `json:"identifier,omitempty"`
	Type                    CodeableConcept     `json:"type"`
	Quantity                Quantity            `json:"quantity"`
	Material                []CodeableConcept   `json:"material,omitempty"`
	AlternateMaterial       []CodeableConcept   `json:"alternateMaterial,omitempty"`
	Device                  []Reference         `json:"device,omitempty"`
	ManufacturedItem        []Reference         `json:"manufacturedItem,omitempty"`
	PhysicalCharacteristics *ProdCharacteristic `json:"physicalCharacteristics,omitempty"`
	OtherCharacteristics    []CodeableConcept   `json:"otherCharacteristics,omitempty"`
	ShelfLifeStorage        []ProductShelfLife  `json:"shelfLifeStorage,omitempty"`
	Manufacturer            []Reference         `json:"manufacturer,omitempty"`
}

type OtherMedicinalProductPackaged MedicinalProductPackaged

// on convert struct to json, automatically add resourceType=MedicinalProductPackaged
func (r MedicinalProductPackaged) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductPackaged
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductPackaged: OtherMedicinalProductPackaged(r),
		ResourceType:                  "MedicinalProductPackaged",
	})
}
func (r MedicinalProductPackaged) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "MedicinalProductPackaged/" + *r.Id
		ref.Reference = &refStr
	}
	if len(r.Identifier) != 0 {
		ref.Identifier = &r.Identifier[0]
	}
	rtype := "MedicinalProductPackaged"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *MedicinalProductPackaged) T_Description(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("description", nil, htmlAttrs)
	}
	return StringInput("description", resource.Description, htmlAttrs)
}
func (resource *MedicinalProductPackaged) T_LegalStatusOfSupply(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("legalStatusOfSupply", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("legalStatusOfSupply", resource.LegalStatusOfSupply, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPackaged) T_PackageItemType(numPackageItem int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPackageItem >= len(resource.PackageItem) {
		return CodeableConceptSelect("packageItem["+strconv.Itoa(numPackageItem)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("packageItem["+strconv.Itoa(numPackageItem)+"].type", &resource.PackageItem[numPackageItem].Type, optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPackaged) T_PackageItemMaterial(numPackageItem int, numMaterial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPackageItem >= len(resource.PackageItem) || numMaterial >= len(resource.PackageItem[numPackageItem].Material) {
		return CodeableConceptSelect("packageItem["+strconv.Itoa(numPackageItem)+"].material["+strconv.Itoa(numMaterial)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("packageItem["+strconv.Itoa(numPackageItem)+"].material["+strconv.Itoa(numMaterial)+"]", &resource.PackageItem[numPackageItem].Material[numMaterial], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPackaged) T_PackageItemAlternateMaterial(numPackageItem int, numAlternateMaterial int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPackageItem >= len(resource.PackageItem) || numAlternateMaterial >= len(resource.PackageItem[numPackageItem].AlternateMaterial) {
		return CodeableConceptSelect("packageItem["+strconv.Itoa(numPackageItem)+"].alternateMaterial["+strconv.Itoa(numAlternateMaterial)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("packageItem["+strconv.Itoa(numPackageItem)+"].alternateMaterial["+strconv.Itoa(numAlternateMaterial)+"]", &resource.PackageItem[numPackageItem].AlternateMaterial[numAlternateMaterial], optionsValueSet, htmlAttrs)
}
func (resource *MedicinalProductPackaged) T_PackageItemOtherCharacteristics(numPackageItem int, numOtherCharacteristics int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numPackageItem >= len(resource.PackageItem) || numOtherCharacteristics >= len(resource.PackageItem[numPackageItem].OtherCharacteristics) {
		return CodeableConceptSelect("packageItem["+strconv.Itoa(numPackageItem)+"].otherCharacteristics["+strconv.Itoa(numOtherCharacteristics)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("packageItem["+strconv.Itoa(numPackageItem)+"].otherCharacteristics["+strconv.Itoa(numOtherCharacteristics)+"]", &resource.PackageItem[numPackageItem].OtherCharacteristics[numOtherCharacteristics], optionsValueSet, htmlAttrs)
}
