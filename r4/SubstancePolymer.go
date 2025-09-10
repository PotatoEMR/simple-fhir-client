package r4

//generated with command go run ./bultaoreune
//inputs https://www.hl7.org/fhir/r4/[profiles-resources.json profiles-types.json valuesets.json]
//for details see https://github.com/PotatoEMR/simple-fhir-client

import (
	"encoding/json"
	"strconv"

	"github.com/a-h/templ"
)

// http://hl7.org/fhir/r4/StructureDefinition/SubstancePolymer
type SubstancePolymer struct {
	Id                    *string                      `json:"id,omitempty"`
	Meta                  *Meta                        `json:"meta,omitempty"`
	ImplicitRules         *string                      `json:"implicitRules,omitempty"`
	Language              *string                      `json:"language,omitempty"`
	Text                  *Narrative                   `json:"text,omitempty"`
	Contained             []Resource                   `json:"contained,omitempty"`
	Extension             []Extension                  `json:"extension,omitempty"`
	ModifierExtension     []Extension                  `json:"modifierExtension,omitempty"`
	Class                 *CodeableConcept             `json:"class,omitempty"`
	Geometry              *CodeableConcept             `json:"geometry,omitempty"`
	CopolymerConnectivity []CodeableConcept            `json:"copolymerConnectivity,omitempty"`
	Modification          []string                     `json:"modification,omitempty"`
	MonomerSet            []SubstancePolymerMonomerSet `json:"monomerSet,omitempty"`
	Repeat                []SubstancePolymerRepeat     `json:"repeat,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstancePolymer
type SubstancePolymerMonomerSet struct {
	Id                *string                                      `json:"id,omitempty"`
	Extension         []Extension                                  `json:"extension,omitempty"`
	ModifierExtension []Extension                                  `json:"modifierExtension,omitempty"`
	RatioType         *CodeableConcept                             `json:"ratioType,omitempty"`
	StartingMaterial  []SubstancePolymerMonomerSetStartingMaterial `json:"startingMaterial,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstancePolymer
type SubstancePolymerMonomerSetStartingMaterial struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Material          *CodeableConcept `json:"material,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	IsDefining        *bool            `json:"isDefining,omitempty"`
	Amount            *SubstanceAmount `json:"amount,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstancePolymer
type SubstancePolymerRepeat struct {
	Id                      *string                            `json:"id,omitempty"`
	Extension               []Extension                        `json:"extension,omitempty"`
	ModifierExtension       []Extension                        `json:"modifierExtension,omitempty"`
	NumberOfUnits           *int                               `json:"numberOfUnits,omitempty"`
	AverageMolecularFormula *string                            `json:"averageMolecularFormula,omitempty"`
	RepeatUnitAmountType    *CodeableConcept                   `json:"repeatUnitAmountType,omitempty"`
	RepeatUnit              []SubstancePolymerRepeatRepeatUnit `json:"repeatUnit,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstancePolymer
type SubstancePolymerRepeatRepeatUnit struct {
	Id                          *string                                                    `json:"id,omitempty"`
	Extension                   []Extension                                                `json:"extension,omitempty"`
	ModifierExtension           []Extension                                                `json:"modifierExtension,omitempty"`
	OrientationOfPolymerisation *CodeableConcept                                           `json:"orientationOfPolymerisation,omitempty"`
	RepeatUnit                  *string                                                    `json:"repeatUnit,omitempty"`
	Amount                      *SubstanceAmount                                           `json:"amount,omitempty"`
	DegreeOfPolymerisation      []SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation   `json:"degreeOfPolymerisation,omitempty"`
	StructuralRepresentation    []SubstancePolymerRepeatRepeatUnitStructuralRepresentation `json:"structuralRepresentation,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstancePolymer
type SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Degree            *CodeableConcept `json:"degree,omitempty"`
	Amount            *SubstanceAmount `json:"amount,omitempty"`
}

// http://hl7.org/fhir/r4/StructureDefinition/SubstancePolymer
type SubstancePolymerRepeatRepeatUnitStructuralRepresentation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Representation    *string          `json:"representation,omitempty"`
	Attachment        *Attachment      `json:"attachment,omitempty"`
}

type OtherSubstancePolymer SubstancePolymer

// on convert struct to json, automatically add resourceType=SubstancePolymer
func (r SubstancePolymer) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstancePolymer
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstancePolymer: OtherSubstancePolymer(r),
		ResourceType:          "SubstancePolymer",
	})
}
func (r SubstancePolymer) ToRef() Reference {
	var ref Reference
	if r.Id != nil {
		refStr := "SubstancePolymer/" + *r.Id
		ref.Reference = &refStr
	}

	rtype := "SubstancePolymer"
	ref.Type = &rtype
	//rDisplay := r.String()
	//ref.Display = &rDisplay
	return ref
}
func (resource *SubstancePolymer) T_Class(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("class", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("class", resource.Class, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_Geometry(optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil {
		return CodeableConceptSelect("geometry", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("geometry", resource.Geometry, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_CopolymerConnectivity(numCopolymerConnectivity int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numCopolymerConnectivity >= len(resource.CopolymerConnectivity) {
		return CodeableConceptSelect("copolymerConnectivity["+strconv.Itoa(numCopolymerConnectivity)+"]", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("copolymerConnectivity["+strconv.Itoa(numCopolymerConnectivity)+"]", &resource.CopolymerConnectivity[numCopolymerConnectivity], optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_Modification(numModification int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numModification >= len(resource.Modification) {
		return StringInput("modification["+strconv.Itoa(numModification)+"]", nil, htmlAttrs)
	}
	return StringInput("modification["+strconv.Itoa(numModification)+"]", &resource.Modification[numModification], htmlAttrs)
}
func (resource *SubstancePolymer) T_MonomerSetRatioType(numMonomerSet int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMonomerSet >= len(resource.MonomerSet) {
		return CodeableConceptSelect("monomerSet["+strconv.Itoa(numMonomerSet)+"].ratioType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("monomerSet["+strconv.Itoa(numMonomerSet)+"].ratioType", resource.MonomerSet[numMonomerSet].RatioType, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_MonomerSetStartingMaterialMaterial(numMonomerSet int, numStartingMaterial int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMonomerSet >= len(resource.MonomerSet) || numStartingMaterial >= len(resource.MonomerSet[numMonomerSet].StartingMaterial) {
		return CodeableConceptSelect("monomerSet["+strconv.Itoa(numMonomerSet)+"].startingMaterial["+strconv.Itoa(numStartingMaterial)+"].material", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("monomerSet["+strconv.Itoa(numMonomerSet)+"].startingMaterial["+strconv.Itoa(numStartingMaterial)+"].material", resource.MonomerSet[numMonomerSet].StartingMaterial[numStartingMaterial].Material, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_MonomerSetStartingMaterialType(numMonomerSet int, numStartingMaterial int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMonomerSet >= len(resource.MonomerSet) || numStartingMaterial >= len(resource.MonomerSet[numMonomerSet].StartingMaterial) {
		return CodeableConceptSelect("monomerSet["+strconv.Itoa(numMonomerSet)+"].startingMaterial["+strconv.Itoa(numStartingMaterial)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("monomerSet["+strconv.Itoa(numMonomerSet)+"].startingMaterial["+strconv.Itoa(numStartingMaterial)+"].type", resource.MonomerSet[numMonomerSet].StartingMaterial[numStartingMaterial].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_MonomerSetStartingMaterialIsDefining(numMonomerSet int, numStartingMaterial int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numMonomerSet >= len(resource.MonomerSet) || numStartingMaterial >= len(resource.MonomerSet[numMonomerSet].StartingMaterial) {
		return BoolInput("monomerSet["+strconv.Itoa(numMonomerSet)+"].startingMaterial["+strconv.Itoa(numStartingMaterial)+"].isDefining", nil, htmlAttrs)
	}
	return BoolInput("monomerSet["+strconv.Itoa(numMonomerSet)+"].startingMaterial["+strconv.Itoa(numStartingMaterial)+"].isDefining", resource.MonomerSet[numMonomerSet].StartingMaterial[numStartingMaterial].IsDefining, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatNumberOfUnits(numRepeat int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) {
		return IntInput("repeat["+strconv.Itoa(numRepeat)+"].numberOfUnits", nil, htmlAttrs)
	}
	return IntInput("repeat["+strconv.Itoa(numRepeat)+"].numberOfUnits", resource.Repeat[numRepeat].NumberOfUnits, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatAverageMolecularFormula(numRepeat int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) {
		return StringInput("repeat["+strconv.Itoa(numRepeat)+"].averageMolecularFormula", nil, htmlAttrs)
	}
	return StringInput("repeat["+strconv.Itoa(numRepeat)+"].averageMolecularFormula", resource.Repeat[numRepeat].AverageMolecularFormula, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitAmountType(numRepeat int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) {
		return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnitAmountType", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnitAmountType", resource.Repeat[numRepeat].RepeatUnitAmountType, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitOrientationOfPolymerisation(numRepeat int, numRepeatUnit int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) {
		return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].orientationOfPolymerisation", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].orientationOfPolymerisation", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].OrientationOfPolymerisation, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitRepeatUnit(numRepeat int, numRepeatUnit int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) {
		return StringInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].repeatUnit", nil, htmlAttrs)
	}
	return StringInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].repeatUnit", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].RepeatUnit, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitDegreeOfPolymerisationDegree(numRepeat int, numRepeatUnit int, numDegreeOfPolymerisation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numDegreeOfPolymerisation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation) {
		return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].degreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].degree", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].degreeOfPolymerisation["+strconv.Itoa(numDegreeOfPolymerisation)+"].degree", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].DegreeOfPolymerisation[numDegreeOfPolymerisation].Degree, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitStructuralRepresentationType(numRepeat int, numRepeatUnit int, numStructuralRepresentation int, optionsValueSet []Coding, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numStructuralRepresentation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation) {
		return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].structuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].type", nil, optionsValueSet, htmlAttrs)
	}
	return CodeableConceptSelect("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].structuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].type", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation[numStructuralRepresentation].Type, optionsValueSet, htmlAttrs)
}
func (resource *SubstancePolymer) T_RepeatRepeatUnitStructuralRepresentationRepresentation(numRepeat int, numRepeatUnit int, numStructuralRepresentation int, htmlAttrs templ.Attributes) templ.Component {
	if resource == nil || numRepeat >= len(resource.Repeat) || numRepeatUnit >= len(resource.Repeat[numRepeat].RepeatUnit) || numStructuralRepresentation >= len(resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation) {
		return StringInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].structuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].representation", nil, htmlAttrs)
	}
	return StringInput("repeat["+strconv.Itoa(numRepeat)+"].repeatUnit["+strconv.Itoa(numRepeatUnit)+"].structuralRepresentation["+strconv.Itoa(numStructuralRepresentation)+"].representation", resource.Repeat[numRepeat].RepeatUnit[numRepeatUnit].StructuralRepresentation[numStructuralRepresentation].Representation, htmlAttrs)
}
