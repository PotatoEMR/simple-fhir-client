package r4

//generated with command go run ./bultaoreune -nodownload
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

func (resource *MedicinalProductPackaged) T_Id() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductPackaged.Id", nil)
	}
	return StringInput("MedicinalProductPackaged.Id", resource.Id)
}
func (resource *MedicinalProductPackaged) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductPackaged.ImplicitRules", nil)
	}
	return StringInput("MedicinalProductPackaged.ImplicitRules", resource.ImplicitRules)
}
func (resource *MedicinalProductPackaged) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("MedicinalProductPackaged.Language", nil, optionsValueSet)
	}
	return CodeSelect("MedicinalProductPackaged.Language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductPackaged) T_Description() templ.Component {

	if resource == nil {
		return StringInput("MedicinalProductPackaged.Description", nil)
	}
	return StringInput("MedicinalProductPackaged.Description", resource.Description)
}
func (resource *MedicinalProductPackaged) T_LegalStatusOfSupply(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("MedicinalProductPackaged.LegalStatusOfSupply", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductPackaged.LegalStatusOfSupply", resource.LegalStatusOfSupply, optionsValueSet)
}
func (resource *MedicinalProductPackaged) T_BatchIdentifierId(numBatchIdentifier int) templ.Component {

	if resource == nil || len(resource.BatchIdentifier) >= numBatchIdentifier {
		return StringInput("MedicinalProductPackaged.BatchIdentifier["+strconv.Itoa(numBatchIdentifier)+"].Id", nil)
	}
	return StringInput("MedicinalProductPackaged.BatchIdentifier["+strconv.Itoa(numBatchIdentifier)+"].Id", resource.BatchIdentifier[numBatchIdentifier].Id)
}
func (resource *MedicinalProductPackaged) T_PackageItemId(numPackageItem int) templ.Component {

	if resource == nil || len(resource.PackageItem) >= numPackageItem {
		return StringInput("MedicinalProductPackaged.PackageItem["+strconv.Itoa(numPackageItem)+"].Id", nil)
	}
	return StringInput("MedicinalProductPackaged.PackageItem["+strconv.Itoa(numPackageItem)+"].Id", resource.PackageItem[numPackageItem].Id)
}
func (resource *MedicinalProductPackaged) T_PackageItemType(numPackageItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PackageItem) >= numPackageItem {
		return CodeableConceptSelect("MedicinalProductPackaged.PackageItem["+strconv.Itoa(numPackageItem)+"].Type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductPackaged.PackageItem["+strconv.Itoa(numPackageItem)+"].Type", &resource.PackageItem[numPackageItem].Type, optionsValueSet)
}
func (resource *MedicinalProductPackaged) T_PackageItemMaterial(numPackageItem int, numMaterial int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PackageItem) >= numPackageItem || len(resource.PackageItem[numPackageItem].Material) >= numMaterial {
		return CodeableConceptSelect("MedicinalProductPackaged.PackageItem["+strconv.Itoa(numPackageItem)+"].Material["+strconv.Itoa(numMaterial)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductPackaged.PackageItem["+strconv.Itoa(numPackageItem)+"].Material["+strconv.Itoa(numMaterial)+"]", &resource.PackageItem[numPackageItem].Material[numMaterial], optionsValueSet)
}
func (resource *MedicinalProductPackaged) T_PackageItemAlternateMaterial(numPackageItem int, numAlternateMaterial int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PackageItem) >= numPackageItem || len(resource.PackageItem[numPackageItem].AlternateMaterial) >= numAlternateMaterial {
		return CodeableConceptSelect("MedicinalProductPackaged.PackageItem["+strconv.Itoa(numPackageItem)+"].AlternateMaterial["+strconv.Itoa(numAlternateMaterial)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductPackaged.PackageItem["+strconv.Itoa(numPackageItem)+"].AlternateMaterial["+strconv.Itoa(numAlternateMaterial)+"]", &resource.PackageItem[numPackageItem].AlternateMaterial[numAlternateMaterial], optionsValueSet)
}
func (resource *MedicinalProductPackaged) T_PackageItemOtherCharacteristics(numPackageItem int, numOtherCharacteristics int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.PackageItem) >= numPackageItem || len(resource.PackageItem[numPackageItem].OtherCharacteristics) >= numOtherCharacteristics {
		return CodeableConceptSelect("MedicinalProductPackaged.PackageItem["+strconv.Itoa(numPackageItem)+"].OtherCharacteristics["+strconv.Itoa(numOtherCharacteristics)+"]", nil, optionsValueSet)
	}
	return CodeableConceptSelect("MedicinalProductPackaged.PackageItem["+strconv.Itoa(numPackageItem)+"].OtherCharacteristics["+strconv.Itoa(numOtherCharacteristics)+"]", &resource.PackageItem[numPackageItem].OtherCharacteristics[numOtherCharacteristics], optionsValueSet)
}
