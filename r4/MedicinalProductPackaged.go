package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *MedicinalProductPackaged) MedicinalProductPackagedLanguage(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *MedicinalProductPackaged) MedicinalProductPackagedLegalStatusOfSupply(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("legalStatusOfSupply", nil, optionsValueSet)
	}
	return CodeableConceptSelect("legalStatusOfSupply", resource.LegalStatusOfSupply, optionsValueSet)
}
func (resource *MedicinalProductPackaged) MedicinalProductPackagedPackageItemType(numPackageItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.PackageItem) >= numPackageItem {
		return CodeableConceptSelect("type", nil, optionsValueSet)
	}
	return CodeableConceptSelect("type", &resource.PackageItem[numPackageItem].Type, optionsValueSet)
}
func (resource *MedicinalProductPackaged) MedicinalProductPackagedPackageItemMaterial(numPackageItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.PackageItem) >= numPackageItem {
		return CodeableConceptSelect("material", nil, optionsValueSet)
	}
	return CodeableConceptSelect("material", &resource.PackageItem[numPackageItem].Material[0], optionsValueSet)
}
func (resource *MedicinalProductPackaged) MedicinalProductPackagedPackageItemAlternateMaterial(numPackageItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.PackageItem) >= numPackageItem {
		return CodeableConceptSelect("alternateMaterial", nil, optionsValueSet)
	}
	return CodeableConceptSelect("alternateMaterial", &resource.PackageItem[numPackageItem].AlternateMaterial[0], optionsValueSet)
}
func (resource *MedicinalProductPackaged) MedicinalProductPackagedPackageItemOtherCharacteristics(numPackageItem int, optionsValueSet []Coding) templ.Component {

	if resource == nil && len(resource.PackageItem) >= numPackageItem {
		return CodeableConceptSelect("otherCharacteristics", nil, optionsValueSet)
	}
	return CodeableConceptSelect("otherCharacteristics", &resource.PackageItem[numPackageItem].OtherCharacteristics[0], optionsValueSet)
}
