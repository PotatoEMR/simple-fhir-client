package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceProtein
type SubstanceProtein struct {
	Id                *string                   `json:"id,omitempty"`
	Meta              *Meta                     `json:"meta,omitempty"`
	ImplicitRules     *string                   `json:"implicitRules,omitempty"`
	Language          *string                   `json:"language,omitempty"`
	Text              *Narrative                `json:"text,omitempty"`
	Contained         []Resource                `json:"contained,omitempty"`
	Extension         []Extension               `json:"extension,omitempty"`
	ModifierExtension []Extension               `json:"modifierExtension,omitempty"`
	SequenceType      *CodeableConcept          `json:"sequenceType,omitempty"`
	NumberOfSubunits  *int                      `json:"numberOfSubunits,omitempty"`
	DisulfideLinkage  []string                  `json:"disulfideLinkage,omitempty"`
	Subunit           []SubstanceProteinSubunit `json:"subunit,omitempty"`
}

// http://hl7.org/fhir/r5/StructureDefinition/SubstanceProtein
type SubstanceProteinSubunit struct {
	Id                      *string     `json:"id,omitempty"`
	Extension               []Extension `json:"extension,omitempty"`
	ModifierExtension       []Extension `json:"modifierExtension,omitempty"`
	Subunit                 *int        `json:"subunit,omitempty"`
	Sequence                *string     `json:"sequence,omitempty"`
	Length                  *int        `json:"length,omitempty"`
	SequenceAttachment      *Attachment `json:"sequenceAttachment,omitempty"`
	NTerminalModificationId *Identifier `json:"nTerminalModificationId,omitempty"`
	NTerminalModification   *string     `json:"nTerminalModification,omitempty"`
	CTerminalModificationId *Identifier `json:"cTerminalModificationId,omitempty"`
	CTerminalModification   *string     `json:"cTerminalModification,omitempty"`
}

type OtherSubstanceProtein SubstanceProtein

// on convert struct to json, automatically add resourceType=SubstanceProtein
func (r SubstanceProtein) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceProtein
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceProtein: OtherSubstanceProtein(r),
		ResourceType:          "SubstanceProtein",
	})
}
func (r SubstanceProtein) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "SubstanceProtein/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "SubstanceProtein"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *SubstanceProtein) T_SequenceType(optionsValueSet []Coding, htmlAttrs string) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("SubstanceProtein.SequenceType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("SubstanceProtein.SequenceType", resource.SequenceType, optionsValueSet, htmlAttrs)
}
func (resource *SubstanceProtein) T_NumberOfSubunits(htmlAttrs string) templ.Component {
	if resource == nil {
		return IntInput("SubstanceProtein.NumberOfSubunits", nil, htmlAttrs)
	}
	return IntInput("SubstanceProtein.NumberOfSubunits", resource.NumberOfSubunits, htmlAttrs)
}
func (resource *SubstanceProtein) T_DisulfideLinkage(numDisulfideLinkage int, htmlAttrs string) templ.Component {
	if resource == nil || numDisulfideLinkage >= len(resource.DisulfideLinkage) {
		return StringInput("SubstanceProtein.DisulfideLinkage["+strconv.Itoa(numDisulfideLinkage)+"]", nil, htmlAttrs)
	}
	return StringInput("SubstanceProtein.DisulfideLinkage["+strconv.Itoa(numDisulfideLinkage)+"]", &resource.DisulfideLinkage[numDisulfideLinkage], htmlAttrs)
}
func (resource *SubstanceProtein) T_SubunitSubunit(numSubunit int, htmlAttrs string) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) {
		return IntInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].Subunit", nil, htmlAttrs)
	}
	return IntInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].Subunit", resource.Subunit[numSubunit].Subunit, htmlAttrs)
}
func (resource *SubstanceProtein) T_SubunitSequence(numSubunit int, htmlAttrs string) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) {
		return StringInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].Sequence", nil, htmlAttrs)
	}
	return StringInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].Sequence", resource.Subunit[numSubunit].Sequence, htmlAttrs)
}
func (resource *SubstanceProtein) T_SubunitLength(numSubunit int, htmlAttrs string) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) {
		return IntInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].Length", nil, htmlAttrs)
	}
	return IntInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].Length", resource.Subunit[numSubunit].Length, htmlAttrs)
}
func (resource *SubstanceProtein) T_SubunitNTerminalModification(numSubunit int, htmlAttrs string) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) {
		return StringInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].NTerminalModification", nil, htmlAttrs)
	}
	return StringInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].NTerminalModification", resource.Subunit[numSubunit].NTerminalModification, htmlAttrs)
}
func (resource *SubstanceProtein) T_SubunitCTerminalModification(numSubunit int, htmlAttrs string) templ.Component {
	if resource == nil || numSubunit >= len(resource.Subunit) {
		return StringInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].CTerminalModification", nil, htmlAttrs)
	}
	return StringInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].CTerminalModification", resource.Subunit[numSubunit].CTerminalModification, htmlAttrs)
}
