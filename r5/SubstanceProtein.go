package r5

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r5/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import "encoding/json"
import "github.com/a-h/templ"

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

func (resource *SubstanceProtein) T_Language(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeSelect("language", nil, optionsValueSet)
	}
	return CodeSelect("language", resource.Language, optionsValueSet)
}
func (resource *SubstanceProtein) T_SequenceType(optionsValueSet []Coding) templ.Component {

	if resource == nil {
		return CodeableConceptSelect("sequenceType", nil, optionsValueSet)
	}
	return CodeableConceptSelect("sequenceType", resource.SequenceType, optionsValueSet)
}
