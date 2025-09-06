package r4

//generated with command go run ./bultaoreune -nodownload
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceProtein
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

// http://hl7.org/fhir/r4/StructureDefinition/SubstanceProtein
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

func (resource *SubstanceProtein) T_Id() templ.Component {

	if resource == nil {
		return StringInput("SubstanceProtein.Id", nil)
	}
	return StringInput("SubstanceProtein.Id", resource.Id)
}
func (resource *SubstanceProtein) T_ImplicitRules() templ.Component {

	if resource == nil {
		return StringInput("SubstanceProtein.ImplicitRules", nil)
	}
	return StringInput("SubstanceProtein.ImplicitRules", resource.ImplicitRules)
}
func (resource *SubstanceProtein) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("SubstanceProtein.Language", nil, optionsValueSet)
	}
	return CodeSelect("SubstanceProtein.Language", resource.Language, optionsValueSet)
}
func (resource *SubstanceProtein) T_SequenceType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("SubstanceProtein.SequenceType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("SubstanceProtein.SequenceType", resource.SequenceType, optionsValueSet)
}
func (resource *SubstanceProtein) T_NumberOfSubunits() templ.Component {

	if resource == nil {
		return IntInput("SubstanceProtein.NumberOfSubunits", nil)
	}
	return IntInput("SubstanceProtein.NumberOfSubunits", resource.NumberOfSubunits)
}
func (resource *SubstanceProtein) T_DisulfideLinkage(numDisulfideLinkage int) templ.Component {

	if resource == nil || len(resource.DisulfideLinkage) >= numDisulfideLinkage {
		return StringInput("SubstanceProtein.DisulfideLinkage["+strconv.Itoa(numDisulfideLinkage)+"]", nil)
	}
	return StringInput("SubstanceProtein.DisulfideLinkage["+strconv.Itoa(numDisulfideLinkage)+"]", &resource.DisulfideLinkage[numDisulfideLinkage])
}
func (resource *SubstanceProtein) T_SubunitId(numSubunit int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit {
		return StringInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].Id", nil)
	}
	return StringInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].Id", resource.Subunit[numSubunit].Id)
}
func (resource *SubstanceProtein) T_SubunitSubunit(numSubunit int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit {
		return IntInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].Subunit", nil)
	}
	return IntInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].Subunit", resource.Subunit[numSubunit].Subunit)
}
func (resource *SubstanceProtein) T_SubunitSequence(numSubunit int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit {
		return StringInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].Sequence", nil)
	}
	return StringInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].Sequence", resource.Subunit[numSubunit].Sequence)
}
func (resource *SubstanceProtein) T_SubunitLength(numSubunit int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit {
		return IntInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].Length", nil)
	}
	return IntInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].Length", resource.Subunit[numSubunit].Length)
}
func (resource *SubstanceProtein) T_SubunitNTerminalModification(numSubunit int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit {
		return StringInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].NTerminalModification", nil)
	}
	return StringInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].NTerminalModification", resource.Subunit[numSubunit].NTerminalModification)
}
func (resource *SubstanceProtein) T_SubunitCTerminalModification(numSubunit int) templ.Component {

	if resource == nil || len(resource.Subunit) >= numSubunit {
		return StringInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].CTerminalModification", nil)
	}
	return StringInput("SubstanceProtein.Subunit["+strconv.Itoa(numSubunit)+"].CTerminalModification", resource.Subunit[numSubunit].CTerminalModification)
}
