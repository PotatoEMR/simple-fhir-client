package r5

//generated August 28 2025 with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceNucleicAcid
type SubstanceNucleicAcid struct {
	Id                  *string                       `json:"id,omitempty"`
	Meta                *Meta                         `json:"meta,omitempty"`
	ImplicitRules       *string                       `json:"implicitRules,omitempty"`
	Language            *string                       `json:"language,omitempty"`
	Text                *Narrative                    `json:"text,omitempty"`
	Contained           []Resource                    `json:"contained,omitempty"`
	Extension           []Extension                   `json:"extension,omitempty"`
	ModifierExtension   []Extension                   `json:"modifierExtension,omitempty"`
	SequenceType        *CodeableConcept              `json:"sequenceType,omitempty"`
	NumberOfSubunits    *int                          `json:"numberOfSubunits,omitempty"`
	AreaOfHybridisation *string                       `json:"areaOfHybridisation,omitempty"`
	OligoNucleotideType *CodeableConcept              `json:"oligoNucleotideType,omitempty"`
	Subunit             []SubstanceNucleicAcidSubunit `json:"subunit,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceNucleicAcid
type SubstanceNucleicAcidSubunit struct {
	Id                 *string                              `json:"id,omitempty"`
	Extension          []Extension                          `json:"extension,omitempty"`
	ModifierExtension  []Extension                          `json:"modifierExtension,omitempty"`
	Subunit            *int                                 `json:"subunit,omitempty"`
	Sequence           *string                              `json:"sequence,omitempty"`
	Length             *int                                 `json:"length,omitempty"`
	SequenceAttachment *Attachment                          `json:"sequenceAttachment,omitempty"`
	FivePrime          *CodeableConcept                     `json:"fivePrime,omitempty"`
	ThreePrime         *CodeableConcept                     `json:"threePrime,omitempty"`
	Linkage            []SubstanceNucleicAcidSubunitLinkage `json:"linkage,omitempty"`
	Sugar              []SubstanceNucleicAcidSubunitSugar   `json:"sugar,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceNucleicAcid
type SubstanceNucleicAcidSubunitLinkage struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Connectivity      *string     `json:"connectivity,omitempty"`
	Identifier        *Identifier `json:"identifier,omitempty"`
	Name              *string     `json:"name,omitempty"`
	ResidueSite       *string     `json:"residueSite,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceNucleicAcid
type SubstanceNucleicAcidSubunitSugar struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Identifier        *Identifier `json:"identifier,omitempty"`
	Name              *string     `json:"name,omitempty"`
	ResidueSite       *string     `json:"residueSite,omitempty"`
}

type OtherSubstanceNucleicAcid SubstanceNucleicAcid

// on convert struct to json, automatically add resourceType=SubstanceNucleicAcid
func (r SubstanceNucleicAcid) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceNucleicAcid
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceNucleicAcid: OtherSubstanceNucleicAcid(r),
		ResourceType:              "SubstanceNucleicAcid",
	})
}

func (resource *SubstanceNucleicAcid) SubstanceNucleicAcidLanguage(optionsValueSet []Coding) templ.Component {
	currentVal := ""
	if resource != nil {
		currentVal = *resource.Language
	}
	return CodeSelect("language", currentVal, optionsValueSet)
}
