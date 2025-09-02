package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceNucleicAcid
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

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceNucleicAcid
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

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceNucleicAcid
type SubstanceNucleicAcidSubunitLinkage struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Connectivity      *string     `json:"connectivity,omitempty"`
	Identifier        *Identifier `json:"identifier,omitempty"`
	Name              *string     `json:"name,omitempty"`
	ResidueSite       *string     `json:"residueSite,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceNucleicAcid
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

	if resource != nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *SubstanceNucleicAcid) SubstanceNucleicAcidSequenceType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("sequenceType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("sequenceType", resource.SequenceType, optionsValueSet)
}
func (resource *SubstanceNucleicAcid) SubstanceNucleicAcidOligoNucleotideType(optionsValueSet []Coding) templ.Component {

	if resource != nil {
		return CodeableConceptSelect("oligoNucleotideType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("oligoNucleotideType", resource.OligoNucleotideType, optionsValueSet)
}
func (resource *SubstanceNucleicAcid) SubstanceNucleicAcidSubunitFivePrime(numSubunit int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Subunit) >= numSubunit {
		return CodeableConceptSelect("fivePrime", nil, optionsValueSet)
	}
	return CodeableConceptSelect("fivePrime", resource.Subunit[numSubunit].FivePrime, optionsValueSet)
}
func (resource *SubstanceNucleicAcid) SubstanceNucleicAcidSubunitThreePrime(numSubunit int, optionsValueSet []Coding) templ.Component {

	if resource != nil && len(resource.Subunit) >= numSubunit {
		return CodeableConceptSelect("threePrime", nil, optionsValueSet)
	}
	return CodeableConceptSelect("threePrime", resource.Subunit[numSubunit].ThreePrime, optionsValueSet)
}
