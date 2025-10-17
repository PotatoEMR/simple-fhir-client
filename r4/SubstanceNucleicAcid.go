package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/a-h/templ"
)

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

// struct -> json, automatically add resourceType=Patient
func (r SubstanceNucleicAcid) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceNucleicAcid
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceNucleicAcid: OtherSubstanceNucleicAcid(r),
		ResourceType:              "SubstanceNucleicAcid",
	})
}

// json -> struct, first reject if resourceType != SubstanceNucleicAcid
func (r *SubstanceNucleicAcid) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &checkType); err != nil {
		return err
	} else if checkType.ResourceType != "SubstanceNucleicAcid" {
		return errors.New("resourceType not SubstanceNucleicAcid")
	}
	return json.Unmarshal(data, (*OtherSubstanceNucleicAcid)(r))
}

func (r SubstanceNucleicAcid) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "SubstanceNucleicAcid/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "SubstanceNucleicAcid"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *SubstanceNucleicAcid) T_SequenceType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("sequenceType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("sequenceType", resource.SequenceType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_NumberOfSubunits(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return IntInput("numberOfSubunits", nil, htmlAttrs)
	}
	return IntInput("numberOfSubunits", resource.NumberOfSubunits, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_AreaOfHybridisation(htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return StringInput("areaOfHybridisation", nil, htmlAttrs)
	}
	return StringInput("areaOfHybridisation", resource.AreaOfHybridisation, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_OligoNucleotideType(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("oligoNucleotideType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("oligoNucleotideType", resource.OligoNucleotideType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitSubunit(numSubunit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) {
		return IntInput("subunit["+strconv.Itoa(numSubunit)+"].subunit", nil, htmlAttrs)
	}
	return IntInput("subunit["+strconv.Itoa(numSubunit)+"].subunit", resource.Subunit[numSubunit].Subunit, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitSequence(numSubunit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) {
		return StringInput("subunit["+strconv.Itoa(numSubunit)+"].sequence", nil, htmlAttrs)
	}
	return StringInput("subunit["+strconv.Itoa(numSubunit)+"].sequence", resource.Subunit[numSubunit].Sequence, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitLength(numSubunit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) {
		return IntInput("subunit["+strconv.Itoa(numSubunit)+"].length", nil, htmlAttrs)
	}
	return IntInput("subunit["+strconv.Itoa(numSubunit)+"].length", resource.Subunit[numSubunit].Length, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitSequenceAttachment(numSubunit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) {
		return AttachmentInput("subunit["+strconv.Itoa(numSubunit)+"].sequenceAttachment", nil, htmlAttrs)
	}
	return AttachmentInput("subunit["+strconv.Itoa(numSubunit)+"].sequenceAttachment", resource.Subunit[numSubunit].SequenceAttachment, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitFivePrime(numSubunit int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) {
		return CodeableConceptSelect("subunit["+strconv.Itoa(numSubunit)+"].fivePrime", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subunit["+strconv.Itoa(numSubunit)+"].fivePrime", resource.Subunit[numSubunit].FivePrime, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitThreePrime(numSubunit int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) {
		return CodeableConceptSelect("subunit["+strconv.Itoa(numSubunit)+"].threePrime", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("subunit["+strconv.Itoa(numSubunit)+"].threePrime", resource.Subunit[numSubunit].ThreePrime, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitLinkageConnectivity(numSubunit int, numLinkage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) || numLinkage >= len(resource.Subunit[numSubunit].Linkage) {
		return StringInput("subunit["+strconv.Itoa(numSubunit)+"].linkage["+strconv.Itoa(numLinkage)+"].connectivity", nil, htmlAttrs)
	}
	return StringInput("subunit["+strconv.Itoa(numSubunit)+"].linkage["+strconv.Itoa(numLinkage)+"].connectivity", resource.Subunit[numSubunit].Linkage[numLinkage].Connectivity, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitLinkageName(numSubunit int, numLinkage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) || numLinkage >= len(resource.Subunit[numSubunit].Linkage) {
		return StringInput("subunit["+strconv.Itoa(numSubunit)+"].linkage["+strconv.Itoa(numLinkage)+"].name", nil, htmlAttrs)
	}
	return StringInput("subunit["+strconv.Itoa(numSubunit)+"].linkage["+strconv.Itoa(numLinkage)+"].name", resource.Subunit[numSubunit].Linkage[numLinkage].Name, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitLinkageResidueSite(numSubunit int, numLinkage int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) || numLinkage >= len(resource.Subunit[numSubunit].Linkage) {
		return StringInput("subunit["+strconv.Itoa(numSubunit)+"].linkage["+strconv.Itoa(numLinkage)+"].residueSite", nil, htmlAttrs)
	}
	return StringInput("subunit["+strconv.Itoa(numSubunit)+"].linkage["+strconv.Itoa(numLinkage)+"].residueSite", resource.Subunit[numSubunit].Linkage[numLinkage].ResidueSite, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitSugarName(numSubunit int, numSugar int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) || numSugar >= len(resource.Subunit[numSubunit].Sugar) {
		return StringInput("subunit["+strconv.Itoa(numSubunit)+"].sugar["+strconv.Itoa(numSugar)+"].name", nil, htmlAttrs)
	}
	return StringInput("subunit["+strconv.Itoa(numSubunit)+"].sugar["+strconv.Itoa(numSugar)+"].name", resource.Subunit[numSubunit].Sugar[numSugar].Name, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitSugarResidueSite(numSubunit int, numSugar int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) || numSugar >= len(resource.Subunit[numSubunit].Sugar) {
		return StringInput("subunit["+strconv.Itoa(numSubunit)+"].sugar["+strconv.Itoa(numSugar)+"].residueSite", nil, htmlAttrs)
	}
	return StringInput("subunit["+strconv.Itoa(numSubunit)+"].sugar["+strconv.Itoa(numSugar)+"].residueSite", resource.Subunit[numSubunit].Sugar[numSugar].ResidueSite, htmlAttrs)
}
