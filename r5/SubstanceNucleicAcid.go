package r5

//generated with command go run ./bultaoreune -nodownload
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

func (resource *SubstanceNucleicAcid) T_Id() templ.Component {

	if resource == nil {
		return StringInput("SubstanceNucleicAcid.Id", nil)
	}
	return StringInput("SubstanceNucleicAcid.Id", resource.Id)
}
func (resource *SubstanceNucleicAcid) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("SubstanceNucleicAcid.ImplicitRules", nil)
	}
	return StringInput("SubstanceNucleicAcid.ImplicitRules", resource.ImplicitRules)
}
func (resource *SubstanceNucleicAcid) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("SubstanceNucleicAcid.Language", nil, optionsValueSet)
	}
	return CodeSelect("SubstanceNucleicAcid.Language", resource.Language, optionsValueSet)
}
func (resource *SubstanceNucleicAcid) T_SequenceType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceNucleicAcid.SequenceType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceNucleicAcid.SequenceType", resource.SequenceType, optionsValueSet)
}
func (resource *SubstanceNucleicAcid) T_NumberOfSubunits() templ.Component {

	if resource == nil {
		return IntInput("SubstanceNucleicAcid.NumberOfSubunits", nil)
	}
	return IntInput("SubstanceNucleicAcid.NumberOfSubunits", resource.NumberOfSubunits)
}
func (resource *SubstanceNucleicAcid) T_AreaOfHybridisation() templ.Component {

	if resource == nil {
		return StringInput("SubstanceNucleicAcid.AreaOfHybridisation", nil)
	}
	return StringInput("SubstanceNucleicAcid.AreaOfHybridisation", resource.AreaOfHybridisation)
}
func (resource *SubstanceNucleicAcid) T_OligoNucleotideType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceNucleicAcid.OligoNucleotideType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceNucleicAcid.OligoNucleotideType", resource.OligoNucleotideType, optionsValueSet)
}
func (resource *SubstanceNucleicAcid) T_SubunitId(numSubunit int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit {
		return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Id", nil)
	}
	return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Id", resource.Subunit[numSubunit].Id)
}
func (resource *SubstanceNucleicAcid) T_SubunitSubunit(numSubunit int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit {
		return IntInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Subunit", nil)
	}
	return IntInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Subunit", resource.Subunit[numSubunit].Subunit)
}
func (resource *SubstanceNucleicAcid) T_SubunitSequence(numSubunit int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit {
		return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Sequence", nil)
	}
	return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Sequence", resource.Subunit[numSubunit].Sequence)
}
func (resource *SubstanceNucleicAcid) T_SubunitLength(numSubunit int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit {
		return IntInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Length", nil)
	}
	return IntInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Length", resource.Subunit[numSubunit].Length)
}
func (resource *SubstanceNucleicAcid) T_SubunitFivePrime(numSubunit int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit {
		return CodeableConceptSelect("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].FivePrime", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].FivePrime", resource.Subunit[numSubunit].FivePrime, optionsValueSet)
}
func (resource *SubstanceNucleicAcid) T_SubunitThreePrime(numSubunit int, optionsValueSet []Coding) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit {
		return CodeableConceptSelect("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].ThreePrime", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].ThreePrime", resource.Subunit[numSubunit].ThreePrime, optionsValueSet)
}
func (resource *SubstanceNucleicAcid) T_SubunitLinkageId(numSubunit int, numLinkage int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit || len(resource.Subunit[numSubunit].Linkage) >= numLinkage {
		return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Linkage["+strconv.Itoa(numLinkage)+"].Id", nil)
	}
	return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Linkage["+strconv.Itoa(numLinkage)+"].Id", resource.Subunit[numSubunit].Linkage[numLinkage].Id)
}
func (resource *SubstanceNucleicAcid) T_SubunitLinkageConnectivity(numSubunit int, numLinkage int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit || len(resource.Subunit[numSubunit].Linkage) >= numLinkage {
		return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Linkage["+strconv.Itoa(numLinkage)+"].Connectivity", nil)
	}
	return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Linkage["+strconv.Itoa(numLinkage)+"].Connectivity", resource.Subunit[numSubunit].Linkage[numLinkage].Connectivity)
}
func (resource *SubstanceNucleicAcid) T_SubunitLinkageName(numSubunit int, numLinkage int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit || len(resource.Subunit[numSubunit].Linkage) >= numLinkage {
		return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Linkage["+strconv.Itoa(numLinkage)+"].Name", nil)
	}
	return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Linkage["+strconv.Itoa(numLinkage)+"].Name", resource.Subunit[numSubunit].Linkage[numLinkage].Name)
}
func (resource *SubstanceNucleicAcid) T_SubunitLinkageResidueSite(numSubunit int, numLinkage int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit || len(resource.Subunit[numSubunit].Linkage) >= numLinkage {
		return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Linkage["+strconv.Itoa(numLinkage)+"].ResidueSite", nil)
	}
	return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Linkage["+strconv.Itoa(numLinkage)+"].ResidueSite", resource.Subunit[numSubunit].Linkage[numLinkage].ResidueSite)
}
func (resource *SubstanceNucleicAcid) T_SubunitSugarId(numSubunit int, numSugar int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit || len(resource.Subunit[numSubunit].Sugar) >= numSugar {
		return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Sugar["+strconv.Itoa(numSugar)+"].Id", nil)
	}
	return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Sugar["+strconv.Itoa(numSugar)+"].Id", resource.Subunit[numSubunit].Sugar[numSugar].Id)
}
func (resource *SubstanceNucleicAcid) T_SubunitSugarName(numSubunit int, numSugar int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit || len(resource.Subunit[numSubunit].Sugar) >= numSugar {
		return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Sugar["+strconv.Itoa(numSugar)+"].Name", nil)
	}
	return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Sugar["+strconv.Itoa(numSugar)+"].Name", resource.Subunit[numSubunit].Sugar[numSugar].Name)
}
func (resource *SubstanceNucleicAcid) T_SubunitSugarResidueSite(numSubunit int, numSugar int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit || len(resource.Subunit[numSubunit].Sugar) >= numSugar {
		return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Sugar["+strconv.Itoa(numSugar)+"].ResidueSite", nil)
	}
	return StringInput("SubstanceNucleicAcid.Subunit["+strconv.Itoa(numSubunit)+"].Sugar["+strconv.Itoa(numSugar)+"].ResidueSite", resource.Subunit[numSubunit].Sugar[numSugar].ResidueSite)
}
