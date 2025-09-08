package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

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
func (resource *SubstanceNucleicAcid) T_SequenceType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SequenceType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SequenceType", resource.SequenceType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_NumberOfSubunits(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("NumberOfSubunits", nil, htmlAttrs)
	}
	return IntInput("NumberOfSubunits", resource.NumberOfSubunits, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_AreaOfHybridisation(htmlAttrs string) templ.Component {
	if resource == nil {
		return StringInput("AreaOfHybridisation", nil, htmlAttrs)
	}
	return StringInput("AreaOfHybridisation", resource.AreaOfHybridisation, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_OligoNucleotideType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("OligoNucleotideType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("OligoNucleotideType", resource.OligoNucleotideType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitSubunit(numSubunit int, htmlAttrs string) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) {
		return IntInput("Subunit["+strconv.Itoa(numSubunit)+"]Subunit", nil, htmlAttrs)
	}
	return IntInput("Subunit["+strconv.Itoa(numSubunit)+"]Subunit", resource.Subunit[numSubunit].Subunit, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitSequence(numSubunit int, htmlAttrs string) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) {
		return StringInput("Subunit["+strconv.Itoa(numSubunit)+"]Sequence", nil, htmlAttrs)
	}
	return StringInput("Subunit["+strconv.Itoa(numSubunit)+"]Sequence", resource.Subunit[numSubunit].Sequence, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitLength(numSubunit int, htmlAttrs string) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) {
		return IntInput("Subunit["+strconv.Itoa(numSubunit)+"]Length", nil, htmlAttrs)
	}
	return IntInput("Subunit["+strconv.Itoa(numSubunit)+"]Length", resource.Subunit[numSubunit].Length, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitFivePrime(numSubunit int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) {
		return CodeableConceptSelect("Subunit["+strconv.Itoa(numSubunit)+"]FivePrime", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Subunit["+strconv.Itoa(numSubunit)+"]FivePrime", resource.Subunit[numSubunit].FivePrime, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitThreePrime(numSubunit int, optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) {
		return CodeableConceptSelect("Subunit["+strconv.Itoa(numSubunit)+"]ThreePrime", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("Subunit["+strconv.Itoa(numSubunit)+"]ThreePrime", resource.Subunit[numSubunit].ThreePrime, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitLinkageConnectivity(numSubunit int, numLinkage int, htmlAttrs string) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) || numLinkage >= len(resource.Subunit[numSubunit].Linkage) {
		return StringInput("Subunit["+strconv.Itoa(numSubunit)+"]Linkage["+strconv.Itoa(numLinkage)+"].Connectivity", nil, htmlAttrs)
	}
	return StringInput("Subunit["+strconv.Itoa(numSubunit)+"]Linkage["+strconv.Itoa(numLinkage)+"].Connectivity", resource.Subunit[numSubunit].Linkage[numLinkage].Connectivity, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitLinkageName(numSubunit int, numLinkage int, htmlAttrs string) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) || numLinkage >= len(resource.Subunit[numSubunit].Linkage) {
		return StringInput("Subunit["+strconv.Itoa(numSubunit)+"]Linkage["+strconv.Itoa(numLinkage)+"].Name", nil, htmlAttrs)
	}
	return StringInput("Subunit["+strconv.Itoa(numSubunit)+"]Linkage["+strconv.Itoa(numLinkage)+"].Name", resource.Subunit[numSubunit].Linkage[numLinkage].Name, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitLinkageResidueSite(numSubunit int, numLinkage int, htmlAttrs string) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) || numLinkage >= len(resource.Subunit[numSubunit].Linkage) {
		return StringInput("Subunit["+strconv.Itoa(numSubunit)+"]Linkage["+strconv.Itoa(numLinkage)+"].ResidueSite", nil, htmlAttrs)
	}
	return StringInput("Subunit["+strconv.Itoa(numSubunit)+"]Linkage["+strconv.Itoa(numLinkage)+"].ResidueSite", resource.Subunit[numSubunit].Linkage[numLinkage].ResidueSite, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitSugarName(numSubunit int, numSugar int, htmlAttrs string) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) || numSugar >= len(resource.Subunit[numSubunit].Sugar) {
		return StringInput("Subunit["+strconv.Itoa(numSubunit)+"]Sugar["+strconv.Itoa(numSugar)+"].Name", nil, htmlAttrs)
	}
	return StringInput("Subunit["+strconv.Itoa(numSubunit)+"]Sugar["+strconv.Itoa(numSugar)+"].Name", resource.Subunit[numSubunit].Sugar[numSugar].Name, htmlAttrs)
}
func (resource *SubstanceNucleicAcid) T_SubunitSugarResidueSite(numSubunit int, numSugar int, htmlAttrs string) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) || numSugar >= len(resource.Subunit[numSubunit].Sugar) {
		return StringInput("Subunit["+strconv.Itoa(numSubunit)+"]Sugar["+strconv.Itoa(numSugar)+"].ResidueSite", nil, htmlAttrs)
	}
	return StringInput("Subunit["+strconv.Itoa(numSubunit)+"]Sugar["+strconv.Itoa(numSugar)+"].ResidueSite", resource.Subunit[numSubunit].Sugar[numSugar].ResidueSite, htmlAttrs)
}
